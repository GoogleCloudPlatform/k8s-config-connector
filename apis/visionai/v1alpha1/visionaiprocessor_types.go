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

// API sources for VisionAIProcessor, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/visionai/v1/platform.proto
// +kcc:source:service-docs=https://cloud.google.com/vision-ai/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/vision-ai/docs/reference/rest/v1/projects.locations.processors

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VisionAIProcessorGVK = GroupVersion.WithKind("VisionAIProcessor")

// VisionAIProcessorSpec defines the desired state of VisionAIProcessor
// +kcc:spec:proto=google.cloud.visionai.v1.Processor
// +kcc:required-from-proto
type VisionAIProcessorSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VisionAIProcessor name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. A user friendly display name for the processor.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Illustrative sentences for describing the functionality of the processor.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.description
	Description *string `json:"description,omitempty"`

	// Model Type.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.model_type
	ModelType *string `json:"modelType,omitempty"`

	// Source info for customer created processor.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.custom_processor_source_info
	CustomProcessorSourceInfo *CustomProcessorSourceInfo `json:"customProcessorSourceInfo,omitempty"`

	// Indicates if the processor supports post processing.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.supports_post_processing
	SupportsPostProcessing *bool `json:"supportsPostProcessing,omitempty"`

	// Which instance types this processor supports; if empty, this default to
	//  STREAMING_PREDICTION.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.supported_instance_types
	SupportedInstanceTypes []string `json:"supportedInstanceTypes,omitempty"`
}

// VisionAIProcessorStatus defines the config connector machine state of VisionAIProcessor
type VisionAIProcessorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VisionAIProcessor resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VisionAIProcessorObservedState `json:"observedState,omitempty"`
}

// VisionAIProcessorObservedState is the state of the VisionAIProcessor resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.visionai.v1.Processor
type VisionAIProcessorObservedState struct {
	// Output only. [Output only] Create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Processor Type.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.processor_type
	ProcessorType *string `json:"processorType,omitempty"`

	// Source info for customer created processor.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.custom_processor_source_info
	CustomProcessorSourceInfo *CustomProcessorSourceInfoObservedState `json:"customProcessorSourceInfo,omitempty"`

	// Output only. State of the Processor.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.state
	State *string `json:"state,omitempty"`

	// Output only. [Output only] The input / output specifications of a
	//  processor, each type of processor has fixed input / output specs which
	//  cannot be altered by customer.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.processor_io_spec
	ProcessorIoSpec *ProcessorIoSpec `json:"processorIoSpec,omitempty"`

	// Output only. The corresponding configuration can be used in the Application
	//  to customize the behavior of the processor.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.configuration_typeurl
	ConfigurationTypeurl *string `json:"configurationTypeurl,omitempty"`

	// +kcc:proto:field=google.cloud.visionai.v1.Processor.supported_annotation_types
	SupportedAnnotationTypes []string `json:"supportedAnnotationTypes,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvisionaiprocessor;gcpvisionaiprocessors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VisionAIProcessor is the Schema for the VisionAIProcessor API
// +k8s:openapi-gen=true
type VisionAIProcessor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VisionAIProcessorSpec   `json:"spec,omitempty"`
	Status VisionAIProcessorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VisionAIProcessorList contains a list of VisionAIProcessor
type VisionAIProcessorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VisionAIProcessor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VisionAIProcessor{}, &VisionAIProcessorList{})
}
