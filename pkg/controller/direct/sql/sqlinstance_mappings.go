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

	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"

	api "google.golang.org/api/sqladmin/v1beta4"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
)

func SQLInstanceKRMToGCP(in *krm.SQLInstance, actual *api.DatabaseInstance, unmanagedFields []string) (*api.DatabaseInstance, error) {
	if in == nil {
		return nil, fmt.Errorf("cannot convert nil KRM SQLInstance to GCP DatabaseInstance")
	}

	out := &api.DatabaseInstance{
		DatabaseVersion:             direct.ValueOf(in.Spec.DatabaseVersion),
		DiskEncryptionConfiguration: InstanceEncryptionKMSCryptoKeyRefKRMToGCP(in.Spec.EncryptionKMSCryptoKeyRef),
		// GeminiConfig is not supported in KRM API.
		InstanceType:       direct.ValueOf(in.Spec.InstanceType),
		Kind:               "sql#instance",
		MaintenanceVersion: direct.ValueOf(in.Spec.MaintenanceVersion),
		MasterInstanceName: InstanceMasterInstanceRefKRMToGCP(in.Spec.MasterInstanceRef),
		// MaxDiskSize is not supported in KRM API.
		Name: direct.ValueOf(in.Spec.ResourceID),
		// OnPremisesConfiguration is not supported in KRM API.
		Region:               direct.ValueOf(in.Spec.Region),
		ReplicaConfiguration: InstanceReplicaConfigurationKRMToGCP(in.Spec.ReplicaConfiguration),
		// ReplicationCluster is not supported in KRM API.
		RootPassword: InstanceRootPasswordKRMToGCP(in.Spec.RootPassword),
		Settings:     InstanceSettingsKRMToGCP(in.Spec.Settings, in.Labels),
		// SqlNetworkArchitecture is not supported in KRM API.
		// SwitchTransactionLogsToCloudStorageEnabled is not supported in KRM API.
	}

	// Here be dragons.
	ApplySQLInstanceGCPDefaults(in, out, actual, unmanagedFields)

	return out, nil
}

func InstanceEncryptionKMSCryptoKeyRefKRMToGCP(in *refs.KMSCryptoKeyRef) *api.DiskEncryptionConfiguration {
	if in == nil {
		return nil
	}

	out := &api.DiskEncryptionConfiguration{
		Kind:       "sql#diskEncryptionConfiguration",
		KmsKeyName: in.External,
	}

	return out
}

func InstanceMasterInstanceRefKRMToGCP(in *refs.SQLInstanceRef) string {
	if in == nil {
		return ""
	}

	out := in.External

	return out
}

func InstanceReplicaConfigurationKRMToGCP(in *krm.InstanceReplicaConfiguration) *api.ReplicaConfiguration {
	if in == nil {
		return nil
	}

	out := &api.ReplicaConfiguration{
		Kind: "sql#replicaConfiguration",
		// CascadableReplica is not supported in KRM API.
		FailoverTarget:            direct.ValueOf(in.FailoverTarget),
		MysqlReplicaConfiguration: InstanceMysqlReplicaConfigurationKRMToGCP(in),
	}

	if in.FailoverTarget != nil {
		out.ForceSendFields = append(out.ForceSendFields, "FailoverTarget")
	}

	return out
}

func InstanceRootPasswordKRMToGCP(in *refsv1beta1secret.Legacy) string {
	if in == nil {
		return ""
	}

	out := direct.ValueOf(in.Value)

	return out
}

