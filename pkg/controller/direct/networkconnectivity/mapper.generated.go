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

package networkconnectivity

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func NetworkconnectivityRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krm.NetworkconnectivityRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkconnectivityRouteObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IPCidrRange
	// MISSING: Type
	// MISSING: NextHopVpcNetwork
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Uid
	// MISSING: State
	// MISSING: Spoke
	// MISSING: Location
	// MISSING: Priority
	// MISSING: NextHopVpnTunnel
	// MISSING: NextHopRouterApplianceInstance
	// MISSING: NextHopInterconnectAttachment
	return out
}
func NetworkconnectivityRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkconnectivityRouteObservedState) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IPCidrRange
	// MISSING: Type
	// MISSING: NextHopVpcNetwork
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Uid
	// MISSING: State
	// MISSING: Spoke
	// MISSING: Location
	// MISSING: Priority
	// MISSING: NextHopVpnTunnel
	// MISSING: NextHopRouterApplianceInstance
	// MISSING: NextHopInterconnectAttachment
	return out
}
func NetworkconnectivityRouteSpec_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krm.NetworkconnectivityRouteSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkconnectivityRouteSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IPCidrRange
	// MISSING: Type
	// MISSING: NextHopVpcNetwork
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Uid
	// MISSING: State
	// MISSING: Spoke
	// MISSING: Location
	// MISSING: Priority
	// MISSING: NextHopVpnTunnel
	// MISSING: NextHopRouterApplianceInstance
	// MISSING: NextHopInterconnectAttachment
	return out
}
func NetworkconnectivityRouteSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkconnectivityRouteSpec) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: IPCidrRange
	// MISSING: Type
	// MISSING: NextHopVpcNetwork
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Uid
	// MISSING: State
	// MISSING: Spoke
	// MISSING: Location
	// MISSING: Priority
	// MISSING: NextHopVpnTunnel
	// MISSING: NextHopRouterApplianceInstance
	// MISSING: NextHopInterconnectAttachment
	return out
}
func NextHopInterconnectAttachment_FromProto(mapCtx *direct.MapContext, in *pb.NextHopInterconnectAttachment) *krm.NextHopInterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &krm.NextHopInterconnectAttachment{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	return out
}
func NextHopInterconnectAttachment_ToProto(mapCtx *direct.MapContext, in *krm.NextHopInterconnectAttachment) *pb.NextHopInterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.NextHopInterconnectAttachment{}
	out.Uri = direct.ValueOf(in.URI)
	out.VpcNetwork = direct.ValueOf(in.VpcNetwork)
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	return out
}
func NextHopRouterApplianceInstance_FromProto(mapCtx *direct.MapContext, in *pb.NextHopRouterApplianceInstance) *krm.NextHopRouterApplianceInstance {
	if in == nil {
		return nil
	}
	out := &krm.NextHopRouterApplianceInstance{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	return out
}
func NextHopRouterApplianceInstance_ToProto(mapCtx *direct.MapContext, in *krm.NextHopRouterApplianceInstance) *pb.NextHopRouterApplianceInstance {
	if in == nil {
		return nil
	}
	out := &pb.NextHopRouterApplianceInstance{}
	out.Uri = direct.ValueOf(in.URI)
	out.VpcNetwork = direct.ValueOf(in.VpcNetwork)
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	return out
}
func NextHopVPNTunnel_FromProto(mapCtx *direct.MapContext, in *pb.NextHopVPNTunnel) *krm.NextHopVPNTunnel {
	if in == nil {
		return nil
	}
	out := &krm.NextHopVPNTunnel{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	return out
}
func NextHopVPNTunnel_ToProto(mapCtx *direct.MapContext, in *krm.NextHopVPNTunnel) *pb.NextHopVPNTunnel {
	if in == nil {
		return nil
	}
	out := &pb.NextHopVPNTunnel{}
	out.Uri = direct.ValueOf(in.URI)
	out.VpcNetwork = direct.ValueOf(in.VpcNetwork)
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	return out
}
func NextHopVpcNetwork_FromProto(mapCtx *direct.MapContext, in *pb.NextHopVpcNetwork) *krm.NextHopVpcNetwork {
	if in == nil {
		return nil
	}
	out := &krm.NextHopVpcNetwork{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func NextHopVpcNetwork_ToProto(mapCtx *direct.MapContext, in *krm.NextHopVpcNetwork) *pb.NextHopVpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.NextHopVpcNetwork{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Route_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krm.Route {
	if in == nil {
		return nil
	}
	out := &krm.Route{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.IPCidrRange = direct.LazyPtr(in.GetIpCidrRange())
	// MISSING: Type
	out.NextHopVpcNetwork = NextHopVpcNetwork_FromProto(mapCtx, in.GetNextHopVpcNetwork())
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: Uid
	// MISSING: State
	out.Spoke = direct.LazyPtr(in.GetSpoke())
	// MISSING: Location
	// MISSING: Priority
	out.NextHopVpnTunnel = NextHopVPNTunnel_FromProto(mapCtx, in.GetNextHopVpnTunnel())
	out.NextHopRouterApplianceInstance = NextHopRouterApplianceInstance_FromProto(mapCtx, in.GetNextHopRouterApplianceInstance())
	out.NextHopInterconnectAttachment = NextHopInterconnectAttachment_FromProto(mapCtx, in.GetNextHopInterconnectAttachment())
	return out
}
func Route_ToProto(mapCtx *direct.MapContext, in *krm.Route) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.IpCidrRange = direct.ValueOf(in.IPCidrRange)
	// MISSING: Type
	out.NextHopVpcNetwork = NextHopVpcNetwork_ToProto(mapCtx, in.NextHopVpcNetwork)
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	// MISSING: Uid
	// MISSING: State
	out.Spoke = direct.ValueOf(in.Spoke)
	// MISSING: Location
	// MISSING: Priority
	out.NextHopVpnTunnel = NextHopVPNTunnel_ToProto(mapCtx, in.NextHopVpnTunnel)
	out.NextHopRouterApplianceInstance = NextHopRouterApplianceInstance_ToProto(mapCtx, in.NextHopRouterApplianceInstance)
	out.NextHopInterconnectAttachment = NextHopInterconnectAttachment_ToProto(mapCtx, in.NextHopInterconnectAttachment)
	return out
}
func RouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krm.RouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RouteObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: IPCidrRange
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: NextHopVpcNetwork
	// MISSING: Labels
	// MISSING: Description
	out.Uid = direct.LazyPtr(in.GetUid())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Spoke
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Priority = direct.LazyPtr(in.GetPriority())
	// MISSING: NextHopVpnTunnel
	// MISSING: NextHopRouterApplianceInstance
	// MISSING: NextHopInterconnectAttachment
	return out
}
func RouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RouteObservedState) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: IPCidrRange
	out.Type = direct.Enum_ToProto[pb.RouteType](mapCtx, in.Type)
	// MISSING: NextHopVpcNetwork
	// MISSING: Labels
	// MISSING: Description
	out.Uid = direct.ValueOf(in.Uid)
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	// MISSING: Spoke
	out.Location = direct.ValueOf(in.Location)
	out.Priority = direct.ValueOf(in.Priority)
	// MISSING: NextHopVpnTunnel
	// MISSING: NextHopRouterApplianceInstance
	// MISSING: NextHopInterconnectAttachment
	return out
}
