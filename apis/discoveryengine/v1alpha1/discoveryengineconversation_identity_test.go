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

func TestDiscoveryEngineConversationIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DiscoveryEngineConversationIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/conversations/my-conversation",
			want: &DiscoveryEngineConversationIdentity{
				Project:      "my-project",
				Location:     "global",
				Collection:   "default_collection",
				DataStore:    "my-datastore",
				Conversation: "my-conversation",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://discoveryengine.googleapis.com/projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/conversations/my-conversation",
			want: &DiscoveryEngineConversationIdentity{
				Project:      "my-project",
				Location:     "global",
				Collection:   "default_collection",
				DataStore:    "my-datastore",
				Conversation: "my-conversation",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DiscoveryEngineConversationIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestIsProjectIDMatch(t *testing.T) {
	tests := []struct {
		p1, p2 string
		want   bool
	}{
		{"1234567890", "1234567890", true},
		{"my-project", "my-project", true},
		{"my-project", "1234567890", true},     // alphanumeric project ID and numeric project number match
		{"1234567890", "my-project", true},     // numeric project number and alphanumeric project ID match
		{"1234567890", "9876543210", false},    // both numeric but different
		{"my-project", "other-project", false}, // both alphanumeric but different
		{"", "my-project", false},
		{"my-project", "", false},
	}

	for _, tt := range tests {
		got := IsProjectIDMatch(tt.p1, tt.p2)
		if got != tt.want {
			t.Errorf("IsProjectIDMatch(%q, %q) = %v; want %v", tt.p1, tt.p2, got, tt.want)
		}
	}
}
