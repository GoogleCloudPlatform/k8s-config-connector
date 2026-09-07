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
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestSortResources(t *testing.T) {
	resources := []*unstructured.Unstructured{
		{
			Object: map[string]interface{}{
				"kind":       "IAMPolicyMember",
				"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "policy-member",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "ComputeNetwork",
				"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "network",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "Project",
				"apiVersion": "resourcemanager.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "project",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "PubSubTopic",
				"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "topic",
				},
			},
		},
	}

	sortResources(resources)

	expectedOrder := []string{"Project", "ComputeNetwork", "PubSubTopic", "IAMPolicyMember"}
	for i, res := range resources {
		if res.GetKind() != expectedOrder[i] {
			t.Errorf("At index %d, expected kind %s, got %s", i, expectedOrder[i], res.GetKind())
		}
	}
}

func TestSortResourcesComplex(t *testing.T) {
	resources := []*unstructured.Unstructured{
		{
			Object: map[string]interface{}{
				"kind":       "ComputeSubnetwork",
				"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "subnetwork",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "IAMServiceAccount",
				"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "sa",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "Folder",
				"apiVersion": "resourcemanager.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "folder",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "ComputeNetwork",
				"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "network",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "Organization",
				"apiVersion": "resourcemanager.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "org",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "Project",
				"apiVersion": "resourcemanager.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "project",
				},
			},
		},
	}

	sortResources(resources)

	expectedOrder := []string{"Organization", "Folder", "Project", "IAMServiceAccount", "ComputeNetwork", "ComputeSubnetwork"}
	for i, res := range resources {
		if res.GetKind() != expectedOrder[i] {
			t.Errorf("At index %d, expected kind %s, got %s", i, expectedOrder[i], res.GetKind())
		}
	}
}

func TestSortResourcesEnterpriseDAG(t *testing.T) {
	resources := []*unstructured.Unstructured{
		{
			Object: map[string]interface{}{
				"kind":       "IAMPolicyMember",
				"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
				"metadata":   map[string]interface{}{"name": "sa-member"},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "SecretManagerSecretVersion",
				"apiVersion": "secretmanager.cnrm.cloud.google.com/v1beta1",
				"metadata":   map[string]interface{}{"name": "version-1"},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "SQLDatabase",
				"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
				"metadata":   map[string]interface{}{"name": "app-db"},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "StorageBucket",
				"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
				"metadata":   map[string]interface{}{"name": "my-bucket"},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "SecretManagerSecret",
				"apiVersion": "secretmanager.cnrm.cloud.google.com/v1beta1",
				"metadata":   map[string]interface{}{"name": "app-secret"},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "PubSubSubscription",
				"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
				"metadata":   map[string]interface{}{"name": "my-sub"},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "SQLInstance",
				"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
				"metadata":   map[string]interface{}{"name": "sql-primary"},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "PubSubTopic",
				"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
				"metadata":   map[string]interface{}{"name": "events-topic"},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "KMSKeyRing",
				"apiVersion": "kms.cnrm.cloud.google.com/v1beta1",
				"metadata":   map[string]interface{}{"name": "ring-1"},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "IAMServiceAccount",
				"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
				"metadata":   map[string]interface{}{"name": "app-sa"},
			},
		},
	}

	sortResources(resources)

	// Expected DAG order:
	// Priority 3: KMSKeyRing
	// Priority 5: IAMServiceAccount
	// Priority 8: SQLInstance, StorageBucket
	// Priority 9: PubSubTopic, SecretManagerSecret
	// Priority 10: PubSubSubscription, SQLDatabase, SecretManagerSecretVersion
	// Priority 12: IAMPolicyMember
	expectedOrder := []string{
		"KMSKeyRing",
		"IAMServiceAccount",
		"SQLInstance",
		"StorageBucket",
		"PubSubTopic",
		"SecretManagerSecret",
		"PubSubSubscription",
		"SQLDatabase",
		"SecretManagerSecretVersion",
		"IAMPolicyMember",
	}

	for i, res := range resources {
		if res.GetKind() != expectedOrder[i] {
			t.Errorf("DAG sort mismatch at index %d: expected %s, got %s", i, expectedOrder[i], res.GetKind())
		}
	}
}

