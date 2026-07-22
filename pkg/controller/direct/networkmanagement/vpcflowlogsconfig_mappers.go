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

package networkmanagement

import (
	pb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkmanagement/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkManagementVPCFlowLogsConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VpcFlowLogsConfig) *krm.NetworkManagementVPCFlowLogsConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkManagementVPCFlowLogsConfigObservedState{}
	out.TargetResourceState = direct.Enum_FromProto(mapCtx, in.GetTargetResourceState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func NetworkManagementVPCFlowLogsConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkManagementVPCFlowLogsConfigObservedState) *pb.VpcFlowLogsConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcFlowLogsConfig{}
	if oneof := NetworkManagementVPCFlowLogsConfigObservedState_TargetResourceState_ToProto(mapCtx, in.TargetResourceState); oneof != nil {
		out.TargetResourceState = oneof
	}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func NetworkManagementVPCFlowLogsConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.VpcFlowLogsConfig) *krm.NetworkManagementVPCFlowLogsConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkManagementVPCFlowLogsConfigSpec{}
	out.Description = in.Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.AggregationInterval = direct.Enum_FromProto(mapCtx, in.GetAggregationInterval())
	out.FlowSampling = in.FlowSampling
	out.Metadata = direct.Enum_FromProto(mapCtx, in.GetMetadata())
	out.MetadataFields = in.MetadataFields
	out.FilterExpr = in.FilterExpr
	if in.GetInterconnectAttachment() != "" {
		out.InterconnectAttachmentRef = &computev1beta1.ComputeInterconnectAttachmentRef{External: in.GetInterconnectAttachment()}
	}
	if in.GetVpnTunnel() != "" {
		out.VPNTunnelRef = &computev1beta1.ComputeVPNTunnelRef{External: in.GetVpnTunnel()}
	}
	out.Labels = in.Labels
	return out
}

func NetworkManagementVPCFlowLogsConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkManagementVPCFlowLogsConfigSpec) *pb.VpcFlowLogsConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcFlowLogsConfig{}
	out.Description = in.Description
	if oneof := NetworkManagementVPCFlowLogsConfigSpec_State_ToProto(mapCtx, in.State); oneof != nil {
		out.State = oneof
	}
	if oneof := NetworkManagementVPCFlowLogsConfigSpec_AggregationInterval_ToProto(mapCtx, in.AggregationInterval); oneof != nil {
		out.AggregationInterval = oneof
	}
	out.FlowSampling = in.FlowSampling
	if oneof := NetworkManagementVPCFlowLogsConfigSpec_Metadata_ToProto(mapCtx, in.Metadata); oneof != nil {
		out.Metadata = oneof
	}
	out.MetadataFields = in.MetadataFields
	out.FilterExpr = in.FilterExpr
	if in.InterconnectAttachmentRef != nil {
		out.TargetResource = &pb.VpcFlowLogsConfig_InterconnectAttachment{
			InterconnectAttachment: in.InterconnectAttachmentRef.External,
		}
	}
	if in.VPNTunnelRef != nil {
		out.TargetResource = &pb.VpcFlowLogsConfig_VpnTunnel{
			VpnTunnel: in.VPNTunnelRef.External,
		}
	}
	out.Labels = in.Labels
	return out
}

func NetworkManagementVPCFlowLogsConfigObservedState_TargetResourceState_ToProto(mapCtx *direct.MapContext, in *string) *pb.VpcFlowLogsConfig_TargetResourceState {
	if in == nil {
		return nil
	}
	return direct.PtrTo(direct.Enum_ToProto[pb.VpcFlowLogsConfig_TargetResourceState](mapCtx, in))
}

func NetworkManagementVPCFlowLogsConfigSpec_State_ToProto(mapCtx *direct.MapContext, in *string) *pb.VpcFlowLogsConfig_State {
	if in == nil {
		return nil
	}
	return direct.PtrTo(direct.Enum_ToProto[pb.VpcFlowLogsConfig_State](mapCtx, in))
}

func NetworkManagementVPCFlowLogsConfigSpec_AggregationInterval_ToProto(mapCtx *direct.MapContext, in *string) *pb.VpcFlowLogsConfig_AggregationInterval {
	if in == nil {
		return nil
	}
	return direct.PtrTo(direct.Enum_ToProto[pb.VpcFlowLogsConfig_AggregationInterval](mapCtx, in))
}

func NetworkManagementVPCFlowLogsConfigSpec_Metadata_ToProto(mapCtx *direct.MapContext, in *string) *pb.VpcFlowLogsConfig_Metadata {
	if in == nil {
		return nil
	}
	return direct.PtrTo(direct.Enum_ToProto[pb.VpcFlowLogsConfig_Metadata](mapCtx, in))
}
