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


// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.AccountReference
type AccountReference struct {
	// IFSC of the account's bank branch.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.AccountReference.ifsc
	Ifsc *string `json:"ifsc,omitempty"`

	// Unique number for an account in a bank and branch.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.AccountReference.account_number
	AccountNumber *string `json:"accountNumber,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.RuleMetadataValue
type RuleMetadataValue struct {

	// The value for string metadata.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.RuleMetadataValue.id
	ID *string `json:"id,omitempty"`

	// The value for account reference metadata.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.RuleMetadataValue.account_reference
	AccountReference *AccountReference `json:"accountReference,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.AccountReference
type AccountReferenceObservedState struct {
	// Output only. Type of account. Examples include SAVINGS, CURRENT, etc.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.AccountReference.account_type
	AccountType *string `json:"accountType,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.RuleMetadataValue
type RuleMetadataValueObservedState struct {
	// Output only. The unique identifier for this resource.
	//  Format: projects/{project}/rules/{rule}/metadata/{metadata}/values/{value}
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.RuleMetadataValue.name
	Name *string `json:"name,omitempty"`

	// The value for account reference metadata.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.RuleMetadataValue.account_reference
	AccountReference *AccountReferenceObservedState `json:"accountReference,omitempty"`
}
