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


// +kcc:proto=google.cloud.aiplatform.v1.EncryptionSpec
type EncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EncryptionSpec.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineJob
type PipelineJob struct {

	// The display name of the Pipeline.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The spec of the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.pipeline_spec
	PipelineSpec map[string]string `json:"pipelineSpec,omitempty"`

	// The labels with user-defined metadata to organize PipelineJob.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//
	//  Note there is some reserved label key for Vertex AI Pipelines.
	//  - `vertex-ai-pipelines-run-billing-id`, user set value will get overrided.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Runtime config of the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.runtime_config
	RuntimeConfig *PipelineJob_RuntimeConfig `json:"runtimeConfig,omitempty"`

	// Customer-managed encryption key spec for a pipelineJob. If set, this
	//  PipelineJob and all of its sub-resources will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// The service account that the pipeline workload runs as.
	//  If not specified, the Compute Engine default service account in the project
	//  will be used.
	//  See
	//  https://cloud.google.com/compute/docs/access/service-accounts#default_service_account
	//
	//  Users starting the pipeline must have the `iam.serviceAccounts.actAs`
	//  permission on this service account.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// The full name of the Compute Engine
	//  [network](/compute/docs/networks-and-firewalls#networks) to which the
	//  Pipeline Job's workload should be peered. For example,
	//  `projects/12345/global/networks/myVPC`.
	//  [Format](/compute/docs/reference/rest/v1/networks/insert)
	//  is of the form `projects/{project}/global/networks/{network}`.
	//  Where {project} is a project number, as in `12345`, and {network} is a
	//  network name.
	//
	//  Private services access must already be configured for the network.
	//  Pipeline job will apply the network configuration to the Google Cloud
	//  resources being launched, if applied, such as Vertex AI
	//  Training or Dataflow job. If left unspecified, the workload is not peered
	//  with any network.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.network
	Network *string `json:"network,omitempty"`

	// A list of names for the reserved ip ranges under the VPC network
	//  that can be used for this Pipeline Job's workload.
	//
	//  If set, we will deploy the Pipeline Job's workload within the provided ip
	//  ranges. Otherwise, the job will be deployed to any ip ranges under the
	//  provided VPC network.
	//
	//  Example: ['vertex-ai-ip-range'].
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.reserved_ip_ranges
	ReservedIPRanges []string `json:"reservedIPRanges,omitempty"`

	// A template uri from where the
	//  [PipelineJob.pipeline_spec][google.cloud.aiplatform.v1.PipelineJob.pipeline_spec],
	//  if empty, will be downloaded. Currently, only uri from Vertex Template
	//  Registry & Gallery is supported. Reference to
	//  https://cloud.google.com/vertex-ai/docs/pipelines/create-pipeline-template.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.template_uri
	TemplateURI *string `json:"templateURI,omitempty"`

	// Optional. Whether to do component level validations before job creation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.preflight_validations
	PreflightValidations *bool `json:"preflightValidations,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineJob.RuntimeConfig
type PipelineJob_RuntimeConfig struct {

	// TODO: unsupported map type with key string and value message


	// Required. A path in a Cloud Storage bucket, which will be treated as the
	//  root output directory of the pipeline. It is used by the system to
	//  generate the paths of output artifacts. The artifact paths are generated
	//  with a sub-path pattern `{job_id}/{task_id}/{output_key}` under the
	//  specified output directory. The service account specified in this
	//  pipeline must have the `storage.objects.get` and `storage.objects.create`
	//  permissions for this bucket.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.RuntimeConfig.gcs_output_directory
	GcsOutputDirectory *string `json:"gcsOutputDirectory,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Represents the failure policy of a pipeline. Currently, the default of a
	//  pipeline is that the pipeline will continue to run until no more tasks
	//  can be executed, also known as PIPELINE_FAILURE_POLICY_FAIL_SLOW.
	//  However, if a pipeline is set to PIPELINE_FAILURE_POLICY_FAIL_FAST, it
	//  will stop scheduling any new tasks when a task has failed. Any scheduled
	//  tasks will continue to completion.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.RuntimeConfig.failure_policy
	FailurePolicy *string `json:"failurePolicy,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineJob.RuntimeConfig.InputArtifact
type PipelineJob_RuntimeConfig_InputArtifact struct {
	// Artifact resource id from MLMD. Which is the last portion of an
	//  artifact resource name:
	//  `projects/{project}/locations/{location}/metadataStores/default/artifacts/{artifact_id}`.
	//  The artifact must stay within the same project, location and default
	//  metadatastore as the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.RuntimeConfig.InputArtifact.artifact_id
	ArtifactID *string `json:"artifactID,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineJobDetail
type PipelineJobDetail struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineTaskDetail
type PipelineTaskDetail struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineTaskDetail.ArtifactList
type PipelineTaskDetail_ArtifactList struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineTaskDetail.PipelineTaskStatus
type PipelineTaskDetail_PipelineTaskStatus struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail
type PipelineTaskExecutorDetail struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.ContainerDetail
type PipelineTaskExecutorDetail_ContainerDetail struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.CustomJobDetail
type PipelineTaskExecutorDetail_CustomJobDetail struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineTemplateMetadata
type PipelineTemplateMetadata struct {
	// The version_name in artifact registry.
	//
	//  Will always be presented in output if the
	//  [PipelineJob.template_uri][google.cloud.aiplatform.v1.PipelineJob.template_uri]
	//  is from supported template registry.
	//
	//  Format is "sha256:abcdef123456...".
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTemplateMetadata.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Value
type Value struct {
	// An integer value.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Value.int_value
	IntValue *int64 `json:"intValue,omitempty"`

	// A double value.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Value.double_value
	DoubleValue *float64 `json:"doubleValue,omitempty"`

	// A string value.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineJob
type PipelineJobObservedState struct {
	// Output only. The resource name of the PipelineJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.name
	Name *string `json:"name,omitempty"`

	// Output only. Pipeline creation time.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Pipeline start time.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Pipeline end time.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Timestamp when this PipelineJob was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The detailed state of the job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.state
	State *string `json:"state,omitempty"`

	// Output only. The details of pipeline run. Not available in the list view.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.job_detail
	JobDetail *PipelineJobDetail `json:"jobDetail,omitempty"`

	// Output only. The error that occurred during pipeline execution.
	//  Only populated when the pipeline's state is FAILED or CANCELLED.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.error
	Error *Status `json:"error,omitempty"`

	// Output only. Pipeline template metadata. Will fill up fields if
	//  [PipelineJob.template_uri][google.cloud.aiplatform.v1.PipelineJob.template_uri]
	//  is from supported template registry.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.template_metadata
	TemplateMetadata *PipelineTemplateMetadata `json:"templateMetadata,omitempty"`

	// Output only. The schedule resource name.
	//  Only returned if the Pipeline is created by Schedule API.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.schedule_name
	ScheduleName *string `json:"scheduleName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineJobDetail
type PipelineJobDetailObservedState struct {
	// Output only. The context of the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJobDetail.pipeline_context
	PipelineContext *Context `json:"pipelineContext,omitempty"`

	// Output only. The context of the current pipeline run.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJobDetail.pipeline_run_context
	PipelineRunContext *Context `json:"pipelineRunContext,omitempty"`

	// Output only. The runtime details of the tasks under the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJobDetail.task_details
	TaskDetails []PipelineTaskDetail `json:"taskDetails,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineTaskDetail
type PipelineTaskDetailObservedState struct {
	// Output only. The system generated ID of the task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.task_id
	TaskID *int64 `json:"taskID,omitempty"`

	// Output only. The id of the parent task if the task is within a component
	//  scope. Empty if the task is at the root level.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.parent_task_id
	ParentTaskID *int64 `json:"parentTaskID,omitempty"`

	// Output only. The user specified name of the task that is defined in
	//  [pipeline_spec][google.cloud.aiplatform.v1.PipelineJob.pipeline_spec].
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.task_name
	TaskName *string `json:"taskName,omitempty"`

	// Output only. Task create time.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Task start time.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Task end time.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The detailed execution info.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.executor_detail
	ExecutorDetail *PipelineTaskExecutorDetail `json:"executorDetail,omitempty"`

	// Output only. State of the task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.state
	State *string `json:"state,omitempty"`

	// Output only. The execution metadata of the task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.execution
	Execution *Execution `json:"execution,omitempty"`

	// Output only. The error that occurred during task execution.
	//  Only populated when the task's state is FAILED or CANCELLED.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.error
	Error *Status `json:"error,omitempty"`

	// Output only. A list of task status. This field keeps a record of task
	//  status evolving over time.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.pipeline_task_status
	PipelineTaskStatus []PipelineTaskDetail_PipelineTaskStatus `json:"pipelineTaskStatus,omitempty"`

	// TODO: unsupported map type with key string and value message


	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineTaskDetail.PipelineTaskStatus
type PipelineTaskDetail_PipelineTaskStatusObservedState struct {
	// Output only. Update time of this status.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.PipelineTaskStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.PipelineTaskStatus.state
	State *string `json:"state,omitempty"`

	// Output only. The error that occurred during the state. May be set when
	//  the state is any of the non-final state (PENDING/RUNNING/CANCELLING) or
	//  FAILED state. If the state is FAILED, the error here is final and not
	//  going to be retried. If the state is a non-final state, the error
	//  indicates a system-error being retried.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.PipelineTaskStatus.error
	Error *Status `json:"error,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail
type PipelineTaskExecutorDetailObservedState struct {
	// Output only. The detailed info for a container executor.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.container_detail
	ContainerDetail *PipelineTaskExecutorDetail_ContainerDetail `json:"containerDetail,omitempty"`

	// Output only. The detailed info for a custom job executor.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.custom_job_detail
	CustomJobDetail *PipelineTaskExecutorDetail_CustomJobDetail `json:"customJobDetail,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.ContainerDetail
type PipelineTaskExecutorDetail_ContainerDetailObservedState struct {
	// Output only. The name of the
	//  [CustomJob][google.cloud.aiplatform.v1.CustomJob] for the main container
	//  execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.ContainerDetail.main_job
	MainJob *string `json:"mainJob,omitempty"`

	// Output only. The name of the
	//  [CustomJob][google.cloud.aiplatform.v1.CustomJob] for the
	//  pre-caching-check container execution. This job will be available if the
	//  [PipelineJob.pipeline_spec][google.cloud.aiplatform.v1.PipelineJob.pipeline_spec]
	//  specifies the `pre_caching_check` hook in the lifecycle events.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.ContainerDetail.pre_caching_check_job
	PreCachingCheckJob *string `json:"preCachingCheckJob,omitempty"`

	// Output only. The names of the previously failed
	//  [CustomJob][google.cloud.aiplatform.v1.CustomJob] for the main container
	//  executions. The list includes the all attempts in chronological order.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.ContainerDetail.failed_main_jobs
	FailedMainJobs []string `json:"failedMainJobs,omitempty"`

	// Output only. The names of the previously failed
	//  [CustomJob][google.cloud.aiplatform.v1.CustomJob] for the
	//  pre-caching-check container executions. This job will be available if the
	//  [PipelineJob.pipeline_spec][google.cloud.aiplatform.v1.PipelineJob.pipeline_spec]
	//  specifies the `pre_caching_check` hook in the lifecycle events. The list
	//  includes the all attempts in chronological order.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.ContainerDetail.failed_pre_caching_check_jobs
	FailedPreCachingCheckJobs []string `json:"failedPreCachingCheckJobs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.CustomJobDetail
type PipelineTaskExecutorDetail_CustomJobDetailObservedState struct {
	// Output only. The name of the
	//  [CustomJob][google.cloud.aiplatform.v1.CustomJob].
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.CustomJobDetail.job
	Job *string `json:"job,omitempty"`

	// Output only. The names of the previously failed
	//  [CustomJob][google.cloud.aiplatform.v1.CustomJob]. The list includes the
	//  all attempts in chronological order.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.CustomJobDetail.failed_jobs
	FailedJobs []string `json:"failedJobs,omitempty"`
}
