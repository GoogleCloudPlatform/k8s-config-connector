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

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
)

func fullClusterDescribeJSON() string {
	return `{
		"endpoint": "1.2.3.4",
		"privateClusterConfig": {
			"privateEndpoint": "10.0.0.2",
			"masterGlobalAccessConfig": {
				"enabled": true
			}
		},
		"controlPlaneEndpointsConfig": {
			"dnsEndpointConfig": {
				"endpoint": "gke-123.us-central1.gke.goog",
				"allowExternalTraffic": true
			},
			"ipEndpointsConfig": {
				"privateEndpoint": "10.0.0.2",
				"globalAccess": true
			}
		},
		"networkConfig": {
			"network": "projects/proj-a/global/networks/default"
		}
	}`
}

func TestCheckLeaderEndpointsPassesForExposedEndpoints(t *testing.T) {
	for _, tc := range []struct {
		name         string
		endpointType string
		clusterAJSON string
		clusterBJSON string
	}{
		{
			name:         "Default endpoint exposed",
			endpointType: operatorv1beta1.LeaderEndpointTypeDefault,
			clusterAJSON: `{"endpoint": "1.2.3.4"}`,
			clusterBJSON: `{"endpoint": "5.6.7.8"}`,
		},
		{
			name:         "DNS endpoint exposed",
			endpointType: operatorv1beta1.LeaderEndpointTypeDNS,
			clusterAJSON: `{"controlPlaneEndpointsConfig": {"dnsEndpointConfig": {"endpoint": "a.gke.goog", "allowExternalTraffic": true}}}`,
			clusterBJSON: `{"controlPlaneEndpointsConfig": {"dnsEndpointConfig": {"endpoint": "b.gke.goog", "allowExternalTraffic": true}}}`,
		},
		{
			name:         "PrivateIP exposed via privateClusterConfig",
			endpointType: operatorv1beta1.LeaderEndpointTypePrivateIP,
			clusterAJSON: `{"privateClusterConfig": {"privateEndpoint": "10.0.0.1"}}`,
			clusterBJSON: `{"privateClusterConfig": {"privateEndpoint": "10.0.0.2"}}`,
		},
		{
			name:         "PrivateIP exposed via ipEndpointsConfig",
			endpointType: operatorv1beta1.LeaderEndpointTypePrivateIP,
			clusterAJSON: `{"controlPlaneEndpointsConfig": {"ipEndpointsConfig": {"privateEndpoint": "10.0.0.1"}}}`,
			clusterBJSON: `{"controlPlaneEndpointsConfig": {"ipEndpointsConfig": {"privateEndpoint": "10.0.0.2"}}}`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			handles := []*clusterHandle{
				newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
				newFakeClusterHandle(t, "gke_proj-b_us-central1_cluster-b", newBaseConfigConnector()),
			}
			runner := &recordingCommandRunner{
				describeOutputs: map[string][]byte{
					"cluster-a": []byte(tc.clusterAJSON),
					"cluster-b": []byte(tc.clusterBJSON),
				},
			}
			opts := newTestSetupOptions()
			opts.leaderEndpointType = tc.endpointType
			opts.leaderEndpointTypeExplicit = true

			var stdout bytes.Buffer
			opts.stdout = &stdout

			err := executeSetup(context.Background(), opts, handles, runner)
			if err != nil {
				t.Fatalf("expected executeSetup to pass, got: %v", err)
			}
			wantCheckedLine := "Checked " + tc.endpointType + " leader endpoint on gke_proj-a_us-central1_cluster-a, gke_proj-b_us-central1_cluster-b.\n"
			if !strings.Contains(stdout.String(), wantCheckedLine) {
				t.Fatalf("expected stdout to contain %q, got:\n%s", wantCheckedLine, stdout.String())
			}
		})
	}
}

