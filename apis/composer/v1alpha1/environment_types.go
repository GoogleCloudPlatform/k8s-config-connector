// Copyright 2025 Google LLC
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

package v1alpha1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	computev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComposerEnvironmentGVK = GroupVersion.WithKind("ComposerEnvironment")

type Parent struct {
	/* Immutable. The Project that this resource belongs to. */
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. The name of the location where the Environment will be created.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location is immutable."
	// Required.
	Location string `json:"location"`
}

// ComposerEnvironmentSpec defines the desired state of ComposerEnvironment
// +kcc:proto=google.cloud.orchestration.airflow.service.v1.Environment
type ComposerEnvironmentSpec struct {
	Parent `json:",inline"`

	// The ComposerEnvironment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Configuration parameters for this environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.config
	Config *EnvironmentConfig `json:"config,omitempty"`

	// Optional. User-defined labels for this environment.
	//  The labels map can contain no more than 64 entries. Entries of the labels
	//  map are UTF8 strings that comply with the following restrictions:
	//
	//  * Keys must conform to regexp: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	//  * Values must conform to regexp:  [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	//  * Both keys and values are additionally constrained to be <= 128 bytes in
	//  size.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Storage configuration for this environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.storage_config
	StorageConfig *StorageConfig `json:"storageConfig,omitempty"`
}

// ComposerEnvironmentStatus defines the config connector machine state of ComposerEnvironment
type ComposerEnvironmentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComposerEnvironment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComposerEnvironmentObservedState `json:"observedState,omitempty"`
}

