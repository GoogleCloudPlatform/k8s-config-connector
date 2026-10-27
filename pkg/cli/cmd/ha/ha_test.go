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

package ha_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/ha"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/kubecli"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestHACommandFlagsAndValidation(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantErrSub string
	}{
		{
			name:       "setup fails with fewer than 2 contexts",
			args:       []string{"ha", "setup", "--contexts=cluster-a", "--bucket=my-bucket"},
			wantErrSub: "at least 2 cluster contexts must be specified",
		},
		{
			name:       "setup fails with duplicate contexts",
			args:       []string{"ha", "setup", "--contexts=cluster-a,cluster-a", "--bucket=my-bucket"},
			wantErrSub: "duplicate context \"cluster-a\"",
		},
		{
			name:       "setup fails with duplicate contexts padded with whitespace",
			args:       []string{"ha", "setup", "--contexts=cluster-a, cluster-a ", "--bucket=my-bucket"},
			wantErrSub: "duplicate context \"cluster-a\"",
		},
		{
			name:       "setup fails without bucket",
			args:       []string{"ha", "setup", "--contexts=cluster-a,cluster-b"},
			wantErrSub: "--bucket is required",
		},
		{
			name:       "setup fails with whitespace-only bucket",
			args:       []string{"ha", "setup", "--contexts=cluster-a,cluster-b", "--bucket=   "},
			wantErrSub: "--bucket is required",
		},
		{
			name:       "status fails with fewer than 2 contexts",
			args:       []string{"ha", "status", "--contexts=cluster-a"},
			wantErrSub: "at least 2 cluster contexts must be specified",
		},
		{
			name:       "failover fails without target-cluster",
			args:       []string{"ha", "failover", "--contexts=cluster-a,cluster-b"},
			wantErrSub: "--target-cluster is required",
		},
		{
			name:       "failover fails when target-cluster is not in contexts",
			args:       []string{"ha", "failover", "--contexts=cluster-a,cluster-b", "--target-cluster=cluster-c"},
			wantErrSub: "must be one of the contexts specified in --contexts",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			haCmd := cmd.NewHACmd()
			var stdout, stderr bytes.Buffer
			haCmd.SetOut(&stdout)
			haCmd.SetErr(&stderr)
			// Strip leading "ha" since we invoke NewHACmd() directly.
			haCmd.SetArgs(tc.args[1:])

			err := haCmd.Execute()
			if err == nil {
				t.Fatalf("expected error containing %q, got nil", tc.wantErrSub)
			}
			if !strings.Contains(err.Error(), tc.wantErrSub) {
				t.Fatalf("expected error containing %q, got %v", tc.wantErrSub, err)
			}
		})
	}

	t.Run("whitespace-padded contexts and target-cluster normalize cleanly", func(t *testing.T) {
		baseOpts := ha.NewBaseOptions()
		baseOpts.Contexts = []string{" cluster-a ", " cluster-b"}
		failoverOpts := ha.NewFailoverOptions(baseOpts)
		failoverOpts.TargetCluster = " cluster-b "

		if err := failoverOpts.Validate(); err != nil {
			t.Fatalf("Validate() failed on whitespace-padded contexts: %v", err)
		}
		if baseOpts.Contexts[0] != "cluster-a" || baseOpts.Contexts[1] != "cluster-b" {
			t.Fatalf("expected trimmed contexts [cluster-a cluster-b], got %v", baseOpts.Contexts)
		}
		if failoverOpts.TargetCluster != "cluster-b" {
			t.Fatalf("expected trimmed target-cluster \"cluster-b\", got %q", failoverOpts.TargetCluster)
		}
	})

	t.Run("pre-populated Clusters are validated against target-cluster", func(t *testing.T) {
		baseOpts := ha.NewBaseOptions()
		baseOpts.Clusters = []*ha.ClusterHandle{
			newFakeClusterHandle(t, "cluster-a"),
			newFakeClusterHandle(t, "cluster-b"),
		}
		failoverOpts := ha.NewFailoverOptions(baseOpts)
		failoverOpts.TargetCluster = "cluster-c"

		err := failoverOpts.Validate()
		if err == nil || !strings.Contains(err.Error(), "must be one of the contexts specified in --contexts") {
			t.Fatalf("expected target-cluster mismatch error when using pre-populated Clusters, got %v", err)
		}
	})

	t.Run("rootCmd routes ha subcommand via ExecuteFromTest", func(t *testing.T) {
		opts := &cmd.TestInvocationOptions{
			Args: []string{"config-connector", "ha", "status", "--contexts=cluster-a"},
		}
		err := cmd.ExecuteFromTest(opts)
		if err == nil || !strings.Contains(err.Error(), "at least 2 cluster contexts must be specified") {
			t.Fatalf("expected contexts validation error via rootCmd, got %v", err)
		}
	})
}

