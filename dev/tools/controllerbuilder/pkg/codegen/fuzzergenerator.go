// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codegen

// This file implements the deterministic fuzzer generator (generate-fuzzer --deterministic).
//
// For one KRM kind it pairs every field of the resource's proto message with the kind's
// Spec and ObservedState structs, using the rules the mapper generator uses, and writes
// <kind>_fuzzer.generated.go with the SpecField, StatusField and Unimplemented_* entries
// that make both FuzzSpec and FuzzStatus pass.
//
// How entries are chosen
//
// FuzzSpec fills a random proto, clears every path listed as a Status field or as
// unimplemented, and requires the rest to round-trip through the Spec mappers. FuzzStatus
// does the same through the ObservedState mappers, clearing Spec fields and unimplemented
// fields. Clearing a path clears everything below it.
//
// So FuzzSpec has to clear every field the Spec side doesn't map, and FuzzStatus every field
// the Status side doesn't map. That includes message fields: a message that is still set,
// but empty, after its fields were cleared has to round-trip too, so a side that doesn't
// map the message field must clear it as a whole.
//
// The generator emits the fewest entries that satisfy this:
//   - A field whose leaves are all mapped only by the Spec side is one SpecField entry. Only
//     by the Status side, one StatusField entry. By neither side, one Unimplemented_* entry.
//   - Otherwise the message field is split into its fields. If one side doesn't map the
//     message field itself, it is first listed for the other side, so the side that
//     doesn't map it clears it as a whole. For example SpecField(".a") together with
//     Unimplemented_NotYetTriaged(".a.b").
//   - A leaf that both sides map is listed as a Spec field. FuzzSpec round-trips it through
//     the Spec mappers and FuzzStatus clears it, so both pass, and the field still appears
//     in the generated file.

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

const (
	// fuzzerMaxDepth is the deepest field path, in path segments, that the generator splits.
	// fuzz.FillWithRandom doesn't fill messages nested more than 10 levels below the resource
	// message, so their fields are never set.
	fuzzerMaxDepth = 11

	fuzztestingGoPackage = "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"

	// FuzzerFileAnnotationKey marks files written by the deterministic fuzzer generator.
	FuzzerFileAnnotationKey = "+generated:fuzzer"
)

// FuzzerMethod is the fuzztesting.KRMTypedFuzzer method that a generated entry calls.
type FuzzerMethod string

const (
	FuzzerSpecField                      FuzzerMethod = "SpecField"
	FuzzerStatusField                    FuzzerMethod = "StatusField"
	FuzzerUnimplementedIdentity          FuzzerMethod = "Unimplemented_Identity"
	FuzzerUnimplementedEtag              FuzzerMethod = "Unimplemented_Etag"
	FuzzerUnimplementedLabelsAnnotations FuzzerMethod = "Unimplemented_LabelsAnnotations"
	FuzzerUnimplementedNotYetTriaged     FuzzerMethod = "Unimplemented_NotYetTriaged"
	fuzzerIdentityPath                                = ".name"
	fuzzerEtagPath                                    = ".etag"
)

// FuzzerEntry is one field path listed in a generated fuzzer.
type FuzzerEntry struct {
	Method FuzzerMethod
	Path   string
}

// ClearedInStatus reports whether FuzzStatus clears the path, because it is a Spec field or unimplemented.
func (e FuzzerEntry) ClearedInStatus() bool {
	return e.Method != FuzzerStatusField
}

// ClearedInSpec reports whether FuzzSpec clears the path, because it is a Status field or unimplemented.
func (e FuzzerEntry) ClearedInSpec() bool {
	return e.Method != FuzzerSpecField
}

// fuzzerEntryGroups orders the entries of a generated fuzzer, and gives each group its comment.
var fuzzerEntryGroups = []struct {
	comment string
	methods []FuzzerMethod
}{
	{"Spec fields", []FuzzerMethod{FuzzerSpecField}},
	{"Status fields", []FuzzerMethod{FuzzerStatusField}},
	{"Identity / Special fields", []FuzzerMethod{FuzzerUnimplementedIdentity, FuzzerUnimplementedEtag, FuzzerUnimplementedLabelsAnnotations}},
	{"Unimplemented / Not Yet Triaged fields", []FuzzerMethod{FuzzerUnimplementedNotYetTriaged}},
}

func fuzzerMethodOrder(m FuzzerMethod) int {
	i := 0
	for _, g := range fuzzerEntryGroups {
		for _, method := range g.methods {
			if method == m {
				return i
			}
			i++
		}
	}
	return i
}

func sortFuzzerEntries(entries []FuzzerEntry) {
	sort.Slice(entries, func(i, j int) bool {
		oi, oj := fuzzerMethodOrder(entries[i].Method), fuzzerMethodOrder(entries[j].Method)
		if oi != oj {
			return oi < oj
		}
		return entries[i].Path < entries[j].Path
	})
}

// fuzzerCoverage says how one side (Spec or Status) maps a proto field.
type fuzzerCoverage int

const (
	// fuzzerNotCovered means the side has no KRM field for the proto field.
	fuzzerNotCovered fuzzerCoverage = iota
	// fuzzerCoveredOpaque means the side maps the proto field as one value: a scalar, an enum, a
	// map, a reference, a well-known type, or a message whose KRM type isn't a struct of the KRM
	// package. Its sub-fields, if any, round-trip with it.
	fuzzerCoveredOpaque
	// fuzzerCoveredStruct means the side maps the proto message field to a KRM struct. The
	// message's fields are paired with the struct's fields.
	fuzzerCoveredStruct
)

