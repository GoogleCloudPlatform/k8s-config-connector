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
	"path/filepath"
	"strings"
	"testing"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/kubecli"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	mclclient "github.com/gke-labs/multicluster-leader-election/pkg/client"
	appsv1 "k8s.io/api/apps/v1"
	authorizationv1 "k8s.io/api/authorization/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/yaml"
)

const (
	apiServerLeaseName = "kcc-multicluster-lease"
	// The test cluster is cluster-a in proj-a. Its only peer is a cluster in proj-b.
	apiServerContext  = "gke_proj-a_us-central1_cluster-a"
	apiServerIdentity = "projects/proj-a/locations/us-central1/clusters/cluster-a"
	ownSyncerGSA      = "kcc-syncer@proj-a.iam.gserviceaccount.com"
	peerSyncerGSA     = "kcc-syncer@proj-b.iam.gserviceaccount.com"

	storageAPIGroup   = "storage.cnrm.cloud.google.com"
	kccSystemAPIGroup = "core.cnrm.cloud.google.com"
)

// TestSetupOnAPIServer runs setup's Kubernetes steps against a real API server, and then asks the
// API server what the applied RBAC allows.
func TestSetupOnAPIServer(t *testing.T) {
	server := startAPIServer(t)
	perNamespaceManager := serviceAccount("cnrm-system", "cnrm-controller-manager-team-a")
	clusterManager := serviceAccount("cnrm-system", "cnrm-controller-manager")

	t.Run("unset mode, which the operator runs as namespaced mode", func(t *testing.T) {
		server.createConfigConnector(t)
		server.runSetup(t)
		// The second run updates every object instead of creating it.
		server.runSetup(t)

		server.requireEstablishedCRD(t, "multiclusterleases.multicluster.core.cnrm.cloud.google.com")
		server.requireEstablishedCRD(t, "krmsyncers.syncer.gkelabs.io")
		server.requireMultiClusterLeaseEnabled(t)
		server.requireComponentAccess(t)
		server.requireAccess(t, "a per-namespace manager", perNamespaceManager, true, managerNeeds...)
		server.requireAccess(t, "a per-namespace manager", perNamespaceManager, false, managerDoesNotNeed...)
		// The names of the per-namespace managers depend on the namespaces, so namespaced mode
		// gives this access to every ServiceAccount in cnrm-system, and to no other.
		server.requireAccess(t, "another ServiceAccount in cnrm-system", serviceAccount("cnrm-system", "default"), true, managerNeeds...)
		server.requireAccess(t, "a ServiceAccount outside cnrm-system", serviceAccount("team-a", "default"), false, managerNeeds...)
	})

	t.Run("cluster mode", func(t *testing.T) {
		server.setClusterMode(t)
		server.runSetup(t)

		server.requireComponentAccess(t)
		server.requireAccess(t, "the manager", clusterManager, true, managerNeeds...)
		server.requireAccess(t, "the manager", clusterManager, false, managerDoesNotNeed...)
		server.requireAccess(t, "a per-namespace manager", perNamespaceManager, false, managerNeeds...)
		server.requireAccess(t, "another ServiceAccount in cnrm-system", serviceAccount("cnrm-system", "default"), false, managerNeeds...)
		// In cluster mode, this cluster's syncer can't write the lease or the KRMSyncers.
		server.requireAccess(t, "this cluster's syncer", serviceAccount("cnrm-system", "kcc-syncer"), false,
			access{verb: "patch", group: mclAPIGroup, resource: "multiclusterleases", namespace: "cnrm-system"},
			access{verb: "update", group: syncerAPIGroup, resource: "krmsyncers", namespace: "cnrm-system"},
		)
		server.requireLeaderElectionLockWorksAs(t, clusterManager)
	})
}

// managerNeeds is what Config Connector's manager does for multi-cluster leader election: its
// lock gets, creates and patches the lease, and it gets, creates and updates KRMSyncers.
var managerNeeds = []access{
	{verb: "get", group: mclAPIGroup, resource: "multiclusterleases", namespace: "cnrm-system"},
	{verb: "create", group: mclAPIGroup, resource: "multiclusterleases", namespace: "cnrm-system"},
	{verb: "patch", group: mclAPIGroup, resource: "multiclusterleases", namespace: "cnrm-system"},
	{verb: "get", group: syncerAPIGroup, resource: "krmsyncers", namespace: "cnrm-system"},
	{verb: "create", group: syncerAPIGroup, resource: "krmsyncers", namespace: "cnrm-system"},
	{verb: "update", group: syncerAPIGroup, resource: "krmsyncers", namespace: "cnrm-system"},
}

