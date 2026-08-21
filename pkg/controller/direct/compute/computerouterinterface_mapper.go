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

// ComputeRouterInterfaceSpec_v1beta1_FromProto maps a pb.RouterInterface to a krm.ComputeRouterInterfaceSpec.
// We hand-code this function because custom reference types (like InterconnectAttachmentRef, PrivateIPAddressRef, etc.)
// require non-standard mapping logic that does not match 1:1 with field names.
func ComputeRouterInterfaceSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterInterface) *krm.ComputeRouterInterfaceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRouterInterfaceSpec{}
	out.IPRange = in.IpRange
	if in.GetLinkedInterconnectAttachment() != "" {
		out.InterconnectAttachmentRef = &krm.ComputeInterconnectAttachmentRef{External: in.GetLinkedInterconnectAttachment()}
	}
	if in.GetLinkedVpnTunnel() != "" {
		out.VPNTunnelRef = &krm.ComputeVPNTunnelRef{External: in.GetLinkedVpnTunnel()}
	}
	if in.GetPrivateIpAddress() != "" {
		out.PrivateIPAddressRef = &krm.ComputeAddressRef{External: in.GetPrivateIpAddress()}
	}
	if in.GetRedundantInterface() != "" {
		out.RedundantInterfaceRef = &krm.ComputeRouterInterfaceRef{External: in.GetRedundantInterface()}
	}
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &krm.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	out.ResourceID = in.Name
	return out
}

// ComputeRouterInterfaceSpec_v1beta1_ToProto maps a krm.ComputeRouterInterfaceSpec to a pb.RouterInterface.
// We hand-code this function because custom reference types (like InterconnectAttachmentRef, PrivateIPAddressRef, etc.)
// require non-standard mapping logic that does not match 1:1 with field names.
func ComputeRouterInterfaceSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRouterInterfaceSpec) *pb.RouterInterface {
	if in == nil {
		return nil
	}
	out := &pb.RouterInterface{}
	out.IpRange = in.IPRange
	if in.InterconnectAttachmentRef != nil {
		out.LinkedInterconnectAttachment = direct.LazyPtr(in.InterconnectAttachmentRef.External)
	}
	if in.VPNTunnelRef != nil {
		out.LinkedVpnTunnel = direct.LazyPtr(in.VPNTunnelRef.External)
	}
	if in.PrivateIPAddressRef != nil {
		out.PrivateIpAddress = direct.LazyPtr(in.PrivateIPAddressRef.External)
	}
	if in.RedundantInterfaceRef != nil {
		out.RedundantInterface = direct.LazyPtr(in.RedundantInterfaceRef.External)
	}
	if in.SubnetworkRef != nil {
		out.Subnetwork = direct.LazyPtr(in.SubnetworkRef.External)
	}
	out.Name = in.ResourceID
	return out
}
