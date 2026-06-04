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

// IsRegion returns true if this string refers to a GCP region or multi-region.
func IsRegion(s *string) bool {
	if s == nil {
		return false
	}

	r := regexp.MustCompile(`^[a-z]+-[a-z]+[0-9]+$`)
	if r.MatchString(*s) {
		return true
	}

	// Also support multi-regions/locations (e.g., us, eu, europe, asia, in)
	// but exclude "global" to maintain backward compatibility for global URLs.
	switch *s {
	case "us", "eu", "europe", "asia", "in":
		return true
	}

	return false
}

// IsZone returns true if this string refers to a GCP zone.
func IsZone(s *string) bool {
	if s == nil {
		return false
	}

	r := regexp.MustCompile(`^[a-z]+-[a-z]+[0-9]+-[a-z]+$`)
	return r.MatchString(*s)
}
