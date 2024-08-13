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
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/tests/scenario"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/tests/utils"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestSimpleCompositionCreate(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()
}

func TestSimpleExpansionJob(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	s.VerifyOutputExists()
	s.VerifyOutputSpecMatches()
}

func TestSimpleExpansionGrpc(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	s.VerifyOutputExists()
	s.VerifyOutputSpecMatches()
}

func TestSimpleCompositionUpdate(t *testing.T) {
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Make sure the pre-update CRs have been created to ensure the coming update triggers a new reconcile.
	cmNames := []string{"p-proj-a", "p-proj-b", "s-proj-a", "s-proj-b"}
	cms := make([]*unstructured.Unstructured, 0)
	for _, cmName := range cmNames {
		cms = append(cms, utils.GetConfigMapObj("team-a", cmName))
	}
	s.C.MustExist(cms, scenario.ExistTimeout)

	// Apply the modified Composition
	s.ApplyManifests("modified composition", "modified_composition.yaml")

	// Changing the composition should trigger the expander to re-reconcile all objects.
	s.VerifyOutputExists()
	s.VerifyOutputSpecMatches()
}

func TestSimpleDeleteFacade(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	s.VerifyOutputExists()

	// Delete Facade CR
	cr := utils.GetUnstructuredObj("facade.foocorp.com", "v1alpha1", "PConfig", "team-a", "team-a-config")
	s.C.MustDelete(cr)

	// Check if Plan is deleted
	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	s.C.MustNotExist([]*unstructured.Unstructured{plan}, scenario.DeleteTimeout)

	// Check if expanded ConfigMap is also deleted
	cm := utils.GetConfigMapObj("team-a", "proj-a")
	s.C.MustNotExist([]*unstructured.Unstructured{cm}, scenario.DeleteTimeout)
}

// Test adding config that results in additional expanded resources
func TestSimpleAddFacadeField(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	s.VerifyOutputExists()

	// Update Facade CR to add a field/list entry
	t.Log("Adding proj-b entry to Facade")
	facade := utils.GetUnstructuredObj("facade.foocorp.com", "v1alpha1", "PConfig", "team-a", "team-a-config")
	addProject := map[string]any{
		"op":    "add",
		"path":  "/spec/projects/-",
		"value": "proj-b",
	}
	s.C.MustJSONPatch(facade, addProject)

	// Check if additional ConfigMap is created
	cm := utils.GetConfigMapObj("team-a", "proj-b")
	s.C.MustExist([]*unstructured.Unstructured{cm}, scenario.CompositionReconcileTimeout)
}

// Test removing config that results in removal of some expanded resource
func TestSimpleDeleteFacadeField(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	s.VerifyOutputExists()

	// Update Facade CR to add a field/list entry
	t.Log("Removing proj-b entry from Facade")
	facade := utils.GetUnstructuredObj("facade.foocorp.com", "v1alpha1", "PConfig", "team-a", "team-a-config")
	addProject := map[string]any{
		"op":   "remove",
		"path": "/spec/projects/1",
	}
	s.C.MustJSONPatch(facade, addProject)

	// Check if the second ConfigMap is removed
	cm := utils.GetConfigMapObj("team-a", "proj-b")
	s.C.MustNotExist([]*unstructured.Unstructured{cm}, scenario.DeleteTimeout)
}

func TestSimpleCompositionStatusValidation(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify there is a Validation failure status condition in compositions
	composition := utils.GetCompositionObj("default", "projectconfigmap")
	condition := utils.GetValidationFailedCondition("ExpanderValidationFailed", "")
	s.C.MustHaveCondition(composition, condition, scenario.CompositionReconcileTimeout)

	// Apply the fixed Composition
	s.ApplyManifests("composition without validation error", "fixed_composition.yaml")

	// Check if Validation failure condition is cleared
	composition = utils.GetCompositionObj("default", "projectconfigmap")
	condition = utils.GetValidationFailedCondition("ExpanderValidationFailed", "")
	s.C.MustNotHaveCondition(composition, condition, scenario.CompositionReconcileTimeout)
}

func TestSimpleCompositionJinjaValidationFailure(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify there is an Error capturing the expander config Validation
	composition := utils.GetCompositionObj("default", "projectconfigmap")
	condition := utils.GetErrorCondition("ValidationFailed", "")
	s.C.MustHaveCondition(composition, condition, scenario.CompositionReconcileTimeout)

	// Apply the fixed Composition
	s.ApplyManifests("composition without validation error", "fixed_composition.yaml")

	// Check if Validation failure condition is cleared
	composition = utils.GetCompositionObj("default", "projectconfigmap")
	condition = utils.GetErrorCondition("ValidationFailed", "")
	s.C.MustNotHaveCondition(composition, condition, scenario.CompositionReconcileTimeout)
}

