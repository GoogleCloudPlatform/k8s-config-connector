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
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	secretv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CompliancePostureConfig_Mode_ToProto(mapCtx *direct.MapContext, in *string) *pb.CompliancePostureConfig_Mode {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.CompliancePostureConfig_Mode](mapCtx, in)
	return &v
}

func GPUDriverInstallationConfig_GPUDriverVersion_ToProto(mapCtx *direct.MapContext, in *string) *pb.GPUDriverInstallationConfig_GPUDriverVersion {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.GPUDriverInstallationConfig_GPUDriverVersion](mapCtx, in)
	return &v
}

func GPUSharingConfig_GPUSharingStrategy_ToProto(mapCtx *direct.MapContext, in *string) *pb.GPUSharingConfig_GPUSharingStrategy {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.GPUSharingConfig_GPUSharingStrategy](mapCtx, in)
	return &v
}

func MasterAuth_Password_FromProto(mapCtx *direct.MapContext, in string) *secretv1beta1.Legacy {
	if in == "" {
		return nil
	}
	return &secretv1beta1.Legacy{Value: &in}
}

func MasterAuth_Password_ToProto(mapCtx *direct.MapContext, in *secretv1beta1.Legacy) string {
	if in == nil {
		return ""
	}
	return direct.ValueOf(in.Value)
}

func NetworkConfig_InTransitEncryptionConfig_ToProto(mapCtx *direct.MapContext, in *string) *pb.InTransitEncryptionConfig {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.InTransitEncryptionConfig](mapCtx, in)
	return &v
}

func NetworkConfig_ClusterNetworkPerformanceConfig_TotalEgressBandwidthTier_ToProto(mapCtx *direct.MapContext, in *string) *pb.NetworkConfig_ClusterNetworkPerformanceConfig_Tier {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.NetworkConfig_ClusterNetworkPerformanceConfig_Tier](mapCtx, in)
	return &v
}

func NodeNetworkConfig_NetworkPerformanceConfig_TotalEgressBandwidthTier_ToProto(mapCtx *direct.MapContext, in *string) *pb.NodeNetworkConfig_NetworkPerformanceConfig_Tier {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.NodeNetworkConfig_NetworkPerformanceConfig_Tier](mapCtx, in)
	return &v
}

func NodePool_UpgradeSettings_Strategy_ToProto(mapCtx *direct.MapContext, in *string) *pb.NodePoolUpdateStrategy {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.NodePoolUpdateStrategy](mapCtx, in)
	return &v
}

func PodAutoscaling_HPAProfile_ToProto(mapCtx *direct.MapContext, in *string) *pb.PodAutoscaling_HPAProfile {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.PodAutoscaling_HPAProfile](mapCtx, in)
	return &v
}

func SecurityPostureConfig_Mode_ToProto(mapCtx *direct.MapContext, in *string) *pb.SecurityPostureConfig_Mode {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.SecurityPostureConfig_Mode](mapCtx, in)
	return &v
}

func SecurityPostureConfig_VulnerabilityMode_ToProto(mapCtx *direct.MapContext, in *string) *pb.SecurityPostureConfig_VulnerabilityMode {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.SecurityPostureConfig_VulnerabilityMode](mapCtx, in)
	return &v
}

func MaintenanceExclusionOptions_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceExclusionOptions) *pb.MaintenanceExclusionOptions {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceExclusionOptions{}
	out.Scope = direct.Enum_ToProto[pb.MaintenanceExclusionOptions_Scope](mapCtx, in.Scope)
	return out
}

func AdvancedMachineFeatures_PerformanceMonitoringUnit_ToProto(mapCtx *direct.MapContext, in *string) *pb.AdvancedMachineFeatures_PerformanceMonitoringUnit {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.AdvancedMachineFeatures_PerformanceMonitoringUnit](mapCtx, in)
	return &v
}

func GCPFilestoreCSIDriverConfig_FromProto(mapCtx *direct.MapContext, in *pb.GcpFilestoreCsiDriverConfig) *krm.GCPFilestoreCSIDriverConfig {
	if in == nil {
		return nil
	}
	out := &krm.GCPFilestoreCSIDriverConfig{}
	out.Enabled = &in.Enabled
	return out
}

