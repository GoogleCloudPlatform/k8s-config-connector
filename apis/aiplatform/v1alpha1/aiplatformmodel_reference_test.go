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
	"strings"
	"testing"
)

func TestAIPlatformModelRef(t *testing.T) {
	testCases := []struct {
		name         string
		external     string
		expectError  bool
		expectedType string
	}{
		{
			name:         "valid custom model external reference",
			external:     "projects/test-project/locations/us-central1/models/test-model",
			expectError:  false,
			expectedType: "*v1alpha1.AIPlatformModelIdentity",
		},
		{
			name:         "valid regional publisher model external reference",
			external:     "projects/test-project/locations/us-central1/publishers/google/models/gemini-1.5-pro-001",
			expectError:  false,
			expectedType: "*v1alpha1.AIPlatformPublisherModelIdentity",
		},
		{
			name:         "valid global publisher model external reference",
			external:     "publishers/google/models/gemini-1.5-pro-001",
			expectError:  false,
			expectedType: "*v1alpha1.AIPlatformPublisherModelIdentity",
		},
		{
			name:        "invalid model external reference",
			external:    "invalid/model/path",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ref := &AIPlatformModelRef{
				External: tc.external,
			}

			valErr := ref.ValidateExternal(tc.external)
			parsedIdentity, parseErr := ref.ParseExternalToIdentity()

			if tc.expectError {
				if valErr == nil {
					t.Errorf("ValidateExternal(%q) expected error, got nil", tc.external)
				} else if !strings.Contains(valErr.Error(), "publishers/") {
					t.Errorf("ValidateExternal(%q) error %q does not mention publisher format", tc.external, valErr.Error())
				}
				if parseErr == nil {
					t.Errorf("ParseExternalToIdentity(%q) expected error, got nil", tc.external)
				} else if !strings.Contains(parseErr.Error(), "publishers/") {
					t.Errorf("ParseExternalToIdentity(%q) error %q does not mention publisher format", tc.external, parseErr.Error())
				}
			} else {
				if valErr != nil {
					t.Errorf("ValidateExternal(%q) unexpected error: %v", tc.external, valErr)
				}
				if parseErr != nil {
					t.Errorf("ParseExternalToIdentity() unexpected error: %v", parseErr)
				}
				if parsedIdentity == nil {
					t.Fatalf("ParseExternalToIdentity() returned nil identity")
				}
				if parsedIdentity.String() != tc.external {
					t.Errorf("identity.String() = %q, want %q", parsedIdentity.String(), tc.external)
				}
			}
		})
	}
}
