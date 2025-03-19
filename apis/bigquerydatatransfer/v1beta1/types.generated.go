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
// krm.group: bigquerydatatransfer.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.bigquery.datatransfer.v1
// resource: BigQueryDataTransferConfig:TransferConfig

package v1beta1

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.EmailPreferences
type EmailPreferences struct {
	// If true, email notifications will be sent on transfer run failures.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.EmailPreferences.enable_failure_email
	EnableFailureEmail *bool `json:"enableFailureEmail,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.ManualSchedule
type ManualSchedule struct {
}

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.ScheduleOptions
type ScheduleOptions struct {
	// If true, automatic scheduling of data transfer runs for this configuration
	//  will be disabled. The runs can be started on ad-hoc basis using
	//  StartManualTransferRuns API. When automatic scheduling is disabled, the
	//  TransferConfig.schedule field will be ignored.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.ScheduleOptions.disable_auto_scheduling
	DisableAutoScheduling *bool `json:"disableAutoScheduling,omitempty"`

	// Specifies time to start scheduling transfer runs. The first run will be
	//  scheduled at or after the start time according to a recurrence pattern
	//  defined in the schedule string. The start time can be changed at any
	//  moment. The time when a data transfer can be triggered manually is not
	//  limited by this option.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.ScheduleOptions.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Defines time to stop scheduling transfer runs. A transfer run cannot be
	//  scheduled at or after the end time. The end time can be changed at any
	//  moment. The time when a data transfer can be triggered manually is not
	//  limited by this option.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.ScheduleOptions.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.ScheduleOptionsV2
type ScheduleOptionsV2 struct {
	// Time based transfer schedule options. This is the default schedule
	//  option.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.ScheduleOptionsV2.time_based_schedule
	TimeBasedSchedule *TimeBasedSchedule `json:"timeBasedSchedule,omitempty"`

	// Manual transfer schedule. If set, the transfer run will not be
	//  auto-scheduled by the system, unless the client invokes
	//  StartManualTransferRuns.  This is equivalent to
	//  disable_auto_scheduling = true.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.ScheduleOptionsV2.manual_schedule
	ManualSchedule *ManualSchedule `json:"manualSchedule,omitempty"`

	// Event driven transfer schedule options. If set, the transfer will be
	//  scheduled upon events arrial.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.ScheduleOptionsV2.event_driven_schedule
	EventDrivenSchedule *EventDrivenSchedule `json:"eventDrivenSchedule,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.TimeBasedSchedule
type TimeBasedSchedule struct {
	// Data transfer schedule.
	//  If the data source does not support a custom schedule, this should be
	//  empty. If it is empty, the default value for the data source will be used.
	//  The specified times are in UTC.
	//  Examples of valid format:
	//  `1st,3rd monday of month 15:30`,
	//  `every wed,fri of jan,jun 13:15`, and
	//  `first sunday of quarter 00:00`.
	//  See more explanation about the format here:
	//  https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	//
	//  NOTE: The minimum interval time between recurring transfers depends on the
	//  data source; refer to the documentation for your data source.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TimeBasedSchedule.schedule
	Schedule *string `json:"schedule,omitempty"`

	// Specifies time to start scheduling transfer runs. The first run will be
	//  scheduled at or after the start time according to a recurrence pattern
	//  defined in the schedule string. The start time can be changed at any
	//  moment.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TimeBasedSchedule.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Defines time to stop scheduling transfer runs. A transfer run cannot be
	//  scheduled at or after the end time. The end time can be changed at any
	//  moment.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TimeBasedSchedule.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.UserInfo
type UserInfo struct {
	// E-mail address of the user.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.UserInfo.email
	Email *string `json:"email,omitempty"`
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