// fuzzerSide is how one side maps a proto field, and the KRM struct when there is one.
type fuzzerSide struct {
	coverage fuzzerCoverage
	goStruct *gocode.GoStruct
	// pkg declares goStruct; the types of goStruct's fields are looked up there.
	pkg *fuzzerKRMPackage
}

func (s fuzzerSide) covered() bool {
	return s.coverage != fuzzerNotCovered
}

// fuzzerClass is a set of leaf classes, as bits.
type fuzzerClass uint8

const (
	// fuzzerClassSpec is a leaf that the Spec side maps, whether or not the Status side does too.
	fuzzerClassSpec fuzzerClass = 1 << iota
	// fuzzerClassStatus is a leaf that only the Status side maps.
	fuzzerClassStatus
	// fuzzerClassUnimplemented is a leaf that neither side maps.
	fuzzerClassUnimplemented
)

// leafFuzzerClass classifies a leaf. A leaf that both sides map is a Spec field: FuzzSpec
// round-trips it through the Spec mappers and FuzzStatus clears it, so both pass.
func leafFuzzerClass(spec, status fuzzerSide) fuzzerClass {
	switch {
	case spec.covered():
		return fuzzerClassSpec
	case status.covered():
		return fuzzerClassStatus
	default:
		return fuzzerClassUnimplemented
	}
}

// fuzzerNode is a proto field, as reached from the resource message.
type fuzzerNode struct {
	// path is the field path in the fuzz.Visit grammar, for example ".a.b" or ".a[].b".
	path     string
	field    protoreflect.FieldDescriptor
	spec     fuzzerSide
	status   fuzzerSide
	children []*fuzzerNode
	// classes holds the classes of all the leaves at or below this field.
	classes fuzzerClass
}

// fuzzerKRMPackage is a loaded KRM Go package.
type fuzzerKRMPackage struct {
	importPath string
	structs    map[string]*gocode.GoStruct
	// protoFields maps a struct name, then the full name of a proto field, to the struct field
	// whose +kcc:proto:field annotation names that proto field.
	protoFields map[string]map[string]string
}

func newFuzzerKRMPackage(pkg *gocode.Package, protoFields map[string]map[string]string) *fuzzerKRMPackage {
	out := &fuzzerKRMPackage{
		importPath:  pkg.GoPackage,
		structs:     make(map[string]*gocode.GoStruct),
		protoFields: protoFields,
	}
	for _, s := range pkg.Structs {
		out.structs[s.Name] = s
	}
	return out
}

// loadProtoFieldAnnotations reads the +kcc:proto:field annotations of the struct fields
// declared in the Go files of dir, the files that gocode.LoadPackage reads. It returns struct
// name -> proto field full name -> Go field name.
func loadProtoFieldAnnotations(dir string) (map[string]map[string]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading directory %q: %w", dir, err)
	}
	out := make(map[string]map[string]string)
	fset := token.NewFileSet()
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		p := filepath.Join(dir, name)
		src, err := os.ReadFile(p)
		if err != nil {
			return nil, fmt.Errorf("reading %q: %w", p, err)
		}
		if bytes.Contains(src, []byte("//go:build !ignore_autogenerated")) {
			continue
		}
		file, err := parser.ParseFile(fset, p, src, parser.ParseComments|parser.SkipObjectResolution)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %w", p, err)
		}
		for _, decl := range file.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.TYPE {
				continue
			}
			for _, spec := range gen.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}
				for _, field := range structType.Fields.List {
					if field.Doc == nil || len(field.Names) != 1 {
						continue
					}
					for _, c := range field.Doc.List {
						line := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
						protoField, ok := strings.CutPrefix(line, KCCProtoFieldAnnotation+"=")
						if !ok {
							continue
						}
						fields := out[typeSpec.Name.Name]
						if fields == nil {
							fields = make(map[string]string)
							out[typeSpec.Name.Name] = fields
						}
						fields[strings.TrimSpace(protoField)] = field.Names[0].Name
					}
				}
			}
		}
	}
	return out, nil
}

// fuzzerKRMLoader loads KRM Go packages by import path, and caches them.
type fuzzerKRMLoader struct {
	// apiGoPackagePath is the import path of the directory that holds the KRM packages,
	// for example github.com/GoogleCloudPlatform/k8s-config-connector/apis.
	apiGoPackagePath string
	// apiDirectory is the directory that holds the KRM packages.
	apiDirectory string
	packages     map[string]*fuzzerKRMPackage
}

func newFuzzerKRMLoader(apiGoPackagePath, apiDirectory string) *fuzzerKRMLoader {
	return &fuzzerKRMLoader{
		apiGoPackagePath: strings.TrimSuffix(apiGoPackagePath, "/"),
		apiDirectory:     apiDirectory,
		packages:         make(map[string]*fuzzerKRMPackage),
	}
}

func (l *fuzzerKRMLoader) load(importPath string) (*fuzzerKRMPackage, error) {
	if pkg, ok := l.packages[importPath]; ok {
		return pkg, nil
	}
	rel, ok := strings.CutPrefix(importPath, l.apiGoPackagePath+"/")
	if !ok {
		return nil, fmt.Errorf("KRM package %q is not under %q", importPath, l.apiGoPackagePath)
	}
	dir := filepath.Join(l.apiDirectory, filepath.FromSlash(rel))
	pkg, err := gocode.LoadPackage(importPath, dir)
	if err != nil {
		return nil, fmt.Errorf("loading KRM package %q: %w", importPath, err)
	}
	if pkg == nil {
		return nil, fmt.Errorf("loading KRM package %q: no Go files in %s", importPath, dir)
	}
	protoFields, err := loadProtoFieldAnnotations(dir)
	if err != nil {
		return nil, fmt.Errorf("loading KRM package %q: %w", importPath, err)
	}
	out := newFuzzerKRMPackage(pkg, protoFields)
	l.packages[importPath] = out
	return out, nil
}

