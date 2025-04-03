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
// krm.group: datacatalog.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.datacatalog.v1

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigQueryConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryConnectionSpec) *krmv1alpha1.BigQueryConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryConnectionSpec{}
	out.ConnectionType = direct.Enum_FromProto(mapCtx, in.GetConnectionType())
	out.CloudSQL = CloudSQLBigQueryConnectionSpec_FromProto(mapCtx, in.GetCloudSql())
	out.HasCredential = direct.LazyPtr(in.GetHasCredential())
	return out
}
func BigQueryConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryConnectionSpec) *pb.BigQueryConnectionSpec {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryConnectionSpec{}
	out.ConnectionType = direct.Enum_ToProto[pb.BigQueryConnectionSpec_ConnectionType](mapCtx, in.ConnectionType)
	if oneof := CloudSQLBigQueryConnectionSpec_ToProto(mapCtx, in.CloudSQL); oneof != nil {
		out.ConnectionSpec = &pb.BigQueryConnectionSpec_CloudSql{CloudSql: oneof}
	}
	out.HasCredential = direct.ValueOf(in.HasCredential)
	return out
}
func BigQueryDateShardedSpec_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDateShardedSpec) *krmv1alpha1.BigQueryDateShardedSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryDateShardedSpec{}
	// MISSING: Dataset
	// MISSING: TablePrefix
	// MISSING: ShardCount
	// MISSING: LatestShardResource
	return out
}
func BigQueryDateShardedSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryDateShardedSpec) *pb.BigQueryDateShardedSpec {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDateShardedSpec{}
	// MISSING: Dataset
	// MISSING: TablePrefix
	// MISSING: ShardCount
	// MISSING: LatestShardResource
	return out
}
func BigQueryDateShardedSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDateShardedSpec) *krmv1alpha1.BigQueryDateShardedSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryDateShardedSpecObservedState{}
	out.Dataset = direct.LazyPtr(in.GetDataset())
	out.TablePrefix = direct.LazyPtr(in.GetTablePrefix())
	out.ShardCount = direct.LazyPtr(in.GetShardCount())
	out.LatestShardResource = direct.LazyPtr(in.GetLatestShardResource())
	return out
}
func BigQueryDateShardedSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryDateShardedSpecObservedState) *pb.BigQueryDateShardedSpec {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDateShardedSpec{}
	out.Dataset = direct.ValueOf(in.Dataset)
	out.TablePrefix = direct.ValueOf(in.TablePrefix)
	out.ShardCount = direct.ValueOf(in.ShardCount)
	out.LatestShardResource = direct.ValueOf(in.LatestShardResource)
	return out
}
func BigQueryRoutineSpec_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryRoutineSpec) *krmv1alpha1.BigQueryRoutineSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryRoutineSpec{}
	out.ImportedLibraries = in.ImportedLibraries
	return out
}
func BigQueryRoutineSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryRoutineSpec) *pb.BigQueryRoutineSpec {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryRoutineSpec{}
	out.ImportedLibraries = in.ImportedLibraries
	return out
}
func BigQueryTableSpec_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryTableSpec) *krmv1alpha1.BigQueryTableSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryTableSpec{}
	// MISSING: TableSourceType
	out.ViewSpec = ViewSpec_FromProto(mapCtx, in.GetViewSpec())
	out.TableSpec = TableSpec_FromProto(mapCtx, in.GetTableSpec())
	return out
}
func BigQueryTableSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryTableSpec) *pb.BigQueryTableSpec {
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
func BigQueryTableSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryTableSpec) *krmv1alpha1.BigQueryTableSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryTableSpecObservedState{}
	out.TableSourceType = direct.Enum_FromProto(mapCtx, in.GetTableSourceType())
	out.ViewSpec = ViewSpecObservedState_FromProto(mapCtx, in.GetViewSpec())
	out.TableSpec = TableSpecObservedState_FromProto(mapCtx, in.GetTableSpec())
	return out
}
func BigQueryTableSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryTableSpecObservedState) *pb.BigQueryTableSpec {
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
func BusinessContext_FromProto(mapCtx *direct.MapContext, in *pb.BusinessContext) *krmv1alpha1.BusinessContext {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BusinessContext{}
	out.EntryOverview = EntryOverview_FromProto(mapCtx, in.GetEntryOverview())
	out.Contacts = Contacts_FromProto(mapCtx, in.GetContacts())
	return out
}
func BusinessContext_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BusinessContext) *pb.BusinessContext {
	if in == nil {
		return nil
	}
	out := &pb.BusinessContext{}
	out.EntryOverview = EntryOverview_ToProto(mapCtx, in.EntryOverview)
	out.Contacts = Contacts_ToProto(mapCtx, in.Contacts)
	return out
}
func CloudBigtableInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudBigtableInstanceSpec) *krmv1alpha1.CloudBigtableInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.CloudBigtableInstanceSpec{}
	out.CloudBigtableClusterSpecs = direct.Slice_FromProto(mapCtx, in.CloudBigtableClusterSpecs, CloudBigtableInstanceSpec_CloudBigtableClusterSpec_FromProto)
	return out
}
func CloudBigtableInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.CloudBigtableInstanceSpec) *pb.CloudBigtableInstanceSpec {
	if in == nil {
		return nil
	}
	out := &pb.CloudBigtableInstanceSpec{}
	out.CloudBigtableClusterSpecs = direct.Slice_ToProto(mapCtx, in.CloudBigtableClusterSpecs, CloudBigtableInstanceSpec_CloudBigtableClusterSpec_ToProto)
	return out
}
func CloudBigtableInstanceSpec_CloudBigtableClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudBigtableInstanceSpec_CloudBigtableClusterSpec) *krmv1alpha1.CloudBigtableInstanceSpec_CloudBigtableClusterSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.CloudBigtableInstanceSpec_CloudBigtableClusterSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Type = direct.LazyPtr(in.GetType())
	out.LinkedResource = direct.LazyPtr(in.GetLinkedResource())
	return out
}
func CloudBigtableInstanceSpec_CloudBigtableClusterSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.CloudBigtableInstanceSpec_CloudBigtableClusterSpec) *pb.CloudBigtableInstanceSpec_CloudBigtableClusterSpec {
	if in == nil {
		return nil
	}
	out := &pb.CloudBigtableInstanceSpec_CloudBigtableClusterSpec{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Location = direct.ValueOf(in.Location)
	out.Type = direct.ValueOf(in.Type)
	out.LinkedResource = direct.ValueOf(in.LinkedResource)
	return out
}
func CloudBigtableSystemSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudBigtableSystemSpec) *krmv1alpha1.CloudBigtableSystemSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.CloudBigtableSystemSpec{}
	out.InstanceDisplayName = direct.LazyPtr(in.GetInstanceDisplayName())
	return out
}
func CloudBigtableSystemSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.CloudBigtableSystemSpec) *pb.CloudBigtableSystemSpec {
	if in == nil {
		return nil
	}
	out := &pb.CloudBigtableSystemSpec{}
	out.InstanceDisplayName = direct.ValueOf(in.InstanceDisplayName)
	return out
}
func CloudSQLBigQueryConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlBigQueryConnectionSpec) *krmv1alpha1.CloudSQLBigQueryConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.CloudSQLBigQueryConnectionSpec{}
	out.InstanceID = direct.LazyPtr(in.GetInstanceId())
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func CloudSQLBigQueryConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.CloudSQLBigQueryConnectionSpec) *pb.CloudSqlBigQueryConnectionSpec {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlBigQueryConnectionSpec{}
	out.InstanceId = direct.ValueOf(in.InstanceID)
	out.Database = direct.ValueOf(in.Database)
	out.Type = direct.Enum_ToProto[pb.CloudSqlBigQueryConnectionSpec_DatabaseType](mapCtx, in.Type)
	return out
}
func ColumnSchema_FromProto(mapCtx *direct.MapContext, in *pb.ColumnSchema) *krmv1alpha1.ColumnSchema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ColumnSchema{}
	out.Column = direct.LazyPtr(in.GetColumn())
	out.Type = direct.LazyPtr(in.GetType())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Mode = direct.LazyPtr(in.GetMode())
	out.DefaultValue = direct.LazyPtr(in.GetDefaultValue())
	out.OrdinalPosition = direct.LazyPtr(in.GetOrdinalPosition())
	out.HighestIndexingType = direct.Enum_FromProto(mapCtx, in.GetHighestIndexingType())
	out.Subcolumns = direct.Slice_FromProto(mapCtx, in.Subcolumns, ColumnSchema_FromProto)
	out.LookerColumnSpec = ColumnSchema_LookerColumnSpec_FromProto(mapCtx, in.GetLookerColumnSpec())
	out.RangeElementType = ColumnSchema_FieldElementType_FromProto(mapCtx, in.GetRangeElementType())
	out.GcRule = direct.LazyPtr(in.GetGcRule())
	return out
}
func ColumnSchema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ColumnSchema) *pb.ColumnSchema {
	if in == nil {
		return nil
	}
	out := &pb.ColumnSchema{}
	out.Column = direct.ValueOf(in.Column)
	out.Type = direct.ValueOf(in.Type)
	out.Description = direct.ValueOf(in.Description)
	out.Mode = direct.ValueOf(in.Mode)
	out.DefaultValue = direct.ValueOf(in.DefaultValue)
	out.OrdinalPosition = direct.ValueOf(in.OrdinalPosition)
	out.HighestIndexingType = direct.Enum_ToProto[pb.ColumnSchema_IndexingType](mapCtx, in.HighestIndexingType)
	out.Subcolumns = direct.Slice_ToProto(mapCtx, in.Subcolumns, ColumnSchema_ToProto)
	if oneof := ColumnSchema_LookerColumnSpec_ToProto(mapCtx, in.LookerColumnSpec); oneof != nil {
		out.SystemSpec = &pb.ColumnSchema_LookerColumnSpec_{LookerColumnSpec: oneof}
	}
	out.RangeElementType = ColumnSchema_FieldElementType_ToProto(mapCtx, in.RangeElementType)
	out.GcRule = direct.ValueOf(in.GcRule)
	return out
}
func ColumnSchema_FieldElementType_FromProto(mapCtx *direct.MapContext, in *pb.ColumnSchema_FieldElementType) *krmv1alpha1.ColumnSchema_FieldElementType {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ColumnSchema_FieldElementType{}
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func ColumnSchema_FieldElementType_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ColumnSchema_FieldElementType) *pb.ColumnSchema_FieldElementType {
	if in == nil {
		return nil
	}
	out := &pb.ColumnSchema_FieldElementType{}
	out.Type = direct.ValueOf(in.Type)
	return out
}
func ColumnSchema_LookerColumnSpec_FromProto(mapCtx *direct.MapContext, in *pb.ColumnSchema_LookerColumnSpec) *krmv1alpha1.ColumnSchema_LookerColumnSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ColumnSchema_LookerColumnSpec{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func ColumnSchema_LookerColumnSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ColumnSchema_LookerColumnSpec) *pb.ColumnSchema_LookerColumnSpec {
	if in == nil {
		return nil
	}
	out := &pb.ColumnSchema_LookerColumnSpec{}
	out.Type = direct.Enum_ToProto[pb.ColumnSchema_LookerColumnSpec_LookerColumnType](mapCtx, in.Type)
	return out
}
func CommonUsageStats_FromProto(mapCtx *direct.MapContext, in *pb.CommonUsageStats) *krmv1alpha1.CommonUsageStats {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.CommonUsageStats{}
	out.ViewCount = in.ViewCount
	return out
}
func CommonUsageStats_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.CommonUsageStats) *pb.CommonUsageStats {
	if in == nil {
		return nil
	}
	out := &pb.CommonUsageStats{}
	out.ViewCount = in.ViewCount
	return out
}
func Contacts_FromProto(mapCtx *direct.MapContext, in *pb.Contacts) *krmv1alpha1.Contacts {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Contacts{}
	out.People = direct.Slice_FromProto(mapCtx, in.People, Contacts_Person_FromProto)
	return out
}
func Contacts_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Contacts) *pb.Contacts {
	if in == nil {
		return nil
	}
	out := &pb.Contacts{}
	out.People = direct.Slice_ToProto(mapCtx, in.People, Contacts_Person_ToProto)
	return out
}
func Contacts_Person_FromProto(mapCtx *direct.MapContext, in *pb.Contacts_Person) *krmv1alpha1.Contacts_Person {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Contacts_Person{}
	out.Designation = direct.LazyPtr(in.GetDesignation())
	out.Email = direct.LazyPtr(in.GetEmail())
	return out
}
func Contacts_Person_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Contacts_Person) *pb.Contacts_Person {
	if in == nil {
		return nil
	}
	out := &pb.Contacts_Person{}
	out.Designation = direct.ValueOf(in.Designation)
	out.Email = direct.ValueOf(in.Email)
	return out
}
func DataCatalogEntryGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krmv1alpha1.DataCatalogEntryGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataCatalogEntryGroupObservedState{}
	// MISSING: Name
	out.DataCatalogTimestamps = SystemTimestamps_FromProto(mapCtx, in.GetDataCatalogTimestamps())
	return out
}
func DataCatalogEntryGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataCatalogEntryGroupObservedState) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	// MISSING: Name
	out.DataCatalogTimestamps = SystemTimestamps_ToProto(mapCtx, in.DataCatalogTimestamps)
	return out
}
func DataCatalogEntryGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krmv1alpha1.DataCatalogEntryGroupSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataCatalogEntryGroupSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.TransferredToDataplex = direct.LazyPtr(in.GetTransferredToDataplex())
	return out
}
func DataCatalogEntryGroupSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataCatalogEntryGroupSpec) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.TransferredToDataplex = direct.ValueOf(in.TransferredToDataplex)
	return out
}
func DataCatalogTagTemplateSpec_FromProto(mapCtx *direct.MapContext, in *pb.TagTemplate) *krmv1alpha1.DataCatalogTagTemplateSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataCatalogTagTemplateSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.IsPubliclyReadable = direct.LazyPtr(in.GetIsPubliclyReadable())
	// TODO: map type string message for field Fields
	out.DataplexTransferStatus = direct.Enum_FromProto(mapCtx, in.GetDataplexTransferStatus())
	return out
}
func DataCatalogTagTemplateSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataCatalogTagTemplateSpec) *pb.TagTemplate {
	if in == nil {
		return nil
	}
	out := &pb.TagTemplate{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.IsPubliclyReadable = direct.ValueOf(in.IsPubliclyReadable)
	// TODO: map type string message for field Fields
	out.DataplexTransferStatus = direct.Enum_ToProto[pb.TagTemplate_DataplexTransferStatus](mapCtx, in.DataplexTransferStatus)
	return out
}
func DataSource_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krmv1alpha1.DataSource {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataSource{}
	// MISSING: Service
	// MISSING: Resource
	// MISSING: SourceEntry
	// MISSING: StorageProperties
	return out
}
func DataSource_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataSource) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	// MISSING: Service
	// MISSING: Resource
	// MISSING: SourceEntry
	// MISSING: StorageProperties
	return out
}
func DataSourceConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataSourceConnectionSpec) *krmv1alpha1.DataSourceConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataSourceConnectionSpec{}
	out.BigqueryConnectionSpec = BigQueryConnectionSpec_FromProto(mapCtx, in.GetBigqueryConnectionSpec())
	return out
}
func DataSourceConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataSourceConnectionSpec) *pb.DataSourceConnectionSpec {
	if in == nil {
		return nil
	}
	out := &pb.DataSourceConnectionSpec{}
	out.BigqueryConnectionSpec = BigQueryConnectionSpec_ToProto(mapCtx, in.BigqueryConnectionSpec)
	return out
}
func DataSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krmv1alpha1.DataSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataSourceObservedState{}
	out.Service = direct.Enum_FromProto(mapCtx, in.GetService())
	out.Resource = direct.LazyPtr(in.GetResource())
	out.SourceEntry = direct.LazyPtr(in.GetSourceEntry())
	out.StorageProperties = StorageProperties_FromProto(mapCtx, in.GetStorageProperties())
	return out
}
func DataSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataSourceObservedState) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	out.Service = direct.Enum_ToProto[pb.DataSource_Service](mapCtx, in.Service)
	out.Resource = direct.ValueOf(in.Resource)
	out.SourceEntry = direct.ValueOf(in.SourceEntry)
	if oneof := StorageProperties_ToProto(mapCtx, in.StorageProperties); oneof != nil {
		out.Properties = &pb.DataSource_StorageProperties{StorageProperties: oneof}
	}
	return out
}
func DatabaseTableSpec_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseTableSpec) *krmv1alpha1.DatabaseTableSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DatabaseTableSpec{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: DataplexTable
	out.DatabaseViewSpec = DatabaseTableSpec_DatabaseViewSpec_FromProto(mapCtx, in.GetDatabaseViewSpec())
	return out
}
func DatabaseTableSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DatabaseTableSpec) *pb.DatabaseTableSpec {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseTableSpec{}
	out.Type = direct.Enum_ToProto[pb.DatabaseTableSpec_TableType](mapCtx, in.Type)
	// MISSING: DataplexTable
	out.DatabaseViewSpec = DatabaseTableSpec_DatabaseViewSpec_ToProto(mapCtx, in.DatabaseViewSpec)
	return out
}
func DatabaseTableSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseTableSpec) *krmv1alpha1.DatabaseTableSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DatabaseTableSpecObservedState{}
	// MISSING: Type
	out.DataplexTable = DataplexTableSpec_FromProto(mapCtx, in.GetDataplexTable())
	// MISSING: DatabaseViewSpec
	return out
}
func DatabaseTableSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DatabaseTableSpecObservedState) *pb.DatabaseTableSpec {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseTableSpec{}
	// MISSING: Type
	out.DataplexTable = DataplexTableSpec_ToProto(mapCtx, in.DataplexTable)
	// MISSING: DatabaseViewSpec
	return out
}
func DatabaseTableSpec_DatabaseViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseTableSpec_DatabaseViewSpec) *krmv1alpha1.DatabaseTableSpec_DatabaseViewSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DatabaseTableSpec_DatabaseViewSpec{}
	out.ViewType = direct.Enum_FromProto(mapCtx, in.GetViewType())
	out.BaseTable = direct.LazyPtr(in.GetBaseTable())
	out.SQLQuery = direct.LazyPtr(in.GetSqlQuery())
	return out
}
func DataplexExternalTable_FromProto(mapCtx *direct.MapContext, in *pb.DataplexExternalTable) *krmv1alpha1.DataplexExternalTable {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataplexExternalTable{}
	out.System = direct.Enum_FromProto(mapCtx, in.GetSystem())
	out.FullyQualifiedName = direct.LazyPtr(in.GetFullyQualifiedName())
	out.GoogleCloudResource = direct.LazyPtr(in.GetGoogleCloudResource())
	out.DataCatalogEntry = direct.LazyPtr(in.GetDataCatalogEntry())
	return out
}
func DataplexExternalTable_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataplexExternalTable) *pb.DataplexExternalTable {
	if in == nil {
		return nil
	}
	out := &pb.DataplexExternalTable{}
	out.System = direct.Enum_ToProto[pb.IntegratedSystem](mapCtx, in.System)
	out.FullyQualifiedName = direct.ValueOf(in.FullyQualifiedName)
	out.GoogleCloudResource = direct.ValueOf(in.GoogleCloudResource)
	out.DataCatalogEntry = direct.ValueOf(in.DataCatalogEntry)
	return out
}
func DataplexFilesetSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataplexFilesetSpec) *krmv1alpha1.DataplexFilesetSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataplexFilesetSpec{}
	out.DataplexSpec = DataplexSpec_FromProto(mapCtx, in.GetDataplexSpec())
	return out
}
func DataplexFilesetSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataplexFilesetSpec) *pb.DataplexFilesetSpec {
	if in == nil {
		return nil
	}
	out := &pb.DataplexFilesetSpec{}
	out.DataplexSpec = DataplexSpec_ToProto(mapCtx, in.DataplexSpec)
	return out
}
func DataplexSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataplexSpec) *krmv1alpha1.DataplexSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataplexSpec{}
	out.Asset = direct.LazyPtr(in.GetAsset())
	out.DataFormat = PhysicalSchema_FromProto(mapCtx, in.GetDataFormat())
	out.CompressionFormat = direct.LazyPtr(in.GetCompressionFormat())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	return out
}
func DataplexSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataplexSpec) *pb.DataplexSpec {
	if in == nil {
		return nil
	}
	out := &pb.DataplexSpec{}
	out.Asset = direct.ValueOf(in.Asset)
	out.DataFormat = PhysicalSchema_ToProto(mapCtx, in.DataFormat)
	out.CompressionFormat = direct.ValueOf(in.CompressionFormat)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	return out
}
func DataplexTableSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataplexTableSpec) *krmv1alpha1.DataplexTableSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataplexTableSpec{}
	out.ExternalTables = direct.Slice_FromProto(mapCtx, in.ExternalTables, DataplexExternalTable_FromProto)
	out.DataplexSpec = DataplexSpec_FromProto(mapCtx, in.GetDataplexSpec())
	out.UserManaged = direct.LazyPtr(in.GetUserManaged())
	return out
}
func DataplexTableSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataplexTableSpec) *pb.DataplexTableSpec {
	if in == nil {
		return nil
	}
	out := &pb.DataplexTableSpec{}
	out.ExternalTables = direct.Slice_ToProto(mapCtx, in.ExternalTables, DataplexExternalTable_ToProto)
	out.DataplexSpec = DataplexSpec_ToProto(mapCtx, in.DataplexSpec)
	out.UserManaged = direct.ValueOf(in.UserManaged)
	return out
}
func DatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.DatasetSpec) *krmv1alpha1.DatasetSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DatasetSpec{}
	out.VertexDatasetSpec = VertexDatasetSpec_FromProto(mapCtx, in.GetVertexDatasetSpec())
	return out
}
func DatasetSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DatasetSpec) *pb.DatasetSpec {
	if in == nil {
		return nil
	}
	out := &pb.DatasetSpec{}
	if oneof := VertexDatasetSpec_ToProto(mapCtx, in.VertexDatasetSpec); oneof != nil {
		out.SystemSpec = &pb.DatasetSpec_VertexDatasetSpec{VertexDatasetSpec: oneof}
	}
	return out
}
func EntryOverview_FromProto(mapCtx *direct.MapContext, in *pb.EntryOverview) *krmv1alpha1.EntryOverview {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.EntryOverview{}
	out.Overview = direct.LazyPtr(in.GetOverview())
	return out
}
func EntryOverview_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.EntryOverview) *pb.EntryOverview {
	if in == nil {
		return nil
	}
	out := &pb.EntryOverview{}
	out.Overview = direct.ValueOf(in.Overview)
	return out
}
func FeatureOnlineStoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.FeatureOnlineStoreSpec) *krmv1alpha1.FeatureOnlineStoreSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.FeatureOnlineStoreSpec{}
	// MISSING: StorageType
	return out
}
func FeatureOnlineStoreSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.FeatureOnlineStoreSpec) *pb.FeatureOnlineStoreSpec {
	if in == nil {
		return nil
	}
	out := &pb.FeatureOnlineStoreSpec{}
	// MISSING: StorageType
	return out
}
func FeatureOnlineStoreSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FeatureOnlineStoreSpec) *krmv1alpha1.FeatureOnlineStoreSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.FeatureOnlineStoreSpecObservedState{}
	out.StorageType = direct.Enum_FromProto(mapCtx, in.GetStorageType())
	return out
}
func FeatureOnlineStoreSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.FeatureOnlineStoreSpecObservedState) *pb.FeatureOnlineStoreSpec {
	if in == nil {
		return nil
	}
	out := &pb.FeatureOnlineStoreSpec{}
	out.StorageType = direct.Enum_ToProto[pb.FeatureOnlineStoreSpec_StorageType](mapCtx, in.StorageType)
	return out
}
func FieldType_FromProto(mapCtx *direct.MapContext, in *pb.FieldType) *krmv1alpha1.FieldType {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.FieldType{}
	out.PrimitiveType = direct.Enum_FromProto(mapCtx, in.GetPrimitiveType())
	out.EnumType = FieldType_EnumType_FromProto(mapCtx, in.GetEnumType())
	return out
}

