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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/iamrefs"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
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
	out.QueuedProvisioning = NodePoolQueuedProvisioning_FromProto(mapCtx, in.GetQueuedProvisioning())
	out.UpgradeSettings = NodePoolUpgradeSettings_FromProto(mapCtx, in.GetUpgradeSettings())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}

// ContainerNodePoolSpec_ToProto maps ContainerNodePoolSpec to proto.
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
	out.QueuedProvisioning = NodePoolQueuedProvisioning_ToProto(mapCtx, in.QueuedProvisioning)
	out.UpgradeSettings = NodePoolUpgradeSettings_ToProto(mapCtx, in.UpgradeSettings)
	out.Version = direct.ValueOf(in.Version)
	return out
}

func NodeNetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeNetworkConfig) *krm.NodeNetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeNetworkConfig{}
	out.CreatePodRange = direct.LazyPtr(in.GetCreatePodRange())
	out.PodRange = direct.LazyPtr(in.GetPodRange())
	out.PodIpv4CidrBlock = direct.LazyPtr(in.GetPodIpv4CidrBlock())
	out.EnablePrivateNodes = in.EnablePrivateNodes
	out.PodCidrOverprovisionConfig = PodCIDROverprovisionConfig_FromProto(mapCtx, in.GetPodCidrOverprovisionConfig())
	out.AdditionalNodeNetworkConfigs = direct.Slice_FromProto(mapCtx, in.AdditionalNodeNetworkConfigs, AdditionalNodeNetworkConfig_FromProto)
	out.AdditionalPodNetworkConfigs = direct.Slice_FromProto(mapCtx, in.AdditionalPodNetworkConfigs, AdditionalPodNetworkConfig_FromProto)
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &computev1beta1.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	return out
}

func NodeNetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeNetworkConfig) *pb.NodeNetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeNetworkConfig{}
	out.CreatePodRange = direct.ValueOf(in.CreatePodRange)
	out.PodRange = direct.ValueOf(in.PodRange)
	out.PodIpv4CidrBlock = direct.ValueOf(in.PodIpv4CidrBlock)
	out.EnablePrivateNodes = in.EnablePrivateNodes
	out.PodCidrOverprovisionConfig = PodCIDROverprovisionConfig_ToProto(mapCtx, in.PodCidrOverprovisionConfig)
	out.AdditionalNodeNetworkConfigs = direct.Slice_ToProto(mapCtx, in.AdditionalNodeNetworkConfigs, AdditionalNodeNetworkConfig_ToProto)
	out.AdditionalPodNetworkConfigs = direct.Slice_ToProto(mapCtx, in.AdditionalPodNetworkConfigs, AdditionalPodNetworkConfig_ToProto)
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	return out
}

func CertificateAuthorityDomainConfig_FromProto(mapCtx *direct.MapContext, in *pb.ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig) *krm.CertificateAuthorityDomainConfig {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthorityDomainConfig{}
	out.Fqdns = in.Fqdns
	out.GCPSecretManagerCertificateConfig = GCPSecretManagerCertificateConfig_FromProto(mapCtx, in.GetGcpSecretManagerCertificateConfig())
	return out
}

func CertificateAuthorityDomainConfig_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthorityDomainConfig) *pb.ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig {
	if in == nil {
		return nil
	}
	out := &pb.ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig{}
	out.Fqdns = in.Fqdns
	if in.GCPSecretManagerCertificateConfig != nil {
		out.CertificateConfig = &pb.ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig_GcpSecretManagerCertificateConfig{
			GcpSecretManagerCertificateConfig: GCPSecretManagerCertificateConfig_ToProto(mapCtx, in.GCPSecretManagerCertificateConfig),
		}
	}
	return out
}

func NodeConfig_WorkloadMetadataConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadMetadataConfig) *krm.NodeConfig_WorkloadMetadataConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeConfig_WorkloadMetadataConfig{}
	out.Mode = direct.Enum_FromProto[pb.WorkloadMetadataConfig_Mode](mapCtx, in.GetMode())
	return out
}

func NodeConfig_WorkloadMetadataConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeConfig_WorkloadMetadataConfig) *pb.WorkloadMetadataConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadMetadataConfig{}
	out.Mode = direct.Enum_ToProto[pb.WorkloadMetadataConfig_Mode](mapCtx, in.Mode)
	return out
}