func InstanceSettingsKRMToGCP(in krm.InstanceSettings, labels map[string]string) *api.Settings {
	out := &api.Settings{
		ActivationPolicy:            direct.ValueOf(in.ActivationPolicy),
		ActiveDirectoryConfig:       InstanceActiveDirectoryConfigKRMToGCP(in.ActiveDirectoryConfig),
		AdvancedMachineFeatures:     InstanceAdvancedMachineFeaturesKRMToGCP(in.AdvancedMachineFeatures),
		AuthorizedGaeApplications:   in.AuthorizedGaeApplications,
		AvailabilityType:            direct.ValueOf(in.AvailabilityType),
		BackupConfiguration:         InstanceBackupConfigurationKRMToGCP(in.BackupConfiguration),
		Collation:                   direct.ValueOf(in.Collation),
		ConnectorEnforcement:        direct.ValueOf(in.ConnectorEnforcement),
		CrashSafeReplicationEnabled: direct.ValueOf(in.CrashSafeReplication),
		DataCacheConfig:             InstanceDataCacheConfigKRMToGCP(in.DataCacheConfig),
		DataDiskSizeGb:              direct.ValueOf(in.DiskSize),
		DataDiskType:                direct.ValueOf(in.DiskType),
		DatabaseFlags:               InstanceDatabaseFlagsKRMToGCP(in.DatabaseFlags),
		// DatabaseReplicationEnabled is not supported in KRM API.
		DeletionProtectionEnabled: direct.ValueOf(in.DeletionProtectionEnabled),
		DenyMaintenancePeriods:    InstanceDenyMaintenancePeriodsKRMToGCP(in.DenyMaintenancePeriod),
		Edition:                   direct.ValueOf(in.Edition),
		// EnableDataplexIntegration is not supported in KRM API.
		// EnableGoogleMlIntegration is not supported in KRM API.
		InsightsConfig:           InstanceInsightsConfigKRMToGCP(in.InsightsConfig),
		IpConfiguration:          InstanceIpConfigurationKRMToGCP(in.IpConfiguration),
		Kind:                     "sql#settings",
		LocationPreference:       InstanceLocationPreferenceKRMToGCP(in.LocationPreference),
		MaintenanceWindow:        InstanceMaintenanceWindowKRMToGCP(in.MaintenanceWindow),
		PasswordValidationPolicy: InstancePasswordValidationPolicyKRMToGCP(in.PasswordValidationPolicy),
		PricingPlan:              direct.ValueOf(in.PricingPlan),
		ReplicationType:          direct.ValueOf(in.ReplicationType),
		// SettingsVersion is omitted because it is not part of the "desired state".
		SqlServerAuditConfig:   InstanceSqlServerAuditConfigKRMToGCP(in.SqlServerAuditConfig),
		StorageAutoResize:      in.DiskAutoresize,
		StorageAutoResizeLimit: direct.ValueOf(in.DiskAutoresizeLimit),
		Tier:                   in.Tier,
		TimeZone:               direct.ValueOf(in.TimeZone),
		UserLabels:             label.NewGCPLabelsFromK8sLabels(labels),
	}

	if in.CrashSafeReplication != nil {
		out.ForceSendFields = append(out.ForceSendFields, "CrashSafeReplicationEnabled")
	}
	if in.DeletionProtectionEnabled != nil {
		out.ForceSendFields = append(out.ForceSendFields, "DeletionProtectionEnabled")
	}
	if in.DiskAutoresize != nil {
		out.ForceSendFields = append(out.ForceSendFields, "StorageAutoResize")
	}
	if in.DiskAutoresizeLimit != nil {
		out.ForceSendFields = append(out.ForceSendFields, "StorageAutoResizeLimit")
	}

	return out
}

