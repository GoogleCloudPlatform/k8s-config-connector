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

// API sources for BlockchainNodeEngineBlockchainNode, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/blockchainnodeengine/v1/blockchainnodeengine.proto
// +kcc:source:service-docs=https://cloud.google.com/blockchain-node-engine
// +kcc:source:resource-docs=https://docs.cloud.google.com/blockchain-node-engine/docs/reference/rest/v1/projects.locations.blockchainNodes

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BlockchainNodeEngineBlockchainNodeGVK = GroupVersion.WithKind("BlockchainNodeEngineBlockchainNode")

// BlockchainNodeEngineBlockchainNodeSpec defines the desired state of BlockchainNodeEngineBlockchainNode
// +kcc:spec:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode
// +kcc:required-from-proto
type BlockchainNodeEngineBlockchainNodeSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The BlockchainNodeEngineBlockchainNode name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Ethereum-specific blockchain node details.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.ethereum_details
	EthereumDetails *BlockchainNode_EthereumDetails `json:"ethereumDetails,omitempty"`

	// User-provided key-value pairs.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The blockchain type of the node.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.blockchain_type
	BlockchainType *string `json:"blockchainType,omitempty"`
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

	// Output only. The timestamp at which the blockchain node was first created.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp at which the blockchain node was last updated.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The connection information used to interact with a blockchain
	//  node.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.connection_info
	ConnectionInfo *BlockchainNode_ConnectionInfoObservedState `json:"connectionInfo,omitempty"`

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
