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
	"encoding/json"

	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

func DialogflowCXPageSpec_FromProto(mapCtx *direct.MapContext, in *pb.Page) *krm.DialogflowCXPageSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXPageSpec{}
	out.DisplayName = in.GetDisplayName()
	out.EntryFulfillment = Fulfillment_FromProto(mapCtx, in.GetEntryFulfillment())
	out.Form = Form_FromProto(mapCtx, in.GetForm())
	out.TransitionRouteGroups = in.GetTransitionRouteGroups()
	out.TransitionRoutes = direct.Slice_FromProto(mapCtx, in.GetTransitionRoutes(), TransitionRoute_FromProto)
	out.EventHandlers = direct.Slice_FromProto(mapCtx, in.GetEventHandlers(), EventHandler_FromProto)
	return out
}

func DialogflowCXPageSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXPageSpec) *pb.Page {
	if in == nil {
		return nil
	}
	out := &pb.Page{}
	out.DisplayName = in.DisplayName
	out.EntryFulfillment = Fulfillment_ToProto(mapCtx, in.EntryFulfillment)
	out.Form = Form_ToProto(mapCtx, in.Form)
	out.TransitionRouteGroups = in.TransitionRouteGroups
	out.TransitionRoutes = direct.Slice_ToProto(mapCtx, in.TransitionRoutes, TransitionRoute_ToProto)
	out.EventHandlers = direct.Slice_ToProto(mapCtx, in.EventHandlers, EventHandler_ToProto)
	return out
}

func DialogflowCXPageObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Page) *krm.DialogflowCXPageObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXPageObservedState{}
	return out
}

func DialogflowCXPageObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXPageObservedState) *pb.Page {
	if in == nil {
		return nil
	}
	out := &pb.Page{}
	return out
}

func Fulfillment_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment) *krm.Fulfillment {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment{}
	out.Messages = direct.Slice_FromProto(mapCtx, in.GetMessages(), ResponseMessage_FromProto)
	out.Webhook = direct.LazyPtr(in.GetWebhook())
	out.ReturnPartialResponses = direct.LazyPtr(in.GetReturnPartialResponses())
	out.Tag = direct.LazyPtr(in.GetTag())
	out.SetParameterActions = direct.Slice_FromProto(mapCtx, in.GetSetParameterActions(), Fulfillment_SetParameterAction_FromProto)
	out.ConditionalCases = direct.Slice_FromProto(mapCtx, in.GetConditionalCases(), Fulfillment_ConditionalCases_FromProto)
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
	return out
}

func Fulfillment_SetParameterAction_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment_SetParameterAction) *krm.Fulfillment_SetParameterAction {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment_SetParameterAction{}
	out.Parameter = direct.LazyPtr(in.GetParameter())
	if in.GetValue() != nil {
		b, err := protojson.Marshal(in.GetValue())
		if err != nil {
			mapCtx.Errorf("marshalling SetParameterAction.Value to json: %v", err)
		} else {
			out.Value = direct.LazyPtr(string(b))
		}
	}
	return out
}

func Fulfillment_SetParameterAction_ToProto(mapCtx *direct.MapContext, in *krm.Fulfillment_SetParameterAction) *pb.Fulfillment_SetParameterAction {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment_SetParameterAction{}
	out.Parameter = direct.ValueOf(in.Parameter)
	if in.Value != nil {
		val := &structpb.Value{}
		if err := protojson.Unmarshal([]byte(*in.Value), val); err != nil {
			mapCtx.Errorf("unmarshalling SetParameterAction.Value from json: %v", err)
		} else {
			out.Value = val
		}
	}
	return out
}

func Fulfillment_ConditionalCases_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment_ConditionalCases) *krm.Fulfillment_ConditionalCases {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment_ConditionalCases{}
	if in.Cases != nil {
		var rawCases []json.RawMessage
		for _, c := range in.Cases {
			b, err := protojson.Marshal(c)
			if err != nil {
				mapCtx.Errorf("marshalling case to json: %v", err)
				continue
			}
			rawCases = append(rawCases, b)
		}
		b, err := json.Marshal(rawCases)
		if err != nil {
			mapCtx.Errorf("marshalling cases json list: %v", err)
		} else {
			out.Cases = direct.LazyPtr(string(b))
		}
	}
	return out
}

