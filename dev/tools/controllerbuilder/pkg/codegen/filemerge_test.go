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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestParseGeneratedGoBlocks(t *testing.T) {
	src := []byte(`// Copyright 2026 Google LLC

package v1alpha1

import (
	"fmt"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/foo/v1alpha1"
)

// +kcc:proto=google.cloud.foo.v1.Alpha
type Alpha struct {
	Name *string
}

/* unreachable type Beta
// +kcc:proto=google.cloud.foo.v1.Beta
type Beta struct {
}
*/

// +kcc:observedstate:proto=google.cloud.foo.v1.Alpha
type AlphaObservedState struct {
	State *string
}

func Alpha_FromProto(in *string) *krm.Alpha {
	_ = fmt.Sprintf
	return nil
}

/*
func Beta_ToProto(mapCtx *direct.MapContext, in *krm.Beta) *string {
	return nil
}
*/
`)

	blocks, imports, err := parseGeneratedGoBlocks(src, "v1alpha1", false)
	if err != nil {
		t.Fatalf("parseGeneratedGoBlocks failed: %v", err)
	}

	wantImports := map[string]string{
		"fmt": "",
		"github.com/GoogleCloudPlatform/k8s-config-connector/apis/foo/v1alpha1": "krm",
	}
	if diff := cmp.Diff(wantImports, imports); diff != "" {
		t.Errorf("imports mismatch (-want +got):\n%s", diff)
	}

	wantBlocks := []generatedBlock{
		{
			name:     "Alpha",
			kind:     "type",
			protoFQN: "google.cloud.foo.v1.Alpha",
		},
		{
			name:        "Beta",
			kind:        "type",
			isCommented: true,
			protoFQN:    "google.cloud.foo.v1.Beta",
		},
		{
			name:            "AlphaObservedState",
			kind:            "type",
			protoFQN:        "google.cloud.foo.v1.Alpha",
			isObservedState: true,
		},
		{
			name:        "Alpha_FromProto",
			kind:        "func",
			krmTypeName: "Alpha",
		},
		{
			name:        "Beta_ToProto",
			kind:        "func",
			isCommented: true,
			krmTypeName: "Beta",
		},
	}
	ignoreFields := cmpopts.IgnoreFields(generatedBlock{}, "rawText", "offset")
	if diff := cmp.Diff(wantBlocks, blocks, cmp.AllowUnexported(generatedBlock{}), ignoreFields); diff != "" {
		t.Errorf("blocks mismatch (-want +got):\n%s", diff)
	}
}

func TestShouldKeepExistingBlock(t *testing.T) {
	f := &generatedFile{
		reservedTypeNames: map[string]bool{"ReservedType": true},
	}
	pkgDecls := packageDeclarations{
		types:     map[string]bool{"HandwrittenType": true},
		protoTags: map[string]bool{"google.cloud.foo.v1.ClaimedProto": true},
		funcs:     map[string]bool{"Handwritten_FromProto": true},
	}
	validKRM := map[string]bool{"KeptSpec": true}

	tests := []struct {
		name  string
		block generatedBlock
		want  bool
	}{
		{
			name:  "active type with no collision is kept",
			block: generatedBlock{name: "ActiveType", kind: "type", protoFQN: "google.cloud.foo.v1.ActiveType"},
			want:  true,
		},
		{
			name:  "active type colliding with handwritten type is dropped",
			block: generatedBlock{name: "HandwrittenType", kind: "type"},
			want:  false,
		},
		{
			name:  "active type in reservedTypeNames is dropped",
			block: generatedBlock{name: "ReservedType", kind: "type"},
			want:  false,
		},
		{
			name:  "active base struct whose protoFQN is claimed in handwritten file is dropped",
			block: generatedBlock{name: "ClaimedProtoObservedState", kind: "type", protoFQN: "google.cloud.foo.v1.ClaimedProto", isObservedState: true},
			want:  false,
		},
		{
			name:  "commented type is kept even if handwritten type exists",
			block: generatedBlock{name: "HandwrittenType", kind: "type", isCommented: true},
			want:  true,
		},
		{
			name:  "active func colliding with handwritten func is dropped",
			block: generatedBlock{name: "Handwritten_FromProto", kind: "func", krmTypeName: "KeptSpec"},
			want:  false,
		},
		{
			name:  "func referencing pruned KRM type is dropped",
			block: generatedBlock{name: "PrunedSpec_FromProto", kind: "func", krmTypeName: "PrunedSpec"},
			want:  false,
		},
		{
			name:  "active func referencing valid KRM type is kept",
			block: generatedBlock{name: "KeptSpec_FromProto", kind: "func", krmTypeName: "KeptSpec"},
			want:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := f.shouldKeepExistingBlock(tc.block, pkgDecls, validKRM); got != tc.want {
				t.Errorf("shouldKeepExistingBlock(%+v) = %v, want %v", tc.block, got, tc.want)
			}
		})
	}
}

