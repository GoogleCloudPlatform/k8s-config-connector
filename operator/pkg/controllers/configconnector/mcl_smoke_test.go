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

package configconnector

import (
	"context"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/main"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestMCLSmoke(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer stop()
	c := mgr.GetClient()

	// Create the cnrm-system namespace
	cnrmSystemNs := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "cnrm-system"}}
	if err := c.Create(ctx, cnrmSystemNs); err != nil {
		t.Fatalf("error creating cnrm-system namespace: %v", err)
	}

	// Create a minimal ConfigConnector object with MCL enabled
	cc := &v1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name: "configconnector.core.cnrm.cloud.google.com",
		},
		Spec: v1beta1.ConfigConnectorSpec{
			Mode: "cluster",
			Experiments: &v1beta1.CCExperiments{
				MultiClusterLease: &v1beta1.MultiClusterLeaseSpec{
					LeaseName:                "test-lease",
					Namespace:                "cnrm-system",
					ClusterCandidateIdentity: "test-cluster",
				},
			},
		},
	}
	// Apply the ConfigConnector object
	if err := c.Create(ctx, cc); err != nil {
		t.Fatalf("error creating ConfigConnector: %v", err)
	}

	// Verify MCL components are deployed
	mclDeployment := &appsv1.Deployment{}
	for i := 0; i < 10; i++ {
		err := c.Get(ctx, client.ObjectKey{Namespace: "cnrm-system", Name: "multiclusterlease-controller-manager"}, mclDeployment)
		if err == nil {
			break
		}
		if !apierrors.IsNotFound(err) {
			t.Fatalf("error getting mcl-controller deployment: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
	if mclDeployment.Name == "" {
		t.Fatalf("mcl-controller deployment not found")
	}

	// Verify Syncer components are deployed
	syncerDeployment := &appsv1.Deployment{}
	for i := 0; i < 10; i++ {
		err := c.Get(ctx, client.ObjectKey{Namespace: "cnrm-system", Name: "syncer-controller-manager"}, syncerDeployment)
		if err == nil {
			break
		}
		if !apierrors.IsNotFound(err) {
			t.Fatalf("error getting syncer-controller deployment: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
	if syncerDeployment.Name == "" {
		t.Fatalf("syncer-controller deployment not found")
	}

	// Disable the experiment and verify cleanup
	if err := c.Get(ctx, client.ObjectKey{Name: cc.Name}, cc); err != nil {
		t.Fatalf("error getting ConfigConnector: %v", err)
	}
	cc.Spec.Experiments = nil
	if err := c.Update(ctx, cc); err != nil {
		t.Fatalf("error updating ConfigConnector: %v", err)
	}

	// Verify MCL components are deleted
	if err := c.Get(ctx, client.ObjectKey{Namespace: "cnrm-system", Name: "multiclusterlease-controller-manager"}, mclDeployment); !apierrors.IsNotFound(err) {
		t.Errorf("expected mcl-controller deployment to be deleted, but it still exists")
	}

	// Verify Syncer components are deleted
	if err := c.Get(ctx, client.ObjectKey{Namespace: "cnrm-system", Name: "syncer-controller-manager"}, syncerDeployment); !apierrors.IsNotFound(err) {
		t.Errorf("expected syncer-controller deployment to be deleted, but it still exists")
	}
}

func TestMCLSmoke_NamespacedMode(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer stop()
	c := mgr.GetClient()

	// Create the cnrm-system namespace
	cnrmSystemNs := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "cnrm-system"}}
	if err := c.Create(ctx, cnrmSystemNs); err != nil {
		t.Fatalf("error creating cnrm-system namespace: %v", err)
	}

	// Create a minimal ConfigConnector object with MCL enabled
	cc := &v1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name: "configconnector.core.cnrm.cloud.google.com",
		},
		Spec: v1beta1.ConfigConnectorSpec{
			Mode: "namespaced",
			Experiments: &v1beta1.CCExperiments{
				MultiClusterLease: &v1beta1.MultiClusterLeaseSpec{
					LeaseName:                "test-lease",
					Namespace:                "cnrm-system",
					ClusterCandidateIdentity: "test-cluster",
				},
			},
		},
	}
	if err := c.Create(ctx, cc); err != nil {
		t.Fatalf("error creating ConfigConnector: %v", err)
	}

	// Create the tenant-a namespace
	tenantANs := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "tenant-a"}}
	if err := c.Create(ctx, tenantANs); err != nil {
		t.Fatalf("error creating tenant-a namespace: %v", err)
	}

	// Create a CCC to trigger namespaced manager creation
	ccc := &v1beta1.ConfigConnectorContext{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "configconnectorcontext.core.cnrm.cloud.google.com",
			Namespace: "tenant-a",
		},
		Spec: v1beta1.ConfigConnectorContextSpec{
			GoogleServiceAccount: "test-gsa@test.iam.gserviceaccount.com",
		},
	}
	if err := c.Create(ctx, ccc); err != nil {
		t.Fatalf("error creating ConfigConnectorContext: %v", err)
	}

	// Verify namespaced manager is configured for MCL
	mgrSts := &appsv1.StatefulSet{}
	for i := 0; i < 10; i++ {
		err := c.Get(ctx, client.ObjectKey{Namespace: "cnrm-system", Name: "cnrm-controller-manager-tenant-a"}, mgrSts)
		if err == nil {
			break
		}
		if !apierrors.IsNotFound(err) {
			t.Fatalf("error getting namespaced manager statefulset: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
	if mgrSts.Name == "" {
		t.Fatalf("namespaced manager statefulset not found")
	}
	args := mgrSts.Spec.Template.Spec.Containers[0].Args
	hasLeaderElectionArg := false
	hasSyncingModeArg := false
	for _, arg := range args {
		if arg == "--leader-election-type=multicluster" {
			hasLeaderElectionArg = true
		}
		if arg == "--syncing-mode=pull" {
			hasSyncingModeArg = true
		}
	}
	if !hasLeaderElectionArg {
		t.Errorf("expected --leader-election-type=multicluster arg, but it was not found")
	}
	if !hasSyncingModeArg {
		t.Errorf("expected --syncing-mode=pull arg, but it was not found")
	}
	annotations := mgrSts.Spec.Template.Annotations
	if annotations["cnrm.cloud.google.com/lease-name"] != "test-lease" {
		t.Errorf("unexpected lease-name annotation: got %v, want test-lease", annotations["cnrm.cloud.google.com/lease-name"])
	}
}
