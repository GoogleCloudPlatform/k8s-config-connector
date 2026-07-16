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

package resourcefixture_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	dclextension "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"

	// Ensure built-in types are registered.
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

func TestVerifyLabelsTestCasesExist(t *testing.T) {
	smLoader, err := servicemappingloader.New()
	if err != nil {
		t.Fatalf("error getting service mapping loader: %v", err)
	}
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error getting dcl schema loader: %v", err)
	}
	serviceMetadataLoader := metadata.New()

	gvks, err := supportedgvks.All(smLoader, serviceMetadataLoader)
	if err != nil {
		t.Fatalf("error getting all supported GVKs: %v", err)
	}

	allFixtures := resourcefixture.Load(t)

	var matchingGVKs []schema.GroupVersionKind

	for _, gvk := range gvks {
		gk := schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}
		config, ok := resourceconfig.ControllerConfigStatic[gk]
		if !ok {
			continue
		}

		supportsTF := false
		supportsDCL := false
		for _, ctrl := range config.SupportedControllers {
			if ctrl == k8s.ReconcilerTypeTerraform {
				supportsTF = true
			}
			if ctrl == k8s.ReconcilerTypeDCL {
				supportsDCL = true
			}
		}

		hasLabels := false
		if supportsTF {
			rcs, err := smLoader.GetResourceConfigs(gvk)
			if err == nil {
				for _, rc := range rcs {
					if rc.MetadataMapping.Labels != "" {
						hasLabels = true
						break
					}
				}
			}
		}

		if supportsDCL {
			s, err := dclschemaloader.GetDCLSchemaForGVK(gvk, serviceMetadataLoader, dclSchemaLoader)
			if err == nil && s != nil {
				_, _, found, err := dclextension.GetLabelsFieldSchema(s)
				if err == nil && found {
					hasLabels = true
				}
			}
		}

		if hasLabels {
			// Only require labels test case if there is a basic test case for this GVK.
			if _, found := findBasicFixture(gvk, allFixtures); found {
				matchingGVKs = append(matchingGVKs, gvk)
			}
		}
	}

	createLabels := map[string]string{
		"label-one": "one",
		"label-two": "two",
	}

	updateLabels := map[string]string{
		"label-one":   "two",
		"label-three": "three",
	}

	packagePath := repo.GetCallerPackagePathOrTestFatal(t)
	testDataPath := filepath.Join(packagePath, "testdata")

	missingCount := 0

	for _, gvk := range matchingGVKs {
		kindLower := strings.ToLower(gvk.Kind)
		folderName := fmt.Sprintf("%s-labels", kindLower)
		folderPath := filepath.Join(testDataPath, "labels", folderName)
		createPath := filepath.Join(folderPath, "create.yaml")
		updatePath := filepath.Join(folderPath, "update.yaml")

		_, errCreate := os.Stat(createPath)
		_, errUpdate := os.Stat(updatePath)

		if os.IsNotExist(errCreate) || os.IsNotExist(errUpdate) {
			missingCount++
			t.Errorf("missing labels-only test case files for GVK %v (expected folder: %s)", gvk, folderPath)

			basicFixture, found := findBasicFixture(gvk, allFixtures)
			if !found {
				t.Logf("Warning: could not find any basic fixture for GVK %v, cannot auto-generate labels-only test case", gvk)
				continue
			}

			if err := os.MkdirAll(folderPath, 0755); err != nil {
				t.Fatalf("failed to create directory %s: %v", folderPath, err)
			}

			// Generate create.yaml and update.yaml
			createDocs := testyaml.SplitYAML(t, basicFixture.Create)
			var newCreateDocs [][]byte
			var newUpdateDocs [][]byte

			for _, docBytes := range createDocs {
				u := &unstructured.Unstructured{}
				if err := yaml.Unmarshal(docBytes, u); err != nil {
					t.Fatalf("error unmarshalling document: %v", err)
				}

				if u.GroupVersionKind() == gvk {
					// Main resource under test, set metadata.labels
					u.SetLabels(createLabels)
					cb, err := yaml.Marshal(u)
					if err != nil {
						t.Fatalf("error marshalling create document: %v", err)
					}
					newCreateDocs = append(newCreateDocs, cb)

					u.SetLabels(updateLabels)
					ub, err := yaml.Marshal(u)
					if err != nil {
						t.Fatalf("error marshalling update document: %v", err)
					}
					newUpdateDocs = append(newUpdateDocs, ub)
				} else {
					// Keep other documents unchanged
					newCreateDocs = append(newCreateDocs, docBytes)
					newUpdateDocs = append(newUpdateDocs, docBytes)
				}
			}

			var createStr []string
			for _, d := range newCreateDocs {
				createStr = append(createStr, string(d))
			}
			createYAML := strings.Join(createStr, "\n---\n")

			var updateStr []string
			for _, d := range newUpdateDocs {
				updateStr = append(updateStr, string(d))
			}
			updateYAML := strings.Join(updateStr, "\n---\n")

			if err := os.WriteFile(createPath, []byte(createYAML), 0644); err != nil {
				t.Fatalf("failed to write %s: %v", createPath, err)
			}
			if err := os.WriteFile(updatePath, []byte(updateYAML), 0644); err != nil {
				t.Fatalf("failed to write %s: %v", updatePath, err)
			}

			// If basic fixture has dependencies, copy them
			if len(basicFixture.Dependencies) > 0 {
				depPath := filepath.Join(folderPath, "dependencies.yaml")
				if err := os.WriteFile(depPath, basicFixture.Dependencies, 0644); err != nil {
					t.Fatalf("failed to write %s: %v", depPath, err)
				}
			}

			t.Logf("Auto-generated labels-only test case at %s", folderPath)
		}
	}

	if missingCount > 0 {
		t.Errorf("%d labels-only test cases were missing and have been auto-generated. Please check in the generated files.", missingCount)
	}
}

func findBasicFixture(gvk schema.GroupVersionKind, allFixtures []resourcefixture.ResourceFixture) (resourcefixture.ResourceFixture, bool) {
	// First, try to find one with "basic" in the name
	for _, f := range allFixtures {
		if f.Type == "basic" && f.GVK == gvk && strings.Contains(strings.ToLower(f.Name), "basic") {
			return f, true
		}
	}
	// Fallback to any basic fixture of the same GVK
	for _, f := range allFixtures {
		if f.Type == "basic" && f.GVK == gvk {
			return f, true
		}
	}
	return resourcefixture.ResourceFixture{}, false
}
