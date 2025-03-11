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


// +kcc:proto=google.cloud.dataproc.v1.AcceleratorConfig
type AcceleratorConfig struct {
	// Full URL, partial URI, or short name of the accelerator type resource to
	//  expose to this instance. See
	//  [Compute Engine
	//  AcceleratorTypes](https://cloud.google.com/compute/docs/reference/v1/acceleratorTypes).
	//
	//  Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]/acceleratorTypes/nvidia-tesla-t4`
	//  * `projects/[project_id]/zones/[zone]/acceleratorTypes/nvidia-tesla-t4`
	//  * `nvidia-tesla-t4`
	//
	//  **Auto Zone Exception**: If you are using the Dataproc
	//  [Auto Zone
	//  Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
	//  feature, you must use the short name of the accelerator type
	//  resource, for example, `nvidia-tesla-t4`.
	// +kcc:proto:field=google.cloud.dataproc.v1.AcceleratorConfig.accelerator_type_uri
	AcceleratorTypeURI *string `json:"acceleratorTypeURI,omitempty"`

	// The number of the accelerator cards of this type exposed to this instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.AcceleratorConfig.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.AutoscalingConfig
type AutoscalingConfig struct {
	// Optional. The autoscaling policy used by the cluster.
	//
	//  Only resource names including projectid and location (region) are valid.
	//  Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]`
	//  * `projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]`
	//
	//  Note that the policy must be in the same project and Dataproc region.
	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingConfig.policy_uri
	PolicyURI *string `json:"policyURI,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.AuxiliaryNodeGroup
type AuxiliaryNodeGroup struct {
	// Required. Node group configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.AuxiliaryNodeGroup.node_group
	NodeGroup *NodeGroup `json:"nodeGroup,omitempty"`

	// Optional. A node group ID. Generated if not specified.
	//
	//  The ID must contain only letters (a-z, A-Z), numbers (0-9),
	//  underscores (_), and hyphens (-). Cannot begin or end with underscore
	//  or hyphen. Must consist of from 3 to 33 characters.
	// +kcc:proto:field=google.cloud.dataproc.v1.AuxiliaryNodeGroup.node_group_id
	NodeGroupID *string `json:"nodeGroupID,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ClusterConfig
type ClusterConfig struct {
	// Optional. A Cloud Storage bucket used to stage job
	//  dependencies, config files, and job driver console output.
	//  If you do not specify a staging bucket, Cloud
	//  Dataproc will determine a Cloud Storage location (US,
	//  ASIA, or EU) for your cluster's staging bucket according to the
	//  Compute Engine zone where your cluster is deployed, and then create
	//  and manage this project-level, per-location bucket (see
	//  [Dataproc staging and temp
	//  buckets](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)).
	//  **This field requires a Cloud Storage bucket name, not a `gs://...` URI to
	//  a Cloud Storage bucket.**
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.config_bucket
	ConfigBucket *string `json:"configBucket,omitempty"`

	// Optional. A Cloud Storage bucket used to store ephemeral cluster and jobs
	//  data, such as Spark and MapReduce history files. If you do not specify a
	//  temp bucket, Dataproc will determine a Cloud Storage location (US, ASIA, or
	//  EU) for your cluster's temp bucket according to the Compute Engine zone
	//  where your cluster is deployed, and then create and manage this
	//  project-level, per-location bucket. The default bucket has a TTL of 90
	//  days, but you can use any TTL (or none) if you specify a bucket (see
	//  [Dataproc staging and temp
	//  buckets](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)).
	//  **This field requires a Cloud Storage bucket name, not a `gs://...` URI to
	//  a Cloud Storage bucket.**
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.temp_bucket
	TempBucket *string `json:"tempBucket,omitempty"`

	// Optional. The shared Compute Engine config settings for
	//  all instances in a cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.gce_cluster_config
	GCEClusterConfig *GCEClusterConfig `json:"gceClusterConfig,omitempty"`

	// Optional. The Compute Engine config settings for
	//  the cluster's master instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.master_config
	MasterConfig *InstanceGroupConfig `json:"masterConfig,omitempty"`

	// Optional. The Compute Engine config settings for
	//  the cluster's worker instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.worker_config
	WorkerConfig *InstanceGroupConfig `json:"workerConfig,omitempty"`

	// Optional. The Compute Engine config settings for
	//  a cluster's secondary worker instances
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.secondary_worker_config
	SecondaryWorkerConfig *InstanceGroupConfig `json:"secondaryWorkerConfig,omitempty"`

	// Optional. The config settings for cluster software.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.software_config
	SoftwareConfig *SoftwareConfig `json:"softwareConfig,omitempty"`

	// Optional. Commands to execute on each node after config is
	//  completed. By default, executables are run on master and all worker nodes.
	//  You can test a node's `role` metadata to run an executable on
	//  a master or worker node, as shown below using `curl` (you can also use
	//  `wget`):
	//
	//      ROLE=$(curl -H Metadata-Flavor:Google
	//      http://metadata/computeMetadata/v1/instance/attributes/dataproc-role)
	//      if [[ "${ROLE}" == 'Master' ]]; then
	//        ... master specific actions ...
	//      else
	//        ... worker specific actions ...
	//      fi
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.initialization_actions
	InitializationActions []NodeInitializationAction `json:"initializationActions,omitempty"`

	// Optional. Encryption settings for the cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// Optional. Autoscaling config for the policy associated with the cluster.
	//  Cluster does not autoscale if this field is unset.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.autoscaling_config
	AutoscalingConfig *AutoscalingConfig `json:"autoscalingConfig,omitempty"`

	// Optional. Security settings for the cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.security_config
	SecurityConfig *SecurityConfig `json:"securityConfig,omitempty"`

	// Optional. Lifecycle setting for the cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.lifecycle_config
	LifecycleConfig *LifecycleConfig `json:"lifecycleConfig,omitempty"`

	// Optional. Port/endpoint configuration for this cluster
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.endpoint_config
	EndpointConfig *EndpointConfig `json:"endpointConfig,omitempty"`

	// Optional. Metastore configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.metastore_config
	MetastoreConfig *MetastoreConfig `json:"metastoreConfig,omitempty"`

	// Optional. The config for Dataproc metrics.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.dataproc_metric_config
	DataprocMetricConfig *DataprocMetricConfig `json:"dataprocMetricConfig,omitempty"`

	// Optional. The node group settings.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.auxiliary_node_groups
	AuxiliaryNodeGroups []AuxiliaryNodeGroup `json:"auxiliaryNodeGroups,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ClusterSelector
type ClusterSelector struct {
	// Optional. The zone where workflow process executes. This parameter does not
	//  affect the selection of the cluster.
	//
	//  If unspecified, the zone of the first cluster matching the selector
	//  is used.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterSelector.zone
	Zone *string `json:"zone,omitempty"`

	// Required. The cluster labels. Cluster must have all labels
	//  to match.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterSelector.cluster_labels
	ClusterLabels map[string]string `json:"clusterLabels,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ConfidentialInstanceConfig
type ConfidentialInstanceConfig struct {
	// Optional. Defines whether the instance should have confidential compute
	//  enabled.
	// +kcc:proto:field=google.cloud.dataproc.v1.ConfidentialInstanceConfig.enable_confidential_compute
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.DataprocMetricConfig
type DataprocMetricConfig struct {
	// Required. Metrics sources to enable.
	// +kcc:proto:field=google.cloud.dataproc.v1.DataprocMetricConfig.metrics
	Metrics []DataprocMetricConfig_Metric `json:"metrics,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.DataprocMetricConfig.Metric
type DataprocMetricConfig_Metric struct {
	// Required. A standard set of metrics is collected unless `metricOverrides`
	//  are specified for the metric source (see [Custom metrics]
	//  (https://cloud.google.com/dataproc/docs/guides/dataproc-metrics#custom_metrics)
	//  for more information).
	// +kcc:proto:field=google.cloud.dataproc.v1.DataprocMetricConfig.Metric.metric_source
	MetricSource *string `json:"metricSource,omitempty"`

	// Optional. Specify one or more [Custom metrics]
	//  (https://cloud.google.com/dataproc/docs/guides/dataproc-metrics#custom_metrics)
	//  to collect for the metric course (for the `SPARK` metric source (any
	//  [Spark metric]
	//  (https://spark.apache.org/docs/latest/monitoring.html#metrics) can be
	//  specified).
	//
	//  Provide metrics in the following format:
	//  <code><var>METRIC_SOURCE</var>:<var>INSTANCE</var>:<var>GROUP</var>:<var>METRIC</var></code>
	//  Use camelcase as appropriate.
	//
	//  Examples:
	//
	//  ```
	//  yarn:ResourceManager:QueueMetrics:AppsCompleted
	//  spark:driver:DAGScheduler:job.allJobs
	//  sparkHistoryServer:JVM:Memory:NonHeapMemoryUsage.committed
	//  hiveserver2:JVM:Memory:NonHeapMemoryUsage.used
	//  ```
	//
	//  Notes:
	//
	//  * Only the specified overridden metrics are collected for the
	//    metric source. For example, if one or more `spark:executive` metrics
	//    are listed as metric overrides, other `SPARK` metrics are not
	//    collected. The collection of the metrics for other enabled custom
	//    metric sources is unaffected. For example, if both `SPARK` andd `YARN`
	//    metric sources are enabled, and overrides are provided for Spark
	//    metrics only, all YARN metrics are collected.
	// +kcc:proto:field=google.cloud.dataproc.v1.DataprocMetricConfig.Metric.metric_overrides
	MetricOverrides []string `json:"metricOverrides,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.DiskConfig
type DiskConfig struct {
	// Optional. Type of the boot disk (default is "pd-standard").
	//  Valid values: "pd-balanced" (Persistent Disk Balanced Solid State Drive),
	//  "pd-ssd" (Persistent Disk Solid State Drive),
	//  or "pd-standard" (Persistent Disk Hard Disk Drive).
	//  See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_type
	BootDiskType *string `json:"bootDiskType,omitempty"`

	// Optional. Size in GB of the boot disk (default is 500GB).
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_size_gb
	BootDiskSizeGB *int32 `json:"bootDiskSizeGB,omitempty"`

	// Optional. Number of attached SSDs, from 0 to 8 (default is 0).
	//  If SSDs are not attached, the boot disk is used to store runtime logs and
	//  [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data.
	//  If one or more SSDs are attached, this runtime bulk
	//  data is spread across them, and the boot disk contains only basic
	//  config and installed binaries.
	//
	//  Note: Local SSD options may vary by machine type and number of vCPUs
	//  selected.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.num_local_ssds
	NumLocalSsds *int32 `json:"numLocalSsds,omitempty"`

	// Optional. Interface type of local SSDs (default is "scsi").
	//  Valid values: "scsi" (Small Computer System Interface),
	//  "nvme" (Non-Volatile Memory Express).
	//  See [local SSD
	//  performance](https://cloud.google.com/compute/docs/disks/local-ssd#performance).
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.local_ssd_interface
	LocalSsdInterface *string `json:"localSsdInterface,omitempty"`

	// Optional. Indicates how many IOPS to provision for the disk. This sets the
	//  number of I/O operations per second that the disk can handle. Note: This
	//  field is only supported if boot_disk_type is hyperdisk-balanced.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_provisioned_iops
	BootDiskProvisionedIops *int64 `json:"bootDiskProvisionedIops,omitempty"`

	// Optional. Indicates how much throughput to provision for the disk. This
	//  sets the number of throughput mb per second that the disk can handle.
	//  Values must be greater than or equal to 1. Note: This field is only
	//  supported if boot_disk_type is hyperdisk-balanced.
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_provisioned_throughput
	BootDiskProvisionedThroughput *int64 `json:"bootDiskProvisionedThroughput,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.EncryptionConfig
type EncryptionConfig struct {
	// Optional. The Cloud KMS key resource name to use for persistent disk
	//  encryption for all instances in the cluster. See [Use CMEK with cluster
	//  data]
	//  (https://cloud.google.com//dataproc/docs/concepts/configuring-clusters/customer-managed-encryption#use_cmek_with_cluster_data)
	//  for more information.
	// +kcc:proto:field=google.cloud.dataproc.v1.EncryptionConfig.gce_pd_kms_key_name
	GCEPDKMSKeyName *string `json:"gcePDKMSKeyName,omitempty"`

	// Optional. The Cloud KMS key resource name to use for cluster persistent
	//  disk and job argument encryption. See [Use CMEK with cluster data]
	//  (https://cloud.google.com//dataproc/docs/concepts/configuring-clusters/customer-managed-encryption#use_cmek_with_cluster_data)
	//  for more information.
	//
	//  When this key resource name is provided, the following job arguments of
	//  the following job types submitted to the cluster are encrypted using CMEK:
	//
	//  * [FlinkJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/FlinkJob)
	//  * [HadoopJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/HadoopJob)
	//  * [SparkJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/SparkJob)
	//  * [SparkRJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/SparkRJob)
	//  * [PySparkJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/PySparkJob)
	//  * [SparkSqlJob](https://cloud.google.com/dataproc/docs/reference/rest/v1/SparkSqlJob)
	//    scriptVariables and queryList.queries
	//  * [HiveJob](https://cloud.google.com/dataproc/docs/reference/rest/v1/HiveJob)
	//    scriptVariables and queryList.queries
	//  * [PigJob](https://cloud.google.com/dataproc/docs/reference/rest/v1/PigJob)
	//    scriptVariables and queryList.queries
	//  * [PrestoJob](https://cloud.google.com/dataproc/docs/reference/rest/v1/PrestoJob)
	//    scriptVariables and queryList.queries
	// +kcc:proto:field=google.cloud.dataproc.v1.EncryptionConfig.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.EndpointConfig
type EndpointConfig struct {

	// Optional. If true, enable http access to specific ports on the cluster
	//  from external sources. Defaults to false.
	// +kcc:proto:field=google.cloud.dataproc.v1.EndpointConfig.enable_http_port_access
	EnableHTTPPortAccess *bool `json:"enableHTTPPortAccess,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.FlinkJob
type FlinkJob struct {
	// The HCFS URI of the jar file that contains the main class.
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.main_jar_file_uri
	MainJarFileURI *string `json:"mainJarFileURI,omitempty"`

	// The name of the driver's main class. The jar file that contains the class
	//  must be in the default CLASSPATH or specified in
	//  [jarFileUris][google.cloud.dataproc.v1.FlinkJob.jar_file_uris].
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.main_class
	MainClass *string `json:"mainClass,omitempty"`

	// Optional. The arguments to pass to the driver. Do not include arguments,
	//  such as `--conf`, that can be set as job properties, since a collision
	//  might occur that causes an incorrect job submission.
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.args
	Args []string `json:"args,omitempty"`

	// Optional. HCFS URIs of jar files to add to the CLASSPATHs of the
	//  Flink driver and tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.jar_file_uris
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// Optional. HCFS URI of the savepoint, which contains the last saved progress
	//  for starting the current job.
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.savepoint_uri
	SavepointURI *string `json:"savepointURI,omitempty"`

	// Optional. A mapping of property names to values, used to configure Flink.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in
	//  `/etc/flink/conf/flink-defaults.conf` and classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.FlinkJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GceClusterConfig
type GCEClusterConfig struct {
	// Optional. The Compute Engine zone where the Dataproc cluster will be
	//  located. If omitted, the service will pick a zone in the cluster's Compute
	//  Engine region. On a get request, zone will always be present.
	//
	//  A full URL, partial URI, or short name are valid. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]`
	//  * `projects/[project_id]/zones/[zone]`
	//  * `[zone]`
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.zone_uri
	ZoneURI *string `json:"zoneURI,omitempty"`

	// Optional. The Compute Engine network to be used for machine
	//  communications. Cannot be specified with subnetwork_uri. If neither
	//  `network_uri` nor `subnetwork_uri` is specified, the "default" network of
	//  the project is used, if it exists. Cannot be a "Custom Subnet Network" (see
	//  [Using Subnetworks](https://cloud.google.com/compute/docs/subnetworks) for
	//  more information).
	//
	//  A full URL, partial URI, or short name are valid. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/global/networks/default`
	//  * `projects/[project_id]/global/networks/default`
	//  * `default`
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`

	// Optional. The Compute Engine subnetwork to be used for machine
	//  communications. Cannot be specified with network_uri.
	//
	//  A full URL, partial URI, or short name are valid. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/[region]/subnetworks/sub0`
	//  * `projects/[project_id]/regions/[region]/subnetworks/sub0`
	//  * `sub0`
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.subnetwork_uri
	SubnetworkURI *string `json:"subnetworkURI,omitempty"`

	// Optional. This setting applies to subnetwork-enabled networks. It is set to
	//  `true` by default in clusters created with image versions 2.2.x.
	//
	//  When set to `true`:
	//
	//  * All cluster VMs have internal IP addresses.
	//  * [Google Private Access]
	//  (https://cloud.google.com/vpc/docs/private-google-access)
	//  must be enabled to access Dataproc and other Google Cloud APIs.
	//  * Off-cluster dependencies must be configured to be accessible
	//  without external IP addresses.
	//
	//  When set to `false`:
	//
	//  * Cluster VMs are not restricted to internal IP addresses.
	//  * Ephemeral external IP addresses are assigned to each cluster VM.
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.internal_ip_only
	InternalIPOnly *bool `json:"internalIPOnly,omitempty"`

	// Optional. The type of IPv6 access for a cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.private_ipv6_google_access
	PrivateIPV6GoogleAccess *string `json:"privateIPV6GoogleAccess,omitempty"`

	// Optional. The [Dataproc service
	//  account](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/service-accounts#service_accounts_in_dataproc)
	//  (also see [VM Data Plane
	//  identity](https://cloud.google.com/dataproc/docs/concepts/iam/dataproc-principals#vm_service_account_data_plane_identity))
	//  used by Dataproc cluster VM instances to access Google Cloud Platform
	//  services.
	//
	//  If not specified, the
	//  [Compute Engine default service
	//  account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account)
	//  is used.
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. The URIs of service account scopes to be included in
	//  Compute Engine instances. The following base set of scopes is always
	//  included:
	//
	//  * https://www.googleapis.com/auth/cloud.useraccounts.readonly
	//  * https://www.googleapis.com/auth/devstorage.read_write
	//  * https://www.googleapis.com/auth/logging.write
	//
	//  If no scopes are specified, the following defaults are also provided:
	//
	//  * https://www.googleapis.com/auth/bigquery
	//  * https://www.googleapis.com/auth/bigtable.admin.table
	//  * https://www.googleapis.com/auth/bigtable.data
	//  * https://www.googleapis.com/auth/devstorage.full_control
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.service_account_scopes
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty"`

	// The Compute Engine network tags to add to all instances (see [Tagging
	//  instances](https://cloud.google.com/vpc/docs/add-remove-network-tags)).
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.tags
	Tags []string `json:"tags,omitempty"`

	// Optional. The Compute Engine metadata entries to add to all instances (see
	//  [Project and instance
	//  metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Optional. Reservation Affinity for consuming Zonal reservation.
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.reservation_affinity
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`

	// Optional. Node Group Affinity for sole-tenant clusters.
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.node_group_affinity
	NodeGroupAffinity *NodeGroupAffinity `json:"nodeGroupAffinity,omitempty"`

	// Optional. Shielded Instance Config for clusters using [Compute Engine
	//  Shielded
	//  VMs](https://cloud.google.com/security/shielded-cloud/shielded-vm).
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.shielded_instance_config
	ShieldedInstanceConfig *ShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	// Optional. Confidential Instance Config for clusters using [Confidential
	//  VMs](https://cloud.google.com/compute/confidential-vm/docs).
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.confidential_instance_config
	ConfidentialInstanceConfig *ConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.HadoopJob
type HadoopJob struct {
	// The HCFS URI of the jar file containing the main class.
	//  Examples:
	//      'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar'
	//      'hdfs:/tmp/test-samples/custom-wordcount.jar'
	//      'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.main_jar_file_uri
	MainJarFileURI *string `json:"mainJarFileURI,omitempty"`

	// The name of the driver's main class. The jar file containing the class
	//  must be in the default CLASSPATH or specified in `jar_file_uris`.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.main_class
	MainClass *string `json:"mainClass,omitempty"`

	// Optional. The arguments to pass to the driver. Do not
	//  include arguments, such as `-libjars` or `-Dfoo=bar`, that can be set as
	//  job properties, since a collision might occur that causes an incorrect job
	//  submission.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.args
	Args []string `json:"args,omitempty"`

	// Optional. Jar file URIs to add to the CLASSPATHs of the
	//  Hadoop driver and tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.jar_file_uris
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// Optional. HCFS (Hadoop Compatible Filesystem) URIs of files to be copied
	//  to the working directory of Hadoop drivers and distributed tasks. Useful
	//  for naively parallel tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.file_uris
	FileUris []string `json:"fileUris,omitempty"`

	// Optional. HCFS URIs of archives to be extracted in the working directory of
	//  Hadoop drivers and tasks. Supported file types:
	//  .jar, .tar, .tar.gz, .tgz, or .zip.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.archive_uris
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// Optional. A mapping of property names to values, used to configure Hadoop.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in `/etc/hadoop/conf/*-site` and
	//  classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.HadoopJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.HiveJob
type HiveJob struct {
	// The HCFS URI of the script that contains Hive queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.HiveJob.query_file_uri
	QueryFileURI *string `json:"queryFileURI,omitempty"`

	// A list of queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.HiveJob.query_list
	QueryList *QueryList `json:"queryList,omitempty"`

	// Optional. Whether to continue executing queries if a query fails.
	//  The default value is `false`. Setting to `true` can be useful when
	//  executing independent parallel queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.HiveJob.continue_on_failure
	ContinueOnFailure *bool `json:"continueOnFailure,omitempty"`

	// Optional. Mapping of query variable names to values (equivalent to the
	//  Hive command: `SET name="value";`).
	// +kcc:proto:field=google.cloud.dataproc.v1.HiveJob.script_variables
	ScriptVariables map[string]string `json:"scriptVariables,omitempty"`

	// Optional. A mapping of property names and values, used to configure Hive.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in `/etc/hadoop/conf/*-site.xml`,
	//  /etc/hive/conf/hive-site.xml, and classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.HiveJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. HCFS URIs of jar files to add to the CLASSPATH of the
	//  Hive server and Hadoop MapReduce (MR) tasks. Can contain Hive SerDes
	//  and UDFs.
	// +kcc:proto:field=google.cloud.dataproc.v1.HiveJob.jar_file_uris
	JarFileUris []string `json:"jarFileUris,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.IdentityConfig
type IdentityConfig struct {
	// Required. Map of user to service account.
	// +kcc:proto:field=google.cloud.dataproc.v1.IdentityConfig.user_service_account_mapping
	UserServiceAccountMapping map[string]string `json:"userServiceAccountMapping,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy
type InstanceFlexibilityPolicy struct {
	// Optional. Defines how the Group selects the provisioning model to ensure
	//  required reliability.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.provisioning_model_mix
	ProvisioningModelMix *InstanceFlexibilityPolicy_ProvisioningModelMix `json:"provisioningModelMix,omitempty"`

	// Optional. List of instance selection options that the group will use when
	//  creating new VMs.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.instance_selection_list
	InstanceSelectionList []InstanceFlexibilityPolicy_InstanceSelection `json:"instanceSelectionList,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelection
type InstanceFlexibilityPolicy_InstanceSelection struct {
	// Optional. Full machine-type names, e.g. "n1-standard-16".
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelection.machine_types
	MachineTypes []string `json:"machineTypes,omitempty"`

	// Optional. Preference of this instance selection. Lower number means
	//  higher preference. Dataproc will first try to create a VM based on the
	//  machine-type with priority rank and fallback to next rank based on
	//  availability. Machine types and instance selections with the same
	//  priority have the same preference.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelection.rank
	Rank *int32 `json:"rank,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult
type InstanceFlexibilityPolicy_InstanceSelectionResult struct {
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.ProvisioningModelMix
type InstanceFlexibilityPolicy_ProvisioningModelMix struct {
	// Optional. The base capacity that will always use Standard VMs to avoid
	//  risk of more preemption than the minimum capacity you need. Dataproc will
	//  create only standard VMs until it reaches standard_capacity_base, then it
	//  will start using standard_capacity_percent_above_base to mix Spot with
	//  Standard VMs. eg. If 15 instances are requested and
	//  standard_capacity_base is 5, Dataproc will create 5 standard VMs and then
	//  start mixing spot and standard VMs for remaining 10 instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.ProvisioningModelMix.standard_capacity_base
	StandardCapacityBase *int32 `json:"standardCapacityBase,omitempty"`

	// Optional. The percentage of target capacity that should use Standard VM.
	//  The remaining percentage will use Spot VMs. The percentage applies only
	//  to the capacity above standard_capacity_base. eg. If 15 instances are
	//  requested and standard_capacity_base is 5 and
	//  standard_capacity_percent_above_base is 30, Dataproc will create 5
	//  standard VMs and then start mixing spot and standard VMs for remaining 10
	//  instances. The mix will be 30% standard and 70% spot.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.ProvisioningModelMix.standard_capacity_percent_above_base
	StandardCapacityPercentAboveBase *int32 `json:"standardCapacityPercentAboveBase,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceGroupConfig
type InstanceGroupConfig struct {
	// Optional. The number of VM instances in the instance group.
	//  For [HA
	//  cluster](/dataproc/docs/concepts/configuring-clusters/high-availability)
	//  [master_config](#FIELDS.master_config) groups, **must be set to 3**.
	//  For standard cluster [master_config](#FIELDS.master_config) groups,
	//  **must be set to 1**.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.num_instances
	NumInstances *int32 `json:"numInstances,omitempty"`

	// Optional. The Compute Engine image resource used for cluster instances.
	//
	//  The URI can represent an image or image family.
	//
	//  Image examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/global/images/[image-id]`
	//  * `projects/[project_id]/global/images/[image-id]`
	//  * `image-id`
	//
	//  Image family examples. Dataproc will use the most recent
	//  image from the family:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/global/images/family/[custom-image-family-name]`
	//  * `projects/[project_id]/global/images/family/[custom-image-family-name]`
	//
	//  If the URI is unspecified, it will be inferred from
	//  `SoftwareConfig.image_version` or the system default.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.image_uri
	ImageURI *string `json:"imageURI,omitempty"`

	// Optional. The Compute Engine machine type used for cluster instances.
	//
	//  A full URL, partial URI, or short name are valid. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]/machineTypes/n1-standard-2`
	//  * `projects/[project_id]/zones/[zone]/machineTypes/n1-standard-2`
	//  * `n1-standard-2`
	//
	//  **Auto Zone Exception**: If you are using the Dataproc
	//  [Auto Zone
	//  Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
	//  feature, you must use the short name of the machine type
	//  resource, for example, `n1-standard-2`.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.machine_type_uri
	MachineTypeURI *string `json:"machineTypeURI,omitempty"`

	// Optional. Disk option config settings.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.disk_config
	DiskConfig *DiskConfig `json:"diskConfig,omitempty"`

	// Optional. Specifies the preemptibility of the instance group.
	//
	//  The default value for master and worker groups is
	//  `NON_PREEMPTIBLE`. This default cannot be changed.
	//
	//  The default value for secondary instances is
	//  `PREEMPTIBLE`.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.preemptibility
	Preemptibility *string `json:"preemptibility,omitempty"`

	// Optional. The Compute Engine accelerator configuration for these
	//  instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.accelerators
	Accelerators []AcceleratorConfig `json:"accelerators,omitempty"`

	// Optional. Specifies the minimum cpu platform for the Instance Group.
	//  See [Dataproc -> Minimum CPU
	//  Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.min_cpu_platform
	MinCPUPlatform *string `json:"minCPUPlatform,omitempty"`

	// Optional. The minimum number of primary worker instances to create.
	//  If `min_num_instances` is set, cluster creation will succeed if
	//  the number of primary workers created is at least equal to the
	//  `min_num_instances` number.
	//
	//  Example: Cluster creation request with `num_instances` = `5` and
	//  `min_num_instances` = `3`:
	//
	//  *  If 4 VMs are created and 1 instance fails,
	//     the failed VM is deleted. The cluster is
	//     resized to 4 instances and placed in a `RUNNING` state.
	//  *  If 2 instances are created and 3 instances fail,
	//     the cluster in placed in an `ERROR` state. The failed VMs
	//     are not deleted.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.min_num_instances
	MinNumInstances *int32 `json:"minNumInstances,omitempty"`

	// Optional. Instance flexibility Policy allowing a mixture of VM shapes and
	//  provisioning models.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_flexibility_policy
	InstanceFlexibilityPolicy *InstanceFlexibilityPolicy `json:"instanceFlexibilityPolicy,omitempty"`

	// Optional. Configuration to handle the startup of instances during cluster
	//  create and update process.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.startup_config
	StartupConfig *StartupConfig `json:"startupConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceReference
type InstanceReference struct {
	// The user-friendly name of the Compute Engine instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.instance_name
	InstanceName *string `json:"instanceName,omitempty"`

	// The unique identifier of the Compute Engine instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.instance_id
	InstanceID *string `json:"instanceID,omitempty"`

	// The public RSA key used for sharing data with this instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.public_key
	PublicKey *string `json:"publicKey,omitempty"`

	// The public ECIES key used for sharing data with this instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.public_ecies_key
	PublicEciesKey *string `json:"publicEciesKey,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.JobScheduling
type JobScheduling struct {
	// Optional. Maximum number of times per hour a driver can be restarted as
	//  a result of driver exiting with non-zero code before job is
	//  reported failed.
	//
	//  A job might be reported as thrashing if the driver exits with a non-zero
	//  code four times within a 10-minute window.
	//
	//  Maximum value is 10.
	//
	//  **Note:** This restartable job option is not supported in Dataproc
	//  [workflow templates]
	//  (https://cloud.google.com/dataproc/docs/concepts/workflows/using-workflows#adding_jobs_to_a_template).
	// +kcc:proto:field=google.cloud.dataproc.v1.JobScheduling.max_failures_per_hour
	MaxFailuresPerHour *int32 `json:"maxFailuresPerHour,omitempty"`

	// Optional. Maximum total number of times a driver can be restarted as a
	//  result of the driver exiting with a non-zero code. After the maximum number
	//  is reached, the job will be reported as failed.
	//
	//  Maximum value is 240.
	//
	//  **Note:** Currently, this restartable job option is
	//  not supported in Dataproc
	//  [workflow
	//  templates](https://cloud.google.com/dataproc/docs/concepts/workflows/using-workflows#adding_jobs_to_a_template).
	// +kcc:proto:field=google.cloud.dataproc.v1.JobScheduling.max_failures_total
	MaxFailuresTotal *int32 `json:"maxFailuresTotal,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.KerberosConfig
type KerberosConfig struct {
	// Optional. Flag to indicate whether to Kerberize the cluster (default:
	//  false). Set this field to true to enable Kerberos on a cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.enable_kerberos
	EnableKerberos *bool `json:"enableKerberos,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the root
	//  principal password.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.root_principal_password_uri
	RootPrincipalPasswordURI *string `json:"rootPrincipalPasswordURI,omitempty"`

	// Optional. The URI of the KMS key used to encrypt sensitive
	//  files.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.kms_key_uri
	KMSKeyURI *string `json:"kmsKeyURI,omitempty"`

	// Optional. The Cloud Storage URI of the keystore file used for SSL
	//  encryption. If not provided, Dataproc will provide a self-signed
	//  certificate.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.keystore_uri
	KeystoreURI *string `json:"keystoreURI,omitempty"`

	// Optional. The Cloud Storage URI of the truststore file used for SSL
	//  encryption. If not provided, Dataproc will provide a self-signed
	//  certificate.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.truststore_uri
	TruststoreURI *string `json:"truststoreURI,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  password to the user provided keystore. For the self-signed certificate,
	//  this password is generated by Dataproc.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.keystore_password_uri
	KeystorePasswordURI *string `json:"keystorePasswordURI,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  password to the user provided key. For the self-signed certificate, this
	//  password is generated by Dataproc.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.key_password_uri
	KeyPasswordURI *string `json:"keyPasswordURI,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  password to the user provided truststore. For the self-signed certificate,
	//  this password is generated by Dataproc.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.truststore_password_uri
	TruststorePasswordURI *string `json:"truststorePasswordURI,omitempty"`

	// Optional. The remote realm the Dataproc on-cluster KDC will trust, should
	//  the user enable cross realm trust.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.cross_realm_trust_realm
	CrossRealmTrustRealm *string `json:"crossRealmTrustRealm,omitempty"`

	// Optional. The KDC (IP or hostname) for the remote trusted realm in a cross
	//  realm trust relationship.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.cross_realm_trust_kdc
	CrossRealmTrustKdc *string `json:"crossRealmTrustKdc,omitempty"`

	// Optional. The admin server (IP or hostname) for the remote trusted realm in
	//  a cross realm trust relationship.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.cross_realm_trust_admin_server
	CrossRealmTrustAdminServer *string `json:"crossRealmTrustAdminServer,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  shared password between the on-cluster Kerberos realm and the remote
	//  trusted realm, in a cross realm trust relationship.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.cross_realm_trust_shared_password_uri
	CrossRealmTrustSharedPasswordURI *string `json:"crossRealmTrustSharedPasswordURI,omitempty"`

	// Optional. The Cloud Storage URI of a KMS encrypted file containing the
	//  master key of the KDC database.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.kdc_db_key_uri
	KdcDbKeyURI *string `json:"kdcDbKeyURI,omitempty"`

	// Optional. The lifetime of the ticket granting ticket, in hours.
	//  If not specified, or user specifies 0, then default value 10
	//  will be used.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.tgt_lifetime_hours
	TgtLifetimeHours *int32 `json:"tgtLifetimeHours,omitempty"`

	// Optional. The name of the on-cluster Kerberos realm.
	//  If not specified, the uppercased domain of hostnames will be the realm.
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.realm
	Realm *string `json:"realm,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.LifecycleConfig
type LifecycleConfig struct {
	// Optional. The duration to keep the cluster alive while idling (when no jobs
	//  are running). Passing this threshold will cause the cluster to be
	//  deleted. Minimum value is 5 minutes; maximum value is 14 days (see JSON
	//  representation of
	//  [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	// +kcc:proto:field=google.cloud.dataproc.v1.LifecycleConfig.idle_delete_ttl
	IdleDeleteTtl *string `json:"idleDeleteTtl,omitempty"`

	// Optional. The time when cluster will be auto-deleted (see JSON
	//  representation of
	//  [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	// +kcc:proto:field=google.cloud.dataproc.v1.LifecycleConfig.auto_delete_time
	AutoDeleteTime *string `json:"autoDeleteTime,omitempty"`

	// Optional. The lifetime duration of cluster. The cluster will be
	//  auto-deleted at the end of this period. Minimum value is 10 minutes;
	//  maximum value is 14 days (see JSON representation of
	//  [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	// +kcc:proto:field=google.cloud.dataproc.v1.LifecycleConfig.auto_delete_ttl
	AutoDeleteTtl *string `json:"autoDeleteTtl,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.LoggingConfig
type LoggingConfig struct {

	// TODO: unsupported map type with key string and value enum

}

// +kcc:proto=google.cloud.dataproc.v1.ManagedCluster
type ManagedCluster struct {
	// Required. The cluster name prefix. A unique cluster name will be formed by
	//  appending a random suffix.
	//
	//  The name must contain only lower-case letters (a-z), numbers (0-9),
	//  and hyphens (-). Must begin with a letter. Cannot begin or end with
	//  hyphen. Must consist of between 2 and 35 characters.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedCluster.cluster_name
	ClusterName *string `json:"clusterName,omitempty"`

	// Required. The cluster configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedCluster.config
	Config *ClusterConfig `json:"config,omitempty"`

	// Optional. The labels to associate with this cluster.
	//
	//  Label keys must be between 1 and 63 characters long, and must conform to
	//  the following PCRE regular expression:
	//  [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	//
	//  Label values must be between 1 and 63 characters long, and must conform to
	//  the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	//
	//  No more than 32 labels can be associated with a given cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedCluster.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ManagedGroupConfig
type ManagedGroupConfig struct {
}

// +kcc:proto=google.cloud.dataproc.v1.MetastoreConfig
type MetastoreConfig struct {
	// Required. Resource name of an existing Dataproc Metastore service.
	//
	//  Example:
	//
	//  * `projects/[project_id]/locations/[dataproc_region]/services/[service-name]`
	// +kcc:proto:field=google.cloud.dataproc.v1.MetastoreConfig.dataproc_metastore_service
	DataprocMetastoreService *string `json:"dataprocMetastoreService,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.NodeGroup
type NodeGroup struct {
	// The Node group [resource name](https://aip.dev/122).
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroup.name
	Name *string `json:"name,omitempty"`

	// Required. Node group roles.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroup.roles
	Roles []string `json:"roles,omitempty"`

	// Optional. The node group instance group configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroup.node_group_config
	NodeGroupConfig *InstanceGroupConfig `json:"nodeGroupConfig,omitempty"`

	// Optional. Node group labels.
	//
	//  * Label **keys** must consist of from 1 to 63 characters and conform to
	//    [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  * Label **values** can be empty. If specified, they must consist of from
	//    1 to 63 characters and conform to [RFC 1035]
	//    (https://www.ietf.org/rfc/rfc1035.txt).
	//  * The node group must have no more than 32 labels.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroup.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.NodeGroupAffinity
type NodeGroupAffinity struct {
	// Required. The URI of a
	//  sole-tenant [node group
	//  resource](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups)
	//  that the cluster will be created on.
	//
	//  A full URL, partial URI, or node group name are valid. Examples:
	//
	//  * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]/nodeGroups/node-group-1`
	//  * `projects/[project_id]/zones/[zone]/nodeGroups/node-group-1`
	//  * `node-group-1`
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroupAffinity.node_group_uri
	NodeGroupURI *string `json:"nodeGroupURI,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.NodeInitializationAction
type NodeInitializationAction struct {
	// Required. Cloud Storage URI of executable file.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeInitializationAction.executable_file
	ExecutableFile *string `json:"executableFile,omitempty"`

	// Optional. Amount of time executable has to complete. Default is
	//  10 minutes (see JSON representation of
	//  [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	//
	//  Cluster creation fails with an explanatory error message (the
	//  name of the executable that caused the error and the exceeded timeout
	//  period) if the executable is not completed at end of the timeout period.
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeInitializationAction.execution_timeout
	ExecutionTimeout *string `json:"executionTimeout,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.OrderedJob
type OrderedJob struct {
	// Required. The step id. The id must be unique among all jobs
	//  within the template.
	//
	//  The step id is used as prefix for job id, as job
	//  `goog-dataproc-workflow-step-id` label, and in
	//  [prerequisiteStepIds][google.cloud.dataproc.v1.OrderedJob.prerequisite_step_ids]
	//  field from other steps.
	//
	//  The id must contain only letters (a-z, A-Z), numbers (0-9),
	//  underscores (_), and hyphens (-). Cannot begin or end with underscore
	//  or hyphen. Must consist of between 3 and 50 characters.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.step_id
	StepID *string `json:"stepID,omitempty"`

	// Optional. Job is a Hadoop job.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.hadoop_job
	HadoopJob *HadoopJob `json:"hadoopJob,omitempty"`

	// Optional. Job is a Spark job.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.spark_job
	SparkJob *SparkJob `json:"sparkJob,omitempty"`

	// Optional. Job is a PySpark job.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.pyspark_job
	PysparkJob *PySparkJob `json:"pysparkJob,omitempty"`

	// Optional. Job is a Hive job.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.hive_job
	HiveJob *HiveJob `json:"hiveJob,omitempty"`

	// Optional. Job is a Pig job.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.pig_job
	PigJob *PigJob `json:"pigJob,omitempty"`

	// Optional. Job is a SparkR job.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.spark_r_job
	SparkRJob *SparkRJob `json:"sparkRJob,omitempty"`

	// Optional. Job is a SparkSql job.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.spark_sql_job
	SparkSQLJob *SparkSQLJob `json:"sparkSQLJob,omitempty"`

	// Optional. Job is a Presto job.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.presto_job
	PrestoJob *PrestoJob `json:"prestoJob,omitempty"`

	// Optional. Job is a Trino job.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.trino_job
	TrinoJob *TrinoJob `json:"trinoJob,omitempty"`

	// Optional. Job is a Flink job.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.flink_job
	FlinkJob *FlinkJob `json:"flinkJob,omitempty"`

	// Optional. The labels to associate with this job.
	//
	//  Label keys must be between 1 and 63 characters long, and must conform to
	//  the following regular expression:
	//  [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	//
	//  Label values must be between 1 and 63 characters long, and must conform to
	//  the following regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	//
	//  No more than 32 labels can be associated with a given job.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Job scheduling configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.scheduling
	Scheduling *JobScheduling `json:"scheduling,omitempty"`

	// Optional. The optional list of prerequisite job step_ids.
	//  If not specified, the job will start at the beginning of workflow.
	// +kcc:proto:field=google.cloud.dataproc.v1.OrderedJob.prerequisite_step_ids
	PrerequisiteStepIds []string `json:"prerequisiteStepIds,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ParameterValidation
type ParameterValidation struct {
	// Validation based on regular expressions.
	// +kcc:proto:field=google.cloud.dataproc.v1.ParameterValidation.regex
	Regex *RegexValidation `json:"regex,omitempty"`

	// Validation based on a list of allowed values.
	// +kcc:proto:field=google.cloud.dataproc.v1.ParameterValidation.values
	Values *ValueValidation `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.PigJob
type PigJob struct {
	// The HCFS URI of the script that contains the Pig queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.query_file_uri
	QueryFileURI *string `json:"queryFileURI,omitempty"`

	// A list of queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.query_list
	QueryList *QueryList `json:"queryList,omitempty"`

	// Optional. Whether to continue executing queries if a query fails.
	//  The default value is `false`. Setting to `true` can be useful when
	//  executing independent parallel queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.continue_on_failure
	ContinueOnFailure *bool `json:"continueOnFailure,omitempty"`

	// Optional. Mapping of query variable names to values (equivalent to the Pig
	//  command: `name=[value]`).
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.script_variables
	ScriptVariables map[string]string `json:"scriptVariables,omitempty"`

	// Optional. A mapping of property names to values, used to configure Pig.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in `/etc/hadoop/conf/*-site.xml`,
	//  /etc/pig/conf/pig.properties, and classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. HCFS URIs of jar files to add to the CLASSPATH of
	//  the Pig Client and Hadoop MapReduce (MR) tasks. Can contain Pig UDFs.
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.jar_file_uris
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.PigJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.PrestoJob
type PrestoJob struct {
	// The HCFS URI of the script that contains SQL queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.query_file_uri
	QueryFileURI *string `json:"queryFileURI,omitempty"`

	// A list of queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.query_list
	QueryList *QueryList `json:"queryList,omitempty"`

	// Optional. Whether to continue executing queries if a query fails.
	//  The default value is `false`. Setting to `true` can be useful when
	//  executing independent parallel queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.continue_on_failure
	ContinueOnFailure *bool `json:"continueOnFailure,omitempty"`

	// Optional. The format in which query output will be displayed. See the
	//  Presto documentation for supported output formats
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.output_format
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Optional. Presto client tags to attach to this query
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.client_tags
	ClientTags []string `json:"clientTags,omitempty"`

	// Optional. A mapping of property names to values. Used to set Presto
	//  [session properties](https://prestodb.io/docs/current/sql/set-session.html)
	//  Equivalent to using the --session flag in the Presto CLI
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.PrestoJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.PySparkJob
type PySparkJob struct {
	// Required. The HCFS URI of the main Python file to use as the driver. Must
	//  be a .py file.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.main_python_file_uri
	MainPythonFileURI *string `json:"mainPythonFileURI,omitempty"`

	// Optional. The arguments to pass to the driver.  Do not include arguments,
	//  such as `--conf`, that can be set as job properties, since a collision may
	//  occur that causes an incorrect job submission.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.args
	Args []string `json:"args,omitempty"`

	// Optional. HCFS file URIs of Python files to pass to the PySpark
	//  framework. Supported file types: .py, .egg, and .zip.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.python_file_uris
	PythonFileUris []string `json:"pythonFileUris,omitempty"`

	// Optional. HCFS URIs of jar files to add to the CLASSPATHs of the
	//  Python driver and tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.jar_file_uris
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// Optional. HCFS URIs of files to be placed in the working directory of
	//  each executor. Useful for naively parallel tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.file_uris
	FileUris []string `json:"fileUris,omitempty"`

	// Optional. HCFS URIs of archives to be extracted into the working directory
	//  of each executor. Supported file types:
	//  .jar, .tar, .tar.gz, .tgz, and .zip.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.archive_uris
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// Optional. A mapping of property names to values, used to configure PySpark.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in
	//  /etc/spark/conf/spark-defaults.conf and classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.PySparkJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.QueryList
type QueryList struct {
	// Required. The queries to execute. You do not need to end a query expression
	//  with a semicolon. Multiple queries can be specified in one
	//  string by separating each with a semicolon. Here is an example of a
	//  Dataproc API snippet that uses a QueryList to specify a HiveJob:
	//
	//      "hiveJob": {
	//        "queryList": {
	//          "queries": [
	//            "query1",
	//            "query2",
	//            "query3;query4",
	//          ]
	//        }
	//      }
	// +kcc:proto:field=google.cloud.dataproc.v1.QueryList.queries
	Queries []string `json:"queries,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.RegexValidation
type RegexValidation struct {
	// Required. RE2 regular expressions used to validate the parameter's value.
	//  The value must match the regex in its entirety (substring
	//  matches are not sufficient).
	// +kcc:proto:field=google.cloud.dataproc.v1.RegexValidation.regexes
	Regexes []string `json:"regexes,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ReservationAffinity
type ReservationAffinity struct {
	// Optional. Type of reservation to consume
	// +kcc:proto:field=google.cloud.dataproc.v1.ReservationAffinity.consume_reservation_type
	ConsumeReservationType *string `json:"consumeReservationType,omitempty"`

	// Optional. Corresponds to the label key of reservation resource.
	// +kcc:proto:field=google.cloud.dataproc.v1.ReservationAffinity.key
	Key *string `json:"key,omitempty"`

	// Optional. Corresponds to the label values of reservation resource.
	// +kcc:proto:field=google.cloud.dataproc.v1.ReservationAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SecurityConfig
type SecurityConfig struct {
	// Optional. Kerberos related configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.SecurityConfig.kerberos_config
	KerberosConfig *KerberosConfig `json:"kerberosConfig,omitempty"`

	// Optional. Identity related configuration, including service account based
	//  secure multi-tenancy user mappings.
	// +kcc:proto:field=google.cloud.dataproc.v1.SecurityConfig.identity_config
	IdentityConfig *IdentityConfig `json:"identityConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ShieldedInstanceConfig
type ShieldedInstanceConfig struct {
	// Optional. Defines whether instances have Secure Boot enabled.
	// +kcc:proto:field=google.cloud.dataproc.v1.ShieldedInstanceConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	// Optional. Defines whether instances have the vTPM enabled.
	// +kcc:proto:field=google.cloud.dataproc.v1.ShieldedInstanceConfig.enable_vtpm
	EnableVTPM *bool `json:"enableVTPM,omitempty"`

	// Optional. Defines whether instances have integrity monitoring enabled.
	// +kcc:proto:field=google.cloud.dataproc.v1.ShieldedInstanceConfig.enable_integrity_monitoring
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SoftwareConfig
type SoftwareConfig struct {
	// Optional. The version of software inside the cluster. It must be one of the
	//  supported [Dataproc
	//  Versions](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#supported-dataproc-image-versions),
	//  such as "1.2" (including a subminor version, such as "1.2.29"), or the
	//  ["preview"
	//  version](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#other_versions).
	//  If unspecified, it defaults to the latest Debian version.
	// +kcc:proto:field=google.cloud.dataproc.v1.SoftwareConfig.image_version
	ImageVersion *string `json:"imageVersion,omitempty"`

	// Optional. The properties to set on daemon config files.
	//
	//  Property keys are specified in `prefix:property` format, for example
	//  `core:hadoop.tmp.dir`. The following are supported prefixes
	//  and their mappings:
	//
	//  * capacity-scheduler: `capacity-scheduler.xml`
	//  * core:   `core-site.xml`
	//  * distcp: `distcp-default.xml`
	//  * hdfs:   `hdfs-site.xml`
	//  * hive:   `hive-site.xml`
	//  * mapred: `mapred-site.xml`
	//  * pig:    `pig.properties`
	//  * spark:  `spark-defaults.conf`
	//  * yarn:   `yarn-site.xml`
	//
	//  For more information, see [Cluster
	//  properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties).
	// +kcc:proto:field=google.cloud.dataproc.v1.SoftwareConfig.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The set of components to activate on the cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.SoftwareConfig.optional_components
	OptionalComponents []string `json:"optionalComponents,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SparkJob
type SparkJob struct {
	// The HCFS URI of the jar file that contains the main class.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.main_jar_file_uri
	MainJarFileURI *string `json:"mainJarFileURI,omitempty"`

	// The name of the driver's main class. The jar file that contains the class
	//  must be in the default CLASSPATH or specified in
	//  SparkJob.jar_file_uris.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.main_class
	MainClass *string `json:"mainClass,omitempty"`

	// Optional. The arguments to pass to the driver. Do not include arguments,
	//  such as `--conf`, that can be set as job properties, since a collision may
	//  occur that causes an incorrect job submission.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.args
	Args []string `json:"args,omitempty"`

	// Optional. HCFS URIs of jar files to add to the CLASSPATHs of the
	//  Spark driver and tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.jar_file_uris
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// Optional. HCFS URIs of files to be placed in the working directory of
	//  each executor. Useful for naively parallel tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.file_uris
	FileUris []string `json:"fileUris,omitempty"`

	// Optional. HCFS URIs of archives to be extracted into the working directory
	//  of each executor. Supported file types:
	//  .jar, .tar, .tar.gz, .tgz, and .zip.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.archive_uris
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// Optional. A mapping of property names to values, used to configure Spark.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in
	//  /etc/spark/conf/spark-defaults.conf and classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SparkRJob
type SparkRJob struct {
	// Required. The HCFS URI of the main R file to use as the driver.
	//  Must be a .R file.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRJob.main_r_file_uri
	MainRFileURI *string `json:"mainRFileURI,omitempty"`

	// Optional. The arguments to pass to the driver.  Do not include arguments,
	//  such as `--conf`, that can be set as job properties, since a collision may
	//  occur that causes an incorrect job submission.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRJob.args
	Args []string `json:"args,omitempty"`

	// Optional. HCFS URIs of files to be placed in the working directory of
	//  each executor. Useful for naively parallel tasks.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRJob.file_uris
	FileUris []string `json:"fileUris,omitempty"`

	// Optional. HCFS URIs of archives to be extracted into the working directory
	//  of each executor. Supported file types:
	//  .jar, .tar, .tar.gz, .tgz, and .zip.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRJob.archive_uris
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// Optional. A mapping of property names to values, used to configure SparkR.
	//  Properties that conflict with values set by the Dataproc API might be
	//  overwritten. Can include properties set in
	//  /etc/spark/conf/spark-defaults.conf and classes in user code.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkRJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SparkSqlJob
type SparkSQLJob struct {
	// The HCFS URI of the script that contains SQL queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlJob.query_file_uri
	QueryFileURI *string `json:"queryFileURI,omitempty"`

	// A list of queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlJob.query_list
	QueryList *QueryList `json:"queryList,omitempty"`

	// Optional. Mapping of query variable names to values (equivalent to the
	//  Spark SQL command: SET `name="value";`).
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlJob.script_variables
	ScriptVariables map[string]string `json:"scriptVariables,omitempty"`

	// Optional. A mapping of property names to values, used to configure
	//  Spark SQL's SparkConf. Properties that conflict with values set by the
	//  Dataproc API might be overwritten.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. HCFS URIs of jar files to be added to the Spark CLASSPATH.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlJob.jar_file_uris
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkSqlJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.StartupConfig
type StartupConfig struct {
	// Optional. The config setting to enable cluster creation/ updation to be
	//  successful only after required_registration_fraction of instances are up
	//  and running. This configuration is applicable to only secondary workers for
	//  now. The cluster will fail if required_registration_fraction of instances
	//  are not available. This will include instance creation, agent registration,
	//  and service registration (if enabled).
	// +kcc:proto:field=google.cloud.dataproc.v1.StartupConfig.required_registration_fraction
	RequiredRegistrationFraction *float64 `json:"requiredRegistrationFraction,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.TemplateParameter
type TemplateParameter struct {
	// Required. Parameter name.
	//  The parameter name is used as the key, and paired with the
	//  parameter value, which are passed to the template when the template
	//  is instantiated.
	//  The name must contain only capital letters (A-Z), numbers (0-9), and
	//  underscores (_), and must not start with a number. The maximum length is
	//  40 characters.
	// +kcc:proto:field=google.cloud.dataproc.v1.TemplateParameter.name
	Name *string `json:"name,omitempty"`

	// Required. Paths to all fields that the parameter replaces.
	//  A field is allowed to appear in at most one parameter's list of field
	//  paths.
	//
	//  A field path is similar in syntax to a
	//  [google.protobuf.FieldMask][google.protobuf.FieldMask]. For example, a
	//  field path that references the zone field of a workflow template's cluster
	//  selector would be specified as `placement.clusterSelector.zone`.
	//
	//  Also, field paths can reference fields using the following syntax:
	//
	//  * Values in maps can be referenced by key:
	//      * labels['key']
	//      * placement.clusterSelector.clusterLabels['key']
	//      * placement.managedCluster.labels['key']
	//      * placement.clusterSelector.clusterLabels['key']
	//      * jobs['step-id'].labels['key']
	//
	//  * Jobs in the jobs list can be referenced by step-id:
	//      * jobs['step-id'].hadoopJob.mainJarFileUri
	//      * jobs['step-id'].hiveJob.queryFileUri
	//      * jobs['step-id'].pySparkJob.mainPythonFileUri
	//      * jobs['step-id'].hadoopJob.jarFileUris[0]
	//      * jobs['step-id'].hadoopJob.archiveUris[0]
	//      * jobs['step-id'].hadoopJob.fileUris[0]
	//      * jobs['step-id'].pySparkJob.pythonFileUris[0]
	//
	//  * Items in repeated fields can be referenced by a zero-based index:
	//      * jobs['step-id'].sparkJob.args[0]
	//
	//  * Other examples:
	//      * jobs['step-id'].hadoopJob.properties['key']
	//      * jobs['step-id'].hadoopJob.args[0]
	//      * jobs['step-id'].hiveJob.scriptVariables['key']
	//      * jobs['step-id'].hadoopJob.mainJarFileUri
	//      * placement.clusterSelector.zone
	//
	//  It may not be possible to parameterize maps and repeated fields in their
	//  entirety since only individual map values and individual items in repeated
	//  fields can be referenced. For example, the following field paths are
	//  invalid:
	//
	//  - placement.clusterSelector.clusterLabels
	//  - jobs['step-id'].sparkJob.args
	// +kcc:proto:field=google.cloud.dataproc.v1.TemplateParameter.fields
	Fields []string `json:"fields,omitempty"`

	// Optional. Brief description of the parameter.
	//  Must not exceed 1024 characters.
	// +kcc:proto:field=google.cloud.dataproc.v1.TemplateParameter.description
	Description *string `json:"description,omitempty"`

	// Optional. Validation rules to be applied to this parameter's value.
	// +kcc:proto:field=google.cloud.dataproc.v1.TemplateParameter.validation
	Validation *ParameterValidation `json:"validation,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.TrinoJob
type TrinoJob struct {
	// The HCFS URI of the script that contains SQL queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.query_file_uri
	QueryFileURI *string `json:"queryFileURI,omitempty"`

	// A list of queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.query_list
	QueryList *QueryList `json:"queryList,omitempty"`

	// Optional. Whether to continue executing queries if a query fails.
	//  The default value is `false`. Setting to `true` can be useful when
	//  executing independent parallel queries.
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.continue_on_failure
	ContinueOnFailure *bool `json:"continueOnFailure,omitempty"`

	// Optional. The format in which query output will be displayed. See the
	//  Trino documentation for supported output formats
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.output_format
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Optional. Trino client tags to attach to this query
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.client_tags
	ClientTags []string `json:"clientTags,omitempty"`

	// Optional. A mapping of property names to values. Used to set Trino
	//  [session properties](https://trino.io/docs/current/sql/set-session.html)
	//  Equivalent to using the --session flag in the Trino CLI
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.properties
	Properties map[string]string `json:"properties,omitempty"`

	// Optional. The runtime log config for job execution.
	// +kcc:proto:field=google.cloud.dataproc.v1.TrinoJob.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ValueValidation
type ValueValidation struct {
	// Required. List of allowed values for the parameter.
	// +kcc:proto:field=google.cloud.dataproc.v1.ValueValidation.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.WorkflowTemplate
type WorkflowTemplate struct {
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.id
	ID *string `json:"id,omitempty"`

	// Optional. Used to perform a consistent read-modify-write.
	//
	//  This field should be left blank for a `CreateWorkflowTemplate` request. It
	//  is required for an `UpdateWorkflowTemplate` request, and must match the
	//  current server version. A typical update template flow would fetch the
	//  current template with a `GetWorkflowTemplate` request, which will return
	//  the current template with the `version` field filled in with the
	//  current server version. The user updates other fields in the template,
	//  then returns it as part of the `UpdateWorkflowTemplate` request.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.version
	Version *int32 `json:"version,omitempty"`

	// Optional. The labels to associate with this template. These labels
	//  will be propagated to all jobs and clusters created by the workflow
	//  instance.
	//
	//  Label **keys** must contain 1 to 63 characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//
	//  Label **values** may be empty, but, if present, must contain 1 to 63
	//  characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//
	//  No more than 32 labels can be associated with a template.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. WorkflowTemplate scheduling information.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.placement
	Placement *WorkflowTemplatePlacement `json:"placement,omitempty"`

	// Required. The Directed Acyclic Graph of Jobs to submit.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.jobs
	Jobs []OrderedJob `json:"jobs,omitempty"`

	// Optional. Template parameters whose values are substituted into the
	//  template. Values for parameters must be provided when the template is
	//  instantiated.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.parameters
	Parameters []TemplateParameter `json:"parameters,omitempty"`

	// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see
	//  [JSON representation of
	//  duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	//  The timeout duration must be from 10 minutes ("600s") to 24 hours
	//  ("86400s"). The timer begins when the first job is submitted. If the
	//  workflow is running at the end of the timeout period, any remaining jobs
	//  are cancelled, the workflow is ended, and if the workflow was running on a
	//  [managed
	//  cluster](/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster),
	//  the cluster is deleted.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.dag_timeout
	DagTimeout *string `json:"dagTimeout,omitempty"`

	// Optional. Encryption settings for encrypting workflow template job
	//  arguments.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.encryption_config
	EncryptionConfig *WorkflowTemplate_EncryptionConfig `json:"encryptionConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.WorkflowTemplate.EncryptionConfig
type WorkflowTemplate_EncryptionConfig struct {
	// Optional. The Cloud KMS key name to use for encrypting
	//  workflow template job arguments.
	//
	//  When this this key is provided, the following workflow template
	//  [job arguments]
	//  (https://cloud.google.com/dataproc/docs/concepts/workflows/use-workflows#adding_jobs_to_a_template),
	//  if present, are
	//  [CMEK
	//  encrypted](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/customer-managed-encryption#use_cmek_with_workflow_template_data):
	//
	//  * [FlinkJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/FlinkJob)
	//  * [HadoopJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/HadoopJob)
	//  * [SparkJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/SparkJob)
	//  * [SparkRJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/SparkRJob)
	//  * [PySparkJob
	//  args](https://cloud.google.com/dataproc/docs/reference/rest/v1/PySparkJob)
	//  * [SparkSqlJob](https://cloud.google.com/dataproc/docs/reference/rest/v1/SparkSqlJob)
	//    scriptVariables and queryList.queries
	//  * [HiveJob](https://cloud.google.com/dataproc/docs/reference/rest/v1/HiveJob)
	//    scriptVariables and queryList.queries
	//  * [PigJob](https://cloud.google.com/dataproc/docs/reference/rest/v1/PigJob)
	//    scriptVariables and queryList.queries
	//  * [PrestoJob](https://cloud.google.com/dataproc/docs/reference/rest/v1/PrestoJob)
	//    scriptVariables and queryList.queries
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.EncryptionConfig.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.WorkflowTemplatePlacement
type WorkflowTemplatePlacement struct {
	// A cluster that is managed by the workflow.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplatePlacement.managed_cluster
	ManagedCluster *ManagedCluster `json:"managedCluster,omitempty"`

	// Optional. A selector that chooses target cluster for jobs based
	//  on metadata.
	//
	//  The selector is evaluated at the time each job is submitted.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplatePlacement.cluster_selector
	ClusterSelector *ClusterSelector `json:"clusterSelector,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ClusterConfig
type ClusterConfigObservedState struct {
	// Optional. The Compute Engine config settings for
	//  the cluster's master instance.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.master_config
	MasterConfig *InstanceGroupConfigObservedState `json:"masterConfig,omitempty"`

	// Optional. Lifecycle setting for the cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.lifecycle_config
	LifecycleConfig *LifecycleConfigObservedState `json:"lifecycleConfig,omitempty"`

	// Optional. Port/endpoint configuration for this cluster
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.endpoint_config
	EndpointConfig *EndpointConfigObservedState `json:"endpointConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.EndpointConfig
type EndpointConfigObservedState struct {
	// Output only. The map of port descriptions to URLs. Will only be populated
	//  if enable_http_port_access is true.
	// +kcc:proto:field=google.cloud.dataproc.v1.EndpointConfig.http_ports
	HTTPPorts map[string]string `json:"httpPorts,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy
type InstanceFlexibilityPolicyObservedState struct {
	// Output only. A list of instance selection results in the group.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.instance_selection_results
	InstanceSelectionResults []InstanceFlexibilityPolicy_InstanceSelectionResult `json:"instanceSelectionResults,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult
type InstanceFlexibilityPolicy_InstanceSelectionResultObservedState struct {
	// Output only. Full machine-type names, e.g. "n1-standard-16".
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Output only. Number of VM provisioned with the machine_type.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceFlexibilityPolicy.InstanceSelectionResult.vm_count
	VmCount *int32 `json:"vmCount,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceGroupConfig
type InstanceGroupConfigObservedState struct {
	// Output only. The list of instance names. Dataproc derives the names
	//  from `cluster_name`, `num_instances`, and the instance group.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_names
	InstanceNames []string `json:"instanceNames,omitempty"`

	// Output only. List of references to Compute Engine instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_references
	InstanceReferences []InstanceReference `json:"instanceReferences,omitempty"`

	// Output only. Specifies that this instance group contains preemptible
	//  instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.is_preemptible
	IsPreemptible *bool `json:"isPreemptible,omitempty"`

	// Output only. The config for Compute Engine Instance Group
	//  Manager that manages this group.
	//  This is only used for preemptible instance groups.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.managed_group_config
	ManagedGroupConfig *ManagedGroupConfig `json:"managedGroupConfig,omitempty"`

	// Optional. Instance flexibility Policy allowing a mixture of VM shapes and
	//  provisioning models.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_flexibility_policy
	InstanceFlexibilityPolicy *InstanceFlexibilityPolicyObservedState `json:"instanceFlexibilityPolicy,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.LifecycleConfig
type LifecycleConfigObservedState struct {
	// Output only. The time when cluster became idle (most recent job finished)
	//  and became eligible for deletion due to idleness (see JSON representation
	//  of
	//  [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	// +kcc:proto:field=google.cloud.dataproc.v1.LifecycleConfig.idle_start_time
	IdleStartTime *string `json:"idleStartTime,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ManagedCluster
type ManagedClusterObservedState struct {
	// Required. The cluster configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedCluster.config
	Config *ClusterConfigObservedState `json:"config,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ManagedGroupConfig
type ManagedGroupConfigObservedState struct {
	// Output only. The name of the Instance Template used for the Managed
	//  Instance Group.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedGroupConfig.instance_template_name
	InstanceTemplateName *string `json:"instanceTemplateName,omitempty"`

	// Output only. The name of the Instance Group Manager for this group.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedGroupConfig.instance_group_manager_name
	InstanceGroupManagerName *string `json:"instanceGroupManagerName,omitempty"`

	// Output only. The partial URI to the instance group manager for this group.
	//  E.g. projects/my-project/regions/us-central1/instanceGroupManagers/my-igm.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedGroupConfig.instance_group_manager_uri
	InstanceGroupManagerURI *string `json:"instanceGroupManagerURI,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.WorkflowTemplate
type WorkflowTemplateObservedState struct {
	// Output only. The resource name of the workflow template, as described
	//  in https://cloud.google.com/apis/design/resource_names.
	//
	//  * For `projects.regions.workflowTemplates`, the resource name of the
	//    template has the following format:
	//    `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}`
	//
	//  * For `projects.locations.workflowTemplates`, the resource name of the
	//    template has the following format:
	//    `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}`
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.name
	Name *string `json:"name,omitempty"`

	// Output only. The time template was created.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time template was last updated.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Required. WorkflowTemplate scheduling information.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.placement
	Placement *WorkflowTemplatePlacementObservedState `json:"placement,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.WorkflowTemplatePlacement
type WorkflowTemplatePlacementObservedState struct {
	// A cluster that is managed by the workflow.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplatePlacement.managed_cluster
	ManagedCluster *ManagedClusterObservedState `json:"managedCluster,omitempty"`
}
