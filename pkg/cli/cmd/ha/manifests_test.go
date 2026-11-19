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
	"path/filepath"
	"slices"
	"strings"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	syncercrd "github.com/gke-labs/kube-etl/syncer/config/crd"
	mclcrd "github.com/gke-labs/multicluster-leader-election/config/crd/bases"
	"github.com/google/go-cmp/cmp"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

// newTestManifestBuilder returns the builder for cluster-a in proj-a, whose only peer is a
// cluster in proj-b.
func newTestManifestBuilder(mode string) *manifestBuilder {
	builder := newManifestBuilder(
		"gke_proj-a_us-central1_cluster-a",
		operatorv1beta1.GKECluster{Project: "proj-a", Location: "us-central1", Name: "cluster-a"},
		"proj-a",
		mode,
		testMCLImage,
		testSyncerImage,
		"kcc-multicluster-lease",
		"cnrm-system",
		"my-ha-bucket",
	)
	builder.setPeerProjects([]string{"proj-b"})
	return builder
}

// TestClusterManifestsGolden covers every manifest except the two CRDs, which are the forks'
// (see TestCRDsComeFirstAndKeepTheForksDefinition).
func TestClusterManifestsGolden(t *testing.T) {
	builder := newTestManifestBuilder("cluster")
	cc := &operatorv1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{Name: "configconnector.core.cnrm.cloud.google.com"},
		Spec:       operatorv1beta1.ConfigConnectorSpec{Mode: "cluster"},
	}
	var objs []client.Object
	for _, obj := range builder.buildClusterObjects() {
		if _, isCRD := obj.(*apiextensionsv1.CustomResourceDefinition); !isCRD {
			objs = append(objs, obj)
		}
	}
	objs = append(objs, builder.buildRenderedConfigConnector(cc))

	var got strings.Builder
	for _, obj := range objs {
		got.WriteString("---\n")
		got.WriteString(mustMarshalManifest(t, obj))
	}

	test.CompareGoldenFile(t, filepath.Join("testdata", "manifests.golden.yaml"), got.String())
}

func TestNamespacedModeOnlyChangesWhoGetsLeaseAccess(t *testing.T) {
	clusterObjs := newTestManifestBuilder("cluster").buildClusterObjects()
	namespacedObjs := newTestManifestBuilder("namespaced").buildClusterObjects()
	if len(clusterObjs) != len(namespacedObjs) {
		t.Fatalf("got %d objects in namespaced mode and %d in cluster mode, want the same count", len(namespacedObjs), len(clusterObjs))
	}

	var changed []client.Object
	for i := range namespacedObjs {
		if mustMarshalManifest(t, clusterObjs[i]) != mustMarshalManifest(t, namespacedObjs[i]) {
			changed = append(changed, namespacedObjs[i])
		}
	}

	if len(changed) != 1 {
		t.Fatalf("got %d objects that differ between the modes, want only the manager's RoleBinding: %v", len(changed), changed)
	}
	binding, ok := changed[0].(*rbacv1.RoleBinding)
	if !ok || binding.Name != "cnrm-manager-ha" {
		t.Fatalf("got changed object %T %q, want RoleBinding cnrm-manager-ha", changed[0], changed[0].GetName())
	}
	wantSubjects := []rbacv1.Subject{{APIGroup: "rbac.authorization.k8s.io", Kind: "Group", Name: "system:serviceaccounts:cnrm-system"}}
	if diff := cmp.Diff(wantSubjects, binding.Subjects); diff != "" {
		t.Fatalf("unexpected subjects in namespaced mode (-want +got):\n%s", diff)
	}
}

func TestCRDsComeFirstAndKeepTheForksDefinition(t *testing.T) {
	forkCRDs := []struct {
		name string
		yaml []byte
	}{
		{name: "multiclusterleases.multicluster.core.cnrm.cloud.google.com", yaml: mclcrd.MultiClusterLeaseCRDYAML},
		{name: "krmsyncers.syncer.gkelabs.io", yaml: syncercrd.KRMSyncerCRDYAML},
	}

	objs := newTestManifestBuilder("cluster").buildClusterObjects()

	for i, forkCRD := range forkCRDs {
		crd, ok := objs[i].(*apiextensionsv1.CustomResourceDefinition)
		if !ok || crd.Name != forkCRD.name {
			t.Fatalf("got object %d = %T %q, want CustomResourceDefinition %q", i, objs[i], objs[i].GetName(), forkCRD.name)
		}
		var got, want map[string]interface{}
		if err := yaml.Unmarshal([]byte(mustMarshalManifest(t, crd)), &got); err != nil {
			t.Fatalf("parsing the rendered CRD %q: %v", forkCRD.name, err)
		}
		if err := yaml.Unmarshal(forkCRD.yaml, &want); err != nil {
			t.Fatalf("parsing the fork's CRD %q: %v", forkCRD.name, err)
		}
		if diff := cmp.Diff(want["spec"], got["spec"]); diff != "" {
			t.Fatalf("the rendered CRD %q differs from the fork's (-fork +rendered):\n%s", forkCRD.name, diff)
		}
	}
	for _, obj := range objs[len(forkCRDs):] {
		if _, isCRD := obj.(*apiextensionsv1.CustomResourceDefinition); isCRD {
			t.Fatalf("CustomResourceDefinition %q comes after other objects, want the CRDs first", obj.GetName())
		}
	}
}

