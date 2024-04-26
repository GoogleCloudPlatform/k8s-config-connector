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

	"google.com/composition/tests/scenario"
	"google.com/composition/tests/utils"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestSimpleCompositionCreate(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()
}

func TestSimpleCompositionExpansion(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	s.VerifyOutputExists()
	s.VerifyOutputSpecMatches()
}

func TestSimpleCompositionDeleteFacade(t *testing.T) {
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
func TestSimpleCompositionAddFacadeField(t *testing.T) {
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
	s.C.MustExist([]*unstructured.Unstructured{cm}, scenario.ExistTimeout)
}

// Test removing config that results in removal of some expanded resource
func TestSimpleCompositionDeleteFacadeField(t *testing.T) {
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
	s.C.MustHaveCondition(composition, condition, scenario.ExistTimeout)

	// Apply the fixed Composition
	s.ApplyManifests("composition without validation error", "fixed_composition.yaml")

	// Check if Validation failure condition is cleared
	composition = utils.GetCompositionObj("default", "projectconfigmap")
	condition = utils.GetValidationFailedCondition("ExpanderValidationFailed", "")
	s.C.MustNotHaveCondition(composition, condition, scenario.ExistTimeout)
}

func TestSimpleCompositionStatusFacadeCRDMissing(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify there is a Error status condition in compositions
	composition := utils.GetCompositionObj("default", "projectconfigmap")
	condition := utils.GetErrorCondition("MissingFacadeCRD", "")
	s.C.MustHaveCondition(composition, condition, scenario.ExistTimeout)

	// Apply the fixed Composition
	s.ApplyManifests("Facade CRD", "facade_crd.yaml")

	// Check if Validation failure condition is cleared
	composition = utils.GetCompositionObj("default", "projectconfigmap")
	condition = utils.GetErrorCondition("MissingFacadeCRD", "")
	s.C.MustNotHaveCondition(composition, condition, scenario.ExistTimeout)
}

func TestSimplePlanStatusWaitingForValues(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify there is a Waiting failure status condition in plan
	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition := utils.GetWaitingCondition("FetchValuesFromFailed", "")
	s.C.MustHaveCondition(plan, condition, scenario.ExistTimeout)

	// Apply Configmap with data present
	s.ApplyManifests("Fixed Configmap", "configmap_with_data.yaml")

	// Check if Waiting failure condition is cleared
	plan = utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition = utils.GetWaitingCondition("FetchValuesFromFailed", "")
	s.C.MustNotHaveCondition(plan, condition, 3*scenario.CompositionReconcile)

	// Verify the composition progresses after being unblocked
	s.VerifyOutputExists()
}

func TestSimplePlanStatusErrorExpansionFailed(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify there is a Waiting failure status condition in plan
	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition := utils.GetErrorCondition("ExpansionFailed", "")
	s.C.MustHaveCondition(plan, condition, scenario.ExistTimeout)

	// Apply Configmap with data present
	s.ApplyManifests("Composition without jinja error", "composition_fixed.yaml")

	// Check if Waiting failure condition is cleared
	plan = utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition = utils.GetErrorCondition("ExpansionFailed", "")
	s.C.MustNotHaveCondition(plan, condition, 2*scenario.CompositionReconcile)

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
	s.C.MustHaveCondition(plan, condition, scenario.ExistTimeout)

	// Apply Configmap with data present
	s.ApplyManifests("Composition without yaml error", "composition_fixed.yaml")

	// Check if Waiting failure condition is cleared
	plan = utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition = utils.GetErrorCondition("FailedLoadingManifestsFromPlan", "")
	s.C.MustNotHaveCondition(plan, condition, 2*scenario.CompositionReconcile)

	// Verify the composition progresses after being unblocked
	s.VerifyOutputExists()
}

func TestSimplePlanStatusErrorFailedApplyingManifests(t *testing.T) {
	//t.Parallel()
	s := scenario.NewBasic(t)
	defer s.Cleanup()
	s.Setup()

	// Verify there is a Waiting failure status condition in plan
	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition := utils.GetErrorCondition("FailedApplyingManifests", "")
	s.C.MustHaveCondition(plan, condition, scenario.ExistTimeout)

	// Apply Configmap with data present
	s.ApplyManifests("Composition with correct object", "composition_fixed.yaml")

	// Check if Waiting failure condition is cleared
	plan = utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	condition = utils.GetErrorCondition("FailedApplyingManifests", "")
	s.C.MustNotHaveCondition(plan, condition, 2*scenario.CompositionReconcile)

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
	s.C.MustHaveCondition(plan, condition, scenario.ExistTimeout)
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
	s.C.MustNotHaveCondition(plan, condition, 2*scenario.CompositionReconcile)
}
