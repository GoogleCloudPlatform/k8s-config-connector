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

package leaser_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/lease/leaser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/managementconflict"
	testreconciler "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller/reconciler"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testrunner "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/runner"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	mgr                  manager.Manager
	defaultLeaseDuration = 5 * time.Minute
)

func TestUnsupportedResourceShouldFail(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		switch fixture.GVK.Kind {
		case "SourceRepoRepository", // Resource with no labels field
			"DataflowJob": // Resource for which leasing is explicitly disabled (special case).
			return true
		default:
			return false
		}
	}
	testFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext) {
		leaser := leaser.NewLeaser(systemContext.TFProvider, systemContext.SMLoader, systemContext.Manager.GetClient())
		ok, err := leaser.UnstructuredSupportsLeasing(testContext.CreateUnstruct)
		if err != nil {
			t.Fatalf("error checking for lease support: %v", err)
		}
		if ok {
			t.Fatalf("test should only be run on resources that do not support locking")
		}
		err = leaser.Obtain(context.Background(), testContext.CreateUnstruct, "my-owner-id", defaultLeaseDuration, defaultLeaseDuration)
		if err == nil {
			t.Fatalf("expected error, instead got nil")
		}
		expectedMsg := fmt.Sprintf("gvk '%v' does not support locking", testContext.CreateUnstruct.GroupVersionKind())
		errMsg := err.Error()
		if errMsg != expectedMsg {
			t.Fatalf("unexpected error message: got '%v', want '%v'", errMsg, expectedMsg)
		}
	}
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testFunc)
}

func TestAll(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		switch fixture.GVK.Kind {
		case "Project", "PubSubTopic":
			return true
		default:
			return false
		}
	}
	testFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext) {
		// By default, the test framework creates resources with the management conflict policy of 'resource'. If we allow that,
		// then the namespace will have a lease on the resource. That would interfere with the results of this test as we want to
		// obtain and release the lease for our 'owners' that we generate below. For that reason, we set the conflict policy to
		// 'none' and create and reconcile the resource below.
		k8s.SetAnnotation(managementconflict.FullyQualifiedAnnotation, managementconflict.ManagementConflictPreventionPolicyNone, testContext.CreateUnstruct)
		if err := systemContext.Manager.GetClient().Create(ctx, testContext.CreateUnstruct); err != nil {
			t.Fatalf("error creating resource: %v", err)
		}
		resourceCleanup := systemContext.Reconciler.BuildCleanupFunc(ctx, testContext.CreateUnstruct, testreconciler.CleanupPolicyAlways)
		defer resourceCleanup()
		systemContext.Reconciler.Reconcile(ctx, testContext.CreateUnstruct, testreconciler.ExpectedSuccessfulReconcileResultFor(systemContext.Reconciler, testContext.CreateUnstruct), nil)
		leaser := leaser.NewLeaser(systemContext.TFProvider, systemContext.SMLoader, systemContext.Manager.GetClient())
		uniqueId1 := fmt.Sprintf("l1-%v", testContext.UniqueID)
		uniqueId2 := fmt.Sprintf("l2-%v", testContext.UniqueID)
		initialUnstruct := testContext.CreateUnstruct
		testObtainReleaseShouldSucceed(t, initialUnstruct, uniqueId1, leaser)
		testObtainTwiceShouldSucceed(t, initialUnstruct, uniqueId1, leaser)
		testReleaseUnobtainedShouldFail(t, initialUnstruct, uniqueId1, leaser)
		testReleasingExpiredResourceShouldFail(t, initialUnstruct, uniqueId1, leaser)
		testRenewLease(t, initialUnstruct, uniqueId1, leaser)
		testObtainingPreviouslyReleasedResourceShouldSucceed(t, initialUnstruct, uniqueId1, uniqueId2, leaser)
		testObtainingLockedResourceShouldFail(t, initialUnstruct, uniqueId1, uniqueId2, leaser)
		testReleasingLockedResourceShouldFail(t, initialUnstruct, uniqueId1, uniqueId2, leaser)
		testObtainingExpiredLeaseShouldSucceed(t, initialUnstruct, uniqueId1, uniqueId2, leaser)
	}
	testrunner.RunAllWithDependenciesCreatedButNotObject(ctx, t, mgr, shouldRun, testFunc)
}

func testObtainReleaseShouldSucceed(t *testing.T, u *unstructured.Unstructured, uniqueID string, leaser *leaser.Leaser) {
	obtainAssertSuccess(t, u, uniqueID, defaultLeaseDuration, defaultLeaseDuration, leaser)
	releaseAssertSuccess(t, u, uniqueID, leaser)
}

func testObtainTwiceShouldSucceed(t *testing.T, u *unstructured.Unstructured, uniqueID string, leaser *leaser.Leaser) {
	obtainAssertSuccess(t, u, uniqueID, defaultLeaseDuration, defaultLeaseDuration, leaser)
	obtainAssertSuccess(t, u, uniqueID, defaultLeaseDuration, defaultLeaseDuration, leaser)
	releaseAssertSuccess(t, u, uniqueID, leaser)
}

func testReleaseUnobtainedShouldFail(t *testing.T, u *unstructured.Unstructured, uniqueID string, leaser *leaser.Leaser) {
	releaseAssertError(t, u, uniqueID, leaser)
}

func testObtainingPreviouslyReleasedResourceShouldSucceed(t *testing.T, u *unstructured.Unstructured, uniqueId1, uniqueId2 string, leaser *leaser.Leaser) {
	obtainAssertSuccess(t, u, uniqueId1, defaultLeaseDuration, defaultLeaseDuration, leaser)
	releaseAssertSuccess(t, u, uniqueId1, leaser)
	obtainAssertSuccess(t, u, uniqueId2, defaultLeaseDuration, defaultLeaseDuration, leaser)
	releaseAssertSuccess(t, u, uniqueId2, leaser)
}

