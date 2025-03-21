// Copyright 2024 Google LLC
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

package v1beta1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedClusterError
type AttachedClusterError struct {
	// Human-friendly description of the error.
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedClustersAuthorization
type AttachedClustersAuthorization struct {
	// Optional. Users that can perform operations as a cluster admin. A managed
	//  ClusterRoleBinding will be created to grant the `cluster-admin` ClusterRole
	//  to the users. Up to ten admin users can be provided.
	//
	//  For more info on RBAC, see
	//  https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	AdminUsers []string `json:"adminUsers,omitempty"`

	/*NOTYET
	// Optional. Groups of users that can perform operations as a cluster admin. A
	//  managed ClusterRoleBinding will be created to grant the `cluster-admin`
	//  ClusterRole to the groups. Up to ten admin groups can be provided.
	//
	//  For more info on RBAC, see
	//  https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	AdminGroups []string `json:"adminGroups,omitempty"`
	*/
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedOidcConfig
type AttachedOidcConfig struct {
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="IssuerURL field is immutable"
	// Immutable. A JSON Web Token (JWT) issuer URI. `issuer` must start with `https://`.
	IssuerURL string `json:"issuerUrl"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Jwks field is immutable"
	// Immutable, Optional. OIDC verification keys in JWKS format (RFC 7517).
	//  It contains a list of OIDC verification keys that can be used to verify
	//  OIDC JWTs.
	//
	//  This field is required for cluster that doesn't have a publicly available
	//  discovery endpoint. When provided, it will be directly used
	//  to verify the OIDC JWT asserted by the IDP.
	Jwks []byte `json:"jwks,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedProxyConfig
type AttachedProxyConfig struct {
	// The Kubernetes Secret resource that contains the HTTP(S) proxy
	//  configuration. The secret must be a JSON encoded proxy configuration
	//  as described in
	KubernetesSecret KubernetesSecret `json:"kubernetesSecret"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.BinaryAuthorization
type BinaryAuthorization struct {
	// Mode of operation for binauthz policy evaluation. If unspecified, defaults
	//  to DISABLED.
	// Possible values: ["DISABLED", "PROJECT_SINGLETON_POLICY_ENFORCE"].
	EvaluationMode *string `json:"evaluationMode,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.Fleet
type Fleet struct {
	// The id of the Fleet host project where this cluster will be registered.
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// Output only. The name of the managed Hub Membership resource associated to
	//  this cluster.
	//
	//  Membership names are formatted as
	//  `projects/<project-number>/locations/global/membership/<cluster-id>`.
	Membership *string `json:"membership,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.KubernetesSecret
type KubernetesSecret struct {
	// Name of the kubernetes secret.
	Name string `json:"name"`

	// Namespace in which the kubernetes secret is stored.
	Namespace string `json:"namespace"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.LoggingComponentConfig
type LoggingComponentConfig struct {
	// The components to be enabled. Possible values: ["SYSTEM_COMPONENTS", "WORKLOADS"].
	EnableComponents []string `json:"enableComponents,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.LoggingConfig
type LoggingConfig struct {
	// The configuration of the logging components;
	ComponentConfig *LoggingComponentConfig `json:"componentConfig,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.ManagedPrometheusConfig
type ManagedPrometheusConfig struct {
	// Enable Managed Collection.
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.MonitoringConfig
type MonitoringConfig struct {
	// Enable Google Cloud Managed Service for Prometheus in the cluster.
	ManagedPrometheusConfig *ManagedPrometheusConfig `json:"managedPrometheusConfig,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.WorkloadIdentityConfig
type WorkloadIdentityConfig struct {
	// The OIDC issuer URL for this cluster.
	IssuerUri *string `json:"issuerUri,omitempty"`

	// The Workload Identity Pool associated to the cluster.
	WorkloadPool *string `json:"workloadPool,omitempty"`

	// The ID of the OIDC Identity Provider (IdP) associated to the Workload
	//  Identity Pool.
	IdentityProvider *string `json:"identityProvider,omitempty"`
}
