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


// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer
type Answer struct {
	// Immutable. Fully qualified name
	//  `projects/{project}/locations/global/collections/{collection}/engines/{engine}/sessions/*/answers/*`
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.name
	Name *string `json:"name,omitempty"`

	// The state of the answer generation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.state
	State *string `json:"state,omitempty"`

	// The textual answer.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.answer_text
	AnswerText *string `json:"answerText,omitempty"`

	// Citations.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.citations
	Citations []Answer_Citation `json:"citations,omitempty"`

	// References.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.references
	References []Answer_Reference `json:"references,omitempty"`

	// Suggested related questions.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.related_questions
	RelatedQuestions []string `json:"relatedQuestions,omitempty"`

	// Answer generation steps.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.steps
	Steps []Answer_Step `json:"steps,omitempty"`

	// Query understanding information.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.query_understanding_info
	QueryUnderstandingInfo *Answer_QueryUnderstandingInfo `json:"queryUnderstandingInfo,omitempty"`

	// Additional answer-skipped reasons. This provides the reason for ignored
	//  cases. If nothing is skipped, this field is not set.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.answer_skipped_reasons
	AnswerSkippedReasons []string `json:"answerSkippedReasons,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Citation
type Answer_Citation struct {
	// Index indicates the start of the segment, measured in bytes (UTF-8
	//  unicode).
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Citation.start_index
	StartIndex *int64 `json:"startIndex,omitempty"`

	// End of the attributed segment, exclusive.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Citation.end_index
	EndIndex *int64 `json:"endIndex,omitempty"`

	// Citation sources for the attributed segment.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Citation.sources
	Sources []Answer_CitationSource `json:"sources,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.CitationSource
type Answer_CitationSource struct {
	// ID of the citation source.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.CitationSource.reference_id
	ReferenceID *string `json:"referenceID,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.QueryUnderstandingInfo
type Answer_QueryUnderstandingInfo struct {
	// Query classification information.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.QueryUnderstandingInfo.query_classification_info
	QueryClassificationInfo []Answer_QueryUnderstandingInfo_QueryClassificationInfo `json:"queryClassificationInfo,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.QueryUnderstandingInfo.QueryClassificationInfo
type Answer_QueryUnderstandingInfo_QueryClassificationInfo struct {
	// Query classification type.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.QueryUnderstandingInfo.QueryClassificationInfo.type
	Type *string `json:"type,omitempty"`

	// Classification output.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.QueryUnderstandingInfo.QueryClassificationInfo.positive
	Positive *bool `json:"positive,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Reference
type Answer_Reference struct {
	// Unstructured document information.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.unstructured_document_info
	UnstructuredDocumentInfo *Answer_Reference_UnstructuredDocumentInfo `json:"unstructuredDocumentInfo,omitempty"`

	// Chunk information.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.chunk_info
	ChunkInfo *Answer_Reference_ChunkInfo `json:"chunkInfo,omitempty"`

	// Structured document information.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.structured_document_info
	StructuredDocumentInfo *Answer_Reference_StructuredDocumentInfo `json:"structuredDocumentInfo,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Reference.ChunkInfo
type Answer_Reference_ChunkInfo struct {
	// Chunk resource name.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.ChunkInfo.chunk
	Chunk *string `json:"chunk,omitempty"`

	// Chunk textual content.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.ChunkInfo.content
	Content *string `json:"content,omitempty"`

	// The relevance of the chunk for a given query. Values range from 0.0
	//  (completely irrelevant) to 1.0 (completely relevant).
	//  This value is for informational purpose only. It may change for
	//  the same query and chunk at any time due to a model retraining or
	//  change in implementation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.ChunkInfo.relevance_score
	RelevanceScore *float32 `json:"relevanceScore,omitempty"`

	// Document metadata.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.ChunkInfo.document_metadata
	DocumentMetadata *Answer_Reference_ChunkInfo_DocumentMetadata `json:"documentMetadata,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Reference.ChunkInfo.DocumentMetadata
type Answer_Reference_ChunkInfo_DocumentMetadata struct {
	// Document resource name.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.ChunkInfo.DocumentMetadata.document
	Document *string `json:"document,omitempty"`

	// URI for the document.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.ChunkInfo.DocumentMetadata.uri
	URI *string `json:"uri,omitempty"`

	// Title.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.ChunkInfo.DocumentMetadata.title
	Title *string `json:"title,omitempty"`

	// Page identifier.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.ChunkInfo.DocumentMetadata.page_identifier
	PageIdentifier *string `json:"pageIdentifier,omitempty"`

	// The structured JSON metadata for the document.
	//  It is populated from the struct data from the Chunk in search result.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.ChunkInfo.DocumentMetadata.struct_data
	StructData map[string]string `json:"structData,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Reference.StructuredDocumentInfo
type Answer_Reference_StructuredDocumentInfo struct {
	// Document resource name.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.StructuredDocumentInfo.document
	Document *string `json:"document,omitempty"`

	// Structured search data.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.StructuredDocumentInfo.struct_data
	StructData map[string]string `json:"structData,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Reference.UnstructuredDocumentInfo
type Answer_Reference_UnstructuredDocumentInfo struct {
	// Document resource name.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.UnstructuredDocumentInfo.document
	Document *string `json:"document,omitempty"`

	// URI for the document.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.UnstructuredDocumentInfo.uri
	URI *string `json:"uri,omitempty"`

	// Title.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.UnstructuredDocumentInfo.title
	Title *string `json:"title,omitempty"`

	// List of cited chunk contents derived from document content.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.UnstructuredDocumentInfo.chunk_contents
	ChunkContents []Answer_Reference_UnstructuredDocumentInfo_ChunkContent `json:"chunkContents,omitempty"`

	// The structured JSON metadata for the document.
	//  It is populated from the struct data from the Chunk in search result.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.UnstructuredDocumentInfo.struct_data
	StructData map[string]string `json:"structData,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Reference.UnstructuredDocumentInfo.ChunkContent
type Answer_Reference_UnstructuredDocumentInfo_ChunkContent struct {
	// Chunk textual content.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.UnstructuredDocumentInfo.ChunkContent.content
	Content *string `json:"content,omitempty"`

	// Page identifier.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.UnstructuredDocumentInfo.ChunkContent.page_identifier
	PageIdentifier *string `json:"pageIdentifier,omitempty"`

	// The relevance of the chunk for a given query. Values range from 0.0
	//  (completely irrelevant) to 1.0 (completely relevant).
	//  This value is for informational purpose only. It may change for
	//  the same query and chunk at any time due to a model retraining or
	//  change in implementation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Reference.UnstructuredDocumentInfo.ChunkContent.relevance_score
	RelevanceScore *float32 `json:"relevanceScore,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Step
type Answer_Step struct {
	// The state of the step.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.state
	State *string `json:"state,omitempty"`

	// The description of the step.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.description
	Description *string `json:"description,omitempty"`

	// The thought of the step.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.thought
	Thought *string `json:"thought,omitempty"`

	// Actions.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.actions
	Actions []Answer_Step_Action `json:"actions,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Step.Action
type Answer_Step_Action struct {
	// Search action.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.search_action
	SearchAction *Answer_Step_Action_SearchAction `json:"searchAction,omitempty"`

	// Observation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.observation
	Observation *Answer_Step_Action_Observation `json:"observation,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation
type Answer_Step_Action_Observation struct {
	// Search results observed by the search action, it can be snippets info
	//  or chunk info, depending on the citation type set by the user.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.search_results
	SearchResults []Answer_Step_Action_Observation_SearchResult `json:"searchResults,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult
type Answer_Step_Action_Observation_SearchResult struct {
	// Document resource name.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult.document
	Document *string `json:"document,omitempty"`

	// URI for the document.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult.uri
	URI *string `json:"uri,omitempty"`

	// Title.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult.title
	Title *string `json:"title,omitempty"`

	// If citation_type is DOCUMENT_LEVEL_CITATION, populate document
	//  level snippets.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult.snippet_info
	SnippetInfo []Answer_Step_Action_Observation_SearchResult_SnippetInfo `json:"snippetInfo,omitempty"`

	// If citation_type is CHUNK_LEVEL_CITATION and chunk mode is on,
	//  populate chunk info.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult.chunk_info
	ChunkInfo []Answer_Step_Action_Observation_SearchResult_ChunkInfo `json:"chunkInfo,omitempty"`

	// Data representation.
	//  The structured JSON data for the document.
	//  It's populated from the struct data from the Document, or the
	//  Chunk in search result.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult.struct_data
	StructData map[string]string `json:"structData,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult.ChunkInfo
type Answer_Step_Action_Observation_SearchResult_ChunkInfo struct {
	// Chunk resource name.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult.ChunkInfo.chunk
	Chunk *string `json:"chunk,omitempty"`

	// Chunk textual content.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult.ChunkInfo.content
	Content *string `json:"content,omitempty"`

	// The relevance of the chunk for a given query. Values range from
	//  0.0 (completely irrelevant) to 1.0 (completely relevant).
	//  This value is for informational purpose only. It may change for
	//  the same query and chunk at any time due to a model retraining or
	//  change in implementation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult.ChunkInfo.relevance_score
	RelevanceScore *float32 `json:"relevanceScore,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult.SnippetInfo
type Answer_Step_Action_Observation_SearchResult_SnippetInfo struct {
	// Snippet content.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult.SnippetInfo.snippet
	Snippet *string `json:"snippet,omitempty"`

	// Status of the snippet defined by the search team.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.Observation.SearchResult.SnippetInfo.snippet_status
	SnippetStatus *string `json:"snippetStatus,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer.Step.Action.SearchAction
type Answer_Step_Action_SearchAction struct {
	// The query to search.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.Step.Action.SearchAction.query
	Query *string `json:"query,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Answer
type AnswerObservedState struct {
	// Output only. Answer creation timestamp.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Answer completed timestamp.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Answer.complete_time
	CompleteTime *string `json:"completeTime,omitempty"`
}
