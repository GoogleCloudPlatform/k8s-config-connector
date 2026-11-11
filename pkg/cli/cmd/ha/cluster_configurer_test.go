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
	"fmt"
	"testing"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/google/go-cmp/cmp"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	testMCLImage    = "gcr.io/my-project/multiclusterlease:my-tag"
	testSyncerImage = "gcr.io/my-project/krmsyncer:my-tag"
)

func TestClusterConfigurerApplyIsIdempotent(t *testing.T) {
	handle := newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector())
	cfg, err := newClusterConfigurer(context.Background(), handle, testMCLImage, testSyncerImage, "kcc-multicluster-lease", "cnrm-system", "my-ha-bucket")
	if err != nil {
		t.Fatalf("newClusterConfigurer failed: %v", err)
	}
	if err := cfg.apply(context.Background(), nil, false); err != nil {
		t.Fatalf("first apply failed: %v", err)
	}

	err = cfg.apply(context.Background(), nil, false)

	if err != nil {
		t.Fatalf("second idempotent apply failed: %v", err)
	}
}

func TestClusterConfigurerApplySkipsConfigConnectorUpdateOnManifestError(t *testing.T) {
	handle := newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector())
	cfg, err := newClusterConfigurer(context.Background(), handle, testMCLImage, testSyncerImage, "kcc-multicluster-lease", "cnrm-system", "my-ha-bucket")
	if err != nil {
		t.Fatalf("newClusterConfigurer failed: %v", err)
	}
	handle.client.Client = &failingCreateClient{Client: handle.client.Client}

	err = cfg.apply(context.Background(), nil, false)

	if err == nil {
		t.Fatalf("expected apply to fail when manifest creation fails, got nil")
	}
	cc, fetchErr := fetchConfigConnector(context.Background(), handle)
	if fetchErr != nil {
		t.Fatalf("fetchConfigConnector failed: %v", fetchErr)
	}
	if cc.Spec.Experiments != nil && cc.Spec.Experiments.MultiClusterLease != nil {
		t.Fatalf("expected ConfigConnector.spec.experiments.multiClusterLease not to be updated when HA manifests fail, got %+v", cc.Spec.Experiments.MultiClusterLease)
	}
}

func TestClusterConfigurerApplyLeavesConfigConnectorUnchangedWhenMCLControllerIsUnavailable(t *testing.T) {
	stuckPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "multiclusterlease-controller-manager-5d8f7-abcde",
			Namespace: "cnrm-system",
			Labels:    map[string]string{"app": "mcl-controller"},
		},
		Status: corev1.PodStatus{
			ContainerStatuses: []corev1.ContainerStatus{
				{
					Name: "manager",
					State: corev1.ContainerState{
						Waiting: &corev1.ContainerStateWaiting{Reason: "ImagePullBackOff", Message: "Back-off pulling image"},
					},
				},
			},
		},
	}
	handle := newFakeClusterHandleWithPendingDeployments(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector(), stuckPod)
	cfg, err := newClusterConfigurer(context.Background(), handle, testMCLImage, testSyncerImage, "kcc-multicluster-lease", "cnrm-system", "my-ha-bucket")
	if err != nil {
		t.Fatalf("newClusterConfigurer failed: %v", err)
	}
	cfg.mclRolloutTimeout = 50 * time.Millisecond
	cfg.mclRolloutPollInterval = 5 * time.Millisecond
	var stdout bytes.Buffer

	err = cfg.apply(context.Background(), &stdout, false)

	const wantErr = `cluster "gke_proj-a_us-central1_cluster-a": Deployment "multiclusterlease-controller-manager" (namespace "cnrm-system") did not become available within 50ms, so spec.experiments.multiClusterLease was left unchanged on this cluster: 0 of 1 pods run the latest spec; pod "multiclusterlease-controller-manager-5d8f7-abcde" container "manager" is waiting: ImagePullBackOff (Back-off pulling image)`
	if err == nil || err.Error() != wantErr {
		t.Fatalf("got error:\n%v\nwant:\n%s", err, wantErr)
	}
	cc, fetchErr := fetchConfigConnector(context.Background(), handle)
	if fetchErr != nil {
		t.Fatalf("fetchConfigConnector failed: %v", fetchErr)
	}
	if cc.Spec.Experiments != nil && cc.Spec.Experiments.MultiClusterLease != nil {
		t.Fatalf("expected spec.experiments.multiClusterLease to stay unset while the MCL controller is unavailable, got %+v", cc.Spec.Experiments.MultiClusterLease)
	}
	const wantStdout = "cluster \"gke_proj-a_us-central1_cluster-a\": waiting up to 50ms for Deployment \"multiclusterlease-controller-manager\" to become available\n"
	if stdout.String() != wantStdout {
		t.Fatalf("got stdout %q, want %q", stdout.String(), wantStdout)
	}
}

