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
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func withFakeClock(opts *setupOptions, startTime time.Time) {
	curr := startTime
	opts.nowFunc = func() time.Time { return curr }
	opts.sleepFunc = func(ctx context.Context, d time.Duration) error {
		curr = curr.Add(d)
		return nil
	}
}

func newTestHealthyStatusReport(now time.Time) *statusReport {
	ccA := &operatorv1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{Name: "configconnector.core.cnrm.cloud.google.com"},
		Spec: operatorv1beta1.ConfigConnectorSpec{
			Experiments: &operatorv1beta1.CCExperiments{
				MultiClusterLease: &operatorv1beta1.MultiClusterLeaseSpec{
					GKECluster: &operatorv1beta1.GKECluster{
						Project:  "proj-a",
						Location: "us-central1",
						Name:     "cluster-a",
					},
					LeaseName:               testLeaseName,
					Namespace:               testLeaseNamespace,
					ResourceReplicationMode: "Status",
					LeaderEndpointType:      "Default",
				},
			},
		},
	}
	ccB := &operatorv1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{Name: "configconnector.core.cnrm.cloud.google.com"},
		Spec: operatorv1beta1.ConfigConnectorSpec{
			Experiments: &operatorv1beta1.CCExperiments{
				MultiClusterLease: &operatorv1beta1.MultiClusterLeaseSpec{
					GKECluster: &operatorv1beta1.GKECluster{
						Project:  "proj-b",
						Location: "us-east1",
						Name:     "cluster-b",
					},
					LeaseName:               testLeaseName,
					Namespace:               testLeaseNamespace,
					ResourceReplicationMode: "Status",
					LeaderEndpointType:      "Default",
				},
			},
		},
	}
	renewTime := now.Format(time.RFC3339)
	holder := "projects/proj-a/locations/us-central1/clusters/cluster-a"
	lease := newTestLease(holder, renewTime)

	return &statusReport{
		state: haStateHealthy,
		clusters: []*clusterStatus{
			{
				gke: gkeClusterInfo{contextName: "gke_proj-a_us-central1_cluster-a"},
				state: &clusterHAState{
					contextName:       "gke_proj-a_us-central1_cluster-a",
					configConnector:   ccA,
					multiClusterLease: lease,
				},
			},
			{
				gke: gkeClusterInfo{contextName: "gke_proj-b_us-east1_cluster-b"},
				state: &clusterHAState{
					contextName:       "gke_proj-b_us-east1_cluster-b",
					configConnector:   ccB,
					multiClusterLease: lease,
				},
			},
		},
	}
}

func TestWaitHealthyFirstPoll(t *testing.T) {
	holder := "projects/proj-a/locations/us-central1/clusters/cluster-a"
	leaseA := newTestLease(holder, testRenewTime)
	syncerA := newTestSyncer(testLeaseName, holder, true, syncerv1alpha1.ReasonSuspended, "Syncer is suspended")
	leaseB := newTestLease(holder, testRenewTime)
	syncerB := newTestSyncer(testLeaseName, holder, false, syncerv1alpha1.ReasonSyncing, "Syncing")

	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector(), leaseA, syncerA),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector(), leaseB, syncerB),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.wait = true
	opts.waitTimeout = 5 * time.Minute
	withFakeClock(&opts, testNow)
	var stdout bytes.Buffer
	opts.stdout = &stdout

	err := executeSetup(context.Background(), opts, handles, runner)
	if err != nil {
		t.Fatalf("expected executeSetup to succeed, got: %v", err)
	}

	out := stdout.String()
	wantStartMsg := "Waiting up to 5m0s for multi-cluster HA to become healthy (use --wait=false to skip)...\n"
	if !strings.Contains(out, wantStartMsg) {
		t.Fatalf("expected stdout to contain start message %q, got:\n%s", wantStartMsg, out)
	}

	wantSummary := "Multi-cluster HA is healthy.\n" +
		"  Leader:      gke_proj-a_us-central1_cluster-a\n" +
		"  Standbys:    gke_proj-b_us-east1_cluster-b\n" +
		"  Replication: Status\n"
	if !strings.Contains(out, wantSummary) {
		t.Fatalf("expected stdout to contain summary:\n%s\ngot:\n%s", wantSummary, out)
	}
}

