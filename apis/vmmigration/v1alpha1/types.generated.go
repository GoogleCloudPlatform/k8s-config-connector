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


// +kcc:proto=google.cloud.vmmigration.v1.Group
type Group struct {

	// User-provided description of the group.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Group.description
	Description *string `json:"description,omitempty"`

	// Display name is a user defined name for this group which can be updated.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Group.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.Group
type GroupObservedState struct {
	// Output only. The Group name.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Group.name
	Name *string `json:"name,omitempty"`

	// Output only. The create time timestamp.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Group.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time timestamp.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Group.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