func VirtualNic_FromProto(mapCtx *direct.MapContext, in *pb.VirtualNIC) *krm.VirtualNic {
	if in == nil {
		return nil
	}
	out := &krm.VirtualNic{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}

func VirtualNic_ToProto(mapCtx *direct.MapContext, in *krm.VirtualNic) *pb.VirtualNIC {
	if in == nil {
		return nil
	}
	out := &pb.VirtualNIC{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}

func NodeConfig_LocalNvmeSsdBlockConfig_FromProto(mapCtx *direct.MapContext, in *pb.LocalNvmeSsdBlockConfig) *krm.NodeConfig_LocalNvmeSsdBlockConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeConfig_LocalNvmeSsdBlockConfig{}
	out.LocalSsdCount = direct.LazyPtr(int(in.GetLocalSsdCount()))
	return out
}

func NodeConfig_LocalNvmeSsdBlockConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeConfig_LocalNvmeSsdBlockConfig) *pb.LocalNvmeSsdBlockConfig {
	if in == nil {
		return nil
	}
	out := &pb.LocalNvmeSsdBlockConfig{}
	out.LocalSsdCount = int32(direct.ValueOf(in.LocalSsdCount))
	return out
}

func NodeAffinity_FromProto(mapCtx *direct.MapContext, in *pb.SoleTenantConfig_NodeAffinity) *krm.NodeAffinity {
	if in == nil {
		return nil
	}
	out := &krm.NodeAffinity{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Operator = direct.Enum_FromProto[pb.SoleTenantConfig_NodeAffinity_Operator](mapCtx, in.GetOperator())
	out.Values = in.Values
	return out
}

func NodeAffinity_ToProto(mapCtx *direct.MapContext, in *krm.NodeAffinity) *pb.SoleTenantConfig_NodeAffinity {
	if in == nil {
		return nil
	}
	out := &pb.SoleTenantConfig_NodeAffinity{}
	out.Key = direct.ValueOf(in.Key)
	out.Operator = direct.Enum_ToProto[pb.SoleTenantConfig_NodeAffinity_Operator](mapCtx, in.Operator)
	out.Values = in.Values
	return out
}

func NodeConfig_SoleTenantConfig_FromProto(mapCtx *direct.MapContext, in *pb.SoleTenantConfig) *krm.NodeConfig_SoleTenantConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeConfig_SoleTenantConfig{}
	out.NodeAffinity = direct.Slice_FromProto(mapCtx, in.NodeAffinities, NodeAffinity_FromProto)
	return out
}

func NodeConfig_SoleTenantConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeConfig_SoleTenantConfig) *pb.SoleTenantConfig {
	if in == nil {
		return nil
	}
	out := &pb.SoleTenantConfig{}
	out.NodeAffinities = direct.Slice_ToProto(mapCtx, in.NodeAffinity, NodeAffinity_ToProto)
	return out
}

