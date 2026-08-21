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

func TestCloudSecurityComplianceFrameworkIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name         string
		ref          string
		wantOrg      string
		wantLocation string
		wantID       string
		wantErr      bool
	}{
		{
			name:         "valid external ref",
			ref:          "organizations/my-org/locations/us-central1/frameworks/my-framework",
			wantOrg:      "my-org",
			wantLocation: "us-central1",
			wantID:       "my-framework",
			wantErr:      false,
		},
		{
			name:         "full url",
			ref:          "https://cloudsecuritycompliance.googleapis.com/organizations/my-org/locations/us-central1/frameworks/my-framework",
			wantOrg:      "my-org",
			wantLocation: "us-central1",
			wantID:       "my-framework",
			wantErr:      false,
		},
		{
			name:    "invalid external ref",
			ref:     "organizations/my-org/frameworks/my-framework",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CloudSecurityComplianceFrameworkIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("CloudSecurityComplianceFrameworkIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.Organization != tt.wantOrg {
					t.Errorf("CloudSecurityComplianceFrameworkIdentity.FromExternal() Organization = %v, want %v", i.Organization, tt.wantOrg)
				}
				if i.Location != tt.wantLocation {
					t.Errorf("CloudSecurityComplianceFrameworkIdentity.FromExternal() Location = %v, want %v", i.Location, tt.wantLocation)
				}
				if i.Framework != tt.wantID {
					t.Errorf("CloudSecurityComplianceFrameworkIdentity.FromExternal() Framework = %v, want %v", i.Framework, tt.wantID)
				}
				if got := i.String(); got != "organizations/"+tt.wantOrg+"/locations/"+tt.wantLocation+"/frameworks/"+tt.wantID {
					t.Errorf("CloudSecurityComplianceFrameworkIdentity.String() = %v", got)
				}
				if got := i.ParentString(); got != "organizations/"+tt.wantOrg+"/locations/"+tt.wantLocation {
					t.Errorf("CloudSecurityComplianceFrameworkIdentity.ParentString() = %v", got)
				}
			}
		})
	}
}
