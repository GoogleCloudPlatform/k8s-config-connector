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
	"testing"
	"time"

	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TestFailoverOnAPIServers runs failover against two real API servers. The fleet stands in for
// the controllers, as in the other failover tests, but the clusters' objects are on the API
// servers. That shows what a fake client can't: the API server takes failover's patches, counts
// the generations that failover waits for, and finds the managers and their pods with failover's
// selectors.
func TestFailoverOnAPIServers(t *testing.T) {
	scheme, err := newHAScheme()
	if err != nil {
		t.Fatalf("newHAScheme() failed: %v", err)
	}
	fleet := newEmptyHAFleet(t)
	for _, name := range []string{"cluster-a", "cluster-b"} {
		server := startAPIServer(t)
		server.installHACRDs(t)
		backend, err := client.NewWithWatch(server.restConfig, client.Options{Scheme: scheme})
		if err != nil {
			t.Fatalf("creating the client of %s: %v", name, err)
		}
		c := fleet.addCluster(name, backend)
		c.create(newConfigConnectorWithMCL(fleetLeaseName, fleetLeaseNamespace, name))
		c.create(&mclv1alpha1.MultiClusterLease{ObjectMeta: metav1.ObjectMeta{Namespace: fleetLeaseNamespace, Name: fleetLeaseName}})
		// Namespaced mode. The manager of team-b has a namespace of its own.
		c.create(&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "team-b-system"}})
		c.create(newManagerStatefulSet("cnrm-system", "cnrm-controller-manager-team-a", "team-a"))
		c.create(newManagerStatefulSet("team-b-system", "cnrm-controller-manager-team-b", "team-b"))
	}
	fleet.settle()
	a, b := fleet.cluster("cluster-a"), fleet.cluster("cluster-b")
	fleet.requireLeader(a)

	// The fleet takes longer over each step on an API server than on a fake client.
	out, err := fleet.failOver("cluster-b", func(opts *failoverOptions) {
		opts.timeout = time.Minute
		opts.leaderAgreementTimeout = time.Minute
	})
	if err != nil {
		t.Fatalf("failover failed: %v\noutput:\n%s", err, out)
	}

	fleet.requireLeader(b)
	for _, c := range fleet.clusters {
		c.requireManager("cnrm-system", "cnrm-controller-manager-team-a", 1)
		c.requireManager("team-b-system", "cnrm-controller-manager-team-b", 1)
	}
	fleet.requirePatches(
		`stop cluster-a cnrm-system/cnrm-controller-manager-team-a (annotation "1") while cluster-a leads and cluster-a,cluster-b run`,
		`stop cluster-a team-b-system/cnrm-controller-manager-team-b (annotation "1") while cluster-a leads and cluster-a,cluster-b run`,
		`start cluster-a cnrm-system/cnrm-controller-manager-team-a (replicas 1, annotation "") while cluster-b leads`,
		`start cluster-a team-b-system/cnrm-controller-manager-team-b (replicas 1, annotation "") while cluster-b leads`,
	)
}

// installHACRDs installs the CRDs that setup installs.
func (s *apiServer) installHACRDs(t *testing.T) {
	t.Helper()
	for _, crd := range buildCRDObjects() {
		if err := upsertObject(t.Context(), s.handle.client, crd); err != nil {
			t.Fatalf("installing CRD %s: %v", crd.GetName(), err)
		}
		s.requireEstablishedCRD(t, crd.GetName())
	}
}
