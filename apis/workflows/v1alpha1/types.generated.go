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


// +kcc:proto=google.cloud.workflows.executions.v1beta.Execution
type Execution struct {

	// Input parameters of the execution represented as a JSON string.
	//  The size limit is 32KB.
	// +kcc:proto:field=google.cloud.workflows.executions.v1beta.Execution.argument
	Argument *string `json:"argument,omitempty"`
}

// +kcc:proto=google.cloud.workflows.executions.v1beta.Execution.Error
type Execution_Error struct {
	// Error payload returned by the execution, represented as a JSON string.
	// +kcc:proto:field=google.cloud.workflows.executions.v1beta.Execution.Error.payload
	Payload *string `json:"payload,omitempty"`

	// Human readable error context, helpful for debugging purposes.
	// +kcc:proto:field=google.cloud.workflows.executions.v1beta.Execution.Error.context
	Context *string `json:"context,omitempty"`
}

// +kcc:proto=google.cloud.workflows.executions.v1beta.Execution
type ExecutionObservedState struct {
	// Output only. The resource name of the execution.
	//  Format:
	//  projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
	// +kcc:proto:field=google.cloud.workflows.executions.v1beta.Execution.name
	Name *string `json:"name,omitempty"`

	// Output only. Marks the beginning of execution.
	// +kcc:proto:field=google.cloud.workflows.executions.v1beta.Execution.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Marks the end of execution, successful or not.
	// +kcc:proto:field=google.cloud.workflows.executions.v1beta.Execution.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Current state of the execution.
	// +kcc:proto:field=google.cloud.workflows.executions.v1beta.Execution.state
	State *string `json:"state,omitempty"`

	// Output only. Output of the execution represented as a JSON string. The
	//  value can only be present if the execution's state is `SUCCEEDED`.
	// +kcc:proto:field=google.cloud.workflows.executions.v1beta.Execution.result
	Result *string `json:"result,omitempty"`

	// Output only. The error which caused the execution to finish prematurely.
	//  The value is only present if the execution's state is `FAILED`
	//  or `CANCELLED`.
	// +kcc:proto:field=google.cloud.workflows.executions.v1beta.Execution.error
	Error *Execution_Error `json:"error,omitempty"`

	// Output only. Revision of the workflow this execution is using.
	// +kcc:proto:field=google.cloud.workflows.executions.v1beta.Execution.workflow_revision_id
	WorkflowRevisionID *string `json:"workflowRevisionID,omitempty"`
}
