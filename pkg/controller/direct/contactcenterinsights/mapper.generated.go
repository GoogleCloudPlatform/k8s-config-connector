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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ContactcenterinsightsQaQuestionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QaQuestion) *krm.ContactcenterinsightsQaQuestionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsQaQuestionObservedState{}
	// MISSING: Name
	// MISSING: Abbreviation
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: QuestionBody
	// MISSING: AnswerInstructions
	// MISSING: AnswerChoices
	// MISSING: Tags
	// MISSING: Order
	// MISSING: Metrics
	// MISSING: TuningMetadata
	return out
}
func ContactcenterinsightsQaQuestionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsQaQuestionObservedState) *pb.QaQuestion {
	if in == nil {
		return nil
	}
	out := &pb.QaQuestion{}
	// MISSING: Name
	// MISSING: Abbreviation
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: QuestionBody
	// MISSING: AnswerInstructions
	// MISSING: AnswerChoices
	// MISSING: Tags
	// MISSING: Order
	// MISSING: Metrics
	// MISSING: TuningMetadata
	return out
}
func ContactcenterinsightsQaQuestionSpec_FromProto(mapCtx *direct.MapContext, in *pb.QaQuestion) *krm.ContactcenterinsightsQaQuestionSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsQaQuestionSpec{}
	// MISSING: Name
	// MISSING: Abbreviation
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: QuestionBody
	// MISSING: AnswerInstructions
	// MISSING: AnswerChoices
	// MISSING: Tags
	// MISSING: Order
	// MISSING: Metrics
	// MISSING: TuningMetadata
	return out
}
func ContactcenterinsightsQaQuestionSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsQaQuestionSpec) *pb.QaQuestion {
	if in == nil {
		return nil
	}
	out := &pb.QaQuestion{}
	// MISSING: Name
	// MISSING: Abbreviation
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: QuestionBody
	// MISSING: AnswerInstructions
	// MISSING: AnswerChoices
	// MISSING: Tags
	// MISSING: Order
	// MISSING: Metrics
	// MISSING: TuningMetadata
	return out
}
func QaQuestion_FromProto(mapCtx *direct.MapContext, in *pb.QaQuestion) *krm.QaQuestion {
	if in == nil {
		return nil
	}
	out := &krm.QaQuestion{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Abbreviation = direct.LazyPtr(in.GetAbbreviation())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.QuestionBody = direct.LazyPtr(in.GetQuestionBody())
	out.AnswerInstructions = direct.LazyPtr(in.GetAnswerInstructions())
	out.AnswerChoices = direct.Slice_FromProto(mapCtx, in.AnswerChoices, QaQuestion_AnswerChoice_FromProto)
	out.Tags = in.Tags
	out.Order = direct.LazyPtr(in.GetOrder())
	out.Metrics = QaQuestion_Metrics_FromProto(mapCtx, in.GetMetrics())
	out.TuningMetadata = QaQuestion_TuningMetadata_FromProto(mapCtx, in.GetTuningMetadata())
	return out
}
func QaQuestion_ToProto(mapCtx *direct.MapContext, in *krm.QaQuestion) *pb.QaQuestion {
	if in == nil {
		return nil
	}
	out := &pb.QaQuestion{}
	out.Name = direct.ValueOf(in.Name)
	out.Abbreviation = direct.ValueOf(in.Abbreviation)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.QuestionBody = direct.ValueOf(in.QuestionBody)
	out.AnswerInstructions = direct.ValueOf(in.AnswerInstructions)
	out.AnswerChoices = direct.Slice_ToProto(mapCtx, in.AnswerChoices, QaQuestion_AnswerChoice_ToProto)
	out.Tags = in.Tags
	out.Order = direct.ValueOf(in.Order)
	out.Metrics = QaQuestion_Metrics_ToProto(mapCtx, in.Metrics)
	out.TuningMetadata = QaQuestion_TuningMetadata_ToProto(mapCtx, in.TuningMetadata)
	return out
}
func QaQuestionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QaQuestion) *krm.QaQuestionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QaQuestionObservedState{}
	// MISSING: Name
	// MISSING: Abbreviation
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: QuestionBody
	// MISSING: AnswerInstructions
	// MISSING: AnswerChoices
	// MISSING: Tags
	// MISSING: Order
	out.Metrics = QaQuestion_MetricsObservedState_FromProto(mapCtx, in.GetMetrics())
	// MISSING: TuningMetadata
	return out
}
func QaQuestionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QaQuestionObservedState) *pb.QaQuestion {
	if in == nil {
		return nil
	}
	out := &pb.QaQuestion{}
	// MISSING: Name
	// MISSING: Abbreviation
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: QuestionBody
	// MISSING: AnswerInstructions
	// MISSING: AnswerChoices
	// MISSING: Tags
	// MISSING: Order
	out.Metrics = QaQuestion_MetricsObservedState_ToProto(mapCtx, in.Metrics)
	// MISSING: TuningMetadata
	return out
}
func QaQuestion_AnswerChoice_FromProto(mapCtx *direct.MapContext, in *pb.QaQuestion_AnswerChoice) *krm.QaQuestion_AnswerChoice {
	if in == nil {
		return nil
	}
	out := &krm.QaQuestion_AnswerChoice{}
	out.StrValue = direct.LazyPtr(in.GetStrValue())
	out.NumValue = direct.LazyPtr(in.GetNumValue())
	out.BoolValue = direct.LazyPtr(in.GetBoolValue())
	out.NaValue = direct.LazyPtr(in.GetNaValue())
	out.Key = direct.LazyPtr(in.GetKey())
	out.Score = in.Score
	return out
}
func QaQuestion_AnswerChoice_ToProto(mapCtx *direct.MapContext, in *krm.QaQuestion_AnswerChoice) *pb.QaQuestion_AnswerChoice {
	if in == nil {
		return nil
	}
	out := &pb.QaQuestion_AnswerChoice{}
	if oneof := QaQuestion_AnswerChoice_StrValue_ToProto(mapCtx, in.StrValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := QaQuestion_AnswerChoice_NumValue_ToProto(mapCtx, in.NumValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := QaQuestion_AnswerChoice_BoolValue_ToProto(mapCtx, in.BoolValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := QaQuestion_AnswerChoice_NaValue_ToProto(mapCtx, in.NaValue); oneof != nil {
		out.Value = oneof
	}
	out.Key = direct.ValueOf(in.Key)
	out.Score = in.Score
	return out
}
func QaQuestion_Metrics_FromProto(mapCtx *direct.MapContext, in *pb.QaQuestion_Metrics) *krm.QaQuestion_Metrics {
	if in == nil {
		return nil
	}
	out := &krm.QaQuestion_Metrics{}
	// MISSING: Accuracy
	return out
}
func QaQuestion_Metrics_ToProto(mapCtx *direct.MapContext, in *krm.QaQuestion_Metrics) *pb.QaQuestion_Metrics {
	if in == nil {
		return nil
	}
	out := &pb.QaQuestion_Metrics{}
	// MISSING: Accuracy
	return out
}
func QaQuestion_MetricsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QaQuestion_Metrics) *krm.QaQuestion_MetricsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QaQuestion_MetricsObservedState{}
	out.Accuracy = direct.LazyPtr(in.GetAccuracy())
	return out
}
func QaQuestion_MetricsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QaQuestion_MetricsObservedState) *pb.QaQuestion_Metrics {
	if in == nil {
		return nil
	}
	out := &pb.QaQuestion_Metrics{}
	out.Accuracy = direct.ValueOf(in.Accuracy)
	return out
}
