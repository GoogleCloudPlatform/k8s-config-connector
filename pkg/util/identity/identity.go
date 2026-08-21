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

package identity

import (
	"fmt"
	"regexp"
)

// Project syntax can be found at https://docs.cloud.google.com/resource-manager/docs/creating-managing-projects
// A project ID has the following requirements:
// - Must be 6 to 30 characters in length.
// - Can only contain lowercase letters, numbers, and hyphens.
// - Must start with a letter.
// - Cannot end with a hyphen.
const ProjectIDRegexp string = "[a-z][a-z0-9-]{4,28}[a-z0-9]"

func ParseIdentityMap(ref string, parser *regexp.Regexp, cnt int) (error, map[string]string) {
	raw := parser.FindStringSubmatch(ref)
	if raw == nil {
		return fmt.Errorf("reference %s did not match expected format", ref), nil
	}
	result := make(map[string]string, cnt)
	for i, name := range parser.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = raw[i]
		}
	}
	if len(result) != cnt {
		return fmt.Errorf("reference %s failed to parse %d values", ref, cnt), nil
	}
	return nil, result
}
