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

// ComputePacketMirroringSpec_v1beta1_FromProto is handcoded because we need custom
// handling for value structs, reference types, and int64-to-uint32 conversions (Priority).
func ComputePacketMirroringSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PacketMirroring) *krm.ComputePacketMirroringSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputePacketMirroringSpec{}
	out.CollectorIlb = PacketmirroringCollectorIlb_v1beta1_FromProto(mapCtx, in.CollectorIlb)
	out.Description = in.Description
	out.Enable = in.Enable
	if in.Filter != nil {
		out.Filter = PacketmirroringFilter_v1beta1_FromProto(mapCtx, in.Filter)
	}
	if in.Region != nil {
		out.Location = *in.Region
	}
	out.MirroredResources = PacketmirroringMirroredResources_v1beta1_FromProto(mapCtx, in.MirroredResources)
	out.Network = PacketmirroringNetwork_v1beta1_FromProto(mapCtx, in.Network)
	if in.Priority != nil {
		val := int64(*in.Priority)
		out.Priority = &val
	}
	return out
}

// ComputePacketMirroringSpec_v1beta1_ToProto is handcoded because we need custom
// handling for value structs, reference types, and int64-to-uint32 conversions (Priority).
func ComputePacketMirroringSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputePacketMirroringSpec) *pb.PacketMirroring {
	if in == nil {
		return nil
	}
	out := &pb.PacketMirroring{}
	out.CollectorIlb = PacketmirroringCollectorIlb_v1beta1_ToProto(mapCtx, &in.CollectorIlb)
	out.Description = in.Description
	out.Enable = in.Enable
	out.Filter = PacketmirroringFilter_v1beta1_ToProto(mapCtx, in.Filter)
	out.MirroredResources = PacketmirroringMirroredResources_v1beta1_ToProto(mapCtx, &in.MirroredResources)
	out.Network = PacketmirroringNetwork_v1beta1_ToProto(mapCtx, &in.Network)
	if in.Priority != nil {
		val := uint32(*in.Priority)
		out.Priority = &val
	}
	return out
}

// PacketmirroringCollectorIlb_v1beta1_FromProto is handcoded to map references from proto string URLs.
func PacketmirroringCollectorIlb_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PacketMirroringForwardingRuleInfo) krm.PacketmirroringCollectorIlb {
	if in == nil {
		return krm.PacketmirroringCollectorIlb{}
	}
	return krm.PacketmirroringCollectorIlb{
		UrlRef: krm.ForwardingRuleRef{
			External: in.GetUrl(),
		},
	}
}

// PacketmirroringCollectorIlb_v1beta1_ToProto is handcoded to map references to proto string URLs.
func PacketmirroringCollectorIlb_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.PacketmirroringCollectorIlb) *pb.PacketMirroringForwardingRuleInfo {
	if in == nil || in.UrlRef.External == "" {
		return nil
	}
	out := &pb.PacketMirroringForwardingRuleInfo{}
	out.Url = &in.UrlRef.External
	return out
}

// PacketmirroringFilter_v1beta1_FromProto is handcoded to map fields with exact cases.
func PacketmirroringFilter_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PacketMirroringFilter) *krm.PacketmirroringFilter {
	if in == nil {
		return nil
	}
	out := &krm.PacketmirroringFilter{}
	out.CidrRanges = in.CidrRanges
	out.Direction = in.Direction
	out.IpProtocols = in.IPProtocols
	return out
}

// PacketmirroringFilter_v1beta1_ToProto is handcoded to map fields with exact cases.
func PacketmirroringFilter_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.PacketmirroringFilter) *pb.PacketMirroringFilter {
	if in == nil {
		return nil
	}
	out := &pb.PacketMirroringFilter{}
	out.CidrRanges = in.CidrRanges
	out.Direction = in.Direction
	out.IPProtocols = in.IpProtocols
	return out
}

