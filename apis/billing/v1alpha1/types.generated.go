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


// +kcc:proto=google.cloud.billing.v1.BillingAccount
type BillingAccount struct {

	// The display name given to the billing account, such as `My Billing
	//  Account`. This name is displayed in the Google Cloud Console.
	// +kcc:proto:field=google.cloud.billing.v1.BillingAccount.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// If this account is a
	//  [subaccount](https://cloud.google.com/billing/docs/concepts), then this
	//  will be the resource name of the parent billing account that it is being
	//  resold through.
	//  Otherwise this will be empty.
	// +kcc:proto:field=google.cloud.billing.v1.BillingAccount.master_billing_account
	MasterBillingAccount *string `json:"masterBillingAccount,omitempty"`

	// Optional. The currency in which the billing account is billed and charged,
	//  represented as an ISO 4217 code such as `USD`.
	//
	//  Billing account currency is determined at the time of billing account
	//  creation and cannot be updated subsequently, so this field should not be
	//  set on update requests. In addition, a subaccount always matches the
	//  currency of its parent billing account, so this field should not be set on
	//  subaccount creation requests. Clients can read this field to determine the
	//  currency of an existing billing account.
	// +kcc:proto:field=google.cloud.billing.v1.BillingAccount.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// +kcc:proto=google.cloud.billing.v1.BillingAccount
type BillingAccountObservedState struct {
	// Output only. The resource name of the billing account. The resource name
	//  has the form `billingAccounts/{billing_account_id}`. For example,
	//  `billingAccounts/012345-567890-ABCDEF` would be the resource name for
	//  billing account `012345-567890-ABCDEF`.
	// +kcc:proto:field=google.cloud.billing.v1.BillingAccount.name
	Name *string `json:"name,omitempty"`

	// Output only. True if the billing account is open, and will therefore be
	//  charged for any usage on associated projects. False if the billing account
	//  is closed, and therefore projects associated with it are unable to use paid
	//  services.
	// +kcc:proto:field=google.cloud.billing.v1.BillingAccount.open
	Open *bool `json:"open,omitempty"`

	// Output only. The billing account's parent resource identifier.
	//  Use the `MoveBillingAccount` method to update the account's parent resource
	//  if it is a organization.
	//  Format:
	//    - `organizations/{organization_id}`, for example,
	//      `organizations/12345678`
	//    - `billingAccounts/{billing_account_id}`, for example,
	//      `billingAccounts/012345-567890-ABCDEF`
	// +kcc:proto:field=google.cloud.billing.v1.BillingAccount.parent
	Parent *string `json:"parent,omitempty"`
}
