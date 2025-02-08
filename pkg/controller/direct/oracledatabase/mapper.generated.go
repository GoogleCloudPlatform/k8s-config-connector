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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
)
func DbNode_FromProto(mapCtx *direct.MapContext, in *pb.DbNode) *krm.DbNode {
	if in == nil {
		return nil
	}
	out := &krm.DbNode{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Properties = DbNodeProperties_FromProto(mapCtx, in.GetProperties())
	return out
}
func DbNode_ToProto(mapCtx *direct.MapContext, in *krm.DbNode) *pb.DbNode {
	if in == nil {
		return nil
	}
	out := &pb.DbNode{}
	out.Name = direct.ValueOf(in.Name)
	out.Properties = DbNodeProperties_ToProto(mapCtx, in.Properties)
	return out
}
func DbNodeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DbNode) *krm.DbNodeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DbNodeObservedState{}
	// MISSING: Name
	out.Properties = DbNodePropertiesObservedState_FromProto(mapCtx, in.GetProperties())
	return out
}
func DbNodeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DbNodeObservedState) *pb.DbNode {
	if in == nil {
		return nil
	}
	out := &pb.DbNode{}
	// MISSING: Name
	out.Properties = DbNodePropertiesObservedState_ToProto(mapCtx, in.Properties)
	return out
}
func DbNodeProperties_FromProto(mapCtx *direct.MapContext, in *pb.DbNodeProperties) *krm.DbNodeProperties {
	if in == nil {
		return nil
	}
	out := &krm.DbNodeProperties{}
	// MISSING: Ocid
	out.OcpuCount = direct.LazyPtr(in.GetOcpuCount())
	out.MemorySizeGB = direct.LazyPtr(in.GetMemorySizeGb())
	out.DbNodeStorageSizeGB = direct.LazyPtr(in.GetDbNodeStorageSizeGb())
	out.DbServerOcid = direct.LazyPtr(in.GetDbServerOcid())
	out.Hostname = direct.LazyPtr(in.GetHostname())
	// MISSING: State
	out.TotalCpuCoreCount = direct.LazyPtr(in.GetTotalCpuCoreCount())
	return out
}
func DbNodeProperties_ToProto(mapCtx *direct.MapContext, in *krm.DbNodeProperties) *pb.DbNodeProperties {
	if in == nil {
		return nil
	}
	out := &pb.DbNodeProperties{}
	// MISSING: Ocid
	out.OcpuCount = direct.ValueOf(in.OcpuCount)
	out.MemorySizeGb = direct.ValueOf(in.MemorySizeGB)
	out.DbNodeStorageSizeGb = direct.ValueOf(in.DbNodeStorageSizeGB)
	out.DbServerOcid = direct.ValueOf(in.DbServerOcid)
	out.Hostname = direct.ValueOf(in.Hostname)
	// MISSING: State
	out.TotalCpuCoreCount = direct.ValueOf(in.TotalCpuCoreCount)
	return out
}
func DbNodePropertiesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DbNodeProperties) *krm.DbNodePropertiesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DbNodePropertiesObservedState{}
	out.Ocid = direct.LazyPtr(in.GetOcid())
	// MISSING: OcpuCount
	// MISSING: MemorySizeGB
	// MISSING: DbNodeStorageSizeGB
	// MISSING: DbServerOcid
	// MISSING: Hostname
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: TotalCpuCoreCount
	return out
}
func DbNodePropertiesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DbNodePropertiesObservedState) *pb.DbNodeProperties {
	if in == nil {
		return nil
	}
	out := &pb.DbNodeProperties{}
	out.Ocid = direct.ValueOf(in.Ocid)
	// MISSING: OcpuCount
	// MISSING: MemorySizeGB
	// MISSING: DbNodeStorageSizeGB
	// MISSING: DbServerOcid
	// MISSING: Hostname
	out.State = direct.Enum_ToProto[pb.DbNodeProperties_State](mapCtx, in.State)
	// MISSING: TotalCpuCoreCount
	return out
}
func OracledatabaseDbNodeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DbNode) *krm.OracledatabaseDbNodeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseDbNodeObservedState{}
	// MISSING: Name
	// MISSING: Properties
	return out
}
func OracledatabaseDbNodeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseDbNodeObservedState) *pb.DbNode {
	if in == nil {
		return nil
	}
	out := &pb.DbNode{}
	// MISSING: Name
	// MISSING: Properties
	return out
}
func OracledatabaseDbNodeSpec_FromProto(mapCtx *direct.MapContext, in *pb.DbNode) *krm.OracledatabaseDbNodeSpec {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseDbNodeSpec{}
	// MISSING: Name
	// MISSING: Properties
	return out
}
func OracledatabaseDbNodeSpec_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseDbNodeSpec) *pb.DbNode {
	if in == nil {
		return nil
	}
	out := &pb.DbNode{}
	// MISSING: Name
	// MISSING: Properties
	return out
}
