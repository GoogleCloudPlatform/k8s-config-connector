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

	//+required
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

// +kcc:proto=google.cloud.dataplex.v1.EntryGroup
type EntryGroup struct {

	// Optional. Description of the EntryGroup.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.description
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the EntryGroup.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// This checksum is computed by the service, and might be sent on update and
	//  delete requests to ensure the client has an up-to-date value before
	//  proceeding.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.EntryGroup
type EntryGroupObservedState struct {
	// Output only. The relative resource name of the EntryGroup, in the format
	//  projects/{project_id_or_number}/locations/{location_id}/entryGroups/{entry_group_id}.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the EntryGroup. If you
	//  delete and recreate the EntryGroup with the same name, this ID will be
	//  different.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the EntryGroup was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the EntryGroup was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Denotes the transfer status of the Entry Group. It is
	//  unspecified for Entry Group created from Dataplex API.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryGroup.transfer_status
	TransferStatus *string `json:"transferStatus,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.EntryType
type EntryType struct {

	// Optional. Description of the EntryType.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.description
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the EntryType.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. This checksum is computed by the service, and might be sent on
	//  update and delete requests to ensure the client has an up-to-date value
	//  before proceeding.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Indicates the classes this Entry Type belongs to, for example,
	//  TABLE, DATABASE, MODEL.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.type_aliases
	TypeAliases []string `json:"typeAliases,omitempty"`

	// Optional. The platform that Entries of this type belongs to.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.platform
	Platform *string `json:"platform,omitempty"`

	// Optional. The system that Entries of this type belongs to. Examples include
	//  CloudSQL, MariaDB etc
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.system
	System *string `json:"system,omitempty"`

	// AspectInfo for the entry type.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.required_aspects
	RequiredAspects []EntryType_AspectInfo `json:"requiredAspects,omitempty"`

	// Immutable. Authorization defined for this type.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.authorization
	Authorization *EntryType_Authorization `json:"authorization,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.EntryType.AspectInfo
type EntryType_AspectInfo struct {
	// Required aspect type for the entry type.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.AspectInfo.type
	//+required
	TypeRef *AspectTypeRef `json:"typeRef,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.EntryType.Authorization
type EntryType_Authorization struct {
	// Immutable. The IAM permission grantable on the Entry Group to allow
	//  access to instantiate Entries of Dataplex owned Entry Types, only
	//  settable for Dataplex owned Types.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.Authorization.alternate_use_permission
	AlternateUsePermission *string `json:"alternateUsePermission,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.EntryType
type EntryTypeObservedState struct {
	// Output only. The relative resource name of the EntryType, of the form:
	//  projects/{project_number}/locations/{location_id}/entryTypes/{entry_type_id}.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the EntryType. This ID
	//  will be different if the EntryType is deleted and re-created with the same
	//  name.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the EntryType was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the EntryType was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryType.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Zone
type Zone struct {

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User defined labels for the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. The type of the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.type
	Type *string `json:"type,omitempty"`

	// Optional. Specification of the discovery feature applied to data in this
	//  zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.discovery_spec
	DiscoverySpec *Zone_DiscoverySpec `json:"discoverySpec,omitempty"`

	// Required. Specification of the resources that are referenced by the assets
	//  within this zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.resource_spec
	ResourceSpec *Zone_ResourceSpec `json:"resourceSpec,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Zone.DiscoverySpec
type Zone_DiscoverySpec struct {
	// Required. Whether discovery is enabled.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. The list of patterns to apply for selecting data to include
	//  during discovery if only a subset of the data should considered. For
	//  Cloud Storage bucket assets, these are interpreted as glob patterns used
	//  to match object names. For BigQuery dataset assets, these are interpreted
	//  as patterns to match table names.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.include_patterns
	IncludePatterns []string `json:"includePatterns,omitempty"`

	// Optional. The list of patterns to apply for selecting data to exclude
	//  during discovery.  For Cloud Storage bucket assets, these are interpreted
	//  as glob patterns used to match object names. For BigQuery dataset assets,
	//  these are interpreted as patterns to match table names.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.exclude_patterns
	ExcludePatterns []string `json:"excludePatterns,omitempty"`

	// Optional. Configuration for CSV data.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.csv_options
	CsvOptions *Zone_DiscoverySpec_CsvOptions `json:"csvOptions,omitempty"`

	// Optional. Configuration for Json data.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.json_options
	JsonOptions *Zone_DiscoverySpec_JsonOptions `json:"jsonOptions,omitempty"`

	// Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for
	//  running discovery periodically. Successive discovery runs must be
	//  scheduled at least 60 minutes apart. The default value is to run
	//  discovery every 60 minutes. To explicitly set a timezone to the cron
	//  tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or
	//  TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string
	//  from IANA time zone database. For example, `CRON_TZ=America/New_York 1
	//  * * * *`, or `TZ=America/New_York 1 * * * *`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.schedule
	Schedule *string `json:"schedule,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Zone.DiscoverySpec.CsvOptions
type Zone_DiscoverySpec_CsvOptions struct {
	// Optional. The number of rows to interpret as header rows that should be
	//  skipped when reading data rows.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.CsvOptions.header_rows
	HeaderRows *int32 `json:"headerRows,omitempty"`

	// Optional. The delimiter being used to separate values. This defaults to
	//  ','.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.CsvOptions.delimiter
	Delimiter *string `json:"delimiter,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.CsvOptions.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. Whether to disable the inference of data type for CSV data.
	//  If true, all columns will be registered as strings.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.CsvOptions.disable_type_inference
	DisableTypeInference *bool `json:"disableTypeInference,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Zone.DiscoverySpec.JsonOptions
type Zone_DiscoverySpec_JsonOptions struct {
	// Optional. The character encoding of the data. The default is UTF-8.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.JsonOptions.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Optional. Whether to disable the inference of data type for Json data.
	//  If true, all columns will be registered as their primitive types
	//  (strings, number or boolean).
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.DiscoverySpec.JsonOptions.disable_type_inference
	DisableTypeInference *bool `json:"disableTypeInference,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Zone.ResourceSpec
type Zone_ResourceSpec struct {
	// Required. Immutable. The location type of the resources that are allowed
	//  to be attached to the assets within this zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.ResourceSpec.location_type
	LocationType *string `json:"locationType,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Zone
type ZoneObservedState struct {
	// Output only. The relative resource name of the zone, of the form:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the zone. This ID will
	//  be different if the zone is deleted and re-created with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the zone was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the zone was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.state
	State *string `json:"state,omitempty"`

	// Output only. Aggregated status of the underlying assets of the zone.
	// +kcc:proto:field=google.cloud.dataplex.v1.Zone.asset_status
	AssetStatus *AssetStatus `json:"assetStatus,omitempty"`
>>>>>>> f703b56d0 (conductor: "Generated types and mapper for DataplexZone")
}
