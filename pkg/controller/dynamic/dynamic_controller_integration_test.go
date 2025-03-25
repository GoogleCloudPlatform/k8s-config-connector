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

//go:build integration
// +build integration

package dynamic_test

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceactuation"
	dclextension "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testreconciler "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller/reconciler"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testk8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/k8s"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/contexts"
	testrunner "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/runner"
	testservicemapping "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemapping"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/cenkalti/backoff"
	"github.com/ghodss/yaml"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

type httpRoundTripperKeyType int

// httpRoundTripperKey is the key value for http.RoundTripper in a context.Context
var httpRoundTripperKey httpRoundTripperKeyType

func init() {
	// run-tests and skip-tests allows you to limit the tests that are run by
	// specifying regexes to be used to match test names. See the
	// formatTestName() function to see what test names look like.
	flag.StringVar(&runTestsRegex, "run-tests", "", "run only the tests whose names match the given regex")
	flag.StringVar(&skipTestsRegex, "skip-tests", "", "skip the tests whose names match the given regex, even those that match the run-tests regex")

	// cleanup-resources allows you to disable the cleanup of resources created during testing. This can be useful for debugging test failures.
	// The default value is true.
	//
	// To use this flag, you MUST use an equals sign as follows: go test -tags=integration -cleanup-resources=false
	flag.BoolVar(&cleanupResources, "cleanup-resources", true, "when enabled, "+
		"cloud resources created by tests will be cleaned up at the end of a test")

	// Allow for capture of http requests during a test.
	transport_tpg.DefaultHTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		ret := inner
		if t := ctx.Value(httpRoundTripperKey); t != nil {
			ret = &http.Client{Transport: t.(http.RoundTripper)}
		}
		if artifacts := os.Getenv("ARTIFACTS"); artifacts == "" {
			log := log.FromContext(ctx)
			log.Info("env var ARTIFACTS is not set; will not record http log")
		} else {
			outputDir := filepath.Join(artifacts, "http-logs")
			t := test.NewHTTPRecorder(ret.Transport, test.NewDirectoryEventSink(outputDir))
			ret = &http.Client{Transport: t}
		}
		return ret
	}
}

var (
	mgr              manager.Manager
	runTestsRegex    string
	skipTestsRegex   string
	cleanupResources bool
)

const resourceIDTestVar = "${resourceId}"

func shouldRunBasedOnRunAndSkipRegexes(parentTestName string, fixture resourcefixture.ResourceFixture) bool {
	testName := formatTestName(parentTestName, fixture)

	// If a skip-tests regex has been provided and it matches the test name, skip the test.
	if skipTestsRegex != "" {
		if regexp.MustCompile(skipTestsRegex).MatchString(testName) {
			return false
		}
	}

	// If a run-tests regex has been provided and it doesn't match the test name, skip the test.
	if runTestsRegex != "" {
		if !regexp.MustCompile(runTestsRegex).MatchString(testName) {
			return false
		}
	}

	return true
}

func TestAcquire(t *testing.T) {
	ctx := context.TODO()

	t.Parallel()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		if !shouldRunBasedOnRunAndSkipRegexes("TestAcquire", fixture) {
			return false
		}

		// Never run the acquire test for 'containerannotations' test cases.
		// This is because these test cases intentionally omit required
		// hierarchical references to test that the webhooks default them
		// correctly. However, the acquire test needs to be able to create GCP
		// resources using create.yaml's without applying them onto the K8s API
		// server, which means trying to create GCP resources using
		// create.yaml's that are missing the required references and without
		// going through the defaulting provided by webhooks.
		if fixture.Type == resourcefixture.ContainerAnnotations {
			return false
		}

		// Run the acquire test for 'resourceid' test cases to test that
		// resources can be acquired using `spec.resourceID`.
		if fixture.Type == resourcefixture.ResourceID {
			return true
		}

		// Run the acquire test for a representative set of resource kinds.
		// Note: ensuring that all fields are accounted for and not changed
		// when applying the same YAMLs is handled separately by the NoChange
		// test.
		// TODO(b/239876828): Add "DataflowJob" back to acquisition tests.
		kinds := map[string]bool{
			// basic resource with no dependencies
			"PubSubTopic": true,
			// resource with dependencies
			"PubSubSubscription": true,
			// resource with no labels support
			"BigQueryTable": true,
			// resource acquirable by displayName and parent org/folder ID if
			// server-generated ID (i.e. folder ID) is not specified
			"Folder": true,
			// used as an integration test verifying that falsey values are not
			// incorrectly defaulted. (b/178744782)
			"ComputeNetwork": true,
		}
		return kinds[fixture.GVK.Kind]
	}
	testFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext) {
		context, err := contexts.GetResourceContext(testContext.ResourceFixture, systemContext.DCLConverter.MetadataLoader, systemContext.DCLConverter.SchemaLoader)
		if err != nil {
			t.Fatalf("error getting resource context for gvk %v: %v", testContext.ResourceFixture.GVK, err)
		}
		testReconcileAcquire(ctx, t, testContext, systemContext, context)
	}
	testrunner.RunAllWithDependenciesCreatedButNotObject(ctx, t, mgr, shouldRun, testFunc)
}

