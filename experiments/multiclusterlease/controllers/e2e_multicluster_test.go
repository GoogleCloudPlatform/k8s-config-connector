// Copyright 2025 Google LLC
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

//go:build e2e_multi

package controllers

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/leaderelection"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/api/v1alpha1"
	multiclusterleaselock "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/pkg/client"
)

const (
	cluster1Context = "kind-multiclusterlease-e2e-1"
	cluster2Context = "kind-multiclusterlease-e2e-2"
)

const (
	leaseName     = "e2e-test-lease"
	testNamespace = "e2e-test-ns"
	leaderTimeout = 20 * time.Second
	leaseDuration = 15 * time.Second
	renewDeadline = 10 * time.Second
	retryPeriod   = 2 * time.Second
)

func TestE2E_MultiCluster_LeaderElection(t *testing.T) {
	testCtx, testCancel := context.WithTimeout(context.Background(), 120*time.Second) // Increased timeout for multi-stage test
	defer testCancel()

	// Register our custom scheme.
	err := v1alpha1.AddToScheme(scheme.Scheme)
	require.NoError(t, err)

	// --- Setup Cluster 1 ---
	client1, err := getKubeClient(cluster1Context)
	require.NoError(t, err)
	nsName := fmt.Sprintf("%s-%s", testNamespace, uuid.New().String()[:8])
	require.NoError(t, createNamespace(testCtx, client1, nsName))
	defer func() {
		require.NoError(t, deleteNamespace(context.Background(), client1, nsName))
	}()

	// --- Setup Cluster 2 ---
	client2, err := getKubeClient(cluster2Context)
	require.NoError(t, err)
	require.NoError(t, createNamespace(testCtx, client2, nsName))
	defer func() {
		require.NoError(t, deleteNamespace(context.Background(), client2, nsName))
	}()

	// --- Stage 1: Initial Election ---
	electorCtx1, electorCancel1 := context.WithCancel(testCtx)
	electorCtx2, electorCancel2 := context.WithCancel(testCtx)
	var wg sync.WaitGroup
	wg.Add(2)

	winnerCh := make(chan string, 2)

	id1 := fmt.Sprintf("elector-1-%s", uuid.New().String()[:8])
	go runElector(electorCtx1, t, &wg, winnerCh, client1, id1, nsName)

	id2 := fmt.Sprintf("elector-2-%s", uuid.New().String()[:8])
	go runElector(electorCtx2, t, &wg, winnerCh, client2, id2, nsName)

	// --- Verification for Stage 1 ---
	var winner1 string
	select {
	case winner1 = <-winnerCh:
		t.Logf("[Stage 1] SUCCESS: Elector [%s] won the initial election", winner1)
	case <-time.After(leaderTimeout):
		t.Fatalf("[Stage 1] failed to elect a leader within test timeout")
	}
	require.Contains(t, []string{id1, id2}, winner1, "winner should be one of the two electors")
	time.Sleep(5 * time.Second) // Give time for the loser's status to be updated.
	require.Len(t, winnerCh, 0, "only one elector should have won the initial election")

	verifyCtx, verifyCancel1 := context.WithTimeout(context.Background(), 15*time.Second)
	defer verifyCancel1()
	t.Logf("[Stage 1] Verifying lease status in both clusters...")
	verifyLeaseStatus(verifyCtx, t, client1, nsName, winner1)
	verifyLeaseStatus(verifyCtx, t, client2, nsName, winner1)
	t.Logf("[Stage 1] Verification successful")

	// --- Stage 2: Failover ---
	t.Logf("[Stage 2] Stopping the winning elector [%s] to trigger failover...", winner1)
	if winner1 == id1 {
		electorCancel1()
	} else {
		electorCancel2()
	}

	// --- Verification for Stage 2 ---
	var winner2 string
	select {
	case winner2 = <-winnerCh:
		t.Logf("[Stage 2] SUCCESS: Elector [%s] won the failover election", winner2)
	case <-testCtx.Done():
		t.Fatalf("[Stage 2] failed to elect a new leader within test timeout")
	}

	// The second winner must be the elector that did not win the first time.
	var expectedWinner2 string
	if winner1 == id1 {
		expectedWinner2 = id2
	} else {
		expectedWinner2 = id1
	}
	require.Equal(t, expectedWinner2, winner2, "the second winner should be the other elector")

	// --- Final Verification and Cleanup ---
	electorCancel1()
	electorCancel2()
	wg.Wait()

	verifyCtx2, verifyCancel2 := context.WithTimeout(context.Background(), 15*time.Second)
	defer verifyCancel2()
	t.Logf("[Stage 2] Verifying final lease status in both clusters...")
	verifyLeaseStatus(verifyCtx2, t, client1, nsName, winner2)
	verifyLeaseStatus(verifyCtx2, t, client2, nsName, winner2)
	t.Logf("[Stage 2] Final verification successful")
}

func runElector(ctx context.Context, t *testing.T, wg *sync.WaitGroup, winnerCh chan<- string, kubeClient client.Client, identity, nsName string) {
	defer wg.Done()

	lock := multiclusterleaselock.New(kubeClient, leaseName, nsName, identity, retryPeriod)

	elector, err := leaderelection.NewLeaderElector(leaderelection.LeaderElectionConfig{
		Lock:          lock,
		LeaseDuration: leaseDuration,
		RenewDeadline: renewDeadline,
		RetryPeriod:   retryPeriod,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				t.Logf("elector [%s] started leading", identity)
				winnerCh <- identity
			},
			OnStoppedLeading: func() {
				t.Logf("elector [%s] stopped leading", identity)
			},
		},
		Name: "e2e-test-elector-" + identity,
	})
	require.NoError(t, err)

	elector.Run(ctx)
	t.Logf("elector [%s] has shut down", identity)
}

func verifyLeaseStatus(ctx context.Context, t *testing.T, kubeClient client.Client, nsName, expectedWinner string) {
	var lease v1alpha1.MultiClusterLease
	key := client.ObjectKey{Namespace: nsName, Name: leaseName}

	// Use require.Eventually to poll for the status, making the test more robust.
	require.Eventually(t, func() bool {
		err := kubeClient.Get(ctx, key, &lease)
		if err != nil {
			return false
		}
		/*
			t.Logf("[debug] want winner %s, lease status: %+v", expectedWinner, lease.Status)
			if lease.Status.GlobalHolderIdentity != nil {
				t.Logf("[debug] observed winner %s", *lease.Status.GlobalHolderIdentity)
			}
		*/
		return lease.Status.GlobalHolderIdentity != nil && *lease.Status.GlobalHolderIdentity == expectedWinner
	}, 10*time.Second, 1*time.Second, "timed out waiting for lease status to be updated to winner '%s'", expectedWinner)
}

func getKubeClient(contextName string) (client.Client, error) {
	cfg, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{CurrentContext: contextName},
	).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get client config for context %s: %w", contextName, err)
	}
	return client.New(cfg, client.Options{Scheme: scheme.Scheme})
}

func createNamespace(ctx context.Context, kubeClient client.Client, name string) error {
	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: name}}
	return kubeClient.Create(ctx, ns)
}

func deleteNamespace(ctx context.Context, kubeClient client.Client, name string) error {
	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: name}}
	return kubeClient.Delete(ctx, ns)
}
