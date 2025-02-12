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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudidentity/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/cloudidentity/v1beta1"
)

func CloudIdentityMembershipSpec_FromAPI(mapCtx *direct.MapContext, in *api.Membership) *krm.CloudIdentityMembershipSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityMembershipSpec{}
	// MISSING: CreateTime
	// MISSING: DeliverySetting
	out.MemberKey = EntityKey_FromAPI(mapCtx, in.MemberKey)
	// MISSING: Name
	out.PreferredMemberKey = EntityKey_FromAPI(mapCtx, in.PreferredMemberKey)
	out.Roles = direct.Slice_FromProto(mapCtx, in.Roles, MembershipRoles_FromAPI)
	// MISSING: Type
	// MISSING: UpdateTime
	return out
}
func CloudIdentityMembershipSpec_ToAPI(mapCtx *direct.MapContext, in *krm.CloudIdentityMembershipSpec) *api.Membership {
	if in == nil {
		return nil
	}
	out := &api.Membership{}
	// MISSING: CreateTime
	// MISSING: DeliverySetting
	out.MemberKey = EntityKey_ToAPI(mapCtx, in.MemberKey)
	// MISSING: Name
	out.PreferredMemberKey = EntityKey_ToAPI(mapCtx, in.PreferredMemberKey)
	out.Roles = direct.Slice_ToProto(mapCtx, in.Roles, MembershipRoles_ToAPI)
	// MISSING: Type
	// MISSING: UpdateTime
	return out
}
func CloudIdentityMembershipStatus_FromAPI(mapCtx *direct.MapContext, in *api.Membership) *krm.CloudIdentityMembershipStatus {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityMembershipStatus{}
	out.CreateTime = direct.LazyPtr(in.CreateTime)
	out.DeliverySetting = direct.LazyPtr(in.DeliverySetting)
	// MISSING: MemberKey
	// MISSING: Name
	// MISSING: PreferredMemberKey
	// MISSING: Roles
	out.Type = direct.LazyPtr(in.Type)
	out.UpdateTime = direct.LazyPtr(in.UpdateTime)
	return out
}
func CloudIdentityMembershipStatus_ToAPI(mapCtx *direct.MapContext, in *krm.CloudIdentityMembershipStatus) *api.Membership {
	if in == nil {
		return nil
	}
	out := &api.Membership{}
	out.CreateTime = direct.ValueOf(in.CreateTime)
	out.DeliverySetting = direct.ValueOf(in.DeliverySetting)
	// MISSING: MemberKey
	// MISSING: Name
	// MISSING: PreferredMemberKey
	// MISSING: Roles
	out.Type = direct.ValueOf(in.Type)
	out.UpdateTime = direct.ValueOf(in.UpdateTime)
	return out
}
func MembershipExpiryDetail_FromAPI(mapCtx *direct.MapContext, in *api.ExpiryDetail) *krm.MembershipExpiryDetail {
	if in == nil {
		return nil
	}
	out := &krm.MembershipExpiryDetail{}
	out.ExpireTime = direct.LazyPtr(in.ExpireTime)
	return out
}
func MembershipExpiryDetail_ToAPI(mapCtx *direct.MapContext, in *krm.MembershipExpiryDetail) *api.ExpiryDetail {
	if in == nil {
		return nil
	}
	out := &api.ExpiryDetail{}
	out.ExpireTime = direct.ValueOf(in.ExpireTime)
	return out
}
func MembershipRestrictionEvaluations_FromAPI(mapCtx *direct.MapContext, in *api.RestrictionEvaluations) *krm.MembershipRestrictionEvaluations {
	if in == nil {
		return nil
	}
	out := &krm.MembershipRestrictionEvaluations{}
	out.MemberRestrictionEvaluation = MembershipRoleRestrictionEvaluation_FromAPI(mapCtx, in.MemberRestrictionEvaluation)
	return out
}
func MembershipRestrictionEvaluations_ToAPI(mapCtx *direct.MapContext, in *krm.MembershipRestrictionEvaluations) *api.RestrictionEvaluations {
	if in == nil {
		return nil
	}
	out := &api.RestrictionEvaluations{}
	out.MemberRestrictionEvaluation = MembershipRoleRestrictionEvaluation_ToAPI(mapCtx, in.MemberRestrictionEvaluation)
	return out
}
func MembershipRoleRestrictionEvaluation_FromAPI(mapCtx *direct.MapContext, in *api.MembershipRoleRestrictionEvaluation) *krm.MembershipRoleRestrictionEvaluation {
	if in == nil {
		return nil
	}
	out := &krm.MembershipRoleRestrictionEvaluation{}
	out.State = direct.LazyPtr(in.State)
	return out
}
func MembershipRoleRestrictionEvaluation_ToAPI(mapCtx *direct.MapContext, in *krm.MembershipRoleRestrictionEvaluation) *api.MembershipRoleRestrictionEvaluation {
	if in == nil {
		return nil
	}
	out := &api.MembershipRoleRestrictionEvaluation{}
	out.State = direct.ValueOf(in.State)
	return out
}
func MembershipRoles_FromAPI(mapCtx *direct.MapContext, in *api.MembershipRole) *krm.MembershipRoles {
	if in == nil {
		return nil
	}
	out := &krm.MembershipRoles{}
	out.ExpiryDetail = MembershipExpiryDetail_FromAPI(mapCtx, in.ExpiryDetail)
	out.Name = direct.LazyPtr(in.Name)
	out.RestrictionEvaluations = MembershipRestrictionEvaluations_FromAPI(mapCtx, in.RestrictionEvaluations)
	return out
}
func MembershipRoles_ToAPI(mapCtx *direct.MapContext, in *krm.MembershipRoles) *api.MembershipRole {
	if in == nil {
		return nil
	}
	out := &api.MembershipRole{}
	out.ExpiryDetail = MembershipExpiryDetail_ToAPI(mapCtx, in.ExpiryDetail)
	out.Name = direct.ValueOf(in.Name)
	out.RestrictionEvaluations = MembershipRestrictionEvaluations_ToAPI(mapCtx, in.RestrictionEvaluations)
	return out
}
