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


// +kcc:proto=google.cloud.discoveryengine.v1beta.Evaluation
type Evaluation struct {
	// Identifier. The full resource name of the
	//  [Evaluation][google.cloud.discoveryengine.v1beta.Evaluation], in the format
	//  of `projects/{project}/locations/{location}/evaluations/{evaluation}`.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 1024
	//  characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Evaluation.name
	Name *string `json:"name,omitempty"`

	// Required. The specification of the evaluation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Evaluation.evaluation_spec
	EvaluationSpec *Evaluation_EvaluationSpec `json:"evaluationSpec,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Evaluation.EvaluationSpec
type Evaluation_EvaluationSpec struct {
	// Required. The search request that is used to perform the evaluation.
	//
	//  Only the following fields within SearchRequest are supported; if any
	//  other fields are provided, an UNSUPPORTED error will be returned:
	//
	//  * [SearchRequest.serving_config][google.cloud.discoveryengine.v1beta.SearchRequest.serving_config]
	//  * [SearchRequest.branch][google.cloud.discoveryengine.v1beta.SearchRequest.branch]
	//  * [SearchRequest.canonical_filter][google.cloud.discoveryengine.v1beta.SearchRequest.canonical_filter]
	//  * [SearchRequest.query_expansion_spec][google.cloud.discoveryengine.v1beta.SearchRequest.query_expansion_spec]
	//  * [SearchRequest.spell_correction_spec][google.cloud.discoveryengine.v1beta.SearchRequest.spell_correction_spec]
	//  * [SearchRequest.content_search_spec][google.cloud.discoveryengine.v1beta.SearchRequest.content_search_spec]
	//  * [SearchRequest.user_pseudo_id][google.cloud.discoveryengine.v1beta.SearchRequest.user_pseudo_id]
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Evaluation.EvaluationSpec.search_request
	SearchRequest *SearchRequest `json:"searchRequest,omitempty"`

	// Required. The specification of the query set.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Evaluation.EvaluationSpec.query_set_spec
	QuerySetSpec *Evaluation_EvaluationSpec_QuerySetSpec `json:"querySetSpec,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Evaluation.EvaluationSpec.QuerySetSpec
type Evaluation_EvaluationSpec_QuerySetSpec struct {
	// Required. The full resource name of the
	//  [SampleQuerySet][google.cloud.discoveryengine.v1beta.SampleQuerySet]
	//  used for the evaluation, in the format of
	//  `projects/{project}/locations/{location}/sampleQuerySets/{sampleQuerySet}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Evaluation.EvaluationSpec.QuerySetSpec.sample_query_set
	SampleQuerySet *string `json:"sampleQuerySet,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Interval
type Interval struct {
	// Inclusive lower bound.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Interval.minimum
	Minimum *float64 `json:"minimum,omitempty"`

	// Exclusive lower bound.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Interval.exclusive_minimum
	ExclusiveMinimum *float64 `json:"exclusiveMinimum,omitempty"`

	// Inclusive upper bound.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Interval.maximum
	Maximum *float64 `json:"maximum,omitempty"`

	// Exclusive upper bound.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Interval.exclusive_maximum
	ExclusiveMaximum *float64 `json:"exclusiveMaximum,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.QualityMetrics
type QualityMetrics struct {
	// Recall per document, at various top-k cutoff levels.
	//
	//  Recall is the fraction of relevant documents retrieved out of all
	//  relevant documents.
	//
	//  Example (top-5):
	//   * For a single
	//   [SampleQuery][google.cloud.discoveryengine.v1beta.SampleQuery], If 3 out
	//   of 5 relevant documents are retrieved in the top-5, recall@5 = 3/5 = 0.6
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.QualityMetrics.doc_recall
	DocRecall *QualityMetrics_TopkMetrics `json:"docRecall,omitempty"`

	// Precision per document, at various top-k cutoff levels.
	//
	//  Precision is the fraction of retrieved documents that are relevant.
	//
	//  Example (top-5):
	//   * For a single
	//   [SampleQuery][google.cloud.discoveryengine.v1beta.SampleQuery], If 4 out
	//   of 5 retrieved documents in the top-5 are relevant, precision@5 = 4/5 =
	//   0.8
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.QualityMetrics.doc_precision
	DocPrecision *QualityMetrics_TopkMetrics `json:"docPrecision,omitempty"`

	// Normalized discounted cumulative gain (NDCG) per document, at various top-k
	//  cutoff levels.
	//
	//  NDCG measures the ranking quality, giving higher relevance to top
	//  results.
	//
	//  Example (top-3):
	//   Suppose [SampleQuery][google.cloud.discoveryengine.v1beta.SampleQuery]
	//   with three retrieved documents (D1, D2, D3) and binary relevance
	//   judgements (1 for relevant, 0 for not relevant):
	//
	//    Retrieved:  [D3 (0), D1 (1), D2 (1)]
	//    Ideal:      [D1 (1), D2 (1), D3 (0)]
	//
	//    Calculate NDCG@3 for each
	//    [SampleQuery][google.cloud.discoveryengine.v1beta.SampleQuery]:
	//     * DCG@3: 0/log2(1+1) + 1/log2(2+1) + 1/log2(3+1) = 1.13
	//     * Ideal DCG@3: 1/log2(1+1) + 1/log2(2+1) + 0/log2(3+1) = 1.63
	//     * NDCG@3: 1.13/1.63 = 0.693
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.QualityMetrics.doc_ndcg
	DocNdcg *QualityMetrics_TopkMetrics `json:"docNdcg,omitempty"`

	// Recall per page, at various top-k cutoff levels.
	//
	//  Recall is the fraction of relevant pages retrieved out of all relevant
	//  pages.
	//
	//  Example (top-5):
	//   * For a single
	//   [SampleQuery][google.cloud.discoveryengine.v1beta.SampleQuery], if 3 out
	//   of 5 relevant pages are retrieved in the top-5, recall@5 = 3/5 = 0.6
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.QualityMetrics.page_recall
	PageRecall *QualityMetrics_TopkMetrics `json:"pageRecall,omitempty"`

	// Normalized discounted cumulative gain (NDCG) per page, at various top-k
	//  cutoff levels.
	//
	//  NDCG measures the ranking quality, giving higher relevance to top
	//  results.
	//
	//  Example (top-3):
	//   Suppose [SampleQuery][google.cloud.discoveryengine.v1beta.SampleQuery]
	//   with three retrieved pages (P1, P2, P3) and binary relevance judgements (1
	//   for relevant, 0 for not relevant):
	//
	//    Retrieved:  [P3 (0), P1 (1), P2 (1)]
	//    Ideal:      [P1 (1), P2 (1), P3 (0)]
	//
	//    Calculate NDCG@3 for
	//    [SampleQuery][google.cloud.discoveryengine.v1beta.SampleQuery]:
	//     * DCG@3: 0/log2(1+1) + 1/log2(2+1) + 1/log2(3+1) = 1.13
	//     * Ideal DCG@3: 1/log2(1+1) + 1/log2(2+1) + 0/log2(3+1) = 1.63
	//     * NDCG@3: 1.13/1.63 = 0.693
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.QualityMetrics.page_ndcg
	PageNdcg *QualityMetrics_TopkMetrics `json:"pageNdcg,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.QualityMetrics.TopkMetrics
type QualityMetrics_TopkMetrics struct {
	// The top-1 value.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.QualityMetrics.TopkMetrics.top_1
	Top1 *float64 `json:"top1,omitempty"`

	// The top-3 value.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.QualityMetrics.TopkMetrics.top_3
	Top3 *float64 `json:"top3,omitempty"`

	// The top-5 value.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.QualityMetrics.TopkMetrics.top_5
	Top5 *float64 `json:"top5,omitempty"`

	// The top-10 value.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.QualityMetrics.TopkMetrics.top_10
	Top10 *float64 `json:"top10,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest
type SearchRequest struct {
	// Required. The resource name of the Search serving config, such as
	//  `projects/*/locations/global/collections/default_collection/engines/*/servingConfigs/default_serving_config`,
	//  or
	//  `projects/*/locations/global/collections/default_collection/dataStores/default_data_store/servingConfigs/default_serving_config`.
	//  This field is used to identify the serving configuration name, set
	//  of models used to make the search.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.serving_config
	ServingConfig *string `json:"servingConfig,omitempty"`

	// The branch resource name, such as
	//  `projects/*/locations/global/collections/default_collection/dataStores/default_data_store/branches/0`.
	//
	//  Use `default_branch` as the branch ID or leave this field empty, to search
	//  documents under the default branch.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.branch
	Branch *string `json:"branch,omitempty"`

	// Raw search query.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.query
	Query *string `json:"query,omitempty"`

	// Raw image query.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.image_query
	ImageQuery *SearchRequest_ImageQuery `json:"imageQuery,omitempty"`

	// Maximum number of [Document][google.cloud.discoveryengine.v1beta.Document]s
	//  to return. The maximum allowed value depends on the data type. Values above
	//  the maximum value are coerced to the maximum value.
	//
	//  * Websites with basic indexing: Default `10`, Maximum `25`.
	//  * Websites with advanced indexing: Default `25`, Maximum `50`.
	//  * Other: Default `50`, Maximum `100`.
	//
	//  If this field is negative, an  `INVALID_ARGUMENT` is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.page_size
	PageSize *int32 `json:"pageSize,omitempty"`

	// A page token received from a previous
	//  [SearchService.Search][google.cloud.discoveryengine.v1beta.SearchService.Search]
	//  call. Provide this to retrieve the subsequent page.
	//
	//  When paginating, all other parameters provided to
	//  [SearchService.Search][google.cloud.discoveryengine.v1beta.SearchService.Search]
	//  must match the call that provided the page token. Otherwise, an
	//   `INVALID_ARGUMENT`  error is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.page_token
	PageToken *string `json:"pageToken,omitempty"`

	// A 0-indexed integer that specifies the current offset (that is, starting
	//  result location, amongst the
	//  [Document][google.cloud.discoveryengine.v1beta.Document]s deemed by the API
	//  as relevant) in search results. This field is only considered if
	//  [page_token][google.cloud.discoveryengine.v1beta.SearchRequest.page_token]
	//  is unset.
	//
	//  If this field is negative, an  `INVALID_ARGUMENT`  is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.offset
	Offset *int32 `json:"offset,omitempty"`

	// The maximum number of results to return for OneBox.
	//  This applies to each OneBox type individually.
	//  Default number is 10.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.one_box_page_size
	OneBoxPageSize *int32 `json:"oneBoxPageSize,omitempty"`

	// Specs defining dataStores to filter on in a search call and configurations
	//  for those dataStores. This is only considered for engines with multiple
	//  dataStores use case. For single dataStore within an engine, they should
	//  use the specs at the top level.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.data_store_specs
	DataStoreSpecs []SearchRequest_DataStoreSpec `json:"dataStoreSpecs,omitempty"`

	// The filter syntax consists of an expression language for constructing a
	//  predicate from one or more fields of the documents being filtered. Filter
	//  expression is case-sensitive.
	//
	//  If this field is unrecognizable, an  `INVALID_ARGUMENT`  is returned.
	//
	//  Filtering in Vertex AI Search is done by mapping the LHS filter key to a
	//  key property defined in the Vertex AI Search backend -- this mapping is
	//  defined by the customer in their schema. For example a media customer might
	//  have a field 'name' in their schema. In this case the filter would look
	//  like this: filter --> name:'ANY("king kong")'
	//
	//  For more information about filtering including syntax and filter
	//  operators, see
	//  [Filter](https://cloud.google.com/generative-ai-app-builder/docs/filter-search-metadata)
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.filter
	Filter *string `json:"filter,omitempty"`

	// The default filter that is applied when a user performs a search without
	//  checking any filters on the search page.
	//
	//  The filter applied to every search request when quality improvement such as
	//  query expansion is needed. In the case a query does not have a sufficient
	//  amount of results this filter will be used to determine whether or not to
	//  enable the query expansion flow. The original filter will still be used for
	//  the query expanded search.
	//  This field is strongly recommended to achieve high search quality.
	//
	//  For more information about filter syntax, see
	//  [SearchRequest.filter][google.cloud.discoveryengine.v1beta.SearchRequest.filter].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.canonical_filter
	CanonicalFilter *string `json:"canonicalFilter,omitempty"`

	// The order in which documents are returned. Documents can be ordered by
	//  a field in an [Document][google.cloud.discoveryengine.v1beta.Document]
	//  object. Leave it unset if ordered by relevance. `order_by` expression is
	//  case-sensitive.
	//
	//  For more information on ordering the website search results, see
	//  [Order web search
	//  results](https://cloud.google.com/generative-ai-app-builder/docs/order-web-search-results).
	//  For more information on ordering the healthcare search results, see
	//  [Order healthcare search
	//  results](https://cloud.google.com/generative-ai-app-builder/docs/order-hc-results).
	//  If this field is unrecognizable, an `INVALID_ARGUMENT` is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.order_by
	OrderBy *string `json:"orderBy,omitempty"`

	// Information about the end user.
	//  Highly recommended for analytics.
	//  [UserInfo.user_agent][google.cloud.discoveryengine.v1beta.UserInfo.user_agent]
	//  is used to deduce `device_type` for analytics.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.user_info
	UserInfo *UserInfo `json:"userInfo,omitempty"`

	// The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	//  information, see [Standard
	//  fields](https://cloud.google.com/apis/design/standard_fields). This field
	//  helps to better interpret the query. If a value isn't specified, the query
	//  language code is automatically detected, which may not be accurate.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// The Unicode country/region code (CLDR) of a location, such as "US" and
	//  "419". For more information, see [Standard
	//  fields](https://cloud.google.com/apis/design/standard_fields). If set,
	//  then results will be boosted based on the region_code provided.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.region_code
	RegionCode *string `json:"regionCode,omitempty"`

	// Facet specifications for faceted search. If empty, no facets are returned.
	//
	//  A maximum of 100 values are allowed. Otherwise, an  `INVALID_ARGUMENT`
	//  error is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.facet_specs
	FacetSpecs []SearchRequest_FacetSpec `json:"facetSpecs,omitempty"`

	// Boost specification to boost certain documents.
	//  For more information on boosting, see
	//  [Boosting](https://cloud.google.com/generative-ai-app-builder/docs/boost-search-results)
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.boost_spec
	BoostSpec *SearchRequest_BoostSpec `json:"boostSpec,omitempty"`

	// TODO: unsupported map type with key string and value message


	// The query expansion specification that specifies the conditions under which
	//  query expansion occurs.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.query_expansion_spec
	QueryExpansionSpec *SearchRequest_QueryExpansionSpec `json:"queryExpansionSpec,omitempty"`

	// The spell correction specification that specifies the mode under
	//  which spell correction takes effect.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.spell_correction_spec
	SpellCorrectionSpec *SearchRequest_SpellCorrectionSpec `json:"spellCorrectionSpec,omitempty"`

	// A unique identifier for tracking visitors. For example, this could be
	//  implemented with an HTTP cookie, which should be able to uniquely identify
	//  a visitor on a single device. This unique identifier should not change if
	//  the visitor logs in or out of the website.
	//
	//  This field should NOT have a fixed value such as `unknown_visitor`.
	//
	//  This should be the same identifier as
	//  [UserEvent.user_pseudo_id][google.cloud.discoveryengine.v1beta.UserEvent.user_pseudo_id]
	//  and
	//  [CompleteQueryRequest.user_pseudo_id][google.cloud.discoveryengine.v1beta.CompleteQueryRequest.user_pseudo_id]
	//
	//  The field must be a UTF-8 encoded string with a length limit of 128
	//  characters. Otherwise, an  `INVALID_ARGUMENT`  error is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.user_pseudo_id
	UserPseudoID *string `json:"userPseudoID,omitempty"`

	// A specification for configuring the behavior of content search.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.content_search_spec
	ContentSearchSpec *SearchRequest_ContentSearchSpec `json:"contentSearchSpec,omitempty"`

	// Uses the provided embedding to do additional semantic document retrieval.
	//  The retrieval is based on the dot product of
	//  [SearchRequest.EmbeddingSpec.EmbeddingVector.vector][google.cloud.discoveryengine.v1beta.SearchRequest.EmbeddingSpec.EmbeddingVector.vector]
	//  and the document embedding that is provided in
	//  [SearchRequest.EmbeddingSpec.EmbeddingVector.field_path][google.cloud.discoveryengine.v1beta.SearchRequest.EmbeddingSpec.EmbeddingVector.field_path].
	//
	//  If
	//  [SearchRequest.EmbeddingSpec.EmbeddingVector.field_path][google.cloud.discoveryengine.v1beta.SearchRequest.EmbeddingSpec.EmbeddingVector.field_path]
	//  is not provided, it will use
	//  [ServingConfig.EmbeddingConfig.field_path][google.cloud.discoveryengine.v1beta.ServingConfig.embedding_config].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.embedding_spec
	EmbeddingSpec *SearchRequest_EmbeddingSpec `json:"embeddingSpec,omitempty"`

	// The ranking expression controls the customized ranking on retrieval
	//  documents. This overrides
	//  [ServingConfig.ranking_expression][google.cloud.discoveryengine.v1beta.ServingConfig.ranking_expression].
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
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ranking_expression
	RankingExpression *string `json:"rankingExpression,omitempty"`

	// Whether to turn on safe search. This is only supported for
	//  website search.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.safe_search
	SafeSearch *bool `json:"safeSearch,omitempty"`

	// The user labels applied to a resource must meet the following requirements:
	//
	//  * Each resource can have multiple labels, up to a maximum of 64.
	//  * Each label must be a key-value pair.
	//  * Keys have a minimum length of 1 character and a maximum length of 63
	//    characters and cannot be empty. Values can be empty and have a maximum
	//    length of 63 characters.
	//  * Keys and values can contain only lowercase letters, numeric characters,
	//    underscores, and dashes. All characters must use UTF-8 encoding, and
	//    international characters are allowed.
	//  * The key portion of a label must be unique. However, you can use the same
	//    key with multiple resources.
	//  * Keys must start with a lowercase letter or international character.
	//
	//  See [Google Cloud
	//  Document](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements)
	//  for more details.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.user_labels
	UserLabels map[string]string `json:"userLabels,omitempty"`

	// If `naturalLanguageQueryUnderstandingSpec` is not specified, no additional
	//  natural language query understanding will be done.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.natural_language_query_understanding_spec
	NaturalLanguageQueryUnderstandingSpec *SearchRequest_NaturalLanguageQueryUnderstandingSpec `json:"naturalLanguageQueryUnderstandingSpec,omitempty"`

	// Search as you type configuration. Only supported for the
	//  [IndustryVertical.MEDIA][google.cloud.discoveryengine.v1beta.IndustryVertical.MEDIA]
	//  vertical.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.search_as_you_type_spec
	SearchAsYouTypeSpec *SearchRequest_SearchAsYouTypeSpec `json:"searchAsYouTypeSpec,omitempty"`

	// The session resource name. Optional.
	//
	//  Session allows users to do multi-turn /search API calls or coordination
	//  between /search API calls and /answer API calls.
	//
	//  Example #1 (multi-turn /search API calls):
	//    1. Call /search API with the auto-session mode (see below).
	//    2. Call /search API with the session ID generated in the first call.
	//       Here, the previous search query gets considered in query
	//       standing. I.e., if the first query is "How did Alphabet do in 2022?"
	//       and the current query is "How about 2023?", the current query will
	//       be interpreted as "How did Alphabet do in 2023?".
	//
	//  Example #2 (coordination between /search API calls and /answer API calls):
	//    1. Call /search API with the auto-session mode (see below).
	//    2. Call /answer API with the session ID generated in the first call.
	//       Here, the answer generation happens in the context of the search
	//       results from the first search call.
	//
	//  Auto-session mode: when `projects/.../sessions/-` is used, a new session
	//  gets automatically created. Otherwise, users can use the create-session API
	//  to create a session manually.
	//
	//  Multi-turn Search feature is currently at private GA stage. Please use
	//  v1alpha or v1beta version instead before we launch this feature to public
	//  GA. Or ask for allowlisting through Google Support team.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.session
	Session *string `json:"session,omitempty"`

	// Session specification.
	//
	//  Can be used only when `session` is set.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.session_spec
	SessionSpec *SearchRequest_SessionSpec `json:"sessionSpec,omitempty"`

	// The relevance threshold of the search results.
	//
	//  Default to Google defined threshold, leveraging a balance of
	//  precision and recall to deliver both highly accurate results and
	//  comprehensive coverage of relevant information.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.relevance_threshold
	RelevanceThreshold *string `json:"relevanceThreshold,omitempty"`

	// The specification for personalization.
	//
	//  Notice that if both
	//  [ServingConfig.personalization_spec][google.cloud.discoveryengine.v1beta.ServingConfig.personalization_spec]
	//  and
	//  [SearchRequest.personalization_spec][google.cloud.discoveryengine.v1beta.SearchRequest.personalization_spec]
	//  are set,
	//  [SearchRequest.personalization_spec][google.cloud.discoveryengine.v1beta.SearchRequest.personalization_spec]
	//  overrides
	//  [ServingConfig.personalization_spec][google.cloud.discoveryengine.v1beta.ServingConfig.personalization_spec].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.personalization_spec
	PersonalizationSpec *SearchRequest_PersonalizationSpec `json:"personalizationSpec,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec
type SearchRequest_BoostSpec struct {
	// Condition boost specifications. If a document matches multiple conditions
	//  in the specifictions, boost scores from these specifications are all
	//  applied and combined in a non-linear way. Maximum number of
	//  specifications is 20.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec.condition_boost_specs
	ConditionBoostSpecs []SearchRequest_BoostSpec_ConditionBoostSpec `json:"conditionBoostSpecs,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec.ConditionBoostSpec
type SearchRequest_BoostSpec_ConditionBoostSpec struct {
	// An expression which specifies a boost condition. The syntax and
	//  supported fields are the same as a filter expression. See
	//  [SearchRequest.filter][google.cloud.discoveryengine.v1beta.SearchRequest.filter]
	//  for detail syntax and limitations.
	//
	//  Examples:
	//
	//  * To boost documents with document ID "doc_1" or "doc_2", and
	//  color "Red" or "Blue":
	//  `(document_id: ANY("doc_1", "doc_2")) AND (color: ANY("Red", "Blue"))`
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec.ConditionBoostSpec.condition
	Condition *string `json:"condition,omitempty"`

	// Strength of the condition boost, which should be in [-1, 1]. Negative
	//  boost means demotion. Default is 0.0.
	//
	//  Setting to 1.0 gives the document a big promotion. However, it does
	//  not necessarily mean that the boosted document will be the top result
	//  at all times, nor that other documents will be excluded. Results
	//  could still be shown even when none of them matches the condition.
	//  And results that are significantly more relevant to the search query
	//  can still trump your heavily favored but irrelevant documents.
	//
	//  Setting to -1.0 gives the document a big demotion. However, results
	//  that are deeply relevant might still be shown. The document will have
	//  an upstream battle to get a fairly high ranking, but it is not
	//  blocked out completely.
	//
	//  Setting to 0.0 means no boost applied. The boosting condition is
	//  ignored. Only one of the (condition, boost) combination or the
	//  boost_control_spec below are set. If both are set then the global boost
	//  is ignored and the more fine-grained boost_control_spec is applied.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec.ConditionBoostSpec.boost
	Boost *float32 `json:"boost,omitempty"`

	// Complex specification for custom ranking based on customer defined
	//  attribute value.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec.ConditionBoostSpec.boost_control_spec
	BoostControlSpec *SearchRequest_BoostSpec_ConditionBoostSpec_BoostControlSpec `json:"boostControlSpec,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec.ConditionBoostSpec.BoostControlSpec
type SearchRequest_BoostSpec_ConditionBoostSpec_BoostControlSpec struct {
	// The name of the field whose value will be used to determine the
	//  boost amount.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec.ConditionBoostSpec.BoostControlSpec.field_name
	FieldName *string `json:"fieldName,omitempty"`

	// The attribute type to be used to determine the boost amount. The
	//  attribute value can be derived from the field value of the specified
	//  field_name. In the case of numerical it is straightforward i.e.
	//  attribute_value = numerical_field_value. In the case of freshness
	//  however, attribute_value = (time.now() - datetime_field_value).
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec.ConditionBoostSpec.BoostControlSpec.attribute_type
	AttributeType *string `json:"attributeType,omitempty"`

	// The interpolation type to be applied to connect the control points
	//  listed below.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec.ConditionBoostSpec.BoostControlSpec.interpolation_type
	InterpolationType *string `json:"interpolationType,omitempty"`

	// The control points used to define the curve. The monotonic function
	//  (defined through the interpolation_type above) passes through the
	//  control points listed here.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec.ConditionBoostSpec.BoostControlSpec.control_points
	ControlPoints []SearchRequest_BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint `json:"controlPoints,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec.ConditionBoostSpec.BoostControlSpec.ControlPoint
type SearchRequest_BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint struct {
	// Can be one of:
	//  1. The numerical field value.
	//  2. The duration spec for freshness:
	//  The value must be formatted as an XSD `dayTimeDuration` value (a
	//  restricted subset of an ISO 8601 duration value). The pattern for
	//  this is: `[nD][T[nH][nM][nS]]`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec.ConditionBoostSpec.BoostControlSpec.ControlPoint.attribute_value
	AttributeValue *string `json:"attributeValue,omitempty"`

	// The value between -1 to 1 by which to boost the score if the
	//  attribute_value evaluates to the value specified above.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.BoostSpec.ConditionBoostSpec.BoostControlSpec.ControlPoint.boost_amount
	BoostAmount *float32 `json:"boostAmount,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec
type SearchRequest_ContentSearchSpec struct {
	// If `snippetSpec` is not specified, snippets are not included in the
	//  search response.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.snippet_spec
	SnippetSpec *SearchRequest_ContentSearchSpec_SnippetSpec `json:"snippetSpec,omitempty"`

	// If `summarySpec` is not specified, summaries are not included in the
	//  search response.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.summary_spec
	SummarySpec *SearchRequest_ContentSearchSpec_SummarySpec `json:"summarySpec,omitempty"`

	// If there is no extractive_content_spec provided, there will be no
	//  extractive answer in the search response.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.extractive_content_spec
	ExtractiveContentSpec *SearchRequest_ContentSearchSpec_ExtractiveContentSpec `json:"extractiveContentSpec,omitempty"`

	// Specifies the search result mode. If unspecified, the
	//  search result mode defaults to `DOCUMENTS`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.search_result_mode
	SearchResultMode *string `json:"searchResultMode,omitempty"`

	// Specifies the chunk spec to be returned from the search response.
	//  Only available if the
	//  [SearchRequest.ContentSearchSpec.search_result_mode][google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.search_result_mode]
	//  is set to
	//  [CHUNKS][google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SearchResultMode.CHUNKS]
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.chunk_spec
	ChunkSpec *SearchRequest_ContentSearchSpec_ChunkSpec `json:"chunkSpec,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.ChunkSpec
type SearchRequest_ContentSearchSpec_ChunkSpec struct {
	// The number of previous chunks to be returned of the current chunk. The
	//  maximum allowed value is 3.
	//  If not specified, no previous chunks will be returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.ChunkSpec.num_previous_chunks
	NumPreviousChunks *int32 `json:"numPreviousChunks,omitempty"`

	// The number of next chunks to be returned of the current chunk. The
	//  maximum allowed value is 3.
	//  If not specified, no next chunks will be returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.ChunkSpec.num_next_chunks
	NumNextChunks *int32 `json:"numNextChunks,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.ExtractiveContentSpec
type SearchRequest_ContentSearchSpec_ExtractiveContentSpec struct {
	// The maximum number of extractive answers returned in each search
	//  result.
	//
	//  An extractive answer is a verbatim answer extracted from the original
	//  document, which provides a precise and contextually relevant answer to
	//  the search query.
	//
	//  If the number of matching answers is less than the
	//  `max_extractive_answer_count`, return all of the answers. Otherwise,
	//  return the `max_extractive_answer_count`.
	//
	//  At most five answers are returned for each
	//  [SearchResult][google.cloud.discoveryengine.v1beta.SearchResponse.SearchResult].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.ExtractiveContentSpec.max_extractive_answer_count
	MaxExtractiveAnswerCount *int32 `json:"maxExtractiveAnswerCount,omitempty"`

	// The max number of extractive segments returned in each search result.
	//  Only applied if the
	//  [DataStore][google.cloud.discoveryengine.v1beta.DataStore] is set to
	//  [DataStore.ContentConfig.CONTENT_REQUIRED][google.cloud.discoveryengine.v1beta.DataStore.ContentConfig.CONTENT_REQUIRED]
	//  or
	//  [DataStore.solution_types][google.cloud.discoveryengine.v1beta.DataStore.solution_types]
	//  is
	//  [SOLUTION_TYPE_CHAT][google.cloud.discoveryengine.v1beta.SolutionType.SOLUTION_TYPE_CHAT].
	//
	//  An extractive segment is a text segment extracted from the original
	//  document that is relevant to the search query, and, in general, more
	//  verbose than an extractive answer. The segment could then be used as
	//  input for LLMs to generate summaries and answers.
	//
	//  If the number of matching segments is less than
	//  `max_extractive_segment_count`, return all of the segments. Otherwise,
	//  return the `max_extractive_segment_count`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.ExtractiveContentSpec.max_extractive_segment_count
	MaxExtractiveSegmentCount *int32 `json:"maxExtractiveSegmentCount,omitempty"`

	// Specifies whether to return the confidence score from the extractive
	//  segments in each search result. This feature is available only for new
	//  or allowlisted data stores. To allowlist your data store,
	//  contact your Customer Engineer. The default value is `false`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.ExtractiveContentSpec.return_extractive_segment_score
	ReturnExtractiveSegmentScore *bool `json:"returnExtractiveSegmentScore,omitempty"`

	// Specifies whether to also include the adjacent from each selected
	//  segments.
	//  Return at most `num_previous_segments` segments before each selected
	//  segments.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.ExtractiveContentSpec.num_previous_segments
	NumPreviousSegments *int32 `json:"numPreviousSegments,omitempty"`

	// Return at most `num_next_segments` segments after each selected
	//  segments.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.ExtractiveContentSpec.num_next_segments
	NumNextSegments *int32 `json:"numNextSegments,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SnippetSpec
type SearchRequest_ContentSearchSpec_SnippetSpec struct {
	// [DEPRECATED] This field is deprecated. To control snippet return, use
	//  `return_snippet` field. For backwards compatibility, we will return
	//  snippet if max_snippet_count > 0.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SnippetSpec.max_snippet_count
	MaxSnippetCount *int32 `json:"maxSnippetCount,omitempty"`

	// [DEPRECATED] This field is deprecated and will have no affect on the
	//  snippet.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SnippetSpec.reference_only
	ReferenceOnly *bool `json:"referenceOnly,omitempty"`

	// If `true`, then return snippet. If no snippet can be generated, we
	//  return "No snippet is available for this page." A `snippet_status` with
	//  `SUCCESS` or `NO_SNIPPET_AVAILABLE` will also be returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SnippetSpec.return_snippet
	ReturnSnippet *bool `json:"returnSnippet,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec
type SearchRequest_ContentSearchSpec_SummarySpec struct {
	// The number of top results to generate the summary from. If the number
	//  of results returned is less than `summaryResultCount`, the summary is
	//  generated from all of the results.
	//
	//  At most 10 results for documents mode, or 50 for chunks mode, can be
	//  used to generate a summary. The chunks mode is used when
	//  [SearchRequest.ContentSearchSpec.search_result_mode][google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.search_result_mode]
	//  is set to
	//  [CHUNKS][google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SearchResultMode.CHUNKS].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.summary_result_count
	SummaryResultCount *int32 `json:"summaryResultCount,omitempty"`

	// Specifies whether to include citations in the summary. The default
	//  value is `false`.
	//
	//  When this field is set to `true`, summaries include in-line citation
	//  numbers.
	//
	//  Example summary including citations:
	//
	//  BigQuery is Google Cloud's fully managed and completely serverless
	//  enterprise data warehouse [1]. BigQuery supports all data types, works
	//  across clouds, and has built-in machine learning and business
	//  intelligence, all within a unified platform [2, 3].
	//
	//  The citation numbers refer to the returned search results and are
	//  1-indexed. For example, [1] means that the sentence is attributed to
	//  the first search result. [2, 3] means that the sentence is attributed
	//  to both the second and third search results.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.include_citations
	IncludeCitations *bool `json:"includeCitations,omitempty"`

	// Specifies whether to filter out adversarial queries. The default value
	//  is `false`.
	//
	//  Google employs search-query classification to detect adversarial
	//  queries. No summary is returned if the search query is classified as an
	//  adversarial query. For example, a user might ask a question regarding
	//  negative comments about the company or submit a query designed to
	//  generate unsafe, policy-violating output. If this field is set to
	//  `true`, we skip generating summaries for adversarial queries and return
	//  fallback messages instead.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.ignore_adversarial_query
	IgnoreAdversarialQuery *bool `json:"ignoreAdversarialQuery,omitempty"`

	// Specifies whether to filter out queries that are not summary-seeking.
	//  The default value is `false`.
	//
	//  Google employs search-query classification to detect summary-seeking
	//  queries. No summary is returned if the search query is classified as a
	//  non-summary seeking query. For example, `why is the sky blue` and `Who
	//  is the best soccer player in the world?` are summary-seeking queries,
	//  but `SFO airport` and `world cup 2026` are not. They are most likely
	//  navigational queries. If this field is set to `true`, we skip
	//  generating summaries for non-summary seeking queries and return
	//  fallback messages instead.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.ignore_non_summary_seeking_query
	IgnoreNonSummarySeekingQuery *bool `json:"ignoreNonSummarySeekingQuery,omitempty"`

	// Specifies whether to filter out queries that have low relevance. The
	//  default value is `false`.
	//
	//  If this field is set to `false`, all search results are used regardless
	//  of relevance to generate answers. If set to `true`, only queries with
	//  high relevance search results will generate answers.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.ignore_low_relevant_content
	IgnoreLowRelevantContent *bool `json:"ignoreLowRelevantContent,omitempty"`

	// Optional. Specifies whether to filter out jail-breaking queries. The
	//  default value is `false`.
	//
	//  Google employs search-query classification to detect jail-breaking
	//  queries. No summary is returned if the search query is classified as a
	//  jail-breaking query. A user might add instructions to the query to
	//  change the tone, style, language, content of the answer, or ask the
	//  model to act as a different entity, e.g. "Reply in the tone of a
	//  competing company's CEO". If this field is set to `true`, we skip
	//  generating summaries for jail-breaking queries and return fallback
	//  messages instead.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.ignore_jail_breaking_query
	IgnoreJailBreakingQuery *bool `json:"ignoreJailBreakingQuery,omitempty"`

	// If specified, the spec will be used to modify the prompt provided to
	//  the LLM.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.model_prompt_spec
	ModelPromptSpec *SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec `json:"modelPromptSpec,omitempty"`

	// Language code for Summary. Use language tags defined by
	//  [BCP47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt).
	//  Note: This is an experimental feature.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// If specified, the spec will be used to modify the model specification
	//  provided to the LLM.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.model_spec
	ModelSpec *SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec `json:"modelSpec,omitempty"`

	// If true, answer will be generated from most relevant chunks from top
	//  search results. This feature will improve summary quality.
	//  Note that with this feature enabled, not all top search results
	//  will be referenced and included in the reference list, so the citation
	//  source index only points to the search results listed in the reference
	//  list.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.use_semantic_chunks
	UseSemanticChunks *bool `json:"useSemanticChunks,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.ModelPromptSpec
type SearchRequest_ContentSearchSpec_SummarySpec_ModelPromptSpec struct {
	// Text at the beginning of the prompt that instructs the assistant.
	//  Examples are available in the user guide.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.ModelPromptSpec.preamble
	Preamble *string `json:"preamble,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.ModelSpec
type SearchRequest_ContentSearchSpec_SummarySpec_ModelSpec struct {
	// The model version used to generate the summary.
	//
	//  Supported values are:
	//
	//  * `stable`: string. Default value when no value is specified. Uses a
	//     generally available, fine-tuned model. For more information, see
	//     [Answer generation model versions and
	//     lifecycle](https://cloud.google.com/generative-ai-app-builder/docs/answer-generation-models).
	//  * `preview`: string. (Public preview) Uses a preview model. For more
	//     information, see
	//     [Answer generation model versions and
	//     lifecycle](https://cloud.google.com/generative-ai-app-builder/docs/answer-generation-models).
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.ModelSpec.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.DataStoreSpec
type SearchRequest_DataStoreSpec struct {
	// Required. Full resource name of
	//  [DataStore][google.cloud.discoveryengine.v1beta.DataStore], such as
	//  `projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.DataStoreSpec.data_store
	DataStore *string `json:"dataStore,omitempty"`

	// Optional. Filter specification to filter documents in the data store
	//  specified by data_store field. For more information on filtering, see
	//  [Filtering](https://cloud.google.com/generative-ai-app-builder/docs/filter-search-metadata)
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.DataStoreSpec.filter
	Filter *string `json:"filter,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.EmbeddingSpec
type SearchRequest_EmbeddingSpec struct {
	// The embedding vector used for retrieval. Limit to 1.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.EmbeddingSpec.embedding_vectors
	EmbeddingVectors []SearchRequest_EmbeddingSpec_EmbeddingVector `json:"embeddingVectors,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.EmbeddingSpec.EmbeddingVector
type SearchRequest_EmbeddingSpec_EmbeddingVector struct {
	// Embedding field path in schema.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.EmbeddingSpec.EmbeddingVector.field_path
	FieldPath *string `json:"fieldPath,omitempty"`

	// Query embedding vector.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.EmbeddingSpec.EmbeddingVector.vector
	Vector []float32 `json:"vector,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec
type SearchRequest_FacetSpec struct {
	// Required. The facet key specification.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.facet_key
	FacetKey *SearchRequest_FacetSpec_FacetKey `json:"facetKey,omitempty"`

	// Maximum facet values that are returned for this facet. If
	//  unspecified, defaults to 20. The maximum allowed value is 300. Values
	//  above 300 are coerced to 300.
	//  For aggregation in healthcare search, when the [FacetKey.key] is
	//  "healthcare_aggregation_key", the limit will be overridden to
	//  10,000 internally, regardless of the value set here.
	//
	//  If this field is negative, an  `INVALID_ARGUMENT`  is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.limit
	Limit *int32 `json:"limit,omitempty"`

	// List of keys to exclude when faceting.
	//
	//
	//  By default,
	//  [FacetKey.key][google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.FacetKey.key]
	//  is not excluded from the filter unless it is listed in this field.
	//
	//  Listing a facet key in this field allows its values to appear as facet
	//  results, even when they are filtered out of search results. Using this
	//  field does not affect what search results are returned.
	//
	//  For example, suppose there are 100 documents with the color facet "Red"
	//  and 200 documents with the color facet "Blue". A query containing the
	//  filter "color:ANY("Red")" and having "color" as
	//  [FacetKey.key][google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.FacetKey.key]
	//  would by default return only "Red" documents in the search results, and
	//  also return "Red" with count 100 as the only color facet. Although there
	//  are also blue documents available, "Blue" would not be shown as an
	//  available facet value.
	//
	//  If "color" is listed in "excludedFilterKeys", then the query returns the
	//  facet values "Red" with count 100 and "Blue" with count 200, because the
	//  "color" key is now excluded from the filter. Because this field doesn't
	//  affect search results, the search results are still correctly filtered to
	//  return only "Red" documents.
	//
	//  A maximum of 100 values are allowed. Otherwise, an  `INVALID_ARGUMENT`
	//  error is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.excluded_filter_keys
	ExcludedFilterKeys []string `json:"excludedFilterKeys,omitempty"`

	// Enables dynamic position for this facet. If set to true, the position of
	//  this facet among all facets in the response is determined automatically.
	//  If dynamic facets are enabled, it is ordered together.
	//  If set to false, the position of this facet in the
	//  response is the same as in the request, and it is ranked before
	//  the facets with dynamic position enable and all dynamic facets.
	//
	//  For example, you may always want to have rating facet returned in
	//  the response, but it's not necessarily to always display the rating facet
	//  at the top. In that case, you can set enable_dynamic_position to true so
	//  that the position of rating facet in response is determined
	//  automatically.
	//
	//  Another example, assuming you have the following facets in the request:
	//
	//  * "rating", enable_dynamic_position = true
	//
	//  * "price", enable_dynamic_position = false
	//
	//  * "brands", enable_dynamic_position = false
	//
	//  And also you have a dynamic facets enabled, which generates a facet
	//  `gender`. Then the final order of the facets in the response can be
	//  ("price", "brands", "rating", "gender") or ("price", "brands", "gender",
	//  "rating") depends on how API orders "gender" and "rating" facets.
	//  However, notice that "price" and "brands" are always
	//  ranked at first and second position because their enable_dynamic_position
	//  is false.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.enable_dynamic_position
	EnableDynamicPosition *bool `json:"enableDynamicPosition,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.FacetKey
type SearchRequest_FacetSpec_FacetKey struct {
	// Required. Supported textual and numerical facet keys in
	//  [Document][google.cloud.discoveryengine.v1beta.Document] object, over
	//  which the facet values are computed. Facet key is case-sensitive.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.FacetKey.key
	Key *string `json:"key,omitempty"`

	// Set only if values should be bucketed into intervals. Must be set
	//  for facets with numerical values. Must not be set for facet with text
	//  values. Maximum number of intervals is 30.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.FacetKey.intervals
	Intervals []Interval `json:"intervals,omitempty"`

	// Only get facet for the given restricted values. Only supported on
	//  textual fields. For example, suppose "category" has three values
	//  "Action > 2022", "Action > 2021" and "Sci-Fi > 2022". If set
	//  "restricted_values" to "Action > 2022", the "category" facet only
	//  contains "Action > 2022". Only supported on textual fields. Maximum
	//  is 10.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.FacetKey.restricted_values
	RestrictedValues []string `json:"restrictedValues,omitempty"`

	// Only get facet values that start with the given string prefix. For
	//  example, suppose "category" has three values "Action > 2022",
	//  "Action > 2021" and "Sci-Fi > 2022". If set "prefixes" to "Action", the
	//  "category" facet only contains "Action > 2022" and "Action > 2021".
	//  Only supported on textual fields. Maximum is 10.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.FacetKey.prefixes
	Prefixes []string `json:"prefixes,omitempty"`

	// Only get facet values that contain the given strings. For example,
	//  suppose "category" has three values "Action > 2022",
	//  "Action > 2021" and "Sci-Fi > 2022". If set "contains" to "2022", the
	//  "category" facet only contains "Action > 2022" and "Sci-Fi > 2022".
	//  Only supported on textual fields. Maximum is 10.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.FacetKey.contains
	Contains []string `json:"contains,omitempty"`

	// True to make facet keys case insensitive when getting faceting
	//  values with prefixes or contains; false otherwise.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.FacetKey.case_insensitive
	CaseInsensitive *bool `json:"caseInsensitive,omitempty"`

	// The order in which documents are returned.
	//
	//  Allowed values are:
	//
	//  * "count desc", which means order by
	//  [SearchResponse.Facet.values.count][google.cloud.discoveryengine.v1beta.SearchResponse.Facet.FacetValue.count]
	//  descending.
	//
	//  * "value desc", which means order by
	//  [SearchResponse.Facet.values.value][google.cloud.discoveryengine.v1beta.SearchResponse.Facet.FacetValue.value]
	//  descending.
	//    Only applies to textual facets.
	//
	//  If not set, textual values are sorted in [natural
	//  order](https://en.wikipedia.org/wiki/Natural_sort_order); numerical
	//  intervals are sorted in the order given by
	//  [FacetSpec.FacetKey.intervals][google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.FacetKey.intervals].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.FacetSpec.FacetKey.order_by
	OrderBy *string `json:"orderBy,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.ImageQuery
type SearchRequest_ImageQuery struct {
	// Base64 encoded image bytes. Supported image formats: JPEG, PNG, and
	//  BMP.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.ImageQuery.image_bytes
	ImageBytes *string `json:"imageBytes,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.NaturalLanguageQueryUnderstandingSpec
type SearchRequest_NaturalLanguageQueryUnderstandingSpec struct {
	// The condition under which filter extraction should occur.
	//  Default to [Condition.DISABLED][].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.NaturalLanguageQueryUnderstandingSpec.filter_extraction_condition
	FilterExtractionCondition *string `json:"filterExtractionCondition,omitempty"`

	// Field names used for location-based filtering, where geolocation filters
	//  are detected in natural language search queries.
	//  Only valid when the FilterExtractionCondition is set to `ENABLED`.
	//
	//  If this field is set, it overrides the field names set in
	//  [ServingConfig.geo_search_query_detection_field_names][google.cloud.discoveryengine.v1beta.ServingConfig.geo_search_query_detection_field_names].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.NaturalLanguageQueryUnderstandingSpec.geo_search_query_detection_field_names
	GeoSearchQueryDetectionFieldNames []string `json:"geoSearchQueryDetectionFieldNames,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.PersonalizationSpec
type SearchRequest_PersonalizationSpec struct {
	// The personalization mode of the search request.
	//  Defaults to
	//  [Mode.AUTO][google.cloud.discoveryengine.v1beta.SearchRequest.PersonalizationSpec.Mode.AUTO].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.PersonalizationSpec.mode
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.QueryExpansionSpec
type SearchRequest_QueryExpansionSpec struct {
	// The condition under which query expansion should occur. Default to
	//  [Condition.DISABLED][google.cloud.discoveryengine.v1beta.SearchRequest.QueryExpansionSpec.Condition.DISABLED].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.QueryExpansionSpec.condition
	Condition *string `json:"condition,omitempty"`

	// Whether to pin unexpanded results. If this field is set to true,
	//  unexpanded products are always at the top of the search results, followed
	//  by the expanded results.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.QueryExpansionSpec.pin_unexpanded_results
	PinUnexpandedResults *bool `json:"pinUnexpandedResults,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.SearchAsYouTypeSpec
type SearchRequest_SearchAsYouTypeSpec struct {
	// The condition under which search as you type should occur.
	//  Default to
	//  [Condition.DISABLED][google.cloud.discoveryengine.v1beta.SearchRequest.SearchAsYouTypeSpec.Condition.DISABLED].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.SearchAsYouTypeSpec.condition
	Condition *string `json:"condition,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.SessionSpec
type SearchRequest_SessionSpec struct {
	// If set, the search result gets stored to the "turn" specified by this
	//  query ID.
	//
	//  Example: Let's say the session looks like this:
	//    session {
	//      name: ".../sessions/xxx"
	//      turns {
	//        query { text: "What is foo?" query_id: ".../questions/yyy" }
	//        answer: "Foo is ..."
	//      }
	//      turns {
	//        query { text: "How about bar then?" query_id: ".../questions/zzz" }
	//      }
	//    }
	//
	//  The user can call /search API with a request like this:
	//
	//     session: ".../sessions/xxx"
	//     session_spec { query_id: ".../questions/zzz" }
	//
	//  Then, the API stores the search result, associated with the last turn.
	//  The stored search result can be used by a subsequent /answer API call
	//  (with the session ID and the query ID specified). Also, it is possible
	//  to call /search and /answer in parallel with the same session ID & query
	//  ID.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.SessionSpec.query_id
	QueryID *string `json:"queryID,omitempty"`

	// The number of top search results to persist. The persisted search results
	//  can be used for the subsequent /answer api call.
	//
	//  This field is simliar to the `summary_result_count` field in
	//  [SearchRequest.ContentSearchSpec.SummarySpec.summary_result_count][google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.SummarySpec.summary_result_count].
	//
	//  At most 10 results for documents mode, or 50 for chunks mode.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.SessionSpec.search_result_persistence_count
	SearchResultPersistenceCount *int32 `json:"searchResultPersistenceCount,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.SpellCorrectionSpec
type SearchRequest_SpellCorrectionSpec struct {
	// The mode under which spell correction
	//  replaces the original search query. Defaults to
	//  [Mode.AUTO][google.cloud.discoveryengine.v1beta.SearchRequest.SpellCorrectionSpec.Mode.AUTO].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.SpellCorrectionSpec.mode
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.UserInfo
type UserInfo struct {
	// Highly recommended for logged-in users. Unique identifier for logged-in
	//  user, such as a user name. Don't set for anonymous users.
	//
	//  Always use a hashed value for this ID.
	//
	//  Don't set the field to the same fixed ID for different users. This mixes
	//  the event history of those users together, which results in degraded
	//  model quality.
	//
	//  The field must be a UTF-8 encoded string with a length limit of 128
	//  characters. Otherwise, an `INVALID_ARGUMENT` error is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.UserInfo.user_id
	UserID *string `json:"userID,omitempty"`

	// User agent as included in the HTTP header.
	//
	//  The field must be a UTF-8 encoded string with a length limit of 1,000
	//  characters. Otherwise, an `INVALID_ARGUMENT` error is returned.
	//
	//  This should not be set when using the client side event reporting with
	//  GTM or JavaScript tag in
	//  [UserEventService.CollectUserEvent][google.cloud.discoveryengine.v1beta.UserEventService.CollectUserEvent]
	//  or if
	//  [UserEvent.direct_user_request][google.cloud.discoveryengine.v1beta.UserEvent.direct_user_request]
	//  is set.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.UserInfo.user_agent
	UserAgent *string `json:"userAgent,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Evaluation
type EvaluationObservedState struct {
	// Output only. The metrics produced by the evaluation, averaged across all
	//  [SampleQuery][google.cloud.discoveryengine.v1beta.SampleQuery]s in the
	//  [SampleQuerySet][google.cloud.discoveryengine.v1beta.SampleQuerySet].
	//
	//  Only populated when the evaluation's state is SUCCEEDED.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Evaluation.quality_metrics
	QualityMetrics *QualityMetrics `json:"qualityMetrics,omitempty"`

	// Output only. The state of the evaluation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Evaluation.state
	State *string `json:"state,omitempty"`

	// Output only. The error that occurred during evaluation. Only populated when
	//  the evaluation's state is FAILED.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Evaluation.error
	Error *Status `json:"error,omitempty"`

	// Output only. Timestamp the
	//  [Evaluation][google.cloud.discoveryengine.v1beta.Evaluation] was created
	//  at.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Evaluation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp the
	//  [Evaluation][google.cloud.discoveryengine.v1beta.Evaluation] was completed
	//  at.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Evaluation.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. A sample of errors encountered while processing the request.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Evaluation.error_samples
	ErrorSamples []Status `json:"errorSamples,omitempty"`
}