func TestDeploymentRolloutStatus(t *testing.T) {
	one := int32(1)
	tests := []struct {
		name        string
		generation  int64
		status      appsv1.DeploymentStatus
		wantPending string
		wantDone    bool
	}{
		{
			name:        "latest spec not processed yet",
			generation:  2,
			status:      appsv1.DeploymentStatus{ObservedGeneration: 1, Replicas: 1, UpdatedReplicas: 1, AvailableReplicas: 1},
			wantPending: "the Deployment controller has not processed the latest spec yet",
		},
		{
			name:        "no pod of the latest spec yet",
			generation:  2,
			status:      appsv1.DeploymentStatus{ObservedGeneration: 2},
			wantPending: "0 of 1 pods run the latest spec",
		},
		{
			name:        "pod of an older spec still terminating",
			generation:  2,
			status:      appsv1.DeploymentStatus{ObservedGeneration: 2, Replicas: 2, UpdatedReplicas: 1, AvailableReplicas: 1},
			wantPending: "1 pods of an older spec are still terminating",
		},
		{
			name:        "new pod not available",
			generation:  2,
			status:      appsv1.DeploymentStatus{ObservedGeneration: 2, Replicas: 1, UpdatedReplicas: 1},
			wantPending: "0 of 1 pods are available",
		},
		{
			name:       "rolled out",
			generation: 2,
			status:     appsv1.DeploymentStatus{ObservedGeneration: 2, Replicas: 1, UpdatedReplicas: 1, AvailableReplicas: 1},
			wantDone:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			deployment := &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{Generation: tc.generation},
				Spec:       appsv1.DeploymentSpec{Replicas: &one},
				Status:     tc.status,
			}

			gotPending, gotDone := deploymentRolloutStatus(deployment)

			if gotPending != tc.wantPending || gotDone != tc.wantDone {
				t.Fatalf("got (%q, %v), want (%q, %v)", gotPending, gotDone, tc.wantPending, tc.wantDone)
			}
		})
	}
}

func TestClusterConfigurerApplyGrantsLeaseAccessToTheManagersOfEachMode(t *testing.T) {
	managerServiceAccount := []rbacv1.Subject{{Kind: "ServiceAccount", Name: "cnrm-controller-manager", Namespace: "cnrm-system"}}
	everyServiceAccountInSystemNamespace := []rbacv1.Subject{{APIGroup: "rbac.authorization.k8s.io", Kind: "Group", Name: "system:serviceaccounts:cnrm-system"}}
	tests := []struct {
		name         string
		mode         string
		wantSubjects []rbacv1.Subject
	}{
		{name: "cluster mode", mode: "cluster", wantSubjects: managerServiceAccount},
		{name: "namespaced mode", mode: "namespaced", wantSubjects: everyServiceAccountInSystemNamespace},
		// The operator runs Config Connector in namespaced mode when spec.mode is unset.
		{name: "unset mode", mode: "", wantSubjects: everyServiceAccountInSystemNamespace},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cc := newBaseConfigConnector()
			cc.Spec.Mode = tc.mode
			handle := newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", cc)
			cfg, err := newClusterConfigurer(context.Background(), handle, testMCLImage, testSyncerImage, "kcc-multicluster-lease", "lease-ns", "my-ha-bucket")
			if err != nil {
				t.Fatalf("newClusterConfigurer failed: %v", err)
			}

			err = cfg.apply(context.Background(), nil, false)

			if err != nil {
				t.Fatalf("apply failed: %v", err)
			}
			binding := &rbacv1.RoleBinding{}
			if err := handle.client.Get(context.Background(), types.NamespacedName{Namespace: "lease-ns", Name: "cnrm-manager-ha"}, binding); err != nil {
				t.Fatalf("getting RoleBinding cnrm-manager-ha in the lease namespace: %v", err)
			}
			if diff := cmp.Diff(tc.wantSubjects, binding.Subjects); diff != "" {
				t.Fatalf("unexpected subjects (-want +got):\n%s", diff)
			}
		})
	}
}

