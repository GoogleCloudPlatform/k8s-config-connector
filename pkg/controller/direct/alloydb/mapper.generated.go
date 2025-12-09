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

// +generated:mapper
// krm.group: alloydb.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.alloydb.v1

package alloydb

import (
	pb "cloud.google.com/go/alloydb/apiv1/alloydbpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AlloyDBBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.AlloyDBBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBBackupObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: CreateCompletionTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Type
	// MISSING: ClusterUid
	// MISSING: Reconciling
	// MISSING: EncryptionInfo
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: SizeBytes
	// MISSING: ExpiryTime
	// MISSING: ExpiryQuantity
	// MISSING: SatisfiesPzs
	// MISSING: DatabaseVersion
	// MISSING: Tags
	return out
}
func AlloyDBBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: CreateCompletionTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Type
	// MISSING: ClusterUid
	// MISSING: Reconciling
	// MISSING: EncryptionInfo
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: SizeBytes
	// MISSING: ExpiryTime
	// MISSING: ExpiryQuantity
	// MISSING: SatisfiesPzs
	// MISSING: DatabaseVersion
	// MISSING: Tags
	return out
}
func AlloyDBBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.AlloyDBBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBBackupSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: CreateCompletionTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Type
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: ClusterUid
	if in.GetClusterName() != "" {
		out.ClusterNameRef = &v1alpha1.ResourceRef{External: in.GetClusterName()}
	}
	// MISSING: Reconciling
	out.EncryptionConfig = BackupEncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	// MISSING: EncryptionInfo
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: SizeBytes
	// MISSING: ExpiryTime
	// MISSING: ExpiryQuantity
	// MISSING: SatisfiesPzs
	// MISSING: DatabaseVersion
	// MISSING: Tags
	return out
}
func AlloyDBBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: CreateCompletionTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Type
	out.Description = direct.ValueOf(in.Description)
	// MISSING: ClusterUid
	if in.ClusterNameRef != nil {
		out.ClusterName = in.ClusterNameRef.External
	}
	// MISSING: Reconciling
	out.EncryptionConfig = BackupEncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	// MISSING: EncryptionInfo
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: SizeBytes
	// MISSING: ExpiryTime
	// MISSING: ExpiryQuantity
	// MISSING: SatisfiesPzs
	// MISSING: DatabaseVersion
	// MISSING: Tags
	return out
}
func AlloyDBUserObservedState_FromProto(mapCtx *direct.MapContext, in *pb.User) *krm.AlloyDBUserObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBUserObservedState{}
	// MISSING: Name
	// MISSING: Password
	// MISSING: DatabaseRoles
	// MISSING: UserType
	// MISSING: KeepExtraRoles
	return out
}
func AlloyDBUserObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBUserObservedState) *pb.User {
	if in == nil {
		return nil
	}
	out := &pb.User{}
	// MISSING: Name
	// MISSING: Password
	// MISSING: DatabaseRoles
	// MISSING: UserType
	// MISSING: KeepExtraRoles
	return out
}
func AlloyDBUserSpec_FromProto(mapCtx *direct.MapContext, in *pb.User) *krm.AlloyDBUserSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBUserSpec{}
	// MISSING: Name
	// MISSING: Password
	// MISSING: DatabaseRoles
	// MISSING: UserType
	// MISSING: KeepExtraRoles
	return out
}
func AlloyDBUserSpec_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBUserSpec) *pb.User {
	if in == nil {
		return nil
	}
	out := &pb.User{}
	// MISSING: Name
	// MISSING: Password
	// MISSING: DatabaseRoles
	// MISSING: UserType
	// MISSING: KeepExtraRoles
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
func Backup_QuantityBasedExpiry_FromProto(mapCtx *direct.MapContext, in *pb.Backup_QuantityBasedExpiry) *krm.Backup_QuantityBasedExpiry {
	if in == nil {
		return nil
	}
	out := &krm.Backup_QuantityBasedExpiry{}
	// MISSING: RetentionCount
	// MISSING: TotalRetentionCount
	return out
}
func Backup_QuantityBasedExpiry_ToProto(mapCtx *direct.MapContext, in *krm.Backup_QuantityBasedExpiry) *pb.Backup_QuantityBasedExpiry {
	if in == nil {
		return nil
	}
	out := &pb.Backup_QuantityBasedExpiry{}
	// MISSING: RetentionCount
	// MISSING: TotalRetentionCount
	return out
}
func Backup_QuantityBasedExpiryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup_QuantityBasedExpiry) *krm.Backup_QuantityBasedExpiryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Backup_QuantityBasedExpiryObservedState{}
	out.RetentionCount = direct.LazyPtr(in.GetRetentionCount())
	out.TotalRetentionCount = direct.LazyPtr(in.GetTotalRetentionCount())
	return out
}
func Backup_QuantityBasedExpiryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Backup_QuantityBasedExpiryObservedState) *pb.Backup_QuantityBasedExpiry {
	if in == nil {
		return nil
	}
	out := &pb.Backup_QuantityBasedExpiry{}
	out.RetentionCount = direct.ValueOf(in.RetentionCount)
	out.TotalRetentionCount = direct.ValueOf(in.TotalRetentionCount)
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
	out.SSLConfig = SSLConfig_FromProto(mapCtx, in.GetSslConfig())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	// MISSING: EncryptionInfo
	out.ContinuousBackupConfig = ContinuousBackupConfig_FromProto(mapCtx, in.GetContinuousBackupConfig())
	// MISSING: ContinuousBackupInfo
	out.SecondaryConfig = Cluster_SecondaryConfig_FromProto(mapCtx, in.GetSecondaryConfig())
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	out.PSCConfig = Cluster_PSCConfig_FromProto(mapCtx, in.GetPscConfig())
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
	out.SslConfig = SSLConfig_ToProto(mapCtx, in.SSLConfig)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	// MISSING: EncryptionInfo
	out.ContinuousBackupConfig = ContinuousBackupConfig_ToProto(mapCtx, in.ContinuousBackupConfig)
	// MISSING: ContinuousBackupInfo
	out.SecondaryConfig = Cluster_SecondaryConfig_ToProto(mapCtx, in.SecondaryConfig)
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	out.PscConfig = Cluster_PSCConfig_ToProto(mapCtx, in.PSCConfig)
	out.MaintenanceUpdatePolicy = MaintenanceUpdatePolicy_ToProto(mapCtx, in.MaintenanceUpdatePolicy)
	// MISSING: MaintenanceSchedule
	out.SubscriptionType = direct.Enum_ToProto[pb.SubscriptionType](mapCtx, in.SubscriptionType)
	// MISSING: TrialMetadata
	out.Tags = in.Tags
	return out
}
func Cluster_PSCConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_PscConfig) *krm.Cluster_PSCConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_PSCConfig{}
	out.PSCEnabled = direct.LazyPtr(in.GetPscEnabled())
	// MISSING: ServiceOwnedProjectNumber
	return out
}
func Cluster_PSCConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_PSCConfig) *pb.Cluster_PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PscConfig{}
	out.PscEnabled = direct.ValueOf(in.PSCEnabled)
	// MISSING: ServiceOwnedProjectNumber
	return out
}
func Cluster_PSCConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_PscConfig) *krm.Cluster_PSCConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_PSCConfigObservedState{}
	// MISSING: PSCEnabled
	out.ServiceOwnedProjectNumber = direct.LazyPtr(in.GetServiceOwnedProjectNumber())
	return out
}
func Cluster_PSCConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_PSCConfigObservedState) *pb.Cluster_PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PscConfig{}
	// MISSING: PSCEnabled
	out.ServiceOwnedProjectNumber = direct.ValueOf(in.ServiceOwnedProjectNumber)
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
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	// MISSING: State
	out.InstanceType = direct.Enum_FromProto(mapCtx, in.GetInstanceType())
	out.MachineConfig = Instance_MachineConfig_FromProto(mapCtx, in.GetMachineConfig())
	out.AvailabilityType = direct.Enum_FromProto(mapCtx, in.GetAvailabilityType())
	out.GCEZone = direct.LazyPtr(in.GetGceZone())
	out.DatabaseFlags = in.DatabaseFlags
	// MISSING: WritableNode
	// MISSING: Nodes
	out.QueryInsightsConfig = Instance_QueryInsightsInstanceConfig_FromProto(mapCtx, in.GetQueryInsightsConfig())
	out.ObservabilityConfig = Instance_ObservabilityInstanceConfig_FromProto(mapCtx, in.GetObservabilityConfig())
	out.ReadPoolConfig = Instance_ReadPoolConfig_FromProto(mapCtx, in.GetReadPoolConfig())
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: Reconciling
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Annotations = in.Annotations
	out.ClientConnectionConfig = Instance_ClientConnectionConfig_FromProto(mapCtx, in.GetClientConnectionConfig())
	// MISSING: SatisfiesPzs
	out.PSCInstanceConfig = Instance_PSCInstanceConfig_FromProto(mapCtx, in.GetPscInstanceConfig())
	out.NetworkConfig = Instance_InstanceNetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	// MISSING: OutboundPublicIPAddresses
	out.ActivationPolicy = direct.Enum_FromProto(mapCtx, in.GetActivationPolicy())
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	// MISSING: State
	out.InstanceType = direct.Enum_ToProto[pb.Instance_InstanceType](mapCtx, in.InstanceType)
	out.MachineConfig = Instance_MachineConfig_ToProto(mapCtx, in.MachineConfig)
	out.AvailabilityType = direct.Enum_ToProto[pb.Instance_AvailabilityType](mapCtx, in.AvailabilityType)
	out.GceZone = direct.ValueOf(in.GCEZone)
	out.DatabaseFlags = in.DatabaseFlags
	// MISSING: WritableNode
	// MISSING: Nodes
	out.QueryInsightsConfig = Instance_QueryInsightsInstanceConfig_ToProto(mapCtx, in.QueryInsightsConfig)
	out.ObservabilityConfig = Instance_ObservabilityInstanceConfig_ToProto(mapCtx, in.ObservabilityConfig)
	out.ReadPoolConfig = Instance_ReadPoolConfig_ToProto(mapCtx, in.ReadPoolConfig)
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: Reconciling
	out.Etag = direct.ValueOf(in.Etag)
	out.Annotations = in.Annotations
	out.ClientConnectionConfig = Instance_ClientConnectionConfig_ToProto(mapCtx, in.ClientConnectionConfig)
	// MISSING: SatisfiesPzs
	out.PscInstanceConfig = Instance_PSCInstanceConfig_ToProto(mapCtx, in.PSCInstanceConfig)
	out.NetworkConfig = Instance_InstanceNetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	// MISSING: OutboundPublicIPAddresses
	out.ActivationPolicy = direct.Enum_ToProto[pb.Instance_ActivationPolicy](mapCtx, in.ActivationPolicy)
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: InstanceType
	// MISSING: MachineConfig
	// MISSING: AvailabilityType
	// MISSING: GCEZone
	// MISSING: DatabaseFlags
	out.WritableNode = Instance_Node_FromProto(mapCtx, in.GetWritableNode())
	out.Nodes = direct.Slice_FromProto(mapCtx, in.Nodes, Instance_Node_FromProto)
	// MISSING: QueryInsightsConfig
	out.ObservabilityConfig = Instance_ObservabilityInstanceConfigObservedState_FromProto(mapCtx, in.GetObservabilityConfig())
	// MISSING: ReadPoolConfig
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.PublicIPAddress = direct.LazyPtr(in.GetPublicIpAddress())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: ClientConnectionConfig
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.PSCInstanceConfig = Instance_PSCInstanceConfigObservedState_FromProto(mapCtx, in.GetPscInstanceConfig())
	out.NetworkConfig = Instance_InstanceNetworkConfigObservedState_FromProto(mapCtx, in.GetNetworkConfig())
	out.OutboundPublicIPAddresses = in.OutboundPublicIpAddresses
	// MISSING: ActivationPolicy
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	// MISSING: InstanceType
	// MISSING: MachineConfig
	// MISSING: AvailabilityType
	// MISSING: GCEZone
	// MISSING: DatabaseFlags
	out.WritableNode = Instance_Node_ToProto(mapCtx, in.WritableNode)
	out.Nodes = direct.Slice_ToProto(mapCtx, in.Nodes, Instance_Node_ToProto)
	// MISSING: QueryInsightsConfig
	out.ObservabilityConfig = Instance_ObservabilityInstanceConfigObservedState_ToProto(mapCtx, in.ObservabilityConfig)
	// MISSING: ReadPoolConfig
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.PublicIpAddress = direct.ValueOf(in.PublicIPAddress)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: ClientConnectionConfig
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.PscInstanceConfig = Instance_PSCInstanceConfigObservedState_ToProto(mapCtx, in.PSCInstanceConfig)
	out.NetworkConfig = Instance_InstanceNetworkConfigObservedState_ToProto(mapCtx, in.NetworkConfig)
	out.OutboundPublicIpAddresses = in.OutboundPublicIPAddresses
	// MISSING: ActivationPolicy
	return out
}
func Instance_ClientConnectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ClientConnectionConfig) *krm.Instance_ClientConnectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_ClientConnectionConfig{}
	out.RequireConnectors = direct.LazyPtr(in.GetRequireConnectors())
	out.SSLConfig = SSLConfig_FromProto(mapCtx, in.GetSslConfig())
	return out
}
func Instance_ClientConnectionConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_ClientConnectionConfig) *pb.Instance_ClientConnectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ClientConnectionConfig{}
	out.RequireConnectors = direct.ValueOf(in.RequireConnectors)
	out.SslConfig = SSLConfig_ToProto(mapCtx, in.SSLConfig)
	return out
}
func Instance_InstanceNetworkConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceNetworkConfig) *krm.Instance_InstanceNetworkConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_InstanceNetworkConfigObservedState{}
	// MISSING: AuthorizedExternalNetworks
	// MISSING: EnablePublicIP
	// MISSING: EnableOutboundPublicIP
	out.Network = direct.LazyPtr(in.GetNetwork())
	// MISSING: AllocatedIPRangeOverride
	return out
}
func Instance_InstanceNetworkConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_InstanceNetworkConfigObservedState) *pb.Instance_InstanceNetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceNetworkConfig{}
	// MISSING: AuthorizedExternalNetworks
	// MISSING: EnablePublicIP
	// MISSING: EnableOutboundPublicIP
	out.Network = direct.ValueOf(in.Network)
	// MISSING: AllocatedIPRangeOverride
	return out
}
func Instance_InstanceNetworkConfig_AuthorizedNetwork_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceNetworkConfig_AuthorizedNetwork) *krm.Instance_InstanceNetworkConfig_AuthorizedNetwork {
	if in == nil {
		return nil
	}
	out := &krm.Instance_InstanceNetworkConfig_AuthorizedNetwork{}
	out.CIDRRange = direct.LazyPtr(in.GetCidrRange())
	return out
}
func Instance_InstanceNetworkConfig_AuthorizedNetwork_ToProto(mapCtx *direct.MapContext, in *krm.Instance_InstanceNetworkConfig_AuthorizedNetwork) *pb.Instance_InstanceNetworkConfig_AuthorizedNetwork {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceNetworkConfig_AuthorizedNetwork{}
	out.CidrRange = direct.ValueOf(in.CIDRRange)
	return out
}
func Instance_MachineConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_MachineConfig) *krm.Instance_MachineConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_MachineConfig{}
	out.CPUCount = direct.LazyPtr(in.GetCpuCount())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	return out
}
func Instance_MachineConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_MachineConfig) *pb.Instance_MachineConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_MachineConfig{}
	out.CpuCount = direct.ValueOf(in.CPUCount)
	out.MachineType = direct.ValueOf(in.MachineType)
	return out
}
func Instance_Node_FromProto(mapCtx *direct.MapContext, in *pb.Instance_Node) *krm.Instance_Node {
	if in == nil {
		return nil
	}
	out := &krm.Instance_Node{}
	// MISSING: ZoneID
	// MISSING: ID
	// MISSING: IP
	// MISSING: State
	return out
}
func Instance_Node_ToProto(mapCtx *direct.MapContext, in *krm.Instance_Node) *pb.Instance_Node {
	if in == nil {
		return nil
	}
	out := &pb.Instance_Node{}
	// MISSING: ZoneID
	// MISSING: ID
	// MISSING: IP
	// MISSING: State
	return out
}
func Instance_NodeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_Node) *krm.Instance_NodeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_NodeObservedState{}
	out.ZoneID = direct.LazyPtr(in.GetZoneId())
	out.ID = direct.LazyPtr(in.GetId())
	out.IP = direct.LazyPtr(in.GetIp())
	out.State = direct.LazyPtr(in.GetState())
	return out
}
func Instance_NodeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_NodeObservedState) *pb.Instance_Node {
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
	// MISSING: TrackWaitEventTypes
	out.MaxQueryStringLength = in.MaxQueryStringLength
	out.RecordApplicationTags = in.RecordApplicationTags
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	out.TrackActiveQueries = in.TrackActiveQueries
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
	// MISSING: TrackWaitEventTypes
	out.MaxQueryStringLength = in.MaxQueryStringLength
	out.RecordApplicationTags = in.RecordApplicationTags
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	out.TrackActiveQueries = in.TrackActiveQueries
	return out
}
func Instance_ObservabilityInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ObservabilityInstanceConfig) *krm.Instance_ObservabilityInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_ObservabilityInstanceConfigObservedState{}
	// MISSING: Enabled
	// MISSING: PreserveComments
	// MISSING: TrackWaitEvents
	out.TrackWaitEventTypes = in.TrackWaitEventTypes
	// MISSING: MaxQueryStringLength
	// MISSING: RecordApplicationTags
	// MISSING: QueryPlansPerMinute
	// MISSING: TrackActiveQueries
	return out
}
func Instance_ObservabilityInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_ObservabilityInstanceConfigObservedState) *pb.Instance_ObservabilityInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ObservabilityInstanceConfig{}
	// MISSING: Enabled
	// MISSING: PreserveComments
	// MISSING: TrackWaitEvents
	out.TrackWaitEventTypes = in.TrackWaitEventTypes
	// MISSING: MaxQueryStringLength
	// MISSING: RecordApplicationTags
	// MISSING: QueryPlansPerMinute
	// MISSING: TrackActiveQueries
	return out
}
func Instance_PSCAutoConnectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscAutoConnectionConfig) *krm.Instance_PSCAutoConnectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PSCAutoConnectionConfig{}
	out.ConsumerProject = direct.LazyPtr(in.GetConsumerProject())
	out.ConsumerNetwork = direct.LazyPtr(in.GetConsumerNetwork())
	// MISSING: IPAddress
	// MISSING: Status
	// MISSING: ConsumerNetworkStatus
	return out
}
func Instance_PSCAutoConnectionConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PSCAutoConnectionConfig) *pb.Instance_PscAutoConnectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscAutoConnectionConfig{}
	out.ConsumerProject = direct.ValueOf(in.ConsumerProject)
	out.ConsumerNetwork = direct.ValueOf(in.ConsumerNetwork)
	// MISSING: IPAddress
	// MISSING: Status
	// MISSING: ConsumerNetworkStatus
	return out
}
func Instance_PSCAutoConnectionConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscAutoConnectionConfig) *krm.Instance_PSCAutoConnectionConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PSCAutoConnectionConfigObservedState{}
	// MISSING: ConsumerProject
	// MISSING: ConsumerNetwork
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.Status = direct.LazyPtr(in.GetStatus())
	out.ConsumerNetworkStatus = direct.LazyPtr(in.GetConsumerNetworkStatus())
	return out
}
func Instance_PSCAutoConnectionConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PSCAutoConnectionConfigObservedState) *pb.Instance_PscAutoConnectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscAutoConnectionConfig{}
	// MISSING: ConsumerProject
	// MISSING: ConsumerNetwork
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.Status = direct.ValueOf(in.Status)
	out.ConsumerNetworkStatus = direct.ValueOf(in.ConsumerNetworkStatus)
	return out
}
func Instance_PSCInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscInstanceConfig) *krm.Instance_PSCInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PSCInstanceConfig{}
	// MISSING: ServiceAttachmentLink
	out.AllowedConsumerProjects = in.AllowedConsumerProjects
	// MISSING: PSCDNSName
	out.PSCInterfaceConfigs = direct.Slice_FromProto(mapCtx, in.PscInterfaceConfigs, Instance_PSCInterfaceConfig_FromProto)
	out.PSCAutoConnections = direct.Slice_FromProto(mapCtx, in.PscAutoConnections, Instance_PSCAutoConnectionConfig_FromProto)
	return out
}
func Instance_PSCInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PSCInstanceConfig) *pb.Instance_PscInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscInstanceConfig{}
	// MISSING: ServiceAttachmentLink
	out.AllowedConsumerProjects = in.AllowedConsumerProjects
	// MISSING: PSCDNSName
	out.PscInterfaceConfigs = direct.Slice_ToProto(mapCtx, in.PSCInterfaceConfigs, Instance_PSCInterfaceConfig_ToProto)
	out.PscAutoConnections = direct.Slice_ToProto(mapCtx, in.PSCAutoConnections, Instance_PSCAutoConnectionConfig_ToProto)
	return out
}
func Instance_PSCInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscInstanceConfig) *krm.Instance_PSCInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PSCInstanceConfigObservedState{}
	out.ServiceAttachmentLink = direct.LazyPtr(in.GetServiceAttachmentLink())
	// MISSING: AllowedConsumerProjects
	out.PSCDNSName = direct.LazyPtr(in.GetPscDnsName())
	// MISSING: PSCInterfaceConfigs
	out.PSCAutoConnections = direct.Slice_FromProto(mapCtx, in.PscAutoConnections, Instance_PSCAutoConnectionConfigObservedState_FromProto)
	return out
}
func Instance_PSCInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PSCInstanceConfigObservedState) *pb.Instance_PscInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscInstanceConfig{}
	out.ServiceAttachmentLink = direct.ValueOf(in.ServiceAttachmentLink)
	// MISSING: AllowedConsumerProjects
	out.PscDnsName = direct.ValueOf(in.PSCDNSName)
	// MISSING: PSCInterfaceConfigs
	out.PscAutoConnections = direct.Slice_ToProto(mapCtx, in.PSCAutoConnections, Instance_PSCAutoConnectionConfigObservedState_ToProto)
	return out
}
func Instance_PSCInterfaceConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PscInterfaceConfig) *krm.Instance_PSCInterfaceConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PSCInterfaceConfig{}
	out.NetworkAttachmentResource = direct.LazyPtr(in.GetNetworkAttachmentResource())
	return out
}
func Instance_PSCInterfaceConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PSCInterfaceConfig) *pb.Instance_PscInterfaceConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PscInterfaceConfig{}
	out.NetworkAttachmentResource = direct.ValueOf(in.NetworkAttachmentResource)
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
	out.DenyMaintenancePeriods = direct.Slice_FromProto(mapCtx, in.DenyMaintenancePeriods, MaintenanceUpdatePolicy_DenyMaintenancePeriod_FromProto)
	return out
}
func MaintenanceUpdatePolicy_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceUpdatePolicy) *pb.MaintenanceUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceUpdatePolicy{}
	out.MaintenanceWindows = direct.Slice_ToProto(mapCtx, in.MaintenanceWindows, MaintenanceUpdatePolicy_MaintenanceWindow_ToProto)
	out.DenyMaintenancePeriods = direct.Slice_ToProto(mapCtx, in.DenyMaintenancePeriods, MaintenanceUpdatePolicy_DenyMaintenancePeriod_ToProto)
	return out
}
func MaintenanceUpdatePolicy_DenyMaintenancePeriod_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceUpdatePolicy_DenyMaintenancePeriod) *krm.MaintenanceUpdatePolicy_DenyMaintenancePeriod {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceUpdatePolicy_DenyMaintenancePeriod{}
	out.StartDate = Date_FromProto(mapCtx, in.GetStartDate())
	out.EndDate = Date_FromProto(mapCtx, in.GetEndDate())
	out.Time = TimeOfDay_FromProto(mapCtx, in.GetTime())
	return out
}
func MaintenanceUpdatePolicy_DenyMaintenancePeriod_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceUpdatePolicy_DenyMaintenancePeriod) *pb.MaintenanceUpdatePolicy_DenyMaintenancePeriod {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceUpdatePolicy_DenyMaintenancePeriod{}
	out.StartDate = Date_ToProto(mapCtx, in.StartDate)
	out.EndDate = Date_ToProto(mapCtx, in.EndDate)
	out.Time = TimeOfDay_ToProto(mapCtx, in.Time)
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
func SSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.SslConfig) *krm.SSLConfig {
	if in == nil {
		return nil
	}
	out := &krm.SSLConfig{}
	out.SSLMode = direct.Enum_FromProto(mapCtx, in.GetSslMode())
	out.CASource = direct.Enum_FromProto(mapCtx, in.GetCaSource())
	return out
}
func SSLConfig_ToProto(mapCtx *direct.MapContext, in *krm.SSLConfig) *pb.SslConfig {
	if in == nil {
		return nil
	}
	out := &pb.SslConfig{}
	out.SslMode = direct.Enum_ToProto[pb.SslConfig_SslMode](mapCtx, in.SSLMode)
	out.CaSource = direct.Enum_ToProto[pb.SslConfig_CaSource](mapCtx, in.CASource)
	return out
}
