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


// +kcc:proto=google.cloud.backupdr.v1.BackupRule
type BackupRule struct {
	// Required. Immutable. The unique id of this `BackupRule`. The `rule_id` is
	//  unique per `BackupPlan`.The `rule_id` must start with a lowercase letter
	//  followed by up to 62 lowercase letters, numbers, or hyphens. Pattern,
	//  /[a-z][a-z0-9-]{,62}/.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupRule.rule_id
	RuleID *string `json:"ruleID,omitempty"`

	// Required. Configures the duration for which backup data will be kept. It is
	//  defined in “days”. The value should be greater than or equal to minimum
	//  enforced retention of the backup vault.
	//
	//  Minimum value is 1 and maximum value is 90 for hourly backups.
	//  Minimum value is 1 and maximum value is 90 for daily backups.
	//  Minimum value is 7 and maximum value is 186 for weekly backups.
	//  Minimum value is 30 and maximum value is 732 for monthly backups.
	//  Minimum value is 365 and maximum value is 36159 for yearly backups.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupRule.backup_retention_days
	BackupRetentionDays *int32 `json:"backupRetentionDays,omitempty"`

	// Required. Defines a schedule that runs within the confines of a defined
	//  window of time.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupRule.standard_schedule
	StandardSchedule *StandardSchedule `json:"standardSchedule,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.BackupVault
type BackupVault struct {

	// Optional. The description of the BackupVault instance (2048 characters or
	//  less).
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.description
	Description *string `json:"description,omitempty"`

	// Optional. Resource labels to represent user provided metadata.
	//  No labels currently defined:
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The default and minimum enforced retention for each backup within
	//  the backup vault.  The enforced retention for each backup can be extended.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.backup_minimum_enforced_retention_duration
	BackupMinimumEnforcedRetentionDuration *string `json:"backupMinimumEnforcedRetentionDuration,omitempty"`

	// Optional. Server specified ETag for the backup vault resource to
	//  prevent simultaneous updates from overwiting each other.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Time after which the BackupVault resource is locked.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.effective_time
	EffectiveTime *string `json:"effectiveTime,omitempty"`

	// Optional. User annotations. See https://google.aip.dev/128#annotations
	//  Stores small amounts of arbitrary data.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Note: This field is added for future use case and will not be
	//  supported in the current release.
	//
	//  Access restriction for the backup vault.
	//  Default value is WITHIN_ORGANIZATION if not provided during creation.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.access_restriction
	AccessRestriction *string `json:"accessRestriction,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.BackupWindow
type BackupWindow struct {
	// Required. The hour of day (0-23) when the window starts for e.g. if value
	//  of start hour of day is 6 that mean backup window start at 6:00.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupWindow.start_hour_of_day
	StartHourOfDay *int32 `json:"startHourOfDay,omitempty"`

	// Required. The hour of day (1-24) when the window end for e.g. if value of
	//  end hour of day is 10 that mean backup window end time is 10:00.
	//
	//  End hour of day should be greater than start hour of day.
	//  0 <= start_hour_of_day < end_hour_of_day <= 24
	//
	//  End hour of day is not include in backup window that mean if
	//  end_hour_of_day= 10 jobs should start before 10:00.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupWindow.end_hour_of_day
	EndHourOfDay *int32 `json:"endHourOfDay,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.StandardSchedule
type StandardSchedule struct {
	// Required. Specifies the `RecurrenceType` for the schedule.
	// +kcc:proto:field=google.cloud.backupdr.v1.StandardSchedule.recurrence_type
	RecurrenceType *string `json:"recurrenceType,omitempty"`

	// Optional. Specifies frequency for hourly backups. A hourly frequency of 2
	//  means jobs will run every 2 hours from start time till end time defined.
	//
	//  This is required for `recurrence_type`, `HOURLY` and is not applicable
	//  otherwise. A validation error will occur if a value is supplied and
	//  `recurrence_type` is not `HOURLY`.
	//
	//  Value of hourly frequency should be between 6 and 23.
	//
	//  Reason for limit : We found that there is bandwidth limitation of 3GB/S for
	//  GMI while taking a backup and 5GB/S while doing a restore. Given the amount
	//  of parallel backups and restore we are targeting, this will potentially
	//  take the backup time to mins and hours (in worst case scenario).
	// +kcc:proto:field=google.cloud.backupdr.v1.StandardSchedule.hourly_frequency
	HourlyFrequency *int32 `json:"hourlyFrequency,omitempty"`

	// Optional. Specifies days of week like, MONDAY or TUESDAY, on which jobs
	//  will run.
	//
	//  This is required for `recurrence_type`, `WEEKLY` and is not applicable
	//  otherwise. A validation error will occur if a value is supplied and
	//  `recurrence_type` is not `WEEKLY`.
	// +kcc:proto:field=google.cloud.backupdr.v1.StandardSchedule.days_of_week
	DaysOfWeek []string `json:"daysOfWeek,omitempty"`

	// Optional. Specifies days of months like 1, 5, or 14 on which jobs will run.
	//
	//  Values for `days_of_month` are only applicable for `recurrence_type`,
	//  `MONTHLY` and `YEARLY`. A validation error will occur if other values are
	//  supplied.
	// +kcc:proto:field=google.cloud.backupdr.v1.StandardSchedule.days_of_month
	DaysOfMonth []int32 `json:"daysOfMonth,omitempty"`

	// Optional. Specifies a week day of the month like, FIRST SUNDAY or LAST
	//  MONDAY, on which jobs will run. This will be specified by two fields in
	//  `WeekDayOfMonth`, one for the day, e.g. `MONDAY`, and one for the week,
	//  e.g. `LAST`.
	//
	//  This field is only applicable for `recurrence_type`, `MONTHLY` and
	//  `YEARLY`. A validation error will occur if other values are supplied.
	// +kcc:proto:field=google.cloud.backupdr.v1.StandardSchedule.week_day_of_month
	WeekDayOfMonth *WeekDayOfMonth `json:"weekDayOfMonth,omitempty"`

	// Optional. Specifies the months of year, like `FEBRUARY` and/or `MAY`, on
	//  which jobs will run.
	//
	//  This field is only applicable when `recurrence_type` is `YEARLY`. A
	//  validation error will occur if other values are supplied.
	// +kcc:proto:field=google.cloud.backupdr.v1.StandardSchedule.months
	Months []string `json:"months,omitempty"`

	// Required. A BackupWindow defines the window of day during which backup jobs
	//  will run. Jobs are queued at the beginning of the window and will be marked
	//  as `NOT_RUN` if they do not start by the end of the window.
	//
	//  Note: running jobs will not be cancelled at the end of the window.
	// +kcc:proto:field=google.cloud.backupdr.v1.StandardSchedule.backup_window
	BackupWindow *BackupWindow `json:"backupWindow,omitempty"`

	// Required. The time zone to be used when interpreting the schedule.
	//  The value of this field must be a time zone name from the IANA tz database.
	//  See https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for the
	//  list of valid timezone names. For e.g., Europe/Paris.
	// +kcc:proto:field=google.cloud.backupdr.v1.StandardSchedule.time_zone
	TimeZone *string `json:"timeZone,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.WeekDayOfMonth
type WeekDayOfMonth struct {
	// Required. Specifies the week of the month.
	// +kcc:proto:field=google.cloud.backupdr.v1.WeekDayOfMonth.week_of_month
	WeekOfMonth *string `json:"weekOfMonth,omitempty"`

	// Required. Specifies the day of the week.
	// +kcc:proto:field=google.cloud.backupdr.v1.WeekDayOfMonth.day_of_week
	DayOfWeek *string `json:"dayOfWeek,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.WorkforceIdentityBasedManagementURI
type WorkforceIdentityBasedManagementURI struct {
}

// +kcc:proto=google.cloud.backupdr.v1.BackupVault
type BackupVaultObservedState struct {
	// Output only. Identifier. Name of the backup vault to create. It must have
	//  the
	//  format`"projects/{project}/locations/{location}/backupVaults/{backupvault}"`.
	//  `{backupvault}` cannot be changed after creation. It must be between 3-63
	//  characters long and must be unique within the project and location.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.name
	Name *string `json:"name,omitempty"`

	// Output only. The time when the instance was created.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the instance was updated.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Set to true when there are no backups nested under this
	//  resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.deletable
	Deletable *bool `json:"deletable,omitempty"`

	// Output only. The BackupVault resource instance state.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.state
	State *string `json:"state,omitempty"`

	// Output only. The number of backups in this backup vault.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.backup_count
	BackupCount *int64 `json:"backupCount,omitempty"`

	// Output only. Service account used by the BackupVault Service for this
	//  BackupVault.  The user should grant this account permissions in their
	//  workload project to enable the service to run backups and restores there.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Output only. Total size of the storage used by all backup resources.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.total_stored_bytes
	TotalStoredBytes *int64 `json:"totalStoredBytes,omitempty"`

	// Output only. Immutable after resource creation until resource deletion.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.uid
	Uid *string `json:"uid,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.WorkforceIdentityBasedManagementURI
type WorkforceIdentityBasedManagementURIObservedState struct {
	// Output only. First party Management URI for Google Identities.
	// +kcc:proto:field=google.cloud.backupdr.v1.WorkforceIdentityBasedManagementURI.first_party_management_uri
	FirstPartyManagementURI *string `json:"firstPartyManagementURI,omitempty"`

	// Output only. Third party Management URI for External Identity Providers.
	// +kcc:proto:field=google.cloud.backupdr.v1.WorkforceIdentityBasedManagementURI.third_party_management_uri
	ThirdPartyManagementURI *string `json:"thirdPartyManagementURI,omitempty"`
}
