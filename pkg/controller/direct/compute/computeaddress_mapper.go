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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeAddressSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeAddressSpec) *pb.Address {
	if in == nil {
		return nil
	}
	out := &pb.Address{}
	out.Address = in.Address
	out.AddressType = in.AddressType
	out.Description = in.Description
	out.IpVersion = in.IpVersion
	out.Ipv6EndpointType = in.Ipv6EndpointType
	if in.NetworkRef != nil {
		out.Network = direct.PtrTo(in.NetworkRef.External)
	}
	out.NetworkTier = in.NetworkTier
	if in.PrefixLength != nil {
		out.PrefixLength = direct.PtrTo(int32(*in.PrefixLength))
	}
	out.Purpose = in.Purpose
	if in.SubnetworkRef != nil {
		out.Subnetwork = direct.PtrTo(in.SubnetworkRef.External)
	}
	return out
}

func ComputeAddressSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Address) *krm.ComputeAddressSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeAddressSpec{}
	out.Address = in.Address
	out.AddressType = in.AddressType
	out.Description = in.Description
	out.IpVersion = in.IpVersion
	out.Ipv6EndpointType = in.Ipv6EndpointType
	if in.GetNetwork() != "" {
		out.NetworkRef = &krm.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.NetworkTier = in.NetworkTier
	if in.PrefixLength != nil {
		out.PrefixLength = direct.PtrTo(int64(*in.PrefixLength))
	}
	out.Purpose = in.Purpose
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &refsv1beta1.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	return out
}

func ComputeAddressStatus_FromProto(mapCtx *direct.MapContext, in *pb.Address, out *krm.ComputeAddressStatus) {
	if in == nil {
		return
	}
	out.CreationTimestamp = in.CreationTimestamp
	out.LabelFingerprint = in.LabelFingerprint
	out.SelfLink = in.SelfLink
	out.Users = in.Users

	if in.Address != nil {
		if out.ObservedState == nil {
			out.ObservedState = &krm.AddressObservedStateStatus{}
		}
		out.ObservedState.Address = in.Address
	}
}
