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

package dclcontroller

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/kccstate"
	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/metrics"
	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/ratelimiter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceactuation"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourcewatcher"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	dclclientconfig "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/clientconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	dclcontainer "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension/container"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/kcclite"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/livestate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/execution"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/lease/leasable"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/lease/leaser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/managementconflict"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	"github.com/go-logr/logr"
	"github.com/nasa9084/go-openapi"
	"golang.org/x/sync/semaphore"
	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var logger = log.Log

// The DCL lifecycles define what operations are acceptable during Apply.
// KCC doesn't allow delete-and-recreate operation to enforce modifications on immutable fields.
var LifecycleParams = []mmdcl.ApplyOption{
	mmdcl.WithLifecycleParam(mmdcl.BlockDestruction),
}

type Reconciler struct {
	lifecyclehandler.LifecycleHandler
	metrics.ReconcilerMetrics
	resourceLeaser *leaser.ResourceLeaser
	defaulters     []k8s.Defaulter
	mgr            manager.Manager
	schemaRef      *k8s.SchemaReference
	schemaRefMu    sync.RWMutex
	logger         logr.Logger
	// DCL related fields
	schema    *openapi.Schema
	dclConfig *mmdcl.Config
	converter *conversion.Converter
	// TF related fields
	serviceMappingLoader *servicemappingloader.ServiceMappingLoader
	// Fields used for triggering reconciliations when dependencies are ready
	immediateReconcileRequests chan event.GenericEvent
	resourceWatcherRoutines    *semaphore.Weighted // Used to cap number of goroutines watching unready dependencies
	jitterGenerator            jitter.Generator
}

