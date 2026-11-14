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
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestBackupLifecycleAndRestore(t *testing.T) {
	tmpDir := t.TempDir()
	inputDir := filepath.Join(tmpDir, "exported")
	outputDir := filepath.Join(tmpDir, "snapshots")
	restoreDir := filepath.Join(tmpDir, "restored")

	if err := os.MkdirAll(inputDir, 0o755); err != nil {
		t.Fatalf("MkdirAll failed: %v", err)
	}

	rawManifests := `apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubSubscription
metadata:
  name: orders-sub
  namespace: team-a
  uid: "1111-2222"
  resourceVersion: "998877"
  generation: 4
  annotations:
    cnrm.cloud.google.com/project-id: source-proj
    kubectl.kubernetes.io/last-applied-configuration: '{"old":"config"}'
spec:
  topicRef:
    name: orders-topic
status:
  externalRef: projects/source-proj/subscriptions/orders-sub-real
---
apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubTopic
metadata:
  name: orders-topic
  namespace: team-a
  uid: "3333-4444"
  resourceVersion: "112233"
spec:
  resourceID: orders-topic
---
apiVersion: resourcemanager.cnrm.cloud.google.com/v1beta1
kind: Project
metadata:
  name: source-proj
  namespace: team-a
spec:
  resourceID: source-proj
---
apiVersion: sql.cnrm.cloud.google.com/v1beta1
kind: SQLInstance
metadata:
  name: main-db
  namespace: team-a
spec:
  region: us-central1
  projectRef:
    name: source-proj
`
	if err := os.WriteFile(filepath.Join(inputDir, "resources.yaml"), []byte(rawManifests), 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	var buf bytes.Buffer
	err := RunCreate(&CreateOptions{
		Cluster:   "prod-cluster",
		Location:  "us-central1",
		Project:   "source-proj",
		InputDir:  inputDir,
		OutputDir: outputDir,
		Timestamp: "2026-10-25",
	}, &buf)
	if err != nil {
		t.Fatalf("RunCreate failed: %v", err)
	}

	snapshotDir := filepath.Join(outputDir, "2026-10-25")
	buf.Reset()
	if err := RunStatus(&StatusOptions{BackupDir: snapshotDir}, &buf); err != nil {
		t.Fatalf("RunStatus failed: %v", err)
	}
	if !strings.Contains(buf.String(), "Status: Succeeded") || !strings.Contains(buf.String(), "Resources: 4") {
		t.Errorf("unexpected status output: %s", buf.String())
	}

	buf.Reset()
	if err := RunRestore(&RestoreOptions{
		BackupDir:      snapshotDir,
		TargetProject:  "dr-proj",
		TargetLocation: "us-east1",
		OutputDir:      restoreDir,
		DryRun:         true,
	}, &buf); err != nil {
		t.Fatalf("RunRestore failed: %v", err)
	}

	restoredObjects, err := loadYAMLObjectsFromDir(restoreDir)
	if err != nil {
		t.Fatalf("loadYAMLObjectsFromDir failed: %v", err)
	}

	// Verify first 4 individual ordered files (Project -> PubSubTopic -> PubSubSubscription / SQLInstance)
	if len(restoredObjects) < 4 {
		t.Fatalf("expected at least 4 restored objects, got %d", len(restoredObjects))
	}
	if restoredObjects[0].GetKind() != "Project" {
		t.Errorf("expected first restored kind to be Project, got %s", restoredObjects[0].GetKind())
	}
	if restoredObjects[1].GetKind() != "PubSubTopic" {
		t.Errorf("expected second restored kind to be PubSubTopic, got %s", restoredObjects[1].GetKind())
	}

	for _, obj := range restoredObjects[:4] {
		if obj.GetUID() != "" || obj.GetResourceVersion() != "" {
			t.Errorf("expected runtime metadata stripped on %s, got uid=%q rv=%q", obj.GetName(), obj.GetUID(), obj.GetResourceVersion())
		}
		if _, found, _ := unstructured.NestedMap(obj.Object, "status"); found {
			t.Errorf("expected status stripped on %s", obj.GetName())
		}
		ann := obj.GetAnnotations()
		if ann[AnnotationDeletionPolicy] != DeletionPolicyAbandon {
			t.Errorf("expected deletion-policy=abandon on %s, got %q", obj.GetName(), ann[AnnotationDeletionPolicy])
		}
		if ann[AnnotationConflictPreventionPolicy] != ConflictPreventionPolicyNone {
			t.Errorf("expected conflict-prevention-policy=none on %s, got %q", obj.GetName(), ann[AnnotationConflictPreventionPolicy])
		}
		if obj.GetName() == "orders-sub" {
			gotID, _, _ := unstructured.NestedString(obj.Object, "spec", "resourceID")
			if gotID != "orders-sub-real" {
				t.Errorf("expected spec.resourceID preserved from status.externalRef as 'orders-sub-real', got %q", gotID)
			}
			if ann[AnnotationProjectID] != "dr-proj" {
				t.Errorf("expected remapped project-id 'dr-proj', got %q", ann[AnnotationProjectID])
			}
		}
		if obj.GetName() == "main-db" {
			gotRegion, _, _ := unstructured.NestedString(obj.Object, "spec", "region")
			if gotRegion != "us-east1" {
				t.Errorf("expected remapped region 'us-east1', got %q", gotRegion)
			}
			gotProjExt, _, _ := unstructured.NestedString(obj.Object, "spec", "projectRef", "external")
			if gotProjExt != "projects/dr-proj" {
				t.Errorf("expected remapped spec.projectRef.external 'projects/dr-proj', got %q", gotProjExt)
			}
		}
	}
}

func TestRunConfigure(t *testing.T) {
	var buf bytes.Buffer
	err := RunConfigure(&ConfigureOptions{
		Cluster:        "prod-cluster",
		Location:       "us-central1",
		Bucket:         "my-kcc-dr-bucket",
		BucketLocation: "US",
		Frequency:      "daily",
		Project:        "my-dr-project",
	}, &buf)
	if err != nil {
		t.Fatalf("RunConfigure failed: %v", err)
	}
	out := buf.String()
	for _, want := range []string{"kind: StorageBucket", "kind: ServiceAccount", "kind: CronJob", `schedule: "0 2 * * *"`} {
		if !strings.Contains(out, want) {
			t.Errorf("expected configure output to contain %q, got:\n%s", want, out)
		}
	}
}
