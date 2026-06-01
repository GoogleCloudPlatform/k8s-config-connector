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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DiskConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.DiskConfig) *krm.DiskConfig {
	if in == nil {
		return nil
	}
	out := &krm.DiskConfig{}
	out.BootDiskType = direct.LazyPtr(in.GetBootDiskType())
	out.BootDiskSizeGB = direct.LazyPtr(in.GetBootDiskSizeGb())
	out.NumLocalSSDs = direct.LazyPtr(in.GetNumLocalSsds())
	out.LocalSSDInterface = direct.LazyPtr(in.GetLocalSsdInterface())
	out.BootDiskProvisionedIOPs = in.BootDiskProvisionedIops
	out.BootDiskProvisionedThroughput = in.BootDiskProvisionedThroughput
	return out
}
func DiskConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.DiskConfig) *pb.DiskConfig {
	if in == nil {
		return nil
	}
	out := &pb.DiskConfig{}
	out.BootDiskType = direct.ValueOf(in.BootDiskType)
	out.BootDiskSizeGb = direct.ValueOf(in.BootDiskSizeGB)
	out.NumLocalSsds = direct.ValueOf(in.NumLocalSSDs)
	out.LocalSsdInterface = direct.ValueOf(in.LocalSSDInterface)
	out.BootDiskProvisionedIops = in.BootDiskProvisionedIOPs
	out.BootDiskProvisionedThroughput = in.BootDiskProvisionedThroughput
	return out
}

func AcceleratorConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorConfig) *krm.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krm.AcceleratorConfig{}
	if in.GetAcceleratorTypeUri() != "" {
		out.AcceleratorTypeRef = &computev1beta1.ComputeAcceleratorTypeRef{
			External: in.GetAcceleratorTypeUri(),
		}
	}
	out.AcceleratorCount = direct.LazyPtr(in.GetAcceleratorCount())
	return out
}
func AcceleratorConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AcceleratorConfig) *pb.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorConfig{}
	if in.AcceleratorTypeRef != nil {
		out.AcceleratorTypeUri = in.AcceleratorTypeRef.External
	}
	out.AcceleratorCount = direct.ValueOf(in.AcceleratorCount)
	return out
}

func InstanceGroupConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupConfig) *krm.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupConfig{}
	out.NumInstances = direct.LazyPtr(in.GetNumInstances())
	if in.GetImageUri() != "" {
		out.ImageRef = &computev1beta1.ComputeImageRef{
			External: in.GetImageUri(),
		}
	}
	if in.GetMachineTypeUri() != "" {
		out.MachineTypeRef = &computev1beta1.ComputeMachineTypeRef{
			External: in.GetMachineTypeUri(),
		}
	}
	out.DiskConfig = DiskConfig_v1alpha1_FromProto(mapCtx, in.GetDiskConfig())
	out.Preemptibility = direct.Enum_FromProto(mapCtx, in.GetPreemptibility())
	out.Accelerators = direct.Slice_FromProto(mapCtx, in.Accelerators, AcceleratorConfig_v1alpha1_FromProto)
	out.MinCPUPlatform = direct.LazyPtr(in.GetMinCpuPlatform())
	out.MinNumInstances = direct.LazyPtr(in.GetMinNumInstances())
	out.InstanceFlexibilityPolicy = InstanceFlexibilityPolicy_v1alpha1_FromProto(mapCtx, in.GetInstanceFlexibilityPolicy())
	out.StartupConfig = StartupConfig_v1alpha1_FromProto(mapCtx, in.GetStartupConfig())
	return out
}
func InstanceGroupConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupConfig) *pb.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupConfig{}
	out.NumInstances = direct.ValueOf(in.NumInstances)
	if in.ImageRef != nil {
		out.ImageUri = in.ImageRef.External
	}
	if in.MachineTypeRef != nil {
		out.MachineTypeUri = in.MachineTypeRef.External
	}
	out.DiskConfig = DiskConfig_v1alpha1_ToProto(mapCtx, in.DiskConfig)
	out.Preemptibility = direct.Enum_ToProto[pb.InstanceGroupConfig_Preemptibility](mapCtx, in.Preemptibility)
	out.Accelerators = direct.Slice_ToProto(mapCtx, in.Accelerators, AcceleratorConfig_v1alpha1_ToProto)
	out.MinCpuPlatform = direct.ValueOf(in.MinCPUPlatform)
	out.MinNumInstances = direct.ValueOf(in.MinNumInstances)
	out.InstanceFlexibilityPolicy = InstanceFlexibilityPolicy_v1alpha1_ToProto(mapCtx, in.InstanceFlexibilityPolicy)
	out.StartupConfig = StartupConfig_v1alpha1_ToProto(mapCtx, in.StartupConfig)
	return out
}

func InstanceGroupConfigObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupConfig) *krm.InstanceGroupConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupConfigObservedState{}
	out.InstanceNames = in.GetInstanceNames()
	out.InstanceReferences = direct.Slice_FromProto(mapCtx, in.InstanceReferences, InstanceReference_v1alpha1_FromProto)
	out.IsPreemptible = direct.LazyPtr(in.GetIsPreemptible())
	out.ManagedGroupConfig = ManagedGroupConfig_v1alpha1_FromProto(mapCtx, in.GetManagedGroupConfig())
	out.InstanceFlexibilityPolicy = InstanceFlexibilityPolicyObservedState_v1alpha1_FromProto(mapCtx, in.GetInstanceFlexibilityPolicy())
	return out
}
func InstanceGroupConfigObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupConfigObservedState) *pb.InstanceGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupConfig{}
	out.InstanceNames = in.InstanceNames
	out.InstanceReferences = direct.Slice_ToProto(mapCtx, in.InstanceReferences, InstanceReference_v1alpha1_ToProto)
	out.IsPreemptible = direct.ValueOf(in.IsPreemptible)
	out.ManagedGroupConfig = ManagedGroupConfig_v1alpha1_ToProto(mapCtx, in.ManagedGroupConfig)
	return out
}

func InstanceReference_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceReference) *krm.InstanceReference {
	if in == nil {
		return nil
	}
	out := &krm.InstanceReference{}
	out.InstanceName = direct.LazyPtr(in.GetInstanceName())
	out.InstanceID = direct.LazyPtr(in.GetInstanceId())
	out.PublicKey = direct.LazyPtr(in.GetPublicKey())
	return out
}
func InstanceReference_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceReference) *pb.InstanceReference {
	if in == nil {
		return nil
	}
	out := &pb.InstanceReference{}
	out.InstanceName = direct.ValueOf(in.InstanceName)
	out.InstanceId = direct.ValueOf(in.InstanceID)
	out.PublicKey = direct.ValueOf(in.PublicKey)
	return out
}

func ManagedGroupConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ManagedGroupConfig) *krm.ManagedGroupConfig {
	if in == nil {
		return nil
	}
	out := &krm.ManagedGroupConfig{}
	out.InstanceTemplateName = direct.LazyPtr(in.GetInstanceTemplateName())
	out.InstanceGroupManagerName = direct.LazyPtr(in.GetInstanceGroupManagerName())
	return out
}
func ManagedGroupConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ManagedGroupConfig) *pb.ManagedGroupConfig {
	if in == nil {
		return nil
	}
	out := &pb.ManagedGroupConfig{}
	out.InstanceTemplateName = direct.ValueOf(in.InstanceTemplateName)
	out.InstanceGroupManagerName = direct.ValueOf(in.InstanceGroupManagerName)
	return out
}
