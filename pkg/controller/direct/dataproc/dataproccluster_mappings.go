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

	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	containerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1beta1"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ClusterAccelerators_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorConfig) *krm.ClusterAccelerators {
	if in == nil {
		return nil
	}
	out := &krm.ClusterAccelerators{}
	out.AcceleratorCount = direct.LazyPtr(int64(in.GetAcceleratorCount()))
	out.AcceleratorType = direct.LazyPtr(in.GetAcceleratorTypeUri())
	return out
}

func ClusterAccelerators_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterAccelerators) *pb.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorConfig{}
	out.AcceleratorCount = int32(direct.ValueOf(in.AcceleratorCount))
	out.AcceleratorTypeUri = direct.ValueOf(in.AcceleratorType)
	return out
}

func ClusterAutoscaling_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GkeNodePoolConfig_GkeNodePoolAutoscalingConfig) *krm.ClusterAutoscaling {
	if in == nil {
		return nil
	}
	out := &krm.ClusterAutoscaling{}
	out.MinNodeCount = direct.LazyPtr(int64(in.GetMinNodeCount()))
	out.MaxNodeCount = direct.LazyPtr(int64(in.GetMaxNodeCount()))
	return out
}

func ClusterAutoscaling_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterAutoscaling) *pb.GkeNodePoolConfig_GkeNodePoolAutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &pb.GkeNodePoolConfig_GkeNodePoolAutoscalingConfig{}
	out.MinNodeCount = int32(direct.ValueOf(in.MinNodeCount))
	out.MaxNodeCount = int32(direct.ValueOf(in.MaxNodeCount))
	return out
}

func ClusterDiskConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.DiskConfig) *krm.ClusterDiskConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterDiskConfig{}
	out.BootDiskType = direct.LazyPtr(in.GetBootDiskType())
	out.BootDiskSizeGb = direct.LazyPtr(int64(in.GetBootDiskSizeGb()))
	out.NumLocalSsds = direct.LazyPtr(int64(in.GetNumLocalSsds()))
	out.LocalSsdInterface = direct.LazyPtr(in.GetLocalSsdInterface())
	return out
}

func ClusterDiskConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterDiskConfig) *pb.DiskConfig {
	if in == nil {
		return nil
	}
	out := &pb.DiskConfig{}
	out.BootDiskType = direct.ValueOf(in.BootDiskType)
	out.BootDiskSizeGb = int32(direct.ValueOf(in.BootDiskSizeGb))
	out.NumLocalSsds = int32(direct.ValueOf(in.NumLocalSsds))
	out.LocalSsdInterface = direct.ValueOf(in.LocalSsdInterface)
	return out
}

func ClusterGkeNodeConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GkeNodePoolConfig_GkeNodeConfig) *krm.ClusterGkeNodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterGkeNodeConfig{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.LocalSsdCount = direct.LazyPtr(int64(in.GetLocalSsdCount()))
	out.Preemptible = direct.LazyPtr(in.GetPreemptible())
	out.Accelerators = direct.Slice_FromProto(mapCtx, in.Accelerators, ClusterGkeNodePoolAccelerators_v1beta1_FromProto)
	out.MinCpuPlatform = direct.LazyPtr(in.GetMinCpuPlatform())
	out.BootDiskKmsKey = direct.LazyPtr(in.GetBootDiskKmsKey())
	out.Spot = direct.LazyPtr(in.GetSpot())
	return out
}

func ClusterGkeNodeConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterGkeNodeConfig) *pb.GkeNodePoolConfig_GkeNodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.GkeNodePoolConfig_GkeNodeConfig{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.LocalSsdCount = int32(direct.ValueOf(in.LocalSsdCount))
	out.Preemptible = direct.ValueOf(in.Preemptible)
	out.Accelerators = direct.Slice_ToProto(mapCtx, in.Accelerators, ClusterGkeNodePoolAccelerators_v1beta1_ToProto)
	out.MinCpuPlatform = direct.ValueOf(in.MinCpuPlatform)
	out.BootDiskKmsKey = direct.ValueOf(in.BootDiskKmsKey)
	out.Spot = direct.ValueOf(in.Spot)
	return out
}

