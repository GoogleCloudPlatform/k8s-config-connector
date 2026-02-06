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

package e2e

import (
	"context"
	"crypto/sha256"
	"fmt"
	rand "math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TestFindPathDependencies runs tests that try to find two paths to the same point in KRM space,
// that have different representations in GCP space.
func TestFindPathDependencies(t *testing.T) {
	if os.Getenv("RUN_PATH_DEPENDENCY_TEST") == "" {
		t.Skip("RUN_PATH_DEPENDENCY_TEST not set; skipping")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	uniqueID := "12345"
	namespace := "path-dependencies"

	waitForReadyTimeout := 15 * time.Second

	for i := 0; i < 50; i++ {
		t.Run(fmt.Sprintf("iteration-%d", i), func(t *testing.T) {
			ctx, cleanupContext := context.WithCancel(ctx)
			t.Cleanup(cleanupContext)

			h := NewHarness(ctx, t)
			project := h.Project

			testcontroller.SetupNamespaceForProject(h.T, h.GetClient(), namespace, project.ProjectID)

			obj := &unstructured.Unstructured{}
			obj.SetGroupVersionKind(schema.GroupVersionKind{
				Group:   "sql.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "SQLInstance",
			})
			obj.SetNamespace(namespace)
			obj.SetName(fmt.Sprintf("test-%v", uniqueID))

			unstructured.SetNestedField(obj.Object, "us-central1", "spec", "region")

			databaseVersion := pick([]string{"MYSQL_5_7", "MYSQL_8_0", "MYSQL_8_4"})
			unstructured.SetNestedField(obj.Object, databaseVersion, "spec", "databaseVersion")

			unstructured.SetNestedField(obj.Object, "CLOUD_SQL_INSTANCE", "spec", "instanceType")

			tier := pick([]string{"db-f1-micro", "db-g1-small", "db-n1-standard-1", "db-n1-standard-2", "db-custom-1-3840", "db-custom-2-7680", "db-custom-2-13312"})
			unstructured.SetNestedField(obj.Object, tier, "spec", "settings", "tier")

			applied := obj.DeepCopy()

			h.ApplyObject(obj)
			h.WaitForReady(waitForReadyTimeout, obj)

			exported := h.ExportObject(obj)
			if exported == nil {
				t.Fatalf("FAIL: exported object is nil")
			}

			appliedYAML := h.AsYAML(applied)
			hash := sha256.Sum256(appliedYAML)

			dir := filepath.Join("fuzzdata", "find_path_dependencies", fmt.Sprintf("%0x", hash[:]))
			if err := os.MkdirAll(dir, 0755); err != nil {
				t.Fatalf("FAIL: error creating directory %q: %v", dir, err)
			}

			appliedFile := filepath.Join(dir, "applied.yaml")
			if err := os.WriteFile(appliedFile, appliedYAML, 0644); err != nil {
				t.Fatalf("FAIL: error writing applied object to file %q: %v", appliedFile, err)
			}

			exportedYAML := h.AsYAML(exported)

			// pathFileHash is a hash of the applied+exported contents, to uniquely identify
			// this path + outcome.
			pathFileContents := string(appliedYAML) + "\n---\n#exported\n" + string(exportedYAML)
			pathFileHash := sha256.Sum256([]byte(pathFileContents))
			pathFile := filepath.Join(dir, fmt.Sprintf("path_%0x.yaml", pathFileHash[:]))
			if err := os.WriteFile(pathFile, []byte(pathFileContents), 0644); err != nil {
				t.Fatalf("FAIL: error writing path file to %q: %v", pathFile, err)
			}

			exportedFile := filepath.Join(dir, "exported.yaml")
			existing, err := os.ReadFile(exportedFile)
			if err == nil {
				if string(existing) != string(exportedYAML) {
					t.Fatalf("FAIL: found path dependency at %q: same applied object exported to different objects", exportedFile)
				}
			}

			t.Logf("writing exported object to %q", exportedFile)
			if err := os.WriteFile(exportedFile, exportedYAML, 0644); err != nil {
				t.Fatalf("FAIL: error writing exported object to file %q: %v", exportedFile, err)
			}
		})
	}
}

func pick(choices []string) string {
	n := rand.Intn(len(choices))
	return choices[n]
}
