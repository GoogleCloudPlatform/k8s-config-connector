// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package identity

import "testing"

func TestStripReferencePrefixes(t *testing.T) {
	tests := []struct {
		name     string
		ref      string
		host     string
		expected string
	}{
		{
			name:     "raw projects reference",
			ref:      "projects/my-project/locations/us-central1/clusters/my-cluster",
			host:     "container.googleapis.com",
			expected: "projects/my-project/locations/us-central1/clusters/my-cluster",
		},
		{
			name:     "https scheme and host and v1",
			ref:      "https://container.googleapis.com/v1/projects/my-project/locations/us-central1/clusters/my-cluster",
			host:     "container.googleapis.com",
			expected: "projects/my-project/locations/us-central1/clusters/my-cluster",
		},
		{
			name:     "https scheme and host and v1beta1",
			ref:      "https://container.googleapis.com/v1beta1/projects/my-project/locations/us-central1/clusters/my-cluster",
			host:     "container.googleapis.com",
			expected: "projects/my-project/locations/us-central1/clusters/my-cluster",
		},
		{
			name:     "http scheme and host",
			ref:      "http://container.googleapis.com/projects/my-project/locations/us-central1/clusters/my-cluster",
			host:     "container.googleapis.com",
			expected: "projects/my-project/locations/us-central1/clusters/my-cluster",
		},
		{
			name:     "double slash prefix and host",
			ref:      "//container.googleapis.com/projects/my-project/locations/us-central1/clusters/my-cluster",
			host:     "container.googleapis.com",
			expected: "projects/my-project/locations/us-central1/clusters/my-cluster",
		},
		{
			name:     "unrelated host does not match",
			ref:      "https://other.googleapis.com/v1/projects/my-project/locations/us-central1/clusters/my-cluster",
			host:     "container.googleapis.com",
			expected: "other.googleapis.com/v1/projects/my-project/locations/us-central1/clusters/my-cluster",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StripReferencePrefixes(tt.ref, tt.host)
			if got != tt.expected {
				t.Errorf("StripReferencePrefixes() = %q, expected %q", got, tt.expected)
			}
		})
	}
}
