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
	networkconnectivitypb "cloud.google.com/go/networkconnectivity/apiv1/networkconnectivitypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Destination_DestinationEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.Destination_DestinationEndpoint) *krm.Destination_DestinationEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Destination_DestinationEndpointObservedState{}
	// MISSING: Asn
	// MISSING: Csp
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func Destination_DestinationEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Destination_DestinationEndpointObservedState) *networkconnectivitypb.Destination_DestinationEndpoint {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.Destination_DestinationEndpoint{}
	// MISSING: Asn
	// MISSING: Csp
	out.State = direct.Enum_ToProto[networkconnectivitypb.Destination_DestinationEndpoint_State](mapCtx, in.State)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func Destination_DestinationEndpoint_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.Destination_DestinationEndpoint) *krm.Destination_DestinationEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.Destination_DestinationEndpoint{}
	out.Asn = direct.LazyPtr(in.GetAsn())
	out.Csp = direct.LazyPtr(in.GetCsp())
	// MISSING: State
	// MISSING: UpdateTime
	return out
}

func Destination_DestinationEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.Destination_DestinationEndpoint) *networkconnectivitypb.Destination_DestinationEndpoint {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.Destination_DestinationEndpoint{}
	out.Asn = direct.ValueOf(in.Asn)
	out.Csp = direct.ValueOf(in.Csp)
	// MISSING: State
	// MISSING: UpdateTime
	return out
}

func NetworkConnectivityDestinationObservedState_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.Destination) *krm.NetworkConnectivityDestinationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityDestinationObservedState{}
	// MISSING: Name
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Endpoints = direct.Slice_FromProto(mapCtx, in.Endpoints, Destination_DestinationEndpointObservedState_FromProto)
	out.StateTimeline = StateTimelineObservedState_FromProto(mapCtx, in.GetStateTimeline())
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}

func NetworkConnectivityDestinationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityDestinationObservedState) *networkconnectivitypb.Destination {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.Destination{}
	// MISSING: Name
	out.Etag = direct.ValueOf(in.Etag)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Endpoints = direct.Slice_ToProto(mapCtx, in.Endpoints, Destination_DestinationEndpointObservedState_ToProto)
	out.StateTimeline = StateTimelineObservedState_ToProto(mapCtx, in.StateTimeline)
	out.Uid = direct.ValueOf(in.Uid)
	return out
}

func NetworkConnectivityDestinationSpec_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.Destination) *krm.NetworkConnectivityDestinationSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityDestinationSpec{}
	// MISSING: Name
	// MISSING: ParentRef
	out.Description = direct.LazyPtr(in.GetDescription())
	out.IPPrefix = direct.LazyPtr(in.GetIpPrefix())
	out.Endpoints = direct.Slice_FromProto(mapCtx, in.Endpoints, Destination_DestinationEndpoint_FromProto)
	return out
}

func NetworkConnectivityDestinationSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityDestinationSpec) *networkconnectivitypb.Destination {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.Destination{}
	// MISSING: Name
	// MISSING: ParentRef
	out.Description = direct.ValueOf(in.Description)
	out.IpPrefix = direct.ValueOf(in.IPPrefix)
	out.Endpoints = direct.Slice_ToProto(mapCtx, in.Endpoints, Destination_DestinationEndpoint_ToProto)
	return out
}

func NetworkConnectivityPolicyBasedRouteObservedState_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.PolicyBasedRoute) *krm.NetworkConnectivityPolicyBasedRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityPolicyBasedRouteObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Warnings = direct.Slice_FromProto(mapCtx, in.Warnings, PolicyBasedRoute_WarningsObservedState_FromProto)
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.Kind = direct.LazyPtr(in.GetKind())
	return out
}

func NetworkConnectivityPolicyBasedRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityPolicyBasedRouteObservedState) *networkconnectivitypb.PolicyBasedRoute {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.PolicyBasedRoute{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Warnings = direct.Slice_ToProto(mapCtx, in.Warnings, PolicyBasedRoute_WarningsObservedState_ToProto)
	out.SelfLink = direct.ValueOf(in.SelfLink)
	out.Kind = direct.ValueOf(in.Kind)
	return out
}

func NetworkConnectivityPolicyBasedRouteSpec_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.PolicyBasedRoute) *krm.NetworkConnectivityPolicyBasedRouteSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityPolicyBasedRouteSpec{}
	out.VirtualMachine = PolicyBasedRoute_VirtualMachine_FromProto(mapCtx, in.GetVirtualMachine())
	out.InterconnectAttachment = PolicyBasedRoute_InterconnectAttachment_FromProto(mapCtx, in.GetInterconnectAttachment())
	out.NextHopIlbIP = direct.LazyPtr(in.GetNextHopIlbIp())
	out.NextHopOtherRoutes = direct.Enum_FromProto(mapCtx, in.GetNextHopOtherRoutes())
	// MISSING: Name
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Filter = PolicyBasedRoute_Filter_FromProto(mapCtx, in.GetFilter())
	out.Priority = direct.LazyPtr(in.GetPriority())
	return out
}

func NetworkConnectivityPolicyBasedRouteSpec_NextHopIlbIp_ToProto(mapCtx *direct.MapContext, in *string) *networkconnectivitypb.PolicyBasedRoute_NextHopIlbIp {
	if in == nil {
		return nil
	}
	return &networkconnectivitypb.PolicyBasedRoute_NextHopIlbIp{NextHopIlbIp: *in}
}

func NetworkConnectivityPolicyBasedRouteSpec_NextHopOtherRoutes_ToProto(mapCtx *direct.MapContext, in *string) *networkconnectivitypb.PolicyBasedRoute_NextHopOtherRoutes {
	if in == nil {
		return nil
	}
	return &networkconnectivitypb.PolicyBasedRoute_NextHopOtherRoutes{NextHopOtherRoutes: direct.Enum_ToProto[networkconnectivitypb.PolicyBasedRoute_OtherRoutes](mapCtx, in)}
}

func NetworkConnectivityPolicyBasedRouteSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityPolicyBasedRouteSpec) *networkconnectivitypb.PolicyBasedRoute {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.PolicyBasedRoute{}
	if oneof := PolicyBasedRoute_VirtualMachine_ToProto(mapCtx, in.VirtualMachine); oneof != nil {
		out.Target = &networkconnectivitypb.PolicyBasedRoute_VirtualMachine_{VirtualMachine: oneof}
	}
	if oneof := PolicyBasedRoute_InterconnectAttachment_ToProto(mapCtx, in.InterconnectAttachment); oneof != nil {
		out.Target = &networkconnectivitypb.PolicyBasedRoute_InterconnectAttachment_{InterconnectAttachment: oneof}
	}
	if oneof := NetworkConnectivityPolicyBasedRouteSpec_NextHopIlbIp_ToProto(mapCtx, in.NextHopIlbIP); oneof != nil {
		out.NextHop = oneof
	}
	if oneof := NetworkConnectivityPolicyBasedRouteSpec_NextHopOtherRoutes_ToProto(mapCtx, in.NextHopOtherRoutes); oneof != nil {
		out.NextHop = oneof
	}
	// MISSING: Name
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Network = direct.ValueOf(in.Network)
	out.Filter = PolicyBasedRoute_Filter_ToProto(mapCtx, in.Filter)
	out.Priority = direct.ValueOf(in.Priority)
	return out
}

func NetworkConnectivityServiceConnectionTokenObservedState_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.ServiceConnectionToken) *krm.NetworkConnectivityServiceConnectionTokenObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityServiceConnectionTokenObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Token = direct.LazyPtr(in.GetToken())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	return out
}

func NetworkConnectivityServiceConnectionTokenObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityServiceConnectionTokenObservedState) *networkconnectivitypb.ServiceConnectionToken {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.ServiceConnectionToken{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Token = direct.ValueOf(in.Token)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	return out
}

func NetworkConnectivityServiceConnectionTokenSpec_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.ServiceConnectionToken) *krm.NetworkConnectivityServiceConnectionTokenSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityServiceConnectionTokenSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Etag = in.Etag
	return out
}

func NetworkConnectivityServiceConnectionTokenSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityServiceConnectionTokenSpec) *networkconnectivitypb.ServiceConnectionToken {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.ServiceConnectionToken{}
	// MISSING: Name
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Network = direct.ValueOf(in.Network)
	out.Etag = in.Etag
	return out
}

