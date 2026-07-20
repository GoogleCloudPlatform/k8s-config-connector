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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EdgeContainerClusterGVK = GroupVersion.WithKind("EdgeContainerCluster")

// EdgeContainerClusterSpec defines the desired state of EdgeContainerCluster
// +kcc:spec:proto=google.cloud.edgecontainer.v1.Cluster
type EdgeContainerClusterSpec struct {
	// Immutable. RBAC policy that will be applied and managed by GEC.
	// +required
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.authorization
	Authorization *Authorization `json:"authorization"`

	// Optional. The configuration of the cluster control plane.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.control_plane
	ControlPlane *Cluster_ControlPlane `json:"controlPlane,omitempty"`

	// Optional. Remote control plane disk encryption options. This field is only used when enabling CMEK support.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.control_plane_encryption
	ControlPlaneEncryption *Cluster_ControlPlaneEncryption `json:"controlPlaneEncryption,omitempty"`

	// Optional. The default maximum number of pods per node used if a maximum value is not specified explicitly for a node pool in this cluster.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.default_max_pods_per_node
	DefaultMaxPodsPerNode *int32 `json:"defaultMaxPodsPerNode,omitempty"`

	// Optional. IPv4 address pools for cluster data plane external load balancing.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.external_load_balancer_ipv4_address_pools
	ExternalLoadBalancerIpv4AddressPools []string `json:"externalLoadBalancerIpv4AddressPools,omitempty"`

	// Immutable. Fleet related configuration.
	// +required
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.fleet
	Fleet *Fleet `json:"fleet"`

	// Immutable. The location of the resource.
	// +required
	Location string `json:"location"`

	// Cluster-wide maintenance policy configuration.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.maintenance_policy
	MaintenancePolicy *MaintenancePolicy `json:"maintenancePolicy,omitempty"`

	// Fleet related configuration.
	// +required
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.networking
	Networking *ClusterNetworking `json:"networking"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *parent.ProjectRef `json:"projectRef"`

	// Optional. The release channel a cluster is subscribed to.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.release_channel
	ReleaseChannel *string `json:"releaseChannel,omitempty"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Config that customers are allowed to define for GDCE system add-ons.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.system_addons_config
	SystemAddonsConfig *Cluster_SystemAddonsConfig `json:"systemAddonsConfig,omitempty"`

	// Optional. The target cluster version. For example: "1.5.0".
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.target_version
	TargetVersion *string `json:"targetVersion,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Authorization
type Authorization struct {
	// User that will be granted the cluster-admin role on the cluster, providing full access to the cluster.
	// +required
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Authorization.admin_users
	AdminUsers *ClusterUser `json:"adminUsers"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.ClusterUser
type ClusterUser struct {
	// +required
	// +kcc:proto:field=google.cloud.edgecontainer.v1.ClusterUser.username
	UsernameRef *refsv1beta1.IAMServiceAccountRef `json:"usernameRef"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ControlPlane
type Cluster_ControlPlane struct {
	// Immutable. Local control plane configuration.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.local
	Local *Cluster_ControlPlane_Local `json:"local,omitempty"`

	// Immutable. Remote control plane configuration.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.remote
	Remote *Cluster_ControlPlane_Remote `json:"remote,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Local
type Cluster_ControlPlane_Local struct {
	// Only machines matching this filter will be allowed to host control plane nodes.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Local.machine_filter
	MachineFilter *string `json:"machineFilter,omitempty"`

	// The number of nodes to serve as replicas of the Control Plane. Only 1 and 3 are supported.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Local.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Immutable. Name of the Google Distributed Cloud Edge zones where this node pool will be created.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Local.node_location
	NodeLocation *string `json:"nodeLocation,omitempty"`

	// Policy configuration about how user applications are deployed.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Local.shared_deployment_policy
	SharedDeploymentPolicy *string `json:"sharedDeploymentPolicy,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Remote
type Cluster_ControlPlane_Remote struct {
	// Immutable. Name of the Google Distributed Cloud Edge zones where this node pool will be created.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Remote.node_location
	NodeLocation *string `json:"nodeLocation,omitempty"`
}

type Cluster_ControlPlaneEncryption struct {
	// Output only. The Cloud KMS CryptoKeyVersion currently in use for protecting control plane disks. Only applicable if kms_key is set.
	KMSKeyActiveVersion *string `json:"kmsKeyActiveVersion,omitempty"`

	// The Cloud KMS CryptoKey to use for protecting control plane disks.
	KMSKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// Output only. Availability of the Cloud KMS CryptoKey.
	KMSKeyState *string `json:"kmsKeyState,omitempty"`

	// Output only. Error status returned by Cloud KMS when using this key.
	KMSStatus []KMSStatus `json:"kmsStatus,omitempty"`
}

type KMSStatus struct {
	// The status code, which should be an enum value of google.rpc.Code.
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English.
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Fleet
type Fleet struct {
	// The name of the managed Hub Membership resource associated to this cluster.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Fleet.membership
	Membership *string `json:"membership,omitempty"`

	// Required. The number of the Fleet host project where this cluster will be registered.
	// +required
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Fleet.project
	ProjectRef *parent.ProjectRef `json:"projectRef"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.MaintenancePolicy
type MaintenancePolicy struct {
	// Specifies the maintenance window in which maintenance may be performed.
	// +required
	// +kcc:proto:field=google.cloud.edgecontainer.v1.MaintenancePolicy.window
	Window *MaintenanceWindow `json:"window"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.MaintenanceWindow
type MaintenanceWindow struct {
	// Represents an arbitrary window of time that recurs.
	// +required
	// +kcc:proto:field=google.cloud.edgecontainer.v1.MaintenanceWindow.recurring_window
	RecurringWindow *RecurringTimeWindow `json:"recurringWindow"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.RecurringTimeWindow
type RecurringTimeWindow struct {
	// An RRULE for how this window recurs.
	Recurrence *string `json:"recurrence,omitempty"`

	// Represents an arbitrary window of time.
	Window *TimeWindow `json:"window,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.TimeWindow
type TimeWindow struct {
	// The time that the window ends.
	EndTime *string `json:"endTime,omitempty"`

	// The time that the window first starts.
	StartTime *string `json:"startTime,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.ClusterNetworking
type ClusterNetworking struct {
	// Immutable. All pods in the cluster are assigned an RFC1918 IPv4 address from these blocks. Only a single block is supported. This field cannot be changed after creation.
	// +required
	// +kcc:proto:field=google.cloud.edgecontainer.v1.ClusterNetworking.cluster_ipv4_cidr_blocks
	ClusterIpv4CidrBlocks []string `json:"clusterIpv4CidrBlocks"`

	// Immutable. If specified, dual stack mode is enabled and all pods in the cluster are assigned an IPv6 address from these blocks alongside from an IPv4 address. Only a single block is supported. This field cannot be changed after creation.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.ClusterNetworking.cluster_ipv6_cidr_blocks
	ClusterIpv6CidrBlocks []string `json:"clusterIpv6CidrBlocks,omitempty"`

	// IP addressing type of this cluster.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.ClusterNetworking.network_type
	NetworkType *string `json:"networkType,omitempty"`

	// Immutable. All services in the cluster are assigned an RFC1918 IPv4 address from these blocks. Only a single block is supported. This field cannot be changed after creation.
	// +required
	// +kcc:proto:field=google.cloud.edgecontainer.v1.ClusterNetworking.services_ipv4_cidr_blocks
	ServicesIpv4CidrBlocks []string `json:"servicesIpv4CidrBlocks"`

	// Immutable. If specified, dual stack mode is enabled and all services in the cluster are assigned an RFC1918 IPv4 address from these blocks. Only a single block is supported. This field cannot be changed after creation.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.ClusterNetworking.services_ipv6_cidr_blocks
	ServicesIpv6CidrBlocks []string `json:"servicesIpv6CidrBlocks,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig
type Cluster_SystemAddonsConfig struct {
	// Config for the Ingress add-on.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig.ingress
	Ingress *Cluster_SystemAddonsConfig_Ingress `json:"ingress,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig.Ingress
type Cluster_SystemAddonsConfig_Ingress struct {
	// Whether Ingress is disabled.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig.Ingress.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Ingress VIP.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig.Ingress.ipv4_vip
	Ipv4Vip *string `json:"ipv4Vip,omitempty"`
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

	// Output only. The PEM-encoded public certificate of the cluster's CA.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.cluster_ca_certificate
	ClusterCaCertificate *string `json:"clusterCaCertificate,omitempty"`

	// Output only. The control plane release version.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.control_plane_version
	ControlPlaneVersion *string `json:"controlPlaneVersion,omitempty"`

	// Output only. The time the cluster was created, in RFC3339 text format.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The IP address of the Kubernetes API server.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.endpoint
	Endpoint *string `json:"endpoint,omitempty"`

	// Output only. All the maintenance events scheduled for the cluster.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.maintenance_events
	MaintenanceEvents []Cluster_MaintenanceEventObservedState `json:"maintenanceEvents,omitempty"`

	// Output only. The lowest release version among all worker nodes. This field can be empty if the cluster does not have any worker nodes.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.node_version
	NodeVersion *string `json:"nodeVersion,omitempty"`

	// Output only. The port number of the Kubernetes API server.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.port
	Port *int32 `json:"port,omitempty"`

	// Output only. Indicates the status of the cluster.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.status
	Status *string `json:"status,omitempty"`

	// Output only. The time the cluster was last updated, in RFC3339 text format.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

type Cluster_MaintenanceEventObservedState struct {
	// Output only. The time when the maintenance event request was created.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the maintenance event ended, either successfully or not.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The operation for running the maintenance event.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.operation
	Operation *string `json:"operation,omitempty"`

	// Output only. The schedule of the maintenance event.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.schedule
	Schedule *string `json:"schedule,omitempty"`

	// Output only. The time when the maintenance event started.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Indicates the maintenance event state.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.state
	State *string `json:"state,omitempty"`

	// Output only. The target version of the cluster.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.target_version
	TargetVersion *string `json:"targetVersion,omitempty"`

	// Output only. Indicates the maintenance event type.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.type
	Type *string `json:"type,omitempty"`

	// Output only. The time when the maintenance event message was updated.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. UUID of the maintenance event.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent.uuid
	Uuid *string `json:"uuid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpedgecontainercluster;gcpedgecontainerclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
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