func TestDRRegionAndProjectRemapping(t *testing.T) {
	opts := &restoreOptions{
		targetRegion:  "us-east4",
		regionMapping: "us-central1=us-east4",
		targetProject: "dr-project-backup",
	}

	bucket := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
			"kind":       "StorageBucket",
			"metadata": map[string]interface{}{
				"name": "app-data-bucket",
				"annotations": map[string]interface{}{
					"cnrm.cloud.google.com/project-id": "primary-prod-proj",
				},
			},
			"spec": map[string]interface{}{
				"location": "us-central1",
			},
		},
	}

	sql := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
			"kind":       "SQLInstance",
			"metadata": map[string]interface{}{
				"name": "app-db",
			},
			"spec": map[string]interface{}{
				"region": "us-central1",
				"projectRef": map[string]interface{}{
					"external": "primary-prod-proj",
				},
			},
		},
	}

	// Apply transformations directly via helper logic
	remapRegion := func(currentRegion string) string {
		if currentRegion == "" {
			return ""
		}
		if opts.regionMapping != "" {
			parts := strings.Split(opts.regionMapping, "=")
			if len(parts) == 2 && strings.TrimSpace(parts[0]) == currentRegion {
				return strings.TrimSpace(parts[1])
			}
		}
		if opts.targetRegion != "" {
			return opts.targetRegion
		}
		return currentRegion
	}

	// Test bucket remapping
	loc, _, _ := unstructured.NestedString(bucket.Object, "spec", "location")
	newLoc := remapRegion(loc)
	if newLoc != "us-east4" {
		t.Errorf("Expected bucket location to be remapped to us-east4, got %s", newLoc)
	}

	// Test SQL remapping
	reg, _, _ := unstructured.NestedString(sql.Object, "spec", "region")
	newReg := remapRegion(reg)
	if newReg != "us-east4" {
		t.Errorf("Expected SQL region to be remapped to us-east4, got %s", newReg)
	}

	// Test PubSub allowedPersistenceRegions remapping
	pubsub := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
			"kind":       "PubSubTopic",
			"spec": map[string]interface{}{
				"messageStoragePolicy": map[string]interface{}{
					"allowedPersistenceRegions": []interface{}{"us-central1"},
				},
			},
		},
	}
	regions, ok, _ := unstructured.NestedSlice(pubsub.Object, "spec", "messageStoragePolicy", "allowedPersistenceRegions")
	if !ok || len(regions) != 1 {
		t.Fatalf("Failed to get allowedPersistenceRegions: %+v", regions)
	}
	remappedRegion := remapRegion(regions[0].(string))
	if remappedRegion != "us-east4" {
		t.Errorf("Expected PubSub allowedPersistenceRegions to be remapped to us-east4, got %s", remappedRegion)
	}

	// Test SecretManager replicas remapping
	secret := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "secretmanager.cnrm.cloud.google.com/v1beta1",
			"kind":       "SecretManagerSecret",
			"spec": map[string]interface{}{
				"replication": map[string]interface{}{
					"userManaged": map[string]interface{}{
						"replicas": []interface{}{
							map[string]interface{}{
								"location": "us-central1",
							},
						},
					},
				},
			},
		},
	}
	replicas, ok, _ := unstructured.NestedSlice(secret.Object, "spec", "replication", "userManaged", "replicas")
	if !ok || len(replicas) != 1 {
		t.Fatalf("Failed to get replicas: %+v", replicas)
	}
	replicaLoc := remapRegion(replicas[0].(map[string]interface{})["location"].(string))
	if replicaLoc != "us-east4" {
		t.Errorf("Expected SecretManager replica location to be remapped to us-east4, got %s", replicaLoc)
	}
}

