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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kcc:proto=google.cloud.compute.v1.BackendServiceCdnPolicyBypassCacheOnRequestHeader
type BackendserviceBypassCacheOnRequestHeaders struct {
	/* The header field name to match on when bypassing cache. Values are case-insensitive. */
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicyBypassCacheOnRequestHeader.header_name
	HeaderName *string `json:"headerName"`
}

// +kcc:proto=google.cloud.compute.v1.CacheKeyPolicy
type BackendserviceCacheKeyPolicy struct {
	/* If true requests to different hosts will be cached separately. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.include_host
	IncludeHost *bool `json:"includeHost,omitempty"`

	/* Allows HTTP request headers (by name) to be used in the cache key. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.include_http_headers
	IncludeHttpHeaders []string `json:"includeHttpHeaders,omitempty"`

	/* Names of cookies to include in cache keys. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.include_named_cookies
	IncludeNamedCookies []string `json:"includeNamedCookies,omitempty"`

	/* If true, http and https requests will be cached separately. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.include_protocol
	IncludeProtocol *bool `json:"includeProtocol,omitempty"`

	/* If true, include query string parameters in the cache key according to query_string_whitelist and query_string_blacklist. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.include_query_string
	IncludeQueryString *bool `json:"includeQueryString,omitempty"`

	/* Names of query string parameters to exclude in cache keys. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.query_string_blacklist
	QueryStringBlacklist []string `json:"queryStringBlacklist,omitempty"`

	/* Names of query string parameters to include in cache keys. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CacheKeyPolicy.query_string_whitelist
	QueryStringWhitelist []string `json:"queryStringWhitelist,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceCdnPolicyNegativeCachingPolicy
type BackendserviceNegativeCachingPolicy struct {
	/* The HTTP status code to define a TTL against. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicyNegativeCachingPolicy.code
	Code *int32 `json:"code,omitempty"`

	/* The TTL (in seconds) for which to cache responses with the corresponding status code. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicyNegativeCachingPolicy.ttl
	Ttl *BackendserviceTtl `json:"ttl,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceCdnPolicy
type BackendserviceCdnPolicy struct {
	/* Bypass the cache when the specified request headers are matched. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.bypass_cache_on_request_headers
	BypassCacheOnRequestHeaders []BackendserviceBypassCacheOnRequestHeaders `json:"bypassCacheOnRequestHeaders,omitempty"`

	/* The CacheKeyPolicy for this CdnPolicy. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.cache_key_policy
	CacheKeyPolicy *BackendserviceCacheKeyPolicy `json:"cacheKeyPolicy,omitempty"`

	/* Specifies the cache setting for all responses from this backend. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.cache_mode
	CacheMode *string `json:"cacheMode,omitempty"`

	/* Specifies the maximum allowed TTL for cached content served by this origin. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.client_ttl
	ClientTtl *int32 `json:"clientTtl,omitempty"`

	/* Specifies the default TTL for cached content served by this origin for responses. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.default_ttl
	DefaultTtl *int32 `json:"defaultTtl,omitempty"`

	/* Specifies the maximum allowed TTL for cached content served by this origin. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.max_ttl
	MaxTtl *int32 `json:"maxTtl,omitempty"`

	/* Negative caching allows per-status code TTLs to be set. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.negative_caching
	NegativeCaching *bool `json:"negativeCaching,omitempty"`

	/* Sets a cache TTL for the specified HTTP status code. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.negative_caching_policy
	NegativeCachingPolicy []BackendserviceNegativeCachingPolicy `json:"negativeCachingPolicy,omitempty"`

	/* Serve existing content from the cache (if available) when revalidating. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.serve_while_stale
	ServeWhileStale *int32 `json:"serveWhileStale,omitempty"`

	/* Maximum number of seconds the response to a signed URL request will be considered fresh. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceCdnPolicy.signed_url_cache_max_age_sec
	SignedUrlCacheMaxAgeSec *int64 `json:"signedUrlCacheMaxAgeSec,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.Duration
type BackendserviceConnectTimeout struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Duration.nanos
	Nanos *int32 `json:"nanos,omitempty"`
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Duration.seconds
	Seconds *int64 `json:"seconds"`
}

// +kcc:proto=google.cloud.compute.v1.CircuitBreakers
type BackendserviceCircuitBreakers struct {
	/* The timeout for new network connections to hosts. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CircuitBreakers.connect_timeout
	ConnectTimeout *BackendserviceConnectTimeout `json:"connectTimeout,omitempty"`

	/* The maximum number of connections to the backend cluster. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CircuitBreakers.max_connections
	MaxConnections *int32 `json:"maxConnections,omitempty"`

	/* The maximum number of pending requests to the backend cluster. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CircuitBreakers.max_pending_requests
	MaxPendingRequests *int32 `json:"maxPendingRequests,omitempty"`

	/* The maximum number of parallel requests to the backend cluster. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CircuitBreakers.max_requests
	MaxRequests *int32 `json:"maxRequests,omitempty"`

	/* Maximum requests for a single backend connection. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CircuitBreakers.max_requests_per_connection
	MaxRequestsPerConnection *int32 `json:"maxRequestsPerConnection,omitempty"`

	/* The maximum number of parallel retries to the backend cluster. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.CircuitBreakers.max_retries
	MaxRetries *int32 `json:"maxRetries,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceConnectionTrackingPolicy
type BackendserviceConnectionTrackingPolicy struct {
	/* Specifies connection persistence when backends are unhealthy. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceConnectionTrackingPolicy.connection_persistence_on_unhealthy_backends
	ConnectionPersistenceOnUnhealthyBackends *string `json:"connectionPersistenceOnUnhealthyBackends,omitempty"`

	/* Enable Strong Session Affinity for Network Load Balancing. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceConnectionTrackingPolicy.enable_strong_affinity
	EnableStrongAffinity *bool `json:"enableStrongAffinity,omitempty"`

	/* Specifies how long to keep a Connection Tracking entry while there is no matching traffic. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceConnectionTrackingPolicy.idle_timeout_sec
	IdleTimeoutSec *int32 `json:"idleTimeoutSec,omitempty"`

	/* Specifies the key used for connection tracking. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceConnectionTrackingPolicy.tracking_mode
	TrackingMode *string `json:"trackingMode,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.Duration
type BackendserviceTtl struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Duration.nanos
	Nanos *int32 `json:"nanos,omitempty"`
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Duration.seconds
	Seconds *int64 `json:"seconds"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceHTTPCookie
type BackendserviceHttpCookie struct {
	/* Name of the cookie. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceHTTPCookie.name
	Name *string `json:"name,omitempty"`

	/* Path to set for the cookie. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceHTTPCookie.path
	Path *string `json:"path,omitempty"`

	/* Lifetime of the cookie. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceHTTPCookie.ttl
	Ttl *BackendserviceTtl `json:"ttl,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ConsistentHashLoadBalancerSettings
type BackendserviceConsistentHash struct {
	/* Hash is based on HTTP Cookie. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.ConsistentHashLoadBalancerSettings.http_cookie
	HttpCookie *BackendserviceHttpCookie `json:"httpCookie,omitempty"`

	/* The hash based on the value of the specified header field. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.ConsistentHashLoadBalancerSettings.http_header_name
	HttpHeaderName *string `json:"httpHeaderName,omitempty"`

	/* The minimum number of virtual nodes to use for the hash ring. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.ConsistentHashLoadBalancerSettings.minimum_ring_size
	MinimumRingSize *int64 `json:"minimumRingSize,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceFailoverPolicy
type BackendserviceFailoverPolicy struct {
	/* On failover or failback, this field indicates whether connection drain will be honored. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceFailoverPolicy.disable_connection_drain_on_failover
	DisableConnectionDrainOnFailover *bool `json:"disableConnectionDrainOnFailover,omitempty"`

	/* Drop traffic if no healthy VMs are detected. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceFailoverPolicy.drop_traffic_if_unhealthy
	DropTrafficIfUnhealthy *bool `json:"dropTrafficIfUnhealthy,omitempty"`

	/* Failover ratio. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceFailoverPolicy.failover_ratio
	FailoverRatio *float32 `json:"failoverRatio,omitempty"`
}

type BackendserviceGroup struct {
	// +optional
	InstanceGroupRef *refsv1beta1.ComputeInstanceGroupRef `json:"instanceGroupRef,omitempty"`

	// +optional
	NetworkEndpointGroupRef *refsv1beta1.ComputeNetworkEndpointGroupRef `json:"networkEndpointGroupRef,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.Backend
type BackendserviceBackend struct {
	/* Specifies the balancing mode for this backend. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Backend.balancing_mode
	BalancingMode *string `json:"balancingMode,omitempty"`

	/* A multiplier applied to the group's maximum servicing capacity. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Backend.capacity_scaler
	CapacityScaler *float32 `json:"capacityScaler,omitempty"`

	/* An optional description of this resource. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Backend.description
	Description *string `json:"description,omitempty"`

	/* This field designates whether this is a failover backend. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Backend.failover
	Failover *bool `json:"failover,omitempty"`

	/* Reference to a ComputeInstanceGroup or ComputeNetworkEndpointGroup resource. */
	// +required
	// +kcc:proto:field=-
	BackendGroup BackendserviceGroup `json:"group"`

	/* The max number of simultaneous connections for the group. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_connections
	MaxConnections *int32 `json:"maxConnections,omitempty"`

	/* The max number of simultaneous connections that a single backend network endpoint can handle. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_connections_per_endpoint
	MaxConnectionsPerEndpoint *int32 `json:"maxConnectionsPerEndpoint,omitempty"`

	/* The max number of simultaneous connections that a single backend instance can handle. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_connections_per_instance
	MaxConnectionsPerInstance *int32 `json:"maxConnectionsPerInstance,omitempty"`

	/* The max requests per second (RPS) of the group. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_rate
	MaxRate *int32 `json:"maxRate,omitempty"`

	/* The max requests per second (RPS) that a single backend network endpoint can handle. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_rate_per_endpoint
	MaxRatePerEndpoint *float32 `json:"maxRatePerEndpoint,omitempty"`

	/* The max requests per second (RPS) that a single backend instance can handle. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_rate_per_instance
	MaxRatePerInstance *float32 `json:"maxRatePerInstance,omitempty"`

	/* CPU utilization target for the group. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Backend.max_utilization
	MaxUtilization *float32 `json:"maxUtilization,omitempty"`
}

type BackendserviceHealthChecks struct {
	// +optional
	HealthCheckRef *refsv1beta1.ComputeHealthCheckRef `json:"healthCheckRef,omitempty"`

	// +optional
	HttpHealthCheckRef *refsv1beta1.ComputeHTTPHealthCheckRef `json:"httpHealthCheckRef,omitempty"`
}

type BackendserviceValueFrom struct {
	/* Reference to a value with the given key in the given Secret. */
	// +optional
	SecretKeyRef *k8sv1alpha1.SecretKeyRef `json:"secretKeyRef,omitempty"`
}

