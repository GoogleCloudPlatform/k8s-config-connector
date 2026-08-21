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

package contactcenterinsights

import (
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Conversation_Transcript_TranscriptSegment_DialogflowSegmentMetadata_FromProto(mapCtx *direct.MapContext, in *pb.Conversation_Transcript_TranscriptSegment_DialogflowSegmentMetadata) *krm.Conversation_Transcript_TranscriptSegment_DialogflowSegmentMetadata {
	if in == nil {
		return nil
	}
	out := &krm.Conversation_Transcript_TranscriptSegment_DialogflowSegmentMetadata{}
	out.SmartReplyAllowlistCovered = direct.LazyPtr(in.GetSmartReplyAllowlistCovered())
	return out
}

func Conversation_Transcript_TranscriptSegment_DialogflowSegmentMetadata_ToProto(mapCtx *direct.MapContext, in *krm.Conversation_Transcript_TranscriptSegment_DialogflowSegmentMetadata) *pb.Conversation_Transcript_TranscriptSegment_DialogflowSegmentMetadata {
	if in == nil {
		return nil
	}
	out := &pb.Conversation_Transcript_TranscriptSegment_DialogflowSegmentMetadata{}
	out.SmartReplyAllowlistCovered = direct.ValueOf(in.SmartReplyAllowlistCovered)
	return out
}

func Conversation_QualityMetadata_AgentInfo_FromProto(mapCtx *direct.MapContext, in *pb.Conversation_QualityMetadata_AgentInfo) *krm.Conversation_QualityMetadata_AgentInfo {
	if in == nil {
		return nil
	}
	out := &krm.Conversation_QualityMetadata_AgentInfo{}
	out.AgentID = direct.LazyPtr(in.GetAgentId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Team = direct.LazyPtr(in.GetTeam())
	out.DispositionCode = direct.LazyPtr(in.GetDispositionCode())
	out.AgentType = direct.Enum_FromProto(mapCtx, in.GetAgentType())
	return out
}

func Conversation_QualityMetadata_AgentInfo_ToProto(mapCtx *direct.MapContext, in *krm.Conversation_QualityMetadata_AgentInfo) *pb.Conversation_QualityMetadata_AgentInfo {
	if in == nil {
		return nil
	}
	out := &pb.Conversation_QualityMetadata_AgentInfo{}
	out.AgentId = direct.ValueOf(in.AgentID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Team = direct.ValueOf(in.Team)
	out.DispositionCode = direct.ValueOf(in.DispositionCode)
	out.AgentType = direct.Enum_ToProto[pb.ConversationParticipant_Role](mapCtx, in.AgentType)
	return out
}

func Conversation_QualityMetadata_FromProto(mapCtx *direct.MapContext, in *pb.Conversation_QualityMetadata) *krm.Conversation_QualityMetadata {
	if in == nil {
		return nil
	}
	out := &krm.Conversation_QualityMetadata{}
	out.CustomerSatisfactionRating = direct.LazyPtr(in.GetCustomerSatisfactionRating())
	out.WaitDuration = direct.StringDuration_FromProto(mapCtx, in.GetWaitDuration())
	out.MenuPath = direct.LazyPtr(in.GetMenuPath())
	out.AgentInfo = direct.Slice_FromProto(mapCtx, in.GetAgentInfo(), Conversation_QualityMetadata_AgentInfo_FromProto)
	return out
}

func Conversation_QualityMetadata_ToProto(mapCtx *direct.MapContext, in *krm.Conversation_QualityMetadata) *pb.Conversation_QualityMetadata {
	if in == nil {
		return nil
	}
	out := &pb.Conversation_QualityMetadata{}
	out.CustomerSatisfactionRating = direct.ValueOf(in.CustomerSatisfactionRating)
	out.WaitDuration = direct.StringDuration_ToProto(mapCtx, in.WaitDuration)
	out.MenuPath = direct.ValueOf(in.MenuPath)
	out.AgentInfo = direct.Slice_ToProto(mapCtx, in.AgentInfo, Conversation_QualityMetadata_AgentInfo_ToProto)
	return out
}

func Conversation_CallMetadata_FromProto(mapCtx *direct.MapContext, in *pb.Conversation_CallMetadata) *krm.Conversation_CallMetadata {
	if in == nil {
		return nil
	}
	out := &krm.Conversation_CallMetadata{}
	out.CustomerChannel = direct.LazyPtr(in.GetCustomerChannel())
	out.AgentChannel = direct.LazyPtr(in.GetAgentChannel())
	return out
}

func Conversation_CallMetadata_ToProto(mapCtx *direct.MapContext, in *krm.Conversation_CallMetadata) *pb.Conversation_CallMetadata {
	if in == nil {
		return nil
	}
	out := &pb.Conversation_CallMetadata{}
	out.CustomerChannel = direct.ValueOf(in.CustomerChannel)
	out.AgentChannel = direct.ValueOf(in.AgentChannel)
	return out
}

func AnalysisResult_CallAnalysisMetadata_FromProto(mapCtx *direct.MapContext, in *pb.AnalysisResult_CallAnalysisMetadata) *krm.AnalysisResult_CallAnalysisMetadata {
	if in == nil {
		return nil
	}
	out := &krm.AnalysisResult_CallAnalysisMetadata{}
	out.Annotations = direct.Slice_FromProto(mapCtx, in.GetAnnotations(), CallAnnotation_FromProto)
	out.Sentiments = direct.Slice_FromProto(mapCtx, in.GetSentiments(), ConversationLevelSentiment_FromProto)
	out.Silence = ConversationLevelSilence_FromProto(mapCtx, in.GetSilence())
	return out
}

func AnalysisResult_CallAnalysisMetadata_ToProto(mapCtx *direct.MapContext, in *krm.AnalysisResult_CallAnalysisMetadata) *pb.AnalysisResult_CallAnalysisMetadata {
	if in == nil {
		return nil
	}
	out := &pb.AnalysisResult_CallAnalysisMetadata{}
	out.Annotations = direct.Slice_ToProto(mapCtx, in.Annotations, CallAnnotation_ToProto)
	out.Sentiments = direct.Slice_ToProto(mapCtx, in.Sentiments, ConversationLevelSentiment_ToProto)
	out.Silence = ConversationLevelSilence_ToProto(mapCtx, in.Silence)
	return out
}

func AnalysisResult_CallAnalysisMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AnalysisResult_CallAnalysisMetadata) *krm.AnalysisResult_CallAnalysisMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalysisResult_CallAnalysisMetadataObservedState{}
	out.Annotations = direct.Slice_FromProto(mapCtx, in.GetAnnotations(), CallAnnotation_FromProto)
	out.Sentiments = direct.Slice_FromProto(mapCtx, in.GetSentiments(), ConversationLevelSentiment_FromProto)
	out.Silence = ConversationLevelSilence_FromProto(mapCtx, in.GetSilence())
	return out
}

func AnalysisResult_CallAnalysisMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalysisResult_CallAnalysisMetadataObservedState) *pb.AnalysisResult_CallAnalysisMetadata {
	if in == nil {
		return nil
	}
	out := &pb.AnalysisResult_CallAnalysisMetadata{}
	out.Annotations = direct.Slice_ToProto(mapCtx, in.Annotations, CallAnnotation_ToProto)
	out.Sentiments = direct.Slice_ToProto(mapCtx, in.Sentiments, ConversationLevelSentiment_ToProto)
	out.Silence = ConversationLevelSilence_ToProto(mapCtx, in.Silence)
	return out
}
