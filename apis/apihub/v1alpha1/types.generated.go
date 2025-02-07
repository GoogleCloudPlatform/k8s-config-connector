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


// +kcc:proto=google.cloud.apihub.v1.Attribute.AllowedValue
type Attribute_AllowedValue struct {
	// Required. The ID of the allowed value.
	//  * If provided, the same will be used. The service will throw an error if
	//  the specified id is already used by another allowed value in the same
	//  attribute resource.
	//  * If not provided, a system generated id derived from the display name
	//  will be used. In this case, the service will handle conflict resolution
	//  by adding a system generated suffix in case of duplicates.
	//
	//  This value should be 4-63 characters, and valid characters
	//  are /[a-z][0-9]-/.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.AllowedValue.id
	ID *string `json:"id,omitempty"`

	// Required. The display name of the allowed value.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.AllowedValue.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The detailed description of the allowed value.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.AllowedValue.description
	Description *string `json:"description,omitempty"`

	// Optional. When set to true, the allowed value cannot be updated or
	//  deleted by the user. It can only be true for System defined attributes.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.AllowedValue.immutable
	Immutable *bool `json:"immutable,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.AttributeValues
type AttributeValues struct {
	// The attribute values associated with a resource in case attribute data
	//  type is enum.
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.enum_values
	EnumValues *AttributeValues_EnumAttributeValues `json:"enumValues,omitempty"`

	// The attribute values associated with a resource in case attribute data
	//  type is string.
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.string_values
	StringValues *AttributeValues_StringAttributeValues `json:"stringValues,omitempty"`

	// The attribute values associated with a resource in case attribute data
	//  type is JSON.
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.json_values
	JsonValues *AttributeValues_StringAttributeValues `json:"jsonValues,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.AttributeValues.EnumAttributeValues
type AttributeValues_EnumAttributeValues struct {
	// Required. The attribute values in case attribute data type is enum.
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.EnumAttributeValues.values
	Values []Attribute_AllowedValue `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.AttributeValues.StringAttributeValues
type AttributeValues_StringAttributeValues struct {
	// Required. The attribute values in case attribute data type is string or
	//  JSON.
	// +kcc:proto:field=google.cloud.apihub.v1.AttributeValues.StringAttributeValues.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Definition
type Definition struct {

	// Identifier. The name of the definition.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}/definitions/{definition}`
	// +kcc:proto:field=google.cloud.apihub.v1.Definition.name
	Name *string `json:"name,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.apihub.v1.Schema
type Schema struct {
}

// +kcc:proto=google.cloud.apihub.v1.Definition
type DefinitionObservedState struct {
	// Output only. The value of a schema definition.
	// +kcc:proto:field=google.cloud.apihub.v1.Definition.schema
	Schema *Schema `json:"schema,omitempty"`

	// Output only. The name of the spec from where the definition was parsed.
	//  Format is
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}/specs/{spec}`
	// +kcc:proto:field=google.cloud.apihub.v1.Definition.spec
	Spec *string `json:"spec,omitempty"`

	// Output only. The type of the definition.
	// +kcc:proto:field=google.cloud.apihub.v1.Definition.type
	Type *string `json:"type,omitempty"`

	// Output only. The time at which the definition was created.
	// +kcc:proto:field=google.cloud.apihub.v1.Definition.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the definition was last updated.
	// +kcc:proto:field=google.cloud.apihub.v1.Definition.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Schema
type SchemaObservedState struct {
	// Output only. The display name of the schema.
	//  This will map to the name of the schema in the spec.
	// +kcc:proto:field=google.cloud.apihub.v1.Schema.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. The raw value of the schema definition corresponding to the
	//  schema name in the spec.
	// +kcc:proto:field=google.cloud.apihub.v1.Schema.raw_value
	RawValue []byte `json:"rawValue,omitempty"`
}
