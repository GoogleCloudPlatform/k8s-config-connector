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
	"github.com/google/go-cmp/cmp"
)

func TestNetworkSecurityUrlListIdentity(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     *NetworkSecurityUrlListIdentity
		wantErr  bool
	}{
		{
			name:     "basic",
			external: "projects/my-project/locations/us-central1/urlLists/my-urllist",
			want: &NetworkSecurityUrlListIdentity{
				Project:  "my-project",
				Location: "us-central1",
				UrlList:  "my-urllist",
			},
		},
		{
			name:     "invalid format",
			external: "invalid/my-project/locations/us-central1/urlLists/my-urllist",
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := &NetworkSecurityUrlListIdentity{}
			err := got.FromExternal(tc.external)
			if (err != nil) != tc.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !tc.wantErr {
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got.String() != tc.external {
					t.Errorf("String() = %q, want %q", got.String(), tc.external)
				}
			}
		})
	}
}

func TestNetworkSecurityUrlListIdentity_Interfaces(t *testing.T) {
	var _ identity.IdentityV2 = &NetworkSecurityUrlListIdentity{}
	var _ identity.Resource = &NetworkSecurityUrlList{}
}
