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

func TestFirebaseHostingChannelIdentity(t *testing.T) {
	tests := []struct {
		name        string
		ref         string
		expected    *FirebaseHostingChannelIdentity
		expectedErr bool
	}{
		{
			name: "valid relative reference",
			ref:  "sites/test-site/channels/test-channel",
			expected: &FirebaseHostingChannelIdentity{
				Site:    "test-site",
				Channel: "test-channel",
			},
		},
		{
			name: "valid external reference with host",
			ref:  "//firebasehosting.googleapis.com/sites/test-site/channels/test-channel",
			expected: &FirebaseHostingChannelIdentity{
				Site:    "test-site",
				Channel: "test-channel",
			},
		},
		{
			name: "valid external reference with https",
			ref:  "https://firebasehosting.googleapis.com/sites/test-site/channels/test-channel",
			expected: &FirebaseHostingChannelIdentity{
				Site:    "test-site",
				Channel: "test-channel",
			},
		},
		{
			name:        "invalid format",
			ref:         "sites/test-site",
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &FirebaseHostingChannelIdentity{}
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
			if id.Site != tc.expected.Site || id.Channel != tc.expected.Channel {
				t.Fatalf("expected %+v, got %+v", tc.expected, id)
			}
		})
	}
}