func InstanceMysqlReplicaConfigurationKRMToGCP(in *krm.InstanceReplicaConfiguration) *api.MySqlReplicaConfiguration {
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
		Password:                InstancePasswordKRMToGCP(in.Password),
		SslCipher:               direct.ValueOf(in.SslCipher),
		Username:                direct.ValueOf(in.Username),
		VerifyServerCertificate: direct.ValueOf(in.VerifyServerCertificate),
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

func InstancePasswordKRMToGCP(in *refsv1beta1secret.Legacy) string {
	if in == nil {
		return ""
	}

	out := direct.ValueOf(in.Value)

	return out
}

func InstanceActiveDirectoryConfigKRMToGCP(in *krm.InstanceActiveDirectoryConfig) *api.SqlActiveDirectoryConfig {
	if in == nil {
		return nil
	}

	out := &api.SqlActiveDirectoryConfig{
		Domain: in.Domain,
		Kind:   "sql#activeDirectoryConfig",
	}

	return out
}

func InstanceAdvancedMachineFeaturesKRMToGCP(in *krm.InstanceAdvancedMachineFeatures) *api.AdvancedMachineFeatures {
	if in == nil {
		return nil
	}

	out := &api.AdvancedMachineFeatures{
		ThreadsPerCore: direct.ValueOf(in.ThreadsPerCore),
	}

	return out
}

func InstanceBackupConfigurationKRMToGCP(in *krm.InstanceBackupConfiguration) *api.BackupConfiguration {
	if in == nil {
		return nil
	}

	out := &api.BackupConfiguration{
		BackupRetentionSettings:    InstanceBackupRetentionSettingsKRMToGCP(in.BackupRetentionSettings),
		BinaryLogEnabled:           direct.ValueOf(in.BinaryLogEnabled),
		Enabled:                    direct.ValueOf(in.Enabled),
		Kind:                       "sql#backupConfiguration",
		Location:                   direct.ValueOf(in.Location),
		PointInTimeRecoveryEnabled: direct.ValueOf(in.PointInTimeRecoveryEnabled),
		// ReplicationLogArchivingEnabled is not supported in KRM API.
		StartTime:                   direct.ValueOf(in.StartTime),
		TransactionLogRetentionDays: direct.ValueOf(in.TransactionLogRetentionDays),
		// TransactionalLogStorageState is not supported in KRM API.
	}

	if in.BinaryLogEnabled != nil {
		out.ForceSendFields = append(out.ForceSendFields, "BinaryLogEnabled")
	}
	if in.Enabled != nil {
		out.ForceSendFields = append(out.ForceSendFields, "Enabled")
	}
	if in.PointInTimeRecoveryEnabled != nil {
		out.ForceSendFields = append(out.ForceSendFields, "PointInTimeRecoveryEnabled")
	}
	if in.TransactionLogRetentionDays != nil {
		out.ForceSendFields = append(out.ForceSendFields, "TransactionLogRetentionDays")
	}

	return out
}

func InstanceBackupRetentionSettingsKRMToGCP(in *krm.InstanceBackupRetentionSettings) *api.BackupRetentionSettings {
	if in == nil {
		return nil
	}

	out := &api.BackupRetentionSettings{
		RetainedBackups: in.RetainedBackups,
		RetentionUnit:   direct.ValueOf(in.RetentionUnit),
	}

	if in.RetainedBackups == 0 {
		out.ForceSendFields = append(out.ForceSendFields, "RetainedBackups")
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

func InstanceDatabaseFlagsKRMToGCP(in []krm.InstanceDatabaseFlags) []*api.DatabaseFlags {
	out := []*api.DatabaseFlags{}
	for _, flag := range in {
		out = append(out, &api.DatabaseFlags{
			Name:  flag.Name,
			Value: flag.Value,
		})
	}
	return out
}

func InstanceDenyMaintenancePeriodsKRMToGCP(in *krm.InstanceDenyMaintenancePeriod) []*api.DenyMaintenancePeriod {
	if in == nil {
		return nil
	}

	// Note:  For some reason, the KRM API allows for only a single *InstanceDenyMaintenancePeriod. However,
	// in the GCP proto there is a list of []*api.DenyMaintenancePeriod. Though, in the GCP UI there is only
	// an option to specify a single deny maintenance period. For now, we'll only allow for specifying one.
	out := []*api.DenyMaintenancePeriod{
		{
			EndDate:   in.EndDate,
			StartDate: in.StartDate,
			Time:      in.Time,
		},
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

func InstanceIpConfigurationKRMToGCP(in *krm.InstanceIpConfiguration) *api.IpConfiguration {
	if in == nil {
		return nil
	}

	out := &api.IpConfiguration{
		AllocatedIpRange:                        direct.ValueOf(in.AllocatedIpRange),
		AuthorizedNetworks:                      InstanceAuthorizedNetworksKRMToGCP(in.AuthorizedNetworks),
		EnablePrivatePathForGoogleCloudServices: direct.ValueOf(in.EnablePrivatePathForGoogleCloudServices),
		Ipv4Enabled:                             direct.ValueOf(in.Ipv4Enabled),
		PrivateNetwork:                          InstancePrivateNetworkRefKRMToGCP(in.PrivateNetworkRef),
		PscConfig:                               InstancePscConfigKRMToGCP(in.PscConfig),
		RequireSsl:                              direct.ValueOf(in.RequireSsl),
		SslMode:                                 direct.ValueOf(in.SslMode),
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

func InstancePrivateNetworkRefKRMToGCP(in *refs.ComputeNetworkRef) string {
	if in == nil {
		return ""
	}

	out := in.External

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

func InstanceSqlServerAuditConfigKRMToGCP(in *krm.InstanceSqlServerAuditConfig) *api.SqlServerAuditConfig {
	if in == nil {
		return nil
	}

	out := &api.SqlServerAuditConfig{
		Bucket:            InstanceAuditConfigBucketRefKRMToGCP(in.BucketRef),
		Kind:              "sql#sqlServerAuditConfig",
		RetentionInterval: direct.ValueOf(in.RetentionInterval),
		UploadInterval:    direct.ValueOf(in.UploadInterval),
	}

	return out
}

func InstanceAuditConfigBucketRefKRMToGCP(in *storagev1beta1.StorageBucketRef) string {
	if in == nil {
		return ""
	}

	out := in.External

	return out
}

func SQLInstanceCloneKRMToGCP(in *krm.SQLInstance) (*api.InstancesCloneRequest, error) {
	if in == nil {
		return nil, fmt.Errorf("cannot convert nil KRM SQLInstance to GCP InstancesCloneRequest")
	}

	if in.Spec.CloneSource == nil {
		// If spec.cloneSource is not specified, it's invalid to convert krm.SQLInstance -> api.InstancesCloneRequest.
		// Instead, the krm.SQLInstance should be converted to an api.DatabaseInstance.
		return nil, fmt.Errorf("cannot convert SQLInstance to InstancesCloneRequest without CloneSource specified")
	}

	out := &api.InstancesCloneRequest{
		CloneContext: &api.CloneContext{
			AllocatedIpRange:        CloneAllocatedIpRangeKRMToGCP(in),
			BinLogCoordinates:       CloneBinLogCoordinatesKRMToGCP(in),
			DatabaseNames:           in.Spec.CloneSource.DatabaseNames,
			DestinationInstanceName: direct.ValueOf(in.Spec.ResourceID),
			Kind:                    "sql#cloneContext",
			PointInTime:             direct.ValueOf(in.Spec.CloneSource.PointInTime),
			PreferredSecondaryZone:  ClonePreferredSecondaryZoneKRMToGCP(in),
			PreferredZone:           ClonePreferredZoneKRMToGCP(in),
		},
	}

	return out, nil
}

func CloneAllocatedIpRangeKRMToGCP(in *krm.SQLInstance) string {
	if in.Spec.Settings.IpConfiguration == nil {
		return ""
	}

	out := direct.ValueOf(in.Spec.Settings.IpConfiguration.AllocatedIpRange)

	return out
}

func CloneBinLogCoordinatesKRMToGCP(in *krm.SQLInstance) *api.BinLogCoordinates {
	if in.Spec.CloneSource.BinLogCoordinates == nil {
		return nil
	}

	out := &api.BinLogCoordinates{
		BinLogFileName: in.Spec.CloneSource.BinLogCoordinates.BinLogFileName,
		BinLogPosition: in.Spec.CloneSource.BinLogCoordinates.BinLogPosition,
	}

	return out
}

func ClonePreferredZoneKRMToGCP(in *krm.SQLInstance) string {
	if in.Spec.Settings.LocationPreference == nil {
		return ""
	}

	out := direct.ValueOf(in.Spec.Settings.LocationPreference.Zone)

	return out
}

func ClonePreferredSecondaryZoneKRMToGCP(in *krm.SQLInstance) string {
	if in.Spec.Settings.LocationPreference == nil {
		return ""
	}

	out := direct.ValueOf(in.Spec.Settings.LocationPreference.SecondaryZone)

	return out
}

func SQLInstanceGCPToKRM(in *api.DatabaseInstance) (*krm.SQLInstance, error) {
	if in == nil {
		return nil, fmt.Errorf("cannot convert nil GCP DatabaseInstance to KRM SQLInstance")
	}

	out := &krm.SQLInstance{
		ObjectMeta: metav1.ObjectMeta{
			Labels: in.Settings.UserLabels,
		},
		Spec: krm.SQLInstanceSpec{
			DatabaseVersion:           direct.LazyPtr(in.DatabaseVersion),
			EncryptionKMSCryptoKeyRef: InstanceEncryptionKMSCryptoKeyRefGCPToKRM(in.DiskEncryptionConfiguration),
			// GeminiConfig is not supported in KRM API.
			InstanceType:       direct.LazyPtr(in.InstanceType),
			MaintenanceVersion: direct.LazyPtr(in.MaintenanceVersion),
			MasterInstanceRef:  InstanceMasterInstanceRefGCPToKRM(in.MasterInstanceName),
			// MaxDiskSize is not supported in KRM API.
			ResourceID: direct.LazyPtr(in.Name),
			// OnPremisesConfiguration is not supported in KRM API.
			Region:               direct.LazyPtr(in.Region),
			ReplicaConfiguration: InstanceReplicaConfigurationGCPToKRM(in.ReplicaConfiguration),
			// ReplicationCluster is not supported in KRM API.
			// RootPassword is not exported.
			Settings: InstanceSettingsGCPToKRM(in.Settings),
			// SqlNetworkArchitecture is not supported in KRM API.
			// SwitchTransactionLogsToCloudStorageEnabled is not supported in KRM API.
		},
	}

	return out, nil
}

func InstanceEncryptionKMSCryptoKeyRefGCPToKRM(in *api.DiskEncryptionConfiguration) *refs.KMSCryptoKeyRef {
	if in == nil {
		return nil
	}

	out := &refs.KMSCryptoKeyRef{
		External: in.KmsKeyName,
	}

	return out
}

func InstanceMasterInstanceRefGCPToKRM(in string) *refs.SQLInstanceRef {
	if in == "" {
		return nil
	}

	out := &refs.SQLInstanceRef{
		External: in,
	}

	return out
}

func InstanceSettingsGCPToKRM(in *api.Settings) krm.InstanceSettings {
	if in == nil {
		return krm.InstanceSettings{}
	}

	out := krm.InstanceSettings{
		ActivationPolicy:          direct.LazyPtr(in.ActivationPolicy),
		ActiveDirectoryConfig:     InstanceActiveDirectoryConfigGCPToKRM(in.ActiveDirectoryConfig),
		AdvancedMachineFeatures:   InstanceAdvancedMachineFeaturesGCPToKRM(in.AdvancedMachineFeatures),
		AuthorizedGaeApplications: in.AuthorizedGaeApplications,
		AvailabilityType:          direct.LazyPtr(in.AvailabilityType),
		BackupConfiguration:       InstanceBackupConfigurationGCPToKRM(in.BackupConfiguration),
		Collation:                 direct.LazyPtr(in.Collation),
		ConnectorEnforcement:      direct.LazyPtr(in.ConnectorEnforcement),
		CrashSafeReplication:      direct.PtrTo(in.CrashSafeReplicationEnabled),
		DataCacheConfig:           InstanceDataCacheConfigGCPToKRM(in.DataCacheConfig),
		DiskSize:                  direct.LazyPtr(in.DataDiskSizeGb),
		DiskType:                  direct.LazyPtr(in.DataDiskType),
		DatabaseFlags:             InstanceDatabaseFlagsGCPToKRM(in.DatabaseFlags),
		// DatabaseReplicationEnabled is not supported in KRM API.
		DeletionProtectionEnabled: direct.PtrTo(in.DeletionProtectionEnabled),
		DenyMaintenancePeriod:     InstanceDenyMaintenancePeriodsGCPToKRM(in.DenyMaintenancePeriods),
		Edition:                   direct.LazyPtr(in.Edition),
		// EnableDataplexIntegration is not supported in KRM API.
		// EnableGoogleMlIntegration is not supported in KRM API.
		InsightsConfig:           InstanceInsightsConfigGCPToKRM(in.InsightsConfig),
		IpConfiguration:          InstanceIpConfigurationGCPToKRM(in.IpConfiguration),
		LocationPreference:       InstanceLocationPreferenceGCPToKRM(in.LocationPreference),
		MaintenanceWindow:        InstanceMaintenanceWindowGCPToKRM(in.MaintenanceWindow),
		PasswordValidationPolicy: InstancePasswordValidationPolicyGCPToKRM(in.PasswordValidationPolicy),
		PricingPlan:              direct.LazyPtr(in.PricingPlan),
		ReplicationType:          direct.LazyPtr(in.ReplicationType),
		// SettingsVersion is omitted because it is not part of the "desired state".
		SqlServerAuditConfig: InstanceSqlServerAuditConfigGCPToKRM(in.SqlServerAuditConfig),
		DiskAutoresize:       in.StorageAutoResize,
		DiskAutoresizeLimit:  direct.PtrTo(in.StorageAutoResizeLimit),
		Tier:                 in.Tier,
		TimeZone:             direct.LazyPtr(in.TimeZone),
	}

	return out
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
		CaCertificate:         mrc.CaCertificate,
		ClientCertificate:     mrc.ClientCertificate,
		ClientKey:             mrc.ClientKey,
		ConnectRetryInterval:  mrc.ConnectRetryInterval,
		DumpFilePath:          mrc.DumpFilePath,
		FailoverTarget:        irc.FailoverTarget,
		MasterHeartbeatPeriod: mrc.MasterHeartbeatPeriod,
		// Password is not exported.
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
		CaCertificate:         direct.LazyPtr(in.CaCertificate),
		ClientCertificate:     direct.LazyPtr(in.ClientCertificate),
		ClientKey:             direct.LazyPtr(in.ClientKey),
		ConnectRetryInterval:  direct.PtrTo(in.ConnectRetryInterval),
		DumpFilePath:          direct.LazyPtr(in.DumpFilePath),
		MasterHeartbeatPeriod: direct.PtrTo(in.MasterHeartbeatPeriod),
		// Password is not exported.
		SslCipher:               direct.LazyPtr(in.SslCipher),
		Username:                direct.LazyPtr(in.Username),
		VerifyServerCertificate: direct.PtrTo(in.VerifyServerCertificate),
	}

	return out
}

func InstanceActiveDirectoryConfigGCPToKRM(in *api.SqlActiveDirectoryConfig) *krm.InstanceActiveDirectoryConfig {
	if in == nil {
		return nil
	}

	out := &krm.InstanceActiveDirectoryConfig{
		Domain: in.Domain,
	}

	return out
}

func InstanceAdvancedMachineFeaturesGCPToKRM(in *api.AdvancedMachineFeatures) *krm.InstanceAdvancedMachineFeatures {
	if in == nil {
		return nil
	}

	out := &krm.InstanceAdvancedMachineFeatures{
		ThreadsPerCore: direct.LazyPtr(in.ThreadsPerCore),
	}

	return out
}

func InstanceBackupConfigurationGCPToKRM(in *api.BackupConfiguration) *krm.InstanceBackupConfiguration {
	if in == nil {
		return nil
	}

	out := &krm.InstanceBackupConfiguration{
		BackupRetentionSettings:    InstanceBackupRetentionSettingsGCPToKRM(in.BackupRetentionSettings),
		BinaryLogEnabled:           direct.PtrTo(in.BinaryLogEnabled),
		Enabled:                    direct.PtrTo(in.Enabled),
		Location:                   direct.LazyPtr(in.Location),
		PointInTimeRecoveryEnabled: direct.PtrTo(in.PointInTimeRecoveryEnabled),
		// ReplicationLogArchivingEnabled is not supported in KRM API.
		StartTime:                   direct.LazyPtr(in.StartTime),
		TransactionLogRetentionDays: direct.PtrTo(in.TransactionLogRetentionDays),
		// TransactionalLogStorageState is not supported in KRM API.
	}

	return out
}
func InstanceBackupRetentionSettingsGCPToKRM(in *api.BackupRetentionSettings) *krm.InstanceBackupRetentionSettings {
	if in == nil {
		return nil
	}

	out := &krm.InstanceBackupRetentionSettings{
		RetainedBackups: in.RetainedBackups,
		RetentionUnit:   direct.LazyPtr(in.RetentionUnit),
	}

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

func InstanceDatabaseFlagsGCPToKRM(in []*api.DatabaseFlags) []krm.InstanceDatabaseFlags {
	out := []krm.InstanceDatabaseFlags{}
	for _, flag := range in {
		out = append(out, krm.InstanceDatabaseFlags{
			Name:  flag.Name,
			Value: flag.Value,
		})
	}
	return out
}

func InstanceDenyMaintenancePeriodsGCPToKRM(in []*api.DenyMaintenancePeriod) *krm.InstanceDenyMaintenancePeriod {
	if in == nil || len(in) < 1 {
		return nil
	}

	// Note:  For some reason, the KRM API allows for only a single *InstanceDenyMaintenancePeriod. However,
	// in the GCP proto there is a list of []*api.DenyMaintenancePeriod. Though, in the GCP UI there is only
	// an option to specify a single deny maintenance period. For now, we'll only allow for specifying one.
	out := &krm.InstanceDenyMaintenancePeriod{
		EndDate:   in[0].EndDate,
		StartDate: in[0].StartDate,
		Time:      in[0].Time,
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
		PrivateNetworkRef:                       InstancePrivateNetworkRefRefGCPToKRM(in.PrivateNetwork),
		PscConfig:                               InstancePscConfigGCPToKRM(in.PscConfig),
		RequireSsl:                              direct.PtrTo(in.RequireSsl),
		SslMode:                                 direct.LazyPtr(in.SslMode),
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

func InstancePrivateNetworkRefRefGCPToKRM(in string) *refs.ComputeNetworkRef {
	if in == "" {
		return nil
	}

	out := &refs.ComputeNetworkRef{
		External: in,
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
		BucketRef:         InstanceAuditConfigBucketRefGCPToKRM(in.Bucket),
		RetentionInterval: direct.LazyPtr(in.RetentionInterval),
		UploadInterval:    direct.LazyPtr(in.UploadInterval),
	}

	return out
}

func InstanceAuditConfigBucketRefGCPToKRM(in string) *storagev1beta1.StorageBucketRef {
	if in == "" {
		return nil
	}

	out := &storagev1beta1.StorageBucketRef{
		External: in,
	}

	return out
}

func SQLInstanceStatusGCPToKRM(in *api.DatabaseInstance) (*krm.SQLInstanceStatus, error) {
	if in == nil {
		return nil, fmt.Errorf("cannot convert nil DatabaseInstance")
	}

	out := &krm.SQLInstanceStatus{
		AvailableMaintenanceVersions: in.AvailableMaintenanceVersions,
		ConnectionName:               direct.LazyPtr(in.ConnectionName),
		DnsName:                      direct.LazyPtr(in.DnsName),
		FirstIpAddress:               SQLInstanceFirstIpAddressGCPToKRM(in.IpAddresses),
		InstanceType:                 direct.LazyPtr(in.InstanceType),
		IpAddress:                    SQLInstanceIpAddressesGCPToKRM(in.IpAddresses),
		PrivateIpAddress:             SQLInstancePrivateIpAddressGCPToKRM(in.IpAddresses),
		PscServiceAttachmentLink:     direct.LazyPtr(in.PscServiceAttachmentLink),
		PublicIpAddress:              SQLInstancePublicIpAddressGCPToKRM(in.IpAddresses),
		SelfLink:                     direct.LazyPtr(in.SelfLink),
		ServerCaCert:                 SQLInstanceServerCaCertGCPToKRM(in.ServerCaCert),
		ServiceAccountEmailAddress:   direct.LazyPtr(in.ServiceAccountEmailAddress),
	}

	return out, nil
}

func SQLInstanceFirstIpAddressGCPToKRM(in []*api.IpMapping) *string {
	if len(in) == 0 {
		return nil
	}

	return direct.LazyPtr(in[0].IpAddress)
}

func SQLInstanceIpAddressesGCPToKRM(in []*api.IpMapping) []krm.InstanceIpAddressStatus {
	if in == nil {
		return nil
	}

	var out []krm.InstanceIpAddressStatus
	for _, ia := range in {
		ipAddr := krm.InstanceIpAddressStatus{
			IpAddress:    direct.LazyPtr(ia.IpAddress),
			TimeToRetire: direct.LazyPtr(ia.TimeToRetire),
			Type:         direct.LazyPtr(ia.Type),
		}
		out = append(out, ipAddr)
	}

	return out
}

func SQLInstancePublicIpAddressGCPToKRM(in []*api.IpMapping) *string {
	if in == nil {
		return nil
	}

	for _, ia := range in {
		if ia.Type == "PRIMARY" {
			return direct.LazyPtr(ia.IpAddress)
		}
	}

	return nil
}

func SQLInstancePrivateIpAddressGCPToKRM(in []*api.IpMapping) *string {
	if in == nil {
		return nil
	}

	for _, ia := range in {
		if ia.Type == "PRIVATE" {
			return direct.LazyPtr(ia.IpAddress)
		}
	}

	return nil
}

func SQLInstanceServerCaCertGCPToKRM(in *api.SslCert) *krm.InstanceServerCaCertStatus {
	if in == nil {
		return nil
	}

	out := &krm.InstanceServerCaCertStatus{
		Cert:            direct.LazyPtr(in.Cert),
		CommonName:      direct.LazyPtr(in.CommonName),
		CreateTime:      direct.LazyPtr(in.CreateTime),
		ExpirationTime:  direct.LazyPtr(in.ExpirationTime),
		Sha1Fingerprint: direct.LazyPtr(in.Sha1Fingerprint),
	}

	return out
}