func TestWaitHealthyLaterPoll(t *testing.T) {
	holder := "projects/proj-a/locations/us-central1/clusters/cluster-a"
	leaseA := newTestLease("", "")
	syncerA := newTestSyncer(testLeaseName, holder, true, syncerv1alpha1.ReasonSuspended, "Syncer is suspended")
	leaseB := newTestLease("", "")
	syncerB := newTestSyncer(testLeaseName, holder, false, syncerv1alpha1.ReasonSyncing, "Syncing")

	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector(), leaseA, syncerA),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector(), leaseB, syncerB),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.wait = true
	opts.waitTimeout = 1 * time.Minute
	opts.waitPollInterval = 5 * time.Second

	currTime := testNow
	opts.nowFunc = func() time.Time { return currTime }
	sleepCount := 0
	opts.sleepFunc = func(ctx context.Context, d time.Duration) error {
		currTime = currTime.Add(d)
		sleepCount++
		if sleepCount == 2 {
			renewTime := testRenewTime
			leaseA.Status.GlobalHolderIdentity = &holder
			leaseA.Status.GlobalRenewTime = &renewTime
			if err := handles[0].client.Update(ctx, leaseA); err != nil {
				t.Fatalf("updating lease on cluster A: %v", err)
			}
			leaseB.Status.GlobalHolderIdentity = &holder
			leaseB.Status.GlobalRenewTime = &renewTime
			if err := handles[1].client.Update(ctx, leaseB); err != nil {
				t.Fatalf("updating lease on cluster B: %v", err)
			}
		}
		return nil
	}

	var stdout bytes.Buffer
	opts.stdout = &stdout

	err := executeSetup(context.Background(), opts, handles, runner)
	if err != nil {
		t.Fatalf("expected executeSetup to succeed on later poll, got: %v", err)
	}
	if sleepCount != 2 {
		t.Fatalf("expected 2 sleeps (3 polls), got %d sleeps", sleepCount)
	}

	out := stdout.String()
	wantSummary := "Multi-cluster HA is healthy.\n" +
		"  Leader:      gke_proj-a_us-central1_cluster-a\n" +
		"  Standbys:    gke_proj-b_us-east1_cluster-b\n" +
		"  Replication: Status\n"
	if !strings.Contains(out, wantSummary) {
		t.Fatalf("expected stdout to contain summary:\n%s\ngot:\n%s", wantSummary, out)
	}
}

func TestWaitTimeoutTwoClustersDisagreeOnHolder(t *testing.T) {
	holderA := "projects/proj-a/locations/us-central1/clusters/cluster-a"
	holderB := "projects/proj-b/locations/us-east1/clusters/cluster-b"
	leaseA := newTestLease(holderA, testRenewTime)
	leaseB := newTestLease(holderB, testRenewTime)

	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector(), leaseA),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector(), leaseB),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.contexts = []string{handles[0].contextName, handles[1].contextName}
	opts.wait = true
	opts.waitTimeout = 10 * time.Second
	opts.waitPollInterval = 5 * time.Second
	withFakeClock(&opts, testNow)

	disagreeProblem := `clusters disagree on the lease holder ("projects/proj-a/locations/us-central1/clusters/cluster-a" reported by [gke_proj-a_us-central1_cluster-a]; "projects/proj-b/locations/us-east1/clusters/cluster-b" reported by [gke_proj-b_us-east1_cluster-b])`

	err := executeSetup(context.Background(), opts, handles, runner)
	if err == nil {
		t.Fatalf("expected executeSetup to fail on timeout, got nil")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, disagreeProblem) {
		t.Fatalf("expected error to list disagreement problem %q, got: %v", disagreeProblem, errMsg)
	}
	if !strings.Contains(errMsg, "The HA configuration is applied and was not rolled back.") {
		t.Fatalf("expected error to state configuration was applied and not rolled back, got: %v", errMsg)
	}
	wantCmd := "config-connector ha status --contexts=gke_proj-a_us-central1_cluster-a,gke_proj-b_us-east1_cluster-b"
	if !strings.Contains(errMsg, wantCmd) {
		t.Fatalf("expected error to provide status command %q, got: %v", wantCmd, errMsg)
	}
}

