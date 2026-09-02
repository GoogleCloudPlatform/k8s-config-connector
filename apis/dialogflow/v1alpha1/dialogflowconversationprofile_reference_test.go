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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

func TestDialogflowConversationProfileRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/locations/us-central1/conversationProfiles/my-profile",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/locations/us-central1/conversationProfiles/my-profile",
			wantErr: true,
		},
		{
			name:    "missing location",
			ref:     "projects/my-project/conversationProfiles/my-profile",
			wantErr: true,
		},
		{
			name:    "missing profile",
			ref:     "projects/my-project/locations/us-central1",
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DialogflowConversationProfileRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("DialogflowConversationProfileRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDialogflowConversationProfileRef_ParseExternalToIdentity(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		wantStr string
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/locations/us-central1/conversationProfiles/my-profile",
			wantStr: "projects/my-project/locations/us-central1/conversationProfiles/my-profile",
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DialogflowConversationProfileRef{External: tt.ref}
			id, err := r.ParseExternalToIdentity()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseExternalToIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if id.String() != tt.wantStr {
					t.Errorf("id.String() = %v, want %v", id.String(), tt.wantStr)
				}
				if idV2, ok := id.(identity.IdentityV2); ok {
					if idV2.Host() != "dialogflow.googleapis.com" {
						t.Errorf("idV2.Host() = %v, want %v", idV2.Host(), "dialogflow.googleapis.com")
					}
				} else {
					t.Errorf("id does not implement identity.IdentityV2")
				}
			}
		})
	}
}
