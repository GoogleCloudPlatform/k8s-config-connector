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

func TestParsePrivateConnectionExternal(t *testing.T) {
	tests := []struct {
		name         string
		external     string
		wantErr      bool
		wantProject  string
		wantLocation string
		wantID       string
	}{
		{
			name:         "valid external name",
			external:     "projects/my-project/locations/us-central1/privateConnections/my-private-connection",
			wantErr:      false,
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantID:       "my-private-connection",
		},
		{
			name:     "invalid external prefix",
			external: "project/my-project/locations/us-central1/privateConnections/my-private-connection",
			wantErr:  true,
		},
		{
			name:     "invalid format too short",
			external: "projects/my-project/locations/us-central1",
			wantErr:  true,
		},
		{
			name:     "invalid format too long",
			external: "projects/my-project/locations/us-central1/privateConnections/my-private-connection/extra",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, resourceID, err := ParsePrivateConnectionExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePrivateConnectionExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if p.ProjectID != tt.wantProject {
					t.Errorf("ProjectID = %v, want %v", p.ProjectID, tt.wantProject)
				}
				if p.Location != tt.wantLocation {
					t.Errorf("Location = %v, want %v", p.Location, tt.wantLocation)
				}
				if resourceID != tt.wantID {
					t.Errorf("ID = %v, want %v", resourceID, tt.wantID)
				}
			}
		})
	}
}
