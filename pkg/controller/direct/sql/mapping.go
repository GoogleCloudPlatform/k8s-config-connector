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

	api "google.golang.org/api/sqladmin/v1beta4"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func SQLInstanceKRMToGCP(in *krm.SQLInstance, refs *SQLInstanceInternalRefs) (*api.DatabaseInstance, error) {
	out := &api.DatabaseInstance{}

	if in == nil {
		return nil, fmt.Errorf("cannot convert nil SQLInstance")
	}

	if in.Spec.CloneSource != nil {
		// If spec.cloneSource is specified, it's invalid to convert krm.SQLInstance -> api.DatabaseInstance.
		// Instead, the krm.SQLInstance should be converted to an api.InstancesCloneRequest.
		return nil, fmt.Errorf("cannot convert SQLInstance with CloneSource specified")
	}

	if in.Spec.DatabaseVersion != nil {
		out.DatabaseVersion = *in.Spec.DatabaseVersion
	}

	if in.Spec.EncryptionKMSCryptoKeyRef != nil {
		out.DiskEncryptionConfiguration = &api.DiskEncryptionConfiguration{
			Kind:       "sql#diskEncryptionConfiguration",
			KmsKeyName: refs.cryptoKey,
		}
	}

	if in.Spec.InstanceType != nil {
		out.InstanceType = *in.Spec.InstanceType
	}

	if in.Spec.MaintenanceVersion != nil {
		out.MaintenanceVersion = *in.Spec.MaintenanceVersion
	}

	if in.Spec.MasterInstanceRef != nil {
		out.MasterInstanceName = refs.masterInstance
	}

	if in.Spec.Region != nil {
		out.Region = *in.Spec.Region
	}

	out.ReplicaConfiguration = InstanceReplicaConfigurationKRMToGCP(in.Spec.ReplicaConfiguration, refs)

	if in.Spec.ResourceID != nil {
		out.Name = *in.Spec.ResourceID
	} else {
		return nil, fmt.Errorf("resourceID is empty")
	}

	if in.Spec.RootPassword != nil && refs.rootPassword != "" {
		out.RootPassword = refs.rootPassword
	}

	out.Settings = &api.Settings{}

	if in.Spec.Settings.ActivationPolicy != nil {
		out.Settings.ActivationPolicy = *in.Spec.Settings.ActivationPolicy
	}

	if in.Spec.Settings.ActiveDirectoryConfig != nil {
		// todo: requires sqlserver
		// todo: requires private network
		out.Settings.ActiveDirectoryConfig = &api.SqlActiveDirectoryConfig{
			Domain: in.Spec.Settings.ActiveDirectoryConfig.Domain,
		}
	}

	if in.Spec.Settings.AdvancedMachineFeatures != nil {
		// todo: requires sqlserver
		// todo: requires >= 6 cpu cores
		out.Settings.AdvancedMachineFeatures = &api.AdvancedMachineFeatures{}
		if in.Spec.Settings.AdvancedMachineFeatures.ThreadsPerCore != nil {
			out.Settings.AdvancedMachineFeatures.ThreadsPerCore = *in.Spec.Settings.AdvancedMachineFeatures.ThreadsPerCore
		}
	}

	if in.Spec.Settings.AuthorizedGaeApplications != nil {
		// todo: deprecated
		out.Settings.AuthorizedGaeApplications = in.Spec.Settings.AuthorizedGaeApplications
	}

	if in.Spec.Settings.AvailabilityType != nil {
		out.Settings.AvailabilityType = *in.Spec.Settings.AvailabilityType
	}

	if in.Spec.Settings.BackupConfiguration != nil {
		backupConfig := &api.BackupConfiguration{}

		if in.Spec.Settings.BackupConfiguration.BackupRetentionSettings != nil {
			retentionSettings := &api.BackupRetentionSettings{
				RetainedBackups: in.Spec.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups,
			}
			if in.Spec.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit != nil {
				retentionSettings.RetentionUnit = *in.Spec.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit
			}
			backupConfig.BackupRetentionSettings = retentionSettings
		}

		if in.Spec.Settings.BackupConfiguration.BinaryLogEnabled != nil {
			// todo: requires mysql
			backupConfig.BinaryLogEnabled = *in.Spec.Settings.BackupConfiguration.BinaryLogEnabled
		}

		if in.Spec.Settings.BackupConfiguration.Enabled != nil {
			backupConfig.Enabled = *in.Spec.Settings.BackupConfiguration.Enabled
		}

		if in.Spec.Settings.BackupConfiguration.Location != nil {
			backupConfig.Location = *in.Spec.Settings.BackupConfiguration.Location
		}

		if in.Spec.Settings.BackupConfiguration.PointInTimeRecoveryEnabled != nil {
			backupConfig.PointInTimeRecoveryEnabled = *in.Spec.Settings.BackupConfiguration.PointInTimeRecoveryEnabled
		}

		if in.Spec.Settings.BackupConfiguration.StartTime != nil {
			backupConfig.StartTime = *in.Spec.Settings.BackupConfiguration.StartTime
		}

		if in.Spec.Settings.BackupConfiguration.TransactionLogRetentionDays != nil {
			backupConfig.TransactionLogRetentionDays = *in.Spec.Settings.BackupConfiguration.TransactionLogRetentionDays
		}

		out.Settings.BackupConfiguration = backupConfig
	}

	out.Settings.Collation = direct.ValueOf(in.Spec.Settings.Collation)

	if in.Spec.Settings.ConnectorEnforcement != nil {
		out.Settings.ConnectorEnforcement = *in.Spec.Settings.ConnectorEnforcement
	}

	if in.Spec.Settings.CrashSafeReplication != nil {
		// todo: deprecated
		out.Settings.CrashSafeReplicationEnabled = *in.Spec.Settings.CrashSafeReplication
	}

	out.Settings.DataCacheConfig = InstanceDataCacheConfigKRMToGCP(in.Spec.Settings.DataCacheConfig)

	if in.Spec.Settings.DatabaseFlags != nil {
		dbFlags := []*api.DatabaseFlags{}
		for _, dbFlag := range in.Spec.Settings.DatabaseFlags {
			dbFlags = append(dbFlags, &api.DatabaseFlags{
				Name:  dbFlag.Name,
				Value: dbFlag.Value,
			})
		}
		out.Settings.DatabaseFlags = dbFlags
	}

	if in.Spec.Settings.DeletionProtectionEnabled != nil {
		out.Settings.DeletionProtectionEnabled = *in.Spec.Settings.DeletionProtectionEnabled
	}

	if in.Spec.Settings.DenyMaintenancePeriod != nil {
		// todo: handle multiple periods (?)
		out.Settings.DenyMaintenancePeriods = []*api.DenyMaintenancePeriod{
			{
				EndDate:   in.Spec.Settings.DenyMaintenancePeriod.EndDate,
				StartDate: in.Spec.Settings.DenyMaintenancePeriod.StartDate,
				Time:      in.Spec.Settings.DenyMaintenancePeriod.Time,
			},
		}
	}

	if in.Spec.Settings.DiskAutoresize != nil {
		out.Settings.StorageAutoResize = in.Spec.Settings.DiskAutoresize
	}

	if in.Spec.Settings.DiskAutoresizeLimit != nil {
		// todo: requires DiskAutoresize == true
		out.Settings.StorageAutoResizeLimit = *in.Spec.Settings.DiskAutoresizeLimit
	}

	if in.Spec.Settings.DiskSize != nil {
		out.Settings.DataDiskSizeGb = *in.Spec.Settings.DiskSize
	}

	if in.Spec.Settings.DiskType != nil {
		out.Settings.DataDiskType = *in.Spec.Settings.DiskType
	}

	out.Settings.Edition = direct.ValueOf(in.Spec.Settings.Edition)
	out.Settings.InsightsConfig = InstanceInsightsConfigKRMToGCP(in.Spec.Settings.InsightsConfig)
	out.Settings.IpConfiguration = InstanceIpConfigurationKRMToGCP(in.Spec.Settings.IpConfiguration, refs)
	out.Settings.LocationPreference = InstanceLocationPreferenceKRMToGCP(in.Spec.Settings.LocationPreference)
	out.Settings.MaintenanceWindow = InstanceMaintenanceWindowKRMToGCP(in.Spec.Settings.MaintenanceWindow)
	out.Settings.PasswordValidationPolicy = InstancePasswordValidationPolicyKRMToGCP(in.Spec.Settings.PasswordValidationPolicy)
	out.Settings.PricingPlan = direct.ValueOf(in.Spec.Settings.PricingPlan)
	out.Settings.ReplicationType = direct.ValueOf(in.Spec.Settings.ReplicationType)
	out.Settings.SqlServerAuditConfig = InstanceSqlServerAuditConfigKRMToGCP(in.Spec.Settings.SqlServerAuditConfig, refs)
	out.Settings.Tier = in.Spec.Settings.Tier
	out.Settings.TimeZone = direct.ValueOf(in.Spec.Settings.TimeZone)

	if in.Labels != nil {
		out.Settings.UserLabels = make(map[string]string)
		for k, v := range in.Labels {
			out.Settings.UserLabels[k] = v
		}
	}

	return out, nil
}

