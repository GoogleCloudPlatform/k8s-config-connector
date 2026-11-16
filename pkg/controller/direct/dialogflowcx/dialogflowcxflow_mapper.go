// Copyright 2026 Google LLC
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

package dialogflowcx

import (
	"fmt"
	"strings"

	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

func DialogflowCXFlowSpec_FromProto(mapCtx *direct.MapContext, in *pb.Flow) *krm.DialogflowCXFlowSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowSpec{}
	out.DisplayName = in.GetDisplayName()
	out.Description = direct.LazyPtr(in.GetDescription())
	out.TransitionRouteGroups = in.TransitionRouteGroups
	out.NluSettings = DialogflowCXFlowNluSettings_FromProto(mapCtx, in.GetNluSettings())
	for _, tr := range in.GetTransitionRoutes() {
		if tr != nil {
			out.TransitionRoutes = append(out.TransitionRoutes, *DialogflowCXFlowTransitionRoutes_FromProto(mapCtx, tr))
		}
	}
	for _, eh := range in.GetEventHandlers() {
		if eh != nil {
			out.EventHandlers = append(out.EventHandlers, *DialogflowCXFlowEventHandlers_FromProto(mapCtx, eh))
		}
	}
	return out
}

func DialogflowCXFlowSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowSpec) *pb.Flow {
	if in == nil {
		return nil
	}
	out := &pb.Flow{}
	out.DisplayName = in.DisplayName
	out.Description = direct.ValueOf(in.Description)
	out.TransitionRouteGroups = in.TransitionRouteGroups
	out.NluSettings = DialogflowCXFlowNluSettings_ToProto(mapCtx, in.NluSettings)
	for i := range in.TransitionRoutes {
		if tr := DialogflowCXFlowTransitionRoutes_ToProto(mapCtx, &in.TransitionRoutes[i]); tr != nil {
			out.TransitionRoutes = append(out.TransitionRoutes, tr)
		}
	}
	for i := range in.EventHandlers {
		if eh := DialogflowCXFlowEventHandlers_ToProto(mapCtx, &in.EventHandlers[i]); eh != nil {
			out.EventHandlers = append(out.EventHandlers, eh)
		}
	}
	return out
}

func DialogflowCXFlowNluSettings_FromProto(mapCtx *direct.MapContext, in *pb.NluSettings) *krm.DialogflowCXFlowNluSettings {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowNluSettings{}
	out.ClassificationThreshold = direct.LazyPtr(float64(in.GetClassificationThreshold()))
	out.ModelTrainingMode = direct.Enum_FromProto(mapCtx, in.GetModelTrainingMode())
	out.ModelType = direct.Enum_FromProto(mapCtx, in.GetModelType())
	return out
}

func DialogflowCXFlowNluSettings_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowNluSettings) *pb.NluSettings {
	if in == nil {
		return nil
	}
	out := &pb.NluSettings{}
	if in.ClassificationThreshold != nil {
		out.ClassificationThreshold = float32(*in.ClassificationThreshold)
	}
	out.ModelTrainingMode = direct.Enum_ToProto[pb.NluSettings_ModelTrainingMode](mapCtx, in.ModelTrainingMode)
	out.ModelType = direct.Enum_ToProto[pb.NluSettings_ModelType](mapCtx, in.ModelType)
	return out
}

func DialogflowCXFlowTransitionRoutes_FromProto(mapCtx *direct.MapContext, in *pb.TransitionRoute) *krm.DialogflowCXFlowTransitionRoutes {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowTransitionRoutes{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Intent = direct.LazyPtr(in.GetIntent())
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.TriggerFulfillment = DialogflowCXFlowTriggerFulfillment_FromProto(mapCtx, in.GetTriggerFulfillment())
	if in.GetTargetFlow() != "" {
		out.TargetFlow = direct.LazyPtr(in.GetTargetFlow())
	}
	if in.GetTargetPage() != "" {
		out.TargetPage = direct.LazyPtr(in.GetTargetPage())
	}
	return out
}

func DialogflowCXFlowTransitionRoutes_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowTransitionRoutes) *pb.TransitionRoute {
	if in == nil {
		return nil
	}
	out := &pb.TransitionRoute{}
	out.Name = direct.ValueOf(in.Name)
	out.Intent = direct.ValueOf(in.Intent)
	out.Condition = direct.ValueOf(in.Condition)
	out.TriggerFulfillment = DialogflowCXFlowTriggerFulfillment_ToProto(mapCtx, in.TriggerFulfillment)
	if in.TargetFlow != nil {
		out.Target = &pb.TransitionRoute_TargetFlow{TargetFlow: *in.TargetFlow}
	} else if in.TargetPage != nil {
		out.Target = &pb.TransitionRoute_TargetPage{TargetPage: *in.TargetPage}
	}
	return out
}