func Add(mgr manager.Manager, crd *apiextensions.CustomResourceDefinition, converter *conversion.Converter,
	dclConfig *mmdcl.Config, serviceMappingLoader *servicemappingloader.ServiceMappingLoader, defaulters []k8s.Defaulter, jitterGenerator jitter.Generator,
	additionalPredicate predicate.Predicate) (k8s.SchemaReferenceUpdater, error) {
	if jitterGenerator == nil {
		return nil, fmt.Errorf("jitter generator not initialized")
	}
	predicates := []predicate.Predicate{kccpredicate.UnderlyingResourceOutOfSyncPredicate{}}
	if additionalPredicate != nil {
		predicates = append(predicates, additionalPredicate)
	}
	kind := crd.Spec.Names.Kind
	apiVersion := k8s.GetAPIVersionFromCRD(crd)
	controllerName := fmt.Sprintf("%v-controller", strings.ToLower(kind))
	immediateReconcileRequests := make(chan event.GenericEvent, k8s.ImmediateReconcileRequestsBufferSize)
	resourceWatcherRoutines := semaphore.NewWeighted(k8s.MaxNumResourceWatcherRoutines)
	r, err := NewReconciler(mgr, crd, converter, dclConfig, serviceMappingLoader, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jitterGenerator)
	if err != nil {
		return nil, err
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
		WithOptions(controller.Options{MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles, RateLimiter: ratelimiter.NewRateLimiter()}).
		WatchesRawSource(source.TypedChannel(immediateReconcileRequests, &handler.EnqueueRequestForObject{})).
		For(obj, builder.OnlyMetadata, builder.WithPredicates(predicates...)).
		Build(r)
	if err != nil {
		return nil, fmt.Errorf("error creating new controller: %w", err)
	}
	logger.V(2).Info("Registered dcl controller", "kind", kind, "apiVersion", apiVersion)
	return r, nil
}

func NewReconciler(mgr manager.Manager, crd *apiextensions.CustomResourceDefinition, converter *conversion.Converter, dclConfig *mmdcl.Config, serviceMappingLoader *servicemappingloader.ServiceMappingLoader, immediateReconcileRequests chan event.GenericEvent, resourceWatcherRoutines *semaphore.Weighted, defaulters []k8s.Defaulter, jitterGenerator jitter.Generator) (*Reconciler, error) {
	controllerName := fmt.Sprintf("%v-controller", strings.ToLower(crd.Spec.Names.Kind))
	gvk := schema.GroupVersionKind{
		Group:   crd.Spec.Group,
		Version: k8s.GetVersionFromCRD(crd),
		Kind:    crd.Spec.Names.Kind,
	}
	dclSchema, err := dclschemaloader.GetDCLSchemaForGVK(gvk, converter.MetadataLoader, converter.SchemaLoader)
	if err != nil {
		return nil, err
	}

	return &Reconciler{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(
			mgr.GetClient(),
			mgr.GetEventRecorderFor(controllerName),
		),
		ReconcilerMetrics: metrics.ReconcilerMetrics{
			ResourceNameLabel: metrics.ResourceNameLabel,
		},
		mgr: mgr,
		schemaRef: &k8s.SchemaReference{
			CRD:        crd,
			JSONSchema: k8s.GetOpenAPIV3SchemaFromCRD(crd),
			GVK:        gvk,
		},
		resourceLeaser:             leaser.NewResourceLeaser(nil, nil, mgr.GetClient()),
		defaulters:                 defaulters,
		schema:                     dclSchema,
		logger:                     logger.WithName(controllerName),
		dclConfig:                  dclclientconfig.CopyAndModifyForKind(dclConfig, gvk.Kind),
		converter:                  converter,
		serviceMappingLoader:       serviceMappingLoader,
		immediateReconcileRequests: immediateReconcileRequests,
		resourceWatcherRoutines:    resourceWatcherRoutines,
		jitterGenerator:            jitterGenerator,
	}, nil
}

func (r *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (res reconcile.Result, err error) {
	r.schemaRefMu.RLock()
	defer r.schemaRefMu.RUnlock()
	r.logger.Info("starting reconcile", "resource", req.NamespacedName)
	startTime := time.Now()
	r.RecordReconcileWorkers(ctx, r.schemaRef.GVK)
	defer r.AfterReconcile()
	defer r.RecordReconcileMetrics(ctx, r.schemaRef.GVK, req.Namespace, req.Name, startTime, &err)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(r.schemaRef.GVK)

	if err := r.Get(ctx, req.NamespacedName, u); err != nil {
		if apierrors.IsNotFound(err) {
			r.logger.Info("resource not found in API server; finishing reconcile", "resource", req.NamespacedName)
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}
	skip, err := resourceactuation.ShouldSkip(u)
	if err != nil {
		return reconcile.Result{}, err
	}
	if skip {
		r.logger.Info("Skipping reconcile as nothing has changed and 0 reconcile period is set", "resource", req.NamespacedName)
		return reconcile.Result{}, nil
	}
	// u, err = k8s.TriggerManagedFieldsMetadata(ctx, r.Client, u)
	// if err != nil {
	// 	return reconcile.Result{}, fmt.Errorf("error triggering Server-Side Apply (SSA) metadata: %w", err)
	// }

	resource, err := dcl.NewResource(u, r.schema)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("could not parse resource %s: %w", req.NamespacedName.String(), err)
	}
	if err := r.handleDefaults(ctx, resource); err != nil {
		return reconcile.Result{}, fmt.Errorf("error handling default values for resource '%v': %w", k8s.GetNamespacedName(resource), err)
	}
	if err := r.applyChangesForBackwardsCompatibility(ctx, resource); err != nil {
		return reconcile.Result{}, fmt.Errorf("error applying changes to resource '%v' for backwards compatibility: %w", k8s.GetNamespacedName(resource), err)
	}

	cc, ccc, err := kccstate.FetchLiveKCCState(ctx, r.mgr.GetClient(), req.NamespacedName)
	if err != nil {
		return reconcile.Result{}, err
	}

	am := resourceactuation.DecideActuationMode(cc, ccc)
	switch am {
	case v1beta1.Reconciling:
		r.logger.V(2).Info("Actuating a resource as actuation mode is \"Reconciling\"", "resource", req.NamespacedName)
	case v1beta1.Paused:
		jitteredPeriod, err := r.jitterGenerator.JitteredReenqueue(r.schemaRef.GVK, u)
		if err != nil {
			return reconcile.Result{}, err
		}

		if resource.GetDeletionTimestamp().IsZero() {
			// add finalizers for deletion defender to make sure we don't delete cloud provider resources when uninstalling
			if err := r.EnsureFinalizers(ctx, resource.Original, &resource.Resource, k8s.ControllerFinalizerName, k8s.DeletionDefenderFinalizerName); err != nil {
				return reconcile.Result{}, err
			}
		}

		r.logger.Info("Skipping actuation of resource as actuation mode is \"Paused\"", "resource", req.NamespacedName, "time to next reconciliation", jitteredPeriod)
		return reconcile.Result{RequeueAfter: jitteredPeriod}, nil
	default:
		return reconcile.Result{}, fmt.Errorf("unknown actuation mode %v", am)
	}

	// Apply pre-actuation transformation.
	if err := resourceoverrides.Handler.PreActuationTransform(&resource.Resource); err != nil {
		return reconcile.Result{}, r.HandlePreActuationTransformFailed(ctx, &resource.Resource, fmt.Errorf("error applying pre-actuation transformation to resource '%v': %w", req.NamespacedName.String(), err))
	}
	requeue, err := r.sync(ctx, resource)
	if err != nil {
		return reconcile.Result{}, err
	}
	if requeue {
		return reconcile.Result{Requeue: true}, nil
	}
	jitteredPeriod, err := r.jitterGenerator.JitteredReenqueue(r.schemaRef.GVK, u)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error generating reconcile interval for resource %v", resource.GetNamespacedName())
	}
	r.logger.Info("successfully finished reconcile", "resource", resource.GetNamespacedName(), "time to next reconciliation", jitteredPeriod)
	return reconcile.Result{RequeueAfter: jitteredPeriod}, nil
}

func (r *Reconciler) sync(ctx context.Context, resource *dcl.Resource) (requeue bool, err error) {
	// isolate any panics to only this function
	defer execution.RecoverWithInternalError(&err)

	dclConfig := r.dclConfig
	if !resource.GetDeletionTimestamp().IsZero() {
		return r.finalizeResourceDeletion(ctx, resource, dclConfig)
	}

	liveLite, err := livestate.FetchLiveState(ctx, resource, dclConfig, r.converter, r.serviceMappingLoader, r.Client)
	if err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			r.logger.Info(unwrappedErr.Error(), "resource", resource.GetNamespacedName())
			return r.handleUnresolvableDeps(ctx, &resource.Resource, unwrappedErr)
		}
		return false, r.HandleUpdateFailed(ctx, &resource.Resource, fmt.Errorf("error fetching live state: %w", err))
	}

	ok, err := resource.HasServerGeneratedIDAndConfigured()
	if err != nil {
		return false, r.HandleUpdateFailed(ctx, &resource.Resource, err)
	}
	if liveLite == nil && ok {
		// GCP resource with server-generated ID had been created once already,
		// but no longer exists. Don't "recreate" the resource, since
		// "recreating" resources with server-generated IDs technically creates
		// a brand new resource instead.
		return false, r.HandleUpdateFailed(ctx, &resource.Resource,
			fmt.Errorf("underlying resource with server-generated Id %s no longer exists and can't be recreated without creating a brand new resource with a different identifier", resource.Spec[k8s.ResourceIDFieldName]))
	}

	// attempt to obtain the resource lease
	var liveLabels map[string]string
	if liveLite != nil {
		liveLabels = liveLite.GetLabels()
	} else {
		liveLabels = make(map[string]string, 0)
	}
	if err = r.obtainResourceLeaseIfNecessary(ctx, resource, liveLabels); err != nil {
		return false, r.HandleObtainLeaseFailed(ctx, &resource.Resource, err)
	}
	// construct the trimmed desired state by only preserving k8s-managed fields
	desired, err := r.constructDesiredStateWithManagedFields(resource)
	if err != nil {
		return false, r.HandleUpdateFailed(ctx, &resource.Resource, err)
	}
	// KCC Full to KCC Lite
	lite, secretVersions, err := kcclite.ToKCCLiteAndSecretVersions(desired, r.converter.MetadataLoader, r.converter.SchemaLoader, r.serviceMappingLoader, r.Client)
	if err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			r.logger.Info(unwrappedErr.Error(), "resource", resource.GetNamespacedName())
			return r.handleUnresolvableDeps(ctx, &resource.Resource, unwrappedErr)
		}
		return false, r.HandleUpdateFailed(ctx, &resource.Resource, fmt.Errorf("error converting the desired state to a KCC lite resource: %w", err))
	}
	// KCC Lite to DCL resource
	dclResource, err := r.converter.KRMObjectToDCLObject(lite)
	if err != nil {
		return false, r.HandleUpdateFailed(ctx, &resource.Resource, fmt.Errorf("error expanding resource configuration: %w", err))
	}
	// Construct the state hint to check the diffs for mutable-but-unreadable fields.
	// Note that the state hint will be ignored if the resource doesn't support
	// it so we always append it to the life cycle params.
	stateHint, err := r.getStateHint(liveLite)
	if err != nil {
		return false, r.HandleUpdateFailed(ctx, &resource.Resource, fmt.Errorf("error generating the state hint: %w", err))
	}
	stateHintApplyOption := dclunstruct.WithStateHint(stateHint)
	hasDiff, err := dclunstruct.HasDiff(ctx, dclConfig, dclResource, stateHintApplyOption)
	if err != nil {
		return false, r.HandleUpdateFailed(ctx, &resource.Resource, fmt.Errorf("error generating the diffs from desired state: %w", err))
	}

	// ensure the finalizers before apply
	if err := r.EnsureFinalizers(ctx, resource.Original, &resource.Resource, k8s.ControllerFinalizerName, k8s.DeletionDefenderFinalizerName); err != nil {
		return false, err
	}
	// check if there are diffs between the desired state and the underlying resource
	if !hasDiff {
		r.logger.Info("resource is already up to date", "resource", resource.GetNamespacedName())
		return r.updateSpecAndStatusWithLiveState(ctx, liveLite, resource, secretVersions)
	}
	// create or update the underlying resource
	r.logger.Info("creating/updating underlying resource", "resource", resource.GetNamespacedName())
	if err := r.HandleUpdating(ctx, &resource.Resource); err != nil {
		return false, err
	}
	lifecycleParams := append(LifecycleParams, stateHintApplyOption)
	newState, err := dclunstruct.Apply(ctx, dclConfig, dclResource, lifecycleParams...)
	if err != nil {
		r.logger.Error(err, "error applying desired state", "resource", resource.GetNamespacedName())
		return false, r.HandleUpdateFailed(ctx, &resource.Resource, fmt.Errorf("error applying desired state: %w", err))
	}
	// update k8s api server with the new state
	newLite, err := r.converter.DCLObjectToKRMObject(newState)
	if err != nil {
		return false, r.HandleUpdateFailed(ctx, &resource.Resource, err)
	}
	return r.updateSpecAndStatusWithLiveState(ctx, newLite, resource, secretVersions)
}