func InstanceReplicaConfigurationKRMToGCP(in *krm.InstanceReplicaConfiguration, refs *SQLInstanceInternalRefs) *api.ReplicaConfiguration {
	if in == nil {
		return nil
	}

	out := &api.ReplicaConfiguration{
		Kind: "sql#replicaConfiguration",
		// CascadableReplica is not supported in KRM API.
		FailoverTarget:            direct.ValueOf(in.FailoverTarget),
		MysqlReplicaConfiguration: InstanceMysqlReplicaConfigurationKRMToGCP(in, refs),
	}

	if in.FailoverTarget != nil {
		out.ForceSendFields = append(out.ForceSendFields, "FailoverTarget")
	}

	return out
}

func InstanceMysqlReplicaConfigurationKRMToGCP(in *krm.InstanceReplicaConfiguration, refs *SQLInstanceInternalRefs) *api.MySqlReplicaConfiguration {
	if in == nil {
		return nil
	}

	// For some reason, the KRM API embeds all of the MySqlReplicaConfiguration fields into the
	// InstanceReplicaConfiguration object (instead of using a separate object). Therefore, we
	// need to check for each of the individual fields here.
	if in.CaCertificate == nil &&
		in.ClientCertificate == nil &&
		in.ClientKey == nil &&
		in.ConnectRetryInterval == nil &&
		in.DumpFilePath == nil &&
		in.MasterHeartbeatPeriod == nil &&
		in.Password == nil &&
		in.SslCipher == nil &&
		in.Username == nil &&
		in.VerifyServerCertificate == nil {
		return nil
	}

	out := &api.MySqlReplicaConfiguration{
		Kind:                    "sql#mysqlReplicaConfiguration",
		CaCertificate:           direct.ValueOf(in.CaCertificate),
		ClientCertificate:       direct.ValueOf(in.ClientCertificate),
		ClientKey:               direct.ValueOf(in.ClientKey),
		ConnectRetryInterval:    direct.ValueOf(in.ConnectRetryInterval),
		DumpFilePath:            direct.ValueOf(in.DumpFilePath),
		MasterHeartbeatPeriod:   direct.ValueOf(in.MasterHeartbeatPeriod),
		SslCipher:               direct.ValueOf(in.SslCipher),
		Username:                direct.ValueOf(in.Username),
		VerifyServerCertificate: direct.ValueOf(in.VerifyServerCertificate),
	}

	// todo: embed refs in krm object external fields, remove this
	if in.Password != nil {
		out.Password = refs.replicaPassword
	}

	if in.ConnectRetryInterval != nil {
		out.ForceSendFields = append(out.ForceSendFields, "ConnectRetryInterval")
	}
	if in.MasterHeartbeatPeriod != nil {
		out.ForceSendFields = append(out.ForceSendFields, "MasterHeartbeatPeriod")
	}
	if in.VerifyServerCertificate != nil {
		out.ForceSendFields = append(out.ForceSendFields, "VerifyServerCertificate")
	}

	return out
}

