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

// +kcc:proto=google.cloud.dataplex.v1.AssetStatus
type AssetStatus struct {
	// Last update time of the status.
	// +kcc:proto:field=google.cloud.dataplex.v1.AssetStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Number of active assets.
	// +kcc:proto:field=google.cloud.dataplex.v1.AssetStatus.active_assets
	ActiveAssets *int32 `json:"activeAssets,omitempty"`

	// Number of assets that are in process of updating the security policy on
	//  attached resources.
	// +kcc:proto:field=google.cloud.dataplex.v1.AssetStatus.security_policy_applying_assets
	SecurityPolicyApplyingAssets *int32 `json:"securityPolicyApplyingAssets,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Lake
type Lake struct {

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.description
	Description *string `json:"description,omitempty"`

	// Optional. Settings to manage lake and Dataproc Metastore service instance
	//  association.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.metastore
	Metastore *Lake_Metastore `json:"metastore,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Lake.MetastoreStatus
type Lake_MetastoreStatus struct {
	// Current state of association.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.MetastoreStatus.state
	State *string `json:"state,omitempty"`

	// Additional information about the current status.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.MetastoreStatus.message
	Message *string `json:"message,omitempty"`

	// Last update time of the metastore status of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.MetastoreStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// The URI of the endpoint used to access the Metastore service.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.MetastoreStatus.endpoint
	Endpoint *string `json:"endpoint,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Lake
type LakeObservedState struct {
	// Output only. The relative resource name of the lake, of the form:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the lake. This ID will
	//  be different if the lake is deleted and re-created with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the lake was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the lake was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.state
	State *string `json:"state,omitempty"`

	// Output only. Service account associated with this lake. This service
	//  account must be authorized to access or operate on resources managed by the
	//  lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Output only. Aggregated status of the underlying assets of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.asset_status
	AssetStatus *AssetStatus `json:"assetStatus,omitempty"`

	// Output only. Metastore status of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.metastore_status
	MetastoreStatus *Lake_MetastoreStatus `json:"metastoreStatus,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Content
type Content struct {

	// Required. The path for the Content file, represented as directory
	//  structure. Unique within a lake. Limited to alphanumerics, hyphens,
	//  underscores, dots and slashes.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.path
	Path *string `json:"path,omitempty"`

	// Optional. User defined labels for the content.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the content.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.description
	Description *string `json:"description,omitempty"`

	// Required. Content data in string format.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.data_text
	DataText *string `json:"dataText,omitempty"`

	// Sql Script related configurations.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.sql_script
	SQLScript *Content_SQLScript `json:"sqlScript,omitempty"`

	// Notebook related configurations.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.notebook
	Notebook *Content_Notebook `json:"notebook,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Content.Notebook
type Content_Notebook struct {
	// Required. Kernel Type of the notebook.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.Notebook.kernel_type
	KernelType *string `json:"kernelType,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Content.SqlScript
type Content_SQLScript struct {
	// Required. Query Engine to be used for the Sql Query.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.SqlScript.engine
	Engine *string `json:"engine,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Content
type ContentObservedState struct {
	// Output only. The relative resource name of the content, of the form:
	//  projects/{project_id}/locations/{location_id}/lakes/{lake_id}/content/{content_id}
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the content. This ID
	//  will be different if the content is deleted and re-created with the same
	//  name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Content creation time.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the content was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment
type Environment struct {

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User defined labels for the environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.description
	Description *string `json:"description,omitempty"`

	// Required. Infrastructure specification for the Environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.infrastructure_spec
	InfrastructureSpec *Environment_InfrastructureSpec `json:"infrastructureSpec,omitempty"`

	// Optional. Configuration for sessions created for this environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.session_spec
	SessionSpec *Environment_SessionSpec `json:"sessionSpec,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.Endpoints
type Environment_Endpoints struct {
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.InfrastructureSpec
type Environment_InfrastructureSpec struct {
	// Optional. Compute resources needed for analyze interactive workloads.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.compute
	Compute *Environment_InfrastructureSpec_ComputeResources `json:"compute,omitempty"`

	// Required. Software Runtime Configuration for analyze interactive
	//  workloads.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.os_image
	OSImage *Environment_InfrastructureSpec_OSImageRuntime `json:"osImage,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.InfrastructureSpec.ComputeResources
type Environment_InfrastructureSpec_ComputeResources struct {
	// Optional. Size in GB of the disk. Default is 100 GB.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.ComputeResources.disk_size_gb
	DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`

	// Optional. Total number of nodes in the sessions created for this
	//  environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.ComputeResources.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Optional. Max configurable nodes.
	//  If max_node_count > node_count, then auto-scaling is enabled.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.ComputeResources.max_node_count
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.InfrastructureSpec.OsImageRuntime
type Environment_InfrastructureSpec_OSImageRuntime struct {
	// Required. Dataplex Image version.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.OsImageRuntime.image_version
	ImageVersion *string `json:"imageVersion,omitempty"`

	// Optional. List of Java jars to be included in the runtime environment.
	//  Valid input includes Cloud Storage URIs to Jar binaries.
	//  For example, gs://bucket-name/my/path/to/file.jar
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.OsImageRuntime.java_libraries
	JavaLibraries []string `json:"javaLibraries,omitempty"`

	// Optional. A list of python packages to be installed.
	//  Valid formats include Cloud Storage URI to a PIP installable library.
	//  For example, gs://bucket-name/my/path/to/lib.tar.gz
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.OsImageRuntime.python_packages
	PythonPackages []string `json:"pythonPackages,omitempty"`

	// Optional. Spark properties to provide configuration for use in sessions
	//  created for this environment. The properties to set on daemon config
	//  files. Property keys are specified in `prefix:property` format. The
	//  prefix must be "spark".
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.InfrastructureSpec.OsImageRuntime.properties
	Properties map[string]string `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.SessionSpec
type Environment_SessionSpec struct {
	// Optional. The idle time configuration of the session. The session will be
	//  auto-terminated at the end of this period.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.SessionSpec.max_idle_duration
	MaxIdleDuration *string `json:"maxIdleDuration,omitempty"`

	// Optional. If True, this causes sessions to be pre-created and available
	//  for faster startup to enable interactive exploration use-cases. This
	//  defaults to False to avoid additional billed charges. These can only be
	//  set to True for the environment with name set to "default", and with
	//  default configuration.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.SessionSpec.enable_fast_startup
	EnableFastStartup *bool `json:"enableFastStartup,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.SessionStatus
type Environment_SessionStatus struct {
}

// +kcc:proto=google.cloud.dataplex.v1.Environment
type EnvironmentObservedState struct {
	// Output only. The relative resource name of the environment, of the form:
	//  projects/{project_id}/locations/{location_id}/lakes/{lake_id}/environment/{environment_id}
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the environment. This
	//  ID will be different if the environment is deleted and re-created with the
	//  same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Environment creation time.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the environment was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.state
	State *string `json:"state,omitempty"`

	// Output only. Status of sessions created for this environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.session_status
	SessionStatus *Environment_SessionStatusObservedState `json:"sessionStatus,omitempty"`

	// Output only. URI Endpoints to access sessions associated with the
	//  Environment.
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.endpoints
	Endpoints *Environment_EndpointsObservedState `json:"endpoints,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.Endpoints
type Environment_EndpointsObservedState struct {
	// Output only. URI to serve notebook APIs
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.Endpoints.notebooks
	Notebooks *string `json:"notebooks,omitempty"`

	// Output only. URI to serve SQL APIs
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.Endpoints.sql
	SQL *string `json:"sql,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Environment.SessionStatus
type Environment_SessionStatusObservedState struct {
	// Output only. Queries over sessions to mark whether the environment is
	//  currently active or not
	// +kcc:proto:field=google.cloud.dataplex.v1.Environment.SessionStatus.active
	Active *bool `json:"active,omitempty"`
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