func (r *Reconciler) supportsImmediateReconciliations() bool {
	return r.immediateReconcileRequests != nil
}

func (r *Reconciler) handleUnresolvableDeps(ctx context.Context, resource *k8s.Resource, originErr error) (requeue bool, err error) {
	refGVK, refNN, ok := lifecyclehandler.CausedByUnreadyOrNonexistentResourceRefs(originErr)
	if !ok || !r.supportsImmediateReconciliations() {
		// Requeue resource for reconciliation with exponential backoff applied
		return true, r.HandleUnresolvableDeps(ctx, resource, originErr)
	}
	// Don't start a watch on the reference if there
	// are too many ongoing watches already
	if !r.resourceWatcherRoutines.TryAcquire(1) {
		// Requeue resource for reconciliation with exponential backoff applied
		return true, r.HandleUnresolvableDeps(ctx, resource, originErr)
	}

	logger := r.logger.WithValues(
		"resource", resource.GetNamespacedName(),
		"resourceGVK", resource.GroupVersionKind(),
		"reference", refNN,
		"referenceGVK", refGVK)
	// Create a logger for ResourceWatcher that contains info
	// about the referencing resource. This is done since the
	// messages logged by ResourceWatcher only include the
	// information of the resource it is watching by default.
	watcherLogger := r.logger.WithValues(
		"referencingResource", resource.GetNamespacedName(),
		"referencingResourceGVK", resource.GroupVersionKind())
	watcher, err := resourcewatcher.New(r.mgr.GetConfig(), watcherLogger)
	if err != nil {
		return false, r.HandleUpdateFailed(ctx, resource, fmt.Errorf("error initializing new resourcewatcher: %w", err))
	}

	go func() {
		// Decrement the number of active resource watches after
		// the watch finishes
		defer r.resourceWatcherRoutines.Release(1)
		timeoutPeriod := r.jitterGenerator.WatchJitteredTimeout()
		ctx, cancel := context.WithTimeout(ctx, timeoutPeriod)
		defer cancel()
		logger.Info("starting wait with timeout on resource's reference", "timeout", timeoutPeriod)
		if err := watcher.WaitForResourceToBeReady(ctx, refNN, refGVK); err != nil {
			logger.Error(err, "error while waiting for resource's reference to be ready")
			return
		}
		logger.Info("enqueuing resource for immediate reconciliation now that its reference is ready")
		r.enqueueForImmediateReconciliation(resource.GetNamespacedName())
	}()

	// Do not requeue resource immediately for reconciliation. Wait for either
	// the next periodic reconciliation or for the referenced resource to be ready (which
	// triggers a reconciliation), whichever comes first.
	return false, r.HandleUnresolvableDeps(ctx, resource, originErr)
}

