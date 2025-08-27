//go:build e2e

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

package controllers_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/leaderelection"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/api/v1alpha1"
	multiclusterleaselock "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/pkg/client"
)

const (
	leaseName     = "e2e-test-lease"
	testNamespace = "e2e-test-ns"
	leaderTimeout = 30 * time.Second
	leaseDuration = 15 * time.Second
	renewDeadline = 10 * time.Second
	retryPeriod   = 2 * time.Second
)

func TestE2E_LeaderElection(t *testing.T) {
	// This test assumes:
	// 1. A Kubernetes cluster is available and the KUBECONFIG is set.
	// 2. The MultiClusterLease CRD has been installed.
	// 3. The multiclusterlease-controller is running in the cluster.

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.GetConfig()
	require.NoError(t, err)

	// Register our custom scheme.
	err = v1alpha1.AddToScheme(scheme.Scheme)
	require.NoError(t, err)

	kubeClient, err := client.New(cfg, client.Options{Scheme: scheme.Scheme})
	require.NoError(t, err)

	// Create a unique namespace for this test run.
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-%s", testNamespace, uuid.New().String()[:8]),
		},
	}
	t.Logf("creating test namespace: %s", ns.Name)
	require.NoError(t, kubeClient.Create(ctx, ns))
	defer func() {
		t.Logf("deleting test namespace: %s", ns.Name)
		require.NoError(t, kubeClient.Delete(ctx, ns))
	}()

	// Channel to signal when leadership is acquired.
	leaderCh := make(chan struct{})
	identity := uuid.New().String()

	// Create our custom resourcelock.
	lock := multiclusterleaselock.New(kubeClient, leaseName, ns.Name, identity)

	// Configure the LeaderElector.
	elector, err := leaderelection.NewLeaderElector(leaderelection.LeaderElectionConfig{
		Lock:          lock,
		LeaseDuration: leaseDuration,
		RenewDeadline: renewDeadline,
		RetryPeriod:   retryPeriod,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				t.Logf("elector [%s] started leading", identity)
				close(leaderCh) // Signal that we are the leader.
			},
			OnStoppedLeading: func() {
				time.Sleep(2)
				t.Logf("elector [%s] stopped leading", identity)
			},
		},
		Name: "e2e-test-elector",
	})
	require.NoError(t, err)

	// Run the leader elector in the background.
	go elector.Run(ctx)

	// Wait for leadership to be acquired.
	t.Logf("elector [%s] waiting to acquire leadership...", identity)
	select {
	case <-leaderCh:
		t.Logf("elector [%s] successfully acquired leadership", identity)
	case <-time.After(leaderTimeout):
		t.Fatalf("elector [%s] failed to acquire leadership within %s", identity, leaderTimeout)
	}

	// Verification:
	// Once we are the leader, the status of the CR should reflect our identity.
	var finalLease v1alpha1.MultiClusterLease
	key := client.ObjectKey{Namespace: ns.Name, Name: leaseName}
	require.NoError(t, kubeClient.Get(ctx, key, &finalLease))

	require.NotNil(t, finalLease.Status.GlobalHolderIdentity, "GlobalHolderIdentity should not be nil")
	require.Equal(t, identity, *finalLease.Status.GlobalHolderIdentity, "GlobalHolderIdentity should be our identity")
	t.Logf("successfully verified that status.globalHolderIdentity is %s", identity)
}
