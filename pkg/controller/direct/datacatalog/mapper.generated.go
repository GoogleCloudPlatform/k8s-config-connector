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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigQueryConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryConnectionSpec) *krm.BigQueryConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryConnectionSpec{}
	out.ConnectionType = direct.Enum_FromProto(mapCtx, in.GetConnectionType())
	out.CloudSQL = CloudSQLBigQueryConnectionSpec_FromProto(mapCtx, in.GetCloudSql())
	out.HasCredential = direct.LazyPtr(in.GetHasCredential())
	return out
}
func BigQueryConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryConnectionSpec) *pb.BigQueryConnectionSpec {
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
func BigQueryDateShardedSpec_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDateShardedSpec) *krm.BigQueryDateShardedSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDateShardedSpec{}
	// MISSING: Dataset
	// MISSING: TablePrefix
	// MISSING: ShardCount
	// MISSING: LatestShardResource
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
	// MISSING: LatestShardResource
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
	out.LatestShardResource = direct.LazyPtr(in.GetLatestShardResource())
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
	out.LatestShardResource = direct.ValueOf(in.LatestShardResource)
	return out
}
func BigQueryRoutineSpec_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryRoutineSpec) *krm.BigQueryRoutineSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryRoutineSpec{}
	out.ImportedLibraries = in.ImportedLibraries
	return out
}
func BigQueryRoutineSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryRoutineSpec) *pb.BigQueryRoutineSpec {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryRoutineSpec{}
	out.ImportedLibraries = in.ImportedLibraries
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
func BusinessContext_FromProto(mapCtx *direct.MapContext, in *pb.BusinessContext) *krm.BusinessContext {
	if in == nil {
		return nil
	}
	out := &krm.BusinessContext{}
	out.EntryOverview = EntryOverview_FromProto(mapCtx, in.GetEntryOverview())
	out.Contacts = Contacts_FromProto(mapCtx, in.GetContacts())
	return out
}
func BusinessContext_ToProto(mapCtx *direct.MapContext, in *krm.BusinessContext) *pb.BusinessContext {
	if in == nil {
		return nil
	}
	out := &pb.BusinessContext{}
	out.EntryOverview = EntryOverview_ToProto(mapCtx, in.EntryOverview)
	out.Contacts = Contacts_ToProto(mapCtx, in.Contacts)
	return out
}
func CloudBigtableInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudBigtableInstanceSpec) *krm.CloudBigtableInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudBigtableInstanceSpec{}
	out.CloudBigtableClusterSpecs = direct.Slice_FromProto(mapCtx, in.CloudBigtableClusterSpecs, CloudBigtableInstanceSpec_CloudBigtableClusterSpec_FromProto)
	return out
}
func CloudBigtableInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudBigtableInstanceSpec) *pb.CloudBigtableInstanceSpec {
	if in == nil {
		return nil
	}
	out := &pb.CloudBigtableInstanceSpec{}
	out.CloudBigtableClusterSpecs = direct.Slice_ToProto(mapCtx, in.CloudBigtableClusterSpecs, CloudBigtableInstanceSpec_CloudBigtableClusterSpec_ToProto)
	return out
}
func CloudBigtableInstanceSpec_CloudBigtableClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudBigtableInstanceSpec_CloudBigtableClusterSpec) *krm.CloudBigtableInstanceSpec_CloudBigtableClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudBigtableInstanceSpec_CloudBigtableClusterSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Type = direct.LazyPtr(in.GetType())
	out.LinkedResource = direct.LazyPtr(in.GetLinkedResource())
	return out
}
func CloudBigtableInstanceSpec_CloudBigtableClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudBigtableInstanceSpec_CloudBigtableClusterSpec) *pb.CloudBigtableInstanceSpec_CloudBigtableClusterSpec {
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
func CloudBigtableSystemSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudBigtableSystemSpec) *krm.CloudBigtableSystemSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudBigtableSystemSpec{}
	out.InstanceDisplayName = direct.LazyPtr(in.GetInstanceDisplayName())
	return out
}
func CloudBigtableSystemSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudBigtableSystemSpec) *pb.CloudBigtableSystemSpec {
	if in == nil {
		return nil
	}
	out := &pb.CloudBigtableSystemSpec{}
	out.InstanceDisplayName = direct.ValueOf(in.InstanceDisplayName)
	return out
}
func CloudSQLBigQueryConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlBigQueryConnectionSpec) *krm.CloudSQLBigQueryConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudSQLBigQueryConnectionSpec{}
	out.InstanceID = direct.LazyPtr(in.GetInstanceId())
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func CloudSQLBigQueryConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudSQLBigQueryConnectionSpec) *pb.CloudSqlBigQueryConnectionSpec {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlBigQueryConnectionSpec{}
	out.InstanceId = direct.ValueOf(in.InstanceID)
	out.Database = direct.ValueOf(in.Database)
	out.Type = direct.Enum_ToProto[pb.CloudSqlBigQueryConnectionSpec_DatabaseType](mapCtx, in.Type)
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
	out.DefaultValue = direct.LazyPtr(in.GetDefaultValue())
	out.OrdinalPosition = direct.LazyPtr(in.GetOrdinalPosition())
	out.HighestIndexingType = direct.Enum_FromProto(mapCtx, in.GetHighestIndexingType())
	out.Subcolumns = direct.Slice_FromProto(mapCtx, in.Subcolumns, ColumnSchema_FromProto)
	out.LookerColumnSpec = ColumnSchema_LookerColumnSpec_FromProto(mapCtx, in.GetLookerColumnSpec())
	out.RangeElementType = ColumnSchema_FieldElementType_FromProto(mapCtx, in.GetRangeElementType())
	out.GcRule = direct.LazyPtr(in.GetGcRule())
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
func ColumnSchema_FieldElementType_FromProto(mapCtx *direct.MapContext, in *pb.ColumnSchema_FieldElementType) *krm.ColumnSchema_FieldElementType {
	if in == nil {
		return nil
	}
	out := &krm.ColumnSchema_FieldElementType{}
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func ColumnSchema_FieldElementType_ToProto(mapCtx *direct.MapContext, in *krm.ColumnSchema_FieldElementType) *pb.ColumnSchema_FieldElementType {
	if in == nil {
		return nil
	}
	out := &pb.ColumnSchema_FieldElementType{}
	out.Type = direct.ValueOf(in.Type)
	return out
}
func ColumnSchema_LookerColumnSpec_FromProto(mapCtx *direct.MapContext, in *pb.ColumnSchema_LookerColumnSpec) *krm.ColumnSchema_LookerColumnSpec {
	if in == nil {
		return nil
	}
	out := &krm.ColumnSchema_LookerColumnSpec{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func ColumnSchema_LookerColumnSpec_ToProto(mapCtx *direct.MapContext, in *krm.ColumnSchema_LookerColumnSpec) *pb.ColumnSchema_LookerColumnSpec {
	if in == nil {
		return nil
	}
	out := &pb.ColumnSchema_LookerColumnSpec{}
	out.Type = direct.Enum_ToProto[pb.ColumnSchema_LookerColumnSpec_LookerColumnType](mapCtx, in.Type)
	return out
}
func CommonUsageStats_FromProto(mapCtx *direct.MapContext, in *pb.CommonUsageStats) *krm.CommonUsageStats {
	if in == nil {
		return nil
	}
	out := &krm.CommonUsageStats{}
	out.ViewCount = in.ViewCount
	return out
}
func CommonUsageStats_ToProto(mapCtx *direct.MapContext, in *krm.CommonUsageStats) *pb.CommonUsageStats {
	if in == nil {
		return nil
	}
	out := &pb.CommonUsageStats{}
	out.ViewCount = in.ViewCount
	return out
}
func Contacts_FromProto(mapCtx *direct.MapContext, in *pb.Contacts) *krm.Contacts {
	if in == nil {
		return nil
	}
	out := &krm.Contacts{}
	out.People = direct.Slice_FromProto(mapCtx, in.People, Contacts_Person_FromProto)
	return out
}
func Contacts_ToProto(mapCtx *direct.MapContext, in *krm.Contacts) *pb.Contacts {
	if in == nil {
		return nil
	}
	out := &pb.Contacts{}
	out.People = direct.Slice_ToProto(mapCtx, in.People, Contacts_Person_ToProto)
	return out
}
func Contacts_Person_FromProto(mapCtx *direct.MapContext, in *pb.Contacts_Person) *krm.Contacts_Person {
	if in == nil {
		return nil
	}
	out := &krm.Contacts_Person{}
	out.Designation = direct.LazyPtr(in.GetDesignation())
	out.Email = direct.LazyPtr(in.GetEmail())
	return out
}
func Contacts_Person_ToProto(mapCtx *direct.MapContext, in *krm.Contacts_Person) *pb.Contacts_Person {
	if in == nil {
		return nil
	}
	out := &pb.Contacts_Person{}
	out.Designation = direct.ValueOf(in.Designation)
	out.Email = direct.ValueOf(in.Email)
	return out
}
func DataCatalogEntryGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.DataCatalogEntryGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogEntryGroupObservedState{}
	// MISSING: Name
	out.DataCatalogTimestamps = SystemTimestamps_FromProto(mapCtx, in.GetDataCatalogTimestamps())
	return out
}
func DataCatalogEntryGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogEntryGroupObservedState) *pb.EntryGroup {
	if in == nil {
		return nil
	}
	out := &pb.EntryGroup{}
	// MISSING: Name
	out.DataCatalogTimestamps = SystemTimestamps_ToProto(mapCtx, in.DataCatalogTimestamps)
	return out
}
func DataCatalogEntryGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.EntryGroup) *krm.DataCatalogEntryGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogEntryGroupSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.TransferredToDataplex = direct.LazyPtr(in.GetTransferredToDataplex())
	return out
}
func DataCatalogEntryGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogEntryGroupSpec) *pb.EntryGroup {
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
func DataCatalogEntryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entry) *krm.DataCatalogEntryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogEntryObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.IntegratedSystem = direct.Enum_FromProto(mapCtx, in.GetIntegratedSystem())
	out.GCSFilesetSpec = GCSFilesetSpecObservedState_FromProto(mapCtx, in.GetGcsFilesetSpec())
	out.BigqueryTableSpec = BigQueryTableSpec_FromProto(mapCtx, in.GetBigqueryTableSpec())
	out.BigqueryDateShardedSpec = BigQueryDateShardedSpec_FromProto(mapCtx, in.GetBigqueryDateShardedSpec())
	out.DatabaseTableSpec = DatabaseTableSpecObservedState_FromProto(mapCtx, in.GetDatabaseTableSpec())
	out.FeatureOnlineStoreSpec = FeatureOnlineStoreSpecObservedState_FromProto(mapCtx, in.GetFeatureOnlineStoreSpec())
	out.UsageSignal = UsageSignalObservedState_FromProto(mapCtx, in.GetUsageSignal())
	out.DataSource = DataSource_FromProto(mapCtx, in.GetDataSource())
	out.PersonalDetails = PersonalDetails_FromProto(mapCtx, in.GetPersonalDetails())
	return out
}
func DataCatalogEntryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogEntryObservedState) *pb.Entry {
	if in == nil {
		return nil
	}
	out := &pb.Entry{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := DataCatalogEntryObservedState_IntegratedSystem_ToProto(mapCtx, in.IntegratedSystem); oneof != nil {
		out.System = oneof
	}
	if oneof := GCSFilesetSpecObservedState_ToProto(mapCtx, in.GCSFilesetSpec); oneof != nil {
		out.TypeSpec = &pb.Entry_GcsFilesetSpec{GcsFilesetSpec: oneof}
	}
	if oneof := BigQueryTableSpec_ToProto(mapCtx, in.BigqueryTableSpec); oneof != nil {
		out.TypeSpec = &pb.Entry_BigqueryTableSpec{BigqueryTableSpec: oneof}
	}
	if oneof := BigQueryDateShardedSpec_ToProto(mapCtx, in.BigqueryDateShardedSpec); oneof != nil {
		out.TypeSpec = &pb.Entry_BigqueryDateShardedSpec{BigqueryDateShardedSpec: oneof}
	}
	if oneof := DatabaseTableSpecObservedState_ToProto(mapCtx, in.DatabaseTableSpec); oneof != nil {
		out.Spec = &pb.Entry_DatabaseTableSpec{DatabaseTableSpec: oneof}
	}
	if oneof := FeatureOnlineStoreSpecObservedState_ToProto(mapCtx, in.FeatureOnlineStoreSpec); oneof != nil {
		out.Spec = &pb.Entry_FeatureOnlineStoreSpec{FeatureOnlineStoreSpec: oneof}
	}
	out.UsageSignal = UsageSignalObservedState_ToProto(mapCtx, in.UsageSignal)
	out.DataSource = DataSource_ToProto(mapCtx, in.DataSource)
	out.PersonalDetails = PersonalDetails_ToProto(mapCtx, in.PersonalDetails)
	return out
}
func DataCatalogEntrySpec_FromProto(mapCtx *direct.MapContext, in *pb.Entry) *krm.DataCatalogEntrySpec {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogEntrySpec{}
	out.LinkedResource = direct.LazyPtr(in.GetLinkedResource())
	out.FullyQualifiedName = direct.LazyPtr(in.GetFullyQualifiedName())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.UserSpecifiedType = direct.LazyPtr(in.GetUserSpecifiedType())
	out.UserSpecifiedSystem = direct.LazyPtr(in.GetUserSpecifiedSystem())
	out.SQLDatabaseSystemSpec = SQLDatabaseSystemSpec_FromProto(mapCtx, in.GetSqlDatabaseSystemSpec())
	out.LookerSystemSpec = LookerSystemSpec_FromProto(mapCtx, in.GetLookerSystemSpec())
	out.CloudBigtableSystemSpec = CloudBigtableSystemSpec_FromProto(mapCtx, in.GetCloudBigtableSystemSpec())
	out.GCSFilesetSpec = GCSFilesetSpec_FromProto(mapCtx, in.GetGcsFilesetSpec())
	out.DatabaseTableSpec = DatabaseTableSpec_FromProto(mapCtx, in.GetDatabaseTableSpec())
	out.DataSourceConnectionSpec = DataSourceConnectionSpec_FromProto(mapCtx, in.GetDataSourceConnectionSpec())
	out.RoutineSpec = RoutineSpec_FromProto(mapCtx, in.GetRoutineSpec())
	out.DatasetSpec = DatasetSpec_FromProto(mapCtx, in.GetDatasetSpec())
	out.FilesetSpec = FilesetSpec_FromProto(mapCtx, in.GetFilesetSpec())
	out.ServiceSpec = ServiceSpec_FromProto(mapCtx, in.GetServiceSpec())
	out.ModelSpec = ModelSpec_FromProto(mapCtx, in.GetModelSpec())
	out.FeatureOnlineStoreSpec = FeatureOnlineStoreSpec_FromProto(mapCtx, in.GetFeatureOnlineStoreSpec())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.BusinessContext = BusinessContext_FromProto(mapCtx, in.GetBusinessContext())
	out.Schema = Schema_FromProto(mapCtx, in.GetSchema())
	out.SourceSystemTimestamps = SystemTimestamps_FromProto(mapCtx, in.GetSourceSystemTimestamps())
	out.UsageSignal = UsageSignal_FromProto(mapCtx, in.GetUsageSignal())
	out.Labels = in.Labels
	return out
}
func DataCatalogEntrySpec_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogEntrySpec) *pb.Entry {
	if in == nil {
		return nil
	}
	out := &pb.Entry{}
	out.LinkedResource = direct.ValueOf(in.LinkedResource)
	out.FullyQualifiedName = direct.ValueOf(in.FullyQualifiedName)
	if oneof := DataCatalogEntrySpec_Type_ToProto(mapCtx, in.Type); oneof != nil {
		out.EntryType = oneof
	}
	if oneof := DataCatalogEntrySpec_UserSpecifiedType_ToProto(mapCtx, in.UserSpecifiedType); oneof != nil {
		out.EntryType = oneof
	}
	if oneof := DataCatalogEntrySpec_UserSpecifiedSystem_ToProto(mapCtx, in.UserSpecifiedSystem); oneof != nil {
		out.System = oneof
	}
	if oneof := SQLDatabaseSystemSpec_ToProto(mapCtx, in.SQLDatabaseSystemSpec); oneof != nil {
		out.SystemSpec = &pb.Entry_SqlDatabaseSystemSpec{SqlDatabaseSystemSpec: oneof}
	}
	if oneof := LookerSystemSpec_ToProto(mapCtx, in.LookerSystemSpec); oneof != nil {
		out.SystemSpec = &pb.Entry_LookerSystemSpec{LookerSystemSpec: oneof}
	}
	if oneof := CloudBigtableSystemSpec_ToProto(mapCtx, in.CloudBigtableSystemSpec); oneof != nil {
		out.SystemSpec = &pb.Entry_CloudBigtableSystemSpec{CloudBigtableSystemSpec: oneof}
	}
	if oneof := GCSFilesetSpec_ToProto(mapCtx, in.GCSFilesetSpec); oneof != nil {
		out.TypeSpec = &pb.Entry_GcsFilesetSpec{GcsFilesetSpec: oneof}
	}
	if oneof := DatabaseTableSpec_ToProto(mapCtx, in.DatabaseTableSpec); oneof != nil {
		out.Spec = &pb.Entry_DatabaseTableSpec{DatabaseTableSpec: oneof}
	}
	if oneof := DataSourceConnectionSpec_ToProto(mapCtx, in.DataSourceConnectionSpec); oneof != nil {
		out.Spec = &pb.Entry_DataSourceConnectionSpec{DataSourceConnectionSpec: oneof}
	}
	if oneof := RoutineSpec_ToProto(mapCtx, in.RoutineSpec); oneof != nil {
		out.Spec = &pb.Entry_RoutineSpec{RoutineSpec: oneof}
	}
	if oneof := DatasetSpec_ToProto(mapCtx, in.DatasetSpec); oneof != nil {
		out.Spec = &pb.Entry_DatasetSpec{DatasetSpec: oneof}
	}
	if oneof := FilesetSpec_ToProto(mapCtx, in.FilesetSpec); oneof != nil {
		out.Spec = &pb.Entry_FilesetSpec{FilesetSpec: oneof}
	}
	if oneof := ServiceSpec_ToProto(mapCtx, in.ServiceSpec); oneof != nil {
		out.Spec = &pb.Entry_ServiceSpec{ServiceSpec: oneof}
	}
	if oneof := ModelSpec_ToProto(mapCtx, in.ModelSpec); oneof != nil {
		out.Spec = &pb.Entry_ModelSpec{ModelSpec: oneof}
	}
	if oneof := FeatureOnlineStoreSpec_ToProto(mapCtx, in.FeatureOnlineStoreSpec); oneof != nil {
		out.Spec = &pb.Entry_FeatureOnlineStoreSpec{FeatureOnlineStoreSpec: oneof}
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.BusinessContext = BusinessContext_ToProto(mapCtx, in.BusinessContext)
	out.Schema = Schema_ToProto(mapCtx, in.Schema)
	out.SourceSystemTimestamps = SystemTimestamps_ToProto(mapCtx, in.SourceSystemTimestamps)
	out.UsageSignal = UsageSignal_ToProto(mapCtx, in.UsageSignal)
	out.Labels = in.Labels
	return out
}
func DataSource_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krm.DataSource {
	if in == nil {
		return nil
	}
	out := &krm.DataSource{}
	out.Service = direct.Enum_FromProto(mapCtx, in.GetService())
	out.Resource = direct.LazyPtr(in.GetResource())
	// MISSING: SourceEntry
	out.StorageProperties = StorageProperties_FromProto(mapCtx, in.GetStorageProperties())
	return out
}
func DataSource_ToProto(mapCtx *direct.MapContext, in *krm.DataSource) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	out.Service = direct.Enum_ToProto[pb.DataSource_Service](mapCtx, in.Service)
	out.Resource = direct.ValueOf(in.Resource)
	// MISSING: SourceEntry
	if oneof := StorageProperties_ToProto(mapCtx, in.StorageProperties); oneof != nil {
		out.Properties = &pb.DataSource_StorageProperties{StorageProperties: oneof}
	}
	return out
}
func DataSourceConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataSourceConnectionSpec) *krm.DataSourceConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataSourceConnectionSpec{}
	out.BigqueryConnectionSpec = BigQueryConnectionSpec_FromProto(mapCtx, in.GetBigqueryConnectionSpec())
	return out
}
func DataSourceConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataSourceConnectionSpec) *pb.DataSourceConnectionSpec {
	if in == nil {
		return nil
	}
	out := &pb.DataSourceConnectionSpec{}
	out.BigqueryConnectionSpec = BigQueryConnectionSpec_ToProto(mapCtx, in.BigqueryConnectionSpec)
	return out
}
func DataSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krm.DataSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataSourceObservedState{}
	// MISSING: Service
	// MISSING: Resource
	out.SourceEntry = direct.LazyPtr(in.GetSourceEntry())
	// MISSING: StorageProperties
	return out
}
func DataSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataSourceObservedState) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	// MISSING: Service
	// MISSING: Resource
	out.SourceEntry = direct.ValueOf(in.SourceEntry)
	// MISSING: StorageProperties
	return out
}
func DatabaseTableSpec_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseTableSpec) *krm.DatabaseTableSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatabaseTableSpec{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: DataplexTable
	out.DatabaseViewSpec = DatabaseTableSpec_DatabaseViewSpec_FromProto(mapCtx, in.GetDatabaseViewSpec())
	return out
}
func DatabaseTableSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatabaseTableSpec) *pb.DatabaseTableSpec {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseTableSpec{}
	out.Type = direct.Enum_ToProto[pb.DatabaseTableSpec_TableType](mapCtx, in.Type)
	// MISSING: DataplexTable
	out.DatabaseViewSpec = DatabaseTableSpec_DatabaseViewSpec_ToProto(mapCtx, in.DatabaseViewSpec)
	return out
}
func DatabaseTableSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseTableSpec) *krm.DatabaseTableSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatabaseTableSpecObservedState{}
	// MISSING: Type
	out.DataplexTable = DataplexTableSpec_FromProto(mapCtx, in.GetDataplexTable())
	// MISSING: DatabaseViewSpec
	return out
}
func DatabaseTableSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatabaseTableSpecObservedState) *pb.DatabaseTableSpec {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseTableSpec{}
	// MISSING: Type
	out.DataplexTable = DataplexTableSpec_ToProto(mapCtx, in.DataplexTable)
	// MISSING: DatabaseViewSpec
	return out
}
func DatabaseTableSpec_DatabaseViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseTableSpec_DatabaseViewSpec) *krm.DatabaseTableSpec_DatabaseViewSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatabaseTableSpec_DatabaseViewSpec{}
	out.ViewType = direct.Enum_FromProto(mapCtx, in.GetViewType())
	out.BaseTable = direct.LazyPtr(in.GetBaseTable())
	out.SQLQuery = direct.LazyPtr(in.GetSqlQuery())
	return out
}
func DatabaseTableSpec_DatabaseViewSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatabaseTableSpec_DatabaseViewSpec) *pb.DatabaseTableSpec_DatabaseViewSpec {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseTableSpec_DatabaseViewSpec{}
	out.ViewType = direct.Enum_ToProto[pb.DatabaseTableSpec_DatabaseViewSpec_ViewType](mapCtx, in.ViewType)
	if oneof := DatabaseTableSpec_DatabaseViewSpec_BaseTable_ToProto(mapCtx, in.BaseTable); oneof != nil {
		out.SourceDefinition = oneof
	}
	if oneof := DatabaseTableSpec_DatabaseViewSpec_SqlQuery_ToProto(mapCtx, in.SQLQuery); oneof != nil {
		out.SourceDefinition = oneof
	}
	return out
}
func DataplexExternalTable_FromProto(mapCtx *direct.MapContext, in *pb.DataplexExternalTable) *krm.DataplexExternalTable {
	if in == nil {
		return nil
	}
	out := &krm.DataplexExternalTable{}
	out.System = direct.Enum_FromProto(mapCtx, in.GetSystem())
	out.FullyQualifiedName = direct.LazyPtr(in.GetFullyQualifiedName())
	out.GoogleCloudResource = direct.LazyPtr(in.GetGoogleCloudResource())
	out.DataCatalogEntry = direct.LazyPtr(in.GetDataCatalogEntry())
	return out
}
func DataplexExternalTable_ToProto(mapCtx *direct.MapContext, in *krm.DataplexExternalTable) *pb.DataplexExternalTable {
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
func DataplexFilesetSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataplexFilesetSpec) *krm.DataplexFilesetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexFilesetSpec{}
	out.DataplexSpec = DataplexSpec_FromProto(mapCtx, in.GetDataplexSpec())
	return out
}
func DataplexFilesetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexFilesetSpec) *pb.DataplexFilesetSpec {
	if in == nil {
		return nil
	}
	out := &pb.DataplexFilesetSpec{}
	out.DataplexSpec = DataplexSpec_ToProto(mapCtx, in.DataplexSpec)
	return out
}
func DataplexSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataplexSpec) *krm.DataplexSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexSpec{}
	out.Asset = direct.LazyPtr(in.GetAsset())
	out.DataFormat = PhysicalSchema_FromProto(mapCtx, in.GetDataFormat())
	out.CompressionFormat = direct.LazyPtr(in.GetCompressionFormat())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	return out
}
func DataplexSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexSpec) *pb.DataplexSpec {
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
func DataplexTableSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataplexTableSpec) *krm.DataplexTableSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexTableSpec{}
	out.ExternalTables = direct.Slice_FromProto(mapCtx, in.ExternalTables, DataplexExternalTable_FromProto)
	out.DataplexSpec = DataplexSpec_FromProto(mapCtx, in.GetDataplexSpec())
	out.UserManaged = direct.LazyPtr(in.GetUserManaged())
	return out
}
func DataplexTableSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexTableSpec) *pb.DataplexTableSpec {
	if in == nil {
		return nil
	}
	out := &pb.DataplexTableSpec{}
	out.ExternalTables = direct.Slice_ToProto(mapCtx, in.ExternalTables, DataplexExternalTable_ToProto)
	out.DataplexSpec = DataplexSpec_ToProto(mapCtx, in.DataplexSpec)
	out.UserManaged = direct.ValueOf(in.UserManaged)
	return out
}
func DatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.DatasetSpec) *krm.DatasetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatasetSpec{}
	out.VertexDatasetSpec = VertexDatasetSpec_FromProto(mapCtx, in.GetVertexDatasetSpec())
	return out
}
func DatasetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatasetSpec) *pb.DatasetSpec {
	if in == nil {
		return nil
	}
	out := &pb.DatasetSpec{}
	if oneof := VertexDatasetSpec_ToProto(mapCtx, in.VertexDatasetSpec); oneof != nil {
		out.SystemSpec = &pb.DatasetSpec_VertexDatasetSpec{VertexDatasetSpec: oneof}
	}
	return out
}
func EntryOverview_FromProto(mapCtx *direct.MapContext, in *pb.EntryOverview) *krm.EntryOverview {
	if in == nil {
		return nil
	}
	out := &krm.EntryOverview{}
	out.Overview = direct.LazyPtr(in.GetOverview())
	return out
}
func EntryOverview_ToProto(mapCtx *direct.MapContext, in *krm.EntryOverview) *pb.EntryOverview {
	if in == nil {
		return nil
	}
	out := &pb.EntryOverview{}
	out.Overview = direct.ValueOf(in.Overview)
	return out
}
func FeatureOnlineStoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.FeatureOnlineStoreSpec) *krm.FeatureOnlineStoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.FeatureOnlineStoreSpec{}
	// MISSING: StorageType
	return out
}
func FeatureOnlineStoreSpec_ToProto(mapCtx *direct.MapContext, in *krm.FeatureOnlineStoreSpec) *pb.FeatureOnlineStoreSpec {
	if in == nil {
		return nil
	}
	out := &pb.FeatureOnlineStoreSpec{}
	// MISSING: StorageType
	return out
}
func FeatureOnlineStoreSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FeatureOnlineStoreSpec) *krm.FeatureOnlineStoreSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FeatureOnlineStoreSpecObservedState{}
	out.StorageType = direct.Enum_FromProto(mapCtx, in.GetStorageType())
	return out
}
func FeatureOnlineStoreSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FeatureOnlineStoreSpecObservedState) *pb.FeatureOnlineStoreSpec {
	if in == nil {
		return nil
	}
	out := &pb.FeatureOnlineStoreSpec{}
	out.StorageType = direct.Enum_ToProto[pb.FeatureOnlineStoreSpec_StorageType](mapCtx, in.StorageType)
	return out
}
func FilesetSpec_FromProto(mapCtx *direct.MapContext, in *pb.FilesetSpec) *krm.FilesetSpec {
	if in == nil {
		return nil
	}
	out := &krm.FilesetSpec{}
	out.DataplexFileset = DataplexFilesetSpec_FromProto(mapCtx, in.GetDataplexFileset())
	return out
}
func FilesetSpec_ToProto(mapCtx *direct.MapContext, in *krm.FilesetSpec) *pb.FilesetSpec {
	if in == nil {
		return nil
	}
	out := &pb.FilesetSpec{}
	out.DataplexFileset = DataplexFilesetSpec_ToProto(mapCtx, in.DataplexFileset)
	return out
}
func GCSFileSpec_FromProto(mapCtx *direct.MapContext, in *pb.GcsFileSpec) *krm.GCSFileSpec {
	if in == nil {
		return nil
	}
	out := &krm.GCSFileSpec{}
	out.FilePath = direct.LazyPtr(in.GetFilePath())
	// MISSING: GCSTimestamps
	// MISSING: SizeBytes
	return out
}
func GCSFileSpec_ToProto(mapCtx *direct.MapContext, in *krm.GCSFileSpec) *pb.GcsFileSpec {
	if in == nil {
		return nil
	}
	out := &pb.GcsFileSpec{}
	out.FilePath = direct.ValueOf(in.FilePath)
	// MISSING: GCSTimestamps
	// MISSING: SizeBytes
	return out
}
func GCSFileSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GcsFileSpec) *krm.GCSFileSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GCSFileSpecObservedState{}
	// MISSING: FilePath
	out.GCSTimestamps = SystemTimestamps_FromProto(mapCtx, in.GetGcsTimestamps())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	return out
}
func GCSFileSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GCSFileSpecObservedState) *pb.GcsFileSpec {
	if in == nil {
		return nil
	}
	out := &pb.GcsFileSpec{}
	// MISSING: FilePath
	out.GcsTimestamps = SystemTimestamps_ToProto(mapCtx, in.GCSTimestamps)
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	return out
}
func GCSFilesetSpec_FromProto(mapCtx *direct.MapContext, in *pb.GcsFilesetSpec) *krm.GCSFilesetSpec {
	if in == nil {
		return nil
	}
	out := &krm.GCSFilesetSpec{}
	out.FilePatterns = in.FilePatterns
	// MISSING: SampleGCSFileSpecs
	return out
}
func GCSFilesetSpec_ToProto(mapCtx *direct.MapContext, in *krm.GCSFilesetSpec) *pb.GcsFilesetSpec {
	if in == nil {
		return nil
	}
	out := &pb.GcsFilesetSpec{}
	out.FilePatterns = in.FilePatterns
	// MISSING: SampleGCSFileSpecs
	return out
}
func GCSFilesetSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GcsFilesetSpec) *krm.GCSFilesetSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GCSFilesetSpecObservedState{}
	// MISSING: FilePatterns
	out.SampleGCSFileSpecs = direct.Slice_FromProto(mapCtx, in.SampleGCSFileSpecs, GCSFileSpec_FromProto)
	return out
}
func GCSFilesetSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GCSFilesetSpecObservedState) *pb.GcsFilesetSpec {
	if in == nil {
		return nil
	}
	out := &pb.GcsFilesetSpec{}
	// MISSING: FilePatterns
	out.SampleGcsFileSpecs = direct.Slice_ToProto(mapCtx, in.SampleGCSFileSpecs, GCSFileSpec_ToProto)
	return out
}
func LookerSystemSpec_FromProto(mapCtx *direct.MapContext, in *pb.LookerSystemSpec) *krm.LookerSystemSpec {
	if in == nil {
		return nil
	}
	out := &krm.LookerSystemSpec{}
	out.ParentInstanceID = direct.LazyPtr(in.GetParentInstanceId())
	out.ParentInstanceDisplayName = direct.LazyPtr(in.GetParentInstanceDisplayName())
	out.ParentModelID = direct.LazyPtr(in.GetParentModelId())
	out.ParentModelDisplayName = direct.LazyPtr(in.GetParentModelDisplayName())
	out.ParentViewID = direct.LazyPtr(in.GetParentViewId())
	out.ParentViewDisplayName = direct.LazyPtr(in.GetParentViewDisplayName())
	return out
}
func LookerSystemSpec_ToProto(mapCtx *direct.MapContext, in *krm.LookerSystemSpec) *pb.LookerSystemSpec {
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
func ModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.ModelSpec) *krm.ModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.ModelSpec{}
	out.VertexModelSpec = VertexModelSpec_FromProto(mapCtx, in.GetVertexModelSpec())
	return out
}
func ModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.ModelSpec) *pb.ModelSpec {
	if in == nil {
		return nil
	}
	out := &pb.ModelSpec{}
	if oneof := VertexModelSpec_ToProto(mapCtx, in.VertexModelSpec); oneof != nil {
		out.SystemSpec = &pb.ModelSpec_VertexModelSpec{VertexModelSpec: oneof}
	}
	return out
}
func PersonalDetails_FromProto(mapCtx *direct.MapContext, in *pb.PersonalDetails) *krm.PersonalDetails {
	if in == nil {
		return nil
	}
	out := &krm.PersonalDetails{}
	out.Starred = direct.LazyPtr(in.GetStarred())
	out.StarTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStarTime())
	return out
}
func PersonalDetails_ToProto(mapCtx *direct.MapContext, in *krm.PersonalDetails) *pb.PersonalDetails {
	if in == nil {
		return nil
	}
	out := &pb.PersonalDetails{}
	out.Starred = direct.ValueOf(in.Starred)
	out.StarTime = direct.StringTimestamp_ToProto(mapCtx, in.StarTime)
	return out
}
func PhysicalSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema) *krm.PhysicalSchema {
	if in == nil {
		return nil
	}
	out := &krm.PhysicalSchema{}
	out.Avro = PhysicalSchema_AvroSchema_FromProto(mapCtx, in.GetAvro())
	out.Thrift = PhysicalSchema_ThriftSchema_FromProto(mapCtx, in.GetThrift())
	out.Protobuf = PhysicalSchema_ProtobufSchema_FromProto(mapCtx, in.GetProtobuf())
	out.Parquet = PhysicalSchema_ParquetSchema_FromProto(mapCtx, in.GetParquet())
	out.Orc = PhysicalSchema_OrcSchema_FromProto(mapCtx, in.GetOrc())
	out.Csv = PhysicalSchema_CsvSchema_FromProto(mapCtx, in.GetCsv())
	return out
}
func PhysicalSchema_ToProto(mapCtx *direct.MapContext, in *krm.PhysicalSchema) *pb.PhysicalSchema {
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
func PhysicalSchema_AvroSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema_AvroSchema) *krm.PhysicalSchema_AvroSchema {
	if in == nil {
		return nil
	}
	out := &krm.PhysicalSchema_AvroSchema{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func PhysicalSchema_AvroSchema_ToProto(mapCtx *direct.MapContext, in *krm.PhysicalSchema_AvroSchema) *pb.PhysicalSchema_AvroSchema {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalSchema_AvroSchema{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
func PhysicalSchema_CsvSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema_CsvSchema) *krm.PhysicalSchema_CsvSchema {
	if in == nil {
		return nil
	}
	out := &krm.PhysicalSchema_CsvSchema{}
	return out
}
func PhysicalSchema_CsvSchema_ToProto(mapCtx *direct.MapContext, in *krm.PhysicalSchema_CsvSchema) *pb.PhysicalSchema_CsvSchema {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalSchema_CsvSchema{}
	return out
}
func PhysicalSchema_OrcSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema_OrcSchema) *krm.PhysicalSchema_OrcSchema {
	if in == nil {
		return nil
	}
	out := &krm.PhysicalSchema_OrcSchema{}
	return out
}
func PhysicalSchema_OrcSchema_ToProto(mapCtx *direct.MapContext, in *krm.PhysicalSchema_OrcSchema) *pb.PhysicalSchema_OrcSchema {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalSchema_OrcSchema{}
	return out
}
func PhysicalSchema_ParquetSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema_ParquetSchema) *krm.PhysicalSchema_ParquetSchema {
	if in == nil {
		return nil
	}
	out := &krm.PhysicalSchema_ParquetSchema{}
	return out
}
func PhysicalSchema_ParquetSchema_ToProto(mapCtx *direct.MapContext, in *krm.PhysicalSchema_ParquetSchema) *pb.PhysicalSchema_ParquetSchema {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalSchema_ParquetSchema{}
	return out
}
func PhysicalSchema_ProtobufSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema_ProtobufSchema) *krm.PhysicalSchema_ProtobufSchema {
	if in == nil {
		return nil
	}
	out := &krm.PhysicalSchema_ProtobufSchema{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func PhysicalSchema_ProtobufSchema_ToProto(mapCtx *direct.MapContext, in *krm.PhysicalSchema_ProtobufSchema) *pb.PhysicalSchema_ProtobufSchema {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalSchema_ProtobufSchema{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
func PhysicalSchema_ThriftSchema_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalSchema_ThriftSchema) *krm.PhysicalSchema_ThriftSchema {
	if in == nil {
		return nil
	}
	out := &krm.PhysicalSchema_ThriftSchema{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func PhysicalSchema_ThriftSchema_ToProto(mapCtx *direct.MapContext, in *krm.PhysicalSchema_ThriftSchema) *pb.PhysicalSchema_ThriftSchema {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalSchema_ThriftSchema{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
func RoutineSpec_FromProto(mapCtx *direct.MapContext, in *pb.RoutineSpec) *krm.RoutineSpec {
	if in == nil {
		return nil
	}
	out := &krm.RoutineSpec{}
	out.RoutineType = direct.Enum_FromProto(mapCtx, in.GetRoutineType())
	out.Language = direct.LazyPtr(in.GetLanguage())
	out.RoutineArguments = direct.Slice_FromProto(mapCtx, in.RoutineArguments, RoutineSpec_Argument_FromProto)
	out.ReturnType = direct.LazyPtr(in.GetReturnType())
	out.DefinitionBody = direct.LazyPtr(in.GetDefinitionBody())
	out.BigqueryRoutineSpec = BigQueryRoutineSpec_FromProto(mapCtx, in.GetBigqueryRoutineSpec())
	return out
}
func RoutineSpec_ToProto(mapCtx *direct.MapContext, in *krm.RoutineSpec) *pb.RoutineSpec {
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
func RoutineSpec_Argument_FromProto(mapCtx *direct.MapContext, in *pb.RoutineSpec_Argument) *krm.RoutineSpec_Argument {
	if in == nil {
		return nil
	}
	out := &krm.RoutineSpec_Argument{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func RoutineSpec_Argument_ToProto(mapCtx *direct.MapContext, in *krm.RoutineSpec_Argument) *pb.RoutineSpec_Argument {
	if in == nil {
		return nil
	}
	out := &pb.RoutineSpec_Argument{}
	out.Name = direct.ValueOf(in.Name)
	out.Mode = direct.Enum_ToProto[pb.RoutineSpec_Argument_Mode](mapCtx, in.Mode)
	out.Type = direct.ValueOf(in.Type)
	return out
}
func SQLDatabaseSystemSpec_FromProto(mapCtx *direct.MapContext, in *pb.SqlDatabaseSystemSpec) *krm.SQLDatabaseSystemSpec {
	if in == nil {
		return nil
	}
	out := &krm.SQLDatabaseSystemSpec{}
	out.SQLEngine = direct.LazyPtr(in.GetSqlEngine())
	out.DatabaseVersion = direct.LazyPtr(in.GetDatabaseVersion())
	out.InstanceHost = direct.LazyPtr(in.GetInstanceHost())
	return out
}
func SQLDatabaseSystemSpec_ToProto(mapCtx *direct.MapContext, in *krm.SQLDatabaseSystemSpec) *pb.SqlDatabaseSystemSpec {
	if in == nil {
		return nil
	}
	out := &pb.SqlDatabaseSystemSpec{}
	out.SqlEngine = direct.ValueOf(in.SQLEngine)
	out.DatabaseVersion = direct.ValueOf(in.DatabaseVersion)
	out.InstanceHost = direct.ValueOf(in.InstanceHost)
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
func ServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceSpec) *krm.ServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ServiceSpec{}
	out.CloudBigtableInstanceSpec = CloudBigtableInstanceSpec_FromProto(mapCtx, in.GetCloudBigtableInstanceSpec())
	return out
}
func ServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ServiceSpec) *pb.ServiceSpec {
	if in == nil {
		return nil
	}
	out := &pb.ServiceSpec{}
	if oneof := CloudBigtableInstanceSpec_ToProto(mapCtx, in.CloudBigtableInstanceSpec); oneof != nil {
		out.SystemSpec = &pb.ServiceSpec_CloudBigtableInstanceSpec{CloudBigtableInstanceSpec: oneof}
	}
	return out
}
func StorageProperties_FromProto(mapCtx *direct.MapContext, in *pb.StorageProperties) *krm.StorageProperties {
	if in == nil {
		return nil
	}
	out := &krm.StorageProperties{}
	out.FilePattern = in.FilePattern
	out.FileType = direct.LazyPtr(in.GetFileType())
	return out
}
func StorageProperties_ToProto(mapCtx *direct.MapContext, in *krm.StorageProperties) *pb.StorageProperties {
	if in == nil {
		return nil
	}
	out := &pb.StorageProperties{}
	out.FilePattern = in.FilePattern
	out.FileType = direct.ValueOf(in.FileType)
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
	// MISSING: CommonUsageWithinTimeRange
	out.FavoriteCount = in.FavoriteCount
	return out
}
func UsageSignal_ToProto(mapCtx *direct.MapContext, in *krm.UsageSignal) *pb.UsageSignal {
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
func UsageSignalObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UsageSignal) *krm.UsageSignalObservedState {
	if in == nil {
		return nil
	}
	out := &krm.UsageSignalObservedState{}
	// MISSING: UpdateTime
	// MISSING: UsageWithinTimeRange
	// MISSING: CommonUsageWithinTimeRange
	// MISSING: FavoriteCount
	return out
}
func UsageSignalObservedState_ToProto(mapCtx *direct.MapContext, in *krm.UsageSignalObservedState) *pb.UsageSignal {
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
func VertexDatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.VertexDatasetSpec) *krm.VertexDatasetSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexDatasetSpec{}
	out.DataItemCount = direct.LazyPtr(in.GetDataItemCount())
	out.DataType = direct.Enum_FromProto(mapCtx, in.GetDataType())
	return out
}
func VertexDatasetSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexDatasetSpec) *pb.VertexDatasetSpec {
	if in == nil {
		return nil
	}
	out := &pb.VertexDatasetSpec{}
	out.DataItemCount = direct.ValueOf(in.DataItemCount)
	out.DataType = direct.Enum_ToProto[pb.VertexDatasetSpec_DataType](mapCtx, in.DataType)
	return out
}
func VertexModelSourceInfo_FromProto(mapCtx *direct.MapContext, in *pb.VertexModelSourceInfo) *krm.VertexModelSourceInfo {
	if in == nil {
		return nil
	}
	out := &krm.VertexModelSourceInfo{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.Copy = direct.LazyPtr(in.GetCopy())
	return out
}
func VertexModelSourceInfo_ToProto(mapCtx *direct.MapContext, in *krm.VertexModelSourceInfo) *pb.VertexModelSourceInfo {
	if in == nil {
		return nil
	}
	out := &pb.VertexModelSourceInfo{}
	out.SourceType = direct.Enum_ToProto[pb.VertexModelSourceInfo_ModelSourceType](mapCtx, in.SourceType)
	out.Copy = direct.ValueOf(in.Copy)
	return out
}
func VertexModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.VertexModelSpec) *krm.VertexModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexModelSpec{}
	out.VersionID = direct.LazyPtr(in.GetVersionId())
	out.VersionAliases = in.VersionAliases
	out.VersionDescription = direct.LazyPtr(in.GetVersionDescription())
	out.VertexModelSourceInfo = VertexModelSourceInfo_FromProto(mapCtx, in.GetVertexModelSourceInfo())
	out.ContainerImageURI = direct.LazyPtr(in.GetContainerImageUri())
	return out
}
func VertexModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexModelSpec) *pb.VertexModelSpec {
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
