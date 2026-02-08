// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkServicesEdgeCacheServiceGVK = GroupVersion.WithKind("NetworkServicesEdgeCacheService")

// ResourceRef identifies a resource
type ResourceRef struct {
	/* The external name of the referenced resource */
	External string `json:"external,omitempty"`
	/* Kind of the referent. */
	Kind string `json:"kind,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

// NetworkServicesEdgeCacheServiceSpec defines the desired state of NetworkServicesEdgeCacheService
// +kcc:spec:proto=google.cloud.networkservices.v1.EdgeCacheService
// Note: The Direct controller implementation for this resource is currently blocked
// by missing support in the cloud.google.com/go/networkservices client library.
// The resource is currently managed by the Terraform-based controller.
type NetworkServicesEdgeCacheServiceSpec struct {
	// Optional. A human-readable description of the resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.description
	Description *string `json:"description,omitempty"`

	// Optional. Disables HTTP/2.
	// HTTP/2 (h2) is enabled by default and recommended for performance. HTTP/2 improves connection re-use and reduces connection setup overhead by sending multiple streams over the same connection.
	// Some legacy HTTP clients may have issues with HTTP/2 connections due to broken HTTP/2 implementations. Setting this to true will prevent HTTP/2 from being advertised and negotiated.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.disable_http2
	DisableHttp2 *bool `json:"disableHttp2,omitempty"`

	// Optional. HTTP/3 (IETF QUIC) and Google QUIC are enabled by default.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.disable_quic
	DisableQuic *bool `json:"disableQuic,omitempty"`

	// Optional. Resource URL that points at the Cloud Armor edge security policy that is applied on each request against the EdgeCacheService.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.edge_security_policy
	EdgeSecurityPolicyRef *ResourceRef `json:"edgeSecurityPolicyRef,omitempty"`

	// Optional. URLs to sslCertificate resources that are used to authenticate connections between users and the EdgeCacheService.
	// Note that only "global" certificates with a "scope" of "EDGE_CACHE" can be attached to an EdgeCacheService.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.edge_ssl_certificates
	EdgeSslCertificates []ResourceRef `json:"edgeSslCertificates,omitempty"`

	// Optional. Specifies the logging options for the traffic served by this service. If logging is enabled, logs will be exported to Cloud Logging.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.log_config
	LogConfig *EdgeCacheServiceLogConfig `json:"logConfig,omitempty"`

	// Required. The project that this resource belongs to.
	ProjectRef *v1beta1.ProjectRef `json:"projectRef"`

	// Optional. Require TLS (HTTPS) for all clients connecting to this service.
	// Clients who connect over HTTP (port 80) will receive a HTTP 301 to the same URL over HTTPS (port 443).
	// You must have at least one (1) edgeSslCertificate specified to enable this.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.require_tls
	RequireTls *bool `json:"requireTls,omitempty"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Defines how requests are routed, modified, cached and/or which origin content is filled from.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.routing
	Routing EdgeCacheServiceRouting `json:"routing"`

	// Optional. URL of the SslPolicy resource that will be associated with the EdgeCacheService.
	// If not set, the EdgeCacheService has no SSL policy configured, and will default to the "COMPATIBLE" policy.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.ssl_policy
	SslPolicyRef *ResourceRef `json:"sslPolicyRef,omitempty"`
}

type EdgeCacheServiceLogConfig struct {
	// Optional. Specifies whether to enable logging for traffic served by this service.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.LogConfig.enable
	Enable *bool `json:"enable,omitempty"`

	// Optional. Configures the sampling rate of requests, where 1.0 means all logged requests are reported and 0.0 means no logged requests are reported. The default value is 1.0, and the value of the field must be in [0, 1].
	// This field can only be specified if logging is enabled for this service.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.LogConfig.sample_rate
	SampleRate *float64 `json:"sampleRate,omitempty"`
}

type EdgeCacheServiceRouting struct {
	// Required. The list of hostRules to match against. These rules define which hostnames the EdgeCacheService will match against, and which route configurations apply.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.host_rule
	HostRule []EdgeCacheServiceHostRule `json:"hostRule"`

	// Required. The list of pathMatchers referenced via name by hostRules. PathMatcher is used to match the path portion of the URL when a HostRule matches the URL's host portion.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.path_matcher
	PathMatcher []EdgeCacheServicePathMatcher `json:"pathMatcher"`
}

