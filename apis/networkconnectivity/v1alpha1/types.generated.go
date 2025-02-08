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


// +kcc:proto=google.cloud.networkconnectivity.v1.AutoAccept
type AutoAccept struct {
	// A list of project ids or project numbers for which you want
	//  to enable auto-accept. The auto-accept setting is applied to
	//  spokes being created or updated in these projects.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.AutoAccept.auto_accept_projects
	AutoAcceptProjects []string `json:"autoAcceptProjects,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.Group
type Group struct {
	// Immutable. The name of the group. Group names must be unique. They
	//  use the following form:
	//       `projects/{project_number}/locations/global/hubs/{hub}/groups/{group_id}`
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Group.name
	Name *string `json:"name,omitempty"`

	// Optional. Labels in key-value pair format. For more information about
	//  labels, see [Requirements for
	//  labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Group.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The description of the group.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Group.description
	Description *string `json:"description,omitempty"`

	// Optional. The auto-accept setting for this group.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Group.auto_accept
	AutoAccept *AutoAccept `json:"autoAccept,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.Group
type GroupObservedState struct {
	// Output only. The time the group was created.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Group.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time the group was last updated.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Group.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The Google-generated UUID for the group. This value is unique
	//  across all group resources. If a group is deleted and
	//  another with the same name is created, the new route table is assigned
	//  a different unique_id.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Group.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The current lifecycle state of this group.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Group.state
	State *string `json:"state,omitempty"`

	// Output only. The name of the route table that corresponds to this group.
	//  They use the following form:
	//  `projects/{project_number}/locations/global/hubs/{hub_id}/routeTables/{route_table_id}`
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Group.route_table
	RouteTable *string `json:"routeTable,omitempty"`
}
