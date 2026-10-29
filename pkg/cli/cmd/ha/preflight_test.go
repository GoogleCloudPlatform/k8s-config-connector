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
	"strings"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/kubecli"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestFetchClusterHAState(t *testing.T) {
	leaderIdentity := "cluster-a-id"
	handle := newFakeClusterHandle(t, "cluster-a",
		newConfigConnectorWithMCL("kcc-lease", "cnrm-system", leaderIdentity),
		&mclv1alpha1.MultiClusterLease{
			ObjectMeta: metav1.ObjectMeta{Name: "kcc-lease", Namespace: "cnrm-system"},
			Status:     mclv1alpha1.MultiClusterLeaseStatus{GlobalHolderIdentity: &leaderIdentity},
		},
		&syncerv1alpha1.KRMSyncer{
			ObjectMeta: metav1.ObjectMeta{Name: "kcc-lease", Namespace: "cnrm-system"},
		},
	)

	state, err := fetchClusterHAState(context.Background(), handle)

	if err != nil {
		t.Fatalf("fetchClusterHAState failed: %v", err)
	}
	if state.multiClusterLease == nil || state.multiClusterLease.Status.GlobalHolderIdentity == nil || *state.multiClusterLease.Status.GlobalHolderIdentity != leaderIdentity {
		t.Fatalf("expected GlobalHolderIdentity %q, got %+v", leaderIdentity, state.multiClusterLease)
	}
	if len(state.krmSyncers) != 1 {
		t.Fatalf("expected 1 KRMSyncer, got %d", len(state.krmSyncers))
	}
}

func TestFetchClusterHAStateToleratesMissingHACRDs(t *testing.T) {
	schemeWithoutHACRDs := runtime.NewScheme()
	if err := operatorv1beta1.AddToScheme(schemeWithoutHACRDs); err != nil {
		t.Fatalf("AddToScheme failed: %v", err)
	}
	cc := newConfigConnectorWithMCL("kcc-lease", "cnrm-system", "cluster-a-id")
	fakeClient := fake.NewClientBuilder().WithScheme(schemeWithoutHACRDs).WithRuntimeObjects(cc).Build()
	handle := &clusterHandle{
		contextName: "cluster-a",
		client:      &kubecli.Client{Client: fakeClient},
	}

	state, err := fetchClusterHAState(context.Background(), handle)

	if err != nil {
		t.Fatalf("expected fetchClusterHAState to tolerate absent HA CRDs, got: %v", err)
	}
	if state.configConnector == nil {
		t.Fatalf("expected ConfigConnector to be populated")
	}
	if state.multiClusterLease != nil || len(state.krmSyncers) != 0 {
		t.Fatalf("expected nil MultiClusterLease and empty KRMSyncers when CRDs are absent, got %+v", state)
	}
}

