// Copyright 2022 Google LLC
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
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	opcorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestAllInSeries(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("RUN_E2E not set; skipping")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	t.Run("samples", func(t *testing.T) {
		samples := create.ListAllSamples(t)

		for _, sampleKey := range samples {
			sampleKey := sampleKey
			// TODO(b/259496928): Randomize the resource names for parallel execution when/if needed.

			t.Run(sampleKey.Name, func(t *testing.T) {
				// Quickly load the sample with a dummy project, just to see if we should skip it
				{
					dummySample := create.LoadSample(t, sampleKey, testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789})
					create.MaybeSkip(t, sampleKey.Name, dummySample.Resources)
				}

				h := create.NewHarness(ctx, t)
				project := h.Project
				s := create.LoadSample(t, sampleKey, project)

				create.SetupNamespacesAndApplyDefaults(h, s.Resources, project)

				// Hack: set project-id because mockkubeapiserver does not support webhooks
				for _, u := range s.Resources {
					annotations := u.GetAnnotations()
					if annotations == nil {
						annotations = make(map[string]string)
					}
					annotations["cnrm.cloud.google.com/project-id"] = project.ProjectID
					u.SetAnnotations(annotations)
				}

				create.RunCreateDeleteTest(h, create.CreateDeleteTestOptions{Create: s.Resources, CleanupResources: true})
			})
		}
	})

	testFixturesInSeries(ctx, t, false, cancel)
}

// TestPauseInSeries is a basic smoke test to prove that if CC pauses actuation of resources
// via the actuationMode field, then resources are not actuated onto the cloud provider.
// The current test is to make sure that POST requests are not recorded as HTTP events.
func TestPauseInSeries(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	testFixturesInSeries(ctx, t, true, cancel)
}

