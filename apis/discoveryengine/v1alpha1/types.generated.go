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


// +kcc:proto=google.cloud.discoveryengine.v1beta.EmbeddingConfig
type EmbeddingConfig struct {
	// Full field path in the schema mapped as embedding field.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.EmbeddingConfig.field_path
	FieldPath *string `json:"fieldPath,omitempty"`
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

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchRequest.PersonalizationSpec
type SearchRequest_PersonalizationSpec struct {
	// The personalization mode of the search request.
	//  Defaults to
	//  [Mode.AUTO][google.cloud.discoveryengine.v1beta.SearchRequest.PersonalizationSpec.Mode.AUTO].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchRequest.PersonalizationSpec.mode
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.ServingConfig
type ServingConfig struct {
	// The MediaConfig of the serving configuration.
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

// +kcc:proto=google.cloud.discoveryengine.v1beta.ServingConfig.GenericConfig
type ServingConfig_GenericConfig struct {
	// Specifies the expected behavior of content search.
	//  Only valid for content-search enabled data store.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.GenericConfig.content_search_spec
	ContentSearchSpec *SearchRequest_ContentSearchSpec `json:"contentSearchSpec,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.ServingConfig.MediaConfig
type ServingConfig_MediaConfig struct {
	// Specifies the content watched percentage threshold for demotion.
	//  Threshold value must be between [0, 1.0] inclusive.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.MediaConfig.content_watched_percentage_threshold
	ContentWatchedPercentageThreshold *float32 `json:"contentWatchedPercentageThreshold,omitempty"`

	// Specifies the content watched minutes threshold for demotion.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.MediaConfig.content_watched_seconds_threshold
	ContentWatchedSecondsThreshold *float32 `json:"contentWatchedSecondsThreshold,omitempty"`

	// Specifies the event type used for demoting recommendation result.
	//  Currently supported values:
	//
	//  * `view-item`: Item viewed.
	//  * `media-play`: Start/resume watching a video, playing a song, etc.
	//  * `media-complete`: Finished or stopped midway through a video, song,
	//  etc.
	//
	//  If unset, watch history demotion will not be applied. Content freshness
	//  demotion will still be applied.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.MediaConfig.demotion_event_type
	DemotionEventType *string `json:"demotionEventType,omitempty"`

	// Optional. Specifies the number of days to look back for demoting watched
	//  content. If set to zero or unset, defaults to the maximum of 365 days.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.MediaConfig.demote_content_watched_past_days
	DemoteContentWatchedPastDays *int32 `json:"demoteContentWatchedPastDays,omitempty"`

	// Specifies the content freshness used for recommendation result.
	//  Contents will be demoted if contents were published for more than content
	//  freshness cutoff days.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.MediaConfig.content_freshness_cutoff_days
	ContentFreshnessCutoffDays *int32 `json:"contentFreshnessCutoffDays,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.ServingConfig
type ServingConfigObservedState struct {
	// Output only. ServingConfig created timestamp.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. ServingConfig updated timestamp.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ServingConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
