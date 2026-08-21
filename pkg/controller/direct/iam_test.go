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

package direct

import (
	"testing"
)

func TestMemberMatches(t *testing.T) {
	tests := []struct {
		member     string
		wantMember string
		matches    bool
	}{
		{
			member:     "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    true,
		},
		{
			member:     "deleted:serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com?uid=12345",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    true,
		},
		{
			member:     "deleted:serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    true,
		},
		{
			member:     "serviceAccount:other@project.iam.gserviceaccount.com",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    false,
		},
		{
			member:     "deleted:serviceAccount:other@project.iam.gserviceaccount.com?uid=67890",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    false,
		},
		{
			member:     "serviceAccount:KCC-REPRO-SA@project.iam.gserviceaccount.com",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    true,
		},
		{
			member:     "DELETED:serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com?UID=12345",
			wantMember: "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			matches:    true,
		},
		{
			member:     "serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com",
			wantMember: "DELETED:serviceAccount:kcc-repro-sa@project.iam.gserviceaccount.com?UID=12345",
			matches:    true,
		},
	}

	for _, tc := range tests {
		got := memberMatches(tc.member, tc.wantMember)
		if got != tc.matches {
			t.Errorf("memberMatches(%q, %q) = %v, want %v", tc.member, tc.wantMember, got, tc.matches)
		}
	}
}
