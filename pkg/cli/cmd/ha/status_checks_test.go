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
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	operatork8s "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/kubecli"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	mclcontroller "github.com/gke-labs/multicluster-leader-election/controllers"
	mclstorage "github.com/gke-labs/multicluster-leader-election/pkg/storage"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	crfake "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

const (
	testLeaseName      = "kcc-multicluster-lease"
	testLeaseNamespace = "cnrm-system"
	testStatusContextC = "gke_proj-c_europe-west1_cluster-c"
	// Election identities of the test clusters, matching their GKE contexts as `ha setup` sets them.
	testIdentityA = "projects/proj-a/locations/us-central1/clusters/cluster-a"
	testIdentityB = "projects/proj-b/locations/us-east1/clusters/cluster-b"
	testIdentityC = "projects/proj-c/locations/europe-west1/clusters/cluster-c"
	testRenewTime = "2026-11-10T18:05:56Z"
)

// testNow is 4s after testRenewTime, so the test leases are live.
var testNow = time.Date(2026, 11, 10, 18, 6, 0, 0, time.UTC)

// haFixture holds the HA objects of one fake cluster.
type haFixture struct {
	context   string
	cc        *operatorv1beta1.ConfigConnector
	lease     *mclv1alpha1.MultiClusterLease
	syncers   []*syncerv1alpha1.KRMSyncer
	managers  []*appsv1.StatefulSet
	extraObjs []runtime.Object
	// unreachable makes every read from the cluster fail.
	unreachable bool
}

type haFixtures []*haFixture

func ptrInt32(v int32) *int32 { return &v }

func newTestManagerSTS(namespace string) *appsv1.StatefulSet {
	return &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cnrm-controller-manager-system",
			Namespace: namespace,
			Labels:    map[string]string{operatork8s.KCCSystemComponentLabel: operatork8s.KCCControllerManagerComponent},
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: ptrInt32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{operatork8s.KCCSystemComponentLabel: operatork8s.KCCControllerManagerComponent},
			},
		},
		Status: appsv1.StatefulSetStatus{
			Replicas:      1,
			ReadyReplicas: 1,
		},
	}
}

func healthyWorkloadObjects(leaseNamespace string) []runtime.Object {
	role := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      managerHARoleName,
			Namespace: leaseNamespace,
			Labels:    map[string]string{manifestManagedByLabelKey: manifestManagedByLabelValue},
		},
	}
	managerPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cnrm-controller-manager-system-0",
			Namespace: leaseNamespace,
			Labels:    map[string]string{operatork8s.KCCSystemComponentLabel: operatork8s.KCCControllerManagerComponent},
		},
		Status: corev1.PodStatus{
			ContainerStatuses: []corev1.ContainerStatus{
				{
					Name: "manager",
					State: corev1.ContainerState{
						Running: &corev1.ContainerStateRunning{StartedAt: metav1.NewTime(testNow.Add(-10 * time.Minute))},
					},
					RestartCount: 0,
				},
			},
		},
	}
	mclDeploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mclDeploymentName,
			Namespace: manifestSystemNamespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptrInt32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": mclDeploymentName},
			},
		},
		Status: appsv1.DeploymentStatus{
			Replicas:          1,
			AvailableReplicas: 1,
			UpdatedReplicas:   1,
		},
	}
	mclPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mclDeploymentName + "-pod",
			Namespace: manifestSystemNamespace,
			Labels:    map[string]string{"app": mclDeploymentName},
		},
		Status: corev1.PodStatus{
			ContainerStatuses: []corev1.ContainerStatus{
				{
					Name: "manager",
					State: corev1.ContainerState{
						Running: &corev1.ContainerStateRunning{StartedAt: metav1.NewTime(testNow.Add(-10 * time.Minute))},
					},
					RestartCount: 0,
				},
			},
		},
	}
	syncerDeploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      syncerDeploymentName,
			Namespace: manifestSystemNamespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptrInt32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": syncerDeploymentName},
			},
		},
		Status: appsv1.DeploymentStatus{
			Replicas:          1,
			AvailableReplicas: 1,
			UpdatedReplicas:   1,
		},
	}
	syncerPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      syncerDeploymentName + "-pod",
			Namespace: manifestSystemNamespace,
			Labels:    map[string]string{"app": syncerDeploymentName},
		},
		Status: corev1.PodStatus{
			ContainerStatuses: []corev1.ContainerStatus{
				{
					Name: "manager",
					State: corev1.ContainerState{
						Running: &corev1.ContainerStateRunning{StartedAt: metav1.NewTime(testNow.Add(-10 * time.Minute))},
					},
					RestartCount: 0,
				},
			},
		},
	}
	return []runtime.Object{role, managerPod, mclDeploy, mclPod, syncerDeploy, syncerPod}
}

// healthyHAFixtures returns a healthy pair: cluster-a holds the lease and its syncer is
// suspended; cluster-b is a standby whose syncer is syncing from cluster-a.
func healthyHAFixtures() haFixtures {
	return haFixtures{
		newLeaderFixture(testStatusContextA),
		newStandbyFixture(testStatusContextB, testIdentityA),
	}
}

// newLeaderFixture returns a cluster that holds the lease, with a suspended syncer and healthy
// HA workloads.
func newLeaderFixture(context string) *haFixture {
	identity := contextIdentity(context)
	sts := newReadyTestManagerStatefulSet(testLeaseNamespace, "cnrm-controller-manager", "")
	return &haFixture{
		context: context,
		cc:      newHAConfigConnector(context),
		lease:   newTestLease(identity, testRenewTime),
		syncers: []*syncerv1alpha1.KRMSyncer{
			newTestSyncer(testLeaseName, identity, true, syncerv1alpha1.ReasonSuspended, "Syncer is suspended"),
		},
		managers:  []*appsv1.StatefulSet{sts},
		extraObjs: healthyWorkloadObjects(testLeaseNamespace),
	}
}

// newStandbyFixture returns a standby cluster that sees leader holding the lease, with a syncer
// that is syncing from it and healthy HA workloads.
func newStandbyFixture(context, leader string) *haFixture {
	sts := newReadyTestManagerStatefulSet(testLeaseNamespace, "cnrm-controller-manager", "")
	return &haFixture{
		context: context,
		cc:      newHAConfigConnector(context),
		lease:   newTestLease(leader, testRenewTime),
		syncers: []*syncerv1alpha1.KRMSyncer{
			newTestSyncer(testLeaseName, leader, false, syncerv1alpha1.ReasonSyncing, "Syncing"),
		},
		managers:  []*appsv1.StatefulSet{sts},
		extraObjs: healthyWorkloadObjects(testLeaseNamespace),
	}
}

// newHAConfigConnector returns a ConfigConnector with multi-cluster HA configured, with the
// gkeCluster that `ha setup` derives from the GKE context.
func newHAConfigConnector(context string) *operatorv1beta1.ConfigConnector {
	cc := newConfigConnectorWithMCL(testLeaseName, testLeaseNamespace, "cluster")
	cc.Spec.Mode = "cluster"
	gke, _ := parseGKEContext(context)
	gkeCluster := gke.gkeCluster()
	cc.Spec.Experiments.MultiClusterLease.GKECluster = &gkeCluster
	return cc
}