func TestFetchAndValidateMultiClusterLeaseConfigs(t *testing.T) {
	ctx := context.Background()

	t.Run("valid two-cluster HA state", func(t *testing.T) {
		leaderIdentity := "cluster-a-id"
		newLease := func() *mclv1alpha1.MultiClusterLease {
			return &mclv1alpha1.MultiClusterLease{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "kcc-lease",
					Namespace: "cnrm-system",
				},
				Status: mclv1alpha1.MultiClusterLeaseStatus{
					GlobalHolderIdentity: &leaderIdentity,
				},
			}
		}
		newSyncer := func() *syncerv1alpha1.KRMSyncer {
			return &syncerv1alpha1.KRMSyncer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "kcc-lease",
					Namespace: "cnrm-system",
				},
			}
		}

		c1 := newFakeClusterHandle(t, "cluster-a",
			newConfigConnectorWithMCL("kcc-lease", "cnrm-system", "cluster-a-id"),
			newLease(),
			newSyncer(),
		)
		c2 := newFakeClusterHandle(t, "cluster-b",
			newConfigConnectorWithMCL("kcc-lease", "cnrm-system", "cluster-b-id"),
			newLease(),
			newSyncer(),
		)

		states, err := ha.FetchAllClusterHAStates(ctx, []*ha.ClusterHandle{c1, c2})
		if err != nil {
			t.Fatalf("FetchAllClusterHAStates failed: %v", err)
		}
		if err := ha.ValidateMultiClusterLeaseConfigs(states); err != nil {
			t.Fatalf("ValidateMultiClusterLeaseConfigs failed on valid config: %v", err)
		}
		if states[0].MultiClusterLease == nil || states[0].MultiClusterLease.Status.GlobalHolderIdentity == nil || *states[0].MultiClusterLease.Status.GlobalHolderIdentity != leaderIdentity {
			t.Fatalf("expected GlobalHolderIdentity %q, got %+v", leaderIdentity, states[0].MultiClusterLease)
		}
		if len(states[0].KRMSyncers) != 1 {
			t.Fatalf("expected 1 KRMSyncer, got %d", len(states[0].KRMSyncers))
		}
	})

	t.Run("missing HA CRDs on cluster are tolerated during state fetch", func(t *testing.T) {
		// Build a scheme with ConfigConnector registered, but without MultiClusterLease or KRMSyncer CRDs.
		schemeWithoutHACRDs := runtime.NewScheme()
		if err := operatorv1beta1.AddToScheme(schemeWithoutHACRDs); err != nil {
			t.Fatalf("AddToScheme failed: %v", err)
		}
		cc := newConfigConnectorWithMCL("kcc-lease", "cnrm-system", "cluster-a-id")
		fakeClient := fake.NewClientBuilder().WithScheme(schemeWithoutHACRDs).WithRuntimeObjects(cc).Build()
		handle := &ha.ClusterHandle{
			ContextName: "cluster-a",
			Client:      &kubecli.Client{Client: fakeClient},
		}

		state, err := ha.FetchClusterHAState(ctx, handle)
		if err != nil {
			t.Fatalf("expected FetchClusterHAState to tolerate NoKindMatchError, got: %v", err)
		}
		if state.ConfigConnector == nil {
			t.Fatalf("expected ConfigConnector to be populated")
		}
		if state.MultiClusterLease != nil || len(state.KRMSyncers) != 0 {
			t.Fatalf("expected nil MultiClusterLease and empty KRMSyncers when CRDs are absent, got %+v", state)
		}
	})

	t.Run("mismatched leaseName fails validation", func(t *testing.T) {
		c1 := newFakeClusterHandle(t, "cluster-a",
			newConfigConnectorWithMCL("kcc-lease-1", "cnrm-system", "cluster-a-id"),
		)
		c2 := newFakeClusterHandle(t, "cluster-b",
			newConfigConnectorWithMCL("kcc-lease-2", "cnrm-system", "cluster-b-id"),
		)

		states, err := ha.FetchAllClusterHAStates(ctx, []*ha.ClusterHandle{c1, c2})
		if err != nil {
			t.Fatalf("FetchAllClusterHAStates failed: %v", err)
		}
		err = ha.ValidateMultiClusterLeaseConfigs(states)
		if err == nil || !strings.Contains(err.Error(), "mismatched leaseName") {
			t.Fatalf("expected mismatched leaseName error, got %v", err)
		}
	})

	t.Run("mismatched lease namespace fails validation", func(t *testing.T) {
		c1 := newFakeClusterHandle(t, "cluster-a",
			newConfigConnectorWithMCL("kcc-lease", "cnrm-system", "cluster-a-id"),
		)
		c2 := newFakeClusterHandle(t, "cluster-b",
			newConfigConnectorWithMCL("kcc-lease", "other-ns", "cluster-b-id"),
		)

		states, err := ha.FetchAllClusterHAStates(ctx, []*ha.ClusterHandle{c1, c2})
		if err != nil {
			t.Fatalf("FetchAllClusterHAStates failed: %v", err)
		}
		err = ha.ValidateMultiClusterLeaseConfigs(states)
		if err == nil || !strings.Contains(err.Error(), "mismatched lease namespace") {
			t.Fatalf("expected mismatched lease namespace error, got %v", err)
		}
	})

	t.Run("duplicate clusterCandidateIdentity fails validation", func(t *testing.T) {
		c1 := newFakeClusterHandle(t, "cluster-a",
			newConfigConnectorWithMCL("kcc-lease", "cnrm-system", "shared-id"),
		)
		c2 := newFakeClusterHandle(t, "cluster-b",
			newConfigConnectorWithMCL("kcc-lease", "cnrm-system", "shared-id"),
		)

		states, err := ha.FetchAllClusterHAStates(ctx, []*ha.ClusterHandle{c1, c2})
		if err != nil {
			t.Fatalf("FetchAllClusterHAStates failed: %v", err)
		}
		err = ha.ValidateMultiClusterLeaseConfigs(states)
		if err == nil || !strings.Contains(err.Error(), "duplicate clusterCandidateIdentity") {
			t.Fatalf("expected duplicate clusterCandidateIdentity error, got %v", err)
		}
	})

	t.Run("missing ConfigConnector fails validation", func(t *testing.T) {
		c1 := newFakeClusterHandle(t, "cluster-a",
			newConfigConnectorWithMCL("kcc-lease", "cnrm-system", "cluster-a-id"),
		)
		c2 := newFakeClusterHandle(t, "cluster-b")

		states, err := ha.FetchAllClusterHAStates(ctx, []*ha.ClusterHandle{c1, c2})
		if err != nil {
			t.Fatalf("FetchAllClusterHAStates failed: %v", err)
		}
		err = ha.ValidateMultiClusterLeaseConfigs(states)
		if err == nil || !strings.Contains(err.Error(), "missing the \"configconnector.core.cnrm.cloud.google.com\"") {
			t.Fatalf("expected missing ConfigConnector error, got %v", err)
		}
	})

	t.Run("nil cluster state fails validation cleanly", func(t *testing.T) {
		err := ha.ValidateMultiClusterLeaseConfigs([]*ha.ClusterHAState{nil, nil})
		if err == nil || !strings.Contains(err.Error(), "cluster HA state is nil") {
			t.Fatalf("expected nil cluster HA state error, got %v", err)
		}
	})
}

func newFakeClusterHandle(t *testing.T, contextName string, objs ...runtime.Object) *ha.ClusterHandle {
	t.Helper()
	scheme, err := ha.NewHAScheme()
	if err != nil {
		t.Fatalf("NewHAScheme() failed: %v", err)
	}
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(objs...).Build()
	return &ha.ClusterHandle{
		ContextName: contextName,
		Client: &kubecli.Client{
			Client: fakeClient,
		},
	}
}

func newConfigConnectorWithMCL(leaseName, namespace, identity string) *operatorv1beta1.ConfigConnector {
	return &operatorv1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name: ha.ConfigConnectorResourceName,
		},
		Spec: operatorv1beta1.ConfigConnectorSpec{
			Experiments: &operatorv1beta1.CCExperiments{
				MultiClusterLease: &operatorv1beta1.MultiClusterLeaseSpec{
					LeaseName:                leaseName,
					Namespace:                namespace,
					ClusterCandidateIdentity: identity,
					ResourceReplicationMode:  "Status",
				},
			},
		},
	}
}
