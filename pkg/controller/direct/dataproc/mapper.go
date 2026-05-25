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

package dataproc

import (
	"strconv"

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ClusterMetrics_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ClusterMetrics) *krm.ClusterMetrics {
	if in == nil {
		return nil
	}
	out := &krm.ClusterMetrics{}
	if in.HdfsMetrics != nil {
		out.HdfsMetrics = make(map[string]string)
		for k, v := range in.HdfsMetrics {
			out.HdfsMetrics[k] = strconv.FormatInt(v, 10)
		}
	}
	if in.YarnMetrics != nil {
		out.YarnMetrics = make(map[string]string)
		for k, v := range in.YarnMetrics {
			out.YarnMetrics[k] = strconv.FormatInt(v, 10)
		}
	}
	return out
}

func ClusterMetrics_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterMetrics) *pb.ClusterMetrics {
	if in == nil {
		return nil
	}
	out := &pb.ClusterMetrics{}
	if in.HdfsMetrics != nil {
		out.HdfsMetrics = make(map[string]int64)
		for k, v := range in.HdfsMetrics {
			if val, err := strconv.ParseInt(v, 10, 64); err == nil {
				out.HdfsMetrics[k] = val
			}
		}
	}
	if in.YarnMetrics != nil {
		out.YarnMetrics = make(map[string]int64)
		for k, v := range in.YarnMetrics {
			if val, err := strconv.ParseInt(v, 10, 64); err == nil {
				out.YarnMetrics[k] = val
			}
		}
	}
	return out
}

func DataprocMetricConfig_Metric_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.DataprocMetricConfig_Metric) *krm.DataprocMetricConfig_Metric {
	if in == nil {
		return nil
	}
	out := &krm.DataprocMetricConfig_Metric{}
	out.MetricOverrides = in.MetricOverrides
	if val := direct.Enum_FromProto(mapCtx, in.GetMetricSource()); val != nil {
		out.MetricSource = *val
	}
	return out
}

func DataprocMetricConfig_Metric_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DataprocMetricConfig_Metric) *pb.DataprocMetricConfig_Metric {
	if in == nil {
		return nil
	}
	out := &pb.DataprocMetricConfig_Metric{}
	out.MetricOverrides = in.MetricOverrides
	if in.MetricSource != "" {
		v := pb.DataprocMetricConfig_MetricSource(pb.DataprocMetricConfig_MetricSource_value[in.MetricSource])
		out.MetricSource = v
	}
	return out
}

func AcceleratorConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorConfig) *krm.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krm.AcceleratorConfig{}
	if in.AcceleratorCount != 0 {
		val := int64(in.AcceleratorCount)
		out.AcceleratorCount = &val
	}
	out.AcceleratorType = direct.LazyPtr(in.GetAcceleratorTypeUri())
	return out
}

func AcceleratorConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.AcceleratorConfig) *pb.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorConfig{}
	if in.AcceleratorCount != nil {
		out.AcceleratorCount = int32(*in.AcceleratorCount)
	}
	out.AcceleratorTypeUri = direct.ValueOf(in.AcceleratorType)
	return out
}

func DiskConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.DiskConfig) *krm.DiskConfig {
	if in == nil {
		return nil
	}
	out := &krm.DiskConfig{}
	if in.BootDiskSizeGb != 0 {
		val := int64(in.BootDiskSizeGb)
		out.BootDiskSizeGb = &val
	}
	out.BootDiskType = direct.LazyPtr(in.GetBootDiskType())
	out.LocalSsdInterface = direct.LazyPtr(in.GetLocalSsdInterface())
	if in.NumLocalSsds != 0 {
		val := int64(in.NumLocalSsds)
		out.NumLocalSsds = &val
	}
	return out
}

func DiskConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DiskConfig) *pb.DiskConfig {
	if in == nil {
		return nil
	}
	out := &pb.DiskConfig{}
	if in.BootDiskSizeGb != nil {
		out.BootDiskSizeGb = int32(*in.BootDiskSizeGb)
	}
	out.BootDiskType = direct.ValueOf(in.BootDiskType)
	out.LocalSsdInterface = direct.ValueOf(in.LocalSsdInterface)
	if in.NumLocalSsds != nil {
		out.NumLocalSsds = int32(*in.NumLocalSsds)
	}
	return out
}

