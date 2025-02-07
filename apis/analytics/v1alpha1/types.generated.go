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


// +kcc:proto=google.analytics.admin.v1beta.GoogleAdsLink
type GoogleAdsLink struct {

	// Immutable. Google Ads customer ID.
	// +kcc:proto:field=google.analytics.admin.v1beta.GoogleAdsLink.customer_id
	CustomerID *string `json:"customerID,omitempty"`

	// Enable personalized advertising features with this integration.
	//  Automatically publish my Google Analytics audience lists and Google
	//  Analytics remarketing events/parameters to the linked Google Ads account.
	//  If this field is not set on create/update, it will be defaulted to true.
	// +kcc:proto:field=google.analytics.admin.v1beta.GoogleAdsLink.ads_personalization_enabled
	AdsPersonalizationEnabled *bool `json:"adsPersonalizationEnabled,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.GoogleAdsLink
type GoogleAdsLinkObservedState struct {
	// Output only. Format:
	//  properties/{propertyId}/googleAdsLinks/{googleAdsLinkId}
	//
	//  Note: googleAdsLinkId is not the Google Ads customer ID.
	// +kcc:proto:field=google.analytics.admin.v1beta.GoogleAdsLink.name
	Name *string `json:"name,omitempty"`

	// Output only. If true, this link is for a Google Ads manager account.
	// +kcc:proto:field=google.analytics.admin.v1beta.GoogleAdsLink.can_manage_clients
	CanManageClients *bool `json:"canManageClients,omitempty"`

	// Output only. Time when this link was originally created.
	// +kcc:proto:field=google.analytics.admin.v1beta.GoogleAdsLink.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when this link was last updated.
	// +kcc:proto:field=google.analytics.admin.v1beta.GoogleAdsLink.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Email address of the user that created the link.
	//  An empty string will be returned if the email address can't be retrieved.
	// +kcc:proto:field=google.analytics.admin.v1beta.GoogleAdsLink.creator_email_address
	CreatorEmailAddress *string `json:"creatorEmailAddress,omitempty"`
}
