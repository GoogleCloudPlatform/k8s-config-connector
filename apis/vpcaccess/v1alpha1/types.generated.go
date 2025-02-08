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


// +kcc:proto=google.cloud.vpcaccess.v1.Connector
type Connector struct {
	// The resource name in the format `projects/*/locations/*/connectors/*`.
	// +kcc:proto:field=google.cloud.vpcaccess.v1.Connector.name
	Name *string `json:"name,omitempty"`

	// Name of a VPC network.
	// +kcc:proto:field=google.cloud.vpcaccess.v1.Connector.network
	Network *string `json:"network,omitempty"`

	// The range of internal addresses that follows RFC 4632 notation.
	//  Example: `10.132.0.0/28`.
	// +kcc:proto:field=google.cloud.vpcaccess.v1.Connector.ip_cidr_range
	IPCidrRange *string `json:"ipCidrRange,omitempty"`

	// Minimum throughput of the connector in Mbps. Default and min is 200.
	// +kcc:proto:field=google.cloud.vpcaccess.v1.Connector.min_throughput
	MinThroughput *int32 `json:"minThroughput,omitempty"`

	// Maximum throughput of the connector in Mbps. Default is 300, max is 1000.
	// +kcc:proto:field=google.cloud.vpcaccess.v1.Connector.max_throughput
	MaxThroughput *int32 `json:"maxThroughput,omitempty"`

	// The subnet in which to house the VPC Access Connector.
	// +kcc:proto:field=google.cloud.vpcaccess.v1.Connector.subnet
	Subnet *Connector_Subnet `json:"subnet,omitempty"`

	// Machine type of VM Instance underlying connector. Default is e2-micro
	// +kcc:proto:field=google.cloud.vpcaccess.v1.Connector.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Minimum value of instances in autoscaling group underlying the connector.
	// +kcc:proto:field=google.cloud.vpcaccess.v1.Connector.min_instances
	MinInstances *int32 `json:"minInstances,omitempty"`

	// Maximum value of instances in autoscaling group underlying the connector.
	// +kcc:proto:field=google.cloud.vpcaccess.v1.Connector.max_instances
	MaxInstances *int32 `json:"maxInstances,omitempty"`
}

// +kcc:proto=google.cloud.vpcaccess.v1.Connector.Subnet
type Connector_Subnet struct {
	// Subnet name (relative, not fully qualified).
	//  E.g. if the full subnet selfLink is
	//  https://compute.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetName}
	//  the correct input for this field would be {subnetName}
	// +kcc:proto:field=google.cloud.vpcaccess.v1.Connector.Subnet.name
	Name *string `json:"name,omitempty"`

	// Project in which the subnet exists.
	//  If not set, this project is assumed to be the project for which
	//  the connector create request was issued.
	// +kcc:proto:field=google.cloud.vpcaccess.v1.Connector.Subnet.project_id
	ProjectID *string `json:"projectID,omitempty"`
}

// +kcc:proto=google.cloud.vpcaccess.v1.Connector
type ConnectorObservedState struct {
	// Output only. State of the VPC access connector.
	// +kcc:proto:field=google.cloud.vpcaccess.v1.Connector.state
	State *string `json:"state,omitempty"`

	// Output only. List of projects using the connector.
	// +kcc:proto:field=google.cloud.vpcaccess.v1.Connector.connected_projects
	ConnectedProjects []string `json:"connectedProjects,omitempty"`
}
