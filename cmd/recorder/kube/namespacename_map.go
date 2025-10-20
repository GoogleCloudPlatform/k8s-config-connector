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

// Snapshot returns a deep-copy of the map, as a naive map[types.NamespacedName]T.
func (m NamespaceNameMap[T]) Snapshot() map[types.NamespacedName]T {
	out := make(map[types.NamespacedName]T)
	for namespace, namespaceObjects := range m {
		for name, value := range namespaceObjects {
			out[types.NamespacedName{Namespace: namespace, Name: name}] = value
		}
	}
	return out
}

// Walk calls the given function for each key/value pair in the map.
func (m NamespaceNameMap[T]) Walk(fn func(types.NamespacedName, T)) {
	for namespace, namespaceObjects := range m {
		for name, value := range namespaceObjects {
			fn(types.NamespacedName{Namespace: namespace, Name: name}, value)
		}
	}
}
