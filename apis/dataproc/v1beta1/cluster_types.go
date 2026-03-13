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

type DataprocClusterSpec struct {
	// Immutable. The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated.
	Config *DataprocClusterConfigSpec `json:"config,omitempty"`

	/* Immutable. The location of this resource. */
	Location string `json:"location"`

	/* The project that this resource belongs to. */
	ProjectRef *DataprocClusterProjectRef `json:"projectRef"`

	/* The DataprocCluster name. If not given, the metadata.name will be used. */
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The virtual cluster config is used when creating a Dataproc cluster that does not directly control the underlying compute resources.
	VirtualClusterConfig *DataprocClusterVirtualClusterConfigSpec `json:"virtualClusterConfig,omitempty"`
}

type DataprocClusterConfigSpec struct {
	/* Immutable. Optional. Autoscaling config for the policy associated with the cluster. Cluster does not autoscale if this field is unset. */
	AutoscalingConfig *DataprocClusterAutoscalingConfigSpec `json:"autoscalingConfig,omitempty"`

	/* Immutable. Optional. The config for Dataproc metrics. */
	DataprocMetricConfig *DataprocClusterDataprocMetricConfigSpec `json:"dataprocMetricConfig,omitempty"`

	/* Immutable. Optional. Encryption settings for the cluster. */
	EncryptionConfig *DataprocClusterEncryptionConfigSpec `json:"encryptionConfig,omitempty"`

	/* Immutable. Optional. Port/endpoint configuration for this cluster */
	EndpointConfig *DataprocClusterEndpointConfigSpec `json:"endpointConfig,omitempty"`

	/* Immutable. Optional. The shared Compute Engine config settings for all instances in a cluster. */
	GceClusterConfig *DataprocClusterGceClusterConfigSpec `json:"gceClusterConfig,omitempty"`

	/* Immutable. Optional. Commands to execute on each node after config is completed. */
	InitializationActions []DataprocClusterInitializationActionsSpec `json:"initializationActions,omitempty"`

	/* Immutable. Optional. Lifecycle setting for the cluster. */
	LifecycleConfig *DataprocClusterLifecycleConfigSpec `json:"lifecycleConfig,omitempty"`

	/* Immutable. Optional. The Compute Engine config settings for the cluster's master instance. */
	MasterConfig *DataprocClusterInstanceGroupConfigSpec `json:"masterConfig,omitempty"`

	/* Immutable. Optional. Metastore configuration. */
	MetastoreConfig *DataprocClusterMetastoreConfigSpec `json:"metastoreConfig,omitempty"`

	/* Immutable. Optional. The Compute Engine config settings for a cluster's secondary worker instances. */
	SecondaryWorkerConfig *DataprocClusterInstanceGroupConfigSpec `json:"secondaryWorkerConfig,omitempty"`

	/* Immutable. Optional. Security settings for the cluster. */
	SecurityConfig *DataprocClusterSecurityConfigSpec `json:"securityConfig,omitempty"`

	/* Immutable. Optional. The config settings for cluster software. */
	SoftwareConfig *DataprocClusterSoftwareConfigSpec `json:"softwareConfig,omitempty"`

	/* Immutable. */
	StagingBucketRef *storagev1beta1.StorageBucketRef `json:"stagingBucketRef,omitempty"`

	/* Immutable. Optional. The temp bucket used to store ephemeral cluster and jobs data. */
	TempBucketRef *storagev1beta1.StorageBucketRef `json:"tempBucketRef,omitempty"`

	/* Immutable. Optional. The Compute Engine config settings for the cluster's worker instances. */
	WorkerConfig *DataprocClusterInstanceGroupConfigSpec `json:"workerConfig,omitempty"`
}

type DataprocClusterAutoscalingConfigSpec struct {
	/* Immutable. */
	PolicyRef *DataprocAutoscalingPolicyRef `json:"policyRef,omitempty"`
}

type DataprocClusterDataprocMetricConfigSpec struct {
	/* Immutable. Required. Metrics sources to enable. */
	Metrics []DataprocClusterMetricsSpec `json:"metrics"`
}

