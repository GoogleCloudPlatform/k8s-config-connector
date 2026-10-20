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
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
	"github.com/google/go-cmp/cmp"
)

func TestMergeWithExisting(t *testing.T) {
	tests := []struct {
		name         string
		fileName     string
		existingFile string
		setupDir     func(t *testing.T, dir string)
		currentAnn   *annotations.FileAnnotation
		currentImps  map[string]string
		currentBody  string
		wantContains []string
		wantExcluded []string
		wantImps     map[string]string
		wantAnnAttrs map[string][]string
	}{
		{
			name:     "merges blocks, annotations, and disambiguates colliding import aliases",
			fileName: "mapper.generated.go",
			existingFile: `// +generated:mapper
// krm.group: foo.cnrm.cloud.google.com
// proto.service: google.cloud.foo.v1

package foo

import (
	pb "cloud.google.com/go/foo/apiv1/foopb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/foo/v1alpha1"
)

func Alpha_FromProto(in *pb.Alpha) *krm.Alpha {
	// selector in.pb.Field must not be rewritten
	_ = in.pb.Field
	return &krm.Alpha{}
}
`,
			currentAnn: &annotations.FileAnnotation{
				Key: "+generated:mapper",
				Attributes: map[string][]string{
					"krm.group":     {"foo.cnrm.cloud.google.com"},
					"proto.service": {"google.cloud.foo.sub.v1"},
				},
			},
			currentImps: map[string]string{
				"cloud.google.com/go/foo/sub/apiv1/subpb":                               "pb",
				"github.com/GoogleCloudPlatform/k8s-config-connector/apis/foo/v1alpha1": "krmv1alpha1",
			},
			currentBody: `
func Beta_FromProto(in *pb.Beta) *krmv1alpha1.Beta {
	return &krmv1alpha1.Beta{}
}
`,
			wantContains: []string{
				"func Alpha_FromProto(in *pb_2.Alpha) *krmv1alpha1.Alpha",
				"_ = in.pb.Field",
				"func Beta_FromProto(in *pb.Beta) *krmv1alpha1.Beta",
			},
			wantImps: map[string]string{
				"cloud.google.com/go/foo/apiv1/foopb":                                   "pb_2",
				"cloud.google.com/go/foo/sub/apiv1/subpb":                               "pb",
				"github.com/GoogleCloudPlatform/k8s-config-connector/apis/foo/v1alpha1": "krmv1alpha1",
			},
			wantAnnAttrs: map[string][]string{
				"krm.group":     {"foo.cnrm.cloud.google.com"},
				"proto.service": {"google.cloud.foo.v1", "google.cloud.foo.sub.v1"},
			},
		},
		{
			name:     "skips merge when existing file only contains commented-out blocks",
			fileName: "types.generated.go",
			existingFile: `package v1alpha1

/* unreachable type OldType
type OldType struct {}
*/
`,
			currentBody: `
type NewType struct {}
`,
			wantContains: []string{"type NewType struct"},
			wantExcluded: []string{"OldType"},
		},
		{
			name:     "skips merge when resource belongs to a separate *_types.generated.go file",
			fileName: "types.generated.go",
			existingFile: `package v1alpha1

type SharedType struct {}
`,
			setupDir: func(t *testing.T, dir string) {
				sep := `// +generated:types
// resource: FooSeparate:Separate

package v1alpha1
`
				if err := os.WriteFile(filepath.Join(dir, "separate_types.generated.go"), []byte(sep), 0644); err != nil {
					t.Fatalf("writing separate file: %v", err)
				}
			},
			currentAnn: &annotations.FileAnnotation{
				Key:        "+generated:types",
				Attributes: map[string][]string{"resource": {"FooSeparate:Separate"}},
			},
			currentBody: `
type SeparateSpec struct {}
`,
			wantContains: []string{"type SeparateSpec struct"},
			wantExcluded: []string{"SharedType"},
		},
		{
			name:     "skips merge when mapper versioning convention changes",
			fileName: "mapper.generated.go",
			existingFile: `package foo

func Alpha_v1alpha1_FromProto(in *string) *string { return nil }
`,
			currentBody: `
func Alpha_FromProto(in *string) *string { return nil }
`,
			wantContains: []string{"func Alpha_FromProto"},
			wantExcluded: []string{"Alpha_v1alpha1_FromProto"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			if tc.setupDir != nil {
				tc.setupDir(t, dir)
			}

			f := &generatedFile{
				key:            generatedFileKey{FileName: tc.fileName},
				goPackage:      "v1alpha1",
				fileAnnotation: tc.currentAnn,
			}

			for pkg, alias := range tc.currentImps {
				f.addImport(alias, pkg)
			}

			f.body.WriteString(tc.currentBody)
			f.mergeWithExisting([]byte(tc.existingFile), dir)

			got := f.body.String()
			for _, w := range tc.wantContains {
				if !strings.Contains(got, w) {
					t.Errorf("expected body to contain %q, got:\n%s", w, got)
				}
			}

			for _, ex := range tc.wantExcluded {
				if strings.Contains(got, ex) {
					t.Errorf("expected body not to contain %q, got:\n%s", ex, got)
				}
			}

			if tc.wantImps != nil {
				if diff := cmp.Diff(tc.wantImps, f.imports); diff != "" {
					t.Errorf("imports mismatch (-want +got):\n%s", diff)
				}
			}

			if tc.wantAnnAttrs != nil {
				if diff := cmp.Diff(tc.wantAnnAttrs, f.fileAnnotation.Attributes); diff != "" {
					t.Errorf("annotation attributes mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestShouldKeepExistingBlock(t *testing.T) {
	baseDir := t.TempDir()
	apiDir := filepath.Join(baseDir, "apis", "foo", "v1alpha1")
	if err := os.MkdirAll(apiDir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	if err := os.WriteFile(filepath.Join(apiDir, "types.generated.go"), []byte(`package v1alpha1
type KeptSpec struct{}
/* unreachable:
type PrunedSpec struct{}
*/
`), 0644); err != nil {
		t.Fatalf("write types.generated.go: %v", err)
	}

	f := &generatedFile{
		baseDir:           baseDir,
		key:               generatedFileKey{FileName: "mapper.generated.go"},
		reservedTypeNames: map[string]bool{"ReservedType": true},
		imports: map[string]string{
			"github.com/GoogleCloudPlatform/k8s-config-connector/apis/foo/v1alpha1": "krm",
		},
	}

	pkgDecls := packageDeclarations{
		types:     map[string]bool{"HandwrittenType": true},
		protoTags: map[string]bool{"google.cloud.foo.v1.ClaimedProto": true},
		funcs:     map[string]bool{"Handwritten_FromProto": true},
	}

	validKRM := f.resolveValidKRMTypes(nil)

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

func TestShouldCurrentReplaceExistingPreservesNonRequiredTypeAndMapHelper(t *testing.T) {
	nonReqType := generatedBlock{
		name:    "GCSSource",
		kind:    "type",
		rawText: "type GCSSource struct {\n\tUris []string\n}",
	}
	reqType := generatedBlock{
		name:    "GCSSource",
		kind:    "type",
		rawText: "type GCSSource struct {\n\t// +required\n\tURIs []string\n}",
	}
	if shouldCurrentReplaceExisting(nonReqType, reqType) {
		t.Errorf("expected non-required type to be preserved over +required type")
	}

	helperFunc := generatedBlock{
		name:    "DataplexDataProductSpec_FromProto",
		kind:    "func",
		rawText: "func DataplexDataProductSpec_FromProto() {\n\tout.AccessGroups = AccessGroups_FromProto(mapCtx, in.AccessGroups)\n}",
	}
	inlineMapFunc := generatedBlock{
		name:    "DataplexDataProductSpec_FromProto",
		kind:    "func",
		rawText: "func DataplexDataProductSpec_FromProto() {\n\tout.AccessGroups = make(map[string]krm.DataProduct_AccessGroup, len(in.AccessGroups))\n}",
	}
	if shouldCurrentReplaceExisting(helperFunc, inlineMapFunc) {
		t.Errorf("expected custom map helper call in existing mapper func to be preserved over inline map loop")
	}
}

