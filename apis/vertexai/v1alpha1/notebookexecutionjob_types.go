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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAINotebookExecutionJobGVK = GroupVersion.WithKind("VertexAINotebookExecutionJob")

// VertexAINotebookExecutionJobSpec defines the desired state of VertexAINotebookExecutionJob
// +kcc:spec:proto=google.cloud.aiplatform.v1.NotebookExecutionJob
type VertexAINotebookExecutionJobSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The VertexAINotebookExecutionJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

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
	GCSOutputURI *string `json:"gcsOutputURI,omitempty"`

	// The user email to run the execution as. Only supported by Colab runtimes.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.execution_user
	ExecutionUser *string `json:"executionUser,omitempty"`

	// The service account to run the execution as.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// The Workbench runtime configuration to use for the notebook execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NotebookExecutionJob.workbench_runtime
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Type=object
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
	ScheduleResourceName *string `json:"scheduleResourceName,omitempty"`

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

// VertexAINotebookExecutionJobStatus defines the config connector machine state of VertexAINotebookExecutionJob
type VertexAINotebookExecutionJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAINotebookExecutionJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAINotebookExecutionJobObservedState `json:"observedState,omitempty"`
}

// VertexAINotebookExecutionJobObservedState is the state of the VertexAINotebookExecutionJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.NotebookExecutionJob
type VertexAINotebookExecutionJobObservedState struct {
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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexainotebookexecutionjob;gcpvertexainotebookexecutionjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAINotebookExecutionJob is the Schema for the VertexAINotebookExecutionJob API
// +k8s:openapi-gen=true
type VertexAINotebookExecutionJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAINotebookExecutionJobSpec   `json:"spec,omitempty"`
	Status VertexAINotebookExecutionJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAINotebookExecutionJobList contains a list of VertexAINotebookExecutionJob
type VertexAINotebookExecutionJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAINotebookExecutionJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAINotebookExecutionJob{}, &VertexAINotebookExecutionJobList{})
}
