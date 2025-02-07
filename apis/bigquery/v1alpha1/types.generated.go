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


// +kcc:proto=google.cloud.bigquery.datatransfer.v1.EmailPreferences
type EmailPreferences struct {
	// If true, email notifications will be sent on transfer run failures.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.EmailPreferences.enable_failure_email
	EnableFailureEmail *bool `json:"enableFailureEmail,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.EncryptionConfiguration
type EncryptionConfiguration struct {
	// The name of the KMS key used for encrypting BigQuery data.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.EncryptionConfiguration.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.EventDrivenSchedule
type EventDrivenSchedule struct {
	// Pub/Sub subscription name used to receive events.
	//  Only Google Cloud Storage data source support this option.
	//  Format: projects/{project}/subscriptions/{subscription}
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.EventDrivenSchedule.pubsub_subscription
	PubsubSubscription *string `json:"pubsubSubscription,omitempty"`
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

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.TransferConfig
type TransferConfig struct {
	// Identifier. The resource name of the transfer config.
	//  Transfer config names have the form either
	//  `projects/{project_id}/locations/{region}/transferConfigs/{config_id}` or
	//  `projects/{project_id}/transferConfigs/{config_id}`,
	//  where `config_id` is usually a UUID, even though it is not
	//  guaranteed or required. The name is ignored when creating a transfer
	//  config.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.name
	Name *string `json:"name,omitempty"`

	// The BigQuery target dataset id.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.destination_dataset_id
	DestinationDatasetID *string `json:"destinationDatasetID,omitempty"`

	// User specified display name for the data transfer.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Data source ID. This cannot be changed once data transfer is created. The
	//  full list of available data source IDs can be returned through an API call:
	//  https://cloud.google.com/bigquery-transfer/docs/reference/datatransfer/rest/v1/projects.locations.dataSources/list
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.data_source_id
	DataSourceID *string `json:"dataSourceID,omitempty"`

	// Parameters specific to each data source. For more information see the
	//  bq tab in the 'Setting up a data transfer' section for each data source.
	//  For example the parameters for Cloud Storage transfers are listed here:
	//  https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.params
	Params map[string]string `json:"params,omitempty"`

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
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.schedule
	Schedule *string `json:"schedule,omitempty"`

	// Options customizing the data transfer schedule.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.schedule_options
	ScheduleOptions *ScheduleOptions `json:"scheduleOptions,omitempty"`

	// Options customizing different types of data transfer schedule.
	//  This field replaces "schedule" and "schedule_options" fields.
	//  ScheduleOptionsV2 cannot be used together with ScheduleOptions/Schedule.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.schedule_options_v2
	ScheduleOptionsV2 *ScheduleOptionsV2 `json:"scheduleOptionsV2,omitempty"`

	// The number of days to look back to automatically refresh the data.
	//  For example, if `data_refresh_window_days = 10`, then every day
	//  BigQuery reingests data for [today-10, today-1], rather than ingesting data
	//  for just [today-1].
	//  Only valid if the data source supports the feature. Set the value to 0
	//  to use the default value.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.data_refresh_window_days
	DataRefreshWindowDays *int32 `json:"dataRefreshWindowDays,omitempty"`

	// Is this config disabled. When set to true, no runs will be scheduled for
	//  this transfer config.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Deprecated. Unique ID of the user on whose behalf transfer is done.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.user_id
	UserID *int64 `json:"userID,omitempty"`

	// Pub/Sub topic where notifications will be sent after transfer runs
	//  associated with this transfer config finish.
	//
	//  The format for specifying a pubsub topic is:
	//  `projects/{project_id}/topics/{topic_id}`
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.notification_pubsub_topic
	NotificationPubsubTopic *string `json:"notificationPubsubTopic,omitempty"`

	// Email notifications will be sent according to these preferences
	//  to the email address of the user who owns this transfer config.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.email_preferences
	EmailPreferences *EmailPreferences `json:"emailPreferences,omitempty"`

	// The encryption configuration part. Currently, it is only used for the
	//  optional KMS key name. The BigQuery service account of your project must be
	//  granted permissions to use the key. Read methods will return the key name
	//  applied in effect. Write methods will apply the key if it is present, or
	//  otherwise try to apply project default keys if it is absent.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.encryption_configuration
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`
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

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.TransferConfig
type TransferConfigObservedState struct {
	// Output only. Data transfer modification time. Ignored by server on input.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Next time when data transfer will run.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.next_run_time
	NextRunTime *string `json:"nextRunTime,omitempty"`

	// Output only. State of the most recently updated transfer run.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.state
	State *string `json:"state,omitempty"`

	// Output only. Region in which BigQuery dataset is located.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.dataset_region
	DatasetRegion *string `json:"datasetRegion,omitempty"`

	// Output only. Information about the user whose credentials are used to
	//  transfer data. Populated only for `transferConfigs.get` requests. In case
	//  the user information is not available, this field will not be populated.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.owner_info
	OwnerInfo *UserInfo `json:"ownerInfo,omitempty"`

	// Output only. Error code with detailed information about reason of the
	//  latest config failure.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferConfig.error
	Error *Status `json:"error,omitempty"`
}