// isOpaqueFuzzerMessage reports whether a message field is always a leaf: well-known types,
// the messages mapped to KRM scalars, and messages without fields.
func isOpaqueFuzzerMessage(msg protoreflect.MessageDescriptor) bool {
	if _, ok := protoMessagesNotMappedToGoStruct[string(msg.FullName())]; ok {
		return true
	}
	if strings.HasPrefix(string(msg.FullName()), "google.protobuf.") {
		return true
	}
	return msg.Fields().Len() == 0
}

func findStructField(s *gocode.GoStruct, name string) *gocode.StructField {
	for _, f := range s.Fields {
		if f.Name == name {
			return f
		}
	}
	return nil
}

// pairFuzzerField returns how one side maps proto field fd, given how it maps the message
// that contains fd. It follows MapperGenerator.writeMapFunctionsForPair: the KRM field is
// named goFieldName(fd), or <name>Ref for a reference, or <name without a trailing s>Refs for
// a list of references. Failing that, the KRM field is the one whose +kcc:proto:field
// annotation names fd: generate-types writes the annotation, and it stays when a field is
// renamed, for example into a reference or to spell an acronym differently. Such fields are
// mapped by hand-written mappers.
func pairFuzzerField(fd protoreflect.FieldDescriptor, parent fuzzerSide) fuzzerSide {
	switch parent.coverage {
	case fuzzerNotCovered:
		return fuzzerSide{}
	case fuzzerCoveredOpaque:
		return fuzzerSide{coverage: fuzzerCoveredOpaque}
	}

	krmFieldName := goFieldName(fd)
	krmField := findStructField(parent.goStruct, krmFieldName)
	if krmField == nil {
		if findStructField(parent.goStruct, krmFieldName+"Ref") != nil {
			return fuzzerSide{coverage: fuzzerCoveredOpaque}
		}
		if findStructField(parent.goStruct, strings.TrimSuffix(krmFieldName, "s")+"Refs") != nil {
			return fuzzerSide{coverage: fuzzerCoveredOpaque}
		}
		annotated := parent.pkg.protoFields[parent.goStruct.Name][string(fd.FullName())]
		if annotated == "" {
			return fuzzerSide{}
		}
		if krmField = findStructField(parent.goStruct, annotated); krmField == nil {
			return fuzzerSide{}
		}
		if strings.HasSuffix(krmField.Name, "Ref") || strings.HasSuffix(krmField.Name, "Refs") {
			return fuzzerSide{coverage: fuzzerCoveredOpaque}
		}
	}

	isKRMSlice := strings.HasPrefix(krmField.Type, "[]") && krmField.Type != "[]byte"
	if fd.IsList() && !isKRMSlice {
		// The generated mapper keeps only the first element, so longer lists don't round-trip.
		return fuzzerSide{}
	}
	if fd.Kind() != protoreflect.MessageKind || fd.IsMap() {
		return fuzzerSide{coverage: fuzzerCoveredOpaque}
	}
	// A message whose KRM type is a struct of the KRM package is paired field by field, even a
	// well-known type: aiplatform, for example, maps google.protobuf.Value to its own Value
	// struct. Other messages are leaves: the well-known types mapped to strings or JSON, and
	// the types of other packages, such as references. Maps never match a struct name.
	typeName := strings.TrimLeft(krmField.Type, "[]*")
	if st := parent.pkg.structs[typeName]; st != nil && fd.Message().Fields().Len() > 0 {
		return fuzzerSide{coverage: fuzzerCoveredStruct, goStruct: st, pkg: parent.pkg}
	}
	return fuzzerSide{coverage: fuzzerCoveredOpaque}
}

// fuzzerPathKey identifies a message and the structs it is paired with, to stop recursing
// through recursive types.
type fuzzerPathKey struct {
	message protoreflect.FullName
	spec    *gocode.GoStruct
	status  *gocode.GoStruct
}

// buildFuzzerNodes pairs the fields of msg with the two sides, and recurses into message
// fields that at least one side maps to a struct.
func buildFuzzerNodes(prefix string, msg protoreflect.MessageDescriptor, spec, status fuzzerSide, depth int, onPath map[fuzzerPathKey]bool) []*fuzzerNode {
	var nodes []*fuzzerNode
	fields := msg.Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		n := &fuzzerNode{
			path:   prefix + "." + string(fd.Name()),
			field:  fd,
			spec:   pairFuzzerField(fd, spec),
			status: pairFuzzerField(fd, status),
		}

		expand := depth < fuzzerMaxDepth &&
			(n.spec.coverage == fuzzerCoveredStruct || n.status.coverage == fuzzerCoveredStruct)
		if expand {
			key := fuzzerPathKey{message: fd.Message().FullName(), spec: n.spec.goStruct, status: n.status.goStruct}
			if onPath[key] {
				// A KRM type that contains itself. Its fields repeat below this field, down to
				// the depth that fuzz.FillWithRandom fills, and some of them may be unmapped.
				// Rather than list them again at every level, neither side keeps the field.
				n.spec, n.status = fuzzerSide{}, fuzzerSide{}
			} else {
				onPath[key] = true
				childPrefix := n.path
				if fd.IsList() {
					childPrefix += "[]"
				}
				n.children = buildFuzzerNodes(childPrefix, fd.Message(), n.spec, n.status, depth+1, onPath)
				delete(onPath, key)
			}
		}

		if len(n.children) == 0 {
			n.classes = leafFuzzerClass(n.spec, n.status)
		} else {
			for _, child := range n.children {
				n.classes |= child.classes
			}
		}
		nodes = append(nodes, n)
	}
	return nodes
}

