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

	bigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate) *krmv1alpha1.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.DatasetIDPrefix = direct.LazyPtr(in.GetDatasetIdPrefix())
	if in.GetKmsKeyName() != "" {
		out.KMSKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	return out
}
func BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate) *pb.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate{}
	out.Location = direct.ValueOf(in.Location)
	out.DatasetIdPrefix = direct.ValueOf(in.DatasetIDPrefix)
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}
func DestinationConfig_FromProto(mapCtx *direct.MapContext, in *pb.DestinationConfig) *krmv1alpha1.DestinationConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DestinationConfig{}
	if in.GetDestinationConnectionProfile() != "" {
		out.DestinationConnectionProfileRef = &krmv1alpha1.ConnectionProfileRef{External: in.GetDestinationConnectionProfile()}
	}
	out.GCSDestinationConfig = GCSDestinationConfig_FromProto(mapCtx, in.GetGcsDestinationConfig())
	out.BigqueryDestinationConfig = BigQueryDestinationConfig_FromProto(mapCtx, in.GetBigqueryDestinationConfig())
	return out
}
func MySQLDatabase_FromProto(mapCtx *direct.MapContext, in *pb.MysqlDatabase) *krmv1alpha1.MySQLDatabase {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MySQLDatabase{}
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.MySQLTables = direct.Slice_FromProto(mapCtx, in.GetMysqlTables(), MySQLTable_FromProto)
	return out
}
func MySQLDatabase_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MySQLDatabase) *pb.MysqlDatabase {
	if in == nil {
		return nil
	}
	out := &pb.MysqlDatabase{}
	out.Database = direct.ValueOf(in.Database)
	out.MysqlTables = direct.Slice_ToProto(mapCtx, in.MySQLTables, MySQLTable_ToProto)
	return out
}
func MySQLTable_FromProto(mapCtx *direct.MapContext, in *pb.MysqlTable) *krmv1alpha1.MySQLTable {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MySQLTable{}
	out.Table = direct.LazyPtr(in.GetTable())
	out.MySQLColumns = direct.Slice_FromProto(mapCtx, in.GetMysqlColumns(), MySQLColumn_FromProto)
	return out
}
func MySQLTable_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MySQLTable) *pb.MysqlTable {
	if in == nil {
		return nil
	}
	out := &pb.MysqlTable{}
	out.Table = direct.ValueOf(in.Table)
	out.MysqlColumns = direct.Slice_ToProto(mapCtx, in.MySQLColumns, MySQLColumn_ToProto)
	return out
}
func MysqlRdbms_FromProto(mapCtx *direct.MapContext, in *pb.MysqlRdbms) *krmv1alpha1.MysqlRdbms {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MysqlRdbms{}
	out.MySQLDatabases = direct.Slice_FromProto(mapCtx, in.GetMysqlDatabases(), MySQLDatabase_FromProto)
	return out
}
func MysqlRdbms_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MysqlRdbms) *pb.MysqlRdbms {
	if in == nil {
		return nil
	}
	out := &pb.MysqlRdbms{}
	out.MysqlDatabases = direct.Slice_ToProto(mapCtx, in.MySQLDatabases, MySQLDatabase_ToProto)
	return out
}
func PostgreSQLSchema_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlSchema) *krmv1alpha1.PostgreSQLSchema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PostgreSQLSchema{}
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.PostgreSQLTables = direct.Slice_FromProto(mapCtx, in.GetPostgresqlTables(), PostgreSQLTable_FromProto)
	return out
}
func PostgreSQLSchema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PostgreSQLSchema) *pb.PostgresqlSchema {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlSchema{}
	out.Schema = direct.ValueOf(in.Schema)
	out.PostgresqlTables = direct.Slice_ToProto(mapCtx, in.PostgreSQLTables, PostgreSQLTable_ToProto)
	return out
}
func PostgreSQLTable_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlTable) *krmv1alpha1.PostgreSQLTable {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PostgreSQLTable{}
	out.Table = direct.LazyPtr(in.GetTable())
	out.PostgreSQLColumns = direct.Slice_FromProto(mapCtx, in.GetPostgresqlColumns(), PostgreSQLColumn_FromProto)
	return out
}
func PostgreSQLTable_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PostgreSQLTable) *pb.PostgresqlTable {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlTable{}
	out.Table = direct.ValueOf(in.Table)
	out.PostgresqlColumns = direct.Slice_ToProto(mapCtx, in.PostgreSQLColumns, PostgreSQLColumn_ToProto)
	return out
}
func PostgresqlRdbms_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlRdbms) *krmv1alpha1.PostgresqlRdbms {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PostgresqlRdbms{}
	out.PostgreSQLSchemas = direct.Slice_FromProto(mapCtx, in.GetPostgresqlSchemas(), PostgreSQLSchema_FromProto)
	return out
}
func PostgresqlRdbms_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PostgresqlRdbms) *pb.PostgresqlRdbms {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlRdbms{}
	out.PostgresqlSchemas = direct.Slice_ToProto(mapCtx, in.PostgreSQLSchemas, PostgreSQLSchema_ToProto)
	return out
}
func SourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.SourceConfig) *krmv1alpha1.SourceConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SourceConfig{}
	if in.GetSourceConnectionProfile() != "" {
		out.SourceConnectionProfileRef = &krmv1alpha1.ConnectionProfileRef{External: in.GetSourceConnectionProfile()}
	}
	out.OracleSourceConfig = OracleSourceConfig_FromProto(mapCtx, in.GetOracleSourceConfig())
	out.MySQLSourceConfig = MysqlSourceConfig_FromProto(mapCtx, in.GetMysqlSourceConfig())
	out.PostgreSQLSourceConfig = PostgresqlSourceConfig_FromProto(mapCtx, in.GetPostgresqlSourceConfig())
	out.SQLServerSourceConfig = SQLServerSourceConfig_FromProto(mapCtx, in.GetSqlServerSourceConfig())
	return out
}
func SourceConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SourceConfig) *pb.SourceConfig {
	if in == nil {
		return nil
	}
	out := &pb.SourceConfig{}
	if in.SourceConnectionProfileRef != nil {
		out.SourceConnectionProfile = in.SourceConnectionProfileRef.External
	}
	if oneof := OracleSourceConfig_ToProto(mapCtx, in.OracleSourceConfig); oneof != nil {
		out.SourceStreamConfig = &pb.SourceConfig_OracleSourceConfig{OracleSourceConfig: oneof}
	}
	if oneof := MysqlSourceConfig_ToProto(mapCtx, in.MySQLSourceConfig); oneof != nil {
		out.SourceStreamConfig = &pb.SourceConfig_MysqlSourceConfig{MysqlSourceConfig: oneof}
	}
	if oneof := PostgresqlSourceConfig_ToProto(mapCtx, in.PostgreSQLSourceConfig); oneof != nil {
		out.SourceStreamConfig = &pb.SourceConfig_PostgresqlSourceConfig{PostgresqlSourceConfig: oneof}
	}
	if oneof := SQLServerSourceConfig_ToProto(mapCtx, in.SQLServerSourceConfig); oneof != nil {
		out.SourceStreamConfig = &pb.SourceConfig_SqlServerSourceConfig{SqlServerSourceConfig: oneof}
	}
	return out
}
func Stream_BackfillAllStrategy_FromProto(mapCtx *direct.MapContext, in *pb.Stream_BackfillAllStrategy) *krmv1alpha1.Stream_BackfillAllStrategy {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Stream_BackfillAllStrategy{}
	out.OracleExcludedObjects = OracleRdbms_FromProto(mapCtx, in.GetOracleExcludedObjects())
	out.MySQLExcludedObjects = MysqlRdbms_FromProto(mapCtx, in.GetMysqlExcludedObjects())
	out.PostgreSQLExcludedObjects = PostgresqlRdbms_FromProto(mapCtx, in.GetPostgresqlExcludedObjects())
	out.SQLServerExcludedObjects = SQLServerRdbms_FromProto(mapCtx, in.GetSqlServerExcludedObjects())
	return out
}
func Stream_BackfillAllStrategy_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Stream_BackfillAllStrategy) *pb.Stream_BackfillAllStrategy {
	if in == nil {
		return nil
	}
	out := &pb.Stream_BackfillAllStrategy{}
	if oneof := OracleRdbms_ToProto(mapCtx, in.OracleExcludedObjects); oneof != nil {
		out.ExcludedObjects = &pb.Stream_BackfillAllStrategy_OracleExcludedObjects{OracleExcludedObjects: oneof}
	}
	if oneof := MysqlRdbms_ToProto(mapCtx, in.MySQLExcludedObjects); oneof != nil {
		out.ExcludedObjects = &pb.Stream_BackfillAllStrategy_MysqlExcludedObjects{MysqlExcludedObjects: oneof}
	}
	if oneof := PostgresqlRdbms_ToProto(mapCtx, in.PostgreSQLExcludedObjects); oneof != nil {
		out.ExcludedObjects = &pb.Stream_BackfillAllStrategy_PostgresqlExcludedObjects{PostgresqlExcludedObjects: oneof}
	}
	if oneof := SQLServerRdbms_ToProto(mapCtx, in.SQLServerExcludedObjects); oneof != nil {
		out.ExcludedObjects = &pb.Stream_BackfillAllStrategy_SqlServerExcludedObjects{SqlServerExcludedObjects: oneof}
	}
	return out
}
func BigQueryDestinationConfig_SingleTargetDataset_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig_SingleTargetDataset) *krmv1alpha1.BigQueryDestinationConfig_SingleTargetDataset {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryDestinationConfig_SingleTargetDataset{}
	if in.GetDatasetId() != "" {
		out.DatasetRef = &bigqueryv1beta1.DatasetRef{External: in.GetDatasetId()}
	}
	return out
}
func BigQueryDestinationConfig_SingleTargetDataset_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryDestinationConfig_SingleTargetDataset) *pb.BigQueryDestinationConfig_SingleTargetDataset {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_SingleTargetDataset{}
	if in.DatasetRef != nil {
		out.DatasetId = in.DatasetRef.External
	}
	return out
}
