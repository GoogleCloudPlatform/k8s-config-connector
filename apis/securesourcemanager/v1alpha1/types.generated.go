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


// +kcc:proto=google.cloud.securesourcemanager.v1.BranchRule
type BranchRule struct {
	// Optional. A unique identifier for a BranchRule. The name should be of the
	//  format:
	//  `projects/{project}/locations/{location}/repositories/{repository}/branchRules/{branch_rule}`
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.name
	Name *string `json:"name,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	//  the user. See https://google.aip.dev/128#annotations for more details such
	//  as format and size limitations.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. This checksum is computed by the server based on the value of
	//  other fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The pattern of the branch that can match to this BranchRule.
	//  Specified as regex.
	//  .* for all branches. Examples: main, (main|release.*).
	//  Current MVP phase only support `.*` for wildcard.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.include_pattern
	IncludePattern *string `json:"includePattern,omitempty"`

	// Optional. Determines if the branch rule is disabled or not.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Optional. Determines if the branch rule requires a pull request or not.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.require_pull_request
	RequirePullRequest *bool `json:"requirePullRequest,omitempty"`

	// Optional. The minimum number of reviews required for the branch rule to be
	//  matched.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.minimum_reviews_count
	MinimumReviewsCount *int32 `json:"minimumReviewsCount,omitempty"`

	// Optional. The minimum number of approvals required for the branch rule to
	//  be matched.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.minimum_approvals_count
	MinimumApprovalsCount *int32 `json:"minimumApprovalsCount,omitempty"`

	// Optional. Determines if require comments resolved before merging to the
	//  branch.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.require_comments_resolved
	RequireCommentsResolved *bool `json:"requireCommentsResolved,omitempty"`

	// Optional. Determines if allow stale reviews or approvals before merging to
	//  the branch.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.allow_stale_reviews
	AllowStaleReviews *bool `json:"allowStaleReviews,omitempty"`

	// Optional. Determines if require linear history before merging to the
	//  branch.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.require_linear_history
	RequireLinearHistory *bool `json:"requireLinearHistory,omitempty"`

	// Optional. List of required status checks before merging to the branch.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.required_status_checks
	RequiredStatusChecks []BranchRule_Check `json:"requiredStatusChecks,omitempty"`
}

// +kcc:proto=google.cloud.securesourcemanager.v1.BranchRule.Check
type BranchRule_Check struct {
	// Required. The context of the check.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.Check.context
	Context *string `json:"context,omitempty"`
}

// +kcc:proto=google.cloud.securesourcemanager.v1.BranchRule
type BranchRuleObservedState struct {
	// Output only. Unique identifier of the repository.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Create timestamp.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update timestamp.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
