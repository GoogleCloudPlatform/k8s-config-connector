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

package automl

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/automl/apiv1beta1/automlpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/automl/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ArrayStats_FromProto(mapCtx *direct.MapContext, in *pb.ArrayStats) *krm.ArrayStats {
	if in == nil {
		return nil
	}
	out := &krm.ArrayStats{}
	out.MemberStats = DataStats_FromProto(mapCtx, in.GetMemberStats())
	return out
}
func ArrayStats_ToProto(mapCtx *direct.MapContext, in *krm.ArrayStats) *pb.ArrayStats {
	if in == nil {
		return nil
	}
	out := &pb.ArrayStats{}
	out.MemberStats = DataStats_ToProto(mapCtx, in.MemberStats)
	return out
}
func AutomlColumnSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ColumnSpec) *krm.AutomlColumnSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutomlColumnSpecObservedState{}
	// MISSING: Name
	// MISSING: DataType
	// MISSING: DisplayName
	// MISSING: DataStats
	// MISSING: TopCorrelatedColumns
	// MISSING: Etag
	return out
}
func AutomlColumnSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutomlColumnSpecObservedState) *pb.ColumnSpec {
	if in == nil {
		return nil
	}
	out := &pb.ColumnSpec{}
	// MISSING: Name
	// MISSING: DataType
	// MISSING: DisplayName
	// MISSING: DataStats
	// MISSING: TopCorrelatedColumns
	// MISSING: Etag
	return out
}
func AutomlColumnSpecSpec_FromProto(mapCtx *direct.MapContext, in *pb.ColumnSpec) *krm.AutomlColumnSpecSpec {
	if in == nil {
		return nil
	}
	out := &krm.AutomlColumnSpecSpec{}
	// MISSING: Name
	// MISSING: DataType
	// MISSING: DisplayName
	// MISSING: DataStats
	// MISSING: TopCorrelatedColumns
	// MISSING: Etag
	return out
}
func AutomlColumnSpecSpec_ToProto(mapCtx *direct.MapContext, in *krm.AutomlColumnSpecSpec) *pb.ColumnSpec {
	if in == nil {
		return nil
	}
	out := &pb.ColumnSpec{}
	// MISSING: Name
	// MISSING: DataType
	// MISSING: DisplayName
	// MISSING: DataStats
	// MISSING: TopCorrelatedColumns
	// MISSING: Etag
	return out
}
func CategoryStats_FromProto(mapCtx *direct.MapContext, in *pb.CategoryStats) *krm.CategoryStats {
	if in == nil {
		return nil
	}
	out := &krm.CategoryStats{}
	out.TopCategoryStats = direct.Slice_FromProto(mapCtx, in.TopCategoryStats, CategoryStats_SingleCategoryStats_FromProto)
	return out
}
func CategoryStats_ToProto(mapCtx *direct.MapContext, in *krm.CategoryStats) *pb.CategoryStats {
	if in == nil {
		return nil
	}
	out := &pb.CategoryStats{}
	out.TopCategoryStats = direct.Slice_ToProto(mapCtx, in.TopCategoryStats, CategoryStats_SingleCategoryStats_ToProto)
	return out
}
func CategoryStats_SingleCategoryStats_FromProto(mapCtx *direct.MapContext, in *pb.CategoryStats_SingleCategoryStats) *krm.CategoryStats_SingleCategoryStats {
	if in == nil {
		return nil
	}
	out := &krm.CategoryStats_SingleCategoryStats{}
	out.Value = direct.LazyPtr(in.GetValue())
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func CategoryStats_SingleCategoryStats_ToProto(mapCtx *direct.MapContext, in *krm.CategoryStats_SingleCategoryStats) *pb.CategoryStats_SingleCategoryStats {
	if in == nil {
		return nil
	}
	out := &pb.CategoryStats_SingleCategoryStats{}
	out.Value = direct.ValueOf(in.Value)
	out.Count = direct.ValueOf(in.Count)
	return out
}
func ColumnSpec_FromProto(mapCtx *direct.MapContext, in *pb.ColumnSpec) *krm.ColumnSpec {
	if in == nil {
		return nil
	}
	out := &krm.ColumnSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DataType = DataType_FromProto(mapCtx, in.GetDataType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.DataStats = DataStats_FromProto(mapCtx, in.GetDataStats())
	out.TopCorrelatedColumns = direct.Slice_FromProto(mapCtx, in.TopCorrelatedColumns, ColumnSpec_CorrelatedColumn_FromProto)
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func ColumnSpec_ToProto(mapCtx *direct.MapContext, in *krm.ColumnSpec) *pb.ColumnSpec {
	if in == nil {
		return nil
	}
	out := &pb.ColumnSpec{}
	out.Name = direct.ValueOf(in.Name)
	out.DataType = DataType_ToProto(mapCtx, in.DataType)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.DataStats = DataStats_ToProto(mapCtx, in.DataStats)
	out.TopCorrelatedColumns = direct.Slice_ToProto(mapCtx, in.TopCorrelatedColumns, ColumnSpec_CorrelatedColumn_ToProto)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func ColumnSpec_CorrelatedColumn_FromProto(mapCtx *direct.MapContext, in *pb.ColumnSpec_CorrelatedColumn) *krm.ColumnSpec_CorrelatedColumn {
	if in == nil {
		return nil
	}
	out := &krm.ColumnSpec_CorrelatedColumn{}
	out.ColumnSpecID = direct.LazyPtr(in.GetColumnSpecId())
	out.CorrelationStats = CorrelationStats_FromProto(mapCtx, in.GetCorrelationStats())
	return out
}
func ColumnSpec_CorrelatedColumn_ToProto(mapCtx *direct.MapContext, in *krm.ColumnSpec_CorrelatedColumn) *pb.ColumnSpec_CorrelatedColumn {
	if in == nil {
		return nil
	}
	out := &pb.ColumnSpec_CorrelatedColumn{}
	out.ColumnSpecId = direct.ValueOf(in.ColumnSpecID)
	out.CorrelationStats = CorrelationStats_ToProto(mapCtx, in.CorrelationStats)
	return out
}
func CorrelationStats_FromProto(mapCtx *direct.MapContext, in *pb.CorrelationStats) *krm.CorrelationStats {
	if in == nil {
		return nil
	}
	out := &krm.CorrelationStats{}
	out.CramersV = direct.LazyPtr(in.GetCramersV())
	return out
}
func CorrelationStats_ToProto(mapCtx *direct.MapContext, in *krm.CorrelationStats) *pb.CorrelationStats {
	if in == nil {
		return nil
	}
	out := &pb.CorrelationStats{}
	out.CramersV = direct.ValueOf(in.CramersV)
	return out
}
func DataStats_FromProto(mapCtx *direct.MapContext, in *pb.DataStats) *krm.DataStats {
	if in == nil {
		return nil
	}
	out := &krm.DataStats{}
	out.Float64Stats = Float64Stats_FromProto(mapCtx, in.GetFloat64Stats())
	out.StringStats = StringStats_FromProto(mapCtx, in.GetStringStats())
	out.TimestampStats = TimestampStats_FromProto(mapCtx, in.GetTimestampStats())
	out.ArrayStats = ArrayStats_FromProto(mapCtx, in.GetArrayStats())
	out.StructStats = StructStats_FromProto(mapCtx, in.GetStructStats())
	out.CategoryStats = CategoryStats_FromProto(mapCtx, in.GetCategoryStats())
	out.DistinctValueCount = direct.LazyPtr(in.GetDistinctValueCount())
	out.NullValueCount = direct.LazyPtr(in.GetNullValueCount())
	out.ValidValueCount = direct.LazyPtr(in.GetValidValueCount())
	return out
}
func DataStats_ToProto(mapCtx *direct.MapContext, in *krm.DataStats) *pb.DataStats {
	if in == nil {
		return nil
	}
	out := &pb.DataStats{}
	if oneof := Float64Stats_ToProto(mapCtx, in.Float64Stats); oneof != nil {
		out.Stats = &pb.DataStats_Float64Stats{Float64Stats: oneof}
	}
	if oneof := StringStats_ToProto(mapCtx, in.StringStats); oneof != nil {
		out.Stats = &pb.DataStats_StringStats{StringStats: oneof}
	}
	if oneof := TimestampStats_ToProto(mapCtx, in.TimestampStats); oneof != nil {
		out.Stats = &pb.DataStats_TimestampStats{TimestampStats: oneof}
	}
	if oneof := ArrayStats_ToProto(mapCtx, in.ArrayStats); oneof != nil {
		out.Stats = &pb.DataStats_ArrayStats{ArrayStats: oneof}
	}
	if oneof := StructStats_ToProto(mapCtx, in.StructStats); oneof != nil {
		out.Stats = &pb.DataStats_StructStats{StructStats: oneof}
	}
	if oneof := CategoryStats_ToProto(mapCtx, in.CategoryStats); oneof != nil {
		out.Stats = &pb.DataStats_CategoryStats{CategoryStats: oneof}
	}
	out.DistinctValueCount = direct.ValueOf(in.DistinctValueCount)
	out.NullValueCount = direct.ValueOf(in.NullValueCount)
	out.ValidValueCount = direct.ValueOf(in.ValidValueCount)
	return out
}
func DataType_FromProto(mapCtx *direct.MapContext, in *pb.DataType) *krm.DataType {
	if in == nil {
		return nil
	}
	out := &krm.DataType{}
	out.ListElementType = DataType_FromProto(mapCtx, in.GetListElementType())
	out.StructType = StructType_FromProto(mapCtx, in.GetStructType())
	out.TimeFormat = direct.LazyPtr(in.GetTimeFormat())
	out.TypeCode = direct.Enum_FromProto(mapCtx, in.GetTypeCode())
	out.Nullable = direct.LazyPtr(in.GetNullable())
	return out
}
func DataType_ToProto(mapCtx *direct.MapContext, in *krm.DataType) *pb.DataType {
	if in == nil {
		return nil
	}
	out := &pb.DataType{}
	if oneof := DataType_ToProto(mapCtx, in.ListElementType); oneof != nil {
		out.Details = &pb.DataType_ListElementType{ListElementType: oneof}
	}
	if oneof := StructType_ToProto(mapCtx, in.StructType); oneof != nil {
		out.Details = &pb.DataType_StructType{StructType: oneof}
	}
	if oneof := DataType_TimeFormat_ToProto(mapCtx, in.TimeFormat); oneof != nil {
		out.Details = oneof
	}
	out.TypeCode = direct.Enum_ToProto[pb.TypeCode](mapCtx, in.TypeCode)
	out.Nullable = direct.ValueOf(in.Nullable)
	return out
}
func Float64Stats_FromProto(mapCtx *direct.MapContext, in *pb.Float64Stats) *krm.Float64Stats {
	if in == nil {
		return nil
	}
	out := &krm.Float64Stats{}
	out.Mean = direct.LazyPtr(in.GetMean())
	out.StandardDeviation = direct.LazyPtr(in.GetStandardDeviation())
	out.Quantiles = in.Quantiles
	out.HistogramBuckets = direct.Slice_FromProto(mapCtx, in.HistogramBuckets, Float64Stats_HistogramBucket_FromProto)
	return out
}
func Float64Stats_ToProto(mapCtx *direct.MapContext, in *krm.Float64Stats) *pb.Float64Stats {
	if in == nil {
		return nil
	}
	out := &pb.Float64Stats{}
	out.Mean = direct.ValueOf(in.Mean)
	out.StandardDeviation = direct.ValueOf(in.StandardDeviation)
	out.Quantiles = in.Quantiles
	out.HistogramBuckets = direct.Slice_ToProto(mapCtx, in.HistogramBuckets, Float64Stats_HistogramBucket_ToProto)
	return out
}
func Float64Stats_HistogramBucket_FromProto(mapCtx *direct.MapContext, in *pb.Float64Stats_HistogramBucket) *krm.Float64Stats_HistogramBucket {
	if in == nil {
		return nil
	}
	out := &krm.Float64Stats_HistogramBucket{}
	out.Min = direct.LazyPtr(in.GetMin())
	out.Max = direct.LazyPtr(in.GetMax())
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func Float64Stats_HistogramBucket_ToProto(mapCtx *direct.MapContext, in *krm.Float64Stats_HistogramBucket) *pb.Float64Stats_HistogramBucket {
	if in == nil {
		return nil
	}
	out := &pb.Float64Stats_HistogramBucket{}
	out.Min = direct.ValueOf(in.Min)
	out.Max = direct.ValueOf(in.Max)
	out.Count = direct.ValueOf(in.Count)
	return out
}
func StringStats_FromProto(mapCtx *direct.MapContext, in *pb.StringStats) *krm.StringStats {
	if in == nil {
		return nil
	}
	out := &krm.StringStats{}
	out.TopUnigramStats = direct.Slice_FromProto(mapCtx, in.TopUnigramStats, StringStats_UnigramStats_FromProto)
	return out
}
func StringStats_ToProto(mapCtx *direct.MapContext, in *krm.StringStats) *pb.StringStats {
	if in == nil {
		return nil
	}
	out := &pb.StringStats{}
	out.TopUnigramStats = direct.Slice_ToProto(mapCtx, in.TopUnigramStats, StringStats_UnigramStats_ToProto)
	return out
}
func StringStats_UnigramStats_FromProto(mapCtx *direct.MapContext, in *pb.StringStats_UnigramStats) *krm.StringStats_UnigramStats {
	if in == nil {
		return nil
	}
	out := &krm.StringStats_UnigramStats{}
	out.Value = direct.LazyPtr(in.GetValue())
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func StringStats_UnigramStats_ToProto(mapCtx *direct.MapContext, in *krm.StringStats_UnigramStats) *pb.StringStats_UnigramStats {
	if in == nil {
		return nil
	}
	out := &pb.StringStats_UnigramStats{}
	out.Value = direct.ValueOf(in.Value)
	out.Count = direct.ValueOf(in.Count)
	return out
}
func StructStats_FromProto(mapCtx *direct.MapContext, in *pb.StructStats) *krm.StructStats {
	if in == nil {
		return nil
	}
	out := &krm.StructStats{}
	// MISSING: FieldStats
	return out
}
func StructStats_ToProto(mapCtx *direct.MapContext, in *krm.StructStats) *pb.StructStats {
	if in == nil {
		return nil
	}
	out := &pb.StructStats{}
	// MISSING: FieldStats
	return out
}
func StructType_FromProto(mapCtx *direct.MapContext, in *pb.StructType) *krm.StructType {
	if in == nil {
		return nil
	}
	out := &krm.StructType{}
	// MISSING: Fields
	return out
}
func StructType_ToProto(mapCtx *direct.MapContext, in *krm.StructType) *pb.StructType {
	if in == nil {
		return nil
	}
	out := &pb.StructType{}
	// MISSING: Fields
	return out
}
func TimestampStats_FromProto(mapCtx *direct.MapContext, in *pb.TimestampStats) *krm.TimestampStats {
	if in == nil {
		return nil
	}
	out := &krm.TimestampStats{}
	// MISSING: GranularStats
	return out
}
func TimestampStats_ToProto(mapCtx *direct.MapContext, in *krm.TimestampStats) *pb.TimestampStats {
	if in == nil {
		return nil
	}
	out := &pb.TimestampStats{}
	// MISSING: GranularStats
	return out
}
func TimestampStats_GranularStats_FromProto(mapCtx *direct.MapContext, in *pb.TimestampStats_GranularStats) *krm.TimestampStats_GranularStats {
	if in == nil {
		return nil
	}
	out := &krm.TimestampStats_GranularStats{}
	// MISSING: Buckets
	return out
}
func TimestampStats_GranularStats_ToProto(mapCtx *direct.MapContext, in *krm.TimestampStats_GranularStats) *pb.TimestampStats_GranularStats {
	if in == nil {
		return nil
	}
	out := &pb.TimestampStats_GranularStats{}
	// MISSING: Buckets
	return out
}