// enqueueForImmediateReconciliation enqueues the given resource for immediate
// reconciliation. Note that this function only takes in the name and namespace
// of the resource and not its GVK since the controller instance that this
// reconcile instance belongs to can only reconcile resources of one GVK.
func (r *Reconciler) enqueueForImmediateReconciliation(resourceNN types.NamespacedName) {
	genEvent := event.GenericEvent{}
	genEvent.Object = &unstructured.Unstructured{}
	genEvent.Object.SetNamespace(resourceNN.Namespace)
	genEvent.Object.SetName(resourceNN.Name)
	r.immediateReconcileRequests <- genEvent
}

func (r *Reconciler) obtainResourceLeaseIfNecessary(ctx context.Context, resource *dcl.Resource, liveLabels map[string]string) error {
	conflictPolicy, err := managementconflict.GetManagementConflictPreventionPolicy(resource)
	if err != nil {
		return err
	}
	if conflictPolicy != managementconflict.ManagementConflictPreventionPolicyResource {
		return nil
	}
	ok, err := leasable.DCLSchemaSupportsLeasing(resource.Schema)
	if err != nil {
		return err
	}
	if !ok {
		return managementconflict.NewLeasingNotSupportedByKindError(resource.GroupVersionKind())
	}
	// Use SoftObtain instead of Obtain so that obtaining the lease ONLY changes the 'labels' value on the local krmResource and does not write the results
	// to GCP. The reason to do that is to reduce the number of writes to GCP and therefore improve performance and reduce errors.
	// The labels are written to GCP by the main sync(...) function because the changes to the labels show up in the diff.
	if err := r.resourceLeaser.SoftObtain(ctx, &resource.Resource, liveLabels); err != nil {
		return r.HandleObtainLeaseFailed(ctx, &resource.Resource, fmt.Errorf("error obtaining lease on '%v': %w",
			resource.GetNamespacedName(), err))
	}
	return nil
}