func contextIdentity(context string) string {
	gke, _ := parseGKEContext(context)
	gkeCluster := gke.gkeCluster()
	return gkeCluster.ResourceName()
}

// unconfiguredFixture returns a cluster without multi-cluster HA.
func unconfiguredFixture(context string) *haFixture {
	return &haFixture{context: context, cc: newConfigConnectorWithoutMCL()}
}

func (fs haFixtures) handles(t *testing.T) []*clusterHandle {
	t.Helper()
	handles := make([]*clusterHandle, 0, len(fs))
	for _, f := range fs {
		objs := []runtime.Object{f.cc}
		if f.lease != nil {
			objs = append(objs, f.lease)
		}
		for _, s := range f.syncers {
			objs = append(objs, s)
		}
		for _, m := range f.managers {
			objs = append(objs, m)
		}
		objs = append(objs, f.extraObjs...)
		h := newFakeClusterHandle(t, f.context, objs...)
		if f.unreachable {
			h.client = &kubecli.Client{Client: &failingGetClient{Client: h.client.Client, err: errors.New("dial tcp 10.0.0.1:443: i/o timeout")}}
		}
		handles = append(handles, h)
	}
	return handles
}

func (fs haFixtures) handlesWithPendingDeployments(t *testing.T) []*clusterHandle {
	t.Helper()
	handles := make([]*clusterHandle, 0, len(fs))
	for _, f := range fs {
		objs := []runtime.Object{f.cc}
		if f.lease != nil {
			objs = append(objs, f.lease)
		}
		for _, s := range f.syncers {
			objs = append(objs, s)
		}
		for _, m := range f.managers {
			objs = append(objs, m)
		}
		objs = append(objs, f.extraObjs...)
		handles = append(handles, newFakeClusterHandleWithPendingDeployments(t, f.context, objs...))
	}
	return handles
}

func newReadyTestManagerStatefulSet(namespace, name, scopedNamespace string) *appsv1.StatefulSet {
	sts := newManagerStatefulSet(namespace, name, scopedNamespace)
	sts.Generation = 1
	sts.Spec.Selector = &metav1.LabelSelector{
		MatchLabels: map[string]string{operatork8s.KCCSystemComponentLabel: operatork8s.KCCControllerManagerComponent},
	}
	if scopedNamespace != "" {
		sts.Spec.Selector.MatchLabels[scopedNamespaceLabel] = scopedNamespace
	}
	sts.Status = appsv1.StatefulSetStatus{
		ObservedGeneration: 1,
		Replicas:           1,
		ReadyReplicas:      1,
	}
	return sts
}

func newTestLease(holder, renewTime string) *mclv1alpha1.MultiClusterLease {
	lease := &mclv1alpha1.MultiClusterLease{
		ObjectMeta: metav1.ObjectMeta{Name: testLeaseName, Namespace: testLeaseNamespace},
	}
	if holder != "" {
		lease.Status.GlobalHolderIdentity = &holder
		lease.Status.GlobalRenewTime = &renewTime
	}
	return lease
}

// newTestSyncer returns a KRMSyncer that syncs from remote (a GKE cluster resource name), with a
// Ready condition with the given reason. Ready is True for Syncing and Suspended.
func newTestSyncer(name, remote string, suspend bool, reason, message string) *syncerv1alpha1.KRMSyncer {
	s := &syncerv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: testLeaseNamespace},
		Spec:       syncerv1alpha1.KRMSyncerSpec{Suspend: suspend, Mode: syncerv1alpha1.ModePull},
	}
	if gke, err := operatorv1beta1.ParseGKEClusterResourceName(remote); err == nil {
		s.Spec.Remote = &syncerv1alpha1.RemoteConfig{GKECluster: &syncerv1alpha1.GKECluster{Project: gke.Project, Location: gke.Location, Name: gke.Name}}
	}
	if reason != "" {
		status := metav1.ConditionFalse
		if reason == syncerv1alpha1.ReasonSyncing || reason == syncerv1alpha1.ReasonSuspended {
			status = metav1.ConditionTrue
		}
		s.Status.Conditions = []metav1.Condition{{Type: syncerv1alpha1.ConditionReady, Status: status, Reason: reason, Message: message}}
	}
	return s
}

func runHAStatus(t *testing.T, fs haFixtures) (*statusReport, string) {
	t.Helper()
	report := executeStatus(context.Background(), fs.handles(t), testNow)
	var out bytes.Buffer
	report.render(&out)
	return report, out.String()
}

func TestHAStatusHealthy(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(haFixtures)
	}{
		{
			name:   "healthy pair",
			mutate: func(haFixtures) {},
		},
		{
			name: "standby candidate lease is stale in multicluster-leader-election",
			mutate: func(fs haFixtures) {
				fs[1].lease.Status.Conditions = []metav1.Condition{{
					Type:    string(mclv1alpha1.ConditionTypeBackendHealthy),
					Status:  metav1.ConditionFalse,
					Reason:  "BackendError",
					Message: "failed to communicate with backend: candidate lease is stale",
				}}
			},
		},
		{
			name:   "leader without a KRMSyncer",
			mutate: func(fs haFixtures) { fs[0].syncers = nil },
		},
		{
			name: "replication disabled, so no KRMSyncers",
			mutate: func(fs haFixtures) {
				for _, f := range fs {
					f.cc.Spec.Experiments.MultiClusterLease.ResourceReplicationMode = replicationModeDisabled
					f.syncers = nil
				}
			},
		},
		{
			name:   "leader with Config Connector paused",
			mutate: func(fs haFixtures) { fs[0].cc.Spec.Actuation = operatorv1beta1.Paused },
		},
		{
			name: "namespaced mode, one KRMSyncer per namespace",
			mutate: func(fs haFixtures) {
				fs[0].syncers = []*syncerv1alpha1.KRMSyncer{
					newTestSyncer(testLeaseName+"-team-a", testIdentityA, true, syncerv1alpha1.ReasonSuspended, ""),
					newTestSyncer(testLeaseName+"-team-b", testIdentityA, true, syncerv1alpha1.ReasonSuspended, ""),
				}
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{
					newTestSyncer(testLeaseName+"-team-a", testIdentityA, false, syncerv1alpha1.ReasonSyncing, ""),
					newTestSyncer(testLeaseName+"-team-b", testIdentityA, false, syncerv1alpha1.ReasonSyncing, ""),
				}
			},
		},
		{
			name: "namespaced mode with a running manager for team-a and team-b on both clusters",
			mutate: func(fs haFixtures) {
				for _, f := range fs {
					f.cc.Spec.Mode = "namespaced"
					f.managers = []*appsv1.StatefulSet{
						newReadyTestManagerStatefulSet(testLeaseNamespace, "cnrm-controller-manager-team-a", "team-a"),
						newReadyTestManagerStatefulSet(testLeaseNamespace, "cnrm-controller-manager-team-b", "team-b"),
					}
				}
				fs[0].syncers = []*syncerv1alpha1.KRMSyncer{
					newTestSyncer(testLeaseName+"-team-a", testIdentityA, true, syncerv1alpha1.ReasonSuspended, ""),
					newTestSyncer(testLeaseName+"-team-b", testIdentityA, true, syncerv1alpha1.ReasonSuspended, ""),
				}
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{
					newTestSyncer(testLeaseName+"-team-a", testIdentityA, false, syncerv1alpha1.ReasonSyncing, ""),
					newTestSyncer(testLeaseName+"-team-b", testIdentityA, false, syncerv1alpha1.ReasonSyncing, ""),
				}
			},
		},
		{
			name: "unrelated KRMSyncers are ignored",
			mutate: func(fs haFixtures) {
				fs[1].syncers = append(fs[1].syncers, newTestSyncer("other-syncer", "", true, "", ""))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fs := healthyHAFixtures()
			tc.mutate(fs)

			report, out := runHAStatus(t, fs)

			if report.state != haStateHealthy {
				t.Fatalf("got state %q, want %q; output:\n%s", report.state, haStateHealthy, out)
			}
		})
	}
}