func TestCheckLeaderEndpointsFailureTableRows(t *testing.T) {
	tests := []struct {
		name         string
		endpointType string
		clusterBJSON string
		wantErrPart  string
		wantFix      string
	}{
		{
			name:         "DNS fails when endpoint is empty",
			endpointType: operatorv1beta1.LeaderEndpointTypeDNS,
			clusterBJSON: `{"controlPlaneEndpointsConfig": {"dnsEndpointConfig": {"endpoint": "", "allowExternalTraffic": true}}}`,
			wantErrPart:  "cluster \"gke_proj-b_us-east1_cluster-b\": DNS endpoint is not enabled for external traffic",
			wantFix:      "gcloud container clusters update cluster-b --location=us-east1 --project=proj-b --enable-dns-access",
		},
		{
			name:         "DNS fails when allowExternalTraffic is not true",
			endpointType: operatorv1beta1.LeaderEndpointTypeDNS,
			clusterBJSON: `{"controlPlaneEndpointsConfig": {"dnsEndpointConfig": {"endpoint": "b.gke.goog", "allowExternalTraffic": false}}}`,
			wantErrPart:  "cluster \"gke_proj-b_us-east1_cluster-b\": DNS endpoint is not enabled for external traffic",
			wantFix:      "gcloud container clusters update cluster-b --location=us-east1 --project=proj-b --enable-dns-access",
		},
		{
			name:         "PrivateIP fails when both privateEndpoints are empty",
			endpointType: operatorv1beta1.LeaderEndpointTypePrivateIP,
			clusterBJSON: `{"endpoint": "1.2.3.4"}`,
			wantErrPart:  "cluster \"gke_proj-b_us-east1_cluster-b\": cluster has no private endpoint",
			wantFix:      "choose --leader-endpoint-type=DNS or Default",
		},
		{
			name:         "Default fails when endpoint is empty",
			endpointType: operatorv1beta1.LeaderEndpointTypeDefault,
			clusterBJSON: `{"privateClusterConfig": {"privateEndpoint": "10.0.0.2"}}`,
			wantErrPart:  "cluster \"gke_proj-b_us-east1_cluster-b\": cluster has no default endpoint",
			wantFix:      "choose --leader-endpoint-type=DNS",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handles := []*clusterHandle{
				newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
				newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
			}
			runner := &recordingCommandRunner{
				describeOutputs: map[string][]byte{
					"cluster-a": []byte(fullClusterDescribeJSON()),
					"cluster-b": []byte(tc.clusterBJSON),
				},
			}
			opts := newTestSetupOptions()
			opts.leaderEndpointType = tc.endpointType
			opts.leaderEndpointTypeExplicit = true

			err := executeSetup(context.Background(), opts, handles, runner)
			if err == nil {
				t.Fatalf("expected executeSetup to fail, got nil")
			}
			if !strings.Contains(err.Error(), tc.wantErrPart) {
				t.Fatalf("expected error containing %q, got: %v", tc.wantErrPart, err)
			}
			if !strings.Contains(err.Error(), tc.wantFix) {
				t.Fatalf("expected error containing fix %q, got: %v", tc.wantFix, err)
			}
			if len(runner.runCommands) != 0 {
				t.Fatalf("expected 0 mutating commands when endpoint check fails, got: %v", runner.runCommands)
			}
		})
	}
}

func TestCheckLeaderEndpointsTwoFailingClustersReported(t *testing.T) {
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-c_us-west1_cluster-c", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{
		describeOutputs: map[string][]byte{
			"cluster-a": []byte(fullClusterDescribeJSON()),
			"cluster-b": []byte(`{"endpoint": ""}`),
			"cluster-c": []byte(`{"endpoint": ""}`),
		},
	}
	opts := newTestSetupOptions()
	opts.leaderEndpointType = operatorv1beta1.LeaderEndpointTypeDefault
	opts.leaderEndpointTypeExplicit = true

	err := executeSetup(context.Background(), opts, handles, runner)
	if err == nil {
		t.Fatalf("expected executeSetup to fail, got nil")
	}

	if !strings.Contains(err.Error(), "gke_proj-b_us-east1_cluster-b") {
		t.Fatalf("expected error to report cluster-b failure, got: %v", err)
	}
	if !strings.Contains(err.Error(), "gke_proj-c_us-west1_cluster-c") {
		t.Fatalf("expected error to report cluster-c failure, got: %v", err)
	}
	if len(runner.runCommands) != 0 {
		t.Fatalf("expected 0 mutating commands when check fails, got: %v", runner.runCommands)
	}
}

