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

package ha

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	operatork8s "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// failoverReplicasAnnotation marks a StatefulSet of Config Connector's manager that failover
	// has scaled to 0. Its value is the number of replicas that the StatefulSet had, which is
	// what failover scales it back to.
	failoverReplicasAnnotation = "cnrm.cloud.google.com/ha-failover-replicas"

	// replicationModeDisabled is the value of spec.experiments.multiClusterLease.resourceReplicationMode
	// with which the managers don't use KRMSyncers.
	replicationModeDisabled = "Disabled"
)

// failoverCluster is one cluster of the HA group.
type failoverCluster struct {
	handle *clusterHandle

	// gke is the cluster's GKE cluster. Its resource name is the cluster's identity in the
	// election, and a KRMSyncer names the cluster that it pulls from with it.
	gke      operatorv1beta1.GKECluster
	identity string

	// lease is the cluster's MultiClusterLease. The KRMSyncers are in its namespace.
	lease types.NamespacedName

	// syncs says whether the cluster's managers keep KRMSyncers pointed at the leader.
	syncs bool
}

// newFailoverCluster returns the cluster of a state that fetchClusterHAState has validated.
func newFailoverCluster(handle *clusterHandle, state *clusterHAState) *failoverCluster {
	mcl := multiClusterLeaseSpec(state.configConnector)
	identity, _ := mcl.CandidateIdentity()
	return &failoverCluster{
		handle:   handle,
		gke:      *mcl.GKECluster,
		identity: identity,
		lease:    types.NamespacedName{Namespace: mcl.Namespace, Name: mcl.LeaseName},
		syncs:    mcl.ResourceReplicationMode != replicationModeDisabled,
	}
}

func (c *failoverCluster) name() string {
	return c.handle.contextName
}

func (c *failoverCluster) String() string {
	return fmt.Sprintf("cluster %q (%s)", c.name(), c.identity)
}

// failoverView is what failover reads from a cluster each time it looks at it.
type failoverView struct {
	// holder is the identity that the cluster's MCL controller last saw holding the lease. It is
	// empty if the cluster has no MultiClusterLease, or if its status names no holder.
	holder string
	// renewTime is when that holder last renewed the lease, as far as the cluster's MCL
	// controller has seen. It is zero if unknown.
	renewTime time.Time

	// managers are the StatefulSets of Config Connector's manager: one in cluster mode, and one
	// per namespace in namespaced mode.
	managers []*managerStatefulSet
	// syncers are the KRMSyncers in the lease's namespace, where the managers keep theirs. It is
	// empty if the cluster doesn't sync.
	syncers []syncerv1alpha1.KRMSyncer
}

func (c *failoverCluster) view(ctx context.Context) (*failoverView, error) {
	v := &failoverView{}

	lease, err := fetchMultiClusterLease(ctx, c.handle, c.lease.Namespace, c.lease.Name)
	if err != nil {
		return nil, err
	}
	if lease != nil {
		if holder := lease.Status.GlobalHolderIdentity; holder != nil {
			v.holder = *holder
		}
		if renewTime := lease.Status.GlobalRenewTime; renewTime != nil {
			if t, err := time.Parse(time.RFC3339, *renewTime); err == nil {
				v.renewTime = t
			}
		}
	}

	if v.managers, err = c.managers(ctx); err != nil {
		return nil, err
	}

	if c.syncs {
		if v.syncers, err = listKRMSyncers(ctx, c.handle, c.lease.Namespace); err != nil {
			return nil, err
		}
	}
	return v, nil
}

