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


// +kcc:proto=google.cloud.visionai.v1.Analysis
type Analysis struct {
	// The name of resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Analysis.name
	Name *string `json:"name,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.visionai.v1.Analysis.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The definition of the analysis.
	// +kcc:proto:field=google.cloud.visionai.v1.Analysis.analysis_definition
	AnalysisDefinition *AnalysisDefinition `json:"analysisDefinition,omitempty"`

	// Map from the input parameter in the definition to the real stream.
	//  E.g., suppose you had a stream source operator named "input-0" and you try
	//  to receive from the real stream "stream-0". You can add the following
	//  mapping: [input-0: stream-0].
	// +kcc:proto:field=google.cloud.visionai.v1.Analysis.input_streams_mapping
	InputStreamsMapping map[string]string `json:"inputStreamsMapping,omitempty"`

	// Map from the output parameter in the definition to the real stream.
	//  E.g., suppose you had a stream sink operator named "output-0" and you try
	//  to send to the real stream "stream-0". You can add the following
	//  mapping: [output-0: stream-0].
	// +kcc:proto:field=google.cloud.visionai.v1.Analysis.output_streams_mapping
	OutputStreamsMapping map[string]string `json:"outputStreamsMapping,omitempty"`

	// Boolean flag to indicate whether you would like to disable the ability
	//  to automatically start a Process when new event happening in the input
	//  Stream. If you would like to start a Process manually, the field needs
	//  to be set to true.
	// +kcc:proto:field=google.cloud.visionai.v1.Analysis.disable_event_watch
	DisableEventWatch *bool `json:"disableEventWatch,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.AnalysisDefinition
type AnalysisDefinition struct {
	// Analyzer definitions.
	// +kcc:proto:field=google.cloud.visionai.v1.AnalysisDefinition.analyzers
	Analyzers []AnalyzerDefinition `json:"analyzers,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.AnalyzerDefinition
type AnalyzerDefinition struct {
	// The name of this analyzer.
	//
	//  Tentatively [a-z][a-z0-9]*(_[a-z0-9]+)*.
	// +kcc:proto:field=google.cloud.visionai.v1.AnalyzerDefinition.analyzer
	Analyzer *string `json:"analyzer,omitempty"`

	// The name of the operator that this analyzer runs.
	//
	//  Must match the name of a supported operator.
	// +kcc:proto:field=google.cloud.visionai.v1.AnalyzerDefinition.operator
	Operator *string `json:"operator,omitempty"`

	// Input streams.
	// +kcc:proto:field=google.cloud.visionai.v1.AnalyzerDefinition.inputs
	Inputs []AnalyzerDefinition_StreamInput `json:"inputs,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Debug options.
	// +kcc:proto:field=google.cloud.visionai.v1.AnalyzerDefinition.debug_options
	DebugOptions *AnalyzerDefinition_DebugOptions `json:"debugOptions,omitempty"`

	// Operator option.
	// +kcc:proto:field=google.cloud.visionai.v1.AnalyzerDefinition.operator_option
	OperatorOption *AnalyzerDefinition_OperatorOption `json:"operatorOption,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.AnalyzerDefinition.DebugOptions
type AnalyzerDefinition_DebugOptions struct {
	// Environment variables.
	// +kcc:proto:field=google.cloud.visionai.v1.AnalyzerDefinition.DebugOptions.environment_variables
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.AnalyzerDefinition.OperatorOption
type AnalyzerDefinition_OperatorOption struct {
	// Tag of the operator.
	// +kcc:proto:field=google.cloud.visionai.v1.AnalyzerDefinition.OperatorOption.tag
	Tag *string `json:"tag,omitempty"`

	// Registry of the operator. e.g. public, dev.
	// +kcc:proto:field=google.cloud.visionai.v1.AnalyzerDefinition.OperatorOption.registry
	Registry *string `json:"registry,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.AnalyzerDefinition.StreamInput
type AnalyzerDefinition_StreamInput struct {
	// The name of the stream input (as discussed above).
	// +kcc:proto:field=google.cloud.visionai.v1.AnalyzerDefinition.StreamInput.input
	Input *string `json:"input,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.AttributeValue
type AttributeValue struct {
	// int.
	// +kcc:proto:field=google.cloud.visionai.v1.AttributeValue.i
	I *int64 `json:"i,omitempty"`

	// float.
	// +kcc:proto:field=google.cloud.visionai.v1.AttributeValue.f
	F *float32 `json:"f,omitempty"`

	// bool.
	// +kcc:proto:field=google.cloud.visionai.v1.AttributeValue.b
	B *bool `json:"b,omitempty"`

	// string.
	// +kcc:proto:field=google.cloud.visionai.v1.AttributeValue.s
	S []byte `json:"s,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Analysis
type AnalysisObservedState struct {
	// Output only. The create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Analysis.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Analysis.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
