// Copyright 2025 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GKEBackupBackupGVK = GroupVersion.WithKind("GKEBackupBackup")

// GKEBackupBackupSpec defines the desired state of GKEBackupBackup
// +kcc:spec:proto=google.cloud.gkebackup.v1.Backup
type GKEBackupBackupSpec struct {
	// The GKEBackupBackup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The BackupPlan from which this Backup is created.
	// +required
	BackupPlanRef *BackupPlanRef `json:"backupPlanRef,omitempty"`

	// Optional. A set of custom labels supplied by user.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Minimum age for this Backup (in days). If this field is set to a
	//  non-zero value, the Backup will be "locked" against deletion (either manual
	//  or automatic deletion) for the number of days provided (measured from the
	//  creation time of the Backup).  MUST be an integer value between 0-90
	//  (inclusive).
	//
	//  Defaults to parent BackupPlan's
	//  [backup_delete_lock_days][google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_delete_lock_days]
	//  setting and may only be increased
	//  (either at creation time or in a subsequent update).
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.delete_lock_days
	DeleteLockDays *int32 `json:"deleteLockDays,omitempty"`

	// Optional. The age (in days) after which this Backup will be automatically
	//  deleted. Must be an integer value >= 0:
	//
	//  - If 0, no automatic deletion will occur for this Backup.
	//  - If not 0, this must be >=
	//  [delete_lock_days][google.cloud.gkebackup.v1.Backup.delete_lock_days] and
	//  <= 365.
	//
	//  Once a Backup is created, this value may only be increased.
	//
	//  Defaults to the parent BackupPlan's
	//  [backup_retain_days][google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_retain_days]
	//  value.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.retain_days
	RetainDays *int32 `json:"retainDays,omitempty"`

	// Optional. User specified descriptive string for this Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.description
	Description *string `json:"description,omitempty"`
}

// GKEBackupBackupStatus defines the config connector machine state of GKEBackupBackup
type GKEBackupBackupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GKEBackupBackup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GKEBackupBackupObservedState `json:"observedState,omitempty"`
}

