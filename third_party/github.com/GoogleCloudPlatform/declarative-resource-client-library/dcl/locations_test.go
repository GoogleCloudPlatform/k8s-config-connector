// Copyright 2026 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package dcl

import "testing"

func TestIsRegionAndIsZone(t *testing.T) {
	tests := []struct {
		location   string
		wantRegion bool
		wantZone   bool
	}{
		{location: "global", wantRegion: false, wantZone: false},
		{location: "us-central1", wantRegion: true, wantZone: false},
		{location: "northamerica-northeast1", wantRegion: true, wantZone: false},
		{location: "u-region-1", wantRegion: true, wantZone: false},
		{location: "u-france-south1", wantRegion: true, wantZone: false},
		{location: "eu-sovereign-central1", wantRegion: true, wantZone: false},
		{location: "us-central1-a", wantRegion: false, wantZone: true},
		{location: "u-region-1-a", wantRegion: false, wantZone: true},
		{location: "u-france-south1-b", wantRegion: false, wantZone: true},
		{location: "eu-sovereign-central1-c", wantRegion: false, wantZone: true},
	}

	for _, tc := range tests {
		loc := tc.location
		if got := IsRegion(&loc); got != tc.wantRegion {
			t.Errorf("IsRegion(%q) = %v, want %v", tc.location, got, tc.wantRegion)
		}
		if got := IsZone(&loc); got != tc.wantZone {
			t.Errorf("IsZone(%q) = %v, want %v", tc.location, got, tc.wantZone)
		}
	}
	if IsRegion(nil) {
		t.Errorf("IsRegion(nil) = true, want false")
	}
	if IsZone(nil) {
		t.Errorf("IsZone(nil) = true, want false")
	}
}