func threeClusterHAFixtures() haFixtures {
	fs := healthyHAFixtures()
	fs = append(fs, newStandbyFixture("gke_proj-c_us-central1_cluster-c", testIdentityA))
	return fs
}

func TestHAStatusProblems(t *testing.T) {
	unhealthy := metav1.Condition{
		Type:    string(mclv1alpha1.ConditionTypeBackendHealthy),
		Status:  metav1.ConditionFalse,
		Reason:  "BackendError",
		Message: "failed to communicate with backend: storage: bucket doesn't exist",
	}

	tests := []struct {
		name         string
		fixtures     func() haFixtures
		mutate       func(haFixtures)
		wantProblems []string
	}{
		// Election configuration.
		{
			name:   "mismatched leaseName",
			mutate: func(fs haFixtures) { fs[1].cc.Spec.Experiments.MultiClusterLease.LeaseName = "other-lease" },
			wantProblems: []string{
				`mismatched leaseName across clusters (gke_proj-a_us-central1_cluster-a="kcc-multicluster-lease", gke_proj-b_us-east1_cluster-b="other-lease")`,
			},
		},
		{
			name: "duplicate gkeCluster",
			mutate: func(fs haFixtures) {
				*fs[1].cc.Spec.Experiments.MultiClusterLease.GKECluster = *fs[0].cc.Spec.Experiments.MultiClusterLease.GKECluster
			},
			wantProblems: []string{
				`duplicate gkeCluster across clusters (identity "` + testIdentityA + `" shared by [gke_proj-a_us-central1_cluster-a, gke_proj-b_us-east1_cluster-b])`,
			},
		},
		{
			name:   "mismatched resourceReplicationMode",
			mutate: func(fs haFixtures) { fs[1].cc.Spec.Experiments.MultiClusterLease.ResourceReplicationMode = "Full" },
			wantProblems: []string{
				`mismatched resourceReplicationMode across clusters (gke_proj-a_us-central1_cluster-a="Status", gke_proj-b_us-east1_cluster-b="Full")`,
			},
		},
		// Lease.
		{
			name:   "MultiClusterLease missing",
			mutate: func(fs haFixtures) { fs[1].lease = nil },
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": MultiClusterLease cnrm-system/kcc-multicluster-lease not found; check that the multiclusterlease-controller-manager is running`,
			},
		},
		{
			name:   "lease backend unhealthy",
			mutate: func(fs haFixtures) { fs[0].lease.Status.Conditions = []metav1.Condition{unhealthy} },
			wantProblems: []string{
				`cluster "gke_proj-a_us-central1_cluster-a": the lease backend is unhealthy: failed to communicate with backend: storage: bucket doesn't exist`,
			},
		},
		{
			name: "standby lease backend unhealthy with real backend failure",
			mutate: func(fs haFixtures) {
				fs[1].lease.Status.Conditions = []metav1.Condition{{
					Type:    string(mclv1alpha1.ConditionTypeBackendHealthy),
					Status:  metav1.ConditionFalse,
					Reason:  "BackendError",
					Message: "failed to communicate with backend: failed to read lease: storage: bucket doesn't exist",
				}}
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": the lease backend is unhealthy: failed to communicate with backend: failed to read lease: storage: bucket doesn't exist`,
			},
		},
		{
			name: "stale candidate on leader is not ignored",
			mutate: func(fs haFixtures) {
				fs[0].lease.Status.Conditions = []metav1.Condition{{
					Type:    string(mclv1alpha1.ConditionTypeBackendHealthy),
					Status:  metav1.ConditionFalse,
					Reason:  "BackendError",
					Message: "failed to communicate with backend: candidate lease is stale",
				}}
			},
			wantProblems: []string{
				`cluster "gke_proj-a_us-central1_cluster-a": the lease backend is unhealthy: failed to communicate with backend: candidate lease is stale`,
			},
		},
		{
			name: "standby candidate lease is stale but lease expired",
			mutate: func(fs haFixtures) {
				fs[1].lease = newTestLease(testIdentityA, "2026-11-10T18:05:00Z")
				fs[1].lease.Status.Conditions = []metav1.Condition{{
					Type:    string(mclv1alpha1.ConditionTypeBackendHealthy),
					Status:  metav1.ConditionFalse,
					Reason:  "BackendError",
					Message: "failed to communicate with backend: candidate lease is stale",
				}}
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": the lease held by "` + testIdentityA + `" has expired (last renewed at 2026-11-10T18:05:00Z, lease duration 15s)`,
			},
		},
		{
			name: "stale candidate with no lease holder",
			mutate: func(fs haFixtures) {
				fs[0].lease = newTestLease("", "")
				fs[1].lease = newTestLease("", "")
				fs[1].lease.Status.Conditions = []metav1.Condition{{
					Type:    string(mclv1alpha1.ConditionTypeBackendHealthy),
					Status:  metav1.ConditionFalse,
					Reason:  "BackendError",
					Message: "failed to communicate with backend: candidate lease is stale",
				}}
			},
			wantProblems: []string{
				`cluster "gke_proj-a_us-central1_cluster-a": no cluster holds the lease cnrm-system/kcc-multicluster-lease`,
				`cluster "gke_proj-b_us-east1_cluster-b": the lease backend is unhealthy: failed to communicate with backend: candidate lease is stale`,
			},
		},
		{
			name: "no lease holder",
			mutate: func(fs haFixtures) {
				fs[0].lease = newTestLease("", "")
				fs[1].lease = newTestLease("", "")
			},
			wantProblems: []string{
				`cluster "gke_proj-a_us-central1_cluster-a": no cluster holds the lease cnrm-system/kcc-multicluster-lease`,
				`cluster "gke_proj-b_us-east1_cluster-b": no cluster holds the lease cnrm-system/kcc-multicluster-lease`,
			},
		},
		{
			name:   "lease expired",
			mutate: func(fs haFixtures) { fs[1].lease = newTestLease(testIdentityA, "2026-11-10T18:05:00Z") },
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": the lease held by "` + testIdentityA + `" has expired (last renewed at 2026-11-10T18:05:00Z, lease duration 15s)`,
			},
		},
		{
			name:   "clusters disagree on the lease holder",
			mutate: func(fs haFixtures) { fs[1].lease = newTestLease(testIdentityB, testRenewTime) },
			wantProblems: []string{
				`clusters disagree on the lease holder ("` + testIdentityA + `" reported by [gke_proj-a_us-central1_cluster-a]; "` + testIdentityB + `" reported by [gke_proj-b_us-east1_cluster-b])`,
			},
		},
		{
			name: "lease held by a cluster that wasn't given",
			mutate: func(fs haFixtures) {
				other := "projects/other/locations/europe-west1/clusters/cluster-c"
				for _, f := range fs {
					f.lease = newTestLease(other, testRenewTime)
					f.syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, other, false, syncerv1alpha1.ReasonSyncing, "")}
				}
			},
			wantProblems: []string{
				`the lease is held by "projects/other/locations/europe-west1/clusters/cluster-c", which is not one of the clusters in --contexts`,
			},
		},
		{
			name: "lease held by cluster not in --contexts and standby lease expired",
			mutate: func(fs haFixtures) {
				other := "projects/other/locations/europe-west1/clusters/cluster-c"
				fs[0].lease = newTestLease(other, testRenewTime)
				fs[1].lease = newTestLease(testIdentityA, "2026-11-10T18:05:00Z")
				// The syncers are checked against the agreed holder, even outside --contexts.
				for _, f := range fs {
					f.syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, other, false, syncerv1alpha1.ReasonSyncing, "")}
				}
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": the lease held by "` + testIdentityA + `" has expired (last renewed at 2026-11-10T18:05:00Z, lease duration 15s)`,
				`the lease is held by "projects/other/locations/europe-west1/clusters/cluster-c", which is not one of the clusters in --contexts`,
			},
		},
		// KRMSyncers.
		{
			name: "leader KRMSyncer not suspended",
			mutate: func(fs haFixtures) {
				fs[0].syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, testIdentityB, false, syncerv1alpha1.ReasonSyncing, "")}
			},
			wantProblems: []string{
				`cluster "gke_proj-a_us-central1_cluster-a" (leader): KRMSyncer cnrm-system/kcc-multicluster-lease is not suspended`,
			},
		},
		{
			name:   "standby without a KRMSyncer",
			mutate: func(fs haFixtures) { fs[1].syncers = nil },
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": no KRMSyncer for lease "kcc-multicluster-lease" found in namespace "cnrm-system"`,
			},
		},
		{
			name: "standby KRMSyncer suspended",
			mutate: func(fs haFixtures) {
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, testIdentityA, true, syncerv1alpha1.ReasonSuspended, "")}
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": KRMSyncer cnrm-system/kcc-multicluster-lease is suspended, so it is not syncing from the leader`,
			},
		},
		{
			name: "standby KRMSyncer syncs from the wrong cluster",
			mutate: func(fs haFixtures) {
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, testIdentityB, false, syncerv1alpha1.ReasonSyncing, "")}
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": KRMSyncer cnrm-system/kcc-multicluster-lease syncs from "` + testIdentityB + `", but the lease is held by "` + testIdentityA + `"`,
			},
		},
		{
			name: "standby KRMSyncer without a Ready condition",
			mutate: func(fs haFixtures) {
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, testIdentityA, false, "", "")}
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": KRMSyncer cnrm-system/kcc-multicluster-lease has not reported a Ready condition; check that the syncer-controller-manager is running`,
			},
		},
		{
			name: "standby lease expired and standby KRMSyncer suspended",
			mutate: func(fs haFixtures) {
				fs[1].lease = newTestLease(testIdentityA, "2026-11-10T18:05:00Z")
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, testIdentityA, true, syncerv1alpha1.ReasonSuspended, "")}
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": the lease held by "` + testIdentityA + `" has expired (last renewed at 2026-11-10T18:05:00Z, lease duration 15s)`,
				`cluster "gke_proj-b_us-east1_cluster-b": KRMSyncer cnrm-system/kcc-multicluster-lease is suspended, so it is not syncing from the leader`,
			},
		},
		{
			name: "standby lease backend unhealthy and standby KRMSyncer SyncFailed",
			mutate: func(fs haFixtures) {
				fs[1].lease.Status.Conditions = []metav1.Condition{{
					Type:    string(mclv1alpha1.ConditionTypeBackendHealthy),
					Status:  metav1.ConditionFalse,
					Reason:  "BackendError",
					Message: "failed to communicate with backend: failed to read lease: Error 403",
				}}
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, testIdentityA, false, syncerv1alpha1.ReasonSyncFailed, "failed to fetch objects")}
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": the lease backend is unhealthy: failed to communicate with backend: failed to read lease: Error 403`,
				`cluster "gke_proj-b_us-east1_cluster-b": KRMSyncer cnrm-system/kcc-multicluster-lease is failing to sync (SyncFailed): failed to fetch objects`,
			},
		},
		{
			name: "leader missing MultiClusterLease and leader KRMSyncer not suspended",
			mutate: func(fs haFixtures) {
				fs[0].lease = nil
				fs[0].syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, testIdentityB, false, syncerv1alpha1.ReasonSyncing, "")}
			},
			wantProblems: []string{
				`cluster "gke_proj-a_us-central1_cluster-a": MultiClusterLease cnrm-system/kcc-multicluster-lease not found; check that the multiclusterlease-controller-manager is running`,
				`cluster "gke_proj-a_us-central1_cluster-a" (leader): KRMSyncer cnrm-system/kcc-multicluster-lease is not suspended`,
			},
		},
		{
			name:     "three clusters: c lease expired and c KRMSyncer syncs from b instead of leader a",
			fixtures: threeClusterHAFixtures,
			mutate: func(fs haFixtures) {
				fs[2].lease = newTestLease(testIdentityA, "2026-11-10T18:05:00Z")
				fs[2].syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, testIdentityB, false, syncerv1alpha1.ReasonSyncing, "")}
			},
			wantProblems: []string{
				`cluster "gke_proj-c_us-central1_cluster-c": the lease held by "` + testIdentityA + `" has expired (last renewed at 2026-11-10T18:05:00Z, lease duration 15s)`,
				`cluster "gke_proj-c_us-central1_cluster-c": KRMSyncer cnrm-system/kcc-multicluster-lease syncs from "` + testIdentityB + `", but the lease is held by "` + testIdentityA + `"`,
			},
		},
		{
			name:     "three clusters: c lease expired, a and b disagree on holder",
			fixtures: threeClusterHAFixtures,
			mutate: func(fs haFixtures) {
				fs[1].lease = newTestLease(testIdentityB, testRenewTime)
				fs[2].lease = newTestLease(testIdentityA, "2026-11-10T18:05:00Z")
				fs[2].syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, testIdentityB, false, syncerv1alpha1.ReasonSyncing, "")}
			},
			wantProblems: []string{
				`cluster "gke_proj-c_us-central1_cluster-c": the lease held by "` + testIdentityA + `" has expired (last renewed at 2026-11-10T18:05:00Z, lease duration 15s)`,
				`clusters disagree on the lease holder ("` + testIdentityA + `" reported by [gke_proj-a_us-central1_cluster-a]; "` + testIdentityB + `" reported by [gke_proj-b_us-east1_cluster-b])`,
			},
		},
		{
			name: "standby KRMSyncer status is out of date",
			mutate: func(fs haFixtures) {
				s := newTestSyncer(testLeaseName, testIdentityA, false, syncerv1alpha1.ReasonSyncing, "")
				s.Generation = 2
				s.Status.ObservedGeneration = 1
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{s}
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": KRMSyncer cnrm-system/kcc-multicluster-lease has not processed its latest spec yet; check that the syncer-controller-manager is running`,
			},
		},
		{
			name: "standby can't reach the leader",
			mutate: func(fs haFixtures) {
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, testIdentityA, false, syncerv1alpha1.ReasonRemoteUnreachable,
					"remote cluster "+testIdentityA+": Unauthorized")}
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": KRMSyncer cnrm-system/kcc-multicluster-lease is not syncing (RemoteUnreachable): remote cluster ` + testIdentityA + `: Unauthorized`,
			},
		},
		{
			name: "standby KRMSyncer setup failed",
			mutate: func(fs haFixtures) {
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, testIdentityA, false, syncerv1alpha1.ReasonSetupFailed, "discovery failed")}
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": KRMSyncer cnrm-system/kcc-multicluster-lease is not syncing (SetupFailed): discovery failed`,
			},
		},
		{
			name: "standby KRMSyncer failing to sync objects",
			mutate: func(fs haFixtures) {
				s := newTestSyncer(testLeaseName, testIdentityA, false, syncerv1alpha1.ReasonSyncFailed,
					"apply failed (Forbidden): StorageBucket. See the controller logs for details.")
				s.Status.FailedObjects = 3
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{s}
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": KRMSyncer cnrm-system/kcc-multicluster-lease is failing to sync 3 objects (SyncFailed): apply failed (Forbidden): StorageBucket. See the controller logs for details.`,
			},
		},
		{
			name: "standby KRMSyncer failing to sync kinds",
			mutate: func(fs haFixtures) {
				s := newTestSyncer(testLeaseName, testIdentityA, false, syncerv1alpha1.ReasonSyncFailed,
					"watch failed (Forbidden): APIGatewayAPI.apigateway.cnrm.cloud.google.com. See the controller logs for details.")
				s.Status.FailedObjects = 0
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{s}
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": KRMSyncer cnrm-system/kcc-multicluster-lease is failing to sync (SyncFailed): watch failed (Forbidden): APIGatewayAPI.apigateway.cnrm.cloud.google.com. See the controller logs for details.`,
			},
		},
		// Managers.
		{
			name: "namespaced mode, leader manager of team-a scaled to 0",
			mutate: func(fs haFixtures) {
				for _, f := range fs {
					f.cc.Spec.Mode = "namespaced"
					f.managers = []*appsv1.StatefulSet{
						newReadyTestManagerStatefulSet(testLeaseNamespace, "cnrm-controller-manager-team-a", "team-a"),
						newReadyTestManagerStatefulSet(testLeaseNamespace, "cnrm-controller-manager-team-b", "team-b"),
					}
				}
				fs[0].syncers = []*syncerv1alpha1.KRMSyncer{
					newTestSyncer(testLeaseName+"-team-a", testIdentityA, true, syncerv1alpha1.ReasonSuspended, ""),
					newTestSyncer(testLeaseName+"-team-b", testIdentityA, true, syncerv1alpha1.ReasonSuspended, ""),
				}
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{
					newTestSyncer(testLeaseName+"-team-a", testIdentityA, false, syncerv1alpha1.ReasonSyncing, ""),
					newTestSyncer(testLeaseName+"-team-b", testIdentityA, false, syncerv1alpha1.ReasonSyncing, ""),
				}
				zero := int32(0)
				fs[0].managers[0].Spec.Replicas = &zero
				fs[0].managers[0].Status.Replicas = 0
				fs[0].managers[0].Status.ReadyReplicas = 0
			},
			wantProblems: []string{
				`cluster "gke_proj-a_us-central1_cluster-a" (leader): the manager of namespace "team-a" is not running (StatefulSet cnrm-system/cnrm-controller-manager-team-a is scaled to 0); no cluster reconciles that namespace. The cluster keeps the lease while another of its managers runs: start the manager, or move leadership to a cluster where it runs with config-connector ha failover`,
			},
		},
		{
			name: "namespaced mode, standby manager of team-a with 0 of 1 pods ready",
			mutate: func(fs haFixtures) {
				for _, f := range fs {
					f.cc.Spec.Mode = "namespaced"
					f.managers = []*appsv1.StatefulSet{
						newReadyTestManagerStatefulSet(testLeaseNamespace, "cnrm-controller-manager-team-a", "team-a"),
						newReadyTestManagerStatefulSet(testLeaseNamespace, "cnrm-controller-manager-team-b", "team-b"),
					}
				}
				fs[0].syncers = []*syncerv1alpha1.KRMSyncer{
					newTestSyncer(testLeaseName+"-team-a", testIdentityA, true, syncerv1alpha1.ReasonSuspended, ""),
					newTestSyncer(testLeaseName+"-team-b", testIdentityA, true, syncerv1alpha1.ReasonSuspended, ""),
				}
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{
					newTestSyncer(testLeaseName+"-team-a", testIdentityA, false, syncerv1alpha1.ReasonSyncing, ""),
					newTestSyncer(testLeaseName+"-team-b", testIdentityA, false, syncerv1alpha1.ReasonSyncing, ""),
				}
				fs[1].managers[0].Status.ReadyReplicas = 0
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": the manager of namespace "team-a" is not running (StatefulSet cnrm-system/cnrm-controller-manager-team-a has 0 of 1 pods ready); this cluster would not reconcile that namespace as the leader`,
			},
		},
		{
			name: "cluster mode, standby manager scaled to 0",
			mutate: func(fs haFixtures) {
				zero := int32(0)
				fs[1].managers[0].Spec.Replicas = &zero
				fs[1].managers[0].Status.Replicas = 0
				fs[1].managers[0].Status.ReadyReplicas = 0
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Config Connector's manager is not running (StatefulSet cnrm-system/cnrm-controller-manager is scaled to 0); this cluster cannot take over`,
			},
		},
		{
			name: "manager that ha failover stopped",
			mutate: func(fs haFixtures) {
				zero := int32(0)
				fs[1].managers[0].Spec.Replicas = &zero
				fs[1].managers[0].Status.Replicas = 0
				fs[1].managers[0].Status.ReadyReplicas = 0
				fs[1].managers[0].Annotations[failoverReplicasAnnotation] = "1"
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Config Connector's manager is not running (StatefulSet cnrm-system/cnrm-controller-manager was stopped by ha failover; running ha failover again starts it); this cluster cannot take over`,
			},
		},
		{
			name: "manager whose latest spec the StatefulSet controller has not processed",
			mutate: func(fs haFixtures) {
				fs[1].managers[0].Generation = 2
				fs[1].managers[0].Status.ObservedGeneration = 1
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Config Connector's manager is not running (the StatefulSet controller has not processed the latest spec of StatefulSet cnrm-system/cnrm-controller-manager yet); this cluster cannot take over`,
			},
		},
		{
			name: "cluster with no manager StatefulSet",
			mutate: func(fs haFixtures) {
				fs[1].managers = nil
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": no StatefulSet of Config Connector's manager found; this cluster cannot take over`,
			},
		},
		{
			name: "an expired lease and a manager that is down",
			mutate: func(fs haFixtures) {
				fs[1].lease = newTestLease(testIdentityA, "2026-11-10T18:05:00Z")
				zero := int32(0)
				fs[0].managers[0].Spec.Replicas = &zero
				fs[0].managers[0].Status.Replicas = 0
				fs[0].managers[0].Status.ReadyReplicas = 0
			},
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": the lease held by "` + testIdentityA + `" has expired (last renewed at 2026-11-10T18:05:00Z, lease duration 15s)`,
				`cluster "gke_proj-a_us-central1_cluster-a" (leader): Config Connector's manager is not running (StatefulSet cnrm-system/cnrm-controller-manager is scaled to 0); nothing reconciles until it runs again or another cluster takes the lease`,
			},
		},
		{
			name: "all leases expired and a manager that is down",
			mutate: func(fs haFixtures) {
				fs[0].lease = newTestLease(testIdentityA, "2026-11-10T18:05:00Z")
				fs[1].lease = newTestLease(testIdentityA, "2026-11-10T18:05:00Z")
				zero := int32(0)
				fs[0].managers[0].Spec.Replicas = &zero
				fs[0].managers[0].Status.Replicas = 0
				fs[0].managers[0].Status.ReadyReplicas = 0
			},
			wantProblems: []string{
				`cluster "gke_proj-a_us-central1_cluster-a": the lease held by "` + testIdentityA + `" has expired (last renewed at 2026-11-10T18:05:00Z, lease duration 15s)`,
				`cluster "gke_proj-b_us-east1_cluster-b": the lease held by "` + testIdentityA + `" has expired (last renewed at 2026-11-10T18:05:00Z, lease duration 15s)`,
				`cluster "gke_proj-a_us-central1_cluster-a": Config Connector's manager is not running (StatefulSet cnrm-system/cnrm-controller-manager is scaled to 0); this cluster cannot take over`,
			},
		},
		{
			name: "leader with no manager StatefulSet",
			mutate: func(fs haFixtures) {
				fs[0].managers = nil
			},
			wantProblems: []string{
				`cluster "gke_proj-a_us-central1_cluster-a" (leader): no StatefulSet of Config Connector's manager found; nothing reconciles until it runs again or another cluster takes the lease`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fsFn := tc.fixtures
			if fsFn == nil {
				fsFn = healthyHAFixtures
			}
			fs := fsFn()
			if tc.mutate != nil {
				tc.mutate(fs)
			}

			report, out := runHAStatus(t, fs)

			if report.state != haStateError {
				t.Fatalf("got state %q, want %q; output:\n%s", report.state, haStateError, out)
			}
			if len(report.problems) != len(tc.wantProblems) {
				t.Fatalf("got %d problems, want %d; output:\n%s", len(report.problems), len(tc.wantProblems), out)
			}
			for _, want := range tc.wantProblems {
				if !strings.Contains(out, "  - "+want) {
					t.Fatalf("expected output to contain problem %q, got:\n%s", want, out)
				}
			}
		})
	}
}

