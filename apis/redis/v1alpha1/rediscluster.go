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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ProjectRef struct {
	/* The external name of the referenced resource */
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

// type ServiceConnectionPolicyRef struct {
// 	// /* The external name of the referenced resource */
// 	// External string `json:"external,omitempty"`

// 	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
// 	Name string `json:"name,omitempty"`
// 	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
// 	Namespace string `json:"namespace,omitempty"`
// }

type RedisClusterSpec struct {
	/* The project that this resource belongs to. */
	ProjectRef ProjectRef `json:"projectRef"`

	/* Immutable. The location where the cluster should reside. */
	Location string `json:"location"`

	/* Immutable. Optional. The apiId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The number of replica nodes per shard.
	ReplicaCount *int32 `json:"replicaCount,omitempty"`
	// Optional. The authorization mode of the Redis cluster.
	//  If not provided, auth feature is disabled for the cluster.
	AuthorizationMode *string `json:"authorizationMode,omitempty"`
	// Optional. The in-transit encryption for the Redis cluster.
	//  If not provided, encryption  is disabled for the cluster.
	TransitEncryptionMode *string `json:"transitEncryptionMode,omitempty"`
	// Required. Number of shards for the Redis cluster.
	ShardCount *int32 `json:"shardCount,omitempty"`
	// Required. Each PscConfig configures the consumer network where IPs will
	//  be designated to the cluster for client access through Private Service
	//  Connect Automation. Currently, only one PscConfig is supported.
	PscConfigs []PscConfig `json:"pscConfigs,omitempty"`
}

type RedisClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeNetwork's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	ObservedState *RedisClusterObservedState `json:"observedState,omitempty"`
}

type RedisClusterObservedState struct {
	// // Output only. The timestamp associated with the cluster creation request.
	// CreateTime *Timestamp `json:"createTime,omitempty"`

	// // Output only. The current state of this cluster.
	// //  Can be CREATING, READY, UPDATING, DELETING and SUSPENDED
	// State *State `json:"state,omitempty"`

	// // Output only. System assigned, unique identifier for the cluster.
	// Uid *string `json:"uid,omitempty"`

	// Output only. Redis memory size in GB for the entire cluster.
	SizeGb *int32 `json:"sizeGb,omitempty"`

	// Output only. Endpoints created on each given network, for Redis clients to
	//  connect to the cluster. Currently only one discovery endpoint is supported.
	DiscoveryEndpoints []DiscoveryEndpoint `json:"discoveryEndpoints,omitempty"`
	// Output only. PSC connections for discovery of the cluster topology and
	//  accessing the cluster.
	PscConnections []PscConnection `json:"pscConnections,omitempty"`
	// Output only. Additional information about the current state of the cluster.
	StateInfo *Cluster_StateInfo `json:"stateInfo,omitempty"`
}

type Cluster_StateInfo struct {
	// Describes ongoing update on the cluster when cluster state is UPDATING.
	UpdateInfo *Cluster_StateInfo_UpdateInfo `json:"updateInfo,omitempty"`
}

type Cluster_StateInfo_UpdateInfo struct {
	// Target number of shards for redis cluster
	TargetShardCount *int32 `json:"targetShardCount,omitempty"`
	// Target number of replica nodes per shard.
	TargetReplicaCount *int32 `json:"targetReplicaCount,omitempty"`
}

type PscConfig struct {
	// Required. The network where the IP address of the discovery endpoint will
	//  be reserved, in the form of
	//  projects/{network_project}/global/networks/{network_id}.
	Network *string `json:"network,omitempty"`

	// ServiceConnectionPolicyRef *ServiceConnectionPolicyRef `json:"serviceConnectionPolicyRef,omitempty"`
}

type DiscoveryEndpoint struct {
	// Output only. Address of the exposed Redis endpoint used by clients to
	//  connect to the service. The address could be either IP or hostname.
	Address *string `json:"address,omitempty"`
	// Output only. The port number of the exposed Redis endpoint.
	Port *int32 `json:"port,omitempty"`
	// Output only. Customer configuration for where the endpoint is created and
	//  accessed from.
	PscConfig *PscConfig `json:"pscConfig,omitempty"`
}

type PscConnection struct {
	// Output only. The PSC connection id of the forwarding rule connected to the
	//  service attachment.
	PscConnectionId *string `json:"pscConnectionId,omitempty"`
	// Output only. The IP allocated on the consumer network for the PSC
	//  forwarding rule.
	Address *string `json:"address,omitempty"`
	// Output only. The URI of the consumer side forwarding rule.
	//  Example:
	//  projects/{projectNumOrId}/regions/us-east1/forwardingRules/{resourceId}.
	ForwardingRule *string `json:"forwardingRule,omitempty"`
	// Output only. The consumer project_id where the forwarding rule is created
	//  from.
	ProjectId *string `json:"projectId,omitempty"`
	// The consumer network where the IP address resides, in the form of
	//  projects/{project_id}/global/networks/{network_id}.
	Network *string `json:"network,omitempty"`
}

// RedisCluster is the Schema for the redis API
// +k8s:openapi-gen=true
// +kubebuilder:resource:categories=gcp,shortName=gcpredisclusters;gcpredisclusters
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:subresource:status
type RedisCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisClusterSpec   `json:"spec,omitempty"`
	Status RedisClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisClusterList contains a list of RedisCluster
type RedisClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisCluster `json:"items"`
}

// func init() {
// 	SchemeBuilder.Register(&RedisCluster{}, &RedisClusterList{})
// }
