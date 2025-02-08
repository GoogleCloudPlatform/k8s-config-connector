// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.visionai.v1.CustomProcessorSourceInfo
type CustomProcessorSourceInfo struct {
	// The resource name original model hosted in the vertex AI platform.
	// +kcc:proto:field=google.cloud.visionai.v1.CustomProcessorSourceInfo.vertex_model
	VertexModel *string `json:"vertexModel,omitempty"`

	// Artifact for product recognizer.
	// +kcc:proto:field=google.cloud.visionai.v1.CustomProcessorSourceInfo.product_recognizer_artifact
	ProductRecognizerArtifact *CustomProcessorSourceInfo_ProductRecognizerArtifact `json:"productRecognizerArtifact,omitempty"`

	// The original product which holds the custom processor's functionality.
	// +kcc:proto:field=google.cloud.visionai.v1.CustomProcessorSourceInfo.source_type
	SourceType *string `json:"sourceType,omitempty"`

	// Model schema files which specifies the signature of the model.
	//  For VERTEX_CUSTOM models, instances schema is required.
	//  If instances schema is not specified during the processor creation,
	//  VisionAI Platform will try to get it from Vertex, if it doesn't exist, the
	//  creation will fail.
	// +kcc:proto:field=google.cloud.visionai.v1.CustomProcessorSourceInfo.model_schema
	ModelSchema *CustomProcessorSourceInfo_ModelSchema `json:"modelSchema,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.CustomProcessorSourceInfo.ModelSchema
type CustomProcessorSourceInfo_ModelSchema struct {
	// Cloud Storage location to a YAML file that defines the format of a single
	//  instance used in prediction and explanation requests.
	// +kcc:proto:field=google.cloud.visionai.v1.CustomProcessorSourceInfo.ModelSchema.instances_schema
	InstancesSchema *GcsSource `json:"instancesSchema,omitempty"`

	// Cloud Storage location to a YAML file that defines the prediction and
	//  explanation parameters.
	// +kcc:proto:field=google.cloud.visionai.v1.CustomProcessorSourceInfo.ModelSchema.parameters_schema
	ParametersSchema *GcsSource `json:"parametersSchema,omitempty"`

	// Cloud Storage location to a YAML file that defines the format of a single
	//  prediction or explanation.
	// +kcc:proto:field=google.cloud.visionai.v1.CustomProcessorSourceInfo.ModelSchema.predictions_schema
	PredictionsSchema *GcsSource `json:"predictionsSchema,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.CustomProcessorSourceInfo.ProductRecognizerArtifact
type CustomProcessorSourceInfo_ProductRecognizerArtifact struct {
	// Required. Resource name of RetailProductRecognitionIndex.
	//  Format is
	//  'projects/*/locations/*/retailCatalogs/*/retailProductRecognitionIndexes/*'
	// +kcc:proto:field=google.cloud.visionai.v1.CustomProcessorSourceInfo.ProductRecognizerArtifact.retail_product_recognition_index
	RetailProductRecognitionIndex *string `json:"retailProductRecognitionIndex,omitempty"`

	// Optional. The resource name of embedding model hosted in Vertex AI
	//  Platform.
	// +kcc:proto:field=google.cloud.visionai.v1.CustomProcessorSourceInfo.ProductRecognizerArtifact.vertex_model
	VertexModel *string `json:"vertexModel,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.GcsSource
type GcsSource struct {
	// Required. References to a Google Cloud Storage paths.
	// +kcc:proto:field=google.cloud.visionai.v1.GcsSource.uris
	Uris []string `json:"uris,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Processor
type Processor struct {
	// name of resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.name
	Name *string `json:"name,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. A user friendly display name for the processor.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.display_name
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

// +kcc:proto=google.cloud.visionai.v1.ProcessorIOSpec
type ProcessorIOSpec struct {
	// For processors with input_channel_specs, the processor must be explicitly
	//  connected to another processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.graph_input_channel_specs
	GraphInputChannelSpecs []ProcessorIOSpec_GraphInputChannelSpec `json:"graphInputChannelSpecs,omitempty"`

	// The output artifact specifications for the current processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.graph_output_channel_specs
	GraphOutputChannelSpecs []ProcessorIOSpec_GraphOutputChannelSpec `json:"graphOutputChannelSpecs,omitempty"`

	// The input resource that needs to be fed from the application instance.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.instance_resource_input_binding_specs
	InstanceResourceInputBindingSpecs []ProcessorIOSpec_InstanceResourceInputBindingSpec `json:"instanceResourceInputBindingSpecs,omitempty"`

	// The output resource that the processor will generate per instance.
	//  Other than the explicitly listed output bindings here, all the processors'
	//  GraphOutputChannels can be binded to stream resource. The bind name then is
	//  the same as the GraphOutputChannel's name.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.instance_resource_output_binding_specs
	InstanceResourceOutputBindingSpecs []ProcessorIOSpec_InstanceResourceOutputBindingSpec `json:"instanceResourceOutputBindingSpecs,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.ProcessorIOSpec.GraphInputChannelSpec
type ProcessorIOSpec_GraphInputChannelSpec struct {
	// The name of the current input channel.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.GraphInputChannelSpec.name
	Name *string `json:"name,omitempty"`

	// The data types of the current input channel.
	//  When this field has more than 1 value, it means this input channel can be
	//  connected to either of these different data types.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.GraphInputChannelSpec.data_type
	DataType *string `json:"dataType,omitempty"`

	// If specified, only those detailed data types can be connected to the
	//  processor. For example, jpeg stream for MEDIA, or PredictionResult proto
	//  for PROTO type. If unspecified, then any proto is accepted.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.GraphInputChannelSpec.accepted_data_type_uris
	AcceptedDataTypeUris []string `json:"acceptedDataTypeUris,omitempty"`

	// Whether the current input channel is required by the processor.
	//  For example, for a processor with required video input and optional audio
	//  input, if video input is missing, the application will be rejected while
	//  the audio input can be missing as long as the video input exists.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.GraphInputChannelSpec.required
	Required *bool `json:"required,omitempty"`

	// How many input edges can be connected to this input channel. 0 means
	//  unlimited.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.GraphInputChannelSpec.max_connection_allowed
	MaxConnectionAllowed *int64 `json:"maxConnectionAllowed,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.ProcessorIOSpec.GraphOutputChannelSpec
type ProcessorIOSpec_GraphOutputChannelSpec struct {
	// The name of the current output channel.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.GraphOutputChannelSpec.name
	Name *string `json:"name,omitempty"`

	// The data type of the current output channel.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.GraphOutputChannelSpec.data_type
	DataType *string `json:"dataType,omitempty"`

	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.GraphOutputChannelSpec.data_type_uri
	DataTypeURI *string `json:"dataTypeURI,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.ProcessorIOSpec.InstanceResourceInputBindingSpec
type ProcessorIOSpec_InstanceResourceInputBindingSpec struct {
	// The configuration proto that includes the Googleapis resources. I.e.
	//  type.googleapis.com/google.cloud.vision.v1.StreamWithAnnotation
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.InstanceResourceInputBindingSpec.config_type_uri
	ConfigTypeURI *string `json:"configTypeURI,omitempty"`

	// The direct type url of Googleapis resource. i.e.
	//  type.googleapis.com/google.cloud.vision.v1.Asset
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.InstanceResourceInputBindingSpec.resource_type_uri
	ResourceTypeURI *string `json:"resourceTypeURI,omitempty"`

	// Name of the input binding, unique within the processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.InstanceResourceInputBindingSpec.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.ProcessorIOSpec.InstanceResourceOutputBindingSpec
type ProcessorIOSpec_InstanceResourceOutputBindingSpec struct {
	// Name of the output binding, unique within the processor.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.InstanceResourceOutputBindingSpec.name
	Name *string `json:"name,omitempty"`

	// The resource type uri of the acceptable output resource.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.InstanceResourceOutputBindingSpec.resource_type_uri
	ResourceTypeURI *string `json:"resourceTypeURI,omitempty"`

	// Whether the output resource needs to be explicitly set in the instance.
	//  If it is false, the processor will automatically generate it if required.
	// +kcc:proto:field=google.cloud.visionai.v1.ProcessorIOSpec.InstanceResourceOutputBindingSpec.explicit
	Explicit *bool `json:"explicit,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.CustomProcessorSourceInfo
type CustomProcessorSourceInfoObservedState struct {
	// Output only. Additional info related to the imported custom processor.
	//  Data is filled in by app platform during the processor creation.
	// +kcc:proto:field=google.cloud.visionai.v1.CustomProcessorSourceInfo.additional_info
	AdditionalInfo map[string]string `json:"additionalInfo,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Processor
type ProcessorObservedState struct {
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
	ProcessorIoSpec *ProcessorIOSpec `json:"processorIoSpec,omitempty"`

	// Output only. The corresponding configuration can be used in the Application
	//  to customize the behavior of the processor.
	// +kcc:proto:field=google.cloud.visionai.v1.Processor.configuration_typeurl
	ConfigurationTypeurl *string `json:"configurationTypeurl,omitempty"`

	// +kcc:proto:field=google.cloud.visionai.v1.Processor.supported_annotation_types
	SupportedAnnotationTypes []string `json:"supportedAnnotationTypes,omitempty"`
}
