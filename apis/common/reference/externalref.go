// Copyright 2025 Google LLC
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

package reference

import (
	"fmt"
	"strings"
)

// ParseExternalRef parses externalRef and returns a map of its components.
// For example: "projects/test-project/locations/us-central1/instances/test-instance"
// will be parsed into:
//
//	{
//	  "externalRef": "projects/test-project/locations/us-central1/instances/test-instance",
//	  "project": "test-project",
//	  "location": "us-central1",
//	  "instance": "test-instance",
//	  "name": "test-instance",
//	}
func ParseExternalRef(externalRef string) (map[string]string, error) {
	if externalRef == "" {
		return nil, fmt.Errorf("externalRef is empty")
	}
	components := make(map[string]string)
	components["externalRef"] = externalRef

	tokens := strings.Split(externalRef, "/")
	components["name"] = tokens[len(tokens)-1]
	// If the number of parts is even, we can build key-value pairs from parts.
	// Otherwise, we just get the name from the last part.
	// todo: handle "projects/test-project/global/instances/test-instance"
	if len(tokens)%2 == 0 {
		for i := 0; i < len(tokens); i += 2 {
			key := tokens[i]
			value := tokens[i+1]
			// Singularize name for key
			// todo: handle more plural cases?
			singularKey := ""
			if strings.HasSuffix(key, "ves") {
				singularKey = strings.TrimSuffix(key, "ves") + "f"
			} else if strings.HasSuffix(key, "ies") {
				singularKey = strings.TrimSuffix(key, "ies") + "y"
			} else if strings.HasSuffix(key, "s") {
				singularKey = strings.TrimSuffix(key, "s")
			} else {
				singularKey = key
			}
			components[singularKey] = value
		}
	}
	return components, nil
}