func TestWaitTimeoutNoClusterReportsHolder(t *testing.T) {
	leaseA := newTestLease("", "")
	leaseB := newTestLease("", "")

	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector(), leaseA),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector(), leaseB),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.contexts = []string{handles[0].contextName, handles[1].contextName}
	opts.wait = true
	opts.waitTimeout = 10 * time.Second
	opts.waitPollInterval = 5 * time.Second
	withFakeClock(&opts, testNow)

	noHolderProblem := `cluster "gke_proj-a_us-central1_cluster-a": no cluster holds the lease cnrm-system/kcc-multicluster-lease`

	err := executeSetup(context.Background(), opts, handles, runner)
	if err == nil {
		t.Fatalf("expected executeSetup to fail on timeout, got nil")
	}
	if !strings.Contains(err.Error(), noHolderProblem) {
		t.Fatalf("expected error to contain %q, got: %v", noHolderProblem, err)
	}
}

func TestWaitTimeoutStandbySyncsFromWrongCluster(t *testing.T) {
	holder := "projects/proj-a/locations/us-central1/clusters/cluster-a"
	leaseA := newTestLease(holder, testRenewTime)
	syncerA := newTestSyncer(testLeaseName, holder, true, syncerv1alpha1.ReasonSuspended, "Syncer is suspended")
	leaseB := newTestLease(holder, testRenewTime)
	wrongRemote := "projects/proj-c/locations/us-west1/clusters/cluster-c"
	syncerB := newTestSyncer(testLeaseName, wrongRemote, false, syncerv1alpha1.ReasonSyncing, "Syncing")

	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector(), leaseA, syncerA),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector(), leaseB, syncerB),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.contexts = []string{handles[0].contextName, handles[1].contextName}
	opts.wait = true
	opts.waitTimeout = 10 * time.Second
	opts.waitPollInterval = 5 * time.Second
	withFakeClock(&opts, testNow)

	wrongClusterProblem := `cluster "gke_proj-b_us-east1_cluster-b": KRMSyncer cnrm-system/kcc-multicluster-lease syncs from "projects/proj-c/locations/us-west1/clusters/cluster-c", but the lease is held by "projects/proj-a/locations/us-central1/clusters/cluster-a"`

	err := executeSetup(context.Background(), opts, handles, runner)
	if err == nil {
		t.Fatalf("expected executeSetup to fail on timeout, got nil")
	}
	if !strings.Contains(err.Error(), wrongClusterProblem) {
		t.Fatalf("expected error to contain wrong syncer problem, got: %v", err)
	}
}

