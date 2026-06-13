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
	"strconv"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeVPNTunnelSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.VpnTunnel) *krm.ComputeVPNTunnelSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeVPNTunnelSpec{}
	out.Description = in.Description
	out.IkeVersion = direct.PtrInt32ToPtrInt64(in.IkeVersion)
	out.LocalTrafficSelector = in.LocalTrafficSelector
	out.PeerExternalGatewayInterface = direct.PtrInt32ToPtrInt64(in.PeerExternalGatewayInterface)
	if in.GetPeerExternalGateway() != "" {
		out.PeerExternalGatewayRef = &krm.ComputeExternalVPNGatewayRef{External: in.GetPeerExternalGateway()}
	}
	if in.GetPeerGcpGateway() != "" {
		out.PeerGCPGatewayRef = &krm.ComputeVPNGatewayRef{External: in.GetPeerGcpGateway()}
	}
	out.PeerIp = in.PeerIp
	out.Region = in.GetRegion()
	out.RemoteTrafficSelector = in.RemoteTrafficSelector
	if in.GetRouter() != "" {
		out.RouterRef = &krm.ComputeRouterRef{External: in.GetRouter()}
	}
	if in.SharedSecret != nil {
		out.SharedSecret.Value = in.SharedSecret
	}
	if in.GetTargetVpnGateway() != "" {
		out.TargetVPNGatewayRef = &refsv1beta1.ComputeTargetVPNGatewayRef{External: in.GetTargetVpnGateway()}
	}
	out.VpnGatewayInterface = direct.PtrInt32ToPtrInt64(in.VpnGatewayInterface)
	if in.GetVpnGateway() != "" {
		out.VpnGatewayRef = &krm.ComputeVPNGatewayRef{External: in.GetVpnGateway()}
	}

	return out
}

func ComputeVPNTunnelSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeVPNTunnelSpec) *pb.VpnTunnel {
	if in == nil {
		return nil
	}
	out := &pb.VpnTunnel{}
	out.Description = in.Description
	out.IkeVersion = direct.PtrInt64ToPtrInt32(in.IkeVersion)
	out.LocalTrafficSelector = in.LocalTrafficSelector
	out.PeerExternalGatewayInterface = direct.PtrInt64ToPtrInt32(in.PeerExternalGatewayInterface)
	if in.PeerExternalGatewayRef != nil && in.PeerExternalGatewayRef.External != "" {
		out.PeerExternalGateway = &in.PeerExternalGatewayRef.External
	}
	if in.PeerGCPGatewayRef != nil && in.PeerGCPGatewayRef.External != "" {
		out.PeerGcpGateway = &in.PeerGCPGatewayRef.External
	}
	out.PeerIp = in.PeerIp
	if in.Region != "" {
		out.Region = &in.Region
	}
	out.RemoteTrafficSelector = in.RemoteTrafficSelector
	if in.RouterRef != nil && in.RouterRef.External != "" {
		out.Router = &in.RouterRef.External
	}
	out.SharedSecret = in.SharedSecret.Value
	if in.TargetVPNGatewayRef != nil && in.TargetVPNGatewayRef.External != "" {
		out.TargetVpnGateway = &in.TargetVPNGatewayRef.External
	}
	out.VpnGatewayInterface = direct.PtrInt64ToPtrInt32(in.VpnGatewayInterface)
	if in.VpnGatewayRef != nil && in.VpnGatewayRef.External != "" {
		out.VpnGateway = &in.VpnGatewayRef.External
	}

	return out
}

func ComputeVPNTunnelStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.VpnTunnel) *krm.ComputeVPNTunnelStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeVPNTunnelStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.DetailedStatus = in.DetailedStatus
	out.LabelFingerprint = in.LabelFingerprint
	out.SelfLink = in.SelfLink
	out.SharedSecretHash = in.SharedSecretHash
	if in.Id != nil {
		idStr := strconv.FormatUint(*in.Id, 10)
		out.TunnelId = &idStr
	}

	return out
}

func ComputeVPNTunnelStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeVPNTunnelStatus) *pb.VpnTunnel {
	if in == nil {
		return nil
	}
	out := &pb.VpnTunnel{}
	out.CreationTimestamp = in.CreationTimestamp
	out.DetailedStatus = in.DetailedStatus
	out.LabelFingerprint = in.LabelFingerprint
	out.SelfLink = in.SelfLink
	out.SharedSecretHash = in.SharedSecretHash
	if in.TunnelId != nil {
		idVal, err := strconv.ParseUint(*in.TunnelId, 10, 64)
		if err == nil {
			out.Id = &idVal
		}
	}

	return out
}
