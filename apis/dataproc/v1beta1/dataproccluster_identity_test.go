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

func TestDataprocClusterIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DataprocClusterIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/regions/us-central1/clusters/my-cluster",
			want: &DataprocClusterIdentity{
				Project: "my-project",
				Region:  "us-central1",
				Cluster: "my-cluster",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "full url unversioned",
			ref:     "https://dataproc.googleapis.com/projects/my-project/regions/us-central1/clusters/my-cluster",
			wantErr: true,
		},
		{
			name: "full url versioned (v1)",
			ref:  "https://dataproc.googleapis.com/v1/projects/my-project/regions/us-central1/clusters/my-cluster",
			want: &DataprocClusterIdentity{
				Project: "my-project",
				Region:  "us-central1",
				Cluster: "my-cluster",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DataprocClusterIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project {
					t.Errorf("Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Region != tt.want.Region {
					t.Errorf("Region = %v, want %v", i.Region, tt.want.Region)
				}
				if i.Cluster != tt.want.Cluster {
					t.Errorf("Cluster = %v, want %v", i.Cluster, tt.want.Cluster)
				}
			}
		})
	}
}

func TestDataprocClusterIdentity_String(t *testing.T) {
	identity := &DataprocClusterIdentity{
		Project: "my-project",
		Region:  "us-central1",
		Cluster: "my-cluster",
	}
	want := "v1/projects/my-project/regions/us-central1/clusters/my-cluster"
	if got := identity.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}
