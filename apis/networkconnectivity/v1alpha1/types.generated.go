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


// +kcc:proto=google.cloud.networkconnectivity.v1.Hub
type Hub struct {
	// Immutable. The name of the hub. Hub names must be unique. They use the
	//  following form:
	//      `projects/{project_number}/locations/global/hubs/{hub_id}`
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.name
	Name *string `json:"name,omitempty"`

	// Optional labels in key-value pair format. For more information about
	//  labels, see [Requirements for
	//  labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.labels
	Labels map[string]string `json:"labels,omitempty"`

	// An optional description of the hub.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.description
	Description *string `json:"description,omitempty"`

	// The VPC networks associated with this hub's spokes.
	//
	//  This field is read-only. Network Connectivity Center automatically
	//  populates it based on the set of spokes attached to the hub.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.routing_vpcs
	RoutingVpcs []RoutingVPC `json:"routingVpcs,omitempty"`

	// Optional. The policy mode of this hub. This field can be either
	//  PRESET or CUSTOM. If unspecified, the
	//  policy_mode defaults to PRESET.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.policy_mode
	PolicyMode *string `json:"policyMode,omitempty"`

	// Optional. The topology implemented in this hub. Currently, this field is
	//  only used when policy_mode = PRESET. The available preset topologies are
	//  MESH and STAR. If preset_topology is unspecified and policy_mode = PRESET,
	//  the preset_topology defaults to MESH. When policy_mode = CUSTOM,
	//  the preset_topology is set to PRESET_TOPOLOGY_UNSPECIFIED.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.preset_topology
	PresetTopology *string `json:"presetTopology,omitempty"`

	// Optional. Whether Private Service Connect transitivity is enabled for the
	//  hub. If true, Private Service Connect endpoints in VPC spokes attached to
	//  the hub are made accessible to other VPC spokes attached to the hub.
	//  The default value is false.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.export_psc
	ExportPsc *bool `json:"exportPsc,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.RoutingVPC
type RoutingVPC struct {
	// The URI of the VPC network.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.RoutingVPC.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.SpokeSummary
type SpokeSummary struct {
}

// +kcc:proto=google.cloud.networkconnectivity.v1.SpokeSummary.SpokeStateCount
type SpokeSummary_SpokeStateCount struct {
}

// +kcc:proto=google.cloud.networkconnectivity.v1.SpokeSummary.SpokeStateReasonCount
type SpokeSummary_SpokeStateReasonCount struct {
}

// +kcc:proto=google.cloud.networkconnectivity.v1.SpokeSummary.SpokeTypeCount
type SpokeSummary_SpokeTypeCount struct {
}

// +kcc:proto=google.cloud.networkconnectivity.v1.Hub
type HubObservedState struct {
	// Output only. The time the hub was created.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time the hub was last updated.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The Google-generated UUID for the hub. This value is unique
	//  across all hub resources. If a hub is deleted and another with the same
	//  name is created, the new hub is assigned a different unique_id.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.unique_id
	UniqueID *string `json:"uniqueID,omitempty"`

	// Output only. The current lifecycle state of this hub.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.state
	State *string `json:"state,omitempty"`

	// The VPC networks associated with this hub's spokes.
	//
	//  This field is read-only. Network Connectivity Center automatically
	//  populates it based on the set of spokes attached to the hub.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.routing_vpcs
	RoutingVpcs []RoutingVPCObservedState `json:"routingVpcs,omitempty"`

	// Output only. The route tables that belong to this hub. They use the
	//  following form:
	//     `projects/{project_number}/locations/global/hubs/{hub_id}/routeTables/{route_table_id}`
	//
	//  This field is read-only. Network Connectivity Center automatically
	//  populates it based on the route tables nested under the hub.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.route_tables
	RouteTables []string `json:"routeTables,omitempty"`

	// Output only. A summary of the spokes associated with a hub. The
	//  summary includes a count of spokes according to type
	//  and according to state. If any spokes are inactive,
	//  the summary also lists the reasons they are inactive,
	//  including a count for each reason.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Hub.spoke_summary
	SpokeSummary *SpokeSummary `json:"spokeSummary,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.RoutingVPC
type RoutingVPCObservedState struct {
	// Output only. If true, indicates that this VPC network is currently
	//  associated with spokes that use the data transfer feature (spokes where the
	//  site_to_site_data_transfer field is set to true). If you create new spokes
	//  that use data transfer, they must be associated with this VPC network. At
	//  most, one VPC network will have this field set to true.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.RoutingVPC.required_for_new_site_to_site_data_transfer_spokes
	RequiredForNewSiteToSiteDataTransferSpokes *bool `json:"requiredForNewSiteToSiteDataTransferSpokes,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.SpokeSummary
type SpokeSummaryObservedState struct {
	// Output only. Counts the number of spokes of each type that are
	//  associated with a specific hub.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.SpokeSummary.spoke_type_counts
	SpokeTypeCounts []SpokeSummary_SpokeTypeCount `json:"spokeTypeCounts,omitempty"`

	// Output only. Counts the number of spokes that are in each state
	//  and associated with a given hub.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.SpokeSummary.spoke_state_counts
	SpokeStateCounts []SpokeSummary_SpokeStateCount `json:"spokeStateCounts,omitempty"`

	// Output only. Counts the number of spokes that are inactive for each
	//  possible reason and associated with a given hub.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.SpokeSummary.spoke_state_reason_counts
	SpokeStateReasonCounts []SpokeSummary_SpokeStateReasonCount `json:"spokeStateReasonCounts,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.SpokeSummary.SpokeStateCount
type SpokeSummary_SpokeStateCountObservedState struct {
	// Output only. The state of the spokes.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.SpokeSummary.SpokeStateCount.state
	State *string `json:"state,omitempty"`

	// Output only. The total number of spokes that are in this state
	//  and associated with a given hub.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.SpokeSummary.SpokeStateCount.count
	Count *int64 `json:"count,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.SpokeSummary.SpokeStateReasonCount
type SpokeSummary_SpokeStateReasonCountObservedState struct {
	// Output only. The reason that a spoke is inactive.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.SpokeSummary.SpokeStateReasonCount.state_reason_code
	StateReasonCode *string `json:"stateReasonCode,omitempty"`

	// Output only. The total number of spokes that are inactive for a
	//  particular reason and associated with a given hub.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.SpokeSummary.SpokeStateReasonCount.count
	Count *int64 `json:"count,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.SpokeSummary.SpokeTypeCount
type SpokeSummary_SpokeTypeCountObservedState struct {
	// Output only. The type of the spokes.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.SpokeSummary.SpokeTypeCount.spoke_type
	SpokeType *string `json:"spokeType,omitempty"`

	// Output only. The total number of spokes of this type that are
	//  associated with the hub.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.SpokeSummary.SpokeTypeCount.count
	Count *int64 `json:"count,omitempty"`
}
