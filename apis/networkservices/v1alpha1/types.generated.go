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
// krm.group: networkservices.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.networkservices.v1
// resource: NetworkServicesEdgeCacheService:EdgeCacheService

package v1alpha1

// +kcc:proto=google.cloud.networkservices.v1.AddHeader
type AddHeader struct {
	// Required. The name of the header to add.
	// +kcc:proto:field=google.cloud.networkservices.v1.AddHeader.header_name
	HeaderName *string `json:"headerName,omitempty"`

	// Required. The value of the header to add.
	// +kcc:proto:field=google.cloud.networkservices.v1.AddHeader.header_value
	HeaderValue *string `json:"headerValue,omitempty"`

	// Optional. Whether to replace all existing headers with the same name.
	// +kcc:proto:field=google.cloud.networkservices.v1.AddHeader.replace
	Replace *bool `json:"replace,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.AddSignatures
type AddSignatures struct {
	// Required. The actions to take to add signatures to responses.
	// +kcc:proto:field=google.cloud.networkservices.v1.AddSignatures.actions
	Actions []string `json:"actions,omitempty"`

	// Optional. The keyset to use for signature generation.
	// +kcc:proto:field=google.cloud.networkservices.v1.AddSignatures.keyset
	Keyset *string `json:"keyset,omitempty"`

	// Optional. The TTL for the generated token.
	// +kcc:proto:field=google.cloud.networkservices.v1.AddSignatures.token_ttl
	TokenTTL *string `json:"tokenTTL,omitempty"`

	// Optional. The query parameter to use for the generated token.
	// +kcc:proto:field=google.cloud.networkservices.v1.AddSignatures.token_query_parameter
	TokenQueryParameter *string `json:"tokenQueryParameter,omitempty"`

	// Optional. Parameters to copy from the verified token to the generated token.
	// +kcc:proto:field=google.cloud.networkservices.v1.AddSignatures.copied_parameters
	CopiedParameters []string `json:"copiedParameters,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.CacheKeyPolicy
type CacheKeyPolicy struct {
	// Optional. If true, http and https requests will be cached separately.
	// +kcc:proto:field=google.cloud.networkservices.v1.CacheKeyPolicy.include_protocol
	IncludeProtocol *bool `json:"includeProtocol,omitempty"`

	// Optional. If true, requests to different hosts will be cached separately.
	// +kcc:proto:field=google.cloud.networkservices.v1.CacheKeyPolicy.exclude_host
	ExcludeHost *bool `json:"excludeHost,omitempty"`

	// Optional. If true, exclude query string parameters from the cache key.
	// +kcc:proto:field=google.cloud.networkservices.v1.CacheKeyPolicy.exclude_query_string
	ExcludeQueryString *bool `json:"excludeQueryString,omitempty"`

	// Optional. Names of query string parameters to include in cache keys.
	// +kcc:proto:field=google.cloud.networkservices.v1.CacheKeyPolicy.included_query_parameters
	IncludedQueryParameters []string `json:"includedQueryParameters,omitempty"`

	// Optional. Names of query string parameters to exclude from cache keys.
	// +kcc:proto:field=google.cloud.networkservices.v1.CacheKeyPolicy.excluded_query_parameters
	ExcludedQueryParameters []string `json:"excludedQueryParameters,omitempty"`

	// Optional. Names of HTTP request headers to include in cache keys.
	// +kcc:proto:field=google.cloud.networkservices.v1.CacheKeyPolicy.included_header_names
	IncludedHeaderNames []string `json:"includedHeaderNames,omitempty"`

	// Optional. Names of Cookies to include in cache keys.
	// +kcc:proto:field=google.cloud.networkservices.v1.CacheKeyPolicy.included_cookie_names
	IncludedCookieNames []string `json:"includedCookieNames,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.CdnPolicy
type CDNPolicy struct {
	// Optional. The CacheMode for matching requests.
	// +kcc:proto:field=google.cloud.networkservices.v1.CdnPolicy.cache_mode
	CacheMode *string `json:"cacheMode,omitempty"`

	// Optional. Specifies a separate client TTL, separate from the TTL used for
	//  Google's cache for the object.
	// +kcc:proto:field=google.cloud.networkservices.v1.CdnPolicy.client_ttl
	ClientTTL *string `json:"clientTTL,omitempty"`

	// Optional. Specifies the default TTL for cached content served by this origin
	//  for responses that do not have an existing valid TTL (max-age or s-max-age).
	// +kcc:proto:field=google.cloud.networkservices.v1.CdnPolicy.default_ttl
	DefaultTTL *string `json:"defaultTTL,omitempty"`

	// Optional. Specifies the maximum allowed TTL for cached content served by
	//  this origin.
	// +kcc:proto:field=google.cloud.networkservices.v1.CdnPolicy.max_ttl
	MaxTTL *string `json:"maxTTL,omitempty"`

	// Optional. Negative caching allows per-status code cache TTLs to be set,
	//  in order to apply fine-grained caching for common errors or redirects.
	// +kcc:proto:field=google.cloud.networkservices.v1.CdnPolicy.negative_caching
	NegativeCaching *bool `json:"negativeCaching,omitempty"`

	// TODO: unsupported map type with key string and value message

	// Optional. Defines the request parameters that contribute to the cache key.
	// +kcc:proto:field=google.cloud.networkservices.v1.CdnPolicy.cache_key_policy
	CacheKeyPolicy *CacheKeyPolicy `json:"cacheKeyPolicy,omitempty"`

	// Optional. Whether to enforce signed requests.
	// +kcc:proto:field=google.cloud.networkservices.v1.CdnPolicy.signed_request_mode
	SignedRequestMode *string `json:"signedRequestMode,omitempty"`

	// Optional. The EdgeCacheKeyset resource that contains the set of public keys
	//  used to validate signed requests at the edge.
	// +kcc:proto:field=google.cloud.networkservices.v1.CdnPolicy.signed_request_keyset
	SignedRequestKeyset *string `json:"signedRequestKeyset,omitempty"`

	// Optional. Additional options for signed requests.
	// +kcc:proto:field=google.cloud.networkservices.v1.CdnPolicy.signed_token_options
	SignedTokenOptions *SignedTokenOptions `json:"signedTokenOptions,omitempty"`

	// Optional. Enable signature generation or propagation on this route.
	// +kcc:proto:field=google.cloud.networkservices.v1.CdnPolicy.add_signatures
	AddSignatures *AddSignatures `json:"addSignatures,omitempty"`

	// Optional. Defines the maximum allowed TTL for signed requests.
	// +kcc:proto:field=google.cloud.networkservices.v1.CdnPolicy.signed_request_maximum_expiration_ttl
	SignedRequestMaximumExpirationTTL *string `json:"signedRequestMaximumExpirationTTL,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.CorsPolicy
type CorsPolicy struct {
	// Required. Specifies the content for the Access-Control-Max-Age header.
	// +kcc:proto:field=google.cloud.networkservices.v1.CorsPolicy.max_age
	MaxAge *string `json:"maxAge,omitempty"`

	// Optional. Specifies the list of origins that will be allowed to do CORS
	//  requests.
	// +kcc:proto:field=google.cloud.networkservices.v1.CorsPolicy.allow_origins
	AllowOrigins []string `json:"allowOrigins,omitempty"`

	// Optional. Specifies the content for the
	//  Access-Control-Allow-Methods header.
	// +kcc:proto:field=google.cloud.networkservices.v1.CorsPolicy.allow_methods
	AllowMethods []string `json:"allowMethods,omitempty"`

	// Optional. Specifies the content for the
	//  Access-Control-Allow-Headers header.
	// +kcc:proto:field=google.cloud.networkservices.v1.CorsPolicy.allow_headers
	AllowHeaders []string `json:"allowHeaders,omitempty"`

	// Optional. Specifies the content for the
	//  Access-Control-Expose-Headers header.
	// +kcc:proto:field=google.cloud.networkservices.v1.CorsPolicy.expose_headers
	ExposeHeaders []string `json:"exposeHeaders,omitempty"`

	// Optional. Specifies whether credentials are allowed in CORS requests.
	// +kcc:proto:field=google.cloud.networkservices.v1.CorsPolicy.allow_credentials
	AllowCredentials *bool `json:"allowCredentials,omitempty"`

	// Optional. If true, CORS policy is disabled. Defaults to false.
	// +kcc:proto:field=google.cloud.networkservices.v1.CorsPolicy.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.EdgeCacheService
type EdgeCacheService struct {
	// Identifier. Name of the EdgeCacheService resource. Format:
	//  projects/{project}/locations/{location}/edgeCacheServices/{edge_cache_service}
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.name
	Name *string `json:"name,omitempty"`

	// Optional. Set of label tags associated with the EdgeCacheService resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A human-readable description of the resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.description
	Description *string `json:"description,omitempty"`

	// Required. Defines how requests are routed, modified, cached and/or
	//  which origin content is filled from.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.routing
	Routing *Routing `json:"routing,omitempty"`

	// Optional. URLs to sslCertificate resources that are used to authenticate
	//  connections between users and the EdgeCacheService.
	//  Note that only "global" certificates with a "scope" of "EDGE_CACHE"
	//  can be attached to an EdgeCacheService.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.edge_ssl_certificates
	EdgeSSLCertificates []string `json:"edgeSSLCertificates,omitempty"`

	// Optional. HTTP/3 (IETF QUIC) and Google QUIC are enabled by default.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.disable_quic
	DisableQuic *bool `json:"disableQuic,omitempty"`

	// Optional. Specifies the logging options for the traffic served by this service.
	//  If logging is enabled, logs will be exported to Cloud Logging.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.log_config
	LogConfig *LogConfig `json:"logConfig,omitempty"`

	// Optional. Disables HTTP/2.
	//  HTTP/2 (h2) is enabled by default and recommended for performance.
	//  HTTP/2 improves connection re-use and reduces connection setup overhead
	//  by sending multiple streams over the same connection.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.disable_http2
	DisableHttp2 *bool `json:"disableHttp2,omitempty"`

	// Optional. Require TLS (HTTPS) for all clients connecting to this service.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.require_tls
	RequireTLS *bool `json:"requireTLS,omitempty"`

	// Optional. Resource URL that points at the Cloud Armor edge security policy
	//  that is applied on each request against the EdgeCacheService.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.edge_security_policy
	EdgeSecurityPolicy *string `json:"edgeSecurityPolicy,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HeaderAction
type HeaderAction struct {
	// Optional. Describes a header to add.
	// +kcc:proto:field=google.cloud.networkservices.v1.HeaderAction.request_headers_to_add
	RequestHeadersToAdd []AddHeader `json:"requestHeadersToAdd,omitempty"`

	// Optional. A list of header names for headers that need to be removed
	//  from the request prior to forwarding the request to the origin.
	// +kcc:proto:field=google.cloud.networkservices.v1.HeaderAction.request_headers_to_remove
	RequestHeadersToRemove []RemoveHeader `json:"requestHeadersToRemove,omitempty"`

	// Optional. Describes a header to add to the response.
	// +kcc:proto:field=google.cloud.networkservices.v1.HeaderAction.response_headers_to_add
	ResponseHeadersToAdd []AddHeader `json:"responseHeadersToAdd,omitempty"`

	// Optional. A list of header names for headers that need to be removed
	//  from the response prior to sending the response back to the client.
	// +kcc:proto:field=google.cloud.networkservices.v1.HeaderAction.response_headers_to_remove
	ResponseHeadersToRemove []RemoveHeader `json:"responseHeadersToRemove,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HeaderMatch
type HeaderMatch struct {
	// Required. The header name to match on.
	// +kcc:proto:field=google.cloud.networkservices.v1.HeaderMatch.header_name
	HeaderName *string `json:"headerName,omitempty"`

	// Optional. Specifies that the match is negated, i.e. a match will succeed
	//  if the header does NOT match the criteria in this HeaderMatch.
	// +kcc:proto:field=google.cloud.networkservices.v1.HeaderMatch.invert_match
	InvertMatch *bool `json:"invertMatch,omitempty"`

	// Specifies that the header must exist.
	// +kcc:proto:field=google.cloud.networkservices.v1.HeaderMatch.present_match
	PresentMatch *bool `json:"presentMatch,omitempty"`

	// The value of the header must exactly match the contents of exactMatch.
	// +kcc:proto:field=google.cloud.networkservices.v1.HeaderMatch.exact_match
	ExactMatch *string `json:"exactMatch,omitempty"`

	// The value of the header must start with the contents of prefixMatch.
	// +kcc:proto:field=google.cloud.networkservices.v1.HeaderMatch.prefix_match
	PrefixMatch *string `json:"prefixMatch,omitempty"`

	// The value of the header must end with the contents of suffixMatch.
	// +kcc:proto:field=google.cloud.networkservices.v1.HeaderMatch.suffix_match
	SuffixMatch *string `json:"suffixMatch,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.HostRule
type HostRule struct {
	// Optional. A human-readable description of the hostRule.
	// +kcc:proto:field=google.cloud.networkservices.v1.HostRule.description
	Description *string `json:"description,omitempty"`

	// Required. The list of host patterns to match.
	//  Host patterns must be valid hostnames. Ports are not allowed.
	//  Wildcard hosts are supported in the form of "*" or "*-foo.example.com".
	//  This field is used to match the ":authority" header for HTTP/2 requests
	//  or the "Host" header for HTTP/1.1 requests.
	// +kcc:proto:field=google.cloud.networkservices.v1.HostRule.hosts
	Hosts []string `json:"hosts,omitempty"`

	// Required. The name of the pathMatcher associated with this hostRule.
	// +kcc:proto:field=google.cloud.networkservices.v1.HostRule.path_matcher
	PathMatcher *string `json:"pathMatcher,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.LogConfig
type LogConfig struct {
	// Optional. Specifies whether to enable logging for traffic served by this
	//  service.
	// +kcc:proto:field=google.cloud.networkservices.v1.LogConfig.enable
	Enable *bool `json:"enable,omitempty"`

	// Optional. Configures the sampling rate of requests. Must be between 0.0
	//  and 1.0 inclusive.
	// +kcc:proto:field=google.cloud.networkservices.v1.LogConfig.sample_rate
	SampleRate *float32 `json:"sampleRate,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.MatchRule
type MatchRule struct {
	// Optional. For satisfying the matchRule condition, the path of the request
	//  must match the wildcard pattern specified in pathTemplateMatch
	//  after removing any query parameters and anchor that may be part of
	//  the original URL.
	// +kcc:proto:field=google.cloud.networkservices.v1.MatchRule.path_template_match
	PathTemplateMatch *string `json:"pathTemplateMatch,omitempty"`

	// Optional. For satisfying the matchRule condition, the request's path must
	//  begin with the specified prefixMatch. prefixMatch must begin with a /.
	// +kcc:proto:field=google.cloud.networkservices.v1.MatchRule.prefix_match
	PrefixMatch *string `json:"prefixMatch,omitempty"`

	// Optional. For satisfying the matchRule condition, the path of the request
	//  must exactly match the value specified in fullPathMatch after removing
	//  any query parameters and anchor that may be part of the original URL.
	// +kcc:proto:field=google.cloud.networkservices.v1.MatchRule.full_path_match
	FullPathMatch *string `json:"fullPathMatch,omitempty"`

	// Optional. Specifies that prefixMatch and fullPathMatch matches are case
	//  insensitive. The default value is false.
	// +kcc:proto:field=google.cloud.networkservices.v1.MatchRule.ignore_case
	IgnoreCase *bool `json:"ignoreCase,omitempty"`

	// Optional. Specifies a list of header match criteria, all of which must match
	//  corresponding headers in the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.MatchRule.header_matches
	HeaderMatches []HeaderMatch `json:"headerMatches,omitempty"`

	// Optional. Specifies a list of query parameter match criteria, all of which
	//  must match corresponding query parameters in the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.MatchRule.query_parameter_matches
	QueryParameterMatches []QueryParameterMatch `json:"queryParameterMatches,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.PathMatcher
type PathMatcher struct {
	// Required. The name to which this PathMatcher is referred by the HostRule.
	// +kcc:proto:field=google.cloud.networkservices.v1.PathMatcher.name
	Name *string `json:"name,omitempty"`

	// Optional. A human-readable description of the resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.PathMatcher.description
	Description *string `json:"description,omitempty"`

	// Required. The routeRules to match against. Each rule must reference
	//  an origin via the route_action or a urlRedirect.
	// +kcc:proto:field=google.cloud.networkservices.v1.PathMatcher.route_rules
	RouteRules []RouteRule `json:"routeRules,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.QueryParameterMatch
type QueryParameterMatch struct {
	// Required. The name of the query parameter to match.
	// +kcc:proto:field=google.cloud.networkservices.v1.QueryParameterMatch.name
	Name *string `json:"name,omitempty"`

	// Specifies that the query parameter must exist.
	// +kcc:proto:field=google.cloud.networkservices.v1.QueryParameterMatch.present_match
	PresentMatch *bool `json:"presentMatch,omitempty"`

	// The value of the query parameter must exactly match the contents
	//  of exactMatch.
	// +kcc:proto:field=google.cloud.networkservices.v1.QueryParameterMatch.exact_match
	ExactMatch *string `json:"exactMatch,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.RemoveHeader
type RemoveHeader struct {
	// Required. The name of the header to remove.
	// +kcc:proto:field=google.cloud.networkservices.v1.RemoveHeader.header_name
	HeaderName *string `json:"headerName,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.RouteAction
type RouteAction struct {
	// Optional. The policy that determines the caching behavior for the
	//  route.
	// +kcc:proto:field=google.cloud.networkservices.v1.RouteAction.cdn_policy
	CDNPolicy *CDNPolicy `json:"cdnPolicy,omitempty"`

	// Optional. If specified, the request and response are transformed
	//  based on these rules.
	// +kcc:proto:field=google.cloud.networkservices.v1.RouteAction.url_rewrite
	URLRewrite *URLRewrite `json:"urlRewrite,omitempty"`

	// Optional. The policy for managing cross-origin resource sharing.
	// +kcc:proto:field=google.cloud.networkservices.v1.RouteAction.cors_policy
	CorsPolicy *CorsPolicy `json:"corsPolicy,omitempty"`

	// Optional. Specify how to compress the response to the client.
	//  By default, responses are not compressed.
	// +kcc:proto:field=google.cloud.networkservices.v1.RouteAction.compression_mode
	CompressionMode *string `json:"compressionMode,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.RouteMethods
type RouteMethods struct {
	// Required. The list of HTTP methods allowed for this route.
	//  Must be non-empty.
	// +kcc:proto:field=google.cloud.networkservices.v1.RouteMethods.allowed_methods
	AllowedMethods []string `json:"allowedMethods,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.RouteRule
type RouteRule struct {
	// Required. The priority of this route rule, where 1 is the highest priority.
	//  You cannot configure two or more routeRules with the same priority.
	//  Priority for each rule must be set to a number between 1 and 999 inclusive.
	// +kcc:proto:field=google.cloud.networkservices.v1.RouteRule.priority
	Priority *int64 `json:"priority,omitempty"`

	// Optional. A human-readable description of the routeRule.
	// +kcc:proto:field=google.cloud.networkservices.v1.RouteRule.description
	Description *string `json:"description,omitempty"`

	// Required. The list of criteria for matching attributes of a request to this
	//  routeRule. All specified conditions must be satisfied for a match.
	//  Must specify at least one match_rule.
	// +kcc:proto:field=google.cloud.networkservices.v1.RouteRule.match_rules
	MatchRules []MatchRule `json:"matchRules,omitempty"`

	// Optional. The header actions, including adding and removing headers,
	//  for requests that match this route.
	// +kcc:proto:field=google.cloud.networkservices.v1.RouteRule.header_action
	HeaderAction *HeaderAction `json:"headerAction,omitempty"`

	// Optional. The route action that will be taken if this routeRule matches.
	// +kcc:proto:field=google.cloud.networkservices.v1.RouteRule.route_action
	RouteAction *RouteAction `json:"routeAction,omitempty"`

	// Optional. The URL redirect configuration for requests that match this route.
	// +kcc:proto:field=google.cloud.networkservices.v1.RouteRule.url_redirect
	URLRedirect *URLRedirect `json:"urlRedirect,omitempty"`

	// Optional. The origin that requests are routed to when this rule matches.
	//  Origin must be a valid origin resource name. It should be formatted as:
	//  "projects/{project}/locations/{location}/edgeCacheOrigins/{origin}".
	// +kcc:proto:field=google.cloud.networkservices.v1.RouteRule.origin
	Origin *string `json:"origin,omitempty"`

	// Optional. routeMethods defines what HTTP methods may match this route.
	// +kcc:proto:field=google.cloud.networkservices.v1.RouteRule.route_methods
	RouteMethods *RouteMethods `json:"routeMethods,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.Routing
type Routing struct {
	// Required. The list of hostRules to match against. These rules define
	//  which hostnames the EdgeCacheService will match against, and which
	//  route to use for the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.Routing.host_rules
	HostRules []HostRule `json:"hostRules,omitempty"`

	// Required. The list of pathMatchers referenced via name by hostRules.
	//  PathMatcher is used to match the path portion of the URL when a
	//  HostRule matches the URL's host portion.
	// +kcc:proto:field=google.cloud.networkservices.v1.Routing.path_matchers
	PathMatchers []PathMatcher `json:"pathMatchers,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.SignedTokenOptions
type SignedTokenOptions struct {
	// Optional. The query parameter name where the signed token should be found.
	// +kcc:proto:field=google.cloud.networkservices.v1.SignedTokenOptions.token_query_parameter
	TokenQueryParameter *string `json:"tokenQueryParameter,omitempty"`

	// Optional. The allowed signature algorithms.
	// +kcc:proto:field=google.cloud.networkservices.v1.SignedTokenOptions.allowed_signature_algorithms
	AllowedSignatureAlgorithms []string `json:"allowedSignatureAlgorithms,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.UrlRedirect
type URLRedirect struct {
	// Optional. The host that will be used in the redirect response instead of
	//  the one that was supplied in the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.UrlRedirect.host_redirect
	HostRedirect *string `json:"hostRedirect,omitempty"`

	// Optional. The path that will be used in the redirect response instead of
	//  the one that was supplied in the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.UrlRedirect.path_redirect
	PathRedirect *string `json:"pathRedirect,omitempty"`

	// Optional. The prefix that will be used in the redirect response instead of
	//  the one that was supplied in the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.UrlRedirect.prefix_redirect
	PrefixRedirect *string `json:"prefixRedirect,omitempty"`

	// Optional. The HTTP Status code to use for this RedirectAction.
	// +kcc:proto:field=google.cloud.networkservices.v1.UrlRedirect.redirect_response_code
	RedirectResponseCode *string `json:"redirectResponseCode,omitempty"`

	// Optional. If set to true, the URL scheme in the redirected request is set
	//  to https.
	// +kcc:proto:field=google.cloud.networkservices.v1.UrlRedirect.https_redirect
	HTTPSRedirect *bool `json:"httpsRedirect,omitempty"`

	// Optional. If set to true, any accompanying query portion of the original
	//  URL is removed prior to redirecting the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.UrlRedirect.strip_query
	StripQuery *bool `json:"stripQuery,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.UrlRewrite
type URLRewrite struct {
	// Optional. Prior to forwarding the request to the selected origin, the
	//  matching portion of the request's path is replaced by pathPrefixRewrite.
	// +kcc:proto:field=google.cloud.networkservices.v1.UrlRewrite.path_prefix_rewrite
	PathPrefixRewrite *string `json:"pathPrefixRewrite,omitempty"`

	// Optional. Prior to forwarding the request to the selected origin, if the
	//  request matched a pathTemplateMatch, the matching portion of the
	//  request's path is replaced by pathTemplateRewrite.
	// +kcc:proto:field=google.cloud.networkservices.v1.UrlRewrite.path_template_rewrite
	PathTemplateRewrite *string `json:"pathTemplateRewrite,omitempty"`

	// Optional. Prior to forwarding the request to the selected origin, the
	//  request's host header is replaced by hostRewrite.
	// +kcc:proto:field=google.cloud.networkservices.v1.UrlRewrite.host_rewrite
	HostRewrite *string `json:"hostRewrite,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.networkservices.v1.EdgeCacheService
type EdgeCacheServiceObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was last updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The IPv4 addresses that the service will serve traffic on.
	//  These will be static for the lifetime of the service.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.ipv4_addresses
	IPV4Addresses []string `json:"ipv4Addresses,omitempty"`

	// Output only. The IPv6 addresses that the service will serve traffic on.
	//  These will be static for the lifetime of the service.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.ipv6_addresses
	IPV6Addresses []string `json:"ipv6Addresses,omitempty"`
}
