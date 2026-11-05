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
	dialogflowpb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
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

func ResponseMessage_MixedAudio_SegmentObservedState_Audio_ToProto(mapCtx *direct.MapContext, in []byte) *pb.ResponseMessage_MixedAudio_Segment_Audio {
	if in == nil {
		return nil
	}
	return &pb.ResponseMessage_MixedAudio_Segment_Audio{Audio: in}
}

func ArticleSuggestionModelMetadata_FromProto(mapCtx *direct.MapContext, in *dialogflowpb.ArticleSuggestionModelMetadata) *krm.ArticleSuggestionModelMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ArticleSuggestionModelMetadata{}
	out.TrainingModelType = direct.Enum_FromProto(mapCtx, in.GetTrainingModelType())
	return out
}

func ArticleSuggestionModelMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ArticleSuggestionModelMetadata) *dialogflowpb.ArticleSuggestionModelMetadata {
	if in == nil {
		return nil
	}
	out := &dialogflowpb.ArticleSuggestionModelMetadata{}
	out.TrainingModelType = direct.Enum_ToProto[dialogflowpb.ConversationModel_ModelType](mapCtx, in.TrainingModelType)
	return out
}

func SmartReplyModelMetadata_FromProto(mapCtx *direct.MapContext, in *dialogflowpb.SmartReplyModelMetadata) *krm.SmartReplyModelMetadata {
	if in == nil {
		return nil
	}
	out := &krm.SmartReplyModelMetadata{}
	out.TrainingModelType = direct.Enum_FromProto(mapCtx, in.GetTrainingModelType())
	return out
}

func SmartReplyModelMetadata_ToProto(mapCtx *direct.MapContext, in *krm.SmartReplyModelMetadata) *dialogflowpb.SmartReplyModelMetadata {
	if in == nil {
		return nil
	}
	out := &dialogflowpb.SmartReplyModelMetadata{}
	out.TrainingModelType = direct.Enum_ToProto[dialogflowpb.ConversationModel_ModelType](mapCtx, in.TrainingModelType)
	return out
}
