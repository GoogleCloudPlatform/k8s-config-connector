// Copyright 2024 Google LLC
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

package v1beta1_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestArtifactRegistryRepositoryRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid_reference",
			ref:     "projects/my-project/locations/us-central1/repositories/my-repo",
			wantErr: false,
		},
		{
			name:    "invalid_prefix",
			ref:     "invalid/my-project/locations/us-central1/repositories/my-repo",
			wantErr: true,
		},
		{
			name:    "missing_location",
			ref:     "projects/my-project/repositories/my-repo",
			wantErr: true,
		},
		{
			name:    "missing_repository",
			ref:     "projects/my-project/locations/us-central1",
			wantErr: true,
		},
		{
			name:    "empty_string",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := &v1beta1.ArtifactRegistryRepositoryRef{}
			err := r.ValidateExternal(tc.ref)

			if (err != nil) != tc.wantErr {
				t.Errorf("ValidateExternal(%q) error = %v, wantErr %v", tc.ref, err, tc.wantErr)
			}
		})
	}
}

func TestArtifactRegistryRepositoryRef_Normalize_Success(t *testing.T) {
	tests := []struct {
		name         string
		ref          v1beta1.ArtifactRegistryRepositoryRef
		objs         []client.Object
		wantExternal string
	}{
		{
			name: "canonicalizes_host_prefixed_external",
			ref: v1beta1.ArtifactRegistryRepositoryRef{
				External: "//artifactregistry.googleapis.com/projects/my-proj/locations/us-west1/repositories/my-repo",
			},
			wantExternal: "projects/my-proj/locations/us-west1/repositories/my-repo",
		},
		{
			name: "canonicalizes_https_external",
			ref: v1beta1.ArtifactRegistryRepositoryRef{
				External: "https://artifactregistry.googleapis.com/projects/my-proj/locations/us-west1/repositories/my-repo",
			},
			wantExternal: "projects/my-proj/locations/us-west1/repositories/my-repo",
		},
		{
			name: "resolves_ready_repository_with_short_status_name",
			ref: v1beta1.ArtifactRegistryRepositoryRef{
				Name: "k8s-upstream",
			},
			objs: []client.Object{
				newUnstructuredARRepo("k8s-upstream", "test-ns", "my-proj", "us-west1", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "True"},
					},
					"name": "server-repo-id",
				}),
			},
			wantExternal: "projects/my-proj/locations/us-west1/repositories/server-repo-id",
		},
		{
			name: "resolves_ready_repository_with_full_path_status_name",
			ref: v1beta1.ArtifactRegistryRepositoryRef{
				Name: "k8s-upstream",
			},
			objs: []client.Object{
				newUnstructuredARRepo("k8s-upstream", "test-ns", "", "", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "True"},
					},
					"name": "projects/status-proj/locations/europe-west1/repositories/full-repo",
				}),
			},
			wantExternal: "projects/status-proj/locations/europe-west1/repositories/full-repo",
		},
		{
			name: "resolves_repository_with_ready_false_update_failed_from_status_name",
			ref: v1beta1.ArtifactRegistryRepositoryRef{
				Name: "upstream",
			},
			objs: []client.Object{
				newUnstructuredARRepo("upstream", "test-ns", "my-proj", "us-west1", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "False", "reason": "UpdateFailed"},
					},
					"name": "upstream",
				}),
			},
			wantExternal: "projects/my-proj/locations/us-west1/repositories/upstream",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).WithObjects(tc.objs...).Build()
			r := tc.ref
			err := r.Normalize(t.Context(), reader, "test-ns")

			if err != nil {
				t.Fatalf("Normalize(%+v) unexpected error: %v", tc.ref, err)
			}
			if r.External != tc.wantExternal {
				t.Errorf("Normalize(%+v) External = %q, want %q", tc.ref, r.External, tc.wantExternal)
			}
		})
	}
}

func TestArtifactRegistryRepositoryRef_Normalize_NotReadyAndNotFound(t *testing.T) {
	tests := []struct {
		name         string
		ref          v1beta1.ArtifactRegistryRepositoryRef
		objs         []client.Object
		wantNotReady bool
		wantNotFound bool
	}{
		{
			name:         "missing_repository_returns_not_found",
			ref:          v1beta1.ArtifactRegistryRepositoryRef{Name: "missing"},
			wantNotFound: true,
		},
		{
			name: "unreconciled_repository_with_empty_status_returns_not_ready",
			ref:  v1beta1.ArtifactRegistryRepositoryRef{Name: "upstream"},
			objs: []client.Object{
				newUnstructuredARRepo("upstream", "test-ns", "my-proj", "us-west1", nil),
			},
			wantNotReady: true,
		},
		{
			name: "ready_true_without_status_name_returns_not_ready",
			ref:  v1beta1.ArtifactRegistryRepositoryRef{Name: "upstream"},
			objs: []client.Object{
				newUnstructuredARRepo("upstream", "test-ns", "my-proj", "us-west1", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "True"},
					},
				}),
			},
			wantNotReady: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).WithObjects(tc.objs...).Build()
			r := tc.ref
			err := r.Normalize(t.Context(), reader, "test-ns")

			if err == nil {
				t.Fatalf("Normalize(%+v) = nil, want error (External=%q)", tc.ref, r.External)
			}
			if tc.wantNotReady && !k8s.IsReferenceNotReadyError(err) {
				t.Errorf("Normalize(%+v) error = %v, want ReferenceNotReadyError", tc.ref, err)
			}
			if tc.wantNotFound && !k8s.IsReferenceNotFoundError(err) {
				t.Errorf("Normalize(%+v) error = %v, want ReferenceNotFoundError", tc.ref, err)
			}
		})
	}
}

func newUnstructuredARRepo(name, namespace, projectID, location string, status map[string]any) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(v1beta1.ArtifactRegistryRepositoryGVK)
	u.SetName(name)
	u.SetNamespace(namespace)
	if projectID != "" {
		u.SetAnnotations(map[string]string{"cnrm.cloud.google.com/project-id": projectID})
	}
	if location != "" {
		u.Object["spec"] = map[string]any{"location": location}
	}
	if status != nil {
		u.Object["status"] = status
	}
	return u
}