func ClusterGkeNodePoolAccelerators_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GkeNodePoolConfig_GkeNodePoolAcceleratorConfig) *krm.ClusterGkeNodePoolAccelerators {
	if in == nil {
		return nil
	}
	out := &krm.ClusterGkeNodePoolAccelerators{}
	out.AcceleratorCount = direct.LazyPtr(in.GetAcceleratorCount())
	out.AcceleratorType = direct.LazyPtr(in.GetAcceleratorType())
	out.GpuPartitionSize = direct.LazyPtr(in.GetGpuPartitionSize())
	return out
}

func ClusterGkeNodePoolAccelerators_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterGkeNodePoolAccelerators) *pb.GkeNodePoolConfig_GkeNodePoolAcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.GkeNodePoolConfig_GkeNodePoolAcceleratorConfig{}
	out.AcceleratorCount = direct.ValueOf(in.AcceleratorCount)
	out.AcceleratorType = direct.ValueOf(in.AcceleratorType)
	out.GpuPartitionSize = direct.ValueOf(in.GpuPartitionSize)
	return out
}

func ClusterWorkerConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupConfig) *krm.ClusterWorkerConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterWorkerConfig{}
	out.Accelerators = direct.Slice_FromProto(mapCtx, in.Accelerators, ClusterAccelerators_v1beta1_FromProto)
	out.DiskConfig = ClusterDiskConfig_v1beta1_FromProto(mapCtx, in.DiskConfig)
	if in.GetImageUri() != "" {
		out.ImageRef = &computev1beta1.ComputeImageRef{External: in.GetImageUri()}
	}
	out.MachineType = direct.LazyPtr(in.GetMachineTypeUri())
	out.MinCpuPlatform = direct.LazyPtr(in.GetMinCpuPlatform())
	out.NumInstances = direct.LazyPtr(int64(in.GetNumInstances()))
	out.Preemptibility = direct.LazyPtr(in.GetPreemptibility().String())
	return out
}

func ClusterWorkerConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterWorkerConfig) *pb.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupConfig{}
	out.Accelerators = direct.Slice_ToProto(mapCtx, in.Accelerators, ClusterAccelerators_v1beta1_ToProto)
	out.DiskConfig = ClusterDiskConfig_v1beta1_ToProto(mapCtx, in.DiskConfig)
	if in.ImageRef != nil {
		out.ImageUri = in.ImageRef.External
	}
	out.MachineTypeUri = direct.ValueOf(in.MachineType)
	out.MinCpuPlatform = direct.ValueOf(in.MinCpuPlatform)
	out.NumInstances = int32(direct.ValueOf(in.NumInstances))
	if in.Preemptibility != nil {
		if val, ok := pb.InstanceGroupConfig_Preemptibility_value[*in.Preemptibility]; ok {
			out.Preemptibility = pb.InstanceGroupConfig_Preemptibility(val)
		}
	}
	return out
}

func ClusterMasterConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupConfig) *krm.ClusterMasterConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterMasterConfig{}
	out.Accelerators = direct.Slice_FromProto(mapCtx, in.Accelerators, ClusterAccelerators_v1beta1_FromProto)
	out.DiskConfig = ClusterDiskConfig_v1beta1_FromProto(mapCtx, in.DiskConfig)
	if in.GetImageUri() != "" {
		out.ImageRef = &computev1beta1.ComputeImageRef{External: in.GetImageUri()}
	}
	out.MachineType = direct.LazyPtr(in.GetMachineTypeUri())
	out.MinCpuPlatform = direct.LazyPtr(in.GetMinCpuPlatform())
	out.NumInstances = direct.LazyPtr(int64(in.GetNumInstances()))
	out.Preemptibility = direct.LazyPtr(in.GetPreemptibility().String())
	return out
}

func ClusterMasterConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterMasterConfig) *pb.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupConfig{}
	out.Accelerators = direct.Slice_ToProto(mapCtx, in.Accelerators, ClusterAccelerators_v1beta1_ToProto)
	out.DiskConfig = ClusterDiskConfig_v1beta1_ToProto(mapCtx, in.DiskConfig)
	if in.ImageRef != nil {
		out.ImageUri = in.ImageRef.External
	}
	out.MachineTypeUri = direct.ValueOf(in.MachineType)
	out.MinCpuPlatform = direct.ValueOf(in.MinCpuPlatform)
	out.NumInstances = int32(direct.ValueOf(in.NumInstances))
	if in.Preemptibility != nil {
		if val, ok := pb.InstanceGroupConfig_Preemptibility_value[*in.Preemptibility]; ok {
			out.Preemptibility = pb.InstanceGroupConfig_Preemptibility(val)
		}
	}
	return out
}

func ClusterSecondaryWorkerConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupConfig) *krm.ClusterSecondaryWorkerConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterSecondaryWorkerConfig{}
	out.Accelerators = direct.Slice_FromProto(mapCtx, in.Accelerators, ClusterAccelerators_v1beta1_FromProto)
	out.DiskConfig = ClusterDiskConfig_v1beta1_FromProto(mapCtx, in.DiskConfig)
	if in.GetImageUri() != "" {
		out.ImageRef = &computev1beta1.ComputeImageRef{External: in.GetImageUri()}
	}
	out.MachineType = direct.LazyPtr(in.GetMachineTypeUri())
	out.MinCpuPlatform = direct.LazyPtr(in.GetMinCpuPlatform())
	out.NumInstances = direct.LazyPtr(int64(in.GetNumInstances()))
	out.Preemptibility = direct.LazyPtr(in.GetPreemptibility().String())
	return out
}

func ClusterSecondaryWorkerConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterSecondaryWorkerConfig) *pb.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupConfig{}
	out.Accelerators = direct.Slice_ToProto(mapCtx, in.Accelerators, ClusterAccelerators_v1beta1_ToProto)
	out.DiskConfig = ClusterDiskConfig_v1beta1_ToProto(mapCtx, in.DiskConfig)
	if in.ImageRef != nil {
		out.ImageUri = in.ImageRef.External
	}
	out.MachineTypeUri = direct.ValueOf(in.MachineType)
	out.MinCpuPlatform = direct.ValueOf(in.MinCpuPlatform)
	out.NumInstances = int32(direct.ValueOf(in.NumInstances))
	if in.Preemptibility != nil {
		if val, ok := pb.InstanceGroupConfig_Preemptibility_value[*in.Preemptibility]; ok {
			out.Preemptibility = pb.InstanceGroupConfig_Preemptibility(val)
		}
	}
	return out
}

func ClusterMasterConfigStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupConfig) *krm.ClusterMasterConfigStatus {
	if in == nil {
		return nil
	}
	out := &krm.ClusterMasterConfigStatus{}
	out.InstanceNames = in.InstanceNames
	out.InstanceReferences = direct.Slice_FromProto(mapCtx, in.InstanceReferences, ClusterInstanceReferencesStatus_v1beta1_FromProto)
	out.IsPreemptible = direct.LazyPtr(in.IsPreemptible)
	out.ManagedGroupConfig = ClusterManagedGroupConfigStatus_v1beta1_FromProto(mapCtx, in.ManagedGroupConfig)
	return out
}

func ClusterMasterConfigStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterMasterConfigStatus) *pb.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupConfig{}
	out.InstanceNames = in.InstanceNames
	out.InstanceReferences = direct.Slice_ToProto(mapCtx, in.InstanceReferences, ClusterInstanceReferencesStatus_v1beta1_ToProto)
	out.IsPreemptible = direct.ValueOf(in.IsPreemptible)
	out.ManagedGroupConfig = ClusterManagedGroupConfigStatus_v1beta1_ToProto(mapCtx, in.ManagedGroupConfig)
	return out
}

func ClusterWorkerConfigStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupConfig) *krm.ClusterWorkerConfigStatus {
	if in == nil {
		return nil
	}
	out := &krm.ClusterWorkerConfigStatus{}
	out.InstanceNames = in.InstanceNames
	out.InstanceReferences = direct.Slice_FromProto(mapCtx, in.InstanceReferences, ClusterInstanceReferencesStatus_v1beta1_FromProto)
	out.IsPreemptible = direct.LazyPtr(in.IsPreemptible)
	out.ManagedGroupConfig = ClusterManagedGroupConfigStatus_v1beta1_FromProto(mapCtx, in.ManagedGroupConfig)
	return out
}

func ClusterWorkerConfigStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterWorkerConfigStatus) *pb.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupConfig{}
	out.InstanceNames = in.InstanceNames
	out.InstanceReferences = direct.Slice_ToProto(mapCtx, in.InstanceReferences, ClusterInstanceReferencesStatus_v1beta1_ToProto)
	out.IsPreemptible = direct.ValueOf(in.IsPreemptible)
	out.ManagedGroupConfig = ClusterManagedGroupConfigStatus_v1beta1_ToProto(mapCtx, in.ManagedGroupConfig)
	return out
}

func ClusterSecondaryWorkerConfigStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupConfig) *krm.ClusterSecondaryWorkerConfigStatus {
	if in == nil {
		return nil
	}
	out := &krm.ClusterSecondaryWorkerConfigStatus{}
	out.InstanceNames = in.InstanceNames
	out.InstanceReferences = direct.Slice_FromProto(mapCtx, in.InstanceReferences, ClusterInstanceReferencesStatus_v1beta1_FromProto)
	out.IsPreemptible = direct.LazyPtr(in.IsPreemptible)
	out.ManagedGroupConfig = ClusterManagedGroupConfigStatus_v1beta1_FromProto(mapCtx, in.ManagedGroupConfig)
	return out
}

func ClusterSecondaryWorkerConfigStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterSecondaryWorkerConfigStatus) *pb.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupConfig{}
	out.InstanceNames = in.InstanceNames
	out.InstanceReferences = direct.Slice_ToProto(mapCtx, in.InstanceReferences, ClusterInstanceReferencesStatus_v1beta1_ToProto)
	out.IsPreemptible = direct.ValueOf(in.IsPreemptible)
	out.ManagedGroupConfig = ClusterManagedGroupConfigStatus_v1beta1_ToProto(mapCtx, in.ManagedGroupConfig)
	return out
}

func ClusterInitializationActions_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NodeInitializationAction) *krm.ClusterInitializationActions {
	if in == nil {
		return nil
	}
	out := &krm.ClusterInitializationActions{}
	out.ExecutableFile = in.GetExecutableFile()
	out.ExecutionTimeout = direct.StringDuration_FromProto(mapCtx, in.GetExecutionTimeout())
	return out
}

func ClusterInitializationActions_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterInitializationActions) *pb.NodeInitializationAction {
	if in == nil {
		return nil
	}
	out := &pb.NodeInitializationAction{}
	out.ExecutableFile = in.ExecutableFile
	out.ExecutionTimeout = direct.StringDuration_ToProto(mapCtx, in.ExecutionTimeout)
	return out
}

func ClusterKerberosConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.KerberosConfig) *krm.ClusterKerberosConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterKerberosConfig{}
	out.CrossRealmTrustAdminServer = direct.LazyPtr(in.GetCrossRealmTrustAdminServer())
	out.CrossRealmTrustKdc = direct.LazyPtr(in.GetCrossRealmTrustKdc())
	out.CrossRealmTrustRealm = direct.LazyPtr(in.GetCrossRealmTrustRealm())
	out.CrossRealmTrustSharedPassword = direct.LazyPtr(in.GetCrossRealmTrustSharedPasswordUri())
	out.EnableKerberos = direct.LazyPtr(in.GetEnableKerberos())
	out.KdcDbKey = direct.LazyPtr(in.GetKdcDbKeyUri())
	out.KeyPassword = direct.LazyPtr(in.GetKeyPasswordUri())
	out.Keystore = direct.LazyPtr(in.GetKeystoreUri())
	out.KeystorePassword = direct.LazyPtr(in.GetKeystorePasswordUri())
	if in.GetKmsKeyUri() != "" {
		out.KmsKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyUri()}
	}
	out.Realm = direct.LazyPtr(in.GetRealm())
	out.RootPrincipalPassword = direct.LazyPtr(in.GetRootPrincipalPasswordUri())
	out.TgtLifetimeHours = direct.LazyPtr(int64(in.GetTgtLifetimeHours()))
	out.Truststore = direct.LazyPtr(in.GetTruststoreUri())
	out.TruststorePassword = direct.LazyPtr(in.GetTruststorePasswordUri())
	return out
}

func ClusterKerberosConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterKerberosConfig) *pb.KerberosConfig {
	if in == nil {
		return nil
	}
	out := &pb.KerberosConfig{}
	out.CrossRealmTrustAdminServer = direct.ValueOf(in.CrossRealmTrustAdminServer)
	out.CrossRealmTrustKdc = direct.ValueOf(in.CrossRealmTrustKdc)
	out.CrossRealmTrustRealm = direct.ValueOf(in.CrossRealmTrustRealm)
	out.CrossRealmTrustSharedPasswordUri = direct.ValueOf(in.CrossRealmTrustSharedPassword)
	out.EnableKerberos = direct.ValueOf(in.EnableKerberos)
	out.KdcDbKeyUri = direct.ValueOf(in.KdcDbKey)
	out.KeyPasswordUri = direct.ValueOf(in.KeyPassword)
	out.KeystoreUri = direct.ValueOf(in.Keystore)
	out.KeystorePasswordUri = direct.ValueOf(in.KeystorePassword)
	if in.KmsKeyRef != nil {
		out.KmsKeyUri = in.KmsKeyRef.External
	}
	out.Realm = direct.ValueOf(in.Realm)
	out.RootPrincipalPasswordUri = direct.ValueOf(in.RootPrincipalPassword)
	out.TgtLifetimeHours = int32(direct.ValueOf(in.TgtLifetimeHours))
	out.TruststoreUri = direct.ValueOf(in.Truststore)
	out.TruststorePasswordUri = direct.ValueOf(in.TruststorePassword)
	return out
}

func ClusterMetastoreConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.MetastoreConfig) *krm.ClusterMetastoreConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterMetastoreConfig{}
	if in.GetDataprocMetastoreService() != "" {
		out.DataprocMetastoreServiceRef = apirefs.DataprocMetastoreServiceRef{
			External: in.GetDataprocMetastoreService(),
		}
	}
	return out
}

func ClusterMetastoreConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterMetastoreConfig) *pb.MetastoreConfig {
	if in == nil {
		return nil
	}
	out := &pb.MetastoreConfig{}
	out.DataprocMetastoreService = in.DataprocMetastoreServiceRef.External
	return out
}

func ClusterMetrics_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.DataprocMetricConfig_Metric) *krm.ClusterMetrics {
	if in == nil {
		return nil
	}
	out := &krm.ClusterMetrics{}
	out.MetricOverrides = in.MetricOverrides
	out.MetricSource = in.MetricSource.String()
	return out
}

func ClusterMetrics_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterMetrics) *pb.DataprocMetricConfig_Metric {
	if in == nil {
		return nil
	}
	out := &pb.DataprocMetricConfig_Metric{}
	out.MetricOverrides = in.MetricOverrides
	if val, ok := pb.DataprocMetricConfig_MetricSource_value[in.MetricSource]; ok {
		out.MetricSource = pb.DataprocMetricConfig_MetricSource(val)
	}
	return out
}

func ClusterMetricsStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ClusterMetrics) *krm.ClusterMetricsStatus {
	if in == nil {
		return nil
	}
	out := &krm.ClusterMetricsStatus{}
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

func ClusterMetricsStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterMetricsStatus) *pb.ClusterMetrics {
	if in == nil {
		return nil
	}
	out := &pb.ClusterMetrics{}
	if in.HdfsMetrics != nil {
		out.HdfsMetrics = make(map[string]int64)
		for k, v := range in.HdfsMetrics {
			if i, err := strconv.ParseInt(v, 10, 64); err == nil {
				out.HdfsMetrics[k] = i
			} else {
				mapCtx.Errorf("%v", err)
			}
		}
	}
	if in.YarnMetrics != nil {
		out.YarnMetrics = make(map[string]int64)
		for k, v := range in.YarnMetrics {
			if i, err := strconv.ParseInt(v, 10, 64); err == nil {
				out.YarnMetrics[k] = i
			} else {
				mapCtx.Errorf("%v", err)
			}
		}
	}
	return out
}

func ClusterNodePoolTarget_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GkeNodePoolTarget) *krm.ClusterNodePoolTarget {
	if in == nil {
		return nil
	}
	out := &krm.ClusterNodePoolTarget{}
	if in.GetNodePool() != "" {
		out.NodePoolRef = containerv1beta1.ContainerNodePoolRef{External: in.GetNodePool()}
	}
	out.Roles = direct.EnumSlice_FromProto(mapCtx, in.Roles)
	out.NodePoolConfig = ClusterNodePoolConfig_v1beta1_FromProto(mapCtx, in.GetNodePoolConfig())
	return out
}

func ClusterNodePoolTarget_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterNodePoolTarget) *pb.GkeNodePoolTarget {
	if in == nil {
		return nil
	}
	out := &pb.GkeNodePoolTarget{}
	out.NodePool = in.NodePoolRef.External
	out.Roles = direct.EnumSlice_ToProto[pb.GkeNodePoolTarget_Role](mapCtx, in.Roles)
	out.NodePoolConfig = ClusterNodePoolConfig_v1beta1_ToProto(mapCtx, in.NodePoolConfig)
	return out
}

func ClusterVirtualClusterConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.VirtualClusterConfig) *krm.ClusterVirtualClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterVirtualClusterConfig{}
	out.AuxiliaryServicesConfig = ClusterAuxiliaryServicesConfig_v1beta1_FromProto(mapCtx, in.GetAuxiliaryServicesConfig())
	if in.GetKubernetesClusterConfig() != nil {
		out.KubernetesClusterConfig = direct.ValueOf(ClusterKubernetesClusterConfig_v1beta1_FromProto(mapCtx, in.GetKubernetesClusterConfig()))
	}
	if in.GetStagingBucket() != "" {
		out.StagingBucketRef = &storagev1beta1.StorageBucketRef{External: in.GetStagingBucket()}
	}
	return out
}

func ClusterVirtualClusterConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterVirtualClusterConfig) *pb.VirtualClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.VirtualClusterConfig{}
	out.AuxiliaryServicesConfig = ClusterAuxiliaryServicesConfig_v1beta1_ToProto(mapCtx, in.AuxiliaryServicesConfig)
	hasKubeConfig := false
	if in.KubernetesClusterConfig.KubernetesNamespace != nil {
		hasKubeConfig = true
	}
	if in.KubernetesClusterConfig.KubernetesSoftwareConfig != nil {
		if len(in.KubernetesClusterConfig.KubernetesSoftwareConfig.ComponentVersion) > 0 || len(in.KubernetesClusterConfig.KubernetesSoftwareConfig.Properties) > 0 {
			hasKubeConfig = true
		}
	}
	if in.KubernetesClusterConfig.GkeClusterConfig.GkeClusterTargetRef != nil || len(in.KubernetesClusterConfig.GkeClusterConfig.NodePoolTarget) > 0 {
		hasKubeConfig = true
	}
	if hasKubeConfig {
		if protoConfig := ClusterKubernetesClusterConfig_v1beta1_ToProto(mapCtx, &in.KubernetesClusterConfig); protoConfig != nil {
			out.InfrastructureConfig = &pb.VirtualClusterConfig_KubernetesClusterConfig{
				KubernetesClusterConfig: protoConfig,
			}
		}
	}
	if in.StagingBucketRef != nil {
		out.StagingBucket = in.StagingBucketRef.External
	}
	return out
}

func ClusterKubernetesClusterConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesClusterConfig) *krm.ClusterKubernetesClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterKubernetesClusterConfig{}
	out.KubernetesNamespace = direct.LazyPtr(in.GetKubernetesNamespace())
	if in.GetGkeClusterConfig() != nil {
		out.GkeClusterConfig = direct.ValueOf(ClusterGkeClusterConfig_v1beta1_FromProto(mapCtx, in.GetGkeClusterConfig()))
	}
	out.KubernetesSoftwareConfig = ClusterKubernetesSoftwareConfig_v1beta1_FromProto(mapCtx, in.GetKubernetesSoftwareConfig())
	return out
}

func ClusterKubernetesClusterConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterKubernetesClusterConfig) *pb.KubernetesClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesClusterConfig{}
	out.KubernetesNamespace = direct.ValueOf(in.KubernetesNamespace)
	if in.GkeClusterConfig.GkeClusterTargetRef != nil || len(in.GkeClusterConfig.NodePoolTarget) > 0 {
		if protoGkeConfig := ClusterGkeClusterConfig_v1beta1_ToProto(mapCtx, &in.GkeClusterConfig); protoGkeConfig != nil {
			out.Config = &pb.KubernetesClusterConfig_GkeClusterConfig{
				GkeClusterConfig: protoGkeConfig,
			}
		}
	}
	out.KubernetesSoftwareConfig = ClusterKubernetesSoftwareConfig_v1beta1_ToProto(mapCtx, in.KubernetesSoftwareConfig)
	return out
}

func ClusterGkeClusterConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GkeClusterConfig) *krm.ClusterGkeClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterGkeClusterConfig{}
	if in.GetGkeClusterTarget() != "" {
		out.GkeClusterTargetRef = &containerv1beta1.ContainerClusterRef{External: in.GetGkeClusterTarget()}
	}
	out.NodePoolTarget = direct.Slice_FromProto(mapCtx, in.NodePoolTarget, ClusterNodePoolTarget_v1beta1_FromProto)
	return out
}

func ClusterGkeClusterConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterGkeClusterConfig) *pb.GkeClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.GkeClusterConfig{}
	if in.GkeClusterTargetRef != nil {
		out.GkeClusterTarget = in.GkeClusterTargetRef.External
	}
	out.NodePoolTarget = direct.Slice_ToProto(mapCtx, in.NodePoolTarget, ClusterNodePoolTarget_v1beta1_ToProto)
	return out
}

func ClusterAutoscalingConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingConfig) *krm.ClusterAutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterAutoscalingConfig{}
	if in.GetPolicyUri() != "" {
		out.PolicyRef = &krm.DataprocAutoscalingPolicyRef{External: in.GetPolicyUri()}
	}
	return out
}

func ClusterAutoscalingConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterAutoscalingConfig) *pb.AutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingConfig{}
	if in.PolicyRef != nil {
		out.PolicyUri = in.PolicyRef.External
	}
	return out
}

func ClusterGceClusterConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GceClusterConfig) *krm.ClusterGceClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterGceClusterConfig{}
	if in.GetZoneUri() != "" {
		out.Zone = direct.LazyPtr(in.GetZoneUri())
	}
	if in.GetNetworkUri() != "" {
		out.NetworkRef = &computerefs.ComputeNetworkRef{External: in.GetNetworkUri()}
	}
	if in.GetSubnetworkUri() != "" {
		out.SubnetworkRef = &computev1beta1.ComputeSubnetworkRef{External: in.GetSubnetworkUri()}
	}
	out.InternalIPOnly = in.InternalIpOnly
	if in.PrivateIpv6GoogleAccess != pb.GceClusterConfig_PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED {
		out.PrivateIPv6GoogleAccess = direct.LazyPtr(in.PrivateIpv6GoogleAccess.String())
	}
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.ServiceAccountScopes = in.ServiceAccountScopes
	out.Tags = in.Tags
	out.Metadata = in.Metadata
	out.ReservationAffinity = ClusterReservationAffinity_v1beta1_FromProto(mapCtx, in.GetReservationAffinity())
	out.NodeGroupAffinity = ClusterNodeGroupAffinity_v1beta1_FromProto(mapCtx, in.GetNodeGroupAffinity())
	out.ShieldedInstanceConfig = ClusterShieldedInstanceConfig_v1beta1_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.ConfidentialInstanceConfig = ClusterConfidentialInstanceConfig_v1beta1_FromProto(mapCtx, in.GetConfidentialInstanceConfig())
	return out
}

func ClusterGceClusterConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterGceClusterConfig) *pb.GceClusterConfig {
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
	out.InternalIpOnly = in.InternalIPOnly
	if in.PrivateIPv6GoogleAccess != nil {
		if val, ok := pb.GceClusterConfig_PrivateIpv6GoogleAccess_value[*in.PrivateIPv6GoogleAccess]; ok {
			out.PrivateIpv6GoogleAccess = pb.GceClusterConfig_PrivateIpv6GoogleAccess(val)
		}
	}
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.ServiceAccountScopes = in.ServiceAccountScopes
	out.Tags = in.Tags
	out.Metadata = in.Metadata
	out.ReservationAffinity = ClusterReservationAffinity_v1beta1_ToProto(mapCtx, in.ReservationAffinity)
	out.NodeGroupAffinity = ClusterNodeGroupAffinity_v1beta1_ToProto(mapCtx, in.NodeGroupAffinity)
	out.ShieldedInstanceConfig = ClusterShieldedInstanceConfig_v1beta1_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.ConfidentialInstanceConfig = ClusterConfidentialInstanceConfig_v1beta1_ToProto(mapCtx, in.ConfidentialInstanceConfig)
	return out
}

func ClusterConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ClusterConfig) *krm.ClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterConfig{}
	if in.GetConfigBucket() != "" {
		out.StagingBucketRef = &storagev1beta1.StorageBucketRef{External: in.GetConfigBucket()}
	}
	if in.GetTempBucket() != "" {
		out.TempBucketRef = &storagev1beta1.StorageBucketRef{External: in.GetTempBucket()}
	}
	out.GceClusterConfig = ClusterGceClusterConfig_v1beta1_FromProto(mapCtx, in.GetGceClusterConfig())
	out.MasterConfig = ClusterMasterConfig_v1beta1_FromProto(mapCtx, in.GetMasterConfig())
	out.WorkerConfig = ClusterWorkerConfig_v1beta1_FromProto(mapCtx, in.GetWorkerConfig())
	out.SecondaryWorkerConfig = ClusterSecondaryWorkerConfig_v1beta1_FromProto(mapCtx, in.GetSecondaryWorkerConfig())
	out.SoftwareConfig = ClusterSoftwareConfig_v1beta1_FromProto(mapCtx, in.GetSoftwareConfig())
	out.InitializationActions = direct.Slice_FromProto(mapCtx, in.InitializationActions, ClusterInitializationActions_v1beta1_FromProto)
	out.EncryptionConfig = ClusterEncryptionConfig_v1beta1_FromProto(mapCtx, in.GetEncryptionConfig())
	out.AutoscalingConfig = ClusterAutoscalingConfig_v1beta1_FromProto(mapCtx, in.GetAutoscalingConfig())
	out.SecurityConfig = ClusterSecurityConfig_v1beta1_FromProto(mapCtx, in.GetSecurityConfig())
	out.LifecycleConfig = ClusterLifecycleConfig_v1beta1_FromProto(mapCtx, in.GetLifecycleConfig())
	out.EndpointConfig = ClusterEndpointConfig_v1beta1_FromProto(mapCtx, in.GetEndpointConfig())
	out.MetastoreConfig = ClusterMetastoreConfig_v1beta1_FromProto(mapCtx, in.GetMetastoreConfig())
	out.DataprocMetricConfig = ClusterDataprocMetricConfig_v1beta1_FromProto(mapCtx, in.GetDataprocMetricConfig())
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
	out.GceClusterConfig = ClusterGceClusterConfig_v1beta1_ToProto(mapCtx, in.GceClusterConfig)
	out.MasterConfig = ClusterMasterConfig_v1beta1_ToProto(mapCtx, in.MasterConfig)
	out.WorkerConfig = ClusterWorkerConfig_v1beta1_ToProto(mapCtx, in.WorkerConfig)
	out.SecondaryWorkerConfig = ClusterSecondaryWorkerConfig_v1beta1_ToProto(mapCtx, in.SecondaryWorkerConfig)
	out.SoftwareConfig = ClusterSoftwareConfig_v1beta1_ToProto(mapCtx, in.SoftwareConfig)
	out.InitializationActions = direct.Slice_ToProto(mapCtx, in.InitializationActions, ClusterInitializationActions_v1beta1_ToProto)
	out.EncryptionConfig = ClusterEncryptionConfig_v1beta1_ToProto(mapCtx, in.EncryptionConfig)
	out.AutoscalingConfig = ClusterAutoscalingConfig_v1beta1_ToProto(mapCtx, in.AutoscalingConfig)
	out.SecurityConfig = ClusterSecurityConfig_v1beta1_ToProto(mapCtx, in.SecurityConfig)
	out.LifecycleConfig = ClusterLifecycleConfig_v1beta1_ToProto(mapCtx, in.LifecycleConfig)
	out.EndpointConfig = ClusterEndpointConfig_v1beta1_ToProto(mapCtx, in.EndpointConfig)
	out.MetastoreConfig = ClusterMetastoreConfig_v1beta1_ToProto(mapCtx, in.MetastoreConfig)
	out.DataprocMetricConfig = ClusterDataprocMetricConfig_v1beta1_ToProto(mapCtx, in.DataprocMetricConfig)
	return out
}

func ClusterSoftwareConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ClusterSoftwareConfig) *pb.SoftwareConfig {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareConfig{}
	out.ImageVersion = direct.ValueOf(in.ImageVersion)
	out.Properties = in.Properties
	out.OptionalComponents = direct.EnumSlice_ToProto[pb.Component](mapCtx, in.OptionalComponents)
	return out
}

func ClusterSoftwareConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareConfig) *krm.ClusterSoftwareConfig {
	if in == nil {
		return nil
	}
	out := &krm.ClusterSoftwareConfig{}
	out.ImageVersion = direct.LazyPtr(in.GetImageVersion())
	out.Properties = in.Properties
	out.OptionalComponents = direct.EnumSlice_FromProto(mapCtx, in.OptionalComponents)
	return out
}
