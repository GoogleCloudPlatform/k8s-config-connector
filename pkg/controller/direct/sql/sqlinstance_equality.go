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

package sql

import (
	"reflect"
	"sort"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	api "google.golang.org/api/sqladmin/v1beta4"
)

func DiffInstances(desired *api.DatabaseInstance, actual *api.DatabaseInstance) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.DatabaseInstance{}
	}
	if actual == nil {
		actual = &api.DatabaseInstance{}
	}

	if desired.DatabaseVersion != actual.DatabaseVersion {
		diff.AddField(".databaseVersion", actual.DatabaseVersion, desired.DatabaseVersion)
	}
	diff.AddDiff(DiffDiskEncryptionConfiguration(desired.DiskEncryptionConfiguration, actual.DiskEncryptionConfiguration))
	// Ignore GeminiConfig. It is not supported in KRM API.
	if desired.InstanceType != actual.InstanceType {
		diff.AddField(".instanceType", actual.InstanceType, desired.InstanceType)
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.MaintenanceVersion != actual.MaintenanceVersion {
		diff.AddField(".maintenanceVersion", actual.MaintenanceVersion, desired.MaintenanceVersion)
	}
	if desired.MasterInstanceName != actual.MasterInstanceName {
		diff.AddField(".masterInstanceName", actual.MasterInstanceName, desired.MasterInstanceName)
	}
	// Ignore MaxDiskSize. It is not supported in KRM API.
	if desired.Name != actual.Name {
		diff.AddField(".name", actual.Name, desired.Name)
	}
	// Ignore OnPremisesConfiguration. It is not supported in KRM API.
	if desired.Region != actual.Region {
		diff.AddField(".region", actual.Region, desired.Region)
	}
	diff.AddDiff(DiffReplicaConfiguration(desired.ReplicaConfiguration, actual.ReplicaConfiguration))
	diff.AddDiff(DiffReplicationCluster(desired.ReplicationCluster, actual.ReplicationCluster))
	// Ignore RootPassword. It is not exported.
	diff.AddDiff(DiffSettings(desired.Settings, actual.Settings))

	// Ignore SqlNetworkArchitecture. It is not supported in KRM API.
	// Ignore SwitchTransactionLogsToCloudStorageEnabled. It is not supported in KRM API.
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffDiskEncryptionConfiguration(desired *api.DiskEncryptionConfiguration, actual *api.DiskEncryptionConfiguration) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.DiskEncryptionConfiguration{}
	}
	if actual == nil {
		actual = &api.DiskEncryptionConfiguration{}
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.KmsKeyName != actual.KmsKeyName {
		diff.AddField(".diskEncryptionConfiguration.kmsKeyName", actual.KmsKeyName, desired.KmsKeyName)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffReplicaConfiguration(desired *api.ReplicaConfiguration, actual *api.ReplicaConfiguration) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.ReplicaConfiguration{}
	}
	if actual == nil {
		actual = &api.ReplicaConfiguration{}
	}
	// Ignore CascadableReplica. It is not supported in KRM API.
	if desired.FailoverTarget != actual.FailoverTarget {
		diff.AddField(".replicaConfiguration.failoverTarget", actual.FailoverTarget, desired.FailoverTarget)
	}
	// Ignore Kind. It is sometimes not set in API responses.
	diff.AddDiff(DiffMysqlReplicaConfiguration(desired.MysqlReplicaConfiguration, actual.MysqlReplicaConfiguration))
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffSettings(desired *api.Settings, actual *api.Settings) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.Settings{}
	}
	if actual == nil {
		actual = &api.Settings{}
	}
	if desired.ActivationPolicy != actual.ActivationPolicy {
		diff.AddField(".settings.activationPolicy", actual.ActivationPolicy, desired.ActivationPolicy)
	}
	diff.AddDiff(DiffActiveDirectoryConfig(desired.ActiveDirectoryConfig, actual.ActiveDirectoryConfig))
	diff.AddDiff(DiffAdvancedMachineFeatures(desired.AdvancedMachineFeatures, actual.AdvancedMachineFeatures))
	if !slicesMatch(desired.AuthorizedGaeApplications, actual.AuthorizedGaeApplications) {
		diff.AddField(".settings.authorizedGaeApplications", actual.AuthorizedGaeApplications, desired.AuthorizedGaeApplications)
	}
	if desired.AvailabilityType != actual.AvailabilityType {
		diff.AddField(".settings.availabilityType", actual.AvailabilityType, desired.AvailabilityType)
	}
	diff.AddDiff(DiffBackupConfiguration(desired.BackupConfiguration, actual.BackupConfiguration))
	if desired.Collation != actual.Collation {
		diff.AddField(".settings.collation", actual.Collation, desired.Collation)
	}
	if desired.ConnectorEnforcement != actual.ConnectorEnforcement {
		diff.AddField(".settings.connectorEnforcement", actual.ConnectorEnforcement, desired.ConnectorEnforcement)
	}
	// Ignore CrashSafeReplicationEnabled. It is only applicable to first-gen instances.
	diff.AddDiff(DiffDataCacheConfig(desired.DataCacheConfig, actual.DataCacheConfig))
	if desired.DataDiskSizeGb != actual.DataDiskSizeGb {
		diff.AddField(".settings.dataDiskSizeGb", actual.DataDiskSizeGb, desired.DataDiskSizeGb)
	}
	if desired.DataDiskType != actual.DataDiskType {
		diff.AddField(".settings.dataDiskType", actual.DataDiskType, desired.DataDiskType)
	}
	diff.AddDiff(DiffDatabaseFlagLists(desired.DatabaseFlags, actual.DatabaseFlags))
	// Ignore DatabaseReplicationEnabled. It is not supported in KRM API.
	if desired.DeletionProtectionEnabled != actual.DeletionProtectionEnabled {
		diff.AddField(".settings.deletionProtectionEnabled", actual.DeletionProtectionEnabled, desired.DeletionProtectionEnabled)
	}
	diff.AddDiff(DiffDenyMaintenancePeriodLists(desired.DenyMaintenancePeriods, actual.DenyMaintenancePeriods))
	if desired.Edition != actual.Edition {
		diff.AddField(".settings.edition", actual.Edition, desired.Edition)
	}
	// Ignore EnableDataplexIntegration. It is not supported in KRM API.
	// Ignore EnableGoogleMlIntegration. It is not supported in KRM API.
	diff.AddDiff(DiffInsightsConfig(desired.InsightsConfig, actual.InsightsConfig))
	diff.AddDiff(DiffIpConfiguration(desired.IpConfiguration, actual.IpConfiguration))
	// Ignore Kind. It is sometimes not set in API responses.
	diff.AddDiff(DiffLocationPreference(desired.LocationPreference, actual.LocationPreference))
	diff.AddDiff(DiffMaintenanceWindow(desired.MaintenanceWindow, actual.MaintenanceWindow))
	diff.AddDiff(DiffPasswordValidationPolicy(desired.PasswordValidationPolicy, actual.PasswordValidationPolicy))
	if desired.PricingPlan != actual.PricingPlan {
		diff.AddField(".settings.pricingPlan", actual.PricingPlan, desired.PricingPlan)
	}
	if desired.ReplicationType != actual.ReplicationType {
		diff.AddField(".settings.replicationType", actual.ReplicationType, desired.ReplicationType)
	}
	if desired.SettingsVersion != actual.SettingsVersion {
		diff.AddField(".settings.settingsVersion", actual.SettingsVersion, desired.SettingsVersion)
	}
	diff.AddDiff(DiffSqlServerAuditConfig(desired.SqlServerAuditConfig, actual.SqlServerAuditConfig))
	diff.AddDiff(DiffStorageAutoResize(desired.StorageAutoResize, actual.StorageAutoResize))
	if desired.StorageAutoResizeLimit != actual.StorageAutoResizeLimit {
		diff.AddField(".settings.storageAutoResizeLimit", actual.StorageAutoResizeLimit, desired.StorageAutoResizeLimit)
	}
	if desired.Tier != actual.Tier {
		diff.AddField(".settings.tier", actual.Tier, desired.Tier)
	}
	if desired.TimeZone != actual.TimeZone {
		diff.AddField(".settings.timeZone", actual.TimeZone, desired.TimeZone)
	}
	if !mapsMatch(desired.UserLabels, actual.UserLabels) {
		diff.AddField(".settings.userLabels", actual.UserLabels, desired.UserLabels)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

// slicesMatch checks if two slices are equal, matching with reflect.DeepEqual.
// As a special-case, the empty slice is treated the same as the nil slice
func slicesMatch[T any](desired []T, actual []T) bool {
	if len(desired) != len(actual) {
		return false
	}
	if len(desired) == 0 && len(actual) == 0 {
		return true
	}
	return reflect.DeepEqual(desired, actual)
}

// mapsMatch checks if two maps are equal, matching with reflect.DeepEqual.
// As a special-case, the empty map is treated the same as the nil map
func mapsMatch[K comparable, V any](desired map[K]V, actual map[K]V) bool {
	if len(desired) != len(actual) {
		return false
	}
	if len(desired) == 0 && len(actual) == 0 {
		return true
	}
	return reflect.DeepEqual(desired, actual)
}

func DiffMysqlReplicaConfiguration(desired *api.MySqlReplicaConfiguration, actual *api.MySqlReplicaConfiguration) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.MySqlReplicaConfiguration{}
	}
	if actual == nil {
		actual = &api.MySqlReplicaConfiguration{}
	}
	if desired.CaCertificate != actual.CaCertificate {
		diff.AddField(".replicaConfiguration.mysqlReplicaConfiguration.caCertificate", actual.CaCertificate, desired.CaCertificate)
	}
	if desired.ClientCertificate != actual.ClientCertificate {
		diff.AddField(".replicaConfiguration.mysqlReplicaConfiguration.clientCertificate", actual.ClientCertificate, desired.ClientCertificate)
	}
	if desired.ClientKey != actual.ClientKey {
		diff.AddField(".replicaConfiguration.mysqlReplicaConfiguration.clientKey", actual.ClientKey, desired.ClientKey)
	}
	if desired.ConnectRetryInterval != actual.ConnectRetryInterval {
		diff.AddField(".replicaConfiguration.mysqlReplicaConfiguration.connectRetryInterval", actual.ConnectRetryInterval, desired.ConnectRetryInterval)
	}
	if desired.DumpFilePath != actual.DumpFilePath {
		diff.AddField(".replicaConfiguration.mysqlReplicaConfiguration.dumpFilePath", actual.DumpFilePath, desired.DumpFilePath)
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.MasterHeartbeatPeriod != actual.MasterHeartbeatPeriod {
		diff.AddField(".replicaConfiguration.mysqlReplicaConfiguration.masterHeartbeatPeriod", actual.MasterHeartbeatPeriod, desired.MasterHeartbeatPeriod)
	}
	// Ignore Password. It is not exported.
	if desired.SslCipher != actual.SslCipher {
		diff.AddField(".replicaConfiguration.mysqlReplicaConfiguration.sslCipher", actual.SslCipher, desired.SslCipher)
	}
	if desired.Username != actual.Username {
		diff.AddField(".replicaConfiguration.mysqlReplicaConfiguration.username", actual.Username, desired.Username)
	}
	if desired.VerifyServerCertificate != actual.VerifyServerCertificate {
		diff.AddField(".replicaConfiguration.mysqlReplicaConfiguration.verifyServerCertificate", actual.VerifyServerCertificate, desired.VerifyServerCertificate)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffActiveDirectoryConfig(desired *api.SqlActiveDirectoryConfig, actual *api.SqlActiveDirectoryConfig) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.SqlActiveDirectoryConfig{}
	}
	if actual == nil {
		actual = &api.SqlActiveDirectoryConfig{}
	}
	if desired.Domain != actual.Domain {
		diff.AddField(".settings.activeDirectoryConfig.domain", actual.Domain, desired.Domain)
	}
	// Ignore Kind. It is sometimes not set in API responses.
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffAdvancedMachineFeatures(desired *api.AdvancedMachineFeatures, actual *api.AdvancedMachineFeatures) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.AdvancedMachineFeatures{}
	}
	if actual == nil {
		actual = &api.AdvancedMachineFeatures{}
	}
	if desired.ThreadsPerCore != actual.ThreadsPerCore {
		diff.AddField(".settings.advancedMachineFeatures.threadsPerCore", actual.ThreadsPerCore, desired.ThreadsPerCore)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffBackupConfiguration(desired *api.BackupConfiguration, actual *api.BackupConfiguration) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.BackupConfiguration{}
	}
	if actual == nil {
		actual = &api.BackupConfiguration{}
	}
	diff.AddDiff(DiffBackupRetentionSettings(desired.BackupRetentionSettings, actual.BackupRetentionSettings))
	if desired.BinaryLogEnabled != actual.BinaryLogEnabled {
		diff.AddField(".settings.backupConfiguration.binaryLogEnabled", actual.BinaryLogEnabled, desired.BinaryLogEnabled)
	}
	if desired.Enabled != actual.Enabled {
		diff.AddField(".settings.backupConfiguration.enabled", actual.Enabled, desired.Enabled)
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.Location != actual.Location {
		diff.AddField(".settings.backupConfiguration.location", actual.Location, desired.Location)
	}
	if desired.PointInTimeRecoveryEnabled != actual.PointInTimeRecoveryEnabled {
		diff.AddField(".settings.backupConfiguration.pointInTimeRecoveryEnabled", actual.PointInTimeRecoveryEnabled, desired.PointInTimeRecoveryEnabled)
	}
	// Ignore StartTime if it is not set. empty string is not a valid start time.
	if desired.StartTime != "" && desired.StartTime != actual.StartTime {
		diff.AddField(".settings.backupConfiguration.startTime", actual.StartTime, desired.StartTime)
	}
	// Ignore TransactionLogRetentionDays if it is not set. 0 is not a valid transaction log retention days.
	if desired.TransactionLogRetentionDays != 0 && desired.TransactionLogRetentionDays != actual.TransactionLogRetentionDays {
		diff.AddField(".settings.backupConfiguration.transactionLogRetentionDays", actual.TransactionLogRetentionDays, desired.TransactionLogRetentionDays)
	}

	// Ignore ReplicationLogArchivingEnabled. It is not supported in KRM API.
	// Ignore TransactionalLogStorageState. It is not supported in KRM API.
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffBackupRetentionSettings(desired *api.BackupRetentionSettings, actual *api.BackupRetentionSettings) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.BackupRetentionSettings{}
	}
	if actual == nil {
		actual = &api.BackupRetentionSettings{}
	}
	if desired.RetainedBackups != actual.RetainedBackups {
		diff.AddField(".settings.backupConfiguration.backupRetentionSettings.retainedBackups", actual.RetainedBackups, desired.RetainedBackups)
	}
	if desired.RetentionUnit != actual.RetentionUnit {
		diff.AddField(".settings.backupConfiguration.backupRetentionSettings.retentionUnit", actual.RetentionUnit, desired.RetentionUnit)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffDataCacheConfig(desired *api.DataCacheConfig, actual *api.DataCacheConfig) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	// GCP omits DataCacheConfig when DataCacheEnabled is false.
	// Treat nil and false as equivalent.
	desiredEnabled := false
	if desired != nil {
		desiredEnabled = desired.DataCacheEnabled
	}
	actualEnabled := false
	if actual != nil {
		actualEnabled = actual.DataCacheEnabled
	}
	if desiredEnabled != actualEnabled {
		diff.AddField(".settings.dataCacheConfig.dataCacheEnabled", actualEnabled, desiredEnabled)
	}
	return diff
}

