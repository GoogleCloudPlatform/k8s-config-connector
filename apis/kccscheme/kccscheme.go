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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// NewObject returns a new strongly-typed runtime.Object for the given GroupKind.
func NewObject(gk schema.GroupKind) (runtime.Object, error) {
	return refs.NewObject(gk)
}

// PreferredGVK returns the preferred GroupVersionKind for the given GroupKind.
func PreferredGVK(gk schema.GroupKind) (schema.GroupVersionKind, bool) {
	return refs.PreferredGVK(gk)
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
