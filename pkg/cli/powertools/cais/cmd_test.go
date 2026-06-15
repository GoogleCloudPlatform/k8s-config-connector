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

package cais

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dns/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cais"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cais/caistesting"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/objects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

func TestReadObjectsAndGetCAISIdentities(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()

	tempDir, err := os.MkdirTemp("", "cais-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	yamlContent := `
apiVersion: dns.cnrm.cloud.google.com/v1beta1
kind: DNSManagedZone
metadata:
  name: test-zone
  namespace: test-namespace
  annotations:
    cnrm.cloud.google.com/project-id: "my-project-id"
spec:
  dnsName: "test.example.com."
`

	tempFile := filepath.Join(tempDir, "zone.yaml")
	if err := os.WriteFile(tempFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	objects, err := cais.ReadObjects(false, tempFile, "")
	if err != nil {
		t.Fatalf("ReadObjects failed: %v", err)
	}

	if len(objects) != 1 {
		t.Fatalf("expected 1 object, got %d", len(objects))
	}

	u := objects[0]
	if u.GetName() != "test-zone" || u.GetNamespace() != "test-namespace" {
		t.Errorf("unexpected object metadata: %s/%s", u.GetNamespace(), u.GetName())
	}

	reader := cais.NewInMemoryReader(scheme, objects)

	results, err := cais.GetCAISIdentities(ctx, scheme, reader, objects)
	if err != nil {
		t.Fatalf("GetCAISIdentities failed: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	res := results[0]
	if res.Kind != "DNSManagedZone" {
		t.Errorf("expected Kind DNSManagedZone, got %s", res.Kind)
	}
	if res.CAISURL != "//dns.googleapis.com/projects/my-project-id/managedZones/test-zone" {
		t.Errorf("unexpected CAIS URL: %s", res.CAISURL)
	}

	// Test formatting
	var outBuf bytes.Buffer
	if err := formatResults(&outBuf, results, "yaml"); err != nil {
		t.Fatalf("formatting to yaml failed: %v", err)
	}

	wantYAML := "- caisURL: //dns.googleapis.com/projects/my-project-id/managedZones/test-zone\n  group: dns.cnrm.cloud.google.com\n  kind: DNSManagedZone\n  name: test-zone\n  namespace: test-namespace\n  version: v1beta1\n"
	if diff := cmp.Diff(wantYAML, outBuf.String()); diff != "" {
		t.Errorf("YAML output mismatch (-want +got):\n%s", diff)
	}

	outBuf.Reset()
	if err := formatResults(&outBuf, results, "json"); err != nil {
		t.Fatalf("formatting to json failed: %v", err)
	}

	wantJSON := `[
  {
    "group": "dns.cnrm.cloud.google.com",
    "version": "v1beta1",
    "kind": "DNSManagedZone",
    "namespace": "test-namespace",
    "name": "test-zone",
    "caisURL": "//dns.googleapis.com/projects/my-project-id/managedZones/test-zone"
  }
]
`
	if diff := cmp.Diff(wantJSON, outBuf.String()); diff != "" {
		t.Errorf("JSON output mismatch (-want +got):\n%s", diff)
	}

	outBuf.Reset()
	if err := formatResults(&outBuf, results, "text"); err != nil {
		t.Fatalf("formatting to text failed: %v", err)
	}

	wantText := "GROUP                       VERSION   KIND             NAMESPACE        NAME        CAIS_URL                                                             ERROR\ndns.cnrm.cloud.google.com   v1beta1   DNSManagedZone   test-namespace   test-zone   //dns.googleapis.com/projects/my-project-id/managedZones/test-zone   \n"
	if diff := cmp.Diff(wantText, outBuf.String()); diff != "" {
		t.Errorf("Text output mismatch (-want +got):\n%s", diff)
	}
}

func TestInMemoryReaderNotFound(t *testing.T) {
	scheme := runtime.NewScheme()

	reader := cais.NewInMemoryReader(scheme, []*unstructured.Unstructured{})

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "dns.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "DNSManagedZone",
	})

	err := reader.Get(context.Background(), client.ObjectKey{}, u)
	if err == nil {
		t.Fatalf("expected NotFound error, got nil")
	}
}

func TestGoldenIdentitiesYamlFiles(t *testing.T) {
	ctx := context.Background()
	scheme := cais.NewScheme()

	// Initialize the exact fake org/project/billing variables as E2E test harness using our refactored method
	caistesting.InitializeFakeGCPVariables()

	rootPath := filepath.Join("..", "..", "..", "test", "resourcefixture", "testdata")

	// Verify or generate _identities.yaml files for all test folders that contain create.yaml
	err := filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if d.Name() != "create.yaml" {
			return nil
		}

		dir := filepath.Dir(path)

		var allObjs []*unstructured.Unstructured

		gcpProject := testgcp.GCPProject{ProjectID: "mock-project", ProjectNumber: 1234567890}

		var depBytes []byte

		// Parse dependencies.yaml first to exactly match the E2E test sequence/ordering
		depPath := filepath.Join(dir, "dependencies.yaml")
		if _, err := os.Stat(depPath); err == nil {
			var err error
			depBytes, err = os.ReadFile(depPath)
			if err != nil {
				t.Fatalf("failed to read dependencies.yaml for %s: %v", dir, err)
			}
			// Expand placeholders using ReplaceTestVars exactly like E2E tests
			depBytesExpanded := testcontroller.ReplaceTestVars(t, depBytes, "puxvndidajatl5i", gcpProject)
			depBytesExpanded = cleanDoubleQuotes(depBytesExpanded)

			depObjs, err := objects.ParseObjectsFromStream(bytes.NewReader(depBytesExpanded))
			if err != nil {
				t.Fatalf("failed to parse dependencies.yaml for %s: %v", dir, err)
			}
			allObjs = append(allObjs, depObjs...)
		}

		// Parse create.yaml (the primary resource under test) and append it last
		createBytes, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read create.yaml for %s: %v", dir, err)
		}
		createBytesExpanded := testcontroller.ReplaceTestVars(t, createBytes, "puxvndidajatl5i", gcpProject)
		createBytesExpanded = cleanDoubleQuotes(createBytesExpanded)

		createObjs, err := objects.ParseObjectsFromStream(bytes.NewReader(createBytesExpanded))
		if err != nil {
			t.Fatalf("failed to parse create.yaml for %s: %v", dir, err)
		}
		allObjs = append(allObjs, createObjs...)

		// Filter to KCC resources only, exactly as in E2E tests
		var kccObjs []*unstructured.Unstructured
		for _, u := range allObjs {
			group := u.GroupVersionKind().Group
			if strings.HasSuffix(group, ".cnrm.cloud.google.com") && group != "core.cnrm.cloud.google.com" {
				kccObjs = append(kccObjs, u)
			}
		}

		// Pre-populate missing default namespaces and project ID annotations to emulate live Kube E2E tests
		for _, u := range kccObjs {
			if u.GetNamespace() == "" {
				u.SetNamespace("puxvndidajatl5i")
			}
			annotations := u.GetAnnotations()
			if annotations == nil {
				annotations = make(map[string]string)
			}
			if annotations["cnrm.cloud.google.com/project-id"] == "" {
				annotations["cnrm.cloud.google.com/project-id"] = "mock-project"
				u.SetAnnotations(annotations)
			}
		}

		reader := cais.NewInMemoryReader(scheme, kccObjs)

		results, err := cais.GetCAISIdentities(ctx, scheme, reader, kccObjs)
		if err != nil {
			t.Fatalf("failed to get CAIS identities for %s: %v", dir, err)
		}

		if len(results) == 0 {
			return nil
		}

		caisYAML, err := yaml.Marshal(results)
		if err != nil {
			t.Fatalf("failed to marshal CAIS identities to YAML for %s: %v", dir, err)
		}

		caisYAMLStr := string(caisYAML)

		// Leverage our refactored placeholder-replacement method
		caisYAMLStr = caistesting.ReplacePlaceholdersInCAIS(caisYAMLStr, dir, createBytes, depBytes)

		identitiesPath := filepath.Join(dir, "_identities.yaml")

		if _, err := os.Stat(identitiesPath); os.IsNotExist(err) && os.Getenv("WRITE_GOLDEN_OUTPUT") == "" {
			return nil
		}

		test.CompareGoldenFile(t, identitiesPath, caisYAMLStr, caistesting.NormalizeDynamicIDs)

		return nil
	})

	if err != nil {
		t.Fatalf("walking basic test fixtures failed: %v", err)
	}
}

// cleanDoubleQuotes removes any duplicate double quotes created when placeholder expansion replaces inside quoted strings
func cleanDoubleQuotes(b []byte) []byte {
	s := string(b)
	s = strings.ReplaceAll(s, `""example-project-01""`, `"example-project-01"`)
	s = strings.ReplaceAll(s, `""example-project-02""`, `"example-project-02"`)
	s = strings.ReplaceAll(s, `""123450001""`, `"123450001"`)
	s = strings.ReplaceAll(s, `""123450002""`, `"123450002"`)
	s = strings.ReplaceAll(s, `""123451001""`, `"123451001"`)
	s = strings.ReplaceAll(s, `""123451002""`, `"123451002"`)
	s = strings.ReplaceAll(s, `""123456-777777-000001""`, `"123456-777777-000001"`)
	s = strings.ReplaceAll(s, `""123456-777777-000002""`, `"123456-777777-000002"`)
	s = strings.ReplaceAll(s, `""123456-777777-000003""`, `"123456-777777-000003"`)
	s = strings.ReplaceAll(s, `""mock-project""`, `"mock-project"`)
	return []byte(s)
}
