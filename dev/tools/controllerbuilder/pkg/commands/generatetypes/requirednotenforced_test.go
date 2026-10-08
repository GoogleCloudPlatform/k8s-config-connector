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
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

// requiredGapFixture registers a resource with REQUIRED fields at the top
// and in a nested message.
func requiredGapFixture(t *testing.T) *protoregistry.Files {
	t.Helper()
	sp := func(s string) *string { return &s }
	ip := func(i int32) *int32 { return &i }
	required := &descriptorpb.FieldOptions{}
	proto.SetExtension(required, annotations.E_FieldBehavior, []annotations.FieldBehavior{annotations.FieldBehavior_REQUIRED})
	str := func(name string, num int32, opts *descriptorpb.FieldOptions) *descriptorpb.FieldDescriptorProto {
		return &descriptorpb.FieldDescriptorProto{Name: sp(name), Number: ip(num), Type: descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(), Options: opts}
	}
	msg := func(name string, num int32, typeName string, repeated bool) *descriptorpb.FieldDescriptorProto {
		f := &descriptorpb.FieldDescriptorProto{Name: sp(name), Number: ip(num), Type: descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(), TypeName: sp(".google.cloud.test.v1." + typeName)}
		if repeated {
			f.Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
		}
		return f
	}
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    sp("gaps.proto"),
		Package: sp("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: sp("Inner"), Field: []*descriptorpb.FieldDescriptorProto{str("a", 1, required), str("b", 2, nil)}},
			{
				Name: sp("Thing"),
				NestedType: []*descriptorpb.DescriptorProto{{
					Name:    sp("InnerByKeyEntry"),
					Field:   []*descriptorpb.FieldDescriptorProto{str("key", 1, nil), msg("value", 2, "Inner", false)},
					Options: &descriptorpb.MessageOptions{MapEntry: proto.Bool(true)},
				}},
				Field: []*descriptorpb.FieldDescriptorProto{
					str("name", 1, required),
					str("saas", 2, required),
					msg("inner", 3, "Inner", false),
					msg("inners", 4, "Inner", true),
					msg("inner_by_key", 5, "Thing.InnerByKeyEntry", true),
					str("network", 6, required),
					str("done", 7, required),
				},
			},
		},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	files := new(protoregistry.Files)
	if err := files.RegisterFile(fd); err != nil {
		t.Fatalf("registering file: %v", err)
	}
	return files
}

// thingTypes is a hand-written thing_types.go. inner is the struct its
// nested fields hold.
func thingTypes(version string, marked bool, inner string) string {
	marker := ""
	if marked {
		marker = "// " + codegen.RequiredFromProtoMarker + "\n"
	}
	return "package " + version + "\n\n" +
		"// ThingSpec defines the desired state of Thing\n" +
		"// +kcc:spec:proto=google.cloud.test.v1.Thing\n" +
		marker +
		"type ThingSpec struct {\n" +
		"\t// +kcc:proto:field=google.cloud.test.v1.Thing.saas\n" +
		"\tSaas *string `json:\"saas,omitempty\"`\n\n" +
		"\t// +kcc:proto:field=google.cloud.test.v1.Thing.inner\n" +
		"\tInner *" + inner + " `json:\"inner,omitempty\"`\n\n" +
		// No annotation on these two: they are matched by name.
		"\tInners []" + inner + " `json:\"inners,omitempty\"`\n\n" +
		"\tInnerByKey map[string]" + inner + " `json:\"innerByKey,omitempty\"`\n\n" +
		"\tNetworkRef *refsv1beta1.ComputeNetworkRef `json:\"networkRef,omitempty\"`\n\n" +
		"\t// +kcc:proto:field=google.cloud.test.v1.Thing.done\n" +
		"\t// +required\n" +
		"\tDone *string `json:\"done,omitempty\"`\n\n" +
		"\tResourceID *string `json:\"resourceID,omitempty\"`\n" +
		"}\n"
}

