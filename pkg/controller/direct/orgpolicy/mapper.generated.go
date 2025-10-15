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

// +generated:mapper
// krm.group: orgpolicy.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.orgpolicy.v2

package orgpolicy

import (
	pb "cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"
	krmorgpolicyv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/orgpolicy/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/orgpolicy/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AlternatePolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.AlternatePolicySpec) *krmorgpolicyv1alpha1.AlternatePolicySpec {
	if in == nil {
		return nil
	}
	out := &krmorgpolicyv1alpha1.AlternatePolicySpec{}
	out.Launch = direct.LazyPtr(in.GetLaunch())
	out.Spec = PolicySpec_FromProto(mapCtx, in.GetSpec())
	return out
}
func AlternatePolicySpec_ToProto(mapCtx *direct.MapContext, in *krmorgpolicyv1alpha1.AlternatePolicySpec) *pb.AlternatePolicySpec {
	if in == nil {
		return nil
	}
	out := &pb.AlternatePolicySpec{}
	out.Launch = direct.ValueOf(in.Launch)
	out.Spec = PolicySpec_ToProto(mapCtx, in.Spec)
	return out
}
func OrgPolicyCustomConstraintObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomConstraint) *krm.OrgPolicyCustomConstraintObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrgPolicyCustomConstraintObservedState{}
	// MISSING: Name
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func OrgPolicyCustomConstraintObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrgPolicyCustomConstraintObservedState) *pb.CustomConstraint {
	if in == nil {
		return nil
	}
	out := &pb.CustomConstraint{}
	// MISSING: Name
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func OrgPolicyCustomConstraintSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomConstraint) *krm.OrgPolicyCustomConstraintSpec {
	if in == nil {
		return nil
	}
	out := &krm.OrgPolicyCustomConstraintSpec{}
	// MISSING: Name
	out.ResourceTypes = in.ResourceTypes
	out.MethodTypes = direct.EnumSlice_FromProto(mapCtx, in.MethodTypes)
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.ActionType = direct.Enum_FromProto(mapCtx, in.GetActionType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func OrgPolicyCustomConstraintSpec_ToProto(mapCtx *direct.MapContext, in *krm.OrgPolicyCustomConstraintSpec) *pb.CustomConstraint {
	if in == nil {
		return nil
	}
	out := &pb.CustomConstraint{}
	// MISSING: Name
	out.ResourceTypes = in.ResourceTypes
	out.MethodTypes = direct.EnumSlice_ToProto[pb.CustomConstraint_MethodType](mapCtx, in.MethodTypes)
	out.Condition = direct.ValueOf(in.Condition)
	out.ActionType = direct.Enum_ToProto[pb.CustomConstraint_ActionType](mapCtx, in.ActionType)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func OrgPolicyPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krmorgpolicyv1alpha1.OrgPolicyPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krmorgpolicyv1alpha1.OrgPolicyPolicyObservedState{}
	// MISSING: Name
	out.Spec = PolicySpecObservedState_FromProto(mapCtx, in.GetSpec())
	// MISSING: Alternate
	out.DryRunSpec = PolicySpecObservedState_FromProto(mapCtx, in.GetDryRunSpec())
	// MISSING: Etag
	return out
}
func OrgPolicyPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krmorgpolicyv1alpha1.OrgPolicyPolicyObservedState) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	// MISSING: Name
	out.Spec = PolicySpecObservedState_ToProto(mapCtx, in.Spec)
	// MISSING: Alternate
	out.DryRunSpec = PolicySpecObservedState_ToProto(mapCtx, in.DryRunSpec)
	// MISSING: Etag
	return out
}
func OrgPolicyPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krmorgpolicyv1alpha1.OrgPolicyPolicySpec {
	if in == nil {
		return nil
	}
	out := &krmorgpolicyv1alpha1.OrgPolicyPolicySpec{}
	// MISSING: Name
	out.Spec = PolicySpec_FromProto(mapCtx, in.GetSpec())
	// MISSING: Alternate
	out.DryRunSpec = PolicySpec_FromProto(mapCtx, in.GetDryRunSpec())
	// MISSING: Etag
	return out
}
func OrgPolicyPolicySpec_ToProto(mapCtx *direct.MapContext, in *krmorgpolicyv1alpha1.OrgPolicyPolicySpec) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	// MISSING: Name
	out.Spec = PolicySpec_ToProto(mapCtx, in.Spec)
	// MISSING: Alternate
	out.DryRunSpec = PolicySpec_ToProto(mapCtx, in.DryRunSpec)
	// MISSING: Etag
	return out
}
func PolicySpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PolicySpec) *krmorgpolicyv1alpha1.PolicySpecObservedState {
	if in == nil {
		return nil
	}
	out := &krmorgpolicyv1alpha1.PolicySpecObservedState{}
	// MISSING: Etag
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Rules
	// MISSING: InheritFromParent
	// MISSING: Reset
	return out
}
func PolicySpecObservedState_ToProto(mapCtx *direct.MapContext, in *krmorgpolicyv1alpha1.PolicySpecObservedState) *pb.PolicySpec {
	if in == nil {
		return nil
	}
	out := &pb.PolicySpec{}
	// MISSING: Etag
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Rules
	// MISSING: InheritFromParent
	// MISSING: Reset
	return out
}
func PolicySpec_PolicyRule_FromProto(mapCtx *direct.MapContext, in *pb.PolicySpec_PolicyRule) *krmorgpolicyv1alpha1.PolicySpec_PolicyRule {
	if in == nil {
		return nil
	}
	out := &krmorgpolicyv1alpha1.PolicySpec_PolicyRule{}
	out.Values = PolicySpec_PolicyRule_StringValues_FromProto(mapCtx, in.GetValues())
	out.AllowAll = direct.LazyPtr(in.GetAllowAll())
	out.DenyAll = direct.LazyPtr(in.GetDenyAll())
	out.Enforce = direct.LazyPtr(in.GetEnforce())
	out.Condition = Expr_FromProto(mapCtx, in.GetCondition())
	out.Parameters = Parameters_FromProto(mapCtx, in.GetParameters())
	return out
}
func PolicySpec_PolicyRule_ToProto(mapCtx *direct.MapContext, in *krmorgpolicyv1alpha1.PolicySpec_PolicyRule) *pb.PolicySpec_PolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.PolicySpec_PolicyRule{}
	if oneof := PolicySpec_PolicyRule_StringValues_ToProto(mapCtx, in.Values); oneof != nil {
		out.Kind = &pb.PolicySpec_PolicyRule_Values{Values: oneof}
	}
	if oneof := PolicySpec_PolicyRule_AllowAll_ToProto(mapCtx, in.AllowAll); oneof != nil {
		out.Kind = oneof
	}
	if oneof := PolicySpec_PolicyRule_DenyAll_ToProto(mapCtx, in.DenyAll); oneof != nil {
		out.Kind = oneof
	}
	if oneof := PolicySpec_PolicyRule_Enforce_ToProto(mapCtx, in.Enforce); oneof != nil {
		out.Kind = oneof
	}
	out.Condition = Expr_ToProto(mapCtx, in.Condition)
	out.Parameters = Parameters_ToProto(mapCtx, in.Parameters)
	return out
}
func PolicySpec_PolicyRule_AllowAll_ToProto(mapCtx *direct.MapContext, in *bool) *pb.PolicySpec_PolicyRule_AllowAll {
	if in == nil {
		return nil
	}
	if !*in {
		return nil
	}
	return &pb.PolicySpec_PolicyRule_AllowAll{AllowAll: *in}
}
func PolicySpec_PolicyRule_DenyAll_ToProto(mapCtx *direct.MapContext, in *bool) *pb.PolicySpec_PolicyRule_DenyAll {
	if in == nil {
		return nil
	}
	if !*in {
		return nil
	}
	return &pb.PolicySpec_PolicyRule_DenyAll{DenyAll: *in}
}
func PolicySpec_PolicyRule_Enforce_ToProto(mapCtx *direct.MapContext, in *bool) *pb.PolicySpec_PolicyRule_Enforce {
	if in == nil {
		return nil
	}
	if !*in {
		return nil
	}
	return &pb.PolicySpec_PolicyRule_Enforce{Enforce: *in}
}
func PolicySpec_PolicyRule_StringValues_FromProto(mapCtx *direct.MapContext, in *pb.PolicySpec_PolicyRule_StringValues) *krmorgpolicyv1alpha1.PolicySpec_PolicyRule_StringValues {
	if in == nil {
		return nil
	}
	out := &krmorgpolicyv1alpha1.PolicySpec_PolicyRule_StringValues{}
	out.AllowedValues = in.AllowedValues
	out.DeniedValues = in.DeniedValues
	return out
}
func PolicySpec_PolicyRule_StringValues_ToProto(mapCtx *direct.MapContext, in *krmorgpolicyv1alpha1.PolicySpec_PolicyRule_StringValues) *pb.PolicySpec_PolicyRule_StringValues {
	if in == nil {
		return nil
	}
	out := &pb.PolicySpec_PolicyRule_StringValues{}
	out.AllowedValues = in.AllowedValues
	out.DeniedValues = in.DeniedValues
	return out
}
