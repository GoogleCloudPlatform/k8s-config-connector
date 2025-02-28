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

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/kccstate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/metrics"
	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/ratelimiter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceactuation"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourcewatcher"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/execution"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	"golang.org/x/sync/semaphore"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	crcontroller "sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

func AddController(mgr manager.Manager, gvk schema.GroupVersionKind, model Model, deps Deps) error {
	immediateReconcileRequests := make(chan event.GenericEvent, k8s.ImmediateReconcileRequestsBufferSize)
	resourceWatcherRoutines := semaphore.NewWeighted(k8s.MaxNumResourceWatcherRoutines)

	if model == nil {
		return fmt.Errorf("model is nil for gvk %s", gvk)
	}

	reconciler, err := NewReconciler(mgr, immediateReconcileRequests, resourceWatcherRoutines, gvk, model, deps.JitterGenerator)
	if err != nil {
		return err
	}
	return add(mgr, reconciler, deps.ReconcilePredicate)
}

// NewReconciler returns a new reconcile.Reconciler.
func NewReconciler(mgr manager.Manager, immediateReconcileRequests chan event.GenericEvent, resourceWatcherRoutines *semaphore.Weighted,
	gvk schema.GroupVersionKind, model Model, jg jitter.Generator) (*DirectReconciler, error) {
	controllerName := strings.ToLower(gvk.Kind) + "-controller"
	if jg == nil {
		return nil, fmt.Errorf("jitter generator is not initialized")
	}
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
		jitterGenerator: jg,
	}
	return &r, nil
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler.
func add(mgr manager.Manager, r *DirectReconciler, reconcilePredicate predicate.Predicate) error {
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(r.gvk)

	predicateList := []predicate.Predicate{kccpredicate.UnderlyingResourceOutOfSyncPredicate{}}
	if reconcilePredicate != nil {
		predicateList = append(predicateList, reconcilePredicate)
	}

	_, err := builder.
		ControllerManagedBy(mgr).
		Named(r.controllerName).
		WithOptions(crcontroller.Options{MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles, RateLimiter: ratelimiter.NewRateLimiter()}).
		WatchesRawSource(&source.Channel{Source: r.immediateReconcileRequests}, &handler.EnqueueRequestForObject{}).
		For(obj, builder.OnlyMetadata, builder.WithPredicates(predicateList...)).
		Build(r)
	if err != nil {
		return fmt.Errorf("error creating new controller: %w", err)
	}
	return nil
}

var _ reconcile.Reconciler = &DirectReconciler{}

type Deps struct {
	JitterGenerator    jitter.Generator
	ReconcilePredicate predicate.Predicate
}

// DirectReconciler is a reconciler for reconciling resources that support the Model/Adapter pattern.
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
	jitterGenerator            jitter.Generator

	controllerName string
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

	skip, err := resourceactuation.ShouldSkip(obj)
	if err != nil {
		return reconcile.Result{}, err
	}
	if skip {
		logger.Info("Skipping reconcile as nothing has changed and 0 reconcile period is set", "resource", request.NamespacedName)
		return reconcile.Result{}, nil
	}

	requeue, err := runCtx.doReconcile(ctx, obj)
	if err != nil {
		return reconcile.Result{}, err
	}
	if requeue {
		return reconcile.Result{Requeue: true}, nil
	}
	jitteredPeriod, err := r.jitterGenerator.JitteredReenqueue(r.gvk, obj)
	if err != nil {
		return reconcile.Result{}, err
	}
	logger.Info("successfully finished reconcile", "resource", request.NamespacedName, "time to next reconciliation", jitteredPeriod)
	return reconcile.Result{RequeueAfter: jitteredPeriod}, nil
}

