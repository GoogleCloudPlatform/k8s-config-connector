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


// +kcc:proto=google.cloud.migrationcenter.v1.Group
type Group struct {

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Group.labels
	Labels map[string]string `json:"labels,omitempty"`

	// User-friendly display name.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Group.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the resource.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Group.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.Group
type GroupObservedState struct {
	// Output only. The name of the group.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Group.name
	Name *string `json:"name,omitempty"`

	// Output only. The timestamp when the group was created.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Group.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the group was last updated.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Group.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
