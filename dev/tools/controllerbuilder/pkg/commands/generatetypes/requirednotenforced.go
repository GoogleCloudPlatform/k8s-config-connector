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

package generatetypes

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// descriptorFinder looks up a proto descriptor by full name.
// *protoregistry.Files satisfies it.
type descriptorFinder interface {
	FindDescriptorByName(protoreflect.FullName) (protoreflect.Descriptor, error)
}

// goStruct is a struct type of the package, as the Go source declares it.
type goStruct struct {
	name string
	// file is the base name of the file that declares the struct.
	file string
	// generated is true for a struct in a file whose name ends in
	// generated.go, which generate-types writes.
	generated bool
	// proto is the message named by the struct's +kcc:proto annotation, or
	// one of its variants such as +kcc:spec:proto. It is "" if there is none.
	proto string
	// marked is true when the doc comment carries the
	// +kcc:required-from-proto marker.
	marked bool
	fields []goField
}

// goField is one field of a goStruct.
type goField struct {
	goName string
	// json is the name in the json tag, or the Go name when there is no tag.
	json string
	// inline is true for an embedded field with no json name, whose fields
	// sit at the parent's path.
	inline bool
	// typeName is the type the field holds, through pointers, slices and map
	// values. It is "" for a type from another package.
	typeName string
	isList   bool
	isMap    bool
	// protoField is the field named by +kcc:proto:field, or "".
	protoField string
	// required is true when the field carries +required or
	// +kubebuilder:validation:Required.
	required bool
}

// requiredNotEnforced returns a queue entry for each spec field, at any depth,
// that the proto marks REQUIRED and the CRD leaves optional, for each Kind in
// the run.
//
// It reads the Go source of the package, so it sees every field the CRD gets:
// generated, hand-written, and fields of Kinds scaffolded before the flag
// existed. Run it after WriteFiles and prune, so it reads what they wrote.
func requiredNotEnforced(files descriptorFinder, apisDir, goPackage, group string, kinds []string, protoFullNames map[string]string) ([]judgement.Entry, error) {
	pkgDir := filepath.Join(apisDir, goPackage)
	structs, err := loadStructs(pkgDir)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", pkgDir, err)
	}
	var out []judgement.Entry
	for _, kind := range kinds {
		spec := structs[kind+"Spec"]
		if spec == nil {
			continue
		}
		fqn := spec.proto
		if fqn == "" {
			fqn = protoFullNames[kind]
		}
		w := &requiredGapFinder{
			structs:  structs,
			files:    files,
			kind:     kind,
			group:    group,
			marked:   spec.marked,
			servedAs: servedBeyondAlpha(apisDir, goPackage, kind),
		}
		w.walk(spec, w.message(fqn), ".spec", true, nil, map[string]bool{})
		out = append(out, w.entries...)
	}
	return out, nil
}

// loadStructs reads the struct types declared in the Go files of dir. Test
// files and deepcopy are left out.
func loadStructs(dir string) (map[string]*goStruct, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	out := map[string]*goStruct{}
	fset := token.NewFileSet()
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") || strings.HasPrefix(name, "zz_generated") {
			continue
		}
		f, err := parser.ParseFile(fset, filepath.Join(dir, name), nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		for _, decl := range f.Decls {
			gd, ok := decl.(*ast.GenDecl)
			if !ok || gd.Tok != token.TYPE {
				continue
			}
			for _, spec := range gd.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				st, ok := ts.Type.(*ast.StructType)
				if !ok {
					continue
				}
				doc := ts.Doc
				if doc == nil && len(gd.Specs) == 1 {
					doc = gd.Doc
				}
				s := &goStruct{
					name:      ts.Name.Name,
					file:      name,
					generated: strings.HasSuffix(name, "generated.go"),
				}
				for _, line := range commentLines(doc) {
					if msg, _, ok := codegen.GetProtoMessageAndKindFromAnnotation(line); ok && s.proto == "" {
						s.proto = msg
					}
					if codegen.IsRequiredFromProtoMarker(line) {
						s.marked = true
					}
				}
				for _, field := range st.Fields.List {
					s.fields = append(s.fields, parseFields(field)...)
				}
				out[s.name] = s
			}
		}
	}
	return out, nil
}

func commentLines(cg *ast.CommentGroup) []string {
	if cg == nil {
		return nil
	}
	out := make([]string, 0, len(cg.List))
	for _, c := range cg.List {
		out = append(out, c.Text)
	}
	return out
}

