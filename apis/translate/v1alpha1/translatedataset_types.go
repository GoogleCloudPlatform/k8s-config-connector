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

// API sources for TranslateDataset, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/translate/v3/automl_translation.proto
// +kcc:source:service-docs=https://cloud.google.com/translate/docs/quickstarts
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/translate/docs/quickstarts/reference/rest/v3/projects.locations.datasets

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TranslateDatasetGVK = GroupVersion.WithKind("TranslateDataset")

// TranslateDatasetSpec defines the desired state of TranslateDataset
// +kcc:spec:proto=google.cloud.translation.v3.Dataset
// +kcc:required-from-proto
type TranslateDatasetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The TranslateDataset name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The name of the dataset to show in the interface. The name can be
	//  up to 32 characters long and can consist only of ASCII Latin letters A-Z
	//  and a-z, underscores (_), and ASCII digits 0-9.
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The BCP-47 language code of the source language.
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.source_language_code
	SourceLanguageCode *string `json:"sourceLanguageCode,omitempty"`

	// The BCP-47 language code of the target language.
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.target_language_code
	TargetLanguageCode *string `json:"targetLanguageCode,omitempty"`
}

// TranslateDatasetStatus defines the config connector machine state of TranslateDataset
type TranslateDatasetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TranslateDataset resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TranslateDatasetObservedState `json:"observedState,omitempty"`
}

// TranslateDatasetObservedState is the state of the TranslateDataset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.translation.v3.Dataset
type TranslateDatasetObservedState struct {
	// Output only. The number of examples in the dataset.
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.example_count
	ExampleCount *int32 `json:"exampleCount,omitempty"`

	// Output only. Number of training examples (sentence pairs).
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.train_example_count
	TrainExampleCount *int32 `json:"trainExampleCount,omitempty"`

	// Output only. Number of validation examples (sentence pairs).
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.validate_example_count
	ValidateExampleCount *int32 `json:"validateExampleCount,omitempty"`

	// Output only. Number of test examples (sentence pairs).
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.test_example_count
	TestExampleCount *int32 `json:"testExampleCount,omitempty"`

	// Output only. Timestamp when this dataset was created.
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this dataset was last updated.
	// +kcc:proto:field=google.cloud.translation.v3.Dataset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptranslatedataset;gcptranslatedatasets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TranslateDataset is the Schema for the TranslateDataset API
// +k8s:openapi-gen=true
type TranslateDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TranslateDatasetSpec   `json:"spec,omitempty"`
	Status TranslateDatasetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TranslateDatasetList contains a list of TranslateDataset
type TranslateDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TranslateDataset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TranslateDataset{}, &TranslateDatasetList{})
}
