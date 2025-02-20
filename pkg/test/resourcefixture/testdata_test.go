// Copyright 2025 Google LLC
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

package resourcefixture_test

import (
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
)

var (
	mockUniqueID = "unit-test"

	// exemptedTestNames map stores all the test cases that don't have their
	// dependent Service resources abandoned.
	// TODO: Abandon Service resources defined in the test cases below.
	exemptedTestNames = map[string]bool{
		// Test case "service" is testing CRUD of Service resource so it should
		// always be exempted.
		"service":                        true,
		"namespacepolicytoclusterpolicy": true,
		"serviceidentitypolicytoserviceaccountpolicy":               true,
		"privateserviceconnectioncomputeregionnetworkendpointgroup": true,
		"acmfeature":                          true,
		"mcifeature":                          true,
		"mcsdfeature":                         true,
		"iapidentityawareproxyclient":         true,
		"networksecurityclienttlspolicy":      true,
		"subnetconnector":                     true,
		"oauth2clientidcomputebackendservice": true,
	}
)

func TestServiceResources(t *testing.T) {
	testCases := resourcefixture.Load(t)
	for _, tc := range testCases {
		if _, ok := exemptedTestNames[tc.Name]; ok {
			continue
		}
		if tc.GVK.Kind == "Service" {
			createUnstruct := test.ToUnstructWithNamespace(t, tc.Create, mockUniqueID)
			assertServiceAbandoned(t, tc.Name, createUnstruct)
		}
		if tc.Dependencies != nil {
			dependencyYamls := testyaml.SplitYAML(t, tc.Dependencies)
			for _, dependBytes := range dependencyYamls {
				depUnstruct := test.ToUnstructWithNamespace(t, dependBytes, mockUniqueID)
				if depUnstruct.GetKind() == "Service" {
					assertServiceAbandoned(t, tc.Name, depUnstruct)
					assertServiceNotUnderMainTestProject(t, tc.Name, depUnstruct)
				}
			}
		}
	}
}

func assertServiceAbandoned(t *testing.T, testName string, service *unstructured.Unstructured) {
	annotations := service.GetAnnotations()
	value, ok := annotations[k8s.DeletionPolicyAnnotation]
	if !ok {
		t.Errorf("Service resource %q in test %q has no deletion policy annotation, "+
			"but it should have the following annotation: \"%s: abandon\"",
			strings.ReplaceAll(service.GetName(), mockUniqueID, "${uniqueId}"),
			testName, k8s.DeletionPolicyAnnotation)
		return
	}
	if value != k8s.DeletionPolicyAbandon {
		t.Errorf("Service resource %q in test %q should have %q annotation set to %q, "+
			"but it should be \"abandon\"",
			strings.ReplaceAll(service.GetName(), mockUniqueID, "${uniqueId}"),
			testName, k8s.DeletionPolicyAnnotation, value)
	}
}

func assertServiceNotUnderMainTestProject(t *testing.T, testName string, service *unstructured.Unstructured) {
	var projectID string
	var projectRefName string
	annotations := service.GetAnnotations()
	if value, ok := annotations[k8s.ProjectIDAnnotation]; ok {
		projectID = value
	}
	if value, exists, _ := unstructured.NestedString(service.Object, "spec", "projectRef", "name"); exists {
		projectRefName = value
	}
	if value, exists, _ := unstructured.NestedString(service.Object, "spec", "projectRef", "external"); exists {
		projectID = value
	}
	if projectRefName == "" && (projectID == "" || projectID == "${projectId}") {
		t.Errorf("Service resource %q in test %q should not be configured in dependencies.yaml, "+
			"but it should be enabled directly in shared-vars-public.sh",
			strings.ReplaceAll(service.GetName(), mockUniqueID, "${uniqueId}"),
			testName)
	}
}
