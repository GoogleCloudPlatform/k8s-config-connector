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


// +kcc:proto=google.cloud.resourcemanager.v3.Project
type Project struct {

	// Optional. A reference to a parent Resource. eg., `organizations/123` or
	//  `folders/876`.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.parent
	Parent *string `json:"parent,omitempty"`

	// Immutable. The unique, user-assigned id of the project.
	//  It must be 6 to 30 lowercase ASCII letters, digits, or hyphens.
	//  It must start with a letter.
	//  Trailing hyphens are prohibited.
	//
	//  Example: `tokyo-rain-123`
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Optional. A user-assigned display name of the project.
	//  When present it must be between 4 to 30 characters.
	//  Allowed characters are: lowercase and uppercase letters, numbers,
	//  hyphen, single-quote, double-quote, space, and exclamation point.
	//
	//  Example: `My Project`
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The labels associated with this project.
	//
	//  Label keys must be between 1 and 63 characters long and must conform
	//  to the following regular expression: \[a-z\](\[-a-z0-9\]*\[a-z0-9\])?.
	//
	//  Label values must be between 0 and 63 characters long and must conform
	//  to the regular expression (\[a-z\](\[-a-z0-9\]*\[a-z0-9\])?)?.
	//
	//  No more than 64 labels can be associated with a given resource.
	//
	//  Clients should store labels in a representation such as JSON that does not
	//  depend on specific characters being disallowed.
	//
	//  Example: `"myBusinessDimension" : "businessValue"`
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.resourcemanager.v3.Project
type ProjectObservedState struct {
	// Output only. The unique resource name of the project. It is an int64
	//  generated number prefixed by "projects/".
	//
	//  Example: `projects/415104041262`
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.name
	Name *string `json:"name,omitempty"`

	// Output only. The project lifecycle state.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.state
	State *string `json:"state,omitempty"`

	// Output only. Creation time.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time this resource was modified.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time at which this resource was requested for deletion.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. A checksum computed by the server based on the current value
	//  of the Project resource. This may be sent on update and delete requests to
	//  ensure the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.etag
	Etag *string `json:"etag,omitempty"`
}