var managerDoesNotNeed = []access{
	{verb: "list", group: mclAPIGroup, resource: "multiclusterleases", namespace: "cnrm-system"},
	{verb: "watch", group: mclAPIGroup, resource: "multiclusterleases", namespace: "cnrm-system"},
	{verb: "update", group: mclAPIGroup, resource: "multiclusterleases", namespace: "cnrm-system"},
	{verb: "delete", group: mclAPIGroup, resource: "multiclusterleases", namespace: "cnrm-system"},
	{verb: "list", group: syncerAPIGroup, resource: "krmsyncers", namespace: "cnrm-system"},
	{verb: "watch", group: syncerAPIGroup, resource: "krmsyncers", namespace: "cnrm-system"},
	{verb: "patch", group: syncerAPIGroup, resource: "krmsyncers", namespace: "cnrm-system"},
	{verb: "delete", group: syncerAPIGroup, resource: "krmsyncers", namespace: "cnrm-system"},
	// The access is limited to the lease namespace.
	{verb: "get", group: mclAPIGroup, resource: "multiclusterleases", namespace: "team-a"},
	{verb: "get", group: syncerAPIGroup, resource: "krmsyncers", namespace: "team-a"},
}

// requireComponentAccess checks the access of the MCL controller, of this cluster's syncer and of
// the other clusters' syncers. It doesn't depend on the mode, so it looks outside the lease
// namespace (cnrm-system), where namespaced mode gives every ServiceAccount in cnrm-system the
// manager's access.
func (s *apiServer) requireComponentAccess(t *testing.T) {
	t.Helper()

	mclController := serviceAccount("cnrm-system", "kcc-mcl")
	s.requireAccess(t, "the MCL controller", mclController, true,
		access{verb: "update", group: mclAPIGroup, resource: "multiclusterleases", namespace: "team-a"},
		access{verb: "patch", group: mclAPIGroup, resource: "multiclusterleases/status", namespace: "team-a"},
		access{verb: "update", group: mclAPIGroup, resource: "multiclusterleases/finalizers", namespace: "team-a"},
		access{verb: "create", group: "", resource: "events", namespace: "team-a"},
	)
	s.requireAccess(t, "the MCL controller", mclController, false,
		access{verb: "get", group: "", resource: "secrets", namespace: "cnrm-system"},
		access{verb: "patch", group: storageAPIGroup, resource: "storagebuckets", namespace: "team-a"},
	)

	syncer := serviceAccount("cnrm-system", "kcc-syncer")
	s.requireAccess(t, "this cluster's syncer", syncer, true,
		access{verb: "list", group: syncerAPIGroup, resource: "krmsyncers"},
		access{verb: "update", group: syncerAPIGroup, resource: "krmsyncers/status", namespace: "team-a"},
		access{verb: "create", group: "", resource: "namespaces"},
		// It writes the Config Connector resources that it copies from the leader.
		access{verb: "create", group: storageAPIGroup, resource: "storagebuckets", namespace: "team-a"},
		access{verb: "patch", group: storageAPIGroup, resource: "storagebuckets", namespace: "team-a"},
		access{verb: "patch", group: storageAPIGroup, resource: "storagebuckets/status", namespace: "team-a"},
		access{verb: "delete", group: storageAPIGroup, resource: "storagebuckets", namespace: "team-a"},
	)
	s.requireAccess(t, "this cluster's syncer", syncer, false,
		access{verb: "get", group: "", resource: "secrets", namespace: "team-a"},
		access{verb: "update", group: syncerAPIGroup, resource: "krmsyncers", namespace: "team-a"},
		// It can't overwrite the objects that configure this cluster's Config Connector.
		access{verb: "patch", group: kccSystemAPIGroup, resource: "configconnectors"},
		access{verb: "patch", group: kccSystemAPIGroup, resource: "configconnectorcontexts", namespace: "team-a"},
		access{verb: "patch", group: mclAPIGroup, resource: "multiclusterleases", namespace: "team-a"},
	)

	peerSyncer := googleServiceAccount(peerSyncerGSA)
	s.requireAccess(t, "a peer's syncer", peerSyncer, true,
		access{verb: "get", group: storageAPIGroup, resource: "storagebuckets", namespace: "team-a"},
		access{verb: "list", group: storageAPIGroup, resource: "storagebuckets"},
		access{verb: "watch", group: storageAPIGroup, resource: "storagebuckets"},
	)
	// A peer's syncer reads Config Connector resources, and nothing else.
	s.requireAccess(t, "a peer's syncer", peerSyncer, false,
		access{verb: "patch", group: storageAPIGroup, resource: "storagebuckets", namespace: "team-a"},
		access{verb: "list", group: kccSystemAPIGroup, resource: "configconnectors"},
		access{verb: "list", group: kccSystemAPIGroup, resource: "configconnectorcontexts"},
		access{verb: "list", group: mclAPIGroup, resource: "multiclusterleases"},
		access{verb: "list", group: syncerAPIGroup, resource: "krmsyncers"},
		access{verb: "list", group: "", resource: "secrets"},
		access{verb: "list", group: "", resource: "pods"},
		access{verb: "list", group: "", resource: "namespaces"},
	)
	// This cluster's own project has no other cluster in the group, so its syncer isn't a peer.
	s.requireAccess(t, "a syncer that isn't a peer", googleServiceAccount(ownSyncerGSA), false,
		access{verb: "list", group: storageAPIGroup, resource: "storagebuckets"},
	)
}

