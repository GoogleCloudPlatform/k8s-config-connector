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


// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedCluster
type AttachedCluster struct {
	// The name of this resource.
	//
	//  Cluster names are formatted as
	//  `projects/<project-number>/locations/<region>/attachedClusters/<cluster-id>`.
	//
	//  See [Resource Names](https://cloud.google.com/apis/design/resource_names)
	//  for more details on Google Cloud Platform resource names.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.name
	Name *string `json:"name,omitempty"`

	// Optional. A human readable description of this cluster.
	//  Cannot be longer than 255 UTF-8 encoded bytes.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.description
	Description *string `json:"description,omitempty"`

	// Required. OpenID Connect (OIDC) configuration for the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.oidc_config
	OIDCConfig *AttachedOIDCConfig `json:"oidcConfig,omitempty"`

	// Required. The platform version for the cluster (e.g. `1.19.0-gke.1000`).
	//
	//  You can list all supported versions on a given Google Cloud region by
	//  calling
	//  [GetAttachedServerConfig][google.cloud.gkemulticloud.v1.AttachedClusters.GetAttachedServerConfig].
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.platform_version
	PlatformVersion *string `json:"platformVersion,omitempty"`

	// Required. The Kubernetes distribution of the underlying attached cluster.
	//
	//  Supported values: ["eks", "aks", "generic"].
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.distribution
	Distribution *string `json:"distribution,omitempty"`

	// Required. Fleet configuration.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.fleet
	Fleet *Fleet `json:"fleet,omitempty"`

	// Allows clients to perform consistent read-modify-writes
	//  through optimistic concurrency control.
	//
	//  Can be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Annotations on the cluster.
	//
	//  This field has the same restrictions as Kubernetes annotations.
	//  The total size of all keys and values combined is limited to 256k.
	//  Key can have 2 segments: prefix (optional) and name (required),
	//  separated by a slash (/).
	//  Prefix must be a DNS subdomain.
	//  Name must be 63 characters or less, begin and end with alphanumerics,
	//  with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Logging configuration for this cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// Optional. Configuration related to the cluster RBAC settings.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.authorization
	Authorization *AttachedClustersAuthorization `json:"authorization,omitempty"`

	// Optional. Monitoring configuration for this cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.monitoring_config
	MonitoringConfig *MonitoringConfig `json:"monitoringConfig,omitempty"`

	// Optional. Proxy configuration for outbound HTTP(S) traffic.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.proxy_config
	ProxyConfig *AttachedProxyConfig `json:"proxyConfig,omitempty"`

	// Optional. Binary Authorization configuration for this cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.binary_authorization
	BinaryAuthorization *BinaryAuthorization `json:"binaryAuthorization,omitempty"`

	// Optional. Security Posture configuration for this cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.security_posture_config
	SecurityPostureConfig *SecurityPostureConfig `json:"securityPostureConfig,omitempty"`

	// Optional. Input only. Tag keys/values directly bound to this resource.
	//
	//  Tag key must be specified in the format <tag namespace>/<tag key name>
	//  where the tag namespace is the ID of the organization or name of the
	//  project that the tag key is defined in.
	//  The short name of a tag key or value can have a maximum length of 256
	//  characters. The permitted character set for the short name includes UTF-8
	//  encoded Unicode characters except single quotes ('), double quotes ("),
	//  backslashes (\), and forward slashes (/).
	//
	//  See
	//  [Tags](https://cloud.google.com/resource-manager/docs/tags/tags-overview)
	//  for more details on Google Cloud Platform tags.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.tags
	Tags map[string]string `json:"tags,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedClusterError
type AttachedClusterError struct {
	// Human-friendly description of the error.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedClusterError.message
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedClusterGroup
type AttachedClusterGroup struct {
	// Required. The name of the group, e.g. `my-group@domain.com`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedClusterGroup.group
	Group *string `json:"group,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedClusterUser
type AttachedClusterUser struct {
	// Required. The name of the user, e.g. `my-gcp-id@gmail.com`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedClusterUser.username
	Username *string `json:"username,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedClustersAuthorization
type AttachedClustersAuthorization struct {
	// Optional. Users that can perform operations as a cluster admin. A managed
	//  ClusterRoleBinding will be created to grant the `cluster-admin` ClusterRole
	//  to the users. Up to ten admin users can be provided.
	//
	//  For more info on RBAC, see
	//  https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedClustersAuthorization.admin_users
	AdminUsers []AttachedClusterUser `json:"adminUsers,omitempty"`

	// Optional. Groups of users that can perform operations as a cluster admin. A
	//  managed ClusterRoleBinding will be created to grant the `cluster-admin`
	//  ClusterRole to the groups. Up to ten admin groups can be provided.
	//
	//  For more info on RBAC, see
	//  https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedClustersAuthorization.admin_groups
	AdminGroups []AttachedClusterGroup `json:"adminGroups,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedOidcConfig
type AttachedOIDCConfig struct {
	// A JSON Web Token (JWT) issuer URI. `issuer` must start with `https://`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedOidcConfig.issuer_url
	IssuerURL *string `json:"issuerURL,omitempty"`

	// Optional. OIDC verification keys in JWKS format (RFC 7517).
	//  It contains a list of OIDC verification keys that can be used to verify
	//  OIDC JWTs.
	//
	//  This field is required for cluster that doesn't have a publicly available
	//  discovery endpoint. When provided, it will be directly used
	//  to verify the OIDC JWT asserted by the IDP.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedOidcConfig.jwks
	Jwks []byte `json:"jwks,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedProxyConfig
type AttachedProxyConfig struct {
	// The Kubernetes Secret resource that contains the HTTP(S) proxy
	//  configuration. The secret must be a JSON encoded proxy configuration
	//  as described in
	//  https://cloud.google.com/kubernetes-engine/multi-cloud/docs/attached/eks/how-to/use-a-proxy#configure-proxy-support
	//  for EKS clusters and
	//  https://cloud.google.com/kubernetes-engine/multi-cloud/docs/attached/aks/how-to/use-a-proxy#configure-proxy-support
	//  for AKS clusters.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedProxyConfig.kubernetes_secret
	KubernetesSecret *KubernetesSecret `json:"kubernetesSecret,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.BinaryAuthorization
type BinaryAuthorization struct {
	// Mode of operation for binauthz policy evaluation. If unspecified, defaults
	//  to DISABLED.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.BinaryAuthorization.evaluation_mode
	EvaluationMode *string `json:"evaluationMode,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.CloudMonitoringConfig
type CloudMonitoringConfig struct {
	// Enable GKE-native logging and metrics.
	//  Only for Attached Clusters.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.CloudMonitoringConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.Fleet
type Fleet struct {
	// Required. The name of the Fleet host project where this cluster will be
	//  registered.
	//
	//  Project names are formatted as
	//  `projects/<project-number>`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.Fleet.project
	Project *string `json:"project,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.KubernetesSecret
type KubernetesSecret struct {
	// Name of the kubernetes secret.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.KubernetesSecret.name
	Name *string `json:"name,omitempty"`

	// Namespace in which the kubernetes secret is stored.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.KubernetesSecret.namespace
	Namespace *string `json:"namespace,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.LoggingComponentConfig
type LoggingComponentConfig struct {
	// The components to be enabled.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.LoggingComponentConfig.enable_components
	EnableComponents []string `json:"enableComponents,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.LoggingConfig
type LoggingConfig struct {
	// The configuration of the logging components;
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.LoggingConfig.component_config
	ComponentConfig *LoggingComponentConfig `json:"componentConfig,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.ManagedPrometheusConfig
type ManagedPrometheusConfig struct {
	// Enable Managed Collection.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.ManagedPrometheusConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.MonitoringConfig
type MonitoringConfig struct {
	// Enable Google Cloud Managed Service for Prometheus in the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.MonitoringConfig.managed_prometheus_config
	ManagedPrometheusConfig *ManagedPrometheusConfig `json:"managedPrometheusConfig,omitempty"`

	// Optionally enable GKE metrics.
	//  Only for Attached Clusters.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.MonitoringConfig.cloud_monitoring_config
	CloudMonitoringConfig *CloudMonitoringConfig `json:"cloudMonitoringConfig,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.SecurityPostureConfig
type SecurityPostureConfig struct {
	// Sets which mode to use for vulnerability scanning.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.SecurityPostureConfig.vulnerability_mode
	VulnerabilityMode *string `json:"vulnerabilityMode,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.WorkloadIdentityConfig
type WorkloadIdentityConfig struct {
	// The OIDC issuer URL for this cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.WorkloadIdentityConfig.issuer_uri
	IssuerURI *string `json:"issuerURI,omitempty"`

	// The Workload Identity Pool associated to the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.WorkloadIdentityConfig.workload_pool
	WorkloadPool *string `json:"workloadPool,omitempty"`

	// The ID of the OIDC Identity Provider (IdP) associated to the Workload
	//  Identity Pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.WorkloadIdentityConfig.identity_provider
	IdentityProvider *string `json:"identityProvider,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedCluster
type AttachedClusterObservedState struct {
	// Output only. The region where this cluster runs.
	//
	//  For EKS clusters, this is a AWS region. For AKS clusters,
	//  this is an Azure region.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.cluster_region
	ClusterRegion *string `json:"clusterRegion,omitempty"`

	// Required. Fleet configuration.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.fleet
	Fleet *FleetObservedState `json:"fleet,omitempty"`

	// Output only. The current state of the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.state
	State *string `json:"state,omitempty"`

	// Output only. A globally unique identifier for the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. If set, there are currently changes in flight to the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The time at which this cluster was registered.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this cluster was last updated.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The Kubernetes version of the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.kubernetes_version
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`

	// Output only. Workload Identity settings.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.workload_identity_config
	WorkloadIdentityConfig *WorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty"`

	// Output only. A set of errors found in the cluster.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AttachedCluster.errors
	Errors []AttachedClusterError `json:"errors,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.Fleet
type FleetObservedState struct {
	// Output only. The name of the managed Hub Membership resource associated to
	//  this cluster.
	//
	//  Membership names are formatted as
	//  `projects/<project-number>/locations/global/membership/<cluster-id>`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.Fleet.membership
	Membership *string `json:"membership,omitempty"`
}
