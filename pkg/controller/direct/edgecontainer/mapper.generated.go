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

package edgecontainer

import (
	pb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgecontainer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func EdgecontainerMachineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Machine) *krm.EdgecontainerMachineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgecontainerMachineObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: HostedNode
	// MISSING: Zone
	// MISSING: Version
	// MISSING: Disabled
	return out
}
func EdgecontainerMachineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgecontainerMachineObservedState) *pb.Machine {
	if in == nil {
		return nil
	}
	out := &pb.Machine{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: HostedNode
	// MISSING: Zone
	// MISSING: Version
	// MISSING: Disabled
	return out
}
func EdgecontainerMachineSpec_FromProto(mapCtx *direct.MapContext, in *pb.Machine) *krm.EdgecontainerMachineSpec {
	if in == nil {
		return nil
	}
	out := &krm.EdgecontainerMachineSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: HostedNode
	// MISSING: Zone
	// MISSING: Version
	// MISSING: Disabled
	return out
}
func EdgecontainerMachineSpec_ToProto(mapCtx *direct.MapContext, in *krm.EdgecontainerMachineSpec) *pb.Machine {
	if in == nil {
		return nil
	}
	out := &pb.Machine{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: HostedNode
	// MISSING: Zone
	// MISSING: Version
	// MISSING: Disabled
	return out
}
func Machine_FromProto(mapCtx *direct.MapContext, in *pb.Machine) *krm.Machine {
	if in == nil {
		return nil
	}
	out := &krm.Machine{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.HostedNode = direct.LazyPtr(in.GetHostedNode())
	out.Zone = direct.LazyPtr(in.GetZone())
	// MISSING: Version
	// MISSING: Disabled
	return out
}
func Machine_ToProto(mapCtx *direct.MapContext, in *krm.Machine) *pb.Machine {
	if in == nil {
		return nil
	}
	out := &pb.Machine{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.HostedNode = direct.ValueOf(in.HostedNode)
	out.Zone = direct.ValueOf(in.Zone)
	// MISSING: Version
	// MISSING: Disabled
	return out
}
func MachineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Machine) *krm.MachineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MachineObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: HostedNode
	// MISSING: Zone
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}
func MachineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MachineObservedState) *pb.Machine {
	if in == nil {
		return nil
	}
	out := &pb.Machine{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: HostedNode
	// MISSING: Zone
	out.Version = direct.ValueOf(in.Version)
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}
