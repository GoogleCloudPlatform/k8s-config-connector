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

package v1beta1

// +kcc:proto=mockgcp.cloud.cloudidentity.groups.v1beta1.ExpiryDetail
type ExpiryDetail struct {
	// The time at which the `MembershipRole` will expire.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.ExpiryDetail.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}

// +kcc:proto=mockgcp.cloud.cloudidentity.groups.v1beta1.MembershipRole
type MembershipRole struct {
	// The expiry details of the `MembershipRole`. Expiry details are only supported for `MEMBER` `MembershipRoles`. May be set if `name` is `MEMBER`. Must not be set if `name` is any other value.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.MembershipRole.expiry_detail
	ExpiryDetail *ExpiryDetail `json:"expiryDetail,omitempty"`

	// The name of the `MembershipRole`. Must be one of `OWNER`, `MANAGER`, `MEMBER`.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.MembershipRole.name
	Name *string `json:"name,omitempty"`

	// Evaluations of restrictions applied to parent group on this membership.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.MembershipRole.restriction_evaluations
	RestrictionEvaluations *RestrictionEvaluations `json:"restrictionEvaluations,omitempty"`
}

// +kcc:proto=mockgcp.cloud.cloudidentity.groups.v1beta1.MembershipRoleRestrictionEvaluation
type MembershipRoleRestrictionEvaluation struct {
	// Output only. The current state of the restriction
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.MembershipRoleRestrictionEvaluation.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=mockgcp.cloud.cloudidentity.groups.v1beta1.RestrictionEvaluations
type RestrictionEvaluations struct {
	// Evaluation of the member restriction applied to this membership. Empty if the user lacks permission to view the restriction evaluation.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.RestrictionEvaluations.member_restriction_evaluation
	MemberRestrictionEvaluation *MembershipRoleRestrictionEvaluation `json:"memberRestrictionEvaluation,omitempty"`
}