// ComposerEnvironmentObservedState is the state of the ComposerEnvironment resource as most recently observed in GCP.
// +kcc:proto=google.cloud.orchestration.airflow.service.v1.Environment
type ComposerEnvironmentObservedState struct {
	// Optional. Configuration parameters for this environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.config
	Config *EnvironmentConfigObservedState `json:"config,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// Output only. The UUID (Universally Unique IDentifier) associated with this
	//  environment. This value is generated when the environment is created.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.uuid
	Uuid *string `json:"uuid,omitempty"`

	// The current state of the environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.state
	State *string `json:"state,omitempty"`

	// Output only. The time at which this environment was created.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this environment was last modified.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.Environment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomposerenvironment;gcpcomposerenvironments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComposerEnvironment is the Schema for the ComposerEnvironment API
// +k8s:openapi-gen=true
type ComposerEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComposerEnvironmentSpec   `json:"spec,omitempty"`
	Status ComposerEnvironmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComposerEnvironmentList contains a list of ComposerEnvironment
type ComposerEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComposerEnvironment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComposerEnvironment{}, &ComposerEnvironmentList{})
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig
type EnvironmentConfig struct {
	// The number of nodes in the Kubernetes Engine cluster that will be
	//  used to run this environment.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-1.*.*-airflow-*.*.*.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Optional. The configuration settings for software inside the environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.software_config
	SoftwareConfig *SoftwareConfig `json:"softwareConfig,omitempty"`

	// Optional. The configuration used for the Kubernetes Engine cluster.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.node_config
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	// Optional. The configuration used for the Private IP Cloud Composer
	//  environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.private_environment_config
	PrivateEnvironmentConfig *PrivateEnvironmentConfig `json:"privateEnvironmentConfig,omitempty"`

	// Optional. The network-level access control policy for the Airflow web
	//  server. If unspecified, no network-level access restrictions will be
	//  applied.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.web_server_network_access_control
	WebServerNetworkAccessControl *WebServerNetworkAccessControl `json:"webServerNetworkAccessControl,omitempty"`

	// Optional. The configuration settings for Cloud SQL instance used internally
	//  by Apache Airflow software.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.database_config
	DatabaseConfig *DatabaseConfig `json:"databaseConfig,omitempty"`

	// Optional. The configuration settings for the Airflow web server App Engine
	//  instance.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.web_server_config
	WebServerConfig *WebServerConfig `json:"webServerConfig,omitempty"`

	// Optional. The encryption options for the Cloud Composer environment
	//  and its dependencies. Cannot be updated.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// Optional. The maintenance window is the period when Cloud Composer
	//  components may undergo maintenance. It is defined so that maintenance is
	//  not executed during peak hours or critical time periods.
	//
	//  The system will not be under maintenance for every occurrence of this
	//  window, but when maintenance is planned, it will be scheduled
	//  during the window.
	//
	//  The maintenance window period must encompass at least 12 hours per week.
	//  This may be split into multiple chunks, each with a size of
	//  at least 4 hours.
	//
	//  If this value is omitted, the default value for maintenance window is
	//  applied. By default, maintenance windows are from 00:00:00 to 04:00:00
	//  (GMT) on Friday, Saturday, and Sunday every week.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.maintenance_window
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`

	// Optional. The workloads configuration settings for the GKE cluster
	//  associated with the Cloud Composer environment. The GKE cluster runs
	//  Airflow scheduler, web server and workers workloads.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-2.*.*-airflow-*.*.* and newer.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.workloads_config
	WorkloadsConfig *WorkloadsConfig `json:"workloadsConfig,omitempty"`

	// Optional. The size of the Cloud Composer environment.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-2.*.*-airflow-*.*.* and newer.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.environment_size
	EnvironmentSize *string `json:"environmentSize,omitempty"`

	// Optional. The configuration options for GKE cluster master authorized
	//  networks. By default master authorized networks feature is:
	//  - in case of private environment: enabled with no external networks
	//  allowlisted.
	//  - in case of public environment: disabled.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.master_authorized_networks_config
	MasterAuthorizedNetworksConfig *MasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig,omitempty"`

	// Optional. The Recovery settings configuration of an environment.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-2.*.*-airflow-*.*.* and newer.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.recovery_config
	RecoveryConfig *RecoveryConfig `json:"recoveryConfig,omitempty"`

	// Optional. Resilience mode of the Cloud Composer Environment.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-2.2.0-airflow-*.*.* and newer.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.resilience_mode
	ResilienceMode *string `json:"resilienceMode,omitempty"`

	// Optional. The configuration setting for Airflow database data retention
	//  mechanism.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.data_retention_config
	DataRetentionConfig *DataRetentionConfig `json:"dataRetentionConfig,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig
type EnvironmentConfigObservedState struct {
	// Output only. The Kubernetes Engine cluster used to run this environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.gke_cluster
	GkeCluster *string `json:"gkeCluster,omitempty"`

	// Output only. The Cloud Storage prefix of the DAGs for this environment.
	//  Although Cloud Storage objects reside in a flat namespace, a hierarchical
	//  file tree can be simulated using "/"-delimited object name prefixes. DAG
	//  objects for this environment reside in a simulated directory with the given
	//  prefix.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.dag_gcs_prefix
	DagGCSPrefix *string `json:"dagGCSPrefix,omitempty"`

	// Optional. The configuration used for the Private IP Cloud Composer
	//  environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.private_environment_config
	PrivateEnvironmentConfig *PrivateEnvironmentConfigObservedState `json:"privateEnvironmentConfig,omitempty"`

	// Output only. The 'bring your own identity' variant of the URI of the Apache
	//  Airflow Web UI hosted within this environment, to be accessed with external
	//  identities using workforce identity federation (see [Access environments
	//  with workforce identity
	//  federation](/composer/docs/composer-2/access-environments-with-workforce-identity-federation)).
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.airflow_byoid_uri
	AirflowBYOIDURI *string `json:"airflowBYOIDURI,omitempty"`

	// Output only. The URI of the Apache Airflow Web UI hosted within this
	//  environment (see [Airflow web
	//  interface](/composer/docs/how-to/accessing/airflow-web-interface)).
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EnvironmentConfig.airflow_uri
	AirflowURI *string `json:"airflowURI,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.EncryptionConfig
type EncryptionConfig struct {
	// Optional. Customer-managed Encryption Key available through Google's Key
	//  Management Service. Cannot be updated. If not specified, Google-managed key
	//  will be used.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.EncryptionConfig.kms_key_name
	KMSKeyRef *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.NodeConfig
type NodeConfig struct {
	// Optional. The Compute Engine [zone](/compute/docs/regions-zones) in which
	//  to deploy the VMs used to run the Apache Airflow software, specified as a
	//  [relative resource
	//  name](/apis/design/resource_names#relative_resource_name). For example:
	//  "projects/{projectId}/zones/{zoneId}".
	//
	//  This `location` must belong to the enclosing environment's project and
	//  location. If both this field and `nodeConfig.machineType` are specified,
	//  `nodeConfig.machineType` must belong to this `location`; if both are
	//  unspecified, the service will pick a zone in the Compute Engine region
	//  corresponding to the Cloud Composer location, and propagate that choice to
	//  both fields. If only one field (`location` or `nodeConfig.machineType`) is
	//  specified, the location information from the specified field will be
	//  propagated to the unspecified field.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-1.*.*-airflow-*.*.*.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.NodeConfig.location
	Location *string `json:"location,omitempty"`

	// Optional. The Compute Engine
	//  [machine type](/compute/docs/machine-types) used for cluster instances,
	//  specified as a
	//  [relative resource
	//  name](/apis/design/resource_names#relative_resource_name). For example:
	//  "projects/{projectId}/zones/{zoneId}/machineTypes/{machineTypeId}".
	//
	//  The `machineType` must belong to the enclosing environment's project and
	//  location. If both this field and `nodeConfig.location` are specified,
	//  this `machineType` must belong to the `nodeConfig.location`; if both are
	//  unspecified, the service will pick a zone in the Compute Engine region
	//  corresponding to the Cloud Composer location, and propagate that choice to
	//  both fields. If exactly one of this field and `nodeConfig.location` is
	//  specified, the location information from the specified field will be
	//  propagated to the unspecified field.
	//
	//  The `machineTypeId` must not be a [shared-core machine
	//  type](/compute/docs/machine-types#sharedcore).
	//
	//  If this field is unspecified, the `machineTypeId` defaults
	//  to "n1-standard-1".
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-1.*.*-airflow-*.*.*.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.NodeConfig.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Optional. The Compute Engine network to be used for machine
	//  communications, specified as a
	//  [relative resource
	//  name](/apis/design/resource_names#relative_resource_name). For example:
	//  "projects/{projectId}/global/networks/{networkId}".
	//
	//  If unspecified, the "default" network ID in the environment's project is
	//  used. If a [Custom Subnet Network](/vpc/docs/vpc#vpc_networks_and_subnets)
	//  is provided, `nodeConfig.subnetwork` must also be provided. For
	//  [Shared VPC](/vpc/docs/shared-vpc) subnetwork requirements, see
	//  `nodeConfig.subnetwork`.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.NodeConfig.network
	NetworkRef *refs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. The Compute Engine subnetwork to be used for machine
	//  communications, specified as a
	//  [relative resource
	//  name](/apis/design/resource_names#relative_resource_name). For example:
	//  "projects/{projectId}/regions/{regionId}/subnetworks/{subnetworkId}"
	//
	//  If a subnetwork is provided, `nodeConfig.network` must also be provided,
	//  and the subnetwork must belong to the enclosing environment's project and
	//  location.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.NodeConfig.subnetwork
	SubnetworkRef *refs.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	// Optional. The disk size in GB used for node VMs. Minimum size is 30GB.
	//  If unspecified, defaults to 100GB. Cannot be updated.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-1.*.*-airflow-*.*.*.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.NodeConfig.disk_size_gb
	DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`

	// Optional. The set of Google API scopes to be made available on all
	//  node VMs. If `oauth_scopes` is empty, defaults to
	//  ["https://www.googleapis.com/auth/cloud-platform"]. Cannot be updated.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-1.*.*-airflow-*.*.*.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.NodeConfig.oauth_scopes
	OAuthScopes []string `json:"oauthScopes,omitempty"`

	// Optional. The Google Cloud Platform Service Account to be used by the node
	//  VMs. If a service account is not specified, the "default" Compute Engine
	//  service account is used. Cannot be updated.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.NodeConfig.service_account
	ServiceAccountRef *refs.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. The list of instance tags applied to all node VMs. Tags are used
	//  to identify valid sources or targets for network firewalls. Each tag within
	//  the list must comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  Cannot be updated.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.NodeConfig.tags
	Tags []string `json:"tags,omitempty"`

	// Optional. The configuration for controlling how IPs are allocated in the
	//  GKE cluster.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.NodeConfig.ip_allocation_policy
	IPAllocationPolicy *IPAllocationPolicy `json:"ipAllocationPolicy,omitempty"`

	// Optional. Deploys 'ip-masq-agent' daemon set in the GKE cluster and defines
	//  nonMasqueradeCIDRs equals to pod IP range so IP masquerading is used for
	//  all destination addresses, except between pods traffic.
	//
	//  See:
	//  https://cloud.google.com/kubernetes-engine/docs/how-to/ip-masquerade-agent
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.NodeConfig.enable_ip_masq_agent
	EnableIPMasqAgent *bool `json:"enableIPMasqAgent,omitempty"`

	// Optional. Network Attachment that Cloud Composer environment is connected
	//  to, which provides connectivity with a user's VPC network. Takes precedence
	//  over network and subnetwork settings. If not provided, but network and
	//  subnetwork are defined during environment, it will be provisioned. If not
	//  provided and network and subnetwork are also empty, then connectivity to
	//  user's VPC network is disabled. Network attachment must be provided in
	//  format
	//  projects/{project}/regions/{region}/networkAttachments/{networkAttachment}.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-3.*.*-airflow-*.*.* and newer.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.NodeConfig.composer_network_attachment
	ComposerNetworkAttachmentRef *computev1alpha1.ComputeNetworkAttachmentRef `json:"composerNetworkAttachmentRef,omitempty"`

	// Optional. The IP range in CIDR notation to use internally by Cloud
	//  Composer. IP addresses are not reserved - and the same range can be used by
	//  multiple Cloud Composer environments. In case of overlap, IPs from this
	//  range will not be accessible in the user's VPC network. Cannot be updated.
	//  If not specified, the default value of '100.64.128.0/20' is used.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-3.*.*-airflow-*.*.* and newer.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.NodeConfig.composer_internal_ipv4_cidr_block
	ComposerInternalIPv4CIDRBlock *string `json:"composerInternalIPv4CIDRBlock,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig
type PrivateEnvironmentConfig struct {
	// Optional. If `true`, a Private IP Cloud Composer environment is created.
	//  If this field is set to true, `IPAllocationPolicy.use_ip_aliases` must be
	//  set to true for Cloud Composer environments in versions
	//  composer-1.*.*-airflow-*.*.*.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig.enable_private_environment
	EnablePrivateEnvironment *bool `json:"enablePrivateEnvironment,omitempty"`

	// Optional. If `true`, builds performed during operations that install Python
	//  packages have only private connectivity to Google services (including
	//  Artifact Registry) and VPC network (if either `NodeConfig.network` and
	//  `NodeConfig.subnetwork` fields or `NodeConfig.composer_network_attachment`
	//  field are specified). If `false`, the builds also have access to the
	//  internet.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-3.*.*-airflow-*.*.* and newer.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig.enable_private_builds_only
	EnablePrivateBuildsOnly *bool `json:"enablePrivateBuildsOnly,omitempty"`

	// Optional. Configuration for the private GKE cluster for a Private IP
	//  Cloud Composer environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig.private_cluster_config
	PrivateClusterConfig *PrivateClusterConfig `json:"privateClusterConfig,omitempty"`

	// Optional. The CIDR block from which IP range for web server will be
	//  reserved. Needs to be disjoint from
	//  `private_cluster_config.master_ipv4_cidr_block` and
	//  `cloud_sql_ipv4_cidr_block`.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-1.*.*-airflow-*.*.*.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig.web_server_ipv4_cidr_block
	WebServerIPv4CIDRBlock *string `json:"webServerIPv4CIDRBlock,omitempty"`

	// Optional. The CIDR block from which IP range in tenant project will be
	//  reserved for Cloud SQL. Needs to be disjoint from
	//  `web_server_ipv4_cidr_block`.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig.cloud_sql_ipv4_cidr_block
	CloudSQLIPv4CIDRBlock *string `json:"cloudSQLIPv4CIDRBlock,omitempty"`

	// Optional. The CIDR block from which IP range for Cloud Composer Network in
	//  tenant project will be reserved. Needs to be disjoint from
	//  private_cluster_config.master_ipv4_cidr_block and
	//  cloud_sql_ipv4_cidr_block.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-2.*.*-airflow-*.*.* and newer.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig.cloud_composer_network_ipv4_cidr_block
	CloudComposerNetworkIPv4CIDRBlock *string `json:"cloudComposerNetworkIPv4CIDRBlock,omitempty"`

	// Optional. When enabled, IPs from public (non-RFC1918) ranges can be used
	//  for `IPAllocationPolicy.cluster_ipv4_cidr_block` and
	//  `IPAllocationPolicy.service_ipv4_cidr_block`.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig.enable_privately_used_public_ips
	EnablePrivatelyUsedPublicIPs *bool `json:"enablePrivatelyUsedPublicIPs,omitempty"`

	// Optional. When specified, the environment will use Private Service Connect
	//  instead of VPC peerings to connect to Cloud SQL in the Tenant Project,
	//  and the PSC endpoint in the Customer Project will use an IP address from
	//  this subnetwork.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig.cloud_composer_connection_subnetwork
	CloudComposerConnectionSubnetworkRef *refs.ComputeSubnetworkRef `json:"cloudComposerConnectionSubnetworkRef,omitempty"`

	// Optional. Configuration for the network connections configuration in the
	//  environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig.networking_config
	NetworkingConfig *NetworkingConfig `json:"networkingConfig,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig
type PrivateEnvironmentConfigObservedState struct {
	// Optional. Configuration for the private GKE cluster for a Private IP
	//  Cloud Composer environment.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig.private_cluster_config
	PrivateClusterConfig *PrivateClusterConfigObservedState `json:"privateClusterConfig,omitempty"`

	// Output only. The IP range reserved for the tenant project's App Engine VMs.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-1.*.*-airflow-*.*.*.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig.web_server_ipv4_reserved_range
	WebServerIPv4ReservedRange *string `json:"webServerIPv4ReservedRange,omitempty"`

	// Output only. The IP range reserved for the tenant project's Cloud Composer
	//  network.
	//
	//  This field is supported for Cloud Composer environments in versions
	//  composer-2.*.*-airflow-*.*.* and newer.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.PrivateEnvironmentConfig.cloud_composer_network_ipv4_reserved_range
	CloudComposerNetworkIPv4ReservedRange *string `json:"cloudComposerNetworkIPv4ReservedRange,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.StorageConfig
type StorageConfig struct {
	// Optional. The name of the Cloud Storage bucket used by the environment. No
	//  `gs://` prefix.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.StorageConfig.bucket
	BucketRef *refs.StorageBucketRef `json:"bucketRef,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.DagProcessorResource
type WorkloadsConfig_DagProcessorResource struct {
	// Optional. CPU request and limit for a single Airflow DAG processor
	//  replica.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.DagProcessorResource.cpu
	CPU *string `json:"cpu,omitempty"`

	// Optional. Memory (GB) request and limit for a single Airflow DAG
	//  processor replica.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.DagProcessorResource.memory_gb
	MemoryGB *string `json:"memoryGB,omitempty"`

	// Optional. Storage (GB) request and limit for a single Airflow DAG
	//  processor replica.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.DagProcessorResource.storage_gb
	StorageGB *string `json:"storageGB,omitempty"`

	// Optional. The number of DAG processors. If not provided or set to 0, a
	//  single DAG processor instance will be created.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.DagProcessorResource.count
	Count *int32 `json:"count,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.SchedulerResource
type WorkloadsConfig_SchedulerResource struct {
	// Optional. CPU request and limit for a single Airflow scheduler replica.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.SchedulerResource.cpu
	CPU *string `json:"cpu,omitempty"`

	// Optional. Memory (GB) request and limit for a single Airflow scheduler
	//  replica.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.SchedulerResource.memory_gb
	MemoryGB *string `json:"memoryGB,omitempty"`

	// Optional. Storage (GB) request and limit for a single Airflow scheduler
	//  replica.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.SchedulerResource.storage_gb
	StorageGB *string `json:"storageGB,omitempty"`

	// Optional. The number of schedulers.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.SchedulerResource.count
	Count *int32 `json:"count,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.TriggererResource
type WorkloadsConfig_TriggererResource struct {
	// Optional. The number of triggerers.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.TriggererResource.count
	Count *int32 `json:"count,omitempty"`

	// Optional. CPU request and limit for a single Airflow triggerer replica.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.TriggererResource.cpu
	CPU *string `json:"cpu,omitempty"`

	// Optional. Memory (GB) request and limit for a single Airflow triggerer
	//  replica.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.TriggererResource.memory_gb
	MemoryGB *string `json:"memoryGB,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.WebServerResource
type WorkloadsConfig_WebServerResource struct {
	// Optional. CPU request and limit for Airflow web server.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.WebServerResource.cpu
	CPU *string `json:"cpu,omitempty"`

	// Optional. Memory (GB) request and limit for Airflow web server.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.WebServerResource.memory_gb
	MemoryGB *string `json:"memoryGB,omitempty"`

	// Optional. Storage (GB) request and limit for Airflow web server.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.WebServerResource.storage_gb
	StorageGB *string `json:"storageGB,omitempty"`
}

// +kcc:proto=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.WorkerResource
type WorkloadsConfig_WorkerResource struct {
	// Optional. CPU request and limit for a single Airflow worker replica.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.WorkerResource.cpu
	CPU *string `json:"cpu,omitempty"`

	// Optional. Memory (GB) request and limit for a single Airflow worker
	//  replica.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.WorkerResource.memory_gb
	MemoryGB *string `json:"memoryGB,omitempty"`

	// Optional. Storage (GB) request and limit for a single Airflow worker
	//  replica.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.WorkerResource.storage_gb
	StorageGB *string `json:"storageGB,omitempty"`

	// Optional. Minimum number of workers for autoscaling.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.WorkerResource.min_count
	MinCount *int32 `json:"minCount,omitempty"`

	// Optional. Maximum number of workers for autoscaling.
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.WorkloadsConfig.WorkerResource.max_count
	MaxCount *int32 `json:"maxCount,omitempty"`
}
