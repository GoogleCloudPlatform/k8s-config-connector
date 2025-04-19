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

// +generated:types
// krm.group: discoveryengine.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.discoveryengine.v1alpha
// resource: DiscoveryEngineDataStore:DataStore
// resource: DiscoveryEngineEngine:Engine
// resource: DiscoveryEngineTargetSite:TargetSite

package v1alpha1

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ChunkingConfig
type DocumentProcessingConfig_ChunkingConfig struct {
	// Configuration for the layout based chunking.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ChunkingConfig.layout_based_chunking_config
	LayoutBasedChunkingConfig *DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig `json:"layoutBasedChunkingConfig,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ChunkingConfig.LayoutBasedChunkingConfig
type DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig struct {
	// The token size limit for each chunk.
	//
	//  Supported values: 100-500 (inclusive).
	//  Default value: 500.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ChunkingConfig.LayoutBasedChunkingConfig.chunk_size
	ChunkSize *int32 `json:"chunkSize,omitempty"`

	// Whether to include appending different levels of headings to chunks
	//  from the middle of the document to prevent context loss.
	//
	//  Default value: False.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ChunkingConfig.LayoutBasedChunkingConfig.include_ancestor_headings
	IncludeAncestorHeadings *bool `json:"includeAncestorHeadings,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ParsingConfig
type DocumentProcessingConfig_ParsingConfig struct {
	// Configurations applied to digital parser.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ParsingConfig.digital_parsing_config
	DigitalParsingConfig *DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig `json:"digitalParsingConfig,omitempty"`

	// Configurations applied to OCR parser. Currently it only applies to
	//  PDFs.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ParsingConfig.ocr_parsing_config
	OcrParsingConfig *DocumentProcessingConfig_ParsingConfig_OcrParsingConfig `json:"ocrParsingConfig,omitempty"`

	// Configurations applied to layout parser.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ParsingConfig.layout_parsing_config
	LayoutParsingConfig *DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig `json:"layoutParsingConfig,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ParsingConfig.DigitalParsingConfig
type DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig struct {
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ParsingConfig.LayoutParsingConfig
type DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig struct {
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ParsingConfig.OcrParsingConfig
type DocumentProcessingConfig_ParsingConfig_OcrParsingConfig struct {
	// [DEPRECATED] This field is deprecated. To use the additional enhanced
	//  document elements processing, please switch to `layout_parsing_config`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ParsingConfig.OcrParsingConfig.enhanced_document_elements
	EnhancedDocumentElements []string `json:"enhancedDocumentElements,omitempty"`

	// If true, will use native text instead of OCR text on pages containing
	//  native text.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.ParsingConfig.OcrParsingConfig.use_native_text
	UseNativeText *bool `json:"useNativeText,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Engine
type Engine struct {
	// Additional config specs for a `similar-items` engine.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.similar_documents_config
	SimilarDocumentsConfig *Engine_SimilarDocumentsEngineConfig `json:"similarDocumentsConfig,omitempty"`

	// Configurations for the Chat Engine. Only applicable if
	//  [solution_type][google.cloud.discoveryengine.v1alpha.Engine.solution_type]
	//  is
	//  [SOLUTION_TYPE_CHAT][google.cloud.discoveryengine.v1alpha.SolutionType.SOLUTION_TYPE_CHAT].
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.chat_engine_config
	ChatEngineConfig *Engine_ChatEngineConfig `json:"chatEngineConfig,omitempty"`

	// Configurations for the Search Engine. Only applicable if
	//  [solution_type][google.cloud.discoveryengine.v1alpha.Engine.solution_type]
	//  is
	//  [SOLUTION_TYPE_SEARCH][google.cloud.discoveryengine.v1alpha.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.search_engine_config
	SearchEngineConfig *Engine_SearchEngineConfig `json:"searchEngineConfig,omitempty"`

	// Configurations for the Media Engine. Only applicable on the data
	//  stores with
	//  [solution_type][google.cloud.discoveryengine.v1alpha.Engine.solution_type]
	//  [SOLUTION_TYPE_RECOMMENDATION][google.cloud.discoveryengine.v1alpha.SolutionType.SOLUTION_TYPE_RECOMMENDATION]
	//  and
	//  [IndustryVertical.MEDIA][google.cloud.discoveryengine.v1alpha.IndustryVertical.MEDIA]
	//  vertical.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.media_recommendation_engine_config
	MediaRecommendationEngineConfig *Engine_MediaRecommendationEngineConfig `json:"mediaRecommendationEngineConfig,omitempty"`

	// Immutable. The fully qualified resource name of the engine.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 1024
	//  characters.
	//
	//  Format:
	//  `projects/{project_number}/locations/{location}/collections/{collection}/engines/{engine}`
	//  engine should be 1-63 characters, and valid characters are
	//  /[a-z0-9][a-z0-9-_]*/. Otherwise, an INVALID_ARGUMENT error is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.name
	Name *string `json:"name,omitempty"`

	// Required. The display name of the engine. Should be human readable. UTF-8
	//  encoded string with limit of 1024 characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The data stores associated with this engine.
	//
	//  For
	//  [SOLUTION_TYPE_SEARCH][google.cloud.discoveryengine.v1alpha.SolutionType.SOLUTION_TYPE_SEARCH]
	//  and
	//  [SOLUTION_TYPE_RECOMMENDATION][google.cloud.discoveryengine.v1alpha.SolutionType.SOLUTION_TYPE_RECOMMENDATION]
	//  type of engines, they can only associate with at most one data store.
	//
	//  If
	//  [solution_type][google.cloud.discoveryengine.v1alpha.Engine.solution_type]
	//  is
	//  [SOLUTION_TYPE_CHAT][google.cloud.discoveryengine.v1alpha.SolutionType.SOLUTION_TYPE_CHAT],
	//  multiple [DataStore][google.cloud.discoveryengine.v1alpha.DataStore]s in
	//  the same [Collection][google.cloud.discoveryengine.v1alpha.Collection] can
	//  be associated here.
	//
	//  Note that when used in
	//  [CreateEngineRequest][google.cloud.discoveryengine.v1alpha.CreateEngineRequest],
	//  one DataStore id must be provided as the system will use it for necessary
	//  initializations.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.data_store_ids
	DataStoreIds []string `json:"dataStoreIds,omitempty"`

	// Required. The solutions of the engine.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.solution_type
	SolutionType *string `json:"solutionType,omitempty"`

	// The industry vertical that the engine registers.
	//  The restriction of the Engine industry vertical is based on
	//  [DataStore][google.cloud.discoveryengine.v1alpha.DataStore]: If
	//  unspecified, default to `GENERIC`. Vertical on Engine has to match vertical
	//  of the DataStore linked to the engine.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.industry_vertical
	IndustryVertical *string `json:"industryVertical,omitempty"`

	// Common config spec that specifies the metadata of the engine.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.common_config
	CommonConfig *Engine_CommonConfig `json:"commonConfig,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Engine.ChatEngineConfig
type Engine_ChatEngineConfig struct {
	// The configurationt generate the Dialogflow agent that is associated to
	//  this Engine.
	//
	//  Note that these configurations are one-time consumed by
	//  and passed to Dialogflow service. It means they cannot be retrieved using
	//  [EngineService.GetEngine][google.cloud.discoveryengine.v1alpha.EngineService.GetEngine]
	//  or
	//  [EngineService.ListEngines][google.cloud.discoveryengine.v1alpha.EngineService.ListEngines]
	//  API after engine creation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.ChatEngineConfig.agent_creation_config
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
	//  [EngineService.GetEngine][google.cloud.discoveryengine.v1alpha.EngineService.GetEngine]
	//  or
	//  [EngineService.ListEngines][google.cloud.discoveryengine.v1alpha.EngineService.ListEngines]
	//  API after engine creation. Use
	//  [ChatEngineMetadata.dialogflow_agent][google.cloud.discoveryengine.v1alpha.Engine.ChatEngineMetadata.dialogflow_agent]
	//  for actual agent association after Engine is created.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.ChatEngineConfig.dialogflow_agent_to_link
	DialogflowAgentToLink *string `json:"dialogflowAgentToLink,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Engine.ChatEngineConfig.AgentCreationConfig
type Engine_ChatEngineConfig_AgentCreationConfig struct {
	// Name of the company, organization or other entity that the agent
	//  represents. Used for knowledge connector LLM prompt and for knowledge
	//  search.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.ChatEngineConfig.AgentCreationConfig.business
	Business *string `json:"business,omitempty"`

	// Required. The default language of the agent as a language tag.
	//  See [Language
	//  Support](https://cloud.google.com/dialogflow/docs/reference/language)
	//  for a list of the currently supported language codes.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.ChatEngineConfig.AgentCreationConfig.default_language_code
	DefaultLanguageCode *string `json:"defaultLanguageCode,omitempty"`

	// Required. The time zone of the agent from the [time zone
	//  database](https://www.iana.org/time-zones), e.g., America/New_York,
	//  Europe/Paris.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.ChatEngineConfig.AgentCreationConfig.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Agent location for Agent creation, supported values: global/us/eu.
	//  If not provided, us Engine will create Agent using us-central-1 by
	//  default; eu Engine will create Agent using eu-west-1 by default.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.ChatEngineConfig.AgentCreationConfig.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Engine.ChatEngineMetadata
type Engine_ChatEngineMetadata struct {
	// The resource name of a Dialogflow agent, that this Chat Engine refers
	//  to.
	//
	//  Format: `projects/<Project ID>/locations/<Location ID>/agents/<Agent
	//  ID>`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.ChatEngineMetadata.dialogflow_agent
	DialogflowAgent *string `json:"dialogflowAgent,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Engine.CommonConfig
type Engine_CommonConfig struct {
	// The name of the company, business or entity that is associated with the
	//  engine. Setting this may help improve LLM related features.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.CommonConfig.company_name
	CompanyName *string `json:"companyName,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Engine.MediaRecommendationEngineConfig
type Engine_MediaRecommendationEngineConfig struct {
	// Required. The type of engine. e.g., `recommended-for-you`.
	//
	//  This field together with
	//  [optimization_objective][Engine.optimization_objective] describe engine
	//  metadata to use to control engine training and serving.
	//
	//  Currently supported values: `recommended-for-you`, `others-you-may-like`,
	//  `more-like-this`, `most-popular-items`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.MediaRecommendationEngineConfig.type
	Type *string `json:"type,omitempty"`

	// The optimization objective. e.g., `cvr`.
	//
	//  This field together with
	//  [optimization_objective][google.cloud.discoveryengine.v1alpha.Engine.MediaRecommendationEngineConfig.type]
	//  describe engine metadata to use to control engine training and serving.
	//
	//  Currently supported
	//  values: `ctr`, `cvr`.
	//
	//   If not specified, we choose default based on engine type.
	//  Default depends on type of recommendation:
	//
	//  `recommended-for-you` => `ctr`
	//
	//  `others-you-may-like` => `ctr`
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.MediaRecommendationEngineConfig.optimization_objective
	OptimizationObjective *string `json:"optimizationObjective,omitempty"`

	// Name and value of the custom threshold for cvr optimization_objective.
	//  For target_field `watch-time`, target_field_value must be an integer
	//  value indicating the media progress time in seconds between (0, 86400]
	//  (excludes 0, includes 86400) (e.g., 90).
	//  For target_field `watch-percentage`, the target_field_value must be a
	//  valid float value between (0, 1.0] (excludes 0, includes 1.0) (e.g.,
	//  0.5).
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.MediaRecommendationEngineConfig.optimization_objective_config
	OptimizationObjectiveConfig *Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig `json:"optimizationObjectiveConfig,omitempty"`

	// The training state that the engine is in (e.g.
	//  `TRAINING` or `PAUSED`).
	//
	//  Since part of the cost of running the service
	//  is frequency of training - this can be used to determine when to train
	//  engine in order to control cost. If not specified: the default value for
	//  `CreateEngine` method is `TRAINING`. The default value for
	//  `UpdateEngine` method is to keep the state the same as before.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.MediaRecommendationEngineConfig.training_state
	TrainingState *string `json:"trainingState,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Engine.MediaRecommendationEngineConfig.OptimizationObjectiveConfig
type Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig struct {
	// Required. The name of the field to target. Currently supported
	//  values: `watch-percentage`, `watch-time`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.MediaRecommendationEngineConfig.OptimizationObjectiveConfig.target_field
	TargetField *string `json:"targetField,omitempty"`

	// Required. The threshold to be applied to the target (e.g., 0.5).
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.MediaRecommendationEngineConfig.OptimizationObjectiveConfig.target_field_value_float
	TargetFieldValueFloat *float32 `json:"targetFieldValueFloat,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Engine.RecommendationMetadata
type Engine_RecommendationMetadata struct {
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Engine.SearchEngineConfig
type Engine_SearchEngineConfig struct {
	// The search feature tier of this engine.
	//
	//  Different tiers might have different
	//  pricing. To learn more, check the pricing documentation.
	//
	//  Defaults to
	//  [SearchTier.SEARCH_TIER_STANDARD][google.cloud.discoveryengine.v1alpha.SearchTier.SEARCH_TIER_STANDARD]
	//  if not specified.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.SearchEngineConfig.search_tier
	SearchTier *string `json:"searchTier,omitempty"`

	// The add-on that this search engine enables.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.SearchEngineConfig.search_add_ons
	SearchAddOns []string `json:"searchAddOns,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Engine.SimilarDocumentsEngineConfig
type Engine_SimilarDocumentsEngineConfig struct {
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.FieldConfig
type FieldConfig struct {
	// Required. Field path of the schema field.
	//  For example: `title`, `description`, `release_info.release_year`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.FieldConfig.field_path
	FieldPath *string `json:"fieldPath,omitempty"`

	// If
	//  [indexable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.indexable_option]
	//  is
	//  [INDEXABLE_ENABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.IndexableOption.INDEXABLE_ENABLED],
	//  field values are indexed so that it can be filtered or faceted in
	//  [SearchService.Search][google.cloud.discoveryengine.v1alpha.SearchService.Search].
	//
	//  If
	//  [indexable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.indexable_option]
	//  is unset, the server behavior defaults to
	//  [INDEXABLE_DISABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.IndexableOption.INDEXABLE_DISABLED]
	//  for fields that support setting indexable options. For those fields that do
	//  not support setting indexable options, such as `object` and `boolean` and
	//  key properties, the server will skip
	//  [indexable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.indexable_option]
	//  setting, and setting
	//  [indexable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.indexable_option]
	//  for those fields will throw `INVALID_ARGUMENT` error.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.FieldConfig.indexable_option
	IndexableOption *string `json:"indexableOption,omitempty"`

	// If
	//  [dynamic_facetable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.dynamic_facetable_option]
	//  is
	//  [DYNAMIC_FACETABLE_ENABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.DynamicFacetableOption.DYNAMIC_FACETABLE_ENABLED],
	//  field values are available for dynamic facet. Could only be
	//  [DYNAMIC_FACETABLE_DISABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.DynamicFacetableOption.DYNAMIC_FACETABLE_DISABLED]
	//  if
	//  [FieldConfig.indexable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.indexable_option]
	//  is
	//  [INDEXABLE_DISABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.IndexableOption.INDEXABLE_DISABLED].
	//  Otherwise, an `INVALID_ARGUMENT` error will be returned.
	//
	//  If
	//  [dynamic_facetable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.dynamic_facetable_option]
	//  is unset, the server behavior defaults to
	//  [DYNAMIC_FACETABLE_DISABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.DynamicFacetableOption.DYNAMIC_FACETABLE_DISABLED]
	//  for fields that support setting dynamic facetable options. For those fields
	//  that do not support setting dynamic facetable options, such as `object` and
	//  `boolean`, the server will skip dynamic facetable option setting, and
	//  setting
	//  [dynamic_facetable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.dynamic_facetable_option]
	//  for those fields will throw `INVALID_ARGUMENT` error.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.FieldConfig.dynamic_facetable_option
	DynamicFacetableOption *string `json:"dynamicFacetableOption,omitempty"`

	// If
	//  [searchable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.searchable_option]
	//  is
	//  [SEARCHABLE_ENABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.SearchableOption.SEARCHABLE_ENABLED],
	//  field values are searchable by text queries in
	//  [SearchService.Search][google.cloud.discoveryengine.v1alpha.SearchService.Search].
	//
	//  If
	//  [SEARCHABLE_ENABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.SearchableOption.SEARCHABLE_ENABLED]
	//  but field type is numerical, field values will not be searchable by text
	//  queries in
	//  [SearchService.Search][google.cloud.discoveryengine.v1alpha.SearchService.Search],
	//  as there are no text values associated to numerical fields.
	//
	//  If
	//  [searchable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.searchable_option]
	//  is unset, the server behavior defaults to
	//  [SEARCHABLE_DISABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.SearchableOption.SEARCHABLE_DISABLED]
	//  for fields that support setting searchable options. Only `string` fields
	//  that have no key property mapping support setting
	//  [searchable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.searchable_option].
	//
	//  For those fields that do not support setting searchable options, the server
	//  will skip searchable option setting, and setting
	//  [searchable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.searchable_option]
	//  for those fields will throw `INVALID_ARGUMENT` error.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.FieldConfig.searchable_option
	SearchableOption *string `json:"searchableOption,omitempty"`

	// If
	//  [retrievable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.retrievable_option]
	//  is
	//  [RETRIEVABLE_ENABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.RetrievableOption.RETRIEVABLE_ENABLED],
	//  field values are included in the search results.
	//
	//  If
	//  [retrievable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.retrievable_option]
	//  is unset, the server behavior defaults to
	//  [RETRIEVABLE_DISABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.RetrievableOption.RETRIEVABLE_DISABLED]
	//  for fields that support setting retrievable options. For those fields
	//  that do not support setting retrievable options, such as `object` and
	//  `boolean`, the server will skip retrievable option setting, and setting
	//  [retrievable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.retrievable_option]
	//  for those fields will throw `INVALID_ARGUMENT` error.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.FieldConfig.retrievable_option
	RetrievableOption *string `json:"retrievableOption,omitempty"`

	// If
	//  [completable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.completable_option]
	//  is
	//  [COMPLETABLE_ENABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.CompletableOption.COMPLETABLE_ENABLED],
	//  field values are directly used and returned as suggestions for Autocomplete
	//  in
	//  [CompletionService.CompleteQuery][google.cloud.discoveryengine.v1alpha.CompletionService.CompleteQuery].
	//
	//  If
	//  [completable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.completable_option]
	//  is unset, the server behavior defaults to
	//  [COMPLETABLE_DISABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.CompletableOption.COMPLETABLE_DISABLED]
	//  for fields that support setting completable options, which are just
	//  `string` fields. For those fields that do not support setting completable
	//  options, the server will skip completable option setting, and setting
	//  [completable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.completable_option]
	//  for those fields will throw `INVALID_ARGUMENT` error.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.FieldConfig.completable_option
	CompletableOption *string `json:"completableOption,omitempty"`

	// If
	//  [recs_filterable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.recs_filterable_option]
	//  is
	//  [FILTERABLE_ENABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.FilterableOption.FILTERABLE_ENABLED],
	//  field values are filterable by filter expression in
	//  [RecommendationService.Recommend][google.cloud.discoveryengine.v1alpha.RecommendationService.Recommend].
	//
	//  If
	//  [FILTERABLE_ENABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.FilterableOption.FILTERABLE_ENABLED]
	//  but the field type is numerical, field values are not filterable by text
	//  queries in
	//  [RecommendationService.Recommend][google.cloud.discoveryengine.v1alpha.RecommendationService.Recommend].
	//  Only textual fields are supported.
	//
	//  If
	//  [recs_filterable_option][google.cloud.discoveryengine.v1alpha.FieldConfig.recs_filterable_option]
	//  is unset, the default setting is
	//  [FILTERABLE_DISABLED][google.cloud.discoveryengine.v1alpha.FieldConfig.FilterableOption.FILTERABLE_DISABLED]
	//  for fields that support setting filterable options.
	//
	//  When a field set to [FILTERABLE_DISABLED] is filtered, a warning is
	//  generated and an empty result is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.FieldConfig.recs_filterable_option
	RecsFilterableOption *string `json:"recsFilterableOption,omitempty"`

	// If this field is set, only the corresponding source will be indexed for
	//  this field. Otherwise, the values from different sources are merged.
	//
	//  Assuming a page with `<author, a>` in meta tag, and `<author, b>` in page
	//  map:
	//   if this enum is set to METATAGS, we will only index `<author, a>`;
	//   if this enum is not set, we will merge them and index `<author, [a, b]>`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.FieldConfig.advanced_site_search_data_sources
	AdvancedSiteSearchDataSources []string `json:"advancedSiteSearchDataSources,omitempty"`

	// Field paths for indexing custom attribute from schema.org data. More
	//  details of schema.org and its defined types can be found at
	//  [schema.org](https://schema.org).
	//
	//  It is only used on advanced site search schema.
	//
	//  Currently only support full path from root. The full path to a field is
	//  constructed by concatenating field names, starting from `_root`, with
	//  a period `.` as the delimiter. Examples:
	//
	//  * Publish date of the root: _root.datePublished
	//  * Publish date of the reviews: _root.review.datePublished
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.FieldConfig.schema_org_paths
	SchemaOrgPaths []string `json:"schemaOrgPaths,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.IdpConfig.ExternalIdpConfig
type IdpConfig_ExternalIdpConfig struct {
	// Workforce pool name.
	//  Example: "locations/global/workforcePools/pool_id"
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.IdpConfig.ExternalIdpConfig.workforce_pool_name
	WorkforcePoolName *string `json:"workforcePoolName,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Schema
type Schema struct {
	// The structured representation of the schema.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Schema.struct_schema
	StructSchema map[string]string `json:"structSchema,omitempty"`

	// The JSON representation of the schema.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Schema.json_schema
	JsonSchema *string `json:"jsonSchema,omitempty"`

	// Immutable. The full resource name of the schema, in the format of
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/schemas/{schema}`.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 1024
	//  characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Schema.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.SiteVerificationInfo
type SiteVerificationInfo struct {
	// Site verification state indicating the ownership and validity.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.SiteVerificationInfo.site_verification_state
	SiteVerificationState *string `json:"siteVerificationState,omitempty"`

	// Latest site verification time.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.SiteVerificationInfo.verify_time
	VerifyTime *string `json:"verifyTime,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.TargetSite
type TargetSite struct {

	// Required. Input only. The user provided URI pattern from which the
	//  `generated_uri_pattern` is generated.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.TargetSite.provided_uri_pattern
	ProvidedURIPattern *string `json:"providedURIPattern,omitempty"`

	// The type of the target site, e.g., whether the site is to be included or
	//  excluded.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.TargetSite.type
	Type *string `json:"type,omitempty"`

	// Input only. If set to false, a uri_pattern is generated to include all
	//  pages whose address contains the provided_uri_pattern. If set to true, an
	//  uri_pattern is generated to try to be an exact match of the
	//  provided_uri_pattern or just the specific page if the provided_uri_pattern
	//  is a specific one. provided_uri_pattern is always normalized to
	//  generate the URI pattern to be used by the search engine.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.TargetSite.exact_match
	ExactMatch *bool `json:"exactMatch,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.TargetSite.FailureReason
type TargetSite_FailureReason struct {
	// Failed due to insufficient quota.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.TargetSite.FailureReason.quota_failure
	QuotaFailure *TargetSite_FailureReason_QuotaFailure `json:"quotaFailure,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.TargetSite.FailureReason.QuotaFailure
type TargetSite_FailureReason_QuotaFailure struct {
	// This number is an estimation on how much total quota this project needs
	//  to successfully complete indexing.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.TargetSite.FailureReason.QuotaFailure.total_required_quota
	TotalRequiredQuota *int64 `json:"totalRequiredQuota,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Engine
type EngineObservedState struct {
	// Output only. Additional information of a recommendation engine. Only
	//  applicable if
	//  [solution_type][google.cloud.discoveryengine.v1alpha.Engine.solution_type]
	//  is
	//  [SOLUTION_TYPE_RECOMMENDATION][google.cloud.discoveryengine.v1alpha.SolutionType.SOLUTION_TYPE_RECOMMENDATION].
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.recommendation_metadata
	RecommendationMetadata *Engine_RecommendationMetadata `json:"recommendationMetadata,omitempty"`

	// Output only. Additional information of the Chat Engine. Only applicable
	//  if
	//  [solution_type][google.cloud.discoveryengine.v1alpha.Engine.solution_type]
	//  is
	//  [SOLUTION_TYPE_CHAT][google.cloud.discoveryengine.v1alpha.SolutionType.SOLUTION_TYPE_CHAT].
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.chat_engine_metadata
	ChatEngineMetadata *Engine_ChatEngineMetadata `json:"chatEngineMetadata,omitempty"`

	// Output only. Timestamp the Recommendation Engine was created at.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp the Recommendation Engine was last updated.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Engine.RecommendationMetadata
type Engine_RecommendationMetadataObservedState struct {
	// Output only. The serving state of the engine: `ACTIVE`, `NOT_ACTIVE`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.RecommendationMetadata.serving_state
	ServingState *string `json:"servingState,omitempty"`

	// Output only. The state of data requirements for this engine: `DATA_OK`
	//  and `DATA_ERROR`.
	//
	//  Engine cannot be trained if the data is in
	//  `DATA_ERROR` state. Engine can have `DATA_ERROR` state even
	//  if serving state is `ACTIVE`: engines were trained successfully before,
	//  but cannot be refreshed because the underlying engine no longer has
	//  sufficient data for training.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.RecommendationMetadata.data_state
	DataState *string `json:"dataState,omitempty"`

	// Output only. The timestamp when the latest successful tune finished. Only
	//  applicable on Media Recommendation engines.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.RecommendationMetadata.last_tune_time
	LastTuneTime *string `json:"lastTuneTime,omitempty"`

	// Output only. The latest tune operation id associated with the engine.
	//  Only applicable on Media Recommendation engines.
	//
	//  If present, this operation id can be used to determine if there is an
	//  ongoing tune for this engine. To check the operation status, send the
	//  GetOperation request with this operation id in the engine resource
	//  format. If no tuning has happened for this engine, the string is empty.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Engine.RecommendationMetadata.tuning_operation
	TuningOperation *string `json:"tuningOperation,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.FieldConfig
type FieldConfigObservedState struct {
	// Output only. Raw type of the field.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.FieldConfig.field_type
	FieldType *string `json:"fieldType,omitempty"`

	// Output only. Type of the key property that this field is mapped to. Empty
	//  string if this is not annotated as mapped to a key property.
	//
	//  Example types are `title`, `description`. Full list is defined
	//  by `keyPropertyMapping` in the schema field annotation.
	//
	//  If the schema field has a `KeyPropertyMapping` annotation,
	//  `indexable_option` and `searchable_option` of this field cannot be
	//  modified.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.FieldConfig.key_property_type
	KeyPropertyType *string `json:"keyPropertyType,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.LanguageInfo
type LanguageInfoObservedState struct {
	// Output only. This is the normalized form of language_code.
	//  E.g.: language_code of `en-GB`, `en_GB`, `en-UK` or `en-gb`
	//  will have normalized_language_code of `en-GB`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.LanguageInfo.normalized_language_code
	NormalizedLanguageCode *string `json:"normalizedLanguageCode,omitempty"`

	// Output only. Language part of normalized_language_code.
	//  E.g.: `en-US` -> `en`, `zh-Hans-HK` -> `zh`, `en` -> `en`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.LanguageInfo.language
	Language *string `json:"language,omitempty"`

	// Output only. Region part of normalized_language_code, if present.
	//  E.g.: `en-US` -> `US`, `zh-Hans-HK` -> `HK`, `en` -> ``.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.LanguageInfo.region
	Region *string `json:"region,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.Schema
type SchemaObservedState struct {
	// Output only. Configurations for fields of the schema.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.Schema.field_configs
	FieldConfigs []FieldConfig `json:"fieldConfigs,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.TargetSite
type TargetSiteObservedState struct {
	// Output only. The fully qualified resource name of the target site.
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/siteSearchEngine/targetSites/{target_site}`
	//  The `target_site_id` is system-generated.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.TargetSite.name
	Name *string `json:"name,omitempty"`

	// Output only. This is system-generated based on the provided_uri_pattern.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.TargetSite.generated_uri_pattern
	GeneratedURIPattern *string `json:"generatedURIPattern,omitempty"`

	// Output only. Root domain of the provided_uri_pattern.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.TargetSite.root_domain_uri
	RootDomainURI *string `json:"rootDomainURI,omitempty"`

	// Output only. Site ownership and validity verification status.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.TargetSite.site_verification_info
	SiteVerificationInfo *SiteVerificationInfo `json:"siteVerificationInfo,omitempty"`

	// Output only. Indexing status.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.TargetSite.indexing_status
	IndexingStatus *string `json:"indexingStatus,omitempty"`

	// Output only. The target site's last updated time.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.TargetSite.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Failure reason.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.TargetSite.failure_reason
	FailureReason *TargetSite_FailureReason `json:"failureReason,omitempty"`
}
