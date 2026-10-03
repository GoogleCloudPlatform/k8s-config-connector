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

package linter

import (
	"testing"
)

func TestLintResource(t *testing.T) {
	// Valid resource with abandon and conflict none
	validDoc := map[string]interface{}{
		"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
		"kind":       "StorageBucket",
		"metadata": map[string]interface{}{
			"name": "valid-bucket",
			"annotations": map[string]interface{}{
				DeletionPolicyAnnotation: "abandon",
				ConflictPolicyAnnotation: "none",
				StateIntoSpecAnnotation:  "absent",
			},
		},
		"spec": map[string]interface{}{
			"location": "asia-southeast1",
		},
	}

	issues := LintResource("test.yaml", validDoc, true)
	if len(issues) != 0 {
		t.Errorf("Expected 0 issues for validDoc, got %d: %v", len(issues), issues)
	}

	// Invalid resource: missing annotations and has runtime status
	invalidDoc := map[string]interface{}{
		"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
		"kind":       "StorageBucket",
		"metadata": map[string]interface{}{
			"name": "leaky-bucket",
			"uid":  "123-abc",
		},
		"status": map[string]interface{}{
			"conditions": []interface{}{},
		},
	}

	issues = LintResource("test.yaml", invalidDoc, true)
	if len(issues) < 3 {
		t.Errorf("Expected at least 3 issues (missing abandon, missing conflict, status leaked, uid leaked), got %d", len(issues))
	}
}
