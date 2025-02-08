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


// +kcc:proto=google.firestore.admin.v1.BackupSchedule
type BackupSchedule struct {

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

// +kcc:proto=google.firestore.admin.v1.DailyRecurrence
type DailyRecurrence struct {
}

// +kcc:proto=google.firestore.admin.v1.WeeklyRecurrence
type WeeklyRecurrence struct {
	// The day of week to run.
	//
	//  DAY_OF_WEEK_UNSPECIFIED is not allowed.
	// +kcc:proto:field=google.firestore.admin.v1.WeeklyRecurrence.day
	Day *string `json:"day,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.BackupSchedule
type BackupScheduleObservedState struct {
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
