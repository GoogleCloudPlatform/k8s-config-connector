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
	"strings"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeServiceAttachmentSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAttachment) *krm.ComputeServiceAttachmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeServiceAttachmentSpec{}
	out.ConnectionPreference = in.GetConnectionPreference()
	out.ConsumerAcceptLists = direct.Slice_FromProto(mapCtx, in.ConsumerAcceptLists, ServiceattachmentConsumerAcceptLists_v1beta1_FromProto)
	out.ConsumerRejectLists = ComputeServiceAttachmentSpec_ConsumerRejectLists_FromProto(mapCtx, in.ConsumerRejectLists)
	out.Description = in.Description
	out.EnableProxyProtocol = in.EnableProxyProtocol
	if in.GetRegion() != "" {
		tokens := strings.Split(in.GetRegion(), "/")
		out.Location = tokens[len(tokens)-1]
	}
	out.NatSubnets = ComputeServiceAttachmentSpec_NatSubnets_FromProto(mapCtx, in.NatSubnets)
	// ProjectRef is not populated from Proto (belongs to parent/KRM metadata)
	// ResourceID is not populated from Proto
	if in.GetTargetService() != "" {
		out.TargetServiceRef = refsv1beta1.ComputeForwardingRuleRef{External: in.GetTargetService()}
	}
	return out
}

func ComputeServiceAttachmentSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeServiceAttachmentSpec) *pb.ServiceAttachment {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAttachment{}
	out.ConnectionPreference = direct.LazyPtr(in.ConnectionPreference)
	out.ConsumerAcceptLists = direct.Slice_ToProto(mapCtx, in.ConsumerAcceptLists, ServiceattachmentConsumerAcceptLists_v1beta1_ToProto)
	out.ConsumerRejectLists = ComputeServiceAttachmentSpec_ConsumerRejectLists_ToProto(mapCtx, in.ConsumerRejectLists)
	out.Description = in.Description
	out.EnableProxyProtocol = in.EnableProxyProtocol
	out.Region = direct.LazyPtr(in.Location)
	out.NatSubnets = ComputeServiceAttachmentSpec_NatSubnets_ToProto(mapCtx, in.NatSubnets)
	// TargetService
	if in.TargetServiceRef.External != "" {
		out.TargetService = direct.LazyPtr(in.TargetServiceRef.External)
	}
	return out
}

func ComputeServiceAttachmentStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAttachment) *krm.ComputeServiceAttachmentStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeServiceAttachmentStatus{}
	out.ConnectedEndpoints = direct.Slice_FromProto(mapCtx, in.ConnectedEndpoints, ServiceattachmentConnectedEndpointsStatus_v1beta1_FromProto)
	out.Fingerprint = in.Fingerprint
	if in.Id != nil {
		out.Id = direct.LazyPtr(int64(*in.Id))
	}
	out.PscServiceAttachmentId = ServiceattachmentPscServiceAttachmentIdStatus_v1beta1_FromProto(mapCtx, in.PscServiceAttachmentId)
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	return out
}

func ComputeServiceAttachmentStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeServiceAttachmentStatus) *pb.ServiceAttachment {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAttachment{}
	out.ConnectedEndpoints = direct.Slice_ToProto(mapCtx, in.ConnectedEndpoints, ServiceattachmentConnectedEndpointsStatus_v1beta1_ToProto)
	out.Fingerprint = in.Fingerprint
	if in.Id != nil {
		out.Id = direct.PtrInt64ToPtrUint64(in.Id)
	}
	out.PscServiceAttachmentId = ServiceattachmentPscServiceAttachmentIdStatus_v1beta1_ToProto(mapCtx, in.PscServiceAttachmentId)
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	return out
}

func ServiceattachmentConsumerAcceptLists_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAttachmentConsumerProjectLimit) *krm.ServiceattachmentConsumerAcceptLists {
	if in == nil {
		return nil
	}
	out := &krm.ServiceattachmentConsumerAcceptLists{}
	if in.ConnectionLimit != nil {
		out.ConnectionLimit = direct.LazyPtr(int64(*in.ConnectionLimit))
	}
	out.ProjectRef.External = in.GetProjectIdOrNum()
	return out
}

