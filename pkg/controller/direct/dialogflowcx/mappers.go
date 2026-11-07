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
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krmdialogflowcxv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

func structToString(mapCtx *direct.MapContext, in *structpb.Struct) *string {
	if in == nil {
		return nil
	}
	bytes, err := protojson.Marshal(in)
	if err != nil {
		mapCtx.Errorf("marshaling Struct to JSON: %v", err)
		return nil
	}
	s := string(bytes)
	return &s
}

func stringToStruct(mapCtx *direct.MapContext, in *string) *structpb.Struct {
	if in == nil {
		return nil
	}
	out := &structpb.Struct{}
	if err := protojson.Unmarshal([]byte(*in), out); err != nil {
		mapCtx.Errorf("unmarshaling Struct from JSON: %v", err)
		return nil
	}
	return out
}

func Parameter_DefaultValue_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *string {
	if in == nil {
		return nil
	}
	bytes, err := protojson.Marshal(in)
	if err != nil {
		mapCtx.Errorf("marshaling Value to JSON: %v", err)
		return nil
	}
	s := string(bytes)
	return &s
}

func Parameter_DefaultValue_ToProto(mapCtx *direct.MapContext, in *string) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if err := protojson.Unmarshal([]byte(*in), out); err != nil {
		return structpb.NewStringValue(*in)
	}
	return out
}

func SetParameterAction_Value_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *string {
	if in == nil {
		return nil
	}
	bytes, err := protojson.Marshal(in)
	if err != nil {
		mapCtx.Errorf("marshaling Value to JSON: %v", err)
		return nil
	}
	s := string(bytes)
	return &s
}

func SetParameterAction_Value_ToProto(mapCtx *direct.MapContext, in *string) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if err := protojson.Unmarshal([]byte(*in), out); err != nil {
		return structpb.NewStringValue(*in)
	}
	return out
}

func string_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment_ConditionalCases_Case) *string {
	if in == nil {
		return nil
	}
	bytes, err := protojson.Marshal(in)
	if err != nil {
		mapCtx.Errorf("marshaling Case to JSON: %v", err)
		return nil
	}
	s := string(bytes)
	return &s
}

func string_v1alpha1_ToProto(mapCtx *direct.MapContext, in *string) *pb.Fulfillment_ConditionalCases_Case {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment_ConditionalCases_Case{}
	if err := protojson.Unmarshal([]byte(*in), out); err != nil {
		mapCtx.Errorf("unmarshaling Case from JSON: %v", err)
		return nil
	}
	return out
}

func ResponseMessage_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage) *krmdialogflowcxv1alpha1.ResponseMessage {
	if in == nil {
		return nil
	}
	out := &krmdialogflowcxv1alpha1.ResponseMessage{}
	out.Text = ResponseMessage_Text_v1alpha1_FromProto(mapCtx, in.GetText())
	out.Payload = structToString(mapCtx, in.GetPayload())
	out.ConversationSuccess = ResponseMessage_ConversationSuccess_v1alpha1_FromProto(mapCtx, in.GetConversationSuccess())
	out.OutputAudioText = ResponseMessage_OutputAudioText_v1alpha1_FromProto(mapCtx, in.GetOutputAudioText())
	out.LiveAgentHandoff = ResponseMessage_LiveAgentHandoff_v1alpha1_FromProto(mapCtx, in.GetLiveAgentHandoff())
	out.PlayAudio = ResponseMessage_PlayAudio_v1alpha1_FromProto(mapCtx, in.GetPlayAudio())
	out.TelephonyTransferCall = ResponseMessage_TelephonyTransferCall_v1alpha1_FromProto(mapCtx, in.GetTelephonyTransferCall())
	out.Channel = direct.LazyPtr(in.GetChannel())
	return out
}

func ResponseMessage_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmdialogflowcxv1alpha1.ResponseMessage) *pb.ResponseMessage {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage{}
	if oneof := ResponseMessage_Text_v1alpha1_ToProto(mapCtx, in.Text); oneof != nil {
		out.Message = &pb.ResponseMessage_Text_{Text: oneof}
	}
	if oneof := stringToStruct(mapCtx, in.Payload); oneof != nil {
		out.Message = &pb.ResponseMessage_Payload{Payload: oneof}
	}
	if oneof := ResponseMessage_ConversationSuccess_v1alpha1_ToProto(mapCtx, in.ConversationSuccess); oneof != nil {
		out.Message = &pb.ResponseMessage_ConversationSuccess_{ConversationSuccess: oneof}
	}
	if oneof := ResponseMessage_OutputAudioText_v1alpha1_ToProto(mapCtx, in.OutputAudioText); oneof != nil {
		out.Message = &pb.ResponseMessage_OutputAudioText_{OutputAudioText: oneof}
	}
	if oneof := ResponseMessage_LiveAgentHandoff_v1alpha1_ToProto(mapCtx, in.LiveAgentHandoff); oneof != nil {
		out.Message = &pb.ResponseMessage_LiveAgentHandoff_{LiveAgentHandoff: oneof}
	}
	if oneof := ResponseMessage_PlayAudio_v1alpha1_ToProto(mapCtx, in.PlayAudio); oneof != nil {
		out.Message = &pb.ResponseMessage_PlayAudio_{PlayAudio: oneof}
	}
	if oneof := ResponseMessage_TelephonyTransferCall_v1alpha1_ToProto(mapCtx, in.TelephonyTransferCall); oneof != nil {
		out.Message = &pb.ResponseMessage_TelephonyTransferCall_{TelephonyTransferCall: oneof}
	}
	out.Channel = direct.ValueOf(in.Channel)
	return out
}

