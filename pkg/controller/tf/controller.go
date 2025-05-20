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

package tf

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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/execution"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/lease/leaser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/managementconflict"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	tfresource "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/resource"
	"github.com/go-logr/logr"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

type Reconciler struct {
	lifecyclehandler.LifecycleHandler
	metrics.ReconcilerMetrics
	resourceLeaser  *leaser.ResourceLeaser
	defaulters      []k8s.Defaulter
	mgr             manager.Manager
	schemaRef       *k8s.SchemaReference
	schemaRefMu     sync.RWMutex
	provider        *tfschema.Provider
	smLoader        *servicemappingloader.ServiceMappingLoader
	logger          logr.Logger
	jitterGenerator jitter.Generator
	// Fields used for triggering reconciliations when dependencies are ready
	immediateReconcileRequests chan event.GenericEvent
	resourceWatcherRoutines    *semaphore.Weighted // Used to cap number of goroutines watching unready dependencies
}

func Add(mgr manager.Manager, crd *apiextensions.CustomResourceDefinition, provider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader, defaulters []k8s.Defaulter, jitterGenerator jitter.Generator, additionalPredicate predicate.Predicate) (k8s.SchemaReferenceUpdater, error) {
	kind := crd.Spec.Names.Kind
	apiVersion := k8s.GetAPIVersionFromCRD(crd)
	controllerName := fmt.Sprintf("%v-controller", strings.ToLower(kind))
	immediateReconcileRequests := make(chan event.GenericEvent, k8s.ImmediateReconcileRequestsBufferSize)
	resourceWatcherRoutines := semaphore.NewWeighted(k8s.MaxNumResourceWatcherRoutines)
	r, err := NewReconciler(mgr, crd, provider, smLoader, immediateReconcileRequests, resourceWatcherRoutines, defaulters, jitterGenerator)
	if err != nil {
		return nil, err
	}
	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       kind,
			"apiVersion": apiVersion,
		},
	}
	predicateList := []predicate.Predicate{kccpredicate.UnderlyingResourceOutOfSyncPredicate{}}
	if additionalPredicate != nil {
		predicateList = append(predicateList, additionalPredicate)
	}
	_, err = builder.
		ControllerManagedBy(mgr).
		Named(controllerName).
		WithOptions(controller.Options{
			MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles,
			RateLimiter:             ratelimiter.NewRateLimiter(),
		}).
		WatchesRawSource(
			source.TypedChannel(immediateReconcileRequests, &handler.EnqueueRequestForObject{})).
		For(obj, builder.OnlyMetadata, builder.WithPredicates(predicateList...)).
		Build(r)
	if err != nil {
		return nil, fmt.Errorf("error creating new controller: %w", err)
	}
	log := mgr.GetLogger()
	log.Info("Registered controller", "kind", kind, "apiVersion", apiVersion)
	return r, nil
}

