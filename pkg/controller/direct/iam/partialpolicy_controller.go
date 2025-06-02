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

package iam

import (
	"context"
	"fmt"
	"reflect"

	"cloud.google.com/go/iam/apiv1/iampb"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	newiamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	oldiamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/go-logr/logr"
)

const (
	iamPartialPolicyControllerName = "iampartialpolicy-controller"
)

func init() {
	registry.RegisterModel(newiamv1beta1.IAMPartialPolicyGVK, NewIAMPartialPolicyModel)
}

func NewIAMPartialPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelIAMPartialPolicy{config: *config}, nil
}

var _ directbase.Model = &modelIAMPartialPolicy{}

type modelIAMPartialPolicy struct {
	config config.ControllerConfig
}

func validateDeps(deps *directbase.IAMAdapterDeps) error {
	if deps == nil {
		return fmt.Errorf("IAMAdapterDeps is nil")
	}

	if deps.KubeClient == nil {
		return fmt.Errorf("KubeClient is nil")
	}
	if deps.ControllerDeps == nil {
		return fmt.Errorf("ControllerDeps is nil")
	}

	if deps.ControllerDeps.TfProvider == nil {
		return fmt.Errorf("TfProvider is nil")
	}
	if deps.ControllerDeps.TfLoader == nil {
		return fmt.Errorf("TfLoader is nil")
	}
	if deps.ControllerDeps.DclConverter == nil {
		return fmt.Errorf("DclConverter is nil")
	}

	return nil
}
func (m *modelIAMPartialPolicy) IAMAdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured, deps *directbase.IAMAdapterDeps) (directbase.Adapter, error) {
	obj := &newiamv1beta1.IAMPartialPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := validateDeps(deps); err != nil {
		return nil, fmt.Errorf("error validating dependencies: %w", err)
	}

	iamClient := kcciamclient.New(deps.ControllerDeps.TfProvider, deps.ControllerDeps.TfLoader, deps.KubeClient, deps.ControllerDeps.DclConverter, deps.ControllerDeps.DclConfig)
	return &IAMPartialPolicyAdapter{
		iamClient: iamClient,
		desired:   obj,
	}, nil
}

func (m *modelIAMPartialPolicy) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	return nil, fmt.Errorf("AdapterForObject not supported for IAMPartialPolicy, call IAMAdapterForObject")
}

func (m *modelIAMPartialPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// IAMPartialPolicy is not identified or managed via a direct GCP URL in the same way
	// other resources are.
	return nil, fmt.Errorf("AdapterForURL not supported for IAMPartialPolicy")
}

type IAMPartialPolicyAdapter struct {
	iamClient                      *kcciamclient.IAMClient
	desired                        *newiamv1beta1.IAMPartialPolicy
	actualReferencedResourcePolicy *iampb.Policy
}

var _ directbase.Adapter = &IAMPartialPolicyAdapter{}

func getLogger(ctx context.Context) logr.Logger {
	return klog.FromContext(ctx).WithName(iamPartialPolicyControllerName).WithValues("controllerType", "direct")
}

func (a *IAMPartialPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := getLogger(ctx)
	log.V(2).Info("getting IAM policy for resource", "kind", a.desired.Spec.ResourceReference.Kind, "name", a.desired.Spec.ResourceReference.Name, "namespace", a.desired.Spec.ResourceReference.Namespace, "external", a.desired.Spec.ResourceReference.External)
	mapCtx := &direct.MapContext{}

	iamPolicySkeleton := ToOldIAMPolicySkeleton(a.desired)
	oldKRMPolicy, err := a.iamClient.GetPolicy(ctx, iamPolicySkeleton)
	if err != nil {
		if apierrors.IsNotFound(err) || k8s.IsReferenceNotFoundError(err) {
			log.V(2).Info("IAM policy not found or underlying resource not found", "resource", k8s.GetNamespacedName(a.desired))
			return false, nil
		}
		return false, fmt.Errorf("getting IAM policy for %v: %w", k8s.GetNamespacedName(a.desired), err)
	}

	a.actualReferencedResourcePolicy = IAMPolicySpec_ToProto(mapCtx, &oldKRMPolicy.Spec)
	return true, nil
}

// IAMMemberIdentityResolver helps to resolve referenced member identity
type IAMMemberIdentityResolver struct {
	Iamclient *kcciamclient.IAMClient
	Ctx       context.Context
}