// NodePoolNodeConfig_FromProto maps NodePoolNodeConfig from proto.
// Handwritten to map directly to kmsv1beta1 and iamrefs.
func NodePoolNodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeConfig) *krm.NodePoolNodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodePoolNodeConfig{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.DiskSizeGb = direct.LazyPtr(in.GetDiskSizeGb())
	out.OauthScopes = in.OauthScopes
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &iamrefs.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.Metadata = in.Metadata
	out.ImageType = direct.LazyPtr(in.GetImageType())
	out.Labels = in.Labels
	out.LocalSsdCount = direct.LazyPtr(in.GetLocalSsdCount())
	out.Tags = in.Tags
	out.Preemptible = direct.LazyPtr(in.GetPreemptible())
	out.AcceleratorConfig = direct.Slice_FromProto(mapCtx, in.Accelerators, AcceleratorConfig_FromProto)
	out.DiskType = direct.LazyPtr(in.GetDiskType())
	out.MinCPUPlatform = direct.LazyPtr(in.GetMinCpuPlatform())
	out.NodeConfig_WorkloadMetadataConfig = NodeConfig_WorkloadMetadataConfig_FromProto(mapCtx, in.GetWorkloadMetadataConfig())
	out.Taint = direct.Slice_FromProto(mapCtx, in.GetTaints(), NodeTaint_FromProto)
	out.SandboxConfig = SandboxConfig_FromProto(mapCtx, in.GetSandboxConfig())
	if in.GetNodeGroup() != "" {
		out.NodeGroupRef = &computev1beta1.ComputeNodeGroupRef{External: in.GetNodeGroup()}
	}
	out.ReservationAffinity = ReservationAffinity_FromProto(mapCtx, in.GetReservationAffinity())
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.LinuxNodeConfig = LinuxNodeConfig_FromProto(mapCtx, in.GetLinuxNodeConfig())
	out.KubeletConfig = KubeletConfig_FromProto(mapCtx, in.GetKubeletConfig())
	if in.GetBootDiskKmsKey() != "" {
		out.BootDiskKMSCryptoKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: in.GetBootDiskKmsKey()}
	}
	out.GcfsConfig = GcfsConfig_FromProto(mapCtx, in.GetGcfsConfig())
	out.AdvancedMachineFeatures = NodeConfig_AdvancedMachineFeatures_FromProto(mapCtx, in.GetAdvancedMachineFeatures())
	out.VirtualNic = VirtualNic_FromProto(mapCtx, in.GetGvnic())
	out.Spot = direct.LazyPtr(in.GetSpot())
	out.ConfidentialNodes = ConfidentialNodes_FromProto(mapCtx, in.GetConfidentialNodes())
	out.FastSocket = FastSocket_FromProto(mapCtx, in.GetFastSocket())
	out.ResourceLabels = in.ResourceLabels
	if in.GetLoggingConfig().GetVariantConfig() != nil {
		variant := in.GetLoggingConfig().GetVariantConfig().GetVariant().String()
		if variant != "" && variant != "VARIANT_UNSPECIFIED" {
			out.LoggingVariant = &variant
		}
	}
	out.WindowsNodeConfig = WindowsNodeConfig_FromProto(mapCtx, in.GetWindowsNodeConfig())
	out.NodeConfig_LocalNvmeSsdBlockConfig = NodeConfig_LocalNvmeSsdBlockConfig_FromProto(mapCtx, in.GetLocalNvmeSsdBlockConfig())
	out.EphemeralStorageLocalSsdConfig = EphemeralStorageLocalSsdConfig_FromProto(mapCtx, in.GetEphemeralStorageLocalSsdConfig())
	out.NodeConfig_SoleTenantConfig = NodeConfig_SoleTenantConfig_FromProto(mapCtx, in.GetSoleTenantConfig())
	out.ContainerdConfig = ContainerdConfig_FromProto(mapCtx, in.GetContainerdConfig())
	out.ResourceManagerTags = map_string_string_FromProto(mapCtx, in.GetResourceManagerTags())
	return out
}