// DatabaseFlagsByName implements sort.Interface for []*api.DatabaseFlags based on the Name field.
type DatabaseFlagsByName []*api.DatabaseFlags

func (a DatabaseFlagsByName) Len() int           { return len(a) }
func (a DatabaseFlagsByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a DatabaseFlagsByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func DiffDatabaseFlagLists(desired []*api.DatabaseFlags, actual []*api.DatabaseFlags) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if len(desired) != len(actual) {
		diff.AddField(".settings.databaseFlags", actual, desired)
		return diff
	}
	// Copy and sort so we don't modify the arguments
	desiredSorted := make([]*api.DatabaseFlags, len(desired))
	copy(desiredSorted, desired)
	sort.Sort(DatabaseFlagsByName(desiredSorted))

	actualSorted := make([]*api.DatabaseFlags, len(actual))
	copy(actualSorted, actual)
	sort.Sort(DatabaseFlagsByName(actualSorted))

	for i := 0; i < len(desiredSorted); i++ {
		diff.AddDiff(DiffDatabaseFlags(desiredSorted[i], actualSorted[i]))
	}
	return diff
}

func DiffDatabaseFlags(desired *api.DatabaseFlags, actual *api.DatabaseFlags) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.DatabaseFlags{}
	}
	if actual == nil {
		actual = &api.DatabaseFlags{}
	}
	if desired.Name != actual.Name {
		diff.AddField(".settings.databaseFlags.name", actual.Name, desired.Name)
	}
	if desired.Value != actual.Value {
		diff.AddField(".settings.databaseFlags.value", actual.Value, desired.Value)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffDenyMaintenancePeriodLists(desired []*api.DenyMaintenancePeriod, actual []*api.DenyMaintenancePeriod) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if len(desired) != len(actual) {
		diff.AddField(".settings.denyMaintenancePeriods", actual, desired)
		return diff
	}
	for i := 0; i < len(desired); i++ {
		diff.AddDiff(DiffDenyMaintenancePeriods(desired[i], actual[i]))
	}
	return diff
}

