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

package kccscheme

import (
	"fmt"
	"reflect"
	"sync"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	registryMu     sync.Mutex
	objectRegistry = make(map[schema.GroupKind]reflect.Type)
	gvkRegistry    = make(map[schema.GroupKind]schema.GroupVersionKind)
	refRegistry    = make(map[schema.GroupKind]reflect.Type)
)

// RegisterType registers the runtime.Object implementation for a given GroupVersionKind.
func RegisterType(gvk schema.GroupVersionKind, obj runtime.Object) {
	registryMu.Lock()
	defer registryMu.Unlock()
	objectRegistry[gvk.GroupKind()] = reflect.TypeOf(obj).Elem()
	gvkRegistry[gvk.GroupKind()] = gvk
}

// RegisterRef registers a Ref implementation mapping to its GroupVersionKind.
func RegisterRef(ref interface{}, gvk schema.GroupVersionKind) {
	registryMu.Lock()
	defer registryMu.Unlock()
	refRegistry[gvk.GroupKind()] = reflect.TypeOf(ref).Elem()
}

// NewRef returns a new instance of Ref for the given GroupKind.
func NewRef(gk schema.GroupKind) (interface{}, error) {
	registryMu.Lock()
	defer registryMu.Unlock()
	typ, ok := refRegistry[gk]
	if !ok {
		return nil, fmt.Errorf("no Ref registered in kccscheme for GroupKind %v", gk)
	}
	return reflect.New(typ).Interface(), nil
}

// NewRefByKind returns a new instance of Ref for the given Kind.
func NewRefByKind(kind string) (interface{}, error) {
	registryMu.Lock()
	defer registryMu.Unlock()
	var found reflect.Type
	for gk, typ := range refRegistry {
		if gk.Kind == kind {
			if found != nil {
				return nil, fmt.Errorf("multiple Refs registered in kccscheme for Kind %q", kind)
			}
			found = typ
		}
	}
	if found != nil {
		return reflect.New(found).Interface(), nil
	}
	return nil, fmt.Errorf("no Ref registered in kccscheme for Kind %q", kind)
}

// NewObject returns a new strongly-typed runtime.Object for the given GroupKind.
func NewObject(gk schema.GroupKind) (runtime.Object, error) {
	registryMu.Lock()
	defer registryMu.Unlock()
	typ, ok := objectRegistry[gk]
	if !ok {
		return nil, fmt.Errorf("no Object registered in kccscheme for GroupKind %v", gk)
	}
	return reflect.New(typ).Interface().(runtime.Object), nil
}

// PreferredGVK returns the preferred GroupVersionKind for the given GroupKind.
func PreferredGVK(gk schema.GroupKind) (schema.GroupVersionKind, bool) {
	registryMu.Lock()
	defer registryMu.Unlock()
	gvk, ok := gvkRegistry[gk]
	return gvk, ok
}

// ObjectKinds is a helper to find the runtime object GVK for an object.
func ObjectKinds(obj runtime.Object) ([]schema.GroupVersionKind, error) {
	if unstructuredObj, ok := obj.(*runtime.Unknown); ok {
		gvk := unstructuredObj.GetObjectKind().GroupVersionKind()
		if !gvk.Empty() {
			return []schema.GroupVersionKind{gvk}, nil
		}
	}
	gvk := obj.GetObjectKind().GroupVersionKind()
	if !gvk.Empty() {
		return []schema.GroupVersionKind{gvk}, nil
	}
	if gk, ok := PreferredGVK(gvk.GroupKind()); ok {
		return []schema.GroupVersionKind{gk}, nil
	}
	return nil, fmt.Errorf("object %T is not registered in kccscheme", obj)
}
