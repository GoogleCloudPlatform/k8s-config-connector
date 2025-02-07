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

package v1alpha1


// +kcc:proto=google.cloud.channel.v1.ChannelPartnerLink
type ChannelPartnerLink struct {

	// Required. Cloud Identity ID of the linked reseller.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.reseller_cloud_identity_id
	ResellerCloudIdentityID *string `json:"resellerCloudIdentityID,omitempty"`

	// Required. State of the channel partner link.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.link_state
	LinkState *string `json:"linkState,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.CloudIdentityInfo
type CloudIdentityInfo struct {
	// CustomerType indicates verification type needed for using services.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.customer_type
	CustomerType *string `json:"customerType,omitempty"`

	// The alternate email.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.alternate_email
	AlternateEmail *string `json:"alternateEmail,omitempty"`

	// Phone number associated with the Cloud Identity.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.phone_number
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// Language code.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Edu information about the customer.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.edu_data
	EduData *EduData `json:"eduData,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.EduData
type EduData struct {
	// Designated institute type of customer.
	// +kcc:proto:field=google.cloud.channel.v1.EduData.institute_type
	InstituteType *string `json:"instituteType,omitempty"`

	// Size of the institute.
	// +kcc:proto:field=google.cloud.channel.v1.EduData.institute_size
	InstituteSize *string `json:"instituteSize,omitempty"`

	// Web address for the edu customer's institution.
	// +kcc:proto:field=google.cloud.channel.v1.EduData.website
	Website *string `json:"website,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.ChannelPartnerLink
type ChannelPartnerLinkObservedState struct {
	// Output only. Resource name for the channel partner link, in the format
	//  accounts/{account_id}/channelPartnerLinks/{id}.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.name
	Name *string `json:"name,omitempty"`

	// Output only. URI of the web page where partner accepts the link invitation.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.invite_link_uri
	InviteLinkURI *string `json:"inviteLinkURI,omitempty"`

	// Output only. Timestamp of when the channel partner link is created.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp of when the channel partner link is updated.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Public identifier that a customer must use to generate a
	//  transfer token to move to this distributor-reseller combination.
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.public_id
	PublicID *string `json:"publicID,omitempty"`

	// Output only. Cloud Identity info of the channel partner (IR).
	// +kcc:proto:field=google.cloud.channel.v1.ChannelPartnerLink.channel_partner_cloud_identity_info
	ChannelPartnerCloudIdentityInfo *CloudIdentityInfo `json:"channelPartnerCloudIdentityInfo,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.CloudIdentityInfo
type CloudIdentityInfoObservedState struct {
	// Output only. The primary domain name.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.primary_domain
	PrimaryDomain *string `json:"primaryDomain,omitempty"`

	// Output only. Whether the domain is verified.
	//  This field is not returned for a Customer's cloud_identity_info resource.
	//  Partners can use the domains.get() method of the Workspace SDK's
	//  Directory API, or listen to the PRIMARY_DOMAIN_VERIFIED Pub/Sub event in
	//  to track domain verification of their resolve Workspace customers.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.is_domain_verified
	IsDomainVerified *bool `json:"isDomainVerified,omitempty"`

	// Output only. URI of Customer's Admin console dashboard.
	// +kcc:proto:field=google.cloud.channel.v1.CloudIdentityInfo.admin_console_uri
	AdminConsoleURI *string `json:"adminConsoleURI,omitempty"`
}
