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

func ComputeNetworkSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Network) *krm.ComputeNetworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkSpec{}
	out.AutoCreateSubnetworks = in.AutoCreateSubnetworks
	out.Description = in.Description
	out.EnableUlaInternalIpv6 = in.EnableUlaInternalIpv6
	out.InternalIpv6Range = in.InternalIpv6Range
	if in.Mtu != nil {
		mtu := int64(*in.Mtu)
		out.Mtu = &mtu
	}
	out.NetworkFirewallPolicyEnforcementOrder = in.NetworkFirewallPolicyEnforcementOrder
	if in.RoutingConfig != nil {
		out.RoutingMode = in.RoutingConfig.RoutingMode
	}
	out.NetworkProfile = in.NetworkProfile
	out.ResourceID = in.Name
	return out
}

func ComputeNetworkSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkSpec) *pb.Network {
	if in == nil {
		return nil
	}
	out := &pb.Network{}
	out.AutoCreateSubnetworks = in.AutoCreateSubnetworks
	out.Description = in.Description
	out.EnableUlaInternalIpv6 = in.EnableUlaInternalIpv6
	out.InternalIpv6Range = in.InternalIpv6Range
	if in.Mtu != nil {
		mtu := int32(*in.Mtu)
		out.Mtu = &mtu
	}
	out.NetworkFirewallPolicyEnforcementOrder = in.NetworkFirewallPolicyEnforcementOrder
	if in.RoutingMode != nil {
		out.RoutingConfig = &pb.NetworkRoutingConfig{
			RoutingMode: in.RoutingMode,
		}
	}
	out.NetworkProfile = in.NetworkProfile
	out.Name = in.ResourceID
	return out
}

func ComputeNetworkObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Network) *krm.ComputeNetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkObservedState{}
	return out
}

func ComputeNetworkObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkObservedState) *pb.Network {
	if in == nil {
		return nil
	}
	out := &pb.Network{}
	return out
}
