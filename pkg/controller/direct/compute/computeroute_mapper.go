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

func ComputeRouteSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krm.ComputeRouteSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRouteSpec{}
	out.Description = in.Description
	out.DestRange = in.GetDestRange()
	if in.GetNetwork() != "" {
		out.NetworkRef = krm.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.NextHopGateway = in.NextHopGateway
	if in.GetNextHopIlb() != "" {
		out.NextHopILBRef = &krm.ForwardingRuleRef{External: in.GetNextHopIlb()}
	}
	if in.GetNextHopInstance() != "" {
		out.NextHopInstanceRef = &krm.InstanceRef{External: in.GetNextHopInstance()}
	}
	out.NextHopIP = in.NextHopIp
	if in.GetNextHopVpnTunnel() != "" {
		out.NextHopVPNTunnelRef = &krm.ComputeVPNTunnelRef{External: in.GetNextHopVpnTunnel()}
	}
	out.Priority = in.Priority
	out.Tags = in.Tags
	return out
}

func ComputeRouteSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRouteSpec) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	out.Description = in.Description
	if in.DestRange != "" {
		out.DestRange = &in.DestRange
	}
	if in.NetworkRef.External != "" {
		out.Network = &in.NetworkRef.External
	}
	out.NextHopGateway = in.NextHopGateway
	if in.NextHopILBRef != nil {
		out.NextHopIlb = &in.NextHopILBRef.External
	}
	if in.NextHopInstanceRef != nil {
		out.NextHopInstance = &in.NextHopInstanceRef.External
	}
	out.NextHopIp = in.NextHopIP
	if in.NextHopVPNTunnelRef != nil {
		out.NextHopVpnTunnel = &in.NextHopVPNTunnelRef.External
	}
	out.Priority = in.Priority
	out.Tags = in.Tags
	return out
}
