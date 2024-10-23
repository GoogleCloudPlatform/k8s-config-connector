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

package common

import "strings"

func ComputeGCPLabels(labels map[string]string) map[string]string {
	RemoveByPrefixes(labels, "cnrm.cloud.google.com")
	labels["managed-by-cnrm"] = "true"
	return labels
}

func RemoveByPrefixes(a map[string]string, prefixes ...string) {
	for k := range a {
		for i := range prefixes {
			if strings.HasPrefix(k, prefixes[i]) {
				delete(a, k)
			}
		}
	}
}
