// Copyright 2024 Google LLC
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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ContainerAttachedClusterGVK = GroupVersion.WithKind("ContainerAttachedCluster")

// ContainerAttachedClusterSpec defines the desired state of ContainerAttachedCluster
// +kcc:spec:proto=google.cloud.gkemulticloud.v1.AttachedCluster
type ContainerAttachedClusterSpec struct {
	/* The ID of the project in which the resource belongs.*/
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Optional.
	// The ContainerAttachedCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable. The location for the resource.
	Location string `json:"location"`

	// Optional. A human readable description of this Attached cluster.
	//  Cannot be longer than 255 UTF-8 encoded bytes.
	Description *string `json:"description,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="OidcConfig field is immutable"
	/* Required. OpenID Connect (OIDC) discovery information of the target cluster.

	Kubernetes Service Account (KSA) tokens are JWT tokens signed by the cluster
	API server. This field indicates how GCP services	validate KSA tokens in order
	to allow system workloads (such as GKE Connect and telemetry agents) to
	authenticate back to GCP.

	Both clusters with public and private issuer URLs are supported.
	Clusters with public issuers only need to specify the 'issuerUrl' field
	while clusters with private issuers need to provide both 'issuerUrl' and 'jwks'.
	*/
	OidcConfig AttachedOidcConfig `json:"oidcConfig"`

	// Required. The platform version for the cluster (e.g. `1.30.0-gke.1`).
	PlatformVersion string `json:"platformVersion"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Distribution field is immutable"
	// Immutable. The Kubernetes distribution of the underlying attached cluster.
	//
	//  Supported values: ["eks", "aks", "generic"].
	Distribution string `json:"distribution"`

	// Required. Fleet configuration.
	Fleet Fleet `json:"fleet"`

	/*NOTYET
	// Allows clients to perform consistent read-modify-writes
	//  through optimistic concurrency control.
	//
	//  Can be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`
	*/

	// Optional. Annotations on the cluster.
	//
	//  This field has the same restrictions as Kubernetes annotations.
	//  The total size of all keys and values combined is limited to 256k.
	//  Key can have 2 segments: prefix (optional) and name (required),
	//  separated by a slash (/).
	//  Prefix must be a DNS subdomain.
	//  Name must be 63 characters or less, begin and end with alphanumerics,
	//  with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Logging configuration for this cluster.
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// Optional. Configuration related to the cluster RBAC settings.
	Authorization *AttachedClustersAuthorization `json:"authorization,omitempty"`

	// Optional. Monitoring configuration for this cluster.
	MonitoringConfig *MonitoringConfig `json:"monitoringConfig,omitempty"`

	/*NOTYET
	// Optional. Proxy configuration for outbound HTTP(S) traffic.
	ProxyConfig *AttachedProxyConfig `json:"proxyConfig,omitempty"`
	*/

	// Optional. Binary Authorization configuration for this cluster.
	BinaryAuthorization *BinaryAuthorization `json:"binaryAuthorization,omitempty"`

	// Optional. Policy to determine what flags to send on delete.
	DeletionPolicy *string `json:"deletionPolicy,omitempty"`
}

// ContainerAttachedClusterStatus defines the config connector machine state of ContainerAttachedCluster
type ContainerAttachedClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A globally unique identifier for the cluster.
	Uid *string `json:"uid,omitempty"`

	// The region where this cluster runs.
	//
	//  For EKS clusters, this is an AWS region. For AKS clusters,
	//  this is an Azure region.
	ClusterRegion *string `json:"clusterRegion,omitempty"`

	// The Kubernetes version of the cluster.
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`

	// Workload Identity settings.
	WorkloadIdentityConfig []WorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty"`

	// The time at which this cluster was registered.
	CreateTime *string `json:"createTime,omitempty"`

	// The time at which this cluster was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// If set, there are currently changes in flight to the cluster.
	Reconciling *bool `json:"reconciling,omitempty"`

	/* The current state of the cluster.
	Possible values:	STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR,	DEGRADED. */
	State *string `json:"state,omitempty"`

	// A set of errors found in the cluster.
	Errors []AttachedClusterError `json:"errors,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ContainerAttachedClusterObservedState `json:"observedState,omitempty"`
}

// ContainerAttachedClusterObservedState is the state of the ContainerAttachedCluster resource as most recently observed in GCP.
type ContainerAttachedClusterObservedState struct {
	// Output only. The name of the managed Hub Membership resource associated to
	//  this cluster.
	//
	//  Membership names are formatted as
	//  `projects/<project-number>/locations/global/membership/<cluster-id>`.
	//  This field mirrors the Spec.Fleet.Membership field.
	FleetMembership *string `json:"fleetMembership,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcontainerattachedcluster;gcpcontainerattachedclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ContainerAttachedCluster is the Schema for the ContainerAttachedCluster API
// +k8s:openapi-gen=true
type ContainerAttachedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerAttachedClusterSpec   `json:"spec"`
	Status ContainerAttachedClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ContainerAttachedClusterList contains a list of ContainerAttachedCluster
type ContainerAttachedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerAttachedCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerAttachedCluster{}, &ContainerAttachedClusterList{})
}
