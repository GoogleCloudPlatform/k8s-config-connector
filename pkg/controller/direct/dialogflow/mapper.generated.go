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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
)
func AutomatedAgentConfig_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedAgentConfig) *krm.AutomatedAgentConfig {
	if in == nil {
		return nil
	}
	out := &krm.AutomatedAgentConfig{}
	out.Agent = direct.LazyPtr(in.GetAgent())
	out.SessionTtl = direct.StringDuration_FromProto(mapCtx, in.GetSessionTtl())
	return out
}
func AutomatedAgentConfig_ToProto(mapCtx *direct.MapContext, in *krm.AutomatedAgentConfig) *pb.AutomatedAgentConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedAgentConfig{}
	out.Agent = direct.ValueOf(in.Agent)
	out.SessionTtl = direct.StringDuration_ToProto(mapCtx, in.SessionTtl)
	return out
}
func ConversationProfile_FromProto(mapCtx *direct.MapContext, in *pb.ConversationProfile) *krm.ConversationProfile {
	if in == nil {
		return nil
	}
	out := &krm.ConversationProfile{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.AutomatedAgentConfig = AutomatedAgentConfig_FromProto(mapCtx, in.GetAutomatedAgentConfig())
	out.HumanAgentAssistantConfig = HumanAgentAssistantConfig_FromProto(mapCtx, in.GetHumanAgentAssistantConfig())
	out.HumanAgentHandoffConfig = HumanAgentHandoffConfig_FromProto(mapCtx, in.GetHumanAgentHandoffConfig())
	out.NotificationConfig = NotificationConfig_FromProto(mapCtx, in.GetNotificationConfig())
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	out.NewMessageEventNotificationConfig = NotificationConfig_FromProto(mapCtx, in.GetNewMessageEventNotificationConfig())
	out.NewRecognitionResultNotificationConfig = NotificationConfig_FromProto(mapCtx, in.GetNewRecognitionResultNotificationConfig())
	out.SttConfig = SpeechToTextConfig_FromProto(mapCtx, in.GetSttConfig())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.SecuritySettings = direct.LazyPtr(in.GetSecuritySettings())
	out.TtsConfig = SynthesizeSpeechConfig_FromProto(mapCtx, in.GetTtsConfig())
	return out
}
func ConversationProfile_ToProto(mapCtx *direct.MapContext, in *krm.ConversationProfile) *pb.ConversationProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConversationProfile{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.AutomatedAgentConfig = AutomatedAgentConfig_ToProto(mapCtx, in.AutomatedAgentConfig)
	out.HumanAgentAssistantConfig = HumanAgentAssistantConfig_ToProto(mapCtx, in.HumanAgentAssistantConfig)
	out.HumanAgentHandoffConfig = HumanAgentHandoffConfig_ToProto(mapCtx, in.HumanAgentHandoffConfig)
	out.NotificationConfig = NotificationConfig_ToProto(mapCtx, in.NotificationConfig)
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	out.NewMessageEventNotificationConfig = NotificationConfig_ToProto(mapCtx, in.NewMessageEventNotificationConfig)
	out.NewRecognitionResultNotificationConfig = NotificationConfig_ToProto(mapCtx, in.NewRecognitionResultNotificationConfig)
	out.SttConfig = SpeechToTextConfig_ToProto(mapCtx, in.SttConfig)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	out.SecuritySettings = direct.ValueOf(in.SecuritySettings)
	out.TtsConfig = SynthesizeSpeechConfig_ToProto(mapCtx, in.TtsConfig)
	return out
}
func ConversationProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversationProfile) *krm.ConversationProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConversationProfileObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: AutomatedAgentConfig
	// MISSING: HumanAgentAssistantConfig
	// MISSING: HumanAgentHandoffConfig
	// MISSING: NotificationConfig
	// MISSING: LoggingConfig
	// MISSING: NewMessageEventNotificationConfig
	// MISSING: NewRecognitionResultNotificationConfig
	// MISSING: SttConfig
	// MISSING: LanguageCode
	// MISSING: TimeZone
	// MISSING: SecuritySettings
	// MISSING: TtsConfig
	return out
}
func ConversationProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConversationProfileObservedState) *pb.ConversationProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConversationProfile{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: AutomatedAgentConfig
	// MISSING: HumanAgentAssistantConfig
	// MISSING: HumanAgentHandoffConfig
	// MISSING: NotificationConfig
	// MISSING: LoggingConfig
	// MISSING: NewMessageEventNotificationConfig
	// MISSING: NewRecognitionResultNotificationConfig
	// MISSING: SttConfig
	// MISSING: LanguageCode
	// MISSING: TimeZone
	// MISSING: SecuritySettings
	// MISSING: TtsConfig
	return out
}
func DialogflowConversationProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversationProfile) *krm.DialogflowConversationProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationProfileObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: AutomatedAgentConfig
	// MISSING: HumanAgentAssistantConfig
	// MISSING: HumanAgentHandoffConfig
	// MISSING: NotificationConfig
	// MISSING: LoggingConfig
	// MISSING: NewMessageEventNotificationConfig
	// MISSING: NewRecognitionResultNotificationConfig
	// MISSING: SttConfig
	// MISSING: LanguageCode
	// MISSING: TimeZone
	// MISSING: SecuritySettings
	// MISSING: TtsConfig
	return out
}
func DialogflowConversationProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationProfileObservedState) *pb.ConversationProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConversationProfile{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: AutomatedAgentConfig
	// MISSING: HumanAgentAssistantConfig
	// MISSING: HumanAgentHandoffConfig
	// MISSING: NotificationConfig
	// MISSING: LoggingConfig
	// MISSING: NewMessageEventNotificationConfig
	// MISSING: NewRecognitionResultNotificationConfig
	// MISSING: SttConfig
	// MISSING: LanguageCode
	// MISSING: TimeZone
	// MISSING: SecuritySettings
	// MISSING: TtsConfig
	return out
}
func DialogflowConversationProfileSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConversationProfile) *krm.DialogflowConversationProfileSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationProfileSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: AutomatedAgentConfig
	// MISSING: HumanAgentAssistantConfig
	// MISSING: HumanAgentHandoffConfig
	// MISSING: NotificationConfig
	// MISSING: LoggingConfig
	// MISSING: NewMessageEventNotificationConfig
	// MISSING: NewRecognitionResultNotificationConfig
	// MISSING: SttConfig
	// MISSING: LanguageCode
	// MISSING: TimeZone
	// MISSING: SecuritySettings
	// MISSING: TtsConfig
	return out
}
func DialogflowConversationProfileSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationProfileSpec) *pb.ConversationProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConversationProfile{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: AutomatedAgentConfig
	// MISSING: HumanAgentAssistantConfig
	// MISSING: HumanAgentHandoffConfig
	// MISSING: NotificationConfig
	// MISSING: LoggingConfig
	// MISSING: NewMessageEventNotificationConfig
	// MISSING: NewRecognitionResultNotificationConfig
	// MISSING: SttConfig
	// MISSING: LanguageCode
	// MISSING: TimeZone
	// MISSING: SecuritySettings
	// MISSING: TtsConfig
	return out
}
func HumanAgentAssistantConfig_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig) *krm.HumanAgentAssistantConfig {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig{}
	out.NotificationConfig = NotificationConfig_FromProto(mapCtx, in.GetNotificationConfig())
	out.HumanAgentSuggestionConfig = HumanAgentAssistantConfig_SuggestionConfig_FromProto(mapCtx, in.GetHumanAgentSuggestionConfig())
	out.EndUserSuggestionConfig = HumanAgentAssistantConfig_SuggestionConfig_FromProto(mapCtx, in.GetEndUserSuggestionConfig())
	out.MessageAnalysisConfig = HumanAgentAssistantConfig_MessageAnalysisConfig_FromProto(mapCtx, in.GetMessageAnalysisConfig())
	return out
}
func HumanAgentAssistantConfig_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig) *pb.HumanAgentAssistantConfig {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig{}
	out.NotificationConfig = NotificationConfig_ToProto(mapCtx, in.NotificationConfig)
	out.HumanAgentSuggestionConfig = HumanAgentAssistantConfig_SuggestionConfig_ToProto(mapCtx, in.HumanAgentSuggestionConfig)
	out.EndUserSuggestionConfig = HumanAgentAssistantConfig_SuggestionConfig_ToProto(mapCtx, in.EndUserSuggestionConfig)
	out.MessageAnalysisConfig = HumanAgentAssistantConfig_MessageAnalysisConfig_ToProto(mapCtx, in.MessageAnalysisConfig)
	return out
}
func HumanAgentAssistantConfig_ConversationModelConfig_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig_ConversationModelConfig) *krm.HumanAgentAssistantConfig_ConversationModelConfig {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig_ConversationModelConfig{}
	out.Model = direct.LazyPtr(in.GetModel())
	out.BaselineModelVersion = direct.LazyPtr(in.GetBaselineModelVersion())
	return out
}
func HumanAgentAssistantConfig_ConversationModelConfig_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig_ConversationModelConfig) *pb.HumanAgentAssistantConfig_ConversationModelConfig {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig_ConversationModelConfig{}
	out.Model = direct.ValueOf(in.Model)
	out.BaselineModelVersion = direct.ValueOf(in.BaselineModelVersion)
	return out
}
func HumanAgentAssistantConfig_ConversationProcessConfig_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig_ConversationProcessConfig) *krm.HumanAgentAssistantConfig_ConversationProcessConfig {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig_ConversationProcessConfig{}
	out.RecentSentencesCount = direct.LazyPtr(in.GetRecentSentencesCount())
	return out
}
func HumanAgentAssistantConfig_ConversationProcessConfig_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig_ConversationProcessConfig) *pb.HumanAgentAssistantConfig_ConversationProcessConfig {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig_ConversationProcessConfig{}
	out.RecentSentencesCount = direct.ValueOf(in.RecentSentencesCount)
	return out
}
func HumanAgentAssistantConfig_MessageAnalysisConfig_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig_MessageAnalysisConfig) *krm.HumanAgentAssistantConfig_MessageAnalysisConfig {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig_MessageAnalysisConfig{}
	out.EnableEntityExtraction = direct.LazyPtr(in.GetEnableEntityExtraction())
	out.EnableSentimentAnalysis = direct.LazyPtr(in.GetEnableSentimentAnalysis())
	return out
}
func HumanAgentAssistantConfig_MessageAnalysisConfig_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig_MessageAnalysisConfig) *pb.HumanAgentAssistantConfig_MessageAnalysisConfig {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig_MessageAnalysisConfig{}
	out.EnableEntityExtraction = direct.ValueOf(in.EnableEntityExtraction)
	out.EnableSentimentAnalysis = direct.ValueOf(in.EnableSentimentAnalysis)
	return out
}
func HumanAgentAssistantConfig_SuggestionConfig_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig_SuggestionConfig) *krm.HumanAgentAssistantConfig_SuggestionConfig {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig_SuggestionConfig{}
	out.FeatureConfigs = direct.Slice_FromProto(mapCtx, in.FeatureConfigs, HumanAgentAssistantConfig_SuggestionFeatureConfig_FromProto)
	out.GroupSuggestionResponses = direct.LazyPtr(in.GetGroupSuggestionResponses())
	out.Generators = in.Generators
	out.DisableHighLatencyFeaturesSyncDelivery = direct.LazyPtr(in.GetDisableHighLatencyFeaturesSyncDelivery())
	return out
}
func HumanAgentAssistantConfig_SuggestionConfig_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig_SuggestionConfig) *pb.HumanAgentAssistantConfig_SuggestionConfig {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig_SuggestionConfig{}
	out.FeatureConfigs = direct.Slice_ToProto(mapCtx, in.FeatureConfigs, HumanAgentAssistantConfig_SuggestionFeatureConfig_ToProto)
	out.GroupSuggestionResponses = direct.ValueOf(in.GroupSuggestionResponses)
	out.Generators = in.Generators
	out.DisableHighLatencyFeaturesSyncDelivery = direct.ValueOf(in.DisableHighLatencyFeaturesSyncDelivery)
	return out
}
func HumanAgentAssistantConfig_SuggestionFeatureConfig_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig_SuggestionFeatureConfig) *krm.HumanAgentAssistantConfig_SuggestionFeatureConfig {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig_SuggestionFeatureConfig{}
	out.SuggestionFeature = SuggestionFeature_FromProto(mapCtx, in.GetSuggestionFeature())
	out.EnableEventBasedSuggestion = direct.LazyPtr(in.GetEnableEventBasedSuggestion())
	out.DisableAgentQueryLogging = direct.LazyPtr(in.GetDisableAgentQueryLogging())
	out.EnableQuerySuggestionWhenNoAnswer = direct.LazyPtr(in.GetEnableQuerySuggestionWhenNoAnswer())
	out.EnableConversationAugmentedQuery = direct.LazyPtr(in.GetEnableConversationAugmentedQuery())
	out.EnableQuerySuggestionOnly = direct.LazyPtr(in.GetEnableQuerySuggestionOnly())
	out.SuggestionTriggerSettings = HumanAgentAssistantConfig_SuggestionTriggerSettings_FromProto(mapCtx, in.GetSuggestionTriggerSettings())
	out.QueryConfig = HumanAgentAssistantConfig_SuggestionQueryConfig_FromProto(mapCtx, in.GetQueryConfig())
	out.ConversationModelConfig = HumanAgentAssistantConfig_ConversationModelConfig_FromProto(mapCtx, in.GetConversationModelConfig())
	out.ConversationProcessConfig = HumanAgentAssistantConfig_ConversationProcessConfig_FromProto(mapCtx, in.GetConversationProcessConfig())
	return out
}
func HumanAgentAssistantConfig_SuggestionFeatureConfig_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig_SuggestionFeatureConfig) *pb.HumanAgentAssistantConfig_SuggestionFeatureConfig {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig_SuggestionFeatureConfig{}
	out.SuggestionFeature = SuggestionFeature_ToProto(mapCtx, in.SuggestionFeature)
	out.EnableEventBasedSuggestion = direct.ValueOf(in.EnableEventBasedSuggestion)
	out.DisableAgentQueryLogging = direct.ValueOf(in.DisableAgentQueryLogging)
	out.EnableQuerySuggestionWhenNoAnswer = direct.ValueOf(in.EnableQuerySuggestionWhenNoAnswer)
	out.EnableConversationAugmentedQuery = direct.ValueOf(in.EnableConversationAugmentedQuery)
	out.EnableQuerySuggestionOnly = direct.ValueOf(in.EnableQuerySuggestionOnly)
	out.SuggestionTriggerSettings = HumanAgentAssistantConfig_SuggestionTriggerSettings_ToProto(mapCtx, in.SuggestionTriggerSettings)
	out.QueryConfig = HumanAgentAssistantConfig_SuggestionQueryConfig_ToProto(mapCtx, in.QueryConfig)
	out.ConversationModelConfig = HumanAgentAssistantConfig_ConversationModelConfig_ToProto(mapCtx, in.ConversationModelConfig)
	out.ConversationProcessConfig = HumanAgentAssistantConfig_ConversationProcessConfig_ToProto(mapCtx, in.ConversationProcessConfig)
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig_SuggestionQueryConfig) *krm.HumanAgentAssistantConfig_SuggestionQueryConfig {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig_SuggestionQueryConfig{}
	out.KnowledgeBaseQuerySource = HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource_FromProto(mapCtx, in.GetKnowledgeBaseQuerySource())
	out.DocumentQuerySource = HumanAgentAssistantConfig_SuggestionQueryConfig_DocumentQuerySource_FromProto(mapCtx, in.GetDocumentQuerySource())
	out.DialogflowQuerySource = HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_FromProto(mapCtx, in.GetDialogflowQuerySource())
	out.MaxResults = direct.LazyPtr(in.GetMaxResults())
	out.ConfidenceThreshold = direct.LazyPtr(in.GetConfidenceThreshold())
	out.ContextFilterSettings = HumanAgentAssistantConfig_SuggestionQueryConfig_ContextFilterSettings_FromProto(mapCtx, in.GetContextFilterSettings())
	out.Sections = HumanAgentAssistantConfig_SuggestionQueryConfig_Sections_FromProto(mapCtx, in.GetSections())
	out.ContextSize = direct.LazyPtr(in.GetContextSize())
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig_SuggestionQueryConfig) *pb.HumanAgentAssistantConfig_SuggestionQueryConfig {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig_SuggestionQueryConfig{}
	if oneof := HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource_ToProto(mapCtx, in.KnowledgeBaseQuerySource); oneof != nil {
		out.QuerySource = &pb.HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource_{KnowledgeBaseQuerySource: oneof}
	}
	if oneof := HumanAgentAssistantConfig_SuggestionQueryConfig_DocumentQuerySource_ToProto(mapCtx, in.DocumentQuerySource); oneof != nil {
		out.QuerySource = &pb.HumanAgentAssistantConfig_SuggestionQueryConfig_DocumentQuerySource_{DocumentQuerySource: oneof}
	}
	if oneof := HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_ToProto(mapCtx, in.DialogflowQuerySource); oneof != nil {
		out.QuerySource = &pb.HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_{DialogflowQuerySource: oneof}
	}
	out.MaxResults = direct.ValueOf(in.MaxResults)
	out.ConfidenceThreshold = direct.ValueOf(in.ConfidenceThreshold)
	out.ContextFilterSettings = HumanAgentAssistantConfig_SuggestionQueryConfig_ContextFilterSettings_ToProto(mapCtx, in.ContextFilterSettings)
	out.Sections = HumanAgentAssistantConfig_SuggestionQueryConfig_Sections_ToProto(mapCtx, in.Sections)
	out.ContextSize = direct.ValueOf(in.ContextSize)
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_ContextFilterSettings_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig_SuggestionQueryConfig_ContextFilterSettings) *krm.HumanAgentAssistantConfig_SuggestionQueryConfig_ContextFilterSettings {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig_SuggestionQueryConfig_ContextFilterSettings{}
	out.DropHandoffMessages = direct.LazyPtr(in.GetDropHandoffMessages())
	out.DropVirtualAgentMessages = direct.LazyPtr(in.GetDropVirtualAgentMessages())
	out.DropIvrMessages = direct.LazyPtr(in.GetDropIvrMessages())
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_ContextFilterSettings_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig_SuggestionQueryConfig_ContextFilterSettings) *pb.HumanAgentAssistantConfig_SuggestionQueryConfig_ContextFilterSettings {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig_SuggestionQueryConfig_ContextFilterSettings{}
	out.DropHandoffMessages = direct.ValueOf(in.DropHandoffMessages)
	out.DropVirtualAgentMessages = direct.ValueOf(in.DropVirtualAgentMessages)
	out.DropIvrMessages = direct.ValueOf(in.DropIvrMessages)
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource) *krm.HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource{}
	out.Agent = direct.LazyPtr(in.GetAgent())
	out.HumanAgentSideConfig = HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_HumanAgentSideConfig_FromProto(mapCtx, in.GetHumanAgentSideConfig())
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource) *pb.HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource{}
	out.Agent = direct.ValueOf(in.Agent)
	out.HumanAgentSideConfig = HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_HumanAgentSideConfig_ToProto(mapCtx, in.HumanAgentSideConfig)
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_HumanAgentSideConfig_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_HumanAgentSideConfig) *krm.HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_HumanAgentSideConfig {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_HumanAgentSideConfig{}
	out.Agent = direct.LazyPtr(in.GetAgent())
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_HumanAgentSideConfig_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_HumanAgentSideConfig) *pb.HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_HumanAgentSideConfig {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig_SuggestionQueryConfig_DialogflowQuerySource_HumanAgentSideConfig{}
	out.Agent = direct.ValueOf(in.Agent)
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_DocumentQuerySource_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig_SuggestionQueryConfig_DocumentQuerySource) *krm.HumanAgentAssistantConfig_SuggestionQueryConfig_DocumentQuerySource {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig_SuggestionQueryConfig_DocumentQuerySource{}
	out.Documents = in.Documents
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_DocumentQuerySource_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig_SuggestionQueryConfig_DocumentQuerySource) *pb.HumanAgentAssistantConfig_SuggestionQueryConfig_DocumentQuerySource {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig_SuggestionQueryConfig_DocumentQuerySource{}
	out.Documents = in.Documents
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource) *krm.HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource{}
	out.KnowledgeBases = in.KnowledgeBases
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource) *pb.HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource{}
	out.KnowledgeBases = in.KnowledgeBases
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_Sections_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig_SuggestionQueryConfig_Sections) *krm.HumanAgentAssistantConfig_SuggestionQueryConfig_Sections {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig_SuggestionQueryConfig_Sections{}
	out.SectionTypes = direct.EnumSlice_FromProto(mapCtx, in.SectionTypes)
	return out
}
func HumanAgentAssistantConfig_SuggestionQueryConfig_Sections_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig_SuggestionQueryConfig_Sections) *pb.HumanAgentAssistantConfig_SuggestionQueryConfig_Sections {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig_SuggestionQueryConfig_Sections{}
	out.SectionTypes = direct.EnumSlice_ToProto[pb.HumanAgentAssistantConfig_SuggestionQueryConfig_Sections_SectionType](mapCtx, in.SectionTypes)
	return out
}
func HumanAgentAssistantConfig_SuggestionTriggerSettings_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentAssistantConfig_SuggestionTriggerSettings) *krm.HumanAgentAssistantConfig_SuggestionTriggerSettings {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentAssistantConfig_SuggestionTriggerSettings{}
	out.NoSmalltalk = direct.LazyPtr(in.GetNoSmalltalk())
	out.OnlyEndUser = direct.LazyPtr(in.GetOnlyEndUser())
	return out
}
func HumanAgentAssistantConfig_SuggestionTriggerSettings_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentAssistantConfig_SuggestionTriggerSettings) *pb.HumanAgentAssistantConfig_SuggestionTriggerSettings {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentAssistantConfig_SuggestionTriggerSettings{}
	out.NoSmalltalk = direct.ValueOf(in.NoSmalltalk)
	out.OnlyEndUser = direct.ValueOf(in.OnlyEndUser)
	return out
}
func HumanAgentHandoffConfig_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentHandoffConfig) *krm.HumanAgentHandoffConfig {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentHandoffConfig{}
	out.LivePersonConfig = HumanAgentHandoffConfig_LivePersonConfig_FromProto(mapCtx, in.GetLivePersonConfig())
	out.SalesforceLiveAgentConfig = HumanAgentHandoffConfig_SalesforceLiveAgentConfig_FromProto(mapCtx, in.GetSalesforceLiveAgentConfig())
	return out
}
func HumanAgentHandoffConfig_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentHandoffConfig) *pb.HumanAgentHandoffConfig {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentHandoffConfig{}
	if oneof := HumanAgentHandoffConfig_LivePersonConfig_ToProto(mapCtx, in.LivePersonConfig); oneof != nil {
		out.AgentService = &pb.HumanAgentHandoffConfig_LivePersonConfig_{LivePersonConfig: oneof}
	}
	if oneof := HumanAgentHandoffConfig_SalesforceLiveAgentConfig_ToProto(mapCtx, in.SalesforceLiveAgentConfig); oneof != nil {
		out.AgentService = &pb.HumanAgentHandoffConfig_SalesforceLiveAgentConfig_{SalesforceLiveAgentConfig: oneof}
	}
	return out
}
func HumanAgentHandoffConfig_LivePersonConfig_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentHandoffConfig_LivePersonConfig) *krm.HumanAgentHandoffConfig_LivePersonConfig {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentHandoffConfig_LivePersonConfig{}
	out.AccountNumber = direct.LazyPtr(in.GetAccountNumber())
	return out
}
func HumanAgentHandoffConfig_LivePersonConfig_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentHandoffConfig_LivePersonConfig) *pb.HumanAgentHandoffConfig_LivePersonConfig {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentHandoffConfig_LivePersonConfig{}
	out.AccountNumber = direct.ValueOf(in.AccountNumber)
	return out
}
func HumanAgentHandoffConfig_SalesforceLiveAgentConfig_FromProto(mapCtx *direct.MapContext, in *pb.HumanAgentHandoffConfig_SalesforceLiveAgentConfig) *krm.HumanAgentHandoffConfig_SalesforceLiveAgentConfig {
	if in == nil {
		return nil
	}
	out := &krm.HumanAgentHandoffConfig_SalesforceLiveAgentConfig{}
	out.OrganizationID = direct.LazyPtr(in.GetOrganizationId())
	out.DeploymentID = direct.LazyPtr(in.GetDeploymentId())
	out.ButtonID = direct.LazyPtr(in.GetButtonId())
	out.EndpointDomain = direct.LazyPtr(in.GetEndpointDomain())
	return out
}
func HumanAgentHandoffConfig_SalesforceLiveAgentConfig_ToProto(mapCtx *direct.MapContext, in *krm.HumanAgentHandoffConfig_SalesforceLiveAgentConfig) *pb.HumanAgentHandoffConfig_SalesforceLiveAgentConfig {
	if in == nil {
		return nil
	}
	out := &pb.HumanAgentHandoffConfig_SalesforceLiveAgentConfig{}
	out.OrganizationId = direct.ValueOf(in.OrganizationID)
	out.DeploymentId = direct.ValueOf(in.DeploymentID)
	out.ButtonId = direct.ValueOf(in.ButtonID)
	out.EndpointDomain = direct.ValueOf(in.EndpointDomain)
	return out
}
func LoggingConfig_FromProto(mapCtx *direct.MapContext, in *pb.LoggingConfig) *krm.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &krm.LoggingConfig{}
	out.EnableStackdriverLogging = direct.LazyPtr(in.GetEnableStackdriverLogging())
	return out
}
func LoggingConfig_ToProto(mapCtx *direct.MapContext, in *krm.LoggingConfig) *pb.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &pb.LoggingConfig{}
	out.EnableStackdriverLogging = direct.ValueOf(in.EnableStackdriverLogging)
	return out
}
func NotificationConfig_FromProto(mapCtx *direct.MapContext, in *pb.NotificationConfig) *krm.NotificationConfig {
	if in == nil {
		return nil
	}
	out := &krm.NotificationConfig{}
	out.Topic = direct.LazyPtr(in.GetTopic())
	out.MessageFormat = direct.Enum_FromProto(mapCtx, in.GetMessageFormat())
	return out
}
func NotificationConfig_ToProto(mapCtx *direct.MapContext, in *krm.NotificationConfig) *pb.NotificationConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotificationConfig{}
	out.Topic = direct.ValueOf(in.Topic)
	out.MessageFormat = direct.Enum_ToProto[pb.NotificationConfig_MessageFormat](mapCtx, in.MessageFormat)
	return out
}
func SpeechToTextConfig_FromProto(mapCtx *direct.MapContext, in *pb.SpeechToTextConfig) *krm.SpeechToTextConfig {
	if in == nil {
		return nil
	}
	out := &krm.SpeechToTextConfig{}
	out.SpeechModelVariant = direct.Enum_FromProto(mapCtx, in.GetSpeechModelVariant())
	out.Model = direct.LazyPtr(in.GetModel())
	out.PhraseSets = in.PhraseSets
	out.AudioEncoding = direct.Enum_FromProto(mapCtx, in.GetAudioEncoding())
	out.SampleRateHertz = direct.LazyPtr(in.GetSampleRateHertz())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.EnableWordInfo = direct.LazyPtr(in.GetEnableWordInfo())
	out.UseTimeoutBasedEndpointing = direct.LazyPtr(in.GetUseTimeoutBasedEndpointing())
	return out
}
func SpeechToTextConfig_ToProto(mapCtx *direct.MapContext, in *krm.SpeechToTextConfig) *pb.SpeechToTextConfig {
	if in == nil {
		return nil
	}
	out := &pb.SpeechToTextConfig{}
	out.SpeechModelVariant = direct.Enum_ToProto[pb.SpeechModelVariant](mapCtx, in.SpeechModelVariant)
	out.Model = direct.ValueOf(in.Model)
	out.PhraseSets = in.PhraseSets
	out.AudioEncoding = direct.Enum_ToProto[pb.AudioEncoding](mapCtx, in.AudioEncoding)
	out.SampleRateHertz = direct.ValueOf(in.SampleRateHertz)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.EnableWordInfo = direct.ValueOf(in.EnableWordInfo)
	out.UseTimeoutBasedEndpointing = direct.ValueOf(in.UseTimeoutBasedEndpointing)
	return out
}
func SuggestionFeature_FromProto(mapCtx *direct.MapContext, in *pb.SuggestionFeature) *krm.SuggestionFeature {
	if in == nil {
		return nil
	}
	out := &krm.SuggestionFeature{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func SuggestionFeature_ToProto(mapCtx *direct.MapContext, in *krm.SuggestionFeature) *pb.SuggestionFeature {
	if in == nil {
		return nil
	}
	out := &pb.SuggestionFeature{}
	out.Type = direct.Enum_ToProto[pb.SuggestionFeature_Type](mapCtx, in.Type)
	return out
}
func SynthesizeSpeechConfig_FromProto(mapCtx *direct.MapContext, in *pb.SynthesizeSpeechConfig) *krm.SynthesizeSpeechConfig {
	if in == nil {
		return nil
	}
	out := &krm.SynthesizeSpeechConfig{}
	out.SpeakingRate = direct.LazyPtr(in.GetSpeakingRate())
	out.Pitch = direct.LazyPtr(in.GetPitch())
	out.VolumeGainDb = direct.LazyPtr(in.GetVolumeGainDb())
	out.EffectsProfileID = in.EffectsProfileId
	out.Voice = VoiceSelectionParams_FromProto(mapCtx, in.GetVoice())
	return out
}
func SynthesizeSpeechConfig_ToProto(mapCtx *direct.MapContext, in *krm.SynthesizeSpeechConfig) *pb.SynthesizeSpeechConfig {
	if in == nil {
		return nil
	}
	out := &pb.SynthesizeSpeechConfig{}
	out.SpeakingRate = direct.ValueOf(in.SpeakingRate)
	out.Pitch = direct.ValueOf(in.Pitch)
	out.VolumeGainDb = direct.ValueOf(in.VolumeGainDb)
	out.EffectsProfileId = in.EffectsProfileID
	out.Voice = VoiceSelectionParams_ToProto(mapCtx, in.Voice)
	return out
}
func VoiceSelectionParams_FromProto(mapCtx *direct.MapContext, in *pb.VoiceSelectionParams) *krm.VoiceSelectionParams {
	if in == nil {
		return nil
	}
	out := &krm.VoiceSelectionParams{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SsmlGender = direct.Enum_FromProto(mapCtx, in.GetSsmlGender())
	return out
}
func VoiceSelectionParams_ToProto(mapCtx *direct.MapContext, in *krm.VoiceSelectionParams) *pb.VoiceSelectionParams {
	if in == nil {
		return nil
	}
	out := &pb.VoiceSelectionParams{}
	out.Name = direct.ValueOf(in.Name)
	out.SsmlGender = direct.Enum_ToProto[pb.SsmlVoiceGender](mapCtx, in.SsmlGender)
	return out
}