// parseFields returns one goField per name the field declares.
func parseFields(field *ast.Field) []goField {
	var f goField
	t := field.Type
unwrap:
	for {
		switch tt := t.(type) {
		case *ast.StarExpr:
			t = tt.X
		case *ast.ArrayType:
			f.isList = true
			t = tt.Elt
		case *ast.MapType:
			f.isMap = true
			t = tt.Value
		default:
			break unwrap
		}
	}
	if id, ok := t.(*ast.Ident); ok {
		f.typeName = id.Name
	}
	for _, line := range commentLines(field.Doc) {
		marker := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "//"))
		if v, ok := strings.CutPrefix(marker, codegen.KCCProtoFieldAnnotation+"="); ok {
			f.protoField = strings.TrimSpace(v)
		}
		if marker == "+required" || marker == "+kubebuilder:validation:Required" {
			f.required = true
		}
	}
	if field.Tag != nil {
		if tag, err := strconv.Unquote(field.Tag.Value); err == nil {
			f.json, _, _ = strings.Cut(reflect.StructTag(tag).Get("json"), ",")
		}
	}
	if len(field.Names) == 0 {
		f.goName = f.typeName
		f.inline = f.json == ""
		return []goField{f}
	}
	out := make([]goField, 0, len(field.Names))
	for _, n := range field.Names {
		g := f
		g.goName = n.Name
		if g.json == "" {
			g.json = n.Name
		}
		out = append(out, g)
	}
	return out
}

// servedBeyondAlpha returns the version that makes kind beta or GA: the
// package's own version when it is not alpha, or another version of the
// service that also declares the Kind. It returns "" for an alpha-only Kind.
func servedBeyondAlpha(apisDir, goPackage, kind string) string {
	version := path.Base(goPackage)
	if !strings.Contains(version, "alpha") {
		return version
	}
	serviceDir := filepath.Dir(filepath.Join(apisDir, goPackage))
	entries, err := os.ReadDir(serviceDir)
	if err != nil {
		return ""
	}
	for _, e := range entries {
		if !e.IsDir() || e.Name() == version || strings.Contains(e.Name(), "alpha") {
			continue
		}
		if declaresStruct(filepath.Join(serviceDir, e.Name()), kind) {
			return e.Name()
		}
	}
	return ""
}

// declaresStruct reports whether a Go file in dir declares the struct name.
func declaresStruct(dir, name string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	decl := []byte("\ntype " + name + " struct")
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err == nil && bytes.Contains(b, decl) {
			return true
		}
	}
	return false
}

// requiredGapFinder walks the Spec of one Kind and records the fields the
// proto marks REQUIRED that have no +required marker.
type requiredGapFinder struct {
	structs map[string]*goStruct
	files   descriptorFinder
	kind    string
	group   string
	// marked is true when the Kind's Spec carries the marker.
	marked bool
	// servedAs is the beta or GA version the Kind is served at, or "" for an
	// alpha Kind.
	servedAs string
	entries  []judgement.Entry
}

// enteredStruct is where a walk went from a hand-written struct into a
// generated one. For a Kind with the marker, the hand-written field is the
// one to switch to the generated struct's Required copy.
type enteredStruct struct {
	// parent is the hand-written struct, and field the Go name of its field.
	parent *goStruct
	field  string
	// child is the generated struct the field holds.
	child *goStruct
}

func (w *requiredGapFinder) walk(s *goStruct, msg protoreflect.MessageDescriptor, prefix string, top bool, entered *enteredStruct, onPath map[string]bool) {
	if onPath[s.name] {
		return
	}
	onPath[s.name] = true
	defer delete(onPath, s.name)

	for _, f := range s.fields {
		if f.json == "-" {
			continue
		}
		child := w.structs[f.typeName]
		if f.inline {
			if child != nil {
				w.walk(child, msg, prefix, top, w.enter(s, f, child, entered), onPath)
			}
			continue
		}

		fieldPath := prefix + "." + f.json
		fd := w.protoField(f, msg)
		if fd != nil && w.isGap(f, fd, top) {
			w.entries = append(w.entries, judgement.Entry{
				Kind:   w.kind,
				Group:  w.group,
				Field:  fieldPath,
				Reason: judgement.ReasonRequiredNotEnforced,
				Detail: w.detail(s, f, entered),
				Status: judgement.StatusOpen,
			})
		}

		if child == nil {
			continue
		}
		childMsg := w.message(child.proto)
		if childMsg == nil && fd != nil {
			childMsg = fieldMessage(fd)
		}
		// The same path shapes as the reference hints use.
		childPath := fieldPath
		switch {
		case f.isMap:
			childPath += ".KEY"
		case f.isList:
			childPath += "[]"
		}
		w.walk(child, childMsg, childPath, false, w.enter(s, f, child, entered), onPath)
	}
}

// enter returns what the walk below child records as entered. It is reset at
// each hand-written struct, and set where a hand-written struct holds a
// generated one.
func (w *requiredGapFinder) enter(parent *goStruct, f goField, child *goStruct, entered *enteredStruct) *enteredStruct {
	switch {
	case !child.generated:
		return nil
	case !parent.generated:
		return &enteredStruct{parent: parent, field: f.goName, child: child}
	default:
		return entered
	}
}

