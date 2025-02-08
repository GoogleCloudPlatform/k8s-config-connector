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
func EdgenetworkRouterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Router) *krm.EdgenetworkRouterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgenetworkRouterObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Interface
	// MISSING: BgpPeer
	// MISSING: Bgp
	// MISSING: State
	// MISSING: RouteAdvertisements
	return out
}
func EdgenetworkRouterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgenetworkRouterObservedState) *pb.Router {
	if in == nil {
		return nil
	}
	out := &pb.Router{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Interface
	// MISSING: BgpPeer
	// MISSING: Bgp
	// MISSING: State
	// MISSING: RouteAdvertisements
	return out
}
func EdgenetworkRouterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Router) *krm.EdgenetworkRouterSpec {
	if in == nil {
		return nil
	}
	out := &krm.EdgenetworkRouterSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Interface
	// MISSING: BgpPeer
	// MISSING: Bgp
	// MISSING: State
	// MISSING: RouteAdvertisements
	return out
}
func EdgenetworkRouterSpec_ToProto(mapCtx *direct.MapContext, in *krm.EdgenetworkRouterSpec) *pb.Router {
	if in == nil {
		return nil
	}
	out := &pb.Router{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Interface
	// MISSING: BgpPeer
	// MISSING: Bgp
	// MISSING: State
	// MISSING: RouteAdvertisements
	return out
}
func Router_FromProto(mapCtx *direct.MapContext, in *pb.Router) *krm.Router {
	if in == nil {
		return nil
	}
	out := &krm.Router{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Interface = direct.Slice_FromProto(mapCtx, in.Interface, Router_Interface_FromProto)
	out.BgpPeer = direct.Slice_FromProto(mapCtx, in.BgpPeer, Router_BgpPeer_FromProto)
	out.Bgp = Router_Bgp_FromProto(mapCtx, in.GetBgp())
	// MISSING: State
	out.RouteAdvertisements = in.RouteAdvertisements
	return out
}
func Router_ToProto(mapCtx *direct.MapContext, in *krm.Router) *pb.Router {
	if in == nil {
		return nil
	}
	out := &pb.Router{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Network = direct.ValueOf(in.Network)
	out.Interface = direct.Slice_ToProto(mapCtx, in.Interface, Router_Interface_ToProto)
	out.BgpPeer = direct.Slice_ToProto(mapCtx, in.BgpPeer, Router_BgpPeer_ToProto)
	out.Bgp = Router_Bgp_ToProto(mapCtx, in.Bgp)
	// MISSING: State
	out.RouteAdvertisements = in.RouteAdvertisements
	return out
}
func RouterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Router) *krm.RouterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RouterObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Interface
	out.BgpPeer = direct.Slice_FromProto(mapCtx, in.BgpPeer, Router_BgpPeerObservedState_FromProto)
	// MISSING: Bgp
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: RouteAdvertisements
	return out
}
func RouterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RouterObservedState) *pb.Router {
	if in == nil {
		return nil
	}
	out := &pb.Router{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Interface
	out.BgpPeer = direct.Slice_ToProto(mapCtx, in.BgpPeer, Router_BgpPeerObservedState_ToProto)
	// MISSING: Bgp
	out.State = direct.Enum_ToProto[pb.ResourceState](mapCtx, in.State)
	// MISSING: RouteAdvertisements
	return out
}
func Router_Bgp_FromProto(mapCtx *direct.MapContext, in *pb.Router_Bgp) *krm.Router_Bgp {
	if in == nil {
		return nil
	}
	out := &krm.Router_Bgp{}
	out.Asn = direct.LazyPtr(in.GetAsn())
	out.KeepaliveIntervalInSeconds = direct.LazyPtr(in.GetKeepaliveIntervalInSeconds())
	return out
}
func Router_Bgp_ToProto(mapCtx *direct.MapContext, in *krm.Router_Bgp) *pb.Router_Bgp {
	if in == nil {
		return nil
	}
	out := &pb.Router_Bgp{}
	out.Asn = direct.ValueOf(in.Asn)
	out.KeepaliveIntervalInSeconds = direct.ValueOf(in.KeepaliveIntervalInSeconds)
	return out
}
func Router_BgpPeer_FromProto(mapCtx *direct.MapContext, in *pb.Router_BgpPeer) *krm.Router_BgpPeer {
	if in == nil {
		return nil
	}
	out := &krm.Router_BgpPeer{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Interface = direct.LazyPtr(in.GetInterface())
	out.InterfaceIpv4Cidr = direct.LazyPtr(in.GetInterfaceIpv4Cidr())
	out.InterfaceIpv6Cidr = direct.LazyPtr(in.GetInterfaceIpv6Cidr())
	out.PeerIpv4Cidr = direct.LazyPtr(in.GetPeerIpv4Cidr())
	out.PeerIpv6Cidr = direct.LazyPtr(in.GetPeerIpv6Cidr())
	out.PeerAsn = direct.LazyPtr(in.GetPeerAsn())
	// MISSING: LocalAsn
	return out
}
func Router_BgpPeer_ToProto(mapCtx *direct.MapContext, in *krm.Router_BgpPeer) *pb.Router_BgpPeer {
	if in == nil {
		return nil
	}
	out := &pb.Router_BgpPeer{}
	out.Name = direct.ValueOf(in.Name)
	out.Interface = direct.ValueOf(in.Interface)
	out.InterfaceIpv4Cidr = direct.ValueOf(in.InterfaceIpv4Cidr)
	out.InterfaceIpv6Cidr = direct.ValueOf(in.InterfaceIpv6Cidr)
	out.PeerIpv4Cidr = direct.ValueOf(in.PeerIpv4Cidr)
	out.PeerIpv6Cidr = direct.ValueOf(in.PeerIpv6Cidr)
	out.PeerAsn = direct.ValueOf(in.PeerAsn)
	// MISSING: LocalAsn
	return out
}
func Router_BgpPeerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Router_BgpPeer) *krm.Router_BgpPeerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Router_BgpPeerObservedState{}
	// MISSING: Name
	// MISSING: Interface
	// MISSING: InterfaceIpv4Cidr
	// MISSING: InterfaceIpv6Cidr
	// MISSING: PeerIpv4Cidr
	// MISSING: PeerIpv6Cidr
	// MISSING: PeerAsn
	out.LocalAsn = direct.LazyPtr(in.GetLocalAsn())
	return out
}
func Router_BgpPeerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Router_BgpPeerObservedState) *pb.Router_BgpPeer {
	if in == nil {
		return nil
	}
	out := &pb.Router_BgpPeer{}
	// MISSING: Name
	// MISSING: Interface
	// MISSING: InterfaceIpv4Cidr
	// MISSING: InterfaceIpv6Cidr
	// MISSING: PeerIpv4Cidr
	// MISSING: PeerIpv6Cidr
	// MISSING: PeerAsn
	out.LocalAsn = direct.ValueOf(in.LocalAsn)
	return out
}
func Router_Interface_FromProto(mapCtx *direct.MapContext, in *pb.Router_Interface) *krm.Router_Interface {
	if in == nil {
		return nil
	}
	out := &krm.Router_Interface{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Ipv4Cidr = direct.LazyPtr(in.GetIpv4Cidr())
	out.Ipv6Cidr = direct.LazyPtr(in.GetIpv6Cidr())
	out.LinkedInterconnectAttachment = direct.LazyPtr(in.GetLinkedInterconnectAttachment())
	out.Subnetwork = direct.LazyPtr(in.GetSubnetwork())
	out.LoopbackIPAddresses = in.LoopbackIpAddresses
	return out
}
func Router_Interface_ToProto(mapCtx *direct.MapContext, in *krm.Router_Interface) *pb.Router_Interface {
	if in == nil {
		return nil
	}
	out := &pb.Router_Interface{}
	out.Name = direct.ValueOf(in.Name)
	out.Ipv4Cidr = direct.ValueOf(in.Ipv4Cidr)
	out.Ipv6Cidr = direct.ValueOf(in.Ipv6Cidr)
	out.LinkedInterconnectAttachment = direct.ValueOf(in.LinkedInterconnectAttachment)
	out.Subnetwork = direct.ValueOf(in.Subnetwork)
	out.LoopbackIpAddresses = in.LoopbackIPAddresses
	return out
}