func InstanceDataCacheConfigKRMToGCP(in *krm.InstanceDataCacheConfig) *api.DataCacheConfig {
	if in == nil {
		return nil
	}

	out := &api.DataCacheConfig{
		DataCacheEnabled: direct.ValueOf(in.DataCacheEnabled),
	}

	if in.DataCacheEnabled != nil {
		out.ForceSendFields = append(out.ForceSendFields, "DataCacheEnabled")
	}

	return out
}

func InstanceInsightsConfigKRMToGCP(in *krm.InstanceInsightsConfig) *api.InsightsConfig {
	if in == nil {
		return nil
	}

	out := &api.InsightsConfig{
		QueryInsightsEnabled:  direct.ValueOf(in.QueryInsightsEnabled),
		QueryPlansPerMinute:   direct.ValueOf(in.QueryPlansPerMinute),
		QueryStringLength:     direct.ValueOf(in.QueryStringLength),
		RecordApplicationTags: direct.ValueOf(in.RecordApplicationTags),
		RecordClientAddress:   direct.ValueOf(in.RecordClientAddress),
	}

	if in.QueryInsightsEnabled != nil {
		out.ForceSendFields = append(out.ForceSendFields, "QueryInsightsEnabled")
	}
	if in.QueryPlansPerMinute != nil {
		out.ForceSendFields = append(out.ForceSendFields, "QueryPlansPerMinute")
	}
	if in.QueryStringLength != nil {
		out.ForceSendFields = append(out.ForceSendFields, "QueryStringLength")
	}
	if in.RecordApplicationTags != nil {
		out.ForceSendFields = append(out.ForceSendFields, "RecordApplicationTags")
	}
	if in.RecordClientAddress != nil {
		out.ForceSendFields = append(out.ForceSendFields, "RecordClientAddress")
	}

	return out
}

