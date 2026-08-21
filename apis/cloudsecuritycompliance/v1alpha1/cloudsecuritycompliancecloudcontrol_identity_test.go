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

func TestCloudSecurityComplianceCloudControlIdentity_FromExternal(t *testing.T) {
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
			ref:          "organizations/my-org/locations/us-central1/cloudControls/my-control",
			wantOrg:      "my-org",
			wantLocation: "us-central1",
			wantID:       "my-control",
			wantErr:      false,
		},
		{
			name:         "full url",
			ref:          "https://cloudsecuritycompliance.googleapis.com/organizations/my-org/locations/us-central1/cloudControls/my-control",
			wantOrg:      "my-org",
			wantLocation: "us-central1",
			wantID:       "my-control",
			wantErr:      false,
		},
		{
			name:    "invalid external ref",
			ref:     "organizations/my-org/cloudControls/my-control",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CloudSecurityComplianceCloudControlIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("CloudSecurityComplianceCloudControlIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.Organization != tt.wantOrg {
					t.Errorf("CloudSecurityComplianceCloudControlIdentity.FromExternal() Organization = %v, want %v", i.Organization, tt.wantOrg)
				}
				if i.Location != tt.wantLocation {
					t.Errorf("CloudSecurityComplianceCloudControlIdentity.FromExternal() Location = %v, want %v", i.Location, tt.wantLocation)
				}
				if i.CloudControl != tt.wantID {
					t.Errorf("CloudSecurityComplianceCloudControlIdentity.FromExternal() CloudControl = %v, want %v", i.CloudControl, tt.wantID)
				}
				if got := i.String(); got != "organizations/"+tt.wantOrg+"/locations/"+tt.wantLocation+"/cloudControls/"+tt.wantID {
					t.Errorf("CloudSecurityComplianceCloudControlIdentity.String() = %v", got)
				}
			}
		})
	}
}
