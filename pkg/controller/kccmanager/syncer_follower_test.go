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
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

// A standby keeps its KRMSyncer pulling from the leader: it retries a failed update, reports
// the failure and the recovery as events on the KRMSyncer, and repairs a KRMSyncer that was
// changed.
func TestSyncerIntegration_StandbyKeepsPullingFromLeader(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, kubeClient := startSyncerEnvtest(t)

	nsName := "cnrm-system"
	if err := kubeClient.Create(ctx, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: nsName}}); err != nil {
		t.Fatalf("error creating namespace: %v", err)
	}

	// The manager's cluster is a standby from the start: its lease names it as a candidate, and
	// reports that another cluster leads.
	cluster := testGKECluster("standby")
	identity := cluster.ResourceName()
	leader := testGKECluster("leader")
	leaderIdentity := leader.ResourceName()
	leaseDuration := int32(15)
	lease := &mclv1alpha1.MultiClusterLease{
		ObjectMeta: metav1.ObjectMeta{Name: testLeaseName, Namespace: nsName},
		Spec: mclv1alpha1.MultiClusterLeaseSpec{
			HolderIdentity:       &identity,
			LeaseDurationSeconds: &leaseDuration,
			RenewTime:            &metav1.MicroTime{Time: time.Now()},
		},
	}
	if err := kubeClient.Create(ctx, lease); err != nil {
		t.Fatalf("error creating MultiClusterLease: %v", err)
	}
	if err := (&mockLeaseController{client: kubeClient}).elect(ctx, leaderIdentity, nsName); err != nil {
		t.Fatalf("error electing the leader: %v", err)
	}

	mgr, err := New(context.Background(), cfg, Config{
		ManagerOptions: manager.Options{
			Metrics:                metricsserver.Options{BindAddress: "0"},
			HealthProbeBindAddress: "0",
			LeaderElection:         false,
			LeaderElectionID:       cluster.Name,
		},
		MultiClusterLease: true,
		SyncerIntegration: true,
		testConfig: testConfig{
			skipControllerRegistration: true,
			multiClusterLeaseConfig: &operatorv1beta1.MultiClusterLeaseSpec{
				LeaseName:  testLeaseName,
				Namespace:  nsName,
				GKECluster: cluster,
			},
			suppressExitOnLeadershipLoss: true,
		},
	})
	if err != nil {
		t.Fatalf("error creating new manager: %v", err)
	}
	lem, ok := mgr.(*leaderElectionManager)
	if !ok {
		t.Fatalf("manager is not *leaderElectionManager, it is %T", mgr)
	}
	lem.follower.resyncPeriod = 500 * time.Millisecond

	// At first, the standby can't write its KRMSyncer.
	var failWrites atomic.Bool
	failWrites.Store(true)
	lem.syncerInteg.client = &hookedClient{Client: lem.syncerInteg.client, beforeWrite: func() error {
		if failWrites.Load() {
			return errors.New("injected failure")
		}
		return nil
	}}

	go func() {
		if err := lem.Start(ctx); err != nil {
			t.Errorf("mgr failed to start: %v", err)
		}
	}()

	// The failure is reported on the KRMSyncer, which doesn't exist yet.
	syncerNN := types.NamespacedName{Namespace: nsName, Name: testLeaseName}
	warning := waitForSyncerEvent(ctx, t, kubeClient, syncerNN, corev1.EventTypeWarning, reasonPullFromLeaderFailed)
	for _, want := range []string{leaderIdentity, "injected failure"} {
		if !strings.Contains(warning.Message, want) {
			t.Errorf("message of the %s event is %q, want it to contain %q", warning.Reason, warning.Message, want)
		}
	}
	syncer := &unstructured.Unstructured{}
	syncer.SetGroupVersionKind(KRMSyncerGVK)
	if err := kubeClient.Get(ctx, syncerNN, syncer); !apierrors.IsNotFound(err) {
		t.Fatalf("getting KRMSyncer %s: got error %v, want a NotFound error", syncerNN, err)
	}

	// Writes work again. The leader hasn't changed, so it is the periodic retry that creates
	// the KRMSyncer. The recovery is reported on it.
	failWrites.Store(false)
	created := waitForSyncer(ctx, t, kubeClient, syncerNN, leader, false)
	normal := waitForSyncerEvent(ctx, t, kubeClient, syncerNN, corev1.EventTypeNormal, reasonPullingFromLeader)
	if !strings.Contains(normal.Message, leaderIdentity) {
		t.Errorf("message of the %s event is %q, want it to contain %q", normal.Reason, normal.Message, leaderIdentity)
	}
	if got, want := normal.InvolvedObject.UID, created.GetUID(); got != want {
		t.Errorf("the %s event is about an object with UID %q, want the UID of the KRMSyncer, %q", normal.Reason, got, want)
	}

	// Someone suspends the KRMSyncer and points it at another cluster. The next periodic check
	// repairs it.
	changed := created.DeepCopy()
	_ = unstructured.SetNestedField(changed.Object, true, "spec", "suspend")
	_ = unstructured.SetNestedMap(changed.Object, testRemote(testGKECluster("another-cluster"), "Default"), "spec", "remote")
	if err := kubeClient.Update(ctx, changed); err != nil {
		t.Fatalf("error changing KRMSyncer %s: %v", syncerNN, err)
	}
	waitForSyncer(ctx, t, kubeClient, syncerNN, leader, false)
}

