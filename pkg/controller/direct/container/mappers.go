// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package container

import (
	pb "cloud.google.com/go/container/apiv1/containerpb"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// AdditionalPodNetworkConfig_FromProto maps AdditionalPodNetworkConfig from proto.
// Handwritten because max_pods_per_node is flattened from a nested MaxPodsConstraint message.
func AdditionalPodNetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.AdditionalPodNetworkConfig) *krm.AdditionalPodNetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.AdditionalPodNetworkConfig{}
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &computev1beta1.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	out.SecondaryPodRange = direct.LazyPtr(in.GetSecondaryPodRange())
	if in.GetMaxPodsPerNode() != nil {
		out.MaxPodsPerNode = direct.LazyPtr(int(in.GetMaxPodsPerNode().GetMaxPodsPerNode()))
	}
	return out
}

// AdditionalPodNetworkConfig_ToProto maps AdditionalPodNetworkConfig to proto.
// Handwritten because max_pods_per_node is flattened to a nested MaxPodsConstraint message.
func AdditionalPodNetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.AdditionalPodNetworkConfig) *pb.AdditionalPodNetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.AdditionalPodNetworkConfig{}
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	out.SecondaryPodRange = direct.ValueOf(in.SecondaryPodRange)
	if in.MaxPodsPerNode != nil {
		out.MaxPodsPerNode = &pb.MaxPodsConstraint{
			MaxPodsPerNode: int64(direct.ValueOf(in.MaxPodsPerNode)),
		}
	}
	return out
}

// NodeConfig_AdvancedMachineFeatures_FromProto maps NodeConfig_AdvancedMachineFeatures from proto.
// Handwritten because it is unreachable in types.generated.go.
func NodeConfig_AdvancedMachineFeatures_FromProto(mapCtx *direct.MapContext, in *pb.AdvancedMachineFeatures) *krm.NodeConfig_AdvancedMachineFeatures {
	if in == nil {
		return nil
	}
	out := &krm.NodeConfig_AdvancedMachineFeatures{}
	out.EnableNestedVirtualization = in.EnableNestedVirtualization
	if in.ThreadsPerCore != nil {
		out.ThreadsPerCore = direct.LazyPtr(int(*in.ThreadsPerCore))
	}
	return out
}

// NodeConfig_AdvancedMachineFeatures_ToProto maps NodeConfig_AdvancedMachineFeatures to proto.
// Handwritten because it is unreachable in types.generated.go.
func NodeConfig_AdvancedMachineFeatures_ToProto(mapCtx *direct.MapContext, in *krm.NodeConfig_AdvancedMachineFeatures) *pb.AdvancedMachineFeatures {
	if in == nil {
		return nil
	}
	out := &pb.AdvancedMachineFeatures{}
	out.EnableNestedVirtualization = in.EnableNestedVirtualization
	if in.ThreadsPerCore != nil {
		out.ThreadsPerCore = direct.LazyPtr(int64(*in.ThreadsPerCore))
	}
	return out
}

// NodePoolBlueGreenSettings_FromProto maps NodePoolBlueGreenSettings from proto.
// Handwritten because it requires mapping standardRolloutPolicy.
func NodePoolBlueGreenSettings_FromProto(mapCtx *direct.MapContext, in *pb.BlueGreenSettings) *krm.NodePoolBlueGreenSettings {
	if in == nil {
		return nil
	}
	out := &krm.NodePoolBlueGreenSettings{}
	out.NodePoolSoakDuration = direct.StringDuration_FromProto(mapCtx, in.GetNodePoolSoakDuration())
	out.StandardRolloutPolicy = StandardRolloutPolicy_FromProto(mapCtx, in.GetStandardRolloutPolicy())
	return out
}

