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

package networkconnectivity

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func LinkedInterconnectAttachments_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.LinkedInterconnectAttachments) *krm.LinkedInterconnectAttachments {
	if in == nil {
		return nil
	}
	out := &krm.LinkedInterconnectAttachments{}
	out.Uris = LinkedInterconnectAttachments_Uris_v1beta1_FromProto(mapCtx, in.GetUris())
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	return out
}
func LinkedInterconnectAttachments_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.LinkedInterconnectAttachments) *pb.LinkedInterconnectAttachments {
	if in == nil {
		return nil
	}
	out := &pb.LinkedInterconnectAttachments{}
	out.Uris = LinkedInterconnectAttachments_Uris_v1beta1_ToProto(mapCtx, in.Uris)
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	return out
}

func LinkedVPNTunnels_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.LinkedVpnTunnels) *krm.LinkedVPNTunnels {
	if in == nil {
		return nil
	}
	out := &krm.LinkedVPNTunnels{}
	out.Uris = LinkedVPNTunnels_Uris_v1beta1_FromProto(mapCtx, in.GetUris())
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	return out
}
func LinkedVPNTunnels_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.LinkedVPNTunnels) *pb.LinkedVpnTunnels {
	if in == nil {
		return nil
	}
	out := &pb.LinkedVpnTunnels{}
	out.Uris = LinkedVPNTunnels_Uris_v1beta1_ToProto(mapCtx, in.Uris)
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	return out
}

func LinkedRouterApplianceInstances_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.LinkedRouterApplianceInstances) *krm.LinkedRouterApplianceInstances {
	if in == nil {
		return nil
	}
	out := &krm.LinkedRouterApplianceInstances{}
	out.Instances = direct.Slice_FromProto(mapCtx, in.GetInstances(), RouterApplianceInstance_v1beta1_FromProto)
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	return out
}
func LinkedRouterApplianceInstances_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.LinkedRouterApplianceInstances) *pb.LinkedRouterApplianceInstances {
	if in == nil {
		return nil
	}
	out := &pb.LinkedRouterApplianceInstances{}
	out.Instances = direct.Slice_ToProto(mapCtx, in.Instances, RouterApplianceInstance_v1beta1_ToProto)
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	return out
}

func RouterApplianceInstance_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RouterApplianceInstance) *krm.RouterApplianceInstance {
	if in == nil {
		return nil
	}
	out := &krm.RouterApplianceInstance{}
	out.IpAddress = direct.LazyPtr(in.GetIpAddress())
	out.VirtualMachineRef = RouterApplianceInstance_VirtualMachineRef_v1beta1_FromProto(mapCtx, in.GetVirtualMachine())
	return out
}
func RouterApplianceInstance_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.RouterApplianceInstance) *pb.RouterApplianceInstance {
	if in == nil {
		return nil
	}
	out := &pb.RouterApplianceInstance{}
	out.IpAddress = direct.ValueOf(in.IpAddress)
	out.VirtualMachine = RouterApplianceInstance_VirtualMachineRef_v1beta1_ToProto(mapCtx, in.VirtualMachineRef)
	return out
}

func LinkedVPCNetwork_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.LinkedVpcNetwork) *krm.LinkedVPCNetwork {
	if in == nil {
		return nil
	}
	out := &krm.LinkedVPCNetwork{}
	out.ExcludeExportRanges = in.GetExcludeExportRanges()
	out.UriRef = LinkedVPCNetwork_UriRef_v1beta1_FromProto(mapCtx, in.GetUri())
	return out
}
func LinkedVPCNetwork_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.LinkedVPCNetwork) *pb.LinkedVpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.LinkedVpcNetwork{}
	out.ExcludeExportRanges = in.ExcludeExportRanges
	out.Uri = LinkedVPCNetwork_UriRef_v1beta1_ToProto(mapCtx, in.UriRef)
	return out
}

func NetworkConnectivitySpokeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Spoke) *krm.NetworkConnectivitySpokeSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivitySpokeSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.GetHub() != "" {
		out.HubRef = &krm.SpokeHubRef{External: in.GetHub()}
	}
	out.LinkedInterconnectAttachments = LinkedInterconnectAttachments_v1beta1_FromProto(mapCtx, in.GetLinkedInterconnectAttachments())
	out.LinkedRouterApplianceInstances = LinkedRouterApplianceInstances_v1beta1_FromProto(mapCtx, in.GetLinkedRouterApplianceInstances())
	out.LinkedVPCNetwork = LinkedVPCNetwork_v1beta1_FromProto(mapCtx, in.GetLinkedVpcNetwork())
	out.LinkedVpnTunnels = LinkedVPNTunnels_v1beta1_FromProto(mapCtx, in.GetLinkedVpnTunnels())
	return out
}
func NetworkConnectivitySpokeSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivitySpokeSpec) *pb.Spoke {
	if in == nil {
		return nil
	}
	out := &pb.Spoke{}
	out.Description = direct.ValueOf(in.Description)
	if in.HubRef != nil {
		out.Hub = in.HubRef.External
	}
	out.LinkedInterconnectAttachments = LinkedInterconnectAttachments_v1beta1_ToProto(mapCtx, in.LinkedInterconnectAttachments)
	out.LinkedRouterApplianceInstances = LinkedRouterApplianceInstances_v1beta1_ToProto(mapCtx, in.LinkedRouterApplianceInstances)
	out.LinkedVpcNetwork = LinkedVPCNetwork_v1beta1_ToProto(mapCtx, in.LinkedVPCNetwork)
	out.LinkedVpnTunnels = LinkedVPNTunnels_v1beta1_ToProto(mapCtx, in.LinkedVpnTunnels)
	return out
}

func SpokeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Spoke) *krm.SpokeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpokeObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.UniqueId = direct.LazyPtr(in.GetUniqueId())
	out.State = direct.LazyPtr(in.GetState())
	return out
}
func SpokeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpokeObservedState) *pb.Spoke {
	if in == nil {
		return nil
	}
	out := &pb.Spoke{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.UniqueId = direct.ValueOf(in.UniqueId)
	out.State = direct.ValueOf(in.State)
	return out
}

// Helpers for references
func LinkedInterconnectAttachments_Uris_v1beta1_FromProto(mapCtx *direct.MapContext, in []string) []computev1beta1.ComputeInterconnectAttachmentRef {
	if in == nil {
		return nil
	}
	out := make([]computev1beta1.ComputeInterconnectAttachmentRef, len(in))
	for i, s := range in {
		out[i] = computev1beta1.ComputeInterconnectAttachmentRef{External: s}
	}
	return out
}
func LinkedInterconnectAttachments_Uris_v1beta1_ToProto(mapCtx *direct.MapContext, in []computev1beta1.ComputeInterconnectAttachmentRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, ref := range in {
		out[i] = ref.External
	}
	return out
}
func LinkedVPNTunnels_Uris_v1beta1_FromProto(mapCtx *direct.MapContext, in []string) []computev1beta1.ComputeVPNTunnelRef {
	if in == nil {
		return nil
	}
	out := make([]computev1beta1.ComputeVPNTunnelRef, len(in))
	for i, s := range in {
		out[i] = computev1beta1.ComputeVPNTunnelRef{External: s}
	}
	return out
}
func LinkedVPNTunnels_Uris_v1beta1_ToProto(mapCtx *direct.MapContext, in []computev1beta1.ComputeVPNTunnelRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, ref := range in {
		out[i] = ref.External
	}
	return out
}
func RouterApplianceInstance_VirtualMachineRef_v1beta1_FromProto(mapCtx *direct.MapContext, in string) *computev1beta1.InstanceRef {
	if in == "" {
		return nil
	}
	return &computev1beta1.InstanceRef{External: in}
}
func RouterApplianceInstance_VirtualMachineRef_v1beta1_ToProto(mapCtx *direct.MapContext, in *computev1beta1.InstanceRef) string {
	if in == nil {
		return ""
	}
	return in.External
}
func LinkedVPCNetwork_UriRef_v1beta1_FromProto(mapCtx *direct.MapContext, in string) computev1beta1.ComputeNetworkRef {
	return computev1beta1.ComputeNetworkRef{External: in}
}
func LinkedVPCNetwork_UriRef_v1beta1_ToProto(mapCtx *direct.MapContext, in computev1beta1.ComputeNetworkRef) string {
	return in.External
}