func TestCreateNoChangeUpdateDelete(t *testing.T) {
	ctx := context.TODO()

	t.Parallel()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		switch fixture.Type {
		case resourcefixture.IAMExternalOnlyRef, resourcefixture.IAMMemberReferences:
			return false
		}

		// Skip ResourceID test cases for resources with server-generated IDs.
		// These test cases exist solely to test acquisition of resources with
		// server-generated IDs.
		if fixture.Type == resourcefixture.ResourceID {
			switch fixture.Name {
			case "servergeneratedresourceid":
				return false
			case "servergeneratedresourceidfordcl":
				return false
			}
		}

		return shouldRunBasedOnRunAndSkipRegexes("TestCreateNoChangeUpdateDelete", fixture)
	}
	testFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext) {
		context, err := contexts.GetResourceContext(testContext.ResourceFixture, systemContext.DCLConverter.MetadataLoader, systemContext.DCLConverter.SchemaLoader)
		if err != nil {
			t.Fatalf("error getting resource context for gvk %v: %v", testContext.ResourceFixture.GVK, err)
		}
		testReconcileCreateNoChangeUpdateDelete(ctx, t, testContext, systemContext, context)
	}
	testrunner.RunAllWithDependenciesCreatedButNotObject(ctx, t, mgr, shouldRun, testFunc)
}

func formatTestName(parentTestName string, fixture resourcefixture.ResourceFixture) string {
	return fmt.Sprintf("%v/%v", parentTestName, resourcefixture.FormatTestName(fixture))
}
func testCreate(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext, resourceContext contexts.ResourceContext) {
	kubeClient := systemContext.Manager.GetClient()
	initialUnstruct := testContext.CreateUnstruct.DeepCopy()
	if err := kubeClient.Create(ctx, initialUnstruct); err != nil {
		t.Fatalf("error creating %v resource %v: %v", initialUnstruct.GetKind(), initialUnstruct.GetName(), err)
	}
	t.Logf("resource created with %v\r", initialUnstruct)
	systemContext.Reconciler.Reconcile(ctx, initialUnstruct, testreconciler.ExpectedSuccessfulReconcileResultFor(systemContext.Reconciler, initialUnstruct), nil)
	validateCreate(ctx, t, testContext, systemContext, resourceContext, initialUnstruct.GetGeneration())
}

func validateCreate(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext,
	resourceContext contexts.ResourceContext, preReconcileGeneration int64) {
	kubeClient := systemContext.Manager.GetClient()
	initialUnstruct := testContext.CreateUnstruct.DeepCopy()
	// Check labels match on create
	reconciledUnstruct := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       initialUnstruct.GetKind(),
			"apiVersion": initialUnstruct.GetAPIVersion(),
		},
	}
	if err := kubeClient.Get(ctx, testContext.NamespacedName, reconciledUnstruct); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}

	// Hack: Optionally wait before getting the object in GCP. This is to work around some issues with troublesome
	// services in GCP that claim to be done with creating / updating the resource before it is actually available.
	time.Sleep(resourceContext.PostModifyDelay)

	gcpUnstruct, err := resourceContext.Get(ctx, t, reconciledUnstruct, systemContext.TFProvider, kubeClient, systemContext.SMLoader, systemContext.DCLConfig, systemContext.DCLConverter, systemContext.HttpClient)
	if err != nil {
		t.Fatalf("[validateCreate] unexpected error when GET-ing '%v': %v", initialUnstruct.GetName(), err)
	}
	t.Logf("created resource is %v\r", gcpUnstruct)
	if resourceContext.SupportsLabels(systemContext.SMLoader, initialUnstruct) {
		testcontroller.AssertLabelsMatchAndHaveManagedLabel(t, gcpUnstruct.GetLabels(), reconciledUnstruct.GetLabels())
	}

	// Check that an "Updating" event was recorded, indicating that the
	// controller tried to update the resource at all.
	// TODO(acpana): figure out if we want to expose Updating event for direct resources
	rt, err := testreconciler.ReconcilerTypeForObject(initialUnstruct)
	if err != nil {
		t.Fatalf("error getting reconciler type: %v", err)
	}
	if rt != testreconciler.ReconcilerTypeDirect {
		testcontroller.AssertEventRecordedforUnstruct(t, kubeClient, reconciledUnstruct, k8s.Updating)
	}

	// Check that condition is ready and "UpToDate" event was recorded
	// TODO: (eventually) check default fields are propagated correctly
	testcontroller.AssertReadyCondition(t, reconciledUnstruct, preReconcileGeneration)
	testcontroller.AssertEventRecordedforUnstruct(t, kubeClient, reconciledUnstruct, k8s.UpToDate)

	verifyResourceIDIfSupported(t, systemContext, resourceContext, reconciledUnstruct, initialUnstruct)

	generationIncrease := reconciledUnstruct.GetGeneration() - preReconcileGeneration
	if generationIncrease > 1 {
		// Generation should have incremented at most once, only due to defaulted field sync-back.
		t.Fatalf("unexpected generation increase %v", generationIncrease)
	}

	// Check observedGeneration matches with the pre-reconcile generation
	testcontroller.AssertObservedGenerationEquals(t, reconciledUnstruct, preReconcileGeneration)

	// Check 'state-into-spec: absent' is set.
	if testContext.ResourceFixture.Type == resourcefixture.StateAbsentInSpec {
		annotationValue, ok := k8s.GetAnnotation(k8s.StateIntoSpecAnnotation, reconciledUnstruct)
		if !ok || annotationValue != k8s.StateAbsentInSpec {
			t.Errorf("annotation %v should be %v but got %v", k8s.StateIntoSpecAnnotation, k8s.StateAbsentInSpec, annotationValue)
		}
	}
}

