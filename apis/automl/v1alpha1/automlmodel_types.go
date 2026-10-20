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

// API sources for AutoMLModel, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/automl/v1/model.proto
// +kcc:guess=source-link reason=verify-service-docs-link
// +kcc:source:service-docs=
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AutoMLModelGVK = GroupVersion.WithKind("AutoMLModel")

// AutoMLModelSpec defines the desired state of AutoMLModel
// +kcc:spec:proto=google.cloud.automl.v1.Model
// +kcc:required-from-proto
type AutoMLModelSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AutoMLModel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Metadata for translation models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.translation_model_metadata
	TranslationModelMetadata *TranslationModelMetadata `json:"translationModelMetadata,omitempty"`

	// Metadata for image classification models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.image_classification_model_metadata
	ImageClassificationModelMetadata *ImageClassificationModelMetadata `json:"imageClassificationModelMetadata,omitempty"`

	// Metadata for text classification models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.text_classification_model_metadata
	TextClassificationModelMetadata *TextClassificationModelMetadata `json:"textClassificationModelMetadata,omitempty"`

	// Metadata for image object detection models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.image_object_detection_model_metadata
	ImageObjectDetectionModelMetadata *ImageObjectDetectionModelMetadata `json:"imageObjectDetectionModelMetadata,omitempty"`

	// Metadata for text extraction models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.text_extraction_model_metadata
	TextExtractionModelMetadata *TextExtractionModelMetadata `json:"textExtractionModelMetadata,omitempty"`

	// Metadata for text sentiment models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.text_sentiment_model_metadata
	TextSentimentModelMetadata *TextSentimentModelMetadata `json:"textSentimentModelMetadata,omitempty"`

	// Required. The name of the model to show in the interface. The name can be
	//  up to 32 characters long and can consist only of ASCII Latin letters A-Z
	//  and a-z, underscores
	//  (_), and ASCII digits 0-9. It must start with a letter.
	// +kcc:proto:field=google.cloud.automl.v1.Model.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The resource ID of the dataset used to create the model. The dataset must
	//  come from the same ancestor project and location.
	// +kcc:proto:field=google.cloud.automl.v1.Model.dataset_id
	DatasetID *string `json:"datasetID,omitempty"`

	// Output only. Deployment state of the model. A model can only serve
	//  prediction requests after it gets deployed.
	// +kcc:proto:field=google.cloud.automl.v1.Model.deployment_state
	DeploymentState *string `json:"deploymentState,omitempty"`

	// Optional. The labels with user-defined metadata to organize your model.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  Label values are optional. Label keys must start with a letter.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	// +kcc:proto:field=google.cloud.automl.v1.Model.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// AutoMLModelStatus defines the config connector machine state of AutoMLModel
type AutoMLModelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AutoMLModel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AutoMLModelObservedState `json:"observedState,omitempty"`
}

// AutoMLModelObservedState is the state of the AutoMLModel resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.automl.v1.Model
type AutoMLModelObservedState struct {
	// Metadata for image classification models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.image_classification_model_metadata
	ImageClassificationModelMetadata *ImageClassificationModelMetadataObservedState `json:"imageClassificationModelMetadata,omitempty"`

	// Metadata for image object detection models.
	// +kcc:proto:field=google.cloud.automl.v1.Model.image_object_detection_model_metadata
	ImageObjectDetectionModelMetadata *ImageObjectDetectionModelMetadataObservedState `json:"imageObjectDetectionModelMetadata,omitempty"`

	// Output only. Timestamp when the model training finished  and can be used for prediction.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.automl.v1.Model.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this model was last updated.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.automl.v1.Model.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Used to perform a consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.automl.v1.Model.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpautomlmodel;gcpautomlmodels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AutoMLModel is the Schema for the AutoMLModel API
// +k8s:openapi-gen=true
type AutoMLModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AutoMLModelSpec   `json:"spec,omitempty"`
	Status AutoMLModelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AutoMLModelList contains a list of AutoMLModel
type AutoMLModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoMLModel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AutoMLModel{}, &AutoMLModelList{})
}
