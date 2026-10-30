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

package fixture

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadAndDiscoverDirs(t *testing.T) {
	tests := []struct {
		name             string
		files            map[string]string
		wantLoadErr      bool
		wantDiscovered   int
		wantPrimaryKind  string
		wantPrimaryResID string
		wantHasUpdate    bool
		wantDepCount     int
		wantNonDeletable bool
	}{
		{
			name: "single fixture with create.yaml, update.yaml, dependencies.yaml, and golden file",
			files: map[string]string{
				"create.yaml": `apiVersion: widgets.cnrm.cloud.google.com/v1alpha1
kind: ExampleWidget
metadata:
  name: widget-${uniqueId}
spec:
  resourceID: custom-id-${uniqueId}
  description: initial
`,
				"update.yaml": `apiVersion: widgets.cnrm.cloud.google.com/v1alpha1
kind: ExampleWidget
metadata:
  name: widget-${uniqueId}
  annotations:
    cnrm.cloud.google.com/deletion-policy: abandon
spec:
  resourceID: custom-id-${uniqueId}
  description: updated
`,
				"dependencies.yaml": `apiVersion: resourcemanager.cnrm.cloud.google.com/v1beta1
kind: Project
metadata:
  name: dep-project
`,
				"_http.log": "# Phase: CREATE\nGET https://example.googleapis.com/v1/w\n\n200 OK\n",
			},
			wantLoadErr:      false,
			wantDiscovered:   1,
			wantPrimaryKind:  "ExampleWidget",
			wantPrimaryResID: "custom-id-${uniqueId}",
			wantHasUpdate:    true,
			wantDepCount:     1,
			wantNonDeletable: true,
		},
		{
			name: "fixture with only create.yaml",
			files: map[string]string{
				"create.yaml": `apiVersion: widgets.cnrm.cloud.google.com/v1alpha1
kind: ExampleWidget
metadata:
  name: widget-only
spec:
  description: initial
`,
				"_http.log": "# Phase: CREATE\nGET https://example.googleapis.com/v1/w\n\n200 OK\n",
			},
			wantLoadErr:      false,
			wantDiscovered:   1,
			wantPrimaryKind:  "ExampleWidget",
			wantPrimaryResID: "",
			wantHasUpdate:    false,
			wantDepCount:     0,
			wantNonDeletable: false,
		},
		{
			name:           "invalid or missing fixture error",
			files:          map[string]string{},
			wantLoadErr:    true,
			wantDiscovered: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			root := t.TempDir()
			fixDir := filepath.Join(root, "widgets", "v1alpha1", "examplewidget", "widget-case")
			if err := os.MkdirAll(fixDir, 0755); err != nil {
				t.Fatalf("MkdirAll: %v", err)
			}
			for fname, content := range tc.files {
				if err := os.WriteFile(filepath.Join(fixDir, fname), []byte(content), 0644); err != nil {
					t.Fatalf("WriteFile(%s): %v", fname, err)
				}
			}

			// Act
			dirs, discoverErr := DiscoverDirs(root)
			data, loadErr := Load(fixDir)

			// Assert
			if discoverErr != nil {
				t.Fatalf("DiscoverDirs() returned unexpected error: %v", discoverErr)
			}
			if len(dirs) != tc.wantDiscovered {
				t.Errorf("len(DiscoverDirs()) = %d, want %d (%v)", len(dirs), tc.wantDiscovered, dirs)
			}
			if (loadErr != nil) != tc.wantLoadErr {
				t.Fatalf("Load() error = %v, wantLoadErr %v", loadErr, tc.wantLoadErr)
			}
			if tc.wantLoadErr {
				return
			}
			if data.PrimaryCreate == nil || data.PrimaryCreate.Kind != tc.wantPrimaryKind || data.PrimaryCreate.ResourceID != tc.wantPrimaryResID {
				t.Errorf("PrimaryCreate = %+v, want Kind=%q ResourceID=%q", data.PrimaryCreate, tc.wantPrimaryKind, tc.wantPrimaryResID)
			}
			if (data.PrimaryUpdate != nil) != tc.wantHasUpdate {
				t.Errorf("PrimaryUpdate != nil is %v, want %v", data.PrimaryUpdate != nil, tc.wantHasUpdate)
			}
			if len(data.Dependencies) != tc.wantDepCount {
				t.Errorf("len(Dependencies) = %d, want %d", len(data.Dependencies), tc.wantDepCount)
			}
			if data.IsNonDeletable() != tc.wantNonDeletable {
				t.Errorf("IsNonDeletable() = %v, want %v", data.IsNonDeletable(), tc.wantNonDeletable)
			}
		})
	}
}
