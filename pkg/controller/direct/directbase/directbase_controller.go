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

package directbase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/metrics"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/ratelimiter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourcewatcher"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/execution"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/lease/leaser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	"golang.org/x/sync/semaphore"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	crcontroller "sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// Add creates a new IAM Policy Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and start it when the Manager is started.
func Add(mgr manager.Manager, gvk schema.GroupVersionKind, model Model) error {
	immediateReconcileRequests := make(chan event.GenericEvent, k8s.ImmediateReconcileRequestsBufferSize)
	resourceWatcherRoutines := semaphore.NewWeighted(k8s.MaxNumResourceWatcherRoutines)
	reconciler, err := NewReconciler(mgr, immediateReconcileRequests, resourceWatcherRoutines, gvk, model)
	if err != nil {
		return err
	}
	return add(mgr, reconciler)
}

// NewReconciler returns a new reconcile.Reconciler.
func NewReconciler(mgr manager.Manager, immediateReconcileRequests chan event.GenericEvent, resourceWatcherRoutines *semaphore.Weighted,
	gvk schema.GroupVersionKind, model Model) (*DirectReconciler, error) {

	controllerName := strings.ToLower(gvk.Kind) + "-controller"

	// var logger = log.Log.WithName(controllerName)

	r := DirectReconciler{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(
			mgr.GetClient(),
			mgr.GetEventRecorderFor(controllerName),
		),
		Client: mgr.GetClient(),
		//iamClient:                  iamclient.New(provider, smLoader, mgr.GetClient(), converter, dclConfig),
		config:                     mgr.GetConfig(),
		immediateReconcileRequests: immediateReconcileRequests,
		resourceWatcherRoutines:    resourceWatcherRoutines,
		scheme:                     mgr.GetScheme(),
		gvk:                        gvk,
		model:                      model,
		controllerName:             controllerName,
		ReconcilerMetrics: metrics.ReconcilerMetrics{
			ResourceNameLabel: metrics.ResourceNameLabel,
		},
	}
	return &r, nil
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler.
func add(mgr manager.Manager, r *DirectReconciler) error {
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(r.gvk)

	_, err := builder.
		ControllerManagedBy(mgr).
		Named(r.controllerName).
		WithOptions(crcontroller.Options{MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles, RateLimiter: ratelimiter.NewRateLimiter()}).
		WatchesRawSource(&source.Channel{Source: r.immediateReconcileRequests}, &handler.EnqueueRequestForObject{}).
		For(obj, builder.OnlyMetadata, builder.WithPredicates(predicate.UnderlyingResourceOutOfSyncPredicate{})).
		Build(r)
	if err != nil {
		return fmt.Errorf("error creating new controller: %v", err)
	}
	return nil
}

var _ reconcile.Reconciler = &DirectReconciler{}

// DirectReconciler is a reconciler for handling IAM policies.
type DirectReconciler struct {
	lifecyclehandler.LifecycleHandler
	client.Client
	metrics.ReconcilerMetrics
	gvk    schema.GroupVersionKind
	model  Model
	scheme *runtime.Scheme
	config *rest.Config
	// Fields used for triggering reconciliations when dependencies are ready
	immediateReconcileRequests chan event.GenericEvent
	resourceWatcherRoutines    *semaphore.Weighted // Used to cap number of goroutines watching unready dependencies

	controllerName string
}

func (r *DirectReconciler) HandleUpToDate(ctx context.Context, obj *unstructured.Unstructured) error {
	setCondition(obj, corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage)
	if err := r.updateAPIServer(ctx, obj); err != nil {
		return err
	}

	r.recordEvent(obj, corev1.EventTypeNormal, k8s.UpToDate, k8s.UpToDateMessage)
	return nil
}

func (r *DirectReconciler) HandleUnresolvableDeps(ctx context.Context, obj *unstructured.Unstructured, originErr error) error {
	reason, err := lifecyclehandler.ReasonForUnresolvableDeps(originErr)
	if err != nil {
		return r.HandleUpdateFailed(ctx, obj, err)
	}
	msg := originErr.Error()
	// Only update the API server if there's new information
	if !k8s.UnstructuredReadyConditionMatches(obj, corev1.ConditionFalse, reason, msg) {
		setCondition(obj, corev1.ConditionFalse, reason, msg)
		setObservedGeneration(obj, obj.GetGeneration())
		if err := r.updateStatus(ctx, obj); err != nil {
			return err
		}
	}

	r.recordEvent(obj, corev1.EventTypeWarning, reason, msg)
	return nil
}

func (r *DirectReconciler) HandleUpdateFailed(ctx context.Context, obj *unstructured.Unstructured, err error) error {
	msg := fmt.Errorf("Update call failed: %w", err).Error()
	setCondition(obj, corev1.ConditionFalse, k8s.UpdateFailed, msg)
	setObservedGeneration(obj, obj.GetGeneration())
	if err := r.updateStatus(ctx, obj); err != nil {
		return err
	}

	r.recordEvent(obj, corev1.EventTypeWarning, k8s.UpdateFailed, msg)
	return fmt.Errorf("Update call failed: %w", err)
}

func (r *DirectReconciler) HandleDeleted(ctx context.Context, obj *unstructured.Unstructured) error {
	setCondition(obj, corev1.ConditionFalse, k8s.Deleted, k8s.DeletedMessage)
	setObservedGeneration(obj, obj.GetGeneration())
	// Do an explicit status update first to prevent a race between the status update and the API
	// server pruning the resource if there are no more finalizers present.
	if err := r.updateStatus(ctx, obj); err != nil {
		return fmt.Errorf("error updating status: %w", err)
	}

	r.recordEvent(obj, corev1.EventTypeNormal, k8s.Deleted, k8s.DeletedMessage)

	k8s.RemoveFinalizer(obj, k8s.ControllerFinalizerName)
	return r.updateAPIServer(ctx, obj)
}

func (r *DirectReconciler) HandleDeleteFailed(ctx context.Context, obj *unstructured.Unstructured, err error) error {
	msg := fmt.Sprintf(k8s.DeleteFailedMessageTmpl, err)
	setCondition(obj, corev1.ConditionFalse, k8s.DeleteFailed, msg)
	setObservedGeneration(obj, obj.GetGeneration())
	if err := r.updateStatus(ctx, obj); err != nil {
		return err
	}

	r.recordEvent(obj, corev1.EventTypeWarning, k8s.DeleteFailed, msg)
	return fmt.Errorf("Delete call failed: %w", err)
}

func (r *DirectReconciler) recordEvent(obj *unstructured.Unstructured, eventtype, reason, message string) {
	r.Recorder.Event(obj, eventtype, reason, message)
}

// The system sets various labels on the resource that are not user facing and should not be saved in the API server
// this function removes any that may be present
func removeSystemLabels(u client.Object) {
	labels := u.GetLabels()
	if labels == nil {
		return
	}
	keys := leaser.GetLabelKeys()
	keys = append(keys, label.CnrmManagedKey)
	for _, k := range keys {
		delete(labels, k)
	}
	// GetLabels(...) returns a new copy of the labels map so we must overwrite that value with our local value
	u.SetLabels(labels)
}

// WARNING: This function should NOT be exported and invoked directly outside the package.
// Controllers are supposed to call exported functions to handle lifecycle transitions.
func (r *DirectReconciler) updateAPIServer(ctx context.Context, u *unstructured.Unstructured) error {
	// Preserve the intended status, as the client.Update call will ignore the given status
	// and return the stale existing status.
	originalStatus := u.Object["status"]

	// Get the current generation as the observed generation because the following client.Update
	// might increase the generation. We want the next reconciliation to handle the new generation.
	observedGeneration := u.GetGeneration()
	// u, err := resource.MarshalAsUnstructured()
	// if err != nil {
	// 	return err
	// }
	removeSystemLabels(u)

	// TODO: Only if spec/label/annotation changes

	// spec := obj.DeepCopyObject().(T)
	if err := r.Client.Update(ctx, u, r.GetFieldOwner()); err != nil {
		if apierrors.IsConflict(err) {
			return fmt.Errorf("couldn't update the API server due to conflict. Re-enqueue the request for another reconciliation attempt: %w", err)
		}
		return fmt.Errorf("error with update call to API server: %w", err)
	}
	// Restore the status, it was likely removed in the above update
	u.Object["status"] = originalStatus

	setObservedGeneration(u, observedGeneration)

	// obj.SetResourceVersion(spec.GetResourceVersion())

	shouldUpdateStatus := true
	if !u.GetDeletionTimestamp().IsZero() && len(u.GetFinalizers()) == 0 {
		// This resource is already gone, or is about to be already gone, don't set the status.
		// Status updates for successful deletions must be handled before the finalizer is removed.
		shouldUpdateStatus = false
	}

	if shouldUpdateStatus {
		// copyForStatus.SetResourceVersion(obj.GetResourceVersion())
		setObservedGeneration(u, observedGeneration)

		// Workaround for https://github.com/kubernetes-sigs/controller-runtime/issues/2453
		// status := obj.DeepCopyObject()
		if err := r.updateStatus(ctx, u); err != nil {
			return err
		}
		u.SetResourceVersion(u.GetResourceVersion())

		// TODO: this doesn't look right
		// // rejections by validating webhooks won't be returned as an error; instead, they will be
		// // objects of kind "Status" with a "Failure" status.
		// if isFailureStatus(u) {
		// 	return fmt.Errorf("error with update call to API server: %v", u.Object["message"])
		// }

		// // sync the resource up with the updated metadata
		// if err := util.Marshal(u, resource); err != nil {
		// 	return fmt.Errorf("error syncing updated resource metadata: %w", err)
		// }
	}

	return nil
}

func setObservedGeneration(obj *unstructured.Unstructured, observedGeneration int64) {
	k8s.UnstructuredSetObservedGeneration(obj, observedGeneration)
}

func setCondition(obj *unstructured.Unstructured, status corev1.ConditionStatus, reason, msg string) {
	conditions, _ := k8s.UnstructuredGetConditions(obj)

	newReadyCondition := k8s.NewCustomReadyCondition(status, reason, msg)

	done := false
	for i, condition := range conditions {
		if condition.Type != newReadyCondition.Type {
			continue
		}
		// We should only update the ready condition's last transition time if there was a transition
		// since its last state. The function sets it to time.Now(), so let's replace it if there was
		// no transition.
		if newReadyCondition.Status == condition.Status {
			newReadyCondition.LastTransitionTime = condition.LastTransitionTime
		}
		conditions[i] = newReadyCondition
		done = true
	}
	if !done {
		conditions = append(conditions, newReadyCondition)
	}
	k8s.UnstructuredSetConditions(obj, conditions)
}

func (r *DirectReconciler) updateStatus(ctx context.Context, obj *unstructured.Unstructured) error {
	// u, err := resource.MarshalAsUnstructured()
	// if err != nil {
	// 	return err
	// }

	// TODO: Do we need to workaround https://github.com/kubernetes-sigs/controller-runtime/issues/2453
	// We could just copy the spec etc

	// Workaround for https://github.com/kubernetes-sigs/controller-runtime/issues/2453
	statusCopy := &unstructured.Unstructured{Object: make(map[string]interface{})}
	for k, v := range obj.Object {
		statusCopy.Object[k] = v
	}

	if err := r.Client.Status().Update(ctx, statusCopy, r.GetFieldOwner()); err != nil {
		if apierrors.IsConflict(err) {
			return fmt.Errorf("couldn't update the API server due to conflict. Re-enqueue the request for another reconciliation attempt: %w", err)
		}
		return fmt.Errorf("error with status update call to API server: %w", err)
	}

	obj.Object["status"] = statusCopy.Object["status"]
	obj.SetResourceVersion(statusCopy.GetResourceVersion())

	// // rejections by some validating webhooks won't be returned as an error; instead, they will be
	// // objects of kind "Status" with a "Failure" status.
	// if isFailureStatus(u) {
	// 	return fmt.Errorf("error with status update call to API server: %v", u.Object["message"])
	// }
	// // sync the resource up with the updated metadata.
	// if err := util.Marshal(u, resource); err != nil {
	// 	return err
	// }

	// TODO: No transforms in DirectReconcilers
	// return resourceoverrides.Handler.PostUpdateStatusTransform(resource)
	return nil
}

type reconcileContext struct {
	gvk            schema.GroupVersionKind
	Reconciler     *DirectReconciler
	NamespacedName types.NamespacedName
}

// Reconcile checks k8s for the current state of the resource.
func (r *DirectReconciler) Reconcile(ctx context.Context, request reconcile.Request) (result reconcile.Result, err error) {
	logger := log.FromContext(ctx)

	logger.Info("Running reconcile", "resource", request.NamespacedName)
	startTime := time.Now()
	ctx, cancel := context.WithTimeout(ctx, k8s.ReconcileDeadline)
	defer cancel()
	r.RecordReconcileWorkers(ctx, r.gvk)
	defer r.AfterReconcile()
	defer r.RecordReconcileMetrics(ctx, r.gvk, request.Namespace, request.Name, startTime, &err)

	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(r.gvk)
	if err := r.Get(ctx, request.NamespacedName, obj); err != nil {
		if apierrors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	runCtx := &reconcileContext{
		Reconciler:     r,
		gvk:            r.gvk,
		NamespacedName: request.NamespacedName,
	}
	requeue, err := runCtx.doReconcile(ctx, obj)
	if err != nil {
		return reconcile.Result{}, err
	}
	if requeue {
		return reconcile.Result{Requeue: true}, nil
	}
	jitteredPeriod, err := jitter.GenerateJitteredReenqueuePeriod(r.gvk, nil, nil, obj)
	if err != nil {
		return reconcile.Result{}, err
	}
	logger.Info("successfully finished reconcile", "resource", request.NamespacedName, "time to next reconciliation", jitteredPeriod)
	return reconcile.Result{RequeueAfter: jitteredPeriod}, nil
}

func (r *reconcileContext) doReconcile(ctx context.Context, u *unstructured.Unstructured) (requeue bool, err error) {
	logger := log.FromContext(ctx)

	logger.Info("RECONCILE", "obj", u)

	adapter, err := r.Reconciler.model.AdapterForObject(ctx, u)
	if err != nil {
		return false, r.handleUpdateFailed(ctx, u, err)
	}

	defer execution.RecoverWithInternalError(&err)
	if !u.GetDeletionTimestamp().IsZero() {
		if !k8s.HasFinalizer(u, k8s.ControllerFinalizerName) {
			// Resource has no controller finalizer; no finalization necessary
			return false, nil
		}
		if k8s.HasFinalizer(u, k8s.DeletionDefenderFinalizerName) {
			// deletion defender has not yet been finalized; requeuing
			logger.Info("deletion defender has not yet been finalized; requeuing", "resource", k8s.GetNamespacedName(u))
			return true, nil
		}
		if !k8s.HasAbandonAnnotation(u) {
			if _, err := adapter.Delete(ctx); err != nil {
				if !errors.Is(err, kcciamclient.NotFoundError) && !k8s.IsReferenceNotFoundError(err) {
					if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
						logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(u))
						// resource, err := toK8sResource(u)
						// if err != nil {
						// 	return false, fmt.Errorf("error converting k8s resource while handling unresolvable dependencies event: %w", err)
						// }
						// Requeue resource for reconciliation with exponential backoff applied
						return true, r.Reconciler.HandleUnresolvableDeps(ctx, u, unwrappedErr)
					}
					return false, r.handleDeleteFailed(ctx, u, err)
				}
			}
		}
		return false, r.handleDeleted(ctx, u)
	}

	if err := r.ensureFinalizers(ctx, u); err != nil {
		return false, err
	}

	existsAlready, err := adapter.Find(ctx)
	if err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(u))
			return r.handleUnresolvableDeps(ctx, u, unwrappedErr)
		}
		return false, r.handleUpdateFailed(ctx, u, err)
	}

	// set the etag to an empty string, since IAMPolicy is the authoritative intent, KCC wants to overwrite the underlying policy regardless
	//policy.Spec.Etag = ""

	if !existsAlready {
		if err := adapter.Create(ctx, u); err != nil {
			if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
				logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(u))
				return r.handleUnresolvableDeps(ctx, u, unwrappedErr)
			}
			return false, r.handleUpdateFailed(ctx, u, fmt.Errorf("error creating: %w", err))
		}
	} else {
		if _, err = adapter.Update(ctx); err != nil {
			if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
				logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(u))
				return r.handleUnresolvableDeps(ctx, u, unwrappedErr)
			}
			return false, r.handleUpdateFailed(ctx, u, fmt.Errorf("error updating: %w", err))
		}
	}
	if isAPIServerUpdateRequired(u) {
		return false, r.handleUpToDate(ctx, u)
	}
	return false, nil
}

