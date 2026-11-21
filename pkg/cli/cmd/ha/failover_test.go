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
	"slices"
	"strings"
	"testing"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestFailoverValidateNormalizesFields(t *testing.T) {
	opts := newFailoverOptions()
	opts.contexts = []string{" cluster-a ", " cluster-b"}
	opts.targetCluster = " cluster-b "

	err := opts.validate()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantContexts := []string{"cluster-a", "cluster-b"}
	if !slices.Equal(opts.contexts, wantContexts) {
		t.Fatalf("got contexts %v, want %v", opts.contexts, wantContexts)
	}
	if opts.targetCluster != "cluster-b" {
		t.Fatalf("got targetCluster %q, want %q", opts.targetCluster, "cluster-b")
	}
}

func TestFailoverValidateErrors(t *testing.T) {
	tests := []struct {
		name              string
		contexts          []string
		targetCluster     string
		timeout           time.Duration
		wantErrSubstrings []string
	}{
		{
			name:              "missing target-cluster",
			contexts:          []string{"cluster-a", "cluster-b"},
			timeout:           time.Minute,
			wantErrSubstrings: []string{"--target-cluster is required"},
		},
		{
			name:              "target-cluster not in contexts",
			contexts:          []string{"cluster-a", "cluster-b"},
			targetCluster:     "cluster-c",
			timeout:           time.Minute,
			wantErrSubstrings: []string{"must be one of the contexts specified in --contexts"},
		},
		{
			name:              "accumulates contexts and target-cluster errors",
			contexts:          []string{"cluster-a"},
			targetCluster:     "",
			timeout:           time.Minute,
			wantErrSubstrings: []string{"at least 2 cluster contexts", "--target-cluster is required"},
		},
		{
			name:              "timeout must be at least one minute",
			contexts:          []string{"cluster-a", "cluster-b"},
			targetCluster:     "cluster-b",
			timeout:           0,
			wantErrSubstrings: []string{"--timeout must be at least 1m0s (lease expiry takes about a minute), got 0s"},
		},
		{
			name:              "timeout shorter than lease expiry is rejected",
			contexts:          []string{"cluster-a", "cluster-b"},
			targetCluster:     "cluster-b",
			timeout:           30 * time.Second,
			wantErrSubstrings: []string{"--timeout must be at least 1m0s (lease expiry takes about a minute), got 30s"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opts := newFailoverOptions()
			opts.contexts = slices.Clone(tc.contexts)
			opts.targetCluster = tc.targetCluster
			opts.timeout = tc.timeout

			err := opts.validate()

			if err == nil {
				t.Fatalf("expected error containing %v, got nil", tc.wantErrSubstrings)
			}
			for _, sub := range tc.wantErrSubstrings {
				if !strings.Contains(err.Error(), sub) {
					t.Fatalf("expected error %q to contain %q", err.Error(), sub)
				}
			}
		})
	}
}

func TestFailoverFlagDefaults(t *testing.T) {
	flags := NewFailoverCmd().Flags()
	if got, want := flags.Lookup(failoverTimeoutFlag).DefValue, "5m0s"; got != want {
		t.Errorf("default of --%s is %q, want %q", failoverTimeoutFlag, got, want)
	}
	if got, want := flags.Lookup(failoverDryRunFlag).DefValue, "false"; got != want {
		t.Errorf("default of --%s is %q, want %q", failoverDryRunFlag, got, want)
	}

	opts := newFailoverOptions()
	if opts.timeout != 5*time.Minute {
		t.Errorf("opts.timeout = %s, want 5m0s", opts.timeout)
	}
	if opts.pollInterval != 2*time.Second {
		t.Errorf("opts.pollInterval = %s, want 2s", opts.pollInterval)
	}
	if opts.leaderAgreementTimeout != 30*time.Second {
		t.Errorf("opts.leaderAgreementTimeout = %s, want 30s", opts.leaderAgreementTimeout)
	}
	if opts.clusterOpTimeout != time.Minute {
		t.Errorf("opts.clusterOpTimeout = %s, want 1m0s", opts.clusterOpTimeout)
	}
	if opts.signalBounceWindow != time.Second {
		t.Errorf("opts.signalBounceWindow = %s, want 1s", opts.signalBounceWindow)
	}
}

func TestNewFailoverCmdRunE(t *testing.T) {
	t.Run("validation errors keep usage output enabled", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		cmd := NewFailoverCmd()
		cmd.SetOut(&stdout)
		cmd.SetErr(&stderr)
		cmd.SetArgs([]string{"--contexts=cluster-a,cluster-b"})

		err := cmd.ExecuteContext(context.Background())

		requireErrorContains(t, err, "--target-cluster is required")
		if cmd.SilenceUsage {
			t.Errorf("SilenceUsage is true after a flag validation error, want false")
		}
		if !strings.Contains(stdout.String()+stderr.String(), "Usage:") {
			t.Errorf("output does not include command usage:\nstdout:\n%s\nstderr:\n%s", stdout.String(), stderr.String())
		}
	})

	t.Run("cluster errors silence usage output", func(t *testing.T) {
		useTestKubeconfig(t, "cluster-a", "cluster-b")
		var stdout, stderr bytes.Buffer
		cmd := NewFailoverCmd()
		cmd.SetOut(&stdout)
		cmd.SetErr(&stderr)
		cmd.SetArgs([]string{"--contexts=cluster-a,cluster-b", "--target-cluster=cluster-b"})

		err := cmd.ExecuteContext(context.Background())

		requireErrorContains(t, err, `fetching ConfigConnector on cluster "cluster-a"`)
		if !cmd.SilenceUsage {
			t.Errorf("SilenceUsage is false after validation succeeded, want true")
		}
		if strings.Contains(stdout.String()+stderr.String(), "Usage:") {
			t.Errorf("output includes command usage on a cluster error:\nstdout:\n%s\nstderr:\n%s", stdout.String(), stderr.String())
		}
	})
}

func TestBuildClusterHandlesSetsRequestTimeout(t *testing.T) {
	useTestKubeconfig(t, "cluster-a", "cluster-b")

	handles, err := buildClusterHandles(context.Background(), []string{"cluster-a", "cluster-b"})
	if err != nil {
		t.Fatalf("buildClusterHandles failed: %v", err)
	}
	for _, h := range handles {
		if got := h.client.RestConfig.Timeout; got != defaultClusterRequestTimeout {
			t.Errorf("cluster %q RestConfig.Timeout = %s, want %s", h.contextName, got, defaultClusterRequestTimeout)
		}
	}
}

// Failover reads every cluster before it changes any, so it fails when one can't be reached.
func TestRunFailoverUnreachableClusters(t *testing.T) {
	useTestKubeconfig(t, "cluster-a", "cluster-b")
	opts := newFailoverOptions()
	opts.contexts = []string{"cluster-a", "cluster-b"}
	opts.targetCluster = "cluster-b"

	err := runFailover(context.Background(), opts)

	for _, want := range []string{`fetching ConfigConnector on cluster "cluster-a"`, `fetching ConfigConnector on cluster "cluster-b"`} {
		if err == nil || !strings.Contains(err.Error(), want) {
			t.Fatalf("got error %v, want it to contain %q", err, want)
		}
	}
}

const fleetManager = "cnrm-system/cnrm-controller-manager"

func fleetIdentity(cluster string) string {
	return "projects/test-project/locations/us-central1/clusters/" + cluster
}

// failOver runs failover to the target on every cluster of the fleet, and returns its output.
func (f *haFleet) failOver(target string, options ...func(*failoverOptions)) (string, error) {
	f.t.Helper()
	return f.failOverOn(context.Background(), f.handles(), target, options...)
}

func (f *haFleet) failOverOn(ctx context.Context, handles []*clusterHandle, target string, options ...func(*failoverOptions)) (string, error) {
	f.t.Helper()
	var out bytes.Buffer
	opts := newFailoverOptions()
	opts.targetCluster = target
	opts.stdout = &out
	// Long enough that a wait which is to succeed doesn't run out of time on a busy machine.
	opts.timeout = 10 * time.Second
	opts.pollInterval = time.Millisecond
	opts.leaderAgreementTimeout = 10 * time.Second
	for _, option := range options {
		option(&opts)
	}
	err := executeFailover(ctx, opts, handles)
	return out.String(), err
}

// shortTestTimeout is the timeout of a wait that a test expects to run out of time. Failover
// looks at the clusters as soon as a wait starts, and only that first look has to fit in.
const shortTestTimeout = 500 * time.Millisecond

// failOverUntilTimeout runs failover in a test in which one of its waits is to time out. Pods
// then stop and the lease moves within one step of the fleet, so that the waits before that one
// end at failover's first look, however slow the machine is.
func (f *haFleet) failOverUntilTimeout(target string) (string, error) {
	f.t.Helper()
	f.podStopSteps, f.takeOverSteps = 1, 1
	return f.failOver(target, func(opts *failoverOptions) { opts.timeout = shortTestTimeout })
}

