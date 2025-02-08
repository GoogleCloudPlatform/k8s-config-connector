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


// +kcc:proto=google.cloud.oracledatabase.v1.CloudAccountDetails
type CloudAccountDetails struct {
}

// +kcc:proto=google.cloud.oracledatabase.v1.Entitlement
type Entitlement struct {
	// Identifier. The name of the Entitlement resource with the format:
	//  projects/{project}/locations/{region}/entitlements/{entitlement}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.Entitlement.name
	Name *string `json:"name,omitempty"`

	// Details of the OCI Cloud Account.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.Entitlement.cloud_account_details
	CloudAccountDetails *CloudAccountDetails `json:"cloudAccountDetails,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.CloudAccountDetails
type CloudAccountDetailsObservedState struct {
	// Output only. OCI account name.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudAccountDetails.cloud_account
	CloudAccount *string `json:"cloudAccount,omitempty"`

	// Output only. OCI account home region.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudAccountDetails.cloud_account_home_region
	CloudAccountHomeRegion *string `json:"cloudAccountHomeRegion,omitempty"`

	// Output only. URL to link an existing account.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudAccountDetails.link_existing_account_uri
	LinkExistingAccountURI *string `json:"linkExistingAccountURI,omitempty"`

	// Output only. URL to create a new account and link.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudAccountDetails.account_creation_uri
	AccountCreationURI *string `json:"accountCreationURI,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.Entitlement
type EntitlementObservedState struct {
	// Details of the OCI Cloud Account.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.Entitlement.cloud_account_details
	CloudAccountDetails *CloudAccountDetailsObservedState `json:"cloudAccountDetails,omitempty"`

	// Output only. Google Cloud Marketplace order ID (aka entitlement ID)
	// +kcc:proto:field=google.cloud.oracledatabase.v1.Entitlement.entitlement_id
	EntitlementID *string `json:"entitlementID,omitempty"`

	// Output only. Entitlement State.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.Entitlement.state
	State *string `json:"state,omitempty"`
}
