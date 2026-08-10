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

var VertexAITensorboardRunGVK = GroupVersion.WithKind("VertexAITensorboardRun")

// VertexAITensorboardRunSpec defines the desired state of VertexAITensorboardRun
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.TensorboardRun
type VertexAITensorboardRunSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The Tensorboard that this resource belongs to.
	// +required
	TensorboardRef *VertexAITensorboardRef `json:"tensorboardRef"`

	// Required. The TensorboardExperiment ID.
	// +required
	TensorboardExperimentID string `json:"tensorboardExperimentID"`

	// The VertexAITensorboardRun name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. User provided name of this TensorboardRun.
	// This value must be unique among all TensorboardRuns belonging to the same parent TensorboardExperiment.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.TensorboardRun.display_name
	// +required
	DisplayName string `json:"displayName"`

	// Description of this TensorboardRun.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.TensorboardRun.description
	Description *string `json:"description,omitempty"`

	// The labels with user-defined metadata to organize your TensorboardRuns.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.TensorboardRun.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// VertexAITensorboardRunStatus defines the config connector machine state of VertexAITensorboardRun
type VertexAITensorboardRunStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAITensorboardRun resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAITensorboardRunObservedState `json:"observedState,omitempty"`
}

// VertexAITensorboardRunObservedState is the state of the VertexAITensorboardRun resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.TensorboardRun
type VertexAITensorboardRunObservedState struct {
	// Output only. Name of the TensorboardRun.
	// Format: `projects/{project}/locations/{location}/tensorboards/{tensorboard}/experiments/{experiment}/runs/{run}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.TensorboardRun.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this TensorboardRun was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.TensorboardRun.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this TensorboardRun was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.TensorboardRun.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaitensorboardrun;gcpvertexaitensorboardruns
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAITensorboardRun is the Schema for the VertexAITensorboardRun API
// +k8s:openapi-gen=true
type VertexAITensorboardRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAITensorboardRunSpec   `json:"spec,omitempty"`
	Status VertexAITensorboardRunStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAITensorboardRunList contains a list of VertexAITensorboardRun
type VertexAITensorboardRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAITensorboardRun `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAITensorboardRun{}, &VertexAITensorboardRunList{})
}
