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
func AnnotatedMessagePart_FromProto(mapCtx *direct.MapContext, in *pb.AnnotatedMessagePart) *krm.AnnotatedMessagePart {
	if in == nil {
		return nil
	}
	out := &krm.AnnotatedMessagePart{}
	out.Text = direct.LazyPtr(in.GetText())
	out.EntityType = direct.LazyPtr(in.GetEntityType())
	out.FormattedValue = Value_FromProto(mapCtx, in.GetFormattedValue())
	return out
}
func AnnotatedMessagePart_ToProto(mapCtx *direct.MapContext, in *krm.AnnotatedMessagePart) *pb.AnnotatedMessagePart {
	if in == nil {
		return nil
	}
	out := &pb.AnnotatedMessagePart{}
	out.Text = direct.ValueOf(in.Text)
	out.EntityType = direct.ValueOf(in.EntityType)
	out.FormattedValue = Value_ToProto(mapCtx, in.FormattedValue)
	return out
}
func Message_FromProto(mapCtx *direct.MapContext, in *pb.Message) *krm.Message {
	if in == nil {
		return nil
	}
	out := &krm.Message{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Content = direct.LazyPtr(in.GetContent())
	out.ResponseMessages = direct.Slice_FromProto(mapCtx, in.ResponseMessages, ResponseMessage_FromProto)
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	// MISSING: Participant
	// MISSING: ParticipantRole
	// MISSING: CreateTime
	out.SendTime = direct.StringTimestamp_FromProto(mapCtx, in.GetSendTime())
	// MISSING: MessageAnnotation
	// MISSING: SentimentAnalysis
	return out
}
func Message_ToProto(mapCtx *direct.MapContext, in *krm.Message) *pb.Message {
	if in == nil {
		return nil
	}
	out := &pb.Message{}
	out.Name = direct.ValueOf(in.Name)
	out.Content = direct.ValueOf(in.Content)
	out.ResponseMessages = direct.Slice_ToProto(mapCtx, in.ResponseMessages, ResponseMessage_ToProto)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	// MISSING: Participant
	// MISSING: ParticipantRole
	// MISSING: CreateTime
	out.SendTime = direct.StringTimestamp_ToProto(mapCtx, in.SendTime)
	// MISSING: MessageAnnotation
	// MISSING: SentimentAnalysis
	return out
}
func MessageAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.MessageAnnotation) *krm.MessageAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.MessageAnnotation{}
	out.Parts = direct.Slice_FromProto(mapCtx, in.Parts, AnnotatedMessagePart_FromProto)
	out.ContainEntities = direct.LazyPtr(in.GetContainEntities())
	return out
}
func MessageAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.MessageAnnotation) *pb.MessageAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.MessageAnnotation{}
	out.Parts = direct.Slice_ToProto(mapCtx, in.Parts, AnnotatedMessagePart_ToProto)
	out.ContainEntities = direct.ValueOf(in.ContainEntities)
	return out
}
func MessageObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Message) *krm.MessageObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MessageObservedState{}
	// MISSING: Name
	// MISSING: Content
	// MISSING: ResponseMessages
	// MISSING: LanguageCode
	out.Participant = direct.LazyPtr(in.GetParticipant())
	out.ParticipantRole = direct.Enum_FromProto(mapCtx, in.GetParticipantRole())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: SendTime
	out.MessageAnnotation = MessageAnnotation_FromProto(mapCtx, in.GetMessageAnnotation())
	out.SentimentAnalysis = SentimentAnalysisResult_FromProto(mapCtx, in.GetSentimentAnalysis())
	return out
}
func MessageObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MessageObservedState) *pb.Message {
	if in == nil {
		return nil
	}
	out := &pb.Message{}
	// MISSING: Name
	// MISSING: Content
	// MISSING: ResponseMessages
	// MISSING: LanguageCode
	out.Participant = direct.ValueOf(in.Participant)
	out.ParticipantRole = direct.Enum_ToProto[pb.Participant_Role](mapCtx, in.ParticipantRole)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: SendTime
	out.MessageAnnotation = MessageAnnotation_ToProto(mapCtx, in.MessageAnnotation)
	out.SentimentAnalysis = SentimentAnalysisResult_ToProto(mapCtx, in.SentimentAnalysis)
	return out
}
func ResponseMessage_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage) *krm.ResponseMessage {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage{}
	out.Text = ResponseMessage_Text_FromProto(mapCtx, in.GetText())
	out.Payload = Payload_FromProto(mapCtx, in.GetPayload())
	out.LiveAgentHandoff = ResponseMessage_LiveAgentHandoff_FromProto(mapCtx, in.GetLiveAgentHandoff())
	out.EndInteraction = ResponseMessage_EndInteraction_FromProto(mapCtx, in.GetEndInteraction())
	out.MixedAudio = ResponseMessage_MixedAudio_FromProto(mapCtx, in.GetMixedAudio())
	out.TelephonyTransferCall = ResponseMessage_TelephonyTransferCall_FromProto(mapCtx, in.GetTelephonyTransferCall())
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
	if oneof := ResponseMessage_LiveAgentHandoff_ToProto(mapCtx, in.LiveAgentHandoff); oneof != nil {
		out.Message = &pb.ResponseMessage_LiveAgentHandoff_{LiveAgentHandoff: oneof}
	}
	if oneof := ResponseMessage_EndInteraction_ToProto(mapCtx, in.EndInteraction); oneof != nil {
		out.Message = &pb.ResponseMessage_EndInteraction_{EndInteraction: oneof}
	}
	if oneof := ResponseMessage_MixedAudio_ToProto(mapCtx, in.MixedAudio); oneof != nil {
		out.Message = &pb.ResponseMessage_MixedAudio_{MixedAudio: oneof}
	}
	if oneof := ResponseMessage_TelephonyTransferCall_ToProto(mapCtx, in.TelephonyTransferCall); oneof != nil {
		out.Message = &pb.ResponseMessage_TelephonyTransferCall_{TelephonyTransferCall: oneof}
	}
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
func ResponseMessage_MixedAudio_Segment_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_MixedAudio_Segment) *krm.ResponseMessage_MixedAudio_Segment {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_MixedAudio_Segment{}
	out.Audio = in.GetAudio()
	out.URI = direct.LazyPtr(in.GetUri())
	out.AllowPlaybackInterruption = direct.LazyPtr(in.GetAllowPlaybackInterruption())
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
	out.AllowPlaybackInterruption = direct.ValueOf(in.AllowPlaybackInterruption)
	return out
}
func ResponseMessage_TelephonyTransferCall_FromProto(mapCtx *direct.MapContext, in *pb.ResponseMessage_TelephonyTransferCall) *krm.ResponseMessage_TelephonyTransferCall {
	if in == nil {
		return nil
	}
	out := &krm.ResponseMessage_TelephonyTransferCall{}
	out.PhoneNumber = direct.LazyPtr(in.GetPhoneNumber())
	out.SipURI = direct.LazyPtr(in.GetSipUri())
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
	if oneof := ResponseMessage_TelephonyTransferCall_SipUri_ToProto(mapCtx, in.SipURI); oneof != nil {
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
	return out
}
func ResponseMessage_Text_ToProto(mapCtx *direct.MapContext, in *krm.ResponseMessage_Text) *pb.ResponseMessage_Text {
	if in == nil {
		return nil
	}
	out := &pb.ResponseMessage_Text{}
	out.Text = in.Text
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
