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
// proto.service: mockgcp.cloud.edgecacheservice.v1
// resource: NetworkServicesEdgeCacheService:EdgeCacheService

package v1alpha1

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.CdnPolicy
type CDNPolicy struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CdnPolicy.cache_mode
	CacheMode *string `json:"cacheMode,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CdnPolicy.client_ttl
	ClientTTL map[string]string `json:"clientTTL,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CdnPolicy.default_ttl
	DefaultTTL *string `json:"defaultTTL,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CdnPolicy.max_ttl
	MaxTTL *string `json:"maxTTL,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CdnPolicy.cache_negative_callbacks
	CacheNegativeCallbacks *bool `json:"cacheNegativeCallbacks,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CdnPolicy.signed_request_keyset
	SignedRequestKeyset map[string]string `json:"signedRequestKeyset,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CdnPolicy.signed_request_mode
	SignedRequestMode *string `json:"signedRequestMode,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.CorsPolicy
type CorsPolicy struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CorsPolicy.max_age
	MaxAge *string `json:"maxAge,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CorsPolicy.allow_origins
	AllowOrigins []string `json:"allowOrigins,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CorsPolicy.allow_methods
	AllowMethods []string `json:"allowMethods,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CorsPolicy.allow_headers
	AllowHeaders []string `json:"allowHeaders,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CorsPolicy.expose_headers
	ExposeHeaders []string `json:"exposeHeaders,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CorsPolicy.allow_credentials
	AllowCredentials *bool `json:"allowCredentials,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.CorsPolicy.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.HeaderAction
type HeaderAction struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HeaderAction.headers_to_remove
	HeadersToRemove []HeaderToRemove `json:"headersToRemove,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HeaderAction.headers_to_add
	HeadersToAdd []HeaderToAdd `json:"headersToAdd,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.HeaderMatch
type HeaderMatch struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HeaderMatch.header_name
	HeaderName *string `json:"headerName,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HeaderMatch.present_match
	PresentMatch *bool `json:"presentMatch,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HeaderMatch.exact_match
	ExactMatch *string `json:"exactMatch,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HeaderMatch.prefix_match
	PrefixMatch *string `json:"prefixMatch,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HeaderMatch.suffix_match
	SuffixMatch *string `json:"suffixMatch,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HeaderMatch.regex_match
	RegexMatch *string `json:"regexMatch,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HeaderMatch.invert_match
	InvertMatch *bool `json:"invertMatch,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.HeaderToAdd
type HeaderToAdd struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HeaderToAdd.header_name
	HeaderName *string `json:"headerName,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HeaderToAdd.header_value
	HeaderValue *string `json:"headerValue,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HeaderToAdd.replace
	Replace *bool `json:"replace,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.HeaderToRemove
type HeaderToRemove struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HeaderToRemove.header_name
	HeaderName *string `json:"headerName,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.HostRule
type HostRule struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HostRule.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HostRule.hosts
	Hosts []string `json:"hosts,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.HostRule.path_matcher
	PathMatcher *string `json:"pathMatcher,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.LogConfig
type LogConfig struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.LogConfig.enable
	Enable *bool `json:"enable,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.LogConfig.sample_rate
	SampleRate *float64 `json:"sampleRate,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.MatchRule
type MatchRule struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.MatchRule.prefix_match
	PrefixMatch *string `json:"prefixMatch,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.MatchRule.full_path_match
	FullPathMatch *string `json:"fullPathMatch,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.MatchRule.regex_match
	RegexMatch *string `json:"regexMatch,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.MatchRule.ignore_case
	IgnoreCase *bool `json:"ignoreCase,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.MatchRule.header_matches
	HeaderMatches []HeaderMatch `json:"headerMatches,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.MatchRule.query_param_matches
	QueryParamMatches []QueryParamMatch `json:"queryParamMatches,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.PathMatcher
type PathMatcher struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.PathMatcher.name
	Name *string `json:"name,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.PathMatcher.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.PathMatcher.route_rules
	RouteRules []RouteRule `json:"routeRules,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.QueryParamMatch
type QueryParamMatch struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.QueryParamMatch.query_param
	QueryParam *string `json:"queryParam,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.QueryParamMatch.present_match
	PresentMatch *bool `json:"presentMatch,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.QueryParamMatch.exact_match
	ExactMatch *string `json:"exactMatch,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.QueryParamMatch.regex_match
	RegexMatch *string `json:"regexMatch,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.RouteAction
type RouteAction struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.RouteAction.cdn_policy
	CDNPolicy *CDNPolicy `json:"cdnPolicy,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.RouteAction.url_rewrite
	URLRewrite *URLRewrite `json:"urlRewrite,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.RouteAction.cors_policy
	CorsPolicy *CorsPolicy `json:"corsPolicy,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.RouteAction.compression_mode
	CompressionMode *string `json:"compressionMode,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.RouteRule
type RouteRule struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.RouteRule.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.RouteRule.priority
	Priority *uint32 `json:"priority,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.RouteRule.match_rules
	MatchRules []MatchRule `json:"matchRules,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.RouteRule.header_action
	HeaderAction *HeaderAction `json:"headerAction,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.RouteRule.route_action
	RouteAction *RouteAction `json:"routeAction,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.RouteRule.origin
	Origin *string `json:"origin,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.RouteRule.url_redirect
	URLRedirect *URLRedirect `json:"urlRedirect,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.RouteRule.route_methods
	RouteMethods []string `json:"routeMethods,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.Routing
type Routing struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.Routing.host_rules
	HostRules []HostRule `json:"hostRules,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.Routing.path_matchers
	PathMatchers []PathMatcher `json:"pathMatchers,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.UrlRedirect
type URLRedirect struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.UrlRedirect.host_redirect
	HostRedirect *string `json:"hostRedirect,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.UrlRedirect.path_redirect
	PathRedirect *string `json:"pathRedirect,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.UrlRedirect.prefix_redirect
	PrefixRedirect *string `json:"prefixRedirect,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.UrlRedirect.redirect_response_code
	RedirectResponseCode *string `json:"redirectResponseCode,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.UrlRedirect.https_redirect
	HTTPSRedirect *bool `json:"httpsRedirect,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.UrlRedirect.strip_query
	StripQuery *bool `json:"stripQuery,omitempty"`
}

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.UrlRewrite
type URLRewrite struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.UrlRewrite.path_prefix_rewrite
	PathPrefixRewrite *string `json:"pathPrefixRewrite,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.UrlRewrite.host_rewrite
	HostRewrite *string `json:"hostRewrite,omitempty"`
}