// shortAgreement is for tests in which the clusters don't agree on the leader.
func shortAgreement(opts *failoverOptions) {
	opts.leaderAgreementTimeout = shortTestTimeout
}

func requireErrorContains(t *testing.T, err error, wants ...string) {
	t.Helper()
	if err == nil {
		t.Fatalf("got no error, want one containing %q", wants)
	}
	for _, want := range wants {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("got error:\n  %v\nwant it to contain:\n  %s", err, want)
		}
	}
}

func TestFailoverMovesTheLeaseToTheTarget(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
	fleet.requireLeader(a)

	out, err := fleet.failOver("cluster-b")
	if err != nil {
		t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
	}

	fleet.requireLeader(b)
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	// The leader's manager is stopped while the leader holds the lease, and it is started again
	// only once the target holds it. The target's manager isn't touched.
	fleet.requirePatches(
		`stop cluster-a `+fleetManager+` (annotation "1") while cluster-a leads and cluster-a,cluster-b run`,
		`start cluster-a `+fleetManager+` (replicas 1, annotation "") while cluster-b leads`,
	)

	want := `Leader: cluster "cluster-a" (projects/test-project/locations/us-central1/clusters/cluster-a)
Target: cluster "cluster-b" (projects/test-project/locations/us-central1/clusters/cluster-b)
If this command doesn't finish, run it again: it starts the managers that are still stopped.
cluster "cluster-a": stopped Config Connector's manager (StatefulSet cnrm-system/cnrm-controller-manager, scaled from 1 to 0)
cluster "cluster-a": waiting up to 10s for its manager pods to stop
No cluster reconciles resources until cluster "cluster-b" takes the lease. The old lease has to expire first, which takes about a minute.
cluster "cluster-b": waiting up to 10s for it to become the leader
cluster "cluster-b": is the leader
cluster "cluster-a": started Config Connector's manager again (StatefulSet cnrm-system/cnrm-controller-manager, scaled from 0 to 1)
cluster "cluster-a": waiting up to 10s for it to follow the new leader
cluster "cluster-a": follows the new leader
Failover complete: cluster "cluster-b" (projects/test-project/locations/us-central1/clusters/cluster-b) is the leader.
`
	if out != want {
		t.Errorf("got output:\n%s\nwant:\n%s", out, want)
	}
}

// With more than two clusters, any standby can take an expired lease. So failover stops the
// other standbys too, and it stops them before the leader.
func TestFailoverStopsTheOtherStandbysBeforeTheLeader(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
	a, b, c := fleet.cluster("cluster-a"), fleet.cluster("cluster-b"), fleet.cluster("cluster-c")

	// In the fleet, the first cluster with a running manager takes an expired lease. That is
	// cluster-b, unless failover stops it.
	out, err := fleet.failOver("cluster-c")
	if err != nil {
		t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
	}

	fleet.requireLeader(c)
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	b.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	fleet.requirePatches(
		`stop cluster-b `+fleetManager+` (annotation "1") while cluster-a leads and cluster-a,cluster-b,cluster-c run`,
		// The standby's manager is gone by the time the leader's is stopped.
		`stop cluster-a `+fleetManager+` (annotation "1") while cluster-a leads and cluster-a,cluster-c run`,
		`start cluster-a `+fleetManager+` (replicas 1, annotation "") while cluster-c leads`,
		`start cluster-b `+fleetManager+` (replicas 1, annotation "") while cluster-c leads`,
	)
}

// In namespaced mode a cluster has one manager per namespace, and each keeps the lease. Failover
// stops them all, wherever they run.
func TestFailoverInNamespacedMode(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
	for _, c := range fleet.clusters {
		// The manager of team-c has a namespace of its own (spec.managerNamespace).
		c.useNamespacedMode(map[string]string{"team-a": "cnrm-system", "team-b": "cnrm-system", "team-c": "team-c-system"})
	}
	fleet.settle()
	fleet.requireLeader(a)
	if got := len(b.syncers()); got != 3 {
		t.Fatalf("cluster-b has %d KRMSyncers, want one per namespace", got)
	}

	out, err := fleet.failOver("cluster-b")
	if err != nil {
		t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
	}

	fleet.requireLeader(b)
	a.requireManager("cnrm-system", "cnrm-controller-manager-team-a", 1)
	a.requireManager("cnrm-system", "cnrm-controller-manager-team-b", 1)
	a.requireManager("team-c-system", "cnrm-controller-manager-team-c", 1)
	fleet.requirePatches(
		`stop cluster-a cnrm-system/cnrm-controller-manager-team-a (annotation "1") while cluster-a leads and cluster-a,cluster-b run`,
		`stop cluster-a cnrm-system/cnrm-controller-manager-team-b (annotation "1") while cluster-a leads and cluster-a,cluster-b run`,
		`stop cluster-a team-c-system/cnrm-controller-manager-team-c (annotation "1") while cluster-a leads and cluster-a,cluster-b run`,
		`start cluster-a cnrm-system/cnrm-controller-manager-team-a (replicas 1, annotation "") while cluster-b leads`,
		`start cluster-a cnrm-system/cnrm-controller-manager-team-b (replicas 1, annotation "") while cluster-b leads`,
		`start cluster-a team-c-system/cnrm-controller-manager-team-c (replicas 1, annotation "") while cluster-b leads`,
	)
}

// A manager that someone scaled to 0 stays at 0, and so does the KRMSyncer that it maintained.
func TestFailoverLeavesAManagerAloneThatSomeoneElseStopped(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
	for _, c := range fleet.clusters {
		c.useNamespacedMode(map[string]string{"team-a": "cnrm-system", "team-b": "cnrm-system"})
	}
	fleet.settle()
	for _, c := range fleet.clusters {
		c.scaleManager(c.manager("cnrm-system", "cnrm-controller-manager-team-b"), 0)
	}
	fleet.settle()

	out, err := fleet.failOver("cluster-b")
	if err != nil {
		t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
	}

	fleet.requireLeader(b)
	a.requireManager("cnrm-system", "cnrm-controller-manager-team-a", 1)
	a.requireManager("cnrm-system", "cnrm-controller-manager-team-b", 0)
	b.requireManager("cnrm-system", "cnrm-controller-manager-team-b", 0)
	fleet.requirePatches(
		`stop cluster-a cnrm-system/cnrm-controller-manager-team-a (annotation "1") while cluster-a leads and cluster-a,cluster-b run`,
		`start cluster-a cnrm-system/cnrm-controller-manager-team-a (replicas 1, annotation "") while cluster-b leads`,
	)
	// The KRMSyncer of the stopped manager still names the old leader. Nothing maintains it, so
	// failover doesn't wait for it.
	if got := b.syncer(fleetLeaseName + "-team-b").Spec.Remote.GKECluster.Name; got != "cluster-a" {
		t.Errorf("the KRMSyncer of the stopped manager names %q, want it unchanged", got)
	}
	want := `cluster "cluster-b": note: StatefulSet cnrm-system/cnrm-controller-manager-team-b is scaled to 0, so that manager will not run after the failover either`
	if !strings.Contains(out, want) {
		t.Errorf("got output:\n%s\nwant it to contain:\n%s", out, want)
	}
}

// Nothing deletes the KRMSyncer of a namespace that Config Connector no longer manages.
func TestFailoverIgnoresTheKRMSyncerOfAManagerThatIsGone(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	b := fleet.cluster("cluster-b")
	gone := operatorv1beta1.GKECluster{Project: "test-project", Location: "us-central1", Name: "cluster-z"}
	for _, c := range fleet.clusters {
		c.create(&syncerv1alpha1.KRMSyncer{
			ObjectMeta: metav1.ObjectMeta{Namespace: fleetLeaseNamespace, Name: fleetLeaseName + "-team-gone"},
			Spec:       newSyncerSpec(gone, true),
		})
	}

	out, err := fleet.failOver("cluster-b")
	if err != nil {
		t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
	}
	fleet.requireLeader(b)
}

func TestFailoverWithoutResourceReplication(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
	for _, c := range fleet.clusters {
		cc := &operatorv1beta1.ConfigConnector{}
		if err := c.backend.Get(context.Background(), types.NamespacedName{Name: configConnectorResourceName}, cc); err != nil {
			t.Fatalf("getting the ConfigConnector: %v", err)
		}
		cc.Spec.Experiments.MultiClusterLease.ResourceReplicationMode = "Disabled"
		if err := c.backend.Update(context.Background(), cc); err != nil {
			t.Fatalf("updating the ConfigConnector: %v", err)
		}
		// The managers don't maintain these anymore.
		for _, syncer := range c.syncers() {
			if err := c.backend.Delete(context.Background(), &syncer); err != nil {
				t.Fatalf("deleting KRMSyncer %s: %v", syncer.Name, err)
			}
		}
	}

	out, err := fleet.failOver("cluster-b")
	if err != nil {
		t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
	}

	fleet.requireLeader(b)
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	if got := len(a.syncers()) + len(b.syncers()); got != 0 {
		t.Errorf("the clusters have %d KRMSyncers, want none", got)
	}
}

