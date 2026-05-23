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
	"testing"
)

func TestParseManagedFolderExternal(t *testing.T) {
	tests := []struct {
		name       string
		external   string
		wantErr    bool
		wantParent *ManagedFolderParent
		wantID     string
	}{
		{
			name:     "valid external ID",
			external: "projects/my-project/buckets/my-bucket/managedfolders/my-folder",
			wantParent: &ManagedFolderParent{
				ProjectID:  "my-project",
				BucketName: "my-bucket",
			},
			wantID: "my-folder",
		},
		{
			name:     "invalid external format",
			external: "projects/my-project/buckets/my-bucket/folders/my-folder",
			wantErr:  true,
		},
		{
			name:     "too few tokens",
			external: "projects/my-project/buckets/my-bucket",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParent, gotID, err := ParseManagedFolderExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseManagedFolderExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if gotID != tt.wantID {
					t.Errorf("gotID = %q, want %q", gotID, tt.wantID)
				}
				if gotParent.ProjectID != tt.wantParent.ProjectID {
					t.Errorf("gotParent.ProjectID = %q, want %q", gotParent.ProjectID, tt.wantParent.ProjectID)
				}
				if gotParent.BucketName != tt.wantParent.BucketName {
					t.Errorf("gotParent.BucketName = %q, want %q", gotParent.BucketName, tt.wantParent.BucketName)
				}
			}
		})
	}
}
