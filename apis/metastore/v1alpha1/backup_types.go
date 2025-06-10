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

var MetastoreBackupGVK = GroupVersion.WithKind("MetastoreBackup")

// Parent defines the parent resource of a MetastoreBackup.
type MetastoreBackupParent struct {
	// +required
	// The MetastoreService that the backup belongs to.
	ServiceRef ServiceRef `json:"serviceRef"`
}

// MetastoreBackupSpec defines the desired state of MetastoreBackup
// +kcc:spec:proto=google.cloud.metastore.v1.Backup
type MetastoreBackupSpec struct {
	MetastoreBackupParent `json:",inline"`
	// The MetastoreBackup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The description of the backup.
	// +kcc:proto:field=google.cloud.metastore.v1.Backup.description
	Description *string `json:"description,omitempty"`
}

// MetastoreBackupStatus defines the config connector machine state of MetastoreBackup
type MetastoreBackupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MetastoreBackup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MetastoreBackupObservedState `json:"observedState,omitempty"`
}

// MetastoreBackupObservedState is the state of the MetastoreBackup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.metastore.v1.Backup
type MetastoreBackupObservedState struct {
	// Output only. The time when the backup was started.
	// +kcc:proto:field=google.cloud.metastore.v1.Backup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the backup finished creating.
	// +kcc:proto:field=google.cloud.metastore.v1.Backup.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The current state of the backup.
	// +kcc:proto:field=google.cloud.metastore.v1.Backup.state
	State *string `json:"state,omitempty"`

	// Output only. The revision of the service at the time of backup.
	// +kcc:proto:field=google.cloud.metastore.v1.Backup.service_revision
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	ServiceRevision *MetastoreServiceSpec `json:"serviceRevision,omitempty"`

	// Output only. Services that are restoring from the backup.
	// +kcc:proto:field=google.cloud.metastore.v1.Backup.restoring_services
	RestoringServices []string `json:"restoringServices,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmetastorebackup;gcpmetastorebackups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MetastoreBackup is the Schema for the MetastoreBackup API
// +k8s:openapi-gen=true
type MetastoreBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MetastoreBackupSpec   `json:"spec,omitempty"`
	Status MetastoreBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MetastoreBackupList contains a list of MetastoreBackup
type MetastoreBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetastoreBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MetastoreBackup{}, &MetastoreBackupList{})
}
