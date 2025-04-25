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

package kccmanager_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dynamic"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/stateintospec"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testrunner "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/runner"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

var (
	// These managers are used to just get the rest.Config since our testmain package's methods are not easily changed
	// to return a rest.Config
	clusterModeManager    manager.Manager
	namespacedModeManager manager.Manager
)

// The scheme is not thread-safe due to its use and modification of its internal maps. Different managers should not
// share a scheme.
func TestSchemeIsUniqueAcrossManagers(t *testing.T) {
	ctx := context.TODO()

	controllersCfg := kccmanager.Config{
		ManagerOptions: manager.Options{
			// disable prometheus metrics as by default, the metrics server binds to the same port in all instances
			Metrics: server.Options{
				BindAddress: "0",
			},
		},
	}
	schemePtrMap := make(map[*runtime.Scheme]string)
	schemePtrMap[clusterModeManager.GetScheme()] = "clusterModeMgr"
	for i := 0; i < 5; i++ {
		mgr, err := kccmanager.New(ctx, clusterModeManager.GetConfig(), controllersCfg)
		if err != nil {
			t.Fatalf("error creating manager: %v", err)
		}
		mgrName := fmt.Sprintf("mgr-%v", i)
		if val, ok := schemePtrMap[mgr.GetScheme()]; ok {
			t.Fatalf("expected new manager '%v' to have a new, unique scheme, instead it is sharing the scheme with '%v'", mgrName, val)
		}
		schemePtrMap[mgr.GetScheme()] = mgrName
	}
}

func TestClusterModeManager(t *testing.T) {
	ctx := context.TODO()
	mgr, err := kccmanager.New(ctx, clusterModeManager.GetConfig(), kccmanager.Config{StateIntoSpecDefaultValue: stateintospec.StateIntoSpecDefaultValueV1Beta1})
	if err != nil {
		t.Fatalf("error creating manager: %v", err)
	}
	stop := testcontroller.StartMgr(t, mgr)
	defer stop()
	basicPubSubFixture := getBasicPubSubSchemaFixture(t)
	project := testgcp.GetDefaultProject(t)
	for i := 0; i < 2; i++ {
		tstContext := testrunner.NewTestContext(t, basicPubSubFixture, project)
		testcontroller.EnsureNamespaceExistsT(t, mgr.GetClient(), tstContext.CreateUnstruct.GetNamespace())
		if err := mgr.GetClient().Create(context.TODO(), tstContext.CreateUnstruct); err != nil {
			t.Fatalf("error creating '%v': %v", tstContext.CreateUnstruct.GetKind(), err)
		}
		waitForReconcile(t, mgr.GetClient(), tstContext.CreateUnstruct)
	}
}

// Create two resources, one in a managed namespace for which we have started controllers, and another for which we have
// not started controllers. Verify that only the first is reconciled, then start a second set of controllers and verify
// the second is reconciled.
func TestNamespacedModeManager(t *testing.T) {
	ctx := context.TODO()
	basicPubSubFixture := getBasicPubSubSchemaFixture(t)
	project := testgcp.GetDefaultProject(t)
	tstContext1 := testrunner.NewTestContext(t, basicPubSubFixture, project)
	tstContext2 := testrunner.NewTestContext(t, basicPubSubFixture, project)
	controllersCfg1 := kccmanager.Config{
		ManagerOptions: manager.Options{
			// disable prometheus metrics as by default, the metrics server binds to the same port in all instances
			Metrics: server.Options{
				BindAddress: "0",
			},
			Cache: cache.Options{
				DefaultNamespaces: map[string]cache.Config{
					tstContext1.CreateUnstruct.GetNamespace(): {},
				},
			},
		},
	}
	mgr1, err := kccmanager.New(ctx, namespacedModeManager.GetConfig(), controllersCfg1)
	if err != nil {
		t.Fatalf("error creating manager: %v", err)
	}
	testcontroller.StartMgr(t, mgr1)
	// TODO: delete the line above and uncomment the two lines below once we have a fix for the race condition out of
	// client-go, sollyross@ is working on it: https://github.com/kubernetes/kubernetes/pull/95664/files
	//stop1 := testcontroller.StartMgr(t, mgr1)
	//defer stop1()
	kubeClient := namespacedModeManager.GetClient()
	testcontroller.EnsureNamespaceExistsT(t, kubeClient, tstContext1.CreateUnstruct.GetNamespace())
	if err := kubeClient.Create(context.TODO(), tstContext1.CreateUnstruct); err != nil {
		t.Fatalf("error creating '%v': %v", tstContext1.CreateUnstruct.GetKind(), err)
	}
	testcontroller.EnsureNamespaceExistsT(t, kubeClient, tstContext2.CreateUnstruct.GetNamespace())
	if err := kubeClient.Create(context.TODO(), tstContext2.CreateUnstruct); err != nil {
		t.Fatalf("error creating '%v': %v", tstContext2.CreateUnstruct.GetKind(), err)
	}
	waitForReconcile(t, kubeClient, tstContext1.CreateUnstruct)
	// sleep 10 seconds to give this resource 'time to reconcile' even though we expect it to NOT reconcile at all as
	// no controllers are running
	time.Sleep(10 * time.Second)
	if err := kubeClient.Get(context.TODO(), tstContext2.NamespacedName, tstContext2.CreateUnstruct); err != nil {
		t.Fatalf("error getting resource: %v", err)
	}
	var expectedValue interface{} = nil
	actualValue := tstContext2.CreateUnstruct.Object["status"]
	if actualValue != expectedValue {
		t.Fatalf("unexpected value for status: got '%v', want '%v'", actualValue, expectedValue)
	}
	controllersCfg2 := kccmanager.Config{
		ManagerOptions: manager.Options{
			// disable prometheus metrics as by default, the metrics server binds to the same port in all instances
			Metrics: server.Options{
				BindAddress: "0",
			},
			Cache: cache.Options{
				DefaultNamespaces: map[string]cache.Config{
					tstContext2.CreateUnstruct.GetNamespace(): {},
				},
			},
		},
	}
	// start controllers for the second namespace and verify that the second resource does reconcile
	mgr2, err := kccmanager.New(ctx, namespacedModeManager.GetConfig(), controllersCfg2)
	if err != nil {
		t.Fatalf("error creating manager: %v", err)
	}
	testcontroller.StartMgr(t, mgr2)
	// TODO: delete the line above and uncomment the two lines below once we have a fix for the race condition out of
	// client-go, sollyross@ is working on it: https://github.com/kubernetes/kubernetes/pull/95664/files
	//stop2 := testcontroller.StartMgr(t, mgr2)
	//defer stop2()
	waitForReconcile(t, kubeClient, tstContext2.CreateUnstruct)
}

