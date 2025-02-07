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

package backupdr

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AcceleratorConfig_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorConfig) *krm.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krm.AcceleratorConfig{}
	out.AcceleratorType = in.AcceleratorType
	out.AcceleratorCount = in.AcceleratorCount
	return out
}
func AcceleratorConfig_ToProto(mapCtx *direct.MapContext, in *krm.AcceleratorConfig) *pb.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorConfig{}
	out.AcceleratorType = in.AcceleratorType
	out.AcceleratorCount = in.AcceleratorCount
	return out
}
func AccessConfig_FromProto(mapCtx *direct.MapContext, in *pb.AccessConfig) *krm.AccessConfig {
	if in == nil {
		return nil
	}
	out := &krm.AccessConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Name = in.Name
	out.ExternalIP = in.ExternalIp
	out.ExternalIpv6 = in.ExternalIpv6
	out.ExternalIpv6PrefixLength = in.ExternalIpv6PrefixLength
	out.SetPublicPtr = in.SetPublicPtr
	out.PublicPtrDomainName = in.PublicPtrDomainName
	out.NetworkTier = direct.Enum_FromProto(mapCtx, in.GetNetworkTier())
	return out
}
func AccessConfig_ToProto(mapCtx *direct.MapContext, in *krm.AccessConfig) *pb.AccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.AccessConfig{}
	if oneof := AccessConfig_Type_ToProto(mapCtx, in.Type); oneof != nil {
		out.Type = oneof
	}
	out.Name = in.Name
	out.ExternalIp = in.ExternalIP
	out.ExternalIpv6 = in.ExternalIpv6
	out.ExternalIpv6PrefixLength = in.ExternalIpv6PrefixLength
	out.SetPublicPtr = in.SetPublicPtr
	out.PublicPtrDomainName = in.PublicPtrDomainName
	if oneof := AccessConfig_NetworkTier_ToProto(mapCtx, in.NetworkTier); oneof != nil {
		out.NetworkTier = oneof
	}
	return out
}
func AliasIpRange_FromProto(mapCtx *direct.MapContext, in *pb.AliasIpRange) *krm.AliasIpRange {
	if in == nil {
		return nil
	}
	out := &krm.AliasIpRange{}
	out.IPCidrRange = in.IpCidrRange
	out.SubnetworkRangeName = in.SubnetworkRangeName
	return out
}
func AliasIpRange_ToProto(mapCtx *direct.MapContext, in *krm.AliasIpRange) *pb.AliasIpRange {
	if in == nil {
		return nil
	}
	out := &pb.AliasIpRange{}
	out.IpCidrRange = in.IPCidrRange
	out.SubnetworkRangeName = in.SubnetworkRangeName
	return out
}
func AttachedDisk_FromProto(mapCtx *direct.MapContext, in *pb.AttachedDisk) *krm.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &krm.AttachedDisk{}
	out.InitializeParams = AttachedDisk_InitializeParams_FromProto(mapCtx, in.GetInitializeParams())
	out.DeviceName = in.DeviceName
	out.Kind = in.Kind
	out.DiskTypeDeprecated = direct.Enum_FromProto(mapCtx, in.GetDiskTypeDeprecated())
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.Source = in.Source
	out.Index = in.Index
	out.Boot = in.Boot
	out.AutoDelete = in.AutoDelete
	out.License = in.License
	out.DiskInterface = direct.Enum_FromProto(mapCtx, in.GetDiskInterface())
	out.GuestOsFeature = direct.Slice_FromProto(mapCtx, in.GuestOsFeature, GuestOsFeature_FromProto)
	out.DiskEncryptionKey = CustomerEncryptionKey_FromProto(mapCtx, in.GetDiskEncryptionKey())
	out.DiskSizeGB = in.DiskSizeGb
	// MISSING: SavedState
	// MISSING: DiskType
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func AttachedDisk_ToProto(mapCtx *direct.MapContext, in *krm.AttachedDisk) *pb.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AttachedDisk{}
	if oneof := AttachedDisk_InitializeParams_ToProto(mapCtx, in.InitializeParams); oneof != nil {
		out.InitializeParams = &pb.AttachedDisk_InitializeParams_{InitializeParams: oneof}
	}
	out.DeviceName = in.DeviceName
	out.Kind = in.Kind
	if oneof := AttachedDisk_DiskTypeDeprecated_ToProto(mapCtx, in.DiskTypeDeprecated); oneof != nil {
		out.DiskTypeDeprecated = oneof
	}
	if oneof := AttachedDisk_Mode_ToProto(mapCtx, in.Mode); oneof != nil {
		out.Mode = oneof
	}
	out.Source = in.Source
	out.Index = in.Index
	out.Boot = in.Boot
	out.AutoDelete = in.AutoDelete
	out.License = in.License
	if oneof := AttachedDisk_DiskInterface_ToProto(mapCtx, in.DiskInterface); oneof != nil {
		out.DiskInterface = oneof
	}
	out.GuestOsFeature = direct.Slice_ToProto(mapCtx, in.GuestOsFeature, GuestOsFeature_ToProto)
	if oneof := CustomerEncryptionKey_ToProto(mapCtx, in.DiskEncryptionKey); oneof != nil {
		out.DiskEncryptionKey = &pb.AttachedDisk_DiskEncryptionKey{DiskEncryptionKey: oneof}
	}
	out.DiskSizeGb = in.DiskSizeGB
	// MISSING: SavedState
	// MISSING: DiskType
	if oneof := AttachedDisk_Type_ToProto(mapCtx, in.Type); oneof != nil {
		out.Type = oneof
	}
	return out
}
func AttachedDiskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AttachedDisk) *krm.AttachedDiskObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AttachedDiskObservedState{}
	// MISSING: InitializeParams
	// MISSING: DeviceName
	// MISSING: Kind
	// MISSING: DiskTypeDeprecated
	// MISSING: Mode
	// MISSING: Source
	// MISSING: Index
	// MISSING: Boot
	// MISSING: AutoDelete
	// MISSING: License
	// MISSING: DiskInterface
	// MISSING: GuestOsFeature
	// MISSING: DiskEncryptionKey
	// MISSING: DiskSizeGB
	out.SavedState = direct.Enum_FromProto(mapCtx, in.GetSavedState())
	out.DiskType = in.DiskType
	// MISSING: Type
	return out
}
func AttachedDiskObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AttachedDiskObservedState) *pb.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AttachedDisk{}
	// MISSING: InitializeParams
	// MISSING: DeviceName
	// MISSING: Kind
	// MISSING: DiskTypeDeprecated
	// MISSING: Mode
	// MISSING: Source
	// MISSING: Index
	// MISSING: Boot
	// MISSING: AutoDelete
	// MISSING: License
	// MISSING: DiskInterface
	// MISSING: GuestOsFeature
	// MISSING: DiskEncryptionKey
	// MISSING: DiskSizeGB
	if oneof := AttachedDiskObservedState_SavedState_ToProto(mapCtx, in.SavedState); oneof != nil {
		out.SavedState = oneof
	}
	out.DiskType = in.DiskType
	// MISSING: Type
	return out
}
func AttachedDisk_InitializeParams_FromProto(mapCtx *direct.MapContext, in *pb.AttachedDisk_InitializeParams) *krm.AttachedDisk_InitializeParams {
	if in == nil {
		return nil
	}
	out := &krm.AttachedDisk_InitializeParams{}
	out.DiskName = in.DiskName
	out.ReplicaZones = in.ReplicaZones
	return out
}
func AttachedDisk_InitializeParams_ToProto(mapCtx *direct.MapContext, in *krm.AttachedDisk_InitializeParams) *pb.AttachedDisk_InitializeParams {
	if in == nil {
		return nil
	}
	out := &pb.AttachedDisk_InitializeParams{}
	out.DiskName = in.DiskName
	out.ReplicaZones = in.ReplicaZones
	return out
}
func Backup_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.Backup {
	if in == nil {
		return nil
	}
	out := &krm.Backup{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.EnforcedRetentionEndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEnforcedRetentionEndTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: ConsistencyTime
	out.Etag = in.Etag
	// MISSING: State
	// MISSING: ServiceLocks
	out.BackupApplianceLocks = direct.Slice_FromProto(mapCtx, in.BackupApplianceLocks, BackupLock_FromProto)
	// MISSING: ComputeInstanceBackupProperties
	// MISSING: BackupApplianceBackupProperties
	// MISSING: BackupType
	// MISSING: GcpBackupPlanInfo
	// MISSING: ResourceSizeBytes
	return out
}
func Backup_ToProto(mapCtx *direct.MapContext, in *krm.Backup) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.EnforcedRetentionEndTime); oneof != nil {
		out.EnforcedRetentionEndTime = &pb.Backup_EnforcedRetentionEndTime{EnforcedRetentionEndTime: oneof}
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime); oneof != nil {
		out.ExpireTime = &pb.Backup_ExpireTime{ExpireTime: oneof}
	}
	// MISSING: ConsistencyTime
	out.Etag = in.Etag
	// MISSING: State
	// MISSING: ServiceLocks
	out.BackupApplianceLocks = direct.Slice_ToProto(mapCtx, in.BackupApplianceLocks, BackupLock_ToProto)
	// MISSING: ComputeInstanceBackupProperties
	// MISSING: BackupApplianceBackupProperties
	// MISSING: BackupType
	// MISSING: GcpBackupPlanInfo
	// MISSING: ResourceSizeBytes
	return out
}
func BackupApplianceBackupProperties_FromProto(mapCtx *direct.MapContext, in *pb.BackupApplianceBackupProperties) *krm.BackupApplianceBackupProperties {
	if in == nil {
		return nil
	}
	out := &krm.BackupApplianceBackupProperties{}
	// MISSING: GenerationID
	// MISSING: FinalizeTime
	out.RecoveryRangeStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRecoveryRangeStartTime())
	out.RecoveryRangeEndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRecoveryRangeEndTime())
	return out
}
func BackupApplianceBackupProperties_ToProto(mapCtx *direct.MapContext, in *krm.BackupApplianceBackupProperties) *pb.BackupApplianceBackupProperties {
	if in == nil {
		return nil
	}
	out := &pb.BackupApplianceBackupProperties{}
	// MISSING: GenerationID
	// MISSING: FinalizeTime
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.RecoveryRangeStartTime); oneof != nil {
		out.RecoveryRangeStartTime = &pb.BackupApplianceBackupProperties_RecoveryRangeStartTime{RecoveryRangeStartTime: oneof}
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.RecoveryRangeEndTime); oneof != nil {
		out.RecoveryRangeEndTime = &pb.BackupApplianceBackupProperties_RecoveryRangeEndTime{RecoveryRangeEndTime: oneof}
	}
	return out
}
func BackupApplianceBackupPropertiesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupApplianceBackupProperties) *krm.BackupApplianceBackupPropertiesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupApplianceBackupPropertiesObservedState{}
	out.GenerationID = in.GenerationId
	out.FinalizeTime = direct.StringTimestamp_FromProto(mapCtx, in.GetFinalizeTime())
	// MISSING: RecoveryRangeStartTime
	// MISSING: RecoveryRangeEndTime
	return out
}
func BackupApplianceBackupPropertiesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupApplianceBackupPropertiesObservedState) *pb.BackupApplianceBackupProperties {
	if in == nil {
		return nil
	}
	out := &pb.BackupApplianceBackupProperties{}
	out.GenerationId = in.GenerationID
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.FinalizeTime); oneof != nil {
		out.FinalizeTime = &pb.BackupApplianceBackupProperties_FinalizeTime{FinalizeTime: oneof}
	}
	// MISSING: RecoveryRangeStartTime
	// MISSING: RecoveryRangeEndTime
	return out
}
func BackupApplianceLockInfo_FromProto(mapCtx *direct.MapContext, in *pb.BackupApplianceLockInfo) *krm.BackupApplianceLockInfo {
	if in == nil {
		return nil
	}
	out := &krm.BackupApplianceLockInfo{}
	out.BackupApplianceID = direct.LazyPtr(in.GetBackupApplianceId())
	out.BackupApplianceName = direct.LazyPtr(in.GetBackupApplianceName())
	out.LockReason = direct.LazyPtr(in.GetLockReason())
	out.JobName = direct.LazyPtr(in.GetJobName())
	out.BackupImage = direct.LazyPtr(in.GetBackupImage())
	out.SlaID = direct.LazyPtr(in.GetSlaId())
	return out
}
func BackupApplianceLockInfo_ToProto(mapCtx *direct.MapContext, in *krm.BackupApplianceLockInfo) *pb.BackupApplianceLockInfo {
	if in == nil {
		return nil
	}
	out := &pb.BackupApplianceLockInfo{}
	out.BackupApplianceId = direct.ValueOf(in.BackupApplianceID)
	out.BackupApplianceName = direct.ValueOf(in.BackupApplianceName)
	out.LockReason = direct.ValueOf(in.LockReason)
	if oneof := BackupApplianceLockInfo_JobName_ToProto(mapCtx, in.JobName); oneof != nil {
		out.LockSource = oneof
	}
	if oneof := BackupApplianceLockInfo_BackupImage_ToProto(mapCtx, in.BackupImage); oneof != nil {
		out.LockSource = oneof
	}
	if oneof := BackupApplianceLockInfo_SlaId_ToProto(mapCtx, in.SlaID); oneof != nil {
		out.LockSource = oneof
	}
	return out
}
func BackupLock_FromProto(mapCtx *direct.MapContext, in *pb.BackupLock) *krm.BackupLock {
	if in == nil {
		return nil
	}
	out := &krm.BackupLock{}
	out.LockUntilTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLockUntilTime())
	out.BackupApplianceLockInfo = BackupApplianceLockInfo_FromProto(mapCtx, in.GetBackupApplianceLockInfo())
	// MISSING: ServiceLockInfo
	return out
}
func BackupLock_ToProto(mapCtx *direct.MapContext, in *krm.BackupLock) *pb.BackupLock {
	if in == nil {
		return nil
	}
	out := &pb.BackupLock{}
	out.LockUntilTime = direct.StringTimestamp_ToProto(mapCtx, in.LockUntilTime)
	if oneof := BackupApplianceLockInfo_ToProto(mapCtx, in.BackupApplianceLockInfo); oneof != nil {
		out.ClientLockInfo = &pb.BackupLock_BackupApplianceLockInfo{BackupApplianceLockInfo: oneof}
	}
	// MISSING: ServiceLockInfo
	return out
}
func BackupLockObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupLock) *krm.BackupLockObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupLockObservedState{}
	// MISSING: LockUntilTime
	// MISSING: BackupApplianceLockInfo
	out.ServiceLockInfo = ServiceLockInfo_FromProto(mapCtx, in.GetServiceLockInfo())
	return out
}
func BackupLockObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupLockObservedState) *pb.BackupLock {
	if in == nil {
		return nil
	}
	out := &pb.BackupLock{}
	// MISSING: LockUntilTime
	// MISSING: BackupApplianceLockInfo
	if oneof := ServiceLockInfo_ToProto(mapCtx, in.ServiceLockInfo); oneof != nil {
		out.ClientLockInfo = &pb.BackupLock_ServiceLockInfo{ServiceLockInfo: oneof}
	}
	return out
}
func BackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = in.Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: EnforcedRetentionEndTime
	// MISSING: ExpireTime
	out.ConsistencyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetConsistencyTime())
	// MISSING: Etag
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ServiceLocks = direct.Slice_FromProto(mapCtx, in.ServiceLocks, BackupLock_FromProto)
	// MISSING: BackupApplianceLocks
	out.ComputeInstanceBackupProperties = ComputeInstanceBackupProperties_FromProto(mapCtx, in.GetComputeInstanceBackupProperties())
	out.BackupApplianceBackupProperties = BackupApplianceBackupProperties_FromProto(mapCtx, in.GetBackupApplianceBackupProperties())
	out.BackupType = direct.Enum_FromProto(mapCtx, in.GetBackupType())
	out.GcpBackupPlanInfo = Backup_GCPBackupPlanInfo_FromProto(mapCtx, in.GetGcpBackupPlanInfo())
	out.ResourceSizeBytes = direct.LazyPtr(in.GetResourceSizeBytes())
	return out
}
func BackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = in.Description
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.CreateTime); oneof != nil {
		out.CreateTime = &pb.Backup_CreateTime{CreateTime: oneof}
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime); oneof != nil {
		out.UpdateTime = &pb.Backup_UpdateTime{UpdateTime: oneof}
	}
	// MISSING: Labels
	// MISSING: EnforcedRetentionEndTime
	// MISSING: ExpireTime
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.ConsistencyTime); oneof != nil {
		out.ConsistencyTime = &pb.Backup_ConsistencyTime{ConsistencyTime: oneof}
	}
	// MISSING: Etag
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.ServiceLocks = direct.Slice_ToProto(mapCtx, in.ServiceLocks, BackupLock_ToProto)
	// MISSING: BackupApplianceLocks
	if oneof := ComputeInstanceBackupProperties_ToProto(mapCtx, in.ComputeInstanceBackupProperties); oneof != nil {
		out.BackupProperties = &pb.Backup_ComputeInstanceBackupProperties{ComputeInstanceBackupProperties: oneof}
	}
	if oneof := BackupApplianceBackupProperties_ToProto(mapCtx, in.BackupApplianceBackupProperties); oneof != nil {
		out.BackupProperties = &pb.Backup_BackupApplianceBackupProperties{BackupApplianceBackupProperties: oneof}
	}
	out.BackupType = direct.Enum_ToProto[pb.Backup_BackupType](mapCtx, in.BackupType)
	if oneof := Backup_GCPBackupPlanInfo_ToProto(mapCtx, in.GcpBackupPlanInfo); oneof != nil {
		out.PlanInfo = &pb.Backup_GcpBackupPlanInfo{GcpBackupPlanInfo: oneof}
	}
	out.ResourceSizeBytes = direct.ValueOf(in.ResourceSizeBytes)
	return out
}
func Backup_GCPBackupPlanInfo_FromProto(mapCtx *direct.MapContext, in *pb.Backup_GCPBackupPlanInfo) *krm.Backup_GCPBackupPlanInfo {
	if in == nil {
		return nil
	}
	out := &krm.Backup_GCPBackupPlanInfo{}
	out.BackupPlan = direct.LazyPtr(in.GetBackupPlan())
	out.BackupPlanRuleID = direct.LazyPtr(in.GetBackupPlanRuleId())
	return out
}
func Backup_GCPBackupPlanInfo_ToProto(mapCtx *direct.MapContext, in *krm.Backup_GCPBackupPlanInfo) *pb.Backup_GCPBackupPlanInfo {
	if in == nil {
		return nil
	}
	out := &pb.Backup_GCPBackupPlanInfo{}
	out.BackupPlan = direct.ValueOf(in.BackupPlan)
	out.BackupPlanRuleId = direct.ValueOf(in.BackupPlanRuleID)
	return out
}
func BackupdrBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BackupdrBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: EnforcedRetentionEndTime
	// MISSING: ExpireTime
	// MISSING: ConsistencyTime
	// MISSING: Etag
	// MISSING: State
	// MISSING: ServiceLocks
	// MISSING: BackupApplianceLocks
	// MISSING: ComputeInstanceBackupProperties
	// MISSING: BackupApplianceBackupProperties
	// MISSING: BackupType
	// MISSING: GcpBackupPlanInfo
	// MISSING: ResourceSizeBytes
	return out
}
func BackupdrBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: EnforcedRetentionEndTime
	// MISSING: ExpireTime
	// MISSING: ConsistencyTime
	// MISSING: Etag
	// MISSING: State
	// MISSING: ServiceLocks
	// MISSING: BackupApplianceLocks
	// MISSING: ComputeInstanceBackupProperties
	// MISSING: BackupApplianceBackupProperties
	// MISSING: BackupType
	// MISSING: GcpBackupPlanInfo
	// MISSING: ResourceSizeBytes
	return out
}
func BackupdrBackupPlanAssociationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlanAssociation) *krm.BackupdrBackupPlanAssociationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanAssociationObservedState{}
	// MISSING: Name
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupdrBackupPlanAssociationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanAssociationObservedState) *pb.BackupPlanAssociation {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlanAssociation{}
	// MISSING: Name
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupdrBackupPlanAssociationSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlanAssociation) *krm.BackupdrBackupPlanAssociationSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanAssociationSpec{}
	// MISSING: Name
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupdrBackupPlanAssociationSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanAssociationSpec) *pb.BackupPlanAssociation {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlanAssociation{}
	// MISSING: Name
	// MISSING: ResourceType
	// MISSING: Resource
	// MISSING: BackupPlan
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: RulesConfigInfo
	// MISSING: DataSource
	return out
}
func BackupdrBackupPlanObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupdrBackupPlanObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupPlanObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanObservedState) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupPlanSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupdrBackupPlanSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupPlanSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupPlanSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupPlanSpec) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupRules
	// MISSING: State
	// MISSING: ResourceType
	// MISSING: Etag
	// MISSING: BackupVault
	// MISSING: BackupVaultServiceAccount
	return out
}
func BackupdrBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BackupdrBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: EnforcedRetentionEndTime
	// MISSING: ExpireTime
	// MISSING: ConsistencyTime
	// MISSING: Etag
	// MISSING: State
	// MISSING: ServiceLocks
	// MISSING: BackupApplianceLocks
	// MISSING: ComputeInstanceBackupProperties
	// MISSING: BackupApplianceBackupProperties
	// MISSING: BackupType
	// MISSING: GcpBackupPlanInfo
	// MISSING: ResourceSizeBytes
	return out
}
func BackupdrBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: EnforcedRetentionEndTime
	// MISSING: ExpireTime
	// MISSING: ConsistencyTime
	// MISSING: Etag
	// MISSING: State
	// MISSING: ServiceLocks
	// MISSING: BackupApplianceLocks
	// MISSING: ComputeInstanceBackupProperties
	// MISSING: BackupApplianceBackupProperties
	// MISSING: BackupType
	// MISSING: GcpBackupPlanInfo
	// MISSING: ResourceSizeBytes
	return out
}
func BackupdrBackupVaultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.BackupdrBackupVaultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupVaultObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupMinimumEnforcedRetentionDuration
	// MISSING: Deletable
	// MISSING: Etag
	// MISSING: State
	// MISSING: EffectiveTime
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupdrBackupVaultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupVaultObservedState) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupMinimumEnforcedRetentionDuration
	// MISSING: Deletable
	// MISSING: Etag
	// MISSING: State
	// MISSING: EffectiveTime
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupdrBackupVaultSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.BackupdrBackupVaultSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrBackupVaultSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupMinimumEnforcedRetentionDuration
	// MISSING: Deletable
	// MISSING: Etag
	// MISSING: State
	// MISSING: EffectiveTime
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupdrBackupVaultSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrBackupVaultSpec) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupMinimumEnforcedRetentionDuration
	// MISSING: Deletable
	// MISSING: Etag
	// MISSING: State
	// MISSING: EffectiveTime
	// MISSING: BackupCount
	// MISSING: ServiceAccount
	// MISSING: TotalStoredBytes
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: AccessRestriction
	return out
}
func BackupdrDataSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krm.BackupdrDataSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrDataSourceObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupCount
	// MISSING: Etag
	// MISSING: TotalStoredBytes
	// MISSING: ConfigState
	// MISSING: BackupConfigInfo
	// MISSING: DataSourceGcpResource
	// MISSING: DataSourceBackupApplianceApplication
	return out
}
func BackupdrDataSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrDataSourceObservedState) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupCount
	// MISSING: Etag
	// MISSING: TotalStoredBytes
	// MISSING: ConfigState
	// MISSING: BackupConfigInfo
	// MISSING: DataSourceGcpResource
	// MISSING: DataSourceBackupApplianceApplication
	return out
}
func BackupdrDataSourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krm.BackupdrDataSourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrDataSourceSpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupCount
	// MISSING: Etag
	// MISSING: TotalStoredBytes
	// MISSING: ConfigState
	// MISSING: BackupConfigInfo
	// MISSING: DataSourceGcpResource
	// MISSING: DataSourceBackupApplianceApplication
	return out
}
func BackupdrDataSourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrDataSourceSpec) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: BackupCount
	// MISSING: Etag
	// MISSING: TotalStoredBytes
	// MISSING: ConfigState
	// MISSING: BackupConfigInfo
	// MISSING: DataSourceGcpResource
	// MISSING: DataSourceBackupApplianceApplication
	return out
}
func BackupdrManagementServerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.BackupdrManagementServerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrManagementServerObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrManagementServerObservedState) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerSpec_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.BackupdrManagementServerSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrManagementServerSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrManagementServerSpec) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ComputeInstanceBackupProperties_FromProto(mapCtx *direct.MapContext, in *pb.ComputeInstanceBackupProperties) *krm.ComputeInstanceBackupProperties {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceBackupProperties{}
	out.Description = in.Description
	out.Tags = Tags_FromProto(mapCtx, in.GetTags())
	out.MachineType = in.MachineType
	out.CanIPForward = in.CanIpForward
	out.NetworkInterface = direct.Slice_FromProto(mapCtx, in.NetworkInterface, NetworkInterface_FromProto)
	out.Disk = direct.Slice_FromProto(mapCtx, in.Disk, AttachedDisk_FromProto)
	out.Metadata = Metadata_FromProto(mapCtx, in.GetMetadata())
	out.ServiceAccount = direct.Slice_FromProto(mapCtx, in.ServiceAccount, ServiceAccount_FromProto)
	out.Scheduling = Scheduling_FromProto(mapCtx, in.GetScheduling())
	out.GuestAccelerator = direct.Slice_FromProto(mapCtx, in.GuestAccelerator, AcceleratorConfig_FromProto)
	out.MinCpuPlatform = in.MinCpuPlatform
	out.KeyRevocationActionType = direct.Enum_FromProto(mapCtx, in.GetKeyRevocationActionType())
	out.SourceInstance = in.SourceInstance
	out.Labels = in.Labels
	return out
}
func ComputeInstanceBackupProperties_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceBackupProperties) *pb.ComputeInstanceBackupProperties {
	if in == nil {
		return nil
	}
	out := &pb.ComputeInstanceBackupProperties{}
	out.Description = in.Description
	if oneof := Tags_ToProto(mapCtx, in.Tags); oneof != nil {
		out.Tags = &pb.ComputeInstanceBackupProperties_Tags{Tags: oneof}
	}
	out.MachineType = in.MachineType
	out.CanIpForward = in.CanIPForward
	out.NetworkInterface = direct.Slice_ToProto(mapCtx, in.NetworkInterface, NetworkInterface_ToProto)
	out.Disk = direct.Slice_ToProto(mapCtx, in.Disk, AttachedDisk_ToProto)
	if oneof := Metadata_ToProto(mapCtx, in.Metadata); oneof != nil {
		out.Metadata = &pb.ComputeInstanceBackupProperties_Metadata{Metadata: oneof}
	}
	out.ServiceAccount = direct.Slice_ToProto(mapCtx, in.ServiceAccount, ServiceAccount_ToProto)
	if oneof := Scheduling_ToProto(mapCtx, in.Scheduling); oneof != nil {
		out.Scheduling = &pb.ComputeInstanceBackupProperties_Scheduling{Scheduling: oneof}
	}
	out.GuestAccelerator = direct.Slice_ToProto(mapCtx, in.GuestAccelerator, AcceleratorConfig_ToProto)
	out.MinCpuPlatform = in.MinCpuPlatform
	if oneof := ComputeInstanceBackupProperties_KeyRevocationActionType_ToProto(mapCtx, in.KeyRevocationActionType); oneof != nil {
		out.KeyRevocationActionType = oneof
	}
	out.SourceInstance = in.SourceInstance
	out.Labels = in.Labels
	return out
}
func ComputeInstanceBackupPropertiesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ComputeInstanceBackupProperties) *krm.ComputeInstanceBackupPropertiesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceBackupPropertiesObservedState{}
	// MISSING: Description
	// MISSING: Tags
	// MISSING: MachineType
	// MISSING: CanIPForward
	out.NetworkInterface = direct.Slice_FromProto(mapCtx, in.NetworkInterface, NetworkInterfaceObservedState_FromProto)
	out.Disk = direct.Slice_FromProto(mapCtx, in.Disk, AttachedDiskObservedState_FromProto)
	// MISSING: Metadata
	// MISSING: ServiceAccount
	// MISSING: Scheduling
	// MISSING: GuestAccelerator
	// MISSING: MinCpuPlatform
	// MISSING: KeyRevocationActionType
	// MISSING: SourceInstance
	// MISSING: Labels
	return out
}
func ComputeInstanceBackupPropertiesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceBackupPropertiesObservedState) *pb.ComputeInstanceBackupProperties {
	if in == nil {
		return nil
	}
	out := &pb.ComputeInstanceBackupProperties{}
	// MISSING: Description
	// MISSING: Tags
	// MISSING: MachineType
	// MISSING: CanIPForward
	out.NetworkInterface = direct.Slice_ToProto(mapCtx, in.NetworkInterface, NetworkInterfaceObservedState_ToProto)
	out.Disk = direct.Slice_ToProto(mapCtx, in.Disk, AttachedDiskObservedState_ToProto)
	// MISSING: Metadata
	// MISSING: ServiceAccount
	// MISSING: Scheduling
	// MISSING: GuestAccelerator
	// MISSING: MinCpuPlatform
	// MISSING: KeyRevocationActionType
	// MISSING: SourceInstance
	// MISSING: Labels
	return out
}
func CustomerEncryptionKey_FromProto(mapCtx *direct.MapContext, in *pb.CustomerEncryptionKey) *krm.CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &krm.CustomerEncryptionKey{}
	out.RawKey = direct.LazyPtr(in.GetRawKey())
	out.RsaEncryptedKey = direct.LazyPtr(in.GetRsaEncryptedKey())
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	out.KMSKeyServiceAccount = in.KmsKeyServiceAccount
	return out
}
func CustomerEncryptionKey_ToProto(mapCtx *direct.MapContext, in *krm.CustomerEncryptionKey) *pb.CustomerEncryptionKey {
	if in == nil {
		return nil
	}
	out := &pb.CustomerEncryptionKey{}
	if oneof := CustomerEncryptionKey_RawKey_ToProto(mapCtx, in.RawKey); oneof != nil {
		out.Key = oneof
	}
	if oneof := CustomerEncryptionKey_RsaEncryptedKey_ToProto(mapCtx, in.RsaEncryptedKey); oneof != nil {
		out.Key = oneof
	}
	if oneof := CustomerEncryptionKey_KmsKeyName_ToProto(mapCtx, in.KMSKeyName); oneof != nil {
		out.Key = oneof
	}
	out.KmsKeyServiceAccount = in.KMSKeyServiceAccount
	return out
}
func Entry_FromProto(mapCtx *direct.MapContext, in *pb.Entry) *krm.Entry {
	if in == nil {
		return nil
	}
	out := &krm.Entry{}
	out.Key = in.Key
	out.Value = in.Value
	return out
}
func Entry_ToProto(mapCtx *direct.MapContext, in *krm.Entry) *pb.Entry {
	if in == nil {
		return nil
	}
	out := &pb.Entry{}
	out.Key = in.Key
	out.Value = in.Value
	return out
}
func GuestOsFeature_FromProto(mapCtx *direct.MapContext, in *pb.GuestOsFeature) *krm.GuestOsFeature {
	if in == nil {
		return nil
	}
	out := &krm.GuestOsFeature{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func GuestOsFeature_ToProto(mapCtx *direct.MapContext, in *krm.GuestOsFeature) *pb.GuestOsFeature {
	if in == nil {
		return nil
	}
	out := &pb.GuestOsFeature{}
	if oneof := GuestOsFeature_Type_ToProto(mapCtx, in.Type); oneof != nil {
		out.Type = oneof
	}
	return out
}
func NetworkInterface_FromProto(mapCtx *direct.MapContext, in *pb.NetworkInterface) *krm.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &krm.NetworkInterface{}
	out.Network = in.Network
	out.Subnetwork = in.Subnetwork
	out.IPAddress = in.IpAddress
	out.Ipv6Address = in.Ipv6Address
	out.InternalIpv6PrefixLength = in.InternalIpv6PrefixLength
	// MISSING: Name
	out.AccessConfigs = direct.Slice_FromProto(mapCtx, in.AccessConfigs, AccessConfig_FromProto)
	out.Ipv6AccessConfigs = direct.Slice_FromProto(mapCtx, in.Ipv6AccessConfigs, AccessConfig_FromProto)
	out.AliasIPRanges = direct.Slice_FromProto(mapCtx, in.AliasIPRanges, AliasIpRange_FromProto)
	out.StackType = direct.Enum_FromProto(mapCtx, in.GetStackType())
	out.Ipv6AccessType = direct.Enum_FromProto(mapCtx, in.GetIpv6AccessType())
	out.QueueCount = in.QueueCount
	out.NicType = direct.Enum_FromProto(mapCtx, in.GetNicType())
	out.NetworkAttachment = in.NetworkAttachment
	return out
}
func NetworkInterface_ToProto(mapCtx *direct.MapContext, in *krm.NetworkInterface) *pb.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.NetworkInterface{}
	out.Network = in.Network
	out.Subnetwork = in.Subnetwork
	out.IpAddress = in.IPAddress
	out.Ipv6Address = in.Ipv6Address
	out.InternalIpv6PrefixLength = in.InternalIpv6PrefixLength
	// MISSING: Name
	out.AccessConfigs = direct.Slice_ToProto(mapCtx, in.AccessConfigs, AccessConfig_ToProto)
	out.Ipv6AccessConfigs = direct.Slice_ToProto(mapCtx, in.Ipv6AccessConfigs, AccessConfig_ToProto)
	out.AliasIpRanges = direct.Slice_ToProto(mapCtx, in.AliasIPRanges, AliasIpRange_ToProto)
	if oneof := NetworkInterface_StackType_ToProto(mapCtx, in.StackType); oneof != nil {
		out.StackType = oneof
	}
	if oneof := NetworkInterface_Ipv6AccessType_ToProto(mapCtx, in.Ipv6AccessType); oneof != nil {
		out.Ipv6AccessType = oneof
	}
	out.QueueCount = in.QueueCount
	if oneof := NetworkInterface_NicType_ToProto(mapCtx, in.NicType); oneof != nil {
		out.NicType = oneof
	}
	out.NetworkAttachment = in.NetworkAttachment
	return out
}
func NetworkInterfaceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkInterface) *krm.NetworkInterfaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkInterfaceObservedState{}
	// MISSING: Network
	// MISSING: Subnetwork
	// MISSING: IPAddress
	// MISSING: Ipv6Address
	// MISSING: InternalIpv6PrefixLength
	out.Name = in.Name
	// MISSING: AccessConfigs
	// MISSING: Ipv6AccessConfigs
	// MISSING: AliasIPRanges
	// MISSING: StackType
	// MISSING: Ipv6AccessType
	// MISSING: QueueCount
	// MISSING: NicType
	// MISSING: NetworkAttachment
	return out
}
func NetworkInterfaceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkInterfaceObservedState) *pb.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.NetworkInterface{}
	// MISSING: Network
	// MISSING: Subnetwork
	// MISSING: IPAddress
	// MISSING: Ipv6Address
	// MISSING: InternalIpv6PrefixLength
	out.Name = in.Name
	// MISSING: AccessConfigs
	// MISSING: Ipv6AccessConfigs
	// MISSING: AliasIPRanges
	// MISSING: StackType
	// MISSING: Ipv6AccessType
	// MISSING: QueueCount
	// MISSING: NicType
	// MISSING: NetworkAttachment
	return out
}
func Scheduling_FromProto(mapCtx *direct.MapContext, in *pb.Scheduling) *krm.Scheduling {
	if in == nil {
		return nil
	}
	out := &krm.Scheduling{}
	out.OnHostMaintenance = direct.Enum_FromProto(mapCtx, in.GetOnHostMaintenance())
	out.AutomaticRestart = in.AutomaticRestart
	out.Preemptible = in.Preemptible
	out.NodeAffinities = direct.Slice_FromProto(mapCtx, in.NodeAffinities, Scheduling_NodeAffinity_FromProto)
	out.MinNodeCpus = in.MinNodeCpus
	out.ProvisioningModel = direct.Enum_FromProto(mapCtx, in.GetProvisioningModel())
	out.InstanceTerminationAction = direct.Enum_FromProto(mapCtx, in.GetInstanceTerminationAction())
	out.LocalSsdRecoveryTimeout = SchedulingDuration_FromProto(mapCtx, in.GetLocalSsdRecoveryTimeout())
	return out
}
func Scheduling_ToProto(mapCtx *direct.MapContext, in *krm.Scheduling) *pb.Scheduling {
	if in == nil {
		return nil
	}
	out := &pb.Scheduling{}
	if oneof := Scheduling_OnHostMaintenance_ToProto(mapCtx, in.OnHostMaintenance); oneof != nil {
		out.OnHostMaintenance = oneof
	}
	out.AutomaticRestart = in.AutomaticRestart
	out.Preemptible = in.Preemptible
	out.NodeAffinities = direct.Slice_ToProto(mapCtx, in.NodeAffinities, Scheduling_NodeAffinity_ToProto)
	out.MinNodeCpus = in.MinNodeCpus
	if oneof := Scheduling_ProvisioningModel_ToProto(mapCtx, in.ProvisioningModel); oneof != nil {
		out.ProvisioningModel = oneof
	}
	if oneof := Scheduling_InstanceTerminationAction_ToProto(mapCtx, in.InstanceTerminationAction); oneof != nil {
		out.InstanceTerminationAction = oneof
	}
	if oneof := SchedulingDuration_ToProto(mapCtx, in.LocalSsdRecoveryTimeout); oneof != nil {
		out.LocalSsdRecoveryTimeout = &pb.Scheduling_LocalSsdRecoveryTimeout{LocalSsdRecoveryTimeout: oneof}
	}
	return out
}
func SchedulingDuration_FromProto(mapCtx *direct.MapContext, in *pb.SchedulingDuration) *krm.SchedulingDuration {
	if in == nil {
		return nil
	}
	out := &krm.SchedulingDuration{}
	out.Seconds = in.Seconds
	out.Nanos = in.Nanos
	return out
}
func SchedulingDuration_ToProto(mapCtx *direct.MapContext, in *krm.SchedulingDuration) *pb.SchedulingDuration {
	if in == nil {
		return nil
	}
	out := &pb.SchedulingDuration{}
	out.Seconds = in.Seconds
	out.Nanos = in.Nanos
	return out
}
func Scheduling_NodeAffinity_FromProto(mapCtx *direct.MapContext, in *pb.Scheduling_NodeAffinity) *krm.Scheduling_NodeAffinity {
	if in == nil {
		return nil
	}
	out := &krm.Scheduling_NodeAffinity{}
	out.Key = in.Key
	out.Operator = direct.Enum_FromProto(mapCtx, in.GetOperator())
	out.Values = in.Values
	return out
}
func Scheduling_NodeAffinity_ToProto(mapCtx *direct.MapContext, in *krm.Scheduling_NodeAffinity) *pb.Scheduling_NodeAffinity {
	if in == nil {
		return nil
	}
	out := &pb.Scheduling_NodeAffinity{}
	out.Key = in.Key
	if oneof := Scheduling_NodeAffinity_Operator_ToProto(mapCtx, in.Operator); oneof != nil {
		out.Operator = oneof
	}
	out.Values = in.Values
	return out
}
func ServiceAccount_FromProto(mapCtx *direct.MapContext, in *pb.ServiceAccount) *krm.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &krm.ServiceAccount{}
	out.Email = in.Email
	out.Scopes = in.Scopes
	return out
}
func ServiceAccount_ToProto(mapCtx *direct.MapContext, in *krm.ServiceAccount) *pb.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.ServiceAccount{}
	out.Email = in.Email
	out.Scopes = in.Scopes
	return out
}
func ServiceLockInfo_FromProto(mapCtx *direct.MapContext, in *pb.ServiceLockInfo) *krm.ServiceLockInfo {
	if in == nil {
		return nil
	}
	out := &krm.ServiceLockInfo{}
	// MISSING: Operation
	return out
}
func ServiceLockInfo_ToProto(mapCtx *direct.MapContext, in *krm.ServiceLockInfo) *pb.ServiceLockInfo {
	if in == nil {
		return nil
	}
	out := &pb.ServiceLockInfo{}
	// MISSING: Operation
	return out
}
func ServiceLockInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceLockInfo) *krm.ServiceLockInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ServiceLockInfoObservedState{}
	out.Operation = direct.LazyPtr(in.GetOperation())
	return out
}
func ServiceLockInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ServiceLockInfoObservedState) *pb.ServiceLockInfo {
	if in == nil {
		return nil
	}
	out := &pb.ServiceLockInfo{}
	out.Operation = direct.ValueOf(in.Operation)
	return out
}
func Tags_FromProto(mapCtx *direct.MapContext, in *pb.Tags) *krm.Tags {
	if in == nil {
		return nil
	}
	out := &krm.Tags{}
	out.Items = in.Items
	return out
}
func Tags_ToProto(mapCtx *direct.MapContext, in *krm.Tags) *pb.Tags {
	if in == nil {
		return nil
	}
	out := &pb.Tags{}
	out.Items = in.Items
	return out
}