func testNoChangeAfterCreate(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext, resourceContext contexts.ResourceContext) {
	testNoChange(ctx, t, testContext, systemContext, resourceContext, false)
}

// testNoChangeAfterUpdate is enabled only on resources allowlisted inside the function.
func testNoChangeAfterUpdate(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext, resourceContext contexts.ResourceContext) {
	// Do not run for tests with `SkipUpdate` explicitly set to 'true'.
	if resourceContext.SkipUpdate {
		return
	}
	// Do not run for tests without an update.yaml set.
	if testContext.UpdateUnstruct == nil {
		t.Logf("UpdateUnstruct not set; skipping testNoChangeAfterUpdate")
		return
	}
	switch testContext.ResourceFixture.GVK.GroupKind() {
	case schema.GroupKind{Group: "sql.cnrm.cloud.google.com", Kind: "SQLInstance"}: // test coverage for https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/1802
	default:
		return
	}
	testNoChange(ctx, t, testContext, systemContext, resourceContext, true)
}

// testNoChange verifies that reconciling a resource which has not changed does not result in
// any meaningful changes.
func testNoChange(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext, resourceContext contexts.ResourceContext, useUpdateFixture bool) {
	if resourceContext.SkipNoChange {
		return
	}
	kubeClient := systemContext.Manager.GetClient()
	initialUnstruct := testContext.CreateUnstruct.DeepCopy()
	if useUpdateFixture {
		initialUnstruct = testContext.UpdateUnstruct.DeepCopy()
	}
	if err := kubeClient.Get(ctx, testContext.NamespacedName, initialUnstruct); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	preReconcileGeneration := initialUnstruct.GetGeneration()

	// Delete all events for the resource so that we can check later at the end
	// of this test that the right events are recorded.
	testcontroller.DeleteAllEventsForUnstruct(t, kubeClient, initialUnstruct)

	// Reconcile resource without changing anything in its configuration
	reconciledUnstruct := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       initialUnstruct.GetKind(),
			"apiVersion": initialUnstruct.GetAPIVersion(),
		},
	}
	systemContext.Reconciler.Reconcile(ctx, initialUnstruct, testreconciler.ExpectedSuccessfulReconcileResultFor(systemContext.Reconciler, initialUnstruct), nil)
	if err := kubeClient.Get(ctx, testContext.NamespacedName, reconciledUnstruct); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}

	if reconciledUnstruct.GetGeneration() != initialUnstruct.GetGeneration() {
		t.Errorf("generation incremented during expected no-op")
	}

	if testContext.ResourceFixture.GVK.Kind == "DataflowJob" {
		// Check that the Dataflow job has not been updated. This means
		// checking that its status.jobId is still the same since Dataflow jobs
		// are updated by creating a new job (which would have a new job ID) to
		// replace the existing one.
		checkDataflowJobNoChange(t, initialUnstruct, reconciledUnstruct)
	}

	// Check that an "Updating" event was never recorded.
	testcontroller.AssertEventNotRecordedforUnstruct(t, kubeClient, initialUnstruct, k8s.Updating)

	// Check observedGeneration matches with the pre-reconcile generation
	testcontroller.AssertObservedGenerationEquals(t, reconciledUnstruct, preReconcileGeneration)
}

