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

func TestManagedKafkaConnectClusterIdentity_FromExternal(t *testing.T) {
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
			ref:          "projects/my-project/locations/us-central1/connectClusters/my-connect-cluster",
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantID:       "my-connect-cluster",
			wantErr:      false,
		},
		{
			name:         "full url",
			ref:          "https://managedkafka.googleapis.com/projects/my-project/locations/us-central1/connectClusters/my-connect-cluster",
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantID:       "my-connect-cluster",
			wantErr:      false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/connectClusters/my-connect-cluster",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ManagedKafkaConnectClusterIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ManagedKafkaConnectClusterIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.Project != tt.wantProject {
					t.Errorf("ManagedKafkaConnectClusterIdentity.FromExternal() Project = %v, want %v", i.Project, tt.wantProject)
				}
				if i.Location != tt.wantLocation {
					t.Errorf("ManagedKafkaConnectClusterIdentity.FromExternal() Location = %v, want %v", i.Location, tt.wantLocation)
				}
				if i.Connect_cluster != tt.wantID {
					t.Errorf("ManagedKafkaConnectClusterIdentity.FromExternal() Connect_cluster = %v, want %v", i.Connect_cluster, tt.wantID)
				}
				if got := i.String(); got != "projects/"+tt.wantProject+"/locations/"+tt.wantLocation+"/connectClusters/"+tt.wantID {
					t.Errorf("ManagedKafkaConnectClusterIdentity.String() = %v", got)
				}
			}
		})
	}
}
