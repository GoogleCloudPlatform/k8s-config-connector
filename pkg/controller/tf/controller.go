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
	"time"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dependencywatcher"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/metrics"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/ratelimiter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/execution"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/lease/leaser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	tfresource "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/resource"

	"github.com/go-logr/logr"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
	klog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var logger = klog.Log

type Reconciler struct {
	lifecyclehandler.LifecycleHandler
	metrics.ReconcilerMetrics
	resourceLeaser             *leaser.ResourceLeaser
	mgr                        manager.Manager
	crd                        *apiextensions.CustomResourceDefinition
	jsonSchema                 *apiextensions.JSONSchemaProps
	gvk                        schema.GroupVersionKind
	provider                   *tfschema.Provider
	smLoader                   *servicemappingloader.ServiceMappingLoader
	immediateReconcileRequests chan event.GenericEvent
	logger                     logr.Logger
}

func Add(mgr manager.Manager, crd *apiextensions.CustomResourceDefinition, provider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader) error {
	kind := crd.Spec.Names.Kind
	apiVersion := k8s.GetAPIVersionFromCRD(crd)
	controllerName := fmt.Sprintf("%v-controller", strings.ToLower(kind))
	immediateReconcileRequests := make(chan event.GenericEvent, k8s.ImmediateReconcileRequestsBufferSize)
	r, err := NewReconciler(mgr, crd, provider, smLoader, immediateReconcileRequests)
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
		WithOptions(controller.Options{MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles, RateLimiter: ratelimiter.NewRateLimiter()}).
		Watches(&source.Channel{Source: immediateReconcileRequests}, &handler.EnqueueRequestForObject{}).
		For(obj, builder.OnlyMetadata, builder.WithPredicates(predicate.UnderlyingResourceOutOfSyncPredicate{})).
		Build(r)
	if err != nil {
		return fmt.Errorf("error creating new controller: %v", err)
	}
	logger.Info("Registered controller", "kind", kind, "apiVersion", apiVersion)
	return nil
}

func NewReconciler(mgr manager.Manager, crd *apiextensions.CustomResourceDefinition, p *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader, immediateReconcileRequests chan event.GenericEvent) (*Reconciler, error) {
	controllerName := fmt.Sprintf("%v-controller", strings.ToLower(crd.Spec.Names.Kind))
	return &Reconciler{
		LifecycleHandler: lifecyclehandler.LifecycleHandler{
			Client:   mgr.GetClient(),
			Recorder: mgr.GetEventRecorderFor(controllerName),
		},
		resourceLeaser: leaser.NewResourceLeaser(p, smLoader, mgr.GetClient()),
		mgr:            mgr,
		crd:            crd,
		jsonSchema:     k8s.GetOpenAPIV3SchemaFromCRD(crd),
		gvk: schema.GroupVersionKind{
			Group:   crd.Spec.Group,
			Version: k8s.GetVersionFromCRD(crd),
			Kind:    crd.Spec.Names.Kind,
		},
		ReconcilerMetrics: metrics.ReconcilerMetrics{
			ResourceNameLabel: metrics.ResourceNameLabel,
		},
		provider:                   p,
		smLoader:                   smLoader,
		logger:                     logger.WithName(controllerName),
		immediateReconcileRequests: immediateReconcileRequests,
	}, nil
}

