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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataCatalogEntryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entry) *krm.DataCatalogEntryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogEntryObservedState{}
	// MISSING: Name
	out.IntegratedSystem = direct.Enum_FromProto(mapCtx, in.GetIntegratedSystem())
	out.GCSFilesetSpec = GCSFilesetSpecObservedState_FromProto(mapCtx, in.GetGcsFilesetSpec())
	out.BigqueryTableSpec = BigQueryTableSpecObservedState_FromProto(mapCtx, in.GetBigqueryTableSpec())
	out.BigqueryDateShardedSpec = BigQueryDateShardedSpecObservedState_FromProto(mapCtx, in.GetBigqueryDateShardedSpec())
	out.DatabaseTableSpec = DatabaseTableSpecObservedState_FromProto(mapCtx, in.GetDatabaseTableSpec())
	out.FeatureOnlineStoreSpec = FeatureOnlineStoreSpecObservedState_FromProto(mapCtx, in.GetFeatureOnlineStoreSpec())
	out.UsageSignal = UsageSignalObservedState_FromProto(mapCtx, in.GetUsageSignal())
	out.DataSource = DataSourceObservedState_FromProto(mapCtx, in.GetDataSource())
	out.PersonalDetails = PersonalDetails_FromProto(mapCtx, in.GetPersonalDetails())
	return out
}
func DataCatalogEntryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogEntryObservedState) *pb.Entry {
	if in == nil {
		return nil
	}
	out := &pb.Entry{}
	// MISSING: Name
	if val := direct.Enum_ToProto[pb.IntegratedSystem](mapCtx, in.IntegratedSystem); val != pb.IntegratedSystem_INTEGRATED_SYSTEM_UNSPECIFIED {
		out.System = &pb.Entry_IntegratedSystem{IntegratedSystem: val}
	}
	if oneof := GCSFilesetSpecObservedState_ToProto(mapCtx, in.GCSFilesetSpec); oneof != nil {
		out.TypeSpec = &pb.Entry_GcsFilesetSpec{GcsFilesetSpec: oneof}
	}
	if oneof := BigQueryTableSpecObservedState_ToProto(mapCtx, in.BigqueryTableSpec); oneof != nil {
		out.TypeSpec = &pb.Entry_BigqueryTableSpec{BigqueryTableSpec: oneof}
	}
	if oneof := BigQueryDateShardedSpecObservedState_ToProto(mapCtx, in.BigqueryDateShardedSpec); oneof != nil {
		out.TypeSpec = &pb.Entry_BigqueryDateShardedSpec{BigqueryDateShardedSpec: oneof}
	}
	if oneof := DatabaseTableSpecObservedState_ToProto(mapCtx, in.DatabaseTableSpec); oneof != nil {
		out.Spec = &pb.Entry_DatabaseTableSpec{DatabaseTableSpec: oneof}
	}
	if oneof := FeatureOnlineStoreSpecObservedState_ToProto(mapCtx, in.FeatureOnlineStoreSpec); oneof != nil {
		out.Spec = &pb.Entry_FeatureOnlineStoreSpec{FeatureOnlineStoreSpec: oneof}
	}
	out.UsageSignal = UsageSignalObservedState_ToProto(mapCtx, in.UsageSignal)
	out.DataSource = DataSourceObservedState_ToProto(mapCtx, in.DataSource)
	out.PersonalDetails = PersonalDetails_ToProto(mapCtx, in.PersonalDetails)
	return out
}
func DataCatalogEntrySpec_FromProto(mapCtx *direct.MapContext, in *pb.Entry) *krm.DataCatalogEntrySpec {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogEntrySpec{}
	// MISSING: FullyQualifiedName - removed on purpose
	out.LinkedResource = direct.LazyPtr(in.GetLinkedResource())
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
	// MISSING: FullyQualifiedName
	out.LinkedResource = direct.ValueOf(in.LinkedResource)
	if val := direct.Enum_ToProto[pb.EntryType](mapCtx, in.Type); val != pb.EntryType_ENTRY_TYPE_UNSPECIFIED {
		out.EntryType = &pb.Entry_Type{Type: val}
	}
	if val := direct.ValueOf(in.UserSpecifiedType); val != "" {
		out.EntryType = &pb.Entry_UserSpecifiedType{UserSpecifiedType: val}
	}
	if val := direct.ValueOf(in.UserSpecifiedSystem); val != "" {
		out.System = &pb.Entry_UserSpecifiedSystem{UserSpecifiedSystem: val}
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
