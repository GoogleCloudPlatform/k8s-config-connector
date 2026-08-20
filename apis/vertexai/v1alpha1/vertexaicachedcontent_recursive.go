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

// +kcc:proto=google.cloud.aiplatform.v1beta1.Schema
type Schema struct {
	// Optional. The type of the data.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.type
	Type *string `json:"type,omitempty"`

	// Optional. The format of the data.
	//  Supported formats:
	//   for NUMBER type: "float", "double"
	//   for INTEGER type: "int32", "int64"
	//   for STRING type: "email", "byte", etc
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.format
	Format *string `json:"format,omitempty"`

	// Optional. The title of the Schema.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.title
	Title *string `json:"title,omitempty"`

	// Optional. The description of the data.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.description
	Description *string `json:"description,omitempty"`

	// Optional. Indicates if the value may be null.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.nullable
	Nullable *bool `json:"nullable,omitempty"`

	// Optional. Default value of the data.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.default
	Default *Value `json:"default,omitempty"`

	// Optional. SCHEMA FIELDS FOR TYPE ARRAY
	//  Schema of the elements of Type.ARRAY.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.items
	// Items is represented as JSON due to CRD recursion constraints
	Items *apiextensionsv1.JSON `json:"items,omitempty"`

	// Optional. Minimum number of the elements for Type.ARRAY.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.min_items
	MinItems *int64 `json:"minItems,omitempty"`

	// Optional. Maximum number of the elements for Type.ARRAY.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.max_items
	MaxItems *int64 `json:"maxItems,omitempty"`

	// Optional. Possible values of the element of Type.STRING with enum format.
	//  For example we can define an Enum Direction as :
	//  {type:STRING, format:enum, enum:["EAST", NORTH", "SOUTH", "WEST"]}
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.enum
	Enum []string `json:"enum,omitempty"`

	// TODO: unsupported map type with key string and value message

	// Optional. The order of the properties.
	//  Not a standard field in open api spec. Only used to support the order of
	//  the properties.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.property_ordering
	PropertyOrdering []string `json:"propertyOrdering,omitempty"`

	// Optional. Required properties of Type.OBJECT.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.required
	Required []string `json:"required,omitempty"`

	// Optional. Minimum number of the properties for Type.OBJECT.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.min_properties
	MinProperties *int64 `json:"minProperties,omitempty"`

	// Optional. Maximum number of the properties for Type.OBJECT.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.max_properties
	MaxProperties *int64 `json:"maxProperties,omitempty"`

	// Optional. SCHEMA FIELDS FOR TYPE INTEGER and NUMBER
	//  Minimum value of the Type.INTEGER and Type.NUMBER
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.minimum
	Minimum *float64 `json:"minimum,omitempty"`

	// Optional. Maximum value of the Type.INTEGER and Type.NUMBER
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.maximum
	Maximum *float64 `json:"maximum,omitempty"`

	// Optional. SCHEMA FIELDS FOR TYPE STRING
	//  Minimum length of the Type.STRING
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.min_length
	MinLength *int64 `json:"minLength,omitempty"`

	// Optional. Maximum length of the Type.STRING
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.max_length
	MaxLength *int64 `json:"maxLength,omitempty"`

	// Optional. Pattern of the Type.STRING to restrict a string to a regular
	//  expression.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.pattern
	Pattern *string `json:"pattern,omitempty"`

	// Optional. Example of the object. Will only populated when the object is the
	//  root.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.example
	Example *Value `json:"example,omitempty"`

	// Optional. The value should be validated against any (one or more) of the
	//  subschemas in the list.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.any_of
	// AnyOf is represented as JSON due to CRD recursion constraints
	AnyOf []apiextensionsv1.JSON `json:"anyOf,omitempty"`

	// Optional. Can either be a boolean or an object; controls the presence of
	//  additional properties.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schema.additional_properties
	AdditionalProperties *Value `json:"additionalProperties,omitempty"`

	// Optional. Allows indirect references between schema nodes. The value should
	//  be a valid reference to a child of the root `defs`.
	Ref *string `json:"ref,omitempty"`

	// TODO: unsupported map type with key string and value message
}