// NodePoolNodeConfig_ToProto maps NodePoolNodeConfig to proto.
// Handwritten to map directly to kmsv1beta1 and iamrefs.
func NodePoolNodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodePoolNodeConfig) *pb.NodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeConfig{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGb)
	out.OauthScopes = in.OauthScopes
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.Metadata = in.Metadata
	out.ImageType = direct.ValueOf(in.ImageType)
	out.Labels = in.Labels
	out.LocalSsdCount = direct.ValueOf(in.LocalSsdCount)
	out.Tags = in.Tags
	out.Preemptible = direct.ValueOf(in.Preemptible)
	out.Accelerators = direct.Slice_ToProto(mapCtx, in.AcceleratorConfig, AcceleratorConfig_ToProto)
	out.DiskType = direct.ValueOf(in.DiskType)
	out.MinCpuPlatform = direct.ValueOf(in.MinCPUPlatform)
	out.WorkloadMetadataConfig = NodeConfig_WorkloadMetadataConfig_ToProto(mapCtx, in.NodeConfig_WorkloadMetadataConfig)
	out.Taints = direct.Slice_ToProto(mapCtx, in.Taint, NodeTaint_ToProto)
	out.SandboxConfig = SandboxConfig_ToProto(mapCtx, in.SandboxConfig)
	if in.NodeGroupRef != nil {
		out.NodeGroup = in.NodeGroupRef.External
	}
	out.ReservationAffinity = ReservationAffinity_ToProto(mapCtx, in.ReservationAffinity)
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.LinuxNodeConfig = LinuxNodeConfig_ToProto(mapCtx, in.LinuxNodeConfig)
	out.KubeletConfig = KubeletConfig_ToProto(mapCtx, in.KubeletConfig)
	if in.BootDiskKMSCryptoKeyRef != nil {
		out.BootDiskKmsKey = in.BootDiskKMSCryptoKeyRef.External
	}
	out.GcfsConfig = GcfsConfig_ToProto(mapCtx, in.GcfsConfig)
	out.AdvancedMachineFeatures = NodeConfig_AdvancedMachineFeatures_ToProto(mapCtx, in.AdvancedMachineFeatures)
	out.Gvnic = VirtualNic_ToProto(mapCtx, in.VirtualNic)
	out.Spot = direct.ValueOf(in.Spot)
	out.ConfidentialNodes = ConfidentialNodes_ToProto(mapCtx, in.ConfidentialNodes)
	out.FastSocket = FastSocket_ToProto(mapCtx, in.FastSocket)
	out.ResourceLabels = in.ResourceLabels
	if in.LoggingVariant != nil {
		out.LoggingConfig = &pb.NodePoolLoggingConfig{
			VariantConfig: &pb.LoggingVariantConfig{
				Variant: pb.LoggingVariantConfig_Variant(pb.LoggingVariantConfig_Variant_value[*in.LoggingVariant]),
			},
		}
	}
	out.WindowsNodeConfig = WindowsNodeConfig_ToProto(mapCtx, in.WindowsNodeConfig)
	out.LocalNvmeSsdBlockConfig = NodeConfig_LocalNvmeSsdBlockConfig_ToProto(mapCtx, in.NodeConfig_LocalNvmeSsdBlockConfig)
	out.EphemeralStorageLocalSsdConfig = EphemeralStorageLocalSsdConfig_ToProto(mapCtx, in.EphemeralStorageLocalSsdConfig)
	out.SoleTenantConfig = NodeConfig_SoleTenantConfig_ToProto(mapCtx, in.NodeConfig_SoleTenantConfig)
	out.ContainerdConfig = ContainerdConfig_ToProto(mapCtx, in.ContainerdConfig)
	out.ResourceManagerTags = map_string_string_ToProto(mapCtx, in.ResourceManagerTags)
	return out
}

// WindowsNodeConfig_FromProto maps pb.WindowsNodeConfig to krm.WindowsNodeConfig.
// Supports both short names (OS_2022, OS_2019) and full proto enum names.
func WindowsNodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.WindowsNodeConfig) *krm.WindowsNodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.WindowsNodeConfig{}
	switch in.GetOsVersion() {
	case pb.WindowsNodeConfig_OS_VERSION_LTSC2022:
		out.OSVersion = direct.LazyPtr("OS_2022")
	case pb.WindowsNodeConfig_OS_VERSION_LTSC2019:
		out.OSVersion = direct.LazyPtr("OS_2019")
	case pb.WindowsNodeConfig_OS_VERSION_UNSPECIFIED:
		out.OSVersion = nil
	default:
		out.OSVersion = direct.Enum_FromProto(mapCtx, in.GetOsVersion())
	}
	return out
}

// WindowsNodeConfig_ToProto maps krm.WindowsNodeConfig to pb.WindowsNodeConfig.
// Supports both short names (OS_2022, OS_2019) and full proto enum names.
func WindowsNodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.WindowsNodeConfig) *pb.WindowsNodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.WindowsNodeConfig{}
	if in.OSVersion != nil {
		switch *in.OSVersion {
		case "OS_2022", "OS_VERSION_LTSC2022":
			out.OsVersion = pb.WindowsNodeConfig_OS_VERSION_LTSC2022
		case "OS_2019", "OS_VERSION_LTSC2019":
			out.OsVersion = pb.WindowsNodeConfig_OS_VERSION_LTSC2019
		case "OS_VERSION_UNSPECIFIED":
			out.OsVersion = pb.WindowsNodeConfig_OS_VERSION_UNSPECIFIED
		default:
			out.OsVersion = direct.Enum_ToProto[pb.WindowsNodeConfig_OSVersion](mapCtx, in.OSVersion)
		}
	}
	return out
}

