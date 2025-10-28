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

// +generated:types
// krm.group: clouddeploy.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.deploy.v1
// resource: DeployCustomTargetType:CustomTargetType
// resource: CloudDeployDeployPolicy:DeployPolicy

package v1alpha1

// +kcc:proto=google.cloud.deploy.v1.CustomTargetSkaffoldActions
type CustomTargetSkaffoldActions struct {
	// Optional. The Skaffold custom action responsible for render operations. If
	//  not provided then Cloud Deploy will perform the render operations via
	//  `skaffold render`.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetSkaffoldActions.render_action
	RenderAction *string `json:"renderAction,omitempty"`

	// Required. The Skaffold custom action responsible for deploy operations.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetSkaffoldActions.deploy_action
	DeployAction *string `json:"deployAction,omitempty"`

	// Optional. List of Skaffold modules Cloud Deploy will include in the
	//  Skaffold Config as required before performing diagnose.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTargetSkaffoldActions.include_skaffold_modules
	IncludeSkaffoldModules []SkaffoldModules `json:"includeSkaffoldModules,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.DeliveryPipelineAttribute
type DeliveryPipelineAttribute struct {
	// Optional. ID of the `DeliveryPipeline`. The value of this field could be
	//  one of the following:
	//
	//  * The last segment of a pipeline name
	//  * "*", all delivery pipelines in a location
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipelineAttribute.id
	ID *string `json:"id,omitempty"`

	// DeliveryPipeline labels.
	// +kcc:proto:field=google.cloud.deploy.v1.DeliveryPipelineAttribute.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.DeployPolicyResourceSelector
type DeployPolicyResourceSelector struct {
	// Optional. Contains attributes about a delivery pipeline.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicyResourceSelector.delivery_pipeline
	DeliveryPipeline *DeliveryPipelineAttribute `json:"deliveryPipeline,omitempty"`

	// Optional. Contains attributes about a target.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicyResourceSelector.target
	Target *TargetAttribute `json:"target,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.OneTimeWindow
type OneTimeWindow struct {
	// Required. Start date.
	// +kcc:proto:field=google.cloud.deploy.v1.OneTimeWindow.start_date
	StartDate *Date `json:"startDate,omitempty"`

	// Required. Start time (inclusive). Use 00:00 for the beginning of the day.
	// +kcc:proto:field=google.cloud.deploy.v1.OneTimeWindow.start_time
	StartTime *TimeOfDay `json:"startTime,omitempty"`

	// Required. End date.
	// +kcc:proto:field=google.cloud.deploy.v1.OneTimeWindow.end_date
	EndDate *Date `json:"endDate,omitempty"`

	// Required. End time (exclusive). You may use 24:00 for the end of the day.
	// +kcc:proto:field=google.cloud.deploy.v1.OneTimeWindow.end_time
	EndTime *TimeOfDay `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PolicyRule
type PolicyRule struct {
	// Optional. Rollout restrictions.
	// +kcc:proto:field=google.cloud.deploy.v1.PolicyRule.rollout_restriction
	RolloutRestriction *RolloutRestriction `json:"rolloutRestriction,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.RolloutRestriction
type RolloutRestriction struct {
	// Required. Restriction rule ID. Required and must be unique within a
	//  DeployPolicy. The format is `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
	// +kcc:proto:field=google.cloud.deploy.v1.RolloutRestriction.id
	ID *string `json:"id,omitempty"`

	// Optional. What invoked the action. If left empty, all invoker types will be
	//  restricted.
	// +kcc:proto:field=google.cloud.deploy.v1.RolloutRestriction.invokers
	Invokers []string `json:"invokers,omitempty"`

	// Optional. Rollout actions to be restricted as part of the policy. If left
	//  empty, all actions will be restricted.
	// +kcc:proto:field=google.cloud.deploy.v1.RolloutRestriction.actions
	Actions []string `json:"actions,omitempty"`

	// Required. Time window within which actions are restricted.
	// +kcc:proto:field=google.cloud.deploy.v1.RolloutRestriction.time_windows
	TimeWindows *TimeWindows `json:"timeWindows,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SkaffoldModules
type SkaffoldModules struct {
	// Optional. The Skaffold Config modules to use from the specified source.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.configs
	Configs []string `json:"configs,omitempty"`

	// Optional. Remote git repository containing the Skaffold Config modules.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.git
	Git *SkaffoldModules_SkaffoldGitSource `json:"git,omitempty"`

	// Optional. Cloud Storage bucket containing the Skaffold Config modules.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.google_cloud_storage
	GoogleCloudStorage *SkaffoldModules_SkaffoldGCSSource `json:"googleCloudStorage,omitempty"`

	// Optional. Cloud Build V2 repository containing the Skaffold Config
	//  modules.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.google_cloud_build_repo
	GoogleCloudBuildRepo *SkaffoldModules_SkaffoldGcbRepoSource `json:"googleCloudBuildRepo,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCSSource
type SkaffoldModules_SkaffoldGCSSource struct {
	// Required. Cloud Storage source paths to copy recursively. For example,
	//  providing "gs://my-bucket/dir/configs/*" will result in Skaffold copying
	//  all files within the "dir/configs" directory in the bucket "my-bucket".
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCSSource.source
	Source *string `json:"source,omitempty"`

	// Optional. Relative path from the source to the Skaffold file.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGCSSource.path
	Path *string `json:"path,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGitSource
type SkaffoldModules_SkaffoldGitSource struct {
	// Required. Git repository the package should be cloned from.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGitSource.repo
	Repo *string `json:"repo,omitempty"`

	// Optional. Relative path from the repository root to the Skaffold file.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGitSource.path
	Path *string `json:"path,omitempty"`

	// Optional. Git branch or tag to use when cloning the repository.
	// +kcc:proto:field=google.cloud.deploy.v1.SkaffoldModules.SkaffoldGitSource.ref
	Ref *string `json:"ref,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TargetAttribute
type TargetAttribute struct {
	// Optional. ID of the `Target`. The value of this field could be one of the
	//  following:
	//
	//  * The last segment of a target name
	//  * "*", all targets in a location
	// +kcc:proto:field=google.cloud.deploy.v1.TargetAttribute.id
	ID *string `json:"id,omitempty"`

	// Target labels.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetAttribute.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.TimeWindows
type TimeWindows struct {
	// Required. The time zone in IANA format [IANA Time Zone
	//  Database](https://www.iana.org/time-zones) (e.g. America/New_York).
	// +kcc:proto:field=google.cloud.deploy.v1.TimeWindows.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Optional. One-time windows within which actions are restricted.
	// +kcc:proto:field=google.cloud.deploy.v1.TimeWindows.one_time_windows
	OneTimeWindows []OneTimeWindow `json:"oneTimeWindows,omitempty"`

	// Optional. Recurring weekly windows within which actions are restricted.
	// +kcc:proto:field=google.cloud.deploy.v1.TimeWindows.weekly_windows
	WeeklyWindows []WeeklyWindow `json:"weeklyWindows,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.WeeklyWindow
type WeeklyWindow struct {
	// Optional. Days of week. If left empty, all days of the week will be
	//  included.
	// +kcc:proto:field=google.cloud.deploy.v1.WeeklyWindow.days_of_week
	DaysOfWeek []string `json:"daysOfWeek,omitempty"`

	// Optional. Start time (inclusive). Use 00:00 for the beginning of the day.
	//  If you specify start_time you must also specify end_time. If left empty,
	//  this will block for the entire day for the days specified in days_of_week.
	// +kcc:proto:field=google.cloud.deploy.v1.WeeklyWindow.start_time
	StartTime *TimeOfDay `json:"startTime,omitempty"`

	// Optional. End time (exclusive). Use 24:00 to indicate midnight. If you
	//  specify end_time you must also specify start_time. If left empty, this will
	//  block for the entire day for the days specified in days_of_week.
	// +kcc:proto:field=google.cloud.deploy.v1.WeeklyWindow.end_time
	EndTime *TimeOfDay `json:"endTime,omitempty"`
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
