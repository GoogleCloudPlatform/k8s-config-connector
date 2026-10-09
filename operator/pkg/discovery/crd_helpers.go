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

package discovery

import (
	"strings"
)

// ExtractGroupFromCRD parses the API group from a standard CRD name (e.g. "computeinstances.compute.cnrm.cloud.google.com").
func ExtractGroupFromCRD(crdName string) string {
	parts := strings.Split(crdName, ".")
	if len(parts) > 1 {
		return strings.Join(parts[1:], ".")
	}
	return crdName
}

// IsCRDSupported checks if a CRD's API group is permitted by the allowed service groups.
func IsCRDSupported(crdName string, allowedGroups map[string]bool) bool {
	group := ExtractGroupFromCRD(crdName)
	if CoreAPIGroups[group] {
		return true
	}
	return allowedGroups[group]
}
