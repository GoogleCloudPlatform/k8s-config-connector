/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package mockkubeapiserver

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/uuid"
)

type MemoryStorage struct {
	schemaMutex      sync.Mutex
	schema           mockSchema
	resourceStorages map[schema.GroupResource]*resourceStorage

	resourceVersionClock int64
}

func NewMemoryStorage() *MemoryStorage {
	s := &MemoryStorage{
		resourceStorages:     make(map[schema.GroupResource]*resourceStorage),
		resourceVersionClock: 1,
	}
	return s
}

type mockSchema struct {
	resources []*ResourceInfo
}

type ResourceInfo struct {
	API     metav1.APIResource
	GVR     schema.GroupVersionResource
	GVK     schema.GroupVersionKind
	ListGVK schema.GroupVersionKind

	storage *resourceStorage
}

// AddObject pre-creates an object
func (s *MemoryStorage) AddObject(obj *unstructured.Unstructured) error {
	ctx := context.Background()

	gv, err := schema.ParseGroupVersion(obj.GetAPIVersion())
	if err != nil {
		return fmt.Errorf("cannot parse apiVersion %q: %w", obj.GetAPIVersion(), err)
	}
	kind := obj.GetKind()

	id := types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}

	gvk := gv.WithKind(kind)

	resource := s.findResourceByGVK(gvk)
	if resource == nil {
		return fmt.Errorf("object group/version/kind %v not known", gvk)
	}

	return s.CreateObject(ctx, resource, id, obj)
}

func (s *MemoryStorage) GetObject(ctx context.Context, resource *ResourceInfo, id types.NamespacedName) (*unstructured.Unstructured, bool, error) {
	resource.storage.mutex.Lock()
	defer resource.storage.mutex.Unlock()

	object := resource.storage.objects[id]
	if object == nil {
		return nil, false, nil
	}

	return object, true, nil
}

type ListFilter struct {
	Namespace string
}

func (s *MemoryStorage) ListObjects(ctx context.Context, resource *ResourceInfo, filter ListFilter) (*unstructured.UnstructuredList, error) {
	resource.storage.mutex.Lock()
	defer resource.storage.mutex.Unlock()

	ret := &unstructured.UnstructuredList{}

	for _, obj := range resource.storage.objects {
		if filter.Namespace != "" {
			if obj.GetNamespace() != filter.Namespace {
				continue
			}
		}
		ret.Items = append(ret.Items, *obj)
	}

	rv := strconv.FormatInt(s.resourceVersionClock, 10)
	ret.SetResourceVersion(rv)

	return ret, nil
}

func (s *MemoryStorage) CreateObject(ctx context.Context, resource *ResourceInfo, id types.NamespacedName, u *unstructured.Unstructured) error {
	resource.storage.mutex.Lock()
	defer resource.storage.mutex.Unlock()

	_, found := resource.storage.objects[id]
	if found {
		return apierrors.NewAlreadyExists(resource.GVR.GroupResource(), id.Name)
	}

	u.SetCreationTimestamp(v1.Now())

	uid := uuid.NewUUID()
	u.SetUID(uid)

	rv := strconv.FormatInt(s.resourceVersionClock, 10)
	s.resourceVersionClock++
	u.SetResourceVersion(rv)

	resource.storage.objects[id] = u
	s.objectChanged(u)

	resource.storage.broadcastEventHoldingLock(ctx, "ADDED", u)

	return nil
}

func (s *MemoryStorage) UpdateObject(ctx context.Context, resource *ResourceInfo, id types.NamespacedName, u *unstructured.Unstructured) error {
	resource.storage.mutex.Lock()
	defer resource.storage.mutex.Unlock()

	_, found := resource.storage.objects[id]
	if !found {
		return apierrors.NewAlreadyExists(resource.GVR.GroupResource(), id.Name)
	}

	rv := strconv.FormatInt(s.resourceVersionClock, 10)
	s.resourceVersionClock++
	u.SetResourceVersion(rv)

	resource.storage.objects[id] = u
	s.objectChanged(u)

	resource.storage.broadcastEventHoldingLock(ctx, "MODIFIED", u)

	return nil
}

func (s *MemoryStorage) DeleteObject(ctx context.Context, resource *ResourceInfo, id types.NamespacedName) (*unstructured.Unstructured, error) {
	resource.storage.mutex.Lock()
	defer resource.storage.mutex.Unlock()

	deletedObj, found := resource.storage.objects[id]
	if !found {
		// TODO: return apierrors something?
		return nil, apierrors.NewNotFound(resource.GVR.GroupResource(), id.Name)
	}
	delete(resource.storage.objects, id)
	s.objectChanged(deletedObj)

	resource.storage.broadcastEventHoldingLock(ctx, "DELETED", deletedObj)

	return deletedObj, nil
}

// RegisterType registers a type with the schema for the mock kubeapiserver
func (s *MemoryStorage) RegisterType(gvk schema.GroupVersionKind, resource string, scope meta.RESTScope) {
	s.schemaMutex.Lock()
	defer s.schemaMutex.Unlock()

	gvr := gvk.GroupVersion().WithResource(resource)
	gr := gvr.GroupResource()

	storage := &resourceStorage{
		GroupResource: gr,
		objects:       make(map[types.NamespacedName]*unstructured.Unstructured),
	}

	// TODO: share storage across different versions
	s.resourceStorages[gr] = storage

	r := &ResourceInfo{
		API: metav1.APIResource{
			Name:    resource,
			Group:   gvk.Group,
			Version: gvk.Version,
			Kind:    gvk.Kind,
		},
		GVK:     gvk,
		GVR:     gvr,
		storage: storage,
	}
	r.ListGVK = gvk.GroupVersion().WithKind(gvk.Kind + "List")

	if scope.Name() == meta.RESTScopeNameNamespace {
		r.API.Namespaced = true
	}

	s.schema.resources = append(s.schema.resources, r)
}

func (s *MemoryStorage) AllResources() []metav1.APIResource {
	s.schemaMutex.Lock()
	defer s.schemaMutex.Unlock()

	var ret []metav1.APIResource
	for _, resource := range s.schema.resources {
		ret = append(ret, resource.API)
	}
	return ret
}

func (s *MemoryStorage) FindResource(gr schema.GroupResource) *ResourceInfo {
	s.schemaMutex.Lock()
	defer s.schemaMutex.Unlock()

	for _, resource := range s.schema.resources {
		if resource.GVR.GroupResource() == gr {
			return resource
		}
	}
	return nil
}

func (s *MemoryStorage) findResourceByGVK(gvk schema.GroupVersionKind) *ResourceInfo {
	s.schemaMutex.Lock()
	defer s.schemaMutex.Unlock()

	for _, resource := range s.schema.resources {
		if resource.GVK == gvk {
			return resource
		}
	}
	return nil
}