// generatedTypes is types.generated.go. It has Inner unless prune commented
// it out, and InnerRequired when generate-types wrote the copy.
func generatedTypes(version string, plain, requiredCopy bool) string {
	src := "package " + version + "\n"
	if plain {
		src += "\n// +kcc:proto=google.cloud.test.v1.Inner\n" +
			"type Inner struct {\n" +
			"\t// +kcc:proto:field=google.cloud.test.v1.Inner.a\n" +
			"\tA *string `json:\"a,omitempty\"`\n\n" +
			"\t// +kcc:proto:field=google.cloud.test.v1.Inner.b\n" +
			"\tB *string `json:\"b,omitempty\"`\n" +
			"}\n"
	}
	if requiredCopy {
		src += "\n// +kcc:proto=google.cloud.test.v1.Inner\n" +
			"type InnerRequired struct {\n" +
			"\t// +kcc:proto:field=google.cloud.test.v1.Inner.a\n" +
			"\t// +required\n" +
			"\tA *string `json:\"a,omitempty\"`\n\n" +
			"\t// +kcc:proto:field=google.cloud.test.v1.Inner.b\n" +
			"\tB *string `json:\"b,omitempty\"`\n" +
			"}\n"
	}
	return src
}

func TestRequiredNotEnforced(t *testing.T) {
	const (
		handWritten = "To enforce it, add +required to ThingSpec."
		addMarker   = "To enforce it, add +kcc:required-from-proto to ThingSpec and run generate.sh again."
		useCopy     = " in thing_types.go from Inner to InnerRequired and run generate.sh again."
		alpha       = "Thing is alpha, so it may be enforced"
		beta        = "Keep it optional: Thing is v1beta1"
	)
	// The marked Kind holds the plain Inner, with or without a copy on disk:
	// generate-types writes the copy once code names it.
	switchToCopy := map[string][]string{
		".spec.saas":             {handWritten + "Saas in thing_types.go.", alpha},
		".spec.inner.a":          {"To enforce it, change ThingSpec.Inner" + useCopy, alpha},
		".spec.inners[].a":       {"To enforce it, change ThingSpec.Inners" + useCopy, alpha},
		".spec.innerByKey.KEY.a": {"To enforce it, change ThingSpec.InnerByKey" + useCopy, alpha},
		".spec.networkRef":       {handWritten + "NetworkRef in thing_types.go.", alpha},
	}
	// Only the hand-written fields are left once the Kind holds the copy.
	handWrittenOnly := map[string][]string{
		".spec.saas":       {handWritten + "Saas in thing_types.go."},
		".spec.networkRef": {handWritten + "NetworkRef in thing_types.go."},
	}
	for _, tc := range []struct {
		name      string
		goPackage string
		files     map[string]string
		// want maps each field path to text its detail must contain.
		want map[string][]string
	}{
		{
			name:      "an alpha Kind without the marker",
			goPackage: "test/v1alpha1",
			files: map[string]string{
				"test/v1alpha1/thing_types.go":     thingTypes("v1alpha1", false, "Inner"),
				"test/v1alpha1/types.generated.go": generatedTypes("v1alpha1", true, false),
			},
			want: map[string][]string{
				".spec.saas":             {handWritten + "Saas in thing_types.go.", alpha},
				".spec.inner.a":          {addMarker, alpha},
				".spec.inners[].a":       {addMarker, alpha},
				".spec.innerByKey.KEY.a": {addMarker, alpha},
				".spec.networkRef":       {handWritten + "NetworkRef in thing_types.go.", alpha},
			},
		},
		{
			name:      "a marked Kind that holds the plain struct",
			goPackage: "test/v1alpha1",
			files: map[string]string{
				"test/v1alpha1/thing_types.go":     thingTypes("v1alpha1", true, "Inner"),
				"test/v1alpha1/types.generated.go": generatedTypes("v1alpha1", true, true),
			},
			want: switchToCopy,
		},
		{
			name:      "a marked Kind that holds the plain struct, with no copy yet",
			goPackage: "test/v1alpha1",
			files: map[string]string{
				"test/v1alpha1/thing_types.go":     thingTypes("v1alpha1", true, "Inner"),
				"test/v1alpha1/types.generated.go": generatedTypes("v1alpha1", true, false),
			},
			want: switchToCopy,
		},
		{
			name:      "a marked Kind that holds the Required copy",
			goPackage: "test/v1alpha1",
			files: map[string]string{
				"test/v1alpha1/thing_types.go":     thingTypes("v1alpha1", true, "InnerRequired"),
				"test/v1alpha1/types.generated.go": generatedTypes("v1alpha1", true, true),
			},
			want: handWrittenOnly,
		},
		{
			name:      "a marked Kind that holds the Required copy, and nothing holds the plain struct",
			goPackage: "test/v1alpha1",
			files: map[string]string{
				"test/v1alpha1/thing_types.go":     thingTypes("v1alpha1", true, "InnerRequired"),
				"test/v1alpha1/types.generated.go": generatedTypes("v1alpha1", false, true),
			},
			want: handWrittenOnly,
		},
		{
			name:      "a marked Kind whose struct cannot get a copy",
			goPackage: "test/v1alpha1",
			files: map[string]string{
				"test/v1alpha1/thing_types.go":     thingTypes("v1alpha1", true, "Inner"),
				"test/v1alpha1/types.generated.go": generatedTypes("v1alpha1", true, false),
				"test/v1alpha1/other_types.go":     "package v1alpha1\n\ntype InnerRequired struct {\n\tOther *string `json:\"other,omitempty\"`\n}\n",
			},
			want: map[string][]string{
				".spec.saas":             {handWritten + "Saas in thing_types.go."},
				".spec.inner.a":          {"Inner has no Required copy because another type is called InnerRequired."},
				".spec.inners[].a":       {"Inner has no Required copy because another type is called InnerRequired."},
				".spec.innerByKey.KEY.a": {"Inner has no Required copy because another type is called InnerRequired."},
				".spec.networkRef":       {handWritten + "NetworkRef in thing_types.go."},
			},
		},
		{
			name:      "a beta Kind",
			goPackage: "test/v1beta1",
			files: map[string]string{
				"test/v1beta1/thing_types.go":     thingTypes("v1beta1", false, "Inner"),
				"test/v1beta1/types.generated.go": generatedTypes("v1beta1", true, false),
			},
			want: map[string][]string{
				".spec.saas":             {beta},
				".spec.inner.a":          {beta},
				".spec.inners[].a":       {beta},
				".spec.innerByKey.KEY.a": {beta},
				".spec.networkRef":       {beta},
			},
		},
		{
			name:      "an alpha version of a Kind that is also beta",
			goPackage: "test/v1alpha1",
			files: map[string]string{
				"test/v1alpha1/thing_types.go":     thingTypes("v1alpha1", false, "Inner"),
				"test/v1alpha1/types.generated.go": generatedTypes("v1alpha1", true, false),
				"test/v1beta1/thing_types.go":      "package v1beta1\n\ntype Thing struct {\n}\n",
			},
			want: map[string][]string{
				".spec.saas":             {beta},
				".spec.inner.a":          {beta},
				".spec.inners[].a":       {beta},
				".spec.innerByKey.KEY.a": {beta},
				".spec.networkRef":       {beta},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			dir := t.TempDir()
			for name, src := range tc.files {
				p := filepath.Join(dir, name)
				if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(p, []byte(src), 0o644); err != nil {
					t.Fatal(err)
				}
			}

			// Act
			entries, err := requiredNotEnforced(requiredGapFixture(t), dir, tc.goPackage, "test.cnrm.cloud.google.com",
				[]string{"Thing", "Missing"}, map[string]string{"Thing": "google.cloud.test.v1.Thing"})

			// Assert
			if err != nil {
				t.Fatalf("requiredNotEnforced: %v", err)
			}
			got := map[string]string{}
			for _, e := range entries {
				if err := e.Validate(); err != nil {
					t.Errorf("entry %+v does not validate: %v", e, err)
				}
				if e.Kind != "Thing" || e.Group != "test.cnrm.cloud.google.com" || e.Reason != judgement.ReasonRequiredNotEnforced || e.Status != judgement.StatusOpen {
					t.Errorf("unexpected entry %+v", e)
				}
				got[e.Field] = e.Detail
			}
			gotPaths, wantPaths := sortedPaths(got), sortedPaths(tc.want)
			if !reflect.DeepEqual(gotPaths, wantPaths) {
				t.Errorf("paths = %v, want %v", gotPaths, wantPaths)
			}
			for path, wants := range tc.want {
				for _, w := range wants {
					if !strings.Contains(got[path], w) {
						t.Errorf("detail for %s = %q, want it to contain %q", path, got[path], w)
					}
				}
			}
		})
	}
}

func sortedPaths[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func TestKRMNamesFor(t *testing.T) {
	for j, want := range map[string][]string{
		"network":     {"network", "networkRef", "networkRefs", "networkRefs"},
		"kmsKeyName":  {"kmsKeyName", "kmsKeyNameRef", "kmsKeyNameRefs", "kmsKeyNameRefs", "kmsKeyRef"},
		"subnetworks": {"subnetworks", "subnetworksRef", "subnetworksRefs", "subnetworkRefs"},
		"projectIds":  {"projectIds", "projectIdsRef", "projectIdsRefs", "projectIdRefs", "projectRefs"},
	} {
		if got := krmNamesFor(j); !reflect.DeepEqual(got, want) {
			t.Errorf("krmNamesFor(%q) = %v, want %v", j, got, want)
		}
	}
}
