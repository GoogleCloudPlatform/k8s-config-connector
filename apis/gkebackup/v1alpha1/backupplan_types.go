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
	container "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GKEBackupBackupPlanGVK = GroupVersion.WithKind("GKEBackupBackupPlan")

// +kcc:proto=google.cloud.gkebackup.v1.EncryptionKey
type EncryptionKey struct {
	// Optional. Google Cloud KMS encryption key.
	// +kcc:proto:field=google.cloud.gkebackup.v1.EncryptionKey.gcp_kms_encryption_key
	KMSKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.gkebackup.v1.BackupPlan.Schedule
type BackupPlan_ScheduleObservedState struct {
	// Output only. Start time of next scheduled backup under this BackupPlan by
	//  either cron_schedule or rpo config.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.Schedule.next_scheduled_backup_time
	NextScheduledBackupTime *string `json:"nextScheduledBackupTime,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.RpoConfig
type RPOConfig struct {
	// Required. Defines the target RPO for the BackupPlan in minutes, which means
	//  the target maximum data loss in time that is acceptable for this
	//  BackupPlan. This must be at least 60, i.e., 1 hour, and at most 86400,
	//  i.e., 60 days.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RpoConfig.target_rpo_minutes
	TargetRPOMinutes *int32 `json:"targetRPOMinutes,omitempty"`

	// Optional. User specified time windows during which backup can NOT happen
	//  for this BackupPlan - backups should start and finish outside of any given
	//  exclusion window. Note: backup jobs will be scheduled to start and
	//  finish outside the duration of the window as much as possible, but
	//  running jobs will not get canceled when it runs into the window.
	//  All the time and date values in exclusion_windows entry in the API are in
	//  UTC.
	//  We only allow <=1 recurrence (daily or weekly) exclusion window for a
	//  BackupPlan while no restriction on number of single occurrence
	//  windows.
	// +kcc:proto:field=google.cloud.gkebackup.v1.RpoConfig.exclusion_windows
	ExclusionWindows []ExclusionWindow `json:"exclusionWindows,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.ExclusionWindow
type ExclusionWindow struct {
	// Required. Specifies the start time of the window using time of the day in
	//  UTC.
	// +kcc:proto:field=google.cloud.gkebackup.v1.ExclusionWindow.start_time
	// +required
	StartTime *TimeOfDay `json:"startTime,omitempty"`

	// Required. Specifies duration of the window.
	//  Duration must be >= 5 minutes and < (target RPO - 20 minutes).
	//  Additional restrictions based on the recurrence type to allow some time for
	//  backup to happen:
	//  - single_occurrence_date:  no restriction, but UI may warn about this when
	//  duration >= target RPO
	//  - daily window: duration < 24 hours
	//  - weekly window:
	//    - days of week includes all seven days of a week: duration < 24 hours
	//    - all other weekly window: duration < 168 hours (i.e., 24 * 7 hours)
	// +kcc:proto:field=google.cloud.gkebackup.v1.ExclusionWindow.duration
	// +required
	Duration *string `json:"duration,omitempty"`

	// No recurrence. The exclusion window occurs only once and on this
	//  date in UTC.
	// +kcc:proto:field=google.cloud.gkebackup.v1.ExclusionWindow.single_occurrence_date
	SingleOccurrenceDate *Date `json:"singleOccurrenceDate,omitempty"`

	// The exclusion window occurs every day if set to "True".
	//  Specifying this field to "False" is an error.
	// +kcc:proto:field=google.cloud.gkebackup.v1.ExclusionWindow.daily
	Daily *bool `json:"daily,omitempty"`

	// The exclusion window occurs on these days of each week in UTC.
	// +kcc:proto:field=google.cloud.gkebackup.v1.ExclusionWindow.days_of_week
	DaysOfWeek *ExclusionWindow_DayOfWeekList `json:"daysOfWeek,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.BackupPlan.Schedule
type BackupPlan_Schedule struct {
	// Optional. A standard [cron](https://wikipedia.com/wiki/cron) string that
	//  defines a repeating schedule for creating Backups via this BackupPlan.
	//  This is mutually exclusive with the
	//  [rpo_config][google.cloud.gkebackup.v1.BackupPlan.Schedule.rpo_config]
	//  field since at most one schedule can be defined for a BackupPlan. If this
	//  is defined, then
	//  [backup_retain_days][google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_retain_days]
	//  must also be defined.
	//
	//  Default (empty): no automatic backup creation will occur.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.Schedule.cron_schedule
	CronSchedule *string `json:"cronSchedule,omitempty"`

	// Optional. This flag denotes whether automatic Backup creation is paused
	//  for this BackupPlan.
	//
	//  Default: False
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.Schedule.paused
	Paused *bool `json:"paused,omitempty"`

	// Optional. Defines the RPO schedule configuration for this BackupPlan.
	//  This is mutually exclusive with the
	//  [cron_schedule][google.cloud.gkebackup.v1.BackupPlan.Schedule.cron_schedule]
	//  field since at most one schedule can be defined for a BackupPLan. If this
	//  is defined, then
	//  [backup_retain_days][google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_retain_days]
	//  must also be defined.
	//
	//  Default (empty): no automatic backup creation will occur.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.Schedule.rpo_config
	RPOConfig *RPOConfig `json:"rpoConfig,omitempty"`
}

type Parent struct {
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
}

// GKEBackupBackupPlanSpec defines the desired state of GKEBackupBackupPlan
// +kcc:spec:proto=google.cloud.gkebackup.v1.BackupPlan
type GKEBackupBackupPlanSpec struct {
	// The GKEBackupBackupPlan name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Optional. User specified descriptive string for this BackupPlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. The source cluster from which Backups will be created
	//  via this BackupPlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.cluster
	// +required
	ClusterRef *container.ContainerClusterRef `json:"clusterRef,omitempty"`

	// Optional. RetentionPolicy governs lifecycle of Backups created under this
	//  plan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.retention_policy
	RetentionPolicy *BackupPlan_RetentionPolicy `json:"retentionPolicy,omitempty"`

	// Optional. A set of custom labels supplied by user.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Defines a schedule for automatic Backup creation via this
	//  BackupPlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.backup_schedule
	BackupSchedule *BackupPlan_Schedule `json:"backupSchedule,omitempty"`

	// Optional. This flag indicates whether this BackupPlan has been deactivated.
	//  Setting this field to True locks the BackupPlan such that no further
	//  updates will be allowed (except deletes), including the deactivated field
	//  itself. It also prevents any new Backups from being created via this
	//  BackupPlan (including scheduled Backups).
	//
	//  Default: False
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.deactivated
	Deactivated *bool `json:"deactivated,omitempty"`

	// Optional. Defines the configuration of Backups created via this BackupPlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.backup_config
	BackupConfig *BackupPlan_BackupConfig `json:"backupConfig,omitempty"`
}

// GKEBackupBackupPlanStatus defines the config connector machine state of GKEBackupBackupPlan
type GKEBackupBackupPlanStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GKEBackupBackupPlan resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GKEBackupBackupPlanObservedState `json:"observedState,omitempty"`
}

// GKEBackupBackupPlanObservedState is the state of the GKEBackupBackupPlan resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.gkebackup.v1.BackupPlan
type GKEBackupBackupPlanObservedState struct {
	// Output only. The full name of the BackupPlan resource.
	//  Format: `projects/*/locations/*/backupPlans/*`
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. Server generated global unique identifier of
	//  [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.uid
	UID *string `json:"uid,omitempty"`

	// Output only. The timestamp when this BackupPlan resource was created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this BackupPlan resource was last
	//  updated.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Defines a schedule for automatic Backup creation via this
	//  BackupPlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.backup_schedule
	BackupSchedule *BackupPlan_ScheduleObservedState `json:"backupSchedule,omitempty"`

	// Output only. `etag` is used for optimistic concurrency control as a way to
	//  help prevent simultaneous updates of a backup plan from overwriting each
	//  other. It is strongly suggested that systems make use of the 'etag' in the
	//  read-modify-write cycle to perform BackupPlan updates in order to avoid
	//  race conditions: An `etag` is returned in the response to `GetBackupPlan`,
	//  and systems are expected to put that etag in the request to
	//  `UpdateBackupPlan` or `DeleteBackupPlan` to ensure that their change
	//  will be applied to the same version of the resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The number of Kubernetes Pods backed up in the
	//  last successful Backup created via this BackupPlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.protected_pod_count
	ProtectedPodCount *int32 `json:"protectedPodCount,omitempty"`

	// Output only. State of the BackupPlan. This State field reflects the
	//  various stages a BackupPlan can be in
	//  during the Create operation. It will be set to "DEACTIVATED"
	//  if the BackupPlan is deactivated on an Update
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.state
	State *string `json:"state,omitempty"`

	// Output only. Human-readable description of why BackupPlan is in the current
	//  `state`
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.state_reason
	StateReason *string `json:"stateReason,omitempty"`

	// Output only. A number that represents the current risk level of this
	//  BackupPlan from RPO perspective with 1 being no risk and 5 being highest
	//  risk.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.rpo_risk_level
	RPORiskLevel *int32 `json:"rpoRiskLevel,omitempty"`

	// Output only. Human-readable description of why the BackupPlan is in the
	//  current rpo_risk_level and action items if any.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.rpo_risk_reason
	RPORiskReason *string `json:"rpoRiskReason,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkebackupbackupplan;gcpgkebackupbackupplans
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEBackupBackupPlan is the Schema for the GKEBackupBackupPlan API
// +k8s:openapi-gen=true
type GKEBackupBackupPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GKEBackupBackupPlanSpec   `json:"spec,omitempty"`
	Status GKEBackupBackupPlanStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GKEBackupBackupPlanList contains a list of GKEBackupBackupPlan
type GKEBackupBackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEBackupBackupPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEBackupBackupPlan{}, &GKEBackupBackupPlanList{})
}
