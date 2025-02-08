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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/iam/apiv3/iampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func IamPrincipalAccessBoundaryPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrincipalAccessBoundaryPolicy) *krm.IamPrincipalAccessBoundaryPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IamPrincipalAccessBoundaryPolicyObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: DisplayName
	// MISSING: Annotations
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Details
	return out
}
func IamPrincipalAccessBoundaryPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IamPrincipalAccessBoundaryPolicyObservedState) *pb.PrincipalAccessBoundaryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.PrincipalAccessBoundaryPolicy{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: DisplayName
	// MISSING: Annotations
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Details
	return out
}
func IamPrincipalAccessBoundaryPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.PrincipalAccessBoundaryPolicy) *krm.IamPrincipalAccessBoundaryPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.IamPrincipalAccessBoundaryPolicySpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: DisplayName
	// MISSING: Annotations
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Details
	return out
}
func IamPrincipalAccessBoundaryPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.IamPrincipalAccessBoundaryPolicySpec) *pb.PrincipalAccessBoundaryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.PrincipalAccessBoundaryPolicy{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: DisplayName
	// MISSING: Annotations
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Details
	return out
}
func PrincipalAccessBoundaryPolicy_FromProto(mapCtx *direct.MapContext, in *pb.PrincipalAccessBoundaryPolicy) *krm.PrincipalAccessBoundaryPolicy {
	if in == nil {
		return nil
	}
	out := &krm.PrincipalAccessBoundaryPolicy{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Uid
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Annotations = in.Annotations
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Details = PrincipalAccessBoundaryPolicyDetails_FromProto(mapCtx, in.GetDetails())
	return out
}
func PrincipalAccessBoundaryPolicy_ToProto(mapCtx *direct.MapContext, in *krm.PrincipalAccessBoundaryPolicy) *pb.PrincipalAccessBoundaryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.PrincipalAccessBoundaryPolicy{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Uid
	out.Etag = direct.ValueOf(in.Etag)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Annotations = in.Annotations
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Details = PrincipalAccessBoundaryPolicyDetails_ToProto(mapCtx, in.Details)
	return out
}
func PrincipalAccessBoundaryPolicyDetails_FromProto(mapCtx *direct.MapContext, in *pb.PrincipalAccessBoundaryPolicyDetails) *krm.PrincipalAccessBoundaryPolicyDetails {
	if in == nil {
		return nil
	}
	out := &krm.PrincipalAccessBoundaryPolicyDetails{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, PrincipalAccessBoundaryPolicyRule_FromProto)
	out.EnforcementVersion = direct.LazyPtr(in.GetEnforcementVersion())
	return out
}
func PrincipalAccessBoundaryPolicyDetails_ToProto(mapCtx *direct.MapContext, in *krm.PrincipalAccessBoundaryPolicyDetails) *pb.PrincipalAccessBoundaryPolicyDetails {
	if in == nil {
		return nil
	}
	out := &pb.PrincipalAccessBoundaryPolicyDetails{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, PrincipalAccessBoundaryPolicyRule_ToProto)
	out.EnforcementVersion = direct.ValueOf(in.EnforcementVersion)
	return out
}
func PrincipalAccessBoundaryPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrincipalAccessBoundaryPolicy) *krm.PrincipalAccessBoundaryPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrincipalAccessBoundaryPolicyObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Etag
	// MISSING: DisplayName
	// MISSING: Annotations
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Details
	return out
}
func PrincipalAccessBoundaryPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrincipalAccessBoundaryPolicyObservedState) *pb.PrincipalAccessBoundaryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.PrincipalAccessBoundaryPolicy{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Etag
	// MISSING: DisplayName
	// MISSING: Annotations
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Details
	return out
}
func PrincipalAccessBoundaryPolicyRule_FromProto(mapCtx *direct.MapContext, in *pb.PrincipalAccessBoundaryPolicyRule) *krm.PrincipalAccessBoundaryPolicyRule {
	if in == nil {
		return nil
	}
	out := &krm.PrincipalAccessBoundaryPolicyRule{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Resources = in.Resources
	out.Effect = direct.Enum_FromProto(mapCtx, in.GetEffect())
	return out
}
func PrincipalAccessBoundaryPolicyRule_ToProto(mapCtx *direct.MapContext, in *krm.PrincipalAccessBoundaryPolicyRule) *pb.PrincipalAccessBoundaryPolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.PrincipalAccessBoundaryPolicyRule{}
	out.Description = direct.ValueOf(in.Description)
	out.Resources = in.Resources
	out.Effect = direct.Enum_ToProto[pb.PrincipalAccessBoundaryPolicyRule_Effect](mapCtx, in.Effect)
	return out
}
