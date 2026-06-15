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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkmanagement/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkManagementVpcFlowLogsConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.VpcFlowLogsConfig) *krm.NetworkManagementVpcFlowLogsConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkManagementVpcFlowLogsConfigSpec{}
	out.Description = in.Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.AggregationInterval = direct.Enum_FromProto(mapCtx, in.GetAggregationInterval())
	out.FlowSampling = in.FlowSampling
	out.Metadata = direct.Enum_FromProto(mapCtx, in.GetMetadata())
	out.MetadataFields = in.MetadataFields
	out.FilterExpr = in.FilterExpr

	switch target := in.GetTargetResource().(type) {
	case *pb.VpcFlowLogsConfig_InterconnectAttachment:
		if target.InterconnectAttachment != "" {
			out.InterconnectAttachmentRef = &refsv1beta1.ComputeInterconnectAttachmentRef{External: target.InterconnectAttachment}
		}
	case *pb.VpcFlowLogsConfig_VpnTunnel:
		if target.VpnTunnel != "" {
			out.VPNTunnelRef = &refsv1beta1.ComputeVPNTunnelRef{External: target.VpnTunnel}
		}
	}

	out.Labels = in.Labels
	return out
}

func NetworkManagementVpcFlowLogsConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkManagementVpcFlowLogsConfigSpec) *pb.VpcFlowLogsConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcFlowLogsConfig{}
	out.Description = in.Description
	if in.State != nil {
		out.State = direct.PtrTo(direct.Enum_ToProto[pb.VpcFlowLogsConfig_State](mapCtx, in.State))
	}
	if in.AggregationInterval != nil {
		out.AggregationInterval = direct.PtrTo(direct.Enum_ToProto[pb.VpcFlowLogsConfig_AggregationInterval](mapCtx, in.AggregationInterval))
	}
	out.FlowSampling = in.FlowSampling
	if in.Metadata != nil {
		out.Metadata = direct.PtrTo(direct.Enum_ToProto[pb.VpcFlowLogsConfig_Metadata](mapCtx, in.Metadata))
	}
	out.MetadataFields = in.MetadataFields
	out.FilterExpr = in.FilterExpr

	if in.InterconnectAttachmentRef != nil {
		out.TargetResource = &pb.VpcFlowLogsConfig_InterconnectAttachment{
			InterconnectAttachment: in.InterconnectAttachmentRef.External,
		}
	} else if in.VPNTunnelRef != nil {
		out.TargetResource = &pb.VpcFlowLogsConfig_VpnTunnel{
			VpnTunnel: in.VPNTunnelRef.External,
		}
	}

	out.Labels = in.Labels
	return out
}

func NetworkManagementVpcFlowLogsConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VpcFlowLogsConfig) *krm.NetworkManagementVpcFlowLogsConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkManagementVpcFlowLogsConfigObservedState{}
	out.TargetResourceState = direct.Enum_FromProto(mapCtx, in.GetTargetResourceState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func NetworkManagementVpcFlowLogsConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkManagementVpcFlowLogsConfigObservedState) *pb.VpcFlowLogsConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcFlowLogsConfig{}
	if in.TargetResourceState != nil {
		out.TargetResourceState = direct.PtrTo(direct.Enum_ToProto[pb.VpcFlowLogsConfig_TargetResourceState](mapCtx, in.TargetResourceState))
	}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