func testFixturesInSeries(ctx context.Context, t *testing.T, testPause bool, cancel context.CancelFunc) {
	t.Helper()

	if os.Getenv("RUN_E2E") == "" {
		t.Skip("RUN_E2E not set; skipping")
	}
	if testPause && os.Getenv("GOLDEN_REQUEST_CHECKS") == "" {
		t.Skip("GOLDEN_REQUEST_CHECKS not set; skipping as this test relies on the golden files.")
	}

	t.Run("fixtures", func(t *testing.T) {
		fixtures := resourcefixture.Load(t)
		for _, fixture := range fixtures {
			fixture := fixture
			// TODO(b/259496928): Randomize the resource names for parallel execution when/if needed.

			t.Run(fixture.Name, func(t *testing.T) {
				uniqueID := testvariable.NewUniqueID()

				loadFixture := func(project testgcp.GCPProject) (*unstructured.Unstructured, create.CreateDeleteTestOptions) {
					primaryResource := bytesToUnstructured(t, fixture.Create, uniqueID, project)

					opt := create.CreateDeleteTestOptions{CleanupResources: true}
					opt.Create = append(opt.Create, primaryResource)

					if fixture.Dependencies != nil {
						dependencyYamls := testyaml.SplitYAML(t, fixture.Dependencies)
						for _, dependBytes := range dependencyYamls {
							depUnstruct := bytesToUnstructured(t, dependBytes, uniqueID, project)
							opt.Create = append(opt.Create, depUnstruct)
						}
					}

					if fixture.Update != nil {
						u := bytesToUnstructured(t, fixture.Update, uniqueID, project)
						opt.Updates = append(opt.Updates, u)
					}
					return primaryResource, opt
				}

				// Quickly load the fixture with a dummy project, just to see if we should skip it
				{
					_, opt := loadFixture(testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789})
					create.MaybeSkip(t, fixture.Name, opt.Create)
				}

				h := create.NewHarness(ctx, t)
				project := h.Project

				if testPause {
					// we need to modify CC/ CCC state
					createPausedCC(ctx, t, h.GetClient())
				}

				primaryResource, opt := loadFixture(project)

				exportResources := []*unstructured.Unstructured{primaryResource}

				create.SetupNamespacesAndApplyDefaults(h, opt.Create, project)

				opt.CleanupResources = false // We delete explicitly below
				if testPause {
					opt.SkipWaitForReady = true // Paused resources don't send out an event yet.
				}
				create.RunCreateDeleteTest(h, opt)

				for _, exportResource := range exportResources {
					exportURI := ""

					gvk := exportResource.GroupVersionKind()
					switch gvk.GroupKind() {
					case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "Service"}:
						name := exportResource.GetName()
						projectID := project.ProjectID
						exportURI = "//serviceusage.googleapis.com/projects/" + projectID + "/services/" + name
					}

					if exportURI == "" {
						continue
					}

					exportParams := h.ExportParams()
					exportParams.IAMFormat = "partialpolicy"
					exportParams.ResourceFormat = "krm"
					outputDir := h.TempDir()
					outputPath := filepath.Join(outputDir, "export.yaml")
					exportParams.Output = outputPath
					exportParams.URI = exportURI
					if err := export.Execute(h.Ctx, &exportParams); err != nil {
						t.Errorf("error from export.Execute: %v", err)
						continue
					}

					expectedPath := filepath.Join(fixture.SourceDir, "export.yaml")
					output := h.MustReadFile(outputPath)
					h.CompareGoldenFile(expectedPath, string(output), h.IgnoreComments, h.ReplaceString(project.ProjectID, "example-project-id"))
				}

				if testPause {
					opt.SkipWaitForDelete = true
				}
				create.DeleteResources(h, opt)

				// Verify events against golden file
				if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" {
					events := h.Events

					operationIDs := map[string]bool{}
					pathIDs := map[string]string{}

					// Find "easy" operations and resources by looking for fully-qualified methods
					for _, event := range events.HTTPEvents {
						u := event.Request.URL
						if index := strings.Index(u, "?"); index != -1 {
							u = u[:index]
						}
						tokens := strings.Split(u, "/")
						n := len(tokens)
						if n >= 2 {
							kind := tokens[n-2]
							id := tokens[n-1]
							switch kind {
							case "tensorboards":
								pathIDs[id] = "${tensorboardID}"
							case "operations":
								operationIDs[id] = true
								pathIDs[id] = "${operationID}"
							}
						}
					}

					for _, event := range events.HTTPEvents {
						id := ""
						body := event.Response.ParseBody()
						val, ok := body["name"]
						if ok {
							s := val.(string)
							// operation name format: operations/{operationId}
							if strings.HasPrefix(s, "operations/") {
								id = strings.TrimPrefix(s, "operations/")
							}
							// operation name format: {prefix}/operations/{operationId}
							if ix := strings.Index(s, "/operations/"); ix != -1 {
								id = strings.TrimPrefix(s[ix:], "/operations/")
							}
							// operation name format: operation-{operationId}
							if strings.HasPrefix(s, "operation") {
								id = s
							}
						}
						if id != "" {
							operationIDs[id] = true
						}
					}

					for _, event := range events.HTTPEvents {
						if !strings.Contains(event.Request.URL, "/operations/${operationID}") {
							continue
						}
						responseBody := event.Response.ParseBody()
						if responseBody == nil {
							continue
						}
						name, _, _ := unstructured.NestedString(responseBody, "response", "name")
						if strings.HasPrefix(name, "tagKeys/") {
							pathIDs[name] = "tagKeys/${tagKeyID}"
						}
					}

					// Replace any dynamic IDs that appear in URLs
					for _, event := range events.HTTPEvents {
						url := event.Request.URL
						for k, v := range pathIDs {
							url = strings.ReplaceAll(url, "/"+k, "/"+v)
						}
						event.Request.URL = url
					}

					// Remove operation polling requests (ones where the operation is not ready)
					events.RemoveRequests(func(e *test.LogEntry) bool {
						if !strings.Contains(e.Request.URL, "/operations/${operationID}") {
							return false
						}
						responseBody := e.Response.ParseBody()
						if responseBody == nil {
							return false
						}
						if done, _, _ := unstructured.NestedBool(responseBody, "done"); done {
							return false
						}
						// remove if not done - and done can be omitted when false
						return true
					})

					jsonMutators := []test.JSONMutator{}
					addReplacement := func(path string, newValue string) {
						tokens := strings.Split(path, ".")
						jsonMutators = append(jsonMutators, func(obj map[string]any) {
							_, found, _ := unstructured.NestedString(obj, tokens...)
							if found {
								if err := unstructured.SetNestedField(obj, newValue, tokens...); err != nil {
									t.Fatal(err)
								}
							}
						})
					}

					addReplacement("id", "000000000000000000000")
					addReplacement("uniqueId", "111111111111111111111")
					addReplacement("oauth2ClientId", "888888888888888888888")

					addReplacement("etag", "abcdef0123A=")
					addReplacement("serviceAccount.etag", "abcdef0123A=")
					addReplacement("response.etag", "abcdef0123A=")

					addReplacement("createTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("response.createTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("creationTimestamp", "2024-04-01T12:34:56.123456Z")
					addReplacement("metadata.genericMetadata.createTime", "2024-04-01T12:34:56.123456Z")

					addReplacement("updateTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("response.updateTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("metadata.genericMetadata.updateTime", "2024-04-01T12:34:56.123456Z")

					// Specific to vertexai
					addReplacement("blobStoragePathPrefix", "cloud-ai-platform-00000000-1111-2222-3333-444444444444")
					addReplacement("response.blobStoragePathPrefix", "cloud-ai-platform-00000000-1111-2222-3333-444444444444")

					// Replace any empty values in LROs; this is surprisingly difficult to fix in mockgcp
					//
					//     "response": {
					// 	-    "@type": "type.googleapis.com/google.protobuf.Empty"
					// 	+    "@type": "type.googleapis.com/google.protobuf.Empty",
					// 	+    "value": {}
					// 	   }
					jsonMutators = append(jsonMutators, func(obj map[string]any) {
						response := obj["response"]
						if responseMap, ok := response.(map[string]any); ok {
							if responseMap["@type"] == "type.googleapis.com/google.protobuf.Empty" {
								value := responseMap["value"]
								if valueMap, ok := value.(map[string]any); ok && len(valueMap) == 0 {
									delete(responseMap, "value")
								}
							}
						}
					})

					events.PrettifyJSON(jsonMutators...)

					// Remove headers that just aren't very relevant to testing
					events.RemoveHTTPResponseHeader("Date")
					events.RemoveHTTPResponseHeader("Alt-Svc")

					got := events.FormatHTTP()
					expectedPath := filepath.Join(fixture.SourceDir, "_http.log")
					normalizers := []func(string) string{}
					normalizers = append(normalizers, h.IgnoreComments)
					normalizers = append(normalizers, h.ReplaceString(uniqueID, "${uniqueId}"))
					normalizers = append(normalizers, h.ReplaceString(project.ProjectID, "${projectId}"))
					normalizers = append(normalizers, h.ReplaceString(fmt.Sprintf("%d", project.ProjectNumber), "${projectNumber}"))
					for k, v := range pathIDs {
						normalizers = append(normalizers, h.ReplaceString(k, v))
					}
					for k := range operationIDs {
						normalizers = append(normalizers, h.ReplaceString(k, "${operationID}"))
					}

					if testPause {
						assertNoRequest(t, got, normalizers...)
					} else {
						h.CompareGoldenFile(expectedPath, got, normalizers...)
					}
				}
			})
		}
	})

	// Do a cleanup while we can still handle the error.
	t.Logf("shutting down manager")
	cancel()
}