func TestFailoverToTheLeaderDoesNothing(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a := fleet.cluster("cluster-a")

	out, err := fleet.failOver("cluster-a")
	if err != nil {
		t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
	}

	fleet.requireLeader(a)
	fleet.requirePatches()
	want := `Leader: cluster "cluster-a" (projects/test-project/locations/us-central1/clusters/cluster-a)
Cluster "cluster-a" is already the leader. Nothing to do.
`
	if out != want {
		t.Errorf("got output:\n%s\nwant:\n%s", out, want)
	}
}

// MultiClusterLease.status.globalHolderIdentity keeps naming the last holder after its managers
// stop. Failing over to that cluster must not report "Nothing to do." when it has no healthy
// running manager.
func TestFailoverToTheLeaderFailsWhenTheLeaderIsUnhealthy(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(t *testing.T, fleet *haFleet, a, b *fleetCluster)
		wantErr string
	}{
		{
			name: "someone scaled its manager to 0",
			prepare: func(t *testing.T, fleet *haFleet, a, b *fleetCluster) {
				b.cannotLead = true
				a.scaleManager(a.manager("cnrm-system", "cnrm-controller-manager"), 0)
				fleet.settle()
			},
			wantErr: `cluster "cluster-a" holds the lease, but: none of its Config Connector managers runs (StatefulSet ` + fleetManager + ` scaled to 0)`,
		},
		{
			name: "it has no manager StatefulSet",
			prepare: func(t *testing.T, fleet *haFleet, a, b *fleetCluster) {
				b.cannotLead = true
				if err := a.backend.Delete(context.Background(), a.manager("cnrm-system", "cnrm-controller-manager")); err != nil {
					t.Fatalf("deleting the manager: %v", err)
				}
			},
			wantErr: `cluster "cluster-a" holds the lease, but: it has no StatefulSet of Config Connector's manager`,
		},
		{
			name: "its manager pod is not ready",
			prepare: func(t *testing.T, fleet *haFleet, a, _ *fleetCluster) {
				a.unreadyPods = true
				fleet.settle()
			},
			wantErr: `cluster "cluster-a" holds the lease, but: StatefulSet ` + fleetManager + ` has 0 of 1 pods ready`,
		},
		{
			name: "its KRMSyncer is not suspended",
			prepare: func(t *testing.T, fleet *haFleet, a, _ *fleetCluster) {
				a.abandonedSyncers = true
				a.setSyncerSpec(fleetLeaseName, newSyncerSpec(a.gke, false))
			},
			wantErr: `cluster "cluster-a" holds the lease, but: KRMSyncer cnrm-system/kcc-multicluster-lease is not suspended yet`,
		},
		{
			name: "its MultiClusterLease has no valid renew time",
			prepare: func(t *testing.T, fleet *haFleet, a, _ *fleetCluster) {
				a.staleLease = true
				a.setLeaseStatus(fleetIdentity("cluster-a"), "not a time")
			},
			wantErr: `cluster "cluster-a" doesn't report when the lease was last renewed`,
		},
	}

	for _, tc := range tests {
		for _, dryRun := range []bool{false, true} {
			mode := "live"
			if dryRun {
				mode = "dry-run"
			}
			t.Run(tc.name+"/"+mode, func(t *testing.T) {
				fleet := newHAFleet(t, "cluster-a", "cluster-b")
				a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
				tc.prepare(t, fleet, a, b)

				out, err := fleet.failOver("cluster-a", func(opts *failoverOptions) { opts.dryRun = dryRun })

				requireErrorContains(t, err, tc.wantErr)
				if strings.Contains(out, "Nothing to do") {
					t.Errorf("output falsely reported Nothing to do:\n%s", out)
				}
				fleet.requirePatches()
			})
		}
	}
}

func TestFailoverHungClusterDoesNotStarveOtherClusters(t *testing.T) {
	const opTimeout = 50 * time.Millisecond

	t.Run("during startOtherManagers after handover completes", func(t *testing.T) {
		fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
		a, b, c := fleet.cluster("cluster-a"), fleet.cluster("cluster-b"), fleet.cluster("cluster-c")
		fleet.onStep = func() {
			if fleet.holder == c {
				a.rejectRead = func(ctx context.Context, obj runtime.Object) error {
					if _, ok := obj.(*appsv1.StatefulSetList); ok {
						<-ctx.Done()
						return ctx.Err()
					}
					return nil
				}
			}
		}

		_, err := fleet.failOver("cluster-c", func(opts *failoverOptions) {
			opts.clusterOpTimeout = opTimeout
		})

		requireErrorContains(t, err,
			`cluster "cluster-c" is the leader, but not every manager on the other clusters was started again`,
			`any managers that failover stopped on cluster "cluster-a" are still scaled to 0 — run this command again once the cluster is reachable`,
		)
		// cluster-b still gets its own restart budget and comes back up.
		b.requireManager("cnrm-system", "cnrm-controller-manager", 1)
		fleet.requireLeader(c)
	})

	t.Run("during stopManagers after another standby was already stopped", func(t *testing.T) {
		fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
		a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
		fleet.onStep = func() {
			// Once cluster-b has been stopped, cluster-a stalls when stopManagers lists its managers.
			if len(fleet.patches) == 1 {
				a.rejectRead = func(ctx context.Context, obj runtime.Object) error {
					if _, ok := obj.(*appsv1.StatefulSetList); ok {
						<-ctx.Done()
						return ctx.Err()
					}
					return nil
				}
			}
		}

		_, err := fleet.failOver("cluster-c", func(opts *failoverOptions) {
			opts.clusterOpTimeout = opTimeout
		})

		requireErrorContains(t, err,
			`failover to cluster "cluster-c" did not complete:`,
			`listing the StatefulSets of Config Connector's manager on cluster "cluster-a": context deadline exceeded`,
			`any managers that failover stopped on cluster "cluster-a" are still scaled to 0`,
		)
		b.requireManager("cnrm-system", "cnrm-controller-manager", 1)
		fleet.settle()
		fleet.requireLeader(a)
	})

	t.Run("during finishEarlierFailover when target hangs", func(t *testing.T) {
		fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
		a, b, c := fleet.cluster("cluster-a"), fleet.cluster("cluster-b"), fleet.cluster("cluster-c")
		c.cannotLead = true
		a.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager")
		b.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager")
		fleet.settle()
		reads := 0
		a.rejectRead = func(ctx context.Context, obj runtime.Object) error {
			if _, ok := obj.(*appsv1.StatefulSetList); ok {
				reads++
				// Let currentLeader read cluster-a, then stall when finishEarlierFailover starts target.
				if reads > 1 {
					<-ctx.Done()
					return ctx.Err()
				}
			}
			return nil
		}

		_, err := fleet.failOver("cluster-a", func(opts *failoverOptions) {
			opts.clusterOpTimeout = opTimeout
		})

		requireErrorContains(t, err, `any managers that failover stopped on cluster "cluster-a" are still scaled to 0`)
		b.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	})
}

func TestFailoverStandbyWithNoRunningManagersLogsAccurately(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
	a, b, c := fleet.cluster("cluster-a"), fleet.cluster("cluster-b"), fleet.cluster("cluster-c")
	b.scaleManager(b.manager("cnrm-system", "cnrm-controller-manager"), 0)
	fleet.settle()

	out, err := fleet.failOver("cluster-c")
	if err != nil {
		t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
	}

	fleet.requireLeader(c)
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	b.requireManager("cnrm-system", "cnrm-controller-manager", 0)
	if !strings.Contains(out, `cluster "cluster-b": has no running manager to stop`) {
		t.Errorf("expected output to report no running manager to stop on cluster-b:\n%s", out)
	}
	if !strings.Contains(out, `cluster "cluster-b": has no running manager`) {
		t.Errorf("expected output to report no running manager on cluster-b:\n%s", out)
	}
	if strings.Contains(out, `cluster "cluster-b": follows the new leader`) {
		t.Errorf("output falsely reported cluster-b following the new leader:\n%s", out)
	}
}

func TestFailoverDryRunChangesNothing(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
	a := fleet.cluster("cluster-a")

	out, err := fleet.failOver("cluster-c", func(opts *failoverOptions) { opts.dryRun = true })
	if err != nil {
		t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
	}

	fleet.requireLeader(a)
	fleet.requirePatches()
	want := `Leader: cluster "cluster-a" (projects/test-project/locations/us-central1/clusters/cluster-a)
Target: cluster "cluster-c" (projects/test-project/locations/us-central1/clusters/cluster-c)
[dry-run] cluster "cluster-b": would stop StatefulSet cnrm-system/cnrm-controller-manager, and wait up to 10s for its pods to stop
[dry-run] cluster "cluster-a": would stop StatefulSet cnrm-system/cnrm-controller-manager, and wait up to 10s for its pods to stop
[dry-run] cluster "cluster-c": would wait up to 10s for it to become the leader
[dry-run] would then start the stopped managers again, and wait up to 10s for their clusters to follow the new leader
`
	if out != want {
		t.Errorf("got output:\n%s\nwant:\n%s", out, want)
	}
}

