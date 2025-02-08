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


// +kcc:proto=google.cloud.retail.v2beta.AlertConfig
type AlertConfig struct {
	// Required. Immutable. The name of the AlertConfig singleton resource.
	//  Format: projects/*/alertConfig
	// +kcc:proto:field=google.cloud.retail.v2beta.AlertConfig.name
	Name *string `json:"name,omitempty"`

	// Alert policies for a customer.
	//  They must be unique by [AlertPolicy.alert_group]
	// +kcc:proto:field=google.cloud.retail.v2beta.AlertConfig.alert_policies
	AlertPolicies []AlertConfig_AlertPolicy `json:"alertPolicies,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.AlertConfig.AlertPolicy
type AlertConfig_AlertPolicy struct {
	// The feature that provides alerting capability.
	//  Supported value:
	//  - `search-data-quality` for retail search customers.
	//  - `conv-data-quality` for retail conversation customers.
	// +kcc:proto:field=google.cloud.retail.v2beta.AlertConfig.AlertPolicy.alert_group
	AlertGroup *string `json:"alertGroup,omitempty"`

	// The enrollment status of a customer.
	// +kcc:proto:field=google.cloud.retail.v2beta.AlertConfig.AlertPolicy.enroll_status
	EnrollStatus *string `json:"enrollStatus,omitempty"`

	// Recipients for the alert policy.
	//  One alert policy should not exceed 20 recipients.
	// +kcc:proto:field=google.cloud.retail.v2beta.AlertConfig.AlertPolicy.recipients
	Recipients []AlertConfig_AlertPolicy_Recipient `json:"recipients,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.AlertConfig.AlertPolicy.Recipient
type AlertConfig_AlertPolicy_Recipient struct {
	// Email address of the recipient.
	// +kcc:proto:field=google.cloud.retail.v2beta.AlertConfig.AlertPolicy.Recipient.email_address
	EmailAddress *string `json:"emailAddress,omitempty"`
}
