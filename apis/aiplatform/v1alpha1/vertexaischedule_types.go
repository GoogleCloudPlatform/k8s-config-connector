// Copyright 2026 Google LLC
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
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	dataformv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIScheduleGVK = GroupVersion.WithKind("VertexAISchedule")

// VertexAIScheduleSpec defines the desired state of VertexAISchedule
// +kcc:spec:proto=google.cloud.aiplatform.v1.Schedule
type VertexAIScheduleSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAISchedule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled
	//  runs. To explicitly set a timezone to the cron tab, apply a prefix in the
	//  cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or "TZ=${IANA_TIME_ZONE}".
	//  The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone
	//  database. For example, "CRON_TZ=America/New_York 1 * * * *", or
	//  "TZ=America/New_York 1 * * * *".
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.cron
	Cron *string `json:"cron,omitempty"`

	// Request for
	//  [PipelineService.CreatePipelineJob][google.cloud.aiplatform.v1.PipelineService.CreatePipelineJob].
	//  The parent field of CreatePipelineJobRequest is required (format: projects-id/locations-id).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.create_pipeline_job_request
	CreatePipelineJobRequest *CreatePipelineJobRequest `json:"createPipelineJobRequest,omitempty"`

	// Request for
	//  [NotebookService.CreateNotebookExecutionJob][google.cloud.aiplatform.v1.NotebookService.CreateNotebookExecutionJob].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.create_notebook_execution_job_request
	CreateNotebookExecutionJobRequest *CreateNotebookExecutionJobRequest `json:"createNotebookExecutionJobRequest,omitempty"`

	// Required. User provided name of the Schedule.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.display_name
	// +required
	DisplayName *string `json:"displayName"`

	// Optional. Timestamp after which the first run can be scheduled.
	//  Default to Schedule create time if not specified.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. Timestamp after which no new runs can be scheduled.
	//  If specified, The schedule will be completed when either
	//  end_time is reached or when scheduled_run_count >= max_run_count.
	//  If not specified, new runs will keep getting scheduled until this Schedule
	//  is paused or deleted. Already scheduled runs will be allowed to complete.
	//  Unset if not specified.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Optional. Maximum run count of the schedule.
	//  If specified, The schedule will be completed when either
	//  started_run_count >= max_run_count or when end_time is reached.
	//  If not specified, new runs will keep getting scheduled until this Schedule
	//  is paused or deleted. Already scheduled runs will be allowed to complete.
	//  Unset if not specified.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.max_run_count
	MaxRunCount *int64 `json:"maxRunCount,omitempty"`

	// Required. Maximum number of runs that can be started concurrently for this
	//  Schedule. This is the limit for starting the scheduled requests and not the
	//  execution of the operations/jobs created by the requests (if applicable).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.max_concurrent_run_count
	// +required
	MaxConcurrentRunCount *int64 `json:"maxConcurrentRunCount"`

	// Optional. Whether new scheduled runs can be queued when max_concurrent_runs
	//  limit is reached. If set to true, new runs will be queued instead of
	//  skipped. Default to false.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.allow_queueing
	AllowQueueing *bool `json:"allowQueueing,omitempty"`
}

// VertexAIScheduleStatus defines the config connector machine state of VertexAISchedule
type VertexAIScheduleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAISchedule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIScheduleObservedState `json:"observedState,omitempty"`
}