func Fulfillment_ConditionalCases_ToProto(mapCtx *direct.MapContext, in *krm.Fulfillment_ConditionalCases) *pb.Fulfillment_ConditionalCases {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment_ConditionalCases{}
	if in.Cases != nil {
		var rawCases []json.RawMessage
		if err := json.Unmarshal([]byte(*in.Cases), &rawCases); err != nil {
			mapCtx.Errorf("unmarshalling cases json list: %v", err)
		} else {
			for _, raw := range rawCases {
				c := &pb.Fulfillment_ConditionalCases_Case{}
				if err := protojson.Unmarshal(raw, c); err != nil {
					mapCtx.Errorf("unmarshalling case from json: %v", err)
					continue
				}
				out.Cases = append(out.Cases, c)
			}
		}
	}
	return out
}

func ResponseMessage_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage) *krm.ResponseMessage {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage{}
	out.Channel = direct.LazyPtr(in.GetChannel())
	switch m := in.GetMessage().(type) {
	case *pb.ResponseMessage_Text_:
		out.Text = ResponseMessage_Text_FromProto(mapCtx, m.Text)
	case *pb.ResponseMessage_Payload:
		if m.Payload != nil {
			b, err := protojson.Marshal(m.Payload)
			if err != nil {
				mapCtx.Errorf("marshalling ResponseMessage.Payload to json: %v", err)
			} else {
				out.Payload = direct.LazyPtr(string(b))
			}
		}
	case *pb.ResponseMessage_ConversationSuccess_:
		out.ConversationSuccess = ResponseMessage_ConversationSuccess_FromProto(mapCtx, m.ConversationSuccess)
	case *pb.ResponseMessage_OutputAudioText_:
		out.OutputAudioText = ResponseMessage_OutputAudioText_FromProto(mapCtx, m.OutputAudioText)
	case *pb.ResponseMessage_LiveAgentHandoff_:
		out.LiveAgentHandoff = ResponseMessage_LiveAgentHandoff_FromProto(mapCtx, m.LiveAgentHandoff)
	case *pb.ResponseMessage_PlayAudio_:
		out.PlayAudio = ResponseMessage_PlayAudio_FromProto(mapCtx, m.PlayAudio)
	case *pb.ResponseMessage_TelephonyTransferCall_:
		out.TelephonyTransferCall = ResponseMessage_TelephonyTransferCall_FromProto(mapCtx, m.TelephonyTransferCall)
	}
	return out
}

func ResponseMessage_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage) *pb.ResponseMessage {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage{}
	out.Channel = direct.ValueOf(in.Channel)
	if in.Text != nil {
		out.Message = &pb.ResponseMessage_Text_{Text: ResponseMessage_Text_ToProto(mapCtx, in.Text)}
	} else if in.Payload != nil {
		st := &structpb.Struct{}
		if err := protojson.Unmarshal([]byte(*in.Payload), st); err != nil {
			mapCtx.Errorf("unmarshalling ResponseMessage.Payload from json: %v", err)
		} else {
			out.Message = &pb.ResponseMessage_Payload{Payload: st}
		}
	} else if in.ConversationSuccess != nil {
		out.Message = &pb.ResponseMessage_ConversationSuccess_{ConversationSuccess: ResponseMessage_ConversationSuccess_ToProto(mapCtx, in.ConversationSuccess)}
	} else if in.OutputAudioText != nil {
		out.Message = &pb.ResponseMessage_OutputAudioText_{OutputAudioText: ResponseMessage_OutputAudioText_ToProto(mapCtx, in.OutputAudioText)}
	} else if in.LiveAgentHandoff != nil {
		out.Message = &pb.ResponseMessage_LiveAgentHandoff_{LiveAgentHandoff: ResponseMessage_LiveAgentHandoff_ToProto(mapCtx, in.LiveAgentHandoff)}
	} else if in.PlayAudio != nil {
		out.Message = &pb.ResponseMessage_PlayAudio_{PlayAudio: ResponseMessage_PlayAudio_ToProto(mapCtx, in.PlayAudio)}
	} else if in.TelephonyTransferCall != nil {
		out.Message = &pb.ResponseMessage_TelephonyTransferCall_{TelephonyTransferCall: ResponseMessage_TelephonyTransferCall_ToProto(mapCtx, in.TelephonyTransferCall)}
	}
	return out
}

