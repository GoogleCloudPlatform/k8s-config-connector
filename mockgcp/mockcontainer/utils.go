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

package mockcontainer

import "strings"

func PtrTo[T any](t T) *T {
	return &t
}

func ValueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}

// AsZonalLink will convert a "location" link to a "zonal" link, if the location is a zone.
// For example, projects/${projectNumber}/locations/us-central1-a/operations/${operationID}
// will be converted to projects/${projectNumber}/zones/us-central1-a/operations/${operationID}
func AsZonalLink(link string) string {
	tokens := strings.Split(link, "/")

	for i := 0; i+1 < len(tokens); i++ {
		switch tokens[i] {
		case "locations":
			location := tokens[i+1]
			if isZone(location) {
				tokens[i] = "zones"
			}
		}
	}

	return strings.Join(tokens, "/")
}

// isZone returns true if the location appears to be a GCP zone (as oppposed to a region)
// The logic is pretty simple right now, based on the number of hyphens.
func isZone(location string) bool {
	tokens := strings.Split(location, "-")
	return len(tokens) == 3
}
