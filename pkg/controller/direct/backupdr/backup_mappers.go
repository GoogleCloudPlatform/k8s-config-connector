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

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AccessConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AccessConfig) *krm.AccessConfig {
	if in == nil {
		return nil
	}
	out := &krm.AccessConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Name = in.Name
	out.ExternalIP = in.ExternalIp
	out.ExternalIPV6 = in.ExternalIpv6
	out.ExternalIPV6PrefixLength = in.ExternalIpv6PrefixLength
	out.SetPublicPtr = in.SetPublicPtr
	out.PublicPtrDomainName = in.PublicPtrDomainName
	out.NetworkTier = direct.Enum_FromProto(mapCtx, in.GetNetworkTier())
	return out
}

func AccessConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AccessConfig) *pb.AccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.AccessConfig{}
	if in.Type != nil {
		v := direct.Enum_ToProto[pb.AccessConfig_AccessType](mapCtx, in.Type)
		out.Type = &v
	}
	out.Name = in.Name
	out.ExternalIp = in.ExternalIP
	out.ExternalIpv6 = in.ExternalIPV6
	out.ExternalIpv6PrefixLength = in.ExternalIPV6PrefixLength
	out.SetPublicPtr = in.SetPublicPtr
	out.PublicPtrDomainName = in.PublicPtrDomainName
	if in.NetworkTier != nil {
		v := direct.Enum_ToProto[pb.AccessConfig_NetworkTier](mapCtx, in.NetworkTier)
		out.NetworkTier = &v
	}
	return out
}

func AttachedDiskObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AttachedDiskObservedState) *pb.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AttachedDisk{}
	out.InitializeParams = AttachedDisk_InitializeParams_v1alpha1_ToProto(mapCtx, in.InitializeParams)
	out.DeviceName = in.DeviceName
	out.Kind = in.Kind
	if in.DiskTypeDeprecated != nil {
		v := direct.Enum_ToProto[pb.AttachedDisk_DiskType](mapCtx, in.DiskTypeDeprecated)
		out.DiskTypeDeprecated = &v
	}
	if in.Mode != nil {
		v := direct.Enum_ToProto[pb.AttachedDisk_DiskMode](mapCtx, in.Mode)
		out.Mode = &v
	}
	out.Source = in.Source
	out.Index = in.Index
	out.Boot = in.Boot
	out.AutoDelete = in.AutoDelete
	out.License = in.License
	if in.DiskInterface != nil {
		v := direct.Enum_ToProto[pb.AttachedDisk_DiskInterface](mapCtx, in.DiskInterface)
		out.DiskInterface = &v
	}
	out.GuestOsFeature = direct.Slice_ToProto(mapCtx, in.GuestOSFeature, GuestOSFeature_v1alpha1_ToProto)
	out.DiskEncryptionKey = CustomerEncryptionKey_v1alpha1_ToProto(mapCtx, in.DiskEncryptionKey)
	out.DiskSizeGb = in.DiskSizeGB
	if in.SavedState != nil {
		v := direct.Enum_ToProto[pb.AttachedDisk_DiskSavedState](mapCtx, in.SavedState)
		out.SavedState = &v
	}
	out.DiskType = in.DiskType
	if in.Type != nil {
		v := direct.Enum_ToProto[pb.AttachedDisk_DiskType](mapCtx, in.Type)
		out.Type = &v
	}
	return out
}

func ComputeInstanceBackupPropertiesObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceBackupPropertiesObservedState) *pb.ComputeInstanceBackupProperties {
	if in == nil {
		return nil
	}
	out := &pb.ComputeInstanceBackupProperties{}
	out.Description = in.Description
	out.Tags = Tags_v1alpha1_ToProto(mapCtx, in.Tags)
	out.MachineType = in.MachineType
	out.CanIpForward = in.CanIPForward
	out.NetworkInterface = direct.Slice_ToProto(mapCtx, in.NetworkInterface, NetworkInterfaceObservedState_v1alpha1_ToProto)
	out.Disk = direct.Slice_ToProto(mapCtx, in.Disk, AttachedDiskObservedState_v1alpha1_ToProto)
	out.Metadata = Metadata_v1alpha1_ToProto(mapCtx, in.Metadata)
	out.ServiceAccount = direct.Slice_ToProto(mapCtx, in.ServiceAccount, ServiceAccount_v1alpha1_ToProto)
	out.Scheduling = Scheduling_v1alpha1_ToProto(mapCtx, in.Scheduling)
	out.GuestAccelerator = direct.Slice_ToProto(mapCtx, in.GuestAccelerator, AcceleratorConfig_v1alpha1_ToProto)
	out.MinCpuPlatform = in.MinCPUPlatform
	if in.KeyRevocationActionType != nil {
		v := direct.Enum_ToProto[pb.KeyRevocationActionType](mapCtx, in.KeyRevocationActionType)
		out.KeyRevocationActionType = &v
	}
	out.SourceInstance = in.SourceInstance
	out.Labels = in.Labels
	return out
}

func BackupDRBackupObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BackupDRBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Description = in.Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ConsistencyTime = direct.StringTimestamp_ToProto(mapCtx, in.ConsistencyTime)
	out.Etag = in.Etag
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.ServiceLocks = direct.Slice_ToProto(mapCtx, in.ServiceLocks, BackupLockObservedState_v1alpha1_ToProto)
	if in.ComputeInstanceBackupProperties != nil {
		out.BackupProperties = &pb.Backup_ComputeInstanceBackupProperties{
			ComputeInstanceBackupProperties: ComputeInstanceBackupPropertiesObservedState_v1alpha1_ToProto(mapCtx, in.ComputeInstanceBackupProperties),
		}
	}
	if in.CloudSQLInstanceBackupProperties != nil {
		out.BackupProperties = &pb.Backup_CloudSqlInstanceBackupProperties{
			CloudSqlInstanceBackupProperties: CloudSQLInstanceBackupPropertiesObservedState_v1alpha1_ToProto(mapCtx, in.CloudSQLInstanceBackupProperties),
		}
	}
	if in.BackupApplianceBackupProperties != nil {
		out.BackupProperties = &pb.Backup_BackupApplianceBackupProperties{
			BackupApplianceBackupProperties: BackupApplianceBackupPropertiesObservedState_v1alpha1_ToProto(mapCtx, in.BackupApplianceBackupProperties),
		}
	}
	if in.DiskBackupProperties != nil {
		out.BackupProperties = &pb.Backup_DiskBackupProperties{
			DiskBackupProperties: DiskBackupPropertiesObservedState_v1alpha1_ToProto(mapCtx, in.DiskBackupProperties),
		}
	}
	out.BackupType = direct.Enum_ToProto[pb.Backup_BackupType](mapCtx, in.BackupType)
	if in.GcpBackupPlanInfo != nil {
		out.PlanInfo = &pb.Backup_GcpBackupPlanInfo{
			GcpBackupPlanInfo: Backup_GcpBackupPlanInfoObservedState_v1alpha1_ToProto(mapCtx, in.GcpBackupPlanInfo),
		}
	}
	out.ResourceSizeBytes = direct.ValueOf(in.ResourceSizeBytes)
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}

func Metadata_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Metadata) *krm.Metadata {
	if in == nil {
		return nil
	}
	out := &krm.Metadata{}
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, Entry_v1alpha1_FromProto)
	return out
}

func Metadata_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Metadata) *pb.Metadata {
	if in == nil {
		return nil
	}
	out := &pb.Metadata{}
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, Entry_v1alpha1_ToProto)
	return out
}

func DiskBackupPropertiesObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.DiskBackupPropertiesObservedState) *pb.DiskBackupProperties {
	if in == nil {
		return nil
	}
	out := &pb.DiskBackupProperties{}
	out.Description = in.Description
	out.Licenses = in.Licenses
	out.GuestOsFeature = direct.Slice_ToProto(mapCtx, in.GuestOSFeature, GuestOSFeature_v1alpha1_ToProto)
	if in.Architecture != nil {
		v := direct.Enum_ToProto[pb.DiskBackupProperties_Architecture](mapCtx, in.Architecture)
		out.Architecture = &v
	}
	out.Type = in.Type
	out.SizeGb = in.SizeGB
	out.Region = in.Region
	out.Zone = in.Zone
	out.ReplicaZones = in.ReplicaZones
	out.SourceDisk = in.SourceDisk
	return out
}

func GuestOSFeature_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.GuestOSFeature) *pb.GuestOsFeature {
	if in == nil {
		return nil
	}
	out := &pb.GuestOsFeature{}
	if in.Type != nil {
		v := direct.Enum_ToProto[pb.GuestOsFeature_FeatureType](mapCtx, in.Type)
		out.Type = &v
	}
	return out
}

func Scheduling_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Scheduling) *krm.Scheduling {
	if in == nil {
		return nil
	}
	out := &krm.Scheduling{}
	out.OnHostMaintenance = direct.Enum_FromProto(mapCtx, in.GetOnHostMaintenance())
	out.AutomaticRestart = in.AutomaticRestart
	out.Preemptible = in.Preemptible
	out.NodeAffinities = direct.Slice_FromProto(mapCtx, in.NodeAffinities, Scheduling_NodeAffinity_v1alpha1_FromProto)
	out.MinNodeCPUs = in.MinNodeCpus
	out.ProvisioningModel = direct.Enum_FromProto(mapCtx, in.GetProvisioningModel())
	out.InstanceTerminationAction = direct.Enum_FromProto(mapCtx, in.GetInstanceTerminationAction())
	out.LocalSSDRecoveryTimeout = SchedulingDuration_v1alpha1_FromProto(mapCtx, in.GetLocalSsdRecoveryTimeout())
	return out
}

