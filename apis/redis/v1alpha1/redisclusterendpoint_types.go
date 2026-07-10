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
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	redisv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var RedisClusterEndpointGVK = GroupVersion.WithKind("RedisClusterEndpoint")

// RedisClusterEndpointSpec defines the desired state of RedisClusterEndpoint
type RedisClusterEndpointSpec struct {
	// Required. The RedisCluster reference of the endpoint.
	// +required
	ClusterRef *redisv1beta1.RedisClusterRef `json:"clusterRef"`

	// Optional. The RedisClusterEndpoint name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A list of cluster endpoints.
	// +optional
	// +listType=atomic
	ClusterEndpoints []ClusterEndpoint_ClusterEndpoint `json:"clusterEndpoints,omitempty"`
}

type ClusterEndpoint_ClusterEndpoint struct {
	// A group of PSC connections. They are created in the same VPC network, one
	//  for each service attachment in the cluster.
	// +optional
	// +listType=atomic
	Connections []ClusterEndpoint_ConnectionDetail `json:"connections,omitempty"`
}

type ClusterEndpoint_ConnectionDetail struct {
	// Detailed information of a PSC connection that is created by the customer
	//  who owns the cluster.
	// +optional
	PSCConnection *ClusterEndpoint_PSCConnection `json:"pscConnection,omitempty"`
}

type ClusterEndpoint_PSCConnection struct {
	// Required. The PSC connection id of the forwarding rule connected to the
	//  service attachment.
	PSCConnectionID *string `json:"pscConnectionID"`

	// Required. The IP allocated on the consumer network for the PSC forwarding
	//  rule.
	AddressRef *computev1beta1.ComputeAddressRef `json:"addressRef"`

	// Required. The URI of the consumer side forwarding rule.
	//  Example:
	//  projects/{projectNumOrId}/regions/us-east1/forwardingRules/{resourceId}.
	ForwardingRuleRef *computev1beta1.ForwardingRuleRef `json:"forwardingRuleRef"`

	// Optional. Project ID of the consumer project where the forwarding rule is
	//  created in.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// Required. The consumer network where the IP address resides, in the form of
	//  projects/{project_id}/global/networks/{network_id}.
	NetworkRef *computerefs.ComputeNetworkRef `json:"networkRef"`

	// Required. The service attachment which is the target of the PSC connection,
	//  in the form of
	//  projects/{project-id}/regions/{region}/serviceAttachments/{service-attachment-id}.
	ServiceAttachmentRef *refsv1beta1.ComputeServiceAttachmentRef `json:"serviceAttachmentRef"`
}

// RedisClusterEndpointStatus defines the config connector machine state of RedisClusterEndpoint
type RedisClusterEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the RedisClusterEndpoint resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *RedisClusterEndpointObservedState `json:"observedState,omitempty"`
}

// RedisClusterEndpointObservedState is the state of the RedisClusterEndpoint resource as most recently observed in GCP.
type RedisClusterEndpointObservedState struct {
	// Optional. A list of cluster endpoints.
	// +optional
	// +listType=atomic
	ClusterEndpoints []ClusterEndpoint_ClusterEndpointObservedState `json:"clusterEndpoints,omitempty"`
}

type ClusterEndpoint_ClusterEndpointObservedState struct {
	// A group of PSC connections. They are created in the same VPC network, one
	//  for each service attachment in the cluster.
	// +optional
	// +listType=atomic
	Connections []ClusterEndpoint_ConnectionDetailObservedState `json:"connections,omitempty"`
}

type ClusterEndpoint_ConnectionDetailObservedState struct {
	// Detailed information of a PSC connection that is created by the customer
	//  who owns the cluster.
	// +optional
	PSCConnection *ClusterEndpoint_PSCConnectionObservedState `json:"pscConnection,omitempty"`
}

type ClusterEndpoint_PSCConnectionObservedState struct {
	// Output only. The status of the PSC connection.
	//  Please note that this value is updated periodically.
	//  To get the latest status of a PSC connection, follow
	//  https://cloud.google.com/vpc/docs/configure-private-service-connect-services#endpoint-details.
	PSCConnectionStatus *string `json:"pscConnectionStatus,omitempty"`

	// Output only. Type of the PSC connection.
	ConnectionType *string `json:"connectionType,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpredisclusterendpoint;gcpredisclusterendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// RedisClusterEndpoint is the Schema for the RedisClusterEndpoint API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type RedisClusterEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   RedisClusterEndpointSpec   `json:"spec,omitempty"`
	Status RedisClusterEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RedisClusterEndpointList contains a list of RedisClusterEndpoint
type RedisClusterEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisClusterEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedisClusterEndpoint{}, &RedisClusterEndpointList{})
}
