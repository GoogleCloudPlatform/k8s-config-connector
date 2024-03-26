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
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ValidateOrDefaultStateIntoSpecAnnotation validates the value of the
// 'state-into-spec' annotation if it is set and defaults the annotation to the
// passed in defaultValue if it is unset.
func ValidateOrDefaultStateIntoSpecAnnotation(obj client.Object, defaultValue string) error {
	_, found := GetAnnotation(StateIntoSpecAnnotation, obj)
	if !found {
		setStateIntoSpecDefaultValueIfAllowed(defaultValue, obj)
	}
	return validateStateIntoSpecAnnotation(obj)
}

func validateStateIntoSpecAnnotation(obj client.Object) error {
	val, found := GetAnnotation(StateIntoSpecAnnotation, obj)
	if !found {
		return fmt.Errorf("couldn't find the value for '%v' annotation", StateIntoSpecAnnotation)
	}

	if !isAcceptedStateIntoSpecValue(val, obj.GetObjectKind().GroupVersionKind()) {
		return fmt.Errorf("invalid value '%v' for '%v' annotation in kind %v", val, StateIntoSpecAnnotation, obj.GetObjectKind().GroupVersionKind().Kind)
	}
	return nil
}

func setStateIntoSpecDefaultValueIfAllowed(defaultValue string, obj client.Object) {
	if defaultValue == StateMergeIntoSpec && !supportsStateIntoSpecMerge(obj.GetObjectKind().GroupVersionKind()) {
		SetAnnotation(StateIntoSpecAnnotation, StateAbsentInSpec, obj)
		return

	}
	SetAnnotation(StateIntoSpecAnnotation, defaultValue, obj)
}

func isAcceptedStateIntoSpecValue(value string, gvk schema.GroupVersionKind) bool {
	if supportsStateIntoSpecMerge(gvk) {
		return isAcceptedValue(value, StateIntoSpecAnnotationValues)
	}
	return value == StateAbsentInSpec
}

func isAcceptedValue(val string, acceptedValues []string) bool {
	for _, v := range acceptedValues {
		if val == v {
			return true
		}
	}
	return false
}
