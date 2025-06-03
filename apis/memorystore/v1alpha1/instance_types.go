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

var MemorystoreInstanceGVK = GroupVersion.WithKind("MemorystoreInstance")

// MemorystoreInstanceSpec defines the desired state of MemorystoreInstance
// +kcc:spec:proto=google.cloud.memorystore.v1beta.Instance
type MemorystoreInstanceSpec struct {
	// The MemorystoreInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Identifier. Unique name of the instance.
	//  Format: projects/{project}/locations/{location}/instances/{instance}
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.name
	Name *string `json:"name,omitempty"`

	// Optional. Labels to represent user-provided metadata.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Number of replica nodes per shard. If omitted the default is 0
	//  replicas.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.replica_count
	ReplicaCount *int32 `json:"replicaCount,omitempty"`

	// Optional. Immutable. Authorization mode of the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.authorization_mode
	AuthorizationMode *string `json:"authorizationMode,omitempty"`

	// Optional. Immutable. In-transit encryption mode of the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.transit_encryption_mode
	TransitEncryptionMode *string `json:"transitEncryptionMode,omitempty"`

	// Optional. Number of shards for the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.shard_count
	ShardCount *int32 `json:"shardCount,omitempty"`

	// Optional. Immutable. Machine type for individual nodes of the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.node_type
	NodeType *string `json:"nodeType,omitempty"`

	// Optional. Persistence configuration of the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.persistence_config
	PersistenceConfig *PersistenceConfig `json:"persistenceConfig,omitempty"`

	// Optional. Immutable. Engine version of the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.engine_version
	EngineVersion *string `json:"engineVersion,omitempty"`

	// Optional. User-provided engine configurations for the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.engine_configs
	EngineConfigs map[string]string `json:"engineConfigs,omitempty"`

	// Optional. Immutable. Zone distribution configuration of the instance for
	//  node allocation.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.zone_distribution_config
	ZoneDistributionConfig *ZoneDistributionConfig `json:"zoneDistributionConfig,omitempty"`

	// Optional. If set to true deletion of the instance will fail.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.deletion_protection_enabled
	DeletionProtectionEnabled *bool `json:"deletionProtectionEnabled,omitempty"`

	// Required. Immutable. User inputs and resource details of the auto-created
	//  PSC connections.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.psc_auto_connections
	PSCAutoConnectionsSpec []PSCAutoConnectionSpec `json:"pscAutoConnections,omitempty"`

	// Optional. Endpoints for the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.endpoints
	Endpoints []Instance_InstanceEndpointSpec `json:"endpoints,omitempty"`

	// Optional. The mode config for the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.mode
	Mode *string `json:"mode,omitempty"`
}

type Parent struct {
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
}

// +kcc:proto=google.cloud.memorystore.v1beta.Instance.InstanceEndpoint
type Instance_InstanceEndpointSpec struct {
	// Optional. A group of PSC connections. They are created in the same VPC
	//  network, one for each service attachment in the cluster.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.InstanceEndpoint.connections
	Connections []Instance_ConnectionDetailSpec `json:"connections,omitempty"`
}

// +kcc:proto=google.cloud.memorystore.v1beta.Instance.ConnectionDetail
type Instance_ConnectionDetailSpec struct {
	// Detailed information of a PSC connection that is created through
	//  service connectivity automation.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.ConnectionDetail.psc_auto_connection
	PSCAutoConnectionSpec *PSCAutoConnectionSpec `json:"pscAutoConnection,omitempty"`

	// Detailed information of a PSC connection that is created by the user.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.ConnectionDetail.psc_connection
	PSCConnectionSpec *PSCConnectionSpec `json:"pscConnection,omitempty"`
}

// kcc specific struct to separate input and output fields in
// +kcc:proto=google.cloud.memorystore.v1beta.PscAutoConnection
type PscAutoConnectionSpec struct {

	// Required. The consumer project_id where PSC connections are established.
	//  This should be the same project_id that the instance is being created in.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.PscAutoConnection.project_id
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Required. The network where the PSC endpoints are created, in the form of
	//  projects/{project_id}/global/networks/{network_id}.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.PscAutoConnection.network
	NetworkRef *refs.ComputeNetworkRef `json:"networkRef,omitempty"`
}

// google.cloud.memorystore.v1beta.PscConnection
type PSCConnectionSpec struct {

	// Required. The IP allocated on the consumer network for the PSC forwarding
	//  rule.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.PscConnection.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// Required. The URI of the consumer side forwarding rule.
	//  Format:
	//  projects/{project}/regions/{region}/forwardingRules/{forwarding_rule}
	// +kcc:proto:field=google.cloud.memorystore.v1beta.PscConnection.forwarding_rule
	ForwardingRule *string `json:"forwardingRule,omitempty"`

	// Required. The consumer network where the IP address resides, in the form of
	//  projects/{project_id}/global/networks/{network_id}.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.PscConnection.network
	NetworkRef *string `json:"network,omitempty"`

	// Required. The service attachment which is the target of the PSC connection,
	//  in the form of
	//  projects/{project-id}/regions/{region}/serviceAttachments/{service-attachment-id}.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.PscConnection.service_attachment
	ServiceAttachment *string `json:"serviceAttachment,omitempty"`
}

// MemorystoreInstanceStatus defines the config connector machine state of MemorystoreInstance
type MemorystoreInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MemorystoreInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MemorystoreInstanceObservedState `json:"observedState,omitempty"`
}

// MemorystoreInstanceObservedState is the state of the MemorystoreInstance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.memorystore.v1beta.Instance
type MemorystoreInstanceObservedState struct {
	// Output only. Creation timestamp of the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Latest update timestamp of the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the state of the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.state_info
	StateInfo *Instance_StateInfo `json:"stateInfo,omitempty"`

	// Output only. System assigned, unique identifier for the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Endpoints clients can connect to the instance through.
	//  Currently only one discovery endpoint is supported.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.discovery_endpoints
	DiscoveryEndpoints []DiscoveryEndpoint `json:"discoveryEndpoints,omitempty"`

	// Output only. Configuration of individual nodes of the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.node_config
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	// Required. Immutable. User inputs and resource details of the auto-created
	//  PSC connections.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.psc_auto_connections
	PSCAutoConnections []PSCAutoConnectionObservedState `json:"pscAutoConnections,omitempty"`

	// Optional. Endpoints for the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1beta.Instance.endpoints
	Endpoints []Instance_InstanceEndpointObservedState `json:"endpoints,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmemorystoreinstance;gcpmemorystoreinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MemorystoreInstance is the Schema for the MemorystoreInstance API
// +k8s:openapi-gen=true
type MemorystoreInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MemorystoreInstanceSpec   `json:"spec,omitempty"`
	Status MemorystoreInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MemorystoreInstanceList contains a list of MemorystoreInstance
type MemorystoreInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MemorystoreInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MemorystoreInstance{}, &MemorystoreInstanceList{})
}
