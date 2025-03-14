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

// +kcc:proto=google.cloud.vmwareengine.v1.VmwareEngineNetwork
type VmwareEngineNetwork struct {

	// User-provided description for this VMware Engine network.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.description
	Description *string `json:"description,omitempty"`

	// Required. VMware Engine network type.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.type
	Type *string `json:"type,omitempty"`

	// Checksum that may be sent on update and delete requests to ensure that the
	//  user-provided value is up to date before the server processes a request.
	//  The server computes checksums based on the value of other fields in the
	//  request.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.VmwareEngineNetwork.VpcNetwork
type VmwareEngineNetwork_VpcNetwork struct {
}

// +kcc:proto=google.cloud.vmwareengine.v1.VmwareEngineNetwork
type VmwareEngineNetworkObservedState struct {
	// Output only. The resource name of the VMware Engine network.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/global/vmwareEngineNetworks/my-network`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. VMware Engine service VPC networks that provide connectivity
	//  from a private cloud to customer projects, the internet, and other Google
	//  Cloud services.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.vpc_networks
	VpcNetworks []VmwareEngineNetwork_VpcNetwork `json:"vpcNetworks,omitempty"`

	// Output only. State of the VMware Engine network.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.state
	State *string `json:"state,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.uid
	Uid *string `json:"uid,omitempty"`
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