func NodepoolObservedStateStatus_FromProto(mapCtx *direct.MapContext, in *pb.NodePool) *krm.NodepoolObservedStateStatus {
	if in == nil {
		return nil
	}
	out := &krm.NodepoolObservedStateStatus{}
	if in.GetVersion() != "" {
		out.Version = direct.LazyPtr(in.GetVersion())
	}
	if in.GetConfig() != nil && len(in.GetConfig().GetTaints()) > 0 {
		out.NodeConfig = &krm.NodePoolNodeConfigObservedState{
			Taint: direct.Slice_FromProto(mapCtx, in.GetConfig().GetTaints(), NodeTaint_FromProto),
		}
	}
	return out
}

func NodepoolObservedStateStatus_ToProto(mapCtx *direct.MapContext, in *krm.NodepoolObservedStateStatus) *pb.NodePool {
	if in == nil {
		return nil
	}
	out := &pb.NodePool{}
	out.Version = direct.ValueOf(in.Version)
	if in.NodeConfig != nil && len(in.NodeConfig.Taint) > 0 {
		out.Config = &pb.NodeConfig{
			Taints: direct.Slice_ToProto(mapCtx, in.NodeConfig.Taint, NodeTaint_ToProto),
		}
	}
	return out
}

func LinuxNodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.LinuxNodeConfig) *krm.LinuxNodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.LinuxNodeConfig{}
	out.Sysctls = in.Sysctls
	out.CgroupMode = direct.Enum_FromProto(mapCtx, in.GetCgroupMode())
	out.SwapConfig = LinuxNodeConfig_SwapConfig_FromProto(mapCtx, in.GetSwapConfig())
	return out
}

func LinuxNodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.LinuxNodeConfig) *pb.LinuxNodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.LinuxNodeConfig{}
	out.Sysctls = in.Sysctls
	out.CgroupMode = direct.Enum_ToProto[pb.LinuxNodeConfig_CgroupMode](mapCtx, in.CgroupMode)
	out.SwapConfig = LinuxNodeConfig_SwapConfig_ToProto(mapCtx, in.SwapConfig)
	return out
}

func LinuxNodeConfig_SwapConfig_FromProto(mapCtx *direct.MapContext, in *pb.LinuxNodeConfig_SwapConfig) *krm.LinuxNodeConfig_SwapConfig {
	if in == nil {
		return nil
	}
	out := &krm.LinuxNodeConfig_SwapConfig{}
	out.Enabled = in.Enabled
	if in.EncryptionConfig != nil {
		out.EncryptionConfig = &krm.SwapConfig_EncryptionConfig{
			Disabled: in.EncryptionConfig.Disabled,
		}
	}
	switch p := in.PerformanceProfile.(type) {
	case *pb.LinuxNodeConfig_SwapConfig_BootDiskProfile_:
		if bp := p.BootDiskProfile; bp != nil {
			out.BootDiskProfile = &krm.SwapConfig_BootDiskProfile{}
			switch s := bp.SwapSize.(type) {
			case *pb.LinuxNodeConfig_SwapConfig_BootDiskProfile_SwapSizeGib:
				out.BootDiskProfile.SwapSizeGib = direct.LazyPtr(int(s.SwapSizeGib))
			case *pb.LinuxNodeConfig_SwapConfig_BootDiskProfile_SwapSizePercent:
				out.BootDiskProfile.SwapSizePercent = direct.LazyPtr(int(s.SwapSizePercent))
			}
		}
	case *pb.LinuxNodeConfig_SwapConfig_EphemeralLocalSsdProfile_:
		if ep := p.EphemeralLocalSsdProfile; ep != nil {
			out.EphemeralLocalSsdProfile = &krm.SwapConfig_EphemeralLocalSsdProfile{}
			switch s := ep.SwapSize.(type) {
			case *pb.LinuxNodeConfig_SwapConfig_EphemeralLocalSsdProfile_SwapSizeGib:
				out.EphemeralLocalSsdProfile.SwapSizeGib = direct.LazyPtr(int(s.SwapSizeGib))
			case *pb.LinuxNodeConfig_SwapConfig_EphemeralLocalSsdProfile_SwapSizePercent:
				out.EphemeralLocalSsdProfile.SwapSizePercent = direct.LazyPtr(int(s.SwapSizePercent))
			}
		}
	case *pb.LinuxNodeConfig_SwapConfig_DedicatedLocalSsdProfile_:
		if dp := p.DedicatedLocalSsdProfile; dp != nil {
			out.DedicatedLocalSsdProfile = &krm.SwapConfig_DedicatedLocalSsdProfile{
				DiskCount: direct.LazyPtr(int(dp.DiskCount)),
			}
		}
	}
	return out
}