type DataprocClusterMetricsSpec struct {
	/* Immutable. Optional. Specify one or more [available OSS metrics] (https://cloud.google.com/dataproc/docs/guides/monitoring#available_oss_metrics) to collect for the metric course (for the `SPARK` metric source, any [Spark metric] (https://spark.apache.org/docs/latest/monitoring.html#metrics) can be specified). Provide metrics in the following format: `METRIC_SOURCE:INSTANCE:GROUP:METRIC` Use camelcase as appropriate. Examples: ``` yarn:ResourceManager:QueueMetrics:AppsCompleted spark:driver:DAGScheduler:job.allJobs sparkHistoryServer:JVM:Memory:NonHeapMemoryUsage.committed hiveserver2:JVM:Memory:NonHeapMemoryUsage.used ``` Notes: * Only the specified overridden metrics will be collected for the metric source. For example, if one or more `spark:executive` metrics are listed as metric overrides, other `SPARK` metrics will not be collected. The collection of the default metrics for other OSS metric sources is unaffected. For example, if both `SPARK` andd `YARN` metric sources are enabled, and overrides are provided for Spark metrics only, all default YARN metrics will be collected. */
	MetricOverrides []string `json:"metricOverrides,omitempty"`

	/* Immutable. Required. Default metrics are collected unless `metricOverrides` are specified for the metric source (see [Available OSS metrics] (https://cloud.google.com/dataproc/docs/guides/monitoring#available_oss_metrics) for more information). Possible values: METRIC_SOURCE_UNSPECIFIED, MONITORING_AGENT_DEFAULTS, HDFS, SPARK, YARN, SPARK_HISTORY_SERVER, HIVESERVER2 */
	MetricSource string `json:"metricSource"`
}

type DataprocClusterEncryptionConfigSpec struct {
	/* Immutable. */
	GcePdKmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"gcePdKmsKeyRef,omitempty"`
}

type DataprocClusterEndpointConfigSpec struct {
	/* Immutable. Optional. If true, enable http access to specific ports on the cluster from external sources. Defaults to false. */
	EnableHttpPortAccess *bool `json:"enableHttpPortAccess,omitempty"`
}

type DataprocClusterGceClusterConfigSpec struct {
	/* Immutable. Optional. Confidential Instance Config for clusters using [Confidential VMs](https://cloud.google.com/compute/confidential-vm/docs). */
	ConfidentialInstanceConfig *DataprocClusterConfidentialInstanceConfigSpec `json:"confidentialInstanceConfig,omitempty"`

	/* Immutable. Optional. This setting applies to subnetwork-enabled networks. It is set to `true` by default in clusters created with image versions 2.2.x. When set to `true`: * All cluster VMs have internal IP addresses. * [Google Private Access] (https://cloud.google.com/vpc/docs/private-google-access) must be enabled to access Dataproc and other Google Cloud APIs. * Off-cluster dependencies must be configured to be accessible without external IP addresses. When set to `false`: * Cluster VMs are not restricted to internal IP addresses. * Ephemeral external IP addresses are assigned to each cluster VM. */
	InternalIPOnly *bool `json:"internalIPOnly,omitempty"`

	/* Immutable. Optional. The Compute Engine metadata entries to add to all instances (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)). */
	Metadata map[string]string `json:"metadata,omitempty"`

	/* Immutable. Optional. The Compute Engine network to be used for machine communications. Cannot be specified with subnetwork_uri. If neither `network_uri` nor `subnetwork_uri` is specified, the "default" network of the project is used, if it exists. Cannot be a "Custom Subnet Network" (see [Using Subnetworks](https://cloud.google.com/compute/docs/subnetworks) for more information). A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/global/networks/default` * `projects/[project_id]/global/networks/default` * `default` */
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	/* Immutable. Optional. Node Group Affinity for sole-tenant clusters. */
	NodeGroupAffinity *DataprocClusterNodeGroupAffinitySpec `json:"nodeGroupAffinity,omitempty"`

	/* Immutable. Optional. The type of IPv6 access for a cluster. Possible values: PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED, INHERIT_FROM_SUBNET, OUTBOUND_ONLY, BIDIRECTIONAL */
	PrivateIPv6GoogleAccess *string `json:"privateIPv6GoogleAccess,omitempty"`

	/* Immutable. Optional. Reservation Affinity for consuming Zonal reservation. */
	ReservationAffinity *DataprocClusterReservationAffinitySpec `json:"reservationAffinity,omitempty"`

	/* Immutable. Optional. The [Dataproc service account](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/service-accounts#service_accounts_in_dataproc) (also see [VM Data Plane identity](https://cloud.google.com/dataproc/docs/concepts/iam/dataproc-principals#vm_service_account_data_plane_identity)) used by Dataproc cluster VM instances to access Google Cloud Platform services. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used. */
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	/* Immutable. Optional. The URIs of service account scopes to be included in Compute Engine instances. The following base set of scopes is always included: * https://www.googleapis.com/auth/cloud.useraccounts.readonly * https://www.googleapis.com/auth/devstorage.read_write * https://www.googleapis.com/auth/logging.write If no scopes are specified, the following defaults are also provided: * https://www.googleapis.com/auth/bigquery * https://www.googleapis.com/auth/bigtable.admin.table * https://www.googleapis.com/auth/bigtable.data * https://www.googleapis.com/auth/devstorage.full_control */
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty"`

	/* Immutable. Optional. Shielded Instance Config for clusters using [Compute Engine Shielded VMs](https://cloud.google.com/security/shielded-cloud/shielded-vm). */
	ShieldedInstanceConfig *DataprocClusterShieldedInstanceConfigSpec `json:"shieldedInstanceConfig,omitempty"`

	/* Immutable. Optional. The Compute Engine subnetwork to be used for machine communications. Cannot be specified with network_uri. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/[region]/subnetworks/sub0` * `projects/[project_id]/regions/[region]/subnetworks/sub0` * `sub0` */
	SubnetworkRef *refsv1beta1.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	/* Immutable. The Compute Engine network tags to add to all instances (see [Tagging instances](https://cloud.google.com/vpc/docs/add-remove-network-tags)). */
	Tags []string `json:"tags,omitempty"`

	/* Immutable. Optional. The Compute Engine zone where the Dataproc cluster will be located. If omitted, the service will pick a zone in the cluster's Compute Engine region. On a get request, zone will always be present. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]` * `projects/[project_id]/zones/[zone]` * `[zone]` */
	Zone *string `json:"zone,omitempty"`
}