func InstanceIpConfigurationKRMToGCP(in *krm.InstanceIpConfiguration, refs *SQLInstanceInternalRefs) *api.IpConfiguration {
	if in == nil {
		return nil
	}

	out := &api.IpConfiguration{
		AllocatedIpRange:                        direct.ValueOf(in.AllocatedIpRange),
		AuthorizedNetworks:                      InstanceAuthorizedNetworksKRMToGCP(in.AuthorizedNetworks),
		EnablePrivatePathForGoogleCloudServices: direct.ValueOf(in.EnablePrivatePathForGoogleCloudServices),
		Ipv4Enabled:                             direct.ValueOf(in.Ipv4Enabled),
		PscConfig:                               InstancePscConfigKRMToGCP(in.PscConfig),
		RequireSsl:                              direct.ValueOf(in.RequireSsl),
		SslMode:                                 direct.ValueOf(in.SslMode),
	}

	// todo: embed refs in krm object external fields, remove this
	if in.PrivateNetworkRef != nil {
		out.PrivateNetwork = refs.privateNetwork
	}

	if in.EnablePrivatePathForGoogleCloudServices != nil {
		out.ForceSendFields = append(out.ForceSendFields, "EnablePrivatePathForGoogleCloudServices")
	}
	if in.Ipv4Enabled != nil {
		out.ForceSendFields = append(out.ForceSendFields, "Ipv4Enabled")
	}
	if in.RequireSsl != nil {
		out.ForceSendFields = append(out.ForceSendFields, "RequireSsl")
	}

	return out
}

func InstanceAuthorizedNetworksKRMToGCP(in []krm.InstanceAuthorizedNetworks) []*api.AclEntry {
	out := []*api.AclEntry{}
	for _, net := range in {
		out = append(out, &api.AclEntry{
			Kind:           "sql#aclEntry",
			ExpirationTime: direct.ValueOf(net.ExpirationTime),
			Name:           direct.ValueOf(net.Name),
			Value:          net.Value,
		})
	}
	return out
}

func InstancePscConfigKRMToGCP(in []krm.InstancePscConfig) *api.PscConfig {
	if len(in) < 1 {
		return nil
	}

	// Note:  For some reason, the KRM API allows []InstancePscConfig. However, in the GCP proto there is only
	// a single *api.PscConfig. I think there is a mistake in the KRM API; it should not allow a list. For
	// now, we will only use the first item in the []InstancePscConfig list.
	inFixed := in[0]

	out := &api.PscConfig{
		AllowedConsumerProjects: inFixed.AllowedConsumerProjects,
		PscEnabled:              direct.ValueOf(inFixed.PscEnabled),
	}

	if inFixed.PscEnabled != nil {
		out.ForceSendFields = append(out.ForceSendFields, "PscEnabled")
	}

	return out
}

func InstanceLocationPreferenceKRMToGCP(in *krm.InstanceLocationPreference) *api.LocationPreference {
	if in == nil {
		return nil
	}

	out := &api.LocationPreference{
		Kind:                 "sql#locationPreference",
		FollowGaeApplication: direct.ValueOf(in.FollowGaeApplication),
		SecondaryZone:        direct.ValueOf(in.SecondaryZone),
		Zone:                 direct.ValueOf(in.Zone),
	}

	return out
}

func InstanceMaintenanceWindowKRMToGCP(in *krm.InstanceMaintenanceWindow) *api.MaintenanceWindow {
	if in == nil {
		return nil
	}

	out := &api.MaintenanceWindow{
		Kind:        "sql#maintenanceWindow",
		Day:         direct.ValueOf(in.Day),
		Hour:        direct.ValueOf(in.Hour),
		UpdateTrack: direct.ValueOf(in.UpdateTrack),
	}

	if in.Hour != nil {
		out.ForceSendFields = append(out.ForceSendFields, "Hour")
	}

	return out
}

func InstancePasswordValidationPolicyKRMToGCP(in *krm.InstancePasswordValidationPolicy) *api.PasswordValidationPolicy {
	if in == nil {
		return nil
	}

	out := &api.PasswordValidationPolicy{
		Complexity: direct.ValueOf(in.Complexity),
		// DisallowCompromisedCredentials is not supported in KRM API.
		DisallowUsernameSubstring: direct.ValueOf(in.DisallowUsernameSubstring),
		EnablePasswordPolicy:      in.EnablePasswordPolicy,
		MinLength:                 direct.ValueOf(in.MinLength),
		PasswordChangeInterval:    direct.ValueOf(in.PasswordChangeInterval),
		ReuseInterval:             direct.ValueOf(in.ReuseInterval),
	}

	if in.DisallowUsernameSubstring != nil {
		out.ForceSendFields = append(out.ForceSendFields, "DisallowUsernameSubstring")
	}
	if !in.EnablePasswordPolicy {
		out.ForceSendFields = append(out.ForceSendFields, "EnablePasswordPolicy")
	}
	if in.MinLength != nil {
		out.ForceSendFields = append(out.ForceSendFields, "MinLength")
	}
	if in.ReuseInterval != nil {
		out.ForceSendFields = append(out.ForceSendFields, "ReuseInterval")
	}

	return out
}