func Scheduling_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Scheduling) *pb.Scheduling {
	if in == nil {
		return nil
	}
	out := &pb.Scheduling{}
	if in.OnHostMaintenance != nil {
		v := direct.Enum_ToProto[pb.Scheduling_OnHostMaintenance](mapCtx, in.OnHostMaintenance)
		out.OnHostMaintenance = &v
	}
	out.AutomaticRestart = in.AutomaticRestart
	out.Preemptible = in.Preemptible
	out.NodeAffinities = direct.Slice_ToProto(mapCtx, in.NodeAffinities, Scheduling_NodeAffinity_v1alpha1_ToProto)
	out.MinNodeCpus = in.MinNodeCPUs
	if in.ProvisioningModel != nil {
		v := direct.Enum_ToProto[pb.Scheduling_ProvisioningModel](mapCtx, in.ProvisioningModel)
		out.ProvisioningModel = &v
	}
	if in.InstanceTerminationAction != nil {
		v := direct.Enum_ToProto[pb.Scheduling_InstanceTerminationAction](mapCtx, in.InstanceTerminationAction)
		out.InstanceTerminationAction = &v
	}
	out.LocalSsdRecoveryTimeout = SchedulingDuration_v1alpha1_ToProto(mapCtx, in.LocalSSDRecoveryTimeout)
	return out
}

func Scheduling_NodeAffinity_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Scheduling_NodeAffinity) *pb.Scheduling_NodeAffinity {
	if in == nil {
		return nil
	}
	out := &pb.Scheduling_NodeAffinity{}
	out.Key = in.Key
	if in.Operator != nil {
		v := direct.Enum_ToProto[pb.Scheduling_NodeAffinity_Operator](mapCtx, in.Operator)
		out.Operator = &v
	}
	out.Values = in.Values
	return out
}

func NetworkInterfaceObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkInterface) *krm.NetworkInterfaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkInterfaceObservedState{}
	out.Network = in.Network
	out.Subnetwork = in.Subnetwork
	out.IPAddress = in.IpAddress
	out.IPV6Address = in.Ipv6Address
	out.InternalIPV6PrefixLength = in.InternalIpv6PrefixLength
	out.Name = in.Name
	out.AccessConfigs = direct.Slice_FromProto(mapCtx, in.AccessConfigs, AccessConfig_v1alpha1_FromProto)
	out.IPV6AccessConfigs = direct.Slice_FromProto(mapCtx, in.Ipv6AccessConfigs, AccessConfig_v1alpha1_FromProto)
	out.AliasIPRanges = direct.Slice_FromProto(mapCtx, in.AliasIpRanges, AliasIPRange_v1alpha1_FromProto)
	out.StackType = direct.Enum_FromProto(mapCtx, in.GetStackType())
	out.IPV6AccessType = direct.Enum_FromProto(mapCtx, in.GetIpv6AccessType())
	out.QueueCount = in.QueueCount
	out.NicType = direct.Enum_FromProto(mapCtx, in.GetNicType())
	out.NetworkAttachment = in.NetworkAttachment
	return out
}

func NetworkInterfaceObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.NetworkInterfaceObservedState) *pb.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.NetworkInterface{}
	out.Network = in.Network
	out.Subnetwork = in.Subnetwork
	out.IpAddress = in.IPAddress
	out.Ipv6Address = in.IPV6Address
	out.InternalIpv6PrefixLength = in.InternalIPV6PrefixLength
	out.Name = in.Name
	out.AccessConfigs = direct.Slice_ToProto(mapCtx, in.AccessConfigs, AccessConfig_v1alpha1_ToProto)
	out.Ipv6AccessConfigs = direct.Slice_ToProto(mapCtx, in.IPV6AccessConfigs, AccessConfig_v1alpha1_ToProto)
	out.AliasIpRanges = direct.Slice_ToProto(mapCtx, in.AliasIPRanges, AliasIPRange_v1alpha1_ToProto)
	if in.StackType != nil {
		v := direct.Enum_ToProto[pb.NetworkInterface_StackType](mapCtx, in.StackType)
		out.StackType = &v
	}
	if in.IPV6AccessType != nil {
		v := direct.Enum_ToProto[pb.NetworkInterface_Ipv6AccessType](mapCtx, in.IPV6AccessType)
		out.Ipv6AccessType = &v
	}
	out.QueueCount = in.QueueCount
	if in.NicType != nil {
		v := direct.Enum_ToProto[pb.NetworkInterface_NicType](mapCtx, in.NicType)
		out.NicType = &v
	}
	out.NetworkAttachment = in.NetworkAttachment
	return out
}
