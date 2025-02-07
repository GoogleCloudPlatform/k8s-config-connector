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


// +kcc:proto=google.analytics.admin.v1beta.AccountSummary
type AccountSummary struct {
	// Resource name for this account summary.
	//  Format: accountSummaries/{account_id}
	//  Example: "accountSummaries/1000"
	// +kcc:proto:field=google.analytics.admin.v1beta.AccountSummary.name
	Name *string `json:"name,omitempty"`

	// Resource name of account referred to by this account summary
	//  Format: accounts/{account_id}
	//  Example: "accounts/1000"
	// +kcc:proto:field=google.analytics.admin.v1beta.AccountSummary.account
	Account *string `json:"account,omitempty"`

	// Display name for the account referred to in this account summary.
	// +kcc:proto:field=google.analytics.admin.v1beta.AccountSummary.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// List of summaries for child accounts of this account.
	// +kcc:proto:field=google.analytics.admin.v1beta.AccountSummary.property_summaries
	PropertySummaries []PropertySummary `json:"propertySummaries,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.PropertySummary
type PropertySummary struct {
	// Resource name of property referred to by this property summary
	//  Format: properties/{property_id}
	//  Example: "properties/1000"
	// +kcc:proto:field=google.analytics.admin.v1beta.PropertySummary.property
	Property *string `json:"property,omitempty"`

	// Display name for the property referred to in this property summary.
	// +kcc:proto:field=google.analytics.admin.v1beta.PropertySummary.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The property's property type.
	// +kcc:proto:field=google.analytics.admin.v1beta.PropertySummary.property_type
	PropertyType *string `json:"propertyType,omitempty"`

	// Resource name of this property's logical parent.
	//
	//  Note: The Property-Moving UI can be used to change the parent.
	//  Format: accounts/{account}, properties/{property}
	//  Example: "accounts/100", "properties/200"
	// +kcc:proto:field=google.analytics.admin.v1beta.PropertySummary.parent
	Parent *string `json:"parent,omitempty"`
}
