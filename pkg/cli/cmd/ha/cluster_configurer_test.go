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
	"fmt"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestClusterConfigurerApplyIsIdempotent(t *testing.T) {
	const kccImage = "gcr.io/gke-release/cnrm/controller:1.158.0"
	handle := newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector(), newManagerStatefulSet(kccImage))
	cfg, err := newClusterConfigurer(context.Background(), handle, "kcc-multicluster-lease", "cnrm-system", "my-ha-bucket")
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
	const kccImage = "gcr.io/gke-release/cnrm/controller:1.158.0"
	handle := newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", newBaseConfigConnector(), newManagerStatefulSet(kccImage))
	cfg, err := newClusterConfigurer(context.Background(), handle, "kcc-multicluster-lease", "cnrm-system", "my-ha-bucket")
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

func TestClusterConfigurerApplyPreservesExistingMultiClusterLeaseFields(t *testing.T) {
	const kccImage = "gcr.io/gke-release/cnrm/controller:1.158.0"
	existingCC := newBaseConfigConnector()
	existingCC.Spec.Experiments = &operatorv1beta1.CCExperiments{
		MultiClusterLease: &operatorv1beta1.MultiClusterLeaseSpec{
			LeaseName:                "old-lease",
			Namespace:                "old-ns",
			ClusterCandidateIdentity: "custom-identity-a",
			ResourceReplicationMode:  "Full",
		},
	}
	handle := newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", existingCC, newManagerStatefulSet(kccImage))
	cfg, err := newClusterConfigurer(context.Background(), handle, "kcc-multicluster-lease", "cnrm-system", "my-ha-bucket")
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
	if gotMCL.ClusterCandidateIdentity != "custom-identity-a" || gotMCL.ResourceReplicationMode != "Full" {
		t.Fatalf("expected existing identity and replication mode to be preserved, got %+v", gotMCL)
	}
}

func TestDiscoverControllerImagePrefersPrimaryManagerStatefulSet(t *testing.T) {
	namespacedSts := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cnrm-controller-manager-ns1",
			Namespace: controllerManagerNamespace,
			Labels:    map[string]string{controllerComponentLabelKey: controllerManagerStatefulSet},
		},
		Spec: appsv1.StatefulSetSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{Name: "manager", Image: "gcr.io/old:1.0"}},
				},
			},
		},
	}
	primarySts := newManagerStatefulSet("gcr.io/primary:2.0")
	handle := newFakeClusterHandle(t, "gke_proj-a_us-central1_cluster-a", namespacedSts, primarySts)

	got, err := discoverControllerImage(context.Background(), handle)

	if err != nil {
		t.Fatalf("discoverControllerImage failed: %v", err)
	}
	if got != "gcr.io/primary:2.0" {
		t.Fatalf("got image %q, want %q", got, "gcr.io/primary:2.0")
	}
}

type failingCreateClient struct {
	client.Client
}

func (f *failingCreateClient) Create(_ context.Context, _ client.Object, _ ...client.CreateOption) error {
	return fmt.Errorf("simulated create failure")
}
