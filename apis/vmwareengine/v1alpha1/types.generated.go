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


// +kcc:proto=google.cloud.vmwareengine.v1.PrivateConnection
type PrivateConnection struct {

	// Optional. User-provided description for this private connection.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.description
	Description *string `json:"description,omitempty"`

	// Required. The relative resource name of Legacy VMware Engine network.
	//  Specify the name in the following form:
	//  `projects/{project}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`
	//  where `{project}`, `{location}` will be same as specified in private
	//  connection resource name and `{vmware_engine_network_id}` will be in the
	//  form of `{location}`-default e.g.
	//  projects/project/locations/us-central1/vmwareEngineNetworks/us-central1-default.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.vmware_engine_network
	VmwareEngineNetwork *string `json:"vmwareEngineNetwork,omitempty"`

	// Required. Private connection type.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.type
	Type *string `json:"type,omitempty"`

	// Optional. Routing Mode.
	//  Default value is set to GLOBAL.
	//  For type = PRIVATE_SERVICE_ACCESS, this field can be set to GLOBAL or
	//  REGIONAL, for other types only GLOBAL is supported.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.routing_mode
	RoutingMode *string `json:"routingMode,omitempty"`

	// Required. Service network to create private connection.
	//  Specify the name in the following form:
	//  `projects/{project}/global/networks/{network_id}`
	//  For type = PRIVATE_SERVICE_ACCESS, this field represents servicenetworking
	//  VPC, e.g. projects/project-tp/global/networks/servicenetworking.
	//  For type = NETAPP_CLOUD_VOLUME, this field represents NetApp service VPC,
	//  e.g. projects/project-tp/global/networks/netapp-tenant-vpc.
	//  For type = DELL_POWERSCALE, this field represent Dell service VPC, e.g.
	//  projects/project-tp/global/networks/dell-tenant-vpc.
	//  For type= THIRD_PARTY_SERVICE, this field could represent a consumer VPC or
	//  any other producer VPC to which the VMware Engine Network needs to be
	//  connected, e.g. projects/project/global/networks/vpc.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.service_network
	ServiceNetwork *string `json:"serviceNetwork,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.PrivateConnection
type PrivateConnectionObservedState struct {
	// Output only. The resource name of the private connection.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1/privateConnections/my-connection`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the private connection.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.state
	State *string `json:"state,omitempty"`

	// Output only. The canonical name of the VMware Engine network in the form:
	//  `projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.vmware_engine_network_canonical
	VmwareEngineNetworkCanonical *string `json:"vmwareEngineNetworkCanonical,omitempty"`

	// Output only. VPC network peering id between given network VPC and
	//  VMwareEngineNetwork.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.peering_id
	PeeringID *string `json:"peeringID,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Peering state between service network and VMware Engine
	//  network.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.peering_state
	PeeringState *string `json:"peeringState,omitempty"`
}
