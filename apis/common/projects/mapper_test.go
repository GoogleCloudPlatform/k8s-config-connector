// Copyright 2025 Google LLC
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
	"context"
	"testing"
)

func TestReplaceProjectNumberWithIDInLink(t *testing.T) {
	ctx := context.Background()

	mockResolver := func(ctx context.Context, idOrNumber string) (string, error) {
		if idOrNumber == "12345" {
			return "my-project", nil
		}
		if idOrNumber == "67890" {
			return "other-project", nil
		}
		// Return as is if not found (simulating ReplaceProjectNumberWithID behavior for non-numbers or unknown numbers)
		return idOrNumber, nil
	}

	tests := []struct {
		name     string
		link     string
		expected string
	}{
		{
			name:     "simple project number",
			link:     "//artifactregistry.googleapis.com/projects/12345/locations/us-central1/repositories/bar",
			expected: "//artifactregistry.googleapis.com/projects/my-project/locations/us-central1/repositories/bar",
		},
		{
			name:     "already project id",
			link:     "//artifactregistry.googleapis.com/projects/my-project/locations/us-central1/repositories/bar",
			expected: "//artifactregistry.googleapis.com/projects/my-project/locations/us-central1/repositories/bar",
		},
		{
			name:     "unknown project number",
			link:     "//artifactregistry.googleapis.com/projects/99999/locations/us-central1/repositories/bar",
			expected: "//artifactregistry.googleapis.com/projects/99999/locations/us-central1/repositories/bar",
		},
		{
			name:     "multiple projects",
			link:     "projects/12345/regions/us-central1/subnetworks/default/projects/67890", // Hypothetical
			expected: "projects/my-project/regions/us-central1/subnetworks/default/projects/other-project",
		},
		{
			name:     "no projects segment",
			link:     "folders/12345",
			expected: "folders/12345",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := replaceProjectNumberWithIDInLink(ctx, tc.link, mockResolver)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}