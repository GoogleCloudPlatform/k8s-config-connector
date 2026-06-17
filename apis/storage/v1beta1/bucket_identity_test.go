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

func TestStorageBucketIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *StorageBucketIdentity
		wantErr bool
	}{
		{
			name: "valid full URL",
			ref:  "//storage.googleapis.com/projects/my-project/buckets/my-bucket",
			want: &StorageBucketIdentity{
				Project: "my-project",
				Bucket:  "my-bucket",
			},
		},
		{
			name: "valid relative path",
			ref:  "projects/my-project/buckets/my-bucket",
			want: &StorageBucketIdentity{
				Project: "my-project",
				Bucket:  "my-bucket",
			},
		},
		{
			name: "gs scheme - valid",
			ref:  "gs://my-bucket",
			want: &StorageBucketIdentity{
				Project: "",
				Bucket:  "my-bucket",
			},
		},
		{
			name:    "invalid format - missing project",
			ref:     "buckets/my-bucket",
			wantErr: true,
		},
		{
			name:    "invalid format - extra segments",
			ref:     "projects/my-project/buckets/my-bucket/extra",
			wantErr: true,
		},
		{
			name:    "invalid format - wrong collection",
			ref:     "projects/my-project/foos/my-bucket",
			wantErr: true,
		},
		{
			name:    "empty bucket",
			ref:     "projects/my-project/buckets/",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StorageBucketIdentity{}
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

func TestStorageBucketIdentity_Methods(t *testing.T) {
	i := &StorageBucketIdentity{
		Project: "my-project",
		Bucket:  "my-bucket",
	}

	if got := i.String(); got != "projects/my-project/buckets/my-bucket" {
		t.Errorf("String() = %v, want %v", got, "projects/my-project/buckets/my-bucket")
	}

	if got := i.BucketName(); got != "my-bucket" {
		t.Errorf("BucketName() = %v, want %v", got, "my-bucket")
	}
}

func TestStorageBucketGetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)

	obj := &StorageBucket{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-bucket",
			Namespace: "test-namespace",
			Annotations: map[string]string{
				"cnrm.cloud.google.com/project-id": "test-project",
			},
		},
	}

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	gotIdentity, err := obj.GetIdentity(ctx, fakeClient)
	if err != nil {
		t.Fatalf("GetIdentity() returned unexpected error: %v", err)
	}

	wantIdentity := &StorageBucketIdentity{
		Project: "test-project",
		Bucket:  "test-bucket",
	}

	if diff := cmp.Diff(wantIdentity, gotIdentity); diff != "" {
		t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
	}
}
