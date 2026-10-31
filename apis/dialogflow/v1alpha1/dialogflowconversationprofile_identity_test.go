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

func TestDialogflowConversationProfileIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name       string
		ref        string
		wantErr    bool
		want       *DialogflowConversationProfileIdentity
		wantString string
		wantParent string
	}{
		{
			name: "valid regional reference",
			ref:  "projects/my-project/locations/us-central1/conversationProfiles/my-profile",
			want: &DialogflowConversationProfileIdentity{
				Project:             "my-project",
				Location:            "us-central1",
				ConversationProfile: "my-profile",
			},
			wantString: "projects/my-project/locations/us-central1/conversationProfiles/my-profile",
			wantParent: "projects/my-project/locations/us-central1",
		},
		{
			name: "valid global reference",
			ref:  "projects/my-project/conversationProfiles/my-profile",
			want: &DialogflowConversationProfileIdentity{
				Project:             "my-project",
				Location:            "",
				ConversationProfile: "my-profile",
			},
			wantString: "projects/my-project/conversationProfiles/my-profile",
			wantParent: "projects/my-project",
		},
		{
			name: "full regional url reference",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/locations/us-central1/conversationProfiles/my-profile",
			want: &DialogflowConversationProfileIdentity{
				Project:             "my-project",
				Location:            "us-central1",
				ConversationProfile: "my-profile",
			},
			wantString: "projects/my-project/locations/us-central1/conversationProfiles/my-profile",
			wantParent: "projects/my-project/locations/us-central1",
		},
		{
			name: "full global url reference",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/conversationProfiles/my-profile",
			want: &DialogflowConversationProfileIdentity{
				Project:             "my-project",
				Location:            "",
				ConversationProfile: "my-profile",
			},
			wantString: "projects/my-project/conversationProfiles/my-profile",
			wantParent: "projects/my-project",
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "wrong resource type",
			ref:     "projects/my-project/locations/us-central1/knowledgeBases/my-kb",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DialogflowConversationProfileIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal(%q) error = %v, wantErr %v", tt.ref, err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal(%q) mismatch (-want +got):\n%s", tt.ref, diff)
				}
				if gotStr := i.String(); gotStr != tt.wantString {
					t.Errorf("String() = %q, want %q", gotStr, tt.wantString)
				}
				if gotHost := i.Host(); gotHost != "dialogflow.googleapis.com" {
					t.Errorf("Host() = %q, want %q", gotHost, "dialogflow.googleapis.com")
				}
				if gotParent := i.ParentString(); gotParent != tt.wantParent {
					t.Errorf("ParentString() = %q, want %q", gotParent, tt.wantParent)
				}
			}
		})
	}
}

func TestDialogflowConversationProfileRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name: "valid regional reference",
			ref:  "projects/my-project/locations/us-central1/conversationProfiles/my-profile",
		},
		{
			name: "valid global reference",
			ref:  "projects/my-project/conversationProfiles/my-profile",
		},
		{
			name: "full regional url",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/locations/us-central1/conversationProfiles/my-profile",
		},
		{
			name: "full global url",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/conversationProfiles/my-profile",
		},
		{
			name:    "invalid reference",
			ref:     "invalid/reference",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DialogflowConversationProfileRef{
				External: tt.ref,
			}
			err := r.ValidateExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal(%q) error = %v, wantErr %v", tt.ref, err, tt.wantErr)
			}
		})
	}
}
