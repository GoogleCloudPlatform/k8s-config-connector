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


// +kcc:proto=google.cloud.discoveryengine.v1beta.Chunk
type Chunk struct {
	// The full resource name of the chunk.
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/branches/{branch}/documents/{document_id}/chunks/{chunk_id}`.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 1024
	//  characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.name
	Name *string `json:"name,omitempty"`

	// Unique chunk ID of the current chunk.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.id
	ID *string `json:"id,omitempty"`

	// Content is a string from a document (parsed content).
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.content
	Content *string `json:"content,omitempty"`

	// Metadata of the document from the current chunk.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.document_metadata
	DocumentMetadata *Chunk_DocumentMetadata `json:"documentMetadata,omitempty"`

	// Page span of the chunk.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.page_span
	PageSpan *Chunk_PageSpan `json:"pageSpan,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Chunk.ChunkMetadata
type Chunk_ChunkMetadata struct {
	// The previous chunks of the current chunk. The number is controlled by
	//  [SearchRequest.ContentSearchSpec.ChunkSpec.num_previous_chunks][google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.ChunkSpec.num_previous_chunks].
	//  This field is only populated on
	//  [SearchService.Search][google.cloud.discoveryengine.v1beta.SearchService.Search]
	//  API.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.ChunkMetadata.previous_chunks
	PreviousChunks []Chunk `json:"previousChunks,omitempty"`

	// The next chunks of the current chunk. The number is controlled by
	//  [SearchRequest.ContentSearchSpec.ChunkSpec.num_next_chunks][google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.ChunkSpec.num_next_chunks].
	//  This field is only populated on
	//  [SearchService.Search][google.cloud.discoveryengine.v1beta.SearchService.Search]
	//  API.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.ChunkMetadata.next_chunks
	NextChunks []Chunk `json:"nextChunks,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Chunk.DocumentMetadata
type Chunk_DocumentMetadata struct {
	// Uri of the document.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.DocumentMetadata.uri
	URI *string `json:"uri,omitempty"`

	// Title of the document.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.DocumentMetadata.title
	Title *string `json:"title,omitempty"`

	// Data representation.
	//  The structured JSON data for the document. It should conform to the
	//  registered [Schema][google.cloud.discoveryengine.v1beta.Schema] or an
	//  `INVALID_ARGUMENT` error is thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.DocumentMetadata.struct_data
	StructData map[string]string `json:"structData,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Chunk.PageSpan
type Chunk_PageSpan struct {
	// The start page of the chunk.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.PageSpan.page_start
	PageStart *int32 `json:"pageStart,omitempty"`

	// The end page of the chunk.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.PageSpan.page_end
	PageEnd *int32 `json:"pageEnd,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Chunk
type ChunkObservedState struct {
	// Output only. Represents the relevance score based on similarity.
	//  Higher score indicates higher chunk relevance.
	//  The score is in range [-1.0, 1.0].
	//  Only populated on [SearchService.SearchResponse][].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.relevance_score
	RelevanceScore *float64 `json:"relevanceScore,omitempty"`

	// Output only. This field is OUTPUT_ONLY.
	//  It contains derived data that are not in the original input document.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.derived_struct_data
	DerivedStructData map[string]string `json:"derivedStructData,omitempty"`

	// Output only. Represents the relevance score based on similarity.
	//  Higher score indicates higher chunk relevance.
	//  The score is in range [-1.0, 1.0].
	//  Only populated on [SearchService.SearchResponse][].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.relevance_score
	RelevanceScore *float64 `json:"relevanceScore,omitempty"`

	// Output only. This field is OUTPUT_ONLY.
	//  It contains derived data that are not in the original input document.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.derived_struct_data
	DerivedStructData map[string]string `json:"derivedStructData,omitempty"`

	// Output only. Metadata of the current chunk.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.chunk_metadata
	ChunkMetadata *Chunk_ChunkMetadata `json:"chunkMetadata,omitempty"`

	// Output only. Metadata of the current chunk.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.chunk_metadata
	ChunkMetadata *Chunk_ChunkMetadata `json:"chunkMetadata,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Chunk.ChunkMetadata
type Chunk_ChunkMetadataObservedState struct {
	// The previous chunks of the current chunk. The number is controlled by
	//  [SearchRequest.ContentSearchSpec.ChunkSpec.num_previous_chunks][google.cloud.discoveryengine.v1beta.SearchRequest.ContentSearchSpec.ChunkSpec.num_previous_chunks].
	//  This field is only populated on
	//  [SearchService.Search][google.cloud.discoveryengine.v1beta.SearchService.Search]
	//  API.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Chunk.ChunkMetadata.previous_chunks
	PreviousChunks []ChunkObservedState `json:"previousChunks,omitempty"`
}