func testUpdate(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext, resourceContext contexts.ResourceContext) {
	if testContext.UpdateUnstruct == nil {
		t.Logf("UpdateUnstruct not set; skipping update")
		return
	}
	// Tests with `SkipUpdate` explicitly set to 'true' or tests for
	// auto-generated resources don't support update test.
	if resourceContext.SkipUpdate || resourceContext.IsAutoGenerated(systemContext.SMLoader, testContext.UpdateUnstruct) &&
		// TODO: Remove the condition for BigQueryConnectionConnection after it becomes v1beta1.
		// BigQueryConnectionConnection is a v1alpha1 resource supported via the
		// autogen channel. Usually an update test is skipped for an autogen
		// resource, but we do have thorough testing for
		// BigQueryConnectionConnection so we make it an exception here.
		// This exception shouldn't be applied to other autogen resources.
		resourceContext.ResourceKind != "BigQueryConnectionConnection" {
		return
	}

	kubeClient := systemContext.Manager.GetClient()
	initialUnstruct := testContext.CreateUnstruct.DeepCopy()
	if err := kubeClient.Get(ctx, testContext.NamespacedName, initialUnstruct); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}

	// Delete all events for the resource so that we can check later at the end
	// of this test that the right events are recorded.
	testcontroller.DeleteAllEventsForUnstruct(t, kubeClient, initialUnstruct)

	// Update resource from test data
	updateUnstruct := testContext.UpdateUnstruct.DeepCopy()
	updateUnstruct.SetResourceVersion(initialUnstruct.GetResourceVersion())
	// For resources with server-generated IDs, ensure the relevant fields are in the status
	status := initialUnstruct.Object["status"]
	if err := unstructured.SetNestedField(updateUnstruct.Object, status, "status"); err != nil {
		t.Fatalf("error setting status on updateUnstruct: %v", err)
	}
	patch := client.MergeFrom(testContext.CreateUnstruct)
	t.Logf("patching %v with %v\r", updateUnstruct, patch)
	if err := kubeClient.Patch(ctx, updateUnstruct, patch); err != nil {
		t.Fatalf("unexpected error when updating '%v': %v", initialUnstruct.GetName(), err)
	}
	preReconcileGeneration := updateUnstruct.GetGeneration()
	systemContext.Reconciler.Reconcile(ctx, updateUnstruct, testreconciler.ExpectedSuccessfulReconcileResultFor(systemContext.Reconciler, updateUnstruct), nil)

	reconciledUnstruct := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       updateUnstruct.GetKind(),
			"apiVersion": updateUnstruct.GetAPIVersion(),
		},
	}
	if err := kubeClient.Get(ctx, testContext.NamespacedName, reconciledUnstruct); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}

	generationIncrease := reconciledUnstruct.GetGeneration() - updateUnstruct.GetGeneration()
	if generationIncrease > 1 {
		// Generation should have incremented at most once, only due to defaulted field sync-back.
		t.Fatalf("unexpected generation increase %v", generationIncrease)
	}

	// Hack: Optionally wait before getting the object in GCP. This is to work around some issues with troublesome
	// services in GCP that claim to be done with creating / updating the resource before it is actually available.
	time.Sleep(resourceContext.PostModifyDelay)

	// Check labels match on update
	gcpUnstruct, err := resourceContext.Get(ctx, t, reconciledUnstruct, systemContext.TFProvider, kubeClient, systemContext.SMLoader, systemContext.DCLConfig, systemContext.DCLConverter, nil)
	if err != nil {
		t.Fatalf("[testUpdate] unexpected error when GET-ing '%v': %v", updateUnstruct.GetName(), err)
	}
	if resourceContext.SupportsLabels(systemContext.SMLoader, updateUnstruct) {
		testcontroller.AssertLabelsMatchAndHaveManagedLabel(t, gcpUnstruct.GetLabels(), testContext.UpdateUnstruct.GetLabels())
	}

	// If the object has a spec, check that the live GCP resource reflects the
	// updates made by the update struct
	if initialUnstruct.Object["spec"] != nil {
		if gcpUnstruct.Object["spec"] == nil {
			t.Fatalf("GCP resource has a nil spec even though it was created using a resource with a non-nil spec")
		}
		changedSpecFields := getChangedFields(initialUnstruct.Object, reconciledUnstruct.Object, "spec")
		removeSensitiveFields(reconciledUnstruct, changedSpecFields) // remove sensitive fields which are reacted by the GCP API
		assertObjectContains(t, gcpUnstruct.Object["spec"].(map[string]interface{}), changedSpecFields)
	}

	// Check that an "Updating" event was recorded, indicating that the
	// controller tried to update the resource at all.
	// TODO(acpana): figure out if we want to expose Updating event for direct resources
	rt, err := testreconciler.ReconcilerTypeForObject(updateUnstruct)
	if err != nil {
		t.Fatalf("error getting reconciler type: %v", err)
	}
	if rt != testreconciler.ReconcilerTypeDirect {
		testcontroller.AssertEventRecordedforUnstruct(t, kubeClient, reconciledUnstruct, k8s.Updating)
	}
	// Check if condition is ready and update event was recorded
	testcontroller.AssertReadyCondition(t, reconciledUnstruct, preReconcileGeneration)
	testcontroller.AssertEventRecordedforUnstruct(t, kubeClient, reconciledUnstruct, k8s.UpToDate)

	// Check observedGeneration matches with the pre-reconcile generation
	testcontroller.AssertObservedGenerationEquals(t, reconciledUnstruct, preReconcileGeneration)

	verifyResourceIDIfSupported(t, systemContext, resourceContext, reconciledUnstruct, updateUnstruct)
}

