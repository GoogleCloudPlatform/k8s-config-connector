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

func TestManagedKafkaConsumerGroupIdentity_IdentityV2(t *testing.T) {
	id := &ManagedKafkaConsumerGroupIdentity{
		Project:       "my-project",
		Location:      "us-central1",
		Cluster:       "my-cluster",
		ConsumerGroup: "my-group",
	}

	wantStr := "projects/my-project/locations/us-central1/clusters/my-cluster/consumerGroups/my-group"
	if got := id.String(); got != wantStr {
		t.Errorf("String() = %v, want %v", got, wantStr)
	}

	wantHost := "managedkafka.googleapis.com"
	if got := id.Host(); got != wantHost {
		t.Errorf("Host() = %v, want %v", got, wantHost)
	}
}

func TestManagedKafkaConsumerGroupIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ManagedKafkaConsumerGroupIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/clusters/my-cluster/consumerGroups/my-group",
			want: &ManagedKafkaConsumerGroupIdentity{
				Project:       "my-project",
				Location:      "us-central1",
				Cluster:       "my-cluster",
				ConsumerGroup: "my-group",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ManagedKafkaConsumerGroupIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project {
					t.Errorf("Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Location != tt.want.Location {
					t.Errorf("Location = %v, want %v", i.Location, tt.want.Location)
				}
				if i.Cluster != tt.want.Cluster {
					t.Errorf("Cluster = %v, want %v", i.Cluster, tt.want.Cluster)
				}
				if i.ConsumerGroup != tt.want.ConsumerGroup {
					t.Errorf("ConsumerGroup = %v, want %v", i.ConsumerGroup, tt.want.ConsumerGroup)
				}
			}
		})
	}
}
