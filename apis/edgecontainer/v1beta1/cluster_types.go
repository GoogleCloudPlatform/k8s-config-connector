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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EdgeContainerClusterGVK = GroupVersion.WithKind("EdgeContainerCluster")

// EdgeContainerClusterSpec defines the desired state of EdgeContainerCluster
// +kcc:proto=google.cloud.edgecontainer.v1.Cluster
type EdgeContainerClusterSpec struct {

	// Labels associated with this resource.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Fleet configuration.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.fleet
	Fleet *Fleet `json:"fleet,omitempty"`

	// Required. Cluster-wide networking configuration.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.networking
	Networking *ClusterNetworking `json:"networking,omitempty"`

	// Required. Immutable. RBAC policy that will be applied and managed by GEC.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.authorization
	Authorization *Authorization `json:"authorization,omitempty"`

	// Optional. The default maximum number of pods per node used if a maximum
	//  value is not specified explicitly for a node pool in this cluster. If
	//  unspecified, the Kubernetes default value will be used.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.default_max_pods_per_node
	DefaultMaxPodsPerNode *int32 `json:"defaultMaxPodsPerNode,omitempty"`

	// Optional. Cluster-wide maintenance policy configuration.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.maintenance_policy
	MaintenancePolicy *MaintenancePolicy `json:"maintenancePolicy,omitempty"`

	// Optional. The configuration of the cluster control plane.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.control_plane
	ControlPlane *Cluster_ControlPlane `json:"controlPlane,omitempty"`

	// Optional. The configuration of the system add-ons.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.system_addons_config
	SystemAddonsConfig *Cluster_SystemAddonsConfig `json:"systemAddonsConfig,omitempty"`

	// Optional. IPv4 address pools for cluster data plane external load
	//  balancing.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.external_load_balancer_ipv4_address_pools
	ExternalLoadBalancerIPV4AddressPools []string `json:"externalLoadBalancerIPV4AddressPools,omitempty"`

	// Optional. Remote control plane disk encryption options. This field is only
	//  used when enabling CMEK support.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.control_plane_encryption
	ControlPlaneEncryption *Cluster_ControlPlaneEncryption `json:"controlPlaneEncryption,omitempty"`

	// Optional. The target cluster version. For example: "1.5.0".
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.target_version
	TargetVersion *string `json:"targetVersion,omitempty"`

	// Optional. The release channel a cluster is subscribed to.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.release_channel
	ReleaseChannel *string `json:"releaseChannel,omitempty"`

	// Optional. Configuration of the cluster survivability, e.g., for the case
	//  when network connectivity is lost. Note: This only applies to local control
	//  plane clusters.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.survivability_config
	SurvivabilityConfig *Cluster_SurvivabilityConfig `json:"survivabilityConfig,omitempty"`

	// Optional. IPv6 address pools for cluster data plane external load
	//  balancing.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.external_load_balancer_ipv6_address_pools
	ExternalLoadBalancerIPV6AddressPools []string `json:"externalLoadBalancerIPV6AddressPools,omitempty"`

	*Parent `json:",inline"`
	// The EdgeContainerCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// EdgeContainerClusterStatus defines the config connector machine state of EdgeContainerCluster
type EdgeContainerClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the EdgeContainerCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *EdgeContainerClusterObservedState `json:"observedState,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Authorization
type Authorization struct {
	// Required. User that will be granted the cluster-admin role on the cluster,
	//  providing full access to the cluster. Currently, this is a singular field,
	//  but will be expanded to allow multiple admins in the future.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Authorization.admin_users
	AdminUsers *ClusterUser `json:"adminUsers,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ConnectionState
type Cluster_ConnectionState struct {
}

// +kcc:proto=google.cloud.edgecontainer.v1.ClusterNetworking
type ClusterNetworking struct {
	// Required. All pods in the cluster are assigned an RFC1918 IPv4 address from
	//  these blocks. Only a single block is supported. This field cannot be
	//  changed after creation.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.ClusterNetworking.cluster_ipv4_cidr_blocks
	ClusterIPV4CIDRBlocks []string `json:"clusterIPV4CIDRBlocks,omitempty"`

	// Required. All services in the cluster are assigned an RFC1918 IPv4 address
	//  from these blocks. Only a single block is supported. This field cannot be
	//  changed after creation.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.ClusterNetworking.services_ipv4_cidr_blocks
	ServicesIPV4CIDRBlocks []string `json:"servicesIPV4CIDRBlocks,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.MaintenanceExclusionWindow
type MaintenanceExclusionWindow struct {
	// Optional. The time window.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.MaintenanceExclusionWindow.window
	Window *TimeWindow `json:"window,omitempty"`

	// Optional. A unique (per cluster) id for the window.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.MaintenanceExclusionWindow.id
	ID *string `json:"id,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.TimeWindow
type TimeWindow struct {
	// The time that the window first starts.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.TimeWindow.start_time
	StartTime *string `json:"startTime,omitempty"`

	// The time that the window ends. The end time must take place after the
	//  start time.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.TimeWindow.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.MaintenancePolicy
type MaintenancePolicy struct {
	// Specifies the maintenance window in which maintenance may be performed.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.MaintenancePolicy.window
	Window *MaintenanceWindow `json:"window,omitempty"`

	// Optional. Exclusions to automatic maintenance. Non-emergency maintenance
	//  should not occur in these windows. Each exclusion has a unique name and may
	//  be active or expired. The max number of maintenance exclusions allowed at a
	//  given time is 3.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.MaintenancePolicy.maintenance_exclusions
	MaintenanceExclusions []MaintenanceExclusionWindow `json:"maintenanceExclusions,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.MaintenanceWindow
type MaintenanceWindow struct {
	// Configuration of a recurring maintenance window.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.MaintenanceWindow.recurring_window
	RecurringWindow *RecurringTimeWindow `json:"recurringWindow,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.RecurringTimeWindow
type RecurringTimeWindow struct {
	// The window of the first recurrence.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.RecurringTimeWindow.window
	Window *TimeWindow `json:"window,omitempty"`

	// An RRULE (https://tools.ietf.org/html/rfc5545#section-3.8.5.3) for how
	//  this window recurs. They go on for the span of time between the start and
	//  end time.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.RecurringTimeWindow.recurrence
	Recurrence *string `json:"recurrence,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.ClusterUser
type ClusterUser struct {
	// Required. An active Google username.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.ClusterUser.username
	Username *string `json:"username,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Fleet
type Fleet struct {
	// Required. The name of the Fleet host project where this cluster will be
	//  registered.
	//
	//  Project names are formatted as
	//  `projects/<project-number>`.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Fleet.project
	ProjectRef *v1beta1.ProjectRef `json:"projectRef,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ControlPlane
type Cluster_ControlPlane struct {
	// Remote control plane configuration.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.remote
	Remote *Cluster_ControlPlane_Remote `json:"remote,omitempty"`

	// Local control plane configuration.
	//
	//  Warning: Local control plane clusters must be created in their own
	//  project. Local control plane clusters cannot coexist in the same
	//  project with any other type of clusters, including non-GDCE clusters.
	//  Mixing local control plane GDCE clusters with any other type of
	//  clusters in the same project can result in data loss.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.local
	Local *Cluster_ControlPlane_Local `json:"local,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Local
type Cluster_ControlPlane_Local struct {
	// Name of the Google Distributed Cloud Edge zones where this node pool
	//  will be created. For example: `us-central1-edge-customer-a`.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Local.node_location
	NodeLocation *string `json:"nodeLocation,omitempty"`

	// The number of nodes to serve as replicas of the Control Plane.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Local.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Only machines matching this filter will be allowed to host control
	//  plane nodes. The filtering language accepts strings like "name=<name>",
	//  and is documented here: [AIP-160](https://google.aip.dev/160).
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Local.machine_filter
	MachineFilter *string `json:"machineFilter,omitempty"`

	// Policy configuration about how user applications are deployed.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Local.shared_deployment_policy
	SharedDeploymentPolicy *string `json:"sharedDeploymentPolicy,omitempty"`

	// Optional. Name for the storage schema of control plane nodes.
	//
	//  Warning: Configurable node local storage schema feature is an
	//  experimental feature, and is not recommended for general use
	//  in production clusters/nodepools.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Local.control_plane_node_storage_schema
	ControlPlaneNodeStorageSchema *string `json:"controlPlaneNodeStorageSchema,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Remote
type Cluster_ControlPlane_Remote struct {
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ControlPlaneEncryption
type Cluster_ControlPlaneEncryption struct {
	// Optional. The Cloud KMS CryptoKey e.g.
	//  projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
	//  to use for protecting control plane disks. If not specified, a
	//  Google-managed key will be used instead.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlaneEncryption.kms_key
	KMSKeyRef *v1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent
type Cluster_MaintenanceEvent struct {
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.SurvivabilityConfig
type Cluster_SurvivabilityConfig struct {
	// Optional. Time period that allows the cluster nodes to be rebooted and
	//  become functional without network connectivity to Google. The default 0
	//  means not allowed. The maximum is 7 days.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.SurvivabilityConfig.offline_reboot_ttl
	OfflineRebootTTL *string `json:"offlineRebootTTL,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig
type Cluster_SystemAddonsConfig struct {
	// Optional. Config for Ingress.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig.ingress
	Ingress *Cluster_SystemAddonsConfig_Ingress `json:"ingress,omitempty"`

	// Optional. Config for VM Service.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig.vm_service_config
	VmServiceConfig *Cluster_SystemAddonsConfig_VmServiceConfig `json:"vmServiceConfig,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig.Ingress
type Cluster_SystemAddonsConfig_Ingress struct {
	// Optional. Whether Ingress is disabled.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig.Ingress.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Optional. Ingress VIP.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig.Ingress.ipv4_vip
	IPV4Vip *string `json:"ipv4Vip,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig.VMServiceConfig
type Cluster_SystemAddonsConfig_VmServiceConfig struct {
	// Optional. Whether VMM is enabled.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig.VMServiceConfig.vmm_enabled
	VmmEnabled *bool `json:"vmmEnabled,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ConnectionState
type Cluster_ConnectionStateObservedState struct {
	// Output only. The current connection state.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ConnectionState.state
	State *string `json:"state,omitempty"`

	// Output only. The time when the connection state was last changed.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ConnectionState.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ControlPlaneEncryption
type Cluster_ControlPlaneEncryptionObservedState struct {
	// Output only. The Cloud KMS CryptoKeyVersion currently in use for
	//  protecting control plane disks. Only applicable if kms_key is set.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlaneEncryption.kms_key_active_version
	KMSKeyActiveVersion *string `json:"kmsKeyActiveVersion,omitempty"`

	// Output only. Availability of the Cloud KMS CryptoKey. If not
	//  `KEY_AVAILABLE`, then nodes may go offline as they cannot access their
	//  local data. This can be caused by a lack of permissions to use the key,
	//  or if the key is disabled or deleted.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlaneEncryption.kms_key_state
	KMSKeyState *string `json:"kmsKeyState,omitempty"`

	// Output only. Error status returned by Cloud KMS when using this key. This
	//  field may be populated only if `kms_key_state` is not
	//  `KMS_KEY_STATE_KEY_AVAILABLE`. If populated, this field contains the
	//  error status reported by Cloud KMS.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlaneEncryption.kms_status
	KMSStatus *Status `json:"kmsStatus,omitempty"`

	// Output only. The current resource state associated with the cmek.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlaneEncryption.resource_state
	ResourceState *string `json:"resourceState,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent
type Cluster_MaintenanceEventObservedState struct {
	// Output only. UUID of the maintenance event.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.uuid
	Uuid *string `json:"uuid,omitempty"`

	// Output only. The target version of the cluster.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.target_version
	TargetVersion *string `json:"targetVersion,omitempty"`

	// Output only. The operation for running the maintenance event. Specified
	//  in the format projects/*/locations/*/operations/*. If the maintenance
	//  event is split into multiple operations (e.g. due to maintenance
	//  windows), the latest one is recorded.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.operation
	Operation *string `json:"operation,omitempty"`

	// Output only. The type of the maintenance event.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.type
	Type *string `json:"type,omitempty"`

	// Output only. The schedule of the maintenance event.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.schedule
	Schedule *string `json:"schedule,omitempty"`

	// Output only. The state of the maintenance event.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.state
	State *string `json:"state,omitempty"`

	// Output only. The time when the maintenance event request was created.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the maintenance event started.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time when the maintenance event ended, either
	//  successfully or not. If the maintenance event is split into multiple
	//  maintenance windows, end_time is only updated when the whole flow ends.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The time when the maintenance event message was updated.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Fleet
type FleetObservedState struct {
	// Output only. The name of the managed Hub Membership resource associated to
	//  this cluster.
	//
	//  Membership names are formatted as
	//  `projects/<project-number>/locations/global/membership/<cluster-id>`.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Fleet.membership
	Membership *string `json:"membership,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// EdgeContainerClusterObservedState is the state of the EdgeContainerCluster resource as most recently observed in GCP.
// +kcc:proto=google.cloud.edgecontainer.v1.Cluster
type EdgeContainerClusterObservedState struct {
	// Output only. The time when the cluster was created.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the cluster was last updated.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Required. Fleet configuration.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.fleet
	Fleet *FleetObservedState `json:"fleet,omitempty"`

	// Output only. The IP address of the Kubernetes API server.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.endpoint
	Endpoint *string `json:"endpoint,omitempty"`

	// Output only. The port number of the Kubernetes API server.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.port
	Port *int32 `json:"port,omitempty"`

	// Output only. The PEM-encoded public certificate of the cluster's CA.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.cluster_ca_certificate
	ClusterCACertificate *string `json:"clusterCACertificate,omitempty"`

	// Output only. The control plane release version
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.control_plane_version
	ControlPlaneVersion *string `json:"controlPlaneVersion,omitempty"`

	// Output only. The lowest release version among all worker nodes. This field
	//  can be empty if the cluster does not have any worker nodes.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.node_version
	NodeVersion *string `json:"nodeVersion,omitempty"`

	// Optional. Remote control plane disk encryption options. This field is only
	//  used when enabling CMEK support.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.control_plane_encryption
	ControlPlaneEncryption *Cluster_ControlPlaneEncryptionObservedState `json:"controlPlaneEncryption,omitempty"`

	// Output only. The current status of the cluster.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.status
	Status *string `json:"status,omitempty"`

	// Output only. All the maintenance events scheduled for the cluster,
	//  including the ones ongoing, planned for the future and done in the past (up
	//  to 90 days).
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.maintenance_events
	MaintenanceEvents []Cluster_MaintenanceEvent `json:"maintenanceEvents,omitempty"`

	// Output only. The current connection state of the cluster.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.connection_state
	ConnectionState *Cluster_ConnectionState `json:"connectionState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpedgecontainercluster;gcpedgecontainerclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EdgeContainerCluster is the Schema for the EdgeContainerCluster API
// +k8s:openapi-gen=true
type EdgeContainerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   EdgeContainerClusterSpec   `json:"spec,omitempty"`
	Status EdgeContainerClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EdgeContainerClusterList contains a list of EdgeContainerCluster
type EdgeContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeContainerCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EdgeContainerCluster{}, &EdgeContainerClusterList{})
}