type EdgeCacheServiceHostRule struct {
	// Optional. A human-readable description of the hostRule.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.HostRule.description
	Description *string `json:"description,omitempty"`

	// Required. The list of host patterns to match.
	// Host patterns must be valid hostnames. Ports are not allowed. Wildcard hosts are supported in the suffix or prefix form. * matches any string of ([a-z0-9-.]*). It does not match the empty string.
	// When multiple hosts are specified, hosts are matched in the following priority:
	// 1. Exact domain names: ''www.foo.com''.
	// 2. Suffix domain wildcards: ''*.foo.com'' or ''*-bar.foo.com''.
	// 3. Prefix domain wildcards: ''foo.*'' or ''foo-*''.
	// 4. Special wildcard ''*'' matching any domain.
	// Notes:
	// The wildcard will not match the empty string. e.g. ''*-bar.foo.com'' will match ''baz-bar.foo.com'' but not ''-bar.foo.com''. The longest wildcards match first. Only a single host in the entire service can match on ''*''. A domain must be unique across all configured hosts within a service.
	// Hosts are matched against the HTTP Host header, or for HTTP/2 and HTTP/3, the ":authority" header, from the incoming request.
	// You may specify up to 10 hosts.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.HostRule.hosts
	Hosts []string `json:"hosts"`

	// Required. The name of the pathMatcher associated with this hostRule.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.HostRule.path_matcher
	PathMatcher string `json:"pathMatcher"`
}

type EdgeCacheServicePathMatcher struct {
	// Optional. A human-readable description of the resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.PathMatcher.description
	Description *string `json:"description,omitempty"`

	// Required. The name to which this PathMatcher is referred by the HostRule.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.PathMatcher.name
	Name string `json:"name"`

	// Required. The routeRules to match against. routeRules support advanced routing behaviour, and can match on paths, headers and query parameters, as well as status codes and HTTP methods.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.PathMatcher.route_rule
	RouteRule []EdgeCacheServiceRouteRule `json:"routeRule"`
}

type EdgeCacheServiceRouteRule struct {
	// Optional. A human-readable description of the routeRule.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.description
	Description *string `json:"description,omitempty"`

	// Optional. The header actions, including adding & removing headers, for requests that match this route.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.header_action
	HeaderAction *EdgeCacheServiceHeaderAction `json:"headerAction,omitempty"`

	// Required. The list of criteria for matching attributes of a request to this routeRule. This list has OR semantics: the request matches this routeRule when any of the matchRules are satisfied. However predicates
	// within a given matchRule have AND semantics. All predicates within a matchRule must match for the request to match the rule.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.match_rule
	MatchRule []EdgeCacheServiceMatchRule `json:"matchRule"`

	// Optional. The Origin resource that requests to this route should fetch from when a matching response is not in cache. Origins can be defined as short names ("my-origin") or fully-qualified resource URLs - e.g. "networkservices.googleapis.com/projects/my-project/global/edgecacheorigins/my-origin"
	// Only one of origin or urlRedirect can be set.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.origin
	OriginRef *ResourceRef `json:"originRef,omitempty"`

	// Required. The priority of this route rule, where 1 is the highest priority.
	// You cannot configure two or more routeRules with the same priority. Priority for each rule must be set to a number between 1 and 999 inclusive.
	// Priority numbers can have gaps, which enable you to add or remove rules in the future without affecting the rest of the rules. For example, 1, 2, 3, 4, 5, 9, 12, 16 is a valid series of priority numbers
	// to which you could add rules numbered from 6 to 8, 10 to 11, and 13 to 15 in the future without any impact on existing rules.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.priority
	Priority string `json:"priority"`

	// Optional. In response to a matching path, the routeAction performs advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request to the selected origin.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.route_action
	RouteAction *EdgeCacheServiceRouteAction `json:"routeAction,omitempty"`

	// Optional. The URL redirect configuration for requests that match this route.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.url_redirect
	UrlRedirect *EdgeCacheServiceUrlRedirect `json:"urlRedirect,omitempty"`
}

