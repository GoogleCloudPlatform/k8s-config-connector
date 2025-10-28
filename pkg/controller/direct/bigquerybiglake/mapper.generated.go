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
// krm.group: bigquerybiglake.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.bigquery.biglake.v1

package bigquerybiglake

import (
	pb "cloud.google.com/go/bigquery/biglake/apiv1/biglakepb"
	krmbigquerybiglakev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigLakeCatalogObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Catalog) *krmbigquerybiglakev1alpha1.BigLakeCatalogObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigquerybiglakev1alpha1.BigLakeCatalogObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	return out
}
func BigLakeCatalogObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigquerybiglakev1alpha1.BigLakeCatalogObservedState) *pb.Catalog {
	if in == nil {
		return nil
	}
	out := &pb.Catalog{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	return out
}
func BigLakeCatalogSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Catalog) *krmbigquerybiglakev1alpha1.BigLakeCatalogSpec {
	if in == nil {
		return nil
	}
	out := &krmbigquerybiglakev1alpha1.BigLakeCatalogSpec{}
	// MISSING: Name
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	return out
}
func BigLakeCatalogSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigquerybiglakev1alpha1.BigLakeCatalogSpec) *pb.Catalog {
	if in == nil {
		return nil
	}
	out := &pb.Catalog{}
	// MISSING: Name
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	return out
}
func BigLakeDatabaseObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krmbigquerybiglakev1alpha1.BigLakeDatabaseObservedState {
	if in == nil {
		return nil
	}
	out := &krmbigquerybiglakev1alpha1.BigLakeDatabaseObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	return out
}
func BigLakeDatabaseObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigquerybiglakev1alpha1.BigLakeDatabaseObservedState) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	return out
}
func BigLakeDatabaseSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krmbigquerybiglakev1alpha1.BigLakeDatabaseSpec {
	if in == nil {
		return nil
	}
	out := &krmbigquerybiglakev1alpha1.BigLakeDatabaseSpec{}
	out.HiveOptions = HiveDatabaseOptions_v1alpha1_FromProto(mapCtx, in.GetHiveOptions())
	// MISSING: Name
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func BigLakeDatabaseSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigquerybiglakev1alpha1.BigLakeDatabaseSpec) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	if oneof := HiveDatabaseOptions_v1alpha1_ToProto(mapCtx, in.HiveOptions); oneof != nil {
		out.Options = &pb.Database_HiveOptions{HiveOptions: oneof}
	}
	// MISSING: Name
	out.Type = direct.Enum_ToProto[pb.Database_Type](mapCtx, in.Type)
	return out
}
func BigLakeTableObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krm.BigLakeTableObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigLakeTableObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: Etag
	return out
}
func BigLakeTableObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BigLakeTableObservedState) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	// MISSING: Etag
	return out
}
func BigLakeTableSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krm.BigLakeTableSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigLakeTableSpec{}
	out.HiveOptions = HiveTableOptions_v1beta1_FromProto(mapCtx, in.GetHiveOptions())
	// MISSING: Name
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: Etag
	return out
}
func BigLakeTableSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BigLakeTableSpec) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	if oneof := HiveTableOptions_v1beta1_ToProto(mapCtx, in.HiveOptions); oneof != nil {
		out.Options = &pb.Table_HiveOptions{HiveOptions: oneof}
	}
	// MISSING: Name
	out.Type = direct.Enum_ToProto[pb.Table_Type](mapCtx, in.Type)
	// MISSING: Etag
	return out
}
func HiveDatabaseOptions_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.HiveDatabaseOptions) *krmbigquerybiglakev1alpha1.HiveDatabaseOptions {
	if in == nil {
		return nil
	}
	out := &krmbigquerybiglakev1alpha1.HiveDatabaseOptions{}
	out.LocationURI = direct.LazyPtr(in.GetLocationUri())
	out.Parameters = in.Parameters
	return out
}
func HiveDatabaseOptions_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigquerybiglakev1alpha1.HiveDatabaseOptions) *pb.HiveDatabaseOptions {
	if in == nil {
		return nil
	}
	out := &pb.HiveDatabaseOptions{}
	out.LocationUri = direct.ValueOf(in.LocationURI)
	out.Parameters = in.Parameters
	return out
}
func HiveTableOptions_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HiveTableOptions) *krm.HiveTableOptions {
	if in == nil {
		return nil
	}
	out := &krm.HiveTableOptions{}
	out.Parameters = in.Parameters
	out.TableType = direct.LazyPtr(in.GetTableType())
	out.StorageDescriptor = HiveTableOptions_StorageDescriptor_v1beta1_FromProto(mapCtx, in.GetStorageDescriptor())
	return out
}
func HiveTableOptions_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HiveTableOptions) *pb.HiveTableOptions {
	if in == nil {
		return nil
	}
	out := &pb.HiveTableOptions{}
	out.Parameters = in.Parameters
	out.TableType = direct.ValueOf(in.TableType)
	out.StorageDescriptor = HiveTableOptions_StorageDescriptor_v1beta1_ToProto(mapCtx, in.StorageDescriptor)
	return out
}
func HiveTableOptions_SerDeInfo_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HiveTableOptions_SerDeInfo) *krm.HiveTableOptions_SerDeInfo {
	if in == nil {
		return nil
	}
	out := &krm.HiveTableOptions_SerDeInfo{}
	out.SerializationLib = direct.LazyPtr(in.GetSerializationLib())
	return out
}
func HiveTableOptions_SerDeInfo_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HiveTableOptions_SerDeInfo) *pb.HiveTableOptions_SerDeInfo {
	if in == nil {
		return nil
	}
	out := &pb.HiveTableOptions_SerDeInfo{}
	out.SerializationLib = direct.ValueOf(in.SerializationLib)
	return out
}
func HiveTableOptions_StorageDescriptor_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HiveTableOptions_StorageDescriptor) *krm.HiveTableOptions_StorageDescriptor {
	if in == nil {
		return nil
	}
	out := &krm.HiveTableOptions_StorageDescriptor{}
	out.LocationURI = direct.LazyPtr(in.GetLocationUri())
	out.InputFormat = direct.LazyPtr(in.GetInputFormat())
	out.OutputFormat = direct.LazyPtr(in.GetOutputFormat())
	out.SerdeInfo = HiveTableOptions_SerDeInfo_v1beta1_FromProto(mapCtx, in.GetSerdeInfo())
	return out
}
func HiveTableOptions_StorageDescriptor_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HiveTableOptions_StorageDescriptor) *pb.HiveTableOptions_StorageDescriptor {
	if in == nil {
		return nil
	}
	out := &pb.HiveTableOptions_StorageDescriptor{}
	out.LocationUri = direct.ValueOf(in.LocationURI)
	out.InputFormat = direct.ValueOf(in.InputFormat)
	out.OutputFormat = direct.ValueOf(in.OutputFormat)
	out.SerdeInfo = HiveTableOptions_SerDeInfo_v1beta1_ToProto(mapCtx, in.SerdeInfo)
	return out
}