func TestHAStatusRoles(t *testing.T) {
	notConfiguredB := `cluster "gke_proj-b_us-east1_cluster-b" does not have spec.experiments.multiClusterLease configured`
	notConfiguredC := `cluster "gke_proj-c_europe-west1_cluster-c" does not have spec.experiments.multiClusterLease configured`
	unreachableA := `fetching ConfigConnector on cluster "gke_proj-a_us-central1_cluster-a": dial tcp 10.0.0.1:443: i/o timeout`
	unreachableC := `fetching ConfigConnector on cluster "gke_proj-c_europe-west1_cluster-c": dial tcp 10.0.0.1:443: i/o timeout`
	leaseA := "Lease: cnrm-system/kcc-multicluster-lease, held by gke_proj-a_us-central1_cluster-a (renewed 4s ago)"

	tests := []struct {
		name string
		// fixtures returns the clusters; healthyHAFixtures if nil.
		fixtures  func() haFixtures
		mutate    func(haFixtures)
		wantState haState
		wantRoles []clusterRole
		// wantLease is the Lease line; empty means no Lease line.
		wantLease string
		// wantProblems are substrings of the expected problems, one per problem.
		wantProblems []string
	}{
		// Every cluster configured.
		{
			name:      "cluster-a leads",
			wantState: haStateHealthy,
			wantRoles: []clusterRole{roleLeader, roleStandby},
			wantLease: leaseA,
		},
		{
			name: "cluster-b leads",
			fixtures: func() haFixtures {
				return haFixtures{newStandbyFixture(testStatusContextA, testIdentityB), newLeaderFixture(testStatusContextB)}
			},
			wantState: haStateHealthy,
			wantRoles: []clusterRole{roleStandby, roleLeader},
			wantLease: "Lease: cnrm-system/kcc-multicluster-lease, held by gke_proj-b_us-east1_cluster-b (renewed 4s ago)",
		},
		{
			name: "leader known, standby not syncing",
			mutate: func(fs haFixtures) {
				fs[1].syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, testIdentityA, false, syncerv1alpha1.ReasonRemoteUnreachable, "Unauthorized")}
			},
			wantState:    haStateError,
			wantRoles:    []clusterRole{roleLeader, roleStandby},
			wantLease:    leaseA,
			wantProblems: []string{"is not syncing (RemoteUnreachable): Unauthorized"},
		},
		{
			name: "one lease missing, the other is live, so the leader is known",
			mutate: func(fs haFixtures) {
				fs[1].lease = nil
			},
			wantState:    haStateError,
			wantRoles:    []clusterRole{roleLeader, roleStandby},
			wantLease:    leaseA,
			wantProblems: []string{`cluster "gke_proj-b_us-east1_cluster-b": MultiClusterLease cnrm-system/kcc-multicluster-lease not found`},
		},
		{
			name: "all leases expired, so the leader is unknown",
			mutate: func(fs haFixtures) {
				fs[0].lease = newTestLease(testIdentityA, "2026-11-10T18:05:00Z")
				fs[1].lease = newTestLease(testIdentityA, "2026-11-10T18:05:00Z")
			},
			wantState:    haStateError,
			wantRoles:    []clusterRole{roleUnknown, roleUnknown},
			wantProblems: []string{"has expired", "has expired"},
		},
		{
			name:         "clusters disagree on the holder, so the leader is unknown",
			mutate:       func(fs haFixtures) { fs[1].lease = newTestLease(testIdentityB, testRenewTime) },
			wantState:    haStateError,
			wantRoles:    []clusterRole{roleUnknown, roleUnknown},
			wantProblems: []string{"clusters disagree on the lease holder"},
		},
		{
			name: "holder outside --contexts is named by its identity",
			mutate: func(fs haFixtures) {
				for _, f := range fs {
					f.lease = newTestLease(testIdentityC, testRenewTime)
					f.syncers = []*syncerv1alpha1.KRMSyncer{newTestSyncer(testLeaseName, testIdentityC, false, syncerv1alpha1.ReasonSyncing, "")}
				}
			},
			wantState:    haStateError,
			wantRoles:    []clusterRole{roleStandby, roleStandby},
			wantLease:    "Lease: cnrm-system/kcc-multicluster-lease, held by " + testIdentityC + " (renewed 4s ago)",
			wantProblems: []string{"which is not one of the clusters in --contexts"},
		},
		{
			name: "no renew time reported",
			mutate: func(fs haFixtures) {
				for _, f := range fs {
					f.lease.Status.GlobalRenewTime = nil
				}
			},
			wantState: haStateHealthy,
			wantRoles: []clusterRole{roleLeader, roleStandby},
			wantLease: "Lease: cnrm-system/kcc-multicluster-lease, held by gke_proj-a_us-central1_cluster-a\n",
		},
		// Some clusters not configured, or not readable.
		{
			name: "the leader can't be read, and the standby sees its live lease",
			mutate: func(fs haFixtures) {
				fs[0].unreachable = true
			},
			wantState:    haStateError,
			wantRoles:    []clusterRole{roleNone, roleStandby},
			wantLease:    leaseA,
			wantProblems: []string{unreachableA},
		},
		{
			name: "only cluster-a configured, and it holds the lease",
			fixtures: func() haFixtures {
				return haFixtures{newLeaderFixture(testStatusContextA), unconfiguredFixture(testStatusContextB)}
			},
			wantState:    haStateError,
			wantRoles:    []clusterRole{roleLeader, roleNone},
			wantLease:    leaseA,
			wantProblems: []string{notConfiguredB + ` on ConfigConnector "configconnector.core.cnrm.cloud.google.com"; run ` + "`config-connector ha setup`" + ` or configure spec.experiments.multiClusterLease; it reconciles independently of the HA clusters, so more than one cluster may change the same resources`},
		},
		{
			name: "only cluster-a configured, and cluster-b holds the lease",
			fixtures: func() haFixtures {
				return haFixtures{newStandbyFixture(testStatusContextA, testIdentityB), unconfiguredFixture(testStatusContextB)}
			},
			wantState:    haStateError,
			wantRoles:    []clusterRole{roleStandby, roleNone},
			wantLease:    "Lease: cnrm-system/kcc-multicluster-lease, held by gke_proj-b_us-east1_cluster-b (renewed 4s ago)",
			wantProblems: []string{notConfiguredB},
		},
		{
			name: "only cluster-a configured, and its lease expired",
			fixtures: func() haFixtures {
				return haFixtures{newLeaderFixture(testStatusContextA), unconfiguredFixture(testStatusContextB)}
			},
			mutate: func(fs haFixtures) {
				fs[0].lease = newTestLease(testIdentityA, "2026-11-10T18:05:00Z")
			},
			wantState: haStateError,
			wantRoles: []clusterRole{roleUnknown, roleNone},
			wantProblems: []string{
				notConfiguredB,
				`cluster "gke_proj-a_us-central1_cluster-a": the lease held by "` + testIdentityA + `" has expired`,
			},
		},
		{
			name: "only cluster-a configured, and its KRMSyncer isn't suspended",
			fixtures: func() haFixtures {
				return haFixtures{newLeaderFixture(testStatusContextA), unconfiguredFixture(testStatusContextB)}
			},
			mutate: func(fs haFixtures) {
				fs[0].syncers[0].Spec.Suspend = false
			},
			wantState: haStateError,
			wantRoles: []clusterRole{roleLeader, roleNone},
			wantLease: leaseA,
			wantProblems: []string{
				notConfiguredB,
				`cluster "gke_proj-a_us-central1_cluster-a" (leader): KRMSyncer cnrm-system/kcc-multicluster-lease is not suspended`,
			},
		},
		// Three clusters.
		{
			name: "three clusters, one can't be read",
			fixtures: func() haFixtures {
				c := unconfiguredFixture(testStatusContextC)
				c.unreachable = true
				return haFixtures{newLeaderFixture(testStatusContextA), newStandbyFixture(testStatusContextB, testIdentityA), c}
			},
			wantState:    haStateError,
			wantRoles:    []clusterRole{roleLeader, roleStandby, roleNone},
			wantLease:    leaseA,
			wantProblems: []string{unreachableC},
		},
		{
			name: "three clusters, the configured two use different leases",
			fixtures: func() haFixtures {
				b := newLeaderFixture(testStatusContextB)
				b.cc.Spec.Experiments.MultiClusterLease.LeaseName = "other-lease"
				return haFixtures{newLeaderFixture(testStatusContextA), b, unconfiguredFixture(testStatusContextC)}
			},
			wantState: haStateError,
			wantRoles: []clusterRole{roleUnknown, roleUnknown, roleNone},
			wantProblems: []string{
				notConfiguredC,
				`mismatched leaseName across clusters (gke_proj-a_us-central1_cluster-a="kcc-multicluster-lease", gke_proj-b_us-east1_cluster-b="other-lease")`,
			},
		},
		{
			name: "three clusters, the configured two disagree on the holder",
			fixtures: func() haFixtures {
				a := newStandbyFixture(testStatusContextA, testIdentityC)
				a.syncers[0].Status.Conditions[0].Status = metav1.ConditionFalse
				a.syncers[0].Status.Conditions[0].Reason = syncerv1alpha1.ReasonSyncFailed
				c := unconfiguredFixture(testStatusContextC)
				c.unreachable = true
				return haFixtures{a, newLeaderFixture(testStatusContextB), c}
			},
			wantState: haStateError,
			wantRoles: []clusterRole{roleUnknown, roleUnknown, roleNone},
			wantProblems: []string{
				unreachableC,
				`clusters disagree on the lease holder ("` + testIdentityB + `" reported by [gke_proj-b_us-east1_cluster-b]; "` + testIdentityC + `" reported by [gke_proj-a_us-central1_cluster-a])`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fs := healthyHAFixtures()
			if tc.fixtures != nil {
				fs = tc.fixtures()
			}
			if tc.mutate != nil {
				tc.mutate(fs)
			}

			report, out := runHAStatus(t, fs)

			if report.state != tc.wantState {
				t.Fatalf("got state %q, want %q; output:\n%s", report.state, tc.wantState, out)
			}
			for i, c := range report.clusters {
				if c.role != tc.wantRoles[i] {
					t.Fatalf("cluster %q: got role %q, want %q; output:\n%s", c.gke.contextName, c.role, tc.wantRoles[i], out)
				}
			}
			if tc.wantLease == "" {
				if strings.Contains(out, "Lease:") {
					t.Fatalf("expected no Lease line, got:\n%s", out)
				}
			} else if !strings.Contains(out, tc.wantLease) {
				t.Fatalf("expected output to contain %q, got:\n%s", tc.wantLease, out)
			}
			if len(report.problems) != len(tc.wantProblems) {
				t.Fatalf("got %d problems, want %d; output:\n%s", len(report.problems), len(tc.wantProblems), out)
			}
			for i, want := range tc.wantProblems {
				if !strings.Contains(report.problems[i], want) {
					t.Fatalf("problem %d: got %q, want it to contain %q; output:\n%s", i, report.problems[i], want, out)
				}
			}
		})
	}
}

