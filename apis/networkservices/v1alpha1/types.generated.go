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


// +kcc:proto=google.cloud.networkservices.v1.HttpRoute
type HttpRoute struct {
	// Required. Name of the HttpRoute resource. It matches pattern
	//  `projects/*/locations/global/httpRoutes/http_route_name>`.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.name
	Name *string `json:"name,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.description
	Description *string `json:"description,omitempty"`

	// Required. Hostnames define a set of hosts that should match against the
	//  HTTP host header to select a HttpRoute to process the request. Hostname is
	//  the fully qualified domain name of a network host, as defined by RFC 1123
	//  with the exception that:
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
	//  The routes associated with a Mesh or Gateways  must have unique hostnames.
	//  If you attempt to attach multiple routes with conflicting hostnames,
	//  the configuration will be rejected.
	//
	//  For example, while it is acceptable for routes for the hostnames
	//  `*.foo.bar.com` and `*.bar.com` to be associated with the same Mesh (or
	//  Gateways under the same scope), it is not possible to associate two routes
	//  both with `*.bar.com` or both with `bar.com`.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.hostnames
	Hostnames []string `json:"hostnames,omitempty"`

	// Optional. Meshes defines a list of meshes this HttpRoute is attached to, as
	//  one of the routing rules to route the requests served by the mesh.
	//
	//  Each mesh reference should match the pattern:
	//  `projects/*/locations/global/meshes/<mesh_name>`
	//
	//  The attached Mesh should be of a type SIDECAR
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.meshes
	Meshes []string `json:"meshes,omitempty"`

	// Optional. Gateways defines a list of gateways this HttpRoute is attached
	//  to, as one of the routing rules to route the requests served by the
	//  gateway.
	//
	//  Each gateway reference should match the pattern:
	//  `projects/*/locations/global/gateways/<gateway_name>`
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.gateways
	Gateways []string `json:"gateways,omitempty"`

	// Optional. Set of label tags associated with the HttpRoute resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Rules that define how traffic is routed and handled.
	//  Rules will be matched sequentially based on the RouteMatch specified for
	//  the rule.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.rules
	Rules []HttpRoute_RouteRule `json:"rules,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.CorsPolicy
type HttpRoute_CorsPolicy struct {
	// Specifies the list of origins that will be allowed to do CORS requests.
	//  An origin is allowed if it matches either an item in allow_origins or
	//  an item in allow_origin_regexes.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.CorsPolicy.allow_origins
	AllowOrigins []string `json:"allowOrigins,omitempty"`

	// Specifies the regular expression patterns that match allowed origins. For
	//  regular expression grammar, please see
	//  https://github.com/google/re2/wiki/Syntax.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.CorsPolicy.allow_origin_regexes
	AllowOriginRegexes []string `json:"allowOriginRegexes,omitempty"`

	// Specifies the content for Access-Control-Allow-Methods header.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.CorsPolicy.allow_methods
	AllowMethods []string `json:"allowMethods,omitempty"`

	// Specifies the content for Access-Control-Allow-Headers header.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.CorsPolicy.allow_headers
	AllowHeaders []string `json:"allowHeaders,omitempty"`

	// Specifies the content for Access-Control-Expose-Headers header.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.CorsPolicy.expose_headers
	ExposeHeaders []string `json:"exposeHeaders,omitempty"`

	// Specifies how long result of a preflight request can be cached in
	//  seconds. This translates to the Access-Control-Max-Age header.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.CorsPolicy.max_age
	MaxAge *string `json:"maxAge,omitempty"`

	// In response to a preflight request, setting this to true indicates that
	//  the actual request can include user credentials. This translates to the
	//  Access-Control-Allow-Credentials header.
	//
	//  Default value is false.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.CorsPolicy.allow_credentials
	AllowCredentials *bool `json:"allowCredentials,omitempty"`

	// If true, the CORS policy is disabled. The default value is false, which
	//  indicates that the CORS policy is in effect.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.CorsPolicy.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.Destination
type HttpRoute_Destination struct {
	// The URL of a BackendService to route traffic to.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.Destination.service_name
	ServiceName *string `json:"serviceName,omitempty"`

	// Specifies the proportion of requests forwarded to the backend referenced
	//  by the serviceName field. This is computed as:
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
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.Destination.weight
	Weight *int32 `json:"weight,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.FaultInjectionPolicy
type HttpRoute_FaultInjectionPolicy struct {
	// The specification for injecting delay to client requests.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.FaultInjectionPolicy.delay
	Delay *HttpRoute_FaultInjectionPolicy_Delay `json:"delay,omitempty"`

	// The specification for aborting to client requests.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.FaultInjectionPolicy.abort
	Abort *HttpRoute_FaultInjectionPolicy_Abort `json:"abort,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.FaultInjectionPolicy.Abort
type HttpRoute_FaultInjectionPolicy_Abort struct {
	// The HTTP status code used to abort the request.
	//
	//  The value must be between 200 and 599 inclusive.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.FaultInjectionPolicy.Abort.http_status
	HTTPStatus *int32 `json:"httpStatus,omitempty"`

	// The percentage of traffic which will be aborted.
	//
	//  The value must be between [0, 100]
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.FaultInjectionPolicy.Abort.percentage
	Percentage *int32 `json:"percentage,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.FaultInjectionPolicy.Delay
type HttpRoute_FaultInjectionPolicy_Delay struct {
	// Specify a fixed delay before forwarding the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.FaultInjectionPolicy.Delay.fixed_delay
	FixedDelay *string `json:"fixedDelay,omitempty"`

	// The percentage of traffic on which delay will be injected.
	//
	//  The value must be between [0, 100]
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.FaultInjectionPolicy.Delay.percentage
	Percentage *int32 `json:"percentage,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.HeaderMatch
type HttpRoute_HeaderMatch struct {
	// The value of the header should match exactly the content of
	//  exact_match.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.HeaderMatch.exact_match
	ExactMatch *string `json:"exactMatch,omitempty"`

	// The value of the header must match the regular expression specified in
	//  regex_match. For regular expression grammar, please see:
	//  https://github.com/google/re2/wiki/Syntax
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.HeaderMatch.regex_match
	RegexMatch *string `json:"regexMatch,omitempty"`

	// The value of the header must start with the contents of prefix_match.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.HeaderMatch.prefix_match
	PrefixMatch *string `json:"prefixMatch,omitempty"`

	// A header with header_name must exist. The match takes place whether or
	//  not the header has a value.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.HeaderMatch.present_match
	PresentMatch *bool `json:"presentMatch,omitempty"`

	// The value of the header must end with the contents of suffix_match.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.HeaderMatch.suffix_match
	SuffixMatch *string `json:"suffixMatch,omitempty"`

	// If specified, the rule will match if the request header value is within
	//  the range.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.HeaderMatch.range_match
	RangeMatch *HttpRoute_HeaderMatch_IntegerRange `json:"rangeMatch,omitempty"`

	// The name of the HTTP header to match against.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.HeaderMatch.header
	Header *string `json:"header,omitempty"`

	// If specified, the match result will be inverted before checking. Default
	//  value is set to false.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.HeaderMatch.invert_match
	InvertMatch *bool `json:"invertMatch,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.HeaderMatch.IntegerRange
type HttpRoute_HeaderMatch_IntegerRange struct {
	// Start of the range (inclusive)
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.HeaderMatch.IntegerRange.start
	Start *int32 `json:"start,omitempty"`

	// End of the range (exclusive)
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.HeaderMatch.IntegerRange.end
	End *int32 `json:"end,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.HeaderModifier
type HttpRoute_HeaderModifier struct {
	// Completely overwrite/replace the headers with given map where key is the
	//  name of the header, value is the value of the header.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.HeaderModifier.set
	Set map[string]string `json:"set,omitempty"`

	// Add the headers with given map where key is the name of the header, value
	//  is the value of the header.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.HeaderModifier.add
	Add map[string]string `json:"add,omitempty"`

	// Remove headers (matching by header names) specified in the list.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.HeaderModifier.remove
	Remove []string `json:"remove,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.QueryParameterMatch
type HttpRoute_QueryParameterMatch struct {
	// The value of the query parameter must exactly match the contents of
	//  exact_match.
	//
	//  Only one of exact_match, regex_match, or present_match must be set.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.QueryParameterMatch.exact_match
	ExactMatch *string `json:"exactMatch,omitempty"`

	// The value of the query parameter must match the regular expression
	//  specified by regex_match. For regular expression grammar, please see
	//  https://github.com/google/re2/wiki/Syntax
	//
	//  Only one of exact_match, regex_match, or present_match must be set.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.QueryParameterMatch.regex_match
	RegexMatch *string `json:"regexMatch,omitempty"`

	// Specifies that the QueryParameterMatcher matches if request contains
	//  query parameter, irrespective of whether the parameter has a value or
	//  not.
	//
	//  Only one of exact_match, regex_match, or present_match must be set.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.QueryParameterMatch.present_match
	PresentMatch *bool `json:"presentMatch,omitempty"`

	// The name of the query parameter to match.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.QueryParameterMatch.query_parameter
	QueryParameter *string `json:"queryParameter,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.Redirect
type HttpRoute_Redirect struct {
	// The host that will be used in the redirect response instead of the one
	//  that was supplied in the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.Redirect.host_redirect
	HostRedirect *string `json:"hostRedirect,omitempty"`

	// The path that will be used in the redirect response instead of the one
	//  that was supplied in the request.
	//  path_redirect can not be supplied together with prefix_redirect. Supply
	//  one alone or neither. If neither is supplied, the path of the original
	//  request will be used for the redirect.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.Redirect.path_redirect
	PathRedirect *string `json:"pathRedirect,omitempty"`

	// Indicates that during redirection, the matched prefix (or path) should be
	//  swapped with this value. This option allows URLs be dynamically created
	//  based on the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.Redirect.prefix_rewrite
	PrefixRewrite *string `json:"prefixRewrite,omitempty"`

	// The HTTP Status code to use for the redirect.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.Redirect.response_code
	ResponseCode *string `json:"responseCode,omitempty"`

	// If set to true, the URL scheme in the redirected request is set to https.
	//  If set to false, the URL scheme of the redirected request will remain the
	//  same as that of the request.
	//
	//  The default is set to false.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.Redirect.https_redirect
	HTTPSRedirect *bool `json:"httpsRedirect,omitempty"`

	// if set to true, any accompanying query portion of the original URL is
	//  removed prior to redirecting the request. If set to false, the query
	//  portion of the original URL is retained.
	//
	//  The default is set to false.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.Redirect.strip_query
	StripQuery *bool `json:"stripQuery,omitempty"`

	// The port that will be used in the redirected request instead of the one
	//  that was supplied in the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.Redirect.port_redirect
	PortRedirect *int32 `json:"portRedirect,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.RequestMirrorPolicy
type HttpRoute_RequestMirrorPolicy struct {
	// The destination the requests will be mirrored to. The weight of the
	//  destination will be ignored.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RequestMirrorPolicy.destination
	Destination *HttpRoute_Destination `json:"destination,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.RetryPolicy
type HttpRoute_RetryPolicy struct {
	// Specifies one or more conditions when this retry policy applies. Valid
	//  values are:
	//    5xx: Proxy will attempt a retry if the destination service responds
	//      with any 5xx response code, of if the destination service does not
	//      respond at all, example: disconnect, reset, read timeout, connection
	//      failure and refused streams.
	//
	//    gateway-error: Similar to 5xx, but only applies to response codes 502,
	//      503, 504.
	//
	//    reset: Proxy will attempt a retry if the destination service does not
	//      respond at all (disconnect/reset/read timeout)
	//
	//    connect-failure: Proxy will retry on failures connecting to destination
	//      for example due to connection timeouts.
	//
	//    retriable-4xx: Proxy will retry fro retriable 4xx response codes.
	//      Currently the only retriable error supported is 409.
	//
	//    refused-stream: Proxy will retry if the destination resets the stream
	//      with a REFUSED_STREAM error code. This reset type indicates that it
	//      is safe to retry.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RetryPolicy.retry_conditions
	RetryConditions []string `json:"retryConditions,omitempty"`

	// Specifies the allowed number of retries. This number must be > 0. If not
	//  specified, default to 1.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RetryPolicy.num_retries
	NumRetries *int32 `json:"numRetries,omitempty"`

	// Specifies a non-zero timeout per retry attempt.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RetryPolicy.per_try_timeout
	PerTryTimeout *string `json:"perTryTimeout,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.RouteAction
type HttpRoute_RouteAction struct {
	// The destination to which traffic should be forwarded.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteAction.destinations
	Destinations []HttpRoute_Destination `json:"destinations,omitempty"`

	// If set, the request is directed as configured by this field.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteAction.redirect
	Redirect *HttpRoute_Redirect `json:"redirect,omitempty"`

	// The specification for fault injection introduced into traffic to test the
	//  resiliency of clients to backend service failure. As part of fault
	//  injection, when clients send requests to a backend service, delays can be
	//  introduced  on a percentage of requests before sending those requests to
	//  the backend service. Similarly requests from clients can be aborted for a
	//  percentage of requests.
	//
	//  timeout and retry_policy will be ignored by clients that are configured
	//  with a fault_injection_policy
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteAction.fault_injection_policy
	FaultInjectionPolicy *HttpRoute_FaultInjectionPolicy `json:"faultInjectionPolicy,omitempty"`

	// The specification for modifying the headers of a matching request prior
	//  to delivery of the request to the destination. If HeaderModifiers are set
	//  on both the Destination and the RouteAction, they will be merged.
	//  Conflicts between the two will not be resolved on the configuration.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteAction.request_header_modifier
	RequestHeaderModifier *HttpRoute_HeaderModifier `json:"requestHeaderModifier,omitempty"`

	// The specification for modifying the headers of a response prior to
	//  sending the response back to the client. If HeaderModifiers are set
	//  on both the Destination and the RouteAction, they will be merged.
	//  Conflicts between the two will not be resolved on the configuration.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteAction.response_header_modifier
	ResponseHeaderModifier *HttpRoute_HeaderModifier `json:"responseHeaderModifier,omitempty"`

	// The specification for rewrite URL before forwarding requests to the
	//  destination.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteAction.url_rewrite
	URLRewrite *HttpRoute_URLRewrite `json:"urlRewrite,omitempty"`

	// Specifies the timeout for selected route. Timeout is computed from the
	//  time the request has been fully processed (i.e. end of stream) up until
	//  the response has been completely processed. Timeout includes all retries.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteAction.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Specifies the retry policy associated with this route.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteAction.retry_policy
	RetryPolicy *HttpRoute_RetryPolicy `json:"retryPolicy,omitempty"`

	// Specifies the policy on how requests intended for the routes destination
	//  are shadowed to a separate mirrored destination. Proxy will not wait for
	//  the shadow destination to respond before returning the response. Prior to
	//  sending traffic to the shadow service, the host/authority header is
	//  suffixed with -shadow.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteAction.request_mirror_policy
	RequestMirrorPolicy *HttpRoute_RequestMirrorPolicy `json:"requestMirrorPolicy,omitempty"`

	// The specification for allowing client side cross-origin requests.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteAction.cors_policy
	CorsPolicy *HttpRoute_CorsPolicy `json:"corsPolicy,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.RouteMatch
type HttpRoute_RouteMatch struct {
	// The HTTP request path value should exactly match this value.
	//
	//  Only one of full_path_match, prefix_match, or regex_match should be
	//  used.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteMatch.full_path_match
	FullPathMatch *string `json:"fullPathMatch,omitempty"`

	// The HTTP request path value must begin with specified prefix_match.
	//  prefix_match must begin with a /.
	//
	//  Only one of full_path_match, prefix_match, or regex_match should be
	//  used.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteMatch.prefix_match
	PrefixMatch *string `json:"prefixMatch,omitempty"`

	// The HTTP request path value must satisfy the regular expression
	//  specified by regex_match after removing any query parameters and anchor
	//  supplied with the original URL. For regular expression grammar, please
	//  see https://github.com/google/re2/wiki/Syntax
	//
	//  Only one of full_path_match, prefix_match, or regex_match should be
	//  used.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteMatch.regex_match
	RegexMatch *string `json:"regexMatch,omitempty"`

	// Specifies if prefix_match and full_path_match matches are case sensitive.
	//  The default value is false.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteMatch.ignore_case
	IgnoreCase *bool `json:"ignoreCase,omitempty"`

	// Specifies a list of HTTP request headers to match against. ALL of the
	//  supplied headers must be matched.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteMatch.headers
	Headers []HttpRoute_HeaderMatch `json:"headers,omitempty"`

	// Specifies a list of query parameters to match against. ALL of the query
	//  parameters must be matched.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteMatch.query_parameters
	QueryParameters []HttpRoute_QueryParameterMatch `json:"queryParameters,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.RouteRule
type HttpRoute_RouteRule struct {
	// A list of matches define conditions used for matching the rule against
	//  incoming HTTP requests. Each match is independent, i.e. this rule will be
	//  matched if ANY one of the matches is satisfied.
	//
	//  If no matches field is specified, this rule will unconditionally match
	//  traffic.
	//
	//  If a default rule is desired to be configured, add a rule with no matches
	//  specified to the end of the rules list.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteRule.matches
	Matches []HttpRoute_RouteMatch `json:"matches,omitempty"`

	// The detailed rule defining how to route matched traffic.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.RouteRule.action
	Action *HttpRoute_RouteAction `json:"action,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute.URLRewrite
type HttpRoute_URLRewrite struct {
	// Prior to forwarding the request to the selected destination, the matching
	//  portion of the requests path is replaced by this value.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.URLRewrite.path_prefix_rewrite
	PathPrefixRewrite *string `json:"pathPrefixRewrite,omitempty"`

	// Prior to forwarding the request to the selected destination, the requests
	//  host header is replaced by this value.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.URLRewrite.host_rewrite
	HostRewrite *string `json:"hostRewrite,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HttpRoute
type HttpRouteObservedState struct {
	// Output only. Server-defined URL of this resource
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.HttpRoute.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
