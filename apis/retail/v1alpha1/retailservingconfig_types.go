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

// API sources for RetailServingConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/retail/v2/serving_config.proto
// +kcc:source:service-docs=https://cloud.google.com/recommendations
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/recommendations/docs/reference/rest/v2/projects.locations.catalogs.servingConfigs

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var RetailServingConfigGVK = GroupVersion.WithKind("RetailServingConfig")

// RetailServingConfigSpec defines the desired state of RetailServingConfig
// +kcc:spec:proto=google.cloud.retail.v2.ServingConfig
// +kcc:required-from-proto
type RetailServingConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The RetailServingConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The human readable serving config display name. Used in Retail
	//  UI.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 128
	//  characters. Otherwise, an INVALID_ARGUMENT error is returned.
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// The id of the model in the same [Catalog][google.cloud.retail.v2.Catalog]
	//  to use at serving time. Currently only RecommendationModels are supported:
	//  https://cloud.google.com/retail/recommendations-ai/docs/create-models
	//  Can be changed but only to a compatible model (e.g.
	//  others-you-may-like CTR to others-you-may-like CVR).
	//
	//  Required when
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_RECOMMENDATION][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_RECOMMENDATION].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.model_id
	ModelID *string `json:"modelID,omitempty"`

	// How much price ranking we want in serving results.
	//  Price reranking causes product items with a similar
	//  recommendation probability to be ordered by price, with the
	//  highest-priced items first. This setting could result in a decrease in
	//  click-through and conversion rates.
	//   Allowed values are:
	//
	//  * `no-price-reranking`
	//  * `low-price-reranking`
	//  * `medium-price-reranking`
	//  * `high-price-reranking`
	//
	//  If not specified, we choose default based on model type. Default value:
	//  `no-price-reranking`.
	//
	//  Can only be set if
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_RECOMMENDATION][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_RECOMMENDATION].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.price_reranking_level
	PriceRerankingLevel *string `json:"priceRerankingLevel,omitempty"`

	// Facet specifications for faceted search. If empty, no facets are returned.
	//  The ids refer to the ids of [Control][google.cloud.retail.v2.Control]
	//  resources with only the Facet control set. These controls are assumed to be
	//  in the same [Catalog][google.cloud.retail.v2.Catalog] as the
	//  [ServingConfig][google.cloud.retail.v2.ServingConfig].
	//  A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error
	//  is returned.
	//
	//  Can only be set if
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.facet_control_ids
	FacetControlIDs []string `json:"facetControlIDs,omitempty"`

	// The specification for dynamically generated facets. Notice that only
	//  textual facets can be dynamically generated.
	//
	//  Can only be set if
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.dynamic_facet_spec
	DynamicFacetSpec *SearchRequest_DynamicFacetSpec `json:"dynamicFacetSpec,omitempty"`

	// Condition boost specifications. If a product matches multiple conditions
	//  in the specifications, boost scores from these specifications are all
	//  applied and combined in a non-linear way. Maximum number of
	//  specifications is 100.
	//
	//  Notice that if both
	//  [ServingConfig.boost_control_ids][google.cloud.retail.v2.ServingConfig.boost_control_ids]
	//  and
	//  [SearchRequest.boost_spec][google.cloud.retail.v2.SearchRequest.boost_spec]
	//  are set, the boost conditions from both places are evaluated. If a search
	//  request matches multiple boost conditions, the final boost score is equal
	//  to the sum of the boost scores from all matched boost conditions.
	//
	//  Can only be set if
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.boost_control_ids
	BoostControlIDs []string `json:"boostControlIDs,omitempty"`

	// Condition filter specifications. If a product matches multiple conditions
	//  in the specifications, filters from these specifications are all
	//  applied and combined via the AND operator. Maximum number of
	//  specifications is 100.
	//
	//  Can only be set if
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.filter_control_ids
	FilterControlIDs []string `json:"filterControlIDs,omitempty"`

	// Condition redirect specifications. Only the first triggered redirect action
	//  is applied, even if multiple apply. Maximum number of specifications is
	//  1000.
	//
	//  Can only be set if
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.redirect_control_ids
	RedirectControlIDs []string `json:"redirectControlIDs,omitempty"`

	// Condition synonyms specifications. If multiple syonyms conditions match,
	//  all matching synonyms control in the list will execute. Order of controls
	//  in the list will not matter. Maximum number of specifications is
	//  100.
	//
	//  Can only be set if
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.twoway_synonyms_control_ids
	TwowaySynonymsControlIDs []string `json:"twowaySynonymsControlIDs,omitempty"`

	// Condition oneway synonyms specifications. If multiple oneway synonyms
	//  conditions match, all matching oneway synonyms controls in the list will
	//  execute. Order of controls in the list will not matter. Maximum number of
	//  specifications is 100.
	//
	//  Can only be set if
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.oneway_synonyms_control_ids
	OnewaySynonymsControlIDs []string `json:"onewaySynonymsControlIDs,omitempty"`

	// Condition do not associate specifications. If multiple do not associate
	//  conditions match, all matching do not associate controls in the list will
	//  execute.
	//  - Order does not matter.
	//  - Maximum number of specifications is 100.
	//
	//  Can only be set if
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.do_not_associate_control_ids
	DoNotAssociateControlIDs []string `json:"doNotAssociateControlIDs,omitempty"`

	// Condition replacement specifications.
	//  - Applied according to the order in the list.
	//  - A previously replaced term can not be re-replaced.
	//  - Maximum number of specifications is 100.
	//
	//  Can only be set if
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.replacement_control_ids
	ReplacementControlIDs []string `json:"replacementControlIDs,omitempty"`

	// Condition ignore specifications. If multiple ignore
	//  conditions match, all matching ignore controls in the list will
	//  execute.
	//  - Order does not matter.
	//  - Maximum number of specifications is 100.
	//
	//  Can only be set if
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.ignore_control_ids
	IgnoreControlIDs []string `json:"ignoreControlIDs,omitempty"`

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
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_RECOMMENDATION][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_RECOMMENDATION].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.diversity_level
	DiversityLevel *string `json:"diversityLevel,omitempty"`

	// What kind of diversity to use - data driven or rule based. If unset, the
	//  server behavior defaults to
	//  [RULE_BASED_DIVERSITY][google.cloud.retail.v2.ServingConfig.DiversityType.RULE_BASED_DIVERSITY].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.diversity_type
	DiversityType *string `json:"diversityType,omitempty"`

	// Whether to add additional category filters on the `similar-items` model.
	//  If not specified, we enable it by default.
	//   Allowed values are:
	//
	//  * `no-category-match`: No additional filtering of original results from
	//    the model and the customer's filters.
	//  * `relaxed-category-match`: Only keep results with categories that match
	//    at least one item categories in the PredictRequests's context item.
	//    * If customer also sends filters in the PredictRequest, then the results
	//    will satisfy both conditions (user given and category match).
	//
	//  Can only be set if
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_RECOMMENDATION][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_RECOMMENDATION].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.enable_category_filter_level
	EnableCategoryFilterLevel *string `json:"enableCategoryFilterLevel,omitempty"`

	// When the flag is enabled, the products in the denylist will not be filtered
	//  out in the recommendation filtering results.
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.ignore_recs_denylist
	IgnoreRecsDenylist *bool `json:"ignoreRecsDenylist,omitempty"`

	// The specification for personalization spec.
	//
	//  Can only be set if
	//  [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_SEARCH].
	//
	//  Notice that if both
	//  [ServingConfig.personalization_spec][google.cloud.retail.v2.ServingConfig.personalization_spec]
	//  and
	//  [SearchRequest.personalization_spec][google.cloud.retail.v2.SearchRequest.personalization_spec]
	//  are set.
	//  [SearchRequest.personalization_spec][google.cloud.retail.v2.SearchRequest.personalization_spec]
	//  will override
	//  [ServingConfig.personalization_spec][google.cloud.retail.v2.ServingConfig.personalization_spec].
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.personalization_spec
	PersonalizationSpec *SearchRequest_PersonalizationSpec `json:"personalizationSpec,omitempty"`

	// Required. Immutable. Specifies the solution types that a serving config can
	//  be associated with. Currently we support setting only one type of solution.
	// +kcc:proto:field=google.cloud.retail.v2.ServingConfig.solution_types
	// +required
	SolutionTypes []string `json:"solutionTypes,omitempty"`
}

// RetailServingConfigStatus defines the config connector machine state of RetailServingConfig
type RetailServingConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the RetailServingConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *RetailServingConfigObservedState `json:"observedState,omitempty"`
}

// RetailServingConfigObservedState is the state of the RetailServingConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.retail.v2.ServingConfig
type RetailServingConfigObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpretailservingconfig;gcpretailservingconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// RetailServingConfig is the Schema for the RetailServingConfig API
// +k8s:openapi-gen=true
type RetailServingConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   RetailServingConfigSpec   `json:"spec,omitempty"`
	Status RetailServingConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RetailServingConfigList contains a list of RetailServingConfig
type RetailServingConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RetailServingConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RetailServingConfig{}, &RetailServingConfigList{})
}
