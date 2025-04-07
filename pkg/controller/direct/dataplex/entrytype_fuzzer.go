// Copyright 2024 Google LLC
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

// +tool:fuzz-gen
// proto.message: google.cloud.dataplex.v1.EntryType
// api.group: dataplex.cnrm.cloud.google.com

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataplexEntryTypeFuzzer())
}

func dataplexEntryTypeFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.EntryType{},
		DataplexEntryTypeSpec_FromProto, DataplexEntryTypeSpec_ToProto,
		DataplexEntryTypeObservedState_FromProto, DataplexEntryTypeObservedState_ToProto,
	)

	f.SpecFields.Insert("description")
	f.SpecFields.Insert("display_name")
	f.SpecFields.Insert("labels")
	f.SpecFields.Insert("type_aliases")
	f.SpecFields.Insert("platform")
	f.SpecFields.Insert("system")
	f.SpecFields.Insert("required_aspects")
	f.SpecFields.Insert("authorization")

	f.StatusFields.Insert("name")
	f.StatusFields.Insert("uid")
	f.StatusFields.Insert("create_time")
	f.StatusFields.Insert("update_time")
	f.StatusFields.Insert("state")
	f.StatusFields.Insert(".required_aspects")
	f.StatusFields.Insert(".authorization")

	return f
}
func BackendMetastore_FromProto(mapCtx *direct.MapContext, in *pb.BackendMetastore) *krm.BackendMetastore {
	if in == nil {
		return nil
	}
	out := &krm.BackendMetastore{}
	if in.GetName() != "" {
		out.Name = in.GetName()
	}
	out.MetastoreType = direct.Enum_FromProto(mapCtx, in.GetMetastoreType())
	return out
}
func BackendMetastore_ToProto(mapCtx *direct.MapContext, in *krm.BackendMetastore) *pb.BackendMetastore {
	if in == nil {
		return nil
	}
	out := &pb.BackendMetastore{}
	out.Name = in.Name
	out.MetastoreType = direct.Enum_ToProto[pb.BackendMetastore_MetastoreType](mapCtx, in.MetastoreType)
	return out
}

