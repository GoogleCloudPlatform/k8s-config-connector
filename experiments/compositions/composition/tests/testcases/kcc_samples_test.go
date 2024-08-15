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
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/tests/scenario"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/tests/testclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/tests/utils"
	"k8s.io/apimachinery/pkg/util/rand"
)

func gvk(g, v, k string) schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: g, Version: v, Kind: k}
}

// ------------------------ AppTeam --------------------------------------------

func getAppTeamObj(project string) []*unstructured.Unstructured {
	return []*unstructured.Unstructured{
		{
			Object: map[string]interface{}{
				"apiVersion": "facade.facade/v1alpha1",
				"kind":       "AppTeam",
				"metadata": map[string]interface{}{
					"name":      project,
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
		{GroupVersionKind: gvk("composition.google.com", "v1alpha1", "Context"),
			NamespacedName: types.NamespacedName{Name: "context", Namespace: project}},
		{GroupVersionKind: gvk("iam.cnrm.cloud.google.com", "v1beta1", "IAMServiceAccount"),
			NamespacedName: types.NamespacedName{Name: "kcc-" + project, Namespace: "config-control"}},
		{GroupVersionKind: gvk("iam.cnrm.cloud.google.com", "v1beta1", "IAMPartialPolicy"),
			NamespacedName: types.NamespacedName{Name: project + "-sa-workload-identity-binding", Namespace: "config-control"}},
		{GroupVersionKind: gvk("iam.cnrm.cloud.google.com", "v1beta1", "IAMPartialPolicy"),
			NamespacedName: types.NamespacedName{Name: "kcc-owners-permissions-" + project, Namespace: "config-control"}},
		{GroupVersionKind: gvk("storage.cnrm.cloud.google.com", "v1beta1", "StorageBucket"),
			NamespacedName: types.NamespacedName{Name: "test-bucket-" + project, Namespace: project}},
		{GroupVersionKind: gvk("core.cnrm.cloud.google.com", "v1beta1", "ConfigConnectorContext"),
			NamespacedName: types.NamespacedName{Name: "configconnectorcontext.core.cnrm.cloud.google.com", Namespace: project}},
		{GroupVersionKind: gvk("resourcemanager.cnrm.cloud.google.com", "v1beta1", "Project"),
			NamespacedName: types.NamespacedName{Name: project, Namespace: "config-control"}},
		{GroupVersionKind: gvk("", "v1", "Namespace"),
			NamespacedName: types.NamespacedName{Name: project}},
	}

	objs := []*unstructured.Unstructured{}
	for _, gvknn := range gkvnns {
		objs = append(objs, gvknn.MakeObject())
	}

	return objs
}

func TestKCCSampleAppTeam(t *testing.T) {
	//t.Parallel()
	s := scenario.NewFromSample(t, scenario.Sample{Name: "AppTeam", Composition: "appteam.yaml"}, nil, true)
	defer s.Cleanup()
	s.Setup()

	// Apply clearing team facade
	project := fmt.Sprintf("clearing-%s", strings.ToLower(rand.String(8)))
	s.Apply("appteam-object", getAppTeamObj(project))

	// Check plan object has no error
	plan := utils.GetPlanObj("config-control", "appteams-clearing")
	condition := utils.GetReadyCondition("ProcessedAllStages", "")
	s.C.MustHaveCondition(plan, condition, 5*scenario.CompositionReconcileTimeout)

	// Since the Plan says it has processed all stages we should validate KCC resources.
	s.Verify("kcc-objects", false, getAppTeamOutputObjects(project))

	// Verify KCC object status ?
}

// ------------------------ CloudSQL --------------------------------------------

func getCloudSqlObj(namespace, name string) []*unstructured.Unstructured {
	return []*unstructured.Unstructured{
		{
			Object: map[string]interface{}{
				"apiVersion": "facade.compositions.google.com/v1",
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
		{GroupVersionKind: gvk("sql.cnrm.cloud.google.com", "v1beta1", "SQLInstance"),
			NamespacedName: types.NamespacedName{Name: name + "-db-main", Namespace: namespace}},
		{GroupVersionKind: gvk("sql.cnrm.cloud.google.com", "v1beta1", "SQLInstance"),
			NamespacedName: types.NamespacedName{Name: name + "-db-replica-us-central1", Namespace: namespace}},

		{GroupVersionKind: gvk("kms.cnrm.cloud.google.com", "v1beta1", "KMSKeyRing"),
			NamespacedName: types.NamespacedName{Name: "kmscryptokeyring-us-central1", Namespace: namespace}},
		{GroupVersionKind: gvk("kms.cnrm.cloud.google.com", "v1beta1", "KMSKeyRing"),
			NamespacedName: types.NamespacedName{Name: "kmscryptokeyring-us-east1", Namespace: namespace}},
		{GroupVersionKind: gvk("kms.cnrm.cloud.google.com", "v1beta1", "KMSCryptoKey"),
			NamespacedName: types.NamespacedName{Name: "kmscryptokey-enc-us-central1", Namespace: namespace}},
		{GroupVersionKind: gvk("kms.cnrm.cloud.google.com", "v1beta1", "KMSCryptoKey"),
			NamespacedName: types.NamespacedName{Name: "kmscryptokey-enc-us-east1", Namespace: namespace}},

		{GroupVersionKind: gvk("iam.cnrm.cloud.google.com", "v1beta1", "IAMPolicyMember"),
			NamespacedName: types.NamespacedName{Name: "sql-kms-us-east1-policybinding", Namespace: namespace}},
		{GroupVersionKind: gvk("iam.cnrm.cloud.google.com", "v1beta1", "IAMPolicyMember"),
			NamespacedName: types.NamespacedName{Name: "sql-kms-us-central1-policybinding", Namespace: namespace}},

		{GroupVersionKind: gvk("serviceusage.cnrm.cloud.google.com", "v1beta1", "ServiceIdentity"),
			NamespacedName: types.NamespacedName{Name: "sqladmin.googleapis.com", Namespace: namespace}},

		{GroupVersionKind: gvk("serviceusage.cnrm.cloud.google.com", "v1beta1", "Service"),
			NamespacedName: types.NamespacedName{Name: "cloudkms.googleapis.com", Namespace: namespace}},
		{GroupVersionKind: gvk("serviceusage.cnrm.cloud.google.com", "v1beta1", "Service"),
			NamespacedName: types.NamespacedName{Name: "iam.googleapis.com", Namespace: namespace}},
		{GroupVersionKind: gvk("serviceusage.cnrm.cloud.google.com", "v1beta1", "Service"),
			NamespacedName: types.NamespacedName{Name: "serviceusage.googleapis.com", Namespace: namespace}},
		{GroupVersionKind: gvk("serviceusage.cnrm.cloud.google.com", "v1beta1", "Service"),
			NamespacedName: types.NamespacedName{Name: "sqladmin.googleapis.com", Namespace: namespace}},
	}

	objs := []*unstructured.Unstructured{}
	for _, gvknn := range gkvnns {
		objs = append(objs, gvknn.MakeObject())
	}

	return objs
}

func TestKCCSampleCloudSQL(t *testing.T) {
	//t.Parallel()

	// Create a project
	s := scenario.NewFromSample(t,
		scenario.Sample{Name: "CloudSQL", Composition: "hasql.yaml"},
		[]scenario.Sample{
			{Name: "AppTeam", Composition: "appteam.yaml"},
		},
		true,
	)
	defer s.Cleanup()
	s.Setup()

	// TODO: better wait
	time.Sleep(time.Second * 10)

	// ---- Setup App Team for cloudsql user --------
	// Apply clearing team facade
	project := fmt.Sprintf("clearing-%s", strings.ToLower(rand.String(8)))
	s.Apply("appteam-object", getAppTeamObj(project))

	// Check plan object has no error
	planName := fmt.Sprintf("appteams-%s", project)
	plan := utils.GetPlanObj("config-control", planName)
	condition := utils.GetReadyCondition("ProcessedAllStages", "")
	s.C.MustHaveCondition(plan, condition, 5*scenario.CompositionReconcileTimeout)

	// Since the Plan says it has processed all stages we should validate KCC resources.
	s.Verify("appteam-kcc-objects", false, getAppTeamOutputObjects(project))

	// ---- Test CloudSQL -----------------------------
	// Apply cloudsql facade
	namespace := project
	name := "collateral"
	s.Apply("appteam-object", getCloudSqlObj(namespace, name))

	// Check plan object has no error
	cloudsqlPlan := utils.GetPlanObj(namespace, "cloudsqls-collateral")
	condition = utils.GetReadyCondition("ProcessedAllStages", "")
	s.C.MustHaveCondition(cloudsqlPlan, condition, 5*scenario.CompositionReconcileTimeout)

	// Since the Plan says it has processed all stages we should validate KCC resources.
	s.Verify("kcc-objects", false, getCloudSQLOutputObjects(namespace, name))
	// Verify KCC object status ?
}
