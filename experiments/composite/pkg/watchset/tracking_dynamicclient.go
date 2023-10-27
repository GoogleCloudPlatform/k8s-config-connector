// Copyright 2023 Google LLC
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

package watchset

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/klog/v2"
)

var _ dynamic.Interface = &trackingDynamicClient{}

type trackingDynamicClient struct {
	deps *DependencySet

	inner dynamic.Interface
}

func (d *DependencySet) TrackingDynamicClient(inner dynamic.Interface) dynamic.Interface {
	return &trackingDynamicClient{inner: inner, deps: d}
}

func (t *trackingDynamicClient) Resource(resource schema.GroupVersionResource) dynamic.NamespaceableResourceInterface {
	inner := t.inner.Resource(resource)
	return &trackingNamespaceableResourceInterface{
		inner: inner,
		deps:  t.deps,
		gvr:   resource,
		trackingResourceInterface: trackingResourceInterface{
			inner: inner,
			gvr:   resource,
			deps:  t.deps,
		},
	}
}

type trackingNamespaceableResourceInterface struct {
	deps *DependencySet
	gvr  schema.GroupVersionResource

	trackingResourceInterface
	inner dynamic.NamespaceableResourceInterface
}

func (t *trackingNamespaceableResourceInterface) Namespace(ns string) dynamic.ResourceInterface {
	inner := t.inner.Namespace(ns)
	return &trackingResourceInterface{
		deps:  t.deps,
		gvr:   t.gvr,
		ns:    ns,
		inner: inner,
	}
}

type trackingResourceInterface struct {
	deps  *DependencySet
	gvr   schema.GroupVersionResource
	ns    string
	inner dynamic.ResourceInterface
}

func (t *trackingResourceInterface) Create(ctx context.Context, obj *unstructured.Unstructured, options metav1.CreateOptions, subresources ...string) (*unstructured.Unstructured, error) {
	panic("trackingResourceInterface method not implemented")
}

func (t *trackingResourceInterface) Update(ctx context.Context, obj *unstructured.Unstructured, options metav1.UpdateOptions, subresources ...string) (*unstructured.Unstructured, error) {
	panic("trackingResourceInterface method not implemented")
}
func (t *trackingResourceInterface) UpdateStatus(ctx context.Context, obj *unstructured.Unstructured, options metav1.UpdateOptions) (*unstructured.Unstructured, error) {
	panic("trackingResourceInterface method not implemented")
}
func (t *trackingResourceInterface) Delete(ctx context.Context, name string, options metav1.DeleteOptions, subresources ...string) error {
	panic("trackingResourceInterface method not implemented")
}
func (t *trackingResourceInterface) DeleteCollection(ctx context.Context, options metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	panic("trackingResourceInterface method not implemented")
}
func (t *trackingResourceInterface) Get(ctx context.Context, name string, options metav1.GetOptions, subresources ...string) (*unstructured.Unstructured, error) {
	panic("trackingResourceInterface method not implemented")
}

func (t *trackingResourceInterface) List(ctx context.Context, opts metav1.ListOptions) (*unstructured.UnstructuredList, error) {
	klog.Warningf("tracking call to List ns=%v, gvr=%v", t.ns, t.gvr)
	t.deps.WatchList(t.gvr, t.ns, opts)

	u, err := t.inner.List(ctx, opts)
	return u, err
}

func (t *trackingResourceInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	panic("trackingResourceInterface method not implemented")
}

func (t *trackingResourceInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, options metav1.PatchOptions, subresources ...string) (*unstructured.Unstructured, error) {
	nn := types.NamespacedName{
		Namespace: t.ns,
		Name:      name,
	}
	t.deps.WatchObject(t.gvr, nn, "")

	u, err := t.inner.Patch(ctx, name, pt, data, options, subresources...)
	return u, err
}

func (t *trackingResourceInterface) Apply(ctx context.Context, name string, obj *unstructured.Unstructured, options metav1.ApplyOptions, subresources ...string) (*unstructured.Unstructured, error) {
	panic("trackingResourceInterface method not implemented")
}

func (t *trackingResourceInterface) ApplyStatus(ctx context.Context, name string, obj *unstructured.Unstructured, options metav1.ApplyOptions) (*unstructured.Unstructured, error) {
	panic("trackingResourceInterface method not implemented")
}