func TestMergeExistingImports(t *testing.T) {
	f := &generatedFile{}
	f.addImport("pb", "cloud.google.com/go/foo/sub/apiv1/subpb")
	f.addImport("krmv1alpha1", "github.com/GoogleCloudPlatform/k8s-config-connector/apis/foo/v1alpha1")

	existingImports := map[string]string{
		"cloud.google.com/go/foo/apiv1/foopb":                                       "pb",
		"github.com/GoogleCloudPlatform/k8s-config-connector/apis/foo/v1alpha1":     "krm",
		"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct": "",
	}

	renames := f.mergeExistingImports(existingImports)

	wantRenames := map[string]string{
		"pb":  "pb_2",
		"krm": "krmv1alpha1",
	}
	if diff := cmp.Diff(wantRenames, renames); diff != "" {
		t.Errorf("mergeExistingImports renames mismatch (-want +got):\n%s", diff)
	}

	wantImports := map[string]string{
		"cloud.google.com/go/foo/sub/apiv1/subpb":                                   "pb",
		"cloud.google.com/go/foo/apiv1/foopb":                                       "pb_2",
		"github.com/GoogleCloudPlatform/k8s-config-connector/apis/foo/v1alpha1":     "krmv1alpha1",
		"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct": "",
	}
	if diff := cmp.Diff(wantImports, f.imports); diff != "" {
		t.Errorf("f.imports mismatch (-want +got):\n%s", diff)
	}
}

func TestRewritePackageQualifiers(t *testing.T) {
	src := `func Foo_FromProto(in *pb.Foo) *krm.Foo {
	// selector obj.pb.Bar must not be rewritten
	out := &krm.Foo{}
	out.Bar = pb.Bar_FromProto(in.pb.Bar)
	return out
}`
	got := rewritePackageQualifiers(src, map[string]string{
		"pb":  "pb_2",
		"krm": "krmv1alpha1",
	})
	want := `func Foo_FromProto(in *pb_2.Foo) *krmv1alpha1.Foo {
	// selector obj.pb.Bar must not be rewritten
	out := &krmv1alpha1.Foo{}
	out.Bar = pb_2.Bar_FromProto(in.pb.Bar)
	return out
}`
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("rewritePackageQualifiers mismatch (-want +got):\n%s", diff)
	}
}

func TestMergeAndSortTypeBlocks(t *testing.T) {
	existing := []generatedBlock{
		{name: "Zeta", kind: "type", protoFQN: "pkg.v1.Zeta", rawText: "type Zeta struct{}"},
		{name: "AlphaObservedState", kind: "type", protoFQN: "pkg.v1.Alpha", isObservedState: true, rawText: "type AlphaObservedState struct{}"},
		{name: "Shared", kind: "type", protoFQN: "pkg.v1.sub1.Shared", rawText: "type Shared struct{ A string }"},
	}
	current := []generatedBlock{
		{name: "Alpha", kind: "type", protoFQN: "pkg.v1.Alpha", rawText: "type Alpha struct{}"},
		{name: "Shared", kind: "type", isCommented: true, protoFQN: "pkg.v1.sub2.Shared", rawText: "/* unreachable type Shared */"},
	}

	got := mergeAndSortTypeBlocks(existing, current)
	want := []generatedBlock{
		{name: "Alpha", kind: "type", protoFQN: "pkg.v1.Alpha", rawText: "type Alpha struct{}"},
		{name: "Zeta", kind: "type", protoFQN: "pkg.v1.Zeta", rawText: "type Zeta struct{}"},
		{name: "Shared", kind: "type", protoFQN: "pkg.v1.sub1.Shared", rawText: "type Shared struct{ A string }"},
		{name: "AlphaObservedState", kind: "type", protoFQN: "pkg.v1.Alpha", isObservedState: true, rawText: "type AlphaObservedState struct{}"},
	}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(generatedBlock{})); diff != "" {
		t.Errorf("mergeAndSortTypeBlocks mismatch (-want +got):\n%s", diff)
	}
}

