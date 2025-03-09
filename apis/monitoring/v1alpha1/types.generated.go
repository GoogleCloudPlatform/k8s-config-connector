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

// +kcc:proto=google.monitoring.v3.Group
type Group struct {
	// Output only. The name of this group. The format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID]
	//
	//  When creating a group, this field is ignored and a new name is created
	//  consisting of the project specified in the call to `CreateGroup`
	//  and a unique `[GROUP_ID]` that is generated automatically.
	// +kcc:proto:field=google.monitoring.v3.Group.name
	Name *string `json:"name,omitempty"`

	// A user-assigned name for this group, used only for display purposes.
	// +kcc:proto:field=google.monitoring.v3.Group.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The name of the group's parent, if it has one. The format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/groups/[GROUP_ID]
	//
	//  For groups with no parent, `parent_name` is the empty string, `""`.
	// +kcc:proto:field=google.monitoring.v3.Group.parent_name
	ParentName *string `json:"parentName,omitempty"`

	// The filter used to determine which monitored resources belong to this
	//  group.
	// +kcc:proto:field=google.monitoring.v3.Group.filter
	Filter *string `json:"filter,omitempty"`

	// If true, the members of this group are considered to be a cluster.
	//  The system can perform additional analysis on groups that are clusters.
	// +kcc:proto:field=google.monitoring.v3.Group.is_cluster
	IsCluster *bool `json:"isCluster,omitempty"`
}