func LinuxNodeConfig_SwapConfig_ToProto(mapCtx *direct.MapContext, in *krm.LinuxNodeConfig_SwapConfig) *pb.LinuxNodeConfig_SwapConfig {
	if in == nil {
		return nil
	}
	out := &pb.LinuxNodeConfig_SwapConfig{}
	out.Enabled = in.Enabled
	if in.EncryptionConfig != nil {
		out.EncryptionConfig = &pb.LinuxNodeConfig_SwapConfig_EncryptionConfig{
			Disabled: in.EncryptionConfig.Disabled,
		}
	}
	if bp := in.BootDiskProfile; bp != nil {
		p := &pb.LinuxNodeConfig_SwapConfig_BootDiskProfile{}
		if bp.SwapSizeGib != nil {
			p.SwapSize = &pb.LinuxNodeConfig_SwapConfig_BootDiskProfile_SwapSizeGib{
				SwapSizeGib: int64(*bp.SwapSizeGib),
			}
		} else if bp.SwapSizePercent != nil {
			p.SwapSize = &pb.LinuxNodeConfig_SwapConfig_BootDiskProfile_SwapSizePercent{
				SwapSizePercent: int32(*bp.SwapSizePercent),
			}
		}
		out.PerformanceProfile = &pb.LinuxNodeConfig_SwapConfig_BootDiskProfile_{
			BootDiskProfile: p,
		}
	} else if ep := in.EphemeralLocalSsdProfile; ep != nil {
		p := &pb.LinuxNodeConfig_SwapConfig_EphemeralLocalSsdProfile{}
		if ep.SwapSizeGib != nil {
			p.SwapSize = &pb.LinuxNodeConfig_SwapConfig_EphemeralLocalSsdProfile_SwapSizeGib{
				SwapSizeGib: int64(*ep.SwapSizeGib),
			}
		} else if ep.SwapSizePercent != nil {
			p.SwapSize = &pb.LinuxNodeConfig_SwapConfig_EphemeralLocalSsdProfile_SwapSizePercent{
				SwapSizePercent: int32(*ep.SwapSizePercent),
			}
		}
		out.PerformanceProfile = &pb.LinuxNodeConfig_SwapConfig_EphemeralLocalSsdProfile_{
			EphemeralLocalSsdProfile: p,
		}
	} else if dp := in.DedicatedLocalSsdProfile; dp != nil {
		out.PerformanceProfile = &pb.LinuxNodeConfig_SwapConfig_DedicatedLocalSsdProfile_{
			DedicatedLocalSsdProfile: &pb.LinuxNodeConfig_SwapConfig_DedicatedLocalSsdProfile{
				DiskCount: int64(direct.ValueOf(dp.DiskCount)),
			},
		}
	}
	return out
}

func ContainerdConfig_FromProto(mapCtx *direct.MapContext, in *pb.ContainerdConfig) *krm.ContainerdConfig {
	if in == nil {
		return nil
	}
	out := &krm.ContainerdConfig{}
	out.PrivateRegistryAccessConfig = PrivateRegistryAccessConfig_FromProto(mapCtx, in.GetPrivateRegistryAccessConfig())
	out.WritableCgroups = WritableCgroups_FromProto(mapCtx, in.GetWritableCgroups())
	out.RegistryHosts = direct.Slice_FromProto(mapCtx, in.GetRegistryHosts(), RegistryHosts_FromProto)
	return out
}

func ContainerdConfig_ToProto(mapCtx *direct.MapContext, in *krm.ContainerdConfig) *pb.ContainerdConfig {
	if in == nil {
		return nil
	}
	out := &pb.ContainerdConfig{}
	out.PrivateRegistryAccessConfig = PrivateRegistryAccessConfig_ToProto(mapCtx, in.PrivateRegistryAccessConfig)
	out.WritableCgroups = WritableCgroups_ToProto(mapCtx, in.WritableCgroups)
	out.RegistryHosts = direct.Slice_ToProto(mapCtx, in.RegistryHosts, RegistryHosts_ToProto)
	return out
}