func TestClusterConfigurerApplyPreservesExistingMultiClusterLeaseFields(t *testing.T) {
	existingGKECluster := operatorv1beta1.GKECluster{Project: "custom-proj", Location: "us-west1", Name: "custom-cluster"}
	existingCC := newBaseConfigConnector()
	existingCC.Spec.Experiments = &operatorv1beta1.CCExperiments{
		MultiClusterLease: &operatorv1beta1.MultiClusterLeaseSpec{
			LeaseName:               "old-lease",
			Namespace:               "old-ns",
			GKECluster:              &existingGKECluster,
			ResourceReplicationMode: "Full",
		},
	}
	handle := newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", existingCC)
	cfg, err := newClusterConfigurer(context.Background(), handle, testMCLImage, testSyncerImage, "kcc-multicluster-lease", "cnrm-system", "my-ha-bucket")
	if err != nil {
		t.Fatalf("newClusterConfigurer failed: %v", err)
	}

	err = cfg.apply(context.Background(), nil, false)

	if err != nil {
		t.Fatalf("apply failed: %v", err)
	}
	cc, fetchErr := fetchConfigConnector(context.Background(), handle)
	if fetchErr != nil {
		t.Fatalf("fetchConfigConnector failed: %v", fetchErr)
	}
	gotMCL := cc.Spec.Experiments.MultiClusterLease
	if gotMCL.LeaseName != "kcc-multicluster-lease" || gotMCL.Namespace != "cnrm-system" {
		t.Fatalf("got lease %s/%s, want cnrm-system/kcc-multicluster-lease", gotMCL.Namespace, gotMCL.LeaseName)
	}
	if gotMCL.GKECluster == nil || *gotMCL.GKECluster != existingGKECluster || gotMCL.ResourceReplicationMode != "Full" {
		t.Fatalf("expected existing gkeCluster and replication mode to be preserved, got %+v", gotMCL)
	}
}

func TestClusterConfigurerApplyReplacesDeprecatedClusterCandidateIdentity(t *testing.T) {
	existingCC := newBaseConfigConnector()
	existingCC.Spec.Experiments = &operatorv1beta1.CCExperiments{
		MultiClusterLease: &operatorv1beta1.MultiClusterLeaseSpec{
			LeaseName:                "kcc-multicluster-lease",
			Namespace:                "cnrm-system",
			ClusterCandidateIdentity: "legacy-identity-a",
		},
	}
	handle := newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", existingCC)
	cfg, err := newClusterConfigurer(context.Background(), handle, testMCLImage, testSyncerImage, "kcc-multicluster-lease", "cnrm-system", "my-ha-bucket")
	if err != nil {
		t.Fatalf("newClusterConfigurer failed: %v", err)
	}

	err = cfg.apply(context.Background(), nil, false)

	if err != nil {
		t.Fatalf("apply failed: %v", err)
	}
	cc, fetchErr := fetchConfigConnector(context.Background(), handle)
	if fetchErr != nil {
		t.Fatalf("fetchConfigConnector failed: %v", fetchErr)
	}
	gotMCL := cc.Spec.Experiments.MultiClusterLease
	if gotMCL.ClusterCandidateIdentity != "" {
		t.Fatalf("expected deprecated clusterCandidateIdentity to be cleared, got %q", gotMCL.ClusterCandidateIdentity)
	}
	identity, err := gotMCL.CandidateIdentity()
	if err != nil {
		t.Fatalf("CandidateIdentity() failed: %v", err)
	}
	if want := "projects/proj-a/locations/us-central1/clusters/cluster-a"; identity != want {
		t.Fatalf("got identity %q, want %q", identity, want)
	}
}

type failingCreateClient struct {
	client.Client
}

func (f *failingCreateClient) Create(_ context.Context, _ client.Object, _ ...client.CreateOption) error {
	return fmt.Errorf("simulated create failure")
}
