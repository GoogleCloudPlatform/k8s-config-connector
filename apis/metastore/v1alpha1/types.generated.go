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

// +generated:types
// krm.group: metastore.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.metastore.v1
// resource: MetastoreBackup:Backup
// resource: MetastoreService:Service
// resource: MetastoreFederation:Federation

package v1alpha1

// +kcc:proto=google.cloud.metastore.v1.KerberosConfig
type KerberosConfig struct {
	// A Kerberos keytab file that can be used to authenticate a service principal
	//  with a Kerberos Key Distribution Center (KDC).
	// +kcc:proto:field=google.cloud.metastore.v1.KerberosConfig.keytab
	Keytab *Secret `json:"keytab,omitempty"`

	// A Kerberos principal that exists in the both the keytab the KDC
	//  to authenticate as. A typical principal is of the form
	//  `primary/instance@REALM`, but there is no exact format.
	// +kcc:proto:field=google.cloud.metastore.v1.KerberosConfig.principal
	Principal *string `json:"principal,omitempty"`

	// A Cloud Storage URI that specifies the path to a
	//  krb5.conf file. It is of the form `gs://{bucket_name}/path/to/krb5.conf`,
	//  although the file does not need to be named krb5.conf explicitly.
	// +kcc:proto:field=google.cloud.metastore.v1.KerberosConfig.krb5_config_gcs_uri
	Krb5ConfigGCSURI *string `json:"krb5ConfigGCSURI,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1.MaintenanceWindow
type MaintenanceWindow struct {
	// The hour of day (0-23) when the window starts.
	// +kcc:proto:field=google.cloud.metastore.v1.MaintenanceWindow.hour_of_day
	HourOfDay *Int32Value `json:"hourOfDay,omitempty"`

	// The day of week, when the window starts.
	// +kcc:proto:field=google.cloud.metastore.v1.MaintenanceWindow.day_of_week
	DayOfWeek *string `json:"dayOfWeek,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1.MetadataExport
type MetadataExport struct {
}

// +kcc:proto=google.cloud.metastore.v1.MetadataManagementActivity
type MetadataManagementActivity struct {
}

// +kcc:proto=google.cloud.metastore.v1.NetworkConfig
type NetworkConfig struct {
	// Immutable. The consumer-side network configuration for the Dataproc
	//  Metastore instance.
	// +kcc:proto:field=google.cloud.metastore.v1.NetworkConfig.consumers
	Consumers []NetworkConfig_Consumer `json:"consumers,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1.Restore
type Restore struct {
}

// +kcc:proto=google.cloud.metastore.v1.ScalingConfig
type ScalingConfig struct {
	// An enum of readable instance sizes, with each instance size mapping to a
	//  float value (e.g. InstanceSize.EXTRA_SMALL = scaling_factor(0.1))
	// +kcc:proto:field=google.cloud.metastore.v1.ScalingConfig.instance_size
	InstanceSize *string `json:"instanceSize,omitempty"`

	// Scaling factor, increments of 0.1 for values less than 1.0, and
	//  increments of 1.0 for values greater than 1.0.
	// +kcc:proto:field=google.cloud.metastore.v1.ScalingConfig.scaling_factor
	ScalingFactor *float32 `json:"scalingFactor,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1.TelemetryConfig
type TelemetryConfig struct {
	// The output format of the Dataproc Metastore service's logs.
	// +kcc:proto:field=google.cloud.metastore.v1.TelemetryConfig.log_format
	LogFormat *string `json:"logFormat,omitempty"`
}

// +kcc:proto=google.protobuf.Int32Value
type Int32Value struct {
	// The int32 value.
	// +kcc:proto:field=google.protobuf.Int32Value.value
	Value *int32 `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1.MetadataExport
type MetadataExportObservedState struct {
	// Output only. A Cloud Storage URI of a folder that metadata are exported
	//  to, in the form of
	//  `gs://<bucket_name>/<path_inside_bucket>/<export_folder>`, where
	//  `<export_folder>` is automatically generated.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataExport.destination_gcs_uri
	DestinationGCSURI *string `json:"destinationGCSURI,omitempty"`

	// Output only. The time when the export started.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataExport.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time when the export ended.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataExport.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The current state of the export.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataExport.state
	State *string `json:"state,omitempty"`

	// Output only. The type of the database dump.
	// +kcc:proto:field=google.cloud.metastore.v1.MetadataExport.database_dump_type
	DatabaseDumpType *string `json:"databaseDumpType,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1.NetworkConfig
type NetworkConfigObservedState struct {
	// Immutable. The consumer-side network configuration for the Dataproc
	//  Metastore instance.
	// +kcc:proto:field=google.cloud.metastore.v1.NetworkConfig.consumers
	Consumers []NetworkConfig_ConsumerObservedState `json:"consumers,omitempty"`
}

// +kcc:proto=google.cloud.metastore.v1.Restore
type RestoreObservedState struct {
	// Output only. The time when the restore started.
	// +kcc:proto:field=google.cloud.metastore.v1.Restore.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time when the restore ended.
	// +kcc:proto:field=google.cloud.metastore.v1.Restore.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The current state of the restore.
	// +kcc:proto:field=google.cloud.metastore.v1.Restore.state
	State *string `json:"state,omitempty"`

	// Output only. The relative resource name of the metastore service backup to
	//  restore from, in the following form:
	//
	//  `projects/{project_id}/locations/{location_id}/services/{service_id}/backups/{backup_id}`.
	// +kcc:proto:field=google.cloud.metastore.v1.Restore.backup
	Backup *string `json:"backup,omitempty"`

	// Output only. The type of restore.
	// +kcc:proto:field=google.cloud.metastore.v1.Restore.type
	Type *string `json:"type,omitempty"`

	// Output only. The restore details containing the revision of the service to
	//  be restored to, in format of JSON.
	// +kcc:proto:field=google.cloud.metastore.v1.Restore.details
	Details *string `json:"details,omitempty"`
}
