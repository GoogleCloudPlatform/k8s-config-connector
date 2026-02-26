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

func TestAccessContextManagerAccessLevelIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *AccessContextManagerAccessLevelIdentity
	}{
		{
			name: "valid reference",
			ref:  "accessPolicies/12345/accessLevels/my-level",
			want: &AccessContextManagerAccessLevelIdentity{
				AccessPolicy: "12345",
				AccessLevel:  "my-level",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://accesscontextmanager.googleapis.com/accessPolicies/12345/accessLevels/my-level",
			want: &AccessContextManagerAccessLevelIdentity{
				AccessPolicy: "12345",
				AccessLevel:  "my-level",
			},
		},
		{
			name: "full url with //",
			ref:  "//accesscontextmanager.googleapis.com/accessPolicies/12345/accessLevels/my-level",
			want: &AccessContextManagerAccessLevelIdentity{
				AccessPolicy: "12345",
				AccessLevel:  "my-level",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AccessContextManagerAccessLevelIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.AccessPolicy != tt.want.AccessPolicy {
					t.Errorf("AccessPolicy = %v, want %v", i.AccessPolicy, tt.want.AccessPolicy)
				}
				if i.AccessLevel != tt.want.AccessLevel {
					t.Errorf("AccessLevel = %v, want %v", i.AccessLevel, tt.want.AccessLevel)
				}
			}
		})
	}
}
