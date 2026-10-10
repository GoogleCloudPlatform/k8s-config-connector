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
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
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

// goStruct is a struct declared in the package's Go source.
type goStruct struct {
	name string
	// file is the base name of the file that declares the struct.
	file string
	// generated is true for a struct in a file ending in generated.go,
	// which generate-types writes.
	generated bool
	// proto is the message in the struct's +kcc:proto annotation, or in a
	// variant such as +kcc:spec:proto. It is "" if there is none.
	proto string
	// marked is true if the doc comment has the +kcc:required-from-proto
	// marker.
	marked bool
	// stability is the level in the struct's stability-level label, or "" if
	// it has none. Only a Kind has one.
	stability stabilityLevel
	fields    []goField
}

// goField is one field of a goStruct.
type goField struct {
	goName string
	// json is the name in the json tag, or the Go name when there is no tag.
	json string
	// inline is true for an embedded field without a json name. Its fields
	// sit at the parent's path.
	inline bool
	// typeName is the type the field holds, looking through pointers, slices
	// and map values. It is "" for a type from another package.
	typeName string
	// qualified is a type from another package, such as "metav1.TypeMeta",
	// or "".
	qualified string
	isList    bool
	isMap     bool
	// protoField is the field named by +kcc:proto:field, or "".
	protoField string
	// required is true if the field has +required or
	// +kubebuilder:validation:Required.
	required bool
}

// requiredNotEnforced returns a queue entry for each spec field, at any
// depth, that the proto marks REQUIRED but the CRD leaves optional. It checks
// each Kind in the run.
//
// It reads the package's Go source, so it sees every field the CRD gets:
// generated ones, hand-written ones, and those of Kinds scaffolded before the
// marker existed. Run it after WriteFiles and prune. It also reads the other
// versions of the service, to find how stable each Kind is.
func requiredNotEnforced(files descriptorFinder, apisDir, goPackage, group string, kinds []string, protoFullNames map[string]string) ([]judgement.Entry, error) {
	pkgDir := filepath.Join(apisDir, goPackage)
	versions, err := loadServiceVersions(filepath.Dir(pkgDir))
	if err != nil {
		return nil, err
	}
	version := path.Base(goPackage)
	structs, ok := versions[version]
	if !ok {
		return nil, fmt.Errorf("reading %s: no such directory", pkgDir)
	}
	users := structUsers(structs)
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
		stability, stableVersion := resourceStabilityLevel(versions, version, kind)
		w := &requiredGapFinder{
			structs:   structs,
			users:     users,
			files:     files,
			kind:      kind,
			group:     group,
			marked:    spec.marked,
			stability: stability,
			version:   stableVersion,
		}
		w.walk(spec, w.message(fqn), ".spec", true, nil, map[string]bool{})
		out = append(out, w.entries...)
	}
	return out, nil
}

// loadServiceVersions loads the structs of each version of a service, such as
// v1alpha1 and v1beta1, keyed by version.
func loadServiceVersions(serviceDir string) (map[string]map[string]*goStruct, error) {
	entries, err := os.ReadDir(serviceDir)
	if err != nil {
		return nil, err
	}
	out := map[string]map[string]*goStruct{}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		dir := filepath.Join(serviceDir, e.Name())
		structs, err := loadStructs(dir)
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", dir, err)
		}
		out[e.Name()] = structs
	}
	return out, nil
}

