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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/bigquery/storage/apiv1beta2/storagepb"
)
func TableFieldSchema_FromProto(mapCtx *direct.MapContext, in *pb.TableFieldSchema) *krm.TableFieldSchema {
	if in == nil {
		return nil
	}
	out := &krm.TableFieldSchema{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, TableFieldSchema_FromProto)
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func TableFieldSchema_ToProto(mapCtx *direct.MapContext, in *krm.TableFieldSchema) *pb.TableFieldSchema {
	if in == nil {
		return nil
	}
	out := &pb.TableFieldSchema{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.Enum_ToProto[pb.TableFieldSchema_Type](mapCtx, in.Type)
	out.Mode = direct.Enum_ToProto[pb.TableFieldSchema_Mode](mapCtx, in.Mode)
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, TableFieldSchema_ToProto)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func TableSchema_FromProto(mapCtx *direct.MapContext, in *pb.TableSchema) *krm.TableSchema {
	if in == nil {
		return nil
	}
	out := &krm.TableSchema{}
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, TableFieldSchema_FromProto)
	return out
}
func TableSchema_ToProto(mapCtx *direct.MapContext, in *krm.TableSchema) *pb.TableSchema {
	if in == nil {
		return nil
	}
	out := &pb.TableSchema{}
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, TableFieldSchema_ToProto)
	return out
}
func WriteStream_FromProto(mapCtx *direct.MapContext, in *pb.WriteStream) *krm.WriteStream {
	if in == nil {
		return nil
	}
	out := &krm.WriteStream{}
	// MISSING: Name
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: CreateTime
	// MISSING: CommitTime
	// MISSING: TableSchema
	return out
}
func WriteStream_ToProto(mapCtx *direct.MapContext, in *krm.WriteStream) *pb.WriteStream {
	if in == nil {
		return nil
	}
	out := &pb.WriteStream{}
	// MISSING: Name
	out.Type = direct.Enum_ToProto[pb.WriteStream_Type](mapCtx, in.Type)
	// MISSING: CreateTime
	// MISSING: CommitTime
	// MISSING: TableSchema
	return out
}
func WriteStreamObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WriteStream) *krm.WriteStreamObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WriteStreamObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Type
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.CommitTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCommitTime())
	out.TableSchema = TableSchema_FromProto(mapCtx, in.GetTableSchema())
	return out
}
func WriteStreamObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WriteStreamObservedState) *pb.WriteStream {
	if in == nil {
		return nil
	}
	out := &pb.WriteStream{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Type
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.CommitTime = direct.StringTimestamp_ToProto(mapCtx, in.CommitTime)
	out.TableSchema = TableSchema_ToProto(mapCtx, in.TableSchema)
	return out
}
