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

package networkservices

import "testing"

func TestNormalizeGCPResourceName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "my-bucket",
			expected: "my-bucket",
		},
		{
			input:    "gs://my-bucket",
			expected: "my-bucket",
		},
		{
			input:    "https://storage.googleapis.com/storage/v1/b/my-bucket",
			expected: "my-bucket",
		},
		{
			input:    "https://www.googleapis.com/storage/v1/b/my-bucket",
			expected: "my-bucket",
		},
		{
			input:    "https://storage.googleapis.com/my-bucket",
			expected: "my-bucket",
		},
		{
			input:    "https://compute.googleapis.com/compute/v1/projects/my-project/global/backendServices/my-backend",
			expected: "projects/my-project/global/backendServices/my-backend",
		},
		{
			input:    "projects/my-project/global/backendServices/my-backend",
			expected: "projects/my-project/global/backendServices/my-backend",
		},
	}

	for _, tc := range tests {
		actual := normalizeGCPResourceName(tc.input)
		if actual != tc.expected {
			t.Errorf("normalizeGCPResourceName(%q) = %q, want %q", tc.input, actual, tc.expected)
		}
	}
}
