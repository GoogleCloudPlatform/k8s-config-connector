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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetAppBackupVaultGVK = GroupVersion.WithKind("NetAppBackupVault")

// BackupVaultSpec defines the desired state of NetAppBackupVault
// +kcc:proto=google.cloud.netapp.v1.BackupVault
type BackupVaultSpec struct {
	// The NetAppBackupVault name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// +required
	Location string `json:"location"`

	// Description of the backup vault.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupVault.description
	Description *string `json:"description,omitempty"`

	// NOT YET
	// // Resource labels to represent user provided metadata.
	// // +kcc:proto:field=google.cloud.netapp.v1.BackupVault.labels
	// Labels map[string]string `json:"labels,omitempty"`
}

// BackupVaultStatus defines the config connector machine state of NetAppBackupVault
type BackupVaultStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetAppBackupVault resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BackupVaultObservedState `json:"observedState,omitempty"`
}

// NetAppBackupVaultObservedState is the state of the NetAppBackupVault resource as most recently observed in GCP.
// +kcc:proto=google.cloud.netapp.v1.BackupVault// +kcc:proto=google.cloud.netapp.v1.BackupVault
type BackupVaultObservedState struct {
	// Output only. The backup vault state.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupVault.state
	State *string `json:"state,omitempty"`

	// Output only. Create time of the backup vault.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupVault.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetappbackupvault;gcpnetappbackupvaults
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetAppBackupVault is the Schema for the NetAppBackupVault API
// +k8s:openapi-gen=true
type NetAppBackupVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BackupVaultSpec   `json:"spec,omitempty"`
	Status BackupVaultStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetAppBackupVaultList contains a list of NetAppBackupVault
type NetAppBackupVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetAppBackupVault `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetAppBackupVault{}, &NetAppBackupVaultList{})
}
