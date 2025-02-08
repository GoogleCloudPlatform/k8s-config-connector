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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
)
func NodeType_FromProto(mapCtx *direct.MapContext, in *pb.NodeType) *krm.NodeType {
	if in == nil {
		return nil
	}
	out := &krm.NodeType{}
	// MISSING: Name
	// MISSING: NodeTypeID
	// MISSING: DisplayName
	// MISSING: VirtualCpuCount
	// MISSING: TotalCoreCount
	// MISSING: MemoryGB
	// MISSING: DiskSizeGB
	// MISSING: AvailableCustomCoreCounts
	// MISSING: Kind
	// MISSING: Families
	// MISSING: Capabilities
	return out
}
func NodeType_ToProto(mapCtx *direct.MapContext, in *krm.NodeType) *pb.NodeType {
	if in == nil {
		return nil
	}
	out := &pb.NodeType{}
	// MISSING: Name
	// MISSING: NodeTypeID
	// MISSING: DisplayName
	// MISSING: VirtualCpuCount
	// MISSING: TotalCoreCount
	// MISSING: MemoryGB
	// MISSING: DiskSizeGB
	// MISSING: AvailableCustomCoreCounts
	// MISSING: Kind
	// MISSING: Families
	// MISSING: Capabilities
	return out
}
func NodeTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NodeType) *krm.NodeTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NodeTypeObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.NodeTypeID = direct.LazyPtr(in.GetNodeTypeId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.VirtualCpuCount = direct.LazyPtr(in.GetVirtualCpuCount())
	out.TotalCoreCount = direct.LazyPtr(in.GetTotalCoreCount())
	out.MemoryGB = direct.LazyPtr(in.GetMemoryGb())
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	out.AvailableCustomCoreCounts = in.AvailableCustomCoreCounts
	out.Kind = direct.Enum_FromProto(mapCtx, in.GetKind())
	out.Families = in.Families
	out.Capabilities = direct.EnumSlice_FromProto(mapCtx, in.Capabilities)
	return out
}
func NodeTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NodeTypeObservedState) *pb.NodeType {
	if in == nil {
		return nil
	}
	out := &pb.NodeType{}
	out.Name = direct.ValueOf(in.Name)
	out.NodeTypeId = direct.ValueOf(in.NodeTypeID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.VirtualCpuCount = direct.ValueOf(in.VirtualCpuCount)
	out.TotalCoreCount = direct.ValueOf(in.TotalCoreCount)
	out.MemoryGb = direct.ValueOf(in.MemoryGB)
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	out.AvailableCustomCoreCounts = in.AvailableCustomCoreCounts
	out.Kind = direct.Enum_ToProto[pb.NodeType_Kind](mapCtx, in.Kind)
	out.Families = in.Families
	out.Capabilities = direct.EnumSlice_ToProto[pb.NodeType_Capability](mapCtx, in.Capabilities)
	return out
}
func VmwareengineNodeTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NodeType) *krm.VmwareengineNodeTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineNodeTypeObservedState{}
	// MISSING: Name
	// MISSING: NodeTypeID
	// MISSING: DisplayName
	// MISSING: VirtualCpuCount
	// MISSING: TotalCoreCount
	// MISSING: MemoryGB
	// MISSING: DiskSizeGB
	// MISSING: AvailableCustomCoreCounts
	// MISSING: Kind
	// MISSING: Families
	// MISSING: Capabilities
	return out
}
func VmwareengineNodeTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineNodeTypeObservedState) *pb.NodeType {
	if in == nil {
		return nil
	}
	out := &pb.NodeType{}
	// MISSING: Name
	// MISSING: NodeTypeID
	// MISSING: DisplayName
	// MISSING: VirtualCpuCount
	// MISSING: TotalCoreCount
	// MISSING: MemoryGB
	// MISSING: DiskSizeGB
	// MISSING: AvailableCustomCoreCounts
	// MISSING: Kind
	// MISSING: Families
	// MISSING: Capabilities
	return out
}
func VmwareengineNodeTypeSpec_FromProto(mapCtx *direct.MapContext, in *pb.NodeType) *krm.VmwareengineNodeTypeSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineNodeTypeSpec{}
	// MISSING: Name
	// MISSING: NodeTypeID
	// MISSING: DisplayName
	// MISSING: VirtualCpuCount
	// MISSING: TotalCoreCount
	// MISSING: MemoryGB
	// MISSING: DiskSizeGB
	// MISSING: AvailableCustomCoreCounts
	// MISSING: Kind
	// MISSING: Families
	// MISSING: Capabilities
	return out
}
func VmwareengineNodeTypeSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineNodeTypeSpec) *pb.NodeType {
	if in == nil {
		return nil
	}
	out := &pb.NodeType{}
	// MISSING: Name
	// MISSING: NodeTypeID
	// MISSING: DisplayName
	// MISSING: VirtualCpuCount
	// MISSING: TotalCoreCount
	// MISSING: MemoryGB
	// MISSING: DiskSizeGB
	// MISSING: AvailableCustomCoreCounts
	// MISSING: Kind
	// MISSING: Families
	// MISSING: Capabilities
	return out
}
