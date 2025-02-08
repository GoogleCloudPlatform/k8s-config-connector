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

package networkconnectivity

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func LinkedInterconnectAttachments_FromProto(mapCtx *direct.MapContext, in *pb.LinkedInterconnectAttachments) *krm.LinkedInterconnectAttachments {
	if in == nil {
		return nil
	}
	out := &krm.LinkedInterconnectAttachments{}
	out.Uris = in.Uris
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	// MISSING: VpcNetwork
	out.IncludeImportRanges = in.IncludeImportRanges
	return out
}
func LinkedInterconnectAttachments_ToProto(mapCtx *direct.MapContext, in *krm.LinkedInterconnectAttachments) *pb.LinkedInterconnectAttachments {
	if in == nil {
		return nil
	}
	out := &pb.LinkedInterconnectAttachments{}
	out.Uris = in.Uris
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	// MISSING: VpcNetwork
	out.IncludeImportRanges = in.IncludeImportRanges
	return out
}
func LinkedInterconnectAttachmentsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LinkedInterconnectAttachments) *krm.LinkedInterconnectAttachmentsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LinkedInterconnectAttachmentsObservedState{}
	// MISSING: Uris
	// MISSING: SiteToSiteDataTransfer
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
	// MISSING: IncludeImportRanges
	return out
}
func LinkedInterconnectAttachmentsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LinkedInterconnectAttachmentsObservedState) *pb.LinkedInterconnectAttachments {
	if in == nil {
		return nil
	}
	out := &pb.LinkedInterconnectAttachments{}
	// MISSING: Uris
	// MISSING: SiteToSiteDataTransfer
	out.VpcNetwork = direct.ValueOf(in.VpcNetwork)
	// MISSING: IncludeImportRanges
	return out
}
func LinkedProducerVpcNetwork_FromProto(mapCtx *direct.MapContext, in *pb.LinkedProducerVpcNetwork) *krm.LinkedProducerVpcNetwork {
	if in == nil {
		return nil
	}
	out := &krm.LinkedProducerVpcNetwork{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	// MISSING: ServiceConsumerVpcSpoke
	out.Peering = direct.LazyPtr(in.GetPeering())
	// MISSING: ProducerNetwork
	out.ExcludeExportRanges = in.ExcludeExportRanges
	out.IncludeExportRanges = in.IncludeExportRanges
	return out
}
func LinkedProducerVpcNetwork_ToProto(mapCtx *direct.MapContext, in *krm.LinkedProducerVpcNetwork) *pb.LinkedProducerVpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.LinkedProducerVpcNetwork{}
	out.Network = direct.ValueOf(in.Network)
	// MISSING: ServiceConsumerVpcSpoke
	out.Peering = direct.ValueOf(in.Peering)
	// MISSING: ProducerNetwork
	out.ExcludeExportRanges = in.ExcludeExportRanges
	out.IncludeExportRanges = in.IncludeExportRanges
	return out
}
func LinkedProducerVpcNetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LinkedProducerVpcNetwork) *krm.LinkedProducerVpcNetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LinkedProducerVpcNetworkObservedState{}
	// MISSING: Network
	out.ServiceConsumerVpcSpoke = direct.LazyPtr(in.GetServiceConsumerVpcSpoke())
	// MISSING: Peering
	out.ProducerNetwork = direct.LazyPtr(in.GetProducerNetwork())
	// MISSING: ExcludeExportRanges
	// MISSING: IncludeExportRanges
	return out
}
func LinkedProducerVpcNetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LinkedProducerVpcNetworkObservedState) *pb.LinkedProducerVpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.LinkedProducerVpcNetwork{}
	// MISSING: Network
	out.ServiceConsumerVpcSpoke = direct.ValueOf(in.ServiceConsumerVpcSpoke)
	// MISSING: Peering
	out.ProducerNetwork = direct.ValueOf(in.ProducerNetwork)
	// MISSING: ExcludeExportRanges
	// MISSING: IncludeExportRanges
	return out
}
func LinkedRouterApplianceInstances_FromProto(mapCtx *direct.MapContext, in *pb.LinkedRouterApplianceInstances) *krm.LinkedRouterApplianceInstances {
	if in == nil {
		return nil
	}
	out := &krm.LinkedRouterApplianceInstances{}
	out.Instances = direct.Slice_FromProto(mapCtx, in.Instances, RouterApplianceInstance_FromProto)
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	// MISSING: VpcNetwork
	out.IncludeImportRanges = in.IncludeImportRanges
	return out
}
func LinkedRouterApplianceInstances_ToProto(mapCtx *direct.MapContext, in *krm.LinkedRouterApplianceInstances) *pb.LinkedRouterApplianceInstances {
	if in == nil {
		return nil
	}
	out := &pb.LinkedRouterApplianceInstances{}
	out.Instances = direct.Slice_ToProto(mapCtx, in.Instances, RouterApplianceInstance_ToProto)
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	// MISSING: VpcNetwork
	out.IncludeImportRanges = in.IncludeImportRanges
	return out
}
func LinkedRouterApplianceInstancesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LinkedRouterApplianceInstances) *krm.LinkedRouterApplianceInstancesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LinkedRouterApplianceInstancesObservedState{}
	// MISSING: Instances
	// MISSING: SiteToSiteDataTransfer
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
	// MISSING: IncludeImportRanges
	return out
}
func LinkedRouterApplianceInstancesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LinkedRouterApplianceInstancesObservedState) *pb.LinkedRouterApplianceInstances {
	if in == nil {
		return nil
	}
	out := &pb.LinkedRouterApplianceInstances{}
	// MISSING: Instances
	// MISSING: SiteToSiteDataTransfer
	out.VpcNetwork = direct.ValueOf(in.VpcNetwork)
	// MISSING: IncludeImportRanges
	return out
}
func LinkedVpcNetwork_FromProto(mapCtx *direct.MapContext, in *pb.LinkedVpcNetwork) *krm.LinkedVpcNetwork {
	if in == nil {
		return nil
	}
	out := &krm.LinkedVpcNetwork{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.ExcludeExportRanges = in.ExcludeExportRanges
	out.IncludeExportRanges = in.IncludeExportRanges
	// MISSING: ProducerVpcSpokes
	return out
}
func LinkedVpcNetwork_ToProto(mapCtx *direct.MapContext, in *krm.LinkedVpcNetwork) *pb.LinkedVpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.LinkedVpcNetwork{}
	out.Uri = direct.ValueOf(in.URI)
	out.ExcludeExportRanges = in.ExcludeExportRanges
	out.IncludeExportRanges = in.IncludeExportRanges
	// MISSING: ProducerVpcSpokes
	return out
}
func LinkedVpcNetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LinkedVpcNetwork) *krm.LinkedVpcNetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LinkedVpcNetworkObservedState{}
	// MISSING: URI
	// MISSING: ExcludeExportRanges
	// MISSING: IncludeExportRanges
	out.ProducerVpcSpokes = in.ProducerVpcSpokes
	return out
}
func LinkedVpcNetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LinkedVpcNetworkObservedState) *pb.LinkedVpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.LinkedVpcNetwork{}
	// MISSING: URI
	// MISSING: ExcludeExportRanges
	// MISSING: IncludeExportRanges
	out.ProducerVpcSpokes = in.ProducerVpcSpokes
	return out
}
func LinkedVpnTunnels_FromProto(mapCtx *direct.MapContext, in *pb.LinkedVpnTunnels) *krm.LinkedVpnTunnels {
	if in == nil {
		return nil
	}
	out := &krm.LinkedVpnTunnels{}
	out.Uris = in.Uris
	out.SiteToSiteDataTransfer = direct.LazyPtr(in.GetSiteToSiteDataTransfer())
	// MISSING: VpcNetwork
	out.IncludeImportRanges = in.IncludeImportRanges
	return out
}
func LinkedVpnTunnels_ToProto(mapCtx *direct.MapContext, in *krm.LinkedVpnTunnels) *pb.LinkedVpnTunnels {
	if in == nil {
		return nil
	}
	out := &pb.LinkedVpnTunnels{}
	out.Uris = in.Uris
	out.SiteToSiteDataTransfer = direct.ValueOf(in.SiteToSiteDataTransfer)
	// MISSING: VpcNetwork
	out.IncludeImportRanges = in.IncludeImportRanges
	return out
}
func LinkedVpnTunnelsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LinkedVpnTunnels) *krm.LinkedVpnTunnelsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LinkedVpnTunnelsObservedState{}
	// MISSING: Uris
	// MISSING: SiteToSiteDataTransfer
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
	// MISSING: IncludeImportRanges
	return out
}
func LinkedVpnTunnelsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LinkedVpnTunnelsObservedState) *pb.LinkedVpnTunnels {
	if in == nil {
		return nil
	}
	out := &pb.LinkedVpnTunnels{}
	// MISSING: Uris
	// MISSING: SiteToSiteDataTransfer
	out.VpcNetwork = direct.ValueOf(in.VpcNetwork)
	// MISSING: IncludeImportRanges
	return out
}
func NetworkconnectivitySpokeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Spoke) *krm.NetworkconnectivitySpokeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkconnectivitySpokeObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Hub
	// MISSING: Group
	// MISSING: LinkedVpnTunnels
	// MISSING: LinkedInterconnectAttachments
	// MISSING: LinkedRouterApplianceInstances
	// MISSING: LinkedVpcNetwork
	// MISSING: LinkedProducerVpcNetwork
	// MISSING: UniqueID
	// MISSING: State
	// MISSING: Reasons
	// MISSING: SpokeType
	return out
}
func NetworkconnectivitySpokeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkconnectivitySpokeObservedState) *pb.Spoke {
	if in == nil {
		return nil
	}
	out := &pb.Spoke{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Hub
	// MISSING: Group
	// MISSING: LinkedVpnTunnels
	// MISSING: LinkedInterconnectAttachments
	// MISSING: LinkedRouterApplianceInstances
	// MISSING: LinkedVpcNetwork
	// MISSING: LinkedProducerVpcNetwork
	// MISSING: UniqueID
	// MISSING: State
	// MISSING: Reasons
	// MISSING: SpokeType
	return out
}
func NetworkconnectivitySpokeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Spoke) *krm.NetworkconnectivitySpokeSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkconnectivitySpokeSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Hub
	// MISSING: Group
	// MISSING: LinkedVpnTunnels
	// MISSING: LinkedInterconnectAttachments
	// MISSING: LinkedRouterApplianceInstances
	// MISSING: LinkedVpcNetwork
	// MISSING: LinkedProducerVpcNetwork
	// MISSING: UniqueID
	// MISSING: State
	// MISSING: Reasons
	// MISSING: SpokeType
	return out
}
func NetworkconnectivitySpokeSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkconnectivitySpokeSpec) *pb.Spoke {
	if in == nil {
		return nil
	}
	out := &pb.Spoke{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Hub
	// MISSING: Group
	// MISSING: LinkedVpnTunnels
	// MISSING: LinkedInterconnectAttachments
	// MISSING: LinkedRouterApplianceInstances
	// MISSING: LinkedVpcNetwork
	// MISSING: LinkedProducerVpcNetwork
	// MISSING: UniqueID
	// MISSING: State
	// MISSING: Reasons
	// MISSING: SpokeType
	return out
}
func RouterApplianceInstance_FromProto(mapCtx *direct.MapContext, in *pb.RouterApplianceInstance) *krm.RouterApplianceInstance {
	if in == nil {
		return nil
	}
	out := &krm.RouterApplianceInstance{}
	out.VirtualMachine = direct.LazyPtr(in.GetVirtualMachine())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	return out
}
func RouterApplianceInstance_ToProto(mapCtx *direct.MapContext, in *krm.RouterApplianceInstance) *pb.RouterApplianceInstance {
	if in == nil {
		return nil
	}
	out := &pb.RouterApplianceInstance{}
	out.VirtualMachine = direct.ValueOf(in.VirtualMachine)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	return out
}
func Spoke_FromProto(mapCtx *direct.MapContext, in *pb.Spoke) *krm.Spoke {
	if in == nil {
		return nil
	}
	out := &krm.Spoke{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Hub = direct.LazyPtr(in.GetHub())
	out.Group = direct.LazyPtr(in.GetGroup())
	out.LinkedVpnTunnels = LinkedVpnTunnels_FromProto(mapCtx, in.GetLinkedVpnTunnels())
	out.LinkedInterconnectAttachments = LinkedInterconnectAttachments_FromProto(mapCtx, in.GetLinkedInterconnectAttachments())
	out.LinkedRouterApplianceInstances = LinkedRouterApplianceInstances_FromProto(mapCtx, in.GetLinkedRouterApplianceInstances())
	out.LinkedVpcNetwork = LinkedVpcNetwork_FromProto(mapCtx, in.GetLinkedVpcNetwork())
	out.LinkedProducerVpcNetwork = LinkedProducerVpcNetwork_FromProto(mapCtx, in.GetLinkedProducerVpcNetwork())
	// MISSING: UniqueID
	// MISSING: State
	// MISSING: Reasons
	// MISSING: SpokeType
	return out
}
func Spoke_ToProto(mapCtx *direct.MapContext, in *krm.Spoke) *pb.Spoke {
	if in == nil {
		return nil
	}
	out := &pb.Spoke{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Hub = direct.ValueOf(in.Hub)
	out.Group = direct.ValueOf(in.Group)
	out.LinkedVpnTunnels = LinkedVpnTunnels_ToProto(mapCtx, in.LinkedVpnTunnels)
	out.LinkedInterconnectAttachments = LinkedInterconnectAttachments_ToProto(mapCtx, in.LinkedInterconnectAttachments)
	out.LinkedRouterApplianceInstances = LinkedRouterApplianceInstances_ToProto(mapCtx, in.LinkedRouterApplianceInstances)
	out.LinkedVpcNetwork = LinkedVpcNetwork_ToProto(mapCtx, in.LinkedVpcNetwork)
	out.LinkedProducerVpcNetwork = LinkedProducerVpcNetwork_ToProto(mapCtx, in.LinkedProducerVpcNetwork)
	// MISSING: UniqueID
	// MISSING: State
	// MISSING: Reasons
	// MISSING: SpokeType
	return out
}
func SpokeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Spoke) *krm.SpokeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpokeObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Hub
	// MISSING: Group
	out.LinkedVpnTunnels = LinkedVpnTunnelsObservedState_FromProto(mapCtx, in.GetLinkedVpnTunnels())
	out.LinkedInterconnectAttachments = LinkedInterconnectAttachmentsObservedState_FromProto(mapCtx, in.GetLinkedInterconnectAttachments())
	out.LinkedRouterApplianceInstances = LinkedRouterApplianceInstancesObservedState_FromProto(mapCtx, in.GetLinkedRouterApplianceInstances())
	out.LinkedVpcNetwork = LinkedVpcNetworkObservedState_FromProto(mapCtx, in.GetLinkedVpcNetwork())
	out.LinkedProducerVpcNetwork = LinkedProducerVpcNetworkObservedState_FromProto(mapCtx, in.GetLinkedProducerVpcNetwork())
	out.UniqueID = direct.LazyPtr(in.GetUniqueId())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Reasons = direct.Slice_FromProto(mapCtx, in.Reasons, Spoke_StateReason_FromProto)
	out.SpokeType = direct.Enum_FromProto(mapCtx, in.GetSpokeType())
	return out
}
func SpokeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpokeObservedState) *pb.Spoke {
	if in == nil {
		return nil
	}
	out := &pb.Spoke{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Hub
	// MISSING: Group
	out.LinkedVpnTunnels = LinkedVpnTunnelsObservedState_ToProto(mapCtx, in.LinkedVpnTunnels)
	out.LinkedInterconnectAttachments = LinkedInterconnectAttachmentsObservedState_ToProto(mapCtx, in.LinkedInterconnectAttachments)
	out.LinkedRouterApplianceInstances = LinkedRouterApplianceInstancesObservedState_ToProto(mapCtx, in.LinkedRouterApplianceInstances)
	out.LinkedVpcNetwork = LinkedVpcNetworkObservedState_ToProto(mapCtx, in.LinkedVpcNetwork)
	out.LinkedProducerVpcNetwork = LinkedProducerVpcNetworkObservedState_ToProto(mapCtx, in.LinkedProducerVpcNetwork)
	out.UniqueId = direct.ValueOf(in.UniqueID)
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	out.Reasons = direct.Slice_ToProto(mapCtx, in.Reasons, Spoke_StateReason_ToProto)
	out.SpokeType = direct.Enum_ToProto[pb.SpokeType](mapCtx, in.SpokeType)
	return out
}
func Spoke_StateReason_FromProto(mapCtx *direct.MapContext, in *pb.Spoke_StateReason) *krm.Spoke_StateReason {
	if in == nil {
		return nil
	}
	out := &krm.Spoke_StateReason{}
	out.Code = direct.Enum_FromProto(mapCtx, in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.UserDetails = direct.LazyPtr(in.GetUserDetails())
	return out
}
func Spoke_StateReason_ToProto(mapCtx *direct.MapContext, in *krm.Spoke_StateReason) *pb.Spoke_StateReason {
	if in == nil {
		return nil
	}
	out := &pb.Spoke_StateReason{}
	out.Code = direct.Enum_ToProto[pb.Spoke_StateReason_Code](mapCtx, in.Code)
	out.Message = direct.ValueOf(in.Message)
	out.UserDetails = direct.ValueOf(in.UserDetails)
	return out
}
