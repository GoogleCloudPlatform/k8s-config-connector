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

// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.compute.v1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeNetworkAttachmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkAttachment) *krm.ComputeNetworkAttachmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkAttachmentObservedState{}
	out.ConnectionEndpoints = direct.Slice_FromProto(mapCtx, in.ConnectionEndpoints, NetworkAttachmentConnectedEndpoint_FromProto)
	out.CreationTimestamp = in.CreationTimestamp
	out.ID = in.Id
	out.Kind = in.Kind
	// MISSING: Name
	out.Network = in.Network
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	out.SelfLinkWithID = in.SelfLinkWithId
	return out
}
func ComputeNetworkAttachmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkAttachmentObservedState) *pb.NetworkAttachment {
	if in == nil {
		return nil
	}
	out := &pb.NetworkAttachment{}
	out.ConnectionEndpoints = direct.Slice_ToProto(mapCtx, in.ConnectionEndpoints, NetworkAttachmentConnectedEndpoint_ToProto)
	out.CreationTimestamp = in.CreationTimestamp
	out.Id = in.ID
	out.Kind = in.Kind
	// MISSING: Name
	out.Network = in.Network
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithID
	return out
}
func ComputeNetworkAttachmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.NetworkAttachment) *krm.ComputeNetworkAttachmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkAttachmentSpec{}
	out.ConnectionPreference = in.ConnectionPreference
	out.Description = in.Description
	out.Fingerprint = in.Fingerprint
	// MISSING: Name
	out.ProducerAcceptLists = ProjectList_FromProto(mapCtx, in.ProducerAcceptLists)
	out.ProducerRejectLists = ProjectList_FromProto(mapCtx, in.ProducerRejectLists)
	out.SubnetworkRefs = SubnetworkList_FromProto(mapCtx, in.Subnetworks)
	return out
}
func ComputeNetworkAttachmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkAttachmentSpec) *pb.NetworkAttachment {
	if in == nil {
		return nil
	}
	out := &pb.NetworkAttachment{}
	out.ConnectionPreference = in.ConnectionPreference
	out.Description = in.Description
	out.Fingerprint = in.Fingerprint
	// MISSING: Name
	out.ProducerAcceptLists = ProjectList_ToProto(mapCtx, in.ProducerAcceptLists)
	out.ProducerRejectLists = ProjectList_ToProto(mapCtx, in.ProducerRejectLists)
	out.Subnetworks = SubnetworkList_ToProto(mapCtx, in.SubnetworkRefs)
	return out
}
func NetworkAttachmentConnectedEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.NetworkAttachmentConnectedEndpoint) *krm.NetworkAttachmentConnectedEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.NetworkAttachmentConnectedEndpoint{}
	out.IPAddress = in.IpAddress
	out.IPV6Address = in.Ipv6Address
	out.ProjectIDOrNum = in.ProjectIdOrNum
	out.SecondaryIPCIDRRanges = in.SecondaryIpCidrRanges
	out.Status = in.Status
	out.Subnetwork = in.Subnetwork
	out.SubnetworkCIDRRange = in.SubnetworkCidrRange
	return out
}
func NetworkAttachmentConnectedEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.NetworkAttachmentConnectedEndpoint) *pb.NetworkAttachmentConnectedEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.NetworkAttachmentConnectedEndpoint{}
	out.IpAddress = in.IPAddress
	out.Ipv6Address = in.IPV6Address
	out.ProjectIdOrNum = in.ProjectIDOrNum
	out.SecondaryIpCidrRanges = in.SecondaryIPCIDRRanges
	out.Status = in.Status
	out.Subnetwork = in.Subnetwork
	out.SubnetworkCidrRange = in.SubnetworkCIDRRange
	return out
}

func ProjectList_FromProto(mapCtx *direct.MapContext, in []string) []*refs.ProjectRef {
	if in == nil {
		return nil
	}
	var out []*refs.ProjectRef
	for _, i := range in {
		out = append(out, &refs.ProjectRef{
			External: i,
		})
	}
	return out
}

func ProjectList_ToProto(mapCtx *direct.MapContext, in []*refs.ProjectRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, i := range in {
		if i == nil {
			continue
		}
		if i.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", i.Name)
		}
		out = append(out, i.External)
	}
	return out
}

func SubnetworkList_FromProto(mapCtx *direct.MapContext, in []string) []*refs.ComputeSubnetworkRef {
	if in == nil {
		return nil
	}
	var out []*refs.ComputeSubnetworkRef
	for _, i := range in {
		out = append(out, &refs.ComputeSubnetworkRef{
			External: i,
		})
	}
	return out
}

func SubnetworkList_ToProto(mapCtx *direct.MapContext, in []*refs.ComputeSubnetworkRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, i := range in {
		if i == nil {
			continue
		}
		if i.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", i.Name)
		}
		out = append(out, i.External)
	}
	return out
}
