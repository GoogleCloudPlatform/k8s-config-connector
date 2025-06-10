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

package deletiondefender

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/go-logr/logr"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var logger = log.Log

type Reconciler struct {
	client.Client
	// a core k8s client set, this is used for CRD GET requests. This client is used because it will not have caches with
	// any configuration.
	clientSet *clientset.Clientset
	mgr       manager.Manager
	crd       *apiextensions.CustomResourceDefinition
	gvk       schema.GroupVersionKind
	logger    logr.Logger
}

func Add(mgr manager.Manager, crd *apiextensions.CustomResourceDefinition) error {
	kind := crd.Spec.Names.Kind
	apiVersion := k8s.GetAPIVersionFromCRD(crd)
	controllerName := fmt.Sprintf("%v-deletion-defender-controller", strings.ToLower(kind))
	r, err := NewReconciler(mgr, crd)
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
		For(obj, builder.OnlyMetadata).
		Build(r)
	if err != nil {
		return fmt.Errorf("error creating new controller: %w", err)
	}
	log := mgr.GetLogger()
	log.Info("Registered deletion-defender controller", "kind", kind, "apiVersion", apiVersion)
	return nil
}

func NewReconciler(mgr manager.Manager, crd *apiextensions.CustomResourceDefinition) (*Reconciler, error) {
	controllerName := fmt.Sprintf("%v-deletion-defender-controller", strings.ToLower(crd.Spec.Names.Kind))
	clientSet, err := clientset.NewForConfig(mgr.GetConfig())
	if err != nil {
		return nil, fmt.Errorf("error creating new clientset: %w", err)
	}
	return &Reconciler{
		Client:    mgr.GetClient(),
		clientSet: clientSet,
		mgr:       mgr,
		crd:       crd,
		gvk: schema.GroupVersionKind{
			Group:   crd.Spec.Group,
			Version: k8s.GetVersionFromCRD(crd),
			Kind:    crd.Spec.Names.Kind,
		},
		logger: logger.WithName(controllerName),
	}, nil
}

func (r *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (res reconcile.Result, err error) {
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(r.gvk)

	if err := r.Get(ctx, req.NamespacedName, u); err != nil {
		if errors.IsNotFound(err) {
			r.logger.Info("resource not found in API server; finishing reconcile", "resource", req.NamespacedName)
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}
	if u.GetDeletionTimestamp().IsZero() || !k8s.HasFinalizer(u, k8s.DeletionDefenderFinalizerName) {
		return reconcile.Result{}, nil
	}

	// The resource is being deleted, and has the deletion defender finalizer. Determine whether
	// this resource deletion should result in a delete request to the underlying API.
	r.logger.Info("starting deletion defender finalization", "resource", req.NamespacedName)

	uninstalling, err := r.isUninstalling(ctx)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error determining if CRD is uninstalling: %w", err)
	}

	// If we are uninstalling, remove both KCC finalizers and set the resource to abandon. Otherwise,
	// remove just the deletion defender finalizer and allow the controller to delete the underlying
	// resource on GCP.
	k8s.RemoveFinalizer(u, k8s.DeletionDefenderFinalizerName)
	if uninstalling {
		r.logger.Info("resource type is being uninstalled; abandoning by default", "resource", req.NamespacedName)
		k8s.RemoveFinalizer(u, k8s.ControllerFinalizerName)
		k8s.SetAnnotation(k8s.DeletionPolicyAnnotation, k8s.DeletionPolicyAbandon, u)
	}
	if err := r.Update(ctx, u); err != nil {
		if errors.IsConflict(err) {
			return reconcile.Result{}, fmt.Errorf("couldn't update the api server due to conflict. Re-enqueue the request for another reconciliation attempt: %w", err)
		}
		return reconcile.Result{}, fmt.Errorf("error with update call to API server: %w", err)
	}

	r.logger.Info("successfully finalized deletion defense", "resource", req.NamespacedName)
	return reconcile.Result{}, nil
}

func (r *Reconciler) isUninstalling(ctx context.Context) (bool, error) {
	// Check if the associated CRD has its deletion timestamp set.
	// it is important to use the clientset.Clientset here rather than the controller-runtime client.Client, because
	// controller-runtime's client can have caches enabled, disabling them is tricky, and it would be easy for a bug to
	// be introduced that re-enables the cache. We want the latest state of the CRD here so use the basic clientset.
	crd, err := r.clientSet.ApiextensionsV1().CustomResourceDefinitions().Get(ctx, r.crd.GetName(), v1.GetOptions{})
	if err != nil {
		return false, fmt.Errorf("error getting CRD '%v': %w", r.crd.GetName(), err)
	}
	return !crd.GetDeletionTimestamp().IsZero(), nil
}
