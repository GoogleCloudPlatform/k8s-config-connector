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

	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	condition "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/metrics"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	klog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const controllerName = "iampartialpolicy-controller"

var logger = klog.Log.WithName(controllerName)

// Add creates a new IAM Partial Policy Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and start it when the Manager is started.
func Add(mgr manager.Manager, tfProvider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader,
	converter *conversion.Converter, dclConfig *mmdcl.Config) error {
	reconciler, err := NewReconciler(mgr, tfProvider, smLoader, converter, dclConfig)
	if err != nil {
		return err
	}
	return add(mgr, reconciler)
}

// NewReconciler returns a new reconcile.Reconciler.
func NewReconciler(mgr manager.Manager, provider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader,
	converter *conversion.Converter, dclConfig *mmdcl.Config) (*ReconcileIAMPartialPolicy, error) {
	r := ReconcileIAMPartialPolicy{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(
			mgr.GetClient(),
			mgr.GetEventRecorderFor(controllerName),
		),
		Client:    mgr.GetClient(),
		iamClient: iamclient.New(provider, smLoader, mgr.GetClient(), converter, dclConfig),
		scheme:    mgr.GetScheme(),
		ReconcilerMetrics: metrics.ReconcilerMetrics{
			ResourceNameLabel: metrics.ResourceNameLabel,
		},
	}
	return &r, nil
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler.
func add(mgr manager.Manager, r *ReconcileIAMPartialPolicy) error {
	obj := &iamv1beta1.IAMPartialPolicy{}
	_, err := builder.
		ControllerManagedBy(mgr).
		Named(controllerName).
		WithOptions(controller.Options{MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles, RateLimiter: ratelimiter.NewRateLimiter()}).
		For(obj, builder.OnlyMetadata, builder.WithPredicates(predicate.UnderlyingResourceOutOfSyncPredicate{})).
		Build(r)
	if err != nil {
		return fmt.Errorf("error creating new controller: %v", err)
	}
	return nil
}

var _ reconcile.Reconciler = &ReconcileIAMPartialPolicy{}

// ReconcileIAMPartialPolicy is a reconciler for handling IAM partial policies.
type ReconcileIAMPartialPolicy struct {
	lifecyclehandler.LifecycleHandler
	client.Client
	metrics.ReconcilerMetrics
	iamClient *kcciamclient.IAMClient
	scheme    *runtime.Scheme
}

type reconcileContext struct {
	Reconciler     *ReconcileIAMPartialPolicy
	Ctx            context.Context
	NamespacedName types.NamespacedName
}

func (r *ReconcileIAMPartialPolicy) Reconcile(ctx context.Context, request reconcile.Request) (result reconcile.Result, err error) {
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
	jitteredPeriod := jitter.GenerateJitteredReenqueuePeriod()
	logger.Info("successfully finished reconcile", "resource", request.NamespacedName, "time to next reconciliation", jitteredPeriod)
	return reconcile.Result{RequeueAfter: jitteredPeriod}, nil
}

func (r *reconcileContext) doReconcile(pp *iamv1beta1.IAMPartialPolicy) (requeue bool, err error) {
	defer execution.RecoverWithInternalError(&err)
	if !pp.DeletionTimestamp.IsZero() {
		return r.finalizeDeletion(pp)
	}
	iamPolicy := ToIAMPolicySkeleton(pp)
	if iamPolicy, err = r.Reconciler.iamClient.GetPolicy(r.Ctx, iamPolicy); err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(pp))
			return true, r.handleUnresolvableDeps(pp, unwrappedErr)
		}
		return false, r.handleUpdateFailed(pp, err)
	}
	k8s.EnsureFinalizers(pp, k8s.ControllerFinalizerName, k8s.DeletionDefenderFinalizerName)

	resolver := IAMMemberIdentityResolver{Iamclient: r.Reconciler.iamClient, Ctx: r.Ctx}
	desiredPartialPolicy, err := ComputePartialPolicyWithMergedBindings(pp, iamPolicy, &resolver)
	if err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(pp))
			return true, r.handleUnresolvableDeps(pp, unwrappedErr)
		}
		return false, r.handleUpdateFailed(pp, fmt.Errorf("error computing partial policy: %w", err))
	}
	desiredPolicy := toDesiredPolicy(desiredPartialPolicy, iamPolicy)
	if _, err = r.Reconciler.iamClient.SetPolicy(r.Ctx, desiredPolicy); err != nil {
		if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
			logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(pp))
			return true, r.handleUnresolvableDeps(pp, unwrappedErr)
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
			if !errors.Is(err, kcciamclient.NotFoundError) && !k8s.IsReferenceNotFoundError(err) {
				if unwrappedErr, ok := lifecyclehandler.CausedByUnresolvableDeps(err); ok {
					logger.Info(unwrappedErr.Error(), "resource", k8s.GetNamespacedName(pp))
					return true, r.handleUnresolvableDeps(pp, unwrappedErr)
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
				return true, r.handleUnresolvableDeps(pp, unwrappedErr)
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
		return fmt.Errorf(k8s.UpdateFailedMessageTmpl, origErr)
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

func (r *reconcileContext) handleUnresolvableDeps(policy *iamv1beta1.IAMPartialPolicy, origErr error) error {
	resource, err := toK8sResource(policy)
	if err != nil {
		return fmt.Errorf("error converting IAMPartialPolicy to k8s resource while handling unresolvable dependencies event: %w", err)
	}
	return r.Reconciler.HandleUnresolvableDeps(r.Ctx, resource, origErr)
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
	iamclient.SetGVK(policy)
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
