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
			id := &NodeGroupIdentity{}
			err := id.FromExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("NodeGroupIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if id.Project != tt.wantProject {
					t.Errorf("Project = %v, want %v", id.Project, tt.wantProject)
				}
				if id.Region != tt.wantRegion {
					t.Errorf("Region = %v, want %v", id.Region, tt.wantRegion)
				}
				if id.Cluster != tt.wantCluster {
					t.Errorf("Cluster = %v, want %v", id.Cluster, tt.wantCluster)
				}
				if id.NodeGroup != tt.wantID {
					t.Errorf("NodeGroup = %v, want %v", id.NodeGroup, tt.wantID)
				}
				if got := id.String(); got != tt.external {
					t.Errorf("NodeGroupIdentity.String() = %v, want %v", got, tt.external)
				}
			}
		})
	}
}
