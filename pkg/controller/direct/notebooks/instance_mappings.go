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

package notebooks

import (
	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	notebookspb "cloud.google.com/go/notebooks/apiv2/notebookspb"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krmnotebooksv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/notebooks/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VMImage_ImageName_ToProto(mapCtx *direct.MapContext, in *string) *pb.VmImage_ImageName {
	if in == nil {
		return nil
	}
	return &pb.VmImage_ImageName{ImageName: direct.ValueOf(in)}
}

func VMImage_ImageFamily_ToProto(mapCtx *direct.MapContext, in *string) *pb.VmImage_ImageFamily {
	if in == nil {
		return nil
	}
	return &pb.VmImage_ImageFamily{ImageFamily: direct.ValueOf(in)}
}

func InstanceBootDisk_v1alpha1_FromProto(mapCtx *direct.MapContext, in *notebookspb.BootDisk) *krmnotebooksv1alpha1.InstanceBootDisk {
	if in == nil {
		return nil
	}
	out := &krmnotebooksv1alpha1.InstanceBootDisk{}
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	out.DiskType = direct.Enum_FromProto(mapCtx, in.GetDiskType())
	out.DiskEncryption = direct.Enum_FromProto(mapCtx, in.GetDiskEncryption())
	if in.GetDiskEncryption() == notebookspb.DiskEncryption_CMEK {
		if in.GetKmsKey() != "" {
			out.KmsKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKey()}
		}
	}
	return out
}

func InstanceBootDisk_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnotebooksv1alpha1.InstanceBootDisk) *notebookspb.BootDisk {
	if in == nil {
		return nil
	}
	out := &notebookspb.BootDisk{}
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	out.DiskType = direct.Enum_ToProto[notebookspb.DiskType](mapCtx, in.DiskType)
	out.DiskEncryption = direct.Enum_ToProto[notebookspb.DiskEncryption](mapCtx, in.DiskEncryption)
	if in.KmsKeyRef != nil {
		out.KmsKey = in.KmsKeyRef.External
		out.DiskEncryption = notebookspb.DiskEncryption_CMEK
	}
	return out
}

func InstanceDataDisk_v1alpha1_FromProto(mapCtx *direct.MapContext, in *notebookspb.DataDisk) *krmnotebooksv1alpha1.InstanceDataDisk {
	if in == nil {
		return nil
	}
	out := &krmnotebooksv1alpha1.InstanceDataDisk{}
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	out.DiskType = direct.Enum_FromProto(mapCtx, in.GetDiskType())
	out.DiskEncryption = direct.Enum_FromProto(mapCtx, in.GetDiskEncryption())
	if in.GetDiskEncryption() == notebookspb.DiskEncryption_CMEK {
		if in.GetKmsKey() != "" {
			out.KmsKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKey()}
		}
	}
	return out
}

func InstanceDataDisk_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnotebooksv1alpha1.InstanceDataDisk) *notebookspb.DataDisk {
	if in == nil {
		return nil
	}
	out := &notebookspb.DataDisk{}
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	out.DiskType = direct.Enum_ToProto[notebookspb.DiskType](mapCtx, in.DiskType)
	out.DiskEncryption = direct.Enum_ToProto[notebookspb.DiskEncryption](mapCtx, in.DiskEncryption)
	if in.KmsKeyRef != nil {
		out.KmsKey = in.KmsKeyRef.External
		out.DiskEncryption = notebookspb.DiskEncryption_CMEK
	}
	return out
}

