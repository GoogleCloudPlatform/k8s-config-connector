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


// +kcc:proto=google.cloud.gkehub.v1beta1.ApplianceCluster
type ApplianceCluster struct {
	// Immutable. Self-link of the GCP resource for the Appliance Cluster. For
	//  example:
	//
	//  //transferappliance.googleapis.com/projects/my-project/locations/us-west1-a/appliances/my-appliance
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.ApplianceCluster.resource_link
	ResourceLink *string `json:"resourceLink,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.Authority
type Authority struct {
	// Optional. A JSON Web Token (JWT) issuer URI. `issuer` must start with
	//  `https://` and be a valid URL with length <2000 characters.
	//
	//  If set, then Google will allow valid OIDC tokens from this issuer to
	//  authenticate within the workload_identity_pool. OIDC discovery will be
	//  performed on this URI to validate tokens from the issuer.
	//
	//  Clearing `issuer` disables Workload Identity. `issuer` cannot be directly
	//  modified; it must be cleared (and Workload Identity disabled) before using
	//  a new issuer (and re-enabling Workload Identity).
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Authority.issuer
	Issuer *string `json:"issuer,omitempty"`

	// Optional. OIDC verification keys for this Membership in JWKS format (RFC
	//  7517).
	//
	//  When this field is set, OIDC discovery will NOT be performed on `issuer`,
	//  and instead OIDC tokens will be validated using this field.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Authority.oidc_jwks
	OidcJwks []byte `json:"oidcJwks,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.EdgeCluster
type EdgeCluster struct {
	// Immutable. Self-link of the GCP resource for the Edge Cluster. For
	//  example:
	//
	//  //edgecontainer.googleapis.com/projects/my-project/locations/us-west1-a/clusters/my-cluster
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.EdgeCluster.resource_link
	ResourceLink *string `json:"resourceLink,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.GkeCluster
type GkeCluster struct {
	// Immutable. Self-link of the GCP resource for the GKE cluster. For example:
	//
	//      //container.googleapis.com/projects/my-project/locations/us-west1-a/clusters/my-cluster
	//
	//  Zonal clusters are also supported.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.GkeCluster.resource_link
	ResourceLink *string `json:"resourceLink,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.KubernetesMetadata
type KubernetesMetadata struct {
}

// +kcc:proto=google.cloud.gkehub.v1beta1.KubernetesResource
type KubernetesResource struct {
	// Input only. The YAML representation of the Membership CR. This field is
	//  ignored for GKE clusters where Hub can read the CR directly.
	//
	//  Callers should provide the CR that is currently present in the cluster
	//  during CreateMembership or UpdateMembership, or leave this field empty if
	//  none exists. The CR manifest is used to validate the cluster has not been
	//  registered with another Membership.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.KubernetesResource.membership_cr_manifest
	MembershipCrManifest *string `json:"membershipCrManifest,omitempty"`

	// Optional. Options for Kubernetes resource generation.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.KubernetesResource.resource_options
	ResourceOptions *ResourceOptions `json:"resourceOptions,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.Membership
type Membership struct {

	// Optional. GCP labels for this membership.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of this membership, limited to 63 characters.
	//  Must match the regex: `[a-zA-Z0-9][a-zA-Z0-9_\-\.\ ]*`
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.description
	Description *string `json:"description,omitempty"`

	// Optional. Endpoint information to reach this member.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.endpoint
	Endpoint *MembershipEndpoint `json:"endpoint,omitempty"`

	// Optional. How to identify workloads from this Membership.
	//  See the documentation on Workload Identity for more details:
	//  https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.authority
	Authority *Authority `json:"authority,omitempty"`

	// Optional. An externally-generated and managed ID for this Membership. This
	//  ID may be modified after creation, but this is not recommended. For GKE
	//  clusters, external_id is managed by the Hub API and updates will be
	//  ignored.
	//
	//  The ID must match the regex: `[a-zA-Z0-9][a-zA-Z0-9_\-\.]*`
	//
	//  If this Membership represents a Kubernetes cluster, this value should be
	//  set to the UID of the `kube-system` namespace object.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.external_id
	ExternalID *string `json:"externalID,omitempty"`

	// Optional. The infrastructure type this Membership is running on.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.infrastructure_type
	InfrastructureType *string `json:"infrastructureType,omitempty"`

	// Optional. The monitoring config information for this membership.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.monitoring_config
	MonitoringConfig *MonitoringConfig `json:"monitoringConfig,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.MembershipEndpoint
type MembershipEndpoint struct {
	// Optional. Specific information for a GKE-on-GCP cluster.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipEndpoint.gke_cluster
	GkeCluster *GkeCluster `json:"gkeCluster,omitempty"`

	// Optional. Specific information for a GKE On-Prem cluster. An onprem
	//  user-cluster who has no resourceLink is not allowed to use this field, it
	//  should have a nil "type" instead.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipEndpoint.on_prem_cluster
	OnPremCluster *OnPremCluster `json:"onPremCluster,omitempty"`

	// Optional. Specific information for a GKE Multi-Cloud cluster.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipEndpoint.multi_cloud_cluster
	MultiCloudCluster *MultiCloudCluster `json:"multiCloudCluster,omitempty"`

	// Optional. Specific information for a Google Edge cluster.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipEndpoint.edge_cluster
	EdgeCluster *EdgeCluster `json:"edgeCluster,omitempty"`

	// Optional. Specific information for a GDC Edge Appliance cluster.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipEndpoint.appliance_cluster
	ApplianceCluster *ApplianceCluster `json:"applianceCluster,omitempty"`

	// Optional. The in-cluster Kubernetes Resources that should be applied for a
	//  correctly registered cluster, in the steady state. These resources:
	//
	//    * Ensure that the cluster is exclusively registered to one and only one
	//      Hub Membership.
	//    * Propagate Workload Pool Information available in the Membership
	//      Authority field.
	//    * Ensure proper initial configuration of default Hub Features.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipEndpoint.kubernetes_resource
	KubernetesResource *KubernetesResource `json:"kubernetesResource,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.MembershipState
type MembershipState struct {

	// This field is never set by the Hub Service.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipState.description
	Description *string `json:"description,omitempty"`

	// This field is never set by the Hub Service.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipState.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.MonitoringConfig
type MonitoringConfig struct {
	// Immutable. Project used to report Metrics
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MonitoringConfig.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Immutable. Location used to report Metrics
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MonitoringConfig.location
	Location *string `json:"location,omitempty"`

	// Immutable. Cluster name used to report metrics.
	//  For Anthos on VMWare/Baremetal, it would be in format
	//  `memberClusters/cluster_name`; And for Anthos on MultiCloud, it would be in
	//  format
	//  `{azureClusters, awsClusters}/cluster_name`.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MonitoringConfig.cluster
	Cluster *string `json:"cluster,omitempty"`

	// Kubernetes system metrics, if available, are written to this prefix.
	//  This defaults to kubernetes.io for GKE, and kubernetes.io/anthos for Anthos
	//  eventually. Noted: Anthos MultiCloud will have kubernetes.io prefix today
	//  but will migration to be under kubernetes.io/anthos
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MonitoringConfig.kubernetes_metrics_prefix
	KubernetesMetricsPrefix *string `json:"kubernetesMetricsPrefix,omitempty"`

	// Immutable. Cluster hash, this is a unique string generated by google code,
	//  which does not contain any PII, which we can use to reference the cluster.
	//  This is expected to be created by the monitoring stack and persisted into
	//  the Cluster object as well as to GKE-Hub.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MonitoringConfig.cluster_hash
	ClusterHash *string `json:"clusterHash,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.MultiCloudCluster
type MultiCloudCluster struct {
	// Immutable. Self-link of the GCP resource for the GKE Multi-Cloud cluster.
	//  For example:
	//
	//   //gkemulticloud.googleapis.com/projects/my-project/locations/us-west1-a/awsClusters/my-cluster
	//   //gkemulticloud.googleapis.com/projects/my-project/locations/us-west1-a/azureClusters/my-cluster
	//   //gkemulticloud.googleapis.com/projects/my-project/locations/us-west1-a/attachedClusters/my-cluster
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MultiCloudCluster.resource_link
	ResourceLink *string `json:"resourceLink,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.OnPremCluster
type OnPremCluster struct {
	// Immutable. Self-link of the GCP resource for the GKE On-Prem cluster. For
	//  example:
	//
	//   //gkeonprem.googleapis.com/projects/my-project/locations/us-west1-a/vmwareClusters/my-cluster
	//   //gkeonprem.googleapis.com/projects/my-project/locations/us-west1-a/bareMetalClusters/my-cluster
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.OnPremCluster.resource_link
	ResourceLink *string `json:"resourceLink,omitempty"`

	// Immutable. Whether the cluster is an admin cluster.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.OnPremCluster.admin_cluster
	AdminCluster *bool `json:"adminCluster,omitempty"`

	// Immutable. The on prem cluster's type.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.OnPremCluster.cluster_type
	ClusterType *string `json:"clusterType,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.ResourceManifest
type ResourceManifest struct {
	// YAML manifest of the resource.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.ResourceManifest.manifest
	Manifest *string `json:"manifest,omitempty"`

	// Whether the resource provided in the manifest is `cluster_scoped`.
	//  If unset, the manifest is assumed to be namespace scoped.
	//
	//  This field is used for REST mapping when applying the resource in a
	//  cluster.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.ResourceManifest.cluster_scoped
	ClusterScoped *bool `json:"clusterScoped,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.ResourceOptions
type ResourceOptions struct {
	// Optional. The Connect agent version to use for connect_resources. Defaults
	//  to the latest GKE Connect version. The version must be a currently
	//  supported version, obsolete versions will be rejected.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.ResourceOptions.connect_version
	ConnectVersion *string `json:"connectVersion,omitempty"`

	// Optional. Use `apiextensions/v1beta1` instead of `apiextensions/v1` for
	//  CustomResourceDefinition resources.
	//  This option should be set for clusters with Kubernetes apiserver versions
	//  <1.16.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.ResourceOptions.v1beta1_crd
	V1beta1Crd *bool `json:"v1beta1Crd,omitempty"`

	// Optional. Major version of the Kubernetes cluster. This is only used to
	//  determine which version to use for the CustomResourceDefinition resources,
	//  `apiextensions/v1beta1` or`apiextensions/v1`.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.ResourceOptions.k8s_version
	K8sVersion *string `json:"k8sVersion,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.Authority
type AuthorityObservedState struct {
	// Output only. The name of the workload identity pool in which `issuer` will
	//  be recognized.
	//
	//  There is a single Workload Identity Pool per Hub that is shared
	//  between all Memberships that belong to that Hub. For a Hub hosted in
	//  {PROJECT_ID}, the workload pool format is `{PROJECT_ID}.hub.id.goog`,
	//  although this is subject to change in newer versions of this API.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Authority.workload_identity_pool
	WorkloadIdentityPool *string `json:"workloadIdentityPool,omitempty"`

	// Output only. An identity provider that reflects the `issuer` in the
	//  workload identity pool.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Authority.identity_provider
	IdentityProvider *string `json:"identityProvider,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.GkeCluster
type GkeClusterObservedState struct {
	// Output only. If cluster_missing is set then it denotes that the GKE cluster
	//  no longer exists in the GKE Control Plane.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.GkeCluster.cluster_missing
	ClusterMissing *bool `json:"clusterMissing,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.KubernetesMetadata
type KubernetesMetadataObservedState struct {
	// Output only. Kubernetes API server version string as reported by
	//  '/version'.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.KubernetesMetadata.kubernetes_api_server_version
	KubernetesApiServerVersion *string `json:"kubernetesApiServerVersion,omitempty"`

	// Output only. Node providerID as reported by the first node in the list of
	//  nodes on the Kubernetes endpoint. On Kubernetes platforms that support
	//  zero-node clusters (like GKE-on-GCP), the node_count will be zero and the
	//  node_provider_id will be empty.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.KubernetesMetadata.node_provider_id
	NodeProviderID *string `json:"nodeProviderID,omitempty"`

	// Output only. Node count as reported by Kubernetes nodes resources.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.KubernetesMetadata.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Output only. vCPU count as reported by Kubernetes nodes resources.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.KubernetesMetadata.vcpu_count
	VcpuCount *int32 `json:"vcpuCount,omitempty"`

	// Output only. The total memory capacity as reported by the sum of all
	//  Kubernetes nodes resources, defined in MB.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.KubernetesMetadata.memory_mb
	MemoryMb *int32 `json:"memoryMb,omitempty"`

	// Output only. The time at which these details were last updated. This
	//  update_time is different from the Membership-level update_time since
	//  EndpointDetails are updated internally for API consumers.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.KubernetesMetadata.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.KubernetesResource
type KubernetesResourceObservedState struct {
	// Output only. Additional Kubernetes resources that need to be applied to the
	//  cluster after Membership creation, and after every update.
	//
	//  This field is only populated in the Membership returned from a successful
	//  long-running operation from CreateMembership or UpdateMembership. It is not
	//  populated during normal GetMembership or ListMemberships requests. To get
	//  the resource manifest after the initial registration, the caller should
	//  make a UpdateMembership call with an empty field mask.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.KubernetesResource.membership_resources
	MembershipResources []ResourceManifest `json:"membershipResources,omitempty"`

	// Output only. The Kubernetes resources for installing the GKE Connect agent
	//
	//  This field is only populated in the Membership returned from a successful
	//  long-running operation from CreateMembership or UpdateMembership. It is not
	//  populated during normal GetMembership or ListMemberships requests. To get
	//  the resource manifest after the initial registration, the caller should
	//  make a UpdateMembership call with an empty field mask.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.KubernetesResource.connect_resources
	ConnectResources []ResourceManifest `json:"connectResources,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.Membership
type MembershipObservedState struct {
	// Output only. The full, unique name of this Membership resource in the
	//  format `projects/*/locations/*/memberships/{membership_id}`, set during
	//  creation.
	//
	//  `membership_id` must be a valid RFC 1123 compliant DNS label:
	//
	//    1. At most 63 characters in length
	//    2. It must consist of lower case alphanumeric characters or `-`
	//    3. It must start and end with an alphanumeric character
	//
	//  Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`,
	//  with a maximum length of 63 characters.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.name
	Name *string `json:"name,omitempty"`

	// Optional. Endpoint information to reach this member.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.endpoint
	Endpoint *MembershipEndpointObservedState `json:"endpoint,omitempty"`

	// Output only. State of the Membership resource.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.state
	State *MembershipState `json:"state,omitempty"`

	// Optional. How to identify workloads from this Membership.
	//  See the documentation on Workload Identity for more details:
	//  https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.authority
	Authority *AuthorityObservedState `json:"authority,omitempty"`

	// Output only. When the Membership was created.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. When the Membership was last updated.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. When the Membership was deleted.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. For clusters using Connect, the timestamp of the most recent
	//  connection established with Google Cloud. This time is updated every
	//  several minutes, not continuously. For clusters that do not use GKE
	//  Connect, or that have never connected successfully, this field will be
	//  unset.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.last_connection_time
	LastConnectionTime *string `json:"lastConnectionTime,omitempty"`

	// Output only. Google-generated UUID for this resource. This is unique across
	//  all Membership resources. If a Membership resource is deleted and another
	//  resource with the same name is created, it gets a different unique_id.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.Membership.unique_id
	UniqueID *string `json:"uniqueID,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.MembershipEndpoint
type MembershipEndpointObservedState struct {
	// Optional. Specific information for a GKE-on-GCP cluster.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipEndpoint.gke_cluster
	GkeCluster *GkeClusterObservedState `json:"gkeCluster,omitempty"`

	// Optional. Specific information for a GKE On-Prem cluster. An onprem
	//  user-cluster who has no resourceLink is not allowed to use this field, it
	//  should have a nil "type" instead.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipEndpoint.on_prem_cluster
	OnPremCluster *OnPremClusterObservedState `json:"onPremCluster,omitempty"`

	// Optional. Specific information for a GKE Multi-Cloud cluster.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipEndpoint.multi_cloud_cluster
	MultiCloudCluster *MultiCloudClusterObservedState `json:"multiCloudCluster,omitempty"`

	// Output only. Useful Kubernetes-specific metadata.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipEndpoint.kubernetes_metadata
	KubernetesMetadata *KubernetesMetadata `json:"kubernetesMetadata,omitempty"`

	// Optional. The in-cluster Kubernetes Resources that should be applied for a
	//  correctly registered cluster, in the steady state. These resources:
	//
	//    * Ensure that the cluster is exclusively registered to one and only one
	//      Hub Membership.
	//    * Propagate Workload Pool Information available in the Membership
	//      Authority field.
	//    * Ensure proper initial configuration of default Hub Features.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipEndpoint.kubernetes_resource
	KubernetesResource *KubernetesResourceObservedState `json:"kubernetesResource,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.MembershipState
type MembershipStateObservedState struct {
	// Output only. The current state of the Membership resource.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MembershipState.code
	Code *string `json:"code,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.MultiCloudCluster
type MultiCloudClusterObservedState struct {
	// Output only. If cluster_missing is set then it denotes that
	//  API(gkemulticloud.googleapis.com) resource for this GKE Multi-Cloud cluster
	//  no longer exists.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.MultiCloudCluster.cluster_missing
	ClusterMissing *bool `json:"clusterMissing,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta1.OnPremCluster
type OnPremClusterObservedState struct {
	// Output only. If cluster_missing is set then it denotes that
	//  API(gkeonprem.googleapis.com) resource for this GKE On-Prem cluster no
	//  longer exists.
	// +kcc:proto:field=google.cloud.gkehub.v1beta1.OnPremCluster.cluster_missing
	ClusterMissing *bool `json:"clusterMissing,omitempty"`
}