// If the target doesn't take the lease, the old leader's managers are started again, and it leads
// again.
func TestFailoverStartsTheManagersAgainWhenTheTargetDoesNotTakeOver(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
	b.cannotLead = true

	out, err := fleet.failOverUntilTimeout("cluster-b")

	requireErrorContains(t, err,
		"failover to cluster \"cluster-b\" did not complete, and the other clusters' managers were started again; run the command with --dry-run to see which cluster leads now: "+
			`cluster "cluster-b" did not become the leader: timed out after 500ms: its MultiClusterLease cnrm-system/kcc-multicluster-lease still names `+fleetIdentity("cluster-a")+` as the holder`,
	)
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	fleet.requirePatches(
		`stop cluster-a `+fleetManager+` (annotation "1") while cluster-a leads and cluster-a,cluster-b run`,
		`start cluster-a `+fleetManager+` (replicas 1, annotation "") while cluster-a leads`,
	)
	fleet.settle()
	fleet.requireLeader(a)
	if strings.Contains(out, "Failover complete") {
		t.Errorf("the output reports success:\n%s", out)
	}
}

// Ctrl-C cancels the command's context. The managers are still started again.
func TestFailoverStartsTheManagersAgainWhenInterrupted(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a := fleet.cluster("cluster-a")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fleet.onStep = func() {
		if len(fleet.patches) == 1 {
			cancel()
		}
	}

	_, err := fleet.failOverOn(ctx, fleet.handles(), "cluster-b")

	requireErrorContains(t, err, "interrupted: context canceled", "the other clusters' managers were started again")
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	fleet.settle()
	fleet.requireLeader(a)
}

// A pod that doesn't stop still renews the lease, so failover gives up.
func TestFailoverStartsTheManagersAgainWhenAPodDoesNotStop(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a := fleet.cluster("cluster-a")
	a.stuckPods = true

	_, err := fleet.failOverUntilTimeout("cluster-b")

	requireErrorContains(t, err,
		`cluster "cluster-a": Config Connector's manager did not stop: timed out after 500ms: StatefulSet cnrm-system/cnrm-controller-manager still has pods (1)`,
		"the other clusters' managers were started again",
	)
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	fleet.requireLeader(a)
}

// While failover waits, a request that fails is tried again.
func TestFailoverKeepsWaitingWhenARequestFails(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
	// After each write of failover, the next 2 reads from each cluster fail. So reads fail while
	// failover waits for the leader's pods to stop, for the target to lead, and for the old
	// leader to follow.
	type failures struct{ patches, sincePatch, total int }
	failed := map[*fleetCluster]*failures{a: {}, b: {}}
	for _, c := range fleet.clusters {
		c.rejectRead = func(context.Context, runtime.Object) error {
			f := failed[c]
			if n := len(fleet.patches); n != f.patches {
				f.patches, f.sincePatch = n, 0
			}
			if f.patches == 0 || f.sincePatch == 2 {
				return nil
			}
			f.sincePatch++
			f.total++
			return errors.New("the API server is busy")
		}
	}

	out, err := fleet.failOver("cluster-b")
	if err != nil {
		t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
	}

	fleet.requireLeader(b)
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	// Failover reads from cluster-a after it stopped and after it started the manager, and from
	// cluster-b only in between.
	if failed[a].total != 4 || failed[b].total != 2 {
		t.Errorf("%d reads from cluster-a and %d from cluster-b failed, want 4 and 2", failed[a].total, failed[b].total)
	}
}

// Before failover changes anything, a request that fails ends it, with the request's error.
func TestFailoverReportsARequestThatFailsBeforeItChangesAnything(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
	b.rejectRead = func(_ context.Context, obj runtime.Object) error {
		if _, ok := obj.(*appsv1.StatefulSetList); ok {
			return errors.New("statefulsets.apps is forbidden")
		}
		return nil
	}

	_, err := fleet.failOver("cluster-b")

	want := `listing the StatefulSets of Config Connector's manager on cluster "cluster-b": statefulsets.apps is forbidden`
	if err == nil || err.Error() != want {
		t.Errorf("got error:\n  %v\nwant:\n  %s", err, want)
	}
	fleet.requirePatches()
	fleet.requireLeader(a)
}

// The request that is in flight when the time to agree on the leader runs out fails. Failover
// reports that the clusters disagree, not that request.
func TestFailoverReportsTheDisagreementWhenTheTimeRunsOutDuringARequest(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
	b.staleLease = true
	b.setLeaseStatus(fleetIdentity("cluster-z"), fleet.renewTime())
	// Failover has looked at both clusters several times by its 30th read. After that, a read
	// takes until the time runs out.
	reads := 0
	slow := func(ctx context.Context, _ runtime.Object) error {
		if reads++; reads <= 30 {
			return nil
		}
		<-ctx.Done()
		return ctx.Err()
	}
	a.rejectRead, b.rejectRead = slow, slow

	_, err := fleet.failOver("cluster-b", shortAgreement)

	want := `cannot tell which cluster is the leader: timed out after 500ms: the clusters disagree: cluster "cluster-a" reports ` + fleetIdentity("cluster-a") + `; cluster "cluster-b" reports ` + fleetIdentity("cluster-z")
	if err == nil || err.Error() != want {
		t.Errorf("got error:\n  %v\nwant:\n  %s", err, want)
	}
	fleet.requirePatches()
}

// The target's lease status names it as the holder from before its managers were stopped. That
// doesn't make it the leader: its managers have to run, and renew the lease.
func TestFailoverFailsWhenTheTargetHoldsTheLeaseButDoesNotRenewIt(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
	b.cannotLead = true
	a.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager")
	fleet.settle()
	if fleet.holder != a || a.running() {
		t.Fatalf("want cluster-a to hold the lease with its manager stopped; holder is %q, running clusters are %v", fleet.holder.name, fleet.runningClusters())
	}
	a.noPods = true

	out, err := fleet.failOverUntilTimeout("cluster-a")

	requireErrorContains(t, err, `cluster "cluster-a" did not become the leader: timed out after 500ms: its MultiClusterLease names it as the holder, but it has not renewed the lease yet`)
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	if strings.Contains(out, "Failover complete") || strings.Contains(out, "is the leader") {
		t.Errorf("the output reports success:\n%s", out)
	}
}

// The target takes the lease, but its KRMSyncers aren't the leader's. Its managers suspend them
// when they start to lead, so they don't lead yet.
func TestFailoverFailsWhenTheTargetDoesNotSuspendItsKRMSyncers(t *testing.T) {
	tests := []struct {
		name string
		// onLeading runs once, when the target has taken the lease.
		onLeading func(b *fleetCluster)
		wantErr   string
	}{
		{
			name:      "they still pull from the old leader",
			onLeading: func(b *fleetCluster) {},
			wantErr:   "KRMSyncer cnrm-system/kcc-multicluster-lease does not name " + fleetIdentity("cluster-b") + " as its remote cluster yet",
		},
		{
			name: "they name the target, but are not suspended",
			onLeading: func(b *fleetCluster) {
				b.setSyncerSpec(fleetLeaseName, newSyncerSpec(b.gke, false))
			},
			wantErr: "KRMSyncer cnrm-system/kcc-multicluster-lease is not suspended yet",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fleet := newHAFleet(t, "cluster-a", "cluster-b")
			a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
			b.abandonedSyncers = true
			leading := false
			fleet.onStep = func() {
				if fleet.holder == b && !leading {
					leading = true
					tc.onLeading(b)
				}
			}

			out, err := fleet.failOverUntilTimeout("cluster-b")

			requireErrorContains(t, err,
				`cluster "cluster-b" did not become the leader: timed out after 500ms: `+tc.wantErr,
				"the other clusters' managers were started again",
			)
			if fleet.holder != b {
				t.Errorf("cluster %q holds the lease, want cluster-b", fleet.holder.name)
			}
			a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
			if strings.Contains(out, "Failover complete") || strings.Contains(out, "is the leader") {
				t.Errorf("the output reports success:\n%s", out)
			}
		})
	}
}

// If something starts a stopped manager again, the lease may never expire, so failover gives up.
func TestFailoverGivesUpWhenAStoppedManagerRunsAgain(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a := fleet.cluster("cluster-a")
	scaledUp := false
	fleet.onStep = func() {
		if len(fleet.patches) == 1 && !scaledUp {
			scaledUp = true
			a.scaleManager(a.manager("cnrm-system", "cnrm-controller-manager"), 1)
		}
	}

	_, err := fleet.failOver("cluster-b")

	requireErrorContains(t, err, "StatefulSet "+fleetManager+" runs again: something scaled it up, or created it, after failover stopped the managers")
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	fleet.requireLeader(a)
}

