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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	tfmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

func getMinimalSpec(gvk schema.GroupVersionKind) map[string]interface{} {
	switch gvk.Kind {
	case "HealthcareDICOMStore", "HealthcareConsentStore", "HealthcareHL7V2Store":
		return map[string]interface{}{
			"datasetRef": map[string]interface{}{
				"name": "dummy-dataset",
			},
		}
	case "HealthcareFHIRStore":
		return map[string]interface{}{
			"datasetRef": map[string]interface{}{
				"name": "dummy-dataset",
			},
			"version": "DSTU2",
		}
	case "VertexAIFeaturestoreEntityType":
		return map[string]interface{}{
			"featurestoreRef": map[string]interface{}{
				"name": "dummy-featurestore",
			},
		}
	case "VertexAIFeaturestoreEntityTypeFeature":
		return map[string]interface{}{
			"entityTypeRef": map[string]interface{}{
				"name": "dummy-entitytype",
			},
			"valueType": "DOUBLE",
		}
	case "DatastreamStream":
		return map[string]interface{}{
			"location":    "us-central1",
			"displayName": "dummy-stream",
			"sourceConnectionProfileRef": map[string]interface{}{
				"name": "dummy-conn",
			},
			"destinationConnectionProfileRef": map[string]interface{}{
				"name": "dummy-conn",
			},
		}
	case "APIGatewayAPIConfig":
		return map[string]interface{}{
			"apiRef": map[string]interface{}{
				"name": "dummy-api",
			},
			"openapiDocuments": []interface{}{
				map[string]interface{}{
					"document": map[string]interface{}{
						"path":     "openapi.yaml",
						"contents": "swagger: '2.0'\ninfo:\n  title: API Gateway\n  version: 1.0.0",
					},
				},
			},
		}
	case "CloudFunctions2Function":
		return map[string]interface{}{
			"location": "us-central1",
			"buildConfig": map[string]interface{}{
				"entryPoint": "helloGet",
				"source": map[string]interface{}{
					"storageSource": map[string]interface{}{
						"bucketRef": map[string]interface{}{
							"name": "dummy-bucket",
						},
						"object": "dummy-object",
					},
				},
			},
		}
	case "VertexAIIndexEndpoint":
		return map[string]interface{}{
			"projectRef": map[string]interface{}{
				"external": "projects/dummy-project",
			},
			"location":    "us-central1",
			"displayName": "dummy-index-endpoint",
		}
	case "FirebaseHostingChannel":
		return map[string]interface{}{
			"siteRef": map[string]interface{}{
				"name": "dummy-site",
			},
		}
	case "DialogflowCXIntent":
		return map[string]interface{}{
			"parentRef": map[string]interface{}{
				"name": "dummy-parent",
			},
			"displayName": "dummy-intent",
		}
	case "FilestoreSnapshot":
		return map[string]interface{}{
			"instanceRef": map[string]interface{}{
				"name": "dummy-instance",
			},
		}
	case "APIGatewayGateway":
		return map[string]interface{}{
			"apiConfigRef": map[string]interface{}{
				"name": "dummy-apiconfig",
			},
			"gatewayId": "dummy-gateway",
		}
	case "NetworkServicesEdgeCacheService":
		return map[string]interface{}{
			"projectRef": map[string]interface{}{
				"external": "projects/dummy-project",
			},
		}
	default:
		return map[string]interface{}{}
	}
}