// VertexAIScheduleObservedState is the state of the VertexAISchedule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.Schedule
type VertexAIScheduleObservedState struct {
	// Immutable. The resource name of the Schedule.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.name
	Name *string `json:"name,omitempty"`

	// Request for
	//  [PipelineService.CreatePipelineJob][google.cloud.aiplatform.v1.PipelineService.CreatePipelineJob].
	//  CreatePipelineJobRequest.parent field is required (format:
	//  projects/{project}/locations/{location}).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.create_pipeline_job_request
	CreatePipelineJobRequest *CreatePipelineJobRequestObservedState `json:"createPipelineJobRequest,omitempty"`

	// Request for
	//  [NotebookService.CreateNotebookExecutionJob][google.cloud.aiplatform.v1.NotebookService.CreateNotebookExecutionJob].
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.create_notebook_execution_job_request
	CreateNotebookExecutionJobRequest *CreateNotebookExecutionJobRequestObservedState `json:"createNotebookExecutionJobRequest,omitempty"`

	// Output only. The number of runs started by this schedule.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.started_run_count
	StartedRunCount *int64 `json:"startedRunCount,omitempty"`

	// Output only. The state of this Schedule.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.state
	State *string `json:"state,omitempty"`

	// Output only. Timestamp when this Schedule was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Schedule was updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Timestamp when this Schedule should schedule the next run.
	//  Having a next_run_time in the past means the runs are being started
	//  behind schedule.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.next_run_time
	NextRunTime *string `json:"nextRunTime,omitempty"`

	// Output only. Timestamp when this Schedule was last paused.
	//  Unset if never paused.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.last_pause_time
	LastPauseTime *string `json:"lastPauseTime,omitempty"`

	// Output only. Timestamp when this Schedule was last resumed.
	//  Unset if never resumed from pause.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.last_resume_time
	LastResumeTime *string `json:"lastResumeTime,omitempty"`

	// Output only. Whether to backfill missed runs when the schedule is resumed
	//  from PAUSED state. If set to true, all missed runs will be scheduled. New
	//  runs will be scheduled after the backfill is complete. Default to false.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.catch_up
	CatchUp *bool `json:"catchUp,omitempty"`

	// Output only. Response of the last scheduled run.
	//  This is the response for starting the scheduled requests and not the
	//  execution of the operations/jobs created by the requests (if applicable).
	//  Unset if no run has been scheduled yet.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Schedule.last_scheduled_run_response
	LastScheduledRunResponse *Schedule_RunResponse `json:"lastScheduledRunResponse,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaischedule;gcpvertexaischedules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAISchedule is the Schema for the VertexAISchedule API
// +k8s:openapi-gen=true
type VertexAISchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIScheduleSpec   `json:"spec,omitempty"`
	Status VertexAIScheduleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIScheduleList contains a list of VertexAISchedule
type VertexAIScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAISchedule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAISchedule{}, &VertexAIScheduleList{})
}

// +kcc:proto=google.cloud.aiplatform.v1.CreateNotebookExecutionJobRequest
type CreateNotebookExecutionJobRequest struct {
	// Required. Location resource name to create the NotebookExecutionJob.
	//  Format: projects-id/locations-id
	// +kcc:proto:field=google.cloud.aiplatform.v1.CreateNotebookExecutionJobRequest.parent
	Parent *string `json:"parent,omitempty"`

	// Required. The NotebookExecutionJob to create.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CreateNotebookExecutionJobRequest.notebook_execution_job
	NotebookExecutionJob *NotebookExecutionJob `json:"notebookExecutionJob,omitempty"`

	// Optional. User specified ID for the NotebookExecutionJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CreateNotebookExecutionJobRequest.notebook_execution_job_id
	NotebookExecutionJobID *string `json:"notebookExecutionJobID,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.CreatePipelineJobRequest
type CreatePipelineJobRequest struct {
	// Required. Location resource name to create the PipelineJob in.
	//  Format: projects-id/locations-id
	// +kcc:proto:field=google.cloud.aiplatform.v1.CreatePipelineJobRequest.parent
	Parent *string `json:"parent,omitempty"`

	// Required. The PipelineJob to create.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CreatePipelineJobRequest.pipeline_job
	PipelineJob *PipelineJob `json:"pipelineJob,omitempty"`

	// The ID to use for the PipelineJob, which will become the final component of
	//  the PipelineJob name. If not provided, an ID will be automatically
	//  generated.
	//
	//  This value should be less than 128 characters, and valid characters
	//  are `/[a-z][0-9]-/`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CreatePipelineJobRequest.pipeline_job_id
	PipelineJobID *string `json:"pipelineJobID,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookExecutionJob
type NotebookExecutionJob struct {
	// The Dataform Repository pointing to a single file notebook repository.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.dataform_repository_source
	DataformRepositorySource *NotebookExecutionJob_DataformRepositorySource `json:"dataformRepositorySource,omitempty"`

	// The Cloud Storage url pointing to the ipynb file. Format:
	//  `gs://bucket/notebook_file.ipynb`
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.gcs_notebook_source
	GCSNotebookSource *NotebookExecutionJob_GCSNotebookSource `json:"gcsNotebookSource,omitempty"`

	// The contents of an input notebook file.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.direct_notebook_source
	DirectNotebookSource *NotebookExecutionJob_DirectNotebookSource `json:"directNotebookSource,omitempty"`

	// The NotebookRuntimeTemplate to source compute configuration from.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.notebook_runtime_template_resource_name
	NotebookRuntimeTemplateResourceName *string `json:"notebookRuntimeTemplateResourceName,omitempty"`

	// The custom compute configuration for an execution job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.custom_environment_spec
	CustomEnvironmentSpec *NotebookExecutionJob_CustomEnvironmentSpec `json:"customEnvironmentSpec,omitempty"`

	// The Cloud Storage location to upload the result to. Format:
	//  `gs://bucket-name`
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.gcs_output_uri
	GCSOutputRef *storagev1beta1.StorageBucketRef `json:"gcsOutputRef,omitempty"`

	// The user email to run the execution as. Only supported by Colab runtimes.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.execution_user
	ExecutionUser *string `json:"executionUser,omitempty"`

	// The service account to run the execution as.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// The Workbench runtime configuration to use for the notebook execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.workbench_runtime
	WorkbenchRuntime *NotebookExecutionJob_WorkbenchRuntime `json:"workbenchRuntime,omitempty"`

	// The display name of the NotebookExecutionJob. The name can be up to 128
	//  characters long and can consist of any UTF-8 characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Max running time of the execution job in seconds (default 86400s / 24 hrs).
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.execution_timeout
	ExecutionTimeout *string `json:"executionTimeout,omitempty"`

	// The Schedule resource name if this job is triggered by one. Format:
	//  `projects/{project_id}/locations/{location}/schedules/{schedule_id}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.schedule_resource_name
	ScheduleRef *VertexAIScheduleRef `json:"scheduleRef,omitempty"`

	// The labels with user-defined metadata to organize NotebookExecutionJobs.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//  System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	//  and are immutable.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The name of the kernel to use during notebook execution. If unset, the
	//  default kernel is used.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.kernel_name
	KernelName *string `json:"kernelName,omitempty"`

	// Customer-managed encryption key spec for the notebook execution job.
	//  This field is auto-populated if the
	//  [NotebookRuntimeTemplate][google.cloud.aiplatform.v1.NotebookRuntimeTemplate]
	//  has an encryption spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NotebookExecutionJob.DataformRepositorySource
type NotebookExecutionJob_DataformRepositorySource struct {
	// The resource name of the Dataform Repository. Format:
	//  projects/PROJECT_ID/locations/LOCATION_ID/repositories/REPOSITORY_ID
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.DataformRepositorySource.dataform_repository_resource_name
	DataformRepositoryRef *dataformv1beta1.DataformRepositoryRef `json:"dataformRepositoryRef,omitempty"`

	// The commit SHA to read repository with. If unset, the file will be read
	//  at HEAD.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.DataformRepositorySource.commit_sha
	CommitSha *string `json:"commitSha,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NetworkSpec
type NetworkSpec struct {
	// Whether to enable public internet access. Default false.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NetworkSpec.enable_internet_access
	EnableInternetAccess *bool `json:"enableInternetAccess,omitempty"`

	// The full name of the Google Compute Engine
	//  [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks)
	// +kcc:proto:field=google.cloud.aiplatform.v1.NetworkSpec.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// The name of the subnet that this instance is in.
	//  Format:
	//  projects/PROJECT_ID/regions/REGION/subnetworks/SUBNETWORK_ID
	// +kcc:proto:field=google.cloud.aiplatform.v1.NetworkSpec.subnetwork
	SubnetworkRef *computev1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.aiplatform.v1.CreateNotebookExecutionJobRequest
type CreateNotebookExecutionJobRequestObservedState struct {
	// Required. The NotebookExecutionJob to create.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CreateNotebookExecutionJobRequest.notebook_execution_job
	NotebookExecutionJob *NotebookExecutionJobObservedState `json:"notebookExecutionJob,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.aiplatform.v1.CreatePipelineJobRequest
type CreatePipelineJobRequestObservedState struct {
	// Required. The PipelineJob to create.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CreatePipelineJobRequest.pipeline_job
	PipelineJob *PipelineJobObservedState `json:"pipelineJob,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.aiplatform.v1.NotebookExecutionJob
type NotebookExecutionJobObservedState struct {
	// Output only. The resource name of this NotebookExecutionJob. Format:
	//  `projects/{project_id}/locations/{location}/notebookExecutionJobs/{job_id}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.name
	Name *string `json:"name,omitempty"`

	// Output only. The state of the NotebookExecutionJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.job_state
	JobState *string `json:"jobState,omitempty"`

	// Output only. Populated when the NotebookExecutionJob is completed. When
	//  there is an error during notebook execution, the error details are
	//  populated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.status
	Status *common.Status `json:"status,omitempty"`

	// Output only. Timestamp when this NotebookExecutionJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this NotebookExecutionJob was most recently
	//  updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +k8s:deepcopy-gen=false
type PipelineJob = VertexAIPipelineJobSpec

// +k8s:deepcopy-gen=false
type PipelineJobObservedState = VertexAIPipelineJobObservedState

// +kcc:proto=google.cloud.aiplatform.v1.NotebookExecutionJob.GcsNotebookSource
type NotebookExecutionJob_GCSNotebookSource struct {
	// Reference to a StorageBucket.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.GcsNotebookSource.uri
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// Name of the Cloud Storage object.
	Object *string `json:"object,omitempty"`

	// The version of the Cloud Storage object to read. If unset, the current
	//  version of the object is read. See
	//  https://cloud.google.com/storage/docs/metadata#generation-number.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.GcsNotebookSource.generation
	Generation *string `json:"generation,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.EncryptionSpec
type EncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed
	//  encryption key used to protect a resource. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  The key needs to be in the same region as where the compute resource is
	//  created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EncryptionSpec.kms_key_name
	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}
