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

func TestParseAgentRegistryBindingExternal(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		expectedParent   *AgentRegistryBindingParent
		expectedResource string
		hasError         bool
	}{
		{
			name:  "Valid standard external reference",
			input: "projects/my-project/locations/us-central1/bindings/my-binding",
			expectedParent: &AgentRegistryBindingParent{
				ProjectID: "my-project",
				Location:  "us-central1",
			},
			expectedResource: "my-binding",
			hasError:         false,
		},
		{
			name:             "Invalid format - too short",
			input:            "projects/my-project/locations/us-central1/bindings",
			expectedParent:   nil,
			expectedResource: "",
			hasError:         true,
		},
		{
			name:             "Invalid format - wrong keyword",
			input:            "projects/my-project/regions/us-central1/bindings/my-binding",
			expectedParent:   nil,
			expectedResource: "",
			hasError:         true,
		},
		{
			name:             "Invalid format - empty string",
			input:            "",
			expectedParent:   nil,
			expectedResource: "",
			hasError:         true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parent, resourceID, err := ParseAgentRegistryBindingExternal(tc.input)
			if tc.hasError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if parent.ProjectID != tc.expectedParent.ProjectID || parent.Location != tc.expectedParent.Location {
				t.Fatalf("expected parent %+v, got %+v", tc.expectedParent, parent)
			}
			if resourceID != tc.expectedResource {
				t.Fatalf("expected resourceID %q, got %q", tc.expectedResource, resourceID)
			}

			// test identity String format
			id := &AgentRegistryBindingIdentity{
				parent: parent,
				id:     resourceID,
			}
			if id.String() != tc.input {
				t.Fatalf("expected string %q, got %q", tc.input, id.String())
			}
			if id.ID() != resourceID {
				t.Fatalf("expected id %q, got %q", resourceID, id.ID())
			}
			if id.Parent().ProjectID != parent.ProjectID || id.Parent().Location != parent.Location {
				t.Fatalf("expected id parent %+v, got %+v", parent, id.Parent())
			}
		})
	}
}
