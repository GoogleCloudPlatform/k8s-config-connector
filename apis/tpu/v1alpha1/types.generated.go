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


// +kcc:proto=google.cloud.tpu.v1.NetworkEndpoint
type NetworkEndpoint struct {
	// The IP address of this network endpoint.
	// +kcc:proto:field=google.cloud.tpu.v1.NetworkEndpoint.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// The port of this network endpoint.
	// +kcc:proto:field=google.cloud.tpu.v1.NetworkEndpoint.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v1.Node
type Node struct {

	// The user-supplied description of the TPU. Maximum of 512 characters.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.description
	Description *string `json:"description,omitempty"`

	// Required. The type of hardware accelerators associated with this node.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// Output only. DEPRECATED! Use network_endpoints instead.
	//  The network address for the TPU Node as visible to Compute Engine
	//  instances.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// Output only. DEPRECATED! Use network_endpoints instead.
	//  The network port for the TPU Node as visible to Compute Engine instances.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.port
	Port *string `json:"port,omitempty"`

	// Required. The version of Tensorflow running in the Node.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.tensorflow_version
	TensorflowVersion *string `json:"tensorflowVersion,omitempty"`

	// The name of a network they wish to peer the TPU node to. It must be a
	//  preexisting Compute Engine network inside of the project on which this API
	//  has been activated. If none is provided, "default" will be used.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.network
	Network *string `json:"network,omitempty"`

	// The CIDR block that the TPU node will use when selecting an IP address.
	//  This CIDR block must be a /29 block; the Compute Engine networks API
	//  forbids a smaller block, and using a larger block would be wasteful (a
	//  node can only consume one IP address). Errors will occur if the CIDR block
	//  has already been used for a currently existing TPU node, the CIDR block
	//  conflicts with any subnetworks in the user's provided network, or the
	//  provided network is peered with another network that is using that CIDR
	//  block.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.cidr_block
	CidrBlock *string `json:"cidrBlock,omitempty"`

	// The scheduling options for this node.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.scheduling_config
	SchedulingConfig *SchedulingConfig `json:"schedulingConfig,omitempty"`

	// The health status of the TPU node.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.health
	Health *string `json:"health,omitempty"`

	// Resource labels to represent user-provided metadata.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Whether the VPC peering for the node is set up through Service Networking
	//  API. The VPC Peering should be set up before provisioning the node.
	//  If this field is set, cidr_block field should not be specified. If the
	//  network, that you want to peer the TPU Node to, is Shared VPC networks,
	//  the node must be created with this this field enabled.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.use_service_networking
	UseServiceNetworking *bool `json:"useServiceNetworking,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v1.SchedulingConfig
type SchedulingConfig struct {
	// Defines whether the node is preemptible.
	// +kcc:proto:field=google.cloud.tpu.v1.SchedulingConfig.preemptible
	Preemptible *bool `json:"preemptible,omitempty"`

	// Whether the node is created under a reservation.
	// +kcc:proto:field=google.cloud.tpu.v1.SchedulingConfig.reserved
	Reserved *bool `json:"reserved,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v1.Symptom
type Symptom struct {
	// Timestamp when the Symptom is created.
	// +kcc:proto:field=google.cloud.tpu.v1.Symptom.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Type of the Symptom.
	// +kcc:proto:field=google.cloud.tpu.v1.Symptom.symptom_type
	SymptomType *string `json:"symptomType,omitempty"`

	// Detailed information of the current Symptom.
	// +kcc:proto:field=google.cloud.tpu.v1.Symptom.details
	Details *string `json:"details,omitempty"`

	// A string used to uniquely distinguish a worker within a TPU node.
	// +kcc:proto:field=google.cloud.tpu.v1.Symptom.worker_id
	WorkerID *string `json:"workerID,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v1.Node
type NodeObservedState struct {
	// Output only. Immutable. The name of the TPU
	// +kcc:proto:field=google.cloud.tpu.v1.Node.name
	Name *string `json:"name,omitempty"`

	// Output only. The current state for the TPU Node.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.state
	State *string `json:"state,omitempty"`

	// Output only. If this field is populated, it contains a description of why
	//  the TPU Node is unhealthy.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.health_description
	HealthDescription *string `json:"healthDescription,omitempty"`

	// Output only. The service account used to run the tensor flow services
	//  within the node. To share resources, including Google Cloud Storage data,
	//  with the Tensorflow job running in the Node, this account must have
	//  permissions to that data.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Output only. The time when the node was created.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The network endpoints where TPU workers can be accessed and
	//  sent work. It is recommended that Tensorflow clients of the node reach out
	//  to the 0th entry in this map first.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.network_endpoints
	NetworkEndpoints []NetworkEndpoint `json:"networkEndpoints,omitempty"`

	// Output only. The API version that created this Node.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.api_version
	ApiVersion *string `json:"apiVersion,omitempty"`

	// Output only. The Symptoms that have occurred to the TPU Node.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.symptoms
	Symptoms []Symptom `json:"symptoms,omitempty"`
}
