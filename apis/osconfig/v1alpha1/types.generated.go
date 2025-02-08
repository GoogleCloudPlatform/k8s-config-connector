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


// +kcc:proto=google.cloud.osconfig.v1beta.AptSettings
type AptSettings struct {
	// By changing the type to DIST, the patching is performed
	//  using `apt-get dist-upgrade` instead.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.AptSettings.type
	Type *string `json:"type,omitempty"`

	// List of packages to exclude from update. These packages will be excluded
	// +kcc:proto:field=google.cloud.osconfig.v1beta.AptSettings.excludes
	Excludes []string `json:"excludes,omitempty"`

	// An exclusive list of packages to be updated. These are the only packages
	//  that will be updated. If these packages are not installed, they will be
	//  ignored. This field cannot be specified with any other patch configuration
	//  fields.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.AptSettings.exclusive_packages
	ExclusivePackages []string `json:"exclusivePackages,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.ExecStep
type ExecStep struct {
	// The ExecStepConfig for all Linux VMs targeted by the PatchJob.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ExecStep.linux_exec_step_config
	LinuxExecStepConfig *ExecStepConfig `json:"linuxExecStepConfig,omitempty"`

	// The ExecStepConfig for all Windows VMs targeted by the PatchJob.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ExecStep.windows_exec_step_config
	WindowsExecStepConfig *ExecStepConfig `json:"windowsExecStepConfig,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.ExecStepConfig
type ExecStepConfig struct {
	// An absolute path to the executable on the VM.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ExecStepConfig.local_path
	LocalPath *string `json:"localPath,omitempty"`

	// A Google Cloud Storage object containing the executable.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ExecStepConfig.gcs_object
	GcsObject *GcsObject `json:"gcsObject,omitempty"`

	// Defaults to [0]. A list of possible return values that the
	//  execution can return to indicate a success.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ExecStepConfig.allowed_success_codes
	AllowedSuccessCodes []int32 `json:"allowedSuccessCodes,omitempty"`

	// The script interpreter to use to run the script. If no interpreter is
	//  specified the script will be executed directly, which will likely
	//  only succeed for scripts with [shebang lines]
	//  (https://en.wikipedia.org/wiki/Shebang_\(Unix\)).
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ExecStepConfig.interpreter
	Interpreter *string `json:"interpreter,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.FixedOrPercent
type FixedOrPercent struct {
	// Specifies a fixed value.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.FixedOrPercent.fixed
	Fixed *int32 `json:"fixed,omitempty"`

	// Specifies the relative value defined as a percentage, which will be
	//  multiplied by a reference value.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.FixedOrPercent.percent
	Percent *int32 `json:"percent,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.GcsObject
type GcsObject struct {
	// Required. Bucket of the Google Cloud Storage object.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.GcsObject.bucket
	Bucket *string `json:"bucket,omitempty"`

	// Required. Name of the Google Cloud Storage object.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.GcsObject.object
	Object *string `json:"object,omitempty"`

	// Required. Generation number of the Google Cloud Storage object. This is used to
	//  ensure that the ExecStep specified by this PatchJob does not change.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.GcsObject.generation_number
	GenerationNumber *int64 `json:"generationNumber,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.GooSettings
type GooSettings struct {
}

// +kcc:proto=google.cloud.osconfig.v1beta.MonthlySchedule
type MonthlySchedule struct {
	// Required. Week day in a month.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.MonthlySchedule.week_day_of_month
	WeekDayOfMonth *WeekDayOfMonth `json:"weekDayOfMonth,omitempty"`

	// Required. One day of the month. 1-31 indicates the 1st to the 31st day. -1
	//  indicates the last day of the month.
	//  Months without the target day will be skipped. For example, a schedule to
	//  run "every month on the 31st" will not run in February, April, June, etc.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.MonthlySchedule.month_day
	MonthDay *int32 `json:"monthDay,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.OneTimeSchedule
type OneTimeSchedule struct {
	// Required. The desired patch job execution time.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.OneTimeSchedule.execute_time
	ExecuteTime *string `json:"executeTime,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.PatchConfig
type PatchConfig struct {
	// Post-patch reboot settings.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.reboot_config
	RebootConfig *string `json:"rebootConfig,omitempty"`

	// Apt update settings. Use this setting to override the default `apt` patch
	//  rules.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.apt
	Apt *AptSettings `json:"apt,omitempty"`

	// Yum update settings. Use this setting to override the default `yum` patch
	//  rules.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.yum
	Yum *YumSettings `json:"yum,omitempty"`

	// Goo update settings. Use this setting to override the default `goo` patch
	//  rules.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.goo
	Goo *GooSettings `json:"goo,omitempty"`

	// Zypper update settings. Use this setting to override the default `zypper`
	//  patch rules.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.zypper
	Zypper *ZypperSettings `json:"zypper,omitempty"`

	// Windows update settings. Use this override the default windows patch rules.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.windows_update
	WindowsUpdate *WindowsUpdateSettings `json:"windowsUpdate,omitempty"`

	// The `ExecStep` to run before the patch update.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.pre_step
	PreStep *ExecStep `json:"preStep,omitempty"`

	// The `ExecStep` to run after the patch update.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.post_step
	PostStep *ExecStep `json:"postStep,omitempty"`

	// Allows the patch job to run on Managed instance groups (MIGs).
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.mig_instances_allowed
	MigInstancesAllowed *bool `json:"migInstancesAllowed,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.PatchDeployment
type PatchDeployment struct {
	// Unique name for the patch deployment resource in a project. The patch
	//  deployment name is in the form:
	//  `projects/{project_id}/patchDeployments/{patch_deployment_id}`.
	//  This field is ignored when you create a new patch deployment.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchDeployment.name
	Name *string `json:"name,omitempty"`

	// Optional. Description of the patch deployment. Length of the description is limited
	//  to 1024 characters.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchDeployment.description
	Description *string `json:"description,omitempty"`

	// Required. VM instances to patch.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchDeployment.instance_filter
	InstanceFilter *PatchInstanceFilter `json:"instanceFilter,omitempty"`

	// Optional. Patch configuration that is applied.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchDeployment.patch_config
	PatchConfig *PatchConfig `json:"patchConfig,omitempty"`

	// Optional. Duration of the patch. After the duration ends, the patch times out.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchDeployment.duration
	Duration *string `json:"duration,omitempty"`

	// Required. Schedule a one-time execution.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchDeployment.one_time_schedule
	OneTimeSchedule *OneTimeSchedule `json:"oneTimeSchedule,omitempty"`

	// Required. Schedule recurring executions.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchDeployment.recurring_schedule
	RecurringSchedule *RecurringSchedule `json:"recurringSchedule,omitempty"`

	// Optional. Rollout strategy of the patch job.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchDeployment.rollout
	Rollout *PatchRollout `json:"rollout,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.PatchInstanceFilter
type PatchInstanceFilter struct {
	// Target all VM instances in the project. If true, no other criteria is
	//  permitted.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchInstanceFilter.all
	All *bool `json:"all,omitempty"`

	// Targets VM instances matching at least one of these label sets. This allows
	//  targeting of disparate groups, for example "env=prod or env=staging".
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchInstanceFilter.group_labels
	GroupLabels []PatchInstanceFilter_GroupLabel `json:"groupLabels,omitempty"`

	// Targets VM instances in ANY of these zones. Leave empty to target VM
	//  instances in any zone.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchInstanceFilter.zones
	Zones []string `json:"zones,omitempty"`

	// Targets any of the VM instances specified. Instances are specified by their
	//  URI in the form `zones/[ZONE]/instances/[INSTANCE_NAME]`,
	//  `projects/[PROJECT_ID]/zones/[ZONE]/instances/[INSTANCE_NAME]`, or
	//  `https://www.googleapis.com/compute/v1/projects/[PROJECT_ID]/zones/[ZONE]/instances/[INSTANCE_NAME]`
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchInstanceFilter.instances
	Instances []string `json:"instances,omitempty"`

	// Targets VMs whose name starts with one of these prefixes. Similar to
	//  labels, this is another way to group VMs when targeting configs, for
	//  example prefix="prod-".
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchInstanceFilter.instance_name_prefixes
	InstanceNamePrefixes []string `json:"instanceNamePrefixes,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.PatchInstanceFilter.GroupLabel
type PatchInstanceFilter_GroupLabel struct {
	// Compute Engine instance labels that must be present for a VM instance to
	//  be targeted by this filter.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchInstanceFilter.GroupLabel.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.PatchRollout
type PatchRollout struct {
	// Mode of the patch rollout.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchRollout.mode
	Mode *string `json:"mode,omitempty"`

	// The maximum number (or percentage) of VMs per zone to disrupt at any given
	//  moment. The number of VMs calculated from multiplying the percentage by the
	//  total number of VMs in a zone is rounded up.
	//
	//  During patching, a VM is considered disrupted from the time the agent is
	//  notified to begin until patching has completed. This disruption time
	//  includes the time to complete reboot and any post-patch steps.
	//
	//  A VM contributes to the disruption budget if its patching operation fails
	//  either when applying the patches, running pre or post patch steps, or if it
	//  fails to respond with a success notification before timing out. VMs that
	//  are not running or do not have an active agent do not count toward this
	//  disruption budget.
	//
	//  For zone-by-zone rollouts, if the disruption budget in a zone is exceeded,
	//  the patch job stops, because continuing to the next zone requires
	//  completion of the patch process in the previous zone.
	//
	//  For example, if the disruption budget has a fixed value of `10`, and 8 VMs
	//  fail to patch in the current zone, the patch job continues to patch 2 VMs
	//  at a time until the zone is completed. When that zone is completed
	//  successfully, patching begins with 10 VMs at a time in the next zone. If 10
	//  VMs in the next zone fail to patch, the patch job stops.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchRollout.disruption_budget
	DisruptionBudget *FixedOrPercent `json:"disruptionBudget,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.RecurringSchedule
type RecurringSchedule struct {
	// Required. Defines the time zone that `time_of_day` is relative to.
	//  The rules for daylight saving time are determined by the chosen time zone.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.RecurringSchedule.time_zone
	TimeZone *TimeZone `json:"timeZone,omitempty"`

	// Optional. The time that the recurring schedule becomes effective.
	//  Defaults to `create_time` of the patch deployment.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.RecurringSchedule.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. The end time at which a recurring patch deployment schedule is no longer
	//  active.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.RecurringSchedule.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Required. Time of the day to run a recurring deployment.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.RecurringSchedule.time_of_day
	TimeOfDay *TimeOfDay `json:"timeOfDay,omitempty"`

	// Required. The frequency unit of this recurring schedule.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.RecurringSchedule.frequency
	Frequency *string `json:"frequency,omitempty"`

	// Required. Schedule with weekly executions.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.RecurringSchedule.weekly
	Weekly *WeeklySchedule `json:"weekly,omitempty"`

	// Required. Schedule with monthly executions.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.RecurringSchedule.monthly
	Monthly *MonthlySchedule `json:"monthly,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.WeekDayOfMonth
type WeekDayOfMonth struct {
	// Required. Week number in a month. 1-4 indicates the 1st to 4th week of the month. -1
	//  indicates the last week of the month.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.WeekDayOfMonth.week_ordinal
	WeekOrdinal *int32 `json:"weekOrdinal,omitempty"`

	// Required. A day of the week.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.WeekDayOfMonth.day_of_week
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// Optional. Represents the number of days before or after the given week day of month
	//  that the patch deployment is scheduled for. For example if `week_ordinal`
	//  and `day_of_week` values point to the second day of the month and this
	//  `day_offset` value is set to `3`, the patch deployment takes place three
	//  days after the second Tuesday of the month. If this value is negative, for
	//  example -5, the patches  are deployed five days before before the second
	//  Tuesday of the month. Allowed values are in range `[-30, 30]`.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.WeekDayOfMonth.day_offset
	DayOffset *int32 `json:"dayOffset,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.WeeklySchedule
type WeeklySchedule struct {
	// Required. Day of the week.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.WeeklySchedule.day_of_week
	DayOfWeek *string `json:"dayOfWeek,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.WindowsUpdateSettings
type WindowsUpdateSettings struct {
	// Only apply updates of these windows update classifications. If empty, all
	//  updates are applied.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.WindowsUpdateSettings.classifications
	Classifications []string `json:"classifications,omitempty"`

	// List of KBs to exclude from update.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.WindowsUpdateSettings.excludes
	Excludes []string `json:"excludes,omitempty"`

	// An exclusive list of kbs to be updated. These are the only patches
	//  that will be updated. This field must not be used with other
	//  patch configurations.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.WindowsUpdateSettings.exclusive_patches
	ExclusivePatches []string `json:"exclusivePatches,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.YumSettings
type YumSettings struct {
	// Adds the `--security` flag to `yum update`. Not supported on
	//  all platforms.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.YumSettings.security
	Security *bool `json:"security,omitempty"`

	// Will cause patch to run `yum update-minimal` instead.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.YumSettings.minimal
	Minimal *bool `json:"minimal,omitempty"`

	// List of packages to exclude from update. These packages are excluded by
	//  using the yum `--exclude` flag.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.YumSettings.excludes
	Excludes []string `json:"excludes,omitempty"`

	// An exclusive list of packages to be updated. These are the only packages
	//  that will be updated. If these packages are not installed, they will be
	//  ignored. This field must not be specified with any other patch
	//  configuration fields.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.YumSettings.exclusive_packages
	ExclusivePackages []string `json:"exclusivePackages,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.ZypperSettings
type ZypperSettings struct {
	// Adds the `--with-optional` flag to `zypper patch`.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperSettings.with_optional
	WithOptional *bool `json:"withOptional,omitempty"`

	// Adds the `--with-update` flag, to `zypper patch`.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperSettings.with_update
	WithUpdate *bool `json:"withUpdate,omitempty"`

	// Install only patches with these categories.
	//  Common categories include security, recommended, and feature.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperSettings.categories
	Categories []string `json:"categories,omitempty"`

	// Install only patches with these severities.
	//  Common severities include critical, important, moderate, and low.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperSettings.severities
	Severities []string `json:"severities,omitempty"`

	// List of patches to exclude from update.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperSettings.excludes
	Excludes []string `json:"excludes,omitempty"`

	// An exclusive list of patches to be updated. These are the only patches
	//  that will be installed using 'zypper patch patch:<patch_name>' command.
	//  This field must not be used with any other patch configuration fields.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperSettings.exclusive_patches
	ExclusivePatches []string `json:"exclusivePatches,omitempty"`
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

// +kcc:proto=google.type.TimeZone
type TimeZone struct {
	// IANA Time Zone Database time zone, e.g. "America/New_York".
	// +kcc:proto:field=google.type.TimeZone.id
	ID *string `json:"id,omitempty"`

	// Optional. IANA Time Zone Database version number, e.g. "2019a".
	// +kcc:proto:field=google.type.TimeZone.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.PatchDeployment
type PatchDeploymentObservedState struct {
	// Required. Schedule recurring executions.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchDeployment.recurring_schedule
	RecurringSchedule *RecurringScheduleObservedState `json:"recurringSchedule,omitempty"`

	// Output only. Time the patch deployment was created. Timestamp is in
	//  [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchDeployment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time the patch deployment was last updated. Timestamp is in
	//  [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchDeployment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The last time a patch job was started by this deployment.
	//  Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text
	//  format.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchDeployment.last_execute_time
	LastExecuteTime *string `json:"lastExecuteTime,omitempty"`

	// Output only. Current state of the patch deployment.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchDeployment.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.RecurringSchedule
type RecurringScheduleObservedState struct {
	// Output only. The time the last patch job ran successfully.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.RecurringSchedule.last_execute_time
	LastExecuteTime *string `json:"lastExecuteTime,omitempty"`

	// Output only. The time the next patch job is scheduled to run.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.RecurringSchedule.next_execute_time
	NextExecuteTime *string `json:"nextExecuteTime,omitempty"`
}
