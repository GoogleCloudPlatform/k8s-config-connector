// Copyright 2024 Google LLC
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
	StructuredDataSize *int64 `json:"structuredDataSize,omitempty"`

	// Data size for unstructured data in terms of bytes.
	UnstructuredDataSize *int64 `json:"unstructuredDataSize,omitempty"`

	// Data size for websites in terms of bytes.
	WebsiteDataSize *int64 `json:"websiteDataSize,omitempty"`

	// Last updated timestamp for structured data.
	StructuredDataUpdateTime *string `json:"structuredDataUpdateTime,omitempty"`

	// Last updated timestamp for unstructured data.
	UnstructuredDataUpdateTime *string `json:"unstructuredDataUpdateTime,omitempty"`

	// Last updated timestamp for websites.
	WebsiteDataUpdateTime *string `json:"websiteDataUpdateTime,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.DocumentProcessingConfig
type DocumentProcessingConfig struct {
	// The full resource name of the Document Processing Config.
	//  Format:
	//  `projects/*/locations/*/collections/*/dataStores/*/documentProcessingConfig`.
	Name *string `json:"name,omitempty"`

	// Whether chunking mode is enabled.
	ChunkingConfig *DocumentProcessingConfig_ChunkingConfig `json:"chunkingConfig,omitempty"`

	// Configurations for default Document parser.
	//  If not specified, we will configure it as default DigitalParsingConfig, and
	//  the default parsing config will be applied to all file types for Document
	//  parsing.
	DefaultParsingConfig *DocumentProcessingConfig_ParsingConfig `json:"defaultParsingConfig,omitempty"`

	// TODO: map type string message for parsing_config_overrides

}

// +kcc:proto=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ChunkingConfig
type DocumentProcessingConfig_ChunkingConfig struct {
	// Configuration for the layout based chunking.
	LayoutBasedChunkingConfig *DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig `json:"layoutBasedChunkingConfig,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ChunkingConfig.LayoutBasedChunkingConfig
type DocumentProcessingConfig_ChunkingConfig_LayoutBasedChunkingConfig struct {
	// The token size limit for each chunk.
	//
	//  Supported values: 100-500 (inclusive).
	//  Default value: 500.
	ChunkSize *int32 `json:"chunkSize,omitempty"`

	// Whether to include appending different levels of headings to chunks
	//  from the middle of the document to prevent context loss.
	//
	//  Default value: False.
	IncludeAncestorHeadings *bool `json:"includeAncestorHeadings,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.DocumentProcessingConfig.ParsingConfig
type DocumentProcessingConfig_ParsingConfig struct {
	// Configurations applied to digital parser.
	DigitalParsingConfig *DocumentProcessingConfig_ParsingConfig_DigitalParsingConfig `json:"digitalParsingConfig,omitempty"`

	// Configurations applied to OCR parser. Currently it only applies to
	//  PDFs.
	OcrParsingConfig *DocumentProcessingConfig_ParsingConfig_OcrParsingConfig `json:"ocrParsingConfig,omitempty"`

	// Configurations applied to layout parser.
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
	EnhancedDocumentElements []string `json:"enhancedDocumentElements,omitempty"`

	// If true, will use native text instead of OCR text on pages containing
	//  native text.
	UseNativeText *bool `json:"useNativeText,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Schema
type Schema struct {
	// The structured representation of the schema.
	StructSchema map[string]string `json:"structSchema,omitempty"`

	// The JSON representation of the schema.
	JsonSchema *string `json:"jsonSchema,omitempty"`

	// Immutable. The full resource name of the schema, in the format of
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/schemas/{schema}`.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 1024
	//  characters.
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.WorkspaceConfig
type WorkspaceConfig struct {
	// The Google Workspace data source.
	Type *string `json:"type,omitempty"`

	// Obfuscated Dasher customer ID.
	DasherCustomerID *string `json:"dasherCustomerID,omitempty"`

	// Optional. The super admin service account for the workspace that will be
	//  used for access token generation. For now we only use it for Native Google
	//  Drive connector data ingestion.
	SuperAdminServiceAccount *string `json:"superAdminServiceAccount,omitempty"`

	// Optional. The super admin email address for the workspace that will be used
	//  for access token generation. For now we only use it for Native Google Drive
	//  connector data ingestion.
	SuperAdminEmailAddress *string `json:"superAdminEmailAddress,omitempty"`
}
