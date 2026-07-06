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
	computev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIPipelineJobGVK = GroupVersion.WithKind("VertexAIPipelineJob")

type PipelineJob = VertexAIPipelineJobSpec
type PipelineJobObservedState = VertexAIPipelineJobObservedState

// VertexAIPipelineJobSpec defines the desired state of VertexAIPipelineJob
// +kcc:spec:proto=google.cloud.aiplatform.v1.PipelineJob
type VertexAIPipelineJobSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The VertexAIPipelineJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The display name of the Pipeline.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The spec of the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.pipeline_spec
	PipelineSpec *apiextensionsv1.JSON `json:"pipelineSpec,omitempty"`

	// The labels with user-defined metadata to organize PipelineJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Runtime config of the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.runtime_config
	RuntimeConfig *PipelineJobRuntimeConfig `json:"runtimeConfig,omitempty"`

	// Customer-managed encryption key spec for a pipelineJob. If set, this
	//  PipelineJob and all of its sub-resources will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// The service account that the pipeline workload runs as.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Reference to a ComputeNetwork.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// A list of names for the reserved ip ranges under the VPC network
	//  that can be used for this Pipeline Job's workload.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.reserved_ip_ranges
	ReservedIPRanges []string `json:"reservedIPRanges,omitempty"`

	// Optional. Configuration for PSC-I for PipelineJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.psc_interface_config
	PSCInterfaceConfig *PSCInterfaceConfig `json:"pscInterfaceConfig,omitempty"`

	// A template uri from where the PipelineJob.pipeline_spec, if empty, will be downloaded.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.template_uri
	TemplateURI *string `json:"templateURI,omitempty"`

	// Optional. Whether to do component level validations before job creation.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.preflight_validations
	PreflightValidations *bool `json:"preflightValidations,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PipelineJob.RuntimeConfig
type PipelineJobRuntimeConfig struct {
	// Required. A path in a Cloud Storage bucket, which will be treated as the
	//  root output directory of the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.RuntimeConfig.gcs_output_directory
	GCSOutputDirectory *string `json:"gcsOutputDirectory,omitempty"`

	// Represents the failure policy of a pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.RuntimeConfig.failure_policy
	FailurePolicy *string `json:"failurePolicy,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.PscInterfaceConfig
