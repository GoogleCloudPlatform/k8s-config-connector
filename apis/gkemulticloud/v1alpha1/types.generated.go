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
