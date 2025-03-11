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

// +kcc:proto=google.cloud.datastream.v1.Error
type Error struct {
	// A title that explains the reason for the error.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.reason
	Reason *string `json:"reason,omitempty"`

	// A unique identifier for this specific error,
	//  allowing it to be traced throughout the system in logs and API responses.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.error_uuid
	ErrorUuid *string `json:"errorUuid,omitempty"`

	// A message containing more information about the error that occurred.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.message
	Message *string `json:"message,omitempty"`

	// The time when the error occurred.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.error_time
	ErrorTime *string `json:"errorTime,omitempty"`

	// Additional information about the error.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.details
	Details map[string]string `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PrivateConnection
type PrivateConnection struct {

	// Labels.
	// +kcc:proto:field=google.cloud.datastream.v1.PrivateConnection.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Display name.
	// +kcc:proto:field=google.cloud.datastream.v1.PrivateConnection.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// VPC Peering Config.
	// +kcc:proto:field=google.cloud.datastream.v1.PrivateConnection.vpc_peering_config
	VpcPeeringConfig *VpcPeeringConfig `json:"vpcPeeringConfig,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.VpcPeeringConfig
type VpcPeeringConfig struct {
	// Required. Fully qualified name of the VPC that Datastream will peer to.
	//  Format: `projects/{project}/global/{networks}/{name}`
	// +kcc:proto:field=google.cloud.datastream.v1.VpcPeeringConfig.vpc
	Vpc *string `json:"vpc,omitempty"`

	// Required. A free subnet for peering. (CIDR of /29)
	// +kcc:proto:field=google.cloud.datastream.v1.VpcPeeringConfig.subnet
	Subnet *string `json:"subnet,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PrivateConnection
type PrivateConnectionObservedState struct {
	// Output only. The resource's name.
	// +kcc:proto:field=google.cloud.datastream.v1.PrivateConnection.name
	Name *string `json:"name,omitempty"`

	// Output only. The create time of the resource.
	// +kcc:proto:field=google.cloud.datastream.v1.PrivateConnection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time of the resource.
	// +kcc:proto:field=google.cloud.datastream.v1.PrivateConnection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the Private Connection.
	// +kcc:proto:field=google.cloud.datastream.v1.PrivateConnection.state
	State *string `json:"state,omitempty"`

	// Output only. In case of error, the details of the error in a user-friendly
	//  format.
	// +kcc:proto:field=google.cloud.datastream.v1.PrivateConnection.error
	Error *Error `json:"error,omitempty"`
}
