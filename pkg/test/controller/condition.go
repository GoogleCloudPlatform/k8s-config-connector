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

	condition "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
)

// AssertReadyCondition checks that the given conditions slice contains a Ready condition.
func AssertReadyCondition(t *testing.T, conditions []condition.Condition) {
	t.Helper()
	if len(conditions) != 1 {
		t.Fatalf("expected 1 condition, instead have %v", len(conditions))
	}
	readyCondition := conditions[0]
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