func GCPFilestoreCSIDriverConfig_ToProto(mapCtx *direct.MapContext, in *krm.GCPFilestoreCSIDriverConfig) *pb.GcpFilestoreCsiDriverConfig {
	if in == nil {
		return nil
	}
	out := &pb.GcpFilestoreCsiDriverConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}

func GCEPersistentDiskCSIDriverConfig_FromProto(mapCtx *direct.MapContext, in *pb.GcePersistentDiskCsiDriverConfig) *krm.GCEPersistentDiskCSIDriverConfig {
	if in == nil {
		return nil
	}
	out := &krm.GCEPersistentDiskCSIDriverConfig{}
	out.Enabled = &in.Enabled
	return out
}

func GCEPersistentDiskCSIDriverConfig_ToProto(mapCtx *direct.MapContext, in *krm.GCEPersistentDiskCSIDriverConfig) *pb.GcePersistentDiskCsiDriverConfig {
	if in == nil {
		return nil
	}
	out := &pb.GcePersistentDiskCsiDriverConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}

func GKEBackupAgentConfig_FromProto(mapCtx *direct.MapContext, in *pb.GkeBackupAgentConfig) *krm.GKEBackupAgentConfig {
	if in == nil {
		return nil
	}
	out := &krm.GKEBackupAgentConfig{}
	out.Enabled = &in.Enabled
	return out
}

func GKEBackupAgentConfig_ToProto(mapCtx *direct.MapContext, in *krm.GKEBackupAgentConfig) *pb.GkeBackupAgentConfig {
	if in == nil {
		return nil
	}
	out := &pb.GkeBackupAgentConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}

func GCSFuseCSIDriverConfig_FromProto(mapCtx *direct.MapContext, in *pb.GcsFuseCsiDriverConfig) *krm.GCSFuseCSIDriverConfig {
	if in == nil {
		return nil
	}
	out := &krm.GCSFuseCSIDriverConfig{}
	out.Enabled = &in.Enabled
	return out
}

func GCSFuseCSIDriverConfig_ToProto(mapCtx *direct.MapContext, in *krm.GCSFuseCSIDriverConfig) *pb.GcsFuseCsiDriverConfig {
	if in == nil {
		return nil
	}
	out := &pb.GcsFuseCsiDriverConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}

func HTTPLoadBalancing_FromProto(mapCtx *direct.MapContext, in *pb.HttpLoadBalancing) *krm.HTTPLoadBalancing {
	if in == nil {
		return nil
	}
	out := &krm.HTTPLoadBalancing{}
	out.Disabled = &in.Disabled
	return out
}

func HTTPLoadBalancing_ToProto(mapCtx *direct.MapContext, in *krm.HTTPLoadBalancing) *pb.HttpLoadBalancing {
	if in == nil {
		return nil
	}
	out := &pb.HttpLoadBalancing{}
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}

func DNSCacheConfig_FromProto(mapCtx *direct.MapContext, in *pb.DnsCacheConfig) *krm.DNSCacheConfig {
	if in == nil {
		return nil
	}
	out := &krm.DNSCacheConfig{}
	out.Enabled = &in.Enabled
	return out
}

func DNSCacheConfig_ToProto(mapCtx *direct.MapContext, in *krm.DNSCacheConfig) *pb.DnsCacheConfig {
	if in == nil {
		return nil
	}
	out := &pb.DnsCacheConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}

func MasterAuth_FromProto(mapCtx *direct.MapContext, in *pb.MasterAuth) *krm.MasterAuth {
	if in == nil {
		return nil
	}
	out := &krm.MasterAuth{}
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = MasterAuth_Password_FromProto(mapCtx, in.GetPassword())
	out.ClientCertificate = direct.LazyPtr(in.GetClientCertificate())
	out.ClientKey = direct.LazyPtr(in.GetClientKey())
	out.ClusterCACertificate = direct.LazyPtr(in.GetClusterCaCertificate())
	return out
}

func MasterAuth_ToProto(mapCtx *direct.MapContext, in *krm.MasterAuth) *pb.MasterAuth {
	if in == nil {
		return nil
	}
	out := &pb.MasterAuth{}
	out.Username = direct.ValueOf(in.Username)
	out.Password = MasterAuth_Password_ToProto(mapCtx, in.Password)
	out.ClientCertificate = direct.ValueOf(in.ClientCertificate)
	out.ClientKey = direct.ValueOf(in.ClientKey)
	out.ClusterCaCertificate = direct.ValueOf(in.ClusterCACertificate)
	return out
}

