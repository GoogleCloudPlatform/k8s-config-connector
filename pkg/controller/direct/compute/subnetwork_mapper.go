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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeSubnetworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.Subnetwork) *krm.ComputeSubnetworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeSubnetworkSpec{}
	// MISSING: CreationTimestamp
	out.Description = in.Description
	// MISSING: EnableFlowLogs
	// MISSING: ExternalIPV6Prefix
	// MISSING: Fingerprint
	// MISSING: GatewayAddress
	// MISSING: ID
	// MISSING: InternalIPV6Prefix
	out.IPCIDRRange = in.IpCidrRange
	// MISSING: IPCollection
	out.IPV6AccessType = in.Ipv6AccessType
	// MISSING: IPV6CIDRRange
	// MISSING: IPV6GCEEndpoint
	// MISSING: Kind
	out.LogConfig = SubnetworkLogConfig_FromProto(mapCtx, in.GetLogConfig())
	// MISSING: Name
	if in.GetNetwork() != "" {
		out.NetworkRef = &krm.ComputeNetworkRef{External: in.GetNetwork()}
	}
	// MISSING: Params
	out.PrivateIPGoogleAccess = in.PrivateIpGoogleAccess
	out.PrivateIPV6GoogleAccess = in.PrivateIpv6GoogleAccess
	out.Purpose = in.Purpose
	out.Region = in.Region
	// MISSING: ReservedInternalRange
	out.Role = in.Role
	out.SecondaryIPRanges = direct.Slice_FromProto(mapCtx, in.SecondaryIpRanges, SubnetworkSecondaryRange_FromProto)
	// MISSING: SelfLink
	out.StackType = in.StackType
	// MISSING: State
	// MISSING: SystemReservedExternalIPV6Ranges
	// MISSING: SystemReservedInternalIPV6Ranges
	return out
}
func ComputeSubnetworkSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeSubnetworkSpec) *pb.Subnetwork {
	if in == nil {
		return nil
	}
	out := &pb.Subnetwork{}
	// MISSING: CreationTimestamp
	out.Description = in.Description
	// MISSING: EnableFlowLogs
	// MISSING: ExternalIPV6Prefix
	// MISSING: Fingerprint
	// MISSING: GatewayAddress
	// MISSING: ID
	// MISSING: InternalIPV6Prefix
	out.IpCidrRange = in.IPCIDRRange
	// MISSING: IPCollection
	out.Ipv6AccessType = in.IPV6AccessType
	// MISSING: IPV6CIDRRange
	// MISSING: IPV6GCEEndpoint
	// MISSING: Kind
	out.LogConfig = SubnetworkLogConfig_ToProto(mapCtx, in.LogConfig)
	// MISSING: Name
	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}
	// MISSING: Params
	out.PrivateIpGoogleAccess = in.PrivateIPGoogleAccess
	out.PrivateIpv6GoogleAccess = in.PrivateIPV6GoogleAccess
	out.Purpose = in.Purpose
	out.Region = in.Region
	// MISSING: ReservedInternalRange
	out.Role = in.Role
	out.SecondaryIpRanges = direct.Slice_ToProto(mapCtx, in.SecondaryIPRanges, SubnetworkSecondaryRange_ToProto)
	// MISSING: SelfLink
	out.StackType = in.StackType
	// MISSING: State
	// MISSING: SystemReservedExternalIPV6Ranges
	// MISSING: SystemReservedInternalIPV6Ranges
	return out
}
