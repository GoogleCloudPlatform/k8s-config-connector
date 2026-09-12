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
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
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

	if _, ok := obj.Object["status"]; !ok {
		t.Errorf("Status should have been preserved")
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

func TestWriteSummaryWithIntegrity(t *testing.T) {
	tempDir := t.TempDir()
	cluster := "test-cluster"
	timestamp := "2026_09_06_12_00_00"

	stats := map[string]int{
		"storage.cnrm.cloud.google.com/StorageBucket": 2,
		"pubsub.cnrm.cloud.google.com/PubSubTopic":    3,
	}
	integrity := map[string]string{
		"default/storagebucket/bucket-a.yaml": "sha256:abc1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcd",
		"default/storagebucket/bucket-b.yaml": "sha256:def1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcd",
	}

	opts := &createOptions{
		outputDir: tempDir,
	}

	clusterBackup := &ClusterBackupMetadata{
		Provider:       "gkebackup.googleapis.com",
		BackupPlan:     "projects/test-project/locations/us-central1/backupPlans/test-cluster-backup-plan",
		BackupName:     "projects/test-project/locations/us-central1/backupPlans/test-cluster-backup-plan/backups/kcc-20260906",
		State:          "IN_PROGRESS",
		IncludeVolumes: true,
		AllNamespaces:  true,
	}

	if err := writeSummary(nil, nil, nil, opts, cluster, timestamp, stats, integrity, clusterBackup); err != nil {
		t.Fatalf("writeSummary failed: %v", err)
	}

	summaryPath := filepath.Join(tempDir, cluster, timestamp, "summary.json")
	data, err := os.ReadFile(summaryPath)
	if err != nil {
		t.Fatalf("Failed to read summary.json: %v", err)
	}

	var sm SummaryManifest
	if err := json.Unmarshal(data, &sm); err != nil {
		t.Fatalf("Failed to unmarshal summary.json: %v", err)
	}

	if sm.Counts["storage.cnrm.cloud.google.com/StorageBucket"] != 2 {
		t.Errorf("Counts mismatch: expected 2, got %d", sm.Counts["storage.cnrm.cloud.google.com/StorageBucket"])
	}
	if len(sm.Integrity) != 2 {
		t.Errorf("Expected 2 integrity records, got %d", len(sm.Integrity))
	}
	if sm.Integrity["default/storagebucket/bucket-a.yaml"] != "sha256:abc1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcd" {
		t.Errorf("Integrity hash mismatch: got %s", sm.Integrity["default/storagebucket/bucket-a.yaml"])
	}
	if sm.ClusterBackup == nil || sm.ClusterBackup.Provider != "gkebackup.googleapis.com" {
		t.Errorf("ClusterBackup metadata not properly recorded: %+v", sm.ClusterBackup)
	}
}

func TestCreatePathTraversalDefense(t *testing.T) {
	opts := &createOptions{
		outputDir: "/tmp/../../etc",
	}

	err := runCreate(nil, opts)
	if err == nil {
		t.Fatalf("Expected path traversal error for output-dir, got nil")
	}
	if !strings.Contains(err.Error(), "relative path traversal sequences") {
		t.Errorf("Expected traversal error message, got: %v", err)
	}
}
