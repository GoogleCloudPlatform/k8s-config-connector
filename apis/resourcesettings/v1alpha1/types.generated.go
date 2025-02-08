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


// +kcc:proto=google.cloud.resourcesettings.v1.Setting
type Setting struct {
	// The resource name of the setting. Must be in one of the following forms:
	//
	//  * `projects/{project_number}/settings/{setting_name}`
	//  * `folders/{folder_id}/settings/{setting_name}`
	//  * `organizations/{organization_id}/settings/{setting_name}`
	//
	//  For example, "/projects/123/settings/gcp-enableMyFeature"
	// +kcc:proto:field=google.cloud.resourcesettings.v1.Setting.name
	Name *string `json:"name,omitempty"`

	// The configured value of the setting at the given parent resource (ignoring
	//  the resource hierarchy). The data type of
	//  [Value][google.cloud.resourcesettings.v1.Value] must always be consistent
	//  with the data type defined in
	//  [Setting.metadata][google.cloud.resourcesettings.v1.Setting.metadata].
	// +kcc:proto:field=google.cloud.resourcesettings.v1.Setting.local_value
	LocalValue *Value `json:"localValue,omitempty"`

	// A fingerprint used for optimistic concurrency. See
	//  [UpdateSetting][google.cloud.resourcesettings.v1.ResourceSettingsService.UpdateSetting]
	//  for more details.
	// +kcc:proto:field=google.cloud.resourcesettings.v1.Setting.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.resourcesettings.v1.SettingMetadata
type SettingMetadata struct {
	// The human readable name for this setting.
	// +kcc:proto:field=google.cloud.resourcesettings.v1.SettingMetadata.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A detailed description of what this setting does.
	// +kcc:proto:field=google.cloud.resourcesettings.v1.SettingMetadata.description
	Description *string `json:"description,omitempty"`

	// A flag indicating that values of this setting cannot be modified (see
	//  documentation of the specific setting for updates and reasons).
	// +kcc:proto:field=google.cloud.resourcesettings.v1.SettingMetadata.read_only
	ReadOnly *bool `json:"readOnly,omitempty"`

	// The data type for this setting.
	// +kcc:proto:field=google.cloud.resourcesettings.v1.SettingMetadata.data_type
	DataType *string `json:"dataType,omitempty"`

	// The value provided by
	//  [Setting.effective_value][google.cloud.resourcesettings.v1.Setting.effective_value]
	//  if no setting value is explicitly set.
	//
	//  Note: not all settings have a default value.
	// +kcc:proto:field=google.cloud.resourcesettings.v1.SettingMetadata.default_value
	DefaultValue *Value `json:"defaultValue,omitempty"`
}

// +kcc:proto=google.cloud.resourcesettings.v1.Value
type Value struct {
	// Defines this value as being a boolean value.
	// +kcc:proto:field=google.cloud.resourcesettings.v1.Value.boolean_value
	BooleanValue *bool `json:"booleanValue,omitempty"`

	// Defines this value as being a string value.
	// +kcc:proto:field=google.cloud.resourcesettings.v1.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Defines this value as being a StringSet.
	// +kcc:proto:field=google.cloud.resourcesettings.v1.Value.string_set_value
	StringSetValue *Value_StringSet `json:"stringSetValue,omitempty"`

	// Defines this value as being a Enum.
	// +kcc:proto:field=google.cloud.resourcesettings.v1.Value.enum_value
	EnumValue *Value_EnumValue `json:"enumValue,omitempty"`
}

// +kcc:proto=google.cloud.resourcesettings.v1.Value.EnumValue
type Value_EnumValue struct {
	// The value of this enum
	// +kcc:proto:field=google.cloud.resourcesettings.v1.Value.EnumValue.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.resourcesettings.v1.Value.StringSet
type Value_StringSet struct {
	// The strings in the set
	// +kcc:proto:field=google.cloud.resourcesettings.v1.Value.StringSet.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.resourcesettings.v1.Setting
type SettingObservedState struct {
	// Output only. Metadata about a setting which is not editable by the end
	//  user.
	// +kcc:proto:field=google.cloud.resourcesettings.v1.Setting.metadata
	Metadata *SettingMetadata `json:"metadata,omitempty"`

	// Output only. The computed effective value of the setting at the given
	//  parent resource (based on the resource hierarchy).
	//
	//  The effective value evaluates to one of the following options in the given
	//  order (the next option is used if the previous one does not exist):
	//
	//  1. the local setting value on the given resource:
	//  [Setting.local_value][google.cloud.resourcesettings.v1.Setting.local_value]
	//  2. if one of the given resource's ancestors have a local setting value,
	//     the local value at the nearest such ancestor
	//  3. the setting's default value:
	//  [SettingMetadata.default_value][google.cloud.resourcesettings.v1.SettingMetadata.default_value]
	//  4. an empty value (defined as a `Value` with all fields unset)
	//
	//  The data type of [Value][google.cloud.resourcesettings.v1.Value] must
	//  always be consistent with the data type defined in
	//  [Setting.metadata][google.cloud.resourcesettings.v1.Setting.metadata].
	// +kcc:proto:field=google.cloud.resourcesettings.v1.Setting.effective_value
	EffectiveValue *Value `json:"effectiveValue,omitempty"`
}
