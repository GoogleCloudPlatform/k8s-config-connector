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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestStorageBucketRef_Normalize_Success(t *testing.T) {
	tests := []struct {
		name         string
		ref          v1beta1.StorageBucketRef
		objs         []client.Object
		wantExternal string
	}{
		// StorageBucketRef intentionally preserves the user's external form; consumers parse it.
		{
			name: "accepts_bare_bucket_external",
			ref: v1beta1.StorageBucketRef{
				External: "my-bucket",
			},
			wantExternal: "my-bucket",
		},
		{
			name: "accepts_self_link_external",
			ref: v1beta1.StorageBucketRef{
				External: "https://www.googleapis.com/storage/v1/b/my-bucket",
			},
			wantExternal: "https://www.googleapis.com/storage/v1/b/my-bucket",
		},
		{
			name: "accepts_gs_uri_with_trailing_slash",
			ref: v1beta1.StorageBucketRef{
				External: "gs://my-bucket/",
			},
			wantExternal: "gs://my-bucket/",
		},
		{
			name: "resolves_ready_bucket_from_status_url_and_project_annotation",
			ref: v1beta1.StorageBucketRef{
				Name: "k8s-bucket",
			},
			objs: []client.Object{
				newUnstructuredStorageBucket("k8s-bucket", "test-ns", "my-project", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "True"},
					},
					"url":      "gs://actual-gcs-bucket",
					"selfLink": "https://www.googleapis.com/storage/v1/b/actual-gcs-bucket",
				}),
			},
			wantExternal: "projects/my-project/buckets/actual-gcs-bucket",
		},
		{
			name: "resolves_ready_bucket_from_status_self_link_without_project_annotation",
			ref: v1beta1.StorageBucketRef{
				Name: "k8s-bucket",
			},
			objs: []client.Object{
				newUnstructuredStorageBucket("k8s-bucket", "test-ns", "", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "True"},
					},
					"selfLink": "https://www.googleapis.com/storage/v1/b/actual-gcs-bucket",
				}),
			},
			wantExternal: "projects/_/buckets/actual-gcs-bucket",
		},
		{
			name: "resolves_bucket_with_ready_false_delete_failed_from_status_url",
			ref: v1beta1.StorageBucketRef{
				Name: "k8s-bucket",
			},
			objs: []client.Object{
				newUnstructuredStorageBucket("k8s-bucket", "test-ns", "my-project", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "False", "reason": "DeleteFailed"},
					},
					"url": "gs://actual-gcs-bucket",
				}),
			},
			wantExternal: "projects/my-project/buckets/actual-gcs-bucket",
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

func TestStorageBucketRef_Normalize_Errors(t *testing.T) {
	tests := []struct {
		name         string
		ref          v1beta1.StorageBucketRef
		objs         []client.Object
		wantNotReady bool
		wantNotFound bool
	}{
		{
			name:         "missing_bucket_cr_returns_not_found",
			ref:          v1beta1.StorageBucketRef{Name: "missing"},
			wantNotFound: true,
		},
		{
			name: "uncreated_bucket_with_empty_status_returns_not_ready",
			ref:  v1beta1.StorageBucketRef{Name: "uncreated"},
			objs: []client.Object{
				newUnstructuredStorageBucket("uncreated", "test-ns", "my-project", nil),
			},
			wantNotReady: true,
		},
		{
			name: "failed_foreign_bucket_with_ready_false_returns_not_ready",
			ref:  v1beta1.StorageBucketRef{Name: "foreign-bucket"},
			objs: []client.Object{
				newUnstructuredStorageBucket("foreign-bucket", "test-ns", "my-project", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "False", "reason": "UpdateFailed"},
					},
				}),
			},
			wantNotReady: true,
		},
		{
			name: "ready_true_without_status_url_or_self_link_returns_not_ready",
			ref:  v1beta1.StorageBucketRef{Name: "empty-urls"},
			objs: []client.Object{
				newUnstructuredStorageBucket("empty-urls", "test-ns", "my-project", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "True"},
					},
				}),
			},
			wantNotReady: true,
		},
		{
			name: "unparseable_status_url_and_self_link_returns_not_ready",
			ref:  v1beta1.StorageBucketRef{Name: "bad-status"},
			objs: []client.Object{
				newUnstructuredStorageBucket("bad-status", "test-ns", "my-project", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "True"},
					},
					"url":      "gs://",
					"selfLink": "https://www.googleapis.com/storage/v1/b/",
				}),
			},
			wantNotReady: true,
		},
		{
			name: "invalid_project_annotation_returns_not_ready",
			ref:  v1beta1.StorageBucketRef{Name: "bad-project"},
			objs: []client.Object{
				newUnstructuredStorageBucket("bad-project", "test-ns", "my-project/buckets/other", map[string]any{
					"conditions": []any{
						map[string]any{"type": "Ready", "status": "True"},
					},
					"url": "gs://actual-gcs-bucket",
				}),
			},
			wantNotReady: true,
		},
		{
			name: "rejects_empty_gs_uri",
			ref:  v1beta1.StorageBucketRef{External: "gs://"},
		},
		{
			name: "rejects_invalid_gs_uri_with_whitespace",
			ref:  v1beta1.StorageBucketRef{External: "gs://my bucket"},
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

func newUnstructuredStorageBucket(name, namespace, projectID string, status map[string]any) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(v1beta1.StorageBucketGVK)
	u.SetName(name)
	u.SetNamespace(namespace)
	if projectID != "" {
		u.SetAnnotations(map[string]string{"cnrm.cloud.google.com/project-id": projectID})
	}
	u.Object["spec"] = map[string]any{"resourceID": name + "-from-spec"}
	if status != nil {
		u.Object["status"] = status
	}
	return u
}
