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

package edgenetwork

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/edgenetwork/apiv1/edgenetworkpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgenetwork/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func EdgenetworkSubnetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Subnet) *krm.EdgenetworkSubnetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgenetworkSubnetObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Ipv4Cidr
	// MISSING: Ipv6Cidr
	// MISSING: VlanID
	// MISSING: BondingType
	// MISSING: State
	return out
}
func EdgenetworkSubnetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgenetworkSubnetObservedState) *pb.Subnet {
	if in == nil {
		return nil
	}
	out := &pb.Subnet{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Ipv4Cidr
	// MISSING: Ipv6Cidr
	// MISSING: VlanID
	// MISSING: BondingType
	// MISSING: State
	return out
}
func EdgenetworkSubnetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Subnet) *krm.EdgenetworkSubnetSpec {
	if in == nil {
		return nil
	}
	out := &krm.EdgenetworkSubnetSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Ipv4Cidr
	// MISSING: Ipv6Cidr
	// MISSING: VlanID
	// MISSING: BondingType
	// MISSING: State
	return out
}
func EdgenetworkSubnetSpec_ToProto(mapCtx *direct.MapContext, in *krm.EdgenetworkSubnetSpec) *pb.Subnet {
	if in == nil {
		return nil
	}
	out := &pb.Subnet{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Ipv4Cidr
	// MISSING: Ipv6Cidr
	// MISSING: VlanID
	// MISSING: BondingType
	// MISSING: State
	return out
}
func Subnet_FromProto(mapCtx *direct.MapContext, in *pb.Subnet) *krm.Subnet {
	if in == nil {
		return nil
	}
	out := &krm.Subnet{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Ipv4Cidr = in.Ipv4Cidr
	out.Ipv6Cidr = in.Ipv6Cidr
	out.VlanID = direct.LazyPtr(in.GetVlanId())
	out.BondingType = direct.Enum_FromProto(mapCtx, in.GetBondingType())
	// MISSING: State
	return out
}
func Subnet_ToProto(mapCtx *direct.MapContext, in *krm.Subnet) *pb.Subnet {
	if in == nil {
		return nil
	}
	out := &pb.Subnet{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Network = direct.ValueOf(in.Network)
	out.Ipv4Cidr = in.Ipv4Cidr
	out.Ipv6Cidr = in.Ipv6Cidr
	out.VlanId = direct.ValueOf(in.VlanID)
	out.BondingType = direct.Enum_ToProto[pb.Subnet_BondingType](mapCtx, in.BondingType)
	// MISSING: State
	return out
}
func SubnetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Subnet) *krm.SubnetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SubnetObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Ipv4Cidr
	// MISSING: Ipv6Cidr
	// MISSING: VlanID
	// MISSING: BondingType
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func SubnetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SubnetObservedState) *pb.Subnet {
	if in == nil {
		return nil
	}
	out := &pb.Subnet{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Ipv4Cidr
	// MISSING: Ipv6Cidr
	// MISSING: VlanID
	// MISSING: BondingType
	out.State = direct.Enum_ToProto[pb.ResourceState](mapCtx, in.State)
	return out
}
