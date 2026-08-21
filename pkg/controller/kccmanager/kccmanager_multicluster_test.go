// Copyright 2025 Google LLC
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
	"os/exec"
	"path/filepath"
	stdruntime "runtime"
	"strings"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
)

const (
	testNamespace1     = "test-namespace-1"
	testNamespace2     = "test-namespace-2"
	testLeaseName      = "test-lease"
	testGlobalLockName = "test-global-lock"
)

func getModulePath(moduleName string) (string, error) {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}", moduleName)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// startSignal is a manager.Runnable that signals when it has been started.
type startSignal struct {
	started chan struct{}
}

func (s *startSignal) Start(ctx context.Context) error {
	close(s.started)
	<-ctx.Done()
	return nil
}

// mockLeaseController simulates the behavior of the decentralized election controller.
type mockLeaseController struct {
	client client.Client
}

func (c *mockLeaseController) elect(ctx context.Context, winnerIdentity string, namespaces ...string) error {
	for _, ns := range namespaces {
		lease := &mclv1alpha1.MultiClusterLease{}
		nn := types.NamespacedName{
			Namespace: ns,
			Name:      testLeaseName,
		}
		if err := c.client.Get(ctx, nn, lease); err != nil {
			return fmt.Errorf("error getting lease in %s: %w", ns, err)
		}

		// In a real election controller, it would check the Global Lock.
		// Here, we simulate that the Global Lock has been won by `winnerIdentity`.
		lease.Status.GlobalHolderIdentity = &winnerIdentity
		lease.Status.ObservedGeneration = &lease.Generation

		if err := c.client.Status().Update(ctx, lease); err != nil {
			return fmt.Errorf("error updating lease status in %s: %w", ns, err)
		}
	}
	return nil
}

func TestMultiClusterLeaderElection_TwoManagers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mclModPath, err := getModulePath("github.com/gke-labs/multicluster-leader-election")
	if err != nil {
		t.Fatalf("error getting module path for multicluster-leader-election: %v", err)
	}

	// Setup envtest
	testEnv := &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join(repo.GetRootOrTestFatal(t), "operator/config/crd/bases"),
			filepath.Join(mclModPath, "config/crd/bases"),
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

	// Create a scheme and add all the types we need
	scheme := runtime.NewScheme()
	if err := corev1.AddToScheme(scheme); err != nil {
		t.Fatalf("error adding corev1 to scheme: %v", err)
	}
	if err := mclv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatalf("error adding configv1alpha1 to scheme: %v", err)
	}
	if err := operatorv1beta1.AddToScheme(scheme); err != nil {
		t.Fatalf("error adding operatorv1beta1 to scheme: %v", err)
	}

	// Create a client
	kubeClient, err := client.New(cfg, client.Options{Scheme: scheme})
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	// Create test namespaces
	for _, nsName := range []string{testNamespace1, testNamespace2} {
		ns := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: nsName,
			},
		}
		if err := kubeClient.Create(ctx, ns); err != nil {
			t.Fatalf("error creating test namespace %s: %v", nsName, err)
		}
	}

	// Create the mock lease controller
	leaseController := &mockLeaseController{client: kubeClient}

	// Create two managers with different configs (simulating two clusters)
	mgr1, signal1 := newTestManagerWithConfig(t, cfg, "mgr1", testNamespace1)
	mgr2, signal2 := newTestManagerWithConfig(t, cfg, "mgr2", testNamespace2)

	// Start managers in goroutines
	go func() {
		if err := mgr1.Start(ctx); err != nil {
			t.Errorf("mgr1 failed to start: %v", err)
		}
	}()
	go func() {
		if err := mgr2.Start(ctx); err != nil {
			t.Errorf("mgr2 failed to start: %v", err)
		}
	}()

	// Wait for managers to create the lease object in their respective namespaces
	leaseNN1 := types.NamespacedName{Name: testLeaseName, Namespace: testNamespace1}
	lease1 := &mclv1alpha1.MultiClusterLease{}
	if err := waitForObject(ctx, kubeClient, leaseNN1, lease1, 15*time.Second); err != nil {
		t.Fatalf("timed out waiting for lease object 1 to be created: %v", err)
	}

	leaseNN2 := types.NamespacedName{Name: testLeaseName, Namespace: testNamespace2}
	lease2 := &mclv1alpha1.MultiClusterLease{}
	if err := waitForObject(ctx, kubeClient, leaseNN2, lease2, 15*time.Second); err != nil {
		t.Fatalf("timed out waiting for lease object 2 to be created: %v", err)
	}

	// --- Scenario: Wait for election, ensure only leader starts ---
	t.Run("Only leader starts", func(t *testing.T) {
		// Assert neither manager has started yet
		select {
		case <-signal1.started:
			t.Fatal("mgr1 started before election!")
		case <-signal2.started:
			t.Fatal("mgr2 started before election!")
		case <-time.After(3 * time.Second):
			// Success, still waiting
		}

		// Wait for both lease specs to be updated with identities
		// We expect mgr1 to set identity="mgr1" in lease1, and mgr2 to set identity="mgr2" in lease2
		// (Assuming the identity passed to New is what mcleclient uses as candidate ID, which we assume is true after fix)

		if err := waitForCandidate(ctx, kubeClient, leaseNN1); err != nil {
			t.Fatalf("timed out waiting for lease1 candidate update")
		}
		if err := waitForCandidate(ctx, kubeClient, leaseNN2); err != nil {
			t.Fatalf("timed out waiting for lease2 candidate update")
		}

		t.Log("Both candidates have announced themselves")

		// Elect mgr1 as the winner (Global Lock Winner)
		// The mock controller updates ALL leases (in all participating clusters/namespaces) with the winner.
		winner := "mgr1"
		if err := leaseController.elect(ctx, winner, testNamespace1, testNamespace2); err != nil {
			t.Fatalf("error electing candidate: %v", err)
		}

		// Wait to see if mgr1 starts
		select {
		case <-signal1.started:
			t.Log("mgr1 started")
		case <-time.After(10 * time.Second):
			t.Fatal("timed out waiting for mgr1 to start after election")
		}

		// Verify mgr2 does NOT start
		select {
		case <-signal2.started:
			t.Fatal("mgr2 started BUT mgr1 is the leader! SPLIT BRAIN DETECTED!")
		case <-time.After(5 * time.Second):
			t.Log("Confirmed mgr2 did not start")
		}
	})
}