func NewReconciler(mgr manager.Manager,
	crd *apiextensions.CustomResourceDefinition,
	p *tfschema.Provider,
	smLoader *servicemappingloader.ServiceMappingLoader,
	immediateReconcileRequests chan event.GenericEvent,
	resourceWatcherRoutines *semaphore.Weighted,
	defaulters []k8s.Defaulter,
	jitterGenerator jitter.Generator) (*Reconciler, error) {

	if jitterGenerator == nil {
		return nil, fmt.Errorf("jitterGenerator must not be nil")
	}

	controllerName := fmt.Sprintf("%v-controller", strings.ToLower(crd.Spec.Names.Kind))
	return &Reconciler{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(
			mgr.GetClient(),
			mgr.GetEventRecorderFor(controllerName),
		),
		resourceLeaser: leaser.NewResourceLeaser(p, smLoader, mgr.GetClient()),
		defaulters:     defaulters,
		mgr:            mgr,
		schemaRef: &k8s.SchemaReference{
			CRD:        crd,
			JSONSchema: k8s.GetOpenAPIV3SchemaFromCRD(crd),
			GVK: schema.GroupVersionKind{
				Group:   crd.Spec.Group,
				Version: k8s.GetVersionFromCRD(crd),
				Kind:    crd.Spec.Names.Kind,
			},
		},
		ReconcilerMetrics: metrics.ReconcilerMetrics{
			ResourceNameLabel: metrics.ResourceNameLabel,
		},
		provider:                   p,
		smLoader:                   smLoader,
		logger:                     logger.WithName(controllerName),
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

	structuredreporting.ReportReconcileStart(ctx, u)

	skip, err := resourceactuation.ShouldSkip(u)
	if err != nil {
		return reconcile.Result{}, err
	}
	if skip {
		r.logger.Info("Skipping reconcile as nothing has changed and 0 reconcile period is set", "resource", req.NamespacedName)
		return reconcile.Result{}, nil
	}

	sm, err := r.smLoader.GetServiceMapping(u.GroupVersionKind().Group)
	if err != nil {
		return reconcile.Result{}, err
	}
	u, err = k8s.TriggerManagedFieldsMetadata(ctx, r.Client, u)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error triggering Server-Side Apply (SSA) metadata: %w", err)
	}
	resource, err := krmtotf.NewResource(u, sm, r.provider)
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

		// add finalizers for deletion defender to make sure we don't delete cloud provider resources when uninstalling
		if resource.GetDeletionTimestamp().IsZero() {
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
		return reconcile.Result{}, err
	}
	r.logger.Info("successfully finished reconcile", "resource", k8s.GetNamespacedName(resource), "time to next reconciliation", jitteredPeriod)
	return reconcile.Result{RequeueAfter: jitteredPeriod}, nil
}

