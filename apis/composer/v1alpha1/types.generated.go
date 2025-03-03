// Copyright 2025 Google LLC
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

package v1alpha1

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.AirflowMetadataRetentionPolicyConfig
type AirflowMetadataRetentionPolicyConfig struct {
	// Optional. Retention can be either enabled or disabled.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.AirflowMetadataRetentionPolicyConfig.retention_mode
	RetentionMode *string `json:"retentionMode,omitempty"`

	// Optional. How many days data should be retained for.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.AirflowMetadataRetentionPolicyConfig.retention_days
	RetentionDays *int32 `json:"retentionDays,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.CloudDataLineageIntegration
type CloudDataLineageIntegration struct {
	// Optional. Whether or not Cloud Data Lineage integration is enabled.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.CloudDataLineageIntegration.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.DataRetentionConfig
type DataRetentionConfig struct {
	// Optional. The retention policy for airflow metadata database.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.DataRetentionConfig.airflow_metadata_retention_config
	AirflowMetadataRetentionConfig *AirflowMetadataRetentionPolicyConfig `json:"airflowMetadataRetentionConfig,omitempty"`

	// Optional. The configuration settings for task logs retention
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.DataRetentionConfig.task_logs_retention_config
	TaskLogsRetentionConfig *TaskLogsRetentionConfig `json:"taskLogsRetentionConfig,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.DatabaseConfig
type DatabaseConfig struct {
	// Optional. Cloud SQL machine type used by Airflow database.
	//  It has to be one of: db-n1-standard-2, db-n1-standard-4, db-n1-standard-8
	//  or db-n1-standard-16. If not specified, db-n1-standard-2 will be used.
	//  Supported for Cloud Composer environments in versions
	//  composer-1.*.*-airflow-*.*.*.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.DatabaseConfig.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Optional. The Compute Engine zone where the Airflow database is created. If
	//  zone is provided, it must be in the region selected for the environment. If
	//  zone is not provided, a zone is automatically selected. The zone can only
	//  be set during environment creation. Supported for Cloud Composer
	//  environments in versions composer-2.*.*-airflow-*.*.*.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.DatabaseConfig.zone
	Zone *string `json:"zone,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.IPAllocationPolicy
type IPAllocationPolicy struct {
	// Optional. Whether or not to enable Alias IPs in the GKE cluster.
	//  If `true`, a VPC-native cluster is created.
	//
	//  This field is only supported for Cloud Composer environments in versions
	//  composer-1.*.*-airflow-*.*.*. Environments in newer versions always use
	//  VPC-native GKE clusters.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.IPAllocationPolicy.use_ip_aliases
	UseIPAliases *bool `json:"useIPAliases,omitempty"`

	// Optional. The name of the GKE cluster's secondary range used to allocate
	//  IP addresses to pods.
	//
	//  For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*,
	//  this field is applicable only when `use_ip_aliases` is true.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.IPAllocationPolicy.cluster_secondary_range_name
	ClusterSecondaryRangeName *string `json:"clusterSecondaryRangeName,omitempty"`

	// Optional. The IP address range used to allocate IP addresses to pods in
	//  the GKE cluster.
	//
	//  For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*,
	//  this field is applicable only when `use_ip_aliases` is true.
	//
	//  Set to blank to have GKE choose a range with the default size.
	//
	//  Set to /netmask (e.g. `/14`) to have GKE choose a range with a specific
	//  netmask.
	//
	//  Set to a
	//  [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	//  notation (e.g. `10.96.0.0/14`) from the RFC-1918 private networks (e.g.
	//  `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`) to pick a specific range
	//  to use.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.IPAllocationPolicy.cluster_ipv4_cidr_block
	ClusterIPV4CIDRBlock *string `json:"clusterIPV4CIDRBlock,omitempty"`

	// Optional. The name of the services' secondary range used to allocate
	//  IP addresses to the GKE cluster.
	//
	//  For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*,
	//  this field is applicable only when `use_ip_aliases` is true.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.IPAllocationPolicy.services_secondary_range_name
	ServicesSecondaryRangeName *string `json:"servicesSecondaryRangeName,omitempty"`

	// Optional. The IP address range of the services IP addresses in this
	//  GKE cluster.
	//
	//  For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*,
	//  this field is applicable only when `use_ip_aliases` is true.
	//
	//  Set to blank to have GKE choose a range with the default size.
	//
	//  Set to /netmask (e.g. `/14`) to have GKE choose a range with a specific
	//  netmask.
	//
	//  Set to a
	//  [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	//  notation (e.g. `10.96.0.0/14`) from the RFC-1918 private networks (e.g.
	//  `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`) to pick a specific range
	//  to use.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.IPAllocationPolicy.services_ipv4_cidr_block
	ServicesIPV4CIDRBlock *string `json:"servicesIPV4CIDRBlock,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.MaintenanceWindow
type MaintenanceWindow struct {
	// Required. Start time of the first recurrence of the maintenance window.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.MaintenanceWindow.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Required. Maintenance window end time. It is used only to calculate the
	//  duration of the maintenance window. The value for end-time must be in the
	//  future, relative to `start_time`.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.MaintenanceWindow.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Required. Maintenance window recurrence. Format is a subset of
	//  [RFC-5545](https://tools.ietf.org/html/rfc5545) `RRULE`. The only allowed
	//  values for `FREQ` field are `FREQ=DAILY` and `FREQ=WEEKLY;BYDAY=...`
	//  Example values: `FREQ=WEEKLY;BYDAY=TU,WE`, `FREQ=DAILY`.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.MaintenanceWindow.recurrence
	Recurrence *string `json:"recurrence,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.MasterAuthorizedNetworksConfig
type MasterAuthorizedNetworksConfig struct {
	// Whether or not master authorized networks feature is enabled.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.MasterAuthorizedNetworksConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Up to 50 external networks that could access Kubernetes master through
	//  HTTPS.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.MasterAuthorizedNetworksConfig.cidr_blocks
	CIDRBlocks []MasterAuthorizedNetworksConfig_CIDRBlock `json:"cidrBlocks,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.MasterAuthorizedNetworksConfig.CidrBlock
type MasterAuthorizedNetworksConfig_CIDRBlock struct {
	// User-defined name that identifies the CIDR block.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.MasterAuthorizedNetworksConfig.CidrBlock.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// CIDR block that must be specified in CIDR notation.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.MasterAuthorizedNetworksConfig.CidrBlock.cidr_block
	CIDRBlock *string `json:"cidrBlock,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.NetworkingConfig
type NetworkingConfig struct {
	// Optional. Indicates the user requested specifc connection type between
	//  Tenant and Customer projects. You cannot set networking connection type in
	//  public IP environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.NetworkingConfig.connection_type
	ConnectionType *string `json:"connectionType,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.PrivateClusterConfig
type PrivateClusterConfig struct {
	// Optional. If `true`, access to the public endpoint of the GKE cluster is
	//  denied.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateClusterConfig.enable_private_endpoint
	EnablePrivateEndpoint *bool `json:"enablePrivateEndpoint,omitempty"`

	// Optional. The CIDR block from which IPv4 range for GKE master will be
	//  reserved. If left blank, the default value of '172.16.0.0/23' is used.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateClusterConfig.master_ipv4_cidr_block
	MasterIPV4CIDRBlock *string `json:"masterIPV4CIDRBlock,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.RecoveryConfig
type RecoveryConfig struct {
	// Optional. The configuration for scheduled snapshot creation mechanism.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.RecoveryConfig.scheduled_snapshots_config
	ScheduledSnapshotsConfig *ScheduledSnapshotsConfig `json:"scheduledSnapshotsConfig,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.ScheduledSnapshotsConfig
type ScheduledSnapshotsConfig struct {
	// Optional. Whether scheduled snapshots creation is enabled.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.ScheduledSnapshotsConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. The Cloud Storage location for storing automatically created
	//  snapshots.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.ScheduledSnapshotsConfig.snapshot_location
	SnapshotLocation *string `json:"snapshotLocation,omitempty"`

	// Optional. The cron expression representing the time when snapshots creation
	//  mechanism runs. This field is subject to additional validation around
	//  frequency of execution.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.ScheduledSnapshotsConfig.snapshot_creation_schedule
	SnapshotCreationSchedule *string `json:"snapshotCreationSchedule,omitempty"`

	// Optional. Time zone that sets the context to interpret
	//  snapshot_creation_schedule.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.ScheduledSnapshotsConfig.time_zone
	TimeZone *string `json:"timeZone,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.SoftwareConfig
type SoftwareConfig struct {
	// Optional. The version of the software running in the environment.
	//  This encapsulates both the version of Cloud Composer functionality and the
	//  version of Apache Airflow. It must match the regular expression
	//  `composer-([0-9]+(\.[0-9]+\.[0-9]+(-preview\.[0-9]+)?)?|latest)-airflow-([0-9]+(\.[0-9]+(\.[0-9]+)?)?)`.
	//  When used as input, the server also checks if the provided version is
	//  supported and denies the request for an unsupported version.
	//
	//  The Cloud Composer portion of the image version is a full
	//  [semantic version](https://semver.org), or an alias in the form of major
	//  version number or `latest`. When an alias is provided, the server replaces
	//  it with the current Cloud Composer version that satisfies the alias.
	//
	//  The Apache Airflow portion of the image version is a full semantic version
	//  that points to one of the supported Apache Airflow versions, or an alias in
	//  the form of only major or major.minor versions specified. When an alias is
	//  provided, the server replaces it with the latest Apache Airflow version
	//  that satisfies the alias and is supported in the given Cloud Composer
	//  version.
	//
	//  In all cases, the resolved image version is stored in the same field.
	//
	//  See also [version
	//  list](/composer/docs/concepts/versioning/composer-versions) and [versioning
	//  overview](/composer/docs/concepts/versioning/composer-versioning-overview).
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.SoftwareConfig.image_version
	ImageVersion *string `json:"imageVersion,omitempty"`

	// Optional. Apache Airflow configuration properties to override.
	//
	//  Property keys contain the section and property names, separated by a
	//  hyphen, for example "core-dags_are_paused_at_creation". Section names must
	//  not contain hyphens ("-"), opening square brackets ("["),  or closing
	//  square brackets ("]"). The property name must not be empty and must not
	//  contain an equals sign ("=") or semicolon (";"). Section and property names
	//  must not contain a period ("."). Apache Airflow configuration property
	//  names must be written in
	//  [snake_case](https://en.wikipedia.org/wiki/Snake_case). Property values can
	//  contain any character, and can be written in any lower/upper case format.
	//
	//  Certain Apache Airflow configuration property values are
	//  [blocked](/composer/docs/concepts/airflow-configurations),
	//  and cannot be overridden.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.SoftwareConfig.airflow_config_overrides
	AirflowConfigOverrides map[string]string `json:"airflowConfigOverrides,omitempty"`

	// Optional. Custom Python Package Index (PyPI) packages to be installed in
	//  the environment.
	//
	//  Keys refer to the lowercase package name such as "numpy"
	//  and values are the lowercase extras and version specifier such as
	//  "==1.12.0", "[devel,gcp_api]", or "[devel]>=1.8.2, <1.9.2". To specify a
	//  package without pinning it to a version specifier, use the empty string as
	//  the value.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.SoftwareConfig.pypi_packages
	PypiPackages map[string]string `json:"pypiPackages,omitempty"`

	// Optional. Additional environment variables to provide to the Apache Airflow
	//  scheduler, worker, and webserver processes.
	//
	//  Environment variable names must match the regular expression
	//  `[a-zA-Z_][a-zA-Z0-9_]*`. They cannot specify Apache Airflow
	//  software configuration overrides (they cannot match the regular expression
	//  `AIRFLOW__[A-Z0-9_]+__[A-Z0-9_]+`), and they cannot match any of the
	//  following reserved names:
	//
	//  * `AIRFLOW_HOME`
	//  * `C_FORCE_ROOT`
	//  * `CONTAINER_NAME`
	//  * `DAGS_FOLDER`
	//  * `GCP_PROJECT`
	//  * `GCS_BUCKET`
	//  * `GKE_CLUSTER_NAME`
	//  * `SQL_DATABASE`
	//  * `SQL_INSTANCE`
	//  * `SQL_PASSWORD`
	//  * `SQL_PROJECT`
	//  * `SQL_REGION`
	//  * `SQL_USER`
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.SoftwareConfig.env_variables
	EnvVariables map[string]string `json:"envVariables,omitempty"`

	// Optional. The major version of Python used to run the Apache Airflow
	//  scheduler, worker, and webserver processes.
	//
	//  Can be set to '2' or '3'. If not specified, the default is '3'. Cannot be
	//  updated.
	//
	//  This field is only supported for Cloud Composer environments in versions
	//  composer-1.*.*-airflow-*.*.*. Environments in newer versions always use
	//  Python major version 3.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.SoftwareConfig.python_version
	PythonVersion *string `json:"pythonVersion,omitempty"`

	// Optional. The number of schedulers for Airflow.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-1.*.*-airflow-2.*.*.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.SoftwareConfig.scheduler_count
	SchedulerCount *int32 `json:"schedulerCount,omitempty"`

	// Optional. The configuration for Cloud Data Lineage integration.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.SoftwareConfig.cloud_data_lineage_integration
	CloudDataLineageIntegration *CloudDataLineageIntegration `json:"cloudDataLineageIntegration,omitempty"`

	// Optional. Whether or not the web server uses custom plugins.
	//  If unspecified, the field defaults to `PLUGINS_ENABLED`.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-3.*.*-airflow-*.*.* and newer.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.SoftwareConfig.web_server_plugins_mode
	WebServerPluginsMode *string `json:"webServerPluginsMode,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.TaskLogsRetentionConfig
type TaskLogsRetentionConfig struct {
	// Optional. The mode of storage for Airflow workers task logs.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.TaskLogsRetentionConfig.storage_mode
	StorageMode *string `json:"storageMode,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.WebServerConfig
type WebServerConfig struct {
	// Optional. Machine type on which Airflow web server is running.
	//  It has to be one of: composer-n1-webserver-2, composer-n1-webserver-4 or
	//  composer-n1-webserver-8.
	//  If not specified, composer-n1-webserver-2 will be used.
	//  Value custom is returned only in response, if Airflow web server parameters
	//  were manually changed to a non-standard values.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WebServerConfig.machine_type
	MachineType *string `json:"machineType,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.WebServerNetworkAccessControl
type WebServerNetworkAccessControl struct {
	// A collection of allowed IP ranges with descriptions.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WebServerNetworkAccessControl.allowed_ip_ranges
	AllowedIPRanges []WebServerNetworkAccessControl_AllowedIPRange `json:"allowedIPRanges,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.WebServerNetworkAccessControl.AllowedIpRange
type WebServerNetworkAccessControl_AllowedIPRange struct {
	// IP address or range, defined using CIDR notation, of requests that this
	//  rule applies to.
	//  Examples: `192.168.1.1` or `192.168.0.0/16` or `2001:db8::/32`
	//            or `2001:0db8:0000:0042:0000:8a2e:0370:7334`.
	//
	//  IP range prefixes should be properly truncated. For example,
	//  `1.2.3.4/24` should be truncated to `1.2.3.0/24`. Similarly, for IPv6,
	//  `2001:db8::1/32` should be truncated to `2001:db8::/32`.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WebServerNetworkAccessControl.AllowedIpRange.value
	Value *string `json:"value,omitempty"`

	// Optional. User-provided description. It must contain at most 300
	//  characters.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WebServerNetworkAccessControl.AllowedIpRange.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig
type WorkloadsConfig struct {
	// Optional. Resources used by Airflow schedulers.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.scheduler
	Scheduler *WorkloadsConfig_SchedulerResource `json:"scheduler,omitempty"`

	// Optional. Resources used by Airflow web server.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.web_server
	WebServer *WorkloadsConfig_WebServerResource `json:"webServer,omitempty"`

	// Optional. Resources used by Airflow workers.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.worker
	Worker *WorkloadsConfig_WorkerResource `json:"worker,omitempty"`

	// Optional. Resources used by Airflow triggerers.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.triggerer
	Triggerer *WorkloadsConfig_TriggererResource `json:"triggerer,omitempty"`

	// Optional. Resources used by Airflow DAG processors.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-3.*.*-airflow-*.*.* and newer.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.dag_processor
	DagProcessor *WorkloadsConfig_DagProcessorResource `json:"dagProcessor,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.PrivateClusterConfig
type PrivateClusterConfigObservedState struct {
	// Output only. The IP range in CIDR notation to use for the hosted master
	//  network. This range is used for assigning internal IP addresses to the GKE
	//  cluster master or set of masters and to the internal load balancer virtual
	//  IP. This range must not overlap with any other ranges in use within the
	//  cluster's network.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateClusterConfig.master_ipv4_reserved_range
	MasterIPV4ReservedRange *string `json:"masterIPV4ReservedRange,omitempty"`
}
