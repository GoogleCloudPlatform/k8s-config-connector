// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AlloyDBClusterGVK = GroupVersion.WithKind("AlloyDBCluster")

// AlloyDBClusterSpec defines the desired state of AlloyDBCluster
// +kcc:proto=google.cloud.alloydb.v1beta.Cluster
type AlloyDBClusterSpec struct {
	// The AlloyDBCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The type of cluster. If not set, defaults to PRIMARY.
	// Default value: "PRIMARY" Possible values: ["PRIMARY", "SECONDARY"].
	ClusterType *string `json:"clusterType,omitempty"`

	// Policy to determine if the cluster should be deleted forcefully.
	// Deleting a cluster forcefully, deletes the cluster and all its associated
	// instances within the cluster.
	// Deleting a Secondary cluster with a secondary instance REQUIRES setting
	// deletion_policy = "FORCE" otherwise an error is returned. This is needed
	// as there is no support to delete just the secondary instance, and the only
	// way to delete secondary instance is to delete the associated secondary
	// cluster forcefully which also deletes the secondary instance.
	DeletionPolicy *string `json:"deletionPolicy,omitempty"`

	// Immutable. The location where the alloydb cluster should reside.
	// +required
	Location *string `json:"location,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// Immutable. The source when restoring from a backup. Conflicts
	// with 'restoreContinuousBackupSource', both can't be set together.
	RestoreBackupSource *BackupSource `json:"restoreBackupSource,omitempty"`

	// Immutable. The source when restoring via point in time
	// recovery (PITR). Conflicts with 'restoreBackupSource', both can't
	// be set together.
	RestoreContinuousBackupSource *RestoreContinuousBackupSource `json:"restoreContinuousBackupSource,omitempty"`

	// User-settable and human-readable display name for the Cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.display_name
	DisplayName *string `json:"displayName,omitempty"`

	/* NOTYET
	// Optional. The database engine major version. This is an optional field and
	//  it is populated at the Cluster creation time. If a database version is not
	//  supplied at cluster creation time, then a default database version will
	//  be used.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.database_version
	DatabaseVersion *string `json:"databaseVersion,omitempty"`
	*/

	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.network_config
	NetworkConfig *Cluster_NetworkConfig `json:"networkConfig,omitempty"`

	// Required. The resource link for the VPC network in which cluster resources
	//  are created and from which they are accessible via Private IP. The network
	//  must belong to the same project as the cluster. It is specified in the
	//  form: `projects/{project}/global/networks/{network_id}`. This is required
	//  to create a cluster. Deprecated, use network_config.network instead.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.network
	NetworkRef *refs.ComputeNetworkRef `json:"networkRef,omitempty"`

	/* NOTYET
	// For Resource freshness validation (https://google.aip.dev/154)
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.etag
	Etag *string `json:"etag,omitempty"`
	*/

	/* NOTYET
	// Annotations to allow client tools to store small amount of arbitrary data.
	//  This is distinct from labels.
	//  https://google.aip.dev/128
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
	*/

	// Input only. Initial user to setup during cluster creation. Required.
	//  If used in `RestoreCluster` this is ignored.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.initial_user
	InitialUser *UserPassword `json:"initialUser,omitempty"`

	// The automated backup policy for this cluster.
	//
	//  If no policy is provided then the default policy will be used. If backups
	//  are supported for the cluster, the default policy takes one backup a day,
	//  has a backup window of 1 hour, and retains backups for 14 days.
	//  For more information on the defaults, consult the
	//  documentation for the message type.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.automated_backup_policy
	AutomatedBackupPolicy *AutomatedBackupPolicy `json:"automatedBackupPolicy,omitempty"`

	/* NOTYET
	// SSL configuration for this AlloyDB cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.ssl_config
	SSLConfig *SSLConfig `json:"sslConfig,omitempty"`
	*/

	// Optional. The encryption config can be specified to encrypt the data disks
	//  and other persistent data resources of a cluster with a
	//  customer-managed encryption key (CMEK). When this field is not
	//  specified, the cluster will then use default encryption scheme to
	//  protect the user data.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// Optional. Continuous backup configuration for this cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.continuous_backup_config
	ContinuousBackupConfig *ContinuousBackupConfig `json:"continuousBackupConfig,omitempty"`

	// Cross Region replication config specific to SECONDARY cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.secondary_config
	SecondaryConfig *Cluster_SecondaryConfig `json:"secondaryConfig,omitempty"`

	/* NOTYET
	// Optional. The configuration for Private Service Connect (PSC) for the
	//  cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.psc_config
	PSCConfig *Cluster_PSCConfig `json:"pscConfig,omitempty"`
	*/

	// Optional. The maintenance update policy determines when to allow or deny
	//  updates.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.maintenance_update_policy
	MaintenanceUpdatePolicy *MaintenanceUpdatePolicy `json:"maintenanceUpdatePolicy,omitempty"`

	/* NOTYET
	// Optional. Configuration parameters related to the Gemini in Databases
	//  add-on.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.gemini_config
	GeminiConfig *GeminiClusterConfig `json:"geminiConfig,omitempty"`

	// Optional. Subscription type of the cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.subscription_type
	SubscriptionType *string `json:"subscriptionType,omitempty"`

	// Optional. Input only. Immutable. Tag keys/values directly bound to this
	//  resource. For example:
	//  ```
	//  "123/environment": "production",
	//  "123/costCenter": "marketing"
	//  ```
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.tags
	Tags map[string]string `json:"tags,omitempty"`
	*/
}

