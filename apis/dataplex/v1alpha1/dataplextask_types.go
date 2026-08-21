// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexTaskGVK = GroupVersion.WithKind("DataplexTask")

// DataplexTaskSpec defines the desired state of DataplexTask
// +kcc:spec:proto=google.cloud.dataplex.v1.Task
type DataplexTaskSpec struct {
	// The DataplexTask name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	LakeRef *LakeRef `json:"lakeRef,omitempty"`

	// Optional. Description of the task.
	// +optional
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the task.
	// +optional
	// Labels map[string]string `json:"labels,omitempty"`

	// Required. Spec related to how often and when a task should be triggered.
	// +required
	TriggerSpec *Task_TriggerSpec `json:"triggerSpec,omitempty"`

	// Required. Spec related to how a task is executed.
	// +required
	ExecutionSpec *Task_ExecutionSpec `json:"executionSpec,omitempty"`

	// Config related to running custom Spark tasks.
	// Exactly one of spark or notebook must be set.
	// +optional
	Spark *Task_SparkTaskConfig `json:"spark,omitempty"`

	// +required
	// Config related to running scheduled Notebooks.
	// Exactly one of spark or notebook must be set.
	// +optional
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
	// +required
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. The project in which jobs are run. By default, the project
	//  containing the Lake is used. If a project is provided, the
	//  [ExecutionSpec.service_account][google.cloud.dataplex.v1.Task.ExecutionSpec.service_account]
	//  must belong to this project.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.project
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// Optional. The maximum duration after which the job execution is expired.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.max_job_execution_lifetime
	MaxJobExecutionLifetime *string `json:"maxJobExecutionLifetime,omitempty"`

	// Optional. The Cloud KMS key to use for encryption, of the form:
	//  `projects/{project_number}/locations/{location_id}/keyRings/{key-ring-name}/cryptoKeys/{key-name}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.kms_key
	KMSKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// DataplexTaskStatus defines the config connector machine state of DataplexTask
type DataplexTaskStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexTask resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexTaskObservedState `json:"observedState,omitempty"`
}

type DataplexTaskObservedState struct {
	// Output only. System generated globally unique ID for the task. This ID will
	// be different if the task is deleted and re-created with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.uid
	UID *string `json:"uid,omitempty"`
	// Output only. The time when the task was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the task was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the task.
	State *string `json:"state,omitempty"`

	//  Status of the task execution (e.g. Jobs).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionStatus
	ExecutionStatus *Task_ExecutionStatusObservedState `json:"executionStatus,omitempty"`
}

// Duplicate of Task_ExecutionSpec struct. As ServiceAccount and KMSKey cannot be a reference field when it's in status.
// +kcc:proto=google.cloud.dataplex.v1.Task.ExecutionSpec
type Task_ExecutionSpecObservedState struct {
	// The arguments to pass to the task.
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

	// Service account to use to execute a task.
	//  If not provided, the default Compute service account for the project is
	//  used.
	// +required
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// The project in which jobs are run. By default, the project
	//  containing the Lake is used. If a project is provided, the
	//  [ExecutionSpec.service_account][google.cloud.dataplex.v1.Task.ExecutionSpec.service_account]
	//  must belong to this project.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.project
	Project *string `json:"project,omitempty"`

	// The maximum duration after which the job execution is expired.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.max_job_execution_lifetime
	MaxJobExecutionLifetime *string `json:"maxJobExecutionLifetime,omitempty"`

	// The Cloud KMS key to use for encryption, of the form:
	//  `projects/{project_number}/locations/{location_id}/keyRings/{key-ring-name}/cryptoKeys/{key-name}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionSpec.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplextask;gcpdataplextasks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexTask is the Schema for the DataplexTask API
// +k8s:openapi-gen=true
type DataplexTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexTaskSpec   `json:"spec,omitempty"`
	Status DataplexTaskStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexTaskList contains a list of DataplexTask
type DataplexTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexTask `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexTask{}, &DataplexTaskList{})
}

// +kcc:proto=google.cloud.dataplex.v1.Task.NotebookTaskConfig
type Task_NotebookTaskConfig struct {
	// +required
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
	FileUris []string `json:"fileURIs,omitempty"`

	// Optional. Cloud Storage URIs of archives to be extracted into the working
	//  directory of each executor. Supported file types: .jar, .tar, .tar.gz,
	//  .tgz, and .zip.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.NotebookTaskConfig.archive_uris
	ArchiveUris []string `json:"archiveURIs,omitempty"`
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
	SQLScriptFile *string `json:"sqlScriptFile,omitempty"`

	// The query text.
	//  The execution args are used to declare a set of script variables
	//  (`set key="value";`).
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.sql_script
	SQLScript *string `json:"sqlScript,omitempty"`

	// Optional. Cloud Storage URIs of files to be placed in the working
	//  directory of each executor.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.file_uris
	FileUris []string `json:"fileURIs,omitempty"`

	// Optional. Cloud Storage URIs of archives to be extracted into the working
	//  directory of each executor. Supported file types: .jar, .tar, .tar.gz,
	//  .tgz, and .zip.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.archive_uris
	ArchiveUris []string `json:"archiveURIs,omitempty"`

	// Optional. Infrastructure specification for the execution.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.SparkTaskConfig.infrastructure_spec
	InfrastructureSpec *Task_InfrastructureSpec `json:"infrastructureSpec,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataplex.v1.Task.ExecutionStatus
type Task_ExecutionStatusObservedState struct {
	// Output only. Last update time of the status.
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. latest job execution
	// +kcc:proto:field=google.cloud.dataplex.v1.Task.ExecutionStatus.latest_job
	LatestJob *JobObservedState `json:"latestJob,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataplex.v1.Job
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
	ExecutionSpec *Task_ExecutionSpecObservedState `json:"executionSpec,omitempty"`
}
