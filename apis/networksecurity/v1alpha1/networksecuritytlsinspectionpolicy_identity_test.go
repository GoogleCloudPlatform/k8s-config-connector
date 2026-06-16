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
	"reflect"
	"testing"
)

func TestNetworkSecurityTLSInspectionPolicyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *NetworkSecurityTLSInspectionPolicyIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/tlsInspectionPolicies/my-policy",
			want: &NetworkSecurityTLSInspectionPolicyIdentity{
				Project:             "my-project",
				Location:            "us-central1",
				TlsInspectionPolicy: "my-policy",
			},
		},
		{
			name: "full url",
			ref:  "https://networksecurity.googleapis.com/projects/my-project/locations/us-central1/tlsInspectionPolicies/my-policy",
			want: &NetworkSecurityTLSInspectionPolicyIdentity{
				Project:             "my-project",
				Location:            "us-central1",
				TlsInspectionPolicy: "my-policy",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &NetworkSecurityTLSInspectionPolicyIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(i, tt.want) {
					t.Errorf("got = %v, want %v", i, tt.want)
				}
			}
		})
	}
}
