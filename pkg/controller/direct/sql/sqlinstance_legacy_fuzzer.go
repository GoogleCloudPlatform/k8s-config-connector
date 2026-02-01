// Copyright 2026 Google LLC
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
	api "google.golang.org/api/sqladmin/v1beta4"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer_NoProto(sqlInstanceLegacyFuzzer())
}

func sqlInstanceLegacyFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto[*api.DatabaseInstance, krm.SQLInstance, krm.SQLInstanceStatus](&api.DatabaseInstance{},
		func(ctx *direct.MapContext, in *api.DatabaseInstance) *krm.SQLInstance {
			res, err := SQLInstanceGCPToKRM(in)
			if err != nil {
				ctx.Errorf("%v", err)
			}
			return res
		},
		func(ctx *direct.MapContext, in *krm.SQLInstance) *api.DatabaseInstance {
			res, err := SQLInstanceKRMToGCP(in, nil, nil)
			if err != nil {
				ctx.Errorf("%v", err)
			}
			return res
		},
		func(ctx *direct.MapContext, in *api.DatabaseInstance) *krm.SQLInstanceStatus {
			res, err := SQLInstanceStatusGCPToKRM(in)
			if err != nil {
				ctx.Errorf("%v", err)
			}
			return res
		},
		func(ctx *direct.MapContext, in *krm.SQLInstanceStatus) *api.DatabaseInstance {
			return &api.DatabaseInstance{}
		},
	)

	f.FilterSpec = func(in *api.DatabaseInstance) {
		in.Kind = "sql#instance"
		if in.DiskEncryptionConfiguration != nil {
			in.DiskEncryptionConfiguration.Kind = "sql#diskEncryptionConfiguration"
		}
		if in.ReplicaConfiguration != nil {
			in.ReplicaConfiguration.Kind = "sql#replicaConfiguration"
			if in.ReplicaConfiguration.MysqlReplicaConfiguration != nil {
				in.ReplicaConfiguration.MysqlReplicaConfiguration.Kind = "sql#mysqlReplicaConfiguration"
			}
		}
		if in.Settings != nil {
			in.Settings.Kind = "sql#settings"
			if in.Settings.ActiveDirectoryConfig != nil {
				in.Settings.ActiveDirectoryConfig.Kind = "sql#activeDirectoryConfig"
			}
			if in.Settings.BackupConfiguration != nil {
				in.Settings.BackupConfiguration.Kind = "sql#backupConfiguration"
			}
			if in.Settings.IpConfiguration != nil {
				for _, net := range in.Settings.IpConfiguration.AuthorizedNetworks {
					net.Kind = "sql#aclEntry"
				}
			}
			if in.Settings.LocationPreference != nil {
				in.Settings.LocationPreference.Kind = "sql#locationPreference"
			}
			if in.Settings.MaintenanceWindow != nil {
				in.Settings.MaintenanceWindow.Kind = "sql#maintenanceWindow"
			}
			if in.Settings.SqlServerAuditConfig != nil {
				in.Settings.SqlServerAuditConfig.Kind = "sql#sqlServerAuditConfig"
			}
			if in.Settings.UserLabels == nil {
				in.Settings.UserLabels = make(map[string]string)
			}
			in.Settings.UserLabels["managed-by-cnrm"] = "true"
		}
	}

	f.SpecField(".DatabaseVersion")
	f.SpecField(".DiskEncryptionConfiguration")
	f.SpecField(".DiskEncryptionConfiguration.KmsKeyName")
	f.SpecField(".GeminiConfig")
	f.SpecField(".GeminiConfig.ActiveQueryEnabled")
	f.SpecField(".GeminiConfig.Entitled")
	f.SpecField(".GeminiConfig.FlagRecommenderEnabled")
	f.SpecField(".GeminiConfig.GoogleVacuumMgmtEnabled")
	f.SpecField(".GeminiConfig.IndexAdvisorEnabled")
	f.SpecField(".GeminiConfig.OomSessionCancelEnabled")
	f.SpecField(".IncludeReplicasForMajorVersionUpgrade")
	f.SpecField(".InstanceType")
	f.SpecField(".MaintenanceVersion")
	f.SpecField(".MasterInstanceName")
	f.SpecField(".Name")
	f.SpecField(".Region")
	f.SpecField(".ReplicaConfiguration")
	f.SpecField(".ReplicaConfiguration.FailoverTarget")
	f.SpecField(".ReplicaConfiguration.MysqlReplicaConfiguration")
	f.SpecField(".ReplicaConfiguration.MysqlReplicaConfiguration.CaCertificate")
	f.SpecField(".ReplicaConfiguration.MysqlReplicaConfiguration.ClientCertificate")
	f.SpecField(".ReplicaConfiguration.MysqlReplicaConfiguration.ClientKey")
	f.SpecField(".ReplicaConfiguration.MysqlReplicaConfiguration.ConnectRetryInterval")
	f.SpecField(".ReplicaConfiguration.MysqlReplicaConfiguration.DumpFilePath")
	f.SpecField(".ReplicaConfiguration.MysqlReplicaConfiguration.MasterHeartbeatPeriod")
	f.SpecField(".ReplicaConfiguration.MysqlReplicaConfiguration.SslCipher")
	f.SpecField(".ReplicaConfiguration.MysqlReplicaConfiguration.Username")
	f.SpecField(".ReplicaConfiguration.MysqlReplicaConfiguration.VerifyServerCertificate")
	f.SpecField(".ReplicationCluster")
	f.SpecField(".ReplicationCluster.FailoverDrReplicaName")
	f.SpecField(".Settings")
	f.SpecField(".Settings.ActivationPolicy")
	f.SpecField(".Settings.ActiveDirectoryConfig")
	f.SpecField(".Settings.ActiveDirectoryConfig.Domain")
	f.SpecField(".Settings.AdvancedMachineFeatures")
	f.SpecField(".Settings.AdvancedMachineFeatures.ThreadsPerCore")
	f.SpecField(".Settings.AuthorizedGaeApplications")
	f.SpecField(".Settings.AvailabilityType")
	f.SpecField(".Settings.BackupConfiguration")
	f.SpecField(".Settings.BackupConfiguration.BackupRetentionSettings")
	f.SpecField(".Settings.BackupConfiguration.BackupRetentionSettings.RetainedBackups")
	f.SpecField(".Settings.BackupConfiguration.BackupRetentionSettings.RetentionUnit")
	f.SpecField(".Settings.BackupConfiguration.BinaryLogEnabled")
	f.SpecField(".Settings.BackupConfiguration.Enabled")
	f.SpecField(".Settings.BackupConfiguration.Location")
	f.SpecField(".Settings.BackupConfiguration.PointInTimeRecoveryEnabled")
	f.SpecField(".Settings.BackupConfiguration.StartTime")
	f.SpecField(".Settings.BackupConfiguration.TransactionLogRetentionDays")
	f.SpecField(".Settings.Collation")
	f.SpecField(".Settings.ConnectorEnforcement")
	f.SpecField(".Settings.CrashSafeReplicationEnabled")
	f.SpecField(".Settings.DataCacheConfig")
	f.SpecField(".Settings.DataCacheConfig.DataCacheEnabled")
	f.SpecField(".Settings.DataDiskSizeGb")
	f.SpecField(".Settings.DataDiskType")
	f.SpecField(".Settings.DatabaseFlags")
	f.SpecField(".Settings.DeletionProtectionEnabled")
	f.SpecField(".Settings.DenyMaintenancePeriods")
	f.SpecField(".Settings.Edition")
	f.SpecField(".Settings.InsightsConfig")
	f.SpecField(".Settings.InsightsConfig.QueryInsightsEnabled")
	f.SpecField(".Settings.InsightsConfig.QueryPlansPerMinute")
	f.SpecField(".Settings.InsightsConfig.QueryStringLength")
	f.SpecField(".Settings.InsightsConfig.RecordApplicationTags")
	f.SpecField(".Settings.InsightsConfig.RecordClientAddress")
	f.SpecField(".Settings.IpConfiguration")
	f.SpecField(".Settings.IpConfiguration.AllocatedIpRange")
	f.SpecField(".Settings.IpConfiguration.AuthorizedNetworks")
	f.SpecField(".Settings.IpConfiguration.EnablePrivatePathForGoogleCloudServices")
	f.SpecField(".Settings.IpConfiguration.Ipv4Enabled")
	f.SpecField(".Settings.IpConfiguration.PrivateNetwork")
	f.SpecField(".Settings.IpConfiguration.PscConfig")
	f.SpecField(".Settings.IpConfiguration.PscConfig.AllowedConsumerProjects")
	f.SpecField(".Settings.IpConfiguration.PscConfig.PscEnabled")
	f.SpecField(".Settings.IpConfiguration.RequireSsl")
	f.SpecField(".Settings.IpConfiguration.SslMode")
	f.SpecField(".Settings.LocationPreference")
	f.SpecField(".Settings.LocationPreference.FollowGaeApplication")
	f.SpecField(".Settings.LocationPreference.SecondaryZone")
	f.SpecField(".Settings.LocationPreference.Zone")
	f.SpecField(".Settings.MaintenanceWindow")
	f.SpecField(".Settings.MaintenanceWindow.Day")
	f.SpecField(".Settings.MaintenanceWindow.Hour")
	f.SpecField(".Settings.MaintenanceWindow.UpdateTrack")
	f.SpecField(".Settings.PasswordValidationPolicy")
	f.SpecField(".Settings.PasswordValidationPolicy.Complexity")
	f.SpecField(".Settings.PasswordValidationPolicy.DisallowUsernameSubstring")
	f.SpecField(".Settings.PasswordValidationPolicy.EnablePasswordPolicy")
	f.SpecField(".Settings.PasswordValidationPolicy.MinLength")
	f.SpecField(".Settings.PasswordValidationPolicy.PasswordChangeInterval")
	f.SpecField(".Settings.PasswordValidationPolicy.ReuseInterval")
	f.SpecField(".Settings.PricingPlan")
	f.SpecField(".Settings.ReplicationType")
	f.SpecField(".Settings.SqlServerAuditConfig")
	f.SpecField(".Settings.SqlServerAuditConfig.Bucket")
	f.SpecField(".Settings.SqlServerAuditConfig.RetentionInterval")
	f.SpecField(".Settings.SqlServerAuditConfig.UploadInterval")
	f.SpecField(".Settings.StorageAutoResize")
	f.SpecField(".Settings.StorageAutoResizeLimit")
	f.SpecField(".Settings.Tier")
	f.SpecField(".Settings.TimeZone")
	f.SpecField(".Settings.UserLabels")

	f.StatusField(".AvailableMaintenanceVersions")
	f.StatusField(".ConnectionName")
	f.StatusField(".CreateTime")
	f.StatusField(".DatabaseInstalledVersion")
	f.StatusField(".DiskEncryptionStatus")
	f.StatusField(".DnsName")
	f.StatusField(".GceZone")
	f.StatusField(".IpAddresses")
	f.StatusField(".Ipv6Address")
	f.StatusField(".OutOfDiskReport")
	f.StatusField(".PscServiceAttachmentLink")
	f.StatusField(".ReplicaNames")
	f.StatusField(".SatisfiesPzs")
	f.StatusField(".ScheduledMaintenance")
	f.StatusField(".SecondaryGceZone")
	f.StatusField(".SelfLink")
	f.StatusField(".ServerCaCert")
	f.StatusField(".ServiceAccountEmailAddress")
	f.StatusField(".State")
	f.StatusField(".SuspensionReason")
	f.StatusField(".UpgradableDatabaseVersions")
	f.StatusField(".WriteEndpoint")
	f.StatusField(".PrimaryDnsName")
	f.StatusField(".ReplicationCluster.DrReplica")
	f.StatusField(".ReplicationCluster.PsaWriteEndpoint")

	f.Unimplemented_NotYetTriaged(".BackendType")
	f.Unimplemented_NotYetTriaged(".CurrentDiskSize")
	f.Unimplemented_NotYetTriaged(".DnsNames")
	f.Unimplemented_NotYetTriaged(".Etag")
	f.Unimplemented_NotYetTriaged(".FailoverReplica")
	f.Unimplemented_NotYetTriaged(".Kind")
	f.Unimplemented_NotYetTriaged(".MaxDiskSize")
	f.Unimplemented_NotYetTriaged(".NodeCount")
	f.Unimplemented_NotYetTriaged(".Nodes")
	f.Unimplemented_NotYetTriaged(".OnPremisesConfiguration")
	f.Unimplemented_NotYetTriaged(".Project")
	f.Unimplemented_NotYetTriaged(".RootPassword")
	f.Unimplemented_NotYetTriaged(".IncludeReplicasForMajorVersionUpgrade")
	f.Unimplemented_NotYetTriaged(".ServerResponse")
	f.Unimplemented_NotYetTriaged(".SatisfiesPzi")
	f.Unimplemented_NotYetTriaged(".SqlNetworkArchitecture")
	f.Unimplemented_NotYetTriaged(".SwitchTransactionLogsToCloudStorageEnabled")
	f.Unimplemented_NotYetTriaged(".Tags")

	f.Unimplemented_NotYetTriaged(".Settings.DatabaseReplicationEnabled")
	f.Unimplemented_NotYetTriaged(".Settings.EnableDataplexIntegration")
	f.Unimplemented_NotYetTriaged(".Settings.EnableGoogleMlIntegration")
	f.Unimplemented_NotYetTriaged(".Settings.Kind")
	f.Unimplemented_NotYetTriaged(".Settings.SettingsVersion")
	f.Unimplemented_NotYetTriaged(".Settings.AutoUpgradeEnabled")
	f.Unimplemented_NotYetTriaged(".Settings.ConnectionPoolConfig")
	f.Unimplemented_NotYetTriaged(".Settings.DataApiAccess")
	f.Unimplemented_NotYetTriaged(".Settings.DataDiskProvisionedIops")
	f.Unimplemented_NotYetTriaged(".Settings.DataDiskProvisionedThroughput")
	f.Unimplemented_NotYetTriaged(".Settings.EntraidConfig")
	f.Unimplemented_NotYetTriaged(".Settings.FinalBackupConfig")
	f.Unimplemented_NotYetTriaged(".Settings.PerformanceCaptureConfig")
	f.Unimplemented_NotYetTriaged(".Settings.ReadPoolAutoScaleConfig")
	f.Unimplemented_NotYetTriaged(".Settings.ReplicationLagMaxSeconds")
	f.Unimplemented_NotYetTriaged(".Settings.ServerCaMode")
	f.Unimplemented_NotYetTriaged(".Settings.ServerCaPool")
	f.Unimplemented_NotYetTriaged(".Settings.ServerCertificateRotationMode")
	f.Unimplemented_NotYetTriaged(".Settings.RetainBackupsOnDelete")

	f.Unimplemented_NotYetTriaged(".Settings.ActiveDirectoryConfig.AdminCredentialSecretName")
	f.Unimplemented_NotYetTriaged(".Settings.ActiveDirectoryConfig.DnsServers")
	f.Unimplemented_NotYetTriaged(".Settings.ActiveDirectoryConfig.Kind")
	f.Unimplemented_NotYetTriaged(".Settings.ActiveDirectoryConfig.Mode")
	f.Unimplemented_NotYetTriaged(".Settings.ActiveDirectoryConfig.OrganizationalUnit")

	f.Unimplemented_NotYetTriaged(".Settings.BackupConfiguration.BackupTier")
	f.Unimplemented_NotYetTriaged(".Settings.BackupConfiguration.TransactionalLogStorageState")
	f.Unimplemented_NotYetTriaged(".Settings.BackupConfiguration.Kind")

	f.Unimplemented_NotYetTriaged(".Settings.IpConfiguration.CustomSubjectAlternativeNames")
	f.Unimplemented_NotYetTriaged(".Settings.IpConfiguration.ServerCaMode")
	f.Unimplemented_NotYetTriaged(".Settings.IpConfiguration.ServerCaPool")
	f.Unimplemented_NotYetTriaged(".Settings.IpConfiguration.ServerCertificateRotationMode")
	f.Unimplemented_NotYetTriaged(".Settings.IpConfiguration.PscConfig.NetworkAttachmentUri")
	f.Unimplemented_NotYetTriaged(".Settings.IpConfiguration.PscConfig.PscAutoConnections")

	f.Unimplemented_NotYetTriaged(".Settings.LocationPreference.Kind")
	f.Unimplemented_NotYetTriaged(".Settings.MaintenanceWindow.Kind")
	f.Unimplemented_NotYetTriaged(".Settings.SqlServerAuditConfig.Kind")
	f.Unimplemented_NotYetTriaged(".ReplicaConfiguration.Kind")
	f.Unimplemented_NotYetTriaged(".ReplicaConfiguration.MysqlReplicaConfiguration.Kind")
	f.Unimplemented_NotYetTriaged(".DiskEncryptionConfiguration.Kind")

	// Global ignores for ForceSendFields and NullFields which are internal to the api library
	f.Unimplemented_NotYetTriaged(".ForceSendFields")
	f.Unimplemented_NotYetTriaged(".NullFields")
	f.Unimplemented_NotYetTriaged(".DiskEncryptionConfiguration.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".DiskEncryptionConfiguration.NullFields")
	f.Unimplemented_NotYetTriaged(".ReplicaConfiguration.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".ReplicaConfiguration.NullFields")
	f.Unimplemented_NotYetTriaged(".ReplicaConfiguration.MysqlReplicaConfiguration.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".ReplicaConfiguration.MysqlReplicaConfiguration.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Settings.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.ActiveDirectoryConfig.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Settings.ActiveDirectoryConfig.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.AdvancedMachineFeatures.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Settings.AdvancedMachineFeatures.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.BackupConfiguration.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Settings.BackupConfiguration.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.BackupConfiguration.BackupRetentionSettings.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Settings.BackupConfiguration.BackupRetentionSettings.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.DataCacheConfig.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Settings.DataCacheConfig.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.InsightsConfig.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Settings.InsightsConfig.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.IpConfiguration.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Settings.IpConfiguration.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.IpConfiguration.PscConfig.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Settings.IpConfiguration.PscConfig.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.LocationPreference.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Settings.LocationPreference.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.MaintenanceWindow.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Settings.MaintenanceWindow.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.PasswordValidationPolicy.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Settings.PasswordValidationPolicy.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.SqlServerAuditConfig.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Settings.SqlServerAuditConfig.NullFields")
	f.Unimplemented_NotYetTriaged(".GeminiConfig.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".GeminiConfig.NullFields")
	f.Unimplemented_NotYetTriaged(".ReplicationCluster.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".ReplicationCluster.NullFields")
	f.Unimplemented_NotYetTriaged(".Settings.DenyMaintenancePeriods")

	return f
}