func testObtainingLockedResourceShouldFail(t *testing.T, u *unstructured.Unstructured, uniqueId1, uniqueId2 string, leaser *leaser.Leaser) {
	obtainAssertSuccess(t, u, uniqueId1, defaultLeaseDuration, defaultLeaseDuration, leaser)
	obtainAssertError(t, u, uniqueId2, defaultLeaseDuration, defaultLeaseDuration, leaser)
	releaseAssertSuccess(t, u, uniqueId1, leaser)
}

func testReleasingLockedResourceShouldFail(t *testing.T, u *unstructured.Unstructured, uniqueId1, uniqueId2 string, leaser *leaser.Leaser) {
	obtainAssertSuccess(t, u, uniqueId1, defaultLeaseDuration, defaultLeaseDuration, leaser)
	releaseAssertError(t, u, uniqueId2, leaser)
	releaseAssertSuccess(t, u, uniqueId1, leaser)
}

func testReleasingExpiredResourceShouldFail(t *testing.T, u *unstructured.Unstructured, uniqueID string, leaser *leaser.Leaser) {
	shortLeaseDuration := 1 * time.Second
	obtainAssertSuccess(t, u, uniqueID, shortLeaseDuration, shortLeaseDuration, leaser)
	time.Sleep(shortLeaseDuration + 1*time.Second)
	releaseAssertError(t, u, uniqueID, leaser)
}

func testObtainingExpiredLeaseShouldSucceed(t *testing.T, u *unstructured.Unstructured, uniqueId1, uniqueId2 string, leaser *leaser.Leaser) {
	shortLeaseDuration := 10 * time.Second
	obtainAssertSuccess(t, u, uniqueId1, shortLeaseDuration, shortLeaseDuration, leaser)
	obtainAssertError(t, u, uniqueId2, defaultLeaseDuration, defaultLeaseDuration, leaser)
	time.Sleep(shortLeaseDuration)
	obtainAssertSuccess(t, u, uniqueId2, defaultLeaseDuration, defaultLeaseDuration, leaser)
	releaseAssertError(t, u, uniqueId1, leaser)
	releaseAssertSuccess(t, u, uniqueId2, leaser)
}

func testRenewLease(t *testing.T, u *unstructured.Unstructured, uniqueID string, leaser *leaser.Leaser) {
	shortMinRemaining := 2 * time.Second
	obtainAssertSuccess(t, u, uniqueID, defaultLeaseDuration, defaultLeaseDuration, leaser)
	_, originalExpirationTIme := getOwnerAndExpirationTime(t, u, leaser)
	obtainAssertSuccess(t, u, uniqueID, defaultLeaseDuration, shortMinRemaining, leaser)
	_, expirationTime := getOwnerAndExpirationTime(t, u, leaser)
	if expirationTime != originalExpirationTIme {
		t.Fatalf("expected expiration times to be equal, instead '%v' and '%v'", expirationTime, originalExpirationTIme)
	}
	time.Sleep(shortMinRemaining + 1*time.Second)
	obtainAssertSuccess(t, u, uniqueID, defaultLeaseDuration, defaultLeaseDuration-shortMinRemaining, leaser)
	_, expirationTime = getOwnerAndExpirationTime(t, u, leaser)
	if expirationTime == originalExpirationTIme {
		t.Fatalf("expected expiration times to NOT be equal")
	}
	if originalExpirationTIme.After(expirationTime) {
		t.Fatalf("expected original expiration time '%v' to be after updated expiration time '%v'", originalExpirationTIme, expirationTime)
	}
	releaseAssertSuccess(t, u, uniqueID, leaser)
}

func getOwnerAndExpirationTime(t *testing.T, u *unstructured.Unstructured, leaser *leaser.Leaser) (string, time.Time) {
	ownerID, expirationTime, err := leaser.GetOwnerAndExpirationTime(context.Background(), u)
	if err != nil {
		t.Fatalf("error getting owner and expiration time: %v", err)
	}
	return ownerID, expirationTime
}

func obtainAssertSuccess(t *testing.T, u *unstructured.Unstructured, uniqueID string, duration time.Duration, minRemaining time.Duration, leaser *leaser.Leaser) {
	t.Helper()
	err := leaser.Obtain(context.Background(), u, uniqueID, duration, minRemaining)
	if err != nil {
		t.Fatalf("error obtaining lease: %v", err)
	}
}

func obtainAssertError(t *testing.T, u *unstructured.Unstructured, uniqueID string, duration time.Duration, minRemaining time.Duration, leaser *leaser.Leaser) {
	t.Helper()
	err := leaser.Obtain(context.Background(), u, uniqueID, duration, minRemaining)
	if err == nil {
		t.Fatalf("expected error when obtaining, instead got 'nil'")
	}
}

func releaseAssertSuccess(t *testing.T, u *unstructured.Unstructured, uniqueID string, leaser *leaser.Leaser) {
	t.Helper()
	err := leaser.Release(context.Background(), u, uniqueID)
	if err != nil {
		t.Fatalf("error obtaining lease: %v", err)
	}
}

func releaseAssertError(t *testing.T, u *unstructured.Unstructured, uniqueID string, leaser *leaser.Leaser) {
	t.Helper()
	err := leaser.Release(context.Background(), u, uniqueID)
	if err == nil {
		t.Fatalf("expected error when releasing, instead got 'nil'")
	}
}

func TestMain(m *testing.M) {
	testmain.ForIntegrationTests(m, &mgr)
}