func (r *reconcileContext) ensureFinalizers(ctx context.Context, obj *unstructured.Unstructured) error {
	updated := false
	updated = updated || controllerutil.AddFinalizer(obj, k8s.ControllerFinalizerName)
	updated = updated || controllerutil.AddFinalizer(obj, k8s.DeletionDefenderFinalizerName)

	if updated {
		if err := r.Reconciler.Client.Update(ctx, obj); err != nil {
			return fmt.Errorf("updating finalizers: %w", err)
		}
	}

	return nil
}

func (r *reconcileContext) handleUpToDate(ctx context.Context, obj *unstructured.Unstructured) error {
	// resource, err := toK8sResource(u)
	// if err != nil {
	// 	return fmt.Errorf("error converting to k8s resource while handling %v event: %w", k8s.UpToDate, err)
	// }
	return r.Reconciler.HandleUpToDate(ctx, obj)
}

func (r *reconcileContext) handleUpdateFailed(ctx context.Context, u *unstructured.Unstructured, origErr error) error {
	// logger := log.FromContext(ctx)

	// resource, err := toK8sResource(obj)
	// if err != nil {
	// 	logger.Error(err, "error converting to k8s resource while handling event",
	// 		"resource", k8s.GetNamespacedName(obj), "event", k8s.UpdateFailed)
	// 	return fmt.Errorf("Update call failed: %w", origErr)
	// }
	return r.Reconciler.HandleUpdateFailed(ctx, u, origErr)
}