func TestWaitTimeoutReportsLatestPollProblems(t *testing.T) {
	t.Run("updates lastReport when poll completes normally", func(t *testing.T) {
		handles := []*clusterHandle{
			newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
			newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
		}
		runner := &recordingCommandRunner{}
		opts := newTestSetupOptions()
		opts.contexts = []string{handles[0].contextName, handles[1].contextName}
		opts.wait = true
		opts.waitTimeout = 10 * time.Second
		opts.waitPollInterval = 5 * time.Second
		withFakeClock(&opts, testNow)

		pollCount := 0
		poll1Problem := "problem from poll 1: no cluster holds the lease"
		poll2Problem := "problem from poll 2: standby syncer setup failed"
		opts.statusFunc = func(ctx context.Context, h []*clusterHandle, now time.Time) *statusReport {
			pollCount++
			if pollCount == 1 {
				return &statusReport{
					state:    haStateError,
					problems: []string{poll1Problem},
				}
			}
			return &statusReport{
				state:    haStateError,
				problems: []string{poll2Problem},
			}
		}

		err := executeSetup(context.Background(), opts, handles, runner)
		if err == nil {
			t.Fatalf("expected executeSetup to fail on timeout, got nil")
		}
		if pollCount < 2 {
			t.Fatalf("expected at least 2 polls, got %d", pollCount)
		}
		errMsg := err.Error()
		if !strings.Contains(errMsg, poll2Problem) {
			t.Fatalf("expected error to contain poll 2 problem %q, got: %v", poll2Problem, errMsg)
		}
		if strings.Contains(errMsg, poll1Problem) {
			t.Fatalf("expected error not to contain stale poll 1 problem %q, got: %v", poll1Problem, errMsg)
		}
	})

	t.Run("preserves previous lastReport when poll context expires mid-poll", func(t *testing.T) {
		handles := []*clusterHandle{
			newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
			newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
		}
		runner := &recordingCommandRunner{}
		opts := newTestSetupOptions()
		opts.contexts = []string{handles[0].contextName, handles[1].contextName}
		opts.wait = true
		opts.waitTimeout = 20 * time.Millisecond
		opts.waitPollInterval = 10 * time.Millisecond

		pollCount := 0
		poll1Problem := "problem from poll 1: valid diagnosed error"
		opts.statusFunc = func(pollCtx context.Context, h []*clusterHandle, now time.Time) *statusReport {
			pollCount++
			if pollCount == 1 {
				return &statusReport{
					state:    haStateError,
					problems: []string{poll1Problem},
				}
			}
			<-pollCtx.Done()
			return &statusReport{
				state:    haStateError,
				problems: []string{"incomplete poll 2 report due to timeout"},
			}
		}

		err := executeSetup(context.Background(), opts, handles, runner)
		if err == nil {
			t.Fatalf("expected executeSetup to fail on timeout, got nil")
		}
		errMsg := err.Error()
		if !strings.Contains(errMsg, poll1Problem) {
			t.Fatalf("expected error to preserve poll 1 problem %q, got: %v", poll1Problem, errMsg)
		}
	})
}

