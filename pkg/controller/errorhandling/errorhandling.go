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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	// ErrorHandlingModeAnnotation is the annotation key for configuring custom error handling modes.
	ErrorHandlingModeAnnotation = "cnrm.cloud.google.com/error-handling-mode"

	// ErrorHandlingModeContinuousRetry is a mode where reconciliation retries continuously despite terminal client errors.
	ErrorHandlingModeContinuousRetry = "ContinuousRetry"
)

// ShouldSkip returns true if the resource reconciliation should be skipped.
// Reconciliation is skipped if the resource has a Ready condition with reason "UpdateFailedTerminalError",
// the status.observedGeneration matches metadata.generation, and the error-handling-mode
// annotation is not set to "ContinuousRetry".
func ShouldSkip(u *unstructured.Unstructured) bool {
	if u.GetDeletionTimestamp() != nil {
		return false
	}

	conditions, foundCond, errCond := unstructured.NestedSlice(u.Object, "status", "conditions")
	if errCond != nil || !foundCond {
		return false
	}

	for _, condRaw := range conditions {
		cond, ok := condRaw.(map[string]interface{})
		if !ok {
			continue
		}
		if cond["type"] == "Ready" && cond["reason"] == "UpdateFailedTerminalError" {
			observedGen, foundGen, errGen := unstructured.NestedInt64(u.Object, "status", "observedGeneration")
			if errGen == nil && foundGen && u.GetGeneration() == observedGen {
				annotations := u.GetAnnotations()
				if annotations == nil || annotations[ErrorHandlingModeAnnotation] != ErrorHandlingModeContinuousRetry {
					return true
				}
			}
		}
	}

	return false
}
