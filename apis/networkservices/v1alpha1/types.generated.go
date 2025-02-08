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


// +kcc:proto=google.cloud.networkservices.v1.TcpRoute
type TcpRoute struct {
	// Required. Name of the TcpRoute resource. It matches pattern
	//  `projects/*/locations/global/tcpRoutes/tcp_route_name>`.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.name
	Name *string `json:"name,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.description
	Description *string `json:"description,omitempty"`

	// Required. Rules that define how traffic is routed and handled. At least one
	//  RouteRule must be supplied. If there are multiple rules then the action
	//  taken will be the first rule to match.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.rules
	Rules []TcpRoute_RouteRule `json:"rules,omitempty"`

	// Optional. Meshes defines a list of meshes this TcpRoute is attached to, as
	//  one of the routing rules to route the requests served by the mesh.
	//
	//  Each mesh reference should match the pattern:
	//  `projects/*/locations/global/meshes/<mesh_name>`
	//
	//  The attached Mesh should be of a type SIDECAR
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.meshes
	Meshes []string `json:"meshes,omitempty"`

	// Optional. Gateways defines a list of gateways this TcpRoute is attached to,
	//  as one of the routing rules to route the requests served by the gateway.
	//
	//  Each gateway reference should match the pattern:
	//  `projects/*/locations/global/gateways/<gateway_name>`
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.gateways
	Gateways []string `json:"gateways,omitempty"`

	// Optional. Set of label tags associated with the TcpRoute resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.TcpRoute.RouteAction
type TcpRoute_RouteAction struct {
	// Optional. The destination services to which traffic should be forwarded.
	//  At least one destination service is required. Only one of route
	//  destination or original destination can be set.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.RouteAction.destinations
	Destinations []TcpRoute_RouteDestination `json:"destinations,omitempty"`

	// Optional. If true, Router will use the destination IP and port of the
	//  original connection as the destination of the request. Default is false.
	//  Only one of route destinations or original destination can be set.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.RouteAction.original_destination
	OriginalDestination *bool `json:"originalDestination,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.TcpRoute.RouteDestination
type TcpRoute_RouteDestination struct {
	// Required. The URL of a BackendService to route traffic to.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.RouteDestination.service_name
	ServiceName *string `json:"serviceName,omitempty"`

	// Optional. Specifies the proportion of requests forwarded to the backend
	//  referenced by the serviceName field. This is computed as:
	//  - weight/Sum(weights in this destination list).
	//  For non-zero values, there may be some epsilon from the exact proportion
	//  defined here depending on the precision an implementation supports.
	//
	//  If only one serviceName is specified and it has a weight greater than 0,
	//  100% of the traffic is forwarded to that backend.
	//
	//  If weights are specified for any one service name, they need to be
	//  specified for all of them.
	//
	//  If weights are unspecified for all services, then, traffic is distributed
	//  in equal proportions to all of them.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.RouteDestination.weight
	Weight *int32 `json:"weight,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.TcpRoute.RouteMatch
type TcpRoute_RouteMatch struct {
	// Required. Must be specified in the CIDR range format. A CIDR range
	//  consists of an IP Address and a prefix length to construct the subnet
	//  mask. By default, the prefix length is 32 (i.e. matches a single IP
	//  address). Only IPV4 addresses are supported.
	//  Examples:
	//  "10.0.0.1" - matches against this exact IP address.
	//  "10.0.0.0/8" - matches against any IP address within the 10.0.0.0 subnet
	//  and 255.255.255.0 mask.
	//  "0.0.0.0/0" - matches against any IP address'.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.RouteMatch.address
	Address *string `json:"address,omitempty"`

	// Required. Specifies the destination port to match against.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.RouteMatch.port
	Port *string `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.TcpRoute.RouteRule
type TcpRoute_RouteRule struct {
	// Optional. RouteMatch defines the predicate used to match requests to a
	//  given action. Multiple match types are "OR"ed for evaluation. If no
	//  routeMatch field is specified, this rule will unconditionally match
	//  traffic.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.RouteRule.matches
	Matches []TcpRoute_RouteMatch `json:"matches,omitempty"`

	// Required. The detailed rule defining how to route matched traffic.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.RouteRule.action
	Action *TcpRoute_RouteAction `json:"action,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.TcpRoute
type TcpRouteObservedState struct {
	// Output only. Server-defined URL of this resource
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.TcpRoute.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
