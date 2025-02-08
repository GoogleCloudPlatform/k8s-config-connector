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
func Subnet_FromProto(mapCtx *direct.MapContext, in *pb.Subnet) *krm.Subnet {
	if in == nil {
		return nil
	}
	out := &krm.Subnet{}
	// MISSING: Name
	out.IPCidrRange = direct.LazyPtr(in.GetIpCidrRange())
	out.GatewayIP = direct.LazyPtr(in.GetGatewayIp())
	// MISSING: Type
	// MISSING: State
	// MISSING: VlanID
	return out
}
func Subnet_ToProto(mapCtx *direct.MapContext, in *krm.Subnet) *pb.Subnet {
	if in == nil {
		return nil
	}
	out := &pb.Subnet{}
	// MISSING: Name
	out.IpCidrRange = direct.ValueOf(in.IPCidrRange)
	out.GatewayIp = direct.ValueOf(in.GatewayIP)
	// MISSING: Type
	// MISSING: State
	// MISSING: VlanID
	return out
}
func SubnetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Subnet) *krm.SubnetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SubnetObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: IPCidrRange
	// MISSING: GatewayIP
	out.Type = direct.LazyPtr(in.GetType())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.VlanID = direct.LazyPtr(in.GetVlanId())
	return out
}
func SubnetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SubnetObservedState) *pb.Subnet {
	if in == nil {
		return nil
	}
	out := &pb.Subnet{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: IPCidrRange
	// MISSING: GatewayIP
	out.Type = direct.ValueOf(in.Type)
	out.State = direct.Enum_ToProto[pb.Subnet_State](mapCtx, in.State)
	out.VlanId = direct.ValueOf(in.VlanID)
	return out
}
func VmwareengineSubnetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Subnet) *krm.VmwareengineSubnetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineSubnetObservedState{}
	// MISSING: Name
	// MISSING: IPCidrRange
	// MISSING: GatewayIP
	// MISSING: Type
	// MISSING: State
	// MISSING: VlanID
	return out
}
func VmwareengineSubnetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineSubnetObservedState) *pb.Subnet {
	if in == nil {
		return nil
	}
	out := &pb.Subnet{}
	// MISSING: Name
	// MISSING: IPCidrRange
	// MISSING: GatewayIP
	// MISSING: Type
	// MISSING: State
	// MISSING: VlanID
	return out
}
func VmwareengineSubnetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Subnet) *krm.VmwareengineSubnetSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineSubnetSpec{}
	// MISSING: Name
	// MISSING: IPCidrRange
	// MISSING: GatewayIP
	// MISSING: Type
	// MISSING: State
	// MISSING: VlanID
	return out
}
func VmwareengineSubnetSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineSubnetSpec) *pb.Subnet {
	if in == nil {
		return nil
	}
	out := &pb.Subnet{}
	// MISSING: Name
	// MISSING: IPCidrRange
	// MISSING: GatewayIP
	// MISSING: Type
	// MISSING: State
	// MISSING: VlanID
	return out
}
