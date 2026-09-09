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

package common

import (
	"strings"
)

// ExtractProjectID extracts the project ID (or number) from a GCP resource name or URL.
// Supported formats:
// - https://[hostname]/[version]/projects/[projectID]/...
// - //compute.googleapis.com/projects/[projectID]/...
// - projects/[projectID]/...
func ExtractProjectID(resourceName string) string {
	if resourceName == "" {
		return ""
	}
	if strings.HasPrefix(resourceName, "https://") || strings.HasPrefix(resourceName, "//") {
		tokens := strings.Split(resourceName, "/")
		for i, token := range tokens {
			if token == "projects" && i+1 < len(tokens) {
				return tokens[i+1]
			}
		}
	}
	if strings.HasPrefix(resourceName, "projects/") {
		tokens := strings.Split(resourceName, "/")
		if len(tokens) > 1 {
			return tokens[1]
		}
	}
	return ""
}
