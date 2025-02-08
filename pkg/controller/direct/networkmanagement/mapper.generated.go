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

package networkmanagement

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkmanagement/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func NetworkmanagementVpcFlowLogsConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VpcFlowLogsConfig) *krm.NetworkmanagementVpcFlowLogsConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkmanagementVpcFlowLogsConfigObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: AggregationInterval
	// MISSING: FlowSampling
	// MISSING: Metadata
	// MISSING: MetadataFields
	// MISSING: FilterExpr
	// MISSING: TargetResourceState
	// MISSING: InterconnectAttachment
	// MISSING: VpnTunnel
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func NetworkmanagementVpcFlowLogsConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkmanagementVpcFlowLogsConfigObservedState) *pb.VpcFlowLogsConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcFlowLogsConfig{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: AggregationInterval
	// MISSING: FlowSampling
	// MISSING: Metadata
	// MISSING: MetadataFields
	// MISSING: FilterExpr
	// MISSING: TargetResourceState
	// MISSING: InterconnectAttachment
	// MISSING: VpnTunnel
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func NetworkmanagementVpcFlowLogsConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.VpcFlowLogsConfig) *krm.NetworkmanagementVpcFlowLogsConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkmanagementVpcFlowLogsConfigSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: AggregationInterval
	// MISSING: FlowSampling
	// MISSING: Metadata
	// MISSING: MetadataFields
	// MISSING: FilterExpr
	// MISSING: TargetResourceState
	// MISSING: InterconnectAttachment
	// MISSING: VpnTunnel
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func NetworkmanagementVpcFlowLogsConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkmanagementVpcFlowLogsConfigSpec) *pb.VpcFlowLogsConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcFlowLogsConfig{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: AggregationInterval
	// MISSING: FlowSampling
	// MISSING: Metadata
	// MISSING: MetadataFields
	// MISSING: FilterExpr
	// MISSING: TargetResourceState
	// MISSING: InterconnectAttachment
	// MISSING: VpnTunnel
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func VpcFlowLogsConfig_FromProto(mapCtx *direct.MapContext, in *pb.VpcFlowLogsConfig) *krm.VpcFlowLogsConfig {
	if in == nil {
		return nil
	}
	out := &krm.VpcFlowLogsConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = in.Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.AggregationInterval = direct.Enum_FromProto(mapCtx, in.GetAggregationInterval())
	out.FlowSampling = in.FlowSampling
	out.Metadata = direct.Enum_FromProto(mapCtx, in.GetMetadata())
	out.MetadataFields = in.MetadataFields
	out.FilterExpr = in.FilterExpr
	// MISSING: TargetResourceState
	out.InterconnectAttachment = direct.LazyPtr(in.GetInterconnectAttachment())
	out.VpnTunnel = direct.LazyPtr(in.GetVpnTunnel())
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func VpcFlowLogsConfig_ToProto(mapCtx *direct.MapContext, in *krm.VpcFlowLogsConfig) *pb.VpcFlowLogsConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcFlowLogsConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = in.Description
	if oneof := VpcFlowLogsConfig_State_ToProto(mapCtx, in.State); oneof != nil {
		out.State = oneof
	}
	if oneof := VpcFlowLogsConfig_AggregationInterval_ToProto(mapCtx, in.AggregationInterval); oneof != nil {
		out.AggregationInterval = oneof
	}
	out.FlowSampling = in.FlowSampling
	if oneof := VpcFlowLogsConfig_Metadata_ToProto(mapCtx, in.Metadata); oneof != nil {
		out.Metadata = oneof
	}
	out.MetadataFields = in.MetadataFields
	out.FilterExpr = in.FilterExpr
	// MISSING: TargetResourceState
	if oneof := VpcFlowLogsConfig_InterconnectAttachment_ToProto(mapCtx, in.InterconnectAttachment); oneof != nil {
		out.TargetResource = oneof
	}
	if oneof := VpcFlowLogsConfig_VpnTunnel_ToProto(mapCtx, in.VpnTunnel); oneof != nil {
		out.TargetResource = oneof
	}
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func VpcFlowLogsConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VpcFlowLogsConfig) *krm.VpcFlowLogsConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VpcFlowLogsConfigObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: AggregationInterval
	// MISSING: FlowSampling
	// MISSING: Metadata
	// MISSING: MetadataFields
	// MISSING: FilterExpr
	out.TargetResourceState = direct.Enum_FromProto(mapCtx, in.GetTargetResourceState())
	// MISSING: InterconnectAttachment
	// MISSING: VpnTunnel
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func VpcFlowLogsConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VpcFlowLogsConfigObservedState) *pb.VpcFlowLogsConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcFlowLogsConfig{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: AggregationInterval
	// MISSING: FlowSampling
	// MISSING: Metadata
	// MISSING: MetadataFields
	// MISSING: FilterExpr
	if oneof := VpcFlowLogsConfigObservedState_TargetResourceState_ToProto(mapCtx, in.TargetResourceState); oneof != nil {
		out.TargetResourceState = oneof
	}
	// MISSING: InterconnectAttachment
	// MISSING: VpnTunnel
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
