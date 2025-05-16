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
	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

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
	// MISSING: DatabaseVersion
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
	// MISSING: DatabaseVersion
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
	// MISSING: DatabaseVersion
	out.NetworkConfig = Cluster_NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	if in.GetNetwork() != "" {
		out.NetworkRef = &computev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
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
	// MISSING: DatabaseVersion
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
	return out
}
func AlloyDBInstanceStatus_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBInstanceStatus) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DeleteTime
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	// MISSING: InstanceType
	// MISSING: MachineConfig
	// MISSING: AvailabilityType
	// MISSING: GCEZone
	// MISSING: DatabaseFlags
	// MISSING: WritableNode
	// MISSING: Nodes
	// MISSING: QueryInsightsConfig
	// MISSING: ObservabilityConfig
	// MISSING: ReadPoolConfig
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.PublicIpAddress = direct.ValueOf(in.PublicIPAddress)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: UpdatePolicy
	// MISSING: ClientConnectionConfig
	// MISSING: SatisfiesPzs
	// MISSING: PSCInstanceConfig
	// MISSING: NetworkConfig
	// MISSING: GeminiConfig
	out.OutboundPublicIpAddresses = in.OutboundPublicIPAddresses
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
func BackupSource_FromProto(mapCtx *direct.MapContext, in *pb.BackupSource) *krm.BackupSource {
	if in == nil {
		return nil
	}
	out := &krm.BackupSource{}
	// MISSING: BackupUid
	if in.GetBackupName() != "" {
		out.BackupNameRef = &refs.AlloyDBBackupRef{External: in.GetBackupName()}
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
		out.NetworkRef = &computev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
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
	return out
}
func Cluster_PSCConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_PSCConfig) *pb.Cluster_PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_PscConfig{}
	out.PscEnabled = direct.ValueOf(in.PSCEnabled)
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
		out.KMSKeyNameRef = &refs.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
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
	return out
}
func Instance_MachineConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_MachineConfig) *pb.Instance_MachineConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_MachineConfig{}
	out.CpuCount = direct.ValueOf(in.CPUCount)
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
	// MISSING: TrackWaitEventTypes
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
	// MISSING: TrackWaitEventTypes
	out.MaxQueryStringLength = in.MaxQueryStringLength
	out.RecordApplicationTags = in.RecordApplicationTags
	out.QueryPlansPerMinute = in.QueryPlansPerMinute
	out.TrackActiveQueries = in.TrackActiveQueries
	out.TrackClientAddress = in.TrackClientAddress
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
	// MISSING: TrackClientAddress
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
	// MISSING: TrackClientAddress
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
