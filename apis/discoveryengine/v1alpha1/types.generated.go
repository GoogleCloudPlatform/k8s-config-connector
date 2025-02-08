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


// +kcc:proto=google.cloud.discoveryengine.v1beta.DataStore
type DataStore struct {
	// Immutable. The full resource name of the data store.
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}`.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 1024
	//  characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.name
	Name *string `json:"name,omitempty"`

	// Required. The data store display name.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 128
	//  characters. Otherwise, an INVALID_ARGUMENT error is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Immutable. The industry vertical that the data store registers.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.industry_vertical
	IndustryVertical *string `json:"industryVertical,omitempty"`

	// The solutions that the data store enrolls. Available solutions for each
	//  [industry_vertical][google.cloud.discoveryengine.v1beta.DataStore.industry_vertical]:
	//
	//  * `MEDIA`: `SOLUTION_TYPE_RECOMMENDATION` and `SOLUTION_TYPE_SEARCH`.
	//  * `SITE_SEARCH`: `SOLUTION_TYPE_SEARCH` is automatically enrolled. Other
	//    solutions cannot be enrolled.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.solution_types
	SolutionTypes []string `json:"solutionTypes,omitempty"`

	// Immutable. The content config of the data store. If this field is unset,
	//  the server behavior defaults to
	//  [ContentConfig.NO_CONTENT][google.cloud.discoveryengine.v1beta.DataStore.ContentConfig.NO_CONTENT].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.content_config
	ContentConfig *string `json:"contentConfig,omitempty"`

	// Language info for DataStore.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.language_info
	LanguageInfo *LanguageInfo `json:"languageInfo,omitempty"`

	// Optional. Configuration for Natural Language Query Understanding.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.natural_language_query_understanding_config
	NaturalLanguageQueryUnderstandingConfig *NaturalLanguageQueryUnderstandingConfig `json:"naturalLanguageQueryUnderstandingConfig,omitempty"`

	// Config to store data store type configuration for workspace data. This
	//  must be set when
	//  [DataStore.content_config][google.cloud.discoveryengine.v1beta.DataStore.content_config]
	//  is set as
	//  [DataStore.ContentConfig.GOOGLE_WORKSPACE][google.cloud.discoveryengine.v1beta.DataStore.ContentConfig.GOOGLE_WORKSPACE].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.workspace_config
	WorkspaceConfig *WorkspaceConfig `json:"workspaceConfig,omitempty"`

	// Configuration for Document understanding and enrichment.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.document_processing_config
	DocumentProcessingConfig *DocumentProcessingConfig `json:"documentProcessingConfig,omitempty"`

	// The start schema to use for this
	//  [DataStore][google.cloud.discoveryengine.v1beta.DataStore] when
	//  provisioning it. If unset, a default vertical specialized schema will be
	//  used.
	//
	//  This field is only used by [CreateDataStore][] API, and will be ignored if
	//  used in other APIs. This field will be omitted from all API responses
	//  including [CreateDataStore][] API. To retrieve a schema of a
	//  [DataStore][google.cloud.discoveryengine.v1beta.DataStore], use
	//  [SchemaService.GetSchema][google.cloud.discoveryengine.v1beta.SchemaService.GetSchema]
	//  API instead.
	//
	//  The provided schema will be validated against certain rules on schema.
	//  Learn more from [this
	//  doc](https://cloud.google.com/generative-ai-app-builder/docs/provide-schema).
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.starting_schema
	StartingSchema *Schema `json:"startingSchema,omitempty"`

	// Optional. Stores serving config at DataStore level.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.serving_config_data_store
	ServingConfigDataStore *DataStore_ServingConfigDataStore `json:"servingConfigDataStore,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.DataStore.BillingEstimation
