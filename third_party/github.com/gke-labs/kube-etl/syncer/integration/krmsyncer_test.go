// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package integration

import (
	"bytes"
	"context"
	"fmt"
	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"github.com/gke-labs/kube-etl/syncer/controllers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"os"
	"os/exec"
	"path/filepath"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"slices"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	if os.Getenv("KUBEBUILDER_ASSETS") == "" {
		cmd := exec.Command("go", "run", "sigs.k8s.io/controller-runtime/tools/setup-envtest@latest", "use", "-p", "path")
		out, err := cmd.Output()
		if err == nil && len(out) > 0 {
			os.Setenv("KUBEBUILDER_ASSETS", string(bytes.TrimSpace(out)))
		}
	}
	os.Exit(m.Run())
}

func TestKRMSyncerIntegration(t *testing.T) {
	// 1. Setup Environments
	testScheme := scheme.Scheme
	require.NoError(t, krmv1alpha1.AddToScheme(testScheme))

	// Cluster A
	t.Log("Starting Cluster A...")
	cfgA := &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "config", "crd"), filepath.Join("..", "integration", "testcrd")},
		ErrorIfCRDPathMissing: true,
	}
	configA, err := cfgA.Start()
	require.NoError(t, err)
	defer cfgA.Stop()

	k8sClientA, err := client.New(configA, client.Options{Scheme: testScheme})
	require.NoError(t, err)

	// Cluster B
	t.Log("Starting Cluster B...")
	cfgB := &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "integration", "testcrd")},
		ErrorIfCRDPathMissing: true,
	}
	configB, err := cfgB.Start()
	require.NoError(t, err)
	defer cfgB.Stop()

	k8sClientB, err := client.New(configB, client.Options{Scheme: testScheme})
	require.NoError(t, err)

	// 2. Start Manager for Cluster A
	t.Log("Starting Manager on Cluster A...")
	mgr, err := ctrl.NewManager(configA, ctrl.Options{
		Scheme:  testScheme,
		Metrics: metricsserver.Options{BindAddress: "0"},
	})
	require.NoError(t, err)

	reconciler := &controllers.KRMSyncerReconciler{
		Client:  mgr.GetClient(),
		Scheme:  mgr.GetScheme(),
		Manager: mgr,
		// Resolve the GKE cluster referenced by the test cases to Cluster B.
		RemoteConfigProvider: staticConfigProvider{cfg: configB},
	}
	require.NoError(t, reconciler.SetupWithManager(mgr))

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	go func() {
		err := mgr.Start(ctx)
		if err != nil {
			t.Errorf("Manager failed: %v", err)
		}
	}()

	// 3. Run Test Cases
	casesDir := "../integration/cases"
	dirs, err := os.ReadDir(casesDir)
	require.NoError(t, err)

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		caseName := d.Name()
		t.Run(caseName, func(t *testing.T) {
			runTestCase(t, ctx, k8sClientA, k8sClientB, filepath.Join(casesDir, caseName))
		})
	}

	// This manager uses the default resync period (about 10h), so these tests
	// rely on the KRMSyncer event filter alone.
	t.Run("spec-change-without-resync", func(t *testing.T) {
		testSpecChangeWithoutResync(t, ctx, k8sClientA)
	})
	t.Run("status-writes-do-not-rerun-setup", func(t *testing.T) {
		testStatusWritesDoNotRerunSetup(t, ctx, k8sClientA)
	})
}

// testSpecChangeWithoutResync verifies that a spec change (suspend) is set up
// and reported right away, without waiting for a periodic resync.
func testSpecChangeWithoutResync(t *testing.T, ctx context.Context, c client.Client) {
	syncer := newTestSyncer("spec-change-syncer", "spec")
	require.NoError(t, c.Create(ctx, syncer))
	defer func() { _ = c.Delete(ctx, syncer) }()
	waitForReady(t, ctx, c, syncer, "True", krmv1alpha1.ReasonSyncing, 1)

	setSuspend := func(suspend bool) {
		t.Helper()
		require.NoError(t, c.Get(ctx, client.ObjectKeyFromObject(syncer), syncer))
		syncer.Spec.Suspend = suspend
		require.NoError(t, c.Update(ctx, syncer))
	}
	setSuspend(true)
	waitForReady(t, ctx, c, syncer, "True", krmv1alpha1.ReasonSuspended, 2)
	setSuspend(false)
	waitForReady(t, ctx, c, syncer, "True", krmv1alpha1.ReasonSyncing, 3)
}

