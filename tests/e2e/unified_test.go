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
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	opcorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
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

	subtestTimeout := time.Hour
	if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "mock" {
		// We allow a total of 3 minutes: 2 for the test itself (for deep object chains with retries),
		// and 1 minute to shutdown envtest / allow kube-apiserver requests to time-out.
		subtestTimeout = 3 * time.Minute
	}

	t.Run("samples", func(t *testing.T) {
		samples := create.ListAllSamples(t)

		for _, sampleKey := range samples {
			sampleKey := sampleKey
			// TODO(b/259496928): Randomize the resource names for parallel execution when/if needed.

			t.Run(sampleKey.Name, func(t *testing.T) {
				ctx := addTestTimeout(ctx, t, subtestTimeout)
				var harnessOptions []create.HarnessOption

				// Quickly load the sample with a dummy project, just to see if we should skip it
				{
					dummySample := create.LoadSample(t, sampleKey, testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789})
					create.MaybeSkip(t, sampleKey.Name, dummySample.Resources)
					if s := os.Getenv("ONLY_TEST_APIGROUPS"); s != "" {
						t.Skipf("skipping test because cannot determine group for samples, with ONLY_TEST_APIGROUPS=%s", s)
					}

					// Record the CRDs we will use, for faster testing
					keepCRDs := map[schema.GroupKind]bool{}
					for _, obj := range dummySample.Resources {
						keepCRDs[obj.GroupVersionKind().GroupKind()] = true
					}
					harnessOptions = append(harnessOptions, buildCRDFilter(keepCRDs))

				}

				h := create.NewHarness(ctx, t, harnessOptions...)
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

	subtestTimeout := time.Hour
	if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "mock" {
		// We allow a total of 3 minutes: 2 for the test itself (for deep object chains with retries),
		// and 1 minute to shutdown envtest / allow kube-apiserver requests to time-out.
		subtestTimeout = 3 * time.Minute
	}
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
			group := fixture.GVK.Group
			if s := os.Getenv("SKIP_TEST_APIGROUP"); s != "" {
				skippedGroups := strings.Split(s, ",")
				if slice.StringSliceContains(skippedGroups, group) {
					klog.Infof("skipping test %s because group %q matched entries in SKIP_TEST_APIGROUP=%s", fixture.Name, group, s)
					continue
				}
			}
			if s := os.Getenv("ONLY_TEST_APIGROUPS"); s != "" {
				groups := strings.Split(s, ",")
				if !slice.StringSliceContains(groups, group) {
					klog.Infof("skipping test %s because group %q did not match ONLY_TEST_APIGROUPS=%s", fixture.Name, group, s)
					continue
				}
			}
			// TODO(b/259496928): Randomize the resource names for parallel execution when/if needed.
			t.Run(fixture.Name, func(t *testing.T) {
				ctx := addTestTimeout(ctx, t, subtestTimeout)

				loadFixture := func(project testgcp.GCPProject, uniqueID string) (*unstructured.Unstructured, create.CreateDeleteTestOptions) {
					primaryResource := bytesToUnstructured(t, fixture.Create, uniqueID, project)

					opt := create.CreateDeleteTestOptions{CleanupResources: true}

					if fixture.Dependencies != nil {
						dependencyYamls := testyaml.SplitYAML(t, fixture.Dependencies)
						for _, dependBytes := range dependencyYamls {
							depUnstruct := bytesToUnstructured(t, dependBytes, uniqueID, project)
							opt.Create = append(opt.Create, depUnstruct)
						}
					}

					opt.Create = append(opt.Create, primaryResource)

					if fixture.Update != nil {
						u := bytesToUnstructured(t, fixture.Update, uniqueID, project)
						opt.Updates = append(opt.Updates, u)
					}
					return primaryResource, opt
				}

				runScenario(ctx, t, testPause, fixture, loadFixture)
			})
		}
	})

	// Do a cleanup while we can still handle the error.
	t.Logf("shutting down manager")
	cancel()
}

