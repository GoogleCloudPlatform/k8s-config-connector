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

package controllers

import (
	"context"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon"

	addonsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/composite/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/composite/pkg/watchset"
)

var _ reconcile.Reconciler = &CompositeDefinitionReconciler{}

// CompositeDefinitionReconciler reconciles a CompositeDefinition object
type CompositeDefinitionReconciler struct {
	// client.Client
	client client.Client
	mgr    ctrl.Manager

	finalizer string

	watchsets *watchset.Manager

	// TODO: Locking
	instanceReconcilers map[types.NamespacedName]*instanceReconcilerRunner
	controller          controller.Controller
}

//+kubebuilder:rbac:groups=addons.kpt.io,resources=compositedefinitions,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=addons.kpt.io,resources=compositedefinitions/status,verbs=get;update;patch

func (r *CompositeDefinitionReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	//log := klog.FromContext(ctx)

	id := req.NamespacedName
	subject := &addonsv1alpha1.CompositeDefinition{}
	if err := r.client.Get(ctx, id, subject); err != nil {
		if apierrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	if !subject.ObjectMeta.DeletionTimestamp.IsZero() {
		reconciler := r.instanceReconcilers[id]
		if reconciler != nil {
			if err := reconciler.stop(); err != nil {
				return ctrl.Result{}, err
			}
			delete(r.instanceReconcilers, id)
		}

		// remove our finalizer from the list and update it.
		if changed := controllerutil.RemoveFinalizer(subject, r.finalizer); changed {
			if err := r.client.Update(ctx, subject); err != nil {
				return ctrl.Result{}, err
			}
		}

		// Stop reconciliation as the item is being deleted
		return ctrl.Result{}, nil
	}

	if added := controllerutil.AddFinalizer(subject, r.finalizer); added {
		if err := r.client.Update(ctx, subject); err != nil {
			return reconcile.Result{}, err
		}
	}

	reconciler := r.instanceReconcilers[id]
	if reconciler != nil {
		if err := reconciler.stop(); err != nil {
			return ctrl.Result{}, err
		}
		delete(r.instanceReconcilers, id)
	}

	runner, err := newInstanceReconcilerRunner(r.mgr, r.watchsets, subject)
	if err != nil {
		return ctrl.Result{}, err
	}

	runner.start()
	r.instanceReconcilers[id] = runner

	return reconcile.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CompositeDefinitionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	addon.Init()

	r.finalizer = "CompositeDefinitionReconciler"

	r.instanceReconcilers = make(map[types.NamespacedName]*instanceReconcilerRunner)

	// TODO: Share watchset manager across controllers
	watchsets, err := watchset.NewManager(mgr)
	if err != nil {
		return err
	}
	r.watchsets = watchsets

	r.client = mgr.GetClient()
	r.mgr = mgr

	c, err := controller.NewUnmanaged("compositedefinition-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	r.controller = c
	if err := mgr.Add(r); err != nil {
		return err
	}

	// Watch for changes to CompositeDefinition
	err = c.Watch(source.TypedKind(mgr.GetCache(), &addonsv1alpha1.CompositeDefinition{}, &handler.TypedEnqueueRequestForObject[*addonsv1alpha1.CompositeDefinition]{}))
	if err != nil {
		return err
	}

	return nil
}

func (r *CompositeDefinitionReconciler) Start(ctx context.Context) error {
	err := r.controller.Start(ctx)

	// hook to shut down our reconcilers
	for id, reconciler := range r.instanceReconcilers {
		klog.Infof("shutting down controller %v", id)
		reconciler.stop()
		// TODO: Handle/log errors
	}
	return err
}