func TestCheckLeaderEndpointsWarningTableRows(t *testing.T) {
	tests := []struct {
		name         string
		endpointType string
		clusterAJSON string
		clusterBJSON string
		wantWarning  string
	}{
		{
			name:         "PrivateIP cross-region without global access warns",
			endpointType: operatorv1beta1.LeaderEndpointTypePrivateIP,
			clusterAJSON: `{
				"privateClusterConfig": {"privateEndpoint": "10.0.0.1"},
				"networkConfig": {"network": "projects/p/global/networks/vpc1"}
			}`,
			clusterBJSON: `{
				"privateClusterConfig": {
					"privateEndpoint": "10.0.0.2",
					"masterGlobalAccessConfig": {"enabled": false}
				},
				"networkConfig": {"network": "projects/p/global/networks/vpc1"}
			}`,
			wantWarning: "connections from another region to the private endpoint time out unless global access is on: gcloud container clusters update cluster-b --location=us-east1-b --project=proj-b --enable-master-global-access",
		},
		{
			name:         "PrivateIP differing networks warns",
			endpointType: operatorv1beta1.LeaderEndpointTypePrivateIP,
			clusterAJSON: `{
				"privateClusterConfig": {"privateEndpoint": "10.0.0.1", "masterGlobalAccessConfig": {"enabled": true}},
				"networkConfig": {"network": "projects/p/global/networks/vpc-a"}
			}`,
			clusterBJSON: `{
				"privateClusterConfig": {"privateEndpoint": "10.0.0.2", "masterGlobalAccessConfig": {"enabled": true}},
				"networkConfig": {"network": "projects/p/global/networks/vpc-b"}
			}`,
			wantWarning: "the clusters are in different VPC networks: PrivateIP requires Cloud VPN or Cloud Interconnect (VPC Network Peering cannot reach GKE private endpoints); otherwise use --leader-endpoint-type=DNS",
		},
		{
			name:         "PrivateIP privateEndpointEnforcementEnabled warns",
			endpointType: operatorv1beta1.LeaderEndpointTypePrivateIP,
			clusterAJSON: `{
				"privateClusterConfig": {"privateEndpoint": "10.0.0.1", "masterGlobalAccessConfig": {"enabled": true}},
				"networkConfig": {"network": "projects/p/global/networks/vpc1"}
			}`,
			clusterBJSON: `{
				"privateClusterConfig": {"privateEndpoint": "10.0.0.2", "masterGlobalAccessConfig": {"enabled": true}},
				"networkConfig": {"network": "projects/p/global/networks/vpc1"},
				"masterAuthorizedNetworksConfig": {"privateEndpointEnforcementEnabled": true}
			}`,
			wantWarning: "add gke_proj-a_us-central1-a_cluster-a's node and Pod ranges to gke_proj-b_us-east1-b_cluster-b's authorized networks",
		},
		{
			name:         "Default masterAuthorizedNetworksConfig.enabled warns",
			endpointType: operatorv1beta1.LeaderEndpointTypeDefault,
			clusterAJSON: `{"endpoint": "1.2.3.4"}`,
			clusterBJSON: `{
				"endpoint": "5.6.7.8",
				"masterAuthorizedNetworksConfig": {"enabled": true}
			}`,
			wantWarning: "add gke_proj-a_us-central1-a_cluster-a's egress addresses (for example its Cloud NAT addresses) to gke_proj-b_us-east1-b_cluster-b's authorized networks",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handles := []*clusterHandle{
				newFakeClusterHandle(t, "gke_proj-a_us-central1-a_cluster-a", newBaseConfigConnector()),
				newFakeClusterHandle(t, "gke_proj-b_us-east1-b_cluster-b", newBaseConfigConnector()),
			}
			runner := &recordingCommandRunner{
				describeOutputs: map[string][]byte{
					"cluster-a": []byte(tc.clusterAJSON),
					"cluster-b": []byte(tc.clusterBJSON),
				},
			}
			opts := newTestSetupOptions()
			opts.leaderEndpointType = tc.endpointType
			opts.leaderEndpointTypeExplicit = true

			var stderr bytes.Buffer
			opts.stderr = &stderr

			err := executeSetup(context.Background(), opts, handles, runner)
			if err != nil {
				t.Fatalf("expected setup to continue and pass, got: %v", err)
			}
			warnOutput := stderr.String()
			if !strings.Contains(warnOutput, tc.wantWarning) {
				t.Fatalf("expected warning %q, got: %q", tc.wantWarning, warnOutput)
			}
			// Verify warning is printed once
			count := strings.Count(warnOutput, tc.wantWarning)
			if count != 1 {
				t.Fatalf("expected warning to be printed once, got %d occurrences in:\n%s", count, warnOutput)
			}
		})
	}
}