func DialogflowCXFlowEventHandlers_FromProto(mapCtx *direct.MapContext, in *pb.EventHandler) *krm.DialogflowCXFlowEventHandlers {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowEventHandlers{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Event = direct.LazyPtr(in.GetEvent())
	out.TriggerFulfillment = DialogflowCXFlowTriggerFulfillment_FromProto(mapCtx, in.GetTriggerFulfillment())
	if in.GetTargetFlow() != "" {
		out.TargetFlow = direct.LazyPtr(in.GetTargetFlow())
	}
	if in.GetTargetPage() != "" {
		out.TargetPage = direct.LazyPtr(in.GetTargetPage())
	}
	return out
}

func DialogflowCXFlowEventHandlers_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowEventHandlers) *pb.EventHandler {
	if in == nil {
		return nil
	}
	out := &pb.EventHandler{}
	out.Name = direct.ValueOf(in.Name)
	out.Event = direct.ValueOf(in.Event)
	out.TriggerFulfillment = DialogflowCXFlowTriggerFulfillment_ToProto(mapCtx, in.TriggerFulfillment)
	if in.TargetFlow != nil {
		out.Target = &pb.EventHandler_TargetFlow{TargetFlow: *in.TargetFlow}
	} else if in.TargetPage != nil {
		out.Target = &pb.EventHandler_TargetPage{TargetPage: *in.TargetPage}
	}
	return out
}

func DialogflowCXFlowTriggerFulfillment_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment) *krm.DialogflowCXFlowTriggerFulfillment {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowTriggerFulfillment{}
	out.Webhook = direct.LazyPtr(in.GetWebhook())
	out.Tag = direct.LazyPtr(in.GetTag())
	out.ReturnPartialResponses = direct.LazyPtr(in.GetReturnPartialResponses())
	for _, pa := range in.GetSetParameterActions() {
		if pa != nil {
			out.SetParameterActions = append(out.SetParameterActions, *DialogflowCXFlowSetParameterActions_FromProto(mapCtx, pa))
		}
	}
	for _, cc := range in.GetConditionalCases() {
		if cc != nil {
			out.ConditionalCases = append(out.ConditionalCases, *DialogflowCXFlowConditionalCases_FromProto(mapCtx, cc))
		}
	}
	for _, m := range in.GetMessages() {
		if m != nil {
			out.Messages = append(out.Messages, *DialogflowCXFlowMessages_FromProto(mapCtx, m))
		}
	}
	return out
}

func DialogflowCXFlowTriggerFulfillment_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowTriggerFulfillment) *pb.Fulfillment {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment{}
	out.Webhook = direct.ValueOf(in.Webhook)
	out.Tag = direct.ValueOf(in.Tag)
	out.ReturnPartialResponses = direct.ValueOf(in.ReturnPartialResponses)
	for i := range in.SetParameterActions {
		if pa := DialogflowCXFlowSetParameterActions_ToProto(mapCtx, &in.SetParameterActions[i]); pa != nil {
			out.SetParameterActions = append(out.SetParameterActions, pa)
		}
	}
	for i := range in.ConditionalCases {
		if cc := DialogflowCXFlowConditionalCases_ToProto(mapCtx, &in.ConditionalCases[i]); cc != nil {
			out.ConditionalCases = append(out.ConditionalCases, cc)
		}
	}
	for i := range in.Messages {
		if m := DialogflowCXFlowMessages_ToProto(mapCtx, &in.Messages[i]); m != nil {
			out.Messages = append(out.Messages, m)
		}
	}
	return out
}

