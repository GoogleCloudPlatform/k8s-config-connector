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

package k8s

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type contextKeyUnstructured struct{}

// ContextWithUnstructured returns a new context carrying the given Unstructured object.
func ContextWithUnstructured(ctx context.Context, u *unstructured.Unstructured) context.Context {
	if u == nil {
		return ctx
	}
	return context.WithValue(ctx, contextKeyUnstructured{}, u)
}

// UnstructuredFromContext returns a deep copy of the Unstructured object stored in context,
// provided its GroupVersionKind (if specified) and NamespacedName match.
func UnstructuredFromContext(ctx context.Context, nn types.NamespacedName, gvk schema.GroupVersionKind) (*unstructured.Unstructured, bool) {
	val := ctx.Value(contextKeyUnstructured{})
	if val == nil {
		return nil, false
	}
	u, ok := val.(*unstructured.Unstructured)
	if !ok || u == nil {
		return nil, false
	}
	if !gvk.Empty() && u.GroupVersionKind() != gvk {
		return nil, false
	}
	if (types.NamespacedName{Namespace: u.GetNamespace(), Name: u.GetName()}) != nn {
		return nil, false
	}
	return u.DeepCopy(), true
}

// FetchUnstructured retrieves the object identified by nn and u.GroupVersionKind().
// If a matching Unstructured object is present in ctx (e.g. passed from ParentReconciler),
// it populates u with a deep copy of that object. Otherwise, it falls back to c.Get(ctx, nn, u).
func FetchUnstructured(ctx context.Context, c client.Reader, nn types.NamespacedName, u *unstructured.Unstructured) error {
	if cached, ok := UnstructuredFromContext(ctx, nn, u.GroupVersionKind()); ok {
		*u = *cached
		return nil
	}
	return c.Get(ctx, nn, u)
}

// FetchObject retrieves the object identified by nn and obj.GetObjectKind().GroupVersionKind().
// If a matching Unstructured object is present in ctx (e.g. passed from ParentReconciler),
// it populates obj with that object. Otherwise, it falls back to c.Get(ctx, nn, obj).
// Note: For typed client.Object structs created via new(T) or &T{},
// obj.GetObjectKind().GroupVersionKind() is empty unless explicitly set beforehand.
// Callers using typed structs should ensure the GVK is set on obj (e.g. via obj.SetGroupVersionKind)
// if strict GVK validation in UnstructuredFromContext is desired.
func FetchObject(ctx context.Context, c client.Reader, nn types.NamespacedName, obj client.Object) error {
	gvk := obj.GetObjectKind().GroupVersionKind()
	if cached, ok := UnstructuredFromContext(ctx, nn, gvk); ok {
		if u, isUnstructured := obj.(*unstructured.Unstructured); isUnstructured {
			*u = *cached
			return nil
		}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(cached.Object, obj); err != nil {
			return fmt.Errorf("error converting unstructured from context to %T: %w", obj, err)
		}
		return nil
	}
	return c.Get(ctx, nn, obj)
}
