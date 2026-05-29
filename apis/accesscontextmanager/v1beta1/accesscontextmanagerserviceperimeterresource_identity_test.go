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

func TestAccessContextManagerServicePerimeterResourceIdentityParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *AccessContextManagerServicePerimeterResourceIdentity
		hasError bool
	}{
		{
			name:  "Normal parse",
			input: "accessPolicies/123/servicePerimeters/my_perimeter/projects/456",
			expected: &AccessContextManagerServicePerimeterResourceIdentity{
				AccessPolicy:     "123",
				ServicePerimeter: "my_perimeter",
				Project:          "456",
			},
			hasError: false,
		},
		{
			name:  "Parsed with host",
			input: "accesscontextmanager.googleapis.com/accessPolicies/123/servicePerimeters/my_perimeter/projects/456",
			expected: &AccessContextManagerServicePerimeterResourceIdentity{
				AccessPolicy:     "123",
				ServicePerimeter: "my_perimeter",
				Project:          "456",
			},
			hasError: false,
		},
		{
			name:  "Parsed with leading slash",
			input: "/accessPolicies/123/servicePerimeters/my_perimeter/projects/456",
			expected: &AccessContextManagerServicePerimeterResourceIdentity{
				AccessPolicy:     "123",
				ServicePerimeter: "my_perimeter",
				Project:          "456",
			},
			hasError: false,
		},
		{
			name:     "Invalid format - missing projects",
			input:    "accessPolicies/123/servicePerimeters/my_perimeter",
			expected: nil,
			hasError: true,
		},
		{
			name:     "Invalid format - wrong resource type",
			input:    "accessPolicies/123/servicePerimeters/my_perimeter/folders/456",
			expected: nil,
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &AccessContextManagerServicePerimeterResourceIdentity{}
			err := id.FromExternal(tc.input)
			if tc.hasError {
				if err == nil {
					t.Fatalf("expected error, got none")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if id.AccessPolicy != tc.expected.AccessPolicy ||
				id.ServicePerimeter != tc.expected.ServicePerimeter ||
				id.Project != tc.expected.Project {
				t.Fatalf("expected %+v, got %+v", tc.expected, id)
			}
			if id.String() != "accessPolicies/"+tc.expected.AccessPolicy+"/servicePerimeters/"+tc.expected.ServicePerimeter+"/projects/"+tc.expected.Project {
				t.Fatalf("bad String(): %s", id.String())
			}
		})
	}
}
