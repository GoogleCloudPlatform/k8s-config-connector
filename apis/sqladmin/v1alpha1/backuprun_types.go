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
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SQLAdminBackupGVK = GroupVersion.WithKind("SQLAdminBackup")

// SQLAdminBackupSpec defines the desired state of SQLAdminBackup
// +kcc:spec:proto=google.cloud.sql.v1.BackupRun
type SQLAdminBackupSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The SQLInstance that this backup belongs to.
	// +required
	InstanceRef *refsv1beta1.SQLInstanceRef `json:"instanceRef"`

	// Immutable. The location of this resource.
	// +optional
	Location *string `json:"location,omitempty"`

	// The SQLAdminBackup name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The description of this backup run, only applicable to on-demand backups.
	// +optional
	Description *string `json:"description,omitempty"`

	// Immutable. Specifies the kind of backup, PHYSICAL or DEFAULT_SNAPSHOT.
	// +optional
	BackupKind *string `json:"backupKind,omitempty"`
}

// SQLAdminBackupStatus defines the config connector machine state of SQLAdminBackup
type SQLAdminBackupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SQLAdminBackup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SQLAdminBackupObservedState `json:"observedState,omitempty"`
}

// SQLAdminBackupObservedState is the state of the SQLAdminBackup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.sql.v1.BackupRun
type SQLAdminBackupObservedState struct {
	// Output only. The identifier for this backup run. Unique only for a specific Cloud SQL instance.
	// +kcc:proto:field=google.cloud.sql.v1.BackupRun.id
	ID *int64 `json:"id,omitempty"`

	// Output only. The status of this run.
	// +kcc:proto:field=google.cloud.sql.v1.BackupRun.status
	Status *string `json:"status,omitempty"`

	// Output only. The time the run was enqueued in UTC timezone in RFC 3339 format.
	// +kcc:proto:field=google.cloud.sql.v1.BackupRun.enqueued_time
	EnqueuedTime *string `json:"enqueuedTime,omitempty"`

	// Output only. The time the backup operation actually started in UTC timezone in RFC 3339 format.
	// +kcc:proto:field=google.cloud.sql.v1.BackupRun.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time the backup operation completed in UTC timezone in RFC 3339 format.
	// +kcc:proto:field=google.cloud.sql.v1.BackupRun.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Information about why the backup operation failed. This is only present if the run has the FAILED status.
	// +kcc:proto:field=google.cloud.sql.v1.BackupRun.error
	Error *OperationError `json:"error,omitempty"`

	// Output only. The type of this run; can be either "AUTOMATED" or "ON_DEMAND" or "FINAL".
	// +kcc:proto:field=google.cloud.sql.v1.BackupRun.type
	Type *string `json:"type,omitempty"`

	// Output only. The start time of the backup window during which this the backup was attempted in RFC 3339 format.
	// +kcc:proto:field=google.cloud.sql.v1.BackupRun.window_start_time
	WindowStartTime *string `json:"windowStartTime,omitempty"`

	// Output only. The URI of this resource.
	// +kcc:proto:field=google.cloud.sql.v1.BackupRun.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. Encryption configuration specific to a backup.
	// +kcc:proto:field=google.cloud.sql.v1.BackupRun.disk_encryption_configuration
	DiskEncryptionConfiguration *DiskEncryptionConfiguration `json:"diskEncryptionConfiguration,omitempty"`

	// Output only. Encryption status specific to a backup.
	// +kcc:proto:field=google.cloud.sql.v1.BackupRun.disk_encryption_status
	DiskEncryptionStatus *DiskEncryptionStatus `json:"diskEncryptionStatus,omitempty"`

	// Output only. The maximum chargeable bytes for the backup.
	// +kcc:proto:field=google.cloud.sql.v1.BackupRun.max_chargeable_bytes
	MaxChargeableBytes *int64 `json:"maxChargeableBytes,omitempty"`

	// Output only. Backup time zone to prevent restores to an instance with a different time zone. Now relevant only for SQL Server.
	// +kcc:proto:field=google.cloud.sql.v1.BackupRun.time_zone
	TimeZone *string `json:"timeZone,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsqladminbackup;gcpsqladminbackups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SQLAdminBackup is the Schema for the SQLAdminBackup API
// +k8s:openapi-gen=true
type SQLAdminBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SQLAdminBackupSpec   `json:"spec,omitempty"`
	Status SQLAdminBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SQLAdminBackupList contains a list of SQLAdminBackup
type SQLAdminBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLAdminBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SQLAdminBackup{}, &SQLAdminBackupList{})
}

// +kcc:proto=google.cloud.sql.v1.DiskEncryptionConfiguration
type DiskEncryptionConfiguration struct {
	// Resource name of KMS key for disk encryption
	KmsKeyRef *kmsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// This is always `sql#diskEncryptionConfiguration`.
	// +kcc:proto:field=google.cloud.sql.v1.DiskEncryptionConfiguration.kind
	Kind *string `json:"kind,omitempty"`
}
