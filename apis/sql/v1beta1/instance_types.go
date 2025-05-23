// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InstanceActiveDirectoryConfig struct {
	/* Domain name of the Active Directory for SQL Server (e.g., mydomain.com). */
	Domain string `json:"domain"`
}

type InstanceAdvancedMachineFeatures struct {
	/* The number of threads per physical core. Can be 1 or 2. */
	// +optional
	ThreadsPerCore *int64 `json:"threadsPerCore,omitempty"`
}

type InstanceAuthorizedNetworks struct {
	// +optional
	ExpirationTime *string `json:"expirationTime,omitempty"`

	// +optional
	Name *string `json:"name,omitempty"`

	Value string `json:"value"`
}

type InstanceBackupConfiguration struct {
	// +optional
	BackupRetentionSettings *InstanceBackupRetentionSettings `json:"backupRetentionSettings,omitempty"`

	/* True if binary logging is enabled. If settings.backup_configuration.enabled is false, this must be as well. Can only be used with MySQL. */
	// +optional
	BinaryLogEnabled *bool `json:"binaryLogEnabled,omitempty"`

	/* True if backup configuration is enabled. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	/* Location of the backup configuration. */
	// +optional
	Location *string `json:"location,omitempty"`

	/* True if Point-in-time recovery is enabled. */
	// +optional
	PointInTimeRecoveryEnabled *bool `json:"pointInTimeRecoveryEnabled,omitempty"`

	/* HH:MM format time indicating when backup configuration starts. */
	// +optional
	StartTime *string `json:"startTime,omitempty"`

	/* The number of days of transaction logs we retain for point in time restore, from 1-7. (For PostgreSQL Enterprise Plus instances, from 1 to 35.). */
	// +optional
	TransactionLogRetentionDays *int64 `json:"transactionLogRetentionDays,omitempty"`
}

type InstanceBackupRetentionSettings struct {
	/* Number of backups to retain. */
	RetainedBackups int64 `json:"retainedBackups"`

	/* The unit that 'retainedBackups' represents. Defaults to COUNT. */
	// +optional
	RetentionUnit *string `json:"retentionUnit,omitempty"`
}

type InstanceDataCacheConfig struct {
	/* Whether data cache is enabled for the instance. */
	// +optional
	DataCacheEnabled *bool `json:"dataCacheEnabled,omitempty"`
}

type InstanceDatabaseFlags struct {
	/* Name of the flag. */
	Name string `json:"name"`

	/* Value of the flag. */
	Value string `json:"value"`
}

type InstanceDenyMaintenancePeriod struct {
	/* End date before which maintenance will not take place. The date is in format yyyy-mm-dd i.e., 2020-11-01, or mm-dd, i.e., 11-01. */
	EndDate string `json:"endDate"`

	/* Start date after which maintenance will not take place. The date is in format yyyy-mm-dd i.e., 2020-11-01, or mm-dd, i.e., 11-01. */
	StartDate string `json:"startDate"`

	/* Time in UTC when the "deny maintenance period" starts on start_date and ends on end_date. The time is in format: HH:mm:SS, i.e., 00:00:00. */
	Time string `json:"time"`
}

type InstanceInsightsConfig struct {
	/* True if Query Insights feature is enabled. */
	// +optional
	QueryInsightsEnabled *bool `json:"queryInsightsEnabled,omitempty"`

	/* Number of query execution plans captured by Insights per minute for all queries combined. Between 0 and 20. Default to 5. */
	// +optional
	QueryPlansPerMinute *int64 `json:"queryPlansPerMinute,omitempty"`

	/* Maximum query length stored in bytes. Between 256 and 4500. Default to 1024. */
	// +optional
	QueryStringLength *int64 `json:"queryStringLength,omitempty"`

	/* True if Query Insights will record application tags from query when enabled. */
	// +optional
	RecordApplicationTags *bool `json:"recordApplicationTags,omitempty"`

	/* True if Query Insights will record client address when enabled. */
	// +optional
	RecordClientAddress *bool `json:"recordClientAddress,omitempty"`
}

