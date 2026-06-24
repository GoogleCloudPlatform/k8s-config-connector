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
// See the License for the "Licensed" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compute

import (
	"strconv"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeInterconnectAttachmentSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectAttachment) *krm.ComputeInterconnectAttachmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInterconnectAttachmentSpec{}
	out.AdminEnabled = in.AdminEnabled
	out.Bandwidth = in.Bandwidth
	out.CandidateSubnets = in.CandidateSubnets
	out.Description = in.Description
	out.EdgeAvailabilityDomain = in.EdgeAvailabilityDomain
	out.Encryption = in.Encryption
	out.Interconnect = in.Interconnect
	out.IpsecInternalAddresses = ComputeInterconnectAttachmentSpec_IpsecInternalAddresses_FromProto(mapCtx, in.IpsecInternalAddresses)
	out.Mtu = Mtu_FromProto(mapCtx, in.Mtu)
	out.Region = Region_FromProto(mapCtx, in.Region)
	out.RouterRef = ComputeRouterRef_FromProto(mapCtx, in.Router)
	out.Type = in.Type
	out.VlanTag8021q = VlanTag8021q_FromProto(mapCtx, in.VlanTag8021Q)
	return out
}

func ComputeInterconnectAttachmentSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInterconnectAttachmentSpec) *pb.InterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectAttachment{}
	out.AdminEnabled = in.AdminEnabled
	out.Bandwidth = in.Bandwidth
	out.CandidateSubnets = in.CandidateSubnets
	out.Description = in.Description
	out.EdgeAvailabilityDomain = in.EdgeAvailabilityDomain
	out.Encryption = in.Encryption
	out.Interconnect = in.Interconnect
	out.IpsecInternalAddresses = ComputeInterconnectAttachmentSpec_IpsecInternalAddresses_ToProto(mapCtx, in.IpsecInternalAddresses)
	out.Mtu = Mtu_ToProto(mapCtx, in.Mtu)
	out.Region = Region_ToProto(mapCtx, in.Region)
	out.Router = ComputeRouterRef_ToProto(mapCtx, in.RouterRef)
	out.Type = in.Type
	out.VlanTag8021Q = VlanTag8021q_ToProto(mapCtx, in.VlanTag8021q)
	return out
}

func ComputeInterconnectAttachmentStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectAttachment) *krm.ComputeInterconnectAttachmentStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInterconnectAttachmentStatus{}
	out.CloudRouterIpAddress = in.CloudRouterIpAddress
	out.CreationTimestamp = in.CreationTimestamp
	out.CustomerRouterIpAddress = in.CustomerRouterIpAddress
	out.GoogleReferenceId = in.GoogleReferenceId
	out.PairingKey = in.PairingKey
	out.PartnerAsn = PartnerAsn_FromProto(mapCtx, in.PartnerAsn)
	out.PrivateInterconnectInfo = PrivateInterconnectInfo_FromProto(mapCtx, in.PrivateInterconnectInfo)
	out.SelfLink = in.SelfLink
	out.State = in.State
	return out
}

func ComputeInterconnectAttachmentStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInterconnectAttachmentStatus) *pb.InterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectAttachment{}
	out.CloudRouterIpAddress = in.CloudRouterIpAddress
	out.CreationTimestamp = in.CreationTimestamp
	out.CustomerRouterIpAddress = in.CustomerRouterIpAddress
	out.GoogleReferenceId = in.GoogleReferenceId
	out.PairingKey = in.PairingKey
	out.PartnerAsn = PartnerAsn_ToProto(mapCtx, in.PartnerAsn)
	out.PrivateInterconnectInfo = PrivateInterconnectInfo_ToProto(mapCtx, in.PrivateInterconnectInfo)
	out.SelfLink = in.SelfLink
	out.State = in.State
	return out
}

