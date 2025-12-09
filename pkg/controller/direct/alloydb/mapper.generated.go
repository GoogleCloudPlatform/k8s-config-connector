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
// proto.service: google.cloud.alloydb.v1beta

package alloydb

import (
	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
		out.ClusterNameRef = &krm.ClusterRef{External: in.GetClusterName()}
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
func AlloyDBClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.AlloyDBClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBClusterObservedState{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	out.ClusterType = direct.Enum_FromProto(mapCtx, in.GetClusterType())
	out.DatabaseVersion = direct.Enum_FromProto(mapCtx, in.GetDatabaseVersion())
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: SSLConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupInfo
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCConfig
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	// MISSING: ServiceAccountEmail
	return out
}
func AlloyDBClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	out.ClusterType = direct.Enum_ToProto[pb.Cluster_ClusterType](mapCtx, in.ClusterType)
	out.DatabaseVersion = direct.Enum_ToProto[pb.DatabaseVersion](mapCtx, in.DatabaseVersion)
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: SSLConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupInfo
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCConfig
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	// MISSING: ServiceAccountEmail
	return out
}
func AlloyDBClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.AlloyDBClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBClusterSpec{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: CloudsqlBackupRunSource
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	out.ClusterType = direct.Enum_FromProto(mapCtx, in.GetClusterType())
	out.DatabaseVersion = direct.Enum_FromProto(mapCtx, in.GetDatabaseVersion())
	out.NetworkConfig = Cluster_NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	if in.GetNetwork() != "" {
		out.NetworkRef = &krmcomputev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	out.InitialUser = UserPassword_FromProto(mapCtx, in.GetInitialUser())
	out.AutomatedBackupPolicy = AutomatedBackupPolicy_FromProto(mapCtx, in.GetAutomatedBackupPolicy())
	// MISSING: SSLConfig
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	// MISSING: EncryptionInfo
	out.ContinuousBackupConfig = ContinuousBackupConfig_FromProto(mapCtx, in.GetContinuousBackupConfig())
	// MISSING: ContinuousBackupInfo
	out.SecondaryConfig = Cluster_SecondaryConfig_FromProto(mapCtx, in.GetSecondaryConfig())
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCConfig
	out.MaintenanceUpdatePolicy = MaintenanceUpdatePolicy_FromProto(mapCtx, in.GetMaintenanceUpdatePolicy())
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	// MISSING: ServiceAccountEmail
	return out
}
func AlloyDBClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBClusterSpec) *pb.Cluster {
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
	// MISSING: Labels
	// MISSING: State
	out.ClusterType = direct.Enum_ToProto[pb.Cluster_ClusterType](mapCtx, in.ClusterType)
	out.DatabaseVersion = direct.Enum_ToProto[pb.DatabaseVersion](mapCtx, in.DatabaseVersion)
	out.NetworkConfig = Cluster_NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	out.InitialUser = UserPassword_ToProto(mapCtx, in.InitialUser)
	out.AutomatedBackupPolicy = AutomatedBackupPolicy_ToProto(mapCtx, in.AutomatedBackupPolicy)
	// MISSING: SSLConfig
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	// MISSING: EncryptionInfo
	out.ContinuousBackupConfig = ContinuousBackupConfig_ToProto(mapCtx, in.ContinuousBackupConfig)
	// MISSING: ContinuousBackupInfo
	out.SecondaryConfig = Cluster_SecondaryConfig_ToProto(mapCtx, in.SecondaryConfig)
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCConfig
	out.MaintenanceUpdatePolicy = MaintenanceUpdatePolicy_ToProto(mapCtx, in.MaintenanceUpdatePolicy)
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	// MISSING: ServiceAccountEmail
	return out
}
func AlloyDBInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.AlloyDBInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBInstanceObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: WritableNode
	// MISSING: Nodes
	// MISSING: QueryInsightsConfig
	// MISSING: ObservabilityConfig
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: Reconciling
	// MISSING: Etag
	// MISSING: UpdatePolicy
	// MISSING: ClientConnectionConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCInstanceConfig
	// MISSING: GeminiConfig
	// MISSING: OutboundPublicIPAddresses
	// MISSING: ActivationPolicy
	// MISSING: ConnectionPoolConfig
	// MISSING: GcaConfig
	return out
}
func AlloyDBInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: WritableNode
	// MISSING: Nodes
	// MISSING: QueryInsightsConfig
	// MISSING: ObservabilityConfig
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: Reconciling
	// MISSING: Etag
	// MISSING: UpdatePolicy
	// MISSING: ClientConnectionConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCInstanceConfig
	// MISSING: GeminiConfig
	// MISSING: OutboundPublicIPAddresses
	// MISSING: ActivationPolicy
	// MISSING: ConnectionPoolConfig
	// MISSING: GcaConfig
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
func BackupEncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.BackupEncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.BackupEncryptionConfig{}
	// MISSING: KMSKeyName
	// (near miss): "KMSKeyName" vs "KmsKeyName"
	return out
}
func BackupEncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.BackupEncryptionConfig) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	// MISSING: KMSKeyName
	// (near miss): "KMSKeyName" vs "KmsKeyName"
	return out
}
func BackupSource_FromProto(mapCtx *direct.MapContext, in *pb.BackupSource) *krm.BackupSource {
	if in == nil {
		return nil
	}
	out := &krm.BackupSource{}
	// MISSING: BackupUid
	if in.GetBackupName() != "" {
		out.BackupNameRef = &refsv1beta1.AlloyDBBackupRef{External: in.GetBackupName()}
	}
	return out
}
func BackupSource_ToProto(mapCtx *direct.MapContext, in *krm.BackupSource) *pb.BackupSource {
	if in == nil {
		return nil
	}
	out := &pb.BackupSource{}
	// MISSING: BackupUid
	if in.BackupNameRef != nil {
		out.BackupName = in.BackupNameRef.External
	}
	return out
}
func BackupSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupSource) *krm.BackupSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupSourceObservedState{}
	// MISSING: BackupUid
	out.BackupName = direct.LazyPtr(in.GetBackupName())
	return out
}
func BackupSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupSourceObservedState) *pb.BackupSource {
	if in == nil {
		return nil
	}
	out := &pb.BackupSource{}
	// MISSING: BackupUid
	out.BackupName = direct.ValueOf(in.BackupName)
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
func Cluster_NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_NetworkConfig) *krm.Cluster_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_NetworkConfig{}
	if in.GetNetwork() != "" {
		out.NetworkRef = &krmcomputev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.AllocatedIPRange = direct.LazyPtr(in.GetAllocatedIpRange())
	return out
}
func Cluster_NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_NetworkConfig) *pb.Cluster_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_NetworkConfig{}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	out.AllocatedIpRange = direct.ValueOf(in.AllocatedIPRange)
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
func EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfig{}
	if in.GetKmsKeyName() != "" {
		out.KMSKeyNameRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	return out
}
func EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfig) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	if in.KMSKeyNameRef != nil {
		out.KmsKeyName = in.KMSKeyNameRef.External
	}
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
func GcaInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.GCAInstanceConfig) *krm.GcaInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.GcaInstanceConfig{}
	// MISSING: GcaEntitlement
	return out
}
func GcaInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.GcaInstanceConfig) *pb.GCAInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.GCAInstanceConfig{}
	// MISSING: GcaEntitlement
	return out
}
func GcaInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GCAInstanceConfig) *krm.GcaInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GcaInstanceConfigObservedState{}
	out.GcaEntitlement = direct.Enum_FromProto(mapCtx, in.GetGcaEntitlement())
	return out
}
func GcaInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GcaInstanceConfigObservedState) *pb.GCAInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.GCAInstanceConfig{}
	out.GcaEntitlement = direct.Enum_ToProto[pb.GCAEntitlementType](mapCtx, in.GcaEntitlement)
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
	// MISSING: Entitled
	return out
}
func GeminiInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.GeminiInstanceConfig) *pb.GeminiInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeminiInstanceConfig{}
	// MISSING: Entitled
	return out
}
func GeminiInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GeminiInstanceConfig) *krm.GeminiInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GeminiInstanceConfigObservedState{}
	out.Entitled = direct.LazyPtr(in.GetEntitled())
	return out
}
func GeminiInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GeminiInstanceConfigObservedState) *pb.GeminiInstanceConfig {
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
func Instance_ConnectionPoolConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ConnectionPoolConfig) *krm.Instance_ConnectionPoolConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_ConnectionPoolConfig{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.Flags = in.Flags
	return out
}
func Instance_ConnectionPoolConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_ConnectionPoolConfig) *pb.Instance_ConnectionPoolConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ConnectionPoolConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.Flags = in.Flags
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
	// MISSING: Network
	// MISSING: AllocatedIPRangeOverride
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
	// MISSING: Network
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
	out.TrackClientAddress = in.TrackClientAddress
	out.AssistiveExperiencesEnabled = in.AssistiveExperiencesEnabled
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
	out.TrackClientAddress = in.TrackClientAddress
	out.AssistiveExperiencesEnabled = in.AssistiveExperiencesEnabled
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
	// MISSING: DenyMaintenancePeriods
	return out
}
func MaintenanceUpdatePolicy_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceUpdatePolicy) *pb.MaintenanceUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceUpdatePolicy{}
	out.MaintenanceWindows = direct.Slice_ToProto(mapCtx, in.MaintenanceWindows, MaintenanceUpdatePolicy_MaintenanceWindow_ToProto)
	// MISSING: DenyMaintenancePeriods
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
