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


// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccount
type ManagedAccount struct {
	// The name of the account which uniquely identifies the account.
	//  Format:
	//  projects/{project}/accountManagers/{account_manager}/accounts/{account}
	//  When account manager is used for managing UPI Lite transactions,
	//  `{account}` is the Lite Reference Number (LRN).
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccount.name
	Name *string `json:"name,omitempty"`

	// Required. The associated bank account information.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccount.account_reference
	AccountReference *AccountReference `json:"accountReference,omitempty"`

	// Required. Current balance of the account.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccount.balance
	Balance *Money `json:"balance,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.AccountReference
type AccountReference struct {
	// IFSC of the account's bank branch.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.AccountReference.ifsc
	Ifsc *string `json:"ifsc,omitempty"`

	// Unique number for an account in a bank and branch.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.AccountReference.account_number
	AccountNumber *string `json:"accountNumber,omitempty"`
}

// +kcc:proto=google.type.Money
type Money struct {
	// The three-letter currency code defined in ISO 4217.
	// +kcc:proto:field=google.type.Money.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// The whole units of the amount.
	//  For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	// +kcc:proto:field=google.type.Money.units
	Units *int64 `json:"units,omitempty"`

	// Number of nano (10^-9) units of the amount.
	//  The value must be between -999,999,999 and +999,999,999 inclusive.
	//  If `units` is positive, `nanos` must be positive or zero.
	//  If `units` is zero, `nanos` can be positive, zero, or negative.
	//  If `units` is negative, `nanos` must be negative or zero.
	//  For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	// +kcc:proto:field=google.type.Money.nanos
	Nanos *int32 `json:"nanos,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccount
type ManagedAccountObservedState struct {
	// Required. The associated bank account information.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccount.account_reference
	AccountReference *AccountReferenceObservedState `json:"accountReference,omitempty"`

	// Output only. State of the account.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccount.state
	State *string `json:"state,omitempty"`

	// Output only. State of the last reconciliation done on the account.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccount.last_reconciliation_state
	LastReconciliationState *string `json:"lastReconciliationState,omitempty"`

	// Output only. Time at which last reconciliation was done on the account.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccount.last_reconciliation_time
	LastReconciliationTime *string `json:"lastReconciliationTime,omitempty"`

	// Output only. The time at which the account was created by the account
	//  manager.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccount.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the account was last updated by the account
	//  manager.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.ManagedAccount.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.AccountReference
type AccountReferenceObservedState struct {
	// Output only. Type of account. Examples include SAVINGS, CURRENT, etc.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.AccountReference.account_type
	AccountType *string `json:"accountType,omitempty"`
}