// NodePoolBlueGreenSettings_ToProto maps NodePoolBlueGreenSettings to proto.
// Handwritten because it requires mapping standardRolloutPolicy.
func NodePoolBlueGreenSettings_ToProto(mapCtx *direct.MapContext, in *krm.NodePoolBlueGreenSettings) *pb.BlueGreenSettings {
	if in == nil {
		return nil
	}
	out := &pb.BlueGreenSettings{}
	out.NodePoolSoakDuration = direct.StringDuration_ToProto(mapCtx, in.NodePoolSoakDuration)
	if in.StandardRolloutPolicy != nil {
		out.RolloutPolicy = &pb.BlueGreenSettings_StandardRolloutPolicy_{
			StandardRolloutPolicy: StandardRolloutPolicy_ToProto(mapCtx, in.StandardRolloutPolicy),
		}
	}
	return out
}

// NodePool_UpdateConfig_BlueGreenSettings_FromProto maps NodePool_UpdateConfig_BlueGreenSettings from proto.
func NodePool_UpdateConfig_BlueGreenSettings_FromProto(mapCtx *direct.MapContext, in *pb.BlueGreenSettings) *krm.NodePool_UpdateConfig_BlueGreenSettings {
	if in == nil {
		return nil
	}
	out := &krm.NodePool_UpdateConfig_BlueGreenSettings{}
	out.NodePoolSoakDuration = direct.StringDuration_FromProto(mapCtx, in.GetNodePoolSoakDuration())
	out.StandardRolloutPolicy = StandardRolloutPolicy_FromProto(mapCtx, in.GetStandardRolloutPolicy())
	return out
}

// NodePool_UpdateConfig_BlueGreenSettings_ToProto maps NodePool_UpdateConfig_BlueGreenSettings to proto.
func NodePool_UpdateConfig_BlueGreenSettings_ToProto(mapCtx *direct.MapContext, in *krm.NodePool_UpdateConfig_BlueGreenSettings) *pb.BlueGreenSettings {
	if in == nil {
		return nil
	}
	out := &pb.BlueGreenSettings{}
	out.NodePoolSoakDuration = direct.StringDuration_ToProto(mapCtx, in.NodePoolSoakDuration)
	if in.StandardRolloutPolicy != nil {
		out.RolloutPolicy = &pb.BlueGreenSettings_StandardRolloutPolicy_{
			StandardRolloutPolicy: StandardRolloutPolicy_ToProto(mapCtx, in.StandardRolloutPolicy),
		}
	}
	return out
}

// NodePoolUpgradeSettings_Strategy_ToProto maps Strategy enum.
func NodePoolUpgradeSettings_Strategy_ToProto(mapCtx *direct.MapContext, in *string) *pb.NodePoolUpdateStrategy {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.NodePoolUpdateStrategy](mapCtx, in)
	return &v
}

// NodePoolUpgradeSettings_FromProto maps NodePoolUpgradeSettings from proto.
// Handwritten due to direct.Enum conversion.
func NodePoolUpgradeSettings_FromProto(mapCtx *direct.MapContext, in *pb.NodePool_UpgradeSettings) *krm.NodePoolUpgradeSettings {
	if in == nil {
		return nil
	}
	out := &krm.NodePoolUpgradeSettings{}
	out.MaxSurge = direct.LazyPtr(in.GetMaxSurge())
	out.MaxUnavailable = direct.LazyPtr(in.GetMaxUnavailable())
	out.Strategy = direct.Enum_FromProto(mapCtx, in.GetStrategy())
	out.BlueGreenSettings = NodePoolBlueGreenSettings_FromProto(mapCtx, in.GetBlueGreenSettings())
	return out
}

// NodePoolUpgradeSettings_ToProto maps NodePoolUpgradeSettings to proto.
// Handwritten due to direct.Enum conversion.
func NodePoolUpgradeSettings_ToProto(mapCtx *direct.MapContext, in *krm.NodePoolUpgradeSettings) *pb.NodePool_UpgradeSettings {
	if in == nil {
		return nil
	}
	out := &pb.NodePool_UpgradeSettings{}
	out.MaxSurge = direct.ValueOf(in.MaxSurge)
	out.MaxUnavailable = direct.ValueOf(in.MaxUnavailable)
	if oneof := NodePoolUpgradeSettings_Strategy_ToProto(mapCtx, in.Strategy); oneof != nil {
		out.Strategy = oneof
	}
	out.BlueGreenSettings = NodePoolBlueGreenSettings_ToProto(mapCtx, in.BlueGreenSettings)
	return out
}

