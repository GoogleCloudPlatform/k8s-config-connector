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

package projects

import (
	"testing"
)

func TestReplaceProjectWithProjectNumberTemplate(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:     "relative path with trailing suffix",
			input:    "projects/my-project/zones/z",
			expected: "projects/${projectNumber}/zones/z",
		},
		{
			name:     "relative path with number ID",
			input:    "projects/123456789/zones/z",
			expected: "projects/${projectNumber}/zones/z",
		},
		{
			name:     "relative path with no trailing suffix",
			input:    "projects/my-project",
			expected: "projects/${projectNumber}",
		},
		{
			name:     "full URL path",
			input:    "https://container.googleapis.com/v1/projects/my-project/zones/us-central1-a/clusters/cluster-1",
			expected: "https://container.googleapis.com/v1/projects/${projectNumber}/zones/us-central1-a/clusters/cluster-1",
		},
		{
			name:        "no project prefix",
			input:       "no-project-prefix",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ReplaceProjectWithProjectNumberTemplate(tc.input)
			if tc.expectError {
				if err == nil {
					t.Errorf("expected error, got none")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.expected {
				t.Errorf("ReplaceProjectWithProjectNumberTemplate(%q) = %q; want %q", tc.input, got, tc.expected)
			}
		})
	}
}