type EdgeCacheServiceHeaderAction struct {
	// Optional. Describes a header to add.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.HeaderAction.request_header_to_add
	RequestHeaderToAdd []EdgeCacheServiceRequestHeaderToAdd `json:"requestHeaderToAdd,omitempty"`

	// Optional. A list of header names for headers that need to be removed from the request prior to forwarding the request to the origin.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.HeaderAction.request_header_to_remove
	RequestHeaderToRemove []EdgeCacheServiceRequestHeaderToRemove `json:"requestHeaderToRemove,omitempty"`

	// Optional. Headers to add to the response prior to sending it back to the client.
	// Response headers are only sent to the client, and do not have an effect on the cache serving the response.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.HeaderAction.response_header_to_add
	ResponseHeaderToAdd []EdgeCacheServiceResponseHeaderToAdd `json:"responseHeaderToAdd,omitempty"`

	// Optional. A list of header names for headers that need to be removed from the request prior to forwarding the request to the origin.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.HeaderAction.response_header_to_remove
	ResponseHeaderToRemove []EdgeCacheServiceResponseHeaderToRemove `json:"responseHeaderToRemove,omitempty"`
}

type EdgeCacheServiceRequestHeaderToAdd struct {
	// Required. The name of the header to add.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.HeaderAction.RequestHeaderToAdd.header_name
	HeaderName string `json:"headerName"`

	// Required. The value of the header to add.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.HeaderAction.RequestHeaderToAdd.header_value
	HeaderValue string `json:"headerValue"`

	// Optional. Whether to replace all existing headers with the same name.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.HeaderAction.RequestHeaderToAdd.replace
	Replace *bool `json:"replace,omitempty"`
}

type EdgeCacheServiceRequestHeaderToRemove struct {
	// Required. The name of the header to remove.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.HeaderAction.RequestHeaderToRemove.header_name
	HeaderName string `json:"headerName"`
}

type EdgeCacheServiceResponseHeaderToAdd struct {
	// Required. The name of the header to add.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.HeaderAction.ResponseHeaderToAdd.header_name
	HeaderName string `json:"headerName"`

	// Required. The value of the header to add.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.HeaderAction.ResponseHeaderToAdd.header_value
	HeaderValue string `json:"headerValue"`

	// Optional. Whether to replace all existing headers with the same name.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.HeaderAction.ResponseHeaderToAdd.replace
	Replace *bool `json:"replace,omitempty"`
}

type EdgeCacheServiceResponseHeaderToRemove struct {
	// Required. Headers to remove from the response prior to sending it back to the client.
	// Response headers are only sent to the client, and do not have an effect on the cache serving the response.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.HeaderAction.ResponseHeaderToRemove.header_name
	HeaderName string `json:"headerName"`
}

type EdgeCacheServiceMatchRule struct {
	// Optional. For satisfying the matchRule condition, the path of the request must exactly match the value specified in fullPathMatch after removing any query parameters and anchor that may be part of the original URL.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.full_path_match
	FullPathMatch *string `json:"fullPathMatch,omitempty"`

	// Optional. Specifies a list of header match criteria, all of which must match corresponding headers in the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.header_match
	HeaderMatch []EdgeCacheServiceHeaderMatch `json:"headerMatch,omitempty"`

	// Optional. Specifies that prefixMatch and fullPathMatch matches are case sensitive.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.ignore_case
	IgnoreCase *bool `json:"ignoreCase,omitempty"`

	// Optional. For satisfying the matchRule condition, the path of the request
	// must match the wildcard pattern specified in pathTemplateMatch
	// after removing any query parameters and anchor that may be part
	// of the original URL.
	// pathTemplateMatch must be between 1 and 255 characters
	// (inclusive).  The pattern specified by pathTemplateMatch may
	// have at most 5 wildcard operators and at most 5 variable
	// captures in total.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.path_template_match
	PathTemplateMatch *string `json:"pathTemplateMatch,omitempty"`

	// Optional. For satisfying the matchRule condition, the request's path must begin with the specified prefixMatch. prefixMatch must begin with a /.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.prefix_match
	PrefixMatch *string `json:"prefixMatch,omitempty"`

	// Optional. Specifies a list of query parameter match criteria, all of which must match corresponding query parameters in the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.query_parameter_match
	QueryParameterMatch []EdgeCacheServiceQueryParameterMatch `json:"queryParameterMatch,omitempty"`
}

