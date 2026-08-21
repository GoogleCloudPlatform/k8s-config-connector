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

package sorted

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Map[K comparable, V any] struct {
	entries []Entry[K, V]
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func (m *Map[K, V]) Get(k K) V {
	for _, entry := range m.entries {
		if entry.Key == k {
			return entry.Value
		}
	}
	var emptyV V
	return emptyV
}

func (m *Map[K, V]) Keys() []K {
	var keys []K
	for _, entry := range m.entries {
		keys = append(keys, entry.Key)
	}
	return keys
}

func (m *Map[K, V]) Entries() []Entry[K, V] {
	return m.entries
}

func (m *Map[K, V]) UnmarshalJSON(b []byte) error {
	dec := json.NewDecoder(bytes.NewReader(b))
	// {
	token, err := dec.Token()
	if err != nil {
		return err
	}
	if s, ok := token.(json.Delim); !ok || s != '{' {
		return fmt.Errorf("expected {")
	}

	for dec.More() {
		// key, assume it is a primitive
		var k K
		token, err := dec.Token()
		if err != nil {
			return err
		}
		if s, ok := token.(K); !ok {
			return fmt.Errorf("expected primitive value, got %T", token)
		} else {
			k = s
		}

		var v V
		if err := dec.Decode(&v); err != nil {
			return fmt.Errorf("decoding value: %w", err)
		}

		m.entries = append(m.entries, Entry[K, V]{Key: k, Value: v})
	}

	// {
	token, err = dec.Token()
	if err != nil {
		return err
	}
	if s, ok := token.(json.Delim); ok && s != '}' {
		return fmt.Errorf("expected }")
	}

	return nil
}