func ResponseMessage_ConversationSuccess_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_ConversationSuccess) *krmdialogflowcxv1alpha1.ResponseMessage_ConversationSuccess {
	if in == nil {
		return nil
	}
	out := &krmdialogflowcxv1alpha1.ResponseMessage_ConversationSuccess{}
	out.Metadata = structToString(mapCtx, in.GetMetadata())
	return out
}

func ResponseMessage_ConversationSuccess_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmdialogflowcxv1alpha1.ResponseMessage_ConversationSuccess) *pb.ResponseMessage_ConversationSuccess {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_ConversationSuccess{}
	out.Metadata = stringToStruct(mapCtx, in.Metadata)
	return out
}

func ResponseMessage_LiveAgentHandoff_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_LiveAgentHandoff) *krmdialogflowcxv1alpha1.ResponseMessage_LiveAgentHandoff {
	if in == nil {
		return nil
	}
	out := &krmdialogflowcxv1alpha1.ResponseMessage_LiveAgentHandoff{}
	out.Metadata = structToString(mapCtx, in.GetMetadata())
	return out
}

func ResponseMessage_LiveAgentHandoff_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmdialogflowcxv1alpha1.ResponseMessage_LiveAgentHandoff) *pb.ResponseMessage_LiveAgentHandoff {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_LiveAgentHandoff{}
	out.Metadata = stringToStruct(mapCtx, in.Metadata)
	return out
}

func ResponseMessage_TelephonyTransferCall_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_TelephonyTransferCall) *krmdialogflowcxv1alpha1.ResponseMessage_TelephonyTransferCall {
	if in == nil {
		return nil
	}
	out := &krmdialogflowcxv1alpha1.ResponseMessage_TelephonyTransferCall{}
	out.PhoneNumber = in.GetPhoneNumber()
	return out
}

func ResponseMessage_TelephonyTransferCall_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmdialogflowcxv1alpha1.ResponseMessage_TelephonyTransferCall) *pb.ResponseMessage_TelephonyTransferCall {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_TelephonyTransferCall{}
	if in.PhoneNumber != "" {
		out.Endpoint = &pb.ResponseMessage_TelephonyTransferCall_PhoneNumber{PhoneNumber: in.PhoneNumber}
	}
	return out
}

func ResponseMessage_TelephonyTransferCall_PhoneNumber_ToProto(mapCtx *direct.MapContext, in string) *pb.ResponseMessage_TelephonyTransferCall_PhoneNumber {
	if in == "" {
		return nil
	}
	return &pb.ResponseMessage_TelephonyTransferCall_PhoneNumber{PhoneNumber: in}
}

func DialogflowCXPageSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Page) *krmdialogflowcxv1alpha1.DialogflowCXPageSpec {
	if in == nil {
		return nil
	}
	out := &krmdialogflowcxv1alpha1.DialogflowCXPageSpec{}
	out.DisplayName = in.GetDisplayName()
	out.EntryFulfillment = Fulfillment_v1alpha1_FromProto(mapCtx, in.GetEntryFulfillment())
	out.Form = Form_v1alpha1_FromProto(mapCtx, in.GetForm())
	out.TransitionRouteGroups = in.TransitionRouteGroups
	out.TransitionRoutes = direct.Slice_FromProto(mapCtx, in.TransitionRoutes, TransitionRoute_v1alpha1_FromProto)
	out.EventHandlers = direct.Slice_FromProto(mapCtx, in.EventHandlers, EventHandler_v1alpha1_FromProto)
	return out
}

func DialogflowCXPageSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmdialogflowcxv1alpha1.DialogflowCXPageSpec) *pb.Page {
	if in == nil {
		return nil
	}
	out := &pb.Page{}
	out.DisplayName = in.DisplayName
	out.EntryFulfillment = Fulfillment_v1alpha1_ToProto(mapCtx, in.EntryFulfillment)
	out.Form = Form_v1alpha1_ToProto(mapCtx, in.Form)
	out.TransitionRouteGroups = in.TransitionRouteGroups
	out.TransitionRoutes = direct.Slice_ToProto(mapCtx, in.TransitionRoutes, TransitionRoute_v1alpha1_ToProto)
	out.EventHandlers = direct.Slice_ToProto(mapCtx, in.EventHandlers, EventHandler_v1alpha1_ToProto)
	return out
}

func ResponseMessage_PlayAudio_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_PlayAudio) *krmdialogflowcxv1alpha1.ResponseMessage_PlayAudio {
	if in == nil {
		return nil
	}
	out := &krmdialogflowcxv1alpha1.ResponseMessage_PlayAudio{}
	out.AudioUri = in.GetAudioUri()
	out.AllowPlaybackInterruption = direct.LazyPtr(in.GetAllowPlaybackInterruption())
	return out
}

func ResponseMessage_PlayAudio_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmdialogflowcxv1alpha1.ResponseMessage_PlayAudio) *pb.ResponseMessage_PlayAudio {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_PlayAudio{}
	out.AudioUri = in.AudioUri
	out.AllowPlaybackInterruption = direct.ValueOf(in.AllowPlaybackInterruption)
	return out
}
