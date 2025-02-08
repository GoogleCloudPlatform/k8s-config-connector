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


// +kcc:proto=google.cloud.resourcemanager.v3.Organization
type Organization struct {

	// Immutable. The G Suite / Workspace customer id used in the Directory API.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Organization.directory_customer_id
	DirectoryCustomerID *string `json:"directoryCustomerID,omitempty"`
}

// +kcc:proto=google.cloud.resourcemanager.v3.Organization
type OrganizationObservedState struct {
	// Output only. The resource name of the organization. This is the
	//  organization's relative path in the API. Its format is
	//  "organizations/[organization_id]". For example, "organizations/1234".
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Organization.name
	Name *string `json:"name,omitempty"`

	// Output only. A human-readable string that refers to the organization in the
	//  Google Cloud Console. This string is set by the server and cannot be
	//  changed. The string will be set to the primary domain (for example,
	//  "google.com") of the Google Workspace customer that owns the organization.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Organization.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. The organization's current lifecycle state.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Organization.state
	State *string `json:"state,omitempty"`

	// Output only. Timestamp when the Organization was created.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Organization.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the Organization was last modified.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Organization.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Timestamp when the Organization was requested for deletion.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Organization.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. A checksum computed by the server based on the current value
	//  of the Organization resource. This may be sent on update and delete
	//  requests to ensure the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Organization.etag
	Etag *string `json:"etag,omitempty"`
}
