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

package iamclient_test

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"

	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testreconciler "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller/reconciler"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testiam "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/iam"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testrunner "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/runner"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func init() {
	// run-tests and skip-tests allows you to limit the tests that are run by
	// specifying regexes to be used to match test names. See the
	// formatTestName() function to see what test names look like.
	flag.StringVar(&runTestsRegex, "run-tests", "", "run only the tests whose names match the given regex")
	flag.StringVar(&skipTestsRegex, "skip-tests", "", "skip the tests whose names match the given regex, even those that match the run-tests regex")
}

var (
	mgr            manager.Manager
	runTestsRegex  string
	skipTestsRegex string
)

func shouldRunBasedOnRunAndSkipRegexes(parentTestName string, fixture resourcefixture.ResourceFixture) bool {
	testName := fmt.Sprintf("%v/%v", parentTestName, resourcefixture.FormatTestName(fixture))

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

func TestAllGetSetDeletePolicy(t *testing.T) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	serviceMetadataLoader := dclmetadata.New()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating a new DCL schema loader: %v", err)
	}
	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && testiam.FixtureSupportsIAMPolicy(t, smLoader, serviceMetadataLoader, dclSchemaLoader, fixture)
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		refResource := testContext.CreateUnstruct
		resourceRef := testiam.NewResourceRef(refResource)
		iamPolicy := newPolicy(t, refResource, resourceRef, testContext.UniqueID)
		testGetSetDeletePolicy(ctx, t, iamClient, iamPolicy, refResource.GetKind())
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testCaseFunc)
}

func TestAllGetSetDeletePolicyWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	serviceMetadataLoader := dclmetadata.New()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating a new DCL schema loader: %v", err)
	}
	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && testiam.ShouldRunWithExternalRef(fixture) && testiam.FixtureSupportsIAMPolicy(t, smLoader, serviceMetadataLoader, dclSchemaLoader, fixture)
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		refResource := testContext.CreateUnstruct
		resourceRef, err := testiam.NewExternalRef(refResource, sysContext.TFProvider, sysContext.SMLoader)
		if err != nil {
			t.Fatal(err)
		}
		iamPolicy := newPolicy(t, refResource, resourceRef, testContext.UniqueID)
		testGetSetDeletePolicy(ctx, t, iamClient, iamPolicy, refResource.GetKind())
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testCaseFunc)
}

func TestAllGetSetDeletePolicyWithExternalOnlyRef(t *testing.T) {
	ctx := context.TODO()

	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.IAMExternalOnlyRef && fixture.GVK.Kind == "IAMPolicy"
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		iamPolicy := &iamv1beta1.IAMPolicy{}
		if err := util.Marshal(testContext.CreateUnstruct, iamPolicy); err != nil {
			t.Fatalf("error marshaling create unstruct into IAMPolicy object: %v", err)
		}

		testGetSetDeletePolicy(ctx, t, iamClient, iamPolicy, iamPolicy.Spec.ResourceReference.Kind)
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testCaseFunc)
}

func TestAllGetSetDeletePolicyWithIAMCondition(t *testing.T) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	serviceMetadataLoader := dclmetadata.New()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating a new DCL schema loader: %v", err)
	}
	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && testiam.ShouldRunWithIAMConditions(fixture) && testiam.FixtureSupportsIAMPolicy(t, smLoader, serviceMetadataLoader, dclSchemaLoader, fixture)
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		refResource := testContext.CreateUnstruct
		resourceRef := testiam.NewResourceRef(refResource)
		iamPolicy := newPolicyWithIAMConditions(t, refResource, resourceRef, testContext.UniqueID)
		testGetSetDeletePolicy(ctx, t, iamClient, iamPolicy, refResource.GetKind())
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testCaseFunc)
}

func TestAllGetSetPolicyWithAuditConfigs(t *testing.T) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	serviceMetadataLoader := dclmetadata.New()
	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && testiam.FixtureSupportsIAMAuditConfigs(t, smLoader, serviceMetadataLoader, fixture)
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		refResource := testContext.CreateUnstruct
		resourceRef := testiam.NewResourceRef(refResource)
		iamPolicy := newPolicyWithAuditConfigs(t, refResource, resourceRef, testContext.UniqueID)
		testGetSetDeletePolicy(ctx, t, iamClient, iamPolicy, refResource.GetKind())
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testCaseFunc)
}

