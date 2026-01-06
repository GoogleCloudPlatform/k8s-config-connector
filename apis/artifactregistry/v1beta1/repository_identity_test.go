// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"testing"
)

func TestParseArtifactRegistryRepositoryIdentity(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		project   string
		location  string
		resource  string
		wantError bool
	}{
		{
			name:     "valid identity",
			id:       "projects/my-project/locations/us-central1/repositories/my-repo",
			project:  "my-project",
			location: "us-central1",
			resource: "my-repo",
		},
		{
			name:      "invalid identity - too short",
			id:        "projects/my-project/locations/us-central1/repositories",
			wantError: true,
		},
		{
			name:      "invalid identity - wrong separator",
			id:        "projects/my-project/locations/us-central1/repo/my-repo",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &ArtifactRegistryRepository{}
			project, location, resourceID, err := obj.ParseIdentity(tt.id)
			if (err != nil) != tt.wantError {
				t.Errorf("ParseIdentity() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if !tt.wantError {
				if project != tt.project {
					t.Errorf("ParseIdentity() project = %v, want %v", project, tt.project)
				}
				if location != tt.location {
					t.Errorf("ParseIdentity() location = %v, want %v", location, tt.location)
				}
				if resourceID != tt.resource {
					t.Errorf("ParseIdentity() resourceID = %v, want %v", resourceID, tt.resource)
				}
			}
		})
	}
}
