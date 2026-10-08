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
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// requiredSplitFixture builds two resources that share nested messages.
//
//   - Shared has a REQUIRED field. Holder holds a Shared.
//   - OnlyNew has a REQUIRED field, and only NewKind's Spec holds it.
//   - Zone has a REQUIRED field. NewKind's Spec holds one, and its status holds
//     a map of them, which ObservedState writes with the plain struct.
//   - OldKind and NewKind both hold a Holder.
func requiredSplitFixture(t *testing.T) protoreflect.FileDescriptor {
	t.Helper()
	behavior := func(b annotations.FieldBehavior) *descriptorpb.FieldOptions {
		o := &descriptorpb.FieldOptions{}
		proto.SetExtension(o, annotations.E_FieldBehavior, []annotations.FieldBehavior{b})
		return o
	}
	required := behavior(annotations.FieldBehavior_REQUIRED)
	outputOnly := behavior(annotations.FieldBehavior_OUTPUT_ONLY)
	str := func(name string, num int32, opts *descriptorpb.FieldOptions) *descriptorpb.FieldDescriptorProto {
		return &descriptorpb.FieldDescriptorProto{
			Name: protoPtr(name), Number: protoPtr(num),
			Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING), Options: opts,
		}
	}
	msg := func(name string, num int32, typeName string, repeated bool) *descriptorpb.FieldDescriptorProto {
		f := &descriptorpb.FieldDescriptorProto{
			Name: protoPtr(name), Number: protoPtr(num),
			Type:     typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE),
			TypeName: protoPtr(".google.cloud.test.v1." + typeName),
		}
		if repeated {
			f.Label = labelDescriptor(descriptorpb.FieldDescriptorProto_LABEL_REPEATED)
		}
		return f
	}
	zonesByName := msg("zones_by_name", 7, "NewKind.ZonesByNameEntry", true)
	zonesByName.Options = outputOnly

	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("split.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: protoPtr("Shared"), Field: []*descriptorpb.FieldDescriptorProto{str("key", 1, required), str("value", 2, nil)}},
			{Name: protoPtr("Holder"), Field: []*descriptorpb.FieldDescriptorProto{msg("shared", 1, "Shared", false)}},
			{Name: protoPtr("OnlyNew"), Field: []*descriptorpb.FieldDescriptorProto{str("token", 1, required)}},
			{Name: protoPtr("Zone"), Field: []*descriptorpb.FieldDescriptorProto{str("zone_name", 1, required)}},
			{
				Name: protoPtr("OldKind"),
				Field: []*descriptorpb.FieldDescriptorProto{
					str("name", 1, nil),
					msg("holder", 2, "Holder", false),
					str("state", 3, outputOnly),
				},
			},
			{
				Name: protoPtr("NewKind"),
				NestedType: []*descriptorpb.DescriptorProto{{
					Name:    protoPtr("ZonesByNameEntry"),
					Field:   []*descriptorpb.FieldDescriptorProto{str("key", 1, nil), msg("value", 2, "Zone", false)},
					Options: &descriptorpb.MessageOptions{MapEntry: protoPtr(true)},
				}},
				Field: []*descriptorpb.FieldDescriptorProto{
					str("name", 1, nil),
					msg("holder", 2, "Holder", false),
					msg("only_new", 3, "OnlyNew", false),
					msg("shared_list", 4, "Shared", true),
					msg("zone", 5, "Zone", false),
					str("state", 6, outputOnly),
					zonesByName,
				},
			},
		},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	return fd
}

// newKindSpecFields are the fields PrepopulateSpec writes into NewKindSpec.
func newKindSpecFields(fd protoreflect.FileDescriptor) []protoreflect.FieldDescriptor {
	fields := fd.Messages().ByName("NewKind").Fields()
	return []protoreflect.FieldDescriptor{
		fields.ByName("holder"),
		fields.ByName("only_new"),
		fields.ByName("shared_list"),
		fields.ByName("zone"),
	}
}

const testImportPath = "example.com/apis/test"

