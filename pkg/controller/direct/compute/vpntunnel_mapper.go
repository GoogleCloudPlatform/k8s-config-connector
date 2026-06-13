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
	"fmt"

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
	out.IkeVersion = in.IkeVersion
	out.LocalTrafficSelector = in.LocalTrafficSelector
	out.Location = direct.ValueOf(in.Region) // Region in VpnTunnel is location
	if in.PeerExternalGateway != nil {
		out.PeerExternalGatewayRef = &refsv1beta1.ComputeExternalVPNGatewayRef{External: direct.ValueOf(in.PeerExternalGateway)}
	}
	out.PeerExternalGatewayInterface = in.PeerExternalGatewayInterface
	if in.PeerGcpGateway != nil {
		out.PeerGCPGatewayRef = &refsv1beta1.ComputeVPNGatewayRef{External: direct.ValueOf(in.PeerGcpGateway)}
	}
	out.PeerIp = in.PeerIp
	out.RemoteTrafficSelector = in.RemoteTrafficSelector
	if in.Router != nil {
		out.RouterRef = &refsv1beta1.ComputeRouterRef{External: direct.ValueOf(in.Router)}
	}
	if in.SharedSecret != nil {
		out.SharedSecret = krm.VPNTunnelSharedSecret{
			Value: in.SharedSecret,
		}
	}
	if in.TargetVpnGateway != nil {
		out.TargetVPNGatewayRef = &refsv1beta1.ComputeTargetVPNGatewayRef{External: direct.ValueOf(in.TargetVpnGateway)}
	}
	out.VpnGatewayInterface = in.VpnGatewayInterface
	if in.VpnGateway != nil {
		out.VpnGatewayRef = &refsv1beta1.ComputeVPNGatewayRef{External: direct.ValueOf(in.VpnGateway)}
	}
	return out
}

func ComputeVPNTunnelSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeVPNTunnelSpec) *pb.VpnTunnel {
	if in == nil {
		return nil
	}
	out := &pb.VpnTunnel{}
	out.Description = in.Description
	out.IkeVersion = in.IkeVersion
	out.LocalTrafficSelector = in.LocalTrafficSelector
	out.Region = direct.LazyPtr(in.Location) // Location in KRM is Region in GCP VpnTunnel
	if in.PeerExternalGatewayRef != nil {
		out.PeerExternalGateway = direct.LazyPtr(in.PeerExternalGatewayRef.External)
	}
	out.PeerExternalGatewayInterface = in.PeerExternalGatewayInterface
	if in.PeerGCPGatewayRef != nil {
		out.PeerGcpGateway = direct.LazyPtr(in.PeerGCPGatewayRef.External)
	}
	out.PeerIp = in.PeerIp
	out.RemoteTrafficSelector = in.RemoteTrafficSelector
	if in.RouterRef != nil {
		out.Router = direct.LazyPtr(in.RouterRef.External)
	}
	if in.SharedSecret.Value != nil {
		out.SharedSecret = in.SharedSecret.Value
	}
	if in.TargetVPNGatewayRef != nil {
		out.TargetVpnGateway = direct.LazyPtr(in.TargetVPNGatewayRef.External)
	}
	out.VpnGatewayInterface = in.VpnGatewayInterface
	if in.VpnGatewayRef != nil {
		out.VpnGateway = direct.LazyPtr(in.VpnGatewayRef.External)
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
		out.TunnelId = direct.LazyPtr(fmt.Sprintf("%d", direct.ValueOf(in.Id)))
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
		var id uint64
		if _, err := fmt.Sscanf(direct.ValueOf(in.TunnelId), "%d", &id); err == nil {
			out.Id = &id
		} else {
			mapCtx.Errorf("error parsing TunnelId: %v", err)
		}
	}
	return out
}
