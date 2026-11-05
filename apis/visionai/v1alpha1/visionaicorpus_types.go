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

// API sources for VisionAICorpus, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/visionai/v1/warehouse.proto
// +kcc:source:service-docs=https://cloud.google.com/vision-ai/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/vision-ai/docs/reference/rest/v1/projects.locations.corpora

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VisionAICorpusGVK = GroupVersion.WithKind("VisionAICorpus")

// VisionAICorpusSpec defines the desired state of VisionAICorpus
// +kcc:spec:proto=google.cloud.visionai.v1.Corpus
// +kcc:required-from-proto
type VisionAICorpusSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VisionAICorpus name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The corpus name to shown in the UI. The name can be up to 32
	//  characters long.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the corpus. Can be up to 25000 characters long.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.description
	Description *string `json:"description,omitempty"`

	// Optional. The default TTL value for all assets under the corpus without a
	//  asset level user-defined TTL. For STREAM_VIDEO type corpora, this is
	//  required and the maximum allowed
	//    default_ttl is 10 years.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.default_ttl
	DefaultTTL *string `json:"defaultTTL,omitempty"`

	// Optional. Type of the asset inside corpus.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.type
	Type *string `json:"type,omitempty"`

	// Default search capability setting on corpus level.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.search_capability_setting
	SearchCapabilitySetting *SearchCapabilitySetting `json:"searchCapabilitySetting,omitempty"`
}

// VisionAICorpusStatus defines the config connector machine state of VisionAICorpus
type VisionAICorpusStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VisionAICorpus resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VisionAICorpusObservedState `json:"observedState,omitempty"`
}

// VisionAICorpusObservedState is the state of the VisionAICorpus resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.visionai.v1.Corpus
type VisionAICorpusObservedState struct {
	// Output only. This boolean field is only set for projects that have Physical
	//  Zone Separation enabled via an Org Policy constraint. It is set to true
	//  when the corpus is a valid zone separated corpus and false if it isn't.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. This boolean field is only set for projects that have Physical
	//  Zone Isolation enabled via an Org Policy constraint. It is set to true when
	//  the corpus is a valid zone isolated corpus and false if it isn't.
	// +kcc:proto:field=google.cloud.visionai.v1.Corpus.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvisionaicorpus;gcpvisionaicorpora
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VisionAICorpus is the Schema for the VisionAICorpus API
// +k8s:openapi-gen=true
type VisionAICorpus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VisionAICorpusSpec   `json:"spec,omitempty"`
	Status VisionAICorpusStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VisionAICorpusList contains a list of VisionAICorpus
type VisionAICorpusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VisionAICorpus `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VisionAICorpus{}, &VisionAICorpusList{})
}
