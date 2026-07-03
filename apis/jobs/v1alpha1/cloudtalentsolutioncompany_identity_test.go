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

func TestCloudTalentSolutionCompanyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *CloudTalentSolutionCompanyIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/tenants/my-tenant/companies/my-company",
			want: &CloudTalentSolutionCompanyIdentity{
				Project: "my-project",
				Tenant:  "my-tenant",
				Company: "my-company",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://jobs.googleapis.com/projects/my-project/tenants/my-tenant/companies/my-company",
			want: &CloudTalentSolutionCompanyIdentity{
				Project: "my-project",
				Tenant:  "my-tenant",
				Company: "my-company",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CloudTalentSolutionCompanyIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project {
					t.Errorf("Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Tenant != tt.want.Tenant {
					t.Errorf("Tenant = %v, want %v", i.Tenant, tt.want.Tenant)
				}
				if i.Company != tt.want.Company {
					t.Errorf("Company = %v, want %v", i.Company, tt.want.Company)
				}
				if gotParent := i.ParentString(); gotParent != "projects/my-project/tenants/my-tenant" {
					t.Errorf("ParentString() = %v, want %v", gotParent, "projects/my-project/tenants/my-tenant")
				}
			}
		})
	}
}
