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
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIHyperparameterTuningJobGVK = GroupVersion.WithKind("VertexAIHyperparameterTuningJob")

// VertexAIHyperparameterTuningJobSpec defines the desired state of VertexAIHyperparameterTuningJob
// +kcc:spec:proto=google.cloud.aiplatform.v1.HyperparameterTuningJob
type VertexAIHyperparameterTuningJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The VertexAIHyperparameterTuningJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the HyperparameterTuningJob.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.display_name
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName"`

	// Required. Study configuration of the HyperparameterTuningJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.study_spec
	// +kubebuilder:validation:Required
	StudySpec *StudySpec `json:"studySpec"`

	// Required. The desired total number of Trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.max_trial_count
	// +kubebuilder:validation:Required
	MaxTrialCount *int32 `json:"maxTrialCount"`

	// Required. The desired number of Trials to run in parallel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.parallel_trial_count
	// +kubebuilder:validation:Required
	ParallelTrialCount *int32 `json:"parallelTrialCount"`

	// The number of failed Trials that need to be seen before failing
	// the HyperparameterTuningJob.
	//
	// If set to 0, Vertex AI decides how many Trials must fail
	// before the whole job fails.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.max_failed_trial_count
	MaxFailedTrialCount *int32 `json:"maxFailedTrialCount,omitempty"`

	// Required. The spec of a trial job. The same spec applies to the CustomJobs
	// created in all the trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.trial_job_spec
	// +kubebuilder:validation:Required
	TrialJobSpec *CustomJobSpec `json:"trialJobSpec"`

	// The labels with user-defined metadata to organize HyperparameterTuningJobs.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key options for a HyperparameterTuningJob.
	// If this is set, then all resources created by the HyperparameterTuningJob
	// will be encrypted with the provided encryption key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAIHyperparameterTuningJobStatus defines the config connector machine state of VertexAIHyperparameterTuningJob
type VertexAIHyperparameterTuningJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIHyperparameterTuningJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIHyperparameterTuningJobObservedState `json:"observedState,omitempty"`
}

// VertexAIHyperparameterTuningJobObservedState is the state of the VertexAIHyperparameterTuningJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.HyperparameterTuningJob
type VertexAIHyperparameterTuningJobObservedState struct {
	// Output only. Trials of the HyperparameterTuningJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.trials
	Trials []TrialObservedState `json:"trials,omitempty"`

	// Output only. The detailed state of the job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.state
	State *string `json:"state,omitempty"`

	// Output only. Time when the HyperparameterTuningJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the HyperparameterTuningJob for the first time
	//  entered the `JOB_STATE_RUNNING` state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the HyperparameterTuningJob entered any of the
	//  following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`,
	//  `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the HyperparameterTuningJob was most recently
	//  updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Only populated when job's state is JOB_STATE_FAILED or
	//  JOB_STATE_CANCELLED.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.HyperparameterTuningJob.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaihyperparametertuningjob;gcpvertexaihyperparametertuningjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIHyperparameterTuningJob is the Schema for the VertexAIHyperparameterTuningJob API
// +k8s:openapi-gen=true
type VertexAIHyperparameterTuningJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIHyperparameterTuningJobSpec   `json:"spec,omitempty"`
	Status VertexAIHyperparameterTuningJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIHyperparameterTuningJobList contains a list of VertexAIHyperparameterTuningJob
type VertexAIHyperparameterTuningJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIHyperparameterTuningJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIHyperparameterTuningJob{}, &VertexAIHyperparameterTuningJobList{})
}
