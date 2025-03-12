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


// +kcc:proto=google.cloud.dataplex.v1.AspectType
type AspectType struct {

	// Optional. Description of the AspectType.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.description
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the AspectType.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The service computes this checksum. The client may send it on update and
	//  delete requests to ensure it has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.etag
	Etag *string `json:"etag,omitempty"`

	// Immutable. Defines the Authorization for this type.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.authorization
	Authorization *AspectType_Authorization `json:"authorization,omitempty"`

	// Required. MetadataTemplate of the aspect.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.metadata_template
	MetadataTemplate *AspectType_MetadataTemplate `json:"metadataTemplate,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.AspectType.Authorization
type AspectType_Authorization struct {
	// Immutable. The IAM permission grantable on the EntryGroup to allow access
	//  to instantiate Aspects of Dataplex owned AspectTypes, only settable for
	//  Dataplex owned Types.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.Authorization.alternate_use_permission
	AlternateUsePermission *string `json:"alternateUsePermission,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.AspectType.MetadataTemplate
type AspectType_MetadataTemplate struct {
	// Optional. Index is used to encode Template messages. The value of index
	//  can range between 1 and 2,147,483,647. Index must be unique within all
	//  fields in a Template. (Nested Templates can reuse indexes). Once a
	//  Template is defined, the index cannot be changed, because it identifies
	//  the field in the actual storage format. Index is a mandatory field, but
	//  it is optional for top level fields, and map/array "values" definitions.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.index
	Index *int32 `json:"index,omitempty"`

	// Required. The name of the field.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.name
	Name *string `json:"name,omitempty"`

	// Required. The datatype of this field. The following values are supported:
	//
	//  Primitive types:
	//
	//  * string
	//  * integer
	//  * boolean
	//  * double
	//  * datetime. Must be of the format RFC3339 UTC "Zulu" (Examples:
	//  "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z").
	//
	//  Complex types:
	//
	//  * enum
	//  * array
	//  * map
	//  * record
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.type
	Type *string `json:"type,omitempty"`

	// Optional. Field definition. You must specify it if the type is record. It
	//  defines the nested fields.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.record_fields
	RecordFields []AspectType_MetadataTemplate `json:"recordFields,omitempty"`

	// Optional. The list of values for an enum type. You must define it if the
	//  type is enum.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.enum_values
	EnumValues []AspectType_MetadataTemplate_EnumValue `json:"enumValues,omitempty"`

	// Optional. If the type is map, set map_items. map_items can refer to a
	//  primitive field or a complex (record only) field. To specify a primitive
	//  field, you only need to set name and type in the nested
	//  MetadataTemplate. The recommended value for the name field is item, as
	//  this isn't used in the actual payload.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.map_items
	MapItems *AspectType_MetadataTemplate `json:"mapItems,omitempty"`

	// Optional. If the type is array, set array_items. array_items can refer
	//  to a primitive field or a complex (record only) field. To specify a
	//  primitive field, you only need to set name and type in the nested
	//  MetadataTemplate. The recommended value for the name field is item, as
	//  this isn't used in the actual payload.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.array_items
	ArrayItems *AspectType_MetadataTemplate `json:"arrayItems,omitempty"`

	// Optional. You can use type id if this definition of the field needs to be
	//  reused later. The type id must be unique across the entire template. You
	//  can only specify it if the field type is record.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.type_id
	TypeID *string `json:"typeID,omitempty"`

	// Optional. A reference to another field definition (not an inline
	//  definition). The value must be equal to the value of an id field defined
	//  elsewhere in the MetadataTemplate. Only fields with record type can
	//  refer to other fields.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.type_ref
	TypeRef *string `json:"typeRef,omitempty"`

	// Optional. Specifies the constraints on this field.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.constraints
	Constraints *AspectType_MetadataTemplate_Constraints `json:"constraints,omitempty"`

	// Optional. Specifies annotations on this field.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.annotations
	Annotations *AspectType_MetadataTemplate_Annotations `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.AspectType.MetadataTemplate.Annotations
type AspectType_MetadataTemplate_Annotations struct {
	// Optional. Marks a field as deprecated. You can include a deprecation
	//  message.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.Annotations.deprecated
	Deprecated *string `json:"deprecated,omitempty"`

	// Optional. Display name for a field.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.Annotations.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description for a field.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.Annotations.description
	Description *string `json:"description,omitempty"`

	// Optional. Display order for a field. You can use this to reorder where
	//  a field is rendered.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.Annotations.display_order
	DisplayOrder *int32 `json:"displayOrder,omitempty"`

	// Optional. You can use String Type annotations to specify special
	//  meaning to string fields. The following values are supported:
	//
	//  * richText: The field must be interpreted as a rich text field.
	//  * url: A fully qualified URL link.
	//  * resource: A service qualified resource reference.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.Annotations.string_type
	StringType *string `json:"stringType,omitempty"`

	// Optional. Suggested hints for string fields. You can use them to
	//  suggest values to users through console.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.Annotations.string_values
	StringValues []string `json:"stringValues,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.AspectType.MetadataTemplate.Constraints
type AspectType_MetadataTemplate_Constraints struct {
	// Optional. Marks this field as optional or required.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.Constraints.required
	Required *bool `json:"required,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.AspectType.MetadataTemplate.EnumValue
type AspectType_MetadataTemplate_EnumValue struct {
	// Required. Index for the enum value. It can't be modified.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.EnumValue.index
	Index *int32 `json:"index,omitempty"`

	// Required. Name of the enumvalue. This is the actual value that the
	//  aspect can contain.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.EnumValue.name
	Name *string `json:"name,omitempty"`

	// Optional. You can set this message if you need to deprecate an enum
	//  value.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.MetadataTemplate.EnumValue.deprecated
	Deprecated *string `json:"deprecated,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.AspectType
type AspectTypeObservedState struct {
	// Output only. The relative resource name of the AspectType, of the form:
	//  projects/{project_number}/locations/{location_id}/aspectTypes/{aspect_type_id}.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the AspectType.
	//  If you delete and recreate the AspectType with the same name, then this ID
	//  will be different.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the AspectType was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the AspectType was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Denotes the transfer status of the Aspect Type. It is
	//  unspecified for Aspect Types created from Dataplex API.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.transfer_status
	TransferStatus *string `json:"transferStatus,omitempty"`
}
