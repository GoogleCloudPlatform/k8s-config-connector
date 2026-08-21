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

	"github.com/google/go-cmp/cmp"
)

func TestConfigDeploymentGroupIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *ConfigDeploymentGroupIdentity
		hasError bool
	}{
		{
			name:  "Full resource name",
			input: "projects/my-project/locations/global/deploymentGroups/my-group",
			expected: &ConfigDeploymentGroupIdentity{
				Project:         "my-project",
				Location:        "global",
				DeploymentGroup: "my-group",
			},
			hasError: false,
		},
		{
			name:  "Full resource name with host",
			input: "config.googleapis.com/projects/my-project/locations/global/deploymentGroups/my-group",
			expected: &ConfigDeploymentGroupIdentity{
				Project:         "my-project",
				Location:        "global",
				DeploymentGroup: "my-group",
			},
			hasError: false,
		},
		{
			name:     "Invalid format",
			input:    "projects/my-project/locations/global/invalid/my-group",
			expected: nil,
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &ConfigDeploymentGroupIdentity{}
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
			if diff := cmp.Diff(tc.expected, id); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
