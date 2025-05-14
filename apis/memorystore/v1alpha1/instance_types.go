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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MemorystoreInstanceGVK = GroupVersion.WithKind("MemorystoreInstance")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MemorystoreInstanceSpec defines the desired state of MemorystoreInstance
// +kcc:proto=google.cloud.memorystore.v1beta.Instance
type MemorystoreInstanceSpec struct {
	// Optional. Immutable. The MemorystoreInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Optional. Number of replica nodes per shard. If omitted the default is 0
	//  replicas.
	ReplicaCount *int32 `json:"replicaCount,omitempty"`

	// Optional. Immutable. Authorization mode of the instance.
	AuthorizationMode *string `json:"authorizationMode,omitempty"`

	// Optional. Immutable. In-transit encryption mode of the instance.
	TransitEncryptionMode *string `json:"transitEncryptionMode,omitempty"`

	// Optional. Number of shards for the instance.
	ShardCount *int32 `json:"shardCount,omitempty"`

	// Optional. Immutable. Machine type for individual nodes of the instance.
	NodeType *string `json:"nodeType,omitempty"`

	// Optional. Persistence configuration of the instance.
	PersistenceConfig *PersistenceConfig `json:"persistenceConfig,omitempty"`

	// Optional. Immutable. Engine version of the instance.
	EngineVersion *string `json:"engineVersion,omitempty"`

	// Optional. User-provided engine configurations for the instance.
	EngineConfigs map[string]string `json:"engineConfigs,omitempty"`

	// Optional. Immutable. Zone distribution configuration of the instance for
	//  node allocatiteon.
	ZoneDistributionConfig *ZoneDistributionConfig `json:"zoneDistributionConfig,omitempty"`

	// Optional. If set to true deletion of the instance will fail.
	DeletionProtectionEnabled *bool `json:"deletionProtectionEnabled,omitempty"`

	// Required. Immutable. User inputs for the auto-created PSC connections.
	PscAutoConnectionsSpec []PscAutoConnectionSpec `json:"pscAutoConnections,omitempty"`
}

// kcc specific struct to separate input and output fields in
// google.cloud.memorystore.v1beta.PscAutoConnection
type PscAutoConnectionSpec struct {

	// Required. The consumer project_id where PSC connections are established.
	//  This should be the same project_id that the cluster is being created in.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Required. The network where the PSC endpoints are created, in the form of
	//  projects/{project_id}/global/networks/{network_id}.
	// +required
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`
}

type Parent struct {
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
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
type MemorystoreInstanceObservedState struct {
	// Identifier. Unique name of the instance.
	//  Format: projects/{project}/locations/{location}/instances/{instance}
	Name *string `json:"name,omitempty"`

	// Output only. Creation timestamp of the instance.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Latest update timestamp of the instance.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the instance.
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the state of the instance.
	StateInfo *Instance_StateInfo `json:"stateInfo,omitempty"`

	// Output only. System assigned, unique identifier for the instance.
	Uid *string `json:"uid,omitempty"`

	// Optional. Immutable. Authorization mode of the instance.
	AuthorizationMode *string `json:"authorizationMode,omitempty"`

	// Optional. Immutable. In-transit encryption mode of the instance.
	TransitEncryptionMode *string `json:"transitEncryptionMode,omitempty"`

	// Output only. Endpoints clients can connect to the instance through.
	//  Currently only one discovery endpoint is supported.
	DiscoveryEndpoints []DiscoveryEndpoint `json:"discoveryEndpoints,omitempty"`

	// Optional. Immutable. Machine type for individual nodes of the instance.
	NodeType *string `json:"nodeType,omitempty"`

	// Optional. Immutable. Engine version of the instance.
	// https://cloud.google.com/memorystore/docs/valkey/supported-versions
	EngineVersion *string `json:"engineVersion,omitempty"`

	// Output only. Configuration of individual nodes of the instance.
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	// Optional. Immutable. Zone distribution configuration of the instance for
	//  node allocation.
	ZoneDistributionConfig *ZoneDistributionConfig `json:"zoneDistributionConfig,omitempty"`

	// Output only. Resource details of the auto-created PSC connections.
	PscAutoConnections []PscAutoConnection `json:"pscAutoConnections,omitempty"`
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