// PacketmirroringMirroredResources_v1beta1_FromProto is handcoded to iterate nested slices and populate structs.
func PacketmirroringMirroredResources_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PacketMirroringMirroredResourceInfo) krm.PacketmirroringMirroredResources {
	if in == nil {
		return krm.PacketmirroringMirroredResources{}
	}
	out := krm.PacketmirroringMirroredResources{}
	for _, inst := range in.Instances {
		out.Instances = append(out.Instances, PacketmirroringInstances_v1beta1_FromProto(mapCtx, inst))
	}
	for _, subnet := range in.Subnetworks {
		out.Subnetworks = append(out.Subnetworks, PacketmirroringSubnetworks_v1beta1_FromProto(mapCtx, subnet))
	}
	out.Tags = in.Tags
	return out
}

// PacketmirroringMirroredResources_v1beta1_ToProto is handcoded to iterate nested slices and populate proto structs.
func PacketmirroringMirroredResources_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.PacketmirroringMirroredResources) *pb.PacketMirroringMirroredResourceInfo {
	if in == nil || (len(in.Instances) == 0 && len(in.Subnetworks) == 0 && len(in.Tags) == 0) {
		return nil
	}
	out := &pb.PacketMirroringMirroredResourceInfo{}
	for _, inst := range in.Instances {
		out.Instances = append(out.Instances, PacketmirroringInstances_v1beta1_ToProto(mapCtx, &inst))
	}
	for _, subnet := range in.Subnetworks {
		out.Subnetworks = append(out.Subnetworks, PacketmirroringSubnetworks_v1beta1_ToProto(mapCtx, &subnet))
	}
	out.Tags = in.Tags
	return out
}

// PacketmirroringInstances_v1beta1_FromProto is handcoded to map references from proto string URLs.
func PacketmirroringInstances_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PacketMirroringMirroredResourceInfoInstanceInfo) krm.PacketmirroringInstances {
	if in == nil {
		return krm.PacketmirroringInstances{}
	}
	out := krm.PacketmirroringInstances{}
	out.CanonicalUrl = in.CanonicalUrl
	if in.Url != nil {
		out.UrlRef = &krm.InstanceRef{
			External: *in.Url,
		}
	}
	return out
}

// PacketmirroringInstances_v1beta1_ToProto is handcoded to map references to proto string URLs.
func PacketmirroringInstances_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.PacketmirroringInstances) *pb.PacketMirroringMirroredResourceInfoInstanceInfo {
	if in == nil || (in.CanonicalUrl == nil && (in.UrlRef == nil || in.UrlRef.External == "")) {
		return nil
	}
	out := &pb.PacketMirroringMirroredResourceInfoInstanceInfo{}
	out.CanonicalUrl = in.CanonicalUrl
	if in.UrlRef != nil && in.UrlRef.External != "" {
		out.Url = &in.UrlRef.External
	}
	return out
}

// PacketmirroringSubnetworks_v1beta1_FromProto is handcoded to map references from proto string URLs.
func PacketmirroringSubnetworks_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PacketMirroringMirroredResourceInfoSubnetInfo) krm.PacketmirroringSubnetworks {
	if in == nil {
		return krm.PacketmirroringSubnetworks{}
	}
	out := krm.PacketmirroringSubnetworks{}
	out.CanonicalUrl = in.CanonicalUrl
	if in.Url != nil {
		out.UrlRef = &krm.ComputeSubnetworkRef{
			External: *in.Url,
		}
	}
	return out
}

// PacketmirroringSubnetworks_v1beta1_ToProto is handcoded to map references to proto string URLs.
func PacketmirroringSubnetworks_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.PacketmirroringSubnetworks) *pb.PacketMirroringMirroredResourceInfoSubnetInfo {
	if in == nil || (in.CanonicalUrl == nil && (in.UrlRef == nil || in.UrlRef.External == "")) {
		return nil
	}
	out := &pb.PacketMirroringMirroredResourceInfoSubnetInfo{}
	out.CanonicalUrl = in.CanonicalUrl
	if in.UrlRef != nil && in.UrlRef.External != "" {
		out.Url = &in.UrlRef.External
	}
	return out
}

