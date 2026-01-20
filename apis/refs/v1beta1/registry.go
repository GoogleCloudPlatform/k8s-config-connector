// Copyright 2026 Google LLC
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
	"sync"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	registryMu sync.Mutex
	registry   = make(map[schema.GroupKind]reflect.Type)
)

// Register registers a Ref implementation.
// It is thread-safe.
func Register(ref Ref) {
	registryMu.Lock()
	defer registryMu.Unlock()
	registry[ref.GetGVK().GroupKind()] = reflect.TypeOf(ref).Elem()
}

// NewRef returns a new instance of Ref for the given GroupKind.
// It is thread-safe.
func NewRef(gk schema.GroupKind) (Ref, error) {
	registryMu.Lock()
	defer registryMu.Unlock()
	typ, ok := registry[gk]
	if !ok {
		return nil, fmt.Errorf("no Ref registered for GroupKind %v", gk)
	}
	return reflect.New(typ).Interface().(Ref), nil
}

// NewRefByKind returns a new instance of Ref for the given Kind.
// It is thread-safe.
// Note: This iterates over the registry, so it is O(N).
// It also assumes that Kind is unique across Groups.
func NewRefByKind(kind string) (Ref, error) {
	registryMu.Lock()
	defer registryMu.Unlock()
	for gk, typ := range registry {
		if gk.Kind == kind {
			return reflect.New(typ).Interface().(Ref), nil
		}
	}
	return nil, fmt.Errorf("no Ref registered for Kind %q", kind)
}
