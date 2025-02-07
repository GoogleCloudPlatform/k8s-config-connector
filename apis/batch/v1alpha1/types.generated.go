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


// +kcc:proto=google.cloud.batch.v1.StatusEvent
type StatusEvent struct {
	// Type of the event.
	// +kcc:proto:field=google.cloud.batch.v1.StatusEvent.type
	Type *string `json:"type,omitempty"`

	// Description of the event.
	// +kcc:proto:field=google.cloud.batch.v1.StatusEvent.description
	Description *string `json:"description,omitempty"`

	// The time this event occurred.
	// +kcc:proto:field=google.cloud.batch.v1.StatusEvent.event_time
	EventTime *string `json:"eventTime,omitempty"`

	// Task Execution.
	//  This field is only defined for task-level status events where the task
	//  fails.
	// +kcc:proto:field=google.cloud.batch.v1.StatusEvent.task_execution
	TaskExecution *TaskExecution `json:"taskExecution,omitempty"`

	// Task State.
	//  This field is only defined for task-level status events.
	// +kcc:proto:field=google.cloud.batch.v1.StatusEvent.task_state
	TaskState *string `json:"taskState,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.Task
type Task struct {
	// Task name.
	//  The name is generated from the parent TaskGroup name and 'id' field.
	//  For example:
	//  "projects/123456/locations/us-west1/jobs/job01/taskGroups/group01/tasks/task01".
	// +kcc:proto:field=google.cloud.batch.v1.Task.name
	Name *string `json:"name,omitempty"`

	// Task Status.
	// +kcc:proto:field=google.cloud.batch.v1.Task.status
	Status *TaskStatus `json:"status,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.TaskExecution
type TaskExecution struct {
	// The exit code of a finished task.
	//
	//  If the task succeeded, the exit code will be 0. If the task failed but not
	//  due to the following reasons, the exit code will be 50000.
	//
	//  Otherwise, it can be from different sources:
	//  * Batch known failures:
	//  https://cloud.google.com/batch/docs/troubleshooting#reserved-exit-codes.
	//  * Batch runnable execution failures; you can rely on Batch logs to further
	//  diagnose: https://cloud.google.com/batch/docs/analyze-job-using-logs. If
	//  there are multiple runnables failures, Batch only exposes the first error.
	// +kcc:proto:field=google.cloud.batch.v1.TaskExecution.exit_code
	ExitCode *int32 `json:"exitCode,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1.TaskStatus
type TaskStatus struct {
	// Task state.
	// +kcc:proto:field=google.cloud.batch.v1.TaskStatus.state
	State *string `json:"state,omitempty"`

	// Detailed info about why the state is reached.
	// +kcc:proto:field=google.cloud.batch.v1.TaskStatus.status_events
	StatusEvents []StatusEvent `json:"statusEvents,omitempty"`
}
