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

package contactcenterinsights

import (
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ContactcenterinsightsQaScorecardResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecardResult) *krm.ContactcenterinsightsQaScorecardResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsQaScorecardResultObservedState{}
	// MISSING: Name
	// MISSING: QaScorecardRevision
	// MISSING: Conversation
	// MISSING: CreateTime
	// MISSING: AgentID
	// MISSING: QaAnswers
	// MISSING: Score
	// MISSING: PotentialScore
	// MISSING: NormalizedScore
	// MISSING: QaTagResults
	// MISSING: ScoreSources
	return out
}
func ContactcenterinsightsQaScorecardResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsQaScorecardResultObservedState) *pb.QaScorecardResult {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecardResult{}
	// MISSING: Name
	// MISSING: QaScorecardRevision
	// MISSING: Conversation
	// MISSING: CreateTime
	// MISSING: AgentID
	// MISSING: QaAnswers
	// MISSING: Score
	// MISSING: PotentialScore
	// MISSING: NormalizedScore
	// MISSING: QaTagResults
	// MISSING: ScoreSources
	return out
}
func ContactcenterinsightsQaScorecardResultSpec_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecardResult) *krm.ContactcenterinsightsQaScorecardResultSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsQaScorecardResultSpec{}
	// MISSING: Name
	// MISSING: QaScorecardRevision
	// MISSING: Conversation
	// MISSING: CreateTime
	// MISSING: AgentID
	// MISSING: QaAnswers
	// MISSING: Score
	// MISSING: PotentialScore
	// MISSING: NormalizedScore
	// MISSING: QaTagResults
	// MISSING: ScoreSources
	return out
}
func ContactcenterinsightsQaScorecardResultSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsQaScorecardResultSpec) *pb.QaScorecardResult {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecardResult{}
	// MISSING: Name
	// MISSING: QaScorecardRevision
	// MISSING: Conversation
	// MISSING: CreateTime
	// MISSING: AgentID
	// MISSING: QaAnswers
	// MISSING: Score
	// MISSING: PotentialScore
	// MISSING: NormalizedScore
	// MISSING: QaTagResults
	// MISSING: ScoreSources
	return out
}
func QaAnswer_FromProto(mapCtx *direct.MapContext, in *pb.QaAnswer) *krm.QaAnswer {
	if in == nil {
		return nil
	}
	out := &krm.QaAnswer{}
	out.QaQuestion = direct.LazyPtr(in.GetQaQuestion())
	out.Conversation = direct.LazyPtr(in.GetConversation())
	out.QuestionBody = direct.LazyPtr(in.GetQuestionBody())
	out.AnswerValue = QaAnswer_AnswerValue_FromProto(mapCtx, in.GetAnswerValue())
	out.Tags = in.Tags
	out.AnswerSources = direct.Slice_FromProto(mapCtx, in.AnswerSources, QaAnswer_AnswerSource_FromProto)
	return out
}
func QaAnswer_ToProto(mapCtx *direct.MapContext, in *krm.QaAnswer) *pb.QaAnswer {
	if in == nil {
		return nil
	}
	out := &pb.QaAnswer{}
	out.QaQuestion = direct.ValueOf(in.QaQuestion)
	out.Conversation = direct.ValueOf(in.Conversation)
	out.QuestionBody = direct.ValueOf(in.QuestionBody)
	out.AnswerValue = QaAnswer_AnswerValue_ToProto(mapCtx, in.AnswerValue)
	out.Tags = in.Tags
	out.AnswerSources = direct.Slice_ToProto(mapCtx, in.AnswerSources, QaAnswer_AnswerSource_ToProto)
	return out
}
func QaAnswerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QaAnswer) *krm.QaAnswerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QaAnswerObservedState{}
	// MISSING: QaQuestion
	// MISSING: Conversation
	// MISSING: QuestionBody
	out.AnswerValue = QaAnswer_AnswerValueObservedState_FromProto(mapCtx, in.GetAnswerValue())
	// MISSING: Tags
	// MISSING: AnswerSources
	return out
}
func QaAnswerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QaAnswerObservedState) *pb.QaAnswer {
	if in == nil {
		return nil
	}
	out := &pb.QaAnswer{}
	// MISSING: QaQuestion
	// MISSING: Conversation
	// MISSING: QuestionBody
	out.AnswerValue = QaAnswer_AnswerValueObservedState_ToProto(mapCtx, in.AnswerValue)
	// MISSING: Tags
	// MISSING: AnswerSources
	return out
}
func QaAnswer_AnswerSource_FromProto(mapCtx *direct.MapContext, in *pb.QaAnswer_AnswerSource) *krm.QaAnswer_AnswerSource {
	if in == nil {
		return nil
	}
	out := &krm.QaAnswer_AnswerSource{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.AnswerValue = QaAnswer_AnswerValue_FromProto(mapCtx, in.GetAnswerValue())
	return out
}
func QaAnswer_AnswerSource_ToProto(mapCtx *direct.MapContext, in *krm.QaAnswer_AnswerSource) *pb.QaAnswer_AnswerSource {
	if in == nil {
		return nil
	}
	out := &pb.QaAnswer_AnswerSource{}
	out.SourceType = direct.Enum_ToProto[pb.QaAnswer_AnswerSource_SourceType](mapCtx, in.SourceType)
	out.AnswerValue = QaAnswer_AnswerValue_ToProto(mapCtx, in.AnswerValue)
	return out
}
func QaAnswer_AnswerValue_FromProto(mapCtx *direct.MapContext, in *pb.QaAnswer_AnswerValue) *krm.QaAnswer_AnswerValue {
	if in == nil {
		return nil
	}
	out := &krm.QaAnswer_AnswerValue{}
	out.StrValue = direct.LazyPtr(in.GetStrValue())
	out.NumValue = direct.LazyPtr(in.GetNumValue())
	out.BoolValue = direct.LazyPtr(in.GetBoolValue())
	out.NaValue = direct.LazyPtr(in.GetNaValue())
	out.Key = direct.LazyPtr(in.GetKey())
	// MISSING: Score
	// MISSING: PotentialScore
	// MISSING: NormalizedScore
	return out
}
func QaAnswer_AnswerValue_ToProto(mapCtx *direct.MapContext, in *krm.QaAnswer_AnswerValue) *pb.QaAnswer_AnswerValue {
	if in == nil {
		return nil
	}
	out := &pb.QaAnswer_AnswerValue{}
	if oneof := QaAnswer_AnswerValue_StrValue_ToProto(mapCtx, in.StrValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := QaAnswer_AnswerValue_NumValue_ToProto(mapCtx, in.NumValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := QaAnswer_AnswerValue_BoolValue_ToProto(mapCtx, in.BoolValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := QaAnswer_AnswerValue_NaValue_ToProto(mapCtx, in.NaValue); oneof != nil {
		out.Value = oneof
	}
	out.Key = direct.ValueOf(in.Key)
	// MISSING: Score
	// MISSING: PotentialScore
	// MISSING: NormalizedScore
	return out
}
func QaAnswer_AnswerValueObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QaAnswer_AnswerValue) *krm.QaAnswer_AnswerValueObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QaAnswer_AnswerValueObservedState{}
	// MISSING: StrValue
	// MISSING: NumValue
	// MISSING: BoolValue
	// MISSING: NaValue
	// MISSING: Key
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	return out
}
func QaAnswer_AnswerValueObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QaAnswer_AnswerValueObservedState) *pb.QaAnswer_AnswerValue {
	if in == nil {
		return nil
	}
	out := &pb.QaAnswer_AnswerValue{}
	// MISSING: StrValue
	// MISSING: NumValue
	// MISSING: BoolValue
	// MISSING: NaValue
	// MISSING: Key
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	return out
}
func QaScorecardResult_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecardResult) *krm.QaScorecardResult {
	if in == nil {
		return nil
	}
	out := &krm.QaScorecardResult{}
	out.Name = direct.LazyPtr(in.GetName())
	out.QaScorecardRevision = direct.LazyPtr(in.GetQaScorecardRevision())
	out.Conversation = direct.LazyPtr(in.GetConversation())
	// MISSING: CreateTime
	out.AgentID = direct.LazyPtr(in.GetAgentId())
	out.QaAnswers = direct.Slice_FromProto(mapCtx, in.QaAnswers, QaAnswer_FromProto)
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	out.QaTagResults = direct.Slice_FromProto(mapCtx, in.QaTagResults, QaScorecardResult_QaTagResult_FromProto)
	out.ScoreSources = direct.Slice_FromProto(mapCtx, in.ScoreSources, QaScorecardResult_ScoreSource_FromProto)
	return out
}
func QaScorecardResult_ToProto(mapCtx *direct.MapContext, in *krm.QaScorecardResult) *pb.QaScorecardResult {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecardResult{}
	out.Name = direct.ValueOf(in.Name)
	out.QaScorecardRevision = direct.ValueOf(in.QaScorecardRevision)
	out.Conversation = direct.ValueOf(in.Conversation)
	// MISSING: CreateTime
	out.AgentId = direct.ValueOf(in.AgentID)
	out.QaAnswers = direct.Slice_ToProto(mapCtx, in.QaAnswers, QaAnswer_ToProto)
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	out.QaTagResults = direct.Slice_ToProto(mapCtx, in.QaTagResults, QaScorecardResult_QaTagResult_ToProto)
	out.ScoreSources = direct.Slice_ToProto(mapCtx, in.ScoreSources, QaScorecardResult_ScoreSource_ToProto)
	return out
}
func QaScorecardResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecardResult) *krm.QaScorecardResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QaScorecardResultObservedState{}
	// MISSING: Name
	// MISSING: QaScorecardRevision
	// MISSING: Conversation
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: AgentID
	out.QaAnswers = direct.Slice_FromProto(mapCtx, in.QaAnswers, QaAnswerObservedState_FromProto)
	// MISSING: Score
	// MISSING: PotentialScore
	// MISSING: NormalizedScore
	// MISSING: QaTagResults
	// MISSING: ScoreSources
	return out
}
func QaScorecardResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QaScorecardResultObservedState) *pb.QaScorecardResult {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecardResult{}
	// MISSING: Name
	// MISSING: QaScorecardRevision
	// MISSING: Conversation
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: AgentID
	out.QaAnswers = direct.Slice_ToProto(mapCtx, in.QaAnswers, QaAnswerObservedState_ToProto)
	// MISSING: Score
	// MISSING: PotentialScore
	// MISSING: NormalizedScore
	// MISSING: QaTagResults
	// MISSING: ScoreSources
	return out
}
func QaScorecardResult_QaTagResult_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecardResult_QaTagResult) *krm.QaScorecardResult_QaTagResult {
	if in == nil {
		return nil
	}
	out := &krm.QaScorecardResult_QaTagResult{}
	out.Tag = direct.LazyPtr(in.GetTag())
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	return out
}
func QaScorecardResult_QaTagResult_ToProto(mapCtx *direct.MapContext, in *krm.QaScorecardResult_QaTagResult) *pb.QaScorecardResult_QaTagResult {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecardResult_QaTagResult{}
	out.Tag = direct.ValueOf(in.Tag)
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	return out
}
func QaScorecardResult_ScoreSource_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecardResult_ScoreSource) *krm.QaScorecardResult_ScoreSource {
	if in == nil {
		return nil
	}
	out := &krm.QaScorecardResult_ScoreSource{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	out.QaTagResults = direct.Slice_FromProto(mapCtx, in.QaTagResults, QaScorecardResult_QaTagResult_FromProto)
	return out
}
func QaScorecardResult_ScoreSource_ToProto(mapCtx *direct.MapContext, in *krm.QaScorecardResult_ScoreSource) *pb.QaScorecardResult_ScoreSource {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecardResult_ScoreSource{}
	out.SourceType = direct.Enum_ToProto[pb.QaScorecardResult_ScoreSource_SourceType](mapCtx, in.SourceType)
	out.Score = in.Score
	out.PotentialScore = in.PotentialScore
	out.NormalizedScore = in.NormalizedScore
	out.QaTagResults = direct.Slice_ToProto(mapCtx, in.QaTagResults, QaScorecardResult_QaTagResult_ToProto)
	return out
}
