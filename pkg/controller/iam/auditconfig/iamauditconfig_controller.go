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

package auditconfig

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	condition "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/metrics"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/ratelimiter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/execution"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	klog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const controllerName = "iamauditconfig-controller"

var logger = klog.Log.WithName(controllerName)

func Add(mgr manager.Manager, provider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader,
	converter *conversion.Converter, dclConfig *mmdcl.Config) error {
	reconciler, err := NewReconciler(mgr, provider, smLoader, converter, dclConfig)
	if err != nil {
		return err
	}
	return add(mgr, reconciler)
}

func NewReconciler(mgr manager.Manager, provider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader,
	converter *conversion.Converter, dclConfig *mmdcl.Config) (*Reconciler, error) {
	r := Reconciler{
		LifecycleHandler: lifecyclehandler.LifecycleHandler{
			Client:   mgr.GetClient(),
			Recorder: mgr.GetEventRecorderFor(controllerName),
		},
		Client:    mgr.GetClient(),
		iamClient: kcciamclient.New(provider, smLoader, mgr.GetClient(), converter, dclConfig).TFIAMClient,
		scheme:    mgr.GetScheme(),
	}
	return &r, nil
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler.
func add(mgr manager.Manager, r *Reconciler) error {
	obj := &v1beta1.IAMAuditConfig{}
	_, err := builder.
		ControllerManagedBy(mgr).
		Named(controllerName).
		WithOptions(controller.Options{MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles, RateLimiter: ratelimiter.NewRateLimiter()}).
		Watches(&source.Kind{Type: obj}, &handler.EnqueueRequestForObject{}, builder.OnlyMetadata).
		For(obj, builder.OnlyMetadata).
		Build(r)
	if err != nil {
		return fmt.Errorf("error creating new controller: %v", err)
	}
	return nil
}

var _ reconcile.Reconciler = &Reconciler{}

type Reconciler struct {
	lifecyclehandler.LifecycleHandler
	client.Client
	metrics.ReconcilerMetrics
	iamClient *kcciamclient.TFIAMClient
	scheme    *runtime.Scheme
}

type reconcileContext struct {
	Reconciler     *Reconciler
	Ctx            context.Context
	NamespacedName types.NamespacedName
}

func (r *Reconciler) Reconcile(ctx context.Context, request reconcile.Request) (result reconcile.Result, err error) {
	logger.Info("Starting reconcile", "resource", request.NamespacedName)
	startTime := time.Now()
	ctx, cancel := context.WithTimeout(ctx, k8s.ReconcileDeadline)
	defer cancel()
	r.RecordReconcileWorkers(ctx, v1beta1.IAMAuditConfigGVK)
	defer r.AfterReconcile()
	defer r.RecordReconcileMetrics(ctx, v1beta1.IAMAuditConfigGVK, request.Namespace, request.Name, startTime, &err)

	var auditConfig v1beta1.IAMAuditConfig
	if err := r.Get(context.TODO(), request.NamespacedName, &auditConfig); err != nil {
		if apierrors.IsNotFound(err) {
			logger.Info("resource not found in API server; finishing reconcile", "resource", request.NamespacedName)
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}
	reconcileContext := &reconcileContext{
		Reconciler:     r,
		Ctx:            ctx,
		NamespacedName: request.NamespacedName,
	}
	requeue, err := reconcileContext.doReconcile(&auditConfig)
	if err != nil {
		return reconcile.Result{}, err
	}
	if requeue {
		return reconcile.Result{Requeue: true}, nil
	}
	jitteredPeriod := jitter.GenerateJitteredReenqueuePeriod()
	logger.Info("successfully finished reconcile", "resource", request.NamespacedName, "time to next reconciliation", jitteredPeriod)
	return reconcile.Result{RequeueAfter: jitteredPeriod}, nil
}

func (r *reconcileContext) doReconcile(auditConfig *v1beta1.IAMAuditConfig) (requeue bool, err error) {
	defer execution.RecoverWithInternalError(&err)
	if !auditConfig.DeletionTimestamp.IsZero() {
		if !k8s.HasFinalizer(auditConfig, k8s.ControllerFinalizerName) {
			// Resource has no controller finalizer; no finalization necessary
			return false, nil
		}
		if k8s.HasFinalizer(auditConfig, k8s.DeletionDefenderFinalizerName) {
			// Deletion defender has not yet been finalized; requeuing
			logger.Info("deletion defender has not yet been finalized; requeuing", "resource", k8s.GetNamespacedName(auditConfig))
			return true, nil
		}
		if !k8s.HasAbandonAnnotation(auditConfig) {
			if err := r.Reconciler.iamClient.DeleteAuditConfig(r.Ctx, auditConfig); err != nil {
				if !errors.Is(err, kcciamclient.NotFoundError) && !k8s.IsReferenceNotFoundError(err) {
					if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
						logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(auditConfig))
						return true, r.handleUnresolvableDeps(auditConfig, unwrappedErr)
					}
					return false, r.handleDeleteFailed(auditConfig, err)
				}
			}
		}
		return false, r.handleDeleted(auditConfig)
	}
	if _, err := r.Reconciler.iamClient.GetAuditConfig(r.Ctx, auditConfig); err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(auditConfig))
			return true, r.handleUnresolvableDeps(auditConfig, unwrappedErr)
		}
		if !errors.Is(err, kcciamclient.NotFoundError) {
			return false, r.handleUpdateFailed(auditConfig, err)
		}
	}
	if !k8s.EnsureFinalizers(auditConfig, k8s.ControllerFinalizerName, k8s.DeletionDefenderFinalizerName) {
		if err := r.update(auditConfig); err != nil {
			return false, r.handleUpdateFailed(auditConfig, err)
		}
	}
	if _, err := r.Reconciler.iamClient.SetAuditConfig(r.Ctx, auditConfig); err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(auditConfig))
			return true, r.handleUnresolvableDeps(auditConfig, unwrappedErr)
		}
		return false, r.handleUpdateFailed(auditConfig, fmt.Errorf("error setting audit config: %w", err))
	}
	if isAPIServerUpdateRequired(auditConfig) {
		return false, r.handleUpToDate(auditConfig)
	}
	return false, nil
}

