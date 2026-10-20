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

// API sources for VisionAIIndex, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/visionai/v1/warehouse.proto
// +kcc:source:service-docs=https://cloud.google.com/vision-ai/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/vision-ai/docs/reference/rest/v1/projects.locations.corpora.indexes

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VisionAIIndexGVK = GroupVersion.WithKind("VisionAIIndex")

// VisionAIIndexSpec defines the desired state of VisionAIIndex
// +kcc:spec:proto=google.cloud.visionai.v1.Index
// +kcc:required-from-proto
type VisionAIIndexSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project_number}/locations/{location}/corpora/{corpus} this resource belongs to.
	// +kcc:guess
	// CorporaRef *VisionAICorpusRef `json:"corporaRef,omitempty"`

	// The VisionAIIndex name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Include all assets under the corpus.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.entire_corpus
	EntireCorpus *bool `json:"entireCorpus,omitempty"`

	// Optional. Optional user-specified display name of the index.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Optional description of the index.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.description
	Description *string `json:"description,omitempty"`
}

// VisionAIIndexStatus defines the config connector machine state of VisionAIIndex
type VisionAIIndexStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VisionAIIndex resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VisionAIIndexObservedState `json:"observedState,omitempty"`
}

// VisionAIIndexObservedState is the state of the VisionAIIndex resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.visionai.v1.Index
type VisionAIIndexObservedState struct {
	// Output only. State of the index.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.state
	State *string `json:"state,omitempty"`

	// Output only. The create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. References to the deployed index instance.
	//  Index of VIDEO_ON_DEMAND corpus can have at most one deployed index.
	//  Index of IMAGE corpus can have multiple deployed indexes.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.deployed_indexes
	DeployedIndexes []DeployedIndexReference `json:"deployedIndexes,omitempty"`

	// Output only. This boolean field is only set for projects that have Physical
	//  Zone Separation enabled via an Org Policy constraint. It is set to true
	//  when the index is a valid zone separated index and false if it isn't.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. This boolean field is only set for projects that have Physical
	//  Zone Isolation enabled via an Org Policy constraint. It is set to true when
	//  the index is a valid zone isolated index and false if it isn't.
	// +kcc:proto:field=google.cloud.visionai.v1.Index.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvisionaiindex;gcpvisionaiindexes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VisionAIIndex is the Schema for the VisionAIIndex API
// +k8s:openapi-gen=true
type VisionAIIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VisionAIIndexSpec   `json:"spec,omitempty"`
	Status VisionAIIndexStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VisionAIIndexList contains a list of VisionAIIndex
type VisionAIIndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VisionAIIndex `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VisionAIIndex{}, &VisionAIIndexList{})
}