// testStatusWritesDoNotRerunSetup verifies that status writes (here, as
// objects start failing) do not trigger KRMSyncer reconciles.
func testStatusWritesDoNotRerunSetup(t *testing.T, ctx context.Context, c client.Client) {
	// Status mode, and no destination objects: every TestCRD fails to sync.
	syncer := newTestSyncer("status-writes-syncer", "status")
	require.NoError(t, c.Create(ctx, syncer))
	defer func() { _ = c.Delete(ctx, syncer) }()
	waitForReady(t, ctx, c, syncer, "True", krmv1alpha1.ReasonSyncing, 1)

	// Wait for KRMSyncer reconciles (e.g. retries of earlier cases) to settle.
	reconciles := krmSyncerReconciles(t)
	settled := time.Now()
	for time.Since(settled) < 5*time.Second {
		time.Sleep(500 * time.Millisecond)
		if n := krmSyncerReconciles(t); n != reconciles {
			reconciles, settled = n, time.Now()
		}
	}

	// Each new failing object changes status.failedObjects: a status write.
	var objs []*unstructured.Unstructured
	defer func() {
		for _, obj := range objs {
			_ = c.Delete(ctx, obj)
		}
	}()
	for i := 1; i <= 3; i++ {
		obj := &unstructured.Unstructured{Object: map[string]any{
			"apiVersion": "e2e.gkelabs.io/v1alpha1",
			"kind":       "TestCRD",
			"metadata":   map[string]any{"name": fmt.Sprintf("status-writes-%d", i), "namespace": "default"},
			"spec":       map[string]any{"foo": "bar"},
		}}
		require.NoError(t, c.Create(ctx, obj))
		objs = append(objs, obj)
		waitForFailedObjects(t, ctx, c, syncer, int32(i))
	}
	assert.Equal(t, reconciles, krmSyncerReconciles(t), "status writes must not trigger KRMSyncer reconciles")
}

// newTestSyncer returns a Push mode KRMSyncer of TestCRDs to cluster B.
func newTestSyncer(name, syncField string) *krmv1alpha1.KRMSyncer {
	return &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: "default"},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode: krmv1alpha1.ModePush,
			Remote: &krmv1alpha1.RemoteConfig{GKECluster: &krmv1alpha1.GKECluster{
				Project: "test-project", Location: "us-central1", Name: "cluster-b",
			}},
			Rules: []krmv1alpha1.ResourceRule{{
				Group: "e2e.gkelabs.io", Version: "v1alpha1", Kind: "TestCRD", SyncFields: []string{syncField},
			}},
		},
	}
}

// waitForReady waits (at most 10s, much less than a resync) until syncer is
// Ready with status and reason, for generation.
func waitForReady(t *testing.T, ctx context.Context, c client.Client, syncer *krmv1alpha1.KRMSyncer, status metav1.ConditionStatus, reason string, generation int64) {
	t.Helper()
	var last krmv1alpha1.KRMSyncerStatus
	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 10*time.Second, true, func(ctx context.Context) (bool, error) {
		got := &krmv1alpha1.KRMSyncer{}
		if err := c.Get(ctx, client.ObjectKeyFromObject(syncer), got); err != nil {
			return false, nil
		}
		last = got.Status
		ready := meta.FindStatusCondition(got.Status.Conditions, krmv1alpha1.ConditionReady)
		return got.Generation == generation && got.Status.ObservedGeneration == generation &&
			ready != nil && ready.Status == status && ready.Reason == reason, nil
	})
	require.NoError(t, err, "want Ready=%s/%s at generation %d, last status: %+v", status, reason, generation, last)
}

// waitForFailedObjects waits until syncer reports want failed objects.
func waitForFailedObjects(t *testing.T, ctx context.Context, c client.Client, syncer *krmv1alpha1.KRMSyncer, want int32) {
	t.Helper()
	var last int32
	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		got := &krmv1alpha1.KRMSyncer{}
		if err := c.Get(ctx, client.ObjectKeyFromObject(syncer), got); err != nil {
			return false, nil
		}
		last = got.Status.FailedObjects
		return last == want, nil
	})
	require.NoError(t, err, "want failedObjects=%d, last: %d", want, last)
}

// krmSyncerReconciles returns the number of KRMSyncer reconciles so far
// (controller_runtime_reconcile_total for the krmsyncer controller).
func krmSyncerReconciles(t *testing.T) float64 {
	t.Helper()
	families, err := metrics.Registry.Gather()
	require.NoError(t, err)
	var total float64
	for _, family := range families {
		if family.GetName() != "controller_runtime_reconcile_total" {
			continue
		}
		for _, m := range family.GetMetric() {
			for _, label := range m.GetLabel() {
				if label.GetName() == "controller" && label.GetValue() == "krmsyncer" {
					total += m.GetCounter().GetValue()
				}
			}
		}
	}
	return total
}

