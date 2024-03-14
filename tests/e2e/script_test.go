// Copyright 2024 Google LLC
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
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	kccyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/yaml"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
)

// TestE2EScript runs a Scenario test that runs step-by-step.
// See testdata/scenarios/README.md for more information.
func TestE2EScript(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("RUN_E2E not set; skipping")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	t.Run("scenarios", func(t *testing.T) {
		scenarioDir := "testdata/scenarios"
		scenarioPaths := findScripts(t, scenarioDir)

		for _, scenarioPath := range scenarioPaths {
			scenarioPath := scenarioPath

			t.Run(scenarioPath, func(t *testing.T) {
				uniqueID := testvariable.NewUniqueID()

				// Quickly load the sample with a dummy project, just to see if we should skip it
				{
					dummy := loadScript(t, filepath.Join(scenarioDir, scenarioPath), uniqueID, testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789})
					create.MaybeSkip(t, dummy.Name, dummy.Objects)
				}

				h := create.NewHarness(ctx, t)
				project := h.Project
				script := loadScript(t, filepath.Join(scenarioDir, scenarioPath), uniqueID, project)

				create.SetupNamespacesAndApplyDefaults(h, script.Objects, project)

				t.Cleanup(func() {
					create.DeleteResources(h, script.Objects)
				})

				var eventsByStep [][]*test.LogEntry

				eventsBefore := h.Events.HTTPEvents
				for i, obj := range script.Objects {

					testCommand := ""
					v, ok := obj.Object["TEST"]
					if ok {
						testCommand = v.(string)
					}
					if testCommand == "" {
						testCommand = "APPLY"
					}

					exportResource := obj.DeepCopy()
					shouldGetKubeObject := true

					switch testCommand {
					case "APPLY":
						applyObject(h, obj)
						create.WaitForReady(h, obj)

					case "DELETE":
						create.DeleteResources(h, []*unstructured.Unstructured{obj})
						exportResource = nil
						shouldGetKubeObject = false

					case "DELETE-NO-WAIT":
						create.DeleteResourceWithoutWaitingForReady(h, obj)

						// Allow some time for reconcile
						// Maybe we should instead wait for observedState
						time.Sleep(2 * time.Second)

						// The object probably still exists (that's probably
						// why we're using DELETE-NO-WAIT), so export the
						// resource and the kube object.

					case "ABANDON":
						setAnnotation(h, obj, "cnrm.cloud.google.com/deletion-policy", "abandon")
						create.DeleteResources(h, []*unstructured.Unstructured{obj})
						// continue to export the resource
						shouldGetKubeObject = false

					default:
						t.Errorf("unknown TEST command %q", testCommand)
						continue
					}

					if exportResource != nil {
						u := exportResourceAsUnstructured(h, exportResource)
						if u == nil {
							t.Logf("ignoring failure to export resource of gvk %v", exportResource.GroupVersionKind())
							// t.Errorf("failed to export resource of gvk %v", exportResource.GroupVersionKind())
						} else {
							if err := normalizeObject(u, project, uniqueID); err != nil {
								t.Fatalf("error from normalizeObject: %v", err)
							}
							got, err := yaml.Marshal(u)
							if err != nil {
								t.Errorf("failed to convert kube object to yaml: %v", err)
							}

							expectedPath := filepath.Join(script.SourceDir, fmt.Sprintf("_export%d.yaml", i))
							normalizers := []func(string) string{
								IgnoreComments,
							}
							h.CompareGoldenFile(expectedPath, string(got), normalizers...)
						}
					}

					if shouldGetKubeObject {
						u := &unstructured.Unstructured{}
						u.SetGroupVersionKind(obj.GroupVersionKind())
						id := types.NamespacedName{Namespace: obj.GetNamespace(), Name: obj.GetName()}
						if err := h.GetClient().Get(ctx, id, u); err != nil {
							t.Errorf("failed to get kube object: %v", err)
						} else {
							if err := normalizeObject(u, project, uniqueID); err != nil {
								t.Fatalf("error from normalizeObject: %v", err)
							}
							got, err := yaml.Marshal(u)
							if err != nil {
								t.Errorf("failed to convert kube object to yaml: %v", err)
							}
							expectedPath := filepath.Join(script.SourceDir, fmt.Sprintf("_object%02d.yaml", i))
							normalizers := []func(string) string{
								IgnoreComments,
							}
							h.CompareGoldenFile(expectedPath, string(got), normalizers...)
						}
					}

					var stepEvents []*test.LogEntry
					for i := len(eventsBefore); i < len(h.Events.HTTPEvents); i++ {
						stepEvents = append(stepEvents, h.Events.HTTPEvents[i])
					}
					eventsByStep = append(eventsByStep, stepEvents)
					eventsBefore = h.Events.HTTPEvents
				}

				shouldDumpHTTP := os.Getenv("GOLDEN_REQUEST_CHECKS") != ""
				if shouldDumpHTTP {
					x := NewNormalizer(uniqueID, project)

					for _, stepEvents := range eventsByStep {
						x.Preprocess(stepEvents)
					}

					for i, stepEvents := range eventsByStep {
						expectedPath := filepath.Join(script.SourceDir, fmt.Sprintf("_http%02d.log", i))
						got := x.Render(stepEvents)
						h.CompareGoldenFile(expectedPath, got, IgnoreComments)
					}

				}

				create.DeleteResources(h, script.Objects)

				h.NoExtraGoldenFiles(filepath.Join(script.SourceDir, "_*.yaml"))
			})
		}
	})
}

