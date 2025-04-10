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
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetAppBackupPolicyGVK = GroupVersion.WithKind("NetAppBackupPolicy")

// NetAppBackupPolicySpec defines the desired state of NetAppBackupPolicy
// +kcc:spec:proto=google.cloud.netapp.v1.BackupPolicy
type NetAppBackupPolicySpec struct {
	commonv1alpha1.CommonSpec `json:",inline"`

	// +required
	Location string `json:"location"`

	// Number of daily backups to keep. Note that the minimum daily backup limit
	//  is 2.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.daily_backup_limit
	DailyBackupLimit *int32 `json:"dailyBackupLimit,omitempty"`

	// Number of weekly backups to keep. Note that the sum of daily, weekly and
	//  monthly backups should be greater than 1.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.weekly_backup_limit
	WeeklyBackupLimit *int32 `json:"weeklyBackupLimit,omitempty"`

	// Number of monthly backups to keep. Note that the sum of daily, weekly and
	//  monthly backups should be greater than 1.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.monthly_backup_limit
	MonthlyBackupLimit *int32 `json:"monthlyBackupLimit,omitempty"`

	// Description of the backup policy.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.description
	Description *string `json:"description,omitempty"`

	// If enabled, make backups automatically according to the schedules.
	//  This will be applied to all volumes that have this policy attached and
	//  enforced on volume level. If not specified, default is true.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// NOT YET
	// // Resource labels to represent user provided metadata.
	// // +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.labels
	// Labels map[string]string `json:"labels,omitempty"`
}

// NetAppBackupPolicyStatus defines the config connector machine state of NetAppBackupPolicy
type NetAppBackupPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetAppBackupPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetAppBackupPolicyObservedState `json:"observedState,omitempty"`
}

// NetAppBackupPolicyObservedState is the state of the NetAppBackupPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.netapp.v1.BackupPolicy
type NetAppBackupPolicyObservedState struct {
	// Output only. The total number of volumes assigned by this backup policy.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.assigned_volume_count
	AssignedVolumeCount *int32 `json:"assignedVolumeCount,omitempty"`

	// Output only. The time when the backup policy was created.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The backup policy state.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetappbackuppolicy;gcpnetappbackuppolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetAppBackupPolicy is the Schema for the NetAppBackupPolicy API
// +k8s:openapi-gen=true
type NetAppBackupPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetAppBackupPolicySpec   `json:"spec,omitempty"`
	Status NetAppBackupPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetAppBackupPolicyList contains a list of NetAppBackupPolicy
type NetAppBackupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetAppBackupPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetAppBackupPolicy{}, &NetAppBackupPolicyList{})
}
