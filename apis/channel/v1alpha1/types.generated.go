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


// +kcc:proto=google.cloud.channel.v1.ReportJob
type ReportJob struct {
	// Required. The resource name of a report job.
	//  Name uses the format:
	//  `accounts/{account_id}/reportJobs/{report_job_id}`
	// +kcc:proto:field=google.cloud.channel.v1.ReportJob.name
	Name *string `json:"name,omitempty"`

	// The current status of report generation.
	// +kcc:proto:field=google.cloud.channel.v1.ReportJob.report_status
	ReportStatus *ReportStatus `json:"reportStatus,omitempty"`
}

// +kcc:proto=google.cloud.channel.v1.ReportStatus
type ReportStatus struct {
	// The current state of the report generation process.
	// +kcc:proto:field=google.cloud.channel.v1.ReportStatus.state
	State *string `json:"state,omitempty"`

	// The report generation's start time.
	// +kcc:proto:field=google.cloud.channel.v1.ReportStatus.start_time
	StartTime *string `json:"startTime,omitempty"`

	// The report generation's completion time.
	// +kcc:proto:field=google.cloud.channel.v1.ReportStatus.end_time
	EndTime *string `json:"endTime,omitempty"`
}
