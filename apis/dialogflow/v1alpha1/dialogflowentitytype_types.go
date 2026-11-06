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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowEntityTypeGVK = GroupVersion.WithKind("DialogflowEntityType")

// +kcc:proto=google.cloud.dialogflow.v2.EntityType.Entity
type EntityType_Entity struct {
	// A collection of value synonyms. For example, if the entity type is vegetable, and value is scallions, a synonym
	// could be green onions.
	// For KIND_LIST entity types:
	// * This collection must contain exactly one synonym equal to value.
	// +required
	Synonyms []string `json:"synonyms"`

	// The primary value associated with this entity entry. For example, if the entity type is vegetable, the value
	// could be scallions.
	// For KIND_MAP entity types:
	// * A reference value to be used in place of synonyms.
	// For KIND_LIST entity types:
	// * A string that can contain references to other entity types (with or without aliases).
	// +required
	Value string `json:"value"`
}

// +kcc:spec:proto=google.cloud.dialogflow.v2.EntityType
type DialogflowEntityTypeSpec struct {
	// The name of this entity type to be displayed on the console.
	// +required
	DisplayName string `json:"displayName"`

	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction *bool `json:"enableFuzzyExtraction,omitempty"`

	// The collection of entity entries associated with the entity type.
	Entities []EntityType_Entity `json:"entities,omitempty"`

	// Indicates the kind of entity type.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
	// types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values. Possible values: ["KIND_MAP", "KIND_LIST", "KIND_REGEXP"].
	// +required
	Kind string `json:"kind"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`
}

type DialogflowEntityTypeStatus struct {
	// Conditions represent the latest available observation of the resource's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// The unique identifier of the entity type.
	// Format: projects/<Project ID>/agent/entityTypes/<Entity type ID>.
	Name *string `json:"name,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowEntityType resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowentitytype;gcpdialogflowentitytypes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowEntityType is the Schema for the DialogflowEntityType API
// +k8s:openapi-gen=true
type DialogflowEntityType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowEntityTypeSpec   `json:"spec"`
	Status DialogflowEntityTypeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowEntityTypeList contains a list of DialogflowEntityType
type DialogflowEntityTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowEntityType `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowEntityType{}, &DialogflowEntityTypeList{})
}
