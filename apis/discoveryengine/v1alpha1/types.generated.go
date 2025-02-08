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


// +kcc:proto=google.cloud.discoveryengine.v1beta.Conversation
type Conversation struct {
	// Immutable. Fully qualified name
	//  `projects/{project}/locations/global/collections/{collection}/dataStore/*/conversations/*`
	//  or
	//  `projects/{project}/locations/global/collections/{collection}/engines/*/conversations/*`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Conversation.name
	Name *string `json:"name,omitempty"`

	// The state of the Conversation.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Conversation.state
	State *string `json:"state,omitempty"`

	// A unique identifier for tracking users.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Conversation.user_pseudo_id
	UserPseudoID *string `json:"userPseudoID,omitempty"`

	// Conversation messages.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Conversation.messages
	Messages []ConversationMessage `json:"messages,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.ConversationContext
type ConversationContext struct {
	// The current list of documents the user is seeing.
	//  It contains the document resource references.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ConversationContext.context_documents
	ContextDocuments []string `json:"contextDocuments,omitempty"`

	// The current active document the user opened.
	//  It contains the document resource reference.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ConversationContext.active_document
	ActiveDocument *string `json:"activeDocument,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.ConversationMessage
type ConversationMessage struct {
	// User text input.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ConversationMessage.user_input
	UserInput *TextInput `json:"userInput,omitempty"`

	// Search reply.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ConversationMessage.reply
	Reply *Reply `json:"reply,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Reply
type Reply struct {
	// DEPRECATED: use `summary` instead.
	//  Text reply.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Reply.reply
	Reply *string `json:"reply,omitempty"`

	// References in the reply.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Reply.references
	References []Reply_Reference `json:"references,omitempty"`

	// Summary based on search results.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Reply.summary
	Summary *SearchResponse_Summary `json:"summary,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Reply.Reference
type Reply_Reference struct {
	// URI link reference.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Reply.Reference.uri
	URI *string `json:"uri,omitempty"`

	// Anchor text.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Reply.Reference.anchor_text
	AnchorText *string `json:"anchorText,omitempty"`

	// Anchor text start index.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Reply.Reference.start
	Start *int32 `json:"start,omitempty"`

	// Anchor text end index.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Reply.Reference.end
	End *int32 `json:"end,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchResponse.Summary
type SearchResponse_Summary struct {
	// The summary content.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.summary_text
	SummaryText *string `json:"summaryText,omitempty"`

	// Additional summary-skipped reasons. This provides the reason for ignored
	//  cases. If nothing is skipped, this field is not set.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.summary_skipped_reasons
	SummarySkippedReasons []string `json:"summarySkippedReasons,omitempty"`

	// A collection of Safety Attribute categories and their associated
	//  confidence scores.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.safety_attributes
	SafetyAttributes *SearchResponse_Summary_SafetyAttributes `json:"safetyAttributes,omitempty"`

	// Summary with metadata information.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.summary_with_metadata
	SummaryWithMetadata *SearchResponse_Summary_SummaryWithMetadata `json:"summaryWithMetadata,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.Citation
type SearchResponse_Summary_Citation struct {
	// Index indicates the start of the segment, measured in bytes/unicode.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.Citation.start_index
	StartIndex *int64 `json:"startIndex,omitempty"`

	// End of the attributed segment, exclusive.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.Citation.end_index
	EndIndex *int64 `json:"endIndex,omitempty"`

	// Citation sources for the attributed segment.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.Citation.sources
	Sources []SearchResponse_Summary_CitationSource `json:"sources,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.CitationMetadata
type SearchResponse_Summary_CitationMetadata struct {
	// Citations for segments.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.CitationMetadata.citations
	Citations []SearchResponse_Summary_Citation `json:"citations,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.CitationSource
type SearchResponse_Summary_CitationSource struct {
	// Document reference index from SummaryWithMetadata.references.
	//  It is 0-indexed and the value will be zero if the reference_index is
	//  not set explicitly.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.CitationSource.reference_index
	ReferenceIndex *int64 `json:"referenceIndex,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.Reference
type SearchResponse_Summary_Reference struct {
	// Title of the document.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.Reference.title
	Title *string `json:"title,omitempty"`

	// Required.
	//  [Document.name][google.cloud.discoveryengine.v1beta.Document.name] of
	//  the document. Full resource name of the referenced document, in the
	//  format
	//  `projects/*/locations/*/collections/*/dataStores/*/branches/*/documents/*`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.Reference.document
	Document *string `json:"document,omitempty"`

	// Cloud Storage or HTTP uri for the document.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.Reference.uri
	URI *string `json:"uri,omitempty"`

	// List of cited chunk contents derived from document content.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.Reference.chunk_contents
	ChunkContents []SearchResponse_Summary_Reference_ChunkContent `json:"chunkContents,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.Reference.ChunkContent
type SearchResponse_Summary_Reference_ChunkContent struct {
	// Chunk textual content.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.Reference.ChunkContent.content
	Content *string `json:"content,omitempty"`

	// Page identifier.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.Reference.ChunkContent.page_identifier
	PageIdentifier *string `json:"pageIdentifier,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.SafetyAttributes
type SearchResponse_Summary_SafetyAttributes struct {
	// The display names of Safety Attribute categories associated with the
	//  generated content. Order matches the Scores.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.SafetyAttributes.categories
	Categories []string `json:"categories,omitempty"`

	// The confidence scores of the each category, higher
	//  value means higher confidence. Order matches the Categories.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.SafetyAttributes.scores
	Scores []float32 `json:"scores,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.SummaryWithMetadata
type SearchResponse_Summary_SummaryWithMetadata struct {
	// Summary text with no citation information.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.SummaryWithMetadata.summary
	Summary *string `json:"summary,omitempty"`

	// Citation metadata for given summary.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.SummaryWithMetadata.citation_metadata
	CitationMetadata *SearchResponse_Summary_CitationMetadata `json:"citationMetadata,omitempty"`

	// Document References.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SearchResponse.Summary.SummaryWithMetadata.references
	References []SearchResponse_Summary_Reference `json:"references,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.TextInput
type TextInput struct {
	// Text input.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TextInput.input
	Input *string `json:"input,omitempty"`

	// Conversation context of the input.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TextInput.context
	Context *ConversationContext `json:"context,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.Conversation
type ConversationObservedState struct {
	// Conversation messages.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Conversation.messages
	Messages []ConversationMessageObservedState `json:"messages,omitempty"`

	// Output only. The time the conversation started.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Conversation.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time the conversation finished.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.Conversation.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.ConversationMessage
type ConversationMessageObservedState struct {
	// Output only. Message creation timestamp.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.ConversationMessage.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
