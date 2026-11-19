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

package v1beta1_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestContainerClusterRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid_regional_reference",
			ref:     "projects/my-project/locations/us-central1/clusters/my-cluster",
			wantErr: false,
		},
		{
			name:    "valid_zonal_reference",
			ref:     "projects/my-project/zones/us-central1-a/clusters/my-cluster",
			wantErr: false,
		},
		{
			name:    "invalid_prefix",
			ref:     "invalid/my-project/locations/us-central1/clusters/my-cluster",
			wantErr: true,
		},
		{
			name:    "missing_location_or_zone",
			ref:     "projects/my-project/clusters/my-cluster",
			wantErr: true,
		},
		{
			name:    "missing_cluster_name",
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
			r := &v1beta1.ContainerClusterRef{}
			err := r.ValidateExternal(tc.ref)

			if (err != nil) != tc.wantErr {
				t.Errorf("ValidateExternal(%q) error = %v, wantErr %v", tc.ref, err, tc.wantErr)
			}
		})
	}
}

func TestContainerClusterRef_Normalize_Success(t *testing.T) {
	tests := []struct {
		name         string
		ref          v1beta1.ContainerClusterRef
		objs         []client.Object
		wantExternal string
	}{
		{
			name: "canonicalizes_full_zonal_https_external",
			ref: v1beta1.ContainerClusterRef{
				External: "https://container.googleapis.com/v1/projects/proj-a/zones/us-central1-a/clusters/cluster-a",
			},
			wantExternal: "projects/proj-a/locations/us-central1-a/clusters/cluster-a",
		},
		{
			name: "canonicalizes_host_prefixed_external",
			ref: v1beta1.ContainerClusterRef{
				External: "//container.googleapis.com/projects/proj-a/locations/us-central1/clusters/cluster-a",
			},
			wantExternal: "projects/proj-a/locations/us-central1/clusters/cluster-a",
		},
		{
			name: "preserves_zonal_relative_external",
			ref: v1beta1.ContainerClusterRef{
				External: "projects/proj-a/zones/us-central1-a/clusters/cluster-a",
			},
			wantExternal: "projects/proj-a/zones/us-central1-a/clusters/cluster-a",
		},
		{
			name: "preserves_regional_relative_external",
			ref: v1beta1.ContainerClusterRef{
				External: "projects/proj-a/locations/us-central1/clusters/cluster-a",
			},
			wantExternal: "projects/proj-a/locations/us-central1/clusters/cluster-a",
		},
		{
			name: "resolves_ready_cluster_from_status_self_link_without_spec_or_project_annotation",
			ref: v1beta1.ContainerClusterRef{
				Name: "cluster-cr",
			},
			objs: []client.Object{
				newUnstructuredContainerCluster("cluster-cr", "test-ns", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "True"},
					},
					"selfLink": "https://container.googleapis.com/v1beta1/projects/proj-status/zones/us-west1-b/clusters/real-cluster",
				}),
			},
			wantExternal: "projects/proj-status/locations/us-west1-b/clusters/real-cluster",
		},
		{
			name: "resolves_cluster_with_ready_false_updating_from_status_self_link",
			ref: v1beta1.ContainerClusterRef{
				Name: "updating",
			},
			objs: []client.Object{
				newUnstructuredContainerClusterWithSpec("updating", "test-ns", "proj-a", "us-central1", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "False", "reason": "Updating"},
					},
					"selfLink": "https://container.googleapis.com/v1/projects/proj-a/locations/us-central1/clusters/updating",
				}),
			},
			wantExternal: "projects/proj-a/locations/us-central1/clusters/updating",
		},
		{
			name: "resolves_cluster_with_ready_false_delete_failed_from_status_self_link",
			ref: v1beta1.ContainerClusterRef{
				Name: "delete-failed",
			},
			objs: []client.Object{
				newUnstructuredContainerClusterWithSpec("delete-failed", "test-ns", "proj-a", "us-central1", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "False", "reason": "DeleteFailed"},
					},
					"selfLink": "https://container.googleapis.com/v1/projects/proj-a/zones/us-central1-a/clusters/delete-failed",
				}),
			},
			wantExternal: "projects/proj-a/locations/us-central1-a/clusters/delete-failed",
		},
		{
			name: "resolves_from_status_external_ref_and_canonicalizes",
			ref: v1beta1.ContainerClusterRef{
				Name: "cluster-cr",
			},
			objs: []client.Object{
				newUnstructuredContainerCluster("cluster-cr", "test-ns", map[string]any{
					"externalRef": "//container.googleapis.com/projects/proj-ext/zones/us-east1-b/clusters/ext-cluster",
				}),
			},
			wantExternal: "projects/proj-ext/locations/us-east1-b/clusters/ext-cluster",
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

func TestContainerClusterRef_Normalize_NotReadyAndNotFound(t *testing.T) {
	tests := []struct {
		name         string
		ref          v1beta1.ContainerClusterRef
		objs         []client.Object
		wantNotReady bool
		wantNotFound bool
	}{
		{
			name: "missing_cluster_returns_not_found",
			ref:  v1beta1.ContainerClusterRef{Name: "missing"},
			objs: nil,

			wantNotFound: true,
		},
		{
			name: "cluster_with_spec_only_and_no_status_returns_not_ready",
			ref:  v1beta1.ContainerClusterRef{Name: "unreconciled"},
			objs: []client.Object{
				newUnstructuredContainerClusterWithSpec("unreconciled", "test-ns", "proj-a", "us-central1", nil),
			},
			wantNotReady: true,
		},
		{
			name: "cluster_with_ready_false_update_failed_returns_not_ready",
			ref:  v1beta1.ContainerClusterRef{Name: "failed"},
			objs: []client.Object{
				newUnstructuredContainerClusterWithSpec("failed", "test-ns", "proj-a", "us-central1", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "False", "reason": "UpdateFailed"},
					},
				}),
			},
			wantNotReady: true,
		},
		{
			name: "cluster_with_ready_true_but_missing_self_link_returns_not_ready",
			ref:  v1beta1.ContainerClusterRef{Name: "no-selflink"},
			objs: []client.Object{
				newUnstructuredContainerClusterWithSpec("no-selflink", "test-ns", "proj-a", "us-central1", map[string]any{
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

func newUnstructuredContainerCluster(name, namespace string, status map[string]any) *unstructured.Unstructured {
	return newUnstructuredContainerClusterWithSpec(name, namespace, "", "", status)
}

func newUnstructuredContainerClusterWithSpec(name, namespace, projectID, location string, status map[string]any) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(v1beta1.ContainerClusterGVK)
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
