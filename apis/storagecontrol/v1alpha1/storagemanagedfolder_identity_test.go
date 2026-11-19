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

package v1alpha1

import (
	"strings"
	"testing"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestNewManagedFolderIdentity_Success(t *testing.T) {
	strPtr := func(s string) *string { return &s }

	tests := []struct {
		name       string
		bucketRef  *storagev1beta1.StorageBucketRef
		objs       []client.Object
		resourceID *string
		wantString string
	}{
		{
			name:       "external_canonical_bucket_ref",
			bucketRef:  &storagev1beta1.StorageBucketRef{External: "projects/my-proj/buckets/ext-bucket"},
			resourceID: strPtr("folder-1"),
			wantString: "projects/my-proj/buckets/ext-bucket/managedfolders/folder-1",
		},
		{
			name:       "external_gs_uri_bucket_ref",
			bucketRef:  &storagev1beta1.StorageBucketRef{External: "gs://gs-bucket"},
			resourceID: strPtr("folder-1"),
			wantString: "projects/my-proj/buckets/gs-bucket/managedfolders/folder-1",
		},
		{
			name:       "external_bare_bucket_ref",
			bucketRef:  &storagev1beta1.StorageBucketRef{External: "bare-bucket"},
			resourceID: strPtr("folder-1"),
			wantString: "projects/my-proj/buckets/bare-bucket/managedfolders/folder-1",
		},
		{
			name: "krm_bucket_with_overridden_resource_id_in_other_namespace",
			bucketRef: &storagev1beta1.StorageBucketRef{
				Name:      "k8s-bucket-name",
				Namespace: "other-ns",
			},
			objs: []client.Object{
				newReadyBucketUnstructured("k8s-bucket-name", "other-ns", "real-gcs-bucket"),
			},
			resourceID: strPtr("folder-1"),
			wantString: "projects/my-proj/buckets/real-gcs-bucket/managedfolders/folder-1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).WithObjects(tc.objs...).Build()
			obj := &StorageManagedFolder{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-folder",
					Namespace: "default",
				},
				Spec: StorageManagedFolderSpec{
					StorageFolderParent: &StorageFolderParent{
						ProjectRef:       &refsv1beta1.ProjectRef{External: "my-proj"},
						StorageBucketRef: tc.bucketRef,
					},
					ResourceID: tc.resourceID,
				},
			}

			got, err := NewManagedFolderIdentity(t.Context(), reader, obj)

			if err != nil {
				t.Fatalf("NewManagedFolderIdentity(%+v) unexpected error: %v", tc.bucketRef, err)
			}
			if got.String() != tc.wantString {
				t.Errorf("NewManagedFolderIdentity(%+v).String() = %q, want %q", tc.bucketRef, got.String(), tc.wantString)
			}
		})
	}
}

func TestNewManagedFolderIdentity_Errors(t *testing.T) {
	unreadyBucket := &unstructured.Unstructured{}
	unreadyBucket.SetGroupVersionKind(storagev1beta1.StorageBucketGVK)
	unreadyBucket.SetName("unready-bucket")
	unreadyBucket.SetNamespace("default")

	reader := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).WithObjects(unreadyBucket).Build()

	t.Run("unready_bucket_returns_reference_not_ready", func(t *testing.T) {
		obj := &StorageManagedFolder{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "k8s-folder",
				Namespace: "default",
			},
			Spec: StorageManagedFolderSpec{
				StorageFolderParent: &StorageFolderParent{
					ProjectRef:       &refsv1beta1.ProjectRef{External: "my-proj"},
					StorageBucketRef: &storagev1beta1.StorageBucketRef{Name: "unready-bucket"},
				},
			},
		}

		got, err := NewManagedFolderIdentity(t.Context(), reader, obj)

		if err == nil {
			t.Fatalf("NewManagedFolderIdentity() = %+v, want ReferenceNotReadyError", got)
		}
		if !k8s.IsReferenceNotReadyError(err) {
			t.Errorf("NewManagedFolderIdentity() error = %v, want ReferenceNotReadyError", err)
		}
	})

	t.Run("nil_storage_bucket_ref_returns_error", func(t *testing.T) {
		obj := &StorageManagedFolder{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "k8s-folder",
				Namespace: "default",
			},
			Spec: StorageManagedFolderSpec{
				StorageFolderParent: &StorageFolderParent{
					ProjectRef: &refsv1beta1.ProjectRef{External: "my-proj"},
				},
			},
		}

		got, err := NewManagedFolderIdentity(t.Context(), reader, obj)

		if err == nil {
			t.Fatalf("NewManagedFolderIdentity() = %+v, want error", got)
		}
	})

	t.Run("nil_storage_folder_parent_returns_error", func(t *testing.T) {
		obj := &StorageManagedFolder{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "k8s-folder",
				Namespace: "default",
			},
		}

		got, err := NewManagedFolderIdentity(t.Context(), reader, obj)

		if err == nil {
			t.Fatalf("NewManagedFolderIdentity() = %+v, want error", got)
		}
		if !strings.Contains(err.Error(), "spec.projectRef and spec.storagebucketRef are required") {
			t.Errorf("NewManagedFolderIdentity() error = %v, want it to mention the required refs", err)
		}
	})
}

func newReadyBucketUnstructured(name, namespace, gcsBucket string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(storagev1beta1.StorageBucketGVK)
	u.SetName(name)
	u.SetNamespace(namespace)
	u.SetAnnotations(map[string]string{"cnrm.cloud.google.com/project-id": "my-proj"})
	u.Object["status"] = map[string]any{
		"conditions": []any{
			map[string]any{"type": "Ready", "status": "True"},
		},
		"url": "gs://" + gcsBucket,
	}
	return u
}
