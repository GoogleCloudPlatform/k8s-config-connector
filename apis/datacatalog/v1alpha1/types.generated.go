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


// +kcc:proto=google.cloud.datacatalog.v1beta1.FieldType
type FieldType struct {
	// Represents primitive types - string, bool etc.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.FieldType.primitive_type
	PrimitiveType *string `json:"primitiveType,omitempty"`

	// Represents an enum type.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.FieldType.enum_type
	EnumType *FieldType_EnumType `json:"enumType,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.FieldType.EnumType
type FieldType_EnumType struct {
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.FieldType.EnumType.allowed_values
	AllowedValues []FieldType_EnumType_EnumValue `json:"allowedValues,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.FieldType.EnumType.EnumValue
type FieldType_EnumType_EnumValue struct {
	// Required. The display name of the enum value. Must not be an empty
	//  string.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.FieldType.EnumType.EnumValue.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.TagTemplateField
type TagTemplateField struct {

	// The display name for this field. Defaults to an empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.TagTemplateField.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The type of value this tag field can contain.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.TagTemplateField.type
	Type *FieldType `json:"type,omitempty"`

	// Whether this is a required field. Defaults to false.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.TagTemplateField.is_required
	IsRequired *bool `json:"isRequired,omitempty"`

	// The description for this field. Defaults to an empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.TagTemplateField.description
	Description *string `json:"description,omitempty"`

	// The order of this field with respect to other fields in this tag
	//  template.  A higher value indicates a more important field. The value can
	//  be negative. Multiple fields can have the same order, and field orders
	//  within a tag do not have to be sequential.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.TagTemplateField.order
	Order *int32 `json:"order,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.TagTemplateField
type TagTemplateFieldObservedState struct {
	// Output only. Identifier. The resource name of the tag template field in URL
	//  format. Example:
	//
	//  * projects/{project_id}/locations/{location}/tagTemplates/{tag_template}/fields/{field}
	//
	//  Note that this TagTemplateField may not actually be stored in the location
	//  in this name.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.TagTemplateField.name
	Name *string `json:"name,omitempty"`
}
