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

// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeRouteSpec_v1beta1_FromProto maps the proto Route message to the KRM ComputeRouteSpec.
// This function is handcoded because of type mismatches between proto fields and CRD fields:
// - destRange: proto uses *string, but KRM uses string (required)
// - priority: proto uses *uint32, but KRM uses *int32
func ComputeRouteSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Route) *krm.ComputeRouteSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRouteSpec{}
	out.Description = in.Description
	if in.DestRange != nil {
		out.DestRange = *in.DestRange
	}
	if in.GetNetwork() != "" {
		out.NetworkRef = &krm.ComputeRouteNetworkRef{External: in.GetNetwork()}
	}
	out.NextHopGateway = in.NextHopGateway
	if in.GetNextHopInstance() != "" {
		out.NextHopInstanceRef = &krm.ComputeRouteNextHopInstanceRef{External: in.GetNextHopInstance()}
	}
	out.NextHopIp = in.NextHopIp
	if in.GetNextHopVpnTunnel() != "" {
		out.NextHopVPNTunnelRef = &krm.ComputeRouteNextHopVPNTunnelRef{External: in.GetNextHopVpnTunnel()}
	}
	if in.Priority != nil {
		p := int32(*in.Priority)
		out.Priority = &p
	}
	out.Tags = in.Tags
	return out
}

// ComputeRouteSpec_v1beta1_ToProto maps the KRM ComputeRouteSpec to the proto Route message.
// This function is handcoded because of type mismatches between proto fields and CRD fields:
// - destRange: proto uses *string, but KRM uses string (required)
// - priority: proto uses *uint32, but KRM uses *int32
func ComputeRouteSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRouteSpec) *pb.Route {
	if in == nil {
		return nil
	}
	out := &pb.Route{}
	out.Description = in.Description
	if in.DestRange != "" {
		out.DestRange = &in.DestRange
	}
	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}
	out.NextHopGateway = in.NextHopGateway
	if in.NextHopInstanceRef != nil {
		out.NextHopInstance = &in.NextHopInstanceRef.External
	}
	out.NextHopIp = in.NextHopIp
	if in.NextHopVPNTunnelRef != nil {
		out.NextHopVpnTunnel = &in.NextHopVPNTunnelRef.External
	}
	if in.Priority != nil {
		p := uint32(*in.Priority)
		out.Priority = &p
	}
	out.Tags = in.Tags
	return out
}