func DialogflowCXFlowSetParameterActions_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment_SetParameterAction) *krm.DialogflowCXFlowSetParameterActions {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowSetParameterActions{}
	out.Parameter = direct.LazyPtr(in.GetParameter())
	if in.Value != nil {
		normalizeValue(in.Value)
		b, err := protojson.Marshal(in.Value)
		if err != nil {
			mapCtx.Errorf("marshalling set_parameter_actions value: %v", err)
		} else {
			out.Value = direct.LazyPtr(string(b))
		}
	}
	return out
}

func DialogflowCXFlowSetParameterActions_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowSetParameterActions) *pb.Fulfillment_SetParameterAction {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment_SetParameterAction{}
	out.Parameter = direct.ValueOf(in.Parameter)
	if in.Value != nil {
		v := &structpb.Value{}
		if err := protojson.Unmarshal([]byte(*in.Value), v); err != nil {
			mapCtx.Errorf("unmarshalling set_parameter_actions value: %v", err)
		} else {
			out.Value = v
		}
	}
	return out
}

func DialogflowCXFlowConditionalCases_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment_ConditionalCases) *krm.DialogflowCXFlowConditionalCases {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowConditionalCases{}
	if len(in.Cases) > 0 {
		var casesJSON []string
		for _, c := range in.Cases {
			normalizeConditionalCasesCase(c)
			b, err := protojson.Marshal(c)
			if err != nil {
				mapCtx.Errorf("marshalling conditional_cases case: %v", err)
				continue
			}
			casesJSON = append(casesJSON, string(b))
		}
		out.Cases = direct.LazyPtr("[" + strings.Join(casesJSON, ",") + "]")
	}
	return out
}

func DialogflowCXFlowConditionalCases_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowConditionalCases) *pb.Fulfillment_ConditionalCases {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment_ConditionalCases{}
	if in.Cases != nil && *in.Cases != "" {
		wrapped := fmt.Sprintf(`{"cases":%s}`, *in.Cases)
		if err := protojson.Unmarshal([]byte(wrapped), out); err != nil {
			mapCtx.Errorf("unmarshalling conditional_cases: %v", err)
		}
	}
	return out
}

func DialogflowCXFlowMessages_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage) *krm.DialogflowCXFlowMessages {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowMessages{}
	out.Channel = direct.LazyPtr(in.GetChannel())
	if text := in.GetText(); text != nil {
		out.Text = DialogflowCXFlowText_FromProto(mapCtx, text)
	}
	if payload := in.GetPayload(); payload != nil {
		normalizeStruct(payload)
		b, err := protojson.Marshal(payload)
		if err != nil {
			mapCtx.Errorf("marshalling payload: %v", err)
		} else {
			out.Payload = direct.LazyPtr(string(b))
		}
	}
	if cs := in.GetConversationSuccess(); cs != nil {
		out.ConversationSuccess = DialogflowCXFlowConversationSuccess_FromProto(mapCtx, cs)
	}
	if oat := in.GetOutputAudioText(); oat != nil {
		out.OutputAudioText = DialogflowCXFlowOutputAudioText_FromProto(mapCtx, oat)
	}
	if lah := in.GetLiveAgentHandoff(); lah != nil {
		out.LiveAgentHandoff = DialogflowCXFlowLiveAgentHandoff_FromProto(mapCtx, lah)
	}
	if pa := in.GetPlayAudio(); pa != nil {
		out.PlayAudio = DialogflowCXFlowPlayAudio_FromProto(mapCtx, pa)
	}
	if ttc := in.GetTelephonyTransferCall(); ttc != nil {
		out.TelephonyTransferCall = DialogflowCXFlowTelephonyTransferCall_FromProto(mapCtx, ttc)
	}
	return out
}

