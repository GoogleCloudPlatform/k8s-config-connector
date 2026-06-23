// Copyright 2026 Google LLC
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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeRouterPeerSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterBgpPeer) *krm.ComputeRouterPeerSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRouterPeerSpec{}
	out.AdvertiseMode = in.AdvertiseMode
	out.AdvertisedGroups = in.AdvertisedGroups
	if in.AdvertisedIpRanges != nil {
		out.AdvertisedIpRanges = make([]krm.RouterpeerAdvertisedIpRanges, len(in.AdvertisedIpRanges))
		for i, v := range in.AdvertisedIpRanges {
			out.AdvertisedIpRanges[i] = RouterpeerAdvertisedIpRanges_FromProto(v)
		}
	}
	out.AdvertisedRoutePriority = uint32ToInt64Ptr(in.AdvertisedRoutePriority)
	out.Bfd = RouterpeerBfd_FromProto(in.Bfd)
	out.Enable = stringToBoolPtr(in.Enable)
	out.EnableIpv6 = in.EnableIpv6
	out.IpAddress = RouterpeerIpAddress_FromProto(in.IpAddress)
	out.Ipv6NexthopAddress = in.Ipv6NexthopAddress
	out.PeerAsn = peerAsnFromProto(in.PeerAsn)
	out.PeerIpAddress = in.PeerIpAddress
	out.PeerIpv6NexthopAddress = in.PeerIpv6NexthopAddress
	out.RouterApplianceInstanceRef = routerApplianceInstanceFromProto(in.RouterApplianceInstance)
	out.RouterInterfaceRef = routerInterfaceRefFromProto(in.InterfaceName)

	return out
}

func ComputeRouterPeerSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRouterPeerSpec) *pb.RouterBgpPeer {
	if in == nil {
		return nil
	}
	out := &pb.RouterBgpPeer{}
	out.AdvertiseMode = in.AdvertiseMode
	out.AdvertisedGroups = in.AdvertisedGroups
	if in.AdvertisedIpRanges != nil {
		out.AdvertisedIpRanges = make([]*pb.RouterAdvertisedIpRange, len(in.AdvertisedIpRanges))
		for i, v := range in.AdvertisedIpRanges {
			out.AdvertisedIpRanges[i] = RouterpeerAdvertisedIpRanges_ToProto(v)
		}
	}
	out.AdvertisedRoutePriority = int64ToUint32Ptr(in.AdvertisedRoutePriority)
	out.Bfd = RouterpeerBfd_ToProto(in.Bfd)
	out.Enable = boolToStringPtr(in.Enable)
	out.EnableIpv6 = in.EnableIpv6
	out.IpAddress = RouterpeerIpAddress_ToProto(in.IpAddress)
	out.Ipv6NexthopAddress = in.Ipv6NexthopAddress
	out.PeerAsn = peerAsnToProto(in.PeerAsn)
	out.PeerIpAddress = in.PeerIpAddress
	out.PeerIpv6NexthopAddress = in.PeerIpv6NexthopAddress
	out.RouterApplianceInstance = routerApplianceInstanceToProto(in.RouterApplianceInstanceRef)
	out.InterfaceName = routerInterfaceRefToProto(in.RouterInterfaceRef)

	return out
}

func ComputeRouterPeerStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterBgpPeer) *krm.ComputeRouterPeerStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRouterPeerStatus{}
	out.ManagementType = in.ManagementType
	return out
}

func ComputeRouterPeerStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRouterPeerStatus) *pb.RouterBgpPeer {
	if in == nil {
		return nil
	}
	out := &pb.RouterBgpPeer{}
	out.ManagementType = in.ManagementType
	return out
}

func RouterpeerAdvertisedIpRanges_FromProto(in *pb.RouterAdvertisedIpRange) krm.RouterpeerAdvertisedIpRanges {
	if in == nil {
		return krm.RouterpeerAdvertisedIpRanges{}
	}
	out := krm.RouterpeerAdvertisedIpRanges{}
	out.Description = in.Description
	if in.Range != nil {
		out.Range = *in.Range
	}
	return out
}

func RouterpeerAdvertisedIpRanges_ToProto(in krm.RouterpeerAdvertisedIpRanges) *pb.RouterAdvertisedIpRange {
	out := &pb.RouterAdvertisedIpRange{}
	out.Description = in.Description
	if in.Range != "" {
		out.Range = &in.Range
	}
	return out
}

func RouterpeerBfd_FromProto(in *pb.RouterBgpPeerBfd) *krm.RouterpeerBfd {
	if in == nil {
		return nil
	}
	out := &krm.RouterpeerBfd{}
	out.MinReceiveInterval = uint32ToInt64Ptr(in.MinReceiveInterval)
	out.MinTransmitInterval = uint32ToInt64Ptr(in.MinTransmitInterval)
	out.Multiplier = uint32ToInt64Ptr(in.Multiplier)
	if in.SessionInitializationMode != nil {
		out.SessionInitializationMode = *in.SessionInitializationMode
	}
	return out
}

func RouterpeerBfd_ToProto(in *krm.RouterpeerBfd) *pb.RouterBgpPeerBfd {
	if in == nil {
		return nil
	}
	out := &pb.RouterBgpPeerBfd{}
	out.MinReceiveInterval = int64ToUint32Ptr(in.MinReceiveInterval)
	out.MinTransmitInterval = int64ToUint32Ptr(in.MinTransmitInterval)
	out.Multiplier = int64ToUint32Ptr(in.Multiplier)
	if in.SessionInitializationMode != "" {
		out.SessionInitializationMode = &in.SessionInitializationMode
	}
	return out
}

func RouterpeerIpAddress_FromProto(in *string) *krm.RouterpeerIpAddress {
	if in == nil || *in == "" {
		return nil
	}
	return &krm.RouterpeerIpAddress{
		External: in,
	}
}

func RouterpeerIpAddress_ToProto(in *krm.RouterpeerIpAddress) *string {
	if in == nil {
		return nil
	}
	return in.External
}

func routerApplianceInstanceFromProto(in *string) *krm.ComputeRouterPeerRef {
	if in == nil || *in == "" {
		return nil
	}
	return &krm.ComputeRouterPeerRef{
		External: *in,
	}
}

func routerApplianceInstanceToProto(in *krm.ComputeRouterPeerRef) *string {
	if in == nil || in.External == "" {
		return nil
	}
	return &in.External
}

func routerInterfaceRefFromProto(in *string) krm.ComputeRouterPeerRef {
	if in == nil {
		return krm.ComputeRouterPeerRef{}
	}
	return krm.ComputeRouterPeerRef{
		External: *in,
	}
}

func routerInterfaceRefToProto(in krm.ComputeRouterPeerRef) *string {
	if in.External == "" {
		return nil
	}
	return &in.External
}

func stringToBoolPtr(in *string) *bool {
	if in == nil {
		return nil
	}
	val := (*in == "TRUE")
	return &val
}

func boolToStringPtr(in *bool) *string {
	if in == nil {
		return nil
	}
	var val string
	if *in {
		val = "TRUE"
	} else {
		val = "FALSE"
	}
	return &val
}

func int64ToUint32Ptr(in *int64) *uint32 {
	if in == nil {
		return nil
	}
	val := uint32(*in)
	return &val
}

func uint32ToInt64Ptr(in *uint32) *int64 {
	if in == nil {
		return nil
	}
	val := int64(*in)
	return &val
}

func peerAsnFromProto(in *uint32) int64 {
	if in == nil {
		return 0
	}
	return int64(*in)
}

func peerAsnToProto(in int64) *uint32 {
	val := uint32(in)
	return &val
}
