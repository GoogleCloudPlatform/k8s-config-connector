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

package resourcecontention_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cluster"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/lease/leaser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testreconciler "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller/reconciler"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testrunner "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/runner"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	mgr1 manager.Manager
	mgr2 manager.Manager
)

// this test creates a second 'manager', i.e. KCC installation, and then ensures that the same project
// that is being managed by the default manager in systemContext, cannot be managed by the second manager.
func TestResourceContentionIsPreventedForTwoNamespacesMappingToSameProjectInDifferentClusters(t *testing.T) {
	ctx := context.TODO()

	shouldRun := func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		// only need to test contention for a single resource since the logic will apply to all resources
		return fixture.GVK.Kind == "PubSubTopic"
	}
	testFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, systemContext testrunner.SystemContext) {
		if err := systemContext.Manager.GetClient().Create(context.TODO(), testContext.CreateUnstruct); err != nil {
			t.Fatalf("error creating resource: %v", err)
		}
		systemContext.Reconciler.Reconcile(ctx, testContext.UpdateUnstruct, testreconciler.ExpectedSuccessfulReconcileResultFor(systemContext.Reconciler, testContext.UpdateUnstruct), nil)
		assertLeaseLabelsAreNotPresent(t, systemContext.Manager, testContext.CreateUnstruct)
		projectID := testgcp.GetDefaultProjectID(t)
		testcontroller.EnsureNamespaceExistsT(t, mgr2.GetClient(), testContext.UniqueID)
		testcontroller.EnsureNamespaceHasProjectIDAnnotation(t, mgr2.GetClient(), testContext.UniqueID, projectID)
		assertNamespaceIdsAreNotEqual(t, systemContext.Manager, mgr2, testContext.UniqueID, testContext.UniqueID)
		reconciler2 := testreconciler.New(t, mgr2, systemContext.TFProvider)
		if err := mgr2.GetClient().Create(context.TODO(), testContext.UpdateUnstruct); err != nil {
			t.Fatalf("error creating resource: %v", err)
		}
		reconciler2.Reconcile(ctx, testContext.UpdateUnstruct, testreconciler.ExpectedUnsuccessfulReconcileResult, regexp.MustCompile("error obtaining lease"))
		events := testcontroller.CollectEvents(t, mgr2.GetConfig(), testContext.UpdateUnstruct.GetNamespace(), 1, 10*time.Second)
		event := events[0]
		expectedReason := k8s.ManagementConflict
		if event.Reason != expectedReason {
			t.Fatalf("event mismatch: got '%v', want '%v'", event.Reason, expectedReason)
		}
		// Since the controller was unable to obtain the lease it does not write the default finalizer onto the object.
		// Add the finalizer manually so that we can test the deletion resource contention flow.
		ensureFinalizer(t, mgr2, testContext.NamespacedName, testContext.CreateUnstruct)
		if err := mgr2.GetClient().Delete(context.TODO(), testContext.CreateUnstruct); err != nil {
			t.Fatalf("error deleting resource: %v", err)
		}
		reconciler2.Reconcile(ctx, testContext.CreateUnstruct, testreconciler.ExpectedUnsuccessfulReconcileResult, regexp.MustCompile("error obtaining lease"))
		events = testcontroller.CollectEvents(t, mgr2.GetConfig(), testContext.CreateUnstruct.GetNamespace(), 3, 10*time.Second)
		nextEvent := events[2]
		if nextEvent.Reason != expectedReason {
			t.Fatalf("event mismatch: got '%v', want '%v'", nextEvent.Reason, expectedReason)
		}
		if !(event.LastTimestamp == nextEvent.LastTimestamp || event.LastTimestamp.Before(&nextEvent.LastTimestamp)) {
			t.Fatalf("expected the previous event's last timestamp to be before or equal to the next event's last timestamp")
		}
	}
	testrunner.RunAllWithDependenciesCreatedButNotObject(ctx, t, mgr1, shouldRun, testFunc)
}

func ensureFinalizer(t *testing.T, mgr manager.Manager, namespacedName types.NamespacedName, u *unstructured.Unstructured) {
	if err := mgr.GetClient().Get(context.TODO(), namespacedName, u); err != nil {
		t.Fatalf("error getting '%v': %v", namespacedName, err)
	}
	if k8s.EnsureFinalizer(u, k8s.ControllerFinalizerName) {
		t.Fatalf("found a finalizer when none was expected")
	}
	if err := mgr.GetClient().Update(context.TODO(), u); err != nil {
		t.Fatalf("error updating '%v': %v", namespacedName, err)
	}
}

func assertLeaseLabelsAreNotPresent(t *testing.T, mgr manager.Manager, u *unstructured.Unstructured) {
	name := k8s.GetNamespacedName(u)
	if err := mgr.GetClient().Get(context.TODO(), name, u); err != nil {
		t.Fatalf("error getting resource '%v': %v", name, err)
	}
	if u.GetLabels() == nil {
		return
	}
	// the keys are hard-coded here to make sure the test doesn't have any knowledge of the implementation
	leaseKeys := []string{"cnrm-lease-expiration", "cnrm-lease-holder-id"}
	// add a check to make sure the implementation has not changed without the test being updated
	assertLeaseKeysMatch(t, leaseKeys)
	for _, k := range leaseKeys {
		if val, ok := u.GetLabels()[k]; ok {
			t.Fatalf("unexpected value of '%v' for label key '%v': the key should not be present", k, val)
		}
	}
}

func assertLeaseKeysMatch(t *testing.T, expectedKeys []string) {
	keys := leaser.GetLabelKeys()
	keysMap := make(map[string]bool)
	for _, k := range keys {
		keysMap[k] = true
	}
	for _, k := range expectedKeys {
		if _, ok := keysMap[k]; !ok {
			t.Fatalf("missing expected key '%v'", k)
		}
	}
}

func assertNamespaceIdsAreNotEqual(t *testing.T, mgr1, mgr2 manager.Manager, namespace1, namespace2 string) {
	id1 := getNamespaceID(t, mgr1, namespace1)
	id2 := getNamespaceID(t, mgr2, namespace2)
	if id1 == id2 {
		t.Fatalf("expected the ids of both managers to not match, instead, both have an id of '%v'", id1)
	}
}

func getNamespaceID(t *testing.T, mgr manager.Manager, namespace string) string {
	t.Helper()
	id, err := cluster.GetNamespaceID(context.TODO(), k8s.NamespaceIDConfigMapNN, mgr.GetClient(), namespace)
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func TestMain(m *testing.M) {
	mgrs := []*manager.Manager{&mgr1, &mgr2}
	testmain.SetupMultipleEnvironments(m, test.IntegrationTestType, nil, mgrs)
}
