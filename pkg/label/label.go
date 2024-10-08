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

package label

import (
	"strings"
)

func NewGCPLabelsFromK8sLabels(labels map[string]string) map[string]string {
	res := removeLabelsWithKRMPrefix(labels)
	// Apply default label.
	res[CnrmManagedKey] = "true"
	return res
}

func removeLabelsWithKRMPrefix(labels map[string]string) map[string]string {
	res := make(map[string]string)
	for k, v := range labels {
		if len(strings.Split(k, "/")) == 2 {
			// Do not include any KRM-style labels (labels that include a prefix
			// denoted with a '/').
			// TODO(b/137755194): Determine long-term solution.
			continue
		}
		res[k] = v
	}
	return res
}

// Reformat labels map into a JSON-compatible map[string]interface{} type.
func ToJSONCompatibleFormat(labels map[string]string) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range labels {
		res[k] = v
	}
	return res
}
