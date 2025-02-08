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


// +kcc:proto=google.cloud.managedkafka.v1.ConsumerGroup
type ConsumerGroup struct {
	// Identifier. The name of the consumer group. The `consumer_group` segment is
	//  used when connecting directly to the cluster. Structured like:
	//  projects/{project}/locations/{location}/clusters/{cluster}/consumerGroups/{consumer_group}
	// +kcc:proto:field=google.cloud.managedkafka.v1.ConsumerGroup.name
	Name *string `json:"name,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.managedkafka.v1.ConsumerPartitionMetadata
type ConsumerPartitionMetadata struct {
	// Required. The current offset for this partition, or 0 if no offset has been
	//  committed.
	// +kcc:proto:field=google.cloud.managedkafka.v1.ConsumerPartitionMetadata.offset
	Offset *int64 `json:"offset,omitempty"`

	// Optional. The associated metadata for this partition, or empty if it does
	//  not exist.
	// +kcc:proto:field=google.cloud.managedkafka.v1.ConsumerPartitionMetadata.metadata
	Metadata *string `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.ConsumerTopicMetadata
type ConsumerTopicMetadata struct {

	// TODO: unsupported map type with key int32 and value message

}
