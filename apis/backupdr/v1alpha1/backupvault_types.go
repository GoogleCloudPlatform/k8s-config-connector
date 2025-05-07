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

var BackupDRBackupVaultGVK = GroupVersion.WithKind("BackupDRBackupVault")

// BackupDRBackupVaultSpec defines the desired state of BackupDRBackupVault
// +kcc:proto=google.cloud.backupdr.v1.BackupVault
type BackupDRBackupVaultSpec struct {
	// The BackupDRBackupVault name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Optional. The description of the BackupVault instance (2048 characters or
	//  less).
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.description
	Description *string `json:"description,omitempty"`

	// Optional. Resource labels to represent user provided metadata.
	//  No labels currently defined:
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The default and minimum enforced retention for each backup within
	//  the backup vault.  The enforced retention for each backup can be extended.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.backup_minimum_enforced_retention_duration
	// +required
	BackupMinimumEnforcedRetentionDuration *string `json:"backupMinimumEnforcedRetentionDuration,omitempty"`

	// Optional. Server specified ETag for the backup vault resource to
	//  prevent simultaneous updates from overwiting each other.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Time after which the BackupVault resource is locked.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.effective_time
	EffectiveTime *string `json:"effectiveTime,omitempty"`

	// Optional. User annotations. See https://google.aip.dev/128#annotations
	//  Stores small amounts of arbitrary data.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Note: This field is added for future use case and will not be
	//  supported in the current release.
	//
	//  Access restriction for the backup vault.
	//  Default value is WITHIN_ORGANIZATION if not provided during creation.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.access_restriction
	AccessRestriction *string `json:"accessRestriction,omitempty"`

	// Optional. If set to true, allows deletion of a backup vault even when it contains inactive data sources.
	// This overrides the default restriction that prevents deletion of backup vaults with any
	// data sources, even if those data sources are inactive.
	IgnoreInactiveDatasources *bool `json:"ignoreInactiveDatasources,omitempty"`
}

// BackupDRBackupVaultStatus defines the config connector machine state of BackupDRBackupVault
type BackupDRBackupVaultStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BackupDRBackupVault resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BackupDRBackupVaultObservedState `json:"observedState,omitempty"`
}

// BackupDRBackupVaultObservedState is the state of the BackupDRBackupVault resource as most recently observed in GCP.
// +kcc:proto=google.cloud.backupdr.v1.BackupVault
type BackupDRBackupVaultObservedState struct {
	// Output only. Identifier. Name of the backup vault to create. It must have
	//  the
	//  format`"projects/{project}/locations/{location}/backupVaults/{backupvault}"`.
	//  `{backupvault}` cannot be changed after creation. It must be between 3-63
	//  characters long and must be unique within the project and location.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. The time when the instance was created.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the instance was updated.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Set to true when there are no backups nested under this
	//  resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.deletable
	Deletable *bool `json:"deletable,omitempty"`

	// Output only. The BackupVault resource instance state.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.state
	State *string `json:"state,omitempty"`

	// Output only. The number of backups in this backup vault.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.backup_count
	BackupCount *int64 `json:"backupCount,omitempty"`

	// Output only. Service account used by the BackupVault Service for this
	//  BackupVault.  The user should grant this account permissions in their
	//  workload project to enable the service to run backups and restores there.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Output only. Total size of the storage used by all backup resources.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.total_stored_bytes
	TotalStoredBytes *int64 `json:"totalStoredBytes,omitempty"`

	// Output only. Immutable after resource creation until resource deletion.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.uid
	UID *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbackupdrbackupvault;gcpbackupdrbackupvaults
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BackupDRBackupVault is the Schema for the BackupDRBackupVault API
// +k8s:openapi-gen=true
type BackupDRBackupVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BackupDRBackupVaultSpec   `json:"spec,omitempty"`
	Status BackupDRBackupVaultStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BackupDRBackupVaultList contains a list of BackupDRBackupVault
type BackupDRBackupVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupDRBackupVault `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupDRBackupVault{}, &BackupDRBackupVaultList{})
}
