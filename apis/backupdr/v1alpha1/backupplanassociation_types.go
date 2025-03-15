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

var BackupDRBackupPlanAssociationGVK = GroupVersion.WithKind("BackupDRBackupPlanAssociation")

// BackupDRBackupPlanAssociationSpec defines the desired state of BackupDRBackupPlanAssociation
// +kcc:proto=google.cloud.backupdr.v1.BackupPlanAssociation
type BackupDRBackupPlanAssociationSpec struct {
	// The BackupDRBackupPlanAssociation name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Required. Immutable. Resource type of workload on which backupplan is
	//  applied
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupPlanAssociation.resource_type
	// +required
	ResourceType *string `json:"resourceType,omitempty"`

	// Required. Immutable. Resource name of workload on which backupplan is
	//  applied
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupPlanAssociation.resource
	// +required
	Resource *string `json:"resource,omitempty"`

	// Required. The backup plan which needs to be applied on
	//  workload.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupPlanAssociation.backup_plan
	// +required
	BackupPlanRef *BackupPlanRef `json:"backupPlanRef,omitempty"`
}

// BackupDRBackupPlanAssociationStatus defines the config connector machine state of BackupDRBackupPlanAssociation
type BackupDRBackupPlanAssociationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BackupDRBackupPlanAssociation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BackupDRBackupPlanAssociationObservedState `json:"observedState,omitempty"`
}

// BackupDRBackupPlanAssociationObservedState is the state of the BackupDRBackupPlanAssociation resource as most recently observed in GCP.
// +kcc:proto=google.cloud.backupdr.v1.BackupPlanAssociation
type BackupDRBackupPlanAssociationObservedState struct {
	// Output only. Identifier. The resource name of BackupPlanAssociation in
	//  below format Format :
	//  projects/{project}/locations/{location}/backupPlanAssociations/{backupPlanAssociationId}
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupPlanAssociation.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. The time when the instance was created.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupPlanAssociation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the instance was updated.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupPlanAssociation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The BackupPlanAssociation resource state.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupPlanAssociation.state
	State *string `json:"state,omitempty"`

	// Output only. The config info related to backup rules.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupPlanAssociation.rules_config_info
	RulesConfigInfo []RuleConfigInfoObservedState `json:"rulesConfigInfo,omitempty"`

	// Output only. Resource name of data source which will be used as storage
	//  location for backups taken. Format :
	//  projects/{project}/locations/{location}/backupVaults/{backupvault}/dataSources/{datasource}
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupPlanAssociation.data_source
	DataSource *string `json:"dataSource,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbackupdrbackupplanassociation;gcpbackupdrbackupplanassociations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BackupDRBackupPlanAssociation is the Schema for the BackupDRBackupPlanAssociation API
// +k8s:openapi-gen=true
type BackupDRBackupPlanAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BackupDRBackupPlanAssociationSpec   `json:"spec,omitempty"`
	Status BackupDRBackupPlanAssociationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BackupDRBackupPlanAssociationList contains a list of BackupDRBackupPlanAssociation
type BackupDRBackupPlanAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupDRBackupPlanAssociation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupDRBackupPlanAssociation{}, &BackupDRBackupPlanAssociationList{})
}
