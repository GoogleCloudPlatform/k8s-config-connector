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

// API sources for AIPlatformTensorboardRun, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/aiplatform/v1/tensorboard_run.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.tensorboards.experiments.runs

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformTensorboardRunGVK = GroupVersion.WithKind("AIPlatformTensorboardRun")

// AIPlatformTensorboardRunSpec defines the desired state of AIPlatformTensorboardRun
// +kcc:spec:proto=google.cloud.aiplatform.v1.TensorboardRun
// +kcc:required-from-proto
type AIPlatformTensorboardRunSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AIPlatformTensorboardRun name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. User provided name of this TensorboardRun.
	//  This value must be unique among all TensorboardRuns
	//  belonging to the same parent TensorboardExperiment.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardRun.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Description of this TensorboardRun.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardRun.description
	Description *string `json:"description,omitempty"`

	// The labels with user-defined metadata to organize your TensorboardRuns.
	//
	//  This field will be used to filter and visualize Runs in the Tensorboard UI.
	//  For example, a Vertex AI training job can set a label
	//  aiplatform.googleapis.com/training_job_id=xxxxx to all the runs created
	//  within that job. An end user can set a label experiment_id=xxxxx for all
	//  the runs produced in a Jupyter notebook. These runs can be grouped by a
	//  label value and visualized together in the Tensorboard UI.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one TensorboardRun
	//  (System labels are excluded).
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//  System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	//  and are immutable.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardRun.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Used to perform a consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardRun.etag
	Etag *string `json:"etag,omitempty"`
}

// AIPlatformTensorboardRunStatus defines the config connector machine state of AIPlatformTensorboardRun
type AIPlatformTensorboardRunStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformTensorboardRun resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformTensorboardRunObservedState `json:"observedState,omitempty"`
}

// AIPlatformTensorboardRunObservedState is the state of the AIPlatformTensorboardRun resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.TensorboardRun
type AIPlatformTensorboardRunObservedState struct {
	// Output only. Timestamp when this TensorboardRun was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardRun.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this TensorboardRun was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardRun.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformtensorboardrun;gcpaiplatformtensorboardruns
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformTensorboardRun is the Schema for the AIPlatformTensorboardRun API
// +k8s:openapi-gen=true
type AIPlatformTensorboardRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformTensorboardRunSpec   `json:"spec,omitempty"`
	Status AIPlatformTensorboardRunStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformTensorboardRunList contains a list of AIPlatformTensorboardRun
type AIPlatformTensorboardRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformTensorboardRun `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformTensorboardRun{}, &AIPlatformTensorboardRunList{})
}
