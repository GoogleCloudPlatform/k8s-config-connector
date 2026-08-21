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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// The health check resources for health checking this ComputeBackendService. Currently at most one health check can be specified, and a health check is required.
type BackendserviceHealthChecks struct {
	// Reference to a ComputeHealthCheck resource.
	HealthCheckRef *ComputeHealthCheckRef `json:"healthCheckRef,omitempty"`
	// Reference to a ComputeHTTPHealthCheck resource.
	HttpHealthCheckRef *ComputeHTTPHealthCheckRef `json:"httpHealthCheckRef,omitempty"`
}

// Reference to a ComputeInstanceGroup or ComputeNetworkEndpointGroup resource. In case of instance group this defines the list of instances that serve traffic. Member virtual machine instances from each instance group must live in the same zone as the instance group itself. No two backends in a backend service are allowed to use same Instance Group resource. For Network Endpoint Groups this defines list of endpoints. All endpoints of Network Endpoint Group must be hosted on instances located in the same zone as the Network Endpoint Group. Backend services cannot mix Instance Group and Network Endpoint Group backends. When the 'load_balancing_scheme' is INTERNAL, only instance groups are supported.
type BackendGroup struct {
	// Reference to a ComputeInstanceGroup resource.
	InstanceGroupRef *ComputeInstanceGroupRef `json:"instanceGroupRef,omitempty"`

	// Reference to a ComputeNetworkEndpointGroup resource.
	NetworkEndpointGroupRef *ComputeNetworkEndpointGroupRef `json:"networkEndpointGroupRef,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.Backend
type Backend struct {
	// Specifies the balancing mode for this backend.
	// +kcc:proto:field=google.cloud.compute.v1.Backend.balancing_mode
	BalancingMode *string `json:"balancingMode,omitempty"`

	// A multiplier applied to the group's maximum servicing capacity.
	// +kcc:proto:field=google.cloud.compute.v1.Backend.capacity_scaler
	CapacityScaler *float32 `json:"capacityScaler,omitempty"`

	// An optional description of this resource.
	// +kcc:proto:field=google.cloud.compute.v1.Backend.description
	Description *string `json:"description,omitempty"`

	// This field designates whether this is a failover backend.
	// +kcc:proto:field=google.cloud.compute.v1.Backend.failover
	Failover *bool `json:"failover,omitempty"`

	// Reference to a ComputeInstanceGroup or ComputeNetworkEndpointGroup
	//  resource. In case of instance group this defines the list of
	//  instances that serve traffic. Member virtual machine instances from
	//  each instance group must live in the same zone as the instance
	//  group itself. No two backends in a backend service are allowed to
	//  use same Instance Group resource.
	//
	//  For Network Endpoint Groups this defines list of endpoints. All
	//  endpoints of Network Endpoint Group must be hosted on instances
	//  located in the same zone as the Network Endpoint Group.
	//
	//  Backend services cannot mix Instance Group and Network Endpoint
	//  Group backends.
	//
	//  When the 'load_balancing_scheme' is INTERNAL, only instance groups
	//  are supported.
	// +required
	Group *BackendGroup `json:"group"`

	// The max number of simultaneous connections for the group.
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_connections
	MaxConnections *int32 `json:"maxConnections,omitempty"`

	// The max number of simultaneous connections that a single backend network endpoint can handle.
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_connections_per_endpoint
	MaxConnectionsPerEndpoint *int32 `json:"maxConnectionsPerEndpoint,omitempty"`

	// The max number of simultaneous connections that a single backend instance can handle.
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_connections_per_instance
	MaxConnectionsPerInstance *int32 `json:"maxConnectionsPerInstance,omitempty"`

	// The max requests per second (RPS) of the group.
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_rate
	MaxRate *int32 `json:"maxRate,omitempty"`

	// The max requests per second (RPS) that a single backend network endpoint can handle.
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_rate_per_endpoint
	MaxRatePerEndpoint *float32 `json:"maxRatePerEndpoint,omitempty"`

	// The max requests per second (RPS) that a single backend instance can handle.
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_rate_per_instance
	MaxRatePerInstance *float32 `json:"maxRatePerInstance,omitempty"`

	// Used when balancingMode is UTILIZATION. This ratio defines the CPU utilization target for the group.
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_utilization
	MaxUtilization *float32 `json:"maxUtilization,omitempty"`
}

type BackendServiceDuration struct {
	// Span of time at a resolution of a second. Must be from 0 to 315,576,000,000 inclusive.
	// +required
	Seconds int64 `json:"seconds"`
	// Span of time that's a fraction of a second at nanosecond resolution. Must be from 0 to 999,999,999 inclusive.
	Nanos *int32 `json:"nanos,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.CircuitBreakers
type CircuitBreakers struct {
	// Not supported by proxy. Connect timeout for the backend service. The timeout is the maximum amount of time KCC will wait for the proxy to connect to the backend.
	ConnectTimeout *BackendServiceDuration `json:"connectTimeout,omitempty"`

	// The maximum number of connections to the backend service. If not specified, there is no limit. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.CircuitBreakers.max_connections
	MaxConnections *int32 `json:"maxConnections,omitempty"`

	// The maximum number of pending requests allowed to the backend service. If not specified, there is no limit. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.CircuitBreakers.max_pending_requests
	MaxPendingRequests *int32 `json:"maxPendingRequests,omitempty"`

	// The maximum number of parallel requests that allowed to the backend service. If not specified, there is no limit.
	// +kcc:proto:field=google.cloud.compute.v1.CircuitBreakers.max_requests
	MaxRequests *int32 `json:"maxRequests,omitempty"`

	// Maximum requests for a single connection to the backend service. This parameter is respected by both the HTTP/1.1 and HTTP/2 implementations. If not specified, there is no limit. Setting this parameter to 1 will effectively disable keep alive. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.CircuitBreakers.max_requests_per_connection
	MaxRequestsPerConnection *int32 `json:"maxRequestsPerConnection,omitempty"`

	// The maximum number of parallel retries allowed to the backend cluster. If not specified, the default is 1. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.CircuitBreakers.max_retries
	MaxRetries *int32 `json:"maxRetries,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceHttpCookie
type BackendServiceHttpCookie struct {
	// Name of the cookie.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceHttpCookie.name
	Name *string `json:"name,omitempty"`

	// Path to set for the cookie.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceHttpCookie.path
	Path *string `json:"path,omitempty"`

	// Lifetime of the cookie.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceHttpCookie.ttl
	Ttl *BackendServiceDuration `json:"ttl,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ConsistentHashLoadBalancerSettings
type ConsistentHashLoadBalancerSettings struct {
	// Hash session affinity settings when using STRONG_COOKIE_AFFINITY or STRONG_HEADER_AFFINITY.
	// +kcc:proto:field=google.cloud.compute.v1.ConsistentHashLoadBalancerSettings.http_cookie
	HttpCookie *BackendServiceHttpCookie `json:"httpCookie,omitempty"`

	// The name of the HTTP header field to map onto. Required if the load balancing policy is MAGLEV or RING_HASH and session affinity is STRONG_HEADER_AFFINITY.
	// +kcc:proto:field=google.cloud.compute.v1.ConsistentHashLoadBalancerSettings.http_header_name
	HttpHeaderName *string `json:"httpHeaderName,omitempty"`

	// The minimum ring size. Valid values are between 1 and 4096. Required if the load balancing policy is RING_HASH.
	// +kcc:proto:field=google.cloud.compute.v1.ConsistentHashLoadBalancerSettings.minimum_ring_size
	MinimumRingSize *int64 `json:"minimumRingSize,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.OutlierDetection
type OutlierDetection struct {
	// The base time that a host is ejected for. The real time is equal to the base ejection time multiplied by the number of times the host has been ejected.
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.base_ejection_time
	BaseEjectionTime *BackendServiceDuration `json:"baseEjectionTime,omitempty"`

	// Number of consecutive errors before ejection.
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.consecutive_errors
	ConsecutiveErrors *int32 `json:"consecutiveErrors,omitempty"`

	// Number of consecutive gateway failures before ejection.
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.consecutive_gateway_failure
	ConsecutiveGatewayFailure *int32 `json:"consecutiveGatewayFailure,omitempty"`

	// Percentage of consecutive errors to enforce.
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.enforcing_consecutive_errors
	EnforcingConsecutiveErrors *int32 `json:"enforcingConsecutiveErrors,omitempty"`

	// Percentage of consecutive gateway failures to enforce.
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.enforcing_consecutive_gateway_failure
	EnforcingConsecutiveGatewayFailure *int32 `json:"enforcingConsecutiveGatewayFailure,omitempty"`

	// Percentage of success rate to enforce.
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.enforcing_success_rate
	EnforcingSuccessRate *int32 `json:"enforcingSuccessRate,omitempty"`

	// Time interval between ejection sweep analysis.
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.interval
	Interval *BackendServiceDuration `json:"interval,omitempty"`

	// Maximum percentage of hosts that can be ejected.
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.max_ejection_percent
	MaxEjectionPercent *int32 `json:"maxEjectionPercent,omitempty"`

	// Minimum number of hosts for success rate analysis.
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.success_rate_minimum_hosts
	SuccessRateMinimumHosts *int32 `json:"successRateMinimumHosts,omitempty"`

	// Minimum request volume for success rate analysis.
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.success_rate_request_volume
	SuccessRateRequestVolume *int32 `json:"successRateRequestVolume,omitempty"`

	// Standard deviation factor for success rate analysis.
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.success_rate_stdev_factor
	SuccessRateStdevFactor *int32 `json:"successRateStdevFactor,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.CacheKeyPolicy
type CacheKeyPolicy struct {
	// If true, requests to different hosts will be cached separately.
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.include_host
	IncludeHost *bool `json:"includeHost,omitempty"`

	// Allows HTTP request headers (by name) to be used in the cache key.
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.include_http_headers
	IncludeHttpHeaders []string `json:"includeHttpHeaders,omitempty"`

	// Allows HTTP cookies (by name) to be used in the cache key. The name=value pair will be used in the cache key Cloud CDN generates.
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.include_named_cookies
	IncludeNamedCookies []string `json:"includeNamedCookies,omitempty"`

	// If true, http and https requests will be cached separately.
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.include_protocol
	IncludeProtocol *bool `json:"includeProtocol,omitempty"`

	// If true, include query string parameters in the cache key according to query_string_whitelist and query_string_blacklist. If neither is set, the entire query string will be included. If false, the query string will be excluded from the cache key entirely.
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.include_query_string
	IncludeQueryString *bool `json:"includeQueryString,omitempty"`

	// Names of query string parameters to exclude in cache keys. All other parameters will be included. Either specify query_string_whitelist or query_string_blacklist, not both. '&' and '=' will be percent encoded and not treated as delimiters.
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.query_string_blacklist
	QueryStringBlacklist []string `json:"queryStringBlacklist,omitempty"`

	// Names of query string parameters to include in cache keys. All other parameters will be excluded. Either specify query_string_whitelist or query_string_blacklist, not both. '&' and '=' will be percent encoded and not treated as delimiters.
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.query_string_whitelist
	QueryStringWhitelist []string `json:"queryStringWhitelist,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceCdnPolicyBypassCacheOnRequestHeader
type BackendServiceCdnPolicyBypassCacheOnRequestHeader struct {
	// The header field name to match on when bypassing cache. Values are case-insensitive.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicyBypassCacheOnRequestHeader.header_name
	// +required
	HeaderName string `json:"headerName"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceCdnPolicyNegativeCachingPolicy
type BackendServiceCdnPolicyNegativeCachingPolicy struct {
	// The HTTP status code to define a TTL against. Only HTTP status codes 300, 301, 302, 307, 308, 404, 405, 410, 421, 451 and 501 are can be specified as values, and you cannot specify a status code more than once.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicyNegativeCachingPolicy.code
	Code *int32 `json:"code,omitempty"`

	// The TTL (in seconds) for which to cache responses with the corresponding status code. The maximum allowed value is 1800s (30 minutes), noting that infrequently accessed objects may be evicted from the cache before the defined TTL.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicyNegativeCachingPolicy.ttl
	Ttl *int32 `json:"ttl,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceCdnPolicy
type BackendServiceCdnPolicy struct {
	// Bypass the cache when the specified request headers are matched - e.g. Pragma or Authorization headers. Up to 5 headers can be specified. The cache is bypassed for all cdnPolicy.cacheMode settings.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.bypass_cache_on_request_headers
	BypassCacheOnRequestHeaders []BackendServiceCdnPolicyBypassCacheOnRequestHeader `json:"bypassCacheOnRequestHeaders,omitempty"`

	// The CacheKeyPolicy for this CdnPolicy.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.cache_key_policy
	CacheKeyPolicy *CacheKeyPolicy `json:"cacheKeyPolicy,omitempty"`

	// Specifies the cache setting for all responses from this backend. The possible values are: USE_ORIGIN_HEADERS Requires the origin to set valid caching headers to cache content. Responses without these headers will not be cached at Google's edge, and will require a full trip to the origin on every request, potentially impacting performance and increasing load on the origin server. FORCE_CACHE_ALL Cache all content, ignoring any "private", "no-store" or "no-cache" directives in Cache-Control response headers. Warning: this may result in Cloud CDN caching private, per-user (user identifiable) content. CACHE_ALL_STATIC Automatically cache static content, including common image formats, media (video and audio), and web assets (JavaScript and CSS). Requests and responses that are marked as uncacheable, as well as dynamic content (including HTML), will not be cached. If no value is provided for cdnPolicy.cacheMode, it defaults to CACHE_ALL_STATIC.
	//  Check the CacheMode enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.cache_mode
	CacheMode *string `json:"cacheMode,omitempty"`

	// Specifies a separate client (e.g. browser client) maximum TTL. This is used to clamp the max-age (or Expires) value sent to the client. With FORCE_CACHE_ALL, the lesser of client_ttl and default_ttl is used for the response max-age directive, along with a "public" directive. For cacheable content in CACHE_ALL_STATIC mode, client_ttl clamps the max-age from the origin (if specified), or else sets the response max-age directive to the lesser of the client_ttl and default_ttl, and also ensures a "public" cache-control directive is present. If a client TTL is not specified, a default value (1 hour) will be used. The maximum allowed value is 31,622,400s (1 year).
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.client_ttl
	ClientTtl *int32 `json:"clientTtl,omitempty"`

	// Specifies the default TTL for cached content served by this origin for responses that do not have an existing valid TTL (max-age or s-maxage). Setting a TTL of "0" means "always revalidate". The value of defaultTTL cannot be set to a value greater than that of maxTTL, but can be equal. When the cacheMode is set to FORCE_CACHE_ALL, the defaultTTL will overwrite the TTL set in all responses. The maximum allowed value is 31,622,400s (1 year), noting that infrequently accessed objects may be evicted from the cache before the defined TTL.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.default_ttl
	DefaultTtl *int32 `json:"defaultTtl,omitempty"`

	// Specifies the maximum allowed TTL for cached content served by this origin. Cache directives that attempt to set a max-age or s-maxage higher than this, or an Expires header more than maxTTL seconds in the future will be capped at the value of maxTTL, as if it were the value of an s-maxage Cache-Control directive. Headers sent to the client will not be modified. Setting a TTL of "0" means "always revalidate". The maximum allowed value is 31,622,400s (1 year), noting that infrequently accessed objects may be evicted from the cache before the defined TTL.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.max_ttl
	MaxTtl *int32 `json:"maxTtl,omitempty"`

	// Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects. This can reduce the load on your origin and improve end-user experience by reducing response latency. When the cache mode is set to CACHE_ALL_STATIC or USE_ORIGIN_HEADERS, negative caching applies to responses with the specified response code that lack any Cache-Control, Expires, or Pragma: no-cache directives. When the cache mode is set to FORCE_CACHE_ALL, negative caching applies to all responses with the specified response code, and override any caching headers. By default, Cloud CDN will apply the following default TTLs to these status codes: HTTP 300 (Multiple Choice), 301, 308 (Permanent Redirects): 10m HTTP 404 (Not Found), 410 (Gone), 451 (Unavailable For Legal Reasons): 120s HTTP 405 (Method Not Found), 421 (Misdirected Request), 501 (Not Implemented): 60s. These defaults can be overridden in negative_caching_policy.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.negative_caching
	NegativeCaching *bool `json:"negativeCaching,omitempty"`

	// Sets a cache TTL for the specified HTTP status code. negative_caching must be enabled to configure negative_caching_policy. Omitting the policy and leaving negative_caching enabled will use Cloud CDN's default cache TTLs. Note that when specifying an explicit negative_caching_policy, you should take care to specify a cache TTL for all response codes that you wish to cache. Cloud CDN will not apply any default negative caching when a policy exists.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.negative_caching_policy
	NegativeCachingPolicy []BackendServiceCdnPolicyNegativeCachingPolicy `json:"negativeCachingPolicy,omitempty"`

	// Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache. This setting defines the default "max-stale" duration for any cached responses that do not specify a max-stale directive. Stale responses that exceed the TTL configured here will not be served. The default limit (max-stale) is 86400s (1 day), which will allow stale content to be served up to this limit beyond the max-age (or s-maxage) of a cached response. The maximum allowed value is 604800 (1 week). Set this to zero (0) to disable serve-while-stale.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.serve_while_stale
	ServeWhileStale *int32 `json:"serveWhileStale,omitempty"`

	// Maximum number of seconds the response to a signed URL request will be considered fresh. After this time period, the response will be revalidated before being served. Defaults to 1hr (3600s). When serving responses to signed URL requests, Cloud CDN will internally behave as though all responses from this backend had a "Cache-Control: public, max-age=[TTL]" header, regardless of any existing Cache-Control header. The actual headers served in responses will not be altered.
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.signed_url_cache_max_age_sec
	SignedUrlCacheMaxAgeSec *int64 `json:"signedUrlCacheMaxAgeSec,omitempty"`
}

// BackendServiceOauth2ClientIdRef is a hand-coded reference. Since there is no KCC resource managing Google OAuth2 Client IDs, this custom reference allows referencing an external client ID.
type BackendServiceOauth2ClientIdRef struct {
	// Name of the referent.
	Name string `json:"name,omitempty"`
	// Namespace of the referent.
	Namespace string `json:"namespace,omitempty"`
	// External is the Google OAuth2 Client ID, e.g. `CLIENT_ID.apps.googleusercontent.com` matching the Cloud Asset Inventory Format.
	External string `json:"external,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceIap
type BackendServiceIap struct {
	// OAuth2 client ID to use for the authentication flow.
	Oauth2ClientId *string `json:"oauth2ClientId,omitempty"`
	// Reference to a ComputeBackendServiceOauth2ClientId.
	Oauth2ClientIdRef *BackendServiceOauth2ClientIdRef `json:"oauth2ClientIdRef,omitempty"`
	// OAuth2 client secret to use for the authentication flow.
	Oauth2ClientSecret *secret.Legacy `json:"oauth2ClientSecret,omitempty"`
	// [Output Only] SHA256 hash value for the field oauth2_client_secret above.
	Oauth2ClientSecretSha256 *string `json:"oauth2ClientSecretSha256,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy
type BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy struct {
	// An optional, arbitrary JSON object with configuration data.
	Data *string `json:"data,omitempty"`
	// The name of a custom policy.
	// +required
	Name string `json:"name"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceLocalityLoadBalancingPolicyConfigPolicy
type BackendServiceLocalityLoadBalancingPolicyConfigPolicy struct {
	// The name of a locality load balancing policy.
	// +required
	Name string `json:"name"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceLocalityLoadBalancingPolicyConfig
type BackendServiceLocalityLoadBalancingPolicyConfig struct {
	// The configuration for a custom policy.
	CustomPolicy *BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy `json:"customPolicy,omitempty"`
	// The configuration for a predefined policy.
	Policy *BackendServiceLocalityLoadBalancingPolicyConfigPolicy `json:"policy,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceLogConfig
type BackendServiceLogConfig struct {
	// This field denotes whether to enable logging for the load balancer traffic.
	Enable *bool `json:"enable,omitempty"`
	// This field defines the log sampling rate.
	SampleRate *float64 `json:"sampleRate,omitempty"`
}

type ComputeSecuritySettingsClientTLSPolicyRef struct {
	// Name of the referent.
	Name string `json:"name,omitempty"`
	// Namespace of the referent.
	Namespace string `json:"namespace,omitempty"`
	// External is the selfLink of the referent.
	External string `json:"external,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SecuritySettings
type SecuritySettings struct {
	// Reference to a ComputeSecuritySettingsClientTLSPolicy.
	// +required
	ClientTLSPolicyRef *ComputeSecuritySettingsClientTLSPolicyRef `json:"clientTLSPolicyRef"`
	// Subject Alternative Names.
	// +required
	SubjectAltNames []string `json:"subjectAltNames"`
}

// +kcc:proto=google.cloud.compute.v1.Subsetting
type Subsetting struct {
	// The subsetting policy. The only supported policy is CONSISTENT_HASH_SUBSETTING.
	// +kcc:proto:field=google.cloud.compute.v1.Subsetting.policy
	// +required
	Policy string `json:"policy"`
}

// ComputeBackendServiceSpec defines the desired state of ComputeBackendService
// +kcc:spec:proto=google.cloud.compute.v1.BackendService
type ComputeBackendServiceSpec struct {
	// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.affinity_cookie_ttl_sec
	AffinityCookieTtlSec *int32 `json:"affinityCookieTtlSec,omitempty"`

	// The set of backends that serve this BackendService.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.backends
	Backend []Backend `json:"backend,omitempty"`

	// Cloud CDN configuration for this BackendService.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.cdn_policy
	CDNPolicy *BackendServiceCdnPolicy `json:"cdnPolicy,omitempty"`

	// Settings controlling the circuit breaking behavior for this BackendService.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.circuit_breakers
	CircuitBreakers *CircuitBreakers `json:"circuitBreakers,omitempty"`

	// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.compression_mode
	CompressionMode *string `json:"compressionMode,omitempty"`

	// Time for which instance will be drained (not accept new connections, but still work to finish started).
	ConnectionDrainingTimeoutSec *int32 `json:"connectionDrainingTimeoutSec,omitempty"`

	// Connection Tracking configuration for this BackendService. Connection tracking policy settings are only available for external passthrough Network Load Balancers and internal passthrough Network Load Balancers. connectionTrackingPolicy cannot be specified with haPolicy.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.connection_tracking_policy
	ConnectionTrackingPolicy *BackendServiceConnectionTrackingPolicy `json:"connectionTrackingPolicy,omitempty"`

	// Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, HTTP2 or H2C, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.consistent_hash
	ConsistentHash *ConsistentHashLoadBalancerSettings `json:"consistentHash,omitempty"`

	// Headers that the load balancer adds to proxied requests. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.custom_request_headers
	CustomRequestHeaders []string `json:"customRequestHeaders,omitempty"`

	// Headers that the load balancer adds to proxied responses. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.custom_response_headers
	CustomResponseHeaders []string `json:"customResponseHeaders,omitempty"`

	// An optional description of this resource.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.description
	Description *string `json:"description,omitempty"`

	EdgeSecurityPolicyRef *ComputeSecurityPolicyRef `json:"edgeSecurityPolicyRef,omitempty"`

	// If true, enables Cloud CDN for the backend service.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.enable_c_d_n
	EnableCdn *bool `json:"enableCdn,omitempty"`

	// Requires at least one backend instance group to be defined as a backup (failover) backend. For load balancers that have configurable failover: [Internal passthrough Network Load Balancers](https://cloud.google.com/load-balancing/docs/internal/failover-overview) and [external passthrough Network Load Balancers](https://cloud.google.com/load-balancing/docs/network/networklb-failover-overview). failoverPolicy cannot be specified with haPolicy.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.failover_policy
	FailoverPolicy *BackendServiceFailoverPolicy `json:"failoverPolicy,omitempty"`

	HealthChecks []BackendserviceHealthChecks `json:"healthChecks,omitempty"`

	// The configurations for Identity-Aware Proxy on this resource.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.iap
	Iap *BackendServiceIap `json:"iap,omitempty"`

	// Specifies the load balancer type.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.load_balancing_scheme
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`

	// A list of locality load-balancing policies to be used in order of preference. When you use localityLbPolicies, you must set at least one value for either the localityLbPolicies[].policy or the localityLbPolicies[].customPolicy field. localityLbPolicies overrides any value set in the localityLbPolicy field. For an example of how to use this field, see Define a list of preferred policies. Caution: This field and its children are intended for use in a service mesh that includes gRPC clients only. Envoy proxies can't use backend services that have this configuration.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.locality_lb_policies
	LocalityLbPolicies []BackendServiceLocalityLoadBalancingPolicyConfig `json:"localityLbPolicies,omitempty"`

	// The load balancing algorithm used within the scope of the locality.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.locality_lb_policy
	LocalityLbPolicy *string `json:"localityLbPolicy,omitempty"`

	// Location of the resource.
	Location string `json:"location"`

	// This field denotes the logging options for the load balancer traffic.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.log_config
	LogConfig *BackendServiceLogConfig `json:"logConfig,omitempty"`

	NetworkRef *ComputeNetworkRef `json:"networkRef,omitempty"`

	// Settings controlling the ejection of unhealthy backend endpoints.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.outlier_detection
	OutlierDetection *OutlierDetection `json:"outlierDetection,omitempty"`

	// A named port on a backend instance group representing the port for communication to the backend VMs in that group. The named port must be [defined on each backend instance group](https://cloud.google.com/load-balancing/docs/backend-service#named_ports). This parameter has no meaning if the backends are NEGs. For internal passthrough Network Load Balancers and external passthrough Network Load Balancers, omit port_name.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.port_name
	PortName *string `json:"portName,omitempty"`

	// The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, H2C, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancers or for Traffic Director for more information. Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.protocol
	Protocol *string `json:"protocol,omitempty"`

	// The ComputeBackendService name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The URL of the security policy associated with this backend service.
	SecurityPolicy *string `json:"securityPolicy,omitempty"`

	SecurityPolicyRef *ComputeSecurityPolicyRef `json:"securityPolicyRef,omitempty"`

	// This field specifies the security settings that apply to this backend service. This field is applicable to a global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.security_settings
	SecuritySettings *SecuritySettings `json:"securitySettings,omitempty"`

	// Type of session affinity to use. The default is NONE. Only NONE and HEADER_FIELD are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. For more details, see: [Session Affinity](https://cloud.google.com/load-balancing/docs/backend-service#session_affinity). sessionAffinity cannot be specified with haPolicy.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.session_affinity
	SessionAffinity *string `json:"sessionAffinity,omitempty"`

	// subsetting cannot be specified with haPolicy.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.subsetting
	Subsetting *Subsetting `json:"subsetting,omitempty"`

	// The backend service timeout.
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.timeout_sec
	TimeoutSec *int32 `json:"timeoutSec,omitempty"`
}

// ComputeBackendServiceStatus defines the config connector machine state of ComputeBackendService
type ComputeBackendServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Fingerprint of this resource. A hash of the contents stored in this object.
	Fingerprint *string `json:"fingerprint,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	GeneratedId *int64 `json:"generatedId,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputebackendservice;gcpcomputebackendservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeBackendService is the Schema for the ComputeBackendService API
// +k8s:openapi-gen=true
type ComputeBackendService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired state of the ComputeBackendService.
	// +required
	Spec ComputeBackendServiceSpec `json:"spec,omitempty"`

	// The observed state of the ComputeBackendService.
	Status ComputeBackendServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeBackendServiceList contains a list of ComputeBackendService
type ComputeBackendServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeBackendService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeBackendService{}, &ComputeBackendServiceList{})
}
