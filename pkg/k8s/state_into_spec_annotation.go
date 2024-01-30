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
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewStateIntoSpecDefaulter(defaultValue string, userOverride *string) (Defaulter, error) {
	if !isAcceptedValue(defaultValue, StateIntoSpecAnnotationValues) {
		return nil, fmt.Errorf("invalid default value '%v' for '%v' annotation, need to be one of {%v}", defaultValue, StateIntoSpecAnnotation, strings.Join(StateIntoSpecAnnotationValues, ", "))
	}
	if userOverride != nil && !isAcceptedValue(*userOverride, StateIntoSpecAnnotationValues) {
		return nil, fmt.Errorf("invalid user override value '%v' for '%v' annotation, need to be one of {%v}", userOverride, StateIntoSpecAnnotation, strings.Join(StateIntoSpecAnnotationValues, ", "))
	}
	return &StateIntoSpecDefaulter{
		defaultValue: defaultValue,
		userOverride: userOverride,
	}, nil
}

// ValidateOrDefaultStateIntoSpecAnnotation validates the value of the
// 'state-into-spec' annotation if it is set and defaults the annotation to the
// passed in defaultValue if it is unset.
func ValidateOrDefaultStateIntoSpecAnnotation(obj metav1.Object, defaultValue string) error {
	_, found := GetAnnotation(StateIntoSpecAnnotation, obj)
	if !found {
		SetAnnotation(StateIntoSpecAnnotation, defaultValue, obj)
	}
	return validateStateIntoSpecAnnotation(obj)
}

func validateStateIntoSpecAnnotation(obj metav1.Object) error {
	val, found := GetAnnotation(StateIntoSpecAnnotation, obj)
	if !found {
		return fmt.Errorf("couldn't find the value for '%v' annotation", StateIntoSpecAnnotation)
	}

	if !isAcceptedValue(val, StateIntoSpecAnnotationValues) {
		return fmt.Errorf("invalid value '%v' for '%v' annotation, can be one of {%v}", val, StateIntoSpecAnnotation, strings.Join(StateIntoSpecAnnotationValues, ", "))
	}
	return nil
}

func isAcceptedValue(val string, acceptedValues []string) bool {
	for _, v := range acceptedValues {
		if val == v {
			return true
		}
	}
	return false
}