func GkeNodePoolConfig_GkeNodeConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GkeNodePoolConfig_GkeNodeConfig) *krm.GkeNodePoolConfig_GkeNodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.GkeNodePoolConfig_GkeNodeConfig{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	if in.LocalSsdCount != 0 {
		val := int64(in.LocalSsdCount)
		out.LocalSsdCount = &val
	}
	out.Preemptible = direct.LazyPtr(in.GetPreemptible())
	out.Accelerators = direct.Slice_FromProto(mapCtx, in.Accelerators, GkeNodePoolConfig_GkeNodePoolAcceleratorConfig_v1beta1_FromProto)
	out.MinCpuPlatform = direct.LazyPtr(in.GetMinCpuPlatform())
	out.Spot = direct.LazyPtr(in.GetSpot())
	out.BootDiskKmsKey = direct.LazyPtr(in.GetBootDiskKmsKey())
	return out
}

func GkeNodePoolConfig_GkeNodeConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.GkeNodePoolConfig_GkeNodeConfig) *pb.GkeNodePoolConfig_GkeNodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.GkeNodePoolConfig_GkeNodeConfig{}
	out.MachineType = direct.ValueOf(in.MachineType)
	if in.LocalSsdCount != nil {
		out.LocalSsdCount = int32(*in.LocalSsdCount)
	}
	out.Preemptible = direct.ValueOf(in.Preemptible)
	out.Accelerators = direct.Slice_ToProto(mapCtx, in.Accelerators, GkeNodePoolConfig_GkeNodePoolAcceleratorConfig_v1beta1_ToProto)
	out.MinCpuPlatform = direct.ValueOf(in.MinCpuPlatform)
	out.Spot = direct.ValueOf(in.Spot)
	out.BootDiskKmsKey = direct.ValueOf(in.BootDiskKmsKey)
	return out
}

func GkeNodePoolConfig_GkeNodePoolAutoscalingConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GkeNodePoolConfig_GkeNodePoolAutoscalingConfig) *krm.GkeNodePoolConfig_GkeNodePoolAutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &krm.GkeNodePoolConfig_GkeNodePoolAutoscalingConfig{}
	if in.MinNodeCount != 0 {
		val := int64(in.MinNodeCount)
		out.MinNodeCount = &val
	}
	if in.MaxNodeCount != 0 {
		val := int64(in.MaxNodeCount)
		out.MaxNodeCount = &val
	}
	return out
}

func GkeNodePoolConfig_GkeNodePoolAutoscalingConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.GkeNodePoolConfig_GkeNodePoolAutoscalingConfig) *pb.GkeNodePoolConfig_GkeNodePoolAutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &pb.GkeNodePoolConfig_GkeNodePoolAutoscalingConfig{}
	if in.MinNodeCount != nil {
		out.MinNodeCount = int32(*in.MinNodeCount)
	}
	if in.MaxNodeCount != nil {
		out.MaxNodeCount = int32(*in.MaxNodeCount)
	}
	return out
}

func InstanceGroupConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupConfig) *krm.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupConfig{}
	if in.NumInstances != 0 {
		val := int64(in.NumInstances)
		out.NumInstances = &val
	}
	out.DiskConfig = DiskConfig_v1beta1_FromProto(mapCtx, in.GetDiskConfig())
	out.Preemptibility = direct.Enum_FromProto(mapCtx, in.GetPreemptibility())
	out.Accelerators = direct.Slice_FromProto(mapCtx, in.Accelerators, AcceleratorConfig_v1beta1_FromProto)
	out.MinCpuPlatform = direct.LazyPtr(in.GetMinCpuPlatform())
	return out
}

func InstanceGroupConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupConfig) *pb.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupConfig{}
	if in.NumInstances != nil {
		out.NumInstances = int32(*in.NumInstances)
	}
	out.DiskConfig = DiskConfig_v1beta1_ToProto(mapCtx, in.DiskConfig)
	out.Preemptibility = direct.Enum_ToProto[pb.InstanceGroupConfig_Preemptibility](mapCtx, in.Preemptibility)
	out.Accelerators = direct.Slice_ToProto(mapCtx, in.Accelerators, AcceleratorConfig_v1beta1_ToProto)
	out.MinCpuPlatform = direct.ValueOf(in.MinCpuPlatform)
	return out
}

func VirtualClusterConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.VirtualClusterConfig) *krm.VirtualClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.VirtualClusterConfig{}
	out.AuxiliaryServicesConfig = AuxiliaryServicesConfig_v1beta1_FromProto(mapCtx, in.GetAuxiliaryServicesConfig())
	if in.GetKubernetesClusterConfig() != nil {
		if cfg := KubernetesClusterConfig_v1beta1_FromProto(mapCtx, in.GetKubernetesClusterConfig()); cfg != nil {
			out.KubernetesClusterConfig = *cfg
		}
	}
	return out
}

func KubernetesClusterConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesClusterConfig) *krm.KubernetesClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.KubernetesClusterConfig{}
	if in.GetGkeClusterConfig() != nil {
		if cfg := GkeClusterConfig_v1beta1_FromProto(mapCtx, in.GetGkeClusterConfig()); cfg != nil {
			out.GkeClusterConfig = *cfg
		}
	}
	out.KubernetesNamespace = direct.LazyPtr(in.GetKubernetesNamespace())
	out.KubernetesSoftwareConfig = KubernetesSoftwareConfig_v1beta1_FromProto(mapCtx, in.GetKubernetesSoftwareConfig())
	return out
}

func NodeInitializationAction_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NodeInitializationAction) *krm.NodeInitializationAction {
	if in == nil {
		return nil
	}
	out := &krm.NodeInitializationAction{}
	out.ExecutableFile = in.GetExecutableFile()
	out.ExecutionTimeout = direct.StringDuration_FromProto(mapCtx, in.GetExecutionTimeout())
	return out
}

func NodeInitializationAction_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.NodeInitializationAction) *pb.NodeInitializationAction {
	if in == nil {
		return nil
	}
	out := &pb.NodeInitializationAction{}
	out.ExecutableFile = in.ExecutableFile
	out.ExecutionTimeout = direct.StringDuration_ToProto(mapCtx, in.ExecutionTimeout)
	return out
}

func VirtualClusterConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.VirtualClusterConfig) *pb.VirtualClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.VirtualClusterConfig{}
	out.AuxiliaryServicesConfig = AuxiliaryServicesConfig_v1beta1_ToProto(mapCtx, in.AuxiliaryServicesConfig)
	if oneof := KubernetesClusterConfig_v1beta1_ToProto(mapCtx, &in.KubernetesClusterConfig); oneof != nil {
		out.InfrastructureConfig = &pb.VirtualClusterConfig_KubernetesClusterConfig{KubernetesClusterConfig: oneof}
	}
	return out
}

func KubernetesClusterConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.KubernetesClusterConfig) *pb.KubernetesClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesClusterConfig{}
	if oneof := GkeClusterConfig_v1beta1_ToProto(mapCtx, &in.GkeClusterConfig); oneof != nil {
		out.Config = &pb.KubernetesClusterConfig_GkeClusterConfig{GkeClusterConfig: oneof}
	}
	out.KubernetesNamespace = direct.ValueOf(in.KubernetesNamespace)
	out.KubernetesSoftwareConfig = KubernetesSoftwareConfig_v1beta1_ToProto(mapCtx, in.KubernetesSoftwareConfig)
	return out
}

func DataprocClusterSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.DataprocClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataprocClusterSpec{}
	out.Config = ClusterConfig_v1beta1_FromProto(mapCtx, in.GetConfig())
	out.VirtualClusterConfig = VirtualClusterConfig_v1beta1_FromProto(mapCtx, in.GetVirtualClusterConfig())
	return out
}

func DataprocClusterSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.DataprocClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.Config = ClusterConfig_v1beta1_ToProto(mapCtx, in.Config)
	out.VirtualClusterConfig = VirtualClusterConfig_v1beta1_ToProto(mapCtx, in.VirtualClusterConfig)
	return out
}

func ClusterConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ClusterConfig) *krm.ClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterConfig{}
	if in.GetConfigBucket() != "" {
		out.StagingBucketRef = &krm.DataprocStagingBucketRef{External: in.GetConfigBucket()}
	}
	if in.GetTempBucket() != "" {
		out.TempBucketRef = &krm.DataprocTempBucketRef{External: in.GetTempBucket()}
	}
	out.GceClusterConfig = GceClusterConfig_v1beta1_FromProto(mapCtx, in.GetGceClusterConfig())
	out.MasterConfig = InstanceGroupConfig_v1beta1_FromProto(mapCtx, in.GetMasterConfig())
	out.WorkerConfig = InstanceGroupConfig_v1beta1_FromProto(mapCtx, in.GetWorkerConfig())
	out.SecondaryWorkerConfig = InstanceGroupConfig_v1beta1_FromProto(mapCtx, in.GetSecondaryWorkerConfig())
	out.SoftwareConfig = SoftwareConfig_v1beta1_FromProto(mapCtx, in.GetSoftwareConfig())
	out.InitializationActions = direct.Slice_FromProto(mapCtx, in.InitializationActions, NodeInitializationAction_v1beta1_FromProto)
	out.EncryptionConfig = EncryptionConfig_v1beta1_FromProto(mapCtx, in.GetEncryptionConfig())
	out.AutoscalingConfig = AutoscalingConfig_v1beta1_FromProto(mapCtx, in.GetAutoscalingConfig())
	out.SecurityConfig = SecurityConfig_v1beta1_FromProto(mapCtx, in.GetSecurityConfig())
	out.LifecycleConfig = LifecycleConfig_v1beta1_FromProto(mapCtx, in.GetLifecycleConfig())
	out.EndpointConfig = EndpointConfig_v1beta1_FromProto(mapCtx, in.GetEndpointConfig())
	out.MetastoreConfig = MetastoreConfig_v1beta1_FromProto(mapCtx, in.GetMetastoreConfig())
	out.DataprocMetricConfig = DataprocMetricConfig_v1beta1_FromProto(mapCtx, in.GetDataprocMetricConfig())
	return out
}

func ClusterConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterConfig) *pb.ClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.ClusterConfig{}
	if in.StagingBucketRef != nil {
		out.ConfigBucket = in.StagingBucketRef.External
	}
	if in.TempBucketRef != nil {
		out.TempBucket = in.TempBucketRef.External
	}
	out.GceClusterConfig = GceClusterConfig_v1beta1_ToProto(mapCtx, in.GceClusterConfig)
	out.MasterConfig = InstanceGroupConfig_v1beta1_ToProto(mapCtx, in.MasterConfig)
	out.WorkerConfig = InstanceGroupConfig_v1beta1_ToProto(mapCtx, in.WorkerConfig)
	out.SecondaryWorkerConfig = InstanceGroupConfig_v1beta1_ToProto(mapCtx, in.SecondaryWorkerConfig)
	out.SoftwareConfig = SoftwareConfig_v1beta1_ToProto(mapCtx, in.SoftwareConfig)
	out.InitializationActions = direct.Slice_ToProto(mapCtx, in.InitializationActions, NodeInitializationAction_v1beta1_ToProto)
	out.EncryptionConfig = EncryptionConfig_v1beta1_ToProto(mapCtx, in.EncryptionConfig)
	out.AutoscalingConfig = AutoscalingConfig_v1beta1_ToProto(mapCtx, in.AutoscalingConfig)
	out.SecurityConfig = SecurityConfig_v1beta1_ToProto(mapCtx, in.SecurityConfig)
	out.LifecycleConfig = LifecycleConfig_v1beta1_ToProto(mapCtx, in.LifecycleConfig)
	out.EndpointConfig = EndpointConfig_v1beta1_ToProto(mapCtx, in.EndpointConfig)
	out.MetastoreConfig = MetastoreConfig_v1beta1_ToProto(mapCtx, in.MetastoreConfig)
	out.DataprocMetricConfig = DataprocMetricConfig_v1beta1_ToProto(mapCtx, in.DataprocMetricConfig)
	return out
}

func GceClusterConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GceClusterConfig) *krm.GceClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.GceClusterConfig{}
	out.Zone = direct.LazyPtr(in.GetZoneUri())
	if in.GetNetworkUri() != "" {
		out.NetworkRef = &krm.DataprocComputeNetworkRef{External: in.GetNetworkUri()}
	}
	if in.GetSubnetworkUri() != "" {
		out.SubnetworkRef = &krm.DataprocComputeSubnetworkRef{External: in.GetSubnetworkUri()}
	}
	out.InternalIpOnly = in.InternalIpOnly
	out.PrivateIpv6GoogleAccess = direct.Enum_FromProto(mapCtx, in.GetPrivateIpv6GoogleAccess())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &krm.DataprocIamServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.ServiceAccountScopes = in.ServiceAccountScopes
	out.Tags = in.Tags
	out.Metadata = in.Metadata
	out.ReservationAffinity = ReservationAffinity_v1beta1_FromProto(mapCtx, in.GetReservationAffinity())
	out.NodeGroupAffinity = NodeGroupAffinity_v1beta1_FromProto(mapCtx, in.GetNodeGroupAffinity())
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_v1beta1_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.ConfidentialInstanceConfig = ConfidentialInstanceConfig_v1beta1_FromProto(mapCtx, in.GetConfidentialInstanceConfig())
	return out
}

func GceClusterConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.GceClusterConfig) *pb.GceClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.GceClusterConfig{}
	out.ZoneUri = direct.ValueOf(in.Zone)
	if in.NetworkRef != nil {
		out.NetworkUri = in.NetworkRef.External
	}
	if in.SubnetworkRef != nil {
		out.SubnetworkUri = in.SubnetworkRef.External
	}
	out.InternalIpOnly = in.InternalIpOnly
	out.PrivateIpv6GoogleAccess = direct.Enum_ToProto[pb.GceClusterConfig_PrivateIpv6GoogleAccess](mapCtx, in.PrivateIpv6GoogleAccess)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.ServiceAccountScopes = in.ServiceAccountScopes
	out.Tags = in.Tags
	out.Metadata = in.Metadata
	out.ReservationAffinity = ReservationAffinity_v1beta1_ToProto(mapCtx, in.ReservationAffinity)
	out.NodeGroupAffinity = NodeGroupAffinity_v1beta1_ToProto(mapCtx, in.NodeGroupAffinity)
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_v1beta1_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.ConfidentialInstanceConfig = ConfidentialInstanceConfig_v1beta1_ToProto(mapCtx, in.ConfidentialInstanceConfig)
	return out
}

func AutoscalingConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingConfig) *krm.AutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingConfig{}
	if in.GetPolicyUri() != "" {
		out.PolicyRef = &krm.DataprocAutoscalingPolicyRef{External: in.GetPolicyUri()}
	}
	return out
}

func AutoscalingConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingConfig) *pb.AutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingConfig{}
	if in.PolicyRef != nil {
		out.PolicyUri = in.PolicyRef.External
	}
	return out
}

func MetastoreConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.MetastoreConfig) *krm.MetastoreConfig {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreConfig{}
	if in.GetDataprocMetastoreService() != "" {
		out.DataprocMetastoreServiceRef = &krm.DataprocMetastoreServiceRef{External: in.GetDataprocMetastoreService()}
	}
	return out
}

func MetastoreConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreConfig) *pb.MetastoreConfig {
	if in == nil {
		return nil
	}
	out := &pb.MetastoreConfig{}
	if in.DataprocMetastoreServiceRef != nil {
		out.DataprocMetastoreService = in.DataprocMetastoreServiceRef.External
	}
	return out
}

func GkeNodePoolTarget_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GkeNodePoolTarget) *krm.GkeNodePoolTarget {
	if in == nil {
		return nil
	}
	out := &krm.GkeNodePoolTarget{}
	if in.GetNodePool() != "" {
		out.NodePoolRef = &krm.DataprocContainerNodePoolRef{External: in.GetNodePool()}
	}
	out.Roles = direct.EnumSlice_FromProto(mapCtx, in.Roles)
	out.NodePoolConfig = GkeNodePoolConfig_v1beta1_FromProto(mapCtx, in.GetNodePoolConfig())
	return out
}

func GkeNodePoolTarget_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.GkeNodePoolTarget) *pb.GkeNodePoolTarget {
	if in == nil {
		return nil
	}
	out := &pb.GkeNodePoolTarget{}
	if in.NodePoolRef != nil {
		out.NodePool = in.NodePoolRef.External
	}
	out.Roles = direct.EnumSlice_ToProto[pb.GkeNodePoolTarget_Role](mapCtx, in.Roles)
	out.NodePoolConfig = GkeNodePoolConfig_v1beta1_ToProto(mapCtx, in.NodePoolConfig)
	return out
}

func GkeClusterConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GkeClusterConfig) *krm.GkeClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.GkeClusterConfig{}
	if in.GetGkeClusterTarget() != "" {
		out.GkeClusterTargetRef = &krm.DataprocContainerClusterRef{External: in.GetGkeClusterTarget()}
	}
	out.NodePoolTarget = direct.Slice_FromProto(mapCtx, in.NodePoolTarget, GkeNodePoolTarget_v1beta1_FromProto)
	return out
}

func GkeClusterConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.GkeClusterConfig) *pb.GkeClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.GkeClusterConfig{}
	if in.GkeClusterTargetRef != nil {
		out.GkeClusterTarget = in.GkeClusterTargetRef.External
	}
	out.NodePoolTarget = direct.Slice_ToProto(mapCtx, in.NodePoolTarget, GkeNodePoolTarget_v1beta1_ToProto)
	return out
}

func SparkHistoryServerConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SparkHistoryServerConfig) *krm.SparkHistoryServerConfig {
	if in == nil {
		return nil
	}
	out := &krm.SparkHistoryServerConfig{}
	if in.GetDataprocCluster() != "" {
		out.DataprocClusterRef = &krm.DataprocClusterRef{External: in.GetDataprocCluster()}
	}
	return out
}

func SparkHistoryServerConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SparkHistoryServerConfig) *pb.SparkHistoryServerConfig {
	if in == nil {
		return nil
	}
	out := &pb.SparkHistoryServerConfig{}
	if in.DataprocClusterRef != nil {
		out.DataprocCluster = in.DataprocClusterRef.External
	}
	return out
}

func AuxiliaryServicesConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AuxiliaryServicesConfig) *krm.AuxiliaryServicesConfig {
	if in == nil {
		return nil
	}
	out := &krm.AuxiliaryServicesConfig{}
	out.MetastoreConfig = MetastoreConfig_v1beta1_FromProto(mapCtx, in.GetMetastoreConfig())
	out.SparkHistoryServerConfig = SparkHistoryServerConfig_v1beta1_FromProto(mapCtx, in.GetSparkHistoryServerConfig())
	return out
}

func AuxiliaryServicesConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.AuxiliaryServicesConfig) *pb.AuxiliaryServicesConfig {
	if in == nil {
		return nil
	}
	out := &pb.AuxiliaryServicesConfig{}
	out.MetastoreConfig = MetastoreConfig_v1beta1_ToProto(mapCtx, in.MetastoreConfig)
	out.SparkHistoryServerConfig = SparkHistoryServerConfig_v1beta1_ToProto(mapCtx, in.SparkHistoryServerConfig)
	return out
}

func KerberosConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.KerberosConfig) *krm.KerberosConfig {
	if in == nil {
		return nil
	}
	out := &krm.KerberosConfig{}
	out.EnableKerberos = direct.LazyPtr(in.GetEnableKerberos())
	out.CrossRealmTrustRealm = direct.LazyPtr(in.GetCrossRealmTrustRealm())
	out.CrossRealmTrustKdc = direct.LazyPtr(in.GetCrossRealmTrustKdc())
	out.CrossRealmTrustAdminServer = direct.LazyPtr(in.GetCrossRealmTrustAdminServer())
	if in.TgtLifetimeHours != 0 {
		val := int64(in.TgtLifetimeHours)
		out.TgtLifetimeHours = &val
	}
	out.Realm = direct.LazyPtr(in.GetRealm())
	if in.GetKmsKeyUri() != "" {
		out.KmsKeyRef = &krm.DataprocKmsKeyRef{External: in.GetKmsKeyUri()}
	}
	return out
}

func KerberosConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.KerberosConfig) *pb.KerberosConfig {
	if in == nil {
		return nil
	}
	out := &pb.KerberosConfig{}
	out.EnableKerberos = direct.ValueOf(in.EnableKerberos)
	out.CrossRealmTrustRealm = direct.ValueOf(in.CrossRealmTrustRealm)
	out.CrossRealmTrustKdc = direct.ValueOf(in.CrossRealmTrustKdc)
	out.CrossRealmTrustAdminServer = direct.ValueOf(in.CrossRealmTrustAdminServer)
	if in.TgtLifetimeHours != nil {
		out.TgtLifetimeHours = int32(*in.TgtLifetimeHours)
	}
	out.Realm = direct.ValueOf(in.Realm)
	if in.KmsKeyRef != nil {
		out.KmsKeyUri = in.KmsKeyRef.External
	}
	return out
}

func EncryptionConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfig{}
	if in.GetGcePdKmsKeyName() != "" {
		out.GcePdKmsKeyRef = &krm.DataprocPdKmsKeyRef{External: in.GetGcePdKmsKeyName()}
	}
	return out
}

func EncryptionConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfig) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	if in.GcePdKmsKeyRef != nil {
		out.GcePdKmsKeyName = in.GcePdKmsKeyRef.External
	}
	return out
}
