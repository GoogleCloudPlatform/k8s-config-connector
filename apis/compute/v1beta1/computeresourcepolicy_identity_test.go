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

func TestComputeResourcePolicyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeResourcePolicyIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/regions/us-central1/resourcePolicies/my-policy",
			want: &ComputeResourcePolicyIdentity{
				Project:        "my-project",
				Region:         "us-central1",
				ResourcePolicy: "my-policy",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://compute.googleapis.com/projects/my-project/regions/us-central1/resourcePolicies/my-policy",
			want: &ComputeResourcePolicyIdentity{
				Project:        "my-project",
				Region:         "us-central1",
				ResourcePolicy: "my-policy",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeResourcePolicyIdentity{}
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
				if i.ResourcePolicy != tt.want.ResourcePolicy {
					t.Errorf("ResourcePolicy = %v, want %v", i.ResourcePolicy, tt.want.ResourcePolicy)
				}
			}
		})
	}
}
