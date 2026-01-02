// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

// +tool:krm-type-terraform
// proto.message: google.cloud.pubsublite.v1.Topic
// crd.kind: PubSubLiteTopic
// crd.version: v1alpha1
// terraform.src: github.com/hashicorp/terraform-provider-google-beta/google-beta/services/pubsublite/resource_pubsub_lite_topic.go

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PubSubLiteTopicGVK = GroupVersion.WithKind("PubSubLiteTopic")

// PubSubLiteTopicSpec defines the desired state of PubSubLiteTopic
// +kcc:spec:proto=google.cloud.pubsublite.v1.Topic
type PubSubLiteTopicSpec struct {
	// Required. Defines the parent path of the resource.
	// +required
	ProjectRef *parent.ProjectRef `json:"projectRef,omitempty"`

	// The zone of the pubsub lite topic.
	// +required
	Zone *string `json:"zone,omitempty"`

	// The region of the pubsub lite topic.
	Region *string `json:"region,omitempty"`

	// The PubSubLiteTopic name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// // The name of the topic.
	// //  Structured like:
	// //  projects/{project_number}/locations/{location}/topics/{topic_id}
	// // +kcc:proto:field=google.cloud.pubsublite.v1.Topic.name
	// Name *string `json:"name,omitempty"`

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

// Removed from autogen for field naming.
// +kcc:proto=google.cloud.pubsublite.v1.Topic.PartitionConfig.Capacity
type Topic_PartitionConfig_Capacity struct {
	// Publish throughput capacity per partition in MiB/s.
	//  Must be >= 4 and <= 16.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.PartitionConfig.Capacity.publish_mib_per_sec
	// +required
	PublishMIBPerSec *int32 `json:"publishMibPerSec,omitempty"`

	// Subscribe throughput capacity per partition in MiB/s.
	//  Must be >= 4 and <= 32.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.PartitionConfig.Capacity.subscribe_mib_per_sec
	// +required
	SubscribeMIBPerSec *int32 `json:"subscribeMibPerSec,omitempty"`
}

// Removed from autogen to remove depcreated scale field
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

	// // DEPRECATED: Use capacity instead which can express a superset of
	// //  configurations.
	// //
	// //  Every partition in the topic is allocated throughput equivalent to
	// //  `scale` times the standard partition throughput (4 MiB/s). This is also
	// //  reflected in the cost of this topic; a topic with `scale` of 2 and
	// //  count of 10 is charged for 20 partitions. This value must be in the
	// //  range [1,4].
	// // +kcc:proto:field=google.cloud.pubsublite.v1.Topic.PartitionConfig.scale
	// Scale *int32 `json:"scale,omitempty"`

	// The capacity configuration.
	// +kcc:proto:field=google.cloud.pubsublite.v1.Topic.PartitionConfig.capacity
	Capacity *Topic_PartitionConfig_Capacity `json:"capacity,omitempty"`
}

// PubSubLiteTopicStatus defines the config connector machine state of PubSubLiteTopic
type PubSubLiteTopicStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the PubSubLiteTopic resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// NOTYET(terraform)
	// // ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *PubSubLiteTopicObservedState `json:"observedState,omitempty"`
}

// PubSubLiteTopicObservedState is the state of the PubSubLiteTopic resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.pubsublite.v1.Topic
type PubSubLiteTopicObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcppubsublitetopic;gcppubsublitetopics
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// PubSubLiteTopic is the Schema for the PubSubLiteTopic API
// +k8s:openapi-gen=true
type PubSubLiteTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   PubSubLiteTopicSpec   `json:"spec,omitempty"`
	Status PubSubLiteTopicStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// PubSubLiteTopicList contains a list of PubSubLiteTopic
type PubSubLiteTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PubSubLiteTopic `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PubSubLiteTopic{}, &PubSubLiteTopicList{})
}
