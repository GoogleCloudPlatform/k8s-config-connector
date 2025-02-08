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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/orgpolicy/v1alpha1"
)
func AlternatePolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.AlternatePolicySpec) *krm.AlternatePolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.AlternatePolicySpec{}
	out.Launch = direct.LazyPtr(in.GetLaunch())
	out.Spec = PolicySpec_FromProto(mapCtx, in.GetSpec())
	return out
}
func AlternatePolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.AlternatePolicySpec) *pb.AlternatePolicySpec {
	if in == nil {
		return nil
	}
	out := &pb.AlternatePolicySpec{}
	out.Launch = direct.ValueOf(in.Launch)
	out.Spec = PolicySpec_ToProto(mapCtx, in.Spec)
	return out
}
func OrgpolicyPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.OrgpolicyPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrgpolicyPolicyObservedState{}
	// MISSING: Name
	// MISSING: Spec
	// MISSING: Alternate
	// MISSING: DryRunSpec
	// MISSING: Etag
	return out
}
func OrgpolicyPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrgpolicyPolicyObservedState) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	// MISSING: Name
	// MISSING: Spec
	// MISSING: Alternate
	// MISSING: DryRunSpec
	// MISSING: Etag
	return out
}
func OrgpolicyPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.OrgpolicyPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.OrgpolicyPolicySpec{}
	// MISSING: Name
	// MISSING: Spec
	// MISSING: Alternate
	// MISSING: DryRunSpec
	// MISSING: Etag
	return out
}
func OrgpolicyPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.OrgpolicyPolicySpec) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	// MISSING: Name
	// MISSING: Spec
	// MISSING: Alternate
	// MISSING: DryRunSpec
	// MISSING: Etag
	return out
}
func Policy_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.Policy {
	if in == nil {
		return nil
	}
	out := &krm.Policy{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Spec = PolicySpec_FromProto(mapCtx, in.GetSpec())
	out.Alternate = AlternatePolicySpec_FromProto(mapCtx, in.GetAlternate())
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
	out.Alternate = AlternatePolicySpec_ToProto(mapCtx, in.Alternate)
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
	out.Reset = direct.LazyPtr(in.GetReset())
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
	out.Reset = direct.ValueOf(in.Reset)
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
	out.Enforce = direct.LazyPtr(in.GetEnforce())
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
