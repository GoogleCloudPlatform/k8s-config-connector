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

package partialpolicy

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/kccstate"
	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	condition "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	kontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/metrics"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	kccratelimiter "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/ratelimiter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceactuation"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourcewatcher"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/execution"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpwatch"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/sync/semaphore"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const controllerName = "iampartialpolicy-controller"

var logger = log.Log.WithName(controllerName)

// Add creates a new IAM Partial Policy Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and start it when the Manager is started.
func Add(mgr manager.Manager, deps *kontroller.Deps) error {
	if deps.JitterGen == nil {
		var dclML metadata.ServiceMetadataLoader
		if deps.DclConverter != nil {
			dclML = deps.DclConverter.MetadataLoader
		}
		deps.JitterGen = jitter.NewDefaultGenerator(deps.TfLoader, dclML)
	}

	immediateReconcileRequests := make(chan event.GenericEvent, k8s.ImmediateReconcileRequestsBufferSize)
	resourceWatcherRoutines := semaphore.NewWeighted(k8s.MaxNumResourceWatcherRoutines)
	reconciler, err := NewReconciler(mgr, deps.TfProvider, deps.TfLoader, deps.DclConverter, deps.DclConfig, immediateReconcileRequests, resourceWatcherRoutines, deps.Defaulters, deps.JitterGen, deps.DependencyTracker)
	if err != nil {
		return err
	}
	return add(mgr, reconciler)
}