// assertNoRequest checks that no POSTs or GETs are made against the cloud provider (GCP). This
// is helpful for when we want to test that Pause works correctly and doesn't actuate resources.
func assertNoRequest(t *testing.T, got string, normalizers ...func(s string) string) {
	t.Helper()

	for _, normalizer := range normalizers {
		got = normalizer(got)
	}

	if strings.Contains(got, "POST") {
		t.Fatalf("unexpected POST in log: %s", got)
	}

	if strings.Contains(got, "GET") {
		t.Fatalf("unexpected GET in log: %s", got)
	}
}

func bytesToUnstructured(t *testing.T, bytes []byte, testID string, project testgcp.GCPProject) *unstructured.Unstructured {
	t.Helper()
	updatedBytes := testcontroller.ReplaceTestVars(t, bytes, testID, project)
	return test.ToUnstructWithNamespace(t, updatedBytes, testID)
}

func createPausedCC(ctx context.Context, t *testing.T, c client.Client) {
	t.Helper()

	cc := &opcorev1beta1.ConfigConnector{}
	cc.Spec.Mode = "cluster"
	cc.Spec.Actuation = opcorev1beta1.Paused
	cc.Name = "configconnector.core.cnrm.cloud.google.com"

	if err := c.Create(ctx, cc); err != nil {
		t.Fatalf("error creating CC: %v", err)
	}
}
