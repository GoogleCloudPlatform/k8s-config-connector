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

package migrationcenter

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/migrationcenter/v1alpha1"
)
func AssetFrame_FromProto(mapCtx *direct.MapContext, in *pb.AssetFrame) *krm.AssetFrame {
	if in == nil {
		return nil
	}
	out := &krm.AssetFrame{}
	out.MachineDetails = MachineDetails_FromProto(mapCtx, in.GetMachineDetails())
	out.ReportTime = direct.StringTimestamp_FromProto(mapCtx, in.GetReportTime())
	out.Labels = in.Labels
	out.Attributes = in.Attributes
	out.PerformanceSamples = direct.Slice_FromProto(mapCtx, in.PerformanceSamples, PerformanceSample_FromProto)
	out.TraceToken = direct.LazyPtr(in.GetTraceToken())
	return out
}
func AssetFrame_ToProto(mapCtx *direct.MapContext, in *krm.AssetFrame) *pb.AssetFrame {
	if in == nil {
		return nil
	}
	out := &pb.AssetFrame{}
	if oneof := MachineDetails_ToProto(mapCtx, in.MachineDetails); oneof != nil {
		out.FrameData = &pb.AssetFrame_MachineDetails{MachineDetails: oneof}
	}
	out.ReportTime = direct.StringTimestamp_ToProto(mapCtx, in.ReportTime)
	out.Labels = in.Labels
	out.Attributes = in.Attributes
	out.PerformanceSamples = direct.Slice_ToProto(mapCtx, in.PerformanceSamples, PerformanceSample_ToProto)
	out.TraceToken = direct.ValueOf(in.TraceToken)
	return out
}
func AwsEc2PlatformDetails_FromProto(mapCtx *direct.MapContext, in *pb.AwsEc2PlatformDetails) *krm.AwsEc2PlatformDetails {
	if in == nil {
		return nil
	}
	out := &krm.AwsEc2PlatformDetails{}
	out.MachineTypeLabel = direct.LazyPtr(in.GetMachineTypeLabel())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func AwsEc2PlatformDetails_ToProto(mapCtx *direct.MapContext, in *krm.AwsEc2PlatformDetails) *pb.AwsEc2PlatformDetails {
	if in == nil {
		return nil
	}
	out := &pb.AwsEc2PlatformDetails{}
	out.MachineTypeLabel = direct.ValueOf(in.MachineTypeLabel)
	out.Location = direct.ValueOf(in.Location)
	return out
}
func AzureVmPlatformDetails_FromProto(mapCtx *direct.MapContext, in *pb.AzureVmPlatformDetails) *krm.AzureVmPlatformDetails {
	if in == nil {
		return nil
	}
	out := &krm.AzureVmPlatformDetails{}
	out.MachineTypeLabel = direct.LazyPtr(in.GetMachineTypeLabel())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.ProvisioningState = direct.LazyPtr(in.GetProvisioningState())
	return out
}
func AzureVmPlatformDetails_ToProto(mapCtx *direct.MapContext, in *krm.AzureVmPlatformDetails) *pb.AzureVmPlatformDetails {
	if in == nil {
		return nil
	}
	out := &pb.AzureVmPlatformDetails{}
	out.MachineTypeLabel = direct.ValueOf(in.MachineTypeLabel)
	out.Location = direct.ValueOf(in.Location)
	out.ProvisioningState = direct.ValueOf(in.ProvisioningState)
	return out
}
func BiosDetails_FromProto(mapCtx *direct.MapContext, in *pb.BiosDetails) *krm.BiosDetails {
	if in == nil {
		return nil
	}
	out := &krm.BiosDetails{}
	out.BiosName = direct.LazyPtr(in.GetBiosName())
	out.ID = direct.LazyPtr(in.GetId())
	out.Manufacturer = direct.LazyPtr(in.GetManufacturer())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.ReleaseDate = Date_FromProto(mapCtx, in.GetReleaseDate())
	out.SmbiosUuid = direct.LazyPtr(in.GetSmbiosUuid())
	return out
}
func BiosDetails_ToProto(mapCtx *direct.MapContext, in *krm.BiosDetails) *pb.BiosDetails {
	if in == nil {
		return nil
	}
	out := &pb.BiosDetails{}
	out.BiosName = direct.ValueOf(in.BiosName)
	out.Id = direct.ValueOf(in.ID)
	out.Manufacturer = direct.ValueOf(in.Manufacturer)
	out.Version = direct.ValueOf(in.Version)
	out.ReleaseDate = Date_ToProto(mapCtx, in.ReleaseDate)
	out.SmbiosUuid = direct.ValueOf(in.SmbiosUuid)
	return out
}
func CpuUsageSample_FromProto(mapCtx *direct.MapContext, in *pb.CpuUsageSample) *krm.CpuUsageSample {
	if in == nil {
		return nil
	}
	out := &krm.CpuUsageSample{}
	out.UtilizedPercentage = direct.LazyPtr(in.GetUtilizedPercentage())
	return out
}
func CpuUsageSample_ToProto(mapCtx *direct.MapContext, in *krm.CpuUsageSample) *pb.CpuUsageSample {
	if in == nil {
		return nil
	}
	out := &pb.CpuUsageSample{}
	out.UtilizedPercentage = direct.ValueOf(in.UtilizedPercentage)
	return out
}
func DiskEntry_FromProto(mapCtx *direct.MapContext, in *pb.DiskEntry) *krm.DiskEntry {
	if in == nil {
		return nil
	}
	out := &krm.DiskEntry{}
	out.CapacityBytes = direct.LazyPtr(in.GetCapacityBytes())
	out.FreeBytes = direct.LazyPtr(in.GetFreeBytes())
	out.DiskLabel = direct.LazyPtr(in.GetDiskLabel())
	out.DiskLabelType = direct.LazyPtr(in.GetDiskLabelType())
	out.InterfaceType = direct.Enum_FromProto(mapCtx, in.GetInterfaceType())
	out.Partitions = DiskPartitionList_FromProto(mapCtx, in.GetPartitions())
	out.HwAddress = direct.LazyPtr(in.GetHwAddress())
	out.Vmware = VmwareDiskConfig_FromProto(mapCtx, in.GetVmware())
	return out
}
func DiskEntry_ToProto(mapCtx *direct.MapContext, in *krm.DiskEntry) *pb.DiskEntry {
	if in == nil {
		return nil
	}
	out := &pb.DiskEntry{}
	out.CapacityBytes = direct.ValueOf(in.CapacityBytes)
	out.FreeBytes = direct.ValueOf(in.FreeBytes)
	out.DiskLabel = direct.ValueOf(in.DiskLabel)
	out.DiskLabelType = direct.ValueOf(in.DiskLabelType)
	out.InterfaceType = direct.Enum_ToProto[pb.DiskEntry_InterfaceType](mapCtx, in.InterfaceType)
	out.Partitions = DiskPartitionList_ToProto(mapCtx, in.Partitions)
	out.HwAddress = direct.ValueOf(in.HwAddress)
	if oneof := VmwareDiskConfig_ToProto(mapCtx, in.Vmware); oneof != nil {
		out.PlatformSpecific = &pb.DiskEntry_Vmware{Vmware: oneof}
	}
	return out
}
func DiskEntryList_FromProto(mapCtx *direct.MapContext, in *pb.DiskEntryList) *krm.DiskEntryList {
	if in == nil {
		return nil
	}
	out := &krm.DiskEntryList{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, DiskEntry_FromProto)
	return out
}
func DiskEntryList_ToProto(mapCtx *direct.MapContext, in *krm.DiskEntryList) *pb.DiskEntryList {
	if in == nil {
		return nil
	}
	out := &pb.DiskEntryList{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, DiskEntry_ToProto)
	return out
}
func DiskPartition_FromProto(mapCtx *direct.MapContext, in *pb.DiskPartition) *krm.DiskPartition {
	if in == nil {
		return nil
	}
	out := &krm.DiskPartition{}
	out.Type = direct.LazyPtr(in.GetType())
	out.FileSystem = direct.LazyPtr(in.GetFileSystem())
	out.MountPoint = direct.LazyPtr(in.GetMountPoint())
	out.CapacityBytes = direct.LazyPtr(in.GetCapacityBytes())
	out.FreeBytes = direct.LazyPtr(in.GetFreeBytes())
	out.Uuid = direct.LazyPtr(in.GetUuid())
	out.SubPartitions = DiskPartitionList_FromProto(mapCtx, in.GetSubPartitions())
	return out
}
func DiskPartition_ToProto(mapCtx *direct.MapContext, in *krm.DiskPartition) *pb.DiskPartition {
	if in == nil {
		return nil
	}
	out := &pb.DiskPartition{}
	out.Type = direct.ValueOf(in.Type)
	out.FileSystem = direct.ValueOf(in.FileSystem)
	out.MountPoint = direct.ValueOf(in.MountPoint)
	out.CapacityBytes = direct.ValueOf(in.CapacityBytes)
	out.FreeBytes = direct.ValueOf(in.FreeBytes)
	out.Uuid = direct.ValueOf(in.Uuid)
	out.SubPartitions = DiskPartitionList_ToProto(mapCtx, in.SubPartitions)
	return out
}
func DiskPartitionList_FromProto(mapCtx *direct.MapContext, in *pb.DiskPartitionList) *krm.DiskPartitionList {
	if in == nil {
		return nil
	}
	out := &krm.DiskPartitionList{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, DiskPartition_FromProto)
	return out
}
func DiskPartitionList_ToProto(mapCtx *direct.MapContext, in *krm.DiskPartitionList) *pb.DiskPartitionList {
	if in == nil {
		return nil
	}
	out := &pb.DiskPartitionList{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, DiskPartition_ToProto)
	return out
}
func DiskUsageSample_FromProto(mapCtx *direct.MapContext, in *pb.DiskUsageSample) *krm.DiskUsageSample {
	if in == nil {
		return nil
	}
	out := &krm.DiskUsageSample{}
	out.AverageIops = direct.LazyPtr(in.GetAverageIops())
	return out
}
func DiskUsageSample_ToProto(mapCtx *direct.MapContext, in *krm.DiskUsageSample) *pb.DiskUsageSample {
	if in == nil {
		return nil
	}
	out := &pb.DiskUsageSample{}
	out.AverageIops = direct.ValueOf(in.AverageIops)
	return out
}
func ErrorFrame_FromProto(mapCtx *direct.MapContext, in *pb.ErrorFrame) *krm.ErrorFrame {
	if in == nil {
		return nil
	}
	out := &krm.ErrorFrame{}
	// MISSING: Name
	// MISSING: Violations
	// MISSING: OriginalFrame
	// MISSING: IngestionTime
	return out
}
func ErrorFrame_ToProto(mapCtx *direct.MapContext, in *krm.ErrorFrame) *pb.ErrorFrame {
	if in == nil {
		return nil
	}
	out := &pb.ErrorFrame{}
	// MISSING: Name
	// MISSING: Violations
	// MISSING: OriginalFrame
	// MISSING: IngestionTime
	return out
}
func ErrorFrameObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ErrorFrame) *krm.ErrorFrameObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ErrorFrameObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Violations = direct.Slice_FromProto(mapCtx, in.Violations, FrameViolationEntry_FromProto)
	out.OriginalFrame = AssetFrame_FromProto(mapCtx, in.GetOriginalFrame())
	out.IngestionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetIngestionTime())
	return out
}
func ErrorFrameObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ErrorFrameObservedState) *pb.ErrorFrame {
	if in == nil {
		return nil
	}
	out := &pb.ErrorFrame{}
	out.Name = direct.ValueOf(in.Name)
	out.Violations = direct.Slice_ToProto(mapCtx, in.Violations, FrameViolationEntry_ToProto)
	out.OriginalFrame = AssetFrame_ToProto(mapCtx, in.OriginalFrame)
	out.IngestionTime = direct.StringTimestamp_ToProto(mapCtx, in.IngestionTime)
	return out
}
func FrameViolationEntry_FromProto(mapCtx *direct.MapContext, in *pb.FrameViolationEntry) *krm.FrameViolationEntry {
	if in == nil {
		return nil
	}
	out := &krm.FrameViolationEntry{}
	out.Field = direct.LazyPtr(in.GetField())
	out.Violation = direct.LazyPtr(in.GetViolation())
	return out
}
func FrameViolationEntry_ToProto(mapCtx *direct.MapContext, in *krm.FrameViolationEntry) *pb.FrameViolationEntry {
	if in == nil {
		return nil
	}
	out := &pb.FrameViolationEntry{}
	out.Field = direct.ValueOf(in.Field)
	out.Violation = direct.ValueOf(in.Violation)
	return out
}
func FstabEntry_FromProto(mapCtx *direct.MapContext, in *pb.FstabEntry) *krm.FstabEntry {
	if in == nil {
		return nil
	}
	out := &krm.FstabEntry{}
	out.Spec = direct.LazyPtr(in.GetSpec())
	out.File = direct.LazyPtr(in.GetFile())
	out.Vfstype = direct.LazyPtr(in.GetVfstype())
	out.Mntops = direct.LazyPtr(in.GetMntops())
	out.Freq = direct.LazyPtr(in.GetFreq())
	out.Passno = direct.LazyPtr(in.GetPassno())
	return out
}
func FstabEntry_ToProto(mapCtx *direct.MapContext, in *krm.FstabEntry) *pb.FstabEntry {
	if in == nil {
		return nil
	}
	out := &pb.FstabEntry{}
	out.Spec = direct.ValueOf(in.Spec)
	out.File = direct.ValueOf(in.File)
	out.Vfstype = direct.ValueOf(in.Vfstype)
	out.Mntops = direct.ValueOf(in.Mntops)
	out.Freq = direct.ValueOf(in.Freq)
	out.Passno = direct.ValueOf(in.Passno)
	return out
}
func FstabEntryList_FromProto(mapCtx *direct.MapContext, in *pb.FstabEntryList) *krm.FstabEntryList {
	if in == nil {
		return nil
	}
	out := &krm.FstabEntryList{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, FstabEntry_FromProto)
	return out
}
func FstabEntryList_ToProto(mapCtx *direct.MapContext, in *krm.FstabEntryList) *pb.FstabEntryList {
	if in == nil {
		return nil
	}
	out := &pb.FstabEntryList{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, FstabEntry_ToProto)
	return out
}
func GenericPlatformDetails_FromProto(mapCtx *direct.MapContext, in *pb.GenericPlatformDetails) *krm.GenericPlatformDetails {
	if in == nil {
		return nil
	}
	out := &krm.GenericPlatformDetails{}
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func GenericPlatformDetails_ToProto(mapCtx *direct.MapContext, in *krm.GenericPlatformDetails) *pb.GenericPlatformDetails {
	if in == nil {
		return nil
	}
	out := &pb.GenericPlatformDetails{}
	out.Location = direct.ValueOf(in.Location)
	return out
}
func GuestConfigDetails_FromProto(mapCtx *direct.MapContext, in *pb.GuestConfigDetails) *krm.GuestConfigDetails {
	if in == nil {
		return nil
	}
	out := &krm.GuestConfigDetails{}
	out.Issue = direct.LazyPtr(in.GetIssue())
	out.Fstab = FstabEntryList_FromProto(mapCtx, in.GetFstab())
	out.Hosts = HostsEntryList_FromProto(mapCtx, in.GetHosts())
	out.NfsExports = NfsExportList_FromProto(mapCtx, in.GetNfsExports())
	out.SelinuxMode = direct.Enum_FromProto(mapCtx, in.GetSelinuxMode())
	return out
}
func GuestConfigDetails_ToProto(mapCtx *direct.MapContext, in *krm.GuestConfigDetails) *pb.GuestConfigDetails {
	if in == nil {
		return nil
	}
	out := &pb.GuestConfigDetails{}
	out.Issue = direct.ValueOf(in.Issue)
	out.Fstab = FstabEntryList_ToProto(mapCtx, in.Fstab)
	out.Hosts = HostsEntryList_ToProto(mapCtx, in.Hosts)
	out.NfsExports = NfsExportList_ToProto(mapCtx, in.NfsExports)
	out.SelinuxMode = direct.Enum_ToProto[pb.GuestConfigDetails_SeLinuxMode](mapCtx, in.SelinuxMode)
	return out
}
func GuestInstalledApplication_FromProto(mapCtx *direct.MapContext, in *pb.GuestInstalledApplication) *krm.GuestInstalledApplication {
	if in == nil {
		return nil
	}
	out := &krm.GuestInstalledApplication{}
	out.ApplicationName = direct.LazyPtr(in.GetApplicationName())
	out.Vendor = direct.LazyPtr(in.GetVendor())
	out.InstallTime = direct.StringTimestamp_FromProto(mapCtx, in.GetInstallTime())
	out.Path = direct.LazyPtr(in.GetPath())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func GuestInstalledApplication_ToProto(mapCtx *direct.MapContext, in *krm.GuestInstalledApplication) *pb.GuestInstalledApplication {
	if in == nil {
		return nil
	}
	out := &pb.GuestInstalledApplication{}
	out.ApplicationName = direct.ValueOf(in.ApplicationName)
	out.Vendor = direct.ValueOf(in.Vendor)
	out.InstallTime = direct.StringTimestamp_ToProto(mapCtx, in.InstallTime)
	out.Path = direct.ValueOf(in.Path)
	out.Version = direct.ValueOf(in.Version)
	return out
}
func GuestInstalledApplicationList_FromProto(mapCtx *direct.MapContext, in *pb.GuestInstalledApplicationList) *krm.GuestInstalledApplicationList {
	if in == nil {
		return nil
	}
	out := &krm.GuestInstalledApplicationList{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, GuestInstalledApplication_FromProto)
	return out
}
func GuestInstalledApplicationList_ToProto(mapCtx *direct.MapContext, in *krm.GuestInstalledApplicationList) *pb.GuestInstalledApplicationList {
	if in == nil {
		return nil
	}
	out := &pb.GuestInstalledApplicationList{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, GuestInstalledApplication_ToProto)
	return out
}
func GuestOsDetails_FromProto(mapCtx *direct.MapContext, in *pb.GuestOsDetails) *krm.GuestOsDetails {
	if in == nil {
		return nil
	}
	out := &krm.GuestOsDetails{}
	out.OsName = direct.LazyPtr(in.GetOsName())
	out.Family = direct.Enum_FromProto(mapCtx, in.GetFamily())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Config = GuestConfigDetails_FromProto(mapCtx, in.GetConfig())
	out.Runtime = GuestRuntimeDetails_FromProto(mapCtx, in.GetRuntime())
	return out
}
func GuestOsDetails_ToProto(mapCtx *direct.MapContext, in *krm.GuestOsDetails) *pb.GuestOsDetails {
	if in == nil {
		return nil
	}
	out := &pb.GuestOsDetails{}
	out.OsName = direct.ValueOf(in.OsName)
	out.Family = direct.Enum_ToProto[pb.OperatingSystemFamily](mapCtx, in.Family)
	out.Version = direct.ValueOf(in.Version)
	out.Config = GuestConfigDetails_ToProto(mapCtx, in.Config)
	out.Runtime = GuestRuntimeDetails_ToProto(mapCtx, in.Runtime)
	return out
}
func GuestRuntimeDetails_FromProto(mapCtx *direct.MapContext, in *pb.GuestRuntimeDetails) *krm.GuestRuntimeDetails {
	if in == nil {
		return nil
	}
	out := &krm.GuestRuntimeDetails{}
	out.Services = RunningServiceList_FromProto(mapCtx, in.GetServices())
	out.Processes = RunningProcessList_FromProto(mapCtx, in.GetProcesses())
	out.Network = RuntimeNetworkInfo_FromProto(mapCtx, in.GetNetwork())
	out.LastBootTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastBootTime())
	out.Domain = direct.LazyPtr(in.GetDomain())
	out.MachineName = direct.LazyPtr(in.GetMachineName())
	out.InstalledApps = GuestInstalledApplicationList_FromProto(mapCtx, in.GetInstalledApps())
	out.OpenFileList = OpenFileList_FromProto(mapCtx, in.GetOpenFileList())
	return out
}
func GuestRuntimeDetails_ToProto(mapCtx *direct.MapContext, in *krm.GuestRuntimeDetails) *pb.GuestRuntimeDetails {
	if in == nil {
		return nil
	}
	out := &pb.GuestRuntimeDetails{}
	out.Services = RunningServiceList_ToProto(mapCtx, in.Services)
	out.Processes = RunningProcessList_ToProto(mapCtx, in.Processes)
	out.Network = RuntimeNetworkInfo_ToProto(mapCtx, in.Network)
	out.LastBootTime = direct.StringTimestamp_ToProto(mapCtx, in.LastBootTime)
	out.Domain = direct.ValueOf(in.Domain)
	out.MachineName = direct.ValueOf(in.MachineName)
	out.InstalledApps = GuestInstalledApplicationList_ToProto(mapCtx, in.InstalledApps)
	out.OpenFileList = OpenFileList_ToProto(mapCtx, in.OpenFileList)
	return out
}
func HostsEntry_FromProto(mapCtx *direct.MapContext, in *pb.HostsEntry) *krm.HostsEntry {
	if in == nil {
		return nil
	}
	out := &krm.HostsEntry{}
	out.IP = direct.LazyPtr(in.GetIp())
	out.HostNames = in.HostNames
	return out
}
func HostsEntry_ToProto(mapCtx *direct.MapContext, in *krm.HostsEntry) *pb.HostsEntry {
	if in == nil {
		return nil
	}
	out := &pb.HostsEntry{}
	out.Ip = direct.ValueOf(in.IP)
	out.HostNames = in.HostNames
	return out
}
func HostsEntryList_FromProto(mapCtx *direct.MapContext, in *pb.HostsEntryList) *krm.HostsEntryList {
	if in == nil {
		return nil
	}
	out := &krm.HostsEntryList{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, HostsEntry_FromProto)
	return out
}
func HostsEntryList_ToProto(mapCtx *direct.MapContext, in *krm.HostsEntryList) *pb.HostsEntryList {
	if in == nil {
		return nil
	}
	out := &pb.HostsEntryList{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, HostsEntry_ToProto)
	return out
}
func MachineArchitectureDetails_FromProto(mapCtx *direct.MapContext, in *pb.MachineArchitectureDetails) *krm.MachineArchitectureDetails {
	if in == nil {
		return nil
	}
	out := &krm.MachineArchitectureDetails{}
	out.CpuArchitecture = direct.LazyPtr(in.GetCpuArchitecture())
	out.CpuName = direct.LazyPtr(in.GetCpuName())
	out.Vendor = direct.LazyPtr(in.GetVendor())
	out.CpuThreadCount = direct.LazyPtr(in.GetCpuThreadCount())
	out.CpuSocketCount = direct.LazyPtr(in.GetCpuSocketCount())
	out.Bios = BiosDetails_FromProto(mapCtx, in.GetBios())
	out.FirmwareType = direct.Enum_FromProto(mapCtx, in.GetFirmwareType())
	out.Hyperthreading = direct.Enum_FromProto(mapCtx, in.GetHyperthreading())
	return out
}
func MachineArchitectureDetails_ToProto(mapCtx *direct.MapContext, in *krm.MachineArchitectureDetails) *pb.MachineArchitectureDetails {
	if in == nil {
		return nil
	}
	out := &pb.MachineArchitectureDetails{}
	out.CpuArchitecture = direct.ValueOf(in.CpuArchitecture)
	out.CpuName = direct.ValueOf(in.CpuName)
	out.Vendor = direct.ValueOf(in.Vendor)
	out.CpuThreadCount = direct.ValueOf(in.CpuThreadCount)
	out.CpuSocketCount = direct.ValueOf(in.CpuSocketCount)
	out.Bios = BiosDetails_ToProto(mapCtx, in.Bios)
	out.FirmwareType = direct.Enum_ToProto[pb.MachineArchitectureDetails_FirmwareType](mapCtx, in.FirmwareType)
	out.Hyperthreading = direct.Enum_ToProto[pb.MachineArchitectureDetails_CpuHyperThreading](mapCtx, in.Hyperthreading)
	return out
}
func MachineDetails_FromProto(mapCtx *direct.MapContext, in *pb.MachineDetails) *krm.MachineDetails {
	if in == nil {
		return nil
	}
	out := &krm.MachineDetails{}
	out.Uuid = direct.LazyPtr(in.GetUuid())
	out.MachineName = direct.LazyPtr(in.GetMachineName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.CoreCount = direct.LazyPtr(in.GetCoreCount())
	out.MemoryMb = direct.LazyPtr(in.GetMemoryMb())
	out.PowerState = direct.Enum_FromProto(mapCtx, in.GetPowerState())
	out.Architecture = MachineArchitectureDetails_FromProto(mapCtx, in.GetArchitecture())
	out.GuestOs = GuestOsDetails_FromProto(mapCtx, in.GetGuestOs())
	out.Network = MachineNetworkDetails_FromProto(mapCtx, in.GetNetwork())
	out.Disks = MachineDiskDetails_FromProto(mapCtx, in.GetDisks())
	out.Platform = PlatformDetails_FromProto(mapCtx, in.GetPlatform())
	return out
}
func MachineDetails_ToProto(mapCtx *direct.MapContext, in *krm.MachineDetails) *pb.MachineDetails {
	if in == nil {
		return nil
	}
	out := &pb.MachineDetails{}
	out.Uuid = direct.ValueOf(in.Uuid)
	out.MachineName = direct.ValueOf(in.MachineName)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.CoreCount = direct.ValueOf(in.CoreCount)
	out.MemoryMb = direct.ValueOf(in.MemoryMb)
	out.PowerState = direct.Enum_ToProto[pb.MachineDetails_PowerState](mapCtx, in.PowerState)
	out.Architecture = MachineArchitectureDetails_ToProto(mapCtx, in.Architecture)
	out.GuestOs = GuestOsDetails_ToProto(mapCtx, in.GuestOs)
	out.Network = MachineNetworkDetails_ToProto(mapCtx, in.Network)
	out.Disks = MachineDiskDetails_ToProto(mapCtx, in.Disks)
	out.Platform = PlatformDetails_ToProto(mapCtx, in.Platform)
	return out
}
func MachineDiskDetails_FromProto(mapCtx *direct.MapContext, in *pb.MachineDiskDetails) *krm.MachineDiskDetails {
	if in == nil {
		return nil
	}
	out := &krm.MachineDiskDetails{}
	out.TotalCapacityBytes = direct.LazyPtr(in.GetTotalCapacityBytes())
	out.TotalFreeBytes = direct.LazyPtr(in.GetTotalFreeBytes())
	out.Disks = DiskEntryList_FromProto(mapCtx, in.GetDisks())
	return out
}
func MachineDiskDetails_ToProto(mapCtx *direct.MapContext, in *krm.MachineDiskDetails) *pb.MachineDiskDetails {
	if in == nil {
		return nil
	}
	out := &pb.MachineDiskDetails{}
	out.TotalCapacityBytes = direct.ValueOf(in.TotalCapacityBytes)
	out.TotalFreeBytes = direct.ValueOf(in.TotalFreeBytes)
	out.Disks = DiskEntryList_ToProto(mapCtx, in.Disks)
	return out
}
func MachineNetworkDetails_FromProto(mapCtx *direct.MapContext, in *pb.MachineNetworkDetails) *krm.MachineNetworkDetails {
	if in == nil {
		return nil
	}
	out := &krm.MachineNetworkDetails{}
	out.PrimaryIPAddress = direct.LazyPtr(in.GetPrimaryIpAddress())
	out.PublicIPAddress = direct.LazyPtr(in.GetPublicIpAddress())
	out.PrimaryMacAddress = direct.LazyPtr(in.GetPrimaryMacAddress())
	out.Adapters = NetworkAdapterList_FromProto(mapCtx, in.GetAdapters())
	return out
}
func MachineNetworkDetails_ToProto(mapCtx *direct.MapContext, in *krm.MachineNetworkDetails) *pb.MachineNetworkDetails {
	if in == nil {
		return nil
	}
	out := &pb.MachineNetworkDetails{}
	out.PrimaryIpAddress = direct.ValueOf(in.PrimaryIPAddress)
	out.PublicIpAddress = direct.ValueOf(in.PublicIPAddress)
	out.PrimaryMacAddress = direct.ValueOf(in.PrimaryMacAddress)
	out.Adapters = NetworkAdapterList_ToProto(mapCtx, in.Adapters)
	return out
}
func MemoryUsageSample_FromProto(mapCtx *direct.MapContext, in *pb.MemoryUsageSample) *krm.MemoryUsageSample {
	if in == nil {
		return nil
	}
	out := &krm.MemoryUsageSample{}
	out.UtilizedPercentage = direct.LazyPtr(in.GetUtilizedPercentage())
	return out
}
func MemoryUsageSample_ToProto(mapCtx *direct.MapContext, in *krm.MemoryUsageSample) *pb.MemoryUsageSample {
	if in == nil {
		return nil
	}
	out := &pb.MemoryUsageSample{}
	out.UtilizedPercentage = direct.ValueOf(in.UtilizedPercentage)
	return out
}
func MigrationcenterErrorFrameObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ErrorFrame) *krm.MigrationcenterErrorFrameObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterErrorFrameObservedState{}
	// MISSING: Name
	// MISSING: Violations
	// MISSING: OriginalFrame
	// MISSING: IngestionTime
	return out
}
func MigrationcenterErrorFrameObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterErrorFrameObservedState) *pb.ErrorFrame {
	if in == nil {
		return nil
	}
	out := &pb.ErrorFrame{}
	// MISSING: Name
	// MISSING: Violations
	// MISSING: OriginalFrame
	// MISSING: IngestionTime
	return out
}
func MigrationcenterErrorFrameSpec_FromProto(mapCtx *direct.MapContext, in *pb.ErrorFrame) *krm.MigrationcenterErrorFrameSpec {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterErrorFrameSpec{}
	// MISSING: Name
	// MISSING: Violations
	// MISSING: OriginalFrame
	// MISSING: IngestionTime
	return out
}
func MigrationcenterErrorFrameSpec_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterErrorFrameSpec) *pb.ErrorFrame {
	if in == nil {
		return nil
	}
	out := &pb.ErrorFrame{}
	// MISSING: Name
	// MISSING: Violations
	// MISSING: OriginalFrame
	// MISSING: IngestionTime
	return out
}
func NetworkAdapterDetails_FromProto(mapCtx *direct.MapContext, in *pb.NetworkAdapterDetails) *krm.NetworkAdapterDetails {
	if in == nil {
		return nil
	}
	out := &krm.NetworkAdapterDetails{}
	out.AdapterType = direct.LazyPtr(in.GetAdapterType())
	out.MacAddress = direct.LazyPtr(in.GetMacAddress())
	out.Addresses = NetworkAddressList_FromProto(mapCtx, in.GetAddresses())
	return out
}
func NetworkAdapterDetails_ToProto(mapCtx *direct.MapContext, in *krm.NetworkAdapterDetails) *pb.NetworkAdapterDetails {
	if in == nil {
		return nil
	}
	out := &pb.NetworkAdapterDetails{}
	out.AdapterType = direct.ValueOf(in.AdapterType)
	out.MacAddress = direct.ValueOf(in.MacAddress)
	out.Addresses = NetworkAddressList_ToProto(mapCtx, in.Addresses)
	return out
}
func NetworkAdapterList_FromProto(mapCtx *direct.MapContext, in *pb.NetworkAdapterList) *krm.NetworkAdapterList {
	if in == nil {
		return nil
	}
	out := &krm.NetworkAdapterList{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, NetworkAdapterDetails_FromProto)
	return out
}
func NetworkAdapterList_ToProto(mapCtx *direct.MapContext, in *krm.NetworkAdapterList) *pb.NetworkAdapterList {
	if in == nil {
		return nil
	}
	out := &pb.NetworkAdapterList{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, NetworkAdapterDetails_ToProto)
	return out
}
func NetworkAddress_FromProto(mapCtx *direct.MapContext, in *pb.NetworkAddress) *krm.NetworkAddress {
	if in == nil {
		return nil
	}
	out := &krm.NetworkAddress{}
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.SubnetMask = direct.LazyPtr(in.GetSubnetMask())
	out.Bcast = direct.LazyPtr(in.GetBcast())
	out.Fqdn = direct.LazyPtr(in.GetFqdn())
	out.Assignment = direct.Enum_FromProto(mapCtx, in.GetAssignment())
	return out
}
func NetworkAddress_ToProto(mapCtx *direct.MapContext, in *krm.NetworkAddress) *pb.NetworkAddress {
	if in == nil {
		return nil
	}
	out := &pb.NetworkAddress{}
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.SubnetMask = direct.ValueOf(in.SubnetMask)
	out.Bcast = direct.ValueOf(in.Bcast)
	out.Fqdn = direct.ValueOf(in.Fqdn)
	out.Assignment = direct.Enum_ToProto[pb.NetworkAddress_AddressAssignment](mapCtx, in.Assignment)
	return out
}
func NetworkAddressList_FromProto(mapCtx *direct.MapContext, in *pb.NetworkAddressList) *krm.NetworkAddressList {
	if in == nil {
		return nil
	}
	out := &krm.NetworkAddressList{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, NetworkAddress_FromProto)
	return out
}
func NetworkAddressList_ToProto(mapCtx *direct.MapContext, in *krm.NetworkAddressList) *pb.NetworkAddressList {
	if in == nil {
		return nil
	}
	out := &pb.NetworkAddressList{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, NetworkAddress_ToProto)
	return out
}
func NetworkConnection_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConnection) *krm.NetworkConnection {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnection{}
	out.Protocol = direct.LazyPtr(in.GetProtocol())
	out.LocalIPAddress = direct.LazyPtr(in.GetLocalIpAddress())
	out.LocalPort = direct.LazyPtr(in.GetLocalPort())
	out.RemoteIPAddress = direct.LazyPtr(in.GetRemoteIpAddress())
	out.RemotePort = direct.LazyPtr(in.GetRemotePort())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Pid = direct.LazyPtr(in.GetPid())
	out.ProcessName = direct.LazyPtr(in.GetProcessName())
	return out
}
func NetworkConnection_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnection) *pb.NetworkConnection {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConnection{}
	out.Protocol = direct.ValueOf(in.Protocol)
	out.LocalIpAddress = direct.ValueOf(in.LocalIPAddress)
	out.LocalPort = direct.ValueOf(in.LocalPort)
	out.RemoteIpAddress = direct.ValueOf(in.RemoteIPAddress)
	out.RemotePort = direct.ValueOf(in.RemotePort)
	out.State = direct.Enum_ToProto[pb.NetworkConnection_State](mapCtx, in.State)
	out.Pid = direct.ValueOf(in.Pid)
	out.ProcessName = direct.ValueOf(in.ProcessName)
	return out
}
func NetworkConnectionList_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConnectionList) *krm.NetworkConnectionList {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectionList{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, NetworkConnection_FromProto)
	return out
}
func NetworkConnectionList_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectionList) *pb.NetworkConnectionList {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConnectionList{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, NetworkConnection_ToProto)
	return out
}
func NetworkUsageSample_FromProto(mapCtx *direct.MapContext, in *pb.NetworkUsageSample) *krm.NetworkUsageSample {
	if in == nil {
		return nil
	}
	out := &krm.NetworkUsageSample{}
	out.AverageIngressBps = direct.LazyPtr(in.GetAverageIngressBps())
	out.AverageEgressBps = direct.LazyPtr(in.GetAverageEgressBps())
	return out
}
func NetworkUsageSample_ToProto(mapCtx *direct.MapContext, in *krm.NetworkUsageSample) *pb.NetworkUsageSample {
	if in == nil {
		return nil
	}
	out := &pb.NetworkUsageSample{}
	out.AverageIngressBps = direct.ValueOf(in.AverageIngressBps)
	out.AverageEgressBps = direct.ValueOf(in.AverageEgressBps)
	return out
}
func NfsExport_FromProto(mapCtx *direct.MapContext, in *pb.NfsExport) *krm.NfsExport {
	if in == nil {
		return nil
	}
	out := &krm.NfsExport{}
	out.ExportDirectory = direct.LazyPtr(in.GetExportDirectory())
	out.Hosts = in.Hosts
	return out
}
func NfsExport_ToProto(mapCtx *direct.MapContext, in *krm.NfsExport) *pb.NfsExport {
	if in == nil {
		return nil
	}
	out := &pb.NfsExport{}
	out.ExportDirectory = direct.ValueOf(in.ExportDirectory)
	out.Hosts = in.Hosts
	return out
}
func NfsExportList_FromProto(mapCtx *direct.MapContext, in *pb.NfsExportList) *krm.NfsExportList {
	if in == nil {
		return nil
	}
	out := &krm.NfsExportList{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, NfsExport_FromProto)
	return out
}
func NfsExportList_ToProto(mapCtx *direct.MapContext, in *krm.NfsExportList) *pb.NfsExportList {
	if in == nil {
		return nil
	}
	out := &pb.NfsExportList{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, NfsExport_ToProto)
	return out
}
func OpenFileDetails_FromProto(mapCtx *direct.MapContext, in *pb.OpenFileDetails) *krm.OpenFileDetails {
	if in == nil {
		return nil
	}
	out := &krm.OpenFileDetails{}
	out.Command = direct.LazyPtr(in.GetCommand())
	out.User = direct.LazyPtr(in.GetUser())
	out.FileType = direct.LazyPtr(in.GetFileType())
	out.FilePath = direct.LazyPtr(in.GetFilePath())
	return out
}
func OpenFileDetails_ToProto(mapCtx *direct.MapContext, in *krm.OpenFileDetails) *pb.OpenFileDetails {
	if in == nil {
		return nil
	}
	out := &pb.OpenFileDetails{}
	out.Command = direct.ValueOf(in.Command)
	out.User = direct.ValueOf(in.User)
	out.FileType = direct.ValueOf(in.FileType)
	out.FilePath = direct.ValueOf(in.FilePath)
	return out
}
func OpenFileList_FromProto(mapCtx *direct.MapContext, in *pb.OpenFileList) *krm.OpenFileList {
	if in == nil {
		return nil
	}
	out := &krm.OpenFileList{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, OpenFileDetails_FromProto)
	return out
}
func OpenFileList_ToProto(mapCtx *direct.MapContext, in *krm.OpenFileList) *pb.OpenFileList {
	if in == nil {
		return nil
	}
	out := &pb.OpenFileList{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, OpenFileDetails_ToProto)
	return out
}
func PerformanceSample_FromProto(mapCtx *direct.MapContext, in *pb.PerformanceSample) *krm.PerformanceSample {
	if in == nil {
		return nil
	}
	out := &krm.PerformanceSample{}
	out.SampleTime = direct.StringTimestamp_FromProto(mapCtx, in.GetSampleTime())
	out.Memory = MemoryUsageSample_FromProto(mapCtx, in.GetMemory())
	out.Cpu = CpuUsageSample_FromProto(mapCtx, in.GetCpu())
	out.Network = NetworkUsageSample_FromProto(mapCtx, in.GetNetwork())
	out.Disk = DiskUsageSample_FromProto(mapCtx, in.GetDisk())
	return out
}
func PerformanceSample_ToProto(mapCtx *direct.MapContext, in *krm.PerformanceSample) *pb.PerformanceSample {
	if in == nil {
		return nil
	}
	out := &pb.PerformanceSample{}
	out.SampleTime = direct.StringTimestamp_ToProto(mapCtx, in.SampleTime)
	out.Memory = MemoryUsageSample_ToProto(mapCtx, in.Memory)
	out.Cpu = CpuUsageSample_ToProto(mapCtx, in.Cpu)
	out.Network = NetworkUsageSample_ToProto(mapCtx, in.Network)
	out.Disk = DiskUsageSample_ToProto(mapCtx, in.Disk)
	return out
}
func PhysicalPlatformDetails_FromProto(mapCtx *direct.MapContext, in *pb.PhysicalPlatformDetails) *krm.PhysicalPlatformDetails {
	if in == nil {
		return nil
	}
	out := &krm.PhysicalPlatformDetails{}
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func PhysicalPlatformDetails_ToProto(mapCtx *direct.MapContext, in *krm.PhysicalPlatformDetails) *pb.PhysicalPlatformDetails {
	if in == nil {
		return nil
	}
	out := &pb.PhysicalPlatformDetails{}
	out.Location = direct.ValueOf(in.Location)
	return out
}
func PlatformDetails_FromProto(mapCtx *direct.MapContext, in *pb.PlatformDetails) *krm.PlatformDetails {
	if in == nil {
		return nil
	}
	out := &krm.PlatformDetails{}
	out.VmwareDetails = VmwarePlatformDetails_FromProto(mapCtx, in.GetVmwareDetails())
	out.AwsEc2Details = AwsEc2PlatformDetails_FromProto(mapCtx, in.GetAwsEc2Details())
	out.AzureVmDetails = AzureVmPlatformDetails_FromProto(mapCtx, in.GetAzureVmDetails())
	out.GenericDetails = GenericPlatformDetails_FromProto(mapCtx, in.GetGenericDetails())
	out.PhysicalDetails = PhysicalPlatformDetails_FromProto(mapCtx, in.GetPhysicalDetails())
	return out
}
func PlatformDetails_ToProto(mapCtx *direct.MapContext, in *krm.PlatformDetails) *pb.PlatformDetails {
	if in == nil {
		return nil
	}
	out := &pb.PlatformDetails{}
	if oneof := VmwarePlatformDetails_ToProto(mapCtx, in.VmwareDetails); oneof != nil {
		out.VendorDetails = &pb.PlatformDetails_VmwareDetails{VmwareDetails: oneof}
	}
	if oneof := AwsEc2PlatformDetails_ToProto(mapCtx, in.AwsEc2Details); oneof != nil {
		out.VendorDetails = &pb.PlatformDetails_AwsEc2Details{AwsEc2Details: oneof}
	}
	if oneof := AzureVmPlatformDetails_ToProto(mapCtx, in.AzureVmDetails); oneof != nil {
		out.VendorDetails = &pb.PlatformDetails_AzureVmDetails{AzureVmDetails: oneof}
	}
	if oneof := GenericPlatformDetails_ToProto(mapCtx, in.GenericDetails); oneof != nil {
		out.VendorDetails = &pb.PlatformDetails_GenericDetails{GenericDetails: oneof}
	}
	if oneof := PhysicalPlatformDetails_ToProto(mapCtx, in.PhysicalDetails); oneof != nil {
		out.VendorDetails = &pb.PlatformDetails_PhysicalDetails{PhysicalDetails: oneof}
	}
	return out
}
func RunningProcess_FromProto(mapCtx *direct.MapContext, in *pb.RunningProcess) *krm.RunningProcess {
	if in == nil {
		return nil
	}
	out := &krm.RunningProcess{}
	out.Pid = direct.LazyPtr(in.GetPid())
	out.ExePath = direct.LazyPtr(in.GetExePath())
	out.Cmdline = direct.LazyPtr(in.GetCmdline())
	out.User = direct.LazyPtr(in.GetUser())
	out.Attributes = in.Attributes
	return out
}
func RunningProcess_ToProto(mapCtx *direct.MapContext, in *krm.RunningProcess) *pb.RunningProcess {
	if in == nil {
		return nil
	}
	out := &pb.RunningProcess{}
	out.Pid = direct.ValueOf(in.Pid)
	out.ExePath = direct.ValueOf(in.ExePath)
	out.Cmdline = direct.ValueOf(in.Cmdline)
	out.User = direct.ValueOf(in.User)
	out.Attributes = in.Attributes
	return out
}
func RunningProcessList_FromProto(mapCtx *direct.MapContext, in *pb.RunningProcessList) *krm.RunningProcessList {
	if in == nil {
		return nil
	}
	out := &krm.RunningProcessList{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, RunningProcess_FromProto)
	return out
}
func RunningProcessList_ToProto(mapCtx *direct.MapContext, in *krm.RunningProcessList) *pb.RunningProcessList {
	if in == nil {
		return nil
	}
	out := &pb.RunningProcessList{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, RunningProcess_ToProto)
	return out
}
func RunningService_FromProto(mapCtx *direct.MapContext, in *pb.RunningService) *krm.RunningService {
	if in == nil {
		return nil
	}
	out := &krm.RunningService{}
	out.ServiceName = direct.LazyPtr(in.GetServiceName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StartMode = direct.Enum_FromProto(mapCtx, in.GetStartMode())
	out.ExePath = direct.LazyPtr(in.GetExePath())
	out.Cmdline = direct.LazyPtr(in.GetCmdline())
	out.Pid = direct.LazyPtr(in.GetPid())
	return out
}
func RunningService_ToProto(mapCtx *direct.MapContext, in *krm.RunningService) *pb.RunningService {
	if in == nil {
		return nil
	}
	out := &pb.RunningService{}
	out.ServiceName = direct.ValueOf(in.ServiceName)
	out.State = direct.Enum_ToProto[pb.RunningService_State](mapCtx, in.State)
	out.StartMode = direct.Enum_ToProto[pb.RunningService_StartMode](mapCtx, in.StartMode)
	out.ExePath = direct.ValueOf(in.ExePath)
	out.Cmdline = direct.ValueOf(in.Cmdline)
	out.Pid = direct.ValueOf(in.Pid)
	return out
}
func RunningServiceList_FromProto(mapCtx *direct.MapContext, in *pb.RunningServiceList) *krm.RunningServiceList {
	if in == nil {
		return nil
	}
	out := &krm.RunningServiceList{}
	out.Entries = direct.Slice_FromProto(mapCtx, in.Entries, RunningService_FromProto)
	return out
}
func RunningServiceList_ToProto(mapCtx *direct.MapContext, in *krm.RunningServiceList) *pb.RunningServiceList {
	if in == nil {
		return nil
	}
	out := &pb.RunningServiceList{}
	out.Entries = direct.Slice_ToProto(mapCtx, in.Entries, RunningService_ToProto)
	return out
}
func RuntimeNetworkInfo_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeNetworkInfo) *krm.RuntimeNetworkInfo {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeNetworkInfo{}
	out.ScanTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScanTime())
	out.Connections = NetworkConnectionList_FromProto(mapCtx, in.GetConnections())
	return out
}
func RuntimeNetworkInfo_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeNetworkInfo) *pb.RuntimeNetworkInfo {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeNetworkInfo{}
	out.ScanTime = direct.StringTimestamp_ToProto(mapCtx, in.ScanTime)
	out.Connections = NetworkConnectionList_ToProto(mapCtx, in.Connections)
	return out
}
func VmwareDiskConfig_FromProto(mapCtx *direct.MapContext, in *pb.VmwareDiskConfig) *krm.VmwareDiskConfig {
	if in == nil {
		return nil
	}
	out := &krm.VmwareDiskConfig{}
	out.BackingType = direct.Enum_FromProto(mapCtx, in.GetBackingType())
	out.Shared = direct.LazyPtr(in.GetShared())
	out.VmdkMode = direct.Enum_FromProto(mapCtx, in.GetVmdkMode())
	out.RdmCompatibility = direct.Enum_FromProto(mapCtx, in.GetRdmCompatibility())
	return out
}
func VmwareDiskConfig_ToProto(mapCtx *direct.MapContext, in *krm.VmwareDiskConfig) *pb.VmwareDiskConfig {
	if in == nil {
		return nil
	}
	out := &pb.VmwareDiskConfig{}
	out.BackingType = direct.Enum_ToProto[pb.VmwareDiskConfig_BackingType](mapCtx, in.BackingType)
	out.Shared = direct.ValueOf(in.Shared)
	out.VmdkMode = direct.Enum_ToProto[pb.VmwareDiskConfig_VmdkMode](mapCtx, in.VmdkMode)
	out.RdmCompatibility = direct.Enum_ToProto[pb.VmwareDiskConfig_RdmCompatibility](mapCtx, in.RdmCompatibility)
	return out
}
func VmwarePlatformDetails_FromProto(mapCtx *direct.MapContext, in *pb.VmwarePlatformDetails) *krm.VmwarePlatformDetails {
	if in == nil {
		return nil
	}
	out := &krm.VmwarePlatformDetails{}
	out.VcenterVersion = direct.LazyPtr(in.GetVcenterVersion())
	out.EsxVersion = direct.LazyPtr(in.GetEsxVersion())
	out.Osid = direct.LazyPtr(in.GetOsid())
	out.VcenterFolder = direct.LazyPtr(in.GetVcenterFolder())
	out.VcenterURI = direct.LazyPtr(in.GetVcenterUri())
	out.VcenterVmID = direct.LazyPtr(in.GetVcenterVmId())
	return out
}
func VmwarePlatformDetails_ToProto(mapCtx *direct.MapContext, in *krm.VmwarePlatformDetails) *pb.VmwarePlatformDetails {
	if in == nil {
		return nil
	}
	out := &pb.VmwarePlatformDetails{}
	out.VcenterVersion = direct.ValueOf(in.VcenterVersion)
	out.EsxVersion = direct.ValueOf(in.EsxVersion)
	out.Osid = direct.ValueOf(in.Osid)
	out.VcenterFolder = direct.ValueOf(in.VcenterFolder)
	out.VcenterUri = direct.ValueOf(in.VcenterURI)
	out.VcenterVmId = direct.ValueOf(in.VcenterVmID)
	return out
}
