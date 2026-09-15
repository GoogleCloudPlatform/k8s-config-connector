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
	"strings"
	"testing"
	"text/template"

	"sigs.k8s.io/yaml"
)

func TestConfigureTemplateRendering(t *testing.T) {
	data := struct {
		ProjectID              string
		ClusterProjectID       string
		Cluster                string
		Bucket                 string
		BucketLocation         string
		Schedule               string
		Namespace              string
		Version                string
		DualRegion             string
		TurboReplication       bool
		Versioning             bool
		RetentionPeriodSeconds int
		LockRetention          bool
		Autopilot              bool
		ReplicaBucket          string
		ReplicaLocation        string
		IncludeClusterBackup   bool
		GKEBackupPlan          string
	}{
		ProjectID:              "prod-gcp-project",
		ClusterProjectID:       "prod-gcp-project",
		Cluster:                "gke-autopilot-cluster",
		Bucket:                 "kcc-backup-vault-primary",
		BucketLocation:         "nam4",
		Schedule:               "0 0 * * *",
		Namespace:              "cnrm-system",
		Version:                "1.125.0",
		DualRegion:             "nam4",
		TurboReplication:       true,
		Versioning:             true,
		RetentionPeriodSeconds: 30 * 86400,
		LockRetention:          true,
		Autopilot:              true,
		ReplicaBucket:          "kcc-backup-vault-secondary",
		ReplicaLocation:        "us-east1",
		IncludeClusterBackup:   true,
		GKEBackupPlan:          "gke-autopilot-cluster-backup-plan",
	}

	tmpl, err := template.New("configure").Parse(configureTemplate)
	if err != nil {
		t.Fatalf("Failed to parse configureTemplate: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}

	rendered := buf.String()

	// 1. Verify least privilege IAM roles (no objectAdmin)
	if strings.Contains(rendered, "roles/storage.objectAdmin") {
		t.Errorf("Template must not grant roles/storage.objectAdmin; expected least privilege creator/viewer")
	}
	if !strings.Contains(rendered, "roles/storage.objectCreator") {
		t.Errorf("Template missing roles/storage.objectCreator")
	}
	if !strings.Contains(rendered, "roles/storage.objectViewer") {
		t.Errorf("Template missing roles/storage.objectViewer")
	}

	// 2. Verify StorageBucket Data Protection: versioning, WORM retentionPolicy, lock, uniformBucketLevelAccess
	if !strings.Contains(rendered, "versioning:\n    enabled: true") {
		t.Errorf("Template missing versioning.enabled: true")
	}
	if !strings.Contains(rendered, "retentionPeriod: 2592000") {
		t.Errorf("Template missing retentionPeriod: 2592000 (30 days)")
	}
	if !strings.Contains(rendered, "isLocked: true") {
		t.Errorf("Template missing isLocked: true for WORM lock")
	}
	if !strings.Contains(rendered, "uniformBucketLevelAccess: true") {
		t.Errorf("Template missing uniformBucketLevelAccess: true")
	}
	if !strings.Contains(rendered, "publicAccessPrevention: enforced") {
		t.Errorf("Template missing publicAccessPrevention: enforced")
	}

	// 3. Verify Replica Bucket configuration
	if !strings.Contains(rendered, "name: kcc-backup-vault-secondary") {
		t.Errorf("Template missing replica bucket definition")
	}
	if !strings.Contains(rendered, "cnrm-backup-replica-creator") {
		t.Errorf("Template missing replica bucket creator IAM policy")
	}

	// 4. Verify GKE Autopilot & Standard compliant CronJob pod spec
	if !strings.Contains(rendered, "runAsNonRoot: true") {
		t.Errorf("CronJob missing runAsNonRoot: true")
	}
	if !strings.Contains(rendered, "seccompProfile:\n              type: RuntimeDefault") {
		t.Errorf("CronJob missing seccompProfile: RuntimeDefault")
	}
	if !strings.Contains(rendered, "drop:\n                - ALL") {
		t.Errorf("CronJob container missing capabilities drop ALL")
	}
	if !strings.Contains(rendered, "cpu: 250m") {
		t.Errorf("CronJob container missing resource request cpu: 250m")
	}
	if !strings.Contains(rendered, "- --include-cluster-backup") {
		t.Errorf("CronJob missing --include-cluster-backup flag")
	}
	if !strings.Contains(rendered, "- gke-autopilot-cluster-backup-plan") {
		t.Errorf("CronJob missing gke-backup-plan value")
	}

	// 5. Verify all documents unmarshal cleanly as valid YAML
	parts := bytes.Split(buf.Bytes(), []byte("\n---\n"))
	for i, part := range parts {
		if len(bytes.TrimSpace(part)) == 0 {
			continue
		}
		var parsed map[string]interface{}
		if err := yaml.Unmarshal(part, &parsed); err != nil {
			t.Errorf("Document %d failed to unmarshal YAML: %v", i, err)
		}
	}
}
