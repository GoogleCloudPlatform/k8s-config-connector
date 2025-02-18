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

package datastream

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AvroFileFormat_FromProto(mapCtx *direct.MapContext, in *pb.AvroFileFormat) *krm.AvroFileFormat {
	if in == nil {
		return nil
	}
	out := &krm.AvroFileFormat{}
	return out
}
func AvroFileFormat_ToProto(mapCtx *direct.MapContext, in *krm.AvroFileFormat) *pb.AvroFileFormat {
	if in == nil {
		return nil
	}
	out := &pb.AvroFileFormat{}
	return out
}
func BigQueryDestinationConfig_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig) *krm.BigQueryDestinationConfig {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDestinationConfig{}
	out.SingleTargetDataset = BigQueryDestinationConfig_SingleTargetDataset_FromProto(mapCtx, in.GetSingleTargetDataset())
	out.SourceHierarchyDatasets = BigQueryDestinationConfig_SourceHierarchyDatasets_FromProto(mapCtx, in.GetSourceHierarchyDatasets())
	out.DataFreshness = direct.StringDuration_FromProto(mapCtx, in.GetDataFreshness())
	out.Merge = BigQueryDestinationConfig_Merge_FromProto(mapCtx, in.GetMerge())
	out.AppendOnly = BigQueryDestinationConfig_AppendOnly_FromProto(mapCtx, in.GetAppendOnly())
	return out
}
func BigQueryDestinationConfig_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDestinationConfig) *pb.BigQueryDestinationConfig {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig{}
	if oneof := BigQueryDestinationConfig_SingleTargetDataset_ToProto(mapCtx, in.SingleTargetDataset); oneof != nil {
		out.DatasetConfig = &pb.BigQueryDestinationConfig_SingleTargetDataset_{SingleTargetDataset: oneof}
	}
	if oneof := BigQueryDestinationConfig_SourceHierarchyDatasets_ToProto(mapCtx, in.SourceHierarchyDatasets); oneof != nil {
		out.DatasetConfig = &pb.BigQueryDestinationConfig_SourceHierarchyDatasets_{SourceHierarchyDatasets: oneof}
	}
	out.DataFreshness = direct.StringDuration_ToProto(mapCtx, in.DataFreshness)
	if oneof := BigQueryDestinationConfig_Merge_ToProto(mapCtx, in.Merge); oneof != nil {
		out.WriteMode = &pb.BigQueryDestinationConfig_Merge_{Merge: oneof}
	}
	if oneof := BigQueryDestinationConfig_AppendOnly_ToProto(mapCtx, in.AppendOnly); oneof != nil {
		out.WriteMode = &pb.BigQueryDestinationConfig_AppendOnly_{AppendOnly: oneof}
	}
	return out
}
func BigQueryDestinationConfig_AppendOnly_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig_AppendOnly) *krm.BigQueryDestinationConfig_AppendOnly {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDestinationConfig_AppendOnly{}
	return out
}
func BigQueryDestinationConfig_AppendOnly_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDestinationConfig_AppendOnly) *pb.BigQueryDestinationConfig_AppendOnly {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_AppendOnly{}
	return out
}
func BigQueryDestinationConfig_Merge_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig_Merge) *krm.BigQueryDestinationConfig_Merge {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDestinationConfig_Merge{}
	return out
}
func BigQueryDestinationConfig_Merge_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDestinationConfig_Merge) *pb.BigQueryDestinationConfig_Merge {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_Merge{}
	return out
}
func BigQueryDestinationConfig_SingleTargetDataset_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig_SingleTargetDataset) *krm.BigQueryDestinationConfig_SingleTargetDataset {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDestinationConfig_SingleTargetDataset{}
	out.DatasetID = direct.LazyPtr(in.GetDatasetId())
	return out
}
func BigQueryDestinationConfig_SingleTargetDataset_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDestinationConfig_SingleTargetDataset) *pb.BigQueryDestinationConfig_SingleTargetDataset {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_SingleTargetDataset{}
	out.DatasetId = direct.ValueOf(in.DatasetID)
	return out
}
func BigQueryDestinationConfig_SourceHierarchyDatasets_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig_SourceHierarchyDatasets) *krm.BigQueryDestinationConfig_SourceHierarchyDatasets {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDestinationConfig_SourceHierarchyDatasets{}
	out.DatasetTemplate = BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate_FromProto(mapCtx, in.GetDatasetTemplate())
	return out
}
func BigQueryDestinationConfig_SourceHierarchyDatasets_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDestinationConfig_SourceHierarchyDatasets) *pb.BigQueryDestinationConfig_SourceHierarchyDatasets {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_SourceHierarchyDatasets{}
	out.DatasetTemplate = BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate_ToProto(mapCtx, in.DatasetTemplate)
	return out
}
func BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate) *krm.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.DatasetIDPrefix = direct.LazyPtr(in.GetDatasetIdPrefix())
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}
func BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate) *pb.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate{}
	out.Location = direct.ValueOf(in.Location)
	out.DatasetIdPrefix = direct.ValueOf(in.DatasetIDPrefix)
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	return out
}
func DatastreamStreamObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.DatastreamStreamObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamStreamObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: SourceConfig
	// MISSING: DestinationConfig
	// MISSING: State
	// MISSING: BackfillAll
	// MISSING: BackfillNone
	// MISSING: Errors
	// MISSING: CustomerManagedEncryptionKey
	// MISSING: LastRecoveryTime
	return out
}
func DatastreamStreamObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamStreamObservedState) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: SourceConfig
	// MISSING: DestinationConfig
	// MISSING: State
	// MISSING: BackfillAll
	// MISSING: BackfillNone
	// MISSING: Errors
	// MISSING: CustomerManagedEncryptionKey
	// MISSING: LastRecoveryTime
	return out
}
func DatastreamStreamSpec_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.DatastreamStreamSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamStreamSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: SourceConfig
	// MISSING: DestinationConfig
	// MISSING: State
	// MISSING: BackfillAll
	// MISSING: BackfillNone
	// MISSING: Errors
	// MISSING: CustomerManagedEncryptionKey
	// MISSING: LastRecoveryTime
	return out
}
func DatastreamStreamSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamStreamSpec) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: SourceConfig
	// MISSING: DestinationConfig
	// MISSING: State
	// MISSING: BackfillAll
	// MISSING: BackfillNone
	// MISSING: Errors
	// MISSING: CustomerManagedEncryptionKey
	// MISSING: LastRecoveryTime
	return out
}
func DestinationConfig_FromProto(mapCtx *direct.MapContext, in *pb.DestinationConfig) *krm.DestinationConfig {
	if in == nil {
		return nil
	}
	out := &krm.DestinationConfig{}
	out.DestinationConnectionProfile = direct.LazyPtr(in.GetDestinationConnectionProfile())
	out.GcsDestinationConfig = GcsDestinationConfig_FromProto(mapCtx, in.GetGcsDestinationConfig())
	out.BigqueryDestinationConfig = BigQueryDestinationConfig_FromProto(mapCtx, in.GetBigqueryDestinationConfig())
	return out
}
func DestinationConfig_ToProto(mapCtx *direct.MapContext, in *krm.DestinationConfig) *pb.DestinationConfig {
	if in == nil {
		return nil
	}
	out := &pb.DestinationConfig{}
	out.DestinationConnectionProfile = direct.ValueOf(in.DestinationConnectionProfile)
	if oneof := GcsDestinationConfig_ToProto(mapCtx, in.GcsDestinationConfig); oneof != nil {
		out.DestinationStreamConfig = &pb.DestinationConfig_GcsDestinationConfig{GcsDestinationConfig: oneof}
	}
	if oneof := BigQueryDestinationConfig_ToProto(mapCtx, in.BigqueryDestinationConfig); oneof != nil {
		out.DestinationStreamConfig = &pb.DestinationConfig_BigqueryDestinationConfig{BigqueryDestinationConfig: oneof}
	}
	return out
}
func Error_FromProto(mapCtx *direct.MapContext, in *pb.Error) *krm.Error {
	if in == nil {
		return nil
	}
	out := &krm.Error{}
	out.Reason = direct.LazyPtr(in.GetReason())
	out.ErrorUuid = direct.LazyPtr(in.GetErrorUuid())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.ErrorTime = direct.StringTimestamp_FromProto(mapCtx, in.GetErrorTime())
	out.Details = in.Details
	return out
}
func Error_ToProto(mapCtx *direct.MapContext, in *krm.Error) *pb.Error {
	if in == nil {
		return nil
	}
	out := &pb.Error{}
	out.Reason = direct.ValueOf(in.Reason)
	out.ErrorUuid = direct.ValueOf(in.ErrorUuid)
	out.Message = direct.ValueOf(in.Message)
	out.ErrorTime = direct.StringTimestamp_ToProto(mapCtx, in.ErrorTime)
	out.Details = in.Details
	return out
}
func GcsDestinationConfig_FromProto(mapCtx *direct.MapContext, in *pb.GcsDestinationConfig) *krm.GcsDestinationConfig {
	if in == nil {
		return nil
	}
	out := &krm.GcsDestinationConfig{}
	out.Path = direct.LazyPtr(in.GetPath())
	out.FileRotationMb = direct.LazyPtr(in.GetFileRotationMb())
	out.FileRotationInterval = direct.StringDuration_FromProto(mapCtx, in.GetFileRotationInterval())
	out.AvroFileFormat = AvroFileFormat_FromProto(mapCtx, in.GetAvroFileFormat())
	out.JsonFileFormat = JsonFileFormat_FromProto(mapCtx, in.GetJsonFileFormat())
	return out
}
func GcsDestinationConfig_ToProto(mapCtx *direct.MapContext, in *krm.GcsDestinationConfig) *pb.GcsDestinationConfig {
	if in == nil {
		return nil
	}
	out := &pb.GcsDestinationConfig{}
	out.Path = direct.ValueOf(in.Path)
	out.FileRotationMb = direct.ValueOf(in.FileRotationMb)
	out.FileRotationInterval = direct.StringDuration_ToProto(mapCtx, in.FileRotationInterval)
	if oneof := AvroFileFormat_ToProto(mapCtx, in.AvroFileFormat); oneof != nil {
		out.FileFormat = &pb.GcsDestinationConfig_AvroFileFormat{AvroFileFormat: oneof}
	}
	if oneof := JsonFileFormat_ToProto(mapCtx, in.JsonFileFormat); oneof != nil {
		out.FileFormat = &pb.GcsDestinationConfig_JsonFileFormat{JsonFileFormat: oneof}
	}
	return out
}
func JsonFileFormat_FromProto(mapCtx *direct.MapContext, in *pb.JsonFileFormat) *krm.JsonFileFormat {
	if in == nil {
		return nil
	}
	out := &krm.JsonFileFormat{}
	out.SchemaFileFormat = direct.Enum_FromProto(mapCtx, in.GetSchemaFileFormat())
	out.Compression = direct.Enum_FromProto(mapCtx, in.GetCompression())
	return out
}
func JsonFileFormat_ToProto(mapCtx *direct.MapContext, in *krm.JsonFileFormat) *pb.JsonFileFormat {
	if in == nil {
		return nil
	}
	out := &pb.JsonFileFormat{}
	out.SchemaFileFormat = direct.Enum_ToProto[pb.JsonFileFormat_SchemaFileFormat](mapCtx, in.SchemaFileFormat)
	out.Compression = direct.Enum_ToProto[pb.JsonFileFormat_JsonCompression](mapCtx, in.Compression)
	return out
}
func MysqlColumn_FromProto(mapCtx *direct.MapContext, in *pb.MysqlColumn) *krm.MysqlColumn {
	if in == nil {
		return nil
	}
	out := &krm.MysqlColumn{}
	out.Column = direct.LazyPtr(in.GetColumn())
	out.DataType = direct.LazyPtr(in.GetDataType())
	out.Length = direct.LazyPtr(in.GetLength())
	out.Collation = direct.LazyPtr(in.GetCollation())
	out.PrimaryKey = direct.LazyPtr(in.GetPrimaryKey())
	out.Nullable = direct.LazyPtr(in.GetNullable())
	out.OrdinalPosition = direct.LazyPtr(in.GetOrdinalPosition())
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.Scale = direct.LazyPtr(in.GetScale())
	return out
}
func MysqlColumn_ToProto(mapCtx *direct.MapContext, in *krm.MysqlColumn) *pb.MysqlColumn {
	if in == nil {
		return nil
	}
	out := &pb.MysqlColumn{}
	out.Column = direct.ValueOf(in.Column)
	out.DataType = direct.ValueOf(in.DataType)
	out.Length = direct.ValueOf(in.Length)
	out.Collation = direct.ValueOf(in.Collation)
	out.PrimaryKey = direct.ValueOf(in.PrimaryKey)
	out.Nullable = direct.ValueOf(in.Nullable)
	out.OrdinalPosition = direct.ValueOf(in.OrdinalPosition)
	out.Precision = direct.ValueOf(in.Precision)
	out.Scale = direct.ValueOf(in.Scale)
	return out
}
func MysqlDatabase_FromProto(mapCtx *direct.MapContext, in *pb.MysqlDatabase) *krm.MysqlDatabase {
	if in == nil {
		return nil
	}
	out := &krm.MysqlDatabase{}
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.MysqlTables = direct.Slice_FromProto(mapCtx, in.MysqlTables, MysqlTable_FromProto)
	return out
}
func MysqlDatabase_ToProto(mapCtx *direct.MapContext, in *krm.MysqlDatabase) *pb.MysqlDatabase {
	if in == nil {
		return nil
	}
	out := &pb.MysqlDatabase{}
	out.Database = direct.ValueOf(in.Database)
	out.MysqlTables = direct.Slice_ToProto(mapCtx, in.MysqlTables, MysqlTable_ToProto)
	return out
}
func MysqlRdbms_FromProto(mapCtx *direct.MapContext, in *pb.MysqlRdbms) *krm.MysqlRdbms {
	if in == nil {
		return nil
	}
	out := &krm.MysqlRdbms{}
	out.MysqlDatabases = direct.Slice_FromProto(mapCtx, in.MysqlDatabases, MysqlDatabase_FromProto)
	return out
}
func MysqlRdbms_ToProto(mapCtx *direct.MapContext, in *krm.MysqlRdbms) *pb.MysqlRdbms {
	if in == nil {
		return nil
	}
	out := &pb.MysqlRdbms{}
	out.MysqlDatabases = direct.Slice_ToProto(mapCtx, in.MysqlDatabases, MysqlDatabase_ToProto)
	return out
}
func MysqlSourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSourceConfig) *krm.MysqlSourceConfig {
	if in == nil {
		return nil
	}
	out := &krm.MysqlSourceConfig{}
	out.IncludeObjects = MysqlRdbms_FromProto(mapCtx, in.GetIncludeObjects())
	out.ExcludeObjects = MysqlRdbms_FromProto(mapCtx, in.GetExcludeObjects())
	out.MaxConcurrentCdcTasks = direct.LazyPtr(in.GetMaxConcurrentCdcTasks())
	out.MaxConcurrentBackfillTasks = direct.LazyPtr(in.GetMaxConcurrentBackfillTasks())
	out.BinaryLogPosition = MysqlSourceConfig_BinaryLogPosition_FromProto(mapCtx, in.GetBinaryLogPosition())
	out.Gtid = MysqlSourceConfig_Gtid_FromProto(mapCtx, in.GetGtid())
	return out
}
func MysqlSourceConfig_ToProto(mapCtx *direct.MapContext, in *krm.MysqlSourceConfig) *pb.MysqlSourceConfig {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSourceConfig{}
	out.IncludeObjects = MysqlRdbms_ToProto(mapCtx, in.IncludeObjects)
	out.ExcludeObjects = MysqlRdbms_ToProto(mapCtx, in.ExcludeObjects)
	out.MaxConcurrentCdcTasks = direct.ValueOf(in.MaxConcurrentCdcTasks)
	out.MaxConcurrentBackfillTasks = direct.ValueOf(in.MaxConcurrentBackfillTasks)
	if oneof := MysqlSourceConfig_BinaryLogPosition_ToProto(mapCtx, in.BinaryLogPosition); oneof != nil {
		out.CdcMethod = &pb.MysqlSourceConfig_BinaryLogPosition_{BinaryLogPosition: oneof}
	}
	if oneof := MysqlSourceConfig_Gtid_ToProto(mapCtx, in.Gtid); oneof != nil {
		out.CdcMethod = &pb.MysqlSourceConfig_Gtid_{Gtid: oneof}
	}
	return out
}
func MysqlSourceConfig_BinaryLogPosition_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSourceConfig_BinaryLogPosition) *krm.MysqlSourceConfig_BinaryLogPosition {
	if in == nil {
		return nil
	}
	out := &krm.MysqlSourceConfig_BinaryLogPosition{}
	return out
}
func MysqlSourceConfig_BinaryLogPosition_ToProto(mapCtx *direct.MapContext, in *krm.MysqlSourceConfig_BinaryLogPosition) *pb.MysqlSourceConfig_BinaryLogPosition {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSourceConfig_BinaryLogPosition{}
	return out
}
func MysqlSourceConfig_Gtid_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSourceConfig_Gtid) *krm.MysqlSourceConfig_Gtid {
	if in == nil {
		return nil
	}
	out := &krm.MysqlSourceConfig_Gtid{}
	return out
}
func MysqlSourceConfig_Gtid_ToProto(mapCtx *direct.MapContext, in *krm.MysqlSourceConfig_Gtid) *pb.MysqlSourceConfig_Gtid {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSourceConfig_Gtid{}
	return out
}
func MysqlTable_FromProto(mapCtx *direct.MapContext, in *pb.MysqlTable) *krm.MysqlTable {
	if in == nil {
		return nil
	}
	out := &krm.MysqlTable{}
	out.Table = direct.LazyPtr(in.GetTable())
	out.MysqlColumns = direct.Slice_FromProto(mapCtx, in.MysqlColumns, MysqlColumn_FromProto)
	return out
}
func MysqlTable_ToProto(mapCtx *direct.MapContext, in *krm.MysqlTable) *pb.MysqlTable {
	if in == nil {
		return nil
	}
	out := &pb.MysqlTable{}
	out.Table = direct.ValueOf(in.Table)
	out.MysqlColumns = direct.Slice_ToProto(mapCtx, in.MysqlColumns, MysqlColumn_ToProto)
	return out
}
func OracleColumn_FromProto(mapCtx *direct.MapContext, in *pb.OracleColumn) *krm.OracleColumn {
	if in == nil {
		return nil
	}
	out := &krm.OracleColumn{}
	out.Column = direct.LazyPtr(in.GetColumn())
	out.DataType = direct.LazyPtr(in.GetDataType())
	out.Length = direct.LazyPtr(in.GetLength())
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.Scale = direct.LazyPtr(in.GetScale())
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.PrimaryKey = direct.LazyPtr(in.GetPrimaryKey())
	out.Nullable = direct.LazyPtr(in.GetNullable())
	out.OrdinalPosition = direct.LazyPtr(in.GetOrdinalPosition())
	return out
}
func OracleColumn_ToProto(mapCtx *direct.MapContext, in *krm.OracleColumn) *pb.OracleColumn {
	if in == nil {
		return nil
	}
	out := &pb.OracleColumn{}
	out.Column = direct.ValueOf(in.Column)
	out.DataType = direct.ValueOf(in.DataType)
	out.Length = direct.ValueOf(in.Length)
	out.Precision = direct.ValueOf(in.Precision)
	out.Scale = direct.ValueOf(in.Scale)
	out.Encoding = direct.ValueOf(in.Encoding)
	out.PrimaryKey = direct.ValueOf(in.PrimaryKey)
	out.Nullable = direct.ValueOf(in.Nullable)
	out.OrdinalPosition = direct.ValueOf(in.OrdinalPosition)
	return out
}
func OracleRdbms_FromProto(mapCtx *direct.MapContext, in *pb.OracleRdbms) *krm.OracleRdbms {
	if in == nil {
		return nil
	}
	out := &krm.OracleRdbms{}
	out.OracleSchemas = direct.Slice_FromProto(mapCtx, in.OracleSchemas, OracleSchema_FromProto)
	return out
}
func OracleRdbms_ToProto(mapCtx *direct.MapContext, in *krm.OracleRdbms) *pb.OracleRdbms {
	if in == nil {
		return nil
	}
	out := &pb.OracleRdbms{}
	out.OracleSchemas = direct.Slice_ToProto(mapCtx, in.OracleSchemas, OracleSchema_ToProto)
	return out
}
func OracleSchema_FromProto(mapCtx *direct.MapContext, in *pb.OracleSchema) *krm.OracleSchema {
	if in == nil {
		return nil
	}
	out := &krm.OracleSchema{}
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.OracleTables = direct.Slice_FromProto(mapCtx, in.OracleTables, OracleTable_FromProto)
	return out
}
func OracleSchema_ToProto(mapCtx *direct.MapContext, in *krm.OracleSchema) *pb.OracleSchema {
	if in == nil {
		return nil
	}
	out := &pb.OracleSchema{}
	out.Schema = direct.ValueOf(in.Schema)
	out.OracleTables = direct.Slice_ToProto(mapCtx, in.OracleTables, OracleTable_ToProto)
	return out
}
func OracleSourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig) *krm.OracleSourceConfig {
	if in == nil {
		return nil
	}
	out := &krm.OracleSourceConfig{}
	out.IncludeObjects = OracleRdbms_FromProto(mapCtx, in.GetIncludeObjects())
	out.ExcludeObjects = OracleRdbms_FromProto(mapCtx, in.GetExcludeObjects())
	out.MaxConcurrentCdcTasks = direct.LazyPtr(in.GetMaxConcurrentCdcTasks())
	out.MaxConcurrentBackfillTasks = direct.LazyPtr(in.GetMaxConcurrentBackfillTasks())
	out.DropLargeObjects = OracleSourceConfig_DropLargeObjects_FromProto(mapCtx, in.GetDropLargeObjects())
	out.StreamLargeObjects = OracleSourceConfig_StreamLargeObjects_FromProto(mapCtx, in.GetStreamLargeObjects())
	out.LogMiner = OracleSourceConfig_LogMiner_FromProto(mapCtx, in.GetLogMiner())
	out.BinaryLogParser = OracleSourceConfig_BinaryLogParser_FromProto(mapCtx, in.GetBinaryLogParser())
	return out
}
func OracleSourceConfig_ToProto(mapCtx *direct.MapContext, in *krm.OracleSourceConfig) *pb.OracleSourceConfig {
	if in == nil {
		return nil
	}
	out := &pb.OracleSourceConfig{}
	out.IncludeObjects = OracleRdbms_ToProto(mapCtx, in.IncludeObjects)
	out.ExcludeObjects = OracleRdbms_ToProto(mapCtx, in.ExcludeObjects)
	out.MaxConcurrentCdcTasks = direct.ValueOf(in.MaxConcurrentCdcTasks)
	out.MaxConcurrentBackfillTasks = direct.ValueOf(in.MaxConcurrentBackfillTasks)
	if oneof := OracleSourceConfig_DropLargeObjects_ToProto(mapCtx, in.DropLargeObjects); oneof != nil {
		out.LargeObjectsHandling = &pb.OracleSourceConfig_DropLargeObjects_{DropLargeObjects: oneof}
	}
	if oneof := OracleSourceConfig_StreamLargeObjects_ToProto(mapCtx, in.StreamLargeObjects); oneof != nil {
		out.LargeObjectsHandling = &pb.OracleSourceConfig_StreamLargeObjects_{StreamLargeObjects: oneof}
	}
	if oneof := OracleSourceConfig_LogMiner_ToProto(mapCtx, in.LogMiner); oneof != nil {
		out.CdcMethod = &pb.OracleSourceConfig_LogMiner_{LogMiner: oneof}
	}
	if oneof := OracleSourceConfig_BinaryLogParser_ToProto(mapCtx, in.BinaryLogParser); oneof != nil {
		out.CdcMethod = &pb.OracleSourceConfig_BinaryLogParser_{BinaryLogParser: oneof}
	}
	return out
}
func OracleSourceConfig_BinaryLogParser_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig_BinaryLogParser) *krm.OracleSourceConfig_BinaryLogParser {
	if in == nil {
		return nil
	}
	out := &krm.OracleSourceConfig_BinaryLogParser{}
	out.OracleAsmLogFileAccess = OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess_FromProto(mapCtx, in.GetOracleAsmLogFileAccess())
	out.LogFileDirectories = OracleSourceConfig_BinaryLogParser_LogFileDirectories_FromProto(mapCtx, in.GetLogFileDirectories())
	return out
}
func OracleSourceConfig_BinaryLogParser_ToProto(mapCtx *direct.MapContext, in *krm.OracleSourceConfig_BinaryLogParser) *pb.OracleSourceConfig_BinaryLogParser {
	if in == nil {
		return nil
	}
	out := &pb.OracleSourceConfig_BinaryLogParser{}
	if oneof := OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess_ToProto(mapCtx, in.OracleAsmLogFileAccess); oneof != nil {
		out.LogFileAccess = &pb.OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess_{OracleAsmLogFileAccess: oneof}
	}
	if oneof := OracleSourceConfig_BinaryLogParser_LogFileDirectories_ToProto(mapCtx, in.LogFileDirectories); oneof != nil {
		out.LogFileAccess = &pb.OracleSourceConfig_BinaryLogParser_LogFileDirectories_{LogFileDirectories: oneof}
	}
	return out
}
func OracleSourceConfig_BinaryLogParser_LogFileDirectories_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig_BinaryLogParser_LogFileDirectories) *krm.OracleSourceConfig_BinaryLogParser_LogFileDirectories {
	if in == nil {
		return nil
	}
	out := &krm.OracleSourceConfig_BinaryLogParser_LogFileDirectories{}
	out.OnlineLogDirectory = direct.LazyPtr(in.GetOnlineLogDirectory())
	out.ArchivedLogDirectory = direct.LazyPtr(in.GetArchivedLogDirectory())
	return out
}
func OracleSourceConfig_BinaryLogParser_LogFileDirectories_ToProto(mapCtx *direct.MapContext, in *krm.OracleSourceConfig_BinaryLogParser_LogFileDirectories) *pb.OracleSourceConfig_BinaryLogParser_LogFileDirectories {
	if in == nil {
		return nil
	}
	out := &pb.OracleSourceConfig_BinaryLogParser_LogFileDirectories{}
	out.OnlineLogDirectory = direct.ValueOf(in.OnlineLogDirectory)
	out.ArchivedLogDirectory = direct.ValueOf(in.ArchivedLogDirectory)
	return out
}
func OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess) *krm.OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess {
	if in == nil {
		return nil
	}
	out := &krm.OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess{}
	return out
}
func OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess_ToProto(mapCtx *direct.MapContext, in *krm.OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess) *pb.OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess {
	if in == nil {
		return nil
	}
	out := &pb.OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess{}
	return out
}
func OracleSourceConfig_DropLargeObjects_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig_DropLargeObjects) *krm.OracleSourceConfig_DropLargeObjects {
	if in == nil {
		return nil
	}
	out := &krm.OracleSourceConfig_DropLargeObjects{}
	return out
}
func OracleSourceConfig_DropLargeObjects_ToProto(mapCtx *direct.MapContext, in *krm.OracleSourceConfig_DropLargeObjects) *pb.OracleSourceConfig_DropLargeObjects {
	if in == nil {
		return nil
	}
	out := &pb.OracleSourceConfig_DropLargeObjects{}
	return out
}
func OracleSourceConfig_LogMiner_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig_LogMiner) *krm.OracleSourceConfig_LogMiner {
	if in == nil {
		return nil
	}
	out := &krm.OracleSourceConfig_LogMiner{}
	return out
}
func OracleSourceConfig_LogMiner_ToProto(mapCtx *direct.MapContext, in *krm.OracleSourceConfig_LogMiner) *pb.OracleSourceConfig_LogMiner {
	if in == nil {
		return nil
	}
	out := &pb.OracleSourceConfig_LogMiner{}
	return out
}
func OracleSourceConfig_StreamLargeObjects_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig_StreamLargeObjects) *krm.OracleSourceConfig_StreamLargeObjects {
	if in == nil {
		return nil
	}
	out := &krm.OracleSourceConfig_StreamLargeObjects{}
	return out
}
func OracleSourceConfig_StreamLargeObjects_ToProto(mapCtx *direct.MapContext, in *krm.OracleSourceConfig_StreamLargeObjects) *pb.OracleSourceConfig_StreamLargeObjects {
	if in == nil {
		return nil
	}
	out := &pb.OracleSourceConfig_StreamLargeObjects{}
	return out
}
func OracleTable_FromProto(mapCtx *direct.MapContext, in *pb.OracleTable) *krm.OracleTable {
	if in == nil {
		return nil
	}
	out := &krm.OracleTable{}
	out.Table = direct.LazyPtr(in.GetTable())
	out.OracleColumns = direct.Slice_FromProto(mapCtx, in.OracleColumns, OracleColumn_FromProto)
	return out
}
func OracleTable_ToProto(mapCtx *direct.MapContext, in *krm.OracleTable) *pb.OracleTable {
	if in == nil {
		return nil
	}
	out := &pb.OracleTable{}
	out.Table = direct.ValueOf(in.Table)
	out.OracleColumns = direct.Slice_ToProto(mapCtx, in.OracleColumns, OracleColumn_ToProto)
	return out
}
func PostgresqlColumn_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlColumn) *krm.PostgresqlColumn {
	if in == nil {
		return nil
	}
	out := &krm.PostgresqlColumn{}
	out.Column = direct.LazyPtr(in.GetColumn())
	out.DataType = direct.LazyPtr(in.GetDataType())
	out.Length = direct.LazyPtr(in.GetLength())
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.Scale = direct.LazyPtr(in.GetScale())
	out.PrimaryKey = direct.LazyPtr(in.GetPrimaryKey())
	out.Nullable = direct.LazyPtr(in.GetNullable())
	out.OrdinalPosition = direct.LazyPtr(in.GetOrdinalPosition())
	return out
}
func PostgresqlColumn_ToProto(mapCtx *direct.MapContext, in *krm.PostgresqlColumn) *pb.PostgresqlColumn {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlColumn{}
	out.Column = direct.ValueOf(in.Column)
	out.DataType = direct.ValueOf(in.DataType)
	out.Length = direct.ValueOf(in.Length)
	out.Precision = direct.ValueOf(in.Precision)
	out.Scale = direct.ValueOf(in.Scale)
	out.PrimaryKey = direct.ValueOf(in.PrimaryKey)
	out.Nullable = direct.ValueOf(in.Nullable)
	out.OrdinalPosition = direct.ValueOf(in.OrdinalPosition)
	return out
}
func PostgresqlRdbms_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlRdbms) *krm.PostgresqlRdbms {
	if in == nil {
		return nil
	}
	out := &krm.PostgresqlRdbms{}
	out.PostgresqlSchemas = direct.Slice_FromProto(mapCtx, in.PostgresqlSchemas, PostgresqlSchema_FromProto)
	return out
}
func PostgresqlRdbms_ToProto(mapCtx *direct.MapContext, in *krm.PostgresqlRdbms) *pb.PostgresqlRdbms {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlRdbms{}
	out.PostgresqlSchemas = direct.Slice_ToProto(mapCtx, in.PostgresqlSchemas, PostgresqlSchema_ToProto)
	return out
}
func PostgresqlSchema_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlSchema) *krm.PostgresqlSchema {
	if in == nil {
		return nil
	}
	out := &krm.PostgresqlSchema{}
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.PostgresqlTables = direct.Slice_FromProto(mapCtx, in.PostgresqlTables, PostgresqlTable_FromProto)
	return out
}
func PostgresqlSchema_ToProto(mapCtx *direct.MapContext, in *krm.PostgresqlSchema) *pb.PostgresqlSchema {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlSchema{}
	out.Schema = direct.ValueOf(in.Schema)
	out.PostgresqlTables = direct.Slice_ToProto(mapCtx, in.PostgresqlTables, PostgresqlTable_ToProto)
	return out
}
func PostgresqlSourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlSourceConfig) *krm.PostgresqlSourceConfig {
	if in == nil {
		return nil
	}
	out := &krm.PostgresqlSourceConfig{}
	out.IncludeObjects = PostgresqlRdbms_FromProto(mapCtx, in.GetIncludeObjects())
	out.ExcludeObjects = PostgresqlRdbms_FromProto(mapCtx, in.GetExcludeObjects())
	out.ReplicationSlot = direct.LazyPtr(in.GetReplicationSlot())
	out.Publication = direct.LazyPtr(in.GetPublication())
	out.MaxConcurrentBackfillTasks = direct.LazyPtr(in.GetMaxConcurrentBackfillTasks())
	return out
}
func PostgresqlSourceConfig_ToProto(mapCtx *direct.MapContext, in *krm.PostgresqlSourceConfig) *pb.PostgresqlSourceConfig {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlSourceConfig{}
	out.IncludeObjects = PostgresqlRdbms_ToProto(mapCtx, in.IncludeObjects)
	out.ExcludeObjects = PostgresqlRdbms_ToProto(mapCtx, in.ExcludeObjects)
	out.ReplicationSlot = direct.ValueOf(in.ReplicationSlot)
	out.Publication = direct.ValueOf(in.Publication)
	out.MaxConcurrentBackfillTasks = direct.ValueOf(in.MaxConcurrentBackfillTasks)
	return out
}
func PostgresqlTable_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlTable) *krm.PostgresqlTable {
	if in == nil {
		return nil
	}
	out := &krm.PostgresqlTable{}
	out.Table = direct.LazyPtr(in.GetTable())
	out.PostgresqlColumns = direct.Slice_FromProto(mapCtx, in.PostgresqlColumns, PostgresqlColumn_FromProto)
	return out
}
func PostgresqlTable_ToProto(mapCtx *direct.MapContext, in *krm.PostgresqlTable) *pb.PostgresqlTable {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlTable{}
	out.Table = direct.ValueOf(in.Table)
	out.PostgresqlColumns = direct.Slice_ToProto(mapCtx, in.PostgresqlColumns, PostgresqlColumn_ToProto)
	return out
}
func SourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.SourceConfig) *krm.SourceConfig {
	if in == nil {
		return nil
	}
	out := &krm.SourceConfig{}
	out.SourceConnectionProfile = direct.LazyPtr(in.GetSourceConnectionProfile())
	out.OracleSourceConfig = OracleSourceConfig_FromProto(mapCtx, in.GetOracleSourceConfig())
	out.MysqlSourceConfig = MysqlSourceConfig_FromProto(mapCtx, in.GetMysqlSourceConfig())
	out.PostgresqlSourceConfig = PostgresqlSourceConfig_FromProto(mapCtx, in.GetPostgresqlSourceConfig())
	out.SqlServerSourceConfig = SqlServerSourceConfig_FromProto(mapCtx, in.GetSqlServerSourceConfig())
	return out
}
func SourceConfig_ToProto(mapCtx *direct.MapContext, in *krm.SourceConfig) *pb.SourceConfig {
	if in == nil {
		return nil
	}
	out := &pb.SourceConfig{}
	out.SourceConnectionProfile = direct.ValueOf(in.SourceConnectionProfile)
	if oneof := OracleSourceConfig_ToProto(mapCtx, in.OracleSourceConfig); oneof != nil {
		out.SourceStreamConfig = &pb.SourceConfig_OracleSourceConfig{OracleSourceConfig: oneof}
	}
	if oneof := MysqlSourceConfig_ToProto(mapCtx, in.MysqlSourceConfig); oneof != nil {
		out.SourceStreamConfig = &pb.SourceConfig_MysqlSourceConfig{MysqlSourceConfig: oneof}
	}
	if oneof := PostgresqlSourceConfig_ToProto(mapCtx, in.PostgresqlSourceConfig); oneof != nil {
		out.SourceStreamConfig = &pb.SourceConfig_PostgresqlSourceConfig{PostgresqlSourceConfig: oneof}
	}
	if oneof := SqlServerSourceConfig_ToProto(mapCtx, in.SqlServerSourceConfig); oneof != nil {
		out.SourceStreamConfig = &pb.SourceConfig_SqlServerSourceConfig{SqlServerSourceConfig: oneof}
	}
	return out
}
func SqlServerChangeTables_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerChangeTables) *krm.SqlServerChangeTables {
	if in == nil {
		return nil
	}
	out := &krm.SqlServerChangeTables{}
	return out
}
func SqlServerChangeTables_ToProto(mapCtx *direct.MapContext, in *krm.SqlServerChangeTables) *pb.SqlServerChangeTables {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerChangeTables{}
	return out
}
func SqlServerColumn_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerColumn) *krm.SqlServerColumn {
	if in == nil {
		return nil
	}
	out := &krm.SqlServerColumn{}
	out.Column = direct.LazyPtr(in.GetColumn())
	out.DataType = direct.LazyPtr(in.GetDataType())
	out.Length = direct.LazyPtr(in.GetLength())
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.Scale = direct.LazyPtr(in.GetScale())
	out.PrimaryKey = direct.LazyPtr(in.GetPrimaryKey())
	out.Nullable = direct.LazyPtr(in.GetNullable())
	out.OrdinalPosition = direct.LazyPtr(in.GetOrdinalPosition())
	return out
}
func SqlServerColumn_ToProto(mapCtx *direct.MapContext, in *krm.SqlServerColumn) *pb.SqlServerColumn {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerColumn{}
	out.Column = direct.ValueOf(in.Column)
	out.DataType = direct.ValueOf(in.DataType)
	out.Length = direct.ValueOf(in.Length)
	out.Precision = direct.ValueOf(in.Precision)
	out.Scale = direct.ValueOf(in.Scale)
	out.PrimaryKey = direct.ValueOf(in.PrimaryKey)
	out.Nullable = direct.ValueOf(in.Nullable)
	out.OrdinalPosition = direct.ValueOf(in.OrdinalPosition)
	return out
}
func SqlServerRdbms_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerRdbms) *krm.SqlServerRdbms {
	if in == nil {
		return nil
	}
	out := &krm.SqlServerRdbms{}
	out.Schemas = direct.Slice_FromProto(mapCtx, in.Schemas, SqlServerSchema_FromProto)
	return out
}
func SqlServerRdbms_ToProto(mapCtx *direct.MapContext, in *krm.SqlServerRdbms) *pb.SqlServerRdbms {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerRdbms{}
	out.Schemas = direct.Slice_ToProto(mapCtx, in.Schemas, SqlServerSchema_ToProto)
	return out
}
func SqlServerSchema_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerSchema) *krm.SqlServerSchema {
	if in == nil {
		return nil
	}
	out := &krm.SqlServerSchema{}
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.Tables = direct.Slice_FromProto(mapCtx, in.Tables, SqlServerTable_FromProto)
	return out
}
func SqlServerSchema_ToProto(mapCtx *direct.MapContext, in *krm.SqlServerSchema) *pb.SqlServerSchema {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerSchema{}
	out.Schema = direct.ValueOf(in.Schema)
	out.Tables = direct.Slice_ToProto(mapCtx, in.Tables, SqlServerTable_ToProto)
	return out
}
func SqlServerSourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerSourceConfig) *krm.SqlServerSourceConfig {
	if in == nil {
		return nil
	}
	out := &krm.SqlServerSourceConfig{}
	out.IncludeObjects = SqlServerRdbms_FromProto(mapCtx, in.GetIncludeObjects())
	out.ExcludeObjects = SqlServerRdbms_FromProto(mapCtx, in.GetExcludeObjects())
	out.MaxConcurrentCdcTasks = direct.LazyPtr(in.GetMaxConcurrentCdcTasks())
	out.MaxConcurrentBackfillTasks = direct.LazyPtr(in.GetMaxConcurrentBackfillTasks())
	out.TransactionLogs = SqlServerTransactionLogs_FromProto(mapCtx, in.GetTransactionLogs())
	out.ChangeTables = SqlServerChangeTables_FromProto(mapCtx, in.GetChangeTables())
	return out
}
func SqlServerSourceConfig_ToProto(mapCtx *direct.MapContext, in *krm.SqlServerSourceConfig) *pb.SqlServerSourceConfig {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerSourceConfig{}
	out.IncludeObjects = SqlServerRdbms_ToProto(mapCtx, in.IncludeObjects)
	out.ExcludeObjects = SqlServerRdbms_ToProto(mapCtx, in.ExcludeObjects)
	out.MaxConcurrentCdcTasks = direct.ValueOf(in.MaxConcurrentCdcTasks)
	out.MaxConcurrentBackfillTasks = direct.ValueOf(in.MaxConcurrentBackfillTasks)
	if oneof := SqlServerTransactionLogs_ToProto(mapCtx, in.TransactionLogs); oneof != nil {
		out.CdcMethod = &pb.SqlServerSourceConfig_TransactionLogs{TransactionLogs: oneof}
	}
	if oneof := SqlServerChangeTables_ToProto(mapCtx, in.ChangeTables); oneof != nil {
		out.CdcMethod = &pb.SqlServerSourceConfig_ChangeTables{ChangeTables: oneof}
	}
	return out
}
func SqlServerTable_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerTable) *krm.SqlServerTable {
	if in == nil {
		return nil
	}
	out := &krm.SqlServerTable{}
	out.Table = direct.LazyPtr(in.GetTable())
	out.Columns = direct.Slice_FromProto(mapCtx, in.Columns, SqlServerColumn_FromProto)
	return out
}
func SqlServerTable_ToProto(mapCtx *direct.MapContext, in *krm.SqlServerTable) *pb.SqlServerTable {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerTable{}
	out.Table = direct.ValueOf(in.Table)
	out.Columns = direct.Slice_ToProto(mapCtx, in.Columns, SqlServerColumn_ToProto)
	return out
}
func SqlServerTransactionLogs_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerTransactionLogs) *krm.SqlServerTransactionLogs {
	if in == nil {
		return nil
	}
	out := &krm.SqlServerTransactionLogs{}
	return out
}
func SqlServerTransactionLogs_ToProto(mapCtx *direct.MapContext, in *krm.SqlServerTransactionLogs) *pb.SqlServerTransactionLogs {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerTransactionLogs{}
	return out
}
func Stream_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.Stream {
	if in == nil {
		return nil
	}
	out := &krm.Stream{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SourceConfig = SourceConfig_FromProto(mapCtx, in.GetSourceConfig())
	out.DestinationConfig = DestinationConfig_FromProto(mapCtx, in.GetDestinationConfig())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.BackfillAll = Stream_BackfillAllStrategy_FromProto(mapCtx, in.GetBackfillAll())
	out.BackfillNone = Stream_BackfillNoneStrategy_FromProto(mapCtx, in.GetBackfillNone())
	// MISSING: Errors
	out.CustomerManagedEncryptionKey = in.CustomerManagedEncryptionKey
	// MISSING: LastRecoveryTime
	return out
}
func Stream_ToProto(mapCtx *direct.MapContext, in *krm.Stream) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.SourceConfig = SourceConfig_ToProto(mapCtx, in.SourceConfig)
	out.DestinationConfig = DestinationConfig_ToProto(mapCtx, in.DestinationConfig)
	out.State = direct.Enum_ToProto[pb.Stream_State](mapCtx, in.State)
	if oneof := Stream_BackfillAllStrategy_ToProto(mapCtx, in.BackfillAll); oneof != nil {
		out.BackfillStrategy = &pb.Stream_BackfillAll{BackfillAll: oneof}
	}
	if oneof := Stream_BackfillNoneStrategy_ToProto(mapCtx, in.BackfillNone); oneof != nil {
		out.BackfillStrategy = &pb.Stream_BackfillNone{BackfillNone: oneof}
	}
	// MISSING: Errors
	out.CustomerManagedEncryptionKey = in.CustomerManagedEncryptionKey
	// MISSING: LastRecoveryTime
	return out
}
func StreamObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.StreamObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StreamObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: SourceConfig
	// MISSING: DestinationConfig
	// MISSING: State
	// MISSING: BackfillAll
	// MISSING: BackfillNone
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, Error_FromProto)
	// MISSING: CustomerManagedEncryptionKey
	out.LastRecoveryTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastRecoveryTime())
	return out
}
func StreamObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StreamObservedState) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: SourceConfig
	// MISSING: DestinationConfig
	// MISSING: State
	// MISSING: BackfillAll
	// MISSING: BackfillNone
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, Error_ToProto)
	// MISSING: CustomerManagedEncryptionKey
	out.LastRecoveryTime = direct.StringTimestamp_ToProto(mapCtx, in.LastRecoveryTime)
	return out
}
func Stream_BackfillAllStrategy_FromProto(mapCtx *direct.MapContext, in *pb.Stream_BackfillAllStrategy) *krm.Stream_BackfillAllStrategy {
	if in == nil {
		return nil
	}
	out := &krm.Stream_BackfillAllStrategy{}
	out.OracleExcludedObjects = OracleRdbms_FromProto(mapCtx, in.GetOracleExcludedObjects())
	out.MysqlExcludedObjects = MysqlRdbms_FromProto(mapCtx, in.GetMysqlExcludedObjects())
	out.PostgresqlExcludedObjects = PostgresqlRdbms_FromProto(mapCtx, in.GetPostgresqlExcludedObjects())
	out.SqlServerExcludedObjects = SqlServerRdbms_FromProto(mapCtx, in.GetSqlServerExcludedObjects())
	return out
}
func Stream_BackfillAllStrategy_ToProto(mapCtx *direct.MapContext, in *krm.Stream_BackfillAllStrategy) *pb.Stream_BackfillAllStrategy {
	if in == nil {
		return nil
	}
	out := &pb.Stream_BackfillAllStrategy{}
	if oneof := OracleRdbms_ToProto(mapCtx, in.OracleExcludedObjects); oneof != nil {
		out.ExcludedObjects = &pb.Stream_BackfillAllStrategy_OracleExcludedObjects{OracleExcludedObjects: oneof}
	}
	if oneof := MysqlRdbms_ToProto(mapCtx, in.MysqlExcludedObjects); oneof != nil {
		out.ExcludedObjects = &pb.Stream_BackfillAllStrategy_MysqlExcludedObjects{MysqlExcludedObjects: oneof}
	}
	if oneof := PostgresqlRdbms_ToProto(mapCtx, in.PostgresqlExcludedObjects); oneof != nil {
		out.ExcludedObjects = &pb.Stream_BackfillAllStrategy_PostgresqlExcludedObjects{PostgresqlExcludedObjects: oneof}
	}
	if oneof := SqlServerRdbms_ToProto(mapCtx, in.SqlServerExcludedObjects); oneof != nil {
		out.ExcludedObjects = &pb.Stream_BackfillAllStrategy_SqlServerExcludedObjects{SqlServerExcludedObjects: oneof}
	}
	return out
}
func Stream_BackfillNoneStrategy_FromProto(mapCtx *direct.MapContext, in *pb.Stream_BackfillNoneStrategy) *krm.Stream_BackfillNoneStrategy {
	if in == nil {
		return nil
	}
	out := &krm.Stream_BackfillNoneStrategy{}
	return out
}
func Stream_BackfillNoneStrategy_ToProto(mapCtx *direct.MapContext, in *krm.Stream_BackfillNoneStrategy) *pb.Stream_BackfillNoneStrategy {
	if in == nil {
		return nil
	}
	out := &pb.Stream_BackfillNoneStrategy{}
	return out
}
