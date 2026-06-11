// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alpha

import "testing"

func TestIsRegion(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// Standard regions
		{"us-central1", true},
		{"europe-west3", true},
		{"asia-east1", true},
		{"us-east4", true},

		// DLP Multi-regions
		{"us", true},
		{"eu", true},
		{"in", true},
		{"europe", false},
		{"asia", false},

		// Global / Non-regional
		{"global", false},

		// Invalid regions / formats
		{"us-central", false},
		{"europe-west", false},
		{"", false},
		{"us-central1-a", false}, // This is a zone
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			actual := isRegion(&tc.input)
			if actual != tc.expected {
				t.Errorf("isRegion(%q) = %v; want %v", tc.input, actual, tc.expected)
			}
		})
	}
}