// NewReconciler returns a new reconcile.Reconciler.
func NewReconciler(mgr manager.Manager, provider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader, converter *conversion.Converter, dclConfig *mmdcl.Config, immediateReconcileRequests chan event.GenericEvent, resourceWatcherRoutines *semaphore.Weighted, defaulters []k8s.Defaulter, jg jitter.Generator, dependencyTracker *gcpwatch.DependencyTracker) (*ReconcileIAMPartialPolicy, error) {
	r := ReconcileIAMPartialPolicy{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(
			mgr.GetClient(),
			mgr.GetEventRecorderFor(controllerName),
		),
		Client:                     mgr.GetClient(),
		iamClient:                  kcciamclient.New(provider, smLoader, mgr.GetClient(), converter, dclConfig),
		scheme:                     mgr.GetScheme(),
		config:                     mgr.GetConfig(),
		defaulters:                 defaulters,
		immediateReconcileRequests: immediateReconcileRequests,
		resourceWatcherRoutines:    resourceWatcherRoutines,
		ReconcilerMetrics: metrics.ReconcilerMetrics{
			ResourceNameLabel: metrics.ResourceNameLabel,
		},
		requeueRateLimiter: kccratelimiter.RequeueRateLimiter(),
		jitterGen:          jg,
	}

	if r.immediateReconcileRequests != nil && dependencyTracker != nil {
		r.driftTracker = dependencyTracker.RegisterController(controllerName, r.immediateReconcileRequests)
	}
	return &r, nil
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler.
func add(mgr manager.Manager, r *ReconcileIAMPartialPolicy) error {
	obj := &iamv1beta1.IAMPartialPolicy{}
	_, err := builder.
		ControllerManagedBy(mgr).
		Named(controllerName).
		WithOptions(controller.Options{
			MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles,
			RateLimiter:             kccratelimiter.NewRateLimiter(),
		}).
		WatchesRawSource(source.TypedChannel(r.immediateReconcileRequests, &handler.EnqueueRequestForObject{})).
		For(obj, builder.OnlyMetadata, builder.WithPredicates(predicate.UnderlyingResourceOutOfSyncPredicate{})).
		Build(r)
	if err != nil {
		return fmt.Errorf("error creating new controller: %w", err)
	}
	return nil
}

var _ reconcile.Reconciler = &ReconcileIAMPartialPolicy{}

// ReconcileIAMPartialPolicy is a reconciler for handling IAM partial policies.
type ReconcileIAMPartialPolicy struct {
	lifecyclehandler.LifecycleHandler
	client.Client
	metrics.ReconcilerMetrics
	iamClient  *kcciamclient.IAMClient
	scheme     *runtime.Scheme
	config     *rest.Config
	defaulters []k8s.Defaulter
	// Fields used for triggering reconciliations when dependencies are ready
	immediateReconcileRequests chan event.GenericEvent
	resourceWatcherRoutines    *semaphore.Weighted // Used to cap number of goroutines watching unready dependencies

	// rate limit requeues (periodic re-reconciliation), so we don't use the whole rate limit on re-reconciles
	requeueRateLimiter workqueue.TypedRateLimiter[reconcile.Request]
	jitterGen          jitter.Generator
	driftTracker       *gcpwatch.ControllerRegistration
}

type reconcileContext struct {
	Reconciler     *ReconcileIAMPartialPolicy
	Ctx            context.Context
	NamespacedName types.NamespacedName

	objRef *iamv1beta1.IAMPolicy
}

func (r *ReconcileIAMPartialPolicy) Reconcile(ctx context.Context, request reconcile.Request) (result reconcile.Result, err error) {
	log := klog.FromContext(ctx)

	logger.Info("Running reconcile", "resource", request.NamespacedName)
	startTime := time.Now()
	ctx, cancel := context.WithTimeout(ctx, k8s.ReconcileDeadline)
	defer cancel()
	r.RecordReconcileWorkers(ctx, iamv1beta1.IAMPartialPolicyGVK)
	defer r.AfterReconcile()
	defer r.RecordReconcileMetrics(ctx, iamv1beta1.IAMPartialPolicyGVK, request.Namespace, request.Name, startTime, &err)

	policy := &iamv1beta1.IAMPartialPolicy{}
	if err := r.Get(context.TODO(), request.NamespacedName, policy); err != nil {
		if apierrors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	// r.Get() overrides the TypeMeta to empty value, so need to configure it
	// after r.Get().
	policy.SetGroupVersionKind(iamv1beta1.IAMPartialPolicyGVK)
	if err := r.handleDefaults(ctx, policy); err != nil {
		return reconcile.Result{}, fmt.Errorf("error handling default values for IAM partial policy '%v': %w", k8s.GetNamespacedName(policy), err)
	}
	runCtx := &reconcileContext{
		Reconciler:     r,
		Ctx:            ctx,
		NamespacedName: request.NamespacedName,
	}
	requeue, err := runCtx.doReconcile(policy)
	if err != nil {
		return reconcile.Result{}, err
	}
	if requeue {
		return reconcile.Result{Requeue: true}, nil
	}

	// if we can, we will use the GCP watch method to prompt reconcile events via the immediateReconcile channel
	if r.driftTracker != nil && runCtx.objRef != nil && r.driftTracker.Add(ctx, policy, runCtx.objRef.Spec.Etag) {
		log.V(2).Info("using gcp watcher instead of periodic requeue", "resource", request.NamespacedName)
		return reconcile.Result{}, nil
	}

	// otherwise we use the naive 10 minute + jiiter reconciliation for drift detection
	jitteredPeriod, err := r.jitterGen.JitteredReenqueue(iamv1beta1.IAMPolicyGVK, policy)
	if err != nil {
		return reconcile.Result{}, err
	}
	requeueDelay := r.requeueRateLimiter.When(request)
	requeueAfter := jitteredPeriod + requeueDelay
	logger.Info("successfully finished reconcile", "resource", request.NamespacedName, "time to next reconciliation", requeueAfter)
	return reconcile.Result{RequeueAfter: requeueAfter}, nil
}

func (r *ReconcileIAMPartialPolicy) handleDefaults(ctx context.Context, pp *iamv1beta1.IAMPartialPolicy) error {
	for _, defaulter := range r.defaulters {
		if _, err := defaulter.ApplyDefaults(ctx, pp); err != nil {
			return err
		}
	}
	return nil
}

func (r *reconcileContext) doReconcile(pp *iamv1beta1.IAMPartialPolicy) (requeue bool, err error) {
	defer execution.RecoverWithInternalError(&err)

	cc, ccc, err := kccstate.FetchLiveKCCState(r.Ctx, r.Reconciler.Client, r.NamespacedName)
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
		if pp.GetDeletionTimestamp().IsZero() {
			k8s.EnsureFinalizers(pp, k8s.ControllerFinalizerName, k8s.DeletionDefenderFinalizerName)
		}

		return false, nil
	default:
		return false, fmt.Errorf("unknown actuation mode %v", am)
	}

	if !pp.DeletionTimestamp.IsZero() {
		return r.finalizeDeletion(pp)
	}
	iamPolicy := ToIAMPolicySkeleton(pp)
	if iamPolicy, err = r.Reconciler.iamClient.GetPolicy(r.Ctx, iamPolicy); err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(pp))
			return r.handleUnresolvableDeps(pp, unwrappedErr)
		}
		return false, r.handleUpdateFailed(pp, err)
	}

	if r.objRef != nil {
		if k8s.GetNamespacedName(r.objRef) != k8s.GetNamespacedName(iamPolicy) {
			logger.Error(fmt.Errorf("object reference changed"), "old", r.objRef, "new", iamPolicy)
		}
	}
	r.objRef = iamPolicy.DeepCopy()

	k8s.EnsureFinalizers(pp, k8s.ControllerFinalizerName, k8s.DeletionDefenderFinalizerName)

	resolver := IAMMemberIdentityResolver{Iamclient: r.Reconciler.iamClient, Ctx: r.Ctx}
	desiredPartialPolicy, err := ComputePartialPolicyWithMergedBindings(pp, iamPolicy, &resolver)
	if err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(pp))
			return r.handleUnresolvableDeps(pp, unwrappedErr)
		}
		return false, r.handleUpdateFailed(pp, fmt.Errorf("error computing partial policy: %w", err))
	}
	desiredPolicy := toDesiredPolicy(desiredPartialPolicy, iamPolicy)
	if _, err = r.Reconciler.iamClient.SetPolicy(r.Ctx, desiredPolicy); err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(pp))
			return r.handleUnresolvableDeps(pp, unwrappedErr)
		}
		return false, r.handleUpdateFailed(pp, fmt.Errorf("error setting policy: %w", err))
	}
	if isAPIServerUpdateRequired(desiredPartialPolicy, pp) {
		return false, r.handleUpToDate(desiredPartialPolicy)
	}
	return false, nil
}

