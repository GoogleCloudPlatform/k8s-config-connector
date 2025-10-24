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
// proto.service: google.cloud.discoveryengine.v1
// resource: DiscoveryEngineDataStore:DataStore
// resource: DiscoveryEngineEngine:Engine
// resource: DiscoveryEngineTargetSite:TargetSite

package v1alpha1

// +kcc:proto=google.cloud.discoveryengine.v1.AdvancedSiteSearchConfig
type AdvancedSiteSearchConfig struct {
	// If set true, initial indexing is disabled for the DataStore.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.AdvancedSiteSearchConfig.disable_initial_index
	DisableInitialIndex *bool `json:"disableInitialIndex,omitempty"`

	// If set true, automatic refresh is disabled for the DataStore.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.AdvancedSiteSearchConfig.disable_automatic_refresh
	DisableAutomaticRefresh *bool `json:"disableAutomaticRefresh,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.CmekConfig
type CmekConfig struct {
	// Required. The name of the CmekConfig of the form
	//  `projects/{project}/locations/{location}/cmekConfig` or
	//  `projects/{project}/locations/{location}/cmekConfigs/{cmek_config}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.name
	Name *string `json:"name,omitempty"`

	// KMS key resource name which will be used to encrypt resources
	//  `projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{keyId}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`

	// KMS key version resource name which will be used to encrypt resources
	//  `<kms_key>/cryptoKeyVersions/{keyVersion}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.kms_key_version
	KMSKeyVersion *string `json:"kmsKeyVersion,omitempty"`

	// Optional. Single-regional CMEKs that are required for some VAIS features.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.single_region_keys
	SingleRegionKeys []SingleRegionKey `json:"singleRegionKeys,omitempty"`
}

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
	// Optional. If true, the LLM based annotation is added to the table
	//  during parsing.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.LayoutParsingConfig.enable_table_annotation
	EnableTableAnnotation *bool `json:"enableTableAnnotation,omitempty"`

	// Optional. If true, the LLM based annotation is added to the image
	//  during parsing.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.LayoutParsingConfig.enable_image_annotation
	EnableImageAnnotation *bool `json:"enableImageAnnotation,omitempty"`

	// Optional. Contains the required structure types to extract from the
	//  document. Supported values:
	//
	//  * `shareholder-structure`
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.LayoutParsingConfig.structured_content_types
	StructuredContentTypes []string `json:"structuredContentTypes,omitempty"`

	// Optional. List of HTML elements to exclude from the parsed content.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.LayoutParsingConfig.exclude_html_elements
	ExcludeHTMLElements []string `json:"excludeHTMLElements,omitempty"`

	// Optional. List of HTML classes to exclude from the parsed content.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.LayoutParsingConfig.exclude_html_classes
	ExcludeHTMLClasses []string `json:"excludeHTMLClasses,omitempty"`

	// Optional. List of HTML ids to exclude from the parsed content.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig.LayoutParsingConfig.exclude_html_ids
	ExcludeHTMLIds []string `json:"excludeHTMLIds,omitempty"`
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

	// Optional. If the flag set to true, we allow the agent and engine are in
	//  different locations, otherwise the agent and engine are required to be in
	//  the same location. The flag is set to false by default.
	//
	//  Note that the `allow_cross_region` are one-time consumed by and
	//  passed to
	//  [EngineService.CreateEngine][google.cloud.discoveryengine.v1.EngineService.CreateEngine].
	//  It means they cannot be retrieved using
	//  [EngineService.GetEngine][google.cloud.discoveryengine.v1.EngineService.GetEngine]
	//  or
	//  [EngineService.ListEngines][google.cloud.discoveryengine.v1.EngineService.ListEngines]
	//  API after engine creation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.ChatEngineConfig.allow_cross_region
	AllowCrossRegion *bool `json:"allowCrossRegion,omitempty"`
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

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig
type Engine_MediaRecommendationEngineConfig struct {
	// Required. The type of engine. e.g., `recommended-for-you`.
	//
	//  This field together with
	//  [optimization_objective][google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.optimization_objective]
	//  describe engine metadata to use to control engine training and serving.
	//
	//  Currently supported values: `recommended-for-you`, `others-you-may-like`,
	//  `more-like-this`, `most-popular-items`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.type
	Type *string `json:"type,omitempty"`

	// The optimization objective. e.g., `cvr`.
	//
	//  This field together with
	//  [optimization_objective][google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.type]
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
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.optimization_objective
	OptimizationObjective *string `json:"optimizationObjective,omitempty"`

	// Name and value of the custom threshold for cvr optimization_objective.
	//  For target_field `watch-time`, target_field_value must be an integer
	//  value indicating the media progress time in seconds between (0, 86400]
	//  (excludes 0, includes 86400) (e.g., 90).
	//  For target_field `watch-percentage`, the target_field_value must be a
	//  valid float value between (0, 1.0] (excludes 0, includes 1.0) (e.g.,
	//  0.5).
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.optimization_objective_config
	OptimizationObjectiveConfig *Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig `json:"optimizationObjectiveConfig,omitempty"`

	// The training state that the engine is in (e.g.
	//  `TRAINING` or `PAUSED`).
	//
	//  Since part of the cost of running the service
	//  is frequency of training - this can be used to determine when to train
	//  engine in order to control cost. If not specified: the default value for
	//  `CreateEngine` method is `TRAINING`. The default value for
	//  `UpdateEngine` method is to keep the state the same as before.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.training_state
	TrainingState *string `json:"trainingState,omitempty"`

	// Optional. Additional engine features config.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.engine_features_config
	EngineFeaturesConfig *Engine_MediaRecommendationEngineConfig_EngineFeaturesConfig `json:"engineFeaturesConfig,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.EngineFeaturesConfig
type Engine_MediaRecommendationEngineConfig_EngineFeaturesConfig struct {
	// Recommended for you engine feature config.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.EngineFeaturesConfig.recommended_for_you_config
	RecommendedForYouConfig *Engine_MediaRecommendationEngineConfig_RecommendedForYouFeatureConfig `json:"recommendedForYouConfig,omitempty"`

	// Most popular engine feature config.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.EngineFeaturesConfig.most_popular_config
	MostPopularConfig *Engine_MediaRecommendationEngineConfig_MostPopularFeatureConfig `json:"mostPopularConfig,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.MostPopularFeatureConfig
type Engine_MediaRecommendationEngineConfig_MostPopularFeatureConfig struct {
	// The time window of which the engine is queried at training and
	//  prediction time. Positive integers only. The value translates to the
	//  last X days of events. Currently required for the `most-popular-items`
	//  engine.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.MostPopularFeatureConfig.time_window_days
	TimeWindowDays *int64 `json:"timeWindowDays,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.OptimizationObjectiveConfig
type Engine_MediaRecommendationEngineConfig_OptimizationObjectiveConfig struct {
	// Required. The name of the field to target. Currently supported
	//  values: `watch-percentage`, `watch-time`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.OptimizationObjectiveConfig.target_field
	TargetField *string `json:"targetField,omitempty"`

	// Required. The threshold to be applied to the target (e.g., 0.5).
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.OptimizationObjectiveConfig.target_field_value_float
	TargetFieldValueFloat *float32 `json:"targetFieldValueFloat,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.RecommendedForYouFeatureConfig
type Engine_MediaRecommendationEngineConfig_RecommendedForYouFeatureConfig struct {
	// The type of event with which the engine is queried at prediction time.
	//  If set to `generic`, only `view-item`, `media-play`,and
	//  `media-complete` will be used as `context-event` in engine training. If
	//  set to `view-home-page`, `view-home-page` will also be used as
	//  `context-events` in addition to `view-item`, `media-play`, and
	//  `media-complete`. Currently supported for the `recommended-for-you`
	//  engine. Currently supported values: `view-home-page`, `generic`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Engine.MediaRecommendationEngineConfig.RecommendedForYouFeatureConfig.context_event_type
	ContextEventType *string `json:"contextEventType,omitempty"`
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

// +kcc:proto=google.cloud.discoveryengine.v1.HealthcareFhirConfig
type HealthcareFhirConfig struct {
	// Whether to enable configurable schema for `HEALTHCARE_FHIR` vertical.
	//
	//  If set to `true`, the predefined healthcare fhir schema can be extended
	//  for more customized searching and filtering.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.HealthcareFhirConfig.enable_configurable_schema
	EnableConfigurableSchema *bool `json:"enableConfigurableSchema,omitempty"`

	// Whether to enable static indexing for `HEALTHCARE_FHIR` batch
	//  ingestion.
	//
	//  If set to `true`, the batch ingestion will be processed in a static
	//  indexing mode which is slower but more capable of handling larger
	//  volume.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.HealthcareFhirConfig.enable_static_indexing_for_batch_ingestion
	EnableStaticIndexingForBatchIngestion *bool `json:"enableStaticIndexingForBatchIngestion,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Schema
type Schema struct {
	// The structured representation of the schema.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Schema.struct_schema
	StructSchema apiextensionsv1.JSON `json:"structSchema,omitempty"`

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

// +kcc:proto=google.cloud.discoveryengine.v1.SingleRegionKey
type SingleRegionKey struct {
	// Required. Single-regional kms key resource name which will be used to
	//  encrypt resources
	//  `projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{keyId}`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.SingleRegionKey.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
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

// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.CmekConfig
type CmekConfigObservedState struct {
	// Output only. The states of the CmekConfig.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.state
	State *string `json:"state,omitempty"`

	// Output only. The default CmekConfig for the Customer.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.is_default
	IsDefault *bool `json:"isDefault,omitempty"`

	// Output only. The timestamp of the last key rotation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.last_rotation_timestamp_micros
	LastRotationTimestampMicros *int64 `json:"lastRotationTimestampMicros,omitempty"`

	// Output only. Whether the NotebookLM Corpus is ready to be used.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.CmekConfig.notebooklm_state
	NotebooklmState *string `json:"notebooklmState,omitempty"`
}