func DialogflowCXFlowMessages_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowMessages) *pb.ResponseMessage {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage{}
	out.Channel = direct.ValueOf(in.Channel)
	if in.Text != nil {
		out.Message = &pb.ResponseMessage_Text_{Text: DialogflowCXFlowText_ToProto(mapCtx, in.Text)}
	} else if in.Payload != nil {
		st := &structpb.Struct{}
		if err := protojson.Unmarshal([]byte(*in.Payload), st); err != nil {
			mapCtx.Errorf("unmarshalling payload: %v", err)
		} else {
			out.Message = &pb.ResponseMessage_Payload{Payload: st}
		}
	} else if in.ConversationSuccess != nil {
		out.Message = &pb.ResponseMessage_ConversationSuccess_{ConversationSuccess: DialogflowCXFlowConversationSuccess_ToProto(mapCtx, in.ConversationSuccess)}
	} else if in.OutputAudioText != nil {
		out.Message = &pb.ResponseMessage_OutputAudioText_{OutputAudioText: DialogflowCXFlowOutputAudioText_ToProto(mapCtx, in.OutputAudioText)}
	} else if in.LiveAgentHandoff != nil {
		out.Message = &pb.ResponseMessage_LiveAgentHandoff_{LiveAgentHandoff: DialogflowCXFlowLiveAgentHandoff_ToProto(mapCtx, in.LiveAgentHandoff)}
	} else if in.PlayAudio != nil {
		out.Message = &pb.ResponseMessage_PlayAudio_{PlayAudio: DialogflowCXFlowPlayAudio_ToProto(mapCtx, in.PlayAudio)}
	} else if in.TelephonyTransferCall != nil {
		out.Message = &pb.ResponseMessage_TelephonyTransferCall_{TelephonyTransferCall: DialogflowCXFlowTelephonyTransferCall_ToProto(mapCtx, in.TelephonyTransferCall)}
	}
	return out
}

func DialogflowCXFlowText_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_Text) *krm.DialogflowCXFlowText {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowText{}
	out.Text = in.Text
	out.AllowPlaybackInterruption = direct.LazyPtr(in.GetAllowPlaybackInterruption())
	return out
}

func DialogflowCXFlowText_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowText) *pb.ResponseMessage_Text {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_Text{}
	out.Text = in.Text
	out.AllowPlaybackInterruption = direct.ValueOf(in.AllowPlaybackInterruption)
	return out
}

func DialogflowCXFlowConversationSuccess_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_ConversationSuccess) *krm.DialogflowCXFlowConversationSuccess {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowConversationSuccess{}
	if in.Metadata != nil {
		normalizeStruct(in.Metadata)
		b, err := protojson.Marshal(in.Metadata)
		if err != nil {
			mapCtx.Errorf("marshalling conversation_success metadata: %v", err)
		} else {
			out.Metadata = direct.LazyPtr(string(b))
		}
	}
	return out
}

func DialogflowCXFlowConversationSuccess_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowConversationSuccess) *pb.ResponseMessage_ConversationSuccess {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_ConversationSuccess{}
	if in.Metadata != nil {
		st := &structpb.Struct{}
		if err := protojson.Unmarshal([]byte(*in.Metadata), st); err != nil {
			mapCtx.Errorf("unmarshalling conversation_success metadata: %v", err)
		} else {
			out.Metadata = st
		}
	}
	return out
}

func DialogflowCXFlowOutputAudioText_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_OutputAudioText) *krm.DialogflowCXFlowOutputAudioText {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowOutputAudioText{}
	out.AllowPlaybackInterruption = direct.LazyPtr(in.GetAllowPlaybackInterruption())
	if in.GetText() != "" {
		out.Text = direct.LazyPtr(in.GetText())
	}
	if in.GetSsml() != "" {
		out.Ssml = direct.LazyPtr(in.GetSsml())
	}
	return out
}

func DialogflowCXFlowOutputAudioText_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowOutputAudioText) *pb.ResponseMessage_OutputAudioText {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_OutputAudioText{}
	out.AllowPlaybackInterruption = direct.ValueOf(in.AllowPlaybackInterruption)
	if in.Text != nil {
		out.Source = &pb.ResponseMessage_OutputAudioText_Text{Text: *in.Text}
	} else if in.Ssml != nil {
		out.Source = &pb.ResponseMessage_OutputAudioText_Ssml{Ssml: *in.Ssml}
	}
	return out
}

func DialogflowCXFlowLiveAgentHandoff_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_LiveAgentHandoff) *krm.DialogflowCXFlowLiveAgentHandoff {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowLiveAgentHandoff{}
	if in.Metadata != nil {
		normalizeStruct(in.Metadata)
		b, err := protojson.Marshal(in.Metadata)
		if err != nil {
			mapCtx.Errorf("marshalling live_agent_handoff metadata: %v", err)
		} else {
			out.Metadata = direct.LazyPtr(string(b))
		}
	}
	return out
}