func runTestCase(t *testing.T, ctx context.Context, clientA, clientB client.Client, caseDir string) {
	t.Logf("Running test case: %s", filepath.Base(caseDir))

	// Load Syncer
	syncerBytes, err := os.ReadFile(filepath.Join(caseDir, "syncer.yaml"))
	require.NoError(t, err)
	syncer := &krmv1alpha1.KRMSyncer{}
	require.NoError(t, yaml.Unmarshal(syncerBytes, syncer))

	// Clean up Syncer at end of test
	defer func() {
		_ = clientA.Delete(ctx, syncer)
	}()

	// Start the Syncer
	require.NoError(t, clientA.Create(ctx, syncer))

	// Determine Source and Destination clients based on mode
	sourceClient := clientA
	destClient := clientB
	if syncer.Spec.Mode == krmv1alpha1.ModePull {
		sourceClient = clientB
		destClient = clientA
	}

	// Read all Resources from object.yaml
	resources := readObjects(t, "../integration/testdata/object.yaml")

	// Read the expected Destination Resources from expected.yaml
	expectedResources := make(map[string]*unstructured.Unstructured)
	expectedPath := filepath.Join(caseDir, "expected.yaml")
	if _, err := os.Stat(expectedPath); err == nil {
		for _, exp := range readObjects(t, expectedPath) {
			// Use Kind/Name as key to be safe, although names are unique here
			key := fmt.Sprintf("%s/%s", exp.GetKind(), exp.GetName())
			expectedResources[key] = exp
		}
	}

	// Clean up Resources at end of test
	defer func() {
		for _, r := range resources {
			_ = clientA.Delete(ctx, r)
			_ = clientB.Delete(ctx, r) // Cleanup on both just in case
		}
	}()

	// Status-mode syncers (syncFields without "spec") only update objects that
	// already exist on the destination. For those, pre-create (without status)
	// the expected resources in Destination. If the case provides a
	// destination.yaml, those objects are used (e.g. with spec values that
	// differ from the source, to verify that fields outside syncFields are not
	// synced); otherwise the source objects from object.yaml are copied.
	destinationPath := filepath.Join(caseDir, "destination.yaml")
	if _, err := os.Stat(destinationPath); err == nil {
		for _, destObj := range readObjects(t, destinationPath) {
			t.Logf("Creating Resource %s/%s in Destination Cluster...", destObj.GetKind(), destObj.GetName())
			unstructured.RemoveNestedField(destObj.Object, "status")
			require.NoError(t, destClient.Create(ctx, destObj))
		}
	} else if isStatusMode(syncer) {
		for _, resource := range resources {
			key := fmt.Sprintf("%s/%s", resource.GetKind(), resource.GetName())
			if _, ok := expectedResources[key]; !ok {
				continue
			}
			t.Logf("Creating Resource %s in Destination Cluster...", key)
			destObj := resource.DeepCopy()
			unstructured.RemoveNestedField(destObj.Object, "status")
			require.NoError(t, destClient.Create(ctx, destObj))
		}
	}

	// Apply All Resources in Source
	for _, resource := range resources {
		t.Logf("Creating Resource %s/%s in Source Cluster...", resource.GetKind(), resource.GetName())
		// Capture initial status
		initialStatus, hasStatus, err := unstructured.NestedFieldCopy(resource.Object, "status")
		require.NoError(t, err)

		// Create in Source
		resCopy := resource.DeepCopy()
		require.NoError(t, sourceClient.Create(ctx, resCopy))

		// Update Status in Source
		if hasStatus {
			statusObj := resCopy.DeepCopy()
			err := unstructured.SetNestedField(statusObj.Object, initialStatus, "status")
			require.NoError(t, err)
			require.NoError(t, sourceClient.Status().Update(ctx, statusObj))
		}
	}

	// Sleep to allow resources to sync
	time.Sleep(2 * time.Second)

	// Verify Output
	for _, resource := range resources {
		key := fmt.Sprintf("%s/%s", resource.GetKind(), resource.GetName())
		expected, shouldExist := expectedResources[key]

		actual := &unstructured.Unstructured{}
		actual.SetGroupVersionKind(resource.GroupVersionKind())
		err := destClient.Get(ctx, client.ObjectKeyFromObject(resource), actual)

		if shouldExist {
			t.Logf("Verifying expected resource %s in Destination Cluster...", key)
			require.NoError(t, err, "Resource %s should exist in Destination Cluster", key)

			// Compare Spec
			expectedSpec, _, _ := unstructured.NestedMap(expected.Object, "spec")
			actualSpec, _, _ := unstructured.NestedMap(actual.Object, "spec")
			assert.Equal(t, expectedSpec, actualSpec, "Spec mismatch for %s", key)

			// Compare Status
			expectedStatus, _, _ := unstructured.NestedMap(expected.Object, "status")
			actualStatus, _, _ := unstructured.NestedMap(actual.Object, "status")
			assert.Equal(t, expectedStatus, actualStatus, "Status mismatch for %s", key)
		} else {
			t.Logf("Verifying resource %s does NOT exist in Destination Cluster...", key)
			assert.True(t, errors.IsNotFound(err), "Resource %s should NOT exist in Destination Cluster", key)
		}
	}

	// Verify the KRMSyncer status, if the case provides syncer-status.yaml.
	statusPath := filepath.Join(caseDir, "syncer-status.yaml")
	if _, err := os.Stat(statusPath); err == nil {
		verifySyncerStatus(t, ctx, clientA, syncer, statusPath)
	}
}

