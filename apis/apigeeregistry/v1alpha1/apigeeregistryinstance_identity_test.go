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

package v1alpha1

import (
	"testing"
)

func TestInstanceIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *InstanceIdentity
		hasError bool
	}{
		{
			name:  "Full resource name segments",
			input: "projects/my-project/locations/us-central1/instances/my-instance",
			expected: &InstanceIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Instance: "my-instance",
			},
			hasError: false,
		},
		{
			name:  "Full resource name with host segment",
			input: "apigeeregistry.googleapis.com/projects/my-project/locations/us-central1/instances/my-instance",
			expected: &InstanceIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Instance: "my-instance",
			},
			hasError: false,
		},
		{
			name:     "Invalid resource segments",
			input:    "projects/my-project/locations/us-central1/invalid/my-instance",
			expected: nil,
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &InstanceIdentity{}
			err := id.FromExternal(tc.input)
			if tc.hasError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if id.Project != tc.expected.Project || id.Location != tc.expected.Location || id.Instance != tc.expected.Instance {
				t.Fatalf("expected %+v, got %+v", tc.expected, id)
			}
		})
	}
}
