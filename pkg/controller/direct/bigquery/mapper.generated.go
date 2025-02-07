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
	pb "cloud.google.com/go/bigquery/storage/apiv1beta1/storagepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ArrowSchema_FromProto(mapCtx *direct.MapContext, in *pb.ArrowSchema) *krm.ArrowSchema {
	if in == nil {
		return nil
	}
	out := &krm.ArrowSchema{}
	out.SerializedSchema = in.GetSerializedSchema()
	return out
}
func ArrowSchema_ToProto(mapCtx *direct.MapContext, in *krm.ArrowSchema) *pb.ArrowSchema {
	if in == nil {
		return nil
	}
	out := &pb.ArrowSchema{}
	out.SerializedSchema = in.SerializedSchema
	return out
}
func AvroSchema_FromProto(mapCtx *direct.MapContext, in *pb.AvroSchema) *krm.AvroSchema {
	if in == nil {
		return nil
	}
	out := &krm.AvroSchema{}
	out.Schema = direct.LazyPtr(in.GetSchema())
	return out
}
func AvroSchema_ToProto(mapCtx *direct.MapContext, in *krm.AvroSchema) *pb.AvroSchema {
	if in == nil {
		return nil
	}
	out := &pb.AvroSchema{}
	out.Schema = direct.ValueOf(in.Schema)
	return out
}
func ReadSession_FromProto(mapCtx *direct.MapContext, in *pb.ReadSession) *krm.ReadSession {
	if in == nil {
		return nil
	}
	out := &krm.ReadSession{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.AvroSchema = AvroSchema_FromProto(mapCtx, in.GetAvroSchema())
	out.ArrowSchema = ArrowSchema_FromProto(mapCtx, in.GetArrowSchema())
	out.Streams = direct.Slice_FromProto(mapCtx, in.Streams, Stream_FromProto)
	out.TableReference = TableReference_FromProto(mapCtx, in.GetTableReference())
	out.TableModifiers = TableModifiers_FromProto(mapCtx, in.GetTableModifiers())
	out.ShardingStrategy = direct.Enum_FromProto(mapCtx, in.GetShardingStrategy())
	return out
}
func ReadSession_ToProto(mapCtx *direct.MapContext, in *krm.ReadSession) *pb.ReadSession {
	if in == nil {
		return nil
	}
	out := &pb.ReadSession{}
	out.Name = direct.ValueOf(in.Name)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	if oneof := AvroSchema_ToProto(mapCtx, in.AvroSchema); oneof != nil {
		out.Schema = &pb.ReadSession_AvroSchema{AvroSchema: oneof}
	}
	if oneof := ArrowSchema_ToProto(mapCtx, in.ArrowSchema); oneof != nil {
		out.Schema = &pb.ReadSession_ArrowSchema{ArrowSchema: oneof}
	}
	out.Streams = direct.Slice_ToProto(mapCtx, in.Streams, Stream_ToProto)
	out.TableReference = TableReference_ToProto(mapCtx, in.TableReference)
	out.TableModifiers = TableModifiers_ToProto(mapCtx, in.TableModifiers)
	out.ShardingStrategy = direct.Enum_ToProto[pb.ShardingStrategy](mapCtx, in.ShardingStrategy)
	return out
}
func Stream_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.Stream {
	if in == nil {
		return nil
	}
	out := &krm.Stream{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Stream_ToProto(mapCtx *direct.MapContext, in *krm.Stream) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func TableModifiers_FromProto(mapCtx *direct.MapContext, in *pb.TableModifiers) *krm.TableModifiers {
	if in == nil {
		return nil
	}
	out := &krm.TableModifiers{}
	out.SnapshotTime = direct.StringTimestamp_FromProto(mapCtx, in.GetSnapshotTime())
	return out
}
func TableModifiers_ToProto(mapCtx *direct.MapContext, in *krm.TableModifiers) *pb.TableModifiers {
	if in == nil {
		return nil
	}
	out := &pb.TableModifiers{}
	out.SnapshotTime = direct.StringTimestamp_ToProto(mapCtx, in.SnapshotTime)
	return out
}
func TableReference_FromProto(mapCtx *direct.MapContext, in *pb.TableReference) *krm.TableReference {
	if in == nil {
		return nil
	}
	out := &krm.TableReference{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.DatasetID = direct.LazyPtr(in.GetDatasetId())
	out.TableID = direct.LazyPtr(in.GetTableId())
	return out
}
func TableReference_ToProto(mapCtx *direct.MapContext, in *krm.TableReference) *pb.TableReference {
	if in == nil {
		return nil
	}
	out := &pb.TableReference{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.DatasetId = direct.ValueOf(in.DatasetID)
	out.TableId = direct.ValueOf(in.TableID)
	return out
}