func (t IAMMemberIdentityResolver) Resolve(member newiamv1beta1.Member, memberFrom *newiamv1beta1.MemberSource, defaultNamespace string) (string, error) {
	oldMember := oldiamv1beta1.Member(member)
	var oldMemberFrom *oldiamv1beta1.MemberSource
	if memberFrom != nil {
		oldMemberFrom = &oldiamv1beta1.MemberSource{}
		if memberFrom.BigQueryConnectionConnectionRef != nil {
			oldMemberFrom.BigQueryConnectionConnectionRef = memberFrom.BigQueryConnectionConnectionRef
		}
		if memberFrom.ServiceAccountRef != nil {
			oldMemberFrom.ServiceAccountRef = &oldiamv1beta1.MemberReference{
				Name:      memberFrom.ServiceAccountRef.Name,
				Namespace: memberFrom.ServiceAccountRef.Namespace,
			}
		}
		if memberFrom.LogSinkRef != nil {
			oldMemberFrom.LogSinkRef = &oldiamv1beta1.MemberReference{
				Name:      memberFrom.LogSinkRef.Name,
				Namespace: memberFrom.LogSinkRef.Namespace,
			}
		}
		if memberFrom.SQLInstanceRef != nil {
			oldMemberFrom.SQLInstanceRef = &oldiamv1beta1.MemberReference{
				Name:      memberFrom.SQLInstanceRef.Name,
				Namespace: memberFrom.SQLInstanceRef.Namespace,
			}
		}
		if memberFrom.ServiceIdentityRef != nil {
			oldMemberFrom.ServiceIdentityRef = &oldiamv1beta1.MemberReference{
				Name:      memberFrom.ServiceIdentityRef.Name,
				Namespace: memberFrom.ServiceIdentityRef.Namespace,
			}
		}

	}

	return kcciamclient.ResolveMemberIdentity(t.Ctx, oldMember, oldMemberFrom, defaultNamespace, t.Iamclient.TFIAMClient)
}

func (a *IAMPartialPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := getLogger(ctx)
	log.V(2).Info("creating/applying IAMPartialPolicy", "name", k8s.GetNamespacedName(a.desired))

	var livePolicyForMerge *newiamv1beta1.IAMPolicy

	// If the policy truly doesn't exist, livePolicyForMerge should be an empty IAMPolicy struct
	// for ComputePartialPolicyWithMergedBindings to work correctly.
	if a.actualReferencedResourcePolicy == nil {
		livePolicyForMerge = &newiamv1beta1.IAMPolicy{}
	} else {
		// todo acpana round trip ?
		livePolicyForMerge = ToNewIAMPolicySkeleton(a.desired)
	}

	resolver := IAMMemberIdentityResolver{Iamclient: a.iamClient, Ctx: ctx}
	desiredPartialPolicyWithStatus, err := ComputePartialPolicyWithMergedBindings(a.desired, livePolicyForMerge, resolver)
	if err != nil {
		return fmt.Errorf("computing partial policy for create for %v: %w", k8s.GetNamespacedName(a.desired), err)
	}

	finalPolicyToSet := toDesiredPolicy(desiredPartialPolicyWithStatus, livePolicyForMerge)

	_, err = a.iamClient.SetPolicy(ctx, finalPolicyToSet)
	if err != nil {
		return fmt.Errorf("setting IAM policy for %v: %w", k8s.GetNamespacedName(a.desired), err)
	}

	log.V(2).Info("successfully applied IAMPartialPolicy for create", "name", k8s.GetNamespacedName(a.desired))

	// Update KRM status
	newKRMStatus := newiamv1beta1.IAMPartialPolicyStatus{
		ObservedGeneration:  a.desired.GetGeneration(),
		LastAppliedBindings: desiredPartialPolicyWithStatus.Status.LastAppliedBindings,
		AllBindings:         desiredPartialPolicyWithStatus.Status.AllBindings,
	}
	readyCondition := k8s.NewCustomReadyCondition(corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage)
	if err := createOp.UpdateStatus(ctx, &newKRMStatus, &readyCondition); err != nil {
		return fmt.Errorf("updating status for %v: %w", k8s.GetNamespacedName(a.desired), err)
	}
	return nil
}

