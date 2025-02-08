// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.retail.v2beta.AttributesConfig
type AttributesConfig struct {
	// Required. Immutable. The fully qualified resource name of the attribute
	//  config. Format: `projects/*/locations/*/catalogs/*/attributesConfig`
	// +kcc:proto:field=google.cloud.retail.v2beta.AttributesConfig.name
	Name *string `json:"name,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.retail.v2beta.CatalogAttribute
type CatalogAttribute struct {
	// Required. Attribute name.
	//  For example: `color`, `brands`, `attributes.custom_attribute`, such as
	//  `attributes.xyz`.
	//  To be indexable, the attribute name can contain only alpha-numeric
	//  characters and underscores. For example, an attribute named
	//  `attributes.abc_xyz` can be indexed, but an attribute named
	//  `attributes.abc-xyz` cannot be indexed.
	//
	//  If the attribute key starts with `attributes.`, then the attribute is a
	//  custom attribute. Attributes such as `brands`, `patterns`, and `title` are
	//  built-in and called system attributes.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.key
	Key *string `json:"key,omitempty"`

	// When
	//  [AttributesConfig.attribute_config_level][google.cloud.retail.v2beta.AttributesConfig.attribute_config_level]
	//  is CATALOG_LEVEL_ATTRIBUTE_CONFIG, if INDEXABLE_ENABLED attribute values
	//  are indexed so that it can be filtered, faceted, or boosted in
	//  [SearchService.Search][google.cloud.retail.v2beta.SearchService.Search].
	//
	//  Must be specified when
	//  [AttributesConfig.attribute_config_level][google.cloud.retail.v2beta.AttributesConfig.attribute_config_level]
	//  is CATALOG_LEVEL_ATTRIBUTE_CONFIG, otherwise throws INVALID_FORMAT error.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.indexable_option
	IndexableOption *string `json:"indexableOption,omitempty"`

	// If DYNAMIC_FACETABLE_ENABLED, attribute values are available for dynamic
	//  facet. Could only be DYNAMIC_FACETABLE_DISABLED if
	//  [CatalogAttribute.indexable_option][google.cloud.retail.v2beta.CatalogAttribute.indexable_option]
	//  is INDEXABLE_DISABLED. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	//  Must be specified, otherwise throws INVALID_FORMAT error.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.dynamic_facetable_option
	DynamicFacetableOption *string `json:"dynamicFacetableOption,omitempty"`

	// When
	//  [AttributesConfig.attribute_config_level][google.cloud.retail.v2beta.AttributesConfig.attribute_config_level]
	//  is CATALOG_LEVEL_ATTRIBUTE_CONFIG, if SEARCHABLE_ENABLED, attribute values
	//  are searchable by text queries in
	//  [SearchService.Search][google.cloud.retail.v2beta.SearchService.Search].
	//
	//  If SEARCHABLE_ENABLED but attribute type is numerical, attribute values
	//  will not be searchable by text queries in
	//  [SearchService.Search][google.cloud.retail.v2beta.SearchService.Search], as
	//  there are no text values associated to numerical attributes.
	//
	//  Must be specified, when
	//  [AttributesConfig.attribute_config_level][google.cloud.retail.v2beta.AttributesConfig.attribute_config_level]
	//  is CATALOG_LEVEL_ATTRIBUTE_CONFIG, otherwise throws INVALID_FORMAT error.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.searchable_option
	SearchableOption *string `json:"searchableOption,omitempty"`

	// When
	//  [AttributesConfig.attribute_config_level][google.cloud.retail.v2beta.AttributesConfig.attribute_config_level]
	//  is CATALOG_LEVEL_ATTRIBUTE_CONFIG, if RECOMMENDATIONS_FILTERING_ENABLED,
	//  attribute values are filterable for recommendations.
	//  This option works for categorical features only,
	//  does not work for numerical features, inventory filtering.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.recommendations_filtering_option
	RecommendationsFilteringOption *string `json:"recommendationsFilteringOption,omitempty"`

	// If EXACT_SEARCHABLE_ENABLED, attribute values will be exact searchable.
	//  This property only applies to textual custom attributes and requires
	//  indexable set to enabled to enable exact-searchable. If unset, the server
	//  behavior defaults to
	//  [EXACT_SEARCHABLE_DISABLED][google.cloud.retail.v2beta.CatalogAttribute.ExactSearchableOption.EXACT_SEARCHABLE_DISABLED].
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.exact_searchable_option
	ExactSearchableOption *string `json:"exactSearchableOption,omitempty"`

	// If RETRIEVABLE_ENABLED, attribute values are retrievable in the search
	//  results. If unset, the server behavior defaults to
	//  [RETRIEVABLE_DISABLED][google.cloud.retail.v2beta.CatalogAttribute.RetrievableOption.RETRIEVABLE_DISABLED].
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.retrievable_option
	RetrievableOption *string `json:"retrievableOption,omitempty"`

	// Contains facet options.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.facet_config
	FacetConfig *CatalogAttribute_FacetConfig `json:"facetConfig,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig
type CatalogAttribute_FacetConfig struct {
	// If you don't set the facet
	//  [SearchRequest.FacetSpec.FacetKey.intervals][google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.intervals]
	//  in the request to a numerical attribute, then we use the computed
	//  intervals with rounded bounds obtained from all its product numerical
	//  attribute values. The computed intervals might not be ideal for some
	//  attributes. Therefore, we give you the option to overwrite them with the
	//  facet_intervals field. The maximum of facet intervals per
	//  [CatalogAttribute][google.cloud.retail.v2beta.CatalogAttribute] is 40.
	//  Each interval must have a lower bound or an upper bound. If both bounds
	//  are provided, then the lower bound must be smaller or equal than the
	//  upper bound.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.facet_intervals
	FacetIntervals []Interval `json:"facetIntervals,omitempty"`

	// Each instance represents a list of attribute values to ignore as facet
	//  values for a specific time range. The maximum number of instances per
	//  [CatalogAttribute][google.cloud.retail.v2beta.CatalogAttribute] is 25.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.ignored_facet_values
	IgnoredFacetValues []CatalogAttribute_FacetConfig_IgnoredFacetValues `json:"ignoredFacetValues,omitempty"`

	// Each instance replaces a list of facet values by a merged facet
	//  value. If a facet value is not in any list, then it will stay the same.
	//  To avoid conflicts, only paths of length 1 are accepted. In other words,
	//  if "dark_blue" merged into "BLUE", then the latter can't merge into
	//  "blues" because this would create a path of length 2. The maximum number
	//  of instances of MergedFacetValue per
	//  [CatalogAttribute][google.cloud.retail.v2beta.CatalogAttribute] is 100.
	//  This feature is available only for textual custom attributes.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.merged_facet_values
	MergedFacetValues []CatalogAttribute_FacetConfig_MergedFacetValue `json:"mergedFacetValues,omitempty"`

	// Use this field only if you want to merge a facet key into another facet
	//  key.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.merged_facet
	MergedFacet *CatalogAttribute_FacetConfig_MergedFacet `json:"mergedFacet,omitempty"`

	// Set this field only if you want to rerank based on facet values engaged
	//  by the user for the current key. This option is only possible for custom
	//  facetable textual keys.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.rerank_config
	RerankConfig *CatalogAttribute_FacetConfig_RerankConfig `json:"rerankConfig,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.IgnoredFacetValues
type CatalogAttribute_FacetConfig_IgnoredFacetValues struct {
	// List of facet values to ignore for the following time range. The facet
	//  values are the same as the attribute values. There is a limit of 10
	//  values per instance of IgnoredFacetValues. Each value can have at most
	//  128 characters.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.IgnoredFacetValues.values
	Values []string `json:"values,omitempty"`

	// Time range for the current list of facet values to ignore.
	//  If multiple time ranges are specified for an facet value for the
	//  current attribute, consider all of them. If both are empty, ignore
	//  always. If start time and end time are set, then start time
	//  must be before end time.
	//  If start time is not empty and end time is empty, then will ignore
	//  these facet values after the start time.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.IgnoredFacetValues.start_time
	StartTime *string `json:"startTime,omitempty"`

	// If start time is empty and end time is not empty, then ignore these
	//  facet values before end time.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.IgnoredFacetValues.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.MergedFacet
type CatalogAttribute_FacetConfig_MergedFacet struct {
	// The merged facet key should be a valid facet key that is different than
	//  the facet key of the current catalog attribute. We refer this is
	//  merged facet key as the child of the current catalog attribute. This
	//  merged facet key can't be a parent of another facet key (i.e. no
	//  directed path of length 2). This merged facet key needs to be either a
	//  textual custom attribute or a numerical custom attribute.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.MergedFacet.merged_facet_key
	MergedFacetKey *string `json:"mergedFacetKey,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.MergedFacetValue
type CatalogAttribute_FacetConfig_MergedFacetValue struct {
	// All the facet values that are replaces by the same
	//  [merged_value][google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.MergedFacetValue.merged_value]
	//  that follows. The maximum number of values per MergedFacetValue is 25.
	//  Each value can have up to 128 characters.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.MergedFacetValue.values
	Values []string `json:"values,omitempty"`

	// All the previous values are replaced by this merged facet value.
	//  This merged_value must be non-empty and can have up to 128 characters.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.MergedFacetValue.merged_value
	MergedValue *string `json:"mergedValue,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.RerankConfig
type CatalogAttribute_FacetConfig_RerankConfig struct {
	// If set to true, then we also rerank the dynamic facets based on the
	//  facet values engaged by the user for the current attribute key during
	//  serving.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.RerankConfig.rerank_facet
	RerankFacet *bool `json:"rerankFacet,omitempty"`

	// If empty, rerank on all facet values for the current key. Otherwise,
	//  will rerank on the facet values from this list only.
	// +kcc:proto:field=google.cloud.retail.v2beta.CatalogAttribute.FacetConfig.RerankConfig.facet_values
	FacetValues []string `json:"facetValues,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Interval
type Interval struct {
	// Inclusive lower bound.
	// +kcc:proto:field=google.cloud.retail.v2beta.Interval.minimum
	Minimum *float64 `json:"minimum,omitempty"`

	// Exclusive lower bound.
	// +kcc:proto:field=google.cloud.retail.v2beta.Interval.exclusive_minimum
	ExclusiveMinimum *float64 `json:"exclusiveMinimum,omitempty"`

	// Inclusive upper bound.
	// +kcc:proto:field=google.cloud.retail.v2beta.Interval.maximum
	Maximum *float64 `json:"maximum,omitempty"`

	// Exclusive upper bound.
	// +kcc:proto:field=google.cloud.retail.v2beta.Interval.exclusive_maximum
	ExclusiveMaximum *float64 `json:"exclusiveMaximum,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.AttributesConfig
type AttributesConfigObservedState struct {
	// Output only. The
	//  [AttributeConfigLevel][google.cloud.retail.v2beta.AttributeConfigLevel]
	//  used for this catalog.
	// +kcc:proto:field=google.cloud.retail.v2beta.AttributesConfig.attribute_config_level
	AttributeConfigLevel *string `json:"attributeConfigLevel,omitempty"`
}