func testGetSetDeletePolicy(ctx context.Context, t *testing.T, iamClient *kcciamclient.IAMClient, newPolicy *iamv1beta1.IAMPolicy, refResourceKind string) {
	gcpPolicy, err := iamClient.GetPolicy(ctx, newPolicy)
	if err != nil {
		t.Fatalf("error getting iam policy: %v", err)
	}
	// After a fresh GET of a newly-created resource, ensure it has the expected default
	// binding count
	defaultNumBindings := 0
	switch refResourceKind {
	case "StorageBucket":
		defaultNumBindings = 2
	case "Project":
		defaultNumBindings = 1
	case "Folder":
		defaultNumBindings = 2
	case "Organization":
		defaultNumBindings = 3
	case "BillingAccount":
		defaultNumBindings = 1
	}
	if len(gcpPolicy.Spec.Bindings) != defaultNumBindings {
		t.Errorf("unexpected length of bindings: got '%v', want '%v'", len(gcpPolicy.Spec.Bindings), defaultNumBindings)
	}
	if len(gcpPolicy.Spec.AuditConfigs) != 0 {
		t.Errorf("unexpected length of audit configs: got '%v', want '%v'", len(gcpPolicy.Spec.AuditConfigs), 0)
	}
	if !reflect.DeepEqual(gcpPolicy.Spec.ResourceReference, newPolicy.Spec.ResourceReference) {
		t.Errorf("resource reference mismatch: got '%v', want '%v'", gcpPolicy.Spec.ResourceReference, newPolicy.Spec.ResourceReference)
	}
	resultPolicy, err := iamClient.SetPolicy(ctx, newPolicy)
	if err != nil {
		t.Errorf("error setting iam policy: %v", err)
	}
	if resultPolicy == nil {
		t.Fatalf("unexpected nil value for resultPolicy")
	}
	if !reflect.DeepEqual(resultPolicy.Spec.Bindings, newPolicy.Spec.Bindings) {
		t.Errorf("mismatched bindings: got '%v', want '%v'", resultPolicy.Spec.Bindings, newPolicy.Spec.Bindings)
	}
	if !reflect.DeepEqual(resultPolicy.Spec.AuditConfigs, newPolicy.Spec.AuditConfigs) {
		t.Errorf("mismatched audit configs: got '%v', want '%v'", resultPolicy.Spec.AuditConfigs, newPolicy.Spec.AuditConfigs)
	}
	gcpPolicy, err = iamClient.GetPolicy(ctx, newPolicy)
	if err != nil {
		t.Fatalf("error getting iam policy: %v", err)
	}
	// The Bindings and AuditConfigs of the IAMPolicy we GET will not
	// necessarily be in the same order as in the IAMPolicy we set, so use
	// testiam.SameBindings() and testiam.SameAuditConfigs() to check that the
	// Bindings and AuditConfigs are the same instead of reflect.DeepEqual().
	if !testiam.SameBindings(newPolicy.Spec.Bindings, gcpPolicy.Spec.Bindings) {
		t.Errorf("mismatched bindings: got '%v', want '%v'", gcpPolicy.Spec.Bindings, newPolicy.Spec.Bindings)
	}
	if !testiam.SameAuditConfigs(newPolicy.Spec.AuditConfigs, gcpPolicy.Spec.AuditConfigs) {
		t.Errorf("mismatched auditConfigs: got '%v', want '%v'", gcpPolicy.Spec.AuditConfigs, newPolicy.Spec.AuditConfigs)
	}
	// Skip testing deletion of IAMPolicy for StorageBuckets, BillingAccount
	// and Organization since you need IAM permissions on them to be able to
	// manage their IAM.
	switch refResourceKind {
	case "StorageBucket", "Organization", "BillingAccount":
		return
	}
	if err := iamClient.DeletePolicy(ctx, newPolicy); err != nil {
		t.Fatalf("error deleting: %v", err)
	}
	// "Deleting" an IAMPolicy just means clearing it. Check that we can still
	// GET the IAMPolicy and that it has been cleared.
	gcpPolicy, err = iamClient.GetPolicy(ctx, newPolicy)
	if err != nil {
		t.Fatalf("error getting iam policy: %v", err)
	}
	if len(gcpPolicy.Spec.Bindings) > 0 {
		t.Errorf("expected there to be no bindings but there were %v", len(gcpPolicy.Spec.Bindings))
	}
}

func TestResolveMemberIdentity(t *testing.T) {
	ctx := context.TODO()

	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, _ manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && fixture.GVK.Kind == "IAMServiceAccount"
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		testResolveMemberIdentity(ctx, t, iamClient, testContext)
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testCaseFunc)
}