// managers lists the StatefulSets of Config Connector's manager in every namespace. A manager
// that has a namespace of its own (ConfigConnectorContext spec.managerNamespace) runs outside
// cnrm-system.
func (c *failoverCluster) managers(ctx context.Context) ([]*managerStatefulSet, error) {
	list := &appsv1.StatefulSetList{}
	labels := client.MatchingLabels{operatork8s.KCCSystemComponentLabel: operatork8s.KCCControllerManagerComponent}
	if err := c.handle.client.List(ctx, list, labels); err != nil {
		return nil, fmt.Errorf("listing the StatefulSets of Config Connector's manager on cluster %q: %w", c.name(), err)
	}
	managers := make([]*managerStatefulSet, 0, len(list.Items))
	for i := range list.Items {
		managers = append(managers, &managerStatefulSet{StatefulSet: &list.Items[i]})
	}
	sort.Slice(managers, func(i, j int) bool { return managers[i].key() < managers[j].key() })
	return managers, nil
}

// managerStatefulSet is a StatefulSet of Config Connector's manager.
type managerStatefulSet struct {
	*appsv1.StatefulSet
}

func (m *managerStatefulSet) key() string {
	return m.Namespace + "/" + m.Name
}

func (m *managerStatefulSet) replicas() int32 {
	if m.Spec.Replicas == nil {
		return 1
	}
	return *m.Spec.Replicas
}

// replicasBeforeFailover returns the replicas that the manager had when failover stopped it.
// annotated is false if the StatefulSet doesn't have failoverReplicasAnnotation.
func (m *managerStatefulSet) replicasBeforeFailover() (replicas int32, annotated bool, err error) {
	value, annotated := m.Annotations[failoverReplicasAnnotation]
	if !annotated {
		return 0, false, nil
	}
	n, err := strconv.ParseInt(value, 10, 32)
	if err != nil || n < 0 {
		return 0, true, fmt.Errorf("StatefulSet %s has the annotation %s=%q, which is not a number of replicas", m.key(), failoverReplicasAnnotation, value)
	}
	return int32(n), true, nil
}

// pendingReady returns what the manager's pods are still waiting for, or "" if they are ready.
func (m *managerStatefulSet) pendingReady() string {
	switch replicas := m.replicas(); {
	case m.Status.ObservedGeneration < m.Generation:
		return fmt.Sprintf("the StatefulSet controller has not processed the latest spec of StatefulSet %s yet", m.key())
	case m.Status.ReadyReplicas < replicas:
		return fmt.Sprintf("StatefulSet %s has %d of %d pods ready", m.key(), m.Status.ReadyReplicas, replicas)
	}
	return ""
}

// pendingStopped returns what a manager that was scaled to 0 is still waiting for, or "" if its
// pods are gone. status.replicas counts a pod until it is deleted, so also while it terminates.
func (m *managerStatefulSet) pendingStopped() string {
	switch {
	case m.Status.ObservedGeneration < m.Generation:
		return fmt.Sprintf("the StatefulSet controller has not processed the latest spec of StatefulSet %s yet", m.key())
	case m.Status.Replicas > 0:
		return fmt.Sprintf("StatefulSet %s still has pods (%d)", m.key(), m.Status.Replicas)
	}
	return ""
}

// hasRunningManager reports whether at least one manager StatefulSet has replicas > 0.
func hasRunningManager(managers []*managerStatefulSet) bool {
	for _, m := range managers {
		if m.replicas() > 0 {
			return true
		}
	}
	return false
}

// stoppedByFailover returns the managers that failover scaled to 0 and has not started again.
func stoppedByFailover(managers []*managerStatefulSet) []*managerStatefulSet {
	var stopped []*managerStatefulSet
	for _, m := range managers {
		if _, annotated := m.Annotations[failoverReplicasAnnotation]; annotated && m.replicas() == 0 {
			stopped = append(stopped, m)
		}
	}
	return stopped
}

func describeManagers(managers []*managerStatefulSet) string {
	keys := make([]string, 0, len(managers))
	for _, m := range managers {
		keys = append(keys, m.key())
	}
	if len(keys) == 1 {
		return "StatefulSet " + keys[0]
	}
	return "StatefulSets " + strings.Join(keys, ", ")
}