type DataprocClusterConfidentialInstanceConfigSpec struct {
	/* Immutable. Optional. Defines whether the instance should have confidential compute enabled. */
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute,omitempty"`
}

type DataprocClusterNodeGroupAffinitySpec struct {
	/* Immutable. */
	NodeGroupRef *DataprocGKENodePoolRef `json:"nodeGroupRef"`
}

type DataprocClusterReservationAffinitySpec struct {
	/* Immutable. Optional. Type of reservation to consume Possible values: TYPE_UNSPECIFIED, NO_RESERVATION, ANY_RESERVATION, SPECIFIC_RESERVATION */
	ConsumeReservationType *string `json:"consumeReservationType,omitempty"`

	/* Immutable. Optional. Corresponds to the label key of reservation resource. */
	Key *string `json:"key,omitempty"`

	/* Immutable. Optional. Corresponds to the label values of reservation resource. */
	Values []string `json:"values,omitempty"`
}

type DataprocClusterShieldedInstanceConfigSpec struct {
	/* Immutable. Optional. Defines whether instances have integrity monitoring enabled. */
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`

	/* Immutable. Optional. Defines whether instances have Secure Boot enabled. */
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	/* Immutable. Optional. Defines whether instances have the vTPM enabled. */
	EnableVtpm *bool `json:"enableVtpm,omitempty"`
}

type DataprocClusterInitializationActionsSpec struct {
	/* Immutable. Required. Cloud Storage URI of executable file. */
	ExecutableFile string `json:"executableFile"`

	/* Immutable. Optional. Amount of time executable has to complete. Default is 10 minutes (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). Cluster creation fails with an explanatory error message (the name of the executable that caused the error and the exceeded timeout period) if the executable is not completed at end of the timeout period. */
	ExecutionTimeout *string `json:"executionTimeout,omitempty"`
}

type DataprocClusterLifecycleConfigSpec struct {
	/* Immutable. Optional. The time when cluster will be auto-deleted (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	// +kubebuilder:validation:Format=date-time
	AutoDeleteTime *string `json:"autoDeleteTime,omitempty"`

	/* Immutable. Optional. The lifetime duration of cluster. The cluster will be auto-deleted at the end of this period. Minimum value is 10 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	AutoDeleteTtl *string `json:"autoDeleteTtl,omitempty"`

	/* Immutable. Optional. The duration to keep the cluster alive while idling (when no jobs are running). Passing this threshold will cause the cluster to be deleted. Minimum value is 5 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	IdleDeleteTtl *string `json:"idleDeleteTtl,omitempty"`
}

type DataprocClusterInstanceGroupConfigSpec struct {
	/* Immutable. Optional. The Compute Engine accelerator configuration for these instances. */
	Accelerators []DataprocClusterAcceleratorsSpec `json:"accelerators,omitempty"`

	/* Immutable. Optional. Disk option config settings. */
	DiskConfig *DataprocClusterDiskConfigSpec `json:"diskConfig,omitempty"`

	/* Immutable. */
	ImageRef *computev1beta1.ComputeImageRef `json:"imageRef,omitempty"`

	/* Immutable. Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]/machineTypes/n1-standard-2` * `projects/[project_id]/zones/[zone]/machineTypes/n1-standard-2` * `n1-standard-2` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, `n1-standard-2`. */
	MachineType *string `json:"machineType,omitempty"`

	/* Immutable. Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -> Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu). */
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Immutable. Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**. */
	NumInstances *int64 `json:"numInstances,omitempty"`

	/* Immutable. Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is `NON_PREEMPTIBLE`. This default cannot be changed. The default value for secondary instances is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE */
	Preemptibility *string `json:"preemptibility,omitempty"`
}

