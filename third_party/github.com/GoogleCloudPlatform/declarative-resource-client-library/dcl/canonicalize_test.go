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

package dcl

import "testing"

func TestIsSelfLink(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "standard googleapis selfLink",
			input: "https://www.googleapis.com/compute/v1/projects/my-proj/global/networks/default",
			want:  true,
		},
		{
			name:  "service subdomain googleapis selfLink",
			input: "https://compute.googleapis.com/compute/v1/projects/my-proj/global/networks/default",
			want:  true,
		},
		{
			name:  "schemeless googleapis selfLink",
			input: "www.googleapis.com/compute/v1/projects/my-proj/global/networks/default",
			want:  true,
		},
		{
			name:  "custom universe .goog selfLink with scheme",
			input: "https://compute.custom.universe.goog/compute/v1/projects/my-proj/global/networks/default",
			want:  true,
		},
		{
			name:  "custom universe .goog selfLink without scheme",
			input: "compute.custom.universe.goog/compute/v1/projects/my-proj/global/networks/default",
			want:  true,
		},
		{
			name:  "custom sovereign domain selfLink with https",
			input: "https://compute.sovereign.example.org/compute/beta/projects/my-proj/regions/u-reg-1/subnetworks/sub-1",
			want:  true,
		},
		{
			name:  "partial selfLink",
			input: "projects/my-proj/global/networks/default",
			want:  false,
		},
		{
			name:  "plain resource name",
			input: "default",
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsSelfLink(tc.input); got != tc.want {
				t.Errorf("IsSelfLink(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestStringCanonicalizeAndEqualsWithCustomUniverseSelfLink(t *testing.T) {
	betaSelfLink := "https://compute.custom.universe.goog/compute/beta/projects/my-proj/regions/u-reg-1/subnetworks/sub-1"
	v1SelfLink := "https://compute.custom.universe.goog/compute/v1/projects/my-proj/regions/u-reg-1/subnetworks/sub-1"
	partialLink := "projects/my-proj/regions/u-reg-1/subnetworks/sub-1"
	shortName := "sub-1"
	differentSub := "https://compute.custom.universe.goog/compute/v1/projects/my-proj/regions/u-reg-1/subnetworks/sub-2"

	if !StringCanonicalize(&betaSelfLink, &v1SelfLink) {
		t.Errorf("StringCanonicalize(%q, %q) = false, want true", betaSelfLink, v1SelfLink)
	}
	if !StringCanonicalize(&betaSelfLink, &partialLink) {
		t.Errorf("StringCanonicalize(%q, %q) = false, want true", betaSelfLink, partialLink)
	}
	if !StringCanonicalize(&betaSelfLink, &shortName) {
		t.Errorf("StringCanonicalize(%q, %q) = false, want true", betaSelfLink, shortName)
	}
	if StringCanonicalize(&betaSelfLink, &differentSub) {
		t.Errorf("StringCanonicalize(%q, %q) = true, want false", betaSelfLink, differentSub)
	}
	if !StringEqualsWithSelfLink(&betaSelfLink, &v1SelfLink) {
		t.Errorf("StringEqualsWithSelfLink(%q, %q) = false, want true", betaSelfLink, v1SelfLink)
	}
}