func TestLeaseExpired(t *testing.T) {
	seconds := func(n int32) *int32 { return &n }
	tests := []struct {
		name        string
		renewTime   string
		duration    *int32
		wantExpired bool
	}{
		{name: "renewed recently", renewTime: "2026-11-10T18:05:56Z"},
		{name: "within twice the default duration", renewTime: "2026-11-10T18:05:31Z"},
		{name: "older than twice the default duration", renewTime: "2026-11-10T18:05:29Z", wantExpired: true},
		{name: "reported lease duration is used", renewTime: "2026-11-10T18:05:00Z", duration: seconds(60)},
		{name: "no renew time", renewTime: ""},
		{name: "unparseable renew time", renewTime: "yesterday", wantExpired: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lease := newTestLease(testIdentityA, tc.renewTime)
			lease.Status.GlobalLeaseDurationSeconds = tc.duration

			got, _ := leaseExpired(lease, testNow)

			if got != tc.wantExpired {
				t.Fatalf("leaseExpired() = %v, want %v", got, tc.wantExpired)
			}
		})
	}
}

// errorStorage wraps Storage to return an error from ReadLease.
type errorStorage struct {
	mclstorage.Storage
	err error
}

func (s *errorStorage) ReadLease(ctx context.Context, key string) (*mclstorage.LeaseObject, error) {
	return nil, s.err
}

