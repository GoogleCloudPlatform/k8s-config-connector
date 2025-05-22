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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dynamic"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
)

// AssertReadyCondition checks that the given status.conditions slice contains a Ready condition.
// It waits for the specified observedGeneration (or later).
func AssertReadyCondition(t *testing.T, object runtime.Object, minObservedGeneration int64) {
	t.Helper()

	gvk := object.GetObjectKind().GroupVersionKind()

	accessor, err := meta.Accessor(object)
	if err != nil {
		t.Fatalf("unable to get accessor for object of type %T: %v", object, err)
	}
	objectID := gvk.Kind + ":" + accessor.GetName()

	objectStatus, err := dynamic.GetObjectStatus(object)
	if err != nil {
		t.Fatalf("error getting object status: %v", err)
	}
	if objectStatus.ObservedGeneration == nil {
		t.Fatalf("resource %v does not yet have status.observedGeneration", objectID)
	}
	if *objectStatus.ObservedGeneration < minObservedGeneration {
		t.Fatalf("resource %v status.observedGeneration %v is behind minObservedGeneration %v", objectID, *objectStatus.ObservedGeneration, minObservedGeneration)
	}

	if len(objectStatus.Conditions) != 1 {
		t.Fatalf("expected 1 condition, instead have %v", len(objectStatus.Conditions))
	}

	readyCondition := objectStatus.Conditions[0]
	if readyCondition.Type != "Ready" {
		t.Errorf("readyCondition type mismatch: got '%v', want '%v'", readyCondition.Type, "Ready")
	}
	if readyCondition.LastTransitionTime == "" {
		t.Errorf("readyCondition last transition time is empty string, expected value")
	}
	if readyCondition.Status != "True" {
		t.Errorf("status value mismatch: got '%v', want '%v'", readyCondition.Status, "True")
	}
}
