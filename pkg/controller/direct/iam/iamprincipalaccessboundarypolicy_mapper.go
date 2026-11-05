// Copyright 2026 Google LLC
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
	iamv3pb "cloud.google.com/go/iam/apiv3/iampb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func IAMPrincipalAccessBoundaryPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *iamv3pb.PrincipalAccessBoundaryPolicy) *krmv1alpha1.IAMPrincipalAccessBoundaryPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.IAMPrincipalAccessBoundaryPolicyObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func IAMPrincipalAccessBoundaryPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.IAMPrincipalAccessBoundaryPolicyObservedState) *iamv3pb.PrincipalAccessBoundaryPolicy {
	if in == nil {
		return nil
	}
	out := &iamv3pb.PrincipalAccessBoundaryPolicy{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func IAMPrincipalAccessBoundaryPolicySpec_FromProto(mapCtx *direct.MapContext, in *iamv3pb.PrincipalAccessBoundaryPolicy) *krmv1alpha1.IAMPrincipalAccessBoundaryPolicySpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.IAMPrincipalAccessBoundaryPolicySpec{}
	// MISSING: Name
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Annotations = in.Annotations
	out.Details = PrincipalAccessBoundaryPolicyDetails_FromProto(mapCtx, in.GetDetails())
	return out
}

func IAMPrincipalAccessBoundaryPolicySpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.IAMPrincipalAccessBoundaryPolicySpec) *iamv3pb.PrincipalAccessBoundaryPolicy {
	if in == nil {
		return nil
	}
	out := &iamv3pb.PrincipalAccessBoundaryPolicy{}
	// MISSING: Name
	out.Etag = direct.ValueOf(in.Etag)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Annotations = in.Annotations
	out.Details = PrincipalAccessBoundaryPolicyDetails_ToProto(mapCtx, in.Details)
	return out
}

func PrincipalAccessBoundaryPolicyDetails_FromProto(mapCtx *direct.MapContext, in *iamv3pb.PrincipalAccessBoundaryPolicyDetails) *krmv1alpha1.PrincipalAccessBoundaryPolicyDetails {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PrincipalAccessBoundaryPolicyDetails{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, PrincipalAccessBoundaryPolicyRule_FromProto)
	out.EnforcementVersion = direct.LazyPtr(in.GetEnforcementVersion())
	return out
}

func PrincipalAccessBoundaryPolicyDetails_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PrincipalAccessBoundaryPolicyDetails) *iamv3pb.PrincipalAccessBoundaryPolicyDetails {
	if in == nil {
		return nil
	}
	out := &iamv3pb.PrincipalAccessBoundaryPolicyDetails{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, PrincipalAccessBoundaryPolicyRule_ToProto)
	out.EnforcementVersion = direct.ValueOf(in.EnforcementVersion)
	return out
}

func PrincipalAccessBoundaryPolicyRule_FromProto(mapCtx *direct.MapContext, in *iamv3pb.PrincipalAccessBoundaryPolicyRule) *krmv1alpha1.PrincipalAccessBoundaryPolicyRule {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PrincipalAccessBoundaryPolicyRule{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Resources = in.Resources
	out.Effect = direct.Enum_FromProto(mapCtx, in.GetEffect())
	return out
}

func PrincipalAccessBoundaryPolicyRule_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PrincipalAccessBoundaryPolicyRule) *iamv3pb.PrincipalAccessBoundaryPolicyRule {
	if in == nil {
		return nil
	}
	out := &iamv3pb.PrincipalAccessBoundaryPolicyRule{}
	out.Description = direct.ValueOf(in.Description)
	out.Resources = in.Resources
	out.Effect = direct.Enum_ToProto[iamv3pb.PrincipalAccessBoundaryPolicyRule_Effect](mapCtx, in.Effect)
	return out
}
