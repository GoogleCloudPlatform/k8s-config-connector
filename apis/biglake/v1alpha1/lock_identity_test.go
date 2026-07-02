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

func TestBigLakeLockIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *BigLakeLockIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/catalogs/my-catalog/databases/my-database/locks/my-lock",
			want: &BigLakeLockIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Catalog:  "my-catalog",
				Database: "my-database",
				Lock:     "my-lock",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://biglake.googleapis.com/projects/my-project/locations/us-central1/catalogs/my-catalog/databases/my-database/locks/my-lock",
			want: &BigLakeLockIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Catalog:  "my-catalog",
				Database: "my-database",
				Lock:     "my-lock",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BigLakeLockIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
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
				if i.Catalog != tt.want.Catalog {
					t.Errorf("Catalog = %v, want %v", i.Catalog, tt.want.Catalog)
				}
				if i.Database != tt.want.Database {
					t.Errorf("Database = %v, want %v", i.Database, tt.want.Database)
				}
				if i.Lock != tt.want.Lock {
					t.Errorf("Lock = %v, want %v", i.Lock, tt.want.Lock)
				}
			}
		})
	}
}
