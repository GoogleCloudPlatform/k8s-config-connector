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

// +tool:controller
// proto.service: google.cloud.compute.v1.Addresses
// proto.service: google.cloud.compute.v1.GlobalAddresses
// proto.message: google.cloud.compute.v1.Address
// crd.type: ComputeAddress
// crd.version: v1beta1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// +generated:mapper
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1

func ComputeAddressSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Address) *krm.ComputeAddressSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeAddressSpec{}
	out.Address = in.Address
	out.AddressType = in.AddressType
	// MISSING: CreationTimestamp
	out.Description = in.Description
	// MISSING: ID
	out.IPVersion = in.IpVersion
	out.IPV6EndpointType = in.Ipv6EndpointType
	// MISSING: Kind
	// MISSING: LabelFingerprint
	// MISSING: Labels
	// MISSING: Name
	if in.GetNetwork() != "" {
		out.NetworkRef = &computerefs.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.NetworkTier = in.NetworkTier
	out.PrefixLength = in.PrefixLength
	out.Purpose = in.Purpose
	// MISSING: Region
	// MISSING: SelfLink
	// MISSING: Status
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &krm.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	if in.GetIpCollection() != "" {
		out.IpCollectionRef = &krm.ComputePublicDelegatedPrefixRef{External: in.GetIpCollection()}
	}
	// MISSING: Users
	return out
}

func ComputeAddressSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeAddressSpec) *pb.Address {
	if in == nil {
		return nil
	}
	out := &pb.Address{}
	out.Address = in.Address
	out.AddressType = in.AddressType
	// MISSING: CreationTimestamp
	out.Description = in.Description
	// MISSING: ID
	out.IpVersion = in.IPVersion
	out.Ipv6EndpointType = in.IPV6EndpointType
	// MISSING: Kind
	// MISSING: LabelFingerprint
	// MISSING: Labels
	// MISSING: Name
	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}
	out.NetworkTier = in.NetworkTier
	out.PrefixLength = in.PrefixLength
	out.Purpose = in.Purpose
	// MISSING: Region
	// MISSING: SelfLink
	// MISSING: Status
	if in.SubnetworkRef != nil {
		out.Subnetwork = &in.SubnetworkRef.External
	}
	if in.IpCollectionRef != nil {
		out.IpCollection = &in.IpCollectionRef.External
	}
	// MISSING: Users
	return out
}