func (a *IAMPartialPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := getLogger(ctx)
	log.V(2).Info("updating/applying IAMPartialPolicy", "name", k8s.GetNamespacedName(a.desired))
	mapCtx := &direct.MapContext{}

	if a.actualReferencedResourcePolicy == nil {
		return fmt.Errorf("no IAM policy found for referenced resource %s during update for IAMPartialPolicy %v; requeueing for Find", a.desired.Spec.ResourceReference.GetNamespacedName(), k8s.GetNamespacedName(a.desired))
	}

	livePolicyForMerge := &newiamv1beta1.IAMPolicy{}
	livePolicyForMerge.Spec = *IAMPolicySpec_FromProto(mapCtx, a.actualReferencedResourcePolicy)

	resolver := IAMMemberIdentityResolver{Iamclient: a.iamClient, Ctx: ctx}

	desiredPartialPolicyWithStatus, err := ComputePartialPolicyWithMergedBindings(a.desired, livePolicyForMerge, resolver)
	if err != nil {
		return fmt.Errorf("computing partial policy for update for %v: %w", k8s.GetNamespacedName(a.desired), err)
	}

	finalPolicyToSet := toDesiredPolicy(desiredPartialPolicyWithStatus, livePolicyForMerge)

	// todo acpana better comparison, etag?
	gcpUpdateNeeded := !reflect.DeepEqual(livePolicyForMerge.Spec.Bindings, finalPolicyToSet.Spec.Bindings) ||
		!reflect.DeepEqual(livePolicyForMerge.Spec.AuditConfigs, finalPolicyToSet.Spec.AuditConfigs)

	if gcpUpdateNeeded {
		log.V(2).Info("GCP IAM policy change detected, calling SetPolicy", "name", k8s.GetNamespacedName(a.desired))
		_, err = a.iamClient.SetPolicy(ctx, finalPolicyToSet)
		if err != nil {
			return fmt.Errorf("setting IAM policy for %v: %w", k8s.GetNamespacedName(a.desired), err)
		}
		log.V(2).Info("successfully applied IAMPartialPolicy for update", "name", k8s.GetNamespacedName(a.desired))
	} else {
		log.V(2).Info("no change in GCP IAM policy needed", "name", k8s.GetNamespacedName(a.desired))
	}

	// Always update KRM status to reflect observed generation and latest computed bindings,
	// even if GCP didn't change, as spec or other KRM details might have.
	newKRMStatus := newiamv1beta1.IAMPartialPolicyStatus{
		ObservedGeneration:  a.desired.GetGeneration(),
		LastAppliedBindings: desiredPartialPolicyWithStatus.Status.LastAppliedBindings,
		AllBindings:         desiredPartialPolicyWithStatus.Status.AllBindings,
	}
	readyCondition := k8s.NewCustomReadyCondition(corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage)
	if err := updateOp.UpdateStatus(ctx, &newKRMStatus, &readyCondition); err != nil {
		return fmt.Errorf("updating status for %v: %w", k8s.GetNamespacedName(a.desired), err)
	}
	return nil
}

func (a *IAMPartialPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := getLogger(ctx)
	log.V(2).Info("deleting/finalizing IAMPartialPolicy", "name", k8s.GetNamespacedName(a.desired))
	mapCtx := &direct.MapContext{}

	if a.actualReferencedResourcePolicy == nil {
		// If the policy or resource is already gone, there's nothing to prune from GCP.
		log.V(2).Info("actual policy not found during delete, assuming already gone or no-op needed for GCP", "name", k8s.GetNamespacedName(a.desired))
		return true, nil
	}
	livePolicyForPruning := &newiamv1beta1.IAMPolicy{}
	livePolicyForPruning.Spec = direct.ValueOf(IAMPolicySpec_FromProto(mapCtx, a.actualReferencedResourcePolicy))

	desiredPartialPolicyWithRemainingStatus := ComputePartialPolicyWithRemainingBindings(a.desired, livePolicyForPruning)
	finalPolicyToSet := toDesiredPolicy(desiredPartialPolicyWithRemainingStatus, livePolicyForPruning)

	_, err := a.iamClient.SetPolicy(ctx, finalPolicyToSet)
	if err != nil {
		if apierrors.IsNotFound(err) || k8s.IsReferenceNotFoundError(err) {
			log.V(2).Info("IAM policy or underlying resource not found during SetPolicy for deletion, treated as success", "name", k8s.GetNamespacedName(a.desired))
			return true, nil // Policy/resource gone, so our bindings are effectively removed
		}
		return false, fmt.Errorf("setting IAM policy for deletion of %v: %w", k8s.GetNamespacedName(a.desired), err)
	}

	log.V(2).Info("successfully deleted IAMPartialPolicy", "name", k8s.GetNamespacedName(a.desired))
	return true, nil
}

