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

	existingReplicaConfig := actual.ReplicaConfiguration != nil
	if desired.Spec.ReplicaConfiguration != nil {
		existingMysqlReplicaConfig := existingReplicaConfig && actual.ReplicaConfiguration.MysqlReplicaConfiguration != nil

		if !existingReplicaConfig {
			// Add replica configuration
			updateRequired = true
		}

		merged.ReplicaConfiguration = &api.ReplicaConfiguration{
			Kind: "sql#replicaConfiguration",
		}

		if desired.Spec.ReplicaConfiguration.FailoverTarget != nil {
			if existingReplicaConfig {
				if direct.ValueOf(desired.Spec.ReplicaConfiguration.FailoverTarget) != actual.ReplicaConfiguration.FailoverTarget {
					// Change failover target
					updateRequired = true
				}
			} else {
				// Add failover target
				updateRequired = true
			}
			merged.ReplicaConfiguration.FailoverTarget = direct.ValueOf(desired.Spec.ReplicaConfiguration.FailoverTarget)
		} else if existingReplicaConfig {
			// Remove failover target
			updateRequired = true
		}

		if desired.Spec.ReplicaConfiguration.CaCertificate != nil ||
			desired.Spec.ReplicaConfiguration.ClientCertificate != nil ||
			desired.Spec.ReplicaConfiguration.ClientKey != nil ||
			desired.Spec.ReplicaConfiguration.ConnectRetryInterval != nil ||
			desired.Spec.ReplicaConfiguration.DumpFilePath != nil ||
			desired.Spec.ReplicaConfiguration.MasterHeartbeatPeriod != nil ||
			desired.Spec.ReplicaConfiguration.Password != nil ||
			desired.Spec.ReplicaConfiguration.SslCipher != nil ||
			desired.Spec.ReplicaConfiguration.Username != nil ||
			desired.Spec.ReplicaConfiguration.VerifyServerCertificate != nil {
			merged.ReplicaConfiguration.MysqlReplicaConfiguration = &api.MySqlReplicaConfiguration{}
		}

		if desired.Spec.ReplicaConfiguration.CaCertificate != nil {
			if existingMysqlReplicaConfig {
				if direct.ValueOf(desired.Spec.ReplicaConfiguration.CaCertificate) != actual.ReplicaConfiguration.MysqlReplicaConfiguration.CaCertificate {
					// Change CA certificate
					updateRequired = true
				}
			} else {
				// Add CA certificate
				updateRequired = true
			}
			merged.ReplicaConfiguration.MysqlReplicaConfiguration.CaCertificate = direct.ValueOf(desired.Spec.ReplicaConfiguration.CaCertificate)
		} else if existingMysqlReplicaConfig && actual.ReplicaConfiguration.MysqlReplicaConfiguration.CaCertificate != "" {
			// Remove CA certificate
			updateRequired = true
		}

		if desired.Spec.ReplicaConfiguration.ClientCertificate != nil {
			if existingMysqlReplicaConfig {
				if direct.ValueOf(desired.Spec.ReplicaConfiguration.ClientCertificate) != actual.ReplicaConfiguration.MysqlReplicaConfiguration.ClientCertificate {
					// Change client certificate
					updateRequired = true
				}
			} else {
				// Add client certificate
				updateRequired = true
			}
			merged.ReplicaConfiguration.MysqlReplicaConfiguration.ClientCertificate = direct.ValueOf(desired.Spec.ReplicaConfiguration.ClientCertificate)
		} else if existingMysqlReplicaConfig && actual.ReplicaConfiguration.MysqlReplicaConfiguration.ClientCertificate != "" {
			// Remove client certificate
			updateRequired = true
		}

		if desired.Spec.ReplicaConfiguration.ClientKey != nil {
			if existingMysqlReplicaConfig {
				if direct.ValueOf(desired.Spec.ReplicaConfiguration.ClientKey) != actual.ReplicaConfiguration.MysqlReplicaConfiguration.ClientKey {
					// Change client key
					updateRequired = true
				}
			} else {
				// Add client key
				updateRequired = true
			}
			merged.ReplicaConfiguration.MysqlReplicaConfiguration.ClientKey = direct.ValueOf(desired.Spec.ReplicaConfiguration.ClientKey)
		} else if existingMysqlReplicaConfig && actual.ReplicaConfiguration.MysqlReplicaConfiguration.ClientKey != "" {
			// Remove client key
			updateRequired = true
		}

		if desired.Spec.ReplicaConfiguration.ConnectRetryInterval != nil {
			if existingMysqlReplicaConfig {
				if direct.ValueOf(desired.Spec.ReplicaConfiguration.ConnectRetryInterval) != actual.ReplicaConfiguration.MysqlReplicaConfiguration.ConnectRetryInterval {
					// Change connect retry interval
					updateRequired = true
				}
			} else {
				// Add connect retry interval
				updateRequired = true
			}
			merged.ReplicaConfiguration.MysqlReplicaConfiguration.ConnectRetryInterval = direct.ValueOf(desired.Spec.ReplicaConfiguration.ConnectRetryInterval)
		} else if existingMysqlReplicaConfig && actual.ReplicaConfiguration.MysqlReplicaConfiguration.ConnectRetryInterval != 0 {
			// Remove connect retry interval
			updateRequired = true
		}

		if desired.Spec.ReplicaConfiguration.DumpFilePath != nil {
			if existingMysqlReplicaConfig {
				if direct.ValueOf(desired.Spec.ReplicaConfiguration.DumpFilePath) != actual.ReplicaConfiguration.MysqlReplicaConfiguration.DumpFilePath {
					// Change dump file path
					updateRequired = true
				}
			} else {
				// Add dump file path
				updateRequired = true
			}
			merged.ReplicaConfiguration.MysqlReplicaConfiguration.DumpFilePath = direct.ValueOf(desired.Spec.ReplicaConfiguration.DumpFilePath)
		} else if existingMysqlReplicaConfig && actual.ReplicaConfiguration.MysqlReplicaConfiguration.DumpFilePath != "" {
			// Remove dump file path
			updateRequired = true
		}

		if desired.Spec.ReplicaConfiguration.MasterHeartbeatPeriod != nil {
			if existingMysqlReplicaConfig {
				if direct.ValueOf(desired.Spec.ReplicaConfiguration.MasterHeartbeatPeriod) != actual.ReplicaConfiguration.MysqlReplicaConfiguration.MasterHeartbeatPeriod {
					// Change master heartbeat period
					updateRequired = true
				}
			} else {
				// Add master heartbeat period
				updateRequired = true
			}
			merged.ReplicaConfiguration.MysqlReplicaConfiguration.MasterHeartbeatPeriod = direct.ValueOf(desired.Spec.ReplicaConfiguration.MasterHeartbeatPeriod)
		} else if existingMysqlReplicaConfig && actual.ReplicaConfiguration.MysqlReplicaConfiguration.MasterHeartbeatPeriod != 0 {
			// Remove master heartbeat period
			updateRequired = true
		}

		if desired.Spec.ReplicaConfiguration.Password != nil {
			if existingMysqlReplicaConfig {
				if refs.replicaPassword != actual.ReplicaConfiguration.MysqlReplicaConfiguration.Password {
					// Change password
					updateRequired = true
				}
			} else {
				// Add password
				updateRequired = true
			}
			merged.ReplicaConfiguration.MysqlReplicaConfiguration.Password = refs.replicaPassword
		} else if existingMysqlReplicaConfig && actual.ReplicaConfiguration.MysqlReplicaConfiguration.Password != "" {
			// Remove password
			updateRequired = true
		}

		if desired.Spec.ReplicaConfiguration.SslCipher != nil {
			if existingMysqlReplicaConfig {
				if direct.ValueOf(desired.Spec.ReplicaConfiguration.SslCipher) != actual.ReplicaConfiguration.MysqlReplicaConfiguration.SslCipher {
					// Change SSL cipher
					updateRequired = true
				}
			} else {
				// Add SSL cipher
				updateRequired = true
			}
			merged.ReplicaConfiguration.MysqlReplicaConfiguration.SslCipher = direct.ValueOf(desired.Spec.ReplicaConfiguration.SslCipher)
		} else if existingMysqlReplicaConfig && actual.ReplicaConfiguration.MysqlReplicaConfiguration.SslCipher != "" {
			// Remove SSL cipher
			updateRequired = true
		}

		if desired.Spec.ReplicaConfiguration.Username != nil {
			if existingMysqlReplicaConfig {
				if direct.ValueOf(desired.Spec.ReplicaConfiguration.Username) != actual.ReplicaConfiguration.MysqlReplicaConfiguration.Username {
					// Change username
					updateRequired = true
				}
			} else {
				// Add username
				updateRequired = true
			}
			merged.ReplicaConfiguration.MysqlReplicaConfiguration.Username = direct.ValueOf(desired.Spec.ReplicaConfiguration.Username)
		} else if existingMysqlReplicaConfig && actual.ReplicaConfiguration.MysqlReplicaConfiguration.Username != "" {
			// Remove username
			updateRequired = true
		}

		if desired.Spec.ReplicaConfiguration.VerifyServerCertificate != nil {
			if existingMysqlReplicaConfig {
				if direct.ValueOf(desired.Spec.ReplicaConfiguration.VerifyServerCertificate) != actual.ReplicaConfiguration.MysqlReplicaConfiguration.VerifyServerCertificate {
					// Change verify server certificate
					updateRequired = true
				}
			} else {
				// Add verify server certificate
				updateRequired = true
			}
			merged.ReplicaConfiguration.MysqlReplicaConfiguration.VerifyServerCertificate = direct.ValueOf(desired.Spec.ReplicaConfiguration.VerifyServerCertificate)
		} else if existingMysqlReplicaConfig && actual.ReplicaConfiguration.MysqlReplicaConfiguration.VerifyServerCertificate {
			// Remove verify server certificate
			updateRequired = true
		}
	} else if existingReplicaConfig {
		// Remove replica configuration
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

	if desired.Spec.Settings.Collation != nil {
		if direct.ValueOf(desired.Spec.Settings.Collation) != actual.Settings.Collation {
			// Change collation
			updateRequired = true
		}
		merged.Settings.Collation = direct.ValueOf(desired.Spec.Settings.Collation)
	} else {
		// Keep collation
		merged.Settings.Collation = actual.Settings.Collation
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

	if desired.Spec.Settings.DataCacheConfig != nil {
		if actual.Settings.DataCacheConfig == nil {
			// Add data cache config
			updateRequired = true
		} else if direct.ValueOf(desired.Spec.Settings.DataCacheConfig.DataCacheEnabled) != actual.Settings.DataCacheConfig.DataCacheEnabled {
			// Change data cache config
			updateRequired = true
		}
		merged.Settings.DataCacheConfig = &api.DataCacheConfig{
			DataCacheEnabled: direct.ValueOf(desired.Spec.Settings.DataCacheConfig.DataCacheEnabled),
		}
	} else if actual.Settings.DataCacheConfig != nil {
		// Remove data cache config
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

	if desired.Spec.Settings.DenyMaintenancePeriod != nil {
		if actual.Settings.DenyMaintenancePeriods == nil {
			// Add deny maintenance period
			updateRequired = true
		} else if (len(actual.Settings.DenyMaintenancePeriods) == 1) && ((desired.Spec.Settings.DenyMaintenancePeriod.EndDate != actual.Settings.DenyMaintenancePeriods[0].EndDate) ||
			(desired.Spec.Settings.DenyMaintenancePeriod.StartDate != actual.Settings.DenyMaintenancePeriods[0].StartDate) ||
			(desired.Spec.Settings.DenyMaintenancePeriod.Time != actual.Settings.DenyMaintenancePeriods[0].Time)) {
			// Change deny maintenance period
			updateRequired = true
		}
		merged.Settings.DenyMaintenancePeriods = []*api.DenyMaintenancePeriod{
			{
				EndDate:   desired.Spec.Settings.DenyMaintenancePeriod.EndDate,
				StartDate: desired.Spec.Settings.DenyMaintenancePeriod.StartDate,
				Time:      desired.Spec.Settings.DenyMaintenancePeriod.Time,
			},
		}
	} else if actual.Settings.DenyMaintenancePeriods != nil {
		// Remove deny maintenance period
		updateRequired = true
	}

	if desired.Spec.Settings.DiskAutoresize != nil {
		if direct.ValueOf(desired.Spec.Settings.DiskAutoresize) != direct.ValueOf(actual.Settings.StorageAutoResize) {
			// Change disk autoresize
			updateRequired = true
		}
		merged.Settings.StorageAutoResize = desired.Spec.Settings.DiskAutoresize
	} else {
		// Keep disk autoresize
		merged.Settings.StorageAutoResize = actual.Settings.StorageAutoResize
	}

	if desired.Spec.Settings.DiskAutoresizeLimit != nil {
		if direct.ValueOf(desired.Spec.Settings.DiskAutoresizeLimit) != actual.Settings.StorageAutoResizeLimit {
			// Change disk autoresize limit
			updateRequired = true
		}
		merged.Settings.StorageAutoResizeLimit = direct.ValueOf(desired.Spec.Settings.DiskAutoresizeLimit)
	} else {
		// Keep disk autoresize limit
		merged.Settings.StorageAutoResizeLimit = actual.Settings.StorageAutoResizeLimit
	}

	if desired.Spec.Settings.DiskSize != nil {
		if direct.ValueOf(desired.Spec.Settings.DiskSize) != actual.Settings.DataDiskSizeGb {
			// Change disk size
			updateRequired = true
		}
		merged.Settings.DataDiskSizeGb = direct.ValueOf(desired.Spec.Settings.DiskSize)
	} else {
		// Keep disk size
		merged.Settings.DataDiskSizeGb = actual.Settings.DataDiskSizeGb
	}

	if desired.Spec.Settings.DiskType != nil {
		if direct.ValueOf(desired.Spec.Settings.DiskType) != actual.Settings.DataDiskType {
			// Change disk type
			updateRequired = true
		}
		merged.Settings.DataDiskType = direct.ValueOf(desired.Spec.Settings.DiskType)
	} else {
		// Keep disk type
		merged.Settings.DataDiskType = actual.Settings.DataDiskType
	}

	if desired.Spec.Settings.Edition != nil {
		if direct.ValueOf(desired.Spec.Settings.Edition) != actual.Settings.Edition {
			// Change edition
			updateRequired = true
		}
		merged.Settings.Edition = direct.ValueOf(desired.Spec.Settings.Edition)
	} else {
		// Keep edition
		merged.Settings.Edition = actual.Settings.Edition
	}

	if desired.Spec.Settings.InsightsConfig != nil {
		if actual.Settings.InsightsConfig == nil {
			// Add insights config
			updateRequired = true
		} else if (direct.ValueOf(desired.Spec.Settings.InsightsConfig.QueryInsightsEnabled) != actual.Settings.InsightsConfig.QueryInsightsEnabled) ||
			(direct.ValueOf(desired.Spec.Settings.InsightsConfig.QueryPlansPerMinute) != actual.Settings.InsightsConfig.QueryPlansPerMinute) ||
			(direct.ValueOf(desired.Spec.Settings.InsightsConfig.QueryStringLength) != actual.Settings.InsightsConfig.QueryStringLength) ||
			(direct.ValueOf(desired.Spec.Settings.InsightsConfig.RecordApplicationTags) != actual.Settings.InsightsConfig.RecordApplicationTags) ||
			(direct.ValueOf(desired.Spec.Settings.InsightsConfig.RecordClientAddress) != actual.Settings.InsightsConfig.RecordClientAddress) {
			// Change insights config
			updateRequired = true
		}
		merged.Settings.InsightsConfig = &api.InsightsConfig{
			QueryInsightsEnabled:  direct.ValueOf(desired.Spec.Settings.InsightsConfig.QueryInsightsEnabled),
			QueryPlansPerMinute:   direct.ValueOf(desired.Spec.Settings.InsightsConfig.QueryPlansPerMinute),
			QueryStringLength:     direct.ValueOf(desired.Spec.Settings.InsightsConfig.QueryStringLength),
			RecordApplicationTags: direct.ValueOf(desired.Spec.Settings.InsightsConfig.RecordApplicationTags),
			RecordClientAddress:   direct.ValueOf(desired.Spec.Settings.InsightsConfig.RecordClientAddress),
		}
	} else if actual.Settings.InsightsConfig != nil {
		// Keep insights config
		merged.Settings.InsightsConfig = actual.Settings.InsightsConfig
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

	if desired.Spec.Settings.PasswordValidationPolicy != nil {
		if actual.Settings.PasswordValidationPolicy == nil {
			// Add password validation policy
			updateRequired = true
		} else if (direct.ValueOf(desired.Spec.Settings.PasswordValidationPolicy.Complexity) != actual.Settings.PasswordValidationPolicy.Complexity) ||
			(direct.ValueOf(desired.Spec.Settings.PasswordValidationPolicy.DisallowUsernameSubstring) != actual.Settings.PasswordValidationPolicy.DisallowUsernameSubstring) ||
			(desired.Spec.Settings.PasswordValidationPolicy.EnablePasswordPolicy != actual.Settings.PasswordValidationPolicy.EnablePasswordPolicy) ||
			(direct.ValueOf(desired.Spec.Settings.PasswordValidationPolicy.MinLength) != actual.Settings.PasswordValidationPolicy.MinLength) ||
			(direct.ValueOf(desired.Spec.Settings.PasswordValidationPolicy.PasswordChangeInterval) != actual.Settings.PasswordValidationPolicy.PasswordChangeInterval) ||
			(direct.ValueOf(desired.Spec.Settings.PasswordValidationPolicy.ReuseInterval) != actual.Settings.PasswordValidationPolicy.ReuseInterval) {
			// Change password validation policy
			updateRequired = true
		}
		merged.Settings.PasswordValidationPolicy = &api.PasswordValidationPolicy{
			Complexity:                direct.ValueOf(desired.Spec.Settings.PasswordValidationPolicy.Complexity),
			DisallowUsernameSubstring: direct.ValueOf(desired.Spec.Settings.PasswordValidationPolicy.DisallowUsernameSubstring),
			EnablePasswordPolicy:      desired.Spec.Settings.PasswordValidationPolicy.EnablePasswordPolicy,
			MinLength:                 direct.ValueOf(desired.Spec.Settings.PasswordValidationPolicy.MinLength),
			PasswordChangeInterval:    direct.ValueOf(desired.Spec.Settings.PasswordValidationPolicy.PasswordChangeInterval),
			ReuseInterval:             direct.ValueOf(desired.Spec.Settings.PasswordValidationPolicy.ReuseInterval),
		}
	} else if actual.Settings.PasswordValidationPolicy != nil {
		// Remove password validation policy
		updateRequired = true
	}

	if desired.Spec.Settings.PricingPlan != nil {
		if actual.Settings.PricingPlan == "" {
			// Add pricing plan
			updateRequired = true
		} else if direct.ValueOf(desired.Spec.Settings.PricingPlan) != actual.Settings.PricingPlan {
			// Change pricing plan
			updateRequired = true
		}
		merged.Settings.PricingPlan = direct.ValueOf(desired.Spec.Settings.PricingPlan)
	} else if actual.Settings.PricingPlan != "" {
		// Keep pricing plan
		merged.Settings.PricingPlan = actual.Settings.PricingPlan
	}

	if desired.Spec.Settings.ReplicationType != nil {
		if actual.Settings.ReplicationType == "" {
			// Add replication type
			updateRequired = true
		} else if direct.ValueOf(desired.Spec.Settings.ReplicationType) != actual.Settings.ReplicationType {
			// Change replication type
			updateRequired = true
		}
		merged.Settings.ReplicationType = direct.ValueOf(desired.Spec.Settings.ReplicationType)
	} else if actual.Settings.ReplicationType != "" {
		// Keep replication type
		merged.Settings.ReplicationType = actual.Settings.ReplicationType
	}

	if desired.Spec.Settings.SqlServerAuditConfig != nil {
		if actual.Settings.SqlServerAuditConfig == nil {
			// Add sql server audit config
			updateRequired = true
		} else if (refs.auditLogBucket != actual.Settings.SqlServerAuditConfig.Bucket) ||
			(direct.ValueOf(desired.Spec.Settings.SqlServerAuditConfig.RetentionInterval) != actual.Settings.SqlServerAuditConfig.RetentionInterval) ||
			(direct.ValueOf(desired.Spec.Settings.SqlServerAuditConfig.UploadInterval) != actual.Settings.SqlServerAuditConfig.UploadInterval) {
			// Change sql server audit config
			updateRequired = true
		}
		merged.Settings.SqlServerAuditConfig = &api.SqlServerAuditConfig{
			Bucket:            refs.auditLogBucket,
			RetentionInterval: direct.ValueOf(desired.Spec.Settings.SqlServerAuditConfig.RetentionInterval),
			UploadInterval:    direct.ValueOf(desired.Spec.Settings.SqlServerAuditConfig.UploadInterval),
			Kind:              "sql#sqlServerAuditConfig",
		}
	} else if actual.Settings.SqlServerAuditConfig != nil {
		// Remove sql server audit config
		updateRequired = true
	}

	if desired.Spec.Settings.Tier != "" {
		if actual.Settings.Tier == "" {
			// Add tier
			updateRequired = true
		} else if desired.Spec.Settings.Tier != actual.Settings.Tier {
			// Change tier
			updateRequired = true
		}
		merged.Settings.Tier = desired.Spec.Settings.Tier
	} else {
		// Keep tier
		merged.Settings.Tier = actual.Settings.Tier
	}

	if desired.Spec.Settings.TimeZone != nil {
		if actual.Settings.TimeZone == "" {
			// Add time zone
			updateRequired = true
		} else if direct.ValueOf(desired.Spec.Settings.TimeZone) != actual.Settings.TimeZone {
			// Change time zone
			updateRequired = true
		}
		merged.Settings.TimeZone = direct.ValueOf(desired.Spec.Settings.TimeZone)
	} else {
		// Keep time zone
		merged.Settings.TimeZone = actual.Settings.TimeZone
	}

	if !reflect.DeepEqual(desired.Labels, actual.Settings.UserLabels) {
		updateRequired = true
	}
	merged.Settings.UserLabels = desired.Labels

	return merged, updateRequired, nil
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

func PointersMatch[T any](desired *T, actual *T) bool {
	if (desired == nil && actual != nil) || (desired != nil && actual == nil) {
		// Pointers are not matching if one is nil and the other is not nil.
		return false
	}
	// Otherwise, they match. Either both are nil, or both are not nil.
	return true
}
