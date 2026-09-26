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

package codegen

import "strings"

// Singular strips a regular English plural and gives up on anything else. It
// stays narrow on purpose. A match here puts a question in front of a reviewer,
// and no GCP resource name uses an irregular plural.
//
// The scaffolder uses it to name a parent reference from the collection segment
// of a google.api.resource pattern, so "keyRings/{key_ring}" can be reported as
// a KMSKeyRing rather than a KMSKeyRings.
func Singular(s string) string {
	switch {
	case strings.HasSuffix(s, "ies") && len(s) > 4:
		// policies -> policy
		return s[:len(s)-3] + "y"
	case strings.HasSuffix(s, "sses"), strings.HasSuffix(s, "shes"),
		strings.HasSuffix(s, "ches"), strings.HasSuffix(s, "xes"),
		strings.HasSuffix(s, "zes"):
		// addresses -> address
		return s[:len(s)-2]
	case strings.HasSuffix(s, "ss"), strings.HasSuffix(s, "us"),
		strings.HasSuffix(s, "is"):
		// access, status, analysis: not plurals at all.
		return s
	case strings.HasSuffix(s, "s") && len(s) > 2:
		// subnetworks -> subnetwork
		return s[:len(s)-1]
	}
	return s
}
