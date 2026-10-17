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

// API sources for DialogflowSessionEntityType, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dialogflow/cx/v3/session_entity_type.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/docs/reference/rest/v3/projects.locations.agents.sessions.entityTypes

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowSessionEntityTypeGVK = GroupVersion.WithKind("DialogflowSessionEntityType")

// DialogflowSessionEntityTypeSpec defines the desired state of DialogflowSessionEntityType
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.SessionEntityType
// +kcc:required-from-proto
type DialogflowSessionEntityTypeSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DialogflowSessionEntityType name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Indicates whether the additional data should override or
	//  supplement the custom entity type definition.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SessionEntityType.entity_override_mode
	// +required
	EntityOverrideMode *string `json:"entityOverrideMode,omitempty"`

	// Required. The collection of entities to override or supplement the custom
	//  entity type.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.SessionEntityType.entities
	// +required
	Entities []EntityType_Entity `json:"entities,omitempty"`
}

// DialogflowSessionEntityTypeStatus defines the config connector machine state of DialogflowSessionEntityType
type DialogflowSessionEntityTypeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowSessionEntityType resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *DialogflowSessionEntityTypeObservedState `json:"observedState,omitempty"`
}

// DialogflowSessionEntityTypeObservedState is the state of the DialogflowSessionEntityType resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.SessionEntityType
// type DialogflowSessionEntityTypeObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowsessionentitytype;gcpdialogflowsessionentitytypes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowSessionEntityType is the Schema for the DialogflowSessionEntityType API
// +k8s:openapi-gen=true
type DialogflowSessionEntityType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowSessionEntityTypeSpec   `json:"spec,omitempty"`
	Status DialogflowSessionEntityTypeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowSessionEntityTypeList contains a list of DialogflowSessionEntityType
type DialogflowSessionEntityTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowSessionEntityType `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowSessionEntityType{}, &DialogflowSessionEntityTypeList{})
}
