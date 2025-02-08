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

package vmwareengine

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func LoggingServer_FromProto(mapCtx *direct.MapContext, in *pb.LoggingServer) *krm.LoggingServer {
	if in == nil {
		return nil
	}
	out := &krm.LoggingServer{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Protocol = direct.Enum_FromProto(mapCtx, in.GetProtocol())
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	// MISSING: Uid
	return out
}
func LoggingServer_ToProto(mapCtx *direct.MapContext, in *krm.LoggingServer) *pb.LoggingServer {
	if in == nil {
		return nil
	}
	out := &pb.LoggingServer{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Port = direct.ValueOf(in.Port)
	out.Protocol = direct.Enum_ToProto[pb.LoggingServer_Protocol](mapCtx, in.Protocol)
	out.SourceType = direct.Enum_ToProto[pb.LoggingServer_SourceType](mapCtx, in.SourceType)
	// MISSING: Uid
	return out
}
func LoggingServerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LoggingServer) *krm.LoggingServerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LoggingServerObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Protocol
	// MISSING: SourceType
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}
func LoggingServerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LoggingServerObservedState) *pb.LoggingServer {
	if in == nil {
		return nil
	}
	out := &pb.LoggingServer{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Protocol
	// MISSING: SourceType
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
func VmwareengineLoggingServerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LoggingServer) *krm.VmwareengineLoggingServerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineLoggingServerObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Protocol
	// MISSING: SourceType
	// MISSING: Uid
	return out
}
func VmwareengineLoggingServerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineLoggingServerObservedState) *pb.LoggingServer {
	if in == nil {
		return nil
	}
	out := &pb.LoggingServer{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Protocol
	// MISSING: SourceType
	// MISSING: Uid
	return out
}
func VmwareengineLoggingServerSpec_FromProto(mapCtx *direct.MapContext, in *pb.LoggingServer) *krm.VmwareengineLoggingServerSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineLoggingServerSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Protocol
	// MISSING: SourceType
	// MISSING: Uid
	return out
}
func VmwareengineLoggingServerSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineLoggingServerSpec) *pb.LoggingServer {
	if in == nil {
		return nil
	}
	out := &pb.LoggingServer{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Hostname
	// MISSING: Port
	// MISSING: Protocol
	// MISSING: SourceType
	// MISSING: Uid
	return out
}
