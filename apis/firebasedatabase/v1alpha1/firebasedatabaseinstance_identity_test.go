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

func TestFirebaseDatabaseInstanceIdentity(t *testing.T) {
	tests := []struct {
		name        string
		ref         string
		expected    *FirebaseDatabaseInstanceIdentity
		expectedErr bool
	}{
		{
			name: "valid external reference",
			ref:  "projects/test-project/locations/us-central1/instances/test-instance",
			expected: &FirebaseDatabaseInstanceIdentity{
				Project:  "test-project",
				Location: "us-central1",
				Instance: "test-instance",
			},
			expectedErr: false,
		},
		{
			name: "valid full URL",
			ref:  "//firebasedatabase.googleapis.com/projects/test-project/locations/us-central1/instances/test-instance",
			expected: &FirebaseDatabaseInstanceIdentity{
				Project:  "test-project",
				Location: "us-central1",
				Instance: "test-instance",
			},
			expectedErr: false,
		},
		{
			name:        "invalid external reference",
			ref:         "projects/test-project/instances/test-instance",
			expected:    nil,
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &FirebaseDatabaseInstanceIdentity{}
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
			if diff := cmp.Diff(tc.expected, id); diff != "" {
				t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
			}
			if got := id.String(); got != tc.expected.String() {
				t.Fatalf("id.String() got %q, want %q", got, tc.expected.String())
			}
		})
	}
}

func TestFirebaseDatabaseInstanceRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name: "valid relative external ref",
			ref:  "projects/test-project/locations/us-central1/instances/test-instance",
		},
		{
			name: "valid full external ref",
			ref:  "//firebasedatabase.googleapis.com/projects/test-project/locations/us-central1/instances/test-instance",
		},
		{
			name:    "invalid external ref",
			ref:     "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &FirebaseDatabaseInstanceRef{
				External: tt.ref,
			}
			err := r.ValidateExternal(r.External)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