func (r *reconcileContext) doReconcile(ctx context.Context, u *unstructured.Unstructured) (requeue bool, err error) {
	logger := log.FromContext(ctx)

	cc, ccc, err := kccstate.FetchLiveKCCState(ctx, r.Reconciler.Client, r.NamespacedName)
	if err != nil {
		return true, err
	}

	am := resourceactuation.DecideActuationMode(cc, ccc)
	switch am {
	case v1beta1.Reconciling:
		logger.V(2).Info("Actuating a resource as actuation mode is \"Reconciling\"", "resource", r.NamespacedName)
	case v1beta1.Paused:
		logger.Info("Skipping actuation of resource as actuation mode is \"Paused\"", "resource", r.NamespacedName)

		// add finalizers for deletion defender to make sure we don't delete cloud provider resources when uninstalling
		if u.GetDeletionTimestamp().IsZero() {
			if err := r.ensureFinalizers(ctx, u); err != nil {
				return false, nil
			}
		}

		return false, nil
	default:
		return false, fmt.Errorf("unknown actuation mode %v", am)
	}

	adapter, err := r.Reconciler.model.AdapterForObject(ctx, r.Reconciler.Client, u)
	if err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(u))
			return r.handleUnresolvableDeps(ctx, u, unwrappedErr)
		}
		return false, r.handleUpdateFailed(ctx, u, err)
	}

	// To create, update or delete the GCP object, we need to get the GCP object first.
	// Because the object contains the cloud service information like `selfLink` `ID` required to validate
	// the resource uniqueness before updating/deleting.
	existsAlready, err := adapter.Find(ctx)
	if err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(u))
			return r.handleUnresolvableDeps(ctx, u, unwrappedErr)
		}
		return false, r.handleUpdateFailed(ctx, u, err)
	}

	defer execution.RecoverWithInternalError(&err)
	if !u.GetDeletionTimestamp().IsZero() {
		logger.Info("finalizing resource deletion", "resource", k8s.GetNamespacedName(u))
		if !k8s.HasFinalizer(u, k8s.ControllerFinalizerName) {
			logger.Info("no controller finalizer; no finalization necessary",
				"resource", k8s.GetNamespacedName(u))
			return false, nil
		}
		if k8s.HasFinalizer(u, k8s.DeletionDefenderFinalizerName) {
			// deletion defender has not yet finalized; requeuing
			logger.Info("deletion defender has not yet finalized; requeuing", "resource", k8s.GetNamespacedName(u))
			return true, nil
		}
		if k8s.HasAbandonAnnotation(u) {
			logger.Info("deletion policy set to abandon; abandoning underlying resource", "resource", k8s.GetNamespacedName(u))
			return false, r.handleDeleted(ctx, u)
		}
		if !existsAlready {
			logger.Info("underlying resource does not exist; no API call necessary", "resource", k8s.GetNamespacedName(u))
			return false, r.handleDeleted(ctx, u)
		}

		logger.Info("deleting underlying resource", "resource", k8s.GetNamespacedName(u))
		deleteOp := NewDeleteOperation(r.Reconciler.Client, u)
		if _, err := adapter.Delete(ctx, deleteOp); err != nil {
			if !errors.Is(err, k8s.ErrIAMNotFound) && !k8s.IsReferenceNotFoundError(err) {
				if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
					logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(u))
					resource, err := toK8sResource(u)
					if err != nil {
						return false, fmt.Errorf("error converting k8s resource while handling unresolvable dependencies event: %w", err)
					}
					// Requeue resource for reconciliation with exponential backoff applied
					return true, r.Reconciler.HandleUnresolvableDeps(ctx, resource, unwrappedErr)
				}
				return false, r.handleDeleteFailed(ctx, u, err)
			}
		}
		return false, r.handleDeleted(ctx, u)
	}

	if err := r.ensureFinalizers(ctx, u); err != nil {
		return false, err
	}

	// set the etag to an empty string, since IAMPolicy is the authoritative intent, KCC wants to overwrite the underlying policy regardless
	//policy.Spec.Etag = ""

	hasSetReadyCondition := false
	requeueRequested := false

	if !existsAlready {
		createOp := NewCreateOperation(r.Reconciler.LifecycleHandler, r.Reconciler.Client, u)
		if err := adapter.Create(ctx, createOp); err != nil {
			if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
				logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(u))
				return r.handleUnresolvableDeps(ctx, u, unwrappedErr)
			}
			return false, r.handleUpdateFailed(ctx, u, fmt.Errorf("error creating: %w", err))
		}
		hasSetReadyCondition = createOp.HasSetReadyCondition
		requeueRequested = createOp.RequeueRequested
	} else {
		updateOp := NewUpdateOperation(r.Reconciler.LifecycleHandler, r.Reconciler.Client, u)
		if err := adapter.Update(ctx, updateOp); err != nil {
			if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
				logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(u))
				return r.handleUnresolvableDeps(ctx, u, unwrappedErr)
			}
			return false, r.handleUpdateFailed(ctx, u, fmt.Errorf("error updating: %w", err))
		}
		hasSetReadyCondition = updateOp.HasSetReadyCondition
		requeueRequested = updateOp.RequeueRequested
	}

	if !hasSetReadyCondition && isAPIServerUpdateRequired(u) {
		return requeueRequested, r.handleUpToDate(ctx, u)
	}
	return requeueRequested, nil
}

// ensureFinalizers will apply our finalizers to the object if they are not present.
// We update the kube-apiserver immediately if any changes are needed.
func (r *reconcileContext) ensureFinalizers(ctx context.Context, u *unstructured.Unstructured) error {
	if k8s.EnsureFinalizers(u, k8s.ControllerFinalizerName, k8s.DeletionDefenderFinalizerName) {
		// No change
		return nil
	}

	if err := r.Reconciler.Client.Update(ctx, u); err != nil {
		return fmt.Errorf("updating finalizers: %w", err)
	}

	return nil
}

func (r *reconcileContext) handleUpToDate(ctx context.Context, u *unstructured.Unstructured) error {
	resource, err := toK8sResource(u)
	if err != nil {
		return fmt.Errorf("error converting to k8s resource while handling %v event: %w", k8s.UpToDate, err)
	}
	return r.Reconciler.HandleUpToDate(ctx, resource)
}

