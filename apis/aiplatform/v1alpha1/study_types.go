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

var VertexAIStudyGVK = GroupVersion.WithKind("VertexAIStudy")

// VertexAIStudySpec defines the desired state of VertexAIStudy
// +kcc:spec:proto=google.cloud.aiplatform.v1.Study
type VertexAIStudySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The VertexAIStudy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Describes the Study, default value is empty string.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Study.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Configuration of the Study.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Study.study_spec
	StudySpec *StudySpec `json:"studySpec,omitempty"`
}

// VertexAIStudyStatus defines the config connector machine state of VertexAIStudy
type VertexAIStudyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIStudy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIStudyObservedState `json:"observedState,omitempty"`
}

// VertexAIStudyObservedState is the state of the VertexAIStudy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.Study
type VertexAIStudyObservedState struct {
	// Output only. The detailed state of a Study.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Study.state
	State *string `json:"state,omitempty"`

	// Output only. Time at which the study was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Study.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. A human readable reason why the Study is inactive.
	//  This should be empty if a study is ACTIVE or COMPLETED.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Study.inactive_reason
	InactiveReason *string `json:"inactiveReason,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaistudy;gcpvertexaistudys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIStudy is the Schema for the VertexAIStudy API
// +k8s:openapi-gen=true
type VertexAIStudy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIStudySpec   `json:"spec,omitempty"`
	Status VertexAIStudyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIStudyList contains a list of VertexAIStudy
type VertexAIStudyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIStudy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIStudy{}, &VertexAIStudyList{})
}
