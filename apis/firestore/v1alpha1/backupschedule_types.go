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
	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FirestoreBackupScheduleGVK = GroupVersion.WithKind("FirestoreBackupSchedule")

// FirestoreBackupScheduleSpec defines the desired state of FirestoreBackupSchedule
// +kcc:spec:proto=google.firestore.admin.v1.BackupSchedule
type FirestoreBackupScheduleSpec struct {
	/* The database that this resource belongs to. */
	// +required
	DatabaseRef v1beta1.FirestoreDatabaseRef `json:"databaseRef"`

	// At what relative time in the future, compared to its creation time,
	//  the backup should be deleted, e.g. keep backups for 7 days.
	//
	//  The maximum supported retention period is 14 weeks.
	// +kcc:proto:field=google.firestore.admin.v1.BackupSchedule.retention
	Retention *string `json:"retention,omitempty"`

	// For a schedule that runs daily.
	// +kcc:proto:field=google.firestore.admin.v1.BackupSchedule.daily_recurrence
	DailyRecurrence *DailyRecurrence `json:"dailyRecurrence,omitempty"`

	// For a schedule that runs weekly on a specific day.
	// +kcc:proto:field=google.firestore.admin.v1.BackupSchedule.weekly_recurrence
	WeeklyRecurrence *WeeklyRecurrence `json:"weeklyRecurrence,omitempty"`
}

// FirestoreBackupScheduleStatus defines the config connector machine state of FirestoreBackupSchedule
type FirestoreBackupScheduleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the FirestoreBackupSchedule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *FirestoreBackupScheduleObservedState `json:"observedState,omitempty"`
}

// FirestoreBackupScheduleObservedState is the state of the FirestoreBackupSchedule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.firestore.admin.v1.BackupSchedule
type FirestoreBackupScheduleObservedState struct {
	// Output only. The unique backup schedule identifier across all locations and
	//  databases for the given project.
	//
	//  This will be auto-assigned.
	//
	//  Format is
	//  `projects/{project}/databases/{database}/backupSchedules/{backup_schedule}`
	// +kcc:proto:field=google.firestore.admin.v1.BackupSchedule.name
	Name *string `json:"name,omitempty"`

	// Output only. The timestamp at which this backup schedule was created and
	//  effective since.
	//
	//  No backups will be created for this schedule before this time.
	// +kcc:proto:field=google.firestore.admin.v1.BackupSchedule.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp at which this backup schedule was most recently
	//  updated. When a backup schedule is first created, this is the same as
	//  create_time.
	// +kcc:proto:field=google.firestore.admin.v1.BackupSchedule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfirestorebackupschedule;gcpfirestorebackupschedules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FirestoreBackupSchedule is the Schema for the FirestoreBackupSchedule API
// +k8s:openapi-gen=true
type FirestoreBackupSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FirestoreBackupScheduleSpec   `json:"spec,omitempty"`
	Status FirestoreBackupScheduleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FirestoreBackupScheduleList contains a list of FirestoreBackupSchedule
type FirestoreBackupScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirestoreBackupSchedule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FirestoreBackupSchedule{}, &FirestoreBackupScheduleList{})
}
