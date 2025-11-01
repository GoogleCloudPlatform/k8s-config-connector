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
				ctx := addTestTimeout(ctx, t, subtestTimeout, sampleKey.Name)
				var harnessOptions []create.HarnessOption

				// Quickly load the sample with a dummy project, just to see if we should skip it
				{
					dummySample := create.LoadSample(t, sampleKey, testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789})
					create.MaybeSkip(t, sampleKey.Name, dummySample.Resources)

					group := dummySample.APIGroup
					skipTestReason := ""
					if group != "" {
						if s := os.Getenv("SKIP_TEST_APIGROUP"); s != "" {
							skippedGroups := strings.Split(s, ",")
							if slice.StringSliceContains(skippedGroups, group) {
								skipTestReason = fmt.Sprintf("skipping test %s because group %q matched entries in SKIP_TEST_APIGROUP=%s", sampleKey.Name, group, s)
							}
						}
						if s := os.Getenv("ONLY_TEST_APIGROUPS"); s != "" {
							onlyGroups := strings.Split(s, ",")
							if !slice.StringSliceContains(onlyGroups, group) {
								skipTestReason = fmt.Sprintf("skipping test %s because group %q did not match ONLY_TEST_APIGROUPS=%s", sampleKey.Name, group, s)
							}
						}
					} else {
						if s := os.Getenv("ONLY_TEST_APIGROUPS"); s != "" {
							t.Skipf("skipping test because cannot determine group for samples, with ONLY_TEST_APIGROUPS=%s", s)
						}
					}

					if skipTestReason != "" {
						t.Skip(skipTestReason)
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

				opt := create.CreateDeleteTestOptions{Create: s.Resources, CleanupResources: true}
				// samples don't do updates so not using SSA is less problematic
				opt.DoNotUseServerSideApplyForCreate = true
				create.RunCreateDeleteTest(h, opt)
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
		// Skip newly added iam/iampartialpolicy for now as they run under TestIAM_AllInSeries
		lightFilter := func(name string, testType resourcefixture.TestType) bool {
			return !strings.Contains(name, "iam-bigqueryconnectionconnectionref") &&
				!strings.Contains(name, "iam-logsinkref") &&
				!strings.Contains(name, "iam-serviceaccountref") &&
				!strings.Contains(name, "iam-serviceidentityref") &&
				!strings.Contains(name, "iam-sqlinstanceref")
		}
		pathFilter := func(path string) bool {
			return !strings.Contains(path, "testdata/iam/iampartialpolicy")
		}

		fixtures := resourcefixture.LoadWithPathFilter(t, pathFilter, lightFilter, nil)
		for _, fixture := range fixtures {
			fixture := fixture
			group := fixture.GVK.Group

			skipTestReason := ""

			if s := os.Getenv("SKIP_TEST_APIGROUP"); s != "" {
				skippedGroups := strings.Split(s, ",")
				if slice.StringSliceContains(skippedGroups, group) {
					skipTestReason = fmt.Sprintf("skipping test %s because group %q matched entries in SKIP_TEST_APIGROUP=%s", fixture.Name, group, s)
				}
			}
			if s := os.Getenv("ONLY_TEST_APIGROUPS"); s != "" {
				groups := strings.Split(s, ",")
				if !slice.StringSliceContains(groups, group) {
					skipTestReason = fmt.Sprintf("skipping test %s because group %q did not match ONLY_TEST_APIGROUPS=%s", fixture.Name, group, s)
				}
			}
			// TODO(b/259496928): Randomize the resource names for parallel execution when/if needed.
			t.Run(fixture.Name, func(t *testing.T) {
				if skipTestReason != "" {
					t.Skip(skipTestReason)
				}

				ctx := addTestTimeout(ctx, t, subtestTimeout, fixture.Name)

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

					// We want to use SSA everywhere, but some of our tests are broken by SSA
					switch group := primaryResource.GetObjectKind().GroupVersionKind().Group; group {
					case "bigtable.cnrm.cloud.google.com",
						"gkehub.cnrm.cloud.google.com",
						"kms.cnrm.cloud.google.com",
						"orgpolicy.cnrm.cloud.google.com",
						"firestore.cnrm.cloud.google.com":
						// Use SSA

					default:
						t.Logf("not yet using SSA for create of resources in group %q", group)
						opt.DoNotUseServerSideApplyForCreate = true
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
				if os.Getenv("GOLDEN_OBJECT_CHECKS") != "" || os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
					folderID := h.FolderID()

					for _, obj := range exportResources {
						// Check the final state of the object in the kube-apiserver (and compare against golden file)
						var normalizer *objectWalker
						{
							u := &unstructured.Unstructured{}
							u.SetGroupVersionKind(obj.GroupVersionKind())
							id := types.NamespacedName{Namespace: obj.GetNamespace(), Name: obj.GetName()}
							if err := h.GetClient().Get(ctx, id, u); err != nil {
								t.Fatalf("FAIL: failed to get KRM object: %v", err)
							}

							normalizer = buildKRMNormalizer(t, u, project, folderID, uniqueID)
							if err := normalizer.VisitUnstructured(u); err != nil {
								t.Fatalf("FAIL: error from normalizer: %v", err)
							}

							got, err := yaml.Marshal(u)
							if err != nil {
								t.Fatalf("FAIL: failed to convert KRM object to yaml: %v", err)
							}
							expectedPath := filepath.Join(fixture.SourceDir, fmt.Sprintf("_generated_object_%v.golden.yaml", testName))
							test.CompareGoldenObject(t, expectedPath, got)
						}

						// Try to export the resource (and compare against golden file)
						exportedYAML := exportResource(h, obj, &Expectations{})
						if exportedYAML != "" {
							exportedObj := &unstructured.Unstructured{}
							if err := yaml.Unmarshal([]byte(exportedYAML), exportedObj); err != nil {
								t.Fatalf("FAIL: error from yaml.Unmarshal: %v", err)
							}

							// Note: the normalizer for the object has more information, so we reuse that normalizer
							if err := normalizer.VisitUnstructured(exportedObj); err != nil {
								t.Fatalf("FAIL: error from normalizer: %v", err)
							}

							got, err := yaml.Marshal(exportedObj)
							if err != nil {
								t.Fatalf("FAIL: failed to convert KRM object to yaml: %v", err)
							}

							expectedPath := filepath.Join(fixture.SourceDir, fmt.Sprintf("_generated_export_%v.golden", testName))
							h.CompareGoldenFile(expectedPath, string(got), IgnoreComments)
						}

					}
				}

				if ShouldTestRereconiliation(t, testName, primaryResource) {
					h.Log("Testing re-reconciliation...", "test name", testName, "primary GVK", primaryResource.GroupVersionKind().String())
					eventsBefore := h.Events.HTTPEvents

					oldGeneration := getGeneration(h, primaryResource)
					touchObject(h, primaryResource)
					// Pause to allow re-reconciliation
					// (annotations don't change the generation, so we can't wait for observedGeneration)
					time.Sleep(2 * time.Second)

					eventsAfter := h.Events.HTTPEvents

					h.Events.HTTPEvents = eventsBefore

					for i := len(eventsBefore); i < len(eventsAfter); i++ {
						event := eventsAfter[i]
						isReadOnly := false
						switch event.Request.Method {
						case "GET":
							isReadOnly = true
						}
						if !isReadOnly {
							t.Errorf("FAIL: unexpected event during re-reconciliation: %v", event)
						}
					}

					newGeneration := getGeneration(h, primaryResource)
					if oldGeneration != newGeneration {
						t.Errorf("FAIL: re-reconciliation caused generation to change (from %v to %v)", oldGeneration, newGeneration)
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

					got, normalizers := LegacyNormalize(t, h, project, uniqueID, events)
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
	// We only expect full watches on Namespaces, CRDs, CCs and CCCs (currently)
	// and K8s Secret.
	allowedFullWatches := sets.NewString(
		"/apis/core.cnrm.cloud.google.com/v1beta1/configconnectorcontexts",
		"/apis/core.cnrm.cloud.google.com/v1beta1/configconnectors",
		"/apis/apiextensions.k8s.io/v1/customresourcedefinitions",
		"/api/v1/secrets",
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
func addTestTimeout(ctx context.Context, t *testing.T, timeout time.Duration, name string) context.Context {

	if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "real" {
		// If the target is real, check if SUBTEST_TIMEOUT_E2E is present set
		// accordingly, or fallback to original timeouts if there's any error.

		if subtestTimeoutE2EStr := os.Getenv("SUBTEST_TIMEOUT_E2E"); subtestTimeoutE2EStr != "" {
			parsedTimeout, err := time.ParseDuration(subtestTimeoutE2EStr)
			if err == nil {
				timeout = parsedTimeout
			}
		}
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	done := false
	timedOut := false
	t.Cleanup(func() {
		done = true
		if timedOut {
			t.Fatalf("FAIL: subtest %s timeout after %v", name, timeout)
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

func TestIAM_AllInSeries(t *testing.T) {
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
		subtestTimeout = 1 * time.Minute
	}

	t.Run("iam-fixtures", func(t *testing.T) {
		fixtures := resourcefixture.LoadWithPathFilter(t, func(path string) bool {
			// Only run fixtures under iam/iampartialpolicy
			return strings.Contains(path, "testdata/iam/iampartialpolicy") &&
				// todo kcc team: need to implement GetIAM/ SetIAM for mock
				!strings.Contains(path, "computeimage") &&
				// todo kcc team: "Failed to get server metadata from context" for some reason
				!strings.Contains(path, "storagebucket") &&
				// todo acpana exclude failing tests for now; needs setup validation
				!strings.Contains(path, "cloudfunctionsfunction") &&
				!strings.Contains(path, "computedisk") &&
				!strings.Contains(path, "computeinstance") &&
				!strings.Contains(path, "computesubnetwork") &&
				!strings.Contains(path, "dataproccluster") &&
				!strings.Contains(path, "folder") &&
				!strings.Contains(path, "iamserviceaccount") &&
				!strings.Contains(path, "servicedirectoryservice") &&
				!strings.Contains(path, "spannerdatabase")
		}, nil, nil)
		for _, fixture := range fixtures {
			fixture := fixture
			group := fixture.GVK.Group

			skipTestReason := ""

			if s := os.Getenv("SKIP_TEST_APIGROUP"); s != "" {
				skippedGroups := strings.Split(s, ",")
				if slice.StringSliceContains(skippedGroups, group) {
					skipTestReason = fmt.Sprintf("skipping test %s because group %q matched entries in SKIP_TEST_APIGROUP=%s", fixture.Name, group, s)
				}
			}
			if s := os.Getenv("ONLY_TEST_APIGROUPS"); s != "" {
				groups := strings.Split(s, ",")
				if !slice.StringSliceContains(groups, group) {
					skipTestReason = fmt.Sprintf("skipping test %s because group %q did not match ONLY_TEST_APIGROUPS=%s", fixture.Name, group, s)
				}
			}
			t.Run(fixture.Name, func(t *testing.T) {
				if skipTestReason != "" {
					t.Skip(skipTestReason)
				}

				ctx := addTestTimeout(ctx, t, subtestTimeout, fixture.Name)

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

				runScenario(ctx, t, false, fixture, loadFixture)
			})
		}
	})

	t.Logf("shutting down manager")
	cancel()
}