func TestMergeOrderedBlocks(t *testing.T) {
	existing := []generatedBlock{
		{name: "FuncA", kind: "func", rawText: "func FuncA() { /* v1 */ }"},
		{name: "FuncC", kind: "func", rawText: "func FuncC() {}"},
	}
	current := []generatedBlock{
		{name: "FuncA", kind: "func", rawText: "func FuncA() { /* v2 */ }"},
		{name: "FuncB", kind: "func", rawText: "func FuncB() {}"},
	}

	got := mergeOrderedBlocks(existing, current)
	want := []generatedBlock{
		{name: "FuncA", kind: "func", rawText: "func FuncA() { /* v2 */ }"},
		{name: "FuncB", kind: "func", rawText: "func FuncB() {}"},
		{name: "FuncC", kind: "func", rawText: "func FuncC() {}"},
	}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(generatedBlock{})); diff != "" {
		t.Errorf("mergeOrderedBlocks mismatch (-want +got):\n%s", diff)
	}
}

func TestExtractKRMTypeFromFuncSignature(t *testing.T) {
	tests := []struct {
		sig  string
		want string
	}{
		{
			sig:  "func WidgetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Widget) *krm.WidgetSpec {",
			want: "WidgetSpec",
		},
		{
			sig:  "func WidgetSpec_ToProto(mapCtx *direct.MapContext, in *krm.WidgetSpec) *pb.Widget {",
			want: "WidgetSpec",
		},
		{
			sig:  "func WidgetSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.WidgetSpec) *pb.Widget {",
			want: "WidgetSpec",
		},
		{
			sig:  "func HelperFunc(x int) int {",
			want: "",
		},
	}
	for _, tc := range tests {
		if got := extractKRMTypeFromFuncSignature(tc.sig); got != tc.want {
			t.Errorf("extractKRMTypeFromFuncSignature(%q) = %q, want %q", tc.sig, got, tc.want)
		}
	}
}

func TestMergeFileAnnotations(t *testing.T) {
	existing := &annotations.FileAnnotation{
		Key: "+generated:types",
		Attributes: map[string][]string{
			"krm.group":     {"example.cnrm.cloud.google.com"},
			"proto.service": {"google.example.v1"},
			"resource":      {"ExampleAlpha:Alpha"},
		},
	}
	current := &annotations.FileAnnotation{
		Key: "+generated:types",
		Attributes: map[string][]string{
			"krm.group":     {"example.cnrm.cloud.google.com"},
			"proto.service": {"google.example.v1", "google.example. sub.v1"},
			"resource":      {"ExampleBeta:Beta"},
		},
	}

	got := mergeFileAnnotations(existing, current)
	want := &annotations.FileAnnotation{
		Key: "+generated:types",
		Attributes: map[string][]string{
			"krm.group":     {"example.cnrm.cloud.google.com"},
			"proto.service": {"google.example.v1", "google.example. sub.v1"},
			"resource":      {"ExampleAlpha:Alpha", "ExampleBeta:Beta"},
		},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mergeFileAnnotations mismatch (-want +got):\n%s", diff)
	}
}

func TestHasVersionedMapperFuncs(t *testing.T) {
	unversioned := []generatedBlock{{name: "Instance_FromProto", kind: "func"}}
	versioned := []generatedBlock{{name: "Instance_v1beta1_FromProto", kind: "func"}}

	if hasVersionedMapperFuncs(unversioned) {
		t.Errorf("expected unversioned mapper func to return false")
	}
	if !hasVersionedMapperFuncs(versioned) {
		t.Errorf("expected versioned mapper func to return true")
	}
}