type EdgeCacheServiceHeaderMatch struct {
	// Optional. The value of the header should exactly match contents of exactMatch.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.HeaderMatch.exact_match
	ExactMatch *string `json:"exactMatch,omitempty"`

	// Required. The header name to match on.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.HeaderMatch.header_name
	HeaderName string `json:"headerName"`

	// Optional. If set to false (default), the headerMatch is considered a match if the match criteria above are met.
	// If set to true, the headerMatch is considered a match if the match criteria above are NOT met.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.HeaderMatch.invert_match
	InvertMatch *bool `json:"invertMatch,omitempty"`

	// Optional. The value of the header must start with the contents of prefixMatch.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.HeaderMatch.prefix_match
	PrefixMatch *string `json:"prefixMatch,omitempty"`

	// Optional. A header with the contents of headerName must exist. The match takes place whether or not the request's header has a value.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.HeaderMatch.present_match
	PresentMatch *bool `json:"presentMatch,omitempty"`

	// Optional. The value of the header must end with the contents of suffixMatch.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.HeaderMatch.suffix_match
	SuffixMatch *string `json:"suffixMatch,omitempty"`
}

type EdgeCacheServiceQueryParameterMatch struct {
	// Optional. The queryParameterMatch matches if the value of the parameter exactly matches the contents of exactMatch.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.QueryParameterMatch.exact_match
	ExactMatch *string `json:"exactMatch,omitempty"`

	// Required. The name of the query parameter to match. The query parameter must exist in the request, in the absence of which the request match fails.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.QueryParameterMatch.name
	Name string `json:"name"`

	// Optional. Specifies that the queryParameterMatch matches if the request contains the query parameter, irrespective of whether the parameter has a value or not.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.MatchRule.QueryParameterMatch.present_match
	PresentMatch *bool `json:"presentMatch,omitempty"`
}

type EdgeCacheServiceRouteAction struct {
	// Optional. The policy to use for defining caching and signed request behaviour for requests that match this route.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.cdn_policy
	CdnPolicy *EdgeCacheServiceCdnPolicy `json:"cdnPolicy,omitempty"`

	// Optional. CORSPolicy defines Cross-Origin-Resource-Sharing configuration, including which CORS response headers will be set.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.cors_policy
	CorsPolicy *EdgeCacheServiceCorsPolicy `json:"corsPolicy,omitempty"`

	// Optional. The URL rewrite configuration for requests that match this route.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.url_rewrite
	UrlRewrite *EdgeCacheServiceUrlRewrite `json:"urlRewrite,omitempty"`
}

