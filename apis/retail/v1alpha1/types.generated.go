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


// +kcc:proto=google.cloud.retail.v2.Condition
type Condition struct {
	// A list (up to 10 entries) of terms to match the query on. If not
	//  specified, match all queries.
	//  If many query terms are specified, the condition
	//  is matched if any of the terms is a match (i.e. using the OR operator).
	// +kcc:proto:field=google.cloud.retail.v2.Condition.query_terms
	QueryTerms []Condition_QueryTerm `json:"queryTerms,omitempty"`

	// Range of time(s) specifying when Condition is active.
	//  Condition true if any time range matches.
	// +kcc:proto:field=google.cloud.retail.v2.Condition.active_time_range
	ActiveTimeRange []Condition_TimeRange `json:"activeTimeRange,omitempty"`

	// Used to support browse uses cases.
	//  A list (up to 10 entries) of categories or departments.
	//  The format should be the same as
	//  [UserEvent.page_categories][google.cloud.retail.v2.UserEvent.page_categories];
	// +kcc:proto:field=google.cloud.retail.v2.Condition.page_categories
	PageCategories []string `json:"pageCategories,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Condition.QueryTerm
type Condition_QueryTerm struct {
	// The value of the term to match on.
	//  Value cannot be empty.
	//  Value can have at most 3 terms if specified as a partial match. Each
	//  space separated string is considered as one term.
	//  For example, "a b c" is 3 terms and allowed, but " a b c d" is 4 terms
	//  and not allowed for a partial match.
	// +kcc:proto:field=google.cloud.retail.v2.Condition.QueryTerm.value
	Value *string `json:"value,omitempty"`

	// Whether this is supposed to be a full or partial match.
	// +kcc:proto:field=google.cloud.retail.v2.Condition.QueryTerm.full_match
	FullMatch *bool `json:"fullMatch,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Condition.TimeRange
type Condition_TimeRange struct {
	// Start of time range. Range is inclusive.
	// +kcc:proto:field=google.cloud.retail.v2.Condition.TimeRange.start_time
	StartTime *string `json:"startTime,omitempty"`

	// End of time range. Range is inclusive.
	// +kcc:proto:field=google.cloud.retail.v2.Condition.TimeRange.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Control
type Control struct {
	// A rule control - a condition-action pair.
	//  Enacts a set action when the condition is triggered.
	//  For example: Boost "gShoe" when query full matches "Running Shoes".
	// +kcc:proto:field=google.cloud.retail.v2.Control.rule
	Rule *Rule `json:"rule,omitempty"`

	// Immutable. Fully qualified name
	//  `projects/*/locations/global/catalogs/*/controls/*`
	// +kcc:proto:field=google.cloud.retail.v2.Control.name
	Name *string `json:"name,omitempty"`

	// Required. The human readable control display name. Used in Retail UI.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 128
	//  characters. Otherwise, an INVALID_ARGUMENT error is thrown.
	// +kcc:proto:field=google.cloud.retail.v2.Control.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. The solution types that the control is used for.
	//  Currently we support setting only one type of solution at creation time.
	//
	//  Only `SOLUTION_TYPE_SEARCH` value is supported at the moment.
	//  If no solution type is provided at creation time, will default to
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.retail.v2.Control.solution_types
	SolutionTypes []string `json:"solutionTypes,omitempty"`

	// Specifies the use case for the control.
	//  Affects what condition fields can be set.
	//  Only settable by search controls.
	//  Will default to
	//  [SEARCH_SOLUTION_USE_CASE_SEARCH][google.cloud.retail.v2.SearchSolutionUseCase.SEARCH_SOLUTION_USE_CASE_SEARCH]
	//  if not specified. Currently only allow one search_solution_use_case per
	//  control.
	// +kcc:proto:field=google.cloud.retail.v2.Control.search_solution_use_case
	SearchSolutionUseCase []string `json:"searchSolutionUseCase,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Rule
type Rule struct {
	// A boost action.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.boost_action
	BoostAction *Rule_BoostAction `json:"boostAction,omitempty"`

	// Redirects a shopper to a specific page.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.redirect_action
	RedirectAction *Rule_RedirectAction `json:"redirectAction,omitempty"`

	// Treats specific term as a synonym with a group of terms.
	//  Group of terms will not be treated as synonyms with the specific term.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.oneway_synonyms_action
	OnewaySynonymsAction *Rule_OnewaySynonymsAction `json:"onewaySynonymsAction,omitempty"`

	// Prevents term from being associated with other terms.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.do_not_associate_action
	DoNotAssociateAction *Rule_DoNotAssociateAction `json:"doNotAssociateAction,omitempty"`

	// Replaces specific terms in the query.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.replacement_action
	ReplacementAction *Rule_ReplacementAction `json:"replacementAction,omitempty"`

	// Ignores specific terms from query during search.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.ignore_action
	IgnoreAction *Rule_IgnoreAction `json:"ignoreAction,omitempty"`

	// Filters results.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.filter_action
	FilterAction *Rule_FilterAction `json:"filterAction,omitempty"`

	// Treats a set of terms as synonyms of one another.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.twoway_synonyms_action
	TwowaySynonymsAction *Rule_TwowaySynonymsAction `json:"twowaySynonymsAction,omitempty"`

	// Force returns an attribute as a facet in the request.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.force_return_facet_action
	ForceReturnFacetAction *Rule_ForceReturnFacetAction `json:"forceReturnFacetAction,omitempty"`

	// Remove an attribute as a facet in the request (if present).
	// +kcc:proto:field=google.cloud.retail.v2.Rule.remove_facet_action
	RemoveFacetAction *Rule_RemoveFacetAction `json:"removeFacetAction,omitempty"`

	// Required. The condition that triggers the rule.
	//  If the condition is empty, the rule will always apply.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.condition
	Condition *Condition `json:"condition,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Rule.BoostAction
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
	// +kcc:proto:field=google.cloud.retail.v2.Rule.BoostAction.boost
	Boost *float32 `json:"boost,omitempty"`

	// The filter can have a max size of 5000 characters.
	//  An expression which specifies which products to apply an action to.
	//  The syntax and supported fields are the same as a filter expression. See
	//  [SearchRequest.filter][google.cloud.retail.v2.SearchRequest.filter] for
	//  detail syntax and limitations.
	//
	//  Examples:
	//
	//  * To boost products with product ID "product_1" or "product_2", and
	//  color
	//    "Red" or "Blue":<br>
	//    *(id: ANY("product_1", "product_2"))<br>*
	//    *AND<br>*
	//    *(colorFamilies: ANY("Red", "Blue"))<br>*
	// +kcc:proto:field=google.cloud.retail.v2.Rule.BoostAction.products_filter
	ProductsFilter *string `json:"productsFilter,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Rule.DoNotAssociateAction
type Rule_DoNotAssociateAction struct {
	// Terms from the search query.
	//  Will not consider do_not_associate_terms for search if in search query.
	//  Can specify up to 100 terms.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.DoNotAssociateAction.query_terms
	QueryTerms []string `json:"queryTerms,omitempty"`

	// Cannot contain duplicates or the query term.
	//  Can specify up to 100 terms.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.DoNotAssociateAction.do_not_associate_terms
	DoNotAssociateTerms []string `json:"doNotAssociateTerms,omitempty"`

	// Will be [deprecated = true] post migration;
	// +kcc:proto:field=google.cloud.retail.v2.Rule.DoNotAssociateAction.terms
	Terms []string `json:"terms,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Rule.FilterAction
type Rule_FilterAction struct {
	// A filter to apply on the matching condition results. Supported features:
	//
	//  * [filter][google.cloud.retail.v2.Rule.FilterAction.filter] must be set.
	//  * Filter syntax is identical to
	//  [SearchRequest.filter][google.cloud.retail.v2.SearchRequest.filter]. For
	//  more
	//    information, see [Filter](/retail/docs/filter-and-order#filter).
	//  * To filter products with product ID "product_1" or "product_2", and
	//  color
	//    "Red" or "Blue":<br>
	//    *(id: ANY("product_1", "product_2"))<br>*
	//    *AND<br>*
	//    *(colorFamilies: ANY("Red", "Blue"))<br>*
	// +kcc:proto:field=google.cloud.retail.v2.Rule.FilterAction.filter
	Filter *string `json:"filter,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Rule.ForceReturnFacetAction
type Rule_ForceReturnFacetAction struct {
	// Each instance corresponds to a force return attribute for the given
	//  condition. There can't be more 15 instances here.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.ForceReturnFacetAction.facet_position_adjustments
	FacetPositionAdjustments []Rule_ForceReturnFacetAction_FacetPositionAdjustment `json:"facetPositionAdjustments,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Rule.ForceReturnFacetAction.FacetPositionAdjustment
type Rule_ForceReturnFacetAction_FacetPositionAdjustment struct {
	// The attribute name to force return as a facet. Each attribute name
	//  should be a valid attribute name, be non-empty and contain at most 80
	//  characters long.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.ForceReturnFacetAction.FacetPositionAdjustment.attribute_name
	AttributeName *string `json:"attributeName,omitempty"`

	// This is the position in the request as explained above. It should be
	//  strictly positive be at most 100.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.ForceReturnFacetAction.FacetPositionAdjustment.position
	Position *int32 `json:"position,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Rule.IgnoreAction
type Rule_IgnoreAction struct {
	// Terms to ignore in the search query.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.IgnoreAction.ignore_terms
	IgnoreTerms []string `json:"ignoreTerms,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Rule.OnewaySynonymsAction
type Rule_OnewaySynonymsAction struct {
	// Terms from the search query.
	//  Will treat synonyms as their synonyms.
	//  Not themselves synonyms of the synonyms.
	//  Can specify up to 100 terms.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.OnewaySynonymsAction.query_terms
	QueryTerms []string `json:"queryTerms,omitempty"`

	// Defines a set of synonyms.
	//  Cannot contain duplicates.
	//  Can specify up to 100 synonyms.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.OnewaySynonymsAction.synonyms
	Synonyms []string `json:"synonyms,omitempty"`

	// Will be [deprecated = true] post migration;
	// +kcc:proto:field=google.cloud.retail.v2.Rule.OnewaySynonymsAction.oneway_terms
	OnewayTerms []string `json:"onewayTerms,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Rule.RedirectAction
type Rule_RedirectAction struct {
	// URL must have length equal or less than 2000 characters.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.RedirectAction.redirect_uri
	RedirectURI *string `json:"redirectURI,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Rule.RemoveFacetAction
type Rule_RemoveFacetAction struct {
	// The attribute names (i.e. facet keys) to remove from the dynamic facets
	//  (if present in the request). There can't be more 3 attribute names.
	//  Each attribute name should be a valid attribute name, be non-empty and
	//  contain at most 80 characters.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.RemoveFacetAction.attribute_names
	AttributeNames []string `json:"attributeNames,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Rule.ReplacementAction
type Rule_ReplacementAction struct {
	// Terms from the search query.
	//  Will be replaced by replacement term.
	//  Can specify up to 100 terms.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.ReplacementAction.query_terms
	QueryTerms []string `json:"queryTerms,omitempty"`

	// Term that will be used for replacement.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.ReplacementAction.replacement_term
	ReplacementTerm *string `json:"replacementTerm,omitempty"`

	// Will be [deprecated = true] post migration;
	// +kcc:proto:field=google.cloud.retail.v2.Rule.ReplacementAction.term
	Term *string `json:"term,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Rule.TwowaySynonymsAction
type Rule_TwowaySynonymsAction struct {
	// Defines a set of synonyms.
	//  Can specify up to 100 synonyms.
	//  Must specify at least 2 synonyms.
	// +kcc:proto:field=google.cloud.retail.v2.Rule.TwowaySynonymsAction.synonyms
	Synonyms []string `json:"synonyms,omitempty"`
}

// +kcc:proto=google.cloud.retail.v2.Control
type ControlObservedState struct {
	// Output only. List of [serving config][google.cloud.retail.v2.ServingConfig]
	//  ids that are associated with this control in the same
	//  [Catalog][google.cloud.retail.v2.Catalog].
	//
	//  Note the association is managed via the
	//  [ServingConfig][google.cloud.retail.v2.ServingConfig], this is an output
	//  only denormalized view.
	// +kcc:proto:field=google.cloud.retail.v2.Control.associated_serving_config_ids
	AssociatedServingConfigIds []string `json:"associatedServingConfigIds,omitempty"`
}
