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
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
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
func ContactcenterinsightsSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.ContactcenterinsightsSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsSettingsObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LanguageCode
	// MISSING: ConversationTtl
	// MISSING: PubsubNotificationSettings
	// MISSING: AnalysisConfig
	// MISSING: RedactionConfig
	// MISSING: SpeechConfig
	return out
}
func ContactcenterinsightsSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsSettingsObservedState) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LanguageCode
	// MISSING: ConversationTtl
	// MISSING: PubsubNotificationSettings
	// MISSING: AnalysisConfig
	// MISSING: RedactionConfig
	// MISSING: SpeechConfig
	return out
}
func ContactcenterinsightsSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.ContactcenterinsightsSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsSettingsSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LanguageCode
	// MISSING: ConversationTtl
	// MISSING: PubsubNotificationSettings
	// MISSING: AnalysisConfig
	// MISSING: RedactionConfig
	// MISSING: SpeechConfig
	return out
}
func ContactcenterinsightsSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsSettingsSpec) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LanguageCode
	// MISSING: ConversationTtl
	// MISSING: PubsubNotificationSettings
	// MISSING: AnalysisConfig
	// MISSING: RedactionConfig
	// MISSING: SpeechConfig
	return out
}
func RedactionConfig_FromProto(mapCtx *direct.MapContext, in *pb.RedactionConfig) *krm.RedactionConfig {
	if in == nil {
		return nil
	}
	out := &krm.RedactionConfig{}
	out.DeidentifyTemplate = direct.LazyPtr(in.GetDeidentifyTemplate())
	out.InspectTemplate = direct.LazyPtr(in.GetInspectTemplate())
	return out
}
func RedactionConfig_ToProto(mapCtx *direct.MapContext, in *krm.RedactionConfig) *pb.RedactionConfig {
	if in == nil {
		return nil
	}
	out := &pb.RedactionConfig{}
	out.DeidentifyTemplate = direct.ValueOf(in.DeidentifyTemplate)
	out.InspectTemplate = direct.ValueOf(in.InspectTemplate)
	return out
}
func Settings_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.Settings {
	if in == nil {
		return nil
	}
	out := &krm.Settings{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.ConversationTtl = direct.StringDuration_FromProto(mapCtx, in.GetConversationTtl())
	out.PubsubNotificationSettings = in.PubsubNotificationSettings
	out.AnalysisConfig = Settings_AnalysisConfig_FromProto(mapCtx, in.GetAnalysisConfig())
	out.RedactionConfig = RedactionConfig_FromProto(mapCtx, in.GetRedactionConfig())
	out.SpeechConfig = SpeechConfig_FromProto(mapCtx, in.GetSpeechConfig())
	return out
}
func Settings_ToProto(mapCtx *direct.MapContext, in *krm.Settings) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.ConversationTtl = direct.StringDuration_ToProto(mapCtx, in.ConversationTtl)
	out.PubsubNotificationSettings = in.PubsubNotificationSettings
	out.AnalysisConfig = Settings_AnalysisConfig_ToProto(mapCtx, in.AnalysisConfig)
	out.RedactionConfig = RedactionConfig_ToProto(mapCtx, in.RedactionConfig)
	out.SpeechConfig = SpeechConfig_ToProto(mapCtx, in.SpeechConfig)
	return out
}
func SettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.SettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SettingsObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: LanguageCode
	// MISSING: ConversationTtl
	// MISSING: PubsubNotificationSettings
	// MISSING: AnalysisConfig
	// MISSING: RedactionConfig
	// MISSING: SpeechConfig
	return out
}
func SettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SettingsObservedState) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: LanguageCode
	// MISSING: ConversationTtl
	// MISSING: PubsubNotificationSettings
	// MISSING: AnalysisConfig
	// MISSING: RedactionConfig
	// MISSING: SpeechConfig
	return out
}
func Settings_AnalysisConfig_FromProto(mapCtx *direct.MapContext, in *pb.Settings_AnalysisConfig) *krm.Settings_AnalysisConfig {
	if in == nil {
		return nil
	}
	out := &krm.Settings_AnalysisConfig{}
	out.RuntimeIntegrationAnalysisPercentage = direct.LazyPtr(in.GetRuntimeIntegrationAnalysisPercentage())
	out.UploadConversationAnalysisPercentage = direct.LazyPtr(in.GetUploadConversationAnalysisPercentage())
	out.AnnotatorSelector = AnnotatorSelector_FromProto(mapCtx, in.GetAnnotatorSelector())
	return out
}
func Settings_AnalysisConfig_ToProto(mapCtx *direct.MapContext, in *krm.Settings_AnalysisConfig) *pb.Settings_AnalysisConfig {
	if in == nil {
		return nil
	}
	out := &pb.Settings_AnalysisConfig{}
	out.RuntimeIntegrationAnalysisPercentage = direct.ValueOf(in.RuntimeIntegrationAnalysisPercentage)
	out.UploadConversationAnalysisPercentage = direct.ValueOf(in.UploadConversationAnalysisPercentage)
	out.AnnotatorSelector = AnnotatorSelector_ToProto(mapCtx, in.AnnotatorSelector)
	return out
}
func SpeechConfig_FromProto(mapCtx *direct.MapContext, in *pb.SpeechConfig) *krm.SpeechConfig {
	if in == nil {
		return nil
	}
	out := &krm.SpeechConfig{}
	out.SpeechRecognizer = direct.LazyPtr(in.GetSpeechRecognizer())
	return out
}
func SpeechConfig_ToProto(mapCtx *direct.MapContext, in *krm.SpeechConfig) *pb.SpeechConfig {
	if in == nil {
		return nil
	}
	out := &pb.SpeechConfig{}
	out.SpeechRecognizer = direct.ValueOf(in.SpeechRecognizer)
	return out
}
