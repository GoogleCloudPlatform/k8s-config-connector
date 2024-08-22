// Copyright 2022 Google LLC
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

package predicate

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	k8spredicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// ReconcileGate allows controllers to select which resources they are enabled for, based on
// features of the resource to-be-reconciled. This allows for partially enabling a controller.
type ReconcileGate interface {
	// ShouldReconcile returns true if the reconciler should be used to for the resource.
	ShouldReconcile(o *unstructured.Unstructured) bool
}

// ReconcilePredicate generates a controller-runtime predicate based on a ReconcileGate.
type ReconcilePredicate struct {
	gvk schema.GroupVersionKind
	rg  ReconcileGate
	c   client.Client
}

var _ k8spredicate.Predicate = &ReconcilePredicate{}

func NewReconcilePredicate(c client.Client, gvk schema.GroupVersionKind, rg ReconcileGate) *ReconcilePredicate {
	return &ReconcilePredicate{
		c:   c,
		gvk: gvk,
		rg:  rg,
	}
}

func (p *ReconcilePredicate) Create(e event.CreateEvent) bool {
	obj, err := getUnstructuredObjWithGVK(p.c, e.Object, p.gvk)
	if err != nil {
		return false
	}
	return p.rg.ShouldReconcile(obj)
}

func (p *ReconcilePredicate) Delete(e event.DeleteEvent) bool {
	obj, err := getUnstructuredObjWithGVK(p.c, e.Object, p.gvk)
	if err != nil {
		return false
	}
	return p.rg.ShouldReconcile(obj)
}

func (p *ReconcilePredicate) Update(e event.UpdateEvent) bool {
	obj, err := getUnstructuredObjWithGVK(p.c, e.ObjectNew, p.gvk)
	if err != nil {
		return false
	}
	return p.rg.ShouldReconcile(obj)
}

func (p *ReconcilePredicate) Generic(e event.GenericEvent) bool {
	obj, err := getUnstructuredObjWithGVK(p.c, e.Object, p.gvk)
	if err != nil {
		return false
	}
	return p.rg.ShouldReconcile(obj)
}

// InverseReconcilePredicate generates a controller-runtime predicate based on the inverse of a ReconcileGate.
type InverseReconcilePredicate struct {
	gvk schema.GroupVersionKind
	rg  ReconcileGate
	c   client.Client
}

var _ k8spredicate.Predicate = &InverseReconcilePredicate{}

func NewInverseReconcilePredicate(c client.Client, gvk schema.GroupVersionKind, rg ReconcileGate) *InverseReconcilePredicate {
	return &InverseReconcilePredicate{
		c:   c,
		gvk: gvk,
		rg:  rg,
	}
}

func (p *InverseReconcilePredicate) Create(e event.CreateEvent) bool {
	obj, err := getUnstructuredObjWithGVK(p.c, e.Object, p.gvk)
	if err != nil {
		return false
	}
	return !p.rg.ShouldReconcile(obj)
}

func (p *InverseReconcilePredicate) Delete(e event.DeleteEvent) bool {
	obj, err := getUnstructuredObjWithGVK(p.c, e.Object, p.gvk)
	if err != nil {
		return false
	}
	return !p.rg.ShouldReconcile(obj)
}

func (p *InverseReconcilePredicate) Update(e event.UpdateEvent) bool {
	obj, err := getUnstructuredObjWithGVK(p.c, e.ObjectNew, p.gvk)
	if err != nil {
		return false
	}
	return !p.rg.ShouldReconcile(obj)
}

func (p *InverseReconcilePredicate) Generic(e event.GenericEvent) bool {
	obj, err := getUnstructuredObjWithGVK(p.c, e.Object, p.gvk)
	if err != nil {
		return false
	}
	return !p.rg.ShouldReconcile(obj)
}

// getUnstructuredObjWithGVK uses the provided client to fetch an object (as unstructured) with a given GVK.
// This helper fn is necessary for the ReconcileGate predicates because both the direct and terraform
// controllers specify the option `builder.OnlyMetadata` to only cache metadata about the objects they are
// watching. Therefore, the objects provided to the predicate fns are only `metav1.PartialObjectMetadata`
// types, and do not contain any of the spec values.
func getUnstructuredObjWithGVK(c client.Client, o client.Object, gvk schema.GroupVersionKind) (*unstructured.Unstructured, error) {
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(gvk)
	key := types.NamespacedName{
		Name:      o.GetName(),
		Namespace: o.GetNamespace(),
	}
	if err := c.Get(context.Background(), key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}