// planAndWrite runs generate-types the way RunGenerateCRD does: visit both
// resources, plan, then write. files are written under the apis directory
// first, keyed by their path relative to it.
func planAndWrite(t *testing.T, files map[string]string, newSpecFields []protoreflect.FieldDescriptor) (*TypeGenerator, string) {
	t.Helper()
	fd := requiredSplitFixture(t)
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "test"), 0o755); err != nil {
		t.Fatal(err)
	}
	for name, src := range files {
		p := filepath.Join(dir, name)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(src), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	g := NewTypeGenerator("test", dir, nil)
	g.WithWriteOptions(WriteOptions{EmitRequired: true, EmitMessageMaps: true})
	g.WithReservedTypeNames("OldKind", "NewKind", "NewKindObservedState")
	for _, name := range []protoreflect.Name{"OldKind", "NewKind"} {
		if err := g.visitMessage(fd.Messages().ByName(name)); err != nil {
			t.Fatalf("visitMessage(%s): %v", name, err)
		}
	}
	uses, err := ScanRequiredUses(filepath.Join(dir, "test"), dir, testImportPath)
	if err != nil {
		t.Fatalf("ScanRequiredUses: %v", err)
	}
	g.PlanRequiredStructs(uses, newSpecFields)
	if err := g.WriteVisitedMessages(); err != nil {
		t.Fatalf("WriteVisitedMessages: %v", err)
	}
	if err := g.WriteOutputMessages(); err != nil {
		t.Fatalf("WriteOutputMessages: %v", err)
	}
	f := g.generatedFiles[generatedFileKey{GoPackage: "test", FileName: "types.generated.go"}]
	if f == nil {
		t.Fatal("types.generated.go was not generated")
	}
	return g, f.body.String()
}

// structBody returns the declaration of the named struct in body, from its
// proto annotation to the closing brace, or "" if body does not declare it.
func structBody(body, name string) string {
	start := strings.Index(body, "type "+name+" struct {")
	if start < 0 {
		return ""
	}
	if annotation := strings.LastIndex(body[:start], "// +kcc:"); annotation >= 0 && !strings.Contains(body[annotation:start], "}") {
		start = annotation
	}
	end := strings.Index(body[start:], "\n}\n")
	if end < 0 {
		return body[start:]
	}
	return body[start : start+end+2]
}

type structWant struct {
	name    string
	want    []string
	notWant []string
}

func checkStructs(t *testing.T, body string, wants []structWant) {
	t.Helper()
	for _, w := range wants {
		got := structBody(body, w.name)
		if got == "" {
			t.Errorf("%s is not declared; body:\n%s", w.name, body)
			continue
		}
		for _, s := range w.want {
			if !strings.Contains(got, s) {
				t.Errorf("%s is missing %q:\n%s", w.name, s, got)
			}
		}
		for _, s := range w.notWant {
			if strings.Contains(got, s) {
				t.Errorf("%s has %q:\n%s", w.name, s, got)
			}
		}
	}
}

// splitWants is what every case below expects: structs that something other
// than the opted-in Kind uses keep their name and stay optional, and the
// opted-in Kind gets Required copies.
var splitWants = []structWant{
	{name: "Shared", want: []string{"Key *string"}, notWant: []string{"+required"}},
	{name: "SharedRequired", want: []string{"// +kcc:proto=google.cloud.test.v1.Shared\n", "// +required\n\tKey *string"}},
	{name: "Holder", want: []string{"Shared *Shared `"}, notWant: []string{"+required", "SharedRequired"}},
	{name: "HolderRequired", want: []string{"// +kcc:proto=google.cloud.test.v1.Holder\n", "Shared *SharedRequired `"}},
	// Only the opted-in Kind uses OnlyNew, so it keeps its name and gets
	// +required, with no copy.
	{name: "OnlyNew", want: []string{"// +required\n\tToken *string"}},
	// The status map holds the plain Zone, so it has to stay optional.
	{name: "Zone", want: []string{"ZoneName *string"}, notWant: []string{"+required"}},
	{name: "ZoneRequired", want: []string{"// +required\n\tZoneName *string"}},
}

func TestPlanRequiredStructs(t *testing.T) {
	for _, tc := range []struct {
		name  string
		files map[string]string
	}{
		{
			name: "the Spec of a Kind without the marker",
			files: map[string]string{"test/oldkind_types.go": "package test\n\n" +
				"type OldKindSpec struct {\n\tHolder *Holder `json:\"holder,omitempty\"`\n}\n"},
		},
		{
			name: "a hand-written status struct",
			files: map[string]string{"test/oldkind_types.go": "package test\n\n" +
				"type OldKindObservedState struct {\n\tHolder *Holder `json:\"holder,omitempty\"`\n}\n"},
		},
		{
			name: "a struct in another package",
			files: map[string]string{"other/v1/other_types.go": "package v1\n\n" +
				"import krm \"" + testImportPath + "\"\n\n" +
				"type OtherSpec struct {\n\tHolder *krm.Holder `json:\"holder,omitempty\"`\n}\n"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			g, body := planAndWrite(t, tc.files, newKindSpecFields(requiredSplitFixture(t)))

			// Assert
			checkStructs(t, body, splitWants)
			if structBody(body, "OnlyNewRequired") != "" {
				t.Errorf("OnlyNewRequired is declared, but nothing else uses OnlyNew:\n%s", body)
			}

			// The new Kind's Spec holds the Required copies.
			fields := requiredSplitFixture(t).Messages().ByName("NewKind").Fields()
			for field, want := range map[protoreflect.Name]string{
				"holder":      "*HolderRequired",
				"only_new":    "*OnlyNew",
				"shared_list": "[]SharedRequired",
				"zone":        "*ZoneRequired",
			} {
				got, err := GoTypeForField(fields.ByName(field), false, g.StrictWriteOptions())
				if err != nil || got != want {
					t.Errorf("GoTypeForField(%s) with the strict options = %q, %v, want %q", field, got, err, want)
				}
			}

			// Its ObservedState holds the plain Zone, which has no +required.
			details, ok := g.OutputFieldsFor("google.cloud.test.v1.NewKind")
			if !ok {
				t.Fatal("NewKind has no output fields")
			}
			var status bytes.Buffer
			WriteObservedStateFields(&status, details, g.ObservedStateMessages(), nil, g.StrictWriteOptions())
			if !strings.Contains(status.String(), "map[string]Zone `") || strings.Contains(status.String(), "Required") {
				t.Errorf("ObservedState should hold map[string]Zone:\n%s", status.String())
			}
		})
	}
}