type InstanceIpConfiguration struct {
	/* The name of the allocated ip range for the private ip CloudSQL instance. For example: "google-managed-services-default". If set, the instance ip will be created in the allocated range. The range name must comply with RFC 1035. Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])?. */
	// +optional
	AllocatedIpRange *string `json:"allocatedIpRange,omitempty"`

	// +optional
	AuthorizedNetworks []InstanceAuthorizedNetworks `json:"authorizedNetworks,omitempty"`

	/* Whether Google Cloud services such as BigQuery are allowed to access data in this Cloud SQL instance over a private IP connection. SQLSERVER database type is not supported. */
	// +optional
	EnablePrivatePathForGoogleCloudServices *bool `json:"enablePrivatePathForGoogleCloudServices,omitempty"`

	/* Whether this Cloud SQL instance should be assigned a public IPV4 address. At least ipv4_enabled must be enabled or a private_network must be configured. */
	// +optional
	Ipv4Enabled *bool `json:"ipv4Enabled,omitempty"`

	// +optional
	PrivateNetworkRef *refsv1beta1.ComputeNetworkRef `json:"privateNetworkRef,omitempty"`

	/* PSC settings for a Cloud SQL instance. */
	// +optional
	PscConfig []InstancePscConfig `json:"pscConfig,omitempty"`

	// +optional
	RequireSsl *bool `json:"requireSsl,omitempty"`

	/* Specify how SSL connection should be enforced in DB connections. This field provides more SSL enforcment options compared to requireSsl. To change this field, also set the correspoding value in requireSsl if it has been set. */
	// +optional
	SslMode *string `json:"sslMode,omitempty"`
}

type InstanceLocationPreference struct {
	/* A Google App Engine application whose zone to remain in. Must be in the same region as this instance. */
	// +optional
	FollowGaeApplication *string `json:"followGaeApplication,omitempty"`

	/* The preferred Compute Engine zone for the secondary/failover. */
	// +optional
	SecondaryZone *string `json:"secondaryZone,omitempty"`

	/* The preferred compute engine zone. */
	// +optional
	Zone *string `json:"zone,omitempty"`
}

type InstanceMaintenanceWindow struct {
	/* Day of week (1-7), starting on Monday. */
	// +optional
	Day *int64 `json:"day,omitempty"`

	/* Hour of day (0-23), ignored if day not set. */
	// +optional
	Hour *int64 `json:"hour,omitempty"`

	/* Receive updates earlier (canary) or later (stable). */
	// +optional
	UpdateTrack *string `json:"updateTrack,omitempty"`
}

type InstancePasswordValidationPolicy struct {
	/* Password complexity. */
	// +optional
	Complexity *string `json:"complexity,omitempty"`

	/* Disallow username as a part of the password. */
	// +optional
	DisallowUsernameSubstring *bool `json:"disallowUsernameSubstring,omitempty"`

	/* Whether the password policy is enabled or not. */
	EnablePasswordPolicy bool `json:"enablePasswordPolicy"`

	/* Minimum number of characters allowed. */
	// +optional
	MinLength *int64 `json:"minLength,omitempty"`

	/* Minimum interval after which the password can be changed. This flag is only supported for PostgresSQL. */
	// +optional
	PasswordChangeInterval *string `json:"passwordChangeInterval,omitempty"`

	/* Number of previous passwords that cannot be reused. */
	// +optional
	ReuseInterval *int64 `json:"reuseInterval,omitempty"`
}

type InstancePscConfig struct {
	/* List of consumer projects that are allow-listed for PSC connections to this instance. This instance can be connected to with PSC from any network in these projects. Each consumer project in this list may be represented by a project number (numeric) or by a project id (alphanumeric). */
	// +optional
	AllowedConsumerProjects []string `json:"allowedConsumerProjects,omitempty"`

	/* Whether PSC connectivity is enabled for this instance. */
	// +optional
	PscEnabled *bool `json:"pscEnabled,omitempty"`
}

