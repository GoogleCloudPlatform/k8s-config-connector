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

func TestParseNodeGroupExternal(t *testing.T) {
	tests := []struct {
		name        string
		external    string
		wantProject string
		wantRegion  string
		wantCluster string
		wantID      string
		wantErr     bool
	}{
		{
			name:        "valid external ref",
			external:    "projects/my-project/regions/us-central1/clusters/my-cluster/nodeGroups/my-nodegroup",
			wantProject: "my-project",
			wantRegion:  "us-central1",
			wantCluster: "my-cluster",
			wantID:      "my-nodegroup",
			wantErr:     false,
		},
		{
			name:     "invalid external ref - wrong service",
			external: "projects/my-project/locations/us-central1/sessions/my-session",
			wantErr:  true,
		},
		{
			name:     "invalid external ref - too short",
			external: "projects/my-project/regions/us-central1/clusters/my-cluster",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parent, resourceID, err := ParseNodeGroupExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNodeGroupExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if parent.ProjectID != tt.wantProject {
					t.Errorf("ParseNodeGroupExternal() ProjectID = %v, want %v", parent.ProjectID, tt.wantProject)
				}
				if parent.Region != tt.wantRegion {
					t.Errorf("ParseNodeGroupExternal() Region = %v, want %v", parent.Region, tt.wantRegion)
				}
				if parent.Cluster != tt.wantCluster {
					t.Errorf("ParseNodeGroupExternal() Cluster = %v, want %v", parent.Cluster, tt.wantCluster)
				}
				if resourceID != tt.wantID {
					t.Errorf("ParseNodeGroupExternal() resourceID = %v, want %v", resourceID, tt.wantID)
				}
				identity := &NodeGroupIdentity{
					parent: parent,
					id:     resourceID,
				}
				if got := identity.String(); got != tt.external {
					t.Errorf("NodeGroupIdentity.String() = %v, want %v", got, tt.external)
				}
			}
		})
	}
}
