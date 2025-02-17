// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudIdentityMembershipGVK = GroupVersion.WithKind("CloudIdentityMembership")

// CloudIdentityMembershipSpec defines the desired state of CloudIdentityMembership
// +kcc:proto=mockgcp.cloud.cloudidentity.groups.v1beta1.Membership
type CloudIdentityMembershipSpec struct {
	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable.
	// +required
	GroupRef GroupRef `json:"groupRef"`

	// Immutable. The `EntityKey` of the member. Either `member_key` or `preferred_member_key` must be set when calling MembershipsService.CreateMembership but not both; both shall be set when returned.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.Membership.member_key
	MemberKey *EntityKey `json:"memberKey,omitempty"`

	// Required. Immutable. The `EntityKey` of the member. Either `member_key` or `preferred_member_key` must be set when calling MembershipsService.CreateMembership but not both; both shall be set when returned.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.Membership.preferred_member_key
	// +required
	PreferredMemberKey *EntityKey `json:"preferredMemberKey,omitempty"`

	// The `MembershipRole`s that apply to the `Membership`. If unspecified, defaults to a single `MembershipRole` with `name` `MEMBER`. Must not contain duplicate `MembershipRole`s with the same `name`.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.Membership.roles
	// +required
	Roles []MembershipRoles `json:"roles"`
}

// CloudIdentityMembershipStatus defines the config connector machine state of CloudIdentityMembership
type CloudIdentityMembershipStatus struct {
	commonv1alpha1.CommonStatus `json:",inline"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *CloudIdentityMembershipObservedState `json:"observedState,omitempty"`

	// Output only. The time when the `Membership` was created.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.Membership.create_time
	// +kubebuilder:validation:Format=date-time
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. Delivery setting associated with the membership. Possible values: DELIVERY_SETTING_UNSPECIFIED, ALL_MAIL, DIGEST, DAILY, NONE, DISABLED */
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.Membership.delivery_setting
	DeliverySetting *string `json:"deliverySetting,omitempty"`

	// displayName: This field does not exist in the v1/v1beta1/v1alpha1 proto defn.

	// Output only. The display name of this member, if available
	DisplayName *MembershipDisplayNameStatus `json:"displayName,omitempty"`

	// Output only. The type of the membership. Possible values: OWNER_TYPE_UNSPECIFIED, OWNER_TYPE_CUSTOMER, OWNER_TYPE_PARTNER
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.Membership.type
	Type *string `json:"type,omitempty"`

	// Output only. The time when the `Membership` was last updated.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.Membership.update_time
	// +kubebuilder:validation:Format=date-time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// CloudIdentityMembershipObservedState is the state of the CloudIdentityMembership resource as most recently observed in GCP.
type CloudIdentityMembershipObservedState struct {
	// The name of the Membership resource in GCP. Server generated.
	// Shall be of the form `groups/{group_id}/memberships/{membership_id}`.
	//Name string `json:"name,omitempty"`

	// These fields exist in the old dcl generated status struct.
	// So not adding them here:
	// CreateTime
	// UpdateTime
	// Type
	// DeliverySetting

	// The state output field is in a list
	// .spec.[]roles.restrictionEvaluations.memberRestrictionEvaluationstate
	// Not sure if we need to build an array copying everything from the spec here.
	// []MembershipMemberRestrictionEvaluation.State
}

// +kcc:proto=mockgcp.cloud.cloudidentity.groups.v1beta1.ExpiryDetail
type MembershipExpiryDetail struct {
	// The time at which the `MembershipRole` will expire.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.ExpiryDetail.expire_time
	// +kubebuilder:validation:Format=date-time
	ExpireTime *string `json:"expireTime,omitempty"`
}

// +kcc:proto=mockgcp.cloud.cloudidentity.groups.v1beta1.MembershipRoleRestrictionEvaluation
type MembershipRoleRestrictionEvaluation struct {
	// Output only. The current state of the restriction
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.MembershipRoleRestrictionEvaluation.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=mockgcp.cloud.cloudidentity.groups.v1beta1.RestrictionEvaluations
type MembershipRestrictionEvaluations struct {
	// Evaluation of the member restriction applied to this membership. Empty if the user lacks permission to view the restriction evaluation.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.RestrictionEvaluations.member_restriction_evaluation
	MemberRestrictionEvaluation *MembershipRoleRestrictionEvaluation `json:"memberRestrictionEvaluation,omitempty"`
}

// +kcc:proto=mockgcp.cloud.cloudidentity.groups.v1beta1.MembershipRole
type MembershipRoles struct {
	// The expiry details of the `MembershipRole`. Expiry details are only supported for `MEMBER` `MembershipRoles`. May be set if `name` is `MEMBER`. Must not be set if `name` is any other value.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.MembershipRole.expiry_detail
	ExpiryDetail *MembershipExpiryDetail `json:"expiryDetail,omitempty"`

	// The name of the `MembershipRole`. Must be one of `OWNER`, `MANAGER`, `MEMBER`.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.MembershipRole.name
	// +required
	Name *string `json:"name"`

	// Evaluations of restrictions applied to parent group on this membership.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.MembershipRole.restriction_evaluations
	RestrictionEvaluations *MembershipRestrictionEvaluations `json:"restrictionEvaluations,omitempty"`
}

type MembershipDisplayNameStatus struct {
	// Output only. Member's family name
	FamilyName *string `json:"familyName,omitempty"`

	/* Output only. Localized UTF-16 full name for the member. Localization is done based on the language in the request and the language of the stored display name. */
	FullName *string `json:"fullName,omitempty"`

	// Output only. Member's given name
	GivenName *string `json:"givenName,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudidentitymembership;gcpcloudidentitymemberships
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudIdentityMembership is the Schema for the CloudIdentityMembership API
// +k8s:openapi-gen=true
type CloudIdentityMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec CloudIdentityMembershipSpec `json:"spec,omitempty"`

	Status CloudIdentityMembershipStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudIdentityMembershipList contains a list of CloudIdentityMembership
type CloudIdentityMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudIdentityMembership `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudIdentityMembership{}, &CloudIdentityMembershipList{})
}
