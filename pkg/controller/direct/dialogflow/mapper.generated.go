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
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
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
func ConversationTurn_FromProto(mapCtx *direct.MapContext, in *pb.ConversationTurn) *krm.ConversationTurn {
	if in == nil {
		return nil
	}
	out := &krm.ConversationTurn{}
	out.UserInput = ConversationTurn_UserInput_FromProto(mapCtx, in.GetUserInput())
	out.VirtualAgentOutput = ConversationTurn_VirtualAgentOutput_FromProto(mapCtx, in.GetVirtualAgentOutput())
	return out
}
func ConversationTurn_ToProto(mapCtx *direct.MapContext, in *krm.ConversationTurn) *pb.ConversationTurn {
	if in == nil {
		return nil
	}
	out := &pb.ConversationTurn{}
	out.UserInput = ConversationTurn_UserInput_ToProto(mapCtx, in.UserInput)
	out.VirtualAgentOutput = ConversationTurn_VirtualAgentOutput_ToProto(mapCtx, in.VirtualAgentOutput)
	return out
}
func ConversationTurnObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversationTurn) *krm.ConversationTurnObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConversationTurnObservedState{}
	// MISSING: UserInput
	out.VirtualAgentOutput = ConversationTurn_VirtualAgentOutputObservedState_FromProto(mapCtx, in.GetVirtualAgentOutput())
	return out
}
func ConversationTurnObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConversationTurnObservedState) *pb.ConversationTurn {
	if in == nil {
		return nil
	}
	out := &pb.ConversationTurn{}
	// MISSING: UserInput
	out.VirtualAgentOutput = ConversationTurn_VirtualAgentOutputObservedState_ToProto(mapCtx, in.VirtualAgentOutput)
	return out
}
func ConversationTurn_UserInput_FromProto(mapCtx *direct.MapContext, in *pb.ConversationTurn_UserInput) *krm.ConversationTurn_UserInput {
	if in == nil {
		return nil
	}
	out := &krm.ConversationTurn_UserInput{}
	out.Input = QueryInput_FromProto(mapCtx, in.GetInput())
	out.InjectedParameters = InjectedParameters_FromProto(mapCtx, in.GetInjectedParameters())
	out.IsWebhookEnabled = direct.LazyPtr(in.GetIsWebhookEnabled())
	out.EnableSentimentAnalysis = direct.LazyPtr(in.GetEnableSentimentAnalysis())
	return out
}
func ConversationTurn_UserInput_ToProto(mapCtx *direct.MapContext, in *krm.ConversationTurn_UserInput) *pb.ConversationTurn_UserInput {
	if in == nil {
		return nil
	}
	out := &pb.ConversationTurn_UserInput{}
	out.Input = QueryInput_ToProto(mapCtx, in.Input)
	out.InjectedParameters = InjectedParameters_ToProto(mapCtx, in.InjectedParameters)
	out.IsWebhookEnabled = direct.ValueOf(in.IsWebhookEnabled)
	out.EnableSentimentAnalysis = direct.ValueOf(in.EnableSentimentAnalysis)
	return out
}
func ConversationTurn_VirtualAgentOutput_FromProto(mapCtx *direct.MapContext, in *pb.ConversationTurn_VirtualAgentOutput) *krm.ConversationTurn_VirtualAgentOutput {
	if in == nil {
		return nil
	}
	out := &krm.ConversationTurn_VirtualAgentOutput{}
	out.SessionParameters = SessionParameters_FromProto(mapCtx, in.GetSessionParameters())
	// MISSING: Differences
	out.DiagnosticInfo = DiagnosticInfo_FromProto(mapCtx, in.GetDiagnosticInfo())
	out.TriggeredIntent = Intent_FromProto(mapCtx, in.GetTriggeredIntent())
	out.CurrentPage = Page_FromProto(mapCtx, in.GetCurrentPage())
	out.TextResponses = direct.Slice_FromProto(mapCtx, in.TextResponses, ResponseMessage_Text_FromProto)
	out.Status = Status_FromProto(mapCtx, in.GetStatus())
	return out
}
func ConversationTurn_VirtualAgentOutput_ToProto(mapCtx *direct.MapContext, in *krm.ConversationTurn_VirtualAgentOutput) *pb.ConversationTurn_VirtualAgentOutput {
	if in == nil {
		return nil
	}
	out := &pb.ConversationTurn_VirtualAgentOutput{}
	out.SessionParameters = SessionParameters_ToProto(mapCtx, in.SessionParameters)
	// MISSING: Differences
	out.DiagnosticInfo = DiagnosticInfo_ToProto(mapCtx, in.DiagnosticInfo)
	out.TriggeredIntent = Intent_ToProto(mapCtx, in.TriggeredIntent)
	out.CurrentPage = Page_ToProto(mapCtx, in.CurrentPage)
	out.TextResponses = direct.Slice_ToProto(mapCtx, in.TextResponses, ResponseMessage_Text_ToProto)
	out.Status = Status_ToProto(mapCtx, in.Status)
	return out
}
func ConversationTurn_VirtualAgentOutputObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversationTurn_VirtualAgentOutput) *krm.ConversationTurn_VirtualAgentOutputObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConversationTurn_VirtualAgentOutputObservedState{}
	// MISSING: SessionParameters
	out.Differences = direct.Slice_FromProto(mapCtx, in.Differences, TestRunDifference_FromProto)
	// MISSING: DiagnosticInfo
	// MISSING: TriggeredIntent
	out.CurrentPage = PageObservedState_FromProto(mapCtx, in.GetCurrentPage())
	// MISSING: TextResponses
	// MISSING: Status
	return out
}
func ConversationTurn_VirtualAgentOutputObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConversationTurn_VirtualAgentOutputObservedState) *pb.ConversationTurn_VirtualAgentOutput {
	if in == nil {
		return nil
	}
	out := &pb.ConversationTurn_VirtualAgentOutput{}
	// MISSING: SessionParameters
	out.Differences = direct.Slice_ToProto(mapCtx, in.Differences, TestRunDifference_ToProto)
	// MISSING: DiagnosticInfo
	// MISSING: TriggeredIntent
	out.CurrentPage = PageObservedState_ToProto(mapCtx, in.CurrentPage)
	// MISSING: TextResponses
	// MISSING: Status
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
func DialogflowTestCaseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TestCase) *krm.DialogflowTestCaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowTestCaseObservedState{}
	// MISSING: Name
	// MISSING: Tags
	// MISSING: DisplayName
	// MISSING: Notes
	// MISSING: TestConfig
	// MISSING: TestCaseConversationTurns
	// MISSING: CreationTime
	// MISSING: LastTestResult
	return out
}
func DialogflowTestCaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowTestCaseObservedState) *pb.TestCase {
	if in == nil {
		return nil
	}
	out := &pb.TestCase{}
	// MISSING: Name
	// MISSING: Tags
	// MISSING: DisplayName
	// MISSING: Notes
	// MISSING: TestConfig
	// MISSING: TestCaseConversationTurns
	// MISSING: CreationTime
	// MISSING: LastTestResult
	return out
}
func DialogflowTestCaseSpec_FromProto(mapCtx *direct.MapContext, in *pb.TestCase) *krm.DialogflowTestCaseSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowTestCaseSpec{}
	// MISSING: Name
	// MISSING: Tags
	// MISSING: DisplayName
	// MISSING: Notes
	// MISSING: TestConfig
	// MISSING: TestCaseConversationTurns
	// MISSING: CreationTime
	// MISSING: LastTestResult
	return out
}
func DialogflowTestCaseSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowTestCaseSpec) *pb.TestCase {
	if in == nil {
		return nil
	}
	out := &pb.TestCase{}
	// MISSING: Name
	// MISSING: Tags
	// MISSING: DisplayName
	// MISSING: Notes
	// MISSING: TestConfig
	// MISSING: TestCaseConversationTurns
	// MISSING: CreationTime
	// MISSING: LastTestResult
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
func FormObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Form) *krm.FormObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FormObservedState{}
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, Form_ParameterObservedState_FromProto)
	return out
}
func FormObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FormObservedState) *pb.Form {
	if in == nil {
		return nil
	}
	out := &pb.Form{}
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, Form_ParameterObservedState_ToProto)
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
func Form_ParameterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Form_Parameter) *krm.Form_ParameterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Form_ParameterObservedState{}
	// MISSING: DisplayName
	// MISSING: Required
	// MISSING: EntityType
	// MISSING: IsList
	out.FillBehavior = Form_Parameter_FillBehaviorObservedState_FromProto(mapCtx, in.GetFillBehavior())
	// MISSING: DefaultValue
	// MISSING: Redact
	// MISSING: AdvancedSettings
	return out
}
func Form_ParameterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Form_ParameterObservedState) *pb.Form_Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Form_Parameter{}
	// MISSING: DisplayName
	// MISSING: Required
	// MISSING: EntityType
	// MISSING: IsList
	out.FillBehavior = Form_Parameter_FillBehaviorObservedState_ToProto(mapCtx, in.FillBehavior)
	// MISSING: DefaultValue
	// MISSING: Redact
	// MISSING: AdvancedSettings
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
func Form_Parameter_FillBehaviorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Form_Parameter_FillBehavior) *krm.Form_Parameter_FillBehaviorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Form_Parameter_FillBehaviorObservedState{}
	// MISSING: InitialPromptFulfillment
	out.RepromptEventHandlers = direct.Slice_FromProto(mapCtx, in.RepromptEventHandlers, EventHandlerObservedState_FromProto)
	return out
}
func Form_Parameter_FillBehaviorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Form_Parameter_FillBehaviorObservedState) *pb.Form_Parameter_FillBehavior {
	if in == nil {
		return nil
	}
	out := &pb.Form_Parameter_FillBehavior{}
	// MISSING: InitialPromptFulfillment
	out.RepromptEventHandlers = direct.Slice_ToProto(mapCtx, in.RepromptEventHandlers, EventHandlerObservedState_ToProto)
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
func PageObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Page) *krm.PageObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PageObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.EntryFulfillment = FulfillmentObservedState_FromProto(mapCtx, in.GetEntryFulfillment())
	out.Form = FormObservedState_FromProto(mapCtx, in.GetForm())
	// MISSING: TransitionRouteGroups
	out.TransitionRoutes = direct.Slice_FromProto(mapCtx, in.TransitionRoutes, TransitionRouteObservedState_FromProto)
	// MISSING: EventHandlers
	// MISSING: AdvancedSettings
	// MISSING: KnowledgeConnectorSettings
	return out
}
func PageObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PageObservedState) *pb.Page {
	if in == nil {
		return nil
	}
	out := &pb.Page{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.EntryFulfillment = FulfillmentObservedState_ToProto(mapCtx, in.EntryFulfillment)
	out.Form = FormObservedState_ToProto(mapCtx, in.Form)
	// MISSING: TransitionRouteGroups
	out.TransitionRoutes = direct.Slice_ToProto(mapCtx, in.TransitionRoutes, TransitionRouteObservedState_ToProto)
	// MISSING: EventHandlers
	// MISSING: AdvancedSettings
	// MISSING: KnowledgeConnectorSettings
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
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
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
	out.ResponseType = direct.Enum_FromProto(mapCtx, in.GetResponseType())
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
	out.ResponseType = direct.Enum_ToProto[pb.ResponseMessage_ResponseType](mapCtx, in.ResponseType)
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
	// MISSING: ResponseType
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
	// MISSING: ResponseType
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
func TestCase_FromProto(mapCtx *direct.MapContext, in *pb.TestCase) *krm.TestCase {
	if in == nil {
		return nil
	}
	out := &krm.TestCase{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Tags = in.Tags
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Notes = direct.LazyPtr(in.GetNotes())
	out.TestConfig = TestConfig_FromProto(mapCtx, in.GetTestConfig())
	out.TestCaseConversationTurns = direct.Slice_FromProto(mapCtx, in.TestCaseConversationTurns, ConversationTurn_FromProto)
	// MISSING: CreationTime
	out.LastTestResult = TestCaseResult_FromProto(mapCtx, in.GetLastTestResult())
	return out
}
func TestCase_ToProto(mapCtx *direct.MapContext, in *krm.TestCase) *pb.TestCase {
	if in == nil {
		return nil
	}
	out := &pb.TestCase{}
	out.Name = direct.ValueOf(in.Name)
	out.Tags = in.Tags
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Notes = direct.ValueOf(in.Notes)
	out.TestConfig = TestConfig_ToProto(mapCtx, in.TestConfig)
	out.TestCaseConversationTurns = direct.Slice_ToProto(mapCtx, in.TestCaseConversationTurns, ConversationTurn_ToProto)
	// MISSING: CreationTime
	out.LastTestResult = TestCaseResult_ToProto(mapCtx, in.LastTestResult)
	return out
}
func TestCaseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TestCase) *krm.TestCaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TestCaseObservedState{}
	// MISSING: Name
	// MISSING: Tags
	// MISSING: DisplayName
	// MISSING: Notes
	// MISSING: TestConfig
	out.TestCaseConversationTurns = direct.Slice_FromProto(mapCtx, in.TestCaseConversationTurns, ConversationTurnObservedState_FromProto)
	out.CreationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreationTime())
	// MISSING: LastTestResult
	return out
}
func TestCaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TestCaseObservedState) *pb.TestCase {
	if in == nil {
		return nil
	}
	out := &pb.TestCase{}
	// MISSING: Name
	// MISSING: Tags
	// MISSING: DisplayName
	// MISSING: Notes
	// MISSING: TestConfig
	out.TestCaseConversationTurns = direct.Slice_ToProto(mapCtx, in.TestCaseConversationTurns, ConversationTurnObservedState_ToProto)
	out.CreationTime = direct.StringTimestamp_ToProto(mapCtx, in.CreationTime)
	// MISSING: LastTestResult
	return out
}
func TestCaseResult_FromProto(mapCtx *direct.MapContext, in *pb.TestCaseResult) *krm.TestCaseResult {
	if in == nil {
		return nil
	}
	out := &krm.TestCaseResult{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Environment = direct.LazyPtr(in.GetEnvironment())
	out.ConversationTurns = direct.Slice_FromProto(mapCtx, in.ConversationTurns, ConversationTurn_FromProto)
	out.TestResult = direct.Enum_FromProto(mapCtx, in.GetTestResult())
	out.TestTime = direct.StringTimestamp_FromProto(mapCtx, in.GetTestTime())
	return out
}
func TestCaseResult_ToProto(mapCtx *direct.MapContext, in *krm.TestCaseResult) *pb.TestCaseResult {
	if in == nil {
		return nil
	}
	out := &pb.TestCaseResult{}
	out.Name = direct.ValueOf(in.Name)
	out.Environment = direct.ValueOf(in.Environment)
	out.ConversationTurns = direct.Slice_ToProto(mapCtx, in.ConversationTurns, ConversationTurn_ToProto)
	out.TestResult = direct.Enum_ToProto[pb.TestResult](mapCtx, in.TestResult)
	out.TestTime = direct.StringTimestamp_ToProto(mapCtx, in.TestTime)
	return out
}
func TestConfig_FromProto(mapCtx *direct.MapContext, in *pb.TestConfig) *krm.TestConfig {
	if in == nil {
		return nil
	}
	out := &krm.TestConfig{}
	out.TrackingParameters = in.TrackingParameters
	out.Flow = direct.LazyPtr(in.GetFlow())
	out.Page = direct.LazyPtr(in.GetPage())
	return out
}
func TestConfig_ToProto(mapCtx *direct.MapContext, in *krm.TestConfig) *pb.TestConfig {
	if in == nil {
		return nil
	}
	out := &pb.TestConfig{}
	out.TrackingParameters = in.TrackingParameters
	out.Flow = direct.ValueOf(in.Flow)
	out.Page = direct.ValueOf(in.Page)
	return out
}
func TestRunDifference_FromProto(mapCtx *direct.MapContext, in *pb.TestRunDifference) *krm.TestRunDifference {
	if in == nil {
		return nil
	}
	out := &krm.TestRunDifference{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func TestRunDifference_ToProto(mapCtx *direct.MapContext, in *krm.TestRunDifference) *pb.TestRunDifference {
	if in == nil {
		return nil
	}
	out := &pb.TestRunDifference{}
	out.Type = direct.Enum_ToProto[pb.TestRunDifference_DiffType](mapCtx, in.Type)
	out.Description = direct.ValueOf(in.Description)
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
	// MISSING: TriggerFulfillment
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
	// MISSING: TriggerFulfillment
	// MISSING: TargetPage
	// MISSING: TargetFlow
	return out
}