func (r *Reconciler) handleDefaults(ctx context.Context, resource *dcl.Resource) error {
	for _, defaulter := range r.defaulters {
		if _, err := defaulter.ApplyDefaults(ctx, resource); err != nil {
			return err
		}
	}
	return nil
}

func (r *Reconciler) applyChangesForBackwardsCompatibility(ctx context.Context, resource *dcl.Resource) error {
	// Ensure the resource has a hierarchical reference. This is done to be
	// backwards compatible with resources created before the webhook for
	// defaulting hierarchical references was added.
	gvk := resource.GroupVersionKind()
	containers, err := dclcontainer.GetContainersForGVK(gvk, r.converter.MetadataLoader, r.converter.SchemaLoader)
	if err != nil {
		return fmt.Errorf("error getting containers supported by GroupVersionKind %v: %w", gvk, err)
	}
	hierarchicalRefs, err := dcl.GetHierarchicalReferencesForGVK(gvk, r.converter.MetadataLoader, r.converter.SchemaLoader)
	if err != nil {
		return fmt.Errorf("error getting hierarchical references supported by GroupVersionKind %v: %w", gvk, err)
	}
	if err := k8s.EnsureHierarchicalReference(ctx, &resource.Resource, hierarchicalRefs, containers, r.Client); err != nil {
		return fmt.Errorf("error ensuring resource '%v' has a hierarchical reference: %w", k8s.GetNamespacedName(resource), err)
	}
	return nil
}

