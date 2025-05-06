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

package managementconflict

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestGetManagementConflictPreventionAnnotationValue(t *testing.T) {
	testCases := []struct {
		Name           string
		Annotations    map[string]string
		ExpectedPolicy ManagementConflictPreventionPolicy
		ShouldSucceed  bool
	}{
		{
			Name:           "nil annotations should no longer error",
			Annotations:    nil,
			ExpectedPolicy: ManagementConflictPreventionPolicyNone,
			ShouldSucceed:  true,
		},
		{
			Name:           "missing annotation should no longer error",
			Annotations:    make(map[string]string),
			ExpectedPolicy: ManagementConflictPreventionPolicyNone,
			ShouldSucceed:  true,
		},
		{
			Name:           "invalid annotation should error",
			Annotations:    newManagementConflictAnnotations("my invalid policy name"),
			ExpectedPolicy: ManagementConflictPreventionPolicyNone,
			ShouldSucceed:  false,
		},
		{
			Name:           "valid value should succeed",
			Annotations:    newManagementConflictAnnotations(ManagementConflictPreventionPolicyResource),
			ExpectedPolicy: ManagementConflictPreventionPolicyResource,
			ShouldSucceed:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			obj := unstructured.Unstructured{}
			obj.SetAnnotations(tc.Annotations)
			policy, err := GetManagementConflictPreventionPolicy(&obj)
			if tc.ShouldSucceed != (err == nil) {
				t.Fatalf("expected success to be '%v', instead got error mismatch: %v", tc.ShouldSucceed, err)
			}
			if policy != tc.ExpectedPolicy {
				t.Fatalf("policy mismatch: got '%v', want '%v'", policy, tc.ExpectedPolicy)
			}
		})
	}
}

func newManagementConflictAnnotations(policy string) map[string]string {
	annotations := make(map[string]string)
	if policy != "" {
		annotations[FullyQualifiedAnnotation] = policy
	}
	return annotations
}
