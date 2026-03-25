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
	if desired == nil && actual == nil {
		return diff
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
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".diskEncryptionConfiguration", actual, desired)
		return diff
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
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".replicaConfiguration", actual, desired)
		return diff
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
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings", actual, desired)
		return diff
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
	if !DatabaseFlagListsMatch(desired.DatabaseFlags, actual.DatabaseFlags) {
		diff.AddField(".settings.databaseFlags", actual.DatabaseFlags, desired.DatabaseFlags)
	}
	// Ignore DatabaseReplicationEnabled. It is not supported in KRM API.
	if desired.DeletionProtectionEnabled != actual.DeletionProtectionEnabled {
		diff.AddField(".settings.deletionProtectionEnabled", actual.DeletionProtectionEnabled, desired.DeletionProtectionEnabled)
	}
	if !DenyMaintenancePeriodListsMatch(desired.DenyMaintenancePeriods, actual.DenyMaintenancePeriods) {
		diff.AddField(".settings.denyMaintenancePeriods", actual.DenyMaintenancePeriods, desired.DenyMaintenancePeriods)
	}
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
	if !reflect.DeepEqual(desired.UserLabels, actual.UserLabels) {
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

func DiffMysqlReplicaConfiguration(desired *api.MySqlReplicaConfiguration, actual *api.MySqlReplicaConfiguration) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".replicaConfiguration.mysqlReplicaConfiguration", actual, desired)
		return diff
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
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings.activeDirectoryConfig", actual, desired)
		return diff
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
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings.advancedMachineFeatures", actual, desired)
		return diff
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
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings.backupConfiguration", actual, desired)
		return diff
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
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings.backupConfiguration.backupRetentionSettings", actual, desired)
		return diff
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

func DatabaseFlagListsMatch(desired []*api.DatabaseFlags, actual []*api.DatabaseFlags) bool {
	if len(desired) != len(actual) {
		return false
	}
	for i := 0; i < len(desired); i++ {
		if !DatabaseFlagsMatch(desired[i], actual[i]) {
			return false
		}
	}
	return true
}

func DatabaseFlagsMatch(desired *api.DatabaseFlags, actual *api.DatabaseFlags) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.Name != actual.Name {
		return false
	}
	if desired.Value != actual.Value {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func DenyMaintenancePeriodListsMatch(desired []*api.DenyMaintenancePeriod, actual []*api.DenyMaintenancePeriod) bool {
	if len(desired) != len(actual) {
		return false
	}
	for i := 0; i < len(desired); i++ {
		if !DenyMaintenancePeriodsMatch(desired[i], actual[i]) {
			return false
		}
	}
	return true
}

func DenyMaintenancePeriodsMatch(desired *api.DenyMaintenancePeriod, actual *api.DenyMaintenancePeriod) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.EndDate != actual.EndDate {
		return false
	}
	if desired.StartDate != actual.StartDate {
		return false
	}
	if desired.Time != actual.Time {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func DiffInsightsConfig(desired *api.InsightsConfig, actual *api.InsightsConfig) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings.insightsConfig", actual, desired)
		return diff
	}
	if desired.QueryInsightsEnabled != actual.QueryInsightsEnabled {
		diff.AddField(".settings.insightsConfig.queryInsightsEnabled", actual.QueryInsightsEnabled, desired.QueryInsightsEnabled)
	}
	if desired.QueryPlansPerMinute != actual.QueryPlansPerMinute {
		diff.AddField(".settings.insightsConfig.queryPlansPerMinute", actual.QueryPlansPerMinute, desired.QueryPlansPerMinute)
	}
	if desired.QueryStringLength != actual.QueryStringLength {
		diff.AddField(".settings.insightsConfig.queryStringLength", actual.QueryStringLength, desired.QueryStringLength)
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
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings.ipConfiguration", actual, desired)
		return diff
	}
	if desired.AllocatedIpRange != actual.AllocatedIpRange {
		diff.AddField(".settings.ipConfiguration.allocatedIpRange", actual.AllocatedIpRange, desired.AllocatedIpRange)
	}
	if !AclEntryListsMatch(desired.AuthorizedNetworks, actual.AuthorizedNetworks) {
		diff.AddField(".settings.ipConfiguration.authorizedNetworks", actual.AuthorizedNetworks, desired.AuthorizedNetworks)
	}
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

func (a AclEntriesByName) Len() int           { return len(a) }
func (a AclEntriesByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AclEntriesByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func AclEntryListsMatch(desired []*api.AclEntry, actual []*api.AclEntry) bool {
	if len(desired) != len(actual) {
		return false
	}
	// We mustiterate over the AclEntry lists in sorted order,
	// so that the comparison is deterministic.
	sort.Sort(AclEntriesByName(desired))
	sort.Sort(AclEntriesByName(actual))
	// Compare the AclEntry lists.
	for i := 0; i < len(desired); i++ {
		if !AclEntriesMatch(desired[i], actual[i]) {
			return false
		}
	}
	return true
}

func AclEntriesMatch(desired *api.AclEntry, actual *api.AclEntry) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.ExpirationTime != actual.ExpirationTime {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.Name != actual.Name {
		return false
	}
	if desired.Value != actual.Value {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func DiffPscConfig(desired *api.PscConfig, actual *api.PscConfig) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings.ipConfiguration.pscConfig", actual, desired)
		return diff
	}
	if !reflect.DeepEqual(desired.AllowedConsumerProjects, actual.AllowedConsumerProjects) {
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
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings.locationPreference", actual, desired)
		return diff
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
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings.maintenanceWindow", actual, desired)
		return diff
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
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings.passwordValidationPolicy", actual, desired)
		return diff
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
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings.sqlServerAuditConfig", actual, desired)
		return diff
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
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings.storageAutoResize", actual, desired)
		return diff
	}
	if *desired != *actual {
		diff.AddField(".settings.storageAutoResize", *actual, *desired)
	}
	return diff
}

func PointersMatch[T any](desired *T, actual *T) bool {
	if (desired == nil && actual != nil) || (desired != nil && actual == nil) {
		// Pointers are not matching if one is nil and the other is not nil.
		return false
	}
	// Otherwise, they match. Either both are nil, or both are not nil.
	return true
}

func DiffReplicationCluster(desired *api.ReplicationCluster, actual *api.ReplicationCluster) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	if desired == nil && actual == nil {
		return diff
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".replicationCluster", actual, desired)
		return diff
	}
	if desired.FailoverDrReplicaName != actual.FailoverDrReplicaName {
		diff.AddField(".replicationCluster.failoverDrReplicaName", actual.FailoverDrReplicaName, desired.FailoverDrReplicaName)
	}
	// Ignore PsaWriteEndpoint. It is output only.
	// Ignore DrReplica. It is output only.
	return diff
}