// +kcc:proto=google.cloud.alloydb.v1beta.EncryptionConfig
type EncryptionConfig struct {
	// The fully-qualified resource name of the KMS key.
	//  Each Cloud KMS key is regionalized and has the following format:
	//  projects/{{PROJECT}}/locations/{{REGION}}/keyRings/{{RING}}/cryptoKeys/{{KEY_NAME}}
	// +kcc:proto:field=google.cloud.alloydb.v1beta.EncryptionConfig.kms_key_name
	KMSKeyNameRef *kmsv1beta1.KMSCryptoKeyRef `json:"kmsKeyNameRef,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.WeeklySchedule
type AutomatedBackupPolicy_WeeklySchedule struct {
	// TODO: Remove "// +required" marker for StartTime after using direct controller.

	// The times during the day to start a backup. The start times are assumed
	//  to be in UTC and to be an exact hour (e.g., 04:00:00).
	//
	//  If no start times are provided, a single fixed start time is chosen
	//  arbitrarily.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.WeeklySchedule.start_times
	// +required
	StartTimes []TimeOfDay `json:"startTimes,omitempty"`

	// The days of the week to perform a backup.
	//
	//  If this field is left empty, the default of every day of the week is
	//  used.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.AutomatedBackupPolicy.WeeklySchedule.days_of_week
	DaysOfWeek []string `json:"daysOfWeek,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.UserPassword
type UserPassword struct {
	// The database username.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.UserPassword.user
	User *string `json:"user,omitempty"`

	// TODO: Verify if "// +required" marker is needed for Password after using direct controller.

	// The initial password for the user.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.UserPassword.password
	// +required
	Password *refsv1beta1secret.Legacy `json:"password,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.MaintenanceUpdatePolicy.MaintenanceWindow
type MaintenanceUpdatePolicy_MaintenanceWindow struct {
	// TODO: Verify if "// +required" marker is needed for Day after using direct controller.

	// Preferred day of the week for maintenance, e.g. MONDAY, TUESDAY, etc.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MaintenanceUpdatePolicy.MaintenanceWindow.day
	// +required
	Day *string `json:"day,omitempty"`

	// TODO: Verify if "// +required" marker is needed for StartTime after using direct controller.

	// Preferred time to start the maintenance operation on the specified day.
	//  Maintenance will start within 1 hour of this time.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MaintenanceUpdatePolicy.MaintenanceWindow.start_time
	// +required
	StartTime *TimeOfDay `json:"startTime,omitempty"`
}

// +kcc:proto=google.type.TimeOfDay
type TimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose
	//  to allow the value "24:00:00" for scenarios like business closing time.
	// +kcc:proto:field=google.type.TimeOfDay.hours
	Hours *int32 `json:"hours,omitempty"`

	// Minutes of hour of day. Must be from 0 to 59.
	// +kcc:proto:field=google.type.TimeOfDay.minutes
	Minutes *int32 `json:"minutes,omitempty"`

	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may
	//  allow the value 60 if it allows leap-seconds.
	// +kcc:proto:field=google.type.TimeOfDay.seconds
	Seconds *int32 `json:"seconds,omitempty"`

	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	// +kcc:proto:field=google.type.TimeOfDay.nanos
	Nanos *int32 `json:"nanos,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Cluster.NetworkConfig
type Cluster_NetworkConfig struct {
	// Optional. The resource link for the VPC network in which cluster
	//  resources are created and from which they are accessible via Private IP.
	//  The network must belong to the same project as the cluster. It is
	//  specified in the form:
	//  `projects/{project_number}/global/networks/{network_id}`. This is
	//  required to create a cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.NetworkConfig.network
	NetworkRef *refs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. Name of the allocated IP range for the private IP AlloyDB
	//  cluster, for example: "google-managed-services-default". If set, the
	//  instance IPs for this cluster will be created in the allocated range. The
	//  range name must comply with RFC 1035. Specifically, the name must be 1-63
	//  characters long and match the regular expression
	//  `[a-z]([-a-z0-9]*[a-z0-9])?`.
	//  Field name is intended to be consistent with Cloud SQL.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.NetworkConfig.allocated_ip_range
	AllocatedIPRange *string `json:"allocatedIpRange,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.BackupSource
type BackupSource struct {

	// Required. The name of the backup resource with the format:
	//   * projects/{project}/locations/{region}/backups/{backup_id}
	// +kcc:proto:field=google.cloud.alloydb.v1beta.BackupSource.backup_name
	// +required
	BackupNameRef *refs.AlloyDBBackupRef `json:"backupNameRef,omitempty"`
}

type RestoreContinuousBackupSource struct {
	// (Required) The name of the source cluster that this cluster is restored from.
	// +required
	ClusterRef *refs.AlloyDBClusterRef `json:"clusterRef,omitempty"`

	// Immutable. The point in time that this cluster is restored to, in RFC 3339 format.
	// +required
	PointInTime *string `json:"pointInTime,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Cluster.SecondaryConfig
type Cluster_SecondaryConfig struct {
	// TODO: Verify if "// +required" marker is needed for PrimaryClusterNameRef after using direct controller.

	// The name of the primary cluster name with the format:
	//  * projects/{project}/locations/{region}/clusters/{cluster_id}
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.SecondaryConfig.primary_cluster_name
	// +required
	PrimaryClusterNameRef *refs.AlloyDBClusterRef `json:"primaryClusterNameRef,omitempty"`
}

// AlloyDBClusterStatus defines the config connector machine state of AlloyDBCluster
type AlloyDBClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/*
		// A unique specifier for the AlloyDBCluster resource in GCP.
		ExternalRef *string `json:"externalRef,omitempty"`
	*/

	// Output only. Cluster created from backup.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.backup_source
	BackupSource []*BackupSourceObservedState `json:"backupSource,omitempty"`

	// Output only. Continuous backup properties for this cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.continuous_backup_info
	ContinuousBackupInfo []*ContinuousBackupInfoObservedState `json:"continuousBackupInfo,omitempty"`

	// The database engine major version. This is an output-only
	// field and it's populated at the Cluster creation time. This field
	// cannot be changed after cluster creation.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.database_version
	DatabaseVersion *string `json:"databaseVersion,omitempty"`

	// Output only. The encryption information for the cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.encryption_info
	EncryptionInfo []*EncryptionInfoObservedState `json:"encryptionInfo,omitempty"`

	// Output only. Cluster created via DMS migration.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.migration_source
	MigrationSource []*MigrationSourceObservedState `json:"migrationSource,omitempty"`

	// Output only. The name of the cluster resource with the format:
	//   * projects/{project}/locations/{region}/clusters/{cluster_id}
	//  where the cluster ID segment should satisfy the regex expression
	//  `[a-z0-9-]+`. For more details see https://google.aip.dev/122.
	//  The prefix of the cluster resource name is the name of the parent resource:
	//   * projects/{project}/locations/{region}
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.name
	Name *string `json:"name,omitempty"`

	// Output only. The system-generated UID of the resource. The UID is assigned
	//  when the resource is created, and it is retained until it is deleted.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.uid
	Uid *string `json:"uid,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AlloyDBClusterObservedState `json:"observedState,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.BackupSource
type BackupSourceObservedState struct {
	// The name of the backup resource.
	BackupName *string `json:"backupName,omitempty"`

	/* NOTYET
	// Output only. The system-generated UID of the backup which was used to
	//  create this resource. The UID is generated when the backup is created, and
	//  it is retained until the backup is deleted.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.BackupSource.backup_uid
	BackupUid *string `json:"backupUid,omitempty"`
	*/
}

// +kcc:proto=google.cloud.alloydb.v1beta.ContinuousBackupInfo
type ContinuousBackupInfoObservedState struct {
	// Output only. The encryption information for the WALs and backups required
	//  for ContinuousBackup.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ContinuousBackupInfo.encryption_info
	EncryptionInfo []*EncryptionInfoObservedState `json:"encryptionInfo,omitempty"`

	// Output only. When ContinuousBackup was most recently enabled. Set to null
	//  if ContinuousBackup is not enabled.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ContinuousBackupInfo.enabled_time
	EnabledTime *string `json:"enabledTime,omitempty"`

	// Output only. Days of the week on which a continuous backup is taken. Output
	//  only field. Ignored if passed into the request.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ContinuousBackupInfo.schedule
	Schedule []string `json:"schedule,omitempty"`

	// Output only. The earliest restorable time that can be restored to. Output
	//  only field.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.ContinuousBackupInfo.earliest_restorable_time
	EarliestRestorableTime *string `json:"earliestRestorableTime,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.MigrationSource
type MigrationSourceObservedState struct {
	// Output only. The host and port of the on-premises instance in host:port
	//  format
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MigrationSource.host_port
	HostPort *string `json:"hostPort,omitempty"`

	// Output only. Place holder for the external source identifier(e.g DMS job
	//  name) that created the cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MigrationSource.reference_id
	ReferenceID *string `json:"referenceId,omitempty"`

	// Output only. Type of migration source.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.MigrationSource.source_type
	SourceType *string `json:"sourceType,omitempty"`
}

// AlloyDBClusterSpec defines the desired state of AlloyDBCluster
// +kcc:proto=google.cloud.alloydb.v1beta.Cluster
// AlloyDBClusterObservedState is the state of the AlloyDBCluster resource as most recently observed in GCP.
type AlloyDBClusterObservedState struct {

	// Output only. The type of the cluster. This is an output-only field and it's
	//  populated at the Cluster creation time or the Cluster promotion
	//  time. The cluster type is determined by which RPC was used to create
	//  the cluster (i.e. `CreateCluster` vs. `CreateSecondaryCluster`
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.cluster_type
	ClusterType *string `json:"clusterType,omitempty"`
}

/* NOTYET
// +kcc:proto=google.cloud.alloydb.v1beta.Cluster
type ClusterObservedState struct {

	// Output only. Cluster created from CloudSQL snapshot.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.cloudsql_backup_run_source
	CloudsqlBackupRunSource *CloudSQLBackupRunSource `json:"cloudsqlBackupRunSource,omitempty"`

	// Output only. Create time stamp
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Delete time stamp
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The current serving state of the cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.state
	State *string `json:"state,omitempty"`

	// Output only. Reconciling (https://google.aip.dev/128#reconciliation).
	//  Set to true if the current state of Cluster does not match the user's
	//  intended state, and the service is actively updating the resource to
	//  reconcile them. This can happen due to user-triggered updates or
	//  system actions like failover or maintenance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. Cross Region replication config specific to PRIMARY cluster.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.primary_config
	PrimaryConfig *Cluster_PrimaryConfig `json:"primaryConfig,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. The maintenance schedule for the cluster, generated for a
	//  specific rollout if a maintenance window is set.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.maintenance_schedule
	MaintenanceSchedule *MaintenanceSchedule `json:"maintenanceSchedule,omitempty"`

	// Optional. Configuration parameters related to the Gemini in Databases
	//  add-on.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.gemini_config
	GeminiConfig *GeminiClusterConfigObservedState `json:"geminiConfig,omitempty"`

	// Output only. Metadata for free trial clusters
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Cluster.trial_metadata
	TrialMetadata *Cluster_TrialMetadata `json:"trialMetadata,omitempty"`
}
*/

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpalloydbcluster;gcpalloydbclusters
// +kubebuilder:subresource:status// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AlloyDBCluster is the Schema for the AlloyDBCluster API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type AlloyDBCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AlloyDBClusterSpec   `json:"spec,omitempty"`
	Status AlloyDBClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AlloyDBClusterList contains a list of AlloyDBCluster
type AlloyDBClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlloyDBCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlloyDBCluster{}, &AlloyDBClusterList{})
}
