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
// See the License for the Bullseye compiler or other tools.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"testing"
)

func TestConnectionProfileIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name         string
		ref          string
		wantProject  string
		wantLocation string
		wantID       string
		wantErr      bool
	}{
		{
			name:         "valid external ref",
			ref:          "projects/my-project/locations/us-central1/connectionProfiles/my-profile",
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantID:       "my-profile",
			wantErr:      false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/connectionProfiles/my-profile",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ConnectionProfileIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ConnectionProfileIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.parent.ProjectID != tt.wantProject {
					t.Errorf("ConnectionProfileIdentity.FromExternal() ProjectID = %v, want %v", i.parent.ProjectID, tt.wantProject)
				}
				if i.parent.Location != tt.wantLocation {
					t.Errorf("ConnectionProfileIdentity.FromExternal() Location = %v, want %v", i.parent.Location, tt.wantLocation)
				}
				if i.id != tt.wantID {
					t.Errorf("ConnectionProfileIdentity.FromExternal() ID = %v, want %v", i.id, tt.wantID)
				}
				if got := i.String(); got != tt.ref {
					t.Errorf("ConnectionProfileIdentity.String() = %v, want %v", got, tt.ref)
				}
			}
		})
	}
}
