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

// +kcc:proto=google.cloud.discoveryengine.v1.DataStore.BillingEstimation
type DataStore_BillingEstimation struct {
	// Data size for structured data in terms of bytes.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataStore.BillingEstimation.structured_data_size
	StructuredDataSize *int64 `json:"structuredDataSize,omitempty"`

	// Data size for unstructured data in terms of bytes.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataStore.BillingEstimation.unstructured_data_size
	UnstructuredDataSize *int64 `json:"unstructuredDataSize,omitempty"`

	// Data size for websites in terms of bytes.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataStore.BillingEstimation.website_data_size
	WebsiteDataSize *int64 `json:"websiteDataSize,omitempty"`

	// Last updated timestamp for structured data.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataStore.BillingEstimation.structured_data_update_time
	StructuredDataUpdateTime *string `json:"structuredDataUpdateTime,omitempty"`

	// Last updated timestamp for unstructured data.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataStore.BillingEstimation.unstructured_data_update_time
	UnstructuredDataUpdateTime *string `json:"unstructuredDataUpdateTime,omitempty"`

	// Last updated timestamp for websites.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataStore.BillingEstimation.website_data_update_time
	WebsiteDataUpdateTime *string `json:"websiteDataUpdateTime,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.DocumentProcessingConfig
type DocumentProcessingConfig struct {
	// The full resource name of the Document Processing Config.
	//  Format:
	//  `projects/*/locations/*/collections/*/dataStores/*/documentProcessingConfig`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.name
	Name *string `json:"name,omitempty"`

	// Whether chunking mode is enabled.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.chunking_config
	ChunkingConfig *DocumentProcessingConfig_ChunkingConfig `json:"chunkingConfig,omitempty"`

	// Configurations for default Document parser.
	//  If not specified, we will configure it as default DigitalParsingConfig, and
	//  the default parsing config will be applied to all file types for Document
	//  parsing.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.default_parsing_config
	DefaultParsingConfig *DocumentProcessingConfig_ParsingConfig `json:"defaultParsingConfig,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ChunkingConfig
type DocumentProcessingConfig_ChunkingConfig struct {
	// Configuration for the layout based chunking.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ChunkingConfig.layout_based_chunking_config
	LayoutBasedChunkingConfig *DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig `json:"layoutBasedChunkingConfig,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ChunkingConfig.LayoutBasedChunkingConfig
type DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig struct {
	// The token size limit for each chunk.
	//
	//  Supported values: 100-500 (inclusive).
	//  Default value: 500.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ChunkingConfig.LayoutBasedChunkingConfig.chunk_size
	ChunkSize *int32 `json:"chunkSize,omitempty"`

	// Whether to include appending different levels of headings to chunks
	//  from the middle of the document to prevent context loss.
	//
	//  Default value: False.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ChunkingConfig.LayoutBasedChunkingConfig.include_ancestor_headings
	IncludeAncestorHeadings *bool `json:"includeAncestorHeadings,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig
type DocumentProcessingConfig_ParsingConfig struct {
	// Configurations applied to digital parser.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.digital_parsing_config
	DigitalParsingConfig *DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig `json:"digitalParsingConfig,omitempty"`

	// Configurations applied to OCR parser. Currently it only applies to
	//  PDFs.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.ocr_parsing_config
	OcrParsingConfig *DocumentProcessingConfig_ParsingConfig_OcrParsingConfig `json:"ocrParsingConfig,omitempty"`

	// Configurations applied to layout parser.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.layout_parsing_config
	LayoutParsingConfig *DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig `json:"layoutParsingConfig,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.DigitalParsingConfig
type DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig struct {
}

// +kcc:proto=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.LayoutParsingConfig
type DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig struct {
}

// +kcc:proto=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.OcrParsingConfig
type DocumentProcessingConfig_ParsingConfig_OcrParsingConfig struct {
	// [DEPRECATED] This field is deprecated. To use the additional enhanced
	//  document elements processing, please switch to `layout_parsing_config`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.OcrParsingConfig.enhanced_document_elements
	EnhancedDocumentElements []string `json:"enhancedDocumentElements,omitempty"`

	// If true, will use native text instead of OCR text on pages containing
	//  native text.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.OcrParsingConfig.use_native_text
	UseNativeText *bool `json:"useNativeText,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig
type Engine_ChatEngineConfig struct {
	// The configurationt generate the Dialogflow agent that is associated to
	//  this Engine.
	//
	//  Note that these configurations are one-time consumed by
	//  and passed to Dialogflow service. It means they cannot be retrieved using
	//  [EngineService.GetEngine][google.cloud.discoveryengine.v1.EngineService.GetEngine]
	//  or
	//  [EngineService.ListEngines][google.cloud.discoveryengine.v1.EngineService.ListEngines]
	//  API after engine creation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.agent_creation_config
	AgentCreationConfig *Engine_ChatEngineConfig_AgentCreationConfig `json:"agentCreationConfig,omitempty"`

	// The resource name of an exist Dialogflow agent to link to this Chat
	//  Engine. Customers can either provide `agent_creation_config` to create
	//  agent or provide an agent name that links the agent with the Chat engine.
	//
	//  Format: `projects/<Project ID>/locations/<Location ID>/agents/<Agent
	//  ID>`.
	//
	//  Note that the `dialogflow_agent_to_link` are one-time consumed by and
	//  passed to Dialogflow service. It means they cannot be retrieved using
	//  [EngineService.GetEngine][google.cloud.discoveryengine.v1.EngineService.GetEngine]
	//  or
	//  [EngineService.ListEngines][google.cloud.discoveryengine.v1.EngineService.ListEngines]
	//  API after engine creation. Use
	//  [ChatEngineMetadata.dialogflow_agent][google.cloud.discoveryengine.v1.Engine.ChatEngineMetadata.dialogflow_agent]
	//  for actual agent association after Engine is created.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.dialogflow_agent_to_link
	DialogflowAgentToLink *string `json:"dialogflowAgentToLink,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.AgentCreationConfig
type Engine_ChatEngineConfig_AgentCreationConfig struct {
	// Name of the company, organization or other entity that the agent
	//  represents. Used for knowledge connector LLM prompt and for knowledge
	//  search.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.AgentCreationConfig.business
	Business *string `json:"business,omitempty"`

	// Required. The default language of the agent as a language tag.
	//  See [Language
	//  Support](https://cloud.google.com/dialogflow/docs/reference/language)
	//  for a list of the currently supported language codes.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.AgentCreationConfig.default_language_code
	DefaultLanguageCode *string `json:"defaultLanguageCode,omitempty"`

	// Required. The time zone of the agent from the [time zone
	//  database](https://www.iana.org/time-zones), e.g., America/New_York,
	//  Europe/Paris.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.AgentCreationConfig.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Agent location for Agent creation, supported values: global/us/eu.
	//  If not provided, us Engine will create Agent using us-central-1 by
	//  default; eu Engine will create Agent using eu-west-1 by default.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.AgentCreationConfig.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.ChatEngineMetadata
type Engine_ChatEngineMetadata struct {
	// The resource name of a Dialogflow agent, that this Chat Engine refers
	//  to.
	//
	//  Format: `projects/<Project ID>/locations/<Location ID>/agents/<Agent
	//  ID>`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineMetadata.dialogflow_agent
	DialogflowAgent *string `json:"dialogflowAgent,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.CommonConfig
type Engine_CommonConfig struct {
	// The name of the company, business or entity that is associated with the
	//  engine. Setting this may help improve LLM related features.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.CommonConfig.company_name
	CompanyName *string `json:"companyName,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.SearchEngineConfig
type Engine_SearchEngineConfig struct {
	// The search feature tier of this engine.
	//
	//  Different tiers might have different
	//  pricing. To learn more, check the pricing documentation.
	//
	//  Defaults to
	//  [SearchTier.SEARCH_TIER_STANDARD][google.cloud.discoveryengine.v1.SearchTier.SEARCH_TIER_STANDARD]
	//  if not specified.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.SearchEngineConfig.search_tier
	SearchTier *string `json:"searchTier,omitempty"`

	// The add-on that this search engine enables.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.SearchEngineConfig.search_add_ons
	SearchAddOns []string `json:"searchAddOns,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Schema
type Schema struct {
	// The structured representation of the schema.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Schema.struct_schema
	StructSchema map[string]string `json:"structSchema,omitempty"`

	// The JSON representation of the schema.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Schema.json_schema
	JsonSchema *string `json:"jsonSchema,omitempty"`

	// Immutable. The full resource name of the schema, in the format of
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/schemas/{schema}`.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 1024
	//  characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Schema.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.SiteVerificationInfo
type SiteVerificationInfo struct {
	// Site verification state indicating the ownership and validity.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.SiteVerificationInfo.site_verification_state
	SiteVerificationState *string `json:"siteVerificationState,omitempty"`

	// Latest site verification time.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.SiteVerificationInfo.verify_time
	VerifyTime *string `json:"verifyTime,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.TargetSite.FailureReason
type TargetSite_FailureReason struct {
	// Failed due to insufficient quota.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.TargetSite.FailureReason.quota_failure
	QuotaFailure *TargetSite_FailureReason_QuotaFailure `json:"quotaFailure,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.TargetSite.FailureReason.QuotaFailure
type TargetSite_FailureReason_QuotaFailure struct {
	// This number is an estimation on how much total quota this project needs
	//  to successfully complete indexing.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.TargetSite.FailureReason.QuotaFailure.total_required_quota
	TotalRequiredQuota *int64 `json:"totalRequiredQuota,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.WorkspaceConfig
type WorkspaceConfig struct {
	// The Google Workspace data source.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WorkspaceConfig.type
	Type *string `json:"type,omitempty"`

	// Obfuscated Dasher customer ID.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WorkspaceConfig.dasher_customer_id
	DasherCustomerID *string `json:"dasherCustomerID,omitempty"`

	// Optional. The super admin service account for the workspace that will be
	//  used for access token generation. For now we only use it for Native Google
	//  Drive connector data ingestion.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WorkspaceConfig.super_admin_service_account
	SuperAdminServiceAccount *string `json:"superAdminServiceAccount,omitempty"`

	// Optional. The super admin email address for the workspace that will be used
	//  for access token generation. For now we only use it for Native Google Drive
	//  connector data ingestion.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.WorkspaceConfig.super_admin_email_address
	SuperAdminEmailAddress *string `json:"superAdminEmailAddress,omitempty"`
}

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
