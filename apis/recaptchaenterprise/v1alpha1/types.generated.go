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


// +kcc:proto=google.cloud.recaptchaenterprise.v1.RelatedAccountGroupMembership
type RelatedAccountGroupMembership struct {
	// Required. Identifier. The resource name for this membership in the format
	//  `projects/{project}/relatedaccountgroups/{relatedaccountgroup}/memberships/{membership}`.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.RelatedAccountGroupMembership.name
	Name *string `json:"name,omitempty"`

	// The unique stable account identifier of the member. The identifier
	//  corresponds to an `account_id` provided in a previous `CreateAssessment` or
	//  `AnnotateAssessment` call.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.RelatedAccountGroupMembership.account_id
	AccountID *string `json:"accountID,omitempty"`

	// Deprecated: use `account_id` instead.
	//  The unique stable hashed account identifier of the member. The identifier
	//  corresponds to a `hashed_account_id` provided in a previous
	//  `CreateAssessment` or `AnnotateAssessment` call.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.RelatedAccountGroupMembership.hashed_account_id
	HashedAccountID []byte `json:"hashedAccountID,omitempty"`
}