func TestSimpleCompositionStatusFacadeCRDMissing(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify there is a Error status condition in compositions
	composition := utils.GetCompositionObj("default", "projectconfigmap")
	condition := utils.GetErrorCondition("MissingFacadeCRD", "")
	s.C.MustHaveCondition(composition, condition, scenario.CompositionReconcileTimeout)

	// Apply the fixed Composition
	s.ApplyManifests("Facade CRD", "facade_crd.yaml")

	// Check if Validation failure condition is cleared
	composition = utils.GetCompositionObj("default", "projectconfigmap")
	condition = utils.GetErrorCondition("MissingFacadeCRD", "")
	s.C.MustNotHaveCondition(composition, condition, scenario.CompositionReconcileTimeout)
}

func TestSimplePlanStatusWaitingForValues(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify there is a Waiting failure status condition in plan
	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition := utils.GetWaitingCondition("EvaluateStatusWait", "")
	s.C.MustHaveCondition(plan, condition, scenario.CompositionReconcileTimeout)

	// Apply Configmap with data present
	s.ApplyManifests("Fixed Configmap", "configmap_with_data.yaml")

	// Check if Waiting failure condition is cleared
	plan = utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition = utils.GetWaitingCondition("EvaluateStatusWait", "")
	s.C.MustNotHaveCondition(plan, condition, 3*scenario.CompositionReconcileTimeout)

	// Verify the composition progresses after being unblocked
	s.VerifyOutputExists()
}

func TestSimplePlanStatusErrorExpansionFailed(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify there is a Error failure status condition in plan
	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition := utils.GetErrorCondition("EvaluateStatusFailed", "")
	s.C.MustHaveCondition(plan, condition, scenario.CompositionReconcileTimeout)

	// Apply Configmap with data present
	s.ApplyManifests("Composition without jinja error", "composition_fixed.yaml")

	// Check if Error failure condition is cleared
	plan = utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition = utils.GetErrorCondition("EvaluateStatusFailed", "")
	s.C.MustNotHaveCondition(plan, condition, 2*scenario.CompositionReconcileTimeout)

	// Verify the composition progresses after being unblocked
	s.VerifyOutputExists()
}

func TestSimpleExpanderJinjaWithQuotes(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify there is a Error failure status condition in plan
	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition := utils.GetErrorCondition("EvaluateStatusFailed", "")
	s.C.MustNotHaveCondition(plan, condition, 2*scenario.CompositionReconcileTimeout)

	// Verify the composition progresses after being unblocked
	s.VerifyOutputExists()
}

func TestSimplePlanStatusErrorFailedLoadingManifestsFromPlan(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify there is a Waiting failure status condition in plan
	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition := utils.GetErrorCondition("FailedLoadingManifestsFromPlan", "")
	s.C.MustHaveCondition(plan, condition, scenario.CompositionReconcileTimeout)

	// Apply Configmap with data present
	s.ApplyManifests("Composition without yaml error", "composition_fixed.yaml")

	// Check if Waiting failure condition is cleared
	plan = utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition = utils.GetErrorCondition("FailedLoadingManifestsFromPlan", "")
	s.C.MustNotHaveCondition(plan, condition, 2*scenario.CompositionReconcileTimeout)

	// Verify the composition progresses after being unblocked
	s.VerifyOutputExists()
}

func TestSimplePlanStatusErrorFailedApplyingManifests(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify there is a failed apply failure status condition in plan
	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition := utils.GetErrorCondition("FailedApplyingManifests", "")
	s.C.MustHaveCondition(plan, condition, scenario.CompositionReconcileTimeout)

	// Apply Configmap with data present
	s.ApplyManifests("Composition with correct object", "composition_fixed.yaml")

	// Check if failed apply failure condition is cleared
	plan = utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition = utils.GetErrorCondition("FailedApplyingManifests", "")
	s.C.MustNotHaveCondition(plan, condition, 2*scenario.CompositionReconcileTimeout)

	// Verify the composition progresses after being unblocked
	s.VerifyOutputExists()
}

func TestSimpleNamespaceInherit(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	s.VerifyOutputExists()
	s.VerifyOutputSpecMatches()

	// Verify presence of an error because one of the manifests is cluster scoped
	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition := utils.GetErrorCondition("FailedApplyingManifests", "")
	s.C.MustHaveCondition(plan, condition, scenario.CompositionReconcileTimeout)
}