func (r *reconcileContext) handleDeleted(ctx context.Context, u *unstructured.Unstructured) error {
	// resource, err := toK8sResource(policy)
	// if err != nil {
	// 	return fmt.Errorf("error converting to k8s resource while handling %v event: %w", k8s.Deleted, err)
	// }
	return r.Reconciler.HandleDeleted(ctx, u)
}

func (r *reconcileContext) handleDeleteFailed(ctx context.Context, u *unstructured.Unstructured, origErr error) error {
	// logger := log.FromContext(ctx)

	// resource, err := toK8sResource(policy)
	// if err != nil {
	// 	logger.Error(err, "error converting to k8s resource while handling event",
	// 		"resource", k8s.GetNamespacedName(policy), "event", k8s.DeleteFailed)
	// 	return fmt.Errorf(k8s.DeleteFailedMessageTmpl, origErr)
	// }
	return r.Reconciler.HandleDeleteFailed(ctx, u, origErr)
}

func (r *DirectReconciler) supportsImmediateReconciliations() bool {
	return r.immediateReconcileRequests != nil
}

func (r *reconcileContext) handleUnresolvableDeps(ctx context.Context, u *unstructured.Unstructured, origErr error) (requeue bool, err error) {
	logger := log.FromContext(ctx)

	// resource, err := toK8sResource(policy)
	// if err != nil {
	// 	return false, fmt.Errorf("error converting to k8s resource while handling unresolvable dependencies event: %w", err)
	// }
	refGVK, refNN, ok := lifecyclehandler.CausedByUnreadyOrNonexistentResourceRefs(origErr)
	if !ok || !r.Reconciler.supportsImmediateReconciliations() {
		// Requeue resource for reconciliation with exponential backoff applied
		return true, r.Reconciler.HandleUnresolvableDeps(ctx, u, origErr)
	}
	// Check that the number of active resource watches
	// does not exceed the controller's cap. If the
	// capacity is not exceeded, The number of active
	// resource watches is incremented by one and a watch
	// is started
	if !r.Reconciler.resourceWatcherRoutines.TryAcquire(1) {
		// Requeue resource for reconciliation with exponential backoff applied
		return true, r.Reconciler.HandleUnresolvableDeps(ctx, u, origErr)
	}
	// Create a logger for ResourceWatcher that contains info
	// about the referencing resource. This is done since the
	// messages logged by ResourceWatcher only include the
	// information of the resource it is watching by default.
	watcherLogger := logger.WithValues(
		"referencingResource", k8s.GetNamespacedName(u),
		"referencingResourceGVK", u.GroupVersionKind())
	watcher, err := resourcewatcher.New(r.Reconciler.config, watcherLogger)
	if err != nil {
		return false, r.Reconciler.HandleUpdateFailed(ctx, u, fmt.Errorf("error initializing new resourcewatcher: %w", err))
	}

	logger = logger.WithValues(
		"resource", k8s.GetNamespacedName(u),
		"resourceGVK", u.GroupVersionKind(),
		"reference", refNN,
		"referenceGVK", refGVK)
	go func() {
		// Decrement the count of active resource watches after
		// the watch finishes
		defer r.Reconciler.resourceWatcherRoutines.Release(1)
		timeoutPeriod := jitter.GenerateWatchJitteredTimeoutPeriod()
		ctx, cancel := context.WithTimeout(context.TODO(), timeoutPeriod)
		defer cancel()
		logger.Info("starting wait with timeout on resource's reference", "timeout", timeoutPeriod)
		if err := watcher.WaitForResourceToBeReady(ctx, refNN, refGVK); err != nil {
			logger.Error(err, "error while waiting for resource's reference to be ready")
			return
		}
		logger.Info("enqueuing resource for immediate reconciliation now that its reference is ready")
		r.Reconciler.enqueueForImmediateReconciliation(k8s.GetNamespacedName(u))
	}()

	// Do not requeue resource for immediate reconciliation. Wait for either
	// the next periodic reconciliation or for the referenced resource to be ready (which
	// triggers a reconciliation), whichever comes first.
	return false, r.Reconciler.HandleUnresolvableDeps(ctx, u, origErr)
}