func TestWaitTimeoutNamespacedModeHint(t *testing.T) {
	const wantHint = "the manager of each namespace restarts on its own, so a cluster with many namespaces can need a longer --wait-timeout"

	t.Run("cluster mode does not include namespaced hint", func(t *testing.T) {
		ccA := newBaseConfigConnector()
		ccA.Spec.Mode = "cluster"
		ccB := newBaseConfigConnector()
		ccB.Spec.Mode = "cluster"

		handles := []*clusterHandle{
			newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", ccA),
			newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", ccB),
		}
		runner := &recordingCommandRunner{}
		opts := newTestSetupOptions()
		opts.contexts = []string{handles[0].contextName, handles[1].contextName}
		opts.wait = true
		opts.waitTimeout = 10 * time.Second
		opts.waitPollInterval = 5 * time.Second
		withFakeClock(&opts, testNow)
		opts.statusFunc = func(ctx context.Context, h []*clusterHandle, now time.Time) *statusReport {
			return &statusReport{
				state:    haStateError,
				problems: []string{"lease not acquired"},
			}
		}

		err := executeSetup(context.Background(), opts, handles, runner)
		if err == nil {
			t.Fatalf("expected timeout error, got nil")
		}
		if strings.Contains(err.Error(), wantHint) {
			t.Fatalf("expected cluster mode error not to contain namespaced hint, got: %v", err)
		}
	})

	t.Run("namespaced mode includes namespaced hint", func(t *testing.T) {
		ccA := newBaseConfigConnector()
		ccA.Spec.Mode = "namespaced"
		ccB := newBaseConfigConnector()
		ccB.Spec.Mode = "cluster"

		handles := []*clusterHandle{
			newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", ccA),
			newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", ccB),
		}
		runner := &recordingCommandRunner{}
		opts := newTestSetupOptions()
		opts.contexts = []string{handles[0].contextName, handles[1].contextName}
		opts.wait = true
		opts.waitTimeout = 10 * time.Second
		opts.waitPollInterval = 5 * time.Second
		withFakeClock(&opts, testNow)
		opts.statusFunc = func(ctx context.Context, h []*clusterHandle, now time.Time) *statusReport {
			return &statusReport{
				state:    haStateError,
				problems: []string{"lease not acquired"},
			}
		}

		err := executeSetup(context.Background(), opts, handles, runner)
		if err == nil {
			t.Fatalf("expected timeout error, got nil")
		}
		if !strings.Contains(err.Error(), wantHint) {
			t.Fatalf("expected namespaced mode error to contain %q, got: %v", wantHint, err)
		}
	})

	t.Run("unset mode defaults to namespaced mode and includes namespaced hint", func(t *testing.T) {
		ccA := newBaseConfigConnector()
		ccA.Spec.Mode = ""
		ccB := newBaseConfigConnector()
		ccB.Spec.Mode = "cluster"

		handles := []*clusterHandle{
			newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", ccA),
			newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", ccB),
		}
		runner := &recordingCommandRunner{}
		opts := newTestSetupOptions()
		opts.contexts = []string{handles[0].contextName, handles[1].contextName}
		opts.wait = true
		opts.waitTimeout = 10 * time.Second
		opts.waitPollInterval = 5 * time.Second
		withFakeClock(&opts, testNow)
		opts.statusFunc = func(ctx context.Context, h []*clusterHandle, now time.Time) *statusReport {
			return &statusReport{
				state:    haStateError,
				problems: []string{"lease not acquired"},
			}
		}

		err := executeSetup(context.Background(), opts, handles, runner)
		if err == nil {
			t.Fatalf("expected timeout error, got nil")
		}
		if !strings.Contains(err.Error(), wantHint) {
			t.Fatalf("expected unset mode error to contain %q, got: %v", wantHint, err)
		}
	})
}

func TestWaitContextCanceledReturnsContextError(t *testing.T) {
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.contexts = []string{handles[0].contextName, handles[1].contextName}
	opts.wait = true
	opts.waitTimeout = 5 * time.Minute
	withFakeClock(&opts, testNow)

	ctx, cancel := context.WithCancel(context.Background())
	opts.statusFunc = func(pollCtx context.Context, h []*clusterHandle, now time.Time) *statusReport {
		cancel()
		return &statusReport{
			state:    haStateError,
			problems: []string{"context canceled mid-poll"},
		}
	}

	err := executeSetup(ctx, opts, handles, runner)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled error, got: %v", err)
	}
}

