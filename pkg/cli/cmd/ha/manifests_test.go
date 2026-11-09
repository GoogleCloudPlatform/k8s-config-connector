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
	"slices"
	"strings"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
)

func TestBuildClusterObjectsRBAC(t *testing.T) {
	b := newManifestBuilder(
		"gke_proj-a_us-central1_cluster-a",
		operatorv1beta1.GKECluster{
			Project:  "proj-a",
			Location: "us-central1",
			Name:     "cluster-a",
		},
		"proj-a",
		"gcr.io/gke-release/cnrm/controller:1.158.0",
		"kcc-multicluster-lease",
		"cnrm-system",
		"my-bucket",
	)

	objs := b.buildClusterObjects()
	if len(objs) != 11 {
		t.Fatalf("expected 11 cluster objects, got %d", len(objs))
	}

	clusterRoles := make(map[string]*rbacv1.ClusterRole)
	clusterRoleBindings := make(map[string]*rbacv1.ClusterRoleBinding)

	for _, obj := range objs {
		switch o := obj.(type) {
		case *rbacv1.ClusterRole:
			clusterRoles[o.Name] = o
		case *rbacv1.ClusterRoleBinding:
			clusterRoleBindings[o.Name] = o
		}
	}

	// 1. MCL ClusterRole checks (upstream rules: events, finalizers)
	mclRole, ok := clusterRoles[mclClusterRoleName]
	if !ok {
		t.Fatalf("missing ClusterRole %q", mclClusterRoleName)
	}
	assertHasRule(t, mclRole, []string{""}, []string{"events"}, []string{"create", "patch"})
	assertHasRule(t, mclRole, []string{"multicluster.core.cnrm.cloud.google.com"}, []string{"multiclusterleases", "multiclusterleases/status"}, []string{"get", "list", "watch", "create", "update", "patch", "delete"})
	assertHasRule(t, mclRole, []string{"multicluster.core.cnrm.cloud.google.com"}, []string{"multiclusterleases/finalizers"}, []string{"update"})

	// 2. MCL ClusterRoleBinding checks (bound ONLY to kcc-mcl service account, not cnrm-system group)
	mclBinding, ok := clusterRoleBindings[mclClusterRoleName]
	if !ok {
		t.Fatalf("missing ClusterRoleBinding %q", mclClusterRoleName)
	}
	if mclBinding.RoleRef.Name != mclClusterRoleName {
		t.Errorf("mclBinding roleRef = %q, want %q", mclBinding.RoleRef.Name, mclClusterRoleName)
	}
	assertExactSubjects(t, mclBinding, []rbacv1.Subject{
		{Kind: "ServiceAccount", Name: mclServiceAccountName, Namespace: manifestSystemNamespace},
	})

	// 3. Syncer ClusterRole checks (syncer.gkelabs.io only; no wildcard *.cnrm.cloud.google.com)
	syncerRole, ok := clusterRoles[syncerClusterRoleName]
	if !ok {
		t.Fatalf("missing ClusterRole %q", syncerClusterRoleName)
	}
	assertHasRule(t, syncerRole, []string{"syncer.gkelabs.io"}, []string{"krmsyncers", "krmsyncers/status"}, []string{"get", "list", "watch", "create", "update", "patch", "delete"})
	for _, rule := range syncerRole.Rules {
		for _, g := range rule.APIGroups {
			if strings.Contains(g, "cnrm.cloud.google.com") {
				t.Errorf("syncer ClusterRole must not contain rule for group %q; KCC access must come from cnrm-admin binding", g)
			}
		}
	}

	// 4. Syncer ClusterRoleBinding checks (bound ONLY to kcc-syncer service account)
	syncerBinding, ok := clusterRoleBindings[syncerClusterRoleName]
	if !ok {
		t.Fatalf("missing ClusterRoleBinding %q", syncerClusterRoleName)
	}
	if syncerBinding.RoleRef.Name != syncerClusterRoleName {
		t.Errorf("syncerBinding roleRef = %q, want %q", syncerBinding.RoleRef.Name, syncerClusterRoleName)
	}
	assertExactSubjects(t, syncerBinding, []rbacv1.Subject{
		{Kind: "ServiceAccount", Name: syncerServiceAccountName, Namespace: manifestSystemNamespace},
	})

	// 5. Syncer binding to cnrm-admin
	syncerAdminBinding, ok := clusterRoleBindings[syncerAdminBindingName]
	if !ok {
		t.Fatalf("missing ClusterRoleBinding %q for cnrm-admin binding", syncerAdminBindingName)
	}
	if syncerAdminBinding.RoleRef.Name != cnrmAdminClusterRoleName {
		t.Errorf("syncerAdminBinding roleRef = %q, want %q", syncerAdminBinding.RoleRef.Name, cnrmAdminClusterRoleName)
	}
	assertExactSubjects(t, syncerAdminBinding, []rbacv1.Subject{
		{Kind: "ServiceAccount", Name: syncerServiceAccountName, Namespace: manifestSystemNamespace},
	})

	// 6. Manager HA ClusterRole checks (covering only multiclusterleases and krmsyncers)
	managerHARole, ok := clusterRoles[managerHAClusterRoleName]
	if !ok {
		t.Fatalf("missing ClusterRole %q", managerHAClusterRoleName)
	}
	if len(managerHARole.Rules) != 2 {
		t.Fatalf("managerHARole should have exactly 2 rules, got %d", len(managerHARole.Rules))
	}
	assertHasRule(t, managerHARole, []string{"multicluster.core.cnrm.cloud.google.com"}, []string{"multiclusterleases"}, []string{"get", "list", "watch", "create", "update", "patch", "delete"})
	assertHasRule(t, managerHARole, []string{"syncer.gkelabs.io"}, []string{"krmsyncers"}, []string{"get", "list", "watch", "create", "update", "patch", "delete"})

	// 7. Manager HA ClusterRoleBinding checks (bound to system:serviceaccounts:cnrm-system)
	managerHABinding, ok := clusterRoleBindings[managerHAClusterRoleName]
	if !ok {
		t.Fatalf("missing ClusterRoleBinding %q", managerHAClusterRoleName)
	}
	if managerHABinding.RoleRef.Name != managerHAClusterRoleName {
		t.Errorf("managerHABinding roleRef = %q, want %q", managerHABinding.RoleRef.Name, managerHAClusterRoleName)
	}
	assertExactSubjects(t, managerHABinding, []rbacv1.Subject{
		{Kind: "Group", APIGroup: "rbac.authorization.k8s.io", Name: manifestSystemSAGroup},
	})
}

