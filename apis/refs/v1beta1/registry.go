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

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	registryMu     sync.Mutex
	registry       = make(map[schema.GroupKind]reflect.Type)
	objectRegistry = make(map[schema.GroupKind]reflect.Type)
	gvkRegistry    = make(map[schema.GroupKind]schema.GroupVersionKind)
)

// Register registers a Ref implementation and optionally its corresponding runtime.Object.
// It is thread-safe.
func Register(ref Ref, objs ...runtime.Object) {
	registryMu.Lock()
	defer registryMu.Unlock()
	registry[ref.GetGVK().GroupKind()] = reflect.TypeOf(ref).Elem()
	if len(objs) > 0 && objs[0] != nil {
		obj := objs[0]
		objectRegistry[ref.GetGVK().GroupKind()] = reflect.TypeOf(obj).Elem()
		gvkRegistry[ref.GetGVK().GroupKind()] = ref.GetGVK()
	}
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
	var found reflect.Type
	for gk, typ := range registry {
		if gk.Kind == kind {
			if found != nil {
				return nil, fmt.Errorf("multiple Refs registered for Kind %q", kind)
			}
			found = typ
		}
	}
	if found != nil {
		return reflect.New(found).Interface().(Ref), nil
	}
	return nil, fmt.Errorf("no Ref registered for Kind %q", kind)
}

// NewObject returns a new instance of runtime.Object for the given GroupKind.
// It is thread-safe.
func NewObject(gk schema.GroupKind) (runtime.Object, error) {
	registryMu.Lock()
	defer registryMu.Unlock()
	typ, ok := objectRegistry[gk]
	if !ok {
		return nil, fmt.Errorf("no Object registered for GroupKind %v", gk)
	}
	return reflect.New(typ).Interface().(runtime.Object), nil
}

// PreferredGVK returns the preferred GroupVersionKind for the given GroupKind.
// It is thread-safe.
func PreferredGVK(gk schema.GroupKind) (schema.GroupVersionKind, bool) {
	registryMu.Lock()
	defer registryMu.Unlock()
	gvk, ok := gvkRegistry[gk]
	return gvk, ok
}
