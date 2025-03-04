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

// +kcc:proto=google.cloud.tpu.v2.NetworkEndpoint
type NetworkEndpoint struct {
	// The internal IP address of this network endpoint.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkEndpoint.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// The port of this network endpoint.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkEndpoint.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.NetworkEndpoint
type NetworkEndpointObservedState struct {
	// The access config for the TPU worker.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkEndpoint.access_config
	AccessConfig *AccessConfigObservedState `json:"accessConfig,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.AccessConfig
type AccessConfigObservedState struct {

	// Output only. An external IP address associated with the TPU worker.
	ExternalIP *string `json:"externalIP,omitempty"`
}