type EdgeCacheServiceCdnPolicy struct {
	// Optional. Enable signature generation or propagation on this route.
	// This field may only be specified when signedRequestMode is set to REQUIRE_TOKENS.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.add_signatures
	AddSignatures *EdgeCacheServiceAddSignatures `json:"addSignatures,omitempty"`

	// Optional. Defines the request parameters that contribute to the cache key.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.cache_key_policy
	CacheKeyPolicy *EdgeCacheServiceCacheKeyPolicy `json:"cacheKeyPolicy,omitempty"`

	// Optional. Cache modes allow users to control the behaviour of the cache, what content it should cache automatically, whether to respect origin headers, or whether to unconditionally cache all responses.
	// For all cache modes, Cache-Control headers will be passed to the client. Use clientTtl to override what is sent to the client. Possible values: ["CACHE_ALL_STATIC", "USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", "BYPASS_CACHE"].
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.cache_mode
	CacheMode *string `json:"cacheMode,omitempty"`

	// Optional. Specifies a separate client (e.g. browser client) TTL, separate from the TTL used by the edge caches. Leaving this empty will use the same cache TTL for both the CDN and the client-facing response.
	// - The TTL must be > 0 and <= 86400s (1 day)
	// - The clientTtl cannot be larger than the defaultTtl (if set)
	// - Fractions of a second are not allowed.
	// Omit this field to use the defaultTtl, or the max-age set by the origin, as the client-facing TTL.
	// When the cache mode is set to "USE_ORIGIN_HEADERS" or "BYPASS_CACHE", you must omit this field.
	// A duration in seconds terminated by 's'. Example: "3s".
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.client_ttl
	ClientTtl *string `json:"clientTtl,omitempty"`

	// Optional. Specifies the default TTL for cached content served by this origin for responses that do not have an existing valid TTL (max-age or s-max-age).
	// Defaults to 3600s (1 hour).
	// - The TTL must be >= 0 and <= 31,536,000 seconds (1 year)
	// - Setting a TTL of "0" means "always revalidate" (equivalent to must-revalidate)
	// - The value of defaultTTL cannot be set to a value greater than that of maxTTL.
	// - Fractions of a second are not allowed.
	// - When the cacheMode is set to FORCE_CACHE_ALL, the defaultTTL will overwrite the TTL set in all responses.
	// Note that infrequently accessed objects may be evicted from the cache before the defined TTL. Objects that expire will be revalidated with the origin.
	// When the cache mode is set to "USE_ORIGIN_HEADERS" or "BYPASS_CACHE", you must omit this field.
	// A duration in seconds terminated by 's'. Example: "3s".
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.default_ttl
	DefaultTtl *string `json:"defaultTtl,omitempty"`

	// Optional. Specifies the maximum allowed TTL for cached content served by this origin.
	// Defaults to 86400s (1 day).
	// Cache directives that attempt to set a max-age or s-maxage higher than this, or an Expires header more than maxTtl seconds in the future will be capped at the value of maxTTL, as if it were the value of an s-maxage Cache-Control directive.
	// - The TTL must be >= 0 and <= 31,536,000 seconds (1 year)
	// - Setting a TTL of "0" means "always revalidate"
	// - The value of maxTtl must be equal to or greater than defaultTtl.
	// - Fractions of a second are not allowed.
	// When the cache mode is set to "USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", or "BYPASS_CACHE", you must omit this field.
	// A duration in seconds terminated by 's'. Example: "3s".
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.max_ttl
	MaxTtl *string `json:"maxTtl,omitempty"`

	// Optional. Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects. This can reduce the load on your origin and improve end-user experience by reducing response latency.
	// By default, the CDNPolicy will apply the following default TTLs to these status codes:
	// - HTTP 300 (Multiple Choice), 301, 308 (Permanent Redirects): 10m
	// - HTTP 404 (Not Found), 410 (Gone), 451 (Unavailable For Legal Reasons): 120s
	// - HTTP 405 (Method Not Found), 414 (URI Too Long), 501 (Not Implemented): 60s
	// These defaults can be overridden in negativeCachingPolicy.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.negative_caching
	NegativeCaching *bool `json:"negativeCaching,omitempty"`

	// Optional. Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.
	// - Omitting the policy and leaving negativeCaching enabled will use the default TTLs for each status code, defined in negativeCaching.
	// - TTLs must be >= 0 (where 0 is "always revalidate") and <= 86400s (1 day)
	// Note that when specifying an explicit negativeCachingPolicy, you should take care to specify a cache TTL for all response codes that you wish to cache. The CDNPolicy will not apply any default negative caching when a policy exists.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.negative_caching_policy
	NegativeCachingPolicy map[string]string `json:"negativeCachingPolicy,omitempty"`

	// Optional. The EdgeCacheKeyset containing the set of public keys used to validate signed requests at the edge.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.signed_request_keyset
	SignedRequestKeysetRef *ResourceRef `json:"signedRequestKeysetRef,omitempty"`

	// Optional. Limit how far into the future the expiration time of a signed request may be.
	// When set, a signed request is rejected if its expiration time is later than now + signedRequestMaximumExpirationTtl, where now is the time at which the signed request is first handled by the CDN.
	// - The TTL must be > 0.
	// - Fractions of a second are not allowed.
	// By default, signedRequestMaximumExpirationTtl is not set and the expiration time of a signed request may be arbitrarily far into future.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.signed_request_maximum_expiration_ttl
	SignedRequestMaximumExpirationTtl *string `json:"signedRequestMaximumExpirationTtl,omitempty"`

	// Optional. Whether to enforce signed requests. The default value is DISABLED, which means all content is public, and does not authorize access.
	// You must also set a signedRequestKeyset to enable signed requests.
	// When set to REQUIRE_SIGNATURES, all matching requests will have their signature validated. Requests that were not signed with the corresponding private key, or that are otherwise invalid (expired, do not match the signature, IP address, or header) will be rejected with a HTTP 403 and (if enabled) logged. Possible values: ["DISABLED", "REQUIRE_SIGNATURES", "REQUIRE_TOKENS"].
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.signed_request_mode
	SignedRequestMode *string `json:"signedRequestMode,omitempty"`

	// Optional. Additional options for signed tokens.
	// signedTokenOptions may only be specified when signedRequestMode is REQUIRE_TOKENS.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.signed_token_options
	SignedTokenOptions *EdgeCacheServiceSignedTokenOptions `json:"signedTokenOptions,omitempty"`
}