func GPUSharingConfig_FromProto(mapCtx *direct.MapContext, in *pb.GPUSharingConfig) *krm.GPUSharingConfig {
	if in == nil {
		return nil
	}
	out := &krm.GPUSharingConfig{}
	out.GPUSharingStrategy = direct.LazyPtr(in.GetGpuSharingStrategy().String())
	if in.GetMaxSharedClientsPerGpu() != 0 {
		out.MaxSharedClientsPerGPU = direct.LazyPtr(int(in.GetMaxSharedClientsPerGpu()))
	}
	return out
}

func GPUSharingConfig_ToProto(mapCtx *direct.MapContext, in *krm.GPUSharingConfig) *pb.GPUSharingConfig {
	if in == nil {
		return nil
	}
	out := &pb.GPUSharingConfig{}
	out.GpuSharingStrategy = GPUSharingConfig_GPUSharingStrategy_ToProto(mapCtx, in.GPUSharingStrategy)
	out.MaxSharedClientsPerGpu = int64(direct.ValueOf(in.MaxSharedClientsPerGPU))
	return out
}

func ResourceLimits_FromProto(mapCtx *direct.MapContext, in *pb.ResourceLimit) *krm.ResourceLimits {
	if in == nil {
		return nil
	}
	out := &krm.ResourceLimits{}
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	out.Maximum = direct.LazyPtr(int(in.GetMaximum()))
	out.Minimum = direct.LazyPtr(int(in.GetMinimum()))
	return out
}

func ResourceLimits_ToProto(mapCtx *direct.MapContext, in *krm.ResourceLimits) *pb.ResourceLimit {
	if in == nil {
		return nil
	}
	out := &pb.ResourceLimit{}
	out.ResourceType = direct.ValueOf(in.ResourceType)
	out.Maximum = int64(direct.ValueOf(in.Maximum))
	out.Minimum = int64(direct.ValueOf(in.Minimum))
	return out
}

func KubeletConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeKubeletConfig) *krm.KubeletConfig {
	if in == nil {
		return nil
	}
	out := &krm.KubeletConfig{}
	out.CPUManagerPolicy = direct.LazyPtr(in.GetCpuManagerPolicy())
	out.CPUCfsQuota = direct.BoolValue_FromProto(mapCtx, in.GetCpuCfsQuota())
	out.CPUCfsQuotaPeriod = direct.LazyPtr(in.GetCpuCfsQuotaPeriod())
	out.PodPidsLimit = direct.LazyPtr(int(in.GetPodPidsLimit()))
	out.ImageGcLowThresholdPercent = direct.LazyPtr(int(in.GetImageGcLowThresholdPercent()))
	out.ImageGcHighThresholdPercent = direct.LazyPtr(int(in.GetImageGcHighThresholdPercent()))
	out.ImageMinimumGcAge = direct.LazyPtr(in.GetImageMinimumGcAge())
	out.ImageMaximumGcAge = direct.LazyPtr(in.GetImageMaximumGcAge())
	return out
}

func KubeletConfig_ToProto(mapCtx *direct.MapContext, in *krm.KubeletConfig) *pb.NodeKubeletConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeKubeletConfig{}
	out.CpuManagerPolicy = direct.ValueOf(in.CPUManagerPolicy)
	out.CpuCfsQuota = direct.BoolValue_ToProto(mapCtx, in.CPUCfsQuota)
	out.CpuCfsQuotaPeriod = direct.ValueOf(in.CPUCfsQuotaPeriod)
	out.PodPidsLimit = int64(direct.ValueOf(in.PodPidsLimit))
	out.ImageGcLowThresholdPercent = int32(direct.ValueOf(in.ImageGcLowThresholdPercent))
	out.ImageGcHighThresholdPercent = int32(direct.ValueOf(in.ImageGcHighThresholdPercent))
	out.ImageMinimumGcAge = direct.ValueOf(in.ImageMinimumGcAge)
	out.ImageMaximumGcAge = direct.ValueOf(in.ImageMaximumGcAge)
	return out
}