func TestWaitNoWaitCasesDoNotReadStatus(t *testing.T) {
	t.Run("wait=false prints ha status command and does not call statusFunc", func(t *testing.T) {
		handles := []*clusterHandle{
			newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
			newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
		}
		runner := &recordingCommandRunner{}
		opts := newTestSetupOptions()
		opts.contexts = []string{handles[0].contextName, handles[1].contextName}
		opts.wait = false
		var stdout bytes.Buffer
		opts.stdout = &stdout

		statusCalled := false
		opts.statusFunc = func(ctx context.Context, h []*clusterHandle, now time.Time) *statusReport {
			statusCalled = true
			return newTestHealthyStatusReport(now)
		}

		err := executeSetup(context.Background(), opts, handles, runner)
		if err != nil {
			t.Fatalf("expected executeSetup to succeed, got: %v", err)
		}
		if statusCalled {
			t.Fatalf("expected statusFunc not to be called when wait=false")
		}
		wantStatusCmd := "config-connector ha status --contexts=gke_proj-a_us-central1_cluster-a,gke_proj-b_us-east1_cluster-b\n"
		if !strings.Contains(stdout.String(), wantStatusCmd) {
			t.Fatalf("expected stdout to contain %q, got:\n%s", wantStatusCmd, stdout.String())
		}
	})

	t.Run("dry-run prints dry-run wait line and does not call statusFunc", func(t *testing.T) {
		handles := []*clusterHandle{
			newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
			newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
		}
		runner := &recordingCommandRunner{}
		opts := newTestSetupOptions()
		opts.contexts = []string{handles[0].contextName, handles[1].contextName}
		opts.dryRun = true
		opts.wait = true
		var stdout bytes.Buffer
		opts.stdout = &stdout

		statusCalled := false
		opts.statusFunc = func(ctx context.Context, h []*clusterHandle, now time.Time) *statusReport {
			statusCalled = true
			return newTestHealthyStatusReport(now)
		}

		err := executeSetup(context.Background(), opts, handles, runner)
		if err != nil {
			t.Fatalf("expected executeSetup to succeed, got: %v", err)
		}
		if statusCalled {
			t.Fatalf("expected statusFunc not to be called in dry-run mode")
		}
		wantDryRunLine := "[dry-run] would wait for multi-cluster HA to become healthy\n"
		if !strings.Contains(stdout.String(), wantDryRunLine) {
			t.Fatalf("expected stdout to contain %q, got:\n%s", wantDryRunLine, stdout.String())
		}
	})

	t.Run("yaml-only does not call statusFunc", func(t *testing.T) {
		handles := []*clusterHandle{
			newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
			newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
		}
		runner := &recordingCommandRunner{}
		opts := newTestSetupOptions()
		opts.contexts = []string{handles[0].contextName, handles[1].contextName}
		opts.yamlOnly = true
		opts.wait = true
		var stdout bytes.Buffer
		opts.stdout = &stdout

		statusCalled := false
		opts.statusFunc = func(ctx context.Context, h []*clusterHandle, now time.Time) *statusReport {
			statusCalled = true
			return newTestHealthyStatusReport(now)
		}

		err := executeSetup(context.Background(), opts, handles, runner)
		if err != nil {
			t.Fatalf("expected executeSetup to succeed, got: %v", err)
		}
		if statusCalled {
			t.Fatalf("expected statusFunc not to be called in yaml-only mode")
		}
	})
}

func TestWaitFlagsDefaultsAndValidation(t *testing.T) {
	// Defaults are true and 5m
	opts := newSetupOptions()
	if !opts.wait {
		t.Fatalf("expected default wait to be true, got %v", opts.wait)
	}
	if opts.waitTimeout != 5*time.Minute {
		t.Fatalf("expected default wait-timeout to be 5m, got %v", opts.waitTimeout)
	}

	// wait-timeout = 0 is rejected
	optsZero := newTestSetupOptions()
	optsZero.contexts = []string{"cluster-a", "cluster-b"}
	optsZero.waitTimeout = 0
	err := optsZero.validate()
	if err == nil || !strings.Contains(err.Error(), "--wait-timeout must be greater than 0; got 0s") {
		t.Fatalf("expected error rejecting 0 timeout, got: %v", err)
	}

	// negative wait-timeout is rejected
	optsNeg := newTestSetupOptions()
	optsNeg.contexts = []string{"cluster-a", "cluster-b"}
	optsNeg.waitTimeout = -5 * time.Second
	err = optsNeg.validate()
	if err == nil || !strings.Contains(err.Error(), "--wait-timeout must be greater than 0; got -5s") {
		t.Fatalf("expected error rejecting negative timeout, got: %v", err)
	}
}
