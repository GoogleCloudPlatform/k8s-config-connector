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

func TestNetworkConnectivityTransportIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name         string
		ref          string
		wantIdentity NetworkConnectivityTransportIdentity
		wantErr      bool
	}{
		{
			name: "valid external ref",
			ref:  "projects/my-project/locations/us-central1/transports/my-transport",
			wantIdentity: NetworkConnectivityTransportIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				Transport: "my-transport",
			},
			wantErr: false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/transports/my-transport",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &NetworkConnectivityTransportIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Fatalf("NetworkConnectivityTransportIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.wantIdentity, *i); diff != "" {
					t.Errorf("NetworkConnectivityTransportIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != tt.ref {
					t.Errorf("NetworkConnectivityTransportIdentity.String() = %v, want %v", got, tt.ref)
				}
				wantParent := "projects/my-project/locations/us-central1"
				if gotParent := i.ParentString(); gotParent != wantParent {
					t.Errorf("NetworkConnectivityTransportIdentity.ParentString() = %v, want %v", gotParent, wantParent)
				}
			}
		})
	}
}
