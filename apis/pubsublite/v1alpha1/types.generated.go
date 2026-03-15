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

// +generated:types
// krm.group: pubsublite.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.pubsublite.v1
// resource: PubSubLiteSubscription:Subscription
// resource: PubSubLiteTopic:Topic

package v1alpha1

// +kcc:proto=google.cloud.pubsublite.v1.ExportConfig
type ExportConfig struct {
	// The desired state of this export. Setting this to values other than
	//  `ACTIVE` and `PAUSED` will result in an error.
	// +kcc:proto:field=google.cloud.pubsublite.v1.ExportConfig.desired_state
	DesiredState *string `json:"desiredState,omitempty"`

	// Optional. The name of an optional Pub/Sub Lite topic to publish messages
	//  that can not be exported to the destination. For example, the message can
	//  not be published to the Pub/Sub service because it does not satisfy the
	//  constraints documented at https://cloud.google.com/pubsub/docs/publisher.
	//
	//  Structured like:
	//  projects/{project_number}/locations/{location}/topics/{topic_id}.
	//  Must be within the same project and location as the subscription. The topic
	//  may be changed or removed.
	// +kcc:proto:field=google.cloud.pubsublite.v1.ExportConfig.dead_letter_topic
	DeadLetterTopic *string `json:"deadLetterTopic,omitempty"`

	// Messages are automatically written from the Pub/Sub Lite topic associated
	//  with this subscription to a Pub/Sub topic.
	// +kcc:proto:field=google.cloud.pubsublite.v1.ExportConfig.pubsub_config
	PubsubConfig *ExportConfig_PubSubConfig `json:"pubsubConfig,omitempty"`
}

// +kcc:proto=google.cloud.pubsublite.v1.ExportConfig.PubSubConfig
type ExportConfig_PubSubConfig struct {
	// The name of the Pub/Sub topic.
	//  Structured like: projects/{project_number}/topics/{topic_id}.
	//  The topic may be changed.
	// +kcc:proto:field=google.cloud.pubsublite.v1.ExportConfig.PubSubConfig.topic
	Topic *string `json:"topic,omitempty"`
}

// +kcc:proto=google.cloud.pubsublite.v1.Subscription.DeliveryConfig
type Subscription_DeliveryConfig struct {
	// The DeliveryRequirement for this subscription.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Subscription.DeliveryConfig.delivery_requirement
	DeliveryRequirement *string `json:"deliveryRequirement,omitempty"`
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

// +kcc:observedstate:proto=google.cloud.pubsublite.v1.ExportConfig
type ExportConfigObservedState struct {
	// Output only. The current state of the export, which may be different to the
	//  desired state due to errors. This field is output only.
	// +kcc:proto:field=google.cloud.pubsublite.v1.ExportConfig.current_state
	CurrentState *string `json:"currentState,omitempty"`
}
