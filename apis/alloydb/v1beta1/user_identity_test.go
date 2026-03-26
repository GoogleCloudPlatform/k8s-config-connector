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

func TestAlloyDBUserIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     *AlloyDBUserIdentity
		wantErr  bool
	}{
		{
			name:     "valid external",
			external: "projects/my-project/locations/us-central1/clusters/my-cluster/users/my-user",
			want: &AlloyDBUserIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Cluster:  "my-cluster",
				User:     "my-user",
			},
			wantErr: false,
		},
		{
			name:     "invalid external - too few tokens",
			external: "projects/my-project/locations/us-central1/clusters/my-cluster",
			wantErr:  true,
		},
		{
			name:     "invalid external - wrong prefix",
			external: "organizations/my-org/locations/us-central1/clusters/my-cluster/users/my-user",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AlloyDBUserIdentity{}
			if err := i.FromExternal(tt.external); (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project || i.Location != tt.want.Location || i.Cluster != tt.want.Cluster || i.User != tt.want.User {
					t.Errorf("FromExternal() got = %v, want %v", i, tt.want)
				}
			}
		})
	}
}

func TestAlloyDBUserIdentity_String(t *testing.T) {
	identity := &AlloyDBUserIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Cluster:  "my-cluster",
		User:     "my-user",
	}
	want := "projects/my-project/locations/us-central1/clusters/my-cluster/users/my-user"
	if got := identity.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}
