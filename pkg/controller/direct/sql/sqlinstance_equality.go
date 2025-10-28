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

func InstancesMatch(desired *api.DatabaseInstance, actual *api.DatabaseInstance, diff *structuredreporting.Diff) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.DatabaseVersion != actual.DatabaseVersion {
		diff.AddField(".databaseVersion", actual.DatabaseVersion, desired.DatabaseVersion)
		return false
	}
	if !DiskEncryptionConfigurationsMatch(desired.DiskEncryptionConfiguration, actual.DiskEncryptionConfiguration) {
		diff.AddField(".diskEncryptionConfiguration", actual.DiskEncryptionConfiguration, desired.DiskEncryptionConfiguration)
		return false
	}
	// Ignore GeminiConfig. It is not supported in KRM API.
	if desired.InstanceType != actual.InstanceType {
		diff.AddField(".instanceType", actual.InstanceType, desired.InstanceType)
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.MaintenanceVersion != actual.MaintenanceVersion {
		diff.AddField(".maintenanceVersion", actual.MaintenanceVersion, desired.MaintenanceVersion)
		return false
	}
	if desired.MasterInstanceName != actual.MasterInstanceName {
		diff.AddField(".masterInstanceName", actual.MasterInstanceName, desired.MasterInstanceName)
		return false
	}
	// Ignore MaxDiskSize. It is not supported in KRM API.
	if desired.Name != actual.Name {
		diff.AddField(".name", actual.Name, desired.Name)
		return false
	}
	// Ignore OnPremisesConfiguration. It is not supported in KRM API.
	if desired.Region != actual.Region {
		diff.AddField(".region", actual.Region, desired.Region)
		return false
	}
	if !ReplicaConfigurationsMatch(desired.ReplicaConfiguration, actual.ReplicaConfiguration) {
		diff.AddField(".replicaConfiguration", actual.ReplicaConfiguration, desired.ReplicaConfiguration)
		return false
	}
	// Ignore ReplicationCluster. It is not supported in KRM API.
	// Ignore RootPassword. It is not exported.
	if !SettingsMatch(desired.Settings, actual.Settings, diff) {
		return false
	}
	// Ignore SqlNetworkArchitecture. It is not supported in KRM API.
	// Ignore SwitchTransactionLogsToCloudStorageEnabled. It is not supported in KRM API.
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func DiskEncryptionConfigurationsMatch(desired *api.DiskEncryptionConfiguration, actual *api.DiskEncryptionConfiguration) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.KmsKeyName != actual.KmsKeyName {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func ReplicaConfigurationsMatch(desired *api.ReplicaConfiguration, actual *api.ReplicaConfiguration) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	// Ignore CascadableReplica. It is not supported in KRM API.
	if desired.FailoverTarget != actual.FailoverTarget {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if !MysqlReplicaConfigurationsMatch(desired.MysqlReplicaConfiguration, actual.MysqlReplicaConfiguration) {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func SettingsMatch(desired *api.Settings, actual *api.Settings, diff *structuredreporting.Diff) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings", actual, desired)
		return false
	}
	if desired.ActivationPolicy != actual.ActivationPolicy {
		diff.AddField(".settings.activationPolicy", actual.ActivationPolicy, desired.ActivationPolicy)
		return false
	}
	if !ActiveDirectoryConfigsMatch(desired.ActiveDirectoryConfig, actual.ActiveDirectoryConfig) {
		diff.AddField(".settings.activeDirectoryConfig", actual.ActiveDirectoryConfig, desired.ActiveDirectoryConfig)
		return false
	}
	if !AdvancedMachineFeaturesMatch(desired.AdvancedMachineFeatures, actual.AdvancedMachineFeatures) {
		diff.AddField(".settings.advancedMachineFeatures", actual.AdvancedMachineFeatures, desired.AdvancedMachineFeatures)
		return false
	}
	if !slicesMatch(desired.AuthorizedGaeApplications, actual.AuthorizedGaeApplications) {
		diff.AddField(".settings.authorizedGaeApplications", actual.AuthorizedGaeApplications, desired.AuthorizedGaeApplications)
		return false
	}
	if desired.AvailabilityType != actual.AvailabilityType {
		diff.AddField(".settings.availabilityType", actual.AvailabilityType, desired.AvailabilityType)
		return false
	}
	if !BackupConfigurationsMatch(desired.BackupConfiguration, actual.BackupConfiguration) {
		diff.AddField(".settings.backupConfiguration", actual.BackupConfiguration, desired.BackupConfiguration)
		return false
	}
	if desired.Collation != actual.Collation {
		diff.AddField(".settings.collation", actual.Collation, desired.Collation)
		return false
	}
	if desired.ConnectorEnforcement != actual.ConnectorEnforcement {
		diff.AddField(".settings.connectorEnforcement", actual.ConnectorEnforcement, desired.ConnectorEnforcement)
		return false
	}
	// Ignore CrashSafeReplicationEnabled. It is only applicable to first-gen instances.
	if !DataCacheConfigsMatch(desired.DataCacheConfig, actual.DataCacheConfig) {
		diff.AddField(".settings.dataCacheConfig", actual.DataCacheConfig, desired.DataCacheConfig)
		return false
	}
	if desired.DataDiskSizeGb != actual.DataDiskSizeGb {
		diff.AddField(".settings.dataDiskSizeGb", actual.DataDiskSizeGb, desired.DataDiskSizeGb)
		return false
	}
	if desired.DataDiskType != actual.DataDiskType {
		diff.AddField(".settings.dataDiskType", actual.DataDiskType, desired.DataDiskType)
		return false
	}
	if !DatabaseFlagListsMatch(desired.DatabaseFlags, actual.DatabaseFlags) {
		diff.AddField(".settings.databaseFlags", actual.DatabaseFlags, desired.DatabaseFlags)
		return false
	}
	// Ignore DatabaseReplicationEnabled. It is not supported in KRM API.
	if desired.DeletionProtectionEnabled != actual.DeletionProtectionEnabled {
		diff.AddField(".settings.deletionProtectionEnabled", actual.DeletionProtectionEnabled, desired.DeletionProtectionEnabled)
		return false
	}
	if !DenyMaintenancePeriodListsMatch(desired.DenyMaintenancePeriods, actual.DenyMaintenancePeriods) {
		diff.AddField(".settings.denyMaintenancePeriods", actual.DenyMaintenancePeriods, desired.DenyMaintenancePeriods)
		return false
	}
	if desired.Edition != actual.Edition {
		diff.AddField(".settings.edition", actual.Edition, desired.Edition)
		return false
	}
	// Ignore EnableDataplexIntegration. It is not supported in KRM API.
	// Ignore EnableGoogleMlIntegration. It is not supported in KRM API.
	if !InsightsConfigsMatch(desired.InsightsConfig, actual.InsightsConfig) {
		diff.AddField(".settings.insightsConfig", actual.InsightsConfig, desired.InsightsConfig)
		return false
	}
	if !IpConfigurationsMatch(desired.IpConfiguration, actual.IpConfiguration) {
		diff.AddField(".settings.ipConfiguration", actual.IpConfiguration, desired.IpConfiguration)
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if !LocationPreferencesMatch(desired.LocationPreference, actual.LocationPreference) {
		diff.AddField(".settings.locationPreference", actual.LocationPreference, desired.LocationPreference)
		return false
	}
	if !MaintenanceWindowsMatch(desired.MaintenanceWindow, actual.MaintenanceWindow) {
		diff.AddField(".settings.maintenanceWindow", actual.MaintenanceWindow, desired.MaintenanceWindow)
		return false
	}
	if !PasswordValidationPoliciesMatch(desired.PasswordValidationPolicy, actual.PasswordValidationPolicy) {
		diff.AddField(".settings.passwordValidationPolicy", actual.PasswordValidationPolicy, desired.PasswordValidationPolicy)
		return false
	}
	if desired.PricingPlan != actual.PricingPlan {
		diff.AddField(".settings.pricingPlan", actual.PricingPlan, desired.PricingPlan)
		return false
	}
	if desired.ReplicationType != actual.ReplicationType {
		diff.AddField(".settings.replicationType", actual.ReplicationType, desired.ReplicationType)
		return false
	}
	if desired.SettingsVersion != actual.SettingsVersion {
		diff.AddField(".settings.settingsVersion", actual.SettingsVersion, desired.SettingsVersion)
		return false
	}
	if !SqlServerAuditConfigsMatch(desired.SqlServerAuditConfig, actual.SqlServerAuditConfig) {
		diff.AddField(".settings.sqlServerAuditConfig", actual.SqlServerAuditConfig, desired.SqlServerAuditConfig)
		return false
	}
	if !StorageAutoResizesMatch(desired.StorageAutoResize, actual.StorageAutoResize) {
		diff.AddField(".settings.storageAutoResize", actual.StorageAutoResize, desired.StorageAutoResize)
		return false
	}
	if desired.StorageAutoResizeLimit != actual.StorageAutoResizeLimit {
		diff.AddField(".settings.storageAutoResizeLimit", actual.StorageAutoResizeLimit, desired.StorageAutoResizeLimit)
		return false
	}
	if desired.Tier != actual.Tier {
		diff.AddField(".settings.tier", actual.Tier, desired.Tier)
		return false
	}
	if desired.TimeZone != actual.TimeZone {
		diff.AddField(".settings.timeZone", actual.TimeZone, desired.TimeZone)
		return false
	}
	if !reflect.DeepEqual(desired.UserLabels, actual.UserLabels) {
		diff.AddField(".settings.userLabels", actual.UserLabels, desired.UserLabels)
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
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

func MysqlReplicaConfigurationsMatch(desired *api.MySqlReplicaConfiguration, actual *api.MySqlReplicaConfiguration) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.CaCertificate != actual.CaCertificate {
		return false
	}
	if desired.ClientCertificate != actual.ClientCertificate {
		return false
	}
	if desired.ClientKey != actual.ClientKey {
		return false
	}
	if desired.ConnectRetryInterval != actual.ConnectRetryInterval {
		return false
	}
	if desired.DumpFilePath != actual.DumpFilePath {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.MasterHeartbeatPeriod != actual.MasterHeartbeatPeriod {
		return false
	}
	// Ignore Password. It is not exported.
	if desired.SslCipher != actual.SslCipher {
		return false
	}
	if desired.Username != actual.Username {
		return false
	}
	if desired.VerifyServerCertificate != actual.VerifyServerCertificate {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func ActiveDirectoryConfigsMatch(desired *api.SqlActiveDirectoryConfig, actual *api.SqlActiveDirectoryConfig) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.Domain != actual.Domain {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func AdvancedMachineFeaturesMatch(desired *api.AdvancedMachineFeatures, actual *api.AdvancedMachineFeatures) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.ThreadsPerCore != actual.ThreadsPerCore {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func BackupConfigurationsMatch(desired *api.BackupConfiguration, actual *api.BackupConfiguration) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if !BackupRetentionSettingsMatch(desired.BackupRetentionSettings, actual.BackupRetentionSettings) {
		return false
	}
	if desired.BinaryLogEnabled != actual.BinaryLogEnabled {
		return false
	}
	if desired.Enabled != actual.Enabled {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.Location != actual.Location {
		return false
	}
	if desired.PointInTimeRecoveryEnabled != actual.PointInTimeRecoveryEnabled {
		return false
	}
	// Ignore StartTime if it is not set. empty string is not a valid start time.
	if desired.StartTime != "" && desired.StartTime != actual.StartTime {
		return false
	}
	// Ignore TransactionLogRetentionDays if it is not set. 0 is not a valid transaction log retention days.
	if desired.TransactionLogRetentionDays != 0 && desired.TransactionLogRetentionDays != actual.TransactionLogRetentionDays {
		return false
	}

	// Ignore ReplicationLogArchivingEnabled. It is not supported in KRM API.
	// Ignore TransactionalLogStorageState. It is not supported in KRM API.
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}
func BackupRetentionSettingsMatch(desired *api.BackupRetentionSettings, actual *api.BackupRetentionSettings) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.RetainedBackups != actual.RetainedBackups {
		return false
	}
	if desired.RetentionUnit != actual.RetentionUnit {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func DataCacheConfigsMatch(desired *api.DataCacheConfig, actual *api.DataCacheConfig) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.DataCacheEnabled != actual.DataCacheEnabled {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
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

func InsightsConfigsMatch(desired *api.InsightsConfig, actual *api.InsightsConfig) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.QueryInsightsEnabled != actual.QueryInsightsEnabled {
		return false
	}
	if desired.QueryPlansPerMinute != actual.QueryPlansPerMinute {
		return false
	}
	if desired.QueryStringLength != actual.QueryStringLength {
		return false
	}
	if desired.RecordApplicationTags != actual.RecordApplicationTags {
		return false
	}
	if desired.RecordClientAddress != actual.RecordClientAddress {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func IpConfigurationsMatch(desired *api.IpConfiguration, actual *api.IpConfiguration) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.AllocatedIpRange != actual.AllocatedIpRange {
		return false
	}
	if !AclEntryListsMatch(desired.AuthorizedNetworks, actual.AuthorizedNetworks) {
		return false
	}
	if desired.EnablePrivatePathForGoogleCloudServices != actual.EnablePrivatePathForGoogleCloudServices {
		return false
	}
	if desired.Ipv4Enabled != actual.Ipv4Enabled {
		return false
	}
	if desired.PrivateNetwork != actual.PrivateNetwork {
		return false
	}
	if !PscConfigsMatch(desired.PscConfig, actual.PscConfig) {
		return false
	}
	if desired.RequireSsl != actual.RequireSsl {
		return false
	}
	// Ignore ServerCaMode. It is not supported in KRM API.
	if desired.SslMode != actual.SslMode {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
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

func PscConfigsMatch(desired *api.PscConfig, actual *api.PscConfig) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if !reflect.DeepEqual(desired.AllowedConsumerProjects, actual.AllowedConsumerProjects) {
		return false
	}
	if desired.PscEnabled != actual.PscEnabled {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func LocationPreferencesMatch(desired *api.LocationPreference, actual *api.LocationPreference) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.FollowGaeApplication != actual.FollowGaeApplication {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.SecondaryZone != actual.SecondaryZone {
		return false
	}
	if desired.Zone != actual.Zone {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func MaintenanceWindowsMatch(desired *api.MaintenanceWindow, actual *api.MaintenanceWindow) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.Day != actual.Day {
		return false
	}
	if desired.Hour != actual.Hour {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.UpdateTrack != actual.UpdateTrack {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func PasswordValidationPoliciesMatch(desired *api.PasswordValidationPolicy, actual *api.PasswordValidationPolicy) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.Complexity != actual.Complexity {
		return false
	}
	// Ignore DisallowCompromisedCredentials. It is not supported in KRM API.
	if desired.DisallowUsernameSubstring != actual.DisallowUsernameSubstring {
		return false
	}
	if desired.EnablePasswordPolicy != actual.EnablePasswordPolicy {
		return false
	}
	if desired.MinLength != actual.MinLength {
		return false
	}
	if desired.PasswordChangeInterval != actual.PasswordChangeInterval {
		return false
	}
	if desired.ReuseInterval != actual.ReuseInterval {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func SqlServerAuditConfigsMatch(desired *api.SqlServerAuditConfig, actual *api.SqlServerAuditConfig) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if desired.Bucket != actual.Bucket {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if desired.RetentionInterval != actual.RetentionInterval {
		return false
	}
	if desired.UploadInterval != actual.UploadInterval {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func StorageAutoResizesMatch(desired *bool, actual *bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if *desired != *actual {
		return false
	}
	return true
}

func PointersMatch[T any](desired *T, actual *T) bool {
	if (desired == nil && actual != nil) || (desired != nil && actual == nil) {
		// Pointers are not matching if one is nil and the other is not nil.
		return false
	}
	// Otherwise, they match. Either both are nil, or both are not nil.
	return true
}
