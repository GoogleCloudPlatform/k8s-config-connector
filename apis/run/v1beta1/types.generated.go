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

// +generated:types
// krm.group: run.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.run.v2
// resource: RunJob:Job

package v1beta1

// +kcc:proto=google.cloud.run.v2.BuildInfo
type BuildInfo struct {
}

// +kcc:proto=google.cloud.run.v2.Condition
type Condition struct {
	// type is used to communicate the status of the reconciliation process.
	//  See also:
	//  https://github.com/knative/serving/blob/main/docs/spec/errors.md#error-conditions-and-reporting
	//  Types common to all resources include:
	//  * "Ready": True when the Resource is ready.
	// +kcc:proto:field=google.cloud.run.v2.Condition.type
	Type *string `json:"type,omitempty"`

	// State of the condition.
	// +kcc:proto:field=google.cloud.run.v2.Condition.state
	State *string `json:"state,omitempty"`

	// Human readable message indicating details about the current status.
	// +kcc:proto:field=google.cloud.run.v2.Condition.message
	Message *string `json:"message,omitempty"`

	// Last time the condition transitioned from one status to another.
	// +kcc:proto:field=google.cloud.run.v2.Condition.last_transition_time
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`

	// How to interpret failures of this condition, one of Error, Warning, Info
	// +kcc:proto:field=google.cloud.run.v2.Condition.severity
	Severity *string `json:"severity,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.run.v2.BuildInfo
type BuildInfoObservedState struct {
	// Output only. Entry point of the function when the image is a Cloud Run
	//  function.
	// +kcc:proto:field=google.cloud.run.v2.BuildInfo.function_target
	FunctionTarget *string `json:"functionTarget,omitempty"`

	// Output only. Source code location of the image.
	// +kcc:proto:field=google.cloud.run.v2.BuildInfo.source_location
	SourceLocation *string `json:"sourceLocation,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.run.v2.Condition
type ConditionObservedState struct {
	// Output only. A common (service-level) reason for this condition.
	// +kcc:proto:field=google.cloud.run.v2.Condition.reason
	Reason *string `json:"reason,omitempty"`

	// Output only. A reason for the revision condition.
	// +kcc:proto:field=google.cloud.run.v2.Condition.revision_reason
	RevisionReason *string `json:"revisionReason,omitempty"`

	// Output only. A reason for the execution condition.
	// +kcc:proto:field=google.cloud.run.v2.Condition.execution_reason
	ExecutionReason *string `json:"executionReason,omitempty"`
}
