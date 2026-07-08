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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	// +kubebuilder:validation:items:Type=object
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue apiextensionsv1.JSON `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	// ListValue is temporarily disabled due to CRD instability
	// ListValue *ListValue `json:"listValue,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExplanationParameters
type ExplanationParameters struct {
	// An attribution method that approximates Shapley values for features that
	//  contribute to the label being predicted. A sampling strategy is used to
	//  approximate the value rather than considering all subsets of features.
	//  Refer to this paper for model details: https://arxiv.org/abs/1306.4265.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.sampled_shapley_attribution
	SampledShapleyAttribution *SampledShapleyAttribution `json:"sampledShapleyAttribution,omitempty"`

	// An attribution method that computes Aumann-Shapley values taking
	//  advantage of the model's fully differentiable structure. Refer to this
	//  paper for more details: https://arxiv.org/abs/1703.01365
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.integrated_gradients_attribution
	IntegratedGradientsAttribution *IntegratedGradientsAttribution `json:"integratedGradientsAttribution,omitempty"`

	// An attribution method that redistributes Integrated Gradients
	//  attribution to segmented regions, taking advantage of the model's fully
	//  differentiable structure. Refer to this paper for
	//  more details: https://arxiv.org/abs/1906.02825
	//
	//  XRAI currently performs better on natural images, like a picture of a
	//  house or an animal. If the images are taken in artificial environments,
	//  like a lab or manufacturing line, or from diagnostic equipment, like
	//  x-rays or quality-control cameras, use Integrated Gradients instead.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.xrai_attribution
	XraiAttribution *XraiAttribution `json:"xraiAttribution,omitempty"`

	// Example-based explanations that returns the nearest neighbors from the
	//  provided dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.examples
	Examples *Examples `json:"examples,omitempty"`

	// If populated, returns attributions for top K indices of outputs
	//  (defaults to 1). Only applies to Models that predicts more than one outputs
	//  (e,g, multi-class Models). When set to -1, returns explanations for all
	//  outputs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.top_k
	TopK *int32 `json:"topK,omitempty"`

	// If populated, only returns attributions that have
	//  [output_index][google.cloud.aiplatform.v1.Attribution.output_index]
	//  contained in output_indices. It must be an ndarray of integers, with the
	//  same shape of the output it's explaining.
	//
	//  If not populated, returns attributions for
	//  [top_k][google.cloud.aiplatform.v1.ExplanationParameters.top_k] indices of
	//  outputs. If neither top_k nor output_indices is populated, returns the
	//  argmax index of the outputs.
	//
	//  Only applicable to Models that predict multiple outputs (e,g, multi-class
	//  Models that predict multiple classes).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ExplanationParameters.output_indices
	// OutputIndices is temporarily disabled due to CRD instability
	// OutputIndices *ListValue `json:"outputIndices,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Schema
type Schema struct {
	// Optional. The type of the data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.type
	Type *string `json:"type,omitempty"`

	// Optional. The format of the data.
	//  Supported formats:
	//   for NUMBER type: "float", "double"
	//   for INTEGER type: "int32", "int64"
	//   for STRING type: "email", "byte", etc
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.format
	Format *string `json:"format,omitempty"`

	// Optional. The title of the Schema.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.title
	Title *string `json:"title,omitempty"`

	// Optional. The description of the data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.description
	Description *string `json:"description,omitempty"`

	// Optional. Indicates if the value may be null.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.nullable
	Nullable *bool `json:"nullable,omitempty"`

	// Optional. Default value of the data.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.default
	Default *Value `json:"default,omitempty"`

	// Optional. SCHEMA FIELDS FOR TYPE ARRAY
	//  Schema of the elements of Type.ARRAY.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.items
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	Items *Schema `json:"items,omitempty"`

	// Optional. Minimum number of the elements for Type.ARRAY.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.min_items
	MinItems *int64 `json:"minItems,omitempty"`

	// Optional. Maximum number of the elements for Type.ARRAY.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.max_items
	MaxItems *int64 `json:"maxItems,omitempty"`

	// Optional. Possible values of the element of primitive type with enum
	//  format. Examples:
	//  1. We can define direction as :
	//  {type:STRING, format:enum, enum:["EAST", NORTH", "SOUTH", "WEST"]}
	//  2. We can define apartment number as :
	//  {type:INTEGER, format:enum, enum:["101", "201", "301"]}
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.enum
	Enum []string `json:"enum,omitempty"`

	// TODO: unsupported map type with key string and value message

	// Optional. The order of the properties.
	//  Not a standard field in open api spec. Only used to support the order of
	//  the properties.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.property_ordering
	PropertyOrdering []string `json:"propertyOrdering,omitempty"`

	// Optional. Required properties of Type.OBJECT.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.required
	Required []string `json:"required,omitempty"`

	// Optional. Minimum number of the properties for Type.OBJECT.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.min_properties
	MinProperties *int64 `json:"minProperties,omitempty"`

	// Optional. Maximum number of the properties for Type.OBJECT.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.max_properties
	MaxProperties *int64 `json:"maxProperties,omitempty"`

	// Optional. SCHEMA FIELDS FOR TYPE INTEGER and NUMBER
	//  Minimum value of the Type.INTEGER and Type.NUMBER
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.minimum
	Minimum *float64 `json:"minimum,omitempty"`

	// Optional. Maximum value of the Type.INTEGER and Type.NUMBER
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.maximum
	Maximum *float64 `json:"maximum,omitempty"`

	// Optional. SCHEMA FIELDS FOR TYPE STRING
	//  Minimum length of the Type.STRING
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.min_length
	MinLength *int64 `json:"minLength,omitempty"`

	// Optional. Maximum length of the Type.STRING
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.max_length
	MaxLength *int64 `json:"maxLength,omitempty"`

	// Optional. Pattern of the Type.STRING to restrict a string to a regular
	//  expression.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.pattern
	Pattern *string `json:"pattern,omitempty"`

	// Optional. Example of the object. Will only populated when the object is the
	//  root.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.example
	Example *Value `json:"example,omitempty"`

	// Optional. The value should be validated against any (one or more) of the
	//  subschemas in the list.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.any_of
	// +kubebuilder:validation:items:XPreserveUnknownFields
	// +kubebuilder:validation:items:Type=object
	AnyOf []Schema `json:"anyOf,omitempty"`

	// Optional. Can either be a boolean or an object; controls the presence of
	//  additional properties.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.additional_properties
	AdditionalProperties *Value `json:"additionalProperties,omitempty"`

	// Optional. Allows indirect references between schema nodes. The value should
	//  be a valid reference to a child of the root `defs`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schema.ref
	Ref *string `json:"ref,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.ExactMatchSpec
// +kubebuilder:validation:XPreserveUnknownFields
type ExactMatchSpec struct {
}
