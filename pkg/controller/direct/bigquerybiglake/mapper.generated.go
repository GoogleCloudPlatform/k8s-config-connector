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

package bigquerybiglake

import (
	pb "cloud.google.com/go/bigquery/biglake/apiv1/biglakepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func HiveTableOptions_FromProto(mapCtx *direct.MapContext, in *pb.HiveTableOptions) *krm.HiveTableOptions {
	if in == nil {
		return nil
	}
	out := &krm.HiveTableOptions{}
	out.Parameters = in.Parameters
	out.TableType = direct.LazyPtr(in.GetTableType())
	out.StorageDescriptor = HiveTableOptions_StorageDescriptor_FromProto(mapCtx, in.GetStorageDescriptor())
	return out
}
func HiveTableOptions_ToProto(mapCtx *direct.MapContext, in *krm.HiveTableOptions) *pb.HiveTableOptions {
	if in == nil {
		return nil
	}
	out := &pb.HiveTableOptions{}
	out.Parameters = in.Parameters
	out.TableType = direct.ValueOf(in.TableType)
	out.StorageDescriptor = HiveTableOptions_StorageDescriptor_ToProto(mapCtx, in.StorageDescriptor)
	return out
}
func HiveTableOptions_SerDeInfo_FromProto(mapCtx *direct.MapContext, in *pb.HiveTableOptions_SerDeInfo) *krm.HiveTableOptions_SerDeInfo {
	if in == nil {
		return nil
	}
	out := &krm.HiveTableOptions_SerDeInfo{}
	out.SerializationLib = direct.LazyPtr(in.GetSerializationLib())
	return out
}
func HiveTableOptions_SerDeInfo_ToProto(mapCtx *direct.MapContext, in *krm.HiveTableOptions_SerDeInfo) *pb.HiveTableOptions_SerDeInfo {
	if in == nil {
		return nil
	}
	out := &pb.HiveTableOptions_SerDeInfo{}
	out.SerializationLib = direct.ValueOf(in.SerializationLib)
	return out
}
func HiveTableOptions_StorageDescriptor_FromProto(mapCtx *direct.MapContext, in *pb.HiveTableOptions_StorageDescriptor) *krm.HiveTableOptions_StorageDescriptor {
	if in == nil {
		return nil
	}
	out := &krm.HiveTableOptions_StorageDescriptor{}
	out.LocationURI = direct.LazyPtr(in.GetLocationUri())
	out.InputFormat = direct.LazyPtr(in.GetInputFormat())
	out.OutputFormat = direct.LazyPtr(in.GetOutputFormat())
	out.SerdeInfo = HiveTableOptions_SerDeInfo_FromProto(mapCtx, in.GetSerdeInfo())
	return out
}
func HiveTableOptions_StorageDescriptor_ToProto(mapCtx *direct.MapContext, in *krm.HiveTableOptions_StorageDescriptor) *pb.HiveTableOptions_StorageDescriptor {
	if in == nil {
		return nil
	}
	out := &pb.HiveTableOptions_StorageDescriptor{}
	out.LocationUri = direct.ValueOf(in.LocationURI)
	out.InputFormat = direct.ValueOf(in.InputFormat)
	out.OutputFormat = direct.ValueOf(in.OutputFormat)
	out.SerdeInfo = HiveTableOptions_SerDeInfo_ToProto(mapCtx, in.SerdeInfo)
	return out
}
func Table_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krm.Table {
	if in == nil {
		return nil
	}
	out := &krm.Table{}
	out.HiveOptions = HiveTableOptions_FromProto(mapCtx, in.GetHiveOptions())
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func Table_ToProto(mapCtx *direct.MapContext, in *krm.Table) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	if oneof := HiveTableOptions_ToProto(mapCtx, in.HiveOptions); oneof != nil {
		out.Options = &pb.Table_HiveOptions{HiveOptions: oneof}
	}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	out.Type = direct.Enum_ToProto[pb.Table_Type](mapCtx, in.Type)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func TableObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krm.TableObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TableObservedState{}
	// MISSING: HiveOptions
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: Type
	// MISSING: Etag
	return out
}
func TableObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TableObservedState) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	// MISSING: HiveOptions
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	// MISSING: Type
	// MISSING: Etag
	return out
}
