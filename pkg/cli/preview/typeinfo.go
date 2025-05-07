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

package preview

import (
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// typeStore is a store of type information.
// It is similar to runtime.Scheme, but avoids a lot of the complexity / weight.
type typeStore struct {
	restMapper meta.RESTMapper
	scheme     *runtime.Scheme
}

// typeInfo is the type information for a Kubernetes kind/resource
type typeInfo struct {
	factory func() Object
	gvr     GroupVersionResource
	gvk     schema.GroupVersionKind
}

// getTypeInfoForGVK returns the type information for the given GVK.
func (s *typeStore) getTypeInfoForGVK(gvk schema.GroupVersionKind) (*typeInfo, error) {
	obj, err := s.scheme.New(gvk)
	if err != nil {
		return nil, fmt.Errorf("unhandled gvk trying to create informer by gvk for %v: %w", gvk, err)
	}
	return s.getTypeInfo(obj.(client.Object))
}

// getTypeInfo returns the type information for the given object.
func (s *typeStore) getTypeInfo(obj client.Object) (*typeInfo, error) {
	typeInfo := &typeInfo{}

	gvk := obj.GetObjectKind().GroupVersionKind()
	if gvk.Kind == "" {
		kinds, isUnversioned, err := s.scheme.ObjectKinds(obj)
		if err != nil {
			return nil, fmt.Errorf("getting kinds for %T: %w", obj, err)
		}
		if isUnversioned {
			return nil, fmt.Errorf("got unversioned type %T", obj)
		}
		if len(kinds) == 0 {
			return nil, fmt.Errorf("cannot find kind for %T", obj)
		}
		if len(kinds) > 1 {
			return nil, fmt.Errorf("found multiple kinds for %T: %v", obj, kinds)
		}
		gvk = kinds[0]

		rt := reflect.TypeOf(obj)
		if rt.Kind() == reflect.Ptr {
			rt = rt.Elem()
		}
		// Idea for the future, we could register well known types here and avoid the need for a scheme?
		// name := rt.Name()
		// pkgPath := rt.PkgPath()

		// switch pkgPath {
		// case "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1":
		// 	switch name {
		// 	case "CustomResourceDefinition":
		// 		gvk = schema.GroupVersionKind{Kind: "CustomResourceDefinition", Group: "apiextensions.k8s.io", Version: "v1"}
		// 	}
		// }

		typeInfo.factory = func() Object {
			return reflect.New(rt).Interface().(Object)
		}
	} else {
		switch obj := obj.(type) {
		case *metav1.PartialObjectMetadata:
			// TODO: Do partial fetch
			typeInfo.factory = func() Object {
				return &metav1.PartialObjectMetadata{}
			}
		case *unstructured.Unstructured:
			typeInfo.factory = func() Object {
				return &unstructured.Unstructured{}
			}
		default:
			return nil, fmt.Errorf("cannot build factory for %T", obj)
		}
	}

	restMapping, err := s.restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return nil, fmt.Errorf("getting rest mapping: %w", err)
	}
	typeInfo.gvk = gvk
	typeInfo.gvr = GroupVersionResource(restMapping.Resource)
	return typeInfo, nil
}

// GroupResource returns the group and resource for the type info.
func (t *typeInfo) GroupResource() schema.GroupResource {
	return schema.GroupResource{Group: t.gvr.Group, Resource: t.gvr.Resource}
}
