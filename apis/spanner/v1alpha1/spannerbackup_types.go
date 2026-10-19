// Copyright 2026 Google LLC
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

// API sources for SpannerBackup, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/spanner/admin/database/v1/backup.proto
// +kcc:source:service-docs=https://cloud.google.com/spanner/
// +kcc:source:resource-docs=https://docs.cloud.google.com/spanner/docs/reference/rest/v1/projects.instances.backups

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SpannerBackupGVK = GroupVersion.WithKind("SpannerBackup")

// SpannerBackupSpec defines the desired state of SpannerBackup
// +kcc:spec:proto=google.spanner.admin.database.v1.Backup
// +kcc:required-from-proto
type SpannerBackupSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The SpannerBackup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required for the
	//  [CreateBackup][google.spanner.admin.database.v1.DatabaseAdmin.CreateBackup]
	//  operation. Name of the database from which this backup was created. This
	//  needs to be in the same instance as the backup. Values are of the form
	//  `projects/<project>/instances/<instance>/databases/<database>`.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.database
	Database *string `json:"database,omitempty"`

	// The backup will contain an externally consistent copy of the database at
	//  the timestamp specified by `version_time`. If `version_time` is not
	//  specified, the system will set `version_time` to the `create_time` of the
	//  backup.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.version_time
	VersionTime *string `json:"versionTime,omitempty"`

	// Required for the
	//  [CreateBackup][google.spanner.admin.database.v1.DatabaseAdmin.CreateBackup]
	//  operation. The expiration time of the backup, with microseconds
	//  granularity that must be at least 6 hours and at most 366 days
	//  from the time the CreateBackup request is processed. Once the `expire_time`
	//  has passed, the backup is eligible to be automatically deleted by Cloud
	//  Spanner to free the resources used by the backup.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}

// SpannerBackupStatus defines the config connector machine state of SpannerBackup
type SpannerBackupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SpannerBackup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SpannerBackupObservedState `json:"observedState,omitempty"`
}

// SpannerBackupObservedState is the state of the SpannerBackup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.spanner.admin.database.v1.Backup
type SpannerBackupObservedState struct {
	// Output only. The time the
	//  [CreateBackup][google.spanner.admin.database.v1.DatabaseAdmin.CreateBackup]
	//  request is received. If the request does not specify `version_time`, the
	//  `version_time` of the backup will be equivalent to the `create_time`.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Size of the backup in bytes.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.size_bytes
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// Output only. The number of bytes that will be freed by deleting this
	//  backup. This value will be zero if, for example, this backup is part of an
	//  incremental backup chain and younger backups in the chain require that we
	//  keep its data. For backups not in an incremental backup chain, this is
	//  always the size of the backup. This value may change if backups on the same
	//  chain get created, deleted or expired.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.freeable_size_bytes
	FreeableSizeBytes *int64 `json:"freeableSizeBytes,omitempty"`

	// Output only. For a backup in an incremental backup chain, this is the
	//  storage space needed to keep the data that has changed since the previous
	//  backup. For all other backups, this is always the size of the backup. This
	//  value may change if backups on the same chain get deleted or expired.
	//
	//  This field can be used to calculate the total storage space used by a set
	//  of backups. For example, the total space used by all backups of a database
	//  can be computed by summing up this field.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.exclusive_size_bytes
	ExclusiveSizeBytes *int64 `json:"exclusiveSizeBytes,omitempty"`

	// Output only. The current state of the backup.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. The names of the restored databases that reference the backup.
	//  The database names are of
	//  the form `projects/<project>/instances/<instance>/databases/<database>`.
	//  Referencing databases may exist in different instances. The existence of
	//  any referencing database prevents the backup from being deleted. When a
	//  restored database from the backup enters the `READY` state, the reference
	//  to the backup is removed.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.referencing_databases
	ReferencingDatabases []string `json:"referencingDatabases,omitempty"`

	// Output only. The encryption information for the backup.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.encryption_info
	EncryptionInfo *EncryptionInfoObservedState `json:"encryptionInfo,omitempty"`

	// Output only. The encryption information for the backup, whether it is
	//  protected by one or more KMS keys. The information includes all Cloud
	//  KMS key versions used to encrypt the backup. The `encryption_status' field
	//  inside of each `EncryptionInfo` is not populated. At least one of the key
	//  versions must be available for the backup to be restored. If a key version
	//  is revoked in the middle of a restore, the restore behavior is undefined.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.encryption_information
	EncryptionInformation []EncryptionInfoObservedState `json:"encryptionInformation,omitempty"`

	// Output only. The database dialect information for the backup.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.database_dialect
	DatabaseDialect *string `json:"databaseDialect,omitempty"`

	// Output only. The names of the destination backups being created by copying
	//  this source backup. The backup names are of the form
	//  `projects/<project>/instances/<instance>/backups/<backup>`.
	//  Referencing backups may exist in different instances. The existence of
	//  any referencing backup prevents the backup from being deleted. When the
	//  copy operation is done (either successfully completed or cancelled or the
	//  destination backup is deleted), the reference to the backup is removed.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.referencing_backups
	ReferencingBackups []string `json:"referencingBackups,omitempty"`

	// Output only. The max allowed expiration time of the backup, with
	//  microseconds granularity. A backup's expiration time can be configured in
	//  multiple APIs: CreateBackup, UpdateBackup, CopyBackup. When updating or
	//  copying an existing backup, the expiration time specified must be
	//  less than `Backup.max_expire_time`.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.max_expire_time
	MaxExpireTime *string `json:"maxExpireTime,omitempty"`

	// Output only. List of backup schedule URIs that are associated with
	//  creating this backup. This is only applicable for scheduled backups, and
	//  is empty for on-demand backups.
	//
	//  To optimize for storage, whenever possible, multiple schedules are
	//  collapsed together to create one backup. In such cases, this field captures
	//  the list of all backup schedule URIs that are associated with creating
	//  this backup. If collapsing is not done, then this field captures the
	//  single backup schedule URI associated with creating this backup.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.backup_schedules
	BackupSchedules []string `json:"backupSchedules,omitempty"`

	// Output only. Populated only for backups in an incremental backup chain.
	//  Backups share the same chain id if and only if they belong to the same
	//  incremental backup chain. Use this field to determine which backups are
	//  part of the same incremental backup chain. The ordering of backups in the
	//  chain can be determined by ordering the backup `version_time`.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.incremental_backup_chain_id
	IncrementalBackupChainID *string `json:"incrementalBackupChainID,omitempty"`

	// Output only. Data deleted at a time older than this is guaranteed not to be
	//  retained in order to support this backup. For a backup in an incremental
	//  backup chain, this is the version time of the oldest backup that exists or
	//  ever existed in the chain. For all other backups, this is the version time
	//  of the backup. This field can be used to understand what data is being
	//  retained by the backup system.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.oldest_version_time
	OldestVersionTime *string `json:"oldestVersionTime,omitempty"`

	// Output only. The instance partition(s) storing the backup.
	//
	//  This is the same as the list of the instance partition(s) that the database
	//  had footprint in at the backup's `version_time`.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.instance_partitions
	InstancePartitions []BackupInstancePartition `json:"instancePartitions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpspannerbackup;gcpspannerbackups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SpannerBackup is the Schema for the SpannerBackup API
// +k8s:openapi-gen=true
type SpannerBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SpannerBackupSpec   `json:"spec,omitempty"`
	Status SpannerBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SpannerBackupList contains a list of SpannerBackup
type SpannerBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpannerBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpannerBackup{}, &SpannerBackupList{})
}
