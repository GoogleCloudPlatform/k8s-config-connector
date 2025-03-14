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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineServingConfigGVK = GroupVersion.WithKind("DiscoveryEngineServingConfig")

// DiscoveryEngineServingConfigSpec defines the desired state of DiscoveryEngineServingConfig
// +kcc:proto=google.cloud.discoveryengine.v1beta.ServingConfig
type DiscoveryEngineServingConfigSpec struct {
	// Immutable. The location where the TPU virtual machine should reside.
	// +required
	Location string `json:"location,omitempty"`

	// The project that the TPU virtual machine belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// The DiscoveryEngineServingConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.media_config
	MediaConfig *ServingConfig_MediaConfig `json:"mediaConfig,omitempty"`

	// The GenericConfig of the serving configuration.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.generic_config
	GenericConfig *ServingConfig_GenericConfig `json:"genericConfig,omitempty"`

	// Immutable. Fully qualified name
	//  `projects/{project}/locations/{location}/collections/{collection_id}/engines/{engine_id}/servingConfigs/{serving_config_id}`
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.name
	Name *string `json:"name,omitempty"`

	// Required. The human readable serving config display name. Used in Discovery
	//  UI.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 128
	//  characters. Otherwise, an INVALID_ARGUMENT error is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. Specifies the solution type that a serving config can
	//  be associated with.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.solution_type
	SolutionType *string `json:"solutionType,omitempty"`

	// The id of the model to use at serving time.
	//  Currently only RecommendationModels are supported.
	//  Can be changed but only to a compatible model (e.g.
	//  others-you-may-like CTR to others-you-may-like CVR).
	//
	//  Required when
	//  [SolutionType][google.cloud.discoveryengine.v1beta.SolutionType] is
	//  [SOLUTION_TYPE_RECOMMENDATION][google.cloud.discoveryengine.v1beta.SolutionType.SOLUTION_TYPE_RECOMMENDATION].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.model_id
	ModelID *string `json:"modelID,omitempty"`

	// How much diversity to use in recommendation model results e.g.
	//  `medium-diversity` or `high-diversity`. Currently supported values:
	//
	//  * `no-diversity`
	//  * `low-diversity`
	//  * `medium-diversity`
	//  * `high-diversity`
	//  * `auto-diversity`
	//
	//  If not specified, we choose default based on recommendation model
	//  type. Default value: `no-diversity`.
	//
	//  Can only be set if
	//  [SolutionType][google.cloud.discoveryengine.v1beta.SolutionType] is
	//  [SOLUTION_TYPE_RECOMMENDATION][google.cloud.discoveryengine.v1beta.SolutionType.SOLUTION_TYPE_RECOMMENDATION].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.diversity_level
	DiversityLevel *string `json:"diversityLevel,omitempty"`

	// Bring your own embedding config. The config is used for search semantic
	//  retrieval. The retrieval is based on the dot product of
	//  [SearchRequest.EmbeddingSpec.EmbeddingVector.vector][google.cloud.discoveryengine.v1beta.SearchRequest.EmbeddingSpec.EmbeddingVector.vector]
	//  and the document embeddings that are provided by this EmbeddingConfig. If
	//  [SearchRequest.EmbeddingSpec.EmbeddingVector.vector][google.cloud.discoveryengine.v1beta.SearchRequest.EmbeddingSpec.EmbeddingVector.vector]
	//  is provided, it overrides this
	//  [ServingConfig.embedding_config][google.cloud.discoveryengine.v1beta.ServingConfig.embedding_config].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.embedding_config
	EmbeddingConfig *EmbeddingConfig `json:"embeddingConfig,omitempty"`

	// The ranking expression controls the customized ranking on retrieval
	//  documents. To leverage this, document embedding is required. The ranking
	//  expression setting in ServingConfig applies to all search requests served
	//  by the serving config. However, if
	//  [SearchRequest.ranking_expression][google.cloud.discoveryengine.v1beta.SearchRequest.ranking_expression]
	//  is specified, it overrides the ServingConfig ranking expression.
	//
	//  The ranking expression is a single function or multiple functions that are
	//  joined by "+".
	//
	//    * ranking_expression = function, { " + ", function };
	//
	//  Supported functions:
	//
	//    * double * relevance_score
	//    * double * dotProduct(embedding_field_path)
	//
	//  Function variables:
	//
	//    * `relevance_score`: pre-defined keywords, used for measure relevance
	//    between query and document.
	//    * `embedding_field_path`: the document embedding field
	//    used with query embedding vector.
	//    * `dotProduct`: embedding function between embedding_field_path and query
	//    embedding vector.
	//
	//   Example ranking expression:
	//
	//     If document has an embedding field doc_embedding, the ranking expression
	//     could be `0.5 * relevance_score + 0.3 * dotProduct(doc_embedding)`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.ranking_expression
	RankingExpression *string `json:"rankingExpression,omitempty"`

	// Filter controls to use in serving path.
	//  All triggered filter controls will be applied.
	//  Filter controls must be in the same data store as the serving config.
	//  Maximum of 20 filter controls.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.filter_control_ids
	FilterControlIds []string `json:"filterControlIds,omitempty"`

	// Boost controls to use in serving path.
	//  All triggered boost controls will be applied.
	//  Boost controls must be in the same data store as the serving config.
	//  Maximum of 20 boost controls.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.boost_control_ids
	BoostControlIds []string `json:"boostControlIds,omitempty"`

	// IDs of the redirect controls. Only the first triggered redirect
	//  action is applied, even if multiple apply. Maximum number of
	//  specifications is 100.
	//
	//  Can only be set if
	//  [SolutionType][google.cloud.discoveryengine.v1beta.SolutionType] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.discoveryengine.v1beta.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.redirect_control_ids
	RedirectControlIds []string `json:"redirectControlIds,omitempty"`

	// Condition synonyms specifications. If multiple synonyms conditions
	//  match, all matching synonyms controls in the list will execute.
	//  Maximum number of specifications is 100.
	//
	//  Can only be set if
	//  [SolutionType][google.cloud.discoveryengine.v1beta.SolutionType] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.discoveryengine.v1beta.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.synonyms_control_ids
	SynonymsControlIds []string `json:"synonymsControlIds,omitempty"`

	// Condition oneway synonyms specifications. If multiple oneway synonyms
	//  conditions match, all matching oneway synonyms controls in the list
	//  will execute. Maximum number of specifications is 100.
	//
	//  Can only be set if
	//  [SolutionType][google.cloud.discoveryengine.v1beta.SolutionType] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.discoveryengine.v1beta.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.oneway_synonyms_control_ids
	OnewaySynonymsControlIds []string `json:"onewaySynonymsControlIds,omitempty"`

	// Condition do not associate specifications. If multiple do not
	//  associate conditions match, all matching do not associate controls in
	//  the list will execute.
	//  Order does not matter.
	//  Maximum number of specifications is 100.
	//
	//  Can only be set if
	//  [SolutionType][google.cloud.discoveryengine.v1beta.SolutionType] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.discoveryengine.v1beta.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.dissociate_control_ids
	DissociateControlIds []string `json:"dissociateControlIds,omitempty"`

	// Condition replacement specifications.
	//  Applied according to the order in the list.
	//  A previously replaced term can not be re-replaced.
	//  Maximum number of specifications is 100.
	//
	//  Can only be set if
	//  [SolutionType][google.cloud.discoveryengine.v1beta.SolutionType] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.discoveryengine.v1beta.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.replacement_control_ids
	ReplacementControlIds []string `json:"replacementControlIds,omitempty"`

	// Condition ignore specifications. If multiple ignore
	//  conditions match, all matching ignore controls in the list will
	//  execute.
	//  Order does not matter.
	//  Maximum number of specifications is 100.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.ignore_control_ids
	IgnoreControlIds []string `json:"ignoreControlIds,omitempty"`

	// The specification for personalization spec.
	//
	//  Notice that if both
	//  [ServingConfig.personalization_spec][google.cloud.discoveryengine.v1beta.ServingConfig.personalization_spec]
	//  and
	//  [SearchRequest.personalization_spec][google.cloud.discoveryengine.v1beta.SearchRequest.personalization_spec]
	//  are set,
	//  [SearchRequest.personalization_spec][google.cloud.discoveryengine.v1beta.SearchRequest.personalization_spec]
	//  overrides
	//  [ServingConfig.personalization_spec][google.cloud.discoveryengine.v1beta.ServingConfig.personalization_spec].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.personalization_spec
	PersonalizationSpec *SearchRequest_PersonalizationSpec `json:"personalizationSpec,omitempty"`
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
// +kcc:proto=google.cloud.discoveryengine.v1beta.ServingConfig
type DiscoveryEngineServingConfigObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryengineservingconfig;gcpdiscoveryengineservingconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
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
