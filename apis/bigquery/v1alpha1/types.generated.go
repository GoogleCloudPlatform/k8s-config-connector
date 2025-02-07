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

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.TransferRun
type TransferRun struct {
	// Identifier. The resource name of the transfer run.
	//  Transfer run names have the form
	//  `projects/{project_id}/locations/{location}/transferConfigs/{config_id}/runs/{run_id}`.
	//  The name is ignored when creating a transfer run.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.name
	Name *string `json:"name,omitempty"`

	// Minimum time after which a transfer run can be started.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.schedule_time
	ScheduleTime *string `json:"scheduleTime,omitempty"`

	// For batch transfer runs, specifies the date and time of the data should be
	//  ingested.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.run_time
	RunTime *string `json:"runTime,omitempty"`

	// Status of the transfer run.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.error_status
	ErrorStatus *Status `json:"errorStatus,omitempty"`

	// Data transfer run state. Ignored for input requests.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.state
	State *string `json:"state,omitempty"`

	// Deprecated. Unique ID of the user on whose behalf transfer is done.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.user_id
	UserID *int64 `json:"userID,omitempty"`
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

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.TransferRun
type TransferRunObservedState struct {
	// Output only. Time when transfer run was started.
	//  Parameter ignored by server for input requests.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when transfer run ended.
	//  Parameter ignored by server for input requests.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Last time the data transfer run state was updated.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Parameters specific to each data source. For more information
	//  see the bq tab in the 'Setting up a data transfer' section for each data
	//  source. For example the parameters for Cloud Storage transfers are listed
	//  here:
	//  https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.params
	Params map[string]string `json:"params,omitempty"`

	// Output only. The BigQuery target dataset id.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.destination_dataset_id
	DestinationDatasetID *string `json:"destinationDatasetID,omitempty"`

	// Output only. Data source id.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.data_source_id
	DataSourceID *string `json:"dataSourceID,omitempty"`

	// Output only. Describes the schedule of this transfer run if it was
	//  created as part of a regular schedule. For batch transfer runs that are
	//  scheduled manually, this is empty.
	//  NOTE: the system might choose to delay the schedule depending on the
	//  current load, so `schedule_time` doesn't always match this.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.schedule
	Schedule *string `json:"schedule,omitempty"`

	// Output only. Pub/Sub topic where a notification will be sent after this
	//  transfer run finishes.
	//
	//  The format for specifying a pubsub topic is:
	//  `projects/{project_id}/topics/{topic_id}`
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.notification_pubsub_topic
	NotificationPubsubTopic *string `json:"notificationPubsubTopic,omitempty"`

	// Output only. Email notifications will be sent according to these
	//  preferences to the email address of the user who owns the transfer config
	//  this run was derived from.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.TransferRun.email_preferences
	EmailPreferences *EmailPreferences `json:"emailPreferences,omitempty"`
}
