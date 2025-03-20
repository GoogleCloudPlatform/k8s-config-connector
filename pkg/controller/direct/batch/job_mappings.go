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

package batch

import (
	pb "cloud.google.com/go/batch/apiv1/batchpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/batch/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BatchJobStatus_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.BatchJobStatus {
	if in == nil {
		return nil
	}
	out := &krm.BatchJobStatus{}
	out.ObservedState = BatchJobObservedState_FromProto(mapCtx, in)
	return out
}
func BatchJobStatus_ToProto(mapCtx *direct.MapContext, in *krm.BatchJobStatus) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out = BatchJobObservedState_ToProto(mapCtx, in.ObservedState)
	return out
}
func BatchJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.BatchJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BatchJobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.TaskGroups = direct.Slice_FromProto(mapCtx, in.TaskGroups, TaskGroupObservedState_FromProto)
	out.Status = JobStatus_FromProto(mapCtx, in.GetStatus())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func BatchJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BatchJobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.TaskGroups = direct.Slice_ToProto(mapCtx, in.TaskGroups, TaskGroupObservedState_ToProto)
	out.Status = JobStatus_ToProto(mapCtx, in.Status)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func BatchJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.BatchJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.BatchJobSpec{}
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.TaskGroups = direct.Slice_FromProto(mapCtx, in.TaskGroups, TaskGroup_FromProto)
	out.AllocationPolicy = AllocationPolicy_FromProto(mapCtx, in.GetAllocationPolicy())
	out.Labels = in.Labels
	out.LogsPolicy = LogsPolicy_FromProto(mapCtx, in.GetLogsPolicy())
	out.Notifications = direct.Slice_FromProto(mapCtx, in.Notifications, JobNotification_FromProto)
	return out
}
func BatchJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.BatchJobSpec) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Priority = direct.ValueOf(in.Priority)
	out.TaskGroups = direct.Slice_ToProto(mapCtx, in.TaskGroups, TaskGroup_ToProto)
	out.AllocationPolicy = AllocationPolicy_ToProto(mapCtx, in.AllocationPolicy)
	out.Labels = in.Labels
	out.LogsPolicy = LogsPolicy_ToProto(mapCtx, in.LogsPolicy)
	out.Notifications = direct.Slice_ToProto(mapCtx, in.Notifications, JobNotification_ToProto)
	return out
}
func AllocationPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AllocationPolicy) *krm.AllocationPolicy {
	if in == nil {
		return nil
	}
	out := &krm.AllocationPolicy{}
	out.Location = AllocationPolicy_LocationPolicy_FromProto(mapCtx, in.GetLocation())
	out.Instances = direct.Slice_FromProto(mapCtx, in.Instances, AllocationPolicy_InstancePolicyOrTemplate_FromProto)
	out.ServiceAccountRef = &v1beta1.IAMServiceAccountRef{
		External: in.GetServiceAccount().String(),
	}
	out.Labels = in.Labels
	out.Network = AllocationPolicy_NetworkPolicy_FromProto(mapCtx, in.GetNetwork())
	out.Placement = AllocationPolicy_PlacementPolicy_FromProto(mapCtx, in.GetPlacement())
	out.Tags = in.Tags
	return out
}
func AllocationPolicy_ToProto(mapCtx *direct.MapContext, in *krm.AllocationPolicy) *pb.AllocationPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy{}
	out.Location = AllocationPolicy_LocationPolicy_ToProto(mapCtx, in.Location)
	out.Instances = direct.Slice_ToProto(mapCtx, in.Instances, AllocationPolicy_InstancePolicyOrTemplate_ToProto)
	out.ServiceAccount = &pb.ServiceAccount{
		Email: in.ServiceAccountRef.External,
	}
	out.Labels = in.Labels
	out.Network = AllocationPolicy_NetworkPolicy_ToProto(mapCtx, in.Network)
	out.Placement = AllocationPolicy_PlacementPolicy_ToProto(mapCtx, in.Placement)
	out.Tags = in.Tags
	return out
}
func Environment_KMSEnvMap_FromProto(mapCtx *direct.MapContext, in *pb.Environment_KMSEnvMap) *krm.Environment_KMSEnvMap {
	if in == nil {
		return nil
	}
	out := &krm.Environment_KMSEnvMap{}
	out.KMSKeyRef = &v1beta1.KMSCryptoKeyRef{
		External: in.KeyName,
	}
	out.CipherText = direct.LazyPtr(in.GetCipherText())
	return out
}
func Environment_KMSEnvMap_ToProto(mapCtx *direct.MapContext, in *krm.Environment_KMSEnvMap) *pb.Environment_KMSEnvMap {
	if in == nil {
		return nil
	}
	out := &pb.Environment_KMSEnvMap{}
	if in.KMSKeyRef != nil {
		out.KeyName = in.KMSKeyRef.External
	}
	out.CipherText = direct.ValueOf(in.CipherText)
	return out
}
func ComputeResource_FromProto(mapCtx *direct.MapContext, in *pb.ComputeResource) *krm.ComputeResource {
	if in == nil {
		return nil
	}
	out := &krm.ComputeResource{}
	out.CPUMilli = direct.LazyPtr(in.GetCpuMilli())
	out.BootDiskMiB = direct.LazyPtr(in.GetBootDiskMib())
	out.MemoryMiB = direct.LazyPtr(in.GetMemoryMib())
	return out
}
func ComputeResource_ToProto(mapCtx *direct.MapContext, in *krm.ComputeResource) *pb.ComputeResource {
	if in == nil {
		return nil
	}
	out := &pb.ComputeResource{}
	out.CpuMilli = direct.ValueOf(in.CPUMilli)
	out.MemoryMib = direct.ValueOf(in.MemoryMiB)
	out.BootDiskMib = direct.ValueOf(in.BootDiskMiB)
	return out
}
func JobNotification_FromProto(mapCtx *direct.MapContext, in *pb.JobNotification) *krm.JobNotification {
	if in == nil {
		return nil
	}
	out := &krm.JobNotification{}
	if in.GetPubsubTopic() != "" {
		out.PubsubTopicRef = &v1beta1.PubSubTopicRef{External: in.GetPubsubTopic()}
	}
	out.Message = JobNotification_Message_FromProto(mapCtx, in.GetMessage())
	return out
}
func Runnable_Script_Path_ToProto(mapCtx *direct.MapContext, in *string) *pb.Runnable_Script_Path {
	if in == nil {
		return nil
	}
	out := &pb.Runnable_Script_Path{
		Path: direct.ValueOf(in),
	}
	return out
}
func Runnable_Script_Text_ToProto(mapCtx *direct.MapContext, in *string) *pb.Runnable_Script_Text {
	if in == nil {
		return nil
	}
	out := &pb.Runnable_Script_Text{
		Text: direct.ValueOf(in),
	}
	return out
}
func Volume_DeviceName_ToProto(mapCtx *direct.MapContext, in *string) *pb.Volume_DeviceName {
	if in == nil {
		return nil
	}
	out := &pb.Volume_DeviceName{
		DeviceName: direct.ValueOf(in),
	}
	return out
}
func AllocationPolicy_Disk_Image_ToProto(mapCtx *direct.MapContext, in *string) *pb.AllocationPolicy_Disk_Image {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_Disk_Image{
		Image: direct.ValueOf(in),
	}
	return out
}
func AllocationPolicy_Disk_Snapshot_ToProto(mapCtx *direct.MapContext, in *string) *pb.AllocationPolicy_Disk_Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_Disk_Snapshot{
		Snapshot: direct.ValueOf(in),
	}
	return out
}
func AllocationPolicy_AttachedDisk_ExistingDisk_ToProto(mapCtx *direct.MapContext, in *string) *pb.AllocationPolicy_AttachedDisk_ExistingDisk {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_AttachedDisk_ExistingDisk{
		ExistingDisk: direct.ValueOf(in),
	}
	return out
}
func AllocationPolicy_InstancePolicyOrTemplate_InstanceTemplate_ToProto(mapCtx *direct.MapContext, in *string) *pb.AllocationPolicy_InstancePolicyOrTemplate_Policy {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_InstancePolicyOrTemplate_Policy{
		Policy: &pb.AllocationPolicy_InstancePolicy{
			MachineType: direct.ValueOf(in),
		},
	}
	return out
}

