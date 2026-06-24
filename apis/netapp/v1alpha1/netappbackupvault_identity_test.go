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

func TestParseBackupVaultExternal(t *testing.T) {
	tests := []struct {
		name        string
		external    string
		wantErr     bool
		wantProj    string
		wantLoc     string
		wantVaultID string
	}{
		{
			name:        "valid external with lowercase backupvaults",
			external:    "projects/my-project/locations/us-central1/backupvaults/my-vault",
			wantProj:    "my-project",
			wantLoc:     "us-central1",
			wantVaultID: "my-vault",
		},
		{
			name:        "valid external with camelcase backupVaults",
			external:    "projects/my-project/locations/us-central1/backupVaults/my-vault",
			wantProj:    "my-project",
			wantLoc:     "us-central1",
			wantVaultID: "my-vault",
		},
		{
			name:     "invalid external - short",
			external: "projects/my-project/locations/us-central1",
			wantErr:  true,
		},
		{
			name:     "invalid external - bad prefix",
			external: "gcp/my-project/locations/us-central1/backupvaults/my-vault",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parent, id, err := ParseBackupVaultExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBackupVaultExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if parent.ProjectID != tt.wantProj {
					t.Errorf("ProjectID = %v, want %v", parent.ProjectID, tt.wantProj)
				}
				if parent.Location != tt.wantLoc {
					t.Errorf("Location = %v, want %v", parent.Location, tt.wantLoc)
				}
				if id != tt.wantVaultID {
					t.Errorf("id = %v, want %v", id, tt.wantVaultID)
				}
			}
		})
	}
}