type DataprocClusterAcceleratorsSpec struct {
	/* Immutable. The number of the accelerator cards of this type exposed to this instance. */
	AcceleratorCount *int64 `json:"acceleratorCount,omitempty"`

	/* Immutable. Full URL, partial URI, or short name of the accelerator type resource to expose to this instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/v1/acceleratorTypes). Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]/acceleratorTypes/nvidia-tesla-t4` * `projects/[project_id]/zones/[zone]/acceleratorTypes/nvidia-tesla-t4` * `nvidia-tesla-t4` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the accelerator type resource, for example, `nvidia-tesla-t4`. */
	AcceleratorType *string `json:"acceleratorType,omitempty"`
}

type DataprocClusterDiskConfigSpec struct {
	/* Immutable. Optional. Size in GB of the boot disk (default is 500GB). */
	BootDiskSizeGb *int64 `json:"bootDiskSizeGb,omitempty"`

	/* Immutable. Optional. Type of the boot disk (default is "pd-standard"). Valid values: "pd-balanced" (Persistent Disk Balanced Solid State Drive), "pd-ssd" (Persistent Disk Solid State Drive), or "pd-standard" (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types). */
	BootDiskType *string `json:"bootDiskType,omitempty"`

	/* Immutable. Optional. Interface type of local SSDs (default is "scsi"). Valid values: "scsi" (Small Computer System Interface), "nvme" (Non-Volatile Memory Express). See [local SSD performance](https://cloud.google.com/compute/docs/disks/local-ssd#performance). */
	LocalSsdInterface *string `json:"localSsdInterface,omitempty"`

	/* Immutable. Optional. Number of attached SSDs, from 0 to 8 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries. Note: Local SSD options may vary by machine type and number of vCPUs selected. */
	NumLocalSsds *int64 `json:"numLocalSsds,omitempty"`
}

type DataprocClusterMetastoreConfigSpec struct {
	/* Immutable. */
	DataprocMetastoreServiceRef *DataprocMetastoreServiceRef `json:"dataprocMetastoreServiceRef"`
}

type DataprocClusterSecurityConfigSpec struct {
	/* Immutable. Optional. Identity related configuration, including service account based secure multi-tenancy user mappings. */
	IdentityConfig *DataprocClusterIdentityConfigSpec `json:"identityConfig,omitempty"`

	/* Immutable. Optional. Kerberos related configuration. */
	KerberosConfig *DataprocClusterKerberosConfigSpec `json:"kerberosConfig,omitempty"`
}

type DataprocClusterIdentityConfigSpec struct {
	/* Immutable. Required. Map of user to service account. */
	UserServiceAccountMapping map[string]string `json:"userServiceAccountMapping"`
}