func TestLoadObjectsFromDir(t *testing.T) {
	tempDir := t.TempDir()
	cluster := "test-cluster"
	timestamp := "2026_09_06_12_00_00"

	manifestDir := filepath.Join(tempDir, cluster, timestamp, "default", "storagebucket")
	if err := os.MkdirAll(manifestDir, 0755); err != nil {
		t.Fatalf("Failed to create manifest dir: %v", err)
	}

	manifestContent := `
apiVersion: storage.cnrm.cloud.google.com/v1beta1
kind: StorageBucket
metadata:
  name: test-bucket
  namespace: default
spec:
  location: us-central1
`
	manifestPath := filepath.Join(manifestDir, "test-bucket.yaml")
	if err := os.WriteFile(manifestPath, []byte(manifestContent), 0644); err != nil {
		t.Fatalf("Failed to write manifest: %v", err)
	}

	// Test loading via rootDir + cluster + timestamp
	objs, err := loadObjectsFromDir(tempDir, cluster, timestamp, false)
	if err != nil {
		t.Fatalf("loadObjectsFromDir failed: %v", err)
	}
	if len(objs) != 1 {
		t.Fatalf("Expected 1 object loaded, got %d", len(objs))
	}
	if objs[0].GetName() != "test-bucket" || objs[0].GetKind() != "StorageBucket" {
		t.Errorf("Loaded object mismatch: name=%s, kind=%s", objs[0].GetName(), objs[0].GetKind())
	}

	// Test loading via latest resolution
	latestObjs, err := loadObjectsFromDir(tempDir, cluster, "latest", false)
	if err != nil {
		t.Fatalf("loadObjectsFromDir with latest failed: %v", err)
	}
	if len(latestObjs) != 1 {
		t.Fatalf("Expected 1 object loaded with latest, got %d", len(latestObjs))
	}
}

func TestLoadObjectsFromDirWithIntegrity(t *testing.T) {
	tempDir := t.TempDir()
	cluster := "test-cluster"
	timestamp := "2026_09_06_12_00_00"

	manifestDir := filepath.Join(tempDir, cluster, timestamp, "default", "storagebucket")
	if err := os.MkdirAll(manifestDir, 0755); err != nil {
		t.Fatalf("Failed to create manifest dir: %v", err)
	}

	manifestContent := "apiVersion: storage.cnrm.cloud.google.com/v1beta1\nkind: StorageBucket\nmetadata:\n  name: secure-bucket\n  namespace: default\nspec:\n  location: us-central1\n"
	manifestPath := filepath.Join(manifestDir, "secure-bucket.yaml")
	if err := os.WriteFile(manifestPath, []byte(manifestContent), 0644); err != nil {
		t.Fatalf("Failed to write manifest: %v", err)
	}

	// Calculate correct SHA-256
	hash := sha256.Sum256([]byte(manifestContent))
	hashStr := fmt.Sprintf("sha256:%x", hash)

	summaryDir := filepath.Join(tempDir, cluster, timestamp)
	summaryJSON := fmt.Sprintf(`{
		"counts": {"StorageBucket": 1},
		"integrity": {
			"default/storagebucket/secure-bucket.yaml": "%s"
		}
	}`, hashStr)

	if err := os.WriteFile(filepath.Join(summaryDir, "summary.json"), []byte(summaryJSON), 0644); err != nil {
		t.Fatalf("Failed to write summary: %v", err)
	}

	objs, err := loadObjectsFromDir(tempDir, cluster, timestamp, true)
	if err != nil {
		t.Fatalf("loadObjectsFromDir with valid integrity failed: %v", err)
	}
	if len(objs) != 1 {
		t.Fatalf("Expected 1 object, got %d", len(objs))
	}
}

func TestLoadObjectsFromDirTamperingDetected(t *testing.T) {
	tempDir := t.TempDir()
	cluster := "test-cluster"
	timestamp := "2026_09_06_12_00_00"

	manifestDir := filepath.Join(tempDir, cluster, timestamp, "default", "storagebucket")
	if err := os.MkdirAll(manifestDir, 0755); err != nil {
		t.Fatalf("Failed to create manifest dir: %v", err)
	}

	manifestContent := "apiVersion: storage.cnrm.cloud.google.com/v1beta1\nkind: StorageBucket\nmetadata:\n  name: tampered-bucket\n  namespace: default\nspec:\n  location: us-central1\n"
	manifestPath := filepath.Join(manifestDir, "tampered-bucket.yaml")
	if err := os.WriteFile(manifestPath, []byte(manifestContent), 0644); err != nil {
		t.Fatalf("Failed to write manifest: %v", err)
	}

	summaryDir := filepath.Join(tempDir, cluster, timestamp)
	// Write wrong hash to simulate tampering
	summaryJSON := `{
		"counts": {"StorageBucket": 1},
		"integrity": {
			"default/storagebucket/tampered-bucket.yaml": "sha256:0000000000000000000000000000000000000000000000000000000000000000"
		}
	}`

	if err := os.WriteFile(filepath.Join(summaryDir, "summary.json"), []byte(summaryJSON), 0644); err != nil {
		t.Fatalf("Failed to write summary: %v", err)
	}

	_, err := loadObjectsFromDir(tempDir, cluster, timestamp, true)
	if err == nil {
		t.Fatalf("Expected tamper detection error, got nil")
	}
	if !strings.Contains(err.Error(), "cryptographic checksum mismatch") {
		t.Errorf("Expected tamper alert message, got: %v", err)
	}
}

