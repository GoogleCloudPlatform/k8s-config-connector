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


// +kcc:proto=google.cloud.managedkafka.v1.RebalanceConfig
type RebalanceConfig struct {
	// Optional. The rebalance behavior for the cluster.
	//  When not specified, defaults to `NO_REBALANCE`.
	// +kcc:proto:field=google.cloud.managedkafka.v1.RebalanceConfig.mode
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.Topic
type Topic struct {
	// Identifier. The name of the topic. The `topic` segment is used when
	//  connecting directly to the cluster. Structured like:
	//  projects/{project}/locations/{location}/clusters/{cluster}/topics/{topic}
	// +kcc:proto:field=google.cloud.managedkafka.v1.Topic.name
	Name *string `json:"name,omitempty"`

	// Required. The number of partitions this topic has. The partition count can
	//  only be increased, not decreased. Please note that if partitions are
	//  increased for a topic that has a key, the partitioning logic or the
	//  ordering of the messages will be affected.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Topic.partition_count
	PartitionCount *int32 `json:"partitionCount,omitempty"`

	// Required. Immutable. The number of replicas of each partition. A
	//  replication factor of 3 is recommended for high availability.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Topic.replication_factor
	ReplicationFactor *int32 `json:"replicationFactor,omitempty"`

	// Optional. Configurations for the topic that are overridden from the cluster
	//  defaults. The key of the map is a Kafka topic property name, for example:
	//  `cleanup.policy`, `compression.type`.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Topic.configs
	Configs map[string]string `json:"configs,omitempty"`
}