func (r *Reconciler) sync(ctx context.Context, krmResource *krmtotf.Resource) (requeue bool, err error) {
	// isolate any panics to only this function
	defer execution.RecoverWithInternalError(&err)
	if !krmResource.GetDeletionTimestamp().IsZero() {
		// Deleting
		r.logger.Info("finalizing resource deletion", "resource", k8s.GetNamespacedName(krmResource))
		if !k8s.HasFinalizer(krmResource, k8s.ControllerFinalizerName) {
			r.logger.Info("no controller finalizer is present; no finalization necessary",
				"resource", k8s.GetNamespacedName(krmResource))
			return false, nil
		}
		if k8s.HasFinalizer(krmResource, k8s.DeletionDefenderFinalizerName) {
			r.logger.Info("deletion defender has not yet been finalized; requeuing", "resource", k8s.GetNamespacedName(krmResource))
			return true, nil
		}
		if err := r.HandleDeleting(ctx, &krmResource.Resource); err != nil {
			return false, err
		}
		if k8s.HasAbandonAnnotation(krmResource) {
			r.logger.Info("deletion policy set to abandon; abandoning underlying resource", "resource", k8s.GetNamespacedName(krmResource))
			return false, r.handleDeleted(ctx, krmResource)
		}

		if krmtotf.ShouldResolveParentForDelete(krmResource) {
			orphaned, parent, err := r.isOrphaned(krmResource)
			// Handle orphaned resources
			if err != nil {
				return false, err
			}
			if orphaned {
				r.logger.Info("resource has been orphaned; no API call necessary", "resource", k8s.GetNamespacedName(krmResource))
				return false, r.handleDeleted(ctx, krmResource)
			}

			if parent != nil && !k8s.IsResourceReady(parent) {
				if krmtotf.ShouldCheckParentReadyForDelete(krmResource, parent) {
					// If this resource has a parent and is not orphaned, ensure its parent
					// is ready before attempting deletion.
					// Requeue resource for reconciliation with exponential backoff applied
					return true, r.HandleUnresolvableDeps(ctx, &krmResource.Resource, k8s.NewReferenceNotReadyErrorForResource(parent))
				}
			}
		}
		liveState, err := krmtotf.FetchLiveStateForDelete(ctx, krmResource, r.provider, r, r.smLoader)
		if err != nil {
			return false, r.HandleDeleteFailed(ctx, &krmResource.Resource, fmt.Errorf("error fetching live state: %w", err))
		}
		if liveState.Empty() {
			r.logger.Info("underlying resource does not exist; no API call necessary", "resource", k8s.GetNamespacedName(krmResource))
			return false, r.handleDeleted(ctx, krmResource)
		}
		if err := r.obtainResourceLeaseIfNecessary(ctx, krmResource, liveState); err != nil {
			return false, err
		}
		r.logger.Info("deleting underlying resource", "resource", k8s.GetNamespacedName(krmResource))
		if _, err := krmResource.TFResource.Apply(ctx, liveState, &terraform.InstanceDiff{Destroy: true}, r.provider.Meta()); err != nil {
			return false, r.HandleDeleteFailed(ctx, &krmResource.Resource, fmt.Errorf("error deleting resource: %v", err))
		}
		return false, r.handleDeleted(ctx, krmResource)
	}
	liveState, err := krmtotf.FetchLiveStateForCreateAndUpdate(ctx, krmResource, r.provider, r, r.smLoader)
	if err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			r.logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(krmResource))
			return r.handleUnresolvableDeps(ctx, &krmResource.Resource, unwrappedErr)
		}
		return false, r.HandleUpdateFailed(ctx, &krmResource.Resource, fmt.Errorf("error fetching live state: %w", err))
	}
	if err := r.obtainResourceLeaseIfNecessary(ctx, krmResource, liveState); err != nil {
		return false, err
	}
	ok, err := r.hasServerGeneratedIDAndHadBeenCreatedOnceAlready(krmResource)
	if err != nil {
		return false, r.HandleUpdateFailed(ctx, &krmResource.Resource, err)
	}
	if ok && liveState.Empty() {
		// GCP resource with server-generated ID had been created once already,
		// but no longer exists. Don't "recreate" the resource, since
		// "recreating" resources with server-generated IDs technically creates
		// a brand new resource instead.
		return false, r.HandleUpdateFailed(ctx, &krmResource.Resource,
			fmt.Errorf("underlying resource no longer exists and can't be recreated without creating a brand new resource"))
	}
	config, secretVersions, err := krmtotf.KRMResourceToTFResourceConfigFull(
		krmResource, r, r.smLoader, liveState, r.schemaRef.JSONSchema, true,
	)
	if err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			r.logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(krmResource))
			return r.handleUnresolvableDeps(ctx, &krmResource.Resource, unwrappedErr)
		}
		return false, r.HandleUpdateFailed(ctx, &krmResource.Resource, fmt.Errorf("error expanding resource configuration for kind %s: %w", krmResource.Kind, err))
	}
	// Apply last-minute apply overrides
	if err := resourceoverrides.Handler.PreTerraformApply(ctx, krmResource.GroupVersionKind(), &operations.PreTerraformApply{KRMResource: krmResource, TerraformConfig: config, LiveState: liveState}); err != nil {
		return false, r.HandleUpdateFailed(ctx, &krmResource.Resource, fmt.Errorf("error applying pre-apply transformation to resource: %w", err))
	}
	diff, err := krmResource.TFResource.Diff(ctx, liveState, config, r.provider.Meta())
	if err != nil {
		return false, r.HandleUpdateFailed(ctx, &krmResource.Resource, fmt.Errorf("error calculating diff: %w", err))
	}
	if !liveState.Empty() && diff.RequiresNew() {
		return false, r.HandleUpdateFailed(ctx, &krmResource.Resource,
			k8s.NewImmutableFieldsMutationError(tfresource.ImmutableFieldsFromDiff(diff)))
	}
	if err := r.EnsureFinalizers(ctx, krmResource.Original, &krmResource.Resource, k8s.ControllerFinalizerName, k8s.DeletionDefenderFinalizerName); err != nil {
		return false, err
	}
	if diff.Empty() {
		r.logger.Info("underlying resource already up to date", "resource", k8s.GetNamespacedName(krmResource))
		return false, r.handleUpToDate(ctx, krmResource, liveState, secretVersions)
	}

	// Report diff to structured-reporting subsystem
	{
		report := &structuredreporting.Diff{}
		u, err := krmResource.MarshalAsUnstructured()
		if err != nil {
			log := log.FromContext(ctx)
			log.Error(err, "error reporting diff")
		}
		report.Object = u
		if diff != nil {
			for k, attr := range diff.Attributes {
				report.Fields = append(report.Fields, structuredreporting.DiffField{
					ID:  k,
					Old: attr.Old,
					New: attr.New,
				})
			}
		}
		report.IsNewObject = liveState.Empty()
		structuredreporting.ReportDiff(ctx, report)
	}

	r.logger.Info("creating/updating underlying resource", "resource", k8s.GetNamespacedName(krmResource))
	if err := r.HandleUpdating(ctx, &krmResource.Resource); err != nil {
		return false, err
	}
	// If creating a new resource, turn off RequiresNew in the diff. This is
	// done to prevent TF from clearing providerMeta (which contains blueprint
	// attribution information) during Apply() (b/193567082#comment15).
	if liveState.Empty() {
		for _, d := range diff.Attributes {
			d.RequiresNew = false
		}
	}
	newState, diagnostics := krmResource.TFResource.Apply(ctx, liveState, diff, r.provider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		r.logger.Error(err, "error applying desired state", "resource", krmResource.GetNamespacedName())
		return false, r.HandleUpdateFailed(ctx, &krmResource.Resource, fmt.Errorf("error applying desired state: %w", err))
	}
	return false, r.handleUpToDate(ctx, krmResource, newState, secretVersions)
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
		// Decrement the count of active resource watches after
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

	// Do not requeue resource for immediate reconciliation. Wait for either
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

