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

// API sources for BackupDRRestoreVerificationPlan, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/backupdr/v1/restore_verification_plan.proto
// +kcc:source:service-docs=https://cloud.google.com/backup-disaster-recovery

package v1alpha1

import (
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BackupDRRestoreVerificationPlanGVK = GroupVersion.WithKind("BackupDRRestoreVerificationPlan")

// BackupDRRestoreVerificationPlanSpec defines the desired state of BackupDRRestoreVerificationPlan
// +kcc:spec:proto=google.cloud.backupdr.v1.RestoreVerificationPlan
// +kcc:required-from-proto
type BackupDRRestoreVerificationPlanSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The BackupDRRestoreVerificationPlan name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The description of the `RestoreVerificationPlan` resource.
	//
	//  The description allows for additional details about
	//  `RestoreVerificationPlan` and its use cases to be provided. An example
	//  description is the following: "This is a restore verification plan that
	//  performs a daily verification and cleans up the resources immediately
	//  after." The description must be at most 2048 characters.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreVerificationPlan.description
	Description *string `json:"description,omitempty"`

	// Required. Describes the schedule on which the restore verification will be
	//  executed.
	//
	//  The schedule is specified as a
	//  [Crontab](https://en.wikipedia.org/wiki/Cron#Overview) string.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreVerificationPlan.schedule
	// +required
	Schedule *string `json:"schedule,omitempty"`

	// Optional. Specifies the time zone to be used in interpreting
	//  [schedule][google.cloud.backupdr.v1.RestoreVerificationPlan.schedule].
	//  The value of this field must be a time zone name from the [tz
	//  database](https://en.wikipedia.org/wiki/Tz_database), like
	//  "America/New_York" or "UTC". Defaults to `UTC` if not specified.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreVerificationPlan.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Optional. The cleanup rule to apply if the restore verification succeeds.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreVerificationPlan.success_cleanup_rule
	SuccessCleanupRule *RestoreVerificationCleanupRule `json:"successCleanupRule,omitempty"`

	// Optional. The cleanup rule to apply if the restore verification fails.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreVerificationPlan.failure_cleanup_rule
	FailureCleanupRule *RestoreVerificationCleanupRule `json:"failureCleanupRule,omitempty"`
}

// BackupDRRestoreVerificationPlanStatus defines the config connector machine state of BackupDRRestoreVerificationPlan
type BackupDRRestoreVerificationPlanStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BackupDRRestoreVerificationPlan resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BackupDRRestoreVerificationPlanObservedState `json:"observedState,omitempty"`
}

// BackupDRRestoreVerificationPlanObservedState is the state of the BackupDRRestoreVerificationPlan resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.backupdr.v1.RestoreVerificationPlan
type BackupDRRestoreVerificationPlanObservedState struct {
	// Output only. When the `RestoreVerificationPlan` was created.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreVerificationPlan.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. When the `RestoreVerificationPlan` was last updated.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreVerificationPlan.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The `State` for the `RestoreVerificationPlan`.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreVerificationPlan.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbackupdrrestoreverificationplan;gcpbackupdrrestoreverificationplans
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BackupDRRestoreVerificationPlan is the Schema for the BackupDRRestoreVerificationPlan API
// +k8s:openapi-gen=true
type BackupDRRestoreVerificationPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BackupDRRestoreVerificationPlanSpec   `json:"spec,omitempty"`
	Status BackupDRRestoreVerificationPlanStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BackupDRRestoreVerificationPlanList contains a list of BackupDRRestoreVerificationPlan
type BackupDRRestoreVerificationPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupDRRestoreVerificationPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupDRRestoreVerificationPlan{}, &BackupDRRestoreVerificationPlanList{})
}
