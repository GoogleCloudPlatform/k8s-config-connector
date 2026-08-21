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

package kube

import "k8s.io/apimachinery/pkg/types"

// NamespaceNameMap is a map from types.NamespacedName to some value of type T.
// We specialize for this case because it lets us share the memory for the namespace string values.
type NamespaceNameMap[T any] map[string]objectsInNamespace[T]

type objectsInNamespace[T any] map[string]T

// NewNamespaceNameMap creates a new NamespaceNameMap.
func NewNamespaceNameMap[T any]() NamespaceNameMap[T] {
	return make(NamespaceNameMap[T])
}

// Set sets the value for the given namespace and name.
func (m NamespaceNameMap[T]) Set(key types.NamespacedName, value T) {
	namespaceObjects, found := m[key.Namespace]
	if !found {
		namespaceObjects = make(objectsInNamespace[T])
		m[key.Namespace] = namespaceObjects
	}
	namespaceObjects[key.Name] = value
}

// Get gets the value for the given namespace and name.
func (m NamespaceNameMap[T]) Get(key types.NamespacedName) (T, bool) {
	namespaceObjects, found := m[key.Namespace]
	if !found {
		var zero T
		return zero, false
	}
	value, ok := namespaceObjects[key.Name]
	return value, ok
}

// Has returns whether the given namespace and name exists.
func (m NamespaceNameMap[T]) Has(key types.NamespacedName) bool {
	namespaceObjects, found := m[key.Namespace]
	if !found {
		return false
	}
	_, ok := namespaceObjects[key.Name]
	return ok
}

// Delete deletes the value for the given namespace and name.
func (m NamespaceNameMap[T]) Delete(key types.NamespacedName) {
	namespaceObjects, found := m[key.Namespace]
	if !found {
		return
	}
	delete(namespaceObjects, key.Name)
}

// Walk calls the callback function for each key/value pair in the map.
func (m NamespaceNameMap[T]) Walk(callback func(types.NamespacedName, T)) {
	for namespace, namespaceObjects := range m {
		for name, value := range namespaceObjects {
			callback(types.NamespacedName{Namespace: namespace, Name: name}, value)
		}
	}
}