func DiffDenyMaintenancePeriods(desired *api.DenyMaintenancePeriod, actual *api.DenyMaintenancePeriod) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.DenyMaintenancePeriod{}
	}
	if actual == nil {
		actual = &api.DenyMaintenancePeriod{}
	}
	if desired.EndDate != actual.EndDate {
		diff.AddField(".settings.denyMaintenancePeriods.endDate", actual.EndDate, desired.EndDate)
	}
	if desired.StartDate != actual.StartDate {
		diff.AddField(".settings.denyMaintenancePeriods.startDate", actual.StartDate, desired.StartDate)
	}
	if desired.Time != actual.Time {
		diff.AddField(".settings.denyMaintenancePeriods.time", actual.Time, desired.Time)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffInsightsConfig(desired *api.InsightsConfig, actual *api.InsightsConfig) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.InsightsConfig{}
	}
	if actual == nil {
		actual = &api.InsightsConfig{}
	}
	if desired.QueryInsightsEnabled != actual.QueryInsightsEnabled {
		diff.AddField(".settings.insightsConfig.queryInsightsEnabled", actual.QueryInsightsEnabled, desired.QueryInsightsEnabled)
	}
	// Defaults to 5.
	desiredQueryPlansPerMinute := int64(5)
	if desired.QueryPlansPerMinute != 0 {
		desiredQueryPlansPerMinute = desired.QueryPlansPerMinute
	}
	actualQueryPlansPerMinute := int64(5)
	if actual.QueryPlansPerMinute != 0 {
		actualQueryPlansPerMinute = actual.QueryPlansPerMinute
	}
	if desiredQueryPlansPerMinute != actualQueryPlansPerMinute {
		diff.AddField(".settings.insightsConfig.queryPlansPerMinute", actualQueryPlansPerMinute, desiredQueryPlansPerMinute)
	}
	// Defaults to 1024.
	desiredQueryStringLength := int64(1024)
	if desired.QueryStringLength != 0 {
		desiredQueryStringLength = desired.QueryStringLength
	}
	actualQueryStringLength := int64(1024)
	if actual.QueryStringLength != 0 {
		actualQueryStringLength = actual.QueryStringLength
	}
	if desiredQueryStringLength != actualQueryStringLength {
		diff.AddField(".settings.insightsConfig.queryStringLength", actualQueryStringLength, desiredQueryStringLength)
	}
	if desired.RecordApplicationTags != actual.RecordApplicationTags {
		diff.AddField(".settings.insightsConfig.recordApplicationTags", actual.RecordApplicationTags, desired.RecordApplicationTags)
	}
	if desired.RecordClientAddress != actual.RecordClientAddress {
		diff.AddField(".settings.insightsConfig.recordClientAddress", actual.RecordClientAddress, desired.RecordClientAddress)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffIpConfiguration(desired *api.IpConfiguration, actual *api.IpConfiguration) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.IpConfiguration{
			Ipv4Enabled: true,
			SslMode:     "ALLOW_UNENCRYPTED_AND_ENCRYPTED",
		}
	}
	if actual == nil {
		actual = &api.IpConfiguration{
			Ipv4Enabled: true,
			SslMode:     "ALLOW_UNENCRYPTED_AND_ENCRYPTED",
		}
	}
	if desired.AllocatedIpRange != actual.AllocatedIpRange {
		diff.AddField(".settings.ipConfiguration.allocatedIpRange", actual.AllocatedIpRange, desired.AllocatedIpRange)
	}
	diff.AddDiff(DiffAclEntryLists(desired.AuthorizedNetworks, actual.AuthorizedNetworks))
	if desired.EnablePrivatePathForGoogleCloudServices != actual.EnablePrivatePathForGoogleCloudServices {
		diff.AddField(".settings.ipConfiguration.enablePrivatePathForGoogleCloudServices", actual.EnablePrivatePathForGoogleCloudServices, desired.EnablePrivatePathForGoogleCloudServices)
	}
	if desired.Ipv4Enabled != actual.Ipv4Enabled {
		diff.AddField(".settings.ipConfiguration.ipv4Enabled", actual.Ipv4Enabled, desired.Ipv4Enabled)
	}
	if desired.PrivateNetwork != actual.PrivateNetwork {
		diff.AddField(".settings.ipConfiguration.privateNetwork", actual.PrivateNetwork, desired.PrivateNetwork)
	}
	diff.AddDiff(DiffPscConfig(desired.PscConfig, actual.PscConfig))
	if desired.RequireSsl != actual.RequireSsl {
		diff.AddField(".settings.ipConfiguration.requireSsl", actual.RequireSsl, desired.RequireSsl)
	}
	// Ignore ServerCaMode. It is not supported in KRM API.
	if desired.SslMode != actual.SslMode {
		diff.AddField(".settings.ipConfiguration.sslMode", actual.SslMode, desired.SslMode)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

// AclEntriesByName implements sort.Interface for []*api.AclEntry based on the Name field.
type AclEntriesByName []*api.AclEntry

func (a AclEntriesByName) Len() int      { return len(a) }
func (a AclEntriesByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a AclEntriesByName) Less(i, j int) bool {
	if a[i].Name != a[j].Name {
		return a[i].Name < a[j].Name
	}
	return a[i].Value < a[j].Value
}

func DiffAclEntryLists(desired []*api.AclEntry, actual []*api.AclEntry) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if len(desired) != len(actual) {
		diff.AddField(".settings.ipConfiguration.authorizedNetworks", actual, desired)
		return diff
	}
	// Copy and sort so we don't modify the arguments
	desiredSorted := make([]*api.AclEntry, len(desired))
	copy(desiredSorted, desired)
	sort.Sort(AclEntriesByName(desiredSorted))

	actualSorted := make([]*api.AclEntry, len(actual))
	copy(actualSorted, actual)
	sort.Sort(AclEntriesByName(actualSorted))

	for i := 0; i < len(desiredSorted); i++ {
		diff.AddDiff(DiffAclEntries(desiredSorted[i], actualSorted[i]))
	}
	return diff
}

