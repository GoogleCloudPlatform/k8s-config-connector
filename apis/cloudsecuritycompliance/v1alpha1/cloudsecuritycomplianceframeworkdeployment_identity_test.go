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

func TestCloudSecurityComplianceFrameworkDeploymentIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *CloudSecurityComplianceFrameworkDeploymentIdentity
	}{
		{
			name: "valid reference",
			ref:  "organizations/123456789/locations/global/frameworkDeployments/my-framework-deployment",
			want: &CloudSecurityComplianceFrameworkDeploymentIdentity{
				Organization:        "123456789",
				Location:            "global",
				FrameworkDeployment: "my-framework-deployment",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://cloudsecuritycompliance.googleapis.com/organizations/123456789/locations/global/frameworkDeployments/my-framework-deployment",
			want: &CloudSecurityComplianceFrameworkDeploymentIdentity{
				Organization:        "123456789",
				Location:            "global",
				FrameworkDeployment: "my-framework-deployment",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CloudSecurityComplianceFrameworkDeploymentIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Organization != tt.want.Organization {
					t.Errorf("Organization = %v, want %v", i.Organization, tt.want.Organization)
				}
				if i.Location != tt.want.Location {
					t.Errorf("Location = %v, want %v", i.Location, tt.want.Location)
				}
				if i.FrameworkDeployment != tt.want.FrameworkDeployment {
					t.Errorf("FrameworkDeployment = %v, want %v", i.FrameworkDeployment, tt.want.FrameworkDeployment)
				}
			}
		})
	}
}

func TestCloudSecurityComplianceFrameworkDeploymentIdentity_String(t *testing.T) {
	i := &CloudSecurityComplianceFrameworkDeploymentIdentity{
		Organization:        "123456789",
		Location:            "global",
		FrameworkDeployment: "my-framework-deployment",
	}
	want := "organizations/123456789/locations/global/frameworkDeployments/my-framework-deployment"
	if got := i.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}
