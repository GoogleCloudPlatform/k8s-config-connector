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

// +kcc:proto=google.cloud.apihub.v1.Attribute
type Attribute struct {
	// Identifier. The name of the attribute in the API Hub.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/attributes/{attribute}`
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.name
	Name *string `json:"name,omitempty"`

	// Required. The display name of the attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.description
	Description *string `json:"description,omitempty"`

	// Required. The scope of the attribute. It represents the resource in the API
	//  Hub to which the attribute can be linked.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.scope
	Scope *string `json:"scope,omitempty"`

	// Required. The type of the data of the attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.data_type
	DataType *string `json:"dataType,omitempty"`

	// Optional. The list of allowed values when the attribute value is of type
	//  enum. This is required when the data_type of the attribute is ENUM. The
	//  maximum number of allowed values of an attribute will be 1000.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.allowed_values
	AllowedValues []Attribute_AllowedValue `json:"allowedValues,omitempty"`

	// Optional. The maximum number of values that the attribute can have when
	//  associated with an API Hub resource. Cardinality 1 would represent a
	//  single-valued attribute. It must not be less than 1 or greater than 20. If
	//  not specified, the cardinality would be set to 1 by default and represent a
	//  single-valued attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.cardinality
	Cardinality *int32 `json:"cardinality,omitempty"`
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

// +kcc:proto=google.cloud.apihub.v1.Attribute
type AttributeObservedState struct {
	// Output only. The definition type of the attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.definition_type
	DefinitionType *string `json:"definitionType,omitempty"`

	// Output only. When mandatory is true, the attribute is mandatory for the
	//  resource specified in the scope. Only System defined attributes can be
	//  mandatory.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.mandatory
	Mandatory *bool `json:"mandatory,omitempty"`

	// Output only. The time at which the attribute was created.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the attribute was last updated.
	// +kcc:proto:field=google.cloud.apihub.v1.Attribute.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
