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


// +kcc:proto=google.cloud.discoveryengine.v1beta.Condition
type Condition struct {
	// Search only
	//  A list of terms to match the query on.
	//  Cannot be set when
	//  [Condition.query_regex][google.cloud.discoveryengine.v1beta.Condition.query_regex]
	//  is set.
	//
	//  Maximum of 10 query terms.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Condition.query_terms
	QueryTerms []Condition_QueryTerm `json:"queryTerms,omitempty"`

	// Range of time(s) specifying when condition is active.
	//
	//  Maximum of 10 time ranges.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Condition.active_time_range
	ActiveTimeRange []Condition_TimeRange `json:"activeTimeRange,omitempty"`

	// Optional. Query regex to match the whole search query.
	//  Cannot be set when
	//  [Condition.query_terms][google.cloud.discoveryengine.v1beta.Condition.query_terms]
	//  is set. This is currently supporting promotion use case.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Condition.query_regex
	QueryRegex *string `json:"queryRegex,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Condition.QueryTerm
type Condition_QueryTerm struct {
	// The specific query value to match against
	//
	//  Must be lowercase, must be UTF-8.
	//  Can have at most 3 space separated terms if full_match is true.
	//  Cannot be an empty string.
	//  Maximum length of 5000 characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Condition.QueryTerm.value
	Value *string `json:"value,omitempty"`

	// Whether the search query needs to exactly match the query term.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Condition.QueryTerm.full_match
	FullMatch *bool `json:"fullMatch,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Condition.TimeRange
type Condition_TimeRange struct {
	// Start of time range.
	//
	//  Range is inclusive.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Condition.TimeRange.start_time
	StartTime *string `json:"startTime,omitempty"`

	// End of time range.
	//
	//  Range is inclusive.
	//  Must be in the future.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Condition.TimeRange.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Control
type Control struct {
	// Defines a boost-type control
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.boost_action
	BoostAction *Control_BoostAction `json:"boostAction,omitempty"`

	// Defines a filter-type control
	//  Currently not supported by Recommendation
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.filter_action
	FilterAction *Control_FilterAction `json:"filterAction,omitempty"`

	// Defines a redirect-type control.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.redirect_action
	RedirectAction *Control_RedirectAction `json:"redirectAction,omitempty"`

	// Treats a group of terms as synonyms of one another.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.synonyms_action
	SynonymsAction *Control_SynonymsAction `json:"synonymsAction,omitempty"`

	// Immutable. Fully qualified name
	//  `projects/*/locations/global/dataStore/*/controls/*`
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.name
	Name *string `json:"name,omitempty"`

	// Required. Human readable name. The identifier used in UI views.
	//
	//  Must be UTF-8 encoded string. Length limit is 128 characters.
	//  Otherwise an INVALID ARGUMENT error is thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. What solution the control belongs to.
	//
	//  Must be compatible with vertical of resource.
	//  Otherwise an INVALID ARGUMENT error is thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.solution_type
	SolutionType *string `json:"solutionType,omitempty"`

	// Specifies the use case for the control.
	//  Affects what condition fields can be set.
	//  Only applies to
	//  [SOLUTION_TYPE_SEARCH][google.cloud.discoveryengine.v1beta.SolutionType.SOLUTION_TYPE_SEARCH].
	//  Currently only allow one use case per control.
	//  Must be set when solution_type is
	//  [SolutionType.SOLUTION_TYPE_SEARCH][google.cloud.discoveryengine.v1beta.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.use_cases
	UseCases []string `json:"useCases,omitempty"`

	// Determines when the associated action will trigger.
	//
	//  Omit to always apply the action.
	//  Currently only a single condition may be specified.
	//  Otherwise an INVALID ARGUMENT error is thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.conditions
	Conditions []Condition `json:"conditions,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Control.BoostAction
type Control_BoostAction struct {
	// Required. Strength of the boost, which should be in [-1, 1]. Negative
	//  boost means demotion. Default is 0.0 (No-op).
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.BoostAction.boost
	Boost *float32 `json:"boost,omitempty"`

	// Required. Specifies which products to apply the boost to.
	//
	//  If no filter is provided all products will be boosted (No-op).
	//  Syntax documentation:
	//  https://cloud.google.com/retail/docs/filter-and-order
	//  Maximum length is 5000 characters.
	//  Otherwise an INVALID ARGUMENT error is thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.BoostAction.filter
	Filter *string `json:"filter,omitempty"`

	// Required. Specifies which data store's documents can be boosted by this
	//  control. Full data store name e.g.
	//  projects/123/locations/global/collections/default_collection/dataStores/default_data_store
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.BoostAction.data_store
	DataStore *string `json:"dataStore,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Control.FilterAction
type Control_FilterAction struct {
	// Required. A filter to apply on the matching condition results.
	//
	//  Required
	//  Syntax documentation:
	//  https://cloud.google.com/retail/docs/filter-and-order
	//  Maximum length is 5000 characters. Otherwise an INVALID
	//  ARGUMENT error is thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.FilterAction.filter
	Filter *string `json:"filter,omitempty"`

	// Required. Specifies which data store's documents can be filtered by this
	//  control. Full data store name e.g.
	//  projects/123/locations/global/collections/default_collection/dataStores/default_data_store
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.FilterAction.data_store
	DataStore *string `json:"dataStore,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Control.RedirectAction
type Control_RedirectAction struct {
	// Required. The URI to which the shopper will be redirected.
	//
	//  Required.
	//  URI must have length equal or less than 2000 characters.
	//  Otherwise an INVALID ARGUMENT error is thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.RedirectAction.redirect_uri
	RedirectURI *string `json:"redirectURI,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Control.SynonymsAction
type Control_SynonymsAction struct {
	// Defines a set of synonyms.
	//  Can specify up to 100 synonyms.
	//  Must specify at least 2 synonyms. Otherwise an INVALID ARGUMENT error is
	//  thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.SynonymsAction.synonyms
	Synonyms []string `json:"synonyms,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Control
type ControlObservedState struct {
	// Output only. List of all
	//  [ServingConfig][google.cloud.discoveryengine.v1beta.ServingConfig] IDs this
	//  control is attached to. May take up to 10 minutes to update after changes.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Control.associated_serving_config_ids
	AssociatedServingConfigIds []string `json:"associatedServingConfigIds,omitempty"`
}
