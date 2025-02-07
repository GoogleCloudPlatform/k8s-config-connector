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
	pb "cloud.google.com/go/bigquery/storage/apiv1/storagepb"
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
func ArrowSerializationOptions_FromProto(mapCtx *direct.MapContext, in *pb.ArrowSerializationOptions) *krm.ArrowSerializationOptions {
	if in == nil {
		return nil
	}
	out := &krm.ArrowSerializationOptions{}
	out.BufferCompression = direct.Enum_FromProto(mapCtx, in.GetBufferCompression())
	return out
}
func ArrowSerializationOptions_ToProto(mapCtx *direct.MapContext, in *krm.ArrowSerializationOptions) *pb.ArrowSerializationOptions {
	if in == nil {
		return nil
	}
	out := &pb.ArrowSerializationOptions{}
	out.BufferCompression = direct.Enum_ToProto[pb.ArrowSerializationOptions_CompressionCodec](mapCtx, in.BufferCompression)
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
func AvroSerializationOptions_FromProto(mapCtx *direct.MapContext, in *pb.AvroSerializationOptions) *krm.AvroSerializationOptions {
	if in == nil {
		return nil
	}
	out := &krm.AvroSerializationOptions{}
	out.EnableDisplayNameAttribute = direct.LazyPtr(in.GetEnableDisplayNameAttribute())
	return out
}
func AvroSerializationOptions_ToProto(mapCtx *direct.MapContext, in *krm.AvroSerializationOptions) *pb.AvroSerializationOptions {
	if in == nil {
		return nil
	}
	out := &pb.AvroSerializationOptions{}
	out.EnableDisplayNameAttribute = direct.ValueOf(in.EnableDisplayNameAttribute)
	return out
}
func BigqueryReadSessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReadSession) *krm.BigqueryReadSessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryReadSessionObservedState{}
	// MISSING: Name
	// MISSING: ExpireTime
	// MISSING: DataFormat
	// MISSING: AvroSchema
	// MISSING: ArrowSchema
	// MISSING: Table
	// MISSING: TableModifiers
	// MISSING: ReadOptions
	// MISSING: Streams
	// MISSING: EstimatedTotalBytesScanned
	// MISSING: EstimatedTotalPhysicalFileSize
	// MISSING: EstimatedRowCount
	// MISSING: TraceID
	return out
}
func BigqueryReadSessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryReadSessionObservedState) *pb.ReadSession {
	if in == nil {
		return nil
	}
	out := &pb.ReadSession{}
	// MISSING: Name
	// MISSING: ExpireTime
	// MISSING: DataFormat
	// MISSING: AvroSchema
	// MISSING: ArrowSchema
	// MISSING: Table
	// MISSING: TableModifiers
	// MISSING: ReadOptions
	// MISSING: Streams
	// MISSING: EstimatedTotalBytesScanned
	// MISSING: EstimatedTotalPhysicalFileSize
	// MISSING: EstimatedRowCount
	// MISSING: TraceID
	return out
}
func BigqueryReadSessionSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReadSession) *krm.BigqueryReadSessionSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryReadSessionSpec{}
	// MISSING: Name
	// MISSING: ExpireTime
	// MISSING: DataFormat
	// MISSING: AvroSchema
	// MISSING: ArrowSchema
	// MISSING: Table
	// MISSING: TableModifiers
	// MISSING: ReadOptions
	// MISSING: Streams
	// MISSING: EstimatedTotalBytesScanned
	// MISSING: EstimatedTotalPhysicalFileSize
	// MISSING: EstimatedRowCount
	// MISSING: TraceID
	return out
}
func BigqueryReadSessionSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryReadSessionSpec) *pb.ReadSession {
	if in == nil {
		return nil
	}
	out := &pb.ReadSession{}
	// MISSING: Name
	// MISSING: ExpireTime
	// MISSING: DataFormat
	// MISSING: AvroSchema
	// MISSING: ArrowSchema
	// MISSING: Table
	// MISSING: TableModifiers
	// MISSING: ReadOptions
	// MISSING: Streams
	// MISSING: EstimatedTotalBytesScanned
	// MISSING: EstimatedTotalPhysicalFileSize
	// MISSING: EstimatedRowCount
	// MISSING: TraceID
	return out
}
func ReadSession_FromProto(mapCtx *direct.MapContext, in *pb.ReadSession) *krm.ReadSession {
	if in == nil {
		return nil
	}
	out := &krm.ReadSession{}
	// MISSING: Name
	// MISSING: ExpireTime
	out.DataFormat = direct.Enum_FromProto(mapCtx, in.GetDataFormat())
	// MISSING: AvroSchema
	// MISSING: ArrowSchema
	out.Table = direct.LazyPtr(in.GetTable())
	out.TableModifiers = ReadSession_TableModifiers_FromProto(mapCtx, in.GetTableModifiers())
	out.ReadOptions = ReadSession_TableReadOptions_FromProto(mapCtx, in.GetReadOptions())
	// MISSING: Streams
	// MISSING: EstimatedTotalBytesScanned
	// MISSING: EstimatedTotalPhysicalFileSize
	// MISSING: EstimatedRowCount
	out.TraceID = direct.LazyPtr(in.GetTraceId())
	return out
}
func ReadSession_ToProto(mapCtx *direct.MapContext, in *krm.ReadSession) *pb.ReadSession {
	if in == nil {
		return nil
	}
	out := &pb.ReadSession{}
	// MISSING: Name
	// MISSING: ExpireTime
	out.DataFormat = direct.Enum_ToProto[pb.DataFormat](mapCtx, in.DataFormat)
	// MISSING: AvroSchema
	// MISSING: ArrowSchema
	out.Table = direct.ValueOf(in.Table)
	out.TableModifiers = ReadSession_TableModifiers_ToProto(mapCtx, in.TableModifiers)
	out.ReadOptions = ReadSession_TableReadOptions_ToProto(mapCtx, in.ReadOptions)
	// MISSING: Streams
	// MISSING: EstimatedTotalBytesScanned
	// MISSING: EstimatedTotalPhysicalFileSize
	// MISSING: EstimatedRowCount
	out.TraceId = direct.ValueOf(in.TraceID)
	return out
}
func ReadSessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReadSession) *krm.ReadSessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ReadSessionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: DataFormat
	out.AvroSchema = AvroSchema_FromProto(mapCtx, in.GetAvroSchema())
	out.ArrowSchema = ArrowSchema_FromProto(mapCtx, in.GetArrowSchema())
	// MISSING: Table
	// MISSING: TableModifiers
	// MISSING: ReadOptions
	out.Streams = direct.Slice_FromProto(mapCtx, in.Streams, ReadStream_FromProto)
	out.EstimatedTotalBytesScanned = direct.LazyPtr(in.GetEstimatedTotalBytesScanned())
	out.EstimatedTotalPhysicalFileSize = direct.LazyPtr(in.GetEstimatedTotalPhysicalFileSize())
	out.EstimatedRowCount = direct.LazyPtr(in.GetEstimatedRowCount())
	// MISSING: TraceID
	return out
}
func ReadSessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReadSessionObservedState) *pb.ReadSession {
	if in == nil {
		return nil
	}
	out := &pb.ReadSession{}
	out.Name = direct.ValueOf(in.Name)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	// MISSING: DataFormat
	if oneof := AvroSchema_ToProto(mapCtx, in.AvroSchema); oneof != nil {
		out.Schema = &pb.ReadSession_AvroSchema{AvroSchema: oneof}
	}
	if oneof := ArrowSchema_ToProto(mapCtx, in.ArrowSchema); oneof != nil {
		out.Schema = &pb.ReadSession_ArrowSchema{ArrowSchema: oneof}
	}
	// MISSING: Table
	// MISSING: TableModifiers
	// MISSING: ReadOptions
	out.Streams = direct.Slice_ToProto(mapCtx, in.Streams, ReadStream_ToProto)
	out.EstimatedTotalBytesScanned = direct.ValueOf(in.EstimatedTotalBytesScanned)
	out.EstimatedTotalPhysicalFileSize = direct.ValueOf(in.EstimatedTotalPhysicalFileSize)
	out.EstimatedRowCount = direct.ValueOf(in.EstimatedRowCount)
	// MISSING: TraceID
	return out
}
func ReadSession_TableModifiers_FromProto(mapCtx *direct.MapContext, in *pb.ReadSession_TableModifiers) *krm.ReadSession_TableModifiers {
	if in == nil {
		return nil
	}
	out := &krm.ReadSession_TableModifiers{}
	out.SnapshotTime = direct.StringTimestamp_FromProto(mapCtx, in.GetSnapshotTime())
	return out
}
func ReadSession_TableModifiers_ToProto(mapCtx *direct.MapContext, in *krm.ReadSession_TableModifiers) *pb.ReadSession_TableModifiers {
	if in == nil {
		return nil
	}
	out := &pb.ReadSession_TableModifiers{}
	out.SnapshotTime = direct.StringTimestamp_ToProto(mapCtx, in.SnapshotTime)
	return out
}
func ReadSession_TableReadOptions_FromProto(mapCtx *direct.MapContext, in *pb.ReadSession_TableReadOptions) *krm.ReadSession_TableReadOptions {
	if in == nil {
		return nil
	}
	out := &krm.ReadSession_TableReadOptions{}
	out.SelectedFields = in.SelectedFields
	out.RowRestriction = direct.LazyPtr(in.GetRowRestriction())
	out.ArrowSerializationOptions = ArrowSerializationOptions_FromProto(mapCtx, in.GetArrowSerializationOptions())
	out.AvroSerializationOptions = AvroSerializationOptions_FromProto(mapCtx, in.GetAvroSerializationOptions())
	out.SamplePercentage = in.SamplePercentage
	out.ResponseCompressionCodec = direct.Enum_FromProto(mapCtx, in.GetResponseCompressionCodec())
	return out
}
func ReadSession_TableReadOptions_ToProto(mapCtx *direct.MapContext, in *krm.ReadSession_TableReadOptions) *pb.ReadSession_TableReadOptions {
	if in == nil {
		return nil
	}
	out := &pb.ReadSession_TableReadOptions{}
	out.SelectedFields = in.SelectedFields
	out.RowRestriction = direct.ValueOf(in.RowRestriction)
	if oneof := ArrowSerializationOptions_ToProto(mapCtx, in.ArrowSerializationOptions); oneof != nil {
		out.OutputFormatSerializationOptions = &pb.ReadSession_TableReadOptions_ArrowSerializationOptions{ArrowSerializationOptions: oneof}
	}
	if oneof := AvroSerializationOptions_ToProto(mapCtx, in.AvroSerializationOptions); oneof != nil {
		out.OutputFormatSerializationOptions = &pb.ReadSession_TableReadOptions_AvroSerializationOptions{AvroSerializationOptions: oneof}
	}
	out.SamplePercentage = in.SamplePercentage
	if oneof := ReadSession_TableReadOptions_ResponseCompressionCodec_ToProto(mapCtx, in.ResponseCompressionCodec); oneof != nil {
		out.ResponseCompressionCodec = oneof
	}
	return out
}
func ReadStream_FromProto(mapCtx *direct.MapContext, in *pb.ReadStream) *krm.ReadStream {
	if in == nil {
		return nil
	}
	out := &krm.ReadStream{}
	// MISSING: Name
	return out
}
func ReadStream_ToProto(mapCtx *direct.MapContext, in *krm.ReadStream) *pb.ReadStream {
	if in == nil {
		return nil
	}
	out := &pb.ReadStream{}
	// MISSING: Name
	return out
}
func ReadStreamObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReadStream) *krm.ReadStreamObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ReadStreamObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func ReadStreamObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReadStreamObservedState) *pb.ReadStream {
	if in == nil {
		return nil
	}
	out := &pb.ReadStream{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