type DataprocClusterKerberosConfigSpec struct {
	/* Immutable. Optional. The admin server (IP or hostname) for the remote trusted realm in a cross realm trust relationship. */
	CrossRealmTrustAdminServer *string `json:"crossRealmTrustAdminServer,omitempty"`

	/* Immutable. Optional. The KDC (IP or hostname) for the remote trusted realm in a cross realm trust relationship. */
	CrossRealmTrustKdc *string `json:"crossRealmTrustKdc,omitempty"`

	/* Immutable. Optional. The remote realm the Dataproc on-cluster KDC will trust, should the user enable cross realm trust. */
	CrossRealmTrustRealm *string `json:"crossRealmTrustRealm,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the shared password between the on-cluster Kerberos realm and the remote trusted realm, in a cross realm trust relationship. */
	CrossRealmTrustSharedPassword *string `json:"crossRealmTrustSharedPassword,omitempty"`

	/* Immutable. Optional. Flag to indicate whether to Kerberize the cluster (default: false). Set this field to true to enable Kerberos on a cluster. */
	EnableKerberos *bool `json:"enableKerberos,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the master key of the KDC database. */
	KdcDbKey *string `json:"kdcDbKey,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided key. For the self-signed certificate, this password is generated by Dataproc. */
	KeyPassword *string `json:"keyPassword,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided keystore. For the self-signed certificate, this password is generated by Dataproc. */
	KeystorePassword *string `json:"keystorePassword,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of the keystore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate. */
	Keystore *string `json:"keystore,omitempty"`

	/* Immutable. */
	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	/* Immutable. Optional. The name of the on-cluster Kerberos realm. If not specified, the uppercased domain of hostnames will be the realm. */
	Realm *string `json:"realm,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the root principal password. */
	RootPrincipalPassword *string `json:"rootPrincipalPassword,omitempty"`

	/* Immutable. Optional. The lifetime of the ticket granting ticket, in hours. If not specified, or user specifies 0, then default value 10 will be used. */
	TgtLifetimeHours *int64 `json:"tgtLifetimeHours,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided truststore. For the self-signed certificate, this password is generated by Dataproc. */
	TruststorePassword *string `json:"truststorePassword,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of the truststore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate. */
	Truststore *string `json:"truststore,omitempty"`
}

type DataprocClusterSoftwareConfigSpec struct {
	/* Immutable. Optional. The version of software inside the cluster. It must be one of the supported [Dataproc Versions](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#supported-dataproc-image-versions), such as "1.2" (including a subminor version, such as "1.2.29"), or the ["preview" version](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#other_versions). If unspecified, it defaults to the latest Debian version. */
	ImageVersion *string `json:"imageVersion,omitempty"`

	/* Immutable. Optional. The set of components to activate on the cluster. */
	OptionalComponents []string `json:"optionalComponents,omitempty"`

	/* Immutable. Optional. The properties to set on daemon config files. Property keys are specified in `prefix:property` format, for example `core:hadoop.tmp.dir`. The following are supported prefixes and their mappings: * capacity-scheduler: `capacity-scheduler.xml` * core: `core-site.xml` * distcp: `distcp-default.xml` * hdfs: `hdfs-site.xml` * hive: `hive-site.xml` * mapred: `mapred-site.xml` * pig: `pig.properties` * spark: `spark-defaults.conf` * yarn: `yarn-site.xml` For more information, see [Cluster properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties). */
	Properties map[string]string `json:"properties,omitempty"`
}

type DataprocClusterVirtualClusterConfigSpec struct {
	/* Immutable. Optional. Configuration of auxiliary services used by this cluster. */
	AuxiliaryServicesConfig *DataprocClusterAuxiliaryServicesConfigSpec `json:"auxiliaryServicesConfig,omitempty"`

	/* Immutable. Required. The configuration for running the Dataproc cluster on Kubernetes. */
	KubernetesClusterConfig DataprocClusterKubernetesClusterConfigSpec `json:"kubernetesClusterConfig"`

	/* Immutable. */
	StagingBucketRef *storagev1beta1.StorageBucketRef `json:"stagingBucketRef,omitempty"`
}

type DataprocClusterAuxiliaryServicesConfigSpec struct {
	/* Immutable. Optional. The Hive Metastore configuration for this workload. */
	MetastoreConfig *DataprocClusterMetastoreConfigSpec `json:"metastoreConfig,omitempty"`

	/* Immutable. Optional. The Spark History Server configuration for the workload. */
	SparkHistoryServerConfig *DataprocClusterSparkHistoryServerConfigSpec `json:"sparkHistoryServerConfig,omitempty"`
}