// NodePool_UpgradeSettings_FromProto maps NodePool_UpgradeSettings from proto.
// Handwritten due to direct.Enum conversion and type mismatch.
func NodePool_UpgradeSettings_FromProto(mapCtx *direct.MapContext, in *pb.NodePool_UpgradeSettings) *krm.NodePool_UpgradeSettings {
	if in == nil {
		return nil
	}
	out := &krm.NodePool_UpgradeSettings{}
	out.MaxSurge = direct.LazyPtr(int(in.GetMaxSurge()))
	out.MaxUnavailable = direct.LazyPtr(int(in.GetMaxUnavailable()))
	out.Strategy = direct.Enum_FromProto(mapCtx, in.GetStrategy())
	out.NodePool_UpdateConfig_BlueGreenSettings = NodePool_UpdateConfig_BlueGreenSettings_FromProto(mapCtx, in.GetBlueGreenSettings())
	return out
}

// NodePool_UpgradeSettings_ToProto maps NodePool_UpgradeSettings to proto.
// Handwritten due to direct.Enum conversion and type mismatch.
func NodePool_UpgradeSettings_ToProto(mapCtx *direct.MapContext, in *krm.NodePool_UpgradeSettings) *pb.NodePool_UpgradeSettings {
	if in == nil {
		return nil
	}
	out := &pb.NodePool_UpgradeSettings{}
	out.MaxSurge = int32(direct.ValueOf(in.MaxSurge))
	out.MaxUnavailable = int32(direct.ValueOf(in.MaxUnavailable))
	if oneof := NodePoolUpgradeSettings_Strategy_ToProto(mapCtx, in.Strategy); oneof != nil {
		out.Strategy = oneof
	}
	out.BlueGreenSettings = NodePool_UpdateConfig_BlueGreenSettings_ToProto(mapCtx, in.NodePool_UpdateConfig_BlueGreenSettings)
	return out
}

// NodePool_PlacementPolicy_FromProto maps NodePool_PlacementPolicy from proto.
// Handwritten because type is a required string in KRM but Enum in proto.
func NodePool_PlacementPolicy_FromProto(mapCtx *direct.MapContext, in *pb.NodePool_PlacementPolicy) *krm.NodePool_PlacementPolicy {
	if in == nil {
		return nil
	}
	out := &krm.NodePool_PlacementPolicy{}
	out.Type = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetType()))
	out.TpuTopology = direct.LazyPtr(in.GetTpuTopology())
	if in.GetPolicyName() != "" {
		out.PolicyNameRef = &computev1beta1.ComputeResourcePolicyRef{External: in.GetPolicyName()}
	}
	return out
}

// NodePool_PlacementPolicy_ToProto maps NodePool_PlacementPolicy to proto.
// Handwritten because type is a required string in KRM but Enum in proto.
func NodePool_PlacementPolicy_ToProto(mapCtx *direct.MapContext, in *krm.NodePool_PlacementPolicy) *pb.NodePool_PlacementPolicy {
	if in == nil {
		return nil
	}
	out := &pb.NodePool_PlacementPolicy{}
	out.Type = direct.Enum_ToProto[pb.NodePool_PlacementPolicy_Type](mapCtx, &in.Type)
	out.TpuTopology = direct.ValueOf(in.TpuTopology)
	if in.PolicyNameRef != nil {
		out.PolicyName = in.PolicyNameRef.External
	}
	return out
}

// StandardRolloutPolicy_FromProto maps StandardRolloutPolicy from proto.
func StandardRolloutPolicy_FromProto(mapCtx *direct.MapContext, in *pb.BlueGreenSettings_StandardRolloutPolicy) *krm.StandardRolloutPolicy {
	if in == nil {
		return nil
	}
	out := &krm.StandardRolloutPolicy{}
	if in.GetBatchNodeCount() != 0 {
		out.BatchNodeCount = direct.LazyPtr(int(in.GetBatchNodeCount()))
	}
	if in.GetBatchPercentage() != 0 {
		out.BatchPercentage = direct.LazyPtr(float64(in.GetBatchPercentage()))
	}
	out.BatchSoakDuration = direct.StringDuration_FromProto(mapCtx, in.GetBatchSoakDuration())
	return out
}

