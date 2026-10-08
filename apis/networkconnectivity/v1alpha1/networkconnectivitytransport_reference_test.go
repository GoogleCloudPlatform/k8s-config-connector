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

func TestNetworkConnectivityTransportRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/locations/us-central1/transports/my-transport",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/locations/us-central1/transports/my-transport",
			wantErr: true,
		},
		{
			name:    "missing location",
			ref:     "projects/my-project/transports/my-transport",
			wantErr: true,
		},
		{
			name:    "missing transport",
			ref:     "projects/my-project/locations/us-central1",
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &NetworkConnectivityTransportRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("NetworkConnectivityTransportRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetworkConnectivityRemoteTransportProfileRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/locations/us-central1/remoteTransportProfiles/my-profile",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/locations/us-central1/remoteTransportProfiles/my-profile",
			wantErr: true,
		},
		{
			name:    "missing location",
			ref:     "projects/my-project/remoteTransportProfiles/my-profile",
			wantErr: true,
		},
		{
			name:    "missing remoteTransportProfile",
			ref:     "projects/my-project/locations/us-central1",
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &NetworkConnectivityRemoteTransportProfileRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("NetworkConnectivityRemoteTransportProfileRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
