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

package common

import (
	"fmt"
	"testing"
)

// TestFixStaleComputeExternalFormat tests the possible external formats
// that the function is expected to handle.
func TestFixStaleComputeExternalFormat(t *testing.T) {
	const relativePath = "projects/projectId/location/us/resources/resourceId"

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "external is the deprecated format",
			input:    fmt.Sprintf("/compute.googleapis.com/%s", relativePath),
			expected: relativePath,
		},
		{
			name:     "external is full url(v1)",
			input:    fmt.Sprintf("https://www.googleapis.com/compute/v1/%s", relativePath),
			expected: relativePath,
		},
		{
			name:     "external is full url(v1beta1)",
			input:    fmt.Sprintf("https://www.googleapis.com/compute/v1beta1/%s", relativePath),
			expected: relativePath,
		},
		{
			name:     "external is full url(beta)",
			input:    fmt.Sprintf("https://www.googleapis.com/compute/beta/%s", relativePath),
			expected: relativePath,
		},
		{
			// Test passes with warning:
			// WARNING: received Compute selfLink with unknown version otherVersion, accepted versions are v1, v1beta1 and beta.
			// The output looks like `otherVersion/projects/projectId/location/us/resources/resourceId` and will fail the compute external format validation later.
			name:     "external is full url(otherVersion)",
			input:    fmt.Sprintf("https://www.googleapis.com/compute/otherVersion/%s", relativePath),
			expected: fmt.Sprintf("otherVersion/%s", relativePath),
		},
		{
			// Test passes.
			// The output looks like `https://www.googleapis.com/storage/v1/bucket` and will fail the compute external format validation later.
			name:     "external is full url but not compute",
			input:    "https://www.googleapis.com/storage/v1/bucket",
			expected: "https://www.googleapis.com/storage/v1/bucket",
		},
		{
			name:     "external is relative path with leading slash",
			input:    fmt.Sprintf("/%s", relativePath),
			expected: relativePath,
		},
		{
			name:     "external is relative path without leading slash",
			input:    relativePath,
			expected: relativePath,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FixStaleComputeExternalFormat(tt.input)

			if got != tt.expected {
				t.Errorf("got: %q; want: %q", got, tt.expected)
			}
		})
	}
}