func TestFetchAllClusterHAStatesErrors(t *testing.T) {
	tests := []struct {
		name              string
		handles           func(t *testing.T) []*clusterHandle
		wantErrSubstrings []string
	}{
		{
			name: "missing ConfigConnector across clusters accumulates errors",
			handles: func(t *testing.T) []*clusterHandle {
				return []*clusterHandle{
					newFakeClusterHandle(t, "cluster-a"),
					newFakeClusterHandle(t, "cluster-b"),
				}
			},
			wantErrSubstrings: []string{"cluster \"cluster-a\"", "cluster \"cluster-b\""},
		},
		{
			name: "empty MultiClusterLease spec fields are accumulated",
			handles: func(t *testing.T) []*clusterHandle {
				return []*clusterHandle{
					newFakeClusterHandle(t, "cluster-a", newConfigConnectorWithMCL("", "", "")),
				}
			},
			wantErrSubstrings: []string{"leaseName", "namespace", "gkeCluster"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handles := tc.handles(t)

			_, err := fetchAllClusterHAStates(context.Background(), handles)

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

func TestValidateMultiClusterLeaseConfigsSuccess(t *testing.T) {
	states := []*clusterHAState{
		{contextName: "cluster-a", configConnector: newConfigConnectorWithMCL("kcc-lease", "cnrm-system", "cluster-a-id")},
		{contextName: "cluster-b", configConnector: newConfigConnectorWithMCL("kcc-lease", "cnrm-system", "cluster-b-id")},
	}

	err := validateMultiClusterLeaseConfigs(states)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateMultiClusterLeaseConfigsErrors(t *testing.T) {
	tests := []struct {
		name              string
		states            []*clusterHAState
		wantErrSubstrings []string
	}{
		{
			name: "mismatched leaseName, namespace, and duplicate gkeCluster include per-cluster values and remediation",
			states: []*clusterHAState{
				{contextName: "cluster-a", configConnector: newConfigConnectorWithMCL("kcc-lease-1", "cnrm-system", "shared")},
				{contextName: "cluster-b", configConnector: newConfigConnectorWithMCL("kcc-lease-2", "other-ns", "shared")},
			},
			wantErrSubstrings: []string{
				`mismatched leaseName across clusters (cluster-a="kcc-lease-1", cluster-b="kcc-lease-2")`,
				`mismatched lease namespace across clusters (cluster-a="cnrm-system", cluster-b="other-ns")`,
				`duplicate gkeCluster across clusters (identity "projects/test-project/locations/us-central1/clusters/shared" shared by [cluster-a, cluster-b])`,
				`leases/<namespace>/<leaseName>`,
			},
		},
		{
			name: "deprecated clusterCandidateIdentity fails validation",
			states: []*clusterHAState{
				{contextName: "cluster-a", configConnector: func() *operatorv1beta1.ConfigConnector {
					cc := newConfigConnectorWithMCL("kcc-lease", "cnrm-system", "cluster-a")
					cc.Spec.Experiments.MultiClusterLease.ClusterCandidateIdentity = "cluster-a-id"
					return cc
				}()},
				{contextName: "cluster-b", configConnector: newConfigConnectorWithMCL("kcc-lease", "cnrm-system", "cluster-b")},
			},
			wantErrSubstrings: []string{
				`cluster "cluster-a"`,
				`clusterCandidateIdentity is no longer supported`,
			},
		},
		{
			name:              "nil cluster states fail cleanly",
			states:            []*clusterHAState{nil, nil},
			wantErrSubstrings: []string{"cluster HA state is nil"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			states := tc.states

			err := validateMultiClusterLeaseConfigs(states)

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

func newFakeClusterHandle(t *testing.T, contextName string, objs ...runtime.Object) *clusterHandle {
	t.Helper()
	scheme, err := newHAScheme()
	if err != nil {
		t.Fatalf("newHAScheme() failed: %v", err)
	}
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(objs...).Build()
	return &clusterHandle{
		contextName: contextName,
		client: &kubecli.Client{
			Client: fakeClient,
		},
	}
}

// newConfigConnectorWithMCL returns a ConfigConnector whose MultiClusterLease runs in the GKE
// cluster clusterName. An empty clusterName leaves gkeCluster unset.
func newConfigConnectorWithMCL(leaseName, namespace, clusterName string) *operatorv1beta1.ConfigConnector {
	var gkeCluster *operatorv1beta1.GKECluster
	if clusterName != "" {
		gkeCluster = &operatorv1beta1.GKECluster{Project: "test-project", Location: "us-central1", Name: clusterName}
	}
	return &operatorv1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name: configConnectorResourceName,
		},
		Spec: operatorv1beta1.ConfigConnectorSpec{
			Experiments: &operatorv1beta1.CCExperiments{
				MultiClusterLease: &operatorv1beta1.MultiClusterLeaseSpec{
					LeaseName:               leaseName,
					Namespace:               namespace,
					GKECluster:              gkeCluster,
					ResourceReplicationMode: "Status",
				},
			},
		},
	}
}
