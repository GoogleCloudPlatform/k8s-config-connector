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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BackupDRBackupGVK = GroupVersion.WithKind("BackupDRBackup")

// +kcc:proto=google.cloud.backupdr.v1.Scheduling
type Scheduling struct {
	// Optional. Defines the maintenance behavior for this instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.on_host_maintenance
	OnHostMaintenance *string `json:"onHostMaintenance,omitempty"`

	// Optional. Specifies whether the instance should be automatically restarted
	//  if it is terminated by Compute Engine (not terminated by a user).
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.automatic_restart
	AutomaticRestart *bool `json:"automaticRestart,omitempty"`

	// Optional. Defines whether the instance is preemptible.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.preemptible
	Preemptible *bool `json:"preemptible,omitempty"`

	// Optional. A set of node affinity and anti-affinity configurations.
	//  Overrides reservationAffinity.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.node_affinities
	NodeAffinities []Scheduling_NodeAffinity `json:"nodeAffinities,omitempty"`

	// Optional. The minimum number of virtual CPUs this instance will consume
	//  when running on a sole-tenant node.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.min_node_cpus
	MinNodeCPUs *int32 `json:"minNodeCPUs,omitempty"`

	// Optional. Specifies the provisioning model of the instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.provisioning_model
	ProvisioningModel *string `json:"provisioningModel,omitempty"`

	// Optional. Specifies the termination action for the instance.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.instance_termination_action
	InstanceTerminationAction *string `json:"instanceTerminationAction,omitempty"`

	// Optional. Specifies the maximum amount of time a Local Ssd Vm should wait for
	//   recovery on the host.
	// +kcc:proto:field=google.cloud.backupdr.v1.Scheduling.local_ssd_recovery_timeout
	LocalSSDRecoveryTimeout *SchedulingDuration `json:"localSSDRecoveryTimeout,omitempty"`
}

// BackupDRBackupSpec defines the desired state of BackupDRBackup
// +kcc:spec:proto=google.cloud.backupdr.v1.Backup
type BackupDRBackupSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of the resource.
	// +required
	Location *string `json:"location"`

	// Immutable. The BackupVault of the resource.
	// +required
	BackupVault *string `json:"backupVault"`

	// Immutable. The DataSource of the resource.
	// +required
	DataSource *string `json:"dataSource"`

	// The BackupDRBackup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Resource labels to represent user provided metadata.
	//  No labels currently defined.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The backup can not be deleted before this time.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.enforced_retention_end_time
	EnforcedRetentionEndTime *string `json:"enforcedRetentionEndTime,omitempty"`

	// Optional. When this backup is automatically expired.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Optional. The list of BackupLocks taken by the accessor Backup Appliance.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.backup_appliance_locks
	BackupApplianceLocks []BackupLock `json:"backupApplianceLocks,omitempty"`
}

// BackupDRBackupStatus defines the config connector machine state of BackupDRBackup
type BackupDRBackupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BackupDRBackup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BackupDRBackupObservedState `json:"observedState,omitempty"`
}

// BackupDRBackupObservedState is the state of the BackupDRBackup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.backupdr.v1.Backup
type BackupDRBackupObservedState struct {
	// Output only. The description of the Backup instance (2048 characters or
	//  less).
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.description
	Description *string `json:"description,omitempty"`

	// Output only. The time when the instance was created.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the instance was updated.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The point in time when this backup was captured from the
	//  source.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.consistency_time
	ConsistencyTime *string `json:"consistencyTime,omitempty"`

	// Output only. Server specified ETag to prevent updates from overwriting each
	//  other.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The Backup resource instance state.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. The list of BackupLocks taken by the service to prevent the
	//  deletion of the backup.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.service_locks
	ServiceLocks []BackupLockObservedState `json:"serviceLocks,omitempty"`

	// Output only. Compute Engine specific backup properties.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.compute_instance_backup_properties
	ComputeInstanceBackupProperties *ComputeInstanceBackupPropertiesObservedState `json:"computeInstanceBackupProperties,omitempty"`

	// Output only. Cloud SQL specific backup properties.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.cloud_sql_instance_backup_properties
	CloudSQLInstanceBackupProperties *CloudSQLInstanceBackupPropertiesObservedState `json:"cloudSQLInstanceBackupProperties,omitempty"`

	// Output only. Backup Appliance specific backup properties.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.backup_appliance_backup_properties
	BackupApplianceBackupProperties *BackupApplianceBackupPropertiesObservedState `json:"backupApplianceBackupProperties,omitempty"`

	// Output only. Disk specific backup properties.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.disk_backup_properties
	DiskBackupProperties *DiskBackupProperties `json:"diskBackupProperties,omitempty"`

	// Output only. Type of the backup, unspecified, scheduled or ondemand.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.backup_type
	BackupType *string `json:"backupType,omitempty"`

	// Output only. Configuration for a Google Cloud resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.gcp_backup_plan_info
	GcpBackupPlanInfo *Backup_GcpBackupPlanInfo `json:"gcpBackupPlanInfo,omitempty"`

	// Output only. source resource size in bytes at the time of the backup.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.resource_size_bytes
	ResourceSizeBytes *int64 `json:"resourceSizeBytes,omitempty"`

	// Optional. Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Optional. Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbackupdrbackup;gcpbackupdrbackups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BackupDRBackup is the Schema for the BackupDRBackup API
// +k8s:openapi-gen=true
type BackupDRBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupDRBackupSpec   `json:"spec,omitempty"`
	Status BackupDRBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BackupDRBackupList contains a list of BackupDRBackup
type BackupDRBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupDRBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupDRBackup{}, &BackupDRBackupList{})
}