func (r *Reconciler) handleDefaults(ctx context.Context, resource *krmtotf.Resource) error {
	for _, defaulter := range r.defaulters {
		if _, err := defaulter.ApplyDefaults(ctx, resource); err != nil {
			return err
		}
	}
	return nil
}

func (r *Reconciler) applyChangesForBackwardsCompatibility(ctx context.Context, resource *krmtotf.Resource) error {
	rc := resource.ResourceConfig

	// Ensure the resource has a management-conflict-prevention-policy
	// annotation. This is done to be backwards compatible with resources
	// created before the webhook for defaulting the annotation was added.
	if err := managementconflict.EnsureManagementConflictPreventionAnnotationForTFBasedResource(ctx, r.Client, resource, &rc, r.provider.ResourcesMap); err != nil {
		return fmt.Errorf("error ensuring resource '%v' has a management conflict policy: %w", k8s.GetNamespacedName(resource), err)
	}

	// Ensure the resource has a hierarchical reference. This is done to be
	// backwards compatible with resources created before the webhook for
	// defaulting hierarchical references was added.
	if err := k8s.EnsureHierarchicalReference(ctx, &resource.Resource, rc.HierarchicalReferences, rc.Containers, r.Client); err != nil {
		return fmt.Errorf("error ensuring resource '%v' has a hierarchical reference: %w", k8s.GetNamespacedName(resource), err)
	}
	return nil
}

func (r *Reconciler) obtainResourceLeaseIfNecessary(ctx context.Context, krmResource *krmtotf.Resource, liveState *terraform.InstanceState) error {
	conflictPolicy, err := managementconflict.GetManagementConflictPreventionAnnotationValue(krmResource)
	if err != nil {
		return err
	}
	if conflictPolicy != managementconflict.ManagementConflictPreventionPolicyResource {
		return nil
	}
	ok, err := r.resourceLeaser.IsLeasable(krmResource)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("kind '%v' does not support usage of %v='%v'", krmResource.GroupVersionKind(),
			managementconflict.FullyQualifiedAnnotation, conflictPolicy)
	}
	// Use SoftObtain instead of Obtain so that obtaining the lease ONLY changes the 'labels' value on the local krmResource and does not write the results
	// to GCP. The reason to do that is to reduce the number of writes to GCP and therefore improve performance and reduce errors.
	// The labels are written to GCP by the main sync(...) function because the changes to the labels show up in the diff.
	if err := r.resourceLeaser.SoftObtain(ctx, &krmResource.Resource, krmtotf.GetLabelsFromState(krmResource, liveState)); err != nil {
		return r.HandleObtainLeaseFailed(ctx, &krmResource.Resource, fmt.Errorf("error obtaining lease on '%v': %w",
			k8s.GetNamespacedName(krmResource), err))
	}
	return nil
}

func (r *Reconciler) handleDeleted(ctx context.Context, resource *krmtotf.Resource) error {
	if err := resourceoverrides.Handler.PostActuationTransform(resource.Original, &resource.Resource, nil, nil); err != nil {
		return r.HandlePostActuationTransformFailed(ctx, &resource.Resource, fmt.Errorf("error applying post-actuation transformation to resource '%v': %w", resource.GetNamespacedName(), err))
	}
	return r.HandleDeleted(ctx, &resource.Resource)
}

