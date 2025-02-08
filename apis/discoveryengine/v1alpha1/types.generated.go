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
