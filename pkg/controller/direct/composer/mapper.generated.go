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

package composer

import (
	pb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/composer/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AirflowMetadataRetentionPolicyConfig_FromProto(mapCtx *direct.MapContext, in *pb.AirflowMetadataRetentionPolicyConfig) *krm.AirflowMetadataRetentionPolicyConfig {
	if in == nil {
		return nil
	}
	out := &krm.AirflowMetadataRetentionPolicyConfig{}
	out.RetentionMode = direct.Enum_FromProto(mapCtx, in.GetRetentionMode())
	out.RetentionDays = direct.LazyPtr(in.GetRetentionDays())
	return out
}
func AirflowMetadataRetentionPolicyConfig_ToProto(mapCtx *direct.MapContext, in *krm.AirflowMetadataRetentionPolicyConfig) *pb.AirflowMetadataRetentionPolicyConfig {
	if in == nil {
		return nil
	}
	out := &pb.AirflowMetadataRetentionPolicyConfig{}
	out.RetentionMode = direct.Enum_ToProto[pb.AirflowMetadataRetentionPolicyConfig_RetentionMode](mapCtx, in.RetentionMode)
	out.RetentionDays = direct.ValueOf(in.RetentionDays)
	return out
}
func CloudDataLineageIntegration_FromProto(mapCtx *direct.MapContext, in *pb.CloudDataLineageIntegration) *krm.CloudDataLineageIntegration {
	if in == nil {
		return nil
	}
	out := &krm.CloudDataLineageIntegration{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}
func CloudDataLineageIntegration_ToProto(mapCtx *direct.MapContext, in *krm.CloudDataLineageIntegration) *pb.CloudDataLineageIntegration {
	if in == nil {
		return nil
	}
	out := &pb.CloudDataLineageIntegration{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
func ComposerEnvironmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.ComposerEnvironmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComposerEnvironmentObservedState{}
	// MISSING: Name
	out.Config = EnvironmentConfigObservedState_FromProto(mapCtx, in.GetConfig())
	out.Uuid = direct.LazyPtr(in.GetUuid())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func ComposerEnvironmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ComposerEnvironmentObservedState) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	// MISSING: Name
	out.Config = EnvironmentConfigObservedState_ToProto(mapCtx, in.Config)
	out.Uuid = direct.ValueOf(in.Uuid)
	out.State = direct.Enum_ToProto[pb.Environment_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func ComposerEnvironmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.ComposerEnvironmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComposerEnvironmentSpec{}
	// MISSING: Name
	out.Config = EnvironmentConfig_FromProto(mapCtx, in.GetConfig())
	out.Labels = in.Labels
	out.StorageConfig = StorageConfig_FromProto(mapCtx, in.GetStorageConfig())
	return out
}
func ComposerEnvironmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComposerEnvironmentSpec) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	// MISSING: Name
	out.Config = EnvironmentConfig_ToProto(mapCtx, in.Config)
	out.Labels = in.Labels
	out.StorageConfig = StorageConfig_ToProto(mapCtx, in.StorageConfig)
	return out
}
func DataRetentionConfig_FromProto(mapCtx *direct.MapContext, in *pb.DataRetentionConfig) *krm.DataRetentionConfig {
	if in == nil {
		return nil
	}
	out := &krm.DataRetentionConfig{}
	out.AirflowMetadataRetentionConfig = AirflowMetadataRetentionPolicyConfig_FromProto(mapCtx, in.GetAirflowMetadataRetentionConfig())
	out.TaskLogsRetentionConfig = TaskLogsRetentionConfig_FromProto(mapCtx, in.GetTaskLogsRetentionConfig())
	return out
}
func DataRetentionConfig_ToProto(mapCtx *direct.MapContext, in *krm.DataRetentionConfig) *pb.DataRetentionConfig {
	if in == nil {
		return nil
	}
	out := &pb.DataRetentionConfig{}
	out.AirflowMetadataRetentionConfig = AirflowMetadataRetentionPolicyConfig_ToProto(mapCtx, in.AirflowMetadataRetentionConfig)
	out.TaskLogsRetentionConfig = TaskLogsRetentionConfig_ToProto(mapCtx, in.TaskLogsRetentionConfig)
	return out
}
func DatabaseConfig_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseConfig) *krm.DatabaseConfig {
	if in == nil {
		return nil
	}
	out := &krm.DatabaseConfig{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.Zone = direct.LazyPtr(in.GetZone())
	return out
}
func DatabaseConfig_ToProto(mapCtx *direct.MapContext, in *krm.DatabaseConfig) *pb.DatabaseConfig {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseConfig{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.Zone = direct.ValueOf(in.Zone)
	return out
}
func EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfig{}
	// MISSING: KMSKeyName
	return out
}
func EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfig) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	// MISSING: KMSKeyName
	return out
}
func EnvironmentConfig_FromProto(mapCtx *direct.MapContext, in *pb.EnvironmentConfig) *krm.EnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &krm.EnvironmentConfig{}
	// MISSING: GkeCluster
	// MISSING: DagGcsPrefix
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.SoftwareConfig = SoftwareConfig_FromProto(mapCtx, in.GetSoftwareConfig())
	out.NodeConfig = NodeConfig_FromProto(mapCtx, in.GetNodeConfig())
	out.PrivateEnvironmentConfig = PrivateEnvironmentConfig_FromProto(mapCtx, in.GetPrivateEnvironmentConfig())
	out.WebServerNetworkAccessControl = WebServerNetworkAccessControl_FromProto(mapCtx, in.GetWebServerNetworkAccessControl())
	out.DatabaseConfig = DatabaseConfig_FromProto(mapCtx, in.GetDatabaseConfig())
	out.WebServerConfig = WebServerConfig_FromProto(mapCtx, in.GetWebServerConfig())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	out.MaintenanceWindow = MaintenanceWindow_FromProto(mapCtx, in.GetMaintenanceWindow())
	out.WorkloadsConfig = WorkloadsConfig_FromProto(mapCtx, in.GetWorkloadsConfig())
	out.EnvironmentSize = direct.Enum_FromProto(mapCtx, in.GetEnvironmentSize())
	// MISSING: AirflowURI
	// MISSING: AirflowByoidURI
	out.MasterAuthorizedNetworksConfig = MasterAuthorizedNetworksConfig_FromProto(mapCtx, in.GetMasterAuthorizedNetworksConfig())
	out.RecoveryConfig = RecoveryConfig_FromProto(mapCtx, in.GetRecoveryConfig())
	out.ResilienceMode = direct.Enum_FromProto(mapCtx, in.GetResilienceMode())
	out.DataRetentionConfig = DataRetentionConfig_FromProto(mapCtx, in.GetDataRetentionConfig())
	return out
}
func EnvironmentConfig_ToProto(mapCtx *direct.MapContext, in *krm.EnvironmentConfig) *pb.EnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &pb.EnvironmentConfig{}
	// MISSING: GkeCluster
	// MISSING: DagGcsPrefix
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.SoftwareConfig = SoftwareConfig_ToProto(mapCtx, in.SoftwareConfig)
	out.NodeConfig = NodeConfig_ToProto(mapCtx, in.NodeConfig)
	out.PrivateEnvironmentConfig = PrivateEnvironmentConfig_ToProto(mapCtx, in.PrivateEnvironmentConfig)
	out.WebServerNetworkAccessControl = WebServerNetworkAccessControl_ToProto(mapCtx, in.WebServerNetworkAccessControl)
	out.DatabaseConfig = DatabaseConfig_ToProto(mapCtx, in.DatabaseConfig)
	out.WebServerConfig = WebServerConfig_ToProto(mapCtx, in.WebServerConfig)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	out.MaintenanceWindow = MaintenanceWindow_ToProto(mapCtx, in.MaintenanceWindow)
	out.WorkloadsConfig = WorkloadsConfig_ToProto(mapCtx, in.WorkloadsConfig)
	out.EnvironmentSize = direct.Enum_ToProto[pb.EnvironmentConfig_EnvironmentSize](mapCtx, in.EnvironmentSize)
	// MISSING: AirflowURI
	// MISSING: AirflowByoidURI
	out.MasterAuthorizedNetworksConfig = MasterAuthorizedNetworksConfig_ToProto(mapCtx, in.MasterAuthorizedNetworksConfig)
	out.RecoveryConfig = RecoveryConfig_ToProto(mapCtx, in.RecoveryConfig)
	out.ResilienceMode = direct.Enum_ToProto[pb.EnvironmentConfig_ResilienceMode](mapCtx, in.ResilienceMode)
	out.DataRetentionConfig = DataRetentionConfig_ToProto(mapCtx, in.DataRetentionConfig)
	return out
}
func EnvironmentConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EnvironmentConfig) *krm.EnvironmentConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EnvironmentConfigObservedState{}
	out.GkeCluster = direct.LazyPtr(in.GetGkeCluster())
	out.DagGcsPrefix = direct.LazyPtr(in.GetDagGcsPrefix())
	// MISSING: NodeCount
	// MISSING: SoftwareConfig
	// MISSING: NodeConfig
	out.PrivateEnvironmentConfig = PrivateEnvironmentConfigObservedState_FromProto(mapCtx, in.GetPrivateEnvironmentConfig())
	// MISSING: WebServerNetworkAccessControl
	// MISSING: DatabaseConfig
	// MISSING: WebServerConfig
	// MISSING: EncryptionConfig
	// MISSING: MaintenanceWindow
	// MISSING: WorkloadsConfig
	// MISSING: EnvironmentSize
	out.AirflowURI = direct.LazyPtr(in.GetAirflowUri())
	out.AirflowByoidURI = direct.LazyPtr(in.GetAirflowByoidUri())
	// MISSING: MasterAuthorizedNetworksConfig
	// MISSING: RecoveryConfig
	// MISSING: ResilienceMode
	// MISSING: DataRetentionConfig
	return out
}
func EnvironmentConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EnvironmentConfigObservedState) *pb.EnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &pb.EnvironmentConfig{}
	out.GkeCluster = direct.ValueOf(in.GkeCluster)
	out.DagGcsPrefix = direct.ValueOf(in.DagGcsPrefix)
	// MISSING: NodeCount
	// MISSING: SoftwareConfig
	// MISSING: NodeConfig
	out.PrivateEnvironmentConfig = PrivateEnvironmentConfigObservedState_ToProto(mapCtx, in.PrivateEnvironmentConfig)
	// MISSING: WebServerNetworkAccessControl
	// MISSING: DatabaseConfig
	// MISSING: WebServerConfig
	// MISSING: EncryptionConfig
	// MISSING: MaintenanceWindow
	// MISSING: WorkloadsConfig
	// MISSING: EnvironmentSize
	out.AirflowUri = direct.ValueOf(in.AirflowURI)
	out.AirflowByoidUri = direct.ValueOf(in.AirflowByoidURI)
	// MISSING: MasterAuthorizedNetworksConfig
	// MISSING: RecoveryConfig
	// MISSING: ResilienceMode
	// MISSING: DataRetentionConfig
	return out
}
func MaintenanceWindow_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceWindow) *krm.MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceWindow{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.Recurrence = direct.LazyPtr(in.GetRecurrence())
	return out
}
func MaintenanceWindow_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceWindow) *pb.MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceWindow{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.Recurrence = direct.ValueOf(in.Recurrence)
	return out
}
func MasterAuthorizedNetworksConfig_CIDRBlock_FromProto(mapCtx *direct.MapContext, in *pb.MasterAuthorizedNetworksConfig_CidrBlock) *krm.MasterAuthorizedNetworksConfig_CIDRBlock {
	if in == nil {
		return nil
	}
	out := &krm.MasterAuthorizedNetworksConfig_CIDRBlock{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.CIDRBlock = direct.LazyPtr(in.GetCidrBlock())
	return out
}
func MasterAuthorizedNetworksConfig_CIDRBlock_ToProto(mapCtx *direct.MapContext, in *krm.MasterAuthorizedNetworksConfig_CIDRBlock) *pb.MasterAuthorizedNetworksConfig_CidrBlock {
	if in == nil {
		return nil
	}
	out := &pb.MasterAuthorizedNetworksConfig_CidrBlock{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.CidrBlock = direct.ValueOf(in.CIDRBlock)
	return out
}
func NetworkingConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkingConfig) *krm.NetworkingConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkingConfig{}
	out.ConnectionType = direct.Enum_FromProto(mapCtx, in.GetConnectionType())
	return out
}
func NetworkingConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkingConfig) *pb.NetworkingConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkingConfig{}
	out.ConnectionType = direct.Enum_ToProto[pb.NetworkingConfig_ConnectionType](mapCtx, in.ConnectionType)
	return out
}
func NodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeConfig) *krm.NodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeConfig{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	if in.GetNetwork() != "" {
		out.NetworkRef = &refs.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &refs.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	out.OauthScopes = in.OauthScopes
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refs.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.Tags = in.Tags
	out.IPAllocationPolicy = IPAllocationPolicy_FromProto(mapCtx, in.GetIpAllocationPolicy())
	out.EnableIPMasqAgent = direct.LazyPtr(in.GetEnableIpMasqAgent())
	if in.GetComposerNetworkAttachment() != "" {
		out.ComposerNetworkAttachmentRef = &refs.ComputeNetworkAttachmentRef{External: in.GetComposerNetworkAttachment()}
	}
	out.ComposerInternalIPV4CIDRBlock = direct.LazyPtr(in.GetComposerInternalIpv4CidrBlock())
	return out
}
func NodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeConfig) *pb.NodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeConfig{}
	out.Location = direct.ValueOf(in.Location)
	out.MachineType = direct.ValueOf(in.MachineType)
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	out.OauthScopes = in.OauthScopes
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.Tags = in.Tags
	out.IpAllocationPolicy = IPAllocationPolicy_ToProto(mapCtx, in.IPAllocationPolicy)
	out.EnableIpMasqAgent = direct.ValueOf(in.EnableIPMasqAgent)
	if in.ComposerNetworkAttachmentRef != nil {
		out.ComposerNetworkAttachment = in.ComposerNetworkAttachmentRef.External
	}
	out.ComposerInternalIpv4CidrBlock = direct.ValueOf(in.ComposerInternalIPV4CIDRBlock)
	return out
}
func PrivateClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.PrivateClusterConfig) *krm.PrivateClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.PrivateClusterConfig{}
	out.EnablePrivateEndpoint = direct.LazyPtr(in.GetEnablePrivateEndpoint())
	out.MasterIPV4CIDRBlock = direct.LazyPtr(in.GetMasterIpv4CidrBlock())
	// MISSING: MasterIPV4ReservedRange
	return out
}
func PrivateClusterConfig_ToProto(mapCtx *direct.MapContext, in *krm.PrivateClusterConfig) *pb.PrivateClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivateClusterConfig{}
	out.EnablePrivateEndpoint = direct.ValueOf(in.EnablePrivateEndpoint)
	out.MasterIpv4CidrBlock = direct.ValueOf(in.MasterIPV4CIDRBlock)
	// MISSING: MasterIPV4ReservedRange
	return out
}
func PrivateClusterConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateClusterConfig) *krm.PrivateClusterConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivateClusterConfigObservedState{}
	// MISSING: EnablePrivateEndpoint
	// MISSING: MasterIPV4CIDRBlock
	out.MasterIPV4ReservedRange = direct.LazyPtr(in.GetMasterIpv4ReservedRange())
	return out
}
func PrivateClusterConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivateClusterConfigObservedState) *pb.PrivateClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivateClusterConfig{}
	// MISSING: EnablePrivateEndpoint
	// MISSING: MasterIPV4CIDRBlock
	out.MasterIpv4ReservedRange = direct.ValueOf(in.MasterIPV4ReservedRange)
	return out
}
func PrivateEnvironmentConfig_FromProto(mapCtx *direct.MapContext, in *pb.PrivateEnvironmentConfig) *krm.PrivateEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &krm.PrivateEnvironmentConfig{}
	out.EnablePrivateEnvironment = direct.LazyPtr(in.GetEnablePrivateEnvironment())
	out.EnablePrivateBuildsOnly = direct.LazyPtr(in.GetEnablePrivateBuildsOnly())
	out.PrivateClusterConfig = PrivateClusterConfig_FromProto(mapCtx, in.GetPrivateClusterConfig())
	out.WebServerIPV4CIDRBlock = direct.LazyPtr(in.GetWebServerIpv4CidrBlock())
	out.CloudSQLIPV4CIDRBlock = direct.LazyPtr(in.GetCloudSqlIpv4CidrBlock())
	// MISSING: WebServerIPV4ReservedRange
	out.CloudComposerNetworkIPV4CIDRBlock = direct.LazyPtr(in.GetCloudComposerNetworkIpv4CidrBlock())
	// MISSING: CloudComposerNetworkIPV4ReservedRange
	out.EnablePrivatelyUsedPublicIps = direct.LazyPtr(in.GetEnablePrivatelyUsedPublicIps())
	if in.GetCloudComposerConnectionSubnetwork() != "" {
		out.CloudComposerConnectionSubnetworkRef = &refs.ComputeSubnetworkRef{External: in.GetCloudComposerConnectionSubnetwork()}
	}
	out.NetworkingConfig = NetworkingConfig_FromProto(mapCtx, in.GetNetworkingConfig())
	return out
}
func PrivateEnvironmentConfig_ToProto(mapCtx *direct.MapContext, in *krm.PrivateEnvironmentConfig) *pb.PrivateEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivateEnvironmentConfig{}
	out.EnablePrivateEnvironment = direct.ValueOf(in.EnablePrivateEnvironment)
	out.EnablePrivateBuildsOnly = direct.ValueOf(in.EnablePrivateBuildsOnly)
	out.PrivateClusterConfig = PrivateClusterConfig_ToProto(mapCtx, in.PrivateClusterConfig)
	out.WebServerIpv4CidrBlock = direct.ValueOf(in.WebServerIPV4CIDRBlock)
	out.CloudSqlIpv4CidrBlock = direct.ValueOf(in.CloudSQLIPV4CIDRBlock)
	// MISSING: WebServerIPV4ReservedRange
	out.CloudComposerNetworkIpv4CidrBlock = direct.ValueOf(in.CloudComposerNetworkIPV4CIDRBlock)
	// MISSING: CloudComposerNetworkIPV4ReservedRange
	out.EnablePrivatelyUsedPublicIps = direct.ValueOf(in.EnablePrivatelyUsedPublicIps)
	if in.CloudComposerConnectionSubnetworkRef != nil {
		out.CloudComposerConnectionSubnetwork = in.CloudComposerConnectionSubnetworkRef.External
	}
	out.NetworkingConfig = NetworkingConfig_ToProto(mapCtx, in.NetworkingConfig)
	return out
}
func PrivateEnvironmentConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateEnvironmentConfig) *krm.PrivateEnvironmentConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivateEnvironmentConfigObservedState{}
	// MISSING: EnablePrivateEnvironment
	// MISSING: EnablePrivateBuildsOnly
	out.PrivateClusterConfig = PrivateClusterConfigObservedState_FromProto(mapCtx, in.GetPrivateClusterConfig())
	// MISSING: WebServerIPV4CIDRBlock
	// MISSING: CloudSQLIPV4CIDRBlock
	out.WebServerIPV4ReservedRange = direct.LazyPtr(in.GetWebServerIpv4ReservedRange())
	// MISSING: CloudComposerNetworkIPV4CIDRBlock
	out.CloudComposerNetworkIPV4ReservedRange = direct.LazyPtr(in.GetCloudComposerNetworkIpv4ReservedRange())
	// MISSING: EnablePrivatelyUsedPublicIps
	// MISSING: CloudComposerConnectionSubnetwork
	// MISSING: NetworkingConfig
	return out
}
func PrivateEnvironmentConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivateEnvironmentConfigObservedState) *pb.PrivateEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivateEnvironmentConfig{}
	// MISSING: EnablePrivateEnvironment
	// MISSING: EnablePrivateBuildsOnly
	out.PrivateClusterConfig = PrivateClusterConfigObservedState_ToProto(mapCtx, in.PrivateClusterConfig)
	// MISSING: WebServerIPV4CIDRBlock
	// MISSING: CloudSQLIPV4CIDRBlock
	out.WebServerIpv4ReservedRange = direct.ValueOf(in.WebServerIPV4ReservedRange)
	// MISSING: CloudComposerNetworkIPV4CIDRBlock
	out.CloudComposerNetworkIpv4ReservedRange = direct.ValueOf(in.CloudComposerNetworkIPV4ReservedRange)
	// MISSING: EnablePrivatelyUsedPublicIps
	// MISSING: CloudComposerConnectionSubnetwork
	// MISSING: NetworkingConfig
	return out
}
func RecoveryConfig_FromProto(mapCtx *direct.MapContext, in *pb.RecoveryConfig) *krm.RecoveryConfig {
	if in == nil {
		return nil
	}
	out := &krm.RecoveryConfig{}
	out.ScheduledSnapshotsConfig = ScheduledSnapshotsConfig_FromProto(mapCtx, in.GetScheduledSnapshotsConfig())
	return out
}
func RecoveryConfig_ToProto(mapCtx *direct.MapContext, in *krm.RecoveryConfig) *pb.RecoveryConfig {
	if in == nil {
		return nil
	}
	out := &pb.RecoveryConfig{}
	out.ScheduledSnapshotsConfig = ScheduledSnapshotsConfig_ToProto(mapCtx, in.ScheduledSnapshotsConfig)
	return out
}
func ScheduledSnapshotsConfig_FromProto(mapCtx *direct.MapContext, in *pb.ScheduledSnapshotsConfig) *krm.ScheduledSnapshotsConfig {
	if in == nil {
		return nil
	}
	out := &krm.ScheduledSnapshotsConfig{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.SnapshotLocation = direct.LazyPtr(in.GetSnapshotLocation())
	out.SnapshotCreationSchedule = direct.LazyPtr(in.GetSnapshotCreationSchedule())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	return out
}
func ScheduledSnapshotsConfig_ToProto(mapCtx *direct.MapContext, in *krm.ScheduledSnapshotsConfig) *pb.ScheduledSnapshotsConfig {
	if in == nil {
		return nil
	}
	out := &pb.ScheduledSnapshotsConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.SnapshotLocation = direct.ValueOf(in.SnapshotLocation)
	out.SnapshotCreationSchedule = direct.ValueOf(in.SnapshotCreationSchedule)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	return out
}
func SoftwareConfig_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareConfig) *krm.SoftwareConfig {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareConfig{}
	out.ImageVersion = direct.LazyPtr(in.GetImageVersion())
	out.AirflowConfigOverrides = in.AirflowConfigOverrides
	out.PypiPackages = in.PypiPackages
	out.EnvVariables = in.EnvVariables
	out.PythonVersion = direct.LazyPtr(in.GetPythonVersion())
	out.SchedulerCount = direct.LazyPtr(in.GetSchedulerCount())
	out.CloudDataLineageIntegration = CloudDataLineageIntegration_FromProto(mapCtx, in.GetCloudDataLineageIntegration())
	out.WebServerPluginsMode = direct.Enum_FromProto(mapCtx, in.GetWebServerPluginsMode())
	return out
}
func SoftwareConfig_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareConfig) *pb.SoftwareConfig {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareConfig{}
	out.ImageVersion = direct.ValueOf(in.ImageVersion)
	out.AirflowConfigOverrides = in.AirflowConfigOverrides
	out.PypiPackages = in.PypiPackages
	out.EnvVariables = in.EnvVariables
	out.PythonVersion = direct.ValueOf(in.PythonVersion)
	out.SchedulerCount = direct.ValueOf(in.SchedulerCount)
	out.CloudDataLineageIntegration = CloudDataLineageIntegration_ToProto(mapCtx, in.CloudDataLineageIntegration)
	out.WebServerPluginsMode = direct.Enum_ToProto[pb.SoftwareConfig_WebServerPluginsMode](mapCtx, in.WebServerPluginsMode)
	return out
}
func StorageConfig_FromProto(mapCtx *direct.MapContext, in *pb.StorageConfig) *krm.StorageConfig {
	if in == nil {
		return nil
	}
	out := &krm.StorageConfig{}
	if in.GetBucket() != "" {
		out.BucketRef = &refs.StorageBucketRef{External: in.GetBucket()}
	}
	return out
}
func StorageConfig_ToProto(mapCtx *direct.MapContext, in *krm.StorageConfig) *pb.StorageConfig {
	if in == nil {
		return nil
	}
	out := &pb.StorageConfig{}
	if in.BucketRef != nil {
		out.Bucket = in.BucketRef.External
	}
	return out
}
func TaskLogsRetentionConfig_FromProto(mapCtx *direct.MapContext, in *pb.TaskLogsRetentionConfig) *krm.TaskLogsRetentionConfig {
	if in == nil {
		return nil
	}
	out := &krm.TaskLogsRetentionConfig{}
	out.StorageMode = direct.Enum_FromProto(mapCtx, in.GetStorageMode())
	return out
}
func TaskLogsRetentionConfig_ToProto(mapCtx *direct.MapContext, in *krm.TaskLogsRetentionConfig) *pb.TaskLogsRetentionConfig {
	if in == nil {
		return nil
	}
	out := &pb.TaskLogsRetentionConfig{}
	out.StorageMode = direct.Enum_ToProto[pb.TaskLogsRetentionConfig_TaskLogsStorageMode](mapCtx, in.StorageMode)
	return out
}
func WebServerConfig_FromProto(mapCtx *direct.MapContext, in *pb.WebServerConfig) *krm.WebServerConfig {
	if in == nil {
		return nil
	}
	out := &krm.WebServerConfig{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	return out
}
func WebServerConfig_ToProto(mapCtx *direct.MapContext, in *krm.WebServerConfig) *pb.WebServerConfig {
	if in == nil {
		return nil
	}
	out := &pb.WebServerConfig{}
	out.MachineType = direct.ValueOf(in.MachineType)
	return out
}
func WebServerNetworkAccessControl_AllowedIPRange_FromProto(mapCtx *direct.MapContext, in *pb.WebServerNetworkAccessControl_AllowedIpRange) *krm.WebServerNetworkAccessControl_AllowedIPRange {
	if in == nil {
		return nil
	}
	out := &krm.WebServerNetworkAccessControl_AllowedIPRange{}
	out.Value = direct.LazyPtr(in.GetValue())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func WebServerNetworkAccessControl_AllowedIPRange_ToProto(mapCtx *direct.MapContext, in *krm.WebServerNetworkAccessControl_AllowedIPRange) *pb.WebServerNetworkAccessControl_AllowedIpRange {
	if in == nil {
		return nil
	}
	out := &pb.WebServerNetworkAccessControl_AllowedIpRange{}
	out.Value = direct.ValueOf(in.Value)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func WorkloadsConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadsConfig) *krm.WorkloadsConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadsConfig{}
	out.Scheduler = WorkloadsConfig_SchedulerResource_FromProto(mapCtx, in.GetScheduler())
	out.WebServer = WorkloadsConfig_WebServerResource_FromProto(mapCtx, in.GetWebServer())
	out.Worker = WorkloadsConfig_WorkerResource_FromProto(mapCtx, in.GetWorker())
	out.Triggerer = WorkloadsConfig_TriggererResource_FromProto(mapCtx, in.GetTriggerer())
	out.DagProcessor = WorkloadsConfig_DagProcessorResource_FromProto(mapCtx, in.GetDagProcessor())
	return out
}
func WorkloadsConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadsConfig) *pb.WorkloadsConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadsConfig{}
	out.Scheduler = WorkloadsConfig_SchedulerResource_ToProto(mapCtx, in.Scheduler)
	out.WebServer = WorkloadsConfig_WebServerResource_ToProto(mapCtx, in.WebServer)
	out.Worker = WorkloadsConfig_WorkerResource_ToProto(mapCtx, in.Worker)
	out.Triggerer = WorkloadsConfig_TriggererResource_ToProto(mapCtx, in.Triggerer)
	out.DagProcessor = WorkloadsConfig_DagProcessorResource_ToProto(mapCtx, in.DagProcessor)
	return out
}
func WorkloadsConfig_DagProcessorResource_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadsConfig_DagProcessorResource) *krm.WorkloadsConfig_DagProcessorResource {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadsConfig_DagProcessorResource{}
	out.CPU = direct.LazyPtr(in.GetCpu())
	out.MemoryGB = direct.LazyPtr(in.GetMemoryGb())
	out.StorageGB = direct.LazyPtr(in.GetStorageGb())
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func WorkloadsConfig_DagProcessorResource_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadsConfig_DagProcessorResource) *pb.WorkloadsConfig_DagProcessorResource {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadsConfig_DagProcessorResource{}
	out.Cpu = direct.ValueOf(in.CPU)
	out.MemoryGb = direct.ValueOf(in.MemoryGB)
	out.StorageGb = direct.ValueOf(in.StorageGB)
	out.Count = direct.ValueOf(in.Count)
	return out
}
func WorkloadsConfig_SchedulerResource_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadsConfig_SchedulerResource) *krm.WorkloadsConfig_SchedulerResource {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadsConfig_SchedulerResource{}
	out.CPU = direct.LazyPtr(in.GetCpu())
	out.MemoryGB = direct.LazyPtr(in.GetMemoryGb())
	out.StorageGB = direct.LazyPtr(in.GetStorageGb())
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}
func WorkloadsConfig_SchedulerResource_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadsConfig_SchedulerResource) *pb.WorkloadsConfig_SchedulerResource {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadsConfig_SchedulerResource{}
	out.Cpu = direct.ValueOf(in.CPU)
	out.MemoryGb = direct.ValueOf(in.MemoryGB)
	out.StorageGb = direct.ValueOf(in.StorageGB)
	out.Count = direct.ValueOf(in.Count)
	return out
}
func WorkloadsConfig_TriggererResource_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadsConfig_TriggererResource) *krm.WorkloadsConfig_TriggererResource {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadsConfig_TriggererResource{}
	out.Count = direct.LazyPtr(in.GetCount())
	out.CPU = direct.LazyPtr(in.GetCpu())
	out.MemoryGB = direct.LazyPtr(in.GetMemoryGb())
	return out
}
func WorkloadsConfig_TriggererResource_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadsConfig_TriggererResource) *pb.WorkloadsConfig_TriggererResource {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadsConfig_TriggererResource{}
	out.Count = direct.ValueOf(in.Count)
	out.Cpu = direct.ValueOf(in.CPU)
	out.MemoryGb = direct.ValueOf(in.MemoryGB)
	return out
}
func WorkloadsConfig_WebServerResource_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadsConfig_WebServerResource) *krm.WorkloadsConfig_WebServerResource {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadsConfig_WebServerResource{}
	out.CPU = direct.LazyPtr(in.GetCpu())
	out.MemoryGB = direct.LazyPtr(in.GetMemoryGb())
	out.StorageGB = direct.LazyPtr(in.GetStorageGb())
	return out
}
func WorkloadsConfig_WebServerResource_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadsConfig_WebServerResource) *pb.WorkloadsConfig_WebServerResource {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadsConfig_WebServerResource{}
	out.Cpu = direct.ValueOf(in.CPU)
	out.MemoryGb = direct.ValueOf(in.MemoryGB)
	out.StorageGb = direct.ValueOf(in.StorageGB)
	return out
}
func WorkloadsConfig_WorkerResource_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadsConfig_WorkerResource) *krm.WorkloadsConfig_WorkerResource {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadsConfig_WorkerResource{}
	out.CPU = direct.LazyPtr(in.GetCpu())
	out.MemoryGB = direct.LazyPtr(in.GetMemoryGb())
	out.StorageGB = direct.LazyPtr(in.GetStorageGb())
	out.MinCount = direct.LazyPtr(in.GetMinCount())
	out.MaxCount = direct.LazyPtr(in.GetMaxCount())
	return out
}
func WorkloadsConfig_WorkerResource_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadsConfig_WorkerResource) *pb.WorkloadsConfig_WorkerResource {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadsConfig_WorkerResource{}
	out.Cpu = direct.ValueOf(in.CPU)
	out.MemoryGb = direct.ValueOf(in.MemoryGB)
	out.StorageGb = direct.ValueOf(in.StorageGB)
	out.MinCount = direct.ValueOf(in.MinCount)
	out.MaxCount = direct.ValueOf(in.MaxCount)
	return out
}