func K8SBetaAPIConfig_FromProto(mapCtx *direct.MapContext, in *pb.K8SBetaAPIConfig) *krm.K8SBetaAPIConfig {
	if in == nil {
		return nil
	}
	out := &krm.K8SBetaAPIConfig{}
	out.EnabledApis = in.EnabledApis
	return out
}

func K8SBetaAPIConfig_ToProto(mapCtx *direct.MapContext, in *krm.K8SBetaAPIConfig) *pb.K8SBetaAPIConfig {
	if in == nil {
		return nil
	}
	out := &pb.K8SBetaAPIConfig{}
	out.EnabledApis = in.EnabledApis
	return out
}

func ContainerClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.ContainerClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContainerClusterSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.InitialNodeCount = direct.LazyPtr(int(in.GetInitialNodeCount()))
	out.NodeConfig = NodeConfig_FromProto(mapCtx, in.GetNodeConfig())
	out.MasterAuth = MasterAuth_FromProto(mapCtx, in.GetMasterAuth())
	out.LoggingService = direct.LazyPtr(in.GetLoggingService())
	out.MonitoringService = direct.LazyPtr(in.GetMonitoringService())
	if in.GetNetwork() != "" {
		out.NetworkRef = &computerefs.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.ClusterIPV4CIDR = direct.LazyPtr(in.GetClusterIpv4Cidr())
	out.AddonsConfig = AddonsConfig_FromProto(mapCtx, in.GetAddonsConfig())
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &computev1beta1.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	out.EnableKubernetesAlpha = direct.LazyPtr(in.GetEnableKubernetesAlpha())
	out.NetworkPolicy = NetworkPolicy_FromProto(mapCtx, in.GetNetworkPolicy())
	out.IPAllocationPolicy = IPAllocationPolicy_FromProto(mapCtx, in.GetIpAllocationPolicy())
	out.MasterAuthorizedNetworksConfig = MasterAuthorizedNetworksConfig_FromProto(mapCtx, in.GetMasterAuthorizedNetworksConfig())
	out.MaintenancePolicy = MaintenancePolicy_FromProto(mapCtx, in.GetMaintenancePolicy())
	out.BinaryAuthorization = BinaryAuthorization_FromProto(mapCtx, in.GetBinaryAuthorization())
	out.ResourceUsageExportConfig = ResourceUsageExportConfig_FromProto(mapCtx, in.GetResourceUsageExportConfig())
	out.AuthenticatorGroupsConfig = AuthenticatorGroupsConfig_FromProto(mapCtx, in.GetAuthenticatorGroupsConfig())
	out.PrivateClusterConfig = PrivateClusterConfig_FromProto(mapCtx, in.GetPrivateClusterConfig())
	out.DatabaseEncryption = DatabaseEncryption_FromProto(mapCtx, in.GetDatabaseEncryption())
	out.VerticalPodAutoscaling = VerticalPodAutoscaling_FromProto(mapCtx, in.GetVerticalPodAutoscaling())
	out.ReleaseChannel = ReleaseChannel_FromProto(mapCtx, in.GetReleaseChannel())
	out.WorkloadIdentityConfig = WorkloadIdentityConfig_FromProto(mapCtx, in.GetWorkloadIdentityConfig())
	out.MeshCertificates = MeshCertificates_FromProto(mapCtx, in.GetMeshCertificates())
	out.CostManagementConfig = CostManagementConfig_FromProto(mapCtx, in.GetCostManagementConfig())
	out.NotificationConfig = NotificationConfig_FromProto(mapCtx, in.GetNotificationConfig())
	out.ConfidentialNodes = ConfidentialNodes_FromProto(mapCtx, in.GetConfidentialNodes())
	out.IdentityServiceConfig = IdentityServiceConfig_FromProto(mapCtx, in.GetIdentityServiceConfig())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.EnableTpu = direct.LazyPtr(in.GetEnableTpu())
	out.NodePoolDefaults = NodePoolDefaults_FromProto(mapCtx, in.GetNodePoolDefaults())
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	out.MonitoringConfig = MonitoringConfig_FromProto(mapCtx, in.GetMonitoringConfig())
	out.NodePoolAutoConfig = NodePoolAutoConfig_FromProto(mapCtx, in.GetNodePoolAutoConfig())
	out.SecurityPostureConfig = SecurityPostureConfig_FromProto(mapCtx, in.GetSecurityPostureConfig())
	out.ControlPlaneEndpointsConfig = ControlPlaneEndpointsConfig_FromProto(mapCtx, in.GetControlPlaneEndpointsConfig())
	out.EnableK8SBetaApis = K8SBetaAPIConfig_FromProto(mapCtx, in.GetEnableK8SBetaApis())

	// Flattened fields
	if in.GetAutopilot() != nil {
		out.EnableAutopilot = direct.LazyPtr(in.GetAutopilot().GetEnabled())
	}
	if in.GetLegacyAbac() != nil {
		out.EnableLegacyAbac = direct.LazyPtr(in.GetLegacyAbac().GetEnabled())
	}
	if in.GetShieldedNodes() != nil {
		out.EnableShieldedNodes = direct.LazyPtr(in.GetShieldedNodes().GetEnabled())
	}
	if in.GetDefaultMaxPodsConstraint() != nil {
		out.DefaultMaxPodsPerNode = direct.LazyPtr(int(in.GetDefaultMaxPodsConstraint().GetMaxPodsPerNode()))
	}
	if in.GetNetworkConfig() != nil {
		nc := in.GetNetworkConfig()
		out.EnableL4ILBSubsetting = direct.LazyPtr(nc.GetEnableL4IlbSubsetting())
		out.EnableIntranodeVisibility = direct.LazyPtr(nc.GetEnableIntraNodeVisibility())
		out.EnableCiliumClusterwideNetworkPolicy = direct.LazyPtr(nc.GetEnableCiliumClusterwideNetworkPolicy())
		out.EnableFQDNNetworkPolicy = direct.LazyPtr(nc.GetEnableFqdnNetworkPolicy())
		out.EnableMultiNetworking = direct.LazyPtr(nc.GetEnableMultiNetworking())
		out.DatapathProvider = direct.Enum_FromProto[pb.DatapathProvider](mapCtx, nc.GetDatapathProvider())
	}

	return out
}

func ContainerClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContainerClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.Description = direct.ValueOf(in.Description)
	out.InitialNodeCount = int32(direct.ValueOf(in.InitialNodeCount))
	out.NodeConfig = NodeConfig_ToProto(mapCtx, in.NodeConfig)
	out.MasterAuth = MasterAuth_ToProto(mapCtx, in.MasterAuth)
	out.LoggingService = direct.ValueOf(in.LoggingService)
	out.MonitoringService = direct.ValueOf(in.MonitoringService)
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	out.ClusterIpv4Cidr = direct.ValueOf(in.ClusterIPV4CIDR)
	out.AddonsConfig = AddonsConfig_ToProto(mapCtx, in.AddonsConfig)
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	out.EnableKubernetesAlpha = direct.ValueOf(in.EnableKubernetesAlpha)
	out.NetworkPolicy = NetworkPolicy_ToProto(mapCtx, in.NetworkPolicy)
	out.IpAllocationPolicy = IPAllocationPolicy_ToProto(mapCtx, in.IPAllocationPolicy)
	out.MasterAuthorizedNetworksConfig = MasterAuthorizedNetworksConfig_ToProto(mapCtx, in.MasterAuthorizedNetworksConfig)
	out.MaintenancePolicy = MaintenancePolicy_ToProto(mapCtx, in.MaintenancePolicy)
	out.BinaryAuthorization = BinaryAuthorization_ToProto(mapCtx, in.BinaryAuthorization)
	out.ResourceUsageExportConfig = ResourceUsageExportConfig_ToProto(mapCtx, in.ResourceUsageExportConfig)
	out.AuthenticatorGroupsConfig = AuthenticatorGroupsConfig_ToProto(mapCtx, in.AuthenticatorGroupsConfig)
	out.PrivateClusterConfig = PrivateClusterConfig_ToProto(mapCtx, in.PrivateClusterConfig)
	out.DatabaseEncryption = DatabaseEncryption_ToProto(mapCtx, in.DatabaseEncryption)
	out.VerticalPodAutoscaling = VerticalPodAutoscaling_ToProto(mapCtx, in.VerticalPodAutoscaling)
	out.ReleaseChannel = ReleaseChannel_ToProto(mapCtx, in.ReleaseChannel)
	out.WorkloadIdentityConfig = WorkloadIdentityConfig_ToProto(mapCtx, in.WorkloadIdentityConfig)
	out.MeshCertificates = MeshCertificates_ToProto(mapCtx, in.MeshCertificates)
	out.CostManagementConfig = CostManagementConfig_ToProto(mapCtx, in.CostManagementConfig)
	out.NotificationConfig = NotificationConfig_ToProto(mapCtx, in.NotificationConfig)
	out.ConfidentialNodes = ConfidentialNodes_ToProto(mapCtx, in.ConfidentialNodes)
	out.IdentityServiceConfig = IdentityServiceConfig_ToProto(mapCtx, in.IdentityServiceConfig)
	out.Location = direct.ValueOf(in.Location)
	out.EnableTpu = direct.ValueOf(in.EnableTpu)
	out.NodePoolDefaults = NodePoolDefaults_ToProto(mapCtx, in.NodePoolDefaults)
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	out.MonitoringConfig = MonitoringConfig_ToProto(mapCtx, in.MonitoringConfig)
	out.NodePoolAutoConfig = NodePoolAutoConfig_ToProto(mapCtx, in.NodePoolAutoConfig)
	out.SecurityPostureConfig = SecurityPostureConfig_ToProto(mapCtx, in.SecurityPostureConfig)
	out.ControlPlaneEndpointsConfig = ControlPlaneEndpointsConfig_ToProto(mapCtx, in.ControlPlaneEndpointsConfig)
	out.EnableK8SBetaApis = K8SBetaAPIConfig_ToProto(mapCtx, in.EnableK8SBetaApis)

	// Flattened fields
	if in.EnableAutopilot != nil {
		out.Autopilot = &pb.Autopilot{Enabled: direct.ValueOf(in.EnableAutopilot)}
	}
	if in.EnableLegacyAbac != nil {
		out.LegacyAbac = &pb.LegacyAbac{Enabled: direct.ValueOf(in.EnableLegacyAbac)}
	}
	if in.EnableShieldedNodes != nil {
		out.ShieldedNodes = &pb.ShieldedNodes{Enabled: direct.ValueOf(in.EnableShieldedNodes)}
	}
	if in.DefaultMaxPodsPerNode != nil {
		out.DefaultMaxPodsConstraint = &pb.MaxPodsConstraint{MaxPodsPerNode: int64(direct.ValueOf(in.DefaultMaxPodsPerNode))}
	}
	if in.EnableL4ILBSubsetting != nil || in.EnableIntranodeVisibility != nil || in.EnableCiliumClusterwideNetworkPolicy != nil || in.EnableFQDNNetworkPolicy != nil || in.EnableMultiNetworking != nil || in.DatapathProvider != nil {
		out.NetworkConfig = &pb.NetworkConfig{}
		out.NetworkConfig.EnableL4IlbSubsetting = direct.ValueOf(in.EnableL4ILBSubsetting)
		out.NetworkConfig.EnableIntraNodeVisibility = direct.ValueOf(in.EnableIntranodeVisibility)
		out.NetworkConfig.EnableCiliumClusterwideNetworkPolicy = in.EnableCiliumClusterwideNetworkPolicy
		out.NetworkConfig.EnableFqdnNetworkPolicy = in.EnableFQDNNetworkPolicy
		out.NetworkConfig.EnableMultiNetworking = direct.ValueOf(in.EnableMultiNetworking)
		out.NetworkConfig.DatapathProvider = direct.Enum_ToProto[pb.DatapathProvider](mapCtx, in.DatapathProvider)
	}

	return out
}