// A cluster that isn't in --contexts keeps its managers, and can take the lease.
func TestFailoverFailsWhenAnotherClusterTakesTheLease(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
	a, b, c := fleet.cluster("cluster-a"), fleet.cluster("cluster-b"), fleet.cluster("cluster-c")

	_, err := fleet.failOverOn(context.Background(), []*clusterHandle{a.handle, c.handle}, "cluster-c")

	requireErrorContains(t, err,
		`cluster "cluster-c" did not become the leader: `+fleetIdentity("cluster-b")+` took the lease instead; it is not one of the clusters in --contexts, so failover did not stop its managers`,
		"the other clusters' managers were started again",
	)
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	fleet.settle()
	fleet.requireLeader(b)
}

// The MCL controller of a stopped cluster keeps asking for the lease for a while, and can get it
// when the leader's expires. Nothing renews it then, so failover waits for it to expire too.
func TestFailoverWaitsWhenAStoppedClusterTakesTheLeaseFirst(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
	a, b, c := fleet.cluster("cluster-a"), fleet.cluster("cluster-b"), fleet.cluster("cluster-c")
	b.stillAsks = true

	out, err := fleet.failOver("cluster-c")
	if err != nil {
		t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
	}

	if got, want := strings.Join(fleet.holders, ","), "cluster-a,cluster-b,cluster-c"; got != want {
		t.Errorf("the lease went to %s, want %s", got, want)
	}
	fleet.requireLeader(c)
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	b.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	// Nothing is started again before the target leads.
	fleet.requirePatches(
		`stop cluster-b `+fleetManager+` (annotation "1") while cluster-a leads and cluster-a,cluster-b,cluster-c run`,
		`stop cluster-a `+fleetManager+` (annotation "1") while cluster-a leads and cluster-a,cluster-c run`,
		`start cluster-a `+fleetManager+` (replicas 1, annotation "") while cluster-c leads`,
		`start cluster-b `+fleetManager+` (replicas 1, annotation "") while cluster-c leads`,
	)
}

// The target leads, but a cluster whose managers were started again doesn't become its standby.
func TestFailoverReportsAClusterThatDoesNotFollowTheNewLeader(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(a *fleetCluster)
		wantErr string
	}{
		{
			name:    "its manager does not become ready",
			prepare: func(a *fleetCluster) { a.unreadyPods = true },
			wantErr: "StatefulSet " + fleetManager + " has 0 of 1 pods ready",
		},
		{
			name:    "its KRMSyncer does not pull from the new leader",
			prepare: func(a *fleetCluster) { a.abandonedSyncers = true },
			wantErr: "KRMSyncer cnrm-system/kcc-multicluster-lease does not name " + fleetIdentity("cluster-b") + " as its remote cluster yet",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fleet := newHAFleet(t, "cluster-a", "cluster-b")
			a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
			tc.prepare(a)

			out, err := fleet.failOverUntilTimeout("cluster-b")

			requireErrorContains(t, err, `cluster "cluster-b" is the leader, but: cluster "cluster-a" does not follow it yet: timed out after 500ms: `+tc.wantErr)
			if fleet.holder != b {
				t.Errorf("cluster %q holds the lease, want cluster-b", fleet.holder.name)
			}
			a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
			if strings.Contains(out, "Failover complete") {
				t.Errorf("the output reports success:\n%s", out)
			}
		})
	}
}

// Failover reports every cluster that doesn't follow, not only the first one.
func TestFailoverReportsEveryClusterThatDoesNotFollowTheNewLeader(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
	a, b, c := fleet.cluster("cluster-a"), fleet.cluster("cluster-b"), fleet.cluster("cluster-c")
	a.unreadyPods = true
	c.abandonedSyncers = true

	_, err := fleet.failOverUntilTimeout("cluster-b")

	requireErrorContains(t, err,
		`cluster "cluster-b" is the leader, but: cluster "cluster-a" does not follow it yet: timed out after 500ms: StatefulSet `+fleetManager+` has 0 of 1 pods ready`,
		`cluster "cluster-c" does not follow it yet: timed out after 500ms: KRMSyncer cnrm-system/kcc-multicluster-lease does not name `+fleetIdentity("cluster-b")+` as its remote cluster yet`,
	)
	if fleet.holder != b {
		t.Errorf("cluster %q holds the lease, want cluster-b", fleet.holder.name)
	}
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	c.requireManager("cnrm-system", "cnrm-controller-manager", 1)
}

// If a manager can't be stopped, the ones that were stopped already are started again.
func TestFailoverStartsTheManagersAgainWhenOneCannotBeStopped(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a := fleet.cluster("cluster-a")
	for _, c := range fleet.clusters {
		c.useNamespacedMode(map[string]string{"team-a": "cnrm-system", "team-b": "cnrm-system"})
	}
	fleet.settle()
	a.rejectPatch = func(obj client.Object, patch string) error {
		if obj.GetName() == "cnrm-controller-manager-team-b" {
			return errors.New("the API server said no")
		}
		return nil
	}

	_, err := fleet.failOver("cluster-b")

	requireErrorContains(t, err,
		`failover to cluster "cluster-b" did not complete, and the other clusters' managers were started again`,
		`stopping StatefulSet cnrm-system/cnrm-controller-manager-team-b on cluster "cluster-a": the API server said no`,
	)
	a.requireManager("cnrm-system", "cnrm-controller-manager-team-a", 1)
	a.requireManager("cnrm-system", "cnrm-controller-manager-team-b", 1)
	fleet.requirePatches(
		`stop cluster-a cnrm-system/cnrm-controller-manager-team-a (annotation "1") while cluster-a leads and cluster-a,cluster-b run`,
		`start cluster-a cnrm-system/cnrm-controller-manager-team-a (replicas 1, annotation "") while cluster-a leads`,
	)
	fleet.settle()
	fleet.requireLeader(a)
}

// A manager that can't be started again keeps the annotation, so that the next run starts it.
func TestFailoverReportsAManagerThatCannotBeStartedAgain(t *testing.T) {
	rejectStart := func(obj client.Object, patch string) error {
		if strings.Contains(patch, `"replicas":1`) {
			return errors.New("the API server said no")
		}
		return nil
	}
	const wantStartErr = `starting StatefulSet ` + fleetManager + ` on cluster "cluster-a" again: the API server said no; run this command again, or start it with ` +
		"`kubectl --context cluster-a --namespace cnrm-system scale statefulset cnrm-controller-manager --replicas=1`"

	t.Run("after the target took over", func(t *testing.T) {
		fleet := newHAFleet(t, "cluster-a", "cluster-b")
		a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
		a.rejectPatch = rejectStart

		_, err := fleet.failOver("cluster-b")

		requireErrorContains(t, err, `cluster "cluster-b" is the leader, but not every manager on the other clusters was started again: `+wantStartErr)
		stopped := a.manager("cnrm-system", "cnrm-controller-manager")
		if got := *stopped.Spec.Replicas; got != 0 {
			t.Errorf("the manager has %d replicas, want 0", got)
		}
		if got, want := stopped.Annotations[failoverReplicasAnnotation], "1"; got != want {
			t.Errorf("the manager has the annotation %s=%q, want %q", failoverReplicasAnnotation, got, want)
		}

		// The next run starts it.
		a.rejectPatch = nil
		out, err := fleet.failOver("cluster-b")
		if err != nil {
			t.Fatalf("the second failover failed: %v\noutput:\n%s", err, out)
		}
		fleet.requireLeader(b)
		a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	})

	t.Run("after the target did not take over", func(t *testing.T) {
		fleet := newHAFleet(t, "cluster-a", "cluster-b")
		a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
		a.rejectPatch = rejectStart
		b.cannotLead = true

		_, err := fleet.failOverUntilTimeout("cluster-b")

		requireErrorContains(t, err,
			`failover to cluster "cluster-b" did not complete: cluster "cluster-b" did not become the leader: timed out after 500ms`,
			wantStartErr,
		)
		if strings.Contains(err.Error(), "were started again") {
			t.Errorf("the error says that the managers were started again:\n%v", err)
		}
	})
}

