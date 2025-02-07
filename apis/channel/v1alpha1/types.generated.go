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


// +kcc:proto=google.cloud.channel.v1.Column
type Column struct {
	// The unique name of the column (for example, customer_domain,
	//  channel_partner, customer_cost). You can use column IDs in
	//  [RunReportJobRequest.filter][google.cloud.channel.v1.RunReportJobRequest.filter].
	//  To see all reports and their columns, call
	//  [CloudChannelReportsService.ListReports][google.cloud.channel.v1.CloudChannelReportsService.ListReports].
	// +kcc:proto:field=google.cloud.channel.v1.Column.column_id
	ColumnID *string `json:"columnID,omitempty"`

	// The column's display name.
	// +kcc:proto:field=google.cloud.channel.v1.Column.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The type of the values for this column.
	// +kcc:proto:field=google.cloud.channel.v1.Column.data_type
	DataType *string `json:"dataType,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.Report
type Report struct {
	// Required. The report's resource name. Specifies the account and report used
	//  to generate report data. The report_id identifier is a UID (for example,
	//  `613bf59q`).
	//
	//  Name uses the format:
	//  accounts/{account_id}/reports/{report_id}
	// +kcc:proto:field=google.cloud.channel.v1.Report.name
	Name *string `json:"name,omitempty"`

	// A human-readable name for this report.
	// +kcc:proto:field=google.cloud.channel.v1.Report.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The list of columns included in the report. This defines the schema of
	//  the report results.
	// +kcc:proto:field=google.cloud.channel.v1.Report.columns
	Columns []Column `json:"columns,omitempty"`

	// A description of other aspects of the report, such as the products
	//  it supports.
	// +kcc:proto:field=google.cloud.channel.v1.Report.description
	Description *string `json:"description,omitempty"`
}
