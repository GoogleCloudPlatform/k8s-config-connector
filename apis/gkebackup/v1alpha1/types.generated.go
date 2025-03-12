// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

// +kcc:proto=google.cloud.gkebackup.v1.BackupPlan.BackupConfig
type BackupPlan_BackupConfig struct {
	// If True, include all namespaced resources
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.all_namespaces
	AllNamespaces *bool `json:"allNamespaces,omitempty"`

	// If set, include just the resources in the listed namespaces.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.selected_namespaces
	SelectedNamespaces *Namespaces `json:"selectedNamespaces,omitempty"`

	// If set, include just the resources referenced by the listed
	//  ProtectedApplications.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.selected_applications
	SelectedApplications *NamespacedNames `json:"selectedApplications,omitempty"`

	// Optional. This flag specifies whether volume data should be backed up
	//  when PVCs are included in the scope of a Backup.
	//
	//  Default: False
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.include_volume_data
	IncludeVolumeData *bool `json:"includeVolumeData,omitempty"`

	// Optional. This flag specifies whether Kubernetes Secret resources should
	//  be included when they fall into the scope of Backups.
	//
	//  Default: False
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.include_secrets
	IncludeSecrets *bool `json:"includeSecrets,omitempty"`

	// Optional. This defines a customer managed encryption key that will be
	//  used to encrypt the "config" portion (the Kubernetes resources) of
	//  Backups created via this plan.
	//
	//  Default (empty): Config backup artifacts will not be encrypted.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.encryption_key
	EncryptionKey *EncryptionKey `json:"encryptionKey,omitempty"`

	// Optional. If false, Backups will fail when Backup for GKE detects
	//  Kubernetes configuration that is non-standard or
	//  requires additional setup to restore.
	//
	//  Default: False
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.BackupConfig.permissive_mode
	PermissiveMode *bool `json:"permissiveMode,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy
type BackupPlan_RetentionPolicy struct {
	// Optional. Minimum age for Backups created via this BackupPlan (in days).
	//  This field MUST be an integer value between 0-90 (inclusive).
	//  A Backup created under this BackupPlan will NOT be deletable until it
	//  reaches Backup's (create_time + backup_delete_lock_days).
	//  Updating this field of a BackupPlan does NOT affect existing Backups
	//  under it. Backups created AFTER a successful update will inherit
	//  the new value.
	//
	//  Default: 0 (no delete blocking)
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_delete_lock_days
	BackupDeleteLockDays *int32 `json:"backupDeleteLockDays,omitempty"`

	// Optional. The default maximum age of a Backup created via this
	//  BackupPlan. This field MUST be an integer value >= 0 and <= 365. If
	//  specified, a Backup created under this BackupPlan will be automatically
	//  deleted after its age reaches (create_time + backup_retain_days). If not
	//  specified, Backups created under this BackupPlan will NOT be subject to
	//  automatic deletion. Updating this field does NOT affect existing Backups
	//  under it. Backups created AFTER a successful update will automatically
	//  pick up the new value. NOTE: backup_retain_days must be >=
	//  [backup_delete_lock_days][google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_delete_lock_days].
	//  If
	//  [cron_schedule][google.cloud.gkebackup.v1.BackupPlan.Schedule.cron_schedule]
	//  is defined, then this must be
	//  <= 360 * the creation interval. If
	//  [rpo_config][google.cloud.gkebackup.v1.BackupPlan.Schedule.rpo_config] is
	//  defined, then this must be
	//  <= 360 * [target_rpo_minutes][Schedule.rpo_config.target_rpo_minutes] /
	//  (1440minutes/day).
	//
	//  Default: 0 (no automatic deletion)
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_retain_days
	BackupRetainDays *int32 `json:"backupRetainDays,omitempty"`

	// Optional. This flag denotes whether the retention policy of this
	//  BackupPlan is locked.  If set to True, no further update is allowed on
	//  this policy, including the `locked` field itself.
	//
	//  Default: False
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.locked
	Locked *bool `json:"locked,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.EncryptionKey
type EncryptionKey struct {
	// Optional. Google Cloud KMS encryption key. Format:
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*`
	// +kcc:proto:field=google.cloud.gkebackup.v1.EncryptionKey.gcp_kms_encryption_key
	GcpKMSEncryptionKey *string `json:"gcpKMSEncryptionKey,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.ExclusionWindow
type ExclusionWindow struct {
	// Required. Specifies the start time of the window using time of the day in
	//  UTC.
	// +kcc:proto:field=google.cloud.gkebackup.v1.ExclusionWindow.start_time
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

// +kcc:proto=google.cloud.gkebackup.v1.ExclusionWindow.DayOfWeekList
type ExclusionWindow_DayOfWeekList struct {
	// Optional. A list of days of week.
	// +kcc:proto:field=google.cloud.gkebackup.v1.ExclusionWindow.DayOfWeekList.days_of_week
	DaysOfWeek []string `json:"daysOfWeek,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.NamespacedName
type NamespacedName struct {
	// Optional. The Namespace of the Kubernetes resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.NamespacedName.namespace
	Namespace *string `json:"namespace,omitempty"`

	// Optional. The name of the Kubernetes resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.NamespacedName.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.NamespacedNames
type NamespacedNames struct {
	// Optional. A list of namespaced Kubernetes resources.
	// +kcc:proto:field=google.cloud.gkebackup.v1.NamespacedNames.namespaced_names
	NamespacedNames []NamespacedName `json:"namespacedNames,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.Namespaces
type Namespaces struct {
	// Optional. A list of Kubernetes Namespaces
	// +kcc:proto:field=google.cloud.gkebackup.v1.Namespaces.namespaces
	Namespaces []string `json:"namespaces,omitempty"`
}

// +kcc:proto=google.type.Date
type Date struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without
	//  a year.
	// +kcc:proto:field=google.type.Date.year
	Year *int32 `json:"year,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a
	//  month and day.
	// +kcc:proto:field=google.type.Date.month
	Month *int32 `json:"month,omitempty"`

	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
	//  to specify a year by itself or a year and month where the day isn't
	//  significant.
	// +kcc:proto:field=google.type.Date.day
	Day *int32 `json:"day,omitempty"`
}

// +kcc:proto=google.type.TimeOfDay
type TimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose
	//  to allow the value "24:00:00" for scenarios like business closing time.
	// +kcc:proto:field=google.type.TimeOfDay.hours
	Hours *int32 `json:"hours,omitempty"`

	// Minutes of hour of day. Must be from 0 to 59.
	// +kcc:proto:field=google.type.TimeOfDay.minutes
	Minutes *int32 `json:"minutes,omitempty"`

	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may
	//  allow the value 60 if it allows leap-seconds.
	// +kcc:proto:field=google.type.TimeOfDay.seconds
	Seconds *int32 `json:"seconds,omitempty"`

	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	// +kcc:proto:field=google.type.TimeOfDay.nanos
	Nanos *int32 `json:"nanos,omitempty"`
}

// +kcc:proto=google.cloud.gkebackup.v1.BackupPlan.Schedule
type BackupPlan_ScheduleObservedState struct {
	// Output only. Start time of next scheduled backup under this BackupPlan by
	//  either cron_schedule or rpo config.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.Schedule.next_scheduled_backup_time
	NextScheduledBackupTime *string `json:"nextScheduledBackupTime,omitempty"`
}
