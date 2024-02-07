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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	"hash/fnv"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	opcorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/sets"
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

				if os.Getenv("E2E_GCP_TARGET") == "vcr" {
					hash := func(s string) uint64 {
						h := fnv.New64a()
						h.Write([]byte(s))
						return h.Sum64()
					}
					uniqueID = strconv.FormatUint(hash(t.Name()), 36)
					// Stop recording after tests finish and write to cassette
					t.Cleanup(func() {
						err := h.VCRRecorder.Stop()
						if err != nil {
							t.Errorf("[VCR] Failed stop vcr recorder: %v", err)
						}
					})

					replaceWellKnownValues := func(s string) string {
						// Replace project id
						result := strings.Replace(s, project.ProjectID, "example-project", -1)

						// Replace user info
						obj := make(map[string]any)
						if err := json.Unmarshal([]byte(s), &obj); err == nil {
							toReplace, _, _ := unstructured.NestedString(obj, "user")
							if len(toReplace) != 0 {
								result = strings.Replace(result, toReplace, "user@google.com", -1)
							}
						}
						return result
					}

					unique := make(map[string]bool)

					hook := func(i *cassette.Interaction) error {
						// Discard failed interactions
						resCode := i.Response.Code
						if resCode == 404 || resCode == 400 || resCode == 403 {
							i.DiscardOnSave = true
						}
						// Discard repeated operation retry interactions
						reqURL := i.Request.URL
						resBody := i.Response.Body

						if strings.Contains(reqURL, "operations") {
							sorted, _ := sortJSON(resBody)
							if _, exists := unique[sorted]; !exists {
								unique[sorted] = true // Mark as seen
							} else {
								i.DiscardOnSave = true
							}
						}

						var requestHeadersToRemove = []string{
							"Authorization",
							"User-Agent",
						}
						for _, header := range requestHeadersToRemove {
							delete(i.Request.Headers, header)
						}

						var responseHeadersToRemove = []string{
							"Cache-Control",
							"Server",
							"Vary",
							"X-Content-Type-Options",
							"X-Frame-Options",
							"X-Xss-Protection",
							"Date",
							"Etag",
						}
						for _, header := range responseHeadersToRemove {
							delete(i.Response.Headers, header)
						}

						i.Request.Body = replaceWellKnownValues(i.Request.Body)
						i.Response.Body = replaceWellKnownValues(i.Response.Body)
						i.Request.URL = replaceWellKnownValues(i.Request.URL)

						return nil
					}
					h.VCRRecorder.AddHook(hook, recorder.BeforeSaveHook)

					h.VCRRecorder.SetMatcher(func(r *http.Request, i cassette.Request) bool {
						// We applied BeforeSaveHook, need to modify the incoming request,
						// so that incoming request matches the saved request.
						modifiedURL := replaceWellKnownValues(r.URL.String())

						if r.Method != i.Method || modifiedURL != i.URL {
							return false
						}

						// Default matcher only checks the request URL and Method. If request body exists, check the body as well.
						// This guarantees that the replayed response matches what the real service would return for that particular request.
						if r.Body != nil && r.Body != http.NoBody {
							var reqBody []byte
							var err error
							reqBody, err = io.ReadAll(r.Body)
							if err != nil {
								t.Fatal("[VCR] Failed to read request body")
							}
							r.Body.Close()
							r.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

							modifiedBody := replaceWellKnownValues(string(reqBody))
							if modifiedBody != i.Body {
								return false
							}
						}
						return true
					})
				}

				primaryResource, opt := loadFixture(project)

				exportResources := []*unstructured.Unstructured{primaryResource}

				create.SetupNamespacesAndApplyDefaults(h, opt.Create, project)

				opt.CleanupResources = false // We delete explicitly below
				if testPause {
					opt.SkipWaitForReady = true // Paused resources don't send out an event yet.
				}
				create.RunCreateDeleteTest(h, opt)

				for _, obj := range exportResources {
					got := exportResource(h, obj)
					if got != "" {
						expectedPath := filepath.Join(fixture.SourceDir, "export.yaml")
						h.CompareGoldenFile(expectedPath, string(got), IgnoreComments, ReplaceString(project.ProjectID, "example-project-id"))
					}
				}

				if testPause {
					opt.SkipWaitForDelete = true
				}
				create.DeleteResources(h, opt)

				// Verify kube events
				if h.KubeEvents != nil {
					verifyKubeWatches(h)
				}

				// Verify events against golden file
				if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" {
					events := test.LogEntries(h.Events.HTTPEvents)

					operationIDs := map[string]bool{}
					pathIDs := map[string]string{}

					// Find "easy" operations and resources by looking for fully-qualified methods
					for _, event := range events {
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

					for _, event := range events {
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

					for _, event := range events {
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
					for _, event := range events {
						url := event.Request.URL
						for k, v := range pathIDs {
							url = strings.ReplaceAll(url, "/"+k, "/"+v)
						}
						event.Request.URL = url
					}

					// Remove operation polling requests (ones where the operation is not ready)
					events = events.KeepIf(func(e *test.LogEntry) bool {
						if !strings.Contains(e.Request.URL, "/operations/${operationID}") {
							return true
						}
						responseBody := e.Response.ParseBody()
						if responseBody == nil {
							return true
						}
						if done, _, _ := unstructured.NestedBool(responseBody, "done"); done {
							return true
						}
						// remove if not done - and done can be omitted when false
						return false
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
					events.RemoveHTTPResponseHeader("Server-Timing")

					got := events.FormatHTTP()
					expectedPath := filepath.Join(fixture.SourceDir, "_http.log")
					normalizers := []func(string) string{}
					normalizers = append(normalizers, IgnoreComments)
					normalizers = append(normalizers, ReplaceString(uniqueID, "${uniqueId}"))
					normalizers = append(normalizers, ReplaceString(project.ProjectID, "${projectId}"))
					normalizers = append(normalizers, ReplaceString(fmt.Sprintf("%d", project.ProjectNumber), "${projectNumber}"))
					for k, v := range pathIDs {
						normalizers = append(normalizers, ReplaceString(k, v))
					}
					for k := range operationIDs {
						normalizers = append(normalizers, ReplaceString(k, "${operationID}"))
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

func verifyKubeWatches(h *create.Harness) {
	// Gather all the watch requests, using the Accept header to determine if it's a metadata-only watch.
	metadataWatches := sets.NewString()
	fullWatches := sets.NewString()
	objectWatches := sets.NewString()
	for _, event := range h.KubeEvents.HTTPEvents {
		if !strings.Contains(event.Request.URL, "watch=true") {
			continue
		}
		u, err := url.Parse(event.Request.URL)
		if err != nil {
			h.Fatalf("cannot parse url %q: %v", event.Request.URL, err)
		}

		metadataWatch := false
		acceptHeader := event.Request.Header.Get("Accept")
		if strings.Contains(acceptHeader, ";as=PartialObjectMetadata") {
			metadataWatch = true
		} else if acceptHeader == "application/json, */*" {
			metadataWatch = false
		} else if acceptHeader == "application/json" {
			metadataWatch = false
		} else if acceptHeader == "application/vnd.kubernetes.protobuf, */*" {
			metadataWatch = false
		} else if acceptHeader == "application/vnd.kubernetes.protobuf" {
			metadataWatch = false
		} else {
			h.Errorf("unhandled Accept header %q", acceptHeader)
		}

		fieldSelector := u.Query().Get("fieldSelector")
		if fieldSelector != "" {
			if strings.HasPrefix(fieldSelector, "metadata.name=") {
				objectName := strings.TrimPrefix(fieldSelector, "metadata.name=")
				objectWatches.Insert(u.Path + "/" + objectName)
				continue
			} else {
				h.Errorf("unhandled fieldSelector %q", fieldSelector)
			}
		}

		if metadataWatch {
			metadataWatches.Insert(u.Path)
		} else {
			fullWatches.Insert(u.Path)
		}
	}

	// Make sure we aren't opening both metadata-only watches and a full watch.
	// If we do this, we will have two caches, we'll get subtle race conditions
	// if we read from both of them.
	for metadataWatch := range metadataWatches {
		if fullWatches.Has(metadataWatch) {
			h.Errorf("two watches on %q (metadata and full); likely to cause race conditions", metadataWatch)
		}
	}

	// Validate the full watches we do have.
	// We only expect full watches on Namespaces, CRDs, CCs and CCCs (currently).
	allowedFullWatches := sets.NewString(
		"/apis/core.cnrm.cloud.google.com/v1beta1/configconnectorcontexts",
		"/apis/core.cnrm.cloud.google.com/v1beta1/configconnectors",
		"/apis/apiextensions.k8s.io/v1/customresourcedefinitions",
	)
	for fullWatch := range fullWatches {
		if !allowedFullWatches.Has(fullWatch) {
			h.Errorf("unexpected full watch on %q", fullWatch)
		}
	}
}

// JSON might be the same, but reordered. Try to sort it before comparing
func sortJSON(s string) (string, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		return "", err
	}
	sortedJSON, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(sortedJSON), nil
}
