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

func TestParseMemoryLayerExternal(t *testing.T) {
	tests := []struct {
		name       string
		external   string
		wantErr    bool
		projectID  string
		instanceID string
		clusterID  string
	}{
		{
			name:       "valid external name",
			external:   "projects/my-project/instances/my-instance/clusters/my-cluster/memoryLayer",
			wantErr:    false,
			projectID:  "my-project",
			instanceID: "my-instance",
			clusterID:  "my-cluster",
		},
		{
			name:       "valid external name with leading slash",
			external:   "/projects/my-project/instances/my-instance/clusters/my-cluster/memoryLayer",
			wantErr:    false,
			projectID:  "my-project",
			instanceID: "my-instance",
			clusterID:  "my-cluster",
		},
		{
			name:     "invalid external prefix",
			external: "project/my-project/instances/my-instance/clusters/my-cluster/memoryLayer",
			wantErr:  true,
		},
		{
			name:     "missing memoryLayer suffix",
			external: "projects/my-project/instances/my-instance/clusters/my-cluster",
			wantErr:  true,
		},
		{
			name:     "memory layer is a singleton, so it takes no ID",
			external: "projects/my-project/instances/my-instance/clusters/my-cluster/memoryLayers/my-layer",
			wantErr:  true,
		},
		{
			name:     "invalid format too long",
			external: "projects/my-project/instances/my-instance/clusters/my-cluster/memoryLayer/extra",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := ParseMemoryLayerExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMemoryLayerExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got := id.Parent().parent.Parent.ProjectID; got != tt.projectID {
				t.Errorf("ProjectID = %v, want %v", got, tt.projectID)
			}
			if got := id.Parent().parent.Id; got != tt.instanceID {
				t.Errorf("InstanceID = %v, want %v", got, tt.instanceID)
			}
			if got := id.Parent().ID(); got != tt.clusterID {
				t.Errorf("ClusterID = %v, want %v", got, tt.clusterID)
			}

			wantExternal := "projects/" + tt.projectID + "/instances/" + tt.instanceID + "/clusters/" + tt.clusterID + "/memoryLayer"
			if got := id.String(); got != wantExternal {
				t.Errorf("String() = %v, want %v", got, wantExternal)
			}
		})
	}
}
