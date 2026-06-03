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
	"fmt"
	"unicode"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// This function should be called if the typed object has `spec.labels` field.
func ComputeLabels(u *unstructured.Unstructured) error {
	var newLabels map[string]string
	specLabels, found, err := unstructured.NestedStringMap(u.Object, "spec", "labels")
	if err != nil {
		return fmt.Errorf("retrieve %s: %s `spec.labels` field: %w", u.GroupVersionKind().Kind, u.GetName(), err)
	}
	if specLabels != nil {
		newLabels = specLabels
	} else if found {
		newLabels = map[string]string{}
	} else {
		newLabels = u.GetLabels()
	}
	// No matter where the labels come from, sanitize them based on GCP label validation.
	newLabels = SanitizeGCPLabels(newLabels)
	newLabels[CnrmManagedKey] = "true"
	return unstructured.SetNestedStringMap(u.Object, newLabels, "spec", "labels")
}

func NewGCPLabelsFromK8sLabels(labels map[string]string) map[string]string {
	res := SanitizeGCPLabels(labels)
	// Apply default label.
	res[CnrmManagedKey] = "true"
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
		if len(k) < 1 || len(k) > 63 {
			continue
		}
		if !unicode.IsLower(rune(k[0])) {
			continue
		}
		for _, r := range k {
			if !(unicode.IsLower(r) || unicode.IsDigit(r) || r == '_' || r == '-') {
				continue keyLoop
			}
		}

		// Value validation: 0-63 characters, lowercase letters, digits, underscores, hyphens.
		if len(v) > 63 {
			continue
		}
		for _, r := range v {
			if !(unicode.IsLower(r) || unicode.IsDigit(r) || r == '_' || r == '-') {
				continue keyLoop
			}
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
