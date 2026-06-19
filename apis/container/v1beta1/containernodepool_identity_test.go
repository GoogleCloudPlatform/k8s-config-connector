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

	"github.com/google/go-cmp/cmp"
)

func TestContainerNodePoolIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ContainerNodePoolIdentity
	}{
		{
			name: "valid regional reference",
			ref:  "projects/my-project/locations/us-central1/clusters/my-cluster/nodePools/my-nodepool",
			want: &ContainerNodePoolIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Cluster:  "my-cluster",
				NodePool: "my-nodepool",
			},
		},
		{
			name: "valid zonal reference",
			ref:  "projects/my-project/zones/us-central1-a/clusters/my-cluster/nodePools/my-nodepool",
			want: &ContainerNodePoolIdentity{
				Project:  "my-project",
				Location: "us-central1-a",
				Cluster:  "my-cluster",
				NodePool: "my-nodepool",
			},
		},
		{
			name: "valid full regional url",
			ref:  "https://container.googleapis.com/v1/projects/my-project/locations/us-central1/clusters/my-cluster/nodePools/my-nodepool",
			want: &ContainerNodePoolIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Cluster:  "my-cluster",
				NodePool: "my-nodepool",
			},
		},
		{
			name: "valid full zonal url",
			ref:  "https://container.googleapis.com/v1/projects/my-project/zones/us-central1-a/clusters/my-cluster/nodePools/my-nodepool",
			want: &ContainerNodePoolIdentity{
				Project:  "my-project",
				Location: "us-central1-a",
				Cluster:  "my-cluster",
				NodePool: "my-nodepool",
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
			i := &ContainerNodePoolIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("ContainerNodePoolIdentity mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestContainerNodePoolIdentity_String(t *testing.T) {
	tests := []struct {
		name     string
		identity *ContainerNodePoolIdentity
		expected string
	}{
		{
			name: "regional",
			identity: &ContainerNodePoolIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Cluster:  "my-cluster",
				NodePool: "my-nodepool",
			},
			expected: "projects/my-project/locations/us-central1/clusters/my-cluster/nodePools/my-nodepool",
		},
		{
			name: "zonal",
			identity: &ContainerNodePoolIdentity{
				Project:  "my-project",
				Location: "us-central1-a",
				Cluster:  "my-cluster",
				NodePool: "my-nodepool",
			},
			expected: "projects/my-project/zones/us-central1-a/clusters/my-cluster/nodePools/my-nodepool",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := tt.identity.String(); actual != tt.expected {
				t.Errorf("String() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestContainerNodePoolIdentity_ParentString(t *testing.T) {
	tests := []struct {
		name     string
		identity *ContainerNodePoolIdentity
		expected string
	}{
		{
			name: "regional parent",
			identity: &ContainerNodePoolIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Cluster:  "my-cluster",
				NodePool: "my-nodepool",
			},
			expected: "projects/my-project/locations/us-central1/clusters/my-cluster",
		},
		{
			name: "zonal parent",
			identity: &ContainerNodePoolIdentity{
				Project:  "my-project",
				Location: "us-central1-a",
				Cluster:  "my-cluster",
				NodePool: "my-nodepool",
			},
			expected: "projects/my-project/zones/us-central1-a/clusters/my-cluster",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := tt.identity.ParentString(); actual != tt.expected {
				t.Errorf("ParentString() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
