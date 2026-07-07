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

func TestSecurityCenterManagementEventThreatDetectionCustomModuleIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *SecurityCenterManagementEventThreatDetectionCustomModuleIdentity
	}{
		{
			name: "valid reference",
			ref:  "organizations/123456789/locations/global/eventThreatDetectionCustomModules/my-module",
			want: &SecurityCenterManagementEventThreatDetectionCustomModuleIdentity{
				Organization:                     "123456789",
				Location:                         "global",
				EventThreatDetectionCustomModule: "my-module",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://securitycentermanagement.googleapis.com/organizations/123456789/locations/global/eventThreatDetectionCustomModules/my-module",
			want: &SecurityCenterManagementEventThreatDetectionCustomModuleIdentity{
				Organization:                     "123456789",
				Location:                         "global",
				EventThreatDetectionCustomModule: "my-module",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &SecurityCenterManagementEventThreatDetectionCustomModuleIdentity{}
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
				if i.EventThreatDetectionCustomModule != tt.want.EventThreatDetectionCustomModule {
					t.Errorf("EventThreatDetectionCustomModule = %v, want %v", i.EventThreatDetectionCustomModule, tt.want.EventThreatDetectionCustomModule)
				}
			}
		})
	}
}
