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

func TestParseFolderExternal(t *testing.T) {
	tests := []struct {
		name           string
		external       string
		wantErr        bool
		wantProjectID  string
		wantBucketName string
		wantResourceID string
	}{
		{
			name:           "valid external ref",
			external:       "projects/my-project/buckets/my-bucket/folders/my-folder",
			wantProjectID:  "my-project",
			wantBucketName: "my-bucket",
			wantResourceID: "my-folder",
		},
		{
			name:     "invalid external - wrong segments count",
			external: "projects/my-project/buckets/my-bucket/folders",
			wantErr:  true,
		},
		{
			name:     "invalid external - wrong path",
			external: "projects/my-project/buckets/my-bucket/invalid/my-folder",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parent, resourceID, err := ParseFolderExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFolderExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if parent.ProjectID != tt.wantProjectID {
					t.Errorf("ProjectID = %v, want %v", parent.ProjectID, tt.wantProjectID)
				}
				if parent.BucketName != tt.wantBucketName {
					t.Errorf("BucketName = %v, want %v", parent.BucketName, tt.wantBucketName)
				}
				if resourceID != tt.wantResourceID {
					t.Errorf("ResourceID = %v, want %v", resourceID, tt.wantResourceID)
				}
			}
		})
	}
}