func (r *Reconciler) handleDeleted(ctx context.Context, resource *dcl.Resource) error {
	if err := resourceoverrides.Handler.PostActuationTransform(resource.Original, &resource.Resource, nil, nil); err != nil {
		return r.HandlePostActuationTransformFailed(ctx, &resource.Resource, fmt.Errorf("error applying post-actuation transformation to resource '%v': %w", resource.GetNamespacedName(), err))
	}
	return r.HandleDeleted(ctx, &resource.Resource)
}

func (r *Reconciler) updateSpecAndStatusWithLiveState(ctx context.Context, liveLite *unstructured.Unstructured, resource *dcl.Resource, secretVersions map[string]string) (requeue bool, err error) {
	newSpec, newStatus, err := kcclite.ResolveSpecAndStatus(liveLite, resource, r.converter.MetadataLoader)
	if err != nil {
		return false, r.HandleUpdateFailed(ctx, &resource.Resource, fmt.Errorf("error resolving the live state: %w", err))
	}
	resource.Spec, resource.Status = newSpec, newStatus

	if err := updateMutableButUnreadableFieldsAnnotationFor(resource); err != nil {
		return false, err
	}
	if err := updateObservedSecretVersionsAnnotationFor(resource, secretVersions); err != nil {
		return false, err
	}

	if err := resourceoverrides.Handler.PostActuationTransform(resource.Original, &resource.Resource, nil, liveLite); err != nil {
		return false, r.HandlePostActuationTransformFailed(ctx, &resource.Resource, fmt.Errorf("error applying post-actuation transformation to resource '%v': %w", resource.GetNamespacedName(), err))
	}

	if !k8s.IsSpecOrStatusUpdateRequired(&resource.Resource, resource.Original) &&
		!k8s.IsAnnotationsUpdateRequired(&resource.Resource, resource.Original) &&
		k8s.ReadyConditionMatches(&resource.Resource, corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage) {
		return false, nil
	}

	return false, r.HandleUpToDate(ctx, &resource.Resource)
}

func (r *Reconciler) constructDesiredStateWithManagedFields(original *dcl.Resource) (*dcl.Resource, error) {
	gvk := original.GroupVersionKind()
	hierarchicalRefs, err := dcl.GetHierarchicalReferencesForGVK(gvk, r.converter.MetadataLoader, r.converter.SchemaLoader)
	if err != nil {
		return nil, fmt.Errorf("error getting hierarchical references supported by GroupVersionKind %v: %w", gvk, err)
	}
	trimmed, err := k8s.ConstructTrimmedSpecWithManagedFields(&original.Resource, r.schemaRef.JSONSchema, hierarchicalRefs)
	if err != nil {
		return nil, err
	}
	u, err := original.MarshalAsUnstructured()
	if err != nil {
		return nil, nil
	}
	u.Object["spec"] = trimmed
	res := &dcl.Resource{
		Schema: original.Schema,
	}
	if err := util.Marshal(u, res); err != nil {
		return nil, err
	}
	if val, ok := original.Spec[k8s.ResourceIDFieldName]; ok {
		if res.Spec == nil {
			res.Spec = make(map[string]interface{})
		}
		res.Spec[k8s.ResourceIDFieldName] = val
	}
	return res, nil
}