func TestFailoverFinishesAnInterruptedFailover(t *testing.T) {
	t.Run("the lease has not moved yet", func(t *testing.T) {
		fleet := newHAFleet(t, "cluster-a", "cluster-b")
		a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
		a.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager")
		for i := 0; i < fleet.podStopSteps; i++ {
			fleet.step()
		}
		if fleet.holder != a || a.running() {
			t.Fatalf("want cluster-a to hold the lease with its manager stopped; holder is %q, running clusters are %v", fleet.holder.name, fleet.runningClusters())
		}

		out, err := fleet.failOver("cluster-b")
		if err != nil {
			t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
		}

		fleet.requireLeader(b)
		a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
		fleet.requirePatches(`start cluster-a ` + fleetManager + ` (replicas 1, annotation "") while cluster-b leads`)
	})

	t.Run("the target has taken the lease", func(t *testing.T) {
		fleet := newHAFleet(t, "cluster-a", "cluster-b")
		a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
		a.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager")
		fleet.settle()
		if fleet.holder != b {
			t.Fatalf("want cluster-b to hold the lease, holder is %q", fleet.holder.name)
		}

		out, err := fleet.failOver("cluster-b")
		if err != nil {
			t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
		}

		fleet.requireLeader(b)
		a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
		fleet.requirePatches(`start cluster-a ` + fleetManager + ` (replicas 1, annotation "") while cluster-b leads`)
		want := `Leader: cluster "cluster-b" (projects/test-project/locations/us-central1/clusters/cluster-b)
Cluster "cluster-b" already holds the lease. Starting the managers that an earlier failover stopped.
cluster "cluster-a": started Config Connector's manager again (StatefulSet cnrm-system/cnrm-controller-manager, scaled from 0 to 1)
cluster "cluster-b": waiting up to 10s for it to become the leader
cluster "cluster-b": is the leader
cluster "cluster-a": waiting up to 10s for it to follow the new leader
cluster "cluster-a": follows the new leader
Failover complete: cluster "cluster-b" (projects/test-project/locations/us-central1/clusters/cluster-b) is the leader.
`
		if out != want {
			t.Errorf("got output:\n%s\nwant:\n%s", out, want)
		}
	})

	// The target of the interrupted failover never took the lease. Failing over to the old
	// leader starts its managers again, and it renews the lease that it still holds.
	t.Run("back to the old leader", func(t *testing.T) {
		fleet := newHAFleet(t, "cluster-a", "cluster-b")
		a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
		b.cannotLead = true
		a.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager")
		fleet.settle()
		if fleet.holder != a || a.running() {
			t.Fatalf("want cluster-a to hold the lease with its manager stopped; holder is %q, running clusters are %v", fleet.holder.name, fleet.runningClusters())
		}

		out, err := fleet.failOver("cluster-a")
		if err != nil {
			t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
		}

		fleet.requireLeader(a)
		a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
		fleet.requirePatches(`start cluster-a ` + fleetManager + ` (replicas 1, annotation "") while cluster-a leads`)
	})

	// The target's own managers are stopped, because an earlier failover to another cluster was
	// interrupted. They are started before the leader is stopped.
	t.Run("the target's managers are stopped", func(t *testing.T) {
		fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
		a, b, c := fleet.cluster("cluster-a"), fleet.cluster("cluster-b"), fleet.cluster("cluster-c")
		c.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager")
		fleet.settle()

		out, err := fleet.failOver("cluster-c")
		if err != nil {
			t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
		}

		fleet.requireLeader(c)
		for _, cluster := range []*fleetCluster{a, b, c} {
			cluster.requireManager("cnrm-system", "cnrm-controller-manager", 1)
		}
		fleet.requirePatches(
			`start cluster-c `+fleetManager+` (replicas 1, annotation "") while cluster-a leads`,
			`stop cluster-b `+fleetManager+` (annotation "1") while cluster-a leads and cluster-a,cluster-b,cluster-c run`,
			`stop cluster-a `+fleetManager+` (annotation "1") while cluster-a leads and cluster-a,cluster-c run`,
			`start cluster-a `+fleetManager+` (replicas 1, annotation "") while cluster-c leads`,
			`start cluster-b `+fleetManager+` (replicas 1, annotation "") while cluster-c leads`,
		)
	})
}

// One StatefulSet that can't be started doesn't keep the others stopped.
func TestFailoverStartsTheOtherManagersWhenOneAnnotationIsInvalid(t *testing.T) {
	const wantErr = `cluster "cluster-a": StatefulSet cnrm-system/cnrm-controller-manager-team-a has the annotation cnrm.cloud.google.com/ha-failover-replicas="one", which is not a number of replicas`
	breakAnnotation := func(a *fleetCluster) {
		broken := a.manager("cnrm-system", "cnrm-controller-manager-team-a")
		broken.Annotations[failoverReplicasAnnotation] = "one"
		a.scaleManager(broken, 0)
	}

	t.Run("on a cluster other than the target", func(t *testing.T) {
		fleet := newHAFleet(t, "cluster-a", "cluster-b")
		a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
		for _, c := range fleet.clusters {
			c.useNamespacedMode(map[string]string{"team-a": "cnrm-system", "team-b": "cnrm-system"})
		}
		fleet.settle()
		a.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager-team-b")
		breakAnnotation(a)
		fleet.settle()
		if fleet.holder != b {
			t.Fatalf("want cluster-b to hold the lease, holder is %q", fleet.holder.name)
		}

		_, err := fleet.failOver("cluster-b")

		requireErrorContains(t, err, wantErr)
		a.requireManager("cnrm-system", "cnrm-controller-manager-team-b", 1)
		if got := *a.manager("cnrm-system", "cnrm-controller-manager-team-a").Spec.Replicas; got != 0 {
			t.Errorf("the StatefulSet with the invalid annotation has %d replicas, want it left at 0", got)
		}
	})

	// An earlier failover from cluster-a stopped cluster-a and cluster-b, and its target didn't
	// take the lease. Failing over back to cluster-a starts cluster-b's managers, even though
	// one of cluster-a's can't be started.
	t.Run("on the target", func(t *testing.T) {
		fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
		a, b, c := fleet.cluster("cluster-a"), fleet.cluster("cluster-b"), fleet.cluster("cluster-c")
		c.cannotLead = true
		for _, cluster := range fleet.clusters {
			cluster.useNamespacedMode(map[string]string{"team-a": "cnrm-system"})
		}
		fleet.settle()
		b.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager-team-a")
		breakAnnotation(a)
		fleet.settle()
		if fleet.holder != a || a.running() || b.running() {
			t.Fatalf("want cluster-a to hold the lease with the managers of cluster-a and cluster-b stopped; holder is %q, running clusters are %v", fleet.holder.name, fleet.runningClusters())
		}

		_, err := fleet.failOver("cluster-a")

		requireErrorContains(t, err, wantErr)
		b.requireManager("cnrm-system", "cnrm-controller-manager-team-a", 1)
	})
}

// Someone started a manager by hand after an interrupted failover. It keeps the replicas that it
// has, and only loses the annotation.
func TestFailoverLeavesAManagerRunningThatSomeoneStartedAgain(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
	for _, c := range fleet.clusters {
		c.useNamespacedMode(map[string]string{"team-a": "cnrm-system", "team-b": "cnrm-system"})
	}
	fleet.settle()
	a.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager-team-a")
	a.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager-team-b")
	fleet.settle()
	if fleet.holder != b {
		t.Fatalf("want cluster-b to hold the lease, holder is %q", fleet.holder.name)
	}
	startedByHand := a.manager("cnrm-system", "cnrm-controller-manager-team-b")
	startedByHand.Annotations[failoverReplicasAnnotation] = "3"
	a.scaleManager(startedByHand, 1)
	fleet.settle()

	out, err := fleet.failOver("cluster-b")
	if err != nil {
		t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
	}

	fleet.requireLeader(b)
	a.requireManager("cnrm-system", "cnrm-controller-manager-team-a", 1)
	a.requireManager("cnrm-system", "cnrm-controller-manager-team-b", 1)
	if got := strings.Count(out, "started Config Connector's manager again"); got != 1 {
		t.Errorf("the output reports %d managers started again, want only team-a's:\n%s", got, out)
	}
}

// The target held the lease when its managers were started again, but another cluster with
// running managers got the lease first. That cluster keeps it, so failover doesn't wait.
func TestFailoverFailsWhenARunningClusterTakesTheLeaseBack(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b")
	a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
	b.cannotLead = true
	a.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager")
	fleet.settle()
	if fleet.holder != a || a.running() {
		t.Fatalf("want cluster-a to hold the lease with its manager stopped; holder is %q, running clusters are %v", fleet.holder.name, fleet.runningClusters())
	}
	fleet.onStep = func() {
		if len(fleet.patches) == 1 && fleet.holder == a {
			fleet.giveLease(b)
		}
	}

	_, err := fleet.failOver("cluster-a")

	requireErrorContains(t, err, `cluster "cluster-a" did not become the leader: cluster "cluster-b" took the lease instead`)
	a.requireManager("cnrm-system", "cnrm-controller-manager", 1)
	fleet.requirePatches(`start cluster-a ` + fleetManager + ` (replicas 1, annotation "") while cluster-a leads`)
}

