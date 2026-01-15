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
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
)

func TestStorageBucketIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name        string
		ref         string
		wantProject string
		wantBucket  string
		wantErr     bool
	}{
		{
			name:        "valid full URL",
			ref:         "//storage.googleapis.com/projects/my-project/buckets/my-bucket",
			wantProject: "my-project",
			wantBucket:  "my-bucket",
		},
		{
			name:        "valid relative path",
			ref:         "projects/my-project/buckets/my-bucket",
			wantProject: "my-project",
			wantBucket:  "my-bucket",
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
				if i.Project != tt.wantProject {
					t.Errorf("FromExternal() Project got = %v, want %v", i.Project, tt.wantProject)
				}
				if i.Bucket != tt.wantBucket {
					t.Errorf("FromExternal() Bucket got = %v, want %v", i.Bucket, tt.wantBucket)
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

	wantParent := &parent.ProjectParent{ProjectID: "my-project"}
	if got := i.Parent(); !reflect.DeepEqual(got, wantParent) {
		t.Errorf("Parent() = %v, want %v", got, wantParent)
	}
}
