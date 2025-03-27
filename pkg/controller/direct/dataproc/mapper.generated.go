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

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AcceleratorConfig_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorConfig) *krm.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krm.AcceleratorConfig{}
	out.AcceleratorTypeURI = direct.LazyPtr(in.GetAcceleratorTypeUri())
	out.AcceleratorCount = direct.LazyPtr(in.GetAcceleratorCount())
	return out
}
func AcceleratorConfig_ToProto(mapCtx *direct.MapContext, in *krm.AcceleratorConfig) *pb.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorConfig{}
	out.AcceleratorTypeUri = direct.ValueOf(in.AcceleratorTypeURI)
	out.AcceleratorCount = direct.ValueOf(in.AcceleratorCount)
	return out
}
func AutoscalingConfig_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingConfig) *krm.AutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingConfig{}
	out.PolicyURI = direct.LazyPtr(in.GetPolicyUri())
	return out
}
func AutoscalingConfig_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingConfig) *pb.AutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingConfig{}
	out.PolicyUri = direct.ValueOf(in.PolicyURI)
	return out
}
func AuxiliaryNodeGroup_FromProto(mapCtx *direct.MapContext, in *pb.AuxiliaryNodeGroup) *krm.AuxiliaryNodeGroup {
	if in == nil {
		return nil
	}
	out := &krm.AuxiliaryNodeGroup{}
	out.NodeGroup = NodeGroup_FromProto(mapCtx, in.GetNodeGroup())
	out.NodeGroupID = direct.LazyPtr(in.GetNodeGroupId())
	return out
}
func AuxiliaryNodeGroup_ToProto(mapCtx *direct.MapContext, in *krm.AuxiliaryNodeGroup) *pb.AuxiliaryNodeGroup {
	if in == nil {
		return nil
	}
	out := &pb.AuxiliaryNodeGroup{}
	out.NodeGroup = NodeGroup_ToProto(mapCtx, in.NodeGroup)
	out.NodeGroupId = direct.ValueOf(in.NodeGroupID)
	return out
}
func ClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.ClusterConfig) *krm.ClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterConfig{}
	out.ConfigBucket = direct.LazyPtr(in.GetConfigBucket())
	out.TempBucket = direct.LazyPtr(in.GetTempBucket())
	out.GCEClusterConfig = GCEClusterConfig_FromProto(mapCtx, in.GetGceClusterConfig())
	out.MasterConfig = InstanceGroupConfig_FromProto(mapCtx, in.GetMasterConfig())
	out.WorkerConfig = InstanceGroupConfig_FromProto(mapCtx, in.GetWorkerConfig())
	out.SecondaryWorkerConfig = InstanceGroupConfig_FromProto(mapCtx, in.GetSecondaryWorkerConfig())
	out.SoftwareConfig = SoftwareConfig_FromProto(mapCtx, in.GetSoftwareConfig())
	out.InitializationActions = direct.Slice_FromProto(mapCtx, in.InitializationActions, NodeInitializationAction_FromProto)
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	out.AutoscalingConfig = AutoscalingConfig_FromProto(mapCtx, in.GetAutoscalingConfig())
	out.SecurityConfig = SecurityConfig_FromProto(mapCtx, in.GetSecurityConfig())
	out.LifecycleConfig = LifecycleConfig_FromProto(mapCtx, in.GetLifecycleConfig())
	out.EndpointConfig = EndpointConfig_FromProto(mapCtx, in.GetEndpointConfig())
	out.MetastoreConfig = MetastoreConfig_FromProto(mapCtx, in.GetMetastoreConfig())
	out.DataprocMetricConfig = DataprocMetricConfig_FromProto(mapCtx, in.GetDataprocMetricConfig())
	out.AuxiliaryNodeGroups = direct.Slice_FromProto(mapCtx, in.AuxiliaryNodeGroups, AuxiliaryNodeGroup_FromProto)
	return out
}
func ClusterConfig_ToProto(mapCtx *direct.MapContext, in *krm.ClusterConfig) *pb.ClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.ClusterConfig{}
	out.ConfigBucket = direct.ValueOf(in.ConfigBucket)
	out.TempBucket = direct.ValueOf(in.TempBucket)
	out.GceClusterConfig = GCEClusterConfig_ToProto(mapCtx, in.GCEClusterConfig)
	out.MasterConfig = InstanceGroupConfig_ToProto(mapCtx, in.MasterConfig)
	out.WorkerConfig = InstanceGroupConfig_ToProto(mapCtx, in.WorkerConfig)
	out.SecondaryWorkerConfig = InstanceGroupConfig_ToProto(mapCtx, in.SecondaryWorkerConfig)
	out.SoftwareConfig = SoftwareConfig_ToProto(mapCtx, in.SoftwareConfig)
	out.InitializationActions = direct.Slice_ToProto(mapCtx, in.InitializationActions, NodeInitializationAction_ToProto)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	out.AutoscalingConfig = AutoscalingConfig_ToProto(mapCtx, in.AutoscalingConfig)
	out.SecurityConfig = SecurityConfig_ToProto(mapCtx, in.SecurityConfig)
	out.LifecycleConfig = LifecycleConfig_ToProto(mapCtx, in.LifecycleConfig)
	out.EndpointConfig = EndpointConfig_ToProto(mapCtx, in.EndpointConfig)
	out.MetastoreConfig = MetastoreConfig_ToProto(mapCtx, in.MetastoreConfig)
	out.DataprocMetricConfig = DataprocMetricConfig_ToProto(mapCtx, in.DataprocMetricConfig)
	out.AuxiliaryNodeGroups = direct.Slice_ToProto(mapCtx, in.AuxiliaryNodeGroups, AuxiliaryNodeGroup_ToProto)
	return out
}
func ClusterConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ClusterConfig) *krm.ClusterConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClusterConfigObservedState{}
	// MISSING: ConfigBucket
	// MISSING: TempBucket
	// MISSING: GCEClusterConfig
	out.MasterConfig = InstanceGroupConfigObservedState_FromProto(mapCtx, in.GetMasterConfig())
	// MISSING: WorkerConfig
	// MISSING: SecondaryWorkerConfig
	// MISSING: SoftwareConfig
	// MISSING: InitializationActions
	// MISSING: EncryptionConfig
	// MISSING: AutoscalingConfig
	// MISSING: SecurityConfig
	out.LifecycleConfig = LifecycleConfigObservedState_FromProto(mapCtx, in.GetLifecycleConfig())
	out.EndpointConfig = EndpointConfigObservedState_FromProto(mapCtx, in.GetEndpointConfig())
	// MISSING: MetastoreConfig
	// MISSING: DataprocMetricConfig
	// MISSING: AuxiliaryNodeGroups
	return out
}
func ClusterConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClusterConfigObservedState) *pb.ClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.ClusterConfig{}
	// MISSING: ConfigBucket
	// MISSING: TempBucket
	// MISSING: GCEClusterConfig
	out.MasterConfig = InstanceGroupConfigObservedState_ToProto(mapCtx, in.MasterConfig)
	// MISSING: WorkerConfig
	// MISSING: SecondaryWorkerConfig
	// MISSING: SoftwareConfig
	// MISSING: InitializationActions
	// MISSING: EncryptionConfig
	// MISSING: AutoscalingConfig
	// MISSING: SecurityConfig
	out.LifecycleConfig = LifecycleConfigObservedState_ToProto(mapCtx, in.LifecycleConfig)
	out.EndpointConfig = EndpointConfigObservedState_ToProto(mapCtx, in.EndpointConfig)
	// MISSING: MetastoreConfig
	// MISSING: DataprocMetricConfig
	// MISSING: AuxiliaryNodeGroups
	return out
}
func ClusterSelector_FromProto(mapCtx *direct.MapContext, in *pb.ClusterSelector) *krm.ClusterSelector {
	if in == nil {
		return nil
	}
	out := &krm.ClusterSelector{}
	out.Zone = direct.LazyPtr(in.GetZone())
	out.ClusterLabels = in.ClusterLabels
	return out
}
func ClusterSelector_ToProto(mapCtx *direct.MapContext, in *krm.ClusterSelector) *pb.ClusterSelector {
	if in == nil {
		return nil
	}
	out := &pb.ClusterSelector{}
	out.Zone = direct.ValueOf(in.Zone)
	out.ClusterLabels = in.ClusterLabels
	return out
}
func ConfidentialInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.ConfidentialInstanceConfig) *krm.ConfidentialInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.ConfidentialInstanceConfig{}
	out.EnableConfidentialCompute = direct.LazyPtr(in.GetEnableConfidentialCompute())
	return out
}
func ConfidentialInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.ConfidentialInstanceConfig) *pb.ConfidentialInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ConfidentialInstanceConfig{}
	out.EnableConfidentialCompute = direct.ValueOf(in.EnableConfidentialCompute)
	return out
}
func DataprocMetricConfig_FromProto(mapCtx *direct.MapContext, in *pb.DataprocMetricConfig) *krm.DataprocMetricConfig {
	if in == nil {
		return nil
	}
	out := &krm.DataprocMetricConfig{}
	out.Metrics = direct.Slice_FromProto(mapCtx, in.Metrics, DataprocMetricConfig_Metric_FromProto)
	return out
}
func DataprocMetricConfig_ToProto(mapCtx *direct.MapContext, in *krm.DataprocMetricConfig) *pb.DataprocMetricConfig {
	if in == nil {
		return nil
	}
	out := &pb.DataprocMetricConfig{}
	out.Metrics = direct.Slice_ToProto(mapCtx, in.Metrics, DataprocMetricConfig_Metric_ToProto)
	return out
}
func DataprocMetricConfig_Metric_FromProto(mapCtx *direct.MapContext, in *pb.DataprocMetricConfig_Metric) *krm.DataprocMetricConfig_Metric {
	if in == nil {
		return nil
	}
	out := &krm.DataprocMetricConfig_Metric{}
	out.MetricSource = direct.Enum_FromProto(mapCtx, in.GetMetricSource())
	out.MetricOverrides = in.MetricOverrides
	return out
}
func DataprocMetricConfig_Metric_ToProto(mapCtx *direct.MapContext, in *krm.DataprocMetricConfig_Metric) *pb.DataprocMetricConfig_Metric {
	if in == nil {
		return nil
	}
	out := &pb.DataprocMetricConfig_Metric{}
	out.MetricSource = direct.Enum_ToProto[pb.DataprocMetricConfig_MetricSource](mapCtx, in.MetricSource)
	out.MetricOverrides = in.MetricOverrides
	return out
}
func DiskConfig_FromProto(mapCtx *direct.MapContext, in *pb.DiskConfig) *krm.DiskConfig {
	if in == nil {
		return nil
	}
	out := &krm.DiskConfig{}
	out.BootDiskType = direct.LazyPtr(in.GetBootDiskType())
	out.BootDiskSizeGB = direct.LazyPtr(in.GetBootDiskSizeGb())
	out.NumLocalSsds = direct.LazyPtr(in.GetNumLocalSsds())
	out.LocalSsdInterface = direct.LazyPtr(in.GetLocalSsdInterface())
	out.BootDiskProvisionedIops = in.BootDiskProvisionedIops
	out.BootDiskProvisionedThroughput = in.BootDiskProvisionedThroughput
	return out
}
func DiskConfig_ToProto(mapCtx *direct.MapContext, in *krm.DiskConfig) *pb.DiskConfig {
	if in == nil {
		return nil
	}
	out := &pb.DiskConfig{}
	out.BootDiskType = direct.ValueOf(in.BootDiskType)
	out.BootDiskSizeGb = direct.ValueOf(in.BootDiskSizeGB)
	out.NumLocalSsds = direct.ValueOf(in.NumLocalSsds)
	out.LocalSsdInterface = direct.ValueOf(in.LocalSsdInterface)
	out.BootDiskProvisionedIops = in.BootDiskProvisionedIops
	out.BootDiskProvisionedThroughput = in.BootDiskProvisionedThroughput
	return out
}
func EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfig{}
	out.GCEPDKMSKeyName = direct.LazyPtr(in.GetGcePdKmsKeyName())
	out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	return out
}
func EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfig) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	out.GcePdKmsKeyName = direct.ValueOf(in.GCEPDKMSKeyName)
	out.KmsKey = direct.ValueOf(in.KMSKey)
	return out
}
func EndpointConfig_FromProto(mapCtx *direct.MapContext, in *pb.EndpointConfig) *krm.EndpointConfig {
	if in == nil {
		return nil
	}
	out := &krm.EndpointConfig{}
	// MISSING: HTTPPorts
	out.EnableHTTPPortAccess = direct.LazyPtr(in.GetEnableHttpPortAccess())
	return out
}
func EndpointConfig_ToProto(mapCtx *direct.MapContext, in *krm.EndpointConfig) *pb.EndpointConfig {
	if in == nil {
		return nil
	}
	out := &pb.EndpointConfig{}
	// MISSING: HTTPPorts
	out.EnableHttpPortAccess = direct.ValueOf(in.EnableHTTPPortAccess)
	return out
}
func EndpointConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EndpointConfig) *krm.EndpointConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EndpointConfigObservedState{}
	out.HTTPPorts = in.HttpPorts
	// MISSING: EnableHTTPPortAccess
	return out
}
func EndpointConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EndpointConfigObservedState) *pb.EndpointConfig {
	if in == nil {
		return nil
	}
	out := &pb.EndpointConfig{}
	out.HttpPorts = in.HTTPPorts
	// MISSING: EnableHTTPPortAccess
	return out
}
func GCEClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.GceClusterConfig) *krm.GCEClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.GCEClusterConfig{}
	out.ZoneURI = direct.LazyPtr(in.GetZoneUri())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.SubnetworkURI = direct.LazyPtr(in.GetSubnetworkUri())
	out.InternalIPOnly = in.InternalIpOnly
	out.PrivateIPV6GoogleAccess = direct.Enum_FromProto(mapCtx, in.GetPrivateIpv6GoogleAccess())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ServiceAccountScopes = in.ServiceAccountScopes
	out.Tags = in.Tags
	out.Metadata = in.Metadata
	out.ReservationAffinity = ReservationAffinity_FromProto(mapCtx, in.GetReservationAffinity())
	out.NodeGroupAffinity = NodeGroupAffinity_FromProto(mapCtx, in.GetNodeGroupAffinity())
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.ConfidentialInstanceConfig = ConfidentialInstanceConfig_FromProto(mapCtx, in.GetConfidentialInstanceConfig())
	return out
}
func GCEClusterConfig_ToProto(mapCtx *direct.MapContext, in *krm.GCEClusterConfig) *pb.GceClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.GceClusterConfig{}
	out.ZoneUri = direct.ValueOf(in.ZoneURI)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	out.SubnetworkUri = direct.ValueOf(in.SubnetworkURI)
	out.InternalIpOnly = in.InternalIPOnly
	out.PrivateIpv6GoogleAccess = direct.Enum_ToProto[pb.GceClusterConfig_PrivateIpv6GoogleAccess](mapCtx, in.PrivateIPV6GoogleAccess)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ServiceAccountScopes = in.ServiceAccountScopes
	out.Tags = in.Tags
	out.Metadata = in.Metadata
	out.ReservationAffinity = ReservationAffinity_ToProto(mapCtx, in.ReservationAffinity)
	out.NodeGroupAffinity = NodeGroupAffinity_ToProto(mapCtx, in.NodeGroupAffinity)
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.ConfidentialInstanceConfig = ConfidentialInstanceConfig_ToProto(mapCtx, in.ConfidentialInstanceConfig)
	return out
}
func IdentityConfig_FromProto(mapCtx *direct.MapContext, in *pb.IdentityConfig) *krm.IdentityConfig {
	if in == nil {
		return nil
	}
	out := &krm.IdentityConfig{}
	out.UserServiceAccountMapping = in.UserServiceAccountMapping
	return out
}
func IdentityConfig_ToProto(mapCtx *direct.MapContext, in *krm.IdentityConfig) *pb.IdentityConfig {
	if in == nil {
		return nil
	}
	out := &pb.IdentityConfig{}
	out.UserServiceAccountMapping = in.UserServiceAccountMapping
	return out
}
func InstanceFlexibilityPolicy_FromProto(mapCtx *direct.MapContext, in *pb.InstanceFlexibilityPolicy) *krm.InstanceFlexibilityPolicy {
	if in == nil {
		return nil
	}
	out := &krm.InstanceFlexibilityPolicy{}
	out.ProvisioningModelMix = InstanceFlexibilityPolicy_ProvisioningModelMix_FromProto(mapCtx, in.GetProvisioningModelMix())
	out.InstanceSelectionList = direct.Slice_FromProto(mapCtx, in.InstanceSelectionList, InstanceFlexibilityPolicy_InstanceSelection_FromProto)
	// MISSING: InstanceSelectionResults
	return out
}
func InstanceFlexibilityPolicy_ToProto(mapCtx *direct.MapContext, in *krm.InstanceFlexibilityPolicy) *pb.InstanceFlexibilityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.InstanceFlexibilityPolicy{}
	out.ProvisioningModelMix = InstanceFlexibilityPolicy_ProvisioningModelMix_ToProto(mapCtx, in.ProvisioningModelMix)
	out.InstanceSelectionList = direct.Slice_ToProto(mapCtx, in.InstanceSelectionList, InstanceFlexibilityPolicy_InstanceSelection_ToProto)
	// MISSING: InstanceSelectionResults
	return out
}
func InstanceFlexibilityPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceFlexibilityPolicy) *krm.InstanceFlexibilityPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceFlexibilityPolicyObservedState{}
	// MISSING: ProvisioningModelMix
	// MISSING: InstanceSelectionList
	out.InstanceSelectionResults = direct.Slice_FromProto(mapCtx, in.InstanceSelectionResults, InstanceFlexibilityPolicy_InstanceSelectionResult_FromProto)
	return out
}
func InstanceFlexibilityPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceFlexibilityPolicyObservedState) *pb.InstanceFlexibilityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.InstanceFlexibilityPolicy{}
	// MISSING: ProvisioningModelMix
	// MISSING: InstanceSelectionList
	out.InstanceSelectionResults = direct.Slice_ToProto(mapCtx, in.InstanceSelectionResults, InstanceFlexibilityPolicy_InstanceSelectionResult_ToProto)
	return out
}
func InstanceFlexibilityPolicy_InstanceSelection_FromProto(mapCtx *direct.MapContext, in *pb.InstanceFlexibilityPolicy_InstanceSelection) *krm.InstanceFlexibilityPolicy_InstanceSelection {
	if in == nil {
		return nil
	}
	out := &krm.InstanceFlexibilityPolicy_InstanceSelection{}
	out.MachineTypes = in.MachineTypes
	out.Rank = direct.LazyPtr(in.GetRank())
	return out
}
func InstanceFlexibilityPolicy_InstanceSelection_ToProto(mapCtx *direct.MapContext, in *krm.InstanceFlexibilityPolicy_InstanceSelection) *pb.InstanceFlexibilityPolicy_InstanceSelection {
	if in == nil {
		return nil
	}
	out := &pb.InstanceFlexibilityPolicy_InstanceSelection{}
	out.MachineTypes = in.MachineTypes
	out.Rank = direct.ValueOf(in.Rank)
	return out
}
func InstanceFlexibilityPolicy_InstanceSelectionResult_FromProto(mapCtx *direct.MapContext, in *pb.InstanceFlexibilityPolicy_InstanceSelectionResult) *krm.InstanceFlexibilityPolicy_InstanceSelectionResult {
	if in == nil {
		return nil
	}
	out := &krm.InstanceFlexibilityPolicy_InstanceSelectionResult{}
	// MISSING: MachineType
	// MISSING: VmCount
	return out
}
func InstanceFlexibilityPolicy_InstanceSelectionResult_ToProto(mapCtx *direct.MapContext, in *krm.InstanceFlexibilityPolicy_InstanceSelectionResult) *pb.InstanceFlexibilityPolicy_InstanceSelectionResult {
	if in == nil {
		return nil
	}
	out := &pb.InstanceFlexibilityPolicy_InstanceSelectionResult{}
	// MISSING: MachineType
	// MISSING: VmCount
	return out
}
func InstanceFlexibilityPolicy_InstanceSelectionResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceFlexibilityPolicy_InstanceSelectionResult) *krm.InstanceFlexibilityPolicy_InstanceSelectionResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceFlexibilityPolicy_InstanceSelectionResultObservedState{}
	out.MachineType = in.MachineType
	out.VmCount = in.VmCount
	return out
}
func InstanceFlexibilityPolicy_InstanceSelectionResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceFlexibilityPolicy_InstanceSelectionResultObservedState) *pb.InstanceFlexibilityPolicy_InstanceSelectionResult {
	if in == nil {
		return nil
	}
	out := &pb.InstanceFlexibilityPolicy_InstanceSelectionResult{}
	out.MachineType = in.MachineType
	out.VmCount = in.VmCount
	return out
}
func InstanceFlexibilityPolicy_ProvisioningModelMix_FromProto(mapCtx *direct.MapContext, in *pb.InstanceFlexibilityPolicy_ProvisioningModelMix) *krm.InstanceFlexibilityPolicy_ProvisioningModelMix {
	if in == nil {
		return nil
	}
	out := &krm.InstanceFlexibilityPolicy_ProvisioningModelMix{}
	out.StandardCapacityBase = in.StandardCapacityBase
	out.StandardCapacityPercentAboveBase = in.StandardCapacityPercentAboveBase
	return out
}
func InstanceFlexibilityPolicy_ProvisioningModelMix_ToProto(mapCtx *direct.MapContext, in *krm.InstanceFlexibilityPolicy_ProvisioningModelMix) *pb.InstanceFlexibilityPolicy_ProvisioningModelMix {
	if in == nil {
		return nil
	}
	out := &pb.InstanceFlexibilityPolicy_ProvisioningModelMix{}
	out.StandardCapacityBase = in.StandardCapacityBase
	out.StandardCapacityPercentAboveBase = in.StandardCapacityPercentAboveBase
	return out
}
func InstanceGroupConfig_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupConfig) *krm.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupConfig{}
	out.NumInstances = direct.LazyPtr(in.GetNumInstances())
	// MISSING: InstanceNames
	// MISSING: InstanceReferences
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.MachineTypeURI = direct.LazyPtr(in.GetMachineTypeUri())
	out.DiskConfig = DiskConfig_FromProto(mapCtx, in.GetDiskConfig())
	// MISSING: IsPreemptible
	out.Preemptibility = direct.Enum_FromProto(mapCtx, in.GetPreemptibility())
	// MISSING: ManagedGroupConfig
	out.Accelerators = direct.Slice_FromProto(mapCtx, in.Accelerators, AcceleratorConfig_FromProto)
	out.MinCPUPlatform = direct.LazyPtr(in.GetMinCpuPlatform())
	out.MinNumInstances = direct.LazyPtr(in.GetMinNumInstances())
	out.InstanceFlexibilityPolicy = InstanceFlexibilityPolicy_FromProto(mapCtx, in.GetInstanceFlexibilityPolicy())
	out.StartupConfig = StartupConfig_FromProto(mapCtx, in.GetStartupConfig())
	return out
}
func InstanceGroupConfig_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupConfig) *pb.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupConfig{}
	out.NumInstances = direct.ValueOf(in.NumInstances)
	// MISSING: InstanceNames
	// MISSING: InstanceReferences
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.MachineTypeUri = direct.ValueOf(in.MachineTypeURI)
	out.DiskConfig = DiskConfig_ToProto(mapCtx, in.DiskConfig)
	// MISSING: IsPreemptible
	out.Preemptibility = direct.Enum_ToProto[pb.InstanceGroupConfig_Preemptibility](mapCtx, in.Preemptibility)
	// MISSING: ManagedGroupConfig
	out.Accelerators = direct.Slice_ToProto(mapCtx, in.Accelerators, AcceleratorConfig_ToProto)
	out.MinCpuPlatform = direct.ValueOf(in.MinCPUPlatform)
	out.MinNumInstances = direct.ValueOf(in.MinNumInstances)
	out.InstanceFlexibilityPolicy = InstanceFlexibilityPolicy_ToProto(mapCtx, in.InstanceFlexibilityPolicy)
	out.StartupConfig = StartupConfig_ToProto(mapCtx, in.StartupConfig)
	return out
}
func InstanceGroupConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupConfig) *krm.InstanceGroupConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupConfigObservedState{}
	// MISSING: NumInstances
	out.InstanceNames = in.InstanceNames
	out.InstanceReferences = direct.Slice_FromProto(mapCtx, in.InstanceReferences, InstanceReference_FromProto)
	// MISSING: ImageURI
	// MISSING: MachineTypeURI
	// MISSING: DiskConfig
	out.IsPreemptible = direct.LazyPtr(in.GetIsPreemptible())
	// MISSING: Preemptibility
	out.ManagedGroupConfig = ManagedGroupConfig_FromProto(mapCtx, in.GetManagedGroupConfig())
	// MISSING: Accelerators
	// MISSING: MinCPUPlatform
	// MISSING: MinNumInstances
	out.InstanceFlexibilityPolicy = InstanceFlexibilityPolicyObservedState_FromProto(mapCtx, in.GetInstanceFlexibilityPolicy())
	// MISSING: StartupConfig
	return out
}
func InstanceGroupConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupConfigObservedState) *pb.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupConfig{}
	// MISSING: NumInstances
	out.InstanceNames = in.InstanceNames
	out.InstanceReferences = direct.Slice_ToProto(mapCtx, in.InstanceReferences, InstanceReference_ToProto)
	// MISSING: ImageURI
	// MISSING: MachineTypeURI
	// MISSING: DiskConfig
	out.IsPreemptible = direct.ValueOf(in.IsPreemptible)
	// MISSING: Preemptibility
	out.ManagedGroupConfig = ManagedGroupConfig_ToProto(mapCtx, in.ManagedGroupConfig)
	// MISSING: Accelerators
	// MISSING: MinCPUPlatform
	// MISSING: MinNumInstances
	out.InstanceFlexibilityPolicy = InstanceFlexibilityPolicyObservedState_ToProto(mapCtx, in.InstanceFlexibilityPolicy)
	// MISSING: StartupConfig
	return out
}
func InstanceReference_FromProto(mapCtx *direct.MapContext, in *pb.InstanceReference) *krm.InstanceReference {
	if in == nil {
		return nil
	}
	out := &krm.InstanceReference{}
	out.InstanceName = direct.LazyPtr(in.GetInstanceName())
	out.InstanceID = direct.LazyPtr(in.GetInstanceId())
	out.PublicKey = direct.LazyPtr(in.GetPublicKey())
	out.PublicEciesKey = direct.LazyPtr(in.GetPublicEciesKey())
	return out
}
func InstanceReference_ToProto(mapCtx *direct.MapContext, in *krm.InstanceReference) *pb.InstanceReference {
	if in == nil {
		return nil
	}
	out := &pb.InstanceReference{}
	out.InstanceName = direct.ValueOf(in.InstanceName)
	out.InstanceId = direct.ValueOf(in.InstanceID)
	out.PublicKey = direct.ValueOf(in.PublicKey)
	out.PublicEciesKey = direct.ValueOf(in.PublicEciesKey)
	return out
}
func KerberosConfig_FromProto(mapCtx *direct.MapContext, in *pb.KerberosConfig) *krm.KerberosConfig {
	if in == nil {
		return nil
	}
	out := &krm.KerberosConfig{}
	out.EnableKerberos = direct.LazyPtr(in.GetEnableKerberos())
	out.RootPrincipalPasswordURI = direct.LazyPtr(in.GetRootPrincipalPasswordUri())
	out.KMSKeyURI = direct.LazyPtr(in.GetKmsKeyUri())
	out.KeystoreURI = direct.LazyPtr(in.GetKeystoreUri())
	out.TruststoreURI = direct.LazyPtr(in.GetTruststoreUri())
	out.KeystorePasswordURI = direct.LazyPtr(in.GetKeystorePasswordUri())
	out.KeyPasswordURI = direct.LazyPtr(in.GetKeyPasswordUri())
	out.TruststorePasswordURI = direct.LazyPtr(in.GetTruststorePasswordUri())
	out.CrossRealmTrustRealm = direct.LazyPtr(in.GetCrossRealmTrustRealm())
	out.CrossRealmTrustKdc = direct.LazyPtr(in.GetCrossRealmTrustKdc())
	out.CrossRealmTrustAdminServer = direct.LazyPtr(in.GetCrossRealmTrustAdminServer())
	out.CrossRealmTrustSharedPasswordURI = direct.LazyPtr(in.GetCrossRealmTrustSharedPasswordUri())
	out.KdcDbKeyURI = direct.LazyPtr(in.GetKdcDbKeyUri())
	out.TgtLifetimeHours = direct.LazyPtr(in.GetTgtLifetimeHours())
	out.Realm = direct.LazyPtr(in.GetRealm())
	return out
}
func KerberosConfig_ToProto(mapCtx *direct.MapContext, in *krm.KerberosConfig) *pb.KerberosConfig {
	if in == nil {
		return nil
	}
	out := &pb.KerberosConfig{}
	out.EnableKerberos = direct.ValueOf(in.EnableKerberos)
	out.RootPrincipalPasswordUri = direct.ValueOf(in.RootPrincipalPasswordURI)
	out.KmsKeyUri = direct.ValueOf(in.KMSKeyURI)
	out.KeystoreUri = direct.ValueOf(in.KeystoreURI)
	out.TruststoreUri = direct.ValueOf(in.TruststoreURI)
	out.KeystorePasswordUri = direct.ValueOf(in.KeystorePasswordURI)
	out.KeyPasswordUri = direct.ValueOf(in.KeyPasswordURI)
	out.TruststorePasswordUri = direct.ValueOf(in.TruststorePasswordURI)
	out.CrossRealmTrustRealm = direct.ValueOf(in.CrossRealmTrustRealm)
	out.CrossRealmTrustKdc = direct.ValueOf(in.CrossRealmTrustKdc)
	out.CrossRealmTrustAdminServer = direct.ValueOf(in.CrossRealmTrustAdminServer)
	out.CrossRealmTrustSharedPasswordUri = direct.ValueOf(in.CrossRealmTrustSharedPasswordURI)
	out.KdcDbKeyUri = direct.ValueOf(in.KdcDbKeyURI)
	out.TgtLifetimeHours = direct.ValueOf(in.TgtLifetimeHours)
	out.Realm = direct.ValueOf(in.Realm)
	return out
}
func LifecycleConfig_FromProto(mapCtx *direct.MapContext, in *pb.LifecycleConfig) *krm.LifecycleConfig {
	if in == nil {
		return nil
	}
	out := &krm.LifecycleConfig{}
	out.IdleDeleteTtl = direct.StringDuration_FromProto(mapCtx, in.GetIdleDeleteTtl())
	out.AutoDeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetAutoDeleteTime())
	out.AutoDeleteTtl = direct.StringDuration_FromProto(mapCtx, in.GetAutoDeleteTtl())
	// MISSING: IdleStartTime
	return out
}
func LifecycleConfig_ToProto(mapCtx *direct.MapContext, in *krm.LifecycleConfig) *pb.LifecycleConfig {
	if in == nil {
		return nil
	}
	out := &pb.LifecycleConfig{}
	out.IdleDeleteTtl = direct.StringDuration_ToProto(mapCtx, in.IdleDeleteTtl)
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.AutoDeleteTime); oneof != nil {
		out.Ttl = &pb.LifecycleConfig_AutoDeleteTime{AutoDeleteTime: oneof}
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.AutoDeleteTtl); oneof != nil {
		out.Ttl = &pb.LifecycleConfig_AutoDeleteTtl{AutoDeleteTtl: oneof}
	}
	// MISSING: IdleStartTime
	return out
}
func LifecycleConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LifecycleConfig) *krm.LifecycleConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LifecycleConfigObservedState{}
	// MISSING: IdleDeleteTtl
	// MISSING: AutoDeleteTime
	// MISSING: AutoDeleteTtl
	out.IdleStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetIdleStartTime())
	return out
}
func LifecycleConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LifecycleConfigObservedState) *pb.LifecycleConfig {
	if in == nil {
		return nil
	}
	out := &pb.LifecycleConfig{}
	// MISSING: IdleDeleteTtl
	// MISSING: AutoDeleteTime
	// MISSING: AutoDeleteTtl
	out.IdleStartTime = direct.StringTimestamp_ToProto(mapCtx, in.IdleStartTime)
	return out
}
func ManagedCluster_FromProto(mapCtx *direct.MapContext, in *pb.ManagedCluster) *krm.ManagedCluster {
	if in == nil {
		return nil
	}
	out := &krm.ManagedCluster{}
	out.ClusterName = direct.LazyPtr(in.GetClusterName())
	out.Config = ClusterConfig_FromProto(mapCtx, in.GetConfig())
	out.Labels = in.Labels
	return out
}
func ManagedCluster_ToProto(mapCtx *direct.MapContext, in *krm.ManagedCluster) *pb.ManagedCluster {
	if in == nil {
		return nil
	}
	out := &pb.ManagedCluster{}
	out.ClusterName = direct.ValueOf(in.ClusterName)
	out.Config = ClusterConfig_ToProto(mapCtx, in.Config)
	out.Labels = in.Labels
	return out
}
func ManagedClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagedCluster) *krm.ManagedClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagedClusterObservedState{}
	// MISSING: ClusterName
	out.Config = ClusterConfigObservedState_FromProto(mapCtx, in.GetConfig())
	// MISSING: Labels
	return out
}
func ManagedClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagedClusterObservedState) *pb.ManagedCluster {
	if in == nil {
		return nil
	}
	out := &pb.ManagedCluster{}
	// MISSING: ClusterName
	out.Config = ClusterConfigObservedState_ToProto(mapCtx, in.Config)
	// MISSING: Labels
	return out
}
func ManagedGroupConfig_FromProto(mapCtx *direct.MapContext, in *pb.ManagedGroupConfig) *krm.ManagedGroupConfig {
	if in == nil {
		return nil
	}
	out := &krm.ManagedGroupConfig{}
	// MISSING: InstanceTemplateName
	// MISSING: InstanceGroupManagerName
	// MISSING: InstanceGroupManagerURI
	return out
}
func ManagedGroupConfig_ToProto(mapCtx *direct.MapContext, in *krm.ManagedGroupConfig) *pb.ManagedGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.ManagedGroupConfig{}
	// MISSING: InstanceTemplateName
	// MISSING: InstanceGroupManagerName
	// MISSING: InstanceGroupManagerURI
	return out
}
func ManagedGroupConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagedGroupConfig) *krm.ManagedGroupConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagedGroupConfigObservedState{}
	out.InstanceTemplateName = direct.LazyPtr(in.GetInstanceTemplateName())
	out.InstanceGroupManagerName = direct.LazyPtr(in.GetInstanceGroupManagerName())
	out.InstanceGroupManagerURI = direct.LazyPtr(in.GetInstanceGroupManagerUri())
	return out
}
func ManagedGroupConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagedGroupConfigObservedState) *pb.ManagedGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.ManagedGroupConfig{}
	out.InstanceTemplateName = direct.ValueOf(in.InstanceTemplateName)
	out.InstanceGroupManagerName = direct.ValueOf(in.InstanceGroupManagerName)
	out.InstanceGroupManagerUri = direct.ValueOf(in.InstanceGroupManagerURI)
	return out
}
func MetastoreConfig_FromProto(mapCtx *direct.MapContext, in *pb.MetastoreConfig) *krm.MetastoreConfig {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreConfig{}
	out.DataprocMetastoreService = direct.LazyPtr(in.GetDataprocMetastoreService())
	return out
}
func MetastoreConfig_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreConfig) *pb.MetastoreConfig {
	if in == nil {
		return nil
	}
	out := &pb.MetastoreConfig{}
	out.DataprocMetastoreService = direct.ValueOf(in.DataprocMetastoreService)
	return out
}
func NodeGroup_FromProto(mapCtx *direct.MapContext, in *pb.NodeGroup) *krm.NodeGroup {
	if in == nil {
		return nil
	}
	out := &krm.NodeGroup{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Roles = direct.EnumSlice_FromProto(mapCtx, in.Roles)
	out.NodeGroupConfig = InstanceGroupConfig_FromProto(mapCtx, in.GetNodeGroupConfig())
	out.Labels = in.Labels
	return out
}
func NodeGroup_ToProto(mapCtx *direct.MapContext, in *krm.NodeGroup) *pb.NodeGroup {
	if in == nil {
		return nil
	}
	out := &pb.NodeGroup{}
	out.Name = direct.ValueOf(in.Name)
	out.Roles = direct.EnumSlice_ToProto[pb.NodeGroup_Role](mapCtx, in.Roles)
	out.NodeGroupConfig = InstanceGroupConfig_ToProto(mapCtx, in.NodeGroupConfig)
	out.Labels = in.Labels
	return out
}
func NodeGroupAffinity_FromProto(mapCtx *direct.MapContext, in *pb.NodeGroupAffinity) *krm.NodeGroupAffinity {
	if in == nil {
		return nil
	}
	out := &krm.NodeGroupAffinity{}
	out.NodeGroupURI = direct.LazyPtr(in.GetNodeGroupUri())
	return out
}
func NodeGroupAffinity_ToProto(mapCtx *direct.MapContext, in *krm.NodeGroupAffinity) *pb.NodeGroupAffinity {
	if in == nil {
		return nil
	}
	out := &pb.NodeGroupAffinity{}
	out.NodeGroupUri = direct.ValueOf(in.NodeGroupURI)
	return out
}
func NodeInitializationAction_FromProto(mapCtx *direct.MapContext, in *pb.NodeInitializationAction) *krm.NodeInitializationAction {
	if in == nil {
		return nil
	}
	out := &krm.NodeInitializationAction{}
	out.ExecutableFile = direct.LazyPtr(in.GetExecutableFile())
	out.ExecutionTimeout = direct.StringDuration_FromProto(mapCtx, in.GetExecutionTimeout())
	return out
}
func NodeInitializationAction_ToProto(mapCtx *direct.MapContext, in *krm.NodeInitializationAction) *pb.NodeInitializationAction {
	if in == nil {
		return nil
	}
	out := &pb.NodeInitializationAction{}
	out.ExecutableFile = direct.ValueOf(in.ExecutableFile)
	out.ExecutionTimeout = direct.StringDuration_ToProto(mapCtx, in.ExecutionTimeout)
	return out
}
func OrderedJob_FromProto(mapCtx *direct.MapContext, in *pb.OrderedJob) *krm.OrderedJob {
	if in == nil {
		return nil
	}
	out := &krm.OrderedJob{}
	out.StepID = direct.LazyPtr(in.GetStepId())
	out.HadoopJob = HadoopJob_FromProto(mapCtx, in.GetHadoopJob())
	out.SparkJob = SparkJob_FromProto(mapCtx, in.GetSparkJob())
	out.PysparkJob = PySparkJob_FromProto(mapCtx, in.GetPysparkJob())
	out.HiveJob = HiveJob_FromProto(mapCtx, in.GetHiveJob())
	out.PigJob = PigJob_FromProto(mapCtx, in.GetPigJob())
	out.SparkRJob = SparkRJob_FromProto(mapCtx, in.GetSparkRJob())
	out.SparkSQLJob = SparkSQLJob_FromProto(mapCtx, in.GetSparkSqlJob())
	out.PrestoJob = PrestoJob_FromProto(mapCtx, in.GetPrestoJob())
	out.TrinoJob = TrinoJob_FromProto(mapCtx, in.GetTrinoJob())
	out.FlinkJob = FlinkJob_FromProto(mapCtx, in.GetFlinkJob())
	out.Labels = in.Labels
	out.Scheduling = JobScheduling_FromProto(mapCtx, in.GetScheduling())
	out.PrerequisiteStepIds = in.PrerequisiteStepIds
	return out
}
func OrderedJob_ToProto(mapCtx *direct.MapContext, in *krm.OrderedJob) *pb.OrderedJob {
	if in == nil {
		return nil
	}
	out := &pb.OrderedJob{}
	out.StepId = direct.ValueOf(in.StepID)
	if oneof := HadoopJob_ToProto(mapCtx, in.HadoopJob); oneof != nil {
		out.JobType = &pb.OrderedJob_HadoopJob{HadoopJob: oneof}
	}
	if oneof := SparkJob_ToProto(mapCtx, in.SparkJob); oneof != nil {
		out.JobType = &pb.OrderedJob_SparkJob{SparkJob: oneof}
	}
	if oneof := PySparkJob_ToProto(mapCtx, in.PysparkJob); oneof != nil {
		out.JobType = &pb.OrderedJob_PysparkJob{PysparkJob: oneof}
	}
	if oneof := HiveJob_ToProto(mapCtx, in.HiveJob); oneof != nil {
		out.JobType = &pb.OrderedJob_HiveJob{HiveJob: oneof}
	}
	if oneof := PigJob_ToProto(mapCtx, in.PigJob); oneof != nil {
		out.JobType = &pb.OrderedJob_PigJob{PigJob: oneof}
	}
	if oneof := SparkRJob_ToProto(mapCtx, in.SparkRJob); oneof != nil {
		out.JobType = &pb.OrderedJob_SparkRJob{SparkRJob: oneof}
	}
	if oneof := SparkSQLJob_ToProto(mapCtx, in.SparkSQLJob); oneof != nil {
		out.JobType = &pb.OrderedJob_SparkSqlJob{SparkSqlJob: oneof}
	}
	if oneof := PrestoJob_ToProto(mapCtx, in.PrestoJob); oneof != nil {
		out.JobType = &pb.OrderedJob_PrestoJob{PrestoJob: oneof}
	}
	if oneof := TrinoJob_ToProto(mapCtx, in.TrinoJob); oneof != nil {
		out.JobType = &pb.OrderedJob_TrinoJob{TrinoJob: oneof}
	}
	if oneof := FlinkJob_ToProto(mapCtx, in.FlinkJob); oneof != nil {
		out.JobType = &pb.OrderedJob_FlinkJob{FlinkJob: oneof}
	}
	out.Labels = in.Labels
	out.Scheduling = JobScheduling_ToProto(mapCtx, in.Scheduling)
	out.PrerequisiteStepIds = in.PrerequisiteStepIds
	return out
}
func ParameterValidation_FromProto(mapCtx *direct.MapContext, in *pb.ParameterValidation) *krm.ParameterValidation {
	if in == nil {
		return nil
	}
	out := &krm.ParameterValidation{}
	out.Regex = RegexValidation_FromProto(mapCtx, in.GetRegex())
	out.Values = ValueValidation_FromProto(mapCtx, in.GetValues())
	return out
}
func ParameterValidation_ToProto(mapCtx *direct.MapContext, in *krm.ParameterValidation) *pb.ParameterValidation {
	if in == nil {
		return nil
	}
	out := &pb.ParameterValidation{}
	if oneof := RegexValidation_ToProto(mapCtx, in.Regex); oneof != nil {
		out.ValidationType = &pb.ParameterValidation_Regex{Regex: oneof}
	}
	if oneof := ValueValidation_ToProto(mapCtx, in.Values); oneof != nil {
		out.ValidationType = &pb.ParameterValidation_Values{Values: oneof}
	}
	return out
}
func RegexValidation_FromProto(mapCtx *direct.MapContext, in *pb.RegexValidation) *krm.RegexValidation {
	if in == nil {
		return nil
	}
	out := &krm.RegexValidation{}
	out.Regexes = in.Regexes
	return out
}
func RegexValidation_ToProto(mapCtx *direct.MapContext, in *krm.RegexValidation) *pb.RegexValidation {
	if in == nil {
		return nil
	}
	out := &pb.RegexValidation{}
	out.Regexes = in.Regexes
	return out
}
func ReservationAffinity_FromProto(mapCtx *direct.MapContext, in *pb.ReservationAffinity) *krm.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &krm.ReservationAffinity{}
	out.ConsumeReservationType = direct.Enum_FromProto(mapCtx, in.GetConsumeReservationType())
	out.Key = direct.LazyPtr(in.GetKey())
	out.Values = in.Values
	return out
}
func ReservationAffinity_ToProto(mapCtx *direct.MapContext, in *krm.ReservationAffinity) *pb.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &pb.ReservationAffinity{}
	out.ConsumeReservationType = direct.Enum_ToProto[pb.ReservationAffinity_Type](mapCtx, in.ConsumeReservationType)
	out.Key = direct.ValueOf(in.Key)
	out.Values = in.Values
	return out
}
func SecurityConfig_FromProto(mapCtx *direct.MapContext, in *pb.SecurityConfig) *krm.SecurityConfig {
	if in == nil {
		return nil
	}
	out := &krm.SecurityConfig{}
	out.KerberosConfig = KerberosConfig_FromProto(mapCtx, in.GetKerberosConfig())
	out.IdentityConfig = IdentityConfig_FromProto(mapCtx, in.GetIdentityConfig())
	return out
}
func SecurityConfig_ToProto(mapCtx *direct.MapContext, in *krm.SecurityConfig) *pb.SecurityConfig {
	if in == nil {
		return nil
	}
	out := &pb.SecurityConfig{}
	out.KerberosConfig = KerberosConfig_ToProto(mapCtx, in.KerberosConfig)
	out.IdentityConfig = IdentityConfig_ToProto(mapCtx, in.IdentityConfig)
	return out
}
func ShieldedInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.ShieldedInstanceConfig) *krm.ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.ShieldedInstanceConfig{}
	out.EnableSecureBoot = in.EnableSecureBoot
	out.EnableVTPM = in.EnableVtpm
	out.EnableIntegrityMonitoring = in.EnableIntegrityMonitoring
	return out
}
func ShieldedInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.ShieldedInstanceConfig) *pb.ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ShieldedInstanceConfig{}
	out.EnableSecureBoot = in.EnableSecureBoot
	out.EnableVtpm = in.EnableVTPM
	out.EnableIntegrityMonitoring = in.EnableIntegrityMonitoring
	return out
}
func SoftwareConfig_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareConfig) *krm.SoftwareConfig {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareConfig{}
	out.ImageVersion = direct.LazyPtr(in.GetImageVersion())
	out.Properties = in.Properties
	out.OptionalComponents = direct.EnumSlice_FromProto(mapCtx, in.OptionalComponents)
	return out
}
func SoftwareConfig_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareConfig) *pb.SoftwareConfig {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareConfig{}
	out.ImageVersion = direct.ValueOf(in.ImageVersion)
	out.Properties = in.Properties
	out.OptionalComponents = direct.EnumSlice_ToProto[pb.Component](mapCtx, in.OptionalComponents)
	return out
}
func StartupConfig_FromProto(mapCtx *direct.MapContext, in *pb.StartupConfig) *krm.StartupConfig {
	if in == nil {
		return nil
	}
	out := &krm.StartupConfig{}
	out.RequiredRegistrationFraction = in.RequiredRegistrationFraction
	return out
}
func StartupConfig_ToProto(mapCtx *direct.MapContext, in *krm.StartupConfig) *pb.StartupConfig {
	if in == nil {
		return nil
	}
	out := &pb.StartupConfig{}
	out.RequiredRegistrationFraction = in.RequiredRegistrationFraction
	return out
}
func TemplateParameter_FromProto(mapCtx *direct.MapContext, in *pb.TemplateParameter) *krm.TemplateParameter {
	if in == nil {
		return nil
	}
	out := &krm.TemplateParameter{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Fields = in.Fields
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Validation = ParameterValidation_FromProto(mapCtx, in.GetValidation())
	return out
}
func TemplateParameter_ToProto(mapCtx *direct.MapContext, in *krm.TemplateParameter) *pb.TemplateParameter {
	if in == nil {
		return nil
	}
	out := &pb.TemplateParameter{}
	out.Name = direct.ValueOf(in.Name)
	out.Fields = in.Fields
	out.Description = direct.ValueOf(in.Description)
	out.Validation = ParameterValidation_ToProto(mapCtx, in.Validation)
	return out
}
func ValueValidation_FromProto(mapCtx *direct.MapContext, in *pb.ValueValidation) *krm.ValueValidation {
	if in == nil {
		return nil
	}
	out := &krm.ValueValidation{}
	out.Values = in.Values
	return out
}
func ValueValidation_ToProto(mapCtx *direct.MapContext, in *krm.ValueValidation) *pb.ValueValidation {
	if in == nil {
		return nil
	}
	out := &pb.ValueValidation{}
	out.Values = in.Values
	return out
}
func WorkflowTemplatePlacement_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowTemplatePlacement) *krm.WorkflowTemplatePlacement {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowTemplatePlacement{}
	out.ManagedCluster = ManagedCluster_FromProto(mapCtx, in.GetManagedCluster())
	out.ClusterSelector = ClusterSelector_FromProto(mapCtx, in.GetClusterSelector())
	return out
}
func WorkflowTemplatePlacement_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowTemplatePlacement) *pb.WorkflowTemplatePlacement {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowTemplatePlacement{}
	if oneof := ManagedCluster_ToProto(mapCtx, in.ManagedCluster); oneof != nil {
		out.Placement = &pb.WorkflowTemplatePlacement_ManagedCluster{ManagedCluster: oneof}
	}
	if oneof := ClusterSelector_ToProto(mapCtx, in.ClusterSelector); oneof != nil {
		out.Placement = &pb.WorkflowTemplatePlacement_ClusterSelector{ClusterSelector: oneof}
	}
	return out
}
func WorkflowTemplatePlacementObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowTemplatePlacement) *krm.WorkflowTemplatePlacementObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowTemplatePlacementObservedState{}
	out.ManagedCluster = ManagedClusterObservedState_FromProto(mapCtx, in.GetManagedCluster())
	// MISSING: ClusterSelector
	return out
}
func WorkflowTemplatePlacementObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowTemplatePlacementObservedState) *pb.WorkflowTemplatePlacement {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowTemplatePlacement{}
	if oneof := ManagedClusterObservedState_ToProto(mapCtx, in.ManagedCluster); oneof != nil {
		out.Placement = &pb.WorkflowTemplatePlacement_ManagedCluster{ManagedCluster: oneof}
	}
	// MISSING: ClusterSelector
	return out
}
func WorkflowTemplate_EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowTemplate_EncryptionConfig) *krm.WorkflowTemplate_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowTemplate_EncryptionConfig{}
	out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	return out
}
func WorkflowTemplate_EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowTemplate_EncryptionConfig) *pb.WorkflowTemplate_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowTemplate_EncryptionConfig{}
	out.KmsKey = direct.ValueOf(in.KMSKey)
	return out
}
