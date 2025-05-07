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

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

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

// +kcc:proto=google.cloud.tpu.v2.NetworkConfig
type NetworkConfig struct {
	// The network for the TPU node. It must be a preexisting Google
	//  Compute Engine network. If none is provided, "default" will be used.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkConfig.network
	NetworkRef *refs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// The subnetwork for the TPU node. It must be a preexisting
	//  Google Compute Engine subnetwork. If none is provided, "default" will be
	//  used.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkConfig.subnetwork
	SubnetworkRef *refs.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`

	// Indicates that external IP addresses would be associated with the TPU
	//  workers. If set to false, the specified subnetwork or network should have
	//  Private Google Access enabled.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkConfig.enable_external_ips
	EnableExternalIps *bool `json:"enableExternalIPs,omitempty"`

	// Allows the TPU node to send and receive packets with non-matching
	//  destination or source IPs. This is required if you plan to use the TPU
	//  workers to forward routes.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkConfig.can_ip_forward
	CanIPForward *bool `json:"canIPForward,omitempty"`

	// Optional. Specifies networking queue count for TPU VM instance's network
	//  interface.
	// +kcc:proto:field=google.cloud.tpu.v2.NetworkConfig.queue_count
	QueueCount *int32 `json:"queueCount,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.ServiceAccount
type ServiceAccount struct {
	// The service account to be used. If empty, the default Compute service
	//  account will be used.
	// +kcc:proto:field=google.cloud.tpu.v2.ServiceAccount.email
	ServiceAccountRef *refs.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// The list of scopes to be made available for this service account. If empty,
	//  access to all Cloud APIs will be allowed.
	// +kcc:proto:field=google.cloud.tpu.v2.ServiceAccount.scope
	Scope []string `json:"scope,omitempty"`
}
