// Copyright 2024 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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
		url  string
		want bool
	}{
		{"https://compute.googleapis.com/compute/v1/projects/my-proj/global/networks/my-net", true},
		{"https://www.googleapis.com/compute/v1/projects/my-proj/global/networks/my-net", true},
		{"https://compute.custom.universe.goog/compute/v1/projects/my-proj/regions/u-region-1/subnetworks/my-subnet", true},
		{"https://custom.universe.goog/compute/v1/projects/my-proj/regions/u-region-1/subnetworks/my-subnet", true},
		{"http://compute.googleapis.com/compute/v1/projects/my-proj/global/networks/my-net", true},
		{"compute.googleapis.com/compute/v1/projects/my-proj/global/networks/my-net", true},
		{"projects/my-proj/global/networks/my-net", false},
		{"my-net", false},
		{"", false},
	}

	for _, tc := range tests {
		t.Run(tc.url, func(t *testing.T) {
			got := IsSelfLink(tc.url)
			if got != tc.want {
				t.Errorf("IsSelfLink(%q) = %v, want %v", tc.url, got, tc.want)
			}
		})
	}
}

func TestStringEqualsWithSelfLink(t *testing.T) {
	strPtr := func(s string) *string {
		return &s
	}

	tests := []struct {
		name  string
		left  *string
		right *string
		want  bool
	}{
		{
			name:  "identical self links",
			left:  strPtr("https://compute.googleapis.com/compute/v1/projects/p/regions/r/serviceAttachments/sa"),
			right: strPtr("https://compute.googleapis.com/compute/v1/projects/p/regions/r/serviceAttachments/sa"),
			want:  true,
		},
		{
			name:  "v1 vs beta self links in standard googleapis",
			left:  strPtr("https://compute.googleapis.com/compute/beta/projects/p/regions/r/serviceAttachments/sa"),
			right: strPtr("https://compute.googleapis.com/compute/v1/projects/p/regions/r/serviceAttachments/sa"),
			want:  true,
		},
		{
			name:  "v1 vs beta self links in custom universe domain",
			left:  strPtr("https://compute.custom.universe.goog/compute/beta/projects/p/regions/r/serviceAttachments/sa"),
			right: strPtr("https://compute.custom.universe.goog/compute/v1/projects/p/regions/r/serviceAttachments/sa"),
			want:  true,
		},
		{
			name:  "partial self link vs full self link in custom universe domain",
			left:  strPtr("projects/p/regions/r/serviceAttachments/sa"),
			right: strPtr("https://compute.custom.universe.goog/compute/v1/projects/p/regions/r/serviceAttachments/sa"),
			want:  true,
		},
		{
			name:  "different resources in custom universe domain",
			left:  strPtr("https://compute.custom.universe.goog/compute/v1/projects/p/regions/r/serviceAttachments/sa1"),
			right: strPtr("https://compute.custom.universe.goog/compute/v1/projects/p/regions/r/serviceAttachments/sa2"),
			want:  false,
		},
		{
			name:  "both nil",
			left:  nil,
			right: nil,
			want:  true,
		},
		{
			name:  "one nil",
			left:  strPtr("sa"),
			right: nil,
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := StringEqualsWithSelfLink(tc.left, tc.right)
			if got != tc.want {
				t.Errorf("StringEqualsWithSelfLink() = %v, want %v", got, tc.want)
			}
		})
	}
}
