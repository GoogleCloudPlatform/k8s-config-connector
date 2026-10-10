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

	"github.com/google/go-cmp/cmp"
)

func TestParseBigtableMemoryLayerExternal(t *testing.T) {
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
			name:     "invalid external prefix",
			external: "project/my-project/instances/my-instance/clusters/my-cluster/memoryLayer",
			wantErr:  true,
		},
		{
			name:     "invalid format too short",
			external: "projects/my-project/instances/my-instance/clusters/my-cluster",
			wantErr:  true,
		},
		{
			name:     "invalid format too long",
			external: "projects/my-project/instances/my-instance/clusters/my-cluster/memoryLayer/extra",
			wantErr:  true,
		},
		{
			name:     "wrong suffix",
			external: "projects/my-project/instances/my-instance/clusters/my-cluster/otherLayer",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := ParseBigtableMemoryLayerExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBigtableMemoryLayerExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if p.ParentString() != "projects/"+tt.projectID+"/instances/"+tt.instanceID {
					t.Errorf("ParentString() = %v, want %v", p.ParentString(), "projects/"+tt.projectID+"/instances/"+tt.instanceID)
				}
				if p.ID() != tt.clusterID {
					t.Errorf("ClusterID = %v, want %v", p.ID(), tt.clusterID)
				}
			}
		})
	}
}

func TestBigtableMemoryLayerIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *BigtableMemoryLayerIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/instances/my-instance/clusters/my-cluster/memoryLayer",
			want: &BigtableMemoryLayerIdentity{
				Project:  "my-project",
				Instance: "my-instance",
				Cluster:  "my-cluster",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://bigtable.googleapis.com/projects/my-project/instances/my-instance/clusters/my-cluster/memoryLayer",
			want: &BigtableMemoryLayerIdentity{
				Project:  "my-project",
				Instance: "my-instance",
				Cluster:  "my-cluster",
			},
		},
		{
			name: "relative url with host",
			ref:  "//bigtable.googleapis.com/projects/my-project/instances/my-instance/clusters/my-cluster/memoryLayer",
			want: &BigtableMemoryLayerIdentity{
				Project:  "my-project",
				Instance: "my-instance",
				Cluster:  "my-cluster",
			},
		},
		{
			name:    "invalid resource suffix",
			ref:     "projects/my-project/instances/my-instance/clusters/my-cluster/otherLayer",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BigtableMemoryLayerIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestBigtableMemoryLayerIdentity_String(t *testing.T) {
	i := &BigtableMemoryLayerIdentity{
		Project:  "my-project",
		Instance: "my-instance",
		Cluster:  "my-cluster",
	}
	want := "projects/my-project/instances/my-instance/clusters/my-cluster/memoryLayer"
	got := i.String()
	if got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func TestBigtableMemoryLayerIdentity_ParentString(t *testing.T) {
	i := &BigtableMemoryLayerIdentity{
		Project:  "my-project",
		Instance: "my-instance",
		Cluster:  "my-cluster",
	}
	want := "projects/my-project/instances/my-instance/clusters/my-cluster"
	got := i.ParentString()
	if got != want {
		t.Errorf("ParentString() = %v, want %v", got, want)
	}
}
