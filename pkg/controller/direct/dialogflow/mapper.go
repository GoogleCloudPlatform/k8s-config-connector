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

package dialogflow

import (
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
)

func Value_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *krm.Value {
	return nil
}

func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *structpb.Value {
	return nil
}

func ResponseMessage_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage) *pb.ResponseMessage {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage{}
	if oneof := ResponseMessage_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Message = &pb.ResponseMessage_Text_{Text: oneof}
	}
	if oneof := direct.Struct_ToProto(mapCtx, &in.Payload); oneof != nil {
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
	if oneof := ResponseMessage_PlayAudio_ToProto(mapCtx, in.PlayAudio); oneof != nil {
		out.Message = &pb.ResponseMessage_PlayAudio_{PlayAudio: oneof}
	}
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

func DialogflowCX_ResponseMessage_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCX_ResponseMessage) *pb.ResponseMessage {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage{}
	if oneof := ResponseMessage_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Message = &pb.ResponseMessage_Text_{Text: oneof}
	}
	if oneof := direct.Struct_ToProto(mapCtx, &in.Payload); oneof != nil {
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
	if oneof := ResponseMessage_PlayAudio_ToProto(mapCtx, in.PlayAudio); oneof != nil {
		out.Message = &pb.ResponseMessage_PlayAudio_{PlayAudio: oneof}
	}
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

func ResponseMessage_MixedAudio_SegmentObservedState_Audio_ToProto(mapCtx *direct.MapContext, in []byte) *pb.ResponseMessage_MixedAudio_Segment_Audio {
	if in == nil {
		return nil
	}
	return &pb.ResponseMessage_MixedAudio_Segment_Audio{Audio: in}
}

func DialogflowCX_EventHandler_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCX_EventHandler) *pb.EventHandler {
	if in == nil {
		return nil
	}
	out := &pb.EventHandler{}
	out.Event = direct.ValueOf(in.Event)
	out.TriggerFulfillment = DialogflowCX_Fulfillment_ToProto(mapCtx, in.TriggerFulfillment)
	if in.TargetPageRef != nil {
		out.Target = &pb.EventHandler_TargetPage{TargetPage: in.TargetPageRef.External}
	}
	if in.TargetFlowRef != nil {
		out.Target = &pb.EventHandler_TargetFlow{TargetFlow: in.TargetFlowRef.External}
	}
	return out
}

func DialogflowCX_KnowledgeConnectorSettings_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCX_KnowledgeConnectorSettings) *pb.KnowledgeConnectorSettings {
	if in == nil {
		return nil
	}
	out := &pb.KnowledgeConnectorSettings{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.TriggerFulfillment = DialogflowCX_Fulfillment_ToProto(mapCtx, in.TriggerFulfillment)
	if in.TargetPageRef != nil {
		out.Target = &pb.KnowledgeConnectorSettings_TargetPage{TargetPage: in.TargetPageRef.External}
	}
	if in.TargetFlowRef != nil {
		out.Target = &pb.KnowledgeConnectorSettings_TargetFlow{TargetFlow: in.TargetFlowRef.External}
	}
	out.DataStoreConnections = direct.Slice_ToProto(mapCtx, in.DataStoreConnections, DialogflowCX_DataStoreConnection_ToProto)
	return out
}

func DialogflowCX_TransitionRoute_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCX_TransitionRoute) *pb.TransitionRoute {
	if in == nil {
		return nil
	}
	out := &pb.TransitionRoute{}
	out.Description = direct.ValueOf(in.Description)
	if in.IntentRef != nil {
		out.Intent = in.IntentRef.External
	}
	out.Condition = direct.ValueOf(in.Condition)
	out.TriggerFulfillment = DialogflowCX_Fulfillment_ToProto(mapCtx, in.TriggerFulfillment)
	if in.TargetPageRef != nil {
		out.Target = &pb.TransitionRoute_TargetPage{TargetPage: in.TargetPageRef.External}
	}
	if in.TargetFlowRef != nil {
		out.Target = &pb.TransitionRoute_TargetFlow{TargetFlow: in.TargetFlowRef.External}
	}
	return out
}
