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


// +kcc:proto=google.cloud.bigquery.datatransfer.v1.DataSource
type DataSource struct {

	// Data source id.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.data_source_id
	DataSourceID *string `json:"dataSourceID,omitempty"`

	// User friendly data source name.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// User friendly data source description string.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.description
	Description *string `json:"description,omitempty"`

	// Data source client id which should be used to receive refresh token.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.client_id
	ClientID *string `json:"clientID,omitempty"`

	// Api auth scopes for which refresh token needs to be obtained. These are
	//  scopes needed by a data source to prepare data and ingest them into
	//  BigQuery, e.g., https://www.googleapis.com/auth/bigquery
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.scopes
	Scopes []string `json:"scopes,omitempty"`

	// Deprecated. This field has no effect.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.transfer_type
	TransferType *string `json:"transferType,omitempty"`

	// Deprecated. This field has no effect.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.supports_multiple_transfers
	SupportsMultipleTransfers *bool `json:"supportsMultipleTransfers,omitempty"`

	// The number of seconds to wait for an update from the data source
	//  before the Data Transfer Service marks the transfer as FAILED.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.update_deadline_seconds
	UpdateDeadlineSeconds *int32 `json:"updateDeadlineSeconds,omitempty"`

	// Default data transfer schedule.
	//  Examples of valid schedules include:
	//  `1st,3rd monday of month 15:30`,
	//  `every wed,fri of jan,jun 13:15`, and
	//  `first sunday of quarter 00:00`.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.default_schedule
	DefaultSchedule *string `json:"defaultSchedule,omitempty"`

	// Specifies whether the data source supports a user defined schedule, or
	//  operates on the default schedule.
	//  When set to `true`, user can override default schedule.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.supports_custom_schedule
	SupportsCustomSchedule *bool `json:"supportsCustomSchedule,omitempty"`

	// Data source parameters.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.parameters
	Parameters []DataSourceParameter `json:"parameters,omitempty"`

	// Url for the help document for this data source.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.help_url
	HelpURL *string `json:"helpURL,omitempty"`

	// Indicates the type of authorization.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.authorization_type
	AuthorizationType *string `json:"authorizationType,omitempty"`

	// Specifies whether the data source supports automatic data refresh for the
	//  past few days, and how it's supported.
	//  For some data sources, data might not be complete until a few days later,
	//  so it's useful to refresh data automatically.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.data_refresh_type
	DataRefreshType *string `json:"dataRefreshType,omitempty"`

	// Default data refresh window on days.
	//  Only meaningful when `data_refresh_type` = `SLIDING_WINDOW`.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.default_data_refresh_window_days
	DefaultDataRefreshWindowDays *int32 `json:"defaultDataRefreshWindowDays,omitempty"`

	// Disables backfilling and manual run scheduling
	//  for the data source.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.manual_runs_disabled
	ManualRunsDisabled *bool `json:"manualRunsDisabled,omitempty"`

	// The minimum interval for scheduler to schedule runs.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.minimum_schedule_interval
	MinimumScheduleInterval *string `json:"minimumScheduleInterval,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.DataSourceParameter
type DataSourceParameter struct {
	// Parameter identifier.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.param_id
	ParamID *string `json:"paramID,omitempty"`

	// Parameter display name in the user interface.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Parameter description.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.description
	Description *string `json:"description,omitempty"`

	// Parameter type.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.type
	Type *string `json:"type,omitempty"`

	// Is parameter required.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.required
	Required *bool `json:"required,omitempty"`

	// Deprecated. This field has no effect.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.repeated
	Repeated *bool `json:"repeated,omitempty"`

	// Regular expression which can be used for parameter validation.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.validation_regex
	ValidationRegex *string `json:"validationRegex,omitempty"`

	// All possible values for the parameter.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.allowed_values
	AllowedValues []string `json:"allowedValues,omitempty"`

	// For integer and double values specifies minimum allowed value.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.min_value
	MinValue *DoubleValue `json:"minValue,omitempty"`

	// For integer and double values specifies maximum allowed value.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.max_value
	MaxValue *DoubleValue `json:"maxValue,omitempty"`

	// Deprecated. This field has no effect.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.fields
	Fields []DataSourceParameter `json:"fields,omitempty"`

	// Description of the requirements for this field, in case the user input does
	//  not fulfill the regex pattern or min/max values.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.validation_description
	ValidationDescription *string `json:"validationDescription,omitempty"`

	// URL to a help document to further explain the naming requirements.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.validation_help_url
	ValidationHelpURL *string `json:"validationHelpURL,omitempty"`

	// Cannot be changed after initial creation.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.immutable
	Immutable *bool `json:"immutable,omitempty"`

	// Deprecated. This field has no effect.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.recurse
	Recurse *bool `json:"recurse,omitempty"`

	// If true, it should not be used in new transfers, and it should not be
	//  visible to users.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSourceParameter.deprecated
	Deprecated *bool `json:"deprecated,omitempty"`
}

// +kcc:proto=google.protobuf.DoubleValue
type DoubleValue struct {
	// The double value.
	// +kcc:proto:field=google.protobuf.DoubleValue.value
	Value *float64 `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.DataSource
type DataSourceObservedState struct {
	// Output only. Data source resource name.
	// +kcc:proto:field=google.cloud.bigquery.datatransfer.v1.DataSource.name
	Name *string `json:"name,omitempty"`
}