func (r *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (res reconcile.Result, err error) {
	r.logger.Info("starting reconcile", "resource", req.NamespacedName)
	startTime := time.Now()
	r.RecordReconcileWorkers(ctx, r.gvk)
	defer r.AfterReconcile()
	defer r.RecordReconcileMetrics(ctx, r.gvk, req.Namespace, req.Name, startTime, &err)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(r.gvk)

	if err := r.Get(ctx, req.NamespacedName, u); err != nil {
		if apierrors.IsNotFound(err) {
			r.logger.Info("resource not found in API server; finishing reconcile", "resource", req.NamespacedName)
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
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
		return reconcile.Result{}, fmt.Errorf("could not parse resource %s: %v", req.NamespacedName.String(), err)
	}
	if err := r.applyChangesForBackwardsCompatibility(ctx, resource); err != nil {
		return reconcile.Result{}, fmt.Errorf("error applying changes to resource '%v' for backwards compatibility: %v", k8s.GetNamespacedName(resource), err)
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
	jitteredPeriod := jitter.GenerateJitteredReenqueuePeriod()
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
		orphaned, parent, err := r.isOrphaned(ctx, krmResource)
		if err != nil {
			return false, err
		}
		if orphaned {
			r.logger.Info("resource has been orphaned; no API call necessary", "resource", k8s.GetNamespacedName(krmResource))
			return false, r.handleDeleted(ctx, krmResource)
		}
		if parent != nil && !k8s.IsResourceReady(parent) {
			// If this resource has a parent and is not orphaned, ensure its parent
			// is ready before attempting deletion.
			return r.handleUnresolvableDeps(ctx, &krmResource.Resource, k8s.NewReferenceNotReadyErrorForResource(parent))
		}
		liveState, err := krmtotf.FetchLiveState(ctx, krmResource, r.provider, r, r.smLoader)
		if err != nil {
			return false, r.HandleDeleteFailed(ctx, &krmResource.Resource, fmt.Errorf("error fetching live state: %v", err))
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
	liveState, err := krmtotf.FetchLiveState(ctx, krmResource, r.provider, r, r.smLoader)
	if err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			r.logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(krmResource))
			return r.handleUnresolvableDeps(ctx, &krmResource.Resource, unwrappedErr)
		}
		return false, r.HandleUpdateFailed(ctx, &krmResource.Resource, fmt.Errorf("error fetching live state: %v", err))
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
		krmResource, r, r.smLoader, liveState, r.jsonSchema, true, label.GetDefaultLabels(),
	)
	if err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			r.logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(krmResource))
			return r.handleUnresolvableDeps(ctx, &krmResource.Resource, unwrappedErr)
		}
		return false, r.HandleUpdateFailed(ctx, &krmResource.Resource, fmt.Errorf("error expanding resource configuration for kind %s: %v", krmResource.Kind, err))
	}
	diff, err := krmResource.TFResource.Diff(ctx, liveState, config, r.provider.Meta())
	if err != nil {
		return false, r.HandleUpdateFailed(ctx, &krmResource.Resource, fmt.Errorf("error calculating diff: %v", err))
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
		return false, r.HandleUpdateFailed(ctx, &krmResource.Resource, fmt.Errorf("error applying desired state: %v", err))
	}
	return false, r.handleUpToDate(ctx, krmResource, newState, secretVersions)
}

// TODO(b/238913094): remove function after 100% of TF resources have been approved for immediate reconcilation
func supportsImmediateReconciliation(resourceKind string) bool {
	switch resourceKind {
	case "ComputeTargetPool",
		"ComputeNetworkEndpointGroup":
		return true
	}
	return false
}

func (r *Reconciler) supportsImmediateReconciliations() bool {
	return r.immediateReconcileRequests != nil
}

