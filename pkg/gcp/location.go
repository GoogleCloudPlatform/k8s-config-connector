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

func LocationToRegion(location string) (string, error) {
	if IsLocationRegional(location) {
		return location, nil
	}
	if !IsLocationZonal(location) {
		return "", fmt.Errorf("provided location is neither regional nor zonal")
	}
	idx := strings.LastIndex(location, "-")
	return location[:idx], nil
}

// IsLocationRegional returns true if the given location is a GCP region name.
// Standard commercial regions have two hyphen-separated segments (e.g. "us-central1"),
// while sovereign and Trusted Partner Cloud (TPC) regions may contain three or more
// hyphen-separated segments ending with a numeric digit (e.g. "u-region-1", "u-france-south1").
func IsLocationRegional(location string) bool {
	parts := strings.Split(location, "-")
	if len(parts) < 2 || hasEmptySegment(parts) {
		return false
	}
	if len(parts) == 2 {
		return true
	}
	return endsWithDigit(parts[len(parts)-1])
}

// IsLocationZonal returns true if the given location is a GCP zone name.
// Standard commercial zones have three hyphen-separated segments (e.g. "us-central1-a"),
// while sovereign and Trusted Partner Cloud (TPC) zones may contain four or more
// hyphen-separated segments where the final segment is a zone letter suffix (e.g. "u-region-1-a", "u-france-south1-b").
func IsLocationZonal(location string) bool {
	parts := strings.Split(location, "-")
	if len(parts) < 3 || hasEmptySegment(parts) {
		return false
	}
	if endsWithDigit(parts[len(parts)-1]) {
		return false
	}
	if len(parts) == 3 {
		return true
	}
	last := parts[len(parts)-1]
	return len(last) == 1 && last[0] >= 'a' && last[0] <= 'z' && endsWithDigit(parts[len(parts)-2])
}

func hasEmptySegment(parts []string) bool {
	for _, p := range parts {
		if p == "" {
			return true
		}
	}
	return false
}

func endsWithDigit(s string) bool {
	if len(s) == 0 {
		return false
	}
	c := s[len(s)-1]
	return c >= '0' && c <= '9'
}

const (
	Global   = "global"
	Regional = "regional"
	Zonal    = "zonal"
)
