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


// +kcc:proto=google.cloud.channel.v1.BillingAccount
type BillingAccount struct {

	// Display name of the billing account.
	// +kcc:proto:field=google.cloud.channel.v1.BillingAccount.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.BillingAccount
type BillingAccountObservedState struct {
	// Output only. Resource name of the billing account.
	//  Format: accounts/{account_id}/billingAccounts/{billing_account_id}.
	// +kcc:proto:field=google.cloud.channel.v1.BillingAccount.name
	Name *string `json:"name,omitempty"`

	// Output only. The time when this billing account was created.
	// +kcc:proto:field=google.cloud.channel.v1.BillingAccount.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The 3-letter currency code defined in ISO 4217.
	// +kcc:proto:field=google.cloud.channel.v1.BillingAccount.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// Output only. The CLDR region code.
	// +kcc:proto:field=google.cloud.channel.v1.BillingAccount.region_code
	RegionCode *string `json:"regionCode,omitempty"`
}