func InstanceGCESetup_v1alpha1_FromProto(mapCtx *direct.MapContext, in *notebookspb.GceSetup) *krmnotebooksv1alpha1.InstanceGCESetup {
	if in == nil {
		return nil
	}
	out := &krmnotebooksv1alpha1.InstanceGCESetup{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.AcceleratorConfigs = direct.Slice_FromProto(mapCtx, in.AcceleratorConfigs, InstanceAcceleratorConfig_v1alpha1_FromProto)
	out.ServiceAccounts = direct.Slice_FromProto(mapCtx, in.ServiceAccounts, InstanceServiceAccount_v1alpha1_FromProto)
	out.VMImage = InstanceVMImage_v1alpha1_FromProto(mapCtx, in.GetVmImage())
	out.ContainerImage = InstanceContainerImage_v1alpha1_FromProto(mapCtx, in.GetContainerImage())
	out.BootDisk = InstanceBootDisk_v1alpha1_FromProto(mapCtx, in.GetBootDisk())
	out.DataDisks = direct.Slice_FromProto(mapCtx, in.DataDisks, InstanceDataDisk_v1alpha1_FromProto)
	out.ShieldedInstanceConfig = InstanceShieldedInstanceConfig_v1alpha1_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.NetworkInterfaces = direct.Slice_FromProto(mapCtx, in.NetworkInterfaces, InstanceNetworkInterface_v1alpha1_FromProto)
	out.DisablePublicIP = direct.LazyPtr(in.GetDisablePublicIp())
	out.Tags = in.Tags
	out.Metadata = in.Metadata
	out.EnableIPForwarding = direct.LazyPtr(in.GetEnableIpForwarding())
	out.GPUDriverConfig = InstanceGPUDriverConfig_v1alpha1_FromProto(mapCtx, in.GetGpuDriverConfig())
	return out
}

func InstanceGCESetup_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnotebooksv1alpha1.InstanceGCESetup) *notebookspb.GceSetup {
	if in == nil {
		return nil
	}
	out := &notebookspb.GceSetup{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.AcceleratorConfigs = direct.Slice_ToProto(mapCtx, in.AcceleratorConfigs, InstanceAcceleratorConfig_v1alpha1_ToProto)
	out.ServiceAccounts = direct.Slice_ToProto(mapCtx, in.ServiceAccounts, InstanceServiceAccount_v1alpha1_ToProto)
	if oneof := InstanceVMImage_v1alpha1_ToProto(mapCtx, in.VMImage); oneof != nil {
		out.Image = &notebookspb.GceSetup_VmImage{VmImage: oneof}
	}
	if oneof := InstanceContainerImage_v1alpha1_ToProto(mapCtx, in.ContainerImage); oneof != nil {
		out.Image = &notebookspb.GceSetup_ContainerImage{ContainerImage: oneof}
	}
	out.BootDisk = InstanceBootDisk_v1alpha1_ToProto(mapCtx, in.BootDisk)
	out.DataDisks = direct.Slice_ToProto(mapCtx, in.DataDisks, InstanceDataDisk_v1alpha1_ToProto)
	out.ShieldedInstanceConfig = InstanceShieldedInstanceConfig_v1alpha1_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.NetworkInterfaces = direct.Slice_ToProto(mapCtx, in.NetworkInterfaces, InstanceNetworkInterface_v1alpha1_ToProto)
	out.DisablePublicIp = direct.ValueOf(in.DisablePublicIP)
	out.Tags = in.Tags
	out.Metadata = in.Metadata
	out.EnableIpForwarding = direct.ValueOf(in.EnableIPForwarding)
	out.GpuDriverConfig = InstanceGPUDriverConfig_v1alpha1_ToProto(mapCtx, in.GPUDriverConfig)
	return out
}

func InstanceServiceAccount_v1alpha1_FromProto(mapCtx *direct.MapContext, in *notebookspb.ServiceAccount) *krmnotebooksv1alpha1.InstanceServiceAccount {
	if in == nil {
		return nil
	}
	out := &krmnotebooksv1alpha1.InstanceServiceAccount{}
	if in.GetEmail() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetEmail()}
	}
	return out
}

func InstanceServiceAccount_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnotebooksv1alpha1.InstanceServiceAccount) *notebookspb.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &notebookspb.ServiceAccount{}
	if in.ServiceAccountRef != nil {
		out.Email = in.ServiceAccountRef.External
	}
	return out
}

func NotebookInstanceV2ObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *notebookspb.Instance) *krmnotebooksv1alpha1.NotebookInstanceV2ObservedState {
	if in == nil {
		return nil
	}
	out := &krmnotebooksv1alpha1.NotebookInstanceV2ObservedState{}
	// MISSING: Name
	out.GCESetup = InstanceGCESetupObservedState_v1alpha1_FromProto(mapCtx, in.GetGceSetup())
	out.ProxyURI = direct.LazyPtr(in.GetProxyUri())
	out.Creator = direct.LazyPtr(in.GetCreator())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UpgradeHistory = direct.Slice_FromProto(mapCtx, in.UpgradeHistory, InstanceUpgradeHistoryEntryObservedState_v1alpha1_FromProto)
	out.GCPID = direct.LazyPtr(in.Id)
	out.HealthState = direct.Enum_FromProto(mapCtx, in.GetHealthState())
	out.HealthInfo = in.HealthInfo
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func NotebookInstanceV2ObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnotebooksv1alpha1.NotebookInstanceV2ObservedState) *notebookspb.Instance {
	if in == nil {
		return nil
	}
	out := &notebookspb.Instance{}
	// MISSING: Name
	if oneof := InstanceGCESetupObservedState_v1alpha1_ToProto(mapCtx, in.GCESetup); oneof != nil {
		out.Infrastructure = &notebookspb.Instance_GceSetup{GceSetup: oneof}
	}
	out.ProxyUri = direct.ValueOf(in.ProxyURI)
	out.Creator = direct.ValueOf(in.Creator)
	out.State = direct.Enum_ToProto[notebookspb.State](mapCtx, in.State)
	out.UpgradeHistory = direct.Slice_ToProto(mapCtx, in.UpgradeHistory, InstanceUpgradeHistoryEntryObservedState_v1alpha1_ToProto)
	out.Id = direct.ValueOf(in.GCPID)
	out.HealthState = direct.Enum_ToProto[notebookspb.HealthState](mapCtx, in.HealthState)
	out.HealthInfo = in.HealthInfo
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
