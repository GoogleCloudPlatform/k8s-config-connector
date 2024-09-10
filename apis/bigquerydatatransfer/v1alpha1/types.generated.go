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

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.UserInfo
type UserInfo struct {
	// E-mail address of the user.
	Email *string `json:"email,omitempty"`
}
