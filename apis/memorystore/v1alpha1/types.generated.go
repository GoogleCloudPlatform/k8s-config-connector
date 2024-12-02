// Copyright 2024 Google LLC
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

// +kcc:proto=google.cloud.memorystore.v1beta.DiscoveryEndpoint
type DiscoveryEndpoint struct {
	// Output only. IP address of the exposed endpoint clients connect to.
	Address *string `json:"address,omitempty"`

	// Output only. The port number of the exposed endpoint.
	Port *int32 `json:"port,omitempty"`

	// Output only. The network where the IP address of the discovery endpoint
	//  will be reserved, in the form of
	//  projects/{network_project}/global/networks/{network_id}.
	Network *string `json:"network,omitempty"`
}

// +kcc:proto=google.cloud.memorystore.v1beta.Instance
type Instance struct {
	// Identifier. Unique name of the instance.
	//  Format: projects/{project}/locations/{location}/instances/{instance}
	Name *string `json:"name,omitempty"`

	// Output only. Creation timestamp of the instance.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Latest update timestamp of the instance.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Labels to represent user-provided metadata.
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. Current state of the instance.
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the state of the instance.
	StateInfo *Instance_StateInfo `json:"stateInfo,omitempty"`

	// Output only. System assigned, unique identifier for the instance.
	Uid *string `json:"uid,omitempty"`

	// Optional. Number of replica nodes per shard. If omitted the default is 0
	//  replicas.
	ReplicaCount *int32 `json:"replicaCount,omitempty"`

	// Optional. Immutable. Authorization mode of the instance.
	AuthorizationMode *string `json:"authorizationMode,omitempty"`

	// Optional. Immutable. In-transit encryption mode of the instance.
	TransitEncryptionMode *string `json:"transitEncryptionMode,omitempty"`

	// Optional. Number of shards for the instance.
	ShardCount *int32 `json:"shardCount,omitempty"`

	// Output only. Endpoints clients can connect to the instance through.
	//  Currently only one discovery endpoint is supported.
	DiscoveryEndpoints []DiscoveryEndpoint `json:"discoveryEndpoints,omitempty"`

	// Optional. Immutable. Machine type for individual nodes of the instance.
	NodeType *string `json:"nodeType,omitempty"`

	// Optional. Persistence configuration of the instance.
	PersistenceConfig *PersistenceConfig `json:"persistenceConfig,omitempty"`

	// Optional. Immutable. Engine version of the instance.
	EngineVersion *string `json:"engineVersion,omitempty"`

	// Optional. User-provided engine configurations for the instance.
	EngineConfigs map[string]string `json:"engineConfigs,omitempty"`

	// Output only. Configuration of individual nodes of the instance.
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	// Optional. Immutable. Zone distribution configuration of the instance for
	//  node allocation.
	ZoneDistributionConfig *ZoneDistributionConfig `json:"zoneDistributionConfig,omitempty"`

	// Optional. If set to true deletion of the instance will fail.
	DeletionProtectionEnabled *bool `json:"deletionProtectionEnabled,omitempty"`

	// Required. Immutable. User inputs and resource details of the auto-created
	//  PSC connections.
	PscAutoConnections []PscAutoConnection `json:"pscAutoConnections,omitempty"`
}

// +kcc:proto=google.cloud.memorystore.v1beta.Instance.StateInfo
type Instance_StateInfo struct {
	// Output only. Describes ongoing update when instance state is UPDATING.
	UpdateInfo *Instance_StateInfo_UpdateInfo `json:"updateInfo,omitempty"`
}

// +kcc:proto=google.cloud.memorystore.v1beta.Instance.StateInfo.UpdateInfo
type Instance_StateInfo_UpdateInfo struct {
	// Output only. Target number of shards for the instance.
	TargetShardCount *int32 `json:"targetShardCount,omitempty"`

	// Output only. Target number of replica nodes per shard for the instance.
	TargetReplicaCount *int32 `json:"targetReplicaCount,omitempty"`
}

