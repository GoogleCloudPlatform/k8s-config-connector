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

package datacatalog

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BigQueryDateShardedSpec_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDateShardedSpec) *krm.BigQueryDateShardedSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDateShardedSpec{}
	// MISSING: Dataset
	// MISSING: TablePrefix
	// MISSING: ShardCount
	return out
}
func BigQueryDateShardedSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDateShardedSpec) *pb.BigQueryDateShardedSpec {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDateShardedSpec{}
	// MISSING: Dataset
	// MISSING: TablePrefix
	// MISSING: ShardCount
	return out
}
func BigQueryDateShardedSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDateShardedSpec) *krm.BigQueryDateShardedSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDateShardedSpecObservedState{}
	out.Dataset = direct.LazyPtr(in.GetDataset())
	out.TablePrefix = direct.LazyPtr(in.GetTablePrefix())
	out.ShardCount = direct.LazyPtr(in.GetShardCount())
	return out
}
func BigQueryDateShardedSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDateShardedSpecObservedState) *pb.BigQueryDateShardedSpec {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDateShardedSpec{}
	out.Dataset = direct.ValueOf(in.Dataset)
	out.TablePrefix = direct.ValueOf(in.TablePrefix)
	out.ShardCount = direct.ValueOf(in.ShardCount)
	return out
}
func BigQueryTableSpec_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryTableSpec) *krm.BigQueryTableSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryTableSpec{}
	// MISSING: TableSourceType
	out.ViewSpec = ViewSpec_FromProto(mapCtx, in.GetViewSpec())
	out.TableSpec = TableSpec_FromProto(mapCtx, in.GetTableSpec())
	return out
}
func BigQueryTableSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryTableSpec) *pb.BigQueryTableSpec {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryTableSpec{}
	// MISSING: TableSourceType
	if oneof := ViewSpec_ToProto(mapCtx, in.ViewSpec); oneof != nil {
		out.TypeSpec = &pb.BigQueryTableSpec_ViewSpec{ViewSpec: oneof}
	}
	if oneof := TableSpec_ToProto(mapCtx, in.TableSpec); oneof != nil {
		out.TypeSpec = &pb.BigQueryTableSpec_TableSpec{TableSpec: oneof}
	}
	return out
}
func BigQueryTableSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryTableSpec) *krm.BigQueryTableSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryTableSpecObservedState{}
	out.TableSourceType = direct.Enum_FromProto(mapCtx, in.GetTableSourceType())
	out.ViewSpec = ViewSpecObservedState_FromProto(mapCtx, in.GetViewSpec())
	out.TableSpec = TableSpecObservedState_FromProto(mapCtx, in.GetTableSpec())
	return out
}
func BigQueryTableSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryTableSpecObservedState) *pb.BigQueryTableSpec {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryTableSpec{}
	out.TableSourceType = direct.Enum_ToProto[pb.TableSourceType](mapCtx, in.TableSourceType)
	if oneof := ViewSpecObservedState_ToProto(mapCtx, in.ViewSpec); oneof != nil {
		out.TypeSpec = &pb.BigQueryTableSpec_ViewSpec{ViewSpec: oneof}
	}
	if oneof := TableSpecObservedState_ToProto(mapCtx, in.TableSpec); oneof != nil {
		out.TypeSpec = &pb.BigQueryTableSpec_TableSpec{TableSpec: oneof}
	}
	return out
}
func ColumnSchema_FromProto(mapCtx *direct.MapContext, in *pb.ColumnSchema) *krm.ColumnSchema {
	if in == nil {
		return nil
	}
	out := &krm.ColumnSchema{}
	out.Column = direct.LazyPtr(in.GetColumn())
	out.Type = direct.LazyPtr(in.GetType())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Mode = direct.LazyPtr(in.GetMode())
	out.Subcolumns = direct.Slice_FromProto(mapCtx, in.Subcolumns, ColumnSchema_FromProto)
	return out
}
func ColumnSchema_ToProto(mapCtx *direct.MapContext, in *krm.ColumnSchema) *pb.ColumnSchema {
	if in == nil {
		return nil
	}
	out := &pb.ColumnSchema{}
	out.Column = direct.ValueOf(in.Column)
	out.Type = direct.ValueOf(in.Type)
	out.Description = direct.ValueOf(in.Description)
	out.Mode = direct.ValueOf(in.Mode)
	out.Subcolumns = direct.Slice_ToProto(mapCtx, in.Subcolumns, ColumnSchema_ToProto)
	return out
}
func Entry_FromProto(mapCtx *direct.MapContext, in *pb.Entry) *krm.Entry {
	if in == nil {
		return nil
	}
	out := &krm.Entry{}
	// MISSING: Name
	out.LinkedResource = direct.LazyPtr(in.GetLinkedResource())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.UserSpecifiedType = direct.LazyPtr(in.GetUserSpecifiedType())
	// MISSING: IntegratedSystem
	out.UserSpecifiedSystem = direct.LazyPtr(in.GetUserSpecifiedSystem())
	out.GcsFilesetSpec = GcsFilesetSpec_FromProto(mapCtx, in.GetGcsFilesetSpec())
	out.BigqueryTableSpec = BigQueryTableSpec_FromProto(mapCtx, in.GetBigqueryTableSpec())
	out.BigqueryDateShardedSpec = BigQueryDateShardedSpec_FromProto(mapCtx, in.GetBigqueryDateShardedSpec())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Schema = Schema_FromProto(mapCtx, in.GetSchema())
	// MISSING: SourceSystemTimestamps
	// MISSING: UsageSignal
	return out
}
func Entry_ToProto(mapCtx *direct.MapContext, in *krm.Entry) *pb.Entry {
	if in == nil {
		return nil
	}
	out := &pb.Entry{}
	// MISSING: Name
	out.LinkedResource = direct.ValueOf(in.LinkedResource)
	if oneof := Entry_Type_ToProto(mapCtx, in.Type); oneof != nil {
		out.EntryType = oneof
	}
	if oneof := Entry_UserSpecifiedType_ToProto(mapCtx, in.UserSpecifiedType); oneof != nil {
		out.EntryType = oneof
	}
	// MISSING: IntegratedSystem
	if oneof := Entry_UserSpecifiedSystem_ToProto(mapCtx, in.UserSpecifiedSystem); oneof != nil {
		out.System = oneof
	}
	if oneof := GcsFilesetSpec_ToProto(mapCtx, in.GcsFilesetSpec); oneof != nil {
		out.TypeSpec = &pb.Entry_GcsFilesetSpec{GcsFilesetSpec: oneof}
	}
	if oneof := BigQueryTableSpec_ToProto(mapCtx, in.BigqueryTableSpec); oneof != nil {
		out.TypeSpec = &pb.Entry_BigqueryTableSpec{BigqueryTableSpec: oneof}
	}
	if oneof := BigQueryDateShardedSpec_ToProto(mapCtx, in.BigqueryDateShardedSpec); oneof != nil {
		out.TypeSpec = &pb.Entry_BigqueryDateShardedSpec{BigqueryDateShardedSpec: oneof}
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Schema = Schema_ToProto(mapCtx, in.Schema)
	// MISSING: SourceSystemTimestamps
	// MISSING: UsageSignal
	return out
}
func EntryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entry) *krm.EntryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EntryObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: LinkedResource
	// MISSING: Type
	// MISSING: UserSpecifiedType
	out.IntegratedSystem = direct.Enum_FromProto(mapCtx, in.GetIntegratedSystem())
	// MISSING: UserSpecifiedSystem
	out.GcsFilesetSpec = GcsFilesetSpecObservedState_FromProto(mapCtx, in.GetGcsFilesetSpec())
	out.BigqueryTableSpec = BigQueryTableSpecObservedState_FromProto(mapCtx, in.GetBigqueryTableSpec())
	out.BigqueryDateShardedSpec = BigQueryDateShardedSpecObservedState_FromProto(mapCtx, in.GetBigqueryDateShardedSpec())
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Schema
	out.SourceSystemTimestamps = SystemTimestamps_FromProto(mapCtx, in.GetSourceSystemTimestamps())
	out.UsageSignal = UsageSignal_FromProto(mapCtx, in.GetUsageSignal())
	return out
}
func EntryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EntryObservedState) *pb.Entry {
	if in == nil {
		return nil
	}
	out := &pb.Entry{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: LinkedResource
	// MISSING: Type
	// MISSING: UserSpecifiedType
	if oneof := EntryObservedState_IntegratedSystem_ToProto(mapCtx, in.IntegratedSystem); oneof != nil {
		out.System = oneof
	}
	// MISSING: UserSpecifiedSystem
	if oneof := GcsFilesetSpecObservedState_ToProto(mapCtx, in.GcsFilesetSpec); oneof != nil {
		out.TypeSpec = &pb.Entry_GcsFilesetSpec{GcsFilesetSpec: oneof}
	}
	if oneof := BigQueryTableSpecObservedState_ToProto(mapCtx, in.BigqueryTableSpec); oneof != nil {
		out.TypeSpec = &pb.Entry_BigqueryTableSpec{BigqueryTableSpec: oneof}
	}
	if oneof := BigQueryDateShardedSpecObservedState_ToProto(mapCtx, in.BigqueryDateShardedSpec); oneof != nil {
		out.TypeSpec = &pb.Entry_BigqueryDateShardedSpec{BigqueryDateShardedSpec: oneof}
	}
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Schema
	out.SourceSystemTimestamps = SystemTimestamps_ToProto(mapCtx, in.SourceSystemTimestamps)
	out.UsageSignal = UsageSignal_ToProto(mapCtx, in.UsageSignal)
	return out
}
func GcsFileSpec_FromProto(mapCtx *direct.MapContext, in *pb.GcsFileSpec) *krm.GcsFileSpec {
	if in == nil {
		return nil
	}
	out := &krm.GcsFileSpec{}
	out.FilePath = direct.LazyPtr(in.GetFilePath())
	// MISSING: GcsTimestamps
	// MISSING: SizeBytes
	return out
}
func GcsFileSpec_ToProto(mapCtx *direct.MapContext, in *krm.GcsFileSpec) *pb.GcsFileSpec {
	if in == nil {
		return nil
	}
	out := &pb.GcsFileSpec{}
	out.FilePath = direct.ValueOf(in.FilePath)
	// MISSING: GcsTimestamps
	// MISSING: SizeBytes
	return out
}
func GcsFileSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GcsFileSpec) *krm.GcsFileSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GcsFileSpecObservedState{}
	// MISSING: FilePath
	out.GcsTimestamps = SystemTimestamps_FromProto(mapCtx, in.GetGcsTimestamps())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	return out
}
func GcsFileSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GcsFileSpecObservedState) *pb.GcsFileSpec {
	if in == nil {
		return nil
	}
	out := &pb.GcsFileSpec{}
	// MISSING: FilePath
	out.GcsTimestamps = SystemTimestamps_ToProto(mapCtx, in.GcsTimestamps)
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	return out
}
func GcsFilesetSpec_FromProto(mapCtx *direct.MapContext, in *pb.GcsFilesetSpec) *krm.GcsFilesetSpec {
	if in == nil {
		return nil
	}
	out := &krm.GcsFilesetSpec{}
	out.FilePatterns = in.FilePatterns
	// MISSING: SampleGcsFileSpecs
	return out
}
func GcsFilesetSpec_ToProto(mapCtx *direct.MapContext, in *krm.GcsFilesetSpec) *pb.GcsFilesetSpec {
	if in == nil {
		return nil
	}
	out := &pb.GcsFilesetSpec{}
	out.FilePatterns = in.FilePatterns
	// MISSING: SampleGcsFileSpecs
	return out
}
func GcsFilesetSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GcsFilesetSpec) *krm.GcsFilesetSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GcsFilesetSpecObservedState{}
	// MISSING: FilePatterns
	out.SampleGcsFileSpecs = direct.Slice_FromProto(mapCtx, in.SampleGcsFileSpecs, GcsFileSpec_FromProto)
	return out
}
func GcsFilesetSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GcsFilesetSpecObservedState) *pb.GcsFilesetSpec {
	if in == nil {
		return nil
	}
	out := &pb.GcsFilesetSpec{}
	// MISSING: FilePatterns
	out.SampleGcsFileSpecs = direct.Slice_ToProto(mapCtx, in.SampleGcsFileSpecs, GcsFileSpec_ToProto)
	return out
}
func Schema_FromProto(mapCtx *direct.MapContext, in *pb.Schema) *krm.Schema {
	if in == nil {
		return nil
	}
	out := &krm.Schema{}
	out.Columns = direct.Slice_FromProto(mapCtx, in.Columns, ColumnSchema_FromProto)
	return out
}
func Schema_ToProto(mapCtx *direct.MapContext, in *krm.Schema) *pb.Schema {
	if in == nil {
		return nil
	}
	out := &pb.Schema{}
	out.Columns = direct.Slice_ToProto(mapCtx, in.Columns, ColumnSchema_ToProto)
	return out
}
func SystemTimestamps_FromProto(mapCtx *direct.MapContext, in *pb.SystemTimestamps) *krm.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &krm.SystemTimestamps{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: ExpireTime
	return out
}
func SystemTimestamps_ToProto(mapCtx *direct.MapContext, in *krm.SystemTimestamps) *pb.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &pb.SystemTimestamps{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: ExpireTime
	return out
}
func SystemTimestampsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SystemTimestamps) *krm.SystemTimestampsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SystemTimestampsObservedState{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	return out
}
func SystemTimestampsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SystemTimestampsObservedState) *pb.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &pb.SystemTimestamps{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	return out
}
func TableSpec_FromProto(mapCtx *direct.MapContext, in *pb.TableSpec) *krm.TableSpec {
	if in == nil {
		return nil
	}
	out := &krm.TableSpec{}
	// MISSING: GroupedEntry
	return out
}
func TableSpec_ToProto(mapCtx *direct.MapContext, in *krm.TableSpec) *pb.TableSpec {
	if in == nil {
		return nil
	}
	out := &pb.TableSpec{}
	// MISSING: GroupedEntry
	return out
}
func TableSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TableSpec) *krm.TableSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TableSpecObservedState{}
	out.GroupedEntry = direct.LazyPtr(in.GetGroupedEntry())
	return out
}
func TableSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TableSpecObservedState) *pb.TableSpec {
	if in == nil {
		return nil
	}
	out := &pb.TableSpec{}
	out.GroupedEntry = direct.ValueOf(in.GroupedEntry)
	return out
}
func UsageSignal_FromProto(mapCtx *direct.MapContext, in *pb.UsageSignal) *krm.UsageSignal {
	if in == nil {
		return nil
	}
	out := &krm.UsageSignal{}
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: UsageWithinTimeRange
	return out
}
func UsageSignal_ToProto(mapCtx *direct.MapContext, in *krm.UsageSignal) *pb.UsageSignal {
	if in == nil {
		return nil
	}
	out := &pb.UsageSignal{}
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: UsageWithinTimeRange
	return out
}
func UsageStats_FromProto(mapCtx *direct.MapContext, in *pb.UsageStats) *krm.UsageStats {
	if in == nil {
		return nil
	}
	out := &krm.UsageStats{}
	out.TotalCompletions = direct.LazyPtr(in.GetTotalCompletions())
	out.TotalFailures = direct.LazyPtr(in.GetTotalFailures())
	out.TotalCancellations = direct.LazyPtr(in.GetTotalCancellations())
	out.TotalExecutionTimeForCompletionsMillis = direct.LazyPtr(in.GetTotalExecutionTimeForCompletionsMillis())
	return out
}
func UsageStats_ToProto(mapCtx *direct.MapContext, in *krm.UsageStats) *pb.UsageStats {
	if in == nil {
		return nil
	}
	out := &pb.UsageStats{}
	out.TotalCompletions = direct.ValueOf(in.TotalCompletions)
	out.TotalFailures = direct.ValueOf(in.TotalFailures)
	out.TotalCancellations = direct.ValueOf(in.TotalCancellations)
	out.TotalExecutionTimeForCompletionsMillis = direct.ValueOf(in.TotalExecutionTimeForCompletionsMillis)
	return out
}
func ViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.ViewSpec) *krm.ViewSpec {
	if in == nil {
		return nil
	}
	out := &krm.ViewSpec{}
	// MISSING: ViewQuery
	return out
}
func ViewSpec_ToProto(mapCtx *direct.MapContext, in *krm.ViewSpec) *pb.ViewSpec {
	if in == nil {
		return nil
	}
	out := &pb.ViewSpec{}
	// MISSING: ViewQuery
	return out
}
func ViewSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ViewSpec) *krm.ViewSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ViewSpecObservedState{}
	out.ViewQuery = direct.LazyPtr(in.GetViewQuery())
	return out
}
func ViewSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ViewSpecObservedState) *pb.ViewSpec {
	if in == nil {
		return nil
	}
	out := &pb.ViewSpec{}
	out.ViewQuery = direct.ValueOf(in.ViewQuery)
	return out
}