func (r *reconcileContext) finalizeDeletion(pp *iamv1beta1.IAMPartialPolicy) (requeue bool, err error) {
	if !k8s.HasFinalizer(pp, k8s.ControllerFinalizerName) {
		// Resource has no controller finalizer; no finalization necessary
		return false, nil
	}
	if k8s.HasFinalizer(pp, k8s.DeletionDefenderFinalizerName) {
		// deletion defender has not yet been finalized; requeuing
		logger.Info("deletion defender has not yet been finalized; requeuing", "resource", k8s.GetNamespacedName(pp))
		return true, nil
	}
	if !k8s.HasAbandonAnnotation(pp) {
		iamPolicy := ToIAMPolicySkeleton(pp)
		if iamPolicy, err = r.Reconciler.iamClient.GetPolicy(r.Ctx, iamPolicy); err != nil {
			if !errors.Is(err, kcciamclient.ErrNotFound) && !k8s.IsReferenceNotFoundError(err) {
				if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
					logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(pp))
					resource, err := toK8sResource(pp)
					if err != nil {
						return false, fmt.Errorf("error converting IAMPartialPolicy to k8s resource while handling unresolvable dependencies event: %w", err)
					}
					// Requeue resource for reconciliation with exponential backoff applied
					return true, r.Reconciler.HandleUnresolvableDeps(r.Ctx, resource, unwrappedErr)
				}
				return false, r.handleDeleteFailed(pp, err)
			}
			// If the GCP resource that the IAM policy is associated with is gone,
			// continue as the Policy is deleted successfully.
			return false, r.handleDeleted(pp)
		}
		// Compute the remaining bindings to set, i.e. pruning last applied bindings (bindings managed by KCC)
		// from all the existing bindings from the underlying IAM Policy.
		desiredPartialPolicy := ComputePartialPolicyWithRemainingBindings(pp, iamPolicy)
		desiredPolicy := toDesiredPolicy(desiredPartialPolicy, iamPolicy)
		if _, err = r.Reconciler.iamClient.SetPolicy(r.Ctx, desiredPolicy); err != nil {
			if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
				logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(pp))
				return r.handleUnresolvableDeps(pp, unwrappedErr)
			}
			return false, r.handleDeleteFailed(pp, fmt.Errorf("error setting policy: %w", err))
		}
		pp = desiredPartialPolicy
	}
	return false, r.handleDeleted(pp)
}

