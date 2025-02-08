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
func NetworkconnectivityPolicyBasedRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBasedRoute) *krm.NetworkconnectivityPolicyBasedRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkconnectivityPolicyBasedRouteObservedState{}
	// MISSING: VirtualMachine
	// MISSING: InterconnectAttachment
	// MISSING: NextHopIlbIP
	// MISSING: NextHopOtherRoutes
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Filter
	// MISSING: Priority
	// MISSING: Warnings
	// MISSING: SelfLink
	// MISSING: Kind
	return out
}
func NetworkconnectivityPolicyBasedRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkconnectivityPolicyBasedRouteObservedState) *pb.PolicyBasedRoute {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBasedRoute{}
	// MISSING: VirtualMachine
	// MISSING: InterconnectAttachment
	// MISSING: NextHopIlbIP
	// MISSING: NextHopOtherRoutes
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Filter
	// MISSING: Priority
	// MISSING: Warnings
	// MISSING: SelfLink
	// MISSING: Kind
	return out
}
func NetworkconnectivityPolicyBasedRouteSpec_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBasedRoute) *krm.NetworkconnectivityPolicyBasedRouteSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkconnectivityPolicyBasedRouteSpec{}
	// MISSING: VirtualMachine
	// MISSING: InterconnectAttachment
	// MISSING: NextHopIlbIP
	// MISSING: NextHopOtherRoutes
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Filter
	// MISSING: Priority
	// MISSING: Warnings
	// MISSING: SelfLink
	// MISSING: Kind
	return out
}
func NetworkconnectivityPolicyBasedRouteSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkconnectivityPolicyBasedRouteSpec) *pb.PolicyBasedRoute {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBasedRoute{}
	// MISSING: VirtualMachine
	// MISSING: InterconnectAttachment
	// MISSING: NextHopIlbIP
	// MISSING: NextHopOtherRoutes
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Filter
	// MISSING: Priority
	// MISSING: Warnings
	// MISSING: SelfLink
	// MISSING: Kind
	return out
}
func PolicyBasedRoute_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBasedRoute) *krm.PolicyBasedRoute {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBasedRoute{}
	out.VirtualMachine = PolicyBasedRoute_VirtualMachine_FromProto(mapCtx, in.GetVirtualMachine())
	out.InterconnectAttachment = PolicyBasedRoute_InterconnectAttachment_FromProto(mapCtx, in.GetInterconnectAttachment())
	out.NextHopIlbIP = direct.LazyPtr(in.GetNextHopIlbIp())
	out.NextHopOtherRoutes = direct.Enum_FromProto(mapCtx, in.GetNextHopOtherRoutes())
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Filter = PolicyBasedRoute_Filter_FromProto(mapCtx, in.GetFilter())
	out.Priority = direct.LazyPtr(in.GetPriority())
	// MISSING: Warnings
	// MISSING: SelfLink
	// MISSING: Kind
	return out
}
func PolicyBasedRoute_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBasedRoute) *pb.PolicyBasedRoute {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBasedRoute{}
	if oneof := PolicyBasedRoute_VirtualMachine_ToProto(mapCtx, in.VirtualMachine); oneof != nil {
		out.Target = &pb.PolicyBasedRoute_VirtualMachine_{VirtualMachine: oneof}
	}
	if oneof := PolicyBasedRoute_InterconnectAttachment_ToProto(mapCtx, in.InterconnectAttachment); oneof != nil {
		out.Target = &pb.PolicyBasedRoute_InterconnectAttachment_{InterconnectAttachment: oneof}
	}
	if oneof := PolicyBasedRoute_NextHopIlbIp_ToProto(mapCtx, in.NextHopIlbIP); oneof != nil {
		out.NextHop = oneof
	}
	if oneof := PolicyBasedRoute_NextHopOtherRoutes_ToProto(mapCtx, in.NextHopOtherRoutes); oneof != nil {
		out.NextHop = oneof
	}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Network = direct.ValueOf(in.Network)
	out.Filter = PolicyBasedRoute_Filter_ToProto(mapCtx, in.Filter)
	out.Priority = direct.ValueOf(in.Priority)
	// MISSING: Warnings
	// MISSING: SelfLink
	// MISSING: Kind
	return out
}
func PolicyBasedRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBasedRoute) *krm.PolicyBasedRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBasedRouteObservedState{}
	// MISSING: VirtualMachine
	// MISSING: InterconnectAttachment
	// MISSING: NextHopIlbIP
	// MISSING: NextHopOtherRoutes
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Filter
	// MISSING: Priority
	out.Warnings = direct.Slice_FromProto(mapCtx, in.Warnings, PolicyBasedRoute_Warnings_FromProto)
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.Kind = direct.LazyPtr(in.GetKind())
	return out
}
func PolicyBasedRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBasedRouteObservedState) *pb.PolicyBasedRoute {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBasedRoute{}
	// MISSING: VirtualMachine
	// MISSING: InterconnectAttachment
	// MISSING: NextHopIlbIP
	// MISSING: NextHopOtherRoutes
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Network
	// MISSING: Filter
	// MISSING: Priority
	out.Warnings = direct.Slice_ToProto(mapCtx, in.Warnings, PolicyBasedRoute_Warnings_ToProto)
	out.SelfLink = direct.ValueOf(in.SelfLink)
	out.Kind = direct.ValueOf(in.Kind)
	return out
}
func PolicyBasedRoute_Filter_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBasedRoute_Filter) *krm.PolicyBasedRoute_Filter {
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
func PolicyBasedRoute_Filter_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBasedRoute_Filter) *pb.PolicyBasedRoute_Filter {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBasedRoute_Filter{}
	out.IpProtocol = direct.ValueOf(in.IPProtocol)
	out.SrcRange = direct.ValueOf(in.SrcRange)
	out.DestRange = direct.ValueOf(in.DestRange)
	out.ProtocolVersion = direct.Enum_ToProto[pb.PolicyBasedRoute_Filter_ProtocolVersion](mapCtx, in.ProtocolVersion)
	return out
}
func PolicyBasedRoute_InterconnectAttachment_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBasedRoute_InterconnectAttachment) *krm.PolicyBasedRoute_InterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBasedRoute_InterconnectAttachment{}
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func PolicyBasedRoute_InterconnectAttachment_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBasedRoute_InterconnectAttachment) *pb.PolicyBasedRoute_InterconnectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBasedRoute_InterconnectAttachment{}
	out.Region = direct.ValueOf(in.Region)
	return out
}
func PolicyBasedRoute_VirtualMachine_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBasedRoute_VirtualMachine) *krm.PolicyBasedRoute_VirtualMachine {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBasedRoute_VirtualMachine{}
	out.Tags = in.Tags
	return out
}
func PolicyBasedRoute_VirtualMachine_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBasedRoute_VirtualMachine) *pb.PolicyBasedRoute_VirtualMachine {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBasedRoute_VirtualMachine{}
	out.Tags = in.Tags
	return out
}
func PolicyBasedRoute_Warnings_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBasedRoute_Warnings) *krm.PolicyBasedRoute_Warnings {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBasedRoute_Warnings{}
	// MISSING: Code
	// MISSING: Data
	// MISSING: WarningMessage
	return out
}
func PolicyBasedRoute_Warnings_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBasedRoute_Warnings) *pb.PolicyBasedRoute_Warnings {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBasedRoute_Warnings{}
	// MISSING: Code
	// MISSING: Data
	// MISSING: WarningMessage
	return out
}
func PolicyBasedRoute_WarningsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBasedRoute_Warnings) *krm.PolicyBasedRoute_WarningsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBasedRoute_WarningsObservedState{}
	out.Code = direct.Enum_FromProto(mapCtx, in.GetCode())
	out.Data = in.Data
	out.WarningMessage = direct.LazyPtr(in.GetWarningMessage())
	return out
}
func PolicyBasedRoute_WarningsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBasedRoute_WarningsObservedState) *pb.PolicyBasedRoute_Warnings {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBasedRoute_Warnings{}
	out.Code = direct.Enum_ToProto[pb.PolicyBasedRoute_Warnings_Code](mapCtx, in.Code)
	out.Data = in.Data
	out.WarningMessage = direct.ValueOf(in.WarningMessage)
	return out
}