func AllocationPolicy_Disk_ToProto(mapCtx *direct.MapContext, in *krm.AllocationPolicy_Disk) *pb.AllocationPolicy_Disk {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_Disk{}
	if in.ImageRef != nil {
		out.DataSource = &pb.AllocationPolicy_Disk_Image{
			Image: in.ImageRef.External,
		}
	}
	if oneof := AllocationPolicy_Disk_Snapshot_ToProto(mapCtx, in.Snapshot); oneof != nil {
		out.DataSource = oneof
	}
	out.Type = direct.ValueOf(in.Type)
	out.SizeGb = direct.ValueOf(in.SizeGB)
	out.DiskInterface = direct.ValueOf(in.DiskInterface)
	return out
}
func AllocationPolicy_Accelerator_FromProto(mapCtx *direct.MapContext, in *pb.AllocationPolicy_Accelerator) *krm.AllocationPolicy_Accelerator {
	if in == nil {
		return nil
	}
	out := &krm.AllocationPolicy_Accelerator{}
	out.Type = direct.LazyPtr(in.GetType())
	out.Count = direct.LazyPtr(in.GetCount())
	out.InstallGpuDrivers = direct.LazyPtr(in.GetInstallGpuDrivers())
	out.DriverVersion = direct.LazyPtr(in.GetDriverVersion())
	return out
}
func AllocationPolicy_Accelerator_ToProto(mapCtx *direct.MapContext, in *krm.AllocationPolicy_Accelerator) *pb.AllocationPolicy_Accelerator {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_Accelerator{}
	out.Type = direct.ValueOf(in.Type)
	out.Count = direct.ValueOf(in.Count)
	out.InstallGpuDrivers = direct.ValueOf(in.InstallGpuDrivers)
	out.DriverVersion = direct.ValueOf(in.DriverVersion)
	return out
}
func AllocationPolicy_AttachedDisk_FromProto(mapCtx *direct.MapContext, in *pb.AllocationPolicy_AttachedDisk) *krm.AllocationPolicy_AttachedDisk {
	if in == nil {
		return nil
	}
	out := &krm.AllocationPolicy_AttachedDisk{}
	out.NewDisk = AllocationPolicy_Disk_FromProto(mapCtx, in.GetNewDisk())
	out.ExistingDisk = direct.LazyPtr(in.GetExistingDisk())
	out.DeviceName = direct.LazyPtr(in.GetDeviceName())
	return out
}
func AllocationPolicy_AttachedDisk_ToProto(mapCtx *direct.MapContext, in *krm.AllocationPolicy_AttachedDisk) *pb.AllocationPolicy_AttachedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_AttachedDisk{}
	if oneof := AllocationPolicy_Disk_ToProto(mapCtx, in.NewDisk); oneof != nil {
		out.Attached = &pb.AllocationPolicy_AttachedDisk_NewDisk{NewDisk: oneof}
	}
	if oneof := AllocationPolicy_AttachedDisk_ExistingDisk_ToProto(mapCtx, in.ExistingDisk); oneof != nil {
		out.Attached = oneof
	}
	out.DeviceName = direct.ValueOf(in.DeviceName)
	return out
}
func AllocationPolicy_Disk_FromProto(mapCtx *direct.MapContext, in *pb.AllocationPolicy_Disk) *krm.AllocationPolicy_Disk {
	if in == nil {
		return nil
	}
	out := &krm.AllocationPolicy_Disk{}
	if in.GetImage() != "" {
		out.ImageRef = &v1alpha1.ResourceRef{External: in.GetImage()}
	}
	out.Snapshot = direct.LazyPtr(in.GetSnapshot())
	out.Type = direct.LazyPtr(in.GetType())
	out.SizeGB = direct.LazyPtr(in.GetSizeGb())
	out.DiskInterface = direct.LazyPtr(in.GetDiskInterface())
	return out
}
func AllocationPolicy_InstancePolicy_FromProto(mapCtx *direct.MapContext, in *pb.AllocationPolicy_InstancePolicy) *krm.AllocationPolicy_InstancePolicy {
	if in == nil {
		return nil
	}
	out := &krm.AllocationPolicy_InstancePolicy{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.MinCPUPlatform = direct.LazyPtr(in.GetMinCpuPlatform())
	out.ProvisioningModel = direct.Enum_FromProto(mapCtx, in.GetProvisioningModel())
	out.Accelerators = direct.Slice_FromProto(mapCtx, in.Accelerators, AllocationPolicy_Accelerator_FromProto)
	out.BootDisk = AllocationPolicy_Disk_FromProto(mapCtx, in.GetBootDisk())
	out.Disks = direct.Slice_FromProto(mapCtx, in.Disks, AllocationPolicy_AttachedDisk_FromProto)
	out.Reservation = direct.LazyPtr(in.GetReservation())
	return out
}
func AllocationPolicy_InstancePolicy_ToProto(mapCtx *direct.MapContext, in *krm.AllocationPolicy_InstancePolicy) *pb.AllocationPolicy_InstancePolicy {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_InstancePolicy{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.MinCpuPlatform = direct.ValueOf(in.MinCPUPlatform)
	out.ProvisioningModel = direct.Enum_ToProto[pb.AllocationPolicy_ProvisioningModel](mapCtx, in.ProvisioningModel)
	out.Accelerators = direct.Slice_ToProto(mapCtx, in.Accelerators, AllocationPolicy_Accelerator_ToProto)
	out.BootDisk = AllocationPolicy_Disk_ToProto(mapCtx, in.BootDisk)
	out.Disks = direct.Slice_ToProto(mapCtx, in.Disks, AllocationPolicy_AttachedDisk_ToProto)
	out.Reservation = direct.ValueOf(in.Reservation)
	return out
}
func AllocationPolicy_InstancePolicyOrTemplate_FromProto(mapCtx *direct.MapContext, in *pb.AllocationPolicy_InstancePolicyOrTemplate) *krm.AllocationPolicy_InstancePolicyOrTemplate {
	if in == nil {
		return nil
	}
	out := &krm.AllocationPolicy_InstancePolicyOrTemplate{}
	out.Policy = AllocationPolicy_InstancePolicy_FromProto(mapCtx, in.GetPolicy())
	out.InstanceTemplate = direct.LazyPtr(in.GetInstanceTemplate())
	out.InstallGpuDrivers = direct.LazyPtr(in.GetInstallGpuDrivers())
	out.InstallOpsAgent = direct.LazyPtr(in.GetInstallOpsAgent())
	out.BlockProjectSSHKeys = direct.LazyPtr(in.GetBlockProjectSshKeys())
	return out
}
func AllocationPolicy_InstancePolicyOrTemplate_ToProto(mapCtx *direct.MapContext, in *krm.AllocationPolicy_InstancePolicyOrTemplate) *pb.AllocationPolicy_InstancePolicyOrTemplate {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_InstancePolicyOrTemplate{}
	if oneof := AllocationPolicy_InstancePolicy_ToProto(mapCtx, in.Policy); oneof != nil {
		out.PolicyTemplate = &pb.AllocationPolicy_InstancePolicyOrTemplate_Policy{Policy: oneof}
	}
	if oneof := AllocationPolicy_InstancePolicyOrTemplate_InstanceTemplate_ToProto(mapCtx, in.InstanceTemplate); oneof != nil {
		out.PolicyTemplate = oneof
	}
	out.InstallGpuDrivers = direct.ValueOf(in.InstallGpuDrivers)
	out.InstallOpsAgent = direct.ValueOf(in.InstallOpsAgent)
	out.BlockProjectSshKeys = direct.ValueOf(in.BlockProjectSSHKeys)
	return out
}
func AllocationPolicy_LocationPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AllocationPolicy_LocationPolicy) *krm.AllocationPolicy_LocationPolicy {
	if in == nil {
		return nil
	}
	out := &krm.AllocationPolicy_LocationPolicy{}
	out.AllowedLocations = in.AllowedLocations
	return out
}
func AllocationPolicy_LocationPolicy_ToProto(mapCtx *direct.MapContext, in *krm.AllocationPolicy_LocationPolicy) *pb.AllocationPolicy_LocationPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_LocationPolicy{}
	out.AllowedLocations = in.AllowedLocations
	return out
}
func AllocationPolicy_NetworkInterface_FromProto(mapCtx *direct.MapContext, in *pb.AllocationPolicy_NetworkInterface) *krm.AllocationPolicy_NetworkInterface {
	if in == nil {
		return nil
	}
	out := &krm.AllocationPolicy_NetworkInterface{}
	if in.GetNetwork() != "" {
		out.NetworkRef = &v1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &v1beta1.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	out.NoExternalIPAddress = direct.LazyPtr(in.GetNoExternalIpAddress())
	return out
}
func AllocationPolicy_NetworkInterface_ToProto(mapCtx *direct.MapContext, in *krm.AllocationPolicy_NetworkInterface) *pb.AllocationPolicy_NetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_NetworkInterface{}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	out.NoExternalIpAddress = direct.ValueOf(in.NoExternalIPAddress)
	return out
}
func AllocationPolicy_NetworkPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AllocationPolicy_NetworkPolicy) *krm.AllocationPolicy_NetworkPolicy {
	if in == nil {
		return nil
	}
	out := &krm.AllocationPolicy_NetworkPolicy{}
	out.NetworkInterfaces = direct.Slice_FromProto(mapCtx, in.NetworkInterfaces, AllocationPolicy_NetworkInterface_FromProto)
	return out
}
func AllocationPolicy_NetworkPolicy_ToProto(mapCtx *direct.MapContext, in *krm.AllocationPolicy_NetworkPolicy) *pb.AllocationPolicy_NetworkPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_NetworkPolicy{}
	out.NetworkInterfaces = direct.Slice_ToProto(mapCtx, in.NetworkInterfaces, AllocationPolicy_NetworkInterface_ToProto)
	return out
}
func AllocationPolicy_PlacementPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AllocationPolicy_PlacementPolicy) *krm.AllocationPolicy_PlacementPolicy {
	if in == nil {
		return nil
	}
	out := &krm.AllocationPolicy_PlacementPolicy{}
	out.Collocation = direct.LazyPtr(in.GetCollocation())
	out.MaxDistance = direct.LazyPtr(in.GetMaxDistance())
	return out
}
func AllocationPolicy_PlacementPolicy_ToProto(mapCtx *direct.MapContext, in *krm.AllocationPolicy_PlacementPolicy) *pb.AllocationPolicy_PlacementPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_PlacementPolicy{}
	out.Collocation = direct.ValueOf(in.Collocation)
	out.MaxDistance = direct.ValueOf(in.MaxDistance)
	return out
}
func Environment_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.Environment {
	if in == nil {
		return nil
	}
	out := &krm.Environment{}
	out.Variables = in.Variables
	out.SecretVariables = in.SecretVariables
	out.EncryptedVariables = Environment_KMSEnvMap_FromProto(mapCtx, in.GetEncryptedVariables())
	return out
}
func Environment_ToProto(mapCtx *direct.MapContext, in *krm.Environment) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	out.Variables = in.Variables
	out.SecretVariables = in.SecretVariables
	out.EncryptedVariables = Environment_KMSEnvMap_ToProto(mapCtx, in.EncryptedVariables)
	return out
}
func GCS_FromProto(mapCtx *direct.MapContext, in *pb.GCS) *krm.GCS {
	if in == nil {
		return nil
	}
	out := &krm.GCS{}
	out.RemotePath = direct.LazyPtr(in.GetRemotePath())
	return out
}
func GCS_ToProto(mapCtx *direct.MapContext, in *krm.GCS) *pb.GCS {
	if in == nil {
		return nil
	}
	out := &pb.GCS{}
	out.RemotePath = direct.ValueOf(in.RemotePath)
	return out
}
func JobNotification_ToProto(mapCtx *direct.MapContext, in *krm.JobNotification) *pb.JobNotification {
	if in == nil {
		return nil
	}
	out := &pb.JobNotification{}
	if in.PubsubTopicRef != nil {
		out.PubsubTopic = in.PubsubTopicRef.External
	}
	out.Message = JobNotification_Message_ToProto(mapCtx, in.Message)
	return out
}
func JobNotification_Message_FromProto(mapCtx *direct.MapContext, in *pb.JobNotification_Message) *krm.JobNotification_Message {
	if in == nil {
		return nil
	}
	out := &krm.JobNotification_Message{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.NewJobState = direct.Enum_FromProto(mapCtx, in.GetNewJobState())
	out.NewTaskState = direct.Enum_FromProto(mapCtx, in.GetNewTaskState())
	return out
}
func JobNotification_Message_ToProto(mapCtx *direct.MapContext, in *krm.JobNotification_Message) *pb.JobNotification_Message {
	if in == nil {
		return nil
	}
	out := &pb.JobNotification_Message{}
	out.Type = direct.Enum_ToProto[pb.JobNotification_Type](mapCtx, in.Type)
	out.NewJobState = direct.Enum_ToProto[pb.JobStatus_State](mapCtx, in.NewJobState)
	out.NewTaskState = direct.Enum_ToProto[pb.TaskStatus_State](mapCtx, in.NewTaskState)
	return out
}
func JobStatus_FromProto(mapCtx *direct.MapContext, in *pb.JobStatus) *krm.JobStatus {
	if in == nil {
		return nil
	}
	out := &krm.JobStatus{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StatusEvents = direct.Slice_FromProto(mapCtx, in.StatusEvents, StatusEvent_FromProto)
	// MISSING: TaskGroups
	out.RunDuration = direct.StringDuration_FromProto(mapCtx, in.GetRunDuration())
	return out
}
func JobStatus_ToProto(mapCtx *direct.MapContext, in *krm.JobStatus) *pb.JobStatus {
	if in == nil {
		return nil
	}
	out := &pb.JobStatus{}
	out.State = direct.Enum_ToProto[pb.JobStatus_State](mapCtx, in.State)
	out.StatusEvents = direct.Slice_ToProto(mapCtx, in.StatusEvents, StatusEvent_ToProto)
	// MISSING: TaskGroups
	out.RunDuration = direct.StringDuration_ToProto(mapCtx, in.RunDuration)
	return out
}
func JobStatus_InstanceStatus_FromProto(mapCtx *direct.MapContext, in *pb.JobStatus_InstanceStatus) *krm.JobStatus_InstanceStatus {
	if in == nil {
		return nil
	}
	out := &krm.JobStatus_InstanceStatus{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.ProvisioningModel = direct.Enum_FromProto(mapCtx, in.GetProvisioningModel())
	out.TaskPack = direct.LazyPtr(in.GetTaskPack())
	out.BootDisk = AllocationPolicy_Disk_FromProto(mapCtx, in.GetBootDisk())
	return out
}
func JobStatus_InstanceStatus_ToProto(mapCtx *direct.MapContext, in *krm.JobStatus_InstanceStatus) *pb.JobStatus_InstanceStatus {
	if in == nil {
		return nil
	}
	out := &pb.JobStatus_InstanceStatus{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.ProvisioningModel = direct.Enum_ToProto[pb.AllocationPolicy_ProvisioningModel](mapCtx, in.ProvisioningModel)
	out.TaskPack = direct.ValueOf(in.TaskPack)
	out.BootDisk = AllocationPolicy_Disk_ToProto(mapCtx, in.BootDisk)
	return out
}
func JobStatus_TaskGroupStatus_FromProto(mapCtx *direct.MapContext, in *pb.JobStatus_TaskGroupStatus) *krm.JobStatus_TaskGroupStatus {
	if in == nil {
		return nil
	}
	out := &krm.JobStatus_TaskGroupStatus{}
	out.Counts = in.Counts
	out.Instances = direct.Slice_FromProto(mapCtx, in.Instances, JobStatus_InstanceStatus_FromProto)
	return out
}
func JobStatus_TaskGroupStatus_ToProto(mapCtx *direct.MapContext, in *krm.JobStatus_TaskGroupStatus) *pb.JobStatus_TaskGroupStatus {
	if in == nil {
		return nil
	}
	out := &pb.JobStatus_TaskGroupStatus{}
	out.Counts = in.Counts
	out.Instances = direct.Slice_ToProto(mapCtx, in.Instances, JobStatus_InstanceStatus_ToProto)
	return out
}
func LifecyclePolicy_FromProto(mapCtx *direct.MapContext, in *pb.LifecyclePolicy) *krm.LifecyclePolicy {
	if in == nil {
		return nil
	}
	out := &krm.LifecyclePolicy{}
	out.Action = direct.Enum_FromProto(mapCtx, in.GetAction())
	out.ActionCondition = LifecyclePolicy_ActionCondition_FromProto(mapCtx, in.GetActionCondition())
	return out
}
func LifecyclePolicy_ToProto(mapCtx *direct.MapContext, in *krm.LifecyclePolicy) *pb.LifecyclePolicy {
	if in == nil {
		return nil
	}
	out := &pb.LifecyclePolicy{}
	out.Action = direct.Enum_ToProto[pb.LifecyclePolicy_Action](mapCtx, in.Action)
	out.ActionCondition = LifecyclePolicy_ActionCondition_ToProto(mapCtx, in.ActionCondition)
	return out
}
func LifecyclePolicy_ActionCondition_FromProto(mapCtx *direct.MapContext, in *pb.LifecyclePolicy_ActionCondition) *krm.LifecyclePolicy_ActionCondition {
	if in == nil {
		return nil
	}
	out := &krm.LifecyclePolicy_ActionCondition{}
	out.ExitCodes = in.ExitCodes
	return out
}
func LifecyclePolicy_ActionCondition_ToProto(mapCtx *direct.MapContext, in *krm.LifecyclePolicy_ActionCondition) *pb.LifecyclePolicy_ActionCondition {
	if in == nil {
		return nil
	}
	out := &pb.LifecyclePolicy_ActionCondition{}
	out.ExitCodes = in.ExitCodes
	return out
}
func LogsPolicy_FromProto(mapCtx *direct.MapContext, in *pb.LogsPolicy) *krm.LogsPolicy {
	if in == nil {
		return nil
	}
	out := &krm.LogsPolicy{}
	out.Destination = direct.Enum_FromProto(mapCtx, in.GetDestination())
	out.LogsPath = direct.LazyPtr(in.GetLogsPath())
	out.CloudLoggingOption = LogsPolicy_CloudLoggingOption_FromProto(mapCtx, in.GetCloudLoggingOption())
	return out
}
func LogsPolicy_ToProto(mapCtx *direct.MapContext, in *krm.LogsPolicy) *pb.LogsPolicy {
	if in == nil {
		return nil
	}
	out := &pb.LogsPolicy{}
	out.Destination = direct.Enum_ToProto[pb.LogsPolicy_Destination](mapCtx, in.Destination)
	out.LogsPath = direct.ValueOf(in.LogsPath)
	out.CloudLoggingOption = LogsPolicy_CloudLoggingOption_ToProto(mapCtx, in.CloudLoggingOption)
	return out
}
func LogsPolicy_CloudLoggingOption_FromProto(mapCtx *direct.MapContext, in *pb.LogsPolicy_CloudLoggingOption) *krm.LogsPolicy_CloudLoggingOption {
	if in == nil {
		return nil
	}
	out := &krm.LogsPolicy_CloudLoggingOption{}
	out.UseGenericTaskMonitoredResource = direct.LazyPtr(in.GetUseGenericTaskMonitoredResource())
	return out
}
func LogsPolicy_CloudLoggingOption_ToProto(mapCtx *direct.MapContext, in *krm.LogsPolicy_CloudLoggingOption) *pb.LogsPolicy_CloudLoggingOption {
	if in == nil {
		return nil
	}
	out := &pb.LogsPolicy_CloudLoggingOption{}
	out.UseGenericTaskMonitoredResource = direct.ValueOf(in.UseGenericTaskMonitoredResource)
	return out
}
func Nfs_FromProto(mapCtx *direct.MapContext, in *pb.NFS) *krm.Nfs {
	if in == nil {
		return nil
	}
	out := &krm.Nfs{}
	out.Server = direct.LazyPtr(in.GetServer())
	out.RemotePath = direct.LazyPtr(in.GetRemotePath())
	return out
}
func Nfs_ToProto(mapCtx *direct.MapContext, in *krm.Nfs) *pb.NFS {
	if in == nil {
		return nil
	}
	out := &pb.NFS{}
	out.Server = direct.ValueOf(in.Server)
	out.RemotePath = direct.ValueOf(in.RemotePath)
	return out
}
func Runnable_FromProto(mapCtx *direct.MapContext, in *pb.Runnable) *krm.Runnable {
	if in == nil {
		return nil
	}
	out := &krm.Runnable{}
	out.Container = Runnable_Container_FromProto(mapCtx, in.GetContainer())
	out.Script = Runnable_Script_FromProto(mapCtx, in.GetScript())
	out.Barrier = Runnable_Barrier_FromProto(mapCtx, in.GetBarrier())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.IgnoreExitStatus = direct.LazyPtr(in.GetIgnoreExitStatus())
	out.Background = direct.LazyPtr(in.GetBackground())
	out.AlwaysRun = direct.LazyPtr(in.GetAlwaysRun())
	out.Environment = Environment_FromProto(mapCtx, in.GetEnvironment())
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.Labels = in.Labels
	return out
}
func Runnable_ToProto(mapCtx *direct.MapContext, in *krm.Runnable) *pb.Runnable {
	if in == nil {
		return nil
	}
	out := &pb.Runnable{}
	if oneof := Runnable_Container_ToProto(mapCtx, in.Container); oneof != nil {
		out.Executable = &pb.Runnable_Container_{Container: oneof}
	}
	if oneof := Runnable_Script_ToProto(mapCtx, in.Script); oneof != nil {
		out.Executable = &pb.Runnable_Script_{Script: oneof}
	}
	if oneof := Runnable_Barrier_ToProto(mapCtx, in.Barrier); oneof != nil {
		out.Executable = &pb.Runnable_Barrier_{Barrier: oneof}
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.IgnoreExitStatus = direct.ValueOf(in.IgnoreExitStatus)
	out.Background = direct.ValueOf(in.Background)
	out.AlwaysRun = direct.ValueOf(in.AlwaysRun)
	out.Environment = Environment_ToProto(mapCtx, in.Environment)
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.Labels = in.Labels
	return out
}
func Runnable_Barrier_FromProto(mapCtx *direct.MapContext, in *pb.Runnable_Barrier) *krm.Runnable_Barrier {
	if in == nil {
		return nil
	}
	out := &krm.Runnable_Barrier{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Runnable_Barrier_ToProto(mapCtx *direct.MapContext, in *krm.Runnable_Barrier) *pb.Runnable_Barrier {
	if in == nil {
		return nil
	}
	out := &pb.Runnable_Barrier{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func Runnable_Container_FromProto(mapCtx *direct.MapContext, in *pb.Runnable_Container) *krm.Runnable_Container {
	if in == nil {
		return nil
	}
	out := &krm.Runnable_Container{}
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.Commands = in.Commands
	out.Entrypoint = direct.LazyPtr(in.GetEntrypoint())
	out.Volumes = in.Volumes
	out.Options = direct.LazyPtr(in.GetOptions())
	out.BlockExternalNetwork = direct.LazyPtr(in.GetBlockExternalNetwork())
	out.Username = direct.LazyPtr(in.GetUsername())
	// MISSING: Password
	out.EnableImageStreaming = direct.LazyPtr(in.GetEnableImageStreaming())
	return out
}
func Runnable_Container_ToProto(mapCtx *direct.MapContext, in *krm.Runnable_Container) *pb.Runnable_Container {
	if in == nil {
		return nil
	}
	out := &pb.Runnable_Container{}
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.Commands = in.Commands
	out.Entrypoint = direct.ValueOf(in.Entrypoint)
	out.Volumes = in.Volumes
	out.Options = direct.ValueOf(in.Options)
	out.BlockExternalNetwork = direct.ValueOf(in.BlockExternalNetwork)
	out.Username = direct.ValueOf(in.Username)
	// MISSING: Password
	out.EnableImageStreaming = direct.ValueOf(in.EnableImageStreaming)
	return out
}
func Runnable_Script_FromProto(mapCtx *direct.MapContext, in *pb.Runnable_Script) *krm.Runnable_Script {
	if in == nil {
		return nil
	}
	out := &krm.Runnable_Script{}
	out.Path = direct.LazyPtr(in.GetPath())
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func Runnable_Script_ToProto(mapCtx *direct.MapContext, in *krm.Runnable_Script) *pb.Runnable_Script {
	if in == nil {
		return nil
	}
	out := &pb.Runnable_Script{}
	if oneof := Runnable_Script_Path_ToProto(mapCtx, in.Path); oneof != nil {
		out.Command = oneof
	}
	if oneof := Runnable_Script_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Command = oneof
	}
	return out
}
func ServiceAccount_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccount) *krm.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &krm.ServiceAccount{}
	out.Email = direct.LazyPtr(in.GetEmail())
	out.Scopes = in.Scopes
	return out
}
func ServiceAccount_ToProto(mapCtx *direct.MapContext, in *krm.ServiceAccount) *pb.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccount{}
	out.Email = direct.ValueOf(in.Email)
	out.Scopes = in.Scopes
	return out
}
func StatusEvent_FromProto(mapCtx *direct.MapContext, in *pb.StatusEvent) *krm.StatusEvent {
	if in == nil {
		return nil
	}
	out := &krm.StatusEvent{}
	out.Type = direct.LazyPtr(in.GetType())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.EventTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEventTime())
	out.TaskExecution = TaskExecution_FromProto(mapCtx, in.GetTaskExecution())
	out.TaskState = direct.Enum_FromProto(mapCtx, in.GetTaskState())
	return out
}
func StatusEvent_ToProto(mapCtx *direct.MapContext, in *krm.StatusEvent) *pb.StatusEvent {
	if in == nil {
		return nil
	}
	out := &pb.StatusEvent{}
	out.Type = direct.ValueOf(in.Type)
	out.Description = direct.ValueOf(in.Description)
	out.EventTime = direct.StringTimestamp_ToProto(mapCtx, in.EventTime)
	out.TaskExecution = TaskExecution_ToProto(mapCtx, in.TaskExecution)
	out.TaskState = direct.Enum_ToProto[pb.TaskStatus_State](mapCtx, in.TaskState)
	return out
}
func TaskExecution_FromProto(mapCtx *direct.MapContext, in *pb.TaskExecution) *krm.TaskExecution {
	if in == nil {
		return nil
	}
	out := &krm.TaskExecution{}
	out.ExitCode = direct.LazyPtr(in.GetExitCode())
	return out
}
func TaskExecution_ToProto(mapCtx *direct.MapContext, in *krm.TaskExecution) *pb.TaskExecution {
	if in == nil {
		return nil
	}
	out := &pb.TaskExecution{}
	out.ExitCode = direct.ValueOf(in.ExitCode)
	return out
}
func TaskGroup_FromProto(mapCtx *direct.MapContext, in *pb.TaskGroup) *krm.TaskGroup {
	if in == nil {
		return nil
	}
	out := &krm.TaskGroup{}
	// MISSING: Name
	out.TaskSpec = TaskSpec_FromProto(mapCtx, in.GetTaskSpec())
	out.TaskCount = direct.LazyPtr(in.GetTaskCount())
	out.Parallelism = direct.LazyPtr(in.GetParallelism())
	out.SchedulingPolicy = direct.Enum_FromProto(mapCtx, in.GetSchedulingPolicy())
	out.TaskEnvironments = direct.Slice_FromProto(mapCtx, in.TaskEnvironments, Environment_FromProto)
	out.TaskCountPerNode = direct.LazyPtr(in.GetTaskCountPerNode())
	out.RequireHostsFile = direct.LazyPtr(in.GetRequireHostsFile())
	out.PermissiveSSH = direct.LazyPtr(in.GetPermissiveSsh())
	out.RunAsNonRoot = direct.LazyPtr(in.GetRunAsNonRoot())
	return out
}
func TaskGroup_ToProto(mapCtx *direct.MapContext, in *krm.TaskGroup) *pb.TaskGroup {
	if in == nil {
		return nil
	}
	out := &pb.TaskGroup{}
	// MISSING: Name
	out.TaskSpec = TaskSpec_ToProto(mapCtx, in.TaskSpec)
	out.TaskCount = direct.ValueOf(in.TaskCount)
	out.Parallelism = direct.ValueOf(in.Parallelism)
	out.SchedulingPolicy = direct.Enum_ToProto[pb.TaskGroup_SchedulingPolicy](mapCtx, in.SchedulingPolicy)
	out.TaskEnvironments = direct.Slice_ToProto(mapCtx, in.TaskEnvironments, Environment_ToProto)
	out.TaskCountPerNode = direct.ValueOf(in.TaskCountPerNode)
	out.RequireHostsFile = direct.ValueOf(in.RequireHostsFile)
	out.PermissiveSsh = direct.ValueOf(in.PermissiveSSH)
	out.RunAsNonRoot = direct.ValueOf(in.RunAsNonRoot)
	return out
}
func TaskGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TaskGroup) *krm.TaskGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TaskGroupObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: TaskSpec
	// MISSING: TaskCount
	// MISSING: Parallelism
	// MISSING: SchedulingPolicy
	// MISSING: TaskEnvironments
	// MISSING: TaskCountPerNode
	// MISSING: RequireHostsFile
	// MISSING: PermissiveSSH
	// MISSING: RunAsNonRoot
	return out
}
func TaskGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TaskGroupObservedState) *pb.TaskGroup {
	if in == nil {
		return nil
	}
	out := &pb.TaskGroup{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: TaskSpec
	// MISSING: TaskCount
	// MISSING: Parallelism
	// MISSING: SchedulingPolicy
	// MISSING: TaskEnvironments
	// MISSING: TaskCountPerNode
	// MISSING: RequireHostsFile
	// MISSING: PermissiveSSH
	// MISSING: RunAsNonRoot
	return out
}
func TaskSpec_FromProto(mapCtx *direct.MapContext, in *pb.TaskSpec) *krm.TaskSpec {
	if in == nil {
		return nil
	}
	out := &krm.TaskSpec{}
	out.Runnables = direct.Slice_FromProto(mapCtx, in.Runnables, Runnable_FromProto)
	out.ComputeResource = ComputeResource_FromProto(mapCtx, in.GetComputeResource())
	out.MaxRunDuration = direct.StringDuration_FromProto(mapCtx, in.GetMaxRunDuration())
	out.MaxRetryCount = direct.LazyPtr(in.GetMaxRetryCount())
	out.LifecyclePolicies = direct.Slice_FromProto(mapCtx, in.LifecyclePolicies, LifecyclePolicy_FromProto)
	out.Environments = in.Environments
	out.Volumes = direct.Slice_FromProto(mapCtx, in.Volumes, Volume_FromProto)
	out.Environment = Environment_FromProto(mapCtx, in.GetEnvironment())
	return out
}
func TaskSpec_ToProto(mapCtx *direct.MapContext, in *krm.TaskSpec) *pb.TaskSpec {
	if in == nil {
		return nil
	}
	out := &pb.TaskSpec{}
	out.Runnables = direct.Slice_ToProto(mapCtx, in.Runnables, Runnable_ToProto)
	out.ComputeResource = ComputeResource_ToProto(mapCtx, in.ComputeResource)
	out.MaxRunDuration = direct.StringDuration_ToProto(mapCtx, in.MaxRunDuration)
	out.MaxRetryCount = direct.ValueOf(in.MaxRetryCount)
	out.LifecyclePolicies = direct.Slice_ToProto(mapCtx, in.LifecyclePolicies, LifecyclePolicy_ToProto)
	out.Environments = in.Environments
	out.Volumes = direct.Slice_ToProto(mapCtx, in.Volumes, Volume_ToProto)
	out.Environment = Environment_ToProto(mapCtx, in.Environment)
	return out
}
func Volume_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.Volume {
	if in == nil {
		return nil
	}
	out := &krm.Volume{}
	out.Nfs = Nfs_FromProto(mapCtx, in.GetNfs())
	out.GCS = GCS_FromProto(mapCtx, in.GetGcs())
	out.DeviceName = direct.LazyPtr(in.GetDeviceName())
	out.MountPath = direct.LazyPtr(in.GetMountPath())
	out.MountOptions = in.MountOptions
	return out
}
func Volume_ToProto(mapCtx *direct.MapContext, in *krm.Volume) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	if oneof := Nfs_ToProto(mapCtx, in.Nfs); oneof != nil {
		out.Source = &pb.Volume_Nfs{Nfs: oneof}
	}
	if oneof := GCS_ToProto(mapCtx, in.GCS); oneof != nil {
		out.Source = &pb.Volume_Gcs{Gcs: oneof}
	}
	if oneof := Volume_DeviceName_ToProto(mapCtx, in.DeviceName); oneof != nil {
		out.Source = oneof
	}
	out.MountPath = direct.ValueOf(in.MountPath)
	out.MountOptions = in.MountOptions
	return out
}
