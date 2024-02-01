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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ValidateOrDefaultStateIntoSpecAnnotation(obj *unstructured.Unstructured) error {
	_, found := GetAnnotation(StateIntoSpecAnnotation, obj)
	if !found {
		SetAnnotation(StateIntoSpecAnnotation, StateMergeIntoSpec, obj)
	}
	return validateStateIntoSpecAnnotation(obj)
}

func EnsureSpecIntoSateAnnotation(obj *Resource) error {
	_, found := GetAnnotation(StateIntoSpecAnnotation, obj)
	if !found {
		SetAnnotation(StateIntoSpecAnnotation, StateMergeIntoSpec, obj)
	}
	return validateStateIntoSpecAnnotation(obj)
}

func validateStateIntoSpecAnnotation(obj metav1.Object) error {
	val, found := GetAnnotation(StateIntoSpecAnnotation, obj)
	if !found {
		return fmt.Errorf("couldn't find the value for '%v' annotation", StateIntoSpecAnnotation)
	}

	if !isAcceptedValue(val) {
		return fmt.Errorf("invalid value '%v' for '%v' annotation, can be one of {%v}", val, StateIntoSpecAnnotation, strings.Join(StateIntoSpecAnnotationValues, ", "))
	}
	return nil
}

func isAcceptedValue(val string) bool {
	for _, v := range StateIntoSpecAnnotationValues {
		if val == v {
			return true
		}
	}
	return false
}