// emitFuzzerEntries appends the entries for the field and the fields below it. clearedInSpec
// and clearedInStatus say whether an entry for a parent field already clears this field in
// FuzzSpec or in FuzzStatus.
func emitFuzzerEntries(n *fuzzerNode, clearedInSpec, clearedInStatus, topLevel bool, out *[]FuzzerEntry) {
	if clearedInSpec && clearedInStatus {
		return
	}

	switch n.classes {
	case fuzzerClassSpec:
		if !clearedInStatus {
			*out = append(*out, FuzzerEntry{Method: FuzzerSpecField, Path: n.path})
		}
		return
	case fuzzerClassStatus:
		if !clearedInSpec {
			*out = append(*out, FuzzerEntry{Method: FuzzerStatusField, Path: n.path})
		}
		return
	case fuzzerClassUnimplemented:
		*out = append(*out, unimplementedFuzzerEntry(n, topLevel))
		return
	}

	// The leaves below this message field are mixed, so the field is split. A side that
	// doesn't map the message field has to clear it as a whole.
	if !n.spec.covered() && !clearedInSpec {
		*out = append(*out, FuzzerEntry{Method: FuzzerStatusField, Path: n.path})
		clearedInSpec = true
	}
	if !n.status.covered() && !clearedInStatus {
		*out = append(*out, FuzzerEntry{Method: FuzzerSpecField, Path: n.path})
		clearedInStatus = true
	}
	for _, child := range n.children {
		emitFuzzerEntries(child, clearedInSpec, clearedInStatus, false, out)
	}
}

// unimplementedFuzzerEntry lists a field that neither side maps. The resource name, the etag,
// and the labels and annotations maps of the resource message get their own methods.
func unimplementedFuzzerEntry(n *fuzzerNode, topLevel bool) FuzzerEntry {
	if topLevel {
		switch n.path {
		case fuzzerIdentityPath:
			return FuzzerEntry{Method: FuzzerUnimplementedIdentity, Path: n.path}
		case fuzzerEtagPath:
			return FuzzerEntry{Method: FuzzerUnimplementedEtag, Path: n.path}
		case ".labels", ".annotations":
			if n.field.IsMap() && n.field.MapKey().Kind() == protoreflect.StringKind && n.field.MapValue().Kind() == protoreflect.StringKind {
				return FuzzerEntry{Method: FuzzerUnimplementedLabelsAnnotations, Path: n.path}
			}
		}
	}
	return FuzzerEntry{Method: FuzzerUnimplementedNotYetTriaged, Path: n.path}
}

// classifyFuzzerFields returns the entries for msg, given the structs that the Spec and
// Status mappers convert it to. A zero fuzzerSide means that side maps nothing.
func classifyFuzzerFields(msg protoreflect.MessageDescriptor, spec, status fuzzerSide) []FuzzerEntry {
	onPath := map[fuzzerPathKey]bool{
		{message: msg.FullName(), spec: spec.goStruct, status: status.goStruct}: true,
	}
	var entries []FuzzerEntry
	for _, n := range buildFuzzerNodes("", msg, spec, status, 1, onPath) {
		emitFuzzerEntries(n, false, false, true, &entries)
	}
	sortFuzzerEntries(entries)
	return entries
}

// goPackageIndex holds the parsed Go files of a direct controller package.
type goPackageIndex struct {
	dir         string
	packageName string
	files       []*goIndexedFile
	funcs       map[string]*goIndexedFunc
}

type goIndexedFile struct {
	path string
	// generated is set for files whose name ends in generated.go, like findFuncDeclaration.
	generated bool
	syntax    *ast.File
	// imports maps each import's name (its alias, or the last element of its path) to its path.
	imports map[string]string
}

type goIndexedFunc struct {
	decl *ast.FuncDecl
	file *goIndexedFile
}

// loadGoPackageIndex parses the non-test Go files of dir. A missing directory gives an empty index.
func loadGoPackageIndex(dir string) (*goPackageIndex, error) {
	idx := &goPackageIndex{dir: dir, funcs: make(map[string]*goIndexedFunc)}
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return idx, nil
		}
		return nil, fmt.Errorf("reading directory %q: %w", dir, err)
	}
	fset := token.NewFileSet()
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		p := filepath.Join(dir, name)
		syntax, err := parser.ParseFile(fset, p, nil, parser.SkipObjectResolution)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %w", p, err)
		}
		file := &goIndexedFile{
			path:      p,
			generated: strings.HasSuffix(name, "generated.go"),
			syntax:    syntax,
			imports:   make(map[string]string),
		}
		for _, imp := range syntax.Imports {
			importPath, err := strconv.Unquote(imp.Path.Value)
			if err != nil {
				return nil, fmt.Errorf("parsing import %s in %q: %w", imp.Path.Value, p, err)
			}
			importName := path.Base(importPath)
			if imp.Name != nil {
				importName = imp.Name.Name
			}
			file.imports[importName] = importPath
		}
		if idx.packageName == "" {
			idx.packageName = syntax.Name.Name
		}
		for _, decl := range syntax.Decls {
			if fd, ok := decl.(*ast.FuncDecl); ok && fd.Recv == nil {
				if _, found := idx.funcs[fd.Name.Name]; !found {
					idx.funcs[fd.Name.Name] = &goIndexedFunc{decl: fd, file: file}
				}
			}
		}
		idx.files = append(idx.files, file)
	}
	return idx, nil
}

