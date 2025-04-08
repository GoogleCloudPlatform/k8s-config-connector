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

// +generated:mapper
// krm.group: datastream.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.datastream.v1

package datastream

import (
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AvroFileFormat_FromProto(mapCtx *direct.MapContext, in *pb.AvroFileFormat) *krmv1alpha1.AvroFileFormat {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.AvroFileFormat{}
	return out
}
func AvroFileFormat_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.AvroFileFormat) *pb.AvroFileFormat {
	if in == nil {
		return nil
	}
	out := &pb.AvroFileFormat{}
	return out
}
func BigQueryDestinationConfig_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig) *krmv1alpha1.BigQueryDestinationConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryDestinationConfig{}
	out.SingleTargetDataset = BigQueryDestinationConfig_SingleTargetDataset_FromProto(mapCtx, in.GetSingleTargetDataset())
	out.SourceHierarchyDatasets = BigQueryDestinationConfig_SourceHierarchyDatasets_FromProto(mapCtx, in.GetSourceHierarchyDatasets())
	out.DataFreshness = direct.StringDuration_FromProto(mapCtx, in.GetDataFreshness())
	out.Merge = BigQueryDestinationConfig_Merge_FromProto(mapCtx, in.GetMerge())
	out.AppendOnly = BigQueryDestinationConfig_AppendOnly_FromProto(mapCtx, in.GetAppendOnly())
	return out
}
func BigQueryDestinationConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryDestinationConfig) *pb.BigQueryDestinationConfig {
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
func BigQueryDestinationConfig_AppendOnly_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig_AppendOnly) *krmv1alpha1.BigQueryDestinationConfig_AppendOnly {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryDestinationConfig_AppendOnly{}
	return out
}
func BigQueryDestinationConfig_AppendOnly_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryDestinationConfig_AppendOnly) *pb.BigQueryDestinationConfig_AppendOnly {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_AppendOnly{}
	return out
}
func BigQueryDestinationConfig_Merge_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig_Merge) *krmv1alpha1.BigQueryDestinationConfig_Merge {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryDestinationConfig_Merge{}
	return out
}
func BigQueryDestinationConfig_Merge_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryDestinationConfig_Merge) *pb.BigQueryDestinationConfig_Merge {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_Merge{}
	return out
}
func BigQueryDestinationConfig_SingleTargetDataset_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig_SingleTargetDataset) *krmv1alpha1.BigQueryDestinationConfig_SingleTargetDataset {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryDestinationConfig_SingleTargetDataset{}
	out.DatasetID = direct.LazyPtr(in.GetDatasetId())
	return out
}
func BigQueryDestinationConfig_SingleTargetDataset_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryDestinationConfig_SingleTargetDataset) *pb.BigQueryDestinationConfig_SingleTargetDataset {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_SingleTargetDataset{}
	out.DatasetId = direct.ValueOf(in.DatasetID)
	return out
}
func BigQueryDestinationConfig_SourceHierarchyDatasets_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig_SourceHierarchyDatasets) *krmv1alpha1.BigQueryDestinationConfig_SourceHierarchyDatasets {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryDestinationConfig_SourceHierarchyDatasets{}
	out.DatasetTemplate = BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate_FromProto(mapCtx, in.GetDatasetTemplate())
	return out
}
func BigQueryDestinationConfig_SourceHierarchyDatasets_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryDestinationConfig_SourceHierarchyDatasets) *pb.BigQueryDestinationConfig_SourceHierarchyDatasets {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_SourceHierarchyDatasets{}
	out.DatasetTemplate = BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate_ToProto(mapCtx, in.DatasetTemplate)
	return out
}
func BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate) *krmv1alpha1.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.DatasetIDPrefix = direct.LazyPtr(in.GetDatasetIdPrefix())
	// MISSING: KMSKeyName
	return out
}
func BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate) *pb.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate{}
	out.Location = direct.ValueOf(in.Location)
	out.DatasetIdPrefix = direct.ValueOf(in.DatasetIDPrefix)
	// MISSING: KMSKeyName
	return out
}
func BigQueryProfile_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryProfile) *krmv1alpha1.BigQueryProfile {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryProfile{}
	return out
}
func BigQueryProfile_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryProfile) *pb.BigQueryProfile {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryProfile{}
	return out
}
func DatastreamPrivateConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krmv1alpha1.DatastreamPrivateConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DatastreamPrivateConnectionObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Error_FromProto(mapCtx, in.GetError())
	// MISSING: VpcPeeringConfig
	return out
}
func DatastreamPrivateConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DatastreamPrivateConnectionObservedState) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.PrivateConnection_State](mapCtx, in.State)
	out.Error = Error_ToProto(mapCtx, in.Error)
	// MISSING: VpcPeeringConfig
	return out
}
func DatastreamRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krmv1alpha1.DatastreamRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DatastreamRouteObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func DatastreamRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DatastreamRouteObservedState) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func DatastreamRouteSpec_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krmv1alpha1.DatastreamRouteSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DatastreamRouteSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.DestinationAddress = direct.LazyPtr(in.GetDestinationAddress())
	out.DestinationPort = direct.LazyPtr(in.GetDestinationPort())
	return out
}
func DatastreamRouteSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DatastreamRouteSpec) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	// MISSING: Name
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.DestinationAddress = direct.ValueOf(in.DestinationAddress)
	out.DestinationPort = direct.ValueOf(in.DestinationPort)
	return out
}
func DatastreamStreamObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krmv1alpha1.DatastreamStreamObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DatastreamStreamObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, Error_FromProto)
	out.LastRecoveryTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastRecoveryTime())
	return out
}
func DatastreamStreamObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DatastreamStreamObservedState) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, Error_ToProto)
	out.LastRecoveryTime = direct.StringTimestamp_ToProto(mapCtx, in.LastRecoveryTime)
	return out
}
func DatastreamStreamSpec_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krmv1alpha1.DatastreamStreamSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DatastreamStreamSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SourceConfig = SourceConfig_FromProto(mapCtx, in.GetSourceConfig())
	out.DestinationConfig = DestinationConfig_FromProto(mapCtx, in.GetDestinationConfig())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.BackfillAll = Stream_BackfillAllStrategy_FromProto(mapCtx, in.GetBackfillAll())
	out.BackfillNone = Stream_BackfillNoneStrategy_FromProto(mapCtx, in.GetBackfillNone())
	out.CustomerManagedEncryptionKey = in.CustomerManagedEncryptionKey
	return out
}
func DatastreamStreamSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DatastreamStreamSpec) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	// MISSING: Name
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
	out.CustomerManagedEncryptionKey = in.CustomerManagedEncryptionKey
	return out
}
func DestinationConfig_FromProto(mapCtx *direct.MapContext, in *pb.DestinationConfig) *krmv1alpha1.DestinationConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DestinationConfig{}
	if in.GetDestinationConnectionProfile() != "" {
		out.DestinationConnectionProfileRef = &refs.*ConnectionProfileRef{External: in.GetDestinationConnectionProfile()}
	}
	out.GCSDestinationConfig = GCSDestinationConfig_FromProto(mapCtx, in.GetGcsDestinationConfig())
	out.BigqueryDestinationConfig = BigQueryDestinationConfig_FromProto(mapCtx, in.GetBigqueryDestinationConfig())
	return out
}
func DestinationConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DestinationConfig) *pb.DestinationConfig {
	if in == nil {
		return nil
	}
	out := &pb.DestinationConfig{}
	if in.DestinationConnectionProfileRef != nil {
		out.DestinationConnectionProfile = in.DestinationConnectionProfileRef.External
	}
	if oneof := GCSDestinationConfig_ToProto(mapCtx, in.GCSDestinationConfig); oneof != nil {
		out.DestinationStreamConfig = &pb.DestinationConfig_GcsDestinationConfig{GcsDestinationConfig: oneof}
	}
	if oneof := BigQueryDestinationConfig_ToProto(mapCtx, in.BigqueryDestinationConfig); oneof != nil {
		out.DestinationStreamConfig = &pb.DestinationConfig_BigqueryDestinationConfig{BigqueryDestinationConfig: oneof}
	}
	return out
}
func GCSDestinationConfig_FromProto(mapCtx *direct.MapContext, in *pb.GcsDestinationConfig) *krmv1alpha1.GCSDestinationConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.GCSDestinationConfig{}
	out.Path = direct.LazyPtr(in.GetPath())
	out.FileRotationMb = direct.LazyPtr(in.GetFileRotationMb())
	out.FileRotationInterval = direct.StringDuration_FromProto(mapCtx, in.GetFileRotationInterval())
	out.AvroFileFormat = AvroFileFormat_FromProto(mapCtx, in.GetAvroFileFormat())
	out.JsonFileFormat = JsonFileFormat_FromProto(mapCtx, in.GetJsonFileFormat())
	return out
}
func GCSDestinationConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.GCSDestinationConfig) *pb.GcsDestinationConfig {
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
func GCSProfile_FromProto(mapCtx *direct.MapContext, in *pb.GcsProfile) *krmv1alpha1.GCSProfile {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.GCSProfile{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.RootPath = direct.LazyPtr(in.GetRootPath())
	return out
}
func GCSProfile_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.GCSProfile) *pb.GcsProfile {
	if in == nil {
		return nil
	}
	out := &pb.GcsProfile{}
	out.Bucket = direct.ValueOf(in.Bucket)
	out.RootPath = direct.ValueOf(in.RootPath)
	return out
}
func JsonFileFormat_FromProto(mapCtx *direct.MapContext, in *pb.JsonFileFormat) *krmv1alpha1.JsonFileFormat {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.JsonFileFormat{}
	out.SchemaFileFormat = direct.Enum_FromProto(mapCtx, in.GetSchemaFileFormat())
	out.Compression = direct.Enum_FromProto(mapCtx, in.GetCompression())
	return out
}
func JsonFileFormat_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.JsonFileFormat) *pb.JsonFileFormat {
	if in == nil {
		return nil
	}
	out := &pb.JsonFileFormat{}
	out.SchemaFileFormat = direct.Enum_ToProto[pb.JsonFileFormat_SchemaFileFormat](mapCtx, in.SchemaFileFormat)
	out.Compression = direct.Enum_ToProto[pb.JsonFileFormat_JsonCompression](mapCtx, in.Compression)
	return out
}
func MySQLColumn_FromProto(mapCtx *direct.MapContext, in *pb.MysqlColumn) *krmv1alpha1.MySQLColumn {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MySQLColumn{}
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
func MySQLColumn_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MySQLColumn) *pb.MysqlColumn {
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
func MySQLDatabase_FromProto(mapCtx *direct.MapContext, in *pb.MysqlDatabase) *krmv1alpha1.MySQLDatabase {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MySQLDatabase{}
	out.Database = direct.LazyPtr(in.GetDatabase())
	// MISSING: MysqlTables
	// (near miss): "MysqlTables" vs "MySQLTables"
	return out
}
func MySQLDatabase_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MySQLDatabase) *pb.MysqlDatabase {
	if in == nil {
		return nil
	}
	out := &pb.MysqlDatabase{}
	out.Database = direct.ValueOf(in.Database)
	// MISSING: MysqlTables
	// (near miss): "MysqlTables" vs "MySQLTables"
	return out
}
func MySQLTable_FromProto(mapCtx *direct.MapContext, in *pb.MysqlTable) *krmv1alpha1.MySQLTable {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MySQLTable{}
	out.Table = direct.LazyPtr(in.GetTable())
	// MISSING: MysqlColumns
	// (near miss): "MysqlColumns" vs "MySQLColumns"
	return out
}
func MySQLTable_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MySQLTable) *pb.MysqlTable {
	if in == nil {
		return nil
	}
	out := &pb.MysqlTable{}
	out.Table = direct.ValueOf(in.Table)
	// MISSING: MysqlColumns
	// (near miss): "MysqlColumns" vs "MySQLColumns"
	return out
}
func MysqlProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MysqlProfile) *krmv1alpha1.MysqlProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MysqlProfileObservedState{}
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	out.SSLConfig = MysqlSSLConfigObservedState_FromProto(mapCtx, in.GetSslConfig())
	return out
}
func MysqlProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MysqlProfileObservedState) *pb.MysqlProfile {
	if in == nil {
		return nil
	}
	out := &pb.MysqlProfile{}
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	out.SslConfig = MysqlSSLConfigObservedState_ToProto(mapCtx, in.SSLConfig)
	return out
}
func MysqlRdbms_FromProto(mapCtx *direct.MapContext, in *pb.MysqlRdbms) *krmv1alpha1.MysqlRdbms {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MysqlRdbms{}
	// MISSING: MysqlDatabases
	// (near miss): "MysqlDatabases" vs "MySQLDatabases"
	return out
}
func MysqlRdbms_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MysqlRdbms) *pb.MysqlRdbms {
	if in == nil {
		return nil
	}
	out := &pb.MysqlRdbms{}
	// MISSING: MysqlDatabases
	// (near miss): "MysqlDatabases" vs "MySQLDatabases"
	return out
}
func MysqlSSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSslConfig) *krmv1alpha1.MysqlSSLConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MysqlSSLConfig{}
	out.ClientKey = direct.LazyPtr(in.GetClientKey())
	// MISSING: ClientKeySet
	out.ClientCertificate = direct.LazyPtr(in.GetClientCertificate())
	// MISSING: ClientCertificateSet
	out.CACertificate = direct.LazyPtr(in.GetCaCertificate())
	// MISSING: CACertificateSet
	return out
}
func MysqlSSLConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MysqlSSLConfig) *pb.MysqlSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSslConfig{}
	out.ClientKey = direct.ValueOf(in.ClientKey)
	// MISSING: ClientKeySet
	out.ClientCertificate = direct.ValueOf(in.ClientCertificate)
	// MISSING: ClientCertificateSet
	out.CaCertificate = direct.ValueOf(in.CACertificate)
	// MISSING: CACertificateSet
	return out
}
func MysqlSSLConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSslConfig) *krmv1alpha1.MysqlSSLConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MysqlSSLConfigObservedState{}
	// MISSING: ClientKey
	out.ClientKeySet = direct.LazyPtr(in.GetClientKeySet())
	// MISSING: ClientCertificate
	out.ClientCertificateSet = direct.LazyPtr(in.GetClientCertificateSet())
	// MISSING: CACertificate
	out.CACertificateSet = direct.LazyPtr(in.GetCaCertificateSet())
	return out
}
func MysqlSSLConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MysqlSSLConfigObservedState) *pb.MysqlSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSslConfig{}
	// MISSING: ClientKey
	out.ClientKeySet = direct.ValueOf(in.ClientKeySet)
	// MISSING: ClientCertificate
	out.ClientCertificateSet = direct.ValueOf(in.ClientCertificateSet)
	// MISSING: CACertificate
	out.CaCertificateSet = direct.ValueOf(in.CACertificateSet)
	return out
}
func MysqlSourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSourceConfig) *krmv1alpha1.MysqlSourceConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MysqlSourceConfig{}
	out.IncludeObjects = MysqlRdbms_FromProto(mapCtx, in.GetIncludeObjects())
	out.ExcludeObjects = MysqlRdbms_FromProto(mapCtx, in.GetExcludeObjects())
	out.MaxConcurrentCdcTasks = direct.LazyPtr(in.GetMaxConcurrentCdcTasks())
	out.MaxConcurrentBackfillTasks = direct.LazyPtr(in.GetMaxConcurrentBackfillTasks())
	out.BinaryLogPosition = MysqlSourceConfig_BinaryLogPosition_FromProto(mapCtx, in.GetBinaryLogPosition())
	out.Gtid = MysqlSourceConfig_Gtid_FromProto(mapCtx, in.GetGtid())
	return out
}
func MysqlSourceConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MysqlSourceConfig) *pb.MysqlSourceConfig {
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
func MysqlSourceConfig_BinaryLogPosition_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSourceConfig_BinaryLogPosition) *krmv1alpha1.MysqlSourceConfig_BinaryLogPosition {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MysqlSourceConfig_BinaryLogPosition{}
	return out
}
func MysqlSourceConfig_BinaryLogPosition_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MysqlSourceConfig_BinaryLogPosition) *pb.MysqlSourceConfig_BinaryLogPosition {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSourceConfig_BinaryLogPosition{}
	return out
}
func MysqlSourceConfig_Gtid_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSourceConfig_Gtid) *krmv1alpha1.MysqlSourceConfig_Gtid {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MysqlSourceConfig_Gtid{}
	return out
}
func MysqlSourceConfig_Gtid_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MysqlSourceConfig_Gtid) *pb.MysqlSourceConfig_Gtid {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSourceConfig_Gtid{}
	return out
}
func OracleColumn_FromProto(mapCtx *direct.MapContext, in *pb.OracleColumn) *krmv1alpha1.OracleColumn {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleColumn{}
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
func OracleColumn_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleColumn) *pb.OracleColumn {
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
func OracleProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OracleProfile) *krmv1alpha1.OracleProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleProfileObservedState{}
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	// MISSING: DatabaseService
	// MISSING: ConnectionAttributes
	out.OracleSSLConfig = OracleSSLConfigObservedState_FromProto(mapCtx, in.GetOracleSslConfig())
	// MISSING: OracleAsmConfig
	// MISSING: SecretManagerStoredPassword
	return out
}
func OracleProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleProfileObservedState) *pb.OracleProfile {
	if in == nil {
		return nil
	}
	out := &pb.OracleProfile{}
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	// MISSING: DatabaseService
	// MISSING: ConnectionAttributes
	out.OracleSslConfig = OracleSSLConfigObservedState_ToProto(mapCtx, in.OracleSSLConfig)
	// MISSING: OracleAsmConfig
	// MISSING: SecretManagerStoredPassword
	return out
}
func OracleRdbms_FromProto(mapCtx *direct.MapContext, in *pb.OracleRdbms) *krmv1alpha1.OracleRdbms {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleRdbms{}
	out.OracleSchemas = direct.Slice_FromProto(mapCtx, in.OracleSchemas, OracleSchema_FromProto)
	return out
}
func OracleRdbms_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleRdbms) *pb.OracleRdbms {
	if in == nil {
		return nil
	}
	out := &pb.OracleRdbms{}
	out.OracleSchemas = direct.Slice_ToProto(mapCtx, in.OracleSchemas, OracleSchema_ToProto)
	return out
}
func OracleSSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.OracleSslConfig) *krmv1alpha1.OracleSSLConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleSSLConfig{}
	out.CACertificate = direct.LazyPtr(in.GetCaCertificate())
	// MISSING: CACertificateSet
	return out
}
func OracleSSLConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleSSLConfig) *pb.OracleSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.OracleSslConfig{}
	out.CaCertificate = direct.ValueOf(in.CACertificate)
	// MISSING: CACertificateSet
	return out
}
func OracleSSLConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OracleSslConfig) *krmv1alpha1.OracleSSLConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleSSLConfigObservedState{}
	// MISSING: CACertificate
	out.CACertificateSet = direct.LazyPtr(in.GetCaCertificateSet())
	return out
}
func OracleSSLConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleSSLConfigObservedState) *pb.OracleSslConfig {
	if in == nil {
		return nil
	}
	out := &pb.OracleSslConfig{}
	// MISSING: CACertificate
	out.CaCertificateSet = direct.ValueOf(in.CACertificateSet)
	return out
}
func OracleSchema_FromProto(mapCtx *direct.MapContext, in *pb.OracleSchema) *krmv1alpha1.OracleSchema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleSchema{}
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.OracleTables = direct.Slice_FromProto(mapCtx, in.OracleTables, OracleTable_FromProto)
	return out
}
func OracleSchema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleSchema) *pb.OracleSchema {
	if in == nil {
		return nil
	}
	out := &pb.OracleSchema{}
	out.Schema = direct.ValueOf(in.Schema)
	out.OracleTables = direct.Slice_ToProto(mapCtx, in.OracleTables, OracleTable_ToProto)
	return out
}
func OracleSourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig) *krmv1alpha1.OracleSourceConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleSourceConfig{}
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
func OracleSourceConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleSourceConfig) *pb.OracleSourceConfig {
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
func OracleSourceConfig_BinaryLogParser_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig_BinaryLogParser) *krmv1alpha1.OracleSourceConfig_BinaryLogParser {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleSourceConfig_BinaryLogParser{}
	out.OracleAsmLogFileAccess = OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess_FromProto(mapCtx, in.GetOracleAsmLogFileAccess())
	out.LogFileDirectories = OracleSourceConfig_BinaryLogParser_LogFileDirectories_FromProto(mapCtx, in.GetLogFileDirectories())
	return out
}
func OracleSourceConfig_BinaryLogParser_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleSourceConfig_BinaryLogParser) *pb.OracleSourceConfig_BinaryLogParser {
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
func OracleSourceConfig_BinaryLogParser_LogFileDirectories_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig_BinaryLogParser_LogFileDirectories) *krmv1alpha1.OracleSourceConfig_BinaryLogParser_LogFileDirectories {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleSourceConfig_BinaryLogParser_LogFileDirectories{}
	out.OnlineLogDirectory = direct.LazyPtr(in.GetOnlineLogDirectory())
	out.ArchivedLogDirectory = direct.LazyPtr(in.GetArchivedLogDirectory())
	return out
}
func OracleSourceConfig_BinaryLogParser_LogFileDirectories_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleSourceConfig_BinaryLogParser_LogFileDirectories) *pb.OracleSourceConfig_BinaryLogParser_LogFileDirectories {
	if in == nil {
		return nil
	}
	out := &pb.OracleSourceConfig_BinaryLogParser_LogFileDirectories{}
	out.OnlineLogDirectory = direct.ValueOf(in.OnlineLogDirectory)
	out.ArchivedLogDirectory = direct.ValueOf(in.ArchivedLogDirectory)
	return out
}
func OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess) *krmv1alpha1.OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess{}
	return out
}
func OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess) *pb.OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess {
	if in == nil {
		return nil
	}
	out := &pb.OracleSourceConfig_BinaryLogParser_OracleAsmLogFileAccess{}
	return out
}
func OracleSourceConfig_DropLargeObjects_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig_DropLargeObjects) *krmv1alpha1.OracleSourceConfig_DropLargeObjects {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleSourceConfig_DropLargeObjects{}
	return out
}
func OracleSourceConfig_DropLargeObjects_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleSourceConfig_DropLargeObjects) *pb.OracleSourceConfig_DropLargeObjects {
	if in == nil {
		return nil
	}
	out := &pb.OracleSourceConfig_DropLargeObjects{}
	return out
}
func OracleSourceConfig_LogMiner_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig_LogMiner) *krmv1alpha1.OracleSourceConfig_LogMiner {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleSourceConfig_LogMiner{}
	return out
}
func OracleSourceConfig_LogMiner_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleSourceConfig_LogMiner) *pb.OracleSourceConfig_LogMiner {
	if in == nil {
		return nil
	}
	out := &pb.OracleSourceConfig_LogMiner{}
	return out
}
func OracleSourceConfig_StreamLargeObjects_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig_StreamLargeObjects) *krmv1alpha1.OracleSourceConfig_StreamLargeObjects {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleSourceConfig_StreamLargeObjects{}
	return out
}
func OracleSourceConfig_StreamLargeObjects_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleSourceConfig_StreamLargeObjects) *pb.OracleSourceConfig_StreamLargeObjects {
	if in == nil {
		return nil
	}
	out := &pb.OracleSourceConfig_StreamLargeObjects{}
	return out
}
func OracleTable_FromProto(mapCtx *direct.MapContext, in *pb.OracleTable) *krmv1alpha1.OracleTable {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.OracleTable{}
	out.Table = direct.LazyPtr(in.GetTable())
	out.OracleColumns = direct.Slice_FromProto(mapCtx, in.OracleColumns, OracleColumn_FromProto)
	return out
}
func OracleTable_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.OracleTable) *pb.OracleTable {
	if in == nil {
		return nil
	}
	out := &pb.OracleTable{}
	out.Table = direct.ValueOf(in.Table)
	out.OracleColumns = direct.Slice_ToProto(mapCtx, in.OracleColumns, OracleColumn_ToProto)
	return out
}
func PostgreSQLColumn_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlColumn) *krmv1alpha1.PostgreSQLColumn {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PostgreSQLColumn{}
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
func PostgreSQLColumn_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PostgreSQLColumn) *pb.PostgresqlColumn {
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
func PostgreSQLSchema_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlSchema) *krmv1alpha1.PostgreSQLSchema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PostgreSQLSchema{}
	out.Schema = direct.LazyPtr(in.GetSchema())
	// MISSING: PostgresqlTables
	// (near miss): "PostgresqlTables" vs "PostgreSQLTables"
	return out
}
func PostgreSQLSchema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PostgreSQLSchema) *pb.PostgresqlSchema {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlSchema{}
	out.Schema = direct.ValueOf(in.Schema)
	// MISSING: PostgresqlTables
	// (near miss): "PostgresqlTables" vs "PostgreSQLTables"
	return out
}
func PostgreSQLTable_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlTable) *krmv1alpha1.PostgreSQLTable {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PostgreSQLTable{}
	out.Table = direct.LazyPtr(in.GetTable())
	// MISSING: PostgresqlColumns
	// (near miss): "PostgresqlColumns" vs "PostgreSQLColumns"
	return out
}
func PostgreSQLTable_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PostgreSQLTable) *pb.PostgresqlTable {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlTable{}
	out.Table = direct.ValueOf(in.Table)
	// MISSING: PostgresqlColumns
	// (near miss): "PostgresqlColumns" vs "PostgreSQLColumns"
	return out
}
func PostgresqlRdbms_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlRdbms) *krmv1alpha1.PostgresqlRdbms {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PostgresqlRdbms{}
	// MISSING: PostgresqlSchemas
	// (near miss): "PostgresqlSchemas" vs "PostgreSQLSchemas"
	return out
}
func PostgresqlRdbms_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PostgresqlRdbms) *pb.PostgresqlRdbms {
	if in == nil {
		return nil
	}
	out := &pb.PostgresqlRdbms{}
	// MISSING: PostgresqlSchemas
	// (near miss): "PostgresqlSchemas" vs "PostgreSQLSchemas"
	return out
}
func PostgresqlSourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.PostgresqlSourceConfig) *krmv1alpha1.PostgresqlSourceConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PostgresqlSourceConfig{}
	out.IncludeObjects = PostgresqlRdbms_FromProto(mapCtx, in.GetIncludeObjects())
	out.ExcludeObjects = PostgresqlRdbms_FromProto(mapCtx, in.GetExcludeObjects())
	out.ReplicationSlot = direct.LazyPtr(in.GetReplicationSlot())
	out.Publication = direct.LazyPtr(in.GetPublication())
	out.MaxConcurrentBackfillTasks = direct.LazyPtr(in.GetMaxConcurrentBackfillTasks())
	return out
}
func PostgresqlSourceConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PostgresqlSourceConfig) *pb.PostgresqlSourceConfig {
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
func PrivateConnectivity_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PrivateConnectivity) *pb.PrivateConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnectivity{}
	if in.PrivateConnectionRef != nil {
		out.PrivateConnection = in.PrivateConnectionRef.External
	}
	return out
}
func SQLServerChangeTables_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerChangeTables) *krmv1alpha1.SQLServerChangeTables {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SQLServerChangeTables{}
	return out
}
func SQLServerChangeTables_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SQLServerChangeTables) *pb.SqlServerChangeTables {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerChangeTables{}
	return out
}
func SQLServerColumn_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerColumn) *krmv1alpha1.SQLServerColumn {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SQLServerColumn{}
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
func SQLServerColumn_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SQLServerColumn) *pb.SqlServerColumn {
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
func SQLServerRdbms_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerRdbms) *krmv1alpha1.SQLServerRdbms {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SQLServerRdbms{}
	out.Schemas = direct.Slice_FromProto(mapCtx, in.Schemas, SQLServerSchema_FromProto)
	return out
}
func SQLServerRdbms_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SQLServerRdbms) *pb.SqlServerRdbms {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerRdbms{}
	out.Schemas = direct.Slice_ToProto(mapCtx, in.Schemas, SQLServerSchema_ToProto)
	return out
}
func SQLServerSchema_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerSchema) *krmv1alpha1.SQLServerSchema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SQLServerSchema{}
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.Tables = direct.Slice_FromProto(mapCtx, in.Tables, SQLServerTable_FromProto)
	return out
}
func SQLServerSchema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SQLServerSchema) *pb.SqlServerSchema {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerSchema{}
	out.Schema = direct.ValueOf(in.Schema)
	out.Tables = direct.Slice_ToProto(mapCtx, in.Tables, SQLServerTable_ToProto)
	return out
}
func SQLServerSourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerSourceConfig) *krmv1alpha1.SQLServerSourceConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SQLServerSourceConfig{}
	out.IncludeObjects = SQLServerRdbms_FromProto(mapCtx, in.GetIncludeObjects())
	out.ExcludeObjects = SQLServerRdbms_FromProto(mapCtx, in.GetExcludeObjects())
	out.MaxConcurrentCdcTasks = direct.LazyPtr(in.GetMaxConcurrentCdcTasks())
	out.MaxConcurrentBackfillTasks = direct.LazyPtr(in.GetMaxConcurrentBackfillTasks())
	out.TransactionLogs = SQLServerTransactionLogs_FromProto(mapCtx, in.GetTransactionLogs())
	out.ChangeTables = SQLServerChangeTables_FromProto(mapCtx, in.GetChangeTables())
	return out
}
func SQLServerSourceConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SQLServerSourceConfig) *pb.SqlServerSourceConfig {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerSourceConfig{}
	out.IncludeObjects = SQLServerRdbms_ToProto(mapCtx, in.IncludeObjects)
	out.ExcludeObjects = SQLServerRdbms_ToProto(mapCtx, in.ExcludeObjects)
	out.MaxConcurrentCdcTasks = direct.ValueOf(in.MaxConcurrentCdcTasks)
	out.MaxConcurrentBackfillTasks = direct.ValueOf(in.MaxConcurrentBackfillTasks)
	if oneof := SQLServerTransactionLogs_ToProto(mapCtx, in.TransactionLogs); oneof != nil {
		out.CdcMethod = &pb.SqlServerSourceConfig_TransactionLogs{TransactionLogs: oneof}
	}
	if oneof := SQLServerChangeTables_ToProto(mapCtx, in.ChangeTables); oneof != nil {
		out.CdcMethod = &pb.SqlServerSourceConfig_ChangeTables{ChangeTables: oneof}
	}
	return out
}
func SQLServerTable_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerTable) *krmv1alpha1.SQLServerTable {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SQLServerTable{}
	out.Table = direct.LazyPtr(in.GetTable())
	out.Columns = direct.Slice_FromProto(mapCtx, in.Columns, SQLServerColumn_FromProto)
	return out
}
func SQLServerTable_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SQLServerTable) *pb.SqlServerTable {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerTable{}
	out.Table = direct.ValueOf(in.Table)
	out.Columns = direct.Slice_ToProto(mapCtx, in.Columns, SQLServerColumn_ToProto)
	return out
}
func SQLServerTransactionLogs_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerTransactionLogs) *krmv1alpha1.SQLServerTransactionLogs {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SQLServerTransactionLogs{}
	return out
}
func SQLServerTransactionLogs_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SQLServerTransactionLogs) *pb.SqlServerTransactionLogs {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerTransactionLogs{}
	return out
}
func SourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.SourceConfig) *krmv1alpha1.SourceConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SourceConfig{}
	if in.GetSourceConnectionProfile() != "" {
		out.SourceConnectionProfileRef = &refs.*ConnectionProfileRef{External: in.GetSourceConnectionProfile()}
	}
	out.OracleSourceConfig = OracleSourceConfig_FromProto(mapCtx, in.GetOracleSourceConfig())
	// MISSING: MysqlSourceConfig
	// (near miss): "MysqlSourceConfig" vs "MySQLSourceConfig"
	// MISSING: PostgresqlSourceConfig
	// (near miss): "PostgresqlSourceConfig" vs "PostgreSQLSourceConfig"
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
	// MISSING: MysqlSourceConfig
	// (near miss): "MysqlSourceConfig" vs "MySQLSourceConfig"
	// MISSING: PostgresqlSourceConfig
	// (near miss): "PostgresqlSourceConfig" vs "PostgreSQLSourceConfig"
	if oneof := SQLServerSourceConfig_ToProto(mapCtx, in.SQLServerSourceConfig); oneof != nil {
		out.SourceStreamConfig = &pb.SourceConfig_SqlServerSourceConfig{SqlServerSourceConfig: oneof}
	}
	return out
}
func StaticServiceIPConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.StaticServiceIpConnectivity) *krmv1alpha1.StaticServiceIPConnectivity {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.StaticServiceIPConnectivity{}
	return out
}
func StaticServiceIPConnectivity_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.StaticServiceIPConnectivity) *pb.StaticServiceIpConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.StaticServiceIpConnectivity{}
	return out
}
func Stream_BackfillAllStrategy_FromProto(mapCtx *direct.MapContext, in *pb.Stream_BackfillAllStrategy) *krmv1alpha1.Stream_BackfillAllStrategy {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Stream_BackfillAllStrategy{}
	out.OracleExcludedObjects = OracleRdbms_FromProto(mapCtx, in.GetOracleExcludedObjects())
	// MISSING: MysqlExcludedObjects
	// (near miss): "MysqlExcludedObjects" vs "MySQLExcludedObjects"
	// MISSING: PostgresqlExcludedObjects
	// (near miss): "PostgresqlExcludedObjects" vs "PostgreSQLExcludedObjects"
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
	// MISSING: MysqlExcludedObjects
	// (near miss): "MysqlExcludedObjects" vs "MySQLExcludedObjects"
	// MISSING: PostgresqlExcludedObjects
	// (near miss): "PostgresqlExcludedObjects" vs "PostgreSQLExcludedObjects"
	if oneof := SQLServerRdbms_ToProto(mapCtx, in.SQLServerExcludedObjects); oneof != nil {
		out.ExcludedObjects = &pb.Stream_BackfillAllStrategy_SqlServerExcludedObjects{SqlServerExcludedObjects: oneof}
	}
	return out
}
func Stream_BackfillNoneStrategy_FromProto(mapCtx *direct.MapContext, in *pb.Stream_BackfillNoneStrategy) *krmv1alpha1.Stream_BackfillNoneStrategy {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Stream_BackfillNoneStrategy{}
	return out
}
func Stream_BackfillNoneStrategy_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Stream_BackfillNoneStrategy) *pb.Stream_BackfillNoneStrategy {
	if in == nil {
		return nil
	}
	out := &pb.Stream_BackfillNoneStrategy{}
	return out
}