type InstanceReplicaConfiguration struct {
	/* Immutable. PEM representation of the trusted CA's x509 certificate. */
	// +optional
	CaCertificate *string `json:"caCertificate,omitempty"`

	/* Immutable. PEM representation of the replica's x509 certificate. */
	// +optional
	ClientCertificate *string `json:"clientCertificate,omitempty"`

	/* Immutable. PEM representation of the replica's private key. The corresponding public key in encoded in the client_certificate. */
	// +optional
	ClientKey *string `json:"clientKey,omitempty"`

	/* Immutable. The number of seconds between connect retries. MySQL's default is 60 seconds. */
	// +optional
	ConnectRetryInterval *int64 `json:"connectRetryInterval,omitempty"`

	/* Immutable. Path to a SQL file in Google Cloud Storage from which replica instances are created. Format is gs://bucket/filename. */
	// +optional
	DumpFilePath *string `json:"dumpFilePath,omitempty"`

	/* Immutable. Specifies if the replica is the failover target. If the field is set to true the replica will be designated as a failover replica. If the master instance fails, the replica instance will be promoted as the new master instance. Not supported for Postgres. */
	// +optional
	FailoverTarget *bool `json:"failoverTarget,omitempty"`

	/* Immutable. Time in ms between replication heartbeats. */
	// +optional
	MasterHeartbeatPeriod *int64 `json:"masterHeartbeatPeriod,omitempty"`

	/* Immutable. Password for the replication connection. */
	// +optional
	Password *refsv1beta1secret.Legacy `json:"password,omitempty"`

	/* Immutable. Permissible ciphers for use in SSL encryption. */
	// +optional
	SslCipher *string `json:"sslCipher,omitempty"`

	/* Immutable. Username for replication connection. */
	// +optional
	Username *string `json:"username,omitempty"`

	/* Immutable. True if the master's common name value is checked during the SSL handshake. */
	// +optional
	VerifyServerCertificate *bool `json:"verifyServerCertificate,omitempty"`
}

type InstanceSettings struct {
	/* This specifies when the instance should be active. Can be either ALWAYS, NEVER or ON_DEMAND. */
	// +optional
	ActivationPolicy *string `json:"activationPolicy,omitempty"`

	// +optional
	ActiveDirectoryConfig *InstanceActiveDirectoryConfig `json:"activeDirectoryConfig,omitempty"`

	// +optional
	AdvancedMachineFeatures *InstanceAdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty"`

	/* DEPRECATED. This property is only applicable to First Generation instances, and First Generation instances are now deprecated. see https://cloud.google.com/sql/docs/mysql/deprecation-notice for information on how to upgrade to Second Generation instances.
	Specifying this field has no-ops; it's recommended to remove this field from your configuration. */
	// +optional
	AuthorizedGaeApplications []string `json:"authorizedGaeApplications,omitempty"`

	/* The availability type of the Cloud SQL instance, high availability
	(REGIONAL) or single zone (ZONAL). For all instances, ensure that
	settings.backup_configuration.enabled is set to true.
	For MySQL instances, ensure that settings.backup_configuration.binary_log_enabled is set to true.
	For Postgres instances, ensure that settings.backup_configuration.point_in_time_recovery_enabled
	is set to true. Defaults to ZONAL. */
	// +optional
	AvailabilityType *string `json:"availabilityType,omitempty"`

	// +optional
	BackupConfiguration *InstanceBackupConfiguration `json:"backupConfiguration,omitempty"`

	/* Immutable. The name of server instance collation. */
	// +optional
	Collation *string `json:"collation,omitempty"`

	/* Specifies if connections must use Cloud SQL connectors. */
	// +optional
	ConnectorEnforcement *string `json:"connectorEnforcement,omitempty"`

	/* DEPRECATED. This property is only applicable to First Generation instances, and First Generation instances are now deprecated. see https://cloud.google.com/sql/docs/mysql/deprecation-notice for information on how to upgrade to Second Generation instances.
	Specifying this field has no-ops; it's recommended to remove this field from your configuration. */
	// +optional
	CrashSafeReplication *bool `json:"crashSafeReplication,omitempty"`

	/* Data cache configurations. */
	// +optional
	DataCacheConfig *InstanceDataCacheConfig `json:"dataCacheConfig,omitempty"`

	// +optional
	DatabaseFlags []InstanceDatabaseFlags `json:"databaseFlags,omitempty"`

	/* Configuration to protect against accidental instance deletion. */
	// +optional
	DeletionProtectionEnabled *bool `json:"deletionProtectionEnabled,omitempty"`

	// +optional
	DenyMaintenancePeriod *InstanceDenyMaintenancePeriod `json:"denyMaintenancePeriod,omitempty"`

	/* Enables auto-resizing of the storage size. Defaults to true. */
	// +optional
	DiskAutoresize *bool `json:"diskAutoresize,omitempty"`

	/* The maximum size, in GB, to which storage capacity can be automatically increased. The default value is 0, which specifies that there is no limit. */
	// +optional
	DiskAutoresizeLimit *int64 `json:"diskAutoresizeLimit,omitempty"`

	/* The size of data disk, in GB. Size of a running instance cannot be reduced but can be increased. The minimum value is 10GB. */
	// +optional
	DiskSize *int64 `json:"diskSize,omitempty"`

	/* Immutable. The type of data disk: PD_SSD or PD_HDD. Defaults to PD_SSD. */
	// +optional
	DiskType *string `json:"diskType,omitempty"`

	/* The edition of the instance, can be ENTERPRISE or ENTERPRISE_PLUS. */
	// +optional
	Edition *string `json:"edition,omitempty"`

	/* Configuration of Query Insights. */
	// +optional
	InsightsConfig *InstanceInsightsConfig `json:"insightsConfig,omitempty"`

	// +optional
	IpConfiguration *InstanceIpConfiguration `json:"ipConfiguration,omitempty"`

	// +optional
	LocationPreference *InstanceLocationPreference `json:"locationPreference,omitempty"`

	/* Declares a one-hour maintenance window when an Instance can automatically restart to apply updates. The maintenance window is specified in UTC time. */
	// +optional
	MaintenanceWindow *InstanceMaintenanceWindow `json:"maintenanceWindow,omitempty"`

	// +optional
	PasswordValidationPolicy *InstancePasswordValidationPolicy `json:"passwordValidationPolicy,omitempty"`

	/* Pricing plan for this instance, can only be PER_USE. */
	// +optional
	PricingPlan *string `json:"pricingPlan,omitempty"`

	/* DEPRECATED. This property is only applicable to First Generation instances, and First Generation instances are now deprecated. see https://cloud.google.com/sql/docs/mysql/deprecation-notice for information on how to upgrade to Second Generation instances.
	Specifying this field has no-ops; it's recommended to remove this field from your configuration. */
	// +optional
	ReplicationType *string `json:"replicationType,omitempty"`

	// +optional
	SqlServerAuditConfig *InstanceSqlServerAuditConfig `json:"sqlServerAuditConfig,omitempty"`

	/* The machine type to use. See tiers for more details and supported versions. Postgres supports only shared-core machine types, and custom machine types such as db-custom-2-13312. See the Custom Machine Type Documentation to learn about specifying custom machine types. */
	Tier string `json:"tier"`

	/* Immutable. The time_zone to be used by the database engine (supported only for SQL Server), in SQL Server timezone format. */
	// +optional
	TimeZone *string `json:"timeZone,omitempty"`
}