// requireLeaderElectionLockWorksAs acquires and renews the lease as Config Connector's manager
// does, with the lock that the manager uses and with the manager's access.
func (s *apiServer) requireLeaderElectionLockWorksAs(t *testing.T, manager subject) {
	t.Helper()
	ctx := t.Context()
	s.confirmLeaseHolders(t)
	lock := mclclient.New(s.clientAs(t, manager), apiServerLeaseName, "cnrm-system", apiServerIdentity, time.Second)
	record := resourcelock.LeaderElectionRecord{
		HolderIdentity:       apiServerIdentity,
		LeaseDurationSeconds: 15,
		RenewTime:            metav1.Now(),
	}

	if err := lock.Create(ctx, record); err != nil {
		t.Fatalf("acquiring the lease as the manager: %v", err)
	}
	record.RenewTime = metav1.NewTime(record.RenewTime.Add(time.Second))
	if err := lock.Update(ctx, record); err != nil {
		t.Fatalf("renewing the lease as the manager: %v", err)
	}
	got, _, err := lock.Get(ctx)
	if err != nil {
		t.Fatalf("reading the lease as the manager: %v", err)
	}
	if got.HolderIdentity != apiServerIdentity {
		t.Fatalf("got lease holder %q, want %q", got.HolderIdentity, apiServerIdentity)
	}
}

// apiServer is a kube-apiserver that looks like a cluster with Config Connector installed: it has
// the operator's CRDs and Config Connector's cnrm-admin and cnrm-viewer ClusterRoles. It runs no
// controllers.
type apiServer struct {
	restConfig *rest.Config
	// admin is a cluster-admin client for the built-in types.
	admin client.Client
	// handle is the cluster-admin handle that setup uses.
	handle *clusterHandle
}

func startAPIServer(t *testing.T) *apiServer {
	t.Helper()
	ctx := t.Context()
	root := repo.GetRootOrTestFatal(t)

	env := &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join(root, "operator", "config", "crd", "bases")},
		ErrorIfCRDPathMissing: true,
	}
	restConfig, err := env.Start()
	if err != nil {
		t.Fatalf("starting the API server: %v", err)
	}
	t.Cleanup(func() {
		if err := env.Stop(); err != nil {
			t.Errorf("stopping the API server: %v", err)
		}
	})

	admin, err := client.New(restConfig, client.Options{Scheme: clientgoscheme.Scheme})
	if err != nil {
		t.Fatalf("creating the admin client: %v", err)
	}
	for _, namespace := range []string{"cnrm-system", "team-a"} {
		if err := admin.Create(ctx, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}}); err != nil {
			t.Fatalf("creating namespace %q: %v", namespace, err)
		}
	}
	// Config Connector's aggregate roles, as generated. The release kustomizations add the cnrm-
	// prefix to their names.
	for file, name := range map[string]string{"cnrm_admin.yaml": "cnrm-admin", "cnrm_viewer.yaml": "cnrm-viewer"} {
		path := filepath.Join(root, "config", "installbundle", "components", "clusterroles", file)
		role := &rbacv1.ClusterRole{}
		if err := yaml.Unmarshal(test.MustReadFile(t, path), role); err != nil {
			t.Fatalf("parsing %s: %v", path, err)
		}
		role.Name = name
		if err := admin.Create(ctx, role); err != nil {
			t.Fatalf("creating ClusterRole %q: %v", name, err)
		}
	}

	haScheme, err := newHAScheme()
	if err != nil {
		t.Fatalf("newHAScheme() failed: %v", err)
	}
	haClient, err := client.New(restConfig, client.Options{Scheme: haScheme})
	if err != nil {
		t.Fatalf("creating the setup client: %v", err)
	}

	server := &apiServer{
		restConfig: restConfig,
		admin:      admin,
		handle:     &clusterHandle{contextName: apiServerContext, client: &kubecli.Client{Client: haClient}},
	}
	server.rollOutDeployments(t)
	return server
}