func (r *reconcileContext) handleUpToDate(policy *iamv1beta1.IAMPartialPolicy) error {
	resource, err := toK8sResource(policy)
	if err != nil {
		return fmt.Errorf("error converting IAMPartialPolicy to k8s resource while handling %v event: %w", k8s.UpToDate, err)
	}
	return r.Reconciler.HandleUpToDate(r.Ctx, resource)
}

func (r *reconcileContext) handleUpdateFailed(policy *iamv1beta1.IAMPartialPolicy, origErr error) error {
	resource, err := toK8sResource(policy)
	if err != nil {
		logger.Error(err, "error converting IAMPartialPolicy to k8s resource while handling event",
			"resource", k8s.GetNamespacedName(policy), "event", k8s.UpdateFailed)
		return fmt.Errorf("update call failed: %w", origErr)
	}
	return r.Reconciler.HandleUpdateFailed(r.Ctx, resource, origErr)
}

func (r *reconcileContext) handleDeleted(policy *iamv1beta1.IAMPartialPolicy) error {
	resource, err := toK8sResource(policy)
	if err != nil {
		return fmt.Errorf("error converting IAMPartialPolicy to k8s resource while handling %v event: %w", k8s.Deleted, err)
	}
	return r.Reconciler.HandleDeleted(r.Ctx, resource)
}

func (r *reconcileContext) handleDeleteFailed(policy *iamv1beta1.IAMPartialPolicy, origErr error) error {
	resource, err := toK8sResource(policy)
	if err != nil {
		logger.Error(err, "error converting IAMPartialPolicy to k8s resource while handling event",
			"resource", k8s.GetNamespacedName(policy), "event", k8s.DeleteFailed)
		return fmt.Errorf(k8s.DeleteFailedMessageTmpl, origErr)
	}
	return r.Reconciler.HandleDeleteFailed(r.Ctx, resource, origErr)
}

func (r *ReconcileIAMPartialPolicy) supportsImmediateReconciliations() bool {
	return r.immediateReconcileRequests != nil
}

