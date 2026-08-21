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
	s := strings.Split(location, "-")
	return s[0] + "-" + s[1], nil
}

func IsLocationRegional(location string) bool {
	return len(strings.Split(location, "-")) == 2
}

func IsLocationZonal(location string) bool {
	return len(strings.Split(location, "-")) == 3
}

const (
	Global   = "global"
	Regional = "regional"
	Zonal    = "zonal"
)