func TestRenderClusterYAMLIncludesExpectedRBAC(t *testing.T) {
	b := newManifestBuilder(
		"gke_proj-a_us-central1_cluster-a",
		operatorv1beta1.GKECluster{
			Project:  "proj-a",
			Location: "us-central1",
			Name:     "cluster-a",
		},
		"proj-a",
		"gcr.io/gke-release/cnrm/controller:1.158.0",
		"kcc-multicluster-lease",
		"cnrm-system",
		"my-bucket",
	)

	yamlStr, err := b.renderClusterYAML(nil)
	if err != nil {
		t.Fatalf("renderClusterYAML failed: %v", err)
	}

	for _, want := range []string{
		"name: cnrm-mcl-controller",
		"name: cnrm-syncer-controller",
		"name: cnrm-syncer-cnrm-admin",
		"name: cnrm-manager-ha",
		"name: cnrm-admin",
		"system:serviceaccounts:cnrm-system",
		"multiclusterleases/finalizers",
	} {
		if !strings.Contains(yamlStr, want) {
			t.Errorf("rendered YAML missing expected snippet %q", want)
		}
	}

	if strings.Contains(yamlStr, "*.cnrm.cloud.google.com") {
		t.Errorf("rendered YAML contains invalid RBAC apiGroup *.cnrm.cloud.google.com")
	}
}

func assertHasRule(t *testing.T, role *rbacv1.ClusterRole, apiGroups, resources, verbs []string) {
	t.Helper()
	for _, rule := range role.Rules {
		if slices.Equal(rule.APIGroups, apiGroups) && slices.Equal(rule.Resources, resources) && slices.Equal(rule.Verbs, verbs) {
			return
		}
	}
	t.Errorf("ClusterRole %q missing expected rule with APIGroups=%v Resources=%v Verbs=%v", role.Name, apiGroups, resources, verbs)
}

func assertExactSubjects(t *testing.T, binding *rbacv1.ClusterRoleBinding, want []rbacv1.Subject) {
	t.Helper()
	if len(binding.Subjects) != len(want) {
		t.Fatalf("binding %q subjects count = %d, want %d (got %v)", binding.Name, len(binding.Subjects), len(want), binding.Subjects)
	}
	for i, s := range want {
		got := binding.Subjects[i]
		if got.Kind != s.Kind || got.Name != s.Name || got.Namespace != s.Namespace || got.APIGroup != s.APIGroup {
			t.Errorf("binding %q subject[%d] = %+v, want %+v", binding.Name, i, got, s)
		}
	}
}
