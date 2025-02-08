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
	pb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Action_FromProto(mapCtx *direct.MapContext, in *pb.Action) *krm.Action {
	if in == nil {
		return nil
	}
	out := &krm.Action{}
	out.UserUtterance = UserUtterance_FromProto(mapCtx, in.GetUserUtterance())
	out.AgentUtterance = AgentUtterance_FromProto(mapCtx, in.GetAgentUtterance())
	out.ToolUse = ToolUse_FromProto(mapCtx, in.GetToolUse())
	out.PlaybookInvocation = PlaybookInvocation_FromProto(mapCtx, in.GetPlaybookInvocation())
	out.FlowInvocation = FlowInvocation_FromProto(mapCtx, in.GetFlowInvocation())
	return out
}
func Action_ToProto(mapCtx *direct.MapContext, in *krm.Action) *pb.Action {
	if in == nil {
		return nil
	}
	out := &pb.Action{}
	if oneof := UserUtterance_ToProto(mapCtx, in.UserUtterance); oneof != nil {
		out.Action = &pb.Action_UserUtterance{UserUtterance: oneof}
	}
	if oneof := AgentUtterance_ToProto(mapCtx, in.AgentUtterance); oneof != nil {
		out.Action = &pb.Action_AgentUtterance{AgentUtterance: oneof}
	}
	if oneof := ToolUse_ToProto(mapCtx, in.ToolUse); oneof != nil {
		out.Action = &pb.Action_ToolUse{ToolUse: oneof}
	}
	if oneof := PlaybookInvocation_ToProto(mapCtx, in.PlaybookInvocation); oneof != nil {
		out.Action = &pb.Action_PlaybookInvocation{PlaybookInvocation: oneof}
	}
	if oneof := FlowInvocation_ToProto(mapCtx, in.FlowInvocation); oneof != nil {
		out.Action = &pb.Action_FlowInvocation{FlowInvocation: oneof}
	}
	return out
}
func AdvancedSettings_FromProto(mapCtx *direct.MapContext, in *pb.AdvancedSettings) *krm.AdvancedSettings {
	if in == nil {
		return nil
	}
	out := &krm.AdvancedSettings{}
	out.AudioExportGcsDestination = GcsDestination_FromProto(mapCtx, in.GetAudioExportGcsDestination())
	out.SpeechSettings = AdvancedSettings_SpeechSettings_FromProto(mapCtx, in.GetSpeechSettings())
	out.DtmfSettings = AdvancedSettings_DtmfSettings_FromProto(mapCtx, in.GetDtmfSettings())
	out.LoggingSettings = AdvancedSettings_LoggingSettings_FromProto(mapCtx, in.GetLoggingSettings())
	return out
}
func AdvancedSettings_ToProto(mapCtx *direct.MapContext, in *krm.AdvancedSettings) *pb.AdvancedSettings {
	if in == nil {
		return nil
	}
	out := &pb.AdvancedSettings{}
	out.AudioExportGcsDestination = GcsDestination_ToProto(mapCtx, in.AudioExportGcsDestination)
	out.SpeechSettings = AdvancedSettings_SpeechSettings_ToProto(mapCtx, in.SpeechSettings)
	out.DtmfSettings = AdvancedSettings_DtmfSettings_ToProto(mapCtx, in.DtmfSettings)
	out.LoggingSettings = AdvancedSettings_LoggingSettings_ToProto(mapCtx, in.LoggingSettings)
	return out
}
func AdvancedSettings_DtmfSettings_FromProto(mapCtx *direct.MapContext, in *pb.AdvancedSettings_DtmfSettings) *krm.AdvancedSettings_DtmfSettings {
	if in == nil {
		return nil
	}
	out := &krm.AdvancedSettings_DtmfSettings{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.MaxDigits = direct.LazyPtr(in.GetMaxDigits())
	out.FinishDigit = direct.LazyPtr(in.GetFinishDigit())
	out.InterdigitTimeoutDuration = direct.StringDuration_FromProto(mapCtx, in.GetInterdigitTimeoutDuration())
	out.EndpointingTimeoutDuration = direct.StringDuration_FromProto(mapCtx, in.GetEndpointingTimeoutDuration())
	return out
}
func AdvancedSettings_DtmfSettings_ToProto(mapCtx *direct.MapContext, in *krm.AdvancedSettings_DtmfSettings) *pb.AdvancedSettings_DtmfSettings {
	if in == nil {
		return nil
	}
	out := &pb.AdvancedSettings_DtmfSettings{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.MaxDigits = direct.ValueOf(in.MaxDigits)
	out.FinishDigit = direct.ValueOf(in.FinishDigit)
	out.InterdigitTimeoutDuration = direct.StringDuration_ToProto(mapCtx, in.InterdigitTimeoutDuration)
	out.EndpointingTimeoutDuration = direct.StringDuration_ToProto(mapCtx, in.EndpointingTimeoutDuration)
	return out
}
func AdvancedSettings_LoggingSettings_FromProto(mapCtx *direct.MapContext, in *pb.AdvancedSettings_LoggingSettings) *krm.AdvancedSettings_LoggingSettings {
	if in == nil {
		return nil
	}
	out := &krm.AdvancedSettings_LoggingSettings{}
	out.EnableStackdriverLogging = direct.LazyPtr(in.GetEnableStackdriverLogging())
	out.EnableInteractionLogging = direct.LazyPtr(in.GetEnableInteractionLogging())
	out.EnableConsentBasedRedaction = direct.LazyPtr(in.GetEnableConsentBasedRedaction())
	return out
}
func AdvancedSettings_LoggingSettings_ToProto(mapCtx *direct.MapContext, in *krm.AdvancedSettings_LoggingSettings) *pb.AdvancedSettings_LoggingSettings {
	if in == nil {
		return nil
	}
	out := &pb.AdvancedSettings_LoggingSettings{}
	out.EnableStackdriverLogging = direct.ValueOf(in.EnableStackdriverLogging)
	out.EnableInteractionLogging = direct.ValueOf(in.EnableInteractionLogging)
	out.EnableConsentBasedRedaction = direct.ValueOf(in.EnableConsentBasedRedaction)
	return out
}
func AdvancedSettings_SpeechSettings_FromProto(mapCtx *direct.MapContext, in *pb.AdvancedSettings_SpeechSettings) *krm.AdvancedSettings_SpeechSettings {
	if in == nil {
		return nil
	}
	out := &krm.AdvancedSettings_SpeechSettings{}
	out.EndpointerSensitivity = direct.LazyPtr(in.GetEndpointerSensitivity())
	out.NoSpeechTimeout = direct.StringDuration_FromProto(mapCtx, in.GetNoSpeechTimeout())
	out.UseTimeoutBasedEndpointing = direct.LazyPtr(in.GetUseTimeoutBasedEndpointing())
	out.Models = in.Models
	return out
}
func AdvancedSettings_SpeechSettings_ToProto(mapCtx *direct.MapContext, in *krm.AdvancedSettings_SpeechSettings) *pb.AdvancedSettings_SpeechSettings {
	if in == nil {
		return nil
	}
	out := &pb.AdvancedSettings_SpeechSettings{}
	out.EndpointerSensitivity = direct.ValueOf(in.EndpointerSensitivity)
	out.NoSpeechTimeout = direct.StringDuration_ToProto(mapCtx, in.NoSpeechTimeout)
	out.UseTimeoutBasedEndpointing = direct.ValueOf(in.UseTimeoutBasedEndpointing)
	out.Models = in.Models
	return out
}
func AgentUtterance_FromProto(mapCtx *direct.MapContext, in *pb.AgentUtterance) *krm.AgentUtterance {
	if in == nil {
		return nil
	}
	out := &krm.AgentUtterance{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func AgentUtterance_ToProto(mapCtx *direct.MapContext, in *krm.AgentUtterance) *pb.AgentUtterance {
	if in == nil {
		return nil
	}
	out := &pb.AgentUtterance{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
func AudioInput_FromProto(mapCtx *direct.MapContext, in *pb.AudioInput) *krm.AudioInput {
	if in == nil {
		return nil
	}
	out := &krm.AudioInput{}
	out.Config = InputAudioConfig_FromProto(mapCtx, in.GetConfig())
	out.Audio = in.GetAudio()
	return out
}
func AudioInput_ToProto(mapCtx *direct.MapContext, in *krm.AudioInput) *pb.AudioInput {
	if in == nil {
		return nil
	}
	out := &pb.AudioInput{}
	out.Config = InputAudioConfig_ToProto(mapCtx, in.Config)
	out.Audio = in.Audio
	return out
}
func BargeInConfig_FromProto(mapCtx *direct.MapContext, in *pb.BargeInConfig) *krm.BargeInConfig {
	if in == nil {
		return nil
	}
	out := &krm.BargeInConfig{}
	out.NoBargeInDuration = direct.StringDuration_FromProto(mapCtx, in.GetNoBargeInDuration())
	out.TotalDuration = direct.StringDuration_FromProto(mapCtx, in.GetTotalDuration())
	return out
}
func BargeInConfig_ToProto(mapCtx *direct.MapContext, in *krm.BargeInConfig) *pb.BargeInConfig {
	if in == nil {
		return nil
	}
	out := &pb.BargeInConfig{}
	out.NoBargeInDuration = direct.StringDuration_ToProto(mapCtx, in.NoBargeInDuration)
	out.TotalDuration = direct.StringDuration_ToProto(mapCtx, in.TotalDuration)
	return out
}
func BoostSpec_FromProto(mapCtx *direct.MapContext, in *pb.BoostSpec) *krm.BoostSpec {
	if in == nil {
		return nil
	}
	out := &krm.BoostSpec{}
	out.ConditionBoostSpecs = direct.Slice_FromProto(mapCtx, in.ConditionBoostSpecs, BoostSpec_ConditionBoostSpec_FromProto)
	return out
}
func BoostSpec_ToProto(mapCtx *direct.MapContext, in *krm.BoostSpec) *pb.BoostSpec {
	if in == nil {
		return nil
	}
	out := &pb.BoostSpec{}
	out.ConditionBoostSpecs = direct.Slice_ToProto(mapCtx, in.ConditionBoostSpecs, BoostSpec_ConditionBoostSpec_ToProto)
	return out
}
func BoostSpec_ConditionBoostSpec_FromProto(mapCtx *direct.MapContext, in *pb.BoostSpec_ConditionBoostSpec) *krm.BoostSpec_ConditionBoostSpec {
	if in == nil {
		return nil
	}
	out := &krm.BoostSpec_ConditionBoostSpec{}
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.Boost = direct.LazyPtr(in.GetBoost())
	out.BoostControlSpec = BoostSpec_ConditionBoostSpec_BoostControlSpec_FromProto(mapCtx, in.GetBoostControlSpec())
	return out
}
func BoostSpec_ConditionBoostSpec_ToProto(mapCtx *direct.MapContext, in *krm.BoostSpec_ConditionBoostSpec) *pb.BoostSpec_ConditionBoostSpec {
	if in == nil {
		return nil
	}
	out := &pb.BoostSpec_ConditionBoostSpec{}
	out.Condition = direct.ValueOf(in.Condition)
	out.Boost = direct.ValueOf(in.Boost)
	out.BoostControlSpec = BoostSpec_ConditionBoostSpec_BoostControlSpec_ToProto(mapCtx, in.BoostControlSpec)
	return out
}
func BoostSpec_ConditionBoostSpec_BoostControlSpec_FromProto(mapCtx *direct.MapContext, in *pb.BoostSpec_ConditionBoostSpec_BoostControlSpec) *krm.BoostSpec_ConditionBoostSpec_BoostControlSpec {
	if in == nil {
		return nil
	}
	out := &krm.BoostSpec_ConditionBoostSpec_BoostControlSpec{}
	out.FieldName = direct.LazyPtr(in.GetFieldName())
	out.AttributeType = direct.Enum_FromProto(mapCtx, in.GetAttributeType())
	out.InterpolationType = direct.Enum_FromProto(mapCtx, in.GetInterpolationType())
	out.ControlPoints = direct.Slice_FromProto(mapCtx, in.ControlPoints, BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint_FromProto)
	return out
}
func BoostSpec_ConditionBoostSpec_BoostControlSpec_ToProto(mapCtx *direct.MapContext, in *krm.BoostSpec_ConditionBoostSpec_BoostControlSpec) *pb.BoostSpec_ConditionBoostSpec_BoostControlSpec {
	if in == nil {
		return nil
	}
	out := &pb.BoostSpec_ConditionBoostSpec_BoostControlSpec{}
	out.FieldName = direct.ValueOf(in.FieldName)
	out.AttributeType = direct.Enum_ToProto[pb.BoostSpec_ConditionBoostSpec_BoostControlSpec_AttributeType](mapCtx, in.AttributeType)
	out.InterpolationType = direct.Enum_ToProto[pb.BoostSpec_ConditionBoostSpec_BoostControlSpec_InterpolationType](mapCtx, in.InterpolationType)
	out.ControlPoints = direct.Slice_ToProto(mapCtx, in.ControlPoints, BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint_ToProto)
	return out
}
func BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint_FromProto(mapCtx *direct.MapContext, in *pb.BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint) *krm.BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint {
	if in == nil {
		return nil
	}
	out := &krm.BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint{}
	out.AttributeValue = direct.LazyPtr(in.GetAttributeValue())
	out.BoostAmount = direct.LazyPtr(in.GetBoostAmount())
	return out
}
func BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint_ToProto(mapCtx *direct.MapContext, in *krm.BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint) *pb.BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint {
	if in == nil {
		return nil
	}
	out := &pb.BoostSpec_ConditionBoostSpec_BoostControlSpec_ControlPoint{}
	out.AttributeValue = direct.ValueOf(in.AttributeValue)
	out.BoostAmount = direct.ValueOf(in.BoostAmount)
	return out
}
func BoostSpecs_FromProto(mapCtx *direct.MapContext, in *pb.BoostSpecs) *krm.BoostSpecs {
	if in == nil {
		return nil
	}
	out := &krm.BoostSpecs{}
	out.DataStores = in.DataStores
	out.Spec = direct.Slice_FromProto(mapCtx, in.Spec, BoostSpec_FromProto)
	return out
}
func BoostSpecs_ToProto(mapCtx *direct.MapContext, in *krm.BoostSpecs) *pb.BoostSpecs {
	if in == nil {
		return nil
	}
	out := &pb.BoostSpecs{}
	out.DataStores = in.DataStores
	out.Spec = direct.Slice_ToProto(mapCtx, in.Spec, BoostSpec_ToProto)
	return out
}
func Conversation_FromProto(mapCtx *direct.MapContext, in *pb.Conversation) *krm.Conversation {
	if in == nil {
		return nil
	}
	out := &krm.Conversation{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.Metrics = Conversation_Metrics_FromProto(mapCtx, in.GetMetrics())
	out.Intents = direct.Slice_FromProto(mapCtx, in.Intents, Intent_FromProto)
	out.Flows = direct.Slice_FromProto(mapCtx, in.Flows, Flow_FromProto)
	out.Pages = direct.Slice_FromProto(mapCtx, in.Pages, Page_FromProto)
	out.Interactions = direct.Slice_FromProto(mapCtx, in.Interactions, Conversation_Interaction_FromProto)
	out.Environment = Environment_FromProto(mapCtx, in.GetEnvironment())
	out.FlowVersions = in.FlowVersions
	return out
}
func Conversation_ToProto(mapCtx *direct.MapContext, in *krm.Conversation) *pb.Conversation {
	if in == nil {
		return nil
	}
	out := &pb.Conversation{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.Enum_ToProto[pb.Conversation_Type](mapCtx, in.Type)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	out.Metrics = Conversation_Metrics_ToProto(mapCtx, in.Metrics)
	out.Intents = direct.Slice_ToProto(mapCtx, in.Intents, Intent_ToProto)
	out.Flows = direct.Slice_ToProto(mapCtx, in.Flows, Flow_ToProto)
	out.Pages = direct.Slice_ToProto(mapCtx, in.Pages, Page_ToProto)
	out.Interactions = direct.Slice_ToProto(mapCtx, in.Interactions, Conversation_Interaction_ToProto)
	out.Environment = Environment_ToProto(mapCtx, in.Environment)
	out.FlowVersions = in.FlowVersions
	return out
}
func ConversationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Conversation) *krm.ConversationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConversationObservedState{}
	// MISSING: Name
	// MISSING: Type
	// MISSING: LanguageCode
	// MISSING: StartTime
	// MISSING: Duration
	// MISSING: Metrics
	// MISSING: Intents
	out.Flows = direct.Slice_FromProto(mapCtx, in.Flows, FlowObservedState_FromProto)
	// MISSING: Pages
	out.Interactions = direct.Slice_FromProto(mapCtx, in.Interactions, Conversation_InteractionObservedState_FromProto)
	out.Environment = EnvironmentObservedState_FromProto(mapCtx, in.GetEnvironment())
	// MISSING: FlowVersions
	return out
}
func ConversationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConversationObservedState) *pb.Conversation {
	if in == nil {
		return nil
	}
	out := &pb.Conversation{}
	// MISSING: Name
	// MISSING: Type
	// MISSING: LanguageCode
	// MISSING: StartTime
	// MISSING: Duration
	// MISSING: Metrics
	// MISSING: Intents
	out.Flows = direct.Slice_ToProto(mapCtx, in.Flows, FlowObservedState_ToProto)
	// MISSING: Pages
	out.Interactions = direct.Slice_ToProto(mapCtx, in.Interactions, Conversation_InteractionObservedState_ToProto)
	out.Environment = EnvironmentObservedState_ToProto(mapCtx, in.Environment)
	// MISSING: FlowVersions
	return out
}
func Conversation_Interaction_FromProto(mapCtx *direct.MapContext, in *pb.Conversation_Interaction) *krm.Conversation_Interaction {
	if in == nil {
		return nil
	}
	out := &krm.Conversation_Interaction{}
	out.Request = DetectIntentRequest_FromProto(mapCtx, in.GetRequest())
	out.Response = DetectIntentResponse_FromProto(mapCtx, in.GetResponse())
	out.PartialResponses = direct.Slice_FromProto(mapCtx, in.PartialResponses, DetectIntentResponse_FromProto)
	out.RequestUtterances = direct.LazyPtr(in.GetRequestUtterances())
	out.ResponseUtterances = direct.LazyPtr(in.GetResponseUtterances())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.MissingTransition = Conversation_Interaction_MissingTransition_FromProto(mapCtx, in.GetMissingTransition())
	return out
}
func Conversation_Interaction_ToProto(mapCtx *direct.MapContext, in *krm.Conversation_Interaction) *pb.Conversation_Interaction {
	if in == nil {
		return nil
	}
	out := &pb.Conversation_Interaction{}
	out.Request = DetectIntentRequest_ToProto(mapCtx, in.Request)
	out.Response = DetectIntentResponse_ToProto(mapCtx, in.Response)
	out.PartialResponses = direct.Slice_ToProto(mapCtx, in.PartialResponses, DetectIntentResponse_ToProto)
	out.RequestUtterances = direct.ValueOf(in.RequestUtterances)
	out.ResponseUtterances = direct.ValueOf(in.ResponseUtterances)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.MissingTransition = Conversation_Interaction_MissingTransition_ToProto(mapCtx, in.MissingTransition)
	return out
}
func Conversation_InteractionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Conversation_Interaction) *krm.Conversation_InteractionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Conversation_InteractionObservedState{}
	// MISSING: Request
	out.Response = DetectIntentResponseObservedState_FromProto(mapCtx, in.GetResponse())
	// MISSING: PartialResponses
	// MISSING: RequestUtterances
	// MISSING: ResponseUtterances
	// MISSING: CreateTime
	// MISSING: MissingTransition
	return out
}
func Conversation_InteractionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Conversation_InteractionObservedState) *pb.Conversation_Interaction {
	if in == nil {
		return nil
	}
	out := &pb.Conversation_Interaction{}
	// MISSING: Request
	out.Response = DetectIntentResponseObservedState_ToProto(mapCtx, in.Response)
	// MISSING: PartialResponses
	// MISSING: RequestUtterances
	// MISSING: ResponseUtterances
	// MISSING: CreateTime
	// MISSING: MissingTransition
	return out
}
func Conversation_Interaction_MissingTransition_FromProto(mapCtx *direct.MapContext, in *pb.Conversation_Interaction_MissingTransition) *krm.Conversation_Interaction_MissingTransition {
	if in == nil {
		return nil
	}
	out := &krm.Conversation_Interaction_MissingTransition{}
	out.IntentDisplayName = direct.LazyPtr(in.GetIntentDisplayName())
	out.Score = direct.LazyPtr(in.GetScore())
	return out
}
func Conversation_Interaction_MissingTransition_ToProto(mapCtx *direct.MapContext, in *krm.Conversation_Interaction_MissingTransition) *pb.Conversation_Interaction_MissingTransition {
	if in == nil {
		return nil
	}
	out := &pb.Conversation_Interaction_MissingTransition{}
	out.IntentDisplayName = direct.ValueOf(in.IntentDisplayName)
	out.Score = direct.ValueOf(in.Score)
	return out
}
func Conversation_Metrics_FromProto(mapCtx *direct.MapContext, in *pb.Conversation_Metrics) *krm.Conversation_Metrics {
	if in == nil {
		return nil
	}
	out := &krm.Conversation_Metrics{}
	out.InteractionCount = direct.LazyPtr(in.GetInteractionCount())
	out.InputAudioDuration = direct.StringDuration_FromProto(mapCtx, in.GetInputAudioDuration())
	out.OutputAudioDuration = direct.StringDuration_FromProto(mapCtx, in.GetOutputAudioDuration())
	out.MaxWebhookLatency = direct.StringDuration_FromProto(mapCtx, in.GetMaxWebhookLatency())
	out.HasEndInteraction = direct.LazyPtr(in.GetHasEndInteraction())
	out.HasLiveAgentHandoff = direct.LazyPtr(in.GetHasLiveAgentHandoff())
	out.AverageMatchConfidence = direct.LazyPtr(in.GetAverageMatchConfidence())
	out.QueryInputCount = Conversation_Metrics_QueryInputCount_FromProto(mapCtx, in.GetQueryInputCount())
	out.MatchTypeCount = Conversation_Metrics_MatchTypeCount_FromProto(mapCtx, in.GetMatchTypeCount())
	return out
}
func Conversation_Metrics_ToProto(mapCtx *direct.MapContext, in *krm.Conversation_Metrics) *pb.Conversation_Metrics {
	if in == nil {
		return nil
	}
	out := &pb.Conversation_Metrics{}
	out.InteractionCount = direct.ValueOf(in.InteractionCount)
	out.InputAudioDuration = direct.StringDuration_ToProto(mapCtx, in.InputAudioDuration)
	out.OutputAudioDuration = direct.StringDuration_ToProto(mapCtx, in.OutputAudioDuration)
	out.MaxWebhookLatency = direct.StringDuration_ToProto(mapCtx, in.MaxWebhookLatency)
	out.HasEndInteraction = direct.ValueOf(in.HasEndInteraction)
	out.HasLiveAgentHandoff = direct.ValueOf(in.HasLiveAgentHandoff)
	out.AverageMatchConfidence = direct.ValueOf(in.AverageMatchConfidence)
	out.QueryInputCount = Conversation_Metrics_QueryInputCount_ToProto(mapCtx, in.QueryInputCount)
	out.MatchTypeCount = Conversation_Metrics_MatchTypeCount_ToProto(mapCtx, in.MatchTypeCount)
	return out
}
func Conversation_Metrics_MatchTypeCount_FromProto(mapCtx *direct.MapContext, in *pb.Conversation_Metrics_MatchTypeCount) *krm.Conversation_Metrics_MatchTypeCount {
	if in == nil {
		return nil
	}
	out := &krm.Conversation_Metrics_MatchTypeCount{}
	out.UnspecifiedCount = direct.LazyPtr(in.GetUnspecifiedCount())
	out.IntentCount = direct.LazyPtr(in.GetIntentCount())
	out.DirectIntentCount = direct.LazyPtr(in.GetDirectIntentCount())
	out.ParameterFillingCount = direct.LazyPtr(in.GetParameterFillingCount())
	out.NoMatchCount = direct.LazyPtr(in.GetNoMatchCount())
	out.NoInputCount = direct.LazyPtr(in.GetNoInputCount())
	out.EventCount = direct.LazyPtr(in.GetEventCount())
	return out
}
func Conversation_Metrics_MatchTypeCount_ToProto(mapCtx *direct.MapContext, in *krm.Conversation_Metrics_MatchTypeCount) *pb.Conversation_Metrics_MatchTypeCount {
	if in == nil {
		return nil
	}
	out := &pb.Conversation_Metrics_MatchTypeCount{}
	out.UnspecifiedCount = direct.ValueOf(in.UnspecifiedCount)
	out.IntentCount = direct.ValueOf(in.IntentCount)
	out.DirectIntentCount = direct.ValueOf(in.DirectIntentCount)
	out.ParameterFillingCount = direct.ValueOf(in.ParameterFillingCount)
	out.NoMatchCount = direct.ValueOf(in.NoMatchCount)
	out.NoInputCount = direct.ValueOf(in.NoInputCount)
	out.EventCount = direct.ValueOf(in.EventCount)
	return out
}
func Conversation_Metrics_QueryInputCount_FromProto(mapCtx *direct.MapContext, in *pb.Conversation_Metrics_QueryInputCount) *krm.Conversation_Metrics_QueryInputCount {
	if in == nil {
		return nil
	}
	out := &krm.Conversation_Metrics_QueryInputCount{}
	out.TextCount = direct.LazyPtr(in.GetTextCount())
	out.IntentCount = direct.LazyPtr(in.GetIntentCount())
	out.AudioCount = direct.LazyPtr(in.GetAudioCount())
	out.EventCount = direct.LazyPtr(in.GetEventCount())
	out.DtmfCount = direct.LazyPtr(in.GetDtmfCount())
	return out
}
func Conversation_Metrics_QueryInputCount_ToProto(mapCtx *direct.MapContext, in *krm.Conversation_Metrics_QueryInputCount) *pb.Conversation_Metrics_QueryInputCount {
	if in == nil {
		return nil
	}
	out := &pb.Conversation_Metrics_QueryInputCount{}
	out.TextCount = direct.ValueOf(in.TextCount)
	out.IntentCount = direct.ValueOf(in.IntentCount)
	out.AudioCount = direct.ValueOf(in.AudioCount)
	out.EventCount = direct.ValueOf(in.EventCount)
	out.DtmfCount = direct.ValueOf(in.DtmfCount)
	return out
}
func DataStoreConnection_FromProto(mapCtx *direct.MapContext, in *pb.DataStoreConnection) *krm.DataStoreConnection {
	if in == nil {
		return nil
	}
	out := &krm.DataStoreConnection{}
	out.DataStoreType = direct.Enum_FromProto(mapCtx, in.GetDataStoreType())
	out.DataStore = direct.LazyPtr(in.GetDataStore())
	return out
}
func DataStoreConnection_ToProto(mapCtx *direct.MapContext, in *krm.DataStoreConnection) *pb.DataStoreConnection {
	if in == nil {
		return nil
	}
	out := &pb.DataStoreConnection{}
	out.DataStoreType = direct.Enum_ToProto[pb.DataStoreType](mapCtx, in.DataStoreType)
	out.DataStore = direct.ValueOf(in.DataStore)
	return out
}
func DataStoreConnectionSignals_FromProto(mapCtx *direct.MapContext, in *pb.DataStoreConnectionSignals) *krm.DataStoreConnectionSignals {
	if in == nil {
		return nil
	}
	out := &krm.DataStoreConnectionSignals{}
	out.RewriterModelCallSignals = DataStoreConnectionSignals_RewriterModelCallSignals_FromProto(mapCtx, in.GetRewriterModelCallSignals())
	out.RewrittenQuery = direct.LazyPtr(in.GetRewrittenQuery())
	out.SearchSnippets = direct.Slice_FromProto(mapCtx, in.SearchSnippets, DataStoreConnectionSignals_SearchSnippet_FromProto)
	out.AnswerGenerationModelCallSignals = DataStoreConnectionSignals_AnswerGenerationModelCallSignals_FromProto(mapCtx, in.GetAnswerGenerationModelCallSignals())
	out.Answer = direct.LazyPtr(in.GetAnswer())
	out.AnswerParts = direct.Slice_FromProto(mapCtx, in.AnswerParts, DataStoreConnectionSignals_AnswerPart_FromProto)
	out.CitedSnippets = direct.Slice_FromProto(mapCtx, in.CitedSnippets, DataStoreConnectionSignals_CitedSnippet_FromProto)
	out.GroundingSignals = DataStoreConnectionSignals_GroundingSignals_FromProto(mapCtx, in.GetGroundingSignals())
	out.SafetySignals = DataStoreConnectionSignals_SafetySignals_FromProto(mapCtx, in.GetSafetySignals())
	return out
}
func DataStoreConnectionSignals_ToProto(mapCtx *direct.MapContext, in *krm.DataStoreConnectionSignals) *pb.DataStoreConnectionSignals {
	if in == nil {
		return nil
	}
	out := &pb.DataStoreConnectionSignals{}
	out.RewriterModelCallSignals = DataStoreConnectionSignals_RewriterModelCallSignals_ToProto(mapCtx, in.RewriterModelCallSignals)
	out.RewrittenQuery = direct.ValueOf(in.RewrittenQuery)
	out.SearchSnippets = direct.Slice_ToProto(mapCtx, in.SearchSnippets, DataStoreConnectionSignals_SearchSnippet_ToProto)
	out.AnswerGenerationModelCallSignals = DataStoreConnectionSignals_AnswerGenerationModelCallSignals_ToProto(mapCtx, in.AnswerGenerationModelCallSignals)
	out.Answer = direct.ValueOf(in.Answer)
	out.AnswerParts = direct.Slice_ToProto(mapCtx, in.AnswerParts, DataStoreConnectionSignals_AnswerPart_ToProto)
	out.CitedSnippets = direct.Slice_ToProto(mapCtx, in.CitedSnippets, DataStoreConnectionSignals_CitedSnippet_ToProto)
	out.GroundingSignals = DataStoreConnectionSignals_GroundingSignals_ToProto(mapCtx, in.GroundingSignals)
	out.SafetySignals = DataStoreConnectionSignals_SafetySignals_ToProto(mapCtx, in.SafetySignals)
	return out
}
func DataStoreConnectionSignals_AnswerGenerationModelCallSignals_FromProto(mapCtx *direct.MapContext, in *pb.DataStoreConnectionSignals_AnswerGenerationModelCallSignals) *krm.DataStoreConnectionSignals_AnswerGenerationModelCallSignals {
	if in == nil {
		return nil
	}
	out := &krm.DataStoreConnectionSignals_AnswerGenerationModelCallSignals{}
	out.RenderedPrompt = direct.LazyPtr(in.GetRenderedPrompt())
	out.ModelOutput = direct.LazyPtr(in.GetModelOutput())
	out.Model = direct.LazyPtr(in.GetModel())
	return out
}
func DataStoreConnectionSignals_AnswerGenerationModelCallSignals_ToProto(mapCtx *direct.MapContext, in *krm.DataStoreConnectionSignals_AnswerGenerationModelCallSignals) *pb.DataStoreConnectionSignals_AnswerGenerationModelCallSignals {
	if in == nil {
		return nil
	}
	out := &pb.DataStoreConnectionSignals_AnswerGenerationModelCallSignals{}
	out.RenderedPrompt = direct.ValueOf(in.RenderedPrompt)
	out.ModelOutput = direct.ValueOf(in.ModelOutput)
	out.Model = direct.ValueOf(in.Model)
	return out
}
func DataStoreConnectionSignals_AnswerPart_FromProto(mapCtx *direct.MapContext, in *pb.DataStoreConnectionSignals_AnswerPart) *krm.DataStoreConnectionSignals_AnswerPart {
	if in == nil {
		return nil
	}
	out := &krm.DataStoreConnectionSignals_AnswerPart{}
	out.Text = direct.LazyPtr(in.GetText())
	out.SupportingIndices = in.SupportingIndices
	return out
}
func DataStoreConnectionSignals_AnswerPart_ToProto(mapCtx *direct.MapContext, in *krm.DataStoreConnectionSignals_AnswerPart) *pb.DataStoreConnectionSignals_AnswerPart {
	if in == nil {
		return nil
	}
	out := &pb.DataStoreConnectionSignals_AnswerPart{}
	out.Text = direct.ValueOf(in.Text)
	out.SupportingIndices = in.SupportingIndices
	return out
}
func DataStoreConnectionSignals_CitedSnippet_FromProto(mapCtx *direct.MapContext, in *pb.DataStoreConnectionSignals_CitedSnippet) *krm.DataStoreConnectionSignals_CitedSnippet {
	if in == nil {
		return nil
	}
	out := &krm.DataStoreConnectionSignals_CitedSnippet{}
	out.SearchSnippet = DataStoreConnectionSignals_SearchSnippet_FromProto(mapCtx, in.GetSearchSnippet())
	out.SnippetIndex = direct.LazyPtr(in.GetSnippetIndex())
	return out
}
func DataStoreConnectionSignals_CitedSnippet_ToProto(mapCtx *direct.MapContext, in *krm.DataStoreConnectionSignals_CitedSnippet) *pb.DataStoreConnectionSignals_CitedSnippet {
	if in == nil {
		return nil
	}
	out := &pb.DataStoreConnectionSignals_CitedSnippet{}
	out.SearchSnippet = DataStoreConnectionSignals_SearchSnippet_ToProto(mapCtx, in.SearchSnippet)
	out.SnippetIndex = direct.ValueOf(in.SnippetIndex)
	return out
}
func DataStoreConnectionSignals_GroundingSignals_FromProto(mapCtx *direct.MapContext, in *pb.DataStoreConnectionSignals_GroundingSignals) *krm.DataStoreConnectionSignals_GroundingSignals {
	if in == nil {
		return nil
	}
	out := &krm.DataStoreConnectionSignals_GroundingSignals{}
	out.Decision = direct.Enum_FromProto(mapCtx, in.GetDecision())
	out.Score = direct.Enum_FromProto(mapCtx, in.GetScore())
	return out
}
func DataStoreConnectionSignals_GroundingSignals_ToProto(mapCtx *direct.MapContext, in *krm.DataStoreConnectionSignals_GroundingSignals) *pb.DataStoreConnectionSignals_GroundingSignals {
	if in == nil {
		return nil
	}
	out := &pb.DataStoreConnectionSignals_GroundingSignals{}
	out.Decision = direct.Enum_ToProto[pb.DataStoreConnectionSignals_GroundingSignals_GroundingDecision](mapCtx, in.Decision)
	out.Score = direct.Enum_ToProto[pb.DataStoreConnectionSignals_GroundingSignals_GroundingScoreBucket](mapCtx, in.Score)
	return out
}
func DataStoreConnectionSignals_RewriterModelCallSignals_FromProto(mapCtx *direct.MapContext, in *pb.DataStoreConnectionSignals_RewriterModelCallSignals) *krm.DataStoreConnectionSignals_RewriterModelCallSignals {
	if in == nil {
		return nil
	}
	out := &krm.DataStoreConnectionSignals_RewriterModelCallSignals{}
	out.RenderedPrompt = direct.LazyPtr(in.GetRenderedPrompt())
	out.ModelOutput = direct.LazyPtr(in.GetModelOutput())
	out.Model = direct.LazyPtr(in.GetModel())
	return out
}
func DataStoreConnectionSignals_RewriterModelCallSignals_ToProto(mapCtx *direct.MapContext, in *krm.DataStoreConnectionSignals_RewriterModelCallSignals) *pb.DataStoreConnectionSignals_RewriterModelCallSignals {
	if in == nil {
		return nil
	}
	out := &pb.DataStoreConnectionSignals_RewriterModelCallSignals{}
	out.RenderedPrompt = direct.ValueOf(in.RenderedPrompt)
	out.ModelOutput = direct.ValueOf(in.ModelOutput)
	out.Model = direct.ValueOf(in.Model)
	return out
}
func DataStoreConnectionSignals_SafetySignals_FromProto(mapCtx *direct.MapContext, in *pb.DataStoreConnectionSignals_SafetySignals) *krm.DataStoreConnectionSignals_SafetySignals {
	if in == nil {
		return nil
	}
	out := &krm.DataStoreConnectionSignals_SafetySignals{}
	out.Decision = direct.Enum_FromProto(mapCtx, in.GetDecision())
	out.BannedPhraseMatch = direct.Enum_FromProto(mapCtx, in.GetBannedPhraseMatch())
	out.MatchedBannedPhrase = direct.LazyPtr(in.GetMatchedBannedPhrase())
	return out
}
func DataStoreConnectionSignals_SafetySignals_ToProto(mapCtx *direct.MapContext, in *krm.DataStoreConnectionSignals_SafetySignals) *pb.DataStoreConnectionSignals_SafetySignals {
	if in == nil {
		return nil
	}
	out := &pb.DataStoreConnectionSignals_SafetySignals{}
	out.Decision = direct.Enum_ToProto[pb.DataStoreConnectionSignals_SafetySignals_SafetyDecision](mapCtx, in.Decision)
	out.BannedPhraseMatch = direct.Enum_ToProto[pb.DataStoreConnectionSignals_SafetySignals_BannedPhraseMatch](mapCtx, in.BannedPhraseMatch)
	out.MatchedBannedPhrase = direct.ValueOf(in.MatchedBannedPhrase)
	return out
}
func DataStoreConnectionSignals_SearchSnippet_FromProto(mapCtx *direct.MapContext, in *pb.DataStoreConnectionSignals_SearchSnippet) *krm.DataStoreConnectionSignals_SearchSnippet {
	if in == nil {
		return nil
	}
	out := &krm.DataStoreConnectionSignals_SearchSnippet{}
	out.DocumentTitle = direct.LazyPtr(in.GetDocumentTitle())
	out.DocumentURI = direct.LazyPtr(in.GetDocumentUri())
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func DataStoreConnectionSignals_SearchSnippet_ToProto(mapCtx *direct.MapContext, in *krm.DataStoreConnectionSignals_SearchSnippet) *pb.DataStoreConnectionSignals_SearchSnippet {
	if in == nil {
		return nil
	}
	out := &pb.DataStoreConnectionSignals_SearchSnippet{}
	out.DocumentTitle = direct.ValueOf(in.DocumentTitle)
	out.DocumentUri = direct.ValueOf(in.DocumentURI)
	out.Text = direct.ValueOf(in.Text)
	return out
}
func DialogflowConversationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Conversation) *krm.DialogflowConversationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationObservedState{}
	// MISSING: Name
	// MISSING: Type
	// MISSING: LanguageCode
	// MISSING: StartTime
	// MISSING: Duration
	// MISSING: Metrics
	// MISSING: Intents
	// MISSING: Flows
	// MISSING: Pages
	// MISSING: Interactions
	// MISSING: Environment
	// MISSING: FlowVersions
	return out
}
func DialogflowConversationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationObservedState) *pb.Conversation {
	if in == nil {
		return nil
	}
	out := &pb.Conversation{}
	// MISSING: Name
	// MISSING: Type
	// MISSING: LanguageCode
	// MISSING: StartTime
	// MISSING: Duration
	// MISSING: Metrics
	// MISSING: Intents
	// MISSING: Flows
	// MISSING: Pages
	// MISSING: Interactions
	// MISSING: Environment
	// MISSING: FlowVersions
	return out
}
func DialogflowConversationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Conversation) *krm.DialogflowConversationSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationSpec{}
	// MISSING: Name
	// MISSING: Type
	// MISSING: LanguageCode
	// MISSING: StartTime
	// MISSING: Duration
	// MISSING: Metrics
	// MISSING: Intents
	// MISSING: Flows
	// MISSING: Pages
	// MISSING: Interactions
	// MISSING: Environment
	// MISSING: FlowVersions
	return out
}
func DialogflowConversationSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationSpec) *pb.Conversation {
	if in == nil {
		return nil
	}
	out := &pb.Conversation{}
	// MISSING: Name
	// MISSING: Type
	// MISSING: LanguageCode
	// MISSING: StartTime
	// MISSING: Duration
	// MISSING: Metrics
	// MISSING: Intents
	// MISSING: Flows
	// MISSING: Pages
	// MISSING: Interactions
	// MISSING: Environment
	// MISSING: FlowVersions
	return out
}
func DtmfInput_FromProto(mapCtx *direct.MapContext, in *pb.DtmfInput) *krm.DtmfInput {
	if in == nil {
		return nil
	}
	out := &krm.DtmfInput{}
	out.Digits = direct.LazyPtr(in.GetDigits())
	out.FinishDigit = direct.LazyPtr(in.GetFinishDigit())
	return out
}
func DtmfInput_ToProto(mapCtx *direct.MapContext, in *krm.DtmfInput) *pb.DtmfInput {
	if in == nil {
		return nil
	}
	out := &pb.DtmfInput{}
	out.Digits = direct.ValueOf(in.Digits)
	out.FinishDigit = direct.ValueOf(in.FinishDigit)
	return out
}
func Environment_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.Environment {
	if in == nil {
		return nil
	}
	out := &krm.Environment{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.VersionConfigs = direct.Slice_FromProto(mapCtx, in.VersionConfigs, Environment_VersionConfig_FromProto)
	// MISSING: UpdateTime
	out.TestCasesConfig = Environment_TestCasesConfig_FromProto(mapCtx, in.GetTestCasesConfig())
	out.WebhookConfig = Environment_WebhookConfig_FromProto(mapCtx, in.GetWebhookConfig())
	return out
}
func Environment_ToProto(mapCtx *direct.MapContext, in *krm.Environment) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.VersionConfigs = direct.Slice_ToProto(mapCtx, in.VersionConfigs, Environment_VersionConfig_ToProto)
	// MISSING: UpdateTime
	out.TestCasesConfig = Environment_TestCasesConfig_ToProto(mapCtx, in.TestCasesConfig)
	out.WebhookConfig = Environment_WebhookConfig_ToProto(mapCtx, in.WebhookConfig)
	return out
}
func EnvironmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.EnvironmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EnvironmentObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: VersionConfigs
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: TestCasesConfig
	// MISSING: WebhookConfig
	return out
}
func EnvironmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EnvironmentObservedState) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: VersionConfigs
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: TestCasesConfig
	// MISSING: WebhookConfig
	return out
}
func Environment_TestCasesConfig_FromProto(mapCtx *direct.MapContext, in *pb.Environment_TestCasesConfig) *krm.Environment_TestCasesConfig {
	if in == nil {
		return nil
	}
	out := &krm.Environment_TestCasesConfig{}
	out.TestCases = in.TestCases
	out.EnableContinuousRun = direct.LazyPtr(in.GetEnableContinuousRun())
	out.EnablePredeploymentRun = direct.LazyPtr(in.GetEnablePredeploymentRun())
	return out
}
func Environment_TestCasesConfig_ToProto(mapCtx *direct.MapContext, in *krm.Environment_TestCasesConfig) *pb.Environment_TestCasesConfig {
	if in == nil {
		return nil
	}
	out := &pb.Environment_TestCasesConfig{}
	out.TestCases = in.TestCases
	out.EnableContinuousRun = direct.ValueOf(in.EnableContinuousRun)
	out.EnablePredeploymentRun = direct.ValueOf(in.EnablePredeploymentRun)
	return out
}
func Environment_VersionConfig_FromProto(mapCtx *direct.MapContext, in *pb.Environment_VersionConfig) *krm.Environment_VersionConfig {
	if in == nil {
		return nil
	}
	out := &krm.Environment_VersionConfig{}
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func Environment_VersionConfig_ToProto(mapCtx *direct.MapContext, in *krm.Environment_VersionConfig) *pb.Environment_VersionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Environment_VersionConfig{}
	out.Version = direct.ValueOf(in.Version)
	return out
}
func Environment_WebhookConfig_FromProto(mapCtx *direct.MapContext, in *pb.Environment_WebhookConfig) *krm.Environment_WebhookConfig {
	if in == nil {
		return nil
	}
	out := &krm.Environment_WebhookConfig{}
	out.WebhookOverrides = direct.Slice_FromProto(mapCtx, in.WebhookOverrides, Webhook_FromProto)
	return out
}
func Environment_WebhookConfig_ToProto(mapCtx *direct.MapContext, in *krm.Environment_WebhookConfig) *pb.Environment_WebhookConfig {
	if in == nil {
		return nil
	}
	out := &pb.Environment_WebhookConfig{}
	out.WebhookOverrides = direct.Slice_ToProto(mapCtx, in.WebhookOverrides, Webhook_ToProto)
	return out
}
func EventHandler_FromProto(mapCtx *direct.MapContext, in *pb.EventHandler) *krm.EventHandler {
	if in == nil {
		return nil
	}
	out := &krm.EventHandler{}
	// MISSING: Name
	out.Event = direct.LazyPtr(in.GetEvent())
	out.TriggerFulfillment = Fulfillment_FromProto(mapCtx, in.GetTriggerFulfillment())
	out.TargetPage = direct.LazyPtr(in.GetTargetPage())
	out.TargetFlow = direct.LazyPtr(in.GetTargetFlow())
	out.TargetPlaybook = direct.LazyPtr(in.GetTargetPlaybook())
	return out
}
func EventHandler_ToProto(mapCtx *direct.MapContext, in *krm.EventHandler) *pb.EventHandler {
	if in == nil {
		return nil
	}
	out := &pb.EventHandler{}
	// MISSING: Name
	out.Event = direct.ValueOf(in.Event)
	out.TriggerFulfillment = Fulfillment_ToProto(mapCtx, in.TriggerFulfillment)
	if oneof := EventHandler_TargetPage_ToProto(mapCtx, in.TargetPage); oneof != nil {
		out.Target = oneof
	}
	if oneof := EventHandler_TargetFlow_ToProto(mapCtx, in.TargetFlow); oneof != nil {
		out.Target = oneof
	}
	if oneof := EventHandler_TargetPlaybook_ToProto(mapCtx, in.TargetPlaybook); oneof != nil {
		out.Target = oneof
	}
	return out
}
func EventHandlerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EventHandler) *krm.EventHandlerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventHandlerObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Event
	// MISSING: TriggerFulfillment
	// MISSING: TargetPage
	// MISSING: TargetFlow
	// MISSING: TargetPlaybook
	return out
}
func EventHandlerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventHandlerObservedState) *pb.EventHandler {
	if in == nil {
		return nil
	}
	out := &pb.EventHandler{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Event
	// MISSING: TriggerFulfillment
	// MISSING: TargetPage
	// MISSING: TargetFlow
	// MISSING: TargetPlaybook
	return out
}
func EventInput_FromProto(mapCtx *direct.MapContext, in *pb.EventInput) *krm.EventInput {
	if in == nil {
		return nil
	}
	out := &krm.EventInput{}
	out.Event = direct.LazyPtr(in.GetEvent())
	return out
}
func EventInput_ToProto(mapCtx *direct.MapContext, in *krm.EventInput) *pb.EventInput {
	if in == nil {
		return nil
	}
	out := &pb.EventInput{}
	out.Event = direct.ValueOf(in.Event)
	return out
}
func Example_FromProto(mapCtx *direct.MapContext, in *pb.Example) *krm.Example {
	if in == nil {
		return nil
	}
	out := &krm.Example{}
	out.Name = direct.LazyPtr(in.GetName())
	out.PlaybookInput = PlaybookInput_FromProto(mapCtx, in.GetPlaybookInput())
	out.PlaybookOutput = PlaybookOutput_FromProto(mapCtx, in.GetPlaybookOutput())
	out.Actions = direct.Slice_FromProto(mapCtx, in.Actions, Action_FromProto)
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ConversationState = direct.Enum_FromProto(mapCtx, in.GetConversationState())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	return out
}
func Example_ToProto(mapCtx *direct.MapContext, in *krm.Example) *pb.Example {
	if in == nil {
		return nil
	}
	out := &pb.Example{}
	out.Name = direct.ValueOf(in.Name)
	out.PlaybookInput = PlaybookInput_ToProto(mapCtx, in.PlaybookInput)
	out.PlaybookOutput = PlaybookOutput_ToProto(mapCtx, in.PlaybookOutput)
	out.Actions = direct.Slice_ToProto(mapCtx, in.Actions, Action_ToProto)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: TokenCount
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ConversationState = direct.Enum_ToProto[pb.OutputState](mapCtx, in.ConversationState)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	return out
}
func ExampleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Example) *krm.ExampleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ExampleObservedState{}
	// MISSING: Name
	// MISSING: PlaybookInput
	// MISSING: PlaybookOutput
	// MISSING: Actions
	// MISSING: DisplayName
	// MISSING: Description
	out.TokenCount = direct.LazyPtr(in.GetTokenCount())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: ConversationState
	// MISSING: LanguageCode
	return out
}
func ExampleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ExampleObservedState) *pb.Example {
	if in == nil {
		return nil
	}
	out := &pb.Example{}
	// MISSING: Name
	// MISSING: PlaybookInput
	// MISSING: PlaybookOutput
	// MISSING: Actions
	// MISSING: DisplayName
	// MISSING: Description
	out.TokenCount = direct.ValueOf(in.TokenCount)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: ConversationState
	// MISSING: LanguageCode
	return out
}
func FilterSpecs_FromProto(mapCtx *direct.MapContext, in *pb.FilterSpecs) *krm.FilterSpecs {
	if in == nil {
		return nil
	}
	out := &krm.FilterSpecs{}
	out.DataStores = in.DataStores
	out.Filter = direct.LazyPtr(in.GetFilter())
	return out
}
func FilterSpecs_ToProto(mapCtx *direct.MapContext, in *krm.FilterSpecs) *pb.FilterSpecs {
	if in == nil {
		return nil
	}
	out := &pb.FilterSpecs{}
	out.DataStores = in.DataStores
	out.Filter = direct.ValueOf(in.Filter)
	return out
}
func Flow_FromProto(mapCtx *direct.MapContext, in *pb.Flow) *krm.Flow {
	if in == nil {
		return nil
	}
	out := &krm.Flow{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.TransitionRoutes = direct.Slice_FromProto(mapCtx, in.TransitionRoutes, TransitionRoute_FromProto)
	out.EventHandlers = direct.Slice_FromProto(mapCtx, in.EventHandlers, EventHandler_FromProto)
	out.TransitionRouteGroups = in.TransitionRouteGroups
	out.NluSettings = NluSettings_FromProto(mapCtx, in.GetNluSettings())
	out.AdvancedSettings = AdvancedSettings_FromProto(mapCtx, in.GetAdvancedSettings())
	out.KnowledgeConnectorSettings = KnowledgeConnectorSettings_FromProto(mapCtx, in.GetKnowledgeConnectorSettings())
	out.MultiLanguageSettings = Flow_MultiLanguageSettings_FromProto(mapCtx, in.GetMultiLanguageSettings())
	out.Locked = direct.LazyPtr(in.GetLocked())
	return out
}
func Flow_ToProto(mapCtx *direct.MapContext, in *krm.Flow) *pb.Flow {
	if in == nil {
		return nil
	}
	out := &pb.Flow{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.TransitionRoutes = direct.Slice_ToProto(mapCtx, in.TransitionRoutes, TransitionRoute_ToProto)
	out.EventHandlers = direct.Slice_ToProto(mapCtx, in.EventHandlers, EventHandler_ToProto)
	out.TransitionRouteGroups = in.TransitionRouteGroups
	out.NluSettings = NluSettings_ToProto(mapCtx, in.NluSettings)
	out.AdvancedSettings = AdvancedSettings_ToProto(mapCtx, in.AdvancedSettings)
	out.KnowledgeConnectorSettings = KnowledgeConnectorSettings_ToProto(mapCtx, in.KnowledgeConnectorSettings)
	out.MultiLanguageSettings = Flow_MultiLanguageSettings_ToProto(mapCtx, in.MultiLanguageSettings)
	out.Locked = direct.ValueOf(in.Locked)
	return out
}
func FlowInvocation_FromProto(mapCtx *direct.MapContext, in *pb.FlowInvocation) *krm.FlowInvocation {
	if in == nil {
		return nil
	}
	out := &krm.FlowInvocation{}
	out.Flow = direct.LazyPtr(in.GetFlow())
	out.InputActionParameters = InputActionParameters_FromProto(mapCtx, in.GetInputActionParameters())
	out.OutputActionParameters = OutputActionParameters_FromProto(mapCtx, in.GetOutputActionParameters())
	out.FlowState = direct.Enum_FromProto(mapCtx, in.GetFlowState())
	return out
}
func FlowInvocation_ToProto(mapCtx *direct.MapContext, in *krm.FlowInvocation) *pb.FlowInvocation {
	if in == nil {
		return nil
	}
	out := &pb.FlowInvocation{}
	out.Flow = direct.ValueOf(in.Flow)
	out.InputActionParameters = InputActionParameters_ToProto(mapCtx, in.InputActionParameters)
	out.OutputActionParameters = OutputActionParameters_ToProto(mapCtx, in.OutputActionParameters)
	out.FlowState = direct.Enum_ToProto[pb.OutputState](mapCtx, in.FlowState)
	return out
}
func FlowObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Flow) *krm.FlowObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FlowObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.TransitionRoutes = direct.Slice_FromProto(mapCtx, in.TransitionRoutes, TransitionRouteObservedState_FromProto)
	out.EventHandlers = direct.Slice_FromProto(mapCtx, in.EventHandlers, EventHandlerObservedState_FromProto)
	// MISSING: TransitionRouteGroups
	// MISSING: NluSettings
	// MISSING: AdvancedSettings
	// MISSING: KnowledgeConnectorSettings
	// MISSING: MultiLanguageSettings
	// MISSING: Locked
	return out
}
func FlowObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FlowObservedState) *pb.Flow {
	if in == nil {
		return nil
	}
	out := &pb.Flow{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.TransitionRoutes = direct.Slice_ToProto(mapCtx, in.TransitionRoutes, TransitionRouteObservedState_ToProto)
	out.EventHandlers = direct.Slice_ToProto(mapCtx, in.EventHandlers, EventHandlerObservedState_ToProto)
	// MISSING: TransitionRouteGroups
	// MISSING: NluSettings
	// MISSING: AdvancedSettings
	// MISSING: KnowledgeConnectorSettings
	// MISSING: MultiLanguageSettings
	// MISSING: Locked
	return out
}
func Flow_MultiLanguageSettings_FromProto(mapCtx *direct.MapContext, in *pb.Flow_MultiLanguageSettings) *krm.Flow_MultiLanguageSettings {
	if in == nil {
		return nil
	}
	out := &krm.Flow_MultiLanguageSettings{}
	out.EnableMultiLanguageDetection = direct.LazyPtr(in.GetEnableMultiLanguageDetection())
	out.SupportedResponseLanguageCodes = in.SupportedResponseLanguageCodes
	return out
}
func Flow_MultiLanguageSettings_ToProto(mapCtx *direct.MapContext, in *krm.Flow_MultiLanguageSettings) *pb.Flow_MultiLanguageSettings {
	if in == nil {
		return nil
	}
	out := &pb.Flow_MultiLanguageSettings{}
	out.EnableMultiLanguageDetection = direct.ValueOf(in.EnableMultiLanguageDetection)
	out.SupportedResponseLanguageCodes = in.SupportedResponseLanguageCodes
	return out
}
func Form_FromProto(mapCtx *direct.MapContext, in *pb.Form) *krm.Form {
	if in == nil {
		return nil
	}
	out := &krm.Form{}
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, Form_Parameter_FromProto)
	return out
}
func Form_ToProto(mapCtx *direct.MapContext, in *krm.Form) *pb.Form {
	if in == nil {
		return nil
	}
	out := &pb.Form{}
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, Form_Parameter_ToProto)
	return out
}
func Form_Parameter_FromProto(mapCtx *direct.MapContext, in *pb.Form_Parameter) *krm.Form_Parameter {
	if in == nil {
		return nil
	}
	out := &krm.Form_Parameter{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Required = direct.LazyPtr(in.GetRequired())
	out.EntityType = direct.LazyPtr(in.GetEntityType())
	out.IsList = direct.LazyPtr(in.GetIsList())
	out.FillBehavior = Form_Parameter_FillBehavior_FromProto(mapCtx, in.GetFillBehavior())
	out.DefaultValue = Value_FromProto(mapCtx, in.GetDefaultValue())
	out.Redact = direct.LazyPtr(in.GetRedact())
	out.AdvancedSettings = AdvancedSettings_FromProto(mapCtx, in.GetAdvancedSettings())
	return out
}
func Form_Parameter_ToProto(mapCtx *direct.MapContext, in *krm.Form_Parameter) *pb.Form_Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Form_Parameter{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Required = direct.ValueOf(in.Required)
	out.EntityType = direct.ValueOf(in.EntityType)
	out.IsList = direct.ValueOf(in.IsList)
	out.FillBehavior = Form_Parameter_FillBehavior_ToProto(mapCtx, in.FillBehavior)
	out.DefaultValue = Value_ToProto(mapCtx, in.DefaultValue)
	out.Redact = direct.ValueOf(in.Redact)
	out.AdvancedSettings = AdvancedSettings_ToProto(mapCtx, in.AdvancedSettings)
	return out
}
func Form_Parameter_FillBehavior_FromProto(mapCtx *direct.MapContext, in *pb.Form_Parameter_FillBehavior) *krm.Form_Parameter_FillBehavior {
	if in == nil {
		return nil
	}
	out := &krm.Form_Parameter_FillBehavior{}
	out.InitialPromptFulfillment = Fulfillment_FromProto(mapCtx, in.GetInitialPromptFulfillment())
	out.RepromptEventHandlers = direct.Slice_FromProto(mapCtx, in.RepromptEventHandlers, EventHandler_FromProto)
	return out
}
func Form_Parameter_FillBehavior_ToProto(mapCtx *direct.MapContext, in *krm.Form_Parameter_FillBehavior) *pb.Form_Parameter_FillBehavior {
	if in == nil {
		return nil
	}
	out := &pb.Form_Parameter_FillBehavior{}
	out.InitialPromptFulfillment = Fulfillment_ToProto(mapCtx, in.InitialPromptFulfillment)
	out.RepromptEventHandlers = direct.Slice_ToProto(mapCtx, in.RepromptEventHandlers, EventHandler_ToProto)
	return out
}
func Fulfillment_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment) *krm.Fulfillment {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment{}
	out.Messages = direct.Slice_FromProto(mapCtx, in.Messages, ResponseMessage_FromProto)
	out.Webhook = direct.LazyPtr(in.GetWebhook())
	out.ReturnPartialResponses = direct.LazyPtr(in.GetReturnPartialResponses())
	out.Tag = direct.LazyPtr(in.GetTag())
	out.SetParameterActions = direct.Slice_FromProto(mapCtx, in.SetParameterActions, Fulfillment_SetParameterAction_FromProto)
	out.ConditionalCases = direct.Slice_FromProto(mapCtx, in.ConditionalCases, Fulfillment_ConditionalCases_FromProto)
	out.AdvancedSettings = AdvancedSettings_FromProto(mapCtx, in.GetAdvancedSettings())
	out.EnableGenerativeFallback = direct.LazyPtr(in.GetEnableGenerativeFallback())
	return out
}
func Fulfillment_ToProto(mapCtx *direct.MapContext, in *krm.Fulfillment) *pb.Fulfillment {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment{}
	out.Messages = direct.Slice_ToProto(mapCtx, in.Messages, ResponseMessage_ToProto)
	out.Webhook = direct.ValueOf(in.Webhook)
	out.ReturnPartialResponses = direct.ValueOf(in.ReturnPartialResponses)
	out.Tag = direct.ValueOf(in.Tag)
	out.SetParameterActions = direct.Slice_ToProto(mapCtx, in.SetParameterActions, Fulfillment_SetParameterAction_ToProto)
	out.ConditionalCases = direct.Slice_ToProto(mapCtx, in.ConditionalCases, Fulfillment_ConditionalCases_ToProto)
	out.AdvancedSettings = AdvancedSettings_ToProto(mapCtx, in.AdvancedSettings)
	out.EnableGenerativeFallback = direct.ValueOf(in.EnableGenerativeFallback)
	return out
}
func FulfillmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment) *krm.FulfillmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FulfillmentObservedState{}
	out.Messages = direct.Slice_FromProto(mapCtx, in.Messages, ResponseMessageObservedState_FromProto)
	// MISSING: Webhook
	// MISSING: ReturnPartialResponses
	// MISSING: Tag
	// MISSING: SetParameterActions
	// MISSING: ConditionalCases
	// MISSING: AdvancedSettings
	// MISSING: EnableGenerativeFallback
	return out
}
func FulfillmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FulfillmentObservedState) *pb.Fulfillment {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment{}
	out.Messages = direct.Slice_ToProto(mapCtx, in.Messages, ResponseMessageObservedState_ToProto)
	// MISSING: Webhook
	// MISSING: ReturnPartialResponses
	// MISSING: Tag
	// MISSING: SetParameterActions
	// MISSING: ConditionalCases
	// MISSING: AdvancedSettings
	// MISSING: EnableGenerativeFallback
	return out
}
func Fulfillment_ConditionalCases_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment_ConditionalCases) *krm.Fulfillment_ConditionalCases {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment_ConditionalCases{}
	out.Cases = direct.Slice_FromProto(mapCtx, in.Cases, Fulfillment_ConditionalCases_Case_FromProto)
	return out
}
func Fulfillment_ConditionalCases_ToProto(mapCtx *direct.MapContext, in *krm.Fulfillment_ConditionalCases) *pb.Fulfillment_ConditionalCases {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment_ConditionalCases{}
	out.Cases = direct.Slice_ToProto(mapCtx, in.Cases, Fulfillment_ConditionalCases_Case_ToProto)
	return out
}
func Fulfillment_ConditionalCases_Case_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment_ConditionalCases_Case) *krm.Fulfillment_ConditionalCases_Case {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment_ConditionalCases_Case{}
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.CaseContent = direct.Slice_FromProto(mapCtx, in.CaseContent, Fulfillment_ConditionalCases_Case_CaseContent_FromProto)
	return out
}
func Fulfillment_ConditionalCases_Case_ToProto(mapCtx *direct.MapContext, in *krm.Fulfillment_ConditionalCases_Case) *pb.Fulfillment_ConditionalCases_Case {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment_ConditionalCases_Case{}
	out.Condition = direct.ValueOf(in.Condition)
	out.CaseContent = direct.Slice_ToProto(mapCtx, in.CaseContent, Fulfillment_ConditionalCases_Case_CaseContent_ToProto)
	return out
}
func Fulfillment_ConditionalCases_Case_CaseContent_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment_ConditionalCases_Case_CaseContent) *krm.Fulfillment_ConditionalCases_Case_CaseContent {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment_ConditionalCases_Case_CaseContent{}
	out.Message = ResponseMessage_FromProto(mapCtx, in.GetMessage())
	out.AdditionalCases = Fulfillment_ConditionalCases_FromProto(mapCtx, in.GetAdditionalCases())
	return out
}
func Fulfillment_ConditionalCases_Case_CaseContent_ToProto(mapCtx *direct.MapContext, in *krm.Fulfillment_ConditionalCases_Case_CaseContent) *pb.Fulfillment_ConditionalCases_Case_CaseContent {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment_ConditionalCases_Case_CaseContent{}
	if oneof := ResponseMessage_ToProto(mapCtx, in.Message); oneof != nil {
		out.CasesOrMessage = &pb.Fulfillment_ConditionalCases_Case_CaseContent_Message{Message: oneof}
	}
	if oneof := Fulfillment_ConditionalCases_ToProto(mapCtx, in.AdditionalCases); oneof != nil {
		out.CasesOrMessage = &pb.Fulfillment_ConditionalCases_Case_CaseContent_AdditionalCases{AdditionalCases: oneof}
	}
	return out
}
func Fulfillment_SetParameterAction_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment_SetParameterAction) *krm.Fulfillment_SetParameterAction {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment_SetParameterAction{}
	out.Parameter = direct.LazyPtr(in.GetParameter())
	out.Value = Value_FromProto(mapCtx, in.GetValue())
	return out
}
func Fulfillment_SetParameterAction_ToProto(mapCtx *direct.MapContext, in *krm.Fulfillment_SetParameterAction) *pb.Fulfillment_SetParameterAction {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment_SetParameterAction{}
	out.Parameter = direct.ValueOf(in.Parameter)
	out.Value = Value_ToProto(mapCtx, in.Value)
	return out
}
func GcsDestination_FromProto(mapCtx *direct.MapContext, in *pb.GcsDestination) *krm.GcsDestination {
	if in == nil {
		return nil
	}
	out := &krm.GcsDestination{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func GcsDestination_ToProto(mapCtx *direct.MapContext, in *krm.GcsDestination) *pb.GcsDestination {
	if in == nil {
		return nil
	}
	out := &pb.GcsDestination{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func GenerativeInfo_FromProto(mapCtx *direct.MapContext, in *pb.GenerativeInfo) *krm.GenerativeInfo {
	if in == nil {
		return nil
	}
	out := &krm.GenerativeInfo{}
	out.CurrentPlaybooks = in.CurrentPlaybooks
	out.ActionTracingInfo = Example_FromProto(mapCtx, in.GetActionTracingInfo())
	return out
}
func GenerativeInfo_ToProto(mapCtx *direct.MapContext, in *krm.GenerativeInfo) *pb.GenerativeInfo {
	if in == nil {
		return nil
	}
	out := &pb.GenerativeInfo{}
	out.CurrentPlaybooks = in.CurrentPlaybooks
	out.ActionTracingInfo = Example_ToProto(mapCtx, in.ActionTracingInfo)
	return out
}
func GenerativeInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GenerativeInfo) *krm.GenerativeInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GenerativeInfoObservedState{}
	// MISSING: CurrentPlaybooks
	out.ActionTracingInfo = ExampleObservedState_FromProto(mapCtx, in.GetActionTracingInfo())
	return out
}
func GenerativeInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GenerativeInfoObservedState) *pb.GenerativeInfo {
	if in == nil {
		return nil
	}
	out := &pb.GenerativeInfo{}
	// MISSING: CurrentPlaybooks
	out.ActionTracingInfo = ExampleObservedState_ToProto(mapCtx, in.ActionTracingInfo)
	return out
}
func InputAudioConfig_FromProto(mapCtx *direct.MapContext, in *pb.InputAudioConfig) *krm.InputAudioConfig {
	if in == nil {
		return nil
	}
	out := &krm.InputAudioConfig{}
	out.AudioEncoding = direct.Enum_FromProto(mapCtx, in.GetAudioEncoding())
	out.SampleRateHertz = direct.LazyPtr(in.GetSampleRateHertz())
	out.EnableWordInfo = direct.LazyPtr(in.GetEnableWordInfo())
	out.PhraseHints = in.PhraseHints
	out.Model = direct.LazyPtr(in.GetModel())
	out.ModelVariant = direct.Enum_FromProto(mapCtx, in.GetModelVariant())
	out.SingleUtterance = direct.LazyPtr(in.GetSingleUtterance())
	out.BargeInConfig = BargeInConfig_FromProto(mapCtx, in.GetBargeInConfig())
	out.OptOutConformerModelMigration = direct.LazyPtr(in.GetOptOutConformerModelMigration())
	return out
}
func InputAudioConfig_ToProto(mapCtx *direct.MapContext, in *krm.InputAudioConfig) *pb.InputAudioConfig {
	if in == nil {
		return nil
	}
	out := &pb.InputAudioConfig{}
	out.AudioEncoding = direct.Enum_ToProto[pb.AudioEncoding](mapCtx, in.AudioEncoding)
	out.SampleRateHertz = direct.ValueOf(in.SampleRateHertz)
	out.EnableWordInfo = direct.ValueOf(in.EnableWordInfo)
	out.PhraseHints = in.PhraseHints
	out.Model = direct.ValueOf(in.Model)
	out.ModelVariant = direct.Enum_ToProto[pb.SpeechModelVariant](mapCtx, in.ModelVariant)
	out.SingleUtterance = direct.ValueOf(in.SingleUtterance)
	out.BargeInConfig = BargeInConfig_ToProto(mapCtx, in.BargeInConfig)
	out.OptOutConformerModelMigration = direct.ValueOf(in.OptOutConformerModelMigration)
	return out
}
func Intent_FromProto(mapCtx *direct.MapContext, in *pb.Intent) *krm.Intent {
	if in == nil {
		return nil
	}
	out := &krm.Intent{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.TrainingPhrases = direct.Slice_FromProto(mapCtx, in.TrainingPhrases, Intent_TrainingPhrase_FromProto)
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, Intent_Parameter_FromProto)
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.IsFallback = direct.LazyPtr(in.GetIsFallback())
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func Intent_ToProto(mapCtx *direct.MapContext, in *krm.Intent) *pb.Intent {
	if in == nil {
		return nil
	}
	out := &pb.Intent{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.TrainingPhrases = direct.Slice_ToProto(mapCtx, in.TrainingPhrases, Intent_TrainingPhrase_ToProto)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, Intent_Parameter_ToProto)
	out.Priority = direct.ValueOf(in.Priority)
	out.IsFallback = direct.ValueOf(in.IsFallback)
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	return out
}
func IntentInput_FromProto(mapCtx *direct.MapContext, in *pb.IntentInput) *krm.IntentInput {
	if in == nil {
		return nil
	}
	out := &krm.IntentInput{}
	out.Intent = direct.LazyPtr(in.GetIntent())
	return out
}
func IntentInput_ToProto(mapCtx *direct.MapContext, in *krm.IntentInput) *pb.IntentInput {
	if in == nil {
		return nil
	}
	out := &pb.IntentInput{}
	out.Intent = direct.ValueOf(in.Intent)
	return out
}
func Intent_Parameter_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Parameter) *krm.Intent_Parameter {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Parameter{}
	out.ID = direct.LazyPtr(in.GetId())
	out.EntityType = direct.LazyPtr(in.GetEntityType())
	out.IsList = direct.LazyPtr(in.GetIsList())
	out.Redact = direct.LazyPtr(in.GetRedact())
	return out
}
func Intent_Parameter_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Parameter) *pb.Intent_Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Parameter{}
	out.Id = direct.ValueOf(in.ID)
	out.EntityType = direct.ValueOf(in.EntityType)
	out.IsList = direct.ValueOf(in.IsList)
	out.Redact = direct.ValueOf(in.Redact)
	return out
}
func Intent_TrainingPhrase_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase) *krm.Intent_TrainingPhrase {
	if in == nil {
		return nil
	}
	out := &krm.Intent_TrainingPhrase{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Parts = direct.Slice_FromProto(mapCtx, in.Parts, Intent_TrainingPhrase_Part_FromProto)
	out.RepeatCount = direct.LazyPtr(in.GetRepeatCount())
	return out
}
func Intent_TrainingPhrase_ToProto(mapCtx *direct.MapContext, in *krm.Intent_TrainingPhrase) *pb.Intent_TrainingPhrase {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase{}
	out.Id = direct.ValueOf(in.ID)
	out.Parts = direct.Slice_ToProto(mapCtx, in.Parts, Intent_TrainingPhrase_Part_ToProto)
	out.RepeatCount = direct.ValueOf(in.RepeatCount)
	return out
}
func Intent_TrainingPhrase_Part_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase_Part) *krm.Intent_TrainingPhrase_Part {
	if in == nil {
		return nil
	}
	out := &krm.Intent_TrainingPhrase_Part{}
	out.Text = direct.LazyPtr(in.GetText())
	out.ParameterID = direct.LazyPtr(in.GetParameterId())
	return out
}
func Intent_TrainingPhrase_Part_ToProto(mapCtx *direct.MapContext, in *krm.Intent_TrainingPhrase_Part) *pb.Intent_TrainingPhrase_Part {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase_Part{}
	out.Text = direct.ValueOf(in.Text)
	out.ParameterId = direct.ValueOf(in.ParameterID)
	return out
}
func KnowledgeConnectorSettings_FromProto(mapCtx *direct.MapContext, in *pb.KnowledgeConnectorSettings) *krm.KnowledgeConnectorSettings {
	if in == nil {
		return nil
	}
	out := &krm.KnowledgeConnectorSettings{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.TriggerFulfillment = Fulfillment_FromProto(mapCtx, in.GetTriggerFulfillment())
	out.TargetPage = direct.LazyPtr(in.GetTargetPage())
	out.TargetFlow = direct.LazyPtr(in.GetTargetFlow())
	out.DataStoreConnections = direct.Slice_FromProto(mapCtx, in.DataStoreConnections, DataStoreConnection_FromProto)
	return out
}
func KnowledgeConnectorSettings_ToProto(mapCtx *direct.MapContext, in *krm.KnowledgeConnectorSettings) *pb.KnowledgeConnectorSettings {
	if in == nil {
		return nil
	}
	out := &pb.KnowledgeConnectorSettings{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.TriggerFulfillment = Fulfillment_ToProto(mapCtx, in.TriggerFulfillment)
	if oneof := KnowledgeConnectorSettings_TargetPage_ToProto(mapCtx, in.TargetPage); oneof != nil {
		out.Target = oneof
	}
	if oneof := KnowledgeConnectorSettings_TargetFlow_ToProto(mapCtx, in.TargetFlow); oneof != nil {
		out.Target = oneof
	}
	out.DataStoreConnections = direct.Slice_ToProto(mapCtx, in.DataStoreConnections, DataStoreConnection_ToProto)
	return out
}
func LlmModelSettings_FromProto(mapCtx *direct.MapContext, in *pb.LlmModelSettings) *krm.LlmModelSettings {
	if in == nil {
		return nil
	}
	out := &krm.LlmModelSettings{}
	out.Model = direct.LazyPtr(in.GetModel())
	out.PromptText = direct.LazyPtr(in.GetPromptText())
	return out
}
func LlmModelSettings_ToProto(mapCtx *direct.MapContext, in *krm.LlmModelSettings) *pb.LlmModelSettings {
	if in == nil {
		return nil
	}
	out := &pb.LlmModelSettings{}
	out.Model = direct.ValueOf(in.Model)
	out.PromptText = direct.ValueOf(in.PromptText)
	return out
}
func Match_FromProto(mapCtx *direct.MapContext, in *pb.Match) *krm.Match {
	if in == nil {
		return nil
	}
	out := &krm.Match{}
	out.Intent = Intent_FromProto(mapCtx, in.GetIntent())
	out.Event = direct.LazyPtr(in.GetEvent())
	out.Parameters = Parameters_FromProto(mapCtx, in.GetParameters())
	out.ResolvedInput = direct.LazyPtr(in.GetResolvedInput())
	out.MatchType = direct.Enum_FromProto(mapCtx, in.GetMatchType())
	out.Confidence = direct.LazyPtr(in.GetConfidence())
	return out
}
func Match_ToProto(mapCtx *direct.MapContext, in *krm.Match) *pb.Match {
	if in == nil {
		return nil
	}
	out := &pb.Match{}
	out.Intent = Intent_ToProto(mapCtx, in.Intent)
	out.Event = direct.ValueOf(in.Event)
	out.Parameters = Parameters_ToProto(mapCtx, in.Parameters)
	out.ResolvedInput = direct.ValueOf(in.ResolvedInput)
	out.MatchType = direct.Enum_ToProto[pb.Match_MatchType](mapCtx, in.MatchType)
	out.Confidence = direct.ValueOf(in.Confidence)
	return out
}
func NluSettings_FromProto(mapCtx *direct.MapContext, in *pb.NluSettings) *krm.NluSettings {
	if in == nil {
		return nil
	}
	out := &krm.NluSettings{}
	out.ModelType = direct.Enum_FromProto(mapCtx, in.GetModelType())
	out.ClassificationThreshold = direct.LazyPtr(in.GetClassificationThreshold())
	out.ModelTrainingMode = direct.Enum_FromProto(mapCtx, in.GetModelTrainingMode())
	return out
}
func NluSettings_ToProto(mapCtx *direct.MapContext, in *krm.NluSettings) *pb.NluSettings {
	if in == nil {
		return nil
	}
	out := &pb.NluSettings{}
	out.ModelType = direct.Enum_ToProto[pb.NluSettings_ModelType](mapCtx, in.ModelType)
	out.ClassificationThreshold = direct.ValueOf(in.ClassificationThreshold)
	out.ModelTrainingMode = direct.Enum_ToProto[pb.NluSettings_ModelTrainingMode](mapCtx, in.ModelTrainingMode)
	return out
}
func OutputAudioConfig_FromProto(mapCtx *direct.MapContext, in *pb.OutputAudioConfig) *krm.OutputAudioConfig {
	if in == nil {
		return nil
	}
	out := &krm.OutputAudioConfig{}
	out.AudioEncoding = direct.Enum_FromProto(mapCtx, in.GetAudioEncoding())
	out.SampleRateHertz = direct.LazyPtr(in.GetSampleRateHertz())
	out.SynthesizeSpeechConfig = SynthesizeSpeechConfig_FromProto(mapCtx, in.GetSynthesizeSpeechConfig())
	return out
}
func OutputAudioConfig_ToProto(mapCtx *direct.MapContext, in *krm.OutputAudioConfig) *pb.OutputAudioConfig {
	if in == nil {
		return nil
	}
	out := &pb.OutputAudioConfig{}
	out.AudioEncoding = direct.Enum_ToProto[pb.OutputAudioEncoding](mapCtx, in.AudioEncoding)
	out.SampleRateHertz = direct.ValueOf(in.SampleRateHertz)
	out.SynthesizeSpeechConfig = SynthesizeSpeechConfig_ToProto(mapCtx, in.SynthesizeSpeechConfig)
	return out
}
func Page_FromProto(mapCtx *direct.MapContext, in *pb.Page) *krm.Page {
	if in == nil {
		return nil
	}
	out := &krm.Page{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.EntryFulfillment = Fulfillment_FromProto(mapCtx, in.GetEntryFulfillment())
	out.Form = Form_FromProto(mapCtx, in.GetForm())
	out.TransitionRouteGroups = in.TransitionRouteGroups
	out.TransitionRoutes = direct.Slice_FromProto(mapCtx, in.TransitionRoutes, TransitionRoute_FromProto)
	out.EventHandlers = direct.Slice_FromProto(mapCtx, in.EventHandlers, EventHandler_FromProto)
	out.AdvancedSettings = AdvancedSettings_FromProto(mapCtx, in.GetAdvancedSettings())
	out.KnowledgeConnectorSettings = KnowledgeConnectorSettings_FromProto(mapCtx, in.GetKnowledgeConnectorSettings())
	return out
}
func Page_ToProto(mapCtx *direct.MapContext, in *krm.Page) *pb.Page {
	if in == nil {
		return nil
	}
	out := &pb.Page{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.EntryFulfillment = Fulfillment_ToProto(mapCtx, in.EntryFulfillment)
	out.Form = Form_ToProto(mapCtx, in.Form)
	out.TransitionRouteGroups = in.TransitionRouteGroups
	out.TransitionRoutes = direct.Slice_ToProto(mapCtx, in.TransitionRoutes, TransitionRoute_ToProto)
	out.EventHandlers = direct.Slice_ToProto(mapCtx, in.EventHandlers, EventHandler_ToProto)
	out.AdvancedSettings = AdvancedSettings_ToProto(mapCtx, in.AdvancedSettings)
	out.KnowledgeConnectorSettings = KnowledgeConnectorSettings_ToProto(mapCtx, in.KnowledgeConnectorSettings)
	return out
}
func PlaybookInput_FromProto(mapCtx *direct.MapContext, in *pb.PlaybookInput) *krm.PlaybookInput {
	if in == nil {
		return nil
	}
	out := &krm.PlaybookInput{}
	out.PrecedingConversationSummary = direct.LazyPtr(in.GetPrecedingConversationSummary())
	out.ActionParameters = ActionParameters_FromProto(mapCtx, in.GetActionParameters())
	return out
}
func PlaybookInput_ToProto(mapCtx *direct.MapContext, in *krm.PlaybookInput) *pb.PlaybookInput {
	if in == nil {
		return nil
	}
	out := &pb.PlaybookInput{}
	out.PrecedingConversationSummary = direct.ValueOf(in.PrecedingConversationSummary)
	out.ActionParameters = ActionParameters_ToProto(mapCtx, in.ActionParameters)
	return out
}
func PlaybookInvocation_FromProto(mapCtx *direct.MapContext, in *pb.PlaybookInvocation) *krm.PlaybookInvocation {
	if in == nil {
		return nil
	}
	out := &krm.PlaybookInvocation{}
	out.Playbook = direct.LazyPtr(in.GetPlaybook())
	out.PlaybookInput = PlaybookInput_FromProto(mapCtx, in.GetPlaybookInput())
	out.PlaybookOutput = PlaybookOutput_FromProto(mapCtx, in.GetPlaybookOutput())
	out.PlaybookState = direct.Enum_FromProto(mapCtx, in.GetPlaybookState())
	return out
}
func PlaybookInvocation_ToProto(mapCtx *direct.MapContext, in *krm.PlaybookInvocation) *pb.PlaybookInvocation {
	if in == nil {
		return nil
	}
	out := &pb.PlaybookInvocation{}
	out.Playbook = direct.ValueOf(in.Playbook)
	out.PlaybookInput = PlaybookInput_ToProto(mapCtx, in.PlaybookInput)
	out.PlaybookOutput = PlaybookOutput_ToProto(mapCtx, in.PlaybookOutput)
	out.PlaybookState = direct.Enum_ToProto[pb.OutputState](mapCtx, in.PlaybookState)
	return out
}
func PlaybookOutput_FromProto(mapCtx *direct.MapContext, in *pb.PlaybookOutput) *krm.PlaybookOutput {
	if in == nil {
		return nil
	}
	out := &krm.PlaybookOutput{}
	out.ExecutionSummary = direct.LazyPtr(in.GetExecutionSummary())
	out.ActionParameters = ActionParameters_FromProto(mapCtx, in.GetActionParameters())
	return out
}
func PlaybookOutput_ToProto(mapCtx *direct.MapContext, in *krm.PlaybookOutput) *pb.PlaybookOutput {
	if in == nil {
		return nil
	}
	out := &pb.PlaybookOutput{}
	out.ExecutionSummary = direct.ValueOf(in.ExecutionSummary)
	out.ActionParameters = ActionParameters_ToProto(mapCtx, in.ActionParameters)
	return out
}
func QueryInput_FromProto(mapCtx *direct.MapContext, in *pb.QueryInput) *krm.QueryInput {
	if in == nil {
		return nil
	}
	out := &krm.QueryInput{}
	out.Text = TextInput_FromProto(mapCtx, in.GetText())
	out.Intent = IntentInput_FromProto(mapCtx, in.GetIntent())
	out.Audio = AudioInput_FromProto(mapCtx, in.GetAudio())
	out.Event = EventInput_FromProto(mapCtx, in.GetEvent())
	out.Dtmf = DtmfInput_FromProto(mapCtx, in.GetDtmf())
	out.ToolCallResult = ToolCallResult_FromProto(mapCtx, in.GetToolCallResult())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	return out
}
func QueryInput_ToProto(mapCtx *direct.MapContext, in *krm.QueryInput) *pb.QueryInput {
	if in == nil {
		return nil
	}
	out := &pb.QueryInput{}
	if oneof := TextInput_ToProto(mapCtx, in.Text); oneof != nil {
		out.Input = &pb.QueryInput_Text{Text: oneof}
	}
	if oneof := IntentInput_ToProto(mapCtx, in.Intent); oneof != nil {
		out.Input = &pb.QueryInput_Intent{Intent: oneof}
	}
	if oneof := AudioInput_ToProto(mapCtx, in.Audio); oneof != nil {
		out.Input = &pb.QueryInput_Audio{Audio: oneof}
	}
	if oneof := EventInput_ToProto(mapCtx, in.Event); oneof != nil {
		out.Input = &pb.QueryInput_Event{Event: oneof}
	}
	if oneof := DtmfInput_ToProto(mapCtx, in.Dtmf); oneof != nil {
		out.Input = &pb.QueryInput_Dtmf{Dtmf: oneof}
	}
	if oneof := ToolCallResult_ToProto(mapCtx, in.ToolCallResult); oneof != nil {
		out.Input = &pb.QueryInput_ToolCallResult{ToolCallResult: oneof}
	}
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	return out
}
func QueryParameters_FromProto(mapCtx *direct.MapContext, in *pb.QueryParameters) *krm.QueryParameters {
	if in == nil {
		return nil
	}
	out := &krm.QueryParameters{}
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.GeoLocation = LatLng_FromProto(mapCtx, in.GetGeoLocation())
	out.SessionEntityTypes = direct.Slice_FromProto(mapCtx, in.SessionEntityTypes, SessionEntityType_FromProto)
	out.Payload = Payload_FromProto(mapCtx, in.GetPayload())
	out.Parameters = Parameters_FromProto(mapCtx, in.GetParameters())
	out.CurrentPage = direct.LazyPtr(in.GetCurrentPage())
	out.DisableWebhook = direct.LazyPtr(in.GetDisableWebhook())
	out.AnalyzeQueryTextSentiment = direct.LazyPtr(in.GetAnalyzeQueryTextSentiment())
	out.WebhookHeaders = in.WebhookHeaders
	out.FlowVersions = in.FlowVersions
	out.CurrentPlaybook = direct.LazyPtr(in.GetCurrentPlaybook())
	out.LlmModelSettings = LlmModelSettings_FromProto(mapCtx, in.GetLlmModelSettings())
	out.Channel = direct.LazyPtr(in.GetChannel())
	out.SessionTtl = direct.StringDuration_FromProto(mapCtx, in.GetSessionTtl())
	out.EndUserMetadata = EndUserMetadata_FromProto(mapCtx, in.GetEndUserMetadata())
	out.SearchConfig = SearchConfig_FromProto(mapCtx, in.GetSearchConfig())
	out.PopulateDataStoreConnectionSignals = direct.LazyPtr(in.GetPopulateDataStoreConnectionSignals())
	return out
}
func QueryParameters_ToProto(mapCtx *direct.MapContext, in *krm.QueryParameters) *pb.QueryParameters {
	if in == nil {
		return nil
	}
	out := &pb.QueryParameters{}
	out.TimeZone = direct.ValueOf(in.TimeZone)
	out.GeoLocation = LatLng_ToProto(mapCtx, in.GeoLocation)
	out.SessionEntityTypes = direct.Slice_ToProto(mapCtx, in.SessionEntityTypes, SessionEntityType_ToProto)
	out.Payload = Payload_ToProto(mapCtx, in.Payload)
	out.Parameters = Parameters_ToProto(mapCtx, in.Parameters)
	out.CurrentPage = direct.ValueOf(in.CurrentPage)
	out.DisableWebhook = direct.ValueOf(in.DisableWebhook)
	out.AnalyzeQueryTextSentiment = direct.ValueOf(in.AnalyzeQueryTextSentiment)
	out.WebhookHeaders = in.WebhookHeaders
	out.FlowVersions = in.FlowVersions
	out.CurrentPlaybook = direct.ValueOf(in.CurrentPlaybook)
	out.LlmModelSettings = LlmModelSettings_ToProto(mapCtx, in.LlmModelSettings)
	out.Channel = direct.ValueOf(in.Channel)
	out.SessionTtl = direct.StringDuration_ToProto(mapCtx, in.SessionTtl)
	out.EndUserMetadata = EndUserMetadata_ToProto(mapCtx, in.EndUserMetadata)
	out.SearchConfig = SearchConfig_ToProto(mapCtx, in.SearchConfig)
	out.PopulateDataStoreConnectionSignals = direct.ValueOf(in.PopulateDataStoreConnectionSignals)
	return out
}
func QueryResult_FromProto(mapCtx *direct.MapContext, in *pb.QueryResult) *krm.QueryResult {
	if in == nil {
		return nil
	}
	out := &krm.QueryResult{}
	out.Text = direct.LazyPtr(in.GetText())
	out.TriggerIntent = direct.LazyPtr(in.GetTriggerIntent())
	out.Transcript = direct.LazyPtr(in.GetTranscript())
	out.TriggerEvent = direct.LazyPtr(in.GetTriggerEvent())
	out.Dtmf = DtmfInput_FromProto(mapCtx, in.GetDtmf())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.Parameters = Parameters_FromProto(mapCtx, in.GetParameters())
	out.ResponseMessages = direct.Slice_FromProto(mapCtx, in.ResponseMessages, ResponseMessage_FromProto)
	out.WebhookIds = in.WebhookIds
	out.WebhookDisplayNames = in.WebhookDisplayNames
	out.WebhookLatencies = direct.Slice_FromProto(mapCtx, in.WebhookLatencies, string_FromProto)
	out.WebhookTags = in.WebhookTags
	out.WebhookStatuses = direct.Slice_FromProto(mapCtx, in.WebhookStatuses, Status_FromProto)
	out.WebhookPayloads = direct.Slice_FromProto(mapCtx, in.WebhookPayloads, map[string]string_FromProto)
	out.CurrentPage = Page_FromProto(mapCtx, in.GetCurrentPage())
	out.CurrentFlow = Flow_FromProto(mapCtx, in.GetCurrentFlow())
	out.Intent = Intent_FromProto(mapCtx, in.GetIntent())
	out.IntentDetectionConfidence = direct.LazyPtr(in.GetIntentDetectionConfidence())
	out.Match = Match_FromProto(mapCtx, in.GetMatch())
	out.DiagnosticInfo = DiagnosticInfo_FromProto(mapCtx, in.GetDiagnosticInfo())
	out.GenerativeInfo = GenerativeInfo_FromProto(mapCtx, in.GetGenerativeInfo())
	out.SentimentAnalysisResult = SentimentAnalysisResult_FromProto(mapCtx, in.GetSentimentAnalysisResult())
	out.AdvancedSettings = AdvancedSettings_FromProto(mapCtx, in.GetAdvancedSettings())
	out.AllowAnswerFeedback = direct.LazyPtr(in.GetAllowAnswerFeedback())
	out.DataStoreConnectionSignals = DataStoreConnectionSignals_FromProto(mapCtx, in.GetDataStoreConnectionSignals())
	return out
}
func QueryResult_ToProto(mapCtx *direct.MapContext, in *krm.QueryResult) *pb.QueryResult {
	if in == nil {
		return nil
	}
	out := &pb.QueryResult{}
	if oneof := QueryResult_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Query = oneof
	}
	if oneof := QueryResult_TriggerIntent_ToProto(mapCtx, in.TriggerIntent); oneof != nil {
		out.Query = oneof
	}
	if oneof := QueryResult_Transcript_ToProto(mapCtx, in.Transcript); oneof != nil {
		out.Query = oneof
	}
	if oneof := QueryResult_TriggerEvent_ToProto(mapCtx, in.TriggerEvent); oneof != nil {
		out.Query = oneof
	}
	if oneof := DtmfInput_ToProto(mapCtx, in.Dtmf); oneof != nil {
		out.Query = &pb.QueryResult_Dtmf{Dtmf: oneof}
	}
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.Parameters = Parameters_ToProto(mapCtx, in.Parameters)
	out.ResponseMessages = direct.Slice_ToProto(mapCtx, in.ResponseMessages, ResponseMessage_ToProto)
	out.WebhookIds = in.WebhookIds
	out.WebhookDisplayNames = in.WebhookDisplayNames
	out.WebhookLatencies = direct.Slice_ToProto(mapCtx, in.WebhookLatencies, string_ToProto)
	out.WebhookTags = in.WebhookTags
	out.WebhookStatuses = direct.Slice_ToProto(mapCtx, in.WebhookStatuses, Status_ToProto)
	out.WebhookPayloads = direct.Slice_ToProto(mapCtx, in.WebhookPayloads, map[string]string_ToProto)
	out.CurrentPage = Page_ToProto(mapCtx, in.CurrentPage)
	out.CurrentFlow = Flow_ToProto(mapCtx, in.CurrentFlow)
	out.Intent = Intent_ToProto(mapCtx, in.Intent)
	out.IntentDetectionConfidence = direct.ValueOf(in.IntentDetectionConfidence)
	out.Match = Match_ToProto(mapCtx, in.Match)
	out.DiagnosticInfo = DiagnosticInfo_ToProto(mapCtx, in.DiagnosticInfo)
	out.GenerativeInfo = GenerativeInfo_ToProto(mapCtx, in.GenerativeInfo)
	out.SentimentAnalysisResult = SentimentAnalysisResult_ToProto(mapCtx, in.SentimentAnalysisResult)
	out.AdvancedSettings = AdvancedSettings_ToProto(mapCtx, in.AdvancedSettings)
	out.AllowAnswerFeedback = direct.ValueOf(in.AllowAnswerFeedback)
	out.DataStoreConnectionSignals = DataStoreConnectionSignals_ToProto(mapCtx, in.DataStoreConnectionSignals)
	return out
}
func QueryResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QueryResult) *krm.QueryResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QueryResultObservedState{}
	// MISSING: Text
	// MISSING: TriggerIntent
	// MISSING: Transcript
	// MISSING: TriggerEvent
	// MISSING: Dtmf
	// MISSING: LanguageCode
	// MISSING: Parameters
	// MISSING: ResponseMessages
	// MISSING: WebhookIds
	// MISSING: WebhookDisplayNames
	// MISSING: WebhookLatencies
	// MISSING: WebhookTags
	// MISSING: WebhookStatuses
	// MISSING: WebhookPayloads
	// MISSING: CurrentPage
	// MISSING: CurrentFlow
	// MISSING: Intent
	// MISSING: IntentDetectionConfidence
	// MISSING: Match
	// MISSING: DiagnosticInfo
	out.GenerativeInfo = GenerativeInfoObservedState_FromProto(mapCtx, in.GetGenerativeInfo())
	// MISSING: SentimentAnalysisResult
	// MISSING: AdvancedSettings
	// MISSING: AllowAnswerFeedback
	// MISSING: DataStoreConnectionSignals
	return out
}
func QueryResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QueryResultObservedState) *pb.QueryResult {
	if in == nil {
		return nil
	}
	out := &pb.QueryResult{}
	// MISSING: Text
	// MISSING: TriggerIntent
	// MISSING: Transcript
	// MISSING: TriggerEvent
	// MISSING: Dtmf
	// MISSING: LanguageCode
	// MISSING: Parameters
	// MISSING: ResponseMessages
	// MISSING: WebhookIds
	// MISSING: WebhookDisplayNames
	// MISSING: WebhookLatencies
	// MISSING: WebhookTags
	// MISSING: WebhookStatuses
	// MISSING: WebhookPayloads
	// MISSING: CurrentPage
	// MISSING: CurrentFlow
	// MISSING: Intent
	// MISSING: IntentDetectionConfidence
	// MISSING: Match
	// MISSING: DiagnosticInfo
	out.GenerativeInfo = GenerativeInfoObservedState_ToProto(mapCtx, in.GenerativeInfo)
	// MISSING: SentimentAnalysisResult
	// MISSING: AdvancedSettings
	// MISSING: AllowAnswerFeedback
	// MISSING: DataStoreConnectionSignals
	return out
}
func ResponseMessage_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage) *krm.ResponseMessage {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage{}
	out.Text = ResponseMessage_Text_FromProto(mapCtx, in.GetText())
	out.Payload = Payload_FromProto(mapCtx, in.GetPayload())
	out.ConversationSuccess = ResponseMessage_ConversationSuccess_FromProto(mapCtx, in.GetConversationSuccess())
	out.OutputAudioText = ResponseMessage_OutputAudioText_FromProto(mapCtx, in.GetOutputAudioText())
	out.LiveAgentHandoff = ResponseMessage_LiveAgentHandoff_FromProto(mapCtx, in.GetLiveAgentHandoff())
	// MISSING: EndInteraction
	out.PlayAudio = ResponseMessage_PlayAudio_FromProto(mapCtx, in.GetPlayAudio())
	// MISSING: MixedAudio
	out.TelephonyTransferCall = ResponseMessage_TelephonyTransferCall_FromProto(mapCtx, in.GetTelephonyTransferCall())
	out.KnowledgeInfoCard = ResponseMessage_KnowledgeInfoCard_FromProto(mapCtx, in.GetKnowledgeInfoCard())
	out.ToolCall = ToolCall_FromProto(mapCtx, in.GetToolCall())
	out.Channel = direct.LazyPtr(in.GetChannel())
	return out
}
func ResponseMessage_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage) *pb.ResponseMessage {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage{}
	if oneof := ResponseMessage_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Message = &pb.ResponseMessage_Text_{Text: oneof}
	}
	if oneof := Payload_ToProto(mapCtx, in.Payload); oneof != nil {
		out.Message = &pb.ResponseMessage_Payload{Payload: oneof}
	}
	if oneof := ResponseMessage_ConversationSuccess_ToProto(mapCtx, in.ConversationSuccess); oneof != nil {
		out.Message = &pb.ResponseMessage_ConversationSuccess_{ConversationSuccess: oneof}
	}
	if oneof := ResponseMessage_OutputAudioText_ToProto(mapCtx, in.OutputAudioText); oneof != nil {
		out.Message = &pb.ResponseMessage_OutputAudioText_{OutputAudioText: oneof}
	}
	if oneof := ResponseMessage_LiveAgentHandoff_ToProto(mapCtx, in.LiveAgentHandoff); oneof != nil {
		out.Message = &pb.ResponseMessage_LiveAgentHandoff_{LiveAgentHandoff: oneof}
	}
	// MISSING: EndInteraction
	if oneof := ResponseMessage_PlayAudio_ToProto(mapCtx, in.PlayAudio); oneof != nil {
		out.Message = &pb.ResponseMessage_PlayAudio_{PlayAudio: oneof}
	}
	// MISSING: MixedAudio
	if oneof := ResponseMessage_TelephonyTransferCall_ToProto(mapCtx, in.TelephonyTransferCall); oneof != nil {
		out.Message = &pb.ResponseMessage_TelephonyTransferCall_{TelephonyTransferCall: oneof}
	}
	if oneof := ResponseMessage_KnowledgeInfoCard_ToProto(mapCtx, in.KnowledgeInfoCard); oneof != nil {
		out.Message = &pb.ResponseMessage_KnowledgeInfoCard_{KnowledgeInfoCard: oneof}
	}
	if oneof := ToolCall_ToProto(mapCtx, in.ToolCall); oneof != nil {
		out.Message = &pb.ResponseMessage_ToolCall{ToolCall: oneof}
	}
	out.Channel = direct.ValueOf(in.Channel)
	return out
}
func ResponseMessageObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage) *krm.ResponseMessageObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessageObservedState{}
	out.Text = ResponseMessage_TextObservedState_FromProto(mapCtx, in.GetText())
	// MISSING: Payload
	// MISSING: ConversationSuccess
	out.OutputAudioText = ResponseMessage_OutputAudioTextObservedState_FromProto(mapCtx, in.GetOutputAudioText())
	// MISSING: LiveAgentHandoff
	out.EndInteraction = ResponseMessage_EndInteraction_FromProto(mapCtx, in.GetEndInteraction())
	out.PlayAudio = ResponseMessage_PlayAudioObservedState_FromProto(mapCtx, in.GetPlayAudio())
	out.MixedAudio = ResponseMessage_MixedAudio_FromProto(mapCtx, in.GetMixedAudio())
	// MISSING: TelephonyTransferCall
	// MISSING: KnowledgeInfoCard
	// MISSING: ToolCall
	// MISSING: Channel
	return out
}
func ResponseMessageObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessageObservedState) *pb.ResponseMessage {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage{}
	if oneof := ResponseMessage_TextObservedState_ToProto(mapCtx, in.Text); oneof != nil {
		out.Message = &pb.ResponseMessage_Text_{Text: oneof}
	}
	// MISSING: Payload
	// MISSING: ConversationSuccess
	if oneof := ResponseMessage_OutputAudioTextObservedState_ToProto(mapCtx, in.OutputAudioText); oneof != nil {
		out.Message = &pb.ResponseMessage_OutputAudioText_{OutputAudioText: oneof}
	}
	// MISSING: LiveAgentHandoff
	if oneof := ResponseMessage_EndInteraction_ToProto(mapCtx, in.EndInteraction); oneof != nil {
		out.Message = &pb.ResponseMessage_EndInteraction_{EndInteraction: oneof}
	}
	if oneof := ResponseMessage_PlayAudioObservedState_ToProto(mapCtx, in.PlayAudio); oneof != nil {
		out.Message = &pb.ResponseMessage_PlayAudio_{PlayAudio: oneof}
	}
	if oneof := ResponseMessage_MixedAudio_ToProto(mapCtx, in.MixedAudio); oneof != nil {
		out.Message = &pb.ResponseMessage_MixedAudio_{MixedAudio: oneof}
	}
	// MISSING: TelephonyTransferCall
	// MISSING: KnowledgeInfoCard
	// MISSING: ToolCall
	// MISSING: Channel
	return out
}
func ResponseMessage_ConversationSuccess_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_ConversationSuccess) *krm.ResponseMessage_ConversationSuccess {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_ConversationSuccess{}
	out.Metadata = Metadata_FromProto(mapCtx, in.GetMetadata())
	return out
}
func ResponseMessage_ConversationSuccess_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_ConversationSuccess) *pb.ResponseMessage_ConversationSuccess {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_ConversationSuccess{}
	out.Metadata = Metadata_ToProto(mapCtx, in.Metadata)
	return out
}
func ResponseMessage_EndInteraction_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_EndInteraction) *krm.ResponseMessage_EndInteraction {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_EndInteraction{}
	return out
}
func ResponseMessage_EndInteraction_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_EndInteraction) *pb.ResponseMessage_EndInteraction {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_EndInteraction{}
	return out
}
func ResponseMessage_KnowledgeInfoCard_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_KnowledgeInfoCard) *krm.ResponseMessage_KnowledgeInfoCard {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_KnowledgeInfoCard{}
	return out
}
func ResponseMessage_KnowledgeInfoCard_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_KnowledgeInfoCard) *pb.ResponseMessage_KnowledgeInfoCard {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_KnowledgeInfoCard{}
	return out
}
func ResponseMessage_LiveAgentHandoff_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_LiveAgentHandoff) *krm.ResponseMessage_LiveAgentHandoff {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_LiveAgentHandoff{}
	out.Metadata = Metadata_FromProto(mapCtx, in.GetMetadata())
	return out
}
func ResponseMessage_LiveAgentHandoff_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_LiveAgentHandoff) *pb.ResponseMessage_LiveAgentHandoff {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_LiveAgentHandoff{}
	out.Metadata = Metadata_ToProto(mapCtx, in.Metadata)
	return out
}
func ResponseMessage_MixedAudio_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_MixedAudio) *krm.ResponseMessage_MixedAudio {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_MixedAudio{}
	out.Segments = direct.Slice_FromProto(mapCtx, in.Segments, ResponseMessage_MixedAudio_Segment_FromProto)
	return out
}
func ResponseMessage_MixedAudio_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_MixedAudio) *pb.ResponseMessage_MixedAudio {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_MixedAudio{}
	out.Segments = direct.Slice_ToProto(mapCtx, in.Segments, ResponseMessage_MixedAudio_Segment_ToProto)
	return out
}
func ResponseMessage_MixedAudioObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_MixedAudio) *krm.ResponseMessage_MixedAudioObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_MixedAudioObservedState{}
	out.Segments = direct.Slice_FromProto(mapCtx, in.Segments, ResponseMessage_MixedAudio_SegmentObservedState_FromProto)
	return out
}
func ResponseMessage_MixedAudioObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_MixedAudioObservedState) *pb.ResponseMessage_MixedAudio {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_MixedAudio{}
	out.Segments = direct.Slice_ToProto(mapCtx, in.Segments, ResponseMessage_MixedAudio_SegmentObservedState_ToProto)
	return out
}
func ResponseMessage_MixedAudio_Segment_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_MixedAudio_Segment) *krm.ResponseMessage_MixedAudio_Segment {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_MixedAudio_Segment{}
	out.Audio = in.GetAudio()
	out.URI = direct.LazyPtr(in.GetUri())
	// MISSING: AllowPlaybackInterruption
	return out
}
func ResponseMessage_MixedAudio_Segment_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_MixedAudio_Segment) *pb.ResponseMessage_MixedAudio_Segment {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_MixedAudio_Segment{}
	if oneof := ResponseMessage_MixedAudio_Segment_Audio_ToProto(mapCtx, in.Audio); oneof != nil {
		out.Content = oneof
	}
	if oneof := ResponseMessage_MixedAudio_Segment_Uri_ToProto(mapCtx, in.URI); oneof != nil {
		out.Content = oneof
	}
	// MISSING: AllowPlaybackInterruption
	return out
}
func ResponseMessage_MixedAudio_SegmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_MixedAudio_Segment) *krm.ResponseMessage_MixedAudio_SegmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_MixedAudio_SegmentObservedState{}
	// MISSING: Audio
	// MISSING: URI
	out.AllowPlaybackInterruption = direct.LazyPtr(in.GetAllowPlaybackInterruption())
	return out
}
func ResponseMessage_MixedAudio_SegmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_MixedAudio_SegmentObservedState) *pb.ResponseMessage_MixedAudio_Segment {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_MixedAudio_Segment{}
	// MISSING: Audio
	// MISSING: URI
	out.AllowPlaybackInterruption = direct.ValueOf(in.AllowPlaybackInterruption)
	return out
}
func ResponseMessage_OutputAudioText_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_OutputAudioText) *krm.ResponseMessage_OutputAudioText {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_OutputAudioText{}
	out.Text = direct.LazyPtr(in.GetText())
	out.Ssml = direct.LazyPtr(in.GetSsml())
	// MISSING: AllowPlaybackInterruption
	return out
}
func ResponseMessage_OutputAudioText_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_OutputAudioText) *pb.ResponseMessage_OutputAudioText {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_OutputAudioText{}
	if oneof := ResponseMessage_OutputAudioText_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Source = oneof
	}
	if oneof := ResponseMessage_OutputAudioText_Ssml_ToProto(mapCtx, in.Ssml); oneof != nil {
		out.Source = oneof
	}
	// MISSING: AllowPlaybackInterruption
	return out
}
func ResponseMessage_OutputAudioTextObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_OutputAudioText) *krm.ResponseMessage_OutputAudioTextObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_OutputAudioTextObservedState{}
	// MISSING: Text
	// MISSING: Ssml
	out.AllowPlaybackInterruption = direct.LazyPtr(in.GetAllowPlaybackInterruption())
	return out
}
func ResponseMessage_OutputAudioTextObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_OutputAudioTextObservedState) *pb.ResponseMessage_OutputAudioText {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_OutputAudioText{}
	// MISSING: Text
	// MISSING: Ssml
	out.AllowPlaybackInterruption = direct.ValueOf(in.AllowPlaybackInterruption)
	return out
}
func ResponseMessage_PlayAudio_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_PlayAudio) *krm.ResponseMessage_PlayAudio {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_PlayAudio{}
	out.AudioURI = direct.LazyPtr(in.GetAudioUri())
	// MISSING: AllowPlaybackInterruption
	return out
}
func ResponseMessage_PlayAudio_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_PlayAudio) *pb.ResponseMessage_PlayAudio {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_PlayAudio{}
	out.AudioUri = direct.ValueOf(in.AudioURI)
	// MISSING: AllowPlaybackInterruption
	return out
}
func ResponseMessage_PlayAudioObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_PlayAudio) *krm.ResponseMessage_PlayAudioObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_PlayAudioObservedState{}
	// MISSING: AudioURI
	out.AllowPlaybackInterruption = direct.LazyPtr(in.GetAllowPlaybackInterruption())
	return out
}
func ResponseMessage_PlayAudioObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_PlayAudioObservedState) *pb.ResponseMessage_PlayAudio {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_PlayAudio{}
	// MISSING: AudioURI
	out.AllowPlaybackInterruption = direct.ValueOf(in.AllowPlaybackInterruption)
	return out
}
func ResponseMessage_TelephonyTransferCall_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_TelephonyTransferCall) *krm.ResponseMessage_TelephonyTransferCall {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_TelephonyTransferCall{}
	out.PhoneNumber = direct.LazyPtr(in.GetPhoneNumber())
	return out
}
func ResponseMessage_TelephonyTransferCall_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_TelephonyTransferCall) *pb.ResponseMessage_TelephonyTransferCall {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_TelephonyTransferCall{}
	if oneof := ResponseMessage_TelephonyTransferCall_PhoneNumber_ToProto(mapCtx, in.PhoneNumber); oneof != nil {
		out.Endpoint = oneof
	}
	return out
}
func ResponseMessage_Text_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_Text) *krm.ResponseMessage_Text {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_Text{}
	out.Text = in.Text
	// MISSING: AllowPlaybackInterruption
	return out
}
func ResponseMessage_Text_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_Text) *pb.ResponseMessage_Text {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_Text{}
	out.Text = in.Text
	// MISSING: AllowPlaybackInterruption
	return out
}
func ResponseMessage_TextObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_Text) *krm.ResponseMessage_TextObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_TextObservedState{}
	// MISSING: Text
	out.AllowPlaybackInterruption = direct.LazyPtr(in.GetAllowPlaybackInterruption())
	return out
}
func ResponseMessage_TextObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_TextObservedState) *pb.ResponseMessage_Text {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_Text{}
	// MISSING: Text
	out.AllowPlaybackInterruption = direct.ValueOf(in.AllowPlaybackInterruption)
	return out
}
func SearchConfig_FromProto(mapCtx *direct.MapContext, in *pb.SearchConfig) *krm.SearchConfig {
	if in == nil {
		return nil
	}
	out := &krm.SearchConfig{}
	out.BoostSpecs = direct.Slice_FromProto(mapCtx, in.BoostSpecs, BoostSpecs_FromProto)
	out.FilterSpecs = direct.Slice_FromProto(mapCtx, in.FilterSpecs, FilterSpecs_FromProto)
	return out
}
func SearchConfig_ToProto(mapCtx *direct.MapContext, in *krm.SearchConfig) *pb.SearchConfig {
	if in == nil {
		return nil
	}
	out := &pb.SearchConfig{}
	out.BoostSpecs = direct.Slice_ToProto(mapCtx, in.BoostSpecs, BoostSpecs_ToProto)
	out.FilterSpecs = direct.Slice_ToProto(mapCtx, in.FilterSpecs, FilterSpecs_ToProto)
	return out
}
func SentimentAnalysisResult_FromProto(mapCtx *direct.MapContext, in *pb.SentimentAnalysisResult) *krm.SentimentAnalysisResult {
	if in == nil {
		return nil
	}
	out := &krm.SentimentAnalysisResult{}
	out.Score = direct.LazyPtr(in.GetScore())
	out.Magnitude = direct.LazyPtr(in.GetMagnitude())
	return out
}
func SentimentAnalysisResult_ToProto(mapCtx *direct.MapContext, in *krm.SentimentAnalysisResult) *pb.SentimentAnalysisResult {
	if in == nil {
		return nil
	}
	out := &pb.SentimentAnalysisResult{}
	out.Score = direct.ValueOf(in.Score)
	out.Magnitude = direct.ValueOf(in.Magnitude)
	return out
}
func SessionEntityType_FromProto(mapCtx *direct.MapContext, in *pb.SessionEntityType) *krm.SessionEntityType {
	if in == nil {
		return nil
	}
	out := &krm.SessionEntityType{}
	out.Name = direct.LazyPtr(in.GetName())
	out.EntityOverrideMode = direct.Enum_FromProto(mapCtx, in.GetEntityOverrideMode())
	out.Entities = direct.Slice_FromProto(mapCtx, in.Entities, EntityType_Entity_FromProto)
	return out
}
func SessionEntityType_ToProto(mapCtx *direct.MapContext, in *krm.SessionEntityType) *pb.SessionEntityType {
	if in == nil {
		return nil
	}
	out := &pb.SessionEntityType{}
	out.Name = direct.ValueOf(in.Name)
	out.EntityOverrideMode = direct.Enum_ToProto[pb.SessionEntityType_EntityOverrideMode](mapCtx, in.EntityOverrideMode)
	out.Entities = direct.Slice_ToProto(mapCtx, in.Entities, EntityType_Entity_ToProto)
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
func TextInput_FromProto(mapCtx *direct.MapContext, in *pb.TextInput) *krm.TextInput {
	if in == nil {
		return nil
	}
	out := &krm.TextInput{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func TextInput_ToProto(mapCtx *direct.MapContext, in *krm.TextInput) *pb.TextInput {
	if in == nil {
		return nil
	}
	out := &pb.TextInput{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
func ToolCall_FromProto(mapCtx *direct.MapContext, in *pb.ToolCall) *krm.ToolCall {
	if in == nil {
		return nil
	}
	out := &krm.ToolCall{}
	out.Tool = direct.LazyPtr(in.GetTool())
	out.Action = direct.LazyPtr(in.GetAction())
	out.InputParameters = InputParameters_FromProto(mapCtx, in.GetInputParameters())
	return out
}
func ToolCall_ToProto(mapCtx *direct.MapContext, in *krm.ToolCall) *pb.ToolCall {
	if in == nil {
		return nil
	}
	out := &pb.ToolCall{}
	out.Tool = direct.ValueOf(in.Tool)
	out.Action = direct.ValueOf(in.Action)
	out.InputParameters = InputParameters_ToProto(mapCtx, in.InputParameters)
	return out
}
func ToolCallResult_FromProto(mapCtx *direct.MapContext, in *pb.ToolCallResult) *krm.ToolCallResult {
	if in == nil {
		return nil
	}
	out := &krm.ToolCallResult{}
	out.Tool = direct.LazyPtr(in.GetTool())
	out.Action = direct.LazyPtr(in.GetAction())
	out.Error = ToolCallResult_Error_FromProto(mapCtx, in.GetError())
	out.OutputParameters = OutputParameters_FromProto(mapCtx, in.GetOutputParameters())
	return out
}
func ToolCallResult_ToProto(mapCtx *direct.MapContext, in *krm.ToolCallResult) *pb.ToolCallResult {
	if in == nil {
		return nil
	}
	out := &pb.ToolCallResult{}
	out.Tool = direct.ValueOf(in.Tool)
	out.Action = direct.ValueOf(in.Action)
	if oneof := ToolCallResult_Error_ToProto(mapCtx, in.Error); oneof != nil {
		out.Result = &pb.ToolCallResult_Error_{Error: oneof}
	}
	if oneof := OutputParameters_ToProto(mapCtx, in.OutputParameters); oneof != nil {
		out.Result = &pb.ToolCallResult_OutputParameters{OutputParameters: oneof}
	}
	return out
}
func ToolCallResult_Error_FromProto(mapCtx *direct.MapContext, in *pb.ToolCallResult_Error) *krm.ToolCallResult_Error {
	if in == nil {
		return nil
	}
	out := &krm.ToolCallResult_Error{}
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func ToolCallResult_Error_ToProto(mapCtx *direct.MapContext, in *krm.ToolCallResult_Error) *pb.ToolCallResult_Error {
	if in == nil {
		return nil
	}
	out := &pb.ToolCallResult_Error{}
	out.Message = direct.ValueOf(in.Message)
	return out
}
func ToolUse_FromProto(mapCtx *direct.MapContext, in *pb.ToolUse) *krm.ToolUse {
	if in == nil {
		return nil
	}
	out := &krm.ToolUse{}
	out.Tool = direct.LazyPtr(in.GetTool())
	out.Action = direct.LazyPtr(in.GetAction())
	out.InputActionParameters = InputActionParameters_FromProto(mapCtx, in.GetInputActionParameters())
	out.OutputActionParameters = OutputActionParameters_FromProto(mapCtx, in.GetOutputActionParameters())
	return out
}
func ToolUse_ToProto(mapCtx *direct.MapContext, in *krm.ToolUse) *pb.ToolUse {
	if in == nil {
		return nil
	}
	out := &pb.ToolUse{}
	out.Tool = direct.ValueOf(in.Tool)
	out.Action = direct.ValueOf(in.Action)
	out.InputActionParameters = InputActionParameters_ToProto(mapCtx, in.InputActionParameters)
	out.OutputActionParameters = OutputActionParameters_ToProto(mapCtx, in.OutputActionParameters)
	return out
}
func TransitionRoute_FromProto(mapCtx *direct.MapContext, in *pb.TransitionRoute) *krm.TransitionRoute {
	if in == nil {
		return nil
	}
	out := &krm.TransitionRoute{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Intent = direct.LazyPtr(in.GetIntent())
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.TriggerFulfillment = Fulfillment_FromProto(mapCtx, in.GetTriggerFulfillment())
	out.TargetPage = direct.LazyPtr(in.GetTargetPage())
	out.TargetFlow = direct.LazyPtr(in.GetTargetFlow())
	return out
}
func TransitionRoute_ToProto(mapCtx *direct.MapContext, in *krm.TransitionRoute) *pb.TransitionRoute {
	if in == nil {
		return nil
	}
	out := &pb.TransitionRoute{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Intent = direct.ValueOf(in.Intent)
	out.Condition = direct.ValueOf(in.Condition)
	out.TriggerFulfillment = Fulfillment_ToProto(mapCtx, in.TriggerFulfillment)
	if oneof := TransitionRoute_TargetPage_ToProto(mapCtx, in.TargetPage); oneof != nil {
		out.Target = oneof
	}
	if oneof := TransitionRoute_TargetFlow_ToProto(mapCtx, in.TargetFlow); oneof != nil {
		out.Target = oneof
	}
	return out
}
func TransitionRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TransitionRoute) *krm.TransitionRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TransitionRouteObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	// MISSING: Intent
	// MISSING: Condition
	out.TriggerFulfillment = FulfillmentObservedState_FromProto(mapCtx, in.GetTriggerFulfillment())
	// MISSING: TargetPage
	// MISSING: TargetFlow
	return out
}
func TransitionRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TransitionRouteObservedState) *pb.TransitionRoute {
	if in == nil {
		return nil
	}
	out := &pb.TransitionRoute{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	// MISSING: Intent
	// MISSING: Condition
	out.TriggerFulfillment = FulfillmentObservedState_ToProto(mapCtx, in.TriggerFulfillment)
	// MISSING: TargetPage
	// MISSING: TargetFlow
	return out
}
func UserUtterance_FromProto(mapCtx *direct.MapContext, in *pb.UserUtterance) *krm.UserUtterance {
	if in == nil {
		return nil
	}
	out := &krm.UserUtterance{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func UserUtterance_ToProto(mapCtx *direct.MapContext, in *krm.UserUtterance) *pb.UserUtterance {
	if in == nil {
		return nil
	}
	out := &pb.UserUtterance{}
	out.Text = direct.ValueOf(in.Text)
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
func Webhook_FromProto(mapCtx *direct.MapContext, in *pb.Webhook) *krm.Webhook {
	if in == nil {
		return nil
	}
	out := &krm.Webhook{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.GenericWebService = Webhook_GenericWebService_FromProto(mapCtx, in.GetGenericWebService())
	out.ServiceDirectory = Webhook_ServiceDirectoryConfig_FromProto(mapCtx, in.GetServiceDirectory())
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}
func Webhook_ToProto(mapCtx *direct.MapContext, in *krm.Webhook) *pb.Webhook {
	if in == nil {
		return nil
	}
	out := &pb.Webhook{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if oneof := Webhook_GenericWebService_ToProto(mapCtx, in.GenericWebService); oneof != nil {
		out.Webhook = &pb.Webhook_GenericWebService_{GenericWebService: oneof}
	}
	if oneof := Webhook_ServiceDirectoryConfig_ToProto(mapCtx, in.ServiceDirectory); oneof != nil {
		out.Webhook = &pb.Webhook_ServiceDirectory{ServiceDirectory: oneof}
	}
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}
func Webhook_GenericWebService_FromProto(mapCtx *direct.MapContext, in *pb.Webhook_GenericWebService) *krm.Webhook_GenericWebService {
	if in == nil {
		return nil
	}
	out := &krm.Webhook_GenericWebService{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.RequestHeaders = in.RequestHeaders
	out.AllowedCaCerts = in.AllowedCaCerts
	out.OauthConfig = Webhook_GenericWebService_OAuthConfig_FromProto(mapCtx, in.GetOauthConfig())
	out.ServiceAgentAuth = direct.Enum_FromProto(mapCtx, in.GetServiceAgentAuth())
	out.WebhookType = direct.Enum_FromProto(mapCtx, in.GetWebhookType())
	out.HTTPMethod = direct.Enum_FromProto(mapCtx, in.GetHttpMethod())
	out.RequestBody = direct.LazyPtr(in.GetRequestBody())
	out.ParameterMapping = in.ParameterMapping
	return out
}
func Webhook_GenericWebService_ToProto(mapCtx *direct.MapContext, in *krm.Webhook_GenericWebService) *pb.Webhook_GenericWebService {
	if in == nil {
		return nil
	}
	out := &pb.Webhook_GenericWebService{}
	out.Uri = direct.ValueOf(in.URI)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.RequestHeaders = in.RequestHeaders
	out.AllowedCaCerts = in.AllowedCaCerts
	out.OauthConfig = Webhook_GenericWebService_OAuthConfig_ToProto(mapCtx, in.OauthConfig)
	out.ServiceAgentAuth = direct.Enum_ToProto[pb.Webhook_GenericWebService_ServiceAgentAuth](mapCtx, in.ServiceAgentAuth)
	out.WebhookType = direct.Enum_ToProto[pb.Webhook_GenericWebService_WebhookType](mapCtx, in.WebhookType)
	out.HttpMethod = direct.Enum_ToProto[pb.Webhook_GenericWebService_HttpMethod](mapCtx, in.HTTPMethod)
	out.RequestBody = direct.ValueOf(in.RequestBody)
	out.ParameterMapping = in.ParameterMapping
	return out
}
func Webhook_GenericWebService_OAuthConfig_FromProto(mapCtx *direct.MapContext, in *pb.Webhook_GenericWebService_OAuthConfig) *krm.Webhook_GenericWebService_OAuthConfig {
	if in == nil {
		return nil
	}
	out := &krm.Webhook_GenericWebService_OAuthConfig{}
	out.ClientID = direct.LazyPtr(in.GetClientId())
	out.ClientSecret = direct.LazyPtr(in.GetClientSecret())
	out.TokenEndpoint = direct.LazyPtr(in.GetTokenEndpoint())
	out.Scopes = in.Scopes
	return out
}
func Webhook_GenericWebService_OAuthConfig_ToProto(mapCtx *direct.MapContext, in *krm.Webhook_GenericWebService_OAuthConfig) *pb.Webhook_GenericWebService_OAuthConfig {
	if in == nil {
		return nil
	}
	out := &pb.Webhook_GenericWebService_OAuthConfig{}
	out.ClientId = direct.ValueOf(in.ClientID)
	out.ClientSecret = direct.ValueOf(in.ClientSecret)
	out.TokenEndpoint = direct.ValueOf(in.TokenEndpoint)
	out.Scopes = in.Scopes
	return out
}
func Webhook_ServiceDirectoryConfig_FromProto(mapCtx *direct.MapContext, in *pb.Webhook_ServiceDirectoryConfig) *krm.Webhook_ServiceDirectoryConfig {
	if in == nil {
		return nil
	}
	out := &krm.Webhook_ServiceDirectoryConfig{}
	out.Service = direct.LazyPtr(in.GetService())
	out.GenericWebService = Webhook_GenericWebService_FromProto(mapCtx, in.GetGenericWebService())
	return out
}
func Webhook_ServiceDirectoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.Webhook_ServiceDirectoryConfig) *pb.Webhook_ServiceDirectoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Webhook_ServiceDirectoryConfig{}
	out.Service = direct.ValueOf(in.Service)
	out.GenericWebService = Webhook_GenericWebService_ToProto(mapCtx, in.GenericWebService)
	return out
}
