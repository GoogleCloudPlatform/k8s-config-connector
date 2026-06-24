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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataLabelingEvaluationJobGVK = GroupVersion.WithKind("DataLabelingEvaluationJob")

// DataLabelingEvaluationJobSpec defines the desired state of DataLabelingEvaluationJob
// +kcc:spec:proto=google.cloud.datalabeling.v1beta1.EvaluationJob
type DataLabelingEvaluationJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The DataLabelingEvaluationJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Description of the job. The description can be up to 25,000 characters long.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.description
	Description *string `json:"description,omitempty"`

	// Required. Describes the interval at which the job runs. This interval must be at least 1 day, and it is rounded to the nearest day.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.schedule
	Schedule *string `json:"schedule,omitempty"`

	// Required. The AI Platform Prediction model version to be evaluated.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.model_version
	ModelVersion *string `json:"modelVersion,omitempty"`

	// Required. Configuration details for the evaluation job.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.evaluation_job_config
	EvaluationJobConfig *EvaluationJobConfig `json:"evaluationJobConfig,omitempty"`

	// Required. Name of the AnnotationSpecSet describing all the labels that your machine learning model outputs.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.annotation_spec_set
	AnnotationSpecSet *string `json:"annotationSpecSet,omitempty"`

	// Required. Whether you want Data Labeling Service to provide ground truth labels for prediction input.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.label_missing_ground_truth
	LabelMissingGroundTruth *bool `json:"labelMissingGroundTruth,omitempty"`
}

// DataLabelingEvaluationJobStatus defines the config connector machine state of DataLabelingEvaluationJob
type DataLabelingEvaluationJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataLabelingEvaluationJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataLabelingEvaluationJobObservedState `json:"observedState,omitempty"`
}

// DataLabelingEvaluationJobObservedState is the state of the DataLabelingEvaluationJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.datalabeling.v1beta1.EvaluationJob
type DataLabelingEvaluationJobObservedState struct {
	// Output only. Describes the current state of the job.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.state
	State *string `json:"state,omitempty"`

	// Output only. Every time the evaluation job runs and an error occurs, the failed attempt is appended to this array.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.attempts
	Attempts []Attempt `json:"attempts,omitempty"`

	// Output only. Timestamp of when this evaluation job was created.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.EvaluationJob.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatalabelingevaluationjob;gcpdatalabelingevaluationjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataLabelingEvaluationJob is the Schema for the DataLabelingEvaluationJob API
// +k8s:openapi-gen=true
type DataLabelingEvaluationJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataLabelingEvaluationJobSpec   `json:"spec,omitempty"`
	Status DataLabelingEvaluationJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataLabelingEvaluationJobList contains a list of DataLabelingEvaluationJob
type DataLabelingEvaluationJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataLabelingEvaluationJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataLabelingEvaluationJob{}, &DataLabelingEvaluationJobList{})
}
