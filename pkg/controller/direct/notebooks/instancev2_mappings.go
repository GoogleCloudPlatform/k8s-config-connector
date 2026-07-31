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

package notebooks

import (
	pb "cloud.google.com/go/notebooks/apiv2/notebookspb"
	kms "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/notebooks/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// We define manual mappers here for v2 types where the auto-generator got near-misses.

func InstanceServiceAccount_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccount) *krm.InstanceServiceAccount {
	if in == nil {
		return nil
	}
	out := &krm.InstanceServiceAccount{}
	if in.Email != "" {
		out.ServiceAccountRef = &refs.IAMServiceAccountRef{External: in.Email}
	}
	return out
}

func InstanceServiceAccount_ToProto(mapCtx *direct.MapContext, in *krm.InstanceServiceAccount) *pb.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccount{}
	if in.ServiceAccountRef != nil {
		out.Email = in.ServiceAccountRef.External
	}
	return out
}

func InstanceBootDisk_FromProto(mapCtx *direct.MapContext, in *pb.BootDisk) *krm.InstanceBootDisk {
	if in == nil {
		return nil
	}
	out := &krm.InstanceBootDisk{}
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	out.DiskType = direct.Enum_FromProto(mapCtx, in.GetDiskType())
	out.DiskEncryption = direct.Enum_FromProto(mapCtx, in.GetDiskEncryption())
	if in.KmsKey != "" {
		out.KmsKeyRef = &kms.KMSCryptoKeyRef{External: in.KmsKey}
	}
	return out
}

func InstanceBootDisk_ToProto(mapCtx *direct.MapContext, in *krm.InstanceBootDisk) *pb.BootDisk {
	if in == nil {
		return nil
	}
	out := &pb.BootDisk{}
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	out.DiskType = direct.Enum_ToProto[pb.DiskType](mapCtx, in.DiskType)
	out.DiskEncryption = direct.Enum_ToProto[pb.DiskEncryption](mapCtx, in.DiskEncryption)
	if in.KmsKeyRef != nil {
		out.KmsKey = in.KmsKeyRef.External
	}
	return out
}

func InstanceDataDisk_FromProto(mapCtx *direct.MapContext, in *pb.DataDisk) *krm.InstanceDataDisk {
	if in == nil {
		return nil
	}
	out := &krm.InstanceDataDisk{}
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	out.DiskType = direct.Enum_FromProto(mapCtx, in.GetDiskType())
	out.DiskEncryption = direct.Enum_FromProto(mapCtx, in.GetDiskEncryption())
	if in.KmsKey != "" {
		out.KmsKeyRef = &kms.KMSCryptoKeyRef{External: in.KmsKey}
	}
	return out
}

func InstanceDataDisk_ToProto(mapCtx *direct.MapContext, in *krm.InstanceDataDisk) *pb.DataDisk {
	if in == nil {
		return nil
	}
	out := &pb.DataDisk{}
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	out.DiskType = direct.Enum_ToProto[pb.DiskType](mapCtx, in.DiskType)
	out.DiskEncryption = direct.Enum_ToProto[pb.DiskEncryption](mapCtx, in.DiskEncryption)
	if in.KmsKeyRef != nil {
		out.KmsKey = in.KmsKeyRef.External
	}
	return out
}

func InstanceGCESetup_FromProto(mapCtx *direct.MapContext, in *pb.GceSetup) *krm.InstanceGCESetup {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGCESetup{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.AcceleratorConfigs = direct.Slice_FromProto(mapCtx, in.AcceleratorConfigs, InstanceAcceleratorConfig_FromProto)
	out.ServiceAccounts = direct.Slice_FromProto(mapCtx, in.ServiceAccounts, InstanceServiceAccount_FromProto)
	out.VMImage = InstanceVMImage_FromProto(mapCtx, in.GetVmImage())
	out.ContainerImage = InstanceContainerImage_FromProto(mapCtx, in.GetContainerImage())
	out.BootDisk = InstanceBootDisk_FromProto(mapCtx, in.GetBootDisk())
	out.DataDisks = direct.Slice_FromProto(mapCtx, in.DataDisks, InstanceDataDisk_FromProto)
	out.ShieldedInstanceConfig = InstanceShieldedInstanceConfig_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.NetworkInterfaces = direct.Slice_FromProto(mapCtx, in.NetworkInterfaces, InstanceNetworkInterface_FromProto)
	out.DisablePublicIP = direct.LazyPtr(in.GetDisablePublicIp())
	out.Tags = in.Tags
	out.Metadata = in.Metadata
	out.EnableIPForwarding = direct.LazyPtr(in.GetEnableIpForwarding())
	out.GPUDriverConfig = InstanceGPUDriverConfig_FromProto(mapCtx, in.GetGpuDriverConfig())
	return out
}

func InstanceGCESetup_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGCESetup) *pb.GceSetup {
	if in == nil {
		return nil
	}
	out := &pb.GceSetup{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.AcceleratorConfigs = direct.Slice_ToProto(mapCtx, in.AcceleratorConfigs, InstanceAcceleratorConfig_ToProto)
	out.ServiceAccounts = direct.Slice_ToProto(mapCtx, in.ServiceAccounts, InstanceServiceAccount_ToProto)
	if in.VMImage != nil {
		out.Image = &pb.GceSetup_VmImage{VmImage: InstanceVMImage_ToProto(mapCtx, in.VMImage)}
	}
	if in.ContainerImage != nil {
		out.Image = &pb.GceSetup_ContainerImage{ContainerImage: InstanceContainerImage_ToProto(mapCtx, in.ContainerImage)}
	}
	out.BootDisk = InstanceBootDisk_ToProto(mapCtx, in.BootDisk)
	out.DataDisks = direct.Slice_ToProto(mapCtx, in.DataDisks, InstanceDataDisk_ToProto)
	out.ShieldedInstanceConfig = InstanceShieldedInstanceConfig_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.NetworkInterfaces = direct.Slice_ToProto(mapCtx, in.NetworkInterfaces, InstanceNetworkInterface_ToProto)
	out.DisablePublicIp = direct.ValueOf(in.DisablePublicIP)
	out.Tags = in.Tags
	out.Metadata = in.Metadata
	out.EnableIpForwarding = direct.ValueOf(in.EnableIPForwarding)
	out.GpuDriverConfig = InstanceGPUDriverConfig_ToProto(mapCtx, in.GPUDriverConfig)
	return out
}
