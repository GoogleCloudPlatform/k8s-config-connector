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

// +kcc:proto=google.cloud.workflows.executions.v1.Execution.Error
type Execution_Error struct {
	// Error message and data returned represented as a JSON string.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.Error.payload
	Payload *string `json:"payload,omitempty"`

	// Human-readable stack trace string.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.Error.context
	Context *string `json:"context,omitempty"`

	// Stack trace with detailed information of where error was generated.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.Error.stack_trace
	StackTrace *Execution_StackTrace `json:"stackTrace,omitempty"`
}

// +kcc:proto=google.cloud.workflows.executions.v1.Execution.StackTrace
type Execution_StackTrace struct {
	// An array of stack elements.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.StackTrace.elements
	Elements []Execution_StackTraceElement `json:"elements,omitempty"`
}

// +kcc:proto=google.cloud.workflows.executions.v1.Execution.StackTraceElement
type Execution_StackTraceElement struct {
	// The step the error occurred at.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.StackTraceElement.step
	Step *string `json:"step,omitempty"`

	// The routine where the error occurred.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.StackTraceElement.routine
	Routine *string `json:"routine,omitempty"`

	// The source position information of the stack trace element.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.StackTraceElement.position
	Position *Execution_StackTraceElement_Position `json:"position,omitempty"`
}

// +kcc:proto=google.cloud.workflows.executions.v1.Execution.StackTraceElement.Position
type Execution_StackTraceElement_Position struct {
	// The source code line number the current instruction was generated from.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.StackTraceElement.Position.line
	Line *int64 `json:"line,omitempty"`

	// The source code column position (of the line) the current instruction
	//  was generated from.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.StackTraceElement.Position.column
	Column *int64 `json:"column,omitempty"`

	// The number of bytes of source code making up this stack trace element.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.StackTraceElement.Position.length
	Length *int64 `json:"length,omitempty"`
}

// +kcc:proto=google.cloud.workflows.executions.v1.Execution.StateError
type Execution_StateError struct {
	// Provides specifics about the error.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.StateError.details
	Details *string `json:"details,omitempty"`

	// The type of this state error.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.StateError.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.workflows.executions.v1.Execution.Status
type Execution_Status struct {
	// A list of currently executing or last executed step names for the
	//  workflow execution currently running. If the workflow has succeeded or
	//  failed, this is the last attempted or executed step. Presently, if the
	//  current step is inside a subworkflow, the list only includes that step.
	//  In the future, the list will contain items for each step in the call
	//  stack, starting with the outermost step in the `main` subworkflow, and
	//  ending with the most deeply nested step.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.Status.current_steps
	CurrentSteps []Execution_Status_Step `json:"currentSteps,omitempty"`
}

// +kcc:proto=google.cloud.workflows.executions.v1.Execution.Status.Step
type Execution_Status_Step struct {
	// Name of a routine within the workflow.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.Status.Step.routine
	Routine *string `json:"routine,omitempty"`

	// Name of a step within the routine.
	// +kcc:proto:field=google.cloud.workflows.executions.v1.Execution.Status.Step.step
	Step *string `json:"step,omitempty"`
}
