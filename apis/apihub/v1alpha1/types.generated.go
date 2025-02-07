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

// +kcc:proto=google.cloud.apihub.v1.Dependency
type Dependency struct {
	// Identifier. The name of the dependency in the API Hub.
	//
	//  Format: `projects/{project}/locations/{location}/dependencies/{dependency}`
	// +kcc:proto:field=google.cloud.apihub.v1.Dependency.name
	Name *string `json:"name,omitempty"`

	// Required. Immutable. The entity acting as the consumer in the dependency.
	// +kcc:proto:field=google.cloud.apihub.v1.Dependency.consumer
	Consumer *DependencyEntityReference `json:"consumer,omitempty"`

	// Required. Immutable. The entity acting as the supplier in the dependency.
	// +kcc:proto:field=google.cloud.apihub.v1.Dependency.supplier
	Supplier *DependencyEntityReference `json:"supplier,omitempty"`

	// Optional. Human readable description corresponding of the dependency.
	// +kcc:proto:field=google.cloud.apihub.v1.Dependency.description
	Description *string `json:"description,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.apihub.v1.DependencyEntityReference
type DependencyEntityReference struct {
	// The resource name of an operation in the API Hub.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}/operations/{operation}`
	// +kcc:proto:field=google.cloud.apihub.v1.DependencyEntityReference.operation_resource_name
	OperationResourceName *string `json:"operationResourceName,omitempty"`

	// The resource name of an external API in the API Hub.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/externalApis/{external_api}`
	// +kcc:proto:field=google.cloud.apihub.v1.DependencyEntityReference.external_api_resource_name
	ExternalApiResourceName *string `json:"externalApiResourceName,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.DependencyErrorDetail
type DependencyErrorDetail struct {
	// Optional. Error in the dependency.
	// +kcc:proto:field=google.cloud.apihub.v1.DependencyErrorDetail.error
	Error *string `json:"error,omitempty"`

	// Optional. Timestamp at which the error was found.
	// +kcc:proto:field=google.cloud.apihub.v1.DependencyErrorDetail.error_time
	ErrorTime *string `json:"errorTime,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.Dependency
type DependencyObservedState struct {
	// Required. Immutable. The entity acting as the consumer in the dependency.
	// +kcc:proto:field=google.cloud.apihub.v1.Dependency.consumer
	Consumer *DependencyEntityReferenceObservedState `json:"consumer,omitempty"`

	// Output only. State of the dependency.
	// +kcc:proto:field=google.cloud.apihub.v1.Dependency.state
	State *string `json:"state,omitempty"`

	// Output only. Discovery mode of the dependency.
	// +kcc:proto:field=google.cloud.apihub.v1.Dependency.discovery_mode
	DiscoveryMode *string `json:"discoveryMode,omitempty"`

	// Output only. Error details of a dependency if the system has detected it
	//  internally.
	// +kcc:proto:field=google.cloud.apihub.v1.Dependency.error_detail
	ErrorDetail *DependencyErrorDetail `json:"errorDetail,omitempty"`

	// Output only. The time at which the dependency was created.
	// +kcc:proto:field=google.cloud.apihub.v1.Dependency.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the dependency was last updated.
	// +kcc:proto:field=google.cloud.apihub.v1.Dependency.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.DependencyEntityReference
type DependencyEntityReferenceObservedState struct {
	// Output only. Display name of the entity.
	// +kcc:proto:field=google.cloud.apihub.v1.DependencyEntityReference.display_name
	DisplayName *string `json:"displayName,omitempty"`
}
