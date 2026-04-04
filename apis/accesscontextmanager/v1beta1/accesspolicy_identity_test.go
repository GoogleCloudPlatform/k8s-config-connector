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

func TestAccessPolicyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name         string
		ref          string
		wantID       string
		wantString   string
		wantErr      bool
	}{
		{
			name:       "simple ID",
			ref:        "accessPolicies/123",
			wantID:     "123",
			wantString: "accessPolicies/123",
			wantErr:    false,
		},
		{
			name:       "full URL",
			ref:        "https://accesscontextmanager.googleapis.com/accessPolicies/123",
			wantID:     "123",
			wantString: "accessPolicies/123",
			wantErr:    false,
		},
		{
			name:       "full URL with double slash",
			ref:        "//accesscontextmanager.googleapis.com/accessPolicies/123",
			wantID:     "123",
			wantString: "accessPolicies/123",
			wantErr:    false,
		},
		{
			name:       "invalid format",
			ref:        "organizations/123",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AccessPolicyIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.AccessPolicy != tt.wantID {
					t.Errorf("FromExternal() AccessPolicy = %v, want %v", i.AccessPolicy, tt.wantID)
				}
				if i.String() != tt.wantString {
					t.Errorf("String() = %v, want %v", i.String(), tt.wantString)
				}
			}
		})
	}
}