type BackendserviceOauth2ClientSecret struct {
	/* Value of the field. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. */
	// +optional
	ValueFrom *BackendserviceValueFrom `json:"valueFrom,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceIAP
type BackendserviceIap struct {
	/* DEPRECATED. Although this field is still available, there is limited support. We recommend that you use `spec.iap.oauth2ClientIdRef` instead. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceIAP.oauth2_client_id
	Oauth2ClientId *string `json:"oauth2ClientId,omitempty"`

	/* OAuth2 Client ID for IAP. */
	// +optional
	Oauth2ClientIdRef *refsv1beta1.IAPIdentityAwareProxyClientRef `json:"oauth2ClientIdRef,omitempty"`

	/* OAuth2 Client Secret for IAP. */
	// +optional
	// +kcc:proto:field=-
	Oauth2ClientSecretRef *BackendserviceOauth2ClientSecret `json:"oauth2ClientSecret,omitempty"`

	/* OAuth2 Client Secret SHA-256 for IAP. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceIAP.oauth2_client_secret_sha256
	Oauth2ClientSecretSha256 *string `json:"oauth2ClientSecretSha256,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy
type BackendserviceCustomPolicy struct {
	/* An optional, arbitrary JSON object with configuration data. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy.data
	Data *string `json:"data,omitempty"`
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy.name
	Name *string `json:"name"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceLocalityLoadBalancingPolicyConfigPolicy
type BackendservicePolicy struct {
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceLocalityLoadBalancingPolicyConfigPolicy.name
	Name *string `json:"name"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceLocalityLoadBalancingPolicyConfig
type BackendserviceLocalityLbPolicies struct {
	/* The configuration for a custom policy. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceLocalityLoadBalancingPolicyConfig.custom_policy
	CustomPolicy *BackendserviceCustomPolicy `json:"customPolicy,omitempty"`

	/* The configuration for a built-in load balancing policy. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceLocalityLoadBalancingPolicyConfig.policy
	Policy *BackendservicePolicy `json:"policy,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.BackendServiceLogConfig
type BackendserviceLogConfig struct {
	/* Whether to enable logging for the load balancer traffic. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceLogConfig.enable
	Enable *bool `json:"enable,omitempty"`

	/* This configures the sampling rate of requests to the load balancer. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendServiceLogConfig.sample_rate
	SampleRate *float32 `json:"sampleRate,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.Duration
type BackendserviceBaseEjectionTime struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Duration.nanos
	Nanos *int32 `json:"nanos,omitempty"`
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Duration.seconds
	Seconds *int64 `json:"seconds"`
}

// +kcc:proto=google.cloud.compute.v1.Duration
type BackendserviceInterval struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Duration.nanos
	Nanos *int32 `json:"nanos,omitempty"`
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Duration.seconds
	Seconds *int64 `json:"seconds"`
}

// +kcc:proto=google.cloud.compute.v1.OutlierDetection
type BackendserviceOutlierDetection struct {
	/* The base time that a host is ejected for. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.base_ejection_time
	BaseEjectionTime *BackendserviceBaseEjectionTime `json:"baseEjectionTime,omitempty"`

	/* Number of errors before a host is ejected. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.consecutive_errors
	ConsecutiveErrors *int32 `json:"consecutiveErrors,omitempty"`

	/* Number of consecutive gateway failures. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.consecutive_gateway_failure
	ConsecutiveGatewayFailure *int32 `json:"consecutiveGatewayFailure,omitempty"`

	/* Percentage chance of ejection through consecutive 5xx. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.enforcing_consecutive_errors
	EnforcingConsecutiveErrors *int32 `json:"enforcingConsecutiveErrors,omitempty"`

	/* Percentage chance of ejection through consecutive gateway failures. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.enforcing_consecutive_gateway_failure
	EnforcingConsecutiveGatewayFailure *int32 `json:"enforcingConsecutiveGatewayFailure,omitempty"`

	/* Percentage chance of ejection through success rate statistics. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.enforcing_success_rate
	EnforcingSuccessRate *int32 `json:"enforcingSuccessRate,omitempty"`

	/* Time interval between ejection sweep analysis. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.interval
	Interval *BackendserviceInterval `json:"interval,omitempty"`

	/* Maximum percentage of hosts that can be ejected. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.max_ejection_percent
	MaxEjectionPercent *int32 `json:"maxEjectionPercent,omitempty"`

	/* Minimum number of hosts in a cluster. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.success_rate_minimum_hosts
	SuccessRateMinimumHosts *int32 `json:"successRateMinimumHosts,omitempty"`

	/* Minimum number of total requests. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.success_rate_request_volume
	SuccessRateRequestVolume *int32 `json:"successRateRequestVolume,omitempty"`

	/* Factor used to determine ejection threshold. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.OutlierDetection.success_rate_stdev_factor
	SuccessRateStdevFactor *int32 `json:"successRateStdevFactor,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SecuritySettings
type BackendserviceSecuritySettings struct {
	/* ClientTlsPolicy is a resource that specifies how a client should authenticate. */
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.SecuritySettings.client_tls_policy
	ClientTLSPolicyRef *refsv1beta1.NetworkSecurityClientTLSPolicyRef `json:"clientTLSPolicyRef"`

	/* Alternate names to verify subject identity. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.SecuritySettings.subject_alt_names
	SubjectAltNames []string `json:"subjectAltNames,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.Subsetting
type BackendserviceSubsetting struct {
	/* The algorithm used for subsetting. */
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Subsetting.policy
	Policy *string `json:"policy"`
}

// ComputeBackendServiceSpec defines the desired state of ComputeBackendService
// +kcc:spec:proto=google.cloud.compute.v1.BackendService
type ComputeBackendServiceSpec struct {
	/* Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.affinity_cookie_ttl_sec
	AffinityCookieTtlSec *int32 `json:"affinityCookieTtlSec,omitempty"`

	/* The set of backends that serve this BackendService. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.backends
	Backend []BackendserviceBackend `json:"backend,omitempty"`

	/* Cloud CDN configuration for this BackendService. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.cdn_policy
	CdnPolicy *BackendserviceCdnPolicy `json:"cdnPolicy,omitempty"`

	/* Settings controlling the volume of connections to a backend service. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.circuit_breakers
	CircuitBreakers *BackendserviceCircuitBreakers `json:"circuitBreakers,omitempty"`

	/* Compress text responses using Brotli or gzip compression. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.compression_mode
	CompressionMode *string `json:"compressionMode,omitempty"`

	/* Time for which instance will be drained. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.connection_draining
	ConnectionDrainingTimeoutSec *int32 `json:"connectionDrainingTimeoutSec,omitempty"`

	/* Connection Tracking configuration for this BackendService. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.connection_tracking_policy
	ConnectionTrackingPolicy *BackendserviceConnectionTrackingPolicy `json:"connectionTrackingPolicy,omitempty"`

	/* Consistent Hash-based load balancing parameters. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.consistent_hash
	ConsistentHash *BackendserviceConsistentHash `json:"consistentHash,omitempty"`

	/* Headers that the HTTP/S load balancer should add to proxied requests. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.custom_request_headers
	CustomRequestHeaders []string `json:"customRequestHeaders,omitempty"`

	/* Headers that the HTTP/S load balancer should add to proxied responses. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.custom_response_headers
	CustomResponseHeaders []string `json:"customResponseHeaders,omitempty"`

	/* An optional description of this resource. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.description
	Description *string `json:"description,omitempty"`

	/* The resource URL for the edge security policy associated with this backend service. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.edge_security_policy
	EdgeSecurityPolicyRef *ComputeSecurityPolicyRef `json:"edgeSecurityPolicyRef,omitempty"`

	/* If true, enable Cloud CDN for this BackendService. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.enable_c_d_n
	EnableCdn *bool `json:"enableCdn,omitempty"`

	/* Policy for failovers. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.failover_policy
	FailoverPolicy *BackendserviceFailoverPolicy `json:"failoverPolicy,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.health_checks
	HealthChecks []BackendserviceHealthChecks `json:"healthChecks,omitempty"`

	/* Settings for enabling Cloud Identity Aware Proxy. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.iap
	Iap *BackendserviceIap `json:"iap,omitempty"`

	/* Immutable. Indicates whether the backend service will be used with internal or external load balancing. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.load_balancing_scheme
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`

	/* A list of locality load balancing policies. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.locality_lb_policies
	LocalityLbPolicies []BackendserviceLocalityLbPolicies `json:"localityLbPolicies,omitempty"`

	/* The load balancing algorithm used within the scope of the locality. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.locality_lb_policy
	LocalityLbPolicy *string `json:"localityLbPolicy,omitempty"`

	/* Location represents the geographical location of the ComputeBackendService. Specify a region name or "global". */
	// +required
	Location string `json:"location"`

	/* Logging options. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.log_config
	LogConfig *BackendserviceLogConfig `json:"logConfig,omitempty"`

	/* The network to which this backend service belongs. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.network
	NetworkRef *ComputeNetworkRef `json:"networkRef,omitempty"`

	/* Settings controlling eviction of unhealthy hosts. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.outlier_detection
	OutlierDetection *BackendserviceOutlierDetection `json:"outlierDetection,omitempty"`

	/* Name of backend port. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.port_name
	PortName *string `json:"portName,omitempty"`

	/* The protocol this BackendService uses to communicate with backends. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.protocol
	Protocol *string `json:"protocol,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The security policy associated with this backend service. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.security_policy
	SecurityPolicy *string `json:"securityPolicy,omitempty"`

	/* The security policy associated with this backend service. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.security_policy
	SecurityPolicyRef *ComputeSecurityPolicyRef `json:"securityPolicyRef,omitempty"`

	/* The security settings that apply to this backend service. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.security_settings
	SecuritySettings *BackendserviceSecuritySettings `json:"securitySettings,omitempty"`

	/* Type of session affinity to use. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.session_affinity
	SessionAffinity *string `json:"sessionAffinity,omitempty"`

	/* Subsetting configuration for this BackendService. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.subsetting
	Subsetting *BackendserviceSubsetting `json:"subsetting,omitempty"`

	/* How many seconds to wait for the backend before considering it a failed request. */
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.BackendService.timeout_sec
	TimeoutSec *int32 `json:"timeoutSec,omitempty"`
}

// ComputeBackendServiceStatus defines the config connector machine state of ComputeBackendService
type ComputeBackendServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* A unique specifier for the ComputeBackendService resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *ComputeBackendServiceObservedState `json:"observedState,omitempty"`

	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* Fingerprint of this resource. A hash of the contents stored in this object. */
	// +optional
	Fingerprint *string `json:"fingerprint,omitempty"`

	/* The unique identifier for the resource. This identifier is defined by the server. */
	// +optional
	GeneratedId *int64 `json:"generatedId,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// ComputeBackendServiceObservedState is the state of the ComputeBackendService resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.BackendService
type ComputeBackendServiceObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputebackendservice;gcpcomputebackendservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeBackendService is the Schema for the ComputeBackendService API
// +k8s:openapi-gen=true
type ComputeBackendService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeBackendServiceSpec   `json:"spec,omitempty"`
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
