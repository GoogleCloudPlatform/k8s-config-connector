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

package gcp

import "testing"

func TestIsLocationRegionalAndZonal(t *testing.T) {
	tests := []struct {
		location   string
		wantRegion bool
		wantZone   bool
		region     string
	}{
		{
			location:   "global",
			wantRegion: false,
			wantZone:   false,
		},
		{
			location:   "us-central1",
			wantRegion: true,
			wantZone:   false,
			region:     "us-central1",
		},
		{
			location:   "europe-west9",
			wantRegion: true,
			wantZone:   false,
			region:     "europe-west9",
		},
		{
			location:   "u-region-1",
			wantRegion: true,
			wantZone:   false,
			region:     "u-region-1",
		},
		{
			location:   "u-france-south1",
			wantRegion: true,
			wantZone:   false,
			region:     "u-france-south1",
		},
		{
			location:   "eu-sovereign-central1",
			wantRegion: true,
			wantZone:   false,
			region:     "eu-sovereign-central1",
		},
		{
			location:   "us-central1-a",
			wantRegion: false,
			wantZone:   true,
			region:     "us-central1",
		},
		{
			location:   "u-region-1-a",
			wantRegion: false,
			wantZone:   true,
			region:     "u-region-1",
		},
		{
			location:   "u-france-south1-b",
			wantRegion: false,
			wantZone:   true,
			region:     "u-france-south1",
		},
		{
			location:   "eu-sovereign-central1-c",
			wantRegion: false,
			wantZone:   true,
			region:     "eu-sovereign-central1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.location, func(t *testing.T) {
			if got := IsLocationRegional(tc.location); got != tc.wantRegion {
				t.Errorf("IsLocationRegional(%q) = %v, want %v", tc.location, got, tc.wantRegion)
			}
			if got := IsLocationZonal(tc.location); got != tc.wantZone {
				t.Errorf("IsLocationZonal(%q) = %v, want %v", tc.location, got, tc.wantZone)
			}
			gotRegion, err := LocationToRegion(tc.location)
			if tc.wantRegion || tc.wantZone {
				if err != nil {
					t.Errorf("LocationToRegion(%q) returned unexpected error: %v", tc.location, err)
				}
				if gotRegion != tc.region {
					t.Errorf("LocationToRegion(%q) = %q, want %q", tc.location, gotRegion, tc.region)
				}
			} else if err == nil {
				t.Errorf("LocationToRegion(%q) = %q, expected error", tc.location, gotRegion)
			}
		})
	}
}