// getBasicPubSubSchemaFixture returns the basic/pubsubschema fixture.
// This is a relatively quick resource to create, that does not have any dependencies that must be created.
func getBasicPubSubSchemaFixture(t *testing.T) resourcefixture.ResourceFixture {
	lightFilter := func(name string, testType resourcefixture.TestType) bool {
		return name == "pubsubschema" && testType == resourcefixture.Basic
	}
	fixtures := resourcefixture.LoadWithFilter(t, lightFilter, nil)
	if len(fixtures) != 1 {
		t.Fatalf("unexpected number of fixtures: got '%v', want '%v'", len(fixtures), 1)
	}
	return fixtures[0]
}

func waitForReconcile(t *testing.T, kubeClient client.Client, resource *unstructured.Unstructured) {
	// return value of true means 'done'
	condFunc := func() (bool, error) {
		nn := types.NamespacedName{
			Namespace: resource.GetNamespace(),
			Name:      resource.GetName(),
		}
		u := unstructured.Unstructured{}
		u.SetGroupVersionKind(resource.GroupVersionKind())
		if err := kubeClient.Get(context.TODO(), nn, &u); err != nil {
			return false, fmt.Errorf("error getting '%v': %v", nn, err)
		}
		if u.Object["status"] == nil {
			klog.Infof("Waiting for 'status' on %v '%v'", u.GetKind(), u.GetName())
			return false, nil
		}
		objectStatus := dynamic.GetObjectStatus(t, &u)
		if objectStatus.ObservedGeneration == nil {
			klog.InfoS("resource does not yet have status.observedGeneration", "kind", u.GetKind(), "name", u.GetName())
			return false, nil
		}
		if *objectStatus.ObservedGeneration < objectStatus.Generation {
			klog.InfoS("resource status.observedGeneration is behind current generation",
				"kind", u.GetKind(), "name", u.GetName(),
				"status.observedGeneration", *objectStatus.ObservedGeneration, "generation", objectStatus.Generation)
			return false, nil
		}
		for _, c := range objectStatus.Conditions {
			if c.Type == "Ready" && c.Status == "True" {
				klog.InfoS("resource is ready", "kind", u.GetKind(), "name", u.GetName())
				return true, nil
			}
		}
		klog.InfoS("resource is not yet ready", "kind", u.GetKind(), "name", u.GetName(), "conditions", objectStatus.Conditions)
		return true, nil
	}
	if err := wait.PollImmediate(10*time.Second, 5*time.Minute, condFunc); err != nil {
		t.Fatalf("error waiting for reconcile of '%v' to complete: %v'", resource.GetKind(), err)
	}
}

func TestMain(m *testing.M) {
	managers := []*manager.Manager{
		&clusterModeManager,
		&namespacedModeManager,
	}
	testmain.SetupMultipleEnvironments(m, test.IntegrationTestType, nil, managers)
}
