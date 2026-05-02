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

package container

import (
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/container/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AdvancedMachineFeatures_PerformanceMonitoringUnit_ToProto(mapCtx *direct.MapContext, in *string) *pb.AdvancedMachineFeatures_PerformanceMonitoringUnit {
	if in == nil {
		return nil
	}
	val := pb.AdvancedMachineFeatures_PerformanceMonitoringUnit(pb.AdvancedMachineFeatures_PerformanceMonitoringUnit_value[*in])
	return &val
}

func CompliancePostureConfig_Mode_ToProto(mapCtx *direct.MapContext, in *string) *pb.CompliancePostureConfig_Mode {
	if in == nil {
		return nil
	}
	val := pb.CompliancePostureConfig_Mode(pb.CompliancePostureConfig_Mode_value[*in])
	return &val
}

func DatabaseEncryptionObservedState_CurrentState_ToProto(mapCtx *direct.MapContext, in *string) *pb.DatabaseEncryption_CurrentState {
	if in == nil {
		return nil
	}
	val := pb.DatabaseEncryption_CurrentState(pb.DatabaseEncryption_CurrentState_value[*in])
	return &val
}

func GpuDriverInstallationConfig_GpuDriverVersion_ToProto(mapCtx *direct.MapContext, in *string) *pb.GPUDriverInstallationConfig_GPUDriverVersion {
	if in == nil {
		return nil
	}
	val := pb.GPUDriverInstallationConfig_GPUDriverVersion(pb.GPUDriverInstallationConfig_GPUDriverVersion_value[*in])
	return &val
}

func GpuSharingConfig_GpuSharingStrategy_ToProto(mapCtx *direct.MapContext, in *string) *pb.GPUSharingConfig_GPUSharingStrategy {
	if in == nil {
		return nil
	}
	val := pb.GPUSharingConfig_GPUSharingStrategy(pb.GPUSharingConfig_GPUSharingStrategy_value[*in])
	return &val
}

func HostMaintenancePolicy_MaintenanceInterval_ToProto(mapCtx *direct.MapContext, in *string) *pb.HostMaintenancePolicy_MaintenanceInterval {
	if in == nil {
		return nil
	}
	val := pb.HostMaintenancePolicy_MaintenanceInterval(pb.HostMaintenancePolicy_MaintenanceInterval_value[*in])
	return &val
}

func NetworkConfig_InTransitEncryptionConfig_ToProto(mapCtx *direct.MapContext, in *string) *pb.InTransitEncryptionConfig {
	if in == nil {
		return nil
	}
	val := pb.InTransitEncryptionConfig(pb.InTransitEncryptionConfig_value[*in])
	return &val
}

func NetworkConfig_ClusterNetworkPerformanceConfig_TotalEgressBandwidthTier_ToProto(mapCtx *direct.MapContext, in *string) *pb.NetworkConfig_ClusterNetworkPerformanceConfig_Tier {
	if in == nil {
		return nil
	}
	val := pb.NetworkConfig_ClusterNetworkPerformanceConfig_Tier(pb.NetworkConfig_ClusterNetworkPerformanceConfig_Tier_value[*in])
	return &val
}

func NodeConfig_LocalSsdEncryptionMode_ToProto(mapCtx *direct.MapContext, in *string) *pb.NodeConfig_LocalSsdEncryptionMode {
	if in == nil {
		return nil
	}
	val := pb.NodeConfig_LocalSsdEncryptionMode(pb.NodeConfig_LocalSsdEncryptionMode_value[*in])
	return &val
}

func NodeNetworkConfig_NetworkPerformanceConfig_TotalEgressBandwidthTier_ToProto(mapCtx *direct.MapContext, in *string) *pb.NodeNetworkConfig_NetworkPerformanceConfig_Tier {
	if in == nil {
		return nil
	}
	val := pb.NodeNetworkConfig_NetworkPerformanceConfig_Tier(pb.NodeNetworkConfig_NetworkPerformanceConfig_Tier_value[*in])
	return &val
}

func NodeNetworkConfig_NetworkPerformanceConfig_ExternalIpEgressBandwidthTier_ToProto(mapCtx *direct.MapContext, in *string) *pb.NodeNetworkConfig_NetworkPerformanceConfig_Tier {
	if in == nil {
		return nil
	}
	val := pb.NodeNetworkConfig_NetworkPerformanceConfig_Tier(pb.NodeNetworkConfig_NetworkPerformanceConfig_Tier_value[*in])
	return &val
}

func NodePool_UpgradeSettings_Strategy_ToProto(mapCtx *direct.MapContext, in *string) *pb.NodePoolUpdateStrategy {
	if in == nil {
		return nil
	}
	val := pb.NodePoolUpdateStrategy(pb.NodePoolUpdateStrategy_value[*in])
	return &val
}

func PodAutoscaling_HpaProfile_ToProto(mapCtx *direct.MapContext, in *string) *pb.PodAutoscaling_HPAProfile {
	if in == nil {
		return nil
	}
	val := pb.PodAutoscaling_HPAProfile(pb.PodAutoscaling_HPAProfile_value[*in])
	return &val
}

func ProtectConfig_WorkloadVulnerabilityMode_ToProto(mapCtx *direct.MapContext, in *string) *pb.ProtectConfig_WorkloadVulnerabilityMode {
	if in == nil {
		return nil
	}
	val := pb.ProtectConfig_WorkloadVulnerabilityMode(pb.ProtectConfig_WorkloadVulnerabilityMode_value[*in])
	return &val
}

func SecurityPostureConfig_Mode_ToProto(mapCtx *direct.MapContext, in *string) *pb.SecurityPostureConfig_Mode {
	if in == nil {
		return nil
	}
	val := pb.SecurityPostureConfig_Mode(pb.SecurityPostureConfig_Mode_value[*in])
	return &val
}

func SecurityPostureConfig_VulnerabilityMode_ToProto(mapCtx *direct.MapContext, in *string) *pb.SecurityPostureConfig_VulnerabilityMode {
	if in == nil {
		return nil
	}
	val := pb.SecurityPostureConfig_VulnerabilityMode(pb.SecurityPostureConfig_VulnerabilityMode_value[*in])
	return &val
}

func WorkloadConfig_AuditMode_ToProto(mapCtx *direct.MapContext, in *string) *pb.WorkloadConfig_Mode {
	if in == nil {
		return nil
	}
	val := pb.WorkloadConfig_Mode(pb.WorkloadConfig_Mode_value[*in])
	return &val
}
