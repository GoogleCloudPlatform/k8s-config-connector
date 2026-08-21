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

package kccmanager

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
)

func TestSyncerIntegration_OnStartedLeading_CacheSafety(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mclModPath, err := getModulePath("github.com/gke-labs/multicluster-leader-election")
	if err != nil {
		t.Fatalf("error getting module path for multicluster-leader-election: %v", err)
	}

	kubeEtlModPath, err := getModulePath("github.com/gke-labs/kube-etl/syncer")
	if err != nil {
		t.Fatalf("error getting module path for kube-etl: %v", err)
	}

	// Setup envtest with KRMSyncer CRD
	testEnv := &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join(repo.GetRootOrTestFatal(t), "operator/config/crd/bases"),
			filepath.Join(mclModPath, "config/crd/bases"),
			filepath.Join(kubeEtlModPath, "config/crd"),
		},
		ErrorIfCRDPathMissing: true,
	}

	cfg, err := testEnv.Start()
	if err != nil {
		t.Fatalf("error starting envtest: %v", err)
	}
	defer func() {
		if err := testEnv.Stop(); err != nil {
			t.Errorf("error stopping envtest: %v", err)
		}
	}()

	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = mclv1alpha1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)
	_ = syncerv1alpha1.AddToScheme(scheme)

	kubeClient, err := client.New(cfg, client.Options{Scheme: scheme})
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	nsName := "test-syncer-ns"
	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: nsName}}
	if err := kubeClient.Create(ctx, ns); err != nil {
		t.Fatalf("error creating test namespace: %v", err)
	}

	// Pre-create the Syncer object so it exists when EnsureSuspended is called
	syncer := &unstructured.Unstructured{}
	syncer.SetGroupVersionKind(KRMSyncerGVK)
	syncer.SetName(testLeaseName)
	syncer.SetNamespace(nsName)
	_ = unstructured.SetNestedField(syncer.Object, false, "spec", "suspend")
	_ = unstructured.SetNestedField(syncer.Object, "dummy-secret", "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name")
	_ = unstructured.SetNestedSlice(syncer.Object, []interface{}{
		map[string]interface{}{
			"group":   "example.com",
			"version": "v1",
			"kind":    "Foo",
		},
	}, "spec", "rules")

	if err := kubeClient.Create(ctx, syncer); err != nil {
		t.Fatalf("error creating syncer: %v", err)
	}

	identity := "test-leader"
	mclConfig := &operatorv1beta1.MultiClusterLeaseSpec{
		LeaseName:                testLeaseName,
		Namespace:                nsName,
		ClusterCandidateIdentity: identity,
	}

	kccCfg := Config{
		ManagerOptions: manager.Options{
			Metrics:                metricsserver.Options{BindAddress: "0"},
			HealthProbeBindAddress: "0",
			LeaderElection:         false,
			LeaderElectionID:       identity,
		},
		MultiClusterLease: true,
		SyncerIntegration: true, // IMPORTANT: Turn on syncer integration
		testConfig: testConfig{
			skipControllerRegistration:   true,
			multiClusterLeaseConfig:      mclConfig,
			suppressExitOnLeadershipLoss: true,
		},
	}

	mgr, err := New(context.Background(), cfg, kccCfg)
	if err != nil {
		t.Fatalf("error creating new manager: %v", err)
	}

	signal := &startSignal{started: make(chan struct{})}
	if err := mgr.Add(signal); err != nil {
		t.Fatalf("error adding start signal: %v", err)
	}

	leaseController := &mockLeaseController{client: kubeClient}

	// Start manager in goroutine
	go func() {
		if err := mgr.Start(ctx); err != nil {
			t.Errorf("mgr failed to start: %v", err)
		}
	}()

	// Wait for lease to be created
	leaseNN := types.NamespacedName{Name: testLeaseName, Namespace: nsName}
	lease := &mclv1alpha1.MultiClusterLease{}
	if err := waitForObject(ctx, kubeClient, leaseNN, lease, 15*time.Second); err != nil {
		t.Fatalf("timed out waiting for lease object to be created: %v", err)
	}

	if err := waitForCandidate(ctx, kubeClient, leaseNN); err != nil {
		t.Fatalf("timed out waiting for lease candidate update")
	}

	// Elect the manager as leader. This will trigger the OnStartedLeading callback.
	if err := leaseController.elect(ctx, identity, nsName); err != nil {
		t.Fatalf("error electing candidate: %v", err)
	}

	// Wait for the manager to fully start
	select {
	case <-signal.started:
		t.Log("mgr started successfully without cache deadlock")
	case <-time.After(15 * time.Second):
		t.Fatal("timed out waiting for mgr to start. If using cached client in EnsureSuspended, this is a cache deadlock!")
	}

	// Verify the syncer was actually suspended by the OnStartedLeading hook
	updatedSyncer := &unstructured.Unstructured{}
	updatedSyncer.SetGroupVersionKind(KRMSyncerGVK)
	if err := kubeClient.Get(ctx, leaseNN, updatedSyncer); err != nil {
		t.Fatalf("failed to get updated syncer: %v", err)
	}
	suspend, _, _ := unstructured.NestedBool(updatedSyncer.Object, "spec", "suspend")
	if !suspend {
		t.Errorf("expected syncer to be suspended after manager became leader")
	}
}

