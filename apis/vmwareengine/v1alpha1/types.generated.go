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

// +generated:types
// krm.group: vmwareengine.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.vmwareengine.v1
// resource: VMwareEngineNetwork:VmwareEngineNetwork
// resource: VMwareEngineNetworkPeering:NetworkPeering
// resource: VMwareEngineNetworkPolicy:NetworkPolicy
// resource: VMwareEngineExternalAccessRule:ExternalAccessRule
// resource: VMwareEnginePrivateCloud:PrivateCloud
// resource: VMwareEngineExternalAddress:ExternalAddress

package v1alpha1

// +kcc:proto=google.cloud.vmwareengine.v1.ExternalAddress
type ExternalAddress struct {

	// The internal IP address of a workload VM.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.internal_ip
	InternalIP *string `json:"internalIP,omitempty"`

	// User-provided description for this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.description
	Description *string `json:"description,omitempty"`
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

// +kcc:proto=google.cloud.vmwareengine.v1.ExternalAddress
type ExternalAddressObservedState struct {
	// Output only. The resource name of this external IP address.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1-a/privateClouds/my-cloud/externalAddresses/my-address`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The external IP address of a workload VM.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.external_ip
	ExternalIP *string `json:"externalIP,omitempty"`

	// Output only. The state of the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.state
	State *string `json:"state,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.uid
	Uid *string `json:"uid,omitempty"`
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