// +kcc:proto=google.cloud.memorystore.v1beta.NodeConfig
type NodeConfig struct {
	// Output only. Memory size in GB of the node.
	SizeGb *float64 `json:"sizeGb,omitempty"`
}

// +kcc:proto=google.cloud.memorystore.v1beta.PersistenceConfig
type PersistenceConfig struct {
	// Optional. Current persistence mode.
	Mode *string `json:"mode,omitempty"`

	// Optional. RDB configuration. This field will be ignored if mode is not RDB.
	RdbConfig *PersistenceConfig_RDBConfig `json:"rdbConfig,omitempty"`

	// Optional. AOF configuration. This field will be ignored if mode is not AOF.
	AofConfig *PersistenceConfig_AOFConfig `json:"aofConfig,omitempty"`
}

// +kcc:proto=google.cloud.memorystore.v1beta.PersistenceConfig.AOFConfig
type PersistenceConfig_AOFConfig struct {
	// Optional. The fsync mode.
	AppendFsync *string `json:"appendFsync,omitempty"`
}

// +kcc:proto=google.cloud.memorystore.v1beta.PersistenceConfig.RDBConfig
type PersistenceConfig_RDBConfig struct {
	// Optional. Period between RDB snapshots.
	RdbSnapshotPeriod *string `json:"rdbSnapshotPeriod,omitempty"`

	// Optional. Time that the first snapshot was/will be attempted, and to
	//  which future snapshots will be aligned. If not provided, the current time
	//  will be used.
	RdbSnapshotStartTime *string `json:"rdbSnapshotStartTime,omitempty"`
}

// +kcc:proto=google.cloud.memorystore.v1beta.PscAutoConnection
type PscAutoConnection struct {
	// Optional. Output only. port will only be set for Primary/Reader or
	//  Discovery endpoint.
	Port *int32 `json:"port,omitempty"`

	// Output only. The PSC connection id of the forwarding rule connected to the
	//  service attachment.
	PscConnectionID *string `json:"pscConnectionID,omitempty"`

	// Output only. The IP allocated on the consumer network for the PSC
	//  forwarding rule.
	IpAddress *string `json:"ipAddress,omitempty"`

	// Output only. The URI of the consumer side forwarding rule.
	//  Format:
	//  projects/{project}/regions/{region}/forwardingRules/{forwarding_rule}
	ForwardingRule *string `json:"forwardingRule,omitempty"`

	// Required. The consumer project_id where PSC connections are established.
	//  This should be the same project_id that the cluster is being created in.
	ProjectID *string `json:"projectID,omitempty"`

	// Required. The network where the PSC endpoints are created, in the form of
	//  projects/{project_id}/global/networks/{network_id}.
	Network *string `json:"network,omitempty"`

	// Output only. The service attachment which is the target of the PSC
	//  connection, in the form of
	//  projects/{project-id}/regions/{region}/serviceAttachments/{service-attachment-id}.
	ServiceAttachment *string `json:"serviceAttachment,omitempty"`

	// Output only. The status of the PSC connection: whether a connection exists
	//  and ACTIVE or it no longer exists. Please note that this value is updated
	//  periodically. Please use Private Service Connect APIs for the latest
	//  status.
	PscConnectionStatus *string `json:"pscConnectionStatus,omitempty"`

	// Output only. Type of the PSC connection.
	ConnectionType *string `json:"connectionType,omitempty"`
}

// +kcc:proto=google.cloud.memorystore.v1beta.ZoneDistributionConfig
type ZoneDistributionConfig struct {
	// Optional. Defines zone where all resources will be allocated with
	//  SINGLE_ZONE mode. Ignored for MULTI_ZONE mode.
	Zone *string `json:"zone,omitempty"`

	// Optional. Current zone distribution mode. Defaults to MULTI_ZONE.
	Mode *string `json:"mode,omitempty"`
}