func DiffAclEntries(desired *api.AclEntry, actual *api.AclEntry) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.AclEntry{}
	}
	if actual == nil {
		actual = &api.AclEntry{}
	}
	if desired.ExpirationTime != actual.ExpirationTime {
		diff.AddField(".settings.ipConfiguration.authorizedNetworks.expirationTime", actual.ExpirationTime, desired.ExpirationTime)
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.Name != actual.Name {
		diff.AddField(".settings.ipConfiguration.authorizedNetworks.name", actual.Name, desired.Name)
	}
	if desired.Value != actual.Value {
		diff.AddField(".settings.ipConfiguration.authorizedNetworks.value", actual.Value, desired.Value)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffPscConfig(desired *api.PscConfig, actual *api.PscConfig) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.PscConfig{}
	}
	if actual == nil {
		actual = &api.PscConfig{}
	}
	if !slicesMatch(desired.AllowedConsumerProjects, actual.AllowedConsumerProjects) {
		diff.AddField(".settings.ipConfiguration.pscConfig.allowedConsumerProjects", actual.AllowedConsumerProjects, desired.AllowedConsumerProjects)
	}
	if desired.PscEnabled != actual.PscEnabled {
		diff.AddField(".settings.ipConfiguration.pscConfig.pscEnabled", actual.PscEnabled, desired.PscEnabled)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffLocationPreference(desired *api.LocationPreference, actual *api.LocationPreference) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.LocationPreference{}
	}
	if actual == nil {
		actual = &api.LocationPreference{}
	}
	if desired.FollowGaeApplication != actual.FollowGaeApplication {
		diff.AddField(".settings.locationPreference.followGaeApplication", actual.FollowGaeApplication, desired.FollowGaeApplication)
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.SecondaryZone != actual.SecondaryZone {
		diff.AddField(".settings.locationPreference.secondaryZone", actual.SecondaryZone, desired.SecondaryZone)
	}
	if desired.Zone != actual.Zone {
		diff.AddField(".settings.locationPreference.zone", actual.Zone, desired.Zone)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffMaintenanceWindow(desired *api.MaintenanceWindow, actual *api.MaintenanceWindow) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.MaintenanceWindow{}
	}
	if actual == nil {
		actual = &api.MaintenanceWindow{}
	}
	if desired.Day != actual.Day {
		diff.AddField(".settings.maintenanceWindow.day", actual.Day, desired.Day)
	}
	if desired.Hour != actual.Hour {
		diff.AddField(".settings.maintenanceWindow.hour", actual.Hour, desired.Hour)
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.UpdateTrack != actual.UpdateTrack {
		diff.AddField(".settings.maintenanceWindow.updateTrack", actual.UpdateTrack, desired.UpdateTrack)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffPasswordValidationPolicy(desired *api.PasswordValidationPolicy, actual *api.PasswordValidationPolicy) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.PasswordValidationPolicy{}
	}
	if actual == nil {
		actual = &api.PasswordValidationPolicy{}
	}
	if desired.Complexity != actual.Complexity {
		diff.AddField(".settings.passwordValidationPolicy.complexity", actual.Complexity, desired.Complexity)
	}
	// Ignore DisallowCompromisedCredentials. It is not supported in KRM API.
	if desired.DisallowUsernameSubstring != actual.DisallowUsernameSubstring {
		diff.AddField(".settings.passwordValidationPolicy.disallowUsernameSubstring", actual.DisallowUsernameSubstring, desired.DisallowUsernameSubstring)
	}
	if desired.EnablePasswordPolicy != actual.EnablePasswordPolicy {
		diff.AddField(".settings.passwordValidationPolicy.enablePasswordPolicy", actual.EnablePasswordPolicy, desired.EnablePasswordPolicy)
	}
	if desired.MinLength != actual.MinLength {
		diff.AddField(".settings.passwordValidationPolicy.minLength", actual.MinLength, desired.MinLength)
	}
	if desired.PasswordChangeInterval != actual.PasswordChangeInterval {
		diff.AddField(".settings.passwordValidationPolicy.passwordChangeInterval", actual.PasswordChangeInterval, desired.PasswordChangeInterval)
	}
	if desired.ReuseInterval != actual.ReuseInterval {
		diff.AddField(".settings.passwordValidationPolicy.reuseInterval", actual.ReuseInterval, desired.ReuseInterval)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffSqlServerAuditConfig(desired *api.SqlServerAuditConfig, actual *api.SqlServerAuditConfig) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.SqlServerAuditConfig{}
	}
	if actual == nil {
		actual = &api.SqlServerAuditConfig{}
	}
	if desired.Bucket != actual.Bucket {
		diff.AddField(".settings.sqlServerAuditConfig.bucket", actual.Bucket, desired.Bucket)
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.RetentionInterval != actual.RetentionInterval {
		diff.AddField(".settings.sqlServerAuditConfig.retentionInterval", actual.RetentionInterval, desired.RetentionInterval)
	}
	if desired.UploadInterval != actual.UploadInterval {
		diff.AddField(".settings.sqlServerAuditConfig.uploadInterval", actual.UploadInterval, desired.UploadInterval)
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return diff
}

func DiffStorageAutoResize(desired *bool, actual *bool) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	d := true
	if desired != nil {
		d = *desired
	}
	a := true
	if actual != nil {
		a = *actual
	}
	if d != a {
		diff.AddField(".settings.storageAutoResize", a, d)
	}
	return diff
}

func DiffReplicationCluster(desired *api.ReplicationCluster, actual *api.ReplicationCluster) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil {
		desired = &api.ReplicationCluster{}
	}
	if actual == nil {
		actual = &api.ReplicationCluster{}
	}
	if desired.FailoverDrReplicaName != actual.FailoverDrReplicaName {
		diff.AddField(".replicationCluster.failoverDrReplicaName", actual.FailoverDrReplicaName, desired.FailoverDrReplicaName)
	}
	// Ignore PsaWriteEndpoint. It is output only.
	// Ignore DrReplica. It is output only.
	return diff
}
