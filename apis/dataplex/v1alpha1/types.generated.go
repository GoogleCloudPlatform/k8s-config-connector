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


// +kcc:proto=google.cloud.dataplex.v1.Job
type Job struct {
}

// +kcc:proto=google.cloud.dataplex.v1.Task.ExecutionSpec
type Task_ExecutionSpec struct {
	// Optional. The arguments to pass to the task.
	//  The args can use placeholders of the format ${placeholder} as
	//  part of key/value string. These will be interpolated before passing the
	//  args to the driver. Currently supported placeholders:
	//  - ${task_id}
	//  - ${job_time}
	//  To pass positional args, set the key as TASK_ARGS. The value should be a
	//  comma-separated string of all the positional arguments. To use a
	//  delimiter other than comma, refer to
	//  https://cloud.google.com/sdk/gcloud/reference/topic/escaping. In case of
	//  other keys being present in the args, then TASK_ARGS will be passed as
	//  the last argument.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.args
	Args map[string]string `json:"args,omitempty"`

	// Required. Service account to use to execute a task.
	//  If not provided, the default Compute service account for the project is
	//  used.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. The project in which jobs are run. By default, the project
	//  containing the Lake is used. If a project is provided, the
	//  [ExecutionSpec.service_account][google.cloud.dataplex.v1.Task.ExecutionSpec.service_account]
	//  must belong to this project.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.project
	Project *string `json:"project,omitempty"`

	// Optional. The maximum duration after which the job execution is expired.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.max_job_execution_lifetime
	MaxJobExecutionLifetime *string `json:"maxJobExecutionLifetime,omitempty"`

	// Optional. The Cloud KMS key to use for encryption, of the form:
	//  `projects/{project_number}/locations/{location_id}/keyRings/{key-ring-name}/cryptoKeys/{key-name}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Job
type JobObservedState struct {
	// Output only. The relative resource name of the job, of the form:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}/tasks/{task_id}/jobs/{job_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the job.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the job was started.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time when the job ended.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Execution state for the job.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.state
	State *string `json:"state,omitempty"`

	// Output only. The number of times the job has been retried (excluding the
	//  initial attempt).
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.retry_count
	RetryCount *uint32 `json:"retryCount,omitempty"`

	// Output only. The underlying service running a job.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.service
	Service *string `json:"service,omitempty"`

	// Output only. The full resource name for the job run under a particular
	//  service.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.service_job
	ServiceJob *string `json:"serviceJob,omitempty"`

	// Output only. Additional information about the current state.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.message
	Message *string `json:"message,omitempty"`

	// Output only. User-defined labels for the task.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. Job execution trigger.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.trigger
	Trigger *string `json:"trigger,omitempty"`

	// Output only. Spec related to how a task is executed.
	// +kcc:proto:field=google.cloud.dataplex.v1.Job.execution_spec
	ExecutionSpec *Task_ExecutionSpec `json:"executionSpec,omitempty"`
}
