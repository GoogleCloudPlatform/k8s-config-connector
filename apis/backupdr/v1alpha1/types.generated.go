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

// +kcc:proto=google.cloud.backupdr.v1.RuleConfigInfo
type RuleConfigInfo struct {
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

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.RuleConfigInfo
type RuleConfigInfoObservedState struct {
	// Output only. Backup Rule id fetched from backup plan.
	// +kcc:proto:field=google.cloud.backupdr.v1.RuleConfigInfo.rule_id
	RuleID *string `json:"ruleID,omitempty"`

	// Output only. The last backup state for rule.
	// +kcc:proto:field=google.cloud.backupdr.v1.RuleConfigInfo.last_backup_state
	LastBackupState *string `json:"lastBackupState,omitempty"`

	// Output only. google.rpc.Status object to store the last backup error.
	// +kcc:proto:field=google.cloud.backupdr.v1.RuleConfigInfo.last_backup_error
	LastBackupError *Status `json:"lastBackupError,omitempty"`

	// Output only. The point in time when the last successful backup was captured
	//  from the source.
	// +kcc:proto:field=google.cloud.backupdr.v1.RuleConfigInfo.last_successful_backup_consistency_time
	LastSuccessfulBackupConsistencyTime *string `json:"lastSuccessfulBackupConsistencyTime,omitempty"`
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