func testResolveMemberIdentity(ctx context.Context, t *testing.T, iamClient *kcciamclient.IAMClient, testContext testrunner.TestContext) {
	member := iamv1beta1.Member("user:user2@example.com")

	resolvedId, err := kcciamclient.ResolveMemberIdentity(ctx, member, nil, testContext.NamespacedName.Namespace, iamClient.TFIAMClient)
	if err != nil {
		t.Fatalf("error resolving member identity with member")
	}

	if resolvedId != string(member) {
		t.Fatalf("error resolving member identity with member, expect: %v, received: %v", string(member), resolvedId)
	}

	memberFrom := &iamv1beta1.MemberSource{
		ServiceAccountRef: &iamv1beta1.MemberReference{
			Namespace: "cnrm-foo",
			Name:      "cnrm-sa",
		},
	}
	_, err = kcciamclient.ResolveMemberIdentity(ctx, "", memberFrom, testContext.NamespacedName.Namespace, iamClient.TFIAMClient)
	if err == nil || !k8s.IsReferenceNotFoundError(err) {
		t.Fatalf("resolving member identity with nonexistent memberFrom reference: got error: %v, want ReferenceNotFoundError", err)
	}

	memberFrom.ServiceAccountRef.Name = testContext.NamespacedName.Name
	memberFrom.ServiceAccountRef.Namespace = testContext.NamespacedName.Namespace

	// TODO: We want to figure out a way to verify resolved id for memberFrom
	_, err = kcciamclient.ResolveMemberIdentity(ctx, "", memberFrom, testContext.NamespacedName.Namespace, iamClient.TFIAMClient)
	if err != nil {
		t.Fatalf("error resolving member identity with valid memberFrom")
	}

	_, err = kcciamclient.ResolveMemberIdentity(ctx, member, memberFrom, testContext.NamespacedName.Namespace, iamClient.TFIAMClient)
	if err == nil {
		t.Fatalf("got no error when ResolveMemberIdentity() was given both a valid member and memberFrom; want error")
	}
}

func TestGetSetDeletePolicyReferenceNotFound(t *testing.T) {
	ctx := context.TODO()

	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, _ manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && fixture.GVK.Kind == "PubSubTopic"
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		refResource := testContext.CreateUnstruct
		resourceRef := testiam.NewResourceRef(refResource)
		newPolicy := newPolicy(t, refResource, resourceRef, testContext.UniqueID)
		testGetSetDeletePolicyReferenceNotFound(ctx, t, iamClient, newPolicy)
	}
	testrunner.RunAllWithDependenciesCreatedButNotObject(ctx, t, mgr, shouldRun, testCaseFunc)
}

func testGetSetDeletePolicyReferenceNotFound(ctx context.Context, t *testing.T, iamClient *kcciamclient.IAMClient, newPolicy *iamv1beta1.IAMPolicy) {
	_, err := iamClient.GetPolicy(ctx, newPolicy)
	if !k8s.IsReferenceNotFoundError(err) {
		t.Errorf("getting policy when referenced resource not found: got error: %v, want ReferenceNotFoundError", err)
	}
	_, err = iamClient.SetPolicy(ctx, newPolicy)
	if !k8s.IsReferenceNotFoundError(err) {
		t.Errorf("setting policy when referenced resource not found: got error: %v, want ReferenceNotFoundError", err)
	}
	err = iamClient.DeletePolicy(ctx, newPolicy)
	if !k8s.IsReferenceNotFoundError(err) {
		t.Errorf("deleting policy when referenced resource not found: got error: %v, want ReferenceNotFoundError", err)
	}
}

