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

//go:build e2e

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"cloud.google.com/go/storage"
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
	lock := multiclusterleaselock.New(kubeClient, leaseName, ns.Name, identity, retryPeriod)

	testfinished := make(chan struct{})

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
				t.Logf("elector [%s] stopped leading", identity)
				close(testfinished)
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

	// Verification 1: Initial Acquisition
	// Once we are the leader, the status of the CR should reflect our identity.
	var lease v1alpha1.MultiClusterLease
	key := client.ObjectKey{Namespace: ns.Name, Name: leaseName}
	require.NoError(t, kubeClient.Get(ctx, key, &lease))

	require.NotNil(t, lease.Status.GlobalHolderIdentity, "GlobalHolderIdentity should not be nil after acquisition")
	require.Equal(t, identity, *lease.Status.GlobalHolderIdentity, "GlobalHolderIdentity should be our identity after acquisition")
	t.Logf("successfully verified initial acquisition")

	// Keep track of the renew time after the first acquisition.
	initialRenewTime, err := time.Parse(time.RFC3339, *lease.Status.GlobalRenewTime)
	require.NoError(t, err)

	// Verification 2: Lease Renewal
	// Wait for a period longer than the RetryPeriod to ensure a renewal must have happened.
	renewalWait := retryPeriod + (1 * time.Second)
	t.Logf("waiting %s to verify lease renewal...", renewalWait)
	time.Sleep(renewalWait)

	// Get the lease again and check that the renew time has been updated.
	require.NoError(t, kubeClient.Get(ctx, key, &lease))
	require.NotNil(t, lease.Status.GlobalRenewTime, "GlobalRenewTime should not be nil after renewal")
	renewedTime, err := time.Parse(time.RFC3339, *lease.Status.GlobalRenewTime)
	require.NoError(t, err)

	require.True(t, renewedTime.After(initialRenewTime), "renewedTime (%s) should be after initialRenewTime (%s)", renewedTime, initialRenewTime)
	t.Logf("successfully verified lease renewal")

	// Wait for the test to finish
	<-testfinished
}

var gcsBucketName = "multiclusterlease-test"

func TestE2E_LeaseAlreadyHeld(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// --- Setup Kubernetes and GCS Clients ---
	cfg, err := config.GetConfig()
	require.NoError(t, err)
	err = v1alpha1.AddToScheme(scheme.Scheme)
	require.NoError(t, err)
	kubeClient, err := client.New(cfg, client.Options{Scheme: scheme.Scheme})
	require.NoError(t, err)
	gcsClient, err := storage.NewClient(ctx)
	require.NoError(t, err)

	// --- Create Test Namespace ---
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

	// --- Pre-create the Global Lock in GCS ---
	incumbentLeader := "the-incumbent-leader"
	gcsObjectKey := fmt.Sprintf("leases/%s/%s", ns.Name, leaseName)
	t.Logf("pre-creating GCS lock object '%s' with holder '%s'", gcsObjectKey, incumbentLeader)
	err = writeToGCS(ctx, gcsClient, gcsBucketName, gcsObjectKey, incumbentLeader, time.Now())
	require.NoError(t, err)
	defer func() {
		t.Logf("deleting GCS lock object '%s'", gcsObjectKey)
		deleteFromGCS(ctx, gcsClient, gcsBucketName, gcsObjectKey)
	}()

	// --- Start the Elector ---
	leaderCh := make(chan struct{})
	identity := uuid.New().String()
	lock := multiclusterleaselock.New(kubeClient, leaseName, ns.Name, identity, retryPeriod)
	elector, err := leaderelection.NewLeaderElector(leaderelection.LeaderElectionConfig{
		Lock:          lock,
		LeaseDuration: leaseDuration,
		RenewDeadline: renewDeadline,
		RetryPeriod:   retryPeriod,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				// This should never be called.
				t.Errorf("elector [%s] started leading but should have been locked out", identity)
				close(leaderCh)
			},
			OnStoppedLeading: func() {
				t.Logf("elector [%s] stopped leading", identity)
			},
		},
		Name: "e2e-test-locked-out-elector",
	})
	require.NoError(t, err)
	go elector.Run(ctx)

	// --- Verification ---
	t.Logf("elector [%s] waiting to see if it incorrectly acquires leadership...", identity)
	select {
	case <-leaderCh:
		t.Fatalf("elector [%s] acquired leadership, but should have been locked out", identity)
	case <-time.After(leaderTimeout):
		t.Logf("SUCCESS: elector [%s] did not acquire leadership within %s", identity, leaderTimeout)
	}

	// Verify that the CR status correctly reflects the incumbent leader.
	var lease v1alpha1.MultiClusterLease
	key := client.ObjectKey{Namespace: ns.Name, Name: leaseName}
	require.NoError(t, kubeClient.Get(ctx, key, &lease))
	require.NotNil(t, lease.Status.GlobalHolderIdentity)
	require.Equal(t, incumbentLeader, *lease.Status.GlobalHolderIdentity)
	t.Logf("successfully verified that status shows incumbent leader '%s'", incumbentLeader)
}

// leaseData is a helper struct for GCS object content.
type leaseData struct {
	HolderIdentity   string    `json:"holderIdentity"`
	RenewTime        time.Time `json:"renewTime"`
	LeaseTransitions int32     `json:"leaseTransitions"`
}

func writeToGCS(ctx context.Context, client *storage.Client, bucket, object, holder string, renewTime time.Time) error {
	w := client.Bucket(bucket).Object(object).NewWriter(ctx)
	w.ContentType = "application/json"
	w.CacheControl = "no-cache, no-store, must-revalidate"

	data := leaseData{
		HolderIdentity:   holder,
		RenewTime:        renewTime,
		LeaseTransitions: 1,
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if _, err := w.Write(bytes); err != nil {
		return err
	}
	return w.Close()
}

func deleteFromGCS(ctx context.Context, client *storage.Client, bucket, object string) {
	_ = client.Bucket(bucket).Object(object).Delete(ctx)
}
