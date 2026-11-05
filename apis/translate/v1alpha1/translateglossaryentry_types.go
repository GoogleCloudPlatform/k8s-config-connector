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

// API sources for TranslateGlossaryEntry, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/translate/v3/common.proto
// +kcc:source:service-docs=https://cloud.google.com/translate/docs/quickstarts
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/translate/docs/quickstarts/reference/rest/v3/projects.locations.glossaries.glossaryEntries

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TranslateGlossaryEntryGVK = GroupVersion.WithKind("TranslateGlossaryEntry")

// TranslateGlossaryEntrySpec defines the desired state of TranslateGlossaryEntry
// +kcc:spec:proto=google.cloud.translation.v3.GlossaryEntry
// +kcc:required-from-proto
type TranslateGlossaryEntrySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The TranslateGlossaryEntry name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Used for an unidirectional glossary.
	// +kcc:proto:field=google.cloud.translation.v3.GlossaryEntry.terms_pair
	TermsPair *GlossaryEntry_GlossaryTermsPair `json:"termsPair,omitempty"`

	// Used for an equivalent term sets glossary.
	// +kcc:proto:field=google.cloud.translation.v3.GlossaryEntry.terms_set
	TermsSet *GlossaryEntry_GlossaryTermsSet `json:"termsSet,omitempty"`

	// Describes the glossary entry.
	// +kcc:proto:field=google.cloud.translation.v3.GlossaryEntry.description
	Description *string `json:"description,omitempty"`
}

// TranslateGlossaryEntryStatus defines the config connector machine state of TranslateGlossaryEntry
type TranslateGlossaryEntryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TranslateGlossaryEntry resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *TranslateGlossaryEntryObservedState `json:"observedState,omitempty"`
}

// TranslateGlossaryEntryObservedState is the state of the TranslateGlossaryEntry resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.translation.v3.GlossaryEntry
// type TranslateGlossaryEntryObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptranslateglossaryentry;gcptranslateglossaryentries
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TranslateGlossaryEntry is the Schema for the TranslateGlossaryEntry API
// +k8s:openapi-gen=true
type TranslateGlossaryEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TranslateGlossaryEntrySpec   `json:"spec,omitempty"`
	Status TranslateGlossaryEntryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TranslateGlossaryEntryList contains a list of TranslateGlossaryEntry
type TranslateGlossaryEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TranslateGlossaryEntry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TranslateGlossaryEntry{}, &TranslateGlossaryEntryList{})
}
