// Copyright 2024 Google LLC
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
	EnableFailureEmail *bool `json:"enableFailureEmail,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.EncryptionConfiguration
type EncryptionConfiguration struct {
	// The name of the KMS key used for encrypting BigQuery data.
	KmsKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.ScheduleOptions
type ScheduleOptions struct {
	// If true, automatic scheduling of data transfer runs for this configuration
	//  will be disabled. The runs can be started on ad-hoc basis using
	//  StartManualTransferRuns API. When automatic scheduling is disabled, the
	//  TransferConfig.schedule field will be ignored.
	DisableAutoScheduling *bool `json:"disableAutoScheduling,omitempty"`

	// Specifies time to start scheduling transfer runs. The first run will be
	//  scheduled at or after the start time according to a recurrence pattern
	//  defined in the schedule string. The start time can be changed at any
	//  moment. The time when a data transfer can be triggered manually is not
	//  limited by this option.
	StartTime *string `json:"startTime,omitempty"`

	// Defines time to stop scheduling transfer runs. A transfer run cannot be
	//  scheduled at or after the end time. The end time can be changed at any
	//  moment. The time when a data transfer can be triggered manually is not
	//  limited by this option.
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
	Name *string `json:"name,omitempty"`

	// The BigQuery target dataset id.
	DestinationDatasetID *string `json:"destinationDatasetID,omitempty"`

	// User specified display name for the data transfer.
	DisplayName *string `json:"displayName,omitempty"`

	// Data source ID. This cannot be changed once data transfer is created. The
	//  full list of available data source IDs can be returned through an API call:
	//  https://cloud.google.com/bigquery-transfer/docs/reference/datatransfer/rest/v1/projects.locations.dataSources/list
	DataSourceID *string `json:"dataSourceID,omitempty"`

	// Parameters specific to each data source. For more information see the
	//  bq tab in the 'Setting up a data transfer' section for each data source.
	//  For example the parameters for Cloud Storage transfers are listed here:
	//  https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
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
	Schedule *string `json:"schedule,omitempty"`

	// Options customizing the data transfer schedule.
	ScheduleOptions *ScheduleOptions `json:"scheduleOptions,omitempty"`

	// The number of days to look back to automatically refresh the data.
	//  For example, if `data_refresh_window_days = 10`, then every day
	//  BigQuery reingests data for [today-10, today-1], rather than ingesting data
	//  for just [today-1].
	//  Only valid if the data source supports the feature. Set the value to 0
	//  to use the default value.
	DataRefreshWindowDays *int32 `json:"dataRefreshWindowDays,omitempty"`

	// Is this config disabled. When set to true, no runs will be scheduled for
	//  this transfer config.
	Disabled *bool `json:"disabled,omitempty"`

	// Output only. Data transfer modification time. Ignored by server on input.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Next time when data transfer will run.
	NextRunTime *string `json:"nextRunTime,omitempty"`

	// Output only. State of the most recently updated transfer run.
	State *string `json:"state,omitempty"`

	// Deprecated. Unique ID of the user on whose behalf transfer is done.
	UserID *int64 `json:"userID,omitempty"`

	// Output only. Region in which BigQuery dataset is located.
	DatasetRegion *string `json:"datasetRegion,omitempty"`

	// Pub/Sub topic where notifications will be sent after transfer runs
	//  associated with this transfer config finish.
	//
	//  The format for specifying a pubsub topic is:
	//  `projects/{project_id}/topics/{topic_id}`
	NotificationPubsubTopic *string `json:"notificationPubsubTopic,omitempty"`

	// Email notifications will be sent according to these preferences
	//  to the email address of the user who owns this transfer config.
	EmailPreferences *EmailPreferences `json:"emailPreferences,omitempty"`

	// Output only. Information about the user whose credentials are used to
	//  transfer data. Populated only for `transferConfigs.get` requests. In case
	//  the user information is not available, this field will not be populated.
	OwnerInfo *UserInfo `json:"ownerInfo,omitempty"`

	// The encryption configuration part. Currently, it is only used for the
	//  optional KMS key name. The BigQuery service account of your project must be
	//  granted permissions to use the key. Read methods will return the key name
	//  applied in effect. Write methods will apply the key if it is present, or
	//  otherwise try to apply project default keys if it is absent.
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.UserInfo
type UserInfo struct {
	// E-mail address of the user.
	Email *string `json:"email,omitempty"`
}
