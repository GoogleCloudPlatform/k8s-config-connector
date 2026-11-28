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

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/kubecli"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	testStatusContextA = "gke_proj-a_us-central1_cluster-a"
	testStatusContextB = "gke_proj-b_us-east1_cluster-b"
)

func TestStatusValidateNormalizesFields(t *testing.T) {
	opts := newStatusOptions()
	opts.contexts = []string{" " + testStatusContextA + " ", " " + testStatusContextB}
	opts.bucket = "gs://my-ha-bucket/"

	err := opts.validate()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantContexts := []string{testStatusContextA, testStatusContextB}
	if !slices.Equal(opts.contexts, wantContexts) {
		t.Fatalf("got contexts %v, want %v", opts.contexts, wantContexts)
	}
	if opts.bucket != "my-ha-bucket" {
		t.Fatalf("got bucket %q, want %q", opts.bucket, "my-ha-bucket")
	}
}

func TestStatusValidateErrors(t *testing.T) {
	tests := []struct {
		name              string
		contexts          []string
		wantErrSubstrings []string
	}{
		{
			name:              "fewer than 2 contexts",
			contexts:          []string{testStatusContextA},
			wantErrSubstrings: []string{"at least 2 cluster contexts must be specified"},
		},
		{
			name:              "duplicate contexts",
			contexts:          []string{testStatusContextA, testStatusContextA},
			wantErrSubstrings: []string{"duplicate context \"gke_proj-a_us-central1_cluster-a\""},
		},
		{
			name:              "empty context",
			contexts:          []string{testStatusContextA, ""},
			wantErrSubstrings: []string{"context name in --contexts cannot be empty"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opts := newStatusOptions()
			opts.contexts = slices.Clone(tc.contexts)

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

func TestStatusCmdSilenceUsage(t *testing.T) {
	t.Run("shows usage when flag validation fails", func(t *testing.T) {
		cmd := NewStatusCmd()
		var stdout, stderr bytes.Buffer
		cmd.SetOut(&stdout)
		cmd.SetErr(&stderr)
		cmd.SetArgs([]string{"--contexts=cluster-a"})

		err := cmd.ExecuteContext(context.Background())
		if err == nil || !strings.Contains(err.Error(), "at least 2 cluster contexts must be specified") {
			t.Fatalf("expected at least 2 contexts error, got: %v", err)
		}
		if cmd.SilenceUsage {
			t.Errorf("SilenceUsage is true after flag validation error, want false")
		}
		if !strings.Contains(stdout.String()+stderr.String(), "Usage:") {
			t.Errorf("output does not include command usage:\nstdout:\n%s\nstderr:\n%s", stdout.String(), stderr.String())
		}
	})

	t.Run("silences usage once validation succeeds", func(t *testing.T) {
		cmd := NewStatusCmd()
		var stdout, stderr bytes.Buffer
		cmd.SetOut(&stdout)
		cmd.SetErr(&stderr)
		cmd.SetArgs([]string{"--contexts=gke_proj-a_us-central1_cluster-a,gke_proj-b_us-east1_cluster-b"})

		_ = cmd.ExecuteContext(context.Background())
		if !cmd.SilenceUsage {
			t.Errorf("SilenceUsage is false after validation succeeded, want true")
		}
	})
}

func TestRunStatusUnreachableClusters(t *testing.T) {
	useTestKubeconfig(t, testStatusContextA, testStatusContextB)
	var stdout bytes.Buffer
	opts := newStatusOptions()
	opts.contexts = []string{testStatusContextA, testStatusContextB}
	opts.stdout = &stdout

	err := runStatus(context.Background(), opts)

	if !errors.Is(err, errHAStatusError) {
		t.Fatalf("got error %v, want %v", err, errHAStatusError)
	}
	for _, want := range []string{
		"gke_proj-a_us-central1_cluster-a  proj-a   us-central1  Error  -\n",
		"gke_proj-b_us-east1_cluster-b     proj-b   us-east1     Error  -\n",
		"HA Status: ERROR\n",
		"  - fetching ConfigConnector on cluster \"gke_proj-a_us-central1_cluster-a\": ",
	} {
		if !strings.Contains(stdout.String(), want) {
			t.Fatalf("expected output to contain %q, got:\n%s", want, stdout.String())
		}
	}
}

func TestExecuteStatus(t *testing.T) {
	notConfigured := newConfigConnectorWithoutMCL()

	tests := []struct {
		name      string
		handles   func(t *testing.T) []*clusterHandle
		wantState haState
		want      string
	}{
		{
			name: "not detected",
			handles: func(t *testing.T) []*clusterHandle {
				return []*clusterHandle{
					newFakeClusterHandle(t, testStatusContextA, notConfigured.DeepCopy()),
					newFakeClusterHandle(t, testStatusContextB, notConfigured.DeepCopy()),
				}
			},
			wantState: haStateNotDetected,
			want: `CONTEXT                           PROJECT  LOCATION     HA              ROLE
gke_proj-a_us-central1_cluster-a  proj-a   us-central1  Not configured  -
gke_proj-b_us-east1_cluster-b     proj-b   us-east1     Not configured  -

HA Status: NOT DETECTED
Next step: config-connector ha setup --contexts=gke_proj-a_us-central1_cluster-a,gke_proj-b_us-east1_cluster-b --bucket=<BUCKET>
`,
		},
		{
			name: "healthy",
			handles: func(t *testing.T) []*clusterHandle {
				return healthyHAFixtures().handles(t)
			},
			wantState: haStateHealthy,
			want: `CONTEXT                           PROJECT  LOCATION     HA          ROLE
gke_proj-a_us-central1_cluster-a  proj-a   us-central1  Configured  Leader
gke_proj-b_us-east1_cluster-b     proj-b   us-east1     Configured  Standby

Lease: cnrm-system/kcc-multicluster-lease, held by gke_proj-a_us-central1_cluster-a (renewed 4s ago)
HA Status: HEALTHY
`,
		},
		{
			name: "configured on only some clusters",
			handles: func(t *testing.T) []*clusterHandle {
				return []*clusterHandle{
					healthyHAFixtures()[:1].handles(t)[0],
					newFakeClusterHandle(t, testStatusContextB, notConfigured.DeepCopy()),
				}
			},
			wantState: haStateError,
			want: `CONTEXT                           PROJECT  LOCATION     HA              ROLE
gke_proj-a_us-central1_cluster-a  proj-a   us-central1  Configured      Leader
gke_proj-b_us-east1_cluster-b     proj-b   us-east1     Not configured  -

Lease: cnrm-system/kcc-multicluster-lease, held by gke_proj-a_us-central1_cluster-a (renewed 4s ago)
HA Status: ERROR
Problems:
  - cluster "gke_proj-b_us-east1_cluster-b" does not have spec.experiments.multiClusterLease configured on ConfigConnector "configconnector.core.cnrm.cloud.google.com"; run ` + "`config-connector ha setup`" + ` or configure spec.experiments.multiClusterLease; it reconciles independently of the HA clusters, so more than one cluster may change the same resources
`,
		},
		{
			name: "missing ConfigConnector",
			handles: func(t *testing.T) []*clusterHandle {
				return []*clusterHandle{
					newFakeClusterHandle(t, testStatusContextA, notConfigured.DeepCopy()),
					newFakeClusterHandle(t, testStatusContextB),
				}
			},
			wantState: haStateError,
			want: `CONTEXT                           PROJECT  LOCATION     HA                  ROLE
gke_proj-a_us-central1_cluster-a  proj-a   us-central1  Not configured      -
gke_proj-b_us-east1_cluster-b     proj-b   us-east1     No ConfigConnector  -

HA Status: ERROR
Problems:
  - cluster "gke_proj-b_us-east1_cluster-b" is missing the "configconnector.core.cnrm.cloud.google.com" ConfigConnector object; install Config Connector and create the "configconnector.core.cnrm.cloud.google.com" resource before running HA commands
`,
		},
		{
			name: "unreachable cluster",
			handles: func(t *testing.T) []*clusterHandle {
				unreachable := newFakeClusterHandle(t, testStatusContextB)
				unreachable.client = &kubecli.Client{Client: &failingGetClient{Client: unreachable.client.Client, err: errors.New("dial tcp 10.0.0.1:443: i/o timeout")}}
				return []*clusterHandle{
					healthyHAFixtures()[:1].handles(t)[0],
					unreachable,
				}
			},
			wantState: haStateError,
			want: `CONTEXT                           PROJECT  LOCATION     HA          ROLE
gke_proj-a_us-central1_cluster-a  proj-a   us-central1  Configured  Leader
gke_proj-b_us-east1_cluster-b     proj-b   us-east1     Error       -

Lease: cnrm-system/kcc-multicluster-lease, held by gke_proj-a_us-central1_cluster-a (renewed 4s ago)
HA Status: ERROR
Problems:
  - fetching ConfigConnector on cluster "gke_proj-b_us-east1_cluster-b": dial tcp 10.0.0.1:443: i/o timeout
`,
		},
		{
			name: "invalid multiClusterLease spec",
			handles: func(t *testing.T) []*clusterHandle {
				return []*clusterHandle{
					healthyHAFixtures()[:1].handles(t)[0],
					newFakeClusterHandle(t, testStatusContextB, newConfigConnectorWithMCL("kcc-multicluster-lease", "cnrm-system", "")),
				}
			},
			wantState: haStateError,
			want: `CONTEXT                           PROJECT  LOCATION     HA          ROLE
gke_proj-a_us-central1_cluster-a  proj-a   us-central1  Configured  Leader
gke_proj-b_us-east1_cluster-b     proj-b   us-east1     Error       -

Lease: cnrm-system/kcc-multicluster-lease, held by gke_proj-a_us-central1_cluster-a (renewed 4s ago)
HA Status: ERROR
Problems:
  - cluster "gke_proj-b_us-east1_cluster-b": spec.experiments.multiClusterLease.gkeCluster is required
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			report := executeStatus(context.Background(), tc.handles(t), testNow)
			var out bytes.Buffer
			report.render(&out)

			if report.state != tc.wantState {
				t.Fatalf("got state %q, want %q", report.state, tc.wantState)
			}
			if got := out.String(); got != tc.want {
				t.Fatalf("got output:\n%s\nwant:\n%s", got, tc.want)
			}
		})
	}
}

// failingGetClient fails every Get with err, to simulate an unreachable cluster.
type failingGetClient struct {
	client.Client
	err error
}

func (f *failingGetClient) Get(_ context.Context, _ client.ObjectKey, _ client.Object, _ ...client.GetOption) error {
	return f.err
}

func newConfigConnectorWithoutMCL() *operatorv1beta1.ConfigConnector {
	return &operatorv1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{Name: configConnectorResourceName},
		Spec:       operatorv1beta1.ConfigConnectorSpec{Mode: "cluster"},
	}
}

func TestStatusConnectGateway_HealthyConfiguredZeroGCloudCalls(t *testing.T) {
	fA := newLeaderFixture(testStatusContextA)
	fA.context = "connectgateway_fleet-proj_global_cluster-a"
	handleA := haFixtures{fA}.handles(t)[0]
	handleA.client.RestConfig = &rest.Config{
		Host: "https://connectgateway.googleapis.com/v1/projects/fleet-proj/locations/global/gkeMemberships/cluster-a",
	}

	fB := newStandbyFixture(testStatusContextB, testIdentityA)
	fB.context = "connectgateway_fleet-proj_global_cluster-b"
	handleB := haFixtures{fB}.handles(t)[0]
	handleB.client.RestConfig = &rest.Config{
		Host: "https://connectgateway.googleapis.com/v1/projects/fleet-proj/locations/global/gkeMemberships/cluster-b",
	}

	runner := &recordingCommandRunner{}
	report := executeStatusWithRunner(context.Background(), []*clusterHandle{handleA, handleB}, testNow, runner)

	if len(runner.commands) != 0 {
		t.Fatalf("expected 0 gcloud commands for healthy configured clusters, got: %v", runner.commands)
	}

	if report.state != haStateHealthy {
		t.Fatalf("got state %q, want %q; problems: %v", report.state, haStateHealthy, report.problems)
	}

	var out bytes.Buffer
	report.render(&out)
	got := out.String()
	for _, want := range []string{
		"connectgateway_fleet-proj_global_cluster-a  proj-a   us-central1  Configured  Leader",
		"connectgateway_fleet-proj_global_cluster-b  proj-b   us-east1     Configured  Standby",
		"Lease: cnrm-system/kcc-multicluster-lease, held by connectgateway_fleet-proj_global_cluster-a",
		"HA Status: HEALTHY",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("output missing %q:\n%s", want, got)
		}
	}
}

func TestStatusConnectGateway_UnreachableCluster(t *testing.T) {
	fA := newStandbyFixture(testStatusContextA, testIdentityB)
	fA.context = "connectgateway_fleet-proj_global_cluster-a"
	handleA := haFixtures{fA}.handles(t)[0]
	handleA.client.RestConfig = &rest.Config{
		Host: "https://connectgateway.googleapis.com/v1/projects/fleet-proj/locations/global/gkeMemberships/cluster-a",
	}

	unreachableClient := &failingGetClient{err: errors.New("connection refused")}
	handleB := &clusterHandle{
		contextName: "connectgateway_fleet-proj_global_cluster-b",
		client: &kubecli.Client{
			Client: unreachableClient,
			RestConfig: &rest.Config{
				Host: "https://connectgateway.googleapis.com/v1/projects/fleet-proj/locations/global/gkeMemberships/cluster-b",
			},
		},
	}

	runner := &recordingCommandRunner{
		outputFn: func(ctx context.Context, name string, args ...string) ([]byte, error) {
			return []byte(`{
				"endpoint": {
					"gkeCluster": {
						"resourceLink": "//container.googleapis.com/projects/proj-b/locations/us-east1/clusters/cluster-b"
					}
				}
			}`), nil
		},
	}

	report := executeStatusWithRunner(context.Background(), []*clusterHandle{handleA, handleB}, testNow, runner)

	if report.clusters[1].gke.project != "proj-b" || report.clusters[1].gke.location != "us-east1" {
		t.Fatalf("expected cluster B GKE info to be proj-b/us-east1, got %+v", report.clusters[1].gke)
	}

	if report.lease == nil || report.lease.holder != "connectgateway_fleet-proj_global_cluster-b" {
		t.Fatalf("expected lease holder to be %q, got %+v", "connectgateway_fleet-proj_global_cluster-b", report.lease)
	}

	for _, prob := range report.problems {
		if strings.Contains(prob, "leader") && strings.Contains(prob, "outside") {
			t.Fatalf("unexpected false problem about leader outside --contexts: %v", prob)
		}
	}

	var out bytes.Buffer
	report.render(&out)
	got := out.String()
	for _, want := range []string{
		"connectgateway_fleet-proj_global_cluster-a  proj-a   us-central1  Configured  Standby",
		"connectgateway_fleet-proj_global_cluster-b  proj-b   us-east1     Error       -",
		"Lease: cnrm-system/kcc-multicluster-lease, held by connectgateway_fleet-proj_global_cluster-b",
		"HA Status: ERROR",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("output missing %q:\n%s", want, got)
		}
	}
}
