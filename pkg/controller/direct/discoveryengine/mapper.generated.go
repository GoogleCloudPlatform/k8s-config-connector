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

package discoveryengine

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Answer_FromProto(mapCtx *direct.MapContext, in *pb.Answer) *krm.Answer {
	if in == nil {
		return nil
	}
	out := &krm.Answer{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.AnswerText = direct.LazyPtr(in.GetAnswerText())
	out.Citations = direct.Slice_FromProto(mapCtx, in.Citations, Answer_Citation_FromProto)
	out.References = direct.Slice_FromProto(mapCtx, in.References, Answer_Reference_FromProto)
	out.RelatedQuestions = in.RelatedQuestions
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, Answer_Step_FromProto)
	out.QueryUnderstandingInfo = Answer_QueryUnderstandingInfo_FromProto(mapCtx, in.GetQueryUnderstandingInfo())
	out.AnswerSkippedReasons = direct.EnumSlice_FromProto(mapCtx, in.AnswerSkippedReasons)
	// MISSING: CreateTime
	// MISSING: CompleteTime
	return out
}
func Answer_ToProto(mapCtx *direct.MapContext, in *krm.Answer) *pb.Answer {
	if in == nil {
		return nil
	}
	out := &pb.Answer{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.Answer_State](mapCtx, in.State)
	out.AnswerText = direct.ValueOf(in.AnswerText)
	out.Citations = direct.Slice_ToProto(mapCtx, in.Citations, Answer_Citation_ToProto)
	out.References = direct.Slice_ToProto(mapCtx, in.References, Answer_Reference_ToProto)
	out.RelatedQuestions = in.RelatedQuestions
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, Answer_Step_ToProto)
	out.QueryUnderstandingInfo = Answer_QueryUnderstandingInfo_ToProto(mapCtx, in.QueryUnderstandingInfo)
	out.AnswerSkippedReasons = direct.EnumSlice_ToProto[pb.Answer_AnswerSkippedReason](mapCtx, in.AnswerSkippedReasons)
	// MISSING: CreateTime
	// MISSING: CompleteTime
	return out
}
func AnswerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Answer) *krm.AnswerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnswerObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: AnswerText
	// MISSING: Citations
	// MISSING: References
	// MISSING: RelatedQuestions
	// MISSING: Steps
	// MISSING: QueryUnderstandingInfo
	// MISSING: AnswerSkippedReasons
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.CompleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCompleteTime())
	return out
}
func AnswerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnswerObservedState) *pb.Answer {
	if in == nil {
		return nil
	}
	out := &pb.Answer{}
	// MISSING: Name
	// MISSING: State
	// MISSING: AnswerText
	// MISSING: Citations
	// MISSING: References
	// MISSING: RelatedQuestions
	// MISSING: Steps
	// MISSING: QueryUnderstandingInfo
	// MISSING: AnswerSkippedReasons
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.CompleteTime = direct.StringTimestamp_ToProto(mapCtx, in.CompleteTime)
	return out
}
func Answer_Citation_FromProto(mapCtx *direct.MapContext, in *pb.Answer_Citation) *krm.Answer_Citation {
	if in == nil {
		return nil
	}
	out := &krm.Answer_Citation{}
	out.StartIndex = direct.LazyPtr(in.GetStartIndex())
	out.EndIndex = direct.LazyPtr(in.GetEndIndex())
	out.Sources = direct.Slice_FromProto(mapCtx, in.Sources, Answer_CitationSource_FromProto)
	return out
}
func Answer_Citation_ToProto(mapCtx *direct.MapContext, in *krm.Answer_Citation) *pb.Answer_Citation {
	if in == nil {
		return nil
	}
	out := &pb.Answer_Citation{}
	out.StartIndex = direct.ValueOf(in.StartIndex)
	out.EndIndex = direct.ValueOf(in.EndIndex)
	out.Sources = direct.Slice_ToProto(mapCtx, in.Sources, Answer_CitationSource_ToProto)
	return out
}
func Answer_CitationSource_FromProto(mapCtx *direct.MapContext, in *pb.Answer_CitationSource) *krm.Answer_CitationSource {
	if in == nil {
		return nil
	}
	out := &krm.Answer_CitationSource{}
	out.ReferenceID = direct.LazyPtr(in.GetReferenceId())
	return out
}
func Answer_CitationSource_ToProto(mapCtx *direct.MapContext, in *krm.Answer_CitationSource) *pb.Answer_CitationSource {
	if in == nil {
		return nil
	}
	out := &pb.Answer_CitationSource{}
	out.ReferenceId = direct.ValueOf(in.ReferenceID)
	return out
}
func Answer_QueryUnderstandingInfo_FromProto(mapCtx *direct.MapContext, in *pb.Answer_QueryUnderstandingInfo) *krm.Answer_QueryUnderstandingInfo {
	if in == nil {
		return nil
	}
	out := &krm.Answer_QueryUnderstandingInfo{}
	out.QueryClassificationInfo = direct.Slice_FromProto(mapCtx, in.QueryClassificationInfo, Answer_QueryUnderstandingInfo_QueryClassificationInfo_FromProto)
	return out
}
func Answer_QueryUnderstandingInfo_ToProto(mapCtx *direct.MapContext, in *krm.Answer_QueryUnderstandingInfo) *pb.Answer_QueryUnderstandingInfo {
	if in == nil {
		return nil
	}
	out := &pb.Answer_QueryUnderstandingInfo{}
	out.QueryClassificationInfo = direct.Slice_ToProto(mapCtx, in.QueryClassificationInfo, Answer_QueryUnderstandingInfo_QueryClassificationInfo_ToProto)
	return out
}
func Answer_QueryUnderstandingInfo_QueryClassificationInfo_FromProto(mapCtx *direct.MapContext, in *pb.Answer_QueryUnderstandingInfo_QueryClassificationInfo) *krm.Answer_QueryUnderstandingInfo_QueryClassificationInfo {
	if in == nil {
		return nil
	}
	out := &krm.Answer_QueryUnderstandingInfo_QueryClassificationInfo{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Positive = direct.LazyPtr(in.GetPositive())
	return out
}
func Answer_QueryUnderstandingInfo_QueryClassificationInfo_ToProto(mapCtx *direct.MapContext, in *krm.Answer_QueryUnderstandingInfo_QueryClassificationInfo) *pb.Answer_QueryUnderstandingInfo_QueryClassificationInfo {
	if in == nil {
		return nil
	}
	out := &pb.Answer_QueryUnderstandingInfo_QueryClassificationInfo{}
	out.Type = direct.Enum_ToProto[pb.Answer_QueryUnderstandingInfo_QueryClassificationInfo_Type](mapCtx, in.Type)
	out.Positive = direct.ValueOf(in.Positive)
	return out
}
func Answer_Reference_FromProto(mapCtx *direct.MapContext, in *pb.Answer_Reference) *krm.Answer_Reference {
	if in == nil {
		return nil
	}
	out := &krm.Answer_Reference{}
	out.UnstructuredDocumentInfo = Answer_Reference_UnstructuredDocumentInfo_FromProto(mapCtx, in.GetUnstructuredDocumentInfo())
	out.ChunkInfo = Answer_Reference_ChunkInfo_FromProto(mapCtx, in.GetChunkInfo())
	out.StructuredDocumentInfo = Answer_Reference_StructuredDocumentInfo_FromProto(mapCtx, in.GetStructuredDocumentInfo())
	return out
}
func Answer_Reference_ToProto(mapCtx *direct.MapContext, in *krm.Answer_Reference) *pb.Answer_Reference {
	if in == nil {
		return nil
	}
	out := &pb.Answer_Reference{}
	if oneof := Answer_Reference_UnstructuredDocumentInfo_ToProto(mapCtx, in.UnstructuredDocumentInfo); oneof != nil {
		out.Content = &pb.Answer_Reference_UnstructuredDocumentInfo_{UnstructuredDocumentInfo: oneof}
	}
	if oneof := Answer_Reference_ChunkInfo_ToProto(mapCtx, in.ChunkInfo); oneof != nil {
		out.Content = &pb.Answer_Reference_ChunkInfo_{ChunkInfo: oneof}
	}
	if oneof := Answer_Reference_StructuredDocumentInfo_ToProto(mapCtx, in.StructuredDocumentInfo); oneof != nil {
		out.Content = &pb.Answer_Reference_StructuredDocumentInfo_{StructuredDocumentInfo: oneof}
	}
	return out
}
func Answer_Reference_ChunkInfo_FromProto(mapCtx *direct.MapContext, in *pb.Answer_Reference_ChunkInfo) *krm.Answer_Reference_ChunkInfo {
	if in == nil {
		return nil
	}
	out := &krm.Answer_Reference_ChunkInfo{}
	out.Chunk = direct.LazyPtr(in.GetChunk())
	out.Content = direct.LazyPtr(in.GetContent())
	out.RelevanceScore = in.RelevanceScore
	out.DocumentMetadata = Answer_Reference_ChunkInfo_DocumentMetadata_FromProto(mapCtx, in.GetDocumentMetadata())
	return out
}
func Answer_Reference_ChunkInfo_ToProto(mapCtx *direct.MapContext, in *krm.Answer_Reference_ChunkInfo) *pb.Answer_Reference_ChunkInfo {
	if in == nil {
		return nil
	}
	out := &pb.Answer_Reference_ChunkInfo{}
	out.Chunk = direct.ValueOf(in.Chunk)
	out.Content = direct.ValueOf(in.Content)
	out.RelevanceScore = in.RelevanceScore
	out.DocumentMetadata = Answer_Reference_ChunkInfo_DocumentMetadata_ToProto(mapCtx, in.DocumentMetadata)
	return out
}
func Answer_Reference_StructuredDocumentInfo_FromProto(mapCtx *direct.MapContext, in *pb.Answer_Reference_StructuredDocumentInfo) *krm.Answer_Reference_StructuredDocumentInfo {
	if in == nil {
		return nil
	}
	out := &krm.Answer_Reference_StructuredDocumentInfo{}
	out.Document = direct.LazyPtr(in.GetDocument())
	out.StructData = StructData_FromProto(mapCtx, in.GetStructData())
	return out
}
func Answer_Reference_StructuredDocumentInfo_ToProto(mapCtx *direct.MapContext, in *krm.Answer_Reference_StructuredDocumentInfo) *pb.Answer_Reference_StructuredDocumentInfo {
	if in == nil {
		return nil
	}
	out := &pb.Answer_Reference_StructuredDocumentInfo{}
	out.Document = direct.ValueOf(in.Document)
	out.StructData = StructData_ToProto(mapCtx, in.StructData)
	return out
}
func Answer_Reference_UnstructuredDocumentInfo_FromProto(mapCtx *direct.MapContext, in *pb.Answer_Reference_UnstructuredDocumentInfo) *krm.Answer_Reference_UnstructuredDocumentInfo {
	if in == nil {
		return nil
	}
	out := &krm.Answer_Reference_UnstructuredDocumentInfo{}
	out.Document = direct.LazyPtr(in.GetDocument())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.ChunkContents = direct.Slice_FromProto(mapCtx, in.ChunkContents, Answer_Reference_UnstructuredDocumentInfo_ChunkContent_FromProto)
	out.StructData = StructData_FromProto(mapCtx, in.GetStructData())
	return out
}
func Answer_Reference_UnstructuredDocumentInfo_ToProto(mapCtx *direct.MapContext, in *krm.Answer_Reference_UnstructuredDocumentInfo) *pb.Answer_Reference_UnstructuredDocumentInfo {
	if in == nil {
		return nil
	}
	out := &pb.Answer_Reference_UnstructuredDocumentInfo{}
	out.Document = direct.ValueOf(in.Document)
	out.Uri = direct.ValueOf(in.URI)
	out.Title = direct.ValueOf(in.Title)
	out.ChunkContents = direct.Slice_ToProto(mapCtx, in.ChunkContents, Answer_Reference_UnstructuredDocumentInfo_ChunkContent_ToProto)
	out.StructData = StructData_ToProto(mapCtx, in.StructData)
	return out
}
func Answer_Reference_UnstructuredDocumentInfo_ChunkContent_FromProto(mapCtx *direct.MapContext, in *pb.Answer_Reference_UnstructuredDocumentInfo_ChunkContent) *krm.Answer_Reference_UnstructuredDocumentInfo_ChunkContent {
	if in == nil {
		return nil
	}
	out := &krm.Answer_Reference_UnstructuredDocumentInfo_ChunkContent{}
	out.Content = direct.LazyPtr(in.GetContent())
	out.PageIdentifier = direct.LazyPtr(in.GetPageIdentifier())
	out.RelevanceScore = in.RelevanceScore
	return out
}
func Answer_Reference_UnstructuredDocumentInfo_ChunkContent_ToProto(mapCtx *direct.MapContext, in *krm.Answer_Reference_UnstructuredDocumentInfo_ChunkContent) *pb.Answer_Reference_UnstructuredDocumentInfo_ChunkContent {
	if in == nil {
		return nil
	}
	out := &pb.Answer_Reference_UnstructuredDocumentInfo_ChunkContent{}
	out.Content = direct.ValueOf(in.Content)
	out.PageIdentifier = direct.ValueOf(in.PageIdentifier)
	out.RelevanceScore = in.RelevanceScore
	return out
}
func Answer_Step_FromProto(mapCtx *direct.MapContext, in *pb.Answer_Step) *krm.Answer_Step {
	if in == nil {
		return nil
	}
	out := &krm.Answer_Step{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Thought = direct.LazyPtr(in.GetThought())
	out.Actions = direct.Slice_FromProto(mapCtx, in.Actions, Answer_Step_Action_FromProto)
	return out
}
func Answer_Step_ToProto(mapCtx *direct.MapContext, in *krm.Answer_Step) *pb.Answer_Step {
	if in == nil {
		return nil
	}
	out := &pb.Answer_Step{}
	out.State = direct.Enum_ToProto[pb.Answer_Step_State](mapCtx, in.State)
	out.Description = direct.ValueOf(in.Description)
	out.Thought = direct.ValueOf(in.Thought)
	out.Actions = direct.Slice_ToProto(mapCtx, in.Actions, Answer_Step_Action_ToProto)
	return out
}
func Answer_Step_Action_FromProto(mapCtx *direct.MapContext, in *pb.Answer_Step_Action) *krm.Answer_Step_Action {
	if in == nil {
		return nil
	}
	out := &krm.Answer_Step_Action{}
	out.SearchAction = Answer_Step_Action_SearchAction_FromProto(mapCtx, in.GetSearchAction())
	out.Observation = Answer_Step_Action_Observation_FromProto(mapCtx, in.GetObservation())
	return out
}
func Answer_Step_Action_ToProto(mapCtx *direct.MapContext, in *krm.Answer_Step_Action) *pb.Answer_Step_Action {
	if in == nil {
		return nil
	}
	out := &pb.Answer_Step_Action{}
	if oneof := Answer_Step_Action_SearchAction_ToProto(mapCtx, in.SearchAction); oneof != nil {
		out.Action = &pb.Answer_Step_Action_SearchAction_{SearchAction: oneof}
	}
	out.Observation = Answer_Step_Action_Observation_ToProto(mapCtx, in.Observation)
	return out
}
func Answer_Step_Action_Observation_FromProto(mapCtx *direct.MapContext, in *pb.Answer_Step_Action_Observation) *krm.Answer_Step_Action_Observation {
	if in == nil {
		return nil
	}
	out := &krm.Answer_Step_Action_Observation{}
	out.SearchResults = direct.Slice_FromProto(mapCtx, in.SearchResults, Answer_Step_Action_Observation_SearchResult_FromProto)
	return out
}
func Answer_Step_Action_Observation_ToProto(mapCtx *direct.MapContext, in *krm.Answer_Step_Action_Observation) *pb.Answer_Step_Action_Observation {
	if in == nil {
		return nil
	}
	out := &pb.Answer_Step_Action_Observation{}
	out.SearchResults = direct.Slice_ToProto(mapCtx, in.SearchResults, Answer_Step_Action_Observation_SearchResult_ToProto)
	return out
}
func Answer_Step_Action_Observation_SearchResult_FromProto(mapCtx *direct.MapContext, in *pb.Answer_Step_Action_Observation_SearchResult) *krm.Answer_Step_Action_Observation_SearchResult {
	if in == nil {
		return nil
	}
	out := &krm.Answer_Step_Action_Observation_SearchResult{}
	out.Document = direct.LazyPtr(in.GetDocument())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.SnippetInfo = direct.Slice_FromProto(mapCtx, in.SnippetInfo, Answer_Step_Action_Observation_SearchResult_SnippetInfo_FromProto)
	out.ChunkInfo = direct.Slice_FromProto(mapCtx, in.ChunkInfo, Answer_Step_Action_Observation_SearchResult_ChunkInfo_FromProto)
	out.StructData = StructData_FromProto(mapCtx, in.GetStructData())
	return out
}
func Answer_Step_Action_Observation_SearchResult_ToProto(mapCtx *direct.MapContext, in *krm.Answer_Step_Action_Observation_SearchResult) *pb.Answer_Step_Action_Observation_SearchResult {
	if in == nil {
		return nil
	}
	out := &pb.Answer_Step_Action_Observation_SearchResult{}
	out.Document = direct.ValueOf(in.Document)
	out.Uri = direct.ValueOf(in.URI)
	out.Title = direct.ValueOf(in.Title)
	out.SnippetInfo = direct.Slice_ToProto(mapCtx, in.SnippetInfo, Answer_Step_Action_Observation_SearchResult_SnippetInfo_ToProto)
	out.ChunkInfo = direct.Slice_ToProto(mapCtx, in.ChunkInfo, Answer_Step_Action_Observation_SearchResult_ChunkInfo_ToProto)
	out.StructData = StructData_ToProto(mapCtx, in.StructData)
	return out
}
func Answer_Step_Action_Observation_SearchResult_ChunkInfo_FromProto(mapCtx *direct.MapContext, in *pb.Answer_Step_Action_Observation_SearchResult_ChunkInfo) *krm.Answer_Step_Action_Observation_SearchResult_ChunkInfo {
	if in == nil {
		return nil
	}
	out := &krm.Answer_Step_Action_Observation_SearchResult_ChunkInfo{}
	out.Chunk = direct.LazyPtr(in.GetChunk())
	out.Content = direct.LazyPtr(in.GetContent())
	out.RelevanceScore = in.RelevanceScore
	return out
}
func Answer_Step_Action_Observation_SearchResult_ChunkInfo_ToProto(mapCtx *direct.MapContext, in *krm.Answer_Step_Action_Observation_SearchResult_ChunkInfo) *pb.Answer_Step_Action_Observation_SearchResult_ChunkInfo {
	if in == nil {
		return nil
	}
	out := &pb.Answer_Step_Action_Observation_SearchResult_ChunkInfo{}
	out.Chunk = direct.ValueOf(in.Chunk)
	out.Content = direct.ValueOf(in.Content)
	out.RelevanceScore = in.RelevanceScore
	return out
}
func Answer_Step_Action_Observation_SearchResult_SnippetInfo_FromProto(mapCtx *direct.MapContext, in *pb.Answer_Step_Action_Observation_SearchResult_SnippetInfo) *krm.Answer_Step_Action_Observation_SearchResult_SnippetInfo {
	if in == nil {
		return nil
	}
	out := &krm.Answer_Step_Action_Observation_SearchResult_SnippetInfo{}
	out.Snippet = direct.LazyPtr(in.GetSnippet())
	out.SnippetStatus = direct.LazyPtr(in.GetSnippetStatus())
	return out
}
func Answer_Step_Action_Observation_SearchResult_SnippetInfo_ToProto(mapCtx *direct.MapContext, in *krm.Answer_Step_Action_Observation_SearchResult_SnippetInfo) *pb.Answer_Step_Action_Observation_SearchResult_SnippetInfo {
	if in == nil {
		return nil
	}
	out := &pb.Answer_Step_Action_Observation_SearchResult_SnippetInfo{}
	out.Snippet = direct.ValueOf(in.Snippet)
	out.SnippetStatus = direct.ValueOf(in.SnippetStatus)
	return out
}
func Answer_Step_Action_SearchAction_FromProto(mapCtx *direct.MapContext, in *pb.Answer_Step_Action_SearchAction) *krm.Answer_Step_Action_SearchAction {
	if in == nil {
		return nil
	}
	out := &krm.Answer_Step_Action_SearchAction{}
	out.Query = direct.LazyPtr(in.GetQuery())
	return out
}
func Answer_Step_Action_SearchAction_ToProto(mapCtx *direct.MapContext, in *krm.Answer_Step_Action_SearchAction) *pb.Answer_Step_Action_SearchAction {
	if in == nil {
		return nil
	}
	out := &pb.Answer_Step_Action_SearchAction{}
	out.Query = direct.ValueOf(in.Query)
	return out
}