// StandardRolloutPolicy_ToProto maps StandardRolloutPolicy to proto.
func StandardRolloutPolicy_ToProto(mapCtx *direct.MapContext, in *krm.StandardRolloutPolicy) *pb.BlueGreenSettings_StandardRolloutPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BlueGreenSettings_StandardRolloutPolicy{}
	if in.BatchNodeCount != nil {
		out.UpdateBatchSize = &pb.BlueGreenSettings_StandardRolloutPolicy_BatchNodeCount{
			BatchNodeCount: int32(*in.BatchNodeCount),
		}
	}
	if in.BatchPercentage != nil {
		out.UpdateBatchSize = &pb.BlueGreenSettings_StandardRolloutPolicy_BatchPercentage{
			BatchPercentage: float32(*in.BatchPercentage),
		}
	}
	out.BatchSoakDuration = direct.StringDuration_ToProto(mapCtx, in.BatchSoakDuration)
	return out
}

// ContainerNodePoolSpec_FromProto maps ContainerNodePoolSpec from proto.
// Handwritten because max_pods_per_node is flattened from a nested MaxPodsConstraint message,
// and parent-level fields like clusterRef are resolved by the controller.
func ContainerNodePoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.NodePool) *krm.ContainerNodePoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContainerNodePoolSpec{}
	out.Autoscaling = NodePoolAutoscaling_FromProto(mapCtx, in.GetAutoscaling())
	out.InitialNodeCount = direct.LazyPtr(in.GetInitialNodeCount())
	out.Management = NodePoolManagement_FromProto(mapCtx, in.GetManagement())
	if in.GetMaxPodsConstraint() != nil {
		out.MaxPodsPerNode = direct.LazyPtr(int(in.GetMaxPodsConstraint().GetMaxPodsPerNode()))
	}
	out.NetworkConfig = NodeNetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	out.NodeConfig = NodePoolNodeConfig_FromProto(mapCtx, in.GetConfig())
	out.NodeLocations = in.GetLocations()
	out.PlacementPolicy = NodePool_PlacementPolicy_FromProto(mapCtx, in.GetPlacementPolicy())
	out.UpgradeSettings = NodePoolUpgradeSettings_FromProto(mapCtx, in.GetUpgradeSettings())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}

// ContainerNodePoolSpec_ToProto maps ContainerNodePoolSpec to proto.
// Handwritten because max_pods_per_node is flattened to a nested MaxPodsConstraint message,
// and parent-level fields like clusterRef are resolved by the controller.
func ContainerNodePoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContainerNodePoolSpec) *pb.NodePool {
	if in == nil {
		return nil
	}
	out := &pb.NodePool{}
	out.Autoscaling = NodePoolAutoscaling_ToProto(mapCtx, in.Autoscaling)
	out.InitialNodeCount = direct.ValueOf(in.InitialNodeCount)
	out.Management = NodePoolManagement_ToProto(mapCtx, in.Management)
	if in.MaxPodsPerNode != nil {
		out.MaxPodsConstraint = &pb.MaxPodsConstraint{
			MaxPodsPerNode: int64(direct.ValueOf(in.MaxPodsPerNode)),
		}
	}
	out.NetworkConfig = NodeNetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	out.Config = NodePoolNodeConfig_ToProto(mapCtx, in.NodeConfig)
	out.Locations = in.NodeLocations
	out.PlacementPolicy = NodePool_PlacementPolicy_ToProto(mapCtx, in.PlacementPolicy)
	out.UpgradeSettings = NodePoolUpgradeSettings_ToProto(mapCtx, in.UpgradeSettings)
	out.Version = direct.ValueOf(in.Version)
	return out
}
