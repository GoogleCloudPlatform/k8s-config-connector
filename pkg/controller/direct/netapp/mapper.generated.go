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

package netapp

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BackupConfig_FromProto(mapCtx *direct.MapContext, in *pb.BackupConfig) *krm.BackupConfig {
	if in == nil {
		return nil
	}
	out := &krm.BackupConfig{}
	out.BackupPolicies = in.BackupPolicies
	out.BackupVault = direct.LazyPtr(in.GetBackupVault())
	out.ScheduledBackupEnabled = in.ScheduledBackupEnabled
	// MISSING: BackupChainBytes
	return out
}
func BackupConfig_ToProto(mapCtx *direct.MapContext, in *krm.BackupConfig) *pb.BackupConfig {
	if in == nil {
		return nil
	}
	out := &pb.BackupConfig{}
	out.BackupPolicies = in.BackupPolicies
	out.BackupVault = direct.ValueOf(in.BackupVault)
	out.ScheduledBackupEnabled = in.ScheduledBackupEnabled
	// MISSING: BackupChainBytes
	return out
}
func BackupConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupConfig) *krm.BackupConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupConfigObservedState{}
	// MISSING: BackupPolicies
	// MISSING: BackupVault
	// MISSING: ScheduledBackupEnabled
	out.BackupChainBytes = in.BackupChainBytes
	return out
}
func BackupConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupConfigObservedState) *pb.BackupConfig {
	if in == nil {
		return nil
	}
	out := &pb.BackupConfig{}
	// MISSING: BackupPolicies
	// MISSING: BackupVault
	// MISSING: ScheduledBackupEnabled
	out.BackupChainBytes = in.BackupChainBytes
	return out
}
func DailySchedule_FromProto(mapCtx *direct.MapContext, in *pb.DailySchedule) *krm.DailySchedule {
	if in == nil {
		return nil
	}
	out := &krm.DailySchedule{}
	out.SnapshotsToKeep = in.SnapshotsToKeep
	out.Minute = in.Minute
	out.Hour = in.Hour
	return out
}
func DailySchedule_ToProto(mapCtx *direct.MapContext, in *krm.DailySchedule) *pb.DailySchedule {
	if in == nil {
		return nil
	}
	out := &pb.DailySchedule{}
	out.SnapshotsToKeep = in.SnapshotsToKeep
	out.Minute = in.Minute
	out.Hour = in.Hour
	return out
}
func ExportPolicy_FromProto(mapCtx *direct.MapContext, in *pb.ExportPolicy) *krm.ExportPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ExportPolicy{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, SimpleExportPolicyRule_FromProto)
	return out
}
func ExportPolicy_ToProto(mapCtx *direct.MapContext, in *krm.ExportPolicy) *pb.ExportPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ExportPolicy{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, SimpleExportPolicyRule_ToProto)
	return out
}
func HourlySchedule_FromProto(mapCtx *direct.MapContext, in *pb.HourlySchedule) *krm.HourlySchedule {
	if in == nil {
		return nil
	}
	out := &krm.HourlySchedule{}
	out.SnapshotsToKeep = in.SnapshotsToKeep
	out.Minute = in.Minute
	return out
}
func HourlySchedule_ToProto(mapCtx *direct.MapContext, in *krm.HourlySchedule) *pb.HourlySchedule {
	if in == nil {
		return nil
	}
	out := &pb.HourlySchedule{}
	out.SnapshotsToKeep = in.SnapshotsToKeep
	out.Minute = in.Minute
	return out
}
func HybridReplicationParameters_FromProto(mapCtx *direct.MapContext, in *pb.HybridReplicationParameters) *krm.HybridReplicationParameters {
	if in == nil {
		return nil
	}
	out := &krm.HybridReplicationParameters{}
	out.Replication = direct.LazyPtr(in.GetReplication())
	out.PeerVolumeName = direct.LazyPtr(in.GetPeerVolumeName())
	out.PeerClusterName = direct.LazyPtr(in.GetPeerClusterName())
	out.PeerSvmName = direct.LazyPtr(in.GetPeerSvmName())
	out.PeerIPAddresses = in.PeerIpAddresses
	out.ClusterLocation = direct.LazyPtr(in.GetClusterLocation())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	return out
}
func HybridReplicationParameters_ToProto(mapCtx *direct.MapContext, in *krm.HybridReplicationParameters) *pb.HybridReplicationParameters {
	if in == nil {
		return nil
	}
	out := &pb.HybridReplicationParameters{}
	out.Replication = direct.ValueOf(in.Replication)
	out.PeerVolumeName = direct.ValueOf(in.PeerVolumeName)
	out.PeerClusterName = direct.ValueOf(in.PeerClusterName)
	out.PeerSvmName = direct.ValueOf(in.PeerSvmName)
	out.PeerIpAddresses = in.PeerIPAddresses
	out.ClusterLocation = direct.ValueOf(in.ClusterLocation)
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	return out
}
func MonthlySchedule_FromProto(mapCtx *direct.MapContext, in *pb.MonthlySchedule) *krm.MonthlySchedule {
	if in == nil {
		return nil
	}
	out := &krm.MonthlySchedule{}
	out.SnapshotsToKeep = in.SnapshotsToKeep
	out.Minute = in.Minute
	out.Hour = in.Hour
	out.DaysOfMonth = in.DaysOfMonth
	return out
}
func MonthlySchedule_ToProto(mapCtx *direct.MapContext, in *krm.MonthlySchedule) *pb.MonthlySchedule {
	if in == nil {
		return nil
	}
	out := &pb.MonthlySchedule{}
	out.SnapshotsToKeep = in.SnapshotsToKeep
	out.Minute = in.Minute
	out.Hour = in.Hour
	out.DaysOfMonth = in.DaysOfMonth
	return out
}
func MountOption_FromProto(mapCtx *direct.MapContext, in *pb.MountOption) *krm.MountOption {
	if in == nil {
		return nil
	}
	out := &krm.MountOption{}
	out.Export = direct.LazyPtr(in.GetExport())
	out.ExportFull = direct.LazyPtr(in.GetExportFull())
	out.Protocol = direct.Enum_FromProto(mapCtx, in.GetProtocol())
	out.Instructions = direct.LazyPtr(in.GetInstructions())
	return out
}
func MountOption_ToProto(mapCtx *direct.MapContext, in *krm.MountOption) *pb.MountOption {
	if in == nil {
		return nil
	}
	out := &pb.MountOption{}
	out.Export = direct.ValueOf(in.Export)
	out.ExportFull = direct.ValueOf(in.ExportFull)
	out.Protocol = direct.Enum_ToProto[pb.Protocols](mapCtx, in.Protocol)
	out.Instructions = direct.ValueOf(in.Instructions)
	return out
}
func NetappVolumeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.NetappVolumeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetappVolumeObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	// MISSING: ShareName
	// MISSING: PsaRange
	// MISSING: StoragePool
	// MISSING: Network
	// MISSING: ServiceLevel
	// MISSING: CapacityGib
	// MISSING: ExportPolicy
	// MISSING: Protocols
	// MISSING: SmbSettings
	// MISSING: MountOptions
	// MISSING: UnixPermissions
	// MISSING: Labels
	// MISSING: Description
	// MISSING: SnapshotPolicy
	// MISSING: SnapReserve
	// MISSING: SnapshotDirectory
	// MISSING: UsedGib
	// MISSING: SecurityStyle
	// MISSING: KerberosEnabled
	// MISSING: LdapEnabled
	// MISSING: ActiveDirectory
	// MISSING: RestoreParameters
	// MISSING: KMSConfig
	// MISSING: EncryptionType
	// MISSING: HasReplication
	// MISSING: BackupConfig
	// MISSING: RestrictedActions
	// MISSING: LargeCapacity
	// MISSING: MultipleEndpoints
	// MISSING: TieringPolicy
	// MISSING: ReplicaZone
	// MISSING: Zone
	// MISSING: ColdTierSizeGib
	// MISSING: HybridReplicationParameters
	return out
}
func NetappVolumeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetappVolumeObservedState) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	// MISSING: Name
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	// MISSING: ShareName
	// MISSING: PsaRange
	// MISSING: StoragePool
	// MISSING: Network
	// MISSING: ServiceLevel
	// MISSING: CapacityGib
	// MISSING: ExportPolicy
	// MISSING: Protocols
	// MISSING: SmbSettings
	// MISSING: MountOptions
	// MISSING: UnixPermissions
	// MISSING: Labels
	// MISSING: Description
	// MISSING: SnapshotPolicy
	// MISSING: SnapReserve
	// MISSING: SnapshotDirectory
	// MISSING: UsedGib
	// MISSING: SecurityStyle
	// MISSING: KerberosEnabled
	// MISSING: LdapEnabled
	// MISSING: ActiveDirectory
	// MISSING: RestoreParameters
	// MISSING: KMSConfig
	// MISSING: EncryptionType
	// MISSING: HasReplication
	// MISSING: BackupConfig
	// MISSING: RestrictedActions
	// MISSING: LargeCapacity
	// MISSING: MultipleEndpoints
	// MISSING: TieringPolicy
	// MISSING: ReplicaZone
	// MISSING: Zone
	// MISSING: ColdTierSizeGib
	// MISSING: HybridReplicationParameters
	return out
}
func NetappVolumeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.NetappVolumeSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetappVolumeSpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	// MISSING: ShareName
	// MISSING: PsaRange
	// MISSING: StoragePool
	// MISSING: Network
	// MISSING: ServiceLevel
	// MISSING: CapacityGib
	// MISSING: ExportPolicy
	// MISSING: Protocols
	// MISSING: SmbSettings
	// MISSING: MountOptions
	// MISSING: UnixPermissions
	// MISSING: Labels
	// MISSING: Description
	// MISSING: SnapshotPolicy
	// MISSING: SnapReserve
	// MISSING: SnapshotDirectory
	// MISSING: UsedGib
	// MISSING: SecurityStyle
	// MISSING: KerberosEnabled
	// MISSING: LdapEnabled
	// MISSING: ActiveDirectory
	// MISSING: RestoreParameters
	// MISSING: KMSConfig
	// MISSING: EncryptionType
	// MISSING: HasReplication
	// MISSING: BackupConfig
	// MISSING: RestrictedActions
	// MISSING: LargeCapacity
	// MISSING: MultipleEndpoints
	// MISSING: TieringPolicy
	// MISSING: ReplicaZone
	// MISSING: Zone
	// MISSING: ColdTierSizeGib
	// MISSING: HybridReplicationParameters
	return out
}
func NetappVolumeSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetappVolumeSpec) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	// MISSING: Name
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	// MISSING: ShareName
	// MISSING: PsaRange
	// MISSING: StoragePool
	// MISSING: Network
	// MISSING: ServiceLevel
	// MISSING: CapacityGib
	// MISSING: ExportPolicy
	// MISSING: Protocols
	// MISSING: SmbSettings
	// MISSING: MountOptions
	// MISSING: UnixPermissions
	// MISSING: Labels
	// MISSING: Description
	// MISSING: SnapshotPolicy
	// MISSING: SnapReserve
	// MISSING: SnapshotDirectory
	// MISSING: UsedGib
	// MISSING: SecurityStyle
	// MISSING: KerberosEnabled
	// MISSING: LdapEnabled
	// MISSING: ActiveDirectory
	// MISSING: RestoreParameters
	// MISSING: KMSConfig
	// MISSING: EncryptionType
	// MISSING: HasReplication
	// MISSING: BackupConfig
	// MISSING: RestrictedActions
	// MISSING: LargeCapacity
	// MISSING: MultipleEndpoints
	// MISSING: TieringPolicy
	// MISSING: ReplicaZone
	// MISSING: Zone
	// MISSING: ColdTierSizeGib
	// MISSING: HybridReplicationParameters
	return out
}
func RestoreParameters_FromProto(mapCtx *direct.MapContext, in *pb.RestoreParameters) *krm.RestoreParameters {
	if in == nil {
		return nil
	}
	out := &krm.RestoreParameters{}
	out.SourceSnapshot = direct.LazyPtr(in.GetSourceSnapshot())
	out.SourceBackup = direct.LazyPtr(in.GetSourceBackup())
	return out
}
func RestoreParameters_ToProto(mapCtx *direct.MapContext, in *krm.RestoreParameters) *pb.RestoreParameters {
	if in == nil {
		return nil
	}
	out := &pb.RestoreParameters{}
	if oneof := RestoreParameters_SourceSnapshot_ToProto(mapCtx, in.SourceSnapshot); oneof != nil {
		out.Source = oneof
	}
	if oneof := RestoreParameters_SourceBackup_ToProto(mapCtx, in.SourceBackup); oneof != nil {
		out.Source = oneof
	}
	return out
}
func SimpleExportPolicyRule_FromProto(mapCtx *direct.MapContext, in *pb.SimpleExportPolicyRule) *krm.SimpleExportPolicyRule {
	if in == nil {
		return nil
	}
	out := &krm.SimpleExportPolicyRule{}
	out.AllowedClients = in.AllowedClients
	out.HasRootAccess = in.HasRootAccess
	out.AccessType = direct.Enum_FromProto(mapCtx, in.GetAccessType())
	out.Nfsv3 = in.Nfsv3
	out.Nfsv4 = in.Nfsv4
	out.Kerberos5ReadOnly = in.Kerberos5ReadOnly
	out.Kerberos5ReadWrite = in.Kerberos5ReadWrite
	out.Kerberos5iReadOnly = in.Kerberos5iReadOnly
	out.Kerberos5iReadWrite = in.Kerberos5iReadWrite
	out.Kerberos5pReadOnly = in.Kerberos5pReadOnly
	out.Kerberos5pReadWrite = in.Kerberos5pReadWrite
	return out
}
func SimpleExportPolicyRule_ToProto(mapCtx *direct.MapContext, in *krm.SimpleExportPolicyRule) *pb.SimpleExportPolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.SimpleExportPolicyRule{}
	out.AllowedClients = in.AllowedClients
	out.HasRootAccess = in.HasRootAccess
	if oneof := SimpleExportPolicyRule_AccessType_ToProto(mapCtx, in.AccessType); oneof != nil {
		out.AccessType = oneof
	}
	out.Nfsv3 = in.Nfsv3
	out.Nfsv4 = in.Nfsv4
	out.Kerberos5ReadOnly = in.Kerberos5ReadOnly
	out.Kerberos5ReadWrite = in.Kerberos5ReadWrite
	out.Kerberos5iReadOnly = in.Kerberos5iReadOnly
	out.Kerberos5iReadWrite = in.Kerberos5iReadWrite
	out.Kerberos5pReadOnly = in.Kerberos5pReadOnly
	out.Kerberos5pReadWrite = in.Kerberos5pReadWrite
	return out
}
func SnapshotPolicy_FromProto(mapCtx *direct.MapContext, in *pb.SnapshotPolicy) *krm.SnapshotPolicy {
	if in == nil {
		return nil
	}
	out := &krm.SnapshotPolicy{}
	out.Enabled = in.Enabled
	out.HourlySchedule = HourlySchedule_FromProto(mapCtx, in.GetHourlySchedule())
	out.DailySchedule = DailySchedule_FromProto(mapCtx, in.GetDailySchedule())
	out.WeeklySchedule = WeeklySchedule_FromProto(mapCtx, in.GetWeeklySchedule())
	out.MonthlySchedule = MonthlySchedule_FromProto(mapCtx, in.GetMonthlySchedule())
	return out
}
func SnapshotPolicy_ToProto(mapCtx *direct.MapContext, in *krm.SnapshotPolicy) *pb.SnapshotPolicy {
	if in == nil {
		return nil
	}
	out := &pb.SnapshotPolicy{}
	out.Enabled = in.Enabled
	if oneof := HourlySchedule_ToProto(mapCtx, in.HourlySchedule); oneof != nil {
		out.HourlySchedule = &pb.SnapshotPolicy_HourlySchedule{HourlySchedule: oneof}
	}
	if oneof := DailySchedule_ToProto(mapCtx, in.DailySchedule); oneof != nil {
		out.DailySchedule = &pb.SnapshotPolicy_DailySchedule{DailySchedule: oneof}
	}
	if oneof := WeeklySchedule_ToProto(mapCtx, in.WeeklySchedule); oneof != nil {
		out.WeeklySchedule = &pb.SnapshotPolicy_WeeklySchedule{WeeklySchedule: oneof}
	}
	if oneof := MonthlySchedule_ToProto(mapCtx, in.MonthlySchedule); oneof != nil {
		out.MonthlySchedule = &pb.SnapshotPolicy_MonthlySchedule{MonthlySchedule: oneof}
	}
	return out
}
func TieringPolicy_FromProto(mapCtx *direct.MapContext, in *pb.TieringPolicy) *krm.TieringPolicy {
	if in == nil {
		return nil
	}
	out := &krm.TieringPolicy{}
	out.TierAction = direct.Enum_FromProto(mapCtx, in.GetTierAction())
	out.CoolingThresholdDays = in.CoolingThresholdDays
	return out
}
func TieringPolicy_ToProto(mapCtx *direct.MapContext, in *krm.TieringPolicy) *pb.TieringPolicy {
	if in == nil {
		return nil
	}
	out := &pb.TieringPolicy{}
	if oneof := TieringPolicy_TierAction_ToProto(mapCtx, in.TierAction); oneof != nil {
		out.TierAction = oneof
	}
	out.CoolingThresholdDays = in.CoolingThresholdDays
	return out
}
func Volume_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.Volume {
	if in == nil {
		return nil
	}
	out := &krm.Volume{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	out.ShareName = direct.LazyPtr(in.GetShareName())
	// MISSING: PsaRange
	out.StoragePool = direct.LazyPtr(in.GetStoragePool())
	// MISSING: Network
	// MISSING: ServiceLevel
	out.CapacityGib = direct.LazyPtr(in.GetCapacityGib())
	out.ExportPolicy = ExportPolicy_FromProto(mapCtx, in.GetExportPolicy())
	out.Protocols = direct.EnumSlice_FromProto(mapCtx, in.Protocols)
	out.SmbSettings = direct.EnumSlice_FromProto(mapCtx, in.SmbSettings)
	// MISSING: MountOptions
	out.UnixPermissions = direct.LazyPtr(in.GetUnixPermissions())
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.SnapshotPolicy = SnapshotPolicy_FromProto(mapCtx, in.GetSnapshotPolicy())
	out.SnapReserve = direct.LazyPtr(in.GetSnapReserve())
	out.SnapshotDirectory = direct.LazyPtr(in.GetSnapshotDirectory())
	// MISSING: UsedGib
	out.SecurityStyle = direct.Enum_FromProto(mapCtx, in.GetSecurityStyle())
	out.KerberosEnabled = direct.LazyPtr(in.GetKerberosEnabled())
	// MISSING: LdapEnabled
	// MISSING: ActiveDirectory
	out.RestoreParameters = RestoreParameters_FromProto(mapCtx, in.GetRestoreParameters())
	// MISSING: KMSConfig
	// MISSING: EncryptionType
	// MISSING: HasReplication
	out.BackupConfig = BackupConfig_FromProto(mapCtx, in.GetBackupConfig())
	out.RestrictedActions = direct.EnumSlice_FromProto(mapCtx, in.RestrictedActions)
	out.LargeCapacity = direct.LazyPtr(in.GetLargeCapacity())
	out.MultipleEndpoints = direct.LazyPtr(in.GetMultipleEndpoints())
	out.TieringPolicy = TieringPolicy_FromProto(mapCtx, in.GetTieringPolicy())
	// MISSING: ReplicaZone
	// MISSING: Zone
	// MISSING: ColdTierSizeGib
	out.HybridReplicationParameters = HybridReplicationParameters_FromProto(mapCtx, in.GetHybridReplicationParameters())
	return out
}
func Volume_ToProto(mapCtx *direct.MapContext, in *krm.Volume) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	out.ShareName = direct.ValueOf(in.ShareName)
	// MISSING: PsaRange
	out.StoragePool = direct.ValueOf(in.StoragePool)
	// MISSING: Network
	// MISSING: ServiceLevel
	out.CapacityGib = direct.ValueOf(in.CapacityGib)
	out.ExportPolicy = ExportPolicy_ToProto(mapCtx, in.ExportPolicy)
	out.Protocols = direct.EnumSlice_ToProto[pb.Protocols](mapCtx, in.Protocols)
	out.SmbSettings = direct.EnumSlice_ToProto[pb.SMBSettings](mapCtx, in.SmbSettings)
	// MISSING: MountOptions
	out.UnixPermissions = direct.ValueOf(in.UnixPermissions)
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.SnapshotPolicy = SnapshotPolicy_ToProto(mapCtx, in.SnapshotPolicy)
	out.SnapReserve = direct.ValueOf(in.SnapReserve)
	out.SnapshotDirectory = direct.ValueOf(in.SnapshotDirectory)
	// MISSING: UsedGib
	out.SecurityStyle = direct.Enum_ToProto[pb.SecurityStyle](mapCtx, in.SecurityStyle)
	out.KerberosEnabled = direct.ValueOf(in.KerberosEnabled)
	// MISSING: LdapEnabled
	// MISSING: ActiveDirectory
	out.RestoreParameters = RestoreParameters_ToProto(mapCtx, in.RestoreParameters)
	// MISSING: KMSConfig
	// MISSING: EncryptionType
	// MISSING: HasReplication
	if oneof := BackupConfig_ToProto(mapCtx, in.BackupConfig); oneof != nil {
		out.BackupConfig = &pb.Volume_BackupConfig{BackupConfig: oneof}
	}
	out.RestrictedActions = direct.EnumSlice_ToProto[pb.RestrictedAction](mapCtx, in.RestrictedActions)
	out.LargeCapacity = direct.ValueOf(in.LargeCapacity)
	out.MultipleEndpoints = direct.ValueOf(in.MultipleEndpoints)
	if oneof := TieringPolicy_ToProto(mapCtx, in.TieringPolicy); oneof != nil {
		out.TieringPolicy = &pb.Volume_TieringPolicy{TieringPolicy: oneof}
	}
	// MISSING: ReplicaZone
	// MISSING: Zone
	// MISSING: ColdTierSizeGib
	out.HybridReplicationParameters = HybridReplicationParameters_ToProto(mapCtx, in.HybridReplicationParameters)
	return out
}
func VolumeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.VolumeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VolumeObservedState{}
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateDetails = direct.LazyPtr(in.GetStateDetails())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: ShareName
	out.PsaRange = direct.LazyPtr(in.GetPsaRange())
	// MISSING: StoragePool
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.ServiceLevel = direct.Enum_FromProto(mapCtx, in.GetServiceLevel())
	// MISSING: CapacityGib
	// MISSING: ExportPolicy
	// MISSING: Protocols
	// MISSING: SmbSettings
	out.MountOptions = direct.Slice_FromProto(mapCtx, in.MountOptions, MountOption_FromProto)
	// MISSING: UnixPermissions
	// MISSING: Labels
	// MISSING: Description
	// MISSING: SnapshotPolicy
	// MISSING: SnapReserve
	// MISSING: SnapshotDirectory
	out.UsedGib = direct.LazyPtr(in.GetUsedGib())
	// MISSING: SecurityStyle
	// MISSING: KerberosEnabled
	out.LdapEnabled = direct.LazyPtr(in.GetLdapEnabled())
	out.ActiveDirectory = direct.LazyPtr(in.GetActiveDirectory())
	// MISSING: RestoreParameters
	out.KMSConfig = direct.LazyPtr(in.GetKmsConfig())
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	out.HasReplication = direct.LazyPtr(in.GetHasReplication())
	out.BackupConfig = BackupConfigObservedState_FromProto(mapCtx, in.GetBackupConfig())
	// MISSING: RestrictedActions
	// MISSING: LargeCapacity
	// MISSING: MultipleEndpoints
	// MISSING: TieringPolicy
	out.ReplicaZone = direct.LazyPtr(in.GetReplicaZone())
	out.Zone = direct.LazyPtr(in.GetZone())
	out.ColdTierSizeGib = direct.LazyPtr(in.GetColdTierSizeGib())
	// MISSING: HybridReplicationParameters
	return out
}
func VolumeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VolumeObservedState) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.Volume_State](mapCtx, in.State)
	out.StateDetails = direct.ValueOf(in.StateDetails)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: ShareName
	out.PsaRange = direct.ValueOf(in.PsaRange)
	// MISSING: StoragePool
	out.Network = direct.ValueOf(in.Network)
	out.ServiceLevel = direct.Enum_ToProto[pb.ServiceLevel](mapCtx, in.ServiceLevel)
	// MISSING: CapacityGib
	// MISSING: ExportPolicy
	// MISSING: Protocols
	// MISSING: SmbSettings
	out.MountOptions = direct.Slice_ToProto(mapCtx, in.MountOptions, MountOption_ToProto)
	// MISSING: UnixPermissions
	// MISSING: Labels
	// MISSING: Description
	// MISSING: SnapshotPolicy
	// MISSING: SnapReserve
	// MISSING: SnapshotDirectory
	out.UsedGib = direct.ValueOf(in.UsedGib)
	// MISSING: SecurityStyle
	// MISSING: KerberosEnabled
	out.LdapEnabled = direct.ValueOf(in.LdapEnabled)
	out.ActiveDirectory = direct.ValueOf(in.ActiveDirectory)
	// MISSING: RestoreParameters
	out.KmsConfig = direct.ValueOf(in.KMSConfig)
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionType](mapCtx, in.EncryptionType)
	out.HasReplication = direct.ValueOf(in.HasReplication)
	if oneof := BackupConfigObservedState_ToProto(mapCtx, in.BackupConfig); oneof != nil {
		out.BackupConfig = &pb.Volume_BackupConfig{BackupConfig: oneof}
	}
	// MISSING: RestrictedActions
	// MISSING: LargeCapacity
	// MISSING: MultipleEndpoints
	// MISSING: TieringPolicy
	out.ReplicaZone = direct.ValueOf(in.ReplicaZone)
	out.Zone = direct.ValueOf(in.Zone)
	out.ColdTierSizeGib = direct.ValueOf(in.ColdTierSizeGib)
	// MISSING: HybridReplicationParameters
	return out
}
func WeeklySchedule_FromProto(mapCtx *direct.MapContext, in *pb.WeeklySchedule) *krm.WeeklySchedule {
	if in == nil {
		return nil
	}
	out := &krm.WeeklySchedule{}
	out.SnapshotsToKeep = in.SnapshotsToKeep
	out.Minute = in.Minute
	out.Hour = in.Hour
	out.Day = in.Day
	return out
}
func WeeklySchedule_ToProto(mapCtx *direct.MapContext, in *krm.WeeklySchedule) *pb.WeeklySchedule {
	if in == nil {
		return nil
	}
	out := &pb.WeeklySchedule{}
	out.SnapshotsToKeep = in.SnapshotsToKeep
	out.Minute = in.Minute
	out.Hour = in.Hour
	out.Day = in.Day
	return out
}