// verifySyncerStatus waits until the status of syncer matches the expected
// status in path: failedObjects, and the type, status, reason and (if set)
// message of each expected condition. observedGeneration must match the
// KRMSyncer generation.
func verifySyncerStatus(t *testing.T, ctx context.Context, c client.Client, syncer *krmv1alpha1.KRMSyncer, path string) {
	t.Helper()
	b, err := os.ReadFile(path)
	require.NoError(t, err)
	want := krmv1alpha1.KRMSyncerStatus{}
	require.NoError(t, yaml.Unmarshal(b, &want))

	t.Logf("Verifying KRMSyncer status...")
	var mismatch string
	err = wait.PollUntilContextTimeout(ctx, 200*time.Millisecond, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		got := &krmv1alpha1.KRMSyncer{}
		if err := c.Get(ctx, client.ObjectKeyFromObject(syncer), got); err != nil {
			mismatch = err.Error()
			return false, nil
		}
		mismatch = syncerStatusMismatch(want, got)
		return mismatch == "", nil
	})
	require.NoError(t, err, "KRMSyncer status mismatch: %s", mismatch)
}

// syncerStatusMismatch describes how the status of got differs from want, or
// returns "" if it matches.
func syncerStatusMismatch(want krmv1alpha1.KRMSyncerStatus, got *krmv1alpha1.KRMSyncer) string {
	if got.Status.ObservedGeneration != got.Generation {
		return fmt.Sprintf("observedGeneration is %d, want %d", got.Status.ObservedGeneration, got.Generation)
	}
	if got.Status.FailedObjects != want.FailedObjects {
		return fmt.Sprintf("failedObjects is %d, want %d", got.Status.FailedObjects, want.FailedObjects)
	}
	for _, w := range want.Conditions {
		g := meta.FindStatusCondition(got.Status.Conditions, w.Type)
		switch {
		case g == nil:
			return fmt.Sprintf("condition %s is not set", w.Type)
		case g.Status != w.Status || g.Reason != w.Reason:
			return fmt.Sprintf("condition %s is %s/%s (%q), want %s/%s", w.Type, g.Status, g.Reason, g.Message, w.Status, w.Reason)
		case w.Message != "" && g.Message != w.Message:
			return fmt.Sprintf("condition %s message is %q, want %q", w.Type, g.Message, w.Message)
		case g.ObservedGeneration != got.Generation:
			return fmt.Sprintf("condition %s observedGeneration is %d, want %d", w.Type, g.ObservedGeneration, got.Generation)
		}
	}
	return ""
}

// isStatusMode reports whether no rule of the syncer syncs the full spec, i.e.
// the syncer only updates objects that already exist on the destination.
func isStatusMode(syncer *krmv1alpha1.KRMSyncer) bool {
	for _, rule := range syncer.Spec.Rules {
		if slices.Contains(rule.SyncFields, "spec") {
			return false
		}
	}
	return true
}

// readObjects decodes all objects from a multi-document YAML file.
func readObjects(t *testing.T, path string) []*unstructured.Unstructured {
	b, err := os.ReadFile(path)
	require.NoError(t, err)

	objs := []*unstructured.Unstructured{}
	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(b), 4096)
	for {
		obj := &unstructured.Unstructured{}
		if err := decoder.Decode(obj); err != nil {
			if err == io.EOF {
				break
			}
			require.NoError(t, err)
		}
		if obj.Object == nil {
			continue
		}
		objs = append(objs, obj)
	}
	return objs
}

// staticConfigProvider resolves every remote cluster to a fixed envtest cluster.
type staticConfigProvider struct {
	cfg *rest.Config
}

// Key uses the same connection key as the GKE provider.
func (p staticConfigProvider) Key(namespace string, remote *krmv1alpha1.RemoteConfig) (string, error) {
	return (&controllers.GKEConfigProvider{}).Key(namespace, remote)
}

func (p staticConfigProvider) RESTConfig(_ context.Context, _ string, remote *krmv1alpha1.RemoteConfig) (*rest.Config, error) {
	if remote == nil || remote.GKECluster == nil {
		return nil, fmt.Errorf("spec.remote.gkeCluster must be set")
	}
	return rest.CopyConfig(p.cfg), nil
}
