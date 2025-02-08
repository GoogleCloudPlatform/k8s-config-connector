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


// +kcc:proto=google.cloud.storageinsights.v1.CSVOptions
type CSVOptions struct {
	// Record separator characters in CSV.
	// +kcc:proto:field=google.cloud.storageinsights.v1.CSVOptions.record_separator
	RecordSeparator *string `json:"recordSeparator,omitempty"`

	// Delimiter characters in CSV.
	// +kcc:proto:field=google.cloud.storageinsights.v1.CSVOptions.delimiter
	Delimiter *string `json:"delimiter,omitempty"`

	// If set, will include a header row in the CSV report.
	// +kcc:proto:field=google.cloud.storageinsights.v1.CSVOptions.header_required
	HeaderRequired *bool `json:"headerRequired,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.CloudStorageDestinationOptions
type CloudStorageDestinationOptions struct {
	// Destination bucket.
	// +kcc:proto:field=google.cloud.storageinsights.v1.CloudStorageDestinationOptions.bucket
	Bucket *string `json:"bucket,omitempty"`

	// Destination path is the path in the bucket where the report should be
	//  generated.
	// +kcc:proto:field=google.cloud.storageinsights.v1.CloudStorageDestinationOptions.destination_path
	DestinationPath *string `json:"destinationPath,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.CloudStorageFilters
type CloudStorageFilters struct {
	// Bucket for which the report will be generated.
	// +kcc:proto:field=google.cloud.storageinsights.v1.CloudStorageFilters.bucket
	Bucket *string `json:"bucket,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.FrequencyOptions
type FrequencyOptions struct {
	// Frequency of report generation.
	// +kcc:proto:field=google.cloud.storageinsights.v1.FrequencyOptions.frequency
	Frequency *string `json:"frequency,omitempty"`

	// The date from which report generation should start.
	//  UTC time zone.
	// +kcc:proto:field=google.cloud.storageinsights.v1.FrequencyOptions.start_date
	StartDate *Date `json:"startDate,omitempty"`

	// The date on which report generation should stop (Inclusive).
	//  UTC time zone.
	// +kcc:proto:field=google.cloud.storageinsights.v1.FrequencyOptions.end_date
	EndDate *Date `json:"endDate,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.ObjectMetadataReportOptions
type ObjectMetadataReportOptions struct {
	// Metadata fields to be included in the report.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ObjectMetadataReportOptions.metadata_fields
	MetadataFields []string `json:"metadataFields,omitempty"`

	// Cloud Storage as the storage system.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ObjectMetadataReportOptions.storage_filters
	StorageFilters *CloudStorageFilters `json:"storageFilters,omitempty"`

	// Cloud Storage as the storage system.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ObjectMetadataReportOptions.storage_destination_options
	StorageDestinationOptions *CloudStorageDestinationOptions `json:"storageDestinationOptions,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.ParquetOptions
type ParquetOptions struct {
}

// +kcc:proto=google.cloud.storageinsights.v1.ReportConfig
type ReportConfig struct {
	// name of resource. It will be of form
	//  projects/<project>/locations/<location>/reportConfigs/<report-config-id>.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.name
	Name *string `json:"name,omitempty"`

	// The frequency of report generation.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.frequency_options
	FrequencyOptions *FrequencyOptions `json:"frequencyOptions,omitempty"`

	// Options for CSV formatted reports.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.csv_options
	CsvOptions *CSVOptions `json:"csvOptions,omitempty"`

	// Options for Parquet formatted reports.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.parquet_options
	ParquetOptions *ParquetOptions `json:"parquetOptions,omitempty"`

	// Report for exporting object metadata.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.object_metadata_report_options
	ObjectMetadataReportOptions *ObjectMetadataReportOptions `json:"objectMetadataReportOptions,omitempty"`

	// Labels as key value pairs
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.labels
	Labels map[string]string `json:"labels,omitempty"`

	// User provided display name which can be empty and limited to 256 characters
	//  that is editable.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`
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

// +kcc:proto=google.cloud.storageinsights.v1.ReportConfig
type ReportConfigObservedState struct {
	// Output only. [Output only] Create time stamp
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
