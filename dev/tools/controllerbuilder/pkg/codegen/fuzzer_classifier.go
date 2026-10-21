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

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	// fuzzerMaxDepth is the deepest field path, in path segments, that the generator splits.
	// fuzz.FillWithRandom doesn't fill messages nested more than 10 levels below the resource
	// message, so their fields are never set.
	fuzzerMaxDepth = 11

	fuzzerIdentityPath = ".name"
	fuzzerEtagPath     = ".etag"
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

// fuzzerPathKey identifies a message and the structs it is paired with, to stop recursing
// through recursive types.
type fuzzerPathKey struct {
	message protoreflect.FullName
	spec    *gocode.GoStruct
	status  *gocode.GoStruct
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
		expandFuzzerNode(n, depth, onPath)
		n.classes = computeNodeClasses(n)
		nodes = append(nodes, n)
	}
	return nodes
}

func expandFuzzerNode(n *fuzzerNode, depth int, onPath map[fuzzerPathKey]bool) {
	expand := depth < fuzzerMaxDepth &&
		(n.spec.coverage == fuzzerCoveredStruct || n.status.coverage == fuzzerCoveredStruct)
	if !expand {
		return
	}
	key := fuzzerPathKey{message: n.field.Message().FullName(), spec: n.spec.goStruct, status: n.status.goStruct}
	if onPath[key] {
		// A KRM type that contains itself. Its fields repeat below this field, down to
		// the depth that fuzz.FillWithRandom fills, and some of them may be unmapped.
		// Rather than list them again at every level, neither side keeps the field.
		n.spec, n.status = fuzzerSide{}, fuzzerSide{}
		return
	}
	onPath[key] = true
	childPrefix := n.path
	if n.field.IsList() {
		childPrefix += "[]"
	}
	n.children = buildFuzzerNodes(childPrefix, n.field.Message(), n.spec, n.status, depth+1, onPath)
	delete(onPath, key)
}

func computeNodeClasses(n *fuzzerNode) fuzzerClass {
	if len(n.children) == 0 {
		return leafFuzzerClass(n.spec, n.status)
	}
	var classes fuzzerClass
	for _, child := range n.children {
		classes |= child.classes
	}
	return classes
}

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

// pairFuzzerField returns how one side maps proto field fd, given how it maps the message
// that contains fd. It follows MapperGenerator.writeMapFunctionsForPair: the KRM field is
// named goFieldName(fd), or <name>Ref for a reference, or <name without a trailing s>Refs for
// a list of references. Failing that, the KRM field is the one whose +kcc:proto:field
// annotation names fd: generate-types writes the annotation, and it stays when a field is
// renamed, for example into a reference or to spell an acronym differently.
func pairFuzzerField(fd protoreflect.FieldDescriptor, parent fuzzerSide) fuzzerSide {
	switch parent.coverage {
	case fuzzerNotCovered:
		return fuzzerSide{}
	case fuzzerCoveredOpaque:
		return fuzzerSide{coverage: fuzzerCoveredOpaque}
	}

	krmField, isRef := lookupKRMFieldForProto(fd, parent)
	if isRef {
		return fuzzerSide{coverage: fuzzerCoveredOpaque}
	}
	if krmField == nil {
		return fuzzerSide{}
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

func lookupKRMFieldForProto(fd protoreflect.FieldDescriptor, parent fuzzerSide) (field *gocode.StructField, isRef bool) {
	krmFieldName := goFieldName(fd)
	if f := findStructField(parent.goStruct, krmFieldName); f != nil {
		return f, false
	}
	if findStructField(parent.goStruct, krmFieldName+"Ref") != nil ||
		findStructField(parent.goStruct, strings.TrimSuffix(krmFieldName, "s")+"Refs") != nil {
		return nil, true
	}
	annotated := parent.pkg.protoFields[parent.goStruct.Name][string(fd.FullName())]
	if annotated == "" {
		return nil, false
	}
	f := findStructField(parent.goStruct, annotated)
	if f == nil {
		return nil, false
	}
	if strings.HasSuffix(f.Name, "Ref") || strings.HasSuffix(f.Name, "Refs") {
		return f, true
	}
	return f, false
}

func findStructField(s *gocode.GoStruct, name string) *gocode.StructField {
	for _, f := range s.Fields {
		if f.Name == name {
			return f
		}
	}
	return nil
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

func sortFuzzerEntries(entries []FuzzerEntry) {
	sort.Slice(entries, func(i, j int) bool {
		oi, oj := fuzzerMethodOrder(entries[i].Method), fuzzerMethodOrder(entries[j].Method)
		if oi != oj {
			return oi < oj
		}
		return entries[i].Path < entries[j].Path
	})
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
		if err := collectFileProtoFieldAnnotations(fset, p, out); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func collectFileProtoFieldAnnotations(fset *token.FileSet, filePath string, out map[string]map[string]string) error {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("reading %q: %w", filePath, err)
	}
	if bytes.Contains(src, []byte("//go:build !ignore_autogenerated")) {
		return nil
	}
	file, err := parser.ParseFile(fset, filePath, src, parser.ParseComments|parser.SkipObjectResolution)
	if err != nil {
		return fmt.Errorf("parsing %q: %w", filePath, err)
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
			collectStructProtoFieldAnnotations(typeSpec.Name.Name, structType, out)
		}
	}
	return nil
}

func collectStructProtoFieldAnnotations(structName string, structType *ast.StructType, out map[string]map[string]string) {
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
			fields := out[structName]
			if fields == nil {
				fields = make(map[string]string)
				out[structName] = fields
			}
			fields[strings.TrimSpace(protoField)] = field.Names[0].Name
		}
	}
}