func TestCheckLeaderEndpointsFailsNothingMutated(t *testing.T) {
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{
		describeOutputs: map[string][]byte{
			"cluster-a": []byte(fullClusterDescribeJSON()),
			"cluster-b": []byte(`{"endpoint": ""}`),
		},
	}
	opts := newTestSetupOptions()
	opts.leaderEndpointType = operatorv1beta1.LeaderEndpointTypeDefault
	opts.leaderEndpointTypeExplicit = true

	err := executeSetup(context.Background(), opts, handles, runner)
	if err == nil {
		t.Fatalf("expected executeSetup to fail, got nil")
	}

	// No mutating gcloud commands ran
	if len(runner.runCommands) != 0 {
		t.Fatalf("expected 0 mutating commands when endpoint check fails, got: %v", runner.runCommands)
	}

	// No Deployment or updated ConfigConnector on clusters
	mclDeploy := &appsv1.Deployment{}
	for _, h := range handles {
		if err := h.client.Get(context.Background(), types.NamespacedName{Namespace: "cnrm-system", Name: "multiclusterlease-controller-manager"}, mclDeploy); err == nil {
			t.Fatalf("expected no Deployment created on %s, but found one", h.contextName)
		}
		cc, err := fetchConfigConnector(context.Background(), h)
		if err != nil {
			t.Fatalf("fetching ConfigConnector on %s: %v", h.contextName, err)
		}
		if cc.Spec.Experiments != nil && cc.Spec.Experiments.MultiClusterLease != nil {
			t.Fatalf("expected ConfigConnector on %s not to be modified, got: %+v", h.contextName, cc.Spec.Experiments.MultiClusterLease)
		}
	}
}

func TestCheckLeaderEndpointsReplicationDisabledAddsNoRequirement(t *testing.T) {
	// Cluster A has replicationMode: Disabled, leaderEndpointType: DNS.
	// Cluster B does NOT expose DNS (so if A required DNS on B, setup would fail).
	// But because A has replicationMode: Disabled, A runs no syncer, so B is not required to expose DNS.
	ccA := newBaseConfigConnector()
	ccA.Spec.Experiments = &operatorv1beta1.CCExperiments{
		MultiClusterLease: &operatorv1beta1.MultiClusterLeaseSpec{
			ResourceReplicationMode: "Disabled",
			LeaderEndpointType:      operatorv1beta1.LeaderEndpointTypeDNS,
		},
	}

	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", ccA),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{
		describeOutputs: map[string][]byte{
			// Cluster A exposes Default endpoint (needed by B)
			"cluster-a": []byte(`{"endpoint": "1.2.3.4"}`),
			// Cluster B exposes Default endpoint, but does NOT expose DNS
			"cluster-b": []byte(`{"endpoint": "5.6.7.8"}`),
		},
	}
	opts := newTestSetupOptions()
	// Leave leaderEndpointType as Default for B; A preserves its own ConfigConnector settings

	err := executeSetup(context.Background(), opts, handles, runner)
	if err != nil {
		t.Fatalf("expected executeSetup to pass because cluster A is Disabled, got: %v", err)
	}
}