type PSCInterfaceConfig struct {
	// Optional. The name of the Compute Engine network attachment to attach to the resource within the region and user project.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PscInterfaceConfig.network_attachment
	NetworkAttachmentRef *computev1alpha1.ComputeNetworkAttachmentRef `json:"networkAttachmentRef,omitempty"`

	// Optional. DNS peering configurations.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PscInterfaceConfig.dns_peering_configs
	DNSPeeringConfigs []DNSPeeringConfig `json:"dnsPeeringConfigs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.DnsPeeringConfig
type DNSPeeringConfig struct {
	// Required. The DNS name suffix of the zone being peered to, e.g., "my-internal-domain.corp.". Must end with a dot.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DnsPeeringConfig.domain
	Domain *string `json:"domain,omitempty"`

	// Required. The project ID hosting the Cloud DNS managed zone that contains the 'domain'.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DnsPeeringConfig.target_project
	TargetProject *string `json:"targetProject,omitempty"`

	// Required. The VPC network name in the target_project where the DNS zone specified by 'domain' is visible.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DnsPeeringConfig.target_network
	TargetNetworkRef *computev1beta1.ComputeNetworkRef `json:"targetNetworkRef,omitempty"`
}

// VertexAIPipelineJobStatus defines the config connector machine state of VertexAIPipelineJob
type VertexAIPipelineJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIPipelineJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIPipelineJobObservedState `json:"observedState,omitempty"`
}

// VertexAIPipelineJobObservedState is the state of the VertexAIPipelineJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.PipelineJob
type VertexAIPipelineJobObservedState struct {
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
	JobDetail *PipelineJobDetailObservedState `json:"jobDetail,omitempty"`

	// Output only. The error that occurred during pipeline execution.
	//  Only populated when the pipeline's state is FAILED or CANCELLED.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJob.error
	Error *common.Status `json:"error,omitempty"`

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

// +kcc:observedstate:proto=google.cloud.aiplatform.v1.PipelineJobDetail
type PipelineJobDetailObservedState struct {
	// Output only. The context of the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJobDetail.pipeline_context
	PipelineContext *ContextObservedState `json:"pipelineContext,omitempty"`

	// Output only. The context of the current pipeline run.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJobDetail.pipeline_run_context
	PipelineRunContext *ContextObservedState `json:"pipelineRunContext,omitempty"`

	// Output only. The runtime details of the tasks under the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineJobDetail.task_details
	TaskDetails []PipelineTaskDetailObservedState `json:"taskDetails,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.aiplatform.v1.Context
type ContextObservedState struct {
	// Immutable. The resource name of the Context.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.name
	Name *string `json:"name,omitempty"`

	// User provided display name of the Context.
	//  May be up to 128 Unicode characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// An eTag used to perform consistent read-modify-write updates. If not set, a
	//  blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Contexts.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. Timestamp when this Context was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Context was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A list of resource names of Contexts that are parents of this Context.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.parent_contexts
	ParentContexts []string `json:"parentContexts,omitempty"`

	// The title of the schema describing the metadata.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.schema_title
	SchemaTitle *string `json:"schemaTitle,omitempty"`

	// The version of the schema in schema_name to use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Properties of the Context.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.metadata
	Metadata *apiextensionsv1.JSON `json:"metadata,omitempty"`

	// Description of the Context
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.description
	Description *string `json:"description,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.aiplatform.v1.PipelineTaskDetail
type PipelineTaskDetailObservedState struct {
	// Output only. The system generated ID of the task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.task_id
	TaskID *int64 `json:"taskID,omitempty"`

	// Output only. The id of the parent task if the task is within a component scope.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.parent_task_id
	ParentTaskID *int64 `json:"parentTaskID,omitempty"`

	// Output only. The user specified name of the task.
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
	ExecutorDetail *PipelineTaskExecutorDetailObservedState `json:"executorDetail,omitempty"`

	// Output only. State of the task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.state
	State *string `json:"state,omitempty"`

	// Output only. The execution metadata of the task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.execution
	Execution *ExecutionObservedState `json:"execution,omitempty"`

	// Output only. The error that occurred during task execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. A list of task status.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.pipeline_task_status
	PipelineTaskStatus []PipelineTaskDetail_PipelineTaskStatusObservedState `json:"pipelineTaskStatus,omitempty"`

	// Output only. The unique name of a task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.task_unique_name
	TaskUniqueName *string `json:"taskUniqueName,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.aiplatform.v1.Execution
type ExecutionObservedState struct {
	// Output only. The resource name of the Execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.name
	Name *string `json:"name,omitempty"`

	// User provided display name of the Execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The state of this Execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.state
	State *string `json:"state,omitempty"`

	// An eTag used to perform consistent read-modify-write updates.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Executions.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. Timestamp when this Execution was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Execution was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// The title of the schema describing the metadata.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.schema_title
	SchemaTitle *string `json:"schemaTitle,omitempty"`

	// The version of the schema in `schema_title` to use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Properties of the Execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.metadata
	Metadata *apiextensionsv1.JSON `json:"metadata,omitempty"`

	// Description of the Execution
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.description
	Description *string `json:"description,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.aiplatform.v1.PipelineTaskDetail.PipelineTaskStatus
type PipelineTaskDetail_PipelineTaskStatusObservedState struct {
	// Output only. Update time of this status.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.PipelineTaskStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the task.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.PipelineTaskStatus.state
	State *string `json:"state,omitempty"`

	// Output only. The error that occurred during the state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskDetail.PipelineTaskStatus.error
	Error *common.Status `json:"error,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail
type PipelineTaskExecutorDetailObservedState struct {
	// Output only. The detailed info for a container executor.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.container_detail
	ContainerDetail *PipelineTaskExecutorDetail_ContainerDetailObservedState `json:"containerDetail,omitempty"`

	// Output only. The detailed info for a custom job executor.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.custom_job_detail
	CustomJobDetail *PipelineTaskExecutorDetail_CustomJobDetailObservedState `json:"customJobDetail,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.ContainerDetail
type PipelineTaskExecutorDetail_ContainerDetailObservedState struct {
	// Output only. The name of the CustomJob for the main container execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.ContainerDetail.main_job
	MainJob *string `json:"mainJob,omitempty"`

	// Output only. The name of the CustomJob for the pre-caching-check container execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.ContainerDetail.pre_caching_check_job
	PreCachingCheckJob *string `json:"preCachingCheckJob,omitempty"`

	// Output only. The names of the previously failed CustomJob for the main container executions.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.ContainerDetail.failed_main_jobs
	FailedMainJobs []string `json:"failedMainJobs,omitempty"`

	// Output only. The names of the previously failed CustomJob for the pre-caching-check container executions.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.ContainerDetail.failed_pre_caching_check_jobs
	FailedPreCachingCheckJobs []string `json:"failedPreCachingCheckJobs,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.CustomJobDetail
type PipelineTaskExecutorDetail_CustomJobDetailObservedState struct {
	// Output only. The name of the CustomJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.CustomJobDetail.job
	Job *string `json:"job,omitempty"`

	// Output only. The names of the previously failed CustomJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTaskExecutorDetail.CustomJobDetail.failed_jobs
	FailedJobs []string `json:"failedJobs,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.aiplatform.v1.PipelineTemplateMetadata
type PipelineTemplateMetadata struct {
	// The version_name in artifact registry.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PipelineTemplateMetadata.version
	Version *string `json:"version,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaipipelinejob;gcpvertexaipipelinejobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIPipelineJob is the Schema for the VertexAIPipelineJob API
// +k8s:openapi-gen=true
type VertexAIPipelineJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIPipelineJobSpec   `json:"spec,omitempty"`
	Status VertexAIPipelineJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIPipelineJobList contains a list of VertexAIPipelineJob
type VertexAIPipelineJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIPipelineJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIPipelineJob{}, &VertexAIPipelineJobList{})
}