func waitForCandidate(ctx context.Context, c client.Client, nn types.NamespacedName) error {
	return wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 10*time.Second, true, func(ctx context.Context) (bool, error) {
		lease := &mclv1alpha1.MultiClusterLease{}
		if err := c.Get(ctx, nn, lease); err != nil {
			return false, nil
		}
		if lease.Spec.HolderIdentity != nil && *lease.Spec.HolderIdentity != "" {
			return true, nil
		}
		return false, nil
	})
}

func waitForObject(ctx context.Context, c client.Client, nn types.NamespacedName, obj client.Object, timeout time.Duration) error {
	return wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, timeout, true, func(ctx context.Context) (bool, error) {
		if err := c.Get(ctx, nn, obj); err != nil {
			return false, nil
		}
		return true, nil
	})
}

func newTestManagerWithConfig(t *testing.T, cfg *rest.Config, identity string, namespace string) (manager.Manager, *startSignal) {
	t.Helper()

	mclConfig := &operatorv1beta1.MultiClusterLeaseSpec{
		LeaseName:                testLeaseName,
		Namespace:                namespace,
		ClusterCandidateIdentity: identity,
	}

	kccCfg := Config{
		ManagerOptions: manager.Options{
			// Scheme: scheme, // Do not provide scheme, let New() build it using addSchemes()
			Metrics: metricsserver.Options{
				BindAddress: "0",
			},
			HealthProbeBindAddress: "0",
			LeaderElection:         false,
			LeaderElectionID:       identity,
		},
		MultiClusterLease: true,
		testConfig: testConfig{
			skipControllerRegistration:   true,
			multiClusterLeaseConfig:      mclConfig,
			suppressExitOnLeadershipLoss: true,
		},
	}

	mgr, err := New(context.Background(), cfg, kccCfg)
	if err != nil {
		t.Fatalf("error creating new manager for %s: %v", identity, err)
	}

	signal := &startSignal{started: make(chan struct{})}
	if err := mgr.Add(signal); err != nil {
		t.Fatalf("error adding start signal to manager %s: %v", identity, err)
	}

	return mgr, signal
}

// MockLock implements resourcelock.Interface
type MockLock struct {
	identity string
}

func (m *MockLock) Get(ctx context.Context) (*resourcelock.LeaderElectionRecord, []byte, error) {
	// Always say we are the leader
	return &resourcelock.LeaderElectionRecord{
		HolderIdentity:       m.identity,
		LeaseDurationSeconds: 15,
		AcquireTime:          metav1.Now(),
		RenewTime:            metav1.Now(),
	}, nil, nil
}

func (m *MockLock) Create(ctx context.Context, ler resourcelock.LeaderElectionRecord) error {
	return nil
}

func (m *MockLock) Update(ctx context.Context, ler resourcelock.LeaderElectionRecord) error {
	return nil
}

func (m *MockLock) RecordEvent(string) {}

func (m *MockLock) Identity() string {
	return m.identity
}

func (m *MockLock) Describe() string {
	return "MockLock"
}