// isGap reports whether the proto marks fd REQUIRED and f does not enforce it.
func (w *requiredGapFinder) isGap(f goField, fd protoreflect.FieldDescriptor, top bool) bool {
	if f.required {
		return false
	}
	if !codegen.IsFieldBehavior(fd, annotations.FieldBehavior_REQUIRED) || codegen.IsFieldBehavior(fd, annotations.FieldBehavior_OUTPUT_ONLY) {
		return false
	}
	// The resource's own name is metadata.name or spec.resourceID in KRM.
	return !(top && fd.Name() == "name")
}

// protoField finds the proto field a Go field maps to: the one its
// +kcc:proto:field annotation names, or else the field of msg whose KRM name
// matches, allowing for the Ref and Refs names of references.
func (w *requiredGapFinder) protoField(f goField, msg protoreflect.MessageDescriptor) protoreflect.FieldDescriptor {
	if f.protoField != "" {
		if d, err := w.files.FindDescriptorByName(protoreflect.FullName(f.protoField)); err == nil {
			if fd, ok := d.(protoreflect.FieldDescriptor); ok {
				return fd
			}
		}
	}
	if msg == nil {
		return nil
	}
	for i := 0; i < msg.Fields().Len(); i++ {
		fd := msg.Fields().Get(i)
		for _, plural := range []bool{false, true} {
			for _, name := range krmNamesFor(codegen.GetJSONForKRM(fd, codegen.WriteOptions{EmitPluralAcronyms: plural})) {
				if name == f.json {
					return fd
				}
			}
		}
	}
	return nil
}

// krmNamesFor returns the KRM names a proto field with JSON name j can have:
// j itself, or a reference to what it names.
func krmNamesFor(j string) []string {
	names := []string{j, j + "Ref", j + "Refs", strings.TrimSuffix(j, "s") + "Refs"}
	for _, suffix := range []string{"Name", "Id"} {
		if base, ok := strings.CutSuffix(j, suffix); ok && base != "" {
			names = append(names, base+"Ref")
		}
		if base, ok := strings.CutSuffix(j, suffix+"s"); ok && base != "" {
			names = append(names, base+"Refs")
		}
	}
	return names
}

func (w *requiredGapFinder) message(fqn string) protoreflect.MessageDescriptor {
	if fqn == "" {
		return nil
	}
	d, err := w.files.FindDescriptorByName(protoreflect.FullName(fqn))
	if err != nil {
		return nil
	}
	msg, _ := d.(protoreflect.MessageDescriptor)
	return msg
}

// fieldMessage returns the message a field holds: its own, or its map's
// value message.
func fieldMessage(fd protoreflect.FieldDescriptor) protoreflect.MessageDescriptor {
	if fd.IsMap() {
		if v := fd.MapValue(); v.Kind() == protoreflect.MessageKind {
			return v.Message()
		}
		return nil
	}
	if fd.Kind() == protoreflect.MessageKind {
		return fd.Message()
	}
	return nil
}

// isRequiredCopy reports whether s is the Required copy generate-types writes
// of a message.
func (w *requiredGapFinder) isRequiredCopy(s *goStruct) bool {
	msg := w.message(s.proto)
	return s.generated && msg != nil && s.name == codegen.RequiredStructName(msg)
}

// requiredNameTaken reports whether another type has the name of s's
// Required copy. generate-types then writes no copy, and s stays optional for
// every Kind.
func (w *requiredGapFinder) requiredNameTaken(s *goStruct) bool {
	other := w.structs[s.name+"Required"]
	return other != nil && !w.isRequiredCopy(other)
}

// detail says why the field is optional, how to enforce it, and whether that
// is allowed for this Kind.
func (w *requiredGapFinder) detail(s *goStruct, f goField, entered *enteredStruct) string {
	const what = "The proto marks this field REQUIRED, but the CRD leaves it optional."
	if w.servedAs != "" {
		return fmt.Sprintf("%s Keep it optional: %s is %s, and making a field required breaks objects that leave it out.", what, w.kind, w.servedAs)
	}
	var how string
	switch {
	case !s.generated:
		how = fmt.Sprintf("To enforce it, add +required to %s.%s in %s.", s.name, f.goName, s.file)
	case !w.marked:
		how = fmt.Sprintf("To enforce it, add %s to %sSpec and run generate.sh again.", codegen.RequiredFromProtoMarker, w.kind)
	case w.requiredNameTaken(s):
		how = fmt.Sprintf("%s has no Required copy because another type is called %sRequired. To enforce it, rename that type and run generate.sh again.", s.name, s.name)
	case entered != nil && !w.isRequiredCopy(entered.child):
		// generate-types writes the copy once code names it.
		how = fmt.Sprintf("To enforce it, change %s.%s in %s from %s to %sRequired and run generate.sh again.",
			entered.parent.name, entered.field, entered.parent.file, entered.child.name, entered.child.name)
	default:
		how = fmt.Sprintf("%s has no +required markers because something else also uses it: a Kind without the marker, a status struct, a hand-written type or another package.", s.name)
	}
	return fmt.Sprintf("%s %s %s is alpha, so it may be enforced, but that breaks objects that leave the field out if the Kind has shipped.", what, how, w.kind)
}
