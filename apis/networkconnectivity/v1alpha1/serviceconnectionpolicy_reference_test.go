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

func TestNetworkConnectivityServiceConnectionPolicyRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/serviceConnectionPolicies/my-policy",
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://networkconnectivity.googleapis.com/projects/my-project/locations/us-central1/serviceConnectionPolicies/my-policy",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &NetworkConnectivityServiceConnectionPolicyRef{
				External: tt.ref,
			}
			err := r.ValidateExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetworkConnectivityServiceConnectionPolicyRef_GVK(t *testing.T) {
	r := &NetworkConnectivityServiceConnectionPolicyRef{}
	gvk := r.GetGVK()
	if gvk != NetworkConnectivityServiceConnectionPolicyGVK {
		t.Errorf("GetGVK() = %v, want %v", gvk, NetworkConnectivityServiceConnectionPolicyGVK)
	}
}