func (r *reconcileContext) handleUpdateFailed(ctx context.Context, policy *unstructured.Unstructured, origErr error) error {
	logger := log.FromContext(ctx)

	resource, err := toK8sResource(policy)
	if err != nil {
		logger.Error(err, "error converting to k8s resource while handling event",
			"resource", k8s.GetNamespacedName(policy), "event", k8s.UpdateFailed)
		return fmt.Errorf("Update call failed: %w", origErr)
	}
	return r.Reconciler.HandleUpdateFailed(ctx, resource, origErr)
}

func (r *reconcileContext) handleDeleted(ctx context.Context, policy *unstructured.Unstructured) error {
	resource, err := toK8sResource(policy)
	if err != nil {
		return fmt.Errorf("error converting to k8s resource while handling %v event: %w", k8s.Deleted, err)
	}
	return r.Reconciler.HandleDeleted(ctx, resource)
}

func (r *reconcileContext) handleDeleteFailed(ctx context.Context, policy *unstructured.Unstructured, origErr error) error {
	logger := log.FromContext(ctx)

	resource, err := toK8sResource(policy)
	if err != nil {
		logger.Error(err, "error converting to k8s resource while handling event",
			"resource", k8s.GetNamespacedName(policy), "event", k8s.DeleteFailed)
		return fmt.Errorf(k8s.DeleteFailedMessageTmpl, origErr)
	}
	return r.Reconciler.HandleDeleteFailed(ctx, resource, origErr)
}

func (r *DirectReconciler) supportsImmediateReconciliations() bool {
	return r.immediateReconcileRequests != nil
}

func (r *reconcileContext) handleUnresolvableDeps(ctx context.Context, policy *unstructured.Unstructured, origErr error) (requeue bool, err error) {
	logger := log.FromContext(ctx)

	resource, err := toK8sResource(policy)
	if err != nil {
		return false, fmt.Errorf("error converting to k8s resource while handling unresolvable dependencies event: %w", err)
	}
	refGVK, refNN, ok := lifecyclehandler.CausedByUnreadyOrNonexistentResourceRefs(origErr)
	if !ok || !r.Reconciler.supportsImmediateReconciliations() {
		// Requeue resource for reconciliation with exponential backoff applied
		return true, r.Reconciler.HandleUnresolvableDeps(ctx, resource, origErr)
	}
	// Check that the number of active resource watches
	// does not exceed the controller's cap. If the
	// capacity is not exceeded, The number of active
	// resource watches is incremented by one and a watch
	// is started
	if !r.Reconciler.resourceWatcherRoutines.TryAcquire(1) {
		// Requeue resource for reconciliation with exponential backoff applied
		return true, r.Reconciler.HandleUnresolvableDeps(ctx, resource, origErr)
	}
	// Create a logger for ResourceWatcher that contains info
	// about the referencing resource. This is done since the
	// messages logged by ResourceWatcher only include the
	// information of the resource it is watching by default.
	watcherLogger := logger.WithValues(
		"referencingResource", resource.GetNamespacedName(),
		"referencingResourceGVK", resource.GroupVersionKind())
	watcher, err := resourcewatcher.New(r.Reconciler.config, watcherLogger)
	if err != nil {
		return false, r.Reconciler.HandleUpdateFailed(ctx, resource, fmt.Errorf("error initializing new resourcewatcher: %w", err))
	}

	logger = logger.WithValues(
		"resource", resource.GetNamespacedName(),
		"resourceGVK", resource.GroupVersionKind(),
		"reference", refNN,
		"referenceGVK", refGVK)
	go func() {
		// Decrement the count of active resource watches after
		// the watch finishes
		defer r.Reconciler.resourceWatcherRoutines.Release(1)
		timeoutPeriod := r.Reconciler.jitterGenerator.WatchJitteredTimeout()
		ctx, cancel := context.WithTimeout(context.TODO(), timeoutPeriod)
		defer cancel()
		logger.Info("starting wait with timeout on resource's reference", "timeout", timeoutPeriod)
		if err := watcher.WaitForResourceToBeReady(ctx, refNN, refGVK); err != nil {
			logger.Error(err, "error while waiting for resource's reference to be ready")
			return
		}
		logger.Info("enqueuing resource for immediate reconciliation now that its reference is ready")
		r.Reconciler.enqueueForImmediateReconciliation(resource.GetNamespacedName())
	}()

	// Do not requeue resource for immediate reconciliation. Wait for either
	// the next periodic reconciliation or for the referenced resource to be ready (which
	// triggers a reconciliation), whichever comes first.
	return false, r.Reconciler.HandleUnresolvableDeps(ctx, resource, origErr)
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

func isAPIServerUpdateRequired(_ *unstructured.Unstructured) bool {
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
