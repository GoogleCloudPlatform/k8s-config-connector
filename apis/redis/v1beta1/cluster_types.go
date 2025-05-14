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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var RedisClusterGVK = GroupVersion.WithKind("RedisCluster")

// RedisClusterSpec defines the desired state of RedisCluster
// +kcc:proto=google.cloud.redis.cluster.v1.Cluster
type RedisClusterSpec struct {

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. Location of the resource. */
	Location *string `json:"location"`

	// The RedisCluster name. If not given, the metadata.name will be used.
	// + optional
	ResourceID *string `json:"resourceID,omitempty"`

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
	PscConfigs []PscConfigSpec `json:"pscConfigs,omitempty"`

	// Optional. The type of a redis node in the cluster. NodeType determines the
	//  underlying machine-type of a redis node.
	NodeType *string `json:"nodeType,omitempty"`

	// Optional. Persistence config (RDB, AOF) for the cluster.
	PersistenceConfig *ClusterPersistenceConfig `json:"persistenceConfig,omitempty"`

	// Optional. Key/Value pairs of customer overrides for mutable Redis Configs
	RedisConfigs map[string]string `json:"redisConfigs,omitempty"`

	// Optional. The number of replica nodes per shard.
	ReplicaCount *int32 `json:"replicaCount,omitempty"`

	// Optional. This config will be used to determine how the customer wants us
	//  to distribute cluster resources within the region.
	ZoneDistributionConfig *ZoneDistributionConfig `json:"zoneDistributionConfig,omitempty"`

	// Optional. The delete operation will fail when the value is set to true.
	DeletionProtectionEnabled *bool `json:"deletionProtectionEnabled,omitempty"`
}

type PscConfigSpec struct {
	// Required. The network where the IP address of the discovery endpoint will
	//  be reserved, in the form of
	//  projects/{network_project}/global/networks/{network_id}.
	// +required
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`
}

// RedisClusterStatus defines the config connector machine state of RedisCluster
type RedisClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the RedisCluster resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *RedisClusterObservedState `json:"observedState,omitempty"`
}

// RedisClusterSpec defines the desired state of RedisCluster
// +kcc:proto=google.cloud.redis.cluster.v1.Cluster
type RedisClusterObservedState struct {

	// Output only. The timestamp associated with the cluster creation request.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The current state of this cluster.
	//  Can be CREATING, READY, UPDATING, DELETING and SUSPENDED
	State *string `json:"state,omitempty"`

	// Output only. System assigned, unique identifier for the cluster.
	Uid *string `json:"uid,omitempty"`

	// Output only. Redis memory size in GB for the entire cluster rounded up to
	//  the next integer.
	SizeGb *int32 `json:"sizeGb,omitempty"`

	// Output only. Endpoints created on each given network, for Redis clients to
	//  connect to the cluster. Currently only one discovery endpoint is supported.
	DiscoveryEndpoints []DiscoveryEndpoint `json:"discoveryEndpoints,omitempty"`

	// Output only. PSC connections for discovery of the cluster topology and
	//  accessing the cluster.
	PscConnections []PscConnection `json:"pscConnections,omitempty"`

	// Output only. Additional information about the current state of the cluster.
	StateInfo *Cluster_StateInfo `json:"stateInfo,omitempty"`

	// Output only. Precise value of redis memory size in GB for the entire
	//  cluster.
	PreciseSizeGb *float64 `json:"preciseSizeGb,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcprediscluster;gcpredisclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=beta"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// RedisCluster is the Schema for the RedisCluster API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
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

func init() {
	SchemeBuilder.Register(&RedisCluster{}, &RedisClusterList{})
}
