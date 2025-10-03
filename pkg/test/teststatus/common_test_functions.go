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

package teststatus

import (
	"testing"

	condition "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// GetStatus holds the required fields for computing if an object should be considered ready (fully reconciled).
type ObjectStatus struct {
	Generation         int64
	ObservedGeneration *int64
	Conditions         []condition.Condition
}

// GetObjectStatus extracts the required fields for computing if an object should be considered ready (fully reconciled).
func GetObjectStatus(t *testing.T, object runtime.Object) ObjectStatus {
	// Simple types with the fields we care about, so that we can use the libraries
	type withConditions struct {
		ObservedGeneration *int64                `json:"observedGeneration,omitempty"`
		Conditions         []condition.Condition `json:"conditions"`
	}
	type withStatusConditions struct {
		Status withConditions `json:"status"`
	}

	// Extract information by converting to withConditions

	u, ok := object.(*unstructured.Unstructured)
	if !ok {
		m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(object)
		if err != nil {
			t.Errorf("error from runtime.DefaultUnstructuredConverter.ToUnstructured(%T): %v", object, err)
		}
		u = &unstructured.Unstructured{Object: m}
	}

	generation := u.GetGeneration()
	var statusConditionsObj withStatusConditions
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &statusConditionsObj); err != nil {
		t.Errorf("error converting to object with status.conditions: %v", err)
	}

	return ObjectStatus{
		Generation:         generation,
		ObservedGeneration: statusConditionsObj.Status.ObservedGeneration,
		Conditions:         statusConditionsObj.Status.Conditions,
	}
}
