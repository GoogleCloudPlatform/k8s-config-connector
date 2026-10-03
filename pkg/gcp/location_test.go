// Copyright 2024 Google LLC
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

func TestLocationClassification(t *testing.T) {
	testCases := []struct {
		Location   string
		IsRegional bool
		IsZonal    bool
		Region     string
		ExpectErr  bool
	}{
		// Commercial Google Cloud regions and zones
		{
			Location:   "us-central1",
			IsRegional: true,
			IsZonal:    false,
			Region:     "us-central1",
		},
		{
			Location:   "us-central1-a",
			IsRegional: false,
			IsZonal:    true,
			Region:     "us-central1",
		},
		{
			Location:   "europe-west3",
			IsRegional: true,
			IsZonal:    false,
			Region:     "europe-west3",
		},
		{
			Location:   "europe-west3-c",
			IsRegional: false,
			IsZonal:    true,
			Region:     "europe-west3",
		},
		// Sovereign Cloud and partitioned regions and zones (e.g. u-region-1)
		{
			Location:   "u-region-1",
			IsRegional: true,
			IsZonal:    false,
			Region:     "u-region-1",
		},
		{
			Location:   "u-region-1-a",
			IsRegional: false,
			IsZonal:    true,
			Region:     "u-region-1",
		},
		{
			Location:   "u-region-1-b",
			IsRegional: false,
			IsZonal:    true,
			Region:     "u-region-1",
		},
		{
			Location:   "sec-us-east1",
			IsRegional: true,
			IsZonal:    false,
			Region:     "sec-us-east1",
		},
		{
			Location:   "sec-us-east1-a",
			IsRegional: false,
			IsZonal:    true,
			Region:     "sec-us-east1",
		},
		// Non-regional / Non-zonal topologies
		{
			Location:   "global",
			IsRegional: false,
			IsZonal:    false,
			ExpectErr:  true,
		},
		{
			Location:   "us",
			IsRegional: false,
			IsZonal:    false,
			ExpectErr:  true,
		},
		{
			Location:   "eu",
			IsRegional: false,
			IsZonal:    false,
			ExpectErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Location, func(t *testing.T) {
			regional := gcp.IsLocationRegional(tc.Location)
			if regional != tc.IsRegional {
				t.Errorf("IsLocationRegional(%q) = %v, want %v", tc.Location, regional, tc.IsRegional)
			}

			zonal := gcp.IsLocationZonal(tc.Location)
			if zonal != tc.IsZonal {
				t.Errorf("IsLocationZonal(%q) = %v, want %v", tc.Location, zonal, tc.IsZonal)
			}

			region, err := gcp.LocationToRegion(tc.Location)
			if tc.ExpectErr {
				if err == nil {
					t.Errorf("LocationToRegion(%q) expected error, got nil (region: %q)", tc.Location, region)
				}
			} else {
				if err != nil {
					t.Errorf("LocationToRegion(%q) unexpected error: %v", tc.Location, err)
				}
				if region != tc.Region {
					t.Errorf("LocationToRegion(%q) = %q, want %q", tc.Location, region, tc.Region)
				}
			}
		})
	}
}
