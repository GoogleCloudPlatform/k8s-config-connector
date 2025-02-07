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


// +kcc:proto=google.cloud.apihub.v1.ApiOperation
type ApiOperation struct {
	// Identifier. The name of the operation.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}/operations/{operation}`
	// +kcc:proto:field=google.cloud.apihub.v1.ApiOperation.name
	Name *string `json:"name,omitempty"`

	// TODO: unsupported map type with key string and value message

}

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

// +kcc:proto=google.cloud.apihub.v1.Documentation
type Documentation struct {
	// Optional. The uri of the externally hosted documentation.
	// +kcc:proto:field=google.cloud.apihub.v1.Documentation.external_uri
	ExternalURI *string `json:"externalURI,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.HttpOperation
type HttpOperation struct {
}

// +kcc:proto=google.cloud.apihub.v1.OperationDetails
type OperationDetails struct {
	// The HTTP Operation.
	// +kcc:proto:field=google.cloud.apihub.v1.OperationDetails.http_operation
	HTTPOperation *HttpOperation `json:"httpOperation,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Path
type Path struct {
}

// +kcc:proto=google.cloud.apihub.v1.ApiOperation
type ApiOperationObservedState struct {
	// Output only. The name of the spec from where the operation was parsed.
	//  Format is
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}/specs/{spec}`
	// +kcc:proto:field=google.cloud.apihub.v1.ApiOperation.spec
	Spec *string `json:"spec,omitempty"`

	// Output only. Operation details.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiOperation.details
	Details *OperationDetails `json:"details,omitempty"`

	// Output only. The time at which the operation was created.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiOperation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the operation was last updated.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiOperation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.HttpOperation
type HttpOperationObservedState struct {
	// Output only. The path details for the Operation.
	// +kcc:proto:field=google.cloud.apihub.v1.HttpOperation.path
	Path *Path `json:"path,omitempty"`

	// Output only. Operation method
	// +kcc:proto:field=google.cloud.apihub.v1.HttpOperation.method
	Method *string `json:"method,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.OperationDetails
type OperationDetailsObservedState struct {
	// The HTTP Operation.
	// +kcc:proto:field=google.cloud.apihub.v1.OperationDetails.http_operation
	HTTPOperation *HttpOperationObservedState `json:"httpOperation,omitempty"`

	// Output only. Description of the operation behavior.
	//  For OpenAPI spec, this will map to `operation.description` in the
	//  spec, in case description is empty, `operation.summary` will be used.
	// +kcc:proto:field=google.cloud.apihub.v1.OperationDetails.description
	Description *string `json:"description,omitempty"`

	// Output only. Additional external documentation for this operation.
	//  For OpenAPI spec, this will map to `operation.documentation` in the spec.
	// +kcc:proto:field=google.cloud.apihub.v1.OperationDetails.documentation
	Documentation *Documentation `json:"documentation,omitempty"`

	// Output only. For OpenAPI spec, this will be set if `operation.deprecated`is
	//  marked as `true` in the spec.
	// +kcc:proto:field=google.cloud.apihub.v1.OperationDetails.deprecated
	Deprecated *bool `json:"deprecated,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Path
type PathObservedState struct {
	// Output only. Complete path relative to server endpoint.
	// +kcc:proto:field=google.cloud.apihub.v1.Path.path
	Path *string `json:"path,omitempty"`

	// Output only. A short description for the path applicable to all operations.
	// +kcc:proto:field=google.cloud.apihub.v1.Path.description
	Description *string `json:"description,omitempty"`
}