type EdgeCacheServiceAddSignatures struct {
	// Required. The actions to take to add signatures to responses. Possible values: ["GENERATE_COOKIE", "GENERATE_TOKEN_HLS_COOKIELESS", "PROPAGATE_TOKEN_HLS_COOKIELESS"].
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.AddSignatures.actions
	Actions []string `json:"actions"`

	// Optional. The parameters to copy from the verified token to the generated token.
	// Only the following parameters may be copied:
	// * 'PathGlobs'
	// * 'paths'
	// * 'acl'
	// * 'URLPrefix'
	// * 'IPRanges'
	// * 'SessionID'
	// * 'id'
	// * 'Data'
	// * 'data'
	// * 'payload'
	// * 'Headers'
	// You may specify up to 6 parameters to copy.  A given parameter is be copied only if the parameter exists in the verified token.  Parameter names are matched exactly as specified.  The order of the parameters does not matter.  Duplicates are not allowed.
	// This field may only be specified when the GENERATE_COOKIE or GENERATE_TOKEN_HLS_COOKIELESS actions are specified.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.AddSignatures.copied_parameters
	CopiedParameters []string `json:"copiedParameters,omitempty"`

	// Optional. The keyset to use for signature generation.
	// The following are both valid paths to an EdgeCacheKeyset resource:
	// * 'projects/project/locations/global/edgeCacheKeysets/yourKeyset'
	// * 'yourKeyset'
	// This must be specified when the GENERATE_COOKIE or GENERATE_TOKEN_HLS_COOKIELESS actions are specified.  This field may not be specified otherwise.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.AddSignatures.keyset
	KeysetRef *ResourceRef `json:"keysetRef,omitempty"`

	// Optional. The query parameter in which to put the generated token.
	// If not specified, defaults to 'edge-cache-token'.
	// If specified, the name must be 1-64 characters long and match the regular expression '[a-zA-Z]([a-zA-Z0-9_-])*' which means the first character must be a letter, and all following characters must be a dash, underscore, letter or digit.
	// This field may only be set when the GENERATE_TOKEN_HLS_COOKIELESS or PROPAGATE_TOKEN_HLS_COOKIELESS actions are specified.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.AddSignatures.token_query_parameter
	TokenQueryParameter *string `json:"tokenQueryParameter,omitempty"`

	// Optional. The duration the token is valid starting from the moment the token is first generated.
	// Defaults to '86400s' (1 day).
	// The TTL must be >= 0 and <= 604,800 seconds (1 week).
	// This field may only be specified when the GENERATE_COOKIE or GENERATE_TOKEN_HLS_COOKIELESS actions are specified.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.AddSignatures.token_ttl
	TokenTtl *string `json:"tokenTtl,omitempty"`
}

type EdgeCacheServiceCacheKeyPolicy struct {
	// Optional. If true, requests to different hosts will be cached separately.
	// Note: this should only be enabled if hosts share the same origin and content. Removing the host from the cache key may inadvertently result in different objects being cached than intended, depending on which route the first user matched.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.CacheKeyPolicy.exclude_host
	ExcludeHost *bool `json:"excludeHost,omitempty"`

	// Optional. If true, exclude query string parameters from the cache key
	// If false (the default), include the query string parameters in
	// the cache key according to includeQueryParameters and
	// excludeQueryParameters. If neither includeQueryParameters nor
	// excludeQueryParameters is set, the entire query string will be
	// included.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.CacheKeyPolicy.exclude_query_string
	ExcludeQueryString *bool `json:"excludeQueryString,omitempty"`

	// Optional. Names of query string parameters to exclude from cache keys. All other parameters will be included.
	// Either specify includedQueryParameters or excludedQueryParameters, not both. '&' and '=' will be percent encoded and not treated as delimiters.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.CacheKeyPolicy.excluded_query_parameters
	ExcludedQueryParameters []string `json:"excludedQueryParameters,omitempty"`

	// Optional. If true, http and https requests will be cached separately.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.CacheKeyPolicy.include_protocol
	IncludeProtocol *bool `json:"includeProtocol,omitempty"`

	// Optional. Names of Cookies to include in cache keys.  The cookie name and cookie value of each cookie named will be used as part of the cache key.
	// Cookie names:
	// - must be valid RFC 6265 "cookie-name" tokens
	// - are case sensitive
	// - cannot start with "Edge-Cache-" (case insensitive)
	// Note that specifying several cookies, and/or cookies that have a large range of values (e.g., per-user) will dramatically impact the cache hit rate, and may result in a higher eviction rate and reduced performance.
	// You may specify up to three cookie names.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.CacheKeyPolicy.included_cookie_names
	IncludedCookieNames []string `json:"includedCookieNames,omitempty"`

	// Optional. Names of HTTP request headers to include in cache keys. The value of the header field will be used as part of the cache key.
	// - Header names must be valid HTTP RFC 7230 header field values.
	// - Header field names are case insensitive
	// - To include the HTTP method, use ":method"
	// Note that specifying several headers, and/or headers that have a large range of values (e.g. per-user) will dramatically impact the cache hit rate, and may result in a higher eviction rate and reduced performance.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.CacheKeyPolicy.included_header_names
	IncludedHeaderNames []string `json:"includedHeaderNames,omitempty"`

	// Optional. Names of query string parameters to include in cache keys. All other parameters will be excluded.
	// Either specify includedQueryParameters or excludedQueryParameters, not both. '&' and '=' will be percent encoded and not treated as delimiters.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.CacheKeyPolicy.included_query_parameters
	IncludedQueryParameters []string `json:"includedQueryParameters,omitempty"`
}

