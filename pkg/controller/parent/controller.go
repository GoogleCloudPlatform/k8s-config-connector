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

package parent

import (
	"context"
	"fmt"
	"strings"

	bigquerykrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	computekrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	secretkrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	spannerkrm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/kccstate"
	dclcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/ratelimiter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/tf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type Reconcilers struct {
	TF     *tf.Reconciler
	DCL    *dclcontroller.Reconciler
	Direct *directbase.DirectReconciler
}

// ParentReconciler is a top-level controller that decides which underlying
// reconciler (TF, DCL, Direct) should be used for a given resource.
type ParentReconciler struct {
	client.Client
	mgr         manager.Manager
	gvk         schema.GroupVersionKind
	reconcilers Reconcilers
}

func Add(mgr manager.Manager, gvk schema.GroupVersionKind, tf *tf.Reconciler, dcl *dclcontroller.Reconciler, direct *directbase.DirectReconciler) error {
	controllerName := fmt.Sprintf("%v-parent-controller", strings.ToLower(gvk.Kind))
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(gvk)
	predicates := []predicate.Predicate{kccpredicate.UnderlyingResourceOutOfSyncPredicate{}}
	immediateReconcileRequests := make(chan event.GenericEvent, k8s.ImmediateReconcileRequestsBufferSize)

	r := &ParentReconciler{
		Client: mgr.GetClient(),
		mgr:    mgr,
		gvk:    gvk,
		reconcilers: Reconcilers{
			TF:     tf,
			DCL:    dcl,
			Direct: direct,
		},
	}

	_, err := builder.
		ControllerManagedBy(mgr).
		Named(controllerName).
		For(obj, builder.OnlyMetadata, builder.WithPredicates(predicates...)).
		WatchesRawSource(source.TypedChannel(immediateReconcileRequests, &handler.EnqueueRequestForObject{})).
		WithOptions(controller.Options{MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles, RateLimiter: ratelimiter.NewRateLimiter()}).
		Build(r)
	if err != nil {
		return fmt.Errorf("error creating new parent controller: %w", err)
	}
	return nil
}

func (r *ParentReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("parent reconciler starting", "resource", req.NamespacedName)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(r.gvk)
	if err := r.Get(ctx, req.NamespacedName, u); err != nil {
		return reconcile.Result{}, client.IgnoreNotFound(err)
	}

	controllerType, err := r.determineControllerType(ctx, u)
	if err != nil {
		return reconcile.Result{}, err
	}
	resourceControllerConfig := resourceconfig.LoadConfig()
	if !resourceControllerConfig.IsControllerSupported(r.gvk, controllerType) {
		logger.Info("controller type not supported for this resource", "resource", req.NamespacedName, "type", controllerType)
		config, err := resourceControllerConfig.GetControllersForGVK(r.gvk)
		if err != nil {
			return reconcile.Result{}, fmt.Errorf("error getting controller config found for GroupKind %v", r.gvk.GroupKind())
		}
		logger.Info("supported controller types", "resource", req.NamespacedName, "supportedControllers", config.SupportedControllers)

		// Try to write a status on the CCC
		ccc := &corekccv1alpha1.ConfigConnectorContext{}
		cccNamespacedName := types.NamespacedName{
			Namespace: req.Namespace,
			Name:      "configconnectorcontext",
		}
		if err := r.Get(ctx, cccNamespacedName, ccc); err != nil {
			if !apierrors.IsNotFound(err) {
				logger.Error(err, "error getting ConfigConnectorContext, cannot write status", "resource", req.NamespacedName)
			}
		} else {
			msg := fmt.Sprintf("controller type %q is not supported for resource %q. Supported types are: %v. Falling back to default %q.",
				controllerType, r.gvk.GroupKind().String(), config.SupportedControllers, config.DefaultController)

			cccToUpdate := ccc.DeepCopy()
			cccToUpdate.Status.Healthy = false
			cccToUpdate.Status.Errors = append(cccToUpdate.Status.Errors, msg)

			if err := r.Status().Update(ctx, cccToUpdate); err != nil {
				logger.Error(err, "error updating ConfigConnectorContext status", "resource", req.NamespacedName)
			}
		}

		controllerType = config.DefaultController
		logger.Info("falling back to default controller type", "resource", req.NamespacedName, "type", controllerType)
	}

	logger.Info("routing to controller", "type", controllerType)

	switch controllerType {
	case k8s.ReconcilerTypeTerraform:
		logger.Info("routing to TF reconciler")
		if r.reconcilers.TF == nil {
			return reconcile.Result{}, fmt.Errorf("TF reconciler is not initialized for resource %v", r.gvk)
		}
		return r.reconcilers.TF.Reconcile(ctx, req)
	case k8s.ReconcilerTypeDCL:
		logger.Info("routing to DCL reconciler")
		if r.reconcilers.DCL == nil {
			return reconcile.Result{}, fmt.Errorf("DCL reconciler is not initialized for resource %v", r.gvk)
		}
		return r.reconcilers.DCL.Reconcile(ctx, req)
	case k8s.ReconcilerTypeDirect:
		logger.Info("routing to Direct reconciler")
		if r.reconcilers.Direct == nil {
			return reconcile.Result{}, fmt.Errorf("direct reconciler is not initialized for resource %v", r.gvk)
		}
		return r.reconcilers.Direct.Reconcile(ctx, req)
	default:
		return reconcile.Result{}, fmt.Errorf("unknown controller type: %v", controllerType)
	}
}

func (r *ParentReconciler) determineControllerType(ctx context.Context, u *unstructured.Unstructured) (k8s.ReconcilerType, error) {
	// Check for resource annotation
	annotations := u.GetAnnotations()
	if annotations[k8s.AlphaReconcilerAnnotation] == "direct" {
		return k8s.ReconcilerTypeDirect, nil
	}

	// Special case handling. Will be removed after the resources have turned on direct as default.
	if r.gvk.Kind == "BigQueryTable" {
		obj := &bigquerykrm.BigQueryTable{}
		if _, ok := annotations[kccpredicate.AnnotationUnmanaged]; ok {
			return k8s.ReconcilerTypeDirect, nil
		}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
			return "", fmt.Errorf("error converting to %T: %w", obj, err)
		}
		if obj.Spec.Labels != nil {
			return k8s.ReconcilerTypeDirect, nil
		}
	}
	if r.gvk.Kind == "ComputeForwardingRule" {
		obj := &computekrm.ComputeForwardingRule{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
			return "", fmt.Errorf("error converting to %T: %w", obj, err)
		}
		if obj.Spec.Target != nil && obj.Spec.Target.GoogleAPIsBundle != nil {
			return k8s.ReconcilerTypeDirect, nil
		}
	}
	if r.gvk.Kind == "ComputeTargetTCPProxy" {
		obj := &computekrm.ComputeTargetTCPProxy{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
			return "", fmt.Errorf("error converting to %T: %w", obj, err)
		}
		if obj.Spec.Location != nil && obj.Spec.Location != direct.PtrTo("global") {
			return k8s.ReconcilerTypeDirect, nil
		}
	}
	if r.gvk.Kind == "SecretManagerSecret" {
		obj := &secretkrm.SecretManagerSecret{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
			return "", fmt.Errorf("error converting to %T: %w", obj, err)
		}
		if obj.Spec.Labels != nil {
			return k8s.ReconcilerTypeDirect, nil
		}
	}
	if r.gvk.Kind == "SpannerInstance" {
		obj := &spannerkrm.SpannerInstance{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
			return "", fmt.Errorf("error converting to %T: %w", obj, err)
		}
		if obj.Spec.DefaultBackupScheduleType != nil || obj.Spec.Labels != nil || obj.Spec.Edition != nil || obj.Spec.AutoscalingConfig != nil {
			return k8s.ReconcilerTypeDirect, nil
		}
	}

	// Check for CCC setting
	_, ccc, err := kccstate.FetchLiveKCCState(ctx, r.Client, types.NamespacedName{Namespace: u.GetNamespace(), Name: u.GetName()})
	if err != nil {
		return "", fmt.Errorf("error fetching kcc state: %w", err)
	}
	if ccc.Spec.Experiments != nil {
		for k, v := range ccc.Spec.Experiments.ControllerOverrides {
			if k == r.gvk.GroupKind().String() {
				return v, nil
			}
		}
	}

	// Fallback to static config
	gk := r.gvk.GroupKind()
	resourcesControllersConfig := resourceconfig.LoadConfig()
	config, err := resourcesControllersConfig.GetControllersForGVK(r.gvk)
	if err != nil {
		return "", fmt.Errorf("error getting controller config found for GroupKind %v", gk)
	}
	return config.DefaultController, nil
}
