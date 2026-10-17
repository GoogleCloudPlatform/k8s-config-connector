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
	lastHyphen := strings.LastIndex(location, "-")
	return location[:lastHyphen], nil
}

func IsLocationRegional(location string) bool {
	if location == "" || location == Global {
		return false
	}
	parts := strings.Split(location, "-")
	if len(parts) < 2 {
		return false
	}
	lastPart := parts[len(parts)-1]
	// A zonal location ends in a single lowercase letter suffix (e.g. -a, -b, -c).
	// A regional location does not end in a single-letter zone suffix.
	if len(lastPart) == 1 && lastPart[0] >= 'a' && lastPart[0] <= 'z' {
		return false
	}
	return true
}

func IsLocationZonal(location string) bool {
	if location == "" || location == Global {
		return false
	}
	parts := strings.Split(location, "-")
	// A zonal location must have a region prefix plus a single lowercase letter zone suffix.
	// Minimum format is <prefix>-<region>-<zone> (at least 3 parts for standard, 4 for sovereign with u- prefix).
	if len(parts) < 3 {
		return false
	}
	lastPart := parts[len(parts)-1]
	return len(lastPart) == 1 && lastPart[0] >= 'a' && lastPart[0] <= 'z'
}

const (
	Global   = "global"
	Regional = "regional"
	Zonal    = "zonal"
)
