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

// API sources for DesignCenterConnection, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/547f8fdb4b91dbac70d540b3757520f0608ae356/google/cloud/designcenter/v1/authoring.proto
// +kcc:source:service-docs=https://cloud.google.com/application-design-center/docs
// +kcc:source:resource-docs=https://cloud.google.com/application-design-center/docs/reference/rest/v1/projects.locations.spaces.applicationTemplates.components.connections

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DesignCenterConnectionGVK = GroupVersion.WithKind("DesignCenterConnection")

// DesignCenterConnectionSpec defines the desired state of DesignCenterConnection
// +kcc:spec:proto=google.cloud.designcenter.v1.Connection
// +kcc:required-from-proto
type DesignCenterConnectionSpec struct {
	// The DesignCenterComponent this DesignCenterConnection is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *DesignCenterComponentRef `json:"parentRef"`

	// The DesignCenterConnection name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The destination component used to generate the connection.
	// +kcc:proto:field=google.cloud.designcenter.v1.Connection.destination_component_uri
	// +required
	DestinationComponentRef *DesignCenterComponentRef `json:"destinationComponentRef"`

	// Optional. The parameters of the connection associated with the source
	//  component.
	// +kcc:proto:field=google.cloud.designcenter.v1.Connection.source_component_parameters
	SourceComponentParameters []Parameter `json:"sourceComponentParameters,omitempty"`

	// Optional. The parameters of the connection associated with the destination
	//  component.
	// +kcc:proto:field=google.cloud.designcenter.v1.Connection.destination_component_parameters
	DestinationComponentParameters []Parameter `json:"destinationComponentParameters,omitempty"`
}

// DesignCenterConnectionStatus defines the config connector machine state of DesignCenterConnection
type DesignCenterConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DesignCenterConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DesignCenterConnectionObservedState `json:"observedState,omitempty"`
}

// DesignCenterConnectionObservedState is the state of the DesignCenterConnection resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.designcenter.v1.Connection
type DesignCenterConnectionObservedState struct {
	// Output only. The connection creation timestamp.
	// +kcc:proto:field=google.cloud.designcenter.v1.Connection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The connection update timestamp.
	// +kcc:proto:field=google.cloud.designcenter.v1.Connection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. The parameters of the connection associated with the source
	//  component.
	// +kcc:proto:field=google.cloud.designcenter.v1.Connection.source_component_parameters
	SourceComponentParameters []ParameterObservedState `json:"sourceComponentParameters,omitempty"`

	// Optional. The parameters of the connection associated with the destination
	//  component.
	// +kcc:proto:field=google.cloud.designcenter.v1.Connection.destination_component_parameters
	DestinationComponentParameters []ParameterObservedState `json:"destinationComponentParameters,omitempty"`

	// Output only. The type of the connection (data or logical).
	// +kcc:proto:field=google.cloud.designcenter.v1.Connection.dependency_type
	DependencyType *string `json:"dependencyType,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdesigncenterconnection;gcpdesigncenterconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DesignCenterConnection is the Schema for the DesignCenterConnection API
// +k8s:openapi-gen=true
type DesignCenterConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DesignCenterConnectionSpec   `json:"spec,omitempty"`
	Status DesignCenterConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DesignCenterConnectionList contains a list of DesignCenterConnection
type DesignCenterConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DesignCenterConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DesignCenterConnection{}, &DesignCenterConnectionList{})
}