type DataprocClusterSparkHistoryServerConfigSpec struct {
	/* Immutable. */
	DataprocClusterRef *DataprocClusterRef `json:"dataprocClusterRef,omitempty"`
}

type DataprocClusterKubernetesClusterConfigSpec struct {
	/* Immutable. Required. The configuration for running the Dataproc cluster on GKE. */
	GkeClusterConfig DataprocClusterGkeClusterConfigSpec `json:"gkeClusterConfig"`

	/* Immutable. Optional. A namespace within the Kubernetes cluster to deploy into. If this namespace does not exist, it is created. If it exists, Dataproc verifies that another Dataproc VirtualCluster is not installed into it. If not specified, the name of the Dataproc Cluster is used. */
	KubernetesNamespace *string `json:"kubernetesNamespace,omitempty"`

	/* Immutable. Optional. The software configuration for this Dataproc cluster running on Kubernetes. */
	KubernetesSoftwareConfig *DataprocClusterKubernetesSoftwareConfigSpec `json:"kubernetesSoftwareConfig,omitempty"`
}

type DataprocClusterGkeClusterConfigSpec struct {
	/* Immutable. */
	GkeClusterTargetRef *containerv1beta1.ContainerClusterRef `json:"gkeClusterTargetRef,omitempty"`

	/* Immutable. Optional. GKE node pools where workloads will be scheduled. At least one node pool must be assigned the `DEFAULT` [GkeNodePoolTarget.Role][google.cloud.dataproc.v1.GkeNodePoolTarget.Role]. If a `GkeNodePoolTarget` is not specified, Dataproc constructs a `DEFAULT` `GkeNodePoolTarget`. Each role can be given to only one `GkeNodePoolTarget`. All node pools must have the same location settings. */
	NodePoolTarget []DataprocClusterNodePoolTargetSpec `json:"nodePoolTarget,omitempty"`
}

type DataprocClusterNodePoolTargetSpec struct {
	/* Immutable. Input only. The configuration for the GKE node pool. If specified, Dataproc attempts to create a node pool with the specified shape. If one with the same name already exists, it is verified against all specified fields. If a field differs, the virtual cluster creation will fail. If omitted, any node pool with the specified name is used. If a node pool with the specified name does not exist, Dataproc create a node pool with default values. This is an input only field. It will not be returned by the API. */
	NodePoolConfig *DataprocClusterNodePoolConfigSpec `json:"nodePoolConfig,omitempty"`

	/* Immutable. */
	NodePoolRef *DataprocGKENodePoolRef `json:"nodePoolRef"`

	/* Immutable. Required. The roles associated with the GKE node pool. */
	Roles []string `json:"roles"`
}

type DataprocClusterNodePoolConfigSpec struct {
	/* Immutable. Optional. The autoscaler configuration for this node pool. The autoscaler is enabled only when a valid configuration is present. */
	Autoscaling *DataprocClusterAutoscalingSpec `json:"autoscaling,omitempty"`

	/* Immutable. Optional. The node pool configuration. */
	Config *DataprocClusterConfig_GKENodeConfigSpec `json:"config,omitempty"`

	/* Immutable. Optional. The list of Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) where node pool nodes associated with a Dataproc on GKE virtual cluster will be located. **Note:** All node pools associated with a virtual cluster must be located in the same region as the virtual cluster, and they must be located in the same zone within that region. If a location is not specified during node pool creation, Dataproc on GKE will choose the zone. */
	Locations []string `json:"locations,omitempty"`
}

type DataprocClusterAutoscalingSpec struct {
	/* Immutable. The maximum number of nodes in the node pool. Must be >= min_node_count, and must be > 0. **Note:** Quota must be sufficient to scale up the cluster. */
	MaxNodeCount *int64 `json:"maxNodeCount,omitempty"`

	/* Immutable. The minimum number of nodes in the node pool. Must be >= 0 and <= max_node_count. */
	MinNodeCount *int64 `json:"minNodeCount,omitempty"`
}

