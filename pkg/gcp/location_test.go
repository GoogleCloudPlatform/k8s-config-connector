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

package gcp_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
)

func TestIsLocationRegional(t *testing.T) {
	tests := []struct {
		location string
		want     bool
	}{
		{"us-central1", true},
		{"europe-west3", true},
		{"asia-northeast1", true},
		{"southamerica-east1", true},
		// Sovereign and partitioned regions with prefix
		{"u-region-1", true},
		{"u-germany-northeast1", true},
		{"u-europe-west1", true},
		// Zonal locations should return false
		{"us-central1-a", false},
		{"europe-west3-c", false},
		{"u-region-1-a", false},
		{"u-germany-northeast1-b", false},
		// Global and edge cases
		{"global", false},
		{"", false},
		{"us", false},
	}

	for _, tc := range tests {
		t.Run(tc.location, func(t *testing.T) {
			got := gcp.IsLocationRegional(tc.location)
			if got != tc.want {
				t.Errorf("IsLocationRegional(%q) = %v, want %v", tc.location, got, tc.want)
			}
		})
	}
}

func TestIsLocationZonal(t *testing.T) {
	tests := []struct {
		location string
		want     bool
	}{
		{"us-central1-a", true},
		{"us-central1-b", true},
		{"europe-west3-c", true},
		{"asia-east1-f", true},
		// Sovereign and partitioned zones with prefix
		{"u-region-1-a", true},
		{"u-germany-northeast1-a", true},
		{"u-germany-northeast1-b", true},
		{"u-europe-west1-c", true},
		// Regional locations should return false
		{"us-central1", false},
		{"europe-west3", false},
		{"u-region-1", false},
		{"u-germany-northeast1", false},
		// Global and edge cases
		{"global", false},
		{"", false},
		{"us", false},
		{"us-central1-1", false}, // Trailing digit is not a zone
	}

	for _, tc := range tests {
		t.Run(tc.location, func(t *testing.T) {
			got := gcp.IsLocationZonal(tc.location)
			if got != tc.want {
				t.Errorf("IsLocationZonal(%q) = %v, want %v", tc.location, got, tc.want)
			}
		})
	}
}

func TestLocationToRegion(t *testing.T) {
	tests := []struct {
		location    string
		wantRegion  string
		expectError bool
	}{
		{"us-central1", "us-central1", false},
		{"us-central1-a", "us-central1", false},
		{"europe-west3-c", "europe-west3", false},
		{"u-region-1", "u-region-1", false},
		{"u-region-1-a", "u-region-1", false},
		{"u-germany-northeast1", "u-germany-northeast1", false},
		{"u-germany-northeast1-b", "u-germany-northeast1", false},
		{"global", "", true},
		{"", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.location, func(t *testing.T) {
			gotRegion, err := gcp.LocationToRegion(tc.location)
			if tc.expectError {
				if err == nil {
					t.Errorf("LocationToRegion(%q) expected error, got nil", tc.location)
				}
			} else {
				if err != nil {
					t.Errorf("LocationToRegion(%q) unexpected error: %v", tc.location, err)
				}
				if gotRegion != tc.wantRegion {
					t.Errorf("LocationToRegion(%q) = %q, want %q", tc.location, gotRegion, tc.wantRegion)
				}
			}
		})
	}
}