// An opted-in Spec already on disk counts the same as one scaffolded in this
// run, and using a Required copy makes generate-types write it.
func TestPlanRequiredStructsMarkedSpec(t *testing.T) {
	files := map[string]string{
		"test/oldkind_types.go": "package test\n\n" +
			"type OldKindSpec struct {\n\tHolder *Holder `json:\"holder,omitempty\"`\n}\n",
		"test/newkind_types.go": "package test\n\n" +
			"// NewKindSpec defines the desired state of NewKind\n" +
			"// +kcc:spec:proto=google.cloud.test.v1.NewKind\n" +
			"// " + RequiredFromProtoMarker + "\n" +
			"type NewKindSpec struct {\n" +
			"\tHolder *HolderRequired `json:\"holder,omitempty\"`\n" +
			"\tOnlyNew *OnlyNew `json:\"onlyNew,omitempty\"`\n" +
			"\tSharedList []SharedRequired `json:\"sharedList,omitempty\"`\n" +
			"\tZone *Zone `json:\"zone,omitempty\"`\n" +
			"}\n",
	}

	// Act
	_, body := planAndWrite(t, files, nil)

	// Assert
	checkStructs(t, body, splitWants)
}

// Without an opted-in Kind, the flag changes nothing: no struct gets
// +required and none is copied.
func TestPlanRequiredStructsLeavesUnmarkedKindsAlone(t *testing.T) {
	files := map[string]string{
		"test/oldkind_types.go": "package test\n\n" +
			"type OldKindSpec struct {\n\tHolder *Holder `json:\"holder,omitempty\"`\n}\n",
		"test/newkind_types.go": "package test\n\n" +
			"type NewKindSpec struct {\n" +
			"\tHolder *Holder `json:\"holder,omitempty\"`\n" +
			"\tOnlyNew *OnlyNew `json:\"onlyNew,omitempty\"`\n" +
			"\tSharedList []Shared `json:\"sharedList,omitempty\"`\n" +
			"\tZone *Zone `json:\"zone,omitempty\"`\n" +
			"}\n",
	}

	// Act
	_, body := planAndWrite(t, files, nil)

	// Assert
	if strings.Contains(body, "+required") {
		t.Errorf("a struct got +required without a marked Kind:\n%s", body)
	}
	if strings.Contains(body, "Required struct {") {
		t.Errorf("a struct was copied without a marked Kind:\n%s", body)
	}
}

// If the package already declares <Name>Required by hand, the message is not
// split. Its struct stays optional for every Kind, so nothing breaks, and the
// judgement queue reports the fields.
func TestPlanRequiredStructsNameTaken(t *testing.T) {
	files := map[string]string{
		"test/oldkind_types.go": "package test\n\n" +
			"type OldKindSpec struct {\n\tHolder *Holder `json:\"holder,omitempty\"`\n}\n\n" +
			"type SharedRequired struct {\n\tOther *string `json:\"other,omitempty\"`\n}\n",
	}

	// Act
	_, body := planAndWrite(t, files, newKindSpecFields(requiredSplitFixture(t)))

	// Assert
	checkStructs(t, body, []structWant{
		{name: "Shared", notWant: []string{"+required"}},
		{name: "Holder", want: []string{"Shared *Shared `"}, notWant: []string{"+required"}},
		{name: "OnlyNew", want: []string{"// +required"}},
		{name: "ZoneRequired", want: []string{"// +required"}},
	})
	for _, name := range []string{"SharedRequired", "HolderRequired"} {
		if structBody(body, name) != "" {
			t.Errorf("%s was generated, but SharedRequired is declared by hand:\n%s", name, body)
		}
	}
}