func PolicyBasedRoute_Filter_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.PolicyBasedRoute_Filter) *krm.PolicyBasedRoute_Filter {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBasedRoute_Filter{}
	out.IPProtocol = direct.LazyPtr(in.GetIpProtocol())
	out.SrcRange = direct.LazyPtr(in.GetSrcRange())
	out.DestRange = direct.LazyPtr(in.GetDestRange())
	out.ProtocolVersion = direct.Enum_FromProto(mapCtx, in.GetProtocolVersion())
	return out
}

func PolicyBasedRoute_Filter_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBasedRoute_Filter) *networkconnectivitypb.PolicyBasedRoute_Filter {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.PolicyBasedRoute_Filter{}
	out.IpProtocol = direct.ValueOf(in.IPProtocol)
	out.SrcRange = direct.ValueOf(in.SrcRange)
	out.DestRange = direct.ValueOf(in.DestRange)
	out.ProtocolVersion = direct.Enum_ToProto[networkconnectivitypb.PolicyBasedRoute_Filter_ProtocolVersion](mapCtx, in.ProtocolVersion)
	return out
}

func PolicyBasedRoute_InterconnectAttachment_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.PolicyBasedRoute_InterconnectAttachment) *krm.PolicyBasedRoute_InterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBasedRoute_InterconnectAttachment{}
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}

func PolicyBasedRoute_InterconnectAttachment_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBasedRoute_InterconnectAttachment) *networkconnectivitypb.PolicyBasedRoute_InterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.PolicyBasedRoute_InterconnectAttachment{}
	out.Region = direct.ValueOf(in.Region)
	return out
}

func PolicyBasedRoute_VirtualMachine_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.PolicyBasedRoute_VirtualMachine) *krm.PolicyBasedRoute_VirtualMachine {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBasedRoute_VirtualMachine{}
	out.Tags = in.Tags
	return out
}

func PolicyBasedRoute_VirtualMachine_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBasedRoute_VirtualMachine) *networkconnectivitypb.PolicyBasedRoute_VirtualMachine {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.PolicyBasedRoute_VirtualMachine{}
	out.Tags = in.Tags
	return out
}

func PolicyBasedRoute_WarningsObservedState_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.PolicyBasedRoute_Warnings) *krm.PolicyBasedRoute_WarningsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBasedRoute_WarningsObservedState{}
	out.Code = direct.Enum_FromProto(mapCtx, in.GetCode())
	out.Data = in.Data
	out.WarningMessage = direct.LazyPtr(in.GetWarningMessage())
	return out
}

func PolicyBasedRoute_WarningsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBasedRoute_WarningsObservedState) *networkconnectivitypb.PolicyBasedRoute_Warnings {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.PolicyBasedRoute_Warnings{}
	out.Code = direct.Enum_ToProto[networkconnectivitypb.PolicyBasedRoute_Warnings_Code](mapCtx, in.Code)
	out.Data = in.Data
	out.WarningMessage = direct.ValueOf(in.WarningMessage)
	return out
}

func StateTimelineObservedState_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.StateTimeline) *krm.StateTimelineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StateTimelineObservedState{}
	out.States = direct.Slice_FromProto(mapCtx, in.States, StateTimeline_StateMetadataObservedState_FromProto)
	return out
}

func StateTimelineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StateTimelineObservedState) *networkconnectivitypb.StateTimeline {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.StateTimeline{}
	out.States = direct.Slice_ToProto(mapCtx, in.States, StateTimeline_StateMetadataObservedState_ToProto)
	return out
}

func StateTimeline_StateMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *networkconnectivitypb.StateTimeline_StateMetadata) *krm.StateTimeline_StateMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StateTimeline_StateMetadataObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.EffectiveTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEffectiveTime())
	return out
}

func StateTimeline_StateMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StateTimeline_StateMetadataObservedState) *networkconnectivitypb.StateTimeline_StateMetadata {
	if in == nil {
		return nil
	}
	out := &networkconnectivitypb.StateTimeline_StateMetadata{}
	out.State = direct.Enum_ToProto[networkconnectivitypb.StateTimeline_StateMetadata_State](mapCtx, in.State)
	out.EffectiveTime = direct.StringTimestamp_ToProto(mapCtx, in.EffectiveTime)
	return out
}