// rollOutDeployments stands in for the Deployment controller: it reports every Deployment in
// cnrm-system as rolled out, as if its pods had started and stayed ready.
func (s *apiServer) rollOutDeployments(t *testing.T) {
	ctx := t.Context()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(20 * time.Millisecond):
			}

			deployments := &appsv1.DeploymentList{}
			if err := s.admin.List(ctx, deployments, client.InNamespace("cnrm-system")); err != nil {
				continue
			}
			for i := range deployments.Items {
				deployment := &deployments.Items[i]
				replicas := *deployment.Spec.Replicas
				if deployment.Status.ObservedGeneration == deployment.Generation && deployment.Status.AvailableReplicas == replicas {
					continue
				}
				deployment.Status = appsv1.DeploymentStatus{
					ObservedGeneration: deployment.Generation,
					Replicas:           replicas,
					UpdatedReplicas:    replicas,
					ReadyReplicas:      replicas,
					AvailableReplicas:  replicas,
				}
				// A failed update is retried on the next pass.
				_ = s.admin.Status().Update(ctx, deployment)
			}
		}
	}()
}

// confirmLeaseHolders stands in for the MCL controller: it confirms the candidate in each lease's
// spec as the global holder.
func (s *apiServer) confirmLeaseHolders(t *testing.T) {
	ctx := t.Context()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(20 * time.Millisecond):
			}

			leases := &mclv1alpha1.MultiClusterLeaseList{}
			if err := s.handle.client.List(ctx, leases); err != nil {
				continue
			}
			for i := range leases.Items {
				lease := &leases.Items[i]
				if lease.Status.ObservedGeneration != nil && *lease.Status.ObservedGeneration == lease.Generation {
					continue
				}
				generation := lease.Generation
				lease.Status.ObservedGeneration = &generation
				lease.Status.GlobalHolderIdentity = lease.Spec.HolderIdentity
				// A failed update is retried on the next pass.
				_ = s.handle.client.Status().Update(ctx, lease)
			}
		}
	}()
}

func (s *apiServer) createConfigConnector(t *testing.T) {
	t.Helper()
	cc := &operatorv1beta1.ConfigConnector{ObjectMeta: metav1.ObjectMeta{Name: configConnectorResourceName}}
	if err := s.handle.client.Create(t.Context(), cc); err != nil {
		t.Fatalf("creating the ConfigConnector: %v", err)
	}
}

func (s *apiServer) setClusterMode(t *testing.T) {
	t.Helper()
	cc, err := fetchConfigConnector(t.Context(), s.handle)
	if err != nil {
		t.Fatalf("fetchConfigConnector failed: %v", err)
	}
	cc.Spec.Mode = "cluster"
	cc.Spec.GoogleServiceAccount = "kcc@proj-a.iam.gserviceaccount.com"
	if err := s.handle.client.Update(t.Context(), cc); err != nil {
		t.Fatalf("switching the ConfigConnector to cluster mode: %v", err)
	}
}