func TestProjectIdAsNamespace(t *testing.T) {
	ctx := context.TODO()

	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		// this test does not load dependencies and we only need to verify that this functionality works
		// for a single resource.
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && fixture.GVK.Kind == "IAMServiceAccount"
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, tstCtx testrunner.TestContext, sysCtx testrunner.SystemContext) {
		projectID := testgcp.GetDefaultProjectID(t)
		testcontroller.SetupNamespaceForDefaultProject(t, sysCtx.Manager.GetClient(), projectID)
		refResource := tstCtx.CreateUnstruct
		refResource.SetNamespace(projectID)
		if err := sysCtx.Manager.GetClient().Create(ctx, refResource); err != nil {
			t.Fatalf("error creating resource: %v", err)
		}
		resourceCleanup := sysCtx.Reconciler.BuildCleanupFunc(ctx, refResource, testreconciler.CleanupPolicyAlways)
		defer resourceCleanup()
		sysCtx.Reconciler.Reconcile(ctx, refResource, reconcile.Result{RequeueAfter: k8s.MeanReconcileReenqueuePeriod}, nil)
		iamClient := testiam.NewIAMClient(sysCtx)
		resourceRef := testiam.NewResourceRef(refResource)
		newPolicy := newPolicy(t, refResource, resourceRef, tstCtx.UniqueID)
		policy, err := iamClient.GetPolicy(ctx, newPolicy)
		if err != nil {
			t.Fatalf("error getting iam policy: %v", err)
		}
		if policy.Spec.Bindings != nil {
			t.Errorf("unexpected value for bindings: got '%v' want 'nil'", policy.Spec.Bindings)
		}
		resultPolicy, err := iamClient.SetPolicy(ctx, newPolicy)
		if err != nil {
			t.Errorf("error setting iam policy: %v", err)
		}
		if resultPolicy == nil {
			t.Fatalf("unexpected nil value for resultPolicy")
		}
		if !reflect.DeepEqual(resultPolicy.Spec.Bindings, newPolicy.Spec.Bindings) {
			t.Errorf("mismatched bindings: got '%v', want '%v'", resultPolicy.Spec.Bindings, newPolicy.Spec.Bindings)
		}
	}
	testrunner.RunAllWithDependenciesCreatedButNotObject(ctx, t, mgr, shouldRun, testCaseFunc)
}

func TestAllGetSetDeletePolicyMember(t *testing.T) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	serviceMetadataLoader := dclmetadata.New()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating a new DCL schema loader: %v", err)
	}
	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && testiam.FixtureSupportsIAMPolicy(t, smLoader, serviceMetadataLoader, dclSchemaLoader, fixture)
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		refResource := testContext.CreateUnstruct
		resourceRef := testiam.NewResourceRef(refResource)
		iamPolicyMember := newPolicyMember(t, refResource, resourceRef, testContext.UniqueID)
		testGetSetDeleteIamPolicyMember(ctx, t, iamClient, iamPolicyMember, refResource)
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testCaseFunc)
}

func TestAllGetSetDeletePolicyMemberWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	serviceMetadataLoader := dclmetadata.New()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating a new DCL schema loader: %v", err)
	}
	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && testiam.ShouldRunWithExternalRef(fixture) && testiam.FixtureSupportsIAMPolicy(t, smLoader, serviceMetadataLoader, dclSchemaLoader, fixture)
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		refResource := testContext.CreateUnstruct
		resourceRef, err := testiam.NewExternalRef(refResource, sysContext.TFProvider, sysContext.SMLoader)
		if err != nil {
			t.Fatal(err)
		}
		iamPolicyMember := newPolicyMember(t, refResource, resourceRef, testContext.UniqueID)
		testGetSetDeleteIamPolicyMember(ctx, t, iamClient, iamPolicyMember, refResource)
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testCaseFunc)
}

func TestAllGetSetDeletePolicyMemberWithIAMCondition(t *testing.T) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	serviceMetadataLoader := dclmetadata.New()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating a new DCL schema loader: %v", err)
	}
	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && testiam.ShouldRunWithIAMConditions(fixture) && testiam.FixtureSupportsIAMPolicy(t, smLoader, serviceMetadataLoader, dclSchemaLoader, fixture)
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		refResource := testContext.CreateUnstruct
		resourceRef := testiam.NewResourceRef(refResource)
		iamPolicyMember := newPolicyMemberWithIAMCondition(t, refResource, resourceRef, testContext.UniqueID)
		testGetSetDeleteIamPolicyMember(ctx, t, iamClient, iamPolicyMember, refResource)
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testCaseFunc)
}

func TestAllGetSetDeletePolicyMemberWithExternalOnlyRef(t *testing.T) {
	ctx := context.TODO()

	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.IAMExternalOnlyRef && fixture.GVK.Kind == "IAMPolicyMember"
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		iamPolicyMember := &iamv1beta1.IAMPolicyMember{}
		if err := util.Marshal(testContext.CreateUnstruct, iamPolicyMember); err != nil {
			t.Fatalf("error marshaling create unstruct into IAMPolicyMember object: %v", err)
		}
		testGetSetDeleteIamPolicyMember(ctx, t, iamClient, iamPolicyMember, nil)
	}
	testrunner.RunAllWithDependenciesCreatedButNotObject(ctx, t, mgr, shouldRun, testCaseFunc)
}

