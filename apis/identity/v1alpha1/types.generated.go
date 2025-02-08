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


// +kcc:proto=google.identity.accesscontextmanager.v1.AccessPolicy
type AccessPolicy struct {
	// Output only. Resource name of the `AccessPolicy`. Format:
	//  `accessPolicies/{access_policy}`
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.name
	Name *string `json:"name,omitempty"`

	// Required. The parent of this `AccessPolicy` in the Cloud Resource
	//  Hierarchy. Currently immutable once created. Format:
	//  `organizations/{organization_id}`
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.parent
	Parent *string `json:"parent,omitempty"`

	// Required. Human readable title. Does not affect behavior.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.title
	Title *string `json:"title,omitempty"`

	// The scopes of a policy define which resources an ACM policy can restrict,
	//  and where ACM resources can be referenced.
	//  For example, a policy with scopes=["folders/123"] has the following
	//  behavior:
	//  - vpcsc perimeters can only restrict projects within folders/123
	//  - access levels can only be referenced by resources within folders/123.
	//  If empty, there are no limitations on which resources can be restricted by
	//  an ACM policy, and there are no limitations on where ACM resources can be
	//  referenced.
	//  Only one policy can include a given scope (attempting to create a second
	//  policy which includes "folders/123" will result in an error).
	//  Currently, scopes cannot be modified after a policy is created.
	//  Currently, policies can only have a single scope.
	//  Format: list of `folders/{folder_number}` or `projects/{project_number}`
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.scopes
	Scopes []string `json:"scopes,omitempty"`

	// Output only. Time the `AccessPolicy` was created in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time the `AccessPolicy` was updated in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. An opaque identifier for the current version of the
	//  `AccessPolicy`. This will always be a strongly validated etag, meaning that
	//  two Access Polices will be identical if and only if their etags are
	//  identical. Clients should not expect this to be in any specific format.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.AccessPolicy.etag
	Etag *string `json:"etag,omitempty"`
}
