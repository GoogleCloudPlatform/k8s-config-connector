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

package orgpolicy

import (
	pb "cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/orgpolicy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	expr "google.golang.org/genproto/googleapis/type/expr"
)

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
	out.ResourceTypes = in.ResourceTypes
	out.MethodTypes = direct.EnumSlice_ToProto[pb.CustomConstraint_MethodType](mapCtx, in.MethodTypes)
	out.Condition = direct.ValueOf(in.Condition)
	out.ActionType = direct.Enum_ToProto[pb.CustomConstraint_ActionType](mapCtx, in.ActionType)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func OrgPolicyPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.OrgPolicyPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrgPolicyPolicyObservedState{}
	return out
}
func OrgPolicyPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrgPolicyPolicyObservedState) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	return out
}
func OrgPolicyPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.OrgPolicyPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.OrgPolicyPolicySpec{}
	out.Spec = PolicySpec_FromProto(mapCtx, in.GetSpec())
	out.DryRunSpec = PolicySpec_FromProto(mapCtx, in.GetDryRunSpec())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func OrgPolicyPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.OrgPolicyPolicySpec) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	out.Spec = PolicySpec_ToProto(mapCtx, in.Spec)
	out.DryRunSpec = PolicySpec_ToProto(mapCtx, in.DryRunSpec)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func Policy_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.Policy {
	if in == nil {
		return nil
	}
	out := &krm.Policy{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Spec = PolicySpec_FromProto(mapCtx, in.GetSpec())
	out.DryRunSpec = PolicySpec_FromProto(mapCtx, in.GetDryRunSpec())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func Policy_ToProto(mapCtx *direct.MapContext, in *krm.Policy) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	out.Name = direct.ValueOf(in.Name)
	out.Spec = PolicySpec_ToProto(mapCtx, in.Spec)
	out.DryRunSpec = PolicySpec_ToProto(mapCtx, in.DryRunSpec)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func PolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.PolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PolicyObservedState{}
	// MISSING: Name
	out.Spec = PolicySpecObservedState_FromProto(mapCtx, in.GetSpec())
	// MISSING: Alternate
	// MISSING: DryRunSpec
	return out
}
func PolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PolicyObservedState) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	// MISSING: Name
	out.Spec = PolicySpecObservedState_ToProto(mapCtx, in.Spec)
	// MISSING: Alternate
	// MISSING: DryRunSpec
	return out
}
func PolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.PolicySpec) *krm.PolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.PolicySpec{}
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: UpdateTime
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, PolicySpec_PolicyRule_FromProto)
	out.InheritFromParent = direct.LazyPtr(in.GetInheritFromParent())
	out.Reset = direct.LazyPtr(in.GetReset_())
	return out
}
func PolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.PolicySpec) *pb.PolicySpec {
	if in == nil {
		return nil
	}
	out := &pb.PolicySpec{}
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: UpdateTime
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, PolicySpec_PolicyRule_ToProto)
	out.InheritFromParent = direct.ValueOf(in.InheritFromParent)
	out.Reset_ = direct.ValueOf(in.Reset)
	return out
}
func PolicySpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PolicySpec) *krm.PolicySpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PolicySpecObservedState{}
	// MISSING: Etag
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Rules
	// MISSING: InheritFromParent
	// MISSING: Reset
	return out
}
func PolicySpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PolicySpecObservedState) *pb.PolicySpec {
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
func PolicySpec_PolicyRule_FromProto(mapCtx *direct.MapContext, in *pb.PolicySpec_PolicyRule) *krm.PolicySpec_PolicyRule {
	if in == nil {
		return nil
	}
	out := &krm.PolicySpec_PolicyRule{}
	out.Values = PolicySpec_PolicyRule_StringValues_FromProto(mapCtx, in.GetValues())
	out.AllowAll = direct.LazyPtr(in.GetAllowAll())
	out.DenyAll = direct.LazyPtr(in.GetDenyAll())
	out.Condition = Expr_FromProto(mapCtx, in.GetCondition())
	return out
}
func PolicySpec_PolicyRule_ToProto(mapCtx *direct.MapContext, in *krm.PolicySpec_PolicyRule) *pb.PolicySpec_PolicyRule {
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
	return out
}

func PolicySpec_PolicyRule_AllowAll_ToProto(mapCtx *direct.MapContext, in *bool) *pb.PolicySpec_PolicyRule_AllowAll {
	if in == nil {
		return nil
	}

	out := &pb.PolicySpec_PolicyRule_AllowAll{}
	out.AllowAll = direct.ValueOf(in)
	return out
}

func PolicySpec_PolicyRule_DenyAll_ToProto(mapCtx *direct.MapContext, in *bool) *pb.PolicySpec_PolicyRule_DenyAll {
	if in == nil {
		return nil
	}

	out := &pb.PolicySpec_PolicyRule_DenyAll{}
	out.DenyAll = direct.ValueOf(in)
	return out
}
func PolicySpec_PolicyRule_Enforce_ToProto(mapCtx *direct.MapContext, in *bool) *pb.PolicySpec_PolicyRule_Enforce {
	if in == nil {
		return nil
	}

	out := &pb.PolicySpec_PolicyRule_Enforce{}
	out.Enforce = direct.ValueOf(in)
	return out
}
func Expr_ToProto(mapCtx *direct.MapContext, in *krm.Expr) *expr.Expr {
	if in == nil {
		return nil
	}
	out := &expr.Expr{}
	out.Expression = direct.ValueOf(in.Expression)
	out.Location = direct.ValueOf(in.Location)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	return out
}

func Expr_FromProto(mapCtx *direct.MapContext, in *expr.Expr) *krm.Expr {
	if in == nil {
		return nil
	}
	out := &krm.Expr{}
	out.Expression = direct.LazyPtr(in.Expression)
	out.Location = direct.LazyPtr(in.Location)
	out.Title = direct.LazyPtr(in.Title)
	out.Description = direct.LazyPtr(in.Description)
	return out
}

func PolicySpec_PolicyRule_StringValues_FromProto(mapCtx *direct.MapContext, in *pb.PolicySpec_PolicyRule_StringValues) *krm.PolicySpec_PolicyRule_StringValues {
	if in == nil {
		return nil
	}
	out := &krm.PolicySpec_PolicyRule_StringValues{}
	out.AllowedValues = in.AllowedValues
	out.DeniedValues = in.DeniedValues
	return out
}
func PolicySpec_PolicyRule_StringValues_ToProto(mapCtx *direct.MapContext, in *krm.PolicySpec_PolicyRule_StringValues) *pb.PolicySpec_PolicyRule_StringValues {
	if in == nil {
		return nil
	}
	out := &pb.PolicySpec_PolicyRule_StringValues{}
	out.AllowedValues = in.AllowedValues
	out.DeniedValues = in.DeniedValues
	return out
}