// finalizeResourceDeletion reacts to the KCC resource object deletion from k8s cluster.
// It performs the following operations:
// 1) checks on the finalizers before any operation
// 2) checks the deletion policy and determines whether to abandon the underlying resource
// 3) checks if the resource is orphaned by its parent
// 4) deletes the underlying resources if it owns the resource lease
func (r *Reconciler) finalizeResourceDeletion(ctx context.Context, resource *dcl.Resource, dclConfig *mmdcl.Config) (requeue bool, err error) {
	r.logger.Info("finalizing resource deletion", "resource", resource.GetNamespacedName())
	if !k8s.HasFinalizer(resource, k8s.ControllerFinalizerName) {
		r.logger.Info("no controller finalizer is present; no finalization necessary",
			"resource", resource.GetNamespacedName())
		return false, nil
	}
	if k8s.HasFinalizer(resource, k8s.DeletionDefenderFinalizerName) {
		r.logger.Info("deletion defender has not yet been finalized; requeuing", "resource", resource.GetNamespacedName())
		return true, nil
	}
	if err := r.HandleDeleting(ctx, &resource.Resource); err != nil {
		return false, err
	}
	// abandon the resource if deletion policy is abandon
	if k8s.HasAbandonAnnotation(resource) {
		r.logger.Info("deletion policy set to abandon; abandoning underlying resource", "resource", resource.GetNamespacedName())
		return false, r.handleDeleted(ctx, resource)
	}
	// check if the resource is orphaned by its parent
	orphaned, parent, err := r.isOrphaned(ctx, resource)
	if err != nil {
		return false, err
	}
	if orphaned {
		r.logger.Info("resource has been orphaned; no API call necessary", "resource", resource.GetNamespacedName())
		return false, r.handleDeleted(ctx, resource)
	}
	if parent != nil && !k8s.IsResourceReady(parent) {
		// If this resource has a parent and is not orphaned, ensure its parent
		// is ready before attempting deletion.
		// Requeue resource for reconciliation with exponential backoff applied
		return true, r.HandleUnresolvableDeps(ctx, &resource.Resource, k8s.NewReferenceNotReadyErrorForResource(parent))
	}

	// check if the underlying resource exists
	liveLite, err := livestate.FetchLiveState(ctx, resource, r.dclConfig, r.converter, r.serviceMappingLoader, r.Client)
	if err != nil {
		return false, r.HandleDeleteFailed(ctx, &resource.Resource, fmt.Errorf("error fetching live state: %w", err))
	}
	if liveLite == nil {
		r.logger.Info("underlying resource does not exist; no API call necessary", "resource", k8s.GetNamespacedName(resource))
		return false, r.handleDeleted(ctx, resource)
	}
	// attempt to obtain the resource lease and delete the underlying resource
	if err = r.obtainResourceLeaseIfNecessary(ctx, resource, liveLite.GetLabels()); err != nil {
		return false, r.HandleObtainLeaseFailed(ctx, &resource.Resource, err)
	}
	lite, err := kcclite.ToKCCLiteBestEffort(resource, r.converter.MetadataLoader, r.converter.SchemaLoader, r.serviceMappingLoader, r.Client)
	if err != nil {
		return false, fmt.Errorf("error converting KCC full to KCC lite: %w", err)
	}
	dclResource, err := r.converter.KRMObjectToDCLObject(lite)
	if err != nil {
		return false, fmt.Errorf("error converting KCC lite to dcl resource: %w", err)
	}
	r.logger.Info("deleting underlying resource", "resource", resource.GetNamespacedName())
	if err := dclunstruct.Delete(ctx, dclConfig, dclResource); err != nil {
		if dcl.IsNoSuchMethodError(err) {
			r.logger.Info("underlying resource cannot be deleted since there is no delete API; only clean up the kubernetes resource object", "resource", k8s.GetNamespacedName(resource))
			return false, r.handleDeleted(ctx, resource)
		}
		return false, r.HandleDeleteFailed(ctx, &resource.Resource, fmt.Errorf("error deleting the resource %v: %w", resource.GetNamespacedName(), err))
	}
	return false, r.handleDeleted(ctx, resource)
}

