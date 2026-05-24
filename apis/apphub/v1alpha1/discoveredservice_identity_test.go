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

func TestParseAppHubDiscoveredServiceIdentity(t *testing.T) {
	tests := []struct {
		name      string
		external  string
		wantProj  string
		wantLoc   string
		wantResID string
		wantErr   bool
	}{
		{
			name:      "valid external ref - lowercase segment",
			external:  "projects/test-project/locations/us-central1/discoveredservices/test-service",
			wantProj:  "test-project",
			wantLoc:   "us-central1",
			wantResID: "test-service",
			wantErr:   false,
		},
		{
			name:      "valid external ref - canonical segment",
			external:  "projects/test-project/locations/us-central1/discoveredServices/test-service",
			wantProj:  "test-project",
			wantLoc:   "us-central1",
			wantResID: "test-service",
			wantErr:   false,
		},
		{
			name:     "invalid external ref - wrong service",
			external: "projects/test-project/locations/us-central1/invalid/test-service",
			wantErr:  true,
		},
		{
			name:     "invalid external ref - too short",
			external: "projects/test-project/locations/us-central1",
			wantErr:  true,
		},
		{
			name:     "invalid external ref - empty",
			external: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := ParseAppHubDiscoveredServiceIdentity(tt.external)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseAppHubDiscoveredServiceIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if id == nil {
					t.Fatalf("expected non-nil id")
				}
				if id.Project != tt.wantProj {
					t.Errorf("id.Project = %v, want %v", id.Project, tt.wantProj)
				}
				if id.Location != tt.wantLoc {
					t.Errorf("id.Location = %v, want %v", id.Location, tt.wantLoc)
				}
				if id.DiscoveredService != tt.wantResID {
					t.Errorf("id.DiscoveredService = %v, want %v", id.DiscoveredService, tt.wantResID)
				}
			}
		})
	}
}
