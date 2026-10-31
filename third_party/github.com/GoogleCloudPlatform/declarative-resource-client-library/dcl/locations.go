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

import "regexp"

// IsRegion returns true if this string refers to a GCP region.
// Supports both standard two-token regions (e.g. "us-central1") and
// multi-token sovereign/TPC regions (e.g. "u-region-1", "u-france-south1").
func IsRegion(s *string) bool {
	if s == nil {
		return false
	}

	r := regexp.MustCompile(`^([a-z0-9]+-)+[a-z0-9]*[0-9]+$`)
	return r.MatchString(*s)
}

// IsZone returns true if this string refers to a GCP zone.
// Supports both standard three-token zones (e.g. "us-central1-a") and
// multi-token sovereign/TPC zones (e.g. "u-region-1-a", "u-france-south1-b").
func IsZone(s *string) bool {
	if s == nil {
		return false
	}

	r := regexp.MustCompile(`^([a-z0-9]+-)+[a-z0-9]*[0-9]+-[a-z]+$`)
	return r.MatchString(*s)
}
