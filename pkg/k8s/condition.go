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

package k8s

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewCustomReadyCondition(status v1.ConditionStatus, reason, message string) v1alpha1.Condition {
	return v1alpha1.Condition{
		LastTransitionTime: metav1.Now().Format(time.RFC3339),
		Type:               v1alpha1.ReadyConditionType,
		Status:             status,
		Reason:             reason,
		Message:            message,
	}
}

func NewReadyCondition() v1alpha1.Condition {
	return v1alpha1.Condition{
		LastTransitionTime: metav1.Now().Format(time.RFC3339),
		Type:               v1alpha1.ReadyConditionType,
	}
}

func NewReadyConditionWithError(err error) v1alpha1.Condition {
	var readyCondition v1alpha1.Condition
	errWithReason := &ErrorWithReason{}
	if errors.As(err, errWithReason) {
		readyCondition = NewCustomReadyCondition(v1.ConditionFalse, errWithReason.Reason, errWithReason.Message)
	} else {
		readyCondition = NewReadyCondition()
		readyCondition.Status = v1.ConditionFalse
		readyCondition.Message = err.Error()
	}
	return readyCondition
}

func ConditionsEqualIgnoreTransitionTime(c1, c2 v1alpha1.Condition) bool {
	return c1.Message == c2.Message &&
		c1.Reason == c2.Reason &&
		c1.Status == c2.Status &&
		c1.Type == c2.Type
}

func ConditionSlicesEqual(conditions1, conditions2 []v1alpha1.Condition) bool {
	if len(conditions1) != len(conditions2) {
		return false
	}
	for i, c1 := range conditions1 {
		c2 := conditions2[i]
		if !ConditionsEqualIgnoreTransitionTime(c1, c2) {
			return false
		}
	}
	return true
}

func MarshalAsConditionsSlice(obj []interface{}) ([]v1alpha1.Condition, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	ret := make([]v1alpha1.Condition, 0)
	if err := json.Unmarshal(b, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}
