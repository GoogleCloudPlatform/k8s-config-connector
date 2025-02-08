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


// +kcc:proto=google.cloud.networkservices.v1.GrpcRoute
type GrpcRoute struct {
	// Required. Name of the GrpcRoute resource. It matches pattern
	//  `projects/*/locations/global/grpcRoutes/<grpc_route_name>`
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.name
	Name *string `json:"name,omitempty"`

	// Optional. Set of label tags associated with the GrpcRoute resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.description
	Description *string `json:"description,omitempty"`

	// Required. Service hostnames with an optional port for which this route
	//  describes traffic.
	//
	//  Format: <hostname>[:<port>]
	//
	//  Hostname is the fully qualified domain name of a network host. This matches
	//  the RFC 1123 definition of a hostname with 2 notable exceptions:
	//   - IPs are not allowed.
	//   - A hostname may be prefixed with a wildcard label (`*.`). The wildcard
	//     label must appear by itself as the first label.
	//
	//  Hostname can be "precise" which is a domain name without the terminating
	//  dot of a network host (e.g. `foo.example.com`) or "wildcard", which is a
	//  domain name prefixed with a single wildcard label (e.g. `*.example.com`).
	//
	//  Note that as per RFC1035 and RFC1123, a label must consist of lower case
	//  alphanumeric characters or '-', and must start and end with an alphanumeric
	//  character. No other punctuation is allowed.
	//
	//  The routes associated with a Mesh or Gateway must have unique hostnames. If
	//  you attempt to attach multiple routes with conflicting hostnames, the
	//  configuration will be rejected.
	//
	//  For example, while it is acceptable for routes for the hostnames
	//  `*.foo.bar.com` and `*.bar.com` to be associated with the same route, it is
	//  not possible to associate two routes both with `*.bar.com` or both with
	//  `bar.com`.
	//
	//  If a port is specified, then gRPC clients must use the channel URI with the
	//  port to match this rule (i.e. "xds:///service:123"), otherwise they must
	//  supply the URI without a port (i.e. "xds:///service").
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.hostnames
	Hostnames []string `json:"hostnames,omitempty"`

	// Optional. Meshes defines a list of meshes this GrpcRoute is attached to, as
	//  one of the routing rules to route the requests served by the mesh.
	//
	//  Each mesh reference should match the pattern:
	//  `projects/*/locations/global/meshes/<mesh_name>`
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.meshes
	Meshes []string `json:"meshes,omitempty"`

	// Optional. Gateways defines a list of gateways this GrpcRoute is attached
	//  to, as one of the routing rules to route the requests served by the
	//  gateway.
	//
	//  Each gateway reference should match the pattern:
	//  `projects/*/locations/global/gateways/<gateway_name>`
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.gateways
	Gateways []string `json:"gateways,omitempty"`

	// Required. A list of detailed rules defining how to route traffic.
	//
	//  Within a single GrpcRoute, the GrpcRoute.RouteAction associated with the
	//  first matching GrpcRoute.RouteRule will be executed. At least one rule
	//  must be supplied.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.rules
	Rules []GrpcRoute_RouteRule `json:"rules,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.GrpcRoute.Destination
type GrpcRoute_Destination struct {
	// Required. The URL of a destination service to which to route traffic.
	//  Must refer to either a BackendService or ServiceDirectoryService.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.Destination.service_name
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
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.Destination.weight
	Weight *int32 `json:"weight,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.GrpcRoute.FaultInjectionPolicy
type GrpcRoute_FaultInjectionPolicy struct {
	// The specification for injecting delay to client requests.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.FaultInjectionPolicy.delay
	Delay *GrpcRoute_FaultInjectionPolicy_Delay `json:"delay,omitempty"`

	// The specification for aborting to client requests.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.FaultInjectionPolicy.abort
	Abort *GrpcRoute_FaultInjectionPolicy_Abort `json:"abort,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.GrpcRoute.FaultInjectionPolicy.Abort
type GrpcRoute_FaultInjectionPolicy_Abort struct {
	// The HTTP status code used to abort the request.
	//
	//  The value must be between 200 and 599 inclusive.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.FaultInjectionPolicy.Abort.http_status
	HTTPStatus *int32 `json:"httpStatus,omitempty"`

	// The percentage of traffic which will be aborted.
	//
	//  The value must be between [0, 100]
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.FaultInjectionPolicy.Abort.percentage
	Percentage *int32 `json:"percentage,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.GrpcRoute.FaultInjectionPolicy.Delay
type GrpcRoute_FaultInjectionPolicy_Delay struct {
	// Specify a fixed delay before forwarding the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.FaultInjectionPolicy.Delay.fixed_delay
	FixedDelay *string `json:"fixedDelay,omitempty"`

	// The percentage of traffic on which delay will be injected.
	//
	//  The value must be between [0, 100]
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.FaultInjectionPolicy.Delay.percentage
	Percentage *int32 `json:"percentage,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.GrpcRoute.HeaderMatch
type GrpcRoute_HeaderMatch struct {
	// Optional. Specifies how to match against the value of the header. If not
	//  specified, a default value of EXACT is used.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.HeaderMatch.type
	Type *string `json:"type,omitempty"`

	// Required. The key of the header.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.HeaderMatch.key
	Key *string `json:"key,omitempty"`

	// Required. The value of the header.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.HeaderMatch.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.GrpcRoute.MethodMatch
type GrpcRoute_MethodMatch struct {
	// Optional. Specifies how to match against the name. If not specified, a
	//  default value of "EXACT" is used.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.MethodMatch.type
	Type *string `json:"type,omitempty"`

	// Required. Name of the service to match against. If unspecified, will
	//  match all services.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.MethodMatch.grpc_service
	GrpcService *string `json:"grpcService,omitempty"`

	// Required. Name of the method to match against. If unspecified, will match
	//  all methods.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.MethodMatch.grpc_method
	GrpcMethod *string `json:"grpcMethod,omitempty"`

	// Optional. Specifies that matches are case sensitive.  The default value
	//  is true. case_sensitive must not be used with a type of
	//  REGULAR_EXPRESSION.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.MethodMatch.case_sensitive
	CaseSensitive *bool `json:"caseSensitive,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.GrpcRoute.RetryPolicy
type GrpcRoute_RetryPolicy struct {
	// - connect-failure: Router will retry on failures connecting to Backend
	//     Services, for example due to connection timeouts.
	//  - refused-stream: Router will retry if the backend service resets the
	//  stream
	//     with a REFUSED_STREAM error code. This reset type indicates that it is
	//     safe to retry.
	//  - cancelled: Router will retry if the gRPC status code in the response
	//  header
	//     is set to cancelled
	//  - deadline-exceeded: Router will retry if the gRPC status code in the
	//  response
	//     header is set to deadline-exceeded
	//  - resource-exhausted: Router will retry if the gRPC status code in the
	//     response header is set to resource-exhausted
	//  - unavailable: Router will retry if the gRPC status code in the response
	//     header is set to unavailable
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.RetryPolicy.retry_conditions
	RetryConditions []string `json:"retryConditions,omitempty"`

	// Specifies the allowed number of retries. This number must be > 0. If not
	//  specified, default to 1.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.RetryPolicy.num_retries
	NumRetries *uint32 `json:"numRetries,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.GrpcRoute.RouteAction
type GrpcRoute_RouteAction struct {
	// Optional. The destination services to which traffic should be forwarded.
	//  If multiple destinations are specified, traffic will be split between
	//  Backend Service(s) according to the weight field of these destinations.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.RouteAction.destinations
	Destinations []GrpcRoute_Destination `json:"destinations,omitempty"`

	// Optional. The specification for fault injection introduced into traffic to test the
	//  resiliency of clients to destination service failure. As part of fault
	//  injection, when clients send requests to a destination, delays can be
	//  introduced on a percentage of requests before sending those requests to
	//  the destination service. Similarly requests from clients can be aborted
	//  by for a percentage of requests.
	//
	//  timeout and retry_policy will be ignored by clients that are configured
	//  with a fault_injection_policy
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.RouteAction.fault_injection_policy
	FaultInjectionPolicy *GrpcRoute_FaultInjectionPolicy `json:"faultInjectionPolicy,omitempty"`

	// Optional. Specifies the timeout for selected route. Timeout is computed
	//  from the time the request has been fully processed (i.e. end of stream)
	//  up until the response has been completely processed. Timeout includes all
	//  retries.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.RouteAction.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Optional. Specifies the retry policy associated with this route.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.RouteAction.retry_policy
	RetryPolicy *GrpcRoute_RetryPolicy `json:"retryPolicy,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.GrpcRoute.RouteMatch
type GrpcRoute_RouteMatch struct {
	// Optional. A gRPC method to match against. If this field is empty or
	//  omitted, will match all methods.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.RouteMatch.method
	Method *GrpcRoute_MethodMatch `json:"method,omitempty"`

	// Optional. Specifies a collection of headers to match.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.RouteMatch.headers
	Headers []GrpcRoute_HeaderMatch `json:"headers,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.GrpcRoute.RouteRule
type GrpcRoute_RouteRule struct {
	// Optional. Matches define conditions used for matching the rule against
	//  incoming gRPC requests. Each match is independent, i.e. this rule will be
	//  matched if ANY one of the matches is satisfied.  If no matches field is
	//  specified, this rule will unconditionally match traffic.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.RouteRule.matches
	Matches []GrpcRoute_RouteMatch `json:"matches,omitempty"`

	// Required. A detailed rule defining how to route traffic. This field is
	//  required.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.RouteRule.action
	Action *GrpcRoute_RouteAction `json:"action,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.GrpcRoute
type GrpcRouteObservedState struct {
	// Output only. Server-defined URL of this resource
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.GrpcRoute.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
