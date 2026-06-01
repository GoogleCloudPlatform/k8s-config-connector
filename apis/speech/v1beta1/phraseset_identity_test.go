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
	"testing"
)

func TestParsePhraseSetExternal(t *testing.T) {
	tests := []struct {
		name      string
		external  string
		wantErr   bool
		wantProj  string
		wantLoc   string
		wantResID string
	}{
		{
			name:      "valid external ID",
			external:  "projects/test-project/locations/us-central1/phraseSets/test-phraseset",
			wantProj:  "test-project",
			wantLoc:   "us-central1",
			wantResID: "test-phraseset",
			wantErr:   false,
		},
		{
			name:     "invalid external ID - wrong segments",
			external: "projects/test-project/locations/us-central1/invalid/test-phraseset",
			wantErr:  true,
		},
		{
			name:     "invalid external ID - too short",
			external: "projects/test-project/locations/us-central1",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parent, resourceID, err := ParsePhraseSetExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePhraseSetExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if parent.ProjectID != tt.wantProj {
					t.Errorf("ProjectID = %v, want %v", parent.ProjectID, tt.wantProj)
				}
				if parent.Location != tt.wantLoc {
					t.Errorf("Location = %v, want %v", parent.Location, tt.wantLoc)
				}
				if resourceID != tt.wantResID {
					t.Errorf("resourceID = %v, want %v", resourceID, tt.wantResID)
				}
			}
		})
	}
}
