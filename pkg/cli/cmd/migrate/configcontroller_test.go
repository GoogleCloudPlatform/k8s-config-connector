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

package migrate

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestStripUnwantedFields(t *testing.T) {
	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"metadata": map[string]interface{}{
				"name":              "test-resource",
				"namespace":         "test-ns",
				"creationTimestamp": "2023-01-01T00:00:00Z",
				"generation":        int64(1),
				"resourceVersion":   "12345",
				"uid":               "abcde",
				"managedFields":     []interface{}{"field1"},
				"generateName":      "test-resource-",
				"ownerReferences":   []interface{}{"owner1"},
				"finalizers":        []interface{}{"finalizer1"},
				"annotations": map[string]interface{}{
					"kubectl.kubernetes.io/last-applied-configuration": "{\"kind\":\"Test\"}",
					"other-annotation": "keep-me",
				},
			},
			"spec": map[string]interface{}{
				"field1": "value1",
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{"cond1"},
			},
		},
	}

	stripUnwantedFields(u)

	// Check status is removed
	if _, ok := u.Object["status"]; ok {
		t.Errorf("status should be removed")
	}

	// Check metadata fields are removed
	metadata := u.Object["metadata"].(map[string]interface{})
	unwantedFields := []string{"creationTimestamp", "generation", "resourceVersion", "uid", "managedFields", "generateName", "ownerReferences", "finalizers"}
	for _, field := range unwantedFields {
		if _, ok := metadata[field]; ok {
			t.Errorf("metadata.%s should be removed", field)
		}
	}

	// Check annotations
	annotations := u.GetAnnotations()
	if _, ok := annotations["kubectl.kubernetes.io/last-applied-configuration"]; ok {
		t.Errorf("last-applied-configuration annotation should be removed")
	}
	if annotations["other-annotation"] != "keep-me" {
		t.Errorf("other-annotation should be preserved")
	}
	if annotations["cnrm.cloud.google.com/management-conflict-prevention-policy"] != "none" {
		t.Errorf("management-conflict-prevention-policy should be set to none")
	}
}