func InstanceSqlServerAuditConfigKRMToGCP(in *krm.InstanceSqlServerAuditConfig, refs *SQLInstanceInternalRefs) *api.SqlServerAuditConfig {
	if in == nil {
		return nil
	}

	out := &api.SqlServerAuditConfig{
		Kind:              "sql#sqlServerAuditConfig",
		RetentionInterval: direct.ValueOf(in.RetentionInterval),
		UploadInterval:    direct.ValueOf(in.UploadInterval),
	}

	// todo: embed refs in krm object external fields, remove this
	if in.BucketRef != nil {
		out.Bucket = refs.auditLogBucket
	}

	return out
}

func SQLInstanceGCPToKRM(in *api.DatabaseInstance) (*krm.SQLInstance, error) {
	out := &krm.SQLInstance{}

	if in == nil {
		return nil, fmt.Errorf("cannot convert nil DatabaseInstance")
	}

	if in.DatabaseVersion != "" {
		out.Spec.DatabaseVersion = &in.DatabaseVersion
	}

	if in.DiskEncryptionConfiguration != nil {
		out.Spec.EncryptionKMSCryptoKeyRef = &refs.KMSCryptoKeyRef{
			External: in.DiskEncryptionConfiguration.KmsKeyName,
		}
	}

	if in.InstanceType != "" {
		out.Spec.InstanceType = &in.InstanceType
	}

	if in.MaintenanceVersion != "" {
		out.Spec.MaintenanceVersion = &in.MaintenanceVersion
	}

	if in.MasterInstanceName != "" {
		out.Spec.MasterInstanceRef = &refs.SQLInstanceRef{
			External: in.MasterInstanceName,
		}
	}

	if in.Region != "" {
		out.Spec.Region = &in.Region
	}

	out.Spec.ReplicaConfiguration = InstanceReplicaConfigurationGCPToKRM(in.ReplicaConfiguration)

	out.Spec.ResourceID = &in.Name

	out.Spec.RootPassword = &krm.InstanceRootPassword{
		Value: &in.RootPassword,
	}

	if in.Settings.ActivationPolicy != "" {
		out.Spec.Settings.ActivationPolicy = &in.Settings.ActivationPolicy
	}

	if in.Settings.ActiveDirectoryConfig != nil {
		out.Spec.Settings.ActiveDirectoryConfig = &krm.InstanceActiveDirectoryConfig{
			Domain: in.Settings.ActiveDirectoryConfig.Domain,
		}
	}

	if in.Settings.AdvancedMachineFeatures != nil {
		out.Spec.Settings.AdvancedMachineFeatures = &krm.InstanceAdvancedMachineFeatures{
			ThreadsPerCore: &in.Settings.AdvancedMachineFeatures.ThreadsPerCore,
		}
	}

	if in.Settings.AuthorizedGaeApplications != nil {
		out.Spec.Settings.AuthorizedGaeApplications = in.Settings.AuthorizedGaeApplications
	}

	if in.Settings.AvailabilityType != "" {
		out.Spec.Settings.AvailabilityType = &in.Settings.AvailabilityType
	}

	if in.Settings.BackupConfiguration != nil {
		bc := &krm.InstanceBackupConfiguration{}

		if in.Settings.BackupConfiguration.BackupRetentionSettings != nil {
			bc.BackupRetentionSettings = &krm.InstanceBackupRetentionSettings{
				RetainedBackups: in.Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups,
				RetentionUnit:   &in.Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit,
			}
		}

		bc.BinaryLogEnabled = &in.Settings.BackupConfiguration.BinaryLogEnabled
		bc.Enabled = &in.Settings.BackupConfiguration.Enabled
		bc.Location = &in.Settings.BackupConfiguration.Location
		bc.PointInTimeRecoveryEnabled = &in.Settings.BackupConfiguration.PointInTimeRecoveryEnabled
		bc.StartTime = &in.Settings.BackupConfiguration.StartTime
		bc.TransactionLogRetentionDays = &in.Settings.BackupConfiguration.TransactionLogRetentionDays

		out.Spec.Settings.BackupConfiguration = bc
	}

	out.Spec.Settings.Collation = direct.LazyPtr(in.Settings.Collation)

	if in.Settings.ConnectorEnforcement != "" {
		out.Spec.Settings.ConnectorEnforcement = &in.Settings.ConnectorEnforcement
	}

	out.Spec.Settings.CrashSafeReplication = &in.Settings.CrashSafeReplicationEnabled

	out.Spec.Settings.DataCacheConfig = InstanceDataCacheConfigGCPToKRM(in.Settings.DataCacheConfig)

	if in.Settings.DatabaseFlags != nil {
		dbFlags := []krm.InstanceDatabaseFlags{}
		for _, dbFlag := range in.Settings.DatabaseFlags {
			dbFlags = append(dbFlags, krm.InstanceDatabaseFlags{
				Name:  dbFlag.Name,
				Value: dbFlag.Value,
			})
		}
		out.Spec.Settings.DatabaseFlags = dbFlags
	}

	out.Spec.Settings.DeletionProtectionEnabled = &in.Settings.DeletionProtectionEnabled

	// todo: handle multiple periods (?)
	if in.Settings.DenyMaintenancePeriods != nil && len(in.Settings.DenyMaintenancePeriods) >= 1 {
		out.Spec.Settings.DenyMaintenancePeriod = &krm.InstanceDenyMaintenancePeriod{
			EndDate:   in.Settings.DenyMaintenancePeriods[0].EndDate,
			StartDate: in.Settings.DenyMaintenancePeriods[0].StartDate,
			Time:      in.Settings.DenyMaintenancePeriods[0].Time,
		}
	}

	out.Spec.Settings.DiskAutoresize = in.Settings.StorageAutoResize

	if in.Settings.StorageAutoResizeLimit != 0 {
		out.Spec.Settings.DiskAutoresizeLimit = &in.Settings.StorageAutoResizeLimit
	}

	out.Spec.Settings.DiskSize = &in.Settings.DataDiskSizeGb

	if in.Settings.DataDiskType != "" {
		out.Spec.Settings.DiskType = &in.Settings.DataDiskType
	}

	out.Spec.Settings.Edition = direct.LazyPtr(in.Settings.Edition)
	out.Spec.Settings.InsightsConfig = InstanceInsightsConfigGCPToKRM(in.Settings.InsightsConfig)
	out.Spec.Settings.IpConfiguration = InstanceIpConfigurationGCPToKRM(in.Settings.IpConfiguration)
	out.Spec.Settings.LocationPreference = InstanceLocationPreferenceGCPToKRM(in.Settings.LocationPreference)
	out.Spec.Settings.MaintenanceWindow = InstanceMaintenanceWindowGCPToKRM(in.Settings.MaintenanceWindow)
	out.Spec.Settings.PasswordValidationPolicy = InstancePasswordValidationPolicyGCPToKRM(in.Settings.PasswordValidationPolicy)
	out.Spec.Settings.PricingPlan = direct.LazyPtr(in.Settings.PricingPlan)
	out.Spec.Settings.ReplicationType = direct.LazyPtr(in.Settings.ReplicationType)
	out.Spec.Settings.SqlServerAuditConfig = InstanceSqlServerAuditConfigGCPToKRM(in.Settings.SqlServerAuditConfig)
	out.Spec.Settings.Tier = in.Settings.Tier
	out.Spec.Settings.TimeZone = direct.LazyPtr(in.Settings.TimeZone)

	if in.Settings.UserLabels != nil {
		out.Labels = in.Settings.UserLabels
	}

	return out, nil
}

