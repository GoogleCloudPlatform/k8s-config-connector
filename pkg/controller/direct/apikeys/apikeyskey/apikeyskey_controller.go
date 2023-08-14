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

package apikeyskey

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"k8s.io/klog/v2"

	pb "cnrm.googlesource.com/cnrm/mockgcp/generated/google/api/apikeys/v2"
	"cnrm.googlesource.com/cnrm/pkg/clients/generated/apis/apikeys/v1alpha1"
	"cnrm.googlesource.com/cnrm/pkg/controller/direct/mappings"
	kcciamclient "cnrm.googlesource.com/cnrm/pkg/controller/iam/iamclient"
	"cnrm.googlesource.com/cnrm/pkg/controller/jitter"
	"cnrm.googlesource.com/cnrm/pkg/controller/lifecyclehandler"
	"cnrm.googlesource.com/cnrm/pkg/controller/metrics"
	"cnrm.googlesource.com/cnrm/pkg/controller/predicate"
	"cnrm.googlesource.com/cnrm/pkg/controller/ratelimiter"
	"cnrm.googlesource.com/cnrm/pkg/controller/resourcewatcher"
	"cnrm.googlesource.com/cnrm/pkg/dcl/conversion"
	"cnrm.googlesource.com/cnrm/pkg/execution"
	"cnrm.googlesource.com/cnrm/pkg/k8s"
	"cnrm.googlesource.com/cnrm/pkg/servicemapping/servicemappingloader"
	"cnrm.googlesource.com/cnrm/pkg/util"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"google.golang.org/api/googleapi"
	"google.golang.org/protobuf/encoding/protojson"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/sync/semaphore"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
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

const controllerName = "iampolicy-controller"

var logger = log.Log.WithName(controllerName)

// Add creates a new IAM Policy Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and start it when the Manager is started.
func Add(mgr manager.Manager, tfProvider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader,
	converter *conversion.Converter, dclConfig *mmdcl.Config) error {
	immediateReconcileRequests := make(chan event.GenericEvent, k8s.ImmediateReconcileRequestsBufferSize)
	resourceWatcherRoutines := semaphore.NewWeighted(k8s.MaxNumResourceWatcherRoutines)
	reconciler, err := NewReconciler(mgr, tfProvider, smLoader, converter, dclConfig, immediateReconcileRequests, resourceWatcherRoutines)
	if err != nil {
		return err
	}
	return add(mgr, reconciler)
}

