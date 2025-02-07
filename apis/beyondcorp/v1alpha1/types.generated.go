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


// +kcc:proto=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService
type ClientConnectorService struct {
	// Required. Name of resource. The name is ignored during creation.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.name
	Name *string `json:"name,omitempty"`

	// Optional. User-provided name.
	//  The display name should follow certain format.
	//  * Must be 6 to 30 characters in length.
	//  * Can only contain lowercase letters, numbers, and hyphens.
	//  * Must start with a letter.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The details of the ingress settings.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.ingress
	Ingress *ClientConnectorService_Ingress `json:"ingress,omitempty"`

	// Required. The details of the egress settings.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.egress
	Egress *ClientConnectorService_Egress `json:"egress,omitempty"`
}

// +kcc:proto=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Egress
type ClientConnectorService_Egress struct {
	// A VPC from the consumer project.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Egress.peered_vpc
	PeeredVpc *ClientConnectorService_Egress_PeeredVpc `json:"peeredVpc,omitempty"`
}

// +kcc:proto=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Egress.PeeredVpc
type ClientConnectorService_Egress_PeeredVpc struct {
	// Required. The name of the peered VPC owned by the consumer project.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Egress.PeeredVpc.network_vpc
	NetworkVpc *string `json:"networkVpc,omitempty"`
}

// +kcc:proto=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress
type ClientConnectorService_Ingress struct {
	// The basic ingress config for ClientGateways.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.config
	Config *ClientConnectorService_Ingress_Config `json:"config,omitempty"`
}

// +kcc:proto=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config
type ClientConnectorService_Ingress_Config struct {
	// Required. Immutable. The transport protocol used between the client and
	//  the server.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config.transport_protocol
	TransportProtocol *string `json:"transportProtocol,omitempty"`

	// Required. The settings used to configure basic ClientGateways.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config.destination_routes
	DestinationRoutes []ClientConnectorService_Ingress_Config_DestinationRoute `json:"destinationRoutes,omitempty"`
}

// +kcc:proto=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config.DestinationRoute
type ClientConnectorService_Ingress_Config_DestinationRoute struct {
	// Required. The network address of the subnet
	//  for which the packet is routed to the ClientGateway.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config.DestinationRoute.address
	Address *string `json:"address,omitempty"`

	// Required. The network mask of the subnet
	//  for which the packet is routed to the ClientGateway.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.Ingress.Config.DestinationRoute.netmask
	Netmask *string `json:"netmask,omitempty"`
}

// +kcc:proto=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService
type ClientConnectorServiceObservedState struct {
	// Output only. [Output only] Create time stamp.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The operational state of the ClientConnectorService.
	// +kcc:proto:field=google.cloud.beyondcorp.clientconnectorservices.v1.ClientConnectorService.state
	State *string `json:"state,omitempty"`
}