// TestSplitBrainProtection verifies that the manager crashes/halts if it believes it is the leader
// but the global MultiClusterLease status indicates a different leader.
func TestSplitBrainProtection(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mclModPath, err := getModulePath("github.com/gke-labs/multicluster-leader-election")
	if err != nil {
		t.Fatalf("error getting module path: %v", err)
	}

	testEnv := &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join(repo.GetRootOrTestFatal(t), "operator/config/crd/bases"),
			filepath.Join(mclModPath, "config/crd/bases"),
		},
		ErrorIfCRDPathMissing: true,
	}

	cfg, err := testEnv.Start()
	if err != nil {
		t.Fatalf("error starting envtest: %v", err)
	}
	defer func() {
		if err := testEnv.Stop(); err != nil {
			t.Logf("error stopping envtest: %v", err)
		}
	}()

	scheme := runtime.NewScheme()
	if err := mclv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatalf("error adding mcl to scheme: %v", err)
	}
	if err := operatorv1beta1.AddToScheme(scheme); err != nil {
		t.Fatalf("error adding operator to scheme: %v", err)
	}
	if err := corev1.AddToScheme(scheme); err != nil {
		t.Fatalf("error adding corev1 to scheme: %v", err)
	}

	k8sClient, err := client.New(cfg, client.Options{Scheme: scheme})
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	tests := []struct {
		name                 string
		globalHolderIdentity *string
		expectedErrorSnippet string
	}{
		{
			name:                 "Different global holder",
			globalHolderIdentity: func() *string { s := "imposter-manager"; return &s }(),
			expectedErrorSnippet: `inconsistent state: started leading but MultiClusterLease status.globalHolderIdentity is "imposter-manager"`,
		},
		{
			name:                 "Nil global holder",
			globalHolderIdentity: nil,
			expectedErrorSnippet: `inconsistent state: started leading but MultiClusterLease status.globalHolderIdentity is ""`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Use a unique namespace for each test case
			ns := strings.ReplaceAll(strings.ToLower(tc.name), " ", "-")
			leaseName := "test-lease"
			localIdentity := "my-manager"

			// Create namespace
			namespaceObj := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{Name: ns},
			}
			if err := k8sClient.Create(ctx, namespaceObj); err != nil {
				t.Fatalf("error creating namespace: %v", err)
			}

			leaseDuration := int32(15)
			renewTime := metav1.MicroTime{Time: time.Now()}

			// 1. Create the MultiClusterLease
			mcl := &mclv1alpha1.MultiClusterLease{
				ObjectMeta: metav1.ObjectMeta{
					Name:      leaseName,
					Namespace: ns,
				},
				Spec: mclv1alpha1.MultiClusterLeaseSpec{
					// Spec holder is not strictly relevant for this test, but filling it for realism
					HolderIdentity:       tc.globalHolderIdentity, // Mimicking that the "imposter" tried to take it
					LeaseDurationSeconds: &leaseDuration,
					RenewTime:            &renewTime,
				},
			}
			if tc.globalHolderIdentity == nil {
				// If global holder is nil, maybe spec holder was someone else or empty?
				// Let's just set spec holder to something random or leave it nil.
				// For the "Nil global holder" case, it means the status is empty.
				// Let's set spec holder to "someone-else" to show contention, or nil.
				// The test is checking Status vs Local.
				// Let's leave spec holder nil if tc is nil.
			}

			if err := k8sClient.Create(ctx, mcl); err != nil {
				t.Fatalf("failed to create MCL: %v", err)
			}

			// 2. Update Status with the test case's global holder
			mcl.Status.GlobalHolderIdentity = tc.globalHolderIdentity
			if err := k8sClient.Status().Update(ctx, mcl); err != nil {
				t.Fatalf("failed to update MCL status: %v", err)
			}

			// 3. Create the Manager
			kccCfg := Config{
				ManagerOptions: manager.Options{
					Scheme: scheme,
					Metrics: metricsserver.Options{
						BindAddress: "0",
					},
					LeaderElection:   false,
					LeaderElectionID: localIdentity,
				},
				MultiClusterLease: true,
				testConfig: testConfig{
					skipControllerRegistration: true,
					multiClusterLeaseConfig: &operatorv1beta1.MultiClusterLeaseSpec{
						LeaseName:                leaseName,
						Namespace:                ns,
						ClusterCandidateIdentity: localIdentity,
					},
					suppressExitOnLeadershipLoss: true,
				},
			}

			mgr, err := New(ctx, cfg, kccCfg)
			if err != nil {
				t.Fatalf("error creating manager: %v", err)
			}

			// 4. Hack: Swap the Lock with our MockLock
			lem, ok := mgr.(*leaderElectionManager)
			if !ok {
				t.Fatalf("manager is not *leaderElectionManager, it is %T", mgr)
			}

			lem.leConfig.Lock = &MockLock{identity: localIdentity}

			// 5. Intercept Fatal exit
			fatalC := make(chan string, 1)
			lem.onFatal = func(format string, args ...interface{}) {
				msg := fmt.Sprintf(format, args...)
				fatalC <- msg
				stdruntime.Goexit()
			}

			// 6. Start the manager in a background goroutine
			go func() {
				if err := lem.Start(ctx); err != nil {
					t.Logf("Manager.Start returned: %v", err)
				}
			}()

			// 7. Verify we received the fatal error
			select {
			case msg := <-fatalC:
				t.Logf("Caught expected fatal error: %s", msg)
				if !strings.Contains(msg, tc.expectedErrorSnippet) {
					t.Errorf("Unexpected fatal message: got %q, want substring %q", msg, tc.expectedErrorSnippet)
				}
			case <-time.After(5 * time.Second):
				t.Fatal("Timed out waiting for split-brain protection to trigger")
			}
		})
	}
}
