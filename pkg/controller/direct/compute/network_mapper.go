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

func ComputeNetworkSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Network) *krm.ComputeNetworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkSpec{}
	// MISSING: IPv4Range
	out.AutoCreateSubnetworks = in.AutoCreateSubnetworks
	// MISSING: CreationTimestamp
	out.Description = in.Description
	out.EnableUlaInternalIPV6 = in.EnableUlaInternalIpv6
	// MISSING: FirewallPolicy
	// MISSING: GatewayIPv4
	// MISSING: ID
	out.InternalIPV6Range = in.InternalIpv6Range
	// MISSING: Kind
	out.Mtu = ConvertInt32ToInt(in.Mtu)
	// MISSING: Name
	out.NetworkFirewallPolicyEnforcementOrder = in.NetworkFirewallPolicyEnforcementOrder
	// MISSING: NetworkProfile
	// MISSING: Params
	// MISSING: Peerings
	if in.RoutingConfig != nil {
		out.RoutingMode = in.RoutingConfig.RoutingMode
	}
	// MISSING: SelfLinkWithID
	// MISSING: Subnetworks
	return out
}

func ComputeNetworkSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkSpec) *pb.Network {
	if in == nil {
		return nil
	}
	out := &pb.Network{}
	// MISSING: IPv4Range
	out.AutoCreateSubnetworks = in.AutoCreateSubnetworks
	// MISSING: CreationTimestamp
	out.Description = in.Description
	out.EnableUlaInternalIpv6 = in.EnableUlaInternalIPV6
	// MISSING: FirewallPolicy
	// MISSING: GatewayIPv4
	// MISSING: ID
	out.InternalIpv6Range = in.InternalIPV6Range
	// MISSING: Kind
	out.Mtu = ConvertIntToInt32(in.Mtu)
	// MISSING: Name
	out.NetworkFirewallPolicyEnforcementOrder = in.NetworkFirewallPolicyEnforcementOrder
	// MISSING: NetworkProfile
	// MISSING: Params
	// MISSING: Peerings
	if in.RoutingMode != nil {
		out.RoutingConfig = &pb.NetworkRoutingConfig{
			RoutingMode: in.RoutingMode,
		}
	}
	// MISSING: SelfLinkWithID
	// MISSING: Subnetworks
	return out
}