type InstanceSqlServerAuditConfig struct {
	/* The name of the destination bucket (e.g., gs://mybucket). */
	// +optional
	BucketRef *refsv1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	/* How long to keep generated audit files. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".. */
	// +optional
	RetentionInterval *string `json:"retentionInterval,omitempty"`

	/* How often to upload generated audit files. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". */
	// +optional
	UploadInterval *string `json:"uploadInterval,omitempty"`
}

type BinLogCoordinates struct {
	/* Name of the binary log file for a Cloud SQL instance. */
	BinLogFileName string `json:"binLogFileName,omitempty"`

	/* Position (offset) within the binary log file. */
	BinLogPosition int64 `json:"binLogPosition,omitempty,string"`
}

type CloneSource struct {
	/* Binary log coordinates, if specified, identify the position up to which the source instance is
	cloned. If not specified, the source instance is cloned up to the most recent binary log coordinates. */
	// +optional
	BinLogCoordinates *BinLogCoordinates `json:"binLogCoordinates,omitempty"`

	/* (SQL Server only) Clone only the specified databases from the source instance. Clone all databases if empty. */
	// +optional
	DatabaseNames []string `json:"databaseNames,omitempty"`

	/* Timestamp, if specified, identifies the time to which the source instance is cloned. */
	// +optional
	PointInTime *string `json:"pointInTime,omitempty"`

	/* The source SQLInstance to clone */
	SQLInstanceRef SQLInstanceRef `json:"sqlInstanceRef,omitempty"`
}