// handwrittenFunc returns the non-generated file that declares function name, or "".
func (idx *goPackageIndex) handwrittenFunc(name string) string {
	if f := idx.funcs[name]; f != nil && !f.file.generated {
		return f.file.path
	}
	for _, file := range idx.files {
		if file.generated {
			continue
		}
		for _, decl := range file.syntax.Decls {
			if fd, ok := decl.(*ast.FuncDecl); ok && fd.Recv == nil && fd.Name.Name == name {
				return file.path
			}
		}
	}
	return ""
}

// handwrittenFuzzerUsing returns a non-generated file that imports fuzztesting and refers to
// the function name, such as a fuzzer written by hand or by an LLM, or "".
func (idx *goPackageIndex) handwrittenFuzzerUsing(name string) string {
	for _, file := range idx.files {
		if file.generated {
			continue
		}
		importsFuzztesting := false
		for _, importPath := range file.imports {
			if importPath == fuzztestingGoPackage {
				importsFuzztesting = true
			}
		}
		if !importsFuzztesting {
			continue
		}
		found := false
		ast.Inspect(file.syntax, func(n ast.Node) bool {
			if id, ok := n.(*ast.Ident); ok && id.Name == name {
				found = true
			}
			return !found
		})
		if found {
			return file.path
		}
	}
	return ""
}

// handwrittenFuzzerBelow is handwrittenFuzzerUsing for the packages in the subdirectories of
// dir. For example the fuzzer of automl/automldataset uses the mappers of the automl package.
func handwrittenFuzzerBelow(dir, name string) (string, error) {
	found := ""
	err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() || p == dir {
			return nil
		}
		idx, err := loadGoPackageIndex(p)
		if err != nil {
			return err
		}
		if found = idx.handwrittenFuzzerUsing(name); found != "" {
			return filepath.SkipAll
		}
		return nil
	})
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return "", fmt.Errorf("looking for fuzzers below %q: %w", dir, err)
	}
	return found, nil
}

// fuzzerMapperPair is a <Type>_FromProto / <Type>_ToProto pair found in the direct package.
type fuzzerMapperPair struct {
	FromProto string
	ToProto   string
	// KRMType and KRMGoPackage are the struct that FromProto returns.
	KRMType      string
	KRMGoPackage string
	// ProtoType and ProtoGoPackage are the Go type of the proto message that FromProto takes.
	ProtoType      string
	ProtoGoPackage string
	// Handwritten is set when FromProto or ToProto is declared in a file that isn't generated.
	Handwritten bool
}

// mapperPairBases returns the names, without the _FromProto and _ToProto suffixes, that the
// mappers of the given KRM types can have, in order of preference: <typeName>_<version>, as
// generate-mapper --multiversion names them, then <typeName>.
func mapperPairBases(version string, typeNames ...string) []string {
	var bases []string
	for _, typeName := range typeNames {
		if version != "" {
			bases = append(bases, typeName+"_"+version)
		}
		bases = append(bases, typeName)
	}
	return bases
}

func fieldListArity(fl *ast.FieldList) int {
	if fl == nil {
		return 0
	}
	n := 0
	for _, f := range fl.List {
		if len(f.Names) == 0 {
			n++
		} else {
			n += len(f.Names)
		}
	}
	return n
}

// mapperPair returns <base>_FromProto and <base>_ToProto, or nil when the package doesn't
// declare both.
func (idx *goPackageIndex) mapperPair(base string) (*fuzzerMapperPair, error) {
	from, to := idx.funcs[base+"_FromProto"], idx.funcs[base+"_ToProto"]
	if from == nil || to == nil {
		return nil, nil
	}
	pair := &fuzzerMapperPair{
		FromProto:   base + "_FromProto",
		ToProto:     base + "_ToProto",
		Handwritten: !from.file.generated || !to.file.generated,
	}
	ft := from.decl.Type
	if fieldListArity(ft.Params) != 2 || fieldListArity(ft.Results) != 1 {
		return nil, fmt.Errorf("%s in %s: want func(*direct.MapContext, *<proto>) *<KRM type>", pair.FromProto, from.file.path)
	}
	var ok bool
	params := ft.Params.List
	if pair.ProtoGoPackage, pair.ProtoType, ok = from.file.typeRef(params[len(params)-1].Type); !ok {
		return nil, fmt.Errorf("%s in %s: cannot tell the Go type of its proto parameter", pair.FromProto, from.file.path)
	}
	if pair.KRMGoPackage, pair.KRMType, ok = from.file.typeRef(ft.Results.List[0].Type); !ok {
		return nil, fmt.Errorf("%s in %s: cannot tell the KRM type it returns", pair.FromProto, from.file.path)
	}
	toFT := to.decl.Type
	if fieldListArity(toFT.Params) != 2 || fieldListArity(toFT.Results) != 1 {
		return nil, fmt.Errorf("%s in %s: want func(*direct.MapContext, *%s) *%s", pair.ToProto, to.file.path, pair.KRMType, pair.ProtoType)
	}
	toParams := toFT.Params.List
	if _, krmType, ok := to.file.typeRef(toParams[len(toParams)-1].Type); !ok || krmType != pair.KRMType {
		return nil, fmt.Errorf("%s in %s: want *%s as its second parameter", pair.ToProto, to.file.path, pair.KRMType)
	}
	if _, protoType, ok := to.file.typeRef(toFT.Results.List[0].Type); !ok || protoType != pair.ProtoType {
		return nil, fmt.Errorf("%s in %s: want *%s as its result", pair.ToProto, to.file.path, pair.ProtoType)
	}
	return pair, nil
}

