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

var BlockchainNodeEngineBlockchainNodeGVK = GroupVersion.WithKind("BlockchainNodeEngineBlockchainNode")

// BlockchainNodeEngineBlockchainNodeSpec defines the desired state of BlockchainNodeEngineBlockchainNode
// +kcc:spec:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode
type BlockchainNodeEngineBlockchainNodeSpec struct {
	/* Immutable. The Project that this resource belongs to. */
	ProjectRef refsv1beta1.ProjectRef `json:"projectRef"`

	/* Immutable. The location of this resource. */
	Location string `json:"location"`

	/* Immutable. Optional. The BlockchainNode name. If not given, the metadata.name will be used. */
	ResourceID *string `json:"resourceID,omitempty"`

	// Ethereum-specific blockchain node details.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.ethereum_details
	EthereumDetails *BlockchainNode_EthereumDetails `json:"ethereumDetails,omitempty"`

	// User-provided key-value pairs.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The blockchain type of the node.
	// +kubebuilder:validation:Enum=ETHEREUM
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.blockchain_type
	BlockchainType *string `json:"blockchainType,omitempty"`

	// Optional. When true, the node is only accessible via Private Service
	//  Connect; no public endpoints are exposed. Otherwise, the node is only
	//  accessible via public endpoints. Warning: These nodes are deprecated,
	//  please use public endpoints instead.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.private_service_connect_enabled
	PrivateServiceConnectEnabled *bool `json:"privateServiceConnectEnabled,omitempty"`
}

// BlockchainNodeEngineBlockchainNodeStatus defines the config connector machine state of BlockchainNodeEngineBlockchainNode
type BlockchainNodeEngineBlockchainNodeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BlockchainNodeEngineBlockchainNode resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BlockchainNodeEngineBlockchainNodeObservedState `json:"observedState,omitempty"`
}

// BlockchainNodeEngineBlockchainNodeObservedState is the state of the BlockchainNodeEngineBlockchainNode resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode
type BlockchainNodeEngineBlockchainNodeObservedState struct {
	// Ethereum-specific blockchain node details.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.ethereum_details
	EthereumDetails *BlockchainNode_EthereumDetailsObservedState `json:"ethereumDetails,omitempty"`

	// Output only. The fully qualified name of the blockchain node.
	//  e.g. `projects/my-project/locations/us-central1/blockchainNodes/my-node`.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.name
	Name *string `json:"name,omitempty"`

	// Output only. The timestamp at which the blockchain node was first created.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp at which the blockchain node was last updated.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The connection information used to interact with a blockchain
	//  node.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.connection_info
	ConnectionInfo *BlockchainNode_ConnectionInfo `json:"connectionInfo,omitempty"`

	// Output only. A status representing the state of the node.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpblockchainnodeengineblockchainnode;gcpblockchainnodeengineblockchainnodes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BlockchainNodeEngineBlockchainNode is the Schema for the BlockchainNodeEngineBlockchainNode API
// +k8s:openapi-gen=true
type BlockchainNodeEngineBlockchainNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BlockchainNodeEngineBlockchainNodeSpec   `json:"spec,omitempty"`
	Status BlockchainNodeEngineBlockchainNodeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BlockchainNodeEngineBlockchainNodeList contains a list of BlockchainNodeEngineBlockchainNode
type BlockchainNodeEngineBlockchainNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BlockchainNodeEngineBlockchainNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BlockchainNodeEngineBlockchainNode{}, &BlockchainNodeEngineBlockchainNodeList{})
}
