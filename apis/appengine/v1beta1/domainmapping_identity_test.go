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

func TestParseDomainMappingExternal(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectProject string
		expectDomain  string
		expectError   bool
	}{
		{
			name:          "Normal parse",
			input:         "apps/myProject/domainMappings/example.com",
			expectProject: "myProject",
			expectDomain:  "example.com",
			expectError:   false,
		},
		{
			name:          "Parse with googleapis prefix",
			input:         "//appengine.googleapis.com/apps/myProject/domainMappings/example.com",
			expectProject: "myProject",
			expectDomain:  "example.com",
			expectError:   false,
		},
		{
			name:          "Parse with subdomain",
			input:         "apps/test-project/domainMappings/www.example.com",
			expectProject: "test-project",
			expectDomain:  "www.example.com",
			expectError:   false,
		},
		{
			name:        "Invalid format - missing apps prefix",
			input:       "projects/myProject/domainMappings/example.com",
			expectError: true,
		},
		{
			name:        "Invalid format - wrong resource type",
			input:       "apps/myProject/domains/example.com",
			expectError: true,
		},
		{
			name:        "Invalid format - too few segments",
			input:       "apps/myProject/domainMappings",
			expectError: true,
		},
		{
			name:        "Invalid format - too many segments",
			input:       "apps/myProject/domainMappings/example.com/extra",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parent, resourceID, err := ParseDomainMappingExternal(tc.input)
			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if parent.ProjectID != tc.expectProject {
				t.Errorf("expected project %q, got %q", tc.expectProject, parent.ProjectID)
			}
			if resourceID != tc.expectDomain {
				t.Errorf("expected domain %q, got %q", tc.expectDomain, resourceID)
			}
		})
	}
}
