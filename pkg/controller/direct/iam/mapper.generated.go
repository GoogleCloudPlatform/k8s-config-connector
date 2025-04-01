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
// krm.group: iam.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.iam.v3

package iam

import (
	pb "cloud.google.com/go/iam/apiv3/iampb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func IAMPolicyBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBinding) *krmv1alpha1.IAMPolicyBindingObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.IAMPolicyBindingObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.PolicyUid = direct.LazyPtr(in.GetPolicyUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func IAMPolicyBindingObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.IAMPolicyBindingObservedState) *pb.PolicyBinding {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBinding{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.PolicyUid = direct.ValueOf(in.PolicyUid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func IAMPolicyBindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBinding) *krmv1alpha1.IAMPolicyBindingSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.IAMPolicyBindingSpec{}
	// MISSING: Name
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Annotations = in.Annotations
	out.Target = PolicyBinding_Target_FromProto(mapCtx, in.GetTarget())
	out.PolicyKind = direct.Enum_FromProto(mapCtx, in.GetPolicyKind())
	out.Policy = direct.LazyPtr(in.GetPolicy())
	out.Condition = Expr_FromProto(mapCtx, in.GetCondition())
	return out
}
func IAMPolicyBindingSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.IAMPolicyBindingSpec) *pb.PolicyBinding {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBinding{}
	// MISSING: Name
	out.Etag = direct.ValueOf(in.Etag)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Annotations = in.Annotations
	out.Target = PolicyBinding_Target_ToProto(mapCtx, in.Target)
	out.PolicyKind = direct.Enum_ToProto[pb.PolicyBinding_PolicyKind](mapCtx, in.PolicyKind)
	out.Policy = direct.ValueOf(in.Policy)
	out.Condition = Expr_ToProto(mapCtx, in.Condition)
	return out
}
func PolicyBinding_Target_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBinding_Target) *krmv1alpha1.PolicyBinding_Target {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PolicyBinding_Target{}
	out.PrincipalSet = direct.LazyPtr(in.GetPrincipalSet())
	return out
}
func PolicyBinding_Target_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PolicyBinding_Target) *pb.PolicyBinding_Target {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBinding_Target{}
	if oneof := PolicyBinding_Target_PrincipalSet_ToProto(mapCtx, in.PrincipalSet); oneof != nil {
		out.Target = oneof
	}
	return out
}
