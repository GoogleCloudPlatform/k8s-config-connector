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

func TestCloudIdentityDeviceIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *CloudIdentityDeviceIdentity
	}{
		{
			name: "valid reference",
			ref:  "devices/123456789",
			want: &CloudIdentityDeviceIdentity{
				Device: "123456789",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://cloudidentity.googleapis.com/devices/123456789",
			want: &CloudIdentityDeviceIdentity{
				Device: "123456789",
			},
		},
		{
			name: "full url with version",
			ref:  "https://cloudidentity.googleapis.com/v1beta1/devices/123456789",
			want: &CloudIdentityDeviceIdentity{
				Device: "123456789",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CloudIdentityDeviceIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Device != tt.want.Device {
					t.Errorf("FromExternal() got = %v, want %v", i.Device, tt.want.Device)
				}
			}
		})
	}
}
