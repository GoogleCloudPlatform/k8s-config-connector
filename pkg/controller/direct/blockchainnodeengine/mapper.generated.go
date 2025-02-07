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

package blockchainnodeengine

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/blockchainnodeengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/blockchainnodeengine/apiv1/blockchainnodeenginepb"
)
func BlockchainNode_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode) *krm.BlockchainNode {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainNode{}
	out.EthereumDetails = BlockchainNode_EthereumDetails_FromProto(mapCtx, in.GetEthereumDetails())
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.BlockchainType = direct.Enum_FromProto(mapCtx, in.GetBlockchainType())
	// MISSING: ConnectionInfo
	// MISSING: State
	out.PrivateServiceConnectEnabled = direct.LazyPtr(in.GetPrivateServiceConnectEnabled())
	return out
}
func BlockchainNode_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainNode) *pb.BlockchainNode {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode{}
	if oneof := BlockchainNode_EthereumDetails_ToProto(mapCtx, in.EthereumDetails); oneof != nil {
		out.BlockchainTypeDetails = &pb.BlockchainNode_EthereumDetails_{EthereumDetails: oneof}
	}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	if oneof := BlockchainNode_BlockchainType_ToProto(mapCtx, in.BlockchainType); oneof != nil {
		out.BlockchainType = oneof
	}
	// MISSING: ConnectionInfo
	// MISSING: State
	out.PrivateServiceConnectEnabled = direct.ValueOf(in.PrivateServiceConnectEnabled)
	return out
}
func BlockchainNodeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode) *krm.BlockchainNodeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainNodeObservedState{}
	out.EthereumDetails = BlockchainNode_EthereumDetailsObservedState_FromProto(mapCtx, in.GetEthereumDetails())
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: BlockchainType
	out.ConnectionInfo = BlockchainNode_ConnectionInfo_FromProto(mapCtx, in.GetConnectionInfo())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: PrivateServiceConnectEnabled
	return out
}
func BlockchainNodeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainNodeObservedState) *pb.BlockchainNode {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode{}
	if oneof := BlockchainNode_EthereumDetailsObservedState_ToProto(mapCtx, in.EthereumDetails); oneof != nil {
		out.BlockchainTypeDetails = &pb.BlockchainNode_EthereumDetails_{EthereumDetails: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: BlockchainType
	out.ConnectionInfo = BlockchainNode_ConnectionInfo_ToProto(mapCtx, in.ConnectionInfo)
	out.State = direct.Enum_ToProto[pb.BlockchainNode_State](mapCtx, in.State)
	// MISSING: PrivateServiceConnectEnabled
	return out
}
func BlockchainNode_ConnectionInfo_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode_ConnectionInfo) *krm.BlockchainNode_ConnectionInfo {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainNode_ConnectionInfo{}
	// MISSING: EndpointInfo
	// MISSING: ServiceAttachment
	return out
}
func BlockchainNode_ConnectionInfo_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainNode_ConnectionInfo) *pb.BlockchainNode_ConnectionInfo {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode_ConnectionInfo{}
	// MISSING: EndpointInfo
	// MISSING: ServiceAttachment
	return out
}
func BlockchainNode_ConnectionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode_ConnectionInfo) *krm.BlockchainNode_ConnectionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainNode_ConnectionInfoObservedState{}
	out.EndpointInfo = BlockchainNode_ConnectionInfo_EndpointInfo_FromProto(mapCtx, in.GetEndpointInfo())
	out.ServiceAttachment = direct.LazyPtr(in.GetServiceAttachment())
	return out
}
func BlockchainNode_ConnectionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainNode_ConnectionInfoObservedState) *pb.BlockchainNode_ConnectionInfo {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode_ConnectionInfo{}
	out.EndpointInfo = BlockchainNode_ConnectionInfo_EndpointInfo_ToProto(mapCtx, in.EndpointInfo)
	out.ServiceAttachment = direct.ValueOf(in.ServiceAttachment)
	return out
}
func BlockchainNode_ConnectionInfo_EndpointInfo_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode_ConnectionInfo_EndpointInfo) *krm.BlockchainNode_ConnectionInfo_EndpointInfo {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainNode_ConnectionInfo_EndpointInfo{}
	// MISSING: JsonRpcApiEndpoint
	// MISSING: WebsocketsApiEndpoint
	return out
}
func BlockchainNode_ConnectionInfo_EndpointInfo_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainNode_ConnectionInfo_EndpointInfo) *pb.BlockchainNode_ConnectionInfo_EndpointInfo {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode_ConnectionInfo_EndpointInfo{}
	// MISSING: JsonRpcApiEndpoint
	// MISSING: WebsocketsApiEndpoint
	return out
}
func BlockchainNode_ConnectionInfo_EndpointInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode_ConnectionInfo_EndpointInfo) *krm.BlockchainNode_ConnectionInfo_EndpointInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainNode_ConnectionInfo_EndpointInfoObservedState{}
	out.JsonRpcApiEndpoint = direct.LazyPtr(in.GetJsonRpcApiEndpoint())
	out.WebsocketsApiEndpoint = direct.LazyPtr(in.GetWebsocketsApiEndpoint())
	return out
}
func BlockchainNode_ConnectionInfo_EndpointInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainNode_ConnectionInfo_EndpointInfoObservedState) *pb.BlockchainNode_ConnectionInfo_EndpointInfo {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode_ConnectionInfo_EndpointInfo{}
	out.JsonRpcApiEndpoint = direct.ValueOf(in.JsonRpcApiEndpoint)
	out.WebsocketsApiEndpoint = direct.ValueOf(in.WebsocketsApiEndpoint)
	return out
}
func BlockchainNode_EthereumDetails_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode_EthereumDetails) *krm.BlockchainNode_EthereumDetails {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainNode_EthereumDetails{}
	out.GethDetails = BlockchainNode_EthereumDetails_GethDetails_FromProto(mapCtx, in.GetGethDetails())
	out.Network = direct.Enum_FromProto(mapCtx, in.GetNetwork())
	out.NodeType = direct.Enum_FromProto(mapCtx, in.GetNodeType())
	out.ExecutionClient = direct.Enum_FromProto(mapCtx, in.GetExecutionClient())
	out.ConsensusClient = direct.Enum_FromProto(mapCtx, in.GetConsensusClient())
	out.ApiEnableAdmin = in.ApiEnableAdmin
	out.ApiEnableDebug = in.ApiEnableDebug
	// MISSING: AdditionalEndpoints
	out.ValidatorConfig = BlockchainNode_EthereumDetails_ValidatorConfig_FromProto(mapCtx, in.GetValidatorConfig())
	return out
}
func BlockchainNode_EthereumDetails_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainNode_EthereumDetails) *pb.BlockchainNode_EthereumDetails {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode_EthereumDetails{}
	if oneof := BlockchainNode_EthereumDetails_GethDetails_ToProto(mapCtx, in.GethDetails); oneof != nil {
		out.ExecutionClientDetails = &pb.BlockchainNode_EthereumDetails_GethDetails_{GethDetails: oneof}
	}
	if oneof := BlockchainNode_EthereumDetails_Network_ToProto(mapCtx, in.Network); oneof != nil {
		out.Network = oneof
	}
	if oneof := BlockchainNode_EthereumDetails_NodeType_ToProto(mapCtx, in.NodeType); oneof != nil {
		out.NodeType = oneof
	}
	if oneof := BlockchainNode_EthereumDetails_ExecutionClient_ToProto(mapCtx, in.ExecutionClient); oneof != nil {
		out.ExecutionClient = oneof
	}
	if oneof := BlockchainNode_EthereumDetails_ConsensusClient_ToProto(mapCtx, in.ConsensusClient); oneof != nil {
		out.ConsensusClient = oneof
	}
	out.ApiEnableAdmin = in.ApiEnableAdmin
	out.ApiEnableDebug = in.ApiEnableDebug
	// MISSING: AdditionalEndpoints
	if oneof := BlockchainNode_EthereumDetails_ValidatorConfig_ToProto(mapCtx, in.ValidatorConfig); oneof != nil {
		out.ValidatorConfig = &pb.BlockchainNode_EthereumDetails_ValidatorConfig_{ValidatorConfig: oneof}
	}
	return out
}
func BlockchainNode_EthereumDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode_EthereumDetails) *krm.BlockchainNode_EthereumDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainNode_EthereumDetailsObservedState{}
	// MISSING: GethDetails
	// MISSING: Network
	// MISSING: NodeType
	// MISSING: ExecutionClient
	// MISSING: ConsensusClient
	// MISSING: ApiEnableAdmin
	// MISSING: ApiEnableDebug
	out.AdditionalEndpoints = BlockchainNode_EthereumDetails_EthereumEndpoints_FromProto(mapCtx, in.GetAdditionalEndpoints())
	// MISSING: ValidatorConfig
	return out
}
func BlockchainNode_EthereumDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainNode_EthereumDetailsObservedState) *pb.BlockchainNode_EthereumDetails {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode_EthereumDetails{}
	// MISSING: GethDetails
	// MISSING: Network
	// MISSING: NodeType
	// MISSING: ExecutionClient
	// MISSING: ConsensusClient
	// MISSING: ApiEnableAdmin
	// MISSING: ApiEnableDebug
	if oneof := BlockchainNode_EthereumDetails_EthereumEndpoints_ToProto(mapCtx, in.AdditionalEndpoints); oneof != nil {
		out.AdditionalEndpoints = &pb.BlockchainNode_EthereumDetails_AdditionalEndpoints{AdditionalEndpoints: oneof}
	}
	// MISSING: ValidatorConfig
	return out
}
func BlockchainNode_EthereumDetails_EthereumEndpoints_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode_EthereumDetails_EthereumEndpoints) *krm.BlockchainNode_EthereumDetails_EthereumEndpoints {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainNode_EthereumDetails_EthereumEndpoints{}
	// MISSING: BeaconApiEndpoint
	// MISSING: BeaconPrometheusMetricsApiEndpoint
	// MISSING: ExecutionClientPrometheusMetricsApiEndpoint
	return out
}
func BlockchainNode_EthereumDetails_EthereumEndpoints_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainNode_EthereumDetails_EthereumEndpoints) *pb.BlockchainNode_EthereumDetails_EthereumEndpoints {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode_EthereumDetails_EthereumEndpoints{}
	// MISSING: BeaconApiEndpoint
	// MISSING: BeaconPrometheusMetricsApiEndpoint
	// MISSING: ExecutionClientPrometheusMetricsApiEndpoint
	return out
}
func BlockchainNode_EthereumDetails_EthereumEndpointsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode_EthereumDetails_EthereumEndpoints) *krm.BlockchainNode_EthereumDetails_EthereumEndpointsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainNode_EthereumDetails_EthereumEndpointsObservedState{}
	out.BeaconApiEndpoint = direct.LazyPtr(in.GetBeaconApiEndpoint())
	out.BeaconPrometheusMetricsApiEndpoint = direct.LazyPtr(in.GetBeaconPrometheusMetricsApiEndpoint())
	out.ExecutionClientPrometheusMetricsApiEndpoint = direct.LazyPtr(in.GetExecutionClientPrometheusMetricsApiEndpoint())
	return out
}
func BlockchainNode_EthereumDetails_EthereumEndpointsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainNode_EthereumDetails_EthereumEndpointsObservedState) *pb.BlockchainNode_EthereumDetails_EthereumEndpoints {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode_EthereumDetails_EthereumEndpoints{}
	out.BeaconApiEndpoint = direct.ValueOf(in.BeaconApiEndpoint)
	out.BeaconPrometheusMetricsApiEndpoint = direct.ValueOf(in.BeaconPrometheusMetricsApiEndpoint)
	out.ExecutionClientPrometheusMetricsApiEndpoint = direct.ValueOf(in.ExecutionClientPrometheusMetricsApiEndpoint)
	return out
}
func BlockchainNode_EthereumDetails_GethDetails_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode_EthereumDetails_GethDetails) *krm.BlockchainNode_EthereumDetails_GethDetails {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainNode_EthereumDetails_GethDetails{}
	out.GarbageCollectionMode = direct.Enum_FromProto(mapCtx, in.GetGarbageCollectionMode())
	return out
}
func BlockchainNode_EthereumDetails_GethDetails_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainNode_EthereumDetails_GethDetails) *pb.BlockchainNode_EthereumDetails_GethDetails {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode_EthereumDetails_GethDetails{}
	if oneof := BlockchainNode_EthereumDetails_GethDetails_GarbageCollectionMode_ToProto(mapCtx, in.GarbageCollectionMode); oneof != nil {
		out.GarbageCollectionMode = oneof
	}
	return out
}
func BlockchainNode_EthereumDetails_ValidatorConfig_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode_EthereumDetails_ValidatorConfig) *krm.BlockchainNode_EthereumDetails_ValidatorConfig {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainNode_EthereumDetails_ValidatorConfig{}
	out.MevRelayUrls = in.MevRelayUrls
	out.ManagedValidatorClient = direct.LazyPtr(in.GetManagedValidatorClient())
	out.BeaconFeeRecipient = in.BeaconFeeRecipient
	return out
}
func BlockchainNode_EthereumDetails_ValidatorConfig_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainNode_EthereumDetails_ValidatorConfig) *pb.BlockchainNode_EthereumDetails_ValidatorConfig {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode_EthereumDetails_ValidatorConfig{}
	out.MevRelayUrls = in.MevRelayUrls
	out.ManagedValidatorClient = direct.ValueOf(in.ManagedValidatorClient)
	out.BeaconFeeRecipient = in.BeaconFeeRecipient
	return out
}
func BlockchainnodeengineBlockchainNodeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode) *krm.BlockchainnodeengineBlockchainNodeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainnodeengineBlockchainNodeObservedState{}
	// MISSING: EthereumDetails
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: BlockchainType
	// MISSING: ConnectionInfo
	// MISSING: State
	// MISSING: PrivateServiceConnectEnabled
	return out
}
func BlockchainnodeengineBlockchainNodeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainnodeengineBlockchainNodeObservedState) *pb.BlockchainNode {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode{}
	// MISSING: EthereumDetails
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: BlockchainType
	// MISSING: ConnectionInfo
	// MISSING: State
	// MISSING: PrivateServiceConnectEnabled
	return out
}
func BlockchainnodeengineBlockchainNodeSpec_FromProto(mapCtx *direct.MapContext, in *pb.BlockchainNode) *krm.BlockchainnodeengineBlockchainNodeSpec {
	if in == nil {
		return nil
	}
	out := &krm.BlockchainnodeengineBlockchainNodeSpec{}
	// MISSING: EthereumDetails
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: BlockchainType
	// MISSING: ConnectionInfo
	// MISSING: State
	// MISSING: PrivateServiceConnectEnabled
	return out
}
func BlockchainnodeengineBlockchainNodeSpec_ToProto(mapCtx *direct.MapContext, in *krm.BlockchainnodeengineBlockchainNodeSpec) *pb.BlockchainNode {
	if in == nil {
		return nil
	}
	out := &pb.BlockchainNode{}
	// MISSING: EthereumDetails
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: BlockchainType
	// MISSING: ConnectionInfo
	// MISSING: State
	// MISSING: PrivateServiceConnectEnabled
	return out
}
