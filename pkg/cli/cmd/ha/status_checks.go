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
	"fmt"
	"sort"
	"strings"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// defaultLeaseDuration matches the lease duration that the KCC manager uses for the
	// multi-cluster lease. It is used when the MultiClusterLease does not report one.
	defaultLeaseDuration = 15 * time.Second

	// leaseExpiryFactor is how many lease durations may pass since the last renewal before
	// the lease is reported as expired. It allows for delay in MultiClusterLease status updates.
	leaseExpiryFactor = 2

	replicationModeStatus   = "Status"
	replicationModeDisabled = "Disabled"
)

// haProblems checks a group of clusters that all have multi-cluster HA configured, and
// returns one problem per issue found. It checks, in order:
//  1. the clusters take part in the same election (validateMultiClusterLeaseConfigs);
//  2. each cluster's MultiClusterLease has a live holder, and all clusters agree on it;
//  3. each cluster's KRMSyncers: suspended on the leader, syncing from the leader on standbys.
//
// Later checks are skipped when an earlier one fails, because they depend on it.
func haProblems(states []*clusterHAState, now time.Time) []string {
	if err := validateMultiClusterLeaseConfigs(states); err != nil {
		return splitErrorLines(err)
	}

	problems := replicationModeProblems(states)

	leader, leaseProblems := leaseHolder(states, now)
	problems = append(problems, leaseProblems...)
	if leader == "" {
		return problems
	}

	for _, s := range states {
		problems = append(problems, syncerProblems(s, leader)...)
	}
	return problems
}

// replicationModeProblems reports clusters that replicate resources differently.
func replicationModeProblems(states []*clusterHAState) []string {
	modes := map[string][]string{}
	for _, s := range states {
		mode := replicationMode(multiClusterLeaseSpec(s.configConnector))
		modes[mode] = append(modes[mode], s.contextName)
	}
	if len(modes) <= 1 {
		return nil
	}
	return []string{fmt.Sprintf(
		"mismatched resourceReplicationMode across clusters (%s): set spec.experiments.multiClusterLease.resourceReplicationMode to the same value on every cluster",
		formatClusterLeaseValues(states, replicationMode),
	)}
}

func replicationMode(mcl *operatorv1beta1.MultiClusterLeaseSpec) string {
	if mcl == nil || mcl.ResourceReplicationMode == "" {
		return replicationModeStatus
	}
	return mcl.ResourceReplicationMode
}

// leaseHolder returns the identity of the cluster that holds the lease, or "" if it can't be
// determined, together with the lease problems found.
func leaseHolder(states []*clusterHAState, now time.Time) (string, []string) {
	var problems []string
	holders := map[string][]string{} // holder identity -> contexts that report it
	for _, s := range states {
		holder, problem := liveLeaseHolder(s, now)
		if problem != "" {
			problems = append(problems, problem)
			continue
		}
		holders[holder] = append(holders[holder], s.contextName)
	}
	if len(problems) > 0 {
		return "", problems
	}

	if len(holders) > 1 {
		return "", []string{fmt.Sprintf("clusters disagree on the lease holder (%s)", formatHolders(holders))}
	}

	var holder string
	for h := range holders {
		holder = h
	}
	if !isInspectedCluster(states, holder) {
		return "", []string{fmt.Sprintf("the lease is held by %q, which is not one of the clusters in --%s", holder, contextsFlag)}
	}
	return holder, nil
}

const (
	// staleCandidateMessage is the error string that the multicluster-leader-election
	// controller records in BackendHealthy=False when a cluster's spec.renewTime is older
	// than its staleness threshold.
	//
	// TODO: remove this workaround once multicluster-leader-election stops reporting a
	// stale standby candidate as BackendHealthy=False.
	staleCandidateMessage = "candidate lease is stale"
)