func TestSimpleNamespaceExplicit(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	s.VerifyOutputExists()
	s.VerifyOutputSpecMatches()

	// Verify absence of an error because explicit allows cluster scoped
	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition := utils.GetErrorCondition("FailedApplyingManifests", "")
	s.C.MustNotHaveCondition(plan, condition, 2*scenario.CompositionReconcileTimeout)
}

// Test Bring Your OWN Schema
func TestSimpleFacadeByoSchema(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// TODO Add Ready condition which we should wait for
	time.Sleep(time.Second * 10)

	facade := utils.GetFacadeObj("default", "projectconfigmap")
	// Ensure no Error condition in Composition
	condition := utils.GetErrorCondition("", "")
	s.C.MustNotHaveCondition(facade, condition, 2*scenario.CompositionReconcileTimeout)

	// Verify CRD has been created and defn matches
	s.VerifyManifests("facade pconfig crd", true, "out_crd_pconfigs.yaml")

	// Create a facade from the new CRD
	s.ApplyManifests("facade cr", "in_pconfig.yaml")

	// Ensure no Error condition in Plan
	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition = utils.GetErrorCondition("", "")
	s.C.MustNotHaveCondition(plan, condition, 2*scenario.CompositionReconcileTimeout)

	// Verify the composition progresses after being unblocked
	s.VerifyOutputExists()
}

func TestSimpleExpanderInvalid(t *testing.T) {
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	composition := utils.GetCompositionObj("default", "projectconfigmap")
	condition := utils.GetErrorCondition("ValidationFailed", "")
	s.C.MustHaveCondition(composition, condition, scenario.CompositionReconcileTimeout)

	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition = utils.GetErrorCondition("MissingExpanderCR", "")
	s.C.MustHaveCondition(plan, condition, scenario.CompositionReconcileTimeout)
}

func TestSimpleExpanderVersionInvalid(t *testing.T) {
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	composition := utils.GetCompositionObj("default", "projectconfigmap")
	condition := utils.GetErrorCondition("ValidationFailed", "")
	s.C.MustHaveCondition(composition, condition, scenario.CompositionReconcileTimeout)

	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition = utils.GetErrorCondition("VersionNotFound", "")
	s.C.MustHaveCondition(plan, condition, scenario.CompositionReconcileTimeout)
}

func TestSimpleCompositionExpanderLoggingEnabled(t *testing.T) {
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition := utils.GetWaitingCondition("EvaluateStatusWait", "")
	s.C.MustNotHaveCondition(plan, condition, 3*scenario.CompositionReconcileTimeout)
	s.VerifyOutputExists()
}

func TestUpdateCompositionCRAddStage(t *testing.T) {
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Make sure the pre-update CRs have been created to ensure the coming update triggers a new reconcile.
	cm := utils.GetConfigMapObj("team-a", "common-config")
	s.C.MustExist([]*unstructured.Unstructured{cm}, scenario.ExistTimeout)

	// Apply the modified Composition
	s.ApplyManifests("updated composition", "updated_composition.yaml")

	// Changing the composition should trigger the expander to re-reconcile all objects.
	s.VerifyOutputExists()
	s.VerifyOutputSpecMatches()
}

func TestUpdateCompositionCRRemoveStage(t *testing.T) {
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify that all resources objects are created
	s.VerifyOutputExists()
	s.VerifyOutputSpecMatches()

	// Apply the modified Composition
	s.ApplyManifests("updated composition", "updated_composition.yaml")

	// Changing the composition should trigger the expander to re-reconcile all objects.
	cm := utils.GetConfigMapObj("team-a", "proj-a")
	s.C.MustNotExist([]*unstructured.Unstructured{cm}, scenario.ExistTimeout)
}

func TestUpdateCompositionModifyStage(t *testing.T) {
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify that all resources objects are created
	s.VerifyOutputExists()
	s.VerifyOutputSpecMatches()

	// Apply the modified Composition
	s.ApplyManifests("updated composition", "updated_composition.yaml")

	// Changing the composition should trigger the expander to re-reconcile all objects.
	cm := utils.GetConfigMapObj("team-a", "common-config-2")
	s.C.MustExist([]*unstructured.Unstructured{cm}, scenario.ExistTimeout)
}

func TestCustomStatusCELRule(t *testing.T) {
	s := scenario.NewBasic(t)
	//defer s.Cleanup()
	s.Setup()

	// Check plan object has no error
	plan := utils.GetPlanObj("my-team", "teampages-landing")
	condition := utils.GetReadyCondition("ProcessedAllStages", "")
	s.C.MustHaveCondition(plan, condition, 5*scenario.CompositionReconcileTimeout)

	// Verify that all resources objects are created
	s.VerifyOutputExists()
	s.VerifyOutputSpecMatches()

	// Verify plan has no errors
}