// managerPatch is a merge patch that scales a manager and sets failoverReplicasAnnotation, or
// removes the annotation if it is nil.
func managerPatch(replicas int32, annotation *string) (client.Patch, error) {
	data, err := json.Marshal(map[string]any{
		"metadata": map[string]any{"annotations": map[string]any{failoverReplicasAnnotation: annotation}},
		"spec":     map[string]any{"replicas": replicas},
	})
	if err != nil {
		return nil, err
	}
	return client.RawPatch(types.MergePatchType, data), nil
}

// stopManagers scales the cluster's managers to 0. It keeps the replicas that each one had in
// failoverReplicasAnnotation, so that startManagers can undo it: in this run, or in a later one
// if this one is interrupted.
//
// The operator owns these StatefulSets, but its manifests don't set spec.replicas and it applies
// them with a client-side apply, so it leaves the replicas and the annotation as they are.
//
// A manager that is already at 0 is left alone, whether failover or someone else stopped it.
// waiting is true if any manager was scaled down or still has pods terminating.
func (c *failoverCluster) stopManagers(ctx context.Context, out io.Writer) (waiting bool, err error) {
	managers, err := c.managers(ctx)
	if err != nil {
		return false, err
	}
	for _, m := range managers {
		replicas := m.replicas()
		if replicas == 0 {
			if m.pendingStopped() != "" {
				waiting = true
			}
			continue
		}
		annotation := strconv.Itoa(int(replicas))
		patch, err := managerPatch(0, &annotation)
		if err != nil {
			return false, err
		}
		if err := c.handle.client.Patch(ctx, m.StatefulSet, patch); err != nil {
			return false, fmt.Errorf("stopping StatefulSet %s on cluster %q: %w", m.key(), c.name(), err)
		}
		waiting = true
		fmt.Fprintf(out, "cluster %q: stopped Config Connector's manager (StatefulSet %s, scaled from %d to 0)\n", c.name(), m.key(), replicas)
	}
	return waiting, nil
}

// startManagers undoes stopManagers: it scales every manager that has failoverReplicasAnnotation
// back to the replicas in it, and removes the annotation. It goes on after a failure, so that
// one StatefulSet doesn't keep the others stopped.
func (c *failoverCluster) startManagers(ctx context.Context, out io.Writer) error {
	managers, err := c.managers(ctx)
	if err != nil {
		return fmt.Errorf("%w; any managers that failover stopped on cluster %q are still scaled to 0 — run this command again once the cluster is reachable", err, c.name())
	}
	var errs []error
	for _, m := range managers {
		replicas, annotated, err := m.replicasBeforeFailover()
		if !annotated {
			continue
		}
		if err != nil {
			errs = append(errs, fmt.Errorf("cluster %q: %w; scale the StatefulSet by hand and remove the annotation", c.name(), err))
			continue
		}
		// If someone has started the manager in the meantime, only the annotation is left to
		// remove.
		current := m.replicas()
		if current > 0 {
			replicas = current
		}
		patch, err := managerPatch(replicas, nil)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if err := c.handle.client.Patch(ctx, m.StatefulSet, patch); err != nil {
			errs = append(errs, fmt.Errorf(
				"starting StatefulSet %s on cluster %q again: %w; run this command again, or start it with `kubectl --context %s --namespace %s scale statefulset %s --replicas=%d`",
				m.key(), c.name(), err, c.name(), m.Namespace, m.Name, replicas))
			continue
		}
		if current == 0 {
			fmt.Fprintf(out, "cluster %q: started Config Connector's manager again (StatefulSet %s, scaled from 0 to %d)\n", c.name(), m.key(), replicas)
		}
	}
	return errors.Join(errs...)
}

// pendingStop returns which of the cluster's managers still has a pod after stopManagers, or ""
// once none has. A manager stops renewing the lease only when its pod is gone.
func (c *failoverCluster) pendingStop(ctx context.Context) (string, error) {
	managers, err := c.managers(ctx)
	if err != nil {
		return err.Error(), nil
	}
	for _, m := range managers {
		if m.replicas() > 0 {
			return "", fmt.Errorf("StatefulSet %s runs again: something scaled it up, or created it, after failover stopped the managers", m.key())
		}
		if pending := m.pendingStopped(); pending != "" {
			return pending, nil
		}
	}
	return "", nil
}