```
</out>

</example>




<example>
<in.api.group>sql.cnrm.cloud.google.com</in.api.group>
<in.go.mappers>
func SQLServerAuditConfig_FromProto(mapCtx *direct.MapContext, in *pb.BackupConfiguration_BackupRetentionSettings) *krm.SQLServerAuditConfig {
	if in == nil {
		return nil
	}
	out := &krm.SQLServerAuditConfig{}
	// MISSING: AuditLogPath
	// MISSING: RetentionPeriod
	return out
}
func SQLServerAuditConfig_ToProto(mapCtx *direct.MapContext, in *krm.SQLServerAuditConfig) *pb.BackupConfiguration_BackupRetentionSettings {
	if in == nil {
		return nil
	}
	out := &pb.BackupConfiguration_BackupRetentionSettings{}
	out.RetentionUnit = pb.BackupRetentionSettings_RetentionUnit(pb.BackupRetentionSettings_RetentionUnit_value["COUNT"])
	out.RetainedBackups = in.RetentionPeriod
	return out
}
func SQLServerConfig_FromProto(mapCtx *direct.MapContext, in *pb.BackupConfiguration_BackupRetentionSettings) *krm.SQLServerConfig {
	if in == nil {
		return nil
	}
	out := &krm.SQLServerConfig{}
	// MISSING: CloudSQLVersion
	out.AuthorizedNetwork = direct.Slice_ToProto(mapCtx, in.AuthorizedNetwork)
	out.TimeZone = direct.LazyPtr(in.GetTimezone())
	out.DatabaseFlags = direct.Slice_ToProto(mapCtx, in.DatabaseFlags, DatabaseFlags_ToProto)
	out.BackupConfiguration = BackupConfiguration_FromProto(mapCtx, in.GetBackupConfiguration())
	out.DenyMaintenancePeriod = MaintenanceWindow_FromProto(mapCtx, in.GetDenyMaintenancePeriod())
	out.InstanceType = direct.Enum_FromProto(mapCtx, in.GetInstanceType())
	out.AvailabilityType = direct.Enum_FromProto(mapCtx, in.GetAvailabilityType())
	out.DiskSize = direct.LazyPtr(in.GetDiskSize())
	out.PricingPlan = direct.Enum_FromProto(mapCtx, in.GetPricingPlan())
	out.ActivationPolicy = direct.Enum_FromProto(mapCtx, in.GetActivationPolicy())
	out.StorageAutoResizeLimit = direct.LazyPtr(in.GetStorageAutoResizeLimit())
	out.DatabaseVersion = direct.LazyPtr(in.GetDatabaseVersion())
	out.Settings = InstanceSettings_FromProto(mapCtx, in.GetSettings())
	// MISSING: MasterInstanceName
	out.RootPasswordRef = direct.LazyPtr(in.GetRootPassword())
	out.InstanceName = direct.LazyPtr(in.GetName())
	out.Project = direct.LazyPtr(in.GetProject())
	// MISSING: ServiceAccountEmail
	out.SqlAdminNetwork = SQLAdminNetworkConfig_FromProto(mapCtx, in.GetSqlAdminNetwork())
	// MISSING: BackupRetentionSettings
	// MISSING: DeletionProtection
	return out
}
func SQLServerConfig_ToProto(mapCtx *direct.MapContext, in *krm.SQLServerConfig) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.DatabaseVersion = direct.ValueOf(in.DatabaseVersion)
	out.Settings = InstanceSettings_ToProto(mapCtx, in.Settings)
	if in.MasterInstanceName != nil {
		out.InstanceType = pb.Instance_InstanceType(krm.GetStringValue(in.MasterInstanceName))
	}
	if in.ServiceAccountRef != nil {
		out.ServiceAccountEmailAddress = direct.ValueOf(in.ServiceAccountRef.External)
	}
	return out
}
func BackupConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.BackupConfiguration) *krm.BackupConfiguration {
	if in == nil {
		return nil
	}
	out := &krm.BackupConfiguration{}
	out.PointInTimeRecoveryEnabled = direct.LazyPtr(in.GetPointInTimeRecoveryEnabled())
	out.BackupRetentionSettings = BackupConfiguration_BackupRetentionSettings_FromProto(mapCtx, in.GetBackupRetentionSettings())
	out.StartTime = direct.LazyPtr(in.GetStartTime())
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.BinaryLogEnabled = direct.LazyPtr(in.GetBinaryLogEnabled())
	out.TransactionLogRetentionDays = direct.LazyPtr(in.GetTransactionLogRetentionDays())
	out.TransactionalLogStorageState = direct.Enum_FromProto(mapCtx, in.GetTransactionalLogStorageState())
	out.RetainedBackupsCount = direct.LazyPtr(in.GetRetainedBackupsCount())
	out.LogRetentionDays = direct.LazyPtr(in.GetLogRetentionDays())
	return out
}
func BackupConfiguration_ToProto(mapCtx *direct.MapContext, in *krm.BackupConfiguration) *pb.BackupConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.BackupConfiguration{}
	out.PointInTimeRecoveryEnabled = direct.ValueOf(in.PointInTimeRecoveryEnabled)
	out.BackupRetentionSettings = BackupConfiguration_BackupRetentionSettings_ToProto(mapCtx, in.BackupRetentionSettings)
	out.StartTime = direct.ValueOf(in.StartTime)
	out.Enabled = direct.ValueOf(in.Enabled)
	out.BinaryLogEnabled = direct.ValueOf(in.BinaryLogEnabled)
	out.TransactionLogRetentionDays = direct.ValueOf(in.TransactionLogRetentionDays)
	out.RetainedBackupsCount = direct.ValueOf(in.RetainedBackupsCount)
	return out
}
func InstanceSettings_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.InstanceSettings {
	if in == nil {
		return nil
	}
	out := &krm.InstanceSettings{}
	out.ActivationPolicy = direct.Enum_FromProto(mapCtx, in.GetActivationPolicy())
	out.AuthorizedGaeApplication = in.AuthorizedGaeApplications
	out.AvailabilityType = direct.Enum_FromProto(mapCtx, in.GetAvailabilityType())
	out.DiskAutoresize = direct.LazyPtr(in.GetDiskAutoresize())
	out.DiskAutoresizeLimit = direct.LazyPtr(in.GetDiskAutoresizeLimit())
	out.DiskSize = direct.LazyPtr(in.GetDiskSize())
	out.DiskType = direct.LazyPtr(in.GetDiskType())
	out.PricingPlan = direct.Enum_FromProto(mapCtx, in.GetPricingPlan())
	out.ReplicationType = direct.Enum_FromProto(mapCtx, in.GetReplicationType())
	out.StorageAutoResizeLimit = direct.LazyPtr(in.GetStorageAutoResizeLimit())
	out.UserDataRetrievalTimeout = direct.StringDuration_FromProto(mapCtx, in.GetUserLabels())
	out.DataDiskType = direct.Enum_FromProto(mapCtx, in.GetDataDiskType())
	out.UserLabels = in.Labels
	out.IPConfiguration = IPConfiguration_FromProto(mapCtx, in.GetIpConfiguration())
	out.LocationPreference = LocationPreference_FromProto(mapCtx, in.GetLocationPreference())
	out.DatabaseFlags = direct.Slice_FromProto(mapCtx, in.GetDatabaseFlags(), DatabaseFlags_FromProto)
	out.DataCacheConfig = DataCacheConfig_FromProto(mapCtx, in.GetDataCacheConfig())
	out.BackupConfiguration = BackupConfiguration_FromProto(mapCtx, in.GetBackupConfiguration())
	out.DatabaseReplicationEnabled = direct.LazyPtr(in.GetDatabaseReplicationEnabled())
	out.PasswordValidationPolicy = PasswordValidationPolicy_FromProto(mapCtx, in.GetPasswordValidationPolicy())
	out.SQLServerAuditConfig = SQLServerAuditConfig_FromProto(mapCtx, in.GetSqlServerAuditConfig())
	return out
}
func InstanceSettings_ToProto(mapCtx *direct.MapContext, in *krm.InstanceSettings) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	out.SettingsVersion = direct.ValueOf(in.SettingsVersion)
	out.ActivationPolicy = direct.Enum_ToProto[pb.Settings_ActivationPolicy](mapCtx, in.ActivationPolicy)
	out.AuthorizedGaeApplications = in.AuthorizedGaeApplication
	out.AvailabilityType = direct.Enum_ToProto[pb.Settings_AvailabilityType](mapCtx, in.AvailabilityType)
	out.DiskAutoresize = direct.ValueOf(in.DiskAutoresize)
	out.DiskAutoresizeLimit = direct.ValueOf(in.DiskAutoresizeLimit)
	out.DiskSize = direct.ValueOf(in.DiskSize)
	out.DiskType = direct.Enum_ToProto[pb.Settings_DiskType](mapCtx, in.DiskType)
	out.PricingPlan = direct.Enum_ToProto[pb.Settings_PricingPlan](mapCtx, in.PricingPlan)
	out.ReplicationType = direct.Enum_ToProto[pb.Settings_ReplicationType](mapCtx, in.ReplicationType)
	out.StorageAutoResizeLimit = direct.ValueOf(in.StorageAutoResizeLimit)
	out.UserLabels = in.UserLabels
	out.IpConfiguration = IPConfiguration_ToProto(mapCtx, in.IPConfiguration)
	out.LocationPreference = LocationPreference_ToProto(mapCtx, in.LocationPreference)
	out.DatabaseFlags = direct.Slice_ToProto(mapCtx, in.DatabaseFlags, DatabaseFlags_ToProto)
	out.DataCacheConfig = DataCacheConfig_ToProto(mapCtx, in.DataCacheConfig)
	out.BackupConfiguration = BackupConfiguration_ToProto(mapCtx, in.BackupConfiguration)
	out.PasswordValidationPolicy = PasswordValidationPolicy_ToProto(mapCtx, in.PasswordValidationPolicy)
	out.SqlServerAuditConfig = SQLServerAuditConfig_ToProto(mapCtx, in.SQLServerAuditConfig)
	return out
}
func SQLServerAuditConfig_FromProto(mapCtx *direct.MapContext, in *pb.SQLServerAuditConfig) *krm.SQLServerAuditConfig {
	if in == nil {
		return nil
	}
	out := &krm.SQLServerAuditConfig{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.RetentionInterval = direct.StringDuration_FromProto(mapCtx, in.GetRetentionInterval())
	return out
}
func SQLServerAuditConfig_ToProto(mapCtx *direct.MapContext, in *krm.SQLServerAuditConfig) *pb.SQLServerAuditConfig {
	if in == nil {
		return nil
	}
	out := &pb.SQLServerAuditConfig{}
	out.Bucket = direct.ValueOf(in.Bucket)
	out.RetentionInterval = direct.StringDuration_ToProto(mapCtx, in.RetentionInterval)
	return out
}
func SQLServerDatabaseDetails_FromProto(mapCtx *direct.MapContext, in *pb.SqlServerDatabaseDetails) *krm.SQLServerDatabaseDetails {
	if in == nil {
		return nil
	}
	out := &krm.SQLServerDatabaseDetails{}
	out.CompatibilityLevel = direct.LazyPtr(in.GetCompatibilityLevel())
	out.RecoveryModel = direct.LazyPtr(in.GetRecoveryModel())
	return out
}
func SQLServerDatabaseDetails_ToProto(mapCtx *direct.MapContext, in *krm.SQLServerDatabaseDetails) *pb.SqlServerDatabaseDetails {
	if in == nil {
		return nil
	}
	out := &pb.SqlServerDatabaseDetails{}
	out.CompatibilityLevel = direct.ValueOf(in.CompatibilityLevel)
	out.RecoveryModel = direct.ValueOf(in.RecoveryModel)
	return out
}
func SpannerDatabaseBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.SpannerBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpannerBackupObservedState{}
	// MISSING: Name
	// MISSING: Database
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: ReferencingDatabases
	// MISSING: EncryptionInfo
	// MISSING: VersionTime
	// MISSING: MaxExpiredTime
	return out
}
func SpannerDatabaseBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpannerDatabaseBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Database = in.DatabaseRef.External
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	// MISSING: CreateTime
	// MISSING: SizeBytes
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	// MISSING: ReferencingDatabases
	// MISSING: EncryptionInfo
	return out
}
func SpannerDatabaseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.SpannerDatabaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpannerDatabaseObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.RestoreInfo = RestoreInfo_FromProto(mapCtx, in.GetRestoreInfo())
	out.VersionRetentionPeriod = direct.LazyPtr(in.GetVersionRetentionPeriod())
	out.EarliestVersionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEarliestVersionTime())
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	out.DefaultLeader = direct.LazyPtr(in.GetDefaultLeader())
	return out
}
func SpannerDatabaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpannerDatabaseObservedState) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.Database_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: RestoreInfo
	out.VersionRetentionPeriod = direct.ValueOf(in.VersionRetentionPeriod)
	out.EarliestVersionTime = direct.StringTimestamp_ToProto(mapCtx, in.EarliestVersionTime)
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	out.DefaultLeader = direct.ValueOf(in.DefaultLeader)
	return out
}
func RestoreInfo_FromProto(mapCtx *direct.MapContext, in *pb.RestoreInfo) *krm.RestoreInfo {
	if in == nil {
		return nil
	}
	out := &krm.RestoreInfo{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.BackupInfo = BackupInfo_FromProto(mapCtx, in.GetBackupInfo())
	return out
}
func RestoreInfo_ToProto(mapCtx *direct.MapContext, in *krm.RestoreInfo) *pb.RestoreInfo {
	if in == nil {
		return nil
	}
	out := &pb.RestoreInfo{}
	out.SourceType = direct.Enum_ToProto[pb.RestoreSourceType](mapCtx, in.SourceType)
	if oneof := BackupInfo_ToProto(mapCtx, in.BackupInfo); oneof != nil {
		out.SourceInfo = &pb.RestoreInfo_BackupInfo_{BackupInfo: oneof}
	}
	return out
}
func BackupInfo_FromProto(mapCtx *direct.MapContext, in *pb.BackupInfo) *krm.BackupInfo {
	if in == nil {
		return nil
	}
	out := &krm.BackupInfo{}
	out.Backup = direct.LazyPtr(in.GetBackup())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.SourceDatabase = direct.LazyPtr(in.GetSourceDatabase())
	out.VersionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVersionTime())
	return out
}
func BackupInfo_ToProto(mapCtx *direct.MapContext, in *krm.BackupInfo) *pb.BackupInfo {
	if in == nil {
		return nil
	}
	out := &pb.BackupInfo{}
	out.Backup = direct.ValueOf(in.Backup)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.SourceDatabase = direct.ValueOf(in.SourceDatabase)
	out.SourceBackup = direct.ValueOf(in.SourceBackup)
	return out
}


func InstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceConfig) *krm.InstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceConfigObservedState{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Replicas = direct.Slice_FromProto(mapCtx, in.Replicas, ReplicaInfo_FromProto)
	// MISSING: LeaderOptions
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func RestoreInfo_FromProto(mapCtx *direct.MapContext, in *pb.RestoreInfo) *krm.RestoreInfo {
	if in == nil {
		return nil
	}
	out := &krm.RestoreInfo{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.BackupInfo = BackupInfo_FromProto(mapCtx, in.GetBackupInfo())
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedState{}
	// MISSING: Name
	// MISSING: Config
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: NodeCount
	// MISSING: ProcessingUnits
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Labels = in.Labels
	// MISSING: EndpointUris
	// MISSING: CreateTime
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DisplayDevice = direct.LazyPtr(in.GetDisplayDevice())
	out.SoftwareVersion = direct.LazyPtr(in.GetSoftwareVersion())
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: Config
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: NodeCount
	// MISSING: ProcessingUnits
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.Labels = in.Labels
	// MISSING: EndpointUris
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DisplayDevice
	out.SoftwareVersion = direct.ValueOf(in.SoftwareVersion)
	return out
}
func InstanceRef_FromProto(mapCtx *direct.MapContext, in string) *krm.InstanceRef {
	if in == "" {
		return nil
	}
	out := &krm.InstanceRef{}
	out.External = in
	return out
}
func InstanceRef_ToProto(mapCtx *direct.MapContext, in *krm.InstanceRef) string {
	if in == nil {
		return ""
	}
	return in.External
}
func SpannerDatabaseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.SpannerDatabaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpannerDatabaseObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	out.RestoreInfo = RestoreInfo_FromProto(mapCtx, in.GetRestoreInfo())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.VersionRetentionPeriod = direct.LazyPtr(in.GetVersionRetentionPeriod())
	out.EarliestVersionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEarliestVersionTime())
	out.EncryptionConfig = EncryptionConfigObservedState_FromProto(mapCtx, in.GetEncryptionConfig())
	return out
}
func SpannerDatabaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpannerDatabaseObservedState) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	out.RestoreInfo = RestoreInfo_ToProto(mapCtx, in.RestoreInfo)
	out.State = direct.Enum_ToProto[pb.Database_State](mapCtx, in.State)
	// MISSING: VersionRetentionPeriod
	out.EarliestVersionTime = direct.StringTimestamp_ToProto(mapCtx, in.EarliestVersionTime)
	out.EncryptionConfig = EncryptionConfigObservedState_ToProto(mapCtx, in.EncryptionConfig)
	return out
}


func BackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupObservedState{}
	// MISSING: Name
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: CreateTime
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ReferencingDatabases = in.ReferencingDatabaseRefs
	out.EncryptionInfo = EncryptionInfo_FromProto(mapCtx, in.GetEncryptionInfo())
	out.VersionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVersionTime())
	return out
}
func BackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Database = direct.ValueOf(in.Database)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	// MISSING: CreateTime
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.ReferencingDatabases = in.ReferencingDatabases
	// MISSING: EncryptionInfo
	return out
}
func BackupInfo_FromProto(mapCtx *direct.MapContext, in *pb.BackupInfo) *krm.BackupInfo {
	if in == nil {
		return nil
	}
	out := &krm.BackupInfo{}
	out.Backup = direct.LazyPtr(in.GetBackup())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.SourceDatabase = direct.LazyPtr(in.GetSourceDatabase())
	out.VersionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVersionTime())
	return out
}
func BackupInfo_ToProto(mapCtx *direct.MapContext, in *krm.BackupInfo) *pb.BackupInfo {
	if in == nil {
		return nil
	}
	out := &pb.BackupInfo{}
	out.Backup = direct.ValueOf(in.Backup)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.SourceDatabase = direct.ValueOf(in.SourceDatabase)
	return out
}

func DataplexLakeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Lake) *krm.DataplexLakeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexLakeObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	out.Metastore = MetadataManagementActivityObservedState_FromProto(mapCtx, in.GetMetastore())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: AssetStatus
	return out
}
func DataplexLakeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexLakeObservedState) *pb.Lake {
	if in == nil {
		return nil
	}
	out := &pb.Lake{}
	// MISSING: Name
	// MISSING: CreateTime
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	out.Metastore = MetadataManagementActivityObservedState_ToProto(mapCtx, in.Metastore)
	out.State = direct.Enum_ToProto[pb.Lake_State](mapCtx, in.State)
	return out
}

func InstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceConfigObservedState) *pb.InstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceConfig{}
	// MISSING: Name
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func ManagedFolderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagedFolder) *krm.ManagedFolderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagedFolderObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: State
	return out
}
func ManagedFolderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagedFolderObservedState) *pb.ManagedFolder {
	if in == nil {
		return nil
	}
	out := &pb.ManagedFolder{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func VpcConfig_FromProto(mapCtx *direct.MapContext, in *pb.VpcConfig) *krm.VpcConfig {
	if in == nil {
		return nil
	}
	out := &krm.VpcConfig{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	return out
}
func VpcConfig_ToProto(mapCtx *direct.MapContext, in *krm.VpcConfig) *pb.VpcConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcConfig{}
	out.Network = direct.ValueOf(in.Network)
	return out
}
func VpcConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VpcConfig) *krm.VpcConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VpcConfigObservedState{}
	// MISSING: Network
	out.Subnet = direct.LazyPtr(in.GetSubnet())
	return out
}
func VpcConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VpcConfigObservedState) *pb.VpcConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcConfig{}
	// MISSING: Network
	out.Subnet = direct.ValueOf(in.Subnet)
	return out
}