// PacketmirroringNetwork_v1beta1_FromProto is handcoded to map references from proto string URLs.
func PacketmirroringNetwork_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PacketMirroringNetworkInfo) krm.PacketmirroringNetwork {
	if in == nil {
		return krm.PacketmirroringNetwork{}
	}
	return krm.PacketmirroringNetwork{
		UrlRef: krm.ComputeNetworkRef{
			External: in.GetUrl(),
		},
	}
}

// PacketmirroringNetwork_v1beta1_ToProto is handcoded to map references to proto string URLs.
func PacketmirroringNetwork_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.PacketmirroringNetwork) *pb.PacketMirroringNetworkInfo {
	if in == nil || in.UrlRef.External == "" {
		return nil
	}
	out := &pb.PacketMirroringNetworkInfo{}
	out.Url = &in.UrlRef.External
	return out
}

// ComputePacketMirroringStatus_v1beta1_FromProto is handcoded because of custom handling for value structs and status fields.
func ComputePacketMirroringStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PacketMirroring) *krm.ComputePacketMirroringStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputePacketMirroringStatus{}
	if in.CollectorIlb != nil {
		c := PacketmirroringCollectorIlbStatus_v1beta1_FromProto(mapCtx, in.CollectorIlb)
		out.CollectorIlb = &c
	}
	if in.Id != nil {
		val := int64(*in.Id)
		out.Id = &val
	}
	if in.Network != nil {
		n := PacketmirroringNetworkStatus_v1beta1_FromProto(mapCtx, in.Network)
		out.Network = &n
	}
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	return out
}

// ComputePacketMirroringStatus_v1beta1_ToProto is handcoded because of custom handling for value structs and status fields.
func ComputePacketMirroringStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputePacketMirroringStatus) *pb.PacketMirroring {
	if in == nil {
		return nil
	}
	out := &pb.PacketMirroring{}
	if in.CollectorIlb != nil {
		out.CollectorIlb = PacketmirroringCollectorIlbStatus_v1beta1_ToProto(mapCtx, in.CollectorIlb)
	}
	if in.Id != nil {
		val := uint64(*in.Id)
		out.Id = &val
	}
	if in.Network != nil {
		out.Network = PacketmirroringNetworkStatus_v1beta1_ToProto(mapCtx, in.Network)
	}
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	return out
}

// PacketmirroringCollectorIlbStatus_v1beta1_FromProto maps forwarding rule info status.
func PacketmirroringCollectorIlbStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PacketMirroringForwardingRuleInfo) krm.PacketmirroringCollectorIlbStatus {
	if in == nil {
		return krm.PacketmirroringCollectorIlbStatus{}
	}
	out := krm.PacketmirroringCollectorIlbStatus{}
	out.CanonicalUrl = in.CanonicalUrl
	return out
}

// PacketmirroringCollectorIlbStatus_v1beta1_ToProto maps forwarding rule info status.
func PacketmirroringCollectorIlbStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.PacketmirroringCollectorIlbStatus) *pb.PacketMirroringForwardingRuleInfo {
	if in == nil {
		return nil
	}
	out := &pb.PacketMirroringForwardingRuleInfo{}
	out.CanonicalUrl = in.CanonicalUrl
	return out
}

// PacketmirroringNetworkStatus_v1beta1_FromProto maps network info status.
func PacketmirroringNetworkStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PacketMirroringNetworkInfo) krm.PacketmirroringNetworkStatus {
	if in == nil {
		return krm.PacketmirroringNetworkStatus{}
	}
	out := krm.PacketmirroringNetworkStatus{}
	out.CanonicalUrl = in.CanonicalUrl
	return out
}

// PacketmirroringNetworkStatus_v1beta1_ToProto maps network info status.
func PacketmirroringNetworkStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.PacketmirroringNetworkStatus) *pb.PacketMirroringNetworkInfo {
	if in == nil {
		return nil
	}
	out := &pb.PacketMirroringNetworkInfo{}
	out.CanonicalUrl = in.CanonicalUrl
	return out
}
