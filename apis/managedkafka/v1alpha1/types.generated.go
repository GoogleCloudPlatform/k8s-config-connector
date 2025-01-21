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

// +kcc:proto=google.cloud.managedkafka.v1.AccessConfig
type AccessConfig struct {
	// Required. Virtual Private Cloud (VPC) networks that must be granted direct
	//  access to the Kafka cluster. Minimum of 1 network is required. Maximum 10
	//  networks can be specified.
	// +kcc:proto:field=google.cloud.managedkafka.v1.AccessConfig.network_configs
	NetworkConfigs []NetworkConfig `json:"networkConfigs,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.CapacityConfig
type CapacityConfig struct {
	// Required. The number of vCPUs to provision for the cluster. Minimum: 3.
	// +kcc:proto:field=google.cloud.managedkafka.v1.CapacityConfig.vcpu_count
	VcpuCount *int64 `json:"vcpuCount,omitempty"`

	// Required. The memory to provision for the cluster in bytes.
	//  The CPU:memory ratio (vCPU:GiB) must be between 1:1 and 1:8.
	//  Minimum: 3221225472 (3 GiB).
	// +kcc:proto:field=google.cloud.managedkafka.v1.CapacityConfig.memory_bytes
	MemoryBytes *int64 `json:"memoryBytes,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.Cluster
type Cluster struct {
	// Required. Configuration properties for a Kafka cluster deployed to Google
	//  Cloud Platform.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.gcp_config
	GcpConfig *GcpConfig `json:"gcpConfig,omitempty"`

	// Identifier. The name of the cluster. Structured like:
	//  projects/{project_number}/locations/{location}/clusters/{cluster_id}
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.name
	Name *string `json:"name,omitempty"`

	// Output only. The time when the cluster was created.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the cluster was last updated.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Labels as key value pairs.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Capacity configuration for the Kafka cluster.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.capacity_config
	CapacityConfig *CapacityConfig `json:"capacityConfig,omitempty"`

	// Optional. Rebalance configuration for the Kafka cluster.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.rebalance_config
	RebalanceConfig *RebalanceConfig `json:"rebalanceConfig,omitempty"`

	// Output only. The current state of the cluster.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.state
	State *string `json:"state,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.GcpConfig
type GcpConfig struct {
	// Required. Access configuration for the Kafka cluster.
	// +kcc:proto:field=google.cloud.managedkafka.v1.GcpConfig.access_config
	AccessConfig *AccessConfig `json:"accessConfig,omitempty"`

	// Optional. Immutable. The Cloud KMS Key name to use for encryption. The key
	//  must be located in the same region as the cluster and cannot be changed.
	//  Structured like:
	//  projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
	// +kcc:proto:field=google.cloud.managedkafka.v1.GcpConfig.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.NetworkConfig
type NetworkConfig struct {
	// Required. Name of the VPC subnet in which to create Private Service Connect
	//  (PSC) endpoints for the Kafka brokers and bootstrap address. Structured
	//  like: projects/{project}/regions/{region}/subnetworks/{subnet_id}
	//
	//  The subnet must be located in the same region as the Kafka cluster. The
	//  project may differ. Multiple subnets from the same parent network must not
	//  be specified.
	//
	//  The CIDR range of the subnet must be within the IPv4 address ranges for
	//  private networks, as specified in RFC 1918.
	// +kcc:proto:field=google.cloud.managedkafka.v1.NetworkConfig.subnet
	Subnet *string `json:"subnet,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.RebalanceConfig
type RebalanceConfig struct {
	// Optional. The rebalance behavior for the cluster.
	//  When not specified, defaults to `NO_REBALANCE`.
	// +kcc:proto:field=google.cloud.managedkafka.v1.RebalanceConfig.mode
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.AccessConfig
type AccessConfig struct {
	// Required. Virtual Private Cloud (VPC) networks that must be granted direct
	//  access to the Kafka cluster. Minimum of 1 network is required. Maximum 10
	//  networks can be specified.
	// +kcc:proto:field=google.cloud.managedkafka.v1.AccessConfig.network_configs
	NetworkConfigs []NetworkConfig `json:"networkConfigs,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.CapacityConfig
type CapacityConfig struct {
	// Required. The number of vCPUs to provision for the cluster. Minimum: 3.
	// +kcc:proto:field=google.cloud.managedkafka.v1.CapacityConfig.vcpu_count
	VcpuCount *int64 `json:"vcpuCount,omitempty"`

	// Required. The memory to provision for the cluster in bytes.
	//  The CPU:memory ratio (vCPU:GiB) must be between 1:1 and 1:8.
	//  Minimum: 3221225472 (3 GiB).
	// +kcc:proto:field=google.cloud.managedkafka.v1.CapacityConfig.memory_bytes
	MemoryBytes *int64 `json:"memoryBytes,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.GcpConfig
type GcpConfig struct {
	// Required. Access configuration for the Kafka cluster.
	// +kcc:proto:field=google.cloud.managedkafka.v1.GcpConfig.access_config
	AccessConfig *AccessConfig `json:"accessConfig,omitempty"`

	// Optional. Immutable. The Cloud KMS Key name to use for encryption. The key
	//  must be located in the same region as the cluster and cannot be changed.
	//  Structured like:
	//  projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
	// +kcc:proto:field=google.cloud.managedkafka.v1.GcpConfig.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.NetworkConfig
type NetworkConfig struct {
	// Required. Name of the VPC subnet in which to create Private Service Connect
	//  (PSC) endpoints for the Kafka brokers and bootstrap address. Structured
	//  like: projects/{project}/regions/{region}/subnetworks/{subnet_id}
	//
	//  The subnet must be located in the same region as the Kafka cluster. The
	//  project may differ. Multiple subnets from the same parent network must not
	//  be specified.
	//
	//  The CIDR range of the subnet must be within the IPv4 address ranges for
	//  private networks, as specified in RFC 1918.
	// +kcc:proto:field=google.cloud.managedkafka.v1.NetworkConfig.subnet
	Subnet *string `json:"subnet,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.RebalanceConfig
type RebalanceConfig struct {
	// Optional. The rebalance behavior for the cluster.
	//  When not specified, defaults to `NO_REBALANCE`.
	// +kcc:proto:field=google.cloud.managedkafka.v1.RebalanceConfig.mode
	Mode *string `json:"mode,omitempty"`
}