func TestAllGetSetDeletePolicyMemberWithMemberReference(t *testing.T) {
	ctx := context.TODO()

	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.IAMMemberReferences
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		iamPolicyMember := &iamv1beta1.IAMPolicyMember{}
		if err := util.Marshal(testContext.CreateUnstruct, iamPolicyMember); err != nil {
			t.Fatalf("error marshaling create unstruct into IAMPolicyMember object: %v", err)
		}
		testGetSetDeleteIamPolicyMember(ctx, t, iamClient, iamPolicyMember, nil)
	}
	testrunner.RunAllWithDependenciesCreatedButNotObject(ctx, t, mgr, shouldRun, testCaseFunc)
}

func testGetSetDeleteIamPolicyMember(ctx context.Context, t *testing.T, iamClient *kcciamclient.IAMClient, policyMember *iamv1beta1.IAMPolicyMember, refResource *unstructured.Unstructured) {
	_, err := iamClient.GetPolicyMember(ctx, policyMember)
	if err == nil {
		t.Fatalf("expected an error when retrieving IAMPolicyMember, instead got 'nil'")
	}
	if !isNotFoundError(err) {
		t.Fatalf("unexpected error when retrieving IAMPolicyMember: got '%v', want '%v'", err, kcciamclient.ErrNotFound)
	}
	gcpPolicyMember, err := iamClient.SetPolicyMember(ctx, policyMember)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gcpPolicyMember.Spec.Member != policyMember.Spec.Member {
		t.Errorf("unexpected member value: got '%v', want '%v'",
			gcpPolicyMember.Spec.Member, policyMember.Spec.Member)
	}
	if !reflect.DeepEqual(gcpPolicyMember.Spec.MemberFrom, policyMember.Spec.MemberFrom) {
		t.Errorf("memberFrom mismatch: got '%v', want '%v'",
			gcpPolicyMember.Spec.MemberFrom, policyMember.Spec.MemberFrom)
	}
	if gcpPolicyMember.Spec.Role != policyMember.Spec.Role {
		t.Errorf("unexpected role value: got '%v', want '%v'",
			gcpPolicyMember.Spec.Role, policyMember.Spec.Role)
	}
	if !reflect.DeepEqual(gcpPolicyMember.Spec.ResourceReference, policyMember.Spec.ResourceReference) {
		t.Errorf("resource reference mismatch: got '%v', want '%v'",
			gcpPolicyMember.Spec.ResourceReference, policyMember.Spec.ResourceReference)
	}
	_, err = iamClient.GetPolicyMember(ctx, policyMember)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := deletePolicyMemberWithRetries(ctx, t, iamClient, gcpPolicyMember); err != nil {
		t.Fatalf("error deleting: %v", err)
	}
	_, err = iamClient.GetPolicyMember(ctx, policyMember)
	if err == nil {
		t.Fatalf("expected an error when retrieving IAMPolicyMember, instead got 'nil'")
	}
	if !isNotFoundError(err) {
		t.Fatalf("unexpected error when retrieving IAMPolicyMember: got '%v', want '%v'", err, kcciamclient.ErrNotFound)
	}
}

func deletePolicyMemberWithRetries(ctx context.Context, t *testing.T, iamClient *kcciamclient.IAMClient, policyMember *iamv1beta1.IAMPolicyMember) error {
	attempt := 0
	maxAttempts := 3
	for {
		attempt++
		err := iamClient.DeletePolicyMember(ctx, policyMember)
		if err == nil {
			return err
		}
		msg := err.Error()
		if strings.Contains(msg, "concurrent policy changes") {
			if attempt >= maxAttempts {
				t.Logf("found concurrent policy changes error, but max attempts reached: %v", err)
				return err
			}

			t.Logf("found concurrent policy changes error, will retry: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		return err
	}
}

func TestGetSetDeletePolicyMemberReferenceNotFound(t *testing.T) {
	ctx := context.TODO()

	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, _ manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && fixture.GVK.Kind == "PubSubTopic"
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		refResource := testContext.CreateUnstruct
		resourceRef := testiam.NewResourceRef(refResource)
		iamPolicyMember := newPolicyMemberWithIAMCondition(t, refResource, resourceRef, testContext.UniqueID)
		testGetSetDeletePolicyMemberReferenceNotFound(ctx, t, iamClient, iamPolicyMember)
	}
	testrunner.RunAllWithDependenciesCreatedButNotObject(ctx, t, mgr, shouldRun, testCaseFunc)
}

func testGetSetDeletePolicyMemberReferenceNotFound(ctx context.Context, t *testing.T, iamClient *kcciamclient.IAMClient, policyMember *iamv1beta1.IAMPolicyMember) {
	_, err := iamClient.GetPolicyMember(ctx, policyMember)
	if !k8s.IsReferenceNotFoundError(err) {
		t.Errorf("getting policy member when referenced resource not found: got error: %v, want ReferenceNotFoundError", err)
	}
	_, err = iamClient.SetPolicyMember(ctx, policyMember)
	if !k8s.IsReferenceNotFoundError(err) {
		t.Errorf("setting policy member when referenced resource not found: got error: %v, want ReferenceNotFoundError", err)
	}
	err = iamClient.DeletePolicyMember(ctx, policyMember)
	if !k8s.IsReferenceNotFoundError(err) {
		t.Errorf("deleting policy member when referenced resource not found: got error: %v, want ReferenceNotFoundError", err)
	}
}

func TestAllGetSetDeleteAuditConfig(t *testing.T) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	serviceMetadataLoader := dclmetadata.New()
	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && testiam.FixtureSupportsIAMAuditConfigs(t, smLoader, serviceMetadataLoader, fixture)
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		refResource := testContext.CreateUnstruct
		resourceRef := testiam.NewResourceRef(refResource)
		iamAuditConfig := newAuditConfig(t, refResource, resourceRef, testContext.UniqueID)
		testGetSetDeleteIamAuditConfig(ctx, t, iamClient.TFIAMClient, iamAuditConfig, refResource)
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testCaseFunc)
}