func runScenario(ctx context.Context, t *testing.T, testPause bool, fixture resourcefixture.ResourceFixture, loadFixture func(project testgcp.GCPProject, uniqueID string) (*unstructured.Unstructured, create.CreateDeleteTestOptions)) {
	var harnessOptions []create.HarnessOption

	// Extra indentation to avoid merge conflicts
	{
		{
			{
				uniqueID := testvariable.NewUniqueID()

				// Quickly load the fixture with a dummy project, just to see if we should skip it
				{
					_, opt := loadFixture(testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789}, uniqueID)
					create.MaybeSkip(t, fixture.Name, opt.Create)
					if testPause && containsCCOrCCC(opt.Create) {
						t.Skipf("test case %q contains ConfigConnector or ConfigConnectorContext object(s): "+
							"pause test should not run against test cases already contain ConfigConnector "+
							"or ConfigConnectorContext objects", fixture.Name)
					}

					// If the test contains "${resourceId}", that means it is an acquisition test, which we don't currently support
					for _, create := range opt.Create {
						resourceID, _, err := unstructured.NestedString(create.Object, "spec", "resourceID")
						if err != nil {
							j, _ := json.Marshal(create.Object)
							t.Logf("error reading spec.resourceID, can't check for acquisition test: %v.  object is %v", err, string(j))
						} else if strings.Contains(resourceID, "${resourceId}") {
							t.Skipf("test has ${resourceId} placeholder in spec.resource, indicating an acquisition test.  Not currently supported here; skipping")
						}
					}

					// Record the CRDs we will use, for faster testing
					keepCRDs := map[schema.GroupKind]bool{}
					for _, obj := range opt.Create {
						keepCRDs[obj.GroupVersionKind().GroupKind()] = true
					}
					for _, obj := range opt.Updates {
						keepCRDs[obj.GroupVersionKind().GroupKind()] = true
					}
					harnessOptions = append(harnessOptions, buildCRDFilter(keepCRDs))
				}

				// Create test harness
				var h *create.Harness
				if os.Getenv("E2E_GCP_TARGET") == "vcr" {
					harnessOptions = append(harnessOptions, create.WithVCRPath(fixture.SourceDir))
					h = create.NewHarness(ctx, t, harnessOptions...)
					hash := func(s string) uint64 {
						h := fnv.New64a()
						h.Write([]byte(s))
						return h.Sum64()
					}
					uniqueID = strconv.FormatUint(hash(t.Name()), 36)
					// Stop recording after tests finish and write to cassette
					t.Cleanup(func() {
						err := h.VCRRecorderNonTF.Stop()
						if err != nil {
							t.Errorf("FAIL: [VCR] Failed stop non TF vcr recorder: %v", err)
						}
						err = h.VCRRecorderTF.Stop()
						if err != nil {
							t.Errorf("FAIL: [VCR] Failed stop TF vcr recorder: %v", err)
						}
						err = h.VCRRecorderOauth.Stop()
						if err != nil {
							t.Errorf("FAIL: [VCR] Failed stop Oauth vcr recorder: %v", err)
						}
					})
					configureVCR(t, h)
				} else {
					h = create.NewHarness(ctx, t, harnessOptions...)
				}
				project := h.Project

				if testPause {
					// we need to modify CC/ CCC state
					createPausedCC(ctx, t, h.GetClient())
				}

				primaryResource, opt := loadFixture(project, uniqueID)

				exportResources := []*unstructured.Unstructured{primaryResource}

				create.SetupNamespacesAndApplyDefaults(h, opt.Create, project)

				opt.CleanupResources = false // We delete explicitly below
				if testPause {
					opt.SkipWaitForReady = true // Paused resources don't send out an event yet.
				}
				if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" {
					// If we're doing golden request checks, create synchronously so that it is reproducible.
					// Note that this does introduce a dependency that objects are ordered correctly for creation.
					opt.CreateInOrder = true
				}
				create.RunCreateDeleteTest(h, opt)

				if os.Getenv("GOLDEN_OBJECT_CHECKS") != "" || os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
					for _, obj := range exportResources {
						// Get testName from t.Name()
						// If t.Name() = TestAllInInSeries_fixtures_computenodetemplate
						// the testName should be computenodetemplate
						pieces := strings.Split(t.Name(), "/")
						var testName string
						if len(pieces) > 0 {
							testName = pieces[len(pieces)-1]
						} else {
							t.Fatalf("FAIL: failed to get test name")
						}
						// Golden test exported GCP object
						exportedYAML := exportResource(h, obj, &Expectations{})
						if exportedYAML != "" {
							exportedObj := &unstructured.Unstructured{}
							if err := yaml.Unmarshal([]byte(exportedYAML), exportedObj); err != nil {
								t.Fatalf("FAIL: error from yaml.Unmarshal: %v", err)
							}
							if err := normalizeKRMObject(t, exportedObj, project, uniqueID); err != nil {
								t.Fatalf("FAIL: error from normalizeObject: %v", err)
							}
							got, err := yaml.Marshal(exportedObj)
							if err != nil {
								t.Fatalf("FAIL: failed to convert KRM object to yaml: %v", err)
							}

							expectedPath := filepath.Join(fixture.SourceDir, fmt.Sprintf("_generated_export_%v.golden", testName))
							h.CompareGoldenFile(expectedPath, string(got), IgnoreComments)
						}
						// Golden test created KRM object
						u := &unstructured.Unstructured{}
						u.SetGroupVersionKind(obj.GroupVersionKind())
						id := types.NamespacedName{Namespace: obj.GetNamespace(), Name: obj.GetName()}
						if err := h.GetClient().Get(ctx, id, u); err != nil {
							t.Fatalf("FAIL: failed to get KRM object: %v", err)
						} else {
							if err := normalizeKRMObject(t, u, project, uniqueID); err != nil {
								t.Fatalf("FAIL: error from normalizeObject: %v", err)
							}
							got, err := yaml.Marshal(u)
							if err != nil {
								t.Fatalf("FAIL: failed to convert KRM object to yaml: %v", err)
							}
							expectedPath := filepath.Join(fixture.SourceDir, fmt.Sprintf("_generated_object_%v.golden.yaml", testName))
							test.CompareGoldenObject(t, expectedPath, got)
						}
					}
				}

				if testPause {
					opt.SkipWaitForDelete = true
				}
				if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" {
					// If we're doing golden request checks, delete synchronously so that it is reproducible.
					// Note that this does introduce a dependency that objects are ordered correctly for deletion.
					opt.DeleteInOrder = true
				}
				create.DeleteResources(h, opt)

				// Verify kube events
				if h.KubeEvents != nil {
					verifyKubeWatches(h)
				}

				// Verify HTTP log with static checks
				verifyUserAgent(h)

				// Verify events against golden file or records events
				if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" || os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
					events := test.LogEntries(h.Events.HTTPEvents)

					r := NewReplacements()

					// Find "easy" operations and resources by looking for fully-qualified methods
					for _, event := range events {
						u := event.Request.URL
						if index := strings.Index(u, "?"); index != -1 {
							u = u[:index]
						}
						r.ExtractIDsFromLinks(u)
					}

					for _, event := range events {
						id := ""
						body := event.Response.ParseBody()
						val, ok := body["name"]
						if ok {
							s := val.(string)
							tokens := strings.Split(s, "/")
							// operation name format: operations/{operationId}
							if len(tokens) == 2 && tokens[0] == "operations" {
								id = strings.TrimPrefix(s, "operations/")
							}
							// operation name format: {prefix}/operations/{operationId}
							if len(tokens) > 2 && tokens[len(tokens)-2] == "operations" {
								id = tokens[len(tokens)-1]
							}
							// operation name format: operation-{operationId}
							if len(tokens) == 1 && strings.HasPrefix(tokens[0], "operation") {
								id = s
							}
							// SQL operations require a special case.
							if kind, ok := body["kind"]; ok && kind == "sql#operation" {
								id = s
							}
						}
						if id != "" {
							r.OperationIDs[id] = true
						}
					}

					for _, event := range events {
						body := event.Response.ParseBody()
						if selfLinkWithId, _, _ := unstructured.NestedString(body, "selfLinkWithId"); selfLinkWithId != "" {
							r.ExtractIDsFromLinks(selfLinkWithId)
						}

						if billingAccountName, _, _ := unstructured.NestedString(body, "billingAccountName"); billingAccountName != "" {
							r.ExtractIDsFromLinks(billingAccountName)
						}

						// if targetId, _, _ := unstructured.NestedString(body, "targetId"); targetId != "" {
						// 	extractIDsFromLinks(selfLinkWithId)
						// }

						if conditions, _, _ := unstructured.NestedSlice(body, "conditions"); conditions != nil {
							for _, conditionAny := range conditions {
								condition := conditionAny.(map[string]any)
								name, _, _ := unstructured.NestedString(condition, "name")
								if name != "" {
									r.ExtractIDsFromLinks(name)
								}
							}
						}

						if val, ok := body["projectNumber"]; ok {
							s := val.(string)
							r.PathIDs[s] = "${projectNumber}"
						}
					}

					// Replace any operation IDs that appear in URLs
					for _, event := range events {
						u := event.Request.URL
						for operationID := range r.OperationIDs {
							u = strings.ReplaceAll(u, operationID, "${operationID}")
						}
						event.Request.URL = u
					}

					for _, event := range events {
						if !isGetOperation(event) {
							continue
						}
						responseBody := event.Response.ParseBody()
						if responseBody == nil {
							continue
						}
						if name, _, _ := unstructured.NestedString(responseBody, "response", "name"); name != "" {
							r.ExtractIDsFromLinks(name)
						}
						if targetLink, _, _ := unstructured.NestedString(responseBody, "targetLink"); targetLink != "" {
							r.ExtractIDsFromLinks(targetLink)
						}
					}

					// Replace any dynamic IDs that appear in URLs
					for _, event := range events {
						u := event.Request.URL
						for k, v := range r.PathIDs {
							u = strings.ReplaceAll(u, "/"+k, "/"+v)
						}
						event.Request.URL = u
					}

					// Remove operation polling requests (ones where the operation is not ready)
					events = events.KeepIf(func(e *test.LogEntry) bool {
						if !isGetOperation(e) {
							return true
						}
						responseBody := e.Response.ParseBody()
						if responseBody == nil {
							return true
						}
						if done, _, _ := unstructured.NestedBool(responseBody, "done"); done {
							return true
						}
						if status, _, _ := unstructured.NestedString(responseBody, "status"); status == "DONE" {
							return true
						}
						// remove if not done - and done can be omitted when false
						return false
					})

					jsonMutators := []test.JSONMutator{}
					addReplacement := func(path string, newValue string) {
						tokens := strings.Split(path, ".")
						jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
							_, found, _ := unstructured.NestedString(obj, tokens...)
							if found {
								if err := unstructured.SetNestedField(obj, newValue, tokens...); err != nil {
									t.Fatalf("FAIL: setting nested field: %v", err)
								}
							}
						})
					}

					addSetStringReplacement := func(path string, newValue string) {
						jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
							if err := setStringAtPath(obj, path, newValue); err != nil {
								t.Fatalf("FAIL: error from setStringAtPath(%+v): %v", obj, err)
							}
						})
					}

					addReplacement("id", "000000000000000000000")
					addReplacement("uniqueId", "111111111111111111111")
					addReplacement("oauth2ClientId", "888888888888888888888")
					addReplacement("response.oauth2ClientId", "888888888888888888888")

					addReplacement("createTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("expireTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("response.createTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("response.deleteTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("creationTimestamp", "2024-04-01T12:34:56.123456Z")
					addReplacement("metadata.createTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("metadata.genericMetadata.createTime", "2024-04-01T12:34:56.123456Z")
					addSetStringReplacement(".monitoredProjects[].createTime", "2024-04-01T12:34:56.123456Z")

					addReplacement("updateTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("response.updateTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("metadata.genericMetadata.updateTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("metadata.updateTime", "2024-04-01T12:34:56.123456Z")

					// Specific to cloudbuild
					addReplacement("metadata.completeTime", "2024-04-01T12:34:56.123456Z")

					// Specific to spanner
					addReplacement("metadata.startTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("metadata.endTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("metadata.instance.createTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("metadata.instance.updateTime", "2024-04-01T12:34:56.123456Z")

					// Specific to spanner database
					addReplacement("earliestVersionTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("response.earliestVersionTime", "2024-04-01T12:34:56.123456Z")
					addSetStringReplacement(".metadata.progress[].startTime", "2024-04-01T12:34:56.123456Z")
					addSetStringReplacement(".metadata.progress[].endTime", "2024-04-01T12:34:56.123456Z")
					addSetStringReplacement(".metadata.commitTimestamps[]", "2024-04-01T12:34:56.123456Z")

					// Specific to redis
					addReplacement("metadata.createTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("metadata.endTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("response.host", "10.1.2.3")
					addReplacement("response.reservedIpRange", "10.1.2.0/24")
					addReplacement("host", "10.1.2.3")
					addReplacement("reservedIpRange", "10.1.2.0/24")
					addReplacement("metadata.endTime", "2024-04-01T12:34:56.123456Z")

					// Specific to Compute
					addReplacement("natIP", "192.0.0.10")
					addReplacement("fingerprint", "abcdef0123A=")
					// Matches the mock ip address of Compute forwarding rule
					addReplacement("IPAddress", "8.8.8.8")
					addReplacement("pscConnectionId", "111111111111")

					// Specific to IAM/policy
					addReplacement("policy.etag", "abcdef0123A=")

					// Specific to vertexai
					addReplacement("blobStoragePathPrefix", "cloud-ai-platform-00000000-1111-2222-3333-444444444444")
					addReplacement("response.blobStoragePathPrefix", "cloud-ai-platform-00000000-1111-2222-3333-444444444444")
					addReplacement("state.diskUtilizationBytes", "1")
					for _, event := range events {
						responseBody := event.Response.ParseBody()
						if responseBody == nil {
							continue
						}
						metadataArtifact, _, _ := unstructured.NestedString(responseBody, "metadataArtifact")
						if metadataArtifact != "" {
							tokens := strings.Split(metadataArtifact, "/")
							n := len(tokens)
							if n >= 2 {
								kind := tokens[n-2]
								id := tokens[n-1]
								switch kind {
								case "artifacts":
									r.PathIDs[id] = "${artifactId}"
								}
							}
						}
						gcsBucket, _, _ := unstructured.NestedString(responseBody, "metadata", "gcsBucket")
						if gcsBucket != "" && strings.HasPrefix(gcsBucket, "cloud-ai-platform-") {
							r.PathIDs[gcsBucket] = "cloud-ai-platform-${bucketId}"
						}
					}

					// Specific to AlloyDB
					addReplacement("uid", "111111111111111111111")
					addReplacement("response.uid", "111111111111111111111")
					addReplacement("continuousBackupInfo.enabledTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("response.continuousBackupInfo.enabledTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("ipAddress", "10.1.2.3")
					addReplacement("response.ipAddress", "10.1.2.3")
					addReplacement("primary.createTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("primary.generateTime", "2024-04-01T12:34:56.123456Z")
					jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
						if val, found, _ := unstructured.NestedString(obj, "name"); found {
							if strings.Contains(val, "clusters/alloydb") ||
								strings.Contains(val, "instances/alloydb") ||
								strings.Contains(val, "backups/alloydb") {

								// Explicitly set `reconciling` to `false`.
								if _, found, _ := unstructured.NestedBool(obj, "reconciling"); !found {
									if err := unstructured.SetNestedField(obj, false, "reconciling"); err != nil {
										t.Fatal(err)
									}
								}

								// Replace the IP addresses in `outboundPublicIpAddresses` slice to test IP addresses.
								if _, found, _ := unstructured.NestedSlice(obj, "outboundPublicIpAddresses"); found {
									if err := unstructured.SetNestedStringSlice(obj, []string{"6.6.6.6", "8.8.8.8"}, "outboundPublicIpAddresses"); err != nil {
										t.Fatal(err)
									}
								}
							}
						}
					})
					// Boolean fields in LRO are omitted when false so we need
					// to add them back.
					jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
						if _, found, _ := unstructured.NestedMap(obj, "metadata"); found {
							if val, found, err := unstructured.NestedString(obj, "metadata", "@type"); err == nil && found && val == "type.googleapis.com/google.cloud.alloydb.v1beta.OperationMetadata" {
								if _, found, err := unstructured.NestedString(obj, "done"); err == nil && !found {
									// Explicitly set `done` to `false`.
									if err := unstructured.SetNestedField(obj, false, "done"); err != nil {
										t.Fatal(err)
									}
								}

								if _, found, err := unstructured.NestedString(obj, "metadata", "requestedCancellation"); err == nil && !found {
									// Explicitly set `metadata.requestedCancellation` to `false`.
									if err := unstructured.SetNestedField(obj, false, "metadata", "requestedCancellation"); err != nil {
										t.Fatal(err)
									}
								}

								if _, found, _ := unstructured.NestedMap(obj, "response"); found {
									if val, found, _ := unstructured.NestedString(obj, "response", "@type"); found &&
										val == "type.googleapis.com/google.cloud.alloydb.v1beta.Cluster" ||
										val == "type.googleapis.com/google.cloud.alloydb.v1beta.Instance" ||
										val == "type.googleapis.com/google.cloud.alloydb.v1beta.Backup" {
										// Explicitly set `reconciling` in response to `false`.
										if _, found, _ := unstructured.NestedBool(obj, "response", "reconciling"); !found {
											if err := unstructured.SetNestedField(obj, false, "response", "reconciling"); err != nil {
												t.Fatal(err)
											}
										}

										// Replace the IP addresses in `outboundPublicIpAddresses` slice to test IP addresses.
										if _, found, _ := unstructured.NestedSlice(obj, "response", "outboundPublicIpAddresses"); found {
											if err := unstructured.SetNestedStringSlice(obj, []string{"6.6.6.6", "8.8.8.8"}, "response", "outboundPublicIpAddresses"); err != nil {
												t.Fatal(err)
											}
										}
									}
								}
							}
						}
					})
					// Specific to BigQuery
					addSetStringReplacement(".access[].userByEmail", "user@google.com")

					// Specific to BigTable
					addSetStringReplacement(".instances[].createTime", "2024-04-01T12:34:56.123456Z")
					addSetStringReplacement(".metadata.requestTime", "2024-04-01T12:34:56.123456Z")
					addSetStringReplacement(".metadata.finishTime", "2024-04-01T12:34:56.123456Z")

					// Specific to Firestore
					jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
						if _, found, _ := unstructured.NestedMap(obj, "response"); found {
							// Only run this mutator for firestore database objects.
							if val, found, err := unstructured.NestedString(obj, "response", "@type"); err == nil && found && val == "type.googleapis.com/google.firestore.admin.v1.Database" {
								// Only run this mutator for firestore database objects that have a name set in the response.
								if val, found, err := unstructured.NestedString(obj, "response", "name"); err == nil && found && val != "" {
									// Set name field to use human-readable ID, instead of UID
									// Note: This only works if firestore databases in all resource fixture test cases use the name "firestoredatabase-${uniqueId}"
									if err := unstructured.SetNestedField(obj, "projects/${projectId}/databases/firestoredatabase-${uniqueId}", "response", "name"); err != nil {
										t.Fatalf("FAIL: stting nested field: %v", err)
									}
								}
							}
						}
					})

					// Specific to PAM
					// Boolean fields in LRO are omitted when false so we need
					// to add them back.
					jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
						if _, found, _ := unstructured.NestedMap(obj, "metadata"); found {
							if val, found, err := unstructured.NestedString(obj, "metadata", "@type"); err == nil && found && val == "type.googleapis.com/google.cloud.privilegedaccessmanager.v1.OperationMetadata" {
								if _, found, err := unstructured.NestedString(obj, "done"); err == nil && !found {
									// Explicitly set `done` to `false`.
									if err := unstructured.SetNestedField(obj, false, "done"); err != nil {
										t.Fatalf("FAIL: setting nested field: %v", err)
									}
								}

								if _, found, err := unstructured.NestedString(obj, "metadata", "requestedCancellation"); err == nil && !found {
									// Explicitly set `metadata.requestedCancellation` to `false`.
									if err := unstructured.SetNestedField(obj, false, "metadata", "requestedCancellation"); err != nil {
										t.Fatalf("FAIL: setting nested field: %v", err)
									}
								}
							}

						}
					})

					// Specific to pubsub
					addReplacement("revisionCreateTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("revisionId", "revision-id-placeholder")

					// Specific to monitoring
					addSetStringReplacement(".creationRecord.mutateTime", "2024-04-01T12:34:56.123456Z")
					addSetStringReplacement(".creationRecord.mutatedBy", "user@example.com")
					addSetStringReplacement(".mutationRecord.mutateTime", "2024-04-01T12:34:56.123456Z")
					addSetStringReplacement(".mutationRecord.mutatedBy", "user@example.com")
					addSetStringReplacement(".mutationRecords[].mutateTime", "2024-04-01T12:34:56.123456Z")
					addSetStringReplacement(".mutationRecords[].mutatedBy", "user@example.com")

					// Specific to CertificateManager
					addReplacement("response.dnsResourceRecord.data", uniqueID)
					jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
						if val, found, err := unstructured.NestedString(obj, "kind"); err != nil || !found || val != "sql#instance" {
							// Only run this mutator for sql instance objects.
							return
						}
						if _, found, _ := unstructured.NestedString(obj, "state"); !found {
							// Only run this mutator for response objects. This is a hack to identify response objects
							// for database instances, because they include the state field (as opposed to requests,
							// which do not).
							return
						}
						if _, found, _ := unstructured.NestedMap(obj, "settings"); found {
							if _, found, _ := unstructured.NestedStringSlice(obj, "settings", "authorizedGaeApplications"); !found {
								// Include settings.authorizedGaeApplications in response, even if it's empty.
								var val []string
								if err := unstructured.SetNestedStringSlice(obj, val, "settings", "authorizedGaeApplications"); err != nil {
									t.Fatalf("FAIL: setting nested field: %v", err)
								}
							}
						}
						if _, found, _ := unstructured.NestedMap(obj, "settings", "ipConfiguration"); found {
							if _, found, _ := unstructured.NestedStringSlice(obj, "settings", "ipConfiguration", "authorizedNetworks"); !found {
								// Include settings.ipConfiguration.authorizedNetworks in response, even if it's empty.
								var val []string
								if err := unstructured.SetNestedStringSlice(obj, val, "settings", "ipConfiguration", "authorizedNetworks"); err != nil {
									t.Fatalf("FAIL: setting nested field: %v", err)
								}
							}
						}
						if _, found, _ := unstructured.NestedString(obj, "gceZone"); found {
							// Hardcode the zone. GCP chooses this zone within the
							// region, and it varies based on availability.
							if err := unstructured.SetNestedField(obj, "us-central1-a", "gceZone"); err != nil {
								t.Fatalf("FAIL: setting nested field: %v", err)
							}
						}
						if ipConfig, found, _ := unstructured.NestedMap(obj, "settings", "ipConfiguration"); found {
							// Hack fix: remove unpublished field that's suddenly showing up in real gcp proto responses.
							delete(ipConfig, "serverCaMode")
							if err := unstructured.SetNestedMap(obj, ipConfig, "settings", "ipConfiguration"); err != nil {
								t.Fatalf("FAIL: setting nested field: %v", err)
							}
						}
					})
					jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
						if val, found, err := unstructured.NestedString(obj, "kind"); err != nil || !found || val != "sql#usersList" {
							// Only run this mutator for sql users list objects.
							return
						}
						if items, found, _ := unstructured.NestedSlice(obj, "items"); found {
							// Include items[].host in response, even if it's empty.
							newItems := []interface{}{}
							for _, item := range items {
								if itemMap, ok := item.(map[string]interface{}); ok {
									if _, found, _ := unstructured.NestedStringSlice(itemMap, "host"); !found {
										if err := unstructured.SetNestedField(itemMap, "", "host"); err != nil {
											t.Fatalf("FAIL: setting nested field: %v", err)
										}
									}
									newItems = append(newItems, itemMap)
								}
							}
							if err := unstructured.SetNestedSlice(obj, newItems, "items"); err != nil {
								t.Fatalf("FAIL: setting nested field: %v", err)
							}
						}
					})

					// Specific to KMS
					addReplacement("policy.etag", "abcdef0123A=")
					addSetStringReplacement(".cryptoKeyVersions[].createTime", "2024-04-01T12:34:56.123456Z")
					addSetStringReplacement(".cryptoKeyVersions[].generateTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("destroyTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("generateTime", "2024-04-01T12:34:56.123456Z")

					// Specific to BigQueryConnectionConnection.
					addReplacement("aws.accessRole.identity", "048077221682493034546")
					addReplacement("azure.identity", "117243083562690747295")
					addReplacement("cloudResource.serviceAccountId", "bqcx-${projectNumber}-abcd@gcp-sa-bigquery-condel.iam.gserviceaccount.com")
					addReplacement("cloudSql.serviceAccountId", "service-${projectNumber}@gcp-sa-bigqueryconnection.iam.gserviceaccount.com")
					addReplacement("spark.serviceAccountId", "bqcx-${projectNumber}-abcd@gcp-sa-bigquery-condel.iam.gserviceaccount.com")

					// Replace any empty values in LROs; this is surprisingly difficult to fix in mockgcp
					//
					//     "response": {
					// 	-    "@type": "type.googleapis.com/google.protobuf.Empty"
					// 	+    "@type": "type.googleapis.com/google.protobuf.Empty",
					// 	+    "value": {}
					// 	   }
					jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
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

					// Specific to Apigee
					addReplacement("lastModifiedAt", strconv.FormatInt(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix(), 10))
					addReplacement("createdAt", strconv.FormatInt(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC).Unix(), 10))

					// Specific to BigQueryDataTransferConfig
					addReplacement("nextRunTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("ownerInfo.email", "user@google.com")
					addReplacement("userId", "0000000000000000000")
					jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
						if _, found, err := unstructured.NestedString(obj, "destinationDatasetId"); err != nil || !found {
							// This is a hack to only run this mutator for BigQueryDataTransferConfig objects.
							return
						}
						// special handling because the field includes dot
						if _, found, _ := unstructured.NestedString(obj, "params", "connector.authentication.oauth.clientId"); found {
							if err := unstructured.SetNestedField(obj, "client-id", "params", "connector.authentication.oauth.clientId"); err != nil {
								t.Fatalf("FAIL: setting nested field: %v", err)
							}
						}
						if _, found, _ := unstructured.NestedString(obj, "params", "connector.authentication.oauth.clientSecret"); found {
							if err := unstructured.SetNestedField(obj, "client-secret", "params", "connector.authentication.oauth.clientSecret"); err != nil {
								t.Fatalf("FAIL: setting nested field: %v", err)
							}
						}
						delete(obj, "state") // data transfer run state, which depends on timing
					})

					// Specific to IAPSettings
					jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
						if val, found, _ := unstructured.NestedString(obj, "name"); found {
							tokens := strings.Split(val, "/")
							// e.g. "projects/project-id/iap_web/compute-us-central1/services/service-id"
							if len(tokens) >= 6 && tokens[0] == "projects" && tokens[2] == "iap_web" && strings.Contains(tokens[3], "compute") && tokens[4] == "services" {
								tokens[len(tokens)-1] = "${serviceId}"
								if err := unstructured.SetNestedField(obj, strings.Join(tokens, "/"), "name"); err != nil {
									t.Fatalf("FAIL: setting nested field: %v", err)
								}
							}
						}
					})

					// Specific to DocumentAIProcessor
					jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
						// normalize the processorVersionAliases
						aliases, found, _ := unstructured.NestedSlice(obj, "processorVersionAliases")
						if !found {
							return
						}
						for i := range aliases {
							aliasMap, ok := aliases[i].(map[string]any)
							if !ok {
								continue
							}
							processorVersion, found, _ := unstructured.NestedString(aliasMap, "processorVersion")
							if !found {
								continue
							}
							tokens := strings.Split(processorVersion, "/")
							// e.g. projects/project-id/locations/us/processors/processor-id/processorVersions/pretrained-ocr-v1.0-2020-09-23
							if len(tokens) >= 2 && tokens[len(tokens)-2] == "processorVersions" {
								tokens[len(tokens)-1] = "${processorVersionID}"
								if err := unstructured.SetNestedField(aliasMap, strings.Join(tokens, "/"), "processorVersion"); err != nil {
									t.Fatalf("FAIL: setting nested field: %v", err)
								}
							}
						}
						if err := unstructured.SetNestedField(obj, aliases, "processorVersionAliases"); err != nil {
							t.Fatalf("FAIL: setting nested field: %v", err)
						}

						// normalize the defaultProcessorVersion
						if val, found, _ := unstructured.NestedString(obj, "defaultProcessorVersion"); found {
							tokens := strings.Split(val, "/")
							if len(tokens) >= 2 && tokens[len(tokens)-2] == "processorVersions" {
								tokens[len(tokens)-1] = "${processorVersionID}"
								if err := unstructured.SetNestedField(obj, strings.Join(tokens, "/"), "defaultProcessorVersion"); err != nil {
									t.Fatalf("FAIL: setting nested field: %v", err)
								}
							}
						}
					})

					// Specific to VMwareEngineNetwork
					// normalize "vpcNetworks[].network"
					jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
						if val, found, _ := unstructured.NestedString(obj, "name"); found {
							tokens := strings.Split(val, "/")
							if len(tokens) < 2 || tokens[len(tokens)-2] != "vmwareEngineNetworks" {
								return
							}
						}
						vpcNetworks, found, _ := unstructured.NestedSlice(obj, "vpcNetworks")
						if !found {
							return
						}
						for _, vpcNetwork := range vpcNetworks {
							if vpcNetworkMap, ok := vpcNetwork.(map[string]any); ok {
								if val, found, _ := unstructured.NestedString(vpcNetworkMap, "network"); found {
									tokens := strings.Split(val, "/")
									if len(tokens) >= 2 && tokens[len(tokens)-2] == "networks" {
										tokens[len(tokens)-1] = "${networkId}"
										if err := unstructured.SetNestedField(vpcNetworkMap, strings.Join(tokens, "/"), "network"); err != nil {
											t.Fatalf("FAIL: setting nested field: %v", err)
										}
									}
								}
							}
						}
						if err := unstructured.SetNestedSlice(obj, vpcNetworks, "vpcNetworks"); err != nil {
							t.Fatalf("FAIL: setting nested field: %v", err)
						}
					})
					// normalize "response.vpcNetworks[].network"
					jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
						responseObj, found, _ := unstructured.NestedMap(obj, "response")
						if !found {
							return
						}
						name, found, _ := unstructured.NestedString(responseObj, "name")
						if !found || !strings.Contains(name, "vmwareEngineNetworks") {
							return
						}
						vpcNetworks, found, _ := unstructured.NestedSlice(responseObj, "vpcNetworks")
						if !found {
							return
						}
						for _, vpcNetwork := range vpcNetworks {
							if vpcNetworkMap, ok := vpcNetwork.(map[string]any); ok {
								if val, found, _ := unstructured.NestedString(vpcNetworkMap, "network"); found {
									tokens := strings.Split(val, "/")
									if len(tokens) >= 2 && tokens[len(tokens)-2] == "networks" {
										tokens[len(tokens)-1] = "${networkId}"
										if err := unstructured.SetNestedField(vpcNetworkMap, strings.Join(tokens, "/"), "network"); err != nil {
											t.Fatalf("FAIL: setting nested field: %v", err)
										}
									}
								}
							}
						}
						if err := unstructured.SetNestedSlice(responseObj, vpcNetworks, "vpcNetworks"); err != nil {
							t.Fatalf("FAIL: setting nested field: %v", err)
						}
						if err := unstructured.SetNestedMap(obj, responseObj, "response"); err != nil {
							t.Fatalf("FAIL: setting nested field: %v", err)
						}
					})

					// Remove error details which can contain confidential information
					jsonMutators = append(jsonMutators, func(requestURL string, obj map[string]any) {
						response := obj["error"]
						if responseMap, ok := response.(map[string]any); ok {
							delete(responseMap, "details")
						}
					})
					addReplacement("creationTime", "123456789")
					addReplacement("lastModifiedTime", "123456789")

					events.PrettifyJSON(jsonMutators...)

					NormalizeHTTPLog(t, events, h.RegisteredServices(), project, uniqueID, testgcp.TestFolderID.Get(), testgcp.TestOrgID.Get())

					events = RemoveExtraEvents(events)

					// Remove repeated GET requests (after normalization)
					{
						var previous *test.LogEntry
						events = events.KeepIf(func(e *test.LogEntry) bool {
							keep := true
							if e.Request.Method == "GET" && previous != nil {
								if previous.Request.Method == "GET" && previous.Request.URL == e.Request.URL {
									if previous.Response.Status == e.Response.Status {
										if previous.Response.Body == e.Response.Body {
											keep = false
										}
									}
								}
							}
							previous = e
							return keep
						})
					}

					got := events.FormatHTTP()
					normalizers := []func(string) string{}
					normalizers = append(normalizers, IgnoreComments)
					normalizers = append(normalizers, ReplaceString(uniqueID, "${uniqueId}"))
					normalizers = append(normalizers, ReplaceString(project.ProjectID, "${projectId}"))
					normalizers = append(normalizers, ReplaceString(fmt.Sprintf("%d", project.ProjectNumber), "${projectNumber}"))
					for k, v := range r.PathIDs {
						normalizers = append(normalizers, ReplaceString(k, v))
					}
					for k := range r.OperationIDs {
						normalizers = append(normalizers, ReplaceString(k, "${operationID}"))
					}

					if testPause {
						assertNoRequest(t, got, normalizers...)
					} else {
						expectedPath := filepath.Join(fixture.SourceDir, "_http.log")

						h.CompareGoldenFile(expectedPath, got, normalizers...)
					}
				}
			}
		}
	}
}