func InstanceReplicaConfigurationGCPToKRM(in *api.ReplicaConfiguration) *krm.InstanceReplicaConfiguration {
	if in == nil {
		return nil
	}

	irc := &krm.InstanceReplicaConfiguration{
		// CascadableReplica is not supported in KRM API.
		FailoverTarget: direct.PtrTo(in.FailoverTarget),
	}

	// For some reason, the KRM API embeds all of the MySqlReplicaConfiguration fields into the
	// InstanceReplicaConfiguration object (instead of using a separate object). Therefore, we
	// need to merge all of the fields here.
	mrc := InstanceMysqlReplicaConfigurationGCPToKRM(in.MysqlReplicaConfiguration)
	if mrc == nil {
		mrc = &krm.InstanceReplicaConfiguration{}
	}

	out := &krm.InstanceReplicaConfiguration{
		CaCertificate:           mrc.CaCertificate,
		ClientCertificate:       mrc.ClientCertificate,
		ClientKey:               mrc.ClientKey,
		ConnectRetryInterval:    mrc.ConnectRetryInterval,
		DumpFilePath:            mrc.DumpFilePath,
		FailoverTarget:          irc.FailoverTarget,
		MasterHeartbeatPeriod:   mrc.MasterHeartbeatPeriod,
		Password:                mrc.Password,
		SslCipher:               mrc.SslCipher,
		Username:                mrc.Username,
		VerifyServerCertificate: mrc.VerifyServerCertificate,
	}

	return out
}

func InstanceMysqlReplicaConfigurationGCPToKRM(in *api.MySqlReplicaConfiguration) *krm.InstanceReplicaConfiguration {
	if in == nil {
		return nil
	}

	out := &krm.InstanceReplicaConfiguration{
		CaCertificate:           direct.LazyPtr(in.CaCertificate),
		ClientCertificate:       direct.LazyPtr(in.ClientCertificate),
		ClientKey:               direct.LazyPtr(in.ClientKey),
		ConnectRetryInterval:    direct.PtrTo(in.ConnectRetryInterval),
		DumpFilePath:            direct.LazyPtr(in.DumpFilePath),
		MasterHeartbeatPeriod:   direct.PtrTo(in.MasterHeartbeatPeriod),
		SslCipher:               direct.LazyPtr(in.SslCipher),
		Username:                direct.LazyPtr(in.Username),
		VerifyServerCertificate: direct.PtrTo(in.VerifyServerCertificate),
	}

	// Note: Password is not exported.

	return out
}