func ResponseMessage_ConversationSuccess_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_ConversationSuccess) *krm.ResponseMessage_ConversationSuccess {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_ConversationSuccess{}
	if in.GetMetadata() != nil {
		b, err := protojson.Marshal(in.GetMetadata())
		if err != nil {
			mapCtx.Errorf("marshalling ResponseMessage_ConversationSuccess.Metadata to json: %v", err)
		} else {
			out.Metadata = direct.LazyPtr(string(b))
		}
	}
	return out
}

func ResponseMessage_ConversationSuccess_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_ConversationSuccess) *pb.ResponseMessage_ConversationSuccess {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_ConversationSuccess{}
	if in.Metadata != nil {
		st := &structpb.Struct{}
		if err := protojson.Unmarshal([]byte(*in.Metadata), st); err != nil {
			mapCtx.Errorf("unmarshalling ResponseMessage_ConversationSuccess.Metadata from json: %v", err)
		} else {
			out.Metadata = st
		}
	}
	return out
}

func ResponseMessage_LiveAgentHandoff_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_LiveAgentHandoff) *krm.ResponseMessage_LiveAgentHandoff {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_LiveAgentHandoff{}
	if in.GetMetadata() != nil {
		b, err := protojson.Marshal(in.GetMetadata())
		if err != nil {
			mapCtx.Errorf("marshalling ResponseMessage_LiveAgentHandoff.Metadata to json: %v", err)
		} else {
			out.Metadata = direct.LazyPtr(string(b))
		}
	}
	return out
}

func ResponseMessage_LiveAgentHandoff_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_LiveAgentHandoff) *pb.ResponseMessage_LiveAgentHandoff {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_LiveAgentHandoff{}
	if in.Metadata != nil {
		st := &structpb.Struct{}
		if err := protojson.Unmarshal([]byte(*in.Metadata), st); err != nil {
			mapCtx.Errorf("unmarshalling ResponseMessage_LiveAgentHandoff.Metadata from json: %v", err)
		} else {
			out.Metadata = st
		}
	}
	return out
}

func ResponseMessage_OutputAudioText_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_OutputAudioText) *krm.ResponseMessage_OutputAudioText {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_OutputAudioText{}
	out.AllowPlaybackInterruption = direct.LazyPtr(in.GetAllowPlaybackInterruption())
	switch s := in.GetSource().(type) {
	case *pb.ResponseMessage_OutputAudioText_Text:
		out.Text = direct.LazyPtr(s.Text)
	case *pb.ResponseMessage_OutputAudioText_Ssml:
		out.Ssml = direct.LazyPtr(s.Ssml)
	}
	return out
}

func ResponseMessage_OutputAudioText_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_OutputAudioText) *pb.ResponseMessage_OutputAudioText {
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

func ResponseMessage_PlayAudio_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_PlayAudio) *krm.ResponseMessage_PlayAudio {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_PlayAudio{}
	out.AllowPlaybackInterruption = direct.LazyPtr(in.GetAllowPlaybackInterruption())
	out.AudioUri = in.GetAudioUri()
	return out
}

func ResponseMessage_PlayAudio_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_PlayAudio) *pb.ResponseMessage_PlayAudio {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_PlayAudio{}
	out.AllowPlaybackInterruption = direct.ValueOf(in.AllowPlaybackInterruption)
	out.AudioUri = in.AudioUri
	return out
}

func ResponseMessage_TelephonyTransferCall_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_TelephonyTransferCall) *krm.ResponseMessage_TelephonyTransferCall {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_TelephonyTransferCall{}
	out.PhoneNumber = in.GetPhoneNumber()
	return out
}

func ResponseMessage_TelephonyTransferCall_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_TelephonyTransferCall) *pb.ResponseMessage_TelephonyTransferCall {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_TelephonyTransferCall{}
	if in.PhoneNumber != "" {
		out.Endpoint = &pb.ResponseMessage_TelephonyTransferCall_PhoneNumber{PhoneNumber: in.PhoneNumber}
	}
	return out
}

func ResponseMessage_Text_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_Text) *krm.ResponseMessage_Text {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_Text{}
	out.AllowPlaybackInterruption = direct.LazyPtr(in.GetAllowPlaybackInterruption())
	out.Text = in.GetText()
	return out
}

func ResponseMessage_Text_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_Text) *pb.ResponseMessage_Text {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_Text{}
	out.AllowPlaybackInterruption = direct.ValueOf(in.AllowPlaybackInterruption)
	out.Text = in.Text
	return out
}

