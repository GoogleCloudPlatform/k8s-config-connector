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

// API sources for ManagedKafkaConnector, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/managedkafka/v1/resources.proto
// +kcc:source:service-docs=https://cloud.google.com/managed-service-for-apache-kafka/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/managed-service-for-apache-kafka/docs/reference/rest/v1/projects.locations.connectClusters.connectors

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ManagedKafkaConnectorGVK = GroupVersion.WithKind("ManagedKafkaConnector")

// ManagedKafkaConnectorSpec defines the desired state of ManagedKafkaConnector
// +kcc:spec:proto=google.cloud.managedkafka.v1.Connector
// +kcc:required-from-proto
type ManagedKafkaConnectorSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ManagedKafkaConnector name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Restarts the individual tasks of a Connector.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Connector.task_restart_policy
	TaskRestartPolicy *TaskRetryPolicy `json:"taskRestartPolicy,omitempty"`

	// Optional. Connector config as keys/values.
	//  The keys of the map are connector property names, for example:
	//  `connector.class`, `tasks.max`, `key.converter`.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Connector.configs
	Configs map[string]string `json:"configs,omitempty"`
}

// ManagedKafkaConnectorStatus defines the config connector machine state of ManagedKafkaConnector
type ManagedKafkaConnectorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ManagedKafkaConnector resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ManagedKafkaConnectorObservedState `json:"observedState,omitempty"`
}

// ManagedKafkaConnectorObservedState is the state of the ManagedKafkaConnector resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.managedkafka.v1.Connector
type ManagedKafkaConnectorObservedState struct {
	// Output only. The current state of the connector.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Connector.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmanagedkafkaconnector;gcpmanagedkafkaconnectors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ManagedKafkaConnector is the Schema for the ManagedKafkaConnector API
// +k8s:openapi-gen=true
type ManagedKafkaConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ManagedKafkaConnectorSpec   `json:"spec,omitempty"`
	Status ManagedKafkaConnectorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ManagedKafkaConnectorList contains a list of ManagedKafkaConnector
type ManagedKafkaConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedKafkaConnector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedKafkaConnector{}, &ManagedKafkaConnectorList{})
}