func WritableCgroups_FromProto(mapCtx *direct.MapContext, in *pb.ContainerdConfig_WritableCgroups) *krm.WritableCgroups {
	if in == nil {
		return nil
	}
	out := &krm.WritableCgroups{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}

func WritableCgroups_ToProto(mapCtx *direct.MapContext, in *krm.WritableCgroups) *pb.ContainerdConfig_WritableCgroups {
	if in == nil {
		return nil
	}
	out := &pb.ContainerdConfig_WritableCgroups{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}

func GCPSecretManagerCertificateConfig_FromProto(mapCtx *direct.MapContext, in *pb.ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig_GCPSecretManagerCertificateConfig) *krm.GCPSecretManagerCertificateConfig {
	if in == nil {
		return nil
	}
	out := &krm.GCPSecretManagerCertificateConfig{}
	if in.GetSecretUri() != "" {
		out.SecretRef = &secretmanagerv1beta1.SecretVersionRef{
			External: in.GetSecretUri(),
		}
	}
	return out
}

func GCPSecretManagerCertificateConfig_ToProto(mapCtx *direct.MapContext, in *krm.GCPSecretManagerCertificateConfig) *pb.ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig_GCPSecretManagerCertificateConfig {
	if in == nil {
		return nil
	}
	out := &pb.ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig_GCPSecretManagerCertificateConfig{}
	if in.SecretRef != nil {
		out.SecretUri = in.SecretRef.External
	}
	return out
}

func RegistryHosts_FromProto(mapCtx *direct.MapContext, in *pb.ContainerdConfig_RegistryHostConfig) *krm.RegistryHosts {
	if in == nil {
		return nil
	}
	out := &krm.RegistryHosts{}
	out.Server = direct.LazyPtr(in.GetServer())
	out.Hosts = direct.Slice_FromProto(mapCtx, in.GetHosts(), RegistryHostsConfig_FromProto)
	return out
}

func RegistryHosts_ToProto(mapCtx *direct.MapContext, in *krm.RegistryHosts) *pb.ContainerdConfig_RegistryHostConfig {
	if in == nil {
		return nil
	}
	out := &pb.ContainerdConfig_RegistryHostConfig{}
	out.Server = direct.ValueOf(in.Server)
	out.Hosts = direct.Slice_ToProto(mapCtx, in.Hosts, RegistryHostsConfig_ToProto)
	return out
}

func RegistryHostsConfig_FromProto(mapCtx *direct.MapContext, in *pb.ContainerdConfig_RegistryHostConfig_HostConfig) *krm.RegistryHostsConfig {
	if in == nil {
		return nil
	}
	out := &krm.RegistryHostsConfig{}
	out.Host = direct.LazyPtr(in.GetHost())
	for _, cap := range in.GetCapabilities() {
		if s := direct.Enum_FromProto(mapCtx, cap); s != nil {
			out.Capabilities = append(out.Capabilities, *s)
		}
	}
	out.OverridePath = direct.LazyPtr(in.GetOverridePath())
	out.DialTimeout = direct.Duration_FromProto(mapCtx, in.GetDialTimeout())
	out.Header = direct.Slice_FromProto(mapCtx, in.GetHeader(), RegistryHeader_FromProto)
	out.Ca = direct.Slice_FromProto(mapCtx, in.GetCa(), RegistryCA_FromProto)
	out.Client = direct.Slice_FromProto(mapCtx, in.GetClient(), RegistryClient_FromProto)
	return out
}

func RegistryHostsConfig_ToProto(mapCtx *direct.MapContext, in *krm.RegistryHostsConfig) *pb.ContainerdConfig_RegistryHostConfig_HostConfig {
	if in == nil {
		return nil
	}
	out := &pb.ContainerdConfig_RegistryHostConfig_HostConfig{}
	out.Host = direct.ValueOf(in.Host)
	for _, cap := range in.Capabilities {
		out.Capabilities = append(out.Capabilities, direct.Enum_ToProto[pb.ContainerdConfig_RegistryHostConfig_HostCapability](mapCtx, direct.LazyPtr(cap)))
	}
	out.OverridePath = direct.ValueOf(in.OverridePath)
	out.DialTimeout = direct.StringDuration_ToProto(mapCtx, in.DialTimeout)
	out.Header = direct.Slice_ToProto(mapCtx, in.Header, RegistryHeader_ToProto)
	out.Ca = direct.Slice_ToProto(mapCtx, in.Ca, RegistryCA_ToProto)
	out.Client = direct.Slice_ToProto(mapCtx, in.Client, RegistryClient_ToProto)
	return out
}

func RegistryHeader_FromProto(mapCtx *direct.MapContext, in *pb.ContainerdConfig_RegistryHostConfig_RegistryHeader) *krm.RegistryHeader {
	if in == nil {
		return nil
	}
	out := &krm.RegistryHeader{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Value = in.GetValue()
	return out
}

func RegistryHeader_ToProto(mapCtx *direct.MapContext, in *krm.RegistryHeader) *pb.ContainerdConfig_RegistryHostConfig_RegistryHeader {
	if in == nil {
		return nil
	}
	out := &pb.ContainerdConfig_RegistryHostConfig_RegistryHeader{}
	out.Key = direct.ValueOf(in.Key)
	out.Value = in.Value
	return out
}

func RegistryCA_FromProto(mapCtx *direct.MapContext, in *pb.ContainerdConfig_RegistryHostConfig_CertificateConfig) *krm.RegistryCA {
	if in == nil {
		return nil
	}
	out := &krm.RegistryCA{}
	if uri := in.GetGcpSecretManagerSecretUri(); uri != "" {
		out.SecretRef = &secretmanagerv1beta1.SecretVersionRef{
			External: uri,
		}
	}
	return out
}

func RegistryCA_ToProto(mapCtx *direct.MapContext, in *krm.RegistryCA) *pb.ContainerdConfig_RegistryHostConfig_CertificateConfig {
	if in == nil || in.SecretRef == nil || in.SecretRef.External == "" {
		return nil
	}
	return &pb.ContainerdConfig_RegistryHostConfig_CertificateConfig{
		Certificate: &pb.ContainerdConfig_RegistryHostConfig_CertificateConfig_GcpSecretManagerSecretUri{
			GcpSecretManagerSecretUri: in.SecretRef.External,
		},
	}
}

func RegistryClient_FromProto(mapCtx *direct.MapContext, in *pb.ContainerdConfig_RegistryHostConfig_CertificateConfigPair) *krm.RegistryClient {
	if in == nil {
		return nil
	}
	out := &krm.RegistryClient{}
	if in.Cert != nil {
		if uri := in.Cert.GetGcpSecretManagerSecretUri(); uri != "" {
			out.Cert = &krm.RegistryClientCert{
				SecretRef: &secretmanagerv1beta1.SecretVersionRef{
					External: uri,
				},
			}
		}
	}
	if in.Key != nil {
		if uri := in.Key.GetGcpSecretManagerSecretUri(); uri != "" {
			out.Key = &krm.RegistryClientKey{
				SecretRef: &secretmanagerv1beta1.SecretVersionRef{
					External: uri,
				},
			}
		}
	}
	return out
}

func RegistryClient_ToProto(mapCtx *direct.MapContext, in *krm.RegistryClient) *pb.ContainerdConfig_RegistryHostConfig_CertificateConfigPair {
	if in == nil {
		return nil
	}
	out := &pb.ContainerdConfig_RegistryHostConfig_CertificateConfigPair{}
	if in.Cert != nil && in.Cert.SecretRef != nil && in.Cert.SecretRef.External != "" {
		out.Cert = &pb.ContainerdConfig_RegistryHostConfig_CertificateConfig{
			Certificate: &pb.ContainerdConfig_RegistryHostConfig_CertificateConfig_GcpSecretManagerSecretUri{
				GcpSecretManagerSecretUri: in.Cert.SecretRef.External,
			},
		}
	}
	if in.Key != nil && in.Key.SecretRef != nil && in.Key.SecretRef.External != "" {
		out.Key = &pb.ContainerdConfig_RegistryHostConfig_CertificateConfig{
			Certificate: &pb.ContainerdConfig_RegistryHostConfig_CertificateConfig_GcpSecretManagerSecretUri{
				GcpSecretManagerSecretUri: in.Key.SecretRef.External,
			},
		}
	}
	return out
}