func (r *Reconciler) handleUnresolvableDeps(ctx context.Context, resource *k8s.Resource, originErr error) (requeue bool, err error) {
	refGVK, refNN, ok := lifecyclehandler.CausedByUnresolvableResourceRefs(originErr)
	if !ok || !supportsImmediateReconciliation(resource.Kind) || !r.supportsImmediateReconciliations() {
		// Requeue resource for immediate reconciliation
		// with exponential backoff applied
		return true, r.HandleUnresolvableDeps(ctx, resource, originErr)
	}
	depWatcher, err := dependencywatcher.CreateWatchForResource(resource, r.mgr.GetConfig())
	if err != nil {
		return false, r.HandleUpdateFailed(ctx, resource, fmt.Errorf("error creating dependencyWatcher to watch reference: %v %v: %v", refGVK.Kind, refNN, err))
	}
	go func() {
		logger := r.logger.WithValues(
			"resource", resource.GetNamespacedName(),
			"resourceGVK", resource.GroupVersionKind(),
			"reference", refNN,
			"referenceGVK", refGVK)
		timeoutPeriod := jitter.GenerateJitteredReenqueuePeriod()
		ctx, cancel := context.WithTimeout(ctx, timeoutPeriod)
		defer cancel()
		logger.Info("starting wait with timeout on resource's reference", "timeout", timeoutPeriod)
		if err := depWatcher.WaitForReferenceToBeReady(ctx, refNN, refGVK); err != nil {
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

func (r *Reconciler) applyChangesForBackwardsCompatibility(ctx context.Context, resource *krmtotf.Resource) error {
	rc := resource.ResourceConfig

	// Ensure the resource has a management-conflict-prevention-policy
	// annotation. This is done to be backwards compatible with resources
	// created before the webhook for defaulting the annotation was added.
	if err := k8s.EnsureManagementConflictPreventionAnnotationForTFBasedResource(r.Client, ctx, resource, &rc, r.provider.ResourcesMap); err != nil {
		return fmt.Errorf("error ensuring resource '%v' has a management conflict policy: %v", k8s.GetNamespacedName(resource), err)
	}

	// Ensure the resource has a hierarchical reference. This is done to be
	// backwards compatible with resources created before the webhook for
	// defaulting hierarchical references was added.
	if err := k8s.EnsureHierarchicalReference(ctx, &resource.Resource, rc.HierarchicalReferences, rc.Containers, r.Client); err != nil {
		return fmt.Errorf("error ensuring resource '%v' has a hierarchical reference: %v", k8s.GetNamespacedName(resource), err)
	}

	// Ensure the resource has a state-into-spec annotation.
	// This is done to be backwards compatible with resources
	// created before the webhook for defaulting the annotation was added.
	if err := k8s.EnsureSpecIntoSateAnnotation(&resource.Resource); err != nil {
		return fmt.Errorf("error ensuring resource '%v' has a '%v' annotation: %v", k8s.GetNamespacedName(resource), k8s.StateIntoSpecAnnotation, err)
	}
	return nil
}

func (r *Reconciler) obtainResourceLeaseIfNecessary(ctx context.Context, krmResource *krmtotf.Resource, liveState *terraform.InstanceState) error {
	conflictPolicy, err := k8s.GetManagementConflictPreventionAnnotationValue(krmResource)
	if err != nil {
		return err
	}
	if conflictPolicy != k8s.ManagementConflictPreventionPolicyResource {
		return nil
	}
	ok, err := r.resourceLeaser.IsLeasable(krmResource)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("kind '%v' does not support usage of %v of '%v'", krmResource.GroupVersionKind(),
			k8s.ManagementConflictPreventionPolicyAnnotation, conflictPolicy)
	}
	// Use SoftObtain instead of Obtain so that obtaining the lease ONLY changes the 'labels' value on the local krmResource and does not write the results
	// to GCP. The reason to do that is to reduce the number of writes to GCP and therefore improve performance and reduce errors.
	// The labels are written to GCP by the main sync(...) function because the changes to the labels show up in the diff.
	if err := r.resourceLeaser.SoftObtain(ctx, &krmResource.Resource, krmtotf.GetLabelsFromState(krmResource, liveState)); err != nil {
		return r.HandleObtainLeaseFailed(ctx, &krmResource.Resource, fmt.Errorf("error obtaining lease on '%v': %v",
			k8s.GetNamespacedName(krmResource), err))
	}
	return nil
}

func (r *Reconciler) handleDeleted(ctx context.Context, resource *krmtotf.Resource) error {
	if err := resourceoverrides.Handler.PostActuationTransform(resource.Original, &resource.Resource); err != nil {
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
	if err := resourceoverrides.Handler.PostActuationTransform(resource.Original, &resource.Resource); err != nil {
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
func (r *Reconciler) isOrphaned(ctx context.Context, resource *krmtotf.Resource) (orphaned bool, parent *k8s.Resource, err error) {
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

func updateMutableButUnreadableFieldsAnnotationFor(resource *krmtotf.Resource) error {
	// The annotation should only be set for resources with mutable-but-unreadable fields.
	if len(resource.ResourceConfig.MutableButUnreadableFields) == 0 {
		k8s.RemoveAnnotation(k8s.MutableButUnreadableFieldsAnnotation, resource)
		return nil
	}
	annotationVal, err := krmtotf.MutableButUnreadableFieldsAnnotationFor(resource)
	if err != nil {
		return fmt.Errorf("error constructing value for %v: %v", k8s.MutableButUnreadableFieldsAnnotation, err)
	}
	k8s.SetAnnotation(k8s.MutableButUnreadableFieldsAnnotation, annotationVal, resource)
	return nil
}

func updateObservedSecretVersionsAnnotationFor(resource *krmtotf.Resource, secretVersions map[string]string) error {
	hasSensitiveFields := tfresource.TFResourceHasSensitiveFields(resource.TFResource)
	return k8s.UpdateOrRemoveObservedSecretVersionsAnnotation(&resource.Resource, secretVersions, hasSensitiveFields)
}
