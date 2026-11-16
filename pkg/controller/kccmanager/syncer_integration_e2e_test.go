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

package kccmanager

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
)

func TestSyncerIntegration_RetryOnConflict(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mclModPath, err := getModulePath("github.com/gke-labs/multicluster-leader-election")
	if err != nil {
		t.Fatalf("error getting module path: %v", err)
	}
	kubeEtlModPath, err := getModulePath("github.com/gke-labs/kube-etl/syncer")
	if err != nil {
		t.Fatalf("error getting module path: %v", err)
	}

	testEnv := &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join(repo.GetRootOrTestFatal(t), "operator/config/crd/bases"),
			filepath.Join(mclModPath, "config/crd/bases"),
			filepath.Join(kubeEtlModPath, "config/crd"),
		},
		ErrorIfCRDPathMissing: true,
	}

	cfg, err := testEnv.Start()
	if err != nil {
		t.Fatalf("error starting envtest: %v", err)
	}
	defer func() {
		if err := testEnv.Stop(); err != nil {
			t.Errorf("error stopping envtest: %v", err)
		}
	}()

	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = mclv1alpha1.AddToScheme(scheme)

	kubeClient, err := client.New(cfg, client.Options{Scheme: scheme})
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_ = kubeClient.Create(ctx, ns)

	name := types.NamespacedName{Name: "test-syncer-retry", Namespace: "default"}

	// Pre-create the Syncer object, pulling from another cluster.
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(KRMSyncerGVK)
	u.SetName(name.Name)
	u.SetNamespace(name.Namespace)
	_ = unstructured.SetNestedField(u.Object, false, "spec", "suspend")
	_ = unstructured.SetNestedMap(u.Object, testRemote(testGKECluster("previous-leader"), "Default"), "spec", "remote")
	_ = unstructured.SetNestedSlice(u.Object, []interface{}{
		map[string]interface{}{
			"group":   "example.com",
			"version": "v1",
			"kind":    "Foo",
		},
	}, "spec", "rules")

	if err := kubeClient.Create(ctx, u); err != nil {
		t.Fatalf("error creating syncer: %v", err)
	}

	cluster := testGKECluster("test-leader")
	si := &SyncerIntegration{
		client:          kubeClient,
		apiReader:       kubeClient, // use non-cached client directly
		leaseNN:         name,
		replicationMode: "Status",
		cluster:         cluster,
	}

	// Fetch the object to get its ResourceVersion
	latest := &unstructured.Unstructured{}
	latest.SetGroupVersionKind(KRMSyncerGVK)
	if err := kubeClient.Get(ctx, name, latest); err != nil {
		t.Fatalf("failed to get latest syncer: %v", err)
	}

	// Update it out-of-band to create a conflict condition for the next update
	outOfBand := latest.DeepCopy()
	_ = unstructured.SetNestedField(outOfBand.Object, "push", "spec", "mode")
	if err := kubeClient.Update(ctx, outOfBand); err != nil {
		t.Fatalf("failed to out-of-band update syncer: %v", err)
	}

	// Call EnsureSuspended. It should fetch, see conflict, retry, and succeed.
	if err := si.EnsureSuspended(ctx); err != nil {
		t.Fatalf("EnsureSuspended failed on conflict: %v", err)
	}

	// Verify it was actually suspended, and that it names this cluster, the current leader.
	updated := &unstructured.Unstructured{}
	updated.SetGroupVersionKind(KRMSyncerGVK)
	if err := kubeClient.Get(ctx, name, updated); err != nil {
		t.Fatalf("failed to get updated syncer: %v", err)
	}
	assertSyncer(t, updated, cluster, "Default", true)
}