func TestAllGetSetDeleteAuditConfigWithExternalRef(t *testing.T) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	serviceMetadataLoader := dclmetadata.New()
	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && testiam.ShouldRunWithExternalRef(fixture) && testiam.FixtureSupportsIAMAuditConfigs(t, smLoader, serviceMetadataLoader, fixture)
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		refResource := testContext.CreateUnstruct
		resourceRef, err := testiam.NewExternalRef(refResource, sysContext.TFProvider, sysContext.SMLoader)
		if err != nil {
			t.Fatal(err)
		}
		iamAuditConfig := newAuditConfig(t, refResource, resourceRef, testContext.UniqueID)
		testGetSetDeleteIamAuditConfig(ctx, t, iamClient.TFIAMClient, iamAuditConfig, refResource)
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testCaseFunc)
}

func testGetSetDeleteIamAuditConfig(ctx context.Context, t *testing.T, iamClient *kcciamclient.TFIAMClient, auditConfig *iamv1beta1.IAMAuditConfig, refResource *unstructured.Unstructured) {
	_, err := iamClient.GetAuditConfig(ctx, auditConfig)
	if err == nil {
		t.Fatalf("expected an error when retrieving IAMAuditConfig, instead got 'nil'")
	}
	if !errors.Is(err, kcciamclient.ErrNotFound) {
		t.Fatalf("unexpected error when retrieving IAMAuditConfig: got '%v', want '%v'", err, kcciamclient.ErrNotFound)
	}
	gcpAuditConfig, err := iamClient.SetAuditConfig(ctx, auditConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if auditConfig == gcpAuditConfig {
		t.Fatalf("expected auditConfig pointer values to not be equal")
	}
	if gcpAuditConfig.Spec.Service != auditConfig.Spec.Service {
		t.Errorf("service value mismatch: got '%v', want '%v'", gcpAuditConfig.Spec.Service, auditConfig.Spec.Service)
	}
	if !reflect.DeepEqual(auditConfig.Spec.AuditLogConfigs, gcpAuditConfig.Spec.AuditLogConfigs) {
		t.Errorf("audit log config mismatch: got '%v', want '%v'", gcpAuditConfig.Spec.AuditLogConfigs, auditConfig.Spec.AuditLogConfigs)
	}
	if !reflect.DeepEqual(auditConfig.Spec.ResourceReference, gcpAuditConfig.Spec.ResourceReference) {
		t.Errorf("resource reference mismatch: got '%v', want '%v'", gcpAuditConfig.Spec.ResourceReference, auditConfig.Spec.ResourceReference)
	}
	_, err = iamClient.GetAuditConfig(ctx, auditConfig)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := iamClient.DeleteAuditConfig(ctx, gcpAuditConfig); err != nil {
		t.Fatalf("error deleting: %v", err)
	}
	_, err = iamClient.GetAuditConfig(ctx, auditConfig)
	if err == nil {
		t.Fatalf("expected an error when retrieving IAMAuditConfig, instead got 'nil'")
	}
	if !errors.Is(err, kcciamclient.ErrNotFound) {
		t.Fatalf("unexpected error when retrieving IAMAuditConfig: got '%v', want '%v'", err, kcciamclient.ErrNotFound)
	}
}

func TestGetSetDeleteAuditConfigReferenceNotFound(t *testing.T) {
	ctx := context.TODO()

	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, _ manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) && fixture.Type == resourcefixture.Basic && fixture.GVK.Kind == "Project"
	}
	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		refResource := testContext.CreateUnstruct
		resourceRef := testiam.NewResourceRef(refResource)
		iamAuditConfig := newAuditConfig(t, refResource, resourceRef, testContext.UniqueID)
		testGetSetDeleteAuditConfigReferenceNotFound(ctx, t, iamClient.TFIAMClient, iamAuditConfig)
	}
	testrunner.RunAllWithDependenciesCreatedButNotObject(ctx, t, mgr, shouldRun, testCaseFunc)
}

