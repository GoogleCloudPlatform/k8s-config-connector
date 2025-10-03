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
	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigtableBackupGVK = GroupVersion.WithKind("BigtableBackup")

type BigtableBackupParent struct {
	// + required
	ClusterRef ClusterRef `json:"clusterRef"`
}

// BigtableBackupSpec defines the desired state of BigtableBackup
// +kcc:spec:proto=google.bigtable.admin.v2.Backup
type BigtableBackupSpec struct {
	BigtableBackupParent `json:",inline"`

	// The BigtableBackup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Immutable. Name of the table from which this backup was created.
	//  This needs to be in the same instance as the backup. Values are of the form
	//  `projects/{project}/instances/{instance}/tables/{source_table}`.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.source_table
	SourceTableRef *bigtablev1beta1.TableRef `json:"sourceTableRef,omitempty"`

	// Required. The expiration time of the backup.
	//  When creating a backup or updating its `expire_time`, the value must be
	//  greater than the backup creation time by:
	//  - At least 6 hours
	//  - At most 90 days
	//
	//  Once the `expire_time` has passed, Cloud Bigtable will delete the backup.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Indicates the backup type of the backup.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.backup_type
	BackupType *string `json:"backupType,omitempty"`

	// The time at which the hot backup will be converted to a standard backup.
	//  Once the `hot_to_standard_time` has passed, Cloud Bigtable will convert the
	//  hot backup to a standard backup. This value must be greater than the backup
	//  creation time by:
	//  - At least 24 hours
	//
	//  This field only applies for hot backups. When creating or updating a
	//  standard backup, attempting to set this field will fail the request.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.hot_to_standard_time
	HotToStandardTime *string `json:"hotToStandardTime,omitempty"`
}

// BigtableBackupStatus defines the config connector machine state of BigtableBackup
type BigtableBackupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigtableBackup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigtableBackupObservedState `json:"observedState,omitempty"`
}

// BigtableBackupObservedState is the state of the BigtableBackup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.bigtable.admin.v2.Backup
type BigtableBackupObservedState struct {
	// Output only. Name of the backup from which this backup was copied. If a
	//  backup is not created by copying a backup, this field will be empty. Values
	//  are of the form:
	//  projects/<project>/instances/<instance>/clusters/<cluster>/backups/<backup>
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.source_backup
	SourceBackup *string `json:"sourceBackup,omitempty"`

	// Output only. `start_time` is the time that the backup was started
	//  (i.e. approximately the time the
	//  [CreateBackup][google.bigtable.admin.v2.BigtableTableAdmin.CreateBackup]
	//  request is received).  The row data in this backup will be no older than
	//  this timestamp.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. `end_time` is the time that the backup was finished. The row
	//  data in the backup will be no newer than this timestamp.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Size of the backup in bytes.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.size_bytes
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// Output only. The current state of the backup.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. The encryption information for the backup.
	// +kcc:proto:field=google.bigtable.admin.v2.Backup.encryption_info
	EncryptionInfo *EncryptionInfoObservedState `json:"encryptionInfo,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtablebackup;gcpbigtablebackups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigtableBackup is the Schema for the BigtableBackup API
// +k8s:openapi-gen=true
type BigtableBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigtableBackupSpec   `json:"spec,omitempty"`
	Status BigtableBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigtableBackupList contains a list of BigtableBackup
type BigtableBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigtableBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigtableBackup{}, &BigtableBackupList{})
}
