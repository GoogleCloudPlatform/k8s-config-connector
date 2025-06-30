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
	"strings"

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
		newLabels = removeLabelsWithKRMPrefix(u.GetLabels())
	}
	newLabels[CnrmManagedKey] = "true"
	return unstructured.SetNestedStringMap(u.Object, newLabels, "spec", "labels")
}

// NewGCPLabelsFromK8sLabels removes labels that violate GCP labels restrictions(i.e. KRM-style labels)
// and apply Config Connector default labels.
func NewGCPLabelsFromK8sLabels(labels map[string]string) map[string]string {
	res := removeLabelsWithKRMPrefix(labels)
	// Apply default label "managed-by-cnrm": "true"
	res[CnrmManagedKey] = "true"
	return res
}

func removeLabelsWithKRMPrefix(labels map[string]string) map[string]string {
	res := make(map[string]string)
	for k, v := range labels {
		// Do not include any KRM-style labels (labels that include a prefix denoted with a '/').
		// todo: verify if '/' is sufficient for identifying all system labels,
		// and are there edge cases we need to consider.
		if len(strings.Split(k, "/")) == 2 {
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