func (a *IAMPartialPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

// ToNewIAMPolicySkeleton creates an IAMPolicy struct with ObjectMeta and resource reference
// copied from the partial policy. The skeleton struct can be passed to IAMClient.GetPolicy()
// to fetch the live IAM policy.
func ToNewIAMPolicySkeleton(p *newiamv1beta1.IAMPartialPolicy) *newiamv1beta1.IAMPolicy {
	res := &newiamv1beta1.IAMPolicy{
		TypeMeta: metav1.TypeMeta{
			Kind:       newiamv1beta1.IAMPolicyGVK.Kind,
			APIVersion: newiamv1beta1.IAMAPIVersion,
		},
	}

	res.ObjectMeta = *p.ObjectMeta.DeepCopy()
	res.Spec.ResourceReference.APIVersion = p.Spec.ResourceReference.APIVersion
	res.Spec.ResourceReference.Kind = p.Spec.ResourceReference.Kind
	res.Spec.ResourceReference.Name = p.Spec.ResourceReference.Name
	res.Spec.ResourceReference.Namespace = p.Spec.ResourceReference.Namespace
	res.Spec.ResourceReference.External = p.Spec.ResourceReference.External

	return res
}

// old style v1beta1 Policy representation for the IAM Client
func ToOldIAMPolicySkeleton(p *newiamv1beta1.IAMPartialPolicy) *oldiamv1beta1.IAMPolicy {
	res := &oldiamv1beta1.IAMPolicy{
		TypeMeta: metav1.TypeMeta{
			Kind:       newiamv1beta1.IAMPolicyGVK.Kind,
			APIVersion: newiamv1beta1.IAMAPIVersion,
		},
	}

	res.ObjectMeta = *p.ObjectMeta.DeepCopy()
	res.Spec.ResourceReference.APIVersion = p.Spec.ResourceReference.APIVersion
	res.Spec.ResourceReference.Kind = p.Spec.ResourceReference.Kind
	res.Spec.ResourceReference.Name = p.Spec.ResourceReference.Name
	res.Spec.ResourceReference.Namespace = p.Spec.ResourceReference.Namespace
	res.Spec.ResourceReference.External = p.Spec.ResourceReference.External

	return res
}

func toDesiredPolicy(desiredPartialPolicy *newiamv1beta1.IAMPartialPolicy, livePolicy *newiamv1beta1.IAMPolicy) *oldiamv1beta1.IAMPolicy {
	desiredPolicy := ToOldIAMPolicySkeleton(desiredPartialPolicy)

	if len(desiredPartialPolicy.Status.AllBindings) > 0 {
		desiredPolicy.Spec.Bindings = make([]oldiamv1beta1.IAMPolicyBinding, len(desiredPartialPolicy.Status.AllBindings))
		for i, binding := range desiredPartialPolicy.Status.AllBindings {
			desiredPolicy.Spec.Bindings[i] = oldiamv1beta1.IAMPolicyBinding{
				Role: binding.Role,
			}
			desiredPolicy.Spec.Bindings[i].Members = make([]oldiamv1beta1.Member, len(binding.Members))
			for j, member := range binding.Members {
				desiredPolicy.Spec.Bindings[i].Members[j] = oldiamv1beta1.Member(member)
			}
			if binding.Condition != nil {
				desiredPolicy.Spec.Bindings[i].Condition = &oldiamv1beta1.IAMCondition{
					Description: binding.Condition.Description,
					Expression:  binding.Condition.Expression,
					Title:       binding.Condition.Title,
				}
			}
		}
	}
	// Carry the current etag from read to support concurrent read-modify-write operations from multiple systems.
	// SetPolicy will fail if the policy has been modified by other actors since the controller retrieved it.
	desiredPolicy.Spec.Etag = livePolicy.Spec.Etag
	// Preserve the audit configs if any.
	if len(livePolicy.Spec.AuditConfigs) > 0 {
		desiredPolicy.Spec.AuditConfigs = make([]oldiamv1beta1.IAMPolicyAuditConfig, len(livePolicy.Spec.AuditConfigs))
		for i, auditConfig := range livePolicy.Spec.AuditConfigs {
			desiredPolicy.Spec.AuditConfigs[i] = oldiamv1beta1.IAMPolicyAuditConfig{
				Service: auditConfig.Service,
			}
			if len(auditConfig.AuditLogConfigs) > 0 {
				desiredPolicy.Spec.AuditConfigs[i].AuditLogConfigs = make([]oldiamv1beta1.AuditLogConfig, len(auditConfig.AuditLogConfigs))
				for j, auditLogConfig := range auditConfig.AuditLogConfigs {
					desiredPolicy.Spec.AuditConfigs[i].AuditLogConfigs[j] = oldiamv1beta1.AuditLogConfig{
						LogType: auditLogConfig.LogType,
					}
					if len(auditLogConfig.ExemptedMembers) > 0 {
						desiredPolicy.Spec.AuditConfigs[i].AuditLogConfigs[j].ExemptedMembers = make([]oldiamv1beta1.Member, len(auditLogConfig.ExemptedMembers))
						for k, member := range auditLogConfig.ExemptedMembers {
							desiredPolicy.Spec.AuditConfigs[i].AuditLogConfigs[j].ExemptedMembers[k] = oldiamv1beta1.Member(member)
						}
					}
				}
			}
		}
	}
	return desiredPolicy
}