// GKEBackupBackupObservedState is the state of the GKEBackupBackup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.gkebackup.v1.Backup
type GKEBackupBackupObservedState struct {
	// Output only. The fully qualified name of the Backup.
	//  `projects/*/locations/*/backupPlans/*/backups/*`
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. Server generated global unique identifier of
	//  [UUID4](https://en.wikipedia.org/wiki/Universally_unique_identifier)
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.uid
	UID *string `json:"uid,omitempty"`

	// Output only. The timestamp when this Backup resource was created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this Backup resource was last updated.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. This flag indicates whether this Backup resource was created
	//  manually by a user or via a schedule in the BackupPlan. A value of True
	//  means that the Backup was created manually.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.manual
	Manual *bool `json:"manual,omitempty"`

	// Output only. The time at which an existing delete lock will expire for this
	//  backup (calculated from create_time +
	//  [delete_lock_days][google.cloud.gkebackup.v1.Backup.delete_lock_days]).
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.delete_lock_expire_time
	DeleteLockExpireTime *string `json:"deleteLockExpireTime,omitempty"`

	// Output only. The time at which this Backup will be automatically deleted
	//  (calculated from create_time +
	//  [retain_days][google.cloud.gkebackup.v1.Backup.retain_days]).
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.retain_expire_time
	RetainExpireTime *string `json:"retainExpireTime,omitempty"`

	// Output only. The customer managed encryption key that was used to encrypt
	//  the Backup's artifacts.  Inherited from the parent BackupPlan's
	//  [encryption_key][google.cloud.gkebackup.v1.BackupPlan.BackupConfig.encryption_key]
	//  value.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.encryption_key
	EncryptionKey *Backup_EncryptionKeyObservedState `json:"encryptionKey,omitempty"`

	// Output only. If True, all namespaces were included in the Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.all_namespaces
	AllNamespaces *bool `json:"allNamespaces,omitempty"`

	// Output only. If set, the list of namespaces that were included in the
	//  Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.selected_namespaces
	SelectedNamespaces *Namespaces `json:"selectedNamespaces,omitempty"`

	// Output only. If set, the list of ProtectedApplications whose resources
	//  were included in the Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.selected_applications
	SelectedApplications *NamespacedNames `json:"selectedApplications,omitempty"`

	// Output only. Whether or not the Backup contains volume data.  Controlled by
	//  the parent BackupPlan's
	//  [include_volume_data][google.cloud.gkebackup.v1.BackupPlan.BackupConfig.include_volume_data]
	//  value.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.contains_volume_data
	ContainsVolumeData *bool `json:"containsVolumeData,omitempty"`

	// Output only. Whether or not the Backup contains Kubernetes Secrets.
	//  Controlled by the parent BackupPlan's
	//  [include_secrets][google.cloud.gkebackup.v1.BackupPlan.BackupConfig.include_secrets]
	//  value.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.contains_secrets
	ContainsSecrets *bool `json:"containsSecrets,omitempty"`

	// Output only. Information about the GKE cluster from which this Backup was
	//  created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.cluster_metadata
	ClusterMetadata *Backup_ClusterMetadataObservedState `json:"clusterMetadata,omitempty"`

	// Output only. Current state of the Backup
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. Human-readable description of why the backup is in the current
	//  `state`.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.state_reason
	StateReason *string `json:"stateReason,omitempty"`

	// Output only. Completion time of the Backup
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.complete_time
	CompleteTime *string `json:"completeTime,omitempty"`

	// Output only. The total number of Kubernetes resources included in the
	//  Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.resource_count
	ResourceCount *int32 `json:"resourceCount,omitempty"`

	// Output only. The total number of volume backups contained in the Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.volume_count
	VolumeCount *int32 `json:"volumeCount,omitempty"`

	// Output only. The total size of the Backup in bytes = config backup size +
	//  sum(volume backup sizes)
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.size_bytes
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// Output only. `etag` is used for optimistic concurrency control as a way to
	//  help prevent simultaneous updates of a backup from overwriting each other.
	//  It is strongly suggested that systems make use of the `etag` in the
	//  read-modify-write cycle to perform backup updates in order to avoid
	//  race conditions: An `etag` is returned in the response to `GetBackup`,
	//  and systems are expected to put that etag in the request to
	//  `UpdateBackup` or `DeleteBackup` to ensure that their change will be
	//  applied to the same version of the resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The total number of Kubernetes Pods contained in the Backup.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.pod_count
	PodCount *int32 `json:"podCount,omitempty"`

	// Output only. The size of the config backup in bytes.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.config_backup_size_bytes
	ConfigBackupSizeBytes *int64 `json:"configBackupSizeBytes,omitempty"`

	// Output only. If false, Backup will fail when Backup for GKE detects
	//  Kubernetes configuration that is non-standard or
	//  requires additional setup to restore.
	//
	//  Inherited from the parent BackupPlan's
	//  [permissive_mode][google.cloud.gkebackup.v1.BackupPlan.BackupConfig.permissive_mode]
	//  value.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.permissive_mode
	PermissiveMode *bool `json:"permissiveMode,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkebackupbackup;gcpgkebackupbackups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEBackupBackup is the Schema for the GKEBackupBackup API
// +k8s:openapi-gen=true
type GKEBackupBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GKEBackupBackupSpec   `json:"spec,omitempty"`
	Status GKEBackupBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GKEBackupBackupList contains a list of GKEBackupBackup
type GKEBackupBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEBackupBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEBackupBackup{}, &GKEBackupBackupList{})
}

// +kcc:observedstate:proto=google.cloud.gkebackup.v1.Backup.ClusterMetadata
type Backup_ClusterMetadataObservedState struct {
	// Output only. The source cluster from which this Backup was created.
	//  Valid formats:
	//
	//    - `projects/*/locations/*/clusters/*`
	//    - `projects/*/zones/*/clusters/*`
	//
	//  This is inherited from the parent BackupPlan's
	//  [cluster][google.cloud.gkebackup.v1.BackupPlan.cluster] field.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.ClusterMetadata.cluster
	Cluster *string `json:"cluster,omitempty"`

	// Output only. The Kubernetes server version of the source cluster.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.ClusterMetadata.k8s_version
	K8sVersion *string `json:"k8sVersion,omitempty"`

	// Output only. A list of the Backup for GKE CRD versions found in the
	//  cluster.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.ClusterMetadata.backup_crd_versions
	BackupCRDVersions map[string]string `json:"backupCRDVersions,omitempty"`

	// Output only. GKE version
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.ClusterMetadata.gke_version
	GKEVersion *string `json:"gkeVersion,omitempty"`

	// Output only. Anthos version
	// +kcc:proto:field=google.cloud.gkebackup.v1.Backup.ClusterMetadata.anthos_version
	AnthosVersion *string `json:"anthosVersion,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.gkebackup.v1.EncryptionKey
type Backup_EncryptionKeyObservedState struct {
	// Optional. Google Cloud KMS encryption key. Format:
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*`
	// +kcc:proto:field=google.cloud.gkebackup.v1.EncryptionKey.gcp_kms_encryption_key
	GCPKMSEncryptionKey *string `json:"gcpKMSEncryptionKey,omitempty"`
}
