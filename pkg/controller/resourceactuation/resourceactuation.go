// Copyright 2023 Google LLC
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

package resourceactuation

import (
	"fmt"

	opv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	opk8s "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/reconciliationinterval"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// DecideActuationMode looks at CC and CCC to see if they specify an actuationMode.
// The order of precedence is:
//  1. Check CCC (if in namespaced mode); if unset
//  2. Check CC; if unset
//  3. Use default value "Reconciling"
func DecideActuationMode(cc opv1beta1.ConfigConnector, ccc opv1beta1.ConfigConnectorContext) opv1beta1.ActuationMode {
	// 1. Check CCC (only relevant in Namespaced mode).
	if ccc.Spec.Actuation != "" && cc.Spec.Mode == opk8s.NamespacedMode {
		return ccc.Spec.Actuation
	}

	// 2. Check CC.
	if cc.Spec.Actuation != "" {
		return cc.Spec.Actuation
	}

	// 3. Default.
	return opv1beta1.DefaultActuationMode()
}

// ShouldSkipActuation determines whether actuation against GCP should be skipped
// based on resource annotations and CC/CCC actuation modes.
//
// Rules:
// 1. Resource annotation takes precedence:
//   - "Paused": Skips Create/Update/Drift, but allows Deletion to proceed (returns false if isDeleting).
//   - "Reconciling": Allows actuation to proceed (returns false).
//   - Invalid value: Returns an error.
//
// 2. Fall back to CC/CCC actuation mode (via DecideActuationMode):
//   - "Paused": Skips all operations including deletion (returns true).
//   - "Reconciling": Allows actuation to proceed (returns false).
func ShouldSkipActuation(annotations map[string]string, isDeleting bool, cc opv1beta1.ConfigConnector, ccc opv1beta1.ConfigConnectorContext) (bool, error) {
	if annotations != nil {
		if val, ok := annotations[k8s.ActuationModeAnnotation]; ok && val != "" {
			switch opv1beta1.ActuationMode(val) {
			case opv1beta1.Reconciling:
				return false, nil
			case opv1beta1.Paused:
				// Resource-level pause does NOT pause deletion
				if isDeleting {
					return false, nil
				}
				return true, nil
			default:
				return false, fmt.Errorf("invalid value %q for annotation %s; allowed values are %q and %q",
					val, k8s.ActuationModeAnnotation, opv1beta1.Paused, opv1beta1.Reconciling)
			}
		}
	}

	am := DecideActuationMode(cc, ccc)
	switch am {
	case opv1beta1.Reconciling:
		return false, nil
	case opv1beta1.Paused:
		return true, nil
	default:
		return false, fmt.Errorf("unknown actuation mode %v", am)
	}
}

// ShouldSkip skips a resource actuatation if the ReconcileIntervalInSecondsAnnotation = 0 and the KRM resource has not changed since its last UpToDate.
// This will disable drift correction on corresponding GCP resources since the reconcileInterval is set to 0.
func ShouldSkip(u *unstructured.Unstructured) (bool, error) {
	generation, found, err := unstructured.NestedInt64(u.Object, "metadata", "generation")
	if err != nil {
		return false, fmt.Errorf("error getting the value for 'metadata.generation' %w", err)
	}
	if !found {
		return false, nil
	}
	observedGeneration, found, err := unstructured.NestedInt64(u.Object, "status", "observedGeneration")
	if err != nil {
		return false, fmt.Errorf("error getting the value for 'status.observedGeneration': %w", err)
	}
	if !found {
		return false, nil
	}
	if observedGeneration != generation {
		return false, nil
	}

	if val, ok := k8s.GetAnnotation(k8s.ReconcileIntervalInSecondsAnnotation, u); ok {
		reconcileInterval, err := reconciliationinterval.MeanReconcileReenqueuePeriodFromAnnotation(val)
		if err != nil {
			return false, err
		}
		if reconcileInterval == 0 {
			conditions, found, err := unstructured.NestedSlice(u.Object, "status", "conditions")
			if err != nil {
				return false, fmt.Errorf("error getting object conditions: %w", err)
			}
			if !found {
				return false, nil
			}
			for _, condition := range conditions {
				conditionMap, ok := condition.(map[string]interface{})
				if !ok {
					return false, fmt.Errorf("error converting condition %v to map", condition)
				}
				if status, foundStatus := conditionMap["status"].(string); foundStatus && status == "True" {
					if reason, foundCondition := conditionMap["reason"].(string); foundCondition && reason == k8s.UpToDate {
						return true, nil
					}
				}
			}
		}
	}
	return false, nil
}
