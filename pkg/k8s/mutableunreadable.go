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
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func GenerateMutableButUnreadableFieldsAnnotation(resource *Resource, mutableButUnreadablePaths [][]string) (string, error) {
	state, err := GenerateMutableButUnreadableFieldsState(resource, mutableButUnreadablePaths)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(state)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func GenerateMutableButUnreadableFieldsState(resource *Resource, mutableButUnreadablePaths [][]string) (map[string]interface{}, error) {
	state := make(map[string]interface{})
	for _, path := range mutableButUnreadablePaths {
		krmPath := append([]string{"spec"}, path...)
		krmField := strings.Join(krmPath, ".")
		val, found, err := unstructured.NestedFieldCopy(resource.Spec, krmPath[1:]...)
		if err != nil {
			return nil, fmt.Errorf("error reading %v: %w", krmField, err)
		}
		if !found {
			continue
		}
		if err := unstructured.SetNestedField(state, val, krmPath...); err != nil {
			return nil, fmt.Errorf("error saving %v: %w", krmField, err)
		}
	}
	return state, nil
}

func GetMutableButUnreadableFieldsFromAnnotations(resource *Resource, mutableButUnreadablePaths [][]string) (map[string]interface{}, error) {
	mutableButUnreadableFields := make(map[string]interface{})
	if len(mutableButUnreadablePaths) == 0 {
		// The resource does not have any mutable-but-unreadable field.
		return mutableButUnreadableFields, nil
	}

	var err error
	annotationVal, ok := GetAnnotation(MutableButUnreadableFieldsAnnotation, resource)
	if !ok {
		// If the resource can have mutable-but-unreadable fields but does not
		// have the annotation set at all, then this is either (1) a new resource,
		// (2) a resource acquisition, or (3) a resource created before it supported
		// mutable-but-unreadable fields.
		// To avoid unnecessarily updating the resource in Cases 2 and 3, generate
		// a value for the annotation based on its current spec.
		// Note that while this will also generate an annotation for Case 1
		// (ideally it shouldn't as the resource technically doesn't exist yet),
		// the controller will actually end up ignoring it since the controller
		// will force a resource creation once it detects that the underlying
		// resource doesn't exist.
		annotationVal, err = GenerateMutableButUnreadableFieldsAnnotation(resource, mutableButUnreadablePaths)
		if err != nil {
			return nil, fmt.Errorf("error ensuring resource '%v' which can have mutable-but-unreadable fields has a value for %v: %w",
				resource.GetNamespace(), MutableButUnreadableFieldsAnnotation, err)
		}
	}

	if err := json.Unmarshal([]byte(annotationVal), &mutableButUnreadableFields); err != nil {
		return nil, fmt.Errorf("error unmarshalling value of %v: %w", MutableButUnreadableFieldsAnnotation, err)
	}
	return mutableButUnreadableFields, nil
}
