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

package v1beta1

import (
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ManagedKafkaTopicGVK = GroupVersion.WithKind("ManagedKafkaTopic")

// ManagedKafkaTopicSpec defines the desired state of ManagedKafkaTopic
// +kcc:proto=google.cloud.managedkafka.v1.Topic
type ManagedKafkaTopicSpec struct {
	commonv1alpha1.CommonSpec `json:",inline"`

	// Required. the location of the Kafka resource.
	// See https://cloud.google.com/managed-kafka/docs/locations for a list of supported locations.
	// +required
	Location string `json:"location"`

	// Required. Reference to the Kafka cluster to create the topic in.
	// +required
	ClusterRef *ClusterRef `json:"clusterRef"`

	// The ManagedKafkaTopic name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

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

// ManagedKafkaTopicStatus defines the config connector machine state of ManagedKafkaTopic
type ManagedKafkaTopicStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ManagedKafkaTopic resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// NOTYET: the resource does not have any output only fields
	// ObservedState *ManagedKafkaTopicObservedState `json:"observedState,omitempty"`
}

// ManagedKafkaTopicObservedState is the state of the ManagedKafkaTopic resource as most recently observed in GCP.
// +kcc:proto=google.cloud.managedkafka.v1.Topic
// NOTYET: the resource does not have any output only fields
// type ManagedKafkaTopicObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmanagedkafkatopic;gcpmanagedkafkatopics
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:storageversion

// ManagedKafkaTopic is the Schema for the ManagedKafkaTopic API
// +k8s:openapi-gen=true
type ManagedKafkaTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ManagedKafkaTopicSpec   `json:"spec,omitempty"`
	Status ManagedKafkaTopicStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ManagedKafkaTopicList contains a list of ManagedKafkaTopic
type ManagedKafkaTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedKafkaTopic `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedKafkaTopic{}, &ManagedKafkaTopicList{})
}
