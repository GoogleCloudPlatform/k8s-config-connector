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
func Hub_FromProto(mapCtx *direct.MapContext, in *pb.Hub) *krm.Hub {
	if in == nil {
		return nil
	}
	out := &krm.Hub{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: UniqueID
	// MISSING: State
	out.RoutingVpcs = direct.Slice_FromProto(mapCtx, in.RoutingVpcs, RoutingVPC_FromProto)
	// MISSING: RouteTables
	// MISSING: SpokeSummary
	out.PolicyMode = direct.Enum_FromProto(mapCtx, in.GetPolicyMode())
	out.PresetTopology = direct.Enum_FromProto(mapCtx, in.GetPresetTopology())
	out.ExportPsc = in.ExportPsc
	return out
}
func Hub_ToProto(mapCtx *direct.MapContext, in *krm.Hub) *pb.Hub {
	if in == nil {
		return nil
	}
	out := &pb.Hub{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	// MISSING: UniqueID
	// MISSING: State
	out.RoutingVpcs = direct.Slice_ToProto(mapCtx, in.RoutingVpcs, RoutingVPC_ToProto)
	// MISSING: RouteTables
	// MISSING: SpokeSummary
	out.PolicyMode = direct.Enum_ToProto[pb.PolicyMode](mapCtx, in.PolicyMode)
	out.PresetTopology = direct.Enum_ToProto[pb.PresetTopology](mapCtx, in.PresetTopology)
	out.ExportPsc = in.ExportPsc
	return out
}
func HubObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Hub) *krm.HubObservedState {
	if in == nil {
		return nil
	}
	out := &krm.HubObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	out.UniqueID = direct.LazyPtr(in.GetUniqueId())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.RoutingVpcs = direct.Slice_FromProto(mapCtx, in.RoutingVpcs, RoutingVPCObservedState_FromProto)
	out.RouteTables = in.RouteTables
	out.SpokeSummary = SpokeSummary_FromProto(mapCtx, in.GetSpokeSummary())
	// MISSING: PolicyMode
	// MISSING: PresetTopology
	// MISSING: ExportPsc
	return out
}
func HubObservedState_ToProto(mapCtx *direct.MapContext, in *krm.HubObservedState) *pb.Hub {
	if in == nil {
		return nil
	}
	out := &pb.Hub{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	out.UniqueId = direct.ValueOf(in.UniqueID)
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	out.RoutingVpcs = direct.Slice_ToProto(mapCtx, in.RoutingVpcs, RoutingVPCObservedState_ToProto)
	out.RouteTables = in.RouteTables
	out.SpokeSummary = SpokeSummary_ToProto(mapCtx, in.SpokeSummary)
	// MISSING: PolicyMode
	// MISSING: PresetTopology
	// MISSING: ExportPsc
	return out
}
func NetworkconnectivityHubObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Hub) *krm.NetworkconnectivityHubObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkconnectivityHubObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: UniqueID
	// MISSING: State
	// MISSING: RoutingVpcs
	// MISSING: RouteTables
	// MISSING: SpokeSummary
	// MISSING: PolicyMode
	// MISSING: PresetTopology
	// MISSING: ExportPsc
	return out
}
func NetworkconnectivityHubObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkconnectivityHubObservedState) *pb.Hub {
	if in == nil {
		return nil
	}
	out := &pb.Hub{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: UniqueID
	// MISSING: State
	// MISSING: RoutingVpcs
	// MISSING: RouteTables
	// MISSING: SpokeSummary
	// MISSING: PolicyMode
	// MISSING: PresetTopology
	// MISSING: ExportPsc
	return out
}
func NetworkconnectivityHubSpec_FromProto(mapCtx *direct.MapContext, in *pb.Hub) *krm.NetworkconnectivityHubSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkconnectivityHubSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: UniqueID
	// MISSING: State
	// MISSING: RoutingVpcs
	// MISSING: RouteTables
	// MISSING: SpokeSummary
	// MISSING: PolicyMode
	// MISSING: PresetTopology
	// MISSING: ExportPsc
	return out
}
func NetworkconnectivityHubSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkconnectivityHubSpec) *pb.Hub {
	if in == nil {
		return nil
	}
	out := &pb.Hub{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: UniqueID
	// MISSING: State
	// MISSING: RoutingVpcs
	// MISSING: RouteTables
	// MISSING: SpokeSummary
	// MISSING: PolicyMode
	// MISSING: PresetTopology
	// MISSING: ExportPsc
	return out
}
func RoutingVPC_FromProto(mapCtx *direct.MapContext, in *pb.RoutingVPC) *krm.RoutingVPC {
	if in == nil {
		return nil
	}
	out := &krm.RoutingVPC{}
	out.URI = direct.LazyPtr(in.GetUri())
	// MISSING: RequiredForNewSiteToSiteDataTransferSpokes
	return out
}
func RoutingVPC_ToProto(mapCtx *direct.MapContext, in *krm.RoutingVPC) *pb.RoutingVPC {
	if in == nil {
		return nil
	}
	out := &pb.RoutingVPC{}
	out.Uri = direct.ValueOf(in.URI)
	// MISSING: RequiredForNewSiteToSiteDataTransferSpokes
	return out
}
func RoutingVPCObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RoutingVPC) *krm.RoutingVPCObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RoutingVPCObservedState{}
	// MISSING: URI
	out.RequiredForNewSiteToSiteDataTransferSpokes = direct.LazyPtr(in.GetRequiredForNewSiteToSiteDataTransferSpokes())
	return out
}
func RoutingVPCObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RoutingVPCObservedState) *pb.RoutingVPC {
	if in == nil {
		return nil
	}
	out := &pb.RoutingVPC{}
	// MISSING: URI
	out.RequiredForNewSiteToSiteDataTransferSpokes = direct.ValueOf(in.RequiredForNewSiteToSiteDataTransferSpokes)
	return out
}
func SpokeSummary_FromProto(mapCtx *direct.MapContext, in *pb.SpokeSummary) *krm.SpokeSummary {
	if in == nil {
		return nil
	}
	out := &krm.SpokeSummary{}
	// MISSING: SpokeTypeCounts
	// MISSING: SpokeStateCounts
	// MISSING: SpokeStateReasonCounts
	return out
}
func SpokeSummary_ToProto(mapCtx *direct.MapContext, in *krm.SpokeSummary) *pb.SpokeSummary {
	if in == nil {
		return nil
	}
	out := &pb.SpokeSummary{}
	// MISSING: SpokeTypeCounts
	// MISSING: SpokeStateCounts
	// MISSING: SpokeStateReasonCounts
	return out
}
func SpokeSummaryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SpokeSummary) *krm.SpokeSummaryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpokeSummaryObservedState{}
	out.SpokeTypeCounts = direct.Slice_FromProto(mapCtx, in.SpokeTypeCounts, SpokeSummary_SpokeTypeCount_FromProto)
	out.SpokeStateCounts = direct.Slice_FromProto(mapCtx, in.SpokeStateCounts, SpokeSummary_SpokeStateCount_FromProto)
	out.SpokeStateReasonCounts = direct.Slice_FromProto(mapCtx, in.SpokeStateReasonCounts, SpokeSummary_SpokeStateReasonCount_FromProto)
	return out
}
func SpokeSummaryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpokeSummaryObservedState) *pb.SpokeSummary {
	if in == nil {
		return nil
	}
	out := &pb.SpokeSummary{}
	out.SpokeTypeCounts = direct.Slice_ToProto(mapCtx, in.SpokeTypeCounts, SpokeSummary_SpokeTypeCount_ToProto)
	out.SpokeStateCounts = direct.Slice_ToProto(mapCtx, in.SpokeStateCounts, SpokeSummary_SpokeStateCount_ToProto)
	out.SpokeStateReasonCounts = direct.Slice_ToProto(mapCtx, in.SpokeStateReasonCounts, SpokeSummary_SpokeStateReasonCount_ToProto)
	return out
}
func SpokeSummary_SpokeStateCount_FromProto(mapCtx *direct.MapContext, in *pb.SpokeSummary_SpokeStateCount) *krm.SpokeSummary_SpokeStateCount {
	if in == nil {
		return nil
	}
	out := &krm.SpokeSummary_SpokeStateCount{}
	// MISSING: State
	// MISSING: Count
	return out
}
func SpokeSummary_SpokeStateCount_ToProto(mapCtx *direct.MapContext, in *krm.SpokeSummary_SpokeStateCount) *pb.SpokeSummary_SpokeStateCount {
	if in == nil {
		return nil
	}
	out := &pb.SpokeSummary_SpokeStateCount{}
	// MISSING: State
	// MISSING: Count
	return out
}
func SpokeSummary_SpokeStateCountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SpokeSummary_SpokeStateCount) *krm.SpokeSummary_SpokeStateCountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpokeSummary_SpokeStateCountObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func SpokeSummary_SpokeStateCountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpokeSummary_SpokeStateCountObservedState) *pb.SpokeSummary_SpokeStateCount {
	if in == nil {
		return nil
	}
	out := &pb.SpokeSummary_SpokeStateCount{}
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	out.Count = direct.ValueOf(in.Count)
	return out
}
func SpokeSummary_SpokeStateReasonCount_FromProto(mapCtx *direct.MapContext, in *pb.SpokeSummary_SpokeStateReasonCount) *krm.SpokeSummary_SpokeStateReasonCount {
	if in == nil {
		return nil
	}
	out := &krm.SpokeSummary_SpokeStateReasonCount{}
	// MISSING: StateReasonCode
	// MISSING: Count
	return out
}
func SpokeSummary_SpokeStateReasonCount_ToProto(mapCtx *direct.MapContext, in *krm.SpokeSummary_SpokeStateReasonCount) *pb.SpokeSummary_SpokeStateReasonCount {
	if in == nil {
		return nil
	}
	out := &pb.SpokeSummary_SpokeStateReasonCount{}
	// MISSING: StateReasonCode
	// MISSING: Count
	return out
}
func SpokeSummary_SpokeStateReasonCountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SpokeSummary_SpokeStateReasonCount) *krm.SpokeSummary_SpokeStateReasonCountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpokeSummary_SpokeStateReasonCountObservedState{}
	out.StateReasonCode = direct.Enum_FromProto(mapCtx, in.GetStateReasonCode())
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func SpokeSummary_SpokeStateReasonCountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpokeSummary_SpokeStateReasonCountObservedState) *pb.SpokeSummary_SpokeStateReasonCount {
	if in == nil {
		return nil
	}
	out := &pb.SpokeSummary_SpokeStateReasonCount{}
	out.StateReasonCode = direct.Enum_ToProto[pb.Spoke_StateReason_Code](mapCtx, in.StateReasonCode)
	out.Count = direct.ValueOf(in.Count)
	return out
}
func SpokeSummary_SpokeTypeCount_FromProto(mapCtx *direct.MapContext, in *pb.SpokeSummary_SpokeTypeCount) *krm.SpokeSummary_SpokeTypeCount {
	if in == nil {
		return nil
	}
	out := &krm.SpokeSummary_SpokeTypeCount{}
	// MISSING: SpokeType
	// MISSING: Count
	return out
}
func SpokeSummary_SpokeTypeCount_ToProto(mapCtx *direct.MapContext, in *krm.SpokeSummary_SpokeTypeCount) *pb.SpokeSummary_SpokeTypeCount {
	if in == nil {
		return nil
	}
	out := &pb.SpokeSummary_SpokeTypeCount{}
	// MISSING: SpokeType
	// MISSING: Count
	return out
}
func SpokeSummary_SpokeTypeCountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SpokeSummary_SpokeTypeCount) *krm.SpokeSummary_SpokeTypeCountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpokeSummary_SpokeTypeCountObservedState{}
	out.SpokeType = direct.Enum_FromProto(mapCtx, in.GetSpokeType())
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func SpokeSummary_SpokeTypeCountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpokeSummary_SpokeTypeCountObservedState) *pb.SpokeSummary_SpokeTypeCount {
	if in == nil {
		return nil
	}
	out := &pb.SpokeSummary_SpokeTypeCount{}
	out.SpokeType = direct.Enum_ToProto[pb.SpokeType](mapCtx, in.SpokeType)
	out.Count = direct.ValueOf(in.Count)
	return out
}
