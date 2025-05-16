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

package composer

import (
	pb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/composer/v1alpha1"
	computev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func MasterAuthorizedNetworksConfig_FromProto(mapCtx *direct.MapContext, in *pb.MasterAuthorizedNetworksConfig) *krm.MasterAuthorizedNetworksConfig {
	if in == nil {
		return nil
	}
	out := &krm.MasterAuthorizedNetworksConfig{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.CIDRBlocks = direct.Slice_FromProto(mapCtx, in.CidrBlocks, MasterAuthorizedNetworksConfig_CIDRBlock_FromProto)
	return out
}
func MasterAuthorizedNetworksConfig_ToProto(mapCtx *direct.MapContext, in *krm.MasterAuthorizedNetworksConfig) *pb.MasterAuthorizedNetworksConfig {
	if in == nil {
		return nil
	}
	out := &pb.MasterAuthorizedNetworksConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.CidrBlocks = direct.Slice_ToProto(mapCtx, in.CIDRBlocks, MasterAuthorizedNetworksConfig_CIDRBlock_ToProto)
	return out
}

func WebServerNetworkAccessControl_FromProto(mapCtx *direct.MapContext, in *pb.WebServerNetworkAccessControl) *krm.WebServerNetworkAccessControl {
	if in == nil {
		return nil
	}
	out := &krm.WebServerNetworkAccessControl{}
	out.AllowedIPRanges = direct.Slice_FromProto(mapCtx, in.AllowedIpRanges, WebServerNetworkAccessControl_AllowedIPRange_FromProto)
	return out
}
func WebServerNetworkAccessControl_ToProto(mapCtx *direct.MapContext, in *krm.WebServerNetworkAccessControl) *pb.WebServerNetworkAccessControl {
	if in == nil {
		return nil
	}
	out := &pb.WebServerNetworkAccessControl{}
	out.AllowedIpRanges = direct.Slice_ToProto(mapCtx, in.AllowedIPRanges, WebServerNetworkAccessControl_AllowedIPRange_ToProto)
	return out
}

func IPAllocationPolicy_FromProto(mapCtx *direct.MapContext, in *pb.IPAllocationPolicy) *krm.IPAllocationPolicy {
	if in == nil {
		return nil
	}
	out := &krm.IPAllocationPolicy{}
	out.UseIPAliases = direct.LazyPtr(in.GetUseIpAliases())
	out.ClusterSecondaryRangeName = direct.LazyPtr(in.GetClusterSecondaryRangeName())
	out.ClusterIPV4CIDRBlock = direct.LazyPtr(in.GetClusterIpv4CidrBlock())
	out.ServicesSecondaryRangeName = direct.LazyPtr(in.GetServicesSecondaryRangeName())
	out.ServicesIPV4CIDRBlock = direct.LazyPtr(in.GetServicesIpv4CidrBlock())
	return out
}
func IPAllocationPolicy_ToProto(mapCtx *direct.MapContext, in *krm.IPAllocationPolicy) *pb.IPAllocationPolicy {
	if in == nil {
		return nil
	}
	out := &pb.IPAllocationPolicy{}
	out.UseIpAliases = direct.ValueOf(in.UseIPAliases)
	if in.ClusterSecondaryRangeName != nil {
		out.ClusterIpAllocation = &pb.IPAllocationPolicy_ClusterSecondaryRangeName{
			ClusterSecondaryRangeName: direct.ValueOf(in.ClusterSecondaryRangeName),
		}
	}
	if in.ClusterIPV4CIDRBlock != nil {
		out.ClusterIpAllocation = &pb.IPAllocationPolicy_ClusterIpv4CidrBlock{
			ClusterIpv4CidrBlock: direct.ValueOf(in.ClusterIPV4CIDRBlock),
		}
	}
	if in.ServicesSecondaryRangeName != nil {
		out.ServicesIpAllocation = &pb.IPAllocationPolicy_ServicesSecondaryRangeName{
			ServicesSecondaryRangeName: direct.ValueOf(in.ServicesSecondaryRangeName),
		}
	}
	if in.ServicesIPV4CIDRBlock != nil {
		out.ServicesIpAllocation = &pb.IPAllocationPolicy_ServicesIpv4CidrBlock{
			ServicesIpv4CidrBlock: direct.ValueOf(in.ServicesIPV4CIDRBlock),
		}
	}
	return out
}

func PrivateEnvironmentConfig_FromProto(mapCtx *direct.MapContext, in *pb.PrivateEnvironmentConfig) *krm.PrivateEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &krm.PrivateEnvironmentConfig{}
	out.EnablePrivateEnvironment = direct.LazyPtr(in.GetEnablePrivateEnvironment())
	out.EnablePrivateBuildsOnly = direct.LazyPtr(in.GetEnablePrivateBuildsOnly())
	out.PrivateClusterConfig = PrivateClusterConfig_FromProto(mapCtx, in.GetPrivateClusterConfig())
	out.WebServerIPv4CIDRBlock = direct.LazyPtr(in.GetWebServerIpv4CidrBlock())
	out.CloudSQLIPv4CIDRBlock = direct.LazyPtr(in.GetCloudSqlIpv4CidrBlock())
	out.CloudComposerNetworkIPv4CIDRBlock = direct.LazyPtr(in.GetCloudComposerNetworkIpv4CidrBlock())
	out.EnablePrivatelyUsedPublicIPs = direct.LazyPtr(in.GetEnablePrivatelyUsedPublicIps())
	if in.GetCloudComposerConnectionSubnetwork() != "" {
		out.CloudComposerConnectionSubnetworkRef = &computev1beta1.ComputeSubnetworkRef{External: in.GetCloudComposerConnectionSubnetwork()}
	}
	// MISSING: WebServerIPV4ReservedRange
	// MISSING: CloudComposerNetworkIPV4ReservedRange
	if in.GetCloudComposerConnectionSubnetwork() != "" {
		out.CloudComposerConnectionSubnetworkRef = &computev1beta1.ComputeSubnetworkRef{External: in.GetCloudComposerConnectionSubnetwork()}
	}
	out.NetworkingConfig = NetworkingConfig_FromProto(mapCtx, in.GetNetworkingConfig())
	return out
}
func PrivateEnvironmentConfig_ToProto(mapCtx *direct.MapContext, in *krm.PrivateEnvironmentConfig) *pb.PrivateEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivateEnvironmentConfig{}
	out.EnablePrivateEnvironment = direct.ValueOf(in.EnablePrivateEnvironment)
	out.EnablePrivateBuildsOnly = direct.ValueOf(in.EnablePrivateBuildsOnly)
	out.PrivateClusterConfig = PrivateClusterConfig_ToProto(mapCtx, in.PrivateClusterConfig)
	out.WebServerIpv4CidrBlock = direct.ValueOf(in.WebServerIPv4CIDRBlock)
	out.CloudSqlIpv4CidrBlock = direct.ValueOf(in.CloudSQLIPv4CIDRBlock)
	out.CloudComposerNetworkIpv4CidrBlock = direct.ValueOf(in.CloudComposerNetworkIPv4CIDRBlock)
	out.EnablePrivatelyUsedPublicIps = direct.ValueOf(in.EnablePrivatelyUsedPublicIPs)
	// MISSING: WebServerIPV4ReservedRange
	// MISSING: CloudComposerNetworkIPV4ReservedRange
	if in.CloudComposerConnectionSubnetworkRef != nil {
		out.CloudComposerConnectionSubnetwork = in.CloudComposerConnectionSubnetworkRef.External
	}
	out.NetworkingConfig = NetworkingConfig_ToProto(mapCtx, in.NetworkingConfig)
	return out
}

func PrivateEnvironmentConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateEnvironmentConfig) *krm.PrivateEnvironmentConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivateEnvironmentConfigObservedState{}
	// MISSING: EnablePrivateEnvironment
	// MISSING: EnablePrivateBuildsOnly
	out.PrivateClusterConfig = PrivateClusterConfigObservedState_FromProto(mapCtx, in.GetPrivateClusterConfig())
	// MISSING: WebServerIPV4CIDRBlock
	// MISSING: CloudSQLIPV4CIDRBlock
	out.WebServerIPv4ReservedRange = direct.LazyPtr(in.GetWebServerIpv4ReservedRange())
	out.CloudComposerNetworkIPv4ReservedRange = direct.LazyPtr(in.GetCloudComposerNetworkIpv4ReservedRange())
	// MISSING: CloudComposerNetworkIPV4CIDRBlock
	// MISSING: EnablePrivatelyUsedPublicIps
	// MISSING: CloudComposerConnectionSubnetwork
	// MISSING: NetworkingConfig
	return out
}
func PrivateEnvironmentConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivateEnvironmentConfigObservedState) *pb.PrivateEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivateEnvironmentConfig{}
	// MISSING: EnablePrivateEnvironment
	// MISSING: EnablePrivateBuildsOnly
	out.PrivateClusterConfig = PrivateClusterConfigObservedState_ToProto(mapCtx, in.PrivateClusterConfig)
	// MISSING: WebServerIPV4CIDRBlock
	// MISSING: CloudSQLIPV4CIDRBlock
	out.WebServerIpv4ReservedRange = direct.ValueOf(in.WebServerIPv4ReservedRange)
	out.CloudComposerNetworkIpv4ReservedRange = direct.ValueOf(in.CloudComposerNetworkIPv4ReservedRange)
	// MISSING: CloudComposerNetworkIPV4CIDRBlock
	// MISSING: EnablePrivatelyUsedPublicIps
	// MISSING: CloudComposerConnectionSubnetwork
	// MISSING: NetworkingConfig
	return out
}
func NodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeConfig) *krm.NodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeConfig{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	if in.GetNetwork() != "" {
		out.NetworkRef = &computev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &computev1beta1.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	// MISSING: OauthScopes
	// (near miss): "OauthScopes" vs "OAuthScopes"
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refs.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.Tags = in.Tags
	out.IPAllocationPolicy = IPAllocationPolicy_FromProto(mapCtx, in.GetIpAllocationPolicy())
	out.EnableIPMasqAgent = direct.LazyPtr(in.GetEnableIpMasqAgent())
	if in.GetComposerNetworkAttachment() != "" {
		out.ComposerNetworkAttachmentRef = &computev1alpha1.ComputeNetworkAttachmentRef{External: in.GetComposerNetworkAttachment()}
	}
	out.ComposerInternalIPv4CIDRBlock = direct.LazyPtr(in.ComposerInternalIpv4CidrBlock)
	return out
}
func NodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeConfig) *pb.NodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeConfig{}
	out.Location = direct.ValueOf(in.Location)
	out.MachineType = direct.ValueOf(in.MachineType)
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	out.OauthScopes = in.OAuthScopes
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.Tags = in.Tags
	out.IpAllocationPolicy = IPAllocationPolicy_ToProto(mapCtx, in.IPAllocationPolicy)
	out.EnableIpMasqAgent = direct.ValueOf(in.EnableIPMasqAgent)
	if in.ComposerNetworkAttachmentRef != nil {
		out.ComposerNetworkAttachment = in.ComposerNetworkAttachmentRef.External
	}
	out.ComposerInternalIpv4CidrBlock = direct.ValueOf(in.ComposerInternalIPv4CIDRBlock)
	return out
}

