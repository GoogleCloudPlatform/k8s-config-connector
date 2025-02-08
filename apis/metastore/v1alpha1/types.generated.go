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


// +kcc:proto=google.cloud.metastore.v1beta.AuxiliaryVersionConfig
type AuxiliaryVersionConfig struct {
	// The Hive metastore version of the auxiliary service. It must be less
	//  than the primary Hive metastore service's version.
	// +kcc:proto:field=google.cloud.metastore.v1beta.AuxiliaryVersionConfig.version
	Version *string `json:"version,omitempty"`

	// A mapping of Hive metastore configuration key-value pairs to apply to the
	//  auxiliary Hive metastore (configured in `hive-site.xml`) in addition to
	//  the primary version's overrides. If keys are present in both the auxiliary
	//  version's overrides and the primary version's overrides, the value from
	//  the auxiliary version's overrides takes precedence.
	// +kcc:proto:field=google.cloud.metastore.v1beta.AuxiliaryVersionConfig.config_overrides
	ConfigOverrides map[string]string `json:"configOverrides,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.Backup
type Backup struct {
	// Immutable. The relative resource name of the backup, in the following form:
	//
	//  `projects/{project_number}/locations/{location_id}/services/{service_id}/backups/{backup_id}`
	// +kcc:proto:field=google.cloud.metastore.v1beta.Backup.name
	Name *string `json:"name,omitempty"`

	// The description of the backup.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Backup.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.DataCatalogConfig
type DataCatalogConfig struct {
	// Defines whether the metastore metadata should be synced to Data Catalog.
	//  The default value is to disable syncing metastore metadata to Data Catalog.
	// +kcc:proto:field=google.cloud.metastore.v1beta.DataCatalogConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.DataplexConfig
type DataplexConfig struct {

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.metastore.v1beta.EncryptionConfig
type EncryptionConfig struct {
	// The fully qualified customer provided Cloud KMS key name to use for
	//  customer data encryption, in the following form:
	//
	//  `projects/{project_number}/locations/{location_id}/keyRings/{key_ring_id}/cryptoKeys/{crypto_key_id}`.
	// +kcc:proto:field=google.cloud.metastore.v1beta.EncryptionConfig.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.HiveMetastoreConfig
type HiveMetastoreConfig struct {
	// Immutable. The Hive metastore schema version.
	// +kcc:proto:field=google.cloud.metastore.v1beta.HiveMetastoreConfig.version
	Version *string `json:"version,omitempty"`

	// A mapping of Hive metastore configuration key-value pairs to apply to the
	//  Hive metastore (configured in `hive-site.xml`). The mappings
	//  override system defaults (some keys cannot be overridden). These
	//  overrides are also applied to auxiliary versions and can be further
	//  customized in the auxiliary version's `AuxiliaryVersionConfig`.
	// +kcc:proto:field=google.cloud.metastore.v1beta.HiveMetastoreConfig.config_overrides
	ConfigOverrides map[string]string `json:"configOverrides,omitempty"`

	// Information used to configure the Hive metastore service as a service
	//  principal in a Kerberos realm. To disable Kerberos, use the `UpdateService`
	//  method and specify this field's path
	//  (`hive_metastore_config.kerberos_config`) in the request's `update_mask`
	//  while omitting this field from the request's `service`.
	// +kcc:proto:field=google.cloud.metastore.v1beta.HiveMetastoreConfig.kerberos_config
	KerberosConfig *KerberosConfig `json:"kerberosConfig,omitempty"`

	// The protocol to use for the metastore service endpoint. If unspecified,
	//  defaults to `THRIFT`.
	// +kcc:proto:field=google.cloud.metastore.v1beta.HiveMetastoreConfig.endpoint_protocol
	EndpointProtocol *string `json:"endpointProtocol,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.metastore.v1beta.KerberosConfig
type KerberosConfig struct {
	// A Kerberos keytab file that can be used to authenticate a service principal
	//  with a Kerberos Key Distribution Center (KDC).
	// +kcc:proto:field=google.cloud.metastore.v1beta.KerberosConfig.keytab
	Keytab *Secret `json:"keytab,omitempty"`

	// A Kerberos principal that exists in the both the keytab the KDC
	//  to authenticate as. A typical principal is of the form
	//  `primary/instance@REALM`, but there is no exact format.
	// +kcc:proto:field=google.cloud.metastore.v1beta.KerberosConfig.principal
	Principal *string `json:"principal,omitempty"`

	// A Cloud Storage URI that specifies the path to a
	//  krb5.conf file. It is of the form `gs://{bucket_name}/path/to/krb5.conf`,
	//  although the file does not need to be named krb5.conf explicitly.
	// +kcc:proto:field=google.cloud.metastore.v1beta.KerberosConfig.krb5_config_gcs_uri
	Krb5ConfigGcsURI *string `json:"krb5ConfigGcsURI,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.Lake
type Lake struct {
	// The Lake resource name.
	//  Example:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}`
	// +kcc:proto:field=google.cloud.metastore.v1beta.Lake.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.MaintenanceWindow
type MaintenanceWindow struct {
	// The hour of day (0-23) when the window starts.
	// +kcc:proto:field=google.cloud.metastore.v1beta.MaintenanceWindow.hour_of_day
	HourOfDay *Int32Value `json:"hourOfDay,omitempty"`

	// The day of week, when the window starts.
	// +kcc:proto:field=google.cloud.metastore.v1beta.MaintenanceWindow.day_of_week
	DayOfWeek *string `json:"dayOfWeek,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.MetadataExport
type MetadataExport struct {
}

// +kcc:proto=google.cloud.metastore.v1beta.MetadataIntegration
type MetadataIntegration struct {
	// The integration config for the Data Catalog service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.MetadataIntegration.data_catalog_config
	DataCatalogConfig *DataCatalogConfig `json:"dataCatalogConfig,omitempty"`

	// The integration config for the Dataplex service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.MetadataIntegration.dataplex_config
	DataplexConfig *DataplexConfig `json:"dataplexConfig,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.MetadataManagementActivity
type MetadataManagementActivity struct {
}

// +kcc:proto=google.cloud.metastore.v1beta.NetworkConfig
type NetworkConfig struct {
	// Immutable. The consumer-side network configuration for the Dataproc
	//  Metastore instance.
	// +kcc:proto:field=google.cloud.metastore.v1beta.NetworkConfig.consumers
	Consumers []NetworkConfig_Consumer `json:"consumers,omitempty"`

	// Enables custom routes to be imported and exported for the Dataproc
	//  Metastore service's peered VPC network.
	// +kcc:proto:field=google.cloud.metastore.v1beta.NetworkConfig.custom_routes_enabled
	CustomRoutesEnabled *bool `json:"customRoutesEnabled,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.NetworkConfig.Consumer
type NetworkConfig_Consumer struct {
	// Immutable. The subnetwork of the customer project from which an IP
	//  address is reserved and used as the Dataproc Metastore service's
	//  endpoint. It is accessible to hosts in the subnet and to all
	//  hosts in a subnet in the same region and same network. There must
	//  be at least one IP address available in the subnet's primary range. The
	//  subnet is specified in the following form:
	//
	//  `projects/{project_number}/regions/{region_id}/subnetworks/{subnetwork_id}`
	// +kcc:proto:field=google.cloud.metastore.v1beta.NetworkConfig.Consumer.subnetwork
	Subnetwork *string `json:"subnetwork,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.Restore
type Restore struct {
}

// +kcc:proto=google.cloud.metastore.v1beta.ScalingConfig
type ScalingConfig struct {
	// An enum of readable instance sizes, with each instance size mapping to a
	//  float value (e.g. InstanceSize.EXTRA_SMALL = scaling_factor(0.1))
	// +kcc:proto:field=google.cloud.metastore.v1beta.ScalingConfig.instance_size
	InstanceSize *string `json:"instanceSize,omitempty"`

	// Scaling factor, increments of 0.1 for values less than 1.0, and
	//  increments of 1.0 for values greater than 1.0.
	// +kcc:proto:field=google.cloud.metastore.v1beta.ScalingConfig.scaling_factor
	ScalingFactor *float32 `json:"scalingFactor,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.Secret
type Secret struct {
	// The relative resource name of a Secret Manager secret version, in the
	//  following form:
	//
	//  `projects/{project_number}/secrets/{secret_id}/versions/{version_id}`.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Secret.cloud_secret
	CloudSecret *string `json:"cloudSecret,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.Service
type Service struct {
	// Configuration information specific to running Hive metastore
	//  software as the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.hive_metastore_config
	HiveMetastoreConfig *HiveMetastoreConfig `json:"hiveMetastoreConfig,omitempty"`

	// Immutable. The relative resource name of the metastore service, in the
	//  following format:
	//
	//  `projects/{project_number}/locations/{location_id}/services/{service_id}`.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.name
	Name *string `json:"name,omitempty"`

	// User-defined labels for the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The relative resource name of the VPC network on which the
	//  instance can be accessed. It is specified in the following form:
	//
	//  `projects/{project_number}/global/networks/{network_id}`.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.network
	Network *string `json:"network,omitempty"`

	// The TCP port at which the metastore service is reached. Default: 9083.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.port
	Port *int32 `json:"port,omitempty"`

	// The tier of the service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.tier
	Tier *string `json:"tier,omitempty"`

	// The setting that defines how metastore metadata should be integrated with
	//  external services and systems.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.metadata_integration
	MetadataIntegration *MetadataIntegration `json:"metadataIntegration,omitempty"`

	// The one hour maintenance window of the metastore service. This specifies
	//  when the service can be restarted for maintenance purposes in UTC time.
	//  Maintenance window is not needed for services with the SPANNER
	//  database type.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.maintenance_window
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`

	// Immutable. The release channel of the service.
	//  If unspecified, defaults to `STABLE`.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.release_channel
	ReleaseChannel *string `json:"releaseChannel,omitempty"`

	// Immutable. Information used to configure the Dataproc Metastore service to
	//  encrypt customer data at rest. Cannot be updated.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// The configuration specifying the network settings for the
	//  Dataproc Metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.network_config
	NetworkConfig *NetworkConfig `json:"networkConfig,omitempty"`

	// Immutable. The database type that the Metastore service stores its data.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.database_type
	DatabaseType *string `json:"databaseType,omitempty"`

	// The configuration specifying telemetry settings for the Dataproc Metastore
	//  service. If unspecified defaults to `JSON`.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.telemetry_config
	TelemetryConfig *TelemetryConfig `json:"telemetryConfig,omitempty"`

	// Scaling configuration of the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.scaling_config
	ScalingConfig *ScalingConfig `json:"scalingConfig,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.TelemetryConfig
type TelemetryConfig struct {
	// The output format of the Dataproc Metastore service's logs.
	// +kcc:proto:field=google.cloud.metastore.v1beta.TelemetryConfig.log_format
	LogFormat *string `json:"logFormat,omitempty"`
}

// +kcc:proto=google.protobuf.Int32Value
type Int32Value struct {
	// The int32 value.
	// +kcc:proto:field=google.protobuf.Int32Value.value
	Value *int32 `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.Backup
type BackupObservedState struct {
	// Output only. The time when the backup was started.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Backup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the backup finished creating.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Backup.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The current state of the backup.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. The revision of the service at the time of backup.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Backup.service_revision
	ServiceRevision *Service `json:"serviceRevision,omitempty"`

	// Output only. Services that are restoring from the backup.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Backup.restoring_services
	RestoringServices []string `json:"restoringServices,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.MetadataExport
type MetadataExportObservedState struct {
	// Output only. A Cloud Storage URI of a folder that metadata are exported
	//  to, in the form of
	//  `gs://<bucket_name>/<path_inside_bucket>/<export_folder>`, where
	//  `<export_folder>` is automatically generated.
	// +kcc:proto:field=google.cloud.metastore.v1beta.MetadataExport.destination_gcs_uri
	DestinationGcsURI *string `json:"destinationGcsURI,omitempty"`

	// Output only. The time when the export started.
	// +kcc:proto:field=google.cloud.metastore.v1beta.MetadataExport.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time when the export ended.
	// +kcc:proto:field=google.cloud.metastore.v1beta.MetadataExport.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The current state of the export.
	// +kcc:proto:field=google.cloud.metastore.v1beta.MetadataExport.state
	State *string `json:"state,omitempty"`

	// Output only. The type of the database dump.
	// +kcc:proto:field=google.cloud.metastore.v1beta.MetadataExport.database_dump_type
	DatabaseDumpType *string `json:"databaseDumpType,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.MetadataManagementActivity
type MetadataManagementActivityObservedState struct {
	// Output only. The latest metadata exports of the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.MetadataManagementActivity.metadata_exports
	MetadataExports []MetadataExport `json:"metadataExports,omitempty"`

	// Output only. The latest restores of the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.MetadataManagementActivity.restores
	Restores []Restore `json:"restores,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.NetworkConfig
type NetworkConfigObservedState struct {
	// Immutable. The consumer-side network configuration for the Dataproc
	//  Metastore instance.
	// +kcc:proto:field=google.cloud.metastore.v1beta.NetworkConfig.consumers
	Consumers []NetworkConfig_ConsumerObservedState `json:"consumers,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.NetworkConfig.Consumer
type NetworkConfig_ConsumerObservedState struct {
	// Output only. The URI of the endpoint used to access the metastore
	//  service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.NetworkConfig.Consumer.endpoint_uri
	EndpointURI *string `json:"endpointURI,omitempty"`

	// Output only. The location of the endpoint URI. Format:
	//  `projects/{project}/locations/{location}`.
	// +kcc:proto:field=google.cloud.metastore.v1beta.NetworkConfig.Consumer.endpoint_location
	EndpointLocation *string `json:"endpointLocation,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.Restore
type RestoreObservedState struct {
	// Output only. The time when the restore started.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Restore.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time when the restore ended.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Restore.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The current state of the restore.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Restore.state
	State *string `json:"state,omitempty"`

	// Output only. The relative resource name of the metastore service backup to
	//  restore from, in the following form:
	//
	//  `projects/{project_id}/locations/{location_id}/services/{service_id}/backups/{backup_id}`.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Restore.backup
	Backup *string `json:"backup,omitempty"`

	// Output only. The type of restore.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Restore.type
	Type *string `json:"type,omitempty"`

	// Output only. The restore details containing the revision of the service to
	//  be restored to, in format of JSON.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Restore.details
	Details *string `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1beta.Service
type ServiceObservedState struct {
	// Output only. The time when the metastore service was created.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the metastore service was last updated.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The URI of the endpoint used to access the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.endpoint_uri
	EndpointURI *string `json:"endpointURI,omitempty"`

	// Output only. The current state of the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the current state of the
	//  metastore service, if available.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Output only. A Cloud Storage URI (starting with `gs://`) that specifies
	//  where artifacts related to the metastore service are stored.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.artifact_gcs_uri
	ArtifactGcsURI *string `json:"artifactGcsURI,omitempty"`

	// Output only. The globally unique resource identifier of the metastore
	//  service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The metadata management activities of the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.metadata_management_activity
	MetadataManagementActivity *MetadataManagementActivity `json:"metadataManagementActivity,omitempty"`

	// The configuration specifying the network settings for the
	//  Dataproc Metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1beta.Service.network_config
	NetworkConfig *NetworkConfigObservedState `json:"networkConfig,omitempty"`
}
