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

// API sources for TranslateGlossary, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/translate/v3/translation_service.proto
// +kcc:source:service-docs=https://cloud.google.com/translate/docs/quickstarts
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/translate/docs/quickstarts/reference/rest/v3/projects.locations.glossaries

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TranslateGlossaryGVK = GroupVersion.WithKind("TranslateGlossary")

// TranslateGlossarySpec defines the desired state of TranslateGlossary
// +kcc:spec:proto=google.cloud.translation.v3.Glossary
// +kcc:required-from-proto
type TranslateGlossarySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The TranslateGlossary name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Used with unidirectional glossaries.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.language_pair
	LanguagePair *Glossary_LanguageCodePair `json:"languagePair,omitempty"`

	// Used with equivalent term set glossaries.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.language_codes_set
	LanguageCodesSet *Glossary_LanguageCodesSet `json:"languageCodesSet,omitempty"`

	// Required. Provides examples to build the glossary from.
	//  Total glossary must not exceed 10M Unicode codepoints.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.input_config
	InputConfig *GlossaryInputConfig `json:"inputConfig,omitempty"`

	// Optional. The display name of the glossary.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// TranslateGlossaryStatus defines the config connector machine state of TranslateGlossary
type TranslateGlossaryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TranslateGlossary resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TranslateGlossaryObservedState `json:"observedState,omitempty"`
}

// TranslateGlossaryObservedState is the state of the TranslateGlossary resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.translation.v3.Glossary
type TranslateGlossaryObservedState struct {
	// Output only. The number of entries defined in the glossary.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.entry_count
	EntryCount *int32 `json:"entryCount,omitempty"`

	// Output only. When CreateGlossary was called.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.submit_time
	SubmitTime *string `json:"submitTime,omitempty"`

	// Output only. When the glossary creation was finished.
	// +kcc:proto:field=google.cloud.translation.v3.Glossary.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptranslateglossary;gcptranslateglossaries
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TranslateGlossary is the Schema for the TranslateGlossary API
// +k8s:openapi-gen=true
type TranslateGlossary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TranslateGlossarySpec   `json:"spec,omitempty"`
	Status TranslateGlossaryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TranslateGlossaryList contains a list of TranslateGlossary
type TranslateGlossaryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TranslateGlossary `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TranslateGlossary{}, &TranslateGlossaryList{})
}