type DataStore_BillingEstimation struct {
	// Data size for structured data in terms of bytes.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.BillingEstimation.structured_data_size
	StructuredDataSize *int64 `json:"structuredDataSize,omitempty"`

	// Data size for unstructured data in terms of bytes.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.BillingEstimation.unstructured_data_size
	UnstructuredDataSize *int64 `json:"unstructuredDataSize,omitempty"`

	// Data size for websites in terms of bytes.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.BillingEstimation.website_data_size
	WebsiteDataSize *int64 `json:"websiteDataSize,omitempty"`

	// Last updated timestamp for structured data.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.BillingEstimation.structured_data_update_time
	StructuredDataUpdateTime *string `json:"structuredDataUpdateTime,omitempty"`

	// Last updated timestamp for unstructured data.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.BillingEstimation.unstructured_data_update_time
	UnstructuredDataUpdateTime *string `json:"unstructuredDataUpdateTime,omitempty"`

	// Last updated timestamp for websites.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.BillingEstimation.website_data_update_time
	WebsiteDataUpdateTime *string `json:"websiteDataUpdateTime,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.DataStore.ServingConfigDataStore
type DataStore_ServingConfigDataStore struct {
	// If set true, the DataStore will not be available for serving search
	//  requests.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.ServingConfigDataStore.disabled_for_serving
	DisabledForServing *bool `json:"disabledForServing,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig
type DocumentProcessingConfig struct {
	// The full resource name of the Document Processing Config.
	//  Format:
	//  `projects/*/locations/*/collections/*/dataStores/*/documentProcessingConfig`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.name
	Name *string `json:"name,omitempty"`

	// Whether chunking mode is enabled.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.chunking_config
	ChunkingConfig *DocumentProcessingConfig_ChunkingConfig `json:"chunkingConfig,omitempty"`

	// Configurations for default Document parser.
	//  If not specified, we will configure it as default DigitalParsingConfig, and
	//  the default parsing config will be applied to all file types for Document
	//  parsing.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.default_parsing_config
	DefaultParsingConfig *DocumentProcessingConfig_ParsingConfig `json:"defaultParsingConfig,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ChunkingConfig
type DocumentProcessingConfig_ChunkingConfig struct {
	// Configuration for the layout based chunking.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ChunkingConfig.layout_based_chunking_config
	LayoutBasedChunkingConfig *DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig `json:"layoutBasedChunkingConfig,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ChunkingConfig.LayoutBasedChunkingConfig
type DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig struct {
	// The token size limit for each chunk.
	//
	//  Supported values: 100-500 (inclusive).
	//  Default value: 500.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ChunkingConfig.LayoutBasedChunkingConfig.chunk_size
	ChunkSize *int32 `json:"chunkSize,omitempty"`

	// Whether to include appending different levels of headings to chunks
	//  from the middle of the document to prevent context loss.
	//
	//  Default value: False.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ChunkingConfig.LayoutBasedChunkingConfig.include_ancestor_headings
	IncludeAncestorHeadings *bool `json:"includeAncestorHeadings,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig
type DocumentProcessingConfig_ParsingConfig struct {
	// Configurations applied to digital parser.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.digital_parsing_config
	DigitalParsingConfig *DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig `json:"digitalParsingConfig,omitempty"`

	// Configurations applied to OCR parser. Currently it only applies to
	//  PDFs.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.ocr_parsing_config
	OcrParsingConfig *DocumentProcessingConfig_ParsingConfig_OcrParsingConfig `json:"ocrParsingConfig,omitempty"`

	// Configurations applied to layout parser.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.layout_parsing_config
	LayoutParsingConfig *DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig `json:"layoutParsingConfig,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.DigitalParsingConfig
type DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig struct {
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.LayoutParsingConfig
type DocumentProcessingConfig_ParsingConfig_LayoutParsingConfig struct {
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.OcrParsingConfig
type DocumentProcessingConfig_ParsingConfig_OcrParsingConfig struct {
	// [DEPRECATED] This field is deprecated. To use the additional enhanced
	//  document elements processing, please switch to `layout_parsing_config`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.OcrParsingConfig.enhanced_document_elements
	EnhancedDocumentElements []string `json:"enhancedDocumentElements,omitempty"`

	// If true, will use native text instead of OCR text on pages containing
	//  native text.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DocumentProcessingConfig.ParsingConfig.OcrParsingConfig.use_native_text
	UseNativeText *bool `json:"useNativeText,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.LanguageInfo
type LanguageInfo struct {
	// The language code for the DataStore.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.LanguageInfo.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.NaturalLanguageQueryUnderstandingConfig
type NaturalLanguageQueryUnderstandingConfig struct {
	// Mode of Natural Language Query Understanding. If this field is unset, the
	//  behavior defaults to
	//  [NaturalLanguageQueryUnderstandingConfig.Mode.DISABLED][google.cloud.discoveryengine.v1beta.NaturalLanguageQueryUnderstandingConfig.Mode.DISABLED].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.NaturalLanguageQueryUnderstandingConfig.mode
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Schema
type Schema struct {
	// The structured representation of the schema.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Schema.struct_schema
	StructSchema map[string]string `json:"structSchema,omitempty"`

	// The JSON representation of the schema.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Schema.json_schema
	JsonSchema *string `json:"jsonSchema,omitempty"`

	// Immutable. The full resource name of the schema, in the format of
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/schemas/{schema}`.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 1024
	//  characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Schema.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.WorkspaceConfig
type WorkspaceConfig struct {
	// The Google Workspace data source.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.WorkspaceConfig.type
	Type *string `json:"type,omitempty"`

	// Obfuscated Dasher customer ID.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.WorkspaceConfig.dasher_customer_id
	DasherCustomerID *string `json:"dasherCustomerID,omitempty"`

	// Optional. The super admin service account for the workspace that will be
	//  used for access token generation. For now we only use it for Native Google
	//  Drive connector data ingestion.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.WorkspaceConfig.super_admin_service_account
	SuperAdminServiceAccount *string `json:"superAdminServiceAccount,omitempty"`

	// Optional. The super admin email address for the workspace that will be used
	//  for access token generation. For now we only use it for Native Google Drive
	//  connector data ingestion.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.WorkspaceConfig.super_admin_email_address
	SuperAdminEmailAddress *string `json:"superAdminEmailAddress,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.DataStore
type DataStoreObservedState struct {
	// Output only. The id of the default
	//  [Schema][google.cloud.discoveryengine.v1beta.Schema] asscociated to this
	//  data store.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.default_schema_id
	DefaultSchemaID *string `json:"defaultSchemaID,omitempty"`

	// Output only. Timestamp the
	//  [DataStore][google.cloud.discoveryengine.v1beta.DataStore] was created at.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Language info for DataStore.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.language_info
	LanguageInfo *LanguageInfoObservedState `json:"languageInfo,omitempty"`

	// Output only. Data size estimation for billing.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.DataStore.billing_estimation
	BillingEstimation *DataStore_BillingEstimation `json:"billingEstimation,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.LanguageInfo
type LanguageInfoObservedState struct {
	// Output only. This is the normalized form of language_code.
	//  E.g.: language_code of `en-GB`, `en_GB`, `en-UK` or `en-gb`
	//  will have normalized_language_code of `en-GB`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.LanguageInfo.normalized_language_code
	NormalizedLanguageCode *string `json:"normalizedLanguageCode,omitempty"`

	// Output only. Language part of normalized_language_code.
	//  E.g.: `en-US` -> `en`, `zh-Hans-HK` -> `zh`, `en` -> `en`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.LanguageInfo.language
	Language *string `json:"language,omitempty"`

	// Output only. Region part of normalized_language_code, if present.
	//  E.g.: `en-US` -> `US`, `zh-Hans-HK` -> `HK`, `en` -> ``.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.LanguageInfo.region
	Region *string `json:"region,omitempty"`
}
