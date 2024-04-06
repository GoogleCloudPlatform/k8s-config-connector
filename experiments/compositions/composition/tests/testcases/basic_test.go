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
	s := scenario.New(t, "")
	defer s.ReleaseResources()
	s.ApplyInput()
	s.VerifyInput()
	s.CleanupInput()
	s.CleanupOutput()
	s.GatherLogs()
}

func TestSimpleCompositionExpansion(t *testing.T) {
	s := scenario.New(t, "")
	defer s.ReleaseResources()
	s.ApplyInput()
	s.VerifyInput()
	s.VerifyOutputExists()
	s.VerifyOutputSpecMatches()
	s.CleanupInput()
	s.CleanupOutput()
	s.GatherLogs()
}

func TestSimpleCompositionDeleteFacade(t *testing.T) {
	s := scenario.New(t, "")
	defer s.ReleaseResources()
	s.ApplyInput()
	s.VerifyInput()
	s.VerifyOutputExists()

	// Delete CR
	cr := utils.GetUnstructuredObj("facade.foocorp.com", "v1", "PConfig", "team-a", "team-a-config")
	s.C.MustDelete(cr)

	// Check if Plan is deleted
	plan := utils.GetPlanObj("team-a", "pconfigs-team-a-config")
	s.C.MustNotExist([]*unstructured.Unstructured{plan}, scenario.DeleteTimeout)

	// Check if expanded ConfigMap is also deleted
	cm := utils.GetConfigMapObj("team-a", "proj-a")
	s.C.MustNotExist([]*unstructured.Unstructured{cm}, scenario.DeleteTimeout)

	s.CleanupInput()
	s.CleanupOutput()
	s.GatherLogs()
}
