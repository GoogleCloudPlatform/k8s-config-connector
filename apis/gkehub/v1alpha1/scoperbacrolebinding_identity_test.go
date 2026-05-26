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

func TestGKEHubScopeRBACRoleBindingIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     *GKEHubScopeRBACRoleBindingIdentity
		wantErr  bool
	}{
		{
			name:     "canonical format",
			external: "projects/my-project/locations/global/scopes/my-scope/rbacrolebindings/my-binding",
			want: &GKEHubScopeRBACRoleBindingIdentity{
				ProjectID:         "my-project",
				Location:          "global",
				ScopeID:           "my-scope",
				RBACRoleBindingID: "my-binding",
			},
			wantErr: false,
		},
		{
			name:     "invalid format",
			external: "projects/my-project/locations/global/scopes/my-scope",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &GKEHubScopeRBACRoleBindingIdentity{}
			if err := i.FromExternal(tt.external); (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(i, tt.want) {
					t.Errorf("FromExternal() = %v, want %v", i, tt.want)
				}
				if i.String() != tt.external {
					t.Errorf("String() = %v, want %v", i.String(), tt.external)
				}
			}
		})
	}
}