func ComputeRouterRef_FromProto(mapCtx *direct.MapContext, in *string) krm.ComputeRouterRef {
	if in == nil {
		return krm.ComputeRouterRef{}
	}
	return krm.ComputeRouterRef{
		External: *in,
	}
}

func ComputeRouterRef_ToProto(mapCtx *direct.MapContext, in krm.ComputeRouterRef) *string {
	if in.External == "" {
		if in.Name != "" {
			mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
		}
		return nil
	}
	return direct.LazyPtr(in.External)
}

func ComputeInterconnectAttachmentSpec_IpsecInternalAddresses_FromProto(mapCtx *direct.MapContext, in []string) []krm.ComputeAddressRef {
	if in == nil {
		return nil
	}
	out := make([]krm.ComputeAddressRef, len(in))
	for i, v := range in {
		out[i] = krm.ComputeAddressRef{
			External: v,
		}
	}
	return out
}

func ComputeInterconnectAttachmentSpec_IpsecInternalAddresses_ToProto(mapCtx *direct.MapContext, in []krm.ComputeAddressRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, v := range in {
		if v.External == "" {
			if v.Name != "" {
				mapCtx.Errorf("reference %s was not pre-resolved", v.Name)
			}
		}
		out[i] = v.External
	}
	return out
}

func Mtu_FromProto(mapCtx *direct.MapContext, in *int32) *string {
	if in == nil {
		return nil
	}
	val := strconv.Itoa(int(*in))
	return &val
}

func Mtu_ToProto(mapCtx *direct.MapContext, in *string) *int32 {
	if in == nil {
		return nil
	}
	val, err := strconv.Atoi(*in)
	if err != nil {
		mapCtx.Errorf("invalid mtu value %q: %v", *in, err)
		return nil
	}
	valInt32 := int32(val)
	return &valInt32
}

func VlanTag8021q_FromProto(mapCtx *direct.MapContext, in *int32) *int64 {
	if in == nil {
		return nil
	}
	val := int64(*in)
	return &val
}

func VlanTag8021q_ToProto(mapCtx *direct.MapContext, in *int64) *int32 {
	if in == nil {
		return nil
	}
	val := int32(*in)
	return &val
}

func Region_FromProto(mapCtx *direct.MapContext, in *string) string {
	if in == nil {
		return ""
	}
	return *in
}

func Region_ToProto(mapCtx *direct.MapContext, in string) *string {
	if in == "" {
		return nil
	}
	return &in
}

func PartnerAsn_FromProto(mapCtx *direct.MapContext, in *int64) *string {
	if in == nil {
		return nil
	}
	val := strconv.FormatInt(*in, 10)
	return &val
}

func PartnerAsn_ToProto(mapCtx *direct.MapContext, in *string) *int64 {
	if in == nil {
		return nil
	}
	val, err := strconv.ParseInt(*in, 10, 64)
	if err != nil {
		mapCtx.Errorf("invalid partnerAsn value %q: %v", *in, err)
		return nil
	}
	return &val
}

func PrivateInterconnectInfo_FromProto(mapCtx *direct.MapContext, in *pb.InterconnectAttachmentPrivateInfo) *krm.InterconnectattachmentPrivateInterconnectInfoStatus {
	if in == nil {
		return nil
	}
	out := &krm.InterconnectattachmentPrivateInterconnectInfoStatus{}
	if in.Tag8021Q != nil {
		val := int64(*in.Tag8021Q)
		out.Tag8021q = &val
	}
	return out
}

func PrivateInterconnectInfo_ToProto(mapCtx *direct.MapContext, in *krm.InterconnectattachmentPrivateInterconnectInfoStatus) *pb.InterconnectAttachmentPrivateInfo {
	if in == nil {
		return nil
	}
	out := &pb.InterconnectAttachmentPrivateInfo{}
	if in.Tag8021q != nil {
		val := uint32(*in.Tag8021q)
		out.Tag8021Q = &val
	}
	return out
}