// this test deletes the resource directly on GCP and then reconciles and verifies the resource was recreated correctly
func testDriftCorrection(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext, resourceContext contexts.ResourceContext) {
	if shouldSkipDriftDetection(t, resourceContext, systemContext.SMLoader, systemContext.DCLConverter.MetadataLoader, testContext.CreateUnstruct) {
		return
	}
	kubeClient := systemContext.Manager.GetClient()
	testUnstruct := testContext.CreateUnstruct.DeepCopy()
	if err := kubeClient.Get(ctx, testContext.NamespacedName, testUnstruct); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	// For test cases with `cnrm.cloud.google.com/reconcile-interval-in-seconds` annotation set to 0, we should skip drift correction test.
	if skip, _ := resourceactuation.ShouldSkip(testUnstruct); skip {
		return
	}
	// Delete all events for the resource so that we can check later at the end
	// of this test that the right events are recorded.
	testcontroller.DeleteAllEventsForUnstruct(t, kubeClient, testUnstruct)

	t.Logf("testDriftCorrection: deleting kube object %v", testUnstruct)
	if err := resourceContext.Delete(ctx, t, testUnstruct, systemContext.TFProvider, systemContext.Manager.GetClient(), systemContext.SMLoader, systemContext.DCLConfig, systemContext.DCLConverter, systemContext.HttpClient); err != nil {
		t.Fatalf("error deleting: %v", err)
	}

	// Underlying APIs may not have strongly-consistent reads due to caching. Sleep before attempting a re-reconcile, to
	// give the underlying system some time to propagate the deletion info.
	time.Sleep(time.Second * 10)

	// get the current state
	t.Logf("reconcile with %v\r", testUnstruct)
	systemContext.Reconciler.Reconcile(ctx, testUnstruct, testreconciler.ExpectedSuccessfulReconcileResultFor(systemContext.Reconciler, testUnstruct), nil)
	t.Logf("reconciled with %v\r", testUnstruct)
	validateCreate(ctx, t, testContext, systemContext, resourceContext, testUnstruct.GetGeneration())
}

func shouldSkipDriftDetection(t *testing.T, resourceContext contexts.ResourceContext, smLoader *servicemappingloader.ServiceMappingLoader,
	serviceMetadataLoader dclmetadata.ServiceMetadataLoader, u *unstructured.Unstructured) bool {
	if !testgcp.ResourceSupportsDeletion(u.GetKind()) {
		// The drift correction test relies on being able to delete the underlying resource.
		return true
	}
	if resourceContext.SkipDriftDetection {
		return true
	}

	rt, err := testreconciler.ReconcilerTypeForObject(u)
	if err != nil {
		t.Fatalf("error getting reconciler type: %v", err)
	}

	// Skip drift detection test for dcl-based resources with server-generated id.
	if rt == testreconciler.ReconcilerTypeDCL {
		s, found := dclextension.GetNameFieldSchema(resourceContext.DCLSchema)
		if !found {
			// The resource doesn't have a 'resourceID' field.
			return false
		}
		isServerGenerated, err := dclextension.IsResourceIDFieldServerGenerated(s)
		if err != nil {
			t.Fatalf("error parsing `resourceID` field schema: %v", err)
		}
		return isServerGenerated
	} else if rt == testreconciler.ReconcilerTypeTerraform {
		// Skip drift detection test for tf-based resources with server-generated id.
		rc := testservicemapping.GetResourceConfig(t, smLoader, u)
		return hasServerGeneratedId(*rc)
	} else {
		// Drift detection tests are enabled by default for direct resources.
		return false
	}
}

func hasServerGeneratedId(rc v1alpha1.ResourceConfig) bool {
	return rc.ServerGeneratedIDField != ""
}

func testDelete(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext, resourceContext contexts.ResourceContext) {
	if resourceContext.SkipDelete {
		return
	}
	kubeClient := systemContext.Manager.GetClient()
	testReconciler := systemContext.Reconciler
	initialUnstruct := testContext.CreateUnstruct.DeepCopy()
	if err := kubeClient.Delete(ctx, initialUnstruct); err != nil {
		t.Fatalf("error deleting resource: %v", err)
	}

	// Test that the deletion defender finalizer causes the resource to requeue
	// and still exist on the underlying API
	reconciledUnstruct := testContext.CreateUnstruct.DeepCopy()
	testReconciler.Reconcile(ctx, reconciledUnstruct, testreconciler.ExpectedRequeueReconcileStruct, nil)
	if err := kubeClient.Get(ctx, testContext.NamespacedName, reconciledUnstruct); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}
	if _, err := resourceContext.Get(ctx, t, reconciledUnstruct, systemContext.TFProvider, kubeClient, systemContext.SMLoader, systemContext.DCLConfig, systemContext.DCLConverter, nil); err != nil {
		t.Errorf("expected resource %s to not be deleted with deletion defender finalizer, but got error: %s",
			initialUnstruct.GetName(), err)
	}

	// Perform the deletion on the underlying API
	testk8s.RemoveDeletionDefenderFinalizerForUnstructured(t, reconciledUnstruct, kubeClient)
	testReconciler.Reconcile(ctx, reconciledUnstruct, testreconciler.ExpectedSuccessfulReconcileResultFor(systemContext.Reconciler, reconciledUnstruct), nil)

	if !testgcp.ResourceSupportsDeletion(testContext.ResourceFixture.GVK.Kind) {
		_, err := resourceContext.Get(ctx, t, reconciledUnstruct, systemContext.TFProvider, kubeClient, systemContext.SMLoader, systemContext.DCLConfig, systemContext.DCLConverter, nil)
		if err != nil {
			t.Errorf("expected resource %s to exist after deletion, but got error: %s", initialUnstruct.GetName(), err)
		}
	} else {
		getFunc := func() error {
			// for some resources, Get after Delete is eventually consistent, for that reason we retry until an error is returned
			_, err := resourceContext.Get(ctx, t, reconciledUnstruct, systemContext.TFProvider, kubeClient, systemContext.SMLoader, systemContext.DCLConfig, systemContext.DCLConverter, nil)
			if err == nil {
				return fmt.Errorf("expected error, instead got 'nil'")
			}
			return backoff.Permanent(err)
		}
		expBackoff := backoff.NewExponentialBackOff()
		expBackoff.MaxElapsedTime = 60 * time.Second
		expBackoff.MaxInterval = 10 * time.Second
		err := backoff.Retry(getFunc, expBackoff)
		// TODO: remove gcp.IsNotFoundError(...) once all resources are converted to use terraform for ResourceContext Create / Get
		if !gcp.IsNotFoundError(err) && !contexts.IsNotFoundError(err) {
			t.Errorf("expected GCP client to return NotFound for '%v', instead got: %v", initialUnstruct.GetName(), err)
		}

		err = kubeClient.Get(ctx, testContext.NamespacedName, initialUnstruct)
		if err == nil || !errors.IsNotFound(err) {
			t.Errorf("unexpected error value: '%v'", err)
		}
	}

	// Check that "Deleted" event was recorded
	testcontroller.AssertEventRecordedforUnstruct(t, kubeClient, initialUnstruct, k8s.Deleted)
}