// liveLeaseHolder returns the holder of the cluster's MultiClusterLease, or a problem if the
// lease is missing, its backend is unhealthy, it has no holder, or it has expired.
func liveLeaseHolder(s *clusterHAState, now time.Time) (string, string) {
	mcl := multiClusterLeaseSpec(s.configConnector)
	lease := s.multiClusterLease
	if lease == nil {
		return "", fmt.Sprintf("cluster %q: MultiClusterLease %s/%s not found; check that the multiclusterlease-controller-manager is running", s.contextName, mcl.Namespace, mcl.LeaseName)
	}

	var holder string
	if lease.Status.GlobalHolderIdentity != nil {
		holder = *lease.Status.GlobalHolderIdentity
	}
	candidate, _ := mcl.CandidateIdentity()

	if cond := meta.FindStatusCondition(lease.Status.Conditions, string(mclv1alpha1.ConditionTypeBackendHealthy)); cond != nil && cond.Status == metav1.ConditionFalse {
		if !isStandbyStaleCandidate(cond, candidate, holder) {
			return "", fmt.Sprintf("cluster %q: the lease backend is unhealthy: %s", s.contextName, cond.Message)
		}
	}

	if holder == "" {
		return "", fmt.Sprintf("cluster %q: no cluster holds the lease %s/%s", s.contextName, mcl.Namespace, mcl.LeaseName)
	}

	if expired, detail := leaseExpired(lease, now); expired {
		return "", fmt.Sprintf("cluster %q: the lease held by %q %s", s.contextName, holder, detail)
	}
	return holder, ""
}

// isStandbyStaleCandidate reports whether a BackendHealthy=False condition is the
// multicluster-leader-election controller reporting that a standby's spec.renewTime is
// stale.
//
// A standby doesn't refresh spec.renewTime while another cluster holds a live lease. The
// MCL controller still reads the backend, updates status.globalHolderIdentity and
// status.globalRenewTime, and then returns "candidate lease is stale", which its reconciler
// records as BackendHealthy=False. If the backend read itself fails, the controller returns
// "failed to read lease: ..." instead, so real backend errors are not ignored. On the
// leader (candidate == holder), a stale spec.renewTime means the leader's manager stopped
// heartbeating and is not ignored.
//
// TODO: remove this workaround once multicluster-leader-election stops reporting a stale
// standby candidate as BackendHealthy=False.
func isStandbyStaleCandidate(cond *metav1.Condition, candidate, holder string) bool {
	if holder == "" || candidate == holder {
		return false
	}
	return strings.Contains(cond.Message, staleCandidateMessage)
}

// leaseExpired reports whether the lease has not been renewed for leaseExpiryFactor lease
// durations. A lease without a renew time is not reported as expired.
func leaseExpired(lease *mclv1alpha1.MultiClusterLease, now time.Time) (bool, string) {
	if lease.Status.GlobalRenewTime == nil || *lease.Status.GlobalRenewTime == "" {
		return false, ""
	}
	renewTime, err := time.Parse(time.RFC3339, *lease.Status.GlobalRenewTime)
	if err != nil {
		return true, fmt.Sprintf("has an unparseable renew time %q", *lease.Status.GlobalRenewTime)
	}

	duration := defaultLeaseDuration
	if d := lease.Status.GlobalLeaseDurationSeconds; d != nil && *d > 0 {
		duration = time.Duration(*d) * time.Second
	} else if d := lease.Spec.LeaseDurationSeconds; d != nil && *d > 0 {
		duration = time.Duration(*d) * time.Second
	}

	if now.Sub(renewTime) > leaseExpiryFactor*duration {
		return true, fmt.Sprintf("has expired (last renewed at %s, lease duration %s)", renewTime.UTC().Format(time.RFC3339), duration)
	}
	return false, ""
}

func isInspectedCluster(states []*clusterHAState, identity string) bool {
	for _, s := range states {
		// validateMultiClusterLeaseConfigs checked that every identity is valid.
		if id, _ := multiClusterLeaseSpec(s.configConnector).CandidateIdentity(); id == identity {
			return true
		}
	}
	return false
}

