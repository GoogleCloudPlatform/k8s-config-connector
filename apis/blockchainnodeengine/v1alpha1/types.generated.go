// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode
type BlockchainNode struct {
	// Ethereum-specific blockchain node details.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.ethereum_details
	EthereumDetails *BlockchainNode_EthereumDetails `json:"ethereumDetails,omitempty"`

	// User-provided key-value pairs.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The blockchain type of the node.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.blockchain_type
	BlockchainType *string `json:"blockchainType,omitempty"`

	// Optional. When true, the node is only accessible via Private Service
	//  Connect; no public endpoints are exposed. Otherwise, the node is only
	//  accessible via public endpoints. Warning: Private Service Connect enabled
	//  nodes may require a manual migration effort to remain compatible with
	//  future versions of the product. If this feature is enabled, you will be
	//  notified of these changes along with any required action to avoid
	//  disruption. See https://cloud.google.com/vpc/docs/private-service-connect.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.private_service_connect_enabled
	PrivateServiceConnectEnabled *bool `json:"privateServiceConnectEnabled,omitempty"`
}

// +kcc:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode.ConnectionInfo
type BlockchainNode_ConnectionInfo struct {
}

// +kcc:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode.ConnectionInfo.EndpointInfo
type BlockchainNode_ConnectionInfo_EndpointInfo struct {
}

// +kcc:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails
type BlockchainNode_EthereumDetails struct {
	// Details for the Geth execution client.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.geth_details
	GethDetails *BlockchainNode_EthereumDetails_GethDetails `json:"gethDetails,omitempty"`

	// Immutable. The Ethereum environment being accessed.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.network
	Network *string `json:"network,omitempty"`

	// Immutable. The type of Ethereum node.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.node_type
	NodeType *string `json:"nodeType,omitempty"`

	// Immutable. The execution client
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.execution_client
	ExecutionClient *string `json:"executionClient,omitempty"`

	// Immutable. The consensus client.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.consensus_client
	ConsensusClient *string `json:"consensusClient,omitempty"`

	// Immutable. Enables JSON-RPC access to functions in the `admin` namespace.
	//  Defaults to `false`.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.api_enable_admin
	ApiEnableAdmin *bool `json:"apiEnableAdmin,omitempty"`

	// Immutable. Enables JSON-RPC access to functions in the `debug` namespace.
	//  Defaults to `false`.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.api_enable_debug
	ApiEnableDebug *bool `json:"apiEnableDebug,omitempty"`

	// Configuration for validator-related parameters on the beacon client,
	//  and for any GCP-managed validator client.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.validator_config
	ValidatorConfig *BlockchainNode_EthereumDetails_ValidatorConfig `json:"validatorConfig,omitempty"`
}

// +kcc:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.EthereumEndpoints
type BlockchainNode_EthereumDetails_EthereumEndpoints struct {
}

// +kcc:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.GethDetails
type BlockchainNode_EthereumDetails_GethDetails struct {
	// Immutable. Blockchain garbage collection mode.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.GethDetails.garbage_collection_mode
	GarbageCollectionMode *string `json:"garbageCollectionMode,omitempty"`
}

// +kcc:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.ValidatorConfig
type BlockchainNode_EthereumDetails_ValidatorConfig struct {
	// URLs for MEV-relay services to use for block building. When set, a
	//  GCP-managed MEV-boost service is configured on the beacon client.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.ValidatorConfig.mev_relay_urls
	MevRelayUrls []string `json:"mevRelayUrls,omitempty"`

	// Immutable. When true, deploys a GCP-managed validator client alongside
	//  the beacon client.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.ValidatorConfig.managed_validator_client
	ManagedValidatorClient *bool `json:"managedValidatorClient,omitempty"`

	// An Ethereum address which the beacon client will send fee rewards to if
	//  no recipient is configured in the validator client.
	//
	//  See https://lighthouse-book.sigmaprime.io/suggested-fee-recipient.html
	//  or https://docs.prylabs.network/docs/execution-node/fee-recipient for
	//  examples of how this is used.
	//
	//  Note that while this is often described as "suggested", as we run the
	//  execution node we can trust the execution node, and therefore this is
	//  considered enforced.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.ValidatorConfig.beacon_fee_recipient
	BeaconFeeRecipient *string `json:"beaconFeeRecipient,omitempty"`
}

// +kcc:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode
type BlockchainNodeObservedState struct {
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

// +kcc:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode.ConnectionInfo
type BlockchainNode_ConnectionInfoObservedState struct {
	// Output only. The endpoint information through which to interact with a
	//  blockchain node.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.ConnectionInfo.endpoint_info
	EndpointInfo *BlockchainNode_ConnectionInfo_EndpointInfo `json:"endpointInfo,omitempty"`

	// Output only. A service attachment that exposes a node, and has the
	//  following format:
	//  projects/{project}/regions/{region}/serviceAttachments/{service_attachment_name}
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.ConnectionInfo.service_attachment
	ServiceAttachment *string `json:"serviceAttachment,omitempty"`
}

// +kcc:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode.ConnectionInfo.EndpointInfo
type BlockchainNode_ConnectionInfo_EndpointInfoObservedState struct {
	// Output only. The assigned URL for the node JSON-RPC API endpoint.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.ConnectionInfo.EndpointInfo.json_rpc_api_endpoint
	JsonRpcApiEndpoint *string `json:"jsonRpcApiEndpoint,omitempty"`

	// Output only. The assigned URL for the node WebSockets API endpoint.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.ConnectionInfo.EndpointInfo.websockets_api_endpoint
	WebsocketsApiEndpoint *string `json:"websocketsApiEndpoint,omitempty"`
}

// +kcc:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails
type BlockchainNode_EthereumDetailsObservedState struct {
	// Output only. Ethereum-specific endpoint information.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.additional_endpoints
	AdditionalEndpoints *BlockchainNode_EthereumDetails_EthereumEndpoints `json:"additionalEndpoints,omitempty"`
}

// +kcc:proto=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.EthereumEndpoints
type BlockchainNode_EthereumDetails_EthereumEndpointsObservedState struct {
	// Output only. The assigned URL for the node's Beacon API endpoint.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.EthereumEndpoints.beacon_api_endpoint
	BeaconApiEndpoint *string `json:"beaconApiEndpoint,omitempty"`

	// Output only. The assigned URL for the node's Beacon Prometheus metrics
	//  endpoint. See [Prometheus
	//  Metrics](https://lighthouse-book.sigmaprime.io/advanced_metrics.html)
	//  for more details.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.EthereumEndpoints.beacon_prometheus_metrics_api_endpoint
	BeaconPrometheusMetricsApiEndpoint *string `json:"beaconPrometheusMetricsApiEndpoint,omitempty"`

	// Output only. The assigned URL for the node's execution client's
	//  Prometheus metrics endpoint.
	// +kcc:proto:field=google.cloud.blockchainnodeengine.v1.BlockchainNode.EthereumDetails.EthereumEndpoints.execution_client_prometheus_metrics_api_endpoint
	ExecutionClientPrometheusMetricsApiEndpoint *string `json:"executionClientPrometheusMetricsApiEndpoint,omitempty"`
}
