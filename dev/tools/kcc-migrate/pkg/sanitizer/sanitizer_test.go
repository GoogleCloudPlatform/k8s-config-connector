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

package sanitizer

import (
	"testing"
)

func TestSanitizeResource(t *testing.T) {
	rawDoc := map[string]interface{}{
		"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
		"kind":       "StorageBucket",
		"metadata": map[string]interface{}{
			"name":            "enterprise-test-bucket",
			"namespace":       "default",
			"uid":             "12345-67890",
			"resourceVersion": "987654",
			"generation":      2,
			"annotations": map[string]interface{}{
				"configmanagement.gke.io/cluster-name": "source-cc",
				"custom.enterprise.com/owner":                 "finance",
			},
		},
		"status": map[string]interface{}{
			"conditions": []interface{}{
				map[string]interface{}{
					"status": "True",
					"type":   "Ready",
				},
			},
		},
		"spec": map[string]interface{}{
			"location": "asia-southeast1",
		},
	}

	opts := SanitizeOptions{
		InjectAbandon:      true,
		InjectConflictNone: true,
	}

	sanitized, isKCC := SanitizeResource(rawDoc, opts)
	if !isKCC {
		t.Fatalf("Expected isKCC to be true")
	}

	// Verify status is stripped
	if _, hasStatus := sanitized["status"]; hasStatus {
		t.Errorf("Expected status to be deleted, but found status")
	}

	meta := sanitized["metadata"].(map[string]interface{})
	if _, hasUID := meta["uid"]; hasUID {
		t.Errorf("Expected uid to be deleted")
	}
	if _, hasRV := meta["resourceVersion"]; hasRV {
		t.Errorf("Expected resourceVersion to be deleted")
	}

	ann := meta["annotations"].(map[string]interface{})
	if ann[DeletionPolicyAnnotation] != PolicyAbandonValue {
		t.Errorf("Expected deletion-policy: abandon, got %v", ann[DeletionPolicyAnnotation])
	}
	if ann[ConflictPolicyAnnotation] != PolicyConflictNoneValue {
		t.Errorf("Expected management-conflict-prevention-policy: none, got %v", ann[ConflictPolicyAnnotation])
	}
	if _, hasConfigSync := ann["configmanagement.gke.io/cluster-name"]; hasConfigSync {
		t.Errorf("Expected Config Sync annotation to be stripped")
	}
	if ann["custom.enterprise.com/owner"] != "finance" {
		t.Errorf("Expected custom annotation to be preserved")
	}
}
