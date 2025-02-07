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

package alloydb

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/alloydb/apiv1/alloydbpb"
)
func AlloydbClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.AlloydbClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbClusterObservedState{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SslConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PscConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	return out
}
func AlloydbClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SslConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PscConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	return out
}
func AlloydbClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.AlloydbClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbClusterSpec{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SslConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PscConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	return out
}
func AlloydbClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SslConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PscConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	return out
}
func AutomatedBackupPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupPolicy) *krm.AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &krm.AutomatedBackupPolicy{}
	out.WeeklySchedule = AutomatedBackupPolicy_WeeklySchedule_FromProto(mapCtx, in.GetWeeklySchedule())
	out.TimeBasedRetention = AutomatedBackupPolicy_TimeBasedRetention_FromProto(mapCtx, in.GetTimeBasedRetention())
	out.QuantityBasedRetention = AutomatedBackupPolicy_QuantityBasedRetention_FromProto(mapCtx, in.GetQuantityBasedRetention())
	out.Enabled = in.Enabled
	out.BackupWindow = direct.StringDuration_FromProto(mapCtx, in.GetBackupWindow())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Labels = in.Labels
	return out
}
func AutomatedBackupPolicy_ToProto(mapCtx *direct.MapContext, in *krm.AutomatedBackupPolicy) *pb.AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupPolicy{}
	if oneof := AutomatedBackupPolicy_WeeklySchedule_ToProto(mapCtx, in.WeeklySchedule); oneof != nil {
		out.Schedule = &pb.AutomatedBackupPolicy_WeeklySchedule_{WeeklySchedule: oneof}
	}
	if oneof := AutomatedBackupPolicy_TimeBasedRetention_ToProto(mapCtx, in.TimeBasedRetention); oneof != nil {
		out.Retention = &pb.AutomatedBackupPolicy_TimeBasedRetention_{TimeBasedRetention: oneof}
	}
	if oneof := AutomatedBackupPolicy_QuantityBasedRetention_ToProto(mapCtx, in.QuantityBasedRetention); oneof != nil {
		out.Retention = &pb.AutomatedBackupPolicy_QuantityBasedRetention_{QuantityBasedRetention: oneof}
	}
	out.Enabled = in.Enabled
	out.BackupWindow = direct.StringDuration_ToProto(mapCtx, in.BackupWindow)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	out.Location = direct.ValueOf(in.Location)
	out.Labels = in.Labels
	return out
}
func AutomatedBackupPolicy_QuantityBasedRetention_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupPolicy_QuantityBasedRetention) *krm.AutomatedBackupPolicy_QuantityBasedRetention {
	if in == nil {
		return nil
	}
	out := &krm.AutomatedBackupPolicy_QuantityBasedRetention{}
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func AutomatedBackupPolicy_QuantityBasedRetention_ToProto(mapCtx *direct.MapContext, in *krm.AutomatedBackupPolicy_QuantityBasedRetention) *pb.AutomatedBackupPolicy_QuantityBasedRetention {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupPolicy_QuantityBasedRetention{}
	out.Count = direct.ValueOf(in.Count)
	return out
}
func AutomatedBackupPolicy_TimeBasedRetention_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupPolicy_TimeBasedRetention) *krm.AutomatedBackupPolicy_TimeBasedRetention {
	if in == nil {
		return nil
	}
	out := &krm.AutomatedBackupPolicy_TimeBasedRetention{}
	out.RetentionPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRetentionPeriod())
	return out
}
func AutomatedBackupPolicy_TimeBasedRetention_ToProto(mapCtx *direct.MapContext, in *krm.AutomatedBackupPolicy_TimeBasedRetention) *pb.AutomatedBackupPolicy_TimeBasedRetention {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupPolicy_TimeBasedRetention{}
	out.RetentionPeriod = direct.StringDuration_ToProto(mapCtx, in.RetentionPeriod)
	return out
}
func AutomatedBackupPolicy_WeeklySchedule_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupPolicy_WeeklySchedule) *krm.AutomatedBackupPolicy_WeeklySchedule {
	if in == nil {
		return nil
	}
	out := &krm.AutomatedBackupPolicy_WeeklySchedule{}
	out.StartTimes = direct.Slice_FromProto(mapCtx, in.StartTimes, TimeOfDay_FromProto)
	out.DaysOfWeek = direct.EnumSlice_FromProto(mapCtx, in.DaysOfWeek)
	return out
}
func AutomatedBackupPolicy_WeeklySchedule_ToProto(mapCtx *direct.MapContext, in *krm.AutomatedBackupPolicy_WeeklySchedule) *pb.AutomatedBackupPolicy_WeeklySchedule {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupPolicy_WeeklySchedule{}
	out.StartTimes = direct.Slice_ToProto(mapCtx, in.StartTimes, TimeOfDay_ToProto)
	out.DaysOfWeek = direct.EnumSlice_ToProto[pb.DayOfWeek](mapCtx, in.DaysOfWeek)
	return out
}
func BackupSource_FromProto(mapCtx *direct.MapContext, in *pb.BackupSource) *krm.BackupSource {
	if in == nil {
		return nil
	}
	out := &krm.BackupSource{}
	// MISSING: BackupUid
	out.BackupName = direct.LazyPtr(in.GetBackupName())
	return out
}
func BackupSource_ToProto(mapCtx *direct.MapContext, in *krm.BackupSource) *pb.BackupSource {
	if in == nil {
		return nil
	}
	out := &pb.BackupSource{}
	// MISSING: BackupUid
	out.BackupName = direct.ValueOf(in.BackupName)
	return out
}
func BackupSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupSource) *krm.BackupSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupSourceObservedState{}
	out.BackupUid = direct.LazyPtr(in.GetBackupUid())
	// MISSING: BackupName
	return out
}
func BackupSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupSourceObservedState) *pb.BackupSource {
	if in == nil {
		return nil
	}
	out := &pb.BackupSource{}
	out.BackupUid = direct.ValueOf(in.BackupUid)
	// MISSING: BackupName
	return out
}
func Cluster_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.Cluster {
	if in == nil {
		return nil
	}
	out := &krm.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	// MISSING: State
	// MISSING: ClusterType
	out.DatabaseVersion = direct.Enum_FromProto(mapCtx, in.GetDatabaseVersion())
	out.NetworkConfig = Cluster_NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Annotations = in.Annotations
	// MISSING: Reconciling
	out.InitialUser = UserPassword_FromProto(mapCtx, in.GetInitialUser())
	out.AutomatedBackupPolicy = AutomatedBackupPolicy_FromProto(mapCtx, in.GetAutomatedBackupPolicy())
	out.SslConfig = SslConfig_FromProto(mapCtx, in.GetSslConfig())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	// MISSING: EncryptionInfo
	out.ContinuousBackupConfig = ContinuousBackupConfig_FromProto(mapCtx, in.GetContinuousBackupConfig())
	// MISSING: ContinuousBackupInfo
	out.SecondaryConfig = Cluster_SecondaryConfig_FromProto(mapCtx, in.GetSecondaryConfig())
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	out.PscConfig = Cluster_PscConfig_FromProto(mapCtx, in.GetPscConfig())
	out.MaintenanceUpdatePolicy = MaintenanceUpdatePolicy_FromProto(mapCtx, in.GetMaintenanceUpdatePolicy())
	// MISSING: MaintenanceSchedule
	out.SubscriptionType = direct.Enum_FromProto(mapCtx, in.GetSubscriptionType())
	// MISSING: TrialMetadata
	out.Tags = in.Tags
	return out
}
func Cluster_ToProto(mapCtx *direct.MapContext, in *krm.Cluster) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	// MISSING: State
	// MISSING: ClusterType
	out.DatabaseVersion = direct.Enum_ToProto[pb.DatabaseVersion](mapCtx, in.DatabaseVersion)
	out.NetworkConfig = Cluster_NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	out.Network = direct.ValueOf(in.Network)
	out.Etag = direct.ValueOf(in.Etag)
	out.Annotations = in.Annotations
	// MISSING: Reconciling
	out.InitialUser = UserPassword_ToProto(mapCtx, in.InitialUser)
	out.AutomatedBackupPolicy = AutomatedBackupPolicy_ToProto(mapCtx, in.AutomatedBackupPolicy)
	out.SslConfig = SslConfig_ToProto(mapCtx, in.SslConfig)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	// MISSING: EncryptionInfo
	out.ContinuousBackupConfig = ContinuousBackupConfig_ToProto(mapCtx, in.ContinuousBackupConfig)
	// MISSING: ContinuousBackupInfo
	out.SecondaryConfig = Cluster_SecondaryConfig_ToProto(mapCtx, in.SecondaryConfig)
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	out.PscConfig = Cluster_PscConfig_ToProto(mapCtx, in.PscConfig)
	out.MaintenanceUpdatePolicy = MaintenanceUpdatePolicy_ToProto(mapCtx, in.MaintenanceUpdatePolicy)
	// MISSING: MaintenanceSchedule
	out.SubscriptionType = direct.Enum_ToProto[pb.SubscriptionType](mapCtx, in.SubscriptionType)
	// MISSING: TrialMetadata
	out.Tags = in.Tags
	return out
}
func ClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.ClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClusterObservedState{}
	out.BackupSource = BackupSource_FromProto(mapCtx, in.GetBackupSource())
	out.MigrationSource = MigrationSource_FromProto(mapCtx, in.GetMigrationSource())
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ClusterType = direct.Enum_FromProto(mapCtx, in.GetClusterType())
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SslConfig
	// MISSING: EncryptionConfig
	out.EncryptionInfo = EncryptionInfo_FromProto(mapCtx, in.GetEncryptionInfo())
	// MISSING: ContinuousBackupConfig
	out.ContinuousBackupInfo = ContinuousBackupInfo_FromProto(mapCtx, in.GetContinuousBackupInfo())
	// MISSING: SecondaryConfig
	out.PrimaryConfig = Cluster_PrimaryConfig_FromProto(mapCtx, in.GetPrimaryConfig())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	// MISSING: PscConfig
	// MISSING: MaintenanceUpdatePolicy
	out.MaintenanceSchedule = MaintenanceSchedule_FromProto(mapCtx, in.GetMaintenanceSchedule())
	// MISSING: SubscriptionType
	out.TrialMetadata = Cluster_TrialMetadata_FromProto(mapCtx, in.GetTrialMetadata())
	// MISSING: Tags
	return out
}
func ClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	if oneof := BackupSource_ToProto(mapCtx, in.BackupSource); oneof != nil {
		out.Source = &pb.Cluster_BackupSource{BackupSource: oneof}
	}
	if oneof := MigrationSource_ToProto(mapCtx, in.MigrationSource); oneof != nil {
		out.Source = &pb.Cluster_MigrationSource{MigrationSource: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.Cluster_State](mapCtx, in.State)
	out.ClusterType = direct.Enum_ToProto[pb.Cluster_ClusterType](mapCtx, in.ClusterType)
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	out.Reconciling = direct.ValueOf(in.Reconciling)
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SslConfig
	// MISSING: EncryptionConfig
	out.EncryptionInfo = EncryptionInfo_ToProto(mapCtx, in.EncryptionInfo)
	// MISSING: ContinuousBackupConfig
	out.ContinuousBackupInfo = ContinuousBackupInfo_ToProto(mapCtx, in.ContinuousBackupInfo)
	// MISSING: SecondaryConfig
	out.PrimaryConfig = Cluster_PrimaryConfig_ToProto(mapCtx, in.PrimaryConfig)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	// MISSING: PscConfig
	// MISSING: MaintenanceUpdatePolicy
	out.MaintenanceSchedule = MaintenanceSchedule_ToProto(mapCtx, in.MaintenanceSchedule)
	// MISSING: SubscriptionType
	out.TrialMetadata = Cluster_TrialMetadata_ToProto(mapCtx, in.TrialMetadata)
	// MISSING: Tags
	return out
}
func Cluster_NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_NetworkConfig) *krm.Cluster_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_NetworkConfig{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.AllocatedIPRange = direct.LazyPtr(in.GetAllocatedIpRange())
	return out
}
func Cluster_NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_NetworkConfig) *pb.Cluster_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_NetworkConfig{}
	out.Network = direct.ValueOf(in.Network)
	out.AllocatedIpRange = direct.ValueOf(in.AllocatedIPRange)
	return out
}
func Cluster_PrimaryConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_PrimaryConfig) *krm.Cluster_PrimaryConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_PrimaryConfig{}
	// MISSING: SecondaryClusterNames
	return out
}
func Cluster_PrimaryConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_PrimaryConfig) *pb.Cluster_PrimaryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PrimaryConfig{}
	// MISSING: SecondaryClusterNames
	return out
}
func Cluster_PrimaryConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_PrimaryConfig) *krm.Cluster_PrimaryConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_PrimaryConfigObservedState{}
	out.SecondaryClusterNames = in.SecondaryClusterNames
	return out
}
func Cluster_PrimaryConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_PrimaryConfigObservedState) *pb.Cluster_PrimaryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PrimaryConfig{}
	out.SecondaryClusterNames = in.SecondaryClusterNames
	return out
}
func Cluster_PscConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_PscConfig) *krm.Cluster_PscConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_PscConfig{}
	out.PscEnabled = direct.LazyPtr(in.GetPscEnabled())
	return out
}
func Cluster_PscConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_PscConfig) *pb.Cluster_PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PscConfig{}
	out.PscEnabled = direct.ValueOf(in.PscEnabled)
	return out
}
func Cluster_SecondaryConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_SecondaryConfig) *krm.Cluster_SecondaryConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_SecondaryConfig{}
	out.PrimaryClusterName = direct.LazyPtr(in.GetPrimaryClusterName())
	return out
}
func Cluster_SecondaryConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_SecondaryConfig) *pb.Cluster_SecondaryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_SecondaryConfig{}
	out.PrimaryClusterName = direct.ValueOf(in.PrimaryClusterName)
	return out
}
func ContinuousBackupConfig_FromProto(mapCtx *direct.MapContext, in *pb.ContinuousBackupConfig) *krm.ContinuousBackupConfig {
	if in == nil {
		return nil
	}
	out := &krm.ContinuousBackupConfig{}
	out.Enabled = in.Enabled
	out.RecoveryWindowDays = direct.LazyPtr(in.GetRecoveryWindowDays())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	return out
}
func ContinuousBackupConfig_ToProto(mapCtx *direct.MapContext, in *krm.ContinuousBackupConfig) *pb.ContinuousBackupConfig {
	if in == nil {
		return nil
	}
	out := &pb.ContinuousBackupConfig{}
	out.Enabled = in.Enabled
	out.RecoveryWindowDays = direct.ValueOf(in.RecoveryWindowDays)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	return out
}
func ContinuousBackupInfo_FromProto(mapCtx *direct.MapContext, in *pb.ContinuousBackupInfo) *krm.ContinuousBackupInfo {
	if in == nil {
		return nil
	}
	out := &krm.ContinuousBackupInfo{}
	// MISSING: EncryptionInfo
	// MISSING: EnabledTime
	// MISSING: Schedule
	// MISSING: EarliestRestorableTime
	return out
}
func ContinuousBackupInfo_ToProto(mapCtx *direct.MapContext, in *krm.ContinuousBackupInfo) *pb.ContinuousBackupInfo {
	if in == nil {
		return nil
	}
	out := &pb.ContinuousBackupInfo{}
	// MISSING: EncryptionInfo
	// MISSING: EnabledTime
	// MISSING: Schedule
	// MISSING: EarliestRestorableTime
	return out
}
func ContinuousBackupInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ContinuousBackupInfo) *krm.ContinuousBackupInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContinuousBackupInfoObservedState{}
	out.EncryptionInfo = EncryptionInfo_FromProto(mapCtx, in.GetEncryptionInfo())
	out.EnabledTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEnabledTime())
	out.Schedule = direct.EnumSlice_FromProto(mapCtx, in.Schedule)
	out.EarliestRestorableTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEarliestRestorableTime())
	return out
}
func ContinuousBackupInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContinuousBackupInfoObservedState) *pb.ContinuousBackupInfo {
	if in == nil {
		return nil
	}
	out := &pb.ContinuousBackupInfo{}
	out.EncryptionInfo = EncryptionInfo_ToProto(mapCtx, in.EncryptionInfo)
	out.EnabledTime = direct.StringTimestamp_ToProto(mapCtx, in.EnabledTime)
	out.Schedule = direct.EnumSlice_ToProto[pb.DayOfWeek](mapCtx, in.Schedule)
	out.EarliestRestorableTime = direct.StringTimestamp_ToProto(mapCtx, in.EarliestRestorableTime)
	return out
}
func EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfig{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}
func EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfig) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	return out
}
func EncryptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: KMSKeyVersions
	return out
}
func EncryptionInfo_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfo) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: KMSKeyVersions
	return out
}
func EncryptionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfoObservedState{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	out.KMSKeyVersions = in.KmsKeyVersions
	return out
}
func EncryptionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfoObservedState) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionInfo_Type](mapCtx, in.EncryptionType)
	out.KmsKeyVersions = in.KMSKeyVersions
	return out
}
func MaintenanceSchedule_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceSchedule) *krm.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceSchedule{}
	// MISSING: StartTime
	return out
}
func MaintenanceSchedule_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceSchedule) *pb.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceSchedule{}
	// MISSING: StartTime
	return out
}
func MaintenanceScheduleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceSchedule) *krm.MaintenanceScheduleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceScheduleObservedState{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	return out
}
func MaintenanceScheduleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceScheduleObservedState) *pb.MaintenanceSchedule {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceSchedule{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	return out
}
func MaintenanceUpdatePolicy_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceUpdatePolicy) *krm.MaintenanceUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceUpdatePolicy{}
	out.MaintenanceWindows = direct.Slice_FromProto(mapCtx, in.MaintenanceWindows, MaintenanceUpdatePolicy_MaintenanceWindow_FromProto)
	return out
}
func MaintenanceUpdatePolicy_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceUpdatePolicy) *pb.MaintenanceUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceUpdatePolicy{}
	out.MaintenanceWindows = direct.Slice_ToProto(mapCtx, in.MaintenanceWindows, MaintenanceUpdatePolicy_MaintenanceWindow_ToProto)
	return out
}
func MaintenanceUpdatePolicy_MaintenanceWindow_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceUpdatePolicy_MaintenanceWindow) *krm.MaintenanceUpdatePolicy_MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceUpdatePolicy_MaintenanceWindow{}
	out.Day = direct.Enum_FromProto(mapCtx, in.GetDay())
	out.StartTime = TimeOfDay_FromProto(mapCtx, in.GetStartTime())
	return out
}
func MaintenanceUpdatePolicy_MaintenanceWindow_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceUpdatePolicy_MaintenanceWindow) *pb.MaintenanceUpdatePolicy_MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceUpdatePolicy_MaintenanceWindow{}
	out.Day = direct.Enum_ToProto[pb.DayOfWeek](mapCtx, in.Day)
	out.StartTime = TimeOfDay_ToProto(mapCtx, in.StartTime)
	return out
}
func MigrationSource_FromProto(mapCtx *direct.MapContext, in *pb.MigrationSource) *krm.MigrationSource {
	if in == nil {
		return nil
	}
	out := &krm.MigrationSource{}
	// MISSING: HostPort
	// MISSING: ReferenceID
	// MISSING: SourceType
	return out
}
func MigrationSource_ToProto(mapCtx *direct.MapContext, in *krm.MigrationSource) *pb.MigrationSource {
	if in == nil {
		return nil
	}
	out := &pb.MigrationSource{}
	// MISSING: HostPort
	// MISSING: ReferenceID
	// MISSING: SourceType
	return out
}
func MigrationSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationSource) *krm.MigrationSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationSourceObservedState{}
	out.HostPort = direct.LazyPtr(in.GetHostPort())
	out.ReferenceID = direct.LazyPtr(in.GetReferenceId())
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	return out
}
func MigrationSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationSourceObservedState) *pb.MigrationSource {
	if in == nil {
		return nil
	}
	out := &pb.MigrationSource{}
	out.HostPort = direct.ValueOf(in.HostPort)
	out.ReferenceId = direct.ValueOf(in.ReferenceID)
	out.SourceType = direct.Enum_ToProto[pb.MigrationSource_MigrationSourceType](mapCtx, in.SourceType)
	return out
}
func SslConfig_FromProto(mapCtx *direct.MapContext, in *pb.SslConfig) *krm.SslConfig {
	if in == nil {
		return nil
	}
	out := &krm.SslConfig{}
	out.SslMode = direct.Enum_FromProto(mapCtx, in.GetSslMode())
	out.CaSource = direct.Enum_FromProto(mapCtx, in.GetCaSource())
	return out
}
func SslConfig_ToProto(mapCtx *direct.MapContext, in *krm.SslConfig) *pb.SslConfig {
	if in == nil {
		return nil
	}
	out := &pb.SslConfig{}
	out.SslMode = direct.Enum_ToProto[pb.SslConfig_SslMode](mapCtx, in.SslMode)
	out.CaSource = direct.Enum_ToProto[pb.SslConfig_CaSource](mapCtx, in.CaSource)
	return out
}
func UserPassword_FromProto(mapCtx *direct.MapContext, in *pb.UserPassword) *krm.UserPassword {
	if in == nil {
		return nil
	}
	out := &krm.UserPassword{}
	out.User = direct.LazyPtr(in.GetUser())
	out.Password = direct.LazyPtr(in.GetPassword())
	return out
}
func UserPassword_ToProto(mapCtx *direct.MapContext, in *krm.UserPassword) *pb.UserPassword {
	if in == nil {
		return nil
	}
	out := &pb.UserPassword{}
	out.User = direct.ValueOf(in.User)
	out.Password = direct.ValueOf(in.Password)
	return out
}