func testReconcileCreateNoChangeUpdateDelete(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext, resourceContext contexts.ResourceContext) {
	resourceCleanup := systemContext.Reconciler.BuildCleanupFunc(ctx, testContext.CreateUnstruct, getResourceCleanupPolicy())
	defer resourceCleanup()
	testCreate(ctx, t, testContext, systemContext, resourceContext)
	testNoChangeAfterCreate(ctx, t, testContext, systemContext, resourceContext)
	testUpdate(ctx, t, testContext, systemContext, resourceContext)
	testNoChangeAfterUpdate(ctx, t, testContext, systemContext, resourceContext)
	testDriftCorrection(ctx, t, testContext, systemContext, resourceContext)
	testDelete(ctx, t, testContext, systemContext, resourceContext)
}

func checkComputeNetworkUpdate(t *testing.T, updateUnstruct *unstructured.Unstructured, gcpUnstruct *unstructured.Unstructured) {
	expect, _, err := unstructured.NestedString(updateUnstruct.Object, "spec", "routingMode")
	if err != nil {
		t.Error(err)
	}
	actual, _, err := unstructured.NestedString(gcpUnstruct.Object, "routingConfig", "routingMode")
	if err != nil {
		t.Error(err)
	}
	if expect != actual {
		t.Errorf("unexpected value for routingMode: got %v, want %v", actual, expect)
	}
}

func checkDataflowJobNoChange(t *testing.T, initialUnstruct, reconciledUnstruct *unstructured.Unstructured) {
	expectJobID, found, err := unstructured.NestedString(initialUnstruct.Object, "status", "jobId")
	if err != nil {
		t.Error(err)
	}
	if !found {
		t.Errorf("initial unstruct does not have a status.jobId field")
	}
	actualJobID, found, err := unstructured.NestedString(reconciledUnstruct.Object, "status", "jobId")
	if err != nil {
		t.Error(err)
	}
	if !found {
		t.Errorf("reconciled unstruct does not have a status.jobId field")
	}
	if expectJobID != actualJobID {
		t.Errorf("expected status.jobId to be unchanged: got %v, want %v", actualJobID, expectJobID)
	}
}