func testGetSetDeleteAuditConfigReferenceNotFound(ctx context.Context, t *testing.T, iamClient *kcciamclient.TFIAMClient, auditConfig *iamv1beta1.IAMAuditConfig) {
	_, err := iamClient.GetAuditConfig(ctx, auditConfig)
	if !k8s.IsReferenceNotFoundError(err) {
		t.Errorf("getting audit config when referenced resource not found: got error: %v, want ReferenceNotFoundError", err)
	}
	_, err = iamClient.SetAuditConfig(ctx, auditConfig)
	if !k8s.IsReferenceNotFoundError(err) {
		t.Errorf("setting audit config when referenced resource not found: got error: %v, want ReferenceNotFoundError", err)
	}
	err = iamClient.DeleteAuditConfig(ctx, auditConfig)
	if !k8s.IsReferenceNotFoundError(err) {
		t.Errorf("deleting audit config when referenced resource not found: got error: %v, want ReferenceNotFoundError", err)
	}
}

func TestConflictPreventionWithEtags(t *testing.T) {
	ctx := context.TODO()

	testName := getCurrentFuncName()
	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		return shouldRunBasedOnRunAndSkipRegexes(testName, fixture) &&
			fixture.Type == resourcefixture.Basic &&
			(fixture.GVK.Kind == "PubSubTopic" || fixture.GVK.Kind == "Project" || fixture.GVK.Kind == "DataprocCluster")
	}

	testCaseFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		iamClient := testiam.NewIAMClient(sysContext)
		refResource := testContext.CreateUnstruct
		resourceRef := testiam.NewResourceRef(refResource)
		iamPolicy := newPolicy(t, refResource, resourceRef, testContext.UniqueID)
		testConflictPreventionWithEtags(ctx, t, iamClient, iamPolicy, refResource)
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testCaseFunc)
}

func testConflictPreventionWithEtags(ctx context.Context, t *testing.T, iamClient *kcciamclient.IAMClient, policy *iamv1beta1.IAMPolicy, refResource *unstructured.Unstructured) {
	gcpPolicy, err := iamClient.GetPolicy(ctx, policy)
	if err != nil {
		t.Fatalf("error getting iam policy: %v", err)
	}
	etag := gcpPolicy.Spec.Etag
	if etag == "" {
		t.Fatalf("returned etag should not be empty")
	}

	// Issue a no-ops SetPolicy request
	noChangePolicy, err := iamClient.SetPolicy(ctx, gcpPolicy)
	if err != nil {
		t.Fatalf("error setting iam policy: %v", err)
	}
	if noChangePolicy.Spec.Etag != etag {
		t.Fatalf("expect no change of etag since there is no realy IAM policy change")
	}

	// Set a new IAM policy
	rc := testiam.GetResourceContext(t, refResource.GetKind())
	newBinding := iamv1beta1.IAMPolicyBinding{
		Role:    rc.CreateBindingRole,
		Members: []iamv1beta1.Member{"group:configconnector-test@google.com"},
	}
	newPolicy := gcpPolicy.DeepCopy()
	newPolicy.Spec.Bindings = append(newPolicy.Spec.Bindings, newBinding)
	changedPolicy, err := iamClient.SetPolicy(ctx, newPolicy)
	if err != nil {
		t.Fatalf("error setting iam policy: %v", err)
	}
	if changedPolicy.Spec.Etag == etag {
		t.Fatalf("expect the returned etag changed")
	}

	// Setting the IAM policy with old etag should fail
	_, err = iamClient.SetPolicy(ctx, gcpPolicy)
	if err == nil {
		t.Fatalf("got no error, want an error related to concurrent policy changes")
	}
}

func newPolicyMemberWithIAMCondition(t *testing.T, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference, testID string) *iamv1beta1.IAMPolicyMember {
	iamPolicyMember := newPolicyMember(t, refResource, resourceRef, testID)
	iamPolicyMember.Spec.Condition = newIAMCondition()
	return iamPolicyMember
}

