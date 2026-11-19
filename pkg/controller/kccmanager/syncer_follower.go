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
	"sync"
	"sync/atomic"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
)

const (
	// standbyResyncPeriod is how often a standby checks that its KRMSyncers pull from the
	// leader.
	standbyResyncPeriod = time.Minute

	// pullFromLeaderTimeout bounds one update of the KRMSyncers.
	pullFromLeaderTimeout = 30 * time.Second

	// leaderFollowerEventSource is the component that the events on the KRMSyncers come from.
	leaderFollowerEventSource = "cnrm-controller-manager"

	// Reasons of the events that a standby records on its KRMSyncers.
	reasonPullFromLeaderFailed = "PullFromLeaderFailed"
	reasonPullingFromLeader    = "PullingFromLeader"
)

// leaderFollower keeps the KRMSyncers of a standby cluster pulling from the leader. It updates
// them when the leader changes and checks them again every resyncPeriod, so that a failed
// update is retried, and a KRMSyncer that was changed or deleted is repaired.
type leaderFollower struct {
	syncers      *SyncerIntegration
	recorder     record.EventRecorder
	resyncPeriod time.Duration

	// leader is the cluster to pull from. It is nil when there is nothing to pull from: no
	// cluster leads, or this one does.
	leader atomic.Pointer[operatorv1beta1.GKECluster]
	// leaderChanged wakes up run.
	leaderChanged chan struct{}

	// mu is held while the KRMSyncers are updated; see stop.
	mu sync.Mutex
	// stopped is set when this cluster starts leading.
	stopped bool
	// failing is true if the last update failed.
	failing bool
}

func newLeaderFollower(syncers *SyncerIntegration, recorder record.EventRecorder) *leaderFollower {
	return &leaderFollower{
		syncers:       syncers,
		recorder:      recorder,
		resyncPeriod:  standbyResyncPeriod,
		leaderChanged: make(chan struct{}, 1),
	}
}

// setLeader sets the cluster to pull from and triggers an update of the KRMSyncers. A nil
// leader means that there is nothing to pull from.
func (f *leaderFollower) setLeader(leader *operatorv1beta1.GKECluster) {
	f.leader.Store(leader)
	select {
	case f.leaderChanged <- struct{}{}:
	default: // An update is already pending; it will use the new leader.
	}
}

// run updates the KRMSyncers when the leader changes and every resyncPeriod, until ctx is done.
func (f *leaderFollower) run(ctx context.Context) {
	ticker := time.NewTicker(f.resyncPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-f.leaderChanged:
		case <-ticker.C:
		}
		f.pullFromLeader(ctx)
	}
}

// stop is called when this cluster starts leading. It waits for an update that is in flight,
// and makes later ones do nothing. Otherwise, an update for the previous leader could resume
// the KRMSyncers after this cluster has suspended them, and they would overwrite the resources
// that this cluster now reconciles with another cluster's copy.
func (f *leaderFollower) stop() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.stopped = true
}

// pullFromLeader makes the KRMSyncers pull from the leader. A failure is logged and recorded
// as an event on the KRMSyncers; the next run retries.
func (f *leaderFollower) pullFromLeader(ctx context.Context) {
	f.mu.Lock()
	defer f.mu.Unlock()

	leader := f.leader.Load()
	if f.stopped || leader == nil {
		return
	}

	// The timeout also covers the events, so that stop never waits longer than that.
	updateCtx, cancel := context.WithTimeout(ctx, pullFromLeaderTimeout)
	defer cancel()

	if err := f.syncers.EnsurePullingFromLeader(updateCtx, leader); err != nil {
		if ctx.Err() != nil {
			return // The manager is shutting down.
		}
		klog.Errorf("error ensuring syncer is pulling from leader %s (will retry within %v): %v", leader.ResourceName(), f.resyncPeriod, err)
		f.failing = true
		f.recordEvent(updateCtx, corev1.EventTypeWarning, reasonPullFromLeaderFailed,
			fmt.Sprintf("Could not make this KRMSyncer pull from the leader %s; will retry: %v", leader.ResourceName(), err))
		return
	}

	if f.failing {
		f.failing = false
		f.recordEvent(updateCtx, corev1.EventTypeNormal, reasonPullingFromLeader,
			fmt.Sprintf("This KRMSyncer now pulls from the leader %s", leader.ResourceName()))
	}
}

// recordEvent records an event on each KRMSyncer that this manager keeps up to date. That is
// one, unless the manager watches several namespaces.
func (f *leaderFollower) recordEvent(ctx context.Context, eventType, reason, message string) {
	for _, name := range f.syncers.getDesiredSyncers() {
		f.recorder.Event(f.eventTarget(ctx, name), eventType, reason, message)
	}
}

// eventTarget returns the object to record an event on. It is the KRMSyncer itself if it can
// be read, so that the event has its UID and shows up in "kubectl describe". Otherwise, for
// example when the KRMSyncer couldn't be created, it is a reference to it by name.
func (f *leaderFollower) eventTarget(ctx context.Context, name types.NamespacedName) runtime.Object {
	syncer := &unstructured.Unstructured{}
	syncer.SetGroupVersionKind(KRMSyncerGVK)
	if err := f.syncers.apiReader.Get(ctx, name, syncer); err == nil {
		return syncer
	}
	return &corev1.ObjectReference{
		APIVersion: KRMSyncerGVK.GroupVersion().String(),
		Kind:       KRMSyncerGVK.Kind,
		Namespace:  name.Namespace,
		Name:       name.Name,
	}
}
