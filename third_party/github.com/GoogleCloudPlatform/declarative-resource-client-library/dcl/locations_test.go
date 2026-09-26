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

func TestIsRegion(t *testing.T) {
	tests := []struct {
		location string
		want     bool
	}{
		{"us-central1", true},
		{"europe-west3", true},
		{"asia-northeast1", true},
		{"southamerica-east1", true},
		{"u-region-1", true},
		{"u-germany-northeast1", true},
		{"u-europe-west1", true},
		{"us-central1-a", false},
		{"europe-west3-b", false},
		{"u-region-1-a", false},
		{"u-germany-northeast1-a", false},
		{"global", false},
		{"", false},
	}

	for _, tc := range tests {
		t.Run(tc.location, func(t *testing.T) {
			got := IsRegion(&tc.location)
			if got != tc.want {
				t.Errorf("IsRegion(%q) = %v, want %v", tc.location, got, tc.want)
			}
		})
	}

	if IsRegion(nil) {
		t.Errorf("IsRegion(nil) = true, want false")
	}
}

func TestIsZone(t *testing.T) {
	tests := []struct {
		location string
		want     bool
	}{
		{"us-central1-a", true},
		{"europe-west3-b", true},
		{"asia-northeast1-c", true},
		{"u-region-1-a", true},
		{"u-germany-northeast1-a", true},
		{"u-europe-west1-b", true},
		{"us-central1", false},
		{"europe-west3", false},
		{"u-region-1", false},
		{"u-germany-northeast1", false},
		{"global", false},
		{"", false},
	}

	for _, tc := range tests {
		t.Run(tc.location, func(t *testing.T) {
			got := IsZone(&tc.location)
			if got != tc.want {
				t.Errorf("IsZone(%q) = %v, want %v", tc.location, got, tc.want)
			}
		})
	}

	if IsZone(nil) {
		t.Errorf("IsZone(nil) = true, want false")
	}
}
