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

// API sources for AIPlatformNasJob, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/aiplatform/v1/nas_job.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:source:resource-docs=https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.nasJobs

package v1alpha1

import (
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformNasJobGVK = GroupVersion.WithKind("AIPlatformNasJob")

// AIPlatformNasJobSpec defines the desired state of AIPlatformNasJob
// +kcc:spec:proto=google.cloud.aiplatform.v1.NasJob
// +kcc:required-from-proto
type AIPlatformNasJobSpec struct {
	// Inlined parent reference for Project and Location
	*parent.ProjectAndLocationRef `json:",inline"`

	// The AIPlatformNasJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the NasJob.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The specification of a NasJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.nas_job_spec
	// +required
	NasJobSpec *NasJobSpec `json:"nasJobSpec,omitempty"`

	// The labels with user-defined metadata to organize NasJobs.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key options for a NasJob.
	//  If this is set, then all resources created by the NasJob
	//  will be encrypted with the provided encryption key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobSpec
type NasJobSpec struct {
	// The spec of multi-trial algorithms.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.multi_trial_algorithm_spec
	MultiTrialAlgorithmSpec *NasJobSpec_MultiTrialAlgorithmSpec `json:"multiTrialAlgorithmSpec,omitempty"`

	// The ID of the existing NasJob in the same Project and Location
	//  which will be used to resume search. search_space_spec and
	//  nas_algorithm_spec are obtained from previous NasJob hence should not
	//  provide them again for this NasJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.resume_nas_job_id
	ResumeNasJobRef *AIPlatformNasJobRef `json:"resumeNasJobRef,omitempty"`

	// It defines the search space for Neural Architecture Search (NAS).
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.search_space_spec
	SearchSpaceSpec *string `json:"searchSpaceSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec
type NasJobSpec_MultiTrialAlgorithmSpec struct {
	// The multi-trial Neural Architecture Search (NAS) algorithm
	//  type. Defaults to `REINFORCEMENT_LEARNING`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.multi_trial_algorithm
	MultiTrialAlgorithm *string `json:"multiTrialAlgorithm,omitempty"`

	// Metric specs for the NAS job.
	//  Validation for this field is done at `multi_trial_algorithm_spec` field.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.metric
	Metric *NasJobSpec_MultiTrialAlgorithmSpec_MetricSpec `json:"metric,omitempty"`

	// Required. Spec for search trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.search_trial_spec
	// +required
	SearchTrialSpec *NasJobSpec_MultiTrialAlgorithmSpec_SearchTrialSpec `json:"searchTrialSpec,omitempty"`

	// Spec for train trials. Top N [TrainTrialSpec.max_parallel_trial_count]
	//  search trials will be trained for every M
	//  [TrainTrialSpec.frequency] trials searched.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.train_trial_spec
	TrainTrialSpec *NasJobSpec_MultiTrialAlgorithmSpec_TrainTrialSpec `json:"trainTrialSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.MetricSpec
type NasJobSpec_MultiTrialAlgorithmSpec_MetricSpec struct {
	// Required. The ID of the metric. Must not contain whitespaces.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.MetricSpec.metric_id
	// +required
	MetricID *string `json:"metricID,omitempty"`

	// Required. The optimization goal of the metric.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.MetricSpec.goal
	// +required
	Goal *string `json:"goal,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.SearchTrialSpec
type NasJobSpec_MultiTrialAlgorithmSpec_SearchTrialSpec struct {
	// Required. The spec of a search trial job. The same spec applies to
	//  all search trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.SearchTrialSpec.search_trial_job_spec
	// +required
	SearchTrialJobSpec *CustomJobSpecWithRequired `json:"searchTrialJobSpec,omitempty"`

	// Required. The maximum number of Neural Architecture Search (NAS) trials
	//  to run.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.SearchTrialSpec.max_trial_count
	// +required
	MaxTrialCount *int32 `json:"maxTrialCount,omitempty"`

	// Required. The maximum number of trials to run in parallel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.SearchTrialSpec.max_parallel_trial_count
	// +required
	MaxParallelTrialCount *int32 `json:"maxParallelTrialCount,omitempty"`

	// The number of failed trials that need to be seen before failing
	//  the NasJob.
	//
	//  If set to 0, Vertex AI decides how many trials must fail
	//  before the whole job fails.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.SearchTrialSpec.max_failed_trial_count
	MaxFailedTrialCount *int32 `json:"maxFailedTrialCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.TrainTrialSpec
type NasJobSpec_MultiTrialAlgorithmSpec_TrainTrialSpec struct {
	// Required. The spec of a train trial job. The same spec applies to
	//  all train trials.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.TrainTrialSpec.train_trial_job_spec
	// +required
	TrainTrialJobSpec *CustomJobSpecWithRequired `json:"trainTrialJobSpec,omitempty"`

	// Required. The maximum number of trials to run in parallel.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.TrainTrialSpec.max_parallel_trial_count
	// +required
	MaxParallelTrialCount *int32 `json:"maxParallelTrialCount,omitempty"`

	// Required. Frequency of search trials to start train stage. Top N
	//  [TrainTrialSpec.max_parallel_trial_count]
	//  search trials will be trained for every M
	//  [TrainTrialSpec.frequency] trials searched.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJobSpec.MultiTrialAlgorithmSpec.TrainTrialSpec.frequency
	// +required
	Frequency *int32 `json:"frequency,omitempty"`
}

// AIPlatformNasJobStatus defines the config connector machine state of AIPlatformNasJob
type AIPlatformNasJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformNasJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformNasJobObservedState `json:"observedState,omitempty"`
}

// AIPlatformNasJobObservedState is the state of the AIPlatformNasJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.NasJob
type AIPlatformNasJobObservedState struct {
	// Output only. Output of the NasJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.nas_job_output
	NasJobOutput *NasJobOutputObservedState `json:"nasJobOutput,omitempty"`

	// Output only. The detailed state of the job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.state
	State *string `json:"state,omitempty"`

	// Output only. Time when the NasJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the NasJob for the first time entered the
	//  `JOB_STATE_RUNNING` state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the NasJob entered any of the following states:
	//  `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the NasJob was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Only populated when job's state is JOB_STATE_FAILED or
	//  JOB_STATE_CANCELLED.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformnasjob;gcpaiplatformnasjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformNasJob is the Schema for the AIPlatformNasJob API
// +k8s:openapi-gen=true
type AIPlatformNasJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformNasJobSpec   `json:"spec,omitempty"`
	Status AIPlatformNasJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformNasJobList contains a list of AIPlatformNasJob
type AIPlatformNasJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformNasJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformNasJob{}, &AIPlatformNasJobList{})
}
