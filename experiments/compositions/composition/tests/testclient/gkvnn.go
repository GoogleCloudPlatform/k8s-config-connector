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

// Package util provides utilities for e2e testing.
package testclient

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GVKNN - id for k8s object
type GVKNN struct {
	schema.GroupVersionKind
	types.NamespacedName
}

// ExtractGVKNN - returns GVKNN of obj
func ExtractGVKNN(obj client.Object) *GVKNN {
	return &GVKNN{
		obj.GetObjectKind().GroupVersionKind(),
		types.NamespacedName{
			Namespace: obj.GetNamespace(),
			Name:      obj.GetName(),
		},
	}
}

// ExtractGVKNNs - returns GVKNNs of objs
func ExtractGVKNNs(objs []*unstructured.Unstructured) []*GVKNN {
	ids := make([]*GVKNN, len(objs))
	for i, obj := range objs {
		ids[i] = ExtractGVKNN(obj)
	}
	return ids
}

// MakeObject - create object with GVKNN of id
func (id *GVKNN) MakeObject() *unstructured.Unstructured {
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(id.GroupVersionKind)
	obj.SetNamespace(id.Namespace)
	obj.SetName(id.Name)
	return obj
}

// String - returns GVKNN delimited by "/"
func (id *GVKNN) String() string {
	if id.Namespace == "" {
		return fmt.Sprintf("%s/%s", id.GroupVersionKind.Kind, id.Name)
	}
	return fmt.Sprintf("%s/%s/%s", id.GroupVersionKind.Kind, id.Namespace, id.Name)
}
