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


// +kcc:proto=google.cloud.retail.v2beta.Condition
type Condition struct {
	// A list (up to 10 entries) of terms to match the query on. If not
	//  specified, match all queries.
	//  If many query terms are specified, the condition
	//  is matched if any of the terms is a match (i.e. using the OR operator).
	// +kcc:proto:field=google.cloud.retail.v2beta.Condition.query_terms
	QueryTerms []Condition_QueryTerm `json:"queryTerms,omitempty"`

	// Range of time(s) specifying when Condition is active.
	//  Condition true if any time range matches.
	// +kcc:proto:field=google.cloud.retail.v2beta.Condition.active_time_range
	ActiveTimeRange []Condition_TimeRange `json:"activeTimeRange,omitempty"`

	// Used to support browse uses cases.
	//  A list (up to 10 entries) of categories or departments.
	//  The format should be the same as
	//  [UserEvent.page_categories][google.cloud.retail.v2beta.UserEvent.page_categories];
	// +kcc:proto:field=google.cloud.retail.v2beta.Condition.page_categories
	PageCategories []string `json:"pageCategories,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Condition.QueryTerm
type Condition_QueryTerm struct {
	// The value of the term to match on.
	//  Value cannot be empty.
	//  Value can have at most 3 terms if specified as a partial match. Each
	//  space separated string is considered as one term.
	//  For example, "a b c" is 3 terms and allowed, but " a b c d" is 4 terms
	//  and not allowed for a partial match.
	// +kcc:proto:field=google.cloud.retail.v2beta.Condition.QueryTerm.value
	Value *string `json:"value,omitempty"`

	// Whether this is supposed to be a full or partial match.
	// +kcc:proto:field=google.cloud.retail.v2beta.Condition.QueryTerm.full_match
	FullMatch *bool `json:"fullMatch,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Condition.TimeRange
type Condition_TimeRange struct {
	// Start of time range. Range is inclusive.
	// +kcc:proto:field=google.cloud.retail.v2beta.Condition.TimeRange.start_time
	StartTime *string `json:"startTime,omitempty"`

	// End of time range. Range is inclusive.
	// +kcc:proto:field=google.cloud.retail.v2beta.Condition.TimeRange.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Control
type Control struct {
	// A facet specification to perform faceted search.
	//
	//  Note that this field is deprecated and will throw NOT_IMPLEMENTED if
	//  used for creating a control.
	// +kcc:proto:field=google.cloud.retail.v2beta.Control.facet_spec
	FacetSpec *SearchRequest_FacetSpec `json:"facetSpec,omitempty"`

	// A rule control - a condition-action pair.
	//  Enacts a set action when the condition is triggered.
	//  For example: Boost "gShoe" when query full matches "Running Shoes".
	// +kcc:proto:field=google.cloud.retail.v2beta.Control.rule
	Rule *Rule `json:"rule,omitempty"`

	// Immutable. Fully qualified name
	//  `projects/*/locations/global/catalogs/*/controls/*`
	// +kcc:proto:field=google.cloud.retail.v2beta.Control.name
	Name *string `json:"name,omitempty"`

	// Required. The human readable control display name. Used in Retail UI.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 128
	//  characters. Otherwise, an INVALID_ARGUMENT error is thrown.
	// +kcc:proto:field=google.cloud.retail.v2beta.Control.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. The solution types that the control is used for.
	//  Currently we support setting only one type of solution at creation time.
	//
	//  Only `SOLUTION_TYPE_SEARCH` value is supported at the moment.
	//  If no solution type is provided at creation time, will default to
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2beta.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.retail.v2beta.Control.solution_types
	SolutionTypes []string `json:"solutionTypes,omitempty"`

	// Specifies the use case for the control.
	//  Affects what condition fields can be set.
	//  Only settable by search controls.
	//  Will default to
	//  [SEARCH_SOLUTION_USE_CASE_SEARCH][google.cloud.retail.v2beta.SearchSolutionUseCase.SEARCH_SOLUTION_USE_CASE_SEARCH]
	//  if not specified. Currently only allow one search_solution_use_case per
	//  control.
	// +kcc:proto:field=google.cloud.retail.v2beta.Control.search_solution_use_case
	SearchSolutionUseCase []string `json:"searchSolutionUseCase,omitempty"`
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

// +kcc:proto=google.cloud.retail.v2beta.Rule
type Rule struct {
	// A boost action.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.boost_action
	BoostAction *Rule_BoostAction `json:"boostAction,omitempty"`

	// Redirects a shopper to a specific page.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.redirect_action
	RedirectAction *Rule_RedirectAction `json:"redirectAction,omitempty"`

	// Treats specific term as a synonym with a group of terms.
	//  Group of terms will not be treated as synonyms with the specific term.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.oneway_synonyms_action
	OnewaySynonymsAction *Rule_OnewaySynonymsAction `json:"onewaySynonymsAction,omitempty"`

	// Prevents term from being associated with other terms.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.do_not_associate_action
	DoNotAssociateAction *Rule_DoNotAssociateAction `json:"doNotAssociateAction,omitempty"`

	// Replaces specific terms in the query.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.replacement_action
	ReplacementAction *Rule_ReplacementAction `json:"replacementAction,omitempty"`

	// Ignores specific terms from query during search.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.ignore_action
	IgnoreAction *Rule_IgnoreAction `json:"ignoreAction,omitempty"`

	// Filters results.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.filter_action
	FilterAction *Rule_FilterAction `json:"filterAction,omitempty"`

	// Treats a set of terms as synonyms of one another.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.twoway_synonyms_action
	TwowaySynonymsAction *Rule_TwowaySynonymsAction `json:"twowaySynonymsAction,omitempty"`

	// Force returns an attribute as a facet in the request.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.force_return_facet_action
	ForceReturnFacetAction *Rule_ForceReturnFacetAction `json:"forceReturnFacetAction,omitempty"`

	// Remove an attribute as a facet in the request (if present).
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.remove_facet_action
	RemoveFacetAction *Rule_RemoveFacetAction `json:"removeFacetAction,omitempty"`

	// Required. The condition that triggers the rule.
	//  If the condition is empty, the rule will always apply.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.condition
	Condition *Condition `json:"condition,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Rule.BoostAction
type Rule_BoostAction struct {
	// Strength of the condition boost, which must be in [-1, 1]. Negative
	//  boost means demotion. Default is 0.0.
	//
	//  Setting to 1.0 gives the item a big promotion. However, it does not
	//  necessarily mean that the boosted item will be the top result at all
	//  times, nor that other items will be excluded. Results could still be
	//  shown even when none of them matches the condition. And results that
	//  are significantly more relevant to the search query can still trump
	//  your heavily favored but irrelevant items.
	//
	//  Setting to -1.0 gives the item a big demotion. However, results that
	//  are deeply relevant might still be shown. The item will have an
	//  upstream battle to get a fairly high ranking, but it is not blocked out
	//  completely.
	//
	//  Setting to 0.0 means no boost applied. The boosting condition is
	//  ignored.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.BoostAction.boost
	Boost *float32 `json:"boost,omitempty"`

	// The filter can have a max size of 5000 characters.
	//  An expression which specifies which products to apply an action to.
	//  The syntax and supported fields are the same as a filter expression. See
	//  [SearchRequest.filter][google.cloud.retail.v2beta.SearchRequest.filter]
	//  for detail syntax and limitations.
	//
	//  Examples:
	//
	//  * To boost products with product ID "product_1" or "product_2", and
	//  color
	//    "Red" or "Blue":<br>
	//    *(id: ANY("product_1", "product_2"))<br>*
	//    *AND<br>*
	//    *(colorFamilies: ANY("Red", "Blue"))<br>*
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.BoostAction.products_filter
	ProductsFilter *string `json:"productsFilter,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Rule.DoNotAssociateAction
type Rule_DoNotAssociateAction struct {
	// Terms from the search query.
	//  Will not consider do_not_associate_terms for search if in search query.
	//  Can specify up to 100 terms.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.DoNotAssociateAction.query_terms
	QueryTerms []string `json:"queryTerms,omitempty"`

	// Cannot contain duplicates or the query term.
	//  Can specify up to 100 terms.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.DoNotAssociateAction.do_not_associate_terms
	DoNotAssociateTerms []string `json:"doNotAssociateTerms,omitempty"`

	// Will be [deprecated = true] post migration;
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.DoNotAssociateAction.terms
	Terms []string `json:"terms,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Rule.FilterAction
type Rule_FilterAction struct {
	// A filter to apply on the matching condition results. Supported features:
	//
	//  * [filter][google.cloud.retail.v2beta.Rule.FilterAction.filter] must be
	//  set.
	//  * Filter syntax is identical to
	//  [SearchRequest.filter][google.cloud.retail.v2beta.SearchRequest.filter].
	//  For more
	//    information, see [Filter](/retail/docs/filter-and-order#filter).
	//  * To filter products with product ID "product_1" or "product_2", and
	//  color
	//    "Red" or "Blue":<br>
	//    *(id: ANY("product_1", "product_2"))<br>*
	//    *AND<br>*
	//    *(colorFamilies: ANY("Red", "Blue"))<br>*
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.FilterAction.filter
	Filter *string `json:"filter,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Rule.ForceReturnFacetAction
type Rule_ForceReturnFacetAction struct {
	// Each instance corresponds to a force return attribute for the given
	//  condition. There can't be more 15 instances here.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.ForceReturnFacetAction.facet_position_adjustments
	FacetPositionAdjustments []Rule_ForceReturnFacetAction_FacetPositionAdjustment `json:"facetPositionAdjustments,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Rule.ForceReturnFacetAction.FacetPositionAdjustment
type Rule_ForceReturnFacetAction_FacetPositionAdjustment struct {
	// The attribute name to force return as a facet. Each attribute name
	//  should be a valid attribute name, be non-empty and contain at most 80
	//  characters long.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.ForceReturnFacetAction.FacetPositionAdjustment.attribute_name
	AttributeName *string `json:"attributeName,omitempty"`

	// This is the position in the request as explained above. It should be
	//  strictly positive be at most 100.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.ForceReturnFacetAction.FacetPositionAdjustment.position
	Position *int32 `json:"position,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Rule.IgnoreAction
type Rule_IgnoreAction struct {
	// Terms to ignore in the search query.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.IgnoreAction.ignore_terms
	IgnoreTerms []string `json:"ignoreTerms,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Rule.OnewaySynonymsAction
type Rule_OnewaySynonymsAction struct {
	// Terms from the search query.
	//  Will treat synonyms as their synonyms.
	//  Not themselves synonyms of the synonyms.
	//  Can specify up to 100 terms.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.OnewaySynonymsAction.query_terms
	QueryTerms []string `json:"queryTerms,omitempty"`

	// Defines a set of synonyms.
	//  Cannot contain duplicates.
	//  Can specify up to 100 synonyms.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.OnewaySynonymsAction.synonyms
	Synonyms []string `json:"synonyms,omitempty"`

	// Will be [deprecated = true] post migration;
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.OnewaySynonymsAction.oneway_terms
	OnewayTerms []string `json:"onewayTerms,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Rule.RedirectAction
type Rule_RedirectAction struct {
	// URL must have length equal or less than 2000 characters.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.RedirectAction.redirect_uri
	RedirectURI *string `json:"redirectURI,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Rule.RemoveFacetAction
type Rule_RemoveFacetAction struct {
	// The attribute names (i.e. facet keys) to remove from the dynamic facets
	//  (if present in the request). There can't be more 3 attribute names.
	//  Each attribute name should be a valid attribute name, be non-empty and
	//  contain at most 80 characters.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.RemoveFacetAction.attribute_names
	AttributeNames []string `json:"attributeNames,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Rule.ReplacementAction
type Rule_ReplacementAction struct {
	// Terms from the search query.
	//  Will be replaced by replacement term.
	//  Can specify up to 100 terms.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.ReplacementAction.query_terms
	QueryTerms []string `json:"queryTerms,omitempty"`

	// Term that will be used for replacement.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.ReplacementAction.replacement_term
	ReplacementTerm *string `json:"replacementTerm,omitempty"`

	// Will be [deprecated = true] post migration;
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.ReplacementAction.term
	Term *string `json:"term,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Rule.TwowaySynonymsAction
type Rule_TwowaySynonymsAction struct {
	// Defines a set of synonyms.
	//  Can specify up to 100 synonyms.
	//  Must specify at least 2 synonyms.
	// +kcc:proto:field=google.cloud.retail.v2beta.Rule.TwowaySynonymsAction.synonyms
	Synonyms []string `json:"synonyms,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.SearchRequest.FacetSpec
type SearchRequest_FacetSpec struct {
	// Required. The facet key specification.
	// +kcc:proto:field=google.cloud.retail.v2beta.SearchRequest.FacetSpec.facet_key
	FacetKey *SearchRequest_FacetSpec_FacetKey `json:"facetKey,omitempty"`

	// Maximum of facet values that should be returned for this facet. If
	//  unspecified, defaults to 50. The maximum allowed value is 300. Values
	//  above 300 will be coerced to 300.
	//
	//  If this field is negative, an INVALID_ARGUMENT is returned.
	// +kcc:proto:field=google.cloud.retail.v2beta.SearchRequest.FacetSpec.limit
	Limit *int32 `json:"limit,omitempty"`

	// List of keys to exclude when faceting.
	//
	//
	//  By default,
	//  [FacetKey.key][google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.key]
	//  is not excluded from the filter unless it is listed in this field.
	//
	//  Listing a facet key in this field allows its values to appear as facet
	//  results, even when they are filtered out of search results. Using this
	//  field does not affect what search results are returned.
	//
	//  For example, suppose there are 100 products with the color facet "Red"
	//  and 200 products with the color facet "Blue". A query containing the
	//  filter "colorFamilies:ANY("Red")" and having "colorFamilies" as
	//  [FacetKey.key][google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.key]
	//  would by default return only "Red" products in the search results, and
	//  also return "Red" with count 100 as the only color facet. Although there
	//  are also blue products available, "Blue" would not be shown as an
	//  available facet value.
	//
	//  If "colorFamilies" is listed in "excludedFilterKeys", then the query
	//  returns the facet values "Red" with count 100 and "Blue" with count
	//  200, because the "colorFamilies" key is now excluded from the filter.
	//  Because this field doesn't affect search results, the search results
	//  are still correctly filtered to return only "Red" products.
	//
	//  A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error
	//  is returned.
	// +kcc:proto:field=google.cloud.retail.v2beta.SearchRequest.FacetSpec.excluded_filter_keys
	ExcludedFilterKeys []string `json:"excludedFilterKeys,omitempty"`

	// Enables dynamic position for this facet. If set to true, the position of
	//  this facet among all facets in the response is determined by Google
	//  Retail Search. It is ordered together with dynamic facets if dynamic
	//  facets is enabled. If set to false, the position of this facet in the
	//  response is the same as in the request, and it is ranked before
	//  the facets with dynamic position enable and all dynamic facets.
	//
	//  For example, you may always want to have rating facet returned in
	//  the response, but it's not necessarily to always display the rating facet
	//  at the top. In that case, you can set enable_dynamic_position to true so
	//  that the position of rating facet in response is determined by
	//  Google Retail Search.
	//
	//  Another example, assuming you have the following facets in the request:
	//
	//  * "rating", enable_dynamic_position = true
	//
	//  * "price", enable_dynamic_position = false
	//
	//  * "brands", enable_dynamic_position = false
	//
	//  And also you have a dynamic facets enable, which generates a facet
	//  "gender". Then, the final order of the facets in the response can be
	//  ("price", "brands", "rating", "gender") or ("price", "brands", "gender",
	//  "rating") depends on how Google Retail Search orders "gender" and
	//  "rating" facets. However, notice that "price" and "brands" are always
	//  ranked at first and second position because their enable_dynamic_position
	//  values are false.
	// +kcc:proto:field=google.cloud.retail.v2beta.SearchRequest.FacetSpec.enable_dynamic_position
	EnableDynamicPosition *bool `json:"enableDynamicPosition,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey
type SearchRequest_FacetSpec_FacetKey struct {
	// Required. Supported textual and numerical facet keys in
	//  [Product][google.cloud.retail.v2beta.Product] object, over which the
	//  facet values are computed. Facet key is case-sensitive.
	//
	//  Allowed facet keys when
	//  [FacetKey.query][google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.query]
	//  is not specified:
	//
	//  * textual_field =
	//      * "brands"
	//      * "categories"
	//      * "genders"
	//      * "ageGroups"
	//      * "availability"
	//      * "colorFamilies"
	//      * "colors"
	//      * "sizes"
	//      * "materials"
	//      * "patterns"
	//      * "conditions"
	//      * "attributes.key"
	//      * "pickupInStore"
	//      * "shipToStore"
	//      * "sameDayDelivery"
	//      * "nextDayDelivery"
	//      * "customFulfillment1"
	//      * "customFulfillment2"
	//      * "customFulfillment3"
	//      * "customFulfillment4"
	//      * "customFulfillment5"
	//      * "inventory(place_id,attributes.key)"
	//
	//  * numerical_field =
	//      * "price"
	//      * "discount"
	//      * "rating"
	//      * "ratingCount"
	//      * "attributes.key"
	//      * "inventory(place_id,price)"
	//      * "inventory(place_id,original_price)"
	//      * "inventory(place_id,attributes.key)"
	// +kcc:proto:field=google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.key
	Key *string `json:"key,omitempty"`

	// Set only if values should be bucketized into intervals. Must be set
	//  for facets with numerical values. Must not be set for facet with text
	//  values. Maximum number of intervals is 40.
	//
	//  For all numerical facet keys that appear in the list of products from
	//  the catalog, the percentiles 0, 10, 30, 50, 70, 90, and 100 are
	//  computed from their distribution weekly. If the model assigns a high
	//  score to a numerical facet key and its intervals are not specified in
	//  the search request, these percentiles become the bounds
	//  for its intervals and are returned in the response. If the
	//  facet key intervals are specified in the request, then the specified
	//  intervals are returned instead.
	// +kcc:proto:field=google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.intervals
	Intervals []Interval `json:"intervals,omitempty"`

	// Only get facet for the given restricted values. For example, when using
	//  "pickupInStore" as key and set restricted values to
	//  ["store123", "store456"], only facets for "store123" and "store456" are
	//  returned. Only supported on predefined textual fields, custom textual
	//  attributes and fulfillments. Maximum is 20.
	//
	//  Must be set for the fulfillment facet keys:
	//
	//  * pickupInStore
	//
	//  * shipToStore
	//
	//  * sameDayDelivery
	//
	//  * nextDayDelivery
	//
	//  * customFulfillment1
	//
	//  * customFulfillment2
	//
	//  * customFulfillment3
	//
	//  * customFulfillment4
	//
	//  * customFulfillment5
	// +kcc:proto:field=google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.restricted_values
	RestrictedValues []string `json:"restrictedValues,omitempty"`

	// Only get facet values that start with the given string prefix. For
	//  example, suppose "categories" has three values "Women > Shoe",
	//  "Women > Dress" and "Men > Shoe". If set "prefixes" to "Women", the
	//  "categories" facet gives only "Women > Shoe" and "Women > Dress".
	//  Only supported on textual fields. Maximum is 10.
	// +kcc:proto:field=google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.prefixes
	Prefixes []string `json:"prefixes,omitempty"`

	// Only get facet values that contains the given strings. For example,
	//  suppose "categories" has three values "Women > Shoe",
	//  "Women > Dress" and "Men > Shoe". If set "contains" to "Shoe", the
	//  "categories" facet gives only "Women > Shoe" and "Men > Shoe".
	//  Only supported on textual fields. Maximum is 10.
	// +kcc:proto:field=google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.contains
	Contains []string `json:"contains,omitempty"`

	// True to make facet keys case insensitive when getting faceting
	//  values with prefixes or contains; false otherwise.
	// +kcc:proto:field=google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.case_insensitive
	CaseInsensitive *bool `json:"caseInsensitive,omitempty"`

	// The order in which
	//  [SearchResponse.Facet.values][google.cloud.retail.v2beta.SearchResponse.Facet.values]
	//  are returned.
	//
	//  Allowed values are:
	//
	//  * "count desc", which means order by
	//  [SearchResponse.Facet.values.count][google.cloud.retail.v2beta.SearchResponse.Facet.FacetValue.count]
	//  descending.
	//
	//  * "value desc", which means order by
	//  [SearchResponse.Facet.values.value][google.cloud.retail.v2beta.SearchResponse.Facet.FacetValue.value]
	//  descending.
	//    Only applies to textual facets.
	//
	//  If not set, textual values are sorted in [natural
	//  order](https://en.wikipedia.org/wiki/Natural_sort_order); numerical
	//  intervals are sorted in the order given by
	//  [FacetSpec.FacetKey.intervals][google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.intervals];
	//  [FulfillmentInfo.place_ids][google.cloud.retail.v2beta.FulfillmentInfo.place_ids]
	//  are sorted in the order given by
	//  [FacetSpec.FacetKey.restricted_values][google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.restricted_values].
	// +kcc:proto:field=google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.order_by
	OrderBy *string `json:"orderBy,omitempty"`

	// The query that is used to compute facet for the given facet key.
	//  When provided, it overrides the default behavior of facet
	//  computation. The query syntax is the same as a filter expression. See
	//  [SearchRequest.filter][google.cloud.retail.v2beta.SearchRequest.filter]
	//  for detail syntax and limitations. Notice that there is no limitation
	//  on
	//  [FacetKey.key][google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.key]
	//  when query is specified.
	//
	//  In the response,
	//  [SearchResponse.Facet.values.value][google.cloud.retail.v2beta.SearchResponse.Facet.FacetValue.value]
	//  is always "1" and
	//  [SearchResponse.Facet.values.count][google.cloud.retail.v2beta.SearchResponse.Facet.FacetValue.count]
	//  is the number of results that match the query.
	//
	//  For example, you can set a customized facet for "shipToStore",
	//  where
	//  [FacetKey.key][google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.key]
	//  is "customizedShipToStore", and
	//  [FacetKey.query][google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.query]
	//  is "availability: ANY(\"IN_STOCK\") AND shipToStore: ANY(\"123\")".
	//  Then the facet counts the products that are both in stock and ship
	//  to store "123".
	// +kcc:proto:field=google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.query
	Query *string `json:"query,omitempty"`

	// Returns the min and max value for each numerical facet intervals.
	//  Ignored for textual facets.
	// +kcc:proto:field=google.cloud.retail.v2beta.SearchRequest.FacetSpec.FacetKey.return_min_max
	ReturnMinMax *bool `json:"returnMinMax,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2beta.Control
type ControlObservedState struct {
	// Output only. List of [serving
	//  config][google.cloud.retail.v2beta.ServingConfig] ids that are associated
	//  with this control in the same
	//  [Catalog][google.cloud.retail.v2beta.Catalog].
	//
	//  Note the association is managed via the
	//  [ServingConfig][google.cloud.retail.v2beta.ServingConfig], this is an
	//  output only denormalized view.
	// +kcc:proto:field=google.cloud.retail.v2beta.Control.associated_serving_config_ids
	AssociatedServingConfigIds []string `json:"associatedServingConfigIds,omitempty"`
}
