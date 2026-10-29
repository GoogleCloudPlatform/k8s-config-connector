// Copyright 2026 Google LLC
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

// API sources for ManagedKafkaSchemaRegistry, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/managedkafka/schemaregistry/v1/schema_registry_resources.proto
// +kcc:source:service-docs=https://cloud.google.com/managed-service-for-apache-kafka/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/managed-service-for-apache-kafka/docs/reference/rest/v1/projects.locations.schemaRegistries

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ManagedKafkaSchemaRegistryGVK = GroupVersion.WithKind("ManagedKafkaSchemaRegistry")

// ManagedKafkaSchemaRegistrySpec defines the desired state of ManagedKafkaSchemaRegistry
// +kcc:spec:proto=google.cloud.managedkafka.schemaregistry.v1.SchemaRegistry
// +kcc:required-from-proto
type ManagedKafkaSchemaRegistrySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ManagedKafkaSchemaRegistry name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ManagedKafkaSchemaRegistryStatus defines the config connector machine state of ManagedKafkaSchemaRegistry
type ManagedKafkaSchemaRegistryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ManagedKafkaSchemaRegistry resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ManagedKafkaSchemaRegistryObservedState `json:"observedState,omitempty"`
}

// ManagedKafkaSchemaRegistryObservedState is the state of the ManagedKafkaSchemaRegistry resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.managedkafka.schemaregistry.v1.SchemaRegistry
type ManagedKafkaSchemaRegistryObservedState struct {
	// Output only. The contexts of the schema registry instance.
	// +kcc:proto:field=google.cloud.managedkafka.schemaregistry.v1.SchemaRegistry.contexts
	Contexts []string `json:"contexts,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmanagedkafkaschemaregistry;gcpmanagedkafkaschemaregistries
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ManagedKafkaSchemaRegistry is the Schema for the ManagedKafkaSchemaRegistry API
// +k8s:openapi-gen=true
type ManagedKafkaSchemaRegistry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ManagedKafkaSchemaRegistrySpec   `json:"spec,omitempty"`
	Status ManagedKafkaSchemaRegistryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ManagedKafkaSchemaRegistryList contains a list of ManagedKafkaSchemaRegistry
type ManagedKafkaSchemaRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedKafkaSchemaRegistry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedKafkaSchemaRegistry{}, &ManagedKafkaSchemaRegistryList{})
}
