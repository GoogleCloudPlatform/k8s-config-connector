// Copyright 2026 Google LLC
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

package errorhandling

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// ShouldSkip returns true if reconciliation should be skipped due to an unchanged terminal error state.
func ShouldSkip(u *unstructured.Unstructured) (bool, error) {
	if u.GetDeletionTimestamp() != nil {
		return false, nil
	}

	if val, ok := k8s.GetAnnotation(k8s.ErrorHandlingModeAnnotation, u); ok && val == k8s.ErrorHandlingModeContinuousRetry {
		return false, nil
	}

	generation, foundGen, err := unstructured.NestedInt64(u.Object, "metadata", "generation")
	if err != nil {
		return false, fmt.Errorf("error reading 'metadata.generation': %w", err)
	}
	observedGeneration, foundObsGen, err := unstructured.NestedInt64(u.Object, "status", "observedGeneration")
	if err != nil {
		return false, fmt.Errorf("error reading 'status.observedGeneration': %w", err)
	}
	if !foundGen || !foundObsGen || generation != observedGeneration {
		return false, nil
	}

	conditions, foundCond, err := unstructured.NestedSlice(u.Object, "status", "conditions")
	if err != nil {
		return false, fmt.Errorf("error reading 'status.conditions': %w", err)
	}
	if !foundCond {
		return false, nil
	}

	for _, condRaw := range conditions {
		cond, ok := condRaw.(map[string]interface{})
		if !ok {
			continue
		}
		if cond["type"] == string(k8s.ConditionReady) &&
			cond["status"] == string(corev1.ConditionFalse) &&
			cond["reason"] == k8s.UpdateFailedTerminalError {
			return true, nil
		}
	}

	return false, nil
}
