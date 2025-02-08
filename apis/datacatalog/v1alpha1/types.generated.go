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


// +kcc:proto=google.cloud.datacatalog.v1beta1.Tag
type Tag struct {
	// Identifier. The resource name of the tag in URL format. Example:
	//
	//  * projects/{project_id}/locations/{location}/entrygroups/{entry_group_id}/entries/{entry_id}/tags/{tag_id}
	//
	//  where `tag_id` is a system-generated identifier.
	//  Note that this Tag may not actually be stored in the location in this name.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Tag.name
	Name *string `json:"name,omitempty"`

	// Required. The resource name of the tag template that this tag uses.
	//  Example:
	//
	//  * projects/{project_id}/locations/{location}/tagTemplates/{tag_template_id}
	//
	//  This field cannot be modified after creation.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Tag.template
	Template *string `json:"template,omitempty"`

	// Resources like Entry can have schemas associated with them. This scope
	//  allows users to attach tags to an individual column based on that schema.
	//
	//  For attaching a tag to a nested column, use `.` to separate the column
	//  names. Example:
	//
	//  * `outer_column.inner_column`
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Tag.column
	Column *string `json:"column,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.datacatalog.v1beta1.TagField
type TagField struct {

	// Holds the value for a tag field with double type.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.TagField.double_value
	DoubleValue *float64 `json:"doubleValue,omitempty"`

	// Holds the value for a tag field with string type.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.TagField.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Holds the value for a tag field with boolean type.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.TagField.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Holds the value for a tag field with timestamp type.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.TagField.timestamp_value
	TimestampValue *string `json:"timestampValue,omitempty"`

	// Holds the value for a tag field with enum type. This value must be
	//  one of the allowed values in the definition of this enum.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.TagField.enum_value
	EnumValue *TagField_EnumValue `json:"enumValue,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.TagField.EnumValue
type TagField_EnumValue struct {
	// The display name of the enum value.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.TagField.EnumValue.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.Tag
type TagObservedState struct {
	// Output only. The display name of the tag template.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.Tag.template_display_name
	TemplateDisplayName *string `json:"templateDisplayName,omitempty"`
}