func FieldType_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.FieldType) *pb.FieldType {
	if in == nil {
		return nil
	}
	out := &pb.FieldType{}
	if oneof := FieldType_PrimitiveType_ToProto(mapCtx, in.PrimitiveType); oneof != nil {
		out.TypeDecl = oneof
	}
	if oneof := FieldType_EnumType_ToProto(mapCtx, in.EnumType); oneof != nil {
		out.TypeDecl = &pb.FieldType_EnumType_{EnumType: oneof}
	}
	return out
}
func FieldType_EnumType_FromProto(mapCtx *direct.MapContext, in *pb.FieldType_EnumType) *krmv1alpha1.FieldType_EnumType {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.FieldType_EnumType{}
	out.AllowedValues = direct.Slice_FromProto(mapCtx, in.AllowedValues, FieldType_EnumType_EnumValue_FromProto)
	return out
}
func FieldType_EnumType_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.FieldType_EnumType) *pb.FieldType_EnumType {
	if in == nil {
		return nil
	}
	out := &pb.FieldType_EnumType{}
	out.AllowedValues = direct.Slice_ToProto(mapCtx, in.AllowedValues, FieldType_EnumType_EnumValue_ToProto)
	return out
}
func FieldType_EnumType_EnumValue_FromProto(mapCtx *direct.MapContext, in *pb.FieldType_EnumType_EnumValue) *krmv1alpha1.FieldType_EnumType_EnumValue {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.FieldType_EnumType_EnumValue{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func FieldType_EnumType_EnumValue_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.FieldType_EnumType_EnumValue) *pb.FieldType_EnumType_EnumValue {
	if in == nil {
		return nil
	}
	out := &pb.FieldType_EnumType_EnumValue{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func FilesetSpec_FromProto(mapCtx *direct.MapContext, in *pb.FilesetSpec) *krmv1alpha1.FilesetSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.FilesetSpec{}
	out.DataplexFileset = DataplexFilesetSpec_FromProto(mapCtx, in.GetDataplexFileset())
	return out
}
func FilesetSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.FilesetSpec) *pb.FilesetSpec {
	if in == nil {
		return nil
	}
	out := &pb.FilesetSpec{}
	out.DataplexFileset = DataplexFilesetSpec_ToProto(mapCtx, in.DataplexFileset)
	return out
}
func GCSFileSpec_FromProto(mapCtx *direct.MapContext, in *pb.GcsFileSpec) *krmv1alpha1.GCSFileSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.GCSFileSpec{}
	out.FilePath = direct.LazyPtr(in.GetFilePath())
	// MISSING: GCSTimestamps
	// MISSING: SizeBytes
	return out
}
func GCSFileSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.GCSFileSpec) *pb.GcsFileSpec {
	if in == nil {
		return nil
	}
	out := &pb.GcsFileSpec{}
	out.FilePath = direct.ValueOf(in.FilePath)
	// MISSING: GCSTimestamps
	// MISSING: SizeBytes
	return out
}
func GCSFileSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GcsFileSpec) *krmv1alpha1.GCSFileSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.GCSFileSpecObservedState{}
	// MISSING: FilePath
	out.GCSTimestamps = SystemTimestamps_FromProto(mapCtx, in.GetGcsTimestamps())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	return out
}
func GCSFileSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.GCSFileSpecObservedState) *pb.GcsFileSpec {
	if in == nil {
		return nil
	}
	out := &pb.GcsFileSpec{}
	// MISSING: FilePath
	out.GcsTimestamps = SystemTimestamps_ToProto(mapCtx, in.GCSTimestamps)
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	return out
}
func GCSFilesetSpec_FromProto(mapCtx *direct.MapContext, in *pb.GcsFilesetSpec) *krmv1alpha1.GCSFilesetSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.GCSFilesetSpec{}
	out.FilePatterns = in.FilePatterns
	// MISSING: SampleGCSFileSpecs
	return out
}
func GCSFilesetSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.GCSFilesetSpec) *pb.GcsFilesetSpec {
	if in == nil {
		return nil
	}
	out := &pb.GcsFilesetSpec{}
	out.FilePatterns = in.FilePatterns
	// MISSING: SampleGCSFileSpecs
	return out
}
func GCSFilesetSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.GCSFilesetSpecObservedState) *pb.GcsFilesetSpec {
	if in == nil {
		return nil
	}
	out := &pb.GcsFilesetSpec{}
	// MISSING: FilePatterns
	out.SampleGcsFileSpecs = direct.Slice_ToProto(mapCtx, in.SampleGCSFileSpecs, GCSFileSpec_ToProto)
	return out
}
func LookerSystemSpec_FromProto(mapCtx *direct.MapContext, in *pb.LookerSystemSpec) *krmv1alpha1.LookerSystemSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.LookerSystemSpec{}
	out.ParentInstanceID = direct.LazyPtr(in.GetParentInstanceId())
	out.ParentInstanceDisplayName = direct.LazyPtr(in.GetParentInstanceDisplayName())
	out.ParentModelID = direct.LazyPtr(in.GetParentModelId())
	out.ParentModelDisplayName = direct.LazyPtr(in.GetParentModelDisplayName())
	out.ParentViewID = direct.LazyPtr(in.GetParentViewId())
	out.ParentViewDisplayName = direct.LazyPtr(in.GetParentViewDisplayName())
	return out
}
func LookerSystemSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.LookerSystemSpec) *pb.LookerSystemSpec {
	if in == nil {
		return nil
	}
	out := &pb.LookerSystemSpec{}
	out.ParentInstanceId = direct.ValueOf(in.ParentInstanceID)
	out.ParentInstanceDisplayName = direct.ValueOf(in.ParentInstanceDisplayName)
	out.ParentModelId = direct.ValueOf(in.ParentModelID)
	out.ParentModelDisplayName = direct.ValueOf(in.ParentModelDisplayName)
	out.ParentViewId = direct.ValueOf(in.ParentViewID)
	out.ParentViewDisplayName = direct.ValueOf(in.ParentViewDisplayName)
	return out
}
func ModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.ModelSpec) *krmv1alpha1.ModelSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ModelSpec{}
	out.VertexModelSpec = VertexModelSpec_FromProto(mapCtx, in.GetVertexModelSpec())
	return out
}
func ModelSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ModelSpec) *pb.ModelSpec {
	if in == nil {
		return nil
	}
	out := &pb.ModelSpec{}
	if oneof := VertexModelSpec_ToProto(mapCtx, in.VertexModelSpec); oneof != nil {
		out.SystemSpec = &pb.ModelSpec_VertexModelSpec{VertexModelSpec: oneof}
	}
	return out
}
func PersonalDetails_FromProto(mapCtx *direct.MapContext, in *pb.PersonalDetails) *krmv1alpha1.PersonalDetails {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PersonalDetails{}
	out.Starred = direct.LazyPtr(in.GetStarred())
	out.StarTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStarTime())
	return out
}
func PersonalDetails_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PersonalDetails) *pb.PersonalDetails {
	if in == nil {
		return nil
	}
	out := &pb.PersonalDetails{}
	out.Starred = direct.ValueOf(in.Starred)
	out.StarTime = direct.StringTimestamp_ToProto(mapCtx, in.StarTime)
	return out
}
func PhysicalSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema) *krmv1alpha1.PhysicalSchema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PhysicalSchema{}
	out.Avro = PhysicalSchema_AvroSchema_FromProto(mapCtx, in.GetAvro())
	out.Thrift = PhysicalSchema_ThriftSchema_FromProto(mapCtx, in.GetThrift())
	out.Protobuf = PhysicalSchema_ProtobufSchema_FromProto(mapCtx, in.GetProtobuf())
	out.Parquet = PhysicalSchema_ParquetSchema_FromProto(mapCtx, in.GetParquet())
	out.Orc = PhysicalSchema_OrcSchema_FromProto(mapCtx, in.GetOrc())
	out.Csv = PhysicalSchema_CsvSchema_FromProto(mapCtx, in.GetCsv())
	return out
}
func PhysicalSchema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PhysicalSchema) *pb.PhysicalSchema {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalSchema{}
	if oneof := PhysicalSchema_AvroSchema_ToProto(mapCtx, in.Avro); oneof != nil {
		out.Schema = &pb.PhysicalSchema_Avro{Avro: oneof}
	}
	if oneof := PhysicalSchema_ThriftSchema_ToProto(mapCtx, in.Thrift); oneof != nil {
		out.Schema = &pb.PhysicalSchema_Thrift{Thrift: oneof}
	}
	if oneof := PhysicalSchema_ProtobufSchema_ToProto(mapCtx, in.Protobuf); oneof != nil {
		out.Schema = &pb.PhysicalSchema_Protobuf{Protobuf: oneof}
	}
	if oneof := PhysicalSchema_ParquetSchema_ToProto(mapCtx, in.Parquet); oneof != nil {
		out.Schema = &pb.PhysicalSchema_Parquet{Parquet: oneof}
	}
	if oneof := PhysicalSchema_OrcSchema_ToProto(mapCtx, in.Orc); oneof != nil {
		out.Schema = &pb.PhysicalSchema_Orc{Orc: oneof}
	}
	if oneof := PhysicalSchema_CsvSchema_ToProto(mapCtx, in.Csv); oneof != nil {
		out.Schema = &pb.PhysicalSchema_Csv{Csv: oneof}
	}
	return out
}
func PhysicalSchema_AvroSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema_AvroSchema) *krmv1alpha1.PhysicalSchema_AvroSchema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PhysicalSchema_AvroSchema{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func PhysicalSchema_AvroSchema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PhysicalSchema_AvroSchema) *pb.PhysicalSchema_AvroSchema {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalSchema_AvroSchema{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
func PhysicalSchema_CsvSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema_CsvSchema) *krmv1alpha1.PhysicalSchema_CsvSchema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PhysicalSchema_CsvSchema{}
	return out
}
func PhysicalSchema_CsvSchema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PhysicalSchema_CsvSchema) *pb.PhysicalSchema_CsvSchema {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalSchema_CsvSchema{}
	return out
}
func PhysicalSchema_OrcSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema_OrcSchema) *krmv1alpha1.PhysicalSchema_OrcSchema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PhysicalSchema_OrcSchema{}
	return out
}
func PhysicalSchema_OrcSchema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PhysicalSchema_OrcSchema) *pb.PhysicalSchema_OrcSchema {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalSchema_OrcSchema{}
	return out
}
func PhysicalSchema_ParquetSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema_ParquetSchema) *krmv1alpha1.PhysicalSchema_ParquetSchema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PhysicalSchema_ParquetSchema{}
	return out
}
func PhysicalSchema_ParquetSchema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PhysicalSchema_ParquetSchema) *pb.PhysicalSchema_ParquetSchema {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalSchema_ParquetSchema{}
	return out
}
func PhysicalSchema_ProtobufSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema_ProtobufSchema) *krmv1alpha1.PhysicalSchema_ProtobufSchema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PhysicalSchema_ProtobufSchema{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func PhysicalSchema_ProtobufSchema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PhysicalSchema_ProtobufSchema) *pb.PhysicalSchema_ProtobufSchema {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalSchema_ProtobufSchema{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
func PhysicalSchema_ThriftSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema_ThriftSchema) *krmv1alpha1.PhysicalSchema_ThriftSchema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.PhysicalSchema_ThriftSchema{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func PhysicalSchema_ThriftSchema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.PhysicalSchema_ThriftSchema) *pb.PhysicalSchema_ThriftSchema {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalSchema_ThriftSchema{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
func RoutineSpec_FromProto(mapCtx *direct.MapContext, in *pb.RoutineSpec) *krmv1alpha1.RoutineSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.RoutineSpec{}
	out.RoutineType = direct.Enum_FromProto(mapCtx, in.GetRoutineType())
	out.Language = direct.LazyPtr(in.GetLanguage())
	out.RoutineArguments = direct.Slice_FromProto(mapCtx, in.RoutineArguments, RoutineSpec_Argument_FromProto)
	out.ReturnType = direct.LazyPtr(in.GetReturnType())
	out.DefinitionBody = direct.LazyPtr(in.GetDefinitionBody())
	out.BigqueryRoutineSpec = BigQueryRoutineSpec_FromProto(mapCtx, in.GetBigqueryRoutineSpec())
	return out
}
func RoutineSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.RoutineSpec) *pb.RoutineSpec {
	if in == nil {
		return nil
	}
	out := &pb.RoutineSpec{}
	out.RoutineType = direct.Enum_ToProto[pb.RoutineSpec_RoutineType](mapCtx, in.RoutineType)
	out.Language = direct.ValueOf(in.Language)
	out.RoutineArguments = direct.Slice_ToProto(mapCtx, in.RoutineArguments, RoutineSpec_Argument_ToProto)
	out.ReturnType = direct.ValueOf(in.ReturnType)
	out.DefinitionBody = direct.ValueOf(in.DefinitionBody)
	if oneof := BigQueryRoutineSpec_ToProto(mapCtx, in.BigqueryRoutineSpec); oneof != nil {
		out.SystemSpec = &pb.RoutineSpec_BigqueryRoutineSpec{BigqueryRoutineSpec: oneof}
	}
	return out
}
func RoutineSpec_Argument_FromProto(mapCtx *direct.MapContext, in *pb.RoutineSpec_Argument) *krmv1alpha1.RoutineSpec_Argument {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.RoutineSpec_Argument{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func RoutineSpec_Argument_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.RoutineSpec_Argument) *pb.RoutineSpec_Argument {
	if in == nil {
		return nil
	}
	out := &pb.RoutineSpec_Argument{}
	out.Name = direct.ValueOf(in.Name)
	out.Mode = direct.Enum_ToProto[pb.RoutineSpec_Argument_Mode](mapCtx, in.Mode)
	out.Type = direct.ValueOf(in.Type)
	return out
}
func SQLDatabaseSystemSpec_FromProto(mapCtx *direct.MapContext, in *pb.SqlDatabaseSystemSpec) *krmv1alpha1.SQLDatabaseSystemSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SQLDatabaseSystemSpec{}
	out.SQLEngine = direct.LazyPtr(in.GetSqlEngine())
	out.DatabaseVersion = direct.LazyPtr(in.GetDatabaseVersion())
	out.InstanceHost = direct.LazyPtr(in.GetInstanceHost())
	return out
}
func SQLDatabaseSystemSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SQLDatabaseSystemSpec) *pb.SqlDatabaseSystemSpec {
	if in == nil {
		return nil
	}
	out := &pb.SqlDatabaseSystemSpec{}
	out.SqlEngine = direct.ValueOf(in.SQLEngine)
	out.DatabaseVersion = direct.ValueOf(in.DatabaseVersion)
	out.InstanceHost = direct.ValueOf(in.InstanceHost)
	return out
}
func Schema_FromProto(mapCtx *direct.MapContext, in *pb.Schema) *krmv1alpha1.Schema {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Schema{}
	out.Columns = direct.Slice_FromProto(mapCtx, in.Columns, ColumnSchema_FromProto)
	return out
}
func Schema_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Schema) *pb.Schema {
	if in == nil {
		return nil
	}
	out := &pb.Schema{}
	out.Columns = direct.Slice_ToProto(mapCtx, in.Columns, ColumnSchema_ToProto)
	return out
}
func ServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceSpec) *krmv1alpha1.ServiceSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ServiceSpec{}
	out.CloudBigtableInstanceSpec = CloudBigtableInstanceSpec_FromProto(mapCtx, in.GetCloudBigtableInstanceSpec())
	return out
}
func ServiceSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ServiceSpec) *pb.ServiceSpec {
	if in == nil {
		return nil
	}
	out := &pb.ServiceSpec{}
	if oneof := CloudBigtableInstanceSpec_ToProto(mapCtx, in.CloudBigtableInstanceSpec); oneof != nil {
		out.SystemSpec = &pb.ServiceSpec_CloudBigtableInstanceSpec{CloudBigtableInstanceSpec: oneof}
	}
	return out
}
func StorageProperties_FromProto(mapCtx *direct.MapContext, in *pb.StorageProperties) *krmv1alpha1.StorageProperties {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.StorageProperties{}
	out.FilePattern = in.FilePattern
	out.FileType = direct.LazyPtr(in.GetFileType())
	return out
}
func StorageProperties_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.StorageProperties) *pb.StorageProperties {
	if in == nil {
		return nil
	}
	out := &pb.StorageProperties{}
	out.FilePattern = in.FilePattern
	out.FileType = direct.ValueOf(in.FileType)
	return out
}
func SystemTimestamps_FromProto(mapCtx *direct.MapContext, in *pb.SystemTimestamps) *krmv1alpha1.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SystemTimestamps{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: ExpireTime
	return out
}
func SystemTimestamps_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SystemTimestamps) *pb.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &pb.SystemTimestamps{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: ExpireTime
	return out
}
func SystemTimestampsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SystemTimestamps) *krmv1alpha1.SystemTimestampsObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SystemTimestampsObservedState{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	return out
}
func SystemTimestampsObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SystemTimestampsObservedState) *pb.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &pb.SystemTimestamps{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	return out
}
func TableSpec_FromProto(mapCtx *direct.MapContext, in *pb.TableSpec) *krmv1alpha1.TableSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.TableSpec{}
	// MISSING: GroupedEntry
	return out
}
func TableSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TableSpec) *pb.TableSpec {
	if in == nil {
		return nil
	}
	out := &pb.TableSpec{}
	// MISSING: GroupedEntry
	return out
}
func TableSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TableSpec) *krmv1alpha1.TableSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.TableSpecObservedState{}
	out.GroupedEntry = direct.LazyPtr(in.GetGroupedEntry())
	return out
}
func TableSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TableSpecObservedState) *pb.TableSpec {
	if in == nil {
		return nil
	}
	out := &pb.TableSpec{}
	out.GroupedEntry = direct.ValueOf(in.GroupedEntry)
	return out
}
func TagTemplateField_FromProto(mapCtx *direct.MapContext, in *pb.TagTemplateField) *krmv1alpha1.TagTemplateField {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.TagTemplateField{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Type = FieldType_FromProto(mapCtx, in.GetType())
	out.IsRequired = direct.LazyPtr(in.GetIsRequired())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Order = direct.LazyPtr(in.GetOrder())
	return out
}
func TagTemplateField_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TagTemplateField) *pb.TagTemplateField {
	if in == nil {
		return nil
	}
	out := &pb.TagTemplateField{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Type = FieldType_ToProto(mapCtx, in.Type)
	out.IsRequired = direct.ValueOf(in.IsRequired)
	out.Description = direct.ValueOf(in.Description)
	out.Order = direct.ValueOf(in.Order)
	return out
}
func UsageSignal_FromProto(mapCtx *direct.MapContext, in *pb.UsageSignal) *krmv1alpha1.UsageSignal {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.UsageSignal{}
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: UsageWithinTimeRange
	// MISSING: CommonUsageWithinTimeRange
	out.FavoriteCount = in.FavoriteCount
	return out
}
func UsageSignal_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.UsageSignal) *pb.UsageSignal {
	if in == nil {
		return nil
	}
	out := &pb.UsageSignal{}
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: UsageWithinTimeRange
	// MISSING: CommonUsageWithinTimeRange
	out.FavoriteCount = in.FavoriteCount
	return out
}
func UsageSignalObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UsageSignal) *krmv1alpha1.UsageSignalObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.UsageSignalObservedState{}
	// MISSING: UpdateTime
	// MISSING: UsageWithinTimeRange
	// MISSING: CommonUsageWithinTimeRange
	// MISSING: FavoriteCount
	return out
}
func UsageSignalObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.UsageSignalObservedState) *pb.UsageSignal {
	if in == nil {
		return nil
	}
	out := &pb.UsageSignal{}
	// MISSING: UpdateTime
	// MISSING: UsageWithinTimeRange
	// MISSING: CommonUsageWithinTimeRange
	// MISSING: FavoriteCount
	return out
}
func UsageStats_FromProto(mapCtx *direct.MapContext, in *pb.UsageStats) *krmv1alpha1.UsageStats {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.UsageStats{}
	out.TotalCompletions = direct.LazyPtr(in.GetTotalCompletions())
	out.TotalFailures = direct.LazyPtr(in.GetTotalFailures())
	out.TotalCancellations = direct.LazyPtr(in.GetTotalCancellations())
	out.TotalExecutionTimeForCompletionsMillis = direct.LazyPtr(in.GetTotalExecutionTimeForCompletionsMillis())
	return out
}
func UsageStats_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.UsageStats) *pb.UsageStats {
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
func VertexDatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.VertexDatasetSpec) *krmv1alpha1.VertexDatasetSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VertexDatasetSpec{}
	out.DataItemCount = direct.LazyPtr(in.GetDataItemCount())
	out.DataType = direct.Enum_FromProto(mapCtx, in.GetDataType())
	return out
}
func VertexDatasetSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VertexDatasetSpec) *pb.VertexDatasetSpec {
	if in == nil {
		return nil
	}
	out := &pb.VertexDatasetSpec{}
	out.DataItemCount = direct.ValueOf(in.DataItemCount)
	out.DataType = direct.Enum_ToProto[pb.VertexDatasetSpec_DataType](mapCtx, in.DataType)
	return out
}
func VertexModelSourceInfo_FromProto(mapCtx *direct.MapContext, in *pb.VertexModelSourceInfo) *krmv1alpha1.VertexModelSourceInfo {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VertexModelSourceInfo{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.Copy = direct.LazyPtr(in.GetCopy())
	return out
}
func VertexModelSourceInfo_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VertexModelSourceInfo) *pb.VertexModelSourceInfo {
	if in == nil {
		return nil
	}
	out := &pb.VertexModelSourceInfo{}
	out.SourceType = direct.Enum_ToProto[pb.VertexModelSourceInfo_ModelSourceType](mapCtx, in.SourceType)
	out.Copy = direct.ValueOf(in.Copy)
	return out
}
func VertexModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.VertexModelSpec) *krmv1alpha1.VertexModelSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VertexModelSpec{}
	out.VersionID = direct.LazyPtr(in.GetVersionId())
	out.VersionAliases = in.VersionAliases
	out.VersionDescription = direct.LazyPtr(in.GetVersionDescription())
	out.VertexModelSourceInfo = VertexModelSourceInfo_FromProto(mapCtx, in.GetVertexModelSourceInfo())
	out.ContainerImageURI = direct.LazyPtr(in.GetContainerImageUri())
	return out
}
func VertexModelSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VertexModelSpec) *pb.VertexModelSpec {
	if in == nil {
		return nil
	}
	out := &pb.VertexModelSpec{}
	out.VersionId = direct.ValueOf(in.VersionID)
	out.VersionAliases = in.VersionAliases
	out.VersionDescription = direct.ValueOf(in.VersionDescription)
	out.VertexModelSourceInfo = VertexModelSourceInfo_ToProto(mapCtx, in.VertexModelSourceInfo)
	out.ContainerImageUri = direct.ValueOf(in.ContainerImageURI)
	return out
}
func ViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.ViewSpec) *krmv1alpha1.ViewSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ViewSpec{}
	// MISSING: ViewQuery
	return out
}
func ViewSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ViewSpec) *pb.ViewSpec {
	if in == nil {
		return nil
	}
	out := &pb.ViewSpec{}
	// MISSING: ViewQuery
	return out
}
func ViewSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ViewSpec) *krmv1alpha1.ViewSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ViewSpecObservedState{}
	out.ViewQuery = direct.LazyPtr(in.GetViewQuery())
	return out
}
func ViewSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ViewSpecObservedState) *pb.ViewSpec {
	if in == nil {
		return nil
	}
	out := &pb.ViewSpec{}
	out.ViewQuery = direct.ValueOf(in.ViewQuery)
	return out
}