// A standby's KRMSyncer follows the leader until this cluster starts leading. From then on the
// follower leaves it alone, whether an update for another leader is in flight or comes later.
func TestLeaderFollower_StopsWhenLeading(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, kubeClient := startSyncerEnvtest(t)

	nsName := "cnrm-system"
	if err := kubeClient.Create(ctx, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: nsName}}); err != nil {
		t.Fatalf("error creating namespace: %v", err)
	}

	// The next write to the KRMSyncer waits for release, once holdNextWrite is set.
	var holdNextWrite atomic.Bool
	held := make(chan struct{})
	release := make(chan struct{})

	self := testGKECluster("cluster-b")
	syncerNN := types.NamespacedName{Namespace: nsName, Name: testLeaseName}
	si := &SyncerIntegration{
		client: &hookedClient{Client: kubeClient, beforeWrite: func() error {
			if holdNextWrite.CompareAndSwap(true, false) {
				close(held)
				<-release
			}
			return nil
		}},
		apiReader:       kubeClient,
		leaseNN:         syncerNN,
		replicationMode: "Status",
		cluster:         self,
	}
	recorder := record.NewFakeRecorder(10)
	follower := newLeaderFollower(si, recorder)
	// Only a leader change triggers an update in this test.
	follower.resyncPeriod = time.Hour
	go follower.run(ctx)

	// The KRMSyncer is created when a leader is observed, and follows a leader change.
	leaderA := testGKECluster("cluster-a")
	follower.setLeader(leaderA)
	waitForSyncer(ctx, t, kubeClient, syncerNN, leaderA, false)

	leaderC := testGKECluster("cluster-c")
	follower.setLeader(leaderC)
	waitForSyncer(ctx, t, kubeClient, syncerNN, leaderC, false)

	// An update for yet another leader is in flight when this cluster starts leading. stop
	// waits for it, so that it can't land after the KRMSyncer is suspended.
	holdNextWrite.Store(true)
	follower.setLeader(testGKECluster("cluster-d"))
	select {
	case <-held:
	case <-time.After(20 * time.Second):
		t.Fatal("timed out waiting for the update of the KRMSyncer")
	}
	stopped := make(chan struct{})
	go func() {
		follower.stop()
		close(stopped)
	}()
	select {
	case <-stopped:
		t.Fatal("stop returned while an update of the KRMSyncer was in flight")
	case <-time.After(500 * time.Millisecond):
	}
	close(release)
	select {
	case <-stopped:
	case <-time.After(20 * time.Second):
		t.Fatal("timed out waiting for stop to return")
	}

	// This cluster leads: the KRMSyncer is suspended and names this cluster.
	if err := si.EnsureSuspended(ctx); err != nil {
		t.Fatalf("EnsureSuspended: %v", err)
	}
	suspended := waitForSyncer(ctx, t, kubeClient, syncerNN, self, true)

	// An update for another leader that comes later doesn't touch the KRMSyncer.
	follower.setLeader(testGKECluster("cluster-e"))
	follower.pullFromLeader(ctx)
	got := &unstructured.Unstructured{}
	got.SetGroupVersionKind(KRMSyncerGVK)
	if err := kubeClient.Get(ctx, syncerNN, got); err != nil {
		t.Fatalf("error getting KRMSyncer %s: %v", syncerNN, err)
	}
	assertSyncer(t, got, self, "Default", true)
	if got, want := got.GetResourceVersion(), suspended.GetResourceVersion(); got != want {
		t.Errorf("the KRMSyncer was updated after this cluster started leading: resourceVersion %s, want %s", got, want)
	}

	// No update failed, so there is nothing to report.
	select {
	case event := <-recorder.Events:
		t.Errorf("unexpected event: %s", event)
	default:
	}
}

