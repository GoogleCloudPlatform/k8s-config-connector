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


// +kcc:proto=google.cloud.tpu.v2.AcceleratorConfig
type AcceleratorConfig struct {
	// Required. Type of TPU.
	// +kcc:proto:field=google.cloud.tpu.v2.AcceleratorConfig.type
	Type *string `json:"type,omitempty"`

	// Required. Topology of TPU in chips.
	// +kcc:proto:field=google.cloud.tpu.v2.AcceleratorConfig.topology
	Topology *string `json:"topology,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.AccessConfig
type AccessConfig struct {
}

// +kcc:proto=google.cloud.tpu.v2.AttachedDisk
type AttachedDisk struct {
	// Specifies the full path to an existing disk.
	//  For example: "projects/my-project/zones/us-central1-c/disks/my-disk".
	// +kcc:proto:field=google.cloud.tpu.v2.AttachedDisk.source_disk
	SourceDisk *string `json:"sourceDisk,omitempty"`

	// The mode in which to attach this disk.
	//  If not specified, the default is READ_WRITE mode.
	//  Only applicable to data_disks.
	// +kcc:proto:field=google.cloud.tpu.v2.AttachedDisk.mode
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.NetworkConfig
type NetworkConfig struct {
	// The name of the network for the TPU node. It must be a preexisting Google
	//  Compute Engine network. If none is provided, "default" will be used.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkConfig.network
	Network *string `json:"network,omitempty"`

	// The name of the subnetwork for the TPU node. It must be a preexisting
	//  Google Compute Engine subnetwork. If none is provided, "default" will be
	//  used.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkConfig.subnetwork
	Subnetwork *string `json:"subnetwork,omitempty"`

	// Indicates that external IP addresses would be associated with the TPU
	//  workers. If set to false, the specified subnetwork or network should have
	//  Private Google Access enabled.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkConfig.enable_external_ips
	EnableExternalIps *bool `json:"enableExternalIps,omitempty"`

	// Allows the TPU node to send and receive packets with non-matching
	//  destination or source IPs. This is required if you plan to use the TPU
	//  workers to forward routes.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkConfig.can_ip_forward
	CanIPForward *bool `json:"canIPForward,omitempty"`

	// Optional. Specifies networking queue count for TPU VM instance's network
	//  interface.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkConfig.queue_count
	QueueCount *int32 `json:"queueCount,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.NetworkEndpoint
type NetworkEndpoint struct {
	// The internal IP address of this network endpoint.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkEndpoint.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// The port of this network endpoint.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkEndpoint.port
	Port *int32 `json:"port,omitempty"`

	// The access config for the TPU worker.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkEndpoint.access_config
	AccessConfig *AccessConfig `json:"accessConfig,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.Node
type Node struct {

	// The user-supplied description of the TPU. Maximum of 512 characters.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.description
	Description *string `json:"description,omitempty"`

	// Optional. The type of hardware accelerators associated with this node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// Required. The runtime version running in the Node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.runtime_version
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`

	// Network configurations for the TPU node. network_config and network_configs
	//  are mutually exclusive, you can only specify one of them. If both are
	//  specified, an error will be returned.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.network_config
	NetworkConfig *NetworkConfig `json:"networkConfig,omitempty"`

	// Optional. Repeated network configurations for the TPU node. This field is
	//  used to specify multiple networks configs for the TPU node. network_config
	//  and network_configs are mutually exclusive, you can only specify one of
	//  them. If both are specified, an error will be returned.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.network_configs
	NetworkConfigs []NetworkConfig `json:"networkConfigs,omitempty"`

	// The CIDR block that the TPU node will use when selecting an IP address.
	//  This CIDR block must be a /29 block; the Compute Engine networks API
	//  forbids a smaller block, and using a larger block would be wasteful (a
	//  node can only consume one IP address). Errors will occur if the CIDR block
	//  has already been used for a currently existing TPU node, the CIDR block
	//  conflicts with any subnetworks in the user's provided network, or the
	//  provided network is peered with another network that is using that CIDR
	//  block.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.cidr_block
	CidrBlock *string `json:"cidrBlock,omitempty"`

	// The Google Cloud Platform Service Account to be used by the TPU node VMs.
	//  If None is specified, the default compute service account will be used.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.service_account
	ServiceAccount *ServiceAccount `json:"serviceAccount,omitempty"`

	// The scheduling options for this node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.scheduling_config
	SchedulingConfig *SchedulingConfig `json:"schedulingConfig,omitempty"`

	// The health status of the TPU node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.health
	Health *string `json:"health,omitempty"`

	// Resource labels to represent user-provided metadata.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Custom metadata to apply to the TPU Node.
	//  Can set startup-script and shutdown-script
	// +kcc:proto:field=google.cloud.tpu.v2.Node.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Tags to apply to the TPU Node. Tags are used to identify valid sources or
	//  targets for network firewalls.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.tags
	Tags []string `json:"tags,omitempty"`

	// The additional data disks for the Node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.data_disks
	DataDisks []AttachedDisk `json:"dataDisks,omitempty"`

	// Shielded Instance options.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.shielded_instance_config
	ShieldedInstanceConfig *ShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	// The AccleratorConfig for the TPU Node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.accelerator_config
	AcceleratorConfig *AcceleratorConfig `json:"acceleratorConfig,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResource
type QueuedResource struct {

	// Optional. Defines a TPU resource.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.tpu
	Tpu *QueuedResource_Tpu `json:"tpu,omitempty"`

	// Optional. The Spot tier.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.spot
	Spot *QueuedResource_Spot `json:"spot,omitempty"`

	// Optional. The Guaranteed tier
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.guaranteed
	Guaranteed *QueuedResource_Guaranteed `json:"guaranteed,omitempty"`

	// Optional. The queueing policy of the QueuedRequest.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.queueing_policy
	QueueingPolicy *QueuedResource_QueueingPolicy `json:"queueingPolicy,omitempty"`

	// Optional. Name of the reservation in which the resource should be
	//  provisioned. Format:
	//  projects/{project}/locations/{zone}/reservations/{reservation}
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.reservation_name
	ReservationName *string `json:"reservationName,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResource.Guaranteed
type QueuedResource_Guaranteed struct {
	// Optional. Defines the minimum duration of the guarantee. If specified,
	//  the requested resources will only be provisioned if they can be
	//  allocated for at least the given duration.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.Guaranteed.min_duration
	MinDuration *string `json:"minDuration,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResource.QueueingPolicy
type QueuedResource_QueueingPolicy struct {
	// Optional. A relative time after which resources should not be created.
	//  If the request cannot be fulfilled by this time the request will be
	//  failed.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.QueueingPolicy.valid_until_duration
	ValidUntilDuration *string `json:"validUntilDuration,omitempty"`

	// Optional. An absolute time after which resources should not be created.
	//  If the request cannot be fulfilled by this time the request will be
	//  failed.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.QueueingPolicy.valid_until_time
	ValidUntilTime *string `json:"validUntilTime,omitempty"`

	// Optional. A relative time after which resources may be created.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.QueueingPolicy.valid_after_duration
	ValidAfterDuration *string `json:"validAfterDuration,omitempty"`

	// Optional. An absolute time after which resources may be created.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.QueueingPolicy.valid_after_time
	ValidAfterTime *string `json:"validAfterTime,omitempty"`

	// Optional. An absolute time interval within which resources may be
	//  created.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.QueueingPolicy.valid_interval
	ValidInterval *Interval `json:"validInterval,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResource.Spot
type QueuedResource_Spot struct {
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResource.Tpu
type QueuedResource_Tpu struct {
	// Optional. The TPU node(s) being requested.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.Tpu.node_spec
	NodeSpec []QueuedResource_Tpu_NodeSpec `json:"nodeSpec,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResource.Tpu.NodeSpec
type QueuedResource_Tpu_NodeSpec struct {
	// Required. The parent resource name.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.Tpu.NodeSpec.parent
	Parent *string `json:"parent,omitempty"`

	// Optional. The unqualified resource name. Should follow the
	//  `^[A-Za-z0-9_.~+%-]+$` regex format. This is only specified when
	//  requesting a single node. In case of multislice requests,
	//  multislice_params must be populated instead.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.Tpu.NodeSpec.node_id
	NodeID *string `json:"nodeID,omitempty"`

	// Optional. Fields to specify in case of multislice request.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.Tpu.NodeSpec.multislice_params
	MultisliceParams *QueuedResource_Tpu_NodeSpec_MultisliceParams `json:"multisliceParams,omitempty"`

	// Required. The node.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.Tpu.NodeSpec.node
	Node *Node `json:"node,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResource.Tpu.NodeSpec.MultisliceParams
type QueuedResource_Tpu_NodeSpec_MultisliceParams struct {
	// Required. Number of nodes with this spec. The system will attempt
	//  to provision "node_count" nodes as part of the request.
	//  This needs to be > 1.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.Tpu.NodeSpec.MultisliceParams.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Optional. Prefix of node_ids in case of multislice request.
	//  Should follow the `^[A-Za-z0-9_.~+%-]+$` regex format.
	//  If node_count = 3 and node_id_prefix = "np", node ids of nodes
	//  created will be "np-0", "np-1", "np-2". If this field is not
	//  provided we use queued_resource_id as the node_id_prefix.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.Tpu.NodeSpec.MultisliceParams.node_id_prefix
	NodeIDPrefix *string `json:"nodeIDPrefix,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResourceState
type QueuedResourceState struct {
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResourceState.AcceptedData
type QueuedResourceState_AcceptedData struct {
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResourceState.ActiveData
type QueuedResourceState_ActiveData struct {
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResourceState.CreatingData
type QueuedResourceState_CreatingData struct {
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResourceState.DeletingData
type QueuedResourceState_DeletingData struct {
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResourceState.FailedData
type QueuedResourceState_FailedData struct {
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResourceState.ProvisioningData
type QueuedResourceState_ProvisioningData struct {
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResourceState.SuspendedData
type QueuedResourceState_SuspendedData struct {
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResourceState.SuspendingData
type QueuedResourceState_SuspendingData struct {
}

// +kcc:proto=google.cloud.tpu.v2.SchedulingConfig
type SchedulingConfig struct {
	// Defines whether the node is preemptible.
	// +kcc:proto:field=google.cloud.tpu.v2.SchedulingConfig.preemptible
	Preemptible *bool `json:"preemptible,omitempty"`

	// Whether the node is created under a reservation.
	// +kcc:proto:field=google.cloud.tpu.v2.SchedulingConfig.reserved
	Reserved *bool `json:"reserved,omitempty"`

	// Optional. Defines whether the node is Spot VM.
	// +kcc:proto:field=google.cloud.tpu.v2.SchedulingConfig.spot
	Spot *bool `json:"spot,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.ServiceAccount
type ServiceAccount struct {
	// Email address of the service account. If empty, default Compute service
	//  account will be used.
	// +kcc:proto:field=google.cloud.tpu.v2.ServiceAccount.email
	Email *string `json:"email,omitempty"`

	// The list of scopes to be made available for this service account. If empty,
	//  access to all Cloud APIs will be allowed.
	// +kcc:proto:field=google.cloud.tpu.v2.ServiceAccount.scope
	Scope []string `json:"scope,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.ShieldedInstanceConfig
type ShieldedInstanceConfig struct {
	// Defines whether the instance has Secure Boot enabled.
	// +kcc:proto:field=google.cloud.tpu.v2.ShieldedInstanceConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.Symptom
type Symptom struct {
	// Timestamp when the Symptom is created.
	// +kcc:proto:field=google.cloud.tpu.v2.Symptom.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Type of the Symptom.
	// +kcc:proto:field=google.cloud.tpu.v2.Symptom.symptom_type
	SymptomType *string `json:"symptomType,omitempty"`

	// Detailed information of the current Symptom.
	// +kcc:proto:field=google.cloud.tpu.v2.Symptom.details
	Details *string `json:"details,omitempty"`

	// A string used to uniquely distinguish a worker within a TPU node.
	// +kcc:proto:field=google.cloud.tpu.v2.Symptom.worker_id
	WorkerID *string `json:"workerID,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.type.Interval
type Interval struct {
	// Optional. Inclusive start of the interval.
	//
	//  If specified, a Timestamp matching this interval will have to be the same
	//  or after the start.
	// +kcc:proto:field=google.type.Interval.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. Exclusive end of the interval.
	//
	//  If specified, a Timestamp matching this interval will have to be before the
	//  end.
	// +kcc:proto:field=google.type.Interval.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.AccessConfig
type AccessConfigObservedState struct {
	// Output only. An external IP address associated with the TPU worker.
	// +kcc:proto:field=google.cloud.tpu.v2.AccessConfig.external_ip
	ExternalIP *string `json:"externalIP,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.NetworkEndpoint
type NetworkEndpointObservedState struct {
	// The access config for the TPU worker.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkEndpoint.access_config
	AccessConfig *AccessConfigObservedState `json:"accessConfig,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.Node
type NodeObservedState struct {
	// Output only. Immutable. The name of the TPU.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.name
	Name *string `json:"name,omitempty"`

	// Output only. The current state for the TPU Node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.state
	State *string `json:"state,omitempty"`

	// Output only. If this field is populated, it contains a description of why
	//  the TPU Node is unhealthy.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.health_description
	HealthDescription *string `json:"healthDescription,omitempty"`

	// Output only. The time when the node was created.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The network endpoints where TPU workers can be accessed and
	//  sent work. It is recommended that runtime clients of the node reach out
	//  to the 0th entry in this map first.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.network_endpoints
	NetworkEndpoints []NetworkEndpoint `json:"networkEndpoints,omitempty"`

	// Output only. The unique identifier for the TPU Node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.id
	ID *int64 `json:"id,omitempty"`

	// Output only. The API version that created this Node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.api_version
	ApiVersion *string `json:"apiVersion,omitempty"`

	// Output only. The Symptoms that have occurred to the TPU Node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.symptoms
	Symptoms []Symptom `json:"symptoms,omitempty"`

	// Output only. The qualified name of the QueuedResource that requested this
	//  Node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.queued_resource
	QueuedResource *string `json:"queuedResource,omitempty"`

	// Output only. Whether the Node belongs to a Multislice group.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.multislice_node
	MultisliceNode *bool `json:"multisliceNode,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResource
type QueuedResourceObservedState struct {
	// Output only. Immutable. The name of the QueuedResource.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.name
	Name *string `json:"name,omitempty"`

	// Output only. The time when the QueuedResource was created.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. Defines a TPU resource.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.tpu
	Tpu *QueuedResource_TpuObservedState `json:"tpu,omitempty"`

	// Output only. State of the QueuedResource request.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.state
	State *QueuedResourceState `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResource.Tpu
type QueuedResource_TpuObservedState struct {
	// Optional. The TPU node(s) being requested.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.Tpu.node_spec
	NodeSpec []QueuedResource_Tpu_NodeSpecObservedState `json:"nodeSpec,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResource.Tpu.NodeSpec
type QueuedResource_Tpu_NodeSpecObservedState struct {
	// Required. The node.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.Tpu.NodeSpec.node
	Node *NodeObservedState `json:"node,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResourceState
type QueuedResourceStateObservedState struct {
	// Output only. State of the QueuedResource request.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResourceState.state
	State *string `json:"state,omitempty"`

	// Output only. Further data for the creating state.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResourceState.creating_data
	CreatingData *QueuedResourceState_CreatingData `json:"creatingData,omitempty"`

	// Output only. Further data for the accepted state.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResourceState.accepted_data
	AcceptedData *QueuedResourceState_AcceptedData `json:"acceptedData,omitempty"`

	// Output only. Further data for the provisioning state.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResourceState.provisioning_data
	ProvisioningData *QueuedResourceState_ProvisioningData `json:"provisioningData,omitempty"`

	// Output only. Further data for the failed state.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResourceState.failed_data
	FailedData *QueuedResourceState_FailedData `json:"failedData,omitempty"`

	// Output only. Further data for the deleting state.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResourceState.deleting_data
	DeletingData *QueuedResourceState_DeletingData `json:"deletingData,omitempty"`

	// Output only. Further data for the active state.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResourceState.active_data
	ActiveData *QueuedResourceState_ActiveData `json:"activeData,omitempty"`

	// Output only. Further data for the suspending state.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResourceState.suspending_data
	SuspendingData *QueuedResourceState_SuspendingData `json:"suspendingData,omitempty"`

	// Output only. Further data for the suspended state.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResourceState.suspended_data
	SuspendedData *QueuedResourceState_SuspendedData `json:"suspendedData,omitempty"`

	// Output only. The initiator of the QueuedResources's current state. Used to
	//  indicate whether the SUSPENDING/SUSPENDED state was initiated by the user
	//  or the service.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResourceState.state_initiator
	StateInitiator *string `json:"stateInitiator,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.QueuedResourceState.FailedData
type QueuedResourceState_FailedDataObservedState struct {
	// Output only. The error that caused the queued resource to enter the
	//  FAILED state.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResourceState.FailedData.error
	Error *Status `json:"error,omitempty"`
}
