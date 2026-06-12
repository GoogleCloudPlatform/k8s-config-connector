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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

func TestNetworkSecurityTLSInspectionPolicyIdentity(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     *NetworkSecurityTLSInspectionPolicyIdentity
		wantErr  bool
	}{
		{
			name:     "basic",
			external: "projects/my-project/locations/us-central1/tlsInspectionPolicies/my-policy",
			want: &NetworkSecurityTLSInspectionPolicyIdentity{
				Project:             "my-project",
				Location:            "us-central1",
				TlsInspectionPolicy: "my-policy",
			},
		},
		{
			name:     "invalid format",
			external: "invalid/my-project/locations/us-central1/tlsInspectionPolicies/my-policy",
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := &NetworkSecurityTLSInspectionPolicyIdentity{}
			err := got.FromExternal(tc.external)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !tc.wantErr {
				if got.String() != tc.external {
					t.Errorf("String() = %v, want %v", got.String(), tc.external)
				}
			}
		})
	}
}

func TestNetworkSecurityTLSInspectionPolicyIdentity_Interfaces(t *testing.T) {
	var _ identity.IdentityV2 = &NetworkSecurityTLSInspectionPolicyIdentity{}
	var _ identity.Resource = &NetworkSecurityTLSInspectionPolicy{}
}
