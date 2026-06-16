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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AutoMLDatasetGVK = GroupVersion.WithKind("AutoMLDataset")

// AutoMLDatasetSpec defines the desired state of AutoMLDataset
// +kcc:spec:proto=google.cloud.automl.v1.Dataset
type AutoMLDatasetSpec struct {
	// The project that this resource belongs to.
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location string `json:"location"`

	// The AutoMLDataset name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Metadata for a dataset used for translation.
	// +kubebuilder:validation:Optional
	TranslationDatasetMetadata *TranslationDatasetMetadata `json:"translationDatasetMetadata,omitempty"`

	// Metadata for a dataset used for image classification.
	// +kubebuilder:validation:Optional
	ImageClassificationDatasetMetadata *ImageClassificationDatasetMetadata `json:"imageClassificationDatasetMetadata,omitempty"`

	// Metadata for a dataset used for text classification.
	// +kubebuilder:validation:Optional
	TextClassificationDatasetMetadata *TextClassificationDatasetMetadata `json:"textClassificationDatasetMetadata,omitempty"`

	// Metadata for a dataset used for image object detection.
	// +kubebuilder:validation:Optional
	ImageObjectDetectionDatasetMetadata *ImageObjectDetectionDatasetMetadata `json:"imageObjectDetectionDatasetMetadata,omitempty"`

	// Metadata for a dataset used for text extraction.
	// +kubebuilder:validation:Optional
	TextExtractionDatasetMetadata *TextExtractionDatasetMetadata `json:"textExtractionDatasetMetadata,omitempty"`

	// Metadata for a dataset used for text sentiment.
	// +kubebuilder:validation:Optional
	TextSentimentDatasetMetadata *TextSentimentDatasetMetadata `json:"textSentimentDatasetMetadata,omitempty"`

	// Required. The name of the dataset to show in the interface. The name can be
	//  up to 32 characters long and can consist only of ASCII Latin letters A-Z
	//  and a-z, underscores
	//  (_), and ASCII digits 0-9.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName"`

	// User-provided description of the dataset. The description can be up to
	//  25000 characters long.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Optional. The labels with user-defined metadata to organize your dataset.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  Label values are optional. Label keys must start with a letter.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`
}

// AutoMLDatasetStatus defines the config connector machine state of AutoMLDataset
type AutoMLDatasetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AutoMLDataset resource in Google Cloud.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in Google Cloud.
	ObservedState *AutoMLDatasetObservedState `json:"observedState,omitempty"`
}

// AutoMLDatasetObservedState is the state of the AutoMLDataset resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.cloud.automl.v1.Dataset
type AutoMLDatasetObservedState struct {
	// Output only. The number of examples in the dataset.
	ExampleCount *int32 `json:"exampleCount,omitempty"`

	// Output only. Timestamp when this dataset was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Used to perform consistent read-modify-write updates.
	Etag *string `json:"etag,omitempty"`
}

// Metadata for a dataset used for image classification.
type ImageClassificationDatasetMetadata struct {
	// Required. Type of the classification problem.
	// +kubebuilder:validation:Required
	ClassificationType *string `json:"classificationType"`
}

// Metadata for a dataset used for image object detection.
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:Schemaless
type ImageObjectDetectionDatasetMetadata struct {
}

// Metadata for a dataset used for text classification.
type TextClassificationDatasetMetadata struct {
	// Required. Type of the classification problem.
	// +kubebuilder:validation:Required
	ClassificationType *string `json:"classificationType"`
}

// Metadata for a dataset used for text extraction.
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:Schemaless
type TextExtractionDatasetMetadata struct {
}

// Metadata for a dataset used for text sentiment.
type TextSentimentDatasetMetadata struct {
	// Required. A sentiment is expressed as an integer ordinal, where higher value
	//  means a more positive sentiment. The range of sentiments that will be used
	//  is between 0 and sentiment_max (inclusive on both ends), and all the values
	//  in the range must be represented in the dataset before a model can be
	//  created.
	//  sentiment_max value must be between 1 and 10 (inclusive).
	// +kubebuilder:validation:Required
	SentimentMax *int32 `json:"sentimentMax"`
}

// Metadata for a dataset used for translation.
type TranslationDatasetMetadata struct {
	// Required. The BCP-47 language code of the source language.
	// +kubebuilder:validation:Required
	SourceLanguageCode *string `json:"sourceLanguageCode"`

	// Required. The BCP-47 language code of the target language.
	// +kubebuilder:validation:Required
	TargetLanguageCode *string `json:"targetLanguageCode"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpautomldataset;gcpautomldatasets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AutoMLDataset is the Schema for the AutoMLDataset API
// +k8s:openapi-gen=true
type AutoMLDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AutoMLDatasetSpec   `json:"spec,omitempty"`
	Status AutoMLDatasetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AutoMLDatasetList contains a list of AutoMLDataset
type AutoMLDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoMLDataset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AutoMLDataset{}, &AutoMLDatasetList{})
}
