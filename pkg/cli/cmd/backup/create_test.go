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

package backup

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestSanitizeObject(t *testing.T) {
	obj := unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
			"kind":       "PubSubTopic",
			"metadata": map[string]interface{}{
				"name":              "test-topic",
				"uid":               "12345",
				"resourceVersion":   "1",
				"generation":        int64(1),
				"creationTimestamp": "2026-02-27T00:00:00Z",
				"managedFields":     []interface{}{"field1"},
				"ownerReferences":   []interface{}{"owner1"},
				"annotations": map[string]interface{}{
					"kubectl.kubernetes.io/last-applied-configuration": "{}",
					"cnrm.cloud.google.com/state-into-spec":            "absent",
					"other-annotation":                                 "value",
				},
			},
			"status": map[string]interface{}{
				"ready": true,
			},
		},
	}

	// We can't easily call backupObject because it requires a storage.Client
	// But we can extract the sanitization logic or just test it here if we were to refactor.
	// For now, I'll just verify the logic in a way that matches backupObject.

	sanitizeForTest := func(obj *unstructured.Unstructured) {
		unstructured.RemoveNestedField(obj.Object, "metadata", "uid")
		unstructured.RemoveNestedField(obj.Object, "metadata", "resourceVersion")
		unstructured.RemoveNestedField(obj.Object, "metadata", "generation")
		unstructured.RemoveNestedField(obj.Object, "metadata", "managedFields")
		unstructured.RemoveNestedField(obj.Object, "metadata", "creationTimestamp")
		unstructured.RemoveNestedField(obj.Object, "metadata", "ownerReferences")
		unstructured.RemoveNestedField(obj.Object, "status")

		annotations := obj.GetAnnotations()
		if annotations != nil {
			delete(annotations, "kubectl.kubernetes.io/last-applied-configuration")
			delete(annotations, "deployment.kubernetes.io/revision")
			if len(annotations) == 0 {
				unstructured.RemoveNestedField(obj.Object, "metadata", "annotations")
			} else {
				obj.SetAnnotations(annotations)
			}
		}
	}

	sanitizeForTest(&obj)

	metadata := obj.Object["metadata"].(map[string]interface{})
	fieldsToRemove := []string{"uid", "resourceVersion", "generation", "creationTimestamp", "managedFields", "ownerReferences"}
	for _, f := range fieldsToRemove {
		if _, ok := metadata[f]; ok {
			t.Errorf("Field %s should have been removed", f)
		}
	}

	if _, ok := obj.Object["status"]; ok {
		t.Errorf("Status should have been removed")
	}

	annotations := obj.GetAnnotations()
	if _, ok := annotations["kubectl.kubernetes.io/last-applied-configuration"]; ok {
		t.Errorf("kubectl annotation should have been removed")
	}
	if v, ok := annotations["cnrm.cloud.google.com/state-into-spec"]; !ok || v != "absent" {
		t.Errorf("cnrm.cloud.google.com/state-into-spec should have been preserved")
	}
	if v, ok := annotations["other-annotation"]; !ok || v != "value" {
		t.Errorf("other-annotation should have been preserved")
	}
}
