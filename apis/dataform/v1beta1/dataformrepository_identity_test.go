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

func TestDataformRepositoryIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     *DataformRepositoryIdentity
		wantErr  bool
	}{
		{
			name:     "valid external reference",
			external: "projects/my-project/locations/us-central1/repositories/my-repo",
			want: &DataformRepositoryIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Repository: "my-repo",
			},
			wantErr: false,
		},
		{
			name:     "invalid format",
			external: "projects/my-project/repositories/my-repo",
			wantErr:  true,
		},
		{
			name:     "another invalid format",
			external: "locations/us-central1/repositories/my-repo",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DataformRepositoryIdentity{}
			if err := i.FromExternal(tt.external); (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project {
					t.Errorf("Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Location != tt.want.Location {
					t.Errorf("Location = %v, want %v", i.Location, tt.want.Location)
				}
				if i.Repository != tt.want.Repository {
					t.Errorf("Repository = %v, want %v", i.Repository, tt.want.Repository)
				}
			}
		})
	}
}

func TestDataformRepositoryIdentity_String(t *testing.T) {
	identity := &DataformRepositoryIdentity{
		Project:    "my-project",
		Location:   "us-central1",
		Repository: "my-repo",
	}
	want := "projects/my-project/locations/us-central1/repositories/my-repo"
	if got := identity.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}
