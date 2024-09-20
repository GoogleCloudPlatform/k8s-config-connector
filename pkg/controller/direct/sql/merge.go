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
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/sqladmin/v1beta4"
)

// The goal of this function is to merge the desired state with the actual API state. Some fields in the API are defaulted by GCP.
// Additionally, some fields may be desirable to "unmanage". To handle these quirks, this function implements some merging logic
// for every API field.
func MergeDesiredSQLInstanceWithActual(desired *krm.SQLInstance, refs *SQLInstanceInternalRefs, actual *api.DatabaseInstance) (*api.DatabaseInstance, bool, error) {
	merged := &api.DatabaseInstance{}
	updateRequired := false

	if desired == nil || actual == nil {
		return nil, false, fmt.Errorf("cannot merge nil SQLInstance")
	}

	if desired.Spec.ResourceID != nil {
		merged.Name = direct.ValueOf(desired.Spec.ResourceID)
	} else {
		merged.Name = desired.Name
	}
	if merged.Name != actual.Name {
		return nil, false, fmt.Errorf("cannot rename SQLInstance")
	}

	if desired.Spec.DatabaseVersion != nil {
		if direct.ValueOf(desired.Spec.DatabaseVersion) != actual.DatabaseVersion {
			// Change version
			updateRequired = true
		}
		merged.DatabaseVersion = direct.ValueOf(desired.Spec.DatabaseVersion)
	} else {
		// Keep same version
		merged.DatabaseVersion = actual.DatabaseVersion
	}

	if desired.Spec.EncryptionKMSCryptoKeyRef != nil {
		if actual.DiskEncryptionConfiguration == nil {
			// Add key
			updateRequired = true
		} else if refs.cryptoKey != actual.DiskEncryptionConfiguration.KmsKeyName {
			// Change keys
			updateRequired = true
		}
		merged.DiskEncryptionConfiguration = &api.DiskEncryptionConfiguration{
			Kind:       "sql#diskEncryptionConfiguration",
			KmsKeyName: refs.cryptoKey,
		}
	} else if actual.DiskEncryptionConfiguration != nil {
		// Remove key
		updateRequired = true
	}

	if desired.Spec.InstanceType != nil {
		if direct.ValueOf(desired.Spec.InstanceType) != actual.InstanceType {
			// Change instance type
			updateRequired = true
		}
		merged.InstanceType = direct.ValueOf(desired.Spec.InstanceType)
	} else {
		// Keep instance type
		merged.InstanceType = actual.InstanceType
	}

	if desired.Spec.MaintenanceVersion != nil {
		if direct.ValueOf(desired.Spec.MaintenanceVersion) != actual.MaintenanceVersion {
			// Change maintenance version
			updateRequired = true
		}
		merged.MaintenanceVersion = direct.ValueOf(desired.Spec.MaintenanceVersion)
	} else {
		// Keep maintenance version
		merged.MaintenanceVersion = actual.MaintenanceVersion
	}

	if desired.Spec.MasterInstanceRef != nil {
		if refs.masterInstance != actual.MasterInstanceName {
			// Change master
			updateRequired = true
		}
		merged.MasterInstanceName = refs.masterInstance
	} else if actual.MasterInstanceName != "" {
		// Remove master
		updateRequired = true
	}

	if desired.Spec.Region != nil {
		if direct.ValueOf(desired.Spec.Region) != actual.Region {
			// Change region
			updateRequired = true
		}
		merged.Region = direct.ValueOf(desired.Spec.Region)
	} else {
		// Keep region
		merged.Region = actual.Region
	}

	merged.ReplicaConfiguration = InstanceReplicaConfigurationKRMToGCP(desired.Spec.ReplicaConfiguration, refs)
	if !ReplicaConfigurationsMatch(merged.ReplicaConfiguration, actual.ReplicaConfiguration) {
		updateRequired = true
	}

	if desired.Spec.RootPassword != nil && refs.rootPassword != "" {
		if refs.rootPassword != actual.RootPassword {
			// Change root password
			updateRequired = true
		}
		merged.RootPassword = refs.rootPassword
	} else {
		// Keep root password
		merged.RootPassword = actual.RootPassword
	}

	merged.Settings = &api.Settings{
		SettingsVersion: actual.Settings.SettingsVersion,
	}

	if desired.Spec.Settings.ActivationPolicy != nil {
		if direct.ValueOf(desired.Spec.Settings.ActivationPolicy) != actual.Settings.ActivationPolicy {
			// Change activation policy
			updateRequired = true
		}
		merged.Settings.ActivationPolicy = direct.ValueOf(desired.Spec.Settings.ActivationPolicy)
	} else {
		// Keep activation policy
		merged.Settings.ActivationPolicy = actual.Settings.ActivationPolicy
	}

	if desired.Spec.Settings.ActiveDirectoryConfig != nil {
		if actual.Settings.ActiveDirectoryConfig == nil {
			// Add active directory config
			updateRequired = true
		} else if desired.Spec.Settings.ActiveDirectoryConfig.Domain != actual.Settings.ActiveDirectoryConfig.Domain {
			// Change active directory config
			updateRequired = true
		}
		merged.Settings.ActiveDirectoryConfig.Domain = desired.Spec.Settings.ActiveDirectoryConfig.Domain
	} else if actual.Settings.ActiveDirectoryConfig != nil {
		// Remove active directory config
		updateRequired = true
	}

	if desired.Spec.Settings.AdvancedMachineFeatures != nil {
		if actual.Settings.AdvancedMachineFeatures == nil {
			// Add advanced machine features
			updateRequired = true
		} else if desired.Spec.Settings.AdvancedMachineFeatures.ThreadsPerCore != &actual.Settings.AdvancedMachineFeatures.ThreadsPerCore {
			// Change advanced machine features
			updateRequired = true
		}
		merged.Settings.AdvancedMachineFeatures.ThreadsPerCore = direct.ValueOf(desired.Spec.Settings.AdvancedMachineFeatures.ThreadsPerCore)
	} else if actual.Settings.AdvancedMachineFeatures != nil {
		// Remove advanced machine features
		updateRequired = true
	}

	if desired.Spec.Settings.AuthorizedGaeApplications != nil {
		if actual.Settings.AuthorizedGaeApplications == nil {
			// Add authorized gae applications
			updateRequired = true
		} else if len(desired.Spec.Settings.AuthorizedGaeApplications) != len(actual.Settings.AuthorizedGaeApplications) {
			// todo: fix this
			// Change authorized gae applications
			updateRequired = true
		}
		merged.Settings.AuthorizedGaeApplications = desired.Spec.Settings.AuthorizedGaeApplications
	} else if len(actual.Settings.AuthorizedGaeApplications) > 0 {
		// Remove authorized gae applications
		updateRequired = true
	}

	if desired.Spec.Settings.AvailabilityType != nil {
		if direct.ValueOf(desired.Spec.Settings.AvailabilityType) != actual.Settings.AvailabilityType {
			// Change availability type
			updateRequired = true
		}
		merged.Settings.AvailabilityType = direct.ValueOf(desired.Spec.Settings.AvailabilityType)
	} else {
		// Keep availability type
		merged.Settings.AvailabilityType = actual.Settings.AvailabilityType
	}

	existingBackupConfig := actual.Settings.BackupConfiguration != nil

	if desired.Spec.Settings.BackupConfiguration != nil {
		existingRetentionSettings := existingBackupConfig && actual.Settings.BackupConfiguration.BackupRetentionSettings != nil

		if !existingBackupConfig {
			// Add backup configuration
			updateRequired = true
		}

		merged.Settings.BackupConfiguration = &api.BackupConfiguration{
			Kind: "sql#backupConfiguration",
		}

		if desired.Spec.Settings.BackupConfiguration.BackupRetentionSettings != nil {
			if !existingRetentionSettings {
				// Add retention settings
				updateRequired = true
			} else if (desired.Spec.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups != actual.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups) ||
				(direct.ValueOf(desired.Spec.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit) != actual.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit) {
				// Change retention settings
				updateRequired = true
			}
			merged.Settings.BackupConfiguration.BackupRetentionSettings = &api.BackupRetentionSettings{
				RetainedBackups: desired.Spec.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups,
				RetentionUnit:   direct.ValueOf(desired.Spec.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit),
			}
		} else if existingRetentionSettings {
			// Keep same retention settings
			merged.Settings.BackupConfiguration.BackupRetentionSettings = actual.Settings.BackupConfiguration.BackupRetentionSettings
		}

		if desired.Spec.Settings.BackupConfiguration.BinaryLogEnabled != nil {
			if existingBackupConfig {
				if direct.ValueOf(desired.Spec.Settings.BackupConfiguration.BinaryLogEnabled) != actual.Settings.BackupConfiguration.BinaryLogEnabled {
					// Change binary log enabled
					updateRequired = true
				}
			}
			merged.Settings.BackupConfiguration.BinaryLogEnabled = direct.ValueOf(desired.Spec.Settings.BackupConfiguration.BinaryLogEnabled)
		} else if existingBackupConfig {
			// Keep same binary log enabled
			merged.Settings.BackupConfiguration.BinaryLogEnabled = actual.Settings.BackupConfiguration.BinaryLogEnabled
		}

		if desired.Spec.Settings.BackupConfiguration.Enabled != nil {
			if existingBackupConfig {
				if direct.ValueOf(desired.Spec.Settings.BackupConfiguration.Enabled) != actual.Settings.BackupConfiguration.Enabled {
					// Change enabled
					updateRequired = true
				}
			}
			merged.Settings.BackupConfiguration.Enabled = direct.ValueOf(desired.Spec.Settings.BackupConfiguration.Enabled)
		} else if existingBackupConfig {
			// Keep same enabled
			merged.Settings.BackupConfiguration.Enabled = actual.Settings.BackupConfiguration.Enabled
		}

		if desired.Spec.Settings.BackupConfiguration.Location != nil {
			if existingBackupConfig {
				if direct.ValueOf(desired.Spec.Settings.BackupConfiguration.Location) != actual.Settings.BackupConfiguration.Location {
					// Change location
					updateRequired = true
				}
			}
			merged.Settings.BackupConfiguration.Location = direct.ValueOf(desired.Spec.Settings.BackupConfiguration.Location)
		} else if existingBackupConfig {
			// Keep same location
			merged.Settings.BackupConfiguration.Location = actual.Settings.BackupConfiguration.Location
		}

		if desired.Spec.Settings.BackupConfiguration.PointInTimeRecoveryEnabled != nil {
			if existingBackupConfig {
				if direct.ValueOf(desired.Spec.Settings.BackupConfiguration.PointInTimeRecoveryEnabled) != actual.Settings.BackupConfiguration.PointInTimeRecoveryEnabled {
					// Change point in time recovery enabled
					updateRequired = true
				}
			}
			merged.Settings.BackupConfiguration.PointInTimeRecoveryEnabled = direct.ValueOf(desired.Spec.Settings.BackupConfiguration.PointInTimeRecoveryEnabled)
		} else if existingBackupConfig {
			// Keep same point in time recovery enabled
			merged.Settings.BackupConfiguration.PointInTimeRecoveryEnabled = actual.Settings.BackupConfiguration.PointInTimeRecoveryEnabled
		}

		if desired.Spec.Settings.BackupConfiguration.StartTime != nil {
			if existingBackupConfig {
				if direct.ValueOf(desired.Spec.Settings.BackupConfiguration.StartTime) != actual.Settings.BackupConfiguration.StartTime {
					// Change start time
					updateRequired = true
				}
			}
			merged.Settings.BackupConfiguration.StartTime = direct.ValueOf(desired.Spec.Settings.BackupConfiguration.StartTime)
		} else if existingBackupConfig {
			// Keep same start time
			merged.Settings.BackupConfiguration.StartTime = actual.Settings.BackupConfiguration.StartTime
		}

		if desired.Spec.Settings.BackupConfiguration.TransactionLogRetentionDays != nil {
			if existingBackupConfig {
				if direct.ValueOf(desired.Spec.Settings.BackupConfiguration.TransactionLogRetentionDays) != actual.Settings.BackupConfiguration.TransactionLogRetentionDays {
					// Change transaction log retention days
					updateRequired = true
				}
			}
			merged.Settings.BackupConfiguration.TransactionLogRetentionDays = direct.ValueOf(desired.Spec.Settings.BackupConfiguration.TransactionLogRetentionDays)
		} else if existingBackupConfig {
			// Keep same transaction log retention days
			merged.Settings.BackupConfiguration.TransactionLogRetentionDays = actual.Settings.BackupConfiguration.TransactionLogRetentionDays
		}
	} else if existingBackupConfig {
		// Keep same backup configuration
		merged.Settings.BackupConfiguration = actual.Settings.BackupConfiguration
	}
	if merged.Settings.BackupConfiguration != nil {
		merged.Settings.BackupConfiguration.ForceSendFields = []string{
			"BinaryLogEnabled",
			"Enabled",
			"PointInTimeRecoveryEnabled",
		}
	}

	merged.Settings.Collation = direct.ValueOf(desired.Spec.Settings.Collation)
	if merged.Settings.Collation != actual.Settings.Collation {
		updateRequired = true
	}

	if desired.Spec.Settings.ConnectorEnforcement != nil {
		if direct.ValueOf(desired.Spec.Settings.ConnectorEnforcement) != actual.Settings.ConnectorEnforcement {
			// Change connector enforcement
			updateRequired = true
		}
		merged.Settings.ConnectorEnforcement = direct.ValueOf(desired.Spec.Settings.ConnectorEnforcement)
	} else {
		// Keep connector enforcement
		merged.Settings.ConnectorEnforcement = actual.Settings.ConnectorEnforcement
	}

	if desired.Spec.Settings.CrashSafeReplication != nil {
		if direct.ValueOf(desired.Spec.Settings.CrashSafeReplication) != actual.Settings.CrashSafeReplicationEnabled {
			// Change crash safe replication
			updateRequired = true
		}
		merged.Settings.CrashSafeReplicationEnabled = direct.ValueOf(desired.Spec.Settings.CrashSafeReplication)
	} else {
		// Keep crash safe replication
		merged.Settings.CrashSafeReplicationEnabled = actual.Settings.CrashSafeReplicationEnabled
	}

	merged.Settings.DataCacheConfig = InstanceDataCacheConfigKRMToGCP(desired.Spec.Settings.DataCacheConfig)
	if !DataCacheConfigsMatch(merged.Settings.DataCacheConfig, actual.Settings.DataCacheConfig) {
		updateRequired = true
	}

	if desired.Spec.Settings.DatabaseFlags != nil {
		if actual.Settings.DatabaseFlags == nil {
			// Add database flags
			updateRequired = true
		} else if len(desired.Spec.Settings.DatabaseFlags) != len(actual.Settings.DatabaseFlags) {
			// todo: fix this
			// Change database flags
			updateRequired = true
		}
		merged.Settings.DatabaseFlags = []*api.DatabaseFlags{}
		for _, flag := range desired.Spec.Settings.DatabaseFlags {
			merged.Settings.DatabaseFlags = append(merged.Settings.DatabaseFlags, &api.DatabaseFlags{
				Name:  flag.Name,
				Value: flag.Value,
			})
		}
	} else if actual.Settings.DatabaseFlags != nil {
		// Remove database flags
		updateRequired = true
	}

	if desired.Spec.Settings.DeletionProtectionEnabled != nil {
		if direct.ValueOf(desired.Spec.Settings.DeletionProtectionEnabled) != actual.Settings.DeletionProtectionEnabled {
			// Change deletion protection enabled
			updateRequired = true
		}
		merged.Settings.DeletionProtectionEnabled = direct.ValueOf(desired.Spec.Settings.DeletionProtectionEnabled)
	} else {
		// Keep deletion protection enabled
		merged.Settings.DeletionProtectionEnabled = actual.Settings.DeletionProtectionEnabled
	}

	merged.Settings.DenyMaintenancePeriods = InstanceDenyMaintenancePeriodsKRMToGCP(desired.Spec.Settings.DenyMaintenancePeriod)
	if !DenyMaintenancePeriodListsMatch(merged.Settings.DenyMaintenancePeriods, actual.Settings.DenyMaintenancePeriods) {
		updateRequired = true
	}

	merged.Settings.StorageAutoResize = desired.Spec.Settings.DiskAutoresize
	if merged.Settings.StorageAutoResize != actual.Settings.StorageAutoResize {
		updateRequired = true
	}

	merged.Settings.StorageAutoResizeLimit = direct.ValueOf(desired.Spec.Settings.DiskAutoresizeLimit)
	if merged.Settings.StorageAutoResizeLimit != actual.Settings.StorageAutoResizeLimit {
		updateRequired = true
	}

	merged.Settings.DataDiskSizeGb = direct.ValueOf(desired.Spec.Settings.DiskSize)
	if merged.Settings.DataDiskSizeGb != actual.Settings.DataDiskSizeGb {
		updateRequired = true
	}

	merged.Settings.DataDiskType = direct.ValueOf(desired.Spec.Settings.DiskType)
	if merged.Settings.DataDiskType != actual.Settings.DataDiskType {
		updateRequired = true
	}

	merged.Settings.Edition = direct.ValueOf(desired.Spec.Settings.Edition)
	if merged.Settings.Edition != actual.Settings.Edition {
		updateRequired = true
	}

	merged.Settings.InsightsConfig = InstanceInsightsConfigKRMToGCP(desired.Spec.Settings.InsightsConfig)
	if !InsightsConfigsMatch(merged.Settings.InsightsConfig, actual.Settings.InsightsConfig) {
		updateRequired = true
	}

	merged.Settings.IpConfiguration = InstanceIpConfigurationKRMToGCP(desired.Spec.Settings.IpConfiguration, refs)
	if !IpConfigurationsMatch(merged.Settings.IpConfiguration, actual.Settings.IpConfiguration) {
		updateRequired = true
	}

	merged.Settings.LocationPreference = InstanceLocationPreferenceKRMToGCP(desired.Spec.Settings.LocationPreference)
	if !LocationPreferencesMatch(merged.Settings.LocationPreference, actual.Settings.LocationPreference) {
		updateRequired = true
	}

	merged.Settings.MaintenanceWindow = InstanceMaintenanceWindowKRMToGCP(desired.Spec.Settings.MaintenanceWindow)
	if !MaintenanceWindowsMatch(merged.Settings.MaintenanceWindow, actual.Settings.MaintenanceWindow) {
		updateRequired = true
	}

	merged.Settings.PasswordValidationPolicy = InstancePasswordValidationPolicyKRMToGCP(desired.Spec.Settings.PasswordValidationPolicy)
	if !PasswordValidationPoliciesMatch(merged.Settings.PasswordValidationPolicy, actual.Settings.PasswordValidationPolicy) {
		updateRequired = true
	}

	merged.Settings.PricingPlan = direct.ValueOf(desired.Spec.Settings.PricingPlan)
	if merged.Settings.PricingPlan != actual.Settings.PricingPlan {
		updateRequired = true
	}

	merged.Settings.ReplicationType = direct.ValueOf(desired.Spec.Settings.ReplicationType)
	if merged.Settings.ReplicationType != actual.Settings.ReplicationType {
		updateRequired = true
	}

	merged.Settings.SqlServerAuditConfig = InstanceSqlServerAuditConfigKRMToGCP(desired.Spec.Settings.SqlServerAuditConfig, refs)
	if !SqlServerAuditConfigsMatch(merged.Settings.SqlServerAuditConfig, actual.Settings.SqlServerAuditConfig) {
		updateRequired = true
	}

	merged.Settings.Tier = desired.Spec.Settings.Tier
	if merged.Settings.Tier != actual.Settings.Tier {
		updateRequired = true
	}

	merged.Settings.TimeZone = direct.ValueOf(desired.Spec.Settings.TimeZone)
	if merged.Settings.TimeZone != actual.Settings.TimeZone {
		updateRequired = true
	}

	if !reflect.DeepEqual(desired.Labels, actual.Settings.UserLabels) {
		updateRequired = true
	}
	merged.Settings.UserLabels = desired.Labels

	// todo: Remove this after switching over to use InstanceSettingsKRMToGCP
	if desired.Spec.Settings.DiskAutoresize != nil {
		merged.Settings.ForceSendFields = append(merged.ForceSendFields, "StorageAutoResize")
	}

	return merged, updateRequired, nil
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

func AclEntryListsMatch(desired []*api.AclEntry, actual []*api.AclEntry) bool {
	if len(desired) != len(actual) {
		return false
	}
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

func PointersMatch[T any](desired *T, actual *T) bool {
	if (desired == nil && actual != nil) || (desired != nil && actual == nil) {
		// Pointers are not matching if one is nil and the other is not nil.
		return false
	}
	// Otherwise, they match. Either both are nil, or both are not nil.
	return true
}
