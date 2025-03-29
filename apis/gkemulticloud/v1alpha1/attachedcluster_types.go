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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GkeMultiCloudAttachedClusterGVK = GroupVersion.WithKind("GkeMultiCloudAttachedCluster")

// GkeMultiCloudAttachedClusterSpec defines the desired state of GkeMultiCloudAttachedCluster
// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedCluster
type GkeMultiCloudAttachedClusterSpec struct {
	// The GkeMultiCloudAttachedCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

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

// GkeMultiCloudAttachedClusterStatus defines the config connector machine state of GkeMultiCloudAttachedCluster
type GkeMultiCloudAttachedClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GkeMultiCloudAttachedCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GkeMultiCloudAttachedClusterObservedState `json:"observedState,omitempty"`
}

// GkeMultiCloudAttachedClusterObservedState is the state of the GkeMultiCloudAttachedCluster resource as most recently observed in GCP.
// +kcc:proto=google.cloud.gkemulticloud.v1.AttachedCluster
type GkeMultiCloudAttachedClusterObservedState struct {
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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpgkemulticloudattachedcluster;gcpgkemulticloudattachedclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GkeMultiCloudAttachedCluster is the Schema for the GkeMultiCloudAttachedCluster API
// +k8s:openapi-gen=true
type GkeMultiCloudAttachedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GkeMultiCloudAttachedClusterSpec   `json:"spec,omitempty"`
	Status GkeMultiCloudAttachedClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GkeMultiCloudAttachedClusterList contains a list of GkeMultiCloudAttachedCluster
type GkeMultiCloudAttachedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GkeMultiCloudAttachedCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GkeMultiCloudAttachedCluster{}, &GkeMultiCloudAttachedClusterList{})
}