// assertNoRequest checks that no POSTs or GETs are made against the cloud provider (GCP). This
// is helpful for when we want to test that Pause works correctly and doesn't actuate resources.
func assertNoRequest(t *testing.T, got string, normalizers ...func(s string) string) {
	t.Helper()

	for _, normalizer := range normalizers {
		got = normalizer(got)
	}

	if strings.Contains(got, "POST") {
		t.Fatalf("FAIL: unexpected POST in log: %s", got)
	}

	if strings.Contains(got, "GET") {
		t.Fatalf("FAIL: unexpected GET in log: %s", got)
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
		t.Fatalf("FAIL: error creating CC: %v", err)
	}
}

// verifyUserAgent verifies that the user agent is set to the expected KCC user agent for all requests
func verifyUserAgent(h *create.Harness) {
	for _, event := range h.Events.HTTPEvents {
		userAgent := event.Request.Header.Get("User-Agent")

		// We don't capture the user-agent for GRPC
		if userAgent == "" && event.Request.Method == "GRPC" {
			continue
		}

		tokens := strings.Split(userAgent, " ")
		var keepTokens []string
		for _, token := range tokens {
			// We ignore the google-api-go-client/ prefix; it's added by the GCP client library and difficult to remove.
			if strings.HasPrefix(token, "google-api-go-client/") {
				continue
			}
			// Similarly we ignore the DCL suffix
			if strings.HasPrefix(token, "DeclarativeClientLib/") {
				continue
			}
			keepTokens = append(keepTokens, token)
		}

		got := strings.Join(keepTokens, " ")
		want := gcp.KCCUserAgent()
		if got != want {
			h.Logf("request is %+v", event.Request)
			h.Errorf("FAIL: unexpected user agent for request %v %v.  got %q, expected %q", event.Request.Method, event.Request.URL, got, want)
		}
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

func isOperationDone(s string) bool {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		return false
	}
	return data["status"] == "DONE" || data["done"] == true
}

// addTestTimeout will ensure the test fails if not completed before timeout
func addTestTimeout(ctx context.Context, t *testing.T, timeout time.Duration) context.Context {
	ctx, cancel := context.WithTimeout(ctx, timeout)

	done := false
	timedOut := false
	t.Cleanup(func() {
		done = true
		if timedOut {
			t.Fatalf("FAIL: subtest timeout after %v", timeout)
		}
		cancel()
	})

	go func() {
		<-ctx.Done()
		if !done {
			timedOut = true
		}
	}()

	return ctx
}

func configureVCR(t *testing.T, h *create.Harness) {
	project := h.Project
	replaceWellKnownValues := func(s string) string {
		// Replace project id and number
		result := strings.Replace(s, project.ProjectID, "example-project", -1)
		result = strings.Replace(result, fmt.Sprintf("%d", project.ProjectNumber), "123456789", -1)
		if os.Getenv("TEST_ORG_ID") != "" {
			result = strings.Replace(result, os.Getenv("TEST_ORG_ID"), "123450001", -1)
		}

		addReplacement := func(path string, newValue string) {
			tokens := strings.Split(path, ".")
			obj := make(map[string]any)
			if err := json.Unmarshal([]byte(s), &obj); err == nil {
				toReplace, found, _ := unstructured.NestedString(obj, tokens...)
				if found {
					result = strings.Replace(result, toReplace, newValue, -1)
				}
			}
		}
		// Replace user info
		addReplacement("user", "user@google.com")
		// Replace billing account name
		addReplacement("billingAccountName", "billingAccounts/123456-777777-000001")
		return result
	}

	unique := make(map[string]bool)

	hook := func(i *cassette.Interaction) error {
		// Remove internal error message from failed interactions
		resCode := i.Response.Code
		if resCode == 404 || resCode == 400 || resCode == 403 {
			i.Response.Body = "fake error message"
			// Set Content-Length to zero
			i.Response.ContentLength = 0
		}

		// Discard repeated operation retry interactions
		reqURL := i.Request.URL
		resBody := i.Response.Body

		if strings.Contains(reqURL, "operations") {
			if !isOperationDone(resBody) {
				i.DiscardOnSave = true
			}
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
	h.VCRRecorderNonTF.AddHook(hook, recorder.BeforeSaveHook)
	h.VCRRecorderTF.AddHook(hook, recorder.BeforeSaveHook)
	h.VCRRecorderOauth.AddHook(hook, recorder.BeforeSaveHook)

	matcher := func(r *http.Request, i cassette.Request) bool {
		if r.Method != i.Method || r.URL.String() != i.URL {
			return false
		}

		// Default matcher only checks the request URL and Method. If request body exists, check the body as well.
		// This guarantees that the replayed response matches what the real service would return for that particular request.
		if r.Body != nil && r.Body != http.NoBody {
			var reqBody []byte
			var err error
			reqBody, err = io.ReadAll(r.Body)
			if err != nil {
				t.Fatal("FAIL: [VCR] Failed to read request body")
			}
			r.Body.Close()
			r.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
			if string(reqBody) == i.Body {
				return true
			}

			// If body contains JSON, it might be reordered
			contentType := r.Header.Get("Content-Type")
			if strings.Contains(contentType, "application/json") {
				sortedReqBody, err := sortJSON(string(reqBody))
				if err != nil {
					return false
				}
				sortedBody, err := sortJSON(i.Body)
				if err != nil {
					return false
				}
				return sortedReqBody == sortedBody
			}
		}
		return true
	}
	h.VCRRecorderNonTF.SetMatcher(matcher)
	h.VCRRecorderTF.SetMatcher(matcher)
	h.VCRRecorderOauth.SetMatcher(matcher)
}

func containsCCOrCCC(resources []*unstructured.Unstructured) bool {
	for _, resource := range resources {
		gvk := resource.GroupVersionKind()
		switch gvk.GroupKind() {
		case schema.GroupKind{Group: "core.cnrm.cloud.google.com", Kind: "ConfigConnector"},
			schema.GroupKind{Group: "core.cnrm.cloud.google.com", Kind: "ConfigConnectorContext"}:
			return true
		}
	}
	return false
}

func buildCRDFilter(keepCRDs map[schema.GroupKind]bool) create.HarnessOption {
	return create.FilterCRDs(func(gk schema.GroupKind) bool {
		// Allow core CRDs
		if gk.Group == "core.cnrm.cloud.google.com" {
			return true
		}
		// Otherwise only allow the specified CRDs
		return keepCRDs[gk]
	})
}
