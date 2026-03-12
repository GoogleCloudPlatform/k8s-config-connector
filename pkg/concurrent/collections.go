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

package concurrent

import "sync"

// ExpiringMap is a thread-safe map that automatically expires entries based on a validity function.
type ExpiringMap[K comparable, V any] struct {
	isValid func(value V) bool
	inner   *ConcurrentMap[K, V]
}

// NewExpiringMap creates a new ExpiringMap with the specified validity function.
func NewExpiringMap[K comparable, V any](isValid func(value V) bool) *ExpiringMap[K, V] {
	return &ExpiringMap[K, V]{
		isValid: isValid,
		inner:   NewConcurrentMap[K, V](),
	}
}

// GetOrCompute retrieves the value for the given key from the map,
// or computes it using the provided function if not present or expired.
func (m *ExpiringMap[K, V]) GetOrCompute(key K, compute func() (V, error)) (V, error) {
	value, err := m.inner.GetOrCompute(key, compute)
	if err != nil {
		return value, err
	}
	if m.isValid(value) {
		return value, nil
	}

	m.inner.Delete(key)
	return m.inner.GetOrCompute(key, compute)
}

// Set sets the value for the given key in the map.
func (m *ExpiringMap[K, V]) Set(k K, v V) {
	m.inner.Set(k, v)
}

// ConcurrentMap is a thread-safe map with lazy computation capabilities.
type ConcurrentMap[K comparable, V any] struct {
	mutex sync.Mutex
	data  map[K]*optional[V]
}

// optional represents a value that may or may not be set.
// It also has lazy computation capabilities.
type optional[V any] struct {
	mutex      sync.Mutex
	valueIsSet bool
	value      V
}

// newOptional creates a new optional with the given value set.
func newOptional[V any](value V) *optional[V] {
	return &optional[V]{
		valueIsSet: true,
		value:      value,
	}
}

// newEmptyOptional creates a new optional with no value set.
func newEmptyOptional[V any]() *optional[V] {
	return &optional[V]{
		valueIsSet: false,
	}
}

// NewConcurrentMap creates a new ConcurrentMap.
func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		data: make(map[K]*optional[V]),
	}
}

// Get retrieves the value for the given key from the map.
func (m *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	m.mutex.Lock()
	value, ok := m.data[key]
	m.mutex.Unlock()
	if !ok {
		var defaultV V
		return defaultV, ok
	}

	return value.get()
}

// get retrieves the value and whether it is set.
func (ov *optional[V]) get() (V, bool) {
	ov.mutex.Lock()
	defer ov.mutex.Unlock()

	return ov.value, ov.valueIsSet
}

// Set sets the value for the given key in the map.
func (m *ConcurrentMap[K, V]) Set(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	ov := newOptional(value)
	m.data[key] = ov
}

// Delete removes the key from the map.
func (m *ConcurrentMap[K, V]) Delete(key K) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.data, key)
}

// GetOrCompute retrieves the value for the given key from the map,
// or computes it using the provided function if not present.
func (m *ConcurrentMap[K, V]) GetOrCompute(key K, compute func() (V, error)) (V, error) {
	m.mutex.Lock()
	value := m.data[key]
	if value == nil {
		value = newEmptyOptional[V]()
		m.data[key] = value
	}
	m.mutex.Unlock()

	return value.getOrCompute(compute)
}

// getOrCompute retrieves the value if set, or computes it using the provided function.
func (ov *optional[V]) getOrCompute(compute func() (V, error)) (V, error) {
	ov.mutex.Lock()
	defer ov.mutex.Unlock()

	if ov.valueIsSet {
		return ov.value, nil
	}

	value, err := compute()
	if err != nil {
		var defaultV V
		return defaultV, err
	}

	ov.value = value
	ov.valueIsSet = true
	return value, nil
}