func InstanceDataCacheConfigGCPToKRM(in *api.DataCacheConfig) *krm.InstanceDataCacheConfig {
	if in == nil {
		return nil
	}

	out := &krm.InstanceDataCacheConfig{
		DataCacheEnabled: direct.PtrTo(in.DataCacheEnabled),
	}

	return out
}

func InstanceInsightsConfigGCPToKRM(in *api.InsightsConfig) *krm.InstanceInsightsConfig {
	if in == nil {
		return nil
	}

	out := &krm.InstanceInsightsConfig{
		QueryInsightsEnabled:  direct.PtrTo(in.QueryInsightsEnabled),
		QueryPlansPerMinute:   direct.PtrTo(in.QueryPlansPerMinute),
		QueryStringLength:     direct.PtrTo(in.QueryStringLength),
		RecordApplicationTags: direct.PtrTo(in.RecordApplicationTags),
		RecordClientAddress:   direct.PtrTo(in.RecordClientAddress),
	}

	return out
}

func InstanceIpConfigurationGCPToKRM(in *api.IpConfiguration) *krm.InstanceIpConfiguration {
	if in == nil {
		return nil
	}

	out := &krm.InstanceIpConfiguration{
		AllocatedIpRange:                        direct.LazyPtr(in.AllocatedIpRange),
		AuthorizedNetworks:                      InstanceAuthorizedNetworksGCPToKRM(in.AuthorizedNetworks),
		EnablePrivatePathForGoogleCloudServices: direct.PtrTo(in.EnablePrivatePathForGoogleCloudServices),
		Ipv4Enabled:                             direct.PtrTo(in.Ipv4Enabled),
		PscConfig:                               InstancePscConfigGCPToKRM(in.PscConfig),
		RequireSsl:                              direct.PtrTo(in.RequireSsl),
		SslMode:                                 direct.LazyPtr(in.SslMode),
	}

	if in.PrivateNetwork != "" {
		out.PrivateNetworkRef = &refs.ComputeNetworkRef{
			External: in.PrivateNetwork,
		}
	}

	return out
}

func InstanceAuthorizedNetworksGCPToKRM(in []*api.AclEntry) []krm.InstanceAuthorizedNetworks {
	out := []krm.InstanceAuthorizedNetworks{}
	for _, net := range in {
		out = append(out, krm.InstanceAuthorizedNetworks{
			ExpirationTime: direct.LazyPtr(net.ExpirationTime),
			Name:           direct.LazyPtr(net.Name),
			Value:          net.Value,
		})
	}
	return out
}

func InstancePscConfigGCPToKRM(in *api.PscConfig) []krm.InstancePscConfig {
	if in == nil {
		return nil
	}

	out := []krm.InstancePscConfig{
		{
			AllowedConsumerProjects: in.AllowedConsumerProjects,
			PscEnabled:              direct.PtrTo(in.PscEnabled),
		},
	}

	return out
}

func InstanceLocationPreferenceGCPToKRM(in *api.LocationPreference) *krm.InstanceLocationPreference {
	if in == nil {
		return nil
	}

	out := &krm.InstanceLocationPreference{
		FollowGaeApplication: direct.LazyPtr(in.FollowGaeApplication),
		SecondaryZone:        direct.LazyPtr(in.SecondaryZone),
		Zone:                 direct.LazyPtr(in.Zone),
	}

	return out
}

func InstanceMaintenanceWindowGCPToKRM(in *api.MaintenanceWindow) *krm.InstanceMaintenanceWindow {
	if in == nil {
		return nil
	}

	out := &krm.InstanceMaintenanceWindow{
		Day:         direct.LazyPtr(in.Day),
		Hour:        direct.PtrTo(in.Hour),
		UpdateTrack: direct.LazyPtr(in.UpdateTrack),
	}

	return out
}

func InstancePasswordValidationPolicyGCPToKRM(in *api.PasswordValidationPolicy) *krm.InstancePasswordValidationPolicy {
	if in == nil {
		return nil
	}

	out := &krm.InstancePasswordValidationPolicy{
		Complexity: direct.LazyPtr(in.Complexity),
		// DisallowCompromisedCredentials is not supported in KRM API.
		DisallowUsernameSubstring: direct.PtrTo(in.DisallowUsernameSubstring),
		EnablePasswordPolicy:      in.EnablePasswordPolicy,
		MinLength:                 direct.PtrTo(in.MinLength),
		PasswordChangeInterval:    direct.LazyPtr(in.PasswordChangeInterval),
		ReuseInterval:             direct.PtrTo(in.ReuseInterval),
	}

	return out
}