func (r *Reconciler) handleUpToDate(ctx context.Context, resource *krmtotf.Resource, liveState *terraform.InstanceState, secretVersions map[string]string) error {
	resource.Spec, resource.Status = krmtotf.ResolveSpecAndStatusWithResourceID(resource, liveState)
	if err := updateMutableButUnreadableFieldsAnnotationFor(resource); err != nil {
		return err
	}
	if err := updateObservedSecretVersionsAnnotationFor(resource, secretVersions); err != nil {
		return err
	}
	if err := resourceoverrides.Handler.PostActuationTransform(resource.Original, &resource.Resource, liveState, nil); err != nil {
		return r.HandlePostActuationTransformFailed(ctx, &resource.Resource, fmt.Errorf("error applying post-actuation transformation to resource '%v': %w", resource.GetNamespacedName(), err))
	}
	if !k8s.IsSpecOrStatusUpdateRequired(&resource.Resource, resource.Original) &&
		!k8s.IsAnnotationsUpdateRequired(&resource.Resource, resource.Original) &&
		k8s.ReadyConditionMatches(&resource.Resource, corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage) {
		return nil
	}
	return r.HandleUpToDate(ctx, &resource.Resource)
}

// isOrphaned returns whether the resource has been orphaned (i.e. its parent
// Kubernetes resource has already been deleted). Note:
// * A resource with no parent will always return false.
// * A resource whose parent is an external resource will always return false.
// * Hierarchical resources are also considered parents.
// It is assumed that parent and hierarchical references are always at the top
// level.
func (r *Reconciler) isOrphaned(resource *krmtotf.Resource) (orphaned bool, parent *k8s.Resource, err error) {
	// Currently, it's assumed that parent reference fields only support one resource type.
	parentConfigs := make([]corekccv1alpha1.TypeConfig, 0)
	for _, ref := range resource.ResourceConfig.ResourceReferences {
		if krmtotf.IsRequiredParentReference(ref, resource) {
			parentConfigs = append(parentConfigs, ref.TypeConfig)
		}
	}
	if len(parentConfigs) == 0 {
		return false, nil, nil
	}
	return lifecyclehandler.IsOrphaned(&resource.Resource, parentConfigs, r.Client)
}

func (r *Reconciler) hasServerGeneratedIDAndHadBeenCreatedOnceAlready(resource *krmtotf.Resource) (bool, error) {
	if !resource.HasServerGeneratedIDField() {
		return false, nil
	}
	val, err := resource.GetServerGeneratedID()
	if err != nil {
		if _, ok := k8s.AsServerGeneratedIDNotFoundError(err); ok {
			return false, nil
		}
		return false, err
	}
	return val != "", nil
}

var _ k8s.SchemaReferenceUpdater = &Reconciler{}

func (r *Reconciler) UpdateSchema(crd *apiextensions.CustomResourceDefinition) error {
	r.schemaRefMu.Lock()
	defer r.schemaRefMu.Unlock()
	return k8s.UpdateSchema(r.schemaRef, crd)
}

func updateMutableButUnreadableFieldsAnnotationFor(resource *krmtotf.Resource) error {
	// The annotation should only be set for resources with mutable-but-unreadable fields.
	if len(resource.ResourceConfig.MutableButUnreadableFields) == 0 {
		k8s.RemoveAnnotation(k8s.MutableButUnreadableFieldsAnnotation, resource)
		return nil
	}
	annotationVal, err := krmtotf.MutableButUnreadableFieldsAnnotationFor(resource)
	if err != nil {
		return fmt.Errorf("error constructing value for %v: %w", k8s.MutableButUnreadableFieldsAnnotation, err)
	}
	k8s.SetAnnotation(k8s.MutableButUnreadableFieldsAnnotation, annotationVal, resource)
	return nil
}

func updateObservedSecretVersionsAnnotationFor(resource *krmtotf.Resource, secretVersions map[string]string) error {
	hasSensitiveFields := tfresource.TFResourceHasSensitiveFields(resource.TFResource)
	return k8s.UpdateOrRemoveObservedSecretVersionsAnnotation(&resource.Resource, secretVersions, hasSensitiveFields)
}
