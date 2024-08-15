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

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/tests/scenario"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/tests/testclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/tests/utils"
	"k8s.io/apimachinery/pkg/util/rand"
)

// ------------------------ FirstComposition --------------------------------------------

func getNamespaceObj(ns string) []*unstructured.Unstructured {
	return []*unstructured.Unstructured{
		{
			Object: map[string]interface{}{
				"apiVersion": "v1",
				"kind":       "Namespace",
				"metadata": map[string]interface{}{
					"name": ns,
				},
			},
		},
	}
}

func getTeamPageObj(team string) []*unstructured.Unstructured {
	return []*unstructured.Unstructured{
		{
			Object: map[string]interface{}{
				"apiVersion": "idp.mycompany.com/v1alpha1",
				"kind":       "TeamPage",
				"metadata": map[string]interface{}{
					"name":      team,
					"namespace": team,
				},
				// members:
				"spec": map[string]interface{}{
					"members": []map[string]interface{}{
						{
							"name": "Jo",
							"role": "Eng Manager",
						},
						{
							"name": "Jane",
							"role": "Lead",
						},
						{
							"name": "Bob",
							"role": "Developer",
						},
					},
				},
			},
		},
	}
}

func getTeamPageOutputObjects(team string) []*unstructured.Unstructured {
	gkvnns := []testclient.GVKNN{
		{GroupVersionKind: schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"},
			NamespacedName: types.NamespacedName{Name: "team-" + team, Namespace: team}},
		{GroupVersionKind: schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Service"},
			NamespacedName: types.NamespacedName{Name: "team-" + team + "-landing", Namespace: team}},
		{GroupVersionKind: schema.GroupVersionKind{Group: "", Version: "v1", Kind: "ConfigMap"},
			NamespacedName: types.NamespacedName{Name: "team-" + team + "-page", Namespace: team}},
	}

	objs := []*unstructured.Unstructured{}
	for _, gvknn := range gkvnns {
		objs = append(objs, gvknn.MakeObject())
	}

	return objs
}

func TestFirstCompositionTeamPage(t *testing.T) {
	//t.Parallel()
	s := scenario.NewFromSample(t, scenario.Sample{Name: "FirstComposition", Composition: "teampage.yaml"}, nil, false)
	defer s.Cleanup()
	s.Setup()

	team := fmt.Sprintf("team-%s", strings.ToLower(rand.String(8)))
	// Create Namespace for team first
	s.Apply("namespace", getNamespaceObj(team))

	// Create team page facade
	s.Apply("teampage-object", getTeamPageObj(team))

	// Check plan object has no error
	plan := utils.GetPlanObj(team, "teampages-"+team)
	condition := utils.GetReadyCondition("ProcessedAllStages", "")
	s.C.MustHaveCondition(plan, condition, 5*scenario.CompositionReconcileTimeout)

	// Since the Plan says it has processed all stages we should validate applied resources.
	s.Verify("objects", false, getTeamPageOutputObjects(team))
}
