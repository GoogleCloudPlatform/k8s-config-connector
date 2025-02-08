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


// +kcc:proto=google.cloud.networkconnectivity.v1.RouteTable
type RouteTable struct {
	// Immutable. The name of the route table. Route table names must be unique.
	//  They use the following form:
	//       `projects/{project_number}/locations/global/hubs/{hub}/routeTables/{route_table_id}`
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.RouteTable.name
	Name *string `json:"name,omitempty"`

	// Optional labels in key-value pair format. For more information about
	//  labels, see [Requirements for
	//  labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.RouteTable.labels
	Labels map[string]string `json:"labels,omitempty"`

	// An optional description of the route table.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.RouteTable.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.RouteTable
type RouteTableObservedState struct {
	// Output only. The time the route table was created.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.RouteTable.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time the route table was last updated.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.RouteTable.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The Google-generated UUID for the route table. This value is
	//  unique across all route table resources. If a route table is deleted and
	//  another with the same name is created, the new route table is assigned
	//  a different `uid`.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.RouteTable.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The current lifecycle state of this route table.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.RouteTable.state
	State *string `json:"state,omitempty"`
}
