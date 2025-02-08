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

package dialogflow

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AgentAssistantFeedback_FromProto(mapCtx *direct.MapContext, in *pb.AgentAssistantFeedback) *krm.AgentAssistantFeedback {
	if in == nil {
		return nil
	}
	out := &krm.AgentAssistantFeedback{}
	out.AnswerRelevance = direct.Enum_FromProto(mapCtx, in.GetAnswerRelevance())
	out.DocumentCorrectness = direct.Enum_FromProto(mapCtx, in.GetDocumentCorrectness())
	out.DocumentEfficiency = direct.Enum_FromProto(mapCtx, in.GetDocumentEfficiency())
	out.SummarizationFeedback = AgentAssistantFeedback_SummarizationFeedback_FromProto(mapCtx, in.GetSummarizationFeedback())
	out.KnowledgeSearchFeedback = AgentAssistantFeedback_KnowledgeSearchFeedback_FromProto(mapCtx, in.GetKnowledgeSearchFeedback())
	out.KnowledgeAssistFeedback = AgentAssistantFeedback_KnowledgeAssistFeedback_FromProto(mapCtx, in.GetKnowledgeAssistFeedback())
	return out
}
func AgentAssistantFeedback_ToProto(mapCtx *direct.MapContext, in *krm.AgentAssistantFeedback) *pb.AgentAssistantFeedback {
	if in == nil {
		return nil
	}
	out := &pb.AgentAssistantFeedback{}
	out.AnswerRelevance = direct.Enum_ToProto[pb.AgentAssistantFeedback_AnswerRelevance](mapCtx, in.AnswerRelevance)
	out.DocumentCorrectness = direct.Enum_ToProto[pb.AgentAssistantFeedback_DocumentCorrectness](mapCtx, in.DocumentCorrectness)
	out.DocumentEfficiency = direct.Enum_ToProto[pb.AgentAssistantFeedback_DocumentEfficiency](mapCtx, in.DocumentEfficiency)
	out.SummarizationFeedback = AgentAssistantFeedback_SummarizationFeedback_ToProto(mapCtx, in.SummarizationFeedback)
	out.KnowledgeSearchFeedback = AgentAssistantFeedback_KnowledgeSearchFeedback_ToProto(mapCtx, in.KnowledgeSearchFeedback)
	out.KnowledgeAssistFeedback = AgentAssistantFeedback_KnowledgeAssistFeedback_ToProto(mapCtx, in.KnowledgeAssistFeedback)
	return out
}
func AgentAssistantFeedback_KnowledgeAssistFeedback_FromProto(mapCtx *direct.MapContext, in *pb.AgentAssistantFeedback_KnowledgeAssistFeedback) *krm.AgentAssistantFeedback_KnowledgeAssistFeedback {
	if in == nil {
		return nil
	}
	out := &krm.AgentAssistantFeedback_KnowledgeAssistFeedback{}
	out.AnswerCopied = direct.LazyPtr(in.GetAnswerCopied())
	out.ClickedUris = in.ClickedUris
	return out
}
func AgentAssistantFeedback_KnowledgeAssistFeedback_ToProto(mapCtx *direct.MapContext, in *krm.AgentAssistantFeedback_KnowledgeAssistFeedback) *pb.AgentAssistantFeedback_KnowledgeAssistFeedback {
	if in == nil {
		return nil
	}
	out := &pb.AgentAssistantFeedback_KnowledgeAssistFeedback{}
	out.AnswerCopied = direct.ValueOf(in.AnswerCopied)
	out.ClickedUris = in.ClickedUris
	return out
}
func AgentAssistantFeedback_KnowledgeSearchFeedback_FromProto(mapCtx *direct.MapContext, in *pb.AgentAssistantFeedback_KnowledgeSearchFeedback) *krm.AgentAssistantFeedback_KnowledgeSearchFeedback {
	if in == nil {
		return nil
	}
	out := &krm.AgentAssistantFeedback_KnowledgeSearchFeedback{}
	out.AnswerCopied = direct.LazyPtr(in.GetAnswerCopied())
	out.ClickedUris = in.ClickedUris
	return out
}
func AgentAssistantFeedback_KnowledgeSearchFeedback_ToProto(mapCtx *direct.MapContext, in *krm.AgentAssistantFeedback_KnowledgeSearchFeedback) *pb.AgentAssistantFeedback_KnowledgeSearchFeedback {
	if in == nil {
		return nil
	}
	out := &pb.AgentAssistantFeedback_KnowledgeSearchFeedback{}
	out.AnswerCopied = direct.ValueOf(in.AnswerCopied)
	out.ClickedUris = in.ClickedUris
	return out
}
func AgentAssistantFeedback_SummarizationFeedback_FromProto(mapCtx *direct.MapContext, in *pb.AgentAssistantFeedback_SummarizationFeedback) *krm.AgentAssistantFeedback_SummarizationFeedback {
	if in == nil {
		return nil
	}
	out := &krm.AgentAssistantFeedback_SummarizationFeedback{}
	out.StartTimestamp = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTimestamp())
	out.SubmitTimestamp = direct.StringTimestamp_FromProto(mapCtx, in.GetSubmitTimestamp())
	out.SummaryText = direct.LazyPtr(in.GetSummaryText())
	out.TextSections = in.TextSections
	return out
}
func AgentAssistantFeedback_SummarizationFeedback_ToProto(mapCtx *direct.MapContext, in *krm.AgentAssistantFeedback_SummarizationFeedback) *pb.AgentAssistantFeedback_SummarizationFeedback {
	if in == nil {
		return nil
	}
	out := &pb.AgentAssistantFeedback_SummarizationFeedback{}
	out.StartTimestamp = direct.StringTimestamp_ToProto(mapCtx, in.StartTimestamp)
	out.SubmitTimestamp = direct.StringTimestamp_ToProto(mapCtx, in.SubmitTimestamp)
	out.SummaryText = direct.ValueOf(in.SummaryText)
	out.TextSections = in.TextSections
	return out
}
func AgentAssistantRecord_FromProto(mapCtx *direct.MapContext, in *pb.AgentAssistantRecord) *krm.AgentAssistantRecord {
	if in == nil {
		return nil
	}
	out := &krm.AgentAssistantRecord{}
	// MISSING: ArticleSuggestionAnswer
	// MISSING: FaqAnswer
	// MISSING: DialogflowAssistAnswer
	return out
}
func AgentAssistantRecord_ToProto(mapCtx *direct.MapContext, in *krm.AgentAssistantRecord) *pb.AgentAssistantRecord {
	if in == nil {
		return nil
	}
	out := &pb.AgentAssistantRecord{}
	// MISSING: ArticleSuggestionAnswer
	// MISSING: FaqAnswer
	// MISSING: DialogflowAssistAnswer
	return out
}
func AgentAssistantRecordObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AgentAssistantRecord) *krm.AgentAssistantRecordObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AgentAssistantRecordObservedState{}
	out.ArticleSuggestionAnswer = ArticleAnswer_FromProto(mapCtx, in.GetArticleSuggestionAnswer())
	out.FaqAnswer = FaqAnswer_FromProto(mapCtx, in.GetFaqAnswer())
	out.DialogflowAssistAnswer = DialogflowAssistAnswer_FromProto(mapCtx, in.GetDialogflowAssistAnswer())
	return out
}
func AgentAssistantRecordObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AgentAssistantRecordObservedState) *pb.AgentAssistantRecord {
	if in == nil {
		return nil
	}
	out := &pb.AgentAssistantRecord{}
	if oneof := ArticleAnswer_ToProto(mapCtx, in.ArticleSuggestionAnswer); oneof != nil {
		out.Answer = &pb.AgentAssistantRecord_ArticleSuggestionAnswer{ArticleSuggestionAnswer: oneof}
	}
	if oneof := FaqAnswer_ToProto(mapCtx, in.FaqAnswer); oneof != nil {
		out.Answer = &pb.AgentAssistantRecord_FaqAnswer{FaqAnswer: oneof}
	}
	if oneof := DialogflowAssistAnswer_ToProto(mapCtx, in.DialogflowAssistAnswer); oneof != nil {
		out.Answer = &pb.AgentAssistantRecord_DialogflowAssistAnswer{DialogflowAssistAnswer: oneof}
	}
	return out
}
func AnswerFeedback_FromProto(mapCtx *direct.MapContext, in *pb.AnswerFeedback) *krm.AnswerFeedback {
	if in == nil {
		return nil
	}
	out := &krm.AnswerFeedback{}
	out.CorrectnessLevel = direct.Enum_FromProto(mapCtx, in.GetCorrectnessLevel())
	out.AgentAssistantDetailFeedback = AgentAssistantFeedback_FromProto(mapCtx, in.GetAgentAssistantDetailFeedback())
	out.Clicked = direct.LazyPtr(in.GetClicked())
	out.ClickTime = direct.StringTimestamp_FromProto(mapCtx, in.GetClickTime())
	out.Displayed = direct.LazyPtr(in.GetDisplayed())
	out.DisplayTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDisplayTime())
	return out
}
func AnswerFeedback_ToProto(mapCtx *direct.MapContext, in *krm.AnswerFeedback) *pb.AnswerFeedback {
	if in == nil {
		return nil
	}
	out := &pb.AnswerFeedback{}
	out.CorrectnessLevel = direct.Enum_ToProto[pb.AnswerFeedback_CorrectnessLevel](mapCtx, in.CorrectnessLevel)
	if oneof := AgentAssistantFeedback_ToProto(mapCtx, in.AgentAssistantDetailFeedback); oneof != nil {
		out.DetailFeedback = &pb.AnswerFeedback_AgentAssistantDetailFeedback{AgentAssistantDetailFeedback: oneof}
	}
	out.Clicked = direct.ValueOf(in.Clicked)
	out.ClickTime = direct.StringTimestamp_ToProto(mapCtx, in.ClickTime)
	out.Displayed = direct.ValueOf(in.Displayed)
	out.DisplayTime = direct.StringTimestamp_ToProto(mapCtx, in.DisplayTime)
	return out
}
func AnswerRecord_FromProto(mapCtx *direct.MapContext, in *pb.AnswerRecord) *krm.AnswerRecord {
	if in == nil {
		return nil
	}
	out := &krm.AnswerRecord{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AnswerFeedback = AnswerFeedback_FromProto(mapCtx, in.GetAnswerFeedback())
	out.AgentAssistantRecord = AgentAssistantRecord_FromProto(mapCtx, in.GetAgentAssistantRecord())
	return out
}
func AnswerRecord_ToProto(mapCtx *direct.MapContext, in *krm.AnswerRecord) *pb.AnswerRecord {
	if in == nil {
		return nil
	}
	out := &pb.AnswerRecord{}
	out.Name = direct.ValueOf(in.Name)
	out.AnswerFeedback = AnswerFeedback_ToProto(mapCtx, in.AnswerFeedback)
	if oneof := AgentAssistantRecord_ToProto(mapCtx, in.AgentAssistantRecord); oneof != nil {
		out.Record = &pb.AnswerRecord_AgentAssistantRecord{AgentAssistantRecord: oneof}
	}
	return out
}
func AnswerRecordObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AnswerRecord) *krm.AnswerRecordObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnswerRecordObservedState{}
	// MISSING: Name
	// MISSING: AnswerFeedback
	out.AgentAssistantRecord = AgentAssistantRecordObservedState_FromProto(mapCtx, in.GetAgentAssistantRecord())
	return out
}
func AnswerRecordObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnswerRecordObservedState) *pb.AnswerRecord {
	if in == nil {
		return nil
	}
	out := &pb.AnswerRecord{}
	// MISSING: Name
	// MISSING: AnswerFeedback
	if oneof := AgentAssistantRecordObservedState_ToProto(mapCtx, in.AgentAssistantRecord); oneof != nil {
		out.Record = &pb.AnswerRecord_AgentAssistantRecord{AgentAssistantRecord: oneof}
	}
	return out
}
func ArticleAnswer_FromProto(mapCtx *direct.MapContext, in *pb.ArticleAnswer) *krm.ArticleAnswer {
	if in == nil {
		return nil
	}
	out := &krm.ArticleAnswer{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Snippets = in.Snippets
	out.Metadata = in.Metadata
	out.AnswerRecord = direct.LazyPtr(in.GetAnswerRecord())
	return out
}
func ArticleAnswer_ToProto(mapCtx *direct.MapContext, in *krm.ArticleAnswer) *pb.ArticleAnswer {
	if in == nil {
		return nil
	}
	out := &pb.ArticleAnswer{}
	out.Title = direct.ValueOf(in.Title)
	out.Uri = direct.ValueOf(in.URI)
	out.Snippets = in.Snippets
	out.Metadata = in.Metadata
	out.AnswerRecord = direct.ValueOf(in.AnswerRecord)
	return out
}
func Context_FromProto(mapCtx *direct.MapContext, in *pb.Context) *krm.Context {
	if in == nil {
		return nil
	}
	out := &krm.Context{}
	out.Name = direct.LazyPtr(in.GetName())
	out.LifespanCount = direct.LazyPtr(in.GetLifespanCount())
	out.Parameters = Parameters_FromProto(mapCtx, in.GetParameters())
	return out
}
func Context_ToProto(mapCtx *direct.MapContext, in *krm.Context) *pb.Context {
	if in == nil {
		return nil
	}
	out := &pb.Context{}
	out.Name = direct.ValueOf(in.Name)
	out.LifespanCount = direct.ValueOf(in.LifespanCount)
	out.Parameters = Parameters_ToProto(mapCtx, in.Parameters)
	return out
}
func DialogflowAssistAnswer_FromProto(mapCtx *direct.MapContext, in *pb.DialogflowAssistAnswer) *krm.DialogflowAssistAnswer {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowAssistAnswer{}
	out.QueryResult = QueryResult_FromProto(mapCtx, in.GetQueryResult())
	out.IntentSuggestion = IntentSuggestion_FromProto(mapCtx, in.GetIntentSuggestion())
	out.AnswerRecord = direct.LazyPtr(in.GetAnswerRecord())
	return out
}
func DialogflowAssistAnswer_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowAssistAnswer) *pb.DialogflowAssistAnswer {
	if in == nil {
		return nil
	}
	out := &pb.DialogflowAssistAnswer{}
	if oneof := QueryResult_ToProto(mapCtx, in.QueryResult); oneof != nil {
		out.Result = &pb.DialogflowAssistAnswer_QueryResult{QueryResult: oneof}
	}
	if oneof := IntentSuggestion_ToProto(mapCtx, in.IntentSuggestion); oneof != nil {
		out.Result = &pb.DialogflowAssistAnswer_IntentSuggestion{IntentSuggestion: oneof}
	}
	out.AnswerRecord = direct.ValueOf(in.AnswerRecord)
	return out
}
func DialogflowAssistAnswerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DialogflowAssistAnswer) *krm.DialogflowAssistAnswerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowAssistAnswerObservedState{}
	out.QueryResult = QueryResultObservedState_FromProto(mapCtx, in.GetQueryResult())
	// MISSING: IntentSuggestion
	// MISSING: AnswerRecord
	return out
}
func DialogflowAssistAnswerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowAssistAnswerObservedState) *pb.DialogflowAssistAnswer {
	if in == nil {
		return nil
	}
	out := &pb.DialogflowAssistAnswer{}
	if oneof := QueryResultObservedState_ToProto(mapCtx, in.QueryResult); oneof != nil {
		out.Result = &pb.DialogflowAssistAnswer_QueryResult{QueryResult: oneof}
	}
	// MISSING: IntentSuggestion
	// MISSING: AnswerRecord
	return out
}
func FaqAnswer_FromProto(mapCtx *direct.MapContext, in *pb.FaqAnswer) *krm.FaqAnswer {
	if in == nil {
		return nil
	}
	out := &krm.FaqAnswer{}
	out.Answer = direct.LazyPtr(in.GetAnswer())
	out.Confidence = direct.LazyPtr(in.GetConfidence())
	out.Question = direct.LazyPtr(in.GetQuestion())
	out.Source = direct.LazyPtr(in.GetSource())
	out.Metadata = in.Metadata
	out.AnswerRecord = direct.LazyPtr(in.GetAnswerRecord())
	return out
}
func FaqAnswer_ToProto(mapCtx *direct.MapContext, in *krm.FaqAnswer) *pb.FaqAnswer {
	if in == nil {
		return nil
	}
	out := &pb.FaqAnswer{}
	out.Answer = direct.ValueOf(in.Answer)
	out.Confidence = direct.ValueOf(in.Confidence)
	out.Question = direct.ValueOf(in.Question)
	out.Source = direct.ValueOf(in.Source)
	out.Metadata = in.Metadata
	out.AnswerRecord = direct.ValueOf(in.AnswerRecord)
	return out
}
func Intent_FromProto(mapCtx *direct.MapContext, in *pb.Intent) *krm.Intent {
	if in == nil {
		return nil
	}
	out := &krm.Intent{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.WebhookState = direct.Enum_FromProto(mapCtx, in.GetWebhookState())
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.IsFallback = direct.LazyPtr(in.GetIsFallback())
	out.MlEnabled = direct.LazyPtr(in.GetMlEnabled())
	out.MlDisabled = direct.LazyPtr(in.GetMlDisabled())
	out.LiveAgentHandoff = direct.LazyPtr(in.GetLiveAgentHandoff())
	out.EndInteraction = direct.LazyPtr(in.GetEndInteraction())
	out.InputContextNames = in.InputContextNames
	out.Events = in.Events
	out.TrainingPhrases = direct.Slice_FromProto(mapCtx, in.TrainingPhrases, Intent_TrainingPhrase_FromProto)
	out.Action = direct.LazyPtr(in.GetAction())
	out.OutputContexts = direct.Slice_FromProto(mapCtx, in.OutputContexts, Context_FromProto)
	out.ResetContexts = direct.LazyPtr(in.GetResetContexts())
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, Intent_Parameter_FromProto)
	out.Messages = direct.Slice_FromProto(mapCtx, in.Messages, Intent_Message_FromProto)
	out.DefaultResponsePlatforms = direct.EnumSlice_FromProto(mapCtx, in.DefaultResponsePlatforms)
	// MISSING: RootFollowupIntentName
	out.ParentFollowupIntentName = direct.LazyPtr(in.GetParentFollowupIntentName())
	// MISSING: FollowupIntentInfo
	return out
}
func Intent_ToProto(mapCtx *direct.MapContext, in *krm.Intent) *pb.Intent {
	if in == nil {
		return nil
	}
	out := &pb.Intent{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.WebhookState = direct.Enum_ToProto[pb.Intent_WebhookState](mapCtx, in.WebhookState)
	out.Priority = direct.ValueOf(in.Priority)
	out.IsFallback = direct.ValueOf(in.IsFallback)
	out.MlEnabled = direct.ValueOf(in.MlEnabled)
	out.MlDisabled = direct.ValueOf(in.MlDisabled)
	out.LiveAgentHandoff = direct.ValueOf(in.LiveAgentHandoff)
	out.EndInteraction = direct.ValueOf(in.EndInteraction)
	out.InputContextNames = in.InputContextNames
	out.Events = in.Events
	out.TrainingPhrases = direct.Slice_ToProto(mapCtx, in.TrainingPhrases, Intent_TrainingPhrase_ToProto)
	out.Action = direct.ValueOf(in.Action)
	out.OutputContexts = direct.Slice_ToProto(mapCtx, in.OutputContexts, Context_ToProto)
	out.ResetContexts = direct.ValueOf(in.ResetContexts)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, Intent_Parameter_ToProto)
	out.Messages = direct.Slice_ToProto(mapCtx, in.Messages, Intent_Message_ToProto)
	out.DefaultResponsePlatforms = direct.EnumSlice_ToProto[pb.Intent_Message_Platform](mapCtx, in.DefaultResponsePlatforms)
	// MISSING: RootFollowupIntentName
	out.ParentFollowupIntentName = direct.ValueOf(in.ParentFollowupIntentName)
	// MISSING: FollowupIntentInfo
	return out
}
func IntentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Intent) *krm.IntentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IntentObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: WebhookState
	// MISSING: Priority
	// MISSING: IsFallback
	// MISSING: MlEnabled
	// MISSING: MlDisabled
	// MISSING: LiveAgentHandoff
	// MISSING: EndInteraction
	// MISSING: InputContextNames
	// MISSING: Events
	out.TrainingPhrases = direct.Slice_FromProto(mapCtx, in.TrainingPhrases, Intent_TrainingPhraseObservedState_FromProto)
	// MISSING: Action
	// MISSING: OutputContexts
	// MISSING: ResetContexts
	// MISSING: Parameters
	// MISSING: Messages
	// MISSING: DefaultResponsePlatforms
	out.RootFollowupIntentName = direct.LazyPtr(in.GetRootFollowupIntentName())
	// MISSING: ParentFollowupIntentName
	out.FollowupIntentInfo = direct.Slice_FromProto(mapCtx, in.FollowupIntentInfo, Intent_FollowupIntentInfo_FromProto)
	return out
}
func IntentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IntentObservedState) *pb.Intent {
	if in == nil {
		return nil
	}
	out := &pb.Intent{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: WebhookState
	// MISSING: Priority
	// MISSING: IsFallback
	// MISSING: MlEnabled
	// MISSING: MlDisabled
	// MISSING: LiveAgentHandoff
	// MISSING: EndInteraction
	// MISSING: InputContextNames
	// MISSING: Events
	out.TrainingPhrases = direct.Slice_ToProto(mapCtx, in.TrainingPhrases, Intent_TrainingPhraseObservedState_ToProto)
	// MISSING: Action
	// MISSING: OutputContexts
	// MISSING: ResetContexts
	// MISSING: Parameters
	// MISSING: Messages
	// MISSING: DefaultResponsePlatforms
	out.RootFollowupIntentName = direct.ValueOf(in.RootFollowupIntentName)
	// MISSING: ParentFollowupIntentName
	out.FollowupIntentInfo = direct.Slice_ToProto(mapCtx, in.FollowupIntentInfo, Intent_FollowupIntentInfo_ToProto)
	return out
}
func IntentSuggestion_FromProto(mapCtx *direct.MapContext, in *pb.IntentSuggestion) *krm.IntentSuggestion {
	if in == nil {
		return nil
	}
	out := &krm.IntentSuggestion{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.IntentV2 = direct.LazyPtr(in.GetIntentV2())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func IntentSuggestion_ToProto(mapCtx *direct.MapContext, in *krm.IntentSuggestion) *pb.IntentSuggestion {
	if in == nil {
		return nil
	}
	out := &pb.IntentSuggestion{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if oneof := IntentSuggestion_IntentV2_ToProto(mapCtx, in.IntentV2); oneof != nil {
		out.Intent = oneof
	}
	out.Description = direct.ValueOf(in.Description)
	return out
}
func Intent_FollowupIntentInfo_FromProto(mapCtx *direct.MapContext, in *pb.Intent_FollowupIntentInfo) *krm.Intent_FollowupIntentInfo {
	if in == nil {
		return nil
	}
	out := &krm.Intent_FollowupIntentInfo{}
	out.FollowupIntentName = direct.LazyPtr(in.GetFollowupIntentName())
	out.ParentFollowupIntentName = direct.LazyPtr(in.GetParentFollowupIntentName())
	return out
}
func Intent_FollowupIntentInfo_ToProto(mapCtx *direct.MapContext, in *krm.Intent_FollowupIntentInfo) *pb.Intent_FollowupIntentInfo {
	if in == nil {
		return nil
	}
	out := &pb.Intent_FollowupIntentInfo{}
	out.FollowupIntentName = direct.ValueOf(in.FollowupIntentName)
	out.ParentFollowupIntentName = direct.ValueOf(in.ParentFollowupIntentName)
	return out
}
func Intent_Message_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message) *krm.Intent_Message {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message{}
	out.Text = Intent_Message_Text_FromProto(mapCtx, in.GetText())
	out.Image = Intent_Message_Image_FromProto(mapCtx, in.GetImage())
	out.QuickReplies = Intent_Message_QuickReplies_FromProto(mapCtx, in.GetQuickReplies())
	out.Card = Intent_Message_Card_FromProto(mapCtx, in.GetCard())
	out.Payload = Payload_FromProto(mapCtx, in.GetPayload())
	out.SimpleResponses = Intent_Message_SimpleResponses_FromProto(mapCtx, in.GetSimpleResponses())
	out.BasicCard = Intent_Message_BasicCard_FromProto(mapCtx, in.GetBasicCard())
	out.Suggestions = Intent_Message_Suggestions_FromProto(mapCtx, in.GetSuggestions())
	out.LinkOutSuggestion = Intent_Message_LinkOutSuggestion_FromProto(mapCtx, in.GetLinkOutSuggestion())
	out.ListSelect = Intent_Message_ListSelect_FromProto(mapCtx, in.GetListSelect())
	out.CarouselSelect = Intent_Message_CarouselSelect_FromProto(mapCtx, in.GetCarouselSelect())
	out.TelephonyPlayAudio = Intent_Message_TelephonyPlayAudio_FromProto(mapCtx, in.GetTelephonyPlayAudio())
	out.TelephonySynthesizeSpeech = Intent_Message_TelephonySynthesizeSpeech_FromProto(mapCtx, in.GetTelephonySynthesizeSpeech())
	out.TelephonyTransferCall = Intent_Message_TelephonyTransferCall_FromProto(mapCtx, in.GetTelephonyTransferCall())
	out.RbmText = Intent_Message_RbmText_FromProto(mapCtx, in.GetRbmText())
	out.RbmStandaloneRichCard = Intent_Message_RbmStandaloneCard_FromProto(mapCtx, in.GetRbmStandaloneRichCard())
	out.RbmCarouselRichCard = Intent_Message_RbmCarouselCard_FromProto(mapCtx, in.GetRbmCarouselRichCard())
	out.BrowseCarouselCard = Intent_Message_BrowseCarouselCard_FromProto(mapCtx, in.GetBrowseCarouselCard())
	out.TableCard = Intent_Message_TableCard_FromProto(mapCtx, in.GetTableCard())
	out.MediaContent = Intent_Message_MediaContent_FromProto(mapCtx, in.GetMediaContent())
	out.Platform = direct.Enum_FromProto(mapCtx, in.GetPlatform())
	return out
}
func Intent_Message_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message) *pb.Intent_Message {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message{}
	if oneof := Intent_Message_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Message = &pb.Intent_Message_Text_{Text: oneof}
	}
	if oneof := Intent_Message_Image_ToProto(mapCtx, in.Image); oneof != nil {
		out.Message = &pb.Intent_Message_Image_{Image: oneof}
	}
	if oneof := Intent_Message_QuickReplies_ToProto(mapCtx, in.QuickReplies); oneof != nil {
		out.Message = &pb.Intent_Message_QuickReplies_{QuickReplies: oneof}
	}
	if oneof := Intent_Message_Card_ToProto(mapCtx, in.Card); oneof != nil {
		out.Message = &pb.Intent_Message_Card_{Card: oneof}
	}
	if oneof := Payload_ToProto(mapCtx, in.Payload); oneof != nil {
		out.Message = &pb.Intent_Message_Payload{Payload: oneof}
	}
	if oneof := Intent_Message_SimpleResponses_ToProto(mapCtx, in.SimpleResponses); oneof != nil {
		out.Message = &pb.Intent_Message_SimpleResponses_{SimpleResponses: oneof}
	}
	if oneof := Intent_Message_BasicCard_ToProto(mapCtx, in.BasicCard); oneof != nil {
		out.Message = &pb.Intent_Message_BasicCard_{BasicCard: oneof}
	}
	if oneof := Intent_Message_Suggestions_ToProto(mapCtx, in.Suggestions); oneof != nil {
		out.Message = &pb.Intent_Message_Suggestions_{Suggestions: oneof}
	}
	if oneof := Intent_Message_LinkOutSuggestion_ToProto(mapCtx, in.LinkOutSuggestion); oneof != nil {
		out.Message = &pb.Intent_Message_LinkOutSuggestion_{LinkOutSuggestion: oneof}
	}
	if oneof := Intent_Message_ListSelect_ToProto(mapCtx, in.ListSelect); oneof != nil {
		out.Message = &pb.Intent_Message_ListSelect_{ListSelect: oneof}
	}
	if oneof := Intent_Message_CarouselSelect_ToProto(mapCtx, in.CarouselSelect); oneof != nil {
		out.Message = &pb.Intent_Message_CarouselSelect_{CarouselSelect: oneof}
	}
	if oneof := Intent_Message_TelephonyPlayAudio_ToProto(mapCtx, in.TelephonyPlayAudio); oneof != nil {
		out.Message = &pb.Intent_Message_TelephonyPlayAudio_{TelephonyPlayAudio: oneof}
	}
	if oneof := Intent_Message_TelephonySynthesizeSpeech_ToProto(mapCtx, in.TelephonySynthesizeSpeech); oneof != nil {
		out.Message = &pb.Intent_Message_TelephonySynthesizeSpeech_{TelephonySynthesizeSpeech: oneof}
	}
	if oneof := Intent_Message_TelephonyTransferCall_ToProto(mapCtx, in.TelephonyTransferCall); oneof != nil {
		out.Message = &pb.Intent_Message_TelephonyTransferCall_{TelephonyTransferCall: oneof}
	}
	if oneof := Intent_Message_RbmText_ToProto(mapCtx, in.RbmText); oneof != nil {
		out.Message = &pb.Intent_Message_RbmText_{RbmText: oneof}
	}
	if oneof := Intent_Message_RbmStandaloneCard_ToProto(mapCtx, in.RbmStandaloneRichCard); oneof != nil {
		out.Message = &pb.Intent_Message_RbmStandaloneRichCard{RbmStandaloneRichCard: oneof}
	}
	if oneof := Intent_Message_RbmCarouselCard_ToProto(mapCtx, in.RbmCarouselRichCard); oneof != nil {
		out.Message = &pb.Intent_Message_RbmCarouselRichCard{RbmCarouselRichCard: oneof}
	}
	if oneof := Intent_Message_BrowseCarouselCard_ToProto(mapCtx, in.BrowseCarouselCard); oneof != nil {
		out.Message = &pb.Intent_Message_BrowseCarouselCard_{BrowseCarouselCard: oneof}
	}
	if oneof := Intent_Message_TableCard_ToProto(mapCtx, in.TableCard); oneof != nil {
		out.Message = &pb.Intent_Message_TableCard_{TableCard: oneof}
	}
	if oneof := Intent_Message_MediaContent_ToProto(mapCtx, in.MediaContent); oneof != nil {
		out.Message = &pb.Intent_Message_MediaContent_{MediaContent: oneof}
	}
	out.Platform = direct.Enum_ToProto[pb.Intent_Message_Platform](mapCtx, in.Platform)
	return out
}
func Intent_Message_BasicCard_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_BasicCard) *krm.Intent_Message_BasicCard {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_BasicCard{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Subtitle = direct.LazyPtr(in.GetSubtitle())
	out.FormattedText = direct.LazyPtr(in.GetFormattedText())
	out.Image = Intent_Message_Image_FromProto(mapCtx, in.GetImage())
	out.Buttons = direct.Slice_FromProto(mapCtx, in.Buttons, Intent_Message_BasicCard_Button_FromProto)
	return out
}
func Intent_Message_BasicCard_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_BasicCard) *pb.Intent_Message_BasicCard {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_BasicCard{}
	out.Title = direct.ValueOf(in.Title)
	out.Subtitle = direct.ValueOf(in.Subtitle)
	out.FormattedText = direct.ValueOf(in.FormattedText)
	out.Image = Intent_Message_Image_ToProto(mapCtx, in.Image)
	out.Buttons = direct.Slice_ToProto(mapCtx, in.Buttons, Intent_Message_BasicCard_Button_ToProto)
	return out
}
func Intent_Message_BasicCard_Button_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_BasicCard_Button) *krm.Intent_Message_BasicCard_Button {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_BasicCard_Button{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.OpenURIAction = Intent_Message_BasicCard_Button_OpenUriAction_FromProto(mapCtx, in.GetOpenUriAction())
	return out
}
func Intent_Message_BasicCard_Button_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_BasicCard_Button) *pb.Intent_Message_BasicCard_Button {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_BasicCard_Button{}
	out.Title = direct.ValueOf(in.Title)
	out.OpenUriAction = Intent_Message_BasicCard_Button_OpenUriAction_ToProto(mapCtx, in.OpenURIAction)
	return out
}
func Intent_Message_BasicCard_Button_OpenUriAction_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_BasicCard_Button_OpenUriAction) *krm.Intent_Message_BasicCard_Button_OpenUriAction {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_BasicCard_Button_OpenUriAction{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Intent_Message_BasicCard_Button_OpenUriAction_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_BasicCard_Button_OpenUriAction) *pb.Intent_Message_BasicCard_Button_OpenUriAction {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_BasicCard_Button_OpenUriAction{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Intent_Message_BrowseCarouselCard_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_BrowseCarouselCard) *krm.Intent_Message_BrowseCarouselCard {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_BrowseCarouselCard{}
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_FromProto)
	out.ImageDisplayOptions = direct.Enum_FromProto(mapCtx, in.GetImageDisplayOptions())
	return out
}
func Intent_Message_BrowseCarouselCard_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_BrowseCarouselCard) *pb.Intent_Message_BrowseCarouselCard {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_BrowseCarouselCard{}
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_ToProto)
	out.ImageDisplayOptions = direct.Enum_ToProto[pb.Intent_Message_BrowseCarouselCard_ImageDisplayOptions](mapCtx, in.ImageDisplayOptions)
	return out
}
func Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem) *krm.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem{}
	out.OpenURIAction = Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction_FromProto(mapCtx, in.GetOpenUriAction())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Image = Intent_Message_Image_FromProto(mapCtx, in.GetImage())
	out.Footer = direct.LazyPtr(in.GetFooter())
	return out
}
func Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem) *pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem{}
	out.OpenUriAction = Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction_ToProto(mapCtx, in.OpenURIAction)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Image = Intent_Message_Image_ToProto(mapCtx, in.Image)
	out.Footer = direct.ValueOf(in.Footer)
	return out
}
func Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction) *krm.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction{}
	out.URL = direct.LazyPtr(in.GetUrl())
	out.URLTypeHint = direct.Enum_FromProto(mapCtx, in.GetUrlTypeHint())
	return out
}
func Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction) *pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction{}
	out.Url = direct.ValueOf(in.URL)
	out.UrlTypeHint = direct.Enum_ToProto[pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction_UrlTypeHint](mapCtx, in.URLTypeHint)
	return out
}
func Intent_Message_Card_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_Card) *krm.Intent_Message_Card {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_Card{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Subtitle = direct.LazyPtr(in.GetSubtitle())
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.Buttons = direct.Slice_FromProto(mapCtx, in.Buttons, Intent_Message_Card_Button_FromProto)
	return out
}
func Intent_Message_Card_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_Card) *pb.Intent_Message_Card {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_Card{}
	out.Title = direct.ValueOf(in.Title)
	out.Subtitle = direct.ValueOf(in.Subtitle)
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.Buttons = direct.Slice_ToProto(mapCtx, in.Buttons, Intent_Message_Card_Button_ToProto)
	return out
}
func Intent_Message_Card_Button_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_Card_Button) *krm.Intent_Message_Card_Button {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_Card_Button{}
	out.Text = direct.LazyPtr(in.GetText())
	out.Postback = direct.LazyPtr(in.GetPostback())
	return out
}
func Intent_Message_Card_Button_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_Card_Button) *pb.Intent_Message_Card_Button {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_Card_Button{}
	out.Text = direct.ValueOf(in.Text)
	out.Postback = direct.ValueOf(in.Postback)
	return out
}
func Intent_Message_CarouselSelect_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_CarouselSelect) *krm.Intent_Message_CarouselSelect {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_CarouselSelect{}
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, Intent_Message_CarouselSelect_Item_FromProto)
	return out
}
func Intent_Message_CarouselSelect_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_CarouselSelect) *pb.Intent_Message_CarouselSelect {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_CarouselSelect{}
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, Intent_Message_CarouselSelect_Item_ToProto)
	return out
}
func Intent_Message_CarouselSelect_Item_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_CarouselSelect_Item) *krm.Intent_Message_CarouselSelect_Item {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_CarouselSelect_Item{}
	out.Info = Intent_Message_SelectItemInfo_FromProto(mapCtx, in.GetInfo())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Image = Intent_Message_Image_FromProto(mapCtx, in.GetImage())
	return out
}
func Intent_Message_CarouselSelect_Item_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_CarouselSelect_Item) *pb.Intent_Message_CarouselSelect_Item {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_CarouselSelect_Item{}
	out.Info = Intent_Message_SelectItemInfo_ToProto(mapCtx, in.Info)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Image = Intent_Message_Image_ToProto(mapCtx, in.Image)
	return out
}
func Intent_Message_ColumnProperties_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_ColumnProperties) *krm.Intent_Message_ColumnProperties {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_ColumnProperties{}
	out.Header = direct.LazyPtr(in.GetHeader())
	out.HorizontalAlignment = direct.Enum_FromProto(mapCtx, in.GetHorizontalAlignment())
	return out
}
func Intent_Message_ColumnProperties_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_ColumnProperties) *pb.Intent_Message_ColumnProperties {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_ColumnProperties{}
	out.Header = direct.ValueOf(in.Header)
	out.HorizontalAlignment = direct.Enum_ToProto[pb.Intent_Message_ColumnProperties_HorizontalAlignment](mapCtx, in.HorizontalAlignment)
	return out
}
func Intent_Message_Image_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_Image) *krm.Intent_Message_Image {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_Image{}
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.AccessibilityText = direct.LazyPtr(in.GetAccessibilityText())
	return out
}
func Intent_Message_Image_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_Image) *pb.Intent_Message_Image {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_Image{}
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.AccessibilityText = direct.ValueOf(in.AccessibilityText)
	return out
}
func Intent_Message_LinkOutSuggestion_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_LinkOutSuggestion) *krm.Intent_Message_LinkOutSuggestion {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_LinkOutSuggestion{}
	out.DestinationName = direct.LazyPtr(in.GetDestinationName())
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Intent_Message_LinkOutSuggestion_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_LinkOutSuggestion) *pb.Intent_Message_LinkOutSuggestion {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_LinkOutSuggestion{}
	out.DestinationName = direct.ValueOf(in.DestinationName)
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Intent_Message_ListSelect_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_ListSelect) *krm.Intent_Message_ListSelect {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_ListSelect{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, Intent_Message_ListSelect_Item_FromProto)
	out.Subtitle = direct.LazyPtr(in.GetSubtitle())
	return out
}
func Intent_Message_ListSelect_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_ListSelect) *pb.Intent_Message_ListSelect {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_ListSelect{}
	out.Title = direct.ValueOf(in.Title)
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, Intent_Message_ListSelect_Item_ToProto)
	out.Subtitle = direct.ValueOf(in.Subtitle)
	return out
}
func Intent_Message_ListSelect_Item_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_ListSelect_Item) *krm.Intent_Message_ListSelect_Item {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_ListSelect_Item{}
	out.Info = Intent_Message_SelectItemInfo_FromProto(mapCtx, in.GetInfo())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Image = Intent_Message_Image_FromProto(mapCtx, in.GetImage())
	return out
}
func Intent_Message_ListSelect_Item_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_ListSelect_Item) *pb.Intent_Message_ListSelect_Item {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_ListSelect_Item{}
	out.Info = Intent_Message_SelectItemInfo_ToProto(mapCtx, in.Info)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Image = Intent_Message_Image_ToProto(mapCtx, in.Image)
	return out
}
func Intent_Message_MediaContent_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_MediaContent) *krm.Intent_Message_MediaContent {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_MediaContent{}
	out.MediaType = direct.Enum_FromProto(mapCtx, in.GetMediaType())
	out.MediaObjects = direct.Slice_FromProto(mapCtx, in.MediaObjects, Intent_Message_MediaContent_ResponseMediaObject_FromProto)
	return out
}
func Intent_Message_MediaContent_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_MediaContent) *pb.Intent_Message_MediaContent {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_MediaContent{}
	out.MediaType = direct.Enum_ToProto[pb.Intent_Message_MediaContent_ResponseMediaType](mapCtx, in.MediaType)
	out.MediaObjects = direct.Slice_ToProto(mapCtx, in.MediaObjects, Intent_Message_MediaContent_ResponseMediaObject_ToProto)
	return out
}
func Intent_Message_MediaContent_ResponseMediaObject_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_MediaContent_ResponseMediaObject) *krm.Intent_Message_MediaContent_ResponseMediaObject {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_MediaContent_ResponseMediaObject{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.LargeImage = Intent_Message_Image_FromProto(mapCtx, in.GetLargeImage())
	out.Icon = Intent_Message_Image_FromProto(mapCtx, in.GetIcon())
	out.ContentURL = direct.LazyPtr(in.GetContentUrl())
	return out
}
func Intent_Message_MediaContent_ResponseMediaObject_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_MediaContent_ResponseMediaObject) *pb.Intent_Message_MediaContent_ResponseMediaObject {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_MediaContent_ResponseMediaObject{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	if oneof := Intent_Message_Image_ToProto(mapCtx, in.LargeImage); oneof != nil {
		out.Image = &pb.Intent_Message_MediaContent_ResponseMediaObject_LargeImage{LargeImage: oneof}
	}
	if oneof := Intent_Message_Image_ToProto(mapCtx, in.Icon); oneof != nil {
		out.Image = &pb.Intent_Message_MediaContent_ResponseMediaObject_Icon{Icon: oneof}
	}
	out.ContentUrl = direct.ValueOf(in.ContentURL)
	return out
}
func Intent_Message_QuickReplies_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_QuickReplies) *krm.Intent_Message_QuickReplies {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_QuickReplies{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.QuickReplies = in.QuickReplies
	return out
}
func Intent_Message_QuickReplies_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_QuickReplies) *pb.Intent_Message_QuickReplies {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_QuickReplies{}
	out.Title = direct.ValueOf(in.Title)
	out.QuickReplies = in.QuickReplies
	return out
}
func Intent_Message_RbmCardContent_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_RbmCardContent) *krm.Intent_Message_RbmCardContent {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_RbmCardContent{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Media = Intent_Message_RbmCardContent_RbmMedia_FromProto(mapCtx, in.GetMedia())
	out.Suggestions = direct.Slice_FromProto(mapCtx, in.Suggestions, Intent_Message_RbmSuggestion_FromProto)
	return out
}
func Intent_Message_RbmCardContent_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_RbmCardContent) *pb.Intent_Message_RbmCardContent {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_RbmCardContent{}
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Media = Intent_Message_RbmCardContent_RbmMedia_ToProto(mapCtx, in.Media)
	out.Suggestions = direct.Slice_ToProto(mapCtx, in.Suggestions, Intent_Message_RbmSuggestion_ToProto)
	return out
}
func Intent_Message_RbmCardContent_RbmMedia_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_RbmCardContent_RbmMedia) *krm.Intent_Message_RbmCardContent_RbmMedia {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_RbmCardContent_RbmMedia{}
	out.FileURI = direct.LazyPtr(in.GetFileUri())
	out.ThumbnailURI = direct.LazyPtr(in.GetThumbnailUri())
	out.Height = direct.Enum_FromProto(mapCtx, in.GetHeight())
	return out
}
func Intent_Message_RbmCardContent_RbmMedia_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_RbmCardContent_RbmMedia) *pb.Intent_Message_RbmCardContent_RbmMedia {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_RbmCardContent_RbmMedia{}
	out.FileUri = direct.ValueOf(in.FileURI)
	out.ThumbnailUri = direct.ValueOf(in.ThumbnailURI)
	out.Height = direct.Enum_ToProto[pb.Intent_Message_RbmCardContent_RbmMedia_Height](mapCtx, in.Height)
	return out
}
func Intent_Message_RbmCarouselCard_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_RbmCarouselCard) *krm.Intent_Message_RbmCarouselCard {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_RbmCarouselCard{}
	out.CardWidth = direct.Enum_FromProto(mapCtx, in.GetCardWidth())
	out.CardContents = direct.Slice_FromProto(mapCtx, in.CardContents, Intent_Message_RbmCardContent_FromProto)
	return out
}
func Intent_Message_RbmCarouselCard_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_RbmCarouselCard) *pb.Intent_Message_RbmCarouselCard {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_RbmCarouselCard{}
	out.CardWidth = direct.Enum_ToProto[pb.Intent_Message_RbmCarouselCard_CardWidth](mapCtx, in.CardWidth)
	out.CardContents = direct.Slice_ToProto(mapCtx, in.CardContents, Intent_Message_RbmCardContent_ToProto)
	return out
}
func Intent_Message_RbmStandaloneCard_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_RbmStandaloneCard) *krm.Intent_Message_RbmStandaloneCard {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_RbmStandaloneCard{}
	out.CardOrientation = direct.Enum_FromProto(mapCtx, in.GetCardOrientation())
	out.ThumbnailImageAlignment = direct.Enum_FromProto(mapCtx, in.GetThumbnailImageAlignment())
	out.CardContent = Intent_Message_RbmCardContent_FromProto(mapCtx, in.GetCardContent())
	return out
}
func Intent_Message_RbmStandaloneCard_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_RbmStandaloneCard) *pb.Intent_Message_RbmStandaloneCard {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_RbmStandaloneCard{}
	out.CardOrientation = direct.Enum_ToProto[pb.Intent_Message_RbmStandaloneCard_CardOrientation](mapCtx, in.CardOrientation)
	out.ThumbnailImageAlignment = direct.Enum_ToProto[pb.Intent_Message_RbmStandaloneCard_ThumbnailImageAlignment](mapCtx, in.ThumbnailImageAlignment)
	out.CardContent = Intent_Message_RbmCardContent_ToProto(mapCtx, in.CardContent)
	return out
}
func Intent_Message_RbmSuggestedAction_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_RbmSuggestedAction) *krm.Intent_Message_RbmSuggestedAction {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_RbmSuggestedAction{}
	out.Text = direct.LazyPtr(in.GetText())
	out.PostbackData = direct.LazyPtr(in.GetPostbackData())
	out.Dial = Intent_Message_RbmSuggestedAction_RbmSuggestedActionDial_FromProto(mapCtx, in.GetDial())
	out.OpenURL = Intent_Message_RbmSuggestedAction_RbmSuggestedActionOpenUri_FromProto(mapCtx, in.GetOpenUrl())
	out.ShareLocation = Intent_Message_RbmSuggestedAction_RbmSuggestedActionShareLocation_FromProto(mapCtx, in.GetShareLocation())
	return out
}
func Intent_Message_RbmSuggestedAction_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_RbmSuggestedAction) *pb.Intent_Message_RbmSuggestedAction {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_RbmSuggestedAction{}
	out.Text = direct.ValueOf(in.Text)
	out.PostbackData = direct.ValueOf(in.PostbackData)
	if oneof := Intent_Message_RbmSuggestedAction_RbmSuggestedActionDial_ToProto(mapCtx, in.Dial); oneof != nil {
		out.Action = &pb.Intent_Message_RbmSuggestedAction_Dial{Dial: oneof}
	}
	if oneof := Intent_Message_RbmSuggestedAction_RbmSuggestedActionOpenUri_ToProto(mapCtx, in.OpenURL); oneof != nil {
		out.Action = &pb.Intent_Message_RbmSuggestedAction_OpenUrl{OpenUrl: oneof}
	}
	if oneof := Intent_Message_RbmSuggestedAction_RbmSuggestedActionShareLocation_ToProto(mapCtx, in.ShareLocation); oneof != nil {
		out.Action = &pb.Intent_Message_RbmSuggestedAction_ShareLocation{ShareLocation: oneof}
	}
	return out
}
func Intent_Message_RbmSuggestedAction_RbmSuggestedActionDial_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_RbmSuggestedAction_RbmSuggestedActionDial) *krm.Intent_Message_RbmSuggestedAction_RbmSuggestedActionDial {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_RbmSuggestedAction_RbmSuggestedActionDial{}
	out.PhoneNumber = direct.LazyPtr(in.GetPhoneNumber())
	return out
}
func Intent_Message_RbmSuggestedAction_RbmSuggestedActionDial_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_RbmSuggestedAction_RbmSuggestedActionDial) *pb.Intent_Message_RbmSuggestedAction_RbmSuggestedActionDial {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_RbmSuggestedAction_RbmSuggestedActionDial{}
	out.PhoneNumber = direct.ValueOf(in.PhoneNumber)
	return out
}
func Intent_Message_RbmSuggestedAction_RbmSuggestedActionOpenUri_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_RbmSuggestedAction_RbmSuggestedActionOpenUri) *krm.Intent_Message_RbmSuggestedAction_RbmSuggestedActionOpenUri {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_RbmSuggestedAction_RbmSuggestedActionOpenUri{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Intent_Message_RbmSuggestedAction_RbmSuggestedActionOpenUri_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_RbmSuggestedAction_RbmSuggestedActionOpenUri) *pb.Intent_Message_RbmSuggestedAction_RbmSuggestedActionOpenUri {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_RbmSuggestedAction_RbmSuggestedActionOpenUri{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Intent_Message_RbmSuggestedAction_RbmSuggestedActionShareLocation_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_RbmSuggestedAction_RbmSuggestedActionShareLocation) *krm.Intent_Message_RbmSuggestedAction_RbmSuggestedActionShareLocation {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_RbmSuggestedAction_RbmSuggestedActionShareLocation{}
	return out
}
func Intent_Message_RbmSuggestedAction_RbmSuggestedActionShareLocation_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_RbmSuggestedAction_RbmSuggestedActionShareLocation) *pb.Intent_Message_RbmSuggestedAction_RbmSuggestedActionShareLocation {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_RbmSuggestedAction_RbmSuggestedActionShareLocation{}
	return out
}
func Intent_Message_RbmSuggestedReply_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_RbmSuggestedReply) *krm.Intent_Message_RbmSuggestedReply {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_RbmSuggestedReply{}
	out.Text = direct.LazyPtr(in.GetText())
	out.PostbackData = direct.LazyPtr(in.GetPostbackData())
	return out
}
func Intent_Message_RbmSuggestedReply_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_RbmSuggestedReply) *pb.Intent_Message_RbmSuggestedReply {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_RbmSuggestedReply{}
	out.Text = direct.ValueOf(in.Text)
	out.PostbackData = direct.ValueOf(in.PostbackData)
	return out
}
func Intent_Message_RbmSuggestion_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_RbmSuggestion) *krm.Intent_Message_RbmSuggestion {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_RbmSuggestion{}
	out.Reply = Intent_Message_RbmSuggestedReply_FromProto(mapCtx, in.GetReply())
	out.Action = Intent_Message_RbmSuggestedAction_FromProto(mapCtx, in.GetAction())
	return out
}
func Intent_Message_RbmSuggestion_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_RbmSuggestion) *pb.Intent_Message_RbmSuggestion {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_RbmSuggestion{}
	if oneof := Intent_Message_RbmSuggestedReply_ToProto(mapCtx, in.Reply); oneof != nil {
		out.Suggestion = &pb.Intent_Message_RbmSuggestion_Reply{Reply: oneof}
	}
	if oneof := Intent_Message_RbmSuggestedAction_ToProto(mapCtx, in.Action); oneof != nil {
		out.Suggestion = &pb.Intent_Message_RbmSuggestion_Action{Action: oneof}
	}
	return out
}
func Intent_Message_RbmText_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_RbmText) *krm.Intent_Message_RbmText {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_RbmText{}
	out.Text = direct.LazyPtr(in.GetText())
	out.RbmSuggestion = direct.Slice_FromProto(mapCtx, in.RbmSuggestion, Intent_Message_RbmSuggestion_FromProto)
	return out
}
func Intent_Message_RbmText_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_RbmText) *pb.Intent_Message_RbmText {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_RbmText{}
	out.Text = direct.ValueOf(in.Text)
	out.RbmSuggestion = direct.Slice_ToProto(mapCtx, in.RbmSuggestion, Intent_Message_RbmSuggestion_ToProto)
	return out
}
func Intent_Message_SelectItemInfo_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_SelectItemInfo) *krm.Intent_Message_SelectItemInfo {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_SelectItemInfo{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Synonyms = in.Synonyms
	return out
}
func Intent_Message_SelectItemInfo_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_SelectItemInfo) *pb.Intent_Message_SelectItemInfo {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_SelectItemInfo{}
	out.Key = direct.ValueOf(in.Key)
	out.Synonyms = in.Synonyms
	return out
}
func Intent_Message_SimpleResponses_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_SimpleResponses) *krm.Intent_Message_SimpleResponses {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_SimpleResponses{}
	out.SimpleResponses = direct.Slice_FromProto(mapCtx, in.SimpleResponses, Intent_Message_SimpleResponse_FromProto)
	return out
}
func Intent_Message_SimpleResponses_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_SimpleResponses) *pb.Intent_Message_SimpleResponses {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_SimpleResponses{}
	out.SimpleResponses = direct.Slice_ToProto(mapCtx, in.SimpleResponses, Intent_Message_SimpleResponse_ToProto)
	return out
}
func Intent_Message_Suggestion_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_Suggestion) *krm.Intent_Message_Suggestion {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_Suggestion{}
	out.Title = direct.LazyPtr(in.GetTitle())
	return out
}
func Intent_Message_Suggestion_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_Suggestion) *pb.Intent_Message_Suggestion {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_Suggestion{}
	out.Title = direct.ValueOf(in.Title)
	return out
}
func Intent_Message_Suggestions_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_Suggestions) *krm.Intent_Message_Suggestions {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_Suggestions{}
	out.Suggestions = direct.Slice_FromProto(mapCtx, in.Suggestions, Intent_Message_Suggestion_FromProto)
	return out
}
func Intent_Message_Suggestions_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_Suggestions) *pb.Intent_Message_Suggestions {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_Suggestions{}
	out.Suggestions = direct.Slice_ToProto(mapCtx, in.Suggestions, Intent_Message_Suggestion_ToProto)
	return out
}
func Intent_Message_TableCard_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_TableCard) *krm.Intent_Message_TableCard {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_TableCard{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Subtitle = direct.LazyPtr(in.GetSubtitle())
	out.Image = Intent_Message_Image_FromProto(mapCtx, in.GetImage())
	out.ColumnProperties = direct.Slice_FromProto(mapCtx, in.ColumnProperties, Intent_Message_ColumnProperties_FromProto)
	out.Rows = direct.Slice_FromProto(mapCtx, in.Rows, Intent_Message_TableCardRow_FromProto)
	out.Buttons = direct.Slice_FromProto(mapCtx, in.Buttons, Intent_Message_BasicCard_Button_FromProto)
	return out
}
func Intent_Message_TableCard_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_TableCard) *pb.Intent_Message_TableCard {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_TableCard{}
	out.Title = direct.ValueOf(in.Title)
	out.Subtitle = direct.ValueOf(in.Subtitle)
	out.Image = Intent_Message_Image_ToProto(mapCtx, in.Image)
	out.ColumnProperties = direct.Slice_ToProto(mapCtx, in.ColumnProperties, Intent_Message_ColumnProperties_ToProto)
	out.Rows = direct.Slice_ToProto(mapCtx, in.Rows, Intent_Message_TableCardRow_ToProto)
	out.Buttons = direct.Slice_ToProto(mapCtx, in.Buttons, Intent_Message_BasicCard_Button_ToProto)
	return out
}
func Intent_Message_TableCardCell_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_TableCardCell) *krm.Intent_Message_TableCardCell {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_TableCardCell{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func Intent_Message_TableCardCell_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_TableCardCell) *pb.Intent_Message_TableCardCell {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_TableCardCell{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
func Intent_Message_TableCardRow_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_TableCardRow) *krm.Intent_Message_TableCardRow {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_TableCardRow{}
	out.Cells = direct.Slice_FromProto(mapCtx, in.Cells, Intent_Message_TableCardCell_FromProto)
	out.DividerAfter = direct.LazyPtr(in.GetDividerAfter())
	return out
}
func Intent_Message_TableCardRow_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_TableCardRow) *pb.Intent_Message_TableCardRow {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_TableCardRow{}
	out.Cells = direct.Slice_ToProto(mapCtx, in.Cells, Intent_Message_TableCardCell_ToProto)
	out.DividerAfter = direct.ValueOf(in.DividerAfter)
	return out
}
func Intent_Message_TelephonyPlayAudio_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_TelephonyPlayAudio) *krm.Intent_Message_TelephonyPlayAudio {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_TelephonyPlayAudio{}
	out.AudioURI = direct.LazyPtr(in.GetAudioUri())
	return out
}
func Intent_Message_TelephonyPlayAudio_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_TelephonyPlayAudio) *pb.Intent_Message_TelephonyPlayAudio {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_TelephonyPlayAudio{}
	out.AudioUri = direct.ValueOf(in.AudioURI)
	return out
}
func Intent_Message_TelephonySynthesizeSpeech_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_TelephonySynthesizeSpeech) *krm.Intent_Message_TelephonySynthesizeSpeech {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_TelephonySynthesizeSpeech{}
	out.Text = direct.LazyPtr(in.GetText())
	out.Ssml = direct.LazyPtr(in.GetSsml())
	return out
}
func Intent_Message_TelephonySynthesizeSpeech_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_TelephonySynthesizeSpeech) *pb.Intent_Message_TelephonySynthesizeSpeech {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_TelephonySynthesizeSpeech{}
	if oneof := Intent_Message_TelephonySynthesizeSpeech_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Source = oneof
	}
	if oneof := Intent_Message_TelephonySynthesizeSpeech_Ssml_ToProto(mapCtx, in.Ssml); oneof != nil {
		out.Source = oneof
	}
	return out
}
func Intent_Message_TelephonyTransferCall_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_TelephonyTransferCall) *krm.Intent_Message_TelephonyTransferCall {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_TelephonyTransferCall{}
	out.PhoneNumber = direct.LazyPtr(in.GetPhoneNumber())
	return out
}
func Intent_Message_TelephonyTransferCall_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_TelephonyTransferCall) *pb.Intent_Message_TelephonyTransferCall {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_TelephonyTransferCall{}
	out.PhoneNumber = direct.ValueOf(in.PhoneNumber)
	return out
}
func Intent_Message_Text_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_Text) *krm.Intent_Message_Text {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_Text{}
	out.Text = in.Text
	return out
}
func Intent_Message_Text_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_Text) *pb.Intent_Message_Text {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_Text{}
	out.Text = in.Text
	return out
}
func Intent_Parameter_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Parameter) *krm.Intent_Parameter {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Parameter{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Value = direct.LazyPtr(in.GetValue())
	out.DefaultValue = direct.LazyPtr(in.GetDefaultValue())
	out.EntityTypeDisplayName = direct.LazyPtr(in.GetEntityTypeDisplayName())
	out.Mandatory = direct.LazyPtr(in.GetMandatory())
	out.Prompts = in.Prompts
	out.IsList = direct.LazyPtr(in.GetIsList())
	return out
}
func Intent_Parameter_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Parameter) *pb.Intent_Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Parameter{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Value = direct.ValueOf(in.Value)
	out.DefaultValue = direct.ValueOf(in.DefaultValue)
	out.EntityTypeDisplayName = direct.ValueOf(in.EntityTypeDisplayName)
	out.Mandatory = direct.ValueOf(in.Mandatory)
	out.Prompts = in.Prompts
	out.IsList = direct.ValueOf(in.IsList)
	return out
}
func Intent_TrainingPhrase_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase) *krm.Intent_TrainingPhrase {
	if in == nil {
		return nil
	}
	out := &krm.Intent_TrainingPhrase{}
	// MISSING: Name
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Parts = direct.Slice_FromProto(mapCtx, in.Parts, Intent_TrainingPhrase_Part_FromProto)
	out.TimesAddedCount = direct.LazyPtr(in.GetTimesAddedCount())
	return out
}
func Intent_TrainingPhrase_ToProto(mapCtx *direct.MapContext, in *krm.Intent_TrainingPhrase) *pb.Intent_TrainingPhrase {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase{}
	// MISSING: Name
	out.Type = direct.Enum_ToProto[pb.Intent_TrainingPhrase_Type](mapCtx, in.Type)
	out.Parts = direct.Slice_ToProto(mapCtx, in.Parts, Intent_TrainingPhrase_Part_ToProto)
	out.TimesAddedCount = direct.ValueOf(in.TimesAddedCount)
	return out
}
func Intent_TrainingPhraseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase) *krm.Intent_TrainingPhraseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Intent_TrainingPhraseObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Type
	// MISSING: Parts
	// MISSING: TimesAddedCount
	return out
}
func Intent_TrainingPhraseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Intent_TrainingPhraseObservedState) *pb.Intent_TrainingPhrase {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Type
	// MISSING: Parts
	// MISSING: TimesAddedCount
	return out
}
func Intent_TrainingPhrase_Part_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase_Part) *krm.Intent_TrainingPhrase_Part {
	if in == nil {
		return nil
	}
	out := &krm.Intent_TrainingPhrase_Part{}
	out.Text = direct.LazyPtr(in.GetText())
	out.EntityType = direct.LazyPtr(in.GetEntityType())
	out.Alias = direct.LazyPtr(in.GetAlias())
	out.UserDefined = direct.LazyPtr(in.GetUserDefined())
	return out
}
func Intent_TrainingPhrase_Part_ToProto(mapCtx *direct.MapContext, in *krm.Intent_TrainingPhrase_Part) *pb.Intent_TrainingPhrase_Part {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase_Part{}
	out.Text = direct.ValueOf(in.Text)
	out.EntityType = direct.ValueOf(in.EntityType)
	out.Alias = direct.ValueOf(in.Alias)
	out.UserDefined = direct.ValueOf(in.UserDefined)
	return out
}
func KnowledgeAnswers_FromProto(mapCtx *direct.MapContext, in *pb.KnowledgeAnswers) *krm.KnowledgeAnswers {
	if in == nil {
		return nil
	}
	out := &krm.KnowledgeAnswers{}
	out.Answers = direct.Slice_FromProto(mapCtx, in.Answers, KnowledgeAnswers_Answer_FromProto)
	return out
}
func KnowledgeAnswers_ToProto(mapCtx *direct.MapContext, in *krm.KnowledgeAnswers) *pb.KnowledgeAnswers {
	if in == nil {
		return nil
	}
	out := &pb.KnowledgeAnswers{}
	out.Answers = direct.Slice_ToProto(mapCtx, in.Answers, KnowledgeAnswers_Answer_ToProto)
	return out
}
func KnowledgeAnswers_Answer_FromProto(mapCtx *direct.MapContext, in *pb.KnowledgeAnswers_Answer) *krm.KnowledgeAnswers_Answer {
	if in == nil {
		return nil
	}
	out := &krm.KnowledgeAnswers_Answer{}
	out.Source = direct.LazyPtr(in.GetSource())
	out.FaqQuestion = direct.LazyPtr(in.GetFaqQuestion())
	out.Answer = direct.LazyPtr(in.GetAnswer())
	out.MatchConfidenceLevel = direct.Enum_FromProto(mapCtx, in.GetMatchConfidenceLevel())
	out.MatchConfidence = direct.LazyPtr(in.GetMatchConfidence())
	return out
}
func KnowledgeAnswers_Answer_ToProto(mapCtx *direct.MapContext, in *krm.KnowledgeAnswers_Answer) *pb.KnowledgeAnswers_Answer {
	if in == nil {
		return nil
	}
	out := &pb.KnowledgeAnswers_Answer{}
	out.Source = direct.ValueOf(in.Source)
	out.FaqQuestion = direct.ValueOf(in.FaqQuestion)
	out.Answer = direct.ValueOf(in.Answer)
	out.MatchConfidenceLevel = direct.Enum_ToProto[pb.KnowledgeAnswers_Answer_MatchConfidenceLevel](mapCtx, in.MatchConfidenceLevel)
	out.MatchConfidence = direct.ValueOf(in.MatchConfidence)
	return out
}
func QueryResult_FromProto(mapCtx *direct.MapContext, in *pb.QueryResult) *krm.QueryResult {
	if in == nil {
		return nil
	}
	out := &krm.QueryResult{}
	out.QueryText = direct.LazyPtr(in.GetQueryText())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.SpeechRecognitionConfidence = direct.LazyPtr(in.GetSpeechRecognitionConfidence())
	out.Action = direct.LazyPtr(in.GetAction())
	out.Parameters = Parameters_FromProto(mapCtx, in.GetParameters())
	out.AllRequiredParamsPresent = direct.LazyPtr(in.GetAllRequiredParamsPresent())
	out.CancelsSlotFilling = direct.LazyPtr(in.GetCancelsSlotFilling())
	out.FulfillmentText = direct.LazyPtr(in.GetFulfillmentText())
	out.FulfillmentMessages = direct.Slice_FromProto(mapCtx, in.FulfillmentMessages, Intent_Message_FromProto)
	out.WebhookSource = direct.LazyPtr(in.GetWebhookSource())
	out.WebhookPayload = WebhookPayload_FromProto(mapCtx, in.GetWebhookPayload())
	out.OutputContexts = direct.Slice_FromProto(mapCtx, in.OutputContexts, Context_FromProto)
	out.Intent = Intent_FromProto(mapCtx, in.GetIntent())
	out.IntentDetectionConfidence = direct.LazyPtr(in.GetIntentDetectionConfidence())
	out.DiagnosticInfo = DiagnosticInfo_FromProto(mapCtx, in.GetDiagnosticInfo())
	out.SentimentAnalysisResult = SentimentAnalysisResult_FromProto(mapCtx, in.GetSentimentAnalysisResult())
	out.KnowledgeAnswers = KnowledgeAnswers_FromProto(mapCtx, in.GetKnowledgeAnswers())
	return out
}
func QueryResult_ToProto(mapCtx *direct.MapContext, in *krm.QueryResult) *pb.QueryResult {
	if in == nil {
		return nil
	}
	out := &pb.QueryResult{}
	out.QueryText = direct.ValueOf(in.QueryText)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.SpeechRecognitionConfidence = direct.ValueOf(in.SpeechRecognitionConfidence)
	out.Action = direct.ValueOf(in.Action)
	out.Parameters = Parameters_ToProto(mapCtx, in.Parameters)
	out.AllRequiredParamsPresent = direct.ValueOf(in.AllRequiredParamsPresent)
	out.CancelsSlotFilling = direct.ValueOf(in.CancelsSlotFilling)
	out.FulfillmentText = direct.ValueOf(in.FulfillmentText)
	out.FulfillmentMessages = direct.Slice_ToProto(mapCtx, in.FulfillmentMessages, Intent_Message_ToProto)
	out.WebhookSource = direct.ValueOf(in.WebhookSource)
	out.WebhookPayload = WebhookPayload_ToProto(mapCtx, in.WebhookPayload)
	out.OutputContexts = direct.Slice_ToProto(mapCtx, in.OutputContexts, Context_ToProto)
	out.Intent = Intent_ToProto(mapCtx, in.Intent)
	out.IntentDetectionConfidence = direct.ValueOf(in.IntentDetectionConfidence)
	out.DiagnosticInfo = DiagnosticInfo_ToProto(mapCtx, in.DiagnosticInfo)
	out.SentimentAnalysisResult = SentimentAnalysisResult_ToProto(mapCtx, in.SentimentAnalysisResult)
	out.KnowledgeAnswers = KnowledgeAnswers_ToProto(mapCtx, in.KnowledgeAnswers)
	return out
}
func QueryResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QueryResult) *krm.QueryResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QueryResultObservedState{}
	// MISSING: QueryText
	// MISSING: LanguageCode
	// MISSING: SpeechRecognitionConfidence
	// MISSING: Action
	// MISSING: Parameters
	// MISSING: AllRequiredParamsPresent
	// MISSING: CancelsSlotFilling
	// MISSING: FulfillmentText
	// MISSING: FulfillmentMessages
	// MISSING: WebhookSource
	// MISSING: WebhookPayload
	// MISSING: OutputContexts
	out.Intent = IntentObservedState_FromProto(mapCtx, in.GetIntent())
	// MISSING: IntentDetectionConfidence
	// MISSING: DiagnosticInfo
	// MISSING: SentimentAnalysisResult
	// MISSING: KnowledgeAnswers
	return out
}
func QueryResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QueryResultObservedState) *pb.QueryResult {
	if in == nil {
		return nil
	}
	out := &pb.QueryResult{}
	// MISSING: QueryText
	// MISSING: LanguageCode
	// MISSING: SpeechRecognitionConfidence
	// MISSING: Action
	// MISSING: Parameters
	// MISSING: AllRequiredParamsPresent
	// MISSING: CancelsSlotFilling
	// MISSING: FulfillmentText
	// MISSING: FulfillmentMessages
	// MISSING: WebhookSource
	// MISSING: WebhookPayload
	// MISSING: OutputContexts
	out.Intent = IntentObservedState_ToProto(mapCtx, in.Intent)
	// MISSING: IntentDetectionConfidence
	// MISSING: DiagnosticInfo
	// MISSING: SentimentAnalysisResult
	// MISSING: KnowledgeAnswers
	return out
}
func Sentiment_FromProto(mapCtx *direct.MapContext, in *pb.Sentiment) *krm.Sentiment {
	if in == nil {
		return nil
	}
	out := &krm.Sentiment{}
	out.Score = direct.LazyPtr(in.GetScore())
	out.Magnitude = direct.LazyPtr(in.GetMagnitude())
	return out
}
func Sentiment_ToProto(mapCtx *direct.MapContext, in *krm.Sentiment) *pb.Sentiment {
	if in == nil {
		return nil
	}
	out := &pb.Sentiment{}
	out.Score = direct.ValueOf(in.Score)
	out.Magnitude = direct.ValueOf(in.Magnitude)
	return out
}
func SentimentAnalysisResult_FromProto(mapCtx *direct.MapContext, in *pb.SentimentAnalysisResult) *krm.SentimentAnalysisResult {
	if in == nil {
		return nil
	}
	out := &krm.SentimentAnalysisResult{}
	out.QueryTextSentiment = Sentiment_FromProto(mapCtx, in.GetQueryTextSentiment())
	return out
}
func SentimentAnalysisResult_ToProto(mapCtx *direct.MapContext, in *krm.SentimentAnalysisResult) *pb.SentimentAnalysisResult {
	if in == nil {
		return nil
	}
	out := &pb.SentimentAnalysisResult{}
	out.QueryTextSentiment = Sentiment_ToProto(mapCtx, in.QueryTextSentiment)
	return out
}
