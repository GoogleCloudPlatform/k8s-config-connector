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

func TestAccessContextManagerAccessLevelConditionIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     *AccessContextManagerAccessLevelConditionIdentity
		wantErr  bool
	}{
		{
			name:     "canonical format",
			external: "accessPolicies/123/accessLevels/level1/condition",
			want: &AccessContextManagerAccessLevelConditionIdentity{
				AccessPolicy: "123",
				AccessLevel:  "level1",
			},
			wantErr: false,
		},
		{
			name:     "with host",
			external: "accesscontextmanager.googleapis.com/accessPolicies/123/accessLevels/level1/condition",
			want: &AccessContextManagerAccessLevelConditionIdentity{
				AccessPolicy: "123",
				AccessLevel:  "level1",
			},
			wantErr: false,
		},
		{
			name:     "invalid format",
			external: "accessPolicies/123/accessLevels/level1",
			want:     nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AccessContextManagerAccessLevelConditionIdentity{}
			if err := i.FromExternal(tt.external); (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.AccessPolicy != tt.want.AccessPolicy {
					t.Errorf("FromExternal() AccessPolicy = %v, want %v", i.AccessPolicy, tt.want.AccessPolicy)
				}
				if i.AccessLevel != tt.want.AccessLevel {
					t.Errorf("FromExternal() AccessLevel = %v, want %v", i.AccessLevel, tt.want.AccessLevel)
				}
			}
		})
	}
}

func TestAccessContextManagerAccessLevelConditionIdentity_String(t *testing.T) {
	i := &AccessContextManagerAccessLevelConditionIdentity{
		AccessPolicy: "123",
		AccessLevel:  "level1",
	}
	want := "accessPolicies/123/accessLevels/level1/condition"
	if got := i.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}
