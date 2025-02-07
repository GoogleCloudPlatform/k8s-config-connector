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

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
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
	out.VpcNetwork = direct.ValueOf(in.VpcNetwork)
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
	// MISSING: VpcNetwork
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
	// MISSING: VpcNetwork
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
	out.CpuCount = direct.LazyPtr(in.GetCpuCount())
	return out
}
func AlloyDbSettings_PrimaryInstanceSettings_MachineConfig_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbSettings_PrimaryInstanceSettings_MachineConfig) *pb.AlloyDbSettings_PrimaryInstanceSettings_MachineConfig {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbSettings_PrimaryInstanceSettings_MachineConfig{}
	out.CpuCount = direct.ValueOf(in.CpuCount)
	return out
}
func AlloyDbSettings_UserPassword_FromProto(mapCtx *direct.MapContext, in *pb.AlloyDbSettings_UserPassword) *krm.AlloyDbSettings_UserPassword {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDbSettings_UserPassword{}
	out.User = direct.LazyPtr(in.GetUser())
	out.Password = direct.LazyPtr(in.GetPassword())
	// MISSING: PasswordSet
	return out
}
func AlloyDbSettings_UserPassword_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDbSettings_UserPassword) *pb.AlloyDbSettings_UserPassword {
	if in == nil {
		return nil
	}
	out := &pb.AlloyDbSettings_UserPassword{}
	out.User = direct.ValueOf(in.User)
	out.Password = direct.ValueOf(in.Password)
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
func CloudSqlConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlConnectionProfile) *krm.CloudSqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.CloudSqlConnectionProfile{}
	// MISSING: CloudSqlID
	out.Settings = CloudSqlSettings_FromProto(mapCtx, in.GetSettings())
	// MISSING: PrivateIP
	// MISSING: PublicIP
	// MISSING: AdditionalPublicIP
	return out
}
func CloudSqlConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.CloudSqlConnectionProfile) *pb.CloudSqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlConnectionProfile{}
	// MISSING: CloudSqlID
	out.Settings = CloudSqlSettings_ToProto(mapCtx, in.Settings)
	// MISSING: PrivateIP
	// MISSING: PublicIP
	// MISSING: AdditionalPublicIP
	return out
}
func CloudSqlConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlConnectionProfile) *krm.CloudSqlConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudSqlConnectionProfileObservedState{}
	out.CloudSqlID = direct.LazyPtr(in.GetCloudSqlId())
	out.Settings = CloudSqlSettingsObservedState_FromProto(mapCtx, in.GetSettings())
	out.PrivateIP = direct.LazyPtr(in.GetPrivateIp())
	out.PublicIP = direct.LazyPtr(in.GetPublicIp())
	out.AdditionalPublicIP = direct.LazyPtr(in.GetAdditionalPublicIp())
	return out
}
func CloudSqlConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudSqlConnectionProfileObservedState) *pb.CloudSqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlConnectionProfile{}
	out.CloudSqlId = direct.ValueOf(in.CloudSqlID)
	out.Settings = CloudSqlSettingsObservedState_ToProto(mapCtx, in.Settings)
	out.PrivateIp = direct.ValueOf(in.PrivateIP)
	out.PublicIp = direct.ValueOf(in.PublicIP)
	out.AdditionalPublicIp = direct.ValueOf(in.AdditionalPublicIP)
	return out
}
func CloudSqlSettings_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlSettings) *krm.CloudSqlSettings {
	if in == nil {
		return nil
	}
	out := &krm.CloudSqlSettings{}
	out.DatabaseVersion = direct.Enum_FromProto(mapCtx, in.GetDatabaseVersion())
	out.UserLabels = in.UserLabels
	out.Tier = direct.LazyPtr(in.GetTier())
	out.StorageAutoResizeLimit = direct.Int64Value_FromProto(mapCtx, in.GetStorageAutoResizeLimit())
	out.ActivationPolicy = direct.Enum_FromProto(mapCtx, in.GetActivationPolicy())
	out.IPConfig = SqlIpConfig_FromProto(mapCtx, in.GetIpConfig())
	out.AutoStorageIncrease = direct.BoolValue_FromProto(mapCtx, in.GetAutoStorageIncrease())
	out.DatabaseFlags = in.DatabaseFlags
	out.DataDiskType = direct.Enum_FromProto(mapCtx, in.GetDataDiskType())
	out.DataDiskSizeGB = direct.Int64Value_FromProto(mapCtx, in.GetDataDiskSizeGb())
	out.Zone = direct.LazyPtr(in.GetZone())
	out.SecondaryZone = direct.LazyPtr(in.GetSecondaryZone())
	out.SourceID = direct.LazyPtr(in.GetSourceId())
	out.RootPassword = direct.LazyPtr(in.GetRootPassword())
	// MISSING: RootPasswordSet
	out.Collation = direct.LazyPtr(in.GetCollation())
	out.CmekKeyName = direct.LazyPtr(in.GetCmekKeyName())
	out.AvailabilityType = direct.Enum_FromProto(mapCtx, in.GetAvailabilityType())
	out.Edition = direct.Enum_FromProto(mapCtx, in.GetEdition())
	return out
}
func CloudSqlSettings_ToProto(mapCtx *direct.MapContext, in *krm.CloudSqlSettings) *pb.CloudSqlSettings {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlSettings{}
	out.DatabaseVersion = direct.Enum_ToProto[pb.CloudSqlSettings_SqlDatabaseVersion](mapCtx, in.DatabaseVersion)
	out.UserLabels = in.UserLabels
	out.Tier = direct.ValueOf(in.Tier)
	out.StorageAutoResizeLimit = direct.Int64Value_ToProto(mapCtx, in.StorageAutoResizeLimit)
	out.ActivationPolicy = direct.Enum_ToProto[pb.CloudSqlSettings_SqlActivationPolicy](mapCtx, in.ActivationPolicy)
	out.IpConfig = SqlIpConfig_ToProto(mapCtx, in.IPConfig)
	out.AutoStorageIncrease = direct.BoolValue_ToProto(mapCtx, in.AutoStorageIncrease)
	out.DatabaseFlags = in.DatabaseFlags
	out.DataDiskType = direct.Enum_ToProto[pb.CloudSqlSettings_SqlDataDiskType](mapCtx, in.DataDiskType)
	out.DataDiskSizeGb = direct.Int64Value_ToProto(mapCtx, in.DataDiskSizeGB)
	out.Zone = direct.ValueOf(in.Zone)
	out.SecondaryZone = direct.ValueOf(in.SecondaryZone)
	out.SourceId = direct.ValueOf(in.SourceID)
	out.RootPassword = direct.ValueOf(in.RootPassword)
	// MISSING: RootPasswordSet
	out.Collation = direct.ValueOf(in.Collation)
	out.CmekKeyName = direct.ValueOf(in.CmekKeyName)
	out.AvailabilityType = direct.Enum_ToProto[pb.CloudSqlSettings_SqlAvailabilityType](mapCtx, in.AvailabilityType)
	out.Edition = direct.Enum_ToProto[pb.CloudSqlSettings_Edition](mapCtx, in.Edition)
	return out
}
func CloudSqlSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlSettings) *krm.CloudSqlSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudSqlSettingsObservedState{}
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
func CloudSqlSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudSqlSettingsObservedState) *pb.CloudSqlSettings {
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
func ClouddmsConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.ClouddmsConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsConnectionProfileObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Mysql
	// MISSING: Postgresql
	// MISSING: Oracle
	// MISSING: Cloudsql
	// MISSING: Alloydb
	// MISSING: Error
	// MISSING: Provider
	return out
}
func ClouddmsConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsConnectionProfileObservedState) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Mysql
	// MISSING: Postgresql
	// MISSING: Oracle
	// MISSING: Cloudsql
	// MISSING: Alloydb
	// MISSING: Error
	// MISSING: Provider
	return out
}
func ClouddmsConnectionProfileSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.ClouddmsConnectionProfileSpec {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsConnectionProfileSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Mysql
	// MISSING: Postgresql
	// MISSING: Oracle
	// MISSING: Cloudsql
	// MISSING: Alloydb
	// MISSING: Error
	// MISSING: Provider
	return out
}
func ClouddmsConnectionProfileSpec_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsConnectionProfileSpec) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Mysql
	// MISSING: Postgresql
	// MISSING: Oracle
	// MISSING: Cloudsql
	// MISSING: Alloydb
	// MISSING: Error
	// MISSING: Provider
	return out
}
func ClouddmsMigrationJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob) *krm.ClouddmsMigrationJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsMigrationJobObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Phase
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	// MISSING: Duration
	// MISSING: Error
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	// MISSING: EndTime
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func ClouddmsMigrationJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsMigrationJobObservedState) *pb.MigrationJob {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Phase
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	// MISSING: Duration
	// MISSING: Error
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	// MISSING: EndTime
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func ClouddmsMigrationJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob) *krm.ClouddmsMigrationJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsMigrationJobSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Phase
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	// MISSING: Duration
	// MISSING: Error
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	// MISSING: EndTime
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func ClouddmsMigrationJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsMigrationJobSpec) *pb.MigrationJob {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Phase
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	// MISSING: Duration
	// MISSING: Error
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	// MISSING: EndTime
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func ConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionProfile{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Mysql = MySqlConnectionProfile_FromProto(mapCtx, in.GetMysql())
	out.Postgresql = PostgreSqlConnectionProfile_FromProto(mapCtx, in.GetPostgresql())
	out.Oracle = OracleConnectionProfile_FromProto(mapCtx, in.GetOracle())
	out.Cloudsql = CloudSqlConnectionProfile_FromProto(mapCtx, in.GetCloudsql())
	out.Alloydb = AlloyDbConnectionProfile_FromProto(mapCtx, in.GetAlloydb())
	// MISSING: Error
	out.Provider = direct.Enum_FromProto(mapCtx, in.GetProvider())
	return out
}
func ConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.ConnectionProfile) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.State = direct.Enum_ToProto[pb.ConnectionProfile_State](mapCtx, in.State)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if oneof := MySqlConnectionProfile_ToProto(mapCtx, in.Mysql); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Mysql{Mysql: oneof}
	}
	if oneof := PostgreSqlConnectionProfile_ToProto(mapCtx, in.Postgresql); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Postgresql{Postgresql: oneof}
	}
	if oneof := OracleConnectionProfile_ToProto(mapCtx, in.Oracle); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Oracle{Oracle: oneof}
	}
	if oneof := CloudSqlConnectionProfile_ToProto(mapCtx, in.Cloudsql); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Cloudsql{Cloudsql: oneof}
	}
	if oneof := AlloyDbConnectionProfile_ToProto(mapCtx, in.Alloydb); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Alloydb{Alloydb: oneof}
	}
	// MISSING: Error
	out.Provider = direct.Enum_ToProto[pb.DatabaseProvider](mapCtx, in.Provider)
	return out
}
func ConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.ConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionProfileObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	out.Mysql = MySqlConnectionProfileObservedState_FromProto(mapCtx, in.GetMysql())
	out.Postgresql = PostgreSqlConnectionProfileObservedState_FromProto(mapCtx, in.GetPostgresql())
	out.Oracle = OracleConnectionProfileObservedState_FromProto(mapCtx, in.GetOracle())
	out.Cloudsql = CloudSqlConnectionProfileObservedState_FromProto(mapCtx, in.GetCloudsql())
	out.Alloydb = AlloyDbConnectionProfileObservedState_FromProto(mapCtx, in.GetAlloydb())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	// MISSING: Provider
	return out
}
func ConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectionProfileObservedState) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	if oneof := MySqlConnectionProfileObservedState_ToProto(mapCtx, in.Mysql); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Mysql{Mysql: oneof}
	}
	if oneof := PostgreSqlConnectionProfileObservedState_ToProto(mapCtx, in.Postgresql); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Postgresql{Postgresql: oneof}
	}
	if oneof := OracleConnectionProfileObservedState_ToProto(mapCtx, in.Oracle); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Oracle{Oracle: oneof}
	}
	if oneof := CloudSqlConnectionProfileObservedState_ToProto(mapCtx, in.Cloudsql); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Cloudsql{Cloudsql: oneof}
	}
	if oneof := AlloyDbConnectionProfileObservedState_ToProto(mapCtx, in.Alloydb); oneof != nil {
		out.ConnectionProfile = &pb.ConnectionProfile_Alloydb{Alloydb: oneof}
	}
	out.Error = Status_ToProto(mapCtx, in.Error)
	// MISSING: Provider
	return out
}
func ForwardSshTunnelConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.ForwardSshTunnelConnectivity) *krm.ForwardSshTunnelConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.ForwardSshTunnelConnectivity{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.PrivateKey = direct.LazyPtr(in.GetPrivateKey())
	return out
}
func ForwardSshTunnelConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.ForwardSshTunnelConnectivity) *pb.ForwardSshTunnelConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.ForwardSshTunnelConnectivity{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.Username = direct.ValueOf(in.Username)
	out.Port = direct.ValueOf(in.Port)
	if oneof := ForwardSshTunnelConnectivity_Password_ToProto(mapCtx, in.Password); oneof != nil {
		out.AuthenticationMethod = oneof
	}
	if oneof := ForwardSshTunnelConnectivity_PrivateKey_ToProto(mapCtx, in.PrivateKey); oneof != nil {
		out.AuthenticationMethod = oneof
	}
	return out
}
func MySqlConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.MySqlConnectionProfile) *krm.MySqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.MySqlConnectionProfile{}
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	// MISSING: PasswordSet
	out.Ssl = SslConfig_FromProto(mapCtx, in.GetSsl())
	out.CloudSqlID = direct.LazyPtr(in.GetCloudSqlId())
	return out
}
func MySqlConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.MySqlConnectionProfile) *pb.MySqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.MySqlConnectionProfile{}
	out.Host = direct.ValueOf(in.Host)
	out.Port = direct.ValueOf(in.Port)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	// MISSING: PasswordSet
	out.Ssl = SslConfig_ToProto(mapCtx, in.Ssl)
	out.CloudSqlId = direct.ValueOf(in.CloudSqlID)
	return out
}
func MySqlConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MySqlConnectionProfile) *krm.MySqlConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MySqlConnectionProfileObservedState{}
	// MISSING: Host
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	out.PasswordSet = direct.LazyPtr(in.GetPasswordSet())
	out.Ssl = SslConfigObservedState_FromProto(mapCtx, in.GetSsl())
	// MISSING: CloudSqlID
	return out
}
func MySqlConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MySqlConnectionProfileObservedState) *pb.MySqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.MySqlConnectionProfile{}
	// MISSING: Host
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	out.PasswordSet = direct.ValueOf(in.PasswordSet)
	out.Ssl = SslConfigObservedState_ToProto(mapCtx, in.Ssl)
	// MISSING: CloudSqlID
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
	out.Password = direct.LazyPtr(in.GetPassword())
	// MISSING: PasswordSet
	out.DatabaseService = direct.LazyPtr(in.GetDatabaseService())
	out.Ssl = SslConfig_FromProto(mapCtx, in.GetSsl())
	out.StaticServiceIPConnectivity = StaticServiceIpConnectivity_FromProto(mapCtx, in.GetStaticServiceIpConnectivity())
	out.ForwardSSHConnectivity = ForwardSshTunnelConnectivity_FromProto(mapCtx, in.GetForwardSshConnectivity())
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
	out.Password = direct.ValueOf(in.Password)
	// MISSING: PasswordSet
	out.DatabaseService = direct.ValueOf(in.DatabaseService)
	out.Ssl = SslConfig_ToProto(mapCtx, in.Ssl)
	if oneof := StaticServiceIpConnectivity_ToProto(mapCtx, in.StaticServiceIPConnectivity); oneof != nil {
		out.Connectivity = &pb.OracleConnectionProfile_StaticServiceIpConnectivity{StaticServiceIpConnectivity: oneof}
	}
	if oneof := ForwardSshTunnelConnectivity_ToProto(mapCtx, in.ForwardSSHConnectivity); oneof != nil {
		out.Connectivity = &pb.OracleConnectionProfile_ForwardSshConnectivity{ForwardSshConnectivity: oneof}
	}
	if oneof := PrivateConnectivity_ToProto(mapCtx, in.PrivateConnectivity); oneof != nil {
		out.Connectivity = &pb.OracleConnectionProfile_PrivateConnectivity{PrivateConnectivity: oneof}
	}
	return out
}
func OracleConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OracleConnectionProfile) *krm.OracleConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracleConnectionProfileObservedState{}
	// MISSING: Host
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	out.PasswordSet = direct.LazyPtr(in.GetPasswordSet())
	// MISSING: DatabaseService
	// MISSING: Ssl
	// MISSING: StaticServiceIPConnectivity
	// MISSING: ForwardSSHConnectivity
	// MISSING: PrivateConnectivity
	return out
}
func OracleConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracleConnectionProfileObservedState) *pb.OracleConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.OracleConnectionProfile{}
	// MISSING: Host
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	out.PasswordSet = direct.ValueOf(in.PasswordSet)
	// MISSING: DatabaseService
	// MISSING: Ssl
	// MISSING: StaticServiceIPConnectivity
	// MISSING: ForwardSSHConnectivity
	// MISSING: PrivateConnectivity
	return out
}
func PostgreSqlConnectionProfile_FromProto(mapCtx *direct.MapContext, in *pb.PostgreSqlConnectionProfile) *krm.PostgreSqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &krm.PostgreSqlConnectionProfile{}
	out.Host = direct.LazyPtr(in.GetHost())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	// MISSING: PasswordSet
	out.Ssl = SslConfig_FromProto(mapCtx, in.GetSsl())
	out.CloudSqlID = direct.LazyPtr(in.GetCloudSqlId())
	// MISSING: NetworkArchitecture
	out.StaticIPConnectivity = StaticIpConnectivity_FromProto(mapCtx, in.GetStaticIpConnectivity())
	out.PrivateServiceConnectConnectivity = PrivateServiceConnectConnectivity_FromProto(mapCtx, in.GetPrivateServiceConnectConnectivity())
	return out
}
func PostgreSqlConnectionProfile_ToProto(mapCtx *direct.MapContext, in *krm.PostgreSqlConnectionProfile) *pb.PostgreSqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.PostgreSqlConnectionProfile{}
	out.Host = direct.ValueOf(in.Host)
	out.Port = direct.ValueOf(in.Port)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	// MISSING: PasswordSet
	out.Ssl = SslConfig_ToProto(mapCtx, in.Ssl)
	out.CloudSqlId = direct.ValueOf(in.CloudSqlID)
	// MISSING: NetworkArchitecture
	if oneof := StaticIpConnectivity_ToProto(mapCtx, in.StaticIPConnectivity); oneof != nil {
		out.Connectivity = &pb.PostgreSqlConnectionProfile_StaticIpConnectivity{StaticIpConnectivity: oneof}
	}
	if oneof := PrivateServiceConnectConnectivity_ToProto(mapCtx, in.PrivateServiceConnectConnectivity); oneof != nil {
		out.Connectivity = &pb.PostgreSqlConnectionProfile_PrivateServiceConnectConnectivity{PrivateServiceConnectConnectivity: oneof}
	}
	return out
}
func PostgreSqlConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PostgreSqlConnectionProfile) *krm.PostgreSqlConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PostgreSqlConnectionProfileObservedState{}
	// MISSING: Host
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	out.PasswordSet = direct.LazyPtr(in.GetPasswordSet())
	// MISSING: Ssl
	// MISSING: CloudSqlID
	out.NetworkArchitecture = direct.Enum_FromProto(mapCtx, in.GetNetworkArchitecture())
	// MISSING: StaticIPConnectivity
	// MISSING: PrivateServiceConnectConnectivity
	return out
}
func PostgreSqlConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PostgreSqlConnectionProfileObservedState) *pb.PostgreSqlConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.PostgreSqlConnectionProfile{}
	// MISSING: Host
	// MISSING: Port
	// MISSING: Username
	// MISSING: Password
	out.PasswordSet = direct.ValueOf(in.PasswordSet)
	// MISSING: Ssl
	// MISSING: CloudSqlID
	out.NetworkArchitecture = direct.Enum_ToProto[pb.NetworkArchitecture](mapCtx, in.NetworkArchitecture)
	// MISSING: StaticIPConnectivity
	// MISSING: PrivateServiceConnectConnectivity
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
func SqlAclEntry_FromProto(mapCtx *direct.MapContext, in *pb.SqlAclEntry) *krm.SqlAclEntry {
	if in == nil {
		return nil
	}
	out := &krm.SqlAclEntry{}
	out.Value = direct.LazyPtr(in.GetValue())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Ttl = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	out.Label = direct.LazyPtr(in.GetLabel())
	return out
}
func SqlAclEntry_ToProto(mapCtx *direct.MapContext, in *krm.SqlAclEntry) *pb.SqlAclEntry {
	if in == nil {
		return nil
	}
	out := &pb.SqlAclEntry{}
	out.Value = direct.ValueOf(in.Value)
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime); oneof != nil {
		out.Expiration = &pb.SqlAclEntry_ExpireTime{ExpireTime: oneof}
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.Ttl); oneof != nil {
		out.Expiration = &pb.SqlAclEntry_Ttl{Ttl: oneof}
	}
	out.Label = direct.ValueOf(in.Label)
	return out
}
func SqlIpConfig_FromProto(mapCtx *direct.MapContext, in *pb.SqlIpConfig) *krm.SqlIpConfig {
	if in == nil {
		return nil
	}
	out := &krm.SqlIpConfig{}
	out.EnableIpv4 = direct.BoolValue_FromProto(mapCtx, in.GetEnableIpv4())
	out.PrivateNetwork = direct.LazyPtr(in.GetPrivateNetwork())
	out.AllocatedIPRange = direct.LazyPtr(in.GetAllocatedIpRange())
	out.RequireSsl = direct.BoolValue_FromProto(mapCtx, in.GetRequireSsl())
	out.AuthorizedNetworks = direct.Slice_FromProto(mapCtx, in.AuthorizedNetworks, SqlAclEntry_FromProto)
	return out
}
func SqlIpConfig_ToProto(mapCtx *direct.MapContext, in *krm.SqlIpConfig) *pb.SqlIpConfig {
	if in == nil {
		return nil
	}
	out := &pb.SqlIpConfig{}
	out.EnableIpv4 = direct.BoolValue_ToProto(mapCtx, in.EnableIpv4)
	out.PrivateNetwork = direct.ValueOf(in.PrivateNetwork)
	out.AllocatedIpRange = direct.ValueOf(in.AllocatedIPRange)
	out.RequireSsl = direct.BoolValue_ToProto(mapCtx, in.RequireSsl)
	out.AuthorizedNetworks = direct.Slice_ToProto(mapCtx, in.AuthorizedNetworks, SqlAclEntry_ToProto)
	return out
}
func SslConfig_FromProto(mapCtx *direct.MapContext, in *pb.SslConfig) *krm.SslConfig {
	if in == nil {
		return nil
	}
	out := &krm.SslConfig{}
	// MISSING: Type
	out.ClientKey = direct.LazyPtr(in.GetClientKey())
	out.ClientCertificate = direct.LazyPtr(in.GetClientCertificate())
	out.CaCertificate = direct.LazyPtr(in.GetCaCertificate())
	return out
}
func SslConfig_ToProto(mapCtx *direct.MapContext, in *krm.SslConfig) *pb.SslConfig {
	if in == nil {
		return nil
	}
	out := &pb.SslConfig{}
	// MISSING: Type
	out.ClientKey = direct.ValueOf(in.ClientKey)
	out.ClientCertificate = direct.ValueOf(in.ClientCertificate)
	out.CaCertificate = direct.ValueOf(in.CaCertificate)
	return out
}
func SslConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SslConfig) *krm.SslConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SslConfigObservedState{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: ClientKey
	// MISSING: ClientCertificate
	// MISSING: CaCertificate
	return out
}
func SslConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SslConfigObservedState) *pb.SslConfig {
	if in == nil {
		return nil
	}
	out := &pb.SslConfig{}
	out.Type = direct.Enum_ToProto[pb.SslConfig_SslType](mapCtx, in.Type)
	// MISSING: ClientKey
	// MISSING: ClientCertificate
	// MISSING: CaCertificate
	return out
}
func StaticIpConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.StaticIpConnectivity) *krm.StaticIpConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.StaticIpConnectivity{}
	return out
}
func StaticIpConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.StaticIpConnectivity) *pb.StaticIpConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.StaticIpConnectivity{}
	return out
}
func StaticServiceIpConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.StaticServiceIpConnectivity) *krm.StaticServiceIpConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.StaticServiceIpConnectivity{}
	return out
}
func StaticServiceIpConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.StaticServiceIpConnectivity) *pb.StaticServiceIpConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.StaticServiceIpConnectivity{}
	return out
}
