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


// +kcc:proto=google.cloud.bigquery.reservation.v1.Assignment
type Assignment struct {

	// The resource which will use the reservation. E.g.
	//  `projects/myproject`, `folders/123`, or `organizations/456`.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Assignment.assignee
	Assignee *string `json:"assignee,omitempty"`

	// Which type of jobs will use the reservation.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Assignment.job_type
	JobType *string `json:"jobType,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.reservation.v1.Assignment
type AssignmentObservedState struct {
	// Output only. Name of the resource. E.g.:
	//  `projects/myproject/locations/US/reservations/team1-prod/assignments/123`.
	//  The assignment_id must only contain lower case alphanumeric characters or
	//  dashes and the max length is 64 characters.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Assignment.name
	Name *string `json:"name,omitempty"`

	// Output only. State of the assignment.
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.Assignment.state
	State *string `json:"state,omitempty"`
}
