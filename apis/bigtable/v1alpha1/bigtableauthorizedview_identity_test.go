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

	"github.com/google/go-cmp/cmp"
)

func TestParseAuthorizedViewExternal(t *testing.T) {
	tests := []struct {
		name       string
		external   string
		wantErr    bool
		projectID  string
		instanceID string
		tableID    string
		viewID     string
	}{
		{
			name:       "valid external name",
			external:   "projects/my-project/instances/my-instance/tables/my-table/authorizedViews/my-view",
			wantErr:    false,
			projectID:  "my-project",
			instanceID: "my-instance",
			tableID:    "my-table",
			viewID:     "my-view",
		},
		{
			name:     "invalid external prefix",
			external: "project/my-project/instances/my-instance/tables/my-table/authorizedViews/my-view",
			wantErr:  true,
		},
		{
			name:     "invalid format too short",
			external: "projects/my-project/instances/my-instance/tables/my-table",
			wantErr:  true,
		},
		{
			name:     "invalid format too long",
			external: "projects/my-project/instances/my-instance/tables/my-table/authorizedViews/my-view/extra",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, resourceID, err := ParseAuthorizedViewExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAuthorizedViewExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if p.Parent.Parent.ProjectID != tt.projectID {
					t.Errorf("ProjectID = %v, want %v", p.Parent.Parent.ProjectID, tt.projectID)
				}
				if p.Parent.Id != tt.instanceID {
					t.Errorf("InstanceID = %v, want %v", p.Parent.Id, tt.instanceID)
				}
				if p.Id != tt.tableID {
					t.Errorf("TableID = %v, want %v", p.Id, tt.tableID)
				}
				if resourceID != tt.viewID {
					t.Errorf("ViewID = %v, want %v", resourceID, tt.viewID)
				}
			}
		})
	}
}

func TestBigtableAuthorizedViewIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *AuthorizedViewIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/instances/my-instance/tables/my-table/authorizedViews/my-view",
			want: &AuthorizedViewIdentity{
				Project:        "my-project",
				Instance:       "my-instance",
				Table:          "my-table",
				AuthorizedView: "my-view",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://bigtable.googleapis.com/projects/my-project/instances/my-instance/tables/my-table/authorizedViews/my-view",
			want: &AuthorizedViewIdentity{
				Project:        "my-project",
				Instance:       "my-instance",
				Table:          "my-table",
				AuthorizedView: "my-view",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AuthorizedViewIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
