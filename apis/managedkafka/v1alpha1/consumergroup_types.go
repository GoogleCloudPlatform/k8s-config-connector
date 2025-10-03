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

import (
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ManagedKafkaConsumerGroupGVK = GroupVersion.WithKind("ManagedKafkaConsumerGroup")

// ManagedKafkaConsumerGroupSpec defines the desired state of ManagedKafkaConsumerGroup
// +kcc:spec:proto=google.cloud.managedkafka.v1.ConsumerGroup
type ManagedKafkaConsumerGroupSpec struct {
	*Parent `json:",inline"`

	// The ManagedKafkaConsumerGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.ConsumerPartitionMetadata
type ConsumerPartitionMetadata struct {
	// Required. The current offset for this partition, or 0 if no offset has been
	//  committed.
	//+required
	// +kcc:proto:field=google.cloud.managedkafka.v1.ConsumerPartitionMetadata.offset
	Offset *int64 `json:"offset,omitempty"`

	// Optional. The associated metadata for this partition, or empty if it does
	//  not exist.
	// +kcc:proto:field=google.cloud.managedkafka.v1.ConsumerPartitionMetadata.metadata
	Metadata *string `json:"metadata,omitempty"`

	// Required. Key of the partition index for topic metadata in this consumer group.
	//+required
	Key *int32 `json:"key,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.ConsumerTopicMetadata
type ConsumerTopicMetadata struct {
	// Optional. Metadata for this consumer group and topic for all partition
	// indexes it has metadata for.
	Partitions []*ConsumerPartitionMetadata `json:"partitions,omitempty"`
}

type Parent struct {
	// +required
	Location string `json:"location"`
	// +required
	ClusterRef *ClusterRef `json:"clusterRef"`
	// +optional
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef,omitempty"`
}

// ManagedKafkaConsumerGroupStatus defines the config connector machine state of ManagedKafkaConsumerGroup
type ManagedKafkaConsumerGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ManagedKafkaConsumerGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ManagedKafkaConsumerGroupObservedState `json:"observedState,omitempty"`
}

// ManagedKafkaConsumerGroupObservedState is the state of the ManagedKafkaConsumerGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.managedkafka.v1.ConsumerGroup
type ManagedKafkaConsumerGroupObservedState struct {
	// Optional. Metadata for this consumer group for all topics it has metadata for.
	// The key of the map is a topic name, structured like: projects/{project}/locations/{location}/clusters/{cluster}/topics/{topic}
	Topics map[string]*ConsumerTopicMetadata `json:"topics,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpmanagedkafkaconsumergroup;gcpmanagedkafkaconsumergroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ManagedKafkaConsumerGroup is the Schema for the ManagedKafkaConsumerGroup API
// +k8s:openapi-gen=true
type ManagedKafkaConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ManagedKafkaConsumerGroupSpec   `json:"spec,omitempty"`
	Status ManagedKafkaConsumerGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ManagedKafkaConsumerGroupList contains a list of ManagedKafkaConsumerGroup
type ManagedKafkaConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedKafkaConsumerGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedKafkaConsumerGroup{}, &ManagedKafkaConsumerGroupList{})
}