// runSetup applies setup's Kubernetes configuration, as executeSetup does for each cluster.
func (s *apiServer) runSetup(t *testing.T) {
	t.Helper()
	cfg, err := newClusterConfigurer(t.Context(), s.handle, testMCLImage, testSyncerImage, apiServerLeaseName, "cnrm-system", "my-ha-bucket")
	if err != nil {
		t.Fatalf("newClusterConfigurer failed: %v", err)
	}
	cfg.setPeerProjects([]string{"proj-b"})
	cfg.mclRolloutTimeout = time.Minute
	cfg.mclRolloutPollInterval = 20 * time.Millisecond

	if err := cfg.apply(t.Context(), nil, false); err != nil {
		t.Fatalf("setup failed: %v", err)
	}
}

func (s *apiServer) requireEstablishedCRD(t *testing.T, name string) {
	t.Helper()
	err := wait.PollUntilContextTimeout(t.Context(), 100*time.Millisecond, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		crd := &apiextensionsv1.CustomResourceDefinition{}
		if err := s.handle.client.Get(ctx, types.NamespacedName{Name: name}, crd); err != nil {
			return false, err
		}
		for _, condition := range crd.Status.Conditions {
			if condition.Type == apiextensionsv1.Established && condition.Status == apiextensionsv1.ConditionTrue {
				return true, nil
			}
		}
		return false, nil
	})
	if err != nil {
		t.Fatalf("CRD %q was not established: %v", name, err)
	}
}

func (s *apiServer) requireMultiClusterLeaseEnabled(t *testing.T) {
	t.Helper()
	cc, err := fetchConfigConnector(t.Context(), s.handle)
	if err != nil {
		t.Fatalf("fetchConfigConnector failed: %v", err)
	}
	if cc.Spec.Experiments == nil || cc.Spec.Experiments.MultiClusterLease == nil {
		t.Fatalf("expected spec.experiments.multiClusterLease to be set, got %+v", cc.Spec)
	}
	identity, err := cc.Spec.Experiments.MultiClusterLease.CandidateIdentity()
	if err != nil || identity != apiServerIdentity {
		t.Fatalf("got identity %q (error %v), want %q", identity, err, apiServerIdentity)
	}
}

// subject is an identity as the API server sees it.
type subject struct {
	user   string
	groups []string
}

func serviceAccount(namespace, name string) subject {
	return subject{
		user:   "system:serviceaccount:" + namespace + ":" + name,
		groups: []string{"system:serviceaccounts", "system:serviceaccounts:" + namespace, "system:authenticated"},
	}
}

// googleServiceAccount is how a GKE cluster sees a Google service account that authenticates with
// its own credentials, as a syncer does on another cluster.
func googleServiceAccount(email string) subject {
	return subject{user: email, groups: []string{"system:authenticated"}}
}

// access is one request. resource may name a subresource, as in "storagebuckets/status". An empty
// namespace means all namespaces, or a cluster-scoped resource.
type access struct {
	verb      string
	group     string
	resource  string
	namespace string
}

// requireAccess asks the API server's authorizer whether who may do each of the requests.
func (s *apiServer) requireAccess(t *testing.T, description string, who subject, wantAllowed bool, requests ...access) {
	t.Helper()
	for _, request := range requests {
		resource, subresource, _ := strings.Cut(request.resource, "/")
		review := &authorizationv1.SubjectAccessReview{
			Spec: authorizationv1.SubjectAccessReviewSpec{
				User:   who.user,
				Groups: who.groups,
				ResourceAttributes: &authorizationv1.ResourceAttributes{
					Verb:        request.verb,
					Group:       request.group,
					Resource:    resource,
					Subresource: subresource,
					Namespace:   request.namespace,
				},
			},
		}
		if err := s.admin.Create(t.Context(), review); err != nil {
			t.Fatalf("creating a SubjectAccessReview: %v", err)
		}
		if review.Status.Allowed != wantAllowed {
			t.Errorf("%s (%s): %s %s/%s in namespace %q: allowed=%v, want %v",
				description, who.user, request.verb, request.group, request.resource, request.namespace, review.Status.Allowed, wantAllowed)
		}
	}
}

// clientAs returns a client that acts as who, with who's access.
func (s *apiServer) clientAs(t *testing.T, who subject) client.Client {
	t.Helper()
	config := rest.CopyConfig(s.restConfig)
	config.Impersonate = rest.ImpersonationConfig{UserName: who.user, Groups: who.groups}
	c, err := client.New(config, client.Options{Scheme: s.handle.client.Scheme()})
	if err != nil {
		t.Fatalf("creating a client for %s: %v", who.user, err)
	}
	return c
}
