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


// +kcc:proto=google.cloud.networkservices.v1.TlsRoute
type TlsRoute struct {
	// Required. Name of the TlsRoute resource. It matches pattern
	//  `projects/*/locations/global/tlsRoutes/tls_route_name>`.
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.name
	Name *string `json:"name,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.description
	Description *string `json:"description,omitempty"`

	// Required. Rules that define how traffic is routed and handled. At least one
	//  RouteRule must be supplied. If there are multiple rules then the action
	//  taken will be the first rule to match.
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.rules
	Rules []TlsRoute_RouteRule `json:"rules,omitempty"`

	// Optional. Meshes defines a list of meshes this TlsRoute is attached to, as
	//  one of the routing rules to route the requests served by the mesh.
	//
	//  Each mesh reference should match the pattern:
	//  `projects/*/locations/global/meshes/<mesh_name>`
	//
	//  The attached Mesh should be of a type SIDECAR
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.meshes
	Meshes []string `json:"meshes,omitempty"`

	// Optional. Gateways defines a list of gateways this TlsRoute is attached to,
	//  as one of the routing rules to route the requests served by the gateway.
	//
	//  Each gateway reference should match the pattern:
	//  `projects/*/locations/global/gateways/<gateway_name>`
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.gateways
	Gateways []string `json:"gateways,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.TlsRoute.RouteAction
type TlsRoute_RouteAction struct {
	// Required. The destination services to which traffic should be forwarded.
	//  At least one destination service is required.
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.RouteAction.destinations
	Destinations []TlsRoute_RouteDestination `json:"destinations,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.TlsRoute.RouteDestination
type TlsRoute_RouteDestination struct {
	// Required. The URL of a BackendService to route traffic to.
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.RouteDestination.service_name
	ServiceName *string `json:"serviceName,omitempty"`

	// Optional. Specifies the proportion of requests forwareded to the backend
	//  referenced by the service_name field. This is computed as:
	//  - weight/Sum(weights in destinations)
	//  Weights in all destinations does not need to sum up to 100.
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.RouteDestination.weight
	Weight *int32 `json:"weight,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.TlsRoute.RouteMatch
type TlsRoute_RouteMatch struct {
	// Optional. SNI (server name indicator) to match against.
	//  SNI will be matched against all wildcard domains, i.e. `www.example.com`
	//  will be first matched against `www.example.com`, then `*.example.com`,
	//  then `*.com.`
	//  Partial wildcards are not supported, and values like *w.example.com are
	//  invalid.
	//  At least one of sni_host and alpn is required.
	//  Up to 5 sni hosts across all matches can be set.
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.RouteMatch.sni_host
	SniHost []string `json:"sniHost,omitempty"`

	// Optional. ALPN (Application-Layer Protocol Negotiation) to match against.
	//  Examples: "http/1.1", "h2".
	//  At least one of sni_host and alpn is required.
	//  Up to 5 alpns across all matches can be set.
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.RouteMatch.alpn
	Alpn []string `json:"alpn,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.TlsRoute.RouteRule
type TlsRoute_RouteRule struct {
	// Required. RouteMatch defines the predicate used to match requests to a
	//  given action. Multiple match types are "OR"ed for evaluation.
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.RouteRule.matches
	Matches []TlsRoute_RouteMatch `json:"matches,omitempty"`

	// Required. The detailed rule defining how to route matched traffic.
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.RouteRule.action
	Action *TlsRoute_RouteAction `json:"action,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.TlsRoute
type TlsRouteObservedState struct {
	// Output only. Server-defined URL of this resource
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.TlsRoute.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