// Code that uses a Required copy gets it, even if nothing else uses the
// message, or that code wouldn't compile. This is how a Kind that opts in by
// hand switches a field to the copy.
func TestPlanRequiredStructsNamedCopy(t *testing.T) {
	files := map[string]string{
		"test/newkind_types.go": "package test\n\n" +
			"// " + RequiredFromProtoMarker + "\n" +
			"type NewKindSpec struct {\n" +
			"\tOnlyNew *OnlyNewRequired `json:\"onlyNew,omitempty\"`\n" +
			"}\n",
	}

	// Act
	_, body := planAndWrite(t, files, nil)

	// Assert
	checkStructs(t, body, []structWant{
		{name: "OnlyNew", want: []string{"Token *string"}, notWant: []string{"+required"}},
		{name: "OnlyNewRequired", want: []string{"// +kcc:proto=google.cloud.test.v1.OnlyNew\n", "// +required\n\tToken *string"}},
	})
}

func TestScanRequiredUses(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	files := map[string]string{
		"test/marked_types.go": "package test\n\n" +
			"// " + RequiredFromProtoMarker + "\n" +
			"type MarkedSpec struct {\n\tA *A `json:\"a,omitempty\"`\n\tB []B `json:\"b,omitempty\"`\n\tC map[string]C `json:\"c,omitempty\"`\n}\n\n" +
			"type MarkedObservedState struct {\n\tD *D `json:\"d,omitempty\"`\n}\n",
		"test/plain_types.go": "package test\n\n" +
			"type PlainSpec struct {\n\tE *E `json:\"e,omitempty\"`\n}\n\n" +
			"type helper struct {\n\tF *F\n}\n",
		// Generated files are rewritten from the proto, so they are not uses.
		"test/types.generated.go":       "package test\n\ntype G struct {\n\tH *H\n}\n",
		"test/zz_generated.deepcopy.go": "package test\n\ntype I struct {\n\tJ *J\n}\n",
		"test/marked_test.go":           "package test\n\ntype K struct {\n\tL *L\n}\n",
		"other/v1/aliased.go": "package v1\n\nimport t \"" + testImportPath + "\"\n\n" +
			"type X struct {\n\tM *t.M\n}\n",
		"other/v2/unaliased.go": "package v2\n\nimport \"" + testImportPath + "\"\n\n" +
			"var _ = test.N{}\n",
		"other/v3/blank.go": "package v3\n\nimport _ \"" + testImportPath + "\"\n\n" +
			"type O struct{}\n",
	}
	for name, src := range files {
		p := filepath.Join(dir, name)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(src), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	// Act
	uses, err := ScanRequiredUses(filepath.Join(dir, "test"), dir, testImportPath)

	// Assert
	if err != nil {
		t.Fatalf("ScanRequiredUses: %v", err)
	}
	if want := map[string]bool{"Marked": true}; !reflect.DeepEqual(uses.MarkedKinds, want) {
		t.Errorf("MarkedKinds = %v, want %v", uses.MarkedKinds, want)
	}
	if want := map[string]bool{"A": true, "B": true, "C": true, "string": true}; !reflect.DeepEqual(uses.FromMarkedSpecs, want) {
		t.Errorf("FromMarkedSpecs = %v, want %v", uses.FromMarkedSpecs, want)
	}
	if want := map[string]bool{"D": true, "E": true, "F": true, "M": true, "N": true}; !reflect.DeepEqual(uses.FromElsewhere, want) {
		t.Errorf("FromElsewhere = %v, want %v", uses.FromElsewhere, want)
	}
}

func TestScanRequiredUsesWithoutPackage(t *testing.T) {
	// A service generated for the first time has no package directory yet.
	uses, err := ScanRequiredUses(filepath.Join(t.TempDir(), "missing"), "", testImportPath)
	if err != nil {
		t.Fatalf("ScanRequiredUses: %v", err)
	}
	if len(uses.MarkedKinds)+len(uses.FromMarkedSpecs)+len(uses.FromElsewhere) != 0 {
		t.Errorf("uses = %+v, want none", uses)
	}
}

func TestIsRequiredFromProtoMarker(t *testing.T) {
	for line, want := range map[string]bool{
		"// +kcc:required-from-proto":      true,
		"//+kcc:required-from-proto":       true,
		"\t// +kcc:required-from-proto  ":  true,
		"// +kcc:required-from-proto=true": false,
		"// +kcc:spec:proto=foo":           false,
		"// see +kcc:required-from-proto":  false,
	} {
		if got := IsRequiredFromProtoMarker(line); got != want {
			t.Errorf("IsRequiredFromProtoMarker(%q) = %v, want %v", line, got, want)
		}
	}
}
