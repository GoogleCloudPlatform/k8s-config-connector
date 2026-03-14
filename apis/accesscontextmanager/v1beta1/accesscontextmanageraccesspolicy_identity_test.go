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

func TestAccessContextManagerAccessPolicyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *AccessContextManagerAccessPolicyIdentity
	}{
		{
			name: "valid reference",
			ref:  "accessPolicies/123456789",
			want: &AccessContextManagerAccessPolicyIdentity{
				AccessPolicy: "123456789",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format/too/many",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://accesscontextmanager.googleapis.com/accessPolicies/123456789",
			want: &AccessContextManagerAccessPolicyIdentity{
				AccessPolicy: "123456789",
			},
		},
		{
			name: "full url without scheme",
			ref:  "accesscontextmanager.googleapis.com/accessPolicies/123456789",
			want: &AccessContextManagerAccessPolicyIdentity{
				AccessPolicy: "123456789",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AccessContextManagerAccessPolicyIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.AccessPolicy != tt.want.AccessPolicy {
					t.Errorf("AccessPolicy = %v, want %v", i.AccessPolicy, tt.want.AccessPolicy)
				}
			}
		})
	}
}