// TestStandbyStaleCandidate_Contract tests that liveLeaseHolder handles the real status produced
// by the vendored MultiClusterLeaseReconciler when reconciling a standby whose spec.renewTime is stale.
func TestStandbyStaleCandidate_Contract(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	if err := mclv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme: %v", err)
	}

	reconcileStandbyLease := func(t *testing.T, s mclstorage.Storage) *mclv1alpha1.MultiClusterLease {
		t.Helper()
		fakeClient := crfake.NewClientBuilder().
			WithScheme(scheme).
			WithStatusSubresource(&mclv1alpha1.MultiClusterLease{}).
			Build()

		reconciler := mclcontroller.NewMultiClusterLeaseReconciler(
			fakeClient,
			ctrl.Log.WithName("test-mcl-reconciler"),
			s,
		)

		staleRenewTime := metav1.NewMicroTime(testNow.Add(-60 * time.Second))
		duration := int32(15)
		standbyCandidate := testIdentityB
		mcl := &mclv1alpha1.MultiClusterLease{
			ObjectMeta: metav1.ObjectMeta{
				Name:      testLeaseName,
				Namespace: testLeaseNamespace,
			},
			Spec: mclv1alpha1.MultiClusterLeaseSpec{
				HolderIdentity:       &standbyCandidate,
				RenewTime:            &staleRenewTime,
				LeaseDurationSeconds: &duration,
			},
		}
		if err := fakeClient.Create(ctx, mcl); err != nil {
			t.Fatalf("create MultiClusterLease: %v", err)
		}

		req := ctrl.Request{
			NamespacedName: types.NamespacedName{
				Namespace: testLeaseNamespace,
				Name:      testLeaseName,
			},
		}
		if _, err := reconciler.Reconcile(ctx, req); err != nil {
			t.Fatalf("reconcile: %v", err)
		}

		var reconciled mclv1alpha1.MultiClusterLease
		if err := fakeClient.Get(ctx, req.NamespacedName, &reconciled); err != nil {
			t.Fatalf("get reconciled MultiClusterLease: %v", err)
		}
		return &reconciled
	}

	t.Run("stale standby candidate produces BackendHealthy=False but ha status reports HEALTHY", func(t *testing.T) {
		fakeStorage := mclstorage.NewFakeStorage()
		// Seed fake storage with a live lease held by cluster A.
		_, err := fakeStorage.CreateLease(ctx, testLeaseNamespace+"/"+testLeaseName, mclstorage.LeaseData{
			HolderIdentity:   testIdentityA,
			RenewTime:        testNow.Add(-4 * time.Second),
			LeaseTransitions: 1,
		})
		if err != nil {
			t.Fatalf("seed fake storage: %v", err)
		}

		reconciled := reconcileStandbyLease(t, fakeStorage)

		// Verify the MCL controller recorded BackendHealthy=False with "candidate lease is stale".
		cond := meta.FindStatusCondition(reconciled.Status.Conditions, string(mclv1alpha1.ConditionTypeBackendHealthy))
		if cond == nil || cond.Status != metav1.ConditionFalse || !strings.Contains(cond.Message, "candidate lease is stale") {
			t.Fatalf("expected BackendHealthy=False with candidate lease is stale, got %+v", cond)
		}

		fs := healthyHAFixtures()
		fs[1].lease = reconciled

		report, out := runHAStatus(t, fs)
		if report.state != haStateHealthy {
			t.Fatalf("got state %q, want %q; output:\n%s", report.state, haStateHealthy, out)
		}
	})

	t.Run("storage failure during standby reconcile reports backend error", func(t *testing.T) {
		fakeStorage := mclstorage.NewFakeStorage()
		errStg := &errorStorage{
			Storage: fakeStorage,
			err:     errors.New("storage network error"),
		}

		reconciled := reconcileStandbyLease(t, errStg)

		// Verify the MCL controller recorded BackendHealthy=False with "failed to read lease".
		cond := meta.FindStatusCondition(reconciled.Status.Conditions, string(mclv1alpha1.ConditionTypeBackendHealthy))
		if cond == nil || cond.Status != metav1.ConditionFalse || !strings.Contains(cond.Message, "storage network error") {
			t.Fatalf("expected BackendHealthy=False with storage network error, got %+v", cond)
		}

		fs := healthyHAFixtures()
		fs[1].lease = reconciled

		report, out := runHAStatus(t, fs)
		if report.state != haStateError {
			t.Fatalf("got state %q, want %q; output:\n%s", report.state, haStateError, out)
		}
		if !strings.Contains(out, "the lease backend is unhealthy: failed to communicate with backend: failed to read lease: storage network error") {
			t.Fatalf("expected output to contain storage network error, got:\n%s", out)
		}
	})
}
