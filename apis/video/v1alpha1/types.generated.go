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


// +kcc:proto=google.cloud.video.livestream.v1.Pool
type Pool struct {
	// The resource name of the pool, in the form of:
	//  `projects/{project}/locations/{location}/pools/{poolId}`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Pool.name
	Name *string `json:"name,omitempty"`

	// User-defined key/value metadata.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Pool.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Network configuration for the pool.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Pool.network_config
	NetworkConfig *Pool_NetworkConfig `json:"networkConfig,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Pool.NetworkConfig
type Pool_NetworkConfig struct {
	// peered_network is the network resource URL of the network that is peered
	//  to the service provider network. Must be of the format
	//  projects/NETWORK_PROJECT_NUMBER/global/networks/NETWORK_NAME, where
	//  NETWORK_PROJECT_NUMBER is the project number of the Cloud project that
	//  holds your VPC network and NETWORK_NAME is the name of your VPC network.
	//  If peered_network is omitted or empty, the pool will use endpoints that
	//  are publicly available.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Pool.NetworkConfig.peered_network
	PeeredNetwork *string `json:"peeredNetwork,omitempty"`
}

// +kcc:proto=google.cloud.video.livestream.v1.Pool
type PoolObservedState struct {
	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Pool.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Pool.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
