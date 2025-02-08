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

// +kcc:proto=google.cloud.visionai.v1.Operator
type Operator struct {
	// Name of the resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Operator.name
	Name *string `json:"name,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.visionai.v1.Operator.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The definition of the operator.
	// +kcc:proto:field=google.cloud.visionai.v1.Operator.operator_definition
	OperatorDefinition *OperatorDefinition `json:"operatorDefinition,omitempty"`

	// The link to the docker image of the operator.
	// +kcc:proto:field=google.cloud.visionai.v1.Operator.docker_image
	DockerImage *string `json:"dockerImage,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.OperatorDefinition
type OperatorDefinition struct {
	// The name of this operator.
	//
	//  Tentatively [A-Z][a-zA-Z0-9]*, e.g., BboxCounter, PetDetector,
	//  PetDetector1.
	// +kcc:proto:field=google.cloud.visionai.v1.OperatorDefinition.operator
	Operator *string `json:"operator,omitempty"`

	// Declares input arguments.
	// +kcc:proto:field=google.cloud.visionai.v1.OperatorDefinition.input_args
	InputArgs []OperatorDefinition_ArgumentDefinition `json:"inputArgs,omitempty"`

	// Declares output arguments.
	// +kcc:proto:field=google.cloud.visionai.v1.OperatorDefinition.output_args
	OutputArgs []OperatorDefinition_ArgumentDefinition `json:"outputArgs,omitempty"`

	// Declares the attributes.
	// +kcc:proto:field=google.cloud.visionai.v1.OperatorDefinition.attributes
	Attributes []OperatorDefinition_AttributeDefinition `json:"attributes,omitempty"`

	// The resources for running the operator.
	// +kcc:proto:field=google.cloud.visionai.v1.OperatorDefinition.resources
	Resources *ResourceSpecification `json:"resources,omitempty"`

	// Short description of the operator.
	// +kcc:proto:field=google.cloud.visionai.v1.OperatorDefinition.short_description
	ShortDescription *string `json:"shortDescription,omitempty"`

	// Full description of the operator.
	// +kcc:proto:field=google.cloud.visionai.v1.OperatorDefinition.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.OperatorDefinition.ArgumentDefinition
type OperatorDefinition_ArgumentDefinition struct {
	// The name of the argument.
	//
	//  Tentatively [a-z]([_a-z0-9]*[a-z0-9])?, e.g., video, audio,
	//  high_fps_frame.
	// +kcc:proto:field=google.cloud.visionai.v1.OperatorDefinition.ArgumentDefinition.argument
	Argument *string `json:"argument,omitempty"`

	// The data type of the argument.
	//
	//  This should match the textual representation of a stream/Packet type.
	// +kcc:proto:field=google.cloud.visionai.v1.OperatorDefinition.ArgumentDefinition.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.OperatorDefinition.AttributeDefinition
type OperatorDefinition_AttributeDefinition struct {
	// The name of the attribute.
	//
	//  Tentatively [a-z]([_a-z0-9]*[a-z0-9])?, e.g., max_frames_per_video,
	//  resize_height.
	// +kcc:proto:field=google.cloud.visionai.v1.OperatorDefinition.AttributeDefinition.attribute
	Attribute *string `json:"attribute,omitempty"`

	// The type of this attribute.
	//
	//  See attribute_value.proto for possibilities.
	// +kcc:proto:field=google.cloud.visionai.v1.OperatorDefinition.AttributeDefinition.type
	Type *string `json:"type,omitempty"`

	// The default value for the attribute.
	// +kcc:proto:field=google.cloud.visionai.v1.OperatorDefinition.AttributeDefinition.default_value
	DefaultValue *AttributeValue `json:"defaultValue,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.ResourceSpecification
type ResourceSpecification struct {
	// CPU specification.
	//
	//  Examples:
	//  "100m", "0.5", "1", "2", ... correspond to
	//  0.1, half, 1, or 2 cpus.
	//
	//  Leave empty to let the system decide.
	//
	//  Note that this does *not* determine the cpu vender/make,
	//  or its underlying clock speed and specific SIMD features.
	//  It is only the amount time it requires in timeslicing.
	// +kcc:proto:field=google.cloud.visionai.v1.ResourceSpecification.cpu
	Cpu *string `json:"cpu,omitempty"`

	// CPU limit.
	//
	//  Examples:
	//  "100m", "0.5", "1", "2", ... correspond to
	//  0.1, half, 1, or 2 cpus.
	//
	//  Leave empty to indicate no limit.
	// +kcc:proto:field=google.cloud.visionai.v1.ResourceSpecification.cpu_limits
	CpuLimits *string `json:"cpuLimits,omitempty"`

	// Memory specification (in bytes).
	//
	//  Examples:
	//  "128974848", "129e6", "129M", "123Mi", ... correspond to
	//  128974848 bytes, 129000000 bytes, 129 mebibytes, 123 megabytes.
	//
	//  Leave empty to let the system decide.
	// +kcc:proto:field=google.cloud.visionai.v1.ResourceSpecification.memory
	Memory *string `json:"memory,omitempty"`

	// Memory usage limits.
	//
	//  Examples:
	//  "128974848", "129e6", "129M", "123Mi", ... correspond to
	//  128974848 bytes, 129000000 bytes, 129 mebibytes, 123 megabytes.
	//
	//  Leave empty to indicate no limit.
	// +kcc:proto:field=google.cloud.visionai.v1.ResourceSpecification.memory_limits
	MemoryLimits *string `json:"memoryLimits,omitempty"`

	// Number of gpus.
	// +kcc:proto:field=google.cloud.visionai.v1.ResourceSpecification.gpus
	Gpus *int32 `json:"gpus,omitempty"`

	// The maximum latency that this operator may use to process an element.
	//
	//  If non positive, then a system default will be used.
	//  Operator developers should arrange for the system compute resources to be
	//  aligned with this latency budget; e.g. if you want a ML model to produce
	//  results within 500ms, then you should make sure you request enough
	//  cpu/gpu/memory to achieve that.
	// +kcc:proto:field=google.cloud.visionai.v1.ResourceSpecification.latency_budget_ms
	LatencyBudgetMs *int32 `json:"latencyBudgetMs,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.Operator
type OperatorObservedState struct {
	// Output only. The create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Operator.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Operator.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