func WorkloadsConfig_DagProcessorResource_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadsConfig_DagProcessorResource) *krm.WorkloadsConfig_DagProcessorResource {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadsConfig_DagProcessorResource{}
	out.CPU = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetCpu()))
	out.MemoryGB = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetMemoryGb()))
	out.StorageGB = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetStorageGb()))
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func WorkloadsConfig_DagProcessorResource_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadsConfig_DagProcessorResource) *pb.WorkloadsConfig_DagProcessorResource {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadsConfig_DagProcessorResource{}
	out.Cpu = direct.StringToFloat32(mapCtx, direct.ValueOf(in.CPU))
	out.MemoryGb = direct.StringToFloat32(mapCtx, direct.ValueOf(in.MemoryGB))
	out.StorageGb = direct.StringToFloat32(mapCtx, direct.ValueOf(in.StorageGB))
	out.Count = direct.ValueOf(in.Count)
	return out
}
func WorkloadsConfig_SchedulerResource_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadsConfig_SchedulerResource) *krm.WorkloadsConfig_SchedulerResource {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadsConfig_SchedulerResource{}
	out.CPU = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetCpu()))
	out.MemoryGB = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetMemoryGb()))
	out.StorageGB = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetStorageGb()))
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func WorkloadsConfig_SchedulerResource_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadsConfig_SchedulerResource) *pb.WorkloadsConfig_SchedulerResource {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadsConfig_SchedulerResource{}
	out.Cpu = direct.StringToFloat32(mapCtx, direct.ValueOf(in.CPU))
	out.MemoryGb = direct.StringToFloat32(mapCtx, direct.ValueOf(in.MemoryGB))
	out.StorageGb = direct.StringToFloat32(mapCtx, direct.ValueOf(in.StorageGB))
	out.Count = direct.ValueOf(in.Count)
	return out
}
func WorkloadsConfig_TriggererResource_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadsConfig_TriggererResource) *krm.WorkloadsConfig_TriggererResource {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadsConfig_TriggererResource{}
	out.Count = direct.LazyPtr(in.GetCount())
	out.CPU = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetCpu()))
	out.MemoryGB = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetMemoryGb()))
	return out
}
func WorkloadsConfig_TriggererResource_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadsConfig_TriggererResource) *pb.WorkloadsConfig_TriggererResource {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadsConfig_TriggererResource{}
	out.Count = direct.ValueOf(in.Count)
	out.Cpu = direct.StringToFloat32(mapCtx, direct.ValueOf(in.CPU))
	out.MemoryGb = direct.StringToFloat32(mapCtx, direct.ValueOf(in.MemoryGB))
	return out
}
func WorkloadsConfig_WebServerResource_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadsConfig_WebServerResource) *krm.WorkloadsConfig_WebServerResource {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadsConfig_WebServerResource{}
	out.CPU = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetCpu()))
	out.MemoryGB = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetMemoryGb()))
	out.StorageGB = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetStorageGb()))
	return out
}
func WorkloadsConfig_WebServerResource_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadsConfig_WebServerResource) *pb.WorkloadsConfig_WebServerResource {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadsConfig_WebServerResource{}
	out.Cpu = direct.StringToFloat32(mapCtx, direct.ValueOf(in.CPU))
	out.MemoryGb = direct.StringToFloat32(mapCtx, direct.ValueOf(in.MemoryGB))
	out.StorageGb = direct.StringToFloat32(mapCtx, direct.ValueOf(in.StorageGB))
	return out
}
func WorkloadsConfig_WorkerResource_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadsConfig_WorkerResource) *krm.WorkloadsConfig_WorkerResource {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadsConfig_WorkerResource{}
	out.CPU = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetCpu()))
	out.MemoryGB = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetMemoryGb()))
	out.StorageGB = direct.PtrTo(direct.Float32ToString(mapCtx, in.GetStorageGb()))
	out.MinCount = direct.LazyPtr(in.GetMinCount())
	out.MaxCount = direct.LazyPtr(in.GetMaxCount())
	return out
}
func WorkloadsConfig_WorkerResource_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadsConfig_WorkerResource) *pb.WorkloadsConfig_WorkerResource {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadsConfig_WorkerResource{}
	out.Cpu = direct.StringToFloat32(mapCtx, direct.ValueOf(in.CPU))
	out.MemoryGb = direct.StringToFloat32(mapCtx, direct.ValueOf(in.MemoryGB))
	out.StorageGb = direct.StringToFloat32(mapCtx, direct.ValueOf(in.StorageGB))
	out.MinCount = direct.ValueOf(in.MinCount)
	out.MaxCount = direct.ValueOf(in.MaxCount)
	return out
}