type EdgeCacheServiceSignedTokenOptions struct {
	// Optional. The allowed signature algorithms to use.
	// Defaults to using only ED25519.
	// You may specify up to 3 signature algorithms to use. Possible values: ["ED25519", "HMAC_SHA_256", "HMAC_SHA1"].
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.SignedTokenOptions.allowed_signature_algorithms
	AllowedSignatureAlgorithms []string `json:"allowedSignatureAlgorithms,omitempty"`

	// Optional. The query parameter in which to find the token.
	// The name must be 1-64 characters long and match the regular expression '[a-zA-Z]([a-zA-Z0-9_-])*' which means the first character must be a letter, and all following characters must be a dash, underscore, letter or digit.
	// Defaults to 'edge-cache-token'.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CdnPolicy.SignedTokenOptions.token_query_parameter
	TokenQueryParameter *string `json:"tokenQueryParameter,omitempty"`
}

type EdgeCacheServiceCorsPolicy struct {
	// Optional. In response to a preflight request, setting this to true indicates that the actual request can include user credentials.
	// This translates to the Access-Control-Allow-Credentials response header.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CorsPolicy.allow_credentials
	AllowCredentials *bool `json:"allowCredentials,omitempty"`

	// Optional. Specifies the content for the Access-Control-Allow-Headers response header.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CorsPolicy.allow_headers
	AllowHeaders []string `json:"allowHeaders,omitempty"`

	// Optional. Specifies the content for the Access-Control-Allow-Methods response header.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CorsPolicy.allow_methods
	AllowMethods []string `json:"allowMethods,omitempty"`

	// Optional. Specifies the list of origins that will be allowed to do CORS requests.
	// This translates to the Access-Control-Allow-Origin response header.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CorsPolicy.allow_origins
	AllowOrigins []string `json:"allowOrigins,omitempty"`

	// Optional. If true, specifies the CORS policy is disabled. The default value is false, which indicates that the CORS policy is in effect.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CorsPolicy.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Optional. Specifies the content for the Access-Control-Allow-Headers response header.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CorsPolicy.expose_headers
	ExposeHeaders []string `json:"exposeHeaders,omitempty"`

	// Required. Specifies how long results of a preflight request can be cached by a client in seconds. Note that many browser clients enforce a maximum TTL of 600s (10 minutes).
	// - Setting the value to -1 forces a pre-flight check for all requests (not recommended)
	// - A maximum TTL of 86400s can be set, but note that (as above) some clients may force pre-flight checks at a more regular interval.
	// - This translates to the Access-Control-Max-Age header.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.CorsPolicy.max_age
	MaxAge string `json:"maxAge"`
}

type EdgeCacheServiceUrlRewrite struct {
	// Optional. Prior to forwarding the request to the selected origin, the request's host header is replaced with contents of hostRewrite.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.UrlRewrite.host_rewrite
	HostRewrite *string `json:"hostRewrite,omitempty"`

	// Optional. Prior to forwarding the request to the selected origin, the matching portion of the request's path is replaced by pathPrefixRewrite.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.UrlRewrite.path_prefix_rewrite
	PathPrefixRewrite *string `json:"pathPrefixRewrite,omitempty"`

	// Optional. Prior to forwarding the request to the selected origin, if the
	// request matched a pathTemplateMatch, the matching portion of the
	// request's path is replaced re-written using the pattern specified
	// by pathTemplateRewrite.
	// pathTemplateRewrite must be between 1 and 255 characters
	// (inclusive), must start with a '/', and must only use variables
	// captured by the route's pathTemplate matchers.
	// pathTemplateRewrite may only be used when all of a route's
	// MatchRules specify pathTemplate.
	// Only one of pathPrefixRewrite and pathTemplateRewrite may be
	// specified.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.RouteAction.UrlRewrite.path_template_rewrite
	PathTemplateRewrite *string `json:"pathTemplateRewrite,omitempty"`
}

