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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/yaml"
)

// legacyScenarios lists all existing scenario suites that remain under the V1 runner.
// Note: "storagebucket" is omitted so it is tested by TestE2EScenariosV2 as the pilot suite.
var legacyScenarios = map[string]bool{
	"acquisition":                            true,
	"alloydbcluster":                         true,
	"alloydbinstance":                        true,
	"always-add-finalizers":                  true,
	"bigquerydatatransferconfig_duplication": true,
	"bigqueryreservationreservation":         true,
	"bigtableinstance":                       true,
	"cc_pause_change_reconcile":              true,
	"ccc_pause_change_reconcile":             true,
	"ccc_state_into_spec_absent":             true,
	"ccc_state_into_spec_merge":              true,
	"computefirewallpolicyrule":              true,
	"computetargettcpproxy":                  true,
	"containercluster":                       true,
	"fields":                                 true,
	"gkehubfeaturemembership":                true,
	"iam":                                    true,
	"iam_add_remove":                         true,
	"labels-apigateway":                      true,
	"powertool":                              true,
	"privilegedaccessmanagerentitlement":     true,
	"reconciliation_interval":                true,
	"secretmanagerversionalias":              true,
	"sql":                                    true,
	"sqlinstance-pointers-match":             true,
	"storageanywherecache":                   true,
	"tracker":                                true,
}