// A standby pulls from the cluster that the leader's identity names. It has nothing to pull
// from when no cluster leads or when it leads itself, and it exits when the identity doesn't
// name a GKE cluster.
func TestOnNewLeader(t *testing.T) {
	self := testGKECluster("self")
	leader := testGKECluster("leader")

	for _, tc := range []struct {
		name       string
		identity   string
		wantLeader *operatorv1beta1.GKECluster
		wantFatal  string
	}{
		{name: "another cluster leads", identity: leader.ResourceName(), wantLeader: leader},
		{name: "no cluster leads", identity: "", wantLeader: nil},
		{name: "this cluster leads", identity: self.ResourceName(), wantLeader: nil},
		{
			name:      "the identity doesn't name a GKE cluster",
			identity:  "legacy-cluster",
			wantFatal: `cannot sync from leader "legacy-cluster"`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			follower := newLeaderFollower(&SyncerIntegration{}, &record.FakeRecorder{})
			follower.setLeader(testGKECluster("previous-leader"))

			var fatal string
			m := &leaderElectionManager{
				leConfig: &leaderelection.LeaderElectionConfig{Lock: &MockLock{identity: self.ResourceName()}},
				follower: follower,
				onFatal:  func(format string, args ...interface{}) { fatal = fmt.Sprintf(format, args...) },
			}
			m.onNewLeader(tc.identity)

			if tc.wantFatal != "" {
				if !strings.Contains(fatal, tc.wantFatal) {
					t.Errorf("fatal error is %q, want it to contain %q", fatal, tc.wantFatal)
				}
				return
			}
			if fatal != "" {
				t.Errorf("unexpected fatal error: %s", fatal)
			}
			if got := follower.leader.Load(); !reflect.DeepEqual(got, tc.wantLeader) {
				t.Errorf("the cluster to pull from is %v, want %v", got, tc.wantLeader)
			}
		})
	}
}

// hookedClient calls beforeWrite before each Create and Update, and fails the call if it
// returns an error.
type hookedClient struct {
	client.Client
	beforeWrite func() error
}

func (c *hookedClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	if err := c.beforeWrite(); err != nil {
		return err
	}
	return c.Client.Create(ctx, obj, opts...)
}

func (c *hookedClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	if err := c.beforeWrite(); err != nil {
		return err
	}
	return c.Client.Update(ctx, obj, opts...)
}

// waitForSyncer waits until the KRMSyncer names the given cluster as its remote and has the
// given suspend flag, and returns it.
func waitForSyncer(ctx context.Context, t *testing.T, c client.Client, nn types.NamespacedName, wantCluster *operatorv1beta1.GKECluster, wantSuspend bool) *unstructured.Unstructured {
	t.Helper()
	syncer := &unstructured.Unstructured{}
	syncer.SetGroupVersionKind(KRMSyncerGVK)
	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 20*time.Second, true, func(ctx context.Context) (bool, error) {
		if err := c.Get(ctx, nn, syncer); err != nil {
			return false, nil
		}
		remote, _, _ := unstructured.NestedMap(syncer.Object, "spec", "remote")
		suspend, _, _ := unstructured.NestedBool(syncer.Object, "spec", "suspend")
		return reflect.DeepEqual(remote, testRemote(wantCluster, "Default")) && suspend == wantSuspend, nil
	})
	if err != nil {
		t.Fatalf("timed out waiting for KRMSyncer %s to have the remote %s and suspend=%v; its spec is %v", nn, wantCluster.ResourceName(), wantSuspend, syncer.Object["spec"])
	}
	return syncer
}

// waitForSyncerEvent waits for an event of the given type and reason about the KRMSyncer, and
// returns it.
func waitForSyncerEvent(ctx context.Context, t *testing.T, c client.Client, nn types.NamespacedName, eventType, reason string) *corev1.Event {
	t.Helper()
	var found *corev1.Event
	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 20*time.Second, true, func(ctx context.Context) (bool, error) {
		events := &corev1.EventList{}
		if err := c.List(ctx, events, client.InNamespace(nn.Namespace)); err != nil {
			return false, nil
		}
		for i := range events.Items {
			e := &events.Items[i]
			about := e.InvolvedObject
			if about.APIVersion == KRMSyncerGVK.GroupVersion().String() && about.Kind == KRMSyncerGVK.Kind && about.Name == nn.Name &&
				e.Type == eventType && e.Reason == reason {
				found = e
				return true, nil
			}
		}
		return false, nil
	})
	if err != nil {
		t.Fatalf("timed out waiting for a %s event with reason %s about KRMSyncer %s", eventType, reason, nn)
	}
	return found
}