type DataprocClusterConfig_GKENodeConfigSpec struct {
	/* Immutable. Optional. A list of [hardware accelerators](https://cloud.google.com/compute/docs/gpus) to attach to each node. */
	Accelerators []DataprocClusterAccelerators_GKENodePoolAcceleratorConfigSpec `json:"accelerators,omitempty"`

	/* Immutable. Optional. The [Customer Managed Encryption Key (CMEK)] (https://cloud.google.com/kubernetes-engine/docs/how-to/using-cmek) used to encrypt the boot disk attached to each node in the node pool. Specify the key using the following format: `projects/KEY_PROJECT_ID/locations/LOCATION/keyRings/RING_NAME/cryptoKeys/KEY_NAME`. */
	BootDiskKmsKey *string `json:"bootDiskKmsKey,omitempty"`

	/* Immutable. Optional. Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk. */
	EphemeralStorageConfig *DataprocClusterEphemeralStorageConfigSpec `json:"ephemeralStorageConfig,omitempty"`

	/* Immutable. Optional. The number of local SSD disks to attach to the node, which is limited by the maximum number of disks allowable per zone (see [Adding Local SSDs](https://cloud.google.com/compute/docs/disks/local-ssd)). */
	LocalSsdCount *int64 `json:"localSsdCount,omitempty"`

	/* Immutable. Optional. The name of a Compute Engine [machine type](https://cloud.google.com/compute/docs/machine-types). */
	MachineType *string `json:"machineType,omitempty"`

	/* Immutable. Optional. [Minimum CPU platform](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform) to be used by this instance. The instance may be scheduled on the specified or a newer CPU platform. Specify the friendly names of CPU platforms, such as "Intel Haswell"` or Intel Sandy Bridge". */
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Immutable. Optional. Whether the nodes are created as legacy [preemptible VM instances] (https://cloud.google.com/compute/docs/instances/preemptible). Also see Spot VMs, preemptible VM instances without a maximum lifetime. Legacy and Spot preemptible nodes cannot be used in a node pool with the `CONTROLLER` [role] (/dataproc/docs/reference/rest/v1/projects.regions.clusters#role) or in the DEFAULT node pool if the CONTROLLER role is not assigned (the DEFAULT node pool will assume the CONTROLLER role). */
	Preemptible *bool `json:"preemptible,omitempty"`

	/* Immutable. Optional. Whether the nodes are created as [Spot VM instances] (https://cloud.google.com/compute/docs/instances/spot). Spot VMs are the latest update to legacy preemptible VMs. Spot VMs do not have a maximum lifetime. Legacy and Spot preemptible nodes cannot be used in a node pool with the `CONTROLLER` [role](/dataproc/docs/reference/rest/v1/projects.regions.clusters#role) or in the DEFAULT node pool if the CONTROLLER role is not assigned (the DEFAULT node pool will assume the CONTROLLER role). */
	Spot *bool `json:"spot,omitempty"`
}

type DataprocClusterAccelerators_GKENodePoolAcceleratorConfigSpec struct {
	/* Immutable. The number of accelerator cards exposed to an instance. */
	AcceleratorCount *int64 `json:"acceleratorCount,omitempty"`

	/* Immutable. The accelerator type resource namename (see GPUs on Compute Engine). */
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	/* Immutable. Size of partitions to create on the GPU. Valid values are described in the NVIDIA [mig user guide](https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning). */
	GpuPartitionSize *string `json:"gpuPartitionSize,omitempty"`
}

type DataprocClusterEphemeralStorageConfigSpec struct {
	/* Immutable. Number of local SSDs to use to back ephemeral storage. Uses NVMe interfaces. Each local SSD is 375 GB in size. If zero, it means to disable using local SSDs as ephemeral storage. */
	LocalSsdCount *int64 `json:"localSsdCount,omitempty"`
}

type DataprocClusterKubernetesSoftwareConfigSpec struct {
	/* Immutable. The components that should be installed in this Dataproc cluster. The key must be a string from the KubernetesComponent enumeration. The value is the version of the software to be installed. At least one entry must be specified. */
	ComponentVersion map[string]string `json:"componentVersion,omitempty"`

	/* Immutable. The properties to set on daemon config files. Property keys are specified in `prefix:property` format, for example `spark:spark.kubernetes.container.image`. The following are supported prefixes and their mappings: * spark: `spark-defaults.conf` For more information, see [Cluster properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties). */
	Properties map[string]string `json:"properties,omitempty"`
}

type DataprocClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// Output only. A cluster UUID (Unique Universal Identifier). Dataproc generates this value when it creates the cluster.
	ClusterUuid *string `json:"clusterUuid,omitempty"`

	// Output only. Cluster status.
	Status *ClusterStatusStatus `json:"status,omitempty"`

	// Output only. The previous cluster status.
	StatusHistory []ClusterStatusStatus `json:"statusHistory,omitempty"`

	// Output only. The cluster config for a cluster of Compute Engine Instances.
	Config *ClusterConfigStatus `json:"config,omitempty"`

	// Output only. Contains cluster daemon metrics such as HDFS and YARN stats. **Beta Feature**: This report is available for testing purposes only. It may be changed before final release.
	Metrics *ClusterMetricsStatus `json:"metrics,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

type ClusterStatusStatus struct {
	/* Optional. Output only. Details of cluster's state. */
	Detail *string `json:"detail,omitempty"`

	/* Output only. The cluster's state. Possible values: UNKNOWN, CREATING, RUNNING, ERROR, DELETING, UPDATING, STOPPING, STOPPED, STARTING */
	State *string `json:"state,omitempty"`

	/* Output only. Time when this state was entered (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	// +kubebuilder:validation:Format=date-time
	StateStartTime *string `json:"stateStartTime,omitempty"`

	/* Output only. Additional state information that includes status reported by the agent. Possible values: UNSPECIFIED, UNHEALTHY, STALE_STATUS */
	Substate *string `json:"substate,omitempty"`
}

type ClusterConfigStatus struct {
	/* Output only. Port/endpoint configuration for this cluster */
	EndpointConfig *EndpointConfigStatus `json:"endpointConfig,omitempty"`

	/* Output only. Lifecycle setting for the cluster. */
	LifecycleConfig *LifecycleConfigStatus `json:"lifecycleConfig,omitempty"`

	/* Output only. The Compute Engine config settings for the cluster's master instance. */
	MasterConfig *InstanceGroupConfigStatus `json:"masterConfig,omitempty"`

	/* Output only. The Compute Engine config settings for a cluster's secondary worker instances. */
	SecondaryWorkerConfig *InstanceGroupConfigStatus `json:"secondaryWorkerConfig,omitempty"`

	/* Output only. The Compute Engine config settings for the cluster's worker instances. */
	WorkerConfig *InstanceGroupConfigStatus `json:"workerConfig,omitempty"`
}

type EndpointConfigStatus struct {
	/* Output only. The map of port descriptions to URLs. Will only be populated if enable_http_port_access is true. */
	HttpPorts map[string]string `json:"httpPorts,omitempty"`
}

type LifecycleConfigStatus struct {
	/* Output only. The time when cluster became idle (most recent job finished) and became eligible for deletion due to idleness (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	// +kubebuilder:validation:Format=date-time
	IdleStartTime *string `json:"idleStartTime,omitempty"`
}

type InstanceGroupConfigStatus struct {
	/* Output only. The list of instance names. Dataproc derives the names from `cluster_name`, `num_instances`, and the instance group. */
	InstanceNames []string `json:"instanceNames,omitempty"`

	/* Output only. List of references to Compute Engine instances. */
	InstanceReferences []InstanceReferenceStatus `json:"instanceReferences,omitempty"`

	/* Output only. Specifies that this instance group contains preemptible instances. */
	IsPreemptible *bool `json:"isPreemptible,omitempty"`

	/* Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups. */
	ManagedGroupConfig *ManagedGroupConfigStatus `json:"managedGroupConfig,omitempty"`
}

type InstanceReferenceStatus struct {
	/* The unique identifier of the Compute Engine instance. */
	InstanceId *string `json:"instanceId,omitempty"`

	/* The user-friendly name of the Compute Engine instance. */
	InstanceName *string `json:"instanceName,omitempty"`

	/* The public ECIES key used for sharing data with this instance. */
	PublicEciesKey *string `json:"publicEciesKey,omitempty"`

	/* The public RSA key used for sharing data with this instance. */
	PublicKey *string `json:"publicKey,omitempty"`
}

type ManagedGroupConfigStatus struct {
	/* Output only. The name of the Instance Group Manager for this group. */
	InstanceGroupManagerName *string `json:"instanceGroupManagerName,omitempty"`

	/* Output only. The name of the Instance Template used for the Managed Instance Group. */
	InstanceTemplateName *string `json:"instanceTemplateName,omitempty"`
}

type ClusterMetricsStatus struct {
	/* The HDFS metrics. */
	HdfsMetrics map[string]string `json:"hdfsMetrics,omitempty"`

	/* The YARN metrics. */
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