type EdgeCacheServiceUrlRedirect struct {
	// Optional. The host that will be used in the redirect response instead of the one that was supplied in the request.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.UrlRedirect.host_redirect
	HostRedirect *string `json:"hostRedirect,omitempty"`

	// Optional. If set to true, the URL scheme in the redirected request is set to https. If set to false, the URL scheme of the redirected request will remain the same as that of the request.
	// This can only be set if there is at least one (1) edgeSslCertificate set on the service.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.UrlRedirect.https_redirect
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// Optional. The path that will be used in the redirect response instead of the one that was supplied in the request.
	// pathRedirect cannot be supplied together with prefixRedirect. Supply one alone or neither. If neither is supplied, the path of the original request will be used for the redirect.
	// The path value must be between 1 and 1024 characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.UrlRedirect.path_redirect
	PathRedirect *string `json:"pathRedirect,omitempty"`

	// Optional. The prefix that replaces the prefixMatch specified in the routeRule, retaining the remaining portion of the URL before redirecting the request.
	// prefixRedirect cannot be supplied together with pathRedirect. Supply one alone or neither. If neither is supplied, the path of the original request will be used for the redirect.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.UrlRedirect.prefix_redirect
	PrefixRedirect *string `json:"prefixRedirect,omitempty"`

	// Optional. The HTTP Status code to use for this RedirectAction.
	// The supported values are:
	// - 'MOVED_PERMANENTLY_DEFAULT', which is the default value and corresponds to 301.
	// - 'FOUND', which corresponds to 302.
	// - 'SEE_OTHER' which corresponds to 303.
	// - 'TEMPORARY_REDIRECT', which corresponds to 307. in this case, the request method will be retained.
	// - 'PERMANENT_REDIRECT', which corresponds to 308. in this case, the request method will be retained. Possible values: ["MOVED_PERMANENTLY_DEFAULT", "FOUND", "SEE_OTHER", "TEMPORARY_REDIRECT", "PERMANENT_REDIRECT"].
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.UrlRedirect.redirect_response_code
	RedirectResponseCode *string `json:"redirectResponseCode,omitempty"`

	// Optional. If set to true, any accompanying query portion of the original URL is removed prior to redirecting the request. If set to false, the query portion of the original URL is retained.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.Routing.RouteRule.UrlRedirect.strip_query
	StripQuery *bool `json:"stripQuery,omitempty"`
}

// NetworkServicesEdgeCacheServiceStatus defines the config connector machine state of NetworkServicesEdgeCacheService
type NetworkServicesEdgeCacheServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []k8sv1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkServicesEdgeCacheService resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkServicesEdgeCacheServiceObservedState `json:"observedState,omitempty"`
}

// NetworkServicesEdgeCacheServiceObservedState is the state of the NetworkServicesEdgeCacheService resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networkservices.v1.EdgeCacheService
type NetworkServicesEdgeCacheServiceObservedState struct {
	// The IPv4 addresses associated with this service. Addresses are static for the lifetime of the service.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.ipv4_addresses
	Ipv4Addresses []string `json:"ipv4Addresses,omitempty"`

	// The IPv6 addresses associated with this service. Addresses are static for the lifetime of the service.
	// +kcc:proto:field=google.cloud.networkservices.v1.EdgeCacheService.ipv6_addresses
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkservicesedgecacheservice;gcpnetworkservicesedgecacheservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkServicesEdgeCacheService is the Schema for the NetworkServicesEdgeCacheService API
// +k8s:openapi-gen=true
type NetworkServicesEdgeCacheService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkServicesEdgeCacheServiceSpec   `json:"spec,omitempty"`
	Status NetworkServicesEdgeCacheServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkServicesEdgeCacheServiceList contains a list of NetworkServicesEdgeCacheService
type NetworkServicesEdgeCacheServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesEdgeCacheService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesEdgeCacheService{}, &NetworkServicesEdgeCacheServiceList{})
}