// NewReconciler returns a new reconcile.Reconciler.
func NewReconciler(mgr manager.Manager, provider *tfschema.Provider,
	smLoader *servicemappingloader.ServiceMappingLoader, converter *conversion.Converter,
	dclConfig *mmdcl.Config, immediateReconcileRequests chan event.GenericEvent, resourceWatcherRoutines *semaphore.Weighted) (*DirectReconciler, error) {
	gvk := schema.GroupVersionKind{
		Group:   "apikeys.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "APIKeysKey",
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
		Named(controllerName).
		WithOptions(controller.Options{MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles, RateLimiter: ratelimiter.NewRateLimiter()}).
		Watches(&source.Channel{Source: r.immediateReconcileRequests}, &handler.EnqueueRequestForObject{}).
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
	gvk schema.GroupVersionKind
	//iamClient *kcciamclient.IAMClient
	scheme *runtime.Scheme
	config *rest.Config
	// Fields used for triggering reconciliations when dependencies are ready
	immediateReconcileRequests chan event.GenericEvent
	resourceWatcherRoutines    *semaphore.Weighted // Used to cap number of goroutines watching unready dependencies
}

type reconcileContext struct {
	gvk            schema.GroupVersionKind
	Reconciler     *DirectReconciler
	NamespacedName types.NamespacedName
}

// Reconcile checks k8s for the current state of the resource.
func (r *DirectReconciler) Reconcile(ctx context.Context, request reconcile.Request) (result reconcile.Result, err error) {
	logger.Info("Running reconcile", "resource", request.NamespacedName)
	startTime := time.Now()
	ctx, cancel := context.WithTimeout(ctx, k8s.ReconcileDeadline)
	defer cancel()
	r.RecordReconcileWorkers(ctx, r.gvk)
	defer r.AfterReconcile()
	defer r.RecordReconcileMetrics(ctx, r.gvk, request.Namespace, request.Name, startTime, &err)

	policy := &unstructured.Unstructured{}
	policy.SetGroupVersionKind(r.gvk)
	if err := r.Get(ctx, request.NamespacedName, policy); err != nil {
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
	requeue, err := runCtx.doReconcile(ctx, policy)
	if err != nil {
		return reconcile.Result{}, err
	}
	if requeue {
		return reconcile.Result{Requeue: true}, nil
	}
	jitteredPeriod, err := jitter.GenerateJitteredReenqueuePeriod(r.gvk, nil, nil, policy)
	if err != nil {
		return reconcile.Result{}, err
	}
	logger.Info("successfully finished reconcile", "resource", request.NamespacedName, "time to next reconciliation", jitteredPeriod)
	return reconcile.Result{RequeueAfter: jitteredPeriod}, nil
}

// type UnstructuredKey struct {
// 	Data map[string]interface{}
// }

func (r *reconcileContext) doReconcile(ctx context.Context, u *unstructured.Unstructured) (requeue bool, err error) {
	// log := log.FromContext(ctx)

	adapter, err := AdapterForObject(ctx, u)
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
			if err := adapter.Delete(ctx); err != nil {
				if !errors.Is(err, kcciamclient.NotFoundError) && !k8s.IsReferenceNotFoundError(err) {
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
		}
		return false, r.handleDeleted(ctx, u)
	}

	existsAlready, err := adapter.Find(ctx)
	if err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(u))
			return r.handleUnresolvableDeps(ctx, u, unwrappedErr)
		}
		return false, r.handleUpdateFailed(ctx, u, err)
	}
	k8s.EnsureFinalizers(u, k8s.ControllerFinalizerName, k8s.DeletionDefenderFinalizerName)

	// set the etag to an empty string, since IAMPolicy is the authoritative intent, KCC wants to overwrite the underlying policy regardless
	//policy.Spec.Etag = ""

	if !existsAlready {
		if _, err = adapter.Create(ctx); err != nil {
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

// func mapToKey(in *unstructured.Unstructured, out *pb.Key) error {
// 	spec, ok := in.Object["spec"].(map[string]interface{})
// 	if !ok {
// 		return fmt.Errorf("unexpected type for spec")
// 	}

// 	mapping, err := getMapping(out.ProtoReflect().Descriptor())
// 	if err != nil {
// 		return err
// 	}

// 	if err := mapping.Apply(spec, out.ProtoReflect()); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func mapKubeToCloudUnstructured(kube *v1alpha1.APIKeysKeySpec, out *UnstructuredKey) error {
// 	kubeJSON, _ := json.Marshal(kube)

// 	m := make(map[string]interface{})
// 	m["restrictions"] = kube.Restrictions
// 	// m["annotations"] = kube.Annotations
// 	m["displayName"] = kube.DisplayName

// 	out.Data = m

// 	outJSON, _ := json.Marshal(out.Data)

// 	klog.Infof("mapping %v => %v", string(kubeJSON), string(outJSON))
// 	return nil
// }

func init() {
	mappings.Add(&v1alpha1.APIKeysKeySpec{}, &pb.Key{},
		mappings.Simple("displayName"),
		mappings.Simple("restrictions"),
	)
	// TODO: Auto convert reverse
	mappings.Add(&pb.Key{}, &v1alpha1.APIKeysKeySpec{},
		mappings.Simple("displayName"),
		mappings.Simple("restrictions"),
	)
	mappings.Add(&v1alpha1.KeyRestrictions{}, &pb.Restrictions{},
		mappings.Simple("apiTargets"),
	)
	mappings.Add(&pb.Restrictions{}, &v1alpha1.KeyRestrictions{},
		mappings.Simple("apiTargets"),
	)
	mappings.Add(&v1alpha1.KeyApiTargets{}, &pb.ApiTarget{},
		mappings.Simple("methods"),
		mappings.Simple("service"),
	)
	mappings.Add(&pb.ApiTarget{}, &v1alpha1.KeyApiTargets{},
		mappings.Simple("methods"),
		mappings.Simple("service"),
	)
}

// restrictions, ok := spec["restrictions"].(map[string]interface{})
// if !ok {
// 	return fmt.Errorf("unexpected type for spec.restrictions")
// }
// apiTargets, ok := restrictions["apiTargets"].([]interface{})
// if !ok {
// 	return fmt.Errorf("unexpected type for spec.restrictions.apiTargets")
// }
// for _, apiTarget := range apiTargets {
// 	obj, ok := apiTarget.(map[string]interface{})
// 	if !ok {
// 		return fmt.Errorf("unexpected type for spec.restrictions.apiTargets")
// 	}
// 	apiTargetOut := &pb.ApiTarget{}
// 	if err := mapToApiTarget(obj, &apiTargetOut);  err != nil {
// 		return err
// 	}
// 	out.Restrictions =
// 	out :=
// }

// 	projectRef:
//     external: ${projectId}
//   resourceID: apikeyskeybasic${uniqueId}
//   displayName: "Example Display Name"
//   restrictions:
//     apiTargets:
//     - service: "translate.googleapis.com"
//       methods: [ "GET" ]

// }

func (r *reconcileContext) handleUpToDate(ctx context.Context, policy *unstructured.Unstructured) error {
	resource, err := toK8sResource(policy)
	if err != nil {
		return fmt.Errorf("error converting to k8s resource while handling %v event: %w", k8s.UpToDate, err)
	}
	return r.Reconciler.HandleUpToDate(ctx, resource)
}

func (r *reconcileContext) handleUpdateFailed(ctx context.Context, policy *unstructured.Unstructured, origErr error) error {
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

	logger := logger.WithValues(
		"resource", resource.GetNamespacedName(),
		"resourceGVK", resource.GroupVersionKind(),
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

type adapter struct {
	projectID string
	location  string
	keyID     string

	desired *v1alpha1.APIKeysKeySpec
	actual  *v1alpha1.APIKeysKeySpec

	gcp *gcpClient
}

func AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (*adapter, error) {
	// TODO: Get from context?   We want our hooks though...
	httpClient := &http.Client{
		Transport: http.DefaultTransport,
	}

	gcpClient := &gcpClient{
		httpClient: httpClient,
	}

	// TODO: Just fetch this object?
	obj := &v1alpha1.APIKeysKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID := obj.Spec.ProjectRef.External
	if projectID == "" {
		return nil, fmt.Errorf("unable to determine project")
	}

	keyID := ValueOf(obj.Spec.ResourceID)
	if keyID == "" {
		return nil, fmt.Errorf("unable to determine resourceID")
	}

	location := "global"

	return &adapter{
		projectID: projectID,
		location:  location,
		keyID:     keyID,
		desired:   &obj.Spec,
		gcp:       gcpClient,
	}, nil
}

func ValueOf[T any](p *T) T {
	var v T
	if p != nil {
		v = *p
	}
	return v
}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	key, err := a.gcp.get(ctx, a.projectID, a.location, a.keyID)
	if IsNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	u := &v1alpha1.APIKeysKeySpec{}
	if err := mappings.Map(key, u); err != nil {
		return false, err
	}
	a.actual = u

	return true, nil
}

func (a *adapter) Delete(ctx context.Context) error {
	// TODO: Delete from status?
	err := a.gcp.delete(ctx, a.projectID, a.location, a.keyID)
	if err != nil {
		return err
	}

	return nil
}

// todo is a placeholder for functionality not-yet-implemented.
type todo struct {
}

func (a *adapter) BuildCreate(ctx context.Context) (*pb.CreateKeyRequest, error) {
	desired := &pb.Key{}
	if err := mappings.Map(a.desired, desired); err != nil {
		return nil, err
	}
	return &pb.CreateKeyRequest{Key: desired}, nil
}

func (a *adapter) Create(ctx context.Context) (*todo, error) {
	// desired := &UnstructuredKey{}
	// if err := mapKubeToCloud(a.desired, desired); err != nil {
	// 	return nil, err
	// }

	req, err := a.BuildCreate(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := a.gcp.create(ctx, a.projectID, a.location, a.keyID, req.Key); err != nil {
		return nil, err
	}
	// TODO: Return created object
	return nil, nil
}

func (a *adapter) Update(ctx context.Context) (*todo, error) {
	// TODO: Return updated object

	desiredCloud := &pb.Key{}
	if err := mappings.Map(a.desired, desiredCloud); err != nil {
		return nil, err
	}

	// patch := &pb.Key{} //Data: make(map[string]interface{})}
	// if !reflect.DeepEqual(a.desired.DisplayName, a.actual.DisplayName) {
	// 	patch.Data["displayName"] = desiredCloud.["displayName"]
	// }
	// if !reflect.DeepEqual(a.desired.Restrictions, a.actual.Restrictions) {
	// 	patch.Data["restrictions"] = desiredCloud.Data["restrictions"]
	// }

	delta := &v1alpha1.APIKeysKeySpec{} //Data: make(map[string]interface{})}
	diffCount := 0
	if !reflect.DeepEqual(a.desired.DisplayName, a.actual.DisplayName) {
		delta.DisplayName = a.desired.DisplayName
		diffCount++
	}
	if !reflect.DeepEqual(a.desired.Restrictions, a.actual.Restrictions) {
		delta.Restrictions = a.desired.Restrictions
		diffCount++
	}
	klog.Infof("desired: %v", JSON(a.desired))
	klog.Infof("actual: %v", JSON(a.actual))
	klog.Infof("delta: %v", JSON(delta))

	if diffCount == 0 {
		// TODO: Log/warn/error?
		return nil, nil
	}

	patch := &pb.Key{}
	if err := mappings.Map(delta, patch); err != nil {
		return nil, err
	}
	_, err := a.gcp.patch(ctx, a.projectID, a.location, a.keyID, patch)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

type gcpClient struct {
	httpClient *http.Client
}

func (c *gcpClient) buildURL(projectID, location, keyID string) (string, error) {
	url := "https://apikeys.googleapis.com/v2/projects/" + projectID + "/locations/" + location + "/keys"
	if keyID != "" {
		url += "/" + keyID
	}
	return url, nil
}

func (r *gcpClient) delete(ctx context.Context, projectID, location, keyID string) error {
	log := log.FromContext(ctx)

	url, err := r.buildURL(projectID, location, keyID)
	if err != nil {
		return err
	}
	method := "DELETE"
	var body io.Reader

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return fmt.Errorf("error building http request %s %s: %w", method, url, err)
	}

	b, err := r.do(ctx, req)
	if err != nil {
		return err
	}

	// TODO: Parse operation, wait
	log.Info("got response", "body", string(b))
	return nil
}

func (r *gcpClient) do(ctx context.Context, req *http.Request) ([]byte, error) {
	method := req.Method
	url := req.URL.String()

	httpClient := r.httpClient
	httpClient = transport_tpg.DefaultHTTPClientTransformer(ctx, httpClient)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error from http request %s %s: %w", method, url, err)
	}
	defer resp.Body.Close()
	if err := googleapi.CheckResponse(resp); err != nil {
		return nil, err
		// return nil, fmt.Errorf("unexpected result from http request %s %s: %v", method, url, resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading http response from %s %s: %w", method, url, err)
	}

	return b, nil
}
func (r *gcpClient) get(ctx context.Context, projectID, location, keyID string) (*pb.Key, error) {
	// log := log.FromContext(ctx)

	url, err := r.buildURL(projectID, location, keyID)
	if err != nil {
		return nil, err
	}
	method := "GET"
	var body io.Reader

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("error building http request %s %s: %w", method, url, err)
	}
	b, err := r.do(ctx, req)
	if err != nil {
		return nil, err
	}

	// m := make(map[string]interface{})
	// if err := json.Unmarshal(b, &m); err != nil {
	// 	return nil, fmt.Errorf("error parsing http response from %s %s: %w", method, url, err)
	// }
	// return &Key{Data: m}, nil

	k := &pb.Key{}
	if err := protojson.Unmarshal(b, k); err != nil {
		return nil, fmt.Errorf("error parsing http response from %s %s: %w", method, url, err)
	}
	return k, nil
}

func (r *gcpClient) create(ctx context.Context, projectID, location, keyID string, key *pb.Key) (*string, error) {
	log := log.FromContext(ctx)

	url, err := r.buildURL(projectID, location, "")
	if err != nil {
		return nil, err
	}
	if keyID != "" {
		// TODO: encoding etc
		url += "?keyId=" + keyID
	}
	method := "POST"

	reqBody, err := protojson.Marshal(key)
	// reqBody, err := json.Marshal(key.Data)
	if err != nil {
		return nil, fmt.Errorf("error building request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error building http request %s %s: %w", method, url, err)
	}
	respBody, err := r.do(ctx, req)
	if err != nil {
		return nil, err
	}

	// TODO: Parse operation, wait
	log.Info("got response", "body", string(respBody))
	return nil, nil
}

func (r *gcpClient) patch(ctx context.Context, projectID, location, keyID string, key *pb.Key) (*string, error) {
	log := log.FromContext(ctx)

	url, err := r.buildURL(projectID, location, keyID)
	if err != nil {
		return nil, err
	}
	method := "PATCH"

	reqBody, err := protojson.Marshal(key)
	// reqBody, err := json.Marshal(key.Data)
	if err != nil {
		return nil, fmt.Errorf("error building request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error building http request %s %s: %w", method, url, err)
	}
	respBody, err := r.do(ctx, req)
	if err != nil {
		return nil, err
	}

	// TODO: Parse operation, wait
	log.Info("got response", "body", string(respBody))
	return nil, nil
}

// IsNotFound reports whether err is the result of the
// server replying with http.StatusNotFound.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	ae, ok := err.(*googleapi.Error)
	return ok && ae.Code == http.StatusNotFound
}

// DeferredJSON is a helper that delays JSON formatting until/unless it is needed.
type DeferredJSON struct {
	O interface{}
}

// String is the method that is called to format the object.
func (d DeferredJSON) String() string {
	b, err := json.Marshal(d.O)
	if err != nil {
		return fmt.Sprintf("<error: %v>", err)
	}
	return string(b)
}

// JSON is a helper that prints the object in JSON format.
// We use lazy-evaluation to avoid calling json.Marshal unless it is actually needed.
func JSON(o interface{}) DeferredJSON {
	return DeferredJSON{o}
}

// func mapCloudToKubeUnstructured(in *UnstructuredKey) (*v1alpha1.APIKeysKeySpec, error) {
// 	kube := &v1alpha1.APIKeysKeySpec{}

// 	if v, found := in.Data["restrictions"]; found {
// 		kube.Restrictions = &v1alpha1.KeyRestrictions{}
// 		vMap, ok := v.(map[string]interface{})
// 		if !ok {
// 			return nil, fmt.Errorf("unexpected type for restrictions, got %T, want object", v)
// 		}
// 		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(vMap, &kube.Restrictions); err != nil {
// 			return nil, err
// 		}
// 	}

// 	if v, found := in.Data["displayName"]; found {
// 		vString, ok := v.(string)
// 		if !ok {
// 			return nil, fmt.Errorf("unexpected type for restrictions, got %T, want string", v)
// 		}
// 		kube.DisplayName = &vString
// 	}

// 	return kube, nil
// }
