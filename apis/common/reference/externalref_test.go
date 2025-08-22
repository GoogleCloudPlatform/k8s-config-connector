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

package reference

import (
	"reflect"
	"testing"
)

func TestParseExternalRef(t *testing.T) {
	testCases := []struct {
		name          string
		externalRef   string
		expectedMap   map[string]string
		expectError   bool
		errorContains string
	}{
		{
			name:        "standard externalRef",
			externalRef: "projects/test-project/locations/us-central1/instances/test-instance",
			expectedMap: map[string]string{
				"externalRef": "projects/test-project/locations/us-central1/instances/test-instance",
				"project":     "test-project",
				"location":    "us-central1",
				"instance":    "test-instance",
				"name":        "test-instance",
			},
			expectError: false,
		},
		{
			name:        "externalRef with pluralization(ies)",
			externalRef: "projects/test-project/locations/us-west1/registries/test-registry",
			expectedMap: map[string]string{
				"externalRef": "projects/test-project/locations/us-west1/registries/test-registry",
				"project":     "test-project",
				"location":    "us-west1",
				"registry":    "test-registry",
				"name":        "test-registry",
			},
			expectError: false,
		},
		{
			name:        "externalRef with pluralization(ves)",
			externalRef: "projects/test-project/locations/us-west1/bookshelves/test-bookshelf",
			expectedMap: map[string]string{
				"externalRef": "projects/test-project/locations/us-west1/bookshelves/test-bookshelf",
				"project":     "test-project",
				"location":    "us-west1",
				"bookshelf":   "test-bookshelf",
				"name":        "test-bookshelf",
			},
			expectError: false,
		},
		{
			name:        "name only",
			externalRef: "test-name",
			expectedMap: map[string]string{
				"externalRef": "test-name",
				"name":        "test-name",
			},
			expectError: false,
		},
		{
			name:        "externalRef with global",
			externalRef: "projects/test-project/global/instances/test-instance",
			expectedMap: map[string]string{
				"externalRef": "projects/test-project/global/instances/test-instance",
				"name":        "test-instance",
			},
			expectError: false,
		},
		{
			name:          "empty externalRef",
			externalRef:   "",
			expectedMap:   nil,
			expectError:   true,
			errorContains: "externalRef is empty",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			components, err := ParseExternalRef(tc.externalRef)

			if tc.expectError {
				if err == nil {
					t.Fatalf("expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if !reflect.DeepEqual(components, tc.expectedMap) {
					t.Errorf("map mismatch: got %v, want %v", components, tc.expectedMap)
				}
			}
		})
	}
}