// Failover checks that the target can take over before it stops anything.
func TestFailoverRefusesBeforeItChangesAnything(t *testing.T) {
	other := operatorv1beta1.GKECluster{Project: "test-project", Location: "us-central1", Name: "cluster-z"}

	tests := []struct {
		name    string
		prepare func(t *testing.T, fleet *haFleet, b *fleetCluster)
		wantErr string
	}{
		{
			name: "the target's manager is not ready",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				b.unreadyPods = true
				fleet.settle()
			},
			wantErr: `cluster "cluster-b" is not ready to take over: StatefulSet ` + fleetManager + ` has 0 of 1 pods ready`,
		},
		{
			name: "the target has no manager",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				if err := b.backend.Delete(context.Background(), b.manager("cnrm-system", "cnrm-controller-manager")); err != nil {
					t.Fatalf("deleting the manager: %v", err)
				}
			},
			wantErr: `cluster "cluster-b" is not ready to take over: it has no StatefulSet of Config Connector's manager`,
		},
		{
			name: "someone scaled the target's manager to 0",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				b.scaleManager(b.manager("cnrm-system", "cnrm-controller-manager"), 0)
				fleet.settle()
			},
			wantErr: `cluster "cluster-b" is not ready to take over: none of its Config Connector managers runs (StatefulSet ` + fleetManager + ` scaled to 0)`,
		},
		{
			name: "the target reports no leader",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				b.staleLease = true
				b.setLeaseStatus("", "")
			},
			wantErr: `cannot tell which cluster is the leader: timed out after 500ms: cluster "cluster-b" doesn't report a leader: the status of its MultiClusterLease cnrm-system/kcc-multicluster-lease names no holder`,
		},
		{
			name: "the clusters disagree about the leader",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				b.staleLease = true
				b.setLeaseStatus(other.ResourceName(), fleet.renewTime())
			},
			wantErr: `cannot tell which cluster is the leader: timed out after 500ms: the clusters disagree: cluster "cluster-a" reports ` + fleetIdentity("cluster-a") + `; cluster "cluster-b" reports ` + fleetIdentity("cluster-z"),
		},
		{
			name: "the target's MCL controller has stopped reading the lease",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				b.staleLease = true
				// The leader renews 100 more times, one second apart.
				for i := 0; i < 100; i++ {
					fleet.step()
				}
			},
			wantErr: `cluster "cluster-b" is not ready to take over: its view of the lease is 1m4`,
		},
		{
			name: "the target does not report when the lease was renewed",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				b.staleLease = true
				b.setLeaseStatus(fleetIdentity("cluster-a"), "not a time")
			},
			wantErr: `cluster "cluster-b" doesn't report when the lease was last renewed`,
		},
		{
			name: "the target's KRMSyncer pulls from another cluster",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				b.abandonedSyncers = true
				b.setSyncerSpec(fleetLeaseName, newSyncerSpec(other, false))
			},
			wantErr: `cluster "cluster-b" is not ready to take over: KRMSyncer cnrm-system/kcc-multicluster-lease does not name ` + fleetIdentity("cluster-a") + ` as its remote cluster yet`,
		},
		{
			name: "the target's KRMSyncer pulls from a cluster with the leader's name in another project",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				b.abandonedSyncers = true
				b.setSyncerSpec(fleetLeaseName, newSyncerSpec(operatorv1beta1.GKECluster{Project: "another-project", Location: "us-central1", Name: "cluster-a"}, false))
			},
			wantErr: `cluster "cluster-b" is not ready to take over: KRMSyncer cnrm-system/kcc-multicluster-lease does not name ` + fleetIdentity("cluster-a") + ` as its remote cluster yet`,
		},
		{
			name: "the target's KRMSyncer pulls from a cluster with the leader's name in another location",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				b.abandonedSyncers = true
				b.setSyncerSpec(fleetLeaseName, newSyncerSpec(operatorv1beta1.GKECluster{Project: "test-project", Location: "europe-west1", Name: "cluster-a"}, false))
			},
			wantErr: `cluster "cluster-b" is not ready to take over: KRMSyncer cnrm-system/kcc-multicluster-lease does not name ` + fleetIdentity("cluster-a") + ` as its remote cluster yet`,
		},
		{
			name: "the target's KRMSyncer is suspended",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				b.abandonedSyncers = true
				b.setSyncerSpec(fleetLeaseName, newSyncerSpec(fleet.cluster("cluster-a").gke, true))
			},
			wantErr: `cluster "cluster-b" is not ready to take over: KRMSyncer cnrm-system/kcc-multicluster-lease is still suspended`,
		},
		{
			name: "the target's manager has no KRMSyncer",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				b.abandonedSyncers = true
				if err := b.backend.Delete(context.Background(), b.syncer(fleetLeaseName)); err != nil {
					t.Fatalf("deleting the KRMSyncer: %v", err)
				}
			},
			wantErr: `cluster "cluster-b" is not ready to take over: KRMSyncer cnrm-system/kcc-multicluster-lease, which the manager in StatefulSet ` + fleetManager + ` maintains, does not exist yet`,
		},
		{
			name: "multi-cluster leader election is not configured on a cluster",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				b.updateConfigConnector(func(cc *operatorv1beta1.ConfigConnector) { cc.Spec.Experiments = nil })
			},
			wantErr: `cluster "cluster-b" does not have spec.experiments.multiClusterLease configured`,
		},
		{
			name: "the clusters use different leases",
			prepare: func(t *testing.T, fleet *haFleet, b *fleetCluster) {
				b.updateConfigConnector(func(cc *operatorv1beta1.ConfigConnector) {
					cc.Spec.Experiments.MultiClusterLease.LeaseName = "another-lease"
				})
			},
			wantErr: `mismatched leaseName across clusters (cluster-a="kcc-multicluster-lease", cluster-b="another-lease")`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fleet := newHAFleet(t, "cluster-a", "cluster-b")
			a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
			tc.prepare(t, fleet, b)

			out, err := fleet.failOver("cluster-b", shortAgreement)

			requireErrorContains(t, err, tc.wantErr)
			fleet.requirePatches()
			if fleet.holder != a {
				t.Errorf("cluster %q holds the lease, want cluster-a still", fleet.holder.name)
			}
			if strings.Contains(out, "stopped") {
				t.Errorf("the output says that something was stopped:\n%s", out)
			}
		})
	}
}

// Failover has to stop the leader's managers, so the leader must be one of the clusters.
func TestFailoverRefusesWhenTheLeaderIsNotOneOfTheClusters(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
	b, c := fleet.cluster("cluster-b"), fleet.cluster("cluster-c")

	_, err := fleet.failOverOn(context.Background(), []*clusterHandle{b.handle, c.handle}, "cluster-c")

	requireErrorContains(t, err, "the leader, "+fleetIdentity("cluster-a")+", is not one of the clusters in --contexts")
	fleet.requirePatches()
}

// A dry run reports what a real run would refuse, and what it would start first.
func TestFailoverDryRunOfAnInterruptedFailover(t *testing.T) {
	t.Run("the target holds the lease", func(t *testing.T) {
		fleet := newHAFleet(t, "cluster-a", "cluster-b")
		a := fleet.cluster("cluster-a")
		a.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager")
		fleet.settle()

		out, err := fleet.failOver("cluster-b", func(opts *failoverOptions) { opts.dryRun = true })
		if err != nil {
			t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
		}

		fleet.requirePatches()
		want := `[dry-run] cluster "cluster-a": would start StatefulSet cnrm-system/cnrm-controller-manager, which an earlier failover stopped`
		if !strings.Contains(out, want) {
			t.Errorf("got output:\n%s\nwant it to contain:\n%s", out, want)
		}
	})

	t.Run("the target's managers are stopped", func(t *testing.T) {
		fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
		a, c := fleet.cluster("cluster-a"), fleet.cluster("cluster-c")
		c.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager")
		fleet.settle()

		out, err := fleet.failOver("cluster-c", func(opts *failoverOptions) { opts.dryRun = true })
		if err != nil {
			t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
		}

		fleet.requirePatches()
		fleet.requireLeader(a)
		want := `[dry-run] cluster "cluster-c": would start StatefulSet cnrm-system/cnrm-controller-manager, which an earlier failover stopped, and wait for the cluster to follow the leader
[dry-run] cluster "cluster-b": would stop StatefulSet cnrm-system/cnrm-controller-manager, and wait up to 10s for its pods to stop
[dry-run] cluster "cluster-a": would stop StatefulSet cnrm-system/cnrm-controller-manager, and wait up to 10s for its pods to stop
`
		if !strings.Contains(out, want) {
			t.Errorf("got output:\n%s\nwant it to contain:\n%s", out, want)
		}
	})
}

// An earlier failover left the target's managers stopped. Failover starts them, but doesn't stop
// the leader if they don't come back.
func TestFailoverRefusesWhenTheManagersOfTheTargetDoNotComeBack(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
	a, c := fleet.cluster("cluster-a"), fleet.cluster("cluster-c")
	c.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager")
	fleet.settle()
	c.unreadyPods = true

	_, err := fleet.failOverUntilTimeout("cluster-c")

	requireErrorContains(t, err, `cluster "cluster-c" is not ready to take over: timed out after 500ms: StatefulSet `+fleetManager+` has 0 of 1 pods ready`)
	fleet.requirePatches(`start cluster-c ` + fleetManager + ` (replicas 1, annotation "") while cluster-a leads`)
	fleet.requireLeader(a)
}

