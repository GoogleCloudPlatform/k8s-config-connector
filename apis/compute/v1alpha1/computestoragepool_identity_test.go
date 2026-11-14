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
	"context"
	"testing"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeStoragePoolIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeStoragePoolIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/zones/us-central1-a/storagePools/my-storagepool",
			want: &ComputeStoragePoolIdentity{
				Project:     "my-project",
				Zone:        "us-central1-a",
				StoragePool: "my-storagepool",
			},
		},
		{
			name: "reference with host prefix",
			ref:  "//compute.googleapis.com/projects/my-project/zones/us-central1-a/storagePools/my-storagepool",
			want: &ComputeStoragePoolIdentity{
				Project:     "my-project",
				Zone:        "us-central1-a",
				StoragePool: "my-storagepool",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "missing project",
			ref:     "zones/us-central1-a/storagePools/my-storagepool",
			wantErr: true,
		},
		{
			name:    "missing zone",
			ref:     "projects/my-project/storagePools/my-storagepool",
			wantErr: true,
		},
		{
			name:    "missing storagePool",
			ref:     "projects/my-project/zones/us-central1-a/storagePools/",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeStoragePoolIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestComputeStoragePoolIdentity_ParentString(t *testing.T) {
	id := &ComputeStoragePoolIdentity{
		Project:     "my-project",
		Zone:        "us-central1-a",
		StoragePool: "my-storagepool",
	}
	want := "projects/my-project/zones/us-central1-a"
	got := id.ParentString()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ParentString() mismatch (-want +got):\n%s", diff)
	}
}

func TestComputeStoragePoolIdentity_String(t *testing.T) {
	id := &ComputeStoragePoolIdentity{
		Project:     "my-project",
		Zone:        "us-central1-a",
		StoragePool: "my-storagepool",
	}
	want := "projects/my-project/zones/us-central1-a/storagePools/my-storagepool"
	got := id.String()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("String() mismatch (-want +got):\n%s", diff)
	}
}

func TestComputeStoragePoolIdentity_GetIdentity(t *testing.T) {
	s := runtime.NewScheme()
	s.AddKnownTypes(GroupVersion, &ComputeStoragePool{})
	cl := fake.NewClientBuilder().WithScheme(s).Build()

	tests := []struct {
		name    string
		obj     *ComputeStoragePool
		want    *ComputeStoragePoolIdentity
		wantErr bool
	}{
		{
			name: "from spec with resourceID",
			obj: &ComputeStoragePool{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-pool",
					Namespace: "default",
				},
				Spec: ComputeStoragePoolSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
					Location:   "us-central1-a",
					ResourceID: ptrTo("custom-pool"),
				},
			},
			want: &ComputeStoragePoolIdentity{
				Project:     "my-project",
				Zone:        "us-central1-a",
				StoragePool: "custom-pool",
			},
		},
		{
			name: "from spec defaulting to metadata.name",
			obj: &ComputeStoragePool{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-pool",
					Namespace: "default",
				},
				Spec: ComputeStoragePoolSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
					Location: "us-central1-a",
				},
			},
			want: &ComputeStoragePoolIdentity{
				Project:     "my-project",
				Zone:        "us-central1-a",
				StoragePool: "k8s-pool",
			},
		},
		{
			name: "matching status.externalRef",
			obj: &ComputeStoragePool{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-pool",
					Namespace: "default",
				},
				Spec: ComputeStoragePoolSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
					Location: "us-central1-a",
				},
				Status: ComputeStoragePoolStatus{
					ExternalRef: ptrTo("projects/my-project/zones/us-central1-a/storagePools/k8s-pool"),
				},
			},
			want: &ComputeStoragePoolIdentity{
				Project:     "my-project",
				Zone:        "us-central1-a",
				StoragePool: "k8s-pool",
			},
		},
		{
			name: "conflicting status.externalRef",
			obj: &ComputeStoragePool{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-pool",
					Namespace: "default",
				},
				Spec: ComputeStoragePoolSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
					Location: "us-central1-a",
				},
				Status: ComputeStoragePoolStatus{
					ExternalRef: ptrTo("projects/other-project/zones/us-central1-a/storagePools/k8s-pool"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(context.Background(), cl)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				id, ok := got.(*ComputeStoragePoolIdentity)
				if !ok {
					t.Fatalf("GetIdentity() returned %T, want *ComputeStoragePoolIdentity", got)
				}
				if diff := cmp.Diff(tt.want, id); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