// findMapperPair returns the mappers that a fuzzer uses for one side, choosing among the
// pairs named by mapperPairBases. Generated and hand-written functions both count, but pairs
// that convert a proto type other than protoType, that return a type of another KRM version,
// or whose signature isn't that of a mapper, are skipped and described in skipped. A
// hand-written pair wins over a generated one, because it is the one that the controller
// uses; otherwise the order of mapperPairBases decides. It returns nil when no pair is left.
func (idx *goPackageIndex) findMapperPair(protoType, version string, typeNames ...string) (pair *fuzzerMapperPair, skipped []string) {
	for _, base := range mapperPairBases(version, typeNames...) {
		candidate, err := idx.mapperPair(base)
		if err != nil {
			skipped = append(skipped, err.Error())
			continue
		}
		if candidate == nil {
			continue
		}
		if candidate.ProtoType != protoType {
			skipped = append(skipped, fmt.Sprintf("%s converts *%s, not *%s", candidate.FromProto, candidate.ProtoType, protoType))
			continue
		}
		if version != "" && candidate.KRMGoPackage != "" && path.Base(candidate.KRMGoPackage) != version {
			skipped = append(skipped, fmt.Sprintf("%s returns a type of %s, not of version %s", candidate.FromProto, candidate.KRMGoPackage, version))
			continue
		}
		if pair == nil || (candidate.Handwritten && !pair.Handwritten) {
			pair = candidate
		}
	}
	return pair, skipped
}

// typeRef resolves a pointer type expression such as *krm.FooSpec to its import path and type name.
// The import path is empty for a type of the same package.
func (f *goIndexedFile) typeRef(expr ast.Expr) (goPackage, name string, ok bool) {
	star, isStar := expr.(*ast.StarExpr)
	if !isStar {
		return "", "", false
	}
	switch t := star.X.(type) {
	case *ast.Ident:
		return "", t.Name, true
	case *ast.SelectorExpr:
		x, isIdent := t.X.(*ast.Ident)
		if !isIdent {
			return "", "", false
		}
		importPath, found := f.imports[x.Name]
		if !found {
			return "", "", false
		}
		return importPath, t.Sel.Name, true
	}
	return "", "", false
}

// FuzzerPlan is the fuzzer the generator writes for one kind.
type FuzzerPlan struct {
	Kind    string
	Message protoreflect.MessageDescriptor
	// PackageName is the Go package of the direct controller.
	PackageName string
	// FuncName is the function that builds the fuzzer, <lowerCamelKind>Fuzzer.
	FuncName string
	Spec     *fuzzerMapperPair
	// Status is nil when the package has no ObservedState or Status mappers for the kind; the
	// fuzzer is then spec-only.
	Status  *fuzzerMapperPair
	Entries []FuzzerEntry
	// FilterSpec and FilterStatus are the hand-written hooks <FuncName>FilterSpec and
	// <FuncName>FilterStatus, when the package declares them.
	FilterSpec   string
	FilterStatus string

	// specSide and statusSide are the KRM structs the entries were classified against.
	specSide, statusSide fuzzerSide
}

type fuzzerPlanInput struct {
	kind string
	// version is the KRM version, used to pick <Type>_<version>_FromProto in packages whose
	// mappers are generated with --multiversion.
	version string
	message protoreflect.MessageDescriptor
	direct  *goPackageIndex
	krm     *fuzzerKRMLoader
}

// planFuzzer finds the kind's mappers and the structs they convert to, and classifies the
// fields of the resource message.
func planFuzzer(in fuzzerPlanInput) (*FuzzerPlan, error) {
	protoType := protoNameForType(in.message)
	spec, skipped := in.direct.findMapperPair(protoType, in.version, in.kind+"Spec")
	if spec == nil {
		msg := fmt.Sprintf("no Spec mappers for kind %s in %s: want %sSpec_%s_FromProto and _ToProto, or %sSpec_FromProto and _ToProto, converting *%s",
			in.kind, in.direct.dir, in.kind, in.version, in.kind, protoType)
		if len(skipped) > 0 {
			msg += "; " + strings.Join(skipped, "; ")
		}
		return nil, errors.New(msg)
	}
	if spec.ProtoGoPackage == "" {
		return nil, fmt.Errorf("%s: cannot tell the Go package of *%s", spec.FromProto, spec.ProtoType)
	}

	specSide, err := in.krmSide(spec)
	if err != nil {
		return nil, err
	}

	statusTypeNames := []string{in.kind + "ObservedState"}
	if st := specSide.pkg.structs[strings.TrimSuffix(spec.KRMType, "Spec")+"Status"]; st != nil {
		if f := findStructField(st, "ObservedState"); f != nil {
			if name := strings.TrimPrefix(f.Type, "*"); name != "" && name != statusTypeNames[0] {
				statusTypeNames = append(statusTypeNames, name)
			}
		}
	}
	statusTypeNames = append(statusTypeNames, in.kind+"Status")

	// Kinds that moved from Terraform or DCL map the proto to their Status struct, or to the
	// struct of its ObservedState field when it isn't named <Kind>ObservedState.
	status, _ := in.direct.findMapperPair(protoType, in.version, statusTypeNames...)

	var statusSide fuzzerSide
	if status != nil {
		if statusSide, err = in.krmSide(status); err != nil {
			return nil, err
		}
	} else {
		// Without ObservedState mappers the fuzzer is spec-only, and FuzzStatus never runs.
		// The ObservedState struct, or the Status struct, still tells which fields are status
		// fields.
		for _, name := range statusTypeNames {
			if st := specSide.pkg.structs[name]; st != nil {
				statusSide = fuzzerSide{coverage: fuzzerCoveredStruct, goStruct: st, pkg: specSide.pkg}
				break
			}
		}
	}

	plan := &FuzzerPlan{
		Kind:        in.kind,
		Message:     in.message,
		PackageName: in.direct.packageName,
		FuncName:    fuzzerFuncName(in.kind),
		Spec:        spec,
		Status:      status,
		Entries:     classifyFuzzerFields(in.message, specSide, statusSide),
		specSide:    specSide,
		statusSide:  statusSide,
	}
	if in.direct.handwrittenFunc(plan.FuncName+"FilterSpec") != "" {
		plan.FilterSpec = plan.FuncName + "FilterSpec"
	}
	if in.direct.handwrittenFunc(plan.FuncName+"FilterStatus") != "" {
		plan.FilterStatus = plan.FuncName + "FilterStatus"
	}
	return plan, nil
}

