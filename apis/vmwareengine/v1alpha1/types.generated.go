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

// +kcc:proto=google.cloud.vmwareengine.v1.ExternalAccessRule.IpRange
type ExternalAccessRule_IPRange struct {
	// A single IP address. For example: `10.0.0.5`.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.IpRange.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// An IP address range in the CIDR format. For example: `10.0.0.0/24`.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.IpRange.ip_address_range
	IPAddressRange *string `json:"ipAddressRange,omitempty"`

	// The name of an `ExternalAddress` resource. The external address must
	//  have been reserved in the scope of this external access rule's parent
	//  network policy.  Provide the external address name in the form of
	//  `projects/{project}/locations/{location}/privateClouds/{private_cloud}/externalAddresses/{external_address}`.
	//  For example:
	//  `projects/my-project/locations/us-central1-a/privateClouds/my-cloud/externalAddresses/my-address`.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAccessRule.IpRange.external_address
	ExternalAddress *string `json:"externalAddress,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.NetworkPolicy.NetworkService
type NetworkPolicy_NetworkService struct {
	// True if the service is enabled; false otherwise.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.NetworkService.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.VmwareEngineNetwork.VpcNetwork
type VmwareEngineNetwork_VpcNetwork struct {
}

// +kcc:proto=google.cloud.vmwareengine.v1.NetworkPolicy.NetworkService
type NetworkPolicy_NetworkServiceObservedState struct {
	// Output only. State of the service. New values may be added to this enum
	//  when appropriate.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.NetworkService.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.VmwareEngineNetwork.VpcNetwork
type VmwareEngineNetwork_VpcNetworkObservedState struct {
	// Output only. Type of VPC network (INTRANET, INTERNET, or
	//  GOOGLE_CLOUD)
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.VpcNetwork.type
	Type *string `json:"type,omitempty"`

	// Output only. The relative resource name of the service VPC network this
	//  VMware Engine network is attached to. For example:
	//  `projects/123123/global/networks/my-network`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.VpcNetwork.network
	Network *string `json:"network,omitempty"`
}
