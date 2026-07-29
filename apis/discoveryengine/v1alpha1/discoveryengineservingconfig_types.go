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

var DiscoveryEngineServingConfigGVK = GroupVersion.WithKind("DiscoveryEngineServingConfig")

// DiscoveryEngineServingConfigSpec defines the desired state of DiscoveryEngineServingConfig
// +kcc:spec:proto=google.cloud.discoveryengine.v1beta.ServingConfig
type DiscoveryEngineServingConfigSpec struct {
	// Immutable. The Project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. Location of the resource.
	// +required
	Location string `json:"location"`

	// Immutable. The Engine this serving config belongs to.
	// +required
	EngineRef *DiscoveryEngineEngineRef `json:"engineRef"`

	// Immutable.
	// The DiscoveryEngineServingConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The human readable serving config display name. Used in Discovery UI.
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. Specifies the solution type that a serving config can be associated with.
	// +required
	SolutionType *string `json:"solutionType,omitempty"`

	// The id of the model to use at serving time.
	ModelID *string `json:"modelID,omitempty"`

	// How much diversity to use in recommendation model results.
	DiversityLevel *string `json:"diversityLevel,omitempty"`

	// Bring your own embedding config.
	EmbeddingConfig *EmbeddingConfig `json:"embeddingConfig,omitempty"`

	// The ranking expression controls the customized ranking on retrieval documents.
	RankingExpression *string `json:"rankingExpression,omitempty"`

	// Filter controls to use in serving path.
	FilterControlIDs []string `json:"filterControlIDs,omitempty"`

	// Boost controls to use in serving path.
	BoostControlIDs []string `json:"boostControlIDs,omitempty"`

	// IDs of the redirect controls.
	RedirectControlIDs []string `json:"redirectControlIDs,omitempty"`

	// Condition synonyms specifications.
	SynonymsControlIDs []string `json:"synonymsControlIDs,omitempty"`

	// Condition oneway synonyms specifications.
	OnewaySynonymsControlIDs []string `json:"onewaySynonymsControlIDs,omitempty"`

	// Condition do not associate specifications.
	DissociateControlIDs []string `json:"dissociateControlIDs,omitempty"`

	// Condition replacement specifications.
	ReplacementControlIDs []string `json:"replacementControlIDs,omitempty"`

	// Condition ignore specifications.
	IgnoreControlIDs []string `json:"ignoreControlIDs,omitempty"`

	// The specification for personalization spec.
	PersonalizationSpec *SearchRequest_PersonalizationSpec `json:"personalizationSpec,omitempty"`

	// The GenericConfig of the serving configuration.
	GenericConfig *ServingConfig_GenericConfig `json:"genericConfig,omitempty"`

	// The MediaConfig of the serving configuration.
	MediaConfig *ServingConfig_MediaConfig `json:"mediaConfig,omitempty"`
}

// DiscoveryEngineServingConfigStatus defines the config connector machine state of DiscoveryEngineServingConfig
type DiscoveryEngineServingConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineServingConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineServingConfigObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineServingConfigObservedState is the state of the DiscoveryEngineServingConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1beta.ServingConfig
type DiscoveryEngineServingConfigObservedState struct {
	// Output only. ServingConfig created timestamp.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. ServingConfig updated timestamp.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryengineservingconfig;gcpdiscoveryengineservingconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineServingConfig is the Schema for the DiscoveryEngineServingConfig API
// +k8s:openapi-gen=true
type DiscoveryEngineServingConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineServingConfigSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineServingConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineServingConfigList contains a list of DiscoveryEngineServingConfig
type DiscoveryEngineServingConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineServingConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineServingConfig{}, &DiscoveryEngineServingConfigList{})
}
