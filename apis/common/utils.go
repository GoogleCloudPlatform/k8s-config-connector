// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ValueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}

func LazyPtr[V comparable](v V) *V {
	var defaultV V
	if v == defaultV {
		return nil
	}
	return &v
}

func ToStructuredType[T client.Object](obj client.Object) (T, error) {
	if typed, ok := obj.(T); ok {
		return typed, nil
	}

	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		var zero T
		return zero, fmt.Errorf("object is neither of type %T nor *unstructured.Unstructured: %T", zero, obj)
	}

	// T is a pointer type, e.g., *SecurityCenterMuteConfig.
	// We need to create a new instance of the underlying struct.
	var zero T
	tType := reflect.TypeOf(zero)
	if tType.Kind() != reflect.Ptr {
		return zero, fmt.Errorf("expected pointer type for T, got %v", tType)
	}

	// Create a new instance of the element type, which is the struct.
	elemType := tType.Elem()
	newObjVal := reflect.New(elemType)
	res := newObjVal.Interface().(T)

	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, res); err != nil {
		return zero, fmt.Errorf("error converting unstructured to %T: %w", zero, err)
	}

	return res, nil
}
