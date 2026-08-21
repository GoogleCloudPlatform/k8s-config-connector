// Copyright 2022 Google LLC
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

package testcontroller

import (
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func AssertServerGeneratedResourceIDMatch(t *testing.T, reconciledResourceID string, appliedUnstruct *unstructured.Unstructured) {
	t.Helper()

	appliedResourceID, foundInAppliedSpec := GetResourceID(t, appliedUnstruct)
	if foundInAppliedSpec && reconciledResourceID != appliedResourceID {
		t.Fatalf("resourceID should never be changed if specified; "+
			"resourceID before reconciliation: '%s', resourceID after "+
			"reconciliation: '%s'", appliedResourceID, reconciledResourceID)
	}
}

func AssertUserSpecifiedResourceIDMatch(t *testing.T, reconciledResourceID string, appliedUnstruct *unstructured.Unstructured) {
	t.Helper()

	appliedResourceID, foundInAppliedSpec := GetResourceID(t, appliedUnstruct)

	if !foundInAppliedSpec && reconciledResourceID != appliedUnstruct.GetName() {
		t.Fatalf("resourceID must be the same as the value of "+
			"'metadata.name' if '%s' is not specified; 'metadata.name': "+
			"'%s', resourceID after reconciliation: '%s'",
			k8s.ResourceIDFieldPath, appliedUnstruct.GetName(),
			reconciledResourceID)
	}

	if foundInAppliedSpec && reconciledResourceID != appliedResourceID {
		t.Fatalf("resourceID should never be changed if specified; "+
			"resourceID before reconciliation: '%s', resourceID after "+
			"reconciliation: '%s'", appliedResourceID, reconciledResourceID)
	}
}

// TODO(kcc-eng): Clean up all the duplicate SupportsResourceIDField functions.
func SupportsResourceIDField(rc *v1alpha1.ResourceConfig) bool {
	return rc.ResourceID.TargetField != ""
}

// TODO(kcc-eng): Clean up all the duplicate IsResourceIDFieldServerGenerated functions.
func IsResourceIDFieldServerGenerated(rc *v1alpha1.ResourceConfig) bool {
	return rc.ResourceID.TargetField == rc.ServerGeneratedIDField &&
		rc.ServerGeneratedIDField != ""
}

func GetResourceID(t *testing.T, resourceUnstruct *unstructured.Unstructured) (string, bool) {
	resourceID, found, err := unstructured.NestedString(resourceUnstruct.Object, "spec", k8s.ResourceIDFieldName)
	if err != nil {
		t.Fatalf("error getting '%s': %v", k8s.ResourceIDFieldPath, err)
	}

	return resourceID, found
}

func SetResourceID(t *testing.T, u *unstructured.Unstructured, val string) {
	if err := unstructured.SetNestedField(u.Object, val, strings.Split(k8s.ResourceIDFieldPath, ".")...); err != nil {
		t.Fatalf("error setting '%s' on unstruct: %v", k8s.ResourceIDFieldPath, err)
	}
}

func RemoveResourceID(u *unstructured.Unstructured) {
	unstructured.RemoveNestedField(u.Object, strings.Split(k8s.ResourceIDFieldPath, ".")...)
}
