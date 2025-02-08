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


// +kcc:proto=google.cloud.storageinsights.v1.ReportDetail
type ReportDetail struct {
	// Name of resource. It will be of form
	//  projects/<project>/locations/<location>/reportConfigs/<report-config-id>/reportDetails/<report-detail-id>.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportDetail.name
	Name *string `json:"name,omitempty"`

	// The snapshot time.
	//  All the report data is referenced at this point of time.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportDetail.snapshot_time
	SnapshotTime *string `json:"snapshotTime,omitempty"`

	// Prefix of the object name of each report's shard. This will have full
	//  prefix except the "extension" and "shard_id".
	//  For example, if the `destination_path` is
	//  `{{report-config-id}}/dt={{datetime}}`, the shard object name would be
	//  `gs://my-insights/1A34-F2E456-12B456-1C3D/dt=2022-05-20T06:35/1A34-F2E456-12B456-1C3D_2022-05-20T06:35_5.csv`
	//  and the value of `report_path_prefix` field would be
	//  `gs://my-insights/1A34-F2E456-12B456-1C3D/dt=2022-05-20T06:35/1A34-F2E456-12B456-1C3D_2022-05-20T06:35_`.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportDetail.report_path_prefix
	ReportPathPrefix *string `json:"reportPathPrefix,omitempty"`

	// Total shards generated for the report.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportDetail.shards_count
	ShardsCount *int64 `json:"shardsCount,omitempty"`

	// Status of the ReportDetail.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportDetail.status
	Status *Status `json:"status,omitempty"`

	// Labels as key value pairs
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportDetail.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The date for which report is generated. The time part of target_datetime
	//  will be zero till we support multiple reports per day.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportDetail.target_datetime
	TargetDatetime *DateTime `json:"targetDatetime,omitempty"`

	// Metrics of the report.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportDetail.report_metrics
	ReportMetrics *ReportDetail_Metrics `json:"reportMetrics,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.ReportDetail.Metrics
type ReportDetail_Metrics struct {
	// Count of Cloud Storage objects which are part of the report.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportDetail.Metrics.processed_records_count
	ProcessedRecordsCount *int64 `json:"processedRecordsCount,omitempty"`
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

// +kcc:proto=google.type.DateTime
type DateTime struct {
	// Optional. Year of date. Must be from 1 to 9999, or 0 if specifying a
	//  datetime without a year.
	// +kcc:proto:field=google.type.DateTime.year
	Year *int32 `json:"year,omitempty"`

	// Required. Month of year. Must be from 1 to 12.
	// +kcc:proto:field=google.type.DateTime.month
	Month *int32 `json:"month,omitempty"`

	// Required. Day of month. Must be from 1 to 31 and valid for the year and
	//  month.
	// +kcc:proto:field=google.type.DateTime.day
	Day *int32 `json:"day,omitempty"`

	// Required. Hours of day in 24 hour format. Should be from 0 to 23. An API
	//  may choose to allow the value "24:00:00" for scenarios like business
	//  closing time.
	// +kcc:proto:field=google.type.DateTime.hours
	Hours *int32 `json:"hours,omitempty"`

	// Required. Minutes of hour of day. Must be from 0 to 59.
	// +kcc:proto:field=google.type.DateTime.minutes
	Minutes *int32 `json:"minutes,omitempty"`

	// Required. Seconds of minutes of the time. Must normally be from 0 to 59. An
	//  API may allow the value 60 if it allows leap-seconds.
	// +kcc:proto:field=google.type.DateTime.seconds
	Seconds *int32 `json:"seconds,omitempty"`

	// Required. Fractions of seconds in nanoseconds. Must be from 0 to
	//  999,999,999.
	// +kcc:proto:field=google.type.DateTime.nanos
	Nanos *int32 `json:"nanos,omitempty"`

	// UTC offset. Must be whole seconds, between -18 hours and +18 hours.
	//  For example, a UTC offset of -4:00 would be represented as
	//  { seconds: -14400 }.
	// +kcc:proto:field=google.type.DateTime.utc_offset
	UtcOffset *string `json:"utcOffset,omitempty"`

	// Time zone.
	// +kcc:proto:field=google.type.DateTime.time_zone
	TimeZone *TimeZone `json:"timeZone,omitempty"`
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