func newPolicyWithIAMConditions(t *testing.T, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference, testID string) *iamv1beta1.IAMPolicy {
	iamPolicy := newPolicy(t, refResource, resourceRef, testID)
	bindings := make([]iamv1beta1.IAMPolicyBinding, 0)
	for _, binding := range iamPolicy.Spec.Bindings {
		binding.Condition = newIAMCondition()
		bindings = append(bindings, binding)
	}
	iamPolicy.Spec.Bindings = bindings
	return iamPolicy
}

func newPolicyWithAuditConfigs(t *testing.T, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference, testID string) *iamv1beta1.IAMPolicy {
	iamPolicy := newPolicy(t, refResource, resourceRef, testID)
	iamPolicy.Spec.AuditConfigs = []iamv1beta1.IAMPolicyAuditConfig{
		{
			Service: "allServices",
			AuditLogConfigs: []iamv1beta1.AuditLogConfig{
				{
					LogType: "DATA_WRITE",
				},
				{
					LogType:         "DATA_READ",
					ExemptedMembers: []iamv1beta1.Member{iamv1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
				},
			},
		},
	}
	return iamPolicy
}

func newPolicyMember(t *testing.T, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference, testID string) *iamv1beta1.IAMPolicyMember {
	rc := testiam.GetResourceContext(t, refResource.GetKind())
	return &iamv1beta1.IAMPolicyMember{
		TypeMeta: metav1.TypeMeta{
			APIVersion: iamv1beta1.IAMPolicyMemberGVK.GroupVersion().String(),
			Kind:       iamv1beta1.IAMPolicyMemberGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("iam-policy-member-%v", testID),
			Namespace: refResource.GetNamespace(),
		},
		Spec: iamv1beta1.IAMPolicyMemberSpec{
			ResourceReference: resourceRef,
			Member:            iamv1beta1.Member(testgcp.GetIAMPolicyBindingMember(t)),
			Role:              rc.CreateBindingRole,
		},
	}
}

func newPolicy(t *testing.T, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference, testID string) *iamv1beta1.IAMPolicy {
	rc := testiam.GetResourceContext(t, refResource.GetKind())
	return &iamv1beta1.IAMPolicy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: iamv1beta1.IAMPolicyGVK.GroupVersion().String(),
			Kind:       iamv1beta1.IAMPolicyGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("iam-policy-%v", testID),
			Namespace: refResource.GetNamespace(),
		},
		Spec: iamv1beta1.IAMPolicySpec{
			ResourceReference: resourceRef,
			Bindings: []iamv1beta1.IAMPolicyBinding{
				{
					Members: []iamv1beta1.Member{
						iamv1beta1.Member(testgcp.GetIAMPolicyBindingMember(t)),
					},
					Role: rc.CreateBindingRole,
				},
			},
		},
		Status: iamv1beta1.IAMPolicyStatus{},
	}
}

func newIAMCondition() *iamv1beta1.IAMCondition {
	return &iamv1beta1.IAMCondition{
		Title:       "test-iam-condition",
		Description: "Test IAM Condition",
		Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
	}
}

func newAuditConfig(t *testing.T, refResource *unstructured.Unstructured, resourceRef iamv1beta1.ResourceReference, testID string) *iamv1beta1.IAMAuditConfig {
	return &iamv1beta1.IAMAuditConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: iamv1beta1.IAMAuditConfigGVK.GroupKind().String(),
			Kind:       iamv1beta1.IAMPolicyGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("iam-audit-config-%v", testID),
			Namespace: refResource.GetNamespace(),
		},
		Spec: iamv1beta1.IAMAuditConfigSpec{
			ResourceReference: resourceRef,
			Service:           "allServices",
			AuditLogConfigs: []iamv1beta1.AuditLogConfig{
				{
					LogType: "DATA_WRITE",
				},
				{
					LogType:         "DATA_READ",
					ExemptedMembers: []iamv1beta1.Member{iamv1beta1.Member(testgcp.GetIAMPolicyBindingMember(t))},
				},
			},
		},
	}
}

func getCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}

// TF-based resource Member gets LiveState and returns NotFoundError if the
// LiveState is empty. DCL-based resource Member gets from the API and
// returns a 404 if the Member does not exist.
func isNotFoundError(err error) bool {
	return errors.Is(err, kcciamclient.ErrNotFound) || dcl.IsNotFound(errors.Unwrap(err))
}

func TestMain(m *testing.M) {
	testmain.ForIntegrationTests(m, &mgr)
}