func TestSyncerIntegration_NamespacedMode(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mclModPath, err := getModulePath("github.com/gke-labs/multicluster-leader-election")
	if err != nil {
		t.Fatalf("error getting module path for multicluster-leader-election: %v", err)
	}

	kubeEtlModPath, err := getModulePath("github.com/gke-labs/kube-etl/syncer")
	if err != nil {
		t.Fatalf("error getting module path for kube-etl: %v", err)
	}

	testEnv := &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join(repo.GetRootOrTestFatal(t), "operator/config/crd/bases"),
			filepath.Join(mclModPath, "config/crd/bases"),
			filepath.Join(kubeEtlModPath, "config/crd"),
		},
		ErrorIfCRDPathMissing: true,
	}

	cfg, err := testEnv.Start()
	if err != nil {
		t.Fatalf("error starting envtest: %v", err)
	}
	defer func() {
		if err := testEnv.Stop(); err != nil {
			t.Errorf("error stopping envtest: %v", err)
		}
	}()

	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = mclv1alpha1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)
	_ = syncerv1alpha1.AddToScheme(scheme)

	kubeClient, err := client.New(cfg, client.Options{Scheme: scheme})
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	nsName := "cnrm-system" // where the lease lives
	watchedNs := "tenant-a" // the namespace the manager watches

	for _, n := range []string{nsName, watchedNs} {
		ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: n}}
		if err := kubeClient.Create(ctx, ns); err != nil {
			t.Fatalf("error creating test namespace %s: %v", n, err)
		}
	}

	identity := "test-leader"
	mclConfig := &operatorv1beta1.MultiClusterLeaseSpec{
		LeaseName:                testLeaseName,
		Namespace:                nsName,
		ClusterCandidateIdentity: identity,
	}

	kccCfg := Config{
		ManagerOptions: manager.Options{
			Metrics:                metricsserver.Options{BindAddress: "0"},
			HealthProbeBindAddress: "0",
			LeaderElection:         false,
			LeaderElectionID:       identity,
			Cache: cache.Options{
				DefaultNamespaces: map[string]cache.Config{
					watchedNs: {}, // Configure manager for Namespaced Mode watching 'tenant-a'
				},
			},
		},
		MultiClusterLease: true,
		SyncerIntegration: true,
		testConfig: testConfig{
			skipControllerRegistration:   true,
			multiClusterLeaseConfig:      mclConfig,
			suppressExitOnLeadershipLoss: true,
		},
	}

	mgr, err := New(context.Background(), cfg, kccCfg)
	if err != nil {
		t.Fatalf("error creating new manager: %v", err)
	}

	signal := &startSignal{started: make(chan struct{})}
	if err := mgr.Add(signal); err != nil {
		t.Fatalf("error adding start signal: %v", err)
	}

	leaseController := &mockLeaseController{client: kubeClient}

	// Start manager in goroutine
	go func() {
		if err := mgr.Start(ctx); err != nil {
			t.Errorf("mgr failed to start: %v", err)
		}
	}()

	// Wait for lease to be created
	leaseNN := types.NamespacedName{Name: testLeaseName, Namespace: nsName}
	lease := &mclv1alpha1.MultiClusterLease{}
	if err := waitForObject(ctx, kubeClient, leaseNN, lease, 15*time.Second); err != nil {
		t.Fatalf("timed out waiting for lease object to be created: %v", err)
	}

	if err := waitForCandidate(ctx, kubeClient, leaseNN); err != nil {
		t.Fatalf("timed out waiting for lease candidate update")
	}

	// Elect the manager as leader.
	if err := leaseController.elect(ctx, identity, nsName); err != nil {
		t.Fatalf("error electing candidate: %v", err)
	}

	// Wait for the manager to fully start
	select {
	case <-signal.started:
		t.Log("mgr started successfully")
	case <-time.After(15 * time.Second):
		t.Fatal("timed out waiting for mgr to start")
	}

	// Verify the correct namespaced syncer was created
	syncerNN := types.NamespacedName{
		Name:      fmt.Sprintf("%s-%s", testLeaseName, watchedNs), // Expected name format: <lease>-<namespace>
		Namespace: nsName,                                         // Always lives in the system namespace with the lease
	}

	createdSyncer := &unstructured.Unstructured{}
	createdSyncer.SetGroupVersionKind(KRMSyncerGVK)
	if err := waitForObject(ctx, kubeClient, syncerNN, createdSyncer, 5*time.Second); err != nil {
		t.Fatalf("failed to find expected namespaced syncer %s: %v", syncerNN, err)
	}

	// Verify the rules contain the correct namespace
	rules, found, _ := unstructured.NestedSlice(createdSyncer.Object, "spec", "rules")
	if !found || len(rules) == 0 {
		t.Fatalf("expected rules in created syncer")
	}

	rule := rules[0].(map[string]interface{})
	namespaces, found, _ := unstructured.NestedStringSlice(rule, "namespaces")
	if !found {
		t.Fatalf("expected 'namespaces' array in syncer rule")
	}
	if len(namespaces) != 1 || namespaces[0] != watchedNs {
		t.Errorf("expected syncer rule to target namespace [%s], got %v", watchedNs, namespaces)
	}
}