func TestCheckLeaderEndpointsDryRunRunsCheckAndYamlOnlySkips(t *testing.T) {
	// 1. Dry run runs check: if cluster B fails endpoint check, dry-run fails
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{
		describeOutputs: map[string][]byte{
			"cluster-a": []byte(fullClusterDescribeJSON()),
			"cluster-b": []byte(`{"endpoint": ""}`),
		},
	}
	opts := newTestSetupOptions()
	opts.dryRun = true
	opts.leaderEndpointType = operatorv1beta1.LeaderEndpointTypeDefault
	opts.leaderEndpointTypeExplicit = true

	err := executeSetup(context.Background(), opts, handles, runner)
	if err == nil {
		t.Fatalf("expected dry-run to fail endpoint check on cluster B, got nil")
	}

	// 2. YAML-only skips endpoint check entirely and does not call runner
	yamlRunner := &recordingCommandRunner{
		describeOutputs: map[string][]byte{
			"cluster-a": []byte(`{"endpoint": ""}`),
			"cluster-b": []byte(`{"endpoint": ""}`),
		},
	}
	yamlOpts := newTestSetupOptions()
	yamlOpts.yamlOnly = true
	var stdout bytes.Buffer
	yamlOpts.stdout = &stdout

	err = executeSetup(context.Background(), yamlOpts, handles, yamlRunner)
	if err != nil {
		t.Fatalf("expected --yaml-only to succeed, got: %v", err)
	}
	if len(yamlRunner.commands) != 0 {
		t.Fatalf("expected 0 commands called in yaml-only mode, got: %v", yamlRunner.commands)
	}
}

func TestCheckLeaderEndpointsDescribeFailureFailsSetup(t *testing.T) {
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{
		describeErr: map[string]error{
			"cluster-b": context.DeadlineExceeded,
		},
	}
	opts := newTestSetupOptions()

	err := executeSetup(context.Background(), opts, handles, runner)
	if err == nil {
		t.Fatalf("expected executeSetup to fail when describe fails, got nil")
	}
	if !strings.Contains(err.Error(), "gke_proj-b_us-east1_cluster-b") {
		t.Fatalf("expected error naming cluster-b, got: %v", err)
	}
}

func TestCheckLeaderEndpointsNoWarningsWhenCheckFails(t *testing.T) {
	handles := []*clusterHandle{
		newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector()),
		newFakeClusterHandle(t, "gke_proj-b_us-east1_cluster-b", newBaseConfigConnector()),
	}
	runner := &recordingCommandRunner{
		describeOutputs: map[string][]byte{
			"cluster-a": []byte(fullClusterDescribeJSON()),
			// cluster-b is a standard public cluster (no private endpoint) across regions
			"cluster-b": []byte(`{"endpoint": "1.2.3.4"}`),
		},
	}
	opts := newTestSetupOptions()
	opts.leaderEndpointType = operatorv1beta1.LeaderEndpointTypePrivateIP
	opts.leaderEndpointTypeExplicit = true
	var stderr bytes.Buffer
	opts.stderr = &stderr

	err := executeSetup(context.Background(), opts, handles, runner)
	if err == nil {
		t.Fatalf("expected executeSetup to fail, got nil")
	}
	if !strings.Contains(err.Error(), "cluster has no private endpoint") {
		t.Fatalf("expected error about missing private endpoint, got: %v", err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected no warnings on stderr when endpoint check fails, got:\n%s", stderr.String())
	}
}
