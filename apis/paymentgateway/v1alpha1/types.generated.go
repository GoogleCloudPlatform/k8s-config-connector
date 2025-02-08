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


// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.Rule
type Rule struct {
	// The unique identifier for this resource.
	//  Format: projects/{project}/rules/{rule}
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Rule.name
	Name *string `json:"name,omitempty"`

	// The description of the rule.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Rule.rule_description
	RuleDescription *string `json:"ruleDescription,omitempty"`

	// The API Type for which this rule gets executed. A value of
	//  `API_TYPE_UNSPECIFIED` indicates that the rule is executed for all API
	//  transactions.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Rule.api_type
	ApiType *string `json:"apiType,omitempty"`

	// The transaction type for which this rule gets executed. A value of
	//  `TRANSACTION_TYPE_UNSPECIFIED` indicates that the rule is executed for
	//  all transaction types.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Rule.transaction_type
	TransactionType *string `json:"transactionType,omitempty"`
}
