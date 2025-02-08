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
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
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
func DialogflowMessageObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Message) *krm.DialogflowMessageObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowMessageObservedState{}
	// MISSING: Name
	// MISSING: Content
	// MISSING: LanguageCode
	// MISSING: Participant
	// MISSING: ParticipantRole
	// MISSING: CreateTime
	// MISSING: SendTime
	// MISSING: MessageAnnotation
	// MISSING: SentimentAnalysis
	return out
}
func DialogflowMessageObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowMessageObservedState) *pb.Message {
	if in == nil {
		return nil
	}
	out := &pb.Message{}
	// MISSING: Name
	// MISSING: Content
	// MISSING: LanguageCode
	// MISSING: Participant
	// MISSING: ParticipantRole
	// MISSING: CreateTime
	// MISSING: SendTime
	// MISSING: MessageAnnotation
	// MISSING: SentimentAnalysis
	return out
}
func DialogflowMessageSpec_FromProto(mapCtx *direct.MapContext, in *pb.Message) *krm.DialogflowMessageSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowMessageSpec{}
	// MISSING: Name
	// MISSING: Content
	// MISSING: LanguageCode
	// MISSING: Participant
	// MISSING: ParticipantRole
	// MISSING: CreateTime
	// MISSING: SendTime
	// MISSING: MessageAnnotation
	// MISSING: SentimentAnalysis
	return out
}
func DialogflowMessageSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowMessageSpec) *pb.Message {
	if in == nil {
		return nil
	}
	out := &pb.Message{}
	// MISSING: Name
	// MISSING: Content
	// MISSING: LanguageCode
	// MISSING: Participant
	// MISSING: ParticipantRole
	// MISSING: CreateTime
	// MISSING: SendTime
	// MISSING: MessageAnnotation
	// MISSING: SentimentAnalysis
	return out
}
func Message_FromProto(mapCtx *direct.MapContext, in *pb.Message) *krm.Message {
	if in == nil {
		return nil
	}
	out := &krm.Message{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Content = direct.LazyPtr(in.GetContent())
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
	// MISSING: LanguageCode
	out.Participant = direct.ValueOf(in.Participant)
	out.ParticipantRole = direct.Enum_ToProto[pb.Participant_Role](mapCtx, in.ParticipantRole)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: SendTime
	out.MessageAnnotation = MessageAnnotation_ToProto(mapCtx, in.MessageAnnotation)
	out.SentimentAnalysis = SentimentAnalysisResult_ToProto(mapCtx, in.SentimentAnalysis)
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
