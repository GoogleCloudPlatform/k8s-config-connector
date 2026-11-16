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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowCXEntityTypeGVK = GroupVersion.WithKind("DialogflowCXEntityType")

// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.EntityType
type DialogflowCXEntityTypeSpec struct {
	// Represents kinds of entities.
	// * AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.
	// * AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity. Possible values: ["AUTO_EXPANSION_MODE_DEFAULT", "AUTO_EXPANSION_MODE_UNSPECIFIED"].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EntityType.auto_expansion_mode
	// +optional
	AutoExpansionMode *string `json:"autoExpansionMode,omitempty"`

	// The human-readable name of the entity type, unique within the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EntityType.display_name
	// +required
	DisplayName *string `json:"displayName"`

	// Enables fuzzy entity extraction during classification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EntityType.enable_fuzzy_extraction
	// +optional
	EnableFuzzyExtraction *bool `json:"enableFuzzyExtraction,omitempty"`

	// The collection of entity entries associated with the entity type.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EntityType.entities
	// +required
	Entities []EntityType_Entity `json:"entities"`

	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
	// If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EntityType.excluded_phrases
	// +optional
	ExcludedPhrases []EntityType_ExcludedPhrase `json:"excludedPhrases,omitempty"`

	// Indicates whether the entity type can be automatically expanded.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values. Possible values: ["KIND_MAP", "KIND_LIST", "KIND_REGEXP"].
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EntityType.kind
	// +required
	Kind *string `json:"kind"`

	// Immutable. The language of the following fields in entityType:
	// EntityType.entities.value
	// EntityType.entities.synonyms
	// EntityType.excluded_phrases.value
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	// +optional
	LanguageCode *string `json:"languageCode,omitempty"`

	// Immutable. The agent to create a entity type for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	// +optional
	Parent *string `json:"parent,omitempty"`

	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EntityType.redact
	// +optional
	Redact *bool `json:"redact,omitempty"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.EntityType.Entity
type EntityType_Entity struct {
	// A collection of value synonyms. For example, if the entity type is vegetable, and value is scallions, a synonym could be green onions.
	// For KIND_LIST entity types: This collection must contain exactly one synonym equal to value.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EntityType.Entity.synonyms
	// +optional
	Synonyms []string `json:"synonyms,omitempty"`

	// The primary value associated with this entity entry. For example, if the entity type is vegetable, the value could be scallions.
	// For KIND_MAP entity types: A canonical value to be used in place of synonyms.
	// For KIND_LIST entity types: A string that can contain references to other entity types (with or without aliases).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EntityType.Entity.value
	// +optional
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.EntityType.ExcludedPhrase
type EntityType_ExcludedPhrase struct {
	// The word or phrase to be excluded.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EntityType.ExcludedPhrase.value
	// +optional
	Value *string `json:"value,omitempty"`
}

type DialogflowCXEntityTypeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// The unique identifier of the entity type.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/entityTypes/<Entity Type ID>.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.EntityType.name
	// +optional
	Name *string `json:"name,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowCXEntityType resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`
}

// DialogflowCXEntityTypeObservedState is the state of the DialogflowCXEntityType resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.EntityType
type DialogflowCXEntityTypeObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowcxentitytype;gcpdialogflowcxentitytypes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowCXEntityType is the Schema for the DialogflowCXEntityType API
// +k8s:openapi-gen=true
type DialogflowCXEntityType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowCXEntityTypeSpec   `json:"spec,omitempty"`
	Status DialogflowCXEntityTypeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowCXEntityTypeList contains a list of DialogflowCXEntityType
type DialogflowCXEntityTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowCXEntityType `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowCXEntityType{}, &DialogflowCXEntityTypeList{})
}
