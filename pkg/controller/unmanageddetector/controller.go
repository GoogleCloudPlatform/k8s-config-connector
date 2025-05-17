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

package unmanageddetector

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	v1 "k8s.io/api/apps/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	crlog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type Reconciler struct {
	lifecyclehandler.LifecycleHandler
	mgr manager.Manager
	gvk schema.GroupVersionKind
}

func Add(ctx context.Context, mgr manager.Manager, crd *apiextensions.CustomResourceDefinition) error {
	logger := crlog.FromContext(ctx)

	kind := crd.Spec.Names.Kind
	apiVersion := k8s.GetAPIVersionFromCRD(crd)
	controllerName := fmt.Sprintf("%v-unmanaged-detector", strings.ToLower(kind))

	gvk := schema.GroupVersionKind{
		Group:   crd.Spec.Group,
		Version: k8s.GetVersionFromCRD(crd),
		Kind:    crd.Spec.Names.Kind,
	}
	r, err := NewReconciler(mgr, gvk)
	if err != nil {
		return err
	}
	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       kind,
			"apiVersion": apiVersion,
		},
	}
	_, err = builder.
		ControllerManagedBy(mgr).
		Named(controllerName).
		WithOptions(controller.Options{MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles}).
		For(obj, builder.OnlyMetadata, builder.WithPredicates(Predicate{})).
		Build(r)
	if err != nil {
		return fmt.Errorf("error creating new controller: %w", err)
	}
	logger.Info("Registered unmanaged detector controller", "kind", kind, "apiVersion", apiVersion)
	return nil
}

// NewReconciler creates a new unmanageddetector reconciler, watching objects of the specified GVK
func NewReconciler(mgr manager.Manager, gvk schema.GroupVersionKind) (*Reconciler, error) {
	controllerName := fmt.Sprintf("%v-unmanaged-detector", strings.ToLower(gvk.Kind))
	return &Reconciler{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandlerWithFieldOwner(
			mgr.GetClient(),
			mgr.GetEventRecorderFor(controllerName),
			k8s.UnmanagedDetectorFieldManager,
		),
		mgr: mgr,
		gvk: gvk,
	}, nil
}

func (r *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (res reconcile.Result, err error) {
	logger := crlog.FromContext(ctx)

	logger.Info("starting reconcile", "resource", req.NamespacedName)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(r.gvk)

	if err := r.Get(ctx, req.NamespacedName, u); err != nil {
		if errors.IsNotFound(err) {
			logger.Info("resource not found in API server; finishing reconcile", "resource", req.NamespacedName)
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	yes, err := controllerExistsForNamespace(ctx, u.GetNamespace(), r)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error determining if controller exists for namespace %v: %w", u.GetNamespace(), err)
	}
	if yes {
		// Don't requeue resource for reconciliation; this controller has
		// achieved its purpose.
		logger.Info("controller found for resource's namespace; finishing reconcile", "resource", k8s.GetNamespacedName(u))
		return reconcile.Result{}, nil
	}

	resource, err := k8s.NewResource(u)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("could not parse resource %v: %w", k8s.GetNamespacedName(u), err)
	}

	// Don't requeue resource for reconciliation (unless there's an error
	// during the status update); this controller has achieved its purpose.
	logger.Info("controller not found for resource's namespace; finishing reconcile", "resource", k8s.GetNamespacedName(u))
	return reconcile.Result{}, r.HandleUnmanaged(ctx, resource)
}

func controllerExistsForNamespace(ctx context.Context, namespace string, c client.Client) (yes bool, err error) {
	stsLabelSelectorRaw := fmt.Sprintf("%v=%v,%v=%v",
		k8s.KCCComponentLabel, k8s.ControllerManagerNamePrefix,
		k8s.ScopedNamespaceLabel, namespace,
	)
	stsLabelSelector, err := labels.Parse(stsLabelSelectorRaw)
	if err != nil {
		return false, fmt.Errorf("error parsing '%v' as a label selector: %w", stsLabelSelectorRaw, err)
	}
	stsList := &v1.StatefulSetList{}
	stsOpts := &client.ListOptions{
		Namespace:     k8s.SystemNamespace,
		LabelSelector: stsLabelSelector,
		Limit:         1,
	}
	if err := c.List(ctx, stsList, stsOpts); err != nil {
		return false, fmt.Errorf("error listing controller manager StatefulSets: %w", err)
	}
	return len(stsList.Items) > 0, nil
}