// An earlier failover was killed after it had stopped the leader and the cluster that is the
// target now. The leader's lease expires while the target's managers start, so what failover
// has read about the leader no longer holds.
func TestFailoverGivesUpWhenTheLeaseMovesWhileTheManagersOfTheTargetStart(t *testing.T) {
	fleet := newHAFleet(t, "cluster-a", "cluster-b", "cluster-c")
	a, b, c := fleet.cluster("cluster-a"), fleet.cluster("cluster-b"), fleet.cluster("cluster-c")
	// The lease lasts until after failover has started the target's managers.
	fleet.takeOverSteps = 60
	a.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager")
	b.stopManagerAsFailoverDoes("cnrm-system", "cnrm-controller-manager")
	for i := 0; i < fleet.podStopSteps; i++ {
		fleet.step()
	}
	if fleet.holder != a || a.running() || b.running() {
		t.Fatalf("want cluster-a to hold the lease with the managers of cluster-a and cluster-b stopped; holder is %q, running clusters are %v", fleet.holder.name, fleet.runningClusters())
	}
	// The target's managers take long to become ready, and cluster-c takes the lease first.
	b.unreadyPods = true
	b.cannotLead = true

	_, err := fleet.failOver("cluster-b")

	want := "the lease went from " + fleetIdentity("cluster-a") + " to " + fleetIdentity("cluster-c") + ` while failover started the managers of cluster "cluster-b"; run the command again`
	if err == nil || err.Error() != want {
		t.Errorf("got error:\n  %v\nwant:\n  %s", err, want)
	}
	if fleet.holder != c {
		t.Errorf("cluster %q holds the lease, want cluster-c", fleet.holder.name)
	}
	// The target's managers stay started. Nothing else was touched.
	fleet.requirePatches(`start cluster-b ` + fleetManager + ` (replicas 1, annotation "") while cluster-a leads`)
}

func (c *fleetCluster) syncer(name string) *syncerv1alpha1.KRMSyncer {
	c.fleet.t.Helper()
	syncer := &syncerv1alpha1.KRMSyncer{}
	if err := c.backend.Get(context.Background(), types.NamespacedName{Namespace: fleetLeaseNamespace, Name: name}, syncer); err != nil {
		c.fleet.t.Fatalf("cluster %q: getting KRMSyncer %s: %v", c.name, name, err)
	}
	return syncer
}

func (c *fleetCluster) setSyncerSpec(name string, spec syncerv1alpha1.KRMSyncerSpec) {
	c.fleet.t.Helper()
	syncer := c.syncer(name)
	syncer.Spec = spec
	if err := c.backend.Update(context.Background(), syncer); err != nil {
		c.fleet.t.Fatalf("cluster %q: updating KRMSyncer %s: %v", c.name, name, err)
	}
}

func (c *fleetCluster) updateConfigConnector(update func(*operatorv1beta1.ConfigConnector)) {
	c.fleet.t.Helper()
	cc := &operatorv1beta1.ConfigConnector{}
	if err := c.backend.Get(context.Background(), types.NamespacedName{Name: configConnectorResourceName}, cc); err != nil {
		c.fleet.t.Fatalf("cluster %q: getting the ConfigConnector: %v", c.name, err)
	}
	update(cc)
	if err := c.backend.Update(context.Background(), cc); err != nil {
		c.fleet.t.Fatalf("cluster %q: updating the ConfigConnector: %v", c.name, err)
	}
}

// The StatefulSet controller reports a spec that it has processed, and the pods that are left.
func TestManagerStatefulSetWaitsForTheStatefulSetController(t *testing.T) {
	const key = "cnrm-system/cnrm-controller-manager"
	tests := []struct {
		name     string
		replicas int32
		// generation and status are what the API server and the StatefulSet controller report.
		generation  int64
		status      appsv1.StatefulSetStatus
		wantReady   string
		wantStopped string
	}{
		{
			name:        "running and ready",
			replicas:    1,
			generation:  2,
			status:      appsv1.StatefulSetStatus{ObservedGeneration: 2, Replicas: 1, ReadyReplicas: 1},
			wantStopped: "StatefulSet " + key + " still has pods (1)",
		},
		{
			name:        "running, not ready yet",
			replicas:    1,
			generation:  2,
			status:      appsv1.StatefulSetStatus{ObservedGeneration: 2, Replicas: 1},
			wantReady:   "StatefulSet " + key + " has 0 of 1 pods ready",
			wantStopped: "StatefulSet " + key + " still has pods (1)",
		},
		{
			// The status is from before the manager was scaled to 0.
			name:        "scaled to 0, not processed yet",
			replicas:    0,
			generation:  3,
			status:      appsv1.StatefulSetStatus{ObservedGeneration: 2},
			wantReady:   "the StatefulSet controller has not processed the latest spec of StatefulSet " + key + " yet",
			wantStopped: "the StatefulSet controller has not processed the latest spec of StatefulSet " + key + " yet",
		},
		{
			name:        "scaled to 0, the pod terminates",
			replicas:    0,
			generation:  3,
			status:      appsv1.StatefulSetStatus{ObservedGeneration: 3, Replicas: 1},
			wantStopped: "StatefulSet " + key + " still has pods (1)",
		},
		{
			name:       "scaled to 0, the pod is gone",
			replicas:   0,
			generation: 3,
			status:     appsv1.StatefulSetStatus{ObservedGeneration: 3},
		},
		{
			// The status is from when the manager was at 0.
			name:        "scaled to 1, not processed yet",
			replicas:    1,
			generation:  4,
			status:      appsv1.StatefulSetStatus{ObservedGeneration: 3, ReadyReplicas: 1},
			wantReady:   "the StatefulSet controller has not processed the latest spec of StatefulSet " + key + " yet",
			wantStopped: "the StatefulSet controller has not processed the latest spec of StatefulSet " + key + " yet",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sts := newManagerStatefulSet("cnrm-system", "cnrm-controller-manager", "")
			sts.Spec.Replicas = &tc.replicas
			sts.Generation = tc.generation
			sts.Status = tc.status
			m := &managerStatefulSet{StatefulSet: sts}

			if got := m.pendingReady(); got != tc.wantReady {
				t.Errorf("pendingReady() = %q, want %q", got, tc.wantReady)
			}
			if got := m.pendingStopped(); got != tc.wantStopped {
				t.Errorf("pendingStopped() = %q, want %q", got, tc.wantStopped)
			}
		})
	}
}

func TestPoll(t *testing.T) {
	const (
		// timeout is for the calls that are to run out of time, and long is for the others.
		timeout = 50 * time.Millisecond
		long    = 10 * time.Second
	)

	t.Run("returns when nothing is pending", func(t *testing.T) {
		calls := 0
		err := poll(context.Background(), time.Millisecond, long, func(context.Context) (string, error) {
			calls++
			if calls < 3 {
				return "the thing", nil
			}
			return "", nil
		})
		if err != nil || calls != 3 {
			t.Errorf("got error %v after %d calls, want none after 3", err, calls)
		}
	})

	t.Run("reports what was pending when the time runs out", func(t *testing.T) {
		err := poll(context.Background(), time.Millisecond, timeout, func(context.Context) (string, error) {
			return "the thing", nil
		})
		if want := "timed out after 50ms: the thing"; err == nil || err.Error() != want {
			t.Errorf("got error %v, want %q", err, want)
		}
	})

	// The request that is in flight when the time runs out fails. That isn't what was pending.
	t.Run("reports what was pending before a request that the time cut short", func(t *testing.T) {
		calls := 0
		err := poll(context.Background(), time.Millisecond, timeout, func(ctx context.Context) (string, error) {
			calls++
			if calls == 1 {
				return "the thing", nil
			}
			<-ctx.Done()
			return "request failed: " + ctx.Err().Error(), nil
		})
		if want := "timed out after 50ms: the thing"; err == nil || err.Error() != want {
			t.Errorf("got error %v, want %q", err, want)
		}
	})

	t.Run("reports the request that the time cut short if nothing else was pending", func(t *testing.T) {
		err := poll(context.Background(), time.Millisecond, timeout, func(ctx context.Context) (string, error) {
			<-ctx.Done()
			return "request failed: " + ctx.Err().Error(), nil
		})
		if want := "timed out after 50ms: request failed: context deadline exceeded"; err == nil || err.Error() != want {
			t.Errorf("got error %v, want %q", err, want)
		}
	})

	t.Run("stops at once on an error", func(t *testing.T) {
		calls := 0
		err := poll(context.Background(), time.Millisecond, long, func(context.Context) (string, error) {
			calls++
			return "", errors.New("it broke")
		})
		if err == nil || err.Error() != "it broke" || calls != 1 {
			t.Errorf("got error %v after %d calls, want \"it broke\" after 1", err, calls)
		}
	})

	t.Run("reports an interruption", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		err := poll(ctx, time.Millisecond, long, func(context.Context) (string, error) {
			cancel()
			return "the thing", nil
		})
		if want := "interrupted: context canceled"; err == nil || err.Error() != want {
			t.Errorf("got error %v, want %q", err, want)
		}
	})
}
