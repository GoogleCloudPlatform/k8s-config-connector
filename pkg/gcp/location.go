// Copyright 2022 Google LLC
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

import (
	"fmt"
	"strings"
)

// LocationToRegion normalizes a zonal or regional location string into its containing region.
// For regional locations (e.g. "us-central1" or "u-region-1"), it returns the location unmodified.
// For zonal locations (e.g. "us-central1-a" or "u-region-1-a"), it returns the region prefix without the zone letter.
func LocationToRegion(location string) (string, error) {
	if IsLocationRegional(location) {
		return location, nil
	}
	if !IsLocationZonal(location) {
		return "", fmt.Errorf("provided location is neither regional nor zonal")
	}
	lastHyphen := strings.LastIndex(location, "-")
	return location[:lastHyphen], nil
}

// IsLocationRegional returns true if location represents a regional Google Cloud topology.
// It supports both commercial (e.g. "us-central1") and sovereign/partitioned (e.g. "u-region-1") region formats.
func IsLocationRegional(location string) bool {
	if IsLocationZonal(location) {
		return false
	}
	return strings.Contains(location, "-")
}

// IsLocationZonal returns true if location represents a zonal Google Cloud topology.
// A zone possesses at least one parent region hyphen and terminates with a single-letter zone identifier (e.g. "-a").
func IsLocationZonal(location string) bool {
	s := strings.Split(location, "-")
	if len(s) < 3 {
		return false
	}
	last := s[len(s)-1]
	return len(last) == 1 && last[0] >= 'a' && last[0] <= 'z'
}

const (
	Global   = "global"
	Regional = "regional"
	Zonal    = "zonal"
)