// On a standby, the KRMSyncer pulls from the leader's GKE cluster and follows leader changes.
// When this cluster becomes the leader, the KRMSyncer is suspended and names this cluster.
func TestSyncerIntegration_RemoteFollowsLeader(t *testing.T) {
	ctx := context.Background()
	_, kubeClient := startSyncerEnvtest(t)

	ns := "cnrm-system"
	if err := kubeClient.Create(ctx, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}); err != nil {
		t.Fatalf("error creating namespace: %v", err)
	}

	self := testGKECluster("cluster-b")
	nn := types.NamespacedName{Namespace: ns, Name: "test-lease"}
	si := &SyncerIntegration{
		client:             kubeClient,
		apiReader:          kubeClient,
		leaseNN:            nn,
		replicationMode:    "Status",
		cluster:            self,
		leaderEndpointType: operatorv1beta1.LeaderEndpointTypeDNS,
	}
	getSyncer := func() *unstructured.Unstructured {
		t.Helper()
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(KRMSyncerGVK)
		if err := kubeClient.Get(ctx, nn, u); err != nil {
			t.Fatalf("error getting KRMSyncer %s: %v", nn, err)
		}
		return u
	}

	// This cluster is a standby: the KRMSyncer is created and pulls from the leader.
	leaderA := testGKECluster("cluster-a")
	if err := si.EnsurePullingFromLeader(ctx, leaderA); err != nil {
		t.Fatalf("EnsurePullingFromLeader(%s): %v", leaderA.ResourceName(), err)
	}
	assertSyncer(t, getSyncer(), leaderA, "DNS", false)

	// Another cluster becomes the leader: the KRMSyncer is updated to pull from it.
	leaderC := testGKECluster("cluster-c")
	if err := si.EnsurePullingFromLeader(ctx, leaderC); err != nil {
		t.Fatalf("EnsurePullingFromLeader(%s): %v", leaderC.ResourceName(), err)
	}
	assertSyncer(t, getSyncer(), leaderC, "DNS", false)

	// This cluster becomes the leader: the KRMSyncer is suspended and names this cluster.
	if err := si.EnsureSuspended(ctx); err != nil {
		t.Fatalf("EnsureSuspended: %v", err)
	}
	suspended := getSyncer()
	assertSyncer(t, suspended, self, "DNS", true)

	// Suspending again doesn't change the object.
	if err := si.EnsureSuspended(ctx); err != nil {
		t.Fatalf("EnsureSuspended: %v", err)
	}
	if got, want := getSyncer().GetResourceVersion(), suspended.GetResourceVersion(); got != want {
		t.Errorf("EnsureSuspended updated an already suspended KRMSyncer: resourceVersion %s, want %s", got, want)
	}
}

// The ConfigConnector CRD requires gkeCluster and rejects the deprecated clusterCandidateIdentity.
func TestMultiClusterLeaseValidation(t *testing.T) {
	ctx := context.Background()
	_, kubeClient := startSyncerEnvtest(t)

	newCC := func(lease *operatorv1beta1.MultiClusterLeaseSpec) *operatorv1beta1.ConfigConnector {
		return &operatorv1beta1.ConfigConnector{
			ObjectMeta: metav1.ObjectMeta{Name: operatorv1beta1.ConfigConnectorAllowedName},
			Spec: operatorv1beta1.ConfigConnectorSpec{
				Mode:        "cluster",
				Experiments: &operatorv1beta1.CCExperiments{MultiClusterLease: lease},
			},
		}
	}

	for _, tc := range []struct {
		name    string
		lease   *operatorv1beta1.MultiClusterLeaseSpec
		wantErr string
	}{
		{
			name: "clusterCandidateIdentity is rejected",
			lease: &operatorv1beta1.MultiClusterLeaseSpec{
				LeaseName: "lease", Namespace: "cnrm-system",
				ClusterCandidateIdentity: "cluster-a",
				GKECluster:               testGKECluster("cluster-a"),
			},
			wantErr: "clusterCandidateIdentity is no longer supported",
		},
		{
			name:    "gkeCluster is required",
			lease:   &operatorv1beta1.MultiClusterLeaseSpec{LeaseName: "lease", Namespace: "cnrm-system"},
			wantErr: "gkeCluster",
		},
		{
			name: "gkeCluster parts can't contain a slash",
			lease: &operatorv1beta1.MultiClusterLeaseSpec{
				LeaseName: "lease", Namespace: "cnrm-system",
				GKECluster: &operatorv1beta1.GKECluster{Project: "test-project", Location: "us-central1", Name: "a/b"},
			},
			wantErr: "gkeCluster.name",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := kubeClient.Create(ctx, newCC(tc.lease))
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Errorf("creating ConfigConnector: got error %v, want an error containing %q", err, tc.wantErr)
			}
		})
	}

	// A valid gkeCluster is accepted, and leaderEndpointType defaults to Default.
	cc := newCC(&operatorv1beta1.MultiClusterLeaseSpec{
		LeaseName: "lease", Namespace: "cnrm-system",
		GKECluster: testGKECluster("cluster-a"),
	})
	if err := kubeClient.Create(ctx, cc); err != nil {
		t.Fatalf("creating ConfigConnector: %v", err)
	}
	if got := cc.Spec.Experiments.MultiClusterLease.LeaderEndpointType; got != operatorv1beta1.LeaderEndpointTypeDefault {
		t.Errorf("leaderEndpointType = %q, want %q", got, operatorv1beta1.LeaderEndpointTypeDefault)
	}
}
