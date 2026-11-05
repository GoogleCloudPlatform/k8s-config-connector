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
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The SQLInstance that this backup belongs to.
	// +kubebuilder:validation:Required
	InstanceRef *refsv1beta1.SQLInstanceRef `json:"instanceRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The description of this backup, only applicable to on-demand backups.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Encryption configuration specific to a backup.
	// +kubebuilder:validation:Optional
	DiskEncryptionConfiguration *DiskEncryptionConfiguration `json:"diskEncryptionConfiguration,omitempty"`

	// Specifies the kind of backup, PHYSICAL or DEFAULT_SNAPSHOT.
	// +kubebuilder:validation:Optional
	BackupKind *string `json:"backupKind,omitempty"`

	// The SQLAdminBackup name (ID). If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.sql.v1.DiskEncryptionConfiguration
type DiskEncryptionConfiguration struct {
	// KMS key used to encrypt the backup.
	KMSKeyRef *kmsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
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
	// The time the run was enqueued in UTC timezone in RFC 3339 format.
	EnqueuedTime *string `json:"enqueuedTime,omitempty"`

	// The identifier for this backup run. Unique only for a specific Cloud SQL instance.
	ID *int64 `json:"id,omitempty"`

	// The time the backup operation actually started in UTC timezone in RFC 3339 format.
	StartTime *string `json:"startTime,omitempty"`

	// The time the backup operation completed in UTC timezone in RFC 3339 format.
	EndTime *string `json:"endTime,omitempty"`

	// Information about why the backup operation failed. This is only present if the run has the FAILED status.
	Error *OperationError `json:"error,omitempty"`

	// The status of this run.
	Status *string `json:"status,omitempty"`

	// The type of this run; can be either "AUTOMATED" or "ON_DEMAND" or "FINAL".
	Type *string `json:"type,omitempty"`

	// The start time of the backup window during which this the backup was attempted in RFC 3339 format.
	WindowStartTime *string `json:"windowStartTime,omitempty"`

	// Encryption status specific to a backup.
	DiskEncryptionStatus *DiskEncryptionStatus `json:"diskEncryptionStatus,omitempty"`

	// Backup time zone to prevent restores to an instance with a different time zone. Now relevant only for SQL Server.
	TimeZone *string `json:"timeZone,omitempty"`
}

// +kcc:proto=google.cloud.sql.v1.DiskEncryptionStatus
type DiskEncryptionStatus struct {
	// KMS key version used to encrypt the Cloud SQL instance resource
	// +kcc:proto:field=google.cloud.sql.v1.DiskEncryptionStatus.kms_key_version_name
	KMSKeyVersionName *string `json:"kmsKeyVersionName,omitempty"`

	// This is always `sql#diskEncryptionStatus`.
	// +kcc:proto:field=google.cloud.sql.v1.DiskEncryptionStatus.kind
	Kind *string `json:"kind,omitempty"`
}

// +kcc:proto=google.cloud.sql.v1.OperationError
type OperationError struct {
	// This is always `sql#operationError`.
	// +kcc:proto:field=google.cloud.sql.v1.OperationError.kind
	Kind *string `json:"kind,omitempty"`

	// Identifies the specific error that occurred.
	// +kcc:proto:field=google.cloud.sql.v1.OperationError.code
	Code *string `json:"code,omitempty"`

	// Additional information about the error encountered.
	// +kcc:proto:field=google.cloud.sql.v1.OperationError.message
	Message *string `json:"message,omitempty"`
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
