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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1alpha1"
)
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
func CloudSQLBackupRunSource_FromProto(mapCtx *direct.MapContext, in *pb.CloudSQLBackupRunSource) *krm.CloudSQLBackupRunSource {
	if in == nil {
		return nil
	}
	out := &krm.CloudSQLBackupRunSource{}
	out.Project = direct.LazyPtr(in.GetProject())
	out.InstanceID = direct.LazyPtr(in.GetInstanceId())
	out.BackupRunID = direct.LazyPtr(in.GetBackupRunId())
	return out
}
func CloudSQLBackupRunSource_ToProto(mapCtx *direct.MapContext, in *krm.CloudSQLBackupRunSource) *pb.CloudSQLBackupRunSource {
	if in == nil {
		return nil
	}
	out := &pb.CloudSQLBackupRunSource{}
	out.Project = direct.ValueOf(in.Project)
	out.InstanceId = direct.ValueOf(in.InstanceID)
	out.BackupRunId = direct.ValueOf(in.BackupRunID)
	return out
}
func Cluster_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.Cluster {
	if in == nil {
		return nil
	}
	out := &krm.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
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
	out.GeminiConfig = GeminiClusterConfig_FromProto(mapCtx, in.GetGeminiConfig())
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
	// MISSING: CloudsqlBackupRunSource
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
	out.GeminiConfig = GeminiClusterConfig_ToProto(mapCtx, in.GeminiConfig)
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
	out.CloudsqlBackupRunSource = CloudSQLBackupRunSource_FromProto(mapCtx, in.GetCloudsqlBackupRunSource())
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
	out.GeminiConfig = GeminiClusterConfigObservedState_FromProto(mapCtx, in.GetGeminiConfig())
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
	if oneof := CloudSQLBackupRunSource_ToProto(mapCtx, in.CloudsqlBackupRunSource); oneof != nil {
		out.Source = &pb.Cluster_CloudsqlBackupRunSource{CloudsqlBackupRunSource: oneof}
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
	out.GeminiConfig = GeminiClusterConfigObservedState_ToProto(mapCtx, in.GeminiConfig)
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
func GeminiClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.GeminiClusterConfig) *krm.GeminiClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.GeminiClusterConfig{}
	// MISSING: Entitled
	return out
}
func GeminiClusterConfig_ToProto(mapCtx *direct.MapContext, in *krm.GeminiClusterConfig) *pb.GeminiClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeminiClusterConfig{}
	// MISSING: Entitled
	return out
}
func GeminiClusterConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GeminiClusterConfig) *krm.GeminiClusterConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GeminiClusterConfigObservedState{}
	out.Entitled = direct.LazyPtr(in.GetEntitled())
	return out
}
func GeminiClusterConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GeminiClusterConfigObservedState) *pb.GeminiClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeminiClusterConfig{}
	out.Entitled = direct.ValueOf(in.Entitled)
	return out
}
func GeminiInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.GeminiInstanceConfig) *krm.GeminiInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.GeminiInstanceConfig{}
	out.Entitled = direct.LazyPtr(in.GetEntitled())
	return out
}
func GeminiInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.GeminiInstanceConfig) *pb.GeminiInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeminiInstanceConfig{}
	out.Entitled = direct.ValueOf(in.Entitled)
	return out
}
func Instance_ClientConnectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ClientConnectionConfig) *krm.Instance_ClientConnectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_ClientConnectionConfig{}
	out.RequireConnectors = direct.LazyPtr(in.GetRequireConnectors())
	// MISSING: SslConfig
	// (near miss): "SslConfig" vs "SSLConfig"
	return out
}
func Instance_ClientConnectionConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_ClientConnectionConfig) *pb.Instance_ClientConnectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ClientConnectionConfig{}
	out.RequireConnectors = direct.ValueOf(in.RequireConnectors)
	// MISSING: SslConfig
	// (near miss): "SslConfig" vs "SSLConfig"
	return out
}
func Instance_InstanceNetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceNetworkConfig) *krm.Instance_InstanceNetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_InstanceNetworkConfig{}
	out.AuthorizedExternalNetworks = direct.Slice_FromProto(mapCtx, in.AuthorizedExternalNetworks, Instance_InstanceNetworkConfig_AuthorizedNetwork_FromProto)
	out.EnablePublicIP = direct.LazyPtr(in.GetEnablePublicIp())
	out.EnableOutboundPublicIP = direct.LazyPtr(in.GetEnableOutboundPublicIp())
	return out
}
func Instance_InstanceNetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_InstanceNetworkConfig) *pb.Instance_InstanceNetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceNetworkConfig{}
	out.AuthorizedExternalNetworks = direct.Slice_ToProto(mapCtx, in.AuthorizedExternalNetworks, Instance_InstanceNetworkConfig_AuthorizedNetwork_ToProto)
	out.EnablePublicIp = direct.ValueOf(in.EnablePublicIP)
	out.EnableOutboundPublicIp = direct.ValueOf(in.EnableOutboundPublicIP)
	return out
}
func Instance_InstanceNetworkConfig_AuthorizedNetwork_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceNetworkConfig_AuthorizedNetwork) *krm.Instance_InstanceNetworkConfig_AuthorizedNetwork {
	if in == nil {
		return nil
	}
	out := &krm.Instance_InstanceNetworkConfig_AuthorizedNetwork{}
	out.CidrRange = direct.LazyPtr(in.GetCidrRange())
	return out
}
func Instance_InstanceNetworkConfig_AuthorizedNetwork_ToProto(mapCtx *direct.MapContext, in *krm.Instance_InstanceNetworkConfig_AuthorizedNetwork) *pb.Instance_InstanceNetworkConfig_AuthorizedNetwork {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceNetworkConfig_AuthorizedNetwork{}
	out.CidrRange = direct.ValueOf(in.CidrRange)
	return out
}
func Instance_MachineConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_MachineConfig) *krm.Instance_MachineConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_MachineConfig{}
	// MISSING: CpuCount
	// (near miss): "CpuCount" vs "CPUCount"
	return out
}
func Instance_MachineConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_MachineConfig) *pb.Instance_MachineConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_MachineConfig{}
	// MISSING: CpuCount
	// (near miss): "CpuCount" vs "CPUCount"
	return out
}
func Instance_Node_FromProto(mapCtx *direct.MapContext, in *pb.Instance_Node) *krm.Instance_Node {
	if in == nil {
		return nil
	}
	out := &krm.Instance_Node{}
	out.ZoneID = direct.LazyPtr(in.GetZoneId())
	out.ID = direct.LazyPtr(in.GetId())
	out.IP = direct.LazyPtr(in.GetIp())
	out.State = direct.LazyPtr(in.GetState())
	return out
}
func Instance_Node_ToProto(mapCtx *direct.MapContext, in *krm.Instance_Node) *pb.Instance_Node {
	if in == nil {
		return nil
	}
	out := &pb.Instance_Node{}
	out.ZoneId = direct.ValueOf(in.ZoneID)
	out.Id = direct.ValueOf(in.ID)
	out.Ip = direct.ValueOf(in.IP)
	out.State = direct.ValueOf(in.State)
	return out
}
func Instance_ObservabilityInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ObservabilityInstanceConfig) *krm.Instance_ObservabilityInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_ObservabilityInstanceConfig{}
	out.Enabled = in.Enabled
	out.PreserveComments = in.PreserveComments
	out.TrackWaitEvents = in.TrackWaitEvents
	out.TrackWaitEventTypes = in.TrackWaitEventTypes
	out.MaxQueryStringLength = in.MaxQueryStringLength
	out.RecordApplicationTags = in.RecordApplicationTags
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	out.TrackActiveQueries = in.TrackActiveQueries
	out.TrackClientAddress = in.TrackClientAddress
	return out
}
func Instance_ObservabilityInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_ObservabilityInstanceConfig) *pb.Instance_ObservabilityInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ObservabilityInstanceConfig{}
	out.Enabled = in.Enabled
	out.PreserveComments = in.PreserveComments
	out.TrackWaitEvents = in.TrackWaitEvents
	out.TrackWaitEventTypes = in.TrackWaitEventTypes
	out.MaxQueryStringLength = in.MaxQueryStringLength
	out.RecordApplicationTags = in.RecordApplicationTags
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	out.TrackActiveQueries = in.TrackActiveQueries
	out.TrackClientAddress = in.TrackClientAddress
	return out
}
func Instance_PscInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscInstanceConfig) *krm.Instance_PscInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PscInstanceConfig{}
	out.ServiceAttachmentLink = direct.LazyPtr(in.GetServiceAttachmentLink())
	out.AllowedConsumerProjects = in.AllowedConsumerProjects
	// MISSING: PscDnsName
	// (near miss): "PscDnsName" vs "PSCDNSName"
	return out
}
func Instance_PscInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PscInstanceConfig) *pb.Instance_PscInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscInstanceConfig{}
	out.ServiceAttachmentLink = direct.ValueOf(in.ServiceAttachmentLink)
	out.AllowedConsumerProjects = in.AllowedConsumerProjects
	// MISSING: PscDnsName
	// (near miss): "PscDnsName" vs "PSCDNSName"
	return out
}
func Instance_QueryInsightsInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_QueryInsightsInstanceConfig) *krm.Instance_QueryInsightsInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_QueryInsightsInstanceConfig{}
	out.RecordApplicationTags = in.RecordApplicationTags
	out.RecordClientAddress = in.RecordClientAddress
	out.QueryStringLength = direct.LazyPtr(in.GetQueryStringLength())
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	return out
}
func Instance_QueryInsightsInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_QueryInsightsInstanceConfig) *pb.Instance_QueryInsightsInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_QueryInsightsInstanceConfig{}
	out.RecordApplicationTags = in.RecordApplicationTags
	out.RecordClientAddress = in.RecordClientAddress
	out.QueryStringLength = direct.ValueOf(in.QueryStringLength)
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	return out
}
func Instance_ReadPoolConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ReadPoolConfig) *krm.Instance_ReadPoolConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_ReadPoolConfig{}
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	return out
}
func Instance_ReadPoolConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_ReadPoolConfig) *pb.Instance_ReadPoolConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ReadPoolConfig{}
	out.NodeCount = direct.ValueOf(in.NodeCount)
	return out
}
func Instance_UpdatePolicy_FromProto(mapCtx *direct.MapContext, in *pb.Instance_UpdatePolicy) *krm.Instance_UpdatePolicy {
	if in == nil {
		return nil
	}
	out := &krm.Instance_UpdatePolicy{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func Instance_UpdatePolicy_ToProto(mapCtx *direct.MapContext, in *krm.Instance_UpdatePolicy) *pb.Instance_UpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.Instance_UpdatePolicy{}
	out.Mode = direct.Enum_ToProto[pb.Instance_UpdatePolicy_Mode](mapCtx, in.Mode)
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
func SSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.SslConfig) *krm.SSLConfig {
	if in == nil {
		return nil
	}
	out := &krm.SSLConfig{}
	// MISSING: SslMode
	// (near miss): "SslMode" vs "SSLMode"
	// MISSING: CaSource
	// (near miss): "CaSource" vs "CASource"
	return out
}
func SSLConfig_ToProto(mapCtx *direct.MapContext, in *krm.SSLConfig) *pb.SslConfig {
	if in == nil {
		return nil
	}
	out := &pb.SslConfig{}
	// MISSING: SslMode
	// (near miss): "SslMode" vs "SSLMode"
	// MISSING: CaSource
	// (near miss): "CaSource" vs "CASource"
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
