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
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestNewManagedFolderIdentity_DefensiveChecks(t *testing.T) {
	ctx := context.Background()
	fakeClient := fake.NewClientBuilder().Build()

	tests := []struct {
		name    string
		obj     *StorageManagedFolder
		wantErr string
	}{
		{
			name: "StorageFolderParent is nil",
			obj: &StorageManagedFolder{
				Spec: StorageManagedFolderSpec{
					StorageFolderParent: nil,
				},
			},
			wantErr: "spec.projectRef and spec.storagebucketRef are required",
		},
		{
			name: "ProjectRef is nil",
			obj: &StorageManagedFolder{
				Spec: StorageManagedFolderSpec{
					StorageFolderParent: &StorageFolderParent{
						ProjectRef:       nil,
						StorageBucketRef: &storagev1beta1.StorageBucketRef{},
					},
				},
			},
			wantErr: "spec.projectRef and spec.storagebucketRef are required",
		},
		{
			name: "StorageBucketRef is nil",
			obj: &StorageManagedFolder{
				Spec: StorageManagedFolderSpec{
					StorageFolderParent: &StorageFolderParent{
						ProjectRef:       &v1beta1.ProjectRef{},
						StorageBucketRef: nil,
					},
				},
			},
			wantErr: "spec.projectRef and spec.storagebucketRef are required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewManagedFolderIdentity(ctx, fakeClient, tt.obj)
			if err == nil {
				t.Fatalf("expected error containing %q, but got nil", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("expected error containing %q, but got %v", tt.wantErr, err)
			}
		})
	}
}

func TestNewManagedFolderIdentityFromExternal(t *testing.T) {
	tests := []struct {
		name       string
		external   string
		wantProj   string
		wantBucket string
		wantID     string
		wantErr    bool
	}{
		{
			name:       "standard format",
			external:   "projects/my-project/buckets/my-bucket/managedfolders/my-folder",
			wantProj:   "my-project",
			wantBucket: "my-bucket",
			wantID:     "my-folder",
			wantErr:    false,
		},
		{
			name:       "capital F format",
			external:   "projects/my-project/buckets/my-bucket/managedFolders/my-folder",
			wantProj:   "my-project",
			wantBucket: "my-bucket",
			wantID:     "my-folder",
			wantErr:    false,
		},
		{
			name:       "URL with storage.googleapis.com prefix",
			external:   "storage.googleapis.com/projects/my-project/buckets/my-bucket/managedfolders/my-folder",
			wantProj:   "my-project",
			wantBucket: "my-bucket",
			wantID:     "my-folder",
			wantErr:    false,
		},
		{
			name:       "URL with double slash storage.googleapis.com prefix",
			external:   "//storage.googleapis.com/projects/my-project/buckets/my-bucket/managedFolders/my-folder",
			wantProj:   "my-project",
			wantBucket: "my-bucket",
			wantID:     "my-folder",
			wantErr:    false,
		},
		{
			name:       "trailing slash",
			external:   "projects/my-project/buckets/my-bucket/managedFolders/my-folder/",
			wantProj:   "my-project",
			wantBucket: "my-bucket",
			wantID:     "my-folder",
			wantErr:    false,
		},
		{
			name:     "invalid format",
			external: "projects/my-project/buckets/my-bucket/something-else/my-folder",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := NewManagedFolderIdentityFromExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewManagedFolderIdentityFromExternal(%q) returned err: %v, wantErr: %v", tt.external, err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if id.Parent().ProjectID != tt.wantProj {
				t.Errorf("ProjectID = %q, want %q", id.Parent().ProjectID, tt.wantProj)
			}
			if id.Parent().BucketName != tt.wantBucket {
				t.Errorf("BucketName = %q, want %q", id.Parent().BucketName, tt.wantBucket)
			}
			if id.ID() != tt.wantID {
				t.Errorf("ID = %q, want %q", id.ID(), tt.wantID)
			}
		})
	}
}
