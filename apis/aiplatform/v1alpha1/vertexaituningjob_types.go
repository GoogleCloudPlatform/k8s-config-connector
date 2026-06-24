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

var VertexAITuningJobGVK = GroupVersion.WithKind("VertexAITuningJob")

// VertexAITuningJobSpec defines the desired state of VertexAITuningJob
// +kcc:spec:proto=google.cloud.aiplatform.v1.TuningJob
type VertexAITuningJobSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAITuningJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The base model that is being tuned. See [Supported
	//  models](https://cloud.google.com/vertex-ai/generative-ai/docs/model-reference/tuning#supported_models).
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.base_model
	BaseModel *string `json:"baseModel,omitempty"`

	// Tuning Spec for Supervised Fine Tuning.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.supervised_tuning_spec
	SupervisedTuningSpec *SupervisedTuningSpec `json:"supervisedTuningSpec,omitempty"`

	// Optional. The display name of the
	//  [TunedModel][google.cloud.aiplatform.v1.Model]. The name can be up to 128
	//  characters long and can consist of any UTF-8 characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.tuned_model_display_name
	TunedModelDisplayName *string `json:"tunedModelDisplayName,omitempty"`

	// Optional. The description of the
	//  [TuningJob][google.cloud.aiplatform.v1.TuningJob].
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.description
	Description *string `json:"description,omitempty"`

	// Optional. The labels with user-defined metadata to organize
	//  [TuningJob][google.cloud.aiplatform.v1.TuningJob] and generated resources
	//  such as [Model][google.cloud.aiplatform.v1.Model] and
	//  [Endpoint][google.cloud.aiplatform.v1.Endpoint].
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key options for a TuningJob. If this is set,
	//  then all resources created by the TuningJob will be encrypted with the
	//  provided encryption key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// The service account that the tuningJob workload runs as.
	//  If not specified, the Vertex AI Secure Fine-Tuned Service Agent in the
	//  project will be used. See
	//  https://cloud.google.com/iam/docs/service-agents#vertex-ai-secure-fine-tuning-service-agent
	//
	//  Users starting the pipeline must have the `iam.serviceAccounts.actAs`
	//  permission on this service account.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`
}

// VertexAITuningJobStatus defines the config connector machine state of VertexAITuningJob
type VertexAITuningJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAITuningJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAITuningJobObservedState `json:"observedState,omitempty"`
}

// VertexAITuningJobObservedState is the state of the VertexAITuningJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.TuningJob
type VertexAITuningJobObservedState struct {
	// Output only. Identifier. Resource name of a TuningJob. Format:
	//  `projects/{project}/locations/{location}/tuningJobs/{tuning_job}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.name
	Name *string `json:"name,omitempty"`

	// Output only. The detailed state of the job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.state
	State *string `json:"state,omitempty"`

	// Output only. Time when the
	//  [TuningJob][google.cloud.aiplatform.v1.TuningJob] was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the
	//  [TuningJob][google.cloud.aiplatform.v1.TuningJob] for the first time
	//  entered the `JOB_STATE_RUNNING` state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the TuningJob entered any of the following
	//  [JobStates][google.cloud.aiplatform.v1.JobState]: `JOB_STATE_SUCCEEDED`,
	//  `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`, `JOB_STATE_EXPIRED`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the
	//  [TuningJob][google.cloud.aiplatform.v1.TuningJob] was most recently
	//  updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Only populated when job's state is `JOB_STATE_FAILED` or
	//  `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. The Experiment associated with this
	//  [TuningJob][google.cloud.aiplatform.v1.TuningJob].
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.experiment
	Experiment *string `json:"experiment,omitempty"`

	// Output only. The tuned model that is created as a result of this job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.tuned_model
	TunedModel *TunedModelObservedState `json:"tunedModel,omitempty"`

	// Output only. The tuning data statistics associated with this
	//  [TuningJob][google.cloud.aiplatform.v1.TuningJob].
	// +kcc:proto:field=google.cloud.aiplatform.v1.TuningJob.tuning_data_stats
	TuningDataStats *TuningDataStatsObservedState `json:"tuningDataStats,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaituningjob;gcpvertexaituningjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAITuningJob is the Schema for the VertexAITuningJob API
// +k8s:openapi-gen=true
type VertexAITuningJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAITuningJobSpec   `json:"spec,omitempty"`
	Status VertexAITuningJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAITuningJobList contains a list of VertexAITuningJob
type VertexAITuningJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAITuningJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAITuningJob{}, &VertexAITuningJobList{})
}