func (in fuzzerPlanInput) krmSide(pair *fuzzerMapperPair) (fuzzerSide, error) {
	if pair.KRMGoPackage == "" {
		return fuzzerSide{}, fmt.Errorf("%s returns *%s from the direct package itself, not a KRM package", pair.FromProto, pair.KRMType)
	}
	pkg, err := in.krm.load(pair.KRMGoPackage)
	if err != nil {
		return fuzzerSide{}, err
	}
	st := pkg.structs[pair.KRMType]
	if st == nil {
		return fuzzerSide{}, fmt.Errorf("%s returns %s.%s, but %s declares no such struct", pair.FromProto, pair.KRMGoPackage, pair.KRMType, pkg.importPath)
	}
	return fuzzerSide{coverage: fuzzerCoveredStruct, goStruct: st, pkg: pkg}, nil
}

// fuzzerFuncName returns <lowerCamelKind>Fuzzer. A leading acronym is lowercased as a whole,
// so ComputeTargetVPNGateway gives computeTargetVPNGatewayFuzzer and IAMDenyPolicy gives
// iamDenyPolicyFuzzer.
func fuzzerFuncName(kind string) string {
	runes := []rune(kind)
	upper := 0
	for upper < len(runes) && unicode.IsUpper(runes[upper]) {
		upper++
	}
	lower := upper
	if upper > 1 && upper < len(runes) && unicode.IsLower(runes[upper]) {
		// The last capital starts the next word, as in IAMDeny.
		lower = upper - 1
	}
	for i := 0; i < lower; i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes) + "Fuzzer"
}

// FuzzerGeneratorOptions configures a FuzzerGenerator.
type FuzzerGeneratorOptions struct {
	// Group and Version are the KRM API group and version, for example
	// compute.cnrm.cloud.google.com and v1beta1.
	Group   string
	Version string
	// ProtoService is the proto package of the resources. It is recorded in the file annotation.
	ProtoService string
	// APIGoPackagePath is the import path of the directory that holds the KRM packages,
	// github.com/GoogleCloudPlatform/k8s-config-connector/apis.
	APIGoPackagePath string
	// APIDirectory is the directory that holds the KRM packages, <repo>/apis.
	APIDirectory string
	// DirectDirectory is the direct controller package the fuzzers are written to,
	// <repo>/pkg/controller/direct/<service>.
	DirectDirectory string
	// Year is the copyright year of new files. Zero means the current year.
	Year int
}

// FuzzerGenerator writes <kind>_fuzzer.generated.go files into a direct controller package.
type FuzzerGenerator struct {
	opts FuzzerGeneratorOptions
	krm  *fuzzerKRMLoader
}

// FuzzerResult reports what the generator did for one kind.
type FuzzerResult struct {
	// Path is the generated file. It is written unless Skipped is set.
	Path string
	// Skipped is set when the package already has a fuzzer for the kind that isn't generated.
	// SkipReason says where it is.
	Skipped    bool
	SkipReason string
	// Plan is the generated fuzzer. It is nil when the kind was skipped.
	Plan *FuzzerPlan
}

func NewFuzzerGenerator(opts FuzzerGeneratorOptions) *FuzzerGenerator {
	return &FuzzerGenerator{
		opts: opts,
		krm:  newFuzzerKRMLoader(opts.APIGoPackagePath, opts.APIDirectory),
	}
}

// KRMGoPackage returns the import path of the KRM package for the configured group and
// version, derived the way generate-mapper derives it.
func (g *FuzzerGenerator) KRMGoPackage() string {
	service := strings.TrimSuffix(g.opts.Group, ".cnrm.cloud.google.com")
	return strings.TrimSuffix(g.opts.APIGoPackagePath, "/") + "/" + service + "/" + g.opts.Version
}

// FuzzerFileName returns the name of the generated fuzzer file for kind.
func FuzzerFileName(kind string) string {
	return strings.ToLower(kind) + "_fuzzer.generated.go"
}