func ServiceattachmentConsumerAcceptLists_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ServiceattachmentConsumerAcceptLists) *pb.ServiceAttachmentConsumerProjectLimit {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAttachmentConsumerProjectLimit{}
	if in.ConnectionLimit != nil {
		out.ConnectionLimit = direct.LazyPtr(uint32(*in.ConnectionLimit))
	}
	if in.ProjectRef.External == "" && in.ProjectRef.Name != "" {
		mapCtx.Errorf("reference projectRef was not pre-resolved")
	}
	out.ProjectIdOrNum = direct.LazyPtr(in.ProjectRef.External)
	return out
}

func ComputeServiceAttachmentSpec_ConsumerRejectLists_FromProto(mapCtx *direct.MapContext, in []string) []refs.ProjectRef {
	if in == nil {
		return nil
	}
	var out []refs.ProjectRef
	for _, i := range in {
		out = append(out, refs.ProjectRef{
			External: i,
		})
	}
	return out
}

func ComputeServiceAttachmentSpec_ConsumerRejectLists_ToProto(mapCtx *direct.MapContext, in []refs.ProjectRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, i := range in {
		if i.External == "" && i.Name != "" {
			mapCtx.Errorf("reference %s was not pre-resolved", i.Name)
		}
		out = append(out, i.External)
	}
	return out
}

func ComputeServiceAttachmentSpec_NatSubnets_FromProto(mapCtx *direct.MapContext, in []string) []krm.ComputeSubnetworkRef {
	if in == nil {
		return nil
	}
	var out []krm.ComputeSubnetworkRef
	for _, i := range in {
		out = append(out, krm.ComputeSubnetworkRef{
			External: i,
		})
	}
	return out
}

func ComputeServiceAttachmentSpec_NatSubnets_ToProto(mapCtx *direct.MapContext, in []krm.ComputeSubnetworkRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, i := range in {
		if i.External == "" && i.Name != "" {
			mapCtx.Errorf("reference %s was not pre-resolved", i.Name)
		}
		out = append(out, i.External)
	}
	return out
}

func ServiceattachmentConnectedEndpointsStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAttachmentConnectedEndpoint) *krm.ServiceattachmentConnectedEndpointsStatus {
	if in == nil {
		return nil
	}
	out := &krm.ServiceattachmentConnectedEndpointsStatus{}
	out.Endpoint = in.Endpoint
	if in.PscConnectionId != nil {
		out.PscConnectionId = direct.LazyPtr(int64(*in.PscConnectionId))
	}
	out.Status = in.Status
	return out
}

func ServiceattachmentConnectedEndpointsStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ServiceattachmentConnectedEndpointsStatus) *pb.ServiceAttachmentConnectedEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAttachmentConnectedEndpoint{}
	out.Endpoint = in.Endpoint
	if in.PscConnectionId != nil {
		out.PscConnectionId = direct.LazyPtr(uint64(*in.PscConnectionId))
	}
	out.Status = in.Status
	return out
}

func ServiceattachmentPscServiceAttachmentIdStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Uint128) *krm.ServiceattachmentPscServiceAttachmentIdStatus {
	if in == nil {
		return nil
	}
	out := &krm.ServiceattachmentPscServiceAttachmentIdStatus{}
	if in.High != nil {
		out.High = direct.LazyPtr(int64(*in.High))
	}
	if in.Low != nil {
		out.Low = direct.LazyPtr(int64(*in.Low))
	}
	return out
}

func ServiceattachmentPscServiceAttachmentIdStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ServiceattachmentPscServiceAttachmentIdStatus) *pb.Uint128 {
	if in == nil {
		return nil
	}
	out := &pb.Uint128{}
	if in.High != nil {
		out.High = direct.LazyPtr(uint64(*in.High))
	}
	if in.Low != nil {
		out.Low = direct.LazyPtr(uint64(*in.Low))
	}
	return out
}
