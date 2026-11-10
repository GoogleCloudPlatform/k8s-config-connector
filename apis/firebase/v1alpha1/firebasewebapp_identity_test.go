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

func TestFirebaseWebAppIdentity(t *testing.T) {
	tests := []struct {
		name        string
		ref         string
		expected    *FirebaseWebAppIdentity
		expectedErr bool
	}{
		{
			name: "valid external reference",
			ref:  "projects/test-project/webApps/1:123456789:web:abcdef",
			expected: &FirebaseWebAppIdentity{
				Project: "test-project",
				App:     "1:123456789:web:abcdef",
			},
			expectedErr: false,
		},
		{
			name:        "invalid external reference",
			ref:         "projects/test-project/invalid/123",
			expected:    nil,
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &FirebaseWebAppIdentity{}
			err := id.FromExternal(tc.ref)
			if tc.expectedErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if id.Project != tc.expected.Project || id.App != tc.expected.App {
				t.Fatalf("got %v, want %v", id, tc.expected)
			}
			if got := id.String(); got != tc.ref {
				t.Fatalf("id.String() got %q, want %q", got, tc.ref)
			}
		})
	}
}
