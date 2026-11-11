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
			name:              "non-GKE context",
			contexts:          []string{testStatusContextA, "prod-east"},
			wantErrSubstrings: []string{"context \"prod-east\" does not match expected GKE context format"},
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
		"gke_proj-a_us-central1_cluster-a  proj-a   us-central1  Unreachable\n",
		"gke_proj-b_us-east1_cluster-b     proj-b   us-east1     Unreachable\n",
		"HA Status: ERROR\n",
		"  - cluster-a: could not read the ConfigConnector object: ",
	} {
		if !strings.Contains(stdout.String(), want) {
			t.Fatalf("expected output to contain %q, got:\n%s", want, stdout.String())
		}
	}
}

func TestExecuteStatus(t *testing.T) {
	configured := newConfigConnectorWithMCL("kcc-multicluster-lease", "cnrm-system", "cluster")
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
			want: `CONTEXT                           PROJECT  LOCATION     HA
gke_proj-a_us-central1_cluster-a  proj-a   us-central1  Not configured
gke_proj-b_us-east1_cluster-b     proj-b   us-east1     Not configured

HA Status: NOT DETECTED
Next step: config-connector ha setup --contexts=gke_proj-a_us-central1_cluster-a,gke_proj-b_us-east1_cluster-b --bucket=<BUCKET>
`,
		},
		{
			name: "detected on all clusters",
			handles: func(t *testing.T) []*clusterHandle {
				return []*clusterHandle{
					newFakeClusterHandle(t, testStatusContextA, configured.DeepCopy()),
					newFakeClusterHandle(t, testStatusContextB, configured.DeepCopy()),
				}
			},
			wantState: haStateDetected,
			want: `CONTEXT                           PROJECT  LOCATION     HA
gke_proj-a_us-central1_cluster-a  proj-a   us-central1  Configured
gke_proj-b_us-east1_cluster-b     proj-b   us-east1     Configured

HA Status: DETECTED
Note: leadership and data sync status are not reported yet.
`,
		},
		{
			name: "configured on only some clusters",
			handles: func(t *testing.T) []*clusterHandle {
				return []*clusterHandle{
					newFakeClusterHandle(t, testStatusContextA, configured.DeepCopy()),
					newFakeClusterHandle(t, testStatusContextB, notConfigured.DeepCopy()),
				}
			},
			wantState: haStateError,
			want: `CONTEXT                           PROJECT  LOCATION     HA
gke_proj-a_us-central1_cluster-a  proj-a   us-central1  Configured
gke_proj-b_us-east1_cluster-b     proj-b   us-east1     Not configured

HA Status: ERROR
Problems:
  - cluster-b: spec.experiments.multiClusterLease is not set on the ConfigConnector; run ` + "`config-connector ha setup`" + ` with all clusters
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
			want: `CONTEXT                           PROJECT  LOCATION     HA
gke_proj-a_us-central1_cluster-a  proj-a   us-central1  Not configured
gke_proj-b_us-east1_cluster-b     proj-b   us-east1     No ConfigConnector

HA Status: ERROR
Problems:
  - cluster-b: the "configconnector.core.cnrm.cloud.google.com" ConfigConnector object was not found; install Config Connector with the operator
`,
		},
		{
			name: "unreachable cluster",
			handles: func(t *testing.T) []*clusterHandle {
				unreachable := newFakeClusterHandle(t, testStatusContextB)
				unreachable.client = &kubecli.Client{Client: &failingGetClient{Client: unreachable.client.Client, err: errors.New("dial tcp 10.0.0.1:443: i/o timeout")}}
				return []*clusterHandle{
					newFakeClusterHandle(t, testStatusContextA, configured.DeepCopy()),
					unreachable,
				}
			},
			wantState: haStateError,
			want: `CONTEXT                           PROJECT  LOCATION     HA
gke_proj-a_us-central1_cluster-a  proj-a   us-central1  Configured
gke_proj-b_us-east1_cluster-b     proj-b   us-east1     Unreachable

HA Status: ERROR
Problems:
  - cluster-b: could not read the ConfigConnector object: dial tcp 10.0.0.1:443: i/o timeout
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			report := executeStatus(context.Background(), tc.handles(t))
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
