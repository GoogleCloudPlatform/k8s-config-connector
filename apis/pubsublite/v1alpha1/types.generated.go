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


// +kcc:proto=google.cloud.pubsublite.v1.Topic
type Topic struct {
	// The name of the topic.
	//  Structured like:
	//  projects/{project_number}/locations/{location}/topics/{topic_id}
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.name
	Name *string `json:"name,omitempty"`

	// The settings for this topic's partitions.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.partition_config
	PartitionConfig *Topic_PartitionConfig `json:"partitionConfig,omitempty"`

	// The settings for this topic's message retention.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.retention_config
	RetentionConfig *Topic_RetentionConfig `json:"retentionConfig,omitempty"`

	// The settings for this topic's Reservation usage.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.reservation_config
	ReservationConfig *Topic_ReservationConfig `json:"reservationConfig,omitempty"`
}

// +kcc:proto=google.cloud.pubsublite.v1.Topic.PartitionConfig
type Topic_PartitionConfig struct {
	// The number of partitions in the topic. Must be at least 1.
	//
	//  Once a topic has been created the number of partitions can be increased
	//  but not decreased. Message ordering is not guaranteed across a topic
	//  resize. For more information see
	//  https://cloud.google.com/pubsub/lite/docs/topics#scaling_capacity
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.PartitionConfig.count
	Count *int64 `json:"count,omitempty"`

	// DEPRECATED: Use capacity instead which can express a superset of
	//  configurations.
	//
	//  Every partition in the topic is allocated throughput equivalent to
	//  `scale` times the standard partition throughput (4 MiB/s). This is also
	//  reflected in the cost of this topic; a topic with `scale` of 2 and
	//  count of 10 is charged for 20 partitions. This value must be in the
	//  range [1,4].
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.PartitionConfig.scale
	Scale *int32 `json:"scale,omitempty"`

	// The capacity configuration.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.PartitionConfig.capacity
	Capacity *Topic_PartitionConfig_Capacity `json:"capacity,omitempty"`
}

// +kcc:proto=google.cloud.pubsublite.v1.Topic.PartitionConfig.Capacity
type Topic_PartitionConfig_Capacity struct {
	// Publish throughput capacity per partition in MiB/s.
	//  Must be >= 4 and <= 16.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.PartitionConfig.Capacity.publish_mib_per_sec
	PublishMibPerSec *int32 `json:"publishMibPerSec,omitempty"`

	// Subscribe throughput capacity per partition in MiB/s.
	//  Must be >= 4 and <= 32.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.PartitionConfig.Capacity.subscribe_mib_per_sec
	SubscribeMibPerSec *int32 `json:"subscribeMibPerSec,omitempty"`
}

// +kcc:proto=google.cloud.pubsublite.v1.Topic.ReservationConfig
type Topic_ReservationConfig struct {
	// The Reservation to use for this topic's throughput capacity.
	//  Structured like:
	//  projects/{project_number}/locations/{location}/reservations/{reservation_id}
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.ReservationConfig.throughput_reservation
	ThroughputReservation *string `json:"throughputReservation,omitempty"`
}

// +kcc:proto=google.cloud.pubsublite.v1.Topic.RetentionConfig
type Topic_RetentionConfig struct {
	// The provisioned storage, in bytes, per partition. If the number of bytes
	//  stored in any of the topic's partitions grows beyond this value, older
	//  messages will be dropped to make room for newer ones, regardless of the
	//  value of `period`.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.RetentionConfig.per_partition_bytes
	PerPartitionBytes *int64 `json:"perPartitionBytes,omitempty"`

	// How long a published message is retained. If unset, messages will be
	//  retained as long as the bytes retained for each partition is below
	//  `per_partition_bytes`.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.RetentionConfig.period
	Period *string `json:"period,omitempty"`
}
