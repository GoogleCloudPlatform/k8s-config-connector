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
func BigqueryReadStreamObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReadStream) *krm.BigqueryReadStreamObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryReadStreamObservedState{}
	// MISSING: Name
	return out
}
func BigqueryReadStreamObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryReadStreamObservedState) *pb.ReadStream {
	if in == nil {
		return nil
	}
	out := &pb.ReadStream{}
	// MISSING: Name
	return out
}
func BigqueryReadStreamSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReadStream) *krm.BigqueryReadStreamSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryReadStreamSpec{}
	// MISSING: Name
	return out
}
func BigqueryReadStreamSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryReadStreamSpec) *pb.ReadStream {
	if in == nil {
		return nil
	}
	out := &pb.ReadStream{}
	// MISSING: Name
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