func TestLabelsCoverage(t *testing.T) {
	smLoader, err := servicemappingloader.New()
	if err != nil {
		t.Fatalf("error creating servicemappingloader: %v", err)
	}
	dclMetaLoader := dclmetadata.New()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating dclschemaloader: %v", err)
	}

	allGVKs, err := supportedgvks.All(smLoader, dclMetaLoader)
	if err != nil {
		t.Fatalf("error loading all GVKs: %v", err)
	}

	staticConfig := resourceconfig.LoadConfig()

	var expectedLabelKinds []schema.GroupVersionKind

	for _, gvk := range allGVKs {
		cfg, ok := staticConfig[gvk.GroupKind()]
		if !ok {
			continue
		}

		isTFSupported := false
		isDCLSupported := false
		for _, c := range cfg.SupportedControllers {
			if c == k8s.ReconcilerTypeTerraform {
				isTFSupported = true
			}
			if c == k8s.ReconcilerTypeDCL {
				isDCLSupported = true
			}
		}

		supportsLabels := false

		if isTFSupported {
			var matchedResourceConfig *v1alpha1.ResourceConfig
			for _, sm := range smLoader.GetServiceMappings() {
				for _, resourceConfig := range sm.Spec.Resources {
					if tfmetadata.GVKForResource(&sm, &resourceConfig) == gvk {
						matchedResourceConfig = &resourceConfig
						break
					}
				}
				if matchedResourceConfig != nil {
					break
				}
			}
			if matchedResourceConfig != nil && matchedResourceConfig.MetadataMapping.Labels != "" {
				supportsLabels = true
			}
		}

		if !supportsLabels && isDCLSupported {
			dclSchema, err := dclschemaloader.GetDCLSchemaForGVK(gvk, dclMetaLoader, dclSchemaLoader)
			if err == nil && dclSchema != nil {
				_, _, found, err := extension.GetLabelsFieldSchema(dclSchema)
				if err == nil && found {
					supportsLabels = true
				}
			}
		}

		if supportsLabels {
			expectedLabelKinds = append(expectedLabelKinds, gvk)
		}
	}

	t.Logf("Found %d resources that support labels", len(expectedLabelKinds))

	// Load existing fixtures to find basic blueprints to copy from
	allFixtures := resourcefixture.Load(t)
	basicFixtures := make(map[schema.GroupVersionKind]resourcefixture.ResourceFixture)
	for _, f := range allFixtures {
		if f.Type == resourcefixture.Basic {
			if strings.Contains(strings.ToLower(f.Name), "basic") {
				basicFixtures[f.GVK] = f
			} else {
				if _, ok := basicFixtures[f.GVK]; !ok {
					basicFixtures[f.GVK] = f
				}
			}
		}
	}
	// Fallback to any fixture if no basic fixture exists
	for _, f := range allFixtures {
		if _, ok := basicFixtures[f.GVK]; !ok {
			basicFixtures[f.GVK] = f
		}
	}

	for _, gvk := range expectedLabelKinds {
		kind := gvk.Kind
		folderName := fmt.Sprintf("%s-labels", strings.ToLower(kind))
		t.Run(folderName, func(t *testing.T) {
			labelsDir := filepath.Join("testdata", "labels", folderName)
			if _, err := os.Stat(labelsDir); os.IsNotExist(err) {
				// Auto-generate labels test case!
				var blueprint resourcefixture.ResourceFixture
				blueprintFound := false
				blueprint, ok := basicFixtures[gvk]
				if ok {
					blueprintFound = true
				} else {
					// Fallback 1: try to find any blueprint matching Kind regardless of Group or Version
					for bgvk, b := range basicFixtures {
						if bgvk.Kind == gvk.Kind {
							blueprint = b
							blueprintFound = true
							break
						}
					}
					if !blueprintFound {
						// Fallback 2: Look in allFixtures as a last resort
						for _, f := range allFixtures {
							if f.GVK.Kind == gvk.Kind {
								blueprint = f
								blueprintFound = true
								break
							}
						}
					}
				}

				if err := os.MkdirAll(labelsDir, 0755); err != nil {
					t.Fatalf("error creating labels test case directory %s: %v", labelsDir, err)
				}

				var createObj, updateObj *unstructured.Unstructured

				if blueprintFound {
					// Create create.yaml from blueprint
					createObj = &unstructured.Unstructured{}
					if err := yaml.Unmarshal(blueprint.Create, createObj); err != nil {
						t.Fatalf("error unmarshalling blueprint create for GVK %v: %v", gvk, err)
					}
					createObj.SetGroupVersionKind(gvk)

					// Create update.yaml from blueprint
					updateObj = &unstructured.Unstructured{}
					if err := yaml.Unmarshal(blueprint.Create, updateObj); err != nil {
						t.Fatalf("error unmarshalling blueprint create for GVK %v: %v", gvk, err)
					}
					updateObj.SetGroupVersionKind(gvk)
				} else {
					// Fallback 3: Create from scratch using minimal spec
					createObj = &unstructured.Unstructured{}
					createObj.SetGroupVersionKind(gvk)
					createObj.SetName(fmt.Sprintf("%s-labels-test", strings.ToLower(kind)))
					createObj.SetNamespace("test-namespace")
					createObj.Object["spec"] = getMinimalSpec(gvk)

					updateObj = &unstructured.Unstructured{}
					updateObj.SetGroupVersionKind(gvk)
					updateObj.SetName(fmt.Sprintf("%s-labels-test", strings.ToLower(kind)))
					updateObj.SetNamespace("test-namespace")
					updateObj.Object["spec"] = getMinimalSpec(gvk)
				}

				createObj.SetLabels(map[string]string{
					"label-one": "one",
					"label-two": "two",
				})
				createBytes, err := yaml.Marshal(createObj)
				if err != nil {
					t.Fatalf("error marshalling create.yaml: %v", err)
				}
				if err := os.WriteFile(filepath.Join(labelsDir, "create.yaml"), createBytes, 0644); err != nil {
					t.Fatalf("error writing create.yaml: %v", err)
				}

				updateObj.SetLabels(map[string]string{
					"label-one":   "two",
					"label-three": "three",
				})
				updateBytes, err := yaml.Marshal(updateObj)
				if err != nil {
					t.Fatalf("error marshalling update.yaml: %v", err)
				}
				if err := os.WriteFile(filepath.Join(labelsDir, "update.yaml"), updateBytes, 0644); err != nil {
					t.Fatalf("error writing update.yaml: %v", err)
				}

				// Create dependencies.yaml if they exist
				if blueprintFound && len(blueprint.Dependencies) > 0 {
					if err := os.WriteFile(filepath.Join(labelsDir, "dependencies.yaml"), blueprint.Dependencies, 0644); err != nil {
						t.Fatalf("error writing dependencies.yaml: %v", err)
					}
				}

				t.Logf("Generated labels test case for %s under %s", kind, labelsDir)
			}

			// Read create.yaml
			createPath := filepath.Join(labelsDir, "create.yaml")
			createBytes, err := os.ReadFile(createPath)
			if err != nil {
				t.Errorf("error reading %s: %v", createPath, err)
				return
			}

			// Read update.yaml
			updatePath := filepath.Join(labelsDir, "update.yaml")
			updateBytes, err := os.ReadFile(updatePath)
			if err != nil {
				t.Errorf("error reading %s: %v", updatePath, err)
				return
			}

			// Validate create.yaml
			createObj := &unstructured.Unstructured{}
			if err := yaml.Unmarshal(createBytes, createObj); err != nil {
				t.Errorf("error unmarshalling %s: %v", createPath, err)
				return
			}

			if createObj.GetKind() != kind {
				t.Errorf("kind mismatch in %s: expected %q, got %q", createPath, kind, createObj.GetKind())
			}

			// Verify metadata.labels in create.yaml
			createLabels := createObj.GetLabels()
			if createLabels == nil || createLabels["label-one"] != "one" || createLabels["label-two"] != "two" {
				t.Errorf("%s metadata.labels must contain 'label-one: one' and 'label-two: two', got: %v", createPath, createLabels)
			}

			// Validate update.yaml
			updateObj := &unstructured.Unstructured{}
			if err := yaml.Unmarshal(updateBytes, updateObj); err != nil {
				t.Errorf("error unmarshalling %s: %v", updatePath, err)
				return
			}

			if updateObj.GetKind() != kind {
				t.Errorf("kind mismatch in %s: expected %q, got %q", updatePath, kind, updateObj.GetKind())
			}

			// Verify metadata.labels in update.yaml
			updateLabels := updateObj.GetLabels()
			if updateLabels == nil || updateLabels["label-one"] != "two" || updateLabels["label-three"] != "three" {
				t.Errorf("%s metadata.labels must contain 'label-one: two' and 'label-three: three', got: %v", updatePath, updateLabels)
			}

			// Verify label-two is NOT in update.yaml labels
			if _, exists := updateLabels["label-two"]; exists {
				t.Errorf("%s metadata.labels must NOT contain 'label-two', got: %v", updatePath, updateLabels)
			}
		})
	}
}
