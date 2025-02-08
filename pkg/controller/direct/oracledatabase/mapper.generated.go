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

package oracledatabase

import (
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func DbServer_FromProto(mapCtx *direct.MapContext, in *pb.DbServer) *krm.DbServer {
	if in == nil {
		return nil
	}
	out := &krm.DbServer{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Properties = DbServerProperties_FromProto(mapCtx, in.GetProperties())
	return out
}
func DbServer_ToProto(mapCtx *direct.MapContext, in *krm.DbServer) *pb.DbServer {
	if in == nil {
		return nil
	}
	out := &pb.DbServer{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Properties = DbServerProperties_ToProto(mapCtx, in.Properties)
	return out
}
func DbServerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DbServer) *krm.DbServerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DbServerObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Properties = DbServerPropertiesObservedState_FromProto(mapCtx, in.GetProperties())
	return out
}
func DbServerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DbServerObservedState) *pb.DbServer {
	if in == nil {
		return nil
	}
	out := &pb.DbServer{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Properties = DbServerPropertiesObservedState_ToProto(mapCtx, in.Properties)
	return out
}
func DbServerProperties_FromProto(mapCtx *direct.MapContext, in *pb.DbServerProperties) *krm.DbServerProperties {
	if in == nil {
		return nil
	}
	out := &krm.DbServerProperties{}
	// MISSING: Ocid
	out.OcpuCount = direct.LazyPtr(in.GetOcpuCount())
	out.MaxOcpuCount = direct.LazyPtr(in.GetMaxOcpuCount())
	out.MemorySizeGB = direct.LazyPtr(in.GetMemorySizeGb())
	out.MaxMemorySizeGB = direct.LazyPtr(in.GetMaxMemorySizeGb())
	out.DbNodeStorageSizeGB = direct.LazyPtr(in.GetDbNodeStorageSizeGb())
	out.MaxDbNodeStorageSizeGB = direct.LazyPtr(in.GetMaxDbNodeStorageSizeGb())
	out.VmCount = direct.LazyPtr(in.GetVmCount())
	// MISSING: State
	// MISSING: DbNodeIds
	return out
}
func DbServerProperties_ToProto(mapCtx *direct.MapContext, in *krm.DbServerProperties) *pb.DbServerProperties {
	if in == nil {
		return nil
	}
	out := &pb.DbServerProperties{}
	// MISSING: Ocid
	out.OcpuCount = direct.ValueOf(in.OcpuCount)
	out.MaxOcpuCount = direct.ValueOf(in.MaxOcpuCount)
	out.MemorySizeGb = direct.ValueOf(in.MemorySizeGB)
	out.MaxMemorySizeGb = direct.ValueOf(in.MaxMemorySizeGB)
	out.DbNodeStorageSizeGb = direct.ValueOf(in.DbNodeStorageSizeGB)
	out.MaxDbNodeStorageSizeGb = direct.ValueOf(in.MaxDbNodeStorageSizeGB)
	out.VmCount = direct.ValueOf(in.VmCount)
	// MISSING: State
	// MISSING: DbNodeIds
	return out
}
func DbServerPropertiesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DbServerProperties) *krm.DbServerPropertiesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DbServerPropertiesObservedState{}
	out.Ocid = direct.LazyPtr(in.GetOcid())
	// MISSING: OcpuCount
	// MISSING: MaxOcpuCount
	// MISSING: MemorySizeGB
	// MISSING: MaxMemorySizeGB
	// MISSING: DbNodeStorageSizeGB
	// MISSING: MaxDbNodeStorageSizeGB
	// MISSING: VmCount
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.DbNodeIds = in.DbNodeIds
	return out
}
func DbServerPropertiesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DbServerPropertiesObservedState) *pb.DbServerProperties {
	if in == nil {
		return nil
	}
	out := &pb.DbServerProperties{}
	out.Ocid = direct.ValueOf(in.Ocid)
	// MISSING: OcpuCount
	// MISSING: MaxOcpuCount
	// MISSING: MemorySizeGB
	// MISSING: MaxMemorySizeGB
	// MISSING: DbNodeStorageSizeGB
	// MISSING: MaxDbNodeStorageSizeGB
	// MISSING: VmCount
	out.State = direct.Enum_ToProto[pb.DbServerProperties_State](mapCtx, in.State)
	out.DbNodeIds = in.DbNodeIds
	return out
}
func OracledatabaseDbServerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DbServer) *krm.OracledatabaseDbServerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseDbServerObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Properties
	return out
}
func OracledatabaseDbServerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseDbServerObservedState) *pb.DbServer {
	if in == nil {
		return nil
	}
	out := &pb.DbServer{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Properties
	return out
}
func OracledatabaseDbServerSpec_FromProto(mapCtx *direct.MapContext, in *pb.DbServer) *krm.OracledatabaseDbServerSpec {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseDbServerSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Properties
	return out
}
func OracledatabaseDbServerSpec_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseDbServerSpec) *pb.DbServer {
	if in == nil {
		return nil
	}
	out := &pb.DbServer{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Properties
	return out
}
