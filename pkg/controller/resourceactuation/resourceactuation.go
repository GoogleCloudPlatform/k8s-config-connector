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

// DecideActuationMode looks at annotations, CC, and CCC to see if they specify an actuationMode.
//   - If CC or CCC specifies "Paused" at the namespace/cluster level, that config takes precedence
//     and the resource actuation mode is always "Paused" (individual resources cannot override to "Reconciling").
//   - If CC/CCC level is "Reconciling" (or unset), and the individual resource has the "cnrm.cloud.google.com/actuation-mode: Paused"
//     annotation, then the individual resource actuation mode is "Paused".
//   - Otherwise, the resource actuation mode is "Reconciling".
func DecideActuationMode(annotations map[string]string, cc opv1beta1.ConfigConnector, ccc opv1beta1.ConfigConnectorContext) opv1beta1.ActuationMode {
	// First, determine the base actuation mode from CC/CCC config.
	var baseActuationMode opv1beta1.ActuationMode
	if ccc.Spec.Actuation != "" && cc.Spec.Mode == opk8s.NamespacedMode {
		baseActuationMode = ccc.Spec.Actuation
	} else if cc.Spec.Actuation != "" {
		baseActuationMode = cc.Spec.Actuation
	} else {
		baseActuationMode = opv1beta1.DefaultActuationMode()
	}

	// If the namespace or cluster level actuation is paused, the resource is always paused.
	if baseActuationMode == opv1beta1.Paused {
		return opv1beta1.Paused
	}

	// Otherwise, the base mode is Reconciling. Check if the individual resource annotation pauses it.
	if annotations != nil {
		if val, ok := annotations[k8s.ActuationModeAnnotation]; ok {
			if opv1beta1.ActuationMode(val) == opv1beta1.Paused {
				return opv1beta1.Paused
			}
		}
	}

	return opv1beta1.Reconciling
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
