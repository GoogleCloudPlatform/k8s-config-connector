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
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
func BigQueryDestinationConfig_SingleTargetDataset_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig_SingleTargetDataset) *krm.BigQueryDestinationConfig_SingleTargetDataset {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDestinationConfig_SingleTargetDataset{}
	if in.GetDatasetId() != "" {
		out.DatasetRef = &refs.BigQueryDatasetRef{
			External: in.GetDatasetId(),
		}
	}
	return out
}
func BigQueryDestinationConfig_SingleTargetDataset_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDestinationConfig_SingleTargetDataset) *pb.BigQueryDestinationConfig_SingleTargetDataset {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_SingleTargetDataset{}
	out.DatasetId = in.DatasetRef.External
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
	if in.GetKmsKeyName() != "" {
		out.KMSKeyRef = &refs.KMSCryptoKeyRef{
			External: in.GetKmsKeyName(),
		}
	}
	return out
}
func BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate) *pb.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate{}
	out.Location = direct.ValueOf(in.Location)
	out.DatasetIdPrefix = direct.ValueOf(in.DatasetIDPrefix)
	out.KmsKeyName = in.KMSKeyRef.External
	return out
}
func DatastreamStreamObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.DatastreamStreamObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamStreamObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, Error_FromProto)
	return out
}
func DatastreamStreamObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamStreamObservedState) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Stream_State](mapCtx, in.State)
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, Error_ToProto)
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
	if oneof := AvroFileFormat_FromProto(mapCtx, in.GetAvroFileFormat()); oneof != nil {
		out.AvroFileFormat = oneof
	}
	if oneof := JsonFileFormat_FromProto(mapCtx, in.GetJsonFileFormat()); oneof != nil {
		out.JsonFileFormat = oneof
	}
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
	out.OracleSchemas = direct.Slice_FromProto(mapCtx, in.GetOracleSchemas(), OracleSchema_FromProto)
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
	out.OracleTables = direct.Slice_FromProto(mapCtx, in.GetOracleTables(), OracleTable_FromProto)
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
	out.OracleColumns = direct.Slice_FromProto(mapCtx, in.GetOracleColumns(), OracleColumn_FromProto)
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
	out.PostgresqlSchemas = direct.Slice_FromProto(mapCtx, in.GetPostgresqlSchemas(), PostgresqlSchema_FromProto)
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
	out.PostgresqlTables = direct.Slice_FromProto(mapCtx, in.GetPostgresqlTables(), PostgresqlTable_FromProto)
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
	if oneof := PostgresqlRdbms_FromProto(mapCtx, in.GetIncludeObjects()); oneof != nil {
		out.IncludeObjects = oneof
	}
	if oneof := PostgresqlRdbms_FromProto(mapCtx, in.GetExcludeObjects()); oneof != nil {
		out.ExcludeObjects = oneof
	}
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
	out.PostgresqlColumns = direct.Slice_FromProto(mapCtx, in.GetPostgresqlColumns(), PostgresqlColumn_FromProto)
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
