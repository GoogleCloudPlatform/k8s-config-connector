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

// +kcc:proto=google.cloud.datacatalog.v1.FieldType
type FieldType struct {
	// Primitive types, such as string, boolean, etc.
	// +kcc:proto:field=google.cloud.datacatalog.v1.FieldType.primitive_type
	PrimitiveType *string `json:"primitiveType,omitempty"`

	// An enum type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.FieldType.enum_type
	EnumType *FieldType_EnumType `json:"enumType,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.FieldType.EnumType
type FieldType_EnumType struct {
	// The set of allowed values for this enum.
	//
	//  This set must not be empty and can include up to 100 allowed values.
	//  The display names of the values in this set must not be empty and must
	//  be case-insensitively unique within this set.
	//
	//  The order of items in this set is preserved. This field can be used to
	//  create, remove, and reorder enum values. To rename enum values, use the
	//  `RenameTagTemplateFieldEnumValue` method.
	// +kcc:proto:field=google.cloud.datacatalog.v1.FieldType.EnumType.allowed_values
	AllowedValues []FieldType_EnumType_EnumValue `json:"allowedValues,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.FieldType.EnumType.EnumValue
type FieldType_EnumType_EnumValue struct {
	// Required. The display name of the enum value. Must not be an empty
	// Required. The display name of the enum value. Must not be an empty
	//  string.
	//
	//  The name must contain only Unicode letters, numbers (0-9), underscores
	//  (_), dashes (-), spaces ( ), and can't start or end with spaces. The
	//  maximum length is 200 characters.
	//+required
	// +kcc:proto:field=google.cloud.datacatalog.v1.FieldType.EnumType.EnumValue.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.TagTemplateField
type TagTemplateField struct {
	// Identifier. The resource name of the tag template field in URL format.
	//  Example:
	//
	//  `projects/{PROJECT_ID}/locations/{LOCATION}/tagTemplates/{TAG_TEMPLATE}/fields/{FIELD}`
	//
	//  Note: The tag template field itself might not be stored in the location
	//  specified in its name.
	//
	//  The name must contain only letters (a-z, A-Z), numbers (0-9),
	//  or underscores (_), and must start with a letter or underscore.
	//  The maximum length is 64 characters.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplateField.name
	Name *string `json:"name,omitempty"`

	// The display name for this field. Defaults to an empty string.
	//
	//  The name must contain only Unicode letters, numbers (0-9), underscores (_),
	//  dashes (-), spaces ( ), and can't start or end with spaces.
	//  The maximum length is 200 characters.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplateField.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The type of value this tag field can contain.
	//+required
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplateField.type
	Type *FieldType `json:"type,omitempty"`

	// If true, this field is required. Defaults to false.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplateField.is_required
	IsRequired *bool `json:"isRequired,omitempty"`

	// The description for this field. Defaults to an empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplateField.description
	Description *string `json:"description,omitempty"`

	// The order of this field with respect to other fields in this tag
	//  template.
	//
	//  For example, a higher value can indicate a more important field.
	//  The value can be negative. Multiple fields can have the same order and
	//  field orders within a tag don't have to be sequential.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplateField.order
	Order *int32 `json:"order,omitempty"`
}