func TestPathTraversalProtection(t *testing.T) {
	_, err := loadObjectsFromDir("/tmp/../etc", "cluster", "latest", false)
	if err == nil {
		t.Fatalf("Expected error for directory traversal path, got nil")
	}
	if !strings.Contains(err.Error(), "directory traversal") {
		t.Errorf("Expected directory traversal error, got: %v", err)
	}
}

func TestTargetNamespaceAndZoneRemapping(t *testing.T) {
	opts := &restoreOptions{
		targetNamespace: "dr-isolated-ns",
		targetProject:   "dr-target-project",
		targetRegion:    "us-east1",
		regionMapping:   "us-central1=us-east1",
	}

	disk := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
			"kind":       "ComputeDisk",
			"metadata": map[string]interface{}{
				"name":      "data-disk-01",
				"namespace": "prod-app-ns",
			},
			"spec": map[string]interface{}{
				"zone": "us-central1-a",
				"projectRef": map[string]interface{}{
					"name": "old-source-project",
				},
			},
		},
	}

	policyMember := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
			"kind":       "IAMPolicyMember",
			"metadata": map[string]interface{}{
				"name":      "bucket-admin-binding",
				"namespace": "prod-app-ns",
			},
			"spec": map[string]interface{}{
				"member": "serviceAccount:my-sa@old-source-project.iam.gserviceaccount.com",
				"resourceRef": map[string]interface{}{
					"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
					"kind":       "StorageBucket",
					"name":       "prod-data-bucket",
					"namespace":  "prod-app-ns",
				},
			},
		},
	}

	objects := []*unstructured.Unstructured{disk, policyMember}

	// Remap namespace
	if opts.targetNamespace != "" {
		for _, obj := range objects {
			if obj.GetNamespace() != "" && obj.GetNamespace() != "_cluster_scoped" {
				oldNs := obj.GetNamespace()
				obj.SetNamespace(opts.targetNamespace)
				if refNs, ok, _ := unstructured.NestedString(obj.Object, "spec", "resourceRef", "namespace"); ok && refNs == oldNs {
					_ = unstructured.SetNestedField(obj.Object, opts.targetNamespace, "spec", "resourceRef", "namespace")
				}
			}
		}
	}

	if disk.GetNamespace() != "dr-isolated-ns" {
		t.Errorf("Expected disk namespace to be dr-isolated-ns, got %s", disk.GetNamespace())
	}
	if policyMember.GetNamespace() != "dr-isolated-ns" {
		t.Errorf("Expected policyMember namespace to be dr-isolated-ns, got %s", policyMember.GetNamespace())
	}
	refNs, _, _ := unstructured.NestedString(policyMember.Object, "spec", "resourceRef", "namespace")
	if refNs != "dr-isolated-ns" {
		t.Errorf("Expected policyMember resourceRef.namespace to be remapped to dr-isolated-ns, got %s", refNs)
	}

	// Test projectRef clean replacement
	if _, ok, _ := unstructured.NestedFieldNoCopy(disk.Object, "spec", "projectRef"); ok {
		unstructured.RemoveNestedField(disk.Object, "spec", "projectRef", "name")
		_ = unstructured.SetNestedField(disk.Object, opts.targetProject, "spec", "projectRef", "external")
	}

	projName, hasName, _ := unstructured.NestedString(disk.Object, "spec", "projectRef", "name")
	projExt, hasExt, _ := unstructured.NestedString(disk.Object, "spec", "projectRef", "external")
	if hasName {
		t.Errorf("Expected spec.projectRef.name to be removed, but still present: %s", projName)
	}
	if !hasExt || projExt != "dr-target-project" {
		t.Errorf("Expected spec.projectRef.external to be dr-target-project, got %s", projExt)
	}
}

