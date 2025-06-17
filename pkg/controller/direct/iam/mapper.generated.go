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
// proto.service: google.iam.v2

package iam

import (
	pb "cloud.google.com/go/iam/apiv2/iampb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DenyRule_FromProto(mapCtx *direct.MapContext, in *pb.DenyRule) *krmv1alpha1.DenyRule {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DenyRule{}
	out.DeniedPrincipals = in.DeniedPrincipals
	out.ExceptionPrincipals = in.ExceptionPrincipals
	out.DeniedPermissions = in.DeniedPermissions
	out.ExceptionPermissions = in.ExceptionPermissions
	out.DenialCondition = Expr_FromProto(mapCtx, in.GetDenialCondition())
	return out
}
func DenyRule_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DenyRule) *pb.DenyRule {
	if in == nil {
		return nil
	}
	out := &pb.DenyRule{}
	out.DeniedPrincipals = in.DeniedPrincipals
	out.ExceptionPrincipals = in.ExceptionPrincipals
	out.DeniedPermissions = in.DeniedPermissions
	out.ExceptionPermissions = in.ExceptionPermissions
	out.DenialCondition = Expr_ToProto(mapCtx, in.DenialCondition)
	return out
}
func IAMDenyPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krmv1alpha1.IAMDenyPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.IAMDenyPolicyObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Kind
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ManagingAuthority
	return out
}
func IAMDenyPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.IAMDenyPolicyObservedState) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Kind
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ManagingAuthority
	return out
}
func IAMDenyPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krmv1alpha1.IAMDenyPolicySpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.IAMDenyPolicySpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Kind
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, PolicyRule_FromProto)
	// MISSING: ManagingAuthority
	return out
}
func IAMDenyPolicySpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.IAMDenyPolicySpec) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Kind
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, PolicyRule_ToProto)
	// MISSING: ManagingAuthority
	return out
}
func PolicyRule_FromProto(mapCtx *direct.MapContext, in *pb.PolicyRule) *krmv1alpha1.PolicyRule {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PolicyRule{}
	out.DenyRule = DenyRule_FromProto(mapCtx, in.GetDenyRule())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func PolicyRule_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PolicyRule) *pb.PolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.PolicyRule{}
	if oneof := DenyRule_ToProto(mapCtx, in.DenyRule); oneof != nil {
		out.Kind = &pb.PolicyRule_DenyRule{DenyRule: oneof}
	}
	out.Description = direct.ValueOf(in.Description)
	return out
}
