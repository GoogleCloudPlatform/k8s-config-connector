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


// +kcc:proto=google.analytics.admin.v1beta.Account
type Account struct {

	// Required. Human-readable display name for this account.
	// +kcc:proto:field=google.analytics.admin.v1beta.Account.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Country of business. Must be a Unicode CLDR region code.
	// +kcc:proto:field=google.analytics.admin.v1beta.Account.region_code
	RegionCode *string `json:"regionCode,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.Account
type AccountObservedState struct {
	// Output only. Resource name of this account.
	//  Format: accounts/{account}
	//  Example: "accounts/100"
	// +kcc:proto:field=google.analytics.admin.v1beta.Account.name
	Name *string `json:"name,omitempty"`

	// Output only. Time when this account was originally created.
	// +kcc:proto:field=google.analytics.admin.v1beta.Account.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when account payload fields were last updated.
	// +kcc:proto:field=google.analytics.admin.v1beta.Account.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Indicates whether this Account is soft-deleted or not. Deleted
	//  accounts are excluded from List results unless specifically requested.
	// +kcc:proto:field=google.analytics.admin.v1beta.Account.deleted
	Deleted *bool `json:"deleted,omitempty"`

	// Output only. The URI for a Google Marketing Platform organization resource.
	//  Only set when this account is connected to a GMP organization.
	//  Format: marketingplatformadmin.googleapis.com/organizations/{org_id}
	// +kcc:proto:field=google.analytics.admin.v1beta.Account.gmp_organization
	GmpOrganization *string `json:"gmpOrganization,omitempty"`
}
