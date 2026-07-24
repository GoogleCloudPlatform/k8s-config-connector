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

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BackupDRBackupGVK = GroupVersion.WithKind("BackupDRBackup")

// BackupDRBackupSpec defines the desired state of BackupDRBackup
// +kcc:spec:proto=google.cloud.backupdr.v1.Backup
type BackupDRBackupSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	// +required
	Location *string `json:"location"`

	// Immutable. The BackupDRBackupVault that this resource belongs to.
	// +required
	BackupVaultRef *BackupVaultRef `json:"backupVaultRef"`

	// Immutable. The DataSource ID of this resource.
	// +required
	DataSourceID string `json:"dataSourceID"`

	// The BackupDRBackup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
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

	// Output only. The Backup resource instance state.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. Type of the backup, unspecified, scheduled or ondemand.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.backup_type
	BackupType *string `json:"backupType,omitempty"`

	// Output only. source resource size in bytes at the time of the backup.
	// +kcc:proto:field=google.cloud.backupdr.v1.Backup.resource_size_bytes
	ResourceSizeBytes *int64 `json:"resourceSizeBytes,omitempty"`
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

	// +required
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
