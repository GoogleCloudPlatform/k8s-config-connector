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

package recaptchaenterprise

import (
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recaptchaenterprise/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FirewallAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction) *krm.FirewallAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction{}
	out.Allow = FirewallAction_AllowAction_FromProto(mapCtx, in.GetAllow())
	out.Block = FirewallAction_BlockAction_FromProto(mapCtx, in.GetBlock())
	out.IncludeRecaptchaScript = FirewallAction_IncludeRecaptchaScriptAction_FromProto(mapCtx, in.GetIncludeRecaptchaScript())
	out.Redirect = FirewallAction_RedirectAction_FromProto(mapCtx, in.GetRedirect())
	out.Substitute = FirewallAction_SubstituteAction_FromProto(mapCtx, in.GetSubstitute())
	out.SetHeader = FirewallAction_SetHeaderAction_FromProto(mapCtx, in.GetSetHeader())
	return out
}
func FirewallAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction) *pb.FirewallAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction{}
	if oneof := FirewallAction_AllowAction_ToProto(mapCtx, in.Allow); oneof != nil {
		out.FirewallActionOneof = &pb.FirewallAction_Allow{Allow: oneof}
	}
	if oneof := FirewallAction_BlockAction_ToProto(mapCtx, in.Block); oneof != nil {
		out.FirewallActionOneof = &pb.FirewallAction_Block{Block: oneof}
	}
	if oneof := FirewallAction_IncludeRecaptchaScriptAction_ToProto(mapCtx, in.IncludeRecaptchaScript); oneof != nil {
		out.FirewallActionOneof = &pb.FirewallAction_IncludeRecaptchaScript{IncludeRecaptchaScript: oneof}
	}
	if oneof := FirewallAction_RedirectAction_ToProto(mapCtx, in.Redirect); oneof != nil {
		out.FirewallActionOneof = &pb.FirewallAction_Redirect{Redirect: oneof}
	}
	if oneof := FirewallAction_SubstituteAction_ToProto(mapCtx, in.Substitute); oneof != nil {
		out.FirewallActionOneof = &pb.FirewallAction_Substitute{Substitute: oneof}
	}
	if oneof := FirewallAction_SetHeaderAction_ToProto(mapCtx, in.SetHeader); oneof != nil {
		out.FirewallActionOneof = &pb.FirewallAction_SetHeader{SetHeader: oneof}
	}
	return out
}
func FirewallAction_AllowAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction_AllowAction) *krm.FirewallAction_AllowAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction_AllowAction{}
	return out
}
func FirewallAction_AllowAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction_AllowAction) *pb.FirewallAction_AllowAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction_AllowAction{}
	return out
}
func FirewallAction_BlockAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction_BlockAction) *krm.FirewallAction_BlockAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction_BlockAction{}
	return out
}
func FirewallAction_BlockAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction_BlockAction) *pb.FirewallAction_BlockAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction_BlockAction{}
	return out
}
func FirewallAction_IncludeRecaptchaScriptAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction_IncludeRecaptchaScriptAction) *krm.FirewallAction_IncludeRecaptchaScriptAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction_IncludeRecaptchaScriptAction{}
	return out
}
func FirewallAction_IncludeRecaptchaScriptAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction_IncludeRecaptchaScriptAction) *pb.FirewallAction_IncludeRecaptchaScriptAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction_IncludeRecaptchaScriptAction{}
	return out
}
func FirewallAction_RedirectAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction_RedirectAction) *krm.FirewallAction_RedirectAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction_RedirectAction{}
	return out
}
func FirewallAction_RedirectAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction_RedirectAction) *pb.FirewallAction_RedirectAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction_RedirectAction{}
	return out
}
func FirewallAction_SetHeaderAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction_SetHeaderAction) *krm.FirewallAction_SetHeaderAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction_SetHeaderAction{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func FirewallAction_SetHeaderAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction_SetHeaderAction) *pb.FirewallAction_SetHeaderAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction_SetHeaderAction{}
	out.Key = direct.ValueOf(in.Key)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func FirewallAction_SubstituteAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction_SubstituteAction) *krm.FirewallAction_SubstituteAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction_SubstituteAction{}
	out.Path = direct.LazyPtr(in.GetPath())
	return out
}
func FirewallAction_SubstituteAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction_SubstituteAction) *pb.FirewallAction_SubstituteAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction_SubstituteAction{}
	out.Path = direct.ValueOf(in.Path)
	return out
}
func FirewallPolicy_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicy) *krm.FirewallPolicy {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicy{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Path = direct.LazyPtr(in.GetPath())
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.Actions = direct.Slice_FromProto(mapCtx, in.Actions, FirewallAction_FromProto)
	return out
}
func FirewallPolicy_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicy) *pb.FirewallPolicy {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicy{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Path = direct.ValueOf(in.Path)
	out.Condition = direct.ValueOf(in.Condition)
	out.Actions = direct.Slice_ToProto(mapCtx, in.Actions, FirewallAction_ToProto)
	return out
}
func ReCAPTCHAEnterpriseFirewallPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicy) *krm.ReCAPTCHAEnterpriseFirewallPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ReCAPTCHAEnterpriseFirewallPolicyObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Path
	// MISSING: Condition
	// MISSING: Actions
	return out
}
func ReCAPTCHAEnterpriseFirewallPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReCAPTCHAEnterpriseFirewallPolicyObservedState) *pb.FirewallPolicy {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicy{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Path
	// MISSING: Condition
	// MISSING: Actions
	return out
}
func ReCAPTCHAEnterpriseFirewallPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicy) *krm.ReCAPTCHAEnterpriseFirewallPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.ReCAPTCHAEnterpriseFirewallPolicySpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Path
	// MISSING: Condition
	// MISSING: Actions
	return out
}
func ReCAPTCHAEnterpriseFirewallPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.ReCAPTCHAEnterpriseFirewallPolicySpec) *pb.FirewallPolicy {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicy{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Path
	// MISSING: Condition
	// MISSING: Actions
	return out
}