func EventHandler_FromProto(mapCtx *direct.MapContext, in *pb.EventHandler) *krm.EventHandler {
	if in == nil {
		return nil
	}
	out := &krm.EventHandler{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Event = direct.LazyPtr(in.GetEvent())
	switch t := in.GetTarget().(type) {
	case *pb.EventHandler_TargetPage:
		out.TargetPage = direct.LazyPtr(t.TargetPage)
	case *pb.EventHandler_TargetFlow:
		out.TargetFlow = direct.LazyPtr(t.TargetFlow)
	}
	out.TriggerFulfillment = Fulfillment_FromProto(mapCtx, in.GetTriggerFulfillment())
	return out
}

func EventHandler_ToProto(mapCtx *direct.MapContext, in *krm.EventHandler) *pb.EventHandler {
	if in == nil {
		return nil
	}
	out := &pb.EventHandler{}
	out.Name = direct.ValueOf(in.Name)
	out.Event = direct.ValueOf(in.Event)
	if in.TargetPage != nil {
		out.Target = &pb.EventHandler_TargetPage{TargetPage: *in.TargetPage}
	} else if in.TargetFlow != nil {
		out.Target = &pb.EventHandler_TargetFlow{TargetFlow: *in.TargetFlow}
	}
	out.TriggerFulfillment = Fulfillment_ToProto(mapCtx, in.TriggerFulfillment)
	return out
}

func Form_FromProto(mapCtx *direct.MapContext, in *pb.Form) *krm.Form {
	if in == nil {
		return nil
	}
	out := &krm.Form{}
	out.Parameters = direct.Slice_FromProto(mapCtx, in.GetParameters(), Form_Parameter_FromProto)
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
	if in.GetDefaultValue() != nil {
		b, err := protojson.Marshal(in.GetDefaultValue())
		if err != nil {
			mapCtx.Errorf("marshalling Form_Parameter.DefaultValue to json: %v", err)
		} else {
			out.DefaultValue = direct.LazyPtr(string(b))
		}
	}
	out.Redact = direct.LazyPtr(in.GetRedact())
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
	if in.DefaultValue != nil {
		val := &structpb.Value{}
		if err := protojson.Unmarshal([]byte(*in.DefaultValue), val); err != nil {
			mapCtx.Errorf("unmarshalling Form_Parameter.DefaultValue from json: %v", err)
		} else {
			out.DefaultValue = val
		}
	}
	out.Redact = direct.ValueOf(in.Redact)
	return out
}

func Form_Parameter_FillBehavior_FromProto(mapCtx *direct.MapContext, in *pb.Form_Parameter_FillBehavior) *krm.Form_Parameter_FillBehavior {
	if in == nil {
		return nil
	}
	out := &krm.Form_Parameter_FillBehavior{}
	out.InitialPromptFulfillment = Fulfillment_FromProto(mapCtx, in.GetInitialPromptFulfillment())
	out.RepromptEventHandlers = direct.Slice_FromProto(mapCtx, in.GetRepromptEventHandlers(), EventHandler_FromProto)
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

func TransitionRoute_FromProto(mapCtx *direct.MapContext, in *pb.TransitionRoute) *krm.TransitionRoute {
	if in == nil {
		return nil
	}
	out := &krm.TransitionRoute{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Intent = direct.LazyPtr(in.GetIntent())
	out.Condition = direct.LazyPtr(in.GetCondition())
	switch t := in.GetTarget().(type) {
	case *pb.TransitionRoute_TargetPage:
		out.TargetPage = direct.LazyPtr(t.TargetPage)
	case *pb.TransitionRoute_TargetFlow:
		out.TargetFlow = direct.LazyPtr(t.TargetFlow)
	}
	out.TriggerFulfillment = Fulfillment_FromProto(mapCtx, in.GetTriggerFulfillment())
	return out
}

func TransitionRoute_ToProto(mapCtx *direct.MapContext, in *krm.TransitionRoute) *pb.TransitionRoute {
	if in == nil {
		return nil
	}
	out := &pb.TransitionRoute{}
	out.Name = direct.ValueOf(in.Name)
	out.Intent = direct.ValueOf(in.Intent)
	out.Condition = direct.ValueOf(in.Condition)
	if in.TargetPage != nil {
		out.Target = &pb.TransitionRoute_TargetPage{TargetPage: *in.TargetPage}
	} else if in.TargetFlow != nil {
		out.Target = &pb.TransitionRoute_TargetFlow{TargetFlow: *in.TargetFlow}
	}
	out.TriggerFulfillment = Fulfillment_ToProto(mapCtx, in.TriggerFulfillment)
	return out
}
