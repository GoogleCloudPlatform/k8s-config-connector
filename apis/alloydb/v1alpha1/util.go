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

package v1alpha1

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func isResourceReady(u *unstructured.Unstructured) (bool, error) {
	cond, found, err := getReadyCondition(u)
	if err != nil {
		return false, err
	}
	return found && cond.Status == corev1.ConditionTrue, nil
}

func getReadyCondition(u *unstructured.Unstructured) (condition v1alpha1.Condition, found bool, err error) {
	conditionsRaw, _, err := unstructured.NestedSlice(u.Object, "status", "conditions")
	if err != nil {
		return v1alpha1.Condition{}, false, fmt.Errorf("error reading status.conditions: %w", err)
	}
	if conditionsRaw == nil {
		return v1alpha1.Condition{}, false, nil
	}
	if conditions, err := k8s.MarshalAsConditionsSlice(conditionsRaw); err == nil {
		for _, condition := range conditions {
			if condition.Type == k8sv1alpha1.ReadyConditionType {
				return condition, true, nil
			}
		}
	}

	return v1alpha1.Condition{}, false, nil
}