func TestManagerAccessIsAppliedBeforeTheComponentBindings(t *testing.T) {
	objs := newTestManifestBuilder("cluster").buildClusterObjects()

	position := func(kind, name string) int {
		t.Helper()
		index := slices.IndexFunc(objs, func(obj client.Object) bool {
			return obj.GetObjectKind().GroupVersionKind().Kind == kind && obj.GetName() == name
		})
		if index == -1 {
			t.Fatalf("no %s %q among the objects", kind, name)
		}
		return index
	}

	// A cluster that an earlier version of setup configured gives the manager its access through
	// the component bindings. Applying these bindings takes that access away, so the manager's
	// own RoleBinding has to exist by then.
	managerBinding := position("RoleBinding", "cnrm-manager-ha")
	for _, componentBinding := range []string{"cnrm-mcl-controller", "cnrm-syncer-controller"} {
		if componentPosition := position("ClusterRoleBinding", componentBinding); managerBinding > componentPosition {
			t.Fatalf("RoleBinding cnrm-manager-ha is object %d, after ClusterRoleBinding %q (object %d)", managerBinding, componentBinding, componentPosition)
		}
	}
}

func TestEveryObjectIsLabeledAsManagedBySetup(t *testing.T) {
	objs := newTestManifestBuilder("cluster").buildClusterObjects()

	for _, obj := range objs {
		if got := obj.GetLabels()["app.kubernetes.io/managed-by"]; got != "config-connector-ha-setup" {
			t.Errorf("%T %q has label app.kubernetes.io/managed-by=%q, want %q", obj, obj.GetName(), got, "config-connector-ha-setup")
		}
	}
}

// TestCnrmViewerCoversEveryGroupThatASyncerWatches guards the assumption behind the
// cnrm-syncer-peers binding. A standby's syncer watches every group that its
// "*.cnrm.cloud.google.com" rule matches on the leader, and cnrm-viewer is its only access there.
// The syncer stops when one of its watches can't sync, so a group that cnrm-viewer doesn't cover
// would break every standby.
func TestCnrmViewerCoversEveryGroupThatASyncerWatches(t *testing.T) {
	viewer := &rbacv1.ClusterRole{}
	viewerPath := filepath.Join(repo.GetRootOrTestFatal(t), "config", "installbundle", "components", "clusterroles", "cnrm_viewer.yaml")
	if err := yaml.Unmarshal(test.MustReadFile(t, viewerPath), viewer); err != nil {
		t.Fatalf("parsing %s: %v", viewerPath, err)
	}
	readable := map[string]bool{}
	for _, rule := range viewer.Rules {
		readsEverything := slices.Contains(rule.Resources, "*")
		for _, verb := range []string{"get", "list", "watch"} {
			readsEverything = readsEverything && slices.Contains(rule.Verbs, verb)
		}
		if !readsEverything {
			continue
		}
		for _, group := range rule.APIGroups {
			readable[group] = true
		}
	}
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("loading Config Connector's CRDs: %v", err)
	}

	watched := map[string]bool{}
	for _, crd := range crds {
		if syncerWildcardMatches(crd.Spec.Group) {
			watched[crd.Spec.Group] = true
		}
	}

	if len(watched) == 0 {
		t.Fatalf("found no Config Connector resource groups among %d CRDs", len(crds))
	}
	for group := range watched {
		if !readable[group] {
			t.Errorf("cnrm-viewer doesn't grant get, list and watch on group %q, which a syncer watches", group)
		}
	}
	for _, systemGroup := range []string{"core.cnrm.cloud.google.com", "customize.core.cnrm.cloud.google.com", "multicluster.core.cnrm.cloud.google.com"} {
		if syncerWildcardMatches(systemGroup) {
			t.Errorf("the syncer's wildcard matches system group %q, which cnrm-viewer doesn't cover", systemGroup)
		}
	}
}

// syncerWildcardMatches reports whether the syncer's "*.cnrm.cloud.google.com" rule matches an
// API group. It mirrors kccGlobMatchesGroup in the syncer fork
// (third_party/github.com/gke-labs/kube-etl/syncer/controllers/kcc_glob.go).
func syncerWildcardMatches(group string) bool {
	const systemGroup = "core.cnrm.cloud.google.com"
	if !strings.HasSuffix(group, "cnrm.cloud.google.com") {
		return false
	}
	return group != systemGroup && !strings.HasSuffix(group, "."+systemGroup)
}

func mustMarshalManifest(t *testing.T, obj client.Object) string {
	t.Helper()
	raw, err := marshalManifest(obj)
	if err != nil {
		t.Fatalf("marshaling %T %q: %v", obj, obj.GetName(), err)
	}
	return string(raw)
}
