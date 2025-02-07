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


// +kcc:proto=google.cloud.backupdr.v1.BackupApplianceBackupConfig
type BackupApplianceBackupConfig struct {
	// The name of the backup appliance.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceBackupConfig.backup_appliance_name
	BackupApplianceName *string `json:"backupApplianceName,omitempty"`

	// The ID of the backup appliance.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceBackupConfig.backup_appliance_id
	BackupApplianceID *int64 `json:"backupApplianceID,omitempty"`

	// The ID of the SLA of this application.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceBackupConfig.sla_id
	SlaID *int64 `json:"slaID,omitempty"`

	// The name of the application.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceBackupConfig.application_name
	ApplicationName *string `json:"applicationName,omitempty"`

	// The name of the host where the application is running.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceBackupConfig.host_name
	HostName *string `json:"hostName,omitempty"`

	// The name of the SLT associated with the application.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceBackupConfig.slt_name
	SltName *string `json:"sltName,omitempty"`

	// The name of the SLP associated with the application.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupApplianceBackupConfig.slp_name
	SlpName *string `json:"slpName,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.BackupConfigInfo
type BackupConfigInfo struct {

	// Configuration for a Google Cloud resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupConfigInfo.gcp_backup_config
	GcpBackupConfig *GcpBackupConfig `json:"gcpBackupConfig,omitempty"`

	// Configuration for an application backed up by a Backup Appliance.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupConfigInfo.backup_appliance_backup_config
	BackupApplianceBackupConfig *BackupApplianceBackupConfig `json:"backupApplianceBackupConfig,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.ComputeInstanceDataSourceProperties
type ComputeInstanceDataSourceProperties struct {
	// Name of the compute instance backed up by the datasource.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceDataSourceProperties.name
	Name *string `json:"name,omitempty"`

	// The description of the Compute Engine instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceDataSourceProperties.description
	Description *string `json:"description,omitempty"`

	// The machine type of the instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceDataSourceProperties.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// The total number of disks attached to the Instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceDataSourceProperties.total_disk_count
	TotalDiskCount *int64 `json:"totalDiskCount,omitempty"`

	// The sum of all the disk sizes.
	// +kcc:proto:field=google.cloud.backupdr.v1.ComputeInstanceDataSourceProperties.total_disk_size_gb
	TotalDiskSizeGB *int64 `json:"totalDiskSizeGB,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.DataSource
type DataSource struct {

	// Optional. Resource labels to represent user provided metadata.
	//  No labels currently defined:
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSource.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Number of backups in the data source.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSource.backup_count
	BackupCount *int64 `json:"backupCount,omitempty"`

	// Server specified ETag for the ManagementServer resource to prevent
	//  simultaneous updates from overwiting each other.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSource.etag
	Etag *string `json:"etag,omitempty"`

	// The number of bytes (metadata and data) stored in this datasource.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSource.total_stored_bytes
	TotalStoredBytes *int64 `json:"totalStoredBytes,omitempty"`

	// The backed up resource is a Google Cloud resource.
	//  The word 'DataSource' was included in the names to indicate that this is
	//  the representation of the Google Cloud resource used within the
	//  DataSource object.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSource.data_source_gcp_resource
	DataSourceGcpResource *DataSourceGcpResource `json:"dataSourceGcpResource,omitempty"`

	// The backed up resource is a backup appliance application.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSource.data_source_backup_appliance_application
	DataSourceBackupApplianceApplication *DataSourceBackupApplianceApplication `json:"dataSourceBackupApplianceApplication,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.DataSourceBackupApplianceApplication
type DataSourceBackupApplianceApplication struct {
	// The name of the Application as known to the Backup Appliance.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSourceBackupApplianceApplication.application_name
	ApplicationName *string `json:"applicationName,omitempty"`

	// Appliance name.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSourceBackupApplianceApplication.backup_appliance
	BackupAppliance *string `json:"backupAppliance,omitempty"`

	// Appliance Id of the Backup Appliance.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSourceBackupApplianceApplication.appliance_id
	ApplianceID *int64 `json:"applianceID,omitempty"`

	// The type of the application. e.g. VMBackup
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSourceBackupApplianceApplication.type
	Type *string `json:"type,omitempty"`

	// The appid field of the application within the Backup Appliance.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSourceBackupApplianceApplication.application_id
	ApplicationID *int64 `json:"applicationID,omitempty"`

	// Hostname of the host where the application is running.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSourceBackupApplianceApplication.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Hostid of the application host.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSourceBackupApplianceApplication.host_id
	HostID *int64 `json:"hostID,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.DataSourceGcpResource
type DataSourceGcpResource struct {

	// Location of the resource: <region>/<zone>/"global"/"unspecified".
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSourceGcpResource.location
	Location *string `json:"location,omitempty"`

	// The type of the Google Cloud resource. Use the Unified Resource Type,
	//  eg. compute.googleapis.com/Instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSourceGcpResource.type
	Type *string `json:"type,omitempty"`

	// ComputeInstanceDataSourceProperties has a subset of Compute Instance
	//  properties that are useful at the Datasource level.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSourceGcpResource.compute_instance_datasource_properties
	ComputeInstanceDatasourceProperties *ComputeInstanceDataSourceProperties `json:"computeInstanceDatasourceProperties,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.GcpBackupConfig
type GcpBackupConfig struct {
	// The name of the backup plan.
	// +kcc:proto:field=google.cloud.backupdr.v1.GcpBackupConfig.backup_plan
	BackupPlan *string `json:"backupPlan,omitempty"`

	// The description of the backup plan.
	// +kcc:proto:field=google.cloud.backupdr.v1.GcpBackupConfig.backup_plan_description
	BackupPlanDescription *string `json:"backupPlanDescription,omitempty"`

	// The name of the backup plan association.
	// +kcc:proto:field=google.cloud.backupdr.v1.GcpBackupConfig.backup_plan_association
	BackupPlanAssociation *string `json:"backupPlanAssociation,omitempty"`

	// The names of the backup plan rules which point to this backupvault
	// +kcc:proto:field=google.cloud.backupdr.v1.GcpBackupConfig.backup_plan_rules
	BackupPlanRules []string `json:"backupPlanRules,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.BackupConfigInfo
type BackupConfigInfoObservedState struct {
	// Output only. The status of the last backup to this BackupVault
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupConfigInfo.last_backup_state
	LastBackupState *string `json:"lastBackupState,omitempty"`

	// Output only. If the last backup were successful, this field has the
	//  consistency date.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupConfigInfo.last_successful_backup_consistency_time
	LastSuccessfulBackupConsistencyTime *string `json:"lastSuccessfulBackupConsistencyTime,omitempty"`

	// Output only. If the last backup failed, this field has the error message.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupConfigInfo.last_backup_error
	LastBackupError *Status `json:"lastBackupError,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.DataSource
type DataSourceObservedState struct {
	// Output only. Identifier. Name of the datasource to create.
	//  It must have the
	//  format`"projects/{project}/locations/{location}/backupVaults/{backupvault}/dataSources/{datasource}"`.
	//  `{datasource}` cannot be changed after creation. It must be between 3-63
	//  characters long and must be unique within the backup vault.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSource.name
	Name *string `json:"name,omitempty"`

	// Output only. The DataSource resource instance state.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSource.state
	State *string `json:"state,omitempty"`

	// Output only. The time when the instance was created.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSource.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the instance was updated.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSource.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The backup configuration state.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSource.config_state
	ConfigState *string `json:"configState,omitempty"`

	// Output only. Details of how the resource is configured for backup.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSource.backup_config_info
	BackupConfigInfo *BackupConfigInfo `json:"backupConfigInfo,omitempty"`

	// The backed up resource is a Google Cloud resource.
	//  The word 'DataSource' was included in the names to indicate that this is
	//  the representation of the Google Cloud resource used within the
	//  DataSource object.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSource.data_source_gcp_resource
	DataSourceGcpResource *DataSourceGcpResourceObservedState `json:"dataSourceGcpResource,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.DataSourceGcpResource
type DataSourceGcpResourceObservedState struct {
	// Output only. Full resource pathname URL of the source Google Cloud
	//  resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.DataSourceGcpResource.gcp_resourcename
	GcpResourcename *string `json:"gcpResourcename,omitempty"`
}
