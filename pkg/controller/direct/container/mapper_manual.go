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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
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

func GpuDriverInstallationConfig_GpuDriverVersion_ToProto(mapCtx *direct.MapContext, in *string) *pb.GPUDriverInstallationConfig_GPUDriverVersion {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.GPUDriverInstallationConfig_GPUDriverVersion](mapCtx, in)
	return &v
}

func GpuSharingConfig_GpuSharingStrategy_ToProto(mapCtx *direct.MapContext, in *string) *pb.GPUSharingConfig_GPUSharingStrategy {
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

func PodAutoscaling_HpaProfile_ToProto(mapCtx *direct.MapContext, in *string) *pb.PodAutoscaling_HPAProfile {
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

func GcpFilestoreCsiDriverConfig_FromProto(mapCtx *direct.MapContext, in *pb.GcpFilestoreCsiDriverConfig) *krm.GcpFilestoreCsiDriverConfig {
	if in == nil {
		return nil
	}
	out := &krm.GcpFilestoreCsiDriverConfig{}
	out.Enabled = &in.Enabled
	return out
}

func GcpFilestoreCsiDriverConfig_ToProto(mapCtx *direct.MapContext, in *krm.GcpFilestoreCsiDriverConfig) *pb.GcpFilestoreCsiDriverConfig {
	if in == nil {
		return nil
	}
	out := &pb.GcpFilestoreCsiDriverConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