func testReconcileAcquire(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext, resourceContext contexts.ResourceContext) {
	kubeClient := systemContext.Manager.GetClient()
	initialUnstruct := testContext.CreateUnstruct.DeepCopy()

	// Create the resource on GCP if it doesn't exist.
	//
	// If the unstruct contains a ${resourceId} test variable, that means its
	// spec.resourceID is only meant to be used for acquisition. Therefore,
	// strip out spec.resourceID for creation.
	unstructToCreate := initialUnstruct.DeepCopy()
	if containsResourceIDTestVar(t, unstructToCreate) {
		testcontroller.RemoveResourceID(unstructToCreate)
	}
	var gcpUnstruct *unstructured.Unstructured
	var err error
	gcpUnstruct, err = resourceContext.Get(ctx, t, unstructToCreate, systemContext.TFProvider, kubeClient, systemContext.SMLoader, systemContext.DCLConfig, systemContext.DCLConverter, nil)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			t.Fatalf("[testReconcileAcquire] unexpected error when GET-ing '%v': %v", unstructToCreate.GetName(), err)
		}
		if gcpUnstruct, err = resourceContext.Create(ctx, t, unstructToCreate, systemContext.TFProvider, kubeClient, systemContext.SMLoader, systemContext.DCLConfig, systemContext.DCLConverter); err != nil {
			t.Fatalf("unexpected error when creating GCP resource '%v': %v", unstructToCreate.GetName(), err)
		}
		if unstructToCreate.GroupVersionKind().Kind == "Folder" {
			// We should not be using the search method, it is only eventually consistent.
			t.Logf("created GCP Folder; waiting 60 seconds for eventual consistency to catch up")
			time.Sleep(time.Minute)
		}
	}

	// Acquire the resource using the original unstruct.
	//
	// If the unstruct contains a ${resourceId} test variable, that means its
	// spec.resourceID needs to be set with the live resource's resource ID to
	// to make the unstruct usable for acquiring the live resource.
	if containsResourceIDTestVar(t, initialUnstruct) {
		resourceID, ok := testcontroller.GetResourceID(t, gcpUnstruct)
		if !ok || resourceID == "" {
			t.Fatalf("GCP resource does not have a %v field", k8s.ResourceIDFieldPath)
		}
		testcontroller.SetResourceID(t, initialUnstruct, resourceID)
	}
	// autoCreateSubnetworks defaults to true, rather than false, and acts
	// as an end-to-end regression test for previous behavior defaulting to the wrong value
	// see: b/178744782
	if testContext.ResourceFixture.GVK.Kind == "ComputeNetwork" {
		unstructured.RemoveNestedField(initialUnstruct.Object, "spec", "autoCreateSubnetworks")
	}
	if err := kubeClient.Create(ctx, initialUnstruct); err != nil {
		t.Fatalf("error creating resource: %v", err)
	}
	preReconcileGeneration := initialUnstruct.GetGeneration()
	resourceCleanup := systemContext.Reconciler.BuildCleanupFunc(ctx, initialUnstruct, getResourceCleanupPolicy())
	defer resourceCleanup()
	systemContext.Reconciler.Reconcile(ctx, initialUnstruct, testreconciler.ExpectedSuccessfulReconcileResultFor(systemContext.Reconciler, initialUnstruct), nil)

	// Check labels match
	if resourceContext.SupportsLabels(systemContext.SMLoader, initialUnstruct) {
		gcpUnstruct, err := resourceContext.Get(ctx, t, initialUnstruct, systemContext.TFProvider, kubeClient, systemContext.SMLoader, systemContext.DCLConfig, systemContext.DCLConverter, nil)
		if err != nil {
			t.Fatalf("[testReconcileAcquire 2] unexpected error when GET-ing '%v': %v", initialUnstruct.GetName(), err)
		}
		testcontroller.AssertLabelsMatchAndHaveManagedLabel(t, gcpUnstruct.GetLabels(), initialUnstruct.GetLabels())
	}

	reconciledUnstruct := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       initialUnstruct.GetKind(),
			"apiVersion": initialUnstruct.GetAPIVersion(),
		},
	}
	if err := kubeClient.Get(ctx, testContext.NamespacedName, reconciledUnstruct); err != nil {
		t.Fatalf("unexpected error getting k8s resource: %v", err)
	}

	// Check that condition is ready and "UpToDate" event was recorded
	testcontroller.AssertReadyCondition(t, reconciledUnstruct, preReconcileGeneration)
	testcontroller.AssertEventRecordedforUnstruct(t, kubeClient, reconciledUnstruct, k8s.UpToDate)

	// Check observedGeneration matches with the pre-reconcile generation
	testcontroller.AssertObservedGenerationEquals(t, reconciledUnstruct, preReconcileGeneration)

	verifyResourceIDIfSupported(t, systemContext, resourceContext, reconciledUnstruct, initialUnstruct)
}

// TODO(b/174100391): Compare the resourceID of the retrieved GCP resource and the appliedUnstruct.
func verifyResourceIDIfSupported(t *testing.T, systemContext testrunner.SystemContext, resourceContext contexts.ResourceContext, reconciledUnstruct, appliedUnstruct *unstructured.Unstructured) {
	rt, err := testreconciler.ReconcilerTypeForObject(appliedUnstruct)
	if err != nil {
		t.Fatalf("error getting reconciler type: %v", err)
	}

	if rt == testreconciler.ReconcilerTypeDCL {
		s, found := dclextension.GetNameFieldSchema(resourceContext.DCLSchema)
		if !found {
			// The resource doesn't have a 'resourceID' field.
			return
		}
		isServerGeneratedID, err := dclextension.IsResourceIDFieldServerGenerated(s)
		if err != nil {
			t.Fatalf("error parsing `resourceID` field schema: %v", err)
		}
		verifyResourceID(t, isServerGeneratedID, reconciledUnstruct, appliedUnstruct)
	} else if rt == testreconciler.ReconcilerTypeTerraform {
		rc, err := systemContext.SMLoader.GetResourceConfig(reconciledUnstruct)
		if err != nil {
			t.Fatalf("error getting resource config for Kind '%s', "+
				"Namespace '%s', Name '%s'", reconciledUnstruct.GetKind(),
				reconciledUnstruct.GetNamespace(), reconciledUnstruct.GetName())
		}
		if !testcontroller.SupportsResourceIDField(rc) {
			return
		}
		isServerGeneratedId := testcontroller.IsResourceIDFieldServerGenerated(rc)
		verifyResourceID(t, isServerGeneratedId, reconciledUnstruct, appliedUnstruct)
	}
}

