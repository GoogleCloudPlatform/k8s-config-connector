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

package oracledatabase

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AllConnectionStrings_FromProto(mapCtx *direct.MapContext, in *pb.AllConnectionStrings) *krm.AllConnectionStrings {
	if in == nil {
		return nil
	}
	out := &krm.AllConnectionStrings{}
	// MISSING: High
	// MISSING: Low
	// MISSING: Medium
	return out
}
func AllConnectionStrings_ToProto(mapCtx *direct.MapContext, in *krm.AllConnectionStrings) *pb.AllConnectionStrings {
	if in == nil {
		return nil
	}
	out := &pb.AllConnectionStrings{}
	// MISSING: High
	// MISSING: Low
	// MISSING: Medium
	return out
}
func AllConnectionStringsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AllConnectionStrings) *krm.AllConnectionStringsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AllConnectionStringsObservedState{}
	out.High = direct.LazyPtr(in.GetHigh())
	out.Low = direct.LazyPtr(in.GetLow())
	out.Medium = direct.LazyPtr(in.GetMedium())
	return out
}
func AllConnectionStringsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AllConnectionStringsObservedState) *pb.AllConnectionStrings {
	if in == nil {
		return nil
	}
	out := &pb.AllConnectionStrings{}
	out.High = direct.ValueOf(in.High)
	out.Low = direct.ValueOf(in.Low)
	out.Medium = direct.ValueOf(in.Medium)
	return out
}
func AutonomousDatabase_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabase) *krm.AutonomousDatabase {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabase{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: EntitlementID
	out.AdminPassword = direct.LazyPtr(in.GetAdminPassword())
	out.Properties = AutonomousDatabaseProperties_FromProto(mapCtx, in.GetProperties())
	out.Labels = in.Labels
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Cidr = direct.LazyPtr(in.GetCidr())
	// MISSING: CreateTime
	return out
}
func AutonomousDatabase_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabase) *pb.AutonomousDatabase {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabase{}
	out.Name = direct.ValueOf(in.Name)
	out.Database = direct.ValueOf(in.Database)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: EntitlementID
	out.AdminPassword = direct.ValueOf(in.AdminPassword)
	out.Properties = AutonomousDatabaseProperties_ToProto(mapCtx, in.Properties)
	out.Labels = in.Labels
	out.Network = direct.ValueOf(in.Network)
	out.Cidr = direct.ValueOf(in.Cidr)
	// MISSING: CreateTime
	return out
}
func AutonomousDatabaseApex_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseApex) *krm.AutonomousDatabaseApex {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseApex{}
	// MISSING: ApexVersion
	// MISSING: OrdsVersion
	return out
}
func AutonomousDatabaseApex_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseApex) *pb.AutonomousDatabaseApex {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseApex{}
	// MISSING: ApexVersion
	// MISSING: OrdsVersion
	return out
}
func AutonomousDatabaseApexObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseApex) *krm.AutonomousDatabaseApexObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseApexObservedState{}
	out.ApexVersion = direct.LazyPtr(in.GetApexVersion())
	out.OrdsVersion = direct.LazyPtr(in.GetOrdsVersion())
	return out
}
func AutonomousDatabaseApexObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseApexObservedState) *pb.AutonomousDatabaseApex {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseApex{}
	out.ApexVersion = direct.ValueOf(in.ApexVersion)
	out.OrdsVersion = direct.ValueOf(in.OrdsVersion)
	return out
}
func AutonomousDatabaseConnectionStrings_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseConnectionStrings) *krm.AutonomousDatabaseConnectionStrings {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseConnectionStrings{}
	// MISSING: AllConnectionStrings
	// MISSING: Dedicated
	// MISSING: High
	// MISSING: Low
	// MISSING: Medium
	// MISSING: Profiles
	return out
}
func AutonomousDatabaseConnectionStrings_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseConnectionStrings) *pb.AutonomousDatabaseConnectionStrings {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseConnectionStrings{}
	// MISSING: AllConnectionStrings
	// MISSING: Dedicated
	// MISSING: High
	// MISSING: Low
	// MISSING: Medium
	// MISSING: Profiles
	return out
}
func AutonomousDatabaseConnectionStringsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseConnectionStrings) *krm.AutonomousDatabaseConnectionStringsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseConnectionStringsObservedState{}
	out.AllConnectionStrings = AllConnectionStrings_FromProto(mapCtx, in.GetAllConnectionStrings())
	out.Dedicated = direct.LazyPtr(in.GetDedicated())
	out.High = direct.LazyPtr(in.GetHigh())
	out.Low = direct.LazyPtr(in.GetLow())
	out.Medium = direct.LazyPtr(in.GetMedium())
	out.Profiles = direct.Slice_FromProto(mapCtx, in.Profiles, DatabaseConnectionStringProfile_FromProto)
	return out
}
func AutonomousDatabaseConnectionStringsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseConnectionStringsObservedState) *pb.AutonomousDatabaseConnectionStrings {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseConnectionStrings{}
	out.AllConnectionStrings = AllConnectionStrings_ToProto(mapCtx, in.AllConnectionStrings)
	out.Dedicated = direct.ValueOf(in.Dedicated)
	out.High = direct.ValueOf(in.High)
	out.Low = direct.ValueOf(in.Low)
	out.Medium = direct.ValueOf(in.Medium)
	out.Profiles = direct.Slice_ToProto(mapCtx, in.Profiles, DatabaseConnectionStringProfile_ToProto)
	return out
}
func AutonomousDatabaseConnectionUrls_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseConnectionUrls) *krm.AutonomousDatabaseConnectionUrls {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseConnectionUrls{}
	// MISSING: ApexURI
	// MISSING: DatabaseTransformsURI
	// MISSING: GraphStudioURI
	// MISSING: MachineLearningNotebookURI
	// MISSING: MachineLearningUserManagementURI
	// MISSING: MongoDbURI
	// MISSING: OrdsURI
	// MISSING: SqlDevWebURI
	return out
}
func AutonomousDatabaseConnectionUrls_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseConnectionUrls) *pb.AutonomousDatabaseConnectionUrls {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseConnectionUrls{}
	// MISSING: ApexURI
	// MISSING: DatabaseTransformsURI
	// MISSING: GraphStudioURI
	// MISSING: MachineLearningNotebookURI
	// MISSING: MachineLearningUserManagementURI
	// MISSING: MongoDbURI
	// MISSING: OrdsURI
	// MISSING: SqlDevWebURI
	return out
}
func AutonomousDatabaseConnectionUrlsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseConnectionUrls) *krm.AutonomousDatabaseConnectionUrlsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseConnectionUrlsObservedState{}
	out.ApexURI = direct.LazyPtr(in.GetApexUri())
	out.DatabaseTransformsURI = direct.LazyPtr(in.GetDatabaseTransformsUri())
	out.GraphStudioURI = direct.LazyPtr(in.GetGraphStudioUri())
	out.MachineLearningNotebookURI = direct.LazyPtr(in.GetMachineLearningNotebookUri())
	out.MachineLearningUserManagementURI = direct.LazyPtr(in.GetMachineLearningUserManagementUri())
	out.MongoDbURI = direct.LazyPtr(in.GetMongoDbUri())
	out.OrdsURI = direct.LazyPtr(in.GetOrdsUri())
	out.SqlDevWebURI = direct.LazyPtr(in.GetSqlDevWebUri())
	return out
}
func AutonomousDatabaseConnectionUrlsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseConnectionUrlsObservedState) *pb.AutonomousDatabaseConnectionUrls {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseConnectionUrls{}
	out.ApexUri = direct.ValueOf(in.ApexURI)
	out.DatabaseTransformsUri = direct.ValueOf(in.DatabaseTransformsURI)
	out.GraphStudioUri = direct.ValueOf(in.GraphStudioURI)
	out.MachineLearningNotebookUri = direct.ValueOf(in.MachineLearningNotebookURI)
	out.MachineLearningUserManagementUri = direct.ValueOf(in.MachineLearningUserManagementURI)
	out.MongoDbUri = direct.ValueOf(in.MongoDbURI)
	out.OrdsUri = direct.ValueOf(in.OrdsURI)
	out.SqlDevWebUri = direct.ValueOf(in.SqlDevWebURI)
	return out
}
func AutonomousDatabaseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabase) *krm.AutonomousDatabaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseObservedState{}
	// MISSING: Name
	// MISSING: Database
	// MISSING: DisplayName
	out.EntitlementID = direct.LazyPtr(in.GetEntitlementId())
	// MISSING: AdminPassword
	out.Properties = AutonomousDatabasePropertiesObservedState_FromProto(mapCtx, in.GetProperties())
	// MISSING: Labels
	// MISSING: Network
	// MISSING: Cidr
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func AutonomousDatabaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseObservedState) *pb.AutonomousDatabase {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabase{}
	// MISSING: Name
	// MISSING: Database
	// MISSING: DisplayName
	out.EntitlementId = direct.ValueOf(in.EntitlementID)
	// MISSING: AdminPassword
	out.Properties = AutonomousDatabasePropertiesObservedState_ToProto(mapCtx, in.Properties)
	// MISSING: Labels
	// MISSING: Network
	// MISSING: Cidr
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func AutonomousDatabaseProperties_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseProperties) *krm.AutonomousDatabaseProperties {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseProperties{}
	// MISSING: Ocid
	out.ComputeCount = direct.LazyPtr(in.GetComputeCount())
	out.CpuCoreCount = direct.LazyPtr(in.GetCpuCoreCount())
	out.DataStorageSizeTb = direct.LazyPtr(in.GetDataStorageSizeTb())
	out.DataStorageSizeGB = direct.LazyPtr(in.GetDataStorageSizeGb())
	out.DbWorkload = direct.Enum_FromProto(mapCtx, in.GetDbWorkload())
	out.DbEdition = direct.Enum_FromProto(mapCtx, in.GetDbEdition())
	out.CharacterSet = direct.LazyPtr(in.GetCharacterSet())
	out.NCharacterSet = direct.LazyPtr(in.GetNCharacterSet())
	out.PrivateEndpointIP = direct.LazyPtr(in.GetPrivateEndpointIp())
	out.PrivateEndpointLabel = direct.LazyPtr(in.GetPrivateEndpointLabel())
	out.DbVersion = direct.LazyPtr(in.GetDbVersion())
	out.IsAutoScalingEnabled = direct.LazyPtr(in.GetIsAutoScalingEnabled())
	out.IsStorageAutoScalingEnabled = direct.LazyPtr(in.GetIsStorageAutoScalingEnabled())
	out.LicenseType = direct.Enum_FromProto(mapCtx, in.GetLicenseType())
	out.CustomerContacts = direct.Slice_FromProto(mapCtx, in.CustomerContacts, CustomerContact_FromProto)
	out.SecretID = direct.LazyPtr(in.GetSecretId())
	out.VaultID = direct.LazyPtr(in.GetVaultId())
	out.MaintenanceScheduleType = direct.Enum_FromProto(mapCtx, in.GetMaintenanceScheduleType())
	out.MtlsConnectionRequired = direct.LazyPtr(in.GetMtlsConnectionRequired())
	out.BackupRetentionPeriodDays = direct.LazyPtr(in.GetBackupRetentionPeriodDays())
	// MISSING: ActualUsedDataStorageSizeTb
	// MISSING: AllocatedStorageSizeTb
	// MISSING: ApexDetails
	// MISSING: ArePrimaryAllowlistedIpsUsed
	// MISSING: LifecycleDetails
	// MISSING: State
	// MISSING: AutonomousContainerDatabaseID
	// MISSING: AvailableUpgradeVersions
	// MISSING: ConnectionStrings
	// MISSING: ConnectionUrls
	// MISSING: FailedDataRecoveryDuration
	// MISSING: MemoryTableGbs
	// MISSING: IsLocalDataGuardEnabled
	// MISSING: LocalAdgAutoFailoverMaxDataLossLimit
	// MISSING: LocalStandbyDb
	// MISSING: MemoryPerOracleComputeUnitGbs
	// MISSING: LocalDisasterRecoveryType
	// MISSING: DataSafeState
	// MISSING: DatabaseManagementState
	// MISSING: OpenMode
	// MISSING: OperationsInsightsState
	// MISSING: PeerDbIds
	// MISSING: PermissionLevel
	// MISSING: PrivateEndpoint
	// MISSING: RefreshableMode
	// MISSING: RefreshableState
	// MISSING: Role
	// MISSING: ScheduledOperationDetails
	// MISSING: SqlWebDeveloperURL
	// MISSING: SupportedCloneRegions
	// MISSING: UsedDataStorageSizeTbs
	// MISSING: OciURL
	// MISSING: TotalAutoBackupStorageSizeGbs
	// MISSING: NextLongTermBackupTime
	// MISSING: MaintenanceBeginTime
	// MISSING: MaintenanceEndTime
	return out
}
func AutonomousDatabaseProperties_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseProperties) *pb.AutonomousDatabaseProperties {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseProperties{}
	// MISSING: Ocid
	out.ComputeCount = direct.ValueOf(in.ComputeCount)
	out.CpuCoreCount = direct.ValueOf(in.CpuCoreCount)
	out.DataStorageSizeTb = direct.ValueOf(in.DataStorageSizeTb)
	out.DataStorageSizeGb = direct.ValueOf(in.DataStorageSizeGB)
	out.DbWorkload = direct.Enum_ToProto[pb.DBWorkload](mapCtx, in.DbWorkload)
	out.DbEdition = direct.Enum_ToProto[pb.AutonomousDatabaseProperties_DatabaseEdition](mapCtx, in.DbEdition)
	out.CharacterSet = direct.ValueOf(in.CharacterSet)
	out.NCharacterSet = direct.ValueOf(in.NCharacterSet)
	out.PrivateEndpointIp = direct.ValueOf(in.PrivateEndpointIP)
	out.PrivateEndpointLabel = direct.ValueOf(in.PrivateEndpointLabel)
	out.DbVersion = direct.ValueOf(in.DbVersion)
	out.IsAutoScalingEnabled = direct.ValueOf(in.IsAutoScalingEnabled)
	out.IsStorageAutoScalingEnabled = direct.ValueOf(in.IsStorageAutoScalingEnabled)
	out.LicenseType = direct.Enum_ToProto[pb.AutonomousDatabaseProperties_LicenseType](mapCtx, in.LicenseType)
	out.CustomerContacts = direct.Slice_ToProto(mapCtx, in.CustomerContacts, CustomerContact_ToProto)
	out.SecretId = direct.ValueOf(in.SecretID)
	out.VaultId = direct.ValueOf(in.VaultID)
	out.MaintenanceScheduleType = direct.Enum_ToProto[pb.AutonomousDatabaseProperties_MaintenanceScheduleType](mapCtx, in.MaintenanceScheduleType)
	out.MtlsConnectionRequired = direct.ValueOf(in.MtlsConnectionRequired)
	out.BackupRetentionPeriodDays = direct.ValueOf(in.BackupRetentionPeriodDays)
	// MISSING: ActualUsedDataStorageSizeTb
	// MISSING: AllocatedStorageSizeTb
	// MISSING: ApexDetails
	// MISSING: ArePrimaryAllowlistedIpsUsed
	// MISSING: LifecycleDetails
	// MISSING: State
	// MISSING: AutonomousContainerDatabaseID
	// MISSING: AvailableUpgradeVersions
	// MISSING: ConnectionStrings
	// MISSING: ConnectionUrls
	// MISSING: FailedDataRecoveryDuration
	// MISSING: MemoryTableGbs
	// MISSING: IsLocalDataGuardEnabled
	// MISSING: LocalAdgAutoFailoverMaxDataLossLimit
	// MISSING: LocalStandbyDb
	// MISSING: MemoryPerOracleComputeUnitGbs
	// MISSING: LocalDisasterRecoveryType
	// MISSING: DataSafeState
	// MISSING: DatabaseManagementState
	// MISSING: OpenMode
	// MISSING: OperationsInsightsState
	// MISSING: PeerDbIds
	// MISSING: PermissionLevel
	// MISSING: PrivateEndpoint
	// MISSING: RefreshableMode
	// MISSING: RefreshableState
	// MISSING: Role
	// MISSING: ScheduledOperationDetails
	// MISSING: SqlWebDeveloperURL
	// MISSING: SupportedCloneRegions
	// MISSING: UsedDataStorageSizeTbs
	// MISSING: OciURL
	// MISSING: TotalAutoBackupStorageSizeGbs
	// MISSING: NextLongTermBackupTime
	// MISSING: MaintenanceBeginTime
	// MISSING: MaintenanceEndTime
	return out
}
func AutonomousDatabasePropertiesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseProperties) *krm.AutonomousDatabasePropertiesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabasePropertiesObservedState{}
	out.Ocid = direct.LazyPtr(in.GetOcid())
	// MISSING: ComputeCount
	// MISSING: CpuCoreCount
	// MISSING: DataStorageSizeTb
	// MISSING: DataStorageSizeGB
	// MISSING: DbWorkload
	// MISSING: DbEdition
	// MISSING: CharacterSet
	// MISSING: NCharacterSet
	// MISSING: PrivateEndpointIP
	// MISSING: PrivateEndpointLabel
	// MISSING: DbVersion
	// MISSING: IsAutoScalingEnabled
	// MISSING: IsStorageAutoScalingEnabled
	// MISSING: LicenseType
	// MISSING: CustomerContacts
	// MISSING: SecretID
	// MISSING: VaultID
	// MISSING: MaintenanceScheduleType
	// MISSING: MtlsConnectionRequired
	// MISSING: BackupRetentionPeriodDays
	out.ActualUsedDataStorageSizeTb = direct.LazyPtr(in.GetActualUsedDataStorageSizeTb())
	out.AllocatedStorageSizeTb = direct.LazyPtr(in.GetAllocatedStorageSizeTb())
	out.ApexDetails = AutonomousDatabaseApex_FromProto(mapCtx, in.GetApexDetails())
	out.ArePrimaryAllowlistedIpsUsed = in.ArePrimaryAllowlistedIpsUsed
	out.LifecycleDetails = direct.LazyPtr(in.GetLifecycleDetails())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.AutonomousContainerDatabaseID = direct.LazyPtr(in.GetAutonomousContainerDatabaseId())
	out.AvailableUpgradeVersions = in.AvailableUpgradeVersions
	out.ConnectionStrings = AutonomousDatabaseConnectionStrings_FromProto(mapCtx, in.GetConnectionStrings())
	out.ConnectionUrls = AutonomousDatabaseConnectionUrls_FromProto(mapCtx, in.GetConnectionUrls())
	out.FailedDataRecoveryDuration = direct.StringDuration_FromProto(mapCtx, in.GetFailedDataRecoveryDuration())
	out.MemoryTableGbs = direct.LazyPtr(in.GetMemoryTableGbs())
	out.IsLocalDataGuardEnabled = direct.LazyPtr(in.GetIsLocalDataGuardEnabled())
	out.LocalAdgAutoFailoverMaxDataLossLimit = direct.LazyPtr(in.GetLocalAdgAutoFailoverMaxDataLossLimit())
	out.LocalStandbyDb = AutonomousDatabaseStandbySummary_FromProto(mapCtx, in.GetLocalStandbyDb())
	out.MemoryPerOracleComputeUnitGbs = direct.LazyPtr(in.GetMemoryPerOracleComputeUnitGbs())
	out.LocalDisasterRecoveryType = direct.Enum_FromProto(mapCtx, in.GetLocalDisasterRecoveryType())
	out.DataSafeState = direct.Enum_FromProto(mapCtx, in.GetDataSafeState())
	out.DatabaseManagementState = direct.Enum_FromProto(mapCtx, in.GetDatabaseManagementState())
	out.OpenMode = direct.Enum_FromProto(mapCtx, in.GetOpenMode())
	out.OperationsInsightsState = direct.Enum_FromProto(mapCtx, in.GetOperationsInsightsState())
	out.PeerDbIds = in.PeerDbIds
	out.PermissionLevel = direct.Enum_FromProto(mapCtx, in.GetPermissionLevel())
	out.PrivateEndpoint = direct.LazyPtr(in.GetPrivateEndpoint())
	out.RefreshableMode = direct.Enum_FromProto(mapCtx, in.GetRefreshableMode())
	out.RefreshableState = direct.Enum_FromProto(mapCtx, in.GetRefreshableState())
	out.Role = direct.Enum_FromProto(mapCtx, in.GetRole())
	out.ScheduledOperationDetails = direct.Slice_FromProto(mapCtx, in.ScheduledOperationDetails, ScheduledOperationDetails_FromProto)
	out.SqlWebDeveloperURL = direct.LazyPtr(in.GetSqlWebDeveloperUrl())
	out.SupportedCloneRegions = in.SupportedCloneRegions
	out.UsedDataStorageSizeTbs = direct.LazyPtr(in.GetUsedDataStorageSizeTbs())
	out.OciURL = direct.LazyPtr(in.GetOciUrl())
	out.TotalAutoBackupStorageSizeGbs = direct.LazyPtr(in.GetTotalAutoBackupStorageSizeGbs())
	out.NextLongTermBackupTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNextLongTermBackupTime())
	out.MaintenanceBeginTime = direct.StringTimestamp_FromProto(mapCtx, in.GetMaintenanceBeginTime())
	out.MaintenanceEndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetMaintenanceEndTime())
	return out
}
func AutonomousDatabasePropertiesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabasePropertiesObservedState) *pb.AutonomousDatabaseProperties {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseProperties{}
	out.Ocid = direct.ValueOf(in.Ocid)
	// MISSING: ComputeCount
	// MISSING: CpuCoreCount
	// MISSING: DataStorageSizeTb
	// MISSING: DataStorageSizeGB
	// MISSING: DbWorkload
	// MISSING: DbEdition
	// MISSING: CharacterSet
	// MISSING: NCharacterSet
	// MISSING: PrivateEndpointIP
	// MISSING: PrivateEndpointLabel
	// MISSING: DbVersion
	// MISSING: IsAutoScalingEnabled
	// MISSING: IsStorageAutoScalingEnabled
	// MISSING: LicenseType
	// MISSING: CustomerContacts
	// MISSING: SecretID
	// MISSING: VaultID
	// MISSING: MaintenanceScheduleType
	// MISSING: MtlsConnectionRequired
	// MISSING: BackupRetentionPeriodDays
	out.ActualUsedDataStorageSizeTb = direct.ValueOf(in.ActualUsedDataStorageSizeTb)
	out.AllocatedStorageSizeTb = direct.ValueOf(in.AllocatedStorageSizeTb)
	out.ApexDetails = AutonomousDatabaseApex_ToProto(mapCtx, in.ApexDetails)
	out.ArePrimaryAllowlistedIpsUsed = in.ArePrimaryAllowlistedIpsUsed
	out.LifecycleDetails = direct.ValueOf(in.LifecycleDetails)
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	out.AutonomousContainerDatabaseId = direct.ValueOf(in.AutonomousContainerDatabaseID)
	out.AvailableUpgradeVersions = in.AvailableUpgradeVersions
	out.ConnectionStrings = AutonomousDatabaseConnectionStrings_ToProto(mapCtx, in.ConnectionStrings)
	out.ConnectionUrls = AutonomousDatabaseConnectionUrls_ToProto(mapCtx, in.ConnectionUrls)
	out.FailedDataRecoveryDuration = direct.StringDuration_ToProto(mapCtx, in.FailedDataRecoveryDuration)
	out.MemoryTableGbs = direct.ValueOf(in.MemoryTableGbs)
	out.IsLocalDataGuardEnabled = direct.ValueOf(in.IsLocalDataGuardEnabled)
	out.LocalAdgAutoFailoverMaxDataLossLimit = direct.ValueOf(in.LocalAdgAutoFailoverMaxDataLossLimit)
	out.LocalStandbyDb = AutonomousDatabaseStandbySummary_ToProto(mapCtx, in.LocalStandbyDb)
	out.MemoryPerOracleComputeUnitGbs = direct.ValueOf(in.MemoryPerOracleComputeUnitGbs)
	out.LocalDisasterRecoveryType = direct.Enum_ToProto[pb.AutonomousDatabaseProperties_LocalDisasterRecoveryType](mapCtx, in.LocalDisasterRecoveryType)
	out.DataSafeState = direct.Enum_ToProto[pb.AutonomousDatabaseProperties_DataSafeState](mapCtx, in.DataSafeState)
	out.DatabaseManagementState = direct.Enum_ToProto[pb.AutonomousDatabaseProperties_DatabaseManagementState](mapCtx, in.DatabaseManagementState)
	out.OpenMode = direct.Enum_ToProto[pb.AutonomousDatabaseProperties_OpenMode](mapCtx, in.OpenMode)
	out.OperationsInsightsState = direct.Enum_ToProto[pb.OperationsInsightsState](mapCtx, in.OperationsInsightsState)
	out.PeerDbIds = in.PeerDbIds
	out.PermissionLevel = direct.Enum_ToProto[pb.AutonomousDatabaseProperties_PermissionLevel](mapCtx, in.PermissionLevel)
	out.PrivateEndpoint = direct.ValueOf(in.PrivateEndpoint)
	out.RefreshableMode = direct.Enum_ToProto[pb.AutonomousDatabaseProperties_RefreshableMode](mapCtx, in.RefreshableMode)
	out.RefreshableState = direct.Enum_ToProto[pb.AutonomousDatabaseProperties_RefreshableState](mapCtx, in.RefreshableState)
	out.Role = direct.Enum_ToProto[pb.AutonomousDatabaseProperties_Role](mapCtx, in.Role)
	out.ScheduledOperationDetails = direct.Slice_ToProto(mapCtx, in.ScheduledOperationDetails, ScheduledOperationDetails_ToProto)
	out.SqlWebDeveloperUrl = direct.ValueOf(in.SqlWebDeveloperURL)
	out.SupportedCloneRegions = in.SupportedCloneRegions
	out.UsedDataStorageSizeTbs = direct.ValueOf(in.UsedDataStorageSizeTbs)
	out.OciUrl = direct.ValueOf(in.OciURL)
	out.TotalAutoBackupStorageSizeGbs = direct.ValueOf(in.TotalAutoBackupStorageSizeGbs)
	out.NextLongTermBackupTime = direct.StringTimestamp_ToProto(mapCtx, in.NextLongTermBackupTime)
	out.MaintenanceBeginTime = direct.StringTimestamp_ToProto(mapCtx, in.MaintenanceBeginTime)
	out.MaintenanceEndTime = direct.StringTimestamp_ToProto(mapCtx, in.MaintenanceEndTime)
	return out
}
func AutonomousDatabaseStandbySummary_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseStandbySummary) *krm.AutonomousDatabaseStandbySummary {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseStandbySummary{}
	// MISSING: LagTimeDuration
	// MISSING: LifecycleDetails
	// MISSING: State
	// MISSING: DataGuardRoleChangedTime
	// MISSING: DisasterRecoveryRoleChangedTime
	return out
}
func AutonomousDatabaseStandbySummary_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseStandbySummary) *pb.AutonomousDatabaseStandbySummary {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseStandbySummary{}
	// MISSING: LagTimeDuration
	// MISSING: LifecycleDetails
	// MISSING: State
	// MISSING: DataGuardRoleChangedTime
	// MISSING: DisasterRecoveryRoleChangedTime
	return out
}
func AutonomousDatabaseStandbySummaryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabaseStandbySummary) *krm.AutonomousDatabaseStandbySummaryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutonomousDatabaseStandbySummaryObservedState{}
	out.LagTimeDuration = direct.StringDuration_FromProto(mapCtx, in.GetLagTimeDuration())
	out.LifecycleDetails = direct.LazyPtr(in.GetLifecycleDetails())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.DataGuardRoleChangedTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDataGuardRoleChangedTime())
	out.DisasterRecoveryRoleChangedTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDisasterRecoveryRoleChangedTime())
	return out
}
func AutonomousDatabaseStandbySummaryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutonomousDatabaseStandbySummaryObservedState) *pb.AutonomousDatabaseStandbySummary {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabaseStandbySummary{}
	out.LagTimeDuration = direct.StringDuration_ToProto(mapCtx, in.LagTimeDuration)
	out.LifecycleDetails = direct.ValueOf(in.LifecycleDetails)
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	out.DataGuardRoleChangedTime = direct.StringTimestamp_ToProto(mapCtx, in.DataGuardRoleChangedTime)
	out.DisasterRecoveryRoleChangedTime = direct.StringTimestamp_ToProto(mapCtx, in.DisasterRecoveryRoleChangedTime)
	return out
}
func CustomerContact_FromProto(mapCtx *direct.MapContext, in *pb.CustomerContact) *krm.CustomerContact {
	if in == nil {
		return nil
	}
	out := &krm.CustomerContact{}
	out.Email = direct.LazyPtr(in.GetEmail())
	return out
}
func CustomerContact_ToProto(mapCtx *direct.MapContext, in *krm.CustomerContact) *pb.CustomerContact {
	if in == nil {
		return nil
	}
	out := &pb.CustomerContact{}
	out.Email = direct.ValueOf(in.Email)
	return out
}
func DatabaseConnectionStringProfile_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseConnectionStringProfile) *krm.DatabaseConnectionStringProfile {
	if in == nil {
		return nil
	}
	out := &krm.DatabaseConnectionStringProfile{}
	// MISSING: ConsumerGroup
	// MISSING: DisplayName
	// MISSING: HostFormat
	// MISSING: IsRegional
	// MISSING: Protocol
	// MISSING: SessionMode
	// MISSING: SyntaxFormat
	// MISSING: TlsAuthentication
	// MISSING: Value
	return out
}
func DatabaseConnectionStringProfile_ToProto(mapCtx *direct.MapContext, in *krm.DatabaseConnectionStringProfile) *pb.DatabaseConnectionStringProfile {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseConnectionStringProfile{}
	// MISSING: ConsumerGroup
	// MISSING: DisplayName
	// MISSING: HostFormat
	// MISSING: IsRegional
	// MISSING: Protocol
	// MISSING: SessionMode
	// MISSING: SyntaxFormat
	// MISSING: TlsAuthentication
	// MISSING: Value
	return out
}
func DatabaseConnectionStringProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseConnectionStringProfile) *krm.DatabaseConnectionStringProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatabaseConnectionStringProfileObservedState{}
	out.ConsumerGroup = direct.Enum_FromProto(mapCtx, in.GetConsumerGroup())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.HostFormat = direct.Enum_FromProto(mapCtx, in.GetHostFormat())
	out.IsRegional = direct.LazyPtr(in.GetIsRegional())
	out.Protocol = direct.Enum_FromProto(mapCtx, in.GetProtocol())
	out.SessionMode = direct.Enum_FromProto(mapCtx, in.GetSessionMode())
	out.SyntaxFormat = direct.Enum_FromProto(mapCtx, in.GetSyntaxFormat())
	out.TlsAuthentication = direct.Enum_FromProto(mapCtx, in.GetTlsAuthentication())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func DatabaseConnectionStringProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatabaseConnectionStringProfileObservedState) *pb.DatabaseConnectionStringProfile {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseConnectionStringProfile{}
	out.ConsumerGroup = direct.Enum_ToProto[pb.DatabaseConnectionStringProfile_ConsumerGroup](mapCtx, in.ConsumerGroup)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.HostFormat = direct.Enum_ToProto[pb.DatabaseConnectionStringProfile_HostFormat](mapCtx, in.HostFormat)
	out.IsRegional = direct.ValueOf(in.IsRegional)
	out.Protocol = direct.Enum_ToProto[pb.DatabaseConnectionStringProfile_Protocol](mapCtx, in.Protocol)
	out.SessionMode = direct.Enum_ToProto[pb.DatabaseConnectionStringProfile_SessionMode](mapCtx, in.SessionMode)
	out.SyntaxFormat = direct.Enum_ToProto[pb.DatabaseConnectionStringProfile_SyntaxFormat](mapCtx, in.SyntaxFormat)
	out.TlsAuthentication = direct.Enum_ToProto[pb.DatabaseConnectionStringProfile_TLSAuthentication](mapCtx, in.TlsAuthentication)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func OracledatabaseAutonomousDatabaseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabase) *krm.OracledatabaseAutonomousDatabaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseAutonomousDatabaseObservedState{}
	// MISSING: Name
	// MISSING: Database
	// MISSING: DisplayName
	// MISSING: EntitlementID
	// MISSING: AdminPassword
	// MISSING: Properties
	// MISSING: Labels
	// MISSING: Network
	// MISSING: Cidr
	// MISSING: CreateTime
	return out
}
func OracledatabaseAutonomousDatabaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseAutonomousDatabaseObservedState) *pb.AutonomousDatabase {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabase{}
	// MISSING: Name
	// MISSING: Database
	// MISSING: DisplayName
	// MISSING: EntitlementID
	// MISSING: AdminPassword
	// MISSING: Properties
	// MISSING: Labels
	// MISSING: Network
	// MISSING: Cidr
	// MISSING: CreateTime
	return out
}
func OracledatabaseAutonomousDatabaseSpec_FromProto(mapCtx *direct.MapContext, in *pb.AutonomousDatabase) *krm.OracledatabaseAutonomousDatabaseSpec {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseAutonomousDatabaseSpec{}
	// MISSING: Name
	// MISSING: Database
	// MISSING: DisplayName
	// MISSING: EntitlementID
	// MISSING: AdminPassword
	// MISSING: Properties
	// MISSING: Labels
	// MISSING: Network
	// MISSING: Cidr
	// MISSING: CreateTime
	return out
}
func OracledatabaseAutonomousDatabaseSpec_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseAutonomousDatabaseSpec) *pb.AutonomousDatabase {
	if in == nil {
		return nil
	}
	out := &pb.AutonomousDatabase{}
	// MISSING: Name
	// MISSING: Database
	// MISSING: DisplayName
	// MISSING: EntitlementID
	// MISSING: AdminPassword
	// MISSING: Properties
	// MISSING: Labels
	// MISSING: Network
	// MISSING: Cidr
	// MISSING: CreateTime
	return out
}
func ScheduledOperationDetails_FromProto(mapCtx *direct.MapContext, in *pb.ScheduledOperationDetails) *krm.ScheduledOperationDetails {
	if in == nil {
		return nil
	}
	out := &krm.ScheduledOperationDetails{}
	// MISSING: DayOfWeek
	// MISSING: StartTime
	// MISSING: StopTime
	return out
}
func ScheduledOperationDetails_ToProto(mapCtx *direct.MapContext, in *krm.ScheduledOperationDetails) *pb.ScheduledOperationDetails {
	if in == nil {
		return nil
	}
	out := &pb.ScheduledOperationDetails{}
	// MISSING: DayOfWeek
	// MISSING: StartTime
	// MISSING: StopTime
	return out
}
func ScheduledOperationDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ScheduledOperationDetails) *krm.ScheduledOperationDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ScheduledOperationDetailsObservedState{}
	out.DayOfWeek = direct.Enum_FromProto(mapCtx, in.GetDayOfWeek())
	out.StartTime = TimeOfDay_FromProto(mapCtx, in.GetStartTime())
	out.StopTime = TimeOfDay_FromProto(mapCtx, in.GetStopTime())
	return out
}
func ScheduledOperationDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ScheduledOperationDetailsObservedState) *pb.ScheduledOperationDetails {
	if in == nil {
		return nil
	}
	out := &pb.ScheduledOperationDetails{}
	out.DayOfWeek = direct.Enum_ToProto[pb.DayOfWeek](mapCtx, in.DayOfWeek)
	out.StartTime = TimeOfDay_ToProto(mapCtx, in.StartTime)
	out.StopTime = TimeOfDay_ToProto(mapCtx, in.StopTime)
	return out
}
