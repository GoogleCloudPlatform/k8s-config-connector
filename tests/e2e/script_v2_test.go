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
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/yaml"
)

//go:embed testdata/legacy_scenarios.txt
var legacyScenariosRaw string

var legacyScenarios = func() map[string]bool {
	m := make(map[string]bool)
	lines := strings.Split(legacyScenariosRaw, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		m[line] = true
	}
	return m
}()

// TestE2EScenariosV2 runs a Scenario test that runs step-by-step.
// See testdata/scenarios/README.md for more information.
func TestE2EScenariosV2(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("RUN_E2E not set; skipping")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	logCheckTimeout := 10 * time.Second
	t.Run("scenarios", func(t *testing.T) {
		scenarioDir := "testdata/scenarios"
		scenarioPaths := findScripts(t, scenarioDir)

		for _, scenarioPath := range scenarioPaths {
			scenarioPath := scenarioPath

			t.Run(scenarioPath, func(t *testing.T) {
				if os.Getenv("SKIP_ALL") != "" {
					t.Skip("SKIP_ALL is set")
				}

				suite := strings.Split(scenarioPath, "/")[0]
				if legacyScenarios[suite] {
					t.Skip("legacy scenario, running under TestE2EScript")
				}

				uniqueID := testvariable.NewUniqueID()
				folderID := ""

				var harnessOptions []create.HarnessOption

				// Quickly load the sample with a dummy project, just to see if we should skip it
				{
					dummy := loadScript(t, filepath.Join(scenarioDir, scenarioPath), uniqueID, testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789})
					create.MaybeSkip(t, dummy.Name, dummy.Objects)

					// Record the CRDs we will use, for faster testing
					keepCRDs := map[schema.GroupKind]bool{}
					for _, obj := range dummy.Objects {
						keepCRDs[obj.GroupVersionKind().GroupKind()] = true
					}
					harnessOptions = append(harnessOptions, buildCRDFilter(keepCRDs))
				}

				h := create.NewHarness(ctx, t, harnessOptions...)
				project := h.Project
				script := loadScript(t, filepath.Join(scenarioDir, scenarioPath), uniqueID, project)

				create.SetupNamespacesAndApplyDefaults(h, script.Objects, project)

				var objectsToDelete []*unstructured.Unstructured
				t.Cleanup(func() {
					create.DeleteResources(h, create.CreateDeleteTestOptions{Create: objectsToDelete})
				})

				var eventsByStep []*SkippableLogEntries
				eventsBeforeCount := len(h.Events.GetHTTPEvents())
				captureHTTPLogEvents := func(skip bool, deferCapture bool) {
					var stepEvents []*test.LogEntry
					allEvents := h.Events.GetHTTPEvents()
					if !deferCapture {
						for i := eventsBeforeCount; i < len(allEvents); i++ {
							stepEvents = append(stepEvents, allEvents[i])
						}
					}
					eventsByStep = append(eventsByStep, &SkippableLogEntries{
						SkipCheck: skip || deferCapture,
						Entries:   stepEvents,
					})
					if !deferCapture {
						eventsBeforeCount = len(allEvents)
					}
				}

				// tracks all applied objects (in order, to avoid deletion dependency-ordering issues)
				appliedObjects := []*unstructured.Unstructured{}

				targetGCP := os.Getenv("E2E_GCP_TARGET")
				if targetGCP == "" {
					targetGCP = "mock"
				}

				for i, obj := range script.Objects {
					stepStart := time.Now()
					testCommand := ""
					v, ok := obj.Object["TEST"]
					if ok {
						testCommand = v.(string)
					}
					if testCommand == "" {
						testCommand = "APPLY"
					}

					t.Logf("***/Step %d: %s %s %s/%s", i, testCommand, obj.GroupVersionKind().Kind, obj.GetNamespace(), obj.GetName())

					deferHTTPLog := false
					v, ok = obj.Object["DEFER-HTTP-LOG"]
					if ok {
						deferHTTPLog = v.(bool)
					}

					if obj.GroupVersionKind().Kind == "RunCLI" {
						argsObjects := obj.Object["args"].([]any)
						var args []string
						for _, arg := range argsObjects {
							args = append(args, arg.(string))
						}
						baseOutputPath := filepath.Join(script.SourceDir, fmt.Sprintf("_cli-%d-", i))
						runCLI(h, args, uniqueID, baseOutputPath)
						captureHTTPLogEvents(true, deferHTTPLog)
						t.Logf("***/Step %d finished in %v", i, time.Since(stepStart))
						continue
					}

					if obj.GroupVersionKind().Kind == "MockGCPBackdoor" {
						if h.MockGCP != nil {
							service, _, _ := unstructured.NestedString(obj.Object, "service")
							verb, _, _ := unstructured.NestedString(obj.Object, "verb")

							if err := h.MockGCP.RunTestCommand(ctx, service, verb); err != nil {
								h.Fatalf("running test command: %v", err)
							}
						} else {
							t.Logf("skipping MockGCPBackdoor command, because not running against mockgcp")
						}

						captureHTTPLogEvents(false, deferHTTPLog)
						t.Logf("***/Step %d finished in %v", i, time.Since(stepStart))
						continue
					}

					// SystemRun lets KCC run for a while to observe/gather HTTP logs
					if obj.GroupVersionKind().Kind == "SystemRun" {
						var waitTimeout time.Duration
						if d, ok, _ := unstructured.NestedInt64(obj.Object, "duration"); ok {
							waitTimeout = time.Duration(d) * time.Second
						} else if d, ok, _ := unstructured.NestedFloat64(obj.Object, "duration"); ok {
							waitTimeout = time.Duration(d * float64(time.Second))
						}
						if durationString, ok, _ := unstructured.NestedString(obj.Object, "duration"); ok {
							d, err := time.ParseDuration(durationString)
							if err != nil {
								h.Fatalf("failed to parse duration %q: %v", durationString, err)
							}
							waitTimeout = d
						}
						if h.MockGCP != nil {
							timeoutChan := time.After(waitTimeout)
							ticker := time.NewTicker(30 * time.Second)

							for {
								stopWaiting := false
								select {
								case <-timeoutChan:
									t.Logf("finished waiting for http log collection")
									stopWaiting = true
									break
								case <-ticker.C:
									t.Logf("waiting for http log collection")
								}
								if stopWaiting {
									break
								}
							}
							ticker.Stop()

						} else {
							t.Logf("sleeping for %v for SystemRun", waitTimeout)
							time.Sleep(waitTimeout)
						}

						captureHTTPLogEvents(true, deferHTTPLog)
						t.Logf("***/Step %d finished in %v", i, time.Since(stepStart))
						continue
					}

					// Try to delete this object as part of cleanup
					objectsToDelete = append(objectsToDelete, obj)

					exportResource := obj.DeepCopy()
					shouldGetKubeObject := true
					v, ok = obj.Object["WRITE-KUBE-OBJECT"]
					if ok {
						shouldGetKubeObject = v.(bool)
					}

					var targetStepForReadAndCompare int
					switch testCommand {
					case "APPLY":
						applyObject(h, obj)
						create.WaitForReady(h, create.DefaultWaitForReadyTimeout, obj)
						appliedObjects = append(appliedObjects, obj)

					case "APPLY-10-SEC":
						applyObject(h, obj)
						time.Sleep(10 * time.Second)

					case "WAIT-FOR-OBSERVED-GENERATION":
						create.WaitForObservedGeneration(h, 5*time.Minute, obj)
						appliedObjects = append(appliedObjects, obj)

					case "WAIT-FOR-READY":
						create.WaitForReady(h, create.DefaultWaitForReadyTimeout, obj)
						appliedObjects = append(appliedObjects, obj)

					case "READ-OBJECT":
						appliedObjects = append(appliedObjects, obj)

					case "READ-OBJECT-AND-COMPARE-SPEC":
						v, ok := obj.Object["TARGET_STEP_FOR_READ_AND_COMPARE"]
						if !ok {
							t.Fatalf("did not find key TARGET_STEP_FOR_READ_AND_COMPARE in the READ-OBJECT-AND-COMPARE-SPEC step")
						}
						targetStepForReadAndCompare = int(v.(int64))
						if targetStepForReadAndCompare <= 0 {
							t.Fatalf("value of TARGET_STEP_FOR_READ_AND_COMPARE should be an integer > 0")
						}
						create.WaitForReady(h, create.DefaultWaitForReadyTimeout, obj)
						appliedObjects = append(appliedObjects, obj)

					case "APPLY-NO-WAIT":
						applyObject(h, obj)
						appliedObjects = append(appliedObjects, obj)
						exportResource = nil
						shouldGetKubeObject = false

					case "PATCH-EXTERNALLY-MANAGED-FIELDS":
						patchObjectWithExternallyManagedFields(h, obj)
						create.WaitForReady(h, create.DefaultWaitForReadyTimeout, obj)

					case "TOUCH":
						touchObject(h, obj)
						time.Sleep(2 * time.Second)

					case "DELETE":
						create.DeleteResources(h, create.CreateDeleteTestOptions{Create: []*unstructured.Unstructured{obj}})
						exportResource = nil
						shouldGetKubeObject = false

					case "SLEEP":
						time.Sleep(2 * time.Second)
						exportResource = nil
						shouldGetKubeObject = false

					case "DELETE-NO-WAIT":
						create.DeleteResources(h, create.CreateDeleteTestOptions{Create: []*unstructured.Unstructured{obj}, SkipWaitForDelete: true})
						time.Sleep(2 * time.Second)

					case "WAIT-FOR-HTTP-REQUEST":
						applyObject(h, obj)

						val, ok := obj.Object["VALUE_PRESENT"]
						if !ok {
							t.Fatalf("did not find key VALUE_PRESENT in step")
						}
						sval := val.(string)

						waitCtx, waitCancel := context.WithTimeout(ctx, logCheckTimeout)
						defer waitCancel()

						ticker := time.NewTicker(1 * time.Second)
						defer ticker.Stop()

						found := false
						for !found {
							select {
							case <-waitCtx.Done():
								t.Fatalf("timed out looking for value %s in http log: %v", sval, waitCtx.Err())
							case <-ticker.C:
								for _, l := range h.Events.GetHTTPEvents() {
									if strings.Contains(l.Response.Body, sval) {
										found = true
										break
									}
								}
							}
						}

					case "ABANDON":
						prePatchObj := readObject(h, obj.GroupVersionKind(), obj.GetNamespace(), obj.GetName())
						prePatchRV := prePatchObj.GetResourceVersion()

						setAnnotation(h, obj, "cnrm.cloud.google.com/deletion-policy", "abandon")
						waitForReconciliationAfterPatch(h, obj, prePatchRV)

						create.DeleteResources(h, create.CreateDeleteTestOptions{Create: []*unstructured.Unstructured{obj}})
						shouldGetKubeObject = false

					case "ABANDON-AND-REACQUIRE":
						existing := readObject(h, obj.GroupVersionKind(), obj.GetNamespace(), obj.GetName())
						resourceID, _, _ := unstructured.NestedString(existing.Object, "spec", "resourceID")
						if resourceID == "" {
							h.Fatalf("object did not have spec.resourceID: %v", existing)
						}
						setAnnotation(h, obj, "cnrm.cloud.google.com/deletion-policy", "abandon")
						deleteObj := obj.DeepCopy()
						create.DeleteResources(h, create.CreateDeleteTestOptions{Create: []*unstructured.Unstructured{deleteObj}})
						if err := unstructured.SetNestedField(obj.Object, resourceID, "spec", "resourceID"); err != nil {
							h.Fatalf("error setting spec.resourceID: %v", err)
						}
						applyObject(h, obj)
						create.WaitForReady(h, create.DefaultWaitForReadyTimeout, obj)
						appliedObjects = append(appliedObjects, obj)

					case "ABANDON-AND-REACQUIRE-WITH-GENERATED-ID":
						existing := readObject(h, obj.GroupVersionKind(), obj.GetNamespace(), obj.GetName())
						resourceID, _, _ := unstructured.NestedString(existing.Object, "spec", "resourceID")
						if resourceID == "" {
							externalRef, _, _ := unstructured.NestedString(existing.Object, "status", "externalRef")
							if externalRef == "" {
								h.Fatalf("object did not have spec.resourceID or status.externalRef: %v", existing.Object)
							}
							tokens := strings.Split(externalRef, "/")
							resourceID = tokens[len(tokens)-1]
						}

						deleteObj := &unstructured.Unstructured{}
						deleteObj.SetGroupVersionKind(existing.GroupVersionKind())
						deleteObj.SetNamespace(existing.GetNamespace())
						deleteObj.SetName(existing.GetName())
						deleteObj.SetAnnotations(existing.GetAnnotations())
						setAnnotation(h, deleteObj, "cnrm.cloud.google.com/deletion-policy", "abandon")
						create.DeleteResources(h, create.CreateDeleteTestOptions{Create: []*unstructured.Unstructured{deleteObj}})

						configuredID, _, _ := unstructured.NestedString(obj.Object, "spec", "resourceID")
						if configuredID == "" {
							h.Fatalf("object does not have resourceID configured: %v", obj)
						}
						err := unstructured.SetNestedField(obj.Object, strings.ReplaceAll(configuredID, "${TEST_GENERATED_ID}", resourceID), "spec", "resourceID")
						if err != nil {
							h.Fatalf("error setting spec.resourceID: %v", err)
						}
						applyObject(h, obj)
						create.WaitForReady(h, create.DefaultWaitForReadyTimeout, obj)
						appliedObjects = append(appliedObjects, obj)

					default:
						t.Errorf("FAIL: unknown TEST command %q", testCommand)
						t.Logf("***/Step %d finished in %v", i, time.Since(stepStart))
						continue
					}

					if exportResource != nil {
						u := exportResourceAsUnstructured(h, exportResource)
						if u == nil {
							t.Logf("ignoring failure to export resource of gvk %v", exportResource.GroupVersionKind())
						} else {
							normalizeKRMObject(t, u, project, folderID, uniqueID)
							got, err := yaml.Marshal(u)
							if err != nil {
								t.Errorf("failed to convert kube object to yaml: %v", err)
							}

							expectedPath := filepath.Join(script.SourceDir, fmt.Sprintf("_export%d_mock.yaml", i))
							if targetGCP != "mock" {
								expectedPath = filepath.Join(script.SourceDir, fmt.Sprintf("_export%d.yaml", i))
							}

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
							normalizeKRMObject(t, u, project, folderID, uniqueID)
							got, err := yaml.Marshal(u)
							if err != nil {
								t.Errorf("failed to convert kube object to yaml: %v", err)
							}

							expectedPath := filepath.Join(script.SourceDir, fmt.Sprintf("_object%02d_mock.yaml", i))
							if targetGCP != "mock" {
								expectedPath = filepath.Join(script.SourceDir, fmt.Sprintf("_object%02d.yaml", i))
							}

							normalizers := []func(string) string{
								IgnoreComments,
								IgnoreAnnotations(map[string]struct{}{
									"cnrm.cloud.google.com/mutable-but-unreadable-fields": {},
								}),
							}
							h.CompareGoldenFile(expectedPath, string(got), normalizers...)

							if targetStepForReadAndCompare > 0 {
								wantPath := filepath.Join(script.SourceDir, fmt.Sprintf("_object%02d_mock.yaml", targetStepForReadAndCompare-1))
								if targetGCP != "mock" {
									wantPath = filepath.Join(script.SourceDir, fmt.Sprintf("_object%02d.yaml", targetStepForReadAndCompare-1))
								}
								gotPath := expectedPath
								wantObj, err := getKubeObjectInStringFromFile(wantPath)
								if err != nil {
									h.Fatalf("%s", err.Error())
								}
								gotObj, err := getKubeObjectInStringFromFile(gotPath)
								if err != nil {
									h.Fatalf("%s", err.Error())
								}
								diff := getDiffInSpecs(wantObj, gotObj)
								if diff != "" {
									t.Errorf("unexpected diff when comparing with the kube object spec in step %v: %s", targetStepForReadAndCompare, diff)
								}
							}
						}
					}

					// Register real golden file as referenced if we are running in mock mode and using a mock-specific file
					if targetGCP == "mock" {
						realPath := filepath.Join(script.SourceDir, fmt.Sprintf("_object%02d.yaml", i))
						if fileExists(realPath) {
							b, _ := os.ReadFile(realPath)
							h.CompareGoldenFile(realPath, string(b))
						}
						realExportPath := filepath.Join(script.SourceDir, fmt.Sprintf("_export%d.yaml", i))
						if fileExists(realExportPath) {
							b, _ := os.ReadFile(realExportPath)
							h.CompareGoldenFile(realExportPath, string(b))
						}
					}

					captureHTTPLogEvents(false, deferHTTPLog)
					t.Logf("***/Step %d finished in %v", i, time.Since(stepStart))
				}

				t.Logf("***/Finished Steps")

				if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" || os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
					{
						x := NewNormalizer(uniqueID, project)

						for _, stepEvents := range eventsByStep {
							x.Preprocess(stepEvents.Entries)
						}

						for i, stepEvents := range eventsByStep {
							expectedPath := filepath.Join(script.SourceDir, fmt.Sprintf("_http%02d_mock.log", i))
							if targetGCP != "mock" {
								expectedPath = filepath.Join(script.SourceDir, fmt.Sprintf("_http%02d.log", i))
							}

							NormalizeHTTPLog(t, stepEvents.Entries, h.RegisteredServices(), project, uniqueID, "", "")
							got := x.Render(stepEvents.Entries)
							if stepEvents.SkipCheck {
								if os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
									if err := os.WriteFile(expectedPath, []byte(got), 0644); err != nil {
										t.Fatalf("FAIL: failed to write golden output %s: %v", expectedPath, err)
									}
									t.Logf("wrote updated golden output to %s", expectedPath)
								}
								continue
							}
							h.CompareGoldenFile(expectedPath, got, IgnoreComments)
						}
					}
				}

				{
					// Delete objects in order, de-duplicating objects we might have applied twice
					var objectsToDelete []*unstructured.Unstructured
					seen := make(map[string]bool)
					for _, obj := range appliedObjects {
						k := obj.GroupVersionKind().Group + "::" + obj.GroupVersionKind().Kind + "::" + obj.GetNamespace() + "::" + obj.GetName()
						if seen[k] {
							continue
						}
						objectsToDelete = append(objectsToDelete, obj)
						seen[k] = true
					}
					create.DeleteResources(h, create.CreateDeleteTestOptions{Create: objectsToDelete})
				}

				h.NoExtraGoldenFiles(filepath.Join(script.SourceDir, "_*.yaml"))
			})
		}
	})
}

func fileExists(p string) bool {
	if _, err := os.Stat(p); err != nil {
		return false
	}
	return true
}
