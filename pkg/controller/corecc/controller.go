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

// logic for reading ConfigConnector (CC) and ConfigConnectorContext (CCC) CRs.
package corecc

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// CoreReconciler is a controller that ONLY observers the ConfigConnector (CC) or ConfigConnectorContext (CCC) CRs.
// Only observing changes allows us to use the same reconciler structure for both and will ensure we don't step
// onto the operator's controllers for CC & CCC.
type CoreReconciler struct {
	mgr                     manager.Manager
	lc                      *LiveConfig
	gvk                     schema.GroupVersionKind
}

var _ reconcile.Reconciler = &CoreReconciler{}

func (r *CoreReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	logger := log.FromContext(ctx)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(r.gvk)

	logger.V(2).Info("fetching resource from API server", "resource", req.NamespacedName)
	if err := r.mgr.GetCache().Get(ctx, req.NamespacedName, u); err != nil {
		if apierrors.IsNotFound(err) {
			logger.Info("resource not found in API server; finishing reconcile", "resource", req.NamespacedName)
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, fmt.Errorf("error fetching %v: %w", req.NamespacedName, err)
	}

	if !u.GetDeletionTimestamp().IsZero() { // the resource is being deleted, nothing to do
		logger.Info("resource is being deleted; finishing reconcile", "resource", req.NamespacedName)
		return reconcile.Result{}, nil
	}

	actuationAction, found, err := unstructured.NestedFieldNoCopy(u.Object, "spec", "actuationAction")
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error getting actuationAction field from %v: %w", req.NamespacedName, err)
	}

	if !found {
		logger.V(2).Info("actuationAction field not found in resource", "resource", req.NamespacedName)
		return reconcile.Result{}, nil
	} else {
		logger.V(2).Info("actuationAction field found in resource", "resource", req.NamespacedName, "value", actuationAction)
		ac, ok := actuationAction.(string)
		if !ok {
			return reconcile.Result{}, fmt.Errorf("actuationAction field is not a string in %v", req.NamespacedName)
		}
		r.lc.Set("actuationAction", ac)
	}

	return reconcile.Result{}, nil
}

func Add(mgr manager.Manager, object client.Object, gvk schema.GroupVersionKind, lc *LiveConfig) error {
	controllerName := fmt.Sprintf("%v-controller", strings.ToLower(gvk.Kind))

	r := &CoreReconciler{mgr: mgr, lc: lc, gvk: gvk}

	c, err := controller.New(controllerName, mgr,
		controller.Options{
			Reconciler:              r,
			MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles,
		},
	)
	if err != nil {
		return fmt.Errorf("error creating %v controller: %v", controllerName, err)
	}

	return c.Watch(source.Kind(mgr.GetCache(), object), &handler.EnqueueRequestForObject{})
}