// enqueueForImmediateReconciliation enqueues the given resource for immediate
// reconciliation. Note that this function only takes in the name and namespace
// of the resource and not its GVK since the controller instance that this
// reconcile instance belongs to can only reconcile resources of one GVK.
func (r *DirectReconciler) enqueueForImmediateReconciliation(resourceNN types.NamespacedName) {
	genEvent := event.GenericEvent{}
	genEvent.Object = &unstructured.Unstructured{}
	genEvent.Object.SetNamespace(resourceNN.Namespace)
	genEvent.Object.SetName(resourceNN.Name)
	r.immediateReconcileRequests <- genEvent
}

func isAPIServerUpdateRequired(policy *unstructured.Unstructured) bool {
	// TODO: even in the event of an actual update to GCP, this function will
	// return false because the condition comparison doesn't account for time.
	// conditions := []condition.Condition{
	// 	k8s.NewCustomReadyCondition(corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage),
	// }

	// TODO: Implement these checks
	return true
	// if !k8s.ConditionSlicesEqual(policy.Status.Conditions, conditions) {
	// 	return true
	// }
	// if policy.Status.ObservedGeneration != policy.GetGeneration() {
	// 	return true
	// }
	// return false
}

func toK8sResource(policy *unstructured.Unstructured) (*k8s.Resource, error) {
	resource := k8s.Resource{}
	if err := util.Marshal(policy, &resource); err != nil {
		return nil, fmt.Errorf("error marshalling to k8s resource: %w", err)
	}
	return &resource, nil
}