// syncerProblems checks the cluster's KRMSyncers for the lease, using the Ready condition that
// the syncer controller reports. The leader's syncers must be suspended; a standby's syncers
// must be syncing from the leader.
func syncerProblems(s *clusterHAState, leader string) []string {
	mcl := multiClusterLeaseSpec(s.configConnector)
	if replicationMode(mcl) == replicationModeDisabled {
		return nil
	}

	identity, _ := mcl.CandidateIdentity()
	isLeader := identity == leader
	syncers := leaseSyncers(s.krmSyncers, mcl.LeaseName)
	if len(syncers) == 0 {
		if isLeader {
			// The leader doesn't sync; its KCC manager only creates suspended syncers.
			return nil
		}
		return []string{fmt.Sprintf("cluster %q: no KRMSyncer for lease %q found in namespace %q", s.contextName, mcl.LeaseName, mcl.Namespace)}
	}

	var problems []string
	for _, syncer := range syncers {
		id := syncer.Namespace + "/" + syncer.Name
		if isLeader {
			if !syncer.Spec.Suspend {
				problems = append(problems, fmt.Sprintf("cluster %q (leader): KRMSyncer %s is not suspended", s.contextName, id))
			}
			continue
		}

		if syncer.Spec.Suspend {
			problems = append(problems, fmt.Sprintf("cluster %q: KRMSyncer %s is suspended, so it is not syncing from the leader", s.contextName, id))
			continue
		}
		if remote := syncerRemote(syncer); remote != leader {
			problems = append(problems, fmt.Sprintf("cluster %q: KRMSyncer %s syncs from %q, but the lease is held by %q", s.contextName, id, remote, leader))
			continue
		}
		if problem := syncerReadyProblem(syncer); problem != "" {
			problems = append(problems, fmt.Sprintf("cluster %q: KRMSyncer %s %s", s.contextName, id, problem))
		}
	}
	return problems
}

// syncerReadyProblem describes why a standby KRMSyncer is not syncing, based on its Ready
// condition, or returns "" if it is syncing.
func syncerReadyProblem(syncer syncerv1alpha1.KRMSyncer) string {
	ready := meta.FindStatusCondition(syncer.Status.Conditions, syncerv1alpha1.ConditionReady)
	switch {
	case ready == nil:
		return "has not reported a Ready condition; check that the syncer-controller-manager is running"
	case syncer.Status.ObservedGeneration < syncer.Generation:
		return "has not processed its latest spec yet; check that the syncer-controller-manager is running"
	case ready.Status == metav1.ConditionTrue && ready.Reason == syncerv1alpha1.ReasonSyncing:
		return ""
	case ready.Reason == syncerv1alpha1.ReasonSyncFailed:
		return fmt.Sprintf("is failing to sync %d objects (%s): %s", syncer.Status.FailedObjects, ready.Reason, ready.Message)
	default:
		return fmt.Sprintf("is not syncing (%s): %s", ready.Reason, ready.Message)
	}
}

// leaseSyncers returns the KRMSyncers that the KCC manager creates for the lease:
// "<leaseName>" in cluster mode and "<leaseName>-<namespace>" in namespaced mode.
func leaseSyncers(syncers []syncerv1alpha1.KRMSyncer, leaseName string) []syncerv1alpha1.KRMSyncer {
	var out []syncerv1alpha1.KRMSyncer
	for _, s := range syncers {
		if s.Name == leaseName || strings.HasPrefix(s.Name, leaseName+"-") {
			out = append(out, s)
		}
	}
	return out
}

// syncerRemote returns the resource name of the GKE cluster a KRMSyncer syncs from.
func syncerRemote(s syncerv1alpha1.KRMSyncer) string {
	if s.Spec.Remote == nil || s.Spec.Remote.GKECluster == nil {
		return ""
	}
	r := s.Spec.Remote.GKECluster
	gke := operatorv1beta1.GKECluster{Project: r.Project, Location: r.Location, Name: r.Name}
	return gke.ResourceName()
}

func formatHolders(holders map[string][]string) string {
	keys := make([]string, 0, len(holders))
	for h := range holders {
		keys = append(keys, h)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, h := range keys {
		parts = append(parts, fmt.Sprintf("%q reported by [%s]", h, strings.Join(holders[h], ", ")))
	}
	return strings.Join(parts, "; ")
}

// splitErrorLines turns an error built with errors.Join into one string per line.
func splitErrorLines(err error) []string {
	var lines []string
	for _, line := range strings.Split(err.Error(), "\n") {
		if line = strings.TrimSpace(line); line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