func applyObject(h *create.Harness, obj *unstructured.Unstructured) {
	if err := h.GetClient().Patch(h.Ctx, obj, client.Apply, client.FieldOwner("kcc-tests"), client.ForceOwnership); err != nil {
		h.Fatalf("error applying resource: %v", err)
	}
}

func setAnnotation(h *create.Harness, obj *unstructured.Unstructured, k, v string) {
	patch := &unstructured.Unstructured{}
	patch.SetGroupVersionKind(obj.GroupVersionKind())
	patch.SetNamespace(obj.GetNamespace())
	patch.SetName(obj.GetName())
	annotations := map[string]string{
		k: v,
	}
	patch.SetAnnotations(annotations)

	if err := h.GetClient().Patch(h.Ctx, patch, client.Apply, client.FieldOwner("kcc-tests"), client.ForceOwnership); err != nil {
		h.Fatalf("error setting annotations on resource: %v", err)
	}
}

func findScripts(t *testing.T, rootDir string) []string {
	var relPaths []string
	if err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == "script.yaml" {
			relPath, err := filepath.Rel(rootDir, filepath.Dir(path))
			if err != nil {
				return fmt.Errorf("getting relative path during directory walk: %w", err)
			}
			relPaths = append(relPaths, relPath)
		}
		return nil
	}); err != nil {
		t.Fatalf("error walking directory %q: %v", rootDir, err)
	}
	return relPaths
}

type Script struct {
	Name      string
	SourceDir string
	Objects   []*unstructured.Unstructured
}

func loadScript(t *testing.T, dir string, testID string, project testgcp.GCPProject) *Script {
	s := &Script{
		Name:      dir,
		SourceDir: dir,
	}
	b := test.MustReadFile(t, filepath.Join(dir, "script.yaml"))

	b = testcontroller.ReplaceTestVars(t, b, testID, project)

	// split into yaml objects
	yamls, err := kccyaml.SplitYAML(b)
	if err != nil {
		t.Fatalf("error splitting bytes into YAMLs: %v", err)
	}

	// Parse to objects
	var objects []*unstructured.Unstructured
	for _, y := range yamls {
		obj := &unstructured.Unstructured{}
		if err := yaml.Unmarshal(y, obj); err != nil {
			t.Fatalf("error parsing object: %v", err)
		}

		// Set namespace to match project
		if obj.GetNamespace() == "" {
			obj.SetNamespace(project.ProjectID)
		}

		// Hack: set project-id because mockkubeapiserver does not support webhooks
		if obj.GetAnnotations()["cnrm.cloud.google.com/project-id"] == "" {
			annotations := obj.GetAnnotations()
			if annotations == nil {
				annotations = make(map[string]string)
			}
			annotations["cnrm.cloud.google.com/project-id"] = project.ProjectID
			obj.SetAnnotations(annotations)
		}

		objects = append(objects, obj)
	}
	s.Objects = objects

	return s
}
