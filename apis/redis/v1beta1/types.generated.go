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

package v1beta1

// +kcc:proto=google.cloud.redis.cluster.v1.Cluster.StateInfo
type Cluster_StateInfo struct {
	// Describes ongoing update on the cluster when cluster state is UPDATING.
	UpdateInfo *Cluster_StateInfo_UpdateInfo `json:"updateInfo,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1.Cluster.StateInfo.UpdateInfo
type Cluster_StateInfo_UpdateInfo struct {
	// Target number of shards for redis cluster
	TargetShardCount *int32 `json:"targetShardCount,omitempty"`

	// Target number of replica nodes per shard.
	TargetReplicaCount *int32 `json:"targetReplicaCount,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1.ClusterPersistenceConfig
type ClusterPersistenceConfig struct {
	// Optional. The mode of persistence.
	Mode *string `json:"mode,omitempty"`

	// Optional. RDB configuration. This field will be ignored if mode is not RDB.
	RdbConfig *ClusterPersistenceConfig_RDBConfig `json:"rdbConfig,omitempty"`

	// Optional. AOF configuration. This field will be ignored if mode is not AOF.
	AofConfig *ClusterPersistenceConfig_AOFConfig `json:"aofConfig,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1.ClusterPersistenceConfig.AOFConfig
type ClusterPersistenceConfig_AOFConfig struct {
	// Optional. fsync configuration.
	AppendFsync *string `json:"appendFsync,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1.ClusterPersistenceConfig.RDBConfig
type ClusterPersistenceConfig_RDBConfig struct {
	// Optional. Period between RDB snapshots.
	RdbSnapshotPeriod *string `json:"rdbSnapshotPeriod,omitempty"`

	// Optional. The time that the first snapshot was/will be attempted, and to
	//  which future snapshots will be aligned. If not provided, the current time
	//  will be used.
	RdbSnapshotStartTime *string `json:"rdbSnapshotStartTime,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1.DiscoveryEndpoint
type DiscoveryEndpoint struct {
	// Output only. Address of the exposed Redis endpoint used by clients to
	//  connect to the service. The address could be either IP or hostname.
	Address *string `json:"address,omitempty"`

	// Output only. The port number of the exposed Redis endpoint.
	Port *int32 `json:"port,omitempty"`

	// Output only. Customer configuration for where the endpoint is created and
	//  accessed from.
	PscConfig *PscConfig `json:"pscConfig,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1.PscConfig
type PscConfig struct {
	// Required. The network where the IP address of the discovery endpoint will
	//  be reserved, in the form of
	//  projects/{network_project}/global/networks/{network_id}.
	Network *string `json:"network,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1.PscConnection
type PscConnection struct {
	// Output only. The PSC connection id of the forwarding rule connected to the
	//  service attachment.
	PscConnectionID *string `json:"pscConnectionID,omitempty"`

	// Output only. The IP allocated on the consumer network for the PSC
	//  forwarding rule.
	Address *string `json:"address,omitempty"`

	// Output only. The URI of the consumer side forwarding rule.
	//  Example:
	//  projects/{projectNumOrId}/regions/us-east1/forwardingRules/{resourceId}.
	ForwardingRule *string `json:"forwardingRule,omitempty"`

	// Output only. The consumer project_id where the forwarding rule is created
	//  from.
	ProjectID *string `json:"projectID,omitempty"`

	// The consumer network where the IP address resides, in the form of
	//  projects/{project_id}/global/networks/{network_id}.
	Network *string `json:"network,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1.ZoneDistributionConfig
type ZoneDistributionConfig struct {
	// Optional. The mode of zone distribution. Defaults to MULTI_ZONE, when not
	//  specified.
	Mode *string `json:"mode,omitempty"`

	// Optional. When SINGLE ZONE distribution is selected, zone field would be
	//  used to allocate all resources in that zone. This is not applicable to
	//  MULTI_ZONE, and would be ignored for MULTI_ZONE clusters.
	Zone *string `json:"zone,omitempty"`
}
