// Copyright 2026 Google LLC
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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	containerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DataprocAutoscalingPolicyRef struct {
	External  string `json:"external,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type DataprocMetastoreServiceRef struct {
	External  string `json:"external,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type DataprocGKENodePoolRef struct {
	External  string `json:"external,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type DataprocClusterProjectRef struct {
	External  string `json:"external,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:spec:proto=google.cloud.dataproc.v1.Cluster
type DataprocClusterSpec struct {
	// Immutable. The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.config
	Config *ClusterConfig `json:"config,omitempty"`

	/* Immutable. The location of this resource. */
	Location string `json:"location"`

	/* The project that this resource belongs to. */
	ProjectRef *DataprocClusterProjectRef `json:"projectRef"`

	/* The DataprocCluster name. If not given, the metadata.name will be used. */
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The virtual cluster config is used when creating a Dataproc cluster that does not directly control the underlying compute resources.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.virtual_cluster_config
	VirtualClusterConfig *VirtualClusterConfig `json:"virtualClusterConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ClusterConfig
type ClusterConfig struct {
	/* Immutable. Optional. Autoscaling config for the policy associated with the cluster. Cluster does not autoscale if this field is unset. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.autoscaling_config
	AutoscalingConfig *AutoscalingConfig `json:"autoscalingConfig,omitempty"`

	/* Immutable. Optional. The config for Dataproc metrics. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.dataproc_metric_config
	DataprocMetricConfig *DataprocMetricConfig `json:"dataprocMetricConfig,omitempty"`

	/* Immutable. Optional. Encryption settings for the cluster. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	/* Immutable. Optional. Port/endpoint configuration for this cluster */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.endpoint_config
	EndpointConfig *EndpointConfig `json:"endpointConfig,omitempty"`

	/* Immutable. Optional. The shared Compute Engine config settings for all instances in a cluster. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.gce_cluster_config
	GceClusterConfig *GceClusterConfig `json:"gceClusterConfig,omitempty"`

	/* Immutable. Optional. Commands to execute on each node after config is completed. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.initialization_actions
	InitializationActions []NodeInitializationAction `json:"initializationActions,omitempty"`

	/* Immutable. Optional. Lifecycle setting for the cluster. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.lifecycle_config
	LifecycleConfig *LifecycleConfig `json:"lifecycleConfig,omitempty"`

	/* Immutable. Optional. The Compute Engine config settings for the cluster's master instance. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.master_config
	MasterConfig *InstanceGroupConfig `json:"masterConfig,omitempty"`

	/* Immutable. Optional. Metastore configuration. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.metastore_config
	MetastoreConfig *MetastoreConfig `json:"metastoreConfig,omitempty"`

	/* Immutable. Optional. The Compute Engine config settings for a cluster's secondary worker instances. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.secondary_worker_config
	SecondaryWorkerConfig *InstanceGroupConfig `json:"secondaryWorkerConfig,omitempty"`

	/* Immutable. Optional. Security settings for the cluster. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.security_config
	SecurityConfig *SecurityConfig `json:"securityConfig,omitempty"`

	/* Immutable. Optional. The config settings for cluster software. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.software_config
	SoftwareConfig *SoftwareConfig `json:"softwareConfig,omitempty"`

	/* Immutable. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.config_bucket
	StagingBucketRef *storagev1beta1.StorageBucketRef `json:"stagingBucketRef,omitempty"`

	/* Immutable. Optional. The temp bucket used to store ephemeral cluster and jobs data. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.temp_bucket
	TempBucketRef *storagev1beta1.StorageBucketRef `json:"tempBucketRef,omitempty"`

	/* Immutable. Optional. The Compute Engine config settings for the cluster's worker instances. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.worker_config
	WorkerConfig *InstanceGroupConfig `json:"workerConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.AutoscalingConfig
type AutoscalingConfig struct {
	/* Immutable. */
	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingConfig.policy_uri
	PolicyRef *DataprocAutoscalingPolicyRef `json:"policyRef,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.DataprocMetricConfig
type DataprocMetricConfig struct {
	/* Immutable. Required. Metrics sources to enable. */
	// +kcc:proto:field=google.cloud.dataproc.v1.DataprocMetricConfig.metrics
	Metrics []DataprocMetricConfig_Metric `json:"metrics"`
}

// +kcc:proto=google.cloud.dataproc.v1.DataprocMetricConfig.Metric
type DataprocMetricConfig_Metric struct {
	/* Immutable. Optional. Specify one or more [available OSS metrics] (https://cloud.google.com/dataproc/docs/guides/monitoring#available_oss_metrics) to collect for the metric course (for the `SPARK` metric source, any [Spark metric] (https://spark.apache.org/docs/latest/monitoring.html#metrics) can be specified). Provide metrics in the following format: `METRIC_SOURCE:INSTANCE:GROUP:METRIC` Use camelcase as appropriate. Examples: ``` yarn:ResourceManager:QueueMetrics:AppsCompleted spark:driver:DAGScheduler:job.allJobs sparkHistoryServer:JVM:Memory:NonHeapMemoryUsage.committed hiveserver2:JVM:Memory:NonHeapMemoryUsage.used ``` Notes: * Only the specified overridden metrics will be collected for the metric source. For example, if one or more `spark:executive` metrics are listed as metric overrides, other `SPARK` metrics will not be collected. The collection of the default metrics for other OSS metric sources is unaffected. For example, if both `SPARK` andd `YARN` metric sources are enabled, and overrides are provided for Spark metrics only, all default YARN metrics will be collected. */
	// +kcc:proto:field=google.cloud.dataproc.v1.DataprocMetricConfig.Metric.metric_overrides
	MetricOverrides []string `json:"metricOverrides,omitempty"`

	/* Immutable. Required. Default metrics are collected unless `metricOverrides` are specified for the metric source (see [Available OSS metrics] (https://cloud.google.com/dataproc/docs/guides/monitoring#available_oss_metrics) for more information). Possible values: METRIC_SOURCE_UNSPECIFIED, MONITORING_AGENT_DEFAULTS, HDFS, SPARK, YARN, SPARK_HISTORY_SERVER, HIVESERVER2 */
	// +kcc:proto:field=google.cloud.dataproc.v1.DataprocMetricConfig.Metric.metric_source
	MetricSource string `json:"metricSource"`
}

// +kcc:proto=google.cloud.dataproc.v1.EncryptionConfig
type EncryptionConfig struct {
	/* Immutable. */
	// +kcc:proto:field=google.cloud.dataproc.v1.EncryptionConfig.gce_pd_kms_key_name
	GcePdKmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"gcePdKmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.EndpointConfig
type EndpointConfig struct {
	/* Immutable. Optional. If true, enable http access to specific ports on the cluster from external sources. Defaults to false. */
	// +kcc:proto:field=google.cloud.dataproc.v1.EndpointConfig.enable_http_port_access
	EnableHttpPortAccess *bool `json:"enableHttpPortAccess,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GceClusterConfig
type GceClusterConfig struct {
	/* Immutable. Optional. Confidential Instance Config for clusters using [Confidential VMs](https://cloud.google.com/compute/confidential-vm/docs). */
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.confidential_instance_config
	ConfidentialInstanceConfig *ConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty"`

	/* Immutable. Optional. This setting applies to subnetwork-enabled networks. It is set to `true` by default in clusters created with image versions 2.2.x. When set to `true`: * All cluster VMs have internal IP addresses. * [Google Private Access] (https://cloud.google.com/vpc/docs/private-google-access) must be enabled to access Dataproc and other Google Cloud APIs. * Off-cluster dependencies must be configured to be accessible without external IP addresses. When set to `false`: * Cluster VMs are not restricted to internal IP addresses. * Ephemeral external IP addresses are assigned to each cluster VM. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.internal_ip_only
	InternalIPOnly *bool `json:"internalIPOnly,omitempty"`

	/* Immutable. Optional. The Compute Engine metadata entries to add to all instances (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)). */
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	/* Immutable. Optional. The Compute Engine network to be used for machine communications. Cannot be specified with subnetwork_uri. If neither `network_uri` nor `subnetwork_uri` is specified, the "default" network of the project is used, if it exists. Cannot be a "Custom Subnet Network" (see [Using Subnetworks](https://cloud.google.com/compute/docs/subnetworks) for more information). A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/global/networks/default` * `projects/[project_id]/global/networks/default` * `default` */
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.network_uri
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	/* Immutable. Optional. Node Group Affinity for sole-tenant clusters. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.node_group_affinity
	NodeGroupAffinity *NodeGroupAffinity `json:"nodeGroupAffinity,omitempty"`

	/* Immutable. Optional. The type of IPv6 access for a cluster. Possible values: PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED, INHERIT_FROM_SUBNET, OUTBOUND_ONLY, BIDIRECTIONAL */
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.private_ipv6_google_access
	PrivateIPv6GoogleAccess *string `json:"privateIPv6GoogleAccess,omitempty"`

	/* Immutable. Optional. Reservation Affinity for consuming Zonal reservation. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.reservation_affinity
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`

	/* Immutable. Optional. The [Dataproc service account](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/service-accounts#service_accounts_in_dataproc) (also see [VM Data Plane identity](https://cloud.google.com/dataproc/docs/concepts/iam/dataproc-principals#vm_service_account_data_plane_identity)) used by Dataproc cluster VM instances to access Google Cloud Platform services. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	/* Immutable. Optional. The URIs of service account scopes to be included in Compute Engine instances. The following base set of scopes is always included: * https://www.googleapis.com/auth/cloud.useraccounts.readonly * https://www.googleapis.com/auth/devstorage.read_write * https://www.googleapis.com/auth/logging.write If no scopes are specified, the following defaults are also provided: * https://www.googleapis.com/auth/bigquery * https://www.googleapis.com/auth/bigtable.admin.table * https://www.googleapis.com/auth/bigtable.data * https://www.googleapis.com/auth/devstorage.full_control */
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.service_account_scopes
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty"`

	/* Immutable. Optional. Shielded Instance Config for clusters using [Compute Engine Shielded VMs](https://cloud.google.com/security/shielded-cloud/shielded-vm). */
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.shielded_instance_config
	ShieldedInstanceConfig *ShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	/* Immutable. Optional. The Compute Engine subnetwork to be used for machine communications. Cannot be specified with network_uri. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/[region]/subnetworks/sub0` * `projects/[project_id]/regions/[region]/subnetworks/sub0` * `sub0` */
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.subnetwork_uri
	SubnetworkRef *refsv1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	/* Immutable. The Compute Engine network tags to add to all instances (see [Tagging instances](https://cloud.google.com/vpc/docs/add-remove-network-tags)). */
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.tags
	Tags []string `json:"tags,omitempty"`

	/* Immutable. Optional. The Compute Engine zone where the Dataproc cluster will be located. If omitted, the service will pick a zone in the cluster's Compute Engine region. On a get request, zone will always be present. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]` * `projects/[project_id]/zones/[zone]` * `[zone]` */
	// +kcc:proto:field=google.cloud.dataproc.v1.GceClusterConfig.zone_uri
	Zone *string `json:"zone,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ConfidentialInstanceConfig
type ConfidentialInstanceConfig struct {
	/* Immutable. Optional. Defines whether the instance should have confidential compute enabled. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ConfidentialInstanceConfig.enable_confidential_compute
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.NodeGroupAffinity
type NodeGroupAffinity struct {
	/* Immutable. */
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeGroupAffinity.node_group_uri
	NodeGroupRef *DataprocGKENodePoolRef `json:"nodeGroupRef"`
}

// +kcc:proto=google.cloud.dataproc.v1.ReservationAffinity
type ReservationAffinity struct {
	/* Immutable. Optional. Type of reservation to consume Possible values: TYPE_UNSPECIFIED, NO_RESERVATION, ANY_RESERVATION, SPECIFIC_RESERVATION */
	// +kcc:proto:field=google.cloud.dataproc.v1.ReservationAffinity.consume_reservation_type
	ConsumeReservationType *string `json:"consumeReservationType,omitempty"`

	/* Immutable. Optional. Corresponds to the label key of reservation resource. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ReservationAffinity.key
	Key *string `json:"key,omitempty"`

	/* Immutable. Optional. Corresponds to the label values of reservation resource. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ReservationAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.ShieldedInstanceConfig
type ShieldedInstanceConfig struct {
	/* Immutable. Optional. Defines whether instances have integrity monitoring enabled. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ShieldedInstanceConfig.enable_integrity_monitoring
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`

	/* Immutable. Optional. Defines whether instances have Secure Boot enabled. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ShieldedInstanceConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	/* Immutable. Optional. Defines whether instances have the vTPM enabled. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ShieldedInstanceConfig.enable_vtpm
	EnableVtpm *bool `json:"enableVtpm,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.NodeInitializationAction
type NodeInitializationAction struct {
	/* Immutable. Required. Cloud Storage URI of executable file. */
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeInitializationAction.executable_file
	ExecutableFile string `json:"executableFile"`

	/* Immutable. Optional. Amount of time executable has to complete. Default is 10 minutes (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). Cluster creation fails with an explanatory error message (the name of the executable that caused the error and the exceeded timeout period) if the executable is not completed at end of the timeout period. */
	// +kcc:proto:field=google.cloud.dataproc.v1.NodeInitializationAction.execution_timeout
	ExecutionTimeout *string `json:"executionTimeout,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.LifecycleConfig
type LifecycleConfig struct {
	/* Immutable. Optional. The time when cluster will be auto-deleted (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.cloud.dataproc.v1.LifecycleConfig.auto_delete_time
	AutoDeleteTime *string `json:"autoDeleteTime,omitempty"`

	/* Immutable. Optional. The lifetime duration of cluster. The cluster will be auto-deleted at the end of this period. Minimum value is 10 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	// +kcc:proto:field=google.cloud.dataproc.v1.LifecycleConfig.auto_delete_ttl
	AutoDeleteTtl *string `json:"autoDeleteTtl,omitempty"`

	/* Immutable. Optional. The duration to keep the cluster alive while idling (when no jobs are running). Passing this threshold will cause the cluster to be deleted. Minimum value is 5 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	// +kcc:proto:field=google.cloud.dataproc.v1.LifecycleConfig.idle_delete_ttl
	IdleDeleteTtl *string `json:"idleDeleteTtl,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceGroupConfig
type InstanceGroupConfig struct {
	/* Immutable. Optional. The Compute Engine accelerator configuration for these instances. */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.accelerators
	Accelerators []AcceleratorConfig `json:"accelerators,omitempty"`

	/* Immutable. Optional. Disk option config settings. */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.disk_config
	DiskConfig *DiskConfig `json:"diskConfig,omitempty"`

	/* Immutable. */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.image_uri
	ImageRef *computev1beta1.ComputeImageRef `json:"imageRef,omitempty"`

	/* Immutable. Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]/machineTypes/n1-standard-2` * `projects/[project_id]/zones/[zone]/machineTypes/n1-standard-2` * `n1-standard-2` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, `n1-standard-2`. */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.machine_type_uri
	MachineType *string `json:"machineType,omitempty"`

	/* Immutable. Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -> Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu). */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.min_cpu_platform
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Immutable. Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**. */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.num_instances
	NumInstances *int64 `json:"numInstances,omitempty"`

	/* Immutable. Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is `NON_PREEMPTIBLE`. This default cannot be changed. The default value for secondary instances is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.preemptibility
	Preemptibility *string `json:"preemptibility,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.AcceleratorConfig
type AcceleratorConfig struct {
	/* Immutable. The number of the accelerator cards of this type exposed to this instance. */
	// +kcc:proto:field=google.cloud.dataproc.v1.AcceleratorConfig.accelerator_count
	AcceleratorCount *int64 `json:"acceleratorCount,omitempty"`

	/* Immutable. Full URL, partial URI, or short name of the accelerator type resource to expose to this instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/v1/acceleratorTypes). Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]/acceleratorTypes/nvidia-tesla-t4` * `projects/[project_id]/zones/[zone]/acceleratorTypes/nvidia-tesla-t4` * `nvidia-tesla-t4` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the accelerator type resource, for example, `nvidia-tesla-t4`. */
	// +kcc:proto:field=google.cloud.dataproc.v1.AcceleratorConfig.accelerator_type_uri
	AcceleratorType *string `json:"acceleratorType,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.DiskConfig
type DiskConfig struct {
	/* Immutable. Optional. Size in GB of the boot disk (default is 500GB). */
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_size_gb
	BootDiskSizeGB *int64 `json:"bootDiskSizeGb,omitempty"`

	/* Immutable. Optional. Type of the boot disk (default is "pd-standard"). Valid values: "pd-balanced" (Persistent Disk Balanced Solid State Drive), "pd-ssd" (Persistent Disk Solid State Drive), or "pd-standard" (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types). */
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.boot_disk_type
	BootDiskType *string `json:"bootDiskType,omitempty"`

	/* Immutable. Optional. Interface type of local SSDs (default is "scsi"). Valid values: "scsi" (Small Computer System Interface), "nvme" (Non-Volatile Memory Express). See [local SSD performance](https://cloud.google.com/compute/docs/disks/local-ssd#performance). */
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.local_ssd_interface
	LocalSSDInterface *string `json:"localSsdInterface,omitempty"`

	/* Immutable. Optional. Number of attached SSDs, from 0 to 8 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries. Note: Local SSD options may vary by machine type and number of vCPUs selected. */
	// +kcc:proto:field=google.cloud.dataproc.v1.DiskConfig.num_local_ssds
	NumLocalSSDs *int64 `json:"numLocalSsds,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.MetastoreConfig
type MetastoreConfig struct {
	/* Immutable. */
	// +kcc:proto:field=google.cloud.dataproc.v1.MetastoreConfig.dataproc_metastore_service
	DataprocMetastoreServiceRef *DataprocMetastoreServiceRef `json:"dataprocMetastoreServiceRef"`
}

// +kcc:proto=google.cloud.dataproc.v1.SecurityConfig
type SecurityConfig struct {
	/* Immutable. Optional. Identity related configuration, including service account based secure multi-tenancy user mappings. */
	// +kcc:proto:field=google.cloud.dataproc.v1.SecurityConfig.identity_config
	IdentityConfig *IdentityConfig `json:"identityConfig,omitempty"`

	/* Immutable. Optional. Kerberos related configuration. */
	// +kcc:proto:field=google.cloud.dataproc.v1.SecurityConfig.kerberos_config
	KerberosConfig *KerberosConfig `json:"kerberosConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.IdentityConfig
type IdentityConfig struct {
	/* Immutable. Required. Map of user to service account. */
	// +kcc:proto:field=google.cloud.dataproc.v1.IdentityConfig.user_service_account_mapping
	UserServiceAccountMapping map[string]string `json:"userServiceAccountMapping"`
}

// +kcc:proto=google.cloud.dataproc.v1.KerberosConfig
type KerberosConfig struct {
	/* Immutable. Optional. The admin server (IP or hostname) for the remote trusted realm in a cross realm trust relationship. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.cross_realm_trust_admin_server
	CrossRealmTrustAdminServer *string `json:"crossRealmTrustAdminServer,omitempty"`

	/* Immutable. Optional. The KDC (IP or hostname) for the remote trusted realm in a cross realm trust relationship. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.cross_realm_trust_kdc
	CrossRealmTrustKdc *string `json:"crossRealmTrustKdc,omitempty"`

	/* Immutable. Optional. The remote realm the Dataproc on-cluster KDC will trust, should the user enable cross realm trust. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.cross_realm_trust_realm
	CrossRealmTrustRealm *string `json:"crossRealmTrustRealm,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the shared password between the on-cluster Kerberos realm and the remote trusted realm, in a cross realm trust relationship. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.cross_realm_trust_shared_password_uri
	CrossRealmTrustSharedPassword *string `json:"crossRealmTrustSharedPassword,omitempty"`

	/* Immutable. Optional. Flag to indicate whether to Kerberize the cluster (default: false). Set this field to true to enable Kerberos on a cluster. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.enable_kerberos
	EnableKerberos *bool `json:"enableKerberos,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the master key of the KDC database. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.kdc_db_key_uri
	KdcDbKey *string `json:"kdcDbKey,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided key. For the self-signed certificate, this password is generated by Dataproc. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.key_password_uri
	KeyPassword *string `json:"keyPassword,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided keystore. For the self-signed certificate, this password is generated by Dataproc. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.keystore_password_uri
	KeystorePassword *string `json:"keystorePassword,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of the keystore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.keystore_uri
	Keystore *string `json:"keystore,omitempty"`

	/* Immutable. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.kms_key_uri
	KMSKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	/* Immutable. Optional. The name of the on-cluster Kerberos realm. If not specified, the uppercased domain of hostnames will be the realm. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.realm
	Realm *string `json:"realm,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the root principal password. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.root_principal_password_uri
	RootPrincipalPassword *string `json:"rootPrincipalPassword,omitempty"`

	/* Immutable. Optional. The lifetime of the ticket granting ticket, in hours. If not specified, or user specifies 0, then default value 10 will be used. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.tgt_lifetime_hours
	TgtLifetimeHours *int64 `json:"tgtLifetimeHours,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided truststore. For the self-signed certificate, this password is generated by Dataproc. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.truststore_password_uri
	TruststorePassword *string `json:"truststorePassword,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of the truststore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KerberosConfig.truststore_uri
	Truststore *string `json:"truststore,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SoftwareConfig
type SoftwareConfig struct {
	/* Immutable. Optional. The version of software inside the cluster. It must be one of the supported [Dataproc Versions](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#supported-dataproc-image-versions), such as "1.2" (including a subminor version, such as "1.2.29"), or the ["preview" version](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#other_versions). If unspecified, it defaults to the latest Debian version. */
	// +kcc:proto:field=google.cloud.dataproc.v1.SoftwareConfig.image_version
	ImageVersion *string `json:"imageVersion,omitempty"`

	/* Immutable. Optional. The set of components to activate on the cluster. */
	// +kcc:proto:field=google.cloud.dataproc.v1.SoftwareConfig.optional_components
	OptionalComponents []string `json:"optionalComponents,omitempty"`

	/* Immutable. Optional. The properties to set on daemon config files. Property keys are specified in `prefix:property` format, for example `core:hadoop.tmp.dir`. The following are supported prefixes and their mappings: * capacity-scheduler: `capacity-scheduler.xml` * core: `core-site.xml` * distcp: `distcp-default.xml` * hdfs: `hdfs-site.xml` * hive: `hive-site.xml` * mapred: `mapred-site.xml` * pig: `pig.properties` * spark: `spark-defaults.conf` * yarn: `yarn-site.xml` For more information, see [Cluster properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties). */
	// +kcc:proto:field=google.cloud.dataproc.v1.SoftwareConfig.properties
	Properties map[string]string `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.VirtualClusterConfig
type VirtualClusterConfig struct {
	/* Immutable. Optional. Configuration of auxiliary services used by this cluster. */
	// +kcc:proto:field=google.cloud.dataproc.v1.VirtualClusterConfig.auxiliary_services_config
	AuxiliaryServicesConfig *AuxiliaryServicesConfig `json:"auxiliaryServicesConfig,omitempty"`

	/* Immutable. Required. The configuration for running the Dataproc cluster on Kubernetes. */
	// +kcc:proto:field=google.cloud.dataproc.v1.VirtualClusterConfig.kubernetes_cluster_config
	KubernetesClusterConfig KubernetesClusterConfig `json:"kubernetesClusterConfig"`

	/* Immutable. */
	// +kcc:proto:field=google.cloud.dataproc.v1.VirtualClusterConfig.staging_bucket
	StagingBucketRef *storagev1beta1.StorageBucketRef `json:"stagingBucketRef,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.AuxiliaryServicesConfig
type AuxiliaryServicesConfig struct {
	/* Immutable. Optional. The Hive Metastore configuration for this workload. */
	// +kcc:proto:field=google.cloud.dataproc.v1.AuxiliaryServicesConfig.metastore_config
	MetastoreConfig *MetastoreConfig `json:"metastoreConfig,omitempty"`

	/* Immutable. Optional. The Spark History Server configuration for the workload. */
	// +kcc:proto:field=google.cloud.dataproc.v1.AuxiliaryServicesConfig.spark_history_server_config
	SparkHistoryServerConfig *SparkHistoryServerConfig `json:"sparkHistoryServerConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.SparkHistoryServerConfig
type SparkHistoryServerConfig struct {
	/* Immutable. */
	// +kcc:proto:field=google.cloud.dataproc.v1.SparkHistoryServerConfig.dataproc_cluster
	DataprocClusterRef *DataprocClusterRef `json:"dataprocClusterRef,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.KubernetesClusterConfig
type KubernetesClusterConfig struct {
	/* Immutable. Required. The configuration for running the Dataproc cluster on GKE. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KubernetesClusterConfig.gke_cluster_config
	GKEClusterConfig GkeClusterConfig `json:"gkeClusterConfig"`

	/* Immutable. Optional. A namespace within the Kubernetes cluster to deploy into. If this namespace does not exist, it is created. If it exists, Dataproc verifies that another Dataproc VirtualCluster is not installed into it. If not specified, the name of the Dataproc Cluster is used. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KubernetesClusterConfig.kubernetes_namespace
	KubernetesNamespace *string `json:"kubernetesNamespace,omitempty"`

	/* Immutable. Optional. The software configuration for this Dataproc cluster running on Kubernetes. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KubernetesClusterConfig.kubernetes_software_config
	KubernetesSoftwareConfig *KubernetesSoftwareConfig `json:"kubernetesSoftwareConfig,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeClusterConfig
type GkeClusterConfig struct {
	/* Immutable. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeClusterConfig.gke_cluster_target
	GkeClusterTargetRef *containerv1beta1.ContainerClusterRef `json:"gkeClusterTargetRef,omitempty"`

	/* Immutable. Optional. GKE node pools where workloads will be scheduled. At least one node pool must be assigned the `DEFAULT` [GkeNodePoolTarget.Role][google.cloud.dataproc.v1.GkeNodePoolTarget.Role]. If a `GkeNodePoolTarget` is not specified, Dataproc constructs a `DEFAULT` `GkeNodePoolTarget`. Each role can be given to only one `GkeNodePoolTarget`. All node pools must have the same location settings. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeClusterConfig.node_pool_target
	NodePoolTarget []GkeNodePoolTarget `json:"nodePoolTarget,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeNodePoolTarget
type GkeNodePoolTarget struct {
	/* Immutable. Input only. The configuration for the GKE node pool. If specified, Dataproc attempts to create a node pool with the specified shape. If one with the same name already exists, it is verified against all specified fields. If a field differs, the virtual cluster creation will fail. If omitted, any node pool with the specified name is used. If a node pool with the specified name does not exist, Dataproc create a node pool with default values. This is an input only field. It will not be returned by the API. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolTarget.node_pool_config
	NodePoolConfig *GkeNodePoolConfig `json:"nodePoolConfig,omitempty"`

	/* Immutable. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolTarget.node_pool
	NodePoolRef *DataprocGKENodePoolRef `json:"nodePoolRef"`

	/* Immutable. Required. The roles associated with the GKE node pool. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolTarget.roles
	Roles []string `json:"roles"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeNodePoolConfig
type GkeNodePoolConfig struct {
	/* Immutable. Optional. The autoscaler configuration for this node pool. The autoscaler is enabled only when a valid configuration is present. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.autoscaling
	Autoscaling *GkeNodePoolConfig_GkeNodePoolAutoscalingConfig `json:"autoscaling,omitempty"`

	/* Immutable. Optional. The node pool configuration. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.config
	Config *GkeNodePoolConfig_GkeNodeConfig `json:"config,omitempty"`

	/* Immutable. Optional. The list of Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) where node pool nodes associated with a Dataproc on GKE virtual cluster will be located. **Note:** All node pools associated with a virtual cluster must be located in the same region as the virtual cluster, and they must be located in the same zone within that region. If a location is not specified during node pool creation, Dataproc on GKE will choose the zone. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.locations
	Locations []string `json:"locations,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAutoscalingConfig
type GkeNodePoolConfig_GkeNodePoolAutoscalingConfig struct {
	/* Immutable. The maximum number of nodes in the node pool. Must be >= min_node_count, and must be > 0. **Note:** Quota must be sufficient to scale up the cluster. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAutoscalingConfig.max_node_count
	MaxNodeCount *int64 `json:"maxNodeCount,omitempty"`

	/* Immutable. The minimum number of nodes in the node pool. Must be >= 0 and <= max_node_count. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAutoscalingConfig.min_node_count
	MinNodeCount *int64 `json:"minNodeCount,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig
type GkeNodePoolConfig_GkeNodeConfig struct {
	/* Immutable. Optional. A list of [hardware accelerators](https://cloud.google.com/compute/docs/gpus) to attach to each node. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.accelerators
	Accelerators []GkeNodePoolConfig_GkeNodePoolAcceleratorConfig `json:"accelerators,omitempty"`

	/* Immutable. Optional. The [Customer Managed Encryption Key (CMEK)] (https://cloud.google.com/kubernetes-engine/docs/how-to/using-cmek) used to encrypt the boot disk attached to each node in the node pool. Specify the key using the following format: `projects/KEY_PROJECT_ID/locations/LOCATION/keyRings/RING_NAME/cryptoKeys/KEY_NAME`. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.boot_disk_kms_key
	BootDiskKmsKey *string `json:"bootDiskKmsKey,omitempty"`

	/* Immutable. Optional. Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk. */
	EphemeralStorageConfig *DataprocClusterEphemeralStorageConfigSpec `json:"ephemeralStorageConfig,omitempty"`

	/* Immutable. Optional. The number of local SSD disks to attach to the node, which is limited by the maximum number of disks allowable per zone (see [Adding Local SSDs](https://cloud.google.com/compute/docs/disks/local-ssd)). */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.local_ssd_count
	LocalSsdCount *int64 `json:"localSsdCount,omitempty"`

	/* Immutable. Optional. The name of a Compute Engine [machine type](https://cloud.google.com/compute/docs/machine-types). */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.machine_type
	MachineType *string `json:"machineType,omitempty"`

	/* Immutable. Optional. [Minimum CPU platform](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform) to be used by this instance. The instance may be scheduled on the specified or a newer CPU platform. Specify the friendly names of CPU platforms, such as "Intel Haswell"` or Intel Sandy Bridge". */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.min_cpu_platform
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Immutable. Optional. Whether the nodes are created as legacy [preemptible VM instances] (https://cloud.google.com/compute/docs/instances/preemptible). Also see Spot VMs, preemptible VM instances without a maximum lifetime. Legacy and Spot preemptible nodes cannot be used in a node pool with the `CONTROLLER` [role] (/dataproc/docs/reference/rest/v1/projects.regions.clusters#role) or in the DEFAULT node pool if the CONTROLLER role is not assigned (the DEFAULT node pool will assume the CONTROLLER role). */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.preemptible
	Preemptible *bool `json:"preemptible,omitempty"`

	/* Immutable. Optional. Whether the nodes are created as [Spot VM instances] (https://cloud.google.com/compute/docs/instances/spot). Spot VMs are the latest update to legacy preemptible VMs. Spot VMs do not have a maximum lifetime. Legacy and Spot preemptible nodes cannot be used in a node pool with the `CONTROLLER` [role](/dataproc/docs/reference/rest/v1/projects.regions.clusters#role) or in the DEFAULT node pool if the CONTROLLER role is not assigned (the DEFAULT node pool will assume the CONTROLLER role). */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodeConfig.spot
	Spot *bool `json:"spot,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAcceleratorConfig
type GkeNodePoolConfig_GkeNodePoolAcceleratorConfig struct {
	/* Immutable. The number of accelerator cards exposed to an instance. */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAcceleratorConfig.accelerator_count
	AcceleratorCount *int64 `json:"acceleratorCount,omitempty"`

	/* Immutable. The accelerator type resource namename (see GPUs on Compute Engine). */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAcceleratorConfig.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	/* Immutable. Size of partitions to create on the GPU. Valid values are described in the NVIDIA [mig user guide](https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning). */
	// +kcc:proto:field=google.cloud.dataproc.v1.GkeNodePoolConfig.GkeNodePoolAcceleratorConfig.gpu_partition_size
	GpuPartitionSize *string `json:"gpuPartitionSize,omitempty"`
}

type DataprocClusterEphemeralStorageConfigSpec struct {
	/* Immutable. Number of local SSDs to use to back ephemeral storage. Uses NVMe interfaces. Each local SSD is 375 GB in size. If zero, it means to disable using local SSDs as ephemeral storage. */
	LocalSsdCount *int64 `json:"localSsdCount,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.KubernetesSoftwareConfig
type KubernetesSoftwareConfig struct {
	/* Immutable. The components that should be installed in this Dataproc cluster. The key must be a string from the KubernetesComponent enumeration. The value is the version of the software to be installed. At least one entry must be specified. */
	// +kcc:proto:field=google.cloud.dataproc.v1.KubernetesSoftwareConfig.component_version
	ComponentVersion map[string]string `json:"componentVersion,omitempty"`

	/* Immutable. The properties to set on daemon config files. Property keys are specified in `prefix:property` format, for example `spark:spark.kubernetes.container.image`. The following are supported prefixes and their mappings: * spark: `spark-defaults.conf` For more information, see [Cluster properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties). */
	// +kcc:proto:field=google.cloud.dataproc.v1.KubernetesSoftwareConfig.properties
	Properties map[string]string `json:"properties,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataproc.v1.Cluster
type DataprocClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// Output only. A cluster UUID (Unique Universal Identifier). Dataproc generates this value when it creates the cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.cluster_uuid
	ClusterUuid *string `json:"clusterUuid,omitempty"`

	// Output only. Cluster status.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.status
	Status *ClusterStatus `json:"status,omitempty"`

	// Output only. The previous cluster status.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.status_history
	StatusHistory []ClusterStatus `json:"statusHistory,omitempty"`

	// Output only. The cluster config for a cluster of Compute Engine Instances.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.config
	Config *ClusterConfigObservedState `json:"config,omitempty"`

	// Output only. Contains cluster daemon metrics such as HDFS and YARN stats. **Beta Feature**: This report is available for testing purposes only. It may be changed before final release.
	// +kcc:proto:field=google.cloud.dataproc.v1.Cluster.metrics
	Metrics *ClusterMetrics `json:"metrics,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataproc.v1.ClusterStatus
type ClusterStatus struct {
	/* Optional. Output only. Details of cluster's state. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterStatus.detail
	Detail *string `json:"detail,omitempty"`

	/* Output only. The cluster's state. Possible values: UNKNOWN, CREATING, RUNNING, ERROR, DELETING, UPDATING, STOPPING, STOPPED, STARTING */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterStatus.state
	State *string `json:"state,omitempty"`

	/* Output only. Time when this state was entered (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterStatus.state_start_time
	StateStartTime *string `json:"stateStartTime,omitempty"`

	/* Output only. Additional state information that includes status reported by the agent. Possible values: UNSPECIFIED, UNHEALTHY, STALE_STATUS */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterStatus.substate
	Substate *string `json:"substate,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataproc.v1.ClusterConfig
type ClusterConfigObservedState struct {
	/* Output only. Port/endpoint configuration for this cluster */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.endpoint_config
	EndpointConfig *EndpointConfigObservedState `json:"endpointConfig,omitempty"`

	/* Output only. Lifecycle setting for the cluster. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.lifecycle_config
	LifecycleConfig *LifecycleConfigObservedState `json:"lifecycleConfig,omitempty"`

	/* Output only. The Compute Engine config settings for the cluster's master instance. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.master_config
	MasterConfig *InstanceGroupConfigObservedState `json:"masterConfig,omitempty"`

	/* Output only. The Compute Engine config settings for a cluster's secondary worker instances. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.secondary_worker_config
	SecondaryWorkerConfig *InstanceGroupConfigObservedState `json:"secondaryWorkerConfig,omitempty"`

	/* Output only. The Compute Engine config settings for the cluster's worker instances. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterConfig.worker_config
	WorkerConfig *InstanceGroupConfigObservedState `json:"workerConfig,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataproc.v1.EndpointConfig
type EndpointConfigObservedState struct {
	/* Output only. The map of port descriptions to URLs. Will only be populated if enable_http_port_access is true. */
	// +kcc:proto:field=google.cloud.dataproc.v1.EndpointConfig.http_ports
	HttpPorts map[string]string `json:"httpPorts,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataproc.v1.LifecycleConfig
type LifecycleConfigObservedState struct {
	/* Output only. The time when cluster became idle (most recent job finished) and became eligible for deletion due to idleness (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	// +kubebuilder:validation:Format=date-time
	// +kcc:proto:field=google.cloud.dataproc.v1.LifecycleConfig.idle_start_time
	IdleStartTime *string `json:"idleStartTime,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataproc.v1.InstanceGroupConfig
type InstanceGroupConfigObservedState struct {
	/* Output only. The list of instance names. Dataproc derives the names from `cluster_name`, `num_instances`, and the instance group. */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_names
	InstanceNames []string `json:"instanceNames,omitempty"`

	/* Output only. List of references to Compute Engine instances. */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.instance_references
	InstanceReferences []InstanceReference `json:"instanceReferences,omitempty"`

	/* Output only. Specifies that this instance group contains preemptible instances. */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.is_preemptible
	IsPreemptible *bool `json:"isPreemptible,omitempty"`

	/* Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups. */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupConfig.managed_group_config
	ManagedGroupConfig *ManagedGroupConfig `json:"managedGroupConfig,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataproc.v1.InstanceReference
type InstanceReference struct {
	/* The unique identifier of the Compute Engine instance. */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.instance_id
	InstanceId *string `json:"instanceId,omitempty"`

	/* The user-friendly name of the Compute Engine instance. */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.instance_name
	InstanceName *string `json:"instanceName,omitempty"`

	/* The public ECIES key used for sharing data with this instance. */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.public_ecies_key
	PublicEciesKey *string `json:"publicEciesKey,omitempty"`

	/* The public RSA key used for sharing data with this instance. */
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceReference.public_key
	PublicKey *string `json:"publicKey,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataproc.v1.ManagedGroupConfig
type ManagedGroupConfig struct {
	/* Output only. The name of the Instance Group Manager for this group. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedGroupConfig.instance_group_manager_name
	InstanceGroupManagerName *string `json:"instanceGroupManagerName,omitempty"`

	/* Output only. The name of the Instance Template used for the Managed Instance Group. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedGroupConfig.instance_template_name
	InstanceTemplateName *string `json:"instanceTemplateName,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataproc.v1.ClusterMetrics
type ClusterMetrics struct {
	/* The HDFS metrics. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterMetrics.hdfs_metrics
	HdfsMetrics map[string]string `json:"hdfsMetrics,omitempty"`

	/* The YARN metrics. */
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterMetrics.yarn_metrics
	YarnMetrics map[string]string `json:"yarnMetrics,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataproccluster;gcpdataprocclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataprocCluster is the Schema for the DataprocCluster API
// +k8s:openapi-gen=true
type DataprocCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataprocClusterSpec   `json:"spec"`
	Status DataprocClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataprocClusterList contains a list of DataprocCluster
type DataprocClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataprocCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataprocCluster{}, &DataprocClusterList{})
}