func DialogflowCXFlowLiveAgentHandoff_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowLiveAgentHandoff) *pb.ResponseMessage_LiveAgentHandoff {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_LiveAgentHandoff{}
	if in.Metadata != nil {
		st := &structpb.Struct{}
		if err := protojson.Unmarshal([]byte(*in.Metadata), st); err != nil {
			mapCtx.Errorf("unmarshalling live_agent_handoff metadata: %v", err)
		} else {
			out.Metadata = st
		}
	}
	return out
}

func DialogflowCXFlowPlayAudio_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_PlayAudio) *krm.DialogflowCXFlowPlayAudio {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowPlayAudio{}
	out.AudioUri = in.GetAudioUri()
	out.AllowPlaybackInterruption = direct.LazyPtr(in.GetAllowPlaybackInterruption())
	return out
}

func DialogflowCXFlowPlayAudio_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowPlayAudio) *pb.ResponseMessage_PlayAudio {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_PlayAudio{}
	out.AudioUri = in.AudioUri
	out.AllowPlaybackInterruption = direct.ValueOf(in.AllowPlaybackInterruption)
	return out
}

func DialogflowCXFlowTelephonyTransferCall_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_TelephonyTransferCall) *krm.DialogflowCXFlowTelephonyTransferCall {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXFlowTelephonyTransferCall{}
	out.PhoneNumber = in.GetPhoneNumber()
	return out
}

func DialogflowCXFlowTelephonyTransferCall_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowTelephonyTransferCall) *pb.ResponseMessage_TelephonyTransferCall {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_TelephonyTransferCall{}
	if in.PhoneNumber != "" {
		out.Endpoint = &pb.ResponseMessage_TelephonyTransferCall_PhoneNumber{PhoneNumber: in.PhoneNumber}
	}
	return out
}

func DialogflowCXFlowObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Flow) *krm.DialogflowCXFlowObservedState {
	if in == nil {
		return nil
	}
	return &krm.DialogflowCXFlowObservedState{}
}

func DialogflowCXFlowObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXFlowObservedState) *pb.Flow {
	if in == nil {
		return nil
	}
	return &pb.Flow{}
}

func normalizeStruct(s *structpb.Struct) {
	if s == nil {
		return
	}
	for _, v := range s.Fields {
		normalizeValue(v)
	}
}

func normalizeValue(v *structpb.Value) {
	if v == nil {
		return
	}
	if v.Kind == nil {
		v.Kind = &structpb.Value_NullValue{NullValue: structpb.NullValue_NULL_VALUE}
		return
	}
	switch k := v.Kind.(type) {
	case *structpb.Value_StructValue:
		if k.StructValue != nil {
			normalizeStruct(k.StructValue)
		}
	case *structpb.Value_ListValue:
		if k.ListValue != nil {
			for _, val := range k.ListValue.Values {
				normalizeValue(val)
			}
		}
	}
}

func normalizeFulfillment(f *pb.Fulfillment) {
	if f == nil {
		return
	}
	for _, pa := range f.SetParameterActions {
		if pa != nil {
			normalizeValue(pa.Value)
		}
	}
	for _, m := range f.Messages {
		if m != nil {
			if payload := m.GetPayload(); payload != nil {
				normalizeStruct(payload)
			}
			if cs := m.GetConversationSuccess(); cs != nil {
				normalizeStruct(cs.Metadata)
			}
			if lah := m.GetLiveAgentHandoff(); lah != nil {
				normalizeStruct(lah.Metadata)
			}
		}
	}
	for _, cc := range f.ConditionalCases {
		if cc != nil {
			for _, c := range cc.Cases {
				normalizeConditionalCasesCase(c)
			}
		}
	}
}

func normalizeConditionalCasesCase(c *pb.Fulfillment_ConditionalCases_Case) {
	if c == nil {
		return
	}
	for _, content := range c.CaseContent {
		if content == nil {
			continue
		}
		if msg := content.GetMessage(); msg != nil {
			if payload := msg.GetPayload(); payload != nil {
				normalizeStruct(payload)
			}
			if cs := msg.GetConversationSuccess(); cs != nil {
				normalizeStruct(cs.Metadata)
			}
			if lah := msg.GetLiveAgentHandoff(); lah != nil {
				normalizeStruct(lah.Metadata)
			}
		}
		if addCases := content.GetAdditionalCases(); addCases != nil {
			for _, nested := range addCases.Cases {
				normalizeConditionalCasesCase(nested)
			}
		}
	}
}