func verifyResourceID(t *testing.T, isServerGeneratedID bool, reconciledUnstruct, appliedUnstruct *unstructured.Unstructured) {
	reconciledResourceID, found := testcontroller.GetResourceID(t, reconciledUnstruct)
	if !found {
		t.Fatalf("'%s' not found", k8s.ResourceIDFieldPath)
	}
	if reconciledResourceID == "" {
		t.Fatalf("invalid value for '%s': empty string",
			k8s.ResourceIDFieldPath)
	}

	if isServerGeneratedID {
		testcontroller.AssertServerGeneratedResourceIDMatch(t, reconciledResourceID, appliedUnstruct)
		return
	}
	testcontroller.AssertUserSpecifiedResourceIDMatch(t, reconciledResourceID, appliedUnstruct)
}

func getResourceCleanupPolicy() testreconciler.ResourceCleanupPolicy {
	if cleanupResources {
		return testreconciler.CleanupPolicyAlways
	}
	return testreconciler.CleanupPolicyOnSuccess
}

func assertObjectContains(t *testing.T, obj, changedFields map[string]interface{}) {
	for changedKey, changedVal := range changedFields {
		objVal, ok := obj[changedKey]
		if !ok {
			t.Fatalf("object is missing the field %v", changedKey)
		}

		switch changedVal.(type) {
		case map[string]interface{}:
			if _, ok := objVal.(map[string]interface{}); !ok {
				t.Fatalf("expected object to have a map at %v, but got %v", changedKey, objVal)
			}
			assertObjectContains(t, objVal.(map[string]interface{}), changedVal.(map[string]interface{}))
		default:
			if diff := cmp.Diff(objVal, changedVal, cmpopts.SortSlices(
				func(a, b interface{}) bool {
					return fmt.Sprintf("%v", a) < fmt.Sprintf("%v", b)
				}),
			); diff != "" {
				t.Fatalf("unexpected diff: %v", diff)
			}
		}
	}
}

func getChangedFields(initialObject, updatedObject map[string]interface{}, field string) map[string]interface{} {
	changedFields := make(map[string]interface{})
	initial, _ := initialObject[field].(map[string]interface{})
	updated := updatedObject[field].(map[string]interface{})
	if !reflect.DeepEqual(initial, updated) {
		for k, v := range updated {
			if !reflect.DeepEqual(initial[k], v) {
				switch v.(type) {
				case map[string]interface{}:
					// Skip checking for changes in resource reference fields, because there
					// is no way to export the name and namespace fields (only external).
					if !strings.HasSuffix(k, "Ref") {
						changedFields[k] = getChangedFields(initial, updated, k)
					}
				default:
					// Skip checking for changes in fields contains a list of resource references,
					// because there is no way to export the name and namespace fields (only external).

					// Fields containing a list of resource references do not have the "Ref" suffix in the field name.
					// For example, targetResources field in ComputeFirewallPolicyRule.
					// Manually add those fields to the skipped fields list.
					// todo: Determine if we want to have "Refs" in those field names, like "targetResourceRefs"
					// That would introduce breaking changes to DCL/TF resource
					if updatedObject["kind"] == "ComputeFirewallPolicyRule" {
						computeFirewallPolicyRuleSkippedFields := []string{"targetResources", "targetServiceAccounts"}
						if !slices.Contains(computeFirewallPolicyRuleSkippedFields, k) {
							changedFields[k] = updated[k]
						}
					}
				}
			}
		}
	}
	return changedFields
}

// stripInternalKeysFromManagedFields strips fields that are used for internal
// tracking of managed fields by the api-server. This is useful when trying to
// differentiate between changes made by the server vs changed made by our controllers.
//
// note: this will modify the struct permanently, and will make the schema for
// managedFields invalid. Do not use this method without a deepcopy if you are
// planning on extracting the ManagedField struct later.
func stripInternalKeysFromManagedFields(t *testing.T, unstruct *unstructured.Unstructured) {
	managedFields, found, err := unstructured.NestedFieldNoCopy(unstruct.Object, "metadata", "managedFields")
	if err != nil {
		t.Fatalf("error getting managed fields: %v", err)
	}
	if !found {
		return
	}
	for _, managedField := range managedFields.([]interface{}) {
		managedField := managedField.(map[string]interface{})
		delete(managedField, "time")
	}
}

func containsResourceIDTestVar(t *testing.T, u *unstructured.Unstructured) bool {
	b, err := yaml.Marshal(u)
	if err != nil {
		t.Fatalf("error marshalling unstruct to bytes: %v", err)
	}
	return strings.Contains(string(b), resourceIDTestVar)
}

func removeSensitiveFields(unstruct *unstructured.Unstructured, changedSpecFields map[string]interface{}) {
	switch unstruct.GetKind() {
	case "BigQueryDataTransferConfig":
		// remove these fields which are reacted by the GCP API when reading.
		unstructured.RemoveNestedField(changedSpecFields, "params", "connector.authentication.oauth.clientId")
		unstructured.RemoveNestedField(changedSpecFields, "params", "connector.authentication.oauth.clientSecret")
	}
}

func TestMain(m *testing.M) {
	testmain.ForIntegrationTests(m, &mgr)
}