func (r *reconcileContext) handleUnresolvableDeps(policy *iamv1beta1.IAMPartialPolicy, origErr error) (requeue bool, err error) {
	resource, err := toK8sResource(policy)
	if err != nil {
		return false, fmt.Errorf("error converting IAMPartialPolicy to k8s resource while handling unresolvable dependencies event: %w", err)
	}
	refGVK, refNN, ok := lifecyclehandler.CausedByUnreadyOrNonexistentResourceRefs(origErr)
	if !ok || !r.Reconciler.supportsImmediateReconciliations() {
		// Requeue resource for reconciliation with exponential backoff applied
		return true, r.Reconciler.HandleUnresolvableDeps(r.Ctx, resource, origErr)
	}
	// Check that the number of active resource watches
	// does not exceed the controller's cap. If the
	// capacity is not exceeded, The number of active
	// resource watches is incremented by one and a watch
	// is started
	if !r.Reconciler.resourceWatcherRoutines.TryAcquire(1) {
		// Requeue resource for reconciliation with exponential backoff applied
		return true, r.Reconciler.HandleUnresolvableDeps(r.Ctx, resource, origErr)
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
		return false, r.Reconciler.HandleUpdateFailed(r.Ctx, resource, fmt.Errorf("error initializing new resourcewatcher: %w", err))
	}

	logger := logger.WithValues(
		"resource", resource.GetNamespacedName(),
		"resourceGVK", resource.GroupVersionKind(),
		"reference", refNN,
		"referenceGVK", refGVK)
	go func() {
		// Decrement the count of active resource watches after
		// the watch finishes
		defer r.Reconciler.resourceWatcherRoutines.Release(1)
		timeoutPeriod := r.Reconciler.jitterGen.WatchJitteredTimeout()
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
	return false, r.Reconciler.HandleUnresolvableDeps(r.Ctx, resource, origErr)
}

// enqueueForImmediateReconciliation enqueues the given resource for immediate
// reconciliation. Note that this function only takes in the name and namespace
// of the resource and not its GVK since the controller instance that this
// reconcile instance belongs to can only reconcile resources of one GVK.
func (r *ReconcileIAMPartialPolicy) enqueueForImmediateReconciliation(resourceNN types.NamespacedName) {
	genEvent := event.GenericEvent{}
	genEvent.Object = &unstructured.Unstructured{}
	genEvent.Object.SetNamespace(resourceNN.Namespace)
	genEvent.Object.SetName(resourceNN.Name)
	r.immediateReconcileRequests <- genEvent
}

// IAMMemberIdentityResolver helps to resolve referenced member identity
type IAMMemberIdentityResolver struct {
	Iamclient *kcciamclient.IAMClient
	Ctx       context.Context
}

func (t *IAMMemberIdentityResolver) Resolve(member iamv1beta1.Member, memberFrom *iamv1beta1.MemberSource, defaultNamespace string) (string, error) {
	return kcciamclient.ResolveMemberIdentity(t.Ctx, member, memberFrom, defaultNamespace, t.Iamclient.TFIAMClient)
}

func isAPIServerUpdateRequired(desired, original *iamv1beta1.IAMPartialPolicy) bool {
	// TODO: even in the event of an actual update to GCP, this function will
	// return false because the condition comparison doesn't account for time.
	conditions := []condition.Condition{
		k8s.NewCustomReadyCondition(corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage),
	}
	if !k8s.ConditionSlicesEqual(original.Status.Conditions, conditions) {
		return true
	}
	if original.Status.ObservedGeneration != original.GetGeneration() {
		return true
	}
	if !reflect.DeepEqual(desired.Status.LastAppliedBindings, original.Status.LastAppliedBindings) {
		return true
	}
	if !reflect.DeepEqual(desired.Status.AllBindings, original.Status.AllBindings) {
		return true
	}
	return false
}

func toK8sResource(policy *iamv1beta1.IAMPartialPolicy) (*k8s.Resource, error) {
	kcciamclient.SetGVK(policy)
	resource := k8s.Resource{}
	if err := util.Marshal(policy, &resource); err != nil {
		return nil, fmt.Errorf("error marshalling IAMPartialPolicy to k8s resource: %w", err)
	}
	return &resource, nil
}

// ToIAMPolicySkeleton creates an IAMPolicy struct with ObjectMeta and resource reference
// copied from the partial policy. The skeleton struct can be passed to IAMClient.GetPolicy()
// to fetch the live IAM policy.
func ToIAMPolicySkeleton(p *iamv1beta1.IAMPartialPolicy) *iamv1beta1.IAMPolicy {
	res := &iamv1beta1.IAMPolicy{
		TypeMeta: metav1.TypeMeta{
			Kind:       iamv1beta1.IAMPolicyGVK.Kind,
			APIVersion: iamv1beta1.IAMAPIVersion,
		},
	}
	res.ObjectMeta = *p.ObjectMeta.DeepCopy()
	res.Spec.ResourceReference = p.Spec.ResourceReference
	return res
}

func toDesiredPolicy(desiredPartialPolicy *iamv1beta1.IAMPartialPolicy, livePolicy *iamv1beta1.IAMPolicy) *iamv1beta1.IAMPolicy {
	desiredPolicy := ToIAMPolicySkeleton(desiredPartialPolicy)
	desiredPolicy.Spec.Bindings = desiredPartialPolicy.Status.AllBindings
	// Carry the current etag from read to support concurrent read-modify-write operations from multiple systems.
	// SetPolicy will fail if the policy has been modified by other actors since the controller retrieved it.
	desiredPolicy.Spec.Etag = livePolicy.Spec.Etag
	// Preserve the audit configs if any.
	desiredPolicy.Spec.AuditConfigs = livePolicy.Spec.AuditConfigs
	return desiredPolicy
}
