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


// +kcc:proto=google.iam.v3beta.PrincipalAccessBoundaryPolicy
type PrincipalAccessBoundaryPolicy struct {
	// Identifier. The resource name of the principal access boundary policy.
	//
	//  The following format is supported:
	//  `organizations/{organization_id}/locations/{location}/principalAccessBoundaryPolicies/{policy_id}`
	// +kcc:proto:field=google.iam.v3beta.PrincipalAccessBoundaryPolicy.name
	Name *string `json:"name,omitempty"`

	// Optional. The etag for the principal access boundary.
	//  If this is provided on update, it must match the server's etag.
	// +kcc:proto:field=google.iam.v3beta.PrincipalAccessBoundaryPolicy.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The description of the principal access boundary policy. Must be
	//  less than or equal to 63 characters.
	// +kcc:proto:field=google.iam.v3beta.PrincipalAccessBoundaryPolicy.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User defined annotations. See
	//  https://google.aip.dev/148#annotations for more details such as format and
	//  size limitations
	// +kcc:proto:field=google.iam.v3beta.PrincipalAccessBoundaryPolicy.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. The details for the principal access boundary policy.
	// +kcc:proto:field=google.iam.v3beta.PrincipalAccessBoundaryPolicy.details
	Details *PrincipalAccessBoundaryPolicyDetails `json:"details,omitempty"`
}

// +kcc:proto=google.iam.v3beta.PrincipalAccessBoundaryPolicyDetails
type PrincipalAccessBoundaryPolicyDetails struct {
	// Required. A list of principal access boundary policy rules. The number of
	//  rules in a policy is limited to 500.
	// +kcc:proto:field=google.iam.v3beta.PrincipalAccessBoundaryPolicyDetails.rules
	Rules []PrincipalAccessBoundaryPolicyRule `json:"rules,omitempty"`

	// Optional.
	//  The version number that indicates which Google Cloud
	//  services are included in the enforcement (e.g. "latest", "1", ...). If
	//  empty, the PAB policy version will be set to the current latest version,
	//  and this version won't get updated when new versions are released.
	// +kcc:proto:field=google.iam.v3beta.PrincipalAccessBoundaryPolicyDetails.enforcement_version
	EnforcementVersion *string `json:"enforcementVersion,omitempty"`
}

// +kcc:proto=google.iam.v3beta.PrincipalAccessBoundaryPolicyRule
type PrincipalAccessBoundaryPolicyRule struct {
	// Optional. The description of the principal access boundary policy rule.
	//  Must be less than or equal to 256 characters.
	// +kcc:proto:field=google.iam.v3beta.PrincipalAccessBoundaryPolicyRule.description
	Description *string `json:"description,omitempty"`

	// Required. A list of Cloud Resource Manager resources. The resource and all
	//  the descendants are included. The number of resources in a policy is
	//  limited to 500 across all rules.
	//
	//  The following resource types are supported:
	//
	//  * Organizations, such as
	//  `//cloudresourcemanager.googleapis.com/organizations/123`.
	//  * Folders, such as `//cloudresourcemanager.googleapis.com/folders/123`.
	//  * Projects, such as `//cloudresourcemanager.googleapis.com/projects/123`
	//    or `//cloudresourcemanager.googleapis.com/projects/my-project-id`.
	// +kcc:proto:field=google.iam.v3beta.PrincipalAccessBoundaryPolicyRule.resources
	Resources []string `json:"resources,omitempty"`

	// Required. The access relationship of principals to the resources in this
	//  rule.
	// +kcc:proto:field=google.iam.v3beta.PrincipalAccessBoundaryPolicyRule.effect
	Effect *string `json:"effect,omitempty"`
}

// +kcc:proto=google.iam.v3beta.PrincipalAccessBoundaryPolicy
type PrincipalAccessBoundaryPolicyObservedState struct {
	// Output only. The globally unique ID of the principal access boundary
	//  policy.
	// +kcc:proto:field=google.iam.v3beta.PrincipalAccessBoundaryPolicy.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the principal access boundary policy was
	//  created.
	// +kcc:proto:field=google.iam.v3beta.PrincipalAccessBoundaryPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the principal access boundary policy was most
	//  recently updated.
	// +kcc:proto:field=google.iam.v3beta.PrincipalAccessBoundaryPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