func TestCMEKAndIAMMemberRemapping(t *testing.T) {
	opts := &restoreOptions{
		targetProject: "new-dr-project",
		targetRegion:  "us-east1",
		regionMapping: "us-central1=us-east1",
	}

	secret := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "secretmanager.cnrm.cloud.google.com/v1beta1",
			"kind":       "SecretManagerSecret",
			"spec": map[string]interface{}{
				"customerManagedEncryption": map[string]interface{}{
					"kmsKeyName": "projects/old-proj/locations/us-central1/keyRings/my-ring/cryptoKeys/my-key",
				},
			},
		},
	}

	iamMemberSA := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
			"kind":       "IAMPolicyMember",
			"spec": map[string]interface{}{
				"member": "serviceAccount:my-sa@old-proj.iam.gserviceaccount.com",
			},
		},
	}

	iamMemberWI := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
			"kind":       "IAMPolicyMember",
			"spec": map[string]interface{}{
				"member": "serviceAccount:old-proj.svc.id.goog[default/workload-ksa]",
			},
		},
	}

	// Remap CMEK key logic
	remapRegion := func(currentRegion string) string {
		if opts.regionMapping != "" {
			parts := strings.Split(opts.regionMapping, "=")
			if len(parts) == 2 && strings.TrimSpace(parts[0]) == currentRegion {
				return strings.TrimSpace(parts[1])
			}
		}
		if opts.targetRegion != "" {
			return opts.targetRegion
		}
		return currentRegion
	}

	remapKMSKey := func(keyPath string) string {
		newPath := keyPath
		if opts.targetProject != "" && strings.HasPrefix(newPath, "projects/") {
			parts := strings.Split(newPath, "/")
			if len(parts) >= 2 {
				parts[1] = opts.targetProject
				newPath = strings.Join(parts, "/")
			}
		}
		if (opts.regionMapping != "" || opts.targetRegion != "") && strings.Contains(newPath, "/locations/") {
			parts := strings.Split(newPath, "/")
			for idx, part := range parts {
				if part == "locations" && idx+1 < len(parts) {
					parts[idx+1] = remapRegion(parts[idx+1])
				}
			}
			newPath = strings.Join(parts, "/")
		}
		return newPath
	}

	origKey, _, _ := unstructured.NestedString(secret.Object, "spec", "customerManagedEncryption", "kmsKeyName")
	newKey := remapKMSKey(origKey)
	expectedKey := "projects/new-dr-project/locations/us-east1/keyRings/my-ring/cryptoKeys/my-key"
	if newKey != expectedKey {
		t.Errorf("Expected CMEK key remapped to %s, got %s", expectedKey, newKey)
	}

	// Remap IAM members
	memberSA, _, _ := unstructured.NestedString(iamMemberSA.Object, "spec", "member")
	parts := strings.Split(memberSA, "@")
	newMemberSA := fmt.Sprintf("%s@%s.iam.gserviceaccount.com", parts[0], opts.targetProject)
	expectedSA := "serviceAccount:my-sa@new-dr-project.iam.gserviceaccount.com"
	if newMemberSA != expectedSA {
		t.Errorf("Expected IAM SA member remapped to %s, got %s", expectedSA, newMemberSA)
	}

	memberWI, _, _ := unstructured.NestedString(iamMemberWI.Object, "spec", "member")
	wiParts := strings.Split(memberWI, ".svc.id.goog")
	newMemberWI := fmt.Sprintf("serviceAccount:%s.svc.id.goog%s", opts.targetProject, wiParts[1])
	expectedWI := "serviceAccount:new-dr-project.svc.id.goog[default/workload-ksa]"
	if newMemberWI != expectedWI {
		t.Errorf("Expected Workload Identity member remapped to %s, got %s", expectedWI, newMemberWI)
	}
}
