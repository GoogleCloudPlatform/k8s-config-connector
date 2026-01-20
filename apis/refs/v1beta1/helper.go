// Copyright 2024 Google LLC
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

package v1beta1

import (
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func GetResourceID(u *unstructured.Unstructured) (string, error) {
	resourceID, _, err := unstructured.NestedString(u.Object, "spec", "resourceID")
	if err != nil {
		return "", fmt.Errorf("reading spec.resourceID from %v %v/%v: %w", u.GroupVersionKind().Kind, u.GetNamespace(), u.GetName(), err)
	}
	if resourceID == "" {
		resourceID = u.GetName()
	}
	return resourceID, nil
}

func GetLocation(u *unstructured.Unstructured) (string, error) {
	location, _, err := unstructured.NestedString(u.Object, "spec", "location")
	if err != nil {
		return "", fmt.Errorf("reading spec.location from %v %v/%v: %w", u.GroupVersionKind().Kind, u.GetNamespace(), u.GetName(), err)
	}
	if location == "" {
		return "", fmt.Errorf("spec.location not set in %v %v/%v: %w", u.GroupVersionKind().Kind, u.GetNamespace(), u.GetName(), err)
	}
	return location, nil
}

// SetRefFields sets the Name, Namespace and External fields on a Ref using reflection.
// It returns an error if a field exists but cannot be set.
func SetRefFields(ref Ref, name, namespace, external string) error {
	val := reflect.ValueOf(ref).Elem()

	if f := val.FieldByName("Name"); f.IsValid() {
		if !f.CanSet() {
			return fmt.Errorf("cannot set Name field")
		}
		f.SetString(name)
	}

	if f := val.FieldByName("Namespace"); f.IsValid() {
		if !f.CanSet() {
			return fmt.Errorf("cannot set Namespace field")
		}
		f.SetString(namespace)
	}

	if f := val.FieldByName("External"); f.IsValid() {
		if !f.CanSet() {
			return fmt.Errorf("cannot set External field")
		}
		f.SetString(external)
	}
	return nil
}