func EphemeralStorageLocalSsdConfig_FromProto(mapCtx *direct.MapContext, in *pb.EphemeralStorageLocalSsdConfig) *krm.EphemeralStorageLocalSsdConfig {
	if in == nil {
		return nil
	}
	out := &krm.EphemeralStorageLocalSsdConfig{}
	out.LocalSsdCount = direct.LazyPtr(int(in.GetLocalSsdCount()))
	out.DataCacheCount = direct.LazyPtr(int(in.GetDataCacheCount()))
	return out
}

func EphemeralStorageLocalSsdConfig_ToProto(mapCtx *direct.MapContext, in *krm.EphemeralStorageLocalSsdConfig) *pb.EphemeralStorageLocalSsdConfig {
	if in == nil {
		return nil
	}
	out := &pb.EphemeralStorageLocalSsdConfig{}
	out.LocalSsdCount = int32(direct.ValueOf(in.LocalSsdCount))
	out.DataCacheCount = int32(direct.ValueOf(in.DataCacheCount))
	return out
}

func NodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeConfig) *krm.NodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeConfig{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.OauthScopes = in.OauthScopes
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.Metadata = in.Metadata
	out.ImageType = direct.LazyPtr(in.GetImageType())
	out.Labels = in.Labels
	out.LocalSsdCount = direct.LazyPtr(int(in.GetLocalSsdCount()))
	out.Tags = in.Tags
	out.Preemptible = direct.LazyPtr(in.GetPreemptible())
	out.DiskType = direct.LazyPtr(in.GetDiskType())
	out.MinCPUPlatform = direct.LazyPtr(in.GetMinCpuPlatform())
	out.SandboxConfig = SandboxConfig_FromProto(mapCtx, in.GetSandboxConfig())
	if in.GetNodeGroup() != "" {
		out.NodeGroupRef = &computev1beta1.ComputeNodeGroupRef{External: in.GetNodeGroup()}
	}
	out.ReservationAffinity = ReservationAffinity_FromProto(mapCtx, in.GetReservationAffinity())
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.LinuxNodeConfig = LinuxNodeConfig_FromProto(mapCtx, in.GetLinuxNodeConfig())
	out.KubeletConfig = KubeletConfig_FromProto(mapCtx, in.GetKubeletConfig())
	out.GcfsConfig = GcfsConfig_FromProto(mapCtx, in.GetGcfsConfig())
	out.Spot = direct.LazyPtr(in.GetSpot())
	out.ConfidentialNodes = ConfidentialNodes_FromProto(mapCtx, in.GetConfidentialNodes())
	out.FastSocket = FastSocket_FromProto(mapCtx, in.GetFastSocket())
	out.ResourceLabels = in.ResourceLabels
	out.EphemeralStorageLocalSsdConfig = EphemeralStorageLocalSsdConfig_FromProto(mapCtx, in.GetEphemeralStorageLocalSsdConfig())
	return out
}

func NodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeConfig) *pb.NodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeConfig{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.OauthScopes = in.OauthScopes
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.Metadata = in.Metadata
	out.ImageType = direct.ValueOf(in.ImageType)
	out.Labels = in.Labels
	out.LocalSsdCount = int32(direct.ValueOf(in.LocalSsdCount))
	out.Tags = in.Tags
	out.Preemptible = direct.ValueOf(in.Preemptible)
	out.DiskType = direct.ValueOf(in.DiskType)
	out.MinCpuPlatform = direct.ValueOf(in.MinCPUPlatform)
	out.SandboxConfig = SandboxConfig_ToProto(mapCtx, in.SandboxConfig)
	if in.NodeGroupRef != nil {
		out.NodeGroup = in.NodeGroupRef.External
	}
	out.ReservationAffinity = ReservationAffinity_ToProto(mapCtx, in.ReservationAffinity)
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.LinuxNodeConfig = LinuxNodeConfig_ToProto(mapCtx, in.LinuxNodeConfig)
	out.KubeletConfig = KubeletConfig_ToProto(mapCtx, in.KubeletConfig)
	out.GcfsConfig = GcfsConfig_ToProto(mapCtx, in.GcfsConfig)
	out.Spot = direct.ValueOf(in.Spot)
	out.ConfidentialNodes = ConfidentialNodes_ToProto(mapCtx, in.ConfidentialNodes)
	out.FastSocket = FastSocket_ToProto(mapCtx, in.FastSocket)
	out.ResourceLabels = in.ResourceLabels
	out.EphemeralStorageLocalSsdConfig = EphemeralStorageLocalSsdConfig_ToProto(mapCtx, in.EphemeralStorageLocalSsdConfig)
	return out
}
