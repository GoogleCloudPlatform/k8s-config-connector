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
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeVPNGatewaySpec_v1beta1_FromProto maps a pb.VpnGateway to krm.ComputeVPNGatewaySpec.
// This function is handcoded because Region in KRM is a non-pointer string while it is a pointer string in Proto.
func ComputeVPNGatewaySpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.VpnGateway) *krm.ComputeVPNGatewaySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeVPNGatewaySpec{}
	out.Description = in.Description
	if in.GetNetwork() != "" {
		out.NetworkRef = &computerefs.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.Region = in.GetRegion()
	out.StackType = in.StackType
	out.VPNInterfaces = direct.Slice_FromProto(mapCtx, in.VpnInterfaces, VPNGatewayInterface_v1beta1_FromProto)
	return out
}

// ComputeVPNGatewaySpec_v1beta1_ToProto maps a krm.ComputeVPNGatewaySpec to pb.VpnGateway.
// This function is handcoded because Region in KRM is a non-pointer string while it is a pointer string in Proto.
func ComputeVPNGatewaySpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeVPNGatewaySpec) *pb.VpnGateway {
	if in == nil {
		return nil
	}
	out := &pb.VpnGateway{}
	out.Description = in.Description
	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}
	if in.Region != "" {
		out.Region = &in.Region
	}
	out.StackType = in.StackType
	out.VpnInterfaces = direct.Slice_ToProto(mapCtx, in.VPNInterfaces, VPNGatewayInterface_v1beta1_ToProto)
	return out
}

// VPNGatewayInterface_v1beta1_FromProto maps a pb.VpnGatewayVpnGatewayInterface to krm.VPNGatewayInterface.
// This function is handcoded because ID in KRM is an int64 pointer, whereas ID in Proto is a uint32 pointer.
func VPNGatewayInterface_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.VpnGatewayVpnGatewayInterface) *krm.VPNGatewayInterface {
	if in == nil {
		return nil
	}
	out := &krm.VPNGatewayInterface{}
	if in.Id != nil {
		idVal := int64(*in.Id)
		out.ID = &idVal
	}
	if in.GetInterconnectAttachment() != "" {
		out.InterconnectAttachmentRef = &krm.ComputeInterconnectAttachmentRef{
			External: in.GetInterconnectAttachment(),
		}
	}
	out.IPAddress = in.IpAddress
	return out
}

// VPNGatewayInterface_v1beta1_ToProto maps a krm.VPNGatewayInterface to pb.VpnGatewayVpnGatewayInterface.
// This function is handcoded because ID in KRM is an int64 pointer, whereas ID in Proto is a uint32 pointer.
func VPNGatewayInterface_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.VPNGatewayInterface) *pb.VpnGatewayVpnGatewayInterface {
	if in == nil {
		return nil
	}
	out := &pb.VpnGatewayVpnGatewayInterface{}
	if in.ID != nil {
		idVal := uint32(*in.ID)
		out.Id = &idVal
	}
	if in.InterconnectAttachmentRef != nil {
		out.InterconnectAttachment = &in.InterconnectAttachmentRef.External
	}
	out.IpAddress = in.IPAddress
	return out
}
