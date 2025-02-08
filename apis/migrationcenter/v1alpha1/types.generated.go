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


// +kcc:proto=google.cloud.migrationcenter.v1.ExecutionReport
type ExecutionReport struct {
	// Total number of asset frames reported for the import job.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ExecutionReport.frames_reported
	FramesReported *int32 `json:"framesReported,omitempty"`

	// Validation errors encountered during the execution of the import job.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ExecutionReport.execution_errors
	ExecutionErrors *ValidationReport `json:"executionErrors,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.FileValidationReport
type FileValidationReport struct {
	// The name of the file.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.FileValidationReport.file_name
	FileName *string `json:"fileName,omitempty"`

	// Partial list of rows that encountered validation error.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.FileValidationReport.row_errors
	RowErrors []ImportRowError `json:"rowErrors,omitempty"`

	// Flag indicating that processing was aborted due to maximum number of
	//  errors.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.FileValidationReport.partial_report
	PartialReport *bool `json:"partialReport,omitempty"`

	// List of file level errors.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.FileValidationReport.file_errors
	FileErrors []ImportError `json:"fileErrors,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ImportError
type ImportError struct {
	// The error information.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportError.error_details
	ErrorDetails *string `json:"errorDetails,omitempty"`

	// The severity of the error.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportError.severity
	Severity *string `json:"severity,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ImportJob
type ImportJob struct {

	// User-friendly display name. Maximum length is 63 characters.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Reference to a source.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.asset_source
	AssetSource *string `json:"assetSource,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ImportRowError
type ImportRowError struct {
	// The row number where the error was detected.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportRowError.row_number
	RowNumber *int32 `json:"rowNumber,omitempty"`

	// The name of the VM in the row.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportRowError.vm_name
	VmName *string `json:"vmName,omitempty"`

	// The VM UUID.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportRowError.vm_uuid
	VmUuid *string `json:"vmUuid,omitempty"`

	// The list of errors detected in the row.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportRowError.errors
	Errors []ImportError `json:"errors,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ValidationReport
type ValidationReport struct {
	// List of errors found in files.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ValidationReport.file_validations
	FileValidations []FileValidationReport `json:"fileValidations,omitempty"`

	// List of job level errors.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ValidationReport.job_errors
	JobErrors []ImportError `json:"jobErrors,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ExecutionReport
type ExecutionReportObservedState struct {
	// Output only. Total number of rows in the import job.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ExecutionReport.total_rows_count
	TotalRowsCount *int32 `json:"totalRowsCount,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ImportJob
type ImportJobObservedState struct {
	// Output only. The full name of the import job.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.name
	Name *string `json:"name,omitempty"`

	// Output only. The timestamp when the import job was created.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the import job was last updated.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The timestamp when the import job was completed.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.complete_time
	CompleteTime *string `json:"completeTime,omitempty"`

	// Output only. The state of the import job.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.state
	State *string `json:"state,omitempty"`

	// Output only. The report with the validation results of the import job.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.validation_report
	ValidationReport *ValidationReport `json:"validationReport,omitempty"`

	// Output only. The report with the results of running the import job.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.execution_report
	ExecutionReport *ExecutionReport `json:"executionReport,omitempty"`
}