// Generate writes the fuzzer for kind, whose resource message is msg. A fuzzer that isn't
// generated wins: when the package declares the fuzzer function in a file that isn't
// generated, or registers a fuzzer built from the kind's Spec mappers, nothing is written,
// and a stale generated file is removed.
func (g *FuzzerGenerator) Generate(kind string, msg protoreflect.MessageDescriptor) (*FuzzerResult, error) {
	direct, err := loadGoPackageIndex(g.opts.DirectDirectory)
	if err != nil {
		return nil, err
	}
	result := &FuzzerResult{Path: filepath.Join(g.opts.DirectDirectory, FuzzerFileName(kind))}

	funcName := fuzzerFuncName(kind)
	if p := direct.handwrittenFunc(funcName); p != "" {
		return g.skip(result, fmt.Sprintf("%s declares %s", p, funcName))
	}
	for _, base := range mapperPairBases(g.opts.Version, kind+"Spec") {
		fromProto := base + "_FromProto"
		if direct.funcs[fromProto] == nil {
			continue
		}
		p := direct.handwrittenFuzzerUsing(fromProto)
		if p == "" {
			if p, err = handwrittenFuzzerBelow(g.opts.DirectDirectory, fromProto); err != nil {
				return nil, err
			}
		}
		if p != "" {
			return g.skip(result, fmt.Sprintf("%s registers a fuzzer built from %s", p, fromProto))
		}
	}

	plan, err := planFuzzer(fuzzerPlanInput{
		kind:    kind,
		version: g.opts.Version,
		message: msg,
		direct:  direct,
		krm:     g.krm,
	})
	if err != nil {
		return nil, err
	}
	if want := g.KRMGoPackage(); plan.Spec.KRMGoPackage != want {
		return nil, fmt.Errorf("%s returns a type of %s, but --api-version %s/%s is %s", plan.Spec.FromProto, plan.Spec.KRMGoPackage, g.opts.Group, g.opts.Version, want)
	}
	result.Plan = plan

	year := g.opts.Year
	if year == 0 {
		year = time.Now().Year()
	}
	existing, err := os.ReadFile(result.Path)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("reading %q: %w", result.Path, err)
	}
	if existingYear, ok := extractCopyrightYear(existing); ok {
		year = existingYear
	}

	src, err := renderFuzzer(plan, g.fileAnnotation(msg), year)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(result.Path, src, 0644); err != nil {
		return nil, fmt.Errorf("writing %q: %w", result.Path, err)
	}
	return result, nil
}

func (g *FuzzerGenerator) skip(result *FuzzerResult, reason string) (*FuzzerResult, error) {
	result.Skipped = true
	result.SkipReason = reason
	klog.Infof("not generating %s: %s", result.Path, reason)
	if err := os.Remove(result.Path); err == nil {
		klog.Infof("removed stale %s", result.Path)
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("removing stale %q: %w", result.Path, err)
	}
	return result, nil
}

func (g *FuzzerGenerator) fileAnnotation(msg protoreflect.MessageDescriptor) *annotations.FileAnnotation {
	attributes := map[string][]string{
		"krm.group":     {g.opts.Group},
		"krm.version":   {g.opts.Version},
		"proto.message": {string(msg.FullName())},
	}
	if g.opts.ProtoService != "" {
		attributes["proto.service"] = []string{g.opts.ProtoService}
	}
	return &annotations.FileAnnotation{Key: FuzzerFileAnnotationKey, Attributes: attributes}
}

// renderFuzzer renders the generated fuzzer file, with the header conventions of mapper.generated.go.
func renderFuzzer(plan *FuzzerPlan, annotation *annotations.FileAnnotation, year int) ([]byte, error) {
	var b bytes.Buffer
	writeCopyright(&b, year)
	b.WriteString("// Code generated by dev/tasks/generate-all. DO NOT EDIT.\n")
	b.WriteString("//go:build !ignore_autogenerated\n")
	b.WriteString("// +build !ignore_autogenerated\n\n")
	fmt.Fprintf(&b, "%s\n", annotation.FormatGo())
	fmt.Fprintf(&b, "package %s\n\n", plan.PackageName)
	fmt.Fprintf(&b, "import (\n\tpb %q\n\t%q\n)\n\n", plan.Spec.ProtoGoPackage, fuzztestingGoPackage)

	register := "RegisterKRMFuzzer"
	if plan.Status == nil {
		register = "RegisterKRMSpecFuzzer"
	}
	fmt.Fprintf(&b, "func init() {\n\tfuzztesting.%s(%s())\n}\n\n", register, plan.FuncName)

	fmt.Fprintf(&b, "func %s() fuzztesting.KRMFuzzer {\n", plan.FuncName)
	if plan.Status != nil {
		fmt.Fprintf(&b, "\tf := fuzztesting.NewKRMTypedFuzzer(&pb.%s{},\n", plan.Spec.ProtoType)
		fmt.Fprintf(&b, "\t\t%s, %s,\n", plan.Spec.FromProto, plan.Spec.ToProto)
		fmt.Fprintf(&b, "\t\t%s, %s,\n", plan.Status.FromProto, plan.Status.ToProto)
	} else {
		fmt.Fprintf(&b, "\tf := fuzztesting.NewKRMTypedSpecFuzzer(&pb.%s{},\n", plan.Spec.ProtoType)
		fmt.Fprintf(&b, "\t\t%s, %s,\n", plan.Spec.FromProto, plan.Spec.ToProto)
	}
	b.WriteString("\t)\n")

	for _, group := range fuzzerEntryGroups {
		var lines []string
		for _, e := range plan.Entries {
			for _, m := range group.methods {
				if e.Method != m {
					continue
				}
				if m == FuzzerUnimplementedEtag {
					lines = append(lines, "\tf.Unimplemented_Etag()\n")
				} else {
					lines = append(lines, fmt.Sprintf("\tf.%s(%q)\n", e.Method, e.Path))
				}
			}
		}
		if len(lines) == 0 {
			continue
		}
		fmt.Fprintf(&b, "\n\t// %s\n", group.comment)
		for _, line := range lines {
			b.WriteString(line)
		}
	}

	if plan.FilterSpec != "" || plan.FilterStatus != "" {
		b.WriteString("\n")
	}
	if plan.FilterSpec != "" {
		fmt.Fprintf(&b, "\tf.FilterSpec = %s\n", plan.FilterSpec)
	}
	if plan.FilterStatus != "" {
		fmt.Fprintf(&b, "\tf.FilterStatus = %s\n", plan.FilterStatus)
	}
	b.WriteString("\n\treturn f\n}\n")

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		return nil, fmt.Errorf("formatting the fuzzer for %s: %w\n%s", plan.Kind, err, b.String())
	}
	return formatted, nil
}
