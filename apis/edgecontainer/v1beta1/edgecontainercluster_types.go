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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EdgeContainerClusterGVK = GroupVersion.WithKind("EdgeContainerCluster")

// EdgeContainerClusterSpec defines the desired state of EdgeContainerCluster
// +kcc:spec:proto=google.cloud.edgecontainer.v1.Cluster
type EdgeContainerClusterSpec struct {
	// Immutable. RBAC policy that will be applied and managed by GEC.
	Authorization EdgeContainerClusterAuthorization `json:"authorization"`

	// The configuration of the cluster control plane.
	ControlPlane *EdgeContainerClusterControlPlane `json:"controlPlane,omitempty"`

	// Remote control plane disk encryption options. This field is only used when
	//  enabling CMEK support.
	ControlPlaneEncryption *EdgeContainerClusterControlPlaneEncryption `json:"controlPlaneEncryption,omitempty"`

	// The default maximum number of pods per node used if a maximum
	//  value is not specified explicitly for a node pool in this cluster. If
	//  unspecified, the Kubernetes default value will be used.
	DefaultMaxPodsPerNode *int32 `json:"defaultMaxPodsPerNode,omitempty"`

	// IPv4 address pools for cluster data plane external load
	//  balancing.
	ExternalLoadBalancerIpv4AddressPools []string `json:"externalLoadBalancerIpv4AddressPools,omitempty"`

	// Fleet related configuration.
	//  Fleets are a Google Cloud concept for logically organizing clusters,
	//  letting you use and manage multi-cluster capabilities and apply
	//  consistent policies across your systems.
	Fleet EdgeContainerClusterFleet `json:"fleet"`

	// Immutable. The location of the resource.
	Location string `json:"location"`

	// Cluster-wide maintenance policy configuration.
	MaintenancePolicy *EdgeContainerClusterMaintenancePolicy `json:"maintenancePolicy,omitempty"`

	// Fleet related configuration.
	//  Fleets are a Google Cloud concept for logically organizing clusters,
	//  letting you use and manage multi-cluster capabilities and apply
	//  consistent policies across your systems.
	Networking EdgeContainerClusterNetworking `json:"networking"`

	// The project that this resource belongs to.
	ProjectRef refsv1beta1.ProjectRef `json:"projectRef"`

	// The release channel a cluster is subscribed to.
	ReleaseChannel *string `json:"releaseChannel,omitempty"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Config that customers are allowed to define for GDCE system add-ons.
	SystemAddonsConfig *EdgeContainerClusterSystemAddonsConfig `json:"systemAddonsConfig,omitempty"`

	// The target cluster version. For example: "1.5.0".
	TargetVersion *string `json:"targetVersion,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Authorization
type EdgeContainerClusterAuthorization struct {
	// User that will be granted the cluster-admin role on the cluster, providing
	//  full access to the cluster. Currently, this is a singular field, but will
	//  be expanded to allow multiple admins in the future.
	AdminUsers EdgeContainerClusterAdminUsers `json:"adminUsers"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.ClusterUser
type EdgeContainerClusterAdminUsers struct {
	UsernameRef refsv1beta1.IAMServiceAccountRef `json:"usernameRef"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ControlPlane
type EdgeContainerClusterControlPlane struct {
	// Immutable. Local control plane configuration.
	Local *EdgeContainerClusterControlPlaneLocal `json:"local,omitempty"`

	// Immutable. Remote control plane configuration.
	Remote *EdgeContainerClusterControlPlaneRemote `json:"remote,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Local
type EdgeContainerClusterControlPlaneLocal struct {
	// Only machines matching this filter will be allowed to host control
	//  plane nodes. The filtering language accepts strings like "name=<name>",
	//  and is documented here: [AIP-160](https://google.aip.dev/160).
	MachineFilter *string `json:"machineFilter,omitempty"`

	// The number of nodes to serve as replicas of the Control Plane.
	//  Only 1 and 3 are supported.
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Immutable. Name of the Google Distributed Cloud Edge zones where this node pool
	//  will be created. For example: 'us-central1-edge-customer-a'.
	NodeLocation *string `json:"nodeLocation,omitempty"`

	// Policy configuration about how user applications are deployed. Possible values: ["SHARED_DEPLOYMENT_POLICY_UNSPECIFIED", "ALLOWED", "DISALLOWED"].
	SharedDeploymentPolicy *string `json:"sharedDeploymentPolicy,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ControlPlane.Remote
type EdgeContainerClusterControlPlaneRemote struct {
	// Immutable. Name of the Google Distributed Cloud Edge zones where this node pool
	//  will be created. For example: 'us-central1-edge-customer-a'.
	NodeLocation *string `json:"nodeLocation,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.ControlPlaneEncryption
type EdgeContainerClusterControlPlaneEncryption struct {
	// The Cloud KMS CryptoKeyVersion currently in use for protecting control
	//  plane disks. Only applicable if kms_key is set.
	KmsKeyActiveVersion *string `json:"kmsKeyActiveVersion,omitempty"`

	KmsKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// Availability of the Cloud KMS CryptoKey. If not 'KEY_AVAILABLE', then
	//  nodes may go offline as they cannot access their local data. This can be
	//  caused by a lack of permissions to use the key, or if the key is disabled
	//  or deleted.
	KmsKeyState *string `json:"kmsKeyState,omitempty"`

	// Error status returned by Cloud KMS when using this key. This field may be
	//  populated only if 'kms_key_state' is not 'KMS_KEY_STATE_KEY_AVAILABLE'.
	//  If populated, this field contains the error status reported by Cloud KMS.
	KmsStatus []EdgeContainerClusterKmsStatus `json:"kmsStatus,omitempty"`
}

// +kcc:proto=google.rpc.Status
type EdgeContainerClusterKmsStatus struct {
	// The status code, which should be an enum value of google.rpc.Code.
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Fleet
type EdgeContainerClusterFleet struct {
	// The name of the managed Hub Membership resource associated to this cluster.
	//  Membership names are formatted as 'projects/<project-number>/locations/global/membership/<cluster-id>'.
	Membership *string `json:"membership,omitempty"`

	// The number of the Fleet host project where this cluster will be registered.
	ProjectRef refsv1beta1.ProjectRef `json:"projectRef"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.MaintenancePolicy
type EdgeContainerClusterMaintenancePolicy struct {
	// Specifies the maintenance window in which maintenance may be performed.
	Window EdgeContainerClusterMaintenanceWindow `json:"window"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.MaintenanceWindow
type EdgeContainerClusterMaintenanceWindow struct {
	// Represents an arbitrary window of time that recurs.
	RecurringWindow EdgeContainerClusterRecurringTimeWindow `json:"recurringWindow"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.RecurringTimeWindow
type EdgeContainerClusterRecurringTimeWindow struct {
	// An RRULE (https://tools.ietf.org/html/rfc5545#section-3.8.5.3) for how
	//  this window recurs. They go on for the span of time between the start and
	//  end time.
	Recurrence *string `json:"recurrence,omitempty"`

	// Represents an arbitrary window of time.
	Window *EdgeContainerClusterTimeWindow `json:"window,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.TimeWindow
type EdgeContainerClusterTimeWindow struct {
	// The time that the window ends. The end time must take place after the start time.
	EndTime *string `json:"endTime,omitempty"`

	// The time that the window first starts.
	StartTime *string `json:"startTime,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.ClusterNetworking
type EdgeContainerClusterNetworking struct {
	// Immutable. All pods in the cluster are assigned an RFC1918 IPv4 address from these
	//  blocks. Only a single block is supported. This field cannot be changed
	//  after creation.
	ClusterIpv4CidrBlocks []string `json:"clusterIpv4CidrBlocks"`

	// Immutable. If specified, dual stack mode is enabled and all pods in the cluster are
	//  assigned an IPv6 address from these blocks alongside from an IPv4
	//  address. Only a single block is supported. This field cannot be changed
	//  after creation.
	ClusterIpv6CidrBlocks []string `json:"clusterIpv6CidrBlocks,omitempty"`

	// IP addressing type of this cluster i.e. SINGLESTACK_V4 vs DUALSTACK_V4_V6.
	NetworkType *string `json:"networkType,omitempty"`

	// Immutable. All services in the cluster are assigned an RFC1918 IPv4 address from these
	//  blocks. Only a single block is supported. This field cannot be changed
	//  after creation.
	ServicesIpv4CidrBlocks []string `json:"servicesIpv4CidrBlocks"`

	// Immutable. If specified, dual stack mode is enabled and all services in the cluster are
	//  assigned an IPv6 address from these blocks alongside from an IPv4
	//  address. Only a single block is supported. This field cannot be changed
	//  after creation.
	ServicesIpv6CidrBlocks []string `json:"servicesIpv6CidrBlocks,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig
type EdgeContainerClusterSystemAddonsConfig struct {
	// Config for the Ingress add-on which allows customers to create an Ingress
	//  object to manage external access to the servers in a cluster. The add-on
	//  consists of istiod and istio-ingress.
	Ingress *EdgeContainerClusterSystemAddonsConfigIngress `json:"ingress,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.SystemAddonsConfig.Ingress
type EdgeContainerClusterSystemAddonsConfigIngress struct {
	// Whether Ingress is disabled.
	Disabled *bool `json:"disabled,omitempty"`

	// Ingress VIP.
	Ipv4Vip *string `json:"ipv4Vip,omitempty"`
}

// EdgeContainerClusterStatus defines the config connector machine state of EdgeContainerCluster
type EdgeContainerClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// The PEM-encoded public certificate of the cluster's CA.
	ClusterCaCertificate *string `json:"clusterCaCertificate,omitempty"`

	// The control plane release version.
	ControlPlaneVersion *string `json:"controlPlaneVersion,omitempty"`

	// The time the cluster was created, in RFC3339 text format.
	CreateTime *string `json:"createTime,omitempty"`

	// The IP address of the Kubernetes API server.
	Endpoint *string `json:"endpoint,omitempty"`

	// All the maintenance events scheduled for the cluster, including the ones
	//  ongoing, planned for the future and done in the past (up to 90 days).
	MaintenanceEvents []EdgeContainerClusterMaintenanceEvent `json:"maintenanceEvents,omitempty"`

	// The lowest release version among all worker nodes. This field can be empty if the cluster does not have any worker nodes.
	NodeVersion *string `json:"nodeVersion,omitempty"`

	// The port number of the Kubernetes API server.
	Port *int32 `json:"port,omitempty"`

	// Indicates the status of the cluster.
	Status *string `json:"status,omitempty"`

	// The time the cluster was last updated, in RFC3339 text format.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.Cluster.MaintenanceEvent
type EdgeContainerClusterMaintenanceEvent struct {
	// The time when the maintenance event request was created.
	CreateTime *string `json:"createTime,omitempty"`

	// The time when the maintenance event ended, either successfully or not. If
	//  the maintenance event is split into multiple maintenance windows,
	//  end_time is only updated when the whole flow ends.
	EndTime *string `json:"endTime,omitempty"`

	// The operation for running the maintenance event. Specified in the format
	//  projects/*/locations/*/operations/*. If the maintenance event is split
	//  into multiple operations (e.g. due to maintenance windows), the latest
	//  one is recorded.
	Operation *string `json:"operation,omitempty"`

	// The schedule of the maintenance event.
	Schedule *string `json:"schedule,omitempty"`

	// The time when the maintenance event started.
	StartTime *string `json:"startTime,omitempty"`

	// Indicates the maintenance event state.
	State *string `json:"state,omitempty"`

	// The target version of the cluster.
	TargetVersion *string `json:"targetVersion,omitempty"`

	// Indicates the maintenance event type.
	Type *string `json:"type,omitempty"`

	// The time when the maintenance event message was updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// UUID of the maintenance event.
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
