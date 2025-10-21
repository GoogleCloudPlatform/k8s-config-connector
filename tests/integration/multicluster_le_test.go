// Copyright 2025 Google LLC
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

package integration

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	"github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	"github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func TestMultiClusterLeaderElection(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	// Register the v1beta1 scheme
	err := v1beta1.AddToScheme(scheme.Scheme)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	err = v1alpha1.AddToScheme(scheme.Scheme)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	// Setup the envtest environment
	testEnv := &envtest.Environment{
		CRDDirectoryPaths: []string{filepath.Join("..", "..", "config", "crds"), filepath.Join("..", "..", "operator", "config", "crd", "bases")},
	}
	cfg, err := testEnv.Start()
	g.Expect(err).NotTo(gomega.HaveOccurred())
	defer func() {
		g.Expect(testEnv.Stop()).To(gomega.Succeed())
	}()

	// Create a client
	c, err := crclient.New(cfg, crclient.Options{})
	g.Expect(err).NotTo(gomega.HaveOccurred())

	// Create the kcc-system namespace
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "kcc-system",
		},
	}
	g.Expect(c.Create(context.Background(), ns)).To(gomega.Succeed())

	// Create the ConfigConnector object with multi-cluster leader election config
	cc := &v1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name: "configconnector.core.cnrm.cloud.google.com",
		},
		Spec: v1beta1.ConfigConnectorSpec{
			Experiments: &v1beta1.CCExperiments{
				LeaderElection: &v1beta1.LeaderElectionSpec{
					MultiClusterLease: &v1beta1.MultiClusterLeaseSpec{
						LeaseName:      "kcc-leader-lease",
						Namespace:      "kcc-system",
						GlobalLockName: "kcc-global-lock",
					},
				},
			},
		},
	}
	g.Expect(c.Create(context.Background(), cc)).To(gomega.Succeed())

	// Create a manager
	mgr, opts, err := kccmanager.NewManager(context.Background(), cfg, c, "", false, "")
	g.Expect(err).NotTo(gomega.HaveOccurred())

	// Create a new manager with the leader election options
	mgr, err = manager.New(cfg, opts)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	// Start the manager in a goroutine
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		g.Expect(mgr.Start(ctx)).To(gomega.Succeed())
	}()

	leaseName := types.NamespacedName{
		Name:      "kcc-leader-lease",
		Namespace: "kcc-system",
	}

	// Start a mock election controller to update the status
	go func() {
		var lastProcessedGeneration int64 = -1
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(100 * time.Millisecond):
				lease := &v1alpha1.MultiClusterLease{}
				err := c.Get(context.Background(), leaseName, lease)
				if err != nil {
					continue // Lease might not exist yet
				}

				if lease.Generation > lastProcessedGeneration {
					patch := crclient.MergeFrom(lease.DeepCopy())
					lease.Status.GlobalHolderIdentity = lease.Spec.HolderIdentity
					lease.Status.ObservedGeneration = &lease.Generation
					now := time.Now().UTC().Format(time.RFC3339)
					lease.Status.GlobalRenewTime = &now
					if err := c.Status().Patch(context.Background(), lease, patch); err == nil {
						lastProcessedGeneration = lease.Generation
					}
				}
			}
		}
	}()

	// Wait for the manager to be ready
	g.Eventually(func() bool {
		return mgr.GetCache().WaitForCacheSync(ctx)
	}, 10*time.Second).Should(gomega.BeTrue())

	// Assert that the MultiClusterLease object is created
	lease := &v1alpha1.MultiClusterLease{}
	g.Eventually(func() error {
		return c.Get(context.Background(), leaseName, lease)
	}, 10*time.Second).Should(gomega.Succeed())

	// Assert that the holder identity is set
	g.Expect(*lease.Spec.HolderIdentity).To(gomega.Equal("kcc-global-lock"))

	// Assert that the lease is renewed
	initialRenewTime := lease.Spec.RenewTime
	g.Eventually(func() bool {
		err := c.Get(context.Background(), leaseName, lease)
		if err != nil {
			return false
		}
		return lease.Spec.RenewTime.After(initialRenewTime.Time)
	}, 15*time.Second).Should(gomega.BeTrue())
}
