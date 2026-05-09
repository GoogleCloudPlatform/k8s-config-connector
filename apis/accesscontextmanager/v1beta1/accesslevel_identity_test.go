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

package v1beta1

import (
	"testing"
)

func TestAccessLevelParse(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expectedID   *AccessContextManagerAccessLevelIdentity
		hasError     bool
	}{
		{
			name:       "Normal parse",
			input:      "accessPolicies/val1/accessLevels/val2",
			expectedID: &AccessContextManagerAccessLevelIdentity{AccessPolicy: "val1", AccessLevel: "val2"},
			hasError:   false,
		},
		{
			name:       "Normal parse with leading slash",
			input:      "/accessPolicies/foo/accessLevels/bar",
			expectedID: &AccessContextManagerAccessLevelIdentity{AccessPolicy: "foo", AccessLevel: "bar"},
			hasError:   false,
		},
		{
			name:       "Normal parse with domain",
			input:      "accesscontextmanager.googleapis.com/accessPolicies/policy/accessLevels/level",
			expectedID: &AccessContextManagerAccessLevelIdentity{AccessPolicy: "policy", AccessLevel: "level"},
			hasError:   false,
		},
		{
			name:       "Normal parse with slashed domain",
			input:      "//accesscontextmanager.googleapis.com/accessPolicies/policy/accessLevels/level",
			expectedID: &AccessContextManagerAccessLevelIdentity{AccessPolicy: "policy", AccessLevel: "level"},
			hasError:   false,
		},
		{
			name:       "Normal parse with wrong domain",
			input:      "iam.googleapis.com/accessPolicies/policy/accessLevels/level",
			expectedID: nil,
			hasError:   true,
		},
		{
			name:       "Normal parse with wrong key",
			input:      "accessPolicys/policy/accessLevels/level",
			expectedID: nil,
			hasError:   true,
		},
	}

	for _, tc := range tests {
		id := &AccessContextManagerAccessLevelIdentity{}
		err := id.FromExternal(tc.input)
		if tc.hasError {
			if err == nil {
				t.Fatalf("Test %s expected error but did not get one", tc.name)
			}
			continue
		}
		// Error no expected at this point
		if err != nil {
			t.Fatalf("Test %s did not expect error but got %v", tc.name, err)
		}
		if *id != *tc.expectedID {
			t.Fatalf("Test %s bad result %v != %v", tc.name, *id, *tc.expectedID)
		}
	}
}