func InstanceSqlServerAuditConfigGCPToKRM(in *api.SqlServerAuditConfig) *krm.InstanceSqlServerAuditConfig {
	if in == nil {
		return nil
	}

	out := &krm.InstanceSqlServerAuditConfig{
		RetentionInterval: direct.LazyPtr(in.RetentionInterval),
		UploadInterval:    direct.LazyPtr(in.UploadInterval),
	}

	if in.Bucket != "" {
		out.BucketRef = &refs.StorageBucketRef{
			External: in.Bucket,
		}
	}

	return out
}

func Convert_SQLInstance_API_v1_To_KRM_status(in *api.DatabaseInstance, out *krm.SQLInstanceStatus) error {
	if in == nil {
		return fmt.Errorf("cannot convert nil DatabaseInstance")
	}

	if in.AvailableMaintenanceVersions != nil {
		out.AvailableMaintenanceVersions = append(out.AvailableMaintenanceVersions, in.AvailableMaintenanceVersions...)
	}

	out.ConnectionName = LazyPtr(in.ConnectionName)

	out.DnsName = LazyPtr(in.DnsName)

	if len(in.IpAddresses) >= 1 {
		out.FirstIpAddress = LazyPtr(in.IpAddresses[0].IpAddress)
	}

	out.InstanceType = LazyPtr(in.InstanceType)

	if in.IpAddresses != nil {
		for _, ia := range in.IpAddresses {
			ipAddr := krm.InstanceIpAddressStatus{
				IpAddress:    LazyPtr(ia.IpAddress),
				TimeToRetire: LazyPtr(ia.TimeToRetire),
				Type:         LazyPtr(ia.Type),
			}
			out.IpAddress = append(out.IpAddress, ipAddr)

			if ia.Type == "PRIMARY" {
				out.PublicIpAddress = LazyPtr(ia.IpAddress)
			}

			if ia.Type == "PRIVATE" {
				out.PrivateIpAddress = LazyPtr(ia.IpAddress)
			}
		}
	}

	out.PscServiceAttachmentLink = LazyPtr(in.PscServiceAttachmentLink)

	out.SelfLink = LazyPtr(in.SelfLink)

	if in.ServerCaCert != nil {
		out.ServerCaCert = &krm.InstanceServerCaCertStatus{
			Cert:            LazyPtr(in.ServerCaCert.Cert),
			CommonName:      LazyPtr(in.ServerCaCert.CommonName),
			CreateTime:      LazyPtr(in.ServerCaCert.CreateTime),
			ExpirationTime:  LazyPtr(in.ServerCaCert.ExpirationTime),
			Sha1Fingerprint: LazyPtr(in.ServerCaCert.Sha1Fingerprint),
		}
	}

	out.ServiceAccountEmailAddress = LazyPtr(in.ServiceAccountEmailAddress)

	return nil
}

func SQLInstanceKRMToGCPCloneRequest(in *krm.SQLInstance, refs *SQLInstanceInternalRefs) (*api.InstancesCloneRequest, error) {
	if in == nil {
		return nil, fmt.Errorf("cannot convert nil SQLInstance")
	}

	if in.Spec.CloneSource == nil {
		// spec.cloneSource is required for converting KRM.SQLInstance -> api.InstancesCloneRequest.
		return nil, fmt.Errorf("cannot convert nil CloneSource")
	}

	cloneReq := &api.InstancesCloneRequest{
		CloneContext: &api.CloneContext{
			DatabaseNames: in.Spec.CloneSource.DatabaseNames,
			Kind:          "sql#cloneContext",
			PointInTime:   direct.ValueOf(in.Spec.CloneSource.PointInTime),
		},
	}

	resourceID := ValueOf(in.Spec.ResourceID)
	if resourceID == "" {
		resourceID = in.Name
	}
	cloneReq.CloneContext.DestinationInstanceName = resourceID

	if in.Spec.Settings.IpConfiguration != nil {
		cloneReq.CloneContext.AllocatedIpRange = ValueOf(in.Spec.Settings.IpConfiguration.AllocatedIpRange)
	}

	if in.Spec.CloneSource.BinLogCoordinates != nil {
		cloneReq.CloneContext.BinLogCoordinates = &api.BinLogCoordinates{
			BinLogFileName: in.Spec.CloneSource.BinLogCoordinates.BinLogFileName,
			BinLogPosition: in.Spec.CloneSource.BinLogCoordinates.BinLogPosition,
		}
	}

	if in.Spec.Settings.LocationPreference != nil {
		cloneReq.CloneContext.PreferredZone = ValueOf(in.Spec.Settings.LocationPreference.Zone)
	}

	return cloneReq, nil
}
