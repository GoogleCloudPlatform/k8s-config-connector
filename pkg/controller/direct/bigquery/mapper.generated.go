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

package bigquery

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/bigquery/migration/apiv2/migrationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AzureSynapseDialect_FromProto(mapCtx *direct.MapContext, in *pb.AzureSynapseDialect) *krm.AzureSynapseDialect {
	if in == nil {
		return nil
	}
	out := &krm.AzureSynapseDialect{}
	return out
}
func AzureSynapseDialect_ToProto(mapCtx *direct.MapContext, in *krm.AzureSynapseDialect) *pb.AzureSynapseDialect {
	if in == nil {
		return nil
	}
	out := &pb.AzureSynapseDialect{}
	return out
}
func BigQueryDialect_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDialect) *krm.BigQueryDialect {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDialect{}
	return out
}
func BigQueryDialect_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDialect) *pb.BigQueryDialect {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDialect{}
	return out
}
func BigqueryMigrationWorkflowObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationWorkflow) *krm.BigqueryMigrationWorkflowObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryMigrationWorkflowObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Tasks
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	return out
}
func BigqueryMigrationWorkflowObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryMigrationWorkflowObservedState) *pb.MigrationWorkflow {
	if in == nil {
		return nil
	}
	out := &pb.MigrationWorkflow{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Tasks
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	return out
}
func BigqueryMigrationWorkflowSpec_FromProto(mapCtx *direct.MapContext, in *pb.MigrationWorkflow) *krm.BigqueryMigrationWorkflowSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryMigrationWorkflowSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Tasks
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	return out
}
func BigqueryMigrationWorkflowSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryMigrationWorkflowSpec) *pb.MigrationWorkflow {
	if in == nil {
		return nil
	}
	out := &pb.MigrationWorkflow{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Tasks
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	return out
}
func DB2Dialect_FromProto(mapCtx *direct.MapContext, in *pb.DB2Dialect) *krm.DB2Dialect {
	if in == nil {
		return nil
	}
	out := &krm.DB2Dialect{}
	return out
}
func DB2Dialect_ToProto(mapCtx *direct.MapContext, in *krm.DB2Dialect) *pb.DB2Dialect {
	if in == nil {
		return nil
	}
	out := &pb.DB2Dialect{}
	return out
}
func Dialect_FromProto(mapCtx *direct.MapContext, in *pb.Dialect) *krm.Dialect {
	if in == nil {
		return nil
	}
	out := &krm.Dialect{}
	out.BigqueryDialect = BigQueryDialect_FromProto(mapCtx, in.GetBigqueryDialect())
	out.HiveqlDialect = HiveQLDialect_FromProto(mapCtx, in.GetHiveqlDialect())
	out.RedshiftDialect = RedshiftDialect_FromProto(mapCtx, in.GetRedshiftDialect())
	out.TeradataDialect = TeradataDialect_FromProto(mapCtx, in.GetTeradataDialect())
	out.OracleDialect = OracleDialect_FromProto(mapCtx, in.GetOracleDialect())
	out.SparksqlDialect = SparkSQLDialect_FromProto(mapCtx, in.GetSparksqlDialect())
	out.SnowflakeDialect = SnowflakeDialect_FromProto(mapCtx, in.GetSnowflakeDialect())
	out.NetezzaDialect = NetezzaDialect_FromProto(mapCtx, in.GetNetezzaDialect())
	out.AzureSynapseDialect = AzureSynapseDialect_FromProto(mapCtx, in.GetAzureSynapseDialect())
	out.VerticaDialect = VerticaDialect_FromProto(mapCtx, in.GetVerticaDialect())
	out.SqlServerDialect = SQLServerDialect_FromProto(mapCtx, in.GetSqlServerDialect())
	out.PostgresqlDialect = PostgresqlDialect_FromProto(mapCtx, in.GetPostgresqlDialect())
	out.PrestoDialect = PrestoDialect_FromProto(mapCtx, in.GetPrestoDialect())
	out.MysqlDialect = MySQLDialect_FromProto(mapCtx, in.GetMysqlDialect())
	out.Db2Dialect = DB2Dialect_FromProto(mapCtx, in.GetDb2Dialect())
	out.SqliteDialect = SQLiteDialect_FromProto(mapCtx, in.GetSqliteDialect())
	out.GreenplumDialect = GreenplumDialect_FromProto(mapCtx, in.GetGreenplumDialect())
	return out
}
func Dialect_ToProto(mapCtx *direct.MapContext, in *krm.Dialect) *pb.Dialect {
	if in == nil {
		return nil
	}
	out := &pb.Dialect{}
	if oneof := BigQueryDialect_ToProto(mapCtx, in.BigqueryDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_BigqueryDialect{BigqueryDialect: oneof}
	}
	if oneof := HiveQLDialect_ToProto(mapCtx, in.HiveqlDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_HiveqlDialect{HiveqlDialect: oneof}
	}
	if oneof := RedshiftDialect_ToProto(mapCtx, in.RedshiftDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_RedshiftDialect{RedshiftDialect: oneof}
	}
	if oneof := TeradataDialect_ToProto(mapCtx, in.TeradataDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_TeradataDialect{TeradataDialect: oneof}
	}
	if oneof := OracleDialect_ToProto(mapCtx, in.OracleDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_OracleDialect{OracleDialect: oneof}
	}
	if oneof := SparkSQLDialect_ToProto(mapCtx, in.SparksqlDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_SparksqlDialect{SparksqlDialect: oneof}
	}
	if oneof := SnowflakeDialect_ToProto(mapCtx, in.SnowflakeDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_SnowflakeDialect{SnowflakeDialect: oneof}
	}
	if oneof := NetezzaDialect_ToProto(mapCtx, in.NetezzaDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_NetezzaDialect{NetezzaDialect: oneof}
	}
	if oneof := AzureSynapseDialect_ToProto(mapCtx, in.AzureSynapseDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_AzureSynapseDialect{AzureSynapseDialect: oneof}
	}
	if oneof := VerticaDialect_ToProto(mapCtx, in.VerticaDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_VerticaDialect{VerticaDialect: oneof}
	}
	if oneof := SQLServerDialect_ToProto(mapCtx, in.SqlServerDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_SqlServerDialect{SqlServerDialect: oneof}
	}
	if oneof := PostgresqlDialect_ToProto(mapCtx, in.PostgresqlDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_PostgresqlDialect{PostgresqlDialect: oneof}
	}
	if oneof := PrestoDialect_ToProto(mapCtx, in.PrestoDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_PrestoDialect{PrestoDialect: oneof}
	}
	if oneof := MySQLDialect_ToProto(mapCtx, in.MysqlDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_MysqlDialect{MysqlDialect: oneof}
	}
	if oneof := DB2Dialect_ToProto(mapCtx, in.Db2Dialect); oneof != nil {
		out.DialectValue = &pb.Dialect_Db2Dialect{Db2Dialect: oneof}
	}
	if oneof := SQLiteDialect_ToProto(mapCtx, in.SqliteDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_SqliteDialect{SqliteDialect: oneof}
	}
	if oneof := GreenplumDialect_ToProto(mapCtx, in.GreenplumDialect); oneof != nil {
		out.DialectValue = &pb.Dialect_GreenplumDialect{GreenplumDialect: oneof}
	}
	return out
}
func ErrorDetail_FromProto(mapCtx *direct.MapContext, in *pb.ErrorDetail) *krm.ErrorDetail {
	if in == nil {
		return nil
	}
	out := &krm.ErrorDetail{}
	out.Location = ErrorLocation_FromProto(mapCtx, in.GetLocation())
	out.ErrorInfo = ErrorInfo_FromProto(mapCtx, in.GetErrorInfo())
	return out
}
func ErrorDetail_ToProto(mapCtx *direct.MapContext, in *krm.ErrorDetail) *pb.ErrorDetail {
	if in == nil {
		return nil
	}
	out := &pb.ErrorDetail{}
	out.Location = ErrorLocation_ToProto(mapCtx, in.Location)
	out.ErrorInfo = ErrorInfo_ToProto(mapCtx, in.ErrorInfo)
	return out
}
func ErrorLocation_FromProto(mapCtx *direct.MapContext, in *pb.ErrorLocation) *krm.ErrorLocation {
	if in == nil {
		return nil
	}
	out := &krm.ErrorLocation{}
	out.Line = direct.LazyPtr(in.GetLine())
	out.Column = direct.LazyPtr(in.GetColumn())
	return out
}
func ErrorLocation_ToProto(mapCtx *direct.MapContext, in *krm.ErrorLocation) *pb.ErrorLocation {
	if in == nil {
		return nil
	}
	out := &pb.ErrorLocation{}
	out.Line = direct.ValueOf(in.Line)
	out.Column = direct.ValueOf(in.Column)
	return out
}
func GcsReportLogMessage_FromProto(mapCtx *direct.MapContext, in *pb.GcsReportLogMessage) *krm.GcsReportLogMessage {
	if in == nil {
		return nil
	}
	out := &krm.GcsReportLogMessage{}
	out.Severity = direct.LazyPtr(in.GetSeverity())
	out.Category = direct.LazyPtr(in.GetCategory())
	out.FilePath = direct.LazyPtr(in.GetFilePath())
	out.Filename = direct.LazyPtr(in.GetFilename())
	out.SourceScriptLine = direct.LazyPtr(in.GetSourceScriptLine())
	out.SourceScriptColumn = direct.LazyPtr(in.GetSourceScriptColumn())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.ScriptContext = direct.LazyPtr(in.GetScriptContext())
	out.Action = direct.LazyPtr(in.GetAction())
	out.Effect = direct.LazyPtr(in.GetEffect())
	out.ObjectName = direct.LazyPtr(in.GetObjectName())
	return out
}
func GcsReportLogMessage_ToProto(mapCtx *direct.MapContext, in *krm.GcsReportLogMessage) *pb.GcsReportLogMessage {
	if in == nil {
		return nil
	}
	out := &pb.GcsReportLogMessage{}
	out.Severity = direct.ValueOf(in.Severity)
	out.Category = direct.ValueOf(in.Category)
	out.FilePath = direct.ValueOf(in.FilePath)
	out.Filename = direct.ValueOf(in.Filename)
	out.SourceScriptLine = direct.ValueOf(in.SourceScriptLine)
	out.SourceScriptColumn = direct.ValueOf(in.SourceScriptColumn)
	out.Message = direct.ValueOf(in.Message)
	out.ScriptContext = direct.ValueOf(in.ScriptContext)
	out.Action = direct.ValueOf(in.Action)
	out.Effect = direct.ValueOf(in.Effect)
	out.ObjectName = direct.ValueOf(in.ObjectName)
	return out
}
func GreenplumDialect_FromProto(mapCtx *direct.MapContext, in *pb.GreenplumDialect) *krm.GreenplumDialect {
	if in == nil {
		return nil
	}
	out := &krm.GreenplumDialect{}
	return out
}
func GreenplumDialect_ToProto(mapCtx *direct.MapContext, in *krm.GreenplumDialect) *pb.GreenplumDialect {
	if in == nil {
		return nil
	}
	out := &pb.GreenplumDialect{}
	return out
}
func HiveQLDialect_FromProto(mapCtx *direct.MapContext, in *pb.HiveQLDialect) *krm.HiveQLDialect {
	if in == nil {
		return nil
	}
	out := &krm.HiveQLDialect{}
	return out
}
func HiveQLDialect_ToProto(mapCtx *direct.MapContext, in *krm.HiveQLDialect) *pb.HiveQLDialect {
	if in == nil {
		return nil
	}
	out := &pb.HiveQLDialect{}
	return out
}
func Literal_FromProto(mapCtx *direct.MapContext, in *pb.Literal) *krm.Literal {
	if in == nil {
		return nil
	}
	out := &krm.Literal{}
	out.LiteralString = direct.LazyPtr(in.GetLiteralString())
	out.LiteralBytes = in.GetLiteralBytes()
	out.RelativePath = direct.LazyPtr(in.GetRelativePath())
	return out
}
func Literal_ToProto(mapCtx *direct.MapContext, in *krm.Literal) *pb.Literal {
	if in == nil {
		return nil
	}
	out := &pb.Literal{}
	if oneof := Literal_LiteralString_ToProto(mapCtx, in.LiteralString); oneof != nil {
		out.LiteralData = oneof
	}
	if oneof := Literal_LiteralBytes_ToProto(mapCtx, in.LiteralBytes); oneof != nil {
		out.LiteralData = oneof
	}
	out.RelativePath = direct.ValueOf(in.RelativePath)
	return out
}
func MigrationTask_FromProto(mapCtx *direct.MapContext, in *pb.MigrationTask) *krm.MigrationTask {
	if in == nil {
		return nil
	}
	out := &krm.MigrationTask{}
	out.TranslationConfigDetails = TranslationConfigDetails_FromProto(mapCtx, in.GetTranslationConfigDetails())
	out.TranslationDetails = TranslationDetails_FromProto(mapCtx, in.GetTranslationDetails())
	// MISSING: ID
	out.Type = direct.LazyPtr(in.GetType())
	// MISSING: State
	// MISSING: ProcessingError
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.LastUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastUpdateTime())
	// MISSING: ResourceErrorDetails
	out.ResourceErrorCount = direct.LazyPtr(in.GetResourceErrorCount())
	out.Metrics = direct.Slice_FromProto(mapCtx, in.Metrics, TimeSeries_FromProto)
	// MISSING: TaskResult
	out.TotalProcessingErrorCount = direct.LazyPtr(in.GetTotalProcessingErrorCount())
	out.TotalResourceErrorCount = direct.LazyPtr(in.GetTotalResourceErrorCount())
	return out
}
func MigrationTask_ToProto(mapCtx *direct.MapContext, in *krm.MigrationTask) *pb.MigrationTask {
	if in == nil {
		return nil
	}
	out := &pb.MigrationTask{}
	if oneof := TranslationConfigDetails_ToProto(mapCtx, in.TranslationConfigDetails); oneof != nil {
		out.TaskDetails = &pb.MigrationTask_TranslationConfigDetails{TranslationConfigDetails: oneof}
	}
	if oneof := TranslationDetails_ToProto(mapCtx, in.TranslationDetails); oneof != nil {
		out.TaskDetails = &pb.MigrationTask_TranslationDetails{TranslationDetails: oneof}
	}
	// MISSING: ID
	out.Type = direct.ValueOf(in.Type)
	// MISSING: State
	// MISSING: ProcessingError
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.LastUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.LastUpdateTime)
	// MISSING: ResourceErrorDetails
	out.ResourceErrorCount = direct.ValueOf(in.ResourceErrorCount)
	out.Metrics = direct.Slice_ToProto(mapCtx, in.Metrics, TimeSeries_ToProto)
	// MISSING: TaskResult
	out.TotalProcessingErrorCount = direct.ValueOf(in.TotalProcessingErrorCount)
	out.TotalResourceErrorCount = direct.ValueOf(in.TotalResourceErrorCount)
	return out
}
func MigrationTaskResult_FromProto(mapCtx *direct.MapContext, in *pb.MigrationTaskResult) *krm.MigrationTaskResult {
	if in == nil {
		return nil
	}
	out := &krm.MigrationTaskResult{}
	out.TranslationTaskResult = TranslationTaskResult_FromProto(mapCtx, in.GetTranslationTaskResult())
	return out
}
func MigrationTaskResult_ToProto(mapCtx *direct.MapContext, in *krm.MigrationTaskResult) *pb.MigrationTaskResult {
	if in == nil {
		return nil
	}
	out := &pb.MigrationTaskResult{}
	if oneof := TranslationTaskResult_ToProto(mapCtx, in.TranslationTaskResult); oneof != nil {
		out.Details = &pb.MigrationTaskResult_TranslationTaskResult{TranslationTaskResult: oneof}
	}
	return out
}
func MigrationWorkflow_FromProto(mapCtx *direct.MapContext, in *pb.MigrationWorkflow) *krm.MigrationWorkflow {
	if in == nil {
		return nil
	}
	out := &krm.MigrationWorkflow{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Tasks
	// MISSING: State
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.LastUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastUpdateTime())
	return out
}
func MigrationWorkflow_ToProto(mapCtx *direct.MapContext, in *krm.MigrationWorkflow) *pb.MigrationWorkflow {
	if in == nil {
		return nil
	}
	out := &pb.MigrationWorkflow{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Tasks
	// MISSING: State
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.LastUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.LastUpdateTime)
	return out
}
func MigrationWorkflowObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationWorkflow) *krm.MigrationWorkflowObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationWorkflowObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Tasks
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	return out
}
func MigrationWorkflowObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationWorkflowObservedState) *pb.MigrationWorkflow {
	if in == nil {
		return nil
	}
	out := &pb.MigrationWorkflow{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Tasks
	out.State = direct.Enum_ToProto[pb.MigrationWorkflow_State](mapCtx, in.State)
	// MISSING: CreateTime
	// MISSING: LastUpdateTime
	return out
}
func MySQLDialect_FromProto(mapCtx *direct.MapContext, in *pb.MySQLDialect) *krm.MySQLDialect {
	if in == nil {
		return nil
	}
	out := &krm.MySQLDialect{}
	return out
}
func MySQLDialect_ToProto(mapCtx *direct.MapContext, in *krm.MySQLDialect) *pb.MySQLDialect {
	if in == nil {
		return nil
	}
	out := &pb.MySQLDialect{}
	return out
}
func NameMappingKey_FromProto(mapCtx *direct.MapContext, in *pb.NameMappingKey) *krm.NameMappingKey {
	if in == nil {
		return nil
	}
	out := &krm.NameMappingKey{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.Relation = direct.LazyPtr(in.GetRelation())
	out.Attribute = direct.LazyPtr(in.GetAttribute())
	return out
}
func NameMappingKey_ToProto(mapCtx *direct.MapContext, in *krm.NameMappingKey) *pb.NameMappingKey {
	if in == nil {
		return nil
	}
	out := &pb.NameMappingKey{}
	out.Type = direct.Enum_ToProto[pb.NameMappingKey_Type](mapCtx, in.Type)
	out.Database = direct.ValueOf(in.Database)
	out.Schema = direct.ValueOf(in.Schema)
	out.Relation = direct.ValueOf(in.Relation)
	out.Attribute = direct.ValueOf(in.Attribute)
	return out
}
func NameMappingValue_FromProto(mapCtx *direct.MapContext, in *pb.NameMappingValue) *krm.NameMappingValue {
	if in == nil {
		return nil
	}
	out := &krm.NameMappingValue{}
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.Relation = direct.LazyPtr(in.GetRelation())
	out.Attribute = direct.LazyPtr(in.GetAttribute())
	return out
}
func NameMappingValue_ToProto(mapCtx *direct.MapContext, in *krm.NameMappingValue) *pb.NameMappingValue {
	if in == nil {
		return nil
	}
	out := &pb.NameMappingValue{}
	out.Database = direct.ValueOf(in.Database)
	out.Schema = direct.ValueOf(in.Schema)
	out.Relation = direct.ValueOf(in.Relation)
	out.Attribute = direct.ValueOf(in.Attribute)
	return out
}
func NetezzaDialect_FromProto(mapCtx *direct.MapContext, in *pb.NetezzaDialect) *krm.NetezzaDialect {
	if in == nil {
		return nil
	}
	out := &krm.NetezzaDialect{}
	return out
}
func NetezzaDialect_ToProto(mapCtx *direct.MapContext, in *krm.NetezzaDialect) *pb.NetezzaDialect {
	if in == nil {
		return nil
	}
	out := &pb.NetezzaDialect{}
	return out
}
func ObjectNameMapping_FromProto(mapCtx *direct.MapContext, in *pb.ObjectNameMapping) *krm.ObjectNameMapping {
	if in == nil {
		return nil
	}
	out := &krm.ObjectNameMapping{}
	out.Source = NameMappingKey_FromProto(mapCtx, in.GetSource())
	out.Target = NameMappingValue_FromProto(mapCtx, in.GetTarget())
	return out
}
func ObjectNameMapping_ToProto(mapCtx *direct.MapContext, in *krm.ObjectNameMapping) *pb.ObjectNameMapping {
	if in == nil {
		return nil
	}
	out := &pb.ObjectNameMapping{}
	out.Source = NameMappingKey_ToProto(mapCtx, in.Source)
	out.Target = NameMappingValue_ToProto(mapCtx, in.Target)
	return out
}
func ObjectNameMappingList_FromProto(mapCtx *direct.MapContext, in *pb.ObjectNameMappingList) *krm.ObjectNameMappingList {
	if in == nil {
		return nil
	}
	out := &krm.ObjectNameMappingList{}
	out.NameMap = direct.Slice_FromProto(mapCtx, in.NameMap, ObjectNameMapping_FromProto)
	return out
}
func ObjectNameMappingList_ToProto(mapCtx *direct.MapContext, in *krm.ObjectNameMappingList) *pb.ObjectNameMappingList {
	if in == nil {
		return nil
	}
	out := &pb.ObjectNameMappingList{}
	out.NameMap = direct.Slice_ToProto(mapCtx, in.NameMap, ObjectNameMapping_ToProto)
	return out
}
func OracleDialect_FromProto(mapCtx *direct.MapContext, in *pb.OracleDialect) *krm.OracleDialect {
	if in == nil {
		return nil
	}
	out := &krm.OracleDialect{}
	return out
}
func OracleDialect_ToProto(mapCtx *direct.MapContext, in *krm.OracleDialect) *pb.OracleDialect {
	if in == nil {
		return nil
	}
	out := &pb.OracleDialect{}
	return out
}
func Point_FromProto(mapCtx *direct.MapContext, in *pb.Point) *krm.Point {
	if in == nil {
		return nil
	}
	out := &krm.Point{}
	out.Interval = TimeInterval_FromProto(mapCtx, in.GetInterval())
	out.Value = TypedValue_FromProto(mapCtx, in.GetValue())
	return out
}
func Point_ToProto(mapCtx *direct.MapContext, in *krm.Point) *pb.Point {
	if in == nil {
		return nil
	}
	out := &pb.Point{}
	out.Interval = TimeInterval_ToProto(mapCtx, in.Interval)
	out.Value = TypedValue_ToProto(mapCtx, in.Value)
	return out
}
func PostgresqlDialect_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlDialect) *krm.PostgresqlDialect {
	if in == nil {
		return nil
	}
	out := &krm.PostgresqlDialect{}
	return out
}
func PostgresqlDialect_ToProto(mapCtx *direct.MapContext, in *krm.PostgresqlDialect) *pb.PostgresqlDialect {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlDialect{}
	return out
}
func PrestoDialect_FromProto(mapCtx *direct.MapContext, in *pb.PrestoDialect) *krm.PrestoDialect {
	if in == nil {
		return nil
	}
	out := &krm.PrestoDialect{}
	return out
}
func PrestoDialect_ToProto(mapCtx *direct.MapContext, in *krm.PrestoDialect) *pb.PrestoDialect {
	if in == nil {
		return nil
	}
	out := &pb.PrestoDialect{}
	return out
}
func RedshiftDialect_FromProto(mapCtx *direct.MapContext, in *pb.RedshiftDialect) *krm.RedshiftDialect {
	if in == nil {
		return nil
	}
	out := &krm.RedshiftDialect{}
	return out
}
func RedshiftDialect_ToProto(mapCtx *direct.MapContext, in *krm.RedshiftDialect) *pb.RedshiftDialect {
	if in == nil {
		return nil
	}
	out := &pb.RedshiftDialect{}
	return out
}
func ResourceErrorDetail_FromProto(mapCtx *direct.MapContext, in *pb.ResourceErrorDetail) *krm.ResourceErrorDetail {
	if in == nil {
		return nil
	}
	out := &krm.ResourceErrorDetail{}
	out.ResourceInfo = ResourceInfo_FromProto(mapCtx, in.GetResourceInfo())
	out.ErrorDetails = direct.Slice_FromProto(mapCtx, in.ErrorDetails, ErrorDetail_FromProto)
	out.ErrorCount = direct.LazyPtr(in.GetErrorCount())
	return out
}
func ResourceErrorDetail_ToProto(mapCtx *direct.MapContext, in *krm.ResourceErrorDetail) *pb.ResourceErrorDetail {
	if in == nil {
		return nil
	}
	out := &pb.ResourceErrorDetail{}
	out.ResourceInfo = ResourceInfo_ToProto(mapCtx, in.ResourceInfo)
	out.ErrorDetails = direct.Slice_ToProto(mapCtx, in.ErrorDetails, ErrorDetail_ToProto)
	out.ErrorCount = direct.ValueOf(in.ErrorCount)
	return out
}
func SQLServerDialect_FromProto(mapCtx *direct.MapContext, in *pb.SQLServerDialect) *krm.SQLServerDialect {
	if in == nil {
		return nil
	}
	out := &krm.SQLServerDialect{}
	return out
}
func SQLServerDialect_ToProto(mapCtx *direct.MapContext, in *krm.SQLServerDialect) *pb.SQLServerDialect {
	if in == nil {
		return nil
	}
	out := &pb.SQLServerDialect{}
	return out
}
func SQLiteDialect_FromProto(mapCtx *direct.MapContext, in *pb.SQLiteDialect) *krm.SQLiteDialect {
	if in == nil {
		return nil
	}
	out := &krm.SQLiteDialect{}
	return out
}
func SQLiteDialect_ToProto(mapCtx *direct.MapContext, in *krm.SQLiteDialect) *pb.SQLiteDialect {
	if in == nil {
		return nil
	}
	out := &pb.SQLiteDialect{}
	return out
}
func SnowflakeDialect_FromProto(mapCtx *direct.MapContext, in *pb.SnowflakeDialect) *krm.SnowflakeDialect {
	if in == nil {
		return nil
	}
	out := &krm.SnowflakeDialect{}
	return out
}
func SnowflakeDialect_ToProto(mapCtx *direct.MapContext, in *krm.SnowflakeDialect) *pb.SnowflakeDialect {
	if in == nil {
		return nil
	}
	out := &pb.SnowflakeDialect{}
	return out
}
func SourceEnv_FromProto(mapCtx *direct.MapContext, in *pb.SourceEnv) *krm.SourceEnv {
	if in == nil {
		return nil
	}
	out := &krm.SourceEnv{}
	out.DefaultDatabase = direct.LazyPtr(in.GetDefaultDatabase())
	out.SchemaSearchPath = in.SchemaSearchPath
	out.MetadataStoreDataset = direct.LazyPtr(in.GetMetadataStoreDataset())
	return out
}
func SourceEnv_ToProto(mapCtx *direct.MapContext, in *krm.SourceEnv) *pb.SourceEnv {
	if in == nil {
		return nil
	}
	out := &pb.SourceEnv{}
	out.DefaultDatabase = direct.ValueOf(in.DefaultDatabase)
	out.SchemaSearchPath = in.SchemaSearchPath
	out.MetadataStoreDataset = direct.ValueOf(in.MetadataStoreDataset)
	return out
}
func SourceEnvironment_FromProto(mapCtx *direct.MapContext, in *pb.SourceEnvironment) *krm.SourceEnvironment {
	if in == nil {
		return nil
	}
	out := &krm.SourceEnvironment{}
	out.DefaultDatabase = direct.LazyPtr(in.GetDefaultDatabase())
	out.SchemaSearchPath = in.SchemaSearchPath
	out.MetadataStoreDataset = direct.LazyPtr(in.GetMetadataStoreDataset())
	return out
}
func SourceEnvironment_ToProto(mapCtx *direct.MapContext, in *krm.SourceEnvironment) *pb.SourceEnvironment {
	if in == nil {
		return nil
	}
	out := &pb.SourceEnvironment{}
	out.DefaultDatabase = direct.ValueOf(in.DefaultDatabase)
	out.SchemaSearchPath = in.SchemaSearchPath
	out.MetadataStoreDataset = direct.ValueOf(in.MetadataStoreDataset)
	return out
}
func SourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.SourceSpec) *krm.SourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.SourceSpec{}
	out.BaseURI = direct.LazyPtr(in.GetBaseUri())
	out.Literal = Literal_FromProto(mapCtx, in.GetLiteral())
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	return out
}
func SourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.SourceSpec) *pb.SourceSpec {
	if in == nil {
		return nil
	}
	out := &pb.SourceSpec{}
	if oneof := SourceSpec_BaseUri_ToProto(mapCtx, in.BaseURI); oneof != nil {
		out.Source = oneof
	}
	if oneof := Literal_ToProto(mapCtx, in.Literal); oneof != nil {
		out.Source = &pb.SourceSpec_Literal{Literal: oneof}
	}
	out.Encoding = direct.ValueOf(in.Encoding)
	return out
}
func SourceTargetMapping_FromProto(mapCtx *direct.MapContext, in *pb.SourceTargetMapping) *krm.SourceTargetMapping {
	if in == nil {
		return nil
	}
	out := &krm.SourceTargetMapping{}
	out.SourceSpec = SourceSpec_FromProto(mapCtx, in.GetSourceSpec())
	out.TargetSpec = TargetSpec_FromProto(mapCtx, in.GetTargetSpec())
	return out
}
func SourceTargetMapping_ToProto(mapCtx *direct.MapContext, in *krm.SourceTargetMapping) *pb.SourceTargetMapping {
	if in == nil {
		return nil
	}
	out := &pb.SourceTargetMapping{}
	out.SourceSpec = SourceSpec_ToProto(mapCtx, in.SourceSpec)
	out.TargetSpec = TargetSpec_ToProto(mapCtx, in.TargetSpec)
	return out
}
func SparkSQLDialect_FromProto(mapCtx *direct.MapContext, in *pb.SparkSQLDialect) *krm.SparkSQLDialect {
	if in == nil {
		return nil
	}
	out := &krm.SparkSQLDialect{}
	return out
}
func SparkSQLDialect_ToProto(mapCtx *direct.MapContext, in *krm.SparkSQLDialect) *pb.SparkSQLDialect {
	if in == nil {
		return nil
	}
	out := &pb.SparkSQLDialect{}
	return out
}
func TargetSpec_FromProto(mapCtx *direct.MapContext, in *pb.TargetSpec) *krm.TargetSpec {
	if in == nil {
		return nil
	}
	out := &krm.TargetSpec{}
	out.RelativePath = direct.LazyPtr(in.GetRelativePath())
	return out
}
func TargetSpec_ToProto(mapCtx *direct.MapContext, in *krm.TargetSpec) *pb.TargetSpec {
	if in == nil {
		return nil
	}
	out := &pb.TargetSpec{}
	out.RelativePath = direct.ValueOf(in.RelativePath)
	return out
}
func TeradataDialect_FromProto(mapCtx *direct.MapContext, in *pb.TeradataDialect) *krm.TeradataDialect {
	if in == nil {
		return nil
	}
	out := &krm.TeradataDialect{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func TeradataDialect_ToProto(mapCtx *direct.MapContext, in *krm.TeradataDialect) *pb.TeradataDialect {
	if in == nil {
		return nil
	}
	out := &pb.TeradataDialect{}
	out.Mode = direct.Enum_ToProto[pb.TeradataDialect_Mode](mapCtx, in.Mode)
	return out
}
func TimeInterval_FromProto(mapCtx *direct.MapContext, in *pb.TimeInterval) *krm.TimeInterval {
	if in == nil {
		return nil
	}
	out := &krm.TimeInterval{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func TimeInterval_ToProto(mapCtx *direct.MapContext, in *krm.TimeInterval) *pb.TimeInterval {
	if in == nil {
		return nil
	}
	out := &pb.TimeInterval{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func TimeSeries_FromProto(mapCtx *direct.MapContext, in *pb.TimeSeries) *krm.TimeSeries {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeries{}
	out.Metric = direct.LazyPtr(in.GetMetric())
	out.ValueType = direct.Enum_FromProto(mapCtx, in.GetValueType())
	out.MetricKind = direct.Enum_FromProto(mapCtx, in.GetMetricKind())
	out.Points = direct.Slice_FromProto(mapCtx, in.Points, Point_FromProto)
	return out
}
func TimeSeries_ToProto(mapCtx *direct.MapContext, in *krm.TimeSeries) *pb.TimeSeries {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeries{}
	out.Metric = direct.ValueOf(in.Metric)
	out.ValueType = direct.Enum_ToProto[pb.MetricDescriptor_ValueType](mapCtx, in.ValueType)
	out.MetricKind = direct.Enum_ToProto[pb.MetricDescriptor_MetricKind](mapCtx, in.MetricKind)
	out.Points = direct.Slice_ToProto(mapCtx, in.Points, Point_ToProto)
	return out
}
func TranslationConfigDetails_FromProto(mapCtx *direct.MapContext, in *pb.TranslationConfigDetails) *krm.TranslationConfigDetails {
	if in == nil {
		return nil
	}
	out := &krm.TranslationConfigDetails{}
	out.GcsSourcePath = direct.LazyPtr(in.GetGcsSourcePath())
	out.GcsTargetPath = direct.LazyPtr(in.GetGcsTargetPath())
	out.NameMappingList = ObjectNameMappingList_FromProto(mapCtx, in.GetNameMappingList())
	out.SourceDialect = Dialect_FromProto(mapCtx, in.GetSourceDialect())
	out.TargetDialect = Dialect_FromProto(mapCtx, in.GetTargetDialect())
	out.SourceEnv = SourceEnv_FromProto(mapCtx, in.GetSourceEnv())
	out.RequestSource = direct.LazyPtr(in.GetRequestSource())
	out.TargetTypes = in.TargetTypes
	return out
}
func TranslationConfigDetails_ToProto(mapCtx *direct.MapContext, in *krm.TranslationConfigDetails) *pb.TranslationConfigDetails {
	if in == nil {
		return nil
	}
	out := &pb.TranslationConfigDetails{}
	if oneof := TranslationConfigDetails_GcsSourcePath_ToProto(mapCtx, in.GcsSourcePath); oneof != nil {
		out.SourceLocation = oneof
	}
	if oneof := TranslationConfigDetails_GcsTargetPath_ToProto(mapCtx, in.GcsTargetPath); oneof != nil {
		out.TargetLocation = oneof
	}
	if oneof := ObjectNameMappingList_ToProto(mapCtx, in.NameMappingList); oneof != nil {
		out.OutputNameMapping = &pb.TranslationConfigDetails_NameMappingList{NameMappingList: oneof}
	}
	out.SourceDialect = Dialect_ToProto(mapCtx, in.SourceDialect)
	out.TargetDialect = Dialect_ToProto(mapCtx, in.TargetDialect)
	out.SourceEnv = SourceEnv_ToProto(mapCtx, in.SourceEnv)
	out.RequestSource = direct.ValueOf(in.RequestSource)
	out.TargetTypes = in.TargetTypes
	return out
}
func TranslationDetails_FromProto(mapCtx *direct.MapContext, in *pb.TranslationDetails) *krm.TranslationDetails {
	if in == nil {
		return nil
	}
	out := &krm.TranslationDetails{}
	out.SourceTargetMapping = direct.Slice_FromProto(mapCtx, in.SourceTargetMapping, SourceTargetMapping_FromProto)
	out.TargetBaseURI = direct.LazyPtr(in.GetTargetBaseUri())
	out.SourceEnvironment = SourceEnvironment_FromProto(mapCtx, in.GetSourceEnvironment())
	out.TargetReturnLiterals = in.TargetReturnLiterals
	out.TargetTypes = in.TargetTypes
	return out
}
func TranslationDetails_ToProto(mapCtx *direct.MapContext, in *krm.TranslationDetails) *pb.TranslationDetails {
	if in == nil {
		return nil
	}
	out := &pb.TranslationDetails{}
	out.SourceTargetMapping = direct.Slice_ToProto(mapCtx, in.SourceTargetMapping, SourceTargetMapping_ToProto)
	out.TargetBaseUri = direct.ValueOf(in.TargetBaseURI)
	out.SourceEnvironment = SourceEnvironment_ToProto(mapCtx, in.SourceEnvironment)
	out.TargetReturnLiterals = in.TargetReturnLiterals
	out.TargetTypes = in.TargetTypes
	return out
}
func TranslationTaskResult_FromProto(mapCtx *direct.MapContext, in *pb.TranslationTaskResult) *krm.TranslationTaskResult {
	if in == nil {
		return nil
	}
	out := &krm.TranslationTaskResult{}
	out.TranslatedLiterals = direct.Slice_FromProto(mapCtx, in.TranslatedLiterals, Literal_FromProto)
	out.ReportLogMessages = direct.Slice_FromProto(mapCtx, in.ReportLogMessages, GcsReportLogMessage_FromProto)
	return out
}
func TranslationTaskResult_ToProto(mapCtx *direct.MapContext, in *krm.TranslationTaskResult) *pb.TranslationTaskResult {
	if in == nil {
		return nil
	}
	out := &pb.TranslationTaskResult{}
	out.TranslatedLiterals = direct.Slice_ToProto(mapCtx, in.TranslatedLiterals, Literal_ToProto)
	out.ReportLogMessages = direct.Slice_ToProto(mapCtx, in.ReportLogMessages, GcsReportLogMessage_ToProto)
	return out
}
func TypedValue_FromProto(mapCtx *direct.MapContext, in *pb.TypedValue) *krm.TypedValue {
	if in == nil {
		return nil
	}
	out := &krm.TypedValue{}
	out.BoolValue = direct.LazyPtr(in.GetBoolValue())
	out.Int64Value = direct.LazyPtr(in.GetInt64Value())
	out.DoubleValue = direct.LazyPtr(in.GetDoubleValue())
	out.StringValue = direct.LazyPtr(in.GetStringValue())
	out.DistributionValue = Distribution_FromProto(mapCtx, in.GetDistributionValue())
	return out
}
func TypedValue_ToProto(mapCtx *direct.MapContext, in *krm.TypedValue) *pb.TypedValue {
	if in == nil {
		return nil
	}
	out := &pb.TypedValue{}
	if oneof := TypedValue_BoolValue_ToProto(mapCtx, in.BoolValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := TypedValue_Int64Value_ToProto(mapCtx, in.Int64Value); oneof != nil {
		out.Value = oneof
	}
	if oneof := TypedValue_DoubleValue_ToProto(mapCtx, in.DoubleValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := TypedValue_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := Distribution_ToProto(mapCtx, in.DistributionValue); oneof != nil {
		out.Value = &pb.TypedValue_DistributionValue{DistributionValue: oneof}
	}
	return out
}
func VerticaDialect_FromProto(mapCtx *direct.MapContext, in *pb.VerticaDialect) *krm.VerticaDialect {
	if in == nil {
		return nil
	}
	out := &krm.VerticaDialect{}
	return out
}
func VerticaDialect_ToProto(mapCtx *direct.MapContext, in *krm.VerticaDialect) *pb.VerticaDialect {
	if in == nil {
		return nil
	}
	out := &pb.VerticaDialect{}
	return out
}