type SQLInstanceSpec struct {
	/* Create this database as a clone of a source instance. Immutable. */
	// +optional
	CloneSource *CloneSource `json:"cloneSource,omitempty"`

	/* The MySQL, PostgreSQL or SQL Server (beta) version to use. Supported values include MYSQL_5_6, MYSQL_5_7, MYSQL_8_0, POSTGRES_9_6, POSTGRES_10, POSTGRES_11, POSTGRES_12, POSTGRES_13, POSTGRES_14, POSTGRES_15, SQLSERVER_2017_STANDARD, SQLSERVER_2017_ENTERPRISE, SQLSERVER_2017_EXPRESS, SQLSERVER_2017_WEB. Database Version Policies includes an up-to-date reference of supported versions. */
	// +optional
	DatabaseVersion *string `json:"databaseVersion,omitempty"`

	// +optional
	EncryptionKMSCryptoKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"encryptionKMSCryptoKeyRef,omitempty"`

	/* The type of the instance. The valid values are:- 'SQL_INSTANCE_TYPE_UNSPECIFIED', 'CLOUD_SQL_INSTANCE', 'ON_PREMISES_INSTANCE' and 'READ_REPLICA_INSTANCE'. */
	// +optional
	InstanceType *string `json:"instanceType,omitempty"`

	/* Maintenance version. */
	// +optional
	MaintenanceVersion *string `json:"maintenanceVersion,omitempty"`

	// +optional
	MasterInstanceRef *SQLInstanceRef `json:"masterInstanceRef,omitempty"`

	/* Immutable. The region the instance will sit in. Note, Cloud SQL is not available in all regions. A valid region must be provided to use this resource. If a region is not provided in the resource definition, the provider region will be used instead, but this will be an apply-time error for instances if the provider region is not supported with Cloud SQL. If you choose not to provide the region argument for this resource, make sure you understand this. */
	// +optional
	Region *string `json:"region,omitempty"`

	/* The configuration for replication. */
	// +optional
	ReplicaConfiguration *InstanceReplicaConfiguration `json:"replicaConfiguration,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Initial root password. Required for MS SQL Server. */
	// +optional
	RootPassword *refsv1beta1secret.Legacy `json:"rootPassword,omitempty"`

	/* The settings to use for the database. The configuration is detailed below. */
	Settings InstanceSettings `json:"settings"`
}

type InstanceIpAddressStatus struct {
	// +optional
	IpAddress *string `json:"ipAddress,omitempty"`

	// +optional
	TimeToRetire *string `json:"timeToRetire,omitempty"`

	// +optional
	Type *string `json:"type,omitempty"`
}

type InstanceServerCaCertStatus struct {
	/* The CA Certificate used to connect to the SQL Instance via SSL. */
	// +optional
	Cert *string `json:"cert,omitempty"`

	/* The CN valid for the CA Cert. */
	// +optional
	CommonName *string `json:"commonName,omitempty"`

	/* Creation time of the CA Cert. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Expiration time of the CA Cert. */
	// +optional
	ExpirationTime *string `json:"expirationTime,omitempty"`

	/* SHA Fingerprint of the CA Cert. */
	// +optional
	Sha1Fingerprint *string `json:"sha1Fingerprint,omitempty"`
}

type SQLInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   SQLInstance's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Available Maintenance versions. */
	// +optional
	AvailableMaintenanceVersions []string `json:"availableMaintenanceVersions,omitempty"`

	/* The connection name of the instance to be used in connection strings. For example, when connecting with Cloud SQL Proxy. */
	// +optional
	ConnectionName *string `json:"connectionName,omitempty"`

	/* The dns name of the instance. */
	// +optional
	DnsName *string `json:"dnsName,omitempty"`

	// +optional
	FirstIpAddress *string `json:"firstIpAddress,omitempty"`

	/* The type of the instance. The valid values are:- 'SQL_INSTANCE_TYPE_UNSPECIFIED', 'CLOUD_SQL_INSTANCE', 'ON_PREMISES_INSTANCE' and 'READ_REPLICA_INSTANCE'. */
	// +optional
	InstanceType *string `json:"instanceType,omitempty"`

	// +optional
	IpAddress []InstanceIpAddressStatus `json:"ipAddress,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// +optional
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`

	/* The link to service attachment of PSC instance. */
	// +optional
	PscServiceAttachmentLink *string `json:"pscServiceAttachmentLink,omitempty"`

	// +optional
	PublicIpAddress *string `json:"publicIpAddress,omitempty"`

	/* The URI of the created resource. */
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	// +optional
	ServerCaCert *InstanceServerCaCertStatus `json:"serverCaCert,omitempty"`

	/* The service account email address assigned to the instance. */
	// +optional
	ServiceAccountEmailAddress *string `json:"serviceAccountEmailAddress,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsqlinstance;gcpsqlinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SQLInstance is the Schema for the sql API
// +k8s:openapi-gen=true
type SQLInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SQLInstanceSpec   `json:"spec"`
	Status SQLInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SQLInstanceList contains a list of SQLInstance
type SQLInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SQLInstance{}, &SQLInstanceList{})
}
