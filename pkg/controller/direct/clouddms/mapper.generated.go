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
// krm.group: clouddms.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.clouddms.v1

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AlloyDbConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbConnectionProfile) *krm.AlloyDbConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbConnectionProfile{}
	out.ClusterID = direct.LazyPtr(in.GetClusterId())
	out.Settings = AlloyDbSettings_FromProto(mapCtx, in.GetSettings())
	return out
}
func AlloyDbConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbConnectionProfile) *pb.AlloyDbConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbConnectionProfile{}
	out.ClusterId = direct.ValueOf(in.ClusterID)
	out.Settings = AlloyDbSettings_ToProto(mapCtx, in.Settings)
	return out
}
func AlloyDbConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbConnectionProfile) *krm.AlloyDbConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbConnectionProfileObservedState{}
	// MISSING: ClusterID
	out.Settings = AlloyDbSettingsObservedState_FromProto(mapCtx, in.GetSettings())
	return out
}
func AlloyDbConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbConnectionProfileObservedState) *pb.AlloyDbConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbConnectionProfile{}
	// MISSING: ClusterID
	out.Settings = AlloyDbSettingsObservedState_ToProto(mapCtx, in.Settings)
	return out
}
func AlloyDbSettings_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbSettings) *krm.AlloyDbSettings {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbSettings{}
	out.InitialUser = AlloyDbSettings_UserPassword_FromProto(mapCtx, in.GetInitialUser())
	out.VPCNetwork = direct.LazyPtr(in.GetVpcNetwork())
	out.Labels = in.Labels
	out.PrimaryInstanceSettings = AlloyDbSettings_PrimaryInstanceSettings_FromProto(mapCtx, in.GetPrimaryInstanceSettings())
	out.EncryptionConfig = AlloyDbSettings_EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	return out
}
func AlloyDbSettings_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbSettings) *pb.AlloyDbSettings {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbSettings{}
	out.InitialUser = AlloyDbSettings_UserPassword_ToProto(mapCtx, in.InitialUser)
	out.VpcNetwork = direct.ValueOf(in.VPCNetwork)
	out.Labels = in.Labels
	out.PrimaryInstanceSettings = AlloyDbSettings_PrimaryInstanceSettings_ToProto(mapCtx, in.PrimaryInstanceSettings)
	out.EncryptionConfig = AlloyDbSettings_EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	return out
}
func AlloyDbSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbSettings) *krm.AlloyDbSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbSettingsObservedState{}
	out.InitialUser = AlloyDbSettings_UserPasswordObservedState_FromProto(mapCtx, in.GetInitialUser())
	// MISSING: VPCNetwork
	// MISSING: Labels
	out.PrimaryInstanceSettings = AlloyDbSettings_PrimaryInstanceSettingsObservedState_FromProto(mapCtx, in.GetPrimaryInstanceSettings())
	// MISSING: EncryptionConfig
	return out
}
func AlloyDbSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbSettingsObservedState) *pb.AlloyDbSettings {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbSettings{}
	out.InitialUser = AlloyDbSettings_UserPasswordObservedState_ToProto(mapCtx, in.InitialUser)
	// MISSING: VPCNetwork
	// MISSING: Labels
	out.PrimaryInstanceSettings = AlloyDbSettings_PrimaryInstanceSettingsObservedState_ToProto(mapCtx, in.PrimaryInstanceSettings)
	// MISSING: EncryptionConfig
	return out
}
func AlloyDbSettings_EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbSettings_EncryptionConfig) *krm.AlloyDbSettings_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbSettings_EncryptionConfig{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}
func AlloyDbSettings_EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbSettings_EncryptionConfig) *pb.AlloyDbSettings_EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbSettings_EncryptionConfig{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	return out
}
func AlloyDbSettings_PrimaryInstanceSettings_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbSettings_PrimaryInstanceSettings) *krm.AlloyDbSettings_PrimaryInstanceSettings {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbSettings_PrimaryInstanceSettings{}
	out.ID = direct.LazyPtr(in.GetId())
	out.MachineConfig = AlloyDbSettings_PrimaryInstanceSettings_MachineConfig_FromProto(mapCtx, in.GetMachineConfig())
	out.DatabaseFlags = in.DatabaseFlags
	out.Labels = in.Labels
	// MISSING: PrivateIP
	return out
}
func AlloyDbSettings_PrimaryInstanceSettings_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbSettings_PrimaryInstanceSettings) *pb.AlloyDbSettings_PrimaryInstanceSettings {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbSettings_PrimaryInstanceSettings{}
	out.Id = direct.ValueOf(in.ID)
	out.MachineConfig = AlloyDbSettings_PrimaryInstanceSettings_MachineConfig_ToProto(mapCtx, in.MachineConfig)
	out.DatabaseFlags = in.DatabaseFlags
	out.Labels = in.Labels
	// MISSING: PrivateIP
	return out
}
func AlloyDbSettings_PrimaryInstanceSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbSettings_PrimaryInstanceSettings) *krm.AlloyDbSettings_PrimaryInstanceSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbSettings_PrimaryInstanceSettingsObservedState{}
	// MISSING: ID
	// MISSING: MachineConfig
	// MISSING: DatabaseFlags
	// MISSING: Labels
	out.PrivateIP = direct.LazyPtr(in.GetPrivateIp())
	return out
}
func AlloyDbSettings_PrimaryInstanceSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbSettings_PrimaryInstanceSettingsObservedState) *pb.AlloyDbSettings_PrimaryInstanceSettings {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbSettings_PrimaryInstanceSettings{}
	// MISSING: ID
	// MISSING: MachineConfig
	// MISSING: DatabaseFlags
	// MISSING: Labels
	out.PrivateIp = direct.ValueOf(in.PrivateIP)
	return out
}
func AlloyDbSettings_PrimaryInstanceSettings_MachineConfig_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbSettings_PrimaryInstanceSettings_MachineConfig) *krm.AlloyDbSettings_PrimaryInstanceSettings_MachineConfig {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbSettings_PrimaryInstanceSettings_MachineConfig{}
	out.CPUCount = direct.LazyPtr(in.GetCpuCount())
	return out
}
func AlloyDbSettings_PrimaryInstanceSettings_MachineConfig_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbSettings_PrimaryInstanceSettings_MachineConfig) *pb.AlloyDbSettings_PrimaryInstanceSettings_MachineConfig {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbSettings_PrimaryInstanceSettings_MachineConfig{}
	out.CpuCount = direct.ValueOf(in.CPUCount)
	return out
}
func AlloyDbSettings_UserPassword_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbSettings_UserPassword) *krm.AlloyDbSettings_UserPassword {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbSettings_UserPassword{}
	out.User = direct.LazyPtr(in.GetUser())
	// MISSING: Password
	// MISSING: PasswordSet
	return out
}
func AlloyDbSettings_UserPassword_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbSettings_UserPassword) *pb.AlloyDbSettings_UserPassword {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbSettings_UserPassword{}
	out.User = direct.ValueOf(in.User)
	// MISSING: Password
	// MISSING: PasswordSet
	return out
}
func AlloyDbSettings_UserPasswordObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbSettings_UserPassword) *krm.AlloyDbSettings_UserPasswordObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbSettings_UserPasswordObservedState{}
	// MISSING: User
	// MISSING: Password
	out.PasswordSet = direct.LazyPtr(in.GetPasswordSet())
	return out
}
func AlloyDbSettings_UserPasswordObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbSettings_UserPasswordObservedState) *pb.AlloyDbSettings_UserPassword {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbSettings_UserPassword{}
	// MISSING: User
	// MISSING: Password
	out.PasswordSet = direct.ValueOf(in.PasswordSet)
	return out
}
func CloudDMSConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.CloudDMSConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSConnectionProfileObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	return out
}
func CloudDMSConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSConnectionProfileObservedState) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.ConnectionProfile_State](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	return out
}
func CloudDMSConnectionProfileSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.CloudDMSConnectionProfileSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSConnectionProfileSpec{}
	// MISSING: Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Mysql = MySQLConnectionProfile_FromProto(mapCtx, in.GetMysql())
	out.Postgresql = PostgreSQLConnectionProfile_FromProto(mapCtx, in.GetPostgresql())
	out.Oracle = OracleConnectionProfile_FromProto(mapCtx, in.GetOracle())
	out.Cloudsql = CloudSQLConnectionProfile_FromProto(mapCtx, in.GetCloudsql())
	out.Alloydb = AlloyDbConnectionProfile_FromProto(mapCtx, in.GetAlloydb())
	out.Provider = direct.Enum_FromProto(mapCtx, in.GetProvider())
	return out
}
func CloudDMSConnectionProfileSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSConnectionProfileSpec) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if oneof := MySQLConnectionProfile_ToProto(mapCtx, in.Mysql); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Mysql{Mysql: oneof}
	}
	if oneof := PostgreSQLConnectionProfile_ToProto(mapCtx, in.Postgresql); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Postgresql{Postgresql: oneof}
	}
	if oneof := OracleConnectionProfile_ToProto(mapCtx, in.Oracle); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Oracle{Oracle: oneof}
	}
	if oneof := CloudSQLConnectionProfile_ToProto(mapCtx, in.Cloudsql); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Cloudsql{Cloudsql: oneof}
	}
	if oneof := AlloyDbConnectionProfile_ToProto(mapCtx, in.Alloydb); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Alloydb{Alloydb: oneof}
	}
	out.Provider = direct.Enum_ToProto[pb.DatabaseProvider](mapCtx, in.Provider)
	return out
}
func CloudDMSConversionWorkspaceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspace) *krm.CloudDMSConversionWorkspaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSConversionWorkspaceObservedState{}
	// MISSING: Name
	out.HasUncommittedChanges = direct.LazyPtr(in.GetHasUncommittedChanges())
	out.LatestCommitID = direct.LazyPtr(in.GetLatestCommitId())
	out.LatestCommitTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLatestCommitTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func CloudDMSConversionWorkspaceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSConversionWorkspaceObservedState) *pb.ConversionWorkspace {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspace{}
	// MISSING: Name
	out.HasUncommittedChanges = direct.ValueOf(in.HasUncommittedChanges)
	out.LatestCommitId = direct.ValueOf(in.LatestCommitID)
	out.LatestCommitTime = direct.StringTimestamp_ToProto(mapCtx, in.LatestCommitTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func CloudDMSConversionWorkspaceSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspace) *krm.CloudDMSConversionWorkspaceSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSConversionWorkspaceSpec{}
	// MISSING: Name
	out.Source = DatabaseEngineInfo_FromProto(mapCtx, in.GetSource())
	out.Destination = DatabaseEngineInfo_FromProto(mapCtx, in.GetDestination())
	out.GlobalSettings = in.GlobalSettings
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func CloudDMSConversionWorkspaceSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSConversionWorkspaceSpec) *pb.ConversionWorkspace {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspace{}
	// MISSING: Name
	out.Source = DatabaseEngineInfo_ToProto(mapCtx, in.Source)
	out.Destination = DatabaseEngineInfo_ToProto(mapCtx, in.Destination)
	out.GlobalSettings = in.GlobalSettings
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func CloudDMSMigrationJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob) *krm.CloudDMSMigrationJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSMigrationJobObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Phase
	// MISSING: Duration
	// MISSING: Error
	// MISSING: EndTime
	return out
}
func CloudDMSMigrationJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSMigrationJobObservedState) *pb.MigrationJob {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Phase
	// MISSING: Duration
	// MISSING: Error
	// MISSING: EndTime
	return out
}
func CloudDMSMigrationJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob) *krm.CloudDMSMigrationJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSMigrationJobSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: State
	// MISSING: Phase
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.DumpPath = direct.LazyPtr(in.GetDumpPath())
	out.DumpFlags = MigrationJob_DumpFlags_FromProto(mapCtx, in.GetDumpFlags())
	if in.GetSource() != "" {
		out.SourceRef = &krm.CloudDMSConnectionProfileRef{External: in.GetSource()}
	}
	if in.GetDestination() != "" {
		out.DestinationRef = &krm.CloudDMSConnectionProfileRef{External: in.GetDestination()}
	}
	out.ReverseSSHConnectivity = ReverseSSHConnectivity_FromProto(mapCtx, in.GetReverseSshConnectivity())
	out.VPCPeeringConnectivity = VPCPeeringConnectivity_FromProto(mapCtx, in.GetVpcPeeringConnectivity())
	out.StaticIPConnectivity = StaticIPConnectivity_FromProto(mapCtx, in.GetStaticIpConnectivity())
	// MISSING: Duration
	// MISSING: Error
	out.SourceDatabase = DatabaseType_FromProto(mapCtx, in.GetSourceDatabase())
	out.DestinationDatabase = DatabaseType_FromProto(mapCtx, in.GetDestinationDatabase())
	// MISSING: EndTime
	out.ConversionWorkspace = ConversionWorkspaceInfo_FromProto(mapCtx, in.GetConversionWorkspace())
	out.Filter = direct.LazyPtr(in.GetFilter())
	if in.GetCmekKeyName() != "" {
		out.CmekKeyNameRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetCmekKeyName()}
	}
	out.PerformanceConfig = MigrationJob_PerformanceConfig_FromProto(mapCtx, in.GetPerformanceConfig())
	return out
}
func CloudDMSMigrationJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSMigrationJobSpec) *pb.MigrationJob {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: State
	// MISSING: Phase
	out.Type = direct.Enum_ToProto[pb.MigrationJob_Type](mapCtx, in.Type)
	out.DumpPath = direct.ValueOf(in.DumpPath)
	out.DumpFlags = MigrationJob_DumpFlags_ToProto(mapCtx, in.DumpFlags)
	if in.SourceRef != nil {
		out.Source = in.SourceRef.External
	}
	if in.DestinationRef != nil {
		out.Destination = in.DestinationRef.External
	}
	if oneof := ReverseSSHConnectivity_ToProto(mapCtx, in.ReverseSSHConnectivity); oneof != nil {
		out.Connectivity = &pb.MigrationJob_ReverseSshConnectivity{ReverseSshConnectivity: oneof}
	}
	if oneof := VPCPeeringConnectivity_ToProto(mapCtx, in.VPCPeeringConnectivity); oneof != nil {
		out.Connectivity = &pb.MigrationJob_VpcPeeringConnectivity{VpcPeeringConnectivity: oneof}
	}
	if oneof := StaticIPConnectivity_ToProto(mapCtx, in.StaticIPConnectivity); oneof != nil {
		out.Connectivity = &pb.MigrationJob_StaticIpConnectivity{StaticIpConnectivity: oneof}
	}
	// MISSING: Duration
	// MISSING: Error
	out.SourceDatabase = DatabaseType_ToProto(mapCtx, in.SourceDatabase)
	out.DestinationDatabase = DatabaseType_ToProto(mapCtx, in.DestinationDatabase)
	// MISSING: EndTime
	out.ConversionWorkspace = ConversionWorkspaceInfo_ToProto(mapCtx, in.ConversionWorkspace)
	out.Filter = direct.ValueOf(in.Filter)
	if in.CmekKeyNameRef != nil {
		out.CmekKeyName = in.CmekKeyNameRef.External
	}
	out.PerformanceConfig = MigrationJob_PerformanceConfig_ToProto(mapCtx, in.PerformanceConfig)
	return out
}
func CloudSQLConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlConnectionProfile) *krm.CloudSQLConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.CloudSQLConnectionProfile{}
	// MISSING: CloudSQLID
	out.Settings = CloudSQLSettings_FromProto(mapCtx, in.GetSettings())
	// MISSING: PrivateIP
	// MISSING: PublicIP
	// MISSING: AdditionalPublicIP
	return out
}
func CloudSQLConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.CloudSQLConnectionProfile) *pb.CloudSqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlConnectionProfile{}
	// MISSING: CloudSQLID
	out.Settings = CloudSQLSettings_ToProto(mapCtx, in.Settings)
	// MISSING: PrivateIP
	// MISSING: PublicIP
	// MISSING: AdditionalPublicIP
	return out
}
func CloudSQLConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlConnectionProfile) *krm.CloudSQLConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudSQLConnectionProfileObservedState{}
	out.CloudSQLID = direct.LazyPtr(in.GetCloudSqlId())
	out.Settings = CloudSQLSettingsObservedState_FromProto(mapCtx, in.GetSettings())
	out.PrivateIP = direct.LazyPtr(in.GetPrivateIp())
	out.PublicIP = direct.LazyPtr(in.GetPublicIp())
	out.AdditionalPublicIP = direct.LazyPtr(in.GetAdditionalPublicIp())
	return out
}
func CloudSQLConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudSQLConnectionProfileObservedState) *pb.CloudSqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlConnectionProfile{}
	out.CloudSqlId = direct.ValueOf(in.CloudSQLID)
	out.Settings = CloudSQLSettingsObservedState_ToProto(mapCtx, in.Settings)
	out.PrivateIp = direct.ValueOf(in.PrivateIP)
	out.PublicIp = direct.ValueOf(in.PublicIP)
	out.AdditionalPublicIp = direct.ValueOf(in.AdditionalPublicIP)
	return out
}
func CloudSQLSettings_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlSettings) *krm.CloudSQLSettings {
	if in == nil {
		return nil
	}
	out := &krm.CloudSQLSettings{}
	out.DatabaseVersion = direct.Enum_FromProto(mapCtx, in.GetDatabaseVersion())
	out.UserLabels = in.UserLabels
	out.Tier = direct.LazyPtr(in.GetTier())
	out.StorageAutoResizeLimit = direct.Int64Value_FromProto(mapCtx, in.GetStorageAutoResizeLimit())
	out.ActivationPolicy = direct.Enum_FromProto(mapCtx, in.GetActivationPolicy())
	out.IPConfig = SQLIPConfig_FromProto(mapCtx, in.GetIpConfig())
	out.AutoStorageIncrease = direct.BoolValue_FromProto(mapCtx, in.GetAutoStorageIncrease())
	out.DatabaseFlags = in.DatabaseFlags
	out.DataDiskType = direct.Enum_FromProto(mapCtx, in.GetDataDiskType())
	out.DataDiskSizeGB = direct.Int64Value_FromProto(mapCtx, in.GetDataDiskSizeGb())
	out.Zone = direct.LazyPtr(in.GetZone())
	out.SecondaryZone = direct.LazyPtr(in.GetSecondaryZone())
	out.SourceID = direct.LazyPtr(in.GetSourceId())
	// MISSING: RootPassword
	// MISSING: RootPasswordSet
	out.Collation = direct.LazyPtr(in.GetCollation())
	out.CmekKeyName = direct.LazyPtr(in.GetCmekKeyName())
	out.AvailabilityType = direct.Enum_FromProto(mapCtx, in.GetAvailabilityType())
	out.Edition = direct.Enum_FromProto(mapCtx, in.GetEdition())
	return out
}
func CloudSQLSettings_ToProto(mapCtx *direct.MapContext, in *krm.CloudSQLSettings) *pb.CloudSqlSettings {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlSettings{}
	out.DatabaseVersion = direct.Enum_ToProto[pb.CloudSqlSettings_SqlDatabaseVersion](mapCtx, in.DatabaseVersion)
	out.UserLabels = in.UserLabels
	out.Tier = direct.ValueOf(in.Tier)
	out.StorageAutoResizeLimit = direct.Int64Value_ToProto(mapCtx, in.StorageAutoResizeLimit)
	out.ActivationPolicy = direct.Enum_ToProto[pb.CloudSqlSettings_SqlActivationPolicy](mapCtx, in.ActivationPolicy)
	out.IpConfig = SQLIPConfig_ToProto(mapCtx, in.IPConfig)
	out.AutoStorageIncrease = direct.BoolValue_ToProto(mapCtx, in.AutoStorageIncrease)
	out.DatabaseFlags = in.DatabaseFlags
	out.DataDiskType = direct.Enum_ToProto[pb.CloudSqlSettings_SqlDataDiskType](mapCtx, in.DataDiskType)
	out.DataDiskSizeGb = direct.Int64Value_ToProto(mapCtx, in.DataDiskSizeGB)
	out.Zone = direct.ValueOf(in.Zone)
	out.SecondaryZone = direct.ValueOf(in.SecondaryZone)
	out.SourceId = direct.ValueOf(in.SourceID)
	// MISSING: RootPassword
	// MISSING: RootPasswordSet
	out.Collation = direct.ValueOf(in.Collation)
	out.CmekKeyName = direct.ValueOf(in.CmekKeyName)
	out.AvailabilityType = direct.Enum_ToProto[pb.CloudSqlSettings_SqlAvailabilityType](mapCtx, in.AvailabilityType)
	out.Edition = direct.Enum_ToProto[pb.CloudSqlSettings_Edition](mapCtx, in.Edition)
	return out
}
func CloudSQLSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlSettings) *krm.CloudSQLSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudSQLSettingsObservedState{}
	// MISSING: DatabaseVersion
	// MISSING: UserLabels
	// MISSING: Tier
	// MISSING: StorageAutoResizeLimit
	// MISSING: ActivationPolicy
	// MISSING: IPConfig
	// MISSING: AutoStorageIncrease
	// MISSING: DatabaseFlags
	// MISSING: DataDiskType
	// MISSING: DataDiskSizeGB
	// MISSING: Zone
	// MISSING: SecondaryZone
	// MISSING: SourceID
	// MISSING: RootPassword
	out.RootPasswordSet = direct.LazyPtr(in.GetRootPasswordSet())
	// MISSING: Collation
	// MISSING: CmekKeyName
	// MISSING: AvailabilityType
	// MISSING: Edition
	return out
}
func CloudSQLSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudSQLSettingsObservedState) *pb.CloudSqlSettings {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlSettings{}
	// MISSING: DatabaseVersion
	// MISSING: UserLabels
	// MISSING: Tier
	// MISSING: StorageAutoResizeLimit
	// MISSING: ActivationPolicy
	// MISSING: IPConfig
	// MISSING: AutoStorageIncrease
	// MISSING: DatabaseFlags
	// MISSING: DataDiskType
	// MISSING: DataDiskSizeGB
	// MISSING: Zone
	// MISSING: SecondaryZone
	// MISSING: SourceID
	// MISSING: RootPassword
	out.RootPasswordSet = direct.ValueOf(in.RootPasswordSet)
	// MISSING: Collation
	// MISSING: CmekKeyName
	// MISSING: AvailabilityType
	// MISSING: Edition
	return out
}
func ConversionWorkspaceInfo_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspaceInfo) *krm.ConversionWorkspaceInfo {
	if in == nil {
		return nil
	}
	out := &krm.ConversionWorkspaceInfo{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CommitID = direct.LazyPtr(in.GetCommitId())
	return out
}
func ConversionWorkspaceInfo_ToProto(mapCtx *direct.MapContext, in *krm.ConversionWorkspaceInfo) *pb.ConversionWorkspaceInfo {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspaceInfo{}
	out.Name = direct.ValueOf(in.Name)
	out.CommitId = direct.ValueOf(in.CommitID)
	return out
}
func DatabaseEngineInfo_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseEngineInfo) *krm.DatabaseEngineInfo {
	if in == nil {
		return nil
	}
	out := &krm.DatabaseEngineInfo{}
	out.Engine = direct.Enum_FromProto(mapCtx, in.GetEngine())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func DatabaseEngineInfo_ToProto(mapCtx *direct.MapContext, in *krm.DatabaseEngineInfo) *pb.DatabaseEngineInfo {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseEngineInfo{}
	out.Engine = direct.Enum_ToProto[pb.DatabaseEngine](mapCtx, in.Engine)
	out.Version = direct.ValueOf(in.Version)
	return out
}
func DatabaseType_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseType) *krm.DatabaseType {
	if in == nil {
		return nil
	}
	out := &krm.DatabaseType{}
	out.Provider = direct.Enum_FromProto(mapCtx, in.GetProvider())
	out.Engine = direct.Enum_FromProto(mapCtx, in.GetEngine())
	return out
}
func DatabaseType_ToProto(mapCtx *direct.MapContext, in *krm.DatabaseType) *pb.DatabaseType {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseType{}
	out.Provider = direct.Enum_ToProto[pb.DatabaseProvider](mapCtx, in.Provider)
	out.Engine = direct.Enum_ToProto[pb.DatabaseEngine](mapCtx, in.Engine)
	return out
}
func ForwardSSHTunnelConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.ForwardSshTunnelConnectivity) *krm.ForwardSSHTunnelConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.ForwardSSHTunnelConnectivity{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Port = direct.LazyPtr(in.GetPort())
	// MISSING: Password
	out.PrivateKey = direct.LazyPtr(in.GetPrivateKey())
	return out
}
func ForwardSSHTunnelConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.ForwardSSHTunnelConnectivity) *pb.ForwardSshTunnelConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.ForwardSshTunnelConnectivity{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Username = direct.ValueOf(in.Username)
	out.Port = direct.ValueOf(in.Port)
	// MISSING: Password
	if oneof := ForwardSSHTunnelConnectivity_PrivateKey_ToProto(mapCtx, in.PrivateKey); oneof != nil {
		out.AuthenticationMethod = oneof
	}
	return out
}
func ForwardSSHTunnelConnectivity_PrivateKey_ToProto(mapCtx *direct.MapContext, in *string) *pb.ForwardSshTunnelConnectivity_PrivateKey {
	if in == nil {
		return nil
	}
	return &pb.ForwardSshTunnelConnectivity_PrivateKey{PrivateKey: *in}
}
func MigrationJob_DumpFlag_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob_DumpFlag) *krm.MigrationJob_DumpFlag {
	if in == nil {
		return nil
	}
	out := &krm.MigrationJob_DumpFlag{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func MigrationJob_DumpFlag_ToProto(mapCtx *direct.MapContext, in *krm.MigrationJob_DumpFlag) *pb.MigrationJob_DumpFlag {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob_DumpFlag{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func MigrationJob_DumpFlags_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob_DumpFlags) *krm.MigrationJob_DumpFlags {
	if in == nil {
		return nil
	}
	out := &krm.MigrationJob_DumpFlags{}
	out.DumpFlags = direct.Slice_FromProto(mapCtx, in.DumpFlags, MigrationJob_DumpFlag_FromProto)
	return out
}
func MigrationJob_DumpFlags_ToProto(mapCtx *direct.MapContext, in *krm.MigrationJob_DumpFlags) *pb.MigrationJob_DumpFlags {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob_DumpFlags{}
	out.DumpFlags = direct.Slice_ToProto(mapCtx, in.DumpFlags, MigrationJob_DumpFlag_ToProto)
	return out
}
func MigrationJob_PerformanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob_PerformanceConfig) *krm.MigrationJob_PerformanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.MigrationJob_PerformanceConfig{}
	out.DumpParallelLevel = direct.Enum_FromProto(mapCtx, in.GetDumpParallelLevel())
	return out
}
func MigrationJob_PerformanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.MigrationJob_PerformanceConfig) *pb.MigrationJob_PerformanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob_PerformanceConfig{}
	out.DumpParallelLevel = direct.Enum_ToProto[pb.MigrationJob_PerformanceConfig_DumpParallelLevel](mapCtx, in.DumpParallelLevel)
	return out
}
func MySQLConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.MySqlConnectionProfile) *krm.MySQLConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.MySQLConnectionProfile{}
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	// MISSING: Password
	// MISSING: PasswordSet
	out.SSL = SSLConfig_FromProto(mapCtx, in.GetSsl())
	out.CloudSQLID = direct.LazyPtr(in.GetCloudSqlId())
	return out
}
func MySQLConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.MySQLConnectionProfile) *pb.MySqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.MySqlConnectionProfile{}
	out.Host = direct.ValueOf(in.Host)
	out.Port = direct.ValueOf(in.Port)
	out.Username = direct.ValueOf(in.Username)
	// MISSING: Password
	// MISSING: PasswordSet
	out.Ssl = SSLConfig_ToProto(mapCtx, in.SSL)
	out.CloudSqlId = direct.ValueOf(in.CloudSQLID)
	return out
}
func OracleConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.OracleConnectionProfile) *krm.OracleConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.OracleConnectionProfile{}
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	// MISSING: Password
	// MISSING: PasswordSet
	out.DatabaseService = direct.LazyPtr(in.GetDatabaseService())
	out.SSL = SSLConfig_FromProto(mapCtx, in.GetSsl())
	out.StaticServiceIPConnectivity = StaticServiceIPConnectivity_FromProto(mapCtx, in.GetStaticServiceIpConnectivity())
	out.ForwardSSHConnectivity = ForwardSSHTunnelConnectivity_FromProto(mapCtx, in.GetForwardSshConnectivity())
	out.PrivateConnectivity = PrivateConnectivity_FromProto(mapCtx, in.GetPrivateConnectivity())
	return out
}
func OracleConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.OracleConnectionProfile) *pb.OracleConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.OracleConnectionProfile{}
	out.Host = direct.ValueOf(in.Host)
	out.Port = direct.ValueOf(in.Port)
	out.Username = direct.ValueOf(in.Username)
	// MISSING: Password
	// MISSING: PasswordSet
	out.DatabaseService = direct.ValueOf(in.DatabaseService)
	out.Ssl = SSLConfig_ToProto(mapCtx, in.SSL)
	if oneof := StaticServiceIPConnectivity_ToProto(mapCtx, in.StaticServiceIPConnectivity); oneof != nil {
		out.Connectivity = &pb.OracleConnectionProfile_StaticServiceIpConnectivity{StaticServiceIpConnectivity: oneof}
	}
	if oneof := ForwardSSHTunnelConnectivity_ToProto(mapCtx, in.ForwardSSHConnectivity); oneof != nil {
		out.Connectivity = &pb.OracleConnectionProfile_ForwardSshConnectivity{ForwardSshConnectivity: oneof}
	}
	if oneof := PrivateConnectivity_ToProto(mapCtx, in.PrivateConnectivity); oneof != nil {
		out.Connectivity = &pb.OracleConnectionProfile_PrivateConnectivity{PrivateConnectivity: oneof}
	}
	return out
}
func PostgreSQLConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.PostgreSqlConnectionProfile) *krm.PostgreSQLConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.PostgreSQLConnectionProfile{}
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	// MISSING: Password
	// MISSING: PasswordSet
	out.SSL = SSLConfig_FromProto(mapCtx, in.GetSsl())
	out.CloudSQLID = direct.LazyPtr(in.GetCloudSqlId())
	// MISSING: NetworkArchitecture
	out.StaticIPConnectivity = StaticIPConnectivity_FromProto(mapCtx, in.GetStaticIpConnectivity())
	out.PrivateServiceConnectConnectivity = PrivateServiceConnectConnectivity_FromProto(mapCtx, in.GetPrivateServiceConnectConnectivity())
	return out
}
func PostgreSQLConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.PostgreSQLConnectionProfile) *pb.PostgreSqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.PostgreSqlConnectionProfile{}
	out.Host = direct.ValueOf(in.Host)
	out.Port = direct.ValueOf(in.Port)
	out.Username = direct.ValueOf(in.Username)
	// MISSING: Password
	// MISSING: PasswordSet
	out.Ssl = SSLConfig_ToProto(mapCtx, in.SSL)
	out.CloudSqlId = direct.ValueOf(in.CloudSQLID)
	// MISSING: NetworkArchitecture
	if oneof := StaticIPConnectivity_ToProto(mapCtx, in.StaticIPConnectivity); oneof != nil {
		out.Connectivity = &pb.PostgreSqlConnectionProfile_StaticIpConnectivity{StaticIpConnectivity: oneof}
	}
	if oneof := PrivateServiceConnectConnectivity_ToProto(mapCtx, in.PrivateServiceConnectConnectivity); oneof != nil {
		out.Connectivity = &pb.PostgreSqlConnectionProfile_PrivateServiceConnectConnectivity{PrivateServiceConnectConnectivity: oneof}
	}
	return out
}
func PrivateConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnectivity) *krm.PrivateConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.PrivateConnectivity{}
	out.PrivateConnection = direct.LazyPtr(in.GetPrivateConnection())
	return out
}
func PrivateConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.PrivateConnectivity) *pb.PrivateConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnectivity{}
	out.PrivateConnection = direct.ValueOf(in.PrivateConnection)
	return out
}
func PrivateServiceConnectConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.PrivateServiceConnectConnectivity) *krm.PrivateServiceConnectConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.PrivateServiceConnectConnectivity{}
	out.ServiceAttachment = direct.LazyPtr(in.GetServiceAttachment())
	return out
}
func PrivateServiceConnectConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.PrivateServiceConnectConnectivity) *pb.PrivateServiceConnectConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.PrivateServiceConnectConnectivity{}
	out.ServiceAttachment = direct.ValueOf(in.ServiceAttachment)
	return out
}
func ReverseSSHConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.ReverseSshConnectivity) *krm.ReverseSSHConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.ReverseSSHConnectivity{}
	out.VMIP = direct.LazyPtr(in.GetVmIp())
	out.VMPort = direct.LazyPtr(in.GetVmPort())
	if in.GetVm() != "" {
		out.VMRef = &krmcomputev1beta1.InstanceRef{External: in.GetVm()}
	}
	if in.GetVpc() != "" {
		out.VPCRef = &krmcomputev1beta1.ComputeNetworkRef{External: in.GetVpc()}
	}
	return out
}
func ReverseSSHConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.ReverseSSHConnectivity) *pb.ReverseSshConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.ReverseSshConnectivity{}
	out.VmIp = direct.ValueOf(in.VMIP)
	out.VmPort = direct.ValueOf(in.VMPort)
	if in.VMRef != nil {
		out.Vm = in.VMRef.External
	}
	if in.VPCRef != nil {
		out.Vpc = in.VPCRef.External
	}
	return out
}
func SQLAclEntry_FromProto(mapCtx *direct.MapContext, in *pb.SqlAclEntry) *krm.SQLAclEntry {
	if in == nil {
		return nil
	}
	out := &krm.SQLAclEntry{}
	out.Value = direct.LazyPtr(in.GetValue())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.TTL = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	out.Label = direct.LazyPtr(in.GetLabel())
	return out
}
func SQLAclEntry_ToProto(mapCtx *direct.MapContext, in *krm.SQLAclEntry) *pb.SqlAclEntry {
	if in == nil {
		return nil
	}
	out := &pb.SqlAclEntry{}
	out.Value = direct.ValueOf(in.Value)
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime); oneof != nil {
		out.Expiration = &pb.SqlAclEntry_ExpireTime{ExpireTime: oneof}
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.TTL); oneof != nil {
		out.Expiration = &pb.SqlAclEntry_Ttl{Ttl: oneof}
	}
	out.Label = direct.ValueOf(in.Label)
	return out
}
func SQLIPConfig_FromProto(mapCtx *direct.MapContext, in *pb.SqlIpConfig) *krm.SQLIPConfig {
	if in == nil {
		return nil
	}
	out := &krm.SQLIPConfig{}
	out.EnableIPV4 = direct.BoolValue_FromProto(mapCtx, in.GetEnableIpv4())
	out.PrivateNetwork = direct.LazyPtr(in.GetPrivateNetwork())
	out.AllocatedIPRange = direct.LazyPtr(in.GetAllocatedIpRange())
	out.RequireSSL = direct.BoolValue_FromProto(mapCtx, in.GetRequireSsl())
	out.AuthorizedNetworks = direct.Slice_FromProto(mapCtx, in.AuthorizedNetworks, SQLAclEntry_FromProto)
	return out
}
func SQLIPConfig_ToProto(mapCtx *direct.MapContext, in *krm.SQLIPConfig) *pb.SqlIpConfig {
	if in == nil {
		return nil
	}
	out := &pb.SqlIpConfig{}
	out.EnableIpv4 = direct.BoolValue_ToProto(mapCtx, in.EnableIPV4)
	out.PrivateNetwork = direct.ValueOf(in.PrivateNetwork)
	out.AllocatedIpRange = direct.ValueOf(in.AllocatedIPRange)
	out.RequireSsl = direct.BoolValue_ToProto(mapCtx, in.RequireSSL)
	out.AuthorizedNetworks = direct.Slice_ToProto(mapCtx, in.AuthorizedNetworks, SQLAclEntry_ToProto)
	return out
}
func SSLConfig_FromProto(mapCtx *direct.MapContext, in *pb.SslConfig) *krm.SSLConfig {
	if in == nil {
		return nil
	}
	out := &krm.SSLConfig{}
	// MISSING: Type
	out.ClientKey = direct.LazyPtr(in.GetClientKey())
	out.ClientCertificate = direct.LazyPtr(in.GetClientCertificate())
	out.CACertificate = direct.LazyPtr(in.GetCaCertificate())
	return out
}
func SSLConfig_ToProto(mapCtx *direct.MapContext, in *krm.SSLConfig) *pb.SslConfig {
	if in == nil {
		return nil
	}
	out := &pb.SslConfig{}
	// MISSING: Type
	out.ClientKey = direct.ValueOf(in.ClientKey)
	out.ClientCertificate = direct.ValueOf(in.ClientCertificate)
	out.CaCertificate = direct.ValueOf(in.CACertificate)
	return out
}
func SSLConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SslConfig) *krm.SSLConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SSLConfigObservedState{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: ClientKey
	// MISSING: ClientCertificate
	// MISSING: CACertificate
	return out
}
func SSLConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SSLConfigObservedState) *pb.SslConfig {
	if in == nil {
		return nil
	}
	out := &pb.SslConfig{}
	out.Type = direct.Enum_ToProto[pb.SslConfig_SslType](mapCtx, in.Type)
	// MISSING: ClientKey
	// MISSING: ClientCertificate
	// MISSING: CACertificate
	return out
}
func StaticIPConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.StaticIpConnectivity) *krm.StaticIPConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.StaticIPConnectivity{}
	return out
}
func StaticIPConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.StaticIPConnectivity) *pb.StaticIpConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.StaticIpConnectivity{}
	return out
}
func StaticServiceIPConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.StaticServiceIpConnectivity) *krm.StaticServiceIPConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.StaticServiceIPConnectivity{}
	return out
}
func StaticServiceIPConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.StaticServiceIPConnectivity) *pb.StaticServiceIpConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.StaticServiceIpConnectivity{}
	return out
}
func VPCPeeringConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.VpcPeeringConnectivity) *krm.VPCPeeringConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.VPCPeeringConnectivity{}
	if in.GetVpc() != "" {
		out.VPCRef = &krmcomputev1beta1.ComputeNetworkRef{External: in.GetVpc()}
	}
	return out
}
func VPCPeeringConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.VPCPeeringConnectivity) *pb.VpcPeeringConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.VpcPeeringConnectivity{}
	if in.VPCRef != nil {
		out.Vpc = in.VPCRef.External
	}
	return out
}
