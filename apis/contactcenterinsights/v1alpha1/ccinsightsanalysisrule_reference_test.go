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

func TestCCInsightsPhraseMatcherRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		hasError bool
	}{
		{
			name:     "Valid CCInsightsPhraseMatcher Ref",
			input:    "projects/my-project/locations/us-central1/phraseMatchers/my-phrasematcher",
			hasError: false,
		},
		{
			name:     "Valid CCInsightsPhraseMatcher Ref with Host",
			input:    "contactcenterinsights.googleapis.com/projects/my-project/locations/us-central1/phraseMatchers/my-phrasematcher",
			hasError: false,
		},
		{
			name:     "Invalid CCInsightsPhraseMatcher Ref",
			input:    "projects/my-project/locations/us-central1/invalid/my-phrasematcher",
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ref := &CCInsightsPhraseMatcherRef{External: tc.input}
			err := ref.ValidateExternal(tc.input)
			if tc.hasError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestDialogflowConversationProfileRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		hasError bool
	}{
		{
			name:     "Valid DialogflowConversationProfile Ref (Regional)",
			input:    "projects/my-project/locations/us-central1/conversationProfiles/my-profile",
			hasError: false,
		},
		{
			name:     "Valid DialogflowConversationProfile Ref (Global)",
			input:    "projects/my-project/conversationProfiles/my-profile",
			hasError: false,
		},
		{
			name:     "Invalid DialogflowConversationProfile Ref",
			input:    "projects/my-project/locations/us-central1/invalid/my-profile",
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ref := &DialogflowConversationProfileRef{External: tc.input}
			err := ref.Validate()
			if tc.hasError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
