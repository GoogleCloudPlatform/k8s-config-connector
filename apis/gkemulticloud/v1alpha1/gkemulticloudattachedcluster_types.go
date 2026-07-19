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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GKEMulticloudAttachedClusterGVK = GroupVersion.WithKind("GKEMulticloudAttachedCluster")

// GKEMulticloudAttachedClusterSpec defines the desired state of GKEMulticloudAttachedCluster
// +kcc:spec:proto=google.cloud.gkemulticloud.v1.AttachedCluster
type GKEMulticloudAttachedClusterSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable. The location of this resource.
	// +required
	Location *string `json:"location"`

	// Optional. A human readable description of this cluster.
	//  Cannot be longer than 255 UTF-8 encoded bytes.
	// +optional
	Description *string `json:"description,omitempty"`

	// Required. OpenID Connect (OIDC) configuration for the cluster.
	// +required
	OIDCConfig *AttachedOIDCConfig `json:"oidcConfig,omitempty"`

	// Required. The platform version for the cluster (e.g. `1.19.0-gke.1000`).
	//  You can list all supported versions on a given Google Cloud region by
	//  calling GetAttachedServerConfig.
	// +required
	PlatformVersion *string `json:"platformVersion,omitempty"`

	// Required. The Kubernetes distribution of the underlying attached cluster.
	//  Supported values: ["eks", "aks", "generic"].
	// +required
	Distribution *string `json:"distribution,omitempty"`

	// Required. Fleet configuration.
	// +required
	Fleet *Fleet `json:"fleet,omitempty"`

	// Optional. Annotations on the cluster.
	//  This field has the same restrictions as Kubernetes annotations.
	//  The total size of all keys and values combined is limited to 256k.
	//  Key can have 2 segments: prefix (optional) and name (required),
	//  separated by a slash (/).
	//  Prefix must be a DNS subdomain.
	//  Name must be 63 characters or less, begin and end with alphanumerics,
	//  with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Logging configuration for this cluster.
	// +optional
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// Optional. Configuration related to the cluster RBAC settings.
	// +optional
	Authorization *AttachedClustersAuthorization `json:"authorization,omitempty"`

	// Optional. Monitoring configuration for this cluster.
	// +optional
	MonitoringConfig *MonitoringConfig `json:"monitoringConfig,omitempty"`

	// Optional. Proxy configuration for outbound HTTP(S) traffic.
	// +optional
	ProxyConfig *AttachedProxyConfig `json:"proxyConfig,omitempty"`

	// Optional. Binary Authorization configuration for this cluster.
	// +optional
	BinaryAuthorization *BinaryAuthorization `json:"binaryAuthorization,omitempty"`

	// Optional. Security Posture configuration for this cluster.
	// +optional
	SecurityPostureConfig *SecurityPostureConfig `json:"securityPostureConfig,omitempty"`

	// Optional. Input only. Tag keys/values directly bound to this resource.
	//  See Tags overview for more details on Google Cloud Platform tags.
	// +optional
	Tags map[string]string `json:"tags,omitempty"`

	// The GKEMulticloudAttachedCluster name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.gkemulticloud.v1.Fleet
type Fleet struct {
	// The host project of the fleet.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`
}

// +kcc:observedstate:proto=google.cloud.gkemulticloud.v1.Fleet
type FleetObservedState struct {
	// Output only. The name of the managed Hub Membership resource associated to this cluster.
	// +optional
	Membership *string `json:"membership,omitempty"`
}

// GKEMulticloudAttachedClusterStatus defines the config connector machine state of GKEMulticloudAttachedCluster
type GKEMulticloudAttachedClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GKEMulticloudAttachedCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GKEMulticloudAttachedClusterObservedState `json:"observedState,omitempty"`
}

// GKEMulticloudAttachedClusterObservedState is the state of the GKEMulticloudAttachedCluster resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.gkemulticloud.v1.AttachedCluster
type GKEMulticloudAttachedClusterObservedState struct {
	// Output only. The region where this cluster runs.
	//  For EKS clusters, this is an AWS region. For AKS clusters,
	//  this is an Azure region.
	// +optional
	ClusterRegion *string `json:"clusterRegion,omitempty"`

	// Required. Fleet configuration.
	// +optional
	Fleet *FleetObservedState `json:"fleet,omitempty"`

	// Output only. The current state of the cluster.
	// +optional
	State *string `json:"state,omitempty"`

	// Output only. A globally unique identifier for the cluster.
	// +optional
	Uid *string `json:"uid,omitempty"`

	// Output only. If set, there are currently changes in flight to the cluster.
	// +optional
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The time at which this cluster was registered.
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this cluster was last updated.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The Kubernetes version of the cluster.
	// +optional
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`

	// Output only. Workload Identity settings.
	// +optional
	WorkloadIdentityConfig *WorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty"`

	// Output only. A set of errors found in the cluster.
	// +optional
	Errors []AttachedClusterError `json:"errors,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkemulticloudattachedcluster;gcpgkemulticloudattachedclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEMulticloudAttachedCluster is the Schema for the GKEMulticloudAttachedCluster API
// +k8s:openapi-gen=true
type GKEMulticloudAttachedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GKEMulticloudAttachedClusterSpec   `json:"spec,omitempty"`
	Status GKEMulticloudAttachedClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GKEMulticloudAttachedClusterList contains a list of GKEMulticloudAttachedCluster
type GKEMulticloudAttachedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEMulticloudAttachedCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEMulticloudAttachedCluster{}, &GKEMulticloudAttachedClusterList{})
}
