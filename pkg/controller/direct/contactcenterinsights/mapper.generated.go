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

package contactcenterinsights

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Analysis_FromProto(mapCtx *direct.MapContext, in *pb.Analysis) *krm.Analysis {
	if in == nil {
		return nil
	}
	out := &krm.Analysis{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: RequestTime
	// MISSING: CreateTime
	// MISSING: AnalysisResult
	out.AnnotatorSelector = AnnotatorSelector_FromProto(mapCtx, in.GetAnnotatorSelector())
	return out
}
func Analysis_ToProto(mapCtx *direct.MapContext, in *krm.Analysis) *pb.Analysis {
	if in == nil {
		return nil
	}
	out := &pb.Analysis{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: RequestTime
	// MISSING: CreateTime
	// MISSING: AnalysisResult
	out.AnnotatorSelector = AnnotatorSelector_ToProto(mapCtx, in.AnnotatorSelector)
	return out
}
func AnalysisObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Analysis) *krm.AnalysisObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalysisObservedState{}
	// MISSING: Name
	out.RequestTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRequestTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.AnalysisResult = AnalysisResult_FromProto(mapCtx, in.GetAnalysisResult())
	// MISSING: AnnotatorSelector
	return out
}
func AnalysisObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalysisObservedState) *pb.Analysis {
	if in == nil {
		return nil
	}
	out := &pb.Analysis{}
	// MISSING: Name
	out.RequestTime = direct.StringTimestamp_ToProto(mapCtx, in.RequestTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.AnalysisResult = AnalysisResult_ToProto(mapCtx, in.AnalysisResult)
	// MISSING: AnnotatorSelector
	return out
}
func AnalysisResult_FromProto(mapCtx *direct.MapContext, in *pb.AnalysisResult) *krm.AnalysisResult {
	if in == nil {
		return nil
	}
	out := &krm.AnalysisResult{}
	out.CallAnalysisMetadata = AnalysisResult_CallAnalysisMetadata_FromProto(mapCtx, in.GetCallAnalysisMetadata())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func AnalysisResult_ToProto(mapCtx *direct.MapContext, in *krm.AnalysisResult) *pb.AnalysisResult {
	if in == nil {
		return nil
	}
	out := &pb.AnalysisResult{}
	if oneof := AnalysisResult_CallAnalysisMetadata_ToProto(mapCtx, in.CallAnalysisMetadata); oneof != nil {
		out.Metadata = &pb.AnalysisResult_CallAnalysisMetadata_{CallAnalysisMetadata: oneof}
	}
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func AnalysisResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AnalysisResult) *krm.AnalysisResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalysisResultObservedState{}
	out.CallAnalysisMetadata = AnalysisResult_CallAnalysisMetadataObservedState_FromProto(mapCtx, in.GetCallAnalysisMetadata())
	// MISSING: EndTime
	return out
}
func AnalysisResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalysisResultObservedState) *pb.AnalysisResult {
	if in == nil {
		return nil
	}
	out := &pb.AnalysisResult{}
	if oneof := AnalysisResult_CallAnalysisMetadataObservedState_ToProto(mapCtx, in.CallAnalysisMetadata); oneof != nil {
		out.Metadata = &pb.AnalysisResult_CallAnalysisMetadata_{CallAnalysisMetadata: oneof}
	}
	// MISSING: EndTime
	return out
}
func AnnotationBoundary_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationBoundary) *krm.AnnotationBoundary {
	if in == nil {
		return nil
	}
	out := &krm.AnnotationBoundary{}
	out.WordIndex = direct.LazyPtr(in.GetWordIndex())
	out.TranscriptIndex = direct.LazyPtr(in.GetTranscriptIndex())
	return out
}
func AnnotationBoundary_ToProto(mapCtx *direct.MapContext, in *krm.AnnotationBoundary) *pb.AnnotationBoundary {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationBoundary{}
	if oneof := AnnotationBoundary_WordIndex_ToProto(mapCtx, in.WordIndex); oneof != nil {
		out.DetailedBoundary = oneof
	}
	out.TranscriptIndex = direct.ValueOf(in.TranscriptIndex)
	return out
}
func AnnotatorSelector_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatorSelector) *krm.AnnotatorSelector {
	if in == nil {
		return nil
	}
	out := &krm.AnnotatorSelector{}
	out.RunInterruptionAnnotator = direct.LazyPtr(in.GetRunInterruptionAnnotator())
	out.RunSilenceAnnotator = direct.LazyPtr(in.GetRunSilenceAnnotator())
	out.RunPhraseMatcherAnnotator = direct.LazyPtr(in.GetRunPhraseMatcherAnnotator())
	out.PhraseMatchers = in.PhraseMatchers
	out.RunSentimentAnnotator = direct.LazyPtr(in.GetRunSentimentAnnotator())
	out.RunEntityAnnotator = direct.LazyPtr(in.GetRunEntityAnnotator())
	out.RunIntentAnnotator = direct.LazyPtr(in.GetRunIntentAnnotator())
	out.RunIssueModelAnnotator = direct.LazyPtr(in.GetRunIssueModelAnnotator())
	out.IssueModels = in.IssueModels
	out.RunSummarizationAnnotator = direct.LazyPtr(in.GetRunSummarizationAnnotator())
	out.SummarizationConfig = AnnotatorSelector_SummarizationConfig_FromProto(mapCtx, in.GetSummarizationConfig())
	out.RunQaAnnotator = direct.LazyPtr(in.GetRunQaAnnotator())
	out.QaConfig = AnnotatorSelector_QaConfig_FromProto(mapCtx, in.GetQaConfig())
	return out
}
func AnnotatorSelector_ToProto(mapCtx *direct.MapContext, in *krm.AnnotatorSelector) *pb.AnnotatorSelector {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatorSelector{}
	out.RunInterruptionAnnotator = direct.ValueOf(in.RunInterruptionAnnotator)
	out.RunSilenceAnnotator = direct.ValueOf(in.RunSilenceAnnotator)
	out.RunPhraseMatcherAnnotator = direct.ValueOf(in.RunPhraseMatcherAnnotator)
	out.PhraseMatchers = in.PhraseMatchers
	out.RunSentimentAnnotator = direct.ValueOf(in.RunSentimentAnnotator)
	out.RunEntityAnnotator = direct.ValueOf(in.RunEntityAnnotator)
	out.RunIntentAnnotator = direct.ValueOf(in.RunIntentAnnotator)
	out.RunIssueModelAnnotator = direct.ValueOf(in.RunIssueModelAnnotator)
	out.IssueModels = in.IssueModels
	out.RunSummarizationAnnotator = direct.ValueOf(in.RunSummarizationAnnotator)
	out.SummarizationConfig = AnnotatorSelector_SummarizationConfig_ToProto(mapCtx, in.SummarizationConfig)
	out.RunQaAnnotator = direct.ValueOf(in.RunQaAnnotator)
	out.QaConfig = AnnotatorSelector_QaConfig_ToProto(mapCtx, in.QaConfig)
	return out
}
func AnnotatorSelector_QaConfig_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatorSelector_QaConfig) *krm.AnnotatorSelector_QaConfig {
	if in == nil {
		return nil
	}
	out := &krm.AnnotatorSelector_QaConfig{}
	out.ScorecardList = AnnotatorSelector_QaConfig_ScorecardList_FromProto(mapCtx, in.GetScorecardList())
	return out
}
func AnnotatorSelector_QaConfig_ToProto(mapCtx *direct.MapContext, in *krm.AnnotatorSelector_QaConfig) *pb.AnnotatorSelector_QaConfig {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatorSelector_QaConfig{}
	if oneof := AnnotatorSelector_QaConfig_ScorecardList_ToProto(mapCtx, in.ScorecardList); oneof != nil {
		out.ScorecardSource = &pb.AnnotatorSelector_QaConfig_ScorecardList_{ScorecardList: oneof}
	}
	return out
}
func AnnotatorSelector_QaConfig_ScorecardList_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatorSelector_QaConfig_ScorecardList) *krm.AnnotatorSelector_QaConfig_ScorecardList {
	if in == nil {
		return nil
	}
	out := &krm.AnnotatorSelector_QaConfig_ScorecardList{}
	out.QaScorecardRevisions = in.QaScorecardRevisions
	return out
}
func AnnotatorSelector_QaConfig_ScorecardList_ToProto(mapCtx *direct.MapContext, in *krm.AnnotatorSelector_QaConfig_ScorecardList) *pb.AnnotatorSelector_QaConfig_ScorecardList {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatorSelector_QaConfig_ScorecardList{}
	out.QaScorecardRevisions = in.QaScorecardRevisions
	return out
}
func AnnotatorSelector_SummarizationConfig_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatorSelector_SummarizationConfig) *krm.AnnotatorSelector_SummarizationConfig {
	if in == nil {
		return nil
	}
	out := &krm.AnnotatorSelector_SummarizationConfig{}
	out.ConversationProfile = direct.LazyPtr(in.GetConversationProfile())
	out.SummarizationModel = direct.Enum_FromProto(mapCtx, in.GetSummarizationModel())
	return out
}
func AnnotatorSelector_SummarizationConfig_ToProto(mapCtx *direct.MapContext, in *krm.AnnotatorSelector_SummarizationConfig) *pb.AnnotatorSelector_SummarizationConfig {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatorSelector_SummarizationConfig{}
	if oneof := AnnotatorSelector_SummarizationConfig_ConversationProfile_ToProto(mapCtx, in.ConversationProfile); oneof != nil {
		out.ModelSource = oneof
	}
	if oneof := AnnotatorSelector_SummarizationConfig_SummarizationModel_ToProto(mapCtx, in.SummarizationModel); oneof != nil {
		out.ModelSource = oneof
	}
	return out
}
func CallAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.CallAnnotation) *krm.CallAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.CallAnnotation{}
	out.InterruptionData = InterruptionData_FromProto(mapCtx, in.GetInterruptionData())
	out.SentimentData = SentimentData_FromProto(mapCtx, in.GetSentimentData())
	out.SilenceData = SilenceData_FromProto(mapCtx, in.GetSilenceData())
	out.HoldData = HoldData_FromProto(mapCtx, in.GetHoldData())
	out.EntityMentionData = EntityMentionData_FromProto(mapCtx, in.GetEntityMentionData())
	out.IntentMatchData = IntentMatchData_FromProto(mapCtx, in.GetIntentMatchData())
	out.PhraseMatchData = PhraseMatchData_FromProto(mapCtx, in.GetPhraseMatchData())
	out.IssueMatchData = IssueMatchData_FromProto(mapCtx, in.GetIssueMatchData())
	out.ChannelTag = direct.LazyPtr(in.GetChannelTag())
	out.AnnotationStartBoundary = AnnotationBoundary_FromProto(mapCtx, in.GetAnnotationStartBoundary())
	out.AnnotationEndBoundary = AnnotationBoundary_FromProto(mapCtx, in.GetAnnotationEndBoundary())
	return out
}
func CallAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.CallAnnotation) *pb.CallAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.CallAnnotation{}
	if oneof := InterruptionData_ToProto(mapCtx, in.InterruptionData); oneof != nil {
		out.Data = &pb.CallAnnotation_InterruptionData{InterruptionData: oneof}
	}
	if oneof := SentimentData_ToProto(mapCtx, in.SentimentData); oneof != nil {
		out.Data = &pb.CallAnnotation_SentimentData{SentimentData: oneof}
	}
	if oneof := SilenceData_ToProto(mapCtx, in.SilenceData); oneof != nil {
		out.Data = &pb.CallAnnotation_SilenceData{SilenceData: oneof}
	}
	if oneof := HoldData_ToProto(mapCtx, in.HoldData); oneof != nil {
		out.Data = &pb.CallAnnotation_HoldData{HoldData: oneof}
	}
	if oneof := EntityMentionData_ToProto(mapCtx, in.EntityMentionData); oneof != nil {
		out.Data = &pb.CallAnnotation_EntityMentionData{EntityMentionData: oneof}
	}
	if oneof := IntentMatchData_ToProto(mapCtx, in.IntentMatchData); oneof != nil {
		out.Data = &pb.CallAnnotation_IntentMatchData{IntentMatchData: oneof}
	}
	if oneof := PhraseMatchData_ToProto(mapCtx, in.PhraseMatchData); oneof != nil {
		out.Data = &pb.CallAnnotation_PhraseMatchData{PhraseMatchData: oneof}
	}
	if oneof := IssueMatchData_ToProto(mapCtx, in.IssueMatchData); oneof != nil {
		out.Data = &pb.CallAnnotation_IssueMatchData{IssueMatchData: oneof}
	}
	out.ChannelTag = direct.ValueOf(in.ChannelTag)
	out.AnnotationStartBoundary = AnnotationBoundary_ToProto(mapCtx, in.AnnotationStartBoundary)
	out.AnnotationEndBoundary = AnnotationBoundary_ToProto(mapCtx, in.AnnotationEndBoundary)
	return out
}
func ContactcenterinsightsAnalysisObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Analysis) *krm.ContactcenterinsightsAnalysisObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsAnalysisObservedState{}
	// MISSING: Name
	// MISSING: RequestTime
	// MISSING: CreateTime
	// MISSING: AnalysisResult
	// MISSING: AnnotatorSelector
	return out
}
func ContactcenterinsightsAnalysisObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsAnalysisObservedState) *pb.Analysis {
	if in == nil {
		return nil
	}
	out := &pb.Analysis{}
	// MISSING: Name
	// MISSING: RequestTime
	// MISSING: CreateTime
	// MISSING: AnalysisResult
	// MISSING: AnnotatorSelector
	return out
}
func ContactcenterinsightsAnalysisSpec_FromProto(mapCtx *direct.MapContext, in *pb.Analysis) *krm.ContactcenterinsightsAnalysisSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsAnalysisSpec{}
	// MISSING: Name
	// MISSING: RequestTime
	// MISSING: CreateTime
	// MISSING: AnalysisResult
	// MISSING: AnnotatorSelector
	return out
}
func ContactcenterinsightsAnalysisSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsAnalysisSpec) *pb.Analysis {
	if in == nil {
		return nil
	}
	out := &pb.Analysis{}
	// MISSING: Name
	// MISSING: RequestTime
	// MISSING: CreateTime
	// MISSING: AnalysisResult
	// MISSING: AnnotatorSelector
	return out
}
func ConversationLevelSentiment_FromProto(mapCtx *direct.MapContext, in *pb.ConversationLevelSentiment) *krm.ConversationLevelSentiment {
	if in == nil {
		return nil
	}
	out := &krm.ConversationLevelSentiment{}
	out.ChannelTag = direct.LazyPtr(in.GetChannelTag())
	out.SentimentData = SentimentData_FromProto(mapCtx, in.GetSentimentData())
	return out
}
func ConversationLevelSentiment_ToProto(mapCtx *direct.MapContext, in *krm.ConversationLevelSentiment) *pb.ConversationLevelSentiment {
	if in == nil {
		return nil
	}
	out := &pb.ConversationLevelSentiment{}
	out.ChannelTag = direct.ValueOf(in.ChannelTag)
	out.SentimentData = SentimentData_ToProto(mapCtx, in.SentimentData)
	return out
}
func ConversationLevelSilence_FromProto(mapCtx *direct.MapContext, in *pb.ConversationLevelSilence) *krm.ConversationLevelSilence {
	if in == nil {
		return nil
	}
	out := &krm.ConversationLevelSilence{}
	out.SilenceDuration = direct.StringDuration_FromProto(mapCtx, in.GetSilenceDuration())
	out.SilencePercentage = direct.LazyPtr(in.GetSilencePercentage())
	return out
}
func ConversationLevelSilence_ToProto(mapCtx *direct.MapContext, in *krm.ConversationLevelSilence) *pb.ConversationLevelSilence {
	if in == nil {
		return nil
	}
	out := &pb.ConversationLevelSilence{}
	out.SilenceDuration = direct.StringDuration_ToProto(mapCtx, in.SilenceDuration)
	out.SilencePercentage = direct.ValueOf(in.SilencePercentage)
	return out
}
func Entity_FromProto(mapCtx *direct.MapContext, in *pb.Entity) *krm.Entity {
	if in == nil {
		return nil
	}
	out := &krm.Entity{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Metadata = in.Metadata
	out.Salience = direct.LazyPtr(in.GetSalience())
	out.Sentiment = SentimentData_FromProto(mapCtx, in.GetSentiment())
	return out
}
func Entity_ToProto(mapCtx *direct.MapContext, in *krm.Entity) *pb.Entity {
	if in == nil {
		return nil
	}
	out := &pb.Entity{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Type = direct.Enum_ToProto[pb.Entity_Type](mapCtx, in.Type)
	out.Metadata = in.Metadata
	out.Salience = direct.ValueOf(in.Salience)
	out.Sentiment = SentimentData_ToProto(mapCtx, in.Sentiment)
	return out
}
func EntityMentionData_FromProto(mapCtx *direct.MapContext, in *pb.EntityMentionData) *krm.EntityMentionData {
	if in == nil {
		return nil
	}
	out := &krm.EntityMentionData{}
	out.EntityUniqueID = direct.LazyPtr(in.GetEntityUniqueId())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Sentiment = SentimentData_FromProto(mapCtx, in.GetSentiment())
	return out
}
func EntityMentionData_ToProto(mapCtx *direct.MapContext, in *krm.EntityMentionData) *pb.EntityMentionData {
	if in == nil {
		return nil
	}
	out := &pb.EntityMentionData{}
	out.EntityUniqueId = direct.ValueOf(in.EntityUniqueID)
	out.Type = direct.Enum_ToProto[pb.EntityMentionData_MentionType](mapCtx, in.Type)
	out.Sentiment = SentimentData_ToProto(mapCtx, in.Sentiment)
	return out
}
func HoldData_FromProto(mapCtx *direct.MapContext, in *pb.HoldData) *krm.HoldData {
	if in == nil {
		return nil
	}
	out := &krm.HoldData{}
	return out
}
func HoldData_ToProto(mapCtx *direct.MapContext, in *krm.HoldData) *pb.HoldData {
	if in == nil {
		return nil
	}
	out := &pb.HoldData{}
	return out
}
func Intent_FromProto(mapCtx *direct.MapContext, in *pb.Intent) *krm.Intent {
	if in == nil {
		return nil
	}
	out := &krm.Intent{}
	out.ID = direct.LazyPtr(in.GetId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func Intent_ToProto(mapCtx *direct.MapContext, in *krm.Intent) *pb.Intent {
	if in == nil {
		return nil
	}
	out := &pb.Intent{}
	out.Id = direct.ValueOf(in.ID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func IntentMatchData_FromProto(mapCtx *direct.MapContext, in *pb.IntentMatchData) *krm.IntentMatchData {
	if in == nil {
		return nil
	}
	out := &krm.IntentMatchData{}
	out.IntentUniqueID = direct.LazyPtr(in.GetIntentUniqueId())
	return out
}
func IntentMatchData_ToProto(mapCtx *direct.MapContext, in *krm.IntentMatchData) *pb.IntentMatchData {
	if in == nil {
		return nil
	}
	out := &pb.IntentMatchData{}
	out.IntentUniqueId = direct.ValueOf(in.IntentUniqueID)
	return out
}
func InterruptionData_FromProto(mapCtx *direct.MapContext, in *pb.InterruptionData) *krm.InterruptionData {
	if in == nil {
		return nil
	}
	out := &krm.InterruptionData{}
	return out
}
func InterruptionData_ToProto(mapCtx *direct.MapContext, in *krm.InterruptionData) *pb.InterruptionData {
	if in == nil {
		return nil
	}
	out := &pb.InterruptionData{}
	return out
}
func IssueAssignment_FromProto(mapCtx *direct.MapContext, in *pb.IssueAssignment) *krm.IssueAssignment {
	if in == nil {
		return nil
	}
	out := &krm.IssueAssignment{}
	out.Issue = direct.LazyPtr(in.GetIssue())
	out.Score = direct.LazyPtr(in.GetScore())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func IssueAssignment_ToProto(mapCtx *direct.MapContext, in *krm.IssueAssignment) *pb.IssueAssignment {
	if in == nil {
		return nil
	}
	out := &pb.IssueAssignment{}
	out.Issue = direct.ValueOf(in.Issue)
	out.Score = direct.ValueOf(in.Score)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func IssueMatchData_FromProto(mapCtx *direct.MapContext, in *pb.IssueMatchData) *krm.IssueMatchData {
	if in == nil {
		return nil
	}
	out := &krm.IssueMatchData{}
	out.IssueAssignment = IssueAssignment_FromProto(mapCtx, in.GetIssueAssignment())
	return out
}
func IssueMatchData_ToProto(mapCtx *direct.MapContext, in *krm.IssueMatchData) *pb.IssueMatchData {
	if in == nil {
		return nil
	}
	out := &pb.IssueMatchData{}
	out.IssueAssignment = IssueAssignment_ToProto(mapCtx, in.IssueAssignment)
	return out
}
func IssueModelResult_FromProto(mapCtx *direct.MapContext, in *pb.IssueModelResult) *krm.IssueModelResult {
	if in == nil {
		return nil
	}
	out := &krm.IssueModelResult{}
	out.IssueModel = direct.LazyPtr(in.GetIssueModel())
	out.Issues = direct.Slice_FromProto(mapCtx, in.Issues, IssueAssignment_FromProto)
	return out
}
func IssueModelResult_ToProto(mapCtx *direct.MapContext, in *krm.IssueModelResult) *pb.IssueModelResult {
	if in == nil {
		return nil
	}
	out := &pb.IssueModelResult{}
	out.IssueModel = direct.ValueOf(in.IssueModel)
	out.Issues = direct.Slice_ToProto(mapCtx, in.Issues, IssueAssignment_ToProto)
	return out
}
func PhraseMatchData_FromProto(mapCtx *direct.MapContext, in *pb.PhraseMatchData) *krm.PhraseMatchData {
	if in == nil {
		return nil
	}
	out := &krm.PhraseMatchData{}
	out.PhraseMatcher = direct.LazyPtr(in.GetPhraseMatcher())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func PhraseMatchData_ToProto(mapCtx *direct.MapContext, in *krm.PhraseMatchData) *pb.PhraseMatchData {
	if in == nil {
		return nil
	}
	out := &pb.PhraseMatchData{}
	out.PhraseMatcher = direct.ValueOf(in.PhraseMatcher)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func QaAnswer_FromProto(mapCtx *direct.MapContext, in *pb.QaAnswer) *krm.QaAnswer {
	if in == nil {
		return nil
	}
	out := &krm.QaAnswer{}
	out.QaQuestion = direct.LazyPtr(in.GetQaQuestion())
	out.Conversation = direct.LazyPtr(in.GetConversation())
	out.QuestionBody = direct.LazyPtr(in.GetQuestionBody())
	out.AnswerValue = QaAnswer_AnswerValue_FromProto(mapCtx, in.GetAnswerValue())
	out.Tags = in.Tags
	out.AnswerSources = direct.Slice_FromProto(mapCtx, in.AnswerSources, QaAnswer_AnswerSource_FromProto)
	return out
}
func QaAnswer_ToProto(mapCtx *direct.MapContext, in *krm.QaAnswer) *pb.QaAnswer {
	if in == nil {
		return nil
	}
	out := &pb.QaAnswer{}
	out.QaQuestion = direct.ValueOf(in.QaQuestion)
	out.Conversation = direct.ValueOf(in.Conversation)
	out.QuestionBody = direct.ValueOf(in.QuestionBody)
	out.AnswerValue = QaAnswer_AnswerValue_ToProto(mapCtx, in.AnswerValue)
	out.Tags = in.Tags
	out.AnswerSources = direct.Slice_ToProto(mapCtx, in.AnswerSources, QaAnswer_AnswerSource_ToProto)
	return out
}
func QaAnswerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QaAnswer) *krm.QaAnswerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QaAnswerObservedState{}
	// MISSING: QaQuestion
	// MISSING: Conversation
	// MISSING: QuestionBody
	out.AnswerValue = QaAnswer_AnswerValueObservedState_FromProto(mapCtx, in.GetAnswerValue())
	// MISSING: Tags
	// MISSING: AnswerSources
	return out
}
func QaAnswerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QaAnswerObservedState) *pb.QaAnswer {
	if in == nil {
		return nil
	}
	out := &pb.QaAnswer{}
	// MISSING: QaQuestion
	// MISSING: Conversation
	// MISSING: QuestionBody
	out.AnswerValue = QaAnswer_AnswerValueObservedState_ToProto(mapCtx, in.AnswerValue)
	// MISSING: Tags
	// MISSING: AnswerSources
	return out
}
func QaAnswer_AnswerSource_FromProto(mapCtx *direct.MapContext, in *pb.QaAnswer_AnswerSource) *krm.QaAnswer_AnswerSource {
	if in == nil {
		return nil
	}
	out := &krm.QaAnswer_AnswerSource{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.AnswerValue = QaAnswer_AnswerValue_FromProto(mapCtx, in.GetAnswerValue())
	return out
}
func QaAnswer_AnswerSource_ToProto(mapCtx *direct.MapContext, in *krm.QaAnswer_AnswerSource) *pb.QaAnswer_AnswerSource {
	if in == nil {
		return nil
	}
	out := &pb.QaAnswer_AnswerSource{}
	out.SourceType = direct.Enum_ToProto[pb.QaAnswer_AnswerSource_SourceType](mapCtx, in.SourceType)
	out.AnswerValue = QaAnswer_AnswerValue_ToProto(mapCtx, in.AnswerValue)
	return out
}
func QaAnswer_AnswerValue_FromProto(mapCtx *direct.MapContext, in *pb.QaAnswer_AnswerValue) *krm.QaAnswer_AnswerValue {
	if in == nil {
		return nil
	}
	out := &krm.QaAnswer_AnswerValue{}
	out.StrValue = direct.LazyPtr(in.GetStrValue())
	out.NumValue = direct.LazyPtr(in.GetNumValue())
	out.BoolValue = direct.LazyPtr(in.GetBoolValue())
	out.NaValue = direct.LazyPtr(in.GetNaValue())
	out.Key = direct.LazyPtr(in.GetKey())
	// MISSING: Score
	// MISSING: PotentialScore
	// MISSING: NormalizedScore
	return out
}
func QaAnswer_AnswerValue_ToProto(mapCtx *direct.MapContext, in *krm.QaAnswer_AnswerValue) *pb.QaAnswer_AnswerValue {
	if in == nil {
		return nil
	}
	out := &pb.QaAnswer_AnswerValue{}
	if oneof := QaAnswer_AnswerValue_StrValue_ToProto(mapCtx, in.StrValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := QaAnswer_AnswerValue_NumValue_ToProto(mapCtx, in.NumValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := QaAnswer_AnswerValue_BoolValue_ToProto(mapCtx, in.BoolValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := QaAnswer_AnswerValue_NaValue_ToProto(mapCtx, in.NaValue); oneof != nil {
		out.Value = oneof
	}
	out.Key = direct.ValueOf(in.Key)
	// MISSING: Score
	// MISSING: PotentialScore
	// MISSING: NormalizedScore
	return out
}
func QaAnswer_AnswerValueObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QaAnswer_AnswerValue) *krm.QaAnswer_AnswerValueObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QaAnswer_AnswerValueObservedState{}
	// MISSING: StrValue
	// MISSING: NumValue
	// MISSING: BoolValue
	// MISSING: NaValue
	// MISSING: Key
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	return out
}
func QaAnswer_AnswerValueObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QaAnswer_AnswerValueObservedState) *pb.QaAnswer_AnswerValue {
	if in == nil {
		return nil
	}
	out := &pb.QaAnswer_AnswerValue{}
	// MISSING: StrValue
	// MISSING: NumValue
	// MISSING: BoolValue
	// MISSING: NaValue
	// MISSING: Key
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	return out
}
func QaScorecardResult_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecardResult) *krm.QaScorecardResult {
	if in == nil {
		return nil
	}
	out := &krm.QaScorecardResult{}
	out.Name = direct.LazyPtr(in.GetName())
	out.QaScorecardRevision = direct.LazyPtr(in.GetQaScorecardRevision())
	out.Conversation = direct.LazyPtr(in.GetConversation())
	// MISSING: CreateTime
	out.AgentID = direct.LazyPtr(in.GetAgentId())
	out.QaAnswers = direct.Slice_FromProto(mapCtx, in.QaAnswers, QaAnswer_FromProto)
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	out.QaTagResults = direct.Slice_FromProto(mapCtx, in.QaTagResults, QaScorecardResult_QaTagResult_FromProto)
	out.ScoreSources = direct.Slice_FromProto(mapCtx, in.ScoreSources, QaScorecardResult_ScoreSource_FromProto)
	return out
}
func QaScorecardResult_ToProto(mapCtx *direct.MapContext, in *krm.QaScorecardResult) *pb.QaScorecardResult {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecardResult{}
	out.Name = direct.ValueOf(in.Name)
	out.QaScorecardRevision = direct.ValueOf(in.QaScorecardRevision)
	out.Conversation = direct.ValueOf(in.Conversation)
	// MISSING: CreateTime
	out.AgentId = direct.ValueOf(in.AgentID)
	out.QaAnswers = direct.Slice_ToProto(mapCtx, in.QaAnswers, QaAnswer_ToProto)
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	out.QaTagResults = direct.Slice_ToProto(mapCtx, in.QaTagResults, QaScorecardResult_QaTagResult_ToProto)
	out.ScoreSources = direct.Slice_ToProto(mapCtx, in.ScoreSources, QaScorecardResult_ScoreSource_ToProto)
	return out
}
func QaScorecardResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecardResult) *krm.QaScorecardResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QaScorecardResultObservedState{}
	// MISSING: Name
	// MISSING: QaScorecardRevision
	// MISSING: Conversation
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: AgentID
	out.QaAnswers = direct.Slice_FromProto(mapCtx, in.QaAnswers, QaAnswerObservedState_FromProto)
	// MISSING: Score
	// MISSING: PotentialScore
	// MISSING: NormalizedScore
	// MISSING: QaTagResults
	// MISSING: ScoreSources
	return out
}
func QaScorecardResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QaScorecardResultObservedState) *pb.QaScorecardResult {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecardResult{}
	// MISSING: Name
	// MISSING: QaScorecardRevision
	// MISSING: Conversation
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: AgentID
	out.QaAnswers = direct.Slice_ToProto(mapCtx, in.QaAnswers, QaAnswerObservedState_ToProto)
	// MISSING: Score
	// MISSING: PotentialScore
	// MISSING: NormalizedScore
	// MISSING: QaTagResults
	// MISSING: ScoreSources
	return out
}
func QaScorecardResult_QaTagResult_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecardResult_QaTagResult) *krm.QaScorecardResult_QaTagResult {
	if in == nil {
		return nil
	}
	out := &krm.QaScorecardResult_QaTagResult{}
	out.Tag = direct.LazyPtr(in.GetTag())
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	return out
}
func QaScorecardResult_QaTagResult_ToProto(mapCtx *direct.MapContext, in *krm.QaScorecardResult_QaTagResult) *pb.QaScorecardResult_QaTagResult {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecardResult_QaTagResult{}
	out.Tag = direct.ValueOf(in.Tag)
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	return out
}
func QaScorecardResult_ScoreSource_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecardResult_ScoreSource) *krm.QaScorecardResult_ScoreSource {
	if in == nil {
		return nil
	}
	out := &krm.QaScorecardResult_ScoreSource{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	out.QaTagResults = direct.Slice_FromProto(mapCtx, in.QaTagResults, QaScorecardResult_QaTagResult_FromProto)
	return out
}
func QaScorecardResult_ScoreSource_ToProto(mapCtx *direct.MapContext, in *krm.QaScorecardResult_ScoreSource) *pb.QaScorecardResult_ScoreSource {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecardResult_ScoreSource{}
	out.SourceType = direct.Enum_ToProto[pb.QaScorecardResult_ScoreSource_SourceType](mapCtx, in.SourceType)
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	out.QaTagResults = direct.Slice_ToProto(mapCtx, in.QaTagResults, QaScorecardResult_QaTagResult_ToProto)
	return out
}
func SentimentData_FromProto(mapCtx *direct.MapContext, in *pb.SentimentData) *krm.SentimentData {
	if in == nil {
		return nil
	}
	out := &krm.SentimentData{}
	out.Magnitude = direct.LazyPtr(in.GetMagnitude())
	out.Score = direct.LazyPtr(in.GetScore())
	return out
}
func SentimentData_ToProto(mapCtx *direct.MapContext, in *krm.SentimentData) *pb.SentimentData {
	if in == nil {
		return nil
	}
	out := &pb.SentimentData{}
	out.Magnitude = direct.ValueOf(in.Magnitude)
	out.Score = direct.ValueOf(in.Score)
	return out
}
func SilenceData_FromProto(mapCtx *direct.MapContext, in *pb.SilenceData) *krm.SilenceData {
	if in == nil {
		return nil
	}
	out := &krm.SilenceData{}
	return out
}
func SilenceData_ToProto(mapCtx *direct.MapContext, in *krm.SilenceData) *pb.SilenceData {
	if in == nil {
		return nil
	}
	out := &pb.SilenceData{}
	return out
}
