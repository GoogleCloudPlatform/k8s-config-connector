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
// krm.group: cloudidentity.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.apps.cloudidentity.v1beta1

package cloudidentity

import (
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudidentity/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/apps/cloudidentity/groups/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudIdentityMembershipSpec_FromProto(mapCtx *direct.MapContext, in *pb.Membership) *krmv1beta1.CloudIdentityMembershipSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.CloudIdentityMembershipSpec{}
	// MISSING: CreateTime
	// MISSING: DeliverySetting
	out.MemberKey = EntityKey_FromProto(mapCtx, in.GetMemberKey())
	// MISSING: Name
	out.PreferredMemberKey = EntityKey_FromProto(mapCtx, in.GetPreferredMemberKey())
	out.Roles = direct.Slice_FromProto(mapCtx, in.Roles, MembershipRoles_FromProto)
	// MISSING: Type
	// MISSING: UpdateTime
	return out
}
func CloudIdentityMembershipSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.CloudIdentityMembershipSpec) *pb.Membership {
	if in == nil {
		return nil
	}
	out := &pb.Membership{}
	// MISSING: CreateTime
	// MISSING: DeliverySetting
	out.MemberKey = EntityKey_ToProto(mapCtx, in.MemberKey)
	// MISSING: Name
	out.PreferredMemberKey = EntityKey_ToProto(mapCtx, in.PreferredMemberKey)
	out.Roles = direct.Slice_ToProto(mapCtx, in.Roles, MembershipRoles_ToProto)
	// MISSING: Type
	// MISSING: UpdateTime
	return out
}
func CloudIdentityMembershipStatus_FromProto(mapCtx *direct.MapContext, in *pb.Membership) *krmv1beta1.CloudIdentityMembershipStatus {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.CloudIdentityMembershipStatus{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.DeliverySetting = in.DeliverySetting
	// MISSING: MemberKey
	// MISSING: Name
	// MISSING: PreferredMemberKey
	// MISSING: Roles
	out.Type = in.Type
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func CloudIdentityMembershipStatus_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.CloudIdentityMembershipStatus) *pb.Membership {
	if in == nil {
		return nil
	}
	out := &pb.Membership{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.DeliverySetting = in.DeliverySetting
	// MISSING: MemberKey
	// MISSING: Name
	// MISSING: PreferredMemberKey
	// MISSING: Roles
	out.Type = in.Type
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func MembershipExpiryDetail_FromProto(mapCtx *direct.MapContext, in *pb.ExpiryDetail) *krmv1beta1.MembershipExpiryDetail {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.MembershipExpiryDetail{}
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	return out
}
func MembershipExpiryDetail_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.MembershipExpiryDetail) *pb.ExpiryDetail {
	if in == nil {
		return nil
	}
	out := &pb.ExpiryDetail{}
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	return out
}
func MembershipRestrictionEvaluations_FromProto(mapCtx *direct.MapContext, in *pb.RestrictionEvaluations) *krmv1beta1.MembershipRestrictionEvaluations {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.MembershipRestrictionEvaluations{}
	out.MemberRestrictionEvaluation = MembershipRoleRestrictionEvaluation_FromProto(mapCtx, in.GetMemberRestrictionEvaluation())
	return out
}
func MembershipRestrictionEvaluations_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.MembershipRestrictionEvaluations) *pb.RestrictionEvaluations {
	if in == nil {
		return nil
	}
	out := &pb.RestrictionEvaluations{}
	out.MemberRestrictionEvaluation = MembershipRoleRestrictionEvaluation_ToProto(mapCtx, in.MemberRestrictionEvaluation)
	return out
}
func MembershipRoleRestrictionEvaluation_FromProto(mapCtx *direct.MapContext, in *pb.MembershipRoleRestrictionEvaluation) *krmv1beta1.MembershipRoleRestrictionEvaluation {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.MembershipRoleRestrictionEvaluation{}
	out.State = in.State
	return out
}
func MembershipRoleRestrictionEvaluation_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.MembershipRoleRestrictionEvaluation) *pb.MembershipRoleRestrictionEvaluation {
	if in == nil {
		return nil
	}
	out := &pb.MembershipRoleRestrictionEvaluation{}
	out.State = in.State
	return out
}
func MembershipRoles_FromProto(mapCtx *direct.MapContext, in *pb.MembershipRole) *krmv1beta1.MembershipRoles {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.MembershipRoles{}
	out.ExpiryDetail = MembershipExpiryDetail_FromProto(mapCtx, in.GetExpiryDetail())
	out.Name = in.Name
	out.RestrictionEvaluations = MembershipRestrictionEvaluations_FromProto(mapCtx, in.GetRestrictionEvaluations())
	return out
}
func MembershipRoles_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.MembershipRoles) *pb.MembershipRole {
	if in == nil {
		return nil
	}
	out := &pb.MembershipRole{}
	out.ExpiryDetail = MembershipExpiryDetail_ToProto(mapCtx, in.ExpiryDetail)
	out.Name = in.Name
	out.RestrictionEvaluations = MembershipRestrictionEvaluations_ToProto(mapCtx, in.RestrictionEvaluations)
	return out
}
