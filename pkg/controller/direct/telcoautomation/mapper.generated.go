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

package telcoautomation

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/telcoautomation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/telcoautomation/apiv1/telcoautomationpb"
)
func FullManagementConfig_FromProto(mapCtx *direct.MapContext, in *pb.FullManagementConfig) *krm.FullManagementConfig {
	if in == nil {
		return nil
	}
	out := &krm.FullManagementConfig{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Subnet = direct.LazyPtr(in.GetSubnet())
	out.MasterIpv4CidrBlock = direct.LazyPtr(in.GetMasterIpv4CidrBlock())
	out.ClusterCidrBlock = direct.LazyPtr(in.GetClusterCidrBlock())
	out.ServicesCidrBlock = direct.LazyPtr(in.GetServicesCidrBlock())
	out.ClusterNamedRange = direct.LazyPtr(in.GetClusterNamedRange())
	out.ServicesNamedRange = direct.LazyPtr(in.GetServicesNamedRange())
	out.MasterAuthorizedNetworksConfig = MasterAuthorizedNetworksConfig_FromProto(mapCtx, in.GetMasterAuthorizedNetworksConfig())
	return out
}
func FullManagementConfig_ToProto(mapCtx *direct.MapContext, in *krm.FullManagementConfig) *pb.FullManagementConfig {
	if in == nil {
		return nil
	}
	out := &pb.FullManagementConfig{}
	out.Network = direct.ValueOf(in.Network)
	out.Subnet = direct.ValueOf(in.Subnet)
	out.MasterIpv4CidrBlock = direct.ValueOf(in.MasterIpv4CidrBlock)
	out.ClusterCidrBlock = direct.ValueOf(in.ClusterCidrBlock)
	out.ServicesCidrBlock = direct.ValueOf(in.ServicesCidrBlock)
	out.ClusterNamedRange = direct.ValueOf(in.ClusterNamedRange)
	out.ServicesNamedRange = direct.ValueOf(in.ServicesNamedRange)
	out.MasterAuthorizedNetworksConfig = MasterAuthorizedNetworksConfig_ToProto(mapCtx, in.MasterAuthorizedNetworksConfig)
	return out
}
func ManagementConfig_FromProto(mapCtx *direct.MapContext, in *pb.ManagementConfig) *krm.ManagementConfig {
	if in == nil {
		return nil
	}
	out := &krm.ManagementConfig{}
	out.StandardManagementConfig = StandardManagementConfig_FromProto(mapCtx, in.GetStandardManagementConfig())
	out.FullManagementConfig = FullManagementConfig_FromProto(mapCtx, in.GetFullManagementConfig())
	return out
}
func ManagementConfig_ToProto(mapCtx *direct.MapContext, in *krm.ManagementConfig) *pb.ManagementConfig {
	if in == nil {
		return nil
	}
	out := &pb.ManagementConfig{}
	if oneof := StandardManagementConfig_ToProto(mapCtx, in.StandardManagementConfig); oneof != nil {
		out.OneofConfig = &pb.ManagementConfig_StandardManagementConfig{StandardManagementConfig: oneof}
	}
	if oneof := FullManagementConfig_ToProto(mapCtx, in.FullManagementConfig); oneof != nil {
		out.OneofConfig = &pb.ManagementConfig_FullManagementConfig{FullManagementConfig: oneof}
	}
	return out
}
func MasterAuthorizedNetworksConfig_FromProto(mapCtx *direct.MapContext, in *pb.MasterAuthorizedNetworksConfig) *krm.MasterAuthorizedNetworksConfig {
	if in == nil {
		return nil
	}
	out := &krm.MasterAuthorizedNetworksConfig{}
	out.CidrBlocks = direct.Slice_FromProto(mapCtx, in.CidrBlocks, MasterAuthorizedNetworksConfig_CidrBlock_FromProto)
	return out
}
func MasterAuthorizedNetworksConfig_ToProto(mapCtx *direct.MapContext, in *krm.MasterAuthorizedNetworksConfig) *pb.MasterAuthorizedNetworksConfig {
	if in == nil {
		return nil
	}
	out := &pb.MasterAuthorizedNetworksConfig{}
	out.CidrBlocks = direct.Slice_ToProto(mapCtx, in.CidrBlocks, MasterAuthorizedNetworksConfig_CidrBlock_ToProto)
	return out
}
func MasterAuthorizedNetworksConfig_CidrBlock_FromProto(mapCtx *direct.MapContext, in *pb.MasterAuthorizedNetworksConfig_CidrBlock) *krm.MasterAuthorizedNetworksConfig_CidrBlock {
	if in == nil {
		return nil
	}
	out := &krm.MasterAuthorizedNetworksConfig_CidrBlock{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.CidrBlock = direct.LazyPtr(in.GetCidrBlock())
	return out
}
func MasterAuthorizedNetworksConfig_CidrBlock_ToProto(mapCtx *direct.MapContext, in *krm.MasterAuthorizedNetworksConfig_CidrBlock) *pb.MasterAuthorizedNetworksConfig_CidrBlock {
	if in == nil {
		return nil
	}
	out := &pb.MasterAuthorizedNetworksConfig_CidrBlock{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.CidrBlock = direct.ValueOf(in.CidrBlock)
	return out
}
func OrchestrationCluster_FromProto(mapCtx *direct.MapContext, in *pb.OrchestrationCluster) *krm.OrchestrationCluster {
	if in == nil {
		return nil
	}
	out := &krm.OrchestrationCluster{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ManagementConfig = ManagementConfig_FromProto(mapCtx, in.GetManagementConfig())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: TnaVersion
	// MISSING: State
	return out
}
func OrchestrationCluster_ToProto(mapCtx *direct.MapContext, in *krm.OrchestrationCluster) *pb.OrchestrationCluster {
	if in == nil {
		return nil
	}
	out := &pb.OrchestrationCluster{}
	out.Name = direct.ValueOf(in.Name)
	out.ManagementConfig = ManagementConfig_ToProto(mapCtx, in.ManagementConfig)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: TnaVersion
	// MISSING: State
	return out
}
func OrchestrationClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OrchestrationCluster) *krm.OrchestrationClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrchestrationClusterObservedState{}
	// MISSING: Name
	// MISSING: ManagementConfig
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	out.TnaVersion = direct.LazyPtr(in.GetTnaVersion())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func OrchestrationClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrchestrationClusterObservedState) *pb.OrchestrationCluster {
	if in == nil {
		return nil
	}
	out := &pb.OrchestrationCluster{}
	// MISSING: Name
	// MISSING: ManagementConfig
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	out.TnaVersion = direct.ValueOf(in.TnaVersion)
	out.State = direct.Enum_ToProto[pb.OrchestrationCluster_State](mapCtx, in.State)
	return out
}
func StandardManagementConfig_FromProto(mapCtx *direct.MapContext, in *pb.StandardManagementConfig) *krm.StandardManagementConfig {
	if in == nil {
		return nil
	}
	out := &krm.StandardManagementConfig{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Subnet = direct.LazyPtr(in.GetSubnet())
	out.MasterIpv4CidrBlock = direct.LazyPtr(in.GetMasterIpv4CidrBlock())
	out.ClusterCidrBlock = direct.LazyPtr(in.GetClusterCidrBlock())
	out.ServicesCidrBlock = direct.LazyPtr(in.GetServicesCidrBlock())
	out.ClusterNamedRange = direct.LazyPtr(in.GetClusterNamedRange())
	out.ServicesNamedRange = direct.LazyPtr(in.GetServicesNamedRange())
	out.MasterAuthorizedNetworksConfig = MasterAuthorizedNetworksConfig_FromProto(mapCtx, in.GetMasterAuthorizedNetworksConfig())
	return out
}
func StandardManagementConfig_ToProto(mapCtx *direct.MapContext, in *krm.StandardManagementConfig) *pb.StandardManagementConfig {
	if in == nil {
		return nil
	}
	out := &pb.StandardManagementConfig{}
	out.Network = direct.ValueOf(in.Network)
	out.Subnet = direct.ValueOf(in.Subnet)
	out.MasterIpv4CidrBlock = direct.ValueOf(in.MasterIpv4CidrBlock)
	out.ClusterCidrBlock = direct.ValueOf(in.ClusterCidrBlock)
	out.ServicesCidrBlock = direct.ValueOf(in.ServicesCidrBlock)
	out.ClusterNamedRange = direct.ValueOf(in.ClusterNamedRange)
	out.ServicesNamedRange = direct.ValueOf(in.ServicesNamedRange)
	out.MasterAuthorizedNetworksConfig = MasterAuthorizedNetworksConfig_ToProto(mapCtx, in.MasterAuthorizedNetworksConfig)
	return out
}
func TelcoautomationOrchestrationClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OrchestrationCluster) *krm.TelcoautomationOrchestrationClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TelcoautomationOrchestrationClusterObservedState{}
	// MISSING: Name
	// MISSING: ManagementConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: TnaVersion
	// MISSING: State
	return out
}
func TelcoautomationOrchestrationClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TelcoautomationOrchestrationClusterObservedState) *pb.OrchestrationCluster {
	if in == nil {
		return nil
	}
	out := &pb.OrchestrationCluster{}
	// MISSING: Name
	// MISSING: ManagementConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: TnaVersion
	// MISSING: State
	return out
}
func TelcoautomationOrchestrationClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.OrchestrationCluster) *krm.TelcoautomationOrchestrationClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.TelcoautomationOrchestrationClusterSpec{}
	// MISSING: Name
	// MISSING: ManagementConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: TnaVersion
	// MISSING: State
	return out
}
func TelcoautomationOrchestrationClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.TelcoautomationOrchestrationClusterSpec) *pb.OrchestrationCluster {
	if in == nil {
		return nil
	}
	out := &pb.OrchestrationCluster{}
	// MISSING: Name
	// MISSING: ManagementConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: TnaVersion
	// MISSING: State
	return out
}
