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

package discoveryengine

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DiscoveryengineEvaluationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation) *krm.DiscoveryengineEvaluationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineEvaluationObservedState{}
	// MISSING: Name
	// MISSING: EvaluationSpec
	// MISSING: QualityMetrics
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: ErrorSamples
	return out
}
func DiscoveryengineEvaluationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineEvaluationObservedState) *pb.Evaluation {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation{}
	// MISSING: Name
	// MISSING: EvaluationSpec
	// MISSING: QualityMetrics
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: ErrorSamples
	return out
}
func DiscoveryengineEvaluationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation) *krm.DiscoveryengineEvaluationSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineEvaluationSpec{}
	// MISSING: Name
	// MISSING: EvaluationSpec
	// MISSING: QualityMetrics
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: ErrorSamples
	return out
}
func DiscoveryengineEvaluationSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineEvaluationSpec) *pb.Evaluation {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation{}
	// MISSING: Name
	// MISSING: EvaluationSpec
	// MISSING: QualityMetrics
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: ErrorSamples
	return out
}
func Evaluation_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation) *krm.Evaluation {
	if in == nil {
		return nil
	}
	out := &krm.Evaluation{}
	out.Name = direct.LazyPtr(in.GetName())
	out.EvaluationSpec = Evaluation_EvaluationSpec_FromProto(mapCtx, in.GetEvaluationSpec())
	// MISSING: QualityMetrics
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: ErrorSamples
	return out
}
func Evaluation_ToProto(mapCtx *direct.MapContext, in *krm.Evaluation) *pb.Evaluation {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation{}
	out.Name = direct.ValueOf(in.Name)
	out.EvaluationSpec = Evaluation_EvaluationSpec_ToProto(mapCtx, in.EvaluationSpec)
	// MISSING: QualityMetrics
	// MISSING: State
	// MISSING: Error
	// MISSING: CreateTime
	// MISSING: EndTime
	// MISSING: ErrorSamples
	return out
}
func EvaluationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation) *krm.EvaluationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EvaluationObservedState{}
	// MISSING: Name
	// MISSING: EvaluationSpec
	out.QualityMetrics = QualityMetrics_FromProto(mapCtx, in.GetQualityMetrics())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.ErrorSamples = direct.Slice_FromProto(mapCtx, in.ErrorSamples, Status_FromProto)
	return out
}
func EvaluationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EvaluationObservedState) *pb.Evaluation {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation{}
	// MISSING: Name
	// MISSING: EvaluationSpec
	out.QualityMetrics = QualityMetrics_ToProto(mapCtx, in.QualityMetrics)
	out.State = direct.Enum_ToProto[pb.Evaluation_State](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.ErrorSamples = direct.Slice_ToProto(mapCtx, in.ErrorSamples, Status_ToProto)
	return out
}
func Evaluation_EvaluationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation_EvaluationSpec) *krm.Evaluation_EvaluationSpec {
	if in == nil {
		return nil
	}
	out := &krm.Evaluation_EvaluationSpec{}
	out.SearchRequest = SearchRequest_FromProto(mapCtx, in.GetSearchRequest())
	out.QuerySetSpec = Evaluation_EvaluationSpec_QuerySetSpec_FromProto(mapCtx, in.GetQuerySetSpec())
	return out
}
func Evaluation_EvaluationSpec_ToProto(mapCtx *direct.MapContext, in *krm.Evaluation_EvaluationSpec) *pb.Evaluation_EvaluationSpec {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation_EvaluationSpec{}
	if oneof := SearchRequest_ToProto(mapCtx, in.SearchRequest); oneof != nil {
		out.SearchSpec = &pb.Evaluation_EvaluationSpec_SearchRequest{SearchRequest: oneof}
	}
	out.QuerySetSpec = Evaluation_EvaluationSpec_QuerySetSpec_ToProto(mapCtx, in.QuerySetSpec)
	return out
}
func Evaluation_EvaluationSpec_QuerySetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Evaluation_EvaluationSpec_QuerySetSpec) *krm.Evaluation_EvaluationSpec_QuerySetSpec {
	if in == nil {
		return nil
	}
	out := &krm.Evaluation_EvaluationSpec_QuerySetSpec{}
	out.SampleQuerySet = direct.LazyPtr(in.GetSampleQuerySet())
	return out
}
func Evaluation_EvaluationSpec_QuerySetSpec_ToProto(mapCtx *direct.MapContext, in *krm.Evaluation_EvaluationSpec_QuerySetSpec) *pb.Evaluation_EvaluationSpec_QuerySetSpec {
	if in == nil {
		return nil
	}
	out := &pb.Evaluation_EvaluationSpec_QuerySetSpec{}
	out.SampleQuerySet = direct.ValueOf(in.SampleQuerySet)
	return out
}
func Interval_FromProto(mapCtx *direct.MapContext, in *pb.Interval) *krm.Interval {
	if in == nil {
		return nil
	}
	out := &krm.Interval{}
	out.Minimum = direct.LazyPtr(in.GetMinimum())
	out.ExclusiveMinimum = direct.LazyPtr(in.GetExclusiveMinimum())
	out.Maximum = direct.LazyPtr(in.GetMaximum())
	out.ExclusiveMaximum = direct.LazyPtr(in.GetExclusiveMaximum())
	return out
}
func Interval_ToProto(mapCtx *direct.MapContext, in *krm.Interval) *pb.Interval {
	if in == nil {
		return nil
	}
	out := &pb.Interval{}
	if oneof := Interval_Minimum_ToProto(mapCtx, in.Minimum); oneof != nil {
		out.Min = oneof
	}
	if oneof := Interval_ExclusiveMinimum_ToProto(mapCtx, in.ExclusiveMinimum); oneof != nil {
		out.Min = oneof
	}
	if oneof := Interval_Maximum_ToProto(mapCtx, in.Maximum); oneof != nil {
		out.Max = oneof
	}
	if oneof := Interval_ExclusiveMaximum_ToProto(mapCtx, in.ExclusiveMaximum); oneof != nil {
		out.Max = oneof
	}
	return out
}
func QualityMetrics_FromProto(mapCtx *direct.MapContext, in *pb.QualityMetrics) *krm.QualityMetrics {
	if in == nil {
		return nil
	}
	out := &krm.QualityMetrics{}
	out.DocRecall = QualityMetrics_TopkMetrics_FromProto(mapCtx, in.GetDocRecall())
	out.DocPrecision = QualityMetrics_TopkMetrics_FromProto(mapCtx, in.GetDocPrecision())
	out.DocNdcg = QualityMetrics_TopkMetrics_FromProto(mapCtx, in.GetDocNdcg())
	out.PageRecall = QualityMetrics_TopkMetrics_FromProto(mapCtx, in.GetPageRecall())
	out.PageNdcg = QualityMetrics_TopkMetrics_FromProto(mapCtx, in.GetPageNdcg())
	return out
}
func QualityMetrics_ToProto(mapCtx *direct.MapContext, in *krm.QualityMetrics) *pb.QualityMetrics {
	if in == nil {
		return nil
	}
	out := &pb.QualityMetrics{}
	out.DocRecall = QualityMetrics_TopkMetrics_ToProto(mapCtx, in.DocRecall)
	out.DocPrecision = QualityMetrics_TopkMetrics_ToProto(mapCtx, in.DocPrecision)
	out.DocNdcg = QualityMetrics_TopkMetrics_ToProto(mapCtx, in.DocNdcg)
	out.PageRecall = QualityMetrics_TopkMetrics_ToProto(mapCtx, in.PageRecall)
	out.PageNdcg = QualityMetrics_TopkMetrics_ToProto(mapCtx, in.PageNdcg)
	return out
}
func QualityMetrics_TopkMetrics_FromProto(mapCtx *direct.MapContext, in *pb.QualityMetrics_TopkMetrics) *krm.QualityMetrics_TopkMetrics {
	if in == nil {
		return nil
	}
	out := &krm.QualityMetrics_TopkMetrics{}
	out.Top1 = direct.LazyPtr(in.GetTop1())
	out.Top3 = direct.LazyPtr(in.GetTop3())
	out.Top5 = direct.LazyPtr(in.GetTop5())
	out.Top10 = direct.LazyPtr(in.GetTop10())
	return out
}
func QualityMetrics_TopkMetrics_ToProto(mapCtx *direct.MapContext, in *krm.QualityMetrics_TopkMetrics) *pb.QualityMetrics_TopkMetrics {
	if in == nil {
		return nil
	}
	out := &pb.QualityMetrics_TopkMetrics{}
	out.Top1 = direct.ValueOf(in.Top1)
	out.Top3 = direct.ValueOf(in.Top3)
	out.Top5 = direct.ValueOf(in.Top5)
	out.Top10 = direct.ValueOf(in.Top10)
	return out
}
func UserInfo_FromProto(mapCtx *direct.MapContext, in *pb.UserInfo) *krm.UserInfo {
	if in == nil {
		return nil
	}
	out := &krm.UserInfo{}
	out.UserID = direct.LazyPtr(in.GetUserId())
	out.UserAgent = direct.LazyPtr(in.GetUserAgent())
	return out
}
func UserInfo_ToProto(mapCtx *direct.MapContext, in *krm.UserInfo) *pb.UserInfo {
	if in == nil {
		return nil
	}
	out := &pb.UserInfo{}
	out.UserId = direct.ValueOf(in.UserID)
	out.UserAgent = direct.ValueOf(in.UserAgent)
	return out
}
