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

import "testing"

func TestSQLInstanceIdentityFromExternal(t *testing.T) {
	tests := []struct {
		name         string
		external     string
		wantProject  string
		wantInstance string
		wantErr      bool
	}{
		{
			name:         "full URI",
			external:     "//sqladmin.googleapis.com/projects/my-project/instances/my-instance",
			wantProject:  "my-project",
			wantInstance: "my-instance",
		},
		{
			name:         "relative resource name",
			external:     "projects/my-project/instances/my-instance",
			wantProject:  "my-project",
			wantInstance: "my-instance",
		},
		{
			name:     "invalid location-qualified resource name",
			external: "projects/my-project/locations/us-central1/instances/my-instance",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := &SQLInstanceIdentity{}
			err := identity.FromExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if identity.Project != tt.wantProject {
				t.Errorf("FromExternal() project = %q, want %q", identity.Project, tt.wantProject)
			}
			if identity.Instance != tt.wantInstance {
				t.Errorf("FromExternal() instance = %q, want %q", identity.Instance, tt.wantInstance)
			}
		})
	}
}

func TestSQLInstanceIdentityString(t *testing.T) {
	identity := &SQLInstanceIdentity{
		Project:  "my-project",
		Instance: "my-instance",
	}

	want := "projects/my-project/instances/my-instance"
	if got := identity.String(); got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}
