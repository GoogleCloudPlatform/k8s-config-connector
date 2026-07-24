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

package v1beta1

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestStorageBucketObjectIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *StorageBucketObjectIdentity
		wantErr bool
	}{
		{
			name: "valid full URL",
			ref:  "//storage.googleapis.com/projects/my-project/buckets/my-bucket/objects/my-object",
			want: &StorageBucketObjectIdentity{
				Project: "my-project",
				Bucket:  "my-bucket",
				Object:  "my-object",
			},
		},
		{
			name: "valid relative path",
			ref:  "projects/my-project/buckets/my-bucket/objects/my-object",
			want: &StorageBucketObjectIdentity{
				Project: "my-project",
				Bucket:  "my-bucket",
				Object:  "my-object",
			},
		},
		{
			name: "valid relative path with slashes in object name",
			ref:  "projects/my-project/buckets/my-bucket/objects/folder/subfolder/my-object.png",
			want: &StorageBucketObjectIdentity{
				Project: "my-project",
				Bucket:  "my-bucket",
				Object:  "folder/subfolder/my-object.png",
			},
		},
		{
			name:    "invalid format - missing project",
			ref:     "buckets/my-bucket/objects/my-object",
			wantErr: true,
		},
		{
			name:    "invalid format - missing object segment",
			ref:     "projects/my-project/buckets/my-bucket/my-object",
			wantErr: true,
		},
		{
			name:    "invalid format - wrong collection",
			ref:     "projects/my-project/buckets/my-bucket/foos/my-object",
			wantErr: true,
		},
		{
			name:    "empty object",
			ref:     "projects/my-project/buckets/my-bucket/objects/",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StorageBucketObjectIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestStorageBucketObjectIdentity_Methods(t *testing.T) {
	i := &StorageBucketObjectIdentity{
		Project: "my-project",
		Bucket:  "my-bucket",
		Object:  "my-object",
	}

	if got := i.String(); got != "projects/my-project/buckets/my-bucket/objects/my-object" {
		t.Errorf("String() = %v, want %v", got, "projects/my-project/buckets/my-bucket/objects/my-object")
	}

	if got := i.ParentString(); got != "projects/my-project/buckets/my-bucket" {
		t.Errorf("ParentString() = %v, want %v", got, "projects/my-project/buckets/my-bucket")
	}
}

func TestStorageBucketObjectGetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)

	obj := &StorageBucketObject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-object",
			Namespace: "test-namespace",
		},
		Spec: StorageBucketObjectSpec{
			BucketRef: &StorageBucketRef{
				External: "projects/test-project/buckets/test-bucket",
			},
		},
	}

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	gotIdentity, err := obj.GetIdentity(ctx, fakeClient)
	if err != nil {
		t.Fatalf("GetIdentity() returned unexpected error: %v", err)
	}

	wantIdentity := &StorageBucketObjectIdentity{
		Project: "test-project",
		Bucket:  "test-bucket",
		Object:  "test-object",
	}

	if diff := cmp.Diff(wantIdentity, gotIdentity); diff != "" {
		t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
	}
}
