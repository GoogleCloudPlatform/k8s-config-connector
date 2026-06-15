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

func TestCloudTalentSolutionCompanyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *CloudTalentSolutionCompanyIdentity
		wantErr bool
	}{
		{
			name: "valid external ref",
			ref:  "projects/my-project/tenants/my-tenant/companies/my-company",
			want: &CloudTalentSolutionCompanyIdentity{
				Project: "my-project",
				Tenant:  "my-tenant",
				Company: "my-company",
			},
			wantErr: false,
		},
		{
			name:    "invalid external ref (missing tenant)",
			ref:     "projects/my-project/companies/my-company",
			wantErr: true,
		},
		{
			name:    "invalid external ref (missing company)",
			ref:     "projects/my-project/tenants/my-tenant",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CloudTalentSolutionCompanyIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CloudTalentSolutionCompanyIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("CloudTalentSolutionCompanyIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != tt.ref {
					t.Errorf("CloudTalentSolutionCompanyIdentity.String() = %v, want %v", got, tt.ref)
				}
			}
		})
	}
}
