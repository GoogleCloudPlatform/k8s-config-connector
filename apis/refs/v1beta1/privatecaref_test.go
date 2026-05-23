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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

func TestPrivateCACAPoolIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *PrivateCACAPoolIdentity
		wantErr bool
	}{
		{
			name: "valid caPool reference",
			ref:  "projects/test-project/locations/us-central1/caPools/test-pool",
			want: &PrivateCACAPoolIdentity{
				Project:  "test-project",
				Location: "us-central1",
				CAPool:   "test-pool",
			},
			wantErr: false,
		},
		{
			name:    "invalid format - extra segment",
			ref:     "projects/test-project/locations/us-central1/caPools/test-pool/extra",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid format - missing location",
			ref:     "projects/test-project/caPools/test-pool",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &PrivateCACAPoolIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrivateCACAPoolIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project || i.Location != tt.want.Location || i.CAPool != tt.want.CAPool {
					t.Errorf("PrivateCACAPoolIdentity.FromExternal() = %+v, want %+v", i, tt.want)
				}
			}
		})
	}
}

func TestPrivateCACAPoolRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		wantErr  bool
	}{
		{
			name:     "valid external",
			external: "projects/my-project/locations/global/caPools/my-pool",
			wantErr:  false,
		},
		{
			name:     "invalid external",
			external: "projects/my-project/caPools/my-pool",
			wantErr:  true,
		},
		{
			name:     "empty external",
			external: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PrivateCACAPoolRef{}
			if err := r.ValidateExternal(tt.external); (err != nil) != tt.wantErr {
				t.Errorf("PrivateCACAPoolRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrivateCACAPoolImplementsIdentity(t *testing.T) {
	var _ identity.Identity = &PrivateCACAPoolIdentity{}
	var _ identity.IdentityV2 = &PrivateCACAPoolIdentity{}
}