// loadStructs reads the structs declared in the Go files of dir, skipping
// tests and deepcopy.
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
					if level, ok := stabilityLabel(line); ok {
						s.stability = level
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
	switch tt := t.(type) {
	case *ast.Ident:
		f.typeName = tt.Name
	case *ast.SelectorExpr:
		if pkg, ok := tt.X.(*ast.Ident); ok {
			f.qualified = pkg.Name + "." + tt.Sel.Name
		}
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

// stabilityLevel is the value of a Kind's cnrm.cloud.google.com/stability-level
// label.
type stabilityLevel string

const (
	stabilityAlpha  stabilityLevel = "alpha"
	stabilityBeta   stabilityLevel = "beta"
	stabilityStable stabilityLevel = "stable"
)

// moreStableThan reports whether l is more stable than other.
func (l stabilityLevel) moreStableThan(other stabilityLevel) bool {
	return l.rank() > other.rank()
}

func (l stabilityLevel) rank() int {
	switch l {
	case stabilityAlpha:
		return 1
	case stabilityBeta:
		return 2
	case stabilityStable:
		return 3
	default:
		return 0
	}
}

// stabilityLabelPattern matches the stability-level label in a
// +kubebuilder:metadata:labels marker, which can list several labels.
var stabilityLabelPattern = regexp.MustCompile(`^\s*//\s*\+kubebuilder:metadata:labels=.*"cnrm\.cloud\.google\.com/stability-level=(alpha|beta|stable)"`)

// stabilityLabel returns the level that a doc comment line's
// +kubebuilder:metadata:labels marker sets.
func stabilityLabel(line string) (stabilityLevel, bool) {
	m := stabilityLabelPattern.FindStringSubmatch(line)
	if m == nil {
		return "", false
	}
	return stabilityLevel(m[1]), true
}

// resourceStabilityLevel returns the Kind's most stable level across the
// versions of its service, and the version that has it. version is the
// package of the Kind's Spec. It counts even when it doesn't declare the Kind
// itself.
//
// A Kind's level in a version is its stability-level label. Kinds that
// generate-types scaffolds have no label, so then the version's name decides.
func resourceStabilityLevel(versions map[string]map[string]*goStruct, version, kind string) (stabilityLevel, string) {
	best, bestVersion := levelIn(versions[version][kind], version), version
	for _, v := range sortedVersions(versions) {
		s := versions[v][kind]
		if v == version || s == nil {
			continue
		}
		if level := levelIn(s, v); level.moreStableThan(best) {
			best, bestVersion = level, v
		}
	}
	return best, bestVersion
}

// levelIn returns the level of a Kind in version. s is the Kind's struct in
// that version, or nil if the version doesn't declare it.
func levelIn(s *goStruct, version string) stabilityLevel {
	if s != nil && s.stability != "" {
		return s.stability
	}
	return versionStability(version)
}

// versionStability returns the level a version's name implies: alpha for
// v1alpha1, beta for v1beta1 and stable for v1.
func versionStability(version string) stabilityLevel {
	switch {
	case strings.Contains(version, "alpha"):
		return stabilityAlpha
	case strings.Contains(version, "beta"):
		return stabilityBeta
	default:
		return stabilityStable
	}
}

func sortedVersions(versions map[string]map[string]*goStruct) []string {
	out := make([]string, 0, len(versions))
	for v := range versions {
		out = append(out, v)
	}
	sort.Strings(out)
	return out
}

// structUser is a Kind whose spec or status holds a struct, directly or
// through other structs.
type structUser struct {
	kind string
	// status is true when the Kind's status holds the struct, and false when
	// its spec does.
	status bool
}

// String names the user in a queue detail.
func (u structUser) String() string {
	if u.status {
		return "the status of " + u.kind
	}
	return u.kind
}

// structUsers maps each struct in the package to its users: the Kinds whose
// spec or status holds it, directly or through other structs. A Kind is a
// struct that embeds metav1.TypeMeta.
func structUsers(structs map[string]*goStruct) map[string][]structUser {
	seen := map[string]map[structUser]bool{}
	var visit func(name string, user structUser)
	visit = func(name string, user structUser) {
		s := structs[name]
		if s == nil || seen[name][user] {
			return
		}
		if seen[name] == nil {
			seen[name] = map[structUser]bool{}
		}
		seen[name][user] = true
		for _, f := range s.fields {
			visit(f.typeName, user)
		}
	}
	for _, s := range structs {
		if !isKind(s) {
			continue
		}
		for _, f := range s.fields {
			switch f.json {
			case "spec":
				visit(f.typeName, structUser{kind: s.name})
			case "status":
				visit(f.typeName, structUser{kind: s.name, status: true})
			}
		}
	}
	out := map[string][]structUser{}
	for name, users := range seen {
		for u := range users {
			out[name] = append(out[name], u)
		}
		sortUsers(out[name])
	}
	return out
}

// sortUsers sorts users by how a queue detail names them.
func sortUsers(users []structUser) {
	sort.Slice(users, func(i, j int) bool { return users[i].String() < users[j].String() })
}

func isKind(s *goStruct) bool {
	for _, f := range s.fields {
		if f.inline && strings.HasSuffix(f.qualified, ".TypeMeta") {
			return true
		}
	}
	return false
}

// joinUsers lists users, separated by commas.
func joinUsers(users []structUser) string {
	names := make([]string, 0, len(users))
	for _, u := range users {
		names = append(names, u.String())
	}
	return strings.Join(names, ", ")
}

// requiredGapFinder walks one Kind's Spec and records the fields the proto
// marks REQUIRED that have no +required marker.
type requiredGapFinder struct {
	structs map[string]*goStruct
	// users maps each struct to its users; see structUsers.
	users map[string][]structUser
	files descriptorFinder
	kind  string
	group string
	// marked is true if the Kind's Spec has the marker.
	marked bool
	// stability is the Kind's most stable level, and version the version
	// that has it.
	stability stabilityLevel
	version   string
	entries   []judgement.Entry
}

// sharedWith returns the other users of the named struct: other Kinds, and
// the status of any Kind, including this one. A +required marker in that
// struct would apply to all of them.
func (w *requiredGapFinder) sharedWith(name string) []structUser {
	var out []structUser
	for _, u := range w.users[name] {
		if u.status || u.kind != w.kind {
			out = append(out, u)
		}
	}
	return out
}

// enteredStruct is where the walk crossed from a hand-written struct into a
// generated one. For an opted-in Kind, that hand-written field should switch
// to the generated struct's Required copy.
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
		// Same path format as the reference hints.
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

// enter returns the enteredStruct for the walk below child: nil when child
// is hand-written, a new one when a hand-written parent holds a generated
// child, and entered unchanged otherwise.
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

// protoField finds the proto field a Go field maps to: the one named in its
// +kcc:proto:field annotation, or else the field of msg with a matching KRM
// name, including the Ref and Refs names of references.
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
// j itself, or a Ref or Refs name if the field is a reference.
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

// fieldMessage returns the message a field holds: its own message, or the
// value message of a map.
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

// isRequiredCopy reports whether s is a Required copy that generate-types
// wrote.
func (w *requiredGapFinder) isRequiredCopy(s *goStruct) bool {
	msg := w.message(s.proto)
	return s.generated && msg != nil && s.name == codegen.RequiredStructName(msg)
}

// requiredNameTaken reports whether another type already has the name of s's
// Required copy. If so, generate-types writes no copy, and s stays optional
// for every Kind.
func (w *requiredGapFinder) requiredNameTaken(s *goStruct) bool {
	other := w.structs[s.name+"Required"]
	return other != nil && !w.isRequiredCopy(other)
}

// detail says why the field is optional, how to enforce it, and whether this
// Kind may.
//
// The fix only ever changes this Kind. If another Kind or a status also uses
// the hand-written struct to edit, the detail asks for a copy instead, since
// an edit in place would change them too.
func (w *requiredGapFinder) detail(s *goStruct, f goField, entered *enteredStruct) string {
	const what = "The proto marks this field REQUIRED, but the CRD leaves it optional."
	if w.stability != stabilityAlpha {
		return fmt.Sprintf("%s Keep it optional: %s is %s (%s), and making a field required breaks objects that leave it out.", what, w.kind, w.stability, w.version)
	}
	var how string
	switch {
	case !s.generated:
		if others := w.sharedWith(s.name); len(others) > 0 {
			how = fmt.Sprintf("%s in %s is also used by %s. +required on %s.%s would apply there too. To enforce it for %s only, give %s its own copy of %s.",
				s.name, s.file, joinUsers(others), s.name, f.goName, w.kind, w.kind, s.name)
		} else {
			how = fmt.Sprintf("To enforce it, add +required to %s.%s in %s.", s.name, f.goName, s.file)
		}
	case !w.marked:
		how = fmt.Sprintf("To enforce it, add %s to %sSpec and run generate.sh again.", codegen.RequiredFromProtoMarker, w.kind)
	case w.requiredNameTaken(s):
		how = fmt.Sprintf("%s has no Required copy because another type is called %sRequired. To enforce it, rename that type and run generate.sh again.", s.name, s.name)
	case entered != nil && !w.isRequiredCopy(entered.child):
		// generate-types writes the copy once code uses it.
		p := entered.parent
		if others := w.sharedWith(p.name); len(others) > 0 {
			how = fmt.Sprintf("%s in %s is also used by %s. To enforce it for %s only, give %s its own copy of %s whose %s field holds %sRequired, and run generate.sh again.",
				p.name, p.file, joinUsers(others), w.kind, w.kind, p.name, entered.field, entered.child.name)
		} else {
			how = fmt.Sprintf("To enforce it, change %s.%s in %s from %s to %sRequired and run generate.sh again.",
				p.name, entered.field, p.file, entered.child.name, entered.child.name)
		}
	default:
		how = fmt.Sprintf("%s has no +required markers because something else also uses it: a Kind without the marker, a status struct, a hand-written type or another package.", s.name)
	}
	return fmt.Sprintf("%s %s %s is alpha, so it may be enforced, but that breaks objects that leave the field out if the Kind has shipped.", what, how, w.kind)
}