// pendingAsStandby returns what the cluster is still waiting for to be a standby of leader: its
// managers run, and their KRMSyncers pull from the leader. It returns "" if nothing is pending.
func (c *failoverCluster) pendingAsStandby(v *failoverView, leader *failoverCluster) string {
	for _, m := range v.managers {
		if m.replicas() == 0 {
			continue
		}
		if pending := m.pendingReady(); pending != "" {
			return pending
		}
	}
	return c.pendingSyncers(v, leader)
}

// pendingAsTarget is pendingAsStandby for the cluster that is to take over. That cluster also
// needs a manager that runs.
func (c *failoverCluster) pendingAsTarget(v *failoverView, leader *failoverCluster) string {
	if len(v.managers) == 0 {
		return "it has no StatefulSet of Config Connector's manager"
	}
	if !hasRunningManager(v.managers) {
		return fmt.Sprintf("none of its Config Connector managers runs (%s scaled to 0)", describeManagers(v.managers))
	}
	return c.pendingAsStandby(v, leader)
}

// pendingSyncers returns what the KRMSyncers of the cluster's running managers are still waiting
// for, or "" if nothing. On every cluster they name the leader as their remote cluster. They
// pull from it on a standby, and they are suspended on the leader itself.
//
// Other KRMSyncers in the namespace are ignored: nothing deletes the KRMSyncer of a manager that
// is gone, so there can be stale ones.
func (c *failoverCluster) pendingSyncers(v *failoverView, leader *failoverCluster) string {
	if !c.syncs {
		return ""
	}
	syncers := make(map[string]*syncerv1alpha1.KRMSyncer, len(v.syncers))
	for i := range v.syncers {
		syncers[v.syncers[i].Name] = &v.syncers[i]
	}
	wantSuspended := c == leader
	for _, m := range v.managers {
		if m.replicas() == 0 {
			continue
		}
		key := c.lease.Namespace + "/" + c.syncerName(m)
		syncer, found := syncers[c.syncerName(m)]
		switch {
		case !found:
			return fmt.Sprintf("KRMSyncer %s, which the manager in StatefulSet %s maintains, does not exist yet", key, m.key())
		case !syncerNames(syncer, leader.gke):
			return fmt.Sprintf("KRMSyncer %s does not name %s as its remote cluster yet", key, leader.identity)
		case wantSuspended && !syncer.Spec.Suspend:
			return fmt.Sprintf("KRMSyncer %s is not suspended yet", key)
		case !wantSuspended && syncer.Spec.Suspend:
			return fmt.Sprintf("KRMSyncer %s is still suspended", key)
		}
	}
	return ""
}

// syncerName returns the name of the KRMSyncer that a manager maintains in the lease's
// namespace: the lease's name in cluster mode, and <lease>-<namespace> for the manager of a
// namespace in namespaced mode. It must match SyncerIntegration.getDesiredSyncers in
// pkg/controller/kccmanager.
func (c *failoverCluster) syncerName(m *managerStatefulSet) string {
	if namespace := m.Labels[operatork8s.NamespacedComponentLabel]; namespace != "" {
		return c.lease.Name + "-" + namespace
	}
	return c.lease.Name
}

// syncerNames reports whether the KRMSyncer's remote cluster is the given GKE cluster.
func syncerNames(syncer *syncerv1alpha1.KRMSyncer, cluster operatorv1beta1.GKECluster) bool {
	if syncer.Spec.Remote == nil || syncer.Spec.Remote.GKECluster == nil {
		return false
	}
	remote := syncer.Spec.Remote.GKECluster
	return remote.Project == cluster.Project && remote.Location == cluster.Location && remote.Name == cluster.Name
}
