// Copyright 2024 Google LLC
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

package e2e

import (
	"fmt"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"

	"google.com/composition/tests/scenario"
	"google.com/composition/tests/testclient"
	"google.com/composition/tests/utils"
	"k8s.io/apimachinery/pkg/util/rand"
)

// ------------------------ AppTeam --------------------------------------------

func getAppTeamObj(project string) []*unstructured.Unstructured {
	return []*unstructured.Unstructured{
		&unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "facade.facade/v1alpha1",
				"kind":       "AppTeam",
				"metadata": map[string]interface{}{
					"name":      "clearing",
					"namespace": "config-control",
				},
				"spec": map[string]interface{}{
					"project": project,
				},
			},
		},
	}
}

func getAppTeamOutputObjects(project string) []*unstructured.Unstructured {
	gkvnns := []testclient.GVKNN{
		{schema.GroupVersionKind{Group: "composition.google.com", Version: "v1alpha1", Kind: "Context"},
			types.NamespacedName{Name: "context", Namespace: project}},
		{schema.GroupVersionKind{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMServiceAccount"},
			types.NamespacedName{Name: "kcc-" + project, Namespace: "config-control"}},
		{schema.GroupVersionKind{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMPartialPolicy"},
			types.NamespacedName{Name: project + "-sa-workload-identity-binding", Namespace: "config-control"}},
		{schema.GroupVersionKind{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMPartialPolicy"},
			types.NamespacedName{Name: "kcc-owners-permissions-" + project, Namespace: "config-control"}},
		{schema.GroupVersionKind{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "StorageBucket"},
			types.NamespacedName{Name: "test-bucket-" + project, Namespace: project}},
		{schema.GroupVersionKind{Group: "core.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ConfigConnectorContext"},
			types.NamespacedName{Name: "configconnectorcontext.core.cnrm.cloud.google.com", Namespace: project}},
		{schema.GroupVersionKind{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Project"},
			types.NamespacedName{Name: project, Namespace: "config-control"}},
		{schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"},
			types.NamespacedName{Name: project}},
	}

	objs := []*unstructured.Unstructured{}
	for _, gvknn := range gkvnns {
		objs = append(objs, gvknn.MakeObject())
	}

	return objs
}

func TestKCCSampleAppTeam(t *testing.T) {
	//t.Parallel()
	s := scenario.NewKCCSample(t, "AppTeam", "appteam.yaml")
	defer s.Cleanup()
	s.Setup()

	// Apply clearing team facade
	project := fmt.Sprintf("clearing-%s", strings.ToLower(rand.String(8)))
	s.Apply("appteam-object", getAppTeamObj(project))

	// Check plan object has no error
	plan := utils.GetPlanObj("config-control", "appteams-clearing")
	condition := utils.GetReadyCondition("ProcessedAllStages", "")
	s.C.MustHaveCondition(plan, condition, 5*scenario.CompositionReconcile)

	// Since the Plan says it has processed all stages we should validate KCC resources.
	s.Verify("kcc-objects", false, getAppTeamOutputObjects(project))

	// Verify KCC object status ?
}

// ------------------------ CloudSQL --------------------------------------------

func getCloudSqlObj(namespace, name string) []*unstructured.Unstructured {
	return []*unstructured.Unstructured{
		&unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "facade.facade/v1alpha1",
				"kind":       "CloudSQL",
				"metadata": map[string]interface{}{
					"name":      name,
					"namespace": namespace,
				},
				"spec": map[string]interface{}{
					"regions": []string{"us-east1", "us-central1"},
					"name":    "collateral-db",
				},
			},
		},
	}
}

func getCloudSQLOutputObjects(namespace, name string) []*unstructured.Unstructured {
	gkvnns := []testclient.GVKNN{
		{schema.GroupVersionKind{Group: "sql.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SQLInstance"},
			types.NamespacedName{Name: name + "-db-main", Namespace: namespace}},
		{schema.GroupVersionKind{Group: "sql.cnrm.cloud.google.com", Version: "v1beta1", Kind: "SQLInstance"},
			types.NamespacedName{Name: name + "-db-replica-us-central1", Namespace: namespace}},

		{schema.GroupVersionKind{Group: "kms.cnrm.cloud.google.com", Version: "v1beta1", Kind: "KMSKeyRing"},
			types.NamespacedName{Name: "kmscryptokeyring-us-central1", Namespace: namespace}},
		{schema.GroupVersionKind{Group: "kms.cnrm.cloud.google.com", Version: "v1beta1", Kind: "KMSKeyRing"},
			types.NamespacedName{Name: "kmscryptokeyring-us-east1", Namespace: namespace}},
		{schema.GroupVersionKind{Group: "kms.cnrm.cloud.google.com", Version: "v1beta1", Kind: "KMSCryptoKey"},
			types.NamespacedName{Name: "kmscryptokey-enc-us-central1", Namespace: namespace}},
		{schema.GroupVersionKind{Group: "kms.cnrm.cloud.google.com", Version: "v1beta1", Kind: "KMSCryptoKey"},
			types.NamespacedName{Name: "kmscryptokey-enc-us-east1", Namespace: namespace}},

		{schema.GroupVersionKind{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMPolicyMember"},
			types.NamespacedName{Name: "sql-kms-us-east1-policybinding", Namespace: namespace}},
		{schema.GroupVersionKind{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMPolicyMember"},
			types.NamespacedName{Name: "sql-kms-us-central1-policybinding", Namespace: namespace}},

		{schema.GroupVersionKind{Group: "serviceusage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ServiceIdentity"},
			types.NamespacedName{Name: "sqladmin.googleapis.com", Namespace: namespace}},

		{schema.GroupVersionKind{Group: "serviceusage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Service"},
			types.NamespacedName{Name: "cloudkms.googleapis.com", Namespace: namespace}},
		{schema.GroupVersionKind{Group: "serviceusage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Service"},
			types.NamespacedName{Name: "iam.googleapis.com", Namespace: namespace}},
		{schema.GroupVersionKind{Group: "serviceusage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Service"},
			types.NamespacedName{Name: "serviceusage.googleapis.com", Namespace: namespace}},
		{schema.GroupVersionKind{Group: "serviceusage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Service"},
			types.NamespacedName{Name: "sqladmin.googleapis.com", Namespace: namespace}},
	}

	objs := []*unstructured.Unstructured{}
	for _, gvknn := range gkvnns {
		objs = append(objs, gvknn.MakeObject())
	}

	return objs
}

func TestKCCSampleCloudSQL(t *testing.T) {
	//t.Parallel()
	s := scenario.NewKCCSample(t, "CloudSQL", "hasql.yaml")
	defer s.Cleanup()
	s.Setup()

	// Apply cloudsql facade
	namespace := "config-control"
	name := "collateral"
	s.Apply("appteam-object", getCloudSqlObj(namespace, name))

	// Check plan object has no error
	plan := utils.GetPlanObj("config-control", "cloudsqls-collateral")
	condition := utils.GetReadyCondition("ProcessedAllStages", "")
	s.C.MustHaveCondition(plan, condition, 5*scenario.CompositionReconcile)

	// Since the Plan says it has processed all stages we should validate KCC resources.
	s.Verify("kcc-objects", false, getCloudSQLOutputObjects(namespace, name))

	// Verify KCC object status ?
}