// TestE2EScenariosV2 runs a Scenario test that runs step-by-step.
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

					if obj.GroupVersionKind().Kind == "HTTPRequest" {
						// 1. Get properties
						verbVal, _, _ := unstructured.NestedString(obj.Object, "verb")
						if verbVal == "" {
							verbVal, _, _ = unstructured.NestedString(obj.Object, "method")
						}
						if verbVal == "" {
							verbVal = "GET"
						}
						verbVal = strings.ToUpper(verbVal)

						urlVal, _, _ := unstructured.NestedString(obj.Object, "url")
						if urlVal == "" {
							urlVal, _, _ = unstructured.NestedString(obj.Object, "path")
						}
						if urlVal == "" {
							h.Fatalf("HTTPRequest must specify a url or path")
						}

						// 2. Prepare body
						var reqBody io.Reader
						if bodyVal, ok := obj.Object["body"]; ok {
							switch val := bodyVal.(type) {
							case string:
								reqBody = strings.NewReader(val)
							default:
								bodyBytes, err := json.Marshal(val)
								if err != nil {
									h.Fatalf("failed to marshal HTTPRequest body to JSON: %v", err)
								}
								reqBody = bytes.NewReader(bodyBytes)
							}
						}

						// 3. Build request
						req, err := http.NewRequestWithContext(ctx, verbVal, urlVal, reqBody)
						if err != nil {
							h.Fatalf("failed to create http request: %v", err)
						}

						// 4. Headers
						if headersVal, ok := obj.Object["headers"]; ok {
							if headerMap, ok := headersVal.(map[string]any); ok {
								for k, v := range headerMap {
									req.Header.Set(k, fmt.Sprintf("%v", v))
								}
							}
						}
						if req.Header.Get("Content-Type") == "" && reqBody != nil {
							req.Header.Set("Content-Type", "application/json")
						}

						// 5. Send request and capture response
						resp, err := h.GCPHTTPClient().Do(req)
						if err != nil {
							h.Fatalf("HTTPRequest failed: %v", err)
						}
						defer resp.Body.Close()

						respBodyBytes, err := io.ReadAll(resp.Body)
						if err != nil {
							h.Fatalf("failed to read HTTPRequest response body: %v", err)
						}

						t.Logf("HTTPRequest response status: %s", resp.Status)

						// 6. Check for LRO polling
						if resp.StatusCode >= 200 && resp.StatusCode < 300 {
							var respJSON map[string]any
							if err := json.Unmarshal(respBodyBytes, &respJSON); err == nil {
								if isLRO(respJSON) {
									opName, _ := respJSON["name"].(string)
									if opName != "" {
										pollURL, err := buildPollURL(urlVal, opName)
										if err != nil {
											h.Fatalf("failed to build poll URL from %q and operation %q: %v", urlVal, opName, err)
										}

										t.Logf("Starting LRO polling for operation %q using URL %q", opName, pollURL)

										// Pause event sink logging during polling
										h.Events.Pause()

										pollErr := wait.PollImmediate(2*time.Second, 15*time.Minute, func() (bool, error) {
											pollReq, err := http.NewRequestWithContext(ctx, "GET", pollURL, nil)
											if err != nil {
												return false, err
											}
											pollResp, err := h.GCPHTTPClient().Do(pollReq)
											if err != nil {
												return false, nil // retry
											}
											defer pollResp.Body.Close()

											if pollResp.StatusCode < 200 || pollResp.StatusCode >= 300 {
												return false, nil // retry
											}

											pollBodyBytes, err := io.ReadAll(pollResp.Body)
											if err != nil {
												return false, nil // retry
											}

											var pollJSON map[string]any
											if err := json.Unmarshal(pollBodyBytes, &pollJSON); err != nil {
												return false, nil // retry
											}

											done, _ := pollJSON["done"].(bool)
											if done {
												if opErr, hasErr := pollJSON["error"]; hasErr {
													return true, fmt.Errorf("operation failed: %v", opErr)
												}
												t.Logf("LRO operation %q completed successfully", opName)
												return true, nil
											}
											return false, nil
										})

										// Resume event sink logging
										h.Events.Resume()

										if pollErr != nil {
											h.Fatalf("LRO polling failed: %v", pollErr)
										}
									}
								}
							}
						}

						captureHTTPLogEvents(false, deferHTTPLog)
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

							expectedPath := filepath.Join(script.SourceDir, fmt.Sprintf("_export%02d_mock.yaml", i))
							if targetGCP != "mock" {
								expectedPath = filepath.Join(script.SourceDir, fmt.Sprintf("_export%02d.yaml", i))
							} else {
								if os.Getenv("WRITE_GOLDEN_OUTPUT") == "" {
									if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
										expectedPath = filepath.Join(script.SourceDir, fmt.Sprintf("_export%02d.yaml", i))
									}
								}
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
							} else {
								if os.Getenv("WRITE_GOLDEN_OUTPUT") == "" {
									if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
										expectedPath = filepath.Join(script.SourceDir, fmt.Sprintf("_object%02d.yaml", i))
									}
								}
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
								} else {
									if _, err := os.Stat(wantPath); os.IsNotExist(err) {
										wantPath = filepath.Join(script.SourceDir, fmt.Sprintf("_object%02d.yaml", targetStepForReadAndCompare-1))
									}
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
						realExportPath := filepath.Join(script.SourceDir, fmt.Sprintf("_export%02d.yaml", i))
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
							var expectedPath string
							if targetGCP == "real" {
								expectedPath = filepath.Join(script.SourceDir, fmt.Sprintf("_http%02d.log", i))
							} else {
								expectedPath = filepath.Join(script.SourceDir, fmt.Sprintf("_http%02d_mock.log", i))
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

				// Diff Verification
				if targetGCP == "mock" {
					for i := range script.Objects {
						realHTTPPath := filepath.Join(script.SourceDir, fmt.Sprintf("_http%02d.log", i))
						mockHTTPPath := filepath.Join(script.SourceDir, fmt.Sprintf("_http%02d_mock.log", i))
						diffHTTPPath := filepath.Join(script.SourceDir, fmt.Sprintf("_http%02d_mock.diff", i))

						if fileExists(realHTTPPath) && fileExists(mockHTTPPath) {
							diff := computeDiff(ctx, realHTTPPath, mockHTTPPath)
							h.CompareGoldenFile(diffHTTPPath, diff)
						} else {
							h.AssertGoldenFileNotFound(diffHTTPPath)
						}

						realObjPath := filepath.Join(script.SourceDir, fmt.Sprintf("_object%02d.yaml", i))
						mockObjPath := filepath.Join(script.SourceDir, fmt.Sprintf("_object%02d_mock.yaml", i))
						diffObjPath := filepath.Join(script.SourceDir, fmt.Sprintf("_object%02d_mock.diff", i))

						if fileExists(realObjPath) && fileExists(mockObjPath) {
							diff := computeDiff(ctx, realObjPath, mockObjPath)
							h.CompareGoldenFile(diffObjPath, diff)
						} else {
							h.AssertGoldenFileNotFound(diffObjPath)
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

func computeDiff(ctx context.Context, oldP, newP string) string {
	var out bytes.Buffer

	cmd := exec.CommandContext(ctx, "diff", oldP, newP)
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) && exitError.ExitCode() == 1 {
			// This is expected when files differ
		} else {
			// Some errors might happen if diff is not installed or permissions,
			// but we try to do our best.
		}
	}

	return out.String()
}

func isLRO(jsonMap map[string]any) bool {
	name, ok := jsonMap["name"].(string)
	if !ok {
		return false
	}
	if strings.Contains(name, "/operations/") || strings.HasPrefix(name, "operations/") {
		return true
	}
	if _, ok := jsonMap["done"]; ok {
		return true
	}
	return false
}

func buildPollURL(originalURL, opName string) (string, error) {
	if strings.HasPrefix(opName, "http://") || strings.HasPrefix(opName, "https://") {
		return opName, nil
	}

	u, err := url.Parse(originalURL)
	if err != nil {
		return "", err
	}

	parts := strings.Split(u.Path, "/")
	var versionIndex = -1
	for i, part := range parts {
		if strings.HasPrefix(part, "v") && len(part) > 1 {
			if _, err := strconv.Atoi(part[1:2]); err == nil {
				versionIndex = i
				break
			}
		}
	}

	var basePath string
	if versionIndex != -1 {
		basePath = strings.Join(parts[:versionIndex+1], "/")
	} else {
		basePath = ""
	}

	u.Path = basePath + "/" + strings.TrimPrefix(opName, "/")
	u.RawQuery = ""
	return u.String(), nil
}
