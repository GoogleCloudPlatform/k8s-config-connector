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
	"strings"
	"testing"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.wait = true
	opts.waitTimeout = 5 * time.Minute
	var stdout bytes.Buffer
	opts.stdout = &stdout

	pollCount := 0
	opts.statusFunc = func(ctx context.Context, h []*clusterHandle, now time.Time) *statusReport {
		pollCount++
		return newTestHealthyStatusReport(now)
	}

	err := executeSetup(context.Background(), opts, handles, runner)
	if err != nil {
		t.Fatalf("expected executeSetup to succeed, got: %v", err)
	}
	if pollCount != 1 {
		t.Fatalf("expected 1 poll, got %d", pollCount)
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
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.wait = true
	opts.waitTimeout = 1 * time.Second
	opts.waitPollInterval = 1 * time.Millisecond
	var stdout bytes.Buffer
	opts.stdout = &stdout

	pollCount := 0
	opts.statusFunc = func(ctx context.Context, h []*clusterHandle, now time.Time) *statusReport {
		pollCount++
		if pollCount < 3 {
			return &statusReport{
				state:    haStateError,
				problems: []string{"no cluster holds the lease"},
			}
		}
		return newTestHealthyStatusReport(now)
	}

	err := executeSetup(context.Background(), opts, handles, runner)
	if err != nil {
		t.Fatalf("expected executeSetup to succeed on later poll, got: %v", err)
	}
	if pollCount != 3 {
		t.Fatalf("expected 3 polls, got %d", pollCount)
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
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.contexts = []string{handles[0].contextName, handles[1].contextName}
	opts.wait = true
	opts.waitTimeout = 2 * time.Millisecond
	opts.waitPollInterval = 1 * time.Millisecond

	disagreeProblem := `clusters disagree on the lease holder ("projects/proj-a/locations/us-central1/clusters/cluster-a" reported by [gke_proj-a_us-central1_cluster-a]; "projects/proj-b/locations/us-east1/clusters/cluster-b" reported by [gke_proj-b_us-east1_cluster-b])`
	opts.statusFunc = func(ctx context.Context, h []*clusterHandle, now time.Time) *statusReport {
		return &statusReport{
			state:    haStateError,
			problems: []string{disagreeProblem},
		}
	}

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
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.contexts = []string{handles[0].contextName, handles[1].contextName}
	opts.wait = true
	opts.waitTimeout = 2 * time.Millisecond
	opts.waitPollInterval = 1 * time.Millisecond

	noHolderProblem := `cluster "gke_proj-a_us-central1_cluster-a": no cluster holds the lease cnrm-system/kcc-multicluster-lease`
	opts.statusFunc = func(ctx context.Context, h []*clusterHandle, now time.Time) *statusReport {
		return &statusReport{
			state:    haStateError,
			problems: []string{noHolderProblem},
		}
	}

	err := executeSetup(context.Background(), opts, handles, runner)
	if err == nil {
		t.Fatalf("expected executeSetup to fail on timeout, got nil")
	}
	if !strings.Contains(err.Error(), noHolderProblem) {
		t.Fatalf("expected error to contain %q, got: %v", noHolderProblem, err)
	}
}

func TestWaitTimeoutStandbySyncsFromWrongCluster(t *testing.T) {
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{}
	opts := newTestSetupOptions()
	opts.contexts = []string{handles[0].contextName, handles[1].contextName}
	opts.wait = true
	opts.waitTimeout = 2 * time.Millisecond
	opts.waitPollInterval = 1 * time.Millisecond

	wrongClusterProblem := `cluster "gke_proj-b_us-east1_cluster-b": KRMSyncer cnrm-system/kcc-multicluster-lease syncs from "projects/proj-c/locations/us-west1/clusters/cluster-c", but the lease is held by "projects/proj-a/locations/us-central1/clusters/cluster-a"`
	opts.statusFunc = func(ctx context.Context, h []*clusterHandle, now time.Time) *statusReport {
		return &statusReport{
			state:    haStateError,
			problems: []string{wrongClusterProblem},
		}
	}

	err := executeSetup(context.Background(), opts, handles, runner)
	if err == nil {
		t.Fatalf("expected executeSetup to fail on timeout, got nil")
	}
	if !strings.Contains(err.Error(), wrongClusterProblem) {
		t.Fatalf("expected error to contain wrong syncer problem, got: %v", err)
	}
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
		opts.waitTimeout = 2 * time.Millisecond
		opts.waitPollInterval = 1 * time.Millisecond
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
		opts.waitTimeout = 2 * time.Millisecond
		opts.waitPollInterval = 1 * time.Millisecond
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
