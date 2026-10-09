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

package judgement

import "testing"

// TestIsReferenceReason pins which reasons hide a TestMissingRefs finding
// while their entry is open. reference-not-representable is not one of them:
// TestMissingRefs never fails on such a field, so there is nothing to hide.
func TestIsReferenceReason(t *testing.T) {
	for _, tc := range []struct {
		reason string
		want   bool
	}{
		{ReasonPossibleReference, true},
		{ReasonPossibleReferenceByDescription, true},
		{ReasonPossibleReferenceByDescriptionLoose, true},
		{ReasonPossibleReferenceByName, true},
		{ReasonPossibleReferenceBySibling, true},
		{ReasonReferenceNotRepresentable, false},
		{"untriaged-bulk-generation", false},
	} {
		t.Run(tc.reason, func(t *testing.T) {
			// Act
			got := IsReferenceReason(tc.reason)

			// Assert
			if got != tc.want {
				t.Errorf("IsReferenceReason(%q) = %v, want %v", tc.reason, got, tc.want)
			}
		})
	}
}
