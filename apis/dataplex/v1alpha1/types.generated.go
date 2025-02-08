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

// +kcc:proto=google.cloud.dataplex.v1.Task
type Task struct {

	// Optional. Description of the task.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.description
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the task.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Spec related to how often and when a task should be triggered.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.trigger_spec
	TriggerSpec *Task_TriggerSpec `json:"triggerSpec,omitempty"`

	// Required. Spec related to how a task is executed.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.execution_spec
	ExecutionSpec *Task_ExecutionSpec `json:"executionSpec,omitempty"`

	// Config related to running custom Spark tasks.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.spark
	Spark *Task_SparkTaskConfig `json:"spark,omitempty"`

	// Config related to running scheduled Notebooks.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.notebook
	Notebook *Task_NotebookTaskConfig `json:"notebook,omitempty"`
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

// +kcc:proto=google.cloud.dataplex.v1.Task.ExecutionStatus
type Task_ExecutionStatus struct {
}

// +kcc:proto=google.cloud.dataplex.v1.Task.InfrastructureSpec
type Task_InfrastructureSpec struct {
	// Compute resources needed for a Task when using Dataproc Serverless.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.batch
	Batch *Task_InfrastructureSpec_BatchComputeResources `json:"batch,omitempty"`

	// Container Image Runtime Configuration.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.container_image
	ContainerImage *Task_InfrastructureSpec_ContainerImageRuntime `json:"containerImage,omitempty"`

	// Vpc network.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.vpc_network
	VpcNetwork *Task_InfrastructureSpec_VpcNetwork `json:"vpcNetwork,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.InfrastructureSpec.BatchComputeResources
type Task_InfrastructureSpec_BatchComputeResources struct {
	// Optional. Total number of job executors.
	//  Executor Count should be between 2 and 100. [Default=2]
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.BatchComputeResources.executors_count
	ExecutorsCount *int32 `json:"executorsCount,omitempty"`

	// Optional. Max configurable executors.
	//  If max_executors_count > executors_count, then auto-scaling is enabled.
	//  Max Executor Count should be between 2 and 1000. [Default=1000]
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.BatchComputeResources.max_executors_count
	MaxExecutorsCount *int32 `json:"maxExecutorsCount,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.InfrastructureSpec.ContainerImageRuntime
type Task_InfrastructureSpec_ContainerImageRuntime struct {
	// Optional. Container image to use.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.ContainerImageRuntime.image
	Image *string `json:"image,omitempty"`

	// Optional. A list of Java JARS to add to the classpath.
	//  Valid input includes Cloud Storage URIs to Jar binaries.
	//  For example, gs://bucket-name/my/path/to/file.jar
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.ContainerImageRuntime.java_jars
	JavaJars []string `json:"javaJars,omitempty"`

	// Optional. A list of python packages to be installed.
	//  Valid formats include Cloud Storage URI to a PIP installable library.
	//  For example, gs://bucket-name/my/path/to/lib.tar.gz
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.ContainerImageRuntime.python_packages
	PythonPackages []string `json:"pythonPackages,omitempty"`

	// Optional. Override to common configuration of open source components
	//  installed on the Dataproc cluster. The properties to set on daemon
	//  config files. Property keys are specified in `prefix:property` format,
	//  for example `core:hadoop.tmp.dir`. For more information, see [Cluster
	//  properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.ContainerImageRuntime.properties
	Properties map[string]string `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.InfrastructureSpec.VpcNetwork
type Task_InfrastructureSpec_VpcNetwork struct {
	// Optional. The Cloud VPC network in which the job is run. By default,
	//  the Cloud VPC network named Default within the project is used.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.VpcNetwork.network
	Network *string `json:"network,omitempty"`

	// Optional. The Cloud VPC sub-network in which the job is run.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.VpcNetwork.sub_network
	SubNetwork *string `json:"subNetwork,omitempty"`

	// Optional. List of network tags to apply to the job.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.InfrastructureSpec.VpcNetwork.network_tags
	NetworkTags []string `json:"networkTags,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.NotebookTaskConfig
type Task_NotebookTaskConfig struct {
	// Required. Path to input notebook. This can be the Cloud Storage URI of
	//  the notebook file or the path to a Notebook Content. The execution args
	//  are accessible as environment variables
	//  (`TASK_key=value`).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.NotebookTaskConfig.notebook
	Notebook *string `json:"notebook,omitempty"`

	// Optional. Infrastructure specification for the execution.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.NotebookTaskConfig.infrastructure_spec
	InfrastructureSpec *Task_InfrastructureSpec `json:"infrastructureSpec,omitempty"`

	// Optional. Cloud Storage URIs of files to be placed in the working
	//  directory of each executor.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.NotebookTaskConfig.file_uris
	FileUris []string `json:"fileUris,omitempty"`

	// Optional. Cloud Storage URIs of archives to be extracted into the working
	//  directory of each executor. Supported file types: .jar, .tar, .tar.gz,
	//  .tgz, and .zip.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.NotebookTaskConfig.archive_uris
	ArchiveUris []string `json:"archiveUris,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.SparkTaskConfig
type Task_SparkTaskConfig struct {
	// The Cloud Storage URI of the jar file that contains the main class.
	//  The execution args are passed in as a sequence of named process
	//  arguments (`--key=value`).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.main_jar_file_uri
	MainJarFileURI *string `json:"mainJarFileURI,omitempty"`

	// The name of the driver's main class. The jar file that contains the
	//  class must be in the default CLASSPATH or specified in
	//  `jar_file_uris`.
	//  The execution args are passed in as a sequence of named process
	//  arguments (`--key=value`).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.main_class
	MainClass *string `json:"mainClass,omitempty"`

	// The Gcloud Storage URI of the main Python file to use as the driver.
	//  Must be a .py file. The execution args are passed in as a sequence of
	//  named process arguments (`--key=value`).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.python_script_file
	PythonScriptFile *string `json:"pythonScriptFile,omitempty"`

	// A reference to a query file. This should be the Cloud Storage URI of
	//  the query file. The execution args are used to declare a set of script
	//  variables (`set key="value";`).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.sql_script_file
	SqlScriptFile *string `json:"sqlScriptFile,omitempty"`

	// The query text.
	//  The execution args are used to declare a set of script variables
	//  (`set key="value";`).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.sql_script
	SqlScript *string `json:"sqlScript,omitempty"`

	// Optional. Cloud Storage URIs of files to be placed in the working
	//  directory of each executor.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.file_uris
	FileUris []string `json:"fileUris,omitempty"`

	// Optional. Cloud Storage URIs of archives to be extracted into the working
	//  directory of each executor. Supported file types: .jar, .tar, .tar.gz,
	//  .tgz, and .zip.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.archive_uris
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// Optional. Infrastructure specification for the execution.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.infrastructure_spec
	InfrastructureSpec *Task_InfrastructureSpec `json:"infrastructureSpec,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.TriggerSpec
type Task_TriggerSpec struct {
	// Required. Immutable. Trigger type of the user-specified Task.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.TriggerSpec.type
	Type *string `json:"type,omitempty"`

	// Optional. The first run of the task will be after this time.
	//  If not specified, the task will run shortly after being submitted if
	//  ON_DEMAND and based on the schedule if RECURRING.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.TriggerSpec.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. Prevent the task from executing.
	//  This does not cancel already running tasks. It is intended to temporarily
	//  disable RECURRING tasks.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.TriggerSpec.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Optional. Number of retry attempts before aborting.
	//  Set to zero to never attempt to retry a failed task.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.TriggerSpec.max_retries
	MaxRetries *int32 `json:"maxRetries,omitempty"`

	// Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron) for
	//  running tasks periodically. To explicitly set a timezone to the cron
	//  tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or
	//  "TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid
	//  string from IANA time zone database. For example,
	//  `CRON_TZ=America/New_York 1 * * * *`, or `TZ=America/New_York 1 * * *
	//  *`. This field is required for RECURRING tasks.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.TriggerSpec.schedule
	Schedule *string `json:"schedule,omitempty"`
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

// +kcc:proto=google.cloud.dataplex.v1.Task
type TaskObservedState struct {
	// Output only. The relative resource name of the task, of the form:
	//  projects/{project_number}/locations/{location_id}/lakes/{lake_id}/
	//  tasks/{task_id}.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the task. This ID will
	//  be different if the task is deleted and re-created with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the task was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the task was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the task.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.state
	State *string `json:"state,omitempty"`

	// Output only. Status of the latest task executions.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.execution_status
	ExecutionStatus *Task_ExecutionStatus `json:"executionStatus,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Task.ExecutionStatus
type Task_ExecutionStatusObservedState struct {
	// Output only. Last update time of the status.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. latest job execution
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionStatus.latest_job
	LatestJob *Job `json:"latestJob,omitempty"`
}
