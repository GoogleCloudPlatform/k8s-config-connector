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

var MemorystoreInstanceBackupGVK = GroupVersion.WithKind("MemorystoreInstanceBackup")

// MemorystoreInstanceBackupSpec defines the desired state of MemorystoreInstanceBackup
// +kcc:spec:proto=google.cloud.memorystore.v1.Backup
type MemorystoreInstanceBackupSpec struct {
	// The project that this resource belongs to.
	BackupCollection *string `json:"backupCollection,omitempty"`

	// The MemorystoreInstanceBackup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// MemorystoreInstanceBackupStatus defines the config connector machine state of MemorystoreInstanceBackup
type MemorystoreInstanceBackupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MemorystoreInstanceBackup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MemorystoreInstanceBackupObservedState `json:"observedState,omitempty"`
}

// MemorystoreInstanceBackupObservedState is the state of the MemorystoreInstanceBackup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.memorystore.v1.Backup
type MemorystoreInstanceBackupObservedState struct {
	// Output only. The time when the backup was created.
	// +kcc:proto:field=google.cloud.memorystore.v1.Backup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Instance resource path of this backup.
	// +kcc:proto:field=google.cloud.memorystore.v1.Backup.instance
	Instance *string `json:"instance,omitempty"`

	// Output only. Instance uid of this backup.
	// +kcc:proto:field=google.cloud.memorystore.v1.Backup.instance_uid
	InstanceUid *string `json:"instanceUid,omitempty"`

	// Output only. Total size of the backup in bytes.
	// +kcc:proto:field=google.cloud.memorystore.v1.Backup.total_size_bytes
	TotalSizeBytes *int64 `json:"totalSizeBytes,omitempty"`

	// Output only. The time when the backup will expire.
	// +kcc:proto:field=google.cloud.memorystore.v1.Backup.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. valkey-7.5/valkey-8.0, etc.
	// +kcc:proto:field=google.cloud.memorystore.v1.Backup.engine_version
	EngineVersion *string `json:"engineVersion,omitempty"`

	// Output only. List of backup files of the backup.
	// +kcc:proto:field=google.cloud.memorystore.v1.Backup.backup_files
	BackupFiles []BackupFileObservedState `json:"backupFiles,omitempty"`

	// Output only. Node type of the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1.Backup.node_type
	NodeType *string `json:"nodeType,omitempty"`

	// Output only. Number of replicas for the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1.Backup.replica_count
	ReplicaCount *int32 `json:"replicaCount,omitempty"`

	// Output only. Number of shards for the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1.Backup.shard_count
	ShardCount *int32 `json:"shardCount,omitempty"`

	// Output only. Type of the backup.
	// +kcc:proto:field=google.cloud.memorystore.v1.Backup.backup_type
	BackupType *string `json:"backupType,omitempty"`

	// Output only. State of the backup.
	// +kcc:proto:field=google.cloud.memorystore.v1.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. System assigned unique identifier of the backup.
	// +kcc:proto:field=google.cloud.memorystore.v1.Backup.uid
	Uid *string `json:"uid,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.memorystore.v1.BackupFile
type BackupFileObservedState struct {
	// Output only. e.g: <shard-id>.rdb
	// +kcc:proto:field=google.cloud.memorystore.v1.BackupFile.file_name
	FileName *string `json:"fileName,omitempty"`

	// Output only. Size of the backup file in bytes.
	// +kcc:proto:field=google.cloud.memorystore.v1.BackupFile.size_bytes
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// Output only. The time when the backup file was created.
	// +kcc:proto:field=google.cloud.memorystore.v1.BackupFile.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmemorystoreinstancebackup;gcpmemorystoreinstancebackups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MemorystoreInstanceBackup is the Schema for the MemorystoreInstanceBackup API
// +k8s:openapi-gen=true
type MemorystoreInstanceBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MemorystoreInstanceBackupSpec   `json:"spec,omitempty"`
	Status MemorystoreInstanceBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MemorystoreInstanceBackupList contains a list of MemorystoreInstanceBackup
type MemorystoreInstanceBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MemorystoreInstanceBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MemorystoreInstanceBackup{}, &MemorystoreInstanceBackupList{})
}
