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

func TestGKEHubScopeIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *GKEHubScopeIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/global/scopes/my-scope",
			want: &GKEHubScopeIdentity{
				ProjectID: "my-project",
				Location:  "global",
				ScopeID:   "my-scope",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://gkehub.googleapis.com/projects/my-project/locations/global/scopes/my-scope",
			want: &GKEHubScopeIdentity{
				ProjectID: "my-project",
				Location:  "global",
				ScopeID:   "my-scope",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &GKEHubScopeIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.ProjectID != tt.want.ProjectID {
					t.Errorf("ProjectID = %v, want %v", i.ProjectID, tt.want.ProjectID)
				}
				if i.Location != tt.want.Location {
					t.Errorf("Location = %v, want %v", i.Location, tt.want.Location)
				}
				if i.ScopeID != tt.want.ScopeID {
					t.Errorf("ScopeID = %v, want %v", i.ScopeID, tt.want.ScopeID)
				}
			}
		})
	}
}
