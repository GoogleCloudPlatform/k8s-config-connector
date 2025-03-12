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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
func DataprocNodeGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NodeGroup) *krm.DataprocNodeGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataprocNodeGroupObservedState{}
	// MISSING: Name
	out.NodeGroupConfig = InstanceGroupConfigObservedState_FromProto(mapCtx, in.GetNodeGroupConfig())
	return out
}
func DataprocNodeGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataprocNodeGroupObservedState) *pb.NodeGroup {
	if in == nil {
		return nil
	}
	out := &pb.NodeGroup{}
	// MISSING: Name
	out.NodeGroupConfig = InstanceGroupConfigObservedState_ToProto(mapCtx, in.NodeGroupConfig)
	return out
}
func DataprocNodeGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.NodeGroup) *krm.DataprocNodeGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataprocNodeGroupSpec{}
	// MISSING: Name
	out.Roles = direct.EnumSlice_FromProto(mapCtx, in.Roles)
	out.NodeGroupConfig = InstanceGroupConfig_FromProto(mapCtx, in.GetNodeGroupConfig())
	out.Labels = in.Labels
	return out
}
func DataprocNodeGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataprocNodeGroupSpec) *pb.NodeGroup {
	if in == nil {
		return nil
	}
	out := &pb.NodeGroup{}
	// MISSING: Name
	out.Roles = direct.EnumSlice_ToProto[pb.NodeGroup_Role](mapCtx, in.Roles)
	out.NodeGroupConfig = InstanceGroupConfig_ToProto(mapCtx, in.NodeGroupConfig)
	out.Labels = in.Labels
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