// isOrphaned returns whether the resource has been orphaned (i.e. its parent
// Kubernetes resource has already been deleted). Note:
// * A resource with no parent will always return false.
// * A resource whose parent is an external resource will always return false.
// * Hierarchical resources are also considered parents.
// It is assumed that parent and hierarchical references are always at the top
// level.
func (r *Reconciler) isOrphaned(_ context.Context, resource *dcl.Resource) (orphaned bool, parent *k8s.Resource, err error) {
	gvk := resource.GroupVersionKind()
	resourceMetadata, found := r.converter.MetadataLoader.GetResourceWithGVK(gvk)
	if !found {
		return false, nil, fmt.Errorf("ServiceMetadata for resource with GroupVersionKind %v not found", gvk)
	}
	parentConfigs := make([]corekccv1alpha1.TypeConfig, 0)
	for k, s := range resource.Schema.Properties {
		if s.ReadOnly {
			continue
		}
		if !extension.IsReferenceField(s) {
			continue
		}
		if s.Type != "string" {
			continue
		}

		// TODO(b/186159460): Delete this if-block once all resources support
		// hierarchical references.
		if dcl.IsContainerField([]string{k}) && !resourceMetadata.SupportsHierarchicalReferences {
			continue
		}

		if dcl.IsMultiTypeParentReferenceField([]string{k}) {
			typeConfigs, err := dcl.GetReferenceTypeConfigs(s, r.converter.MetadataLoader)
			if err != nil {
				return false, nil, fmt.Errorf("error getting reference type configs for DCL field '%v': %w", k, err)
			}
			for _, tc := range typeConfigs {
				parentConfigs = append(parentConfigs, tc)
			}
			continue
		}

		refFieldName, err := extension.GetReferenceFieldName([]string{k}, s)
		if err != nil {
			return false, nil, err
		}
		typeConfigs, err := dcl.GetReferenceTypeConfigs(s, r.converter.MetadataLoader)
		if err != nil {
			return false, nil, fmt.Errorf("error getting reference type configs for DCL field '%v': %w", k, err)
		}
		// It is assumed that parent references and single-type hierarchical
		// references only support one resource type.
		if len(typeConfigs) != 1 {
			continue
		}
		tc := typeConfigs[0]
		if !tc.Parent {
			continue
		}
		tc.Key = refFieldName
		parentConfigs = append(parentConfigs, tc)
	}
	return lifecyclehandler.IsOrphaned(&resource.Resource, parentConfigs, r.Client)
}

// getStateHint returns a state hint based on the given resource live state. A
// state hint is a DCL unstructured object that represents the live state of the
// resource.
//
// This function returns the given live state (if it's not nil) as a DCL
// unstructured object. Otherwise, it returns nil.
func (r *Reconciler) getStateHint(liveLite *unstructured.Unstructured) (*dclunstruct.Resource, error) {
	if liveLite == nil {
		return nil, nil
	}

	stateHint, err := r.converter.KRMObjectToDCLObject(liveLite)
	if err != nil {
		return nil, err
	}
	return stateHint, nil
}

var _ k8s.SchemaReferenceUpdater = &Reconciler{}

func (r *Reconciler) UpdateSchema(crd *apiextensions.CustomResourceDefinition) error {
	r.schemaRefMu.Lock()
	defer r.schemaRefMu.Unlock()
	return k8s.UpdateSchema(r.schemaRef, crd)
}

func updateMutableButUnreadableFieldsAnnotationFor(resource *dcl.Resource) error {
	hasMutableButUnreadableFields, err := resource.HasMutableButUnreadableFields()
	if err != nil {
		return fmt.Errorf("error checking if resource has mutable-but-unreadable fields: %w", err)
	}

	// The annotation should only be set for resources with mutable-but-unreadable fields.
	if !hasMutableButUnreadableFields {
		k8s.RemoveAnnotation(k8s.MutableButUnreadableFieldsAnnotation, resource)
		return nil
	}

	annotationVal, err := dcl.MutableButUnreadableFieldsAnnotationFor(resource)
	if err != nil {
		return fmt.Errorf("error constructing value for %v: %w", k8s.MutableButUnreadableFieldsAnnotation, err)
	}
	k8s.SetAnnotation(k8s.MutableButUnreadableFieldsAnnotation, annotationVal, resource)
	return nil
}

func updateObservedSecretVersionsAnnotationFor(resource *dcl.Resource, secretVersions map[string]string) error {
	hasSensitiveFields, err := extension.HasSensitiveFields(resource.Schema)
	if err != nil {
		return fmt.Errorf("error checking if resource has sensitive fields: %w", err)
	}

	return k8s.UpdateOrRemoveObservedSecretVersionsAnnotation(&resource.Resource, secretVersions, hasSensitiveFields)
}
