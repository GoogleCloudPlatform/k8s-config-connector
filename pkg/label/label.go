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
	"unicode"
	"unicode/utf8"
)

func NewGCPLabelsFromK8sLabels(labels map[string]string) map[string]string {
	res := SanitizeGCPLabels(labels)
	return res
}

// Sanitize labels with GCP label validation
func SanitizeGCPLabels(labels map[string]string) map[string]string {
	res := make(map[string]string)
keyLoop:
	for k, v := range labels {
		// GCP labels have strict validation rules. This function filters out any
		// labels from Kubernetes metadata that do not conform to these rules,
		// preventing them from being propagated to GCP.
		// See: https://cloud.google.com/compute/docs/labeling-resources#requirements

		// Key validation: 1-63 characters, lowercase letters, digits, underscores, hyphens. Must start with a letter.
		if utf8.RuneCountInString(k) < 1 || utf8.RuneCountInString(k) > 63 {
			continue
		}
		firstRune, _ := utf8.DecodeRuneInString(k)
		if !unicode.IsLower(firstRune) {
			continue
		}
		for _, r := range k {
			if !(unicode.IsLower(r) || unicode.IsDigit(r) || r == '_' || r == '-') {
				continue keyLoop
			}
		}

		// Value validation: 0-63 characters, lowercase letters, digits, underscores, hyphens.
		if utf8.RuneCountInString(v) > 63 {
			continue
		}
		for _, r := range v {
			if !(unicode.IsLower(r) || unicode.IsDigit(r) || r == '_' || r == '-') {
				continue keyLoop
			}
		}

		res[k] = v
	}
	// Add the managed-by label to indicate that this resource is managed by KCC.
	res[CnrmManagedKey] = "true"
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
