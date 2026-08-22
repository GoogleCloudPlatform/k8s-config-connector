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

func TestCloudSecurityComplianceFrameworkDeploymentIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *CloudSecurityComplianceFrameworkDeploymentIdentity
		wantErr bool
	}{
		{
			name: "valid external ref",
			ref:  "organizations/my-org/locations/us-central1/frameworkDeployments/my-deployment",
			want: &CloudSecurityComplianceFrameworkDeploymentIdentity{
				Organization:        "my-org",
				Location:            "us-central1",
				FrameworkDeployment: "my-deployment",
			},
			wantErr: false,
		},
		{
			name: "full url",
			ref:  "https://cloudsecuritycompliance.googleapis.com/organizations/my-org/locations/us-central1/frameworkDeployments/my-deployment",
			want: &CloudSecurityComplianceFrameworkDeploymentIdentity{
				Organization:        "my-org",
				Location:            "us-central1",
				FrameworkDeployment: "my-deployment",
			},
			wantErr: false,
		},
		{
			name:    "invalid external ref",
			ref:     "organizations/my-org/frameworkDeployments/my-deployment",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CloudSecurityComplianceFrameworkDeploymentIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CloudSecurityComplianceFrameworkDeploymentIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("CloudSecurityComplianceFrameworkDeploymentIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
				expectedString := "organizations/" + tt.want.Organization + "/locations/" + tt.want.Location + "/frameworkDeployments/" + tt.want.FrameworkDeployment
				if got := i.String(); got != expectedString {
					t.Errorf("CloudSecurityComplianceFrameworkDeploymentIdentity.String() = %v, want %v", got, expectedString)
				}
				expectedParent := "organizations/" + tt.want.Organization + "/locations/" + tt.want.Location
				if got := i.ParentString(); got != expectedParent {
					t.Errorf("CloudSecurityComplianceFrameworkDeploymentIdentity.ParentString() = %v, want %v", got, expectedParent)
				}
			}
		})
	}
}