func (r *reconcileContext) update(auditConfig *v1beta1.IAMAuditConfig) error {
	if err := r.Reconciler.Client.Update(r.Ctx, auditConfig); err != nil {
		return fmt.Errorf("error updating '%v' in API server: %w", r.NamespacedName, err)
	}
	return nil
}

func (r *reconcileContext) handleUpToDate(auditConfig *v1beta1.IAMAuditConfig) error {
	resource, err := toK8sResource(auditConfig)
	if err != nil {
		return fmt.Errorf("error converting IAMAuditConfig to k8s resource while handling %v event: %w", k8s.UpToDate, err)
	}
	return r.Reconciler.HandleUpToDate(r.Ctx, resource)
}

func (r *reconcileContext) handleUpdateFailed(auditConfig *v1beta1.IAMAuditConfig, origErr error) error {
	resource, err := toK8sResource(auditConfig)
	if err != nil {
		logger.Error(err, "error converting IAMAuditConfig to k8s resource while handling event",
			"resource", k8s.GetNamespacedName(auditConfig), "event", k8s.UpdateFailed)
		return fmt.Errorf(k8s.UpdateFailedMessageTmpl, origErr)
	}
	return r.Reconciler.HandleUpdateFailed(r.Ctx, resource, origErr)
}

func (r *reconcileContext) handleDeleted(auditConfig *v1beta1.IAMAuditConfig) error {
	resource, err := toK8sResource(auditConfig)
	if err != nil {
		return fmt.Errorf("error converting IAMAuditConfig to k8s resource while handling %v event: %w", k8s.Deleted, err)
	}
	return r.Reconciler.HandleDeleted(r.Ctx, resource)
}

func (r *reconcileContext) handleDeleteFailed(auditConfig *v1beta1.IAMAuditConfig, origErr error) error {
	resource, err := toK8sResource(auditConfig)
	if err != nil {
		logger.Error(err, "error converting IAMAuditConfig to k8s resource while handling event",
			"resource", k8s.GetNamespacedName(auditConfig), "event", k8s.DeleteFailed)
		return fmt.Errorf(k8s.DeleteFailedMessageTmpl, origErr)
	}
	return r.Reconciler.HandleDeleteFailed(r.Ctx, resource, origErr)
}

func (r *reconcileContext) handleUnresolvableDeps(auditConfig *v1beta1.IAMAuditConfig, origErr error) error {
	resource, err := toK8sResource(auditConfig)
	if err != nil {
		return fmt.Errorf("error converting IAMAuditConfig to k8s resource while handling unresolvable dependencies event: %w", err)
	}
	return r.Reconciler.HandleUnresolvableDeps(r.Ctx, resource, origErr)
}

func isAPIServerUpdateRequired(auditConfig *v1beta1.IAMAuditConfig) bool {
	// TODO: even in the event of an actual update to GCP, this function will
	// return false because the condition comparison doesn't account for time.
	conditions := []condition.Condition{
		k8s.NewCustomReadyCondition(corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage),
	}
	if !k8s.ConditionSlicesEqual(auditConfig.Status.Conditions, conditions) {
		return true
	}
	if auditConfig.Status.ObservedGeneration != auditConfig.GetGeneration() {
		return true
	}
	return false
}

func toK8sResource(auditConfig *v1beta1.IAMAuditConfig) (*k8s.Resource, error) {
	kcciamclient.SetGVK(auditConfig)
	resource := k8s.Resource{}
	if err := util.Marshal(auditConfig, &resource); err != nil {
		return nil, fmt.Errorf("error marshalling IAMAuditConfig to k8s resource: %w", err)
	}
	return &resource, nil
}